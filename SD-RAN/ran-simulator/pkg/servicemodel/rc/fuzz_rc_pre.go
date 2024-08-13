package rc

import (
	"context"
	"fmt"
	"net"

	ransimtypes "github.com/onosproject/onos-api/go/onos/ransim/types"
	e2smrcpresm "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre_go/servicemodel"
	e2smrcpreies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre_go/v2/e2sm-rc-pre-v2-go"
	"github.com/onosproject/ran-simulator/pkg/utils"
	indicationutils "github.com/onosproject/ran-simulator/pkg/utils/e2ap/indication"
	rcindicationhdr "github.com/onosproject/ran-simulator/pkg/utils/e2sm/rc/indication/header"
	"google.golang.org/protobuf/proto"
)

func (sm *Client) fuzzSendRicIndication(ctx context.Context) {
	// time.Sleep(10 * time.Minute)
	server, err := net.Listen("tcp", fmt.Sprintf("%s:%d", utils.SERVER_HOST, utils.RC_PRE_SERVER_PORT))
	if err != nil {
		log.Fatalf("Error listening: %v", err)
	}
	defer server.Close()
	log.Infof("fuzzSendRicIndication: Listening on %s:%d", utils.SERVER_HOST, utils.RC_PRE_SERVER_PORT)
	log.Infof("Starting fuzzing process for MHO Indication")

	for {
		sm.fuzzSendOneIndication(ctx, server)
	}
}

func (sm *Client) fuzzSendOneIndication(ctx context.Context, server net.Listener) {
	log.Infof("Sending RC pre fuzzing indications")

	conn, err := server.Accept()
	if err != nil {
		log.Fatalf("Error accepting: %v", err)
	}
	defer conn.Close()
	node := sm.ServiceModel.Node

	var indicationHeaderAsn1Bytes []byte
	for _, ncgi := range node.Cells {

		plmnID := sm.getPlmnID()
		cell, err := sm.ServiceModel.CellStore.Get(ctx, ncgi)
		if err != nil {
			log.Errorf("Error getting cell: %v", err)
		}

		cellEci := ransimtypes.GetNCI(cell.NCGI)
		// Creates RC indication header
		header := rcindicationhdr.NewIndicationHeader(
			rcindicationhdr.WithPlmnID(plmnID.Value()),
			rcindicationhdr.WithNRcellIdentity(uint64(cellEci)))

		indicationHeaderAsn1Bytes, err = header.ToAsn1Bytes()
		if err != nil {
			log.Errorf("Error converting RC indication header to asn1 bytes: %v", err)
			continue
		} else {
			break
		}
	}

	if len(indicationHeaderAsn1Bytes) == 0 {
		log.Errorf("Error creating indication header bytes: %v", err)
		return
	}

	indMsgBytes, err := utils.SocketRecv(conn)
	if err != nil {
		log.Errorf("Error receiving indication message bytes: %v", err)
		return
	}
	indMsg := &e2smrcpreies.E2SmRcPreIndicationMessage{}
	err = proto.Unmarshal(indMsgBytes, indMsg)
	if err != nil {
		log.Errorf("Error unmarshalling E2SmRcPreIndicationMessage: %v", err)
		return
	}

	indicationMessageProtoBytes, err := proto.Marshal(indMsg)
	if err != nil {
		log.Errorf("Error marshalling E2SmRcPreIndicationMessage: %v", err)
		return
	}

	var rcPreServiceModel e2smrcpresm.RcPreServiceModel
	indicationMessageAsn1Bytes, err := rcPreServiceModel.IndicationMessageProtoToASN1(indicationMessageProtoBytes)
	if err != nil {
		log.Errorf("Error converting E2SmRcPreIndicationMessage to asn1 bytes: %v", err)
		return
	}

	subs, err := sm.ServiceModel.Subscriptions.List()
	if err != nil {
		log.Fatalf("Error getting subscription: %v", err)
	}

	for _, sub := range subs {

		eventTrigger, err := getEventTriggerType(sub.Details.GetRicEventTriggerDefinition().Value)
		if err != nil {
			log.Errorf("Error getting event trigger type: %v", err)
		}
		log.Infof("Event trigger type: %v", eventTrigger)
		if eventTrigger != e2smrcpreies.RcPreTriggerType_RC_PRE_TRIGGER_TYPE_UPON_CHANGE {
			continue
		}

		// Creates e2 indication
		indication := indicationutils.NewIndication(
			indicationutils.WithRicInstanceID(sub.ReqID.GetRicInstanceId()),
			indicationutils.WithRanFuncID(sub.FnID.GetValue()),
			indicationutils.WithRequestID(sub.ReqID.GetRicRequestorId()),
			indicationutils.WithIndicationHeader(indicationHeaderAsn1Bytes),
			indicationutils.WithIndicationMessage(indicationMessageAsn1Bytes))

		ricIndication, err := indication.Build()
		if err != nil {
			log.Errorf("creating indication message is failed: %v", err)
			return
		}

		err = sub.E2Channel.RICIndication(ctx, ricIndication)
		if err != nil {
			log.Errorf("Error sending indication: %v", err)
			return
		}

		log.Infof("Successfully sent rc indication to sub: %+v, msg: %+v",
			sub, indMsg)
	}
}

// getEventTriggerType extracts event trigger type
func getEventTriggerType(eventTriggerAsnBytes []byte) (e2smrcpreies.RcPreTriggerType, error) {
	var rcPreServiceModel e2smrcpresm.RcPreServiceModel
	eventTriggerProtoBytes, err := rcPreServiceModel.EventTriggerDefinitionASN1toProto(eventTriggerAsnBytes)
	if err != nil {
		return -1, err
	}
	eventTriggerDefinition := &e2smrcpreies.E2SmRcPreEventTriggerDefinition{}
	err = proto.Unmarshal(eventTriggerProtoBytes, eventTriggerDefinition)
	if err != nil {
		return -1, err
	}
	eventTriggerType := eventTriggerDefinition.GetEventDefinitionFormats().GetEventDefinitionFormat1().TriggerType
	return eventTriggerType, nil
}
