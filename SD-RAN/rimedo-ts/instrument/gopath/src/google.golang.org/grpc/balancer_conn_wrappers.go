//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:19
package grpc

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:19
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:19
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:19
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:19
)

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/internal/balancer/gracefulswitch"
	"google.golang.org/grpc/internal/buffer"
	"google.golang.org/grpc/internal/channelz"
	"google.golang.org/grpc/internal/grpcsync"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/status"
)

// ccBalancerWrapper sits between the ClientConn and the Balancer.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:38
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:38
// ccBalancerWrapper implements methods corresponding to the ones on the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:38
// balancer.Balancer interface. The ClientConn is free to call these methods
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:38
// concurrently and the ccBalancerWrapper ensures that calls from the ClientConn
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:38
// to the Balancer happen synchronously and in order.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:38
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:38
// ccBalancerWrapper also implements the balancer.ClientConn interface and is
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:38
// passed to the Balancer implementations. It invokes unexported methods on the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:38
// ClientConn to handle these calls from the Balancer.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:38
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:38
// It uses the gracefulswitch.Balancer internally to ensure that balancer
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:38
// switches happen in a graceful manner.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:51
type ccBalancerWrapper struct {
	cc	*ClientConn

	// Since these fields are accessed only from handleXxx() methods which are
	// synchronized by the watcher goroutine, we do not need a mutex to protect
	// these fields.
	balancer	*gracefulswitch.Balancer
	curBalancerName	string

	updateCh	*buffer.Unbounded	// Updates written on this channel are processed by watcher().
	resultCh	*buffer.Unbounded	// Results of calls to UpdateClientConnState() are pushed here.
	closed		*grpcsync.Event		// Indicates if close has been called.
	done		*grpcsync.Event		// Indicates if close has completed its work.
}

// newCCBalancerWrapper creates a new balancer wrapper. The underlying balancer
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:66
// is not created until the switchTo() method is invoked.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:68
func newCCBalancerWrapper(cc *ClientConn, bopts balancer.BuildOptions) *ccBalancerWrapper {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:68
	_go_fuzz_dep_.CoverTab[78723]++
												ccb := &ccBalancerWrapper{
		cc:		cc,
		updateCh:	buffer.NewUnbounded(),
		resultCh:	buffer.NewUnbounded(),
		closed:		grpcsync.NewEvent(),
		done:		grpcsync.NewEvent(),
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:75
	_curRoutineNum85_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:75
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum85_)
												go ccb.watcher()
												ccb.balancer = gracefulswitch.NewBalancer(ccb, bopts)
												return ccb
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:78
	// _ = "end of CoverTab[78723]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:85
type ccStateUpdate struct {
	ccs *balancer.ClientConnState
}

type scStateUpdate struct {
	sc	balancer.SubConn
	state	connectivity.State
	err	error
}

type exitIdleUpdate struct{}

type resolverErrorUpdate struct {
	err error
}

type switchToUpdate struct {
	name string
}

type subConnUpdate struct {
	acbw *acBalancerWrapper
}

// watcher is a long-running goroutine which reads updates from a channel and
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:109
// invokes corresponding methods on the underlying balancer. It ensures that
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:109
// these methods are invoked in a synchronous fashion. It also ensures that
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:109
// these methods are invoked in the order in which the updates were received.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:113
func (ccb *ccBalancerWrapper) watcher() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:113
	_go_fuzz_dep_.CoverTab[78724]++
												for {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:114
		_go_fuzz_dep_.CoverTab[78725]++
													select {
		case u := <-ccb.updateCh.Get():
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:116
			_go_fuzz_dep_.CoverTab[78727]++
														ccb.updateCh.Load()
														if ccb.closed.HasFired() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:118
				_go_fuzz_dep_.CoverTab[78730]++
															break
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:119
				// _ = "end of CoverTab[78730]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:120
				_go_fuzz_dep_.CoverTab[78731]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:120
				// _ = "end of CoverTab[78731]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:120
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:120
			// _ = "end of CoverTab[78727]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:120
			_go_fuzz_dep_.CoverTab[78728]++
														switch update := u.(type) {
			case *ccStateUpdate:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:122
				_go_fuzz_dep_.CoverTab[78732]++
															ccb.handleClientConnStateChange(update.ccs)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:123
				// _ = "end of CoverTab[78732]"
			case *scStateUpdate:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:124
				_go_fuzz_dep_.CoverTab[78733]++
															ccb.handleSubConnStateChange(update)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:125
				// _ = "end of CoverTab[78733]"
			case *exitIdleUpdate:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:126
				_go_fuzz_dep_.CoverTab[78734]++
															ccb.handleExitIdle()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:127
				// _ = "end of CoverTab[78734]"
			case *resolverErrorUpdate:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:128
				_go_fuzz_dep_.CoverTab[78735]++
															ccb.handleResolverError(update.err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:129
				// _ = "end of CoverTab[78735]"
			case *switchToUpdate:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:130
				_go_fuzz_dep_.CoverTab[78736]++
															ccb.handleSwitchTo(update.name)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:131
				// _ = "end of CoverTab[78736]"
			case *subConnUpdate:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:132
				_go_fuzz_dep_.CoverTab[78737]++
															ccb.handleRemoveSubConn(update.acbw)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:133
				// _ = "end of CoverTab[78737]"
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:134
				_go_fuzz_dep_.CoverTab[78738]++
															logger.Errorf("ccBalancerWrapper.watcher: unknown update %+v, type %T", update, update)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:135
				// _ = "end of CoverTab[78738]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:136
			// _ = "end of CoverTab[78728]"
		case <-ccb.closed.Done():
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:137
			_go_fuzz_dep_.CoverTab[78729]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:137
			// _ = "end of CoverTab[78729]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:138
		// _ = "end of CoverTab[78725]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:138
		_go_fuzz_dep_.CoverTab[78726]++

													if ccb.closed.HasFired() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:140
			_go_fuzz_dep_.CoverTab[78739]++
														ccb.handleClose()
														return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:142
			// _ = "end of CoverTab[78739]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:143
			_go_fuzz_dep_.CoverTab[78740]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:143
			// _ = "end of CoverTab[78740]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:143
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:143
		// _ = "end of CoverTab[78726]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:144
	// _ = "end of CoverTab[78724]"
}

// updateClientConnState is invoked by grpc to push a ClientConnState update to
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:147
// the underlying balancer.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:147
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:147
// Unlike other methods invoked by grpc to push updates to the underlying
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:147
// balancer, this method cannot simply push the update onto the update channel
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:147
// and return. It needs to return the error returned by the underlying balancer
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:147
// back to grpc which propagates that to the resolver.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:154
func (ccb *ccBalancerWrapper) updateClientConnState(ccs *balancer.ClientConnState) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:154
	_go_fuzz_dep_.CoverTab[78741]++
												ccb.updateCh.Put(&ccStateUpdate{ccs: ccs})

												var res interface{}
												select {
	case res = <-ccb.resultCh.Get():
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:159
		_go_fuzz_dep_.CoverTab[78744]++
													ccb.resultCh.Load()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:160
		// _ = "end of CoverTab[78744]"
	case <-ccb.closed.Done():
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:161
		_go_fuzz_dep_.CoverTab[78745]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:164
		return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:164
		// _ = "end of CoverTab[78745]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:165
	// _ = "end of CoverTab[78741]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:165
	_go_fuzz_dep_.CoverTab[78742]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:168
	if res == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:168
		_go_fuzz_dep_.CoverTab[78746]++
													return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:169
		// _ = "end of CoverTab[78746]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:170
		_go_fuzz_dep_.CoverTab[78747]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:170
		// _ = "end of CoverTab[78747]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:170
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:170
	// _ = "end of CoverTab[78742]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:170
	_go_fuzz_dep_.CoverTab[78743]++
												return res.(error)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:171
	// _ = "end of CoverTab[78743]"
}

// handleClientConnStateChange handles a ClientConnState update from the update
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:174
// channel and invokes the appropriate method on the underlying balancer.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:174
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:174
// If the addresses specified in the update contain addresses of type "grpclb"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:174
// and the selected LB policy is not "grpclb", these addresses will be filtered
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:174
// out and ccs will be modified with the updated address list.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:180
func (ccb *ccBalancerWrapper) handleClientConnStateChange(ccs *balancer.ClientConnState) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:180
	_go_fuzz_dep_.CoverTab[78748]++
												if ccb.curBalancerName != grpclbName {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:181
		_go_fuzz_dep_.CoverTab[78750]++
		// Filter any grpclb addresses since we don't have the grpclb balancer.
		var addrs []resolver.Address
		for _, addr := range ccs.ResolverState.Addresses {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:184
			_go_fuzz_dep_.CoverTab[78752]++
														if addr.Type == resolver.GRPCLB {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:185
				_go_fuzz_dep_.CoverTab[78754]++
															continue
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:186
				// _ = "end of CoverTab[78754]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:187
				_go_fuzz_dep_.CoverTab[78755]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:187
				// _ = "end of CoverTab[78755]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:187
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:187
			// _ = "end of CoverTab[78752]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:187
			_go_fuzz_dep_.CoverTab[78753]++
														addrs = append(addrs, addr)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:188
			// _ = "end of CoverTab[78753]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:189
		// _ = "end of CoverTab[78750]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:189
		_go_fuzz_dep_.CoverTab[78751]++
													ccs.ResolverState.Addresses = addrs
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:190
		// _ = "end of CoverTab[78751]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:191
		_go_fuzz_dep_.CoverTab[78756]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:191
		// _ = "end of CoverTab[78756]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:191
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:191
	// _ = "end of CoverTab[78748]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:191
	_go_fuzz_dep_.CoverTab[78749]++
												ccb.resultCh.Put(ccb.balancer.UpdateClientConnState(*ccs))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:192
	// _ = "end of CoverTab[78749]"
}

// updateSubConnState is invoked by grpc to push a subConn state update to the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:195
// underlying balancer.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:197
func (ccb *ccBalancerWrapper) updateSubConnState(sc balancer.SubConn, s connectivity.State, err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:197
	_go_fuzz_dep_.CoverTab[78757]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:205
	if sc == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:205
		_go_fuzz_dep_.CoverTab[78759]++
													return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:206
		// _ = "end of CoverTab[78759]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:207
		_go_fuzz_dep_.CoverTab[78760]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:207
		// _ = "end of CoverTab[78760]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:207
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:207
	// _ = "end of CoverTab[78757]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:207
	_go_fuzz_dep_.CoverTab[78758]++
												ccb.updateCh.Put(&scStateUpdate{
		sc:	sc,
		state:	s,
		err:	err,
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:212
	// _ = "end of CoverTab[78758]"
}

// handleSubConnStateChange handles a SubConnState update from the update
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:215
// channel and invokes the appropriate method on the underlying balancer.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:217
func (ccb *ccBalancerWrapper) handleSubConnStateChange(update *scStateUpdate) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:217
	_go_fuzz_dep_.CoverTab[78761]++
												ccb.balancer.UpdateSubConnState(update.sc, balancer.SubConnState{ConnectivityState: update.state, ConnectionError: update.err})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:218
	// _ = "end of CoverTab[78761]"
}

func (ccb *ccBalancerWrapper) exitIdle() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:221
	_go_fuzz_dep_.CoverTab[78762]++
												ccb.updateCh.Put(&exitIdleUpdate{})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:222
	// _ = "end of CoverTab[78762]"
}

func (ccb *ccBalancerWrapper) handleExitIdle() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:225
	_go_fuzz_dep_.CoverTab[78763]++
												if ccb.cc.GetState() != connectivity.Idle {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:226
		_go_fuzz_dep_.CoverTab[78765]++
													return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:227
		// _ = "end of CoverTab[78765]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:228
		_go_fuzz_dep_.CoverTab[78766]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:228
		// _ = "end of CoverTab[78766]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:228
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:228
	// _ = "end of CoverTab[78763]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:228
	_go_fuzz_dep_.CoverTab[78764]++
												ccb.balancer.ExitIdle()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:229
	// _ = "end of CoverTab[78764]"
}

func (ccb *ccBalancerWrapper) resolverError(err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:232
	_go_fuzz_dep_.CoverTab[78767]++
												ccb.updateCh.Put(&resolverErrorUpdate{err: err})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:233
	// _ = "end of CoverTab[78767]"
}

func (ccb *ccBalancerWrapper) handleResolverError(err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:236
	_go_fuzz_dep_.CoverTab[78768]++
												ccb.balancer.ResolverError(err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:237
	// _ = "end of CoverTab[78768]"
}

// switchTo is invoked by grpc to instruct the balancer wrapper to switch to the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:240
// LB policy identified by name.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:240
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:240
// ClientConn calls newCCBalancerWrapper() at creation time. Upon receipt of the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:240
// first good update from the name resolver, it determines the LB policy to use
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:240
// and invokes the switchTo() method. Upon receipt of every subsequent update
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:240
// from the name resolver, it invokes this method.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:240
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:240
// the ccBalancerWrapper keeps track of the current LB policy name, and skips
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:240
// the graceful balancer switching process if the name does not change.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:250
func (ccb *ccBalancerWrapper) switchTo(name string) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:250
	_go_fuzz_dep_.CoverTab[78769]++
												ccb.updateCh.Put(&switchToUpdate{name: name})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:251
	// _ = "end of CoverTab[78769]"
}

// handleSwitchTo handles a balancer switch update from the update channel. It
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:254
// calls the SwitchTo() method on the gracefulswitch.Balancer with a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:254
// balancer.Builder corresponding to name. If no balancer.Builder is registered
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:254
// for the given name, it uses the default LB policy which is "pick_first".
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:258
func (ccb *ccBalancerWrapper) handleSwitchTo(name string) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:258
	_go_fuzz_dep_.CoverTab[78770]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:261
	if strings.EqualFold(ccb.curBalancerName, name) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:261
		_go_fuzz_dep_.CoverTab[78774]++
													return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:262
		// _ = "end of CoverTab[78774]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:263
		_go_fuzz_dep_.CoverTab[78775]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:263
		// _ = "end of CoverTab[78775]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:263
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:263
	// _ = "end of CoverTab[78770]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:263
	_go_fuzz_dep_.CoverTab[78771]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:269
	builder := balancer.Get(name)
	if builder == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:270
		_go_fuzz_dep_.CoverTab[78776]++
													channelz.Warningf(logger, ccb.cc.channelzID, "Channel switches to new LB policy %q, since the specified LB policy %q was not registered", PickFirstBalancerName, name)
													builder = newPickfirstBuilder()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:272
		// _ = "end of CoverTab[78776]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:273
		_go_fuzz_dep_.CoverTab[78777]++
													channelz.Infof(logger, ccb.cc.channelzID, "Channel switches to new LB policy %q", name)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:274
		// _ = "end of CoverTab[78777]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:275
	// _ = "end of CoverTab[78771]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:275
	_go_fuzz_dep_.CoverTab[78772]++

												if err := ccb.balancer.SwitchTo(builder); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:277
		_go_fuzz_dep_.CoverTab[78778]++
													channelz.Errorf(logger, ccb.cc.channelzID, "Channel failed to build new LB policy %q: %v", name, err)
													return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:279
		// _ = "end of CoverTab[78778]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:280
		_go_fuzz_dep_.CoverTab[78779]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:280
		// _ = "end of CoverTab[78779]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:280
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:280
	// _ = "end of CoverTab[78772]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:280
	_go_fuzz_dep_.CoverTab[78773]++
												ccb.curBalancerName = builder.Name()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:281
	// _ = "end of CoverTab[78773]"
}

// handleRemoveSucConn handles a request from the underlying balancer to remove
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:284
// a subConn.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:284
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:284
// See comments in RemoveSubConn() for more details.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:288
func (ccb *ccBalancerWrapper) handleRemoveSubConn(acbw *acBalancerWrapper) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:288
	_go_fuzz_dep_.CoverTab[78780]++
												ccb.cc.removeAddrConn(acbw.getAddrConn(), errConnDrain)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:289
	// _ = "end of CoverTab[78780]"
}

func (ccb *ccBalancerWrapper) close() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:292
	_go_fuzz_dep_.CoverTab[78781]++
												ccb.closed.Fire()
												<-ccb.done.Done()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:294
	// _ = "end of CoverTab[78781]"
}

func (ccb *ccBalancerWrapper) handleClose() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:297
	_go_fuzz_dep_.CoverTab[78782]++
												ccb.balancer.Close()
												ccb.done.Fire()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:299
	// _ = "end of CoverTab[78782]"
}

func (ccb *ccBalancerWrapper) NewSubConn(addrs []resolver.Address, opts balancer.NewSubConnOptions) (balancer.SubConn, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:302
	_go_fuzz_dep_.CoverTab[78783]++
												if len(addrs) <= 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:303
		_go_fuzz_dep_.CoverTab[78786]++
													return nil, fmt.Errorf("grpc: cannot create SubConn with empty address list")
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:304
		// _ = "end of CoverTab[78786]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:305
		_go_fuzz_dep_.CoverTab[78787]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:305
		// _ = "end of CoverTab[78787]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:305
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:305
	// _ = "end of CoverTab[78783]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:305
	_go_fuzz_dep_.CoverTab[78784]++
												ac, err := ccb.cc.newAddrConn(addrs, opts)
												if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:307
		_go_fuzz_dep_.CoverTab[78788]++
													channelz.Warningf(logger, ccb.cc.channelzID, "acBalancerWrapper: NewSubConn: failed to newAddrConn: %v", err)
													return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:309
		// _ = "end of CoverTab[78788]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:310
		_go_fuzz_dep_.CoverTab[78789]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:310
		// _ = "end of CoverTab[78789]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:310
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:310
	// _ = "end of CoverTab[78784]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:310
	_go_fuzz_dep_.CoverTab[78785]++
												acbw := &acBalancerWrapper{ac: ac, producers: make(map[balancer.ProducerBuilder]*refCountedProducer)}
												acbw.ac.mu.Lock()
												ac.acbw = acbw
												acbw.ac.mu.Unlock()
												return acbw, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:315
	// _ = "end of CoverTab[78785]"
}

func (ccb *ccBalancerWrapper) RemoveSubConn(sc balancer.SubConn) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:318
	_go_fuzz_dep_.CoverTab[78790]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:329
	acbw, ok := sc.(*acBalancerWrapper)
	if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:330
		_go_fuzz_dep_.CoverTab[78792]++
													return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:331
		// _ = "end of CoverTab[78792]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:332
		_go_fuzz_dep_.CoverTab[78793]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:332
		// _ = "end of CoverTab[78793]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:332
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:332
	// _ = "end of CoverTab[78790]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:332
	_go_fuzz_dep_.CoverTab[78791]++
												ccb.updateCh.Put(&subConnUpdate{acbw: acbw})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:333
	// _ = "end of CoverTab[78791]"
}

func (ccb *ccBalancerWrapper) UpdateAddresses(sc balancer.SubConn, addrs []resolver.Address) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:336
	_go_fuzz_dep_.CoverTab[78794]++
												acbw, ok := sc.(*acBalancerWrapper)
												if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:338
		_go_fuzz_dep_.CoverTab[78796]++
													return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:339
		// _ = "end of CoverTab[78796]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:340
		_go_fuzz_dep_.CoverTab[78797]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:340
		// _ = "end of CoverTab[78797]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:340
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:340
	// _ = "end of CoverTab[78794]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:340
	_go_fuzz_dep_.CoverTab[78795]++
												acbw.UpdateAddresses(addrs)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:341
	// _ = "end of CoverTab[78795]"
}

func (ccb *ccBalancerWrapper) UpdateState(s balancer.State) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:344
	_go_fuzz_dep_.CoverTab[78798]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:350
	ccb.cc.blockingpicker.updatePicker(s.Picker)
												ccb.cc.csMgr.updateState(s.ConnectivityState)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:351
	// _ = "end of CoverTab[78798]"
}

func (ccb *ccBalancerWrapper) ResolveNow(o resolver.ResolveNowOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:354
	_go_fuzz_dep_.CoverTab[78799]++
												ccb.cc.resolveNow(o)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:355
	// _ = "end of CoverTab[78799]"
}

func (ccb *ccBalancerWrapper) Target() string {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:358
	_go_fuzz_dep_.CoverTab[78800]++
												return ccb.cc.target
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:359
	// _ = "end of CoverTab[78800]"
}

// acBalancerWrapper is a wrapper on top of ac for balancers.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:362
// It implements balancer.SubConn interface.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:364
type acBalancerWrapper struct {
	mu		sync.Mutex
	ac		*addrConn
	producers	map[balancer.ProducerBuilder]*refCountedProducer
}

func (acbw *acBalancerWrapper) UpdateAddresses(addrs []resolver.Address) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:370
	_go_fuzz_dep_.CoverTab[78801]++
												acbw.mu.Lock()
												defer acbw.mu.Unlock()
												if len(addrs) <= 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:373
		_go_fuzz_dep_.CoverTab[78803]++
													acbw.ac.cc.removeAddrConn(acbw.ac, errConnDrain)
													return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:375
		// _ = "end of CoverTab[78803]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:376
		_go_fuzz_dep_.CoverTab[78804]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:376
		// _ = "end of CoverTab[78804]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:376
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:376
	// _ = "end of CoverTab[78801]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:376
	_go_fuzz_dep_.CoverTab[78802]++
												if !acbw.ac.tryUpdateAddrs(addrs) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:377
		_go_fuzz_dep_.CoverTab[78805]++
													cc := acbw.ac.cc
													opts := acbw.ac.scopts
													acbw.ac.mu.Lock()

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:386
		acbw.ac.acbw = nil
		acbw.ac.mu.Unlock()
		acState := acbw.ac.getState()
		acbw.ac.cc.removeAddrConn(acbw.ac, errConnDrain)

		if acState == connectivity.Shutdown {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:391
			_go_fuzz_dep_.CoverTab[78808]++
														return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:392
			// _ = "end of CoverTab[78808]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:393
			_go_fuzz_dep_.CoverTab[78809]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:393
			// _ = "end of CoverTab[78809]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:393
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:393
		// _ = "end of CoverTab[78805]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:393
		_go_fuzz_dep_.CoverTab[78806]++

													newAC, err := cc.newAddrConn(addrs, opts)
													if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:396
			_go_fuzz_dep_.CoverTab[78810]++
														channelz.Warningf(logger, acbw.ac.channelzID, "acBalancerWrapper: UpdateAddresses: failed to newAddrConn: %v", err)
														return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:398
			// _ = "end of CoverTab[78810]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:399
			_go_fuzz_dep_.CoverTab[78811]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:399
			// _ = "end of CoverTab[78811]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:399
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:399
		// _ = "end of CoverTab[78806]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:399
		_go_fuzz_dep_.CoverTab[78807]++
													acbw.ac = newAC
													newAC.mu.Lock()
													newAC.acbw = acbw
													newAC.mu.Unlock()
													if acState != connectivity.Idle {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:404
			_go_fuzz_dep_.CoverTab[78812]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:404
			_curRoutineNum86_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:404
			_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum86_)
														go newAC.connect()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:405
			// _ = "end of CoverTab[78812]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:406
			_go_fuzz_dep_.CoverTab[78813]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:406
			// _ = "end of CoverTab[78813]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:406
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:406
		// _ = "end of CoverTab[78807]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:407
		_go_fuzz_dep_.CoverTab[78814]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:407
		// _ = "end of CoverTab[78814]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:407
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:407
	// _ = "end of CoverTab[78802]"
}

func (acbw *acBalancerWrapper) Connect() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:410
	_go_fuzz_dep_.CoverTab[78815]++
												acbw.mu.Lock()
												defer acbw.mu.Unlock()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:412
	_curRoutineNum87_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:412
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum87_)
												go acbw.ac.connect()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:413
	// _ = "end of CoverTab[78815]"
}

func (acbw *acBalancerWrapper) getAddrConn() *addrConn {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:416
	_go_fuzz_dep_.CoverTab[78816]++
												acbw.mu.Lock()
												defer acbw.mu.Unlock()
												return acbw.ac
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:419
	// _ = "end of CoverTab[78816]"
}

var errSubConnNotReady = status.Error(codes.Unavailable, "SubConn not currently connected")

// NewStream begins a streaming RPC on the addrConn.  If the addrConn is not
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:424
// ready, returns errSubConnNotReady.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:426
func (acbw *acBalancerWrapper) NewStream(ctx context.Context, desc *StreamDesc, method string, opts ...CallOption) (ClientStream, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:426
	_go_fuzz_dep_.CoverTab[78817]++
												transport := acbw.ac.getReadyTransport()
												if transport == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:428
		_go_fuzz_dep_.CoverTab[78819]++
													return nil, errSubConnNotReady
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:429
		// _ = "end of CoverTab[78819]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:430
		_go_fuzz_dep_.CoverTab[78820]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:430
		// _ = "end of CoverTab[78820]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:430
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:430
	// _ = "end of CoverTab[78817]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:430
	_go_fuzz_dep_.CoverTab[78818]++
												return newNonRetryClientStream(ctx, desc, method, transport, acbw.ac, opts...)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:431
	// _ = "end of CoverTab[78818]"
}

// Invoke performs a unary RPC.  If the addrConn is not ready, returns
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:434
// errSubConnNotReady.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:436
func (acbw *acBalancerWrapper) Invoke(ctx context.Context, method string, args interface{}, reply interface{}, opts ...CallOption) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:436
	_go_fuzz_dep_.CoverTab[78821]++
												cs, err := acbw.NewStream(ctx, unaryStreamDesc, method, opts...)
												if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:438
		_go_fuzz_dep_.CoverTab[78824]++
													return err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:439
		// _ = "end of CoverTab[78824]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:440
		_go_fuzz_dep_.CoverTab[78825]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:440
		// _ = "end of CoverTab[78825]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:440
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:440
	// _ = "end of CoverTab[78821]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:440
	_go_fuzz_dep_.CoverTab[78822]++
												if err := cs.SendMsg(args); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:441
		_go_fuzz_dep_.CoverTab[78826]++
													return err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:442
		// _ = "end of CoverTab[78826]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:443
		_go_fuzz_dep_.CoverTab[78827]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:443
		// _ = "end of CoverTab[78827]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:443
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:443
	// _ = "end of CoverTab[78822]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:443
	_go_fuzz_dep_.CoverTab[78823]++
												return cs.RecvMsg(reply)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:444
	// _ = "end of CoverTab[78823]"
}

type refCountedProducer struct {
	producer	balancer.Producer
	refs		int	// number of current refs to the producer
	close		func()	// underlying producer's close function
}

func (acbw *acBalancerWrapper) GetOrBuildProducer(pb balancer.ProducerBuilder) (balancer.Producer, func()) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:453
	_go_fuzz_dep_.CoverTab[78828]++
												acbw.mu.Lock()
												defer acbw.mu.Unlock()

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:458
	pData := acbw.producers[pb]
	if pData == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:459
		_go_fuzz_dep_.CoverTab[78831]++

													p, close := pb.Build(acbw)
													pData = &refCountedProducer{producer: p, close: close}
													acbw.producers[pb] = pData
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:463
		// _ = "end of CoverTab[78831]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:464
		_go_fuzz_dep_.CoverTab[78832]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:464
		// _ = "end of CoverTab[78832]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:464
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:464
	// _ = "end of CoverTab[78828]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:464
	_go_fuzz_dep_.CoverTab[78829]++

												pData.refs++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:471
	unref := func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:471
		_go_fuzz_dep_.CoverTab[78833]++
													acbw.mu.Lock()
													pData.refs--
													if pData.refs == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:474
			_go_fuzz_dep_.CoverTab[78835]++
														defer pData.close()
														delete(acbw.producers, pb)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:476
			// _ = "end of CoverTab[78835]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:477
			_go_fuzz_dep_.CoverTab[78836]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:477
			// _ = "end of CoverTab[78836]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:477
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:477
		// _ = "end of CoverTab[78833]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:477
		_go_fuzz_dep_.CoverTab[78834]++
													acbw.mu.Unlock()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:478
		// _ = "end of CoverTab[78834]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:479
	// _ = "end of CoverTab[78829]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:479
	_go_fuzz_dep_.CoverTab[78830]++
												return pData.producer, grpcsync.OnceFunc(unref)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:480
	// _ = "end of CoverTab[78830]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:481
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer_conn_wrappers.go:481
var _ = _go_fuzz_dep_.CoverTab
