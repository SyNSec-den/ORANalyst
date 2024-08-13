//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/base.go:19
// Package base defines a balancer base that can be used to build balancers with
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/base.go:19
// different picking algorithms.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/base.go:19
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/base.go:19
// The base balancer creates a new SubConn for each resolved address. The
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/base.go:19
// provided picker will only be notified about READY SubConns.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/base.go:19
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/base.go:19
// This package is the base of round_robin balancer, its purpose is to be used
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/base.go:19
// to build round_robin like balancers with complex picking algorithms.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/base.go:19
// Balancers with more complicated logic should try to implement a balancer
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/base.go:19
// builder from scratch.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/base.go:19
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/base.go:19
// All APIs in this package are experimental.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/base.go:31
package base

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/base.go:31
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/base.go:31
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/base.go:31
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/base.go:31
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/base.go:31
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/base.go:31
)

import (
	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/resolver"
)

// PickerBuilder creates balancer.Picker.
type PickerBuilder interface {
	// Build returns a picker that will be used by gRPC to pick a SubConn.
	Build(info PickerBuildInfo) balancer.Picker
}

// PickerBuildInfo contains information needed by the picker builder to
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/base.go:44
// construct a picker.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/base.go:46
type PickerBuildInfo struct {
	// ReadySCs is a map from all ready SubConns to the Addresses used to
	// create them.
	ReadySCs map[balancer.SubConn]SubConnInfo
}

// SubConnInfo contains information about a SubConn created by the base
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/base.go:52
// balancer.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/base.go:54
type SubConnInfo struct {
	Address resolver.Address	// the address used to create this SubConn
}

// Config contains the config info about the base balancer builder.
type Config struct {
	// HealthCheck indicates whether health checking should be enabled for this specific balancer.
	HealthCheck bool
}

// NewBalancerBuilder returns a base balancer builder configured by the provided config.
func NewBalancerBuilder(name string, pb PickerBuilder, config Config) balancer.Builder {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/base.go:65
	_go_fuzz_dep_.CoverTab[67502]++
												return &baseBuilder{
		name:		name,
		pickerBuilder:	pb,
		config:		config,
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/base.go:70
	// _ = "end of CoverTab[67502]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/base.go:71
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/base/base.go:71
var _ = _go_fuzz_dep_.CoverTab
