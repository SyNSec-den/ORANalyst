// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/options.go:5
package e2

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/options.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/options.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/options.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/options.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/options.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/options.go:5
)

import (
	"fmt"
	"time"
)

const defaultServicePort = 5150

// Encoding :
type Encoding int

const (
	// ProtoEncoding protobuf
	ProtoEncoding	Encoding	= iota

	// ASN1Encoding asn1
	ASN1Encoding
)

// Option is an E2 client option
type Option interface {
	apply(*Options)
}

// SubscribeOption is an option for subscribe request
type SubscribeOption interface {
	apply(*SubscribeOptions)
}

// EmptyOption is an empty client option
type EmptyOption struct{}

func (EmptyOption) apply(*Options) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/options.go:38
	_go_fuzz_dep_.CoverTab[196598]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/options.go:38
	// _ = "end of CoverTab[196598]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/options.go:38
}

// Options is a set of E2 client options
type Options struct {
	// AppOptions are the options for the application
	App	AppOptions
	// ServiceMode is service model options
	ServiceModel	ServiceModelOptions
	// Service is the E2 termination service configuration
	Service	ServiceOptions
	// Topo is the topology service configuration
	Topo	ServiceOptions
	// Encoding is the default encoding
	Encoding	Encoding
}

// AppID is an application identifier
type AppID string

// InstanceID is an app instance identifier
type InstanceID string

// AppOptions are the options for the application
type AppOptions struct {
	// AppID is the application identifier
	AppID	AppID
	// InstanceID is the application instance identifier
	InstanceID	InstanceID
}

// ServiceOptions are the options for a service
type ServiceOptions struct {
	// Host is the service host
	Host	string
	// Port is the service port
	Port	int
}

// SubscribeOptions are the options for a subscription
type SubscribeOptions struct {
	// Port is the service port
	TransactionTimeout time.Duration
}

// GetHost gets the service host
func (o ServiceOptions) GetHost() string {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/options.go:83
	_go_fuzz_dep_.CoverTab[196599]++
														return o.Host
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/options.go:84
	// _ = "end of CoverTab[196599]"
}

// GetPort gets the service port
func (o ServiceOptions) GetPort() int {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/options.go:88
	_go_fuzz_dep_.CoverTab[196600]++
														if o.Port == 0 {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/options.go:89
		_go_fuzz_dep_.CoverTab[196602]++
															return defaultServicePort
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/options.go:90
		// _ = "end of CoverTab[196602]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/options.go:91
		_go_fuzz_dep_.CoverTab[196603]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/options.go:91
		// _ = "end of CoverTab[196603]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/options.go:91
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/options.go:91
	// _ = "end of CoverTab[196600]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/options.go:91
	_go_fuzz_dep_.CoverTab[196601]++
														return o.Port
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/options.go:92
	// _ = "end of CoverTab[196601]"
}

// GetAddress gets the service address
func (o ServiceOptions) GetAddress() string {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/options.go:96
	_go_fuzz_dep_.CoverTab[196604]++
														return fmt.Sprintf("%s:%d", o.GetHost(), o.GetPort())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/options.go:97
	// _ = "end of CoverTab[196604]"
}

// ServiceModelName is a service model identifier
type ServiceModelName string

// ServiceModelVersion string
type ServiceModelVersion string

// ServiceModelOptions is options for defining a service model
type ServiceModelOptions struct {
	// Name is the service model identifier
	Name	ServiceModelName

	// Version is the service model version
	Version	ServiceModelVersion
}

type funcSubscribeOption struct {
	f func(*SubscribeOptions)
}

func (f funcSubscribeOption) apply(options *SubscribeOptions) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/options.go:119
	_go_fuzz_dep_.CoverTab[196605]++
														f.f(options)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/options.go:120
	// _ = "end of CoverTab[196605]"
}

func newSubscribeOption(f func(*SubscribeOptions)) SubscribeOption {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/options.go:123
	_go_fuzz_dep_.CoverTab[196606]++
														return funcSubscribeOption{
		f: f,
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/options.go:126
	// _ = "end of CoverTab[196606]"
}

type funcOption struct {
	f func(*Options)
}

func (f funcOption) apply(options *Options) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/options.go:133
	_go_fuzz_dep_.CoverTab[196607]++
														f.f(options)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/options.go:134
	// _ = "end of CoverTab[196607]"
}

func newOption(f func(*Options)) Option {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/options.go:137
	_go_fuzz_dep_.CoverTab[196608]++
														return funcOption{
		f: f,
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/options.go:140
	// _ = "end of CoverTab[196608]"
}

// WithOptions sets the client options
func WithOptions(opts Options) Option {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/options.go:144
	_go_fuzz_dep_.CoverTab[196609]++
														return newOption(func(options *Options) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/options.go:145
		_go_fuzz_dep_.CoverTab[196610]++
															*options = opts
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/options.go:146
		// _ = "end of CoverTab[196610]"
	})
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/options.go:147
	// _ = "end of CoverTab[196609]"
}

// WithAppID sets the client application identifier
func WithAppID(appID AppID) Option {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/options.go:151
	_go_fuzz_dep_.CoverTab[196611]++
														return newOption(func(options *Options) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/options.go:152
		_go_fuzz_dep_.CoverTab[196612]++
															options.App.AppID = appID
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/options.go:153
		// _ = "end of CoverTab[196612]"
	})
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/options.go:154
	// _ = "end of CoverTab[196611]"
}

// WithInstanceID sets the client instance identifier
func WithInstanceID(instanceID InstanceID) Option {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/options.go:158
	_go_fuzz_dep_.CoverTab[196613]++
														return newOption(func(options *Options) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/options.go:159
		_go_fuzz_dep_.CoverTab[196614]++
															options.App.InstanceID = instanceID
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/options.go:160
		// _ = "end of CoverTab[196614]"
	})
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/options.go:161
	// _ = "end of CoverTab[196613]"
}

// WithServiceModel sets the client service model
func WithServiceModel(name ServiceModelName, version ServiceModelVersion) Option {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/options.go:165
	_go_fuzz_dep_.CoverTab[196615]++
														return newOption(func(options *Options) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/options.go:166
		_go_fuzz_dep_.CoverTab[196616]++
															options.ServiceModel = ServiceModelOptions{
			Name:		name,
			Version:	version,
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/options.go:170
		// _ = "end of CoverTab[196616]"
	})
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/options.go:171
	// _ = "end of CoverTab[196615]"
}

// WithEncoding sets the client encoding
func WithEncoding(encoding Encoding) Option {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/options.go:175
	_go_fuzz_dep_.CoverTab[196617]++
														return newOption(func(options *Options) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/options.go:176
		_go_fuzz_dep_.CoverTab[196618]++
															options.Encoding = encoding
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/options.go:177
		// _ = "end of CoverTab[196618]"
	})
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/options.go:178
	// _ = "end of CoverTab[196617]"
}

// WithProtoEncoding sets the client encoding to ProtoEncoding
func WithProtoEncoding() Option {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/options.go:182
	_go_fuzz_dep_.CoverTab[196619]++
														return WithEncoding(ProtoEncoding)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/options.go:183
	// _ = "end of CoverTab[196619]"
}

// WithASN1Encoding sets the client encoding to ASN1Encoding
func WithASN1Encoding() Option {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/options.go:187
	_go_fuzz_dep_.CoverTab[196620]++
														return WithEncoding(ASN1Encoding)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/options.go:188
	// _ = "end of CoverTab[196620]"
}

// WithE2TAddress sets the address for the E2T service
func WithE2TAddress(host string, port int) Option {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/options.go:192
	_go_fuzz_dep_.CoverTab[196621]++
														return newOption(func(options *Options) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/options.go:193
		_go_fuzz_dep_.CoverTab[196622]++
															options.Service.Host = host
															options.Service.Port = port
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/options.go:195
		// _ = "end of CoverTab[196622]"
	})
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/options.go:196
	// _ = "end of CoverTab[196621]"
}

// WithE2THost sets the host for the E2T service
func WithE2THost(host string) Option {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/options.go:200
	_go_fuzz_dep_.CoverTab[196623]++
														return newOption(func(options *Options) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/options.go:201
		_go_fuzz_dep_.CoverTab[196624]++
															options.Service.Host = host
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/options.go:202
		// _ = "end of CoverTab[196624]"
	})
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/options.go:203
	// _ = "end of CoverTab[196623]"
}

// WithE2TPort sets the port for the E2T service
func WithE2TPort(port int) Option {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/options.go:207
	_go_fuzz_dep_.CoverTab[196625]++
														return newOption(func(options *Options) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/options.go:208
		_go_fuzz_dep_.CoverTab[196626]++
															options.Service.Port = port
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/options.go:209
		// _ = "end of CoverTab[196626]"
	})
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/options.go:210
	// _ = "end of CoverTab[196625]"
}

// WithTransactionTimeout sets a timeout value for subscriptions
func WithTransactionTimeout(transactionTimeout time.Duration) SubscribeOption {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/options.go:214
	_go_fuzz_dep_.CoverTab[196627]++
														return newSubscribeOption(func(options *SubscribeOptions) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/options.go:215
		_go_fuzz_dep_.CoverTab[196628]++
															options.TransactionTimeout = transactionTimeout
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/options.go:216
		// _ = "end of CoverTab[196628]"
	})
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/options.go:217
	// _ = "end of CoverTab[196627]"
}

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/options.go:218
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/options.go:218
var _ = _go_fuzz_dep_.CoverTab
