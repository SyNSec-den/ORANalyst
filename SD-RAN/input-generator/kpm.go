package main

import (
	"fmt"
	"input-generator/utils"
	"net"
	"syscall"
	"time"
	"unsafe"

	. "github.com/dvyukov/go-fuzz/go-fuzz-defs"
	e2smkpmsm "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/servicemodel"
	e2sm_kpm "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"google.golang.org/protobuf/proto"
)

func kpm() {
	mem, inFD, outFD := setupCommFile()

	input := mem[CoverSize : CoverSize+MaxInputSize]
	vals := make([]byte, 8*3) // res, ns, sonarPos
	CoverTab := (*[CoverSize]byte)(unsafe.Pointer(&mem[0]))

	// get coverage
	serverConn, err := net.Dial("tcp", "onos-kpimon.riab.svc.cluster.local:19999")
	_ = err
	if err != nil {
		panic(err)
	}
	defer serverConn.Close()
	fmt.Println("connected to coverage server")

	for {
		_, n := read(inFD)
		if n > uint64(MaxInputSize) {
			fmt.Println("invalid input length")
			syscall.Exit(1)
		}
		startTime := time.Now()
		for i := range CoverTab {
			CoverTab[i] = 0
		}
		data := input[:n:n]
		CoverTab[genCounter(300000)]++
		if len(data) == 0 {
			write(outFD, uint64(time.Since(startTime)), 1, 10)
			CoverTab[genCounter(300001)]++
			continue
		}

		indMsg := &e2sm_kpm.E2SmKpmIndicationMessage{}
		err := proto.Unmarshal(data, indMsg)
		CoverTab[genCounter(300002)]++
		if err != nil {
			write(outFD, uint64(time.Since(startTime)), 2, 10)
			CoverTab[genCounter(300003)]++
			continue
		}

		asn1Ind, err := kpmMessageToAsn1Bytes(indMsg)
		CoverTab[genCounter(300004)]++
		if err != nil {
			write(outFD, uint64(time.Since(startTime)), 3, 10)
			CoverTab[genCounter(300005)]++
			continue
		}

		_, err = kpmMessageToProtoBytes(asn1Ind)
		CoverTab[genCounter(300006)]++
		if err != nil {
			write(outFD, uint64(time.Since(startTime)), 4, 10)
			CoverTab[genCounter(300007)]++
			continue
		}

		fmt.Printf("found valid indMsg input: %+v\n", indMsg)
		// send input
		conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", utils.SERVER_HOST, utils.KPM_SERVER_PORT))
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

		_, err = readAll(serverConn, mem)
		if err != nil {
			fmt.Printf("failed to read coverage from server: %v\n", err)
			syscall.Exit(1)
		}

		valN, err := readAll(serverConn, vals)
		if err != nil {
			fmt.Printf("failed to read vals from server conn: %v\n", err)
			syscall.Exit(1)
		}
		if valN != 8*3 {
			fmt.Println("invalid vals length")
			syscall.Exit(1)
		}

		// time.Sleep(10 * time.Second)
		write(outFD, 1, uint64(time.Since(startTime))+deserialize64(vals[8:16]), deserialize64(vals[16:24]))
		fmt.Printf("res: %v, ns: %v, sonar: %v\n", 1, uint64(time.Since(startTime))+deserialize64(vals[8:16]), deserialize64(vals[16:24]))
	}
}

// kpmHeaderToAsn1Bytes converts e2sm kpm header to asn1 bytes
func kpmHeaderToAsn1Bytes(header *e2sm_kpm.E2SmKpmIndicationHeader) ([]byte, error) {
	indicationHeaderProtoBytes, err := proto.Marshal(header)
	if err != nil {
		return nil, err
	}

	var kpmServiceModel e2smkpmsm.Kpm2ServiceModel
	indicationHeaderAsn1Bytes, err := kpmServiceModel.IndicationHeaderProtoToASN1(indicationHeaderProtoBytes)

	if err != nil {
		return nil, err
	}
	return indicationHeaderAsn1Bytes, nil
}

// kpmMessageToAsn1Bytes converts e2sm kpm message to asn1 bytes
func kpmMessageToAsn1Bytes(msg *e2sm_kpm.E2SmKpmIndicationMessage) ([]byte, error) {
	indicationMessageProtoBytes, err := proto.Marshal(msg)
	if err != nil {
		return nil, err
	}

	var kpmServiceModel e2smkpmsm.Kpm2ServiceModel
	indicationMessageAsn1Bytes, err := kpmServiceModel.IndicationMessageProtoToASN1(indicationMessageProtoBytes)
	if err != nil {
		return nil, err
	}

	return indicationMessageAsn1Bytes, nil
}

// kpmHeaderToProtoBytes converts e2sm kpm header to proto bytes
func kpmHeaderToProtoBytes(asn1Bytes []byte) ([]byte, error) {
	var kpmServiceModel e2smkpmsm.Kpm2ServiceModel
	return kpmServiceModel.IndicationHeaderASN1toProto(asn1Bytes)
}

// kpmMessageToProtoBytes converts e2sm kpm message to proto bytes
func kpmMessageToProtoBytes(asn1Bytes []byte) ([]byte, error) {
	var kpmServiceModel e2smkpmsm.Kpm2ServiceModel
	return kpmServiceModel.IndicationMessageASN1toProto(asn1Bytes)
}
