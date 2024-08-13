//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/connectivity/connectivity.go:19
// Package connectivity defines connectivity semantics.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/connectivity/connectivity.go:19
// For details, see https://github.com/grpc/grpc/blob/master/doc/connectivity-semantics-and-api.md.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/connectivity/connectivity.go:21
package connectivity

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/connectivity/connectivity.go:21
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/connectivity/connectivity.go:21
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/connectivity/connectivity.go:21
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/connectivity/connectivity.go:21
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/connectivity/connectivity.go:21
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/connectivity/connectivity.go:21
)

import (
	"google.golang.org/grpc/grpclog"
)

var logger = grpclog.Component("core")

// State indicates the state of connectivity.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/connectivity/connectivity.go:29
// It can be the state of a ClientConn or SubConn.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/connectivity/connectivity.go:31
type State int

func (s State) String() string {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/connectivity/connectivity.go:33
	_go_fuzz_dep_.CoverTab[48166]++
													switch s {
	case Idle:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/connectivity/connectivity.go:35
		_go_fuzz_dep_.CoverTab[48167]++
														return "IDLE"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/connectivity/connectivity.go:36
		// _ = "end of CoverTab[48167]"
	case Connecting:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/connectivity/connectivity.go:37
		_go_fuzz_dep_.CoverTab[48168]++
														return "CONNECTING"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/connectivity/connectivity.go:38
		// _ = "end of CoverTab[48168]"
	case Ready:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/connectivity/connectivity.go:39
		_go_fuzz_dep_.CoverTab[48169]++
														return "READY"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/connectivity/connectivity.go:40
		// _ = "end of CoverTab[48169]"
	case TransientFailure:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/connectivity/connectivity.go:41
		_go_fuzz_dep_.CoverTab[48170]++
														return "TRANSIENT_FAILURE"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/connectivity/connectivity.go:42
		// _ = "end of CoverTab[48170]"
	case Shutdown:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/connectivity/connectivity.go:43
		_go_fuzz_dep_.CoverTab[48171]++
														return "SHUTDOWN"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/connectivity/connectivity.go:44
		// _ = "end of CoverTab[48171]"
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/connectivity/connectivity.go:45
		_go_fuzz_dep_.CoverTab[48172]++
														logger.Errorf("unknown connectivity state: %d", s)
														return "INVALID_STATE"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/connectivity/connectivity.go:47
		// _ = "end of CoverTab[48172]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/connectivity/connectivity.go:48
	// _ = "end of CoverTab[48166]"
}

const (
	// Idle indicates the ClientConn is idle.
	Idle	State	= iota
	// Connecting indicates the ClientConn is connecting.
	Connecting
	// Ready indicates the ClientConn is ready for work.
	Ready
	// TransientFailure indicates the ClientConn has seen a failure but expects to recover.
	TransientFailure
	// Shutdown indicates the ClientConn has started shutting down.
	Shutdown
)

// ServingMode indicates the current mode of operation of the server.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/connectivity/connectivity.go:64
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/connectivity/connectivity.go:64
// Only xDS enabled gRPC servers currently report their serving mode.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/connectivity/connectivity.go:67
type ServingMode int

const (
	// ServingModeStarting indicates that the server is starting up.
	ServingModeStarting	ServingMode	= iota
	// ServingModeServing indicates that the server contains all required
	// configuration and is serving RPCs.
	ServingModeServing
	// ServingModeNotServing indicates that the server is not accepting new
	// connections. Existing connections will be closed gracefully, allowing
	// in-progress RPCs to complete. A server enters this mode when it does not
	// contain the required configuration to serve RPCs.
	ServingModeNotServing
)

func (s ServingMode) String() string {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/connectivity/connectivity.go:82
	_go_fuzz_dep_.CoverTab[48173]++
													switch s {
	case ServingModeStarting:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/connectivity/connectivity.go:84
		_go_fuzz_dep_.CoverTab[48174]++
														return "STARTING"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/connectivity/connectivity.go:85
		// _ = "end of CoverTab[48174]"
	case ServingModeServing:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/connectivity/connectivity.go:86
		_go_fuzz_dep_.CoverTab[48175]++
														return "SERVING"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/connectivity/connectivity.go:87
		// _ = "end of CoverTab[48175]"
	case ServingModeNotServing:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/connectivity/connectivity.go:88
		_go_fuzz_dep_.CoverTab[48176]++
														return "NOT_SERVING"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/connectivity/connectivity.go:89
		// _ = "end of CoverTab[48176]"
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/connectivity/connectivity.go:90
		_go_fuzz_dep_.CoverTab[48177]++
														logger.Errorf("unknown serving mode: %d", s)
														return "INVALID_MODE"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/connectivity/connectivity.go:92
		// _ = "end of CoverTab[48177]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/connectivity/connectivity.go:93
	// _ = "end of CoverTab[48173]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/connectivity/connectivity.go:94
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/connectivity/connectivity.go:94
var _ = _go_fuzz_dep_.CoverTab
