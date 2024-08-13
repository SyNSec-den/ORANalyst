//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/backoff.go:22
package grpc

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/backoff.go:22
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/backoff.go:22
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/backoff.go:22
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/backoff.go:22
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/backoff.go:22
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/backoff.go:22
)

import (
	"time"

	"google.golang.org/grpc/backoff"
)

// DefaultBackoffConfig uses values specified for backoff in
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/backoff.go:30
// https://github.com/grpc/grpc/blob/master/doc/connection-backoff.md.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/backoff.go:30
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/backoff.go:30
// Deprecated: use ConnectParams instead. Will be supported throughout 1.x.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/backoff.go:34
var DefaultBackoffConfig = BackoffConfig{
	MaxDelay: 120 * time.Second,
}

// BackoffConfig defines the parameters for the default gRPC backoff strategy.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/backoff.go:38
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/backoff.go:38
// Deprecated: use ConnectParams instead. Will be supported throughout 1.x.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/backoff.go:41
type BackoffConfig struct {
	// MaxDelay is the upper bound of backoff delay.
	MaxDelay time.Duration
}

// ConnectParams defines the parameters for connecting and retrying. Users are
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/backoff.go:46
// encouraged to use this instead of the BackoffConfig type defined above. See
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/backoff.go:46
// here for more details:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/backoff.go:46
// https://github.com/grpc/grpc/blob/master/doc/connection-backoff.md.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/backoff.go:46
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/backoff.go:46
// # Experimental
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/backoff.go:46
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/backoff.go:46
// Notice: This type is EXPERIMENTAL and may be changed or removed in a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/backoff.go:46
// later release.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/backoff.go:55
type ConnectParams struct {
	// Backoff specifies the configuration options for connection backoff.
	Backoff	backoff.Config
	// MinConnectTimeout is the minimum amount of time we are willing to give a
	// connection to complete.
	MinConnectTimeout	time.Duration
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/backoff.go:61
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/backoff.go:61
var _ = _go_fuzz_dep_.CoverTab
