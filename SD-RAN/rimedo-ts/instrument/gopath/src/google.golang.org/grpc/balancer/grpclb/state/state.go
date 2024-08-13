//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/grpclb/state/state.go:19
// Package state declares grpclb types to be set by resolvers wishing to pass
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/grpclb/state/state.go:19
// information to grpclb via resolver.State Attributes.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/grpclb/state/state.go:21
package state

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/grpclb/state/state.go:21
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/grpclb/state/state.go:21
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/grpclb/state/state.go:21
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/grpclb/state/state.go:21
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/grpclb/state/state.go:21
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/grpclb/state/state.go:21
)

import (
	"google.golang.org/grpc/resolver"
)

// keyType is the key to use for storing State in Attributes.
type keyType string

const key = keyType("grpc.grpclb.state")

// State contains gRPCLB-relevant data passed from the name resolver.
type State struct {
	// BalancerAddresses contains the remote load balancer address(es).  If
	// set, overrides any resolver-provided addresses with Type of GRPCLB.
	BalancerAddresses []resolver.Address
}

// Set returns a copy of the provided state with attributes containing s.  s's
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/grpclb/state/state.go:39
// data should not be mutated after calling Set.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/grpclb/state/state.go:41
func Set(state resolver.State, s *State) resolver.State {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/grpclb/state/state.go:41
	_go_fuzz_dep_.CoverTab[68991]++
													state.Attributes = state.Attributes.WithValue(key, s)
													return state
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/grpclb/state/state.go:43
	// _ = "end of CoverTab[68991]"
}

// Get returns the grpclb State in the resolver.State, or nil if not present.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/grpclb/state/state.go:46
// The returned data should not be mutated.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/grpclb/state/state.go:48
func Get(state resolver.State) *State {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/grpclb/state/state.go:48
	_go_fuzz_dep_.CoverTab[68992]++
													s, _ := state.Attributes.Value(key).(*State)
													return s
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/grpclb/state/state.go:50
	// _ = "end of CoverTab[68992]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/grpclb/state/state.go:51
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/grpclb/state/state.go:51
var _ = _go_fuzz_dep_.CoverTab
