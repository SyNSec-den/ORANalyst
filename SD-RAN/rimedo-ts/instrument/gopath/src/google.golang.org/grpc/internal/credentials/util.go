//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/util.go:19
package credentials

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/util.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/util.go:19
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/util.go:19
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/util.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/util.go:19
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/util.go:19
)

import (
	"crypto/tls"
)

const alpnProtoStrH2 = "h2"

// AppendH2ToNextProtos appends h2 to next protos.
func AppendH2ToNextProtos(ps []string) []string {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/util.go:28
	_go_fuzz_dep_.CoverTab[62497]++
													for _, p := range ps {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/util.go:29
		_go_fuzz_dep_.CoverTab[62499]++
														if p == alpnProtoStrH2 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/util.go:30
			_go_fuzz_dep_.CoverTab[62500]++
															return ps
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/util.go:31
			// _ = "end of CoverTab[62500]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/util.go:32
			_go_fuzz_dep_.CoverTab[62501]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/util.go:32
			// _ = "end of CoverTab[62501]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/util.go:32
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/util.go:32
		// _ = "end of CoverTab[62499]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/util.go:33
	// _ = "end of CoverTab[62497]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/util.go:33
	_go_fuzz_dep_.CoverTab[62498]++
													ret := make([]string, 0, len(ps)+1)
													ret = append(ret, ps...)
													return append(ret, alpnProtoStrH2)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/util.go:36
	// _ = "end of CoverTab[62498]"
}

// CloneTLSConfig returns a shallow clone of the exported
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/util.go:39
// fields of cfg, ignoring the unexported sync.Once, which
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/util.go:39
// contains a mutex and must not be copied.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/util.go:39
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/util.go:39
// If cfg is nil, a new zero tls.Config is returned.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/util.go:39
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/util.go:39
// TODO: inline this function if possible.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/util.go:46
func CloneTLSConfig(cfg *tls.Config) *tls.Config {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/util.go:46
	_go_fuzz_dep_.CoverTab[62502]++
													if cfg == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/util.go:47
		_go_fuzz_dep_.CoverTab[62504]++
														return &tls.Config{}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/util.go:48
		// _ = "end of CoverTab[62504]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/util.go:49
		_go_fuzz_dep_.CoverTab[62505]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/util.go:49
		// _ = "end of CoverTab[62505]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/util.go:49
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/util.go:49
	// _ = "end of CoverTab[62502]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/util.go:49
	_go_fuzz_dep_.CoverTab[62503]++

													return cfg.Clone()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/util.go:51
	// _ = "end of CoverTab[62503]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/util.go:52
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/util.go:52
var _ = _go_fuzz_dep_.CoverTab
