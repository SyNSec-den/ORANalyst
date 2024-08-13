//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/backoff/backoff.go:19
// Package backoff implement the backoff strategy for gRPC.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/backoff/backoff.go:19
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/backoff/backoff.go:19
// This is kept in internal until the gRPC project decides whether or not to
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/backoff/backoff.go:19
// allow alternative backoff strategies.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/backoff/backoff.go:23
package backoff

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/backoff/backoff.go:23
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/backoff/backoff.go:23
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/backoff/backoff.go:23
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/backoff/backoff.go:23
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/backoff/backoff.go:23
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/backoff/backoff.go:23
)

import (
	"time"

	grpcbackoff "google.golang.org/grpc/backoff"
	"google.golang.org/grpc/internal/grpcrand"
)

// Strategy defines the methodology for backing off after a grpc connection
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/backoff/backoff.go:32
// failure.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/backoff/backoff.go:34
type Strategy interface {
	// Backoff returns the amount of time to wait before the next retry given
	// the number of consecutive failures.
	Backoff(retries int) time.Duration
}

// DefaultExponential is an exponential backoff implementation using the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/backoff/backoff.go:40
// default values for all the configurable knobs defined in
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/backoff/backoff.go:40
// https://github.com/grpc/grpc/blob/master/doc/connection-backoff.md.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/backoff/backoff.go:43
var DefaultExponential = Exponential{Config: grpcbackoff.DefaultConfig}

// Exponential implements exponential backoff algorithm as defined in
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/backoff/backoff.go:45
// https://github.com/grpc/grpc/blob/master/doc/connection-backoff.md.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/backoff/backoff.go:47
type Exponential struct {
	// Config contains all options to configure the backoff algorithm.
	Config grpcbackoff.Config
}

// Backoff returns the amount of time to wait before the next retry given the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/backoff/backoff.go:52
// number of retries.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/backoff/backoff.go:54
func (bc Exponential) Backoff(retries int) time.Duration {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/backoff/backoff.go:54
	_go_fuzz_dep_.CoverTab[67675]++
													if retries == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/backoff/backoff.go:55
		_go_fuzz_dep_.CoverTab[67680]++
														return bc.Config.BaseDelay
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/backoff/backoff.go:56
		// _ = "end of CoverTab[67680]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/backoff/backoff.go:57
		_go_fuzz_dep_.CoverTab[67681]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/backoff/backoff.go:57
		// _ = "end of CoverTab[67681]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/backoff/backoff.go:57
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/backoff/backoff.go:57
	// _ = "end of CoverTab[67675]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/backoff/backoff.go:57
	_go_fuzz_dep_.CoverTab[67676]++
													backoff, max := float64(bc.Config.BaseDelay), float64(bc.Config.MaxDelay)
													for backoff < max && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/backoff/backoff.go:59
		_go_fuzz_dep_.CoverTab[67682]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/backoff/backoff.go:59
		return retries > 0
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/backoff/backoff.go:59
		// _ = "end of CoverTab[67682]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/backoff/backoff.go:59
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/backoff/backoff.go:59
		_go_fuzz_dep_.CoverTab[67683]++
														backoff *= bc.Config.Multiplier
														retries--
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/backoff/backoff.go:61
		// _ = "end of CoverTab[67683]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/backoff/backoff.go:62
	// _ = "end of CoverTab[67676]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/backoff/backoff.go:62
	_go_fuzz_dep_.CoverTab[67677]++
													if backoff > max {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/backoff/backoff.go:63
		_go_fuzz_dep_.CoverTab[67684]++
														backoff = max
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/backoff/backoff.go:64
		// _ = "end of CoverTab[67684]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/backoff/backoff.go:65
		_go_fuzz_dep_.CoverTab[67685]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/backoff/backoff.go:65
		// _ = "end of CoverTab[67685]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/backoff/backoff.go:65
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/backoff/backoff.go:65
	// _ = "end of CoverTab[67677]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/backoff/backoff.go:65
	_go_fuzz_dep_.CoverTab[67678]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/backoff/backoff.go:68
	backoff *= 1 + bc.Config.Jitter*(grpcrand.Float64()*2-1)
	if backoff < 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/backoff/backoff.go:69
		_go_fuzz_dep_.CoverTab[67686]++
														return 0
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/backoff/backoff.go:70
		// _ = "end of CoverTab[67686]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/backoff/backoff.go:71
		_go_fuzz_dep_.CoverTab[67687]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/backoff/backoff.go:71
		// _ = "end of CoverTab[67687]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/backoff/backoff.go:71
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/backoff/backoff.go:71
	// _ = "end of CoverTab[67678]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/backoff/backoff.go:71
	_go_fuzz_dep_.CoverTab[67679]++
													return time.Duration(backoff)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/backoff/backoff.go:72
	// _ = "end of CoverTab[67679]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/backoff/backoff.go:73
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/backoff/backoff.go:73
var _ = _go_fuzz_dep_.CoverTab
