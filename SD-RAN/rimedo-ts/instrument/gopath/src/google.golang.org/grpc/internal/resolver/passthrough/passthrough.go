//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/passthrough/passthrough.go:19
// Package passthrough implements a pass-through resolver. It sends the target
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/passthrough/passthrough.go:19
// name without scheme back to gRPC as resolved address.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/passthrough/passthrough.go:21
package passthrough

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/passthrough/passthrough.go:21
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/passthrough/passthrough.go:21
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/passthrough/passthrough.go:21
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/passthrough/passthrough.go:21
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/passthrough/passthrough.go:21
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/passthrough/passthrough.go:21
)

import (
	"errors"

	"google.golang.org/grpc/resolver"
)

const scheme = "passthrough"

type passthroughBuilder struct{}

func (*passthroughBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/passthrough/passthrough.go:33
	_go_fuzz_dep_.CoverTab[69151]++
															if target.Endpoint() == "" && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/passthrough/passthrough.go:34
		_go_fuzz_dep_.CoverTab[69153]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/passthrough/passthrough.go:34
		return opts.Dialer == nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/passthrough/passthrough.go:34
		// _ = "end of CoverTab[69153]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/passthrough/passthrough.go:34
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/passthrough/passthrough.go:34
		_go_fuzz_dep_.CoverTab[69154]++
																return nil, errors.New("passthrough: received empty target in Build()")
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/passthrough/passthrough.go:35
		// _ = "end of CoverTab[69154]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/passthrough/passthrough.go:36
		_go_fuzz_dep_.CoverTab[69155]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/passthrough/passthrough.go:36
		// _ = "end of CoverTab[69155]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/passthrough/passthrough.go:36
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/passthrough/passthrough.go:36
	// _ = "end of CoverTab[69151]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/passthrough/passthrough.go:36
	_go_fuzz_dep_.CoverTab[69152]++
															r := &passthroughResolver{
		target:	target,
		cc:	cc,
	}
															r.start()
															return r, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/passthrough/passthrough.go:42
	// _ = "end of CoverTab[69152]"
}

func (*passthroughBuilder) Scheme() string {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/passthrough/passthrough.go:45
	_go_fuzz_dep_.CoverTab[69156]++
															return scheme
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/passthrough/passthrough.go:46
	// _ = "end of CoverTab[69156]"
}

type passthroughResolver struct {
	target	resolver.Target
	cc	resolver.ClientConn
}

func (r *passthroughResolver) start() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/passthrough/passthrough.go:54
	_go_fuzz_dep_.CoverTab[69157]++
															r.cc.UpdateState(resolver.State{Addresses: []resolver.Address{{Addr: r.target.Endpoint()}}})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/passthrough/passthrough.go:55
	// _ = "end of CoverTab[69157]"
}

func (*passthroughResolver) ResolveNow(o resolver.ResolveNowOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/passthrough/passthrough.go:58
	_go_fuzz_dep_.CoverTab[69158]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/passthrough/passthrough.go:58
	// _ = "end of CoverTab[69158]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/passthrough/passthrough.go:58
}

func (*passthroughResolver) Close()	{ _go_fuzz_dep_.CoverTab[69159]++; // _ = "end of CoverTab[69159]" }

func init() {
	resolver.Register(&passthroughBuilder{})
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/passthrough/passthrough.go:64
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/passthrough/passthrough.go:64
var _ = _go_fuzz_dep_.CoverTab
