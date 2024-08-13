// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/compress/flate/deflatefast.go:5
package flate

//line /usr/local/go/src/compress/flate/deflatefast.go:5
import (
//line /usr/local/go/src/compress/flate/deflatefast.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/compress/flate/deflatefast.go:5
)
//line /usr/local/go/src/compress/flate/deflatefast.go:5
import (
//line /usr/local/go/src/compress/flate/deflatefast.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/compress/flate/deflatefast.go:5
)

import "math"

//line /usr/local/go/src/compress/flate/deflatefast.go:12
const (
	tableBits	= 14			// Bits used in the table.
	tableSize	= 1 << tableBits	// Size of the table.
	tableMask	= tableSize - 1		// Mask for table indices. Redundant, but can eliminate bounds checks.
	tableShift	= 32 - tableBits	// Right-shift to get the tableBits most significant bits of a uint32.

	// Reset the buffer offset when reaching this.
	// Offsets are stored between blocks as int32 values.
	// Since the offset we are checking against is at the beginning
	// of the buffer, we need to subtract the current and input
	// buffer to not risk overflowing the int32.
	bufferReset	= math.MaxInt32 - maxStoreBlockSize*2
)

func load32(b []byte, i int32) uint32 {
//line /usr/local/go/src/compress/flate/deflatefast.go:26
	_go_fuzz_dep_.CoverTab[25925]++
								b = b[i : i+4 : len(b)]
								return uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24
//line /usr/local/go/src/compress/flate/deflatefast.go:28
	// _ = "end of CoverTab[25925]"
}

func load64(b []byte, i int32) uint64 {
//line /usr/local/go/src/compress/flate/deflatefast.go:31
	_go_fuzz_dep_.CoverTab[25926]++
								b = b[i : i+8 : len(b)]
								return uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 |
		uint64(b[4])<<32 | uint64(b[5])<<40 | uint64(b[6])<<48 | uint64(b[7])<<56
//line /usr/local/go/src/compress/flate/deflatefast.go:34
	// _ = "end of CoverTab[25926]"
}

func hash(u uint32) uint32 {
//line /usr/local/go/src/compress/flate/deflatefast.go:37
	_go_fuzz_dep_.CoverTab[25927]++
								return (u * 0x1e35a7bd) >> tableShift
//line /usr/local/go/src/compress/flate/deflatefast.go:38
	// _ = "end of CoverTab[25927]"
}

// These constants are defined by the Snappy implementation so that its
//line /usr/local/go/src/compress/flate/deflatefast.go:41
// assembly implementation can fast-path some 16-bytes-at-a-time copies. They
//line /usr/local/go/src/compress/flate/deflatefast.go:41
// aren't necessary in the pure Go implementation, as we don't use those same
//line /usr/local/go/src/compress/flate/deflatefast.go:41
// optimizations, but using the same thresholds doesn't really hurt.
//line /usr/local/go/src/compress/flate/deflatefast.go:45
const (
	inputMargin		= 16 - 1
	minNonLiteralBlockSize	= 1 + 1 + inputMargin
)

type tableEntry struct {
	val	uint32	// Value at destination
	offset	int32
}

// deflateFast maintains the table for matches,
//line /usr/local/go/src/compress/flate/deflatefast.go:55
// and the previous byte block for cross block matching.
//line /usr/local/go/src/compress/flate/deflatefast.go:57
type deflateFast struct {
	table	[tableSize]tableEntry
	prev	[]byte	// Previous block, zero length if unknown.
	cur	int32	// Current match offset.
}

func newDeflateFast() *deflateFast {
//line /usr/local/go/src/compress/flate/deflatefast.go:63
	_go_fuzz_dep_.CoverTab[25928]++
								return &deflateFast{cur: maxStoreBlockSize, prev: make([]byte, 0, maxStoreBlockSize)}
//line /usr/local/go/src/compress/flate/deflatefast.go:64
	// _ = "end of CoverTab[25928]"
}

// encode encodes a block given in src and appends tokens
//line /usr/local/go/src/compress/flate/deflatefast.go:67
// to dst and returns the result.
//line /usr/local/go/src/compress/flate/deflatefast.go:69
func (e *deflateFast) encode(dst []token, src []byte) []token {
//line /usr/local/go/src/compress/flate/deflatefast.go:69
	_go_fuzz_dep_.CoverTab[25929]++

								if e.cur >= bufferReset {
//line /usr/local/go/src/compress/flate/deflatefast.go:71
		_go_fuzz_dep_.CoverTab[25934]++
									e.shiftOffsets()
//line /usr/local/go/src/compress/flate/deflatefast.go:72
		// _ = "end of CoverTab[25934]"
	} else {
//line /usr/local/go/src/compress/flate/deflatefast.go:73
		_go_fuzz_dep_.CoverTab[25935]++
//line /usr/local/go/src/compress/flate/deflatefast.go:73
		// _ = "end of CoverTab[25935]"
//line /usr/local/go/src/compress/flate/deflatefast.go:73
	}
//line /usr/local/go/src/compress/flate/deflatefast.go:73
	// _ = "end of CoverTab[25929]"
//line /usr/local/go/src/compress/flate/deflatefast.go:73
	_go_fuzz_dep_.CoverTab[25930]++

//line /usr/local/go/src/compress/flate/deflatefast.go:77
	if len(src) < minNonLiteralBlockSize {
//line /usr/local/go/src/compress/flate/deflatefast.go:77
		_go_fuzz_dep_.CoverTab[25936]++
									e.cur += maxStoreBlockSize
									e.prev = e.prev[:0]
									return emitLiteral(dst, src)
//line /usr/local/go/src/compress/flate/deflatefast.go:80
		// _ = "end of CoverTab[25936]"
	} else {
//line /usr/local/go/src/compress/flate/deflatefast.go:81
		_go_fuzz_dep_.CoverTab[25937]++
//line /usr/local/go/src/compress/flate/deflatefast.go:81
		// _ = "end of CoverTab[25937]"
//line /usr/local/go/src/compress/flate/deflatefast.go:81
	}
//line /usr/local/go/src/compress/flate/deflatefast.go:81
	// _ = "end of CoverTab[25930]"
//line /usr/local/go/src/compress/flate/deflatefast.go:81
	_go_fuzz_dep_.CoverTab[25931]++

//line /usr/local/go/src/compress/flate/deflatefast.go:86
	sLimit := int32(len(src) - inputMargin)

//line /usr/local/go/src/compress/flate/deflatefast.go:89
	nextEmit := int32(0)
	s := int32(0)
	cv := load32(src, s)
	nextHash := hash(cv)

	for {
//line /usr/local/go/src/compress/flate/deflatefast.go:94
		_go_fuzz_dep_.CoverTab[25938]++

//line /usr/local/go/src/compress/flate/deflatefast.go:110
		skip := int32(32)

		nextS := s
		var candidate tableEntry
		for {
//line /usr/local/go/src/compress/flate/deflatefast.go:114
			_go_fuzz_dep_.CoverTab[25940]++
										s = nextS
										bytesBetweenHashLookups := skip >> 5
										nextS = s + bytesBetweenHashLookups
										skip += bytesBetweenHashLookups
										if nextS > sLimit {
//line /usr/local/go/src/compress/flate/deflatefast.go:119
				_go_fuzz_dep_.CoverTab[25943]++
											goto emitRemainder
//line /usr/local/go/src/compress/flate/deflatefast.go:120
				// _ = "end of CoverTab[25943]"
			} else {
//line /usr/local/go/src/compress/flate/deflatefast.go:121
				_go_fuzz_dep_.CoverTab[25944]++
//line /usr/local/go/src/compress/flate/deflatefast.go:121
				// _ = "end of CoverTab[25944]"
//line /usr/local/go/src/compress/flate/deflatefast.go:121
			}
//line /usr/local/go/src/compress/flate/deflatefast.go:121
			// _ = "end of CoverTab[25940]"
//line /usr/local/go/src/compress/flate/deflatefast.go:121
			_go_fuzz_dep_.CoverTab[25941]++
										candidate = e.table[nextHash&tableMask]
										now := load32(src, nextS)
										e.table[nextHash&tableMask] = tableEntry{offset: s + e.cur, val: cv}
										nextHash = hash(now)

										offset := s - (candidate.offset - e.cur)
										if offset > maxMatchOffset || func() bool {
//line /usr/local/go/src/compress/flate/deflatefast.go:128
				_go_fuzz_dep_.CoverTab[25945]++
//line /usr/local/go/src/compress/flate/deflatefast.go:128
				return cv != candidate.val
//line /usr/local/go/src/compress/flate/deflatefast.go:128
				// _ = "end of CoverTab[25945]"
//line /usr/local/go/src/compress/flate/deflatefast.go:128
			}() {
//line /usr/local/go/src/compress/flate/deflatefast.go:128
				_go_fuzz_dep_.CoverTab[25946]++

											cv = now
											continue
//line /usr/local/go/src/compress/flate/deflatefast.go:131
				// _ = "end of CoverTab[25946]"
			} else {
//line /usr/local/go/src/compress/flate/deflatefast.go:132
				_go_fuzz_dep_.CoverTab[25947]++
//line /usr/local/go/src/compress/flate/deflatefast.go:132
				// _ = "end of CoverTab[25947]"
//line /usr/local/go/src/compress/flate/deflatefast.go:132
			}
//line /usr/local/go/src/compress/flate/deflatefast.go:132
			// _ = "end of CoverTab[25941]"
//line /usr/local/go/src/compress/flate/deflatefast.go:132
			_go_fuzz_dep_.CoverTab[25942]++
										break
//line /usr/local/go/src/compress/flate/deflatefast.go:133
			// _ = "end of CoverTab[25942]"
		}
//line /usr/local/go/src/compress/flate/deflatefast.go:134
		// _ = "end of CoverTab[25938]"
//line /usr/local/go/src/compress/flate/deflatefast.go:134
		_go_fuzz_dep_.CoverTab[25939]++

//line /usr/local/go/src/compress/flate/deflatefast.go:139
		dst = emitLiteral(dst, src[nextEmit:s])

//line /usr/local/go/src/compress/flate/deflatefast.go:149
		for {
//line /usr/local/go/src/compress/flate/deflatefast.go:149
			_go_fuzz_dep_.CoverTab[25948]++

//line /usr/local/go/src/compress/flate/deflatefast.go:155
			s += 4
										t := candidate.offset - e.cur + 4
										l := e.matchLen(s, t, src)

//line /usr/local/go/src/compress/flate/deflatefast.go:160
			dst = append(dst, matchToken(uint32(l+4-baseMatchLength), uint32(s-t-baseMatchOffset)))
			s += l
			nextEmit = s
			if s >= sLimit {
//line /usr/local/go/src/compress/flate/deflatefast.go:163
				_go_fuzz_dep_.CoverTab[25950]++
											goto emitRemainder
//line /usr/local/go/src/compress/flate/deflatefast.go:164
				// _ = "end of CoverTab[25950]"
			} else {
//line /usr/local/go/src/compress/flate/deflatefast.go:165
				_go_fuzz_dep_.CoverTab[25951]++
//line /usr/local/go/src/compress/flate/deflatefast.go:165
				// _ = "end of CoverTab[25951]"
//line /usr/local/go/src/compress/flate/deflatefast.go:165
			}
//line /usr/local/go/src/compress/flate/deflatefast.go:165
			// _ = "end of CoverTab[25948]"
//line /usr/local/go/src/compress/flate/deflatefast.go:165
			_go_fuzz_dep_.CoverTab[25949]++

//line /usr/local/go/src/compress/flate/deflatefast.go:173
			x := load64(src, s-1)
			prevHash := hash(uint32(x))
			e.table[prevHash&tableMask] = tableEntry{offset: e.cur + s - 1, val: uint32(x)}
			x >>= 8
			currHash := hash(uint32(x))
			candidate = e.table[currHash&tableMask]
			e.table[currHash&tableMask] = tableEntry{offset: e.cur + s, val: uint32(x)}

			offset := s - (candidate.offset - e.cur)
			if offset > maxMatchOffset || func() bool {
//line /usr/local/go/src/compress/flate/deflatefast.go:182
				_go_fuzz_dep_.CoverTab[25952]++
//line /usr/local/go/src/compress/flate/deflatefast.go:182
				return uint32(x) != candidate.val
//line /usr/local/go/src/compress/flate/deflatefast.go:182
				// _ = "end of CoverTab[25952]"
//line /usr/local/go/src/compress/flate/deflatefast.go:182
			}() {
//line /usr/local/go/src/compress/flate/deflatefast.go:182
				_go_fuzz_dep_.CoverTab[25953]++
											cv = uint32(x >> 8)
											nextHash = hash(cv)
											s++
											break
//line /usr/local/go/src/compress/flate/deflatefast.go:186
				// _ = "end of CoverTab[25953]"
			} else {
//line /usr/local/go/src/compress/flate/deflatefast.go:187
				_go_fuzz_dep_.CoverTab[25954]++
//line /usr/local/go/src/compress/flate/deflatefast.go:187
				// _ = "end of CoverTab[25954]"
//line /usr/local/go/src/compress/flate/deflatefast.go:187
			}
//line /usr/local/go/src/compress/flate/deflatefast.go:187
			// _ = "end of CoverTab[25949]"
		}
//line /usr/local/go/src/compress/flate/deflatefast.go:188
		// _ = "end of CoverTab[25939]"
	}
//line /usr/local/go/src/compress/flate/deflatefast.go:189
	// _ = "end of CoverTab[25931]"
//line /usr/local/go/src/compress/flate/deflatefast.go:189
	_go_fuzz_dep_.CoverTab[25932]++

emitRemainder:
	if int(nextEmit) < len(src) {
//line /usr/local/go/src/compress/flate/deflatefast.go:192
		_go_fuzz_dep_.CoverTab[25955]++
									dst = emitLiteral(dst, src[nextEmit:])
//line /usr/local/go/src/compress/flate/deflatefast.go:193
		// _ = "end of CoverTab[25955]"
	} else {
//line /usr/local/go/src/compress/flate/deflatefast.go:194
		_go_fuzz_dep_.CoverTab[25956]++
//line /usr/local/go/src/compress/flate/deflatefast.go:194
		// _ = "end of CoverTab[25956]"
//line /usr/local/go/src/compress/flate/deflatefast.go:194
	}
//line /usr/local/go/src/compress/flate/deflatefast.go:194
	// _ = "end of CoverTab[25932]"
//line /usr/local/go/src/compress/flate/deflatefast.go:194
	_go_fuzz_dep_.CoverTab[25933]++
								e.cur += int32(len(src))
								e.prev = e.prev[:len(src)]
								copy(e.prev, src)
								return dst
//line /usr/local/go/src/compress/flate/deflatefast.go:198
	// _ = "end of CoverTab[25933]"
}

func emitLiteral(dst []token, lit []byte) []token {
//line /usr/local/go/src/compress/flate/deflatefast.go:201
	_go_fuzz_dep_.CoverTab[25957]++
								for _, v := range lit {
//line /usr/local/go/src/compress/flate/deflatefast.go:202
		_go_fuzz_dep_.CoverTab[25959]++
									dst = append(dst, literalToken(uint32(v)))
//line /usr/local/go/src/compress/flate/deflatefast.go:203
		// _ = "end of CoverTab[25959]"
	}
//line /usr/local/go/src/compress/flate/deflatefast.go:204
	// _ = "end of CoverTab[25957]"
//line /usr/local/go/src/compress/flate/deflatefast.go:204
	_go_fuzz_dep_.CoverTab[25958]++
								return dst
//line /usr/local/go/src/compress/flate/deflatefast.go:205
	// _ = "end of CoverTab[25958]"
}

// matchLen returns the match length between src[s:] and src[t:].
//line /usr/local/go/src/compress/flate/deflatefast.go:208
// t can be negative to indicate the match is starting in e.prev.
//line /usr/local/go/src/compress/flate/deflatefast.go:208
// We assume that src[s-4:s] and src[t-4:t] already match.
//line /usr/local/go/src/compress/flate/deflatefast.go:211
func (e *deflateFast) matchLen(s, t int32, src []byte) int32 {
//line /usr/local/go/src/compress/flate/deflatefast.go:211
	_go_fuzz_dep_.CoverTab[25960]++
								s1 := int(s) + maxMatchLength - 4
								if s1 > len(src) {
//line /usr/local/go/src/compress/flate/deflatefast.go:213
		_go_fuzz_dep_.CoverTab[25968]++
									s1 = len(src)
//line /usr/local/go/src/compress/flate/deflatefast.go:214
		// _ = "end of CoverTab[25968]"
	} else {
//line /usr/local/go/src/compress/flate/deflatefast.go:215
		_go_fuzz_dep_.CoverTab[25969]++
//line /usr/local/go/src/compress/flate/deflatefast.go:215
		// _ = "end of CoverTab[25969]"
//line /usr/local/go/src/compress/flate/deflatefast.go:215
	}
//line /usr/local/go/src/compress/flate/deflatefast.go:215
	// _ = "end of CoverTab[25960]"
//line /usr/local/go/src/compress/flate/deflatefast.go:215
	_go_fuzz_dep_.CoverTab[25961]++

//line /usr/local/go/src/compress/flate/deflatefast.go:218
	if t >= 0 {
//line /usr/local/go/src/compress/flate/deflatefast.go:218
		_go_fuzz_dep_.CoverTab[25970]++
									b := src[t:]
									a := src[s:s1]
									b = b[:len(a)]

									for i := range a {
//line /usr/local/go/src/compress/flate/deflatefast.go:223
			_go_fuzz_dep_.CoverTab[25972]++
										if a[i] != b[i] {
//line /usr/local/go/src/compress/flate/deflatefast.go:224
				_go_fuzz_dep_.CoverTab[25973]++
											return int32(i)
//line /usr/local/go/src/compress/flate/deflatefast.go:225
				// _ = "end of CoverTab[25973]"
			} else {
//line /usr/local/go/src/compress/flate/deflatefast.go:226
				_go_fuzz_dep_.CoverTab[25974]++
//line /usr/local/go/src/compress/flate/deflatefast.go:226
				// _ = "end of CoverTab[25974]"
//line /usr/local/go/src/compress/flate/deflatefast.go:226
			}
//line /usr/local/go/src/compress/flate/deflatefast.go:226
			// _ = "end of CoverTab[25972]"
		}
//line /usr/local/go/src/compress/flate/deflatefast.go:227
		// _ = "end of CoverTab[25970]"
//line /usr/local/go/src/compress/flate/deflatefast.go:227
		_go_fuzz_dep_.CoverTab[25971]++
									return int32(len(a))
//line /usr/local/go/src/compress/flate/deflatefast.go:228
		// _ = "end of CoverTab[25971]"
	} else {
//line /usr/local/go/src/compress/flate/deflatefast.go:229
		_go_fuzz_dep_.CoverTab[25975]++
//line /usr/local/go/src/compress/flate/deflatefast.go:229
		// _ = "end of CoverTab[25975]"
//line /usr/local/go/src/compress/flate/deflatefast.go:229
	}
//line /usr/local/go/src/compress/flate/deflatefast.go:229
	// _ = "end of CoverTab[25961]"
//line /usr/local/go/src/compress/flate/deflatefast.go:229
	_go_fuzz_dep_.CoverTab[25962]++

//line /usr/local/go/src/compress/flate/deflatefast.go:232
	tp := int32(len(e.prev)) + t
	if tp < 0 {
//line /usr/local/go/src/compress/flate/deflatefast.go:233
		_go_fuzz_dep_.CoverTab[25976]++
									return 0
//line /usr/local/go/src/compress/flate/deflatefast.go:234
		// _ = "end of CoverTab[25976]"
	} else {
//line /usr/local/go/src/compress/flate/deflatefast.go:235
		_go_fuzz_dep_.CoverTab[25977]++
//line /usr/local/go/src/compress/flate/deflatefast.go:235
		// _ = "end of CoverTab[25977]"
//line /usr/local/go/src/compress/flate/deflatefast.go:235
	}
//line /usr/local/go/src/compress/flate/deflatefast.go:235
	// _ = "end of CoverTab[25962]"
//line /usr/local/go/src/compress/flate/deflatefast.go:235
	_go_fuzz_dep_.CoverTab[25963]++

//line /usr/local/go/src/compress/flate/deflatefast.go:238
	a := src[s:s1]
	b := e.prev[tp:]
	if len(b) > len(a) {
//line /usr/local/go/src/compress/flate/deflatefast.go:240
		_go_fuzz_dep_.CoverTab[25978]++
									b = b[:len(a)]
//line /usr/local/go/src/compress/flate/deflatefast.go:241
		// _ = "end of CoverTab[25978]"
	} else {
//line /usr/local/go/src/compress/flate/deflatefast.go:242
		_go_fuzz_dep_.CoverTab[25979]++
//line /usr/local/go/src/compress/flate/deflatefast.go:242
		// _ = "end of CoverTab[25979]"
//line /usr/local/go/src/compress/flate/deflatefast.go:242
	}
//line /usr/local/go/src/compress/flate/deflatefast.go:242
	// _ = "end of CoverTab[25963]"
//line /usr/local/go/src/compress/flate/deflatefast.go:242
	_go_fuzz_dep_.CoverTab[25964]++
								a = a[:len(b)]
								for i := range b {
//line /usr/local/go/src/compress/flate/deflatefast.go:244
		_go_fuzz_dep_.CoverTab[25980]++
									if a[i] != b[i] {
//line /usr/local/go/src/compress/flate/deflatefast.go:245
			_go_fuzz_dep_.CoverTab[25981]++
										return int32(i)
//line /usr/local/go/src/compress/flate/deflatefast.go:246
			// _ = "end of CoverTab[25981]"
		} else {
//line /usr/local/go/src/compress/flate/deflatefast.go:247
			_go_fuzz_dep_.CoverTab[25982]++
//line /usr/local/go/src/compress/flate/deflatefast.go:247
			// _ = "end of CoverTab[25982]"
//line /usr/local/go/src/compress/flate/deflatefast.go:247
		}
//line /usr/local/go/src/compress/flate/deflatefast.go:247
		// _ = "end of CoverTab[25980]"
	}
//line /usr/local/go/src/compress/flate/deflatefast.go:248
	// _ = "end of CoverTab[25964]"
//line /usr/local/go/src/compress/flate/deflatefast.go:248
	_go_fuzz_dep_.CoverTab[25965]++

//line /usr/local/go/src/compress/flate/deflatefast.go:252
	n := int32(len(b))
	if int(s+n) == s1 {
//line /usr/local/go/src/compress/flate/deflatefast.go:253
		_go_fuzz_dep_.CoverTab[25983]++
									return n
//line /usr/local/go/src/compress/flate/deflatefast.go:254
		// _ = "end of CoverTab[25983]"
	} else {
//line /usr/local/go/src/compress/flate/deflatefast.go:255
		_go_fuzz_dep_.CoverTab[25984]++
//line /usr/local/go/src/compress/flate/deflatefast.go:255
		// _ = "end of CoverTab[25984]"
//line /usr/local/go/src/compress/flate/deflatefast.go:255
	}
//line /usr/local/go/src/compress/flate/deflatefast.go:255
	// _ = "end of CoverTab[25965]"
//line /usr/local/go/src/compress/flate/deflatefast.go:255
	_go_fuzz_dep_.CoverTab[25966]++

//line /usr/local/go/src/compress/flate/deflatefast.go:258
	a = src[s+n : s1]
	b = src[:len(a)]
	for i := range a {
//line /usr/local/go/src/compress/flate/deflatefast.go:260
		_go_fuzz_dep_.CoverTab[25985]++
									if a[i] != b[i] {
//line /usr/local/go/src/compress/flate/deflatefast.go:261
			_go_fuzz_dep_.CoverTab[25986]++
										return int32(i) + n
//line /usr/local/go/src/compress/flate/deflatefast.go:262
			// _ = "end of CoverTab[25986]"
		} else {
//line /usr/local/go/src/compress/flate/deflatefast.go:263
			_go_fuzz_dep_.CoverTab[25987]++
//line /usr/local/go/src/compress/flate/deflatefast.go:263
			// _ = "end of CoverTab[25987]"
//line /usr/local/go/src/compress/flate/deflatefast.go:263
		}
//line /usr/local/go/src/compress/flate/deflatefast.go:263
		// _ = "end of CoverTab[25985]"
	}
//line /usr/local/go/src/compress/flate/deflatefast.go:264
	// _ = "end of CoverTab[25966]"
//line /usr/local/go/src/compress/flate/deflatefast.go:264
	_go_fuzz_dep_.CoverTab[25967]++
								return int32(len(a)) + n
//line /usr/local/go/src/compress/flate/deflatefast.go:265
	// _ = "end of CoverTab[25967]"
}

// Reset resets the encoding history.
//line /usr/local/go/src/compress/flate/deflatefast.go:268
// This ensures that no matches are made to the previous block.
//line /usr/local/go/src/compress/flate/deflatefast.go:270
func (e *deflateFast) reset() {
//line /usr/local/go/src/compress/flate/deflatefast.go:270
	_go_fuzz_dep_.CoverTab[25988]++
								e.prev = e.prev[:0]

//line /usr/local/go/src/compress/flate/deflatefast.go:274
	e.cur += maxMatchOffset

//line /usr/local/go/src/compress/flate/deflatefast.go:277
	if e.cur >= bufferReset {
//line /usr/local/go/src/compress/flate/deflatefast.go:277
		_go_fuzz_dep_.CoverTab[25989]++
									e.shiftOffsets()
//line /usr/local/go/src/compress/flate/deflatefast.go:278
		// _ = "end of CoverTab[25989]"
	} else {
//line /usr/local/go/src/compress/flate/deflatefast.go:279
		_go_fuzz_dep_.CoverTab[25990]++
//line /usr/local/go/src/compress/flate/deflatefast.go:279
		// _ = "end of CoverTab[25990]"
//line /usr/local/go/src/compress/flate/deflatefast.go:279
	}
//line /usr/local/go/src/compress/flate/deflatefast.go:279
	// _ = "end of CoverTab[25988]"
}

// shiftOffsets will shift down all match offset.
//line /usr/local/go/src/compress/flate/deflatefast.go:282
// This is only called in rare situations to prevent integer overflow.
//line /usr/local/go/src/compress/flate/deflatefast.go:282
//
//line /usr/local/go/src/compress/flate/deflatefast.go:282
// See https://golang.org/issue/18636 and https://github.com/golang/go/issues/34121.
//line /usr/local/go/src/compress/flate/deflatefast.go:286
func (e *deflateFast) shiftOffsets() {
//line /usr/local/go/src/compress/flate/deflatefast.go:286
	_go_fuzz_dep_.CoverTab[25991]++
								if len(e.prev) == 0 {
//line /usr/local/go/src/compress/flate/deflatefast.go:287
		_go_fuzz_dep_.CoverTab[25994]++

									for i := range e.table[:] {
//line /usr/local/go/src/compress/flate/deflatefast.go:289
			_go_fuzz_dep_.CoverTab[25996]++
										e.table[i] = tableEntry{}
//line /usr/local/go/src/compress/flate/deflatefast.go:290
			// _ = "end of CoverTab[25996]"
		}
//line /usr/local/go/src/compress/flate/deflatefast.go:291
		// _ = "end of CoverTab[25994]"
//line /usr/local/go/src/compress/flate/deflatefast.go:291
		_go_fuzz_dep_.CoverTab[25995]++
									e.cur = maxMatchOffset + 1
									return
//line /usr/local/go/src/compress/flate/deflatefast.go:293
		// _ = "end of CoverTab[25995]"
	} else {
//line /usr/local/go/src/compress/flate/deflatefast.go:294
		_go_fuzz_dep_.CoverTab[25997]++
//line /usr/local/go/src/compress/flate/deflatefast.go:294
		// _ = "end of CoverTab[25997]"
//line /usr/local/go/src/compress/flate/deflatefast.go:294
	}
//line /usr/local/go/src/compress/flate/deflatefast.go:294
	// _ = "end of CoverTab[25991]"
//line /usr/local/go/src/compress/flate/deflatefast.go:294
	_go_fuzz_dep_.CoverTab[25992]++

//line /usr/local/go/src/compress/flate/deflatefast.go:297
	for i := range e.table[:] {
//line /usr/local/go/src/compress/flate/deflatefast.go:297
		_go_fuzz_dep_.CoverTab[25998]++
									v := e.table[i].offset - e.cur + maxMatchOffset + 1
									if v < 0 {
//line /usr/local/go/src/compress/flate/deflatefast.go:299
			_go_fuzz_dep_.CoverTab[26000]++

//line /usr/local/go/src/compress/flate/deflatefast.go:304
			v = 0
//line /usr/local/go/src/compress/flate/deflatefast.go:304
			// _ = "end of CoverTab[26000]"
		} else {
//line /usr/local/go/src/compress/flate/deflatefast.go:305
			_go_fuzz_dep_.CoverTab[26001]++
//line /usr/local/go/src/compress/flate/deflatefast.go:305
			// _ = "end of CoverTab[26001]"
//line /usr/local/go/src/compress/flate/deflatefast.go:305
		}
//line /usr/local/go/src/compress/flate/deflatefast.go:305
		// _ = "end of CoverTab[25998]"
//line /usr/local/go/src/compress/flate/deflatefast.go:305
		_go_fuzz_dep_.CoverTab[25999]++
									e.table[i].offset = v
//line /usr/local/go/src/compress/flate/deflatefast.go:306
		// _ = "end of CoverTab[25999]"
	}
//line /usr/local/go/src/compress/flate/deflatefast.go:307
	// _ = "end of CoverTab[25992]"
//line /usr/local/go/src/compress/flate/deflatefast.go:307
	_go_fuzz_dep_.CoverTab[25993]++
								e.cur = maxMatchOffset + 1
//line /usr/local/go/src/compress/flate/deflatefast.go:308
	// _ = "end of CoverTab[25993]"
}

//line /usr/local/go/src/compress/flate/deflatefast.go:309
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/compress/flate/deflatefast.go:309
var _ = _go_fuzz_dep_.CoverTab
