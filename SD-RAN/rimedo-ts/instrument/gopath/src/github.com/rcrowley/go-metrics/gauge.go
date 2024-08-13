//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:1
package metrics

//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:1
)

import "sync/atomic"

// Gauges hold an int64 value that can be set arbitrarily.
type Gauge interface {
	Snapshot() Gauge
	Update(int64)
	Value() int64
}

// GetOrRegisterGauge returns an existing Gauge or constructs and registers a
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:12
// new StandardGauge.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:14
func GetOrRegisterGauge(name string, r Registry) Gauge {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:14
	_go_fuzz_dep_.CoverTab[96136]++
															if nil == r {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:15
		_go_fuzz_dep_.CoverTab[96138]++
																r = DefaultRegistry
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:16
		// _ = "end of CoverTab[96138]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:17
		_go_fuzz_dep_.CoverTab[96139]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:17
		// _ = "end of CoverTab[96139]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:17
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:17
	// _ = "end of CoverTab[96136]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:17
	_go_fuzz_dep_.CoverTab[96137]++
															return r.GetOrRegister(name, NewGauge).(Gauge)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:18
	// _ = "end of CoverTab[96137]"
}

// NewGauge constructs a new StandardGauge.
func NewGauge() Gauge {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:22
	_go_fuzz_dep_.CoverTab[96140]++
															if UseNilMetrics {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:23
		_go_fuzz_dep_.CoverTab[96142]++
																return NilGauge{}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:24
		// _ = "end of CoverTab[96142]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:25
		_go_fuzz_dep_.CoverTab[96143]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:25
		// _ = "end of CoverTab[96143]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:25
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:25
	// _ = "end of CoverTab[96140]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:25
	_go_fuzz_dep_.CoverTab[96141]++
															return &StandardGauge{0}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:26
	// _ = "end of CoverTab[96141]"
}

// NewRegisteredGauge constructs and registers a new StandardGauge.
func NewRegisteredGauge(name string, r Registry) Gauge {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:30
	_go_fuzz_dep_.CoverTab[96144]++
															c := NewGauge()
															if nil == r {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:32
		_go_fuzz_dep_.CoverTab[96146]++
																r = DefaultRegistry
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:33
		// _ = "end of CoverTab[96146]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:34
		_go_fuzz_dep_.CoverTab[96147]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:34
		// _ = "end of CoverTab[96147]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:34
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:34
	// _ = "end of CoverTab[96144]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:34
	_go_fuzz_dep_.CoverTab[96145]++
															r.Register(name, c)
															return c
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:36
	// _ = "end of CoverTab[96145]"
}

// NewFunctionalGauge constructs a new FunctionalGauge.
func NewFunctionalGauge(f func() int64) Gauge {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:40
	_go_fuzz_dep_.CoverTab[96148]++
															if UseNilMetrics {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:41
		_go_fuzz_dep_.CoverTab[96150]++
																return NilGauge{}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:42
		// _ = "end of CoverTab[96150]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:43
		_go_fuzz_dep_.CoverTab[96151]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:43
		// _ = "end of CoverTab[96151]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:43
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:43
	// _ = "end of CoverTab[96148]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:43
	_go_fuzz_dep_.CoverTab[96149]++
															return &FunctionalGauge{value: f}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:44
	// _ = "end of CoverTab[96149]"
}

// NewRegisteredFunctionalGauge constructs and registers a new StandardGauge.
func NewRegisteredFunctionalGauge(name string, r Registry, f func() int64) Gauge {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:48
	_go_fuzz_dep_.CoverTab[96152]++
															c := NewFunctionalGauge(f)
															if nil == r {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:50
		_go_fuzz_dep_.CoverTab[96154]++
																r = DefaultRegistry
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:51
		// _ = "end of CoverTab[96154]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:52
		_go_fuzz_dep_.CoverTab[96155]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:52
		// _ = "end of CoverTab[96155]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:52
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:52
	// _ = "end of CoverTab[96152]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:52
	_go_fuzz_dep_.CoverTab[96153]++
															r.Register(name, c)
															return c
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:54
	// _ = "end of CoverTab[96153]"
}

// GaugeSnapshot is a read-only copy of another Gauge.
type GaugeSnapshot int64

// Snapshot returns the snapshot.
func (g GaugeSnapshot) Snapshot() Gauge {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:61
	_go_fuzz_dep_.CoverTab[96156]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:61
	return g
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:61
	// _ = "end of CoverTab[96156]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:61
}

// Update panics.
func (GaugeSnapshot) Update(int64) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:64
	_go_fuzz_dep_.CoverTab[96157]++
															panic("Update called on a GaugeSnapshot")
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:65
	// _ = "end of CoverTab[96157]"
}

// Value returns the value at the time the snapshot was taken.
func (g GaugeSnapshot) Value() int64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:69
	_go_fuzz_dep_.CoverTab[96158]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:69
	return int64(g)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:69
	// _ = "end of CoverTab[96158]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:69
}

// NilGauge is a no-op Gauge.
type NilGauge struct{}

// Snapshot is a no-op.
func (NilGauge) Snapshot() Gauge {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:75
	_go_fuzz_dep_.CoverTab[96159]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:75
	return NilGauge{}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:75
	// _ = "end of CoverTab[96159]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:75
}

// Update is a no-op.
func (NilGauge) Update(v int64)	{ _go_fuzz_dep_.CoverTab[96160]++; // _ = "end of CoverTab[96160]" }

// Value is a no-op.
func (NilGauge) Value() int64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:81
	_go_fuzz_dep_.CoverTab[96161]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:81
	return 0
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:81
	// _ = "end of CoverTab[96161]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:81
}

// StandardGauge is the standard implementation of a Gauge and uses the
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:83
// sync/atomic package to manage a single int64 value.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:85
type StandardGauge struct {
	value int64
}

// Snapshot returns a read-only copy of the gauge.
func (g *StandardGauge) Snapshot() Gauge {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:90
	_go_fuzz_dep_.CoverTab[96162]++
															return GaugeSnapshot(g.Value())
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:91
	// _ = "end of CoverTab[96162]"
}

// Update updates the gauge's value.
func (g *StandardGauge) Update(v int64) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:95
	_go_fuzz_dep_.CoverTab[96163]++
															atomic.StoreInt64(&g.value, v)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:96
	// _ = "end of CoverTab[96163]"
}

// Value returns the gauge's current value.
func (g *StandardGauge) Value() int64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:100
	_go_fuzz_dep_.CoverTab[96164]++
															return atomic.LoadInt64(&g.value)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:101
	// _ = "end of CoverTab[96164]"
}

// FunctionalGauge returns value from given function
type FunctionalGauge struct {
	value func() int64
}

// Value returns the gauge's current value.
func (g FunctionalGauge) Value() int64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:110
	_go_fuzz_dep_.CoverTab[96165]++
															return g.value()
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:111
	// _ = "end of CoverTab[96165]"
}

// Snapshot returns the snapshot.
func (g FunctionalGauge) Snapshot() Gauge {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:115
	_go_fuzz_dep_.CoverTab[96166]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:115
	return GaugeSnapshot(g.Value())
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:115
	// _ = "end of CoverTab[96166]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:115
}

// Update panics.
func (FunctionalGauge) Update(int64) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:118
	_go_fuzz_dep_.CoverTab[96167]++
															panic("Update called on a FunctionalGauge")
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:119
	// _ = "end of CoverTab[96167]"
}

//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:120
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/gauge.go:120
var _ = _go_fuzz_dep_.CoverTab
