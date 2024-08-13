//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/interceptor.go:19
package grpc

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/interceptor.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/interceptor.go:19
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/interceptor.go:19
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/interceptor.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/interceptor.go:19
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/interceptor.go:19
)

import (
	"context"
)

// UnaryInvoker is called by UnaryClientInterceptor to complete RPCs.
type UnaryInvoker func(ctx context.Context, method string, req, reply interface{}, cc *ClientConn, opts ...CallOption) error

// UnaryClientInterceptor intercepts the execution of a unary RPC on the client.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/interceptor.go:28
// Unary interceptors can be specified as a DialOption, using
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/interceptor.go:28
// WithUnaryInterceptor() or WithChainUnaryInterceptor(), when creating a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/interceptor.go:28
// ClientConn. When a unary interceptor(s) is set on a ClientConn, gRPC
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/interceptor.go:28
// delegates all unary RPC invocations to the interceptor, and it is the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/interceptor.go:28
// responsibility of the interceptor to call invoker to complete the processing
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/interceptor.go:28
// of the RPC.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/interceptor.go:28
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/interceptor.go:28
// method is the RPC name. req and reply are the corresponding request and
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/interceptor.go:28
// response messages. cc is the ClientConn on which the RPC was invoked. invoker
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/interceptor.go:28
// is the handler to complete the RPC and it is the responsibility of the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/interceptor.go:28
// interceptor to call it. opts contain all applicable call options, including
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/interceptor.go:28
// defaults from the ClientConn as well as per-call options.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/interceptor.go:28
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/interceptor.go:28
// The returned error must be compatible with the status package.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/interceptor.go:43
type UnaryClientInterceptor func(ctx context.Context, method string, req, reply interface{}, cc *ClientConn, invoker UnaryInvoker, opts ...CallOption) error

// Streamer is called by StreamClientInterceptor to create a ClientStream.
type Streamer func(ctx context.Context, desc *StreamDesc, cc *ClientConn, method string, opts ...CallOption) (ClientStream, error)

// StreamClientInterceptor intercepts the creation of a ClientStream. Stream
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/interceptor.go:48
// interceptors can be specified as a DialOption, using WithStreamInterceptor()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/interceptor.go:48
// or WithChainStreamInterceptor(), when creating a ClientConn. When a stream
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/interceptor.go:48
// interceptor(s) is set on the ClientConn, gRPC delegates all stream creations
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/interceptor.go:48
// to the interceptor, and it is the responsibility of the interceptor to call
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/interceptor.go:48
// streamer.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/interceptor.go:48
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/interceptor.go:48
// desc contains a description of the stream. cc is the ClientConn on which the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/interceptor.go:48
// RPC was invoked. streamer is the handler to create a ClientStream and it is
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/interceptor.go:48
// the responsibility of the interceptor to call it. opts contain all applicable
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/interceptor.go:48
// call options, including defaults from the ClientConn as well as per-call
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/interceptor.go:48
// options.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/interceptor.go:48
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/interceptor.go:48
// StreamClientInterceptor may return a custom ClientStream to intercept all I/O
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/interceptor.go:48
// operations. The returned error must be compatible with the status package.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/interceptor.go:63
type StreamClientInterceptor func(ctx context.Context, desc *StreamDesc, cc *ClientConn, method string, streamer Streamer, opts ...CallOption) (ClientStream, error)

// UnaryServerInfo consists of various information about a unary RPC on
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/interceptor.go:65
// server side. All per-rpc information may be mutated by the interceptor.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/interceptor.go:67
type UnaryServerInfo struct {
	// Server is the service implementation the user provides. This is read-only.
	Server	interface{}
	// FullMethod is the full RPC method string, i.e., /package.service/method.
	FullMethod	string
}

// UnaryHandler defines the handler invoked by UnaryServerInterceptor to complete the normal
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/interceptor.go:74
// execution of a unary RPC.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/interceptor.go:74
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/interceptor.go:74
// If a UnaryHandler returns an error, it should either be produced by the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/interceptor.go:74
// status package, or be one of the context errors. Otherwise, gRPC will use
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/interceptor.go:74
// codes.Unknown as the status code and err.Error() as the status message of the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/interceptor.go:74
// RPC.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/interceptor.go:81
type UnaryHandler func(ctx context.Context, req interface{}) (interface{}, error)

// UnaryServerInterceptor provides a hook to intercept the execution of a unary RPC on the server. info
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/interceptor.go:83
// contains all the information of this RPC the interceptor can operate on. And handler is the wrapper
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/interceptor.go:83
// of the service method implementation. It is the responsibility of the interceptor to invoke handler
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/interceptor.go:83
// to complete the RPC.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/interceptor.go:87
type UnaryServerInterceptor func(ctx context.Context, req interface{}, info *UnaryServerInfo, handler UnaryHandler) (resp interface{}, err error)

// StreamServerInfo consists of various information about a streaming RPC on
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/interceptor.go:89
// server side. All per-rpc information may be mutated by the interceptor.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/interceptor.go:91
type StreamServerInfo struct {
	// FullMethod is the full RPC method string, i.e., /package.service/method.
	FullMethod	string
	// IsClientStream indicates whether the RPC is a client streaming RPC.
	IsClientStream	bool
	// IsServerStream indicates whether the RPC is a server streaming RPC.
	IsServerStream	bool
}

// StreamServerInterceptor provides a hook to intercept the execution of a streaming RPC on the server.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/interceptor.go:100
// info contains all the information of this RPC the interceptor can operate on. And handler is the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/interceptor.go:100
// service method implementation. It is the responsibility of the interceptor to invoke handler to
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/interceptor.go:100
// complete the RPC.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/interceptor.go:104
type StreamServerInterceptor func(srv interface{}, ss ServerStream, info *StreamServerInfo, handler StreamHandler) error

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/interceptor.go:104
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/interceptor.go:104
var _ = _go_fuzz_dep_.CoverTab
