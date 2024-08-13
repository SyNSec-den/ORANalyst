package main

// #cgo CFLAGS: -I/home/tianchang/Desktop/proj/oran-sc/oran-input-gen/kpm
// #cgo LDFLAGS: -L/home/tianchang/Desktop/proj/oran-sc/oran-input-gen/kpm/build -lkpm -lm
// #include "mutator/e2ap/mutator.h"
import "C"
import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"syscall"
	"time"
	"unsafe"

	. "github.com/dvyukov/go-fuzz/go-fuzz-defs"
	"github.com/google/uuid"
)

// func main() {
// 	e2ap := readFromFile("corpus/e2ap.bin")
// 	decodeRes := C.decode_msg_print(unsafe.Pointer(&e2ap[0]), C.ulong(len(e2ap)))
// 	if decodeRes != 0 {
// 		fmt.Printf("decode failed: %v\n", decodeRes)
// 		// write(outFD, uint64(time.Since(startTime)), 1, 10)
// 		// continue
// 	}
// }

func main() {
	mem, inFD, outFD := setupCommFile()

	input := mem[CoverSize : CoverSize+MaxInputSize]
	vals := make([]byte, 8*3) // res, ns, sonarPos
	CoverTab := (*[CoverSize]byte)(unsafe.Pointer(&mem[0]))

	// get coverage
	serverConn, err := net.Dial("tcp", "localhost:19999")
	_ = err
	if err != nil {
		panic(err)
	}
	defer serverConn.Close()
	fmt.Println("connected to coverage server")
	time.Sleep(1 * time.Second) // wait for communication setup

	for {
		startTime := time.Now()
		_, n := read(inFD)
		if n > uint64(10240) {
			fmt.Println("invalid input length")
			write(outFD, uint64(time.Since(startTime)), 1, 10)
			continue
		}

		for i := range CoverTab {
			CoverTab[i] = 0
		}
		data := input[:n:n]
		if len(data) == 0 {
			write(outFD, uint64(time.Since(startTime)), 1, 10)
			continue
		}

		// decodeRes := C.decode_msg_print(unsafe.Pointer(&data[0]), C.ulong(len(data)))
		// if decodeRes != 0 {
		// 	fmt.Printf("decode failed: %v\n", decodeRes)
		// 	write(outFD, uint64(time.Since(startTime)), 1, 10)
		// 	continue
		// }

		err = sendSocket(data)
		if err != nil {
			fmt.Printf("Error sending data: %v\n", err)
			// the link broke, wait for someone to fix it
			for {
			}
			write(outFD, uint64(time.Since(startTime)), 1, 10)
			panic(fmt.Sprintf("Error sending data: %v\n", err))
			continue
		}

		realStartTime := time.Now()

		_, err = readAll(serverConn, make([]byte, 1))
		if err != nil {
			fmt.Printf("failed to read end signal from server: %v\n", err)
			panic(fmt.Sprintf("failed to read end signal from server: %v\n", err))
			syscall.Exit(1)
		}
		ns := time.Since(realStartTime)

		_, err = readAll(serverConn, mem)
		if err != nil {
			fmt.Printf("failed to read coverage from server: %v\n", err)
			panic(fmt.Sprintf("failed to read coverage from server: %v\n", err))
			syscall.Exit(1)
		}

		// for _, v := range CoverTab {
		// 	if v != 0 {
		// 		fmt.Printf("%v", v)
		// 		// break
		// 	}
		// }
		valN, err := readAll(serverConn, vals)
		if err != nil {
			fmt.Printf("failed to read vals from server conn: %v\n", err)
			panic(fmt.Sprintf("failed to read vals from server conn: %v\n", err))
			syscall.Exit(1)
		}
		if valN != 8*3 {
			fmt.Println("invalid vals length")
			panic(fmt.Sprintf("invalid vals length"))
			syscall.Exit(1)
		}

		// time.Sleep(10 * time.Second)
		write(outFD, 1, uint64(ns), deserialize64(vals[16:24]))
		fmt.Printf("res: %v, ns: %v, time: %v, sonar: %v\n", 1, uint64(ns), ns, deserialize64(vals[16:24]))
	}

}

func processMutateData(mutateData []byte) {
	err := sendSocket(mutateData)
	if err != nil {
		fmt.Printf("Error sending data: %v\n", err)
	}
	fileName := fmt.Sprintf("bin/%v.bin", uuid.New().String())
	writeToFile(mutateData, fileName)
	fmt.Printf("Wrote mutated data to %s\n", fileName)
}

func listenSocket() {
	listener, err := net.Listen("tcp", "localhost:19961")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	defer listener.Close()

	connChan := make(chan net.Conn, 1)
	errChan := make(chan error, 1)

	go func() {
		conn, err := listener.Accept()
		if err != nil {
			errChan <- err
			return
		}
		connChan <- conn
	}()

	select {
	case conn := <-connChan:
		handleConnection(conn)
	case err := <-errChan:
		fmt.Printf("Error accepting connection: %v", err)
	case <-time.After(15 * time.Second):
		fmt.Printf("Timeout: No connection was made")
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	var buffer bytes.Buffer

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	done := make(chan struct{})

	go func() {
		_, err := io.Copy(&buffer, conn)
		if err != nil {
			fmt.Printf("Error reading from connection: %v", err)
		}
		close(done) // signal that reading is complete
	}()

	select {
	case <-done:
		// Read completed, proceed
	case <-ctx.Done():
		// The read has taken too long, log an error
		fmt.Println("Read from connection timed out")
	}

	data := buffer.Bytes()
	fmt.Printf("Received data: %v\n", data)
}

func readFromFile(filename string) []byte {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", filename, err)
	}
	return data
}

func writeToFile(data []byte, filename string) {
	f, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Error creating file %s: %v\n", filename, err)
	}
	defer f.Close()

	_, err = f.Write(data)
	if err != nil {
		fmt.Printf("Error writing to file %s: %v\n", filename, err)
	}
}

func mutateOnce(fileMap map[string][]byte, sizeLb, sizeUb int) []byte {
	for {
		for n, data := range fileMap {
			if strings.HasPrefix(n, "msg") {
				ret := mutateMessageFork(data)
				// for _, b := range ret {
				// 	fmt.Printf("%02x ", b)
				// }
				// fmt.Println()
				if len(ret) > 0 {
					if (sizeUb != 0) && (len(ret) > sizeUb || len(ret) < sizeLb) {
						continue
					}
					decodeRes := C.decode_msg_print(unsafe.Pointer(&ret[0]), C.ulong(len(ret)))
					if decodeRes == 0 {
						return ret
					}
				}
			}
		}

	}
}

func sendSocket(data []byte) error {
	serverAddr := "localhost:19960"
	var conn net.Conn
	var err error
	// for {
	// 	conn, err = net.Dial("tcp", serverAddr)
	// 	if err == nil {
	// 		break
	// 	}
	// }
	conn, err = net.Dial("tcp", serverAddr)
	if err != nil {
		return fmt.Errorf("Error connecting to server: %v", err)
	}
	defer conn.Close()

	// Send all the content
	_, err = conn.Write(data)
	if err != nil {
		return fmt.Errorf("Error sending data: %v", err)
	}
	return nil
}
