// SPDX-FileCopyrightText: 2022-present Intel Corporation
// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package subscription

import (
	"context"
	"strings"

	"github.com/onosproject/onos-kpimon/pkg/monitoring"
	"github.com/onosproject/onos-kpimon/pkg/store/actions"
	"github.com/onosproject/onos-kpimon/pkg/store/measurements"

	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"

	"github.com/onosproject/onos-kpimon/pkg/utils"

	"github.com/onosproject/onos-ric-sdk-go/pkg/config/event"

	"github.com/onosproject/onos-kpimon/pkg/broker"

	appConfig "github.com/onosproject/onos-kpimon/pkg/config"

	subutils "github.com/onosproject/onos-kpimon/pkg/utils/subscription"
	"github.com/onosproject/onos-lib-go/pkg/errors"

	prototypes "github.com/gogo/protobuf/types"
	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	"github.com/onosproject/onos-kpimon/pkg/rnib"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	e2client "github.com/onosproject/onos-ric-sdk-go/pkg/e2/v1beta1"
)

var log = logging.GetLogger()

const (
	kpmServiceModelOID = "1.3.6.1.4.1.53148.1.2.2.2"
)

// SubManager subscription manager interface
type SubManager interface {
	Start() error
	Stop() error
}

// Manager subscription manager
type Manager struct {
	e2client         e2client.Client
	rnibClient       rnib.Client
	serviceModel     ServiceModelOptions
	appConfig        *appConfig.AppConfig
	streams          broker.Broker
	actionStore      actions.Store
	measurementStore measurements.Store
}

// NewManager creates a new subscription manager
func NewManager(opts ...Option) (Manager, error) {
	options := Options{}

	for _, opt := range opts {
		opt.apply(&options)
	}

	serviceModelName := e2client.ServiceModelName(options.ServiceModel.Name)
	serviceModelVersion := e2client.ServiceModelVersion(options.ServiceModel.Version)
	appID := e2client.AppID(options.App.AppID)
	e2Client := e2client.NewClient(
		e2client.WithServiceModel(serviceModelName, serviceModelVersion),
		e2client.WithAppID(appID),
		e2client.WithE2TAddress(options.E2TService.Host, options.E2TService.Port))

	rnibClient, err := rnib.NewClient()
	if err != nil {
		return Manager{}, err
	}

	return Manager{
		e2client:   e2Client,
		rnibClient: rnibClient,
		serviceModel: ServiceModelOptions{
			Name:    options.ServiceModel.Name,
			Version: options.ServiceModel.Version,
		},
		appConfig:        options.App.AppConfig,
		streams:          options.App.Broker,
		actionStore:      options.App.ActionStore,
		measurementStore: options.App.MeasurementStore,
	}, nil

}

// Start starts subscription manager
func (m *Manager) Start() error {
	go func() {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		err := m.watchE2Connections(ctx)
		if err != nil {
			return
		}
	}()

	go func() {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		err := m.watchConfigChanges(ctx)
		if err != nil {
			return
		}
	}()
	return nil
}

func (m *Manager) watchConfigChanges(ctx context.Context) error {
	ch := make(chan event.Event)
	err := m.appConfig.Watch(ctx, ch)
	if err != nil {
		return err
	}

	// Deletes all of subscriptions
	for configEvent := range ch {
		log.Debugf("Config event is received: %v", configEvent)
		if configEvent.Key == utils.ReportPeriodConfigPath {
			channelIDs := m.streams.ChannelIDs()
			for _, channelID := range channelIDs {
				_, err := m.streams.CloseStream(ctx, channelID)
				if err != nil {
					log.Warn(err)
					return err
				}

			}
		}

	}
	// Gets all of connected E2 nodes and creates new subscriptions based on new report interval
	e2NodeIDs, err := m.rnibClient.E2NodeIDs(ctx, kpmServiceModelOID)
	if err != nil {
		log.Warn(err)
		return err
	}

	for _, e2NodeID := range e2NodeIDs {
		if !m.rnibClient.HasKPMRanFunction(ctx, e2NodeID, kpmServiceModelOID) {
			continue
		}
		go func(e2NodeID topoapi.ID) {
			err := m.newSubscription(ctx, e2NodeID)
			if err != nil {
				log.Warn(err)
			}
		}(e2NodeID)
	}

	return nil

}

func (m *Manager) getReportStyles(serviceModelsInfo map[string]*topoapi.ServiceModelInfo) ([]*topoapi.KPMReportStyle, error) {
	for _, sm := range serviceModelsInfo {
		smName := strings.ToLower(sm.Name)
		if smName == string(m.serviceModel.Name) && sm.OID == kpmServiceModelOID {
			kpmRanFunction := &topoapi.KPMRanFunction{}
			for _, ranFunction := range sm.RanFunctions {
				if ranFunction.TypeUrl == ranFunction.GetTypeUrl() {
					err := prototypes.UnmarshalAny(ranFunction, kpmRanFunction)
					if err != nil {
						return nil, err
					}
					return kpmRanFunction.ReportStyles, nil
				}
			}
		}
	}
	return nil, errors.New(errors.NotFound, "cannot retrieve report styles")
}

func (m *Manager) sendIndicationOnStream(streamID broker.StreamID, ch chan e2api.Indication) {
	streamWriter, err := m.streams.GetWriter(streamID)
	if err != nil {
		return
	}

	for msg := range ch {
		err := streamWriter.Send(msg)
		if err != nil {
			log.Warn(err)
			return
		}
	}
}

func (m *Manager) createSubscription(ctx context.Context, e2nodeID topoapi.ID) error {
	log.Info("Creating subscription for E2 node with ID:", e2nodeID)
	aspects, err := m.rnibClient.GetE2NodeAspects(ctx, e2nodeID)
	if err != nil {
		log.Warn(err)
		return err
	}
	reportStyles, err := m.getReportStyles(aspects.ServiceModels)
	if err != nil {
		log.Warn(err)
		return err
	}

	cells, err := m.rnibClient.GetCells(ctx, e2nodeID)
	if err != nil {
		log.Warn(err)
		return err
	}

	reportPeriod, err := m.appConfig.GetReportPeriod()
	if err != nil {
		log.Warn(err)
		return err
	}
	log.Debugf("Report period: %d", reportPeriod)
	eventTriggerData, err := subutils.CreateEventTriggerData(int64(reportPeriod))
	if err != nil {
		log.Warn(err)
		return err
	}

	granularityPeriod, err := m.appConfig.GetGranularityPeriod()
	if err != nil {
		log.Warn(err)
		return err
	}

	log.Debugf("Report styles:%v", reportStyles)
	// TODO we should check if for each report style a subscription should be created or for all of them
	for _, reportStyle := range reportStyles {
		actions, err := m.createSubscriptionActions(ctx, reportStyle, cells, int64(granularityPeriod))
		if err != nil {
			log.Warn(err)
			return err
		}
		measurements := reportStyle.Measurements

		ch := make(chan e2api.Indication)
		node := m.e2client.Node(e2client.NodeID(e2nodeID))
		subName := "onos-kpimon-subscription"

		subSpec := e2api.SubscriptionSpec{
			Actions: actions,
			EventTrigger: e2api.EventTrigger{
				Payload: eventTriggerData,
			},
		}

		channelID, err := node.Subscribe(ctx, subName, subSpec, ch)
		if err != nil {
			return err
		}

		log.Debugf("Channel ID:%s", channelID)
		streamReader, err := m.streams.OpenReader(ctx, node, subName, channelID, subSpec)
		if err != nil {
			return err
		}

		go m.sendIndicationOnStream(streamReader.StreamID(), ch)
		monitor := monitoring.NewMonitor(monitoring.WithAppConfig(m.appConfig),
			monitoring.WithActionStore(m.actionStore),
			monitoring.WithMeasurements(measurements),
			monitoring.WithNode(node),
			monitoring.WithStreamReader(streamReader),
			monitoring.WithNodeID(e2nodeID),
			monitoring.WithMeasurementStore(m.measurementStore),
			monitoring.WithRNIBClient(m.rnibClient))
		err = monitor.Start(ctx)
		if err != nil {
			log.Warn(err)
		}

	}

	return nil

}

func (m *Manager) newSubscription(ctx context.Context, e2NodeID topoapi.ID) error {
	err := m.createSubscription(ctx, e2NodeID)
	return err
}

func (m *Manager) watchE2Connections(ctx context.Context) error {
	ch := make(chan topoapi.Event)
	err := m.rnibClient.WatchE2Connections(ctx, ch)
	if err != nil {
		log.Warn(err)
		return err
	}

	// creates a new subscription whenever there is a new E2 node connected and supports KPM service model
	for topoEvent := range ch {
		log.Debugf("Received topo event: %v", topoEvent)

		if topoEvent.Type == topoapi.EventType_ADDED || topoEvent.Type == topoapi.EventType_NONE {
			relation := topoEvent.Object.Obj.(*topoapi.Object_Relation)
			e2NodeID := relation.Relation.TgtEntityID
			if !m.rnibClient.HasKPMRanFunction(ctx, e2NodeID, kpmServiceModelOID) {
				log.Debugf("Received topo event does not have KPM RAN function - %v", topoEvent)
				continue
			}

			go func(t topoapi.Event) {
				log.Debugf("start creating subscriptions %v", t)
				err := m.newSubscription(ctx, e2NodeID)
				if err != nil {
					log.Warn(err)
				}
			}(topoEvent)

		} else if topoEvent.Type == topoapi.EventType_REMOVED {
			relation := topoEvent.Object.Obj.(*topoapi.Object_Relation)
			e2NodeID := relation.Relation.TgtEntityID
			if !m.rnibClient.HasKPMRanFunction(ctx, e2NodeID, kpmServiceModelOID) {
				continue
			}
			cellIDs, err := m.rnibClient.GetCells(ctx, e2NodeID)
			if err != nil {
				return err
			}
			for _, coi := range cellIDs {
				key := measurements.Key{
					NodeID: string(e2NodeID),
					CellIdentity: measurements.CellIdentity{
						CellID: coi.CellObjectID,
					},
				}
				err = m.measurementStore.Delete(ctx, key)
				if err != nil {
					log.Warn(err)
				}
			}
		}

	}
	return nil
}

// Stop stops the subscription manager
func (m *Manager) Stop() error {
	panic("implement me")
}

var _ SubManager = &Manager{}
