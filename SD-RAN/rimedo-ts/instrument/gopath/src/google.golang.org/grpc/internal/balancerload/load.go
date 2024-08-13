//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancerload/load.go:17
// Package balancerload defines APIs to parse server loads in trailers. The
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancerload/load.go:17
// parsed loads are sent to balancers in DoneInfo.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancerload/load.go:19
package balancerload

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancerload/load.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancerload/load.go:19
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancerload/load.go:19
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancerload/load.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancerload/load.go:19
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancerload/load.go:19
)

import (
	"google.golang.org/grpc/metadata"
)

// Parser converts loads from metadata into a concrete type.
type Parser interface {
	// Parse parses loads from metadata.
	Parse(md metadata.MD) interface{}
}

var parser Parser

// SetParser sets the load parser.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancerload/load.go:33
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancerload/load.go:33
// Not mutex-protected, should be called before any gRPC functions.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancerload/load.go:36
func SetParser(lr Parser) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancerload/load.go:36
	_go_fuzz_dep_.CoverTab[67786]++
													parser = lr
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancerload/load.go:37
	// _ = "end of CoverTab[67786]"
}

// Parse calls parser.Read().
func Parse(md metadata.MD) interface{} {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancerload/load.go:41
	_go_fuzz_dep_.CoverTab[67787]++
													if parser == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancerload/load.go:42
		_go_fuzz_dep_.CoverTab[67789]++
														return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancerload/load.go:43
		// _ = "end of CoverTab[67789]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancerload/load.go:44
		_go_fuzz_dep_.CoverTab[67790]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancerload/load.go:44
		// _ = "end of CoverTab[67790]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancerload/load.go:44
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancerload/load.go:44
	// _ = "end of CoverTab[67787]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancerload/load.go:44
	_go_fuzz_dep_.CoverTab[67788]++
													return parser.Parse(md)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancerload/load.go:45
	// _ = "end of CoverTab[67788]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancerload/load.go:46
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/balancerload/load.go:46
var _ = _go_fuzz_dep_.CoverTab
