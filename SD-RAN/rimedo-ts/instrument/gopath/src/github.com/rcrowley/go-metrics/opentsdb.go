//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/opentsdb.go:1
package metrics

//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/opentsdb.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/opentsdb.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/opentsdb.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/opentsdb.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/opentsdb.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/opentsdb.go:1
)

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

var shortHostName string = ""

// OpenTSDBConfig provides a container with configuration parameters for
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/opentsdb.go:15
// the OpenTSDB exporter
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/opentsdb.go:17
type OpenTSDBConfig struct {
	Addr		*net.TCPAddr	// Network address to connect to
	Registry	Registry	// Registry to be exported
	FlushInterval	time.Duration	// Flush interval
	DurationUnit	time.Duration	// Time conversion unit for durations
	Prefix		string		// Prefix to be prepended to metric names
}

// OpenTSDB is a blocking exporter function which reports metrics in r
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/opentsdb.go:25
// to a TSDB server located at addr, flushing them every d duration
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/opentsdb.go:25
// and prepending metric names with prefix.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/opentsdb.go:28
func OpenTSDB(r Registry, d time.Duration, prefix string, addr *net.TCPAddr) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/opentsdb.go:28
	_go_fuzz_dep_.CoverTab[96360]++
															OpenTSDBWithConfig(OpenTSDBConfig{
		Addr:		addr,
		Registry:	r,
		FlushInterval:	d,
		DurationUnit:	time.Nanosecond,
		Prefix:		prefix,
	})
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/opentsdb.go:35
	// _ = "end of CoverTab[96360]"
}

// OpenTSDBWithConfig is a blocking exporter function just like OpenTSDB,
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/opentsdb.go:38
// but it takes a OpenTSDBConfig instead.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/opentsdb.go:40
func OpenTSDBWithConfig(c OpenTSDBConfig) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/opentsdb.go:40
	_go_fuzz_dep_.CoverTab[96361]++
															for _ = range time.Tick(c.FlushInterval) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/opentsdb.go:41
		_go_fuzz_dep_.CoverTab[96362]++
																if err := openTSDB(&c); nil != err {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/opentsdb.go:42
			_go_fuzz_dep_.CoverTab[96363]++
																	log.Println(err)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/opentsdb.go:43
			// _ = "end of CoverTab[96363]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/opentsdb.go:44
			_go_fuzz_dep_.CoverTab[96364]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/opentsdb.go:44
			// _ = "end of CoverTab[96364]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/opentsdb.go:44
		}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/opentsdb.go:44
		// _ = "end of CoverTab[96362]"
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/opentsdb.go:45
	// _ = "end of CoverTab[96361]"
}

func getShortHostname() string {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/opentsdb.go:48
	_go_fuzz_dep_.CoverTab[96365]++
															if shortHostName == "" {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/opentsdb.go:49
		_go_fuzz_dep_.CoverTab[96367]++
																host, _ := os.Hostname()
																if index := strings.Index(host, "."); index > 0 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/opentsdb.go:51
			_go_fuzz_dep_.CoverTab[96368]++
																	shortHostName = host[:index]
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/opentsdb.go:52
			// _ = "end of CoverTab[96368]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/opentsdb.go:53
			_go_fuzz_dep_.CoverTab[96369]++
																	shortHostName = host
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/opentsdb.go:54
			// _ = "end of CoverTab[96369]"
		}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/opentsdb.go:55
		// _ = "end of CoverTab[96367]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/opentsdb.go:56
		_go_fuzz_dep_.CoverTab[96370]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/opentsdb.go:56
		// _ = "end of CoverTab[96370]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/opentsdb.go:56
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/opentsdb.go:56
	// _ = "end of CoverTab[96365]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/opentsdb.go:56
	_go_fuzz_dep_.CoverTab[96366]++
															return shortHostName
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/opentsdb.go:57
	// _ = "end of CoverTab[96366]"
}

func openTSDB(c *OpenTSDBConfig) error {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/opentsdb.go:60
	_go_fuzz_dep_.CoverTab[96371]++
															shortHostname := getShortHostname()
															now := time.Now().Unix()
															du := float64(c.DurationUnit)
															conn, err := net.DialTCP("tcp", nil, c.Addr)
															if nil != err {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/opentsdb.go:65
		_go_fuzz_dep_.CoverTab[96374]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/opentsdb.go:66
		// _ = "end of CoverTab[96374]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/opentsdb.go:67
		_go_fuzz_dep_.CoverTab[96375]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/opentsdb.go:67
		// _ = "end of CoverTab[96375]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/opentsdb.go:67
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/opentsdb.go:67
	// _ = "end of CoverTab[96371]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/opentsdb.go:67
	_go_fuzz_dep_.CoverTab[96372]++
															defer conn.Close()
															w := bufio.NewWriter(conn)
															c.Registry.Each(func(name string, i interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/opentsdb.go:70
		_go_fuzz_dep_.CoverTab[96376]++
																switch metric := i.(type) {
		case Counter:
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/opentsdb.go:72
			_go_fuzz_dep_.CoverTab[96378]++
																	fmt.Fprintf(w, "put %s.%s.count %d %d host=%s\n", c.Prefix, name, now, metric.Count(), shortHostname)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/opentsdb.go:73
			// _ = "end of CoverTab[96378]"
		case Gauge:
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/opentsdb.go:74
			_go_fuzz_dep_.CoverTab[96379]++
																	fmt.Fprintf(w, "put %s.%s.value %d %d host=%s\n", c.Prefix, name, now, metric.Value(), shortHostname)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/opentsdb.go:75
			// _ = "end of CoverTab[96379]"
		case GaugeFloat64:
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/opentsdb.go:76
			_go_fuzz_dep_.CoverTab[96380]++
																	fmt.Fprintf(w, "put %s.%s.value %d %f host=%s\n", c.Prefix, name, now, metric.Value(), shortHostname)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/opentsdb.go:77
			// _ = "end of CoverTab[96380]"
		case Histogram:
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/opentsdb.go:78
			_go_fuzz_dep_.CoverTab[96381]++
																	h := metric.Snapshot()
																	ps := h.Percentiles([]float64{0.5, 0.75, 0.95, 0.99, 0.999})
																	fmt.Fprintf(w, "put %s.%s.count %d %d host=%s\n", c.Prefix, name, now, h.Count(), shortHostname)
																	fmt.Fprintf(w, "put %s.%s.min %d %d host=%s\n", c.Prefix, name, now, h.Min(), shortHostname)
																	fmt.Fprintf(w, "put %s.%s.max %d %d host=%s\n", c.Prefix, name, now, h.Max(), shortHostname)
																	fmt.Fprintf(w, "put %s.%s.mean %d %.2f host=%s\n", c.Prefix, name, now, h.Mean(), shortHostname)
																	fmt.Fprintf(w, "put %s.%s.std-dev %d %.2f host=%s\n", c.Prefix, name, now, h.StdDev(), shortHostname)
																	fmt.Fprintf(w, "put %s.%s.50-percentile %d %.2f host=%s\n", c.Prefix, name, now, ps[0], shortHostname)
																	fmt.Fprintf(w, "put %s.%s.75-percentile %d %.2f host=%s\n", c.Prefix, name, now, ps[1], shortHostname)
																	fmt.Fprintf(w, "put %s.%s.95-percentile %d %.2f host=%s\n", c.Prefix, name, now, ps[2], shortHostname)
																	fmt.Fprintf(w, "put %s.%s.99-percentile %d %.2f host=%s\n", c.Prefix, name, now, ps[3], shortHostname)
																	fmt.Fprintf(w, "put %s.%s.999-percentile %d %.2f host=%s\n", c.Prefix, name, now, ps[4], shortHostname)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/opentsdb.go:90
			// _ = "end of CoverTab[96381]"
		case Meter:
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/opentsdb.go:91
			_go_fuzz_dep_.CoverTab[96382]++
																	m := metric.Snapshot()
																	fmt.Fprintf(w, "put %s.%s.count %d %d host=%s\n", c.Prefix, name, now, m.Count(), shortHostname)
																	fmt.Fprintf(w, "put %s.%s.one-minute %d %.2f host=%s\n", c.Prefix, name, now, m.Rate1(), shortHostname)
																	fmt.Fprintf(w, "put %s.%s.five-minute %d %.2f host=%s\n", c.Prefix, name, now, m.Rate5(), shortHostname)
																	fmt.Fprintf(w, "put %s.%s.fifteen-minute %d %.2f host=%s\n", c.Prefix, name, now, m.Rate15(), shortHostname)
																	fmt.Fprintf(w, "put %s.%s.mean %d %.2f host=%s\n", c.Prefix, name, now, m.RateMean(), shortHostname)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/opentsdb.go:97
			// _ = "end of CoverTab[96382]"
		case Timer:
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/opentsdb.go:98
			_go_fuzz_dep_.CoverTab[96383]++
																	t := metric.Snapshot()
																	ps := t.Percentiles([]float64{0.5, 0.75, 0.95, 0.99, 0.999})
																	fmt.Fprintf(w, "put %s.%s.count %d %d host=%s\n", c.Prefix, name, now, t.Count(), shortHostname)
																	fmt.Fprintf(w, "put %s.%s.min %d %d host=%s\n", c.Prefix, name, now, t.Min()/int64(du), shortHostname)
																	fmt.Fprintf(w, "put %s.%s.max %d %d host=%s\n", c.Prefix, name, now, t.Max()/int64(du), shortHostname)
																	fmt.Fprintf(w, "put %s.%s.mean %d %.2f host=%s\n", c.Prefix, name, now, t.Mean()/du, shortHostname)
																	fmt.Fprintf(w, "put %s.%s.std-dev %d %.2f host=%s\n", c.Prefix, name, now, t.StdDev()/du, shortHostname)
																	fmt.Fprintf(w, "put %s.%s.50-percentile %d %.2f host=%s\n", c.Prefix, name, now, ps[0]/du, shortHostname)
																	fmt.Fprintf(w, "put %s.%s.75-percentile %d %.2f host=%s\n", c.Prefix, name, now, ps[1]/du, shortHostname)
																	fmt.Fprintf(w, "put %s.%s.95-percentile %d %.2f host=%s\n", c.Prefix, name, now, ps[2]/du, shortHostname)
																	fmt.Fprintf(w, "put %s.%s.99-percentile %d %.2f host=%s\n", c.Prefix, name, now, ps[3]/du, shortHostname)
																	fmt.Fprintf(w, "put %s.%s.999-percentile %d %.2f host=%s\n", c.Prefix, name, now, ps[4]/du, shortHostname)
																	fmt.Fprintf(w, "put %s.%s.one-minute %d %.2f host=%s\n", c.Prefix, name, now, t.Rate1(), shortHostname)
																	fmt.Fprintf(w, "put %s.%s.five-minute %d %.2f host=%s\n", c.Prefix, name, now, t.Rate5(), shortHostname)
																	fmt.Fprintf(w, "put %s.%s.fifteen-minute %d %.2f host=%s\n", c.Prefix, name, now, t.Rate15(), shortHostname)
																	fmt.Fprintf(w, "put %s.%s.mean-rate %d %.2f host=%s\n", c.Prefix, name, now, t.RateMean(), shortHostname)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/opentsdb.go:114
			// _ = "end of CoverTab[96383]"
		}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/opentsdb.go:115
		// _ = "end of CoverTab[96376]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/opentsdb.go:115
		_go_fuzz_dep_.CoverTab[96377]++
																w.Flush()
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/opentsdb.go:116
		// _ = "end of CoverTab[96377]"
	})
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/opentsdb.go:117
	// _ = "end of CoverTab[96372]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/opentsdb.go:117
	_go_fuzz_dep_.CoverTab[96373]++
															return nil
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/opentsdb.go:118
	// _ = "end of CoverTab[96373]"
}

//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/opentsdb.go:119
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/opentsdb.go:119
var _ = _go_fuzz_dep_.CoverTab
