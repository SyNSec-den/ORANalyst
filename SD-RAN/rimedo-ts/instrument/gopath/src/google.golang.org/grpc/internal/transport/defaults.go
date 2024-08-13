//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/defaults.go:19
package transport

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/defaults.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/defaults.go:19
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/defaults.go:19
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/defaults.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/defaults.go:19
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/defaults.go:19
)

import (
	"math"
	"time"
)

const (
	// The default value of flow control window size in HTTP2 spec.
	defaultWindowSize	= 65535
	// The initial window size for flow control.
	initialWindowSize		= defaultWindowSize	// for an RPC
	infinity			= time.Duration(math.MaxInt64)
	defaultClientKeepaliveTime	= infinity
	defaultClientKeepaliveTimeout	= 20 * time.Second
	defaultMaxStreamsClient		= 100
	defaultMaxConnectionIdle	= infinity
	defaultMaxConnectionAge		= infinity
	defaultMaxConnectionAgeGrace	= infinity
	defaultServerKeepaliveTime	= 2 * time.Hour
	defaultServerKeepaliveTimeout	= 20 * time.Second
	defaultKeepalivePolicyMinTime	= 5 * time.Minute
	// max window limit set by HTTP2 Specs.
	maxWindowSize	= math.MaxInt32
	// defaultWriteQuota is the default value for number of data
	// bytes that each stream can schedule before some of it being
	// flushed out.
	defaultWriteQuota		= 64 * 1024
	defaultClientMaxHeaderListSize	= uint32(16 << 20)
	defaultServerMaxHeaderListSize	= uint32(16 << 20)
)

// MaxStreamID is the upper bound for the stream ID before the current
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/defaults.go:51
// transport gracefully closes and new transport is created for subsequent RPCs.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/defaults.go:51
// This is set to 75% of 2^31-1. Streams are identified with an unsigned 31-bit
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/defaults.go:51
// integer. It's exported so that tests can override it.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/defaults.go:55
var MaxStreamID = uint32(math.MaxInt32 * 3 / 4)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/defaults.go:55
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/defaults.go:55
var _ = _go_fuzz_dep_.CoverTab
