// Copyright 2016 Michal Witkowski. All Rights Reserved.
// See LICENSE for licensing terms.

// gRPC Server Interceptor chaining middleware.

//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:6
package grpc_middleware

//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:6
import (
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:6
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:6
)
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:6
import (
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:6
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:6
)

import (
	"context"

	"google.golang.org/grpc"
)

// ChainUnaryServer creates a single interceptor out of a chain of many interceptors.
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:14
//
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:14
// Execution is done in left-to-right order, including passing of context.
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:14
// For example ChainUnaryServer(one, two, three) will execute one before two before three, and three
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:14
// will see context changes of one and two.
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:19
func ChainUnaryServer(interceptors ...grpc.UnaryServerInterceptor) grpc.UnaryServerInterceptor {
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:19
	_go_fuzz_dep_.CoverTab[183615]++
													n := len(interceptors)

													return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:22
		_go_fuzz_dep_.CoverTab[183616]++
														chainer := func(currentInter grpc.UnaryServerInterceptor, currentHandler grpc.UnaryHandler) grpc.UnaryHandler {
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:23
			_go_fuzz_dep_.CoverTab[183619]++
															return func(currentCtx context.Context, currentReq interface{}) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:24
				_go_fuzz_dep_.CoverTab[183620]++
																return currentInter(currentCtx, currentReq, info, currentHandler)
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:25
				// _ = "end of CoverTab[183620]"
			}
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:26
			// _ = "end of CoverTab[183619]"
		}
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:27
		// _ = "end of CoverTab[183616]"
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:27
		_go_fuzz_dep_.CoverTab[183617]++

														chainedHandler := handler
														for i := n - 1; i >= 0; i-- {
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:30
			_go_fuzz_dep_.CoverTab[183621]++
															chainedHandler = chainer(interceptors[i], chainedHandler)
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:31
			// _ = "end of CoverTab[183621]"
		}
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:32
		// _ = "end of CoverTab[183617]"
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:32
		_go_fuzz_dep_.CoverTab[183618]++

														return chainedHandler(ctx, req)
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:34
		// _ = "end of CoverTab[183618]"
	}
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:35
	// _ = "end of CoverTab[183615]"
}

// ChainStreamServer creates a single interceptor out of a chain of many interceptors.
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:38
//
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:38
// Execution is done in left-to-right order, including passing of context.
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:38
// For example ChainUnaryServer(one, two, three) will execute one before two before three.
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:38
// If you want to pass context between interceptors, use WrapServerStream.
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:43
func ChainStreamServer(interceptors ...grpc.StreamServerInterceptor) grpc.StreamServerInterceptor {
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:43
	_go_fuzz_dep_.CoverTab[183622]++
													n := len(interceptors)

													return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:46
		_go_fuzz_dep_.CoverTab[183623]++
														chainer := func(currentInter grpc.StreamServerInterceptor, currentHandler grpc.StreamHandler) grpc.StreamHandler {
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:47
			_go_fuzz_dep_.CoverTab[183626]++
															return func(currentSrv interface{}, currentStream grpc.ServerStream) error {
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:48
				_go_fuzz_dep_.CoverTab[183627]++
																return currentInter(currentSrv, currentStream, info, currentHandler)
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:49
				// _ = "end of CoverTab[183627]"
			}
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:50
			// _ = "end of CoverTab[183626]"
		}
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:51
		// _ = "end of CoverTab[183623]"
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:51
		_go_fuzz_dep_.CoverTab[183624]++

														chainedHandler := handler
														for i := n - 1; i >= 0; i-- {
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:54
			_go_fuzz_dep_.CoverTab[183628]++
															chainedHandler = chainer(interceptors[i], chainedHandler)
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:55
			// _ = "end of CoverTab[183628]"
		}
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:56
		// _ = "end of CoverTab[183624]"
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:56
		_go_fuzz_dep_.CoverTab[183625]++

														return chainedHandler(srv, ss)
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:58
		// _ = "end of CoverTab[183625]"
	}
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:59
	// _ = "end of CoverTab[183622]"
}

// ChainUnaryClient creates a single interceptor out of a chain of many interceptors.
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:62
//
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:62
// Execution is done in left-to-right order, including passing of context.
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:62
// For example ChainUnaryClient(one, two, three) will execute one before two before three.
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:66
func ChainUnaryClient(interceptors ...grpc.UnaryClientInterceptor) grpc.UnaryClientInterceptor {
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:66
	_go_fuzz_dep_.CoverTab[183629]++
													n := len(interceptors)

													return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:69
		_go_fuzz_dep_.CoverTab[183630]++
														chainer := func(currentInter grpc.UnaryClientInterceptor, currentInvoker grpc.UnaryInvoker) grpc.UnaryInvoker {
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:70
			_go_fuzz_dep_.CoverTab[183633]++
															return func(currentCtx context.Context, currentMethod string, currentReq, currentRepl interface{}, currentConn *grpc.ClientConn, currentOpts ...grpc.CallOption) error {
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:71
				_go_fuzz_dep_.CoverTab[183634]++
																return currentInter(currentCtx, currentMethod, currentReq, currentRepl, currentConn, currentInvoker, currentOpts...)
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:72
				// _ = "end of CoverTab[183634]"
			}
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:73
			// _ = "end of CoverTab[183633]"
		}
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:74
		// _ = "end of CoverTab[183630]"
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:74
		_go_fuzz_dep_.CoverTab[183631]++

														chainedInvoker := invoker
														for i := n - 1; i >= 0; i-- {
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:77
			_go_fuzz_dep_.CoverTab[183635]++
															chainedInvoker = chainer(interceptors[i], chainedInvoker)
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:78
			// _ = "end of CoverTab[183635]"
		}
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:79
		// _ = "end of CoverTab[183631]"
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:79
		_go_fuzz_dep_.CoverTab[183632]++

														return chainedInvoker(ctx, method, req, reply, cc, opts...)
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:81
		// _ = "end of CoverTab[183632]"
	}
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:82
	// _ = "end of CoverTab[183629]"
}

// ChainStreamClient creates a single interceptor out of a chain of many interceptors.
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:85
//
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:85
// Execution is done in left-to-right order, including passing of context.
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:85
// For example ChainStreamClient(one, two, three) will execute one before two before three.
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:89
func ChainStreamClient(interceptors ...grpc.StreamClientInterceptor) grpc.StreamClientInterceptor {
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:89
	_go_fuzz_dep_.CoverTab[183636]++
													n := len(interceptors)

													return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:92
		_go_fuzz_dep_.CoverTab[183637]++
														chainer := func(currentInter grpc.StreamClientInterceptor, currentStreamer grpc.Streamer) grpc.Streamer {
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:93
			_go_fuzz_dep_.CoverTab[183640]++
															return func(currentCtx context.Context, currentDesc *grpc.StreamDesc, currentConn *grpc.ClientConn, currentMethod string, currentOpts ...grpc.CallOption) (grpc.ClientStream, error) {
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:94
				_go_fuzz_dep_.CoverTab[183641]++
																return currentInter(currentCtx, currentDesc, currentConn, currentMethod, currentStreamer, currentOpts...)
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:95
				// _ = "end of CoverTab[183641]"
			}
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:96
			// _ = "end of CoverTab[183640]"
		}
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:97
		// _ = "end of CoverTab[183637]"
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:97
		_go_fuzz_dep_.CoverTab[183638]++

														chainedStreamer := streamer
														for i := n - 1; i >= 0; i-- {
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:100
			_go_fuzz_dep_.CoverTab[183642]++
															chainedStreamer = chainer(interceptors[i], chainedStreamer)
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:101
			// _ = "end of CoverTab[183642]"
		}
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:102
		// _ = "end of CoverTab[183638]"
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:102
		_go_fuzz_dep_.CoverTab[183639]++

														return chainedStreamer(ctx, desc, cc, method, opts...)
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:104
		// _ = "end of CoverTab[183639]"
	}
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:105
	// _ = "end of CoverTab[183636]"
}

// Chain creates a single interceptor out of a chain of many interceptors.
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:108
//
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:108
// WithUnaryServerChain is a grpc.Server config option that accepts multiple unary interceptors.
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:108
// Basically syntactic sugar.
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:112
func WithUnaryServerChain(interceptors ...grpc.UnaryServerInterceptor) grpc.ServerOption {
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:112
	_go_fuzz_dep_.CoverTab[183643]++
													return grpc.UnaryInterceptor(ChainUnaryServer(interceptors...))
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:113
	// _ = "end of CoverTab[183643]"
}

// WithStreamServerChain is a grpc.Server config option that accepts multiple stream interceptors.
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:116
// Basically syntactic sugar.
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:118
func WithStreamServerChain(interceptors ...grpc.StreamServerInterceptor) grpc.ServerOption {
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:118
	_go_fuzz_dep_.CoverTab[183644]++
													return grpc.StreamInterceptor(ChainStreamServer(interceptors...))
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:119
	// _ = "end of CoverTab[183644]"
}

//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:120
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/chain.go:120
var _ = _go_fuzz_dep_.CoverTab
