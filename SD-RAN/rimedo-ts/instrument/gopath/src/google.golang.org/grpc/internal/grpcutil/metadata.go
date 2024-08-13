//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/metadata.go:19
package grpcutil

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/metadata.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/metadata.go:19
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/metadata.go:19
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/metadata.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/metadata.go:19
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/metadata.go:19
)

import (
	"context"

	"google.golang.org/grpc/metadata"
)

type mdExtraKey struct{}

// WithExtraMetadata creates a new context with incoming md attached.
func WithExtraMetadata(ctx context.Context, md metadata.MD) context.Context {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/metadata.go:30
	_go_fuzz_dep_.CoverTab[67628]++
													return context.WithValue(ctx, mdExtraKey{}, md)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/metadata.go:31
	// _ = "end of CoverTab[67628]"
}

// ExtraMetadata returns the incoming metadata in ctx if it exists.  The
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/metadata.go:34
// returned MD should not be modified. Writing to it may cause races.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/metadata.go:34
// Modification should be made to copies of the returned MD.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/metadata.go:37
func ExtraMetadata(ctx context.Context) (md metadata.MD, ok bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/metadata.go:37
	_go_fuzz_dep_.CoverTab[67629]++
													md, ok = ctx.Value(mdExtraKey{}).(metadata.MD)
													return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/metadata.go:39
	// _ = "end of CoverTab[67629]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/metadata.go:40
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/metadata.go:40
var _ = _go_fuzz_dep_.CoverTab
