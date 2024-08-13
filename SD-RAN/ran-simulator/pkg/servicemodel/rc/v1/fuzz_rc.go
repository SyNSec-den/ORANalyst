package v1

import (
	"context"
	"fmt"
	"net"

	e2smrcsm "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc/servicemodel"
	e2smrcies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc/v1/e2sm-rc-ies"
	"github.com/onosproject/ran-simulator/pkg/utils"
	indicationutils "github.com/onosproject/ran-simulator/pkg/utils/e2ap/indication"
	"github.com/onosproject/ran-simulator/pkg/utils/e2sm/rc/v1/indication/headers/format1"
	"google.golang.org/protobuf/proto"
)

func (sm *Client) fuzzSendRicIndication(ctx context.Context) {
	// time.Sleep(10 * time.Minute)
	server, err := net.Listen("tcp", fmt.Sprintf("%s:%d", utils.SERVER_HOST, utils.RC_SERVER_PORT))
	if err != nil {
		log.Fatalf("Error listening: %v", err)
	}
	defer server.Close()
	log.Infof("fuzzSendRicIndication: Listening on %s:%d", utils.SERVER_HOST, utils.RC_SERVER_PORT)
	log.Infof("Starting fuzzing process for RC Indication")

	for {
		sm.fuzzSendOneIndication(ctx, server)
	}
}

func (sm *Client) fuzzSendOneIndication(ctx context.Context, server net.Listener) {
	log.Infof("Sending RC fuzzing indications")

	conn, err := server.Accept()
	if err != nil {
		log.Fatalf("Error accepting: %v", err)
	}
	defer conn.Close()

	headerFormat1 := format1.NewIndicationHeader(format1.WithEventConditionID(1))
	indicationHeaderAsn1Bytes, err := headerFormat1.ToAsn1Bytes()
	if err != nil {
		log.Errorf("Error converting RC indication header to asn1 bytes: %v", err)
		return
	}

	indMsgBytes, err := utils.SocketRecv(conn)
	if err != nil {
		log.Errorf("Error receiving indication message bytes: %v", err)
		return
	}
	indMsg := &e2smrcies.E2SmRcIndicationMessage{}
	err = proto.Unmarshal(indMsgBytes, indMsg)
	if err != nil {
		log.Errorf("Error unmarshalling E2SmRcIndicationMessage: %v", err)
		return
	}

	indicationMessageProtoBytes, err := proto.Marshal(indMsg)
	if err != nil {
		log.Errorf("Error marshalling E2SmRcIndicationMessage: %v", err)
		return
	}

	var rcServiceModel e2smrcsm.RCServiceModel
	indicationMessageAsn1Bytes, err := rcServiceModel.IndicationMessageProtoToASN1(indicationMessageProtoBytes)
	if err != nil {
		log.Errorf("Error converting RC indication message to asn1 bytes: %v", err)
		return
	}

	subs, err := sm.ServiceModel.Subscriptions.List()
	if err != nil {
		log.Fatalf("Error getting subscription: %v", err)
	}

	for _, sub := range subs {
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
