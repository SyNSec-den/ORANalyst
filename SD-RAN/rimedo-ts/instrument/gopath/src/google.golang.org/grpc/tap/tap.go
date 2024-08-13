//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/tap/tap.go:19
// Package tap defines the function handles which are executed on the transport
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/tap/tap.go:19
// layer of gRPC-Go and related information.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/tap/tap.go:19
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/tap/tap.go:19
// # Experimental
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/tap/tap.go:19
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/tap/tap.go:19
// Notice: This API is EXPERIMENTAL and may be changed or removed in a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/tap/tap.go:19
// later release.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/tap/tap.go:26
package tap

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/tap/tap.go:26
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/tap/tap.go:26
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/tap/tap.go:26
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/tap/tap.go:26
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/tap/tap.go:26
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/tap/tap.go:26
)

import (
	"context"
)

// Info defines the relevant information needed by the handles.
type Info struct {
	// FullMethodName is the string of grpc method (in the format of
	// /package.service/method).
	FullMethodName string
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/tap/tap.go:38
}

// ServerInHandle defines the function which runs before a new stream is
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/tap/tap.go:40
// created on the server side. If it returns a non-nil error, the stream will
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/tap/tap.go:40
// not be created and an error will be returned to the client.  If the error
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/tap/tap.go:40
// returned is a status error, that status code and message will be used,
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/tap/tap.go:40
// otherwise PermissionDenied will be the code and err.Error() will be the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/tap/tap.go:40
// message.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/tap/tap.go:40
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/tap/tap.go:40
// It's intended to be used in situations where you don't want to waste the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/tap/tap.go:40
// resources to accept the new stream (e.g. rate-limiting). For other general
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/tap/tap.go:40
// usages, please use interceptors.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/tap/tap.go:40
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/tap/tap.go:40
// Note that it is executed in the per-connection I/O goroutine(s) instead of
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/tap/tap.go:40
// per-RPC goroutine. Therefore, users should NOT have any
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/tap/tap.go:40
// blocking/time-consuming work in this handle. Otherwise all the RPCs would
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/tap/tap.go:40
// slow down. Also, for the same reason, this handle won't be called
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/tap/tap.go:40
// concurrently by gRPC.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/tap/tap.go:56
type ServerInHandle func(ctx context.Context, info *Info) (context.Context, error)

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/tap/tap.go:56
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/tap/tap.go:56
var _ = _go_fuzz_dep_.CoverTab
