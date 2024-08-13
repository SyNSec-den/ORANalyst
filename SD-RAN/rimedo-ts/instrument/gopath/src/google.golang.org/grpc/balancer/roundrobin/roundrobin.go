//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/roundrobin/roundrobin.go:19
// Package roundrobin defines a roundrobin balancer. Roundrobin balancer is
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/roundrobin/roundrobin.go:19
// installed as one of the default balancers in gRPC, users don't need to
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/roundrobin/roundrobin.go:19
// explicitly install this balancer.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/roundrobin/roundrobin.go:22
package roundrobin

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/roundrobin/roundrobin.go:22
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/roundrobin/roundrobin.go:22
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/roundrobin/roundrobin.go:22
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/roundrobin/roundrobin.go:22
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/roundrobin/roundrobin.go:22
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/roundrobin/roundrobin.go:22
)

import (
	"sync/atomic"

	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/balancer/base"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/internal/grpcrand"
)

// Name is the name of round_robin balancer.
const Name = "round_robin"

var logger = grpclog.Component("roundrobin")

// newBuilder creates a new roundrobin balancer builder.
func newBuilder() balancer.Builder {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/roundrobin/roundrobin.go:39
	_go_fuzz_dep_.CoverTab[67509]++
													return base.NewBalancerBuilder(Name, &rrPickerBuilder{}, base.Config{HealthCheck: true})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/roundrobin/roundrobin.go:40
	// _ = "end of CoverTab[67509]"
}

func init() {
	balancer.Register(newBuilder())
}

type rrPickerBuilder struct{}

func (*rrPickerBuilder) Build(info base.PickerBuildInfo) balancer.Picker {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/roundrobin/roundrobin.go:49
	_go_fuzz_dep_.CoverTab[67510]++
													logger.Infof("roundrobinPicker: Build called with info: %v", info)
													if len(info.ReadySCs) == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/roundrobin/roundrobin.go:51
		_go_fuzz_dep_.CoverTab[67513]++
														return base.NewErrPicker(balancer.ErrNoSubConnAvailable)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/roundrobin/roundrobin.go:52
		// _ = "end of CoverTab[67513]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/roundrobin/roundrobin.go:53
		_go_fuzz_dep_.CoverTab[67514]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/roundrobin/roundrobin.go:53
		// _ = "end of CoverTab[67514]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/roundrobin/roundrobin.go:53
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/roundrobin/roundrobin.go:53
	// _ = "end of CoverTab[67510]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/roundrobin/roundrobin.go:53
	_go_fuzz_dep_.CoverTab[67511]++
													scs := make([]balancer.SubConn, 0, len(info.ReadySCs))
													for sc := range info.ReadySCs {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/roundrobin/roundrobin.go:55
		_go_fuzz_dep_.CoverTab[67515]++
														scs = append(scs, sc)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/roundrobin/roundrobin.go:56
		// _ = "end of CoverTab[67515]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/roundrobin/roundrobin.go:57
	// _ = "end of CoverTab[67511]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/roundrobin/roundrobin.go:57
	_go_fuzz_dep_.CoverTab[67512]++
													return &rrPicker{
														subConns:	scs,

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/roundrobin/roundrobin.go:63
		next:	uint32(grpcrand.Intn(len(scs))),
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/roundrobin/roundrobin.go:64
	// _ = "end of CoverTab[67512]"
}

type rrPicker struct {
	// subConns is the snapshot of the roundrobin balancer when this picker was
	// created. The slice is immutable. Each Get() will do a round robin
	// selection from it and return the selected SubConn.
	subConns	[]balancer.SubConn
	next		uint32
}

func (p *rrPicker) Pick(balancer.PickInfo) (balancer.PickResult, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/roundrobin/roundrobin.go:75
	_go_fuzz_dep_.CoverTab[67516]++
													subConnsLen := uint32(len(p.subConns))
													nextIndex := atomic.AddUint32(&p.next, 1)

													sc := p.subConns[nextIndex%subConnsLen]
													return balancer.PickResult{SubConn: sc}, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/roundrobin/roundrobin.go:80
	// _ = "end of CoverTab[67516]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/roundrobin/roundrobin.go:81
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/roundrobin/roundrobin.go:81
var _ = _go_fuzz_dep_.CoverTab
