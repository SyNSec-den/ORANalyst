//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:1
// Package xxh32 implements the very fast XXH hashing algorithm (32 bits version).
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:1
// (https://github.com/Cyan4973/XXH/)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:3
package xxh32

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:3
import (
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:3
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:3
)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:3
import (
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:3
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:3
)

import (
	"encoding/binary"
)

const (
	prime1	uint32	= 2654435761
	prime2	uint32	= 2246822519
	prime3	uint32	= 3266489917
	prime4	uint32	= 668265263
	prime5	uint32	= 374761393

	primeMask	= 0xFFFFFFFF
	prime1plus2	= uint32((uint64(prime1) + uint64(prime2)) & primeMask)	// 606290984
	prime1minus	= uint32((-int64(prime1)) & primeMask)			// 1640531535
)

// XXHZero represents an xxhash32 object with seed 0.
type XXHZero struct {
	v1		uint32
	v2		uint32
	v3		uint32
	v4		uint32
	totalLen	uint64
	buf		[16]byte
	bufused		int
}

// Sum appends the current hash to b and returns the resulting slice.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:32
// It does not change the underlying hash state.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:34
func (xxh XXHZero) Sum(b []byte) []byte {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:34
	_go_fuzz_dep_.CoverTab[95301]++
														h32 := xxh.Sum32()
														return append(b, byte(h32), byte(h32>>8), byte(h32>>16), byte(h32>>24))
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:36
	// _ = "end of CoverTab[95301]"
}

// Reset resets the Hash to its initial state.
func (xxh *XXHZero) Reset() {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:40
	_go_fuzz_dep_.CoverTab[95302]++
														xxh.v1 = prime1plus2
														xxh.v2 = prime2
														xxh.v3 = 0
														xxh.v4 = prime1minus
														xxh.totalLen = 0
														xxh.bufused = 0
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:46
	// _ = "end of CoverTab[95302]"
}

// Size returns the number of bytes returned by Sum().
func (xxh *XXHZero) Size() int {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:50
	_go_fuzz_dep_.CoverTab[95303]++
														return 4
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:51
	// _ = "end of CoverTab[95303]"
}

// BlockSize gives the minimum number of bytes accepted by Write().
func (xxh *XXHZero) BlockSize() int {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:55
	_go_fuzz_dep_.CoverTab[95304]++
														return 1
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:56
	// _ = "end of CoverTab[95304]"
}

// Write adds input bytes to the Hash.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:59
// It never returns an error.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:61
func (xxh *XXHZero) Write(input []byte) (int, error) {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:61
	_go_fuzz_dep_.CoverTab[95305]++
														if xxh.totalLen == 0 {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:62
		_go_fuzz_dep_.CoverTab[95310]++
															xxh.Reset()
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:63
		// _ = "end of CoverTab[95310]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:64
		_go_fuzz_dep_.CoverTab[95311]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:64
		// _ = "end of CoverTab[95311]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:64
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:64
	// _ = "end of CoverTab[95305]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:64
	_go_fuzz_dep_.CoverTab[95306]++
														n := len(input)
														m := xxh.bufused

														xxh.totalLen += uint64(n)

														r := len(xxh.buf) - m
														if n < r {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:71
		_go_fuzz_dep_.CoverTab[95312]++
															copy(xxh.buf[m:], input)
															xxh.bufused += len(input)
															return n, nil
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:74
		// _ = "end of CoverTab[95312]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:75
		_go_fuzz_dep_.CoverTab[95313]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:75
		// _ = "end of CoverTab[95313]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:75
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:75
	// _ = "end of CoverTab[95306]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:75
	_go_fuzz_dep_.CoverTab[95307]++

														p := 0

														v1, v2, v3, v4 := xxh.v1, xxh.v2, xxh.v3, xxh.v4
														if m > 0 {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:80
		_go_fuzz_dep_.CoverTab[95314]++

															copy(xxh.buf[xxh.bufused:], input[:r])
															xxh.bufused += len(input) - r

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:86
		buf := xxh.buf[:16]
															v1 = rol13(v1+binary.LittleEndian.Uint32(buf[:])*prime2) * prime1
															v2 = rol13(v2+binary.LittleEndian.Uint32(buf[4:])*prime2) * prime1
															v3 = rol13(v3+binary.LittleEndian.Uint32(buf[8:])*prime2) * prime1
															v4 = rol13(v4+binary.LittleEndian.Uint32(buf[12:])*prime2) * prime1
															p = r
															xxh.bufused = 0
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:92
		// _ = "end of CoverTab[95314]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:93
		_go_fuzz_dep_.CoverTab[95315]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:93
		// _ = "end of CoverTab[95315]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:93
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:93
	// _ = "end of CoverTab[95307]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:93
	_go_fuzz_dep_.CoverTab[95308]++

														for n := n - 16; p <= n; p += 16 {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:95
		_go_fuzz_dep_.CoverTab[95316]++
															sub := input[p:][:16]
															v1 = rol13(v1+binary.LittleEndian.Uint32(sub[:])*prime2) * prime1
															v2 = rol13(v2+binary.LittleEndian.Uint32(sub[4:])*prime2) * prime1
															v3 = rol13(v3+binary.LittleEndian.Uint32(sub[8:])*prime2) * prime1
															v4 = rol13(v4+binary.LittleEndian.Uint32(sub[12:])*prime2) * prime1
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:100
		// _ = "end of CoverTab[95316]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:101
	// _ = "end of CoverTab[95308]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:101
	_go_fuzz_dep_.CoverTab[95309]++
														xxh.v1, xxh.v2, xxh.v3, xxh.v4 = v1, v2, v3, v4

														copy(xxh.buf[xxh.bufused:], input[p:])
														xxh.bufused += len(input) - p

														return n, nil
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:107
	// _ = "end of CoverTab[95309]"
}

// Sum32 returns the 32 bits Hash value.
func (xxh *XXHZero) Sum32() uint32 {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:111
	_go_fuzz_dep_.CoverTab[95317]++
														h32 := uint32(xxh.totalLen)
														if h32 >= 16 {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:113
		_go_fuzz_dep_.CoverTab[95321]++
															h32 += rol1(xxh.v1) + rol7(xxh.v2) + rol12(xxh.v3) + rol18(xxh.v4)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:114
		// _ = "end of CoverTab[95321]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:115
		_go_fuzz_dep_.CoverTab[95322]++
															h32 += prime5
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:116
		// _ = "end of CoverTab[95322]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:117
	// _ = "end of CoverTab[95317]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:117
	_go_fuzz_dep_.CoverTab[95318]++

														p := 0
														n := xxh.bufused
														buf := xxh.buf
														for n := n - 4; p <= n; p += 4 {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:122
		_go_fuzz_dep_.CoverTab[95323]++
															h32 += binary.LittleEndian.Uint32(buf[p:p+4]) * prime3
															h32 = rol17(h32) * prime4
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:124
		// _ = "end of CoverTab[95323]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:125
	// _ = "end of CoverTab[95318]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:125
	_go_fuzz_dep_.CoverTab[95319]++
														for ; p < n; p++ {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:126
		_go_fuzz_dep_.CoverTab[95324]++
															h32 += uint32(buf[p]) * prime5
															h32 = rol11(h32) * prime1
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:128
		// _ = "end of CoverTab[95324]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:129
	// _ = "end of CoverTab[95319]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:129
	_go_fuzz_dep_.CoverTab[95320]++

														h32 ^= h32 >> 15
														h32 *= prime2
														h32 ^= h32 >> 13
														h32 *= prime3
														h32 ^= h32 >> 16

														return h32
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:137
	// _ = "end of CoverTab[95320]"
}

// ChecksumZero returns the 32bits Hash value.
func ChecksumZero(input []byte) uint32 {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:141
	_go_fuzz_dep_.CoverTab[95325]++
														n := len(input)
														h32 := uint32(n)

														if n < 16 {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:145
		_go_fuzz_dep_.CoverTab[95329]++
															h32 += prime5
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:146
		// _ = "end of CoverTab[95329]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:147
		_go_fuzz_dep_.CoverTab[95330]++
															v1 := prime1plus2
															v2 := prime2
															v3 := uint32(0)
															v4 := prime1minus
															p := 0
															for n := n - 16; p <= n; p += 16 {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:153
			_go_fuzz_dep_.CoverTab[95332]++
																sub := input[p:][:16]
																v1 = rol13(v1+binary.LittleEndian.Uint32(sub[:])*prime2) * prime1
																v2 = rol13(v2+binary.LittleEndian.Uint32(sub[4:])*prime2) * prime1
																v3 = rol13(v3+binary.LittleEndian.Uint32(sub[8:])*prime2) * prime1
																v4 = rol13(v4+binary.LittleEndian.Uint32(sub[12:])*prime2) * prime1
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:158
			// _ = "end of CoverTab[95332]"
		}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:159
		// _ = "end of CoverTab[95330]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:159
		_go_fuzz_dep_.CoverTab[95331]++
															input = input[p:]
															n -= p
															h32 += rol1(v1) + rol7(v2) + rol12(v3) + rol18(v4)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:162
		// _ = "end of CoverTab[95331]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:163
	// _ = "end of CoverTab[95325]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:163
	_go_fuzz_dep_.CoverTab[95326]++

														p := 0
														for n := n - 4; p <= n; p += 4 {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:166
		_go_fuzz_dep_.CoverTab[95333]++
															h32 += binary.LittleEndian.Uint32(input[p:p+4]) * prime3
															h32 = rol17(h32) * prime4
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:168
		// _ = "end of CoverTab[95333]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:169
	// _ = "end of CoverTab[95326]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:169
	_go_fuzz_dep_.CoverTab[95327]++
														for p < n {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:170
		_go_fuzz_dep_.CoverTab[95334]++
															h32 += uint32(input[p]) * prime5
															h32 = rol11(h32) * prime1
															p++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:173
		// _ = "end of CoverTab[95334]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:174
	// _ = "end of CoverTab[95327]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:174
	_go_fuzz_dep_.CoverTab[95328]++

														h32 ^= h32 >> 15
														h32 *= prime2
														h32 ^= h32 >> 13
														h32 *= prime3
														h32 ^= h32 >> 16

														return h32
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:182
	// _ = "end of CoverTab[95328]"
}

// Uint32Zero hashes x with seed 0.
func Uint32Zero(x uint32) uint32 {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:186
	_go_fuzz_dep_.CoverTab[95335]++
														h := prime5 + 4 + x*prime3
														h = rol17(h) * prime4
														h ^= h >> 15
														h *= prime2
														h ^= h >> 13
														h *= prime3
														h ^= h >> 16
														return h
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:194
	// _ = "end of CoverTab[95335]"
}

func rol1(u uint32) uint32 {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:197
	_go_fuzz_dep_.CoverTab[95336]++
														return u<<1 | u>>31
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:198
	// _ = "end of CoverTab[95336]"
}

func rol7(u uint32) uint32 {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:201
	_go_fuzz_dep_.CoverTab[95337]++
														return u<<7 | u>>25
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:202
	// _ = "end of CoverTab[95337]"
}

func rol11(u uint32) uint32 {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:205
	_go_fuzz_dep_.CoverTab[95338]++
														return u<<11 | u>>21
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:206
	// _ = "end of CoverTab[95338]"
}

func rol12(u uint32) uint32 {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:209
	_go_fuzz_dep_.CoverTab[95339]++
														return u<<12 | u>>20
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:210
	// _ = "end of CoverTab[95339]"
}

func rol13(u uint32) uint32 {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:213
	_go_fuzz_dep_.CoverTab[95340]++
														return u<<13 | u>>19
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:214
	// _ = "end of CoverTab[95340]"
}

func rol17(u uint32) uint32 {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:217
	_go_fuzz_dep_.CoverTab[95341]++
														return u<<17 | u>>15
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:218
	// _ = "end of CoverTab[95341]"
}

func rol18(u uint32) uint32 {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:221
	_go_fuzz_dep_.CoverTab[95342]++
														return u<<18 | u>>14
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:222
	// _ = "end of CoverTab[95342]"
}

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:223
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/internal/xxh32/xxh32zero.go:223
var _ = _go_fuzz_dep_.CoverTab
