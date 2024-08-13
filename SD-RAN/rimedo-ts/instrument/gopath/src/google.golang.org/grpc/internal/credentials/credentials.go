//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/credentials.go:17
package credentials

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/credentials.go:17
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/credentials.go:17
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/credentials.go:17
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/credentials.go:17
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/credentials.go:17
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/credentials.go:17
)

import (
	"context"
)

// requestInfoKey is a struct to be used as the key to store RequestInfo in a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/credentials.go:23
// context.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/credentials.go:25
type requestInfoKey struct{}

// NewRequestInfoContext creates a context with ri.
func NewRequestInfoContext(ctx context.Context, ri interface{}) context.Context {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/credentials.go:28
	_go_fuzz_dep_.CoverTab[62457]++
														return context.WithValue(ctx, requestInfoKey{}, ri)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/credentials.go:29
	// _ = "end of CoverTab[62457]"
}

// RequestInfoFromContext extracts the RequestInfo from ctx.
func RequestInfoFromContext(ctx context.Context) interface{} {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/credentials.go:33
	_go_fuzz_dep_.CoverTab[62458]++
														return ctx.Value(requestInfoKey{})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/credentials.go:34
	// _ = "end of CoverTab[62458]"
}

// clientHandshakeInfoKey is a struct used as the key to store
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/credentials.go:37
// ClientHandshakeInfo in a context.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/credentials.go:39
type clientHandshakeInfoKey struct{}

// ClientHandshakeInfoFromContext extracts the ClientHandshakeInfo from ctx.
func ClientHandshakeInfoFromContext(ctx context.Context) interface{} {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/credentials.go:42
	_go_fuzz_dep_.CoverTab[62459]++
														return ctx.Value(clientHandshakeInfoKey{})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/credentials.go:43
	// _ = "end of CoverTab[62459]"
}

// NewClientHandshakeInfoContext creates a context with chi.
func NewClientHandshakeInfoContext(ctx context.Context, chi interface{}) context.Context {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/credentials.go:47
	_go_fuzz_dep_.CoverTab[62460]++
														return context.WithValue(ctx, clientHandshakeInfoKey{}, chi)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/credentials.go:48
	// _ = "end of CoverTab[62460]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/credentials.go:49
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/credentials.go:49
var _ = _go_fuzz_dep_.CoverTab
