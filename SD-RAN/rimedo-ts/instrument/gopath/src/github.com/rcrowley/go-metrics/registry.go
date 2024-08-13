//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:1
package metrics

//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:1
)

import (
	"fmt"
	"reflect"
	"strings"
	"sync"
)

// DuplicateMetric is the error returned by Registry.Register when a metric
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:10
// already exists.  If you mean to Register that metric you must first
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:10
// Unregister the existing metric.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:13
type DuplicateMetric string

func (err DuplicateMetric) Error() string {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:15
	_go_fuzz_dep_.CoverTab[96384]++
															return fmt.Sprintf("duplicate metric: %s", string(err))
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:16
	// _ = "end of CoverTab[96384]"
}

// A Registry holds references to a set of metrics by name and can iterate
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:19
// over them, calling callback functions provided by the user.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:19
//
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:19
// This is an interface so as to encourage other structs to implement
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:19
// the Registry API as appropriate.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:24
type Registry interface {

	// Call the given function for each registered metric.
	Each(func(string, interface{}))

	// Get the metric by the given name or nil if none is registered.
	Get(string) interface{}

	// GetAll metrics in the Registry.
	GetAll() map[string]map[string]interface{}

	// Gets an existing metric or registers the given one.
	// The interface can be the metric to register if not found in registry,
	// or a function returning the metric for lazy instantiation.
	GetOrRegister(string, interface{}) interface{}

	// Register the given metric under the given name.
	Register(string, interface{}) error

	// Run all registered healthchecks.
	RunHealthchecks()

	// Unregister the metric with the given name.
	Unregister(string)

	// Unregister all metrics.  (Mostly for testing.)
	UnregisterAll()
}

// The standard implementation of a Registry is a mutex-protected map
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:53
// of names to metrics.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:55
type StandardRegistry struct {
	metrics	map[string]interface{}
	mutex	sync.RWMutex
}

// Create a new registry.
func NewRegistry() Registry {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:61
	_go_fuzz_dep_.CoverTab[96385]++
															return &StandardRegistry{metrics: make(map[string]interface{})}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:62
	// _ = "end of CoverTab[96385]"
}

// Call the given function for each registered metric.
func (r *StandardRegistry) Each(f func(string, interface{})) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:66
	_go_fuzz_dep_.CoverTab[96386]++
															metrics := r.registered()
															for i := range metrics {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:68
		_go_fuzz_dep_.CoverTab[96387]++
																kv := &metrics[i]
																f(kv.name, kv.value)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:70
		// _ = "end of CoverTab[96387]"
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:71
	// _ = "end of CoverTab[96386]"
}

// Get the metric by the given name or nil if none is registered.
func (r *StandardRegistry) Get(name string) interface{} {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:75
	_go_fuzz_dep_.CoverTab[96388]++
															r.mutex.RLock()
															defer r.mutex.RUnlock()
															return r.metrics[name]
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:78
	// _ = "end of CoverTab[96388]"
}

// Gets an existing metric or creates and registers a new one. Threadsafe
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:81
// alternative to calling Get and Register on failure.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:81
// The interface can be the metric to register if not found in registry,
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:81
// or a function returning the metric for lazy instantiation.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:85
func (r *StandardRegistry) GetOrRegister(name string, i interface{}) interface{} {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:85
	_go_fuzz_dep_.CoverTab[96389]++

															r.mutex.RLock()
															metric, ok := r.metrics[name]
															r.mutex.RUnlock()
															if ok {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:90
		_go_fuzz_dep_.CoverTab[96393]++
																return metric
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:91
		// _ = "end of CoverTab[96393]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:92
		_go_fuzz_dep_.CoverTab[96394]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:92
		// _ = "end of CoverTab[96394]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:92
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:92
	// _ = "end of CoverTab[96389]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:92
	_go_fuzz_dep_.CoverTab[96390]++

//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:95
	r.mutex.Lock()
	defer r.mutex.Unlock()
	if metric, ok := r.metrics[name]; ok {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:97
		_go_fuzz_dep_.CoverTab[96395]++
																return metric
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:98
		// _ = "end of CoverTab[96395]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:99
		_go_fuzz_dep_.CoverTab[96396]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:99
		// _ = "end of CoverTab[96396]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:99
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:99
	// _ = "end of CoverTab[96390]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:99
	_go_fuzz_dep_.CoverTab[96391]++
															if v := reflect.ValueOf(i); v.Kind() == reflect.Func {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:100
		_go_fuzz_dep_.CoverTab[96397]++
																i = v.Call(nil)[0].Interface()
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:101
		// _ = "end of CoverTab[96397]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:102
		_go_fuzz_dep_.CoverTab[96398]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:102
		// _ = "end of CoverTab[96398]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:102
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:102
	// _ = "end of CoverTab[96391]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:102
	_go_fuzz_dep_.CoverTab[96392]++
															r.register(name, i)
															return i
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:104
	// _ = "end of CoverTab[96392]"
}

// Register the given metric under the given name.  Returns a DuplicateMetric
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:107
// if a metric by the given name is already registered.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:109
func (r *StandardRegistry) Register(name string, i interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:109
	_go_fuzz_dep_.CoverTab[96399]++
															r.mutex.Lock()
															defer r.mutex.Unlock()
															return r.register(name, i)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:112
	// _ = "end of CoverTab[96399]"
}

// Run all registered healthchecks.
func (r *StandardRegistry) RunHealthchecks() {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:116
	_go_fuzz_dep_.CoverTab[96400]++
															r.mutex.RLock()
															defer r.mutex.RUnlock()
															for _, i := range r.metrics {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:119
		_go_fuzz_dep_.CoverTab[96401]++
																if h, ok := i.(Healthcheck); ok {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:120
			_go_fuzz_dep_.CoverTab[96402]++
																	h.Check()
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:121
			// _ = "end of CoverTab[96402]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:122
			_go_fuzz_dep_.CoverTab[96403]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:122
			// _ = "end of CoverTab[96403]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:122
		}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:122
		// _ = "end of CoverTab[96401]"
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:123
	// _ = "end of CoverTab[96400]"
}

// GetAll metrics in the Registry
func (r *StandardRegistry) GetAll() map[string]map[string]interface{} {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:127
	_go_fuzz_dep_.CoverTab[96404]++
															data := make(map[string]map[string]interface{})
															r.Each(func(name string, i interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:129
		_go_fuzz_dep_.CoverTab[96406]++
																values := make(map[string]interface{})
																switch metric := i.(type) {
		case Counter:
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:132
			_go_fuzz_dep_.CoverTab[96408]++
																	values["count"] = metric.Count()
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:133
			// _ = "end of CoverTab[96408]"
		case Gauge:
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:134
			_go_fuzz_dep_.CoverTab[96409]++
																	values["value"] = metric.Value()
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:135
			// _ = "end of CoverTab[96409]"
		case GaugeFloat64:
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:136
			_go_fuzz_dep_.CoverTab[96410]++
																	values["value"] = metric.Value()
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:137
			// _ = "end of CoverTab[96410]"
		case Healthcheck:
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:138
			_go_fuzz_dep_.CoverTab[96411]++
																	values["error"] = nil
																	metric.Check()
																	if err := metric.Error(); nil != err {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:141
				_go_fuzz_dep_.CoverTab[96415]++
																		values["error"] = metric.Error().Error()
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:142
				// _ = "end of CoverTab[96415]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:143
				_go_fuzz_dep_.CoverTab[96416]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:143
				// _ = "end of CoverTab[96416]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:143
			}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:143
			// _ = "end of CoverTab[96411]"
		case Histogram:
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:144
			_go_fuzz_dep_.CoverTab[96412]++
																	h := metric.Snapshot()
																	ps := h.Percentiles([]float64{0.5, 0.75, 0.95, 0.99, 0.999})
																	values["count"] = h.Count()
																	values["min"] = h.Min()
																	values["max"] = h.Max()
																	values["mean"] = h.Mean()
																	values["stddev"] = h.StdDev()
																	values["median"] = ps[0]
																	values["75%"] = ps[1]
																	values["95%"] = ps[2]
																	values["99%"] = ps[3]
																	values["99.9%"] = ps[4]
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:156
			// _ = "end of CoverTab[96412]"
		case Meter:
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:157
			_go_fuzz_dep_.CoverTab[96413]++
																	m := metric.Snapshot()
																	values["count"] = m.Count()
																	values["1m.rate"] = m.Rate1()
																	values["5m.rate"] = m.Rate5()
																	values["15m.rate"] = m.Rate15()
																	values["mean.rate"] = m.RateMean()
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:163
			// _ = "end of CoverTab[96413]"
		case Timer:
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:164
			_go_fuzz_dep_.CoverTab[96414]++
																	t := metric.Snapshot()
																	ps := t.Percentiles([]float64{0.5, 0.75, 0.95, 0.99, 0.999})
																	values["count"] = t.Count()
																	values["min"] = t.Min()
																	values["max"] = t.Max()
																	values["mean"] = t.Mean()
																	values["stddev"] = t.StdDev()
																	values["median"] = ps[0]
																	values["75%"] = ps[1]
																	values["95%"] = ps[2]
																	values["99%"] = ps[3]
																	values["99.9%"] = ps[4]
																	values["1m.rate"] = t.Rate1()
																	values["5m.rate"] = t.Rate5()
																	values["15m.rate"] = t.Rate15()
																	values["mean.rate"] = t.RateMean()
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:180
			// _ = "end of CoverTab[96414]"
		}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:181
		// _ = "end of CoverTab[96406]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:181
		_go_fuzz_dep_.CoverTab[96407]++
																data[name] = values
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:182
		// _ = "end of CoverTab[96407]"
	})
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:183
	// _ = "end of CoverTab[96404]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:183
	_go_fuzz_dep_.CoverTab[96405]++
															return data
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:184
	// _ = "end of CoverTab[96405]"
}

// Unregister the metric with the given name.
func (r *StandardRegistry) Unregister(name string) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:188
	_go_fuzz_dep_.CoverTab[96417]++
															r.mutex.Lock()
															defer r.mutex.Unlock()
															r.stop(name)
															delete(r.metrics, name)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:192
	// _ = "end of CoverTab[96417]"
}

// Unregister all metrics.  (Mostly for testing.)
func (r *StandardRegistry) UnregisterAll() {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:196
	_go_fuzz_dep_.CoverTab[96418]++
															r.mutex.Lock()
															defer r.mutex.Unlock()
															for name, _ := range r.metrics {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:199
		_go_fuzz_dep_.CoverTab[96419]++
																r.stop(name)
																delete(r.metrics, name)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:201
		// _ = "end of CoverTab[96419]"
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:202
	// _ = "end of CoverTab[96418]"
}

func (r *StandardRegistry) register(name string, i interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:205
	_go_fuzz_dep_.CoverTab[96420]++
															if _, ok := r.metrics[name]; ok {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:206
		_go_fuzz_dep_.CoverTab[96423]++
																return DuplicateMetric(name)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:207
		// _ = "end of CoverTab[96423]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:208
		_go_fuzz_dep_.CoverTab[96424]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:208
		// _ = "end of CoverTab[96424]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:208
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:208
	// _ = "end of CoverTab[96420]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:208
	_go_fuzz_dep_.CoverTab[96421]++
															switch i.(type) {
	case Counter, Gauge, GaugeFloat64, Healthcheck, Histogram, Meter, Timer:
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:210
		_go_fuzz_dep_.CoverTab[96425]++
																r.metrics[name] = i
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:211
		// _ = "end of CoverTab[96425]"
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:212
	// _ = "end of CoverTab[96421]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:212
	_go_fuzz_dep_.CoverTab[96422]++
															return nil
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:213
	// _ = "end of CoverTab[96422]"
}

type metricKV struct {
	name	string
	value	interface{}
}

func (r *StandardRegistry) registered() []metricKV {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:221
	_go_fuzz_dep_.CoverTab[96426]++
															r.mutex.RLock()
															defer r.mutex.RUnlock()
															metrics := make([]metricKV, 0, len(r.metrics))
															for name, i := range r.metrics {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:225
		_go_fuzz_dep_.CoverTab[96428]++
																metrics = append(metrics, metricKV{
			name:	name,
			value:	i,
		})
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:229
		// _ = "end of CoverTab[96428]"
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:230
	// _ = "end of CoverTab[96426]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:230
	_go_fuzz_dep_.CoverTab[96427]++
															return metrics
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:231
	// _ = "end of CoverTab[96427]"
}

func (r *StandardRegistry) stop(name string) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:234
	_go_fuzz_dep_.CoverTab[96429]++
															if i, ok := r.metrics[name]; ok {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:235
		_go_fuzz_dep_.CoverTab[96430]++
																if s, ok := i.(Stoppable); ok {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:236
			_go_fuzz_dep_.CoverTab[96431]++
																	s.Stop()
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:237
			// _ = "end of CoverTab[96431]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:238
			_go_fuzz_dep_.CoverTab[96432]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:238
			// _ = "end of CoverTab[96432]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:238
		}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:238
		// _ = "end of CoverTab[96430]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:239
		_go_fuzz_dep_.CoverTab[96433]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:239
		// _ = "end of CoverTab[96433]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:239
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:239
	// _ = "end of CoverTab[96429]"
}

// Stoppable defines the metrics which has to be stopped.
type Stoppable interface {
	Stop()
}

type PrefixedRegistry struct {
	underlying	Registry
	prefix		string
}

func NewPrefixedRegistry(prefix string) Registry {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:252
	_go_fuzz_dep_.CoverTab[96434]++
															return &PrefixedRegistry{
		underlying:	NewRegistry(),
		prefix:		prefix,
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:256
	// _ = "end of CoverTab[96434]"
}

func NewPrefixedChildRegistry(parent Registry, prefix string) Registry {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:259
	_go_fuzz_dep_.CoverTab[96435]++
															return &PrefixedRegistry{
		underlying:	parent,
		prefix:		prefix,
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:263
	// _ = "end of CoverTab[96435]"
}

// Call the given function for each registered metric.
func (r *PrefixedRegistry) Each(fn func(string, interface{})) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:267
	_go_fuzz_dep_.CoverTab[96436]++
															wrappedFn := func(prefix string) func(string, interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:268
		_go_fuzz_dep_.CoverTab[96438]++
																return func(name string, iface interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:269
			_go_fuzz_dep_.CoverTab[96439]++
																	if strings.HasPrefix(name, prefix) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:270
				_go_fuzz_dep_.CoverTab[96440]++
																		fn(name, iface)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:271
				// _ = "end of CoverTab[96440]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:272
				_go_fuzz_dep_.CoverTab[96441]++
																		return
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:273
				// _ = "end of CoverTab[96441]"
			}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:274
			// _ = "end of CoverTab[96439]"
		}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:275
		// _ = "end of CoverTab[96438]"
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:276
	// _ = "end of CoverTab[96436]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:276
	_go_fuzz_dep_.CoverTab[96437]++

															baseRegistry, prefix := findPrefix(r, "")
															baseRegistry.Each(wrappedFn(prefix))
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:279
	// _ = "end of CoverTab[96437]"
}

func findPrefix(registry Registry, prefix string) (Registry, string) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:282
	_go_fuzz_dep_.CoverTab[96442]++
															switch r := registry.(type) {
	case *PrefixedRegistry:
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:284
		_go_fuzz_dep_.CoverTab[96444]++
																return findPrefix(r.underlying, r.prefix+prefix)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:285
		// _ = "end of CoverTab[96444]"
	case *StandardRegistry:
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:286
		_go_fuzz_dep_.CoverTab[96445]++
																return r, prefix
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:287
		// _ = "end of CoverTab[96445]"
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:288
	// _ = "end of CoverTab[96442]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:288
	_go_fuzz_dep_.CoverTab[96443]++
															return nil, ""
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:289
	// _ = "end of CoverTab[96443]"
}

// Get the metric by the given name or nil if none is registered.
func (r *PrefixedRegistry) Get(name string) interface{} {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:293
	_go_fuzz_dep_.CoverTab[96446]++
															realName := r.prefix + name
															return r.underlying.Get(realName)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:295
	// _ = "end of CoverTab[96446]"
}

// Gets an existing metric or registers the given one.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:298
// The interface can be the metric to register if not found in registry,
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:298
// or a function returning the metric for lazy instantiation.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:301
func (r *PrefixedRegistry) GetOrRegister(name string, metric interface{}) interface{} {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:301
	_go_fuzz_dep_.CoverTab[96447]++
															realName := r.prefix + name
															return r.underlying.GetOrRegister(realName, metric)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:303
	// _ = "end of CoverTab[96447]"
}

// Register the given metric under the given name. The name will be prefixed.
func (r *PrefixedRegistry) Register(name string, metric interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:307
	_go_fuzz_dep_.CoverTab[96448]++
															realName := r.prefix + name
															return r.underlying.Register(realName, metric)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:309
	// _ = "end of CoverTab[96448]"
}

// Run all registered healthchecks.
func (r *PrefixedRegistry) RunHealthchecks() {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:313
	_go_fuzz_dep_.CoverTab[96449]++
															r.underlying.RunHealthchecks()
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:314
	// _ = "end of CoverTab[96449]"
}

// GetAll metrics in the Registry
func (r *PrefixedRegistry) GetAll() map[string]map[string]interface{} {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:318
	_go_fuzz_dep_.CoverTab[96450]++
															return r.underlying.GetAll()
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:319
	// _ = "end of CoverTab[96450]"
}

// Unregister the metric with the given name. The name will be prefixed.
func (r *PrefixedRegistry) Unregister(name string) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:323
	_go_fuzz_dep_.CoverTab[96451]++
															realName := r.prefix + name
															r.underlying.Unregister(realName)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:325
	// _ = "end of CoverTab[96451]"
}

// Unregister all metrics.  (Mostly for testing.)
func (r *PrefixedRegistry) UnregisterAll() {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:329
	_go_fuzz_dep_.CoverTab[96452]++
															r.underlying.UnregisterAll()
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:330
	// _ = "end of CoverTab[96452]"
}

var DefaultRegistry Registry = NewRegistry()

// Call the given function for each registered metric.
func Each(f func(string, interface{})) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:336
	_go_fuzz_dep_.CoverTab[96453]++
															DefaultRegistry.Each(f)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:337
	// _ = "end of CoverTab[96453]"
}

// Get the metric by the given name or nil if none is registered.
func Get(name string) interface{} {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:341
	_go_fuzz_dep_.CoverTab[96454]++
															return DefaultRegistry.Get(name)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:342
	// _ = "end of CoverTab[96454]"
}

// Gets an existing metric or creates and registers a new one. Threadsafe
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:345
// alternative to calling Get and Register on failure.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:347
func GetOrRegister(name string, i interface{}) interface{} {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:347
	_go_fuzz_dep_.CoverTab[96455]++
															return DefaultRegistry.GetOrRegister(name, i)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:348
	// _ = "end of CoverTab[96455]"
}

// Register the given metric under the given name.  Returns a DuplicateMetric
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:351
// if a metric by the given name is already registered.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:353
func Register(name string, i interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:353
	_go_fuzz_dep_.CoverTab[96456]++
															return DefaultRegistry.Register(name, i)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:354
	// _ = "end of CoverTab[96456]"
}

// Register the given metric under the given name.  Panics if a metric by the
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:357
// given name is already registered.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:359
func MustRegister(name string, i interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:359
	_go_fuzz_dep_.CoverTab[96457]++
															if err := Register(name, i); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:360
		_go_fuzz_dep_.CoverTab[96458]++
																panic(err)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:361
		// _ = "end of CoverTab[96458]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:362
		_go_fuzz_dep_.CoverTab[96459]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:362
		// _ = "end of CoverTab[96459]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:362
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:362
	// _ = "end of CoverTab[96457]"
}

// Run all registered healthchecks.
func RunHealthchecks() {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:366
	_go_fuzz_dep_.CoverTab[96460]++
															DefaultRegistry.RunHealthchecks()
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:367
	// _ = "end of CoverTab[96460]"
}

// Unregister the metric with the given name.
func Unregister(name string) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:371
	_go_fuzz_dep_.CoverTab[96461]++
															DefaultRegistry.Unregister(name)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:372
	// _ = "end of CoverTab[96461]"
}

//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:373
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/registry.go:373
var _ = _go_fuzz_dep_.CoverTab
