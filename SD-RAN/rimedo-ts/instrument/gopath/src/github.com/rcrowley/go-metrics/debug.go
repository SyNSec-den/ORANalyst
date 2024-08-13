//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/debug.go:1
package metrics

//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/debug.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/debug.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/debug.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/debug.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/debug.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/debug.go:1
)

import (
	"runtime/debug"
	"sync"
	"time"
)

var (
	debugMetrics	struct {
		GCStats	struct {
			LastGC	Gauge
			NumGC	Gauge
			Pause	Histogram
			//PauseQuantiles Histogram
			PauseTotal	Gauge
		}
		ReadGCStats	Timer
	}
	gcStats				debug.GCStats
	registerDebugMetricsOnce	= sync.Once{}
)

// Capture new values for the Go garbage collector statistics exported in
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/debug.go:24
// debug.GCStats.  This is designed to be called as a goroutine.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/debug.go:26
func CaptureDebugGCStats(r Registry, d time.Duration) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/debug.go:26
	_go_fuzz_dep_.CoverTab[96101]++
															for _ = range time.Tick(d) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/debug.go:27
		_go_fuzz_dep_.CoverTab[96102]++
																CaptureDebugGCStatsOnce(r)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/debug.go:28
		// _ = "end of CoverTab[96102]"
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/debug.go:29
	// _ = "end of CoverTab[96101]"
}

// Capture new values for the Go garbage collector statistics exported in
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/debug.go:32
// debug.GCStats.  This is designed to be called in a background goroutine.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/debug.go:32
// Giving a registry which has not been given to RegisterDebugGCStats will
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/debug.go:32
// panic.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/debug.go:32
//
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/debug.go:32
// Be careful (but much less so) with this because debug.ReadGCStats calls
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/debug.go:32
// the C function runtime·lock(runtime·mheap) which, while not a stop-the-world
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/debug.go:32
// operation, isn't something you want to be doing all the time.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/debug.go:40
func CaptureDebugGCStatsOnce(r Registry) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/debug.go:40
	_go_fuzz_dep_.CoverTab[96103]++
															lastGC := gcStats.LastGC
															t := time.Now()
															debug.ReadGCStats(&gcStats)
															debugMetrics.ReadGCStats.UpdateSince(t)

															debugMetrics.GCStats.LastGC.Update(int64(gcStats.LastGC.UnixNano()))
															debugMetrics.GCStats.NumGC.Update(int64(gcStats.NumGC))
															if lastGC != gcStats.LastGC && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/debug.go:48
		_go_fuzz_dep_.CoverTab[96105]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/debug.go:48
		return 0 < len(gcStats.Pause)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/debug.go:48
		// _ = "end of CoverTab[96105]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/debug.go:48
	}() {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/debug.go:48
		_go_fuzz_dep_.CoverTab[96106]++
																debugMetrics.GCStats.Pause.Update(int64(gcStats.Pause[0]))
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/debug.go:49
		// _ = "end of CoverTab[96106]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/debug.go:50
		_go_fuzz_dep_.CoverTab[96107]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/debug.go:50
		// _ = "end of CoverTab[96107]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/debug.go:50
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/debug.go:50
	// _ = "end of CoverTab[96103]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/debug.go:50
	_go_fuzz_dep_.CoverTab[96104]++

															debugMetrics.GCStats.PauseTotal.Update(int64(gcStats.PauseTotal))
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/debug.go:52
	// _ = "end of CoverTab[96104]"
}

// Register metrics for the Go garbage collector statistics exported in
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/debug.go:55
// debug.GCStats.  The metrics are named by their fully-qualified Go symbols,
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/debug.go:55
// i.e. debug.GCStats.PauseTotal.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/debug.go:58
func RegisterDebugGCStats(r Registry) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/debug.go:58
	_go_fuzz_dep_.CoverTab[96108]++
															registerDebugMetricsOnce.Do(func() {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/debug.go:59
		_go_fuzz_dep_.CoverTab[96109]++
																debugMetrics.GCStats.LastGC = NewGauge()
																debugMetrics.GCStats.NumGC = NewGauge()
																debugMetrics.GCStats.Pause = NewHistogram(NewExpDecaySample(1028, 0.015))

																debugMetrics.GCStats.PauseTotal = NewGauge()
																debugMetrics.ReadGCStats = NewTimer()

																r.Register("debug.GCStats.LastGC", debugMetrics.GCStats.LastGC)
																r.Register("debug.GCStats.NumGC", debugMetrics.GCStats.NumGC)
																r.Register("debug.GCStats.Pause", debugMetrics.GCStats.Pause)

																r.Register("debug.GCStats.PauseTotal", debugMetrics.GCStats.PauseTotal)
																r.Register("debug.ReadGCStats", debugMetrics.ReadGCStats)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/debug.go:72
		// _ = "end of CoverTab[96109]"
	})
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/debug.go:73
	// _ = "end of CoverTab[96108]"
}

// Allocate an initial slice for gcStats.Pause to avoid allocations during
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/debug.go:76
// normal operation.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/debug.go:78
func init() {
	gcStats.Pause = make([]time.Duration, 11)
}

//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/debug.go:80
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/debug.go:80
var _ = _go_fuzz_dep_.CoverTab
