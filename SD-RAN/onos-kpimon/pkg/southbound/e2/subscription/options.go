// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package subscription

import (
	"github.com/onosproject/onos-kpimon/pkg/broker"
	appConfig "github.com/onosproject/onos-kpimon/pkg/config"
	"github.com/onosproject/onos-kpimon/pkg/monitoring"
	"github.com/onosproject/onos-kpimon/pkg/store/actions"
	"github.com/onosproject/onos-kpimon/pkg/store/measurements"
)

// Options E2 client options
type Options struct {
	E2TService E2TServiceOptions

	E2SubService E2SubServiceOptions

	ServiceModel ServiceModelOptions

	App AppOptions
}

// AppOptions application options
type AppOptions struct {
	AppID string

	AppConfig *appConfig.AppConfig

	Broker broker.Broker

	Monitor *monitoring.Monitor

	ActionStore actions.Store

	MeasurementStore measurements.Store
}

// E2TServiceOptions are the options for a E2T service
type E2TServiceOptions struct {
	// Host is the service host
	Host string
	// Port is the service port
	Port int
}

// E2SubServiceOptions are the options for E2sub service
type E2SubServiceOptions struct {
	// Host is the service host
	Host string
	// Port is the service port
	Port int
}

// ServiceModelName is a service model identifier
type ServiceModelName string

// ServiceModelVersion string
type ServiceModelVersion string

// ServiceModelOptions is options for defining a service model
type ServiceModelOptions struct {
	// Name is the service model identifier
	Name ServiceModelName

	// Version is the service model version
	Version ServiceModelVersion
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

// WithE2TAddress sets the address for the E2T service
func WithE2TAddress(host string, port int) Option {
	return newOption(func(options *Options) {
		options.E2TService.Host = host
		options.E2TService.Port = port
	})
}

// WithE2THost sets the host for the e2t service
func WithE2THost(host string) Option {
	return newOption(func(options *Options) {
		options.E2TService.Host = host
	})
}

// WithE2TPort sets the port for the e2t service
func WithE2TPort(port int) Option {
	return newOption(func(options *Options) {
		options.E2TService.Port = port
	})
}

// WithE2SubAddress sets the address for the E2Sub service
func WithE2SubAddress(host string, port int) Option {
	return newOption(func(options *Options) {
		options.E2SubService.Host = host
		options.E2SubService.Port = port
	})
}

// WithE2SubHost sets the host for the e2sub service
func WithE2SubHost(host string) Option {
	return newOption(func(options *Options) {
		options.E2SubService.Host = host
	})
}

// WithE2SubPort sets the port for the e2sub service
func WithE2SubPort(port int) Option {
	return newOption(func(options *Options) {
		options.E2SubService.Port = port
	})
}

// WithServiceModel sets the client service model
func WithServiceModel(name ServiceModelName, version ServiceModelVersion) Option {
	return newOption(func(options *Options) {
		options.ServiceModel = ServiceModelOptions{
			Name:    name,
			Version: version,
		}
	})
}

// WithAppID sets application ID
func WithAppID(appID string) Option {
	return newOption(func(options *Options) {
		options.App.AppID = appID
	})
}

// WithAppConfig sets the app config interface
func WithAppConfig(appConfig *appConfig.AppConfig) Option {
	return newOption(func(options *Options) {
		options.App.AppConfig = appConfig
	})
}

// WithBroker sets subscription broker
func WithBroker(broker broker.Broker) Option {
	return newOption(func(options *Options) {
		options.App.Broker = broker
	})
}

// WithActionStore sets actions store
func WithActionStore(actionStore actions.Store) Option {
	return newOption(func(options *Options) {
		options.App.ActionStore = actionStore
	})
}

// WithMeasurementStore sets measurement store
func WithMeasurementStore(measurementStore measurements.Store) Option {
	return newOption(func(options *Options) {
		options.App.MeasurementStore = measurementStore
	})
}
