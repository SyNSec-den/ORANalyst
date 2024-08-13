// Copyright 2016 The Snappy-Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:5
package snapref

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:5
)

func load32(b []byte, i int) uint32 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:7
	_go_fuzz_dep_.CoverTab[90611]++
														b = b[i : i+4 : len(b)]
														return uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:9
	// _ = "end of CoverTab[90611]"
}

func load64(b []byte, i int) uint64 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:12
	_go_fuzz_dep_.CoverTab[90612]++
														b = b[i : i+8 : len(b)]
														return uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 |
		uint64(b[4])<<32 | uint64(b[5])<<40 | uint64(b[6])<<48 | uint64(b[7])<<56
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:15
	// _ = "end of CoverTab[90612]"
}

// emitLiteral writes a literal chunk and returns the number of bytes written.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:18
//
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:18
// It assumes that:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:18
//
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:18
//	dst is long enough to hold the encoded bytes
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:18
//	1 <= len(lit) && len(lit) <= 65536
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:23
func emitLiteral(dst, lit []byte) int {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:23
	_go_fuzz_dep_.CoverTab[90613]++
														i, n := 0, uint(len(lit)-1)
														switch {
	case n < 60:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:26
		_go_fuzz_dep_.CoverTab[90615]++
															dst[0] = uint8(n)<<2 | tagLiteral
															i = 1
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:28
		// _ = "end of CoverTab[90615]"
	case n < 1<<8:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:29
		_go_fuzz_dep_.CoverTab[90616]++
															dst[0] = 60<<2 | tagLiteral
															dst[1] = uint8(n)
															i = 2
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:32
		// _ = "end of CoverTab[90616]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:33
		_go_fuzz_dep_.CoverTab[90617]++
															dst[0] = 61<<2 | tagLiteral
															dst[1] = uint8(n)
															dst[2] = uint8(n >> 8)
															i = 3
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:37
		// _ = "end of CoverTab[90617]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:38
	// _ = "end of CoverTab[90613]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:38
	_go_fuzz_dep_.CoverTab[90614]++
														return i + copy(dst[i:], lit)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:39
	// _ = "end of CoverTab[90614]"
}

// emitCopy writes a copy chunk and returns the number of bytes written.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:42
//
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:42
// It assumes that:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:42
//
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:42
//	dst is long enough to hold the encoded bytes
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:42
//	1 <= offset && offset <= 65535
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:42
//	4 <= length && length <= 65535
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:48
func emitCopy(dst []byte, offset, length int) int {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:48
	_go_fuzz_dep_.CoverTab[90618]++
														i := 0

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:59
	for length >= 68 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:59
		_go_fuzz_dep_.CoverTab[90622]++

															dst[i+0] = 63<<2 | tagCopy2
															dst[i+1] = uint8(offset)
															dst[i+2] = uint8(offset >> 8)
															i += 3
															length -= 64
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:65
		// _ = "end of CoverTab[90622]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:66
	// _ = "end of CoverTab[90618]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:66
	_go_fuzz_dep_.CoverTab[90619]++
														if length > 64 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:67
		_go_fuzz_dep_.CoverTab[90623]++

															dst[i+0] = 59<<2 | tagCopy2
															dst[i+1] = uint8(offset)
															dst[i+2] = uint8(offset >> 8)
															i += 3
															length -= 60
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:73
		// _ = "end of CoverTab[90623]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:74
		_go_fuzz_dep_.CoverTab[90624]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:74
		// _ = "end of CoverTab[90624]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:74
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:74
	// _ = "end of CoverTab[90619]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:74
	_go_fuzz_dep_.CoverTab[90620]++
														if length >= 12 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:75
		_go_fuzz_dep_.CoverTab[90625]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:75
		return offset >= 2048
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:75
		// _ = "end of CoverTab[90625]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:75
	}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:75
		_go_fuzz_dep_.CoverTab[90626]++

															dst[i+0] = uint8(length-1)<<2 | tagCopy2
															dst[i+1] = uint8(offset)
															dst[i+2] = uint8(offset >> 8)
															return i + 3
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:80
		// _ = "end of CoverTab[90626]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:81
		_go_fuzz_dep_.CoverTab[90627]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:81
		// _ = "end of CoverTab[90627]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:81
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:81
	// _ = "end of CoverTab[90620]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:81
	_go_fuzz_dep_.CoverTab[90621]++

														dst[i+0] = uint8(offset>>8)<<5 | uint8(length-4)<<2 | tagCopy1
														dst[i+1] = uint8(offset)
														return i + 2
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:85
	// _ = "end of CoverTab[90621]"
}

// extendMatch returns the largest k such that k <= len(src) and that
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:88
// src[i:i+k-j] and src[j:k] have the same contents.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:88
//
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:88
// It assumes that:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:88
//
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:88
//	0 <= i && i < j && j <= len(src)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:93
func extendMatch(src []byte, i, j int) int {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:93
	_go_fuzz_dep_.CoverTab[90628]++
														for ; j < len(src) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:94
		_go_fuzz_dep_.CoverTab[90630]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:94
		return src[i] == src[j]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:94
		// _ = "end of CoverTab[90630]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:94
	}(); i, j = i+1, j+1 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:94
		_go_fuzz_dep_.CoverTab[90631]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:94
		// _ = "end of CoverTab[90631]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:95
	// _ = "end of CoverTab[90628]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:95
	_go_fuzz_dep_.CoverTab[90629]++
														return j
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:96
	// _ = "end of CoverTab[90629]"
}

func hash(u, shift uint32) uint32 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:99
	_go_fuzz_dep_.CoverTab[90632]++
														return (u * 0x1e35a7bd) >> shift
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:100
	// _ = "end of CoverTab[90632]"
}

// encodeBlock encodes a non-empty src to a guaranteed-large-enough dst. It
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:103
// assumes that the varint-encoded length of the decompressed bytes has already
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:103
// been written.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:103
//
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:103
// It also assumes that:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:103
//
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:103
//	len(dst) >= MaxEncodedLen(len(src)) &&
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:103
//	minNonLiteralBlockSize <= len(src) && len(src) <= maxBlockSize
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:110
func encodeBlock(dst, src []byte) (d int) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:110
	_go_fuzz_dep_.CoverTab[90633]++
	// Initialize the hash table. Its size ranges from 1<<8 to 1<<14 inclusive.
	// The table element type is uint16, as s < sLimit and sLimit < len(src)
	// and len(src) <= maxBlockSize and maxBlockSize == 65536.
	const (
		maxTableSize	= 1 << 14
		// tableMask is redundant, but helps the compiler eliminate bounds
		// checks.
		tableMask	= maxTableSize - 1
	)
	shift := uint32(32 - 8)
	for tableSize := 1 << 8; tableSize < maxTableSize && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:121
		_go_fuzz_dep_.CoverTab[90637]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:121
		return tableSize < len(src)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:121
		// _ = "end of CoverTab[90637]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:121
	}(); tableSize *= 2 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:121
		_go_fuzz_dep_.CoverTab[90638]++
															shift--
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:122
		// _ = "end of CoverTab[90638]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:123
	// _ = "end of CoverTab[90633]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:123
	_go_fuzz_dep_.CoverTab[90634]++
	// In Go, all array elements are zero-initialized, so there is no advantage
	// to a smaller tableSize per se. However, it matches the C++ algorithm,
	// and in the asm versions of this code, we can get away with zeroing only
														// the first tableSize elements.
														var table [maxTableSize]uint16

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:133
	sLimit := len(src) - inputMargin

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:136
	nextEmit := 0

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:140
	s := 1
	nextHash := hash(load32(src, s), shift)

	for {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:143
		_go_fuzz_dep_.CoverTab[90639]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:159
		skip := 32

		nextS := s
		candidate := 0
		for {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:163
			_go_fuzz_dep_.CoverTab[90641]++
																s = nextS
																bytesBetweenHashLookups := skip >> 5
																nextS = s + bytesBetweenHashLookups
																skip += bytesBetweenHashLookups
																if nextS > sLimit {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:168
				_go_fuzz_dep_.CoverTab[90643]++
																	goto emitRemainder
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:169
				// _ = "end of CoverTab[90643]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:170
				_go_fuzz_dep_.CoverTab[90644]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:170
				// _ = "end of CoverTab[90644]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:170
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:170
			// _ = "end of CoverTab[90641]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:170
			_go_fuzz_dep_.CoverTab[90642]++
																candidate = int(table[nextHash&tableMask])
																table[nextHash&tableMask] = uint16(s)
																nextHash = hash(load32(src, nextS), shift)
																if load32(src, s) == load32(src, candidate) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:174
				_go_fuzz_dep_.CoverTab[90645]++
																	break
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:175
				// _ = "end of CoverTab[90645]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:176
				_go_fuzz_dep_.CoverTab[90646]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:176
				// _ = "end of CoverTab[90646]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:176
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:176
			// _ = "end of CoverTab[90642]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:177
		// _ = "end of CoverTab[90639]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:177
		_go_fuzz_dep_.CoverTab[90640]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:182
		d += emitLiteral(dst[d:], src[nextEmit:s])

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:192
		for {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:192
			_go_fuzz_dep_.CoverTab[90647]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:195
			base := s

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:201
			s += 4
			for i := candidate + 4; s < len(src) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:202
				_go_fuzz_dep_.CoverTab[90650]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:202
				return src[i] == src[s]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:202
				// _ = "end of CoverTab[90650]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:202
			}(); i, s = i+1, s+1 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:202
				_go_fuzz_dep_.CoverTab[90651]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:202
				// _ = "end of CoverTab[90651]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:203
			// _ = "end of CoverTab[90647]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:203
			_go_fuzz_dep_.CoverTab[90648]++

																d += emitCopy(dst[d:], base-candidate, s-base)
																nextEmit = s
																if s >= sLimit {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:207
				_go_fuzz_dep_.CoverTab[90652]++
																	goto emitRemainder
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:208
				// _ = "end of CoverTab[90652]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:209
				_go_fuzz_dep_.CoverTab[90653]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:209
				// _ = "end of CoverTab[90653]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:209
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:209
			// _ = "end of CoverTab[90648]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:209
			_go_fuzz_dep_.CoverTab[90649]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:217
			x := load64(src, s-1)
			prevHash := hash(uint32(x>>0), shift)
			table[prevHash&tableMask] = uint16(s - 1)
			currHash := hash(uint32(x>>8), shift)
			candidate = int(table[currHash&tableMask])
			table[currHash&tableMask] = uint16(s)
			if uint32(x>>8) != load32(src, candidate) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:223
				_go_fuzz_dep_.CoverTab[90654]++
																	nextHash = hash(uint32(x>>16), shift)
																	s++
																	break
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:226
				// _ = "end of CoverTab[90654]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:227
				_go_fuzz_dep_.CoverTab[90655]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:227
				// _ = "end of CoverTab[90655]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:227
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:227
			// _ = "end of CoverTab[90649]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:228
		// _ = "end of CoverTab[90640]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:229
	// _ = "end of CoverTab[90634]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:229
	_go_fuzz_dep_.CoverTab[90635]++

emitRemainder:
	if nextEmit < len(src) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:232
		_go_fuzz_dep_.CoverTab[90656]++
															d += emitLiteral(dst[d:], src[nextEmit:])
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:233
		// _ = "end of CoverTab[90656]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:234
		_go_fuzz_dep_.CoverTab[90657]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:234
		// _ = "end of CoverTab[90657]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:234
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:234
	// _ = "end of CoverTab[90635]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:234
	_go_fuzz_dep_.CoverTab[90636]++
														return d
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:235
	// _ = "end of CoverTab[90636]"
}

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:236
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode_other.go:236
var _ = _go_fuzz_dep_.CoverTab
