// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:5
package trace

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:5
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:5
)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:5
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:5
)

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:9
import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"math"
	"sync"

	"golang.org/x/net/internal/timeseries"
)

const (
	bucketCount = 38
)

// histogram keeps counts of values in buckets that are spaced
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:24
// out in powers of 2: 0-1, 2-3, 4-7...
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:24
// histogram implements timeseries.Observable
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:27
type histogram struct {
	sum		int64	// running total of measurements
	sumOfSquares	float64	// square of running total
	buckets		[]int64	// bucketed values for histogram
	value		int	// holds a single value as an optimization
	valueCount	int64	// number of values recorded for single value
}

// addMeasurement records a value measurement observation to the histogram.
func (h *histogram) addMeasurement(value int64) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:36
	_go_fuzz_dep_.CoverTab[45345]++

											h.sum += value
											h.sumOfSquares += float64(value) * float64(value)

											bucketIndex := getBucket(value)

											if h.valueCount == 0 || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:43
		_go_fuzz_dep_.CoverTab[45346]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:43
		return (h.valueCount > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:43
			_go_fuzz_dep_.CoverTab[45347]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:43
			return h.value == bucketIndex
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:43
			// _ = "end of CoverTab[45347]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:43
		}())
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:43
		// _ = "end of CoverTab[45346]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:43
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:43
		_go_fuzz_dep_.CoverTab[45348]++
												h.value = bucketIndex
												h.valueCount++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:45
		// _ = "end of CoverTab[45348]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:46
		_go_fuzz_dep_.CoverTab[45349]++
												h.allocateBuckets()
												h.buckets[bucketIndex]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:48
		// _ = "end of CoverTab[45349]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:49
	// _ = "end of CoverTab[45345]"
}

func (h *histogram) allocateBuckets() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:52
	_go_fuzz_dep_.CoverTab[45350]++
											if h.buckets == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:53
		_go_fuzz_dep_.CoverTab[45351]++
												h.buckets = make([]int64, bucketCount)
												h.buckets[h.value] = h.valueCount
												h.value = 0
												h.valueCount = -1
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:57
		// _ = "end of CoverTab[45351]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:58
		_go_fuzz_dep_.CoverTab[45352]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:58
		// _ = "end of CoverTab[45352]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:58
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:58
	// _ = "end of CoverTab[45350]"
}

func log2(i int64) int {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:61
	_go_fuzz_dep_.CoverTab[45353]++
											n := 0
											for ; i >= 0x100; i >>= 8 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:63
		_go_fuzz_dep_.CoverTab[45356]++
												n += 8
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:64
		// _ = "end of CoverTab[45356]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:65
	// _ = "end of CoverTab[45353]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:65
	_go_fuzz_dep_.CoverTab[45354]++
											for ; i > 0; i >>= 1 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:66
		_go_fuzz_dep_.CoverTab[45357]++
												n += 1
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:67
		// _ = "end of CoverTab[45357]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:68
	// _ = "end of CoverTab[45354]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:68
	_go_fuzz_dep_.CoverTab[45355]++
											return n
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:69
	// _ = "end of CoverTab[45355]"
}

func getBucket(i int64) (index int) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:72
	_go_fuzz_dep_.CoverTab[45358]++
											index = log2(i) - 1
											if index < 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:74
		_go_fuzz_dep_.CoverTab[45361]++
												index = 0
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:75
		// _ = "end of CoverTab[45361]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:76
		_go_fuzz_dep_.CoverTab[45362]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:76
		// _ = "end of CoverTab[45362]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:76
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:76
	// _ = "end of CoverTab[45358]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:76
	_go_fuzz_dep_.CoverTab[45359]++
											if index >= bucketCount {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:77
		_go_fuzz_dep_.CoverTab[45363]++
												index = bucketCount - 1
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:78
		// _ = "end of CoverTab[45363]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:79
		_go_fuzz_dep_.CoverTab[45364]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:79
		// _ = "end of CoverTab[45364]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:79
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:79
	// _ = "end of CoverTab[45359]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:79
	_go_fuzz_dep_.CoverTab[45360]++
											return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:80
	// _ = "end of CoverTab[45360]"
}

// Total returns the number of recorded observations.
func (h *histogram) total() (total int64) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:84
	_go_fuzz_dep_.CoverTab[45365]++
											if h.valueCount >= 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:85
		_go_fuzz_dep_.CoverTab[45368]++
												total = h.valueCount
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:86
		// _ = "end of CoverTab[45368]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:87
		_go_fuzz_dep_.CoverTab[45369]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:87
		// _ = "end of CoverTab[45369]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:87
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:87
	// _ = "end of CoverTab[45365]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:87
	_go_fuzz_dep_.CoverTab[45366]++
											for _, val := range h.buckets {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:88
		_go_fuzz_dep_.CoverTab[45370]++
												total += int64(val)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:89
		// _ = "end of CoverTab[45370]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:90
	// _ = "end of CoverTab[45366]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:90
	_go_fuzz_dep_.CoverTab[45367]++
											return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:91
	// _ = "end of CoverTab[45367]"
}

// Average returns the average value of recorded observations.
func (h *histogram) average() float64 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:95
	_go_fuzz_dep_.CoverTab[45371]++
											t := h.total()
											if t == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:97
		_go_fuzz_dep_.CoverTab[45373]++
												return 0
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:98
		// _ = "end of CoverTab[45373]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:99
		_go_fuzz_dep_.CoverTab[45374]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:99
		// _ = "end of CoverTab[45374]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:99
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:99
	// _ = "end of CoverTab[45371]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:99
	_go_fuzz_dep_.CoverTab[45372]++
											return float64(h.sum) / float64(t)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:100
	// _ = "end of CoverTab[45372]"
}

// Variance returns the variance of recorded observations.
func (h *histogram) variance() float64 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:104
	_go_fuzz_dep_.CoverTab[45375]++
											t := float64(h.total())
											if t == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:106
		_go_fuzz_dep_.CoverTab[45377]++
												return 0
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:107
		// _ = "end of CoverTab[45377]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:108
		_go_fuzz_dep_.CoverTab[45378]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:108
		// _ = "end of CoverTab[45378]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:108
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:108
	// _ = "end of CoverTab[45375]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:108
	_go_fuzz_dep_.CoverTab[45376]++
											s := float64(h.sum) / t
											return h.sumOfSquares/t - s*s
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:110
	// _ = "end of CoverTab[45376]"
}

// StandardDeviation returns the standard deviation of recorded observations.
func (h *histogram) standardDeviation() float64 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:114
	_go_fuzz_dep_.CoverTab[45379]++
											return math.Sqrt(h.variance())
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:115
	// _ = "end of CoverTab[45379]"
}

// PercentileBoundary estimates the value that the given fraction of recorded
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:118
// observations are less than.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:120
func (h *histogram) percentileBoundary(percentile float64) int64 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:120
	_go_fuzz_dep_.CoverTab[45380]++
											total := h.total()

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:124
	if total == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:124
		_go_fuzz_dep_.CoverTab[45383]++
												return 0
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:125
		// _ = "end of CoverTab[45383]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:126
		_go_fuzz_dep_.CoverTab[45384]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:126
		if total == 1 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:126
			_go_fuzz_dep_.CoverTab[45385]++
													return int64(h.average())
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:127
			// _ = "end of CoverTab[45385]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:128
			_go_fuzz_dep_.CoverTab[45386]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:128
			// _ = "end of CoverTab[45386]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:128
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:128
		// _ = "end of CoverTab[45384]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:128
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:128
	// _ = "end of CoverTab[45380]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:128
	_go_fuzz_dep_.CoverTab[45381]++

											percentOfTotal := round(float64(total) * percentile)
											var runningTotal int64

											for i := range h.buckets {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:133
		_go_fuzz_dep_.CoverTab[45387]++
												value := h.buckets[i]
												runningTotal += value
												if runningTotal == percentOfTotal {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:136
			_go_fuzz_dep_.CoverTab[45388]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:142
			j := uint8(i + 1)
			min := bucketBoundary(j)
			if runningTotal < total {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:144
				_go_fuzz_dep_.CoverTab[45390]++
														for h.buckets[j] == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:145
					_go_fuzz_dep_.CoverTab[45391]++
															j++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:146
					// _ = "end of CoverTab[45391]"
				}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:147
				// _ = "end of CoverTab[45390]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:148
				_go_fuzz_dep_.CoverTab[45392]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:148
				// _ = "end of CoverTab[45392]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:148
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:148
			// _ = "end of CoverTab[45388]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:148
			_go_fuzz_dep_.CoverTab[45389]++
													max := bucketBoundary(j)
													return min + round(float64(max-min)/2)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:150
			// _ = "end of CoverTab[45389]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:151
			_go_fuzz_dep_.CoverTab[45393]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:151
			if runningTotal > percentOfTotal {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:151
				_go_fuzz_dep_.CoverTab[45394]++

														delta := runningTotal - percentOfTotal
														percentBucket := float64(value-delta) / float64(value)
														bucketMin := bucketBoundary(uint8(i))
														nextBucketMin := bucketBoundary(uint8(i + 1))
														bucketSize := nextBucketMin - bucketMin
														return bucketMin + round(percentBucket*float64(bucketSize))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:158
				// _ = "end of CoverTab[45394]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:159
				_go_fuzz_dep_.CoverTab[45395]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:159
				// _ = "end of CoverTab[45395]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:159
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:159
			// _ = "end of CoverTab[45393]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:159
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:159
		// _ = "end of CoverTab[45387]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:160
	// _ = "end of CoverTab[45381]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:160
	_go_fuzz_dep_.CoverTab[45382]++
											return bucketBoundary(bucketCount - 1)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:161
	// _ = "end of CoverTab[45382]"
}

// Median returns the estimated median of the observed values.
func (h *histogram) median() int64 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:165
	_go_fuzz_dep_.CoverTab[45396]++
											return h.percentileBoundary(0.5)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:166
	// _ = "end of CoverTab[45396]"
}

// Add adds other to h.
func (h *histogram) Add(other timeseries.Observable) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:170
	_go_fuzz_dep_.CoverTab[45397]++
											o := other.(*histogram)
											if o.valueCount == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:172
		_go_fuzz_dep_.CoverTab[45399]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:172
		// _ = "end of CoverTab[45399]"

	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:174
		_go_fuzz_dep_.CoverTab[45400]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:174
		if h.valueCount >= 0 && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:174
			_go_fuzz_dep_.CoverTab[45401]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:174
			return o.valueCount > 0
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:174
			// _ = "end of CoverTab[45401]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:174
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:174
			_go_fuzz_dep_.CoverTab[45402]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:174
			return h.value == o.value
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:174
			// _ = "end of CoverTab[45402]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:174
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:174
			_go_fuzz_dep_.CoverTab[45403]++

													h.valueCount += o.valueCount
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:176
			// _ = "end of CoverTab[45403]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:177
			_go_fuzz_dep_.CoverTab[45404]++

													h.allocateBuckets()
													if o.valueCount >= 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:180
				_go_fuzz_dep_.CoverTab[45405]++
														h.buckets[o.value] += o.valueCount
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:181
				// _ = "end of CoverTab[45405]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:182
				_go_fuzz_dep_.CoverTab[45406]++
														for i := range h.buckets {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:183
					_go_fuzz_dep_.CoverTab[45407]++
															h.buckets[i] += o.buckets[i]
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:184
					// _ = "end of CoverTab[45407]"
				}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:185
				// _ = "end of CoverTab[45406]"
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:186
			// _ = "end of CoverTab[45404]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:187
		// _ = "end of CoverTab[45400]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:187
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:187
	// _ = "end of CoverTab[45397]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:187
	_go_fuzz_dep_.CoverTab[45398]++
											h.sumOfSquares += o.sumOfSquares
											h.sum += o.sum
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:189
	// _ = "end of CoverTab[45398]"
}

// Clear resets the histogram to an empty state, removing all observed values.
func (h *histogram) Clear() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:193
	_go_fuzz_dep_.CoverTab[45408]++
											h.buckets = nil
											h.value = 0
											h.valueCount = 0
											h.sum = 0
											h.sumOfSquares = 0
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:198
	// _ = "end of CoverTab[45408]"
}

// CopyFrom copies from other, which must be a *histogram, into h.
func (h *histogram) CopyFrom(other timeseries.Observable) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:202
	_go_fuzz_dep_.CoverTab[45409]++
											o := other.(*histogram)
											if o.valueCount == -1 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:204
		_go_fuzz_dep_.CoverTab[45411]++
												h.allocateBuckets()
												copy(h.buckets, o.buckets)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:206
		// _ = "end of CoverTab[45411]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:207
		_go_fuzz_dep_.CoverTab[45412]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:207
		// _ = "end of CoverTab[45412]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:207
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:207
	// _ = "end of CoverTab[45409]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:207
	_go_fuzz_dep_.CoverTab[45410]++
											h.sum = o.sum
											h.sumOfSquares = o.sumOfSquares
											h.value = o.value
											h.valueCount = o.valueCount
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:211
	// _ = "end of CoverTab[45410]"
}

// Multiply scales the histogram by the specified ratio.
func (h *histogram) Multiply(ratio float64) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:215
	_go_fuzz_dep_.CoverTab[45413]++
											if h.valueCount == -1 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:216
		_go_fuzz_dep_.CoverTab[45415]++
												for i := range h.buckets {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:217
			_go_fuzz_dep_.CoverTab[45416]++
													h.buckets[i] = int64(float64(h.buckets[i]) * ratio)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:218
			// _ = "end of CoverTab[45416]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:219
		// _ = "end of CoverTab[45415]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:220
		_go_fuzz_dep_.CoverTab[45417]++
												h.valueCount = int64(float64(h.valueCount) * ratio)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:221
		// _ = "end of CoverTab[45417]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:222
	// _ = "end of CoverTab[45413]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:222
	_go_fuzz_dep_.CoverTab[45414]++
											h.sum = int64(float64(h.sum) * ratio)
											h.sumOfSquares = h.sumOfSquares * ratio
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:224
	// _ = "end of CoverTab[45414]"
}

// New creates a new histogram.
func (h *histogram) New() timeseries.Observable {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:228
	_go_fuzz_dep_.CoverTab[45418]++
											r := new(histogram)
											r.Clear()
											return r
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:231
	// _ = "end of CoverTab[45418]"
}

func (h *histogram) String() string {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:234
	_go_fuzz_dep_.CoverTab[45419]++
											return fmt.Sprintf("%d, %f, %d, %d, %v",
		h.sum, h.sumOfSquares, h.value, h.valueCount, h.buckets)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:236
	// _ = "end of CoverTab[45419]"
}

// round returns the closest int64 to the argument
func round(in float64) int64 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:240
	_go_fuzz_dep_.CoverTab[45420]++
											return int64(math.Floor(in + 0.5))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:241
	// _ = "end of CoverTab[45420]"
}

// bucketBoundary returns the first value in the bucket.
func bucketBoundary(bucket uint8) int64 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:245
	_go_fuzz_dep_.CoverTab[45421]++
											if bucket == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:246
		_go_fuzz_dep_.CoverTab[45423]++
												return 0
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:247
		// _ = "end of CoverTab[45423]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:248
		_go_fuzz_dep_.CoverTab[45424]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:248
		// _ = "end of CoverTab[45424]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:248
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:248
	// _ = "end of CoverTab[45421]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:248
	_go_fuzz_dep_.CoverTab[45422]++
											return 1 << bucket
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:249
	// _ = "end of CoverTab[45422]"
}

// bucketData holds data about a specific bucket for use in distTmpl.
type bucketData struct {
	Lower, Upper		int64
	N			int64
	Pct, CumulativePct	float64
	GraphWidth		int
}

// data holds data about a Distribution for use in distTmpl.
type data struct {
	Buckets			[]*bucketData
	Count, Median		int64
	Mean, StandardDeviation	float64
}

// maxHTMLBarWidth is the maximum width of the HTML bar for visualizing buckets.
const maxHTMLBarWidth = 350.0

// newData returns data representing h for use in distTmpl.
func (h *histogram) newData() *data {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:271
	_go_fuzz_dep_.CoverTab[45425]++

											h.allocateBuckets()

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:276
	maxBucket := int64(0)
	for _, n := range h.buckets {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:277
		_go_fuzz_dep_.CoverTab[45429]++
												if n > maxBucket {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:278
			_go_fuzz_dep_.CoverTab[45430]++
													maxBucket = n
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:279
			// _ = "end of CoverTab[45430]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:280
			_go_fuzz_dep_.CoverTab[45431]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:280
			// _ = "end of CoverTab[45431]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:280
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:280
		// _ = "end of CoverTab[45429]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:281
	// _ = "end of CoverTab[45425]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:281
	_go_fuzz_dep_.CoverTab[45426]++
											total := h.total()
											barsizeMult := maxHTMLBarWidth / float64(maxBucket)
											var pctMult float64
											if total == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:285
		_go_fuzz_dep_.CoverTab[45432]++
												pctMult = 1.0
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:286
		// _ = "end of CoverTab[45432]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:287
		_go_fuzz_dep_.CoverTab[45433]++
												pctMult = 100.0 / float64(total)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:288
		// _ = "end of CoverTab[45433]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:289
	// _ = "end of CoverTab[45426]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:289
	_go_fuzz_dep_.CoverTab[45427]++

											buckets := make([]*bucketData, len(h.buckets))
											runningTotal := int64(0)
											for i, n := range h.buckets {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:293
		_go_fuzz_dep_.CoverTab[45434]++
												if n == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:294
			_go_fuzz_dep_.CoverTab[45437]++
													continue
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:295
			// _ = "end of CoverTab[45437]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:296
			_go_fuzz_dep_.CoverTab[45438]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:296
			// _ = "end of CoverTab[45438]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:296
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:296
		// _ = "end of CoverTab[45434]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:296
		_go_fuzz_dep_.CoverTab[45435]++
												runningTotal += n
												var upperBound int64
												if i < bucketCount-1 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:299
			_go_fuzz_dep_.CoverTab[45439]++
													upperBound = bucketBoundary(uint8(i + 1))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:300
			// _ = "end of CoverTab[45439]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:301
			_go_fuzz_dep_.CoverTab[45440]++
													upperBound = math.MaxInt64
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:302
			// _ = "end of CoverTab[45440]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:303
		// _ = "end of CoverTab[45435]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:303
		_go_fuzz_dep_.CoverTab[45436]++
												buckets[i] = &bucketData{
			Lower:		bucketBoundary(uint8(i)),
			Upper:		upperBound,
			N:		n,
			Pct:		float64(n) * pctMult,
			CumulativePct:	float64(runningTotal) * pctMult,
			GraphWidth:	int(float64(n) * barsizeMult),
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:311
		// _ = "end of CoverTab[45436]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:312
	// _ = "end of CoverTab[45427]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:312
	_go_fuzz_dep_.CoverTab[45428]++
											return &data{
		Buckets:		buckets,
		Count:			total,
		Median:			h.median(),
		Mean:			h.average(),
		StandardDeviation:	h.standardDeviation(),
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:319
	// _ = "end of CoverTab[45428]"
}

func (h *histogram) html() template.HTML {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:322
	_go_fuzz_dep_.CoverTab[45441]++
											buf := new(bytes.Buffer)
											if err := distTmpl().Execute(buf, h.newData()); err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:324
		_go_fuzz_dep_.CoverTab[45443]++
												buf.Reset()
												log.Printf("net/trace: couldn't execute template: %v", err)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:326
		// _ = "end of CoverTab[45443]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:327
		_go_fuzz_dep_.CoverTab[45444]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:327
		// _ = "end of CoverTab[45444]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:327
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:327
	// _ = "end of CoverTab[45441]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:327
	_go_fuzz_dep_.CoverTab[45442]++
											return template.HTML(buf.String())
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:328
	// _ = "end of CoverTab[45442]"
}

var distTmplCache *template.Template
var distTmplOnce sync.Once

func distTmpl() *template.Template {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:334
	_go_fuzz_dep_.CoverTab[45445]++
											distTmplOnce.Do(func() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:335
		_go_fuzz_dep_.CoverTab[45447]++

												distTmplCache = template.Must(template.New("distTmpl").Parse(`
<table>
<tr>
    <td style="padding:0.25em">Count: {{.Count}}</td>
    <td style="padding:0.25em">Mean: {{printf "%.0f" .Mean}}</td>
    <td style="padding:0.25em">StdDev: {{printf "%.0f" .StandardDeviation}}</td>
    <td style="padding:0.25em">Median: {{.Median}}</td>
</tr>
</table>
<hr>
<table>
{{range $b := .Buckets}}
{{if $b}}
  <tr>
    <td style="padding:0 0 0 0.25em">[</td>
    <td style="text-align:right;padding:0 0.25em">{{.Lower}},</td>
    <td style="text-align:right;padding:0 0.25em">{{.Upper}})</td>
    <td style="text-align:right;padding:0 0.25em">{{.N}}</td>
    <td style="text-align:right;padding:0 0.25em">{{printf "%#.3f" .Pct}}%</td>
    <td style="text-align:right;padding:0 0.25em">{{printf "%#.3f" .CumulativePct}}%</td>
    <td><div style="background-color: blue; height: 1em; width: {{.GraphWidth}};"></div></td>
  </tr>
{{end}}
{{end}}
</table>
`))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:362
		// _ = "end of CoverTab[45447]"
	})
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:363
	// _ = "end of CoverTab[45445]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:363
	_go_fuzz_dep_.CoverTab[45446]++
											return distTmplCache
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:364
	// _ = "end of CoverTab[45446]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:365
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/histogram.go:365
var _ = _go_fuzz_dep_.CoverTab
