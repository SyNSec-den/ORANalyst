//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/handlers.go:19
package stats

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/handlers.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/handlers.go:19
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/handlers.go:19
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/handlers.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/handlers.go:19
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/handlers.go:19
)

import (
	"context"
	"net"
)

// ConnTagInfo defines the relevant information needed by connection context tagger.
type ConnTagInfo struct {
	// RemoteAddr is the remote address of the corresponding connection.
	RemoteAddr	net.Addr
	// LocalAddr is the local address of the corresponding connection.
	LocalAddr	net.Addr
}

// RPCTagInfo defines the relevant information needed by RPC context tagger.
type RPCTagInfo struct {
	// FullMethodName is the RPC method in the format of /package.service/method.
	FullMethodName	string
	// FailFast indicates if this RPC is failfast.
	// This field is only valid on client side, it's always false on server side.
	FailFast	bool
}

// Handler defines the interface for the related stats handling (e.g., RPCs, connections).
type Handler interface {
	// TagRPC can attach some information to the given context.
	// The context used for the rest lifetime of the RPC will be derived from
	// the returned context.
	TagRPC(context.Context, *RPCTagInfo) context.Context
	// HandleRPC processes the RPC stats.
	HandleRPC(context.Context, RPCStats)

	// TagConn can attach some information to the given context.
	// The returned context will be used for stats handling.
	// For conn stats handling, the context used in HandleConn for this
	// connection will be derived from the context returned.
	// For RPC stats handling,
	//  - On server side, the context used in HandleRPC for all RPCs on this
	// connection will be derived from the context returned.
	//  - On client side, the context is not derived from the context returned.
	TagConn(context.Context, *ConnTagInfo) context.Context
	// HandleConn processes the Conn stats.
	HandleConn(context.Context, ConnStats)
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/handlers.go:63
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/handlers.go:63
var _ = _go_fuzz_dep_.CoverTab
