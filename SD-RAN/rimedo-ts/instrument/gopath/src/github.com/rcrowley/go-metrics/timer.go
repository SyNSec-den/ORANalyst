//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:1
package metrics

//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:1
)

import (
	"sync"
	"time"
)

// Timers capture the duration and rate of events.
type Timer interface {
	Count() int64
	Max() int64
	Mean() float64
	Min() int64
	Percentile(float64) float64
	Percentiles([]float64) []float64
	Rate1() float64
	Rate5() float64
	Rate15() float64
	RateMean() float64
	Snapshot() Timer
	StdDev() float64
	Stop()
	Sum() int64
	Time(func())
	Update(time.Duration)
	UpdateSince(time.Time)
	Variance() float64
}

// GetOrRegisterTimer returns an existing Timer or constructs and registers a
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:30
// new StandardTimer.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:30
// Be sure to unregister the meter from the registry once it is of no use to
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:30
// allow for garbage collection.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:34
func GetOrRegisterTimer(name string, r Registry) Timer {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:34
	_go_fuzz_dep_.CoverTab[96643]++
															if nil == r {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:35
		_go_fuzz_dep_.CoverTab[96645]++
																r = DefaultRegistry
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:36
		// _ = "end of CoverTab[96645]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:37
		_go_fuzz_dep_.CoverTab[96646]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:37
		// _ = "end of CoverTab[96646]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:37
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:37
	// _ = "end of CoverTab[96643]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:37
	_go_fuzz_dep_.CoverTab[96644]++
															return r.GetOrRegister(name, NewTimer).(Timer)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:38
	// _ = "end of CoverTab[96644]"
}

// NewCustomTimer constructs a new StandardTimer from a Histogram and a Meter.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:41
// Be sure to call Stop() once the timer is of no use to allow for garbage collection.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:43
func NewCustomTimer(h Histogram, m Meter) Timer {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:43
	_go_fuzz_dep_.CoverTab[96647]++
															if UseNilMetrics {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:44
		_go_fuzz_dep_.CoverTab[96649]++
																return NilTimer{}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:45
		// _ = "end of CoverTab[96649]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:46
		_go_fuzz_dep_.CoverTab[96650]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:46
		// _ = "end of CoverTab[96650]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:46
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:46
	// _ = "end of CoverTab[96647]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:46
	_go_fuzz_dep_.CoverTab[96648]++
															return &StandardTimer{
		histogram:	h,
		meter:		m,
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:50
	// _ = "end of CoverTab[96648]"
}

// NewRegisteredTimer constructs and registers a new StandardTimer.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:53
// Be sure to unregister the meter from the registry once it is of no use to
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:53
// allow for garbage collection.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:56
func NewRegisteredTimer(name string, r Registry) Timer {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:56
	_go_fuzz_dep_.CoverTab[96651]++
															c := NewTimer()
															if nil == r {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:58
		_go_fuzz_dep_.CoverTab[96653]++
																r = DefaultRegistry
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:59
		// _ = "end of CoverTab[96653]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:60
		_go_fuzz_dep_.CoverTab[96654]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:60
		// _ = "end of CoverTab[96654]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:60
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:60
	// _ = "end of CoverTab[96651]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:60
	_go_fuzz_dep_.CoverTab[96652]++
															r.Register(name, c)
															return c
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:62
	// _ = "end of CoverTab[96652]"
}

// NewTimer constructs a new StandardTimer using an exponentially-decaying
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:65
// sample with the same reservoir size and alpha as UNIX load averages.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:65
// Be sure to call Stop() once the timer is of no use to allow for garbage collection.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:68
func NewTimer() Timer {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:68
	_go_fuzz_dep_.CoverTab[96655]++
															if UseNilMetrics {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:69
		_go_fuzz_dep_.CoverTab[96657]++
																return NilTimer{}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:70
		// _ = "end of CoverTab[96657]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:71
		_go_fuzz_dep_.CoverTab[96658]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:71
		// _ = "end of CoverTab[96658]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:71
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:71
	// _ = "end of CoverTab[96655]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:71
	_go_fuzz_dep_.CoverTab[96656]++
															return &StandardTimer{
		histogram:	NewHistogram(NewExpDecaySample(1028, 0.015)),
		meter:		NewMeter(),
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:75
	// _ = "end of CoverTab[96656]"
}

// NilTimer is a no-op Timer.
type NilTimer struct {
	h	Histogram
	m	Meter
}

// Count is a no-op.
func (NilTimer) Count() int64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:85
	_go_fuzz_dep_.CoverTab[96659]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:85
	return 0
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:85
	// _ = "end of CoverTab[96659]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:85
}

// Max is a no-op.
func (NilTimer) Max() int64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:88
	_go_fuzz_dep_.CoverTab[96660]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:88
	return 0
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:88
	// _ = "end of CoverTab[96660]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:88
}

// Mean is a no-op.
func (NilTimer) Mean() float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:91
	_go_fuzz_dep_.CoverTab[96661]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:91
	return 0.0
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:91
	// _ = "end of CoverTab[96661]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:91
}

// Min is a no-op.
func (NilTimer) Min() int64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:94
	_go_fuzz_dep_.CoverTab[96662]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:94
	return 0
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:94
	// _ = "end of CoverTab[96662]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:94
}

// Percentile is a no-op.
func (NilTimer) Percentile(p float64) float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:97
	_go_fuzz_dep_.CoverTab[96663]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:97
	return 0.0
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:97
	// _ = "end of CoverTab[96663]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:97
}

// Percentiles is a no-op.
func (NilTimer) Percentiles(ps []float64) []float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:100
	_go_fuzz_dep_.CoverTab[96664]++
															return make([]float64, len(ps))
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:101
	// _ = "end of CoverTab[96664]"
}

// Rate1 is a no-op.
func (NilTimer) Rate1() float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:105
	_go_fuzz_dep_.CoverTab[96665]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:105
	return 0.0
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:105
	// _ = "end of CoverTab[96665]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:105
}

// Rate5 is a no-op.
func (NilTimer) Rate5() float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:108
	_go_fuzz_dep_.CoverTab[96666]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:108
	return 0.0
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:108
	// _ = "end of CoverTab[96666]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:108
}

// Rate15 is a no-op.
func (NilTimer) Rate15() float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:111
	_go_fuzz_dep_.CoverTab[96667]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:111
	return 0.0
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:111
	// _ = "end of CoverTab[96667]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:111
}

// RateMean is a no-op.
func (NilTimer) RateMean() float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:114
	_go_fuzz_dep_.CoverTab[96668]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:114
	return 0.0
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:114
	// _ = "end of CoverTab[96668]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:114
}

// Snapshot is a no-op.
func (NilTimer) Snapshot() Timer {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:117
	_go_fuzz_dep_.CoverTab[96669]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:117
	return NilTimer{}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:117
	// _ = "end of CoverTab[96669]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:117
}

// StdDev is a no-op.
func (NilTimer) StdDev() float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:120
	_go_fuzz_dep_.CoverTab[96670]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:120
	return 0.0
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:120
	// _ = "end of CoverTab[96670]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:120
}

// Stop is a no-op.
func (NilTimer) Stop()	{ _go_fuzz_dep_.CoverTab[96671]++; // _ = "end of CoverTab[96671]" }

// Sum is a no-op.
func (NilTimer) Sum() int64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:126
	_go_fuzz_dep_.CoverTab[96672]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:126
	return 0
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:126
	// _ = "end of CoverTab[96672]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:126
}

// Time is a no-op.
func (NilTimer) Time(func())	{ _go_fuzz_dep_.CoverTab[96673]++; // _ = "end of CoverTab[96673]" }

// Update is a no-op.
func (NilTimer) Update(time.Duration) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:132
	_go_fuzz_dep_.CoverTab[96674]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:132
	// _ = "end of CoverTab[96674]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:132
}

// UpdateSince is a no-op.
func (NilTimer) UpdateSince(time.Time) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:135
	_go_fuzz_dep_.CoverTab[96675]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:135
	// _ = "end of CoverTab[96675]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:135
}

// Variance is a no-op.
func (NilTimer) Variance() float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:138
	_go_fuzz_dep_.CoverTab[96676]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:138
	return 0.0
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:138
	// _ = "end of CoverTab[96676]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:138
}

// StandardTimer is the standard implementation of a Timer and uses a Histogram
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:140
// and Meter.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:142
type StandardTimer struct {
	histogram	Histogram
	meter		Meter
	mutex		sync.Mutex
}

// Count returns the number of events recorded.
func (t *StandardTimer) Count() int64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:149
	_go_fuzz_dep_.CoverTab[96677]++
															return t.histogram.Count()
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:150
	// _ = "end of CoverTab[96677]"
}

// Max returns the maximum value in the sample.
func (t *StandardTimer) Max() int64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:154
	_go_fuzz_dep_.CoverTab[96678]++
															return t.histogram.Max()
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:155
	// _ = "end of CoverTab[96678]"
}

// Mean returns the mean of the values in the sample.
func (t *StandardTimer) Mean() float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:159
	_go_fuzz_dep_.CoverTab[96679]++
															return t.histogram.Mean()
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:160
	// _ = "end of CoverTab[96679]"
}

// Min returns the minimum value in the sample.
func (t *StandardTimer) Min() int64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:164
	_go_fuzz_dep_.CoverTab[96680]++
															return t.histogram.Min()
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:165
	// _ = "end of CoverTab[96680]"
}

// Percentile returns an arbitrary percentile of the values in the sample.
func (t *StandardTimer) Percentile(p float64) float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:169
	_go_fuzz_dep_.CoverTab[96681]++
															return t.histogram.Percentile(p)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:170
	// _ = "end of CoverTab[96681]"
}

// Percentiles returns a slice of arbitrary percentiles of the values in the
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:173
// sample.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:175
func (t *StandardTimer) Percentiles(ps []float64) []float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:175
	_go_fuzz_dep_.CoverTab[96682]++
															return t.histogram.Percentiles(ps)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:176
	// _ = "end of CoverTab[96682]"
}

// Rate1 returns the one-minute moving average rate of events per second.
func (t *StandardTimer) Rate1() float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:180
	_go_fuzz_dep_.CoverTab[96683]++
															return t.meter.Rate1()
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:181
	// _ = "end of CoverTab[96683]"
}

// Rate5 returns the five-minute moving average rate of events per second.
func (t *StandardTimer) Rate5() float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:185
	_go_fuzz_dep_.CoverTab[96684]++
															return t.meter.Rate5()
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:186
	// _ = "end of CoverTab[96684]"
}

// Rate15 returns the fifteen-minute moving average rate of events per second.
func (t *StandardTimer) Rate15() float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:190
	_go_fuzz_dep_.CoverTab[96685]++
															return t.meter.Rate15()
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:191
	// _ = "end of CoverTab[96685]"
}

// RateMean returns the meter's mean rate of events per second.
func (t *StandardTimer) RateMean() float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:195
	_go_fuzz_dep_.CoverTab[96686]++
															return t.meter.RateMean()
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:196
	// _ = "end of CoverTab[96686]"
}

// Snapshot returns a read-only copy of the timer.
func (t *StandardTimer) Snapshot() Timer {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:200
	_go_fuzz_dep_.CoverTab[96687]++
															t.mutex.Lock()
															defer t.mutex.Unlock()
															return &TimerSnapshot{
		histogram:	t.histogram.Snapshot().(*HistogramSnapshot),
		meter:		t.meter.Snapshot().(*MeterSnapshot),
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:206
	// _ = "end of CoverTab[96687]"
}

// StdDev returns the standard deviation of the values in the sample.
func (t *StandardTimer) StdDev() float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:210
	_go_fuzz_dep_.CoverTab[96688]++
															return t.histogram.StdDev()
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:211
	// _ = "end of CoverTab[96688]"
}

// Stop stops the meter.
func (t *StandardTimer) Stop() {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:215
	_go_fuzz_dep_.CoverTab[96689]++
															t.meter.Stop()
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:216
	// _ = "end of CoverTab[96689]"
}

// Sum returns the sum in the sample.
func (t *StandardTimer) Sum() int64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:220
	_go_fuzz_dep_.CoverTab[96690]++
															return t.histogram.Sum()
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:221
	// _ = "end of CoverTab[96690]"
}

// Record the duration of the execution of the given function.
func (t *StandardTimer) Time(f func()) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:225
	_go_fuzz_dep_.CoverTab[96691]++
															ts := time.Now()
															f()
															t.Update(time.Since(ts))
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:228
	// _ = "end of CoverTab[96691]"
}

// Record the duration of an event.
func (t *StandardTimer) Update(d time.Duration) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:232
	_go_fuzz_dep_.CoverTab[96692]++
															t.mutex.Lock()
															defer t.mutex.Unlock()
															t.histogram.Update(int64(d))
															t.meter.Mark(1)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:236
	// _ = "end of CoverTab[96692]"
}

// Record the duration of an event that started at a time and ends now.
func (t *StandardTimer) UpdateSince(ts time.Time) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:240
	_go_fuzz_dep_.CoverTab[96693]++
															t.mutex.Lock()
															defer t.mutex.Unlock()
															t.histogram.Update(int64(time.Since(ts)))
															t.meter.Mark(1)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:244
	// _ = "end of CoverTab[96693]"
}

// Variance returns the variance of the values in the sample.
func (t *StandardTimer) Variance() float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:248
	_go_fuzz_dep_.CoverTab[96694]++
															return t.histogram.Variance()
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:249
	// _ = "end of CoverTab[96694]"
}

// TimerSnapshot is a read-only copy of another Timer.
type TimerSnapshot struct {
	histogram	*HistogramSnapshot
	meter		*MeterSnapshot
}

// Count returns the number of events recorded at the time the snapshot was
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:258
// taken.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:260
func (t *TimerSnapshot) Count() int64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:260
	_go_fuzz_dep_.CoverTab[96695]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:260
	return t.histogram.Count()
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:260
	// _ = "end of CoverTab[96695]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:260
}

// Max returns the maximum value at the time the snapshot was taken.
func (t *TimerSnapshot) Max() int64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:263
	_go_fuzz_dep_.CoverTab[96696]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:263
	return t.histogram.Max()
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:263
	// _ = "end of CoverTab[96696]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:263
}

// Mean returns the mean value at the time the snapshot was taken.
func (t *TimerSnapshot) Mean() float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:266
	_go_fuzz_dep_.CoverTab[96697]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:266
	return t.histogram.Mean()
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:266
	// _ = "end of CoverTab[96697]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:266
}

// Min returns the minimum value at the time the snapshot was taken.
func (t *TimerSnapshot) Min() int64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:269
	_go_fuzz_dep_.CoverTab[96698]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:269
	return t.histogram.Min()
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:269
	// _ = "end of CoverTab[96698]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:269
}

// Percentile returns an arbitrary percentile of sampled values at the time the
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:271
// snapshot was taken.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:273
func (t *TimerSnapshot) Percentile(p float64) float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:273
	_go_fuzz_dep_.CoverTab[96699]++
															return t.histogram.Percentile(p)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:274
	// _ = "end of CoverTab[96699]"
}

// Percentiles returns a slice of arbitrary percentiles of sampled values at
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:277
// the time the snapshot was taken.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:279
func (t *TimerSnapshot) Percentiles(ps []float64) []float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:279
	_go_fuzz_dep_.CoverTab[96700]++
															return t.histogram.Percentiles(ps)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:280
	// _ = "end of CoverTab[96700]"
}

// Rate1 returns the one-minute moving average rate of events per second at the
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:283
// time the snapshot was taken.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:285
func (t *TimerSnapshot) Rate1() float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:285
	_go_fuzz_dep_.CoverTab[96701]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:285
	return t.meter.Rate1()
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:285
	// _ = "end of CoverTab[96701]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:285
}

// Rate5 returns the five-minute moving average rate of events per second at
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:287
// the time the snapshot was taken.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:289
func (t *TimerSnapshot) Rate5() float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:289
	_go_fuzz_dep_.CoverTab[96702]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:289
	return t.meter.Rate5()
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:289
	// _ = "end of CoverTab[96702]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:289
}

// Rate15 returns the fifteen-minute moving average rate of events per second
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:291
// at the time the snapshot was taken.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:293
func (t *TimerSnapshot) Rate15() float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:293
	_go_fuzz_dep_.CoverTab[96703]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:293
	return t.meter.Rate15()
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:293
	// _ = "end of CoverTab[96703]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:293
}

// RateMean returns the meter's mean rate of events per second at the time the
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:295
// snapshot was taken.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:297
func (t *TimerSnapshot) RateMean() float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:297
	_go_fuzz_dep_.CoverTab[96704]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:297
	return t.meter.RateMean()
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:297
	// _ = "end of CoverTab[96704]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:297
}

// Snapshot returns the snapshot.
func (t *TimerSnapshot) Snapshot() Timer {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:300
	_go_fuzz_dep_.CoverTab[96705]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:300
	return t
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:300
	// _ = "end of CoverTab[96705]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:300
}

// StdDev returns the standard deviation of the values at the time the snapshot
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:302
// was taken.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:304
func (t *TimerSnapshot) StdDev() float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:304
	_go_fuzz_dep_.CoverTab[96706]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:304
	return t.histogram.StdDev()
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:304
	// _ = "end of CoverTab[96706]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:304
}

// Stop is a no-op.
func (t *TimerSnapshot) Stop()	{ _go_fuzz_dep_.CoverTab[96707]++; // _ = "end of CoverTab[96707]" }

// Sum returns the sum at the time the snapshot was taken.
func (t *TimerSnapshot) Sum() int64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:310
	_go_fuzz_dep_.CoverTab[96708]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:310
	return t.histogram.Sum()
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:310
	// _ = "end of CoverTab[96708]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:310
}

// Time panics.
func (*TimerSnapshot) Time(func()) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:313
	_go_fuzz_dep_.CoverTab[96709]++
															panic("Time called on a TimerSnapshot")
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:314
	// _ = "end of CoverTab[96709]"
}

// Update panics.
func (*TimerSnapshot) Update(time.Duration) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:318
	_go_fuzz_dep_.CoverTab[96710]++
															panic("Update called on a TimerSnapshot")
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:319
	// _ = "end of CoverTab[96710]"
}

// UpdateSince panics.
func (*TimerSnapshot) UpdateSince(time.Time) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:323
	_go_fuzz_dep_.CoverTab[96711]++
															panic("UpdateSince called on a TimerSnapshot")
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:324
	// _ = "end of CoverTab[96711]"
}

// Variance returns the variance of the values at the time the snapshot was
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:327
// taken.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:329
func (t *TimerSnapshot) Variance() float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:329
	_go_fuzz_dep_.CoverTab[96712]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:329
	return t.histogram.Variance()
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:329
	// _ = "end of CoverTab[96712]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:329
}

//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:329
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/timer.go:329
var _ = _go_fuzz_dep_.CoverTab
