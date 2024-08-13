//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:19
package grpc

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:19
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:19
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:19
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:19
)

import (
	"errors"
	"fmt"

	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/connectivity"
)

// PickFirstBalancerName is the name of the pick_first balancer.
const PickFirstBalancerName = "pick_first"

func newPickfirstBuilder() balancer.Builder {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:32
	_go_fuzz_dep_.CoverTab[79540]++
											return &pickfirstBuilder{}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:33
	// _ = "end of CoverTab[79540]"
}

type pickfirstBuilder struct{}

func (*pickfirstBuilder) Build(cc balancer.ClientConn, opt balancer.BuildOptions) balancer.Balancer {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:38
	_go_fuzz_dep_.CoverTab[79541]++
											return &pickfirstBalancer{cc: cc}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:39
	// _ = "end of CoverTab[79541]"
}

func (*pickfirstBuilder) Name() string {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:42
	_go_fuzz_dep_.CoverTab[79542]++
											return PickFirstBalancerName
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:43
	// _ = "end of CoverTab[79542]"
}

type pickfirstBalancer struct {
	state	connectivity.State
	cc	balancer.ClientConn
	subConn	balancer.SubConn
}

func (b *pickfirstBalancer) ResolverError(err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:52
	_go_fuzz_dep_.CoverTab[79543]++
											if logger.V(2) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:53
		_go_fuzz_dep_.CoverTab[79547]++
												logger.Infof("pickfirstBalancer: ResolverError called with error: %v", err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:54
		// _ = "end of CoverTab[79547]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:55
		_go_fuzz_dep_.CoverTab[79548]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:55
		// _ = "end of CoverTab[79548]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:55
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:55
	// _ = "end of CoverTab[79543]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:55
	_go_fuzz_dep_.CoverTab[79544]++
											if b.subConn == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:56
		_go_fuzz_dep_.CoverTab[79549]++
												b.state = connectivity.TransientFailure
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:57
		// _ = "end of CoverTab[79549]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:58
		_go_fuzz_dep_.CoverTab[79550]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:58
		// _ = "end of CoverTab[79550]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:58
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:58
	// _ = "end of CoverTab[79544]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:58
	_go_fuzz_dep_.CoverTab[79545]++

											if b.state != connectivity.TransientFailure {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:60
		_go_fuzz_dep_.CoverTab[79551]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:63
		return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:63
		// _ = "end of CoverTab[79551]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:64
		_go_fuzz_dep_.CoverTab[79552]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:64
		// _ = "end of CoverTab[79552]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:64
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:64
	// _ = "end of CoverTab[79545]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:64
	_go_fuzz_dep_.CoverTab[79546]++
											b.cc.UpdateState(balancer.State{
		ConnectivityState:	connectivity.TransientFailure,
		Picker:			&picker{err: fmt.Errorf("name resolver error: %v", err)},
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:68
	// _ = "end of CoverTab[79546]"
}

func (b *pickfirstBalancer) UpdateClientConnState(state balancer.ClientConnState) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:71
	_go_fuzz_dep_.CoverTab[79553]++
											if len(state.ResolverState.Addresses) == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:72
		_go_fuzz_dep_.CoverTab[79557]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:75
		if b.subConn != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:75
			_go_fuzz_dep_.CoverTab[79559]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:78
			b.cc.RemoveSubConn(b.subConn)
													b.subConn = nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:79
			// _ = "end of CoverTab[79559]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:80
			_go_fuzz_dep_.CoverTab[79560]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:80
			// _ = "end of CoverTab[79560]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:80
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:80
		// _ = "end of CoverTab[79557]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:80
		_go_fuzz_dep_.CoverTab[79558]++
												b.ResolverError(errors.New("produced zero addresses"))
												return balancer.ErrBadResolverState
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:82
		// _ = "end of CoverTab[79558]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:83
		_go_fuzz_dep_.CoverTab[79561]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:83
		// _ = "end of CoverTab[79561]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:83
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:83
	// _ = "end of CoverTab[79553]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:83
	_go_fuzz_dep_.CoverTab[79554]++

											if b.subConn != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:85
		_go_fuzz_dep_.CoverTab[79562]++
												b.cc.UpdateAddresses(b.subConn, state.ResolverState.Addresses)
												return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:87
		// _ = "end of CoverTab[79562]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:88
		_go_fuzz_dep_.CoverTab[79563]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:88
		// _ = "end of CoverTab[79563]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:88
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:88
	// _ = "end of CoverTab[79554]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:88
	_go_fuzz_dep_.CoverTab[79555]++

											subConn, err := b.cc.NewSubConn(state.ResolverState.Addresses, balancer.NewSubConnOptions{})
											if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:91
		_go_fuzz_dep_.CoverTab[79564]++
												if logger.V(2) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:92
			_go_fuzz_dep_.CoverTab[79566]++
													logger.Errorf("pickfirstBalancer: failed to NewSubConn: %v", err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:93
			// _ = "end of CoverTab[79566]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:94
			_go_fuzz_dep_.CoverTab[79567]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:94
			// _ = "end of CoverTab[79567]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:94
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:94
		// _ = "end of CoverTab[79564]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:94
		_go_fuzz_dep_.CoverTab[79565]++
												b.state = connectivity.TransientFailure
												b.cc.UpdateState(balancer.State{
			ConnectivityState:	connectivity.TransientFailure,
			Picker:			&picker{err: fmt.Errorf("error creating connection: %v", err)},
		})
												return balancer.ErrBadResolverState
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:100
		// _ = "end of CoverTab[79565]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:101
		_go_fuzz_dep_.CoverTab[79568]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:101
		// _ = "end of CoverTab[79568]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:101
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:101
	// _ = "end of CoverTab[79555]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:101
	_go_fuzz_dep_.CoverTab[79556]++
											b.subConn = subConn
											b.state = connectivity.Idle
											b.cc.UpdateState(balancer.State{
		ConnectivityState:	connectivity.Connecting,
		Picker:			&picker{err: balancer.ErrNoSubConnAvailable},
	})
											b.subConn.Connect()
											return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:109
	// _ = "end of CoverTab[79556]"
}

func (b *pickfirstBalancer) UpdateSubConnState(subConn balancer.SubConn, state balancer.SubConnState) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:112
	_go_fuzz_dep_.CoverTab[79569]++
											if logger.V(2) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:113
		_go_fuzz_dep_.CoverTab[79573]++
												logger.Infof("pickfirstBalancer: UpdateSubConnState: %p, %v", subConn, state)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:114
		// _ = "end of CoverTab[79573]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:115
		_go_fuzz_dep_.CoverTab[79574]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:115
		// _ = "end of CoverTab[79574]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:115
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:115
	// _ = "end of CoverTab[79569]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:115
	_go_fuzz_dep_.CoverTab[79570]++
											if b.subConn != subConn {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:116
		_go_fuzz_dep_.CoverTab[79575]++
												if logger.V(2) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:117
			_go_fuzz_dep_.CoverTab[79577]++
													logger.Infof("pickfirstBalancer: ignored state change because subConn is not recognized")
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:118
			// _ = "end of CoverTab[79577]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:119
			_go_fuzz_dep_.CoverTab[79578]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:119
			// _ = "end of CoverTab[79578]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:119
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:119
		// _ = "end of CoverTab[79575]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:119
		_go_fuzz_dep_.CoverTab[79576]++
												return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:120
		// _ = "end of CoverTab[79576]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:121
		_go_fuzz_dep_.CoverTab[79579]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:121
		// _ = "end of CoverTab[79579]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:121
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:121
	// _ = "end of CoverTab[79570]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:121
	_go_fuzz_dep_.CoverTab[79571]++
											b.state = state.ConnectivityState
											if state.ConnectivityState == connectivity.Shutdown {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:123
		_go_fuzz_dep_.CoverTab[79580]++
												b.subConn = nil
												return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:125
		// _ = "end of CoverTab[79580]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:126
		_go_fuzz_dep_.CoverTab[79581]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:126
		// _ = "end of CoverTab[79581]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:126
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:126
	// _ = "end of CoverTab[79571]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:126
	_go_fuzz_dep_.CoverTab[79572]++

											switch state.ConnectivityState {
	case connectivity.Ready:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:129
		_go_fuzz_dep_.CoverTab[79582]++
												b.cc.UpdateState(balancer.State{
			ConnectivityState:	state.ConnectivityState,
			Picker:			&picker{result: balancer.PickResult{SubConn: subConn}},
		})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:133
		// _ = "end of CoverTab[79582]"
	case connectivity.Connecting:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:134
		_go_fuzz_dep_.CoverTab[79583]++
												b.cc.UpdateState(balancer.State{
			ConnectivityState:	state.ConnectivityState,
			Picker:			&picker{err: balancer.ErrNoSubConnAvailable},
		})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:138
		// _ = "end of CoverTab[79583]"
	case connectivity.Idle:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:139
		_go_fuzz_dep_.CoverTab[79584]++
												b.cc.UpdateState(balancer.State{
			ConnectivityState:	state.ConnectivityState,
			Picker:			&idlePicker{subConn: subConn},
		})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:143
		// _ = "end of CoverTab[79584]"
	case connectivity.TransientFailure:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:144
		_go_fuzz_dep_.CoverTab[79585]++
												b.cc.UpdateState(balancer.State{
			ConnectivityState:	state.ConnectivityState,
			Picker:			&picker{err: state.ConnectionError},
		})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:148
		// _ = "end of CoverTab[79585]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:148
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:148
		_go_fuzz_dep_.CoverTab[79586]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:148
		// _ = "end of CoverTab[79586]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:149
	// _ = "end of CoverTab[79572]"
}

func (b *pickfirstBalancer) Close() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:152
	_go_fuzz_dep_.CoverTab[79587]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:152
	// _ = "end of CoverTab[79587]"
}

func (b *pickfirstBalancer) ExitIdle() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:155
	_go_fuzz_dep_.CoverTab[79588]++
											if b.subConn != nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:156
		_go_fuzz_dep_.CoverTab[79589]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:156
		return b.state == connectivity.Idle
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:156
		// _ = "end of CoverTab[79589]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:156
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:156
		_go_fuzz_dep_.CoverTab[79590]++
												b.subConn.Connect()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:157
		// _ = "end of CoverTab[79590]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:158
		_go_fuzz_dep_.CoverTab[79591]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:158
		// _ = "end of CoverTab[79591]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:158
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:158
	// _ = "end of CoverTab[79588]"
}

type picker struct {
	result	balancer.PickResult
	err	error
}

func (p *picker) Pick(balancer.PickInfo) (balancer.PickResult, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:166
	_go_fuzz_dep_.CoverTab[79592]++
											return p.result, p.err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:167
	// _ = "end of CoverTab[79592]"
}

// idlePicker is used when the SubConn is IDLE and kicks the SubConn into
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:170
// CONNECTING when Pick is called.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:172
type idlePicker struct {
	subConn balancer.SubConn
}

func (i *idlePicker) Pick(balancer.PickInfo) (balancer.PickResult, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:176
	_go_fuzz_dep_.CoverTab[79593]++
											i.subConn.Connect()
											return balancer.PickResult{}, balancer.ErrNoSubConnAvailable
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:178
	// _ = "end of CoverTab[79593]"
}

func init() {
	balancer.Register(newPickfirstBuilder())
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:183
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/pickfirst.go:183
var _ = _go_fuzz_dep_.CoverTab
