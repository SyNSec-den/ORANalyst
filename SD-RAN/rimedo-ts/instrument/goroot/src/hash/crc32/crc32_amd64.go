// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// AMD64-specific hardware-assisted CRC32 algorithms. See crc32.go for a
// description of the interface that each architecture-specific file
// implements.

//line /usr/local/go/src/hash/crc32/crc32_amd64.go:9
package crc32

//line /usr/local/go/src/hash/crc32/crc32_amd64.go:9
import (
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:9
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:9
)
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:9
import (
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:9
	_atomic_ "sync/atomic"
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:9
)

import (
	"internal/cpu"
	"unsafe"
)

//line /usr/local/go/src/hash/crc32/crc32_amd64.go:22
//go:noescape
func castagnoliSSE42(crc uint32, p []byte) uint32

//line /usr/local/go/src/hash/crc32/crc32_amd64.go:28
//go:noescape
func castagnoliSSE42Triple(
	crcA, crcB, crcC uint32,
	a, b, c []byte,
	rounds uint32,
) (retA uint32, retB uint32, retC uint32)

//line /usr/local/go/src/hash/crc32/crc32_amd64.go:38
//go:noescape
func ieeeCLMUL(crc uint32, p []byte) uint32

const castagnoliK1 = 168
const castagnoliK2 = 1344

type sse42Table [4]Table

var castagnoliSSE42TableK1 *sse42Table
var castagnoliSSE42TableK2 *sse42Table

func archAvailableCastagnoli() bool {
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:49
	_go_fuzz_dep_.CoverTab[26593]++
							return cpu.X86.HasSSE42
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:50
	// _ = "end of CoverTab[26593]"
}

func archInitCastagnoli() {
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:53
	_go_fuzz_dep_.CoverTab[26594]++
							if !cpu.X86.HasSSE42 {
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:54
		_go_fuzz_dep_.CoverTab[26596]++
								panic("arch-specific Castagnoli not available")
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:55
		// _ = "end of CoverTab[26596]"
	} else {
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:56
		_go_fuzz_dep_.CoverTab[26597]++
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:56
		// _ = "end of CoverTab[26597]"
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:56
	}
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:56
	// _ = "end of CoverTab[26594]"
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:56
	_go_fuzz_dep_.CoverTab[26595]++
							castagnoliSSE42TableK1 = new(sse42Table)
							castagnoliSSE42TableK2 = new(sse42Table)

//line /usr/local/go/src/hash/crc32/crc32_amd64.go:65
	var tmp [castagnoliK2]byte
	for b := 0; b < 4; b++ {
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:66
		_go_fuzz_dep_.CoverTab[26598]++
								for i := 0; i < 256; i++ {
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:67
			_go_fuzz_dep_.CoverTab[26599]++
									val := uint32(i) << uint32(b*8)
									castagnoliSSE42TableK1[b][i] = castagnoliSSE42(val, tmp[:castagnoliK1])
									castagnoliSSE42TableK2[b][i] = castagnoliSSE42(val, tmp[:])
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:70
			// _ = "end of CoverTab[26599]"
		}
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:71
		// _ = "end of CoverTab[26598]"
	}
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:72
	// _ = "end of CoverTab[26595]"
}

//line /usr/local/go/src/hash/crc32/crc32_amd64.go:78
func castagnoliShift(table *sse42Table, crc uint32) uint32 {
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:78
	_go_fuzz_dep_.CoverTab[26600]++
							return table[3][crc>>24] ^
		table[2][(crc>>16)&0xFF] ^
		table[1][(crc>>8)&0xFF] ^
		table[0][crc&0xFF]
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:82
	// _ = "end of CoverTab[26600]"
}

func archUpdateCastagnoli(crc uint32, p []byte) uint32 {
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:85
	_go_fuzz_dep_.CoverTab[26601]++
							if !cpu.X86.HasSSE42 {
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:86
		_go_fuzz_dep_.CoverTab[26606]++
								panic("not available")
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:87
		// _ = "end of CoverTab[26606]"
	} else {
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:88
		_go_fuzz_dep_.CoverTab[26607]++
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:88
		// _ = "end of CoverTab[26607]"
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:88
	}
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:88
	// _ = "end of CoverTab[26601]"
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:88
	_go_fuzz_dep_.CoverTab[26602]++

//line /usr/local/go/src/hash/crc32/crc32_amd64.go:148
	crc = ^crc

//line /usr/local/go/src/hash/crc32/crc32_amd64.go:152
	if len(p) >= castagnoliK1*3 {
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:152
		_go_fuzz_dep_.CoverTab[26608]++
								delta := int(uintptr(unsafe.Pointer(&p[0])) & 7)
								if delta != 0 {
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:154
			_go_fuzz_dep_.CoverTab[26609]++
									delta = 8 - delta
									crc = castagnoliSSE42(crc, p[:delta])
									p = p[delta:]
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:157
			// _ = "end of CoverTab[26609]"
		} else {
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:158
			_go_fuzz_dep_.CoverTab[26610]++
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:158
			// _ = "end of CoverTab[26610]"
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:158
		}
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:158
		// _ = "end of CoverTab[26608]"
	} else {
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:159
		_go_fuzz_dep_.CoverTab[26611]++
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:159
		// _ = "end of CoverTab[26611]"
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:159
	}
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:159
	// _ = "end of CoverTab[26602]"
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:159
	_go_fuzz_dep_.CoverTab[26603]++

//line /usr/local/go/src/hash/crc32/crc32_amd64.go:162
	for len(p) >= castagnoliK2*3 {
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:162
		_go_fuzz_dep_.CoverTab[26612]++

								crcA, crcB, crcC := castagnoliSSE42Triple(
			crc, 0, 0,
			p, p[castagnoliK2:], p[castagnoliK2*2:],
			castagnoliK2/24)

//line /usr/local/go/src/hash/crc32/crc32_amd64.go:170
		crcAB := castagnoliShift(castagnoliSSE42TableK2, crcA) ^ crcB

								crc = castagnoliShift(castagnoliSSE42TableK2, crcAB) ^ crcC
								p = p[castagnoliK2*3:]
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:173
		// _ = "end of CoverTab[26612]"
	}
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:174
	// _ = "end of CoverTab[26603]"
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:174
	_go_fuzz_dep_.CoverTab[26604]++

//line /usr/local/go/src/hash/crc32/crc32_amd64.go:177
	for len(p) >= castagnoliK1*3 {
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:177
		_go_fuzz_dep_.CoverTab[26613]++

								crcA, crcB, crcC := castagnoliSSE42Triple(
			crc, 0, 0,
			p, p[castagnoliK1:], p[castagnoliK1*2:],
			castagnoliK1/24)

//line /usr/local/go/src/hash/crc32/crc32_amd64.go:185
		crcAB := castagnoliShift(castagnoliSSE42TableK1, crcA) ^ crcB

								crc = castagnoliShift(castagnoliSSE42TableK1, crcAB) ^ crcC
								p = p[castagnoliK1*3:]
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:188
		// _ = "end of CoverTab[26613]"
	}
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:189
	// _ = "end of CoverTab[26604]"
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:189
	_go_fuzz_dep_.CoverTab[26605]++

//line /usr/local/go/src/hash/crc32/crc32_amd64.go:192
	crc = castagnoliSSE42(crc, p)
							return ^crc
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:193
	// _ = "end of CoverTab[26605]"
}

func archAvailableIEEE() bool {
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:196
	_go_fuzz_dep_.CoverTab[26614]++
							return cpu.X86.HasPCLMULQDQ && func() bool {
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:197
		_go_fuzz_dep_.CoverTab[26615]++
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:197
		return cpu.X86.HasSSE41
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:197
		// _ = "end of CoverTab[26615]"
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:197
	}()
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:197
	// _ = "end of CoverTab[26614]"
}

var archIeeeTable8 *slicing8Table

func archInitIEEE() {
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:202
	_go_fuzz_dep_.CoverTab[26616]++
							if !cpu.X86.HasPCLMULQDQ || func() bool {
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:203
		_go_fuzz_dep_.CoverTab[26618]++
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:203
		return !cpu.X86.HasSSE41
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:203
		// _ = "end of CoverTab[26618]"
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:203
	}() {
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:203
		_go_fuzz_dep_.CoverTab[26619]++
								panic("not available")
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:204
		// _ = "end of CoverTab[26619]"
	} else {
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:205
		_go_fuzz_dep_.CoverTab[26620]++
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:205
		// _ = "end of CoverTab[26620]"
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:205
	}
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:205
	// _ = "end of CoverTab[26616]"
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:205
	_go_fuzz_dep_.CoverTab[26617]++

							archIeeeTable8 = slicingMakeTable(IEEE)
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:207
	// _ = "end of CoverTab[26617]"
}

func archUpdateIEEE(crc uint32, p []byte) uint32 {
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:210
	_go_fuzz_dep_.CoverTab[26621]++
							if !cpu.X86.HasPCLMULQDQ || func() bool {
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:211
		_go_fuzz_dep_.CoverTab[26625]++
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:211
		return !cpu.X86.HasSSE41
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:211
		// _ = "end of CoverTab[26625]"
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:211
	}() {
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:211
		_go_fuzz_dep_.CoverTab[26626]++
								panic("not available")
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:212
		// _ = "end of CoverTab[26626]"
	} else {
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:213
		_go_fuzz_dep_.CoverTab[26627]++
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:213
		// _ = "end of CoverTab[26627]"
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:213
	}
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:213
	// _ = "end of CoverTab[26621]"
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:213
	_go_fuzz_dep_.CoverTab[26622]++

							if len(p) >= 64 {
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:215
		_go_fuzz_dep_.CoverTab[26628]++
								left := len(p) & 15
								do := len(p) - left
								crc = ^ieeeCLMUL(^crc, p[:do])
								p = p[do:]
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:219
		// _ = "end of CoverTab[26628]"
	} else {
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:220
		_go_fuzz_dep_.CoverTab[26629]++
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:220
		// _ = "end of CoverTab[26629]"
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:220
	}
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:220
	// _ = "end of CoverTab[26622]"
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:220
	_go_fuzz_dep_.CoverTab[26623]++
							if len(p) == 0 {
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:221
		_go_fuzz_dep_.CoverTab[26630]++
								return crc
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:222
		// _ = "end of CoverTab[26630]"
	} else {
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:223
		_go_fuzz_dep_.CoverTab[26631]++
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:223
		// _ = "end of CoverTab[26631]"
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:223
	}
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:223
	// _ = "end of CoverTab[26623]"
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:223
	_go_fuzz_dep_.CoverTab[26624]++
							return slicingUpdate(crc, archIeeeTable8, p)
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:224
	// _ = "end of CoverTab[26624]"
}

//line /usr/local/go/src/hash/crc32/crc32_amd64.go:225
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/hash/crc32/crc32_amd64.go:225
var _ = _go_fuzz_dep_.CoverTab
