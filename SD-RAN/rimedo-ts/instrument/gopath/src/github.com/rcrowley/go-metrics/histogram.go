//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:1
package metrics

//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:1
)

// Histograms calculate distribution statistics from a series of int64 values.
type Histogram interface {
	Clear()
	Count() int64
	Max() int64
	Mean() float64
	Min() int64
	Percentile(float64) float64
	Percentiles([]float64) []float64
	Sample() Sample
	Snapshot() Histogram
	StdDev() float64
	Sum() int64
	Update(int64)
	Variance() float64
}

// GetOrRegisterHistogram returns an existing Histogram or constructs and
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:20
// registers a new StandardHistogram.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:22
func GetOrRegisterHistogram(name string, r Registry, s Sample) Histogram {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:22
	_go_fuzz_dep_.CoverTab[96234]++
															if nil == r {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:23
		_go_fuzz_dep_.CoverTab[96236]++
																r = DefaultRegistry
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:24
		// _ = "end of CoverTab[96236]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:25
		_go_fuzz_dep_.CoverTab[96237]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:25
		// _ = "end of CoverTab[96237]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:25
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:25
	// _ = "end of CoverTab[96234]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:25
	_go_fuzz_dep_.CoverTab[96235]++
															return r.GetOrRegister(name, func() Histogram {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:26
		_go_fuzz_dep_.CoverTab[96238]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:26
		return NewHistogram(s)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:26
		// _ = "end of CoverTab[96238]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:26
	}).(Histogram)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:26
	// _ = "end of CoverTab[96235]"
}

// NewHistogram constructs a new StandardHistogram from a Sample.
func NewHistogram(s Sample) Histogram {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:30
	_go_fuzz_dep_.CoverTab[96239]++
															if UseNilMetrics {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:31
		_go_fuzz_dep_.CoverTab[96241]++
																return NilHistogram{}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:32
		// _ = "end of CoverTab[96241]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:33
		_go_fuzz_dep_.CoverTab[96242]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:33
		// _ = "end of CoverTab[96242]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:33
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:33
	// _ = "end of CoverTab[96239]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:33
	_go_fuzz_dep_.CoverTab[96240]++
															return &StandardHistogram{sample: s}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:34
	// _ = "end of CoverTab[96240]"
}

// NewRegisteredHistogram constructs and registers a new StandardHistogram from
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:37
// a Sample.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:39
func NewRegisteredHistogram(name string, r Registry, s Sample) Histogram {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:39
	_go_fuzz_dep_.CoverTab[96243]++
															c := NewHistogram(s)
															if nil == r {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:41
		_go_fuzz_dep_.CoverTab[96245]++
																r = DefaultRegistry
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:42
		// _ = "end of CoverTab[96245]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:43
		_go_fuzz_dep_.CoverTab[96246]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:43
		// _ = "end of CoverTab[96246]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:43
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:43
	// _ = "end of CoverTab[96243]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:43
	_go_fuzz_dep_.CoverTab[96244]++
															r.Register(name, c)
															return c
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:45
	// _ = "end of CoverTab[96244]"
}

// HistogramSnapshot is a read-only copy of another Histogram.
type HistogramSnapshot struct {
	sample *SampleSnapshot
}

// Clear panics.
func (*HistogramSnapshot) Clear() {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:54
	_go_fuzz_dep_.CoverTab[96247]++
															panic("Clear called on a HistogramSnapshot")
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:55
	// _ = "end of CoverTab[96247]"
}

// Count returns the number of samples recorded at the time the snapshot was
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:58
// taken.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:60
func (h *HistogramSnapshot) Count() int64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:60
	_go_fuzz_dep_.CoverTab[96248]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:60
	return h.sample.Count()
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:60
	// _ = "end of CoverTab[96248]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:60
}

// Max returns the maximum value in the sample at the time the snapshot was
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:62
// taken.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:64
func (h *HistogramSnapshot) Max() int64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:64
	_go_fuzz_dep_.CoverTab[96249]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:64
	return h.sample.Max()
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:64
	// _ = "end of CoverTab[96249]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:64
}

// Mean returns the mean of the values in the sample at the time the snapshot
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:66
// was taken.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:68
func (h *HistogramSnapshot) Mean() float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:68
	_go_fuzz_dep_.CoverTab[96250]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:68
	return h.sample.Mean()
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:68
	// _ = "end of CoverTab[96250]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:68
}

// Min returns the minimum value in the sample at the time the snapshot was
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:70
// taken.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:72
func (h *HistogramSnapshot) Min() int64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:72
	_go_fuzz_dep_.CoverTab[96251]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:72
	return h.sample.Min()
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:72
	// _ = "end of CoverTab[96251]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:72
}

// Percentile returns an arbitrary percentile of values in the sample at the
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:74
// time the snapshot was taken.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:76
func (h *HistogramSnapshot) Percentile(p float64) float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:76
	_go_fuzz_dep_.CoverTab[96252]++
															return h.sample.Percentile(p)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:77
	// _ = "end of CoverTab[96252]"
}

// Percentiles returns a slice of arbitrary percentiles of values in the sample
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:80
// at the time the snapshot was taken.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:82
func (h *HistogramSnapshot) Percentiles(ps []float64) []float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:82
	_go_fuzz_dep_.CoverTab[96253]++
															return h.sample.Percentiles(ps)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:83
	// _ = "end of CoverTab[96253]"
}

// Sample returns the Sample underlying the histogram.
func (h *HistogramSnapshot) Sample() Sample {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:87
	_go_fuzz_dep_.CoverTab[96254]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:87
	return h.sample
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:87
	// _ = "end of CoverTab[96254]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:87
}

// Snapshot returns the snapshot.
func (h *HistogramSnapshot) Snapshot() Histogram {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:90
	_go_fuzz_dep_.CoverTab[96255]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:90
	return h
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:90
	// _ = "end of CoverTab[96255]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:90
}

// StdDev returns the standard deviation of the values in the sample at the
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:92
// time the snapshot was taken.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:94
func (h *HistogramSnapshot) StdDev() float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:94
	_go_fuzz_dep_.CoverTab[96256]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:94
	return h.sample.StdDev()
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:94
	// _ = "end of CoverTab[96256]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:94
}

// Sum returns the sum in the sample at the time the snapshot was taken.
func (h *HistogramSnapshot) Sum() int64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:97
	_go_fuzz_dep_.CoverTab[96257]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:97
	return h.sample.Sum()
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:97
	// _ = "end of CoverTab[96257]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:97
}

// Update panics.
func (*HistogramSnapshot) Update(int64) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:100
	_go_fuzz_dep_.CoverTab[96258]++
															panic("Update called on a HistogramSnapshot")
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:101
	// _ = "end of CoverTab[96258]"
}

// Variance returns the variance of inputs at the time the snapshot was taken.
func (h *HistogramSnapshot) Variance() float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:105
	_go_fuzz_dep_.CoverTab[96259]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:105
	return h.sample.Variance()
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:105
	// _ = "end of CoverTab[96259]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:105
}

// NilHistogram is a no-op Histogram.
type NilHistogram struct{}

// Clear is a no-op.
func (NilHistogram) Clear()	{ _go_fuzz_dep_.CoverTab[96260]++; // _ = "end of CoverTab[96260]" }

// Count is a no-op.
func (NilHistogram) Count() int64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:114
	_go_fuzz_dep_.CoverTab[96261]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:114
	return 0
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:114
	// _ = "end of CoverTab[96261]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:114
}

// Max is a no-op.
func (NilHistogram) Max() int64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:117
	_go_fuzz_dep_.CoverTab[96262]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:117
	return 0
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:117
	// _ = "end of CoverTab[96262]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:117
}

// Mean is a no-op.
func (NilHistogram) Mean() float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:120
	_go_fuzz_dep_.CoverTab[96263]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:120
	return 0.0
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:120
	// _ = "end of CoverTab[96263]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:120
}

// Min is a no-op.
func (NilHistogram) Min() int64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:123
	_go_fuzz_dep_.CoverTab[96264]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:123
	return 0
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:123
	// _ = "end of CoverTab[96264]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:123
}

// Percentile is a no-op.
func (NilHistogram) Percentile(p float64) float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:126
	_go_fuzz_dep_.CoverTab[96265]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:126
	return 0.0
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:126
	// _ = "end of CoverTab[96265]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:126
}

// Percentiles is a no-op.
func (NilHistogram) Percentiles(ps []float64) []float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:129
	_go_fuzz_dep_.CoverTab[96266]++
															return make([]float64, len(ps))
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:130
	// _ = "end of CoverTab[96266]"
}

// Sample is a no-op.
func (NilHistogram) Sample() Sample {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:134
	_go_fuzz_dep_.CoverTab[96267]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:134
	return NilSample{}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:134
	// _ = "end of CoverTab[96267]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:134
}

// Snapshot is a no-op.
func (NilHistogram) Snapshot() Histogram {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:137
	_go_fuzz_dep_.CoverTab[96268]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:137
	return NilHistogram{}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:137
	// _ = "end of CoverTab[96268]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:137
}

// StdDev is a no-op.
func (NilHistogram) StdDev() float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:140
	_go_fuzz_dep_.CoverTab[96269]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:140
	return 0.0
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:140
	// _ = "end of CoverTab[96269]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:140
}

// Sum is a no-op.
func (NilHistogram) Sum() int64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:143
	_go_fuzz_dep_.CoverTab[96270]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:143
	return 0
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:143
	// _ = "end of CoverTab[96270]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:143
}

// Update is a no-op.
func (NilHistogram) Update(v int64)	{ _go_fuzz_dep_.CoverTab[96271]++; // _ = "end of CoverTab[96271]" }

// Variance is a no-op.
func (NilHistogram) Variance() float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:149
	_go_fuzz_dep_.CoverTab[96272]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:149
	return 0.0
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:149
	// _ = "end of CoverTab[96272]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:149
}

// StandardHistogram is the standard implementation of a Histogram and uses a
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:151
// Sample to bound its memory use.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:153
type StandardHistogram struct {
	sample Sample
}

// Clear clears the histogram and its sample.
func (h *StandardHistogram) Clear() {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:158
	_go_fuzz_dep_.CoverTab[96273]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:158
	h.sample.Clear()
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:158
	// _ = "end of CoverTab[96273]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:158
}

// Count returns the number of samples recorded since the histogram was last
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:160
// cleared.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:162
func (h *StandardHistogram) Count() int64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:162
	_go_fuzz_dep_.CoverTab[96274]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:162
	return h.sample.Count()
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:162
	// _ = "end of CoverTab[96274]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:162
}

// Max returns the maximum value in the sample.
func (h *StandardHistogram) Max() int64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:165
	_go_fuzz_dep_.CoverTab[96275]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:165
	return h.sample.Max()
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:165
	// _ = "end of CoverTab[96275]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:165
}

// Mean returns the mean of the values in the sample.
func (h *StandardHistogram) Mean() float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:168
	_go_fuzz_dep_.CoverTab[96276]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:168
	return h.sample.Mean()
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:168
	// _ = "end of CoverTab[96276]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:168
}

// Min returns the minimum value in the sample.
func (h *StandardHistogram) Min() int64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:171
	_go_fuzz_dep_.CoverTab[96277]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:171
	return h.sample.Min()
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:171
	// _ = "end of CoverTab[96277]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:171
}

// Percentile returns an arbitrary percentile of the values in the sample.
func (h *StandardHistogram) Percentile(p float64) float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:174
	_go_fuzz_dep_.CoverTab[96278]++
															return h.sample.Percentile(p)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:175
	// _ = "end of CoverTab[96278]"
}

// Percentiles returns a slice of arbitrary percentiles of the values in the
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:178
// sample.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:180
func (h *StandardHistogram) Percentiles(ps []float64) []float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:180
	_go_fuzz_dep_.CoverTab[96279]++
															return h.sample.Percentiles(ps)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:181
	// _ = "end of CoverTab[96279]"
}

// Sample returns the Sample underlying the histogram.
func (h *StandardHistogram) Sample() Sample {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:185
	_go_fuzz_dep_.CoverTab[96280]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:185
	return h.sample
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:185
	// _ = "end of CoverTab[96280]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:185
}

// Snapshot returns a read-only copy of the histogram.
func (h *StandardHistogram) Snapshot() Histogram {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:188
	_go_fuzz_dep_.CoverTab[96281]++
															return &HistogramSnapshot{sample: h.sample.Snapshot().(*SampleSnapshot)}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:189
	// _ = "end of CoverTab[96281]"
}

// StdDev returns the standard deviation of the values in the sample.
func (h *StandardHistogram) StdDev() float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:193
	_go_fuzz_dep_.CoverTab[96282]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:193
	return h.sample.StdDev()
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:193
	// _ = "end of CoverTab[96282]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:193
}

// Sum returns the sum in the sample.
func (h *StandardHistogram) Sum() int64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:196
	_go_fuzz_dep_.CoverTab[96283]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:196
	return h.sample.Sum()
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:196
	// _ = "end of CoverTab[96283]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:196
}

// Update samples a new value.
func (h *StandardHistogram) Update(v int64) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:199
	_go_fuzz_dep_.CoverTab[96284]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:199
	h.sample.Update(v)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:199
	// _ = "end of CoverTab[96284]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:199
}

// Variance returns the variance of the values in the sample.
func (h *StandardHistogram) Variance() float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:202
	_go_fuzz_dep_.CoverTab[96285]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:202
	return h.sample.Variance()
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:202
	// _ = "end of CoverTab[96285]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:202
}

//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:202
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/histogram.go:202
var _ = _go_fuzz_dep_.CoverTab
