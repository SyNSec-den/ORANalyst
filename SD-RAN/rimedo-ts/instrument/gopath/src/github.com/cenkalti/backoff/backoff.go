//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/backoff.go:1
// Package backoff implements backoff algorithms for retrying operations.
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/backoff.go:1
//
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/backoff.go:1
// Use Retry function for retrying operations that may fail.
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/backoff.go:1
// If Retry does not meet your needs,
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/backoff.go:1
// copy/paste the function into your project and modify as you wish.
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/backoff.go:1
//
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/backoff.go:1
// There is also Ticker type similar to time.Ticker.
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/backoff.go:1
// You can use it if you need to work with channels.
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/backoff.go:1
//
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/backoff.go:1
// See Examples section below for usage examples.
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/backoff.go:11
package backoff

//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/backoff.go:11
import (
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/backoff.go:11
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/backoff.go:11
)
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/backoff.go:11
import (
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/backoff.go:11
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/backoff.go:11
)

import "time"

// BackOff is a backoff policy for retrying an operation.
type BackOff interface {
	// NextBackOff returns the duration to wait before retrying the operation,
	// or backoff. Stop to indicate that no more retries should be made.
	//
	// Example usage:
	//
	// 	duration := backoff.NextBackOff();
	// 	if (duration == backoff.Stop) {
	// 		// Do not retry operation.
	// 	} else {
	// 		// Sleep for duration and retry operation.
	// 	}
	//
	NextBackOff() time.Duration

	// Reset to initial state.
	Reset()
}

// Stop indicates that no more retries should be made for use in NextBackOff().
const Stop time.Duration = -1

// ZeroBackOff is a fixed backoff policy whose backoff time is always zero,
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/backoff.go:38
// meaning that the operation is retried immediately without waiting, indefinitely.
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/backoff.go:40
type ZeroBackOff struct{}

func (b *ZeroBackOff) Reset()	{ _go_fuzz_dep_.CoverTab[182663]++; // _ = "end of CoverTab[182663]" }

func (b *ZeroBackOff) NextBackOff() time.Duration {
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/backoff.go:44
	_go_fuzz_dep_.CoverTab[182664]++
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/backoff.go:44
	return 0
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/backoff.go:44
	// _ = "end of CoverTab[182664]"
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/backoff.go:44
}

// StopBackOff is a fixed backoff policy that always returns backoff.Stop for
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/backoff.go:46
// NextBackOff(), meaning that the operation should never be retried.
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/backoff.go:48
type StopBackOff struct{}

func (b *StopBackOff) Reset()	{ _go_fuzz_dep_.CoverTab[182665]++; // _ = "end of CoverTab[182665]" }

func (b *StopBackOff) NextBackOff() time.Duration {
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/backoff.go:52
	_go_fuzz_dep_.CoverTab[182666]++
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/backoff.go:52
	return Stop
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/backoff.go:52
	// _ = "end of CoverTab[182666]"
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/backoff.go:52
}

// ConstantBackOff is a backoff policy that always returns the same backoff delay.
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/backoff.go:54
// This is in contrast to an exponential backoff policy,
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/backoff.go:54
// which returns a delay that grows longer as you call NextBackOff() over and over again.
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/backoff.go:57
type ConstantBackOff struct {
	Interval time.Duration
}

func (b *ConstantBackOff) Reset()	{ _go_fuzz_dep_.CoverTab[182667]++; // _ = "end of CoverTab[182667]" }
func (b *ConstantBackOff) NextBackOff() time.Duration {
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/backoff.go:62
	_go_fuzz_dep_.CoverTab[182668]++
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/backoff.go:62
	return b.Interval
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/backoff.go:62
	// _ = "end of CoverTab[182668]"
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/backoff.go:62
}

func NewConstantBackOff(d time.Duration) *ConstantBackOff {
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/backoff.go:64
	_go_fuzz_dep_.CoverTab[182669]++
													return &ConstantBackOff{Interval: d}
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/backoff.go:65
	// _ = "end of CoverTab[182669]"
}

//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/backoff.go:66
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/backoff.go:66
var _ = _go_fuzz_dep_.CoverTab
