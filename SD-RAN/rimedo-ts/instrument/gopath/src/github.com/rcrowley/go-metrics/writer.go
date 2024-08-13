//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/writer.go:1
package metrics

//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/writer.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/writer.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/writer.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/writer.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/writer.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/writer.go:1
)

import (
	"fmt"
	"io"
	"sort"
	"time"
)

// Write sorts writes each metric in the given registry periodically to the
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/writer.go:10
// given io.Writer.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/writer.go:12
func Write(r Registry, d time.Duration, w io.Writer) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/writer.go:12
	_go_fuzz_dep_.CoverTab[96713]++
															for _ = range time.Tick(d) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/writer.go:13
		_go_fuzz_dep_.CoverTab[96714]++
																WriteOnce(r, w)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/writer.go:14
		// _ = "end of CoverTab[96714]"
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/writer.go:15
	// _ = "end of CoverTab[96713]"
}

// WriteOnce sorts and writes metrics in the given registry to the given
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/writer.go:18
// io.Writer.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/writer.go:20
func WriteOnce(r Registry, w io.Writer) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/writer.go:20
	_go_fuzz_dep_.CoverTab[96715]++
															var namedMetrics namedMetricSlice
															r.Each(func(name string, i interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/writer.go:22
		_go_fuzz_dep_.CoverTab[96717]++
																namedMetrics = append(namedMetrics, namedMetric{name, i})
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/writer.go:23
		// _ = "end of CoverTab[96717]"
	})
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/writer.go:24
	// _ = "end of CoverTab[96715]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/writer.go:24
	_go_fuzz_dep_.CoverTab[96716]++

															sort.Sort(namedMetrics)
															for _, namedMetric := range namedMetrics {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/writer.go:27
		_go_fuzz_dep_.CoverTab[96718]++
																switch metric := namedMetric.m.(type) {
		case Counter:
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/writer.go:29
			_go_fuzz_dep_.CoverTab[96719]++
																	fmt.Fprintf(w, "counter %s\n", namedMetric.name)
																	fmt.Fprintf(w, "  count:       %9d\n", metric.Count())
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/writer.go:31
			// _ = "end of CoverTab[96719]"
		case Gauge:
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/writer.go:32
			_go_fuzz_dep_.CoverTab[96720]++
																	fmt.Fprintf(w, "gauge %s\n", namedMetric.name)
																	fmt.Fprintf(w, "  value:       %9d\n", metric.Value())
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/writer.go:34
			// _ = "end of CoverTab[96720]"
		case GaugeFloat64:
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/writer.go:35
			_go_fuzz_dep_.CoverTab[96721]++
																	fmt.Fprintf(w, "gauge %s\n", namedMetric.name)
																	fmt.Fprintf(w, "  value:       %f\n", metric.Value())
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/writer.go:37
			// _ = "end of CoverTab[96721]"
		case Healthcheck:
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/writer.go:38
			_go_fuzz_dep_.CoverTab[96722]++
																	metric.Check()
																	fmt.Fprintf(w, "healthcheck %s\n", namedMetric.name)
																	fmt.Fprintf(w, "  error:       %v\n", metric.Error())
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/writer.go:41
			// _ = "end of CoverTab[96722]"
		case Histogram:
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/writer.go:42
			_go_fuzz_dep_.CoverTab[96723]++
																	h := metric.Snapshot()
																	ps := h.Percentiles([]float64{0.5, 0.75, 0.95, 0.99, 0.999})
																	fmt.Fprintf(w, "histogram %s\n", namedMetric.name)
																	fmt.Fprintf(w, "  count:       %9d\n", h.Count())
																	fmt.Fprintf(w, "  min:         %9d\n", h.Min())
																	fmt.Fprintf(w, "  max:         %9d\n", h.Max())
																	fmt.Fprintf(w, "  mean:        %12.2f\n", h.Mean())
																	fmt.Fprintf(w, "  stddev:      %12.2f\n", h.StdDev())
																	fmt.Fprintf(w, "  median:      %12.2f\n", ps[0])
																	fmt.Fprintf(w, "  75%%:         %12.2f\n", ps[1])
																	fmt.Fprintf(w, "  95%%:         %12.2f\n", ps[2])
																	fmt.Fprintf(w, "  99%%:         %12.2f\n", ps[3])
																	fmt.Fprintf(w, "  99.9%%:       %12.2f\n", ps[4])
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/writer.go:55
			// _ = "end of CoverTab[96723]"
		case Meter:
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/writer.go:56
			_go_fuzz_dep_.CoverTab[96724]++
																	m := metric.Snapshot()
																	fmt.Fprintf(w, "meter %s\n", namedMetric.name)
																	fmt.Fprintf(w, "  count:       %9d\n", m.Count())
																	fmt.Fprintf(w, "  1-min rate:  %12.2f\n", m.Rate1())
																	fmt.Fprintf(w, "  5-min rate:  %12.2f\n", m.Rate5())
																	fmt.Fprintf(w, "  15-min rate: %12.2f\n", m.Rate15())
																	fmt.Fprintf(w, "  mean rate:   %12.2f\n", m.RateMean())
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/writer.go:63
			// _ = "end of CoverTab[96724]"
		case Timer:
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/writer.go:64
			_go_fuzz_dep_.CoverTab[96725]++
																	t := metric.Snapshot()
																	ps := t.Percentiles([]float64{0.5, 0.75, 0.95, 0.99, 0.999})
																	fmt.Fprintf(w, "timer %s\n", namedMetric.name)
																	fmt.Fprintf(w, "  count:       %9d\n", t.Count())
																	fmt.Fprintf(w, "  min:         %9d\n", t.Min())
																	fmt.Fprintf(w, "  max:         %9d\n", t.Max())
																	fmt.Fprintf(w, "  mean:        %12.2f\n", t.Mean())
																	fmt.Fprintf(w, "  stddev:      %12.2f\n", t.StdDev())
																	fmt.Fprintf(w, "  median:      %12.2f\n", ps[0])
																	fmt.Fprintf(w, "  75%%:         %12.2f\n", ps[1])
																	fmt.Fprintf(w, "  95%%:         %12.2f\n", ps[2])
																	fmt.Fprintf(w, "  99%%:         %12.2f\n", ps[3])
																	fmt.Fprintf(w, "  99.9%%:       %12.2f\n", ps[4])
																	fmt.Fprintf(w, "  1-min rate:  %12.2f\n", t.Rate1())
																	fmt.Fprintf(w, "  5-min rate:  %12.2f\n", t.Rate5())
																	fmt.Fprintf(w, "  15-min rate: %12.2f\n", t.Rate15())
																	fmt.Fprintf(w, "  mean rate:   %12.2f\n", t.RateMean())
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/writer.go:81
			// _ = "end of CoverTab[96725]"
		}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/writer.go:82
		// _ = "end of CoverTab[96718]"
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/writer.go:83
	// _ = "end of CoverTab[96716]"
}

type namedMetric struct {
	name	string
	m	interface{}
}

// namedMetricSlice is a slice of namedMetrics that implements sort.Interface.
type namedMetricSlice []namedMetric

func (nms namedMetricSlice) Len() int {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/writer.go:94
	_go_fuzz_dep_.CoverTab[96726]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/writer.go:94
	return len(nms)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/writer.go:94
	// _ = "end of CoverTab[96726]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/writer.go:94
}

func (nms namedMetricSlice) Swap(i, j int) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/writer.go:96
	_go_fuzz_dep_.CoverTab[96727]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/writer.go:96
	nms[i], nms[j] = nms[j], nms[i]
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/writer.go:96
	// _ = "end of CoverTab[96727]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/writer.go:96
}

func (nms namedMetricSlice) Less(i, j int) bool {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/writer.go:98
	_go_fuzz_dep_.CoverTab[96728]++
															return nms[i].name < nms[j].name
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/writer.go:99
	// _ = "end of CoverTab[96728]"
}

//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/writer.go:100
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/writer.go:100
var _ = _go_fuzz_dep_.CoverTab
