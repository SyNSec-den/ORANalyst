//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/conn_state_evaluator.go:19
package balancer

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/conn_state_evaluator.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/conn_state_evaluator.go:19
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/conn_state_evaluator.go:19
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/conn_state_evaluator.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/conn_state_evaluator.go:19
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/conn_state_evaluator.go:19
)

import "google.golang.org/grpc/connectivity"

// ConnectivityStateEvaluator takes the connectivity states of multiple SubConns
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/conn_state_evaluator.go:23
// and returns one aggregated connectivity state.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/conn_state_evaluator.go:23
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/conn_state_evaluator.go:23
// It's not thread safe.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/conn_state_evaluator.go:27
type ConnectivityStateEvaluator struct {
	numReady		uint64	// Number of addrConns in ready state.
	numConnecting		uint64	// Number of addrConns in connecting state.
	numTransientFailure	uint64	// Number of addrConns in transient failure state.
	numIdle			uint64	// Number of addrConns in idle state.
}

// RecordTransition records state change happening in subConn and based on that
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/conn_state_evaluator.go:34
// it evaluates what aggregated state should be.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/conn_state_evaluator.go:34
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/conn_state_evaluator.go:34
//   - If at least one SubConn in Ready, the aggregated state is Ready;
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/conn_state_evaluator.go:34
//   - Else if at least one SubConn in Connecting, the aggregated state is Connecting;
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/conn_state_evaluator.go:34
//   - Else if at least one SubConn is Idle, the aggregated state is Idle;
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/conn_state_evaluator.go:34
//   - Else if at least one SubConn is TransientFailure (or there are no SubConns), the aggregated state is Transient Failure.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/conn_state_evaluator.go:34
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/conn_state_evaluator.go:34
// Shutdown is not considered.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/conn_state_evaluator.go:43
func (cse *ConnectivityStateEvaluator) RecordTransition(oldState, newState connectivity.State) connectivity.State {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/conn_state_evaluator.go:43
	_go_fuzz_dep_.CoverTab[67410]++

													for idx, state := range []connectivity.State{oldState, newState} {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/conn_state_evaluator.go:45
		_go_fuzz_dep_.CoverTab[67412]++
														updateVal := 2*uint64(idx) - 1
														switch state {
		case connectivity.Ready:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/conn_state_evaluator.go:48
			_go_fuzz_dep_.CoverTab[67413]++
															cse.numReady += updateVal
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/conn_state_evaluator.go:49
			// _ = "end of CoverTab[67413]"
		case connectivity.Connecting:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/conn_state_evaluator.go:50
			_go_fuzz_dep_.CoverTab[67414]++
															cse.numConnecting += updateVal
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/conn_state_evaluator.go:51
			// _ = "end of CoverTab[67414]"
		case connectivity.TransientFailure:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/conn_state_evaluator.go:52
			_go_fuzz_dep_.CoverTab[67415]++
															cse.numTransientFailure += updateVal
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/conn_state_evaluator.go:53
			// _ = "end of CoverTab[67415]"
		case connectivity.Idle:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/conn_state_evaluator.go:54
			_go_fuzz_dep_.CoverTab[67416]++
															cse.numIdle += updateVal
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/conn_state_evaluator.go:55
			// _ = "end of CoverTab[67416]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/conn_state_evaluator.go:55
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/conn_state_evaluator.go:55
			_go_fuzz_dep_.CoverTab[67417]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/conn_state_evaluator.go:55
			// _ = "end of CoverTab[67417]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/conn_state_evaluator.go:56
		// _ = "end of CoverTab[67412]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/conn_state_evaluator.go:57
	// _ = "end of CoverTab[67410]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/conn_state_evaluator.go:57
	_go_fuzz_dep_.CoverTab[67411]++
													return cse.CurrentState()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/conn_state_evaluator.go:58
	// _ = "end of CoverTab[67411]"
}

// CurrentState returns the current aggregate conn state by evaluating the counters
func (cse *ConnectivityStateEvaluator) CurrentState() connectivity.State {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/conn_state_evaluator.go:62
	_go_fuzz_dep_.CoverTab[67418]++

													if cse.numReady > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/conn_state_evaluator.go:64
		_go_fuzz_dep_.CoverTab[67422]++
														return connectivity.Ready
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/conn_state_evaluator.go:65
		// _ = "end of CoverTab[67422]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/conn_state_evaluator.go:66
		_go_fuzz_dep_.CoverTab[67423]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/conn_state_evaluator.go:66
		// _ = "end of CoverTab[67423]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/conn_state_evaluator.go:66
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/conn_state_evaluator.go:66
	// _ = "end of CoverTab[67418]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/conn_state_evaluator.go:66
	_go_fuzz_dep_.CoverTab[67419]++
													if cse.numConnecting > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/conn_state_evaluator.go:67
		_go_fuzz_dep_.CoverTab[67424]++
														return connectivity.Connecting
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/conn_state_evaluator.go:68
		// _ = "end of CoverTab[67424]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/conn_state_evaluator.go:69
		_go_fuzz_dep_.CoverTab[67425]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/conn_state_evaluator.go:69
		// _ = "end of CoverTab[67425]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/conn_state_evaluator.go:69
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/conn_state_evaluator.go:69
	// _ = "end of CoverTab[67419]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/conn_state_evaluator.go:69
	_go_fuzz_dep_.CoverTab[67420]++
													if cse.numIdle > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/conn_state_evaluator.go:70
		_go_fuzz_dep_.CoverTab[67426]++
														return connectivity.Idle
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/conn_state_evaluator.go:71
		// _ = "end of CoverTab[67426]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/conn_state_evaluator.go:72
		_go_fuzz_dep_.CoverTab[67427]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/conn_state_evaluator.go:72
		// _ = "end of CoverTab[67427]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/conn_state_evaluator.go:72
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/conn_state_evaluator.go:72
	// _ = "end of CoverTab[67420]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/conn_state_evaluator.go:72
	_go_fuzz_dep_.CoverTab[67421]++
													return connectivity.TransientFailure
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/conn_state_evaluator.go:73
	// _ = "end of CoverTab[67421]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/conn_state_evaluator.go:74
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/conn_state_evaluator.go:74
var _ = _go_fuzz_dep_.CoverTab
