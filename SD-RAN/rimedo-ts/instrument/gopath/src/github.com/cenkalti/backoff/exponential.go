//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/exponential.go:1
package backoff

//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/exponential.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/exponential.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/exponential.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/exponential.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/exponential.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/exponential.go:1
)

import (
	"math/rand"
	"time"
)

/*
ExponentialBackOff is a backoff implementation that increases the backoff
period for each retry attempt using a randomization function that grows exponentially.

NextBackOff() is calculated using the following formula:

	randomized interval =
	    RetryInterval * (random value in range [1 - RandomizationFactor, 1 + RandomizationFactor])

In other words NextBackOff() will range between the randomization factor
percentage below and above the retry interval.

For example, given the following parameters:

	RetryInterval = 2
	RandomizationFactor = 0.5
	Multiplier = 2

the actual backoff period used in the next retry attempt will range between 1 and 3 seconds,
multiplied by the exponential, that is, between 2 and 6 seconds.

Note: MaxInterval caps the RetryInterval and not the randomized interval.

If the time elapsed since an ExponentialBackOff instance is created goes past the
MaxElapsedTime, then the method NextBackOff() starts returning backoff.Stop.

The elapsed time can be reset by calling Reset().

Example: Given the following default arguments, for 10 tries the sequence will be,
and assuming we go over the MaxElapsedTime on the 10th try:

	Request #  RetryInterval (seconds)  Randomized Interval (seconds)

	 1          0.5                     [0.25,   0.75]
	 2          0.75                    [0.375,  1.125]
	 3          1.125                   [0.562,  1.687]
	 4          1.687                   [0.8435, 2.53]
	 5          2.53                    [1.265,  3.795]
	 6          3.795                   [1.897,  5.692]
	 7          5.692                   [2.846,  8.538]
	 8          8.538                   [4.269, 12.807]
	 9         12.807                   [6.403, 19.210]
	10         19.210                   backoff.Stop

Note: Implementation is not thread-safe.
*/
type ExponentialBackOff struct {
	InitialInterval		time.Duration
	RandomizationFactor	float64
	Multiplier		float64
	MaxInterval		time.Duration
	// After MaxElapsedTime the ExponentialBackOff stops.
	// It never stops if MaxElapsedTime == 0.
	MaxElapsedTime	time.Duration
	Clock		Clock

	currentInterval	time.Duration
	startTime	time.Time
}

// Clock is an interface that returns current time for BackOff.
type Clock interface {
	Now() time.Time
}

// Default values for ExponentialBackOff.
const (
	DefaultInitialInterval		= 500 * time.Millisecond
	DefaultRandomizationFactor	= 0.5
	DefaultMultiplier		= 1.5
	DefaultMaxInterval		= 60 * time.Second
	DefaultMaxElapsedTime		= 15 * time.Minute
)

// NewExponentialBackOff creates an instance of ExponentialBackOff using default values.
func NewExponentialBackOff() *ExponentialBackOff {
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/exponential.go:83
	_go_fuzz_dep_.CoverTab[182690]++
													b := &ExponentialBackOff{
		InitialInterval:	DefaultInitialInterval,
		RandomizationFactor:	DefaultRandomizationFactor,
		Multiplier:		DefaultMultiplier,
		MaxInterval:		DefaultMaxInterval,
		MaxElapsedTime:		DefaultMaxElapsedTime,
		Clock:			SystemClock,
	}
													b.Reset()
													return b
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/exponential.go:93
	// _ = "end of CoverTab[182690]"
}

type systemClock struct{}

func (t systemClock) Now() time.Time {
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/exponential.go:98
	_go_fuzz_dep_.CoverTab[182691]++
													return time.Now()
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/exponential.go:99
	// _ = "end of CoverTab[182691]"
}

// SystemClock implements Clock interface that uses time.Now().
var SystemClock = systemClock{}

// Reset the interval back to the initial retry interval and restarts the timer.
func (b *ExponentialBackOff) Reset() {
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/exponential.go:106
	_go_fuzz_dep_.CoverTab[182692]++
													b.currentInterval = b.InitialInterval
													b.startTime = b.Clock.Now()
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/exponential.go:108
	// _ = "end of CoverTab[182692]"
}

// NextBackOff calculates the next backoff interval using the formula:
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/exponential.go:111
//
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/exponential.go:111
//	Randomized interval = RetryInterval +/- (RandomizationFactor * RetryInterval)
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/exponential.go:113
func (b *ExponentialBackOff) NextBackOff() time.Duration {
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/exponential.go:113
	_go_fuzz_dep_.CoverTab[182693]++

													if b.MaxElapsedTime != 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/exponential.go:115
		_go_fuzz_dep_.CoverTab[182695]++
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/exponential.go:115
		return b.GetElapsedTime() > b.MaxElapsedTime
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/exponential.go:115
		// _ = "end of CoverTab[182695]"
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/exponential.go:115
	}() {
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/exponential.go:115
		_go_fuzz_dep_.CoverTab[182696]++
														return Stop
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/exponential.go:116
		// _ = "end of CoverTab[182696]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/exponential.go:117
		_go_fuzz_dep_.CoverTab[182697]++
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/exponential.go:117
		// _ = "end of CoverTab[182697]"
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/exponential.go:117
	}
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/exponential.go:117
	// _ = "end of CoverTab[182693]"
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/exponential.go:117
	_go_fuzz_dep_.CoverTab[182694]++
													defer b.incrementCurrentInterval()
													return getRandomValueFromInterval(b.RandomizationFactor, rand.Float64(), b.currentInterval)
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/exponential.go:119
	// _ = "end of CoverTab[182694]"
}

// GetElapsedTime returns the elapsed time since an ExponentialBackOff instance
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/exponential.go:122
// is created and is reset when Reset() is called.
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/exponential.go:122
//
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/exponential.go:122
// The elapsed time is computed using time.Now().UnixNano(). It is
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/exponential.go:122
// safe to call even while the backoff policy is used by a running
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/exponential.go:122
// ticker.
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/exponential.go:128
func (b *ExponentialBackOff) GetElapsedTime() time.Duration {
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/exponential.go:128
	_go_fuzz_dep_.CoverTab[182698]++
													return b.Clock.Now().Sub(b.startTime)
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/exponential.go:129
	// _ = "end of CoverTab[182698]"
}

// Increments the current interval by multiplying it with the multiplier.
func (b *ExponentialBackOff) incrementCurrentInterval() {
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/exponential.go:133
	_go_fuzz_dep_.CoverTab[182699]++

													if float64(b.currentInterval) >= float64(b.MaxInterval)/b.Multiplier {
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/exponential.go:135
		_go_fuzz_dep_.CoverTab[182700]++
														b.currentInterval = b.MaxInterval
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/exponential.go:136
		// _ = "end of CoverTab[182700]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/exponential.go:137
		_go_fuzz_dep_.CoverTab[182701]++
														b.currentInterval = time.Duration(float64(b.currentInterval) * b.Multiplier)
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/exponential.go:138
		// _ = "end of CoverTab[182701]"
	}
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/exponential.go:139
	// _ = "end of CoverTab[182699]"
}

// Returns a random value from the following interval:
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/exponential.go:142
//
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/exponential.go:142
//	[randomizationFactor * currentInterval, randomizationFactor * currentInterval].
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/exponential.go:144
func getRandomValueFromInterval(randomizationFactor, random float64, currentInterval time.Duration) time.Duration {
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/exponential.go:144
	_go_fuzz_dep_.CoverTab[182702]++
													var delta = randomizationFactor * float64(currentInterval)
													var minInterval = float64(currentInterval) - delta
													var maxInterval = float64(currentInterval) + delta

//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/exponential.go:152
	return time.Duration(minInterval + (random * (maxInterval - minInterval + 1)))
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/exponential.go:152
	// _ = "end of CoverTab[182702]"
}

//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/exponential.go:153
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/exponential.go:153
var _ = _go_fuzz_dep_.CoverTab
