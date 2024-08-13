// Copyright 2019+ Klaus Post. All rights reserved.
// License information can be found in the LICENSE file.
// Based on work by Yann Collet, released under BSD License.

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:5
package zstd

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:5
)

import (
	"bytes"
	"fmt"

	"github.com/klauspost/compress"
)

const (
	bestLongTableBits	= 22				// Bits used in the long match table
	bestLongTableSize	= 1 << bestLongTableBits	// Size of the table
	bestLongLen		= 8				// Bytes used for table hash

	// Note: Increasing the short table bits or making the hash shorter
	// can actually lead to compression degradation since it will 'steal' more from the
	// long match table and match offsets are quite big.
	// This greatly depends on the type of input.
	bestShortTableBits	= 18				// Bits used in the short match table
	bestShortTableSize	= 1 << bestShortTableBits	// Size of the table
	bestShortLen		= 4				// Bytes used for table hash

)

type match struct {
	offset	int32
	s	int32
	length	int32
	rep	int32
	est	int32
}

const highScore = 25000

// estBits will estimate output bits from predefined tables.
func (m *match) estBits(bitsPerByte int32) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:40
	_go_fuzz_dep_.CoverTab[92096]++
												mlc := mlCode(uint32(m.length - zstdMinMatch))
												var ofc uint8
												if m.rep < 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:43
		_go_fuzz_dep_.CoverTab[92098]++
													ofc = ofCode(uint32(m.s-m.offset) + 3)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:44
		// _ = "end of CoverTab[92098]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:45
		_go_fuzz_dep_.CoverTab[92099]++
													ofc = ofCode(uint32(m.rep))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:46
		// _ = "end of CoverTab[92099]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:47
	// _ = "end of CoverTab[92096]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:47
	_go_fuzz_dep_.CoverTab[92097]++

												ofTT, mlTT := fsePredefEnc[tableOffsets].ct.symbolTT[ofc], fsePredefEnc[tableMatchLengths].ct.symbolTT[mlc]

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:52
	m.est = int32(ofTT.outBits + mlTT.outBits)
	m.est += int32(ofTT.deltaNbBits>>16 + mlTT.deltaNbBits>>16)

	m.est -= (m.length * bitsPerByte) >> 10
	if m.est > 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:56
		_go_fuzz_dep_.CoverTab[92100]++

													m.length = 0
													m.est = highScore
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:59
		// _ = "end of CoverTab[92100]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:60
		_go_fuzz_dep_.CoverTab[92101]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:60
		// _ = "end of CoverTab[92101]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:60
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:60
	// _ = "end of CoverTab[92097]"
}

// bestFastEncoder uses 2 tables, one for short matches (5 bytes) and one for long matches.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:63
// The long match table contains the previous entry with the same hash,
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:63
// effectively making it a "chain" of length 2.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:63
// When we find a long match we choose between the two values and select the longest.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:63
// When we find a short match, after checking the long, we check if we can find a long at n+1
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:63
// and that it is longer (lazy matching).
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:69
type bestFastEncoder struct {
	fastBase
	table		[bestShortTableSize]prevEntry
	longTable	[bestLongTableSize]prevEntry
	dictTable	[]prevEntry
	dictLongTable	[]prevEntry
}

// Encode improves compression...
func (e *bestFastEncoder) Encode(blk *blockEnc, src []byte) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:78
	_go_fuzz_dep_.CoverTab[92102]++
												const (
		// Input margin is the number of bytes we read (8)
		// and the maximum we will read ahead (2)
		inputMargin		= 8 + 4
		minNonLiteralBlockSize	= 16
	)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:87
	for e.cur >= bufferReset {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:87
		_go_fuzz_dep_.CoverTab[92110]++
													if len(e.hist) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:88
			_go_fuzz_dep_.CoverTab[92114]++
														for i := range e.table[:] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:89
				_go_fuzz_dep_.CoverTab[92117]++
															e.table[i] = prevEntry{}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:90
				// _ = "end of CoverTab[92117]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:91
			// _ = "end of CoverTab[92114]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:91
			_go_fuzz_dep_.CoverTab[92115]++
														for i := range e.longTable[:] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:92
				_go_fuzz_dep_.CoverTab[92118]++
															e.longTable[i] = prevEntry{}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:93
				// _ = "end of CoverTab[92118]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:94
			// _ = "end of CoverTab[92115]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:94
			_go_fuzz_dep_.CoverTab[92116]++
														e.cur = e.maxMatchOff
														break
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:96
			// _ = "end of CoverTab[92116]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:97
			_go_fuzz_dep_.CoverTab[92119]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:97
			// _ = "end of CoverTab[92119]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:97
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:97
		// _ = "end of CoverTab[92110]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:97
		_go_fuzz_dep_.CoverTab[92111]++

													minOff := e.cur + int32(len(e.hist)) - e.maxMatchOff
													for i := range e.table[:] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:100
			_go_fuzz_dep_.CoverTab[92120]++
														v := e.table[i].offset
														v2 := e.table[i].prev
														if v < minOff {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:103
				_go_fuzz_dep_.CoverTab[92122]++
															v = 0
															v2 = 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:105
				// _ = "end of CoverTab[92122]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:106
				_go_fuzz_dep_.CoverTab[92123]++
															v = v - e.cur + e.maxMatchOff
															if v2 < minOff {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:108
					_go_fuzz_dep_.CoverTab[92124]++
																v2 = 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:109
					// _ = "end of CoverTab[92124]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:110
					_go_fuzz_dep_.CoverTab[92125]++
																v2 = v2 - e.cur + e.maxMatchOff
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:111
					// _ = "end of CoverTab[92125]"
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:112
				// _ = "end of CoverTab[92123]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:113
			// _ = "end of CoverTab[92120]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:113
			_go_fuzz_dep_.CoverTab[92121]++
														e.table[i] = prevEntry{
				offset:	v,
				prev:	v2,
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:117
			// _ = "end of CoverTab[92121]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:118
		// _ = "end of CoverTab[92111]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:118
		_go_fuzz_dep_.CoverTab[92112]++
													for i := range e.longTable[:] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:119
			_go_fuzz_dep_.CoverTab[92126]++
														v := e.longTable[i].offset
														v2 := e.longTable[i].prev
														if v < minOff {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:122
				_go_fuzz_dep_.CoverTab[92128]++
															v = 0
															v2 = 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:124
				// _ = "end of CoverTab[92128]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:125
				_go_fuzz_dep_.CoverTab[92129]++
															v = v - e.cur + e.maxMatchOff
															if v2 < minOff {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:127
					_go_fuzz_dep_.CoverTab[92130]++
																v2 = 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:128
					// _ = "end of CoverTab[92130]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:129
					_go_fuzz_dep_.CoverTab[92131]++
																v2 = v2 - e.cur + e.maxMatchOff
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:130
					// _ = "end of CoverTab[92131]"
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:131
				// _ = "end of CoverTab[92129]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:132
			// _ = "end of CoverTab[92126]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:132
			_go_fuzz_dep_.CoverTab[92127]++
														e.longTable[i] = prevEntry{
				offset:	v,
				prev:	v2,
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:136
			// _ = "end of CoverTab[92127]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:137
		// _ = "end of CoverTab[92112]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:137
		_go_fuzz_dep_.CoverTab[92113]++
													e.cur = e.maxMatchOff
													break
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:139
		// _ = "end of CoverTab[92113]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:140
	// _ = "end of CoverTab[92102]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:140
	_go_fuzz_dep_.CoverTab[92103]++

												s := e.addBlock(src)
												blk.size = len(src)
												if len(src) < minNonLiteralBlockSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:144
		_go_fuzz_dep_.CoverTab[92132]++
													blk.extraLits = len(src)
													blk.literals = blk.literals[:len(src)]
													copy(blk.literals, src)
													return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:148
		// _ = "end of CoverTab[92132]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:149
		_go_fuzz_dep_.CoverTab[92133]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:149
		// _ = "end of CoverTab[92133]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:149
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:149
	// _ = "end of CoverTab[92103]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:149
	_go_fuzz_dep_.CoverTab[92104]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:153
	bitsPerByte := int32((compress.ShannonEntropyBits(src) * 1024) / len(src))

	if bitsPerByte < 1024 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:155
		_go_fuzz_dep_.CoverTab[92134]++
													bitsPerByte = 1024
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:156
		// _ = "end of CoverTab[92134]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:157
		_go_fuzz_dep_.CoverTab[92135]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:157
		// _ = "end of CoverTab[92135]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:157
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:157
	// _ = "end of CoverTab[92104]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:157
	_go_fuzz_dep_.CoverTab[92105]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:160
	src = e.hist
												sLimit := int32(len(src)) - inputMargin
												const kSearchStrength = 10

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:165
	nextEmit := s
												cv := load6432(src, s)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:169
	offset1 := int32(blk.recentOffsets[0])
	offset2 := int32(blk.recentOffsets[1])
	offset3 := int32(blk.recentOffsets[2])

	addLiterals := func(s *seq, until int32) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:173
		_go_fuzz_dep_.CoverTab[92136]++
													if until == nextEmit {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:174
			_go_fuzz_dep_.CoverTab[92138]++
														return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:175
			// _ = "end of CoverTab[92138]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:176
			_go_fuzz_dep_.CoverTab[92139]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:176
			// _ = "end of CoverTab[92139]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:176
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:176
		// _ = "end of CoverTab[92136]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:176
		_go_fuzz_dep_.CoverTab[92137]++
													blk.literals = append(blk.literals, src[nextEmit:until]...)
													s.litLen = uint32(until - nextEmit)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:178
		// _ = "end of CoverTab[92137]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:179
	// _ = "end of CoverTab[92105]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:179
	_go_fuzz_dep_.CoverTab[92106]++
												_ = addLiterals

												if debugEncoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:182
		_go_fuzz_dep_.CoverTab[92140]++
													println("recent offsets:", blk.recentOffsets)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:183
		// _ = "end of CoverTab[92140]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:184
		_go_fuzz_dep_.CoverTab[92141]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:184
		// _ = "end of CoverTab[92141]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:184
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:184
	// _ = "end of CoverTab[92106]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:184
	_go_fuzz_dep_.CoverTab[92107]++

encodeLoop:
	for {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:187
		_go_fuzz_dep_.CoverTab[92142]++

													canRepeat := len(blk.sequences) > 2

													if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:191
			_go_fuzz_dep_.CoverTab[92159]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:191
			return canRepeat
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:191
			// _ = "end of CoverTab[92159]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:191
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:191
			_go_fuzz_dep_.CoverTab[92160]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:191
			return offset1 == 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:191
			// _ = "end of CoverTab[92160]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:191
		}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:191
			_go_fuzz_dep_.CoverTab[92161]++
														panic("offset0 was 0")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:192
			// _ = "end of CoverTab[92161]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:193
			_go_fuzz_dep_.CoverTab[92162]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:193
			// _ = "end of CoverTab[92162]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:193
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:193
		// _ = "end of CoverTab[92142]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:193
		_go_fuzz_dep_.CoverTab[92143]++

													bestOf := func(a, b match) match {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:195
			_go_fuzz_dep_.CoverTab[92163]++
														if a.est+(a.s-b.s)*bitsPerByte>>10 < b.est+(b.s-a.s)*bitsPerByte>>10 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:196
				_go_fuzz_dep_.CoverTab[92165]++
															return a
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:197
				// _ = "end of CoverTab[92165]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:198
				_go_fuzz_dep_.CoverTab[92166]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:198
				// _ = "end of CoverTab[92166]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:198
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:198
			// _ = "end of CoverTab[92163]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:198
			_go_fuzz_dep_.CoverTab[92164]++
														return b
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:199
			// _ = "end of CoverTab[92164]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:200
		// _ = "end of CoverTab[92143]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:200
		_go_fuzz_dep_.CoverTab[92144]++
													const goodEnough = 100

													nextHashL := hashLen(cv, bestLongTableBits, bestLongLen)
													nextHashS := hashLen(cv, bestShortTableBits, bestShortLen)
													candidateL := e.longTable[nextHashL]
													candidateS := e.table[nextHashS]

													matchAt := func(offset int32, s int32, first uint32, rep int32) match {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:208
			_go_fuzz_dep_.CoverTab[92167]++
														if s-offset >= e.maxMatchOff || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:209
				_go_fuzz_dep_.CoverTab[92170]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:209
				return load3232(src, offset) != first
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:209
				// _ = "end of CoverTab[92170]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:209
			}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:209
				_go_fuzz_dep_.CoverTab[92171]++
															return match{s: s, est: highScore}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:210
				// _ = "end of CoverTab[92171]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:211
				_go_fuzz_dep_.CoverTab[92172]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:211
				// _ = "end of CoverTab[92172]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:211
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:211
			// _ = "end of CoverTab[92167]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:211
			_go_fuzz_dep_.CoverTab[92168]++
														if debugAsserts {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:212
				_go_fuzz_dep_.CoverTab[92173]++
															if !bytes.Equal(src[s:s+4], src[offset:offset+4]) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:213
					_go_fuzz_dep_.CoverTab[92174]++
																panic(fmt.Sprintf("first match mismatch: %v != %v, first: %08x", src[s:s+4], src[offset:offset+4], first))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:214
					// _ = "end of CoverTab[92174]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:215
					_go_fuzz_dep_.CoverTab[92175]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:215
					// _ = "end of CoverTab[92175]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:215
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:215
				// _ = "end of CoverTab[92173]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:216
				_go_fuzz_dep_.CoverTab[92176]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:216
				// _ = "end of CoverTab[92176]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:216
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:216
			// _ = "end of CoverTab[92168]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:216
			_go_fuzz_dep_.CoverTab[92169]++
														m := match{offset: offset, s: s, length: 4 + e.matchlen(s+4, offset+4, src), rep: rep}
														m.estBits(bitsPerByte)
														return m
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:219
			// _ = "end of CoverTab[92169]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:220
		// _ = "end of CoverTab[92144]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:220
		_go_fuzz_dep_.CoverTab[92145]++

													best := bestOf(matchAt(candidateL.offset-e.cur, s, uint32(cv), -1), matchAt(candidateL.prev-e.cur, s, uint32(cv), -1))
													best = bestOf(best, matchAt(candidateS.offset-e.cur, s, uint32(cv), -1))
													best = bestOf(best, matchAt(candidateS.prev-e.cur, s, uint32(cv), -1))

													if canRepeat && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:226
			_go_fuzz_dep_.CoverTab[92177]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:226
			return best.length < goodEnough
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:226
			// _ = "end of CoverTab[92177]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:226
		}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:226
			_go_fuzz_dep_.CoverTab[92178]++
														cv32 := uint32(cv >> 8)
														spp := s + 1
														best = bestOf(best, matchAt(spp-offset1, spp, cv32, 1))
														best = bestOf(best, matchAt(spp-offset2, spp, cv32, 2))
														best = bestOf(best, matchAt(spp-offset3, spp, cv32, 3))
														if best.length > 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:232
				_go_fuzz_dep_.CoverTab[92179]++
															cv32 = uint32(cv >> 24)
															spp += 2
															best = bestOf(best, matchAt(spp-offset1, spp, cv32, 1))
															best = bestOf(best, matchAt(spp-offset2, spp, cv32, 2))
															best = bestOf(best, matchAt(spp-offset3, spp, cv32, 3))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:237
				// _ = "end of CoverTab[92179]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:238
				_go_fuzz_dep_.CoverTab[92180]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:238
				// _ = "end of CoverTab[92180]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:238
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:238
			// _ = "end of CoverTab[92178]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:239
			_go_fuzz_dep_.CoverTab[92181]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:239
			// _ = "end of CoverTab[92181]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:239
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:239
		// _ = "end of CoverTab[92145]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:239
		_go_fuzz_dep_.CoverTab[92146]++

													e.longTable[nextHashL] = prevEntry{offset: s + e.cur, prev: candidateL.offset}
													e.table[nextHashS] = prevEntry{offset: s + e.cur, prev: candidateS.offset}

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:245
		if best.length < goodEnough {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:245
			_go_fuzz_dep_.CoverTab[92182]++

														if best.length < 4 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:247
				_go_fuzz_dep_.CoverTab[92185]++
															s += 1 + (s-nextEmit)>>(kSearchStrength-1)
															if s >= sLimit {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:249
					_go_fuzz_dep_.CoverTab[92187]++
																break encodeLoop
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:250
					// _ = "end of CoverTab[92187]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:251
					_go_fuzz_dep_.CoverTab[92188]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:251
					// _ = "end of CoverTab[92188]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:251
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:251
				// _ = "end of CoverTab[92185]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:251
				_go_fuzz_dep_.CoverTab[92186]++
															cv = load6432(src, s)
															continue
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:253
				// _ = "end of CoverTab[92186]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:254
				_go_fuzz_dep_.CoverTab[92189]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:254
				// _ = "end of CoverTab[92189]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:254
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:254
			// _ = "end of CoverTab[92182]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:254
			_go_fuzz_dep_.CoverTab[92183]++

														s++
														candidateS = e.table[hashLen(cv>>8, bestShortTableBits, bestShortLen)]
														cv = load6432(src, s)
														cv2 := load6432(src, s+1)
														candidateL = e.longTable[hashLen(cv, bestLongTableBits, bestLongLen)]
														candidateL2 := e.longTable[hashLen(cv2, bestLongTableBits, bestLongLen)]

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:264
			best = bestOf(best, matchAt(candidateS.offset-e.cur, s, uint32(cv), -1))

			best = bestOf(best, matchAt(candidateL.offset-e.cur, s, uint32(cv), -1))
			best = bestOf(best, matchAt(candidateL.prev-e.cur, s, uint32(cv), -1))
			best = bestOf(best, matchAt(candidateL2.offset-e.cur, s+1, uint32(cv2), -1))
			best = bestOf(best, matchAt(candidateL2.prev-e.cur, s+1, uint32(cv2), -1))
			if false {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:270
				_go_fuzz_dep_.CoverTab[92190]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:273
				best = bestOf(best, matchAt(e.table[hashLen(cv2>>8, bestShortTableBits, bestShortLen)].offset-e.cur, s+2, uint32(cv2>>8), -1))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:273
				// _ = "end of CoverTab[92190]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:274
				_go_fuzz_dep_.CoverTab[92191]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:274
				// _ = "end of CoverTab[92191]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:274
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:274
			// _ = "end of CoverTab[92183]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:274
			_go_fuzz_dep_.CoverTab[92184]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:277
			if sAt := best.s + best.length; sAt < sLimit {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:277
				_go_fuzz_dep_.CoverTab[92192]++
															nextHashL := hashLen(load6432(src, sAt), bestLongTableBits, bestLongLen)
															candidateEnd := e.longTable[nextHashL]
															if pos := candidateEnd.offset - e.cur - best.length; pos >= 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:280
					_go_fuzz_dep_.CoverTab[92193]++
																bestEnd := bestOf(best, matchAt(pos, best.s, load3232(src, best.s), -1))
																if pos := candidateEnd.prev - e.cur - best.length; pos >= 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:282
						_go_fuzz_dep_.CoverTab[92195]++
																	bestEnd = bestOf(bestEnd, matchAt(pos, best.s, load3232(src, best.s), -1))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:283
						// _ = "end of CoverTab[92195]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:284
						_go_fuzz_dep_.CoverTab[92196]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:284
						// _ = "end of CoverTab[92196]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:284
					}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:284
					// _ = "end of CoverTab[92193]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:284
					_go_fuzz_dep_.CoverTab[92194]++
																best = bestEnd
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:285
					// _ = "end of CoverTab[92194]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:286
					_go_fuzz_dep_.CoverTab[92197]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:286
					// _ = "end of CoverTab[92197]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:286
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:286
				// _ = "end of CoverTab[92192]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:287
				_go_fuzz_dep_.CoverTab[92198]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:287
				// _ = "end of CoverTab[92198]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:287
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:287
			// _ = "end of CoverTab[92184]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:288
			_go_fuzz_dep_.CoverTab[92199]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:288
			// _ = "end of CoverTab[92199]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:288
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:288
		// _ = "end of CoverTab[92146]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:288
		_go_fuzz_dep_.CoverTab[92147]++

													if debugAsserts {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:290
			_go_fuzz_dep_.CoverTab[92200]++
														if !bytes.Equal(src[best.s:best.s+best.length], src[best.offset:best.offset+best.length]) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:291
				_go_fuzz_dep_.CoverTab[92201]++
															panic(fmt.Sprintf("match mismatch: %v != %v", src[best.s:best.s+best.length], src[best.offset:best.offset+best.length]))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:292
				// _ = "end of CoverTab[92201]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:293
				_go_fuzz_dep_.CoverTab[92202]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:293
				// _ = "end of CoverTab[92202]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:293
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:293
			// _ = "end of CoverTab[92200]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:294
			_go_fuzz_dep_.CoverTab[92203]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:294
			// _ = "end of CoverTab[92203]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:294
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:294
		// _ = "end of CoverTab[92147]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:294
		_go_fuzz_dep_.CoverTab[92148]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:297
		if best.rep > 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:297
			_go_fuzz_dep_.CoverTab[92204]++
														s = best.s
														var seq seq
														seq.matchLen = uint32(best.length - zstdMinMatch)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:304
			start := best.s

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:307
			startLimit := nextEmit + 1

			tMin := s - e.maxMatchOff
			if tMin < 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:310
				_go_fuzz_dep_.CoverTab[92211]++
															tMin = 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:311
				// _ = "end of CoverTab[92211]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:312
				_go_fuzz_dep_.CoverTab[92212]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:312
				// _ = "end of CoverTab[92212]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:312
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:312
			// _ = "end of CoverTab[92204]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:312
			_go_fuzz_dep_.CoverTab[92205]++
														repIndex := best.offset
														for repIndex > tMin && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:314
				_go_fuzz_dep_.CoverTab[92213]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:314
				return start > startLimit
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:314
				// _ = "end of CoverTab[92213]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:314
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:314
				_go_fuzz_dep_.CoverTab[92214]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:314
				return src[repIndex-1] == src[start-1]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:314
				// _ = "end of CoverTab[92214]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:314
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:314
				_go_fuzz_dep_.CoverTab[92215]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:314
				return seq.matchLen < maxMatchLength-zstdMinMatch-1
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:314
				// _ = "end of CoverTab[92215]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:314
			}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:314
				_go_fuzz_dep_.CoverTab[92216]++
															repIndex--
															start--
															seq.matchLen++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:317
				// _ = "end of CoverTab[92216]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:318
			// _ = "end of CoverTab[92205]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:318
			_go_fuzz_dep_.CoverTab[92206]++
														addLiterals(&seq, start)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:322
			seq.offset = uint32(best.rep)
			if debugSequences {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:323
				_go_fuzz_dep_.CoverTab[92217]++
															println("repeat sequence", seq, "next s:", s)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:324
				// _ = "end of CoverTab[92217]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:325
				_go_fuzz_dep_.CoverTab[92218]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:325
				// _ = "end of CoverTab[92218]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:325
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:325
			// _ = "end of CoverTab[92206]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:325
			_go_fuzz_dep_.CoverTab[92207]++
														blk.sequences = append(blk.sequences, seq)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:329
			index0 := s
			s = best.s + best.length

			nextEmit = s
			if s >= sLimit {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:333
				_go_fuzz_dep_.CoverTab[92219]++
															if debugEncoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:334
					_go_fuzz_dep_.CoverTab[92221]++
																println("repeat ended", s, best.length)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:335
					// _ = "end of CoverTab[92221]"

				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:337
					_go_fuzz_dep_.CoverTab[92222]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:337
					// _ = "end of CoverTab[92222]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:337
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:337
				// _ = "end of CoverTab[92219]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:337
				_go_fuzz_dep_.CoverTab[92220]++
															break encodeLoop
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:338
				// _ = "end of CoverTab[92220]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:339
				_go_fuzz_dep_.CoverTab[92223]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:339
				// _ = "end of CoverTab[92223]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:339
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:339
			// _ = "end of CoverTab[92207]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:339
			_go_fuzz_dep_.CoverTab[92208]++

														off := index0 + e.cur
														for index0 < s-1 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:342
				_go_fuzz_dep_.CoverTab[92224]++
															cv0 := load6432(src, index0)
															h0 := hashLen(cv0, bestLongTableBits, bestLongLen)
															h1 := hashLen(cv0, bestShortTableBits, bestShortLen)
															e.longTable[h0] = prevEntry{offset: off, prev: e.longTable[h0].offset}
															e.table[h1] = prevEntry{offset: off, prev: e.table[h1].offset}
															off++
															index0++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:349
				// _ = "end of CoverTab[92224]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:350
			// _ = "end of CoverTab[92208]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:350
			_go_fuzz_dep_.CoverTab[92209]++
														switch best.rep {
			case 2:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:352
				_go_fuzz_dep_.CoverTab[92225]++
															offset1, offset2 = offset2, offset1
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:353
				// _ = "end of CoverTab[92225]"
			case 3:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:354
				_go_fuzz_dep_.CoverTab[92226]++
															offset1, offset2, offset3 = offset3, offset1, offset2
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:355
				// _ = "end of CoverTab[92226]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:355
			default:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:355
				_go_fuzz_dep_.CoverTab[92227]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:355
				// _ = "end of CoverTab[92227]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:356
			// _ = "end of CoverTab[92209]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:356
			_go_fuzz_dep_.CoverTab[92210]++
														cv = load6432(src, s)
														continue
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:358
			// _ = "end of CoverTab[92210]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:359
			_go_fuzz_dep_.CoverTab[92228]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:359
			// _ = "end of CoverTab[92228]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:359
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:359
		// _ = "end of CoverTab[92148]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:359
		_go_fuzz_dep_.CoverTab[92149]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:363
		s = best.s
		t := best.offset
		offset1, offset2, offset3 = s-t, offset1, offset2

		if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:367
			_go_fuzz_dep_.CoverTab[92229]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:367
			return s <= t
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:367
			// _ = "end of CoverTab[92229]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:367
		}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:367
			_go_fuzz_dep_.CoverTab[92230]++
														panic(fmt.Sprintf("s (%d) <= t (%d)", s, t))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:368
			// _ = "end of CoverTab[92230]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:369
			_go_fuzz_dep_.CoverTab[92231]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:369
			// _ = "end of CoverTab[92231]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:369
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:369
		// _ = "end of CoverTab[92149]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:369
		_go_fuzz_dep_.CoverTab[92150]++

													if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:371
			_go_fuzz_dep_.CoverTab[92232]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:371
			return int(offset1) > len(src)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:371
			// _ = "end of CoverTab[92232]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:371
		}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:371
			_go_fuzz_dep_.CoverTab[92233]++
														panic("invalid offset")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:372
			// _ = "end of CoverTab[92233]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:373
			_go_fuzz_dep_.CoverTab[92234]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:373
			// _ = "end of CoverTab[92234]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:373
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:373
		// _ = "end of CoverTab[92150]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:373
		_go_fuzz_dep_.CoverTab[92151]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:376
		l := best.length

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:379
		tMin := s - e.maxMatchOff
		if tMin < 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:380
			_go_fuzz_dep_.CoverTab[92235]++
														tMin = 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:381
			// _ = "end of CoverTab[92235]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:382
			_go_fuzz_dep_.CoverTab[92236]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:382
			// _ = "end of CoverTab[92236]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:382
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:382
		// _ = "end of CoverTab[92151]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:382
		_go_fuzz_dep_.CoverTab[92152]++
													for t > tMin && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:383
			_go_fuzz_dep_.CoverTab[92237]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:383
			return s > nextEmit
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:383
			// _ = "end of CoverTab[92237]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:383
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:383
			_go_fuzz_dep_.CoverTab[92238]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:383
			return src[t-1] == src[s-1]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:383
			// _ = "end of CoverTab[92238]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:383
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:383
			_go_fuzz_dep_.CoverTab[92239]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:383
			return l < maxMatchLength
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:383
			// _ = "end of CoverTab[92239]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:383
		}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:383
			_go_fuzz_dep_.CoverTab[92240]++
														s--
														t--
														l++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:386
			// _ = "end of CoverTab[92240]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:387
		// _ = "end of CoverTab[92152]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:387
		_go_fuzz_dep_.CoverTab[92153]++

		// Write our sequence
		var seq seq
		seq.litLen = uint32(s - nextEmit)
		seq.matchLen = uint32(l - zstdMinMatch)
		if seq.litLen > 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:393
			_go_fuzz_dep_.CoverTab[92241]++
														blk.literals = append(blk.literals, src[nextEmit:s]...)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:394
			// _ = "end of CoverTab[92241]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:395
			_go_fuzz_dep_.CoverTab[92242]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:395
			// _ = "end of CoverTab[92242]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:395
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:395
		// _ = "end of CoverTab[92153]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:395
		_go_fuzz_dep_.CoverTab[92154]++
													seq.offset = uint32(s-t) + 3
													s += l
													if debugSequences {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:398
			_go_fuzz_dep_.CoverTab[92243]++
														println("sequence", seq, "next s:", s)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:399
			// _ = "end of CoverTab[92243]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:400
			_go_fuzz_dep_.CoverTab[92244]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:400
			// _ = "end of CoverTab[92244]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:400
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:400
		// _ = "end of CoverTab[92154]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:400
		_go_fuzz_dep_.CoverTab[92155]++
													blk.sequences = append(blk.sequences, seq)
													nextEmit = s
													if s >= sLimit {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:403
			_go_fuzz_dep_.CoverTab[92245]++
														break encodeLoop
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:404
			// _ = "end of CoverTab[92245]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:405
			_go_fuzz_dep_.CoverTab[92246]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:405
			// _ = "end of CoverTab[92246]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:405
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:405
		// _ = "end of CoverTab[92155]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:405
		_go_fuzz_dep_.CoverTab[92156]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:408
		index0 := s - l + 1

		for index0 < s-1 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:410
			_go_fuzz_dep_.CoverTab[92247]++
														cv0 := load6432(src, index0)
														h0 := hashLen(cv0, bestLongTableBits, bestLongLen)
														h1 := hashLen(cv0, bestShortTableBits, bestShortLen)
														off := index0 + e.cur
														e.longTable[h0] = prevEntry{offset: off, prev: e.longTable[h0].offset}
														e.table[h1] = prevEntry{offset: off, prev: e.table[h1].offset}
														index0++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:417
			// _ = "end of CoverTab[92247]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:418
		// _ = "end of CoverTab[92156]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:418
		_go_fuzz_dep_.CoverTab[92157]++

													cv = load6432(src, s)
													if !canRepeat {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:421
			_go_fuzz_dep_.CoverTab[92248]++
														continue
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:422
			// _ = "end of CoverTab[92248]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:423
			_go_fuzz_dep_.CoverTab[92249]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:423
			// _ = "end of CoverTab[92249]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:423
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:423
		// _ = "end of CoverTab[92157]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:423
		_go_fuzz_dep_.CoverTab[92158]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:426
		for {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:426
			_go_fuzz_dep_.CoverTab[92250]++
														o2 := s - offset2
														if load3232(src, o2) != uint32(cv) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:428
				_go_fuzz_dep_.CoverTab[92254]++

															break
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:430
				// _ = "end of CoverTab[92254]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:431
				_go_fuzz_dep_.CoverTab[92255]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:431
				// _ = "end of CoverTab[92255]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:431
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:431
			// _ = "end of CoverTab[92250]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:431
			_go_fuzz_dep_.CoverTab[92251]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:434
			nextHashS := hashLen(cv, bestShortTableBits, bestShortLen)
														nextHashL := hashLen(cv, bestLongTableBits, bestLongLen)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:439
			l := 4 + e.matchlen(s+4, o2+4, src)

														e.longTable[nextHashL] = prevEntry{offset: s + e.cur, prev: e.longTable[nextHashL].offset}
														e.table[nextHashS] = prevEntry{offset: s + e.cur, prev: e.table[nextHashS].offset}
														seq.matchLen = uint32(l) - zstdMinMatch
														seq.litLen = 0

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:447
			seq.offset = 1
			s += l
			nextEmit = s
			if debugSequences {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:450
				_go_fuzz_dep_.CoverTab[92256]++
															println("sequence", seq, "next s:", s)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:451
				// _ = "end of CoverTab[92256]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:452
				_go_fuzz_dep_.CoverTab[92257]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:452
				// _ = "end of CoverTab[92257]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:452
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:452
			// _ = "end of CoverTab[92251]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:452
			_go_fuzz_dep_.CoverTab[92252]++
														blk.sequences = append(blk.sequences, seq)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:456
			offset1, offset2 = offset2, offset1
			if s >= sLimit {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:457
				_go_fuzz_dep_.CoverTab[92258]++

															break encodeLoop
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:459
				// _ = "end of CoverTab[92258]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:460
				_go_fuzz_dep_.CoverTab[92259]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:460
				// _ = "end of CoverTab[92259]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:460
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:460
			// _ = "end of CoverTab[92252]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:460
			_go_fuzz_dep_.CoverTab[92253]++
														cv = load6432(src, s)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:461
			// _ = "end of CoverTab[92253]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:462
		// _ = "end of CoverTab[92158]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:463
	// _ = "end of CoverTab[92107]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:463
	_go_fuzz_dep_.CoverTab[92108]++

												if int(nextEmit) < len(src) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:465
		_go_fuzz_dep_.CoverTab[92260]++
													blk.literals = append(blk.literals, src[nextEmit:]...)
													blk.extraLits = len(src) - int(nextEmit)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:467
		// _ = "end of CoverTab[92260]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:468
		_go_fuzz_dep_.CoverTab[92261]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:468
		// _ = "end of CoverTab[92261]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:468
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:468
	// _ = "end of CoverTab[92108]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:468
	_go_fuzz_dep_.CoverTab[92109]++
												blk.recentOffsets[0] = uint32(offset1)
												blk.recentOffsets[1] = uint32(offset2)
												blk.recentOffsets[2] = uint32(offset3)
												if debugEncoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:472
		_go_fuzz_dep_.CoverTab[92262]++
													println("returning, recent offsets:", blk.recentOffsets, "extra literals:", blk.extraLits)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:473
		// _ = "end of CoverTab[92262]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:474
		_go_fuzz_dep_.CoverTab[92263]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:474
		// _ = "end of CoverTab[92263]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:474
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:474
	// _ = "end of CoverTab[92109]"
}

// EncodeNoHist will encode a block with no history and no following blocks.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:477
// Most notable difference is that src will not be copied for history and
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:477
// we do not need to check for max match length.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:480
func (e *bestFastEncoder) EncodeNoHist(blk *blockEnc, src []byte) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:480
	_go_fuzz_dep_.CoverTab[92264]++
												e.ensureHist(len(src))
												e.Encode(blk, src)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:482
	// _ = "end of CoverTab[92264]"
}

// Reset will reset and set a dictionary if not nil
func (e *bestFastEncoder) Reset(d *dict, singleBlock bool) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:486
	_go_fuzz_dep_.CoverTab[92265]++
												e.resetBase(d, singleBlock)
												if d == nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:488
		_go_fuzz_dep_.CoverTab[92269]++
													return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:489
		// _ = "end of CoverTab[92269]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:490
		_go_fuzz_dep_.CoverTab[92270]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:490
		// _ = "end of CoverTab[92270]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:490
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:490
	// _ = "end of CoverTab[92265]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:490
	_go_fuzz_dep_.CoverTab[92266]++

												if len(e.dictTable) != len(e.table) || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:492
		_go_fuzz_dep_.CoverTab[92271]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:492
		return d.id != e.lastDictID
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:492
		// _ = "end of CoverTab[92271]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:492
	}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:492
		_go_fuzz_dep_.CoverTab[92272]++
													if len(e.dictTable) != len(e.table) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:493
			_go_fuzz_dep_.CoverTab[92275]++
														e.dictTable = make([]prevEntry, len(e.table))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:494
			// _ = "end of CoverTab[92275]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:495
			_go_fuzz_dep_.CoverTab[92276]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:495
			// _ = "end of CoverTab[92276]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:495
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:495
		// _ = "end of CoverTab[92272]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:495
		_go_fuzz_dep_.CoverTab[92273]++
													end := int32(len(d.content)) - 8 + e.maxMatchOff
													for i := e.maxMatchOff; i < end; i += 4 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:497
			_go_fuzz_dep_.CoverTab[92277]++
														const hashLog = bestShortTableBits

														cv := load6432(d.content, i-e.maxMatchOff)
														nextHash := hashLen(cv, hashLog, bestShortLen)
														nextHash1 := hashLen(cv>>8, hashLog, bestShortLen)
														nextHash2 := hashLen(cv>>16, hashLog, bestShortLen)
														nextHash3 := hashLen(cv>>24, hashLog, bestShortLen)
														e.dictTable[nextHash] = prevEntry{
				prev:	e.dictTable[nextHash].offset,
				offset:	i,
			}
			e.dictTable[nextHash1] = prevEntry{
				prev:	e.dictTable[nextHash1].offset,
				offset:	i + 1,
			}
			e.dictTable[nextHash2] = prevEntry{
				prev:	e.dictTable[nextHash2].offset,
				offset:	i + 2,
			}
			e.dictTable[nextHash3] = prevEntry{
				prev:	e.dictTable[nextHash3].offset,
				offset:	i + 3,
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:520
			// _ = "end of CoverTab[92277]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:521
		// _ = "end of CoverTab[92273]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:521
		_go_fuzz_dep_.CoverTab[92274]++
													e.lastDictID = d.id
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:522
		// _ = "end of CoverTab[92274]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:523
		_go_fuzz_dep_.CoverTab[92278]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:523
		// _ = "end of CoverTab[92278]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:523
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:523
	// _ = "end of CoverTab[92266]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:523
	_go_fuzz_dep_.CoverTab[92267]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:526
	if len(e.dictLongTable) != len(e.longTable) || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:526
		_go_fuzz_dep_.CoverTab[92279]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:526
		return d.id != e.lastDictID
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:526
		// _ = "end of CoverTab[92279]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:526
	}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:526
		_go_fuzz_dep_.CoverTab[92280]++
													if len(e.dictLongTable) != len(e.longTable) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:527
			_go_fuzz_dep_.CoverTab[92283]++
														e.dictLongTable = make([]prevEntry, len(e.longTable))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:528
			// _ = "end of CoverTab[92283]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:529
			_go_fuzz_dep_.CoverTab[92284]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:529
			// _ = "end of CoverTab[92284]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:529
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:529
		// _ = "end of CoverTab[92280]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:529
		_go_fuzz_dep_.CoverTab[92281]++
													if len(d.content) >= 8 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:530
			_go_fuzz_dep_.CoverTab[92285]++
														cv := load6432(d.content, 0)
														h := hashLen(cv, bestLongTableBits, bestLongLen)
														e.dictLongTable[h] = prevEntry{
				offset:	e.maxMatchOff,
				prev:	e.dictLongTable[h].offset,
			}

			end := int32(len(d.content)) - 8 + e.maxMatchOff
			off := 8
			for i := e.maxMatchOff + 1; i < end; i++ {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:540
				_go_fuzz_dep_.CoverTab[92286]++
															cv = cv>>8 | (uint64(d.content[off]) << 56)
															h := hashLen(cv, bestLongTableBits, bestLongLen)
															e.dictLongTable[h] = prevEntry{
					offset:	i,
					prev:	e.dictLongTable[h].offset,
				}
															off++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:547
				// _ = "end of CoverTab[92286]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:548
			// _ = "end of CoverTab[92285]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:549
			_go_fuzz_dep_.CoverTab[92287]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:549
			// _ = "end of CoverTab[92287]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:549
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:549
		// _ = "end of CoverTab[92281]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:549
		_go_fuzz_dep_.CoverTab[92282]++
													e.lastDictID = d.id
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:550
		// _ = "end of CoverTab[92282]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:551
		_go_fuzz_dep_.CoverTab[92288]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:551
		// _ = "end of CoverTab[92288]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:551
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:551
	// _ = "end of CoverTab[92267]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:551
	_go_fuzz_dep_.CoverTab[92268]++

												copy(e.longTable[:], e.dictLongTable)

												e.cur = e.maxMatchOff

												copy(e.table[:], e.dictTable)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:557
	// _ = "end of CoverTab[92268]"
}

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:558
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_best.go:558
var _ = _go_fuzz_dep_.CoverTab
