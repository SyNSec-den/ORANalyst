//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/graphite.go:1
package metrics

//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/graphite.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/graphite.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/graphite.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/graphite.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/graphite.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/graphite.go:1
)

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"time"
)

// GraphiteConfig provides a container with configuration parameters for
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/graphite.go:13
// the Graphite exporter
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/graphite.go:15
type GraphiteConfig struct {
	Addr		*net.TCPAddr	// Network address to connect to
	Registry	Registry	// Registry to be exported
	FlushInterval	time.Duration	// Flush interval
	DurationUnit	time.Duration	// Time conversion unit for durations
	Prefix		string		// Prefix to be prepended to metric names
	Percentiles	[]float64	// Percentiles to export from timers and histograms
}

// Graphite is a blocking exporter function which reports metrics in r
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/graphite.go:24
// to a graphite server located at addr, flushing them every d duration
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/graphite.go:24
// and prepending metric names with prefix.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/graphite.go:27
func Graphite(r Registry, d time.Duration, prefix string, addr *net.TCPAddr) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/graphite.go:27
	_go_fuzz_dep_.CoverTab[96200]++
															GraphiteWithConfig(GraphiteConfig{
		Addr:		addr,
		Registry:	r,
		FlushInterval:	d,
		DurationUnit:	time.Nanosecond,
		Prefix:		prefix,
		Percentiles:	[]float64{0.5, 0.75, 0.95, 0.99, 0.999},
	})
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/graphite.go:35
	// _ = "end of CoverTab[96200]"
}

// GraphiteWithConfig is a blocking exporter function just like Graphite,
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/graphite.go:38
// but it takes a GraphiteConfig instead.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/graphite.go:40
func GraphiteWithConfig(c GraphiteConfig) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/graphite.go:40
	_go_fuzz_dep_.CoverTab[96201]++
															log.Printf("WARNING: This go-metrics client has been DEPRECATED! It has been moved to https://github.com/cyberdelia/go-metrics-graphite and will be removed from rcrowley/go-metrics on August 12th 2015")
															for _ = range time.Tick(c.FlushInterval) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/graphite.go:42
		_go_fuzz_dep_.CoverTab[96202]++
																if err := graphite(&c); nil != err {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/graphite.go:43
			_go_fuzz_dep_.CoverTab[96203]++
																	log.Println(err)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/graphite.go:44
			// _ = "end of CoverTab[96203]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/graphite.go:45
			_go_fuzz_dep_.CoverTab[96204]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/graphite.go:45
			// _ = "end of CoverTab[96204]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/graphite.go:45
		}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/graphite.go:45
		// _ = "end of CoverTab[96202]"
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/graphite.go:46
	// _ = "end of CoverTab[96201]"
}

// GraphiteOnce performs a single submission to Graphite, returning a
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/graphite.go:49
// non-nil error on failed connections. This can be used in a loop
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/graphite.go:49
// similar to GraphiteWithConfig for custom error handling.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/graphite.go:52
func GraphiteOnce(c GraphiteConfig) error {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/graphite.go:52
	_go_fuzz_dep_.CoverTab[96205]++
															log.Printf("WARNING: This go-metrics client has been DEPRECATED! It has been moved to https://github.com/cyberdelia/go-metrics-graphite and will be removed from rcrowley/go-metrics on August 12th 2015")
															return graphite(&c)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/graphite.go:54
	// _ = "end of CoverTab[96205]"
}

func graphite(c *GraphiteConfig) error {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/graphite.go:57
	_go_fuzz_dep_.CoverTab[96206]++
															now := time.Now().Unix()
															du := float64(c.DurationUnit)
															conn, err := net.DialTCP("tcp", nil, c.Addr)
															if nil != err {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/graphite.go:61
		_go_fuzz_dep_.CoverTab[96209]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/graphite.go:62
		// _ = "end of CoverTab[96209]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/graphite.go:63
		_go_fuzz_dep_.CoverTab[96210]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/graphite.go:63
		// _ = "end of CoverTab[96210]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/graphite.go:63
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/graphite.go:63
	// _ = "end of CoverTab[96206]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/graphite.go:63
	_go_fuzz_dep_.CoverTab[96207]++
															defer conn.Close()
															w := bufio.NewWriter(conn)
															c.Registry.Each(func(name string, i interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/graphite.go:66
		_go_fuzz_dep_.CoverTab[96211]++
																switch metric := i.(type) {
		case Counter:
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/graphite.go:68
			_go_fuzz_dep_.CoverTab[96213]++
																	fmt.Fprintf(w, "%s.%s.count %d %d\n", c.Prefix, name, metric.Count(), now)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/graphite.go:69
			// _ = "end of CoverTab[96213]"
		case Gauge:
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/graphite.go:70
			_go_fuzz_dep_.CoverTab[96214]++
																	fmt.Fprintf(w, "%s.%s.value %d %d\n", c.Prefix, name, metric.Value(), now)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/graphite.go:71
			// _ = "end of CoverTab[96214]"
		case GaugeFloat64:
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/graphite.go:72
			_go_fuzz_dep_.CoverTab[96215]++
																	fmt.Fprintf(w, "%s.%s.value %f %d\n", c.Prefix, name, metric.Value(), now)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/graphite.go:73
			// _ = "end of CoverTab[96215]"
		case Histogram:
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/graphite.go:74
			_go_fuzz_dep_.CoverTab[96216]++
																	h := metric.Snapshot()
																	ps := h.Percentiles(c.Percentiles)
																	fmt.Fprintf(w, "%s.%s.count %d %d\n", c.Prefix, name, h.Count(), now)
																	fmt.Fprintf(w, "%s.%s.min %d %d\n", c.Prefix, name, h.Min(), now)
																	fmt.Fprintf(w, "%s.%s.max %d %d\n", c.Prefix, name, h.Max(), now)
																	fmt.Fprintf(w, "%s.%s.mean %.2f %d\n", c.Prefix, name, h.Mean(), now)
																	fmt.Fprintf(w, "%s.%s.std-dev %.2f %d\n", c.Prefix, name, h.StdDev(), now)
																	for psIdx, psKey := range c.Percentiles {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/graphite.go:82
				_go_fuzz_dep_.CoverTab[96220]++
																		key := strings.Replace(strconv.FormatFloat(psKey*100.0, 'f', -1, 64), ".", "", 1)
																		fmt.Fprintf(w, "%s.%s.%s-percentile %.2f %d\n", c.Prefix, name, key, ps[psIdx], now)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/graphite.go:84
				// _ = "end of CoverTab[96220]"
			}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/graphite.go:85
			// _ = "end of CoverTab[96216]"
		case Meter:
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/graphite.go:86
			_go_fuzz_dep_.CoverTab[96217]++
																	m := metric.Snapshot()
																	fmt.Fprintf(w, "%s.%s.count %d %d\n", c.Prefix, name, m.Count(), now)
																	fmt.Fprintf(w, "%s.%s.one-minute %.2f %d\n", c.Prefix, name, m.Rate1(), now)
																	fmt.Fprintf(w, "%s.%s.five-minute %.2f %d\n", c.Prefix, name, m.Rate5(), now)
																	fmt.Fprintf(w, "%s.%s.fifteen-minute %.2f %d\n", c.Prefix, name, m.Rate15(), now)
																	fmt.Fprintf(w, "%s.%s.mean %.2f %d\n", c.Prefix, name, m.RateMean(), now)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/graphite.go:92
			// _ = "end of CoverTab[96217]"
		case Timer:
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/graphite.go:93
			_go_fuzz_dep_.CoverTab[96218]++
																	t := metric.Snapshot()
																	ps := t.Percentiles(c.Percentiles)
																	fmt.Fprintf(w, "%s.%s.count %d %d\n", c.Prefix, name, t.Count(), now)
																	fmt.Fprintf(w, "%s.%s.min %d %d\n", c.Prefix, name, t.Min()/int64(du), now)
																	fmt.Fprintf(w, "%s.%s.max %d %d\n", c.Prefix, name, t.Max()/int64(du), now)
																	fmt.Fprintf(w, "%s.%s.mean %.2f %d\n", c.Prefix, name, t.Mean()/du, now)
																	fmt.Fprintf(w, "%s.%s.std-dev %.2f %d\n", c.Prefix, name, t.StdDev()/du, now)
																	for psIdx, psKey := range c.Percentiles {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/graphite.go:101
				_go_fuzz_dep_.CoverTab[96221]++
																		key := strings.Replace(strconv.FormatFloat(psKey*100.0, 'f', -1, 64), ".", "", 1)
																		fmt.Fprintf(w, "%s.%s.%s-percentile %.2f %d\n", c.Prefix, name, key, ps[psIdx], now)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/graphite.go:103
				// _ = "end of CoverTab[96221]"
			}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/graphite.go:104
			// _ = "end of CoverTab[96218]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/graphite.go:104
			_go_fuzz_dep_.CoverTab[96219]++
																	fmt.Fprintf(w, "%s.%s.one-minute %.2f %d\n", c.Prefix, name, t.Rate1(), now)
																	fmt.Fprintf(w, "%s.%s.five-minute %.2f %d\n", c.Prefix, name, t.Rate5(), now)
																	fmt.Fprintf(w, "%s.%s.fifteen-minute %.2f %d\n", c.Prefix, name, t.Rate15(), now)
																	fmt.Fprintf(w, "%s.%s.mean-rate %.2f %d\n", c.Prefix, name, t.RateMean(), now)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/graphite.go:108
			// _ = "end of CoverTab[96219]"
		}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/graphite.go:109
		// _ = "end of CoverTab[96211]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/graphite.go:109
		_go_fuzz_dep_.CoverTab[96212]++
																w.Flush()
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/graphite.go:110
		// _ = "end of CoverTab[96212]"
	})
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/graphite.go:111
	// _ = "end of CoverTab[96207]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/graphite.go:111
	_go_fuzz_dep_.CoverTab[96208]++
															return nil
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/graphite.go:112
	// _ = "end of CoverTab[96208]"
}

//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/graphite.go:113
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/graphite.go:113
var _ = _go_fuzz_dep_.CoverTab
