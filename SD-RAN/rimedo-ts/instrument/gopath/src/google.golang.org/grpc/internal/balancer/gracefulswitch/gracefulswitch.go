//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:19
// Package gracefulswitch implements a graceful switch load balancer.
package gracefulswitch

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:20
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:20
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:20
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:20
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:20
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:20
)

import (
	"errors"
	"fmt"
	"sync"

	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/balancer/base"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/resolver"
)

var errBalancerClosed = errors.New("gracefulSwitchBalancer is closed")
var _ balancer.Balancer = (*Balancer)(nil)

// NewBalancer returns a graceful switch Balancer.
func NewBalancer(cc balancer.ClientConn, opts balancer.BuildOptions) *Balancer {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:37
	_go_fuzz_dep_.CoverTab[67688]++
															return &Balancer{
		cc:	cc,
		bOpts:	opts,
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:41
	// _ = "end of CoverTab[67688]"
}

// Balancer is a utility to gracefully switch from one balancer to
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:44
// a new balancer. It implements the balancer.Balancer interface.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:46
type Balancer struct {
	bOpts	balancer.BuildOptions
	cc	balancer.ClientConn

	// mu protects the following fields and all fields within balancerCurrent
	// and balancerPending. mu does not need to be held when calling into the
	// child balancers, as all calls into these children happen only as a direct
	// result of a call into the gracefulSwitchBalancer, which are also
	// guaranteed to be synchronous. There is one exception: an UpdateState call
	// from a child balancer when current and pending are populated can lead to
	// calling Close() on the current. To prevent that racing with an
	// UpdateSubConnState from the channel, we hold currentMu during Close and
	// UpdateSubConnState calls.
	mu		sync.Mutex
	balancerCurrent	*balancerWrapper
	balancerPending	*balancerWrapper
	closed		bool	// set to true when this balancer is closed

	// currentMu must be locked before mu. This mutex guards against this
	// sequence of events: UpdateSubConnState() called, finds the
	// balancerCurrent, gives up lock, updateState comes in, causes Close() on
	// balancerCurrent before the UpdateSubConnState is called on the
	// balancerCurrent.
	currentMu	sync.Mutex
}

// swap swaps out the current lb with the pending lb and updates the ClientConn.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:72
// The caller must hold gsb.mu.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:74
func (gsb *Balancer) swap() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:74
	_go_fuzz_dep_.CoverTab[67689]++
															gsb.cc.UpdateState(gsb.balancerPending.lastState)
															cur := gsb.balancerCurrent
															gsb.balancerCurrent = gsb.balancerPending
															gsb.balancerPending = nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:78
	_curRoutineNum51_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:78
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum51_)
															go func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:79
		_go_fuzz_dep_.CoverTab[67690]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:79
		defer func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:79
			_go_fuzz_dep_.CoverTab[67691]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:79
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum51_)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:79
			// _ = "end of CoverTab[67691]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:79
		}()
																gsb.currentMu.Lock()
																defer gsb.currentMu.Unlock()
																cur.Close()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:82
		// _ = "end of CoverTab[67690]"
	}()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:83
	// _ = "end of CoverTab[67689]"
}

// Helper function that checks if the balancer passed in is current or pending.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:86
// The caller must hold gsb.mu.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:88
func (gsb *Balancer) balancerCurrentOrPending(bw *balancerWrapper) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:88
	_go_fuzz_dep_.CoverTab[67692]++
															return bw == gsb.balancerCurrent || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:89
		_go_fuzz_dep_.CoverTab[67693]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:89
		return bw == gsb.balancerPending
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:89
		// _ = "end of CoverTab[67693]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:89
	}()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:89
	// _ = "end of CoverTab[67692]"
}

// SwitchTo initializes the graceful switch process, which completes based on
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:92
// connectivity state changes on the current/pending balancer. Thus, the switch
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:92
// process is not complete when this method returns. This method must be called
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:92
// synchronously alongside the rest of the balancer.Balancer methods this
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:92
// Graceful Switch Balancer implements.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:97
func (gsb *Balancer) SwitchTo(builder balancer.Builder) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:97
	_go_fuzz_dep_.CoverTab[67694]++
															gsb.mu.Lock()
															if gsb.closed {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:99
			_go_fuzz_dep_.CoverTab[67698]++
																	gsb.mu.Unlock()
																	return errBalancerClosed
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:101
		// _ = "end of CoverTab[67698]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:102
		_go_fuzz_dep_.CoverTab[67699]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:102
		// _ = "end of CoverTab[67699]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:102
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:102
	// _ = "end of CoverTab[67694]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:102
	_go_fuzz_dep_.CoverTab[67695]++
																bw := &balancerWrapper{
		gsb:	gsb,
		lastState: balancer.State{
			ConnectivityState:	connectivity.Connecting,
			Picker:			base.NewErrPicker(balancer.ErrNoSubConnAvailable),
		},
		subconns:	make(map[balancer.SubConn]bool),
	}
	balToClose := gsb.balancerPending
	if gsb.balancerCurrent == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:112
		_go_fuzz_dep_.CoverTab[67700]++
																	gsb.balancerCurrent = bw
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:113
		// _ = "end of CoverTab[67700]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:114
		_go_fuzz_dep_.CoverTab[67701]++
																	gsb.balancerPending = bw
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:115
		// _ = "end of CoverTab[67701]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:116
	// _ = "end of CoverTab[67695]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:116
	_go_fuzz_dep_.CoverTab[67696]++
																gsb.mu.Unlock()
																balToClose.Close()

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:121
	newBalancer := builder.Build(bw, gsb.bOpts)
	if newBalancer == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:122
		_go_fuzz_dep_.CoverTab[67702]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:125
		gsb.mu.Lock()
		if gsb.balancerPending != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:126
			_go_fuzz_dep_.CoverTab[67704]++
																		gsb.balancerPending = nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:127
			// _ = "end of CoverTab[67704]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:128
			_go_fuzz_dep_.CoverTab[67705]++
																		gsb.balancerCurrent = nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:129
			// _ = "end of CoverTab[67705]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:130
		// _ = "end of CoverTab[67702]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:130
		_go_fuzz_dep_.CoverTab[67703]++
																	gsb.mu.Unlock()
																	return balancer.ErrBadResolverState
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:132
		// _ = "end of CoverTab[67703]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:133
		_go_fuzz_dep_.CoverTab[67706]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:133
		// _ = "end of CoverTab[67706]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:133
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:133
	// _ = "end of CoverTab[67696]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:133
	_go_fuzz_dep_.CoverTab[67697]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:140
	bw.Balancer = newBalancer
																return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:141
	// _ = "end of CoverTab[67697]"
}

// Returns nil if the graceful switch balancer is closed.
func (gsb *Balancer) latestBalancer() *balancerWrapper {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:145
	_go_fuzz_dep_.CoverTab[67707]++
																gsb.mu.Lock()
																defer gsb.mu.Unlock()
																if gsb.balancerPending != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:148
		_go_fuzz_dep_.CoverTab[67709]++
																	return gsb.balancerPending
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:149
		// _ = "end of CoverTab[67709]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:150
		_go_fuzz_dep_.CoverTab[67710]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:150
		// _ = "end of CoverTab[67710]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:150
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:150
	// _ = "end of CoverTab[67707]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:150
	_go_fuzz_dep_.CoverTab[67708]++
																return gsb.balancerCurrent
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:151
	// _ = "end of CoverTab[67708]"
}

// UpdateClientConnState forwards the update to the latest balancer created.
func (gsb *Balancer) UpdateClientConnState(state balancer.ClientConnState) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:155
	_go_fuzz_dep_.CoverTab[67711]++

																balToUpdate := gsb.latestBalancer()
																if balToUpdate == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:158
		_go_fuzz_dep_.CoverTab[67713]++
																	return errBalancerClosed
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:159
		// _ = "end of CoverTab[67713]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:160
		_go_fuzz_dep_.CoverTab[67714]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:160
		// _ = "end of CoverTab[67714]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:160
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:160
	// _ = "end of CoverTab[67711]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:160
	_go_fuzz_dep_.CoverTab[67712]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:164
	return balToUpdate.UpdateClientConnState(state)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:164
	// _ = "end of CoverTab[67712]"
}

// ResolverError forwards the error to the latest balancer created.
func (gsb *Balancer) ResolverError(err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:168
	_go_fuzz_dep_.CoverTab[67715]++

																balToUpdate := gsb.latestBalancer()
																if balToUpdate == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:171
		_go_fuzz_dep_.CoverTab[67717]++
																	return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:172
		// _ = "end of CoverTab[67717]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:173
		_go_fuzz_dep_.CoverTab[67718]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:173
		// _ = "end of CoverTab[67718]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:173
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:173
	// _ = "end of CoverTab[67715]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:173
	_go_fuzz_dep_.CoverTab[67716]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:177
	balToUpdate.ResolverError(err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:177
	// _ = "end of CoverTab[67716]"
}

// ExitIdle forwards the call to the latest balancer created.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:180
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:180
// If the latest balancer does not support ExitIdle, the subConns are
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:180
// re-connected to manually.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:184
func (gsb *Balancer) ExitIdle() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:184
	_go_fuzz_dep_.CoverTab[67719]++
																balToUpdate := gsb.latestBalancer()
																if balToUpdate == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:186
		_go_fuzz_dep_.CoverTab[67722]++
																	return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:187
		// _ = "end of CoverTab[67722]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:188
		_go_fuzz_dep_.CoverTab[67723]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:188
		// _ = "end of CoverTab[67723]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:188
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:188
	// _ = "end of CoverTab[67719]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:188
	_go_fuzz_dep_.CoverTab[67720]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:192
	if ei, ok := balToUpdate.Balancer.(balancer.ExitIdler); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:192
		_go_fuzz_dep_.CoverTab[67724]++
																	ei.ExitIdle()
																	return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:194
		// _ = "end of CoverTab[67724]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:195
		_go_fuzz_dep_.CoverTab[67725]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:195
		// _ = "end of CoverTab[67725]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:195
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:195
	// _ = "end of CoverTab[67720]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:195
	_go_fuzz_dep_.CoverTab[67721]++
																gsb.mu.Lock()
																defer gsb.mu.Unlock()
																for sc := range balToUpdate.subconns {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:198
		_go_fuzz_dep_.CoverTab[67726]++
																	sc.Connect()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:199
		// _ = "end of CoverTab[67726]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:200
	// _ = "end of CoverTab[67721]"
}

// UpdateSubConnState forwards the update to the appropriate child.
func (gsb *Balancer) UpdateSubConnState(sc balancer.SubConn, state balancer.SubConnState) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:204
	_go_fuzz_dep_.CoverTab[67727]++
																gsb.currentMu.Lock()
																defer gsb.currentMu.Unlock()
																gsb.mu.Lock()
	// Forward update to the appropriate child.  Even if there is a pending
	// balancer, the current balancer should continue to get SubConn updates to
	// maintain the proper state while the pending is still connecting.
	var balToUpdate *balancerWrapper
	if gsb.balancerCurrent != nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:212
		_go_fuzz_dep_.CoverTab[67730]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:212
		return gsb.balancerCurrent.subconns[sc]
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:212
		// _ = "end of CoverTab[67730]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:212
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:212
		_go_fuzz_dep_.CoverTab[67731]++
																	balToUpdate = gsb.balancerCurrent
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:213
		// _ = "end of CoverTab[67731]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:214
		_go_fuzz_dep_.CoverTab[67732]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:214
		if gsb.balancerPending != nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:214
			_go_fuzz_dep_.CoverTab[67733]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:214
			return gsb.balancerPending.subconns[sc]
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:214
			// _ = "end of CoverTab[67733]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:214
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:214
			_go_fuzz_dep_.CoverTab[67734]++
																		balToUpdate = gsb.balancerPending
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:215
			// _ = "end of CoverTab[67734]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:216
			_go_fuzz_dep_.CoverTab[67735]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:216
			// _ = "end of CoverTab[67735]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:216
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:216
		// _ = "end of CoverTab[67732]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:216
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:216
	// _ = "end of CoverTab[67727]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:216
	_go_fuzz_dep_.CoverTab[67728]++
																gsb.mu.Unlock()
																if balToUpdate == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:218
		_go_fuzz_dep_.CoverTab[67736]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:221
		return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:221
		// _ = "end of CoverTab[67736]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:222
		_go_fuzz_dep_.CoverTab[67737]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:222
		// _ = "end of CoverTab[67737]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:222
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:222
	// _ = "end of CoverTab[67728]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:222
	_go_fuzz_dep_.CoverTab[67729]++
																balToUpdate.UpdateSubConnState(sc, state)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:223
	// _ = "end of CoverTab[67729]"
}

// Close closes any active child balancers.
func (gsb *Balancer) Close() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:227
	_go_fuzz_dep_.CoverTab[67738]++
																gsb.mu.Lock()
																gsb.closed = true
																currentBalancerToClose := gsb.balancerCurrent
																gsb.balancerCurrent = nil
																pendingBalancerToClose := gsb.balancerPending
																gsb.balancerPending = nil
																gsb.mu.Unlock()

																currentBalancerToClose.Close()
																pendingBalancerToClose.Close()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:237
	// _ = "end of CoverTab[67738]"
}

// balancerWrapper wraps a balancer.Balancer, and overrides some Balancer
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:240
// methods to help cleanup SubConns created by the wrapped balancer.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:240
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:240
// It implements the balancer.ClientConn interface and is passed down in that
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:240
// capacity to the wrapped balancer. It maintains a set of subConns created by
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:240
// the wrapped balancer and calls from the latter to create/update/remove
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:240
// SubConns update this set before being forwarded to the parent ClientConn.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:240
// State updates from the wrapped balancer can result in invocation of the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:240
// graceful switch logic.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:249
type balancerWrapper struct {
	balancer.Balancer
	gsb	*Balancer

	lastState	balancer.State
	subconns	map[balancer.SubConn]bool	// subconns created by this balancer
}

func (bw *balancerWrapper) UpdateSubConnState(sc balancer.SubConn, state balancer.SubConnState) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:257
	_go_fuzz_dep_.CoverTab[67739]++
																if state.ConnectivityState == connectivity.Shutdown {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:258
		_go_fuzz_dep_.CoverTab[67741]++
																	bw.gsb.mu.Lock()
																	delete(bw.subconns, sc)
																	bw.gsb.mu.Unlock()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:261
		// _ = "end of CoverTab[67741]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:262
		_go_fuzz_dep_.CoverTab[67742]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:262
		// _ = "end of CoverTab[67742]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:262
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:262
	// _ = "end of CoverTab[67739]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:262
	_go_fuzz_dep_.CoverTab[67740]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:266
	bw.Balancer.UpdateSubConnState(sc, state)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:266
	// _ = "end of CoverTab[67740]"
}

// Close closes the underlying LB policy and removes the subconns it created. bw
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:269
// must not be referenced via balancerCurrent or balancerPending in gsb when
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:269
// called. gsb.mu must not be held.  Does not panic with a nil receiver.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:272
func (bw *balancerWrapper) Close() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:272
	_go_fuzz_dep_.CoverTab[67743]++

																if bw == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:274
		_go_fuzz_dep_.CoverTab[67746]++
																	return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:275
		// _ = "end of CoverTab[67746]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:276
		_go_fuzz_dep_.CoverTab[67747]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:276
		// _ = "end of CoverTab[67747]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:276
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:276
	// _ = "end of CoverTab[67743]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:276
	_go_fuzz_dep_.CoverTab[67744]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:281
	bw.Balancer.Close()
	bw.gsb.mu.Lock()
	for sc := range bw.subconns {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:283
		_go_fuzz_dep_.CoverTab[67748]++
																	bw.gsb.cc.RemoveSubConn(sc)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:284
		// _ = "end of CoverTab[67748]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:285
	// _ = "end of CoverTab[67744]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:285
	_go_fuzz_dep_.CoverTab[67745]++
																bw.gsb.mu.Unlock()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:286
	// _ = "end of CoverTab[67745]"
}

func (bw *balancerWrapper) UpdateState(state balancer.State) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:289
	_go_fuzz_dep_.CoverTab[67749]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:293
	bw.gsb.mu.Lock()
	defer bw.gsb.mu.Unlock()
	bw.lastState = state

	if !bw.gsb.balancerCurrentOrPending(bw) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:297
		_go_fuzz_dep_.CoverTab[67752]++
																	return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:298
		// _ = "end of CoverTab[67752]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:299
		_go_fuzz_dep_.CoverTab[67753]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:299
		// _ = "end of CoverTab[67753]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:299
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:299
	// _ = "end of CoverTab[67749]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:299
	_go_fuzz_dep_.CoverTab[67750]++

																if bw == bw.gsb.balancerCurrent {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:301
		_go_fuzz_dep_.CoverTab[67754]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:307
		if state.ConnectivityState != connectivity.Ready && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:307
			_go_fuzz_dep_.CoverTab[67756]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:307
			return bw.gsb.balancerPending != nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:307
			// _ = "end of CoverTab[67756]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:307
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:307
			_go_fuzz_dep_.CoverTab[67757]++
																		bw.gsb.swap()
																		return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:309
			// _ = "end of CoverTab[67757]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:310
			_go_fuzz_dep_.CoverTab[67758]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:310
			// _ = "end of CoverTab[67758]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:310
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:310
		// _ = "end of CoverTab[67754]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:310
		_go_fuzz_dep_.CoverTab[67755]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:317
		bw.gsb.cc.UpdateState(state)
																	return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:318
		// _ = "end of CoverTab[67755]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:319
		_go_fuzz_dep_.CoverTab[67759]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:319
		// _ = "end of CoverTab[67759]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:319
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:319
	// _ = "end of CoverTab[67750]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:319
	_go_fuzz_dep_.CoverTab[67751]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:325
	if state.ConnectivityState != connectivity.Connecting || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:325
		_go_fuzz_dep_.CoverTab[67760]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:325
		return bw.gsb.balancerCurrent.lastState.ConnectivityState != connectivity.Ready
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:325
		// _ = "end of CoverTab[67760]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:325
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:325
		_go_fuzz_dep_.CoverTab[67761]++
																	bw.gsb.swap()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:326
		// _ = "end of CoverTab[67761]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:327
		_go_fuzz_dep_.CoverTab[67762]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:327
		// _ = "end of CoverTab[67762]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:327
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:327
	// _ = "end of CoverTab[67751]"
}

func (bw *balancerWrapper) NewSubConn(addrs []resolver.Address, opts balancer.NewSubConnOptions) (balancer.SubConn, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:330
	_go_fuzz_dep_.CoverTab[67763]++
																bw.gsb.mu.Lock()
																if !bw.gsb.balancerCurrentOrPending(bw) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:332
		_go_fuzz_dep_.CoverTab[67767]++
																	bw.gsb.mu.Unlock()
																	return nil, fmt.Errorf("%T at address %p that called NewSubConn is deleted", bw, bw)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:334
		// _ = "end of CoverTab[67767]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:335
		_go_fuzz_dep_.CoverTab[67768]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:335
		// _ = "end of CoverTab[67768]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:335
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:335
	// _ = "end of CoverTab[67763]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:335
	_go_fuzz_dep_.CoverTab[67764]++
																bw.gsb.mu.Unlock()

																sc, err := bw.gsb.cc.NewSubConn(addrs, opts)
																if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:339
		_go_fuzz_dep_.CoverTab[67769]++
																	return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:340
		// _ = "end of CoverTab[67769]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:341
		_go_fuzz_dep_.CoverTab[67770]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:341
		// _ = "end of CoverTab[67770]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:341
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:341
	// _ = "end of CoverTab[67764]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:341
	_go_fuzz_dep_.CoverTab[67765]++
																bw.gsb.mu.Lock()
																if !bw.gsb.balancerCurrentOrPending(bw) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:343
		_go_fuzz_dep_.CoverTab[67771]++
																	bw.gsb.cc.RemoveSubConn(sc)
																	bw.gsb.mu.Unlock()
																	return nil, fmt.Errorf("%T at address %p that called NewSubConn is deleted", bw, bw)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:346
		// _ = "end of CoverTab[67771]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:347
		_go_fuzz_dep_.CoverTab[67772]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:347
		// _ = "end of CoverTab[67772]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:347
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:347
	// _ = "end of CoverTab[67765]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:347
	_go_fuzz_dep_.CoverTab[67766]++
																bw.subconns[sc] = true
																bw.gsb.mu.Unlock()
																return sc, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:350
	// _ = "end of CoverTab[67766]"
}

func (bw *balancerWrapper) ResolveNow(opts resolver.ResolveNowOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:353
	_go_fuzz_dep_.CoverTab[67773]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:356
	if bw != bw.gsb.latestBalancer() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:356
		_go_fuzz_dep_.CoverTab[67775]++
																	return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:357
		// _ = "end of CoverTab[67775]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:358
		_go_fuzz_dep_.CoverTab[67776]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:358
		// _ = "end of CoverTab[67776]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:358
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:358
	// _ = "end of CoverTab[67773]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:358
	_go_fuzz_dep_.CoverTab[67774]++
																bw.gsb.cc.ResolveNow(opts)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:359
	// _ = "end of CoverTab[67774]"
}

func (bw *balancerWrapper) RemoveSubConn(sc balancer.SubConn) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:362
	_go_fuzz_dep_.CoverTab[67777]++
																bw.gsb.mu.Lock()
																if !bw.gsb.balancerCurrentOrPending(bw) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:364
		_go_fuzz_dep_.CoverTab[67779]++
																	bw.gsb.mu.Unlock()
																	return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:366
		// _ = "end of CoverTab[67779]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:367
		_go_fuzz_dep_.CoverTab[67780]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:367
		// _ = "end of CoverTab[67780]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:367
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:367
	// _ = "end of CoverTab[67777]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:367
	_go_fuzz_dep_.CoverTab[67778]++
																bw.gsb.mu.Unlock()
																bw.gsb.cc.RemoveSubConn(sc)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:369
	// _ = "end of CoverTab[67778]"
}

func (bw *balancerWrapper) UpdateAddresses(sc balancer.SubConn, addrs []resolver.Address) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:372
	_go_fuzz_dep_.CoverTab[67781]++
																bw.gsb.mu.Lock()
																if !bw.gsb.balancerCurrentOrPending(bw) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:374
		_go_fuzz_dep_.CoverTab[67783]++
																	bw.gsb.mu.Unlock()
																	return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:376
		// _ = "end of CoverTab[67783]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:377
		_go_fuzz_dep_.CoverTab[67784]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:377
		// _ = "end of CoverTab[67784]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:377
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:377
	// _ = "end of CoverTab[67781]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:377
	_go_fuzz_dep_.CoverTab[67782]++
																bw.gsb.mu.Unlock()
																bw.gsb.cc.UpdateAddresses(sc, addrs)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:379
	// _ = "end of CoverTab[67782]"
}

func (bw *balancerWrapper) Target() string {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:382
	_go_fuzz_dep_.CoverTab[67785]++
																return bw.gsb.cc.Target()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:383
	// _ = "end of CoverTab[67785]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:384
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancer/gracefulswitch/gracefulswitch.go:384
var _ = _go_fuzz_dep_.CoverTab
