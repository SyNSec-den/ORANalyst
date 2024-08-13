//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:1
package metrics

//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:1
)

import (
	"math"
	"math/rand"
	"sort"
	"sync"
	"time"
)

const rescaleThreshold = time.Hour

// Samples maintain a statistically-significant selection of values from
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:13
// a stream.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:15
type Sample interface {
	Clear()
	Count() int64
	Max() int64
	Mean() float64
	Min() int64
	Percentile(float64) float64
	Percentiles([]float64) []float64
	Size() int
	Snapshot() Sample
	StdDev() float64
	Sum() int64
	Update(int64)
	Values() []int64
	Variance() float64
}

// ExpDecaySample is an exponentially-decaying sample using a forward-decaying
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:32
// priority reservoir.  See Cormode et al's "Forward Decay: A Practical Time
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:32
// Decay Model for Streaming Systems".
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:32
//
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:32
// <http://dimacs.rutgers.edu/~graham/pubs/papers/fwddecay.pdf>
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:37
type ExpDecaySample struct {
	alpha		float64
	count		int64
	mutex		sync.Mutex
	reservoirSize	int
	t0, t1		time.Time
	values		*expDecaySampleHeap
}

// NewExpDecaySample constructs a new exponentially-decaying sample with the
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:46
// given reservoir size and alpha.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:48
func NewExpDecaySample(reservoirSize int, alpha float64) Sample {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:48
	_go_fuzz_dep_.CoverTab[96485]++
															if UseNilMetrics {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:49
		_go_fuzz_dep_.CoverTab[96487]++
																return NilSample{}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:50
		// _ = "end of CoverTab[96487]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:51
		_go_fuzz_dep_.CoverTab[96488]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:51
		// _ = "end of CoverTab[96488]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:51
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:51
	// _ = "end of CoverTab[96485]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:51
	_go_fuzz_dep_.CoverTab[96486]++
															s := &ExpDecaySample{
		alpha:		alpha,
		reservoirSize:	reservoirSize,
		t0:		time.Now(),
		values:		newExpDecaySampleHeap(reservoirSize),
	}
															s.t1 = s.t0.Add(rescaleThreshold)
															return s
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:59
	// _ = "end of CoverTab[96486]"
}

// Clear clears all samples.
func (s *ExpDecaySample) Clear() {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:63
	_go_fuzz_dep_.CoverTab[96489]++
															s.mutex.Lock()
															defer s.mutex.Unlock()
															s.count = 0
															s.t0 = time.Now()
															s.t1 = s.t0.Add(rescaleThreshold)
															s.values.Clear()
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:69
	// _ = "end of CoverTab[96489]"
}

// Count returns the number of samples recorded, which may exceed the
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:72
// reservoir size.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:74
func (s *ExpDecaySample) Count() int64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:74
	_go_fuzz_dep_.CoverTab[96490]++
															s.mutex.Lock()
															defer s.mutex.Unlock()
															return s.count
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:77
	// _ = "end of CoverTab[96490]"
}

// Max returns the maximum value in the sample, which may not be the maximum
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:80
// value ever to be part of the sample.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:82
func (s *ExpDecaySample) Max() int64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:82
	_go_fuzz_dep_.CoverTab[96491]++
															return SampleMax(s.Values())
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:83
	// _ = "end of CoverTab[96491]"
}

// Mean returns the mean of the values in the sample.
func (s *ExpDecaySample) Mean() float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:87
	_go_fuzz_dep_.CoverTab[96492]++
															return SampleMean(s.Values())
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:88
	// _ = "end of CoverTab[96492]"
}

// Min returns the minimum value in the sample, which may not be the minimum
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:91
// value ever to be part of the sample.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:93
func (s *ExpDecaySample) Min() int64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:93
	_go_fuzz_dep_.CoverTab[96493]++
															return SampleMin(s.Values())
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:94
	// _ = "end of CoverTab[96493]"
}

// Percentile returns an arbitrary percentile of values in the sample.
func (s *ExpDecaySample) Percentile(p float64) float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:98
	_go_fuzz_dep_.CoverTab[96494]++
															return SamplePercentile(s.Values(), p)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:99
	// _ = "end of CoverTab[96494]"
}

// Percentiles returns a slice of arbitrary percentiles of values in the
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:102
// sample.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:104
func (s *ExpDecaySample) Percentiles(ps []float64) []float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:104
	_go_fuzz_dep_.CoverTab[96495]++
															return SamplePercentiles(s.Values(), ps)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:105
	// _ = "end of CoverTab[96495]"
}

// Size returns the size of the sample, which is at most the reservoir size.
func (s *ExpDecaySample) Size() int {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:109
	_go_fuzz_dep_.CoverTab[96496]++
															s.mutex.Lock()
															defer s.mutex.Unlock()
															return s.values.Size()
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:112
	// _ = "end of CoverTab[96496]"
}

// Snapshot returns a read-only copy of the sample.
func (s *ExpDecaySample) Snapshot() Sample {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:116
	_go_fuzz_dep_.CoverTab[96497]++
															s.mutex.Lock()
															defer s.mutex.Unlock()
															vals := s.values.Values()
															values := make([]int64, len(vals))
															for i, v := range vals {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:121
		_go_fuzz_dep_.CoverTab[96499]++
																values[i] = v.v
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:122
		// _ = "end of CoverTab[96499]"
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:123
	// _ = "end of CoverTab[96497]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:123
	_go_fuzz_dep_.CoverTab[96498]++
															return &SampleSnapshot{
		count:	s.count,
		values:	values,
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:127
	// _ = "end of CoverTab[96498]"
}

// StdDev returns the standard deviation of the values in the sample.
func (s *ExpDecaySample) StdDev() float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:131
	_go_fuzz_dep_.CoverTab[96500]++
															return SampleStdDev(s.Values())
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:132
	// _ = "end of CoverTab[96500]"
}

// Sum returns the sum of the values in the sample.
func (s *ExpDecaySample) Sum() int64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:136
	_go_fuzz_dep_.CoverTab[96501]++
															return SampleSum(s.Values())
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:137
	// _ = "end of CoverTab[96501]"
}

// Update samples a new value.
func (s *ExpDecaySample) Update(v int64) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:141
	_go_fuzz_dep_.CoverTab[96502]++
															s.update(time.Now(), v)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:142
	// _ = "end of CoverTab[96502]"
}

// Values returns a copy of the values in the sample.
func (s *ExpDecaySample) Values() []int64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:146
	_go_fuzz_dep_.CoverTab[96503]++
															s.mutex.Lock()
															defer s.mutex.Unlock()
															vals := s.values.Values()
															values := make([]int64, len(vals))
															for i, v := range vals {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:151
		_go_fuzz_dep_.CoverTab[96505]++
																values[i] = v.v
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:152
		// _ = "end of CoverTab[96505]"
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:153
	// _ = "end of CoverTab[96503]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:153
	_go_fuzz_dep_.CoverTab[96504]++
															return values
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:154
	// _ = "end of CoverTab[96504]"
}

// Variance returns the variance of the values in the sample.
func (s *ExpDecaySample) Variance() float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:158
	_go_fuzz_dep_.CoverTab[96506]++
															return SampleVariance(s.Values())
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:159
	// _ = "end of CoverTab[96506]"
}

// update samples a new value at a particular timestamp.  This is a method all
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:162
// its own to facilitate testing.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:164
func (s *ExpDecaySample) update(t time.Time, v int64) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:164
	_go_fuzz_dep_.CoverTab[96507]++
															s.mutex.Lock()
															defer s.mutex.Unlock()
															s.count++
															if s.values.Size() == s.reservoirSize {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:168
		_go_fuzz_dep_.CoverTab[96509]++
																s.values.Pop()
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:169
		// _ = "end of CoverTab[96509]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:170
		_go_fuzz_dep_.CoverTab[96510]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:170
		// _ = "end of CoverTab[96510]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:170
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:170
	// _ = "end of CoverTab[96507]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:170
	_go_fuzz_dep_.CoverTab[96508]++
															s.values.Push(expDecaySample{
		k:	math.Exp(t.Sub(s.t0).Seconds()*s.alpha) / rand.Float64(),
		v:	v,
	})
	if t.After(s.t1) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:175
		_go_fuzz_dep_.CoverTab[96511]++
																values := s.values.Values()
																t0 := s.t0
																s.values.Clear()
																s.t0 = t
																s.t1 = s.t0.Add(rescaleThreshold)
																for _, v := range values {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:181
			_go_fuzz_dep_.CoverTab[96512]++
																	v.k = v.k * math.Exp(-s.alpha*s.t0.Sub(t0).Seconds())
																	s.values.Push(v)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:183
			// _ = "end of CoverTab[96512]"
		}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:184
		// _ = "end of CoverTab[96511]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:185
		_go_fuzz_dep_.CoverTab[96513]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:185
		// _ = "end of CoverTab[96513]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:185
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:185
	// _ = "end of CoverTab[96508]"
}

// NilSample is a no-op Sample.
type NilSample struct{}

// Clear is a no-op.
func (NilSample) Clear()	{ _go_fuzz_dep_.CoverTab[96514]++; // _ = "end of CoverTab[96514]" }

// Count is a no-op.
func (NilSample) Count() int64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:195
	_go_fuzz_dep_.CoverTab[96515]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:195
	return 0
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:195
	// _ = "end of CoverTab[96515]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:195
}

// Max is a no-op.
func (NilSample) Max() int64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:198
	_go_fuzz_dep_.CoverTab[96516]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:198
	return 0
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:198
	// _ = "end of CoverTab[96516]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:198
}

// Mean is a no-op.
func (NilSample) Mean() float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:201
	_go_fuzz_dep_.CoverTab[96517]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:201
	return 0.0
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:201
	// _ = "end of CoverTab[96517]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:201
}

// Min is a no-op.
func (NilSample) Min() int64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:204
	_go_fuzz_dep_.CoverTab[96518]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:204
	return 0
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:204
	// _ = "end of CoverTab[96518]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:204
}

// Percentile is a no-op.
func (NilSample) Percentile(p float64) float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:207
	_go_fuzz_dep_.CoverTab[96519]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:207
	return 0.0
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:207
	// _ = "end of CoverTab[96519]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:207
}

// Percentiles is a no-op.
func (NilSample) Percentiles(ps []float64) []float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:210
	_go_fuzz_dep_.CoverTab[96520]++
															return make([]float64, len(ps))
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:211
	// _ = "end of CoverTab[96520]"
}

// Size is a no-op.
func (NilSample) Size() int {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:215
	_go_fuzz_dep_.CoverTab[96521]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:215
	return 0
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:215
	// _ = "end of CoverTab[96521]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:215
}

// Sample is a no-op.
func (NilSample) Snapshot() Sample {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:218
	_go_fuzz_dep_.CoverTab[96522]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:218
	return NilSample{}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:218
	// _ = "end of CoverTab[96522]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:218
}

// StdDev is a no-op.
func (NilSample) StdDev() float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:221
	_go_fuzz_dep_.CoverTab[96523]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:221
	return 0.0
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:221
	// _ = "end of CoverTab[96523]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:221
}

// Sum is a no-op.
func (NilSample) Sum() int64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:224
	_go_fuzz_dep_.CoverTab[96524]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:224
	return 0
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:224
	// _ = "end of CoverTab[96524]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:224
}

// Update is a no-op.
func (NilSample) Update(v int64)	{ _go_fuzz_dep_.CoverTab[96525]++; // _ = "end of CoverTab[96525]" }

// Values is a no-op.
func (NilSample) Values() []int64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:230
	_go_fuzz_dep_.CoverTab[96526]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:230
	return []int64{}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:230
	// _ = "end of CoverTab[96526]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:230
}

// Variance is a no-op.
func (NilSample) Variance() float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:233
	_go_fuzz_dep_.CoverTab[96527]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:233
	return 0.0
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:233
	// _ = "end of CoverTab[96527]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:233
}

// SampleMax returns the maximum value of the slice of int64.
func SampleMax(values []int64) int64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:236
	_go_fuzz_dep_.CoverTab[96528]++
															if 0 == len(values) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:237
		_go_fuzz_dep_.CoverTab[96531]++
																return 0
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:238
		// _ = "end of CoverTab[96531]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:239
		_go_fuzz_dep_.CoverTab[96532]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:239
		// _ = "end of CoverTab[96532]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:239
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:239
	// _ = "end of CoverTab[96528]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:239
	_go_fuzz_dep_.CoverTab[96529]++
															var max int64 = math.MinInt64
															for _, v := range values {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:241
		_go_fuzz_dep_.CoverTab[96533]++
																if max < v {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:242
			_go_fuzz_dep_.CoverTab[96534]++
																	max = v
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:243
			// _ = "end of CoverTab[96534]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:244
			_go_fuzz_dep_.CoverTab[96535]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:244
			// _ = "end of CoverTab[96535]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:244
		}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:244
		// _ = "end of CoverTab[96533]"
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:245
	// _ = "end of CoverTab[96529]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:245
	_go_fuzz_dep_.CoverTab[96530]++
															return max
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:246
	// _ = "end of CoverTab[96530]"
}

// SampleMean returns the mean value of the slice of int64.
func SampleMean(values []int64) float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:250
	_go_fuzz_dep_.CoverTab[96536]++
															if 0 == len(values) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:251
		_go_fuzz_dep_.CoverTab[96538]++
																return 0.0
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:252
		// _ = "end of CoverTab[96538]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:253
		_go_fuzz_dep_.CoverTab[96539]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:253
		// _ = "end of CoverTab[96539]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:253
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:253
	// _ = "end of CoverTab[96536]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:253
	_go_fuzz_dep_.CoverTab[96537]++
															return float64(SampleSum(values)) / float64(len(values))
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:254
	// _ = "end of CoverTab[96537]"
}

// SampleMin returns the minimum value of the slice of int64.
func SampleMin(values []int64) int64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:258
	_go_fuzz_dep_.CoverTab[96540]++
															if 0 == len(values) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:259
		_go_fuzz_dep_.CoverTab[96543]++
																return 0
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:260
		// _ = "end of CoverTab[96543]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:261
		_go_fuzz_dep_.CoverTab[96544]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:261
		// _ = "end of CoverTab[96544]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:261
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:261
	// _ = "end of CoverTab[96540]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:261
	_go_fuzz_dep_.CoverTab[96541]++
															var min int64 = math.MaxInt64
															for _, v := range values {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:263
		_go_fuzz_dep_.CoverTab[96545]++
																if min > v {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:264
			_go_fuzz_dep_.CoverTab[96546]++
																	min = v
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:265
			// _ = "end of CoverTab[96546]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:266
			_go_fuzz_dep_.CoverTab[96547]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:266
			// _ = "end of CoverTab[96547]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:266
		}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:266
		// _ = "end of CoverTab[96545]"
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:267
	// _ = "end of CoverTab[96541]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:267
	_go_fuzz_dep_.CoverTab[96542]++
															return min
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:268
	// _ = "end of CoverTab[96542]"
}

// SamplePercentiles returns an arbitrary percentile of the slice of int64.
func SamplePercentile(values int64Slice, p float64) float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:272
	_go_fuzz_dep_.CoverTab[96548]++
															return SamplePercentiles(values, []float64{p})[0]
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:273
	// _ = "end of CoverTab[96548]"
}

// SamplePercentiles returns a slice of arbitrary percentiles of the slice of
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:276
// int64.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:278
func SamplePercentiles(values int64Slice, ps []float64) []float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:278
	_go_fuzz_dep_.CoverTab[96549]++
															scores := make([]float64, len(ps))
															size := len(values)
															if size > 0 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:281
		_go_fuzz_dep_.CoverTab[96551]++
																sort.Sort(values)
																for i, p := range ps {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:283
			_go_fuzz_dep_.CoverTab[96552]++
																	pos := p * float64(size+1)
																	if pos < 1.0 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:285
				_go_fuzz_dep_.CoverTab[96553]++
																		scores[i] = float64(values[0])
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:286
				// _ = "end of CoverTab[96553]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:287
				_go_fuzz_dep_.CoverTab[96554]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:287
				if pos >= float64(size) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:287
					_go_fuzz_dep_.CoverTab[96555]++
																			scores[i] = float64(values[size-1])
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:288
					// _ = "end of CoverTab[96555]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:289
					_go_fuzz_dep_.CoverTab[96556]++
																			lower := float64(values[int(pos)-1])
																			upper := float64(values[int(pos)])
																			scores[i] = lower + (pos-math.Floor(pos))*(upper-lower)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:292
					// _ = "end of CoverTab[96556]"
				}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:293
				// _ = "end of CoverTab[96554]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:293
			}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:293
			// _ = "end of CoverTab[96552]"
		}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:294
		// _ = "end of CoverTab[96551]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:295
		_go_fuzz_dep_.CoverTab[96557]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:295
		// _ = "end of CoverTab[96557]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:295
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:295
	// _ = "end of CoverTab[96549]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:295
	_go_fuzz_dep_.CoverTab[96550]++
															return scores
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:296
	// _ = "end of CoverTab[96550]"
}

// SampleSnapshot is a read-only copy of another Sample.
type SampleSnapshot struct {
	count	int64
	values	[]int64
}

func NewSampleSnapshot(count int64, values []int64) *SampleSnapshot {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:305
	_go_fuzz_dep_.CoverTab[96558]++
															return &SampleSnapshot{
		count:	count,
		values:	values,
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:309
	// _ = "end of CoverTab[96558]"
}

// Clear panics.
func (*SampleSnapshot) Clear() {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:313
	_go_fuzz_dep_.CoverTab[96559]++
															panic("Clear called on a SampleSnapshot")
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:314
	// _ = "end of CoverTab[96559]"
}

// Count returns the count of inputs at the time the snapshot was taken.
func (s *SampleSnapshot) Count() int64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:318
	_go_fuzz_dep_.CoverTab[96560]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:318
	return s.count
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:318
	// _ = "end of CoverTab[96560]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:318
}

// Max returns the maximal value at the time the snapshot was taken.
func (s *SampleSnapshot) Max() int64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:321
	_go_fuzz_dep_.CoverTab[96561]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:321
	return SampleMax(s.values)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:321
	// _ = "end of CoverTab[96561]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:321
}

// Mean returns the mean value at the time the snapshot was taken.
func (s *SampleSnapshot) Mean() float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:324
	_go_fuzz_dep_.CoverTab[96562]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:324
	return SampleMean(s.values)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:324
	// _ = "end of CoverTab[96562]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:324
}

// Min returns the minimal value at the time the snapshot was taken.
func (s *SampleSnapshot) Min() int64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:327
	_go_fuzz_dep_.CoverTab[96563]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:327
	return SampleMin(s.values)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:327
	// _ = "end of CoverTab[96563]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:327
}

// Percentile returns an arbitrary percentile of values at the time the
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:329
// snapshot was taken.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:331
func (s *SampleSnapshot) Percentile(p float64) float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:331
	_go_fuzz_dep_.CoverTab[96564]++
															return SamplePercentile(s.values, p)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:332
	// _ = "end of CoverTab[96564]"
}

// Percentiles returns a slice of arbitrary percentiles of values at the time
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:335
// the snapshot was taken.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:337
func (s *SampleSnapshot) Percentiles(ps []float64) []float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:337
	_go_fuzz_dep_.CoverTab[96565]++
															return SamplePercentiles(s.values, ps)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:338
	// _ = "end of CoverTab[96565]"
}

// Size returns the size of the sample at the time the snapshot was taken.
func (s *SampleSnapshot) Size() int {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:342
	_go_fuzz_dep_.CoverTab[96566]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:342
	return len(s.values)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:342
	// _ = "end of CoverTab[96566]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:342
}

// Snapshot returns the snapshot.
func (s *SampleSnapshot) Snapshot() Sample {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:345
	_go_fuzz_dep_.CoverTab[96567]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:345
	return s
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:345
	// _ = "end of CoverTab[96567]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:345
}

// StdDev returns the standard deviation of values at the time the snapshot was
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:347
// taken.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:349
func (s *SampleSnapshot) StdDev() float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:349
	_go_fuzz_dep_.CoverTab[96568]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:349
	return SampleStdDev(s.values)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:349
	// _ = "end of CoverTab[96568]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:349
}

// Sum returns the sum of values at the time the snapshot was taken.
func (s *SampleSnapshot) Sum() int64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:352
	_go_fuzz_dep_.CoverTab[96569]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:352
	return SampleSum(s.values)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:352
	// _ = "end of CoverTab[96569]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:352
}

// Update panics.
func (*SampleSnapshot) Update(int64) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:355
	_go_fuzz_dep_.CoverTab[96570]++
															panic("Update called on a SampleSnapshot")
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:356
	// _ = "end of CoverTab[96570]"
}

// Values returns a copy of the values in the sample.
func (s *SampleSnapshot) Values() []int64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:360
	_go_fuzz_dep_.CoverTab[96571]++
															values := make([]int64, len(s.values))
															copy(values, s.values)
															return values
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:363
	// _ = "end of CoverTab[96571]"
}

// Variance returns the variance of values at the time the snapshot was taken.
func (s *SampleSnapshot) Variance() float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:367
	_go_fuzz_dep_.CoverTab[96572]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:367
	return SampleVariance(s.values)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:367
	// _ = "end of CoverTab[96572]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:367
}

// SampleStdDev returns the standard deviation of the slice of int64.
func SampleStdDev(values []int64) float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:370
	_go_fuzz_dep_.CoverTab[96573]++
															return math.Sqrt(SampleVariance(values))
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:371
	// _ = "end of CoverTab[96573]"
}

// SampleSum returns the sum of the slice of int64.
func SampleSum(values []int64) int64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:375
	_go_fuzz_dep_.CoverTab[96574]++
															var sum int64
															for _, v := range values {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:377
		_go_fuzz_dep_.CoverTab[96576]++
																sum += v
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:378
		// _ = "end of CoverTab[96576]"
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:379
	// _ = "end of CoverTab[96574]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:379
	_go_fuzz_dep_.CoverTab[96575]++
															return sum
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:380
	// _ = "end of CoverTab[96575]"
}

// SampleVariance returns the variance of the slice of int64.
func SampleVariance(values []int64) float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:384
	_go_fuzz_dep_.CoverTab[96577]++
															if 0 == len(values) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:385
		_go_fuzz_dep_.CoverTab[96580]++
																return 0.0
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:386
		// _ = "end of CoverTab[96580]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:387
		_go_fuzz_dep_.CoverTab[96581]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:387
		// _ = "end of CoverTab[96581]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:387
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:387
	// _ = "end of CoverTab[96577]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:387
	_go_fuzz_dep_.CoverTab[96578]++
															m := SampleMean(values)
															var sum float64
															for _, v := range values {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:390
		_go_fuzz_dep_.CoverTab[96582]++
																d := float64(v) - m
																sum += d * d
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:392
		// _ = "end of CoverTab[96582]"
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:393
	// _ = "end of CoverTab[96578]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:393
	_go_fuzz_dep_.CoverTab[96579]++
															return sum / float64(len(values))
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:394
	// _ = "end of CoverTab[96579]"
}

// A uniform sample using Vitter's Algorithm R.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:397
//
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:397
// <http://www.cs.umd.edu/~samir/498/vitter.pdf>
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:400
type UniformSample struct {
	count		int64
	mutex		sync.Mutex
	reservoirSize	int
	values		[]int64
}

// NewUniformSample constructs a new uniform sample with the given reservoir
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:407
// size.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:409
func NewUniformSample(reservoirSize int) Sample {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:409
	_go_fuzz_dep_.CoverTab[96583]++
															if UseNilMetrics {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:410
		_go_fuzz_dep_.CoverTab[96585]++
																return NilSample{}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:411
		// _ = "end of CoverTab[96585]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:412
		_go_fuzz_dep_.CoverTab[96586]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:412
		// _ = "end of CoverTab[96586]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:412
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:412
	// _ = "end of CoverTab[96583]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:412
	_go_fuzz_dep_.CoverTab[96584]++
															return &UniformSample{
		reservoirSize:	reservoirSize,
		values:		make([]int64, 0, reservoirSize),
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:416
	// _ = "end of CoverTab[96584]"
}

// Clear clears all samples.
func (s *UniformSample) Clear() {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:420
	_go_fuzz_dep_.CoverTab[96587]++
															s.mutex.Lock()
															defer s.mutex.Unlock()
															s.count = 0
															s.values = make([]int64, 0, s.reservoirSize)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:424
	// _ = "end of CoverTab[96587]"
}

// Count returns the number of samples recorded, which may exceed the
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:427
// reservoir size.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:429
func (s *UniformSample) Count() int64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:429
	_go_fuzz_dep_.CoverTab[96588]++
															s.mutex.Lock()
															defer s.mutex.Unlock()
															return s.count
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:432
	// _ = "end of CoverTab[96588]"
}

// Max returns the maximum value in the sample, which may not be the maximum
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:435
// value ever to be part of the sample.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:437
func (s *UniformSample) Max() int64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:437
	_go_fuzz_dep_.CoverTab[96589]++
															s.mutex.Lock()
															defer s.mutex.Unlock()
															return SampleMax(s.values)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:440
	// _ = "end of CoverTab[96589]"
}

// Mean returns the mean of the values in the sample.
func (s *UniformSample) Mean() float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:444
	_go_fuzz_dep_.CoverTab[96590]++
															s.mutex.Lock()
															defer s.mutex.Unlock()
															return SampleMean(s.values)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:447
	// _ = "end of CoverTab[96590]"
}

// Min returns the minimum value in the sample, which may not be the minimum
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:450
// value ever to be part of the sample.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:452
func (s *UniformSample) Min() int64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:452
	_go_fuzz_dep_.CoverTab[96591]++
															s.mutex.Lock()
															defer s.mutex.Unlock()
															return SampleMin(s.values)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:455
	// _ = "end of CoverTab[96591]"
}

// Percentile returns an arbitrary percentile of values in the sample.
func (s *UniformSample) Percentile(p float64) float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:459
	_go_fuzz_dep_.CoverTab[96592]++
															s.mutex.Lock()
															defer s.mutex.Unlock()
															return SamplePercentile(s.values, p)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:462
	// _ = "end of CoverTab[96592]"
}

// Percentiles returns a slice of arbitrary percentiles of values in the
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:465
// sample.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:467
func (s *UniformSample) Percentiles(ps []float64) []float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:467
	_go_fuzz_dep_.CoverTab[96593]++
															s.mutex.Lock()
															defer s.mutex.Unlock()
															return SamplePercentiles(s.values, ps)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:470
	// _ = "end of CoverTab[96593]"
}

// Size returns the size of the sample, which is at most the reservoir size.
func (s *UniformSample) Size() int {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:474
	_go_fuzz_dep_.CoverTab[96594]++
															s.mutex.Lock()
															defer s.mutex.Unlock()
															return len(s.values)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:477
	// _ = "end of CoverTab[96594]"
}

// Snapshot returns a read-only copy of the sample.
func (s *UniformSample) Snapshot() Sample {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:481
	_go_fuzz_dep_.CoverTab[96595]++
															s.mutex.Lock()
															defer s.mutex.Unlock()
															values := make([]int64, len(s.values))
															copy(values, s.values)
															return &SampleSnapshot{
		count:	s.count,
		values:	values,
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:489
	// _ = "end of CoverTab[96595]"
}

// StdDev returns the standard deviation of the values in the sample.
func (s *UniformSample) StdDev() float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:493
	_go_fuzz_dep_.CoverTab[96596]++
															s.mutex.Lock()
															defer s.mutex.Unlock()
															return SampleStdDev(s.values)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:496
	// _ = "end of CoverTab[96596]"
}

// Sum returns the sum of the values in the sample.
func (s *UniformSample) Sum() int64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:500
	_go_fuzz_dep_.CoverTab[96597]++
															s.mutex.Lock()
															defer s.mutex.Unlock()
															return SampleSum(s.values)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:503
	// _ = "end of CoverTab[96597]"
}

// Update samples a new value.
func (s *UniformSample) Update(v int64) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:507
	_go_fuzz_dep_.CoverTab[96598]++
															s.mutex.Lock()
															defer s.mutex.Unlock()
															s.count++
															if len(s.values) < s.reservoirSize {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:511
		_go_fuzz_dep_.CoverTab[96599]++
																s.values = append(s.values, v)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:512
		// _ = "end of CoverTab[96599]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:513
		_go_fuzz_dep_.CoverTab[96600]++
																r := rand.Int63n(s.count)
																if r < int64(len(s.values)) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:515
			_go_fuzz_dep_.CoverTab[96601]++
																	s.values[int(r)] = v
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:516
			// _ = "end of CoverTab[96601]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:517
			_go_fuzz_dep_.CoverTab[96602]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:517
			// _ = "end of CoverTab[96602]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:517
		}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:517
		// _ = "end of CoverTab[96600]"
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:518
	// _ = "end of CoverTab[96598]"
}

// Values returns a copy of the values in the sample.
func (s *UniformSample) Values() []int64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:522
	_go_fuzz_dep_.CoverTab[96603]++
															s.mutex.Lock()
															defer s.mutex.Unlock()
															values := make([]int64, len(s.values))
															copy(values, s.values)
															return values
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:527
	// _ = "end of CoverTab[96603]"
}

// Variance returns the variance of the values in the sample.
func (s *UniformSample) Variance() float64 {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:531
	_go_fuzz_dep_.CoverTab[96604]++
															s.mutex.Lock()
															defer s.mutex.Unlock()
															return SampleVariance(s.values)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:534
	// _ = "end of CoverTab[96604]"
}

// expDecaySample represents an individual sample in a heap.
type expDecaySample struct {
	k	float64
	v	int64
}

func newExpDecaySampleHeap(reservoirSize int) *expDecaySampleHeap {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:543
	_go_fuzz_dep_.CoverTab[96605]++
															return &expDecaySampleHeap{make([]expDecaySample, 0, reservoirSize)}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:544
	// _ = "end of CoverTab[96605]"
}

// expDecaySampleHeap is a min-heap of expDecaySamples.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:547
// The internal implementation is copied from the standard library's container/heap
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:549
type expDecaySampleHeap struct {
	s []expDecaySample
}

func (h *expDecaySampleHeap) Clear() {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:553
	_go_fuzz_dep_.CoverTab[96606]++
															h.s = h.s[:0]
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:554
	// _ = "end of CoverTab[96606]"
}

func (h *expDecaySampleHeap) Push(s expDecaySample) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:557
	_go_fuzz_dep_.CoverTab[96607]++
															n := len(h.s)
															h.s = h.s[0 : n+1]
															h.s[n] = s
															h.up(n)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:561
	// _ = "end of CoverTab[96607]"
}

func (h *expDecaySampleHeap) Pop() expDecaySample {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:564
	_go_fuzz_dep_.CoverTab[96608]++
															n := len(h.s) - 1
															h.s[0], h.s[n] = h.s[n], h.s[0]
															h.down(0, n)

															n = len(h.s)
															s := h.s[n-1]
															h.s = h.s[0 : n-1]
															return s
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:572
	// _ = "end of CoverTab[96608]"
}

func (h *expDecaySampleHeap) Size() int {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:575
	_go_fuzz_dep_.CoverTab[96609]++
															return len(h.s)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:576
	// _ = "end of CoverTab[96609]"
}

func (h *expDecaySampleHeap) Values() []expDecaySample {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:579
	_go_fuzz_dep_.CoverTab[96610]++
															return h.s
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:580
	// _ = "end of CoverTab[96610]"
}

func (h *expDecaySampleHeap) up(j int) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:583
	_go_fuzz_dep_.CoverTab[96611]++
															for {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:584
		_go_fuzz_dep_.CoverTab[96612]++
																i := (j - 1) / 2
																if i == j || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:586
			_go_fuzz_dep_.CoverTab[96614]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:586
			return !(h.s[j].k < h.s[i].k)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:586
			// _ = "end of CoverTab[96614]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:586
		}() {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:586
			_go_fuzz_dep_.CoverTab[96615]++
																	break
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:587
			// _ = "end of CoverTab[96615]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:588
			_go_fuzz_dep_.CoverTab[96616]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:588
			// _ = "end of CoverTab[96616]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:588
		}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:588
		// _ = "end of CoverTab[96612]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:588
		_go_fuzz_dep_.CoverTab[96613]++
																h.s[i], h.s[j] = h.s[j], h.s[i]
																j = i
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:590
		// _ = "end of CoverTab[96613]"
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:591
	// _ = "end of CoverTab[96611]"
}

func (h *expDecaySampleHeap) down(i, n int) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:594
	_go_fuzz_dep_.CoverTab[96617]++
															for {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:595
		_go_fuzz_dep_.CoverTab[96618]++
																j1 := 2*i + 1
																if j1 >= n || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:597
			_go_fuzz_dep_.CoverTab[96622]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:597
			return j1 < 0
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:597
			// _ = "end of CoverTab[96622]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:597
		}() {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:597
			_go_fuzz_dep_.CoverTab[96623]++
																	break
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:598
			// _ = "end of CoverTab[96623]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:599
			_go_fuzz_dep_.CoverTab[96624]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:599
			// _ = "end of CoverTab[96624]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:599
		}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:599
		// _ = "end of CoverTab[96618]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:599
		_go_fuzz_dep_.CoverTab[96619]++
																j := j1
																if j2 := j1 + 1; j2 < n && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:601
			_go_fuzz_dep_.CoverTab[96625]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:601
			return !(h.s[j1].k < h.s[j2].k)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:601
			// _ = "end of CoverTab[96625]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:601
		}() {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:601
			_go_fuzz_dep_.CoverTab[96626]++
																	j = j2
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:602
			// _ = "end of CoverTab[96626]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:603
			_go_fuzz_dep_.CoverTab[96627]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:603
			// _ = "end of CoverTab[96627]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:603
		}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:603
		// _ = "end of CoverTab[96619]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:603
		_go_fuzz_dep_.CoverTab[96620]++
																if !(h.s[j].k < h.s[i].k) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:604
			_go_fuzz_dep_.CoverTab[96628]++
																	break
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:605
			// _ = "end of CoverTab[96628]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:606
			_go_fuzz_dep_.CoverTab[96629]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:606
			// _ = "end of CoverTab[96629]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:606
		}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:606
		// _ = "end of CoverTab[96620]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:606
		_go_fuzz_dep_.CoverTab[96621]++
																h.s[i], h.s[j] = h.s[j], h.s[i]
																i = j
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:608
		// _ = "end of CoverTab[96621]"
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:609
	// _ = "end of CoverTab[96617]"
}

type int64Slice []int64

func (p int64Slice) Len() int {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:614
	_go_fuzz_dep_.CoverTab[96630]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:614
	return len(p)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:614
	// _ = "end of CoverTab[96630]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:614
}
func (p int64Slice) Less(i, j int) bool {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:615
	_go_fuzz_dep_.CoverTab[96631]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:615
	return p[i] < p[j]
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:615
	// _ = "end of CoverTab[96631]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:615
}
func (p int64Slice) Swap(i, j int) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:616
	_go_fuzz_dep_.CoverTab[96632]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:616
	p[i], p[j] = p[j], p[i]
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:616
	// _ = "end of CoverTab[96632]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:616
}

//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:616
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/sample.go:616
var _ = _go_fuzz_dep_.CoverTab
