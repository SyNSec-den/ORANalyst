//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:1
package metrics

//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:1
)

import (
	"math"
	"sync/atomic"
)

// GaugeFloat64s hold a float64 value that can be set arbitrarily.
type GaugeFloat64 interface {
	Snapshot() GaugeFloat64
	Update(float64)
	Value() float64
}

// GetOrRegisterGaugeFloat64 returns an existing GaugeFloat64 or constructs and registers a
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:15
// new StandardGaugeFloat64.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:17
func GetOrRegisterGaugeFloat64(name string, r Registry) GaugeFloat64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:17
	_go_fuzz_dep_.CoverTab[96168]++
																if nil == r {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:18
		_go_fuzz_dep_.CoverTab[96170]++
																	r = DefaultRegistry
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:19
		// _ = "end of CoverTab[96170]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:20
		_go_fuzz_dep_.CoverTab[96171]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:20
		// _ = "end of CoverTab[96171]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:20
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:20
	// _ = "end of CoverTab[96168]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:20
	_go_fuzz_dep_.CoverTab[96169]++
																return r.GetOrRegister(name, NewGaugeFloat64()).(GaugeFloat64)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:21
	// _ = "end of CoverTab[96169]"
}

// NewGaugeFloat64 constructs a new StandardGaugeFloat64.
func NewGaugeFloat64() GaugeFloat64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:25
	_go_fuzz_dep_.CoverTab[96172]++
																if UseNilMetrics {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:26
		_go_fuzz_dep_.CoverTab[96174]++
																	return NilGaugeFloat64{}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:27
		// _ = "end of CoverTab[96174]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:28
		_go_fuzz_dep_.CoverTab[96175]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:28
		// _ = "end of CoverTab[96175]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:28
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:28
	// _ = "end of CoverTab[96172]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:28
	_go_fuzz_dep_.CoverTab[96173]++
																return &StandardGaugeFloat64{
		value: 0.0,
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:31
	// _ = "end of CoverTab[96173]"
}

// NewRegisteredGaugeFloat64 constructs and registers a new StandardGaugeFloat64.
func NewRegisteredGaugeFloat64(name string, r Registry) GaugeFloat64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:35
	_go_fuzz_dep_.CoverTab[96176]++
																c := NewGaugeFloat64()
																if nil == r {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:37
		_go_fuzz_dep_.CoverTab[96178]++
																	r = DefaultRegistry
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:38
		// _ = "end of CoverTab[96178]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:39
		_go_fuzz_dep_.CoverTab[96179]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:39
		// _ = "end of CoverTab[96179]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:39
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:39
	// _ = "end of CoverTab[96176]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:39
	_go_fuzz_dep_.CoverTab[96177]++
																r.Register(name, c)
																return c
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:41
	// _ = "end of CoverTab[96177]"
}

// NewFunctionalGauge constructs a new FunctionalGauge.
func NewFunctionalGaugeFloat64(f func() float64) GaugeFloat64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:45
	_go_fuzz_dep_.CoverTab[96180]++
																if UseNilMetrics {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:46
		_go_fuzz_dep_.CoverTab[96182]++
																	return NilGaugeFloat64{}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:47
		// _ = "end of CoverTab[96182]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:48
		_go_fuzz_dep_.CoverTab[96183]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:48
		// _ = "end of CoverTab[96183]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:48
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:48
	// _ = "end of CoverTab[96180]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:48
	_go_fuzz_dep_.CoverTab[96181]++
																return &FunctionalGaugeFloat64{value: f}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:49
	// _ = "end of CoverTab[96181]"
}

// NewRegisteredFunctionalGauge constructs and registers a new StandardGauge.
func NewRegisteredFunctionalGaugeFloat64(name string, r Registry, f func() float64) GaugeFloat64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:53
	_go_fuzz_dep_.CoverTab[96184]++
																c := NewFunctionalGaugeFloat64(f)
																if nil == r {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:55
		_go_fuzz_dep_.CoverTab[96186]++
																	r = DefaultRegistry
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:56
		// _ = "end of CoverTab[96186]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:57
		_go_fuzz_dep_.CoverTab[96187]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:57
		// _ = "end of CoverTab[96187]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:57
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:57
	// _ = "end of CoverTab[96184]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:57
	_go_fuzz_dep_.CoverTab[96185]++
																r.Register(name, c)
																return c
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:59
	// _ = "end of CoverTab[96185]"
}

// GaugeFloat64Snapshot is a read-only copy of another GaugeFloat64.
type GaugeFloat64Snapshot float64

// Snapshot returns the snapshot.
func (g GaugeFloat64Snapshot) Snapshot() GaugeFloat64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:66
	_go_fuzz_dep_.CoverTab[96188]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:66
	return g
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:66
	// _ = "end of CoverTab[96188]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:66
}

// Update panics.
func (GaugeFloat64Snapshot) Update(float64) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:69
	_go_fuzz_dep_.CoverTab[96189]++
																panic("Update called on a GaugeFloat64Snapshot")
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:70
	// _ = "end of CoverTab[96189]"
}

// Value returns the value at the time the snapshot was taken.
func (g GaugeFloat64Snapshot) Value() float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:74
	_go_fuzz_dep_.CoverTab[96190]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:74
	return float64(g)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:74
	// _ = "end of CoverTab[96190]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:74
}

// NilGauge is a no-op Gauge.
type NilGaugeFloat64 struct{}

// Snapshot is a no-op.
func (NilGaugeFloat64) Snapshot() GaugeFloat64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:80
	_go_fuzz_dep_.CoverTab[96191]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:80
	return NilGaugeFloat64{}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:80
	// _ = "end of CoverTab[96191]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:80
}

// Update is a no-op.
func (NilGaugeFloat64) Update(v float64) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:83
	_go_fuzz_dep_.CoverTab[96192]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:83
	// _ = "end of CoverTab[96192]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:83
}

// Value is a no-op.
func (NilGaugeFloat64) Value() float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:86
	_go_fuzz_dep_.CoverTab[96193]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:86
	return 0.0
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:86
	// _ = "end of CoverTab[96193]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:86
}

// StandardGaugeFloat64 is the standard implementation of a GaugeFloat64 and uses
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:88
// sync.Mutex to manage a single float64 value.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:90
type StandardGaugeFloat64 struct {
	value uint64
}

// Snapshot returns a read-only copy of the gauge.
func (g *StandardGaugeFloat64) Snapshot() GaugeFloat64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:95
	_go_fuzz_dep_.CoverTab[96194]++
																return GaugeFloat64Snapshot(g.Value())
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:96
	// _ = "end of CoverTab[96194]"
}

// Update updates the gauge's value.
func (g *StandardGaugeFloat64) Update(v float64) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:100
	_go_fuzz_dep_.CoverTab[96195]++
																atomic.StoreUint64(&g.value, math.Float64bits(v))
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:101
	// _ = "end of CoverTab[96195]"
}

// Value returns the gauge's current value.
func (g *StandardGaugeFloat64) Value() float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:105
	_go_fuzz_dep_.CoverTab[96196]++
																return math.Float64frombits(atomic.LoadUint64(&g.value))
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:106
	// _ = "end of CoverTab[96196]"
}

// FunctionalGaugeFloat64 returns value from given function
type FunctionalGaugeFloat64 struct {
	value func() float64
}

// Value returns the gauge's current value.
func (g FunctionalGaugeFloat64) Value() float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:115
	_go_fuzz_dep_.CoverTab[96197]++
																return g.value()
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:116
	// _ = "end of CoverTab[96197]"
}

// Snapshot returns the snapshot.
func (g FunctionalGaugeFloat64) Snapshot() GaugeFloat64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:120
	_go_fuzz_dep_.CoverTab[96198]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:120
	return GaugeFloat64Snapshot(g.Value())
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:120
	// _ = "end of CoverTab[96198]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:120
}

// Update panics.
func (FunctionalGaugeFloat64) Update(float64) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:123
	_go_fuzz_dep_.CoverTab[96199]++
																panic("Update called on a FunctionalGaugeFloat64")
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:124
	// _ = "end of CoverTab[96199]"
}

//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:125
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge_float64.go:125
var _ = _go_fuzz_dep_.CoverTab
