//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/unix/unix.go:19
// Package unix implements a resolver for unix targets.
package unix

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/unix/unix.go:20
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/unix/unix.go:20
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/unix/unix.go:20
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/unix/unix.go:20
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/unix/unix.go:20
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/unix/unix.go:20
)

import (
	"fmt"

	"google.golang.org/grpc/internal/transport/networktype"
	"google.golang.org/grpc/resolver"
)

const unixScheme = "unix"
const unixAbstractScheme = "unix-abstract"

type builder struct {
	scheme string
}

func (b *builder) Build(target resolver.Target, cc resolver.ClientConn, _ resolver.BuildOptions) (resolver.Resolver, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/unix/unix.go:36
	_go_fuzz_dep_.CoverTab[69165]++
													if target.URL.Host != "" {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/unix/unix.go:37
		_go_fuzz_dep_.CoverTab[69169]++
														return nil, fmt.Errorf("invalid (non-empty) authority: %v", target.URL.Host)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/unix/unix.go:38
		// _ = "end of CoverTab[69169]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/unix/unix.go:39
		_go_fuzz_dep_.CoverTab[69170]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/unix/unix.go:39
		// _ = "end of CoverTab[69170]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/unix/unix.go:39
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/unix/unix.go:39
	// _ = "end of CoverTab[69165]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/unix/unix.go:39
	_go_fuzz_dep_.CoverTab[69166]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/unix/unix.go:46
	endpoint := target.URL.Path
	if endpoint == "" {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/unix/unix.go:47
		_go_fuzz_dep_.CoverTab[69171]++
														endpoint = target.URL.Opaque
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/unix/unix.go:48
		// _ = "end of CoverTab[69171]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/unix/unix.go:49
		_go_fuzz_dep_.CoverTab[69172]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/unix/unix.go:49
		// _ = "end of CoverTab[69172]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/unix/unix.go:49
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/unix/unix.go:49
	// _ = "end of CoverTab[69166]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/unix/unix.go:49
	_go_fuzz_dep_.CoverTab[69167]++
													addr := resolver.Address{Addr: endpoint}
													if b.scheme == unixAbstractScheme {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/unix/unix.go:51
		_go_fuzz_dep_.CoverTab[69173]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/unix/unix.go:54
		addr.Addr = "@" + addr.Addr
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/unix/unix.go:54
		// _ = "end of CoverTab[69173]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/unix/unix.go:55
		_go_fuzz_dep_.CoverTab[69174]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/unix/unix.go:55
		// _ = "end of CoverTab[69174]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/unix/unix.go:55
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/unix/unix.go:55
	// _ = "end of CoverTab[69167]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/unix/unix.go:55
	_go_fuzz_dep_.CoverTab[69168]++
													cc.UpdateState(resolver.State{Addresses: []resolver.Address{networktype.Set(addr, "unix")}})
													return &nopResolver{}, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/unix/unix.go:57
	// _ = "end of CoverTab[69168]"
}

func (b *builder) Scheme() string {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/unix/unix.go:60
	_go_fuzz_dep_.CoverTab[69175]++
													return b.scheme
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/unix/unix.go:61
	// _ = "end of CoverTab[69175]"
}

type nopResolver struct {
}

func (*nopResolver) ResolveNow(resolver.ResolveNowOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/unix/unix.go:67
	_go_fuzz_dep_.CoverTab[69176]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/unix/unix.go:67
	// _ = "end of CoverTab[69176]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/unix/unix.go:67
}

func (*nopResolver) Close()	{ _go_fuzz_dep_.CoverTab[69177]++; // _ = "end of CoverTab[69177]" }

func init() {
	resolver.Register(&builder{scheme: unixScheme})
	resolver.Register(&builder{scheme: unixAbstractScheme})
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/unix/unix.go:74
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/unix/unix.go:74
var _ = _go_fuzz_dep_.CoverTab
