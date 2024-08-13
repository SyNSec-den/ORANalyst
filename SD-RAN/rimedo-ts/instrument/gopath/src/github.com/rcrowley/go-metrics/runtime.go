//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/runtime.go:1
package metrics

//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/runtime.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/runtime.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/runtime.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/runtime.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/runtime.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/runtime.go:1
)

import (
	"runtime"
	"runtime/pprof"
	"sync"
	"time"
)

var (
	memStats	runtime.MemStats
	runtimeMetrics	struct {
		MemStats	struct {
			Alloc		Gauge
			BuckHashSys	Gauge
			DebugGC		Gauge
			EnableGC	Gauge
			Frees		Gauge
			HeapAlloc	Gauge
			HeapIdle	Gauge
			HeapInuse	Gauge
			HeapObjects	Gauge
			HeapReleased	Gauge
			HeapSys		Gauge
			LastGC		Gauge
			Lookups		Gauge
			Mallocs		Gauge
			MCacheInuse	Gauge
			MCacheSys	Gauge
			MSpanInuse	Gauge
			MSpanSys	Gauge
			NextGC		Gauge
			NumGC		Gauge
			GCCPUFraction	GaugeFloat64
			PauseNs		Histogram
			PauseTotalNs	Gauge
			StackInuse	Gauge
			StackSys	Gauge
			Sys		Gauge
			TotalAlloc	Gauge
		}
		NumCgoCall	Gauge
		NumGoroutine	Gauge
		NumThread	Gauge
		ReadMemStats	Timer
	}
	frees		uint64
	lookups		uint64
	mallocs		uint64
	numGC		uint32
	numCgoCalls	int64

	threadCreateProfile		= pprof.Lookup("threadcreate")
	registerRuntimeMetricsOnce	= sync.Once{}
)

// Capture new values for the Go runtime statistics exported in
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/runtime.go:57
// runtime.MemStats.  This is designed to be called as a goroutine.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/runtime.go:59
func CaptureRuntimeMemStats(r Registry, d time.Duration) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/runtime.go:59
	_go_fuzz_dep_.CoverTab[96462]++
															for _ = range time.Tick(d) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/runtime.go:60
		_go_fuzz_dep_.CoverTab[96463]++
																CaptureRuntimeMemStatsOnce(r)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/runtime.go:61
		// _ = "end of CoverTab[96463]"
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/runtime.go:62
	// _ = "end of CoverTab[96462]"
}

// Capture new values for the Go runtime statistics exported in
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/runtime.go:65
// runtime.MemStats.  This is designed to be called in a background
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/runtime.go:65
// goroutine.  Giving a registry which has not been given to
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/runtime.go:65
// RegisterRuntimeMemStats will panic.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/runtime.go:65
//
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/runtime.go:65
// Be very careful with this because runtime.ReadMemStats calls the C
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/runtime.go:65
// functions runtime·semacquire(&runtime·worldsema) and runtime·stoptheworld()
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/runtime.go:65
// and that last one does what it says on the tin.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/runtime.go:73
func CaptureRuntimeMemStatsOnce(r Registry) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/runtime.go:73
	_go_fuzz_dep_.CoverTab[96464]++
															t := time.Now()
															runtime.ReadMemStats(&memStats)
															runtimeMetrics.ReadMemStats.UpdateSince(t)

															runtimeMetrics.MemStats.Alloc.Update(int64(memStats.Alloc))
															runtimeMetrics.MemStats.BuckHashSys.Update(int64(memStats.BuckHashSys))
															if memStats.DebugGC {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/runtime.go:80
		_go_fuzz_dep_.CoverTab[96468]++
																runtimeMetrics.MemStats.DebugGC.Update(1)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/runtime.go:81
		// _ = "end of CoverTab[96468]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/runtime.go:82
		_go_fuzz_dep_.CoverTab[96469]++
																runtimeMetrics.MemStats.DebugGC.Update(0)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/runtime.go:83
		// _ = "end of CoverTab[96469]"
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/runtime.go:84
	// _ = "end of CoverTab[96464]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/runtime.go:84
	_go_fuzz_dep_.CoverTab[96465]++
															if memStats.EnableGC {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/runtime.go:85
		_go_fuzz_dep_.CoverTab[96470]++
																runtimeMetrics.MemStats.EnableGC.Update(1)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/runtime.go:86
		// _ = "end of CoverTab[96470]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/runtime.go:87
		_go_fuzz_dep_.CoverTab[96471]++
																runtimeMetrics.MemStats.EnableGC.Update(0)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/runtime.go:88
		// _ = "end of CoverTab[96471]"
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/runtime.go:89
	// _ = "end of CoverTab[96465]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/runtime.go:89
	_go_fuzz_dep_.CoverTab[96466]++

															runtimeMetrics.MemStats.Frees.Update(int64(memStats.Frees - frees))
															runtimeMetrics.MemStats.HeapAlloc.Update(int64(memStats.HeapAlloc))
															runtimeMetrics.MemStats.HeapIdle.Update(int64(memStats.HeapIdle))
															runtimeMetrics.MemStats.HeapInuse.Update(int64(memStats.HeapInuse))
															runtimeMetrics.MemStats.HeapObjects.Update(int64(memStats.HeapObjects))
															runtimeMetrics.MemStats.HeapReleased.Update(int64(memStats.HeapReleased))
															runtimeMetrics.MemStats.HeapSys.Update(int64(memStats.HeapSys))
															runtimeMetrics.MemStats.LastGC.Update(int64(memStats.LastGC))
															runtimeMetrics.MemStats.Lookups.Update(int64(memStats.Lookups - lookups))
															runtimeMetrics.MemStats.Mallocs.Update(int64(memStats.Mallocs - mallocs))
															runtimeMetrics.MemStats.MCacheInuse.Update(int64(memStats.MCacheInuse))
															runtimeMetrics.MemStats.MCacheSys.Update(int64(memStats.MCacheSys))
															runtimeMetrics.MemStats.MSpanInuse.Update(int64(memStats.MSpanInuse))
															runtimeMetrics.MemStats.MSpanSys.Update(int64(memStats.MSpanSys))
															runtimeMetrics.MemStats.NextGC.Update(int64(memStats.NextGC))
															runtimeMetrics.MemStats.NumGC.Update(int64(memStats.NumGC - numGC))
															runtimeMetrics.MemStats.GCCPUFraction.Update(gcCPUFraction(&memStats))

//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/runtime.go:110
	i := numGC % uint32(len(memStats.PauseNs))
	ii := memStats.NumGC % uint32(len(memStats.PauseNs))
	if memStats.NumGC-numGC >= uint32(len(memStats.PauseNs)) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/runtime.go:112
		_go_fuzz_dep_.CoverTab[96472]++
																for i = 0; i < uint32(len(memStats.PauseNs)); i++ {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/runtime.go:113
			_go_fuzz_dep_.CoverTab[96473]++
																	runtimeMetrics.MemStats.PauseNs.Update(int64(memStats.PauseNs[i]))
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/runtime.go:114
			// _ = "end of CoverTab[96473]"
		}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/runtime.go:115
		// _ = "end of CoverTab[96472]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/runtime.go:116
		_go_fuzz_dep_.CoverTab[96474]++
																if i > ii {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/runtime.go:117
			_go_fuzz_dep_.CoverTab[96476]++
																	for ; i < uint32(len(memStats.PauseNs)); i++ {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/runtime.go:118
				_go_fuzz_dep_.CoverTab[96478]++
																		runtimeMetrics.MemStats.PauseNs.Update(int64(memStats.PauseNs[i]))
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/runtime.go:119
				// _ = "end of CoverTab[96478]"
			}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/runtime.go:120
			// _ = "end of CoverTab[96476]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/runtime.go:120
			_go_fuzz_dep_.CoverTab[96477]++
																	i = 0
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/runtime.go:121
			// _ = "end of CoverTab[96477]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/runtime.go:122
			_go_fuzz_dep_.CoverTab[96479]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/runtime.go:122
			// _ = "end of CoverTab[96479]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/runtime.go:122
		}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/runtime.go:122
		// _ = "end of CoverTab[96474]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/runtime.go:122
		_go_fuzz_dep_.CoverTab[96475]++
																for ; i < ii; i++ {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/runtime.go:123
			_go_fuzz_dep_.CoverTab[96480]++
																	runtimeMetrics.MemStats.PauseNs.Update(int64(memStats.PauseNs[i]))
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/runtime.go:124
			// _ = "end of CoverTab[96480]"
		}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/runtime.go:125
		// _ = "end of CoverTab[96475]"
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/runtime.go:126
	// _ = "end of CoverTab[96466]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/runtime.go:126
	_go_fuzz_dep_.CoverTab[96467]++
															frees = memStats.Frees
															lookups = memStats.Lookups
															mallocs = memStats.Mallocs
															numGC = memStats.NumGC

															runtimeMetrics.MemStats.PauseTotalNs.Update(int64(memStats.PauseTotalNs))
															runtimeMetrics.MemStats.StackInuse.Update(int64(memStats.StackInuse))
															runtimeMetrics.MemStats.StackSys.Update(int64(memStats.StackSys))
															runtimeMetrics.MemStats.Sys.Update(int64(memStats.Sys))
															runtimeMetrics.MemStats.TotalAlloc.Update(int64(memStats.TotalAlloc))

															currentNumCgoCalls := numCgoCall()
															runtimeMetrics.NumCgoCall.Update(currentNumCgoCalls - numCgoCalls)
															numCgoCalls = currentNumCgoCalls

															runtimeMetrics.NumGoroutine.Update(int64(runtime.NumGoroutine()))

															runtimeMetrics.NumThread.Update(int64(threadCreateProfile.Count()))
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/runtime.go:144
	// _ = "end of CoverTab[96467]"
}

// Register runtimeMetrics for the Go runtime statistics exported in runtime and
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/runtime.go:147
// specifically runtime.MemStats.  The runtimeMetrics are named by their
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/runtime.go:147
// fully-qualified Go symbols, i.e. runtime.MemStats.Alloc.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/runtime.go:150
func RegisterRuntimeMemStats(r Registry) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/runtime.go:150
	_go_fuzz_dep_.CoverTab[96481]++
															registerRuntimeMetricsOnce.Do(func() {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/runtime.go:151
		_go_fuzz_dep_.CoverTab[96482]++
																runtimeMetrics.MemStats.Alloc = NewGauge()
																runtimeMetrics.MemStats.BuckHashSys = NewGauge()
																runtimeMetrics.MemStats.DebugGC = NewGauge()
																runtimeMetrics.MemStats.EnableGC = NewGauge()
																runtimeMetrics.MemStats.Frees = NewGauge()
																runtimeMetrics.MemStats.HeapAlloc = NewGauge()
																runtimeMetrics.MemStats.HeapIdle = NewGauge()
																runtimeMetrics.MemStats.HeapInuse = NewGauge()
																runtimeMetrics.MemStats.HeapObjects = NewGauge()
																runtimeMetrics.MemStats.HeapReleased = NewGauge()
																runtimeMetrics.MemStats.HeapSys = NewGauge()
																runtimeMetrics.MemStats.LastGC = NewGauge()
																runtimeMetrics.MemStats.Lookups = NewGauge()
																runtimeMetrics.MemStats.Mallocs = NewGauge()
																runtimeMetrics.MemStats.MCacheInuse = NewGauge()
																runtimeMetrics.MemStats.MCacheSys = NewGauge()
																runtimeMetrics.MemStats.MSpanInuse = NewGauge()
																runtimeMetrics.MemStats.MSpanSys = NewGauge()
																runtimeMetrics.MemStats.NextGC = NewGauge()
																runtimeMetrics.MemStats.NumGC = NewGauge()
																runtimeMetrics.MemStats.GCCPUFraction = NewGaugeFloat64()
																runtimeMetrics.MemStats.PauseNs = NewHistogram(NewExpDecaySample(1028, 0.015))
																runtimeMetrics.MemStats.PauseTotalNs = NewGauge()
																runtimeMetrics.MemStats.StackInuse = NewGauge()
																runtimeMetrics.MemStats.StackSys = NewGauge()
																runtimeMetrics.MemStats.Sys = NewGauge()
																runtimeMetrics.MemStats.TotalAlloc = NewGauge()
																runtimeMetrics.NumCgoCall = NewGauge()
																runtimeMetrics.NumGoroutine = NewGauge()
																runtimeMetrics.NumThread = NewGauge()
																runtimeMetrics.ReadMemStats = NewTimer()

																r.Register("runtime.MemStats.Alloc", runtimeMetrics.MemStats.Alloc)
																r.Register("runtime.MemStats.BuckHashSys", runtimeMetrics.MemStats.BuckHashSys)
																r.Register("runtime.MemStats.DebugGC", runtimeMetrics.MemStats.DebugGC)
																r.Register("runtime.MemStats.EnableGC", runtimeMetrics.MemStats.EnableGC)
																r.Register("runtime.MemStats.Frees", runtimeMetrics.MemStats.Frees)
																r.Register("runtime.MemStats.HeapAlloc", runtimeMetrics.MemStats.HeapAlloc)
																r.Register("runtime.MemStats.HeapIdle", runtimeMetrics.MemStats.HeapIdle)
																r.Register("runtime.MemStats.HeapInuse", runtimeMetrics.MemStats.HeapInuse)
																r.Register("runtime.MemStats.HeapObjects", runtimeMetrics.MemStats.HeapObjects)
																r.Register("runtime.MemStats.HeapReleased", runtimeMetrics.MemStats.HeapReleased)
																r.Register("runtime.MemStats.HeapSys", runtimeMetrics.MemStats.HeapSys)
																r.Register("runtime.MemStats.LastGC", runtimeMetrics.MemStats.LastGC)
																r.Register("runtime.MemStats.Lookups", runtimeMetrics.MemStats.Lookups)
																r.Register("runtime.MemStats.Mallocs", runtimeMetrics.MemStats.Mallocs)
																r.Register("runtime.MemStats.MCacheInuse", runtimeMetrics.MemStats.MCacheInuse)
																r.Register("runtime.MemStats.MCacheSys", runtimeMetrics.MemStats.MCacheSys)
																r.Register("runtime.MemStats.MSpanInuse", runtimeMetrics.MemStats.MSpanInuse)
																r.Register("runtime.MemStats.MSpanSys", runtimeMetrics.MemStats.MSpanSys)
																r.Register("runtime.MemStats.NextGC", runtimeMetrics.MemStats.NextGC)
																r.Register("runtime.MemStats.NumGC", runtimeMetrics.MemStats.NumGC)
																r.Register("runtime.MemStats.GCCPUFraction", runtimeMetrics.MemStats.GCCPUFraction)
																r.Register("runtime.MemStats.PauseNs", runtimeMetrics.MemStats.PauseNs)
																r.Register("runtime.MemStats.PauseTotalNs", runtimeMetrics.MemStats.PauseTotalNs)
																r.Register("runtime.MemStats.StackInuse", runtimeMetrics.MemStats.StackInuse)
																r.Register("runtime.MemStats.StackSys", runtimeMetrics.MemStats.StackSys)
																r.Register("runtime.MemStats.Sys", runtimeMetrics.MemStats.Sys)
																r.Register("runtime.MemStats.TotalAlloc", runtimeMetrics.MemStats.TotalAlloc)
																r.Register("runtime.NumCgoCall", runtimeMetrics.NumCgoCall)
																r.Register("runtime.NumGoroutine", runtimeMetrics.NumGoroutine)
																r.Register("runtime.NumThread", runtimeMetrics.NumThread)
																r.Register("runtime.ReadMemStats", runtimeMetrics.ReadMemStats)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/runtime.go:214
		// _ = "end of CoverTab[96482]"
	})
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/runtime.go:215
	// _ = "end of CoverTab[96481]"
}

//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/runtime.go:216
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/runtime.go:216
var _ = _go_fuzz_dep_.CoverTab
