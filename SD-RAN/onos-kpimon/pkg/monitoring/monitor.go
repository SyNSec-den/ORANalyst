// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package monitoring

import (
	"context"
	"io"

	"github.com/onosproject/onos-kpimon/pkg/rnib"

	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	"github.com/onosproject/onos-kpimon/pkg/store/actions"

	topoapi "github.com/onosproject/onos-api/go/onos/topo"

	appConfig "github.com/onosproject/onos-kpimon/pkg/config"

	measurmentStore "github.com/onosproject/onos-kpimon/pkg/store/measurements"

	e2smkpmv2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"google.golang.org/protobuf/proto"

	"github.com/onosproject/onos-lib-go/pkg/logging"

	gofuzzdep "github.com/dvyukov/go-fuzz/go-fuzz-dep"
	"github.com/onosproject/onos-kpimon/pkg/broker"
)

var log = logging.GetLogger()

// NewMonitor creates a new indication monitor
func NewMonitor(opts ...Option) *Monitor {
	options := Options{}

	for _, opt := range opts {
		opt.apply(&options)
	}

	return &Monitor{
		appConfig:        options.App.AppConfig,
		measurementStore: options.App.MeasurementStore,
		actionStore:      options.App.ActionStore,
		streamReader:     options.Monitor.StreamReader,
		nodeID:           options.Monitor.NodeID,
		measurements:     options.Monitor.Measurements,
		rnibClient:       options.App.RNIBClient,
	}
}

// Monitor indication monitor
type Monitor struct {
	streamReader     broker.StreamReader
	measurementStore measurmentStore.Store
	actionStore      actions.Store
	appConfig        *appConfig.AppConfig
	measurements     []*topoapi.KPMMeasurement
	nodeID           topoapi.ID
	rnibClient       rnib.Client
}

func (m *Monitor) processIndicationFormat1(ctx context.Context, indication e2api.Indication,
	measurements []*topoapi.KPMMeasurement, nodeID topoapi.ID) error {
	indHeader := e2smkpmv2.E2SmKpmIndicationHeader{}
	err := proto.Unmarshal(indication.Header, &indHeader)
	if err != nil {
		log.Warn(err)
		return err
	}

	indMessage := e2smkpmv2.E2SmKpmIndicationMessage{}
	err = proto.Unmarshal(indication.Payload, &indMessage)
	if err != nil {
		log.Warn(err)
		return err
	}

	indHdrFormat1 := indHeader.GetIndicationHeaderFormats().GetIndicationHeaderFormat1()
	indMsgFormat1 := indMessage.GetIndicationMessageFormats().GetIndicationMessageFormat1()
	log.Debugf("Received indication header format 1 %v:", indHdrFormat1)
	log.Debugf("Received indication message format 1: %v", indMsgFormat1)

	startTime := getTimeStampFromHeader(indHdrFormat1)
	startTimeUnixNano := toUnixNano(int64(startTime))

	granularity, err := m.appConfig.GetGranularityPeriod()
	if err != nil {
		log.Warn(err)
		return err
	}

	var cid string
	if indMsgFormat1.GetCellObjId() == nil {
		// Use the actions store to find cell object Id based on sub ID in action definition
		key := actions.NewKey(actions.SubscriptionID{
			SubID: indMsgFormat1.GetSubscriptId().GetValue(),
		})

		response, err := m.actionStore.Get(ctx, key)
		if err != nil {
			return err
		}

		actionDefinition := response.Value.(*e2smkpmv2.E2SmKpmActionDefinitionFormat1)
		cid = actionDefinition.GetCellObjId().GetValue()

	} else {
		cid = indMsgFormat1.GetCellObjId().Value
	}

	measDataItems := indMsgFormat1.GetMeasData().GetValue()
	measInfoList := indMsgFormat1.GetMeasInfoList().GetValue()

	measItems := make([]measurmentStore.MeasurementItem, 0)
	for i, measDataItem := range measDataItems {
		meadDataRecords := measDataItem.GetMeasRecord().GetValue()
		measRecords := make([]measurmentStore.MeasurementRecord, 0)
		for j, measDataRecord := range meadDataRecords {
			var measValue interface{}
			switch val := measDataRecord.MeasurementRecordItem.(type) {
			case *e2smkpmv2.MeasurementRecordItem_Integer:
				measValue = val.Integer

			case *e2smkpmv2.MeasurementRecordItem_Real:
				measValue = val.Real

			case *e2smkpmv2.MeasurementRecordItem_NoValue:
				measValue = val.NoValue
			default:
				measValue = 0
			}

			timeStamp := uint64(startTimeUnixNano) + granularity*uint64(1000000)*uint64(i)
			if len(measInfoList) > j {
				if measInfoList[j].GetMeasType().GetMeasName().GetValue() != "" {
					measName := measInfoList[j].GetMeasType().GetMeasName().GetValue()
					measRecord := measurmentStore.MeasurementRecord{
						Timestamp:        timeStamp,
						MeasurementName:  measName,
						MeasurementValue: measValue,
					}
					measRecords = append(measRecords, measRecord)
				} else if measInfoList[j].GetMeasType().GetMeasId() != nil {
					measID := measInfoList[j].GetMeasType().GetMeasId().String()
					log.Debugf("Received meas ID in indication message:", measID)
					log.Debugf("List of measurements:", measurements)
					measName := getMeasurementName(measID, measurements)
					measRecord := measurmentStore.MeasurementRecord{
						Timestamp:        timeStamp,
						MeasurementName:  measName,
						MeasurementValue: measValue,
					}
					measRecords = append(measRecords, measRecord)
				}
			}
		}

		measItem := measurmentStore.MeasurementItem{
			MeasurementRecords: measRecords,
		}
		measItems = append(measItems, measItem)

	}

	cellID := measurmentStore.CellIdentity{
		CellID: cid,
	}

	measurementKey := measurmentStore.NewKey(cellID, string(nodeID))
	_, err = m.measurementStore.Put(ctx, measurementKey, measItems)
	if err != nil {
		log.Warn(err)
		return err
	}

	cellTopoID, err := m.rnibClient.GetCellTopoID(ctx, cellID.CellID, nodeID)
	if err != nil {
		return err
	}
	err = m.rnibClient.UpdateCellAspects(ctx, cellTopoID, measItems)
	if err != nil {
		return err
	}

	return nil
}

func (m *Monitor) processIndication(ctx context.Context, indication e2api.Indication,
	measurements []*topoapi.KPMMeasurement, nodeID topoapi.ID) error {
	err := m.processIndicationFormat1(ctx, indication, measurements, nodeID)
	if err != nil {
		log.Warn(err)
		return err
	}

	return nil
}

// Start start monitoring of indication messages for a given subscription ID
func (m *Monitor) Start(ctx context.Context) error {
	go func() {
		for {
			gofuzzdep.LoopPos()
			indMsg, err := m.streamReader.Recv(ctx)
			if err != nil {
				log.Warn(err)
				if err == context.Canceled || err == context.DeadlineExceeded || err == io.EOF {
					return
				}
			}
			err = m.processIndication(ctx, indMsg, m.measurements, m.nodeID)
			if err != nil {
				log.Warn(err)
			}
		}
	}()

	<-ctx.Done()
	return ctx.Err()
}
