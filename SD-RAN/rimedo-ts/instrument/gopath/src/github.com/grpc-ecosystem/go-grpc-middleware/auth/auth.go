// Copyright 2016 Michal Witkowski. All Rights Reserved.
// See LICENSE for licensing terms.

//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/auth.go:4
package grpc_auth

//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/auth.go:4
import (
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/auth.go:4
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/auth.go:4
)
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/auth.go:4
import (
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/auth.go:4
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/auth.go:4
)

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
)

// AuthFunc is the pluggable function that performs authentication.
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/auth.go:13
//
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/auth.go:13
// The passed in `Context` will contain the gRPC metadata.MD object (for header-based authentication) and
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/auth.go:13
// the peer.Peer information that can contain transport-based credentials (e.g. `credentials.AuthInfo`).
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/auth.go:13
//
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/auth.go:13
// The returned context will be propagated to handlers, allowing user changes to `Context`. However,
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/auth.go:13
// please make sure that the `Context` returned is a child `Context` of the one passed in.
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/auth.go:13
//
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/auth.go:13
// If error is returned, its `grpc.Code()` will be returned to the user as well as the verbatim message.
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/auth.go:13
// Please make sure you use `codes.Unauthenticated` (lacking auth) and `codes.PermissionDenied`
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/auth.go:13
// (authed, but lacking perms) appropriately.
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/auth.go:24
type AuthFunc func(ctx context.Context) (context.Context, error)

// ServiceAuthFuncOverride allows a given gRPC service implementation to override the global `AuthFunc`.
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/auth.go:26
//
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/auth.go:26
// If a service implements the AuthFuncOverride method, it takes precedence over the `AuthFunc` method,
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/auth.go:26
// and will be called instead of AuthFunc for all method invocations within that service.
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/auth.go:30
type ServiceAuthFuncOverride interface {
	AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error)
}

// UnaryServerInterceptor returns a new unary server interceptors that performs per-request auth.
func UnaryServerInterceptor(authFunc AuthFunc) grpc.UnaryServerInterceptor {
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/auth.go:35
	_go_fuzz_dep_.CoverTab[183679]++
													return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/auth.go:36
		_go_fuzz_dep_.CoverTab[183680]++
														var newCtx context.Context
														var err error
														if overrideSrv, ok := info.Server.(ServiceAuthFuncOverride); ok {
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/auth.go:39
			_go_fuzz_dep_.CoverTab[183683]++
															newCtx, err = overrideSrv.AuthFuncOverride(ctx, info.FullMethod)
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/auth.go:40
			// _ = "end of CoverTab[183683]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/auth.go:41
			_go_fuzz_dep_.CoverTab[183684]++
															newCtx, err = authFunc(ctx)
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/auth.go:42
			// _ = "end of CoverTab[183684]"
		}
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/auth.go:43
		// _ = "end of CoverTab[183680]"
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/auth.go:43
		_go_fuzz_dep_.CoverTab[183681]++
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/auth.go:44
			_go_fuzz_dep_.CoverTab[183685]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/auth.go:45
			// _ = "end of CoverTab[183685]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/auth.go:46
			_go_fuzz_dep_.CoverTab[183686]++
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/auth.go:46
			// _ = "end of CoverTab[183686]"
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/auth.go:46
		}
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/auth.go:46
		// _ = "end of CoverTab[183681]"
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/auth.go:46
		_go_fuzz_dep_.CoverTab[183682]++
														return handler(newCtx, req)
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/auth.go:47
		// _ = "end of CoverTab[183682]"
	}
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/auth.go:48
	// _ = "end of CoverTab[183679]"
}

// StreamServerInterceptor returns a new unary server interceptors that performs per-request auth.
func StreamServerInterceptor(authFunc AuthFunc) grpc.StreamServerInterceptor {
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/auth.go:52
	_go_fuzz_dep_.CoverTab[183687]++
													return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/auth.go:53
		_go_fuzz_dep_.CoverTab[183688]++
														var newCtx context.Context
														var err error
														if overrideSrv, ok := srv.(ServiceAuthFuncOverride); ok {
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/auth.go:56
			_go_fuzz_dep_.CoverTab[183691]++
															newCtx, err = overrideSrv.AuthFuncOverride(stream.Context(), info.FullMethod)
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/auth.go:57
			// _ = "end of CoverTab[183691]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/auth.go:58
			_go_fuzz_dep_.CoverTab[183692]++
															newCtx, err = authFunc(stream.Context())
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/auth.go:59
			// _ = "end of CoverTab[183692]"
		}
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/auth.go:60
		// _ = "end of CoverTab[183688]"
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/auth.go:60
		_go_fuzz_dep_.CoverTab[183689]++
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/auth.go:61
			_go_fuzz_dep_.CoverTab[183693]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/auth.go:62
			// _ = "end of CoverTab[183693]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/auth.go:63
			_go_fuzz_dep_.CoverTab[183694]++
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/auth.go:63
			// _ = "end of CoverTab[183694]"
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/auth.go:63
		}
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/auth.go:63
		// _ = "end of CoverTab[183689]"
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/auth.go:63
		_go_fuzz_dep_.CoverTab[183690]++
														wrapped := grpc_middleware.WrapServerStream(stream)
														wrapped.WrappedContext = newCtx
														return handler(srv, wrapped)
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/auth.go:66
		// _ = "end of CoverTab[183690]"
	}
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/auth.go:67
	// _ = "end of CoverTab[183687]"
}

//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/auth.go:68
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/auth.go:68
var _ = _go_fuzz_dep_.CoverTab
