package mho

import (
	"context"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	e2smmhosm "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/servicemodel"
	e2sm_mho "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/v2/e2sm-mho-go"
	"github.com/onosproject/ran-simulator/pkg/utils"
	e2apIndicationUtils "github.com/onosproject/ran-simulator/pkg/utils/e2ap/indication"
	"google.golang.org/protobuf/proto"
)

type MhoIndicationBytes struct {
	HeaderBytes []byte
	IndBytes    []byte
}

func (m *Mho) fuzzSendOneIndication(ctx context.Context, server net.Listener) {
	log.Infof("Sending MHO fuzzing indications")
	log.Infof("Waiting for connection...")
	conn, err := server.Accept()
	if err != nil {
		log.Fatalf("Error accepting: %v", err)
	}
	defer conn.Close()
	log.Infof("fuzzSendOneIndication: Accepted connection from %v", conn.RemoteAddr())
	node := m.ServiceModel.Node
	// Creates and sends an indication message for each cell in the node
	var indicationHeaderBytes []byte
	for _, ncgi := range node.Cells {
		log.Infof("Send MHO indications for cell ncgi:%d", ncgi)
		indicationHeaderBytes, err = m.createIndicationHeaderBytes(ctx, ncgi)
		if err != nil {
			log.Errorf("Error creating indication header bytes: %v", err)
		} else {
			// writeToFile("mho_indication_header", indicationHeaderBytes)
			break
		}
	}
	if len(indicationHeaderBytes) == 0 {
		log.Errorf("Error creating indication header bytes: %v", err)
		return
	}

	indMsgBytes, err := utils.SocketRecv(conn)
	if err != nil {
		log.Errorf("Error receiving indication message bytes: %v", err)
		return
	}
	log.Infof("Received indication message bytes: %v", indMsgBytes)
	indMsg := &e2sm_mho.E2SmMhoIndicationMessage{}
	err = proto.Unmarshal(indMsgBytes, indMsg)
	if err != nil {
		log.Errorf("Error unmarshalling E2SmMhoIndicationMessage: %v", err)
		return
	}

	asn1Ind, err := mhoMessageToAsn1Bytes(indMsg)
	if err != nil {
		log.Errorf("Error converting E2SmMhoIndicationMessage to asn1 bytes: %v", err)
		return
	}

	subs, err := m.ServiceModel.Subscriptions.List()
	if err != nil {
		log.Fatalf("Error getting subscription: %v", err)
	}

	for _, sub := range subs {

		eventTrigger, err := getEventTriggerType(sub.Details.GetRicEventTriggerDefinition().Value)
		if err != nil {
			log.Errorf("Error getting event trigger type: %v", err)
		}
		log.Infof("Event trigger type: %v", eventTrigger)
		if eventTrigger != e2sm_mho.MhoTriggerType_MHO_TRIGGER_TYPE_UPON_CHANGE_RRC_STATUS {
			continue
		}

		indication := e2apIndicationUtils.NewIndication(
			e2apIndicationUtils.WithRicInstanceID(sub.ReqID.GetRicInstanceId()),
			e2apIndicationUtils.WithRanFuncID(sub.FnID.GetValue()),
			e2apIndicationUtils.WithRequestID(sub.ReqID.GetRicRequestorId()),
			e2apIndicationUtils.WithIndicationHeader(indicationHeaderBytes),
			e2apIndicationUtils.WithIndicationMessage(asn1Ind))

		ricIndication, err := indication.Build()
		if err != nil {
			log.Errorf("Error building indication: %v", err)
			return
		}

		err = sub.E2Channel.RICIndication(ctx, ricIndication)
		if err != nil {
			log.Errorf("Error sending indication: %v", err)
			return
		}

		log.Infof("Successfully sent mho indication to sub: %+v, msg: %+v",
			sub, indMsg)
	}
}

func writeToFile(dirPath string, data []byte) error {
	err := os.MkdirAll(dirPath, 0755)
	if err != nil {
		log.Fatalf("Failed to create directory: %v", err)
	}

	fileName := fmt.Sprintf("%v.bin", uuid.New().String())
	filePath := filepath.Join(dirPath, fileName)
	err = ioutil.WriteFile(filePath, data, 0644)
	if err != nil {
		log.Fatalf("Failed to write file: %v", err)
	}
	return nil
}

// getEventTriggerType extracts event trigger type
func getEventTriggerType(eventTriggerAsnBytes []byte) (e2sm_mho.MhoTriggerType, error) {
	var mhoServiceModel e2smmhosm.MhoServiceModel
	eventTriggerProtoBytes, err := mhoServiceModel.EventTriggerDefinitionASN1toProto(eventTriggerAsnBytes)
	if err != nil {
		return -1, err
	}
	eventTriggerDefinition := &e2sm_mho.E2SmMhoEventTriggerDefinition{}
	err = proto.Unmarshal(eventTriggerProtoBytes, eventTriggerDefinition)
	if err != nil {
		return -1, err
	}
	eventTriggerType := eventTriggerDefinition.GetEventDefinitionFormats().GetEventDefinitionFormat1().TriggerType
	return eventTriggerType, nil
}

func (m *Mho) fuzzSendRicIndication(ctx context.Context) {
	// time.Sleep(10 * time.Minute)
	server, err := net.Listen("tcp", fmt.Sprintf("%s:%d", utils.SERVER_HOST, utils.MHO_SERVER_PORT))
	if err != nil {
		log.Fatalf("Error listening: %v", err)
	}
	defer server.Close()
	log.Infof("fuzzSendRicIndication: Listening on %s:%d", utils.SERVER_HOST, utils.MHO_SERVER_PORT)
	log.Infof("Starting fuzzing process for MHO Indication")
	time.Sleep(10 * time.Second)

	for {
		m.fuzzSendOneIndication(ctx, server)
	}

	// for {
	// 	mhoIndicationBytes := &MhoIndicationBytes{}
	// 	err = utils.GobUnmarshal(conn, mhoIndicationBytes)
	// 	if err != nil {
	// 		log.Errorf("Error unmarshalling MhoIndicationBytes: %v", err)
	// 		continue
	// 	}
	// 	indHeader := &e2sm_mho.E2SmMhoIndicationHeader{}
	// 	err = proto.Unmarshal(mhoIndicationBytes.HeaderBytes, indHeader)
	// 	if err != nil {
	// 		log.Errorf("Error unmarshalling E2SmMhoIndicationHeader: %v", err)
	// 		continue
	// 	}
	// 	indMsg := &e2sm_mho.E2SmMhoIndicationMessage{}
	// 	err = proto.Unmarshal(mhoIndicationBytes.IndBytes, indMsg)
	// 	if err != nil {
	// 		log.Errorf("Error unmarshalling E2SmMhoIndicationMessage: %v", err)
	// 		continue
	// 	}

	// 	asn1Header, err := mhoHeaderToAsn1Bytes(indHeader)
	// 	if err != nil {
	// 		log.Errorf("Error converting E2SmMhoIndicationHeader to asn1 bytes: %v", err)
	// 		continue
	// 	}
	// 	asn1Ind, err := mhoMessageToAsn1Bytes(indMsg)
	// 	if err != nil {
	// 		log.Errorf("Error converting E2SmMhoIndicationMessage to asn1 bytes: %v", err)
	// 		continue
	// 	}

	// 	subs, err := m.ServiceModel.Subscriptions.List()
	// 	if err != nil {
	// 		log.Fatalf("Error getting subscription: %v", err)
	// 	}

	// 	for _, sub := range subs {

	// 		indication := e2apIndicationUtils.NewIndication(
	// 			e2apIndicationUtils.WithRicInstanceID(sub.ReqID.GetRicInstanceId()),
	// 			e2apIndicationUtils.WithRanFuncID(sub.FnID.GetValue()),
	// 			e2apIndicationUtils.WithRequestID(sub.ReqID.GetRicRequestorId()),
	// 			e2apIndicationUtils.WithIndicationHeader(asn1Header),
	// 			e2apIndicationUtils.WithIndicationMessage(asn1Ind))

	// 		ricIndication, err := indication.Build()
	// 		if err != nil {
	// 			log.Errorf("Error building indication: %v", err)
	// 			continue
	// 		}

	// 		err = sub.E2Channel.RICIndication(ctx, ricIndication)
	// 		if err != nil {
	// 			log.Errorf("Error sending indication: %v", err)
	// 			continue
	// 		}

	// 		log.Infof("Successfully sent mho indication to sub: %+v, header: %+v, msg: %+v",
	// 			sub, indHeader, indMsg)
	// 	}
	// }
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
