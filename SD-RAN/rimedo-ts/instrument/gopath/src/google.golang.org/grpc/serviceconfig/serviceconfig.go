//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/serviceconfig/serviceconfig.go:19
// Package serviceconfig defines types and methods for operating on gRPC
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/serviceconfig/serviceconfig.go:19
// service configs.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/serviceconfig/serviceconfig.go:19
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/serviceconfig/serviceconfig.go:19
// # Experimental
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/serviceconfig/serviceconfig.go:19
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/serviceconfig/serviceconfig.go:19
// Notice: This package is EXPERIMENTAL and may be changed or removed in a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/serviceconfig/serviceconfig.go:19
// later release.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/serviceconfig/serviceconfig.go:26
package serviceconfig

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/serviceconfig/serviceconfig.go:26
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/serviceconfig/serviceconfig.go:26
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/serviceconfig/serviceconfig.go:26
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/serviceconfig/serviceconfig.go:26
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/serviceconfig/serviceconfig.go:26
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/serviceconfig/serviceconfig.go:26
)

// Config represents an opaque data structure holding a service config.
type Config interface {
	isServiceConfig()
}

// LoadBalancingConfig represents an opaque data structure holding a load
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/serviceconfig/serviceconfig.go:33
// balancing config.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/serviceconfig/serviceconfig.go:35
type LoadBalancingConfig interface {
	isLoadBalancingConfig()
}

// ParseResult contains a service config or an error.  Exactly one must be
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/serviceconfig/serviceconfig.go:39
// non-nil.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/serviceconfig/serviceconfig.go:41
type ParseResult struct {
	Config	Config
	Err	error
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/serviceconfig/serviceconfig.go:44
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/serviceconfig/serviceconfig.go:44
var _ = _go_fuzz_dep_.CoverTab
