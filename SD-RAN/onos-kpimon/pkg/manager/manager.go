// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package manager

import (
	"github.com/onosproject/onos-kpimon/pkg/broker"
	appConfig "github.com/onosproject/onos-kpimon/pkg/config"
	nbi "github.com/onosproject/onos-kpimon/pkg/northbound"
	"github.com/onosproject/onos-kpimon/pkg/southbound/e2/subscription"
	"github.com/onosproject/onos-kpimon/pkg/store/actions"
	"github.com/onosproject/onos-kpimon/pkg/store/measurements"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/onos-lib-go/pkg/northbound"
)

var log = logging.GetLogger()

// Config is a manager configuration
type Config struct {
	CAPath      string
	KeyPath     string
	CertPath    string
	E2tEndpoint string
	GRPCPort    int
	RicActionID int32
	ConfigPath  string
	SMName      string
	SMVersion   string
}

// NewManager generates the new KPIMON xAPP manager
func NewManager(config Config) *Manager {
	appCfg, err := appConfig.NewConfig(config.ConfigPath)
	if err != nil {
		log.Warn(err)
	}
	subscriptionBroker := broker.NewBroker()
	measStore := measurements.NewStore()
	actionsStore := actions.NewStore()

	subManager, err := subscription.NewManager(
		subscription.WithE2TAddress("onos-e2t", 5150),
		subscription.WithServiceModel(subscription.ServiceModelName(config.SMName),
			subscription.ServiceModelVersion(config.SMVersion)),
		subscription.WithAppConfig(appCfg),
		subscription.WithAppID("onos-kpimon"),
		subscription.WithBroker(subscriptionBroker),
		subscription.WithActionStore(actionsStore),
		subscription.WithMeasurementStore(measStore))

	if err != nil {
		log.Warn(err)
	}

	manager := &Manager{
		appConfig:        appCfg,
		config:           config,
		subManager:       subManager,
		measurementStore: measStore,
	}
	return manager
}

// Manager is an abstract struct for manager
type Manager struct {
	appConfig        appConfig.Config
	config           Config
	measurementStore measurements.Store
	subManager       subscription.Manager
}

// Run runs KPIMON manager
func (m *Manager) Run() {
	err := m.start()
	if err != nil {
		log.Errorf("Error when starting KPIMON: %v", err)
	}
}

// Close closes manager
func (m *Manager) Close() {
	log.Info("closing Manager")
}

func (m *Manager) start() error {
	err := m.startNorthboundServer()
	if err != nil {
		log.Warn(err)
		return err
	}

	err = m.subManager.Start()
	if err != nil {
		log.Warn(err)
		return err
	}

	return nil
}

func (m *Manager) startNorthboundServer() error {
	s := northbound.NewServer(northbound.NewServerCfg(
		m.config.CAPath,
		m.config.KeyPath,
		m.config.CertPath,
		int16(m.config.GRPCPort),
		true,
		northbound.SecurityConfig{}))

	s.AddService(nbi.NewService(m.measurementStore))

	doneCh := make(chan error)
	go func() {
		err := s.Serve(func(started string) {
			log.Info("Started NBI on ", started)
			close(doneCh)
		})
		if err != nil {
			doneCh <- err
		}
	}()
	return <-doneCh
}

// GetMeasurementStore returns measurement store
func (m *Manager) GetMeasurementStore() measurements.Store {
	return m.measurementStore
}
