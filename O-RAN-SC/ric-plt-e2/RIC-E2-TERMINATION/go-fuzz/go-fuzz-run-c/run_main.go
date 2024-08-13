package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
	"unsafe"
)

const (
	CoverSize       = 1 << 20
	MaxInputSize    = 1 << 20
	SonarRegionSize = 1 << 20

	SonarOpMask = 7
	SonarLength = 1 << 3
	SonarSigned = 1 << 4
	SonarString = 1 << 5
	SonarConst1 = 1 << 6
	SonarConst2 = 1 << 7

	SonarHdrLen = 6
	SonarMaxLen = 20
)

const socketPath = "/tmp/e2.sock"

const (
	IPC_CREAT = 0100000
	IPC_RMID  = 0
)

var runCount int

type monitor struct {
	coverConn   net.Conn
	mem         []byte
	coverRegion []byte
	// inputRegion []byte
	// sonarRegion []byte
	resbuf    [24]byte // reusable results buffer
	coverData []byte

	history []byte
}

func main() {
	var coverConn net.Conn
	mem := make([]byte, CoverSize+MaxInputSize+SonarRegionSize)
	coverRegion := mem[:CoverSize]

	shmID, _, errno := syscall.Syscall(syscall.SYS_SHMGET, 0, CoverSize, 0666|IPC_CREAT)
	if shmID < 0 {
		log.Fatalf("SYS_SHMGET failed: %v", errno)
	}

	shmAddr, _, errno := syscall.Syscall(syscall.SYS_SHMAT, shmID, 0, 0)
	if shmAddr == 0 {
		log.Fatalf("SYS_SHMAT failed: %v", errno)
	}

	coverData := (*[CoverSize]byte)(unsafe.Pointer(shmAddr))[:]

	for i := range coverData {
		coverData[i] = 0
	}
	monitor := &monitor{
		mem:         mem,
		coverConn:   coverConn,
		coverRegion: coverRegion,
		coverData:   coverData,
		history:     make([]byte, CoverSize),
	}

	go monitor.setupCoverServer(coverConn)

	// setup socket for IPC
	os.Remove(socketPath)
	l, err := net.Listen("unix", socketPath)
	if err != nil {
		panic(fmt.Sprintf("failed to listen: %v", err))
	}
	defer l.Close()

	// startup
	shmIDStr := strconv.Itoa(int(shmID))
	os.Setenv("__AFL_SHM_ID", shmIDStr)
	cmd := exec.Command("./startup.sh")
	monitor.start(cmd)

	for {
		conn, err := l.Accept()
		if err != nil {
			panic(fmt.Sprintf("failed to accept: %v", err))
		}

		monitor.handleConnection(conn)
	}
}

func (m *monitor) handleConnection(conn net.Conn) {
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Println("read error:", err)
		return
	}

	message := string(buf[:n])
	log.Println("Received:", message)

	if message == "Processing Done" {
		m.collectCover()
		// Notify e2 that coverage has been captured and it can proceed
		conn.Write([]byte("Coverage Captured"))
	}

	conn.Close()
}

func (m *monitor) start(cmd *exec.Cmd) {
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		panic(err)
	}
}

func (m *monitor) run(cmd *exec.Cmd) error {
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	return err
}

func (m *monitor) collectCover() {
	// time.Sleep(1 * time.Second)
	var err error
	count := 0
	for i, v := range m.history {
		if v < m.coverRegion[i] {
			m.history[i] = m.coverRegion[i]
		}
		if v > 0 {
			count++
		}
	}
	fmt.Printf("edge cover: %d\n", count)

	if m.coverConn != nil {
		err = writeAll(m.coverConn, []byte{1})
		if err != nil {
			fmt.Printf("failed to write to cover server: %v", err)
			return
		}
	}

	for i, v := range m.coverData {
		if v > 0 {
			m.coverRegion[i] = v
		}
	}

	if m.coverConn != nil {
		for _, v := range m.mem {
			if v != 0 {
				fmt.Printf("found non-zero: %d\n", v)
				break
			}
		}
		fmt.Printf("sending cover\n")
		err = writeAll(m.coverConn, m.mem)
		if err != nil {
			fmt.Printf("failed to write to cover server: %v", err)
			return
		}
		err = writeAll(m.coverConn, m.resbuf[:])
		if err != nil {
			fmt.Printf("failed to write to cover server: %v", err)
			return
		}
		fmt.Println("write to cover")
	}

	for i := range m.coverData {
		m.coverData[i] = 0
	}
}

func (m *monitor) collectCoverLcov() {
	var err error
	if m.coverConn != nil {
		err = writeAll(m.coverConn, []byte{1})
		if err != nil {
			fmt.Printf("failed to write to cover server: %v", err)
			return
		}
	}

	captureCmd := exec.Command("lcov", "--capture", "--directory", ".",
		"--output-file", "coverage.info", "--rc", "lcov_branch_coverage=1")
	err = m.run(captureCmd)
	if err != nil {
		fmt.Printf("lcov capture error: %v", err)
		return
	}

	content, err := os.ReadFile("coverage.info")
	if err != nil {
		fmt.Printf("read coverage.info error: %v", err)
		return
	}
	contentString := string(content)

	zeroCmd := exec.Command("lcov", "--zerocounters", "--directory", ".")
	err = m.run(zeroCmd)
	if err != nil {
		fmt.Printf("lcov zero error: %v", err)
		return
	}

	err = m.parseCover(contentString)
	if err != nil {
		fmt.Printf("parse cover error: %v", err)
		return
	}

	if m.coverConn != nil {
		for _, v := range m.mem {
			if v != 0 {
				fmt.Printf("found non-zero: %d\n", v)
				break
			}
		}
		err = writeAll(m.coverConn, m.mem)
		if err != nil {
			fmt.Printf("failed to write to cover server: %v", err)
			return
		}
		err = writeAll(m.coverConn, m.resbuf[:])
		if err != nil {
			fmt.Printf("failed to write to cover server: %v", err)
			return
		}
		fmt.Println("write to cover")
	}
}

func (m *monitor) parseCover(cover string) error {
	daIdx, brdaIdx := 0, 0
	scanner := bufio.NewScanner(strings.NewReader(cover))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "DA:") { // DA:2,1
			fields := strings.Split(line, ",")
			if len(fields) != 2 {
				return fmt.Errorf("invalid DA line: %s", line)
			}
			hits := strings.TrimSpace(fields[1])
			hitsInt, err := strconv.Atoi(hits)
			if err != nil || hitsInt < 0 {
				return fmt.Errorf("invalid DA line: %s", line)
			}
			m.coverRegion[daIdx] = byte(hitsInt % 256)
			daIdx++
		}
	}

	scanner = bufio.NewScanner(strings.NewReader(cover))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "BRDA:") { // BRDA:2,1,1,1 / BRDA:2,1,1,-
			fields := strings.Split(line, ",")
			if len(fields) != 4 {
				return fmt.Errorf("invalid BRDA line: %s", line)
			}
			hits := strings.TrimSpace(fields[3])
			hitsInt, err := strconv.Atoi(hits)
			if err == nil && hitsInt < 0 {
				return fmt.Errorf("invalid BRDA line: %s", line)
			}
			if err != nil {
				if hits != "-" {
					return fmt.Errorf("invalid BRDA line: %s", line)
				}
				hitsInt = -1
			}
			hitsInt++
			m.coverRegion[daIdx+brdaIdx] = byte(hitsInt % 256)
			brdaIdx++
		}
	}
	fmt.Printf("daIdx: %d, brdaIdx: %d\n", daIdx, brdaIdx)
	return nil
}

func (m *monitor) setupCoverServer(coverConn net.Conn) {
	coverServer, err := net.Listen("tcp", ":19999")
	if err != nil {
		panic(err)
	}
	fmt.Println("start cover server")
	for {
		m.coverConn, err = coverServer.Accept()
		if err != nil {
			panic(err)
		}
		fmt.Println("accept cover server")
	}
}

func makeCopy(data []byte) []byte {
	return append([]byte{}, data...)
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
