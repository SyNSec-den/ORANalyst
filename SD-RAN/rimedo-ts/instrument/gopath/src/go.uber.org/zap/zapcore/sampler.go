// Copyright (c) 2016 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:21
package zapcore

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:21
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:21
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:21
)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:21
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:21
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:21
)

import (
	"time"

	"go.uber.org/atomic"
)

const (
	_numLevels		= _maxLevel - _minLevel + 1
	_countersPerLevel	= 4096
)

type counter struct {
	resetAt	atomic.Int64
	counter	atomic.Uint64
}

type counters [_numLevels][_countersPerLevel]counter

func newCounters() *counters {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:41
	_go_fuzz_dep_.CoverTab[131187]++
											return &counters{}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:42
	// _ = "end of CoverTab[131187]"
}

func (cs *counters) get(lvl Level, key string) *counter {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:45
	_go_fuzz_dep_.CoverTab[131188]++
											i := lvl - _minLevel
											j := fnv32a(key) % _countersPerLevel
											return &cs[i][j]
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:48
	// _ = "end of CoverTab[131188]"
}

// fnv32a, adapted from "hash/fnv", but without a []byte(string) alloc
func fnv32a(s string) uint32 {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:52
	_go_fuzz_dep_.CoverTab[131189]++
											const (
		offset32	= 2166136261
		prime32		= 16777619
	)
	hash := uint32(offset32)
	for i := 0; i < len(s); i++ {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:58
		_go_fuzz_dep_.CoverTab[131191]++
												hash ^= uint32(s[i])
												hash *= prime32
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:60
		// _ = "end of CoverTab[131191]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:61
	// _ = "end of CoverTab[131189]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:61
	_go_fuzz_dep_.CoverTab[131190]++
											return hash
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:62
	// _ = "end of CoverTab[131190]"
}

func (c *counter) IncCheckReset(t time.Time, tick time.Duration) uint64 {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:65
	_go_fuzz_dep_.CoverTab[131192]++
											tn := t.UnixNano()
											resetAfter := c.resetAt.Load()
											if resetAfter > tn {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:68
		_go_fuzz_dep_.CoverTab[131195]++
												return c.counter.Inc()
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:69
		// _ = "end of CoverTab[131195]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:70
		_go_fuzz_dep_.CoverTab[131196]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:70
		// _ = "end of CoverTab[131196]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:70
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:70
	// _ = "end of CoverTab[131192]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:70
	_go_fuzz_dep_.CoverTab[131193]++

											c.counter.Store(1)

											newResetAfter := tn + tick.Nanoseconds()
											if !c.resetAt.CAS(resetAfter, newResetAfter) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:75
		_go_fuzz_dep_.CoverTab[131197]++

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:78
		return c.counter.Inc()
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:78
		// _ = "end of CoverTab[131197]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:79
		_go_fuzz_dep_.CoverTab[131198]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:79
		// _ = "end of CoverTab[131198]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:79
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:79
	// _ = "end of CoverTab[131193]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:79
	_go_fuzz_dep_.CoverTab[131194]++

											return 1
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:81
	// _ = "end of CoverTab[131194]"
}

// SamplingDecision is a decision represented as a bit field made by sampler.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:84
// More decisions may be added in the future.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:86
type SamplingDecision uint32

const (
	// LogDropped indicates that the Sampler dropped a log entry.
	LogDropped	SamplingDecision	= 1 << iota
	// LogSampled indicates that the Sampler sampled a log entry.
	LogSampled
)

// optionFunc wraps a func so it satisfies the SamplerOption interface.
type optionFunc func(*sampler)

func (f optionFunc) apply(s *sampler) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:98
	_go_fuzz_dep_.CoverTab[131199]++
											f(s)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:99
	// _ = "end of CoverTab[131199]"
}

// SamplerOption configures a Sampler.
type SamplerOption interface {
	apply(*sampler)
}

// nopSamplingHook is the default hook used by sampler.
func nopSamplingHook(Entry, SamplingDecision) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:108
	_go_fuzz_dep_.CoverTab[131200]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:108
	// _ = "end of CoverTab[131200]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:108
}

// SamplerHook registers a function  which will be called when Sampler makes a
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:110
// decision.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:110
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:110
// This hook may be used to get visibility into the performance of the sampler.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:110
// For example, use it to track metrics of dropped versus sampled logs.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:110
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:110
//	var dropped atomic.Int64
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:110
//	zapcore.SamplerHook(func(ent zapcore.Entry, dec zapcore.SamplingDecision) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:110
//	  if dec&zapcore.LogDropped > 0 {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:110
//	    dropped.Inc()
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:110
//	  }
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:110
//	})
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:122
func SamplerHook(hook func(entry Entry, dec SamplingDecision)) SamplerOption {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:122
	_go_fuzz_dep_.CoverTab[131201]++
											return optionFunc(func(s *sampler) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:123
		_go_fuzz_dep_.CoverTab[131202]++
												s.hook = hook
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:124
		// _ = "end of CoverTab[131202]"
	})
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:125
	// _ = "end of CoverTab[131201]"
}

// NewSamplerWithOptions creates a Core that samples incoming entries, which
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:128
// caps the CPU and I/O load of logging while attempting to preserve a
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:128
// representative subset of your logs.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:128
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:128
// Zap samples by logging the first N entries with a given level and message
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:128
// each tick. If more Entries with the same level and message are seen during
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:128
// the same interval, every Mth message is logged and the rest are dropped.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:128
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:128
// Sampler can be configured to report sampling decisions with the SamplerHook
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:128
// option.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:128
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:128
// Keep in mind that zap's sampling implementation is optimized for speed over
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:128
// absolute precision; under load, each tick may be slightly over- or
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:128
// under-sampled.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:142
func NewSamplerWithOptions(core Core, tick time.Duration, first, thereafter int, opts ...SamplerOption) Core {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:142
	_go_fuzz_dep_.CoverTab[131203]++
											s := &sampler{
		Core:		core,
		tick:		tick,
		counts:		newCounters(),
		first:		uint64(first),
		thereafter:	uint64(thereafter),
		hook:		nopSamplingHook,
	}
	for _, opt := range opts {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:151
		_go_fuzz_dep_.CoverTab[131205]++
												opt.apply(s)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:152
		// _ = "end of CoverTab[131205]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:153
	// _ = "end of CoverTab[131203]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:153
	_go_fuzz_dep_.CoverTab[131204]++

											return s
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:155
	// _ = "end of CoverTab[131204]"
}

type sampler struct {
	Core

	counts			*counters
	tick			time.Duration
	first, thereafter	uint64
	hook			func(Entry, SamplingDecision)
}

// NewSampler creates a Core that samples incoming entries, which
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:167
// caps the CPU and I/O load of logging while attempting to preserve a
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:167
// representative subset of your logs.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:167
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:167
// Zap samples by logging the first N entries with a given level and message
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:167
// each tick. If more Entries with the same level and message are seen during
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:167
// the same interval, every Mth message is logged and the rest are dropped.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:167
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:167
// Keep in mind that zap's sampling implementation is optimized for speed over
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:167
// absolute precision; under load, each tick may be slightly over- or
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:167
// under-sampled.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:167
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:167
// Deprecated: use NewSamplerWithOptions.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:180
func NewSampler(core Core, tick time.Duration, first, thereafter int) Core {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:180
	_go_fuzz_dep_.CoverTab[131206]++
											return NewSamplerWithOptions(core, tick, first, thereafter)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:181
	// _ = "end of CoverTab[131206]"
}

func (s *sampler) With(fields []Field) Core {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:184
	_go_fuzz_dep_.CoverTab[131207]++
											return &sampler{
		Core:		s.Core.With(fields),
		tick:		s.tick,
		counts:		s.counts,
		first:		s.first,
		thereafter:	s.thereafter,
		hook:		s.hook,
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:192
	// _ = "end of CoverTab[131207]"
}

func (s *sampler) Check(ent Entry, ce *CheckedEntry) *CheckedEntry {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:195
	_go_fuzz_dep_.CoverTab[131208]++
											if !s.Enabled(ent.Level) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:196
		_go_fuzz_dep_.CoverTab[131211]++
												return ce
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:197
		// _ = "end of CoverTab[131211]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:198
		_go_fuzz_dep_.CoverTab[131212]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:198
		// _ = "end of CoverTab[131212]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:198
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:198
	// _ = "end of CoverTab[131208]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:198
	_go_fuzz_dep_.CoverTab[131209]++

											counter := s.counts.get(ent.Level, ent.Message)
											n := counter.IncCheckReset(ent.Time, s.tick)
											if n > s.first && func() bool {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:202
		_go_fuzz_dep_.CoverTab[131213]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:202
		return (n-s.first)%s.thereafter != 0
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:202
		// _ = "end of CoverTab[131213]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:202
	}() {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:202
		_go_fuzz_dep_.CoverTab[131214]++
												s.hook(ent, LogDropped)
												return ce
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:204
		// _ = "end of CoverTab[131214]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:205
		_go_fuzz_dep_.CoverTab[131215]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:205
		// _ = "end of CoverTab[131215]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:205
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:205
	// _ = "end of CoverTab[131209]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:205
	_go_fuzz_dep_.CoverTab[131210]++
											s.hook(ent, LogSampled)
											return s.Core.Check(ent, ce)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:207
	// _ = "end of CoverTab[131210]"
}

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:208
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/sampler.go:208
var _ = _go_fuzz_dep_.CoverTab
