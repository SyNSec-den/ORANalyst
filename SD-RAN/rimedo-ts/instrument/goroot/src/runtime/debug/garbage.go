// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/runtime/debug/garbage.go:5
package debug

//line /usr/local/go/src/runtime/debug/garbage.go:5
import (
//line /usr/local/go/src/runtime/debug/garbage.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/runtime/debug/garbage.go:5
)
//line /usr/local/go/src/runtime/debug/garbage.go:5
import (
//line /usr/local/go/src/runtime/debug/garbage.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/runtime/debug/garbage.go:5
)

import (
	"runtime"
	"sort"
	"time"
)

// GCStats collect information about recent garbage collections.
type GCStats struct {
	LastGC		time.Time	// time of last collection
	NumGC		int64		// number of garbage collections
	PauseTotal	time.Duration	// total pause for all collections
	Pause		[]time.Duration	// pause history, most recent first
	PauseEnd	[]time.Time	// pause end times history, most recent first
	PauseQuantiles	[]time.Duration
}

// ReadGCStats reads statistics about garbage collection into stats.
//line /usr/local/go/src/runtime/debug/garbage.go:23
// The number of entries in the pause history is system-dependent;
//line /usr/local/go/src/runtime/debug/garbage.go:23
// stats.Pause slice will be reused if large enough, reallocated otherwise.
//line /usr/local/go/src/runtime/debug/garbage.go:23
// ReadGCStats may use the full capacity of the stats.Pause slice.
//line /usr/local/go/src/runtime/debug/garbage.go:23
// If stats.PauseQuantiles is non-empty, ReadGCStats fills it with quantiles
//line /usr/local/go/src/runtime/debug/garbage.go:23
// summarizing the distribution of pause time. For example, if
//line /usr/local/go/src/runtime/debug/garbage.go:23
// len(stats.PauseQuantiles) is 5, it will be filled with the minimum,
//line /usr/local/go/src/runtime/debug/garbage.go:23
// 25%, 50%, 75%, and maximum pause times.
//line /usr/local/go/src/runtime/debug/garbage.go:31
func ReadGCStats(stats *GCStats) {
//line /usr/local/go/src/runtime/debug/garbage.go:31
	_go_fuzz_dep_.CoverTab[90710]++
	// Create a buffer with space for at least two copies of the
	// pause history tracked by the runtime. One will be returned
	// to the caller and the other will be used as transfer buffer
	// for end times history and as a temporary buffer for
	// computing quantiles.
	const maxPause = len(((*runtime.MemStats)(nil)).PauseNs)
	if cap(stats.Pause) < 2*maxPause+3 {
//line /usr/local/go/src/runtime/debug/garbage.go:38
		_go_fuzz_dep_.CoverTab[90714]++
								stats.Pause = make([]time.Duration, 2*maxPause+3)
//line /usr/local/go/src/runtime/debug/garbage.go:39
		// _ = "end of CoverTab[90714]"
	} else {
//line /usr/local/go/src/runtime/debug/garbage.go:40
		_go_fuzz_dep_.CoverTab[90715]++
//line /usr/local/go/src/runtime/debug/garbage.go:40
		// _ = "end of CoverTab[90715]"
//line /usr/local/go/src/runtime/debug/garbage.go:40
	}
//line /usr/local/go/src/runtime/debug/garbage.go:40
	// _ = "end of CoverTab[90710]"
//line /usr/local/go/src/runtime/debug/garbage.go:40
	_go_fuzz_dep_.CoverTab[90711]++

//line /usr/local/go/src/runtime/debug/garbage.go:48
	readGCStats(&stats.Pause)
	n := len(stats.Pause) - 3
	stats.LastGC = time.Unix(0, int64(stats.Pause[n]))
	stats.NumGC = int64(stats.Pause[n+1])
	stats.PauseTotal = stats.Pause[n+2]
	n /= 2
	stats.Pause = stats.Pause[:n]

	if cap(stats.PauseEnd) < maxPause {
//line /usr/local/go/src/runtime/debug/garbage.go:56
		_go_fuzz_dep_.CoverTab[90716]++
								stats.PauseEnd = make([]time.Time, 0, maxPause)
//line /usr/local/go/src/runtime/debug/garbage.go:57
		// _ = "end of CoverTab[90716]"
	} else {
//line /usr/local/go/src/runtime/debug/garbage.go:58
		_go_fuzz_dep_.CoverTab[90717]++
//line /usr/local/go/src/runtime/debug/garbage.go:58
		// _ = "end of CoverTab[90717]"
//line /usr/local/go/src/runtime/debug/garbage.go:58
	}
//line /usr/local/go/src/runtime/debug/garbage.go:58
	// _ = "end of CoverTab[90711]"
//line /usr/local/go/src/runtime/debug/garbage.go:58
	_go_fuzz_dep_.CoverTab[90712]++
							stats.PauseEnd = stats.PauseEnd[:0]
							for _, ns := range stats.Pause[n : n+n] {
//line /usr/local/go/src/runtime/debug/garbage.go:60
		_go_fuzz_dep_.CoverTab[90718]++
								stats.PauseEnd = append(stats.PauseEnd, time.Unix(0, int64(ns)))
//line /usr/local/go/src/runtime/debug/garbage.go:61
		// _ = "end of CoverTab[90718]"
	}
//line /usr/local/go/src/runtime/debug/garbage.go:62
	// _ = "end of CoverTab[90712]"
//line /usr/local/go/src/runtime/debug/garbage.go:62
	_go_fuzz_dep_.CoverTab[90713]++

							if len(stats.PauseQuantiles) > 0 {
//line /usr/local/go/src/runtime/debug/garbage.go:64
		_go_fuzz_dep_.CoverTab[90719]++
								if n == 0 {
//line /usr/local/go/src/runtime/debug/garbage.go:65
			_go_fuzz_dep_.CoverTab[90720]++
									for i := range stats.PauseQuantiles {
//line /usr/local/go/src/runtime/debug/garbage.go:66
				_go_fuzz_dep_.CoverTab[90721]++
										stats.PauseQuantiles[i] = 0
//line /usr/local/go/src/runtime/debug/garbage.go:67
				// _ = "end of CoverTab[90721]"
			}
//line /usr/local/go/src/runtime/debug/garbage.go:68
			// _ = "end of CoverTab[90720]"
		} else {
//line /usr/local/go/src/runtime/debug/garbage.go:69
			_go_fuzz_dep_.CoverTab[90722]++

//line /usr/local/go/src/runtime/debug/garbage.go:72
			sorted := stats.Pause[n : n+n]
			copy(sorted, stats.Pause)
			sort.Slice(sorted, func(i, j int) bool {
//line /usr/local/go/src/runtime/debug/garbage.go:74
				_go_fuzz_dep_.CoverTab[90725]++
//line /usr/local/go/src/runtime/debug/garbage.go:74
				return sorted[i] < sorted[j]
//line /usr/local/go/src/runtime/debug/garbage.go:74
				// _ = "end of CoverTab[90725]"
//line /usr/local/go/src/runtime/debug/garbage.go:74
			})
//line /usr/local/go/src/runtime/debug/garbage.go:74
			// _ = "end of CoverTab[90722]"
//line /usr/local/go/src/runtime/debug/garbage.go:74
			_go_fuzz_dep_.CoverTab[90723]++
									nq := len(stats.PauseQuantiles) - 1
									for i := 0; i < nq; i++ {
//line /usr/local/go/src/runtime/debug/garbage.go:76
				_go_fuzz_dep_.CoverTab[90726]++
										stats.PauseQuantiles[i] = sorted[len(sorted)*i/nq]
//line /usr/local/go/src/runtime/debug/garbage.go:77
				// _ = "end of CoverTab[90726]"
			}
//line /usr/local/go/src/runtime/debug/garbage.go:78
			// _ = "end of CoverTab[90723]"
//line /usr/local/go/src/runtime/debug/garbage.go:78
			_go_fuzz_dep_.CoverTab[90724]++
									stats.PauseQuantiles[nq] = sorted[len(sorted)-1]
//line /usr/local/go/src/runtime/debug/garbage.go:79
			// _ = "end of CoverTab[90724]"
		}
//line /usr/local/go/src/runtime/debug/garbage.go:80
		// _ = "end of CoverTab[90719]"
	} else {
//line /usr/local/go/src/runtime/debug/garbage.go:81
		_go_fuzz_dep_.CoverTab[90727]++
//line /usr/local/go/src/runtime/debug/garbage.go:81
		// _ = "end of CoverTab[90727]"
//line /usr/local/go/src/runtime/debug/garbage.go:81
	}
//line /usr/local/go/src/runtime/debug/garbage.go:81
	// _ = "end of CoverTab[90713]"
}

// SetGCPercent sets the garbage collection target percentage:
//line /usr/local/go/src/runtime/debug/garbage.go:84
// a collection is triggered when the ratio of freshly allocated data
//line /usr/local/go/src/runtime/debug/garbage.go:84
// to live data remaining after the previous collection reaches this percentage.
//line /usr/local/go/src/runtime/debug/garbage.go:84
// SetGCPercent returns the previous setting.
//line /usr/local/go/src/runtime/debug/garbage.go:84
// The initial setting is the value of the GOGC environment variable
//line /usr/local/go/src/runtime/debug/garbage.go:84
// at startup, or 100 if the variable is not set.
//line /usr/local/go/src/runtime/debug/garbage.go:84
// This setting may be effectively reduced in order to maintain a memory
//line /usr/local/go/src/runtime/debug/garbage.go:84
// limit.
//line /usr/local/go/src/runtime/debug/garbage.go:84
// A negative percentage effectively disables garbage collection, unless
//line /usr/local/go/src/runtime/debug/garbage.go:84
// the memory limit is reached.
//line /usr/local/go/src/runtime/debug/garbage.go:84
// See SetMemoryLimit for more details.
//line /usr/local/go/src/runtime/debug/garbage.go:95
func SetGCPercent(percent int) int {
//line /usr/local/go/src/runtime/debug/garbage.go:95
	_go_fuzz_dep_.CoverTab[90728]++
							return int(setGCPercent(int32(percent)))
//line /usr/local/go/src/runtime/debug/garbage.go:96
	// _ = "end of CoverTab[90728]"
}

// FreeOSMemory forces a garbage collection followed by an
//line /usr/local/go/src/runtime/debug/garbage.go:99
// attempt to return as much memory to the operating system
//line /usr/local/go/src/runtime/debug/garbage.go:99
// as possible. (Even if this is not called, the runtime gradually
//line /usr/local/go/src/runtime/debug/garbage.go:99
// returns memory to the operating system in a background task.)
//line /usr/local/go/src/runtime/debug/garbage.go:103
func FreeOSMemory() {
//line /usr/local/go/src/runtime/debug/garbage.go:103
	_go_fuzz_dep_.CoverTab[90729]++
							freeOSMemory()
//line /usr/local/go/src/runtime/debug/garbage.go:104
	// _ = "end of CoverTab[90729]"
}

// SetMaxStack sets the maximum amount of memory that
//line /usr/local/go/src/runtime/debug/garbage.go:107
// can be used by a single goroutine stack.
//line /usr/local/go/src/runtime/debug/garbage.go:107
// If any goroutine exceeds this limit while growing its stack,
//line /usr/local/go/src/runtime/debug/garbage.go:107
// the program crashes.
//line /usr/local/go/src/runtime/debug/garbage.go:107
// SetMaxStack returns the previous setting.
//line /usr/local/go/src/runtime/debug/garbage.go:107
// The initial setting is 1 GB on 64-bit systems, 250 MB on 32-bit systems.
//line /usr/local/go/src/runtime/debug/garbage.go:107
// There may be a system-imposed maximum stack limit regardless
//line /usr/local/go/src/runtime/debug/garbage.go:107
// of the value provided to SetMaxStack.
//line /usr/local/go/src/runtime/debug/garbage.go:107
//
//line /usr/local/go/src/runtime/debug/garbage.go:107
// SetMaxStack is useful mainly for limiting the damage done by
//line /usr/local/go/src/runtime/debug/garbage.go:107
// goroutines that enter an infinite recursion. It only limits future
//line /usr/local/go/src/runtime/debug/garbage.go:107
// stack growth.
//line /usr/local/go/src/runtime/debug/garbage.go:119
func SetMaxStack(bytes int) int {
//line /usr/local/go/src/runtime/debug/garbage.go:119
	_go_fuzz_dep_.CoverTab[90730]++
							return setMaxStack(bytes)
//line /usr/local/go/src/runtime/debug/garbage.go:120
	// _ = "end of CoverTab[90730]"
}

// SetMaxThreads sets the maximum number of operating system
//line /usr/local/go/src/runtime/debug/garbage.go:123
// threads that the Go program can use. If it attempts to use more than
//line /usr/local/go/src/runtime/debug/garbage.go:123
// this many, the program crashes.
//line /usr/local/go/src/runtime/debug/garbage.go:123
// SetMaxThreads returns the previous setting.
//line /usr/local/go/src/runtime/debug/garbage.go:123
// The initial setting is 10,000 threads.
//line /usr/local/go/src/runtime/debug/garbage.go:123
//
//line /usr/local/go/src/runtime/debug/garbage.go:123
// The limit controls the number of operating system threads, not the number
//line /usr/local/go/src/runtime/debug/garbage.go:123
// of goroutines. A Go program creates a new thread only when a goroutine
//line /usr/local/go/src/runtime/debug/garbage.go:123
// is ready to run but all the existing threads are blocked in system calls, cgo calls,
//line /usr/local/go/src/runtime/debug/garbage.go:123
// or are locked to other goroutines due to use of runtime.LockOSThread.
//line /usr/local/go/src/runtime/debug/garbage.go:123
//
//line /usr/local/go/src/runtime/debug/garbage.go:123
// SetMaxThreads is useful mainly for limiting the damage done by
//line /usr/local/go/src/runtime/debug/garbage.go:123
// programs that create an unbounded number of threads. The idea is
//line /usr/local/go/src/runtime/debug/garbage.go:123
// to take down the program before it takes down the operating system.
//line /usr/local/go/src/runtime/debug/garbage.go:137
func SetMaxThreads(threads int) int {
//line /usr/local/go/src/runtime/debug/garbage.go:137
	_go_fuzz_dep_.CoverTab[90731]++
							return setMaxThreads(threads)
//line /usr/local/go/src/runtime/debug/garbage.go:138
	// _ = "end of CoverTab[90731]"
}

// SetPanicOnFault controls the runtime's behavior when a program faults
//line /usr/local/go/src/runtime/debug/garbage.go:141
// at an unexpected (non-nil) address. Such faults are typically caused by
//line /usr/local/go/src/runtime/debug/garbage.go:141
// bugs such as runtime memory corruption, so the default response is to crash
//line /usr/local/go/src/runtime/debug/garbage.go:141
// the program. Programs working with memory-mapped files or unsafe
//line /usr/local/go/src/runtime/debug/garbage.go:141
// manipulation of memory may cause faults at non-nil addresses in less
//line /usr/local/go/src/runtime/debug/garbage.go:141
// dramatic situations; SetPanicOnFault allows such programs to request
//line /usr/local/go/src/runtime/debug/garbage.go:141
// that the runtime trigger only a panic, not a crash.
//line /usr/local/go/src/runtime/debug/garbage.go:141
// The runtime.Error that the runtime panics with may have an additional method:
//line /usr/local/go/src/runtime/debug/garbage.go:141
//
//line /usr/local/go/src/runtime/debug/garbage.go:141
//	Addr() uintptr
//line /usr/local/go/src/runtime/debug/garbage.go:141
//
//line /usr/local/go/src/runtime/debug/garbage.go:141
// If that method exists, it returns the memory address which triggered the fault.
//line /usr/local/go/src/runtime/debug/garbage.go:141
// The results of Addr are best-effort and the veracity of the result
//line /usr/local/go/src/runtime/debug/garbage.go:141
// may depend on the platform.
//line /usr/local/go/src/runtime/debug/garbage.go:141
// SetPanicOnFault applies only to the current goroutine.
//line /usr/local/go/src/runtime/debug/garbage.go:141
// It returns the previous setting.
//line /usr/local/go/src/runtime/debug/garbage.go:157
func SetPanicOnFault(enabled bool) bool {
//line /usr/local/go/src/runtime/debug/garbage.go:157
	_go_fuzz_dep_.CoverTab[90732]++
							return setPanicOnFault(enabled)
//line /usr/local/go/src/runtime/debug/garbage.go:158
	// _ = "end of CoverTab[90732]"
}

// WriteHeapDump writes a description of the heap and the objects in
//line /usr/local/go/src/runtime/debug/garbage.go:161
// it to the given file descriptor.
//line /usr/local/go/src/runtime/debug/garbage.go:161
//
//line /usr/local/go/src/runtime/debug/garbage.go:161
// WriteHeapDump suspends the execution of all goroutines until the heap
//line /usr/local/go/src/runtime/debug/garbage.go:161
// dump is completely written.  Thus, the file descriptor must not be
//line /usr/local/go/src/runtime/debug/garbage.go:161
// connected to a pipe or socket whose other end is in the same Go
//line /usr/local/go/src/runtime/debug/garbage.go:161
// process; instead, use a temporary file or network socket.
//line /usr/local/go/src/runtime/debug/garbage.go:161
//
//line /usr/local/go/src/runtime/debug/garbage.go:161
// The heap dump format is defined at https://golang.org/s/go15heapdump.
//line /usr/local/go/src/runtime/debug/garbage.go:170
func WriteHeapDump(fd uintptr)

// SetTraceback sets the amount of detail printed by the runtime in
//line /usr/local/go/src/runtime/debug/garbage.go:172
// the traceback it prints before exiting due to an unrecovered panic
//line /usr/local/go/src/runtime/debug/garbage.go:172
// or an internal runtime error.
//line /usr/local/go/src/runtime/debug/garbage.go:172
// The level argument takes the same values as the GOTRACEBACK
//line /usr/local/go/src/runtime/debug/garbage.go:172
// environment variable. For example, SetTraceback("all") ensure
//line /usr/local/go/src/runtime/debug/garbage.go:172
// that the program prints all goroutines when it crashes.
//line /usr/local/go/src/runtime/debug/garbage.go:172
// See the package runtime documentation for details.
//line /usr/local/go/src/runtime/debug/garbage.go:172
// If SetTraceback is called with a level lower than that of the
//line /usr/local/go/src/runtime/debug/garbage.go:172
// environment variable, the call is ignored.
//line /usr/local/go/src/runtime/debug/garbage.go:181
func SetTraceback(level string)

// SetMemoryLimit provides the runtime with a soft memory limit.
//line /usr/local/go/src/runtime/debug/garbage.go:183
//
//line /usr/local/go/src/runtime/debug/garbage.go:183
// The runtime undertakes several processes to try to respect this
//line /usr/local/go/src/runtime/debug/garbage.go:183
// memory limit, including adjustments to the frequency of garbage
//line /usr/local/go/src/runtime/debug/garbage.go:183
// collections and returning memory to the underlying system more
//line /usr/local/go/src/runtime/debug/garbage.go:183
// aggressively. This limit will be respected even if GOGC=off (or,
//line /usr/local/go/src/runtime/debug/garbage.go:183
// if SetGCPercent(-1) is executed).
//line /usr/local/go/src/runtime/debug/garbage.go:183
//
//line /usr/local/go/src/runtime/debug/garbage.go:183
// The input limit is provided as bytes, and includes all memory
//line /usr/local/go/src/runtime/debug/garbage.go:183
// mapped, managed, and not released by the Go runtime. Notably, it
//line /usr/local/go/src/runtime/debug/garbage.go:183
// does not account for space used by the Go binary and memory
//line /usr/local/go/src/runtime/debug/garbage.go:183
// external to Go, such as memory managed by the underlying system
//line /usr/local/go/src/runtime/debug/garbage.go:183
// on behalf of the process, or memory managed by non-Go code inside
//line /usr/local/go/src/runtime/debug/garbage.go:183
// the same process. Examples of excluded memory sources include: OS
//line /usr/local/go/src/runtime/debug/garbage.go:183
// kernel memory held on behalf of the process, memory allocated by
//line /usr/local/go/src/runtime/debug/garbage.go:183
// C code, and memory mapped by syscall.Mmap (because it is not
//line /usr/local/go/src/runtime/debug/garbage.go:183
// managed by the Go runtime).
//line /usr/local/go/src/runtime/debug/garbage.go:183
//
//line /usr/local/go/src/runtime/debug/garbage.go:183
// More specifically, the following expression accurately reflects
//line /usr/local/go/src/runtime/debug/garbage.go:183
// the value the runtime attempts to maintain as the limit:
//line /usr/local/go/src/runtime/debug/garbage.go:183
//
//line /usr/local/go/src/runtime/debug/garbage.go:183
//	runtime.MemStats.Sys - runtime.MemStats.HeapReleased
//line /usr/local/go/src/runtime/debug/garbage.go:183
//
//line /usr/local/go/src/runtime/debug/garbage.go:183
// or in terms of the runtime/metrics package:
//line /usr/local/go/src/runtime/debug/garbage.go:183
//
//line /usr/local/go/src/runtime/debug/garbage.go:183
//	/memory/classes/total:bytes - /memory/classes/heap/released:bytes
//line /usr/local/go/src/runtime/debug/garbage.go:183
//
//line /usr/local/go/src/runtime/debug/garbage.go:183
// A zero limit or a limit that's lower than the amount of memory
//line /usr/local/go/src/runtime/debug/garbage.go:183
// used by the Go runtime may cause the garbage collector to run
//line /usr/local/go/src/runtime/debug/garbage.go:183
// nearly continuously. However, the application may still make
//line /usr/local/go/src/runtime/debug/garbage.go:183
// progress.
//line /usr/local/go/src/runtime/debug/garbage.go:183
//
//line /usr/local/go/src/runtime/debug/garbage.go:183
// The memory limit is always respected by the Go runtime, so to
//line /usr/local/go/src/runtime/debug/garbage.go:183
// effectively disable this behavior, set the limit very high.
//line /usr/local/go/src/runtime/debug/garbage.go:183
// math.MaxInt64 is the canonical value for disabling the limit,
//line /usr/local/go/src/runtime/debug/garbage.go:183
// but values much greater than the available memory on the underlying
//line /usr/local/go/src/runtime/debug/garbage.go:183
// system work just as well.
//line /usr/local/go/src/runtime/debug/garbage.go:183
//
//line /usr/local/go/src/runtime/debug/garbage.go:183
// See https://go.dev/doc/gc-guide for a detailed guide explaining
//line /usr/local/go/src/runtime/debug/garbage.go:183
// the soft memory limit in more detail, as well as a variety of common
//line /usr/local/go/src/runtime/debug/garbage.go:183
// use-cases and scenarios.
//line /usr/local/go/src/runtime/debug/garbage.go:183
//
//line /usr/local/go/src/runtime/debug/garbage.go:183
// The initial setting is math.MaxInt64 unless the GOMEMLIMIT
//line /usr/local/go/src/runtime/debug/garbage.go:183
// environment variable is set, in which case it provides the initial
//line /usr/local/go/src/runtime/debug/garbage.go:183
// setting. GOMEMLIMIT is a numeric value in bytes with an optional
//line /usr/local/go/src/runtime/debug/garbage.go:183
// unit suffix. The supported suffixes include B, KiB, MiB, GiB, and
//line /usr/local/go/src/runtime/debug/garbage.go:183
// TiB. These suffixes represent quantities of bytes as defined by
//line /usr/local/go/src/runtime/debug/garbage.go:183
// the IEC 80000-13 standard. That is, they are based on powers of
//line /usr/local/go/src/runtime/debug/garbage.go:183
// two: KiB means 2^10 bytes, MiB means 2^20 bytes, and so on.
//line /usr/local/go/src/runtime/debug/garbage.go:183
//
//line /usr/local/go/src/runtime/debug/garbage.go:183
// SetMemoryLimit returns the previously set memory limit.
//line /usr/local/go/src/runtime/debug/garbage.go:183
// A negative input does not adjust the limit, and allows for
//line /usr/local/go/src/runtime/debug/garbage.go:183
// retrieval of the currently set memory limit.
//line /usr/local/go/src/runtime/debug/garbage.go:236
func SetMemoryLimit(limit int64) int64 {
//line /usr/local/go/src/runtime/debug/garbage.go:236
	_go_fuzz_dep_.CoverTab[90733]++
							return setMemoryLimit(limit)
//line /usr/local/go/src/runtime/debug/garbage.go:237
	// _ = "end of CoverTab[90733]"
}

//line /usr/local/go/src/runtime/debug/garbage.go:238
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/runtime/debug/garbage.go:238
var _ = _go_fuzz_dep_.CoverTab
