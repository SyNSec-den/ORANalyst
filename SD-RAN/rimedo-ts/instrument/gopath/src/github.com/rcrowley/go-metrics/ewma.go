//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:1
package metrics

//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:1
)

import (
	"math"
	"sync"
	"sync/atomic"
)

// EWMAs continuously calculate an exponentially-weighted moving average
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:9
// based on an outside source of clock ticks.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:11
type EWMA interface {
	Rate() float64
	Snapshot() EWMA
	Tick()
	Update(int64)
}

// NewEWMA constructs a new EWMA with the given alpha.
func NewEWMA(alpha float64) EWMA {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:19
	_go_fuzz_dep_.CoverTab[96110]++
														if UseNilMetrics {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:20
		_go_fuzz_dep_.CoverTab[96112]++
															return NilEWMA{}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:21
		// _ = "end of CoverTab[96112]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:22
		_go_fuzz_dep_.CoverTab[96113]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:22
		// _ = "end of CoverTab[96113]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:22
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:22
	// _ = "end of CoverTab[96110]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:22
	_go_fuzz_dep_.CoverTab[96111]++
														return &StandardEWMA{alpha: alpha}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:23
	// _ = "end of CoverTab[96111]"
}

// NewEWMA1 constructs a new EWMA for a one-minute moving average.
func NewEWMA1() EWMA {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:27
	_go_fuzz_dep_.CoverTab[96114]++
														return NewEWMA(1 - math.Exp(-5.0/60.0/1))
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:28
	// _ = "end of CoverTab[96114]"
}

// NewEWMA5 constructs a new EWMA for a five-minute moving average.
func NewEWMA5() EWMA {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:32
	_go_fuzz_dep_.CoverTab[96115]++
														return NewEWMA(1 - math.Exp(-5.0/60.0/5))
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:33
	// _ = "end of CoverTab[96115]"
}

// NewEWMA15 constructs a new EWMA for a fifteen-minute moving average.
func NewEWMA15() EWMA {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:37
	_go_fuzz_dep_.CoverTab[96116]++
														return NewEWMA(1 - math.Exp(-5.0/60.0/15))
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:38
	// _ = "end of CoverTab[96116]"
}

// EWMASnapshot is a read-only copy of another EWMA.
type EWMASnapshot float64

// Rate returns the rate of events per second at the time the snapshot was
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:44
// taken.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:46
func (a EWMASnapshot) Rate() float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:46
	_go_fuzz_dep_.CoverTab[96117]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:46
	return float64(a)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:46
	// _ = "end of CoverTab[96117]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:46
}

// Snapshot returns the snapshot.
func (a EWMASnapshot) Snapshot() EWMA {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:49
	_go_fuzz_dep_.CoverTab[96118]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:49
	return a
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:49
	// _ = "end of CoverTab[96118]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:49
}

// Tick panics.
func (EWMASnapshot) Tick() {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:52
	_go_fuzz_dep_.CoverTab[96119]++
														panic("Tick called on an EWMASnapshot")
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:53
	// _ = "end of CoverTab[96119]"
}

// Update panics.
func (EWMASnapshot) Update(int64) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:57
	_go_fuzz_dep_.CoverTab[96120]++
														panic("Update called on an EWMASnapshot")
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:58
	// _ = "end of CoverTab[96120]"
}

// NilEWMA is a no-op EWMA.
type NilEWMA struct{}

// Rate is a no-op.
func (NilEWMA) Rate() float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:65
	_go_fuzz_dep_.CoverTab[96121]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:65
	return 0.0
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:65
	// _ = "end of CoverTab[96121]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:65
}

// Snapshot is a no-op.
func (NilEWMA) Snapshot() EWMA {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:68
	_go_fuzz_dep_.CoverTab[96122]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:68
	return NilEWMA{}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:68
	// _ = "end of CoverTab[96122]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:68
}

// Tick is a no-op.
func (NilEWMA) Tick()	{ _go_fuzz_dep_.CoverTab[96123]++; // _ = "end of CoverTab[96123]" }

// Update is a no-op.
func (NilEWMA) Update(n int64)	{ _go_fuzz_dep_.CoverTab[96124]++; // _ = "end of CoverTab[96124]" }

// StandardEWMA is the standard implementation of an EWMA and tracks the number
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:76
// of uncounted events and processes them on each tick.  It uses the
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:76
// sync/atomic package to manage uncounted events.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:79
type StandardEWMA struct {
	uncounted	int64	// /!\ this should be the first member to ensure 64-bit alignment
	alpha		float64
	rate		uint64
	init		uint32
	mutex		sync.Mutex
}

// Rate returns the moving average rate of events per second.
func (a *StandardEWMA) Rate() float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:88
	_go_fuzz_dep_.CoverTab[96125]++
														currentRate := math.Float64frombits(atomic.LoadUint64(&a.rate)) * float64(1e9)
														return currentRate
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:90
	// _ = "end of CoverTab[96125]"
}

// Snapshot returns a read-only copy of the EWMA.
func (a *StandardEWMA) Snapshot() EWMA {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:94
	_go_fuzz_dep_.CoverTab[96126]++
														return EWMASnapshot(a.Rate())
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:95
	// _ = "end of CoverTab[96126]"
}

// Tick ticks the clock to update the moving average.  It assumes it is called
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:98
// every five seconds.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:100
func (a *StandardEWMA) Tick() {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:100
	_go_fuzz_dep_.CoverTab[96127]++

															if atomic.LoadUint32(&a.init) == 1 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:102
		_go_fuzz_dep_.CoverTab[96128]++
																a.updateRate(a.fetchInstantRate())
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:103
		// _ = "end of CoverTab[96128]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:104
		_go_fuzz_dep_.CoverTab[96129]++

//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:109
		a.mutex.Lock()
		if atomic.LoadUint32(&a.init) == 1 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:110
			_go_fuzz_dep_.CoverTab[96131]++

//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:113
			a.updateRate(a.fetchInstantRate())
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:113
			// _ = "end of CoverTab[96131]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:114
			_go_fuzz_dep_.CoverTab[96132]++
																	atomic.StoreUint32(&a.init, 1)
																	atomic.StoreUint64(&a.rate, math.Float64bits(a.fetchInstantRate()))
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:116
			// _ = "end of CoverTab[96132]"
		}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:117
		// _ = "end of CoverTab[96129]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:117
		_go_fuzz_dep_.CoverTab[96130]++
																a.mutex.Unlock()
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:118
		// _ = "end of CoverTab[96130]"
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:119
	// _ = "end of CoverTab[96127]"
}

func (a *StandardEWMA) fetchInstantRate() float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:122
	_go_fuzz_dep_.CoverTab[96133]++
															count := atomic.LoadInt64(&a.uncounted)
															atomic.AddInt64(&a.uncounted, -count)
															instantRate := float64(count) / float64(5e9)
															return instantRate
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:126
	// _ = "end of CoverTab[96133]"
}

func (a *StandardEWMA) updateRate(instantRate float64) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:129
	_go_fuzz_dep_.CoverTab[96134]++
															currentRate := math.Float64frombits(atomic.LoadUint64(&a.rate))
															currentRate += a.alpha * (instantRate - currentRate)
															atomic.StoreUint64(&a.rate, math.Float64bits(currentRate))
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:132
	// _ = "end of CoverTab[96134]"
}

// Update adds n uncounted events.
func (a *StandardEWMA) Update(n int64) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:136
	_go_fuzz_dep_.CoverTab[96135]++
															atomic.AddInt64(&a.uncounted, n)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:137
	// _ = "end of CoverTab[96135]"
}

//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:138
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/ewma.go:138
var _ = _go_fuzz_dep_.CoverTab
