// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:5
// Package timeseries implements a time series structure for stats collection.
package timeseries

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:6
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:6
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:6
)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:6
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:6
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:6
)

import (
	"fmt"
	"log"
	"time"
)

const (
	timeSeriesNumBuckets		= 64
	minuteHourSeriesNumBuckets	= 60
)

var timeSeriesResolutions = []time.Duration{
	1 * time.Second,
	10 * time.Second,
	1 * time.Minute,
	10 * time.Minute,
	1 * time.Hour,
	6 * time.Hour,
	24 * time.Hour,
	7 * 24 * time.Hour,
	4 * 7 * 24 * time.Hour,
	16 * 7 * 24 * time.Hour,
}

var minuteHourSeriesResolutions = []time.Duration{
	1 * time.Second,
	1 * time.Minute,
}

// An Observable is a kind of data that can be aggregated in a time series.
type Observable interface {
	Multiply(ratio float64)		// Multiplies the data in self by a given ratio
	Add(other Observable)		// Adds the data from a different observation to self
	Clear()				// Clears the observation so it can be reused.
	CopyFrom(other Observable)	// Copies the contents of a given observation to self
}

// Float attaches the methods of Observable to a float64.
type Float float64

// NewFloat returns a Float.
func NewFloat() Observable {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:49
	_go_fuzz_dep_.CoverTab[28666]++
													f := Float(0)
													return &f
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:51
	// _ = "end of CoverTab[28666]"
}

// String returns the float as a string.
func (f *Float) String() string {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:55
	_go_fuzz_dep_.CoverTab[28667]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:55
	return fmt.Sprintf("%g", f.Value())
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:55
	// _ = "end of CoverTab[28667]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:55
}

// Value returns the float's value.
func (f *Float) Value() float64 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:58
	_go_fuzz_dep_.CoverTab[28668]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:58
	return float64(*f)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:58
	// _ = "end of CoverTab[28668]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:58
}

func (f *Float) Multiply(ratio float64) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:60
	_go_fuzz_dep_.CoverTab[28669]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:60
	*f *= Float(ratio)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:60
	// _ = "end of CoverTab[28669]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:60
}

func (f *Float) Add(other Observable) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:62
	_go_fuzz_dep_.CoverTab[28670]++
													o := other.(*Float)
													*f += *o
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:64
	// _ = "end of CoverTab[28670]"
}

func (f *Float) Clear()	{ _go_fuzz_dep_.CoverTab[28671]++; *f = 0; // _ = "end of CoverTab[28671]" }

func (f *Float) CopyFrom(other Observable) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:69
	_go_fuzz_dep_.CoverTab[28672]++
													o := other.(*Float)
													*f = *o
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:71
	// _ = "end of CoverTab[28672]"
}

// A Clock tells the current time.
type Clock interface {
	Time() time.Time
}

type defaultClock int

var defaultClockInstance defaultClock

func (defaultClock) Time() time.Time {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:83
	_go_fuzz_dep_.CoverTab[28673]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:83
	return time.Now()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:83
	// _ = "end of CoverTab[28673]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:83
}

// Information kept per level. Each level consists of a circular list of
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:85
// observations. The start of the level may be derived from end and the
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:85
// len(buckets) * sizeInMillis.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:88
type tsLevel struct {
	oldest		int			// index to oldest bucketed Observable
	newest		int			// index to newest bucketed Observable
	end		time.Time		// end timestamp for this level
	size		time.Duration		// duration of the bucketed Observable
	buckets		[]Observable		// collections of observations
	provider	func() Observable	// used for creating new Observable
}

func (l *tsLevel) Clear() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:97
	_go_fuzz_dep_.CoverTab[28674]++
													l.oldest = 0
													l.newest = len(l.buckets) - 1
													l.end = time.Time{}
													for i := range l.buckets {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:101
		_go_fuzz_dep_.CoverTab[28675]++
														if l.buckets[i] != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:102
			_go_fuzz_dep_.CoverTab[28676]++
															l.buckets[i].Clear()
															l.buckets[i] = nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:104
			// _ = "end of CoverTab[28676]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:105
			_go_fuzz_dep_.CoverTab[28677]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:105
			// _ = "end of CoverTab[28677]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:105
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:105
		// _ = "end of CoverTab[28675]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:106
	// _ = "end of CoverTab[28674]"
}

func (l *tsLevel) InitLevel(size time.Duration, numBuckets int, f func() Observable) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:109
	_go_fuzz_dep_.CoverTab[28678]++
													l.size = size
													l.provider = f
													l.buckets = make([]Observable, numBuckets)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:112
	// _ = "end of CoverTab[28678]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:120
// Each level is represented by a sequence of buckets. Each bucket spans an
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:120
// interval equal to the resolution of the level. New observations are added
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:120
// to the last bucket.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:123
type timeSeries struct {
	provider	func() Observable	// make more Observable
	numBuckets	int			// number of buckets in each level
	levels		[]*tsLevel		// levels of bucketed Observable
	lastAdd		time.Time		// time of last Observable tracked
	total		Observable		// convenient aggregation of all Observable
	clock		Clock			// Clock for getting current time
	pending		Observable		// observations not yet bucketed
	pendingTime	time.Time		// what time are we keeping in pending
	dirty		bool			// if there are pending observations
}

// init initializes a level according to the supplied criteria.
func (ts *timeSeries) init(resolutions []time.Duration, f func() Observable, numBuckets int, clock Clock) {
	ts.provider = f
	ts.numBuckets = numBuckets
	ts.clock = clock
	ts.levels = make([]*tsLevel, len(resolutions))

	for i := range resolutions {
		if i > 0 && resolutions[i-1] >= resolutions[i] {
			log.Print("timeseries: resolutions must be monotonically increasing")
			break
		}
		newLevel := new(tsLevel)
		newLevel.InitLevel(resolutions[i], ts.numBuckets, ts.provider)
		ts.levels[i] = newLevel
	}

	ts.Clear()
}

// Clear removes all observations from the time series.
func (ts *timeSeries) Clear() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:156
	_go_fuzz_dep_.CoverTab[28679]++
													ts.lastAdd = time.Time{}
													ts.total = ts.resetObservation(ts.total)
													ts.pending = ts.resetObservation(ts.pending)
													ts.pendingTime = time.Time{}
													ts.dirty = false

													for i := range ts.levels {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:163
		_go_fuzz_dep_.CoverTab[28680]++
														ts.levels[i].Clear()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:164
		// _ = "end of CoverTab[28680]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:165
	// _ = "end of CoverTab[28679]"
}

// Add records an observation at the current time.
func (ts *timeSeries) Add(observation Observable) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:169
	_go_fuzz_dep_.CoverTab[28681]++
													ts.AddWithTime(observation, ts.clock.Time())
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:170
	// _ = "end of CoverTab[28681]"
}

// AddWithTime records an observation at the specified time.
func (ts *timeSeries) AddWithTime(observation Observable, t time.Time) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:174
	_go_fuzz_dep_.CoverTab[28682]++

													smallBucketDuration := ts.levels[0].size

													if t.After(ts.lastAdd) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:178
		_go_fuzz_dep_.CoverTab[28684]++
														ts.lastAdd = t
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:179
		// _ = "end of CoverTab[28684]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:180
		_go_fuzz_dep_.CoverTab[28685]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:180
		// _ = "end of CoverTab[28685]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:180
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:180
	// _ = "end of CoverTab[28682]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:180
	_go_fuzz_dep_.CoverTab[28683]++

													if t.After(ts.pendingTime) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:182
		_go_fuzz_dep_.CoverTab[28686]++
														ts.advance(t)
														ts.mergePendingUpdates()
														ts.pendingTime = ts.levels[0].end
														ts.pending.CopyFrom(observation)
														ts.dirty = true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:187
		// _ = "end of CoverTab[28686]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:188
		_go_fuzz_dep_.CoverTab[28687]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:188
		if t.After(ts.pendingTime.Add(-1 * smallBucketDuration)) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:188
			_go_fuzz_dep_.CoverTab[28688]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:192
			ts.pending.Add(observation)
															ts.dirty = true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:193
			// _ = "end of CoverTab[28688]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:194
			_go_fuzz_dep_.CoverTab[28689]++
															ts.mergeValue(observation, t)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:195
			// _ = "end of CoverTab[28689]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:196
		// _ = "end of CoverTab[28687]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:196
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:196
	// _ = "end of CoverTab[28683]"
}

// mergeValue inserts the observation at the specified time in the past into all levels.
func (ts *timeSeries) mergeValue(observation Observable, t time.Time) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:200
	_go_fuzz_dep_.CoverTab[28690]++
													for _, level := range ts.levels {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:201
		_go_fuzz_dep_.CoverTab[28692]++
														index := (ts.numBuckets - 1) - int(level.end.Sub(t)/level.size)
														if 0 <= index && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:203
			_go_fuzz_dep_.CoverTab[28693]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:203
			return index < ts.numBuckets
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:203
			// _ = "end of CoverTab[28693]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:203
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:203
			_go_fuzz_dep_.CoverTab[28694]++
															bucketNumber := (level.oldest + index) % ts.numBuckets
															if level.buckets[bucketNumber] == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:205
				_go_fuzz_dep_.CoverTab[28696]++
																level.buckets[bucketNumber] = level.provider()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:206
				// _ = "end of CoverTab[28696]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:207
				_go_fuzz_dep_.CoverTab[28697]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:207
				// _ = "end of CoverTab[28697]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:207
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:207
			// _ = "end of CoverTab[28694]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:207
			_go_fuzz_dep_.CoverTab[28695]++
															level.buckets[bucketNumber].Add(observation)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:208
			// _ = "end of CoverTab[28695]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:209
			_go_fuzz_dep_.CoverTab[28698]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:209
			// _ = "end of CoverTab[28698]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:209
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:209
		// _ = "end of CoverTab[28692]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:210
	// _ = "end of CoverTab[28690]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:210
	_go_fuzz_dep_.CoverTab[28691]++
													ts.total.Add(observation)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:211
	// _ = "end of CoverTab[28691]"
}

// mergePendingUpdates applies the pending updates into all levels.
func (ts *timeSeries) mergePendingUpdates() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:215
	_go_fuzz_dep_.CoverTab[28699]++
													if ts.dirty {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:216
		_go_fuzz_dep_.CoverTab[28700]++
														ts.mergeValue(ts.pending, ts.pendingTime)
														ts.pending = ts.resetObservation(ts.pending)
														ts.dirty = false
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:219
		// _ = "end of CoverTab[28700]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:220
		_go_fuzz_dep_.CoverTab[28701]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:220
		// _ = "end of CoverTab[28701]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:220
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:220
	// _ = "end of CoverTab[28699]"
}

// advance cycles the buckets at each level until the latest bucket in
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:223
// each level can hold the time specified.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:225
func (ts *timeSeries) advance(t time.Time) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:225
	_go_fuzz_dep_.CoverTab[28702]++
													if !t.After(ts.levels[0].end) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:226
		_go_fuzz_dep_.CoverTab[28704]++
														return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:227
		// _ = "end of CoverTab[28704]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:228
		_go_fuzz_dep_.CoverTab[28705]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:228
		// _ = "end of CoverTab[28705]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:228
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:228
	// _ = "end of CoverTab[28702]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:228
	_go_fuzz_dep_.CoverTab[28703]++
													for i := 0; i < len(ts.levels); i++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:229
		_go_fuzz_dep_.CoverTab[28706]++
														level := ts.levels[i]
														if !level.end.Before(t) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:231
			_go_fuzz_dep_.CoverTab[28710]++
															break
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:232
			// _ = "end of CoverTab[28710]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:233
			_go_fuzz_dep_.CoverTab[28711]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:233
			// _ = "end of CoverTab[28711]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:233
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:233
		// _ = "end of CoverTab[28706]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:233
		_go_fuzz_dep_.CoverTab[28707]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:237
		if !t.Before(level.end.Add(level.size * time.Duration(ts.numBuckets))) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:237
			_go_fuzz_dep_.CoverTab[28712]++
															for _, b := range level.buckets {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:238
				_go_fuzz_dep_.CoverTab[28714]++
																ts.resetObservation(b)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:239
				// _ = "end of CoverTab[28714]"
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:240
			// _ = "end of CoverTab[28712]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:240
			_go_fuzz_dep_.CoverTab[28713]++
															level.end = time.Unix(0, (t.UnixNano()/level.size.Nanoseconds())*level.size.Nanoseconds())
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:241
			// _ = "end of CoverTab[28713]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:242
			_go_fuzz_dep_.CoverTab[28715]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:242
			// _ = "end of CoverTab[28715]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:242
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:242
		// _ = "end of CoverTab[28707]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:242
		_go_fuzz_dep_.CoverTab[28708]++

														for t.After(level.end) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:244
			_go_fuzz_dep_.CoverTab[28716]++
															level.end = level.end.Add(level.size)
															level.newest = level.oldest
															level.oldest = (level.oldest + 1) % ts.numBuckets
															ts.resetObservation(level.buckets[level.newest])
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:248
			// _ = "end of CoverTab[28716]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:249
		// _ = "end of CoverTab[28708]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:249
		_go_fuzz_dep_.CoverTab[28709]++

														t = level.end
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:251
		// _ = "end of CoverTab[28709]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:252
	// _ = "end of CoverTab[28703]"
}

// Latest returns the sum of the num latest buckets from the level.
func (ts *timeSeries) Latest(level, num int) Observable {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:256
	_go_fuzz_dep_.CoverTab[28717]++
													now := ts.clock.Time()
													if ts.levels[0].end.Before(now) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:258
		_go_fuzz_dep_.CoverTab[28720]++
														ts.advance(now)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:259
		// _ = "end of CoverTab[28720]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:260
		_go_fuzz_dep_.CoverTab[28721]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:260
		// _ = "end of CoverTab[28721]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:260
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:260
	// _ = "end of CoverTab[28717]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:260
	_go_fuzz_dep_.CoverTab[28718]++

													ts.mergePendingUpdates()

													result := ts.provider()
													l := ts.levels[level]
													index := l.newest

													for i := 0; i < num; i++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:268
		_go_fuzz_dep_.CoverTab[28722]++
														if l.buckets[index] != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:269
			_go_fuzz_dep_.CoverTab[28725]++
															result.Add(l.buckets[index])
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:270
			// _ = "end of CoverTab[28725]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:271
			_go_fuzz_dep_.CoverTab[28726]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:271
			// _ = "end of CoverTab[28726]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:271
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:271
		// _ = "end of CoverTab[28722]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:271
		_go_fuzz_dep_.CoverTab[28723]++
														if index == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:272
			_go_fuzz_dep_.CoverTab[28727]++
															index = ts.numBuckets
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:273
			// _ = "end of CoverTab[28727]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:274
			_go_fuzz_dep_.CoverTab[28728]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:274
			// _ = "end of CoverTab[28728]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:274
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:274
		// _ = "end of CoverTab[28723]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:274
		_go_fuzz_dep_.CoverTab[28724]++
														index--
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:275
		// _ = "end of CoverTab[28724]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:276
	// _ = "end of CoverTab[28718]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:276
	_go_fuzz_dep_.CoverTab[28719]++

													return result
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:278
	// _ = "end of CoverTab[28719]"
}

// LatestBuckets returns a copy of the num latest buckets from level.
func (ts *timeSeries) LatestBuckets(level, num int) []Observable {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:282
	_go_fuzz_dep_.CoverTab[28729]++
													if level < 0 || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:283
		_go_fuzz_dep_.CoverTab[28734]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:283
		return level > len(ts.levels)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:283
		// _ = "end of CoverTab[28734]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:283
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:283
		_go_fuzz_dep_.CoverTab[28735]++
														log.Print("timeseries: bad level argument: ", level)
														return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:285
		// _ = "end of CoverTab[28735]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:286
		_go_fuzz_dep_.CoverTab[28736]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:286
		// _ = "end of CoverTab[28736]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:286
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:286
	// _ = "end of CoverTab[28729]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:286
	_go_fuzz_dep_.CoverTab[28730]++
													if num < 0 || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:287
		_go_fuzz_dep_.CoverTab[28737]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:287
		return num >= ts.numBuckets
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:287
		// _ = "end of CoverTab[28737]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:287
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:287
		_go_fuzz_dep_.CoverTab[28738]++
														log.Print("timeseries: bad num argument: ", num)
														return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:289
		// _ = "end of CoverTab[28738]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:290
		_go_fuzz_dep_.CoverTab[28739]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:290
		// _ = "end of CoverTab[28739]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:290
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:290
	// _ = "end of CoverTab[28730]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:290
	_go_fuzz_dep_.CoverTab[28731]++

													results := make([]Observable, num)
													now := ts.clock.Time()
													if ts.levels[0].end.Before(now) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:294
		_go_fuzz_dep_.CoverTab[28740]++
														ts.advance(now)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:295
		// _ = "end of CoverTab[28740]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:296
		_go_fuzz_dep_.CoverTab[28741]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:296
		// _ = "end of CoverTab[28741]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:296
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:296
	// _ = "end of CoverTab[28731]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:296
	_go_fuzz_dep_.CoverTab[28732]++

													ts.mergePendingUpdates()

													l := ts.levels[level]
													index := l.newest

													for i := 0; i < num; i++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:303
		_go_fuzz_dep_.CoverTab[28742]++
														result := ts.provider()
														results[i] = result
														if l.buckets[index] != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:306
			_go_fuzz_dep_.CoverTab[28745]++
															result.CopyFrom(l.buckets[index])
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:307
			// _ = "end of CoverTab[28745]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:308
			_go_fuzz_dep_.CoverTab[28746]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:308
			// _ = "end of CoverTab[28746]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:308
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:308
		// _ = "end of CoverTab[28742]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:308
		_go_fuzz_dep_.CoverTab[28743]++

														if index == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:310
			_go_fuzz_dep_.CoverTab[28747]++
															index = ts.numBuckets
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:311
			// _ = "end of CoverTab[28747]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:312
			_go_fuzz_dep_.CoverTab[28748]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:312
			// _ = "end of CoverTab[28748]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:312
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:312
		// _ = "end of CoverTab[28743]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:312
		_go_fuzz_dep_.CoverTab[28744]++
														index -= 1
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:313
		// _ = "end of CoverTab[28744]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:314
	// _ = "end of CoverTab[28732]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:314
	_go_fuzz_dep_.CoverTab[28733]++
													return results
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:315
	// _ = "end of CoverTab[28733]"
}

// ScaleBy updates observations by scaling by factor.
func (ts *timeSeries) ScaleBy(factor float64) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:319
	_go_fuzz_dep_.CoverTab[28749]++
													for _, l := range ts.levels {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:320
		_go_fuzz_dep_.CoverTab[28751]++
														for i := 0; i < ts.numBuckets; i++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:321
			_go_fuzz_dep_.CoverTab[28752]++
															l.buckets[i].Multiply(factor)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:322
			// _ = "end of CoverTab[28752]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:323
		// _ = "end of CoverTab[28751]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:324
	// _ = "end of CoverTab[28749]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:324
	_go_fuzz_dep_.CoverTab[28750]++

													ts.total.Multiply(factor)
													ts.pending.Multiply(factor)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:327
	// _ = "end of CoverTab[28750]"
}

// Range returns the sum of observations added over the specified time range.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:330
// If start or finish times don't fall on bucket boundaries of the same
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:330
// level, then return values are approximate answers.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:333
func (ts *timeSeries) Range(start, finish time.Time) Observable {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:333
	_go_fuzz_dep_.CoverTab[28753]++
													return ts.ComputeRange(start, finish, 1)[0]
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:334
	// _ = "end of CoverTab[28753]"
}

// Recent returns the sum of observations from the last delta.
func (ts *timeSeries) Recent(delta time.Duration) Observable {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:338
	_go_fuzz_dep_.CoverTab[28754]++
													now := ts.clock.Time()
													return ts.Range(now.Add(-delta), now)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:340
	// _ = "end of CoverTab[28754]"
}

// Total returns the total of all observations.
func (ts *timeSeries) Total() Observable {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:344
	_go_fuzz_dep_.CoverTab[28755]++
													ts.mergePendingUpdates()
													return ts.total
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:346
	// _ = "end of CoverTab[28755]"
}

// ComputeRange computes a specified number of values into a slice using
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:349
// the observations recorded over the specified time period. The return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:349
// values are approximate if the start or finish times don't fall on the
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:349
// bucket boundaries at the same level or if the number of buckets spanning
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:349
// the range is not an integral multiple of num.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:354
func (ts *timeSeries) ComputeRange(start, finish time.Time, num int) []Observable {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:354
	_go_fuzz_dep_.CoverTab[28756]++
													if start.After(finish) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:355
		_go_fuzz_dep_.CoverTab[28760]++
														log.Printf("timeseries: start > finish, %v>%v", start, finish)
														return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:357
		// _ = "end of CoverTab[28760]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:358
		_go_fuzz_dep_.CoverTab[28761]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:358
		// _ = "end of CoverTab[28761]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:358
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:358
	// _ = "end of CoverTab[28756]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:358
	_go_fuzz_dep_.CoverTab[28757]++

													if num < 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:360
		_go_fuzz_dep_.CoverTab[28762]++
														log.Printf("timeseries: num < 0, %v", num)
														return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:362
		// _ = "end of CoverTab[28762]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:363
		_go_fuzz_dep_.CoverTab[28763]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:363
		// _ = "end of CoverTab[28763]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:363
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:363
	// _ = "end of CoverTab[28757]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:363
	_go_fuzz_dep_.CoverTab[28758]++

													results := make([]Observable, num)

													for _, l := range ts.levels {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:367
		_go_fuzz_dep_.CoverTab[28764]++
														if !start.Before(l.end.Add(-l.size * time.Duration(ts.numBuckets))) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:368
			_go_fuzz_dep_.CoverTab[28765]++
															ts.extract(l, start, finish, num, results)
															return results
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:370
			// _ = "end of CoverTab[28765]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:371
			_go_fuzz_dep_.CoverTab[28766]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:371
			// _ = "end of CoverTab[28766]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:371
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:371
		// _ = "end of CoverTab[28764]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:372
	// _ = "end of CoverTab[28758]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:372
	_go_fuzz_dep_.CoverTab[28759]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:377
	ts.extract(ts.levels[len(ts.levels)-1], start, finish, num, results)

													return results
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:379
	// _ = "end of CoverTab[28759]"
}

// RecentList returns the specified number of values in slice over the most
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:382
// recent time period of the specified range.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:384
func (ts *timeSeries) RecentList(delta time.Duration, num int) []Observable {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:384
	_go_fuzz_dep_.CoverTab[28767]++
													if delta < 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:385
		_go_fuzz_dep_.CoverTab[28769]++
														return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:386
		// _ = "end of CoverTab[28769]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:387
		_go_fuzz_dep_.CoverTab[28770]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:387
		// _ = "end of CoverTab[28770]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:387
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:387
	// _ = "end of CoverTab[28767]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:387
	_go_fuzz_dep_.CoverTab[28768]++
													now := ts.clock.Time()
													return ts.ComputeRange(now.Add(-delta), now, num)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:389
	// _ = "end of CoverTab[28768]"
}

// extract returns a slice of specified number of observations from a given
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:392
// level over a given range.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:394
func (ts *timeSeries) extract(l *tsLevel, start, finish time.Time, num int, results []Observable) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:394
	_go_fuzz_dep_.CoverTab[28771]++
													ts.mergePendingUpdates()

													srcInterval := l.size
													dstInterval := finish.Sub(start) / time.Duration(num)
													dstStart := start
													srcStart := l.end.Add(-srcInterval * time.Duration(ts.numBuckets))

													srcIndex := 0

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:405
	if dstStart.After(srcStart) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:405
		_go_fuzz_dep_.CoverTab[28773]++
														advance := int(dstStart.Sub(srcStart) / srcInterval)
														srcIndex += advance
														srcStart = srcStart.Add(time.Duration(advance) * srcInterval)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:408
		// _ = "end of CoverTab[28773]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:409
		_go_fuzz_dep_.CoverTab[28774]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:409
		// _ = "end of CoverTab[28774]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:409
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:409
	// _ = "end of CoverTab[28771]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:409
	_go_fuzz_dep_.CoverTab[28772]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:416
	for i := 0; i < num; i++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:416
		_go_fuzz_dep_.CoverTab[28775]++
														results[i] = ts.resetObservation(results[i])
														dstEnd := dstStart.Add(dstInterval)
														for srcIndex < ts.numBuckets && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:419
			_go_fuzz_dep_.CoverTab[28777]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:419
			return srcStart.Before(dstEnd)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:419
			// _ = "end of CoverTab[28777]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:419
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:419
			_go_fuzz_dep_.CoverTab[28778]++
															srcEnd := srcStart.Add(srcInterval)
															if srcEnd.After(ts.lastAdd) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:421
				_go_fuzz_dep_.CoverTab[28781]++
																srcEnd = ts.lastAdd
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:422
				// _ = "end of CoverTab[28781]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:423
				_go_fuzz_dep_.CoverTab[28782]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:423
				// _ = "end of CoverTab[28782]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:423
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:423
			// _ = "end of CoverTab[28778]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:423
			_go_fuzz_dep_.CoverTab[28779]++

															if !srcEnd.Before(dstStart) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:425
				_go_fuzz_dep_.CoverTab[28783]++
																srcValue := l.buckets[(srcIndex+l.oldest)%ts.numBuckets]
																if !srcStart.Before(dstStart) && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:427
					_go_fuzz_dep_.CoverTab[28785]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:427
					return !srcEnd.After(dstEnd)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:427
					// _ = "end of CoverTab[28785]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:427
				}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:427
					_go_fuzz_dep_.CoverTab[28786]++

																	if srcValue != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:429
						_go_fuzz_dep_.CoverTab[28787]++
																		results[i].Add(srcValue)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:430
						// _ = "end of CoverTab[28787]"
					} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:431
						_go_fuzz_dep_.CoverTab[28788]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:431
						// _ = "end of CoverTab[28788]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:431
					}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:431
					// _ = "end of CoverTab[28786]"
				} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:432
					_go_fuzz_dep_.CoverTab[28789]++

																	overlapStart := maxTime(srcStart, dstStart)
																	overlapEnd := minTime(srcEnd, dstEnd)
																	base := srcEnd.Sub(srcStart)
																	fraction := overlapEnd.Sub(overlapStart).Seconds() / base.Seconds()

																	used := ts.provider()
																	if srcValue != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:440
						_go_fuzz_dep_.CoverTab[28791]++
																		used.CopyFrom(srcValue)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:441
						// _ = "end of CoverTab[28791]"
					} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:442
						_go_fuzz_dep_.CoverTab[28792]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:442
						// _ = "end of CoverTab[28792]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:442
					}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:442
					// _ = "end of CoverTab[28789]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:442
					_go_fuzz_dep_.CoverTab[28790]++
																	used.Multiply(fraction)
																	results[i].Add(used)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:444
					// _ = "end of CoverTab[28790]"
				}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:445
				// _ = "end of CoverTab[28783]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:445
				_go_fuzz_dep_.CoverTab[28784]++

																if srcEnd.After(dstEnd) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:447
					_go_fuzz_dep_.CoverTab[28793]++
																	break
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:448
					// _ = "end of CoverTab[28793]"
				} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:449
					_go_fuzz_dep_.CoverTab[28794]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:449
					// _ = "end of CoverTab[28794]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:449
				}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:449
				// _ = "end of CoverTab[28784]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:450
				_go_fuzz_dep_.CoverTab[28795]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:450
				// _ = "end of CoverTab[28795]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:450
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:450
			// _ = "end of CoverTab[28779]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:450
			_go_fuzz_dep_.CoverTab[28780]++
															srcIndex++
															srcStart = srcStart.Add(srcInterval)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:452
			// _ = "end of CoverTab[28780]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:453
		// _ = "end of CoverTab[28775]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:453
		_go_fuzz_dep_.CoverTab[28776]++
														dstStart = dstStart.Add(dstInterval)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:454
		// _ = "end of CoverTab[28776]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:455
	// _ = "end of CoverTab[28772]"
}

// resetObservation clears the content so the struct may be reused.
func (ts *timeSeries) resetObservation(observation Observable) Observable {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:459
	_go_fuzz_dep_.CoverTab[28796]++
													if observation == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:460
		_go_fuzz_dep_.CoverTab[28798]++
														observation = ts.provider()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:461
		// _ = "end of CoverTab[28798]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:462
		_go_fuzz_dep_.CoverTab[28799]++
														observation.Clear()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:463
		// _ = "end of CoverTab[28799]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:464
	// _ = "end of CoverTab[28796]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:464
	_go_fuzz_dep_.CoverTab[28797]++
													return observation
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:465
	// _ = "end of CoverTab[28797]"
}

// TimeSeries tracks data at granularities from 1 second to 16 weeks.
type TimeSeries struct {
	timeSeries
}

// NewTimeSeries creates a new TimeSeries using the function provided for creating new Observable.
func NewTimeSeries(f func() Observable) *TimeSeries {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:474
	_go_fuzz_dep_.CoverTab[28800]++
													return NewTimeSeriesWithClock(f, defaultClockInstance)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:475
	// _ = "end of CoverTab[28800]"
}

// NewTimeSeriesWithClock creates a new TimeSeries using the function provided for creating new Observable and the clock for
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:478
// assigning timestamps.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:480
func NewTimeSeriesWithClock(f func() Observable, clock Clock) *TimeSeries {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:480
	_go_fuzz_dep_.CoverTab[28801]++
													ts := new(TimeSeries)
													ts.timeSeries.init(timeSeriesResolutions, f, timeSeriesNumBuckets, clock)
													return ts
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:483
	// _ = "end of CoverTab[28801]"
}

// MinuteHourSeries tracks data at granularities of 1 minute and 1 hour.
type MinuteHourSeries struct {
	timeSeries
}

// NewMinuteHourSeries creates a new MinuteHourSeries using the function provided for creating new Observable.
func NewMinuteHourSeries(f func() Observable) *MinuteHourSeries {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:492
	_go_fuzz_dep_.CoverTab[28802]++
													return NewMinuteHourSeriesWithClock(f, defaultClockInstance)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:493
	// _ = "end of CoverTab[28802]"
}

// NewMinuteHourSeriesWithClock creates a new MinuteHourSeries using the function provided for creating new Observable and the clock for
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:496
// assigning timestamps.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:498
func NewMinuteHourSeriesWithClock(f func() Observable, clock Clock) *MinuteHourSeries {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:498
	_go_fuzz_dep_.CoverTab[28803]++
													ts := new(MinuteHourSeries)
													ts.timeSeries.init(minuteHourSeriesResolutions, f,
		minuteHourSeriesNumBuckets, clock)
													return ts
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:502
	// _ = "end of CoverTab[28803]"
}

func (ts *MinuteHourSeries) Minute() Observable {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:505
	_go_fuzz_dep_.CoverTab[28804]++
													return ts.timeSeries.Latest(0, 60)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:506
	// _ = "end of CoverTab[28804]"
}

func (ts *MinuteHourSeries) Hour() Observable {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:509
	_go_fuzz_dep_.CoverTab[28805]++
													return ts.timeSeries.Latest(1, 60)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:510
	// _ = "end of CoverTab[28805]"
}

func minTime(a, b time.Time) time.Time {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:513
	_go_fuzz_dep_.CoverTab[28806]++
													if a.Before(b) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:514
		_go_fuzz_dep_.CoverTab[28808]++
														return a
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:515
		// _ = "end of CoverTab[28808]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:516
		_go_fuzz_dep_.CoverTab[28809]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:516
		// _ = "end of CoverTab[28809]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:516
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:516
	// _ = "end of CoverTab[28806]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:516
	_go_fuzz_dep_.CoverTab[28807]++
													return b
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:517
	// _ = "end of CoverTab[28807]"
}

func maxTime(a, b time.Time) time.Time {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:520
	_go_fuzz_dep_.CoverTab[28810]++
													if a.After(b) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:521
		_go_fuzz_dep_.CoverTab[28812]++
														return a
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:522
		// _ = "end of CoverTab[28812]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:523
		_go_fuzz_dep_.CoverTab[28813]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:523
		// _ = "end of CoverTab[28813]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:523
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:523
	// _ = "end of CoverTab[28810]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:523
	_go_fuzz_dep_.CoverTab[28811]++
													return b
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:524
	// _ = "end of CoverTab[28811]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:525
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/timeseries/timeseries.go:525
var _ = _go_fuzz_dep_.CoverTab
