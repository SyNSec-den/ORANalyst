//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/counter.go:1
package metrics

//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/counter.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/counter.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/counter.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/counter.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/counter.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/counter.go:1
)

import "sync/atomic"

// Counters hold an int64 value that can be incremented and decremented.
type Counter interface {
	Clear()
	Count() int64
	Dec(int64)
	Inc(int64)
	Snapshot() Counter
}

// GetOrRegisterCounter returns an existing Counter or constructs and registers
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/counter.go:14
// a new StandardCounter.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/counter.go:16
func GetOrRegisterCounter(name string, r Registry) Counter {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/counter.go:16
	_go_fuzz_dep_.CoverTab[96074]++
															if nil == r {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/counter.go:17
		_go_fuzz_dep_.CoverTab[96076]++
																r = DefaultRegistry
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/counter.go:18
		// _ = "end of CoverTab[96076]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/counter.go:19
		_go_fuzz_dep_.CoverTab[96077]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/counter.go:19
		// _ = "end of CoverTab[96077]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/counter.go:19
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/counter.go:19
	// _ = "end of CoverTab[96074]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/counter.go:19
	_go_fuzz_dep_.CoverTab[96075]++
															return r.GetOrRegister(name, NewCounter).(Counter)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/counter.go:20
	// _ = "end of CoverTab[96075]"
}

// NewCounter constructs a new StandardCounter.
func NewCounter() Counter {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/counter.go:24
	_go_fuzz_dep_.CoverTab[96078]++
															if UseNilMetrics {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/counter.go:25
		_go_fuzz_dep_.CoverTab[96080]++
																return NilCounter{}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/counter.go:26
		// _ = "end of CoverTab[96080]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/counter.go:27
		_go_fuzz_dep_.CoverTab[96081]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/counter.go:27
		// _ = "end of CoverTab[96081]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/counter.go:27
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/counter.go:27
	// _ = "end of CoverTab[96078]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/counter.go:27
	_go_fuzz_dep_.CoverTab[96079]++
															return &StandardCounter{0}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/counter.go:28
	// _ = "end of CoverTab[96079]"
}

// NewRegisteredCounter constructs and registers a new StandardCounter.
func NewRegisteredCounter(name string, r Registry) Counter {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/counter.go:32
	_go_fuzz_dep_.CoverTab[96082]++
															c := NewCounter()
															if nil == r {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/counter.go:34
		_go_fuzz_dep_.CoverTab[96084]++
																r = DefaultRegistry
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/counter.go:35
		// _ = "end of CoverTab[96084]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/counter.go:36
		_go_fuzz_dep_.CoverTab[96085]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/counter.go:36
		// _ = "end of CoverTab[96085]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/counter.go:36
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/counter.go:36
	// _ = "end of CoverTab[96082]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/counter.go:36
	_go_fuzz_dep_.CoverTab[96083]++
															r.Register(name, c)
															return c
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/counter.go:38
	// _ = "end of CoverTab[96083]"
}

// CounterSnapshot is a read-only copy of another Counter.
type CounterSnapshot int64

// Clear panics.
func (CounterSnapshot) Clear() {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/counter.go:45
	_go_fuzz_dep_.CoverTab[96086]++
															panic("Clear called on a CounterSnapshot")
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/counter.go:46
	// _ = "end of CoverTab[96086]"
}

// Count returns the count at the time the snapshot was taken.
func (c CounterSnapshot) Count() int64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/counter.go:50
	_go_fuzz_dep_.CoverTab[96087]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/counter.go:50
	return int64(c)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/counter.go:50
	// _ = "end of CoverTab[96087]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/counter.go:50
}

// Dec panics.
func (CounterSnapshot) Dec(int64) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/counter.go:53
	_go_fuzz_dep_.CoverTab[96088]++
															panic("Dec called on a CounterSnapshot")
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/counter.go:54
	// _ = "end of CoverTab[96088]"
}

// Inc panics.
func (CounterSnapshot) Inc(int64) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/counter.go:58
	_go_fuzz_dep_.CoverTab[96089]++
															panic("Inc called on a CounterSnapshot")
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/counter.go:59
	// _ = "end of CoverTab[96089]"
}

// Snapshot returns the snapshot.
func (c CounterSnapshot) Snapshot() Counter {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/counter.go:63
	_go_fuzz_dep_.CoverTab[96090]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/counter.go:63
	return c
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/counter.go:63
	// _ = "end of CoverTab[96090]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/counter.go:63
}

// NilCounter is a no-op Counter.
type NilCounter struct{}

// Clear is a no-op.
func (NilCounter) Clear()	{ _go_fuzz_dep_.CoverTab[96091]++; // _ = "end of CoverTab[96091]" }

// Count is a no-op.
func (NilCounter) Count() int64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/counter.go:72
	_go_fuzz_dep_.CoverTab[96092]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/counter.go:72
	return 0
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/counter.go:72
	// _ = "end of CoverTab[96092]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/counter.go:72
}

// Dec is a no-op.
func (NilCounter) Dec(i int64)	{ _go_fuzz_dep_.CoverTab[96093]++; // _ = "end of CoverTab[96093]" }

// Inc is a no-op.
func (NilCounter) Inc(i int64)	{ _go_fuzz_dep_.CoverTab[96094]++; // _ = "end of CoverTab[96094]" }

// Snapshot is a no-op.
func (NilCounter) Snapshot() Counter {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/counter.go:81
	_go_fuzz_dep_.CoverTab[96095]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/counter.go:81
	return NilCounter{}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/counter.go:81
	// _ = "end of CoverTab[96095]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/counter.go:81
}

// StandardCounter is the standard implementation of a Counter and uses the
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/counter.go:83
// sync/atomic package to manage a single int64 value.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/counter.go:85
type StandardCounter struct {
	count int64
}

// Clear sets the counter to zero.
func (c *StandardCounter) Clear() {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/counter.go:90
	_go_fuzz_dep_.CoverTab[96096]++
															atomic.StoreInt64(&c.count, 0)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/counter.go:91
	// _ = "end of CoverTab[96096]"
}

// Count returns the current count.
func (c *StandardCounter) Count() int64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/counter.go:95
	_go_fuzz_dep_.CoverTab[96097]++
															return atomic.LoadInt64(&c.count)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/counter.go:96
	// _ = "end of CoverTab[96097]"
}

// Dec decrements the counter by the given amount.
func (c *StandardCounter) Dec(i int64) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/counter.go:100
	_go_fuzz_dep_.CoverTab[96098]++
															atomic.AddInt64(&c.count, -i)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/counter.go:101
	// _ = "end of CoverTab[96098]"
}

// Inc increments the counter by the given amount.
func (c *StandardCounter) Inc(i int64) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/counter.go:105
	_go_fuzz_dep_.CoverTab[96099]++
															atomic.AddInt64(&c.count, i)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/counter.go:106
	// _ = "end of CoverTab[96099]"
}

// Snapshot returns a read-only copy of the counter.
func (c *StandardCounter) Snapshot() Counter {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/counter.go:110
	_go_fuzz_dep_.CoverTab[96100]++
															return CounterSnapshot(c.Count())
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/counter.go:111
	// _ = "end of CoverTab[96100]"
}

//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/counter.go:112
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/counter.go:112
var _ = _go_fuzz_dep_.CoverTab
