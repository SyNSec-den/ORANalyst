// +build !windows

//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/syslog.go:3
package metrics

//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/syslog.go:3
import (
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/syslog.go:3
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/syslog.go:3
)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/syslog.go:3
import (
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/syslog.go:3
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/syslog.go:3
)

import (
	"fmt"
	"log/syslog"
	"time"
)

// Output each metric in the given registry to syslog periodically using
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/syslog.go:11
// the given syslogger.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/syslog.go:13
func Syslog(r Registry, d time.Duration, w *syslog.Writer) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/syslog.go:13
	_go_fuzz_dep_.CoverTab[96633]++
															for _ = range time.Tick(d) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/syslog.go:14
		_go_fuzz_dep_.CoverTab[96634]++
																r.Each(func(name string, i interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/syslog.go:15
			_go_fuzz_dep_.CoverTab[96635]++
																	switch metric := i.(type) {
			case Counter:
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/syslog.go:17
				_go_fuzz_dep_.CoverTab[96636]++
																		w.Info(fmt.Sprintf("counter %s: count: %d", name, metric.Count()))
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/syslog.go:18
				// _ = "end of CoverTab[96636]"
			case Gauge:
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/syslog.go:19
				_go_fuzz_dep_.CoverTab[96637]++
																		w.Info(fmt.Sprintf("gauge %s: value: %d", name, metric.Value()))
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/syslog.go:20
				// _ = "end of CoverTab[96637]"
			case GaugeFloat64:
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/syslog.go:21
				_go_fuzz_dep_.CoverTab[96638]++
																		w.Info(fmt.Sprintf("gauge %s: value: %f", name, metric.Value()))
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/syslog.go:22
				// _ = "end of CoverTab[96638]"
			case Healthcheck:
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/syslog.go:23
				_go_fuzz_dep_.CoverTab[96639]++
																		metric.Check()
																		w.Info(fmt.Sprintf("healthcheck %s: error: %v", name, metric.Error()))
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/syslog.go:25
				// _ = "end of CoverTab[96639]"
			case Histogram:
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/syslog.go:26
				_go_fuzz_dep_.CoverTab[96640]++
																		h := metric.Snapshot()
																		ps := h.Percentiles([]float64{0.5, 0.75, 0.95, 0.99, 0.999})
																		w.Info(fmt.Sprintf(
					"histogram %s: count: %d min: %d max: %d mean: %.2f stddev: %.2f median: %.2f 75%%: %.2f 95%%: %.2f 99%%: %.2f 99.9%%: %.2f",
					name,
					h.Count(),
					h.Min(),
					h.Max(),
					h.Mean(),
					h.StdDev(),
					ps[0],
					ps[1],
					ps[2],
					ps[3],
					ps[4],
				))
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/syslog.go:42
				// _ = "end of CoverTab[96640]"
			case Meter:
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/syslog.go:43
				_go_fuzz_dep_.CoverTab[96641]++
																		m := metric.Snapshot()
																		w.Info(fmt.Sprintf(
					"meter %s: count: %d 1-min: %.2f 5-min: %.2f 15-min: %.2f mean: %.2f",
					name,
					m.Count(),
					m.Rate1(),
					m.Rate5(),
					m.Rate15(),
					m.RateMean(),
				))
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/syslog.go:53
				// _ = "end of CoverTab[96641]"
			case Timer:
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/syslog.go:54
				_go_fuzz_dep_.CoverTab[96642]++
																		t := metric.Snapshot()
																		ps := t.Percentiles([]float64{0.5, 0.75, 0.95, 0.99, 0.999})
																		w.Info(fmt.Sprintf(
					"timer %s: count: %d min: %d max: %d mean: %.2f stddev: %.2f median: %.2f 75%%: %.2f 95%%: %.2f 99%%: %.2f 99.9%%: %.2f 1-min: %.2f 5-min: %.2f 15-min: %.2f mean-rate: %.2f",
					name,
					t.Count(),
					t.Min(),
					t.Max(),
					t.Mean(),
					t.StdDev(),
					ps[0],
					ps[1],
					ps[2],
					ps[3],
					ps[4],
					t.Rate1(),
					t.Rate5(),
					t.Rate15(),
					t.RateMean(),
				))
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/syslog.go:74
				// _ = "end of CoverTab[96642]"
			}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/syslog.go:75
			// _ = "end of CoverTab[96635]"
		})
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/syslog.go:76
		// _ = "end of CoverTab[96634]"
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/syslog.go:77
	// _ = "end of CoverTab[96633]"
}

//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/syslog.go:78
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/syslog.go:78
var _ = _go_fuzz_dep_.CoverTab
