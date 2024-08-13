package main

import (
	"fmt"
	"input-generator/utils"
	"net"
	"os"
	"syscall"
	"time"

	// . "github.com/dvyukov/go-fuzz/go-fuzz-defs"
	e2sm_mho "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/v2/e2sm-mho-go"
	"google.golang.org/protobuf/proto"
)

var (
	indir = "/app/corpus/"
)

func mho_sender() {
	mem := fakeMem()
	vals := make([]byte, 8*3) // res, ns, sonarPos

	files, err := os.ReadDir(indir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}
	// Map to hold filename and its corresponding binary content
	fileContents := make([][]byte, 0)

	for _, file := range files {
		if !file.IsDir() {
			filePath := indir + file.Name()
			content, err := os.ReadFile(filePath)
			if err != nil {
				fmt.Printf("Error reading file %s: %s\n", file.Name(), err)
				continue
			}
			fileContents = append(fileContents, content)
		}
	}

	// get coverage
	serverConn, err := net.Dial("tcp", "rimedo-ts.riab.svc.cluster.local:19999")
	_ = err
	if err != nil {
		panic(err)
	}
	defer serverConn.Close()
	fmt.Println("connected to coverage server")

	i := 0
	for {
		// startTime := time.Now()
		data := fileContents[i%len(fileContents)]
		if len(data) == 0 {
			fmt.Println("invalid input length")
			continue
		}

		indMsg := &e2sm_mho.E2SmMhoIndicationMessage{}
		err := proto.Unmarshal(data, indMsg)
		if err != nil {
			fmt.Printf("failed to unmarshal indMsg: %v\n", err)
			continue
		}

		asn1Ind, err := mhoMessageToAsn1Bytes(indMsg)
		if err != nil {
			fmt.Printf("failed to convert indMsg to asn1: %v\n", err)
			continue
		}

		_, err = mhoMessageToProtoBytes(asn1Ind)
		if err != nil {
			fmt.Printf("failed to convert indMsg to proto: %v\n", err)
			continue
		}

		fmt.Printf("found valid indMsg input: %+v\n", indMsg)
		// send input
		conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", utils.SERVER_HOST, utils.MHO_SERVER_PORT))
		if err != nil {
			panic(err)
		}

		defer conn.Close()
		fmt.Println("connected to ran sim server")
		_, err = conn.Write(data)
		if err != nil {
			fmt.Printf("failed to write input to conn: %v\n", err)
			syscall.Exit(1)
		}
		fmt.Printf("sent indMsg input: %+v\n", indMsg)

		realStartTime := time.Now()

		_, err = readAll(serverConn, make([]byte, 1))
		if err != nil {
			fmt.Printf("failed to read end signal from server: %v\n", err)
			syscall.Exit(1)
		}
		ns := time.Since(realStartTime)
		fmt.Printf("received end signal from server: %v\n", ns)

		_, err = readAll(serverConn, mem)
		if err != nil {
			fmt.Printf("failed to read coverage from server: %v\n", err)
			syscall.Exit(1)
		}
		fmt.Printf("received coverage from server: \n")

		valN, err := readAll(serverConn, vals)
		if err != nil {
			fmt.Printf("failed to read vals from server conn: %v\n", err)
			syscall.Exit(1)
		}
		if valN != 8*3 {
			fmt.Println("invalid vals length")
			syscall.Exit(1)
		}
		fmt.Printf("received vals from server: %v\n", vals)

		fmt.Printf("res: %v, ns: %v, time: %v, sonar: %v\n", 1, uint64(ns), ns, deserialize64(vals[16:24]))
	}
}
