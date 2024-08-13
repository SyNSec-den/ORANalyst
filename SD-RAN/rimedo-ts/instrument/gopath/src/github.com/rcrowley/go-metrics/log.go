//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/log.go:1
package metrics

//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/log.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/log.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/log.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/log.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/log.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/log.go:1
)

import (
	"time"
)

type Logger interface {
	Printf(format string, v ...interface{})
}

// Log outputs each metric in the given registry periodically using the given logger.
func Log(r Registry, freq time.Duration, l Logger) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/log.go:12
	_go_fuzz_dep_.CoverTab[96291]++
														LogScaled(r, freq, time.Nanosecond, l)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/log.go:13
	// _ = "end of CoverTab[96291]"
}

// LogOnCue outputs each metric in the given registry on demand through the channel
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/log.go:16
// using the given logger
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/log.go:18
func LogOnCue(r Registry, ch chan interface{}, l Logger) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/log.go:18
	_go_fuzz_dep_.CoverTab[96292]++
														LogScaledOnCue(r, ch, time.Nanosecond, l)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/log.go:19
	// _ = "end of CoverTab[96292]"
}

// LogScaled outputs each metric in the given registry periodically using the given
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/log.go:22
// logger. Print timings in `scale` units (eg time.Millisecond) rather than nanos.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/log.go:24
func LogScaled(r Registry, freq time.Duration, scale time.Duration, l Logger) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/log.go:24
	_go_fuzz_dep_.CoverTab[96293]++
														ch := make(chan interface{})
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/log.go:25
	_curRoutineNum110_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/log.go:25
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum110_)
														go func(channel chan interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/log.go:26
		_go_fuzz_dep_.CoverTab[96295]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/log.go:26
		defer func() {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/log.go:26
			_go_fuzz_dep_.CoverTab[96296]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/log.go:26
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum110_)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/log.go:26
			// _ = "end of CoverTab[96296]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/log.go:26
		}()
															for _ = range time.Tick(freq) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/log.go:27
			_go_fuzz_dep_.CoverTab[96297]++
																channel <- struct{}{}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/log.go:28
			// _ = "end of CoverTab[96297]"
		}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/log.go:29
		// _ = "end of CoverTab[96295]"
	}(ch)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/log.go:30
	// _ = "end of CoverTab[96293]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/log.go:30
	_go_fuzz_dep_.CoverTab[96294]++
														LogScaledOnCue(r, ch, scale, l)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/log.go:31
	// _ = "end of CoverTab[96294]"
}

// LogScaledOnCue outputs each metric in the given registry on demand through the channel
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/log.go:34
// using the given logger. Print timings in `scale` units (eg time.Millisecond) rather
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/log.go:34
// than nanos.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/log.go:37
func LogScaledOnCue(r Registry, ch chan interface{}, scale time.Duration, l Logger) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/log.go:37
	_go_fuzz_dep_.CoverTab[96298]++
														du := float64(scale)
														duSuffix := scale.String()[1:]

														for _ = range ch {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/log.go:41
		_go_fuzz_dep_.CoverTab[96299]++
															r.Each(func(name string, i interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/log.go:42
			_go_fuzz_dep_.CoverTab[96300]++
																switch metric := i.(type) {
			case Counter:
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/log.go:44
				_go_fuzz_dep_.CoverTab[96301]++
																	l.Printf("counter %s\n", name)
																	l.Printf("  count:       %9d\n", metric.Count())
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/log.go:46
				// _ = "end of CoverTab[96301]"
			case Gauge:
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/log.go:47
				_go_fuzz_dep_.CoverTab[96302]++
																	l.Printf("gauge %s\n", name)
																	l.Printf("  value:       %9d\n", metric.Value())
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/log.go:49
				// _ = "end of CoverTab[96302]"
			case GaugeFloat64:
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/log.go:50
				_go_fuzz_dep_.CoverTab[96303]++
																	l.Printf("gauge %s\n", name)
																	l.Printf("  value:       %f\n", metric.Value())
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/log.go:52
				// _ = "end of CoverTab[96303]"
			case Healthcheck:
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/log.go:53
				_go_fuzz_dep_.CoverTab[96304]++
																	metric.Check()
																	l.Printf("healthcheck %s\n", name)
																	l.Printf("  error:       %v\n", metric.Error())
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/log.go:56
				// _ = "end of CoverTab[96304]"
			case Histogram:
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/log.go:57
				_go_fuzz_dep_.CoverTab[96305]++
																	h := metric.Snapshot()
																	ps := h.Percentiles([]float64{0.5, 0.75, 0.95, 0.99, 0.999})
																	l.Printf("histogram %s\n", name)
																	l.Printf("  count:       %9d\n", h.Count())
																	l.Printf("  min:         %9d\n", h.Min())
																	l.Printf("  max:         %9d\n", h.Max())
																	l.Printf("  mean:        %12.2f\n", h.Mean())
																	l.Printf("  stddev:      %12.2f\n", h.StdDev())
																	l.Printf("  median:      %12.2f\n", ps[0])
																	l.Printf("  75%%:         %12.2f\n", ps[1])
																	l.Printf("  95%%:         %12.2f\n", ps[2])
																	l.Printf("  99%%:         %12.2f\n", ps[3])
																	l.Printf("  99.9%%:       %12.2f\n", ps[4])
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/log.go:70
				// _ = "end of CoverTab[96305]"
			case Meter:
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/log.go:71
				_go_fuzz_dep_.CoverTab[96306]++
																	m := metric.Snapshot()
																	l.Printf("meter %s\n", name)
																	l.Printf("  count:       %9d\n", m.Count())
																	l.Printf("  1-min rate:  %12.2f\n", m.Rate1())
																	l.Printf("  5-min rate:  %12.2f\n", m.Rate5())
																	l.Printf("  15-min rate: %12.2f\n", m.Rate15())
																	l.Printf("  mean rate:   %12.2f\n", m.RateMean())
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/log.go:78
				// _ = "end of CoverTab[96306]"
			case Timer:
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/log.go:79
				_go_fuzz_dep_.CoverTab[96307]++
																	t := metric.Snapshot()
																	ps := t.Percentiles([]float64{0.5, 0.75, 0.95, 0.99, 0.999})
																	l.Printf("timer %s\n", name)
																	l.Printf("  count:       %9d\n", t.Count())
																	l.Printf("  min:         %12.2f%s\n", float64(t.Min())/du, duSuffix)
																	l.Printf("  max:         %12.2f%s\n", float64(t.Max())/du, duSuffix)
																	l.Printf("  mean:        %12.2f%s\n", t.Mean()/du, duSuffix)
																	l.Printf("  stddev:      %12.2f%s\n", t.StdDev()/du, duSuffix)
																	l.Printf("  median:      %12.2f%s\n", ps[0]/du, duSuffix)
																	l.Printf("  75%%:         %12.2f%s\n", ps[1]/du, duSuffix)
																	l.Printf("  95%%:         %12.2f%s\n", ps[2]/du, duSuffix)
																	l.Printf("  99%%:         %12.2f%s\n", ps[3]/du, duSuffix)
																	l.Printf("  99.9%%:       %12.2f%s\n", ps[4]/du, duSuffix)
																	l.Printf("  1-min rate:  %12.2f\n", t.Rate1())
																	l.Printf("  5-min rate:  %12.2f\n", t.Rate5())
																	l.Printf("  15-min rate: %12.2f\n", t.Rate15())
																	l.Printf("  mean rate:   %12.2f\n", t.RateMean())
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/log.go:96
				// _ = "end of CoverTab[96307]"
			}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/log.go:97
			// _ = "end of CoverTab[96300]"
		})
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/log.go:98
		// _ = "end of CoverTab[96299]"
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/log.go:99
	// _ = "end of CoverTab[96298]"
}

//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/log.go:100
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/log.go:100
var _ = _go_fuzz_dep_.CoverTab
