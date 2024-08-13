//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/channelz/channelz.go:19
// Package channelz exports internals of the channelz implementation as required
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/channelz/channelz.go:19
// by other gRPC packages.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/channelz/channelz.go:19
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/channelz/channelz.go:19
// The implementation of the channelz spec as defined in
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/channelz/channelz.go:19
// https://github.com/grpc/proposal/blob/master/A14-channelz.md, is provided by
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/channelz/channelz.go:19
// the `internal/channelz` package.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/channelz/channelz.go:19
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/channelz/channelz.go:19
// # Experimental
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/channelz/channelz.go:19
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/channelz/channelz.go:19
// Notice: All APIs in this package are experimental and may be removed in a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/channelz/channelz.go:19
// later release.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/channelz/channelz.go:30
package channelz

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/channelz/channelz.go:30
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/channelz/channelz.go:30
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/channelz/channelz.go:30
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/channelz/channelz.go:30
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/channelz/channelz.go:30
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/channelz/channelz.go:30
)

import "google.golang.org/grpc/internal/channelz"

// Identifier is an opaque identifier which uniquely identifies an entity in the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/channelz/channelz.go:34
// channelz database.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/channelz/channelz.go:36
type Identifier = channelz.Identifier

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/channelz/channelz.go:36
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/channelz/channelz.go:36
var _ = _go_fuzz_dep_.CoverTab
