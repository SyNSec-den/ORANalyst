// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// CPU affinity functions

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/affinity_linux.go:7
package unix

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/affinity_linux.go:7
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/affinity_linux.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/affinity_linux.go:7
)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/affinity_linux.go:7
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/affinity_linux.go:7
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/affinity_linux.go:7
)

import (
	"math/bits"
	"unsafe"
)

const cpuSetSize = _CPU_SETSIZE / _NCPUBITS

// CPUSet represents a CPU affinity mask.
type CPUSet [cpuSetSize]cpuMask

func schedAffinity(trap uintptr, pid int, set *CPUSet) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/affinity_linux.go:19
	_go_fuzz_dep_.CoverTab[45696]++
											_, _, e := RawSyscall(trap, uintptr(pid), uintptr(unsafe.Sizeof(*set)), uintptr(unsafe.Pointer(set)))
											if e != 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/affinity_linux.go:21
		_go_fuzz_dep_.CoverTab[45698]++
												return errnoErr(e)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/affinity_linux.go:22
		// _ = "end of CoverTab[45698]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/affinity_linux.go:23
		_go_fuzz_dep_.CoverTab[45699]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/affinity_linux.go:23
		// _ = "end of CoverTab[45699]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/affinity_linux.go:23
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/affinity_linux.go:23
	// _ = "end of CoverTab[45696]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/affinity_linux.go:23
	_go_fuzz_dep_.CoverTab[45697]++
											return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/affinity_linux.go:24
	// _ = "end of CoverTab[45697]"
}

// SchedGetaffinity gets the CPU affinity mask of the thread specified by pid.
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/affinity_linux.go:27
// If pid is 0 the calling thread is used.
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/affinity_linux.go:29
func SchedGetaffinity(pid int, set *CPUSet) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/affinity_linux.go:29
	_go_fuzz_dep_.CoverTab[45700]++
											return schedAffinity(SYS_SCHED_GETAFFINITY, pid, set)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/affinity_linux.go:30
	// _ = "end of CoverTab[45700]"
}

// SchedSetaffinity sets the CPU affinity mask of the thread specified by pid.
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/affinity_linux.go:33
// If pid is 0 the calling thread is used.
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/affinity_linux.go:35
func SchedSetaffinity(pid int, set *CPUSet) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/affinity_linux.go:35
	_go_fuzz_dep_.CoverTab[45701]++
											return schedAffinity(SYS_SCHED_SETAFFINITY, pid, set)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/affinity_linux.go:36
	// _ = "end of CoverTab[45701]"
}

// Zero clears the set s, so that it contains no CPUs.
func (s *CPUSet) Zero() {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/affinity_linux.go:40
	_go_fuzz_dep_.CoverTab[45702]++
											for i := range s {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/affinity_linux.go:41
		_go_fuzz_dep_.CoverTab[45703]++
												s[i] = 0
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/affinity_linux.go:42
		// _ = "end of CoverTab[45703]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/affinity_linux.go:43
	// _ = "end of CoverTab[45702]"
}

func cpuBitsIndex(cpu int) int {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/affinity_linux.go:46
	_go_fuzz_dep_.CoverTab[45704]++
											return cpu / _NCPUBITS
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/affinity_linux.go:47
	// _ = "end of CoverTab[45704]"
}

func cpuBitsMask(cpu int) cpuMask {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/affinity_linux.go:50
	_go_fuzz_dep_.CoverTab[45705]++
											return cpuMask(1 << (uint(cpu) % _NCPUBITS))
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/affinity_linux.go:51
	// _ = "end of CoverTab[45705]"
}

// Set adds cpu to the set s.
func (s *CPUSet) Set(cpu int) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/affinity_linux.go:55
	_go_fuzz_dep_.CoverTab[45706]++
											i := cpuBitsIndex(cpu)
											if i < len(s) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/affinity_linux.go:57
		_go_fuzz_dep_.CoverTab[45707]++
												s[i] |= cpuBitsMask(cpu)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/affinity_linux.go:58
		// _ = "end of CoverTab[45707]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/affinity_linux.go:59
		_go_fuzz_dep_.CoverTab[45708]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/affinity_linux.go:59
		// _ = "end of CoverTab[45708]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/affinity_linux.go:59
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/affinity_linux.go:59
	// _ = "end of CoverTab[45706]"
}

// Clear removes cpu from the set s.
func (s *CPUSet) Clear(cpu int) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/affinity_linux.go:63
	_go_fuzz_dep_.CoverTab[45709]++
											i := cpuBitsIndex(cpu)
											if i < len(s) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/affinity_linux.go:65
		_go_fuzz_dep_.CoverTab[45710]++
												s[i] &^= cpuBitsMask(cpu)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/affinity_linux.go:66
		// _ = "end of CoverTab[45710]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/affinity_linux.go:67
		_go_fuzz_dep_.CoverTab[45711]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/affinity_linux.go:67
		// _ = "end of CoverTab[45711]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/affinity_linux.go:67
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/affinity_linux.go:67
	// _ = "end of CoverTab[45709]"
}

// IsSet reports whether cpu is in the set s.
func (s *CPUSet) IsSet(cpu int) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/affinity_linux.go:71
	_go_fuzz_dep_.CoverTab[45712]++
											i := cpuBitsIndex(cpu)
											if i < len(s) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/affinity_linux.go:73
		_go_fuzz_dep_.CoverTab[45714]++
												return s[i]&cpuBitsMask(cpu) != 0
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/affinity_linux.go:74
		// _ = "end of CoverTab[45714]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/affinity_linux.go:75
		_go_fuzz_dep_.CoverTab[45715]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/affinity_linux.go:75
		// _ = "end of CoverTab[45715]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/affinity_linux.go:75
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/affinity_linux.go:75
	// _ = "end of CoverTab[45712]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/affinity_linux.go:75
	_go_fuzz_dep_.CoverTab[45713]++
											return false
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/affinity_linux.go:76
	// _ = "end of CoverTab[45713]"
}

// Count returns the number of CPUs in the set s.
func (s *CPUSet) Count() int {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/affinity_linux.go:80
	_go_fuzz_dep_.CoverTab[45716]++
											c := 0
											for _, b := range s {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/affinity_linux.go:82
		_go_fuzz_dep_.CoverTab[45718]++
												c += bits.OnesCount64(uint64(b))
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/affinity_linux.go:83
		// _ = "end of CoverTab[45718]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/affinity_linux.go:84
	// _ = "end of CoverTab[45716]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/affinity_linux.go:84
	_go_fuzz_dep_.CoverTab[45717]++
											return c
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/affinity_linux.go:85
	// _ = "end of CoverTab[45717]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/affinity_linux.go:86
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/affinity_linux.go:86
var _ = _go_fuzz_dep_.CoverTab
