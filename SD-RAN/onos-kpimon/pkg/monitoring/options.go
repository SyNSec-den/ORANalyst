// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package monitoring

import (
	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	"github.com/onosproject/onos-kpimon/pkg/broker"
	appConfig "github.com/onosproject/onos-kpimon/pkg/config"
	"github.com/onosproject/onos-kpimon/pkg/rnib"
	"github.com/onosproject/onos-kpimon/pkg/store/actions"
	measurmentStore "github.com/onosproject/onos-kpimon/pkg/store/measurements"
	e2client "github.com/onosproject/onos-ric-sdk-go/pkg/e2/v1beta1"
)

// Options monitor options
type Options struct {
	App AppOptions

	Monitor MonitorOptions
}

// AppOptions application options
type AppOptions struct {
	AppConfig *appConfig.AppConfig

	ActionStore actions.Store

	MeasurementStore measurmentStore.Store

	RNIBClient rnib.Client
}

// MonitorOptions monitoring options
type MonitorOptions struct {
	Node         e2client.Node
	Measurements []*topoapi.KPMMeasurement
	NodeID       topoapi.ID
	StreamReader broker.StreamReader
}

// Option option interface
type Option interface {
	apply(*Options)
}

type funcOption struct {
	f func(*Options)
}

func (f funcOption) apply(options *Options) {
	f.f(options)
}

func newOption(f func(*Options)) Option {
	return funcOption{
		f: f,
	}
}

// WithNodeID sets node ID
func WithNodeID(nodeID topoapi.ID) Option {
	return newOption(func(options *Options) {
		options.Monitor.NodeID = nodeID
	})
}

// WithStreamReader sets stream reader
func WithStreamReader(streamReader broker.StreamReader) Option {
	return newOption(func(options *Options) {
		options.Monitor.StreamReader = streamReader
	})
}

// WithMeasurements sets measurements
func WithMeasurements(measurements []*topoapi.KPMMeasurement) Option {
	return newOption(func(options *Options) {
		options.Monitor.Measurements = measurements
	})
}

// WithNode sets e2 node interface
func WithNode(node e2client.Node) Option {
	return newOption(func(options *Options) {
		options.Monitor.Node = node
	})
}

// WithMeasurementStore sets measurement store
func WithMeasurementStore(measurmentStore measurmentStore.Store) Option {
	return newOption(func(options *Options) {
		options.App.MeasurementStore = measurmentStore
	})
}

// WithAppConfig sets app config
func WithAppConfig(appConfig *appConfig.AppConfig) Option {
	return newOption(func(options *Options) {
		options.App.AppConfig = appConfig
	})
}

// WithActionStore sets actions store
func WithActionStore(actionStore actions.Store) Option {
	return newOption(func(options *Options) {
		options.App.ActionStore = actionStore
	})
}

// WithRNIBClient sets RNIB client
func WithRNIBClient(rnibClient rnib.Client) Option {
	return newOption(func(options *Options) {
		options.App.RNIBClient = rnibClient
	})
}
