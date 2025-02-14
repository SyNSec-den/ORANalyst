// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/compress/flate/token.go:5
package flate

//line /usr/local/go/src/compress/flate/token.go:5
import (
//line /usr/local/go/src/compress/flate/token.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/compress/flate/token.go:5
)
//line /usr/local/go/src/compress/flate/token.go:5
import (
//line /usr/local/go/src/compress/flate/token.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/compress/flate/token.go:5
)

const (
	// 2 bits:   type   0 = literal  1=EOF  2=Match   3=Unused
	// 8 bits:   xlength = length - MIN_MATCH_LENGTH
	// 22 bits   xoffset = offset - MIN_OFFSET_SIZE, or literal
	lengthShift	= 22
	offsetMask	= 1<<lengthShift - 1
	typeMask	= 3 << 30
	literalType	= 0 << 30
	matchType	= 1 << 30
)

// The length code for length X (MIN_MATCH_LENGTH <= X <= MAX_MATCH_LENGTH)
//line /usr/local/go/src/compress/flate/token.go:18
// is lengthCodes[length - MIN_MATCH_LENGTH]
//line /usr/local/go/src/compress/flate/token.go:20
var lengthCodes = [...]uint32{
	0, 1, 2, 3, 4, 5, 6, 7, 8, 8,
	9, 9, 10, 10, 11, 11, 12, 12, 12, 12,
	13, 13, 13, 13, 14, 14, 14, 14, 15, 15,
	15, 15, 16, 16, 16, 16, 16, 16, 16, 16,
	17, 17, 17, 17, 17, 17, 17, 17, 18, 18,
	18, 18, 18, 18, 18, 18, 19, 19, 19, 19,
	19, 19, 19, 19, 20, 20, 20, 20, 20, 20,
	20, 20, 20, 20, 20, 20, 20, 20, 20, 20,
	21, 21, 21, 21, 21, 21, 21, 21, 21, 21,
	21, 21, 21, 21, 21, 21, 22, 22, 22, 22,
	22, 22, 22, 22, 22, 22, 22, 22, 22, 22,
	22, 22, 23, 23, 23, 23, 23, 23, 23, 23,
	23, 23, 23, 23, 23, 23, 23, 23, 24, 24,
	24, 24, 24, 24, 24, 24, 24, 24, 24, 24,
	24, 24, 24, 24, 24, 24, 24, 24, 24, 24,
	24, 24, 24, 24, 24, 24, 24, 24, 24, 24,
	25, 25, 25, 25, 25, 25, 25, 25, 25, 25,
	25, 25, 25, 25, 25, 25, 25, 25, 25, 25,
	25, 25, 25, 25, 25, 25, 25, 25, 25, 25,
	25, 25, 26, 26, 26, 26, 26, 26, 26, 26,
	26, 26, 26, 26, 26, 26, 26, 26, 26, 26,
	26, 26, 26, 26, 26, 26, 26, 26, 26, 26,
	26, 26, 26, 26, 27, 27, 27, 27, 27, 27,
	27, 27, 27, 27, 27, 27, 27, 27, 27, 27,
	27, 27, 27, 27, 27, 27, 27, 27, 27, 27,
	27, 27, 27, 27, 27, 28,
}

var offsetCodes = [...]uint32{
	0, 1, 2, 3, 4, 4, 5, 5, 6, 6, 6, 6, 7, 7, 7, 7,
	8, 8, 8, 8, 8, 8, 8, 8, 9, 9, 9, 9, 9, 9, 9, 9,
	10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12,
	12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12,
	13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13,
	13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13,
	14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14,
	14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14,
	14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14,
	14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
}

type token uint32

// Convert a literal into a literal token.
func literalToken(literal uint32) token {
//line /usr/local/go/src/compress/flate/token.go:71
	_go_fuzz_dep_.CoverTab[26526]++
//line /usr/local/go/src/compress/flate/token.go:71
	return token(literalType + literal)
//line /usr/local/go/src/compress/flate/token.go:71
	// _ = "end of CoverTab[26526]"
//line /usr/local/go/src/compress/flate/token.go:71
}

// Convert a < xlength, xoffset > pair into a match token.
func matchToken(xlength uint32, xoffset uint32) token {
//line /usr/local/go/src/compress/flate/token.go:74
	_go_fuzz_dep_.CoverTab[26527]++
							return token(matchType + xlength<<lengthShift + xoffset)
//line /usr/local/go/src/compress/flate/token.go:75
	// _ = "end of CoverTab[26527]"
}

// Returns the literal of a literal token.
func (t token) literal() uint32 {
//line /usr/local/go/src/compress/flate/token.go:79
	_go_fuzz_dep_.CoverTab[26528]++
//line /usr/local/go/src/compress/flate/token.go:79
	return uint32(t - literalType)
//line /usr/local/go/src/compress/flate/token.go:79
	// _ = "end of CoverTab[26528]"
//line /usr/local/go/src/compress/flate/token.go:79
}

// Returns the extra offset of a match token.
func (t token) offset() uint32 {
//line /usr/local/go/src/compress/flate/token.go:82
	_go_fuzz_dep_.CoverTab[26529]++
//line /usr/local/go/src/compress/flate/token.go:82
	return uint32(t) & offsetMask
//line /usr/local/go/src/compress/flate/token.go:82
	// _ = "end of CoverTab[26529]"
//line /usr/local/go/src/compress/flate/token.go:82
}

func (t token) length() uint32 {
//line /usr/local/go/src/compress/flate/token.go:84
	_go_fuzz_dep_.CoverTab[26530]++
//line /usr/local/go/src/compress/flate/token.go:84
	return uint32((t - matchType) >> lengthShift)
//line /usr/local/go/src/compress/flate/token.go:84
	// _ = "end of CoverTab[26530]"
//line /usr/local/go/src/compress/flate/token.go:84
}

func lengthCode(len uint32) uint32 {
//line /usr/local/go/src/compress/flate/token.go:86
	_go_fuzz_dep_.CoverTab[26531]++
//line /usr/local/go/src/compress/flate/token.go:86
	return lengthCodes[len]
//line /usr/local/go/src/compress/flate/token.go:86
	// _ = "end of CoverTab[26531]"
//line /usr/local/go/src/compress/flate/token.go:86
}

// Returns the offset code corresponding to a specific offset.
func offsetCode(off uint32) uint32 {
//line /usr/local/go/src/compress/flate/token.go:89
	_go_fuzz_dep_.CoverTab[26532]++
							if off < uint32(len(offsetCodes)) {
//line /usr/local/go/src/compress/flate/token.go:90
		_go_fuzz_dep_.CoverTab[26535]++
								return offsetCodes[off]
//line /usr/local/go/src/compress/flate/token.go:91
		// _ = "end of CoverTab[26535]"
	} else {
//line /usr/local/go/src/compress/flate/token.go:92
		_go_fuzz_dep_.CoverTab[26536]++
//line /usr/local/go/src/compress/flate/token.go:92
		// _ = "end of CoverTab[26536]"
//line /usr/local/go/src/compress/flate/token.go:92
	}
//line /usr/local/go/src/compress/flate/token.go:92
	// _ = "end of CoverTab[26532]"
//line /usr/local/go/src/compress/flate/token.go:92
	_go_fuzz_dep_.CoverTab[26533]++
							if off>>7 < uint32(len(offsetCodes)) {
//line /usr/local/go/src/compress/flate/token.go:93
		_go_fuzz_dep_.CoverTab[26537]++
								return offsetCodes[off>>7] + 14
//line /usr/local/go/src/compress/flate/token.go:94
		// _ = "end of CoverTab[26537]"
	} else {
//line /usr/local/go/src/compress/flate/token.go:95
		_go_fuzz_dep_.CoverTab[26538]++
//line /usr/local/go/src/compress/flate/token.go:95
		// _ = "end of CoverTab[26538]"
//line /usr/local/go/src/compress/flate/token.go:95
	}
//line /usr/local/go/src/compress/flate/token.go:95
	// _ = "end of CoverTab[26533]"
//line /usr/local/go/src/compress/flate/token.go:95
	_go_fuzz_dep_.CoverTab[26534]++
							return offsetCodes[off>>14] + 28
//line /usr/local/go/src/compress/flate/token.go:96
	// _ = "end of CoverTab[26534]"
}

//line /usr/local/go/src/compress/flate/token.go:97
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/compress/flate/token.go:97
var _ = _go_fuzz_dep_.CoverTab
