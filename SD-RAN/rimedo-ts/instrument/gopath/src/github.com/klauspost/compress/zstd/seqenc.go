// Copyright 2019+ Klaus Post. All rights reserved.
// License information can be found in the LICENSE file.
// Based on work by Yann Collet, released under BSD License.

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqenc.go:5
package zstd

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqenc.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqenc.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqenc.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqenc.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqenc.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqenc.go:5
)

import "math/bits"

type seqCoders struct {
	llEnc, ofEnc, mlEnc	*fseEncoder
	llPrev, ofPrev, mlPrev	*fseEncoder
}

// swap coders with another (block).
func (s *seqCoders) swap(other *seqCoders) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqenc.go:15
	_go_fuzz_dep_.CoverTab[95042]++
												*s, *other = *other, *s
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqenc.go:16
	// _ = "end of CoverTab[95042]"
}

// setPrev will update the previous encoders to the actually used ones
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqenc.go:19
// and make sure a fresh one is in the main slot.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqenc.go:21
func (s *seqCoders) setPrev(ll, ml, of *fseEncoder) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqenc.go:21
	_go_fuzz_dep_.CoverTab[95043]++
												compareSwap := func(used *fseEncoder, current, prev **fseEncoder) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqenc.go:22
		_go_fuzz_dep_.CoverTab[95045]++

													if *current == used {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqenc.go:24
			_go_fuzz_dep_.CoverTab[95048]++
														*prev, *current = *current, *prev
														c := *current
														p := *prev
														c.reUsed = false
														p.reUsed = true
														return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqenc.go:30
			// _ = "end of CoverTab[95048]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqenc.go:31
			_go_fuzz_dep_.CoverTab[95049]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqenc.go:31
			// _ = "end of CoverTab[95049]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqenc.go:31
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqenc.go:31
		// _ = "end of CoverTab[95045]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqenc.go:31
		_go_fuzz_dep_.CoverTab[95046]++
													if used == *prev {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqenc.go:32
			_go_fuzz_dep_.CoverTab[95050]++
														return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqenc.go:33
			// _ = "end of CoverTab[95050]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqenc.go:34
			_go_fuzz_dep_.CoverTab[95051]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqenc.go:34
			// _ = "end of CoverTab[95051]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqenc.go:34
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqenc.go:34
		// _ = "end of CoverTab[95046]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqenc.go:34
		_go_fuzz_dep_.CoverTab[95047]++

													prevEnc := *prev
													prevEnc.symbolLen = 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqenc.go:37
		// _ = "end of CoverTab[95047]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqenc.go:38
	// _ = "end of CoverTab[95043]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqenc.go:38
	_go_fuzz_dep_.CoverTab[95044]++
												compareSwap(ll, &s.llEnc, &s.llPrev)
												compareSwap(ml, &s.mlEnc, &s.mlPrev)
												compareSwap(of, &s.ofEnc, &s.ofPrev)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqenc.go:41
	// _ = "end of CoverTab[95044]"
}

func highBit(val uint32) (n uint32) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqenc.go:44
	_go_fuzz_dep_.CoverTab[95052]++
												return uint32(bits.Len32(val) - 1)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqenc.go:45
	// _ = "end of CoverTab[95052]"
}

var llCodeTable = [64]byte{0, 1, 2, 3, 4, 5, 6, 7,
	8, 9, 10, 11, 12, 13, 14, 15,
	16, 16, 17, 17, 18, 18, 19, 19,
	20, 20, 20, 20, 21, 21, 21, 21,
	22, 22, 22, 22, 22, 22, 22, 22,
	23, 23, 23, 23, 23, 23, 23, 23,
	24, 24, 24, 24, 24, 24, 24, 24,
	24, 24, 24, 24, 24, 24, 24, 24}

// Up to 6 bits
const maxLLCode = 35

// llBitsTable translates from ll code to number of bits.
var llBitsTable = [maxLLCode + 1]byte{
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	1, 1, 1, 1, 2, 2, 3, 3,
	4, 6, 7, 8, 9, 10, 11, 12,
	13, 14, 15, 16}

// llCode returns the code that represents the literal length requested.
func llCode(litLength uint32) uint8 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqenc.go:69
	_go_fuzz_dep_.CoverTab[95053]++
												const llDeltaCode = 19
												if litLength <= 63 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqenc.go:71
		_go_fuzz_dep_.CoverTab[95055]++

													return llCodeTable[litLength&63]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqenc.go:73
		// _ = "end of CoverTab[95055]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqenc.go:74
		_go_fuzz_dep_.CoverTab[95056]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqenc.go:74
		// _ = "end of CoverTab[95056]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqenc.go:74
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqenc.go:74
	// _ = "end of CoverTab[95053]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqenc.go:74
	_go_fuzz_dep_.CoverTab[95054]++
												return uint8(highBit(litLength)) + llDeltaCode
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqenc.go:75
	// _ = "end of CoverTab[95054]"
}

var mlCodeTable = [128]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15,
	16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 32, 33, 33, 34, 34, 35, 35, 36, 36, 36, 36, 37, 37, 37, 37,
	38, 38, 38, 38, 38, 38, 38, 38, 39, 39, 39, 39, 39, 39, 39, 39,
	40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40,
	41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41,
	42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42,
	42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42}

// Up to 6 bits
const maxMLCode = 52

// mlBitsTable translates from ml code to number of bits.
var mlBitsTable = [maxMLCode + 1]byte{
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	1, 1, 1, 1, 2, 2, 3, 3,
	4, 4, 5, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16}

// note : mlBase = matchLength - MINMATCH;
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqenc.go:100
// because it's the format it's stored in seqStore->sequences
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqenc.go:102
func mlCode(mlBase uint32) uint8 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqenc.go:102
	_go_fuzz_dep_.CoverTab[95057]++
												const mlDeltaCode = 36
												if mlBase <= 127 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqenc.go:104
		_go_fuzz_dep_.CoverTab[95059]++

													return mlCodeTable[mlBase&127]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqenc.go:106
		// _ = "end of CoverTab[95059]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqenc.go:107
		_go_fuzz_dep_.CoverTab[95060]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqenc.go:107
		// _ = "end of CoverTab[95060]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqenc.go:107
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqenc.go:107
	// _ = "end of CoverTab[95057]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqenc.go:107
	_go_fuzz_dep_.CoverTab[95058]++
												return uint8(highBit(mlBase)) + mlDeltaCode
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqenc.go:108
	// _ = "end of CoverTab[95058]"
}

func ofCode(offset uint32) uint8 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqenc.go:111
	_go_fuzz_dep_.CoverTab[95061]++

												return uint8(bits.Len32(offset) - 1)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqenc.go:113
	// _ = "end of CoverTab[95061]"
}

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqenc.go:114
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqenc.go:114
var _ = _go_fuzz_dep_.CoverTab
