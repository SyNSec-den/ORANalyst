//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/backoff/backoff.go:19
// Package backoff provides configuration options for backoff.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/backoff/backoff.go:19
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/backoff/backoff.go:19
// More details can be found at:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/backoff/backoff.go:19
// https://github.com/grpc/grpc/blob/master/doc/connection-backoff.md.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/backoff/backoff.go:19
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/backoff/backoff.go:19
// All APIs in this package are experimental.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/backoff/backoff.go:25
package backoff

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/backoff/backoff.go:25
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/backoff/backoff.go:25
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/backoff/backoff.go:25
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/backoff/backoff.go:25
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/backoff/backoff.go:25
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/backoff/backoff.go:25
)

import "time"

// Config defines the configuration options for backoff.
type Config struct {
	// BaseDelay is the amount of time to backoff after the first failure.
	BaseDelay	time.Duration
	// Multiplier is the factor with which to multiply backoffs after a
	// failed retry. Should ideally be greater than 1.
	Multiplier	float64
	// Jitter is the factor with which backoffs are randomized.
	Jitter	float64
	// MaxDelay is the upper bound of backoff delay.
	MaxDelay	time.Duration
}

// DefaultConfig is a backoff configuration with the default values specfied
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/backoff/backoff.go:42
// at https://github.com/grpc/grpc/blob/master/doc/connection-backoff.md.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/backoff/backoff.go:42
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/backoff/backoff.go:42
// This should be useful for callers who want to configure backoff with
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/backoff/backoff.go:42
// non-default values only for a subset of the options.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/backoff/backoff.go:47
var DefaultConfig = Config{
	BaseDelay:	1.0 * time.Second,
	Multiplier:	1.6,
	Jitter:		0.2,
	MaxDelay:	120 * time.Second,
}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/backoff/backoff.go:52
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/backoff/backoff.go:52
var _ = _go_fuzz_dep_.CoverTab
