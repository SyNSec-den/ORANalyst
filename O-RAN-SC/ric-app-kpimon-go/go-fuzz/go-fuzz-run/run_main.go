package main

import (
	"archive/zip"
	"encoding/binary"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

var CumulativeCover []byte

const (
	CoverSize       = 1 << 20
	MaxInputSize    = 1 << 20
	SonarRegionSize = 1 << 20

	BlockCoverSize = 1 << 19
	EdgeCoverSize  = 1 << 18

	SonarOpMask = 7
	SonarLength = 1 << 3
	SonarSigned = 1 << 4
	SonarString = 1 << 5
	SonarConst1 = 1 << 6
	SonarConst2 = 1 << 7

	SonarHdrLen = 6
	SonarMaxLen = 20
)

var coverConn net.Conn
var runCount int

type testee struct {
	coverRegion []byte
	inputRegion []byte
	sonarRegion []byte
	inPipe      *os.File
	outPipe     *os.File
	stdoutPipe  *os.File
	cmd         *exec.Cmd
	resbuf      [24]byte // reusable results buffer
}

func main() {
	CumulativeCover = make([]byte, CoverSize)
	runCount = 0
	filePtr := flag.String("file", "", "Path to the file")
	trackDir := flag.String("trackdir", "gopath/src/github.com/onosproject/rimedo-ts", "Path to the directory tracked by cover")
	wkDir := flag.String("wkdir", "/go/src/github.com/rimedo-labs/rimedo-ts", "Path to the working directory")
	coverFreq := flag.Int("coverfreq", 1, "Cover frequency")
	args := flag.String("args", "", "Arguments to the test binary")
	flag.Parse()

	trackedDir = *trackDir
	baseDir = *wkDir
	sourceDir = baseDir + "/instrument"
	destDir = baseDir + "/original"
	outputDir = baseDir + "/output"
	lastRunDir = baseDir + "/output/last-run"
	cumulativeDir = baseDir + "/output/cumulative"

	if *coverFreq > 0 {
		setupCover()
		go fileServer()
	}
	coverServer := setupCoverServer()
	go setupCoverConn(coverServer)
	zipfile := *filePtr
	destDir := "unzip"
	unzip(zipfile, destDir)
	coverBin := destDir + "/cover.exe"
	sonarBin := destDir + "/sonar.exe"
	_ = coverBin
	_ = sonarBin

	comm, err := os.Create("comm")
	if err != nil {
		log.Fatalf("failed to create comm file: %v", err)
	}
	comm.Truncate(CoverSize + MaxInputSize + SonarRegionSize)
	comm.Close()
	mappingFile, mem := createMapping(comm.Name(), CoverSize+MaxInputSize+SonarRegionSize)
	_ = mappingFile
	_ = mem

	t := startTestee(sonarBin, mappingFile, mem, *args)
	t.watchResult(mem, *coverFreq)
}

func setupCoverConn(server net.Listener) {
	for {
		var err error
		coverConn, err = server.Accept()
		if err != nil {
			panic(err)
		}
		fmt.Println("accept cover server")
	}
}

func setupCoverServer() net.Listener {
	coverServer, err := net.Listen("tcp", ":19999")
	if err != nil {
		panic(err)
	}
	fmt.Println("start cover server")
	return coverServer
}

func startTestee(sonarBin string, mappingFile *Mapping, mem []byte, args string) *testee {
retry:
	rIn, wIn, err := os.Pipe()
	if err != nil {
		log.Fatalf("failed to pipe: %v", err)
	}
	rOut, wOut, err := os.Pipe()
	if err != nil {
		log.Fatalf("failed to pipe: %v", err)
	}
	os.Chmod(sonarBin, 0755)
	cmd := exec.Command(sonarBin, strings.Split(args, " ")...)
	// For debugging of testee failures.
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	cmd.Env = append([]string{}, os.Environ()...)
	cmd.Env = append(cmd.Env, "GOTRACEBACK=1")
	setupCommMapping(cmd, mappingFile, rOut, wIn)
	fmt.Println("start sonar")
	if err = cmd.Start(); err != nil {
		// This can be a transient failure like "cannot allocate memory" or "text file is busy".
		log.Printf("failed to start test binary: %v", err)
		rIn.Close()
		wIn.Close()
		rOut.Close()
		wOut.Close()
		time.Sleep(time.Second)
		goto retry
	}
	rOut.Close()
	wIn.Close()

	t := &testee{
		coverRegion: mem[:CoverSize],
		inputRegion: mem[CoverSize : CoverSize+MaxInputSize],
		sonarRegion: mem[CoverSize+MaxInputSize:],
		inPipe:      rIn,
		outPipe:     wOut,
		stdoutPipe:  os.Stdout,
	}
	return t
}

func (t *testee) shutdown() {
	// t.cmd.Process.Kill()
	// t.inPipe.Close()
	// t.outPipe.Close()
	// t.stdoutPipe.Close()
}

func (t *testee) watchResult(mem []byte, coverFreq int) {
	fmt.Printf("watchResult: %v\n", coverFreq)
	defer t.shutdown()
	for {
		bbCover := 0
		edgeCover := 0
		fmt.Printf("runCount: %v\n", runCount)
		runCount++
		_, err := io.ReadFull(t.inPipe, t.resbuf[:])
		fmt.Printf("res: %v, ns: %v, sonar: %v\n", deserialize64(t.resbuf[0:8]), deserialize64(t.resbuf[8:16]), deserialize64(t.resbuf[16:24]))
		if err != nil {
			fmt.Print(fmt.Sprintf("failed to read from pipe: %v", err))
		}
		for i, v := range CumulativeCover {
			if v < t.coverRegion[i] {
				CumulativeCover[i] = t.coverRegion[i]
			}
			if v > 0 && i < BlockCoverSize {
				bbCover++
			}
			if v > 0 && i >= BlockCoverSize && i < BlockCoverSize+EdgeCoverSize {
				edgeCover++
			}
		}
		fmt.Printf("bbCover: %v, edgeCover: %v\n", bbCover, edgeCover)
		if coverConn == nil {
			continue
		}
		err = writeAll(coverConn, []byte{1})
		if err != nil {
			fmt.Print(fmt.Sprintf("failed to write to cover server: %v", err))
		}

		if runCount != 1 {
			if coverFreq > 0 && runCount%coverFreq == 0 {
				displayCover(t.coverRegion)
			}
		}

		// fmt.Println("validating sonar")
		// parseSonarData(t.sonarRegion[:deserialize64(t.resbuf[16:24])])
		// fmt.Println("validating success")

		err = writeAll(coverConn, mem)
		if err != nil {
			fmt.Print(fmt.Sprintf("failed to write to cover server: %v", err))
		}
		err = writeAll(coverConn, t.resbuf[:])
		if err != nil {
			fmt.Print(fmt.Sprintf("failed to write to cover server: %v", err))
		}
		fmt.Println("write to cover")

		t.outPipe.Write([]byte{1, 2, 3, 4, 5, 6, 7, 8, 9})
	}
}

func makeCopy(data []byte) []byte {
	return append([]byte{}, data...)
}

func parseSonarData(sonar []byte) {
	fmt.Println("parseSonarData")
	sonar = makeCopy(sonar)
	for len(sonar) > SonarHdrLen {
		id := binary.LittleEndian.Uint32(sonar)
		flags := byte(id)
		id >>= 8
		n1 := sonar[4]
		n2 := sonar[5]
		sonar = sonar[SonarHdrLen:]
		if n1 > SonarMaxLen || n2 > SonarMaxLen || len(sonar) < int(n1)+int(n2) {
			log.Fatalf("corrupted sonar data: hdr=[%v/%v/%v] data=%v", flags, n1, n2, len(sonar))
		}
		v1 := makeCopy(sonar[:n1])
		v2 := makeCopy(sonar[n1 : n1+n2])
		sonar = sonar[n1+n2:]

		// Trim trailing 0x00 and 0xff bytes (we don't know exact size of operands).
		if flags&SonarString == 0 {
			for len(v1) > 0 || len(v2) > 0 {
				i := len(v1) - 1
				if len(v2) > len(v1) {
					i = len(v2) - 1
				}
				var c1, c2 byte
				if i < len(v1) {
					c1 = v1[i]
				}
				if i < len(v2) {
					c2 = v2[i]
				}
				if (c1 == 0 || c1 == 0xff) && (c2 == 0 || c2 == 0xff) {
					if i < len(v1) {
						v1 = v1[:i]
					}
					if i < len(v2) {
						v2 = v2[:i]
					}
					continue
				}
				break
			}
		}
	}
}

func unzip(zipfile string, destDir string) {
	r, err := zip.OpenReader(zipfile)
	if err != nil {
		log.Fatalf("failed to open zip file: %v", err)
	}
	defer r.Close()

	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			panic(err)
		}
		defer rc.Close()

		fpath := filepath.Join(destDir, f.Name)

		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
		} else {
			var dir string
			if lastIndex := strings.LastIndex(fpath, string(os.PathSeparator)); lastIndex > -1 {
				dir = fpath[:lastIndex]
				os.MkdirAll(dir, os.ModePerm)
			}

			outFile, err := os.OpenFile(fpath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
			if err != nil {
				log.Fatalf("failed to open file: %v", err)
			}

			_, err = io.Copy(outFile, rc)
			outFile.Close()

			if err != nil {
				log.Fatalf("failed to copy file: %v", err)
			}
		}
	}
}

func deserialize64(buf []byte) uint64 {
	_ = buf[7]
	return uint64(buf[0])<<0 |
		uint64(buf[1])<<8 |
		uint64(buf[2])<<16 |
		uint64(buf[3])<<24 |
		uint64(buf[4])<<32 |
		uint64(buf[5])<<40 |
		uint64(buf[6])<<48 |
		uint64(buf[7])<<56
}

func writeAll(c net.Conn, b []byte) (err error) {
	if len(b) == 0 || c == nil {
		return nil
	}
	totalBytesWritten := 0
	for totalBytesWritten < len(b) {
		n, err := c.Write(b[totalBytesWritten:])
		if err != nil {
			netError, ok := err.(net.Error)
			if ok && netError.Temporary() {
				// If it's a temporary error, we'll retry.
				continue
			} else {
				// If it's not a temporary error, we return with the error.
				return err
			}
		}
		totalBytesWritten += n
	}
	return nil
}
