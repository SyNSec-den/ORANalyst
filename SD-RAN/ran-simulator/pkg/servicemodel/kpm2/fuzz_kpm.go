package kpm2

import (
	"context"
	"fmt"
	"net"

	e2smkpmv2sm "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/servicemodel"
	e2smkpmv2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"github.com/onosproject/ran-simulator/pkg/utils"
	e2apIndicationUtils "github.com/onosproject/ran-simulator/pkg/utils/e2ap/indication"
	"google.golang.org/protobuf/proto"
)

func (sm *Client) fuzzSendRicIndication(ctx context.Context) {
	// time.Sleep(10 * time.Minute)
	server, err := net.Listen("tcp", fmt.Sprintf("%s:%d", utils.SERVER_HOST, utils.KPM_SERVER_PORT))
	if err != nil {
		log.Fatalf("Error listening: %v", err)
	}
	defer server.Close()
	log.Infof("fuzzSendRicIndication: Listening on %s:%d", utils.SERVER_HOST, utils.KPM_SERVER_PORT)
	log.Infof("Starting fuzzing process for KPM Indication")

	for {
		sm.fuzzSendOneIndication(ctx, server)
	}
}

func (sm *Client) fuzzSendOneIndication(ctx context.Context, server net.Listener) {
	log.Infof("Sending KPM fuzzing indications")

	conn, err := server.Accept()
	if err != nil {
		log.Fatalf("Error accepting: %v", err)
	}
	defer conn.Close()

	indicationHeaderBytes, err := sm.createIndicationHeaderBytes(fileFormatVersion1)
	if err != nil {
		log.Errorf("Error creating indication header bytes: %v", err)
		return
	}

	indMsgBytes, err := utils.SocketRecv(conn)
	if err != nil {
		log.Errorf("Error receiving indication message bytes: %v", err)
		return
	}
	indMsg := &e2smkpmv2.E2SmKpmIndicationMessage{}
	err = proto.Unmarshal(indMsgBytes, indMsg)
	if err != nil {
		log.Errorf("Error unmarshalling E2SmKpmIndicationMessage: %v", err)
		return
	}

	var kpm2ServiceModel e2smkpmv2sm.Kpm2ServiceModel

	indicationMessageAsn1Bytes, err := kpm2ServiceModel.IndicationMessageProtoToASN1(indMsgBytes)
	if err != nil {
		log.Errorf("Error converting indication message from proto to asn1: %v", err)
		return
	}

	subs, err := sm.ServiceModel.Subscriptions.List()
	if err != nil {
		log.Fatalf("Error getting subscription: %v", err)
	}

	for _, sub := range subs {

		indication := e2apIndicationUtils.NewIndication(
			e2apIndicationUtils.WithRicInstanceID(sub.ReqID.GetRicInstanceId()),
			e2apIndicationUtils.WithRanFuncID(sub.FnID.GetValue()),
			e2apIndicationUtils.WithRequestID(sub.ReqID.GetRicRequestorId()),
			e2apIndicationUtils.WithIndicationHeader(indicationHeaderBytes),
			e2apIndicationUtils.WithIndicationMessage(indicationMessageAsn1Bytes))

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
