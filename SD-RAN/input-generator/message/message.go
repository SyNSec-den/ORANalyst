package message

import (
	"fmt"
	"input-generator/utils"
	"net"
	"time"

	e2smmhosm "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/servicemodel"
	e2sm_mho "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/v2/e2sm-mho-go"
	"google.golang.org/protobuf/proto"
)

func Fuzz(data []byte) int {
	if len(data) == 0 {
		return -1
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

	defer time.Sleep(10 * time.Second)

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", utils.SERVER_HOST, utils.SERVER_PORT))
	if err != nil {
		panic(err)
	}

	defer conn.Close()
	_, err = conn.Write(data)
	if err != nil {
		panic(err)
	}

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
