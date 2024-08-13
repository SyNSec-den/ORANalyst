//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:1
package metrics

//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:1
)

import (
	"math"
	"sync"
	"sync/atomic"
	"time"
)

// Meters count events to produce exponentially-weighted moving average rates
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:10
// at one-, five-, and fifteen-minutes and a mean rate.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:12
type Meter interface {
	Count() int64
	Mark(int64)
	Rate1() float64
	Rate5() float64
	Rate15() float64
	RateMean() float64
	Snapshot() Meter
	Stop()
}

// GetOrRegisterMeter returns an existing Meter or constructs and registers a
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:23
// new StandardMeter.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:23
// Be sure to unregister the meter from the registry once it is of no use to
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:23
// allow for garbage collection.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:27
func GetOrRegisterMeter(name string, r Registry) Meter {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:27
	_go_fuzz_dep_.CoverTab[96308]++
															if nil == r {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:28
		_go_fuzz_dep_.CoverTab[96310]++
																r = DefaultRegistry
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:29
		// _ = "end of CoverTab[96310]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:30
		_go_fuzz_dep_.CoverTab[96311]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:30
		// _ = "end of CoverTab[96311]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:30
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:30
	// _ = "end of CoverTab[96308]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:30
	_go_fuzz_dep_.CoverTab[96309]++
															return r.GetOrRegister(name, NewMeter).(Meter)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:31
	// _ = "end of CoverTab[96309]"
}

// NewMeter constructs a new StandardMeter and launches a goroutine.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:34
// Be sure to call Stop() once the meter is of no use to allow for garbage collection.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:36
func NewMeter() Meter {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:36
	_go_fuzz_dep_.CoverTab[96312]++
															if UseNilMetrics {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:37
		_go_fuzz_dep_.CoverTab[96315]++
																return NilMeter{}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:38
		// _ = "end of CoverTab[96315]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:39
		_go_fuzz_dep_.CoverTab[96316]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:39
		// _ = "end of CoverTab[96316]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:39
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:39
	// _ = "end of CoverTab[96312]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:39
	_go_fuzz_dep_.CoverTab[96313]++
															m := newStandardMeter()
															arbiter.Lock()
															defer arbiter.Unlock()
															arbiter.meters[m] = struct{}{}
															if !arbiter.started {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:44
		_go_fuzz_dep_.CoverTab[96317]++
																arbiter.started = true
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:45
		_curRoutineNum111_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:45
		_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum111_)
																go arbiter.tick()
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:46
		// _ = "end of CoverTab[96317]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:47
		_go_fuzz_dep_.CoverTab[96318]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:47
		// _ = "end of CoverTab[96318]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:47
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:47
	// _ = "end of CoverTab[96313]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:47
	_go_fuzz_dep_.CoverTab[96314]++
															return m
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:48
	// _ = "end of CoverTab[96314]"
}

// NewMeter constructs and registers a new StandardMeter and launches a
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:51
// goroutine.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:51
// Be sure to unregister the meter from the registry once it is of no use to
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:51
// allow for garbage collection.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:55
func NewRegisteredMeter(name string, r Registry) Meter {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:55
	_go_fuzz_dep_.CoverTab[96319]++
															c := NewMeter()
															if nil == r {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:57
		_go_fuzz_dep_.CoverTab[96321]++
																r = DefaultRegistry
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:58
		// _ = "end of CoverTab[96321]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:59
		_go_fuzz_dep_.CoverTab[96322]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:59
		// _ = "end of CoverTab[96322]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:59
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:59
	// _ = "end of CoverTab[96319]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:59
	_go_fuzz_dep_.CoverTab[96320]++
															r.Register(name, c)
															return c
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:61
	// _ = "end of CoverTab[96320]"
}

// MeterSnapshot is a read-only copy of another Meter.
type MeterSnapshot struct {
	count				int64
	rate1, rate5, rate15, rateMean	uint64
}

// Count returns the count of events at the time the snapshot was taken.
func (m *MeterSnapshot) Count() int64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:71
	_go_fuzz_dep_.CoverTab[96323]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:71
	return m.count
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:71
	// _ = "end of CoverTab[96323]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:71
}

// Mark panics.
func (*MeterSnapshot) Mark(n int64) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:74
	_go_fuzz_dep_.CoverTab[96324]++
															panic("Mark called on a MeterSnapshot")
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:75
	// _ = "end of CoverTab[96324]"
}

// Rate1 returns the one-minute moving average rate of events per second at the
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:78
// time the snapshot was taken.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:80
func (m *MeterSnapshot) Rate1() float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:80
	_go_fuzz_dep_.CoverTab[96325]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:80
	return math.Float64frombits(m.rate1)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:80
	// _ = "end of CoverTab[96325]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:80
}

// Rate5 returns the five-minute moving average rate of events per second at
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:82
// the time the snapshot was taken.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:84
func (m *MeterSnapshot) Rate5() float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:84
	_go_fuzz_dep_.CoverTab[96326]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:84
	return math.Float64frombits(m.rate5)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:84
	// _ = "end of CoverTab[96326]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:84
}

// Rate15 returns the fifteen-minute moving average rate of events per second
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:86
// at the time the snapshot was taken.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:88
func (m *MeterSnapshot) Rate15() float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:88
	_go_fuzz_dep_.CoverTab[96327]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:88
	return math.Float64frombits(m.rate15)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:88
	// _ = "end of CoverTab[96327]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:88
}

// RateMean returns the meter's mean rate of events per second at the time the
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:90
// snapshot was taken.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:92
func (m *MeterSnapshot) RateMean() float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:92
	_go_fuzz_dep_.CoverTab[96328]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:92
	return math.Float64frombits(m.rateMean)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:92
	// _ = "end of CoverTab[96328]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:92
}

// Snapshot returns the snapshot.
func (m *MeterSnapshot) Snapshot() Meter {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:95
	_go_fuzz_dep_.CoverTab[96329]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:95
	return m
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:95
	// _ = "end of CoverTab[96329]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:95
}

// Stop is a no-op.
func (m *MeterSnapshot) Stop()	{ _go_fuzz_dep_.CoverTab[96330]++; // _ = "end of CoverTab[96330]" }

// NilMeter is a no-op Meter.
type NilMeter struct{}

// Count is a no-op.
func (NilMeter) Count() int64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:104
	_go_fuzz_dep_.CoverTab[96331]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:104
	return 0
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:104
	// _ = "end of CoverTab[96331]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:104
}

// Mark is a no-op.
func (NilMeter) Mark(n int64)	{ _go_fuzz_dep_.CoverTab[96332]++; // _ = "end of CoverTab[96332]" }

// Rate1 is a no-op.
func (NilMeter) Rate1() float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:110
	_go_fuzz_dep_.CoverTab[96333]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:110
	return 0.0
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:110
	// _ = "end of CoverTab[96333]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:110
}

// Rate5 is a no-op.
func (NilMeter) Rate5() float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:113
	_go_fuzz_dep_.CoverTab[96334]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:113
	return 0.0
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:113
	// _ = "end of CoverTab[96334]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:113
}

// Rate15is a no-op.
func (NilMeter) Rate15() float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:116
	_go_fuzz_dep_.CoverTab[96335]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:116
	return 0.0
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:116
	// _ = "end of CoverTab[96335]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:116
}

// RateMean is a no-op.
func (NilMeter) RateMean() float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:119
	_go_fuzz_dep_.CoverTab[96336]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:119
	return 0.0
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:119
	// _ = "end of CoverTab[96336]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:119
}

// Snapshot is a no-op.
func (NilMeter) Snapshot() Meter {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:122
	_go_fuzz_dep_.CoverTab[96337]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:122
	return NilMeter{}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:122
	// _ = "end of CoverTab[96337]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:122
}

// Stop is a no-op.
func (NilMeter) Stop()	{ _go_fuzz_dep_.CoverTab[96338]++; // _ = "end of CoverTab[96338]" }

// StandardMeter is the standard implementation of a Meter.
type StandardMeter struct {
	snapshot	*MeterSnapshot
	a1, a5, a15	EWMA
	startTime	time.Time
	stopped		uint32
}

func newStandardMeter() *StandardMeter {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:135
	_go_fuzz_dep_.CoverTab[96339]++
															return &StandardMeter{
		snapshot:	&MeterSnapshot{},
		a1:		NewEWMA1(),
		a5:		NewEWMA5(),
		a15:		NewEWMA15(),
		startTime:	time.Now(),
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:142
	// _ = "end of CoverTab[96339]"
}

// Stop stops the meter, Mark() will be a no-op if you use it after being stopped.
func (m *StandardMeter) Stop() {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:146
	_go_fuzz_dep_.CoverTab[96340]++
															if atomic.CompareAndSwapUint32(&m.stopped, 0, 1) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:147
		_go_fuzz_dep_.CoverTab[96341]++
																arbiter.Lock()
																delete(arbiter.meters, m)
																arbiter.Unlock()
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:150
		// _ = "end of CoverTab[96341]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:151
		_go_fuzz_dep_.CoverTab[96342]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:151
		// _ = "end of CoverTab[96342]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:151
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:151
	// _ = "end of CoverTab[96340]"
}

// Count returns the number of events recorded.
func (m *StandardMeter) Count() int64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:155
	_go_fuzz_dep_.CoverTab[96343]++
															return atomic.LoadInt64(&m.snapshot.count)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:156
	// _ = "end of CoverTab[96343]"
}

// Mark records the occurance of n events.
func (m *StandardMeter) Mark(n int64) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:160
	_go_fuzz_dep_.CoverTab[96344]++
															if atomic.LoadUint32(&m.stopped) == 1 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:161
		_go_fuzz_dep_.CoverTab[96346]++
																return
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:162
		// _ = "end of CoverTab[96346]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:163
		_go_fuzz_dep_.CoverTab[96347]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:163
		// _ = "end of CoverTab[96347]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:163
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:163
	// _ = "end of CoverTab[96344]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:163
	_go_fuzz_dep_.CoverTab[96345]++

															atomic.AddInt64(&m.snapshot.count, n)

															m.a1.Update(n)
															m.a5.Update(n)
															m.a15.Update(n)
															m.updateSnapshot()
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:170
	// _ = "end of CoverTab[96345]"
}

// Rate1 returns the one-minute moving average rate of events per second.
func (m *StandardMeter) Rate1() float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:174
	_go_fuzz_dep_.CoverTab[96348]++
															return math.Float64frombits(atomic.LoadUint64(&m.snapshot.rate1))
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:175
	// _ = "end of CoverTab[96348]"
}

// Rate5 returns the five-minute moving average rate of events per second.
func (m *StandardMeter) Rate5() float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:179
	_go_fuzz_dep_.CoverTab[96349]++
															return math.Float64frombits(atomic.LoadUint64(&m.snapshot.rate5))
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:180
	// _ = "end of CoverTab[96349]"
}

// Rate15 returns the fifteen-minute moving average rate of events per second.
func (m *StandardMeter) Rate15() float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:184
	_go_fuzz_dep_.CoverTab[96350]++
															return math.Float64frombits(atomic.LoadUint64(&m.snapshot.rate15))
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:185
	// _ = "end of CoverTab[96350]"
}

// RateMean returns the meter's mean rate of events per second.
func (m *StandardMeter) RateMean() float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:189
	_go_fuzz_dep_.CoverTab[96351]++
															return math.Float64frombits(atomic.LoadUint64(&m.snapshot.rateMean))
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:190
	// _ = "end of CoverTab[96351]"
}

// Snapshot returns a read-only copy of the meter.
func (m *StandardMeter) Snapshot() Meter {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:194
	_go_fuzz_dep_.CoverTab[96352]++
															copiedSnapshot := MeterSnapshot{
		count:		atomic.LoadInt64(&m.snapshot.count),
		rate1:		atomic.LoadUint64(&m.snapshot.rate1),
		rate5:		atomic.LoadUint64(&m.snapshot.rate5),
		rate15:		atomic.LoadUint64(&m.snapshot.rate15),
		rateMean:	atomic.LoadUint64(&m.snapshot.rateMean),
	}
															return &copiedSnapshot
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:202
	// _ = "end of CoverTab[96352]"
}

func (m *StandardMeter) updateSnapshot() {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:205
	_go_fuzz_dep_.CoverTab[96353]++
															rate1 := math.Float64bits(m.a1.Rate())
															rate5 := math.Float64bits(m.a5.Rate())
															rate15 := math.Float64bits(m.a15.Rate())
															rateMean := math.Float64bits(float64(m.Count()) / time.Since(m.startTime).Seconds())

															atomic.StoreUint64(&m.snapshot.rate1, rate1)
															atomic.StoreUint64(&m.snapshot.rate5, rate5)
															atomic.StoreUint64(&m.snapshot.rate15, rate15)
															atomic.StoreUint64(&m.snapshot.rateMean, rateMean)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:214
	// _ = "end of CoverTab[96353]"
}

func (m *StandardMeter) tick() {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:217
	_go_fuzz_dep_.CoverTab[96354]++
															m.a1.Tick()
															m.a5.Tick()
															m.a15.Tick()
															m.updateSnapshot()
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:221
	// _ = "end of CoverTab[96354]"
}

// meterArbiter ticks meters every 5s from a single goroutine.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:224
// meters are references in a set for future stopping.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:226
type meterArbiter struct {
	sync.RWMutex
	started	bool
	meters	map[*StandardMeter]struct{}
	ticker	*time.Ticker
}

var arbiter = meterArbiter{ticker: time.NewTicker(5e9), meters: make(map[*StandardMeter]struct{})}

// Ticks meters on the scheduled interval
func (ma *meterArbiter) tick() {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:236
	_go_fuzz_dep_.CoverTab[96355]++
															for {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:237
		_go_fuzz_dep_.CoverTab[96356]++
																select {
		case <-ma.ticker.C:
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:239
			_go_fuzz_dep_.CoverTab[96357]++
																	ma.tickMeters()
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:240
			// _ = "end of CoverTab[96357]"
		}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:241
		// _ = "end of CoverTab[96356]"
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:242
	// _ = "end of CoverTab[96355]"
}

func (ma *meterArbiter) tickMeters() {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:245
	_go_fuzz_dep_.CoverTab[96358]++
															ma.RLock()
															defer ma.RUnlock()
															for meter := range ma.meters {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:248
		_go_fuzz_dep_.CoverTab[96359]++
																meter.tick()
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:249
		// _ = "end of CoverTab[96359]"
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:250
	// _ = "end of CoverTab[96358]"
}

//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:251
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/meter.go:251
var _ = _go_fuzz_dep_.CoverTab
