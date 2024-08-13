//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:19
package base

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:19
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:19
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:19
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:19
)

import (
	"errors"
	"fmt"

	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/resolver"
)

var logger = grpclog.Component("balancer")

type baseBuilder struct {
	name		string
	pickerBuilder	PickerBuilder
	config		Config
}

func (bb *baseBuilder) Build(cc balancer.ClientConn, opt balancer.BuildOptions) balancer.Balancer {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:39
	_go_fuzz_dep_.CoverTab[67428]++
												bal := &baseBalancer{
		cc:		cc,
		pickerBuilder:	bb.pickerBuilder,

		subConns:	resolver.NewAddressMap(),
		scStates:	make(map[balancer.SubConn]connectivity.State),
		csEvltr:	&balancer.ConnectivityStateEvaluator{},
		config:		bb.config,
		state:		connectivity.Connecting,
	}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:53
	bal.picker = NewErrPicker(balancer.ErrNoSubConnAvailable)
												return bal
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:54
	// _ = "end of CoverTab[67428]"
}

func (bb *baseBuilder) Name() string {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:57
	_go_fuzz_dep_.CoverTab[67429]++
												return bb.name
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:58
	// _ = "end of CoverTab[67429]"
}

type baseBalancer struct {
	cc		balancer.ClientConn
	pickerBuilder	PickerBuilder

	csEvltr	*balancer.ConnectivityStateEvaluator
	state	connectivity.State

	subConns	*resolver.AddressMap
	scStates	map[balancer.SubConn]connectivity.State
	picker		balancer.Picker
	config		Config

	resolverErr	error	// the last error reported by the resolver; cleared on successful resolution
	connErr		error	// the last connection error; cleared upon leaving TransientFailure
}

func (b *baseBalancer) ResolverError(err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:77
	_go_fuzz_dep_.CoverTab[67430]++
												b.resolverErr = err
												if b.subConns.Len() == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:79
		_go_fuzz_dep_.CoverTab[67433]++
													b.state = connectivity.TransientFailure
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:80
		// _ = "end of CoverTab[67433]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:81
		_go_fuzz_dep_.CoverTab[67434]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:81
		// _ = "end of CoverTab[67434]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:81
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:81
	// _ = "end of CoverTab[67430]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:81
	_go_fuzz_dep_.CoverTab[67431]++

												if b.state != connectivity.TransientFailure {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:83
		_go_fuzz_dep_.CoverTab[67435]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:86
		return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:86
		// _ = "end of CoverTab[67435]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:87
		_go_fuzz_dep_.CoverTab[67436]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:87
		// _ = "end of CoverTab[67436]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:87
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:87
	// _ = "end of CoverTab[67431]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:87
	_go_fuzz_dep_.CoverTab[67432]++
												b.regeneratePicker()
												b.cc.UpdateState(balancer.State{
		ConnectivityState:	b.state,
		Picker:			b.picker,
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:92
	// _ = "end of CoverTab[67432]"
}

func (b *baseBalancer) UpdateClientConnState(s balancer.ClientConnState) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:95
	_go_fuzz_dep_.CoverTab[67437]++

												if logger.V(2) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:97
		_go_fuzz_dep_.CoverTab[67442]++
													logger.Info("base.baseBalancer: got new ClientConn state: ", s)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:98
		// _ = "end of CoverTab[67442]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:99
		_go_fuzz_dep_.CoverTab[67443]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:99
		// _ = "end of CoverTab[67443]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:99
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:99
	// _ = "end of CoverTab[67437]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:99
	_go_fuzz_dep_.CoverTab[67438]++

												b.resolverErr = nil

												addrsSet := resolver.NewAddressMap()
												for _, a := range s.ResolverState.Addresses {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:104
		_go_fuzz_dep_.CoverTab[67444]++
													addrsSet.Set(a, nil)
													if _, ok := b.subConns.Get(a); !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:106
			_go_fuzz_dep_.CoverTab[67445]++

														sc, err := b.cc.NewSubConn([]resolver.Address{a}, balancer.NewSubConnOptions{HealthCheckEnabled: b.config.HealthCheck})
														if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:109
				_go_fuzz_dep_.CoverTab[67447]++
															logger.Warningf("base.baseBalancer: failed to create new SubConn: %v", err)
															continue
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:111
				// _ = "end of CoverTab[67447]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:112
				_go_fuzz_dep_.CoverTab[67448]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:112
				// _ = "end of CoverTab[67448]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:112
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:112
			// _ = "end of CoverTab[67445]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:112
			_go_fuzz_dep_.CoverTab[67446]++
														b.subConns.Set(a, sc)
														b.scStates[sc] = connectivity.Idle
														b.csEvltr.RecordTransition(connectivity.Shutdown, connectivity.Idle)
														sc.Connect()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:116
			// _ = "end of CoverTab[67446]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:117
			_go_fuzz_dep_.CoverTab[67449]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:117
			// _ = "end of CoverTab[67449]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:117
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:117
		// _ = "end of CoverTab[67444]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:118
	// _ = "end of CoverTab[67438]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:118
	_go_fuzz_dep_.CoverTab[67439]++
												for _, a := range b.subConns.Keys() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:119
		_go_fuzz_dep_.CoverTab[67450]++
													sci, _ := b.subConns.Get(a)
													sc := sci.(balancer.SubConn)

													if _, ok := addrsSet.Get(a); !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:123
			_go_fuzz_dep_.CoverTab[67451]++
														b.cc.RemoveSubConn(sc)
														b.subConns.Delete(a)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:125
			// _ = "end of CoverTab[67451]"

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:128
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:128
			_go_fuzz_dep_.CoverTab[67452]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:128
			// _ = "end of CoverTab[67452]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:128
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:128
		// _ = "end of CoverTab[67450]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:129
	// _ = "end of CoverTab[67439]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:129
	_go_fuzz_dep_.CoverTab[67440]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:134
	if len(s.ResolverState.Addresses) == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:134
		_go_fuzz_dep_.CoverTab[67453]++
													b.ResolverError(errors.New("produced zero addresses"))
													return balancer.ErrBadResolverState
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:136
		// _ = "end of CoverTab[67453]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:137
		_go_fuzz_dep_.CoverTab[67454]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:137
		// _ = "end of CoverTab[67454]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:137
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:137
	// _ = "end of CoverTab[67440]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:137
	_go_fuzz_dep_.CoverTab[67441]++

												b.regeneratePicker()
												b.cc.UpdateState(balancer.State{ConnectivityState: b.state, Picker: b.picker})
												return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:141
	// _ = "end of CoverTab[67441]"
}

// mergeErrors builds an error from the last connection error and the last
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:144
// resolver error.  Must only be called if b.state is TransientFailure.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:146
func (b *baseBalancer) mergeErrors() error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:146
	_go_fuzz_dep_.CoverTab[67455]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:149
	if b.connErr == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:149
		_go_fuzz_dep_.CoverTab[67458]++
													return fmt.Errorf("last resolver error: %v", b.resolverErr)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:150
		// _ = "end of CoverTab[67458]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:151
		_go_fuzz_dep_.CoverTab[67459]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:151
		// _ = "end of CoverTab[67459]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:151
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:151
	// _ = "end of CoverTab[67455]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:151
	_go_fuzz_dep_.CoverTab[67456]++
												if b.resolverErr == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:152
		_go_fuzz_dep_.CoverTab[67460]++
													return fmt.Errorf("last connection error: %v", b.connErr)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:153
		// _ = "end of CoverTab[67460]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:154
		_go_fuzz_dep_.CoverTab[67461]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:154
		// _ = "end of CoverTab[67461]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:154
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:154
	// _ = "end of CoverTab[67456]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:154
	_go_fuzz_dep_.CoverTab[67457]++
												return fmt.Errorf("last connection error: %v; last resolver error: %v", b.connErr, b.resolverErr)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:155
	// _ = "end of CoverTab[67457]"
}

// regeneratePicker takes a snapshot of the balancer, and generates a picker
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:158
// from it. The picker is
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:158
//   - errPicker if the balancer is in TransientFailure,
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:158
//   - built by the pickerBuilder with all READY SubConns otherwise.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:162
func (b *baseBalancer) regeneratePicker() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:162
	_go_fuzz_dep_.CoverTab[67462]++
												if b.state == connectivity.TransientFailure {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:163
		_go_fuzz_dep_.CoverTab[67465]++
													b.picker = NewErrPicker(b.mergeErrors())
													return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:165
		// _ = "end of CoverTab[67465]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:166
		_go_fuzz_dep_.CoverTab[67466]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:166
		// _ = "end of CoverTab[67466]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:166
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:166
	// _ = "end of CoverTab[67462]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:166
	_go_fuzz_dep_.CoverTab[67463]++
												readySCs := make(map[balancer.SubConn]SubConnInfo)

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:170
	for _, addr := range b.subConns.Keys() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:170
		_go_fuzz_dep_.CoverTab[67467]++
													sci, _ := b.subConns.Get(addr)
													sc := sci.(balancer.SubConn)
													if st, ok := b.scStates[sc]; ok && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:173
			_go_fuzz_dep_.CoverTab[67468]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:173
			return st == connectivity.Ready
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:173
			// _ = "end of CoverTab[67468]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:173
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:173
			_go_fuzz_dep_.CoverTab[67469]++
														readySCs[sc] = SubConnInfo{Address: addr}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:174
			// _ = "end of CoverTab[67469]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:175
			_go_fuzz_dep_.CoverTab[67470]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:175
			// _ = "end of CoverTab[67470]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:175
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:175
		// _ = "end of CoverTab[67467]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:176
	// _ = "end of CoverTab[67463]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:176
	_go_fuzz_dep_.CoverTab[67464]++
												b.picker = b.pickerBuilder.Build(PickerBuildInfo{ReadySCs: readySCs})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:177
	// _ = "end of CoverTab[67464]"
}

func (b *baseBalancer) UpdateSubConnState(sc balancer.SubConn, state balancer.SubConnState) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:180
	_go_fuzz_dep_.CoverTab[67471]++
												s := state.ConnectivityState
												if logger.V(2) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:182
		_go_fuzz_dep_.CoverTab[67477]++
													logger.Infof("base.baseBalancer: handle SubConn state change: %p, %v", sc, s)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:183
		// _ = "end of CoverTab[67477]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:184
		_go_fuzz_dep_.CoverTab[67478]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:184
		// _ = "end of CoverTab[67478]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:184
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:184
	// _ = "end of CoverTab[67471]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:184
	_go_fuzz_dep_.CoverTab[67472]++
												oldS, ok := b.scStates[sc]
												if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:186
		_go_fuzz_dep_.CoverTab[67479]++
													if logger.V(2) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:187
			_go_fuzz_dep_.CoverTab[67481]++
														logger.Infof("base.baseBalancer: got state changes for an unknown SubConn: %p, %v", sc, s)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:188
			// _ = "end of CoverTab[67481]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:189
			_go_fuzz_dep_.CoverTab[67482]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:189
			// _ = "end of CoverTab[67482]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:189
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:189
		// _ = "end of CoverTab[67479]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:189
		_go_fuzz_dep_.CoverTab[67480]++
													return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:190
		// _ = "end of CoverTab[67480]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:191
		_go_fuzz_dep_.CoverTab[67483]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:191
		// _ = "end of CoverTab[67483]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:191
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:191
	// _ = "end of CoverTab[67472]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:191
	_go_fuzz_dep_.CoverTab[67473]++
												if oldS == connectivity.TransientFailure && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:192
		_go_fuzz_dep_.CoverTab[67484]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:192
		return (s == connectivity.Connecting || func() bool {
														_go_fuzz_dep_.CoverTab[67485]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:193
			return s == connectivity.Idle
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:193
			// _ = "end of CoverTab[67485]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:193
		}())
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:193
		// _ = "end of CoverTab[67484]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:193
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:193
		_go_fuzz_dep_.CoverTab[67486]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:197
		if s == connectivity.Idle {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:197
			_go_fuzz_dep_.CoverTab[67488]++
														sc.Connect()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:198
			// _ = "end of CoverTab[67488]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:199
			_go_fuzz_dep_.CoverTab[67489]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:199
			// _ = "end of CoverTab[67489]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:199
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:199
		// _ = "end of CoverTab[67486]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:199
		_go_fuzz_dep_.CoverTab[67487]++
													return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:200
		// _ = "end of CoverTab[67487]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:201
		_go_fuzz_dep_.CoverTab[67490]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:201
		// _ = "end of CoverTab[67490]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:201
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:201
	// _ = "end of CoverTab[67473]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:201
	_go_fuzz_dep_.CoverTab[67474]++
												b.scStates[sc] = s
												switch s {
	case connectivity.Idle:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:204
		_go_fuzz_dep_.CoverTab[67491]++
													sc.Connect()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:205
		// _ = "end of CoverTab[67491]"
	case connectivity.Shutdown:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:206
		_go_fuzz_dep_.CoverTab[67492]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:209
		delete(b.scStates, sc)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:209
		// _ = "end of CoverTab[67492]"
	case connectivity.TransientFailure:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:210
		_go_fuzz_dep_.CoverTab[67493]++

													b.connErr = state.ConnectionError
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:212
		// _ = "end of CoverTab[67493]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:212
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:212
		_go_fuzz_dep_.CoverTab[67494]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:212
		// _ = "end of CoverTab[67494]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:213
	// _ = "end of CoverTab[67474]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:213
	_go_fuzz_dep_.CoverTab[67475]++

												b.state = b.csEvltr.RecordTransition(oldS, s)

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:221
	if (s == connectivity.Ready) != (oldS == connectivity.Ready) || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:221
		_go_fuzz_dep_.CoverTab[67495]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:221
		return b.state == connectivity.TransientFailure
													// _ = "end of CoverTab[67495]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:222
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:222
		_go_fuzz_dep_.CoverTab[67496]++
													b.regeneratePicker()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:223
		// _ = "end of CoverTab[67496]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:224
		_go_fuzz_dep_.CoverTab[67497]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:224
		// _ = "end of CoverTab[67497]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:224
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:224
	// _ = "end of CoverTab[67475]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:224
	_go_fuzz_dep_.CoverTab[67476]++
												b.cc.UpdateState(balancer.State{ConnectivityState: b.state, Picker: b.picker})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:225
	// _ = "end of CoverTab[67476]"
}

// Close is a nop because base balancer doesn't have internal state to clean up,
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:228
// and it doesn't need to call RemoveSubConn for the SubConns.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:230
func (b *baseBalancer) Close() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:230
	_go_fuzz_dep_.CoverTab[67498]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:230
	// _ = "end of CoverTab[67498]"
}

// ExitIdle is a nop because the base balancer attempts to stay connected to
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:233
// all SubConns at all times.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:235
func (b *baseBalancer) ExitIdle() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:235
	_go_fuzz_dep_.CoverTab[67499]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:235
	// _ = "end of CoverTab[67499]"
}

// NewErrPicker returns a Picker that always returns err on Pick().
func NewErrPicker(err error) balancer.Picker {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:239
	_go_fuzz_dep_.CoverTab[67500]++
												return &errPicker{err: err}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:240
	// _ = "end of CoverTab[67500]"
}

// NewErrPickerV2 is temporarily defined for backward compatibility reasons.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:243
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:243
// Deprecated: use NewErrPicker instead.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:246
var NewErrPickerV2 = NewErrPicker

type errPicker struct {
	err error	// Pick() always returns this err.
}

func (p *errPicker) Pick(info balancer.PickInfo) (balancer.PickResult, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:252
	_go_fuzz_dep_.CoverTab[67501]++
												return balancer.PickResult{}, p.err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:253
	// _ = "end of CoverTab[67501]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:254
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/balancer.go:254
var _ = _go_fuzz_dep_.CoverTab
