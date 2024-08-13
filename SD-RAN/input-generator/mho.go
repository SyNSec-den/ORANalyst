package main

import (
	"crypto/sha1"
	"fmt"
	"input-generator/utils"
	"net"
	"syscall"
	"time"
	"unsafe"

	. "github.com/dvyukov/go-fuzz/go-fuzz-defs"
	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	e2smmhosm "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/servicemodel"
	e2sm_mho "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/v2/e2sm-mho-go"
	"google.golang.org/protobuf/proto"
)

type MhoIndicationBytes struct {
	HeaderBytes []byte
	IndBytes    []byte
}

type E2NodeIndication struct {
	NodeID      string
	TriggerType e2sm_mho.MhoTriggerType
	IndMsg      e2api.Indication
}

func mho() {
	mem, inFD, outFD := setupCommFile()

	input := mem[CoverSize : CoverSize+MaxInputSize]
	vals := make([]byte, 8*3) // res, ns, sonarPos
	CoverTab := (*[CoverSize]byte)(unsafe.Pointer(&mem[0]))

	// get coverage
	serverConn, err := net.Dial("tcp", "rimedo-ts.riab.svc.cluster.local:19999")
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
		CoverTab[genCounter(200000)]++
		if len(data) == 0 {
			write(outFD, uint64(time.Since(startTime)), 1, 10)
			CoverTab[genCounter(200001)]++
			continue
		}

		indMsg := &e2sm_mho.E2SmMhoIndicationMessage{}
		err := proto.Unmarshal(data, indMsg)
		CoverTab[genCounter(200002)]++
		if err != nil {
			write(outFD, uint64(time.Since(startTime)), 2, 10)
			CoverTab[genCounter(200003)]++
			continue
		}

		asn1Ind, err := mhoMessageToAsn1Bytes(indMsg)
		CoverTab[genCounter(200004)]++
		if err != nil {
			write(outFD, uint64(time.Since(startTime)), 3, 10)
			CoverTab[genCounter(200005)]++
			continue
		}

		_, err = mhoMessageToProtoBytes(asn1Ind)
		CoverTab[genCounter(200006)]++
		if err != nil {
			write(outFD, uint64(time.Since(startTime)), 4, 10)
			CoverTab[genCounter(200007)]++
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
		write(outFD, 1, uint64(ns), deserialize64(vals[16:24]))
		fmt.Printf("res: %v, ns: %v, time: %v, sonar: %v\n", 1, uint64(ns), ns, deserialize64(vals[16:24]))
	}
}

// mhoHeaderToAsn1Bytes converts e2sm mho header to asn1 bytes
func mhoHeaderToAsn1Bytes(header *e2sm_mho.E2SmMhoIndicationHeader) ([]byte, error) {
	indicationHeaderProtoBytes, err := proto.Marshal(header)
	if err != nil {
		return nil, err
	}

	var mhoServiceModel e2smmhosm.MhoServiceModel
	indicationHeaderAsn1Bytes, err := mhoServiceModel.IndicationHeaderProtoToASN1(indicationHeaderProtoBytes)

	if err != nil {
		return nil, err
	}
	return indicationHeaderAsn1Bytes, nil
}

// mhoMessageToAsn1Bytes converts e2sm mho message to asn1 bytes
func mhoMessageToAsn1Bytes(msg *e2sm_mho.E2SmMhoIndicationMessage) ([]byte, error) {
	indicationMessageProtoBytes, err := proto.Marshal(msg)
	if err != nil {
		return nil, err
	}

	var mhoServiceModel e2smmhosm.MhoServiceModel
	indicationMessageAsn1Bytes, err := mhoServiceModel.IndicationMessageProtoToASN1(indicationMessageProtoBytes)
	if err != nil {
		return nil, err
	}

	return indicationMessageAsn1Bytes, nil
}

// mhoHeaderToProtoBytes converts e2sm mho header to proto bytes
func mhoHeaderToProtoBytes(asn1Bytes []byte) ([]byte, error) {
	var mhoServiceModel e2smmhosm.MhoServiceModel
	return mhoServiceModel.IndicationHeaderASN1toProto(asn1Bytes)
}

// mhoMessageToProtoBytes converts e2sm mho message to proto bytes
func mhoMessageToProtoBytes(asn1Bytes []byte) ([]byte, error) {
	var mhoServiceModel e2smmhosm.MhoServiceModel
	return mhoServiceModel.IndicationMessageASN1toProto(asn1Bytes)
}

// func main() {
// 	message.Fuzz([]byte{})
// 	dir := "crashers"

// 	conn, err := utils.SocketConnect()
// 	if err != nil {
// 		log.Fatalf("SocketConnect: %v", err)
// 	}

// 	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
// 		if err != nil {
// 			log.Fatalf("filepath.Walk: %v", err)
// 		}

// 		if !info.IsDir() && !strings.Contains(path, ".") {
// 			time.Sleep(10 * time.Second)
// 			data, err := ioutil.ReadFile(path)
// 			if err != nil {
// 				log.Fatalf("ioutil.ReadFile: %v", err)
// 			}

// 			log.Printf("processing file: %s", path)

// 			var indMsg E2NodeIndication
// 			fuzz.NewFromGoFuzz(data).Fuzz(&indMsg)

// 			indHeaderByte := indMsg.IndMsg.Header
// 			indMessageByte := indMsg.IndMsg.Payload

// 			indHeader := e2sm_mho.E2SmMhoIndicationHeader{}
// 			if err := proto.Unmarshal(indHeaderByte, &indHeader); err != nil {
// 				log.Printf("Error: proto.Unmarshal indHeader: %v", err)
// 				return nil
// 			}

// 			indMessage := e2sm_mho.E2SmMhoIndicationMessage{}
// 			if err := proto.Unmarshal(indMessageByte, &indMessage); err != nil {
// 				log.Printf("Error: proto.Unmarshal indMessage: %v", err)
// 				return nil
// 			}

// 			indHeaderBytes, err := proto.Marshal(&indHeader)
// 			if err != nil {
// 				log.Fatalf("proto.Marshal indHeader: %v", err)
// 			}

// 			indMsgBytes, err := proto.Marshal(&indMessage)
// 			if err != nil {
// 				log.Fatalf("proto.Marshal indMsg: %v", err)
// 			}

// 			indBytes := &MhoIndicationBytes{
// 				HeaderBytes: indHeaderBytes,
// 				IndBytes:    indMsgBytes,
// 			}

// 			err = utils.GobMarshal(conn, indBytes)
// 			if err != nil {
// 				log.Fatalf("GobMarshal: %v", err)
// 			}

// 			log.Printf("Sent indHeader: %+v, indMessage: %+v", indHeader, indMessage)
// 		}
// 		return nil
// 	})

// 	if err != nil {
// 		log.Fatalf("filepath.Walk: %v", err)
// 	}

// 	for {
// 	}

// }

func genCounter(id uint32) int {
	buf := []byte{byte(id), byte(id >> 8), byte(id >> 16), byte(id >> 24)}
	hash := sha1.Sum(buf)
	return int(uint16(hash[0]) | uint16(hash[1])<<8)
}
