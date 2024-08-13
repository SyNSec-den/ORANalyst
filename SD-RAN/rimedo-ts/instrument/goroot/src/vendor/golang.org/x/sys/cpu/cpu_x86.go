// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build 386 || amd64 || amd64p32
// +build 386 amd64 amd64p32

//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:8
package cpu

//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:8
import (
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:8
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:8
)
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:8
import (
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:8
	_atomic_ "sync/atomic"
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:8
)

import "runtime"

const cacheLineSize = 64

func initOptions() {
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:14
	_go_fuzz_dep_.CoverTab[20870]++
									options = []option{
										{Name: "adx", Feature: &X86.HasADX},
										{Name: "aes", Feature: &X86.HasAES},
										{Name: "avx", Feature: &X86.HasAVX},
										{Name: "avx2", Feature: &X86.HasAVX2},
										{Name: "avx512", Feature: &X86.HasAVX512},
										{Name: "avx512f", Feature: &X86.HasAVX512F},
										{Name: "avx512cd", Feature: &X86.HasAVX512CD},
										{Name: "avx512er", Feature: &X86.HasAVX512ER},
										{Name: "avx512pf", Feature: &X86.HasAVX512PF},
										{Name: "avx512vl", Feature: &X86.HasAVX512VL},
										{Name: "avx512bw", Feature: &X86.HasAVX512BW},
										{Name: "avx512dq", Feature: &X86.HasAVX512DQ},
										{Name: "avx512ifma", Feature: &X86.HasAVX512IFMA},
										{Name: "avx512vbmi", Feature: &X86.HasAVX512VBMI},
										{Name: "avx512vnniw", Feature: &X86.HasAVX5124VNNIW},
										{Name: "avx5124fmaps", Feature: &X86.HasAVX5124FMAPS},
										{Name: "avx512vpopcntdq", Feature: &X86.HasAVX512VPOPCNTDQ},
										{Name: "avx512vpclmulqdq", Feature: &X86.HasAVX512VPCLMULQDQ},
										{Name: "avx512vnni", Feature: &X86.HasAVX512VNNI},
										{Name: "avx512gfni", Feature: &X86.HasAVX512GFNI},
										{Name: "avx512vaes", Feature: &X86.HasAVX512VAES},
										{Name: "avx512vbmi2", Feature: &X86.HasAVX512VBMI2},
										{Name: "avx512bitalg", Feature: &X86.HasAVX512BITALG},
										{Name: "avx512bf16", Feature: &X86.HasAVX512BF16},
										{Name: "bmi1", Feature: &X86.HasBMI1},
										{Name: "bmi2", Feature: &X86.HasBMI2},
										{Name: "cx16", Feature: &X86.HasCX16},
										{Name: "erms", Feature: &X86.HasERMS},
										{Name: "fma", Feature: &X86.HasFMA},
										{Name: "osxsave", Feature: &X86.HasOSXSAVE},
										{Name: "pclmulqdq", Feature: &X86.HasPCLMULQDQ},
										{Name: "popcnt", Feature: &X86.HasPOPCNT},
										{Name: "rdrand", Feature: &X86.HasRDRAND},
										{Name: "rdseed", Feature: &X86.HasRDSEED},
										{Name: "sse3", Feature: &X86.HasSSE3},
										{Name: "sse41", Feature: &X86.HasSSE41},
										{Name: "sse42", Feature: &X86.HasSSE42},
										{Name: "ssse3", Feature: &X86.HasSSSE3},

//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:56
		{Name: "sse2", Feature: &X86.HasSSE2, Required: runtime.GOARCH == "amd64"},
	}
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:57
	// _ = "end of CoverTab[20870]"
}

func archInit() {
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:60
	_go_fuzz_dep_.CoverTab[20871]++

									Initialized = true

									maxID, _, _, _ := cpuid(0, 0)

									if maxID < 1 {
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:66
		_go_fuzz_dep_.CoverTab[20875]++
										return
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:67
		// _ = "end of CoverTab[20875]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:68
		_go_fuzz_dep_.CoverTab[20876]++
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:68
		// _ = "end of CoverTab[20876]"
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:68
	}
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:68
	// _ = "end of CoverTab[20871]"
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:68
	_go_fuzz_dep_.CoverTab[20872]++

									_, _, ecx1, edx1 := cpuid(1, 0)
									X86.HasSSE2 = isSet(26, edx1)

									X86.HasSSE3 = isSet(0, ecx1)
									X86.HasPCLMULQDQ = isSet(1, ecx1)
									X86.HasSSSE3 = isSet(9, ecx1)
									X86.HasFMA = isSet(12, ecx1)
									X86.HasCX16 = isSet(13, ecx1)
									X86.HasSSE41 = isSet(19, ecx1)
									X86.HasSSE42 = isSet(20, ecx1)
									X86.HasPOPCNT = isSet(23, ecx1)
									X86.HasAES = isSet(25, ecx1)
									X86.HasOSXSAVE = isSet(27, ecx1)
									X86.HasRDRAND = isSet(30, ecx1)

									var osSupportsAVX, osSupportsAVX512 bool

									if X86.HasOSXSAVE {
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:87
		_go_fuzz_dep_.CoverTab[20877]++
										eax, _ := xgetbv()

										osSupportsAVX = isSet(1, eax) && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:90
			_go_fuzz_dep_.CoverTab[20878]++
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:90
			return isSet(2, eax)
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:90
			// _ = "end of CoverTab[20878]"
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:90
		}()

										if runtime.GOOS == "darwin" {
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:92
			_go_fuzz_dep_.CoverTab[20879]++

//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:96
			osSupportsAVX512 = false
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:96
			// _ = "end of CoverTab[20879]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:97
			_go_fuzz_dep_.CoverTab[20880]++

											osSupportsAVX512 = osSupportsAVX && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:99
				_go_fuzz_dep_.CoverTab[20881]++
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:99
				return isSet(5, eax)
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:99
				// _ = "end of CoverTab[20881]"
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:99
			}() && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:99
				_go_fuzz_dep_.CoverTab[20882]++
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:99
				return isSet(6, eax)
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:99
				// _ = "end of CoverTab[20882]"
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:99
			}() && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:99
				_go_fuzz_dep_.CoverTab[20883]++
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:99
				return isSet(7, eax)
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:99
				// _ = "end of CoverTab[20883]"
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:99
			}()
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:99
			// _ = "end of CoverTab[20880]"
		}
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:100
		// _ = "end of CoverTab[20877]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:101
		_go_fuzz_dep_.CoverTab[20884]++
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:101
		// _ = "end of CoverTab[20884]"
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:101
	}
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:101
	// _ = "end of CoverTab[20872]"
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:101
	_go_fuzz_dep_.CoverTab[20873]++

									X86.HasAVX = isSet(28, ecx1) && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:103
		_go_fuzz_dep_.CoverTab[20885]++
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:103
		return osSupportsAVX
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:103
		// _ = "end of CoverTab[20885]"
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:103
	}()

									if maxID < 7 {
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:105
		_go_fuzz_dep_.CoverTab[20886]++
										return
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:106
		// _ = "end of CoverTab[20886]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:107
		_go_fuzz_dep_.CoverTab[20887]++
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:107
		// _ = "end of CoverTab[20887]"
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:107
	}
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:107
	// _ = "end of CoverTab[20873]"
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:107
	_go_fuzz_dep_.CoverTab[20874]++

									_, ebx7, ecx7, edx7 := cpuid(7, 0)
									X86.HasBMI1 = isSet(3, ebx7)
									X86.HasAVX2 = isSet(5, ebx7) && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:111
		_go_fuzz_dep_.CoverTab[20888]++
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:111
		return osSupportsAVX
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:111
		// _ = "end of CoverTab[20888]"
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:111
	}()
									X86.HasBMI2 = isSet(8, ebx7)
									X86.HasERMS = isSet(9, ebx7)
									X86.HasRDSEED = isSet(18, ebx7)
									X86.HasADX = isSet(19, ebx7)

									X86.HasAVX512 = isSet(16, ebx7) && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:117
		_go_fuzz_dep_.CoverTab[20889]++
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:117
		return osSupportsAVX512
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:117
		// _ = "end of CoverTab[20889]"
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:117
	}()
									if X86.HasAVX512 {
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:118
		_go_fuzz_dep_.CoverTab[20890]++
										X86.HasAVX512F = true
										X86.HasAVX512CD = isSet(28, ebx7)
										X86.HasAVX512ER = isSet(27, ebx7)
										X86.HasAVX512PF = isSet(26, ebx7)
										X86.HasAVX512VL = isSet(31, ebx7)
										X86.HasAVX512BW = isSet(30, ebx7)
										X86.HasAVX512DQ = isSet(17, ebx7)
										X86.HasAVX512IFMA = isSet(21, ebx7)
										X86.HasAVX512VBMI = isSet(1, ecx7)
										X86.HasAVX5124VNNIW = isSet(2, edx7)
										X86.HasAVX5124FMAPS = isSet(3, edx7)
										X86.HasAVX512VPOPCNTDQ = isSet(14, ecx7)
										X86.HasAVX512VPCLMULQDQ = isSet(10, ecx7)
										X86.HasAVX512VNNI = isSet(11, ecx7)
										X86.HasAVX512GFNI = isSet(8, ecx7)
										X86.HasAVX512VAES = isSet(9, ecx7)
										X86.HasAVX512VBMI2 = isSet(6, ecx7)
										X86.HasAVX512BITALG = isSet(12, ecx7)

										eax71, _, _, _ := cpuid(7, 1)
										X86.HasAVX512BF16 = isSet(5, eax71)
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:139
		// _ = "end of CoverTab[20890]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:140
		_go_fuzz_dep_.CoverTab[20891]++
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:140
		// _ = "end of CoverTab[20891]"
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:140
	}
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:140
	// _ = "end of CoverTab[20874]"
}

func isSet(bitpos uint, value uint32) bool {
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:143
	_go_fuzz_dep_.CoverTab[20892]++
									return value&(1<<bitpos) != 0
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:144
	// _ = "end of CoverTab[20892]"
}

//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:145
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_x86.go:145
var _ = _go_fuzz_dep_.CoverTab
