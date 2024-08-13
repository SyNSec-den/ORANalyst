// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/options.go:5
package topo

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/options.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/options.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/options.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/options.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/options.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/options.go:5
)

import (
	"fmt"

	topoapi "github.com/onosproject/onos-api/go/onos/topo"
)

const (
	// DefaultServicePort :
	DefaultServicePort	= 5150

	// DefaultServiceHost :
	DefaultServiceHost	= "onos-topo"
)

// Options topo SDK options
type Options struct {
	// Service is the service options
	Service ServiceOptions
}

// WithTopoAddress sets the address for the topo service
func WithTopoAddress(host string, port int) Option {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/options.go:28
	_go_fuzz_dep_.CoverTab[182970]++
													return newOption(func(options *Options) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/options.go:29
		_go_fuzz_dep_.CoverTab[182971]++
														options.Service.Host = host
														options.Service.Port = port
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/options.go:31
		// _ = "end of CoverTab[182971]"
	})
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/options.go:32
	// _ = "end of CoverTab[182970]"
}

// WithTopoHost sets the host for the topo service
func WithTopoHost(host string) Option {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/options.go:36
	_go_fuzz_dep_.CoverTab[182972]++
													return newOption(func(options *Options) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/options.go:37
		_go_fuzz_dep_.CoverTab[182973]++
														options.Service.Host = host
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/options.go:38
		// _ = "end of CoverTab[182973]"
	})
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/options.go:39
	// _ = "end of CoverTab[182972]"
}

// WithTopoPort sets the port for the topo service
func WithTopoPort(port int) Option {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/options.go:43
	_go_fuzz_dep_.CoverTab[182974]++
													return newOption(func(options *Options) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/options.go:44
		_go_fuzz_dep_.CoverTab[182975]++
														options.Service.Port = port
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/options.go:45
		// _ = "end of CoverTab[182975]"
	})
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/options.go:46
	// _ = "end of CoverTab[182974]"
}

// ServiceOptions are the options for a service
type ServiceOptions struct {
	// Host is the service host
	Host	string
	// Port is the service port
	Port	int

	Insecure	bool
}

// GetHost gets the service host
func (o ServiceOptions) GetHost() string {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/options.go:60
	_go_fuzz_dep_.CoverTab[182976]++
													return o.Host
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/options.go:61
	// _ = "end of CoverTab[182976]"
}

// GetPort gets the service port
func (o ServiceOptions) GetPort() int {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/options.go:65
	_go_fuzz_dep_.CoverTab[182977]++
													if o.Port == 0 {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/options.go:66
		_go_fuzz_dep_.CoverTab[182979]++
														return DefaultServicePort
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/options.go:67
		// _ = "end of CoverTab[182979]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/options.go:68
		_go_fuzz_dep_.CoverTab[182980]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/options.go:68
		// _ = "end of CoverTab[182980]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/options.go:68
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/options.go:68
	// _ = "end of CoverTab[182977]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/options.go:68
	_go_fuzz_dep_.CoverTab[182978]++
													return o.Port
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/options.go:69
	// _ = "end of CoverTab[182978]"
}

// IsInsecure is topo connection secure
func (o ServiceOptions) IsInsecure() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/options.go:73
	_go_fuzz_dep_.CoverTab[182981]++
													return o.Insecure
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/options.go:74
	// _ = "end of CoverTab[182981]"
}

// GetAddress gets the service address
func (o ServiceOptions) GetAddress() string {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/options.go:78
	_go_fuzz_dep_.CoverTab[182982]++
													return fmt.Sprintf("%s:%d", o.GetHost(), o.GetPort())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/options.go:79
	// _ = "end of CoverTab[182982]"
}

// Option topo client
type Option interface {
	apply(*Options)
}

type funcOption struct {
	f func(*Options)
}

func (f funcOption) apply(options *Options) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/options.go:91
	_go_fuzz_dep_.CoverTab[182983]++
													f.f(options)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/options.go:92
	// _ = "end of CoverTab[182983]"
}

func newOption(f func(*Options)) Option {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/options.go:95
	_go_fuzz_dep_.CoverTab[182984]++
													return funcOption{
		f: f,
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/options.go:98
	// _ = "end of CoverTab[182984]"
}

// WithOptions sets the client options
func WithOptions(opts Options) Option {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/options.go:102
	_go_fuzz_dep_.CoverTab[182985]++
														return newOption(func(options *Options) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/options.go:103
		_go_fuzz_dep_.CoverTab[182986]++
															*options = opts
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/options.go:104
		// _ = "end of CoverTab[182986]"
	})
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/options.go:105
	// _ = "end of CoverTab[182985]"
}

// WatchOptions topo client watch method options
type WatchOptions struct {
	filters		*topoapi.Filters
	noReplay	bool
}

// GetFilters get filters
func (w WatchOptions) GetFilters() *topoapi.Filters {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/options.go:115
	_go_fuzz_dep_.CoverTab[182987]++
														return w.filters
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/options.go:116
	// _ = "end of CoverTab[182987]"
}

// GetNoReplay gets no replay option
func (w WatchOptions) GetNoReplay() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/options.go:120
	_go_fuzz_dep_.CoverTab[182988]++
														return w.noReplay
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/options.go:121
	// _ = "end of CoverTab[182988]"
}

// WatchOption topo client watch option
type WatchOption interface {
	apply(*WatchOptions)
}

type funcWatchOption struct {
	f func(*WatchOptions)
}

func (f funcWatchOption) apply(options *WatchOptions) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/options.go:133
	_go_fuzz_dep_.CoverTab[182989]++
														f.f(options)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/options.go:134
	// _ = "end of CoverTab[182989]"
}

func newWatchOption(f func(*WatchOptions)) WatchOption {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/options.go:137
	_go_fuzz_dep_.CoverTab[182990]++
														return funcWatchOption{
		f: f,
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/options.go:140
	// _ = "end of CoverTab[182990]"
}

// WithWatchFilters sets filters for watch method
func WithWatchFilters(filters *topoapi.Filters) WatchOption {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/options.go:144
	_go_fuzz_dep_.CoverTab[182991]++
														return newWatchOption(func(o *WatchOptions) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/options.go:145
		_go_fuzz_dep_.CoverTab[182992]++
															o.filters = filters
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/options.go:146
		// _ = "end of CoverTab[182992]"
	})
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/options.go:147
	// _ = "end of CoverTab[182991]"
}

// WithNoReplay sets no replay option
func WithNoReplay(noReplay bool) WatchOption {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/options.go:151
	_go_fuzz_dep_.CoverTab[182993]++
														return newWatchOption(func(o *WatchOptions) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/options.go:152
		_go_fuzz_dep_.CoverTab[182994]++
															o.noReplay = noReplay
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/options.go:153
		// _ = "end of CoverTab[182994]"
	})
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/options.go:154
	// _ = "end of CoverTab[182993]"
}

// ListOptions topo client get method options
type ListOptions struct {
	filters *topoapi.Filters
}

// ListOption topo client list option
type ListOption interface {
	apply(*ListOptions)
}

type funcListOption struct {
	f func(options *ListOptions)
}

func (f funcListOption) apply(options *ListOptions) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/options.go:171
	_go_fuzz_dep_.CoverTab[182995]++
														f.f(options)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/options.go:172
	// _ = "end of CoverTab[182995]"
}

func newListOption(f func(options *ListOptions)) ListOption {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/options.go:175
	_go_fuzz_dep_.CoverTab[182996]++
														return funcListOption{
		f: f,
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/options.go:178
	// _ = "end of CoverTab[182996]"
}

// GetFilters get filters
func (l ListOptions) GetFilters() *topoapi.Filters {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/options.go:182
	_go_fuzz_dep_.CoverTab[182997]++
														return l.filters
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/options.go:183
	// _ = "end of CoverTab[182997]"
}

// WithListFilters sets filters for list method
func WithListFilters(filters *topoapi.Filters) ListOption {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/options.go:187
	_go_fuzz_dep_.CoverTab[182998]++
														return newListOption(func(o *ListOptions) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/options.go:188
		_go_fuzz_dep_.CoverTab[182999]++
															o.filters = filters
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/options.go:189
		// _ = "end of CoverTab[182999]"
	})
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/options.go:190
	// _ = "end of CoverTab[182998]"

}

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/options.go:192
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/options.go:192
var _ = _go_fuzz_dep_.CoverTab
