package msg

import (
	"crypto/sha1"
	"fmt"

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

func Fuzz(data []byte) int {
	if len(data) == 0 {
		return 0
	}

	indMsg := &e2sm_mho.E2SmMhoIndicationMessage{}
	err := proto.Unmarshal(data, indMsg)
	if err != nil {
		return 0
	}

	asn1Ind, err := mhoMessageToAsn1Bytes(indMsg)
	if err != nil {
		return 0
	}

	_, err = mhoMessageToProtoBytes(asn1Ind)
	if err != nil {
		return 0
	}

	fmt.Printf("found valid indMsg input: %+v\n", indMsg)
	return 1
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