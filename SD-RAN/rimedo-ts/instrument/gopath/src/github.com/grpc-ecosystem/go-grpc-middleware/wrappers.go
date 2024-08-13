// Copyright 2016 Michal Witkowski. All Rights Reserved.
// See LICENSE for licensing terms.

//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/wrappers.go:4
package grpc_middleware

//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/wrappers.go:4
import (
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/wrappers.go:4
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/wrappers.go:4
)
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/wrappers.go:4
import (
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/wrappers.go:4
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/wrappers.go:4
)

import (
	"context"

	"google.golang.org/grpc"
)

// WrappedServerStream is a thin wrapper around grpc.ServerStream that allows modifying context.
type WrappedServerStream struct {
	grpc.ServerStream
	// WrappedContext is the wrapper's own Context. You can assign it.
	WrappedContext	context.Context
}

// Context returns the wrapper's WrappedContext, overwriting the nested grpc.ServerStream.Context()
func (w *WrappedServerStream) Context() context.Context {
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/wrappers.go:20
	_go_fuzz_dep_.CoverTab[183645]++
													return w.WrappedContext
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/wrappers.go:21
	// _ = "end of CoverTab[183645]"
}

// WrapServerStream returns a ServerStream that has the ability to overwrite context.
func WrapServerStream(stream grpc.ServerStream) *WrappedServerStream {
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/wrappers.go:25
	_go_fuzz_dep_.CoverTab[183646]++
													if existing, ok := stream.(*WrappedServerStream); ok {
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/wrappers.go:26
		_go_fuzz_dep_.CoverTab[183648]++
														return existing
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/wrappers.go:27
		// _ = "end of CoverTab[183648]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/wrappers.go:28
		_go_fuzz_dep_.CoverTab[183649]++
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/wrappers.go:28
		// _ = "end of CoverTab[183649]"
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/wrappers.go:28
	}
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/wrappers.go:28
	// _ = "end of CoverTab[183646]"
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/wrappers.go:28
	_go_fuzz_dep_.CoverTab[183647]++
													return &WrappedServerStream{ServerStream: stream, WrappedContext: stream.Context()}
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/wrappers.go:29
	// _ = "end of CoverTab[183647]"
}

//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/wrappers.go:30
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/wrappers.go:30
var _ = _go_fuzz_dep_.CoverTab
