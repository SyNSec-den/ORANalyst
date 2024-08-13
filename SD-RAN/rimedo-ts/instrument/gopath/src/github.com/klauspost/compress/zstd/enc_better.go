// Copyright 2019+ Klaus Post. All rights reserved.
// License information can be found in the LICENSE file.
// Based on work by Yann Collet, released under BSD License.

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:5
package zstd

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:5
)

import "fmt"

const (
	betterLongTableBits	= 19				// Bits used in the long match table
	betterLongTableSize	= 1 << betterLongTableBits	// Size of the table
	betterLongLen		= 8				// Bytes used for table hash

	// Note: Increasing the short table bits or making the hash shorter
	// can actually lead to compression degradation since it will 'steal' more from the
	// long match table and match offsets are quite big.
	// This greatly depends on the type of input.
	betterShortTableBits	= 13				// Bits used in the short match table
	betterShortTableSize	= 1 << betterShortTableBits	// Size of the table
	betterShortLen		= 5				// Bytes used for table hash

	betterLongTableShardCnt		= 1 << (betterLongTableBits - dictShardBits)	// Number of shards in the table
	betterLongTableShardSize	= betterLongTableSize / betterLongTableShardCnt	// Size of an individual shard

	betterShortTableShardCnt	= 1 << (betterShortTableBits - dictShardBits)		// Number of shards in the table
	betterShortTableShardSize	= betterShortTableSize / betterShortTableShardCnt	// Size of an individual shard
)

type prevEntry struct {
	offset	int32
	prev	int32
}

// betterFastEncoder uses 2 tables, one for short matches (5 bytes) and one for long matches.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:34
// The long match table contains the previous entry with the same hash,
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:34
// effectively making it a "chain" of length 2.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:34
// When we find a long match we choose between the two values and select the longest.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:34
// When we find a short match, after checking the long, we check if we can find a long at n+1
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:34
// and that it is longer (lazy matching).
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:40
type betterFastEncoder struct {
	fastBase
	table		[betterShortTableSize]tableEntry
	longTable	[betterLongTableSize]prevEntry
}

type betterFastEncoderDict struct {
	betterFastEncoder
	dictTable		[]tableEntry
	dictLongTable		[]prevEntry
	shortTableShardDirty	[betterShortTableShardCnt]bool
	longTableShardDirty	[betterLongTableShardCnt]bool
	allDirty		bool
}

// Encode improves compression...
func (e *betterFastEncoder) Encode(blk *blockEnc, src []byte) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:56
	_go_fuzz_dep_.CoverTab[92289]++
												const (
		// Input margin is the number of bytes we read (8)
		// and the maximum we will read ahead (2)
		inputMargin		= 8 + 2
		minNonLiteralBlockSize	= 16
	)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:65
	for e.cur >= bufferReset {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:65
		_go_fuzz_dep_.CoverTab[92296]++
													if len(e.hist) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:66
			_go_fuzz_dep_.CoverTab[92300]++
														for i := range e.table[:] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:67
				_go_fuzz_dep_.CoverTab[92303]++
															e.table[i] = tableEntry{}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:68
				// _ = "end of CoverTab[92303]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:69
			// _ = "end of CoverTab[92300]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:69
			_go_fuzz_dep_.CoverTab[92301]++
														for i := range e.longTable[:] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:70
				_go_fuzz_dep_.CoverTab[92304]++
															e.longTable[i] = prevEntry{}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:71
				// _ = "end of CoverTab[92304]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:72
			// _ = "end of CoverTab[92301]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:72
			_go_fuzz_dep_.CoverTab[92302]++
														e.cur = e.maxMatchOff
														break
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:74
			// _ = "end of CoverTab[92302]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:75
			_go_fuzz_dep_.CoverTab[92305]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:75
			// _ = "end of CoverTab[92305]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:75
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:75
		// _ = "end of CoverTab[92296]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:75
		_go_fuzz_dep_.CoverTab[92297]++

													minOff := e.cur + int32(len(e.hist)) - e.maxMatchOff
													for i := range e.table[:] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:78
			_go_fuzz_dep_.CoverTab[92306]++
														v := e.table[i].offset
														if v < minOff {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:80
				_go_fuzz_dep_.CoverTab[92308]++
															v = 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:81
				// _ = "end of CoverTab[92308]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:82
				_go_fuzz_dep_.CoverTab[92309]++
															v = v - e.cur + e.maxMatchOff
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:83
				// _ = "end of CoverTab[92309]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:84
			// _ = "end of CoverTab[92306]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:84
			_go_fuzz_dep_.CoverTab[92307]++
														e.table[i].offset = v
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:85
			// _ = "end of CoverTab[92307]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:86
		// _ = "end of CoverTab[92297]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:86
		_go_fuzz_dep_.CoverTab[92298]++
													for i := range e.longTable[:] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:87
			_go_fuzz_dep_.CoverTab[92310]++
														v := e.longTable[i].offset
														v2 := e.longTable[i].prev
														if v < minOff {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:90
				_go_fuzz_dep_.CoverTab[92312]++
															v = 0
															v2 = 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:92
				// _ = "end of CoverTab[92312]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:93
				_go_fuzz_dep_.CoverTab[92313]++
															v = v - e.cur + e.maxMatchOff
															if v2 < minOff {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:95
					_go_fuzz_dep_.CoverTab[92314]++
																v2 = 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:96
					// _ = "end of CoverTab[92314]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:97
					_go_fuzz_dep_.CoverTab[92315]++
																v2 = v2 - e.cur + e.maxMatchOff
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:98
					// _ = "end of CoverTab[92315]"
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:99
				// _ = "end of CoverTab[92313]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:100
			// _ = "end of CoverTab[92310]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:100
			_go_fuzz_dep_.CoverTab[92311]++
														e.longTable[i] = prevEntry{
				offset:	v,
				prev:	v2,
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:104
			// _ = "end of CoverTab[92311]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:105
		// _ = "end of CoverTab[92298]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:105
		_go_fuzz_dep_.CoverTab[92299]++
													e.cur = e.maxMatchOff
													break
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:107
		// _ = "end of CoverTab[92299]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:108
	// _ = "end of CoverTab[92289]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:108
	_go_fuzz_dep_.CoverTab[92290]++

												s := e.addBlock(src)
												blk.size = len(src)
												if len(src) < minNonLiteralBlockSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:112
		_go_fuzz_dep_.CoverTab[92316]++
													blk.extraLits = len(src)
													blk.literals = blk.literals[:len(src)]
													copy(blk.literals, src)
													return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:116
		// _ = "end of CoverTab[92316]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:117
		_go_fuzz_dep_.CoverTab[92317]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:117
		// _ = "end of CoverTab[92317]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:117
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:117
	// _ = "end of CoverTab[92290]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:117
	_go_fuzz_dep_.CoverTab[92291]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:120
	src = e.hist
	sLimit := int32(len(src)) - inputMargin
	// stepSize is the number of bytes to skip on every main loop iteration.
												// It should be >= 1.
												const stepSize = 1

												const kSearchStrength = 9

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:129
	nextEmit := s
												cv := load6432(src, s)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:133
	offset1 := int32(blk.recentOffsets[0])
	offset2 := int32(blk.recentOffsets[1])

	addLiterals := func(s *seq, until int32) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:136
		_go_fuzz_dep_.CoverTab[92318]++
													if until == nextEmit {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:137
			_go_fuzz_dep_.CoverTab[92320]++
														return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:138
			// _ = "end of CoverTab[92320]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:139
			_go_fuzz_dep_.CoverTab[92321]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:139
			// _ = "end of CoverTab[92321]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:139
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:139
		// _ = "end of CoverTab[92318]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:139
		_go_fuzz_dep_.CoverTab[92319]++
													blk.literals = append(blk.literals, src[nextEmit:until]...)
													s.litLen = uint32(until - nextEmit)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:141
		// _ = "end of CoverTab[92319]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:142
	// _ = "end of CoverTab[92291]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:142
	_go_fuzz_dep_.CoverTab[92292]++
												if debugEncoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:143
		_go_fuzz_dep_.CoverTab[92322]++
													println("recent offsets:", blk.recentOffsets)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:144
		// _ = "end of CoverTab[92322]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:145
		_go_fuzz_dep_.CoverTab[92323]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:145
		// _ = "end of CoverTab[92323]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:145
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:145
	// _ = "end of CoverTab[92292]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:145
	_go_fuzz_dep_.CoverTab[92293]++

encodeLoop:
	for {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:148
		_go_fuzz_dep_.CoverTab[92324]++
													var t int32

													canRepeat := len(blk.sequences) > 2
													var matched int32

													for {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:154
			_go_fuzz_dep_.CoverTab[92336]++
														if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:155
				_go_fuzz_dep_.CoverTab[92343]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:155
				return canRepeat
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:155
				// _ = "end of CoverTab[92343]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:155
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:155
				_go_fuzz_dep_.CoverTab[92344]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:155
				return offset1 == 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:155
				// _ = "end of CoverTab[92344]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:155
			}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:155
				_go_fuzz_dep_.CoverTab[92345]++
															panic("offset0 was 0")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:156
				// _ = "end of CoverTab[92345]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:157
				_go_fuzz_dep_.CoverTab[92346]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:157
				// _ = "end of CoverTab[92346]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:157
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:157
			// _ = "end of CoverTab[92336]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:157
			_go_fuzz_dep_.CoverTab[92337]++

														nextHashS := hashLen(cv, betterShortTableBits, betterShortLen)
														nextHashL := hashLen(cv, betterLongTableBits, betterLongLen)
														candidateL := e.longTable[nextHashL]
														candidateS := e.table[nextHashS]

														const repOff = 1
														repIndex := s - offset1 + repOff
														off := s + e.cur
														e.longTable[nextHashL] = prevEntry{offset: off, prev: candidateL.offset}
														e.table[nextHashS] = tableEntry{offset: off, val: uint32(cv)}

														if canRepeat {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:170
				_go_fuzz_dep_.CoverTab[92347]++
															if repIndex >= 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:171
					_go_fuzz_dep_.CoverTab[92349]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:171
					return load3232(src, repIndex) == uint32(cv>>(repOff*8))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:171
					// _ = "end of CoverTab[92349]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:171
				}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:171
					_go_fuzz_dep_.CoverTab[92350]++
																// Consider history as well.
																var seq seq
																lenght := 4 + e.matchlen(s+4+repOff, repIndex+4, src)

																seq.matchLen = uint32(lenght - zstdMinMatch)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:180
					start := s + repOff

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:183
					startLimit := nextEmit + 1

					tMin := s - e.maxMatchOff
					if tMin < 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:186
						_go_fuzz_dep_.CoverTab[92356]++
																	tMin = 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:187
						// _ = "end of CoverTab[92356]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:188
						_go_fuzz_dep_.CoverTab[92357]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:188
						// _ = "end of CoverTab[92357]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:188
					}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:188
					// _ = "end of CoverTab[92350]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:188
					_go_fuzz_dep_.CoverTab[92351]++
																for repIndex > tMin && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:189
						_go_fuzz_dep_.CoverTab[92358]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:189
						return start > startLimit
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:189
						// _ = "end of CoverTab[92358]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:189
					}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:189
						_go_fuzz_dep_.CoverTab[92359]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:189
						return src[repIndex-1] == src[start-1]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:189
						// _ = "end of CoverTab[92359]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:189
					}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:189
						_go_fuzz_dep_.CoverTab[92360]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:189
						return seq.matchLen < maxMatchLength-zstdMinMatch-1
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:189
						// _ = "end of CoverTab[92360]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:189
					}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:189
						_go_fuzz_dep_.CoverTab[92361]++
																	repIndex--
																	start--
																	seq.matchLen++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:192
						// _ = "end of CoverTab[92361]"
					}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:193
					// _ = "end of CoverTab[92351]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:193
					_go_fuzz_dep_.CoverTab[92352]++
																addLiterals(&seq, start)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:197
					seq.offset = 1
					if debugSequences {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:198
						_go_fuzz_dep_.CoverTab[92362]++
																	println("repeat sequence", seq, "next s:", s)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:199
						// _ = "end of CoverTab[92362]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:200
						_go_fuzz_dep_.CoverTab[92363]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:200
						// _ = "end of CoverTab[92363]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:200
					}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:200
					// _ = "end of CoverTab[92352]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:200
					_go_fuzz_dep_.CoverTab[92353]++
																blk.sequences = append(blk.sequences, seq)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:204
					index0 := s + repOff
					s += lenght + repOff

					nextEmit = s
					if s >= sLimit {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:208
						_go_fuzz_dep_.CoverTab[92364]++
																	if debugEncoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:209
							_go_fuzz_dep_.CoverTab[92366]++
																		println("repeat ended", s, lenght)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:210
							// _ = "end of CoverTab[92366]"

						} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:212
							_go_fuzz_dep_.CoverTab[92367]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:212
							// _ = "end of CoverTab[92367]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:212
						}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:212
						// _ = "end of CoverTab[92364]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:212
						_go_fuzz_dep_.CoverTab[92365]++
																	break encodeLoop
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:213
						// _ = "end of CoverTab[92365]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:214
						_go_fuzz_dep_.CoverTab[92368]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:214
						// _ = "end of CoverTab[92368]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:214
					}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:214
					// _ = "end of CoverTab[92353]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:214
					_go_fuzz_dep_.CoverTab[92354]++

																for index0 < s-1 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:216
						_go_fuzz_dep_.CoverTab[92369]++
																	cv0 := load6432(src, index0)
																	cv1 := cv0 >> 8
																	h0 := hashLen(cv0, betterLongTableBits, betterLongLen)
																	off := index0 + e.cur
																	e.longTable[h0] = prevEntry{offset: off, prev: e.longTable[h0].offset}
																	e.table[hashLen(cv1, betterShortTableBits, betterShortLen)] = tableEntry{offset: off + 1, val: uint32(cv1)}
																	index0 += 2
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:223
						// _ = "end of CoverTab[92369]"
					}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:224
					// _ = "end of CoverTab[92354]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:224
					_go_fuzz_dep_.CoverTab[92355]++
																cv = load6432(src, s)
																continue
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:226
					// _ = "end of CoverTab[92355]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:227
					_go_fuzz_dep_.CoverTab[92370]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:227
					// _ = "end of CoverTab[92370]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:227
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:227
				// _ = "end of CoverTab[92347]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:227
				_go_fuzz_dep_.CoverTab[92348]++
															const repOff2 = 1

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:233
				if false && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:233
					_go_fuzz_dep_.CoverTab[92371]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:233
					return repIndex >= 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:233
					// _ = "end of CoverTab[92371]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:233
				}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:233
					_go_fuzz_dep_.CoverTab[92372]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:233
					return load6432(src, repIndex) == load6432(src, s+repOff)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:233
					// _ = "end of CoverTab[92372]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:233
				}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:233
					_go_fuzz_dep_.CoverTab[92373]++
																// Consider history as well.
																var seq seq
																lenght := 8 + e.matchlen(s+8+repOff2, repIndex+8, src)

																seq.matchLen = uint32(lenght - zstdMinMatch)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:242
					start := s + repOff2

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:245
					startLimit := nextEmit + 1

					tMin := s - e.maxMatchOff
					if tMin < 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:248
						_go_fuzz_dep_.CoverTab[92379]++
																	tMin = 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:249
						// _ = "end of CoverTab[92379]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:250
						_go_fuzz_dep_.CoverTab[92380]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:250
						// _ = "end of CoverTab[92380]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:250
					}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:250
					// _ = "end of CoverTab[92373]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:250
					_go_fuzz_dep_.CoverTab[92374]++
																for repIndex > tMin && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:251
						_go_fuzz_dep_.CoverTab[92381]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:251
						return start > startLimit
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:251
						// _ = "end of CoverTab[92381]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:251
					}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:251
						_go_fuzz_dep_.CoverTab[92382]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:251
						return src[repIndex-1] == src[start-1]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:251
						// _ = "end of CoverTab[92382]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:251
					}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:251
						_go_fuzz_dep_.CoverTab[92383]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:251
						return seq.matchLen < maxMatchLength-zstdMinMatch-1
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:251
						// _ = "end of CoverTab[92383]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:251
					}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:251
						_go_fuzz_dep_.CoverTab[92384]++
																	repIndex--
																	start--
																	seq.matchLen++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:254
						// _ = "end of CoverTab[92384]"
					}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:255
					// _ = "end of CoverTab[92374]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:255
					_go_fuzz_dep_.CoverTab[92375]++
																addLiterals(&seq, start)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:259
					seq.offset = 2
					if debugSequences {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:260
						_go_fuzz_dep_.CoverTab[92385]++
																	println("repeat sequence 2", seq, "next s:", s)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:261
						// _ = "end of CoverTab[92385]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:262
						_go_fuzz_dep_.CoverTab[92386]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:262
						// _ = "end of CoverTab[92386]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:262
					}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:262
					// _ = "end of CoverTab[92375]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:262
					_go_fuzz_dep_.CoverTab[92376]++
																blk.sequences = append(blk.sequences, seq)

																index0 := s + repOff2
																s += lenght + repOff2
																nextEmit = s
																if s >= sLimit {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:268
						_go_fuzz_dep_.CoverTab[92387]++
																	if debugEncoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:269
							_go_fuzz_dep_.CoverTab[92389]++
																		println("repeat ended", s, lenght)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:270
							// _ = "end of CoverTab[92389]"

						} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:272
							_go_fuzz_dep_.CoverTab[92390]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:272
							// _ = "end of CoverTab[92390]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:272
						}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:272
						// _ = "end of CoverTab[92387]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:272
						_go_fuzz_dep_.CoverTab[92388]++
																	break encodeLoop
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:273
						// _ = "end of CoverTab[92388]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:274
						_go_fuzz_dep_.CoverTab[92391]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:274
						// _ = "end of CoverTab[92391]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:274
					}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:274
					// _ = "end of CoverTab[92376]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:274
					_go_fuzz_dep_.CoverTab[92377]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:277
					for index0 < s-1 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:277
						_go_fuzz_dep_.CoverTab[92392]++
																	cv0 := load6432(src, index0)
																	cv1 := cv0 >> 8
																	h0 := hashLen(cv0, betterLongTableBits, betterLongLen)
																	off := index0 + e.cur
																	e.longTable[h0] = prevEntry{offset: off, prev: e.longTable[h0].offset}
																	e.table[hashLen(cv1, betterShortTableBits, betterShortLen)] = tableEntry{offset: off + 1, val: uint32(cv1)}
																	index0 += 2
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:284
						// _ = "end of CoverTab[92392]"
					}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:285
					// _ = "end of CoverTab[92377]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:285
					_go_fuzz_dep_.CoverTab[92378]++
																cv = load6432(src, s)

																offset1, offset2 = offset2, offset1
																continue
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:289
					// _ = "end of CoverTab[92378]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:290
					_go_fuzz_dep_.CoverTab[92393]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:290
					// _ = "end of CoverTab[92393]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:290
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:290
				// _ = "end of CoverTab[92348]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:291
				_go_fuzz_dep_.CoverTab[92394]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:291
				// _ = "end of CoverTab[92394]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:291
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:291
			// _ = "end of CoverTab[92337]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:291
			_go_fuzz_dep_.CoverTab[92338]++

														coffsetL := candidateL.offset - e.cur
														coffsetLP := candidateL.prev - e.cur

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:297
			if s-coffsetL < e.maxMatchOff && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:297
				_go_fuzz_dep_.CoverTab[92395]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:297
				return cv == load6432(src, coffsetL)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:297
				// _ = "end of CoverTab[92395]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:297
			}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:297
				_go_fuzz_dep_.CoverTab[92396]++

															matched = e.matchlen(s+8, coffsetL+8, src) + 8
															t = coffsetL
															if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:301
					_go_fuzz_dep_.CoverTab[92401]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:301
					return s <= t
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:301
					// _ = "end of CoverTab[92401]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:301
				}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:301
					_go_fuzz_dep_.CoverTab[92402]++
																panic(fmt.Sprintf("s (%d) <= t (%d)", s, t))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:302
					// _ = "end of CoverTab[92402]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:303
					_go_fuzz_dep_.CoverTab[92403]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:303
					// _ = "end of CoverTab[92403]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:303
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:303
				// _ = "end of CoverTab[92396]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:303
				_go_fuzz_dep_.CoverTab[92397]++
															if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:304
					_go_fuzz_dep_.CoverTab[92404]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:304
					return s-t > e.maxMatchOff
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:304
					// _ = "end of CoverTab[92404]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:304
				}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:304
					_go_fuzz_dep_.CoverTab[92405]++
																panic("s - t >e.maxMatchOff")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:305
					// _ = "end of CoverTab[92405]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:306
					_go_fuzz_dep_.CoverTab[92406]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:306
					// _ = "end of CoverTab[92406]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:306
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:306
				// _ = "end of CoverTab[92397]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:306
				_go_fuzz_dep_.CoverTab[92398]++
															if debugMatches {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:307
					_go_fuzz_dep_.CoverTab[92407]++
																println("long match")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:308
					// _ = "end of CoverTab[92407]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:309
					_go_fuzz_dep_.CoverTab[92408]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:309
					// _ = "end of CoverTab[92408]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:309
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:309
				// _ = "end of CoverTab[92398]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:309
				_go_fuzz_dep_.CoverTab[92399]++

															if s-coffsetLP < e.maxMatchOff && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:311
					_go_fuzz_dep_.CoverTab[92409]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:311
					return cv == load6432(src, coffsetLP)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:311
					// _ = "end of CoverTab[92409]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:311
				}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:311
					_go_fuzz_dep_.CoverTab[92410]++

																prevMatch := e.matchlen(s+8, coffsetLP+8, src) + 8
																if prevMatch > matched {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:314
						_go_fuzz_dep_.CoverTab[92414]++
																	matched = prevMatch
																	t = coffsetLP
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:316
						// _ = "end of CoverTab[92414]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:317
						_go_fuzz_dep_.CoverTab[92415]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:317
						// _ = "end of CoverTab[92415]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:317
					}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:317
					// _ = "end of CoverTab[92410]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:317
					_go_fuzz_dep_.CoverTab[92411]++
																if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:318
						_go_fuzz_dep_.CoverTab[92416]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:318
						return s <= t
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:318
						// _ = "end of CoverTab[92416]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:318
					}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:318
						_go_fuzz_dep_.CoverTab[92417]++
																	panic(fmt.Sprintf("s (%d) <= t (%d)", s, t))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:319
						// _ = "end of CoverTab[92417]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:320
						_go_fuzz_dep_.CoverTab[92418]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:320
						// _ = "end of CoverTab[92418]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:320
					}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:320
					// _ = "end of CoverTab[92411]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:320
					_go_fuzz_dep_.CoverTab[92412]++
																if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:321
						_go_fuzz_dep_.CoverTab[92419]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:321
						return s-t > e.maxMatchOff
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:321
						// _ = "end of CoverTab[92419]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:321
					}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:321
						_go_fuzz_dep_.CoverTab[92420]++
																	panic("s - t >e.maxMatchOff")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:322
						// _ = "end of CoverTab[92420]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:323
						_go_fuzz_dep_.CoverTab[92421]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:323
						// _ = "end of CoverTab[92421]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:323
					}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:323
					// _ = "end of CoverTab[92412]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:323
					_go_fuzz_dep_.CoverTab[92413]++
																if debugMatches {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:324
						_go_fuzz_dep_.CoverTab[92422]++
																	println("long match")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:325
						// _ = "end of CoverTab[92422]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:326
						_go_fuzz_dep_.CoverTab[92423]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:326
						// _ = "end of CoverTab[92423]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:326
					}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:326
					// _ = "end of CoverTab[92413]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:327
					_go_fuzz_dep_.CoverTab[92424]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:327
					// _ = "end of CoverTab[92424]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:327
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:327
				// _ = "end of CoverTab[92399]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:327
				_go_fuzz_dep_.CoverTab[92400]++
															break
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:328
				// _ = "end of CoverTab[92400]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:329
				_go_fuzz_dep_.CoverTab[92425]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:329
				// _ = "end of CoverTab[92425]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:329
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:329
			// _ = "end of CoverTab[92338]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:329
			_go_fuzz_dep_.CoverTab[92339]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:332
			if s-coffsetLP < e.maxMatchOff && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:332
				_go_fuzz_dep_.CoverTab[92426]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:332
				return cv == load6432(src, coffsetLP)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:332
				// _ = "end of CoverTab[92426]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:332
			}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:332
				_go_fuzz_dep_.CoverTab[92427]++

															matched = e.matchlen(s+8, coffsetLP+8, src) + 8
															t = coffsetLP
															if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:336
					_go_fuzz_dep_.CoverTab[92431]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:336
					return s <= t
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:336
					// _ = "end of CoverTab[92431]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:336
				}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:336
					_go_fuzz_dep_.CoverTab[92432]++
																panic(fmt.Sprintf("s (%d) <= t (%d)", s, t))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:337
					// _ = "end of CoverTab[92432]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:338
					_go_fuzz_dep_.CoverTab[92433]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:338
					// _ = "end of CoverTab[92433]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:338
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:338
				// _ = "end of CoverTab[92427]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:338
				_go_fuzz_dep_.CoverTab[92428]++
															if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:339
					_go_fuzz_dep_.CoverTab[92434]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:339
					return s-t > e.maxMatchOff
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:339
					// _ = "end of CoverTab[92434]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:339
				}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:339
					_go_fuzz_dep_.CoverTab[92435]++
																panic("s - t >e.maxMatchOff")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:340
					// _ = "end of CoverTab[92435]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:341
					_go_fuzz_dep_.CoverTab[92436]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:341
					// _ = "end of CoverTab[92436]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:341
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:341
				// _ = "end of CoverTab[92428]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:341
				_go_fuzz_dep_.CoverTab[92429]++
															if debugMatches {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:342
					_go_fuzz_dep_.CoverTab[92437]++
																println("long match")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:343
					// _ = "end of CoverTab[92437]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:344
					_go_fuzz_dep_.CoverTab[92438]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:344
					// _ = "end of CoverTab[92438]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:344
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:344
				// _ = "end of CoverTab[92429]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:344
				_go_fuzz_dep_.CoverTab[92430]++
															break
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:345
				// _ = "end of CoverTab[92430]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:346
				_go_fuzz_dep_.CoverTab[92439]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:346
				// _ = "end of CoverTab[92439]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:346
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:346
			// _ = "end of CoverTab[92339]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:346
			_go_fuzz_dep_.CoverTab[92340]++

														coffsetS := candidateS.offset - e.cur

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:351
			if s-coffsetS < e.maxMatchOff && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:351
				_go_fuzz_dep_.CoverTab[92440]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:351
				return uint32(cv) == candidateS.val
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:351
				// _ = "end of CoverTab[92440]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:351
			}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:351
				_go_fuzz_dep_.CoverTab[92441]++

															matched = e.matchlen(s+4, coffsetS+4, src) + 4

															// See if we can find a long match at s+1
															const checkAt = 1
															cv := load6432(src, s+checkAt)
															nextHashL = hashLen(cv, betterLongTableBits, betterLongLen)
															candidateL = e.longTable[nextHashL]
															coffsetL = candidateL.offset - e.cur

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:363
				e.longTable[nextHashL] = prevEntry{offset: s + checkAt + e.cur, prev: candidateL.offset}
				if s-coffsetL < e.maxMatchOff && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:364
					_go_fuzz_dep_.CoverTab[92448]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:364
					return cv == load6432(src, coffsetL)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:364
					// _ = "end of CoverTab[92448]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:364
				}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:364
					_go_fuzz_dep_.CoverTab[92449]++

																matchedNext := e.matchlen(s+8+checkAt, coffsetL+8, src) + 8
																if matchedNext > matched {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:367
						_go_fuzz_dep_.CoverTab[92450]++
																	t = coffsetL
																	s += checkAt
																	matched = matchedNext
																	if debugMatches {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:371
							_go_fuzz_dep_.CoverTab[92452]++
																		println("long match (after short)")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:372
							// _ = "end of CoverTab[92452]"
						} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:373
							_go_fuzz_dep_.CoverTab[92453]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:373
							// _ = "end of CoverTab[92453]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:373
						}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:373
						// _ = "end of CoverTab[92450]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:373
						_go_fuzz_dep_.CoverTab[92451]++
																	break
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:374
						// _ = "end of CoverTab[92451]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:375
						_go_fuzz_dep_.CoverTab[92454]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:375
						// _ = "end of CoverTab[92454]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:375
					}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:375
					// _ = "end of CoverTab[92449]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:376
					_go_fuzz_dep_.CoverTab[92455]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:376
					// _ = "end of CoverTab[92455]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:376
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:376
				// _ = "end of CoverTab[92441]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:376
				_go_fuzz_dep_.CoverTab[92442]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:379
				coffsetL = candidateL.prev - e.cur
				if s-coffsetL < e.maxMatchOff && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:380
					_go_fuzz_dep_.CoverTab[92456]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:380
					return cv == load6432(src, coffsetL)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:380
					// _ = "end of CoverTab[92456]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:380
				}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:380
					_go_fuzz_dep_.CoverTab[92457]++

																matchedNext := e.matchlen(s+8+checkAt, coffsetL+8, src) + 8
																if matchedNext > matched {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:383
						_go_fuzz_dep_.CoverTab[92458]++
																	t = coffsetL
																	s += checkAt
																	matched = matchedNext
																	if debugMatches {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:387
							_go_fuzz_dep_.CoverTab[92460]++
																		println("prev long match (after short)")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:388
							// _ = "end of CoverTab[92460]"
						} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:389
							_go_fuzz_dep_.CoverTab[92461]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:389
							// _ = "end of CoverTab[92461]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:389
						}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:389
						// _ = "end of CoverTab[92458]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:389
						_go_fuzz_dep_.CoverTab[92459]++
																	break
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:390
						// _ = "end of CoverTab[92459]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:391
						_go_fuzz_dep_.CoverTab[92462]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:391
						// _ = "end of CoverTab[92462]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:391
					}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:391
					// _ = "end of CoverTab[92457]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:392
					_go_fuzz_dep_.CoverTab[92463]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:392
					// _ = "end of CoverTab[92463]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:392
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:392
				// _ = "end of CoverTab[92442]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:392
				_go_fuzz_dep_.CoverTab[92443]++
															t = coffsetS
															if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:394
					_go_fuzz_dep_.CoverTab[92464]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:394
					return s <= t
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:394
					// _ = "end of CoverTab[92464]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:394
				}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:394
					_go_fuzz_dep_.CoverTab[92465]++
																panic(fmt.Sprintf("s (%d) <= t (%d)", s, t))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:395
					// _ = "end of CoverTab[92465]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:396
					_go_fuzz_dep_.CoverTab[92466]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:396
					// _ = "end of CoverTab[92466]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:396
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:396
				// _ = "end of CoverTab[92443]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:396
				_go_fuzz_dep_.CoverTab[92444]++
															if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:397
					_go_fuzz_dep_.CoverTab[92467]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:397
					return s-t > e.maxMatchOff
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:397
					// _ = "end of CoverTab[92467]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:397
				}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:397
					_go_fuzz_dep_.CoverTab[92468]++
																panic("s - t >e.maxMatchOff")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:398
					// _ = "end of CoverTab[92468]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:399
					_go_fuzz_dep_.CoverTab[92469]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:399
					// _ = "end of CoverTab[92469]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:399
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:399
				// _ = "end of CoverTab[92444]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:399
				_go_fuzz_dep_.CoverTab[92445]++
															if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:400
					_go_fuzz_dep_.CoverTab[92470]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:400
					return t < 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:400
					// _ = "end of CoverTab[92470]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:400
				}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:400
					_go_fuzz_dep_.CoverTab[92471]++
																panic("t<0")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:401
					// _ = "end of CoverTab[92471]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:402
					_go_fuzz_dep_.CoverTab[92472]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:402
					// _ = "end of CoverTab[92472]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:402
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:402
				// _ = "end of CoverTab[92445]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:402
				_go_fuzz_dep_.CoverTab[92446]++
															if debugMatches {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:403
					_go_fuzz_dep_.CoverTab[92473]++
																println("short match")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:404
					// _ = "end of CoverTab[92473]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:405
					_go_fuzz_dep_.CoverTab[92474]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:405
					// _ = "end of CoverTab[92474]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:405
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:405
				// _ = "end of CoverTab[92446]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:405
				_go_fuzz_dep_.CoverTab[92447]++
															break
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:406
				// _ = "end of CoverTab[92447]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:407
				_go_fuzz_dep_.CoverTab[92475]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:407
				// _ = "end of CoverTab[92475]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:407
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:407
			// _ = "end of CoverTab[92340]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:407
			_go_fuzz_dep_.CoverTab[92341]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:410
			s += stepSize + ((s - nextEmit) >> (kSearchStrength - 1))
			if s >= sLimit {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:411
				_go_fuzz_dep_.CoverTab[92476]++
															break encodeLoop
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:412
				// _ = "end of CoverTab[92476]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:413
				_go_fuzz_dep_.CoverTab[92477]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:413
				// _ = "end of CoverTab[92477]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:413
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:413
			// _ = "end of CoverTab[92341]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:413
			_go_fuzz_dep_.CoverTab[92342]++
														cv = load6432(src, s)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:414
			// _ = "end of CoverTab[92342]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:415
		// _ = "end of CoverTab[92324]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:415
		_go_fuzz_dep_.CoverTab[92325]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:418
		if s+matched < sLimit {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:418
			_go_fuzz_dep_.CoverTab[92478]++
														nextHashL := hashLen(load6432(src, s+matched), betterLongTableBits, betterLongLen)
														cv := load3232(src, s)
														candidateL := e.longTable[nextHashL]
														coffsetL := candidateL.offset - e.cur - matched
														if coffsetL >= 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:423
				_go_fuzz_dep_.CoverTab[92480]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:423
				return coffsetL < s
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:423
				// _ = "end of CoverTab[92480]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:423
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:423
				_go_fuzz_dep_.CoverTab[92481]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:423
				return s-coffsetL < e.maxMatchOff
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:423
				// _ = "end of CoverTab[92481]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:423
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:423
				_go_fuzz_dep_.CoverTab[92482]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:423
				return cv == load3232(src, coffsetL)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:423
				// _ = "end of CoverTab[92482]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:423
			}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:423
				_go_fuzz_dep_.CoverTab[92483]++

															matchedNext := e.matchlen(s+4, coffsetL+4, src) + 4
															if matchedNext > matched {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:426
					_go_fuzz_dep_.CoverTab[92484]++
																t = coffsetL
																matched = matchedNext
																if debugMatches {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:429
						_go_fuzz_dep_.CoverTab[92485]++
																	println("long match at end-of-match")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:430
						// _ = "end of CoverTab[92485]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:431
						_go_fuzz_dep_.CoverTab[92486]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:431
						// _ = "end of CoverTab[92486]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:431
					}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:431
					// _ = "end of CoverTab[92484]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:432
					_go_fuzz_dep_.CoverTab[92487]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:432
					// _ = "end of CoverTab[92487]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:432
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:432
				// _ = "end of CoverTab[92483]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:433
				_go_fuzz_dep_.CoverTab[92488]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:433
				// _ = "end of CoverTab[92488]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:433
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:433
			// _ = "end of CoverTab[92478]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:433
			_go_fuzz_dep_.CoverTab[92479]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:436
			if true {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:436
				_go_fuzz_dep_.CoverTab[92489]++
															coffsetL = candidateL.prev - e.cur - matched
															if coffsetL >= 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:438
					_go_fuzz_dep_.CoverTab[92490]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:438
					return coffsetL < s
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:438
					// _ = "end of CoverTab[92490]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:438
				}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:438
					_go_fuzz_dep_.CoverTab[92491]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:438
					return s-coffsetL < e.maxMatchOff
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:438
					// _ = "end of CoverTab[92491]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:438
				}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:438
					_go_fuzz_dep_.CoverTab[92492]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:438
					return cv == load3232(src, coffsetL)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:438
					// _ = "end of CoverTab[92492]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:438
				}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:438
					_go_fuzz_dep_.CoverTab[92493]++

																matchedNext := e.matchlen(s+4, coffsetL+4, src) + 4
																if matchedNext > matched {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:441
						_go_fuzz_dep_.CoverTab[92494]++
																	t = coffsetL
																	matched = matchedNext
																	if debugMatches {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:444
							_go_fuzz_dep_.CoverTab[92495]++
																		println("prev long match at end-of-match")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:445
							// _ = "end of CoverTab[92495]"
						} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:446
							_go_fuzz_dep_.CoverTab[92496]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:446
							// _ = "end of CoverTab[92496]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:446
						}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:446
						// _ = "end of CoverTab[92494]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:447
						_go_fuzz_dep_.CoverTab[92497]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:447
						// _ = "end of CoverTab[92497]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:447
					}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:447
					// _ = "end of CoverTab[92493]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:448
					_go_fuzz_dep_.CoverTab[92498]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:448
					// _ = "end of CoverTab[92498]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:448
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:448
				// _ = "end of CoverTab[92489]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:449
				_go_fuzz_dep_.CoverTab[92499]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:449
				// _ = "end of CoverTab[92499]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:449
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:449
			// _ = "end of CoverTab[92479]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:450
			_go_fuzz_dep_.CoverTab[92500]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:450
			// _ = "end of CoverTab[92500]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:450
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:450
		// _ = "end of CoverTab[92325]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:450
		_go_fuzz_dep_.CoverTab[92326]++

													offset2 = offset1
													offset1 = s - t

													if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:455
			_go_fuzz_dep_.CoverTab[92501]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:455
			return s <= t
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:455
			// _ = "end of CoverTab[92501]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:455
		}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:455
			_go_fuzz_dep_.CoverTab[92502]++
														panic(fmt.Sprintf("s (%d) <= t (%d)", s, t))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:456
			// _ = "end of CoverTab[92502]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:457
			_go_fuzz_dep_.CoverTab[92503]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:457
			// _ = "end of CoverTab[92503]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:457
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:457
		// _ = "end of CoverTab[92326]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:457
		_go_fuzz_dep_.CoverTab[92327]++

													if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:459
			_go_fuzz_dep_.CoverTab[92504]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:459
			return canRepeat
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:459
			// _ = "end of CoverTab[92504]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:459
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:459
			_go_fuzz_dep_.CoverTab[92505]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:459
			return int(offset1) > len(src)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:459
			// _ = "end of CoverTab[92505]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:459
		}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:459
			_go_fuzz_dep_.CoverTab[92506]++
														panic("invalid offset")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:460
			// _ = "end of CoverTab[92506]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:461
			_go_fuzz_dep_.CoverTab[92507]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:461
			// _ = "end of CoverTab[92507]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:461
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:461
		// _ = "end of CoverTab[92327]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:461
		_go_fuzz_dep_.CoverTab[92328]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:464
		l := matched

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:467
		tMin := s - e.maxMatchOff
		if tMin < 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:468
			_go_fuzz_dep_.CoverTab[92508]++
														tMin = 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:469
			// _ = "end of CoverTab[92508]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:470
			_go_fuzz_dep_.CoverTab[92509]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:470
			// _ = "end of CoverTab[92509]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:470
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:470
		// _ = "end of CoverTab[92328]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:470
		_go_fuzz_dep_.CoverTab[92329]++
													for t > tMin && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:471
			_go_fuzz_dep_.CoverTab[92510]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:471
			return s > nextEmit
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:471
			// _ = "end of CoverTab[92510]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:471
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:471
			_go_fuzz_dep_.CoverTab[92511]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:471
			return src[t-1] == src[s-1]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:471
			// _ = "end of CoverTab[92511]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:471
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:471
			_go_fuzz_dep_.CoverTab[92512]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:471
			return l < maxMatchLength
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:471
			// _ = "end of CoverTab[92512]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:471
		}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:471
			_go_fuzz_dep_.CoverTab[92513]++
														s--
														t--
														l++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:474
			// _ = "end of CoverTab[92513]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:475
		// _ = "end of CoverTab[92329]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:475
		_go_fuzz_dep_.CoverTab[92330]++

		// Write our sequence
		var seq seq
		seq.litLen = uint32(s - nextEmit)
		seq.matchLen = uint32(l - zstdMinMatch)
		if seq.litLen > 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:481
			_go_fuzz_dep_.CoverTab[92514]++
														blk.literals = append(blk.literals, src[nextEmit:s]...)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:482
			// _ = "end of CoverTab[92514]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:483
			_go_fuzz_dep_.CoverTab[92515]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:483
			// _ = "end of CoverTab[92515]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:483
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:483
		// _ = "end of CoverTab[92330]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:483
		_go_fuzz_dep_.CoverTab[92331]++
													seq.offset = uint32(s-t) + 3
													s += l
													if debugSequences {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:486
			_go_fuzz_dep_.CoverTab[92516]++
														println("sequence", seq, "next s:", s)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:487
			// _ = "end of CoverTab[92516]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:488
			_go_fuzz_dep_.CoverTab[92517]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:488
			// _ = "end of CoverTab[92517]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:488
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:488
		// _ = "end of CoverTab[92331]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:488
		_go_fuzz_dep_.CoverTab[92332]++
													blk.sequences = append(blk.sequences, seq)
													nextEmit = s
													if s >= sLimit {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:491
			_go_fuzz_dep_.CoverTab[92518]++
														break encodeLoop
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:492
			// _ = "end of CoverTab[92518]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:493
			_go_fuzz_dep_.CoverTab[92519]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:493
			// _ = "end of CoverTab[92519]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:493
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:493
		// _ = "end of CoverTab[92332]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:493
		_go_fuzz_dep_.CoverTab[92333]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:496
		index0 := s - l + 1
		for index0 < s-1 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:497
			_go_fuzz_dep_.CoverTab[92520]++
														cv0 := load6432(src, index0)
														cv1 := cv0 >> 8
														h0 := hashLen(cv0, betterLongTableBits, betterLongLen)
														off := index0 + e.cur
														e.longTable[h0] = prevEntry{offset: off, prev: e.longTable[h0].offset}
														e.table[hashLen(cv1, betterShortTableBits, betterShortLen)] = tableEntry{offset: off + 1, val: uint32(cv1)}
														index0 += 2
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:504
			// _ = "end of CoverTab[92520]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:505
		// _ = "end of CoverTab[92333]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:505
		_go_fuzz_dep_.CoverTab[92334]++

													cv = load6432(src, s)
													if !canRepeat {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:508
			_go_fuzz_dep_.CoverTab[92521]++
														continue
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:509
			// _ = "end of CoverTab[92521]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:510
			_go_fuzz_dep_.CoverTab[92522]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:510
			// _ = "end of CoverTab[92522]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:510
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:510
		// _ = "end of CoverTab[92334]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:510
		_go_fuzz_dep_.CoverTab[92335]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:513
		for {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:513
			_go_fuzz_dep_.CoverTab[92523]++
														o2 := s - offset2
														if load3232(src, o2) != uint32(cv) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:515
				_go_fuzz_dep_.CoverTab[92527]++

															break
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:517
				// _ = "end of CoverTab[92527]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:518
				_go_fuzz_dep_.CoverTab[92528]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:518
				// _ = "end of CoverTab[92528]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:518
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:518
			// _ = "end of CoverTab[92523]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:518
			_go_fuzz_dep_.CoverTab[92524]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:521
			nextHashS := hashLen(cv, betterShortTableBits, betterShortLen)
														nextHashL := hashLen(cv, betterLongTableBits, betterLongLen)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:526
			l := 4 + e.matchlen(s+4, o2+4, src)

														e.longTable[nextHashL] = prevEntry{offset: s + e.cur, prev: e.longTable[nextHashL].offset}
														e.table[nextHashS] = tableEntry{offset: s + e.cur, val: uint32(cv)}
														seq.matchLen = uint32(l) - zstdMinMatch
														seq.litLen = 0

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:534
			seq.offset = 1
			s += l
			nextEmit = s
			if debugSequences {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:537
				_go_fuzz_dep_.CoverTab[92529]++
															println("sequence", seq, "next s:", s)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:538
				// _ = "end of CoverTab[92529]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:539
				_go_fuzz_dep_.CoverTab[92530]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:539
				// _ = "end of CoverTab[92530]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:539
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:539
			// _ = "end of CoverTab[92524]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:539
			_go_fuzz_dep_.CoverTab[92525]++
														blk.sequences = append(blk.sequences, seq)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:543
			offset1, offset2 = offset2, offset1
			if s >= sLimit {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:544
				_go_fuzz_dep_.CoverTab[92531]++

															break encodeLoop
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:546
				// _ = "end of CoverTab[92531]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:547
				_go_fuzz_dep_.CoverTab[92532]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:547
				// _ = "end of CoverTab[92532]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:547
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:547
			// _ = "end of CoverTab[92525]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:547
			_go_fuzz_dep_.CoverTab[92526]++
														cv = load6432(src, s)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:548
			// _ = "end of CoverTab[92526]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:549
		// _ = "end of CoverTab[92335]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:550
	// _ = "end of CoverTab[92293]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:550
	_go_fuzz_dep_.CoverTab[92294]++

												if int(nextEmit) < len(src) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:552
		_go_fuzz_dep_.CoverTab[92533]++
													blk.literals = append(blk.literals, src[nextEmit:]...)
													blk.extraLits = len(src) - int(nextEmit)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:554
		// _ = "end of CoverTab[92533]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:555
		_go_fuzz_dep_.CoverTab[92534]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:555
		// _ = "end of CoverTab[92534]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:555
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:555
	// _ = "end of CoverTab[92294]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:555
	_go_fuzz_dep_.CoverTab[92295]++
												blk.recentOffsets[0] = uint32(offset1)
												blk.recentOffsets[1] = uint32(offset2)
												if debugEncoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:558
		_go_fuzz_dep_.CoverTab[92535]++
													println("returning, recent offsets:", blk.recentOffsets, "extra literals:", blk.extraLits)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:559
		// _ = "end of CoverTab[92535]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:560
		_go_fuzz_dep_.CoverTab[92536]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:560
		// _ = "end of CoverTab[92536]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:560
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:560
	// _ = "end of CoverTab[92295]"
}

// EncodeNoHist will encode a block with no history and no following blocks.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:563
// Most notable difference is that src will not be copied for history and
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:563
// we do not need to check for max match length.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:566
func (e *betterFastEncoder) EncodeNoHist(blk *blockEnc, src []byte) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:566
	_go_fuzz_dep_.CoverTab[92537]++
												e.ensureHist(len(src))
												e.Encode(blk, src)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:568
	// _ = "end of CoverTab[92537]"
}

// Encode improves compression...
func (e *betterFastEncoderDict) Encode(blk *blockEnc, src []byte) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:572
	_go_fuzz_dep_.CoverTab[92538]++
												const (
		// Input margin is the number of bytes we read (8)
		// and the maximum we will read ahead (2)
		inputMargin		= 8 + 2
		minNonLiteralBlockSize	= 16
	)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:581
	for e.cur >= bufferReset {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:581
		_go_fuzz_dep_.CoverTab[92545]++
													if len(e.hist) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:582
			_go_fuzz_dep_.CoverTab[92549]++
														for i := range e.table[:] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:583
				_go_fuzz_dep_.CoverTab[92552]++
															e.table[i] = tableEntry{}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:584
				// _ = "end of CoverTab[92552]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:585
			// _ = "end of CoverTab[92549]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:585
			_go_fuzz_dep_.CoverTab[92550]++
														for i := range e.longTable[:] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:586
				_go_fuzz_dep_.CoverTab[92553]++
															e.longTable[i] = prevEntry{}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:587
				// _ = "end of CoverTab[92553]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:588
			// _ = "end of CoverTab[92550]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:588
			_go_fuzz_dep_.CoverTab[92551]++
														e.cur = e.maxMatchOff
														e.allDirty = true
														break
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:591
			// _ = "end of CoverTab[92551]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:592
			_go_fuzz_dep_.CoverTab[92554]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:592
			// _ = "end of CoverTab[92554]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:592
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:592
		// _ = "end of CoverTab[92545]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:592
		_go_fuzz_dep_.CoverTab[92546]++

													minOff := e.cur + int32(len(e.hist)) - e.maxMatchOff
													for i := range e.table[:] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:595
			_go_fuzz_dep_.CoverTab[92555]++
														v := e.table[i].offset
														if v < minOff {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:597
				_go_fuzz_dep_.CoverTab[92557]++
															v = 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:598
				// _ = "end of CoverTab[92557]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:599
				_go_fuzz_dep_.CoverTab[92558]++
															v = v - e.cur + e.maxMatchOff
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:600
				// _ = "end of CoverTab[92558]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:601
			// _ = "end of CoverTab[92555]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:601
			_go_fuzz_dep_.CoverTab[92556]++
														e.table[i].offset = v
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:602
			// _ = "end of CoverTab[92556]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:603
		// _ = "end of CoverTab[92546]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:603
		_go_fuzz_dep_.CoverTab[92547]++
													for i := range e.longTable[:] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:604
			_go_fuzz_dep_.CoverTab[92559]++
														v := e.longTable[i].offset
														v2 := e.longTable[i].prev
														if v < minOff {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:607
				_go_fuzz_dep_.CoverTab[92561]++
															v = 0
															v2 = 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:609
				// _ = "end of CoverTab[92561]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:610
				_go_fuzz_dep_.CoverTab[92562]++
															v = v - e.cur + e.maxMatchOff
															if v2 < minOff {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:612
					_go_fuzz_dep_.CoverTab[92563]++
																v2 = 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:613
					// _ = "end of CoverTab[92563]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:614
					_go_fuzz_dep_.CoverTab[92564]++
																v2 = v2 - e.cur + e.maxMatchOff
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:615
					// _ = "end of CoverTab[92564]"
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:616
				// _ = "end of CoverTab[92562]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:617
			// _ = "end of CoverTab[92559]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:617
			_go_fuzz_dep_.CoverTab[92560]++
														e.longTable[i] = prevEntry{
				offset:	v,
				prev:	v2,
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:621
			// _ = "end of CoverTab[92560]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:622
		// _ = "end of CoverTab[92547]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:622
		_go_fuzz_dep_.CoverTab[92548]++
													e.allDirty = true
													e.cur = e.maxMatchOff
													break
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:625
		// _ = "end of CoverTab[92548]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:626
	// _ = "end of CoverTab[92538]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:626
	_go_fuzz_dep_.CoverTab[92539]++

												s := e.addBlock(src)
												blk.size = len(src)
												if len(src) < minNonLiteralBlockSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:630
		_go_fuzz_dep_.CoverTab[92565]++
													blk.extraLits = len(src)
													blk.literals = blk.literals[:len(src)]
													copy(blk.literals, src)
													return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:634
		// _ = "end of CoverTab[92565]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:635
		_go_fuzz_dep_.CoverTab[92566]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:635
		// _ = "end of CoverTab[92566]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:635
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:635
	// _ = "end of CoverTab[92539]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:635
	_go_fuzz_dep_.CoverTab[92540]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:638
	src = e.hist
	sLimit := int32(len(src)) - inputMargin
	// stepSize is the number of bytes to skip on every main loop iteration.
												// It should be >= 1.
												const stepSize = 1

												const kSearchStrength = 9

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:647
	nextEmit := s
												cv := load6432(src, s)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:651
	offset1 := int32(blk.recentOffsets[0])
	offset2 := int32(blk.recentOffsets[1])

	addLiterals := func(s *seq, until int32) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:654
		_go_fuzz_dep_.CoverTab[92567]++
													if until == nextEmit {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:655
			_go_fuzz_dep_.CoverTab[92569]++
														return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:656
			// _ = "end of CoverTab[92569]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:657
			_go_fuzz_dep_.CoverTab[92570]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:657
			// _ = "end of CoverTab[92570]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:657
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:657
		// _ = "end of CoverTab[92567]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:657
		_go_fuzz_dep_.CoverTab[92568]++
													blk.literals = append(blk.literals, src[nextEmit:until]...)
													s.litLen = uint32(until - nextEmit)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:659
		// _ = "end of CoverTab[92568]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:660
	// _ = "end of CoverTab[92540]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:660
	_go_fuzz_dep_.CoverTab[92541]++
												if debugEncoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:661
		_go_fuzz_dep_.CoverTab[92571]++
													println("recent offsets:", blk.recentOffsets)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:662
		// _ = "end of CoverTab[92571]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:663
		_go_fuzz_dep_.CoverTab[92572]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:663
		// _ = "end of CoverTab[92572]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:663
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:663
	// _ = "end of CoverTab[92541]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:663
	_go_fuzz_dep_.CoverTab[92542]++

encodeLoop:
	for {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:666
		_go_fuzz_dep_.CoverTab[92573]++
													var t int32

													canRepeat := len(blk.sequences) > 2
													var matched int32

													for {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:672
			_go_fuzz_dep_.CoverTab[92585]++
														if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:673
				_go_fuzz_dep_.CoverTab[92592]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:673
				return canRepeat
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:673
				// _ = "end of CoverTab[92592]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:673
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:673
				_go_fuzz_dep_.CoverTab[92593]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:673
				return offset1 == 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:673
				// _ = "end of CoverTab[92593]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:673
			}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:673
				_go_fuzz_dep_.CoverTab[92594]++
															panic("offset0 was 0")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:674
				// _ = "end of CoverTab[92594]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:675
				_go_fuzz_dep_.CoverTab[92595]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:675
				// _ = "end of CoverTab[92595]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:675
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:675
			// _ = "end of CoverTab[92585]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:675
			_go_fuzz_dep_.CoverTab[92586]++

														nextHashS := hashLen(cv, betterShortTableBits, betterShortLen)
														nextHashL := hashLen(cv, betterLongTableBits, betterLongLen)
														candidateL := e.longTable[nextHashL]
														candidateS := e.table[nextHashS]

														const repOff = 1
														repIndex := s - offset1 + repOff
														off := s + e.cur
														e.longTable[nextHashL] = prevEntry{offset: off, prev: candidateL.offset}
														e.markLongShardDirty(nextHashL)
														e.table[nextHashS] = tableEntry{offset: off, val: uint32(cv)}
														e.markShortShardDirty(nextHashS)

														if canRepeat {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:690
				_go_fuzz_dep_.CoverTab[92596]++
															if repIndex >= 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:691
					_go_fuzz_dep_.CoverTab[92598]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:691
					return load3232(src, repIndex) == uint32(cv>>(repOff*8))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:691
					// _ = "end of CoverTab[92598]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:691
				}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:691
					_go_fuzz_dep_.CoverTab[92599]++
																// Consider history as well.
																var seq seq
																lenght := 4 + e.matchlen(s+4+repOff, repIndex+4, src)

																seq.matchLen = uint32(lenght - zstdMinMatch)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:700
					start := s + repOff

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:703
					startLimit := nextEmit + 1

					tMin := s - e.maxMatchOff
					if tMin < 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:706
						_go_fuzz_dep_.CoverTab[92605]++
																	tMin = 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:707
						// _ = "end of CoverTab[92605]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:708
						_go_fuzz_dep_.CoverTab[92606]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:708
						// _ = "end of CoverTab[92606]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:708
					}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:708
					// _ = "end of CoverTab[92599]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:708
					_go_fuzz_dep_.CoverTab[92600]++
																for repIndex > tMin && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:709
						_go_fuzz_dep_.CoverTab[92607]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:709
						return start > startLimit
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:709
						// _ = "end of CoverTab[92607]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:709
					}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:709
						_go_fuzz_dep_.CoverTab[92608]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:709
						return src[repIndex-1] == src[start-1]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:709
						// _ = "end of CoverTab[92608]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:709
					}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:709
						_go_fuzz_dep_.CoverTab[92609]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:709
						return seq.matchLen < maxMatchLength-zstdMinMatch-1
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:709
						// _ = "end of CoverTab[92609]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:709
					}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:709
						_go_fuzz_dep_.CoverTab[92610]++
																	repIndex--
																	start--
																	seq.matchLen++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:712
						// _ = "end of CoverTab[92610]"
					}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:713
					// _ = "end of CoverTab[92600]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:713
					_go_fuzz_dep_.CoverTab[92601]++
																addLiterals(&seq, start)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:717
					seq.offset = 1
					if debugSequences {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:718
						_go_fuzz_dep_.CoverTab[92611]++
																	println("repeat sequence", seq, "next s:", s)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:719
						// _ = "end of CoverTab[92611]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:720
						_go_fuzz_dep_.CoverTab[92612]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:720
						// _ = "end of CoverTab[92612]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:720
					}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:720
					// _ = "end of CoverTab[92601]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:720
					_go_fuzz_dep_.CoverTab[92602]++
																blk.sequences = append(blk.sequences, seq)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:724
					index0 := s + repOff
					s += lenght + repOff

					nextEmit = s
					if s >= sLimit {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:728
						_go_fuzz_dep_.CoverTab[92613]++
																	if debugEncoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:729
							_go_fuzz_dep_.CoverTab[92615]++
																		println("repeat ended", s, lenght)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:730
							// _ = "end of CoverTab[92615]"

						} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:732
							_go_fuzz_dep_.CoverTab[92616]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:732
							// _ = "end of CoverTab[92616]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:732
						}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:732
						// _ = "end of CoverTab[92613]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:732
						_go_fuzz_dep_.CoverTab[92614]++
																	break encodeLoop
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:733
						// _ = "end of CoverTab[92614]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:734
						_go_fuzz_dep_.CoverTab[92617]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:734
						// _ = "end of CoverTab[92617]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:734
					}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:734
					// _ = "end of CoverTab[92602]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:734
					_go_fuzz_dep_.CoverTab[92603]++

																for index0 < s-1 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:736
						_go_fuzz_dep_.CoverTab[92618]++
																	cv0 := load6432(src, index0)
																	cv1 := cv0 >> 8
																	h0 := hashLen(cv0, betterLongTableBits, betterLongLen)
																	off := index0 + e.cur
																	e.longTable[h0] = prevEntry{offset: off, prev: e.longTable[h0].offset}
																	e.markLongShardDirty(h0)
																	h1 := hashLen(cv1, betterShortTableBits, betterShortLen)
																	e.table[h1] = tableEntry{offset: off + 1, val: uint32(cv1)}
																	e.markShortShardDirty(h1)
																	index0 += 2
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:746
						// _ = "end of CoverTab[92618]"
					}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:747
					// _ = "end of CoverTab[92603]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:747
					_go_fuzz_dep_.CoverTab[92604]++
																cv = load6432(src, s)
																continue
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:749
					// _ = "end of CoverTab[92604]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:750
					_go_fuzz_dep_.CoverTab[92619]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:750
					// _ = "end of CoverTab[92619]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:750
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:750
				// _ = "end of CoverTab[92596]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:750
				_go_fuzz_dep_.CoverTab[92597]++
															const repOff2 = 1

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:756
				if false && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:756
					_go_fuzz_dep_.CoverTab[92620]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:756
					return repIndex >= 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:756
					// _ = "end of CoverTab[92620]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:756
				}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:756
					_go_fuzz_dep_.CoverTab[92621]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:756
					return load6432(src, repIndex) == load6432(src, s+repOff)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:756
					// _ = "end of CoverTab[92621]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:756
				}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:756
					_go_fuzz_dep_.CoverTab[92622]++
																// Consider history as well.
																var seq seq
																lenght := 8 + e.matchlen(s+8+repOff2, repIndex+8, src)

																seq.matchLen = uint32(lenght - zstdMinMatch)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:765
					start := s + repOff2

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:768
					startLimit := nextEmit + 1

					tMin := s - e.maxMatchOff
					if tMin < 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:771
						_go_fuzz_dep_.CoverTab[92628]++
																	tMin = 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:772
						// _ = "end of CoverTab[92628]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:773
						_go_fuzz_dep_.CoverTab[92629]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:773
						// _ = "end of CoverTab[92629]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:773
					}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:773
					// _ = "end of CoverTab[92622]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:773
					_go_fuzz_dep_.CoverTab[92623]++
																for repIndex > tMin && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:774
						_go_fuzz_dep_.CoverTab[92630]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:774
						return start > startLimit
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:774
						// _ = "end of CoverTab[92630]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:774
					}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:774
						_go_fuzz_dep_.CoverTab[92631]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:774
						return src[repIndex-1] == src[start-1]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:774
						// _ = "end of CoverTab[92631]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:774
					}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:774
						_go_fuzz_dep_.CoverTab[92632]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:774
						return seq.matchLen < maxMatchLength-zstdMinMatch-1
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:774
						// _ = "end of CoverTab[92632]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:774
					}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:774
						_go_fuzz_dep_.CoverTab[92633]++
																	repIndex--
																	start--
																	seq.matchLen++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:777
						// _ = "end of CoverTab[92633]"
					}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:778
					// _ = "end of CoverTab[92623]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:778
					_go_fuzz_dep_.CoverTab[92624]++
																addLiterals(&seq, start)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:782
					seq.offset = 2
					if debugSequences {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:783
						_go_fuzz_dep_.CoverTab[92634]++
																	println("repeat sequence 2", seq, "next s:", s)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:784
						// _ = "end of CoverTab[92634]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:785
						_go_fuzz_dep_.CoverTab[92635]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:785
						// _ = "end of CoverTab[92635]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:785
					}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:785
					// _ = "end of CoverTab[92624]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:785
					_go_fuzz_dep_.CoverTab[92625]++
																blk.sequences = append(blk.sequences, seq)

																index0 := s + repOff2
																s += lenght + repOff2
																nextEmit = s
																if s >= sLimit {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:791
						_go_fuzz_dep_.CoverTab[92636]++
																	if debugEncoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:792
							_go_fuzz_dep_.CoverTab[92638]++
																		println("repeat ended", s, lenght)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:793
							// _ = "end of CoverTab[92638]"

						} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:795
							_go_fuzz_dep_.CoverTab[92639]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:795
							// _ = "end of CoverTab[92639]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:795
						}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:795
						// _ = "end of CoverTab[92636]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:795
						_go_fuzz_dep_.CoverTab[92637]++
																	break encodeLoop
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:796
						// _ = "end of CoverTab[92637]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:797
						_go_fuzz_dep_.CoverTab[92640]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:797
						// _ = "end of CoverTab[92640]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:797
					}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:797
					// _ = "end of CoverTab[92625]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:797
					_go_fuzz_dep_.CoverTab[92626]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:800
					for index0 < s-1 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:800
						_go_fuzz_dep_.CoverTab[92641]++
																	cv0 := load6432(src, index0)
																	cv1 := cv0 >> 8
																	h0 := hashLen(cv0, betterLongTableBits, betterLongLen)
																	off := index0 + e.cur
																	e.longTable[h0] = prevEntry{offset: off, prev: e.longTable[h0].offset}
																	e.markLongShardDirty(h0)
																	h1 := hashLen(cv1, betterShortTableBits, betterShortLen)
																	e.table[h1] = tableEntry{offset: off + 1, val: uint32(cv1)}
																	e.markShortShardDirty(h1)
																	index0 += 2
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:810
						// _ = "end of CoverTab[92641]"
					}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:811
					// _ = "end of CoverTab[92626]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:811
					_go_fuzz_dep_.CoverTab[92627]++
																cv = load6432(src, s)

																offset1, offset2 = offset2, offset1
																continue
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:815
					// _ = "end of CoverTab[92627]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:816
					_go_fuzz_dep_.CoverTab[92642]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:816
					// _ = "end of CoverTab[92642]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:816
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:816
				// _ = "end of CoverTab[92597]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:817
				_go_fuzz_dep_.CoverTab[92643]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:817
				// _ = "end of CoverTab[92643]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:817
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:817
			// _ = "end of CoverTab[92586]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:817
			_go_fuzz_dep_.CoverTab[92587]++

														coffsetL := candidateL.offset - e.cur
														coffsetLP := candidateL.prev - e.cur

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:823
			if s-coffsetL < e.maxMatchOff && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:823
				_go_fuzz_dep_.CoverTab[92644]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:823
				return cv == load6432(src, coffsetL)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:823
				// _ = "end of CoverTab[92644]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:823
			}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:823
				_go_fuzz_dep_.CoverTab[92645]++

															matched = e.matchlen(s+8, coffsetL+8, src) + 8
															t = coffsetL
															if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:827
					_go_fuzz_dep_.CoverTab[92650]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:827
					return s <= t
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:827
					// _ = "end of CoverTab[92650]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:827
				}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:827
					_go_fuzz_dep_.CoverTab[92651]++
																panic(fmt.Sprintf("s (%d) <= t (%d)", s, t))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:828
					// _ = "end of CoverTab[92651]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:829
					_go_fuzz_dep_.CoverTab[92652]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:829
					// _ = "end of CoverTab[92652]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:829
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:829
				// _ = "end of CoverTab[92645]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:829
				_go_fuzz_dep_.CoverTab[92646]++
															if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:830
					_go_fuzz_dep_.CoverTab[92653]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:830
					return s-t > e.maxMatchOff
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:830
					// _ = "end of CoverTab[92653]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:830
				}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:830
					_go_fuzz_dep_.CoverTab[92654]++
																panic("s - t >e.maxMatchOff")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:831
					// _ = "end of CoverTab[92654]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:832
					_go_fuzz_dep_.CoverTab[92655]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:832
					// _ = "end of CoverTab[92655]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:832
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:832
				// _ = "end of CoverTab[92646]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:832
				_go_fuzz_dep_.CoverTab[92647]++
															if debugMatches {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:833
					_go_fuzz_dep_.CoverTab[92656]++
																println("long match")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:834
					// _ = "end of CoverTab[92656]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:835
					_go_fuzz_dep_.CoverTab[92657]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:835
					// _ = "end of CoverTab[92657]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:835
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:835
				// _ = "end of CoverTab[92647]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:835
				_go_fuzz_dep_.CoverTab[92648]++

															if s-coffsetLP < e.maxMatchOff && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:837
					_go_fuzz_dep_.CoverTab[92658]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:837
					return cv == load6432(src, coffsetLP)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:837
					// _ = "end of CoverTab[92658]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:837
				}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:837
					_go_fuzz_dep_.CoverTab[92659]++

																prevMatch := e.matchlen(s+8, coffsetLP+8, src) + 8
																if prevMatch > matched {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:840
						_go_fuzz_dep_.CoverTab[92663]++
																	matched = prevMatch
																	t = coffsetLP
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:842
						// _ = "end of CoverTab[92663]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:843
						_go_fuzz_dep_.CoverTab[92664]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:843
						// _ = "end of CoverTab[92664]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:843
					}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:843
					// _ = "end of CoverTab[92659]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:843
					_go_fuzz_dep_.CoverTab[92660]++
																if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:844
						_go_fuzz_dep_.CoverTab[92665]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:844
						return s <= t
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:844
						// _ = "end of CoverTab[92665]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:844
					}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:844
						_go_fuzz_dep_.CoverTab[92666]++
																	panic(fmt.Sprintf("s (%d) <= t (%d)", s, t))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:845
						// _ = "end of CoverTab[92666]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:846
						_go_fuzz_dep_.CoverTab[92667]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:846
						// _ = "end of CoverTab[92667]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:846
					}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:846
					// _ = "end of CoverTab[92660]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:846
					_go_fuzz_dep_.CoverTab[92661]++
																if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:847
						_go_fuzz_dep_.CoverTab[92668]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:847
						return s-t > e.maxMatchOff
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:847
						// _ = "end of CoverTab[92668]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:847
					}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:847
						_go_fuzz_dep_.CoverTab[92669]++
																	panic("s - t >e.maxMatchOff")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:848
						// _ = "end of CoverTab[92669]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:849
						_go_fuzz_dep_.CoverTab[92670]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:849
						// _ = "end of CoverTab[92670]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:849
					}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:849
					// _ = "end of CoverTab[92661]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:849
					_go_fuzz_dep_.CoverTab[92662]++
																if debugMatches {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:850
						_go_fuzz_dep_.CoverTab[92671]++
																	println("long match")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:851
						// _ = "end of CoverTab[92671]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:852
						_go_fuzz_dep_.CoverTab[92672]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:852
						// _ = "end of CoverTab[92672]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:852
					}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:852
					// _ = "end of CoverTab[92662]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:853
					_go_fuzz_dep_.CoverTab[92673]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:853
					// _ = "end of CoverTab[92673]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:853
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:853
				// _ = "end of CoverTab[92648]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:853
				_go_fuzz_dep_.CoverTab[92649]++
															break
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:854
				// _ = "end of CoverTab[92649]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:855
				_go_fuzz_dep_.CoverTab[92674]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:855
				// _ = "end of CoverTab[92674]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:855
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:855
			// _ = "end of CoverTab[92587]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:855
			_go_fuzz_dep_.CoverTab[92588]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:858
			if s-coffsetLP < e.maxMatchOff && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:858
				_go_fuzz_dep_.CoverTab[92675]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:858
				return cv == load6432(src, coffsetLP)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:858
				// _ = "end of CoverTab[92675]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:858
			}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:858
				_go_fuzz_dep_.CoverTab[92676]++

															matched = e.matchlen(s+8, coffsetLP+8, src) + 8
															t = coffsetLP
															if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:862
					_go_fuzz_dep_.CoverTab[92680]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:862
					return s <= t
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:862
					// _ = "end of CoverTab[92680]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:862
				}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:862
					_go_fuzz_dep_.CoverTab[92681]++
																panic(fmt.Sprintf("s (%d) <= t (%d)", s, t))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:863
					// _ = "end of CoverTab[92681]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:864
					_go_fuzz_dep_.CoverTab[92682]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:864
					// _ = "end of CoverTab[92682]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:864
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:864
				// _ = "end of CoverTab[92676]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:864
				_go_fuzz_dep_.CoverTab[92677]++
															if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:865
					_go_fuzz_dep_.CoverTab[92683]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:865
					return s-t > e.maxMatchOff
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:865
					// _ = "end of CoverTab[92683]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:865
				}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:865
					_go_fuzz_dep_.CoverTab[92684]++
																panic("s - t >e.maxMatchOff")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:866
					// _ = "end of CoverTab[92684]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:867
					_go_fuzz_dep_.CoverTab[92685]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:867
					// _ = "end of CoverTab[92685]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:867
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:867
				// _ = "end of CoverTab[92677]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:867
				_go_fuzz_dep_.CoverTab[92678]++
															if debugMatches {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:868
					_go_fuzz_dep_.CoverTab[92686]++
																println("long match")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:869
					// _ = "end of CoverTab[92686]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:870
					_go_fuzz_dep_.CoverTab[92687]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:870
					// _ = "end of CoverTab[92687]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:870
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:870
				// _ = "end of CoverTab[92678]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:870
				_go_fuzz_dep_.CoverTab[92679]++
															break
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:871
				// _ = "end of CoverTab[92679]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:872
				_go_fuzz_dep_.CoverTab[92688]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:872
				// _ = "end of CoverTab[92688]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:872
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:872
			// _ = "end of CoverTab[92588]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:872
			_go_fuzz_dep_.CoverTab[92589]++

														coffsetS := candidateS.offset - e.cur

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:877
			if s-coffsetS < e.maxMatchOff && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:877
				_go_fuzz_dep_.CoverTab[92689]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:877
				return uint32(cv) == candidateS.val
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:877
				// _ = "end of CoverTab[92689]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:877
			}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:877
				_go_fuzz_dep_.CoverTab[92690]++

															matched = e.matchlen(s+4, coffsetS+4, src) + 4

															// See if we can find a long match at s+1
															const checkAt = 1
															cv := load6432(src, s+checkAt)
															nextHashL = hashLen(cv, betterLongTableBits, betterLongLen)
															candidateL = e.longTable[nextHashL]
															coffsetL = candidateL.offset - e.cur

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:889
				e.longTable[nextHashL] = prevEntry{offset: s + checkAt + e.cur, prev: candidateL.offset}
				e.markLongShardDirty(nextHashL)
				if s-coffsetL < e.maxMatchOff && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:891
					_go_fuzz_dep_.CoverTab[92697]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:891
					return cv == load6432(src, coffsetL)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:891
					// _ = "end of CoverTab[92697]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:891
				}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:891
					_go_fuzz_dep_.CoverTab[92698]++

																matchedNext := e.matchlen(s+8+checkAt, coffsetL+8, src) + 8
																if matchedNext > matched {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:894
						_go_fuzz_dep_.CoverTab[92699]++
																	t = coffsetL
																	s += checkAt
																	matched = matchedNext
																	if debugMatches {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:898
							_go_fuzz_dep_.CoverTab[92701]++
																		println("long match (after short)")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:899
							// _ = "end of CoverTab[92701]"
						} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:900
							_go_fuzz_dep_.CoverTab[92702]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:900
							// _ = "end of CoverTab[92702]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:900
						}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:900
						// _ = "end of CoverTab[92699]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:900
						_go_fuzz_dep_.CoverTab[92700]++
																	break
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:901
						// _ = "end of CoverTab[92700]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:902
						_go_fuzz_dep_.CoverTab[92703]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:902
						// _ = "end of CoverTab[92703]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:902
					}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:902
					// _ = "end of CoverTab[92698]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:903
					_go_fuzz_dep_.CoverTab[92704]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:903
					// _ = "end of CoverTab[92704]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:903
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:903
				// _ = "end of CoverTab[92690]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:903
				_go_fuzz_dep_.CoverTab[92691]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:906
				coffsetL = candidateL.prev - e.cur
				if s-coffsetL < e.maxMatchOff && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:907
					_go_fuzz_dep_.CoverTab[92705]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:907
					return cv == load6432(src, coffsetL)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:907
					// _ = "end of CoverTab[92705]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:907
				}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:907
					_go_fuzz_dep_.CoverTab[92706]++

																matchedNext := e.matchlen(s+8+checkAt, coffsetL+8, src) + 8
																if matchedNext > matched {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:910
						_go_fuzz_dep_.CoverTab[92707]++
																	t = coffsetL
																	s += checkAt
																	matched = matchedNext
																	if debugMatches {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:914
							_go_fuzz_dep_.CoverTab[92709]++
																		println("prev long match (after short)")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:915
							// _ = "end of CoverTab[92709]"
						} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:916
							_go_fuzz_dep_.CoverTab[92710]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:916
							// _ = "end of CoverTab[92710]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:916
						}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:916
						// _ = "end of CoverTab[92707]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:916
						_go_fuzz_dep_.CoverTab[92708]++
																	break
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:917
						// _ = "end of CoverTab[92708]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:918
						_go_fuzz_dep_.CoverTab[92711]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:918
						// _ = "end of CoverTab[92711]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:918
					}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:918
					// _ = "end of CoverTab[92706]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:919
					_go_fuzz_dep_.CoverTab[92712]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:919
					// _ = "end of CoverTab[92712]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:919
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:919
				// _ = "end of CoverTab[92691]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:919
				_go_fuzz_dep_.CoverTab[92692]++
															t = coffsetS
															if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:921
					_go_fuzz_dep_.CoverTab[92713]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:921
					return s <= t
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:921
					// _ = "end of CoverTab[92713]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:921
				}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:921
					_go_fuzz_dep_.CoverTab[92714]++
																panic(fmt.Sprintf("s (%d) <= t (%d)", s, t))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:922
					// _ = "end of CoverTab[92714]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:923
					_go_fuzz_dep_.CoverTab[92715]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:923
					// _ = "end of CoverTab[92715]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:923
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:923
				// _ = "end of CoverTab[92692]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:923
				_go_fuzz_dep_.CoverTab[92693]++
															if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:924
					_go_fuzz_dep_.CoverTab[92716]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:924
					return s-t > e.maxMatchOff
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:924
					// _ = "end of CoverTab[92716]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:924
				}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:924
					_go_fuzz_dep_.CoverTab[92717]++
																panic("s - t >e.maxMatchOff")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:925
					// _ = "end of CoverTab[92717]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:926
					_go_fuzz_dep_.CoverTab[92718]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:926
					// _ = "end of CoverTab[92718]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:926
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:926
				// _ = "end of CoverTab[92693]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:926
				_go_fuzz_dep_.CoverTab[92694]++
															if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:927
					_go_fuzz_dep_.CoverTab[92719]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:927
					return t < 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:927
					// _ = "end of CoverTab[92719]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:927
				}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:927
					_go_fuzz_dep_.CoverTab[92720]++
																panic("t<0")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:928
					// _ = "end of CoverTab[92720]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:929
					_go_fuzz_dep_.CoverTab[92721]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:929
					// _ = "end of CoverTab[92721]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:929
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:929
				// _ = "end of CoverTab[92694]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:929
				_go_fuzz_dep_.CoverTab[92695]++
															if debugMatches {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:930
					_go_fuzz_dep_.CoverTab[92722]++
																println("short match")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:931
					// _ = "end of CoverTab[92722]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:932
					_go_fuzz_dep_.CoverTab[92723]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:932
					// _ = "end of CoverTab[92723]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:932
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:932
				// _ = "end of CoverTab[92695]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:932
				_go_fuzz_dep_.CoverTab[92696]++
															break
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:933
				// _ = "end of CoverTab[92696]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:934
				_go_fuzz_dep_.CoverTab[92724]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:934
				// _ = "end of CoverTab[92724]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:934
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:934
			// _ = "end of CoverTab[92589]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:934
			_go_fuzz_dep_.CoverTab[92590]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:937
			s += stepSize + ((s - nextEmit) >> (kSearchStrength - 1))
			if s >= sLimit {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:938
				_go_fuzz_dep_.CoverTab[92725]++
															break encodeLoop
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:939
				// _ = "end of CoverTab[92725]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:940
				_go_fuzz_dep_.CoverTab[92726]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:940
				// _ = "end of CoverTab[92726]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:940
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:940
			// _ = "end of CoverTab[92590]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:940
			_go_fuzz_dep_.CoverTab[92591]++
														cv = load6432(src, s)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:941
			// _ = "end of CoverTab[92591]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:942
		// _ = "end of CoverTab[92573]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:942
		_go_fuzz_dep_.CoverTab[92574]++

													if s+matched < sLimit {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:944
			_go_fuzz_dep_.CoverTab[92727]++
														nextHashL := hashLen(load6432(src, s+matched), betterLongTableBits, betterLongLen)
														cv := load3232(src, s)
														candidateL := e.longTable[nextHashL]
														coffsetL := candidateL.offset - e.cur - matched
														if coffsetL >= 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:949
				_go_fuzz_dep_.CoverTab[92729]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:949
				return coffsetL < s
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:949
				// _ = "end of CoverTab[92729]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:949
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:949
				_go_fuzz_dep_.CoverTab[92730]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:949
				return s-coffsetL < e.maxMatchOff
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:949
				// _ = "end of CoverTab[92730]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:949
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:949
				_go_fuzz_dep_.CoverTab[92731]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:949
				return cv == load3232(src, coffsetL)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:949
				// _ = "end of CoverTab[92731]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:949
			}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:949
				_go_fuzz_dep_.CoverTab[92732]++

															matchedNext := e.matchlen(s+4, coffsetL+4, src) + 4
															if matchedNext > matched {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:952
					_go_fuzz_dep_.CoverTab[92733]++
																t = coffsetL
																matched = matchedNext
																if debugMatches {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:955
						_go_fuzz_dep_.CoverTab[92734]++
																	println("long match at end-of-match")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:956
						// _ = "end of CoverTab[92734]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:957
						_go_fuzz_dep_.CoverTab[92735]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:957
						// _ = "end of CoverTab[92735]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:957
					}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:957
					// _ = "end of CoverTab[92733]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:958
					_go_fuzz_dep_.CoverTab[92736]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:958
					// _ = "end of CoverTab[92736]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:958
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:958
				// _ = "end of CoverTab[92732]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:959
				_go_fuzz_dep_.CoverTab[92737]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:959
				// _ = "end of CoverTab[92737]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:959
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:959
			// _ = "end of CoverTab[92727]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:959
			_go_fuzz_dep_.CoverTab[92728]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:962
			if true {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:962
				_go_fuzz_dep_.CoverTab[92738]++
															coffsetL = candidateL.prev - e.cur - matched
															if coffsetL >= 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:964
					_go_fuzz_dep_.CoverTab[92739]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:964
					return coffsetL < s
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:964
					// _ = "end of CoverTab[92739]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:964
				}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:964
					_go_fuzz_dep_.CoverTab[92740]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:964
					return s-coffsetL < e.maxMatchOff
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:964
					// _ = "end of CoverTab[92740]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:964
				}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:964
					_go_fuzz_dep_.CoverTab[92741]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:964
					return cv == load3232(src, coffsetL)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:964
					// _ = "end of CoverTab[92741]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:964
				}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:964
					_go_fuzz_dep_.CoverTab[92742]++

																matchedNext := e.matchlen(s+4, coffsetL+4, src) + 4
																if matchedNext > matched {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:967
						_go_fuzz_dep_.CoverTab[92743]++
																	t = coffsetL
																	matched = matchedNext
																	if debugMatches {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:970
							_go_fuzz_dep_.CoverTab[92744]++
																		println("prev long match at end-of-match")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:971
							// _ = "end of CoverTab[92744]"
						} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:972
							_go_fuzz_dep_.CoverTab[92745]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:972
							// _ = "end of CoverTab[92745]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:972
						}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:972
						// _ = "end of CoverTab[92743]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:973
						_go_fuzz_dep_.CoverTab[92746]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:973
						// _ = "end of CoverTab[92746]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:973
					}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:973
					// _ = "end of CoverTab[92742]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:974
					_go_fuzz_dep_.CoverTab[92747]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:974
					// _ = "end of CoverTab[92747]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:974
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:974
				// _ = "end of CoverTab[92738]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:975
				_go_fuzz_dep_.CoverTab[92748]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:975
				// _ = "end of CoverTab[92748]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:975
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:975
			// _ = "end of CoverTab[92728]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:976
			_go_fuzz_dep_.CoverTab[92749]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:976
			// _ = "end of CoverTab[92749]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:976
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:976
		// _ = "end of CoverTab[92574]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:976
		_go_fuzz_dep_.CoverTab[92575]++

													offset2 = offset1
													offset1 = s - t

													if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:981
			_go_fuzz_dep_.CoverTab[92750]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:981
			return s <= t
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:981
			// _ = "end of CoverTab[92750]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:981
		}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:981
			_go_fuzz_dep_.CoverTab[92751]++
														panic(fmt.Sprintf("s (%d) <= t (%d)", s, t))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:982
			// _ = "end of CoverTab[92751]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:983
			_go_fuzz_dep_.CoverTab[92752]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:983
			// _ = "end of CoverTab[92752]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:983
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:983
		// _ = "end of CoverTab[92575]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:983
		_go_fuzz_dep_.CoverTab[92576]++

													if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:985
			_go_fuzz_dep_.CoverTab[92753]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:985
			return canRepeat
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:985
			// _ = "end of CoverTab[92753]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:985
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:985
			_go_fuzz_dep_.CoverTab[92754]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:985
			return int(offset1) > len(src)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:985
			// _ = "end of CoverTab[92754]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:985
		}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:985
			_go_fuzz_dep_.CoverTab[92755]++
														panic("invalid offset")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:986
			// _ = "end of CoverTab[92755]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:987
			_go_fuzz_dep_.CoverTab[92756]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:987
			// _ = "end of CoverTab[92756]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:987
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:987
		// _ = "end of CoverTab[92576]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:987
		_go_fuzz_dep_.CoverTab[92577]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:990
		l := matched

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:993
		tMin := s - e.maxMatchOff
		if tMin < 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:994
			_go_fuzz_dep_.CoverTab[92757]++
														tMin = 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:995
			// _ = "end of CoverTab[92757]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:996
			_go_fuzz_dep_.CoverTab[92758]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:996
			// _ = "end of CoverTab[92758]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:996
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:996
		// _ = "end of CoverTab[92577]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:996
		_go_fuzz_dep_.CoverTab[92578]++
													for t > tMin && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:997
			_go_fuzz_dep_.CoverTab[92759]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:997
			return s > nextEmit
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:997
			// _ = "end of CoverTab[92759]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:997
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:997
			_go_fuzz_dep_.CoverTab[92760]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:997
			return src[t-1] == src[s-1]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:997
			// _ = "end of CoverTab[92760]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:997
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:997
			_go_fuzz_dep_.CoverTab[92761]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:997
			return l < maxMatchLength
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:997
			// _ = "end of CoverTab[92761]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:997
		}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:997
				_go_fuzz_dep_.CoverTab[92762]++
															s--
															t--
															l++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1000
			// _ = "end of CoverTab[92762]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1001
		// _ = "end of CoverTab[92578]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1001
		_go_fuzz_dep_.CoverTab[92579]++

		// Write our sequence
		var seq seq
		seq.litLen = uint32(s - nextEmit)
		seq.matchLen = uint32(l - zstdMinMatch)
		if seq.litLen > 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1007
			_go_fuzz_dep_.CoverTab[92763]++
															blk.literals = append(blk.literals, src[nextEmit:s]...)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1008
			// _ = "end of CoverTab[92763]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1009
			_go_fuzz_dep_.CoverTab[92764]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1009
			// _ = "end of CoverTab[92764]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1009
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1009
		// _ = "end of CoverTab[92579]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1009
		_go_fuzz_dep_.CoverTab[92580]++
														seq.offset = uint32(s-t) + 3
														s += l
														if debugSequences {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1012
			_go_fuzz_dep_.CoverTab[92765]++
															println("sequence", seq, "next s:", s)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1013
			// _ = "end of CoverTab[92765]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1014
			_go_fuzz_dep_.CoverTab[92766]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1014
			// _ = "end of CoverTab[92766]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1014
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1014
		// _ = "end of CoverTab[92580]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1014
		_go_fuzz_dep_.CoverTab[92581]++
														blk.sequences = append(blk.sequences, seq)
														nextEmit = s
														if s >= sLimit {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1017
			_go_fuzz_dep_.CoverTab[92767]++
															break encodeLoop
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1018
			// _ = "end of CoverTab[92767]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1019
			_go_fuzz_dep_.CoverTab[92768]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1019
			// _ = "end of CoverTab[92768]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1019
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1019
		// _ = "end of CoverTab[92581]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1019
		_go_fuzz_dep_.CoverTab[92582]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1022
		index0 := s - l + 1
		for index0 < s-1 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1023
			_go_fuzz_dep_.CoverTab[92769]++
															cv0 := load6432(src, index0)
															cv1 := cv0 >> 8
															h0 := hashLen(cv0, betterLongTableBits, betterLongLen)
															off := index0 + e.cur
															e.longTable[h0] = prevEntry{offset: off, prev: e.longTable[h0].offset}
															e.markLongShardDirty(h0)
															h1 := hashLen(cv1, betterShortTableBits, betterShortLen)
															e.table[h1] = tableEntry{offset: off + 1, val: uint32(cv1)}
															e.markShortShardDirty(h1)
															index0 += 2
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1033
			// _ = "end of CoverTab[92769]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1034
		// _ = "end of CoverTab[92582]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1034
		_go_fuzz_dep_.CoverTab[92583]++

														cv = load6432(src, s)
														if !canRepeat {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1037
			_go_fuzz_dep_.CoverTab[92770]++
															continue
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1038
			// _ = "end of CoverTab[92770]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1039
			_go_fuzz_dep_.CoverTab[92771]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1039
			// _ = "end of CoverTab[92771]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1039
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1039
		// _ = "end of CoverTab[92583]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1039
		_go_fuzz_dep_.CoverTab[92584]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1042
		for {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1042
			_go_fuzz_dep_.CoverTab[92772]++
															o2 := s - offset2
															if load3232(src, o2) != uint32(cv) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1044
				_go_fuzz_dep_.CoverTab[92776]++

																break
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1046
				// _ = "end of CoverTab[92776]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1047
				_go_fuzz_dep_.CoverTab[92777]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1047
				// _ = "end of CoverTab[92777]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1047
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1047
			// _ = "end of CoverTab[92772]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1047
			_go_fuzz_dep_.CoverTab[92773]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1050
			nextHashS := hashLen(cv, betterShortTableBits, betterShortLen)
															nextHashL := hashLen(cv, betterLongTableBits, betterLongLen)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1055
			l := 4 + e.matchlen(s+4, o2+4, src)

															e.longTable[nextHashL] = prevEntry{offset: s + e.cur, prev: e.longTable[nextHashL].offset}
															e.markLongShardDirty(nextHashL)
															e.table[nextHashS] = tableEntry{offset: s + e.cur, val: uint32(cv)}
															e.markShortShardDirty(nextHashS)
															seq.matchLen = uint32(l) - zstdMinMatch
															seq.litLen = 0

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1065
			seq.offset = 1
			s += l
			nextEmit = s
			if debugSequences {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1068
				_go_fuzz_dep_.CoverTab[92778]++
																println("sequence", seq, "next s:", s)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1069
				// _ = "end of CoverTab[92778]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1070
				_go_fuzz_dep_.CoverTab[92779]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1070
				// _ = "end of CoverTab[92779]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1070
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1070
			// _ = "end of CoverTab[92773]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1070
			_go_fuzz_dep_.CoverTab[92774]++
															blk.sequences = append(blk.sequences, seq)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1074
			offset1, offset2 = offset2, offset1
			if s >= sLimit {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1075
				_go_fuzz_dep_.CoverTab[92780]++

																break encodeLoop
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1077
				// _ = "end of CoverTab[92780]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1078
				_go_fuzz_dep_.CoverTab[92781]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1078
				// _ = "end of CoverTab[92781]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1078
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1078
			// _ = "end of CoverTab[92774]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1078
			_go_fuzz_dep_.CoverTab[92775]++
															cv = load6432(src, s)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1079
			// _ = "end of CoverTab[92775]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1080
		// _ = "end of CoverTab[92584]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1081
	// _ = "end of CoverTab[92542]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1081
	_go_fuzz_dep_.CoverTab[92543]++

													if int(nextEmit) < len(src) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1083
		_go_fuzz_dep_.CoverTab[92782]++
														blk.literals = append(blk.literals, src[nextEmit:]...)
														blk.extraLits = len(src) - int(nextEmit)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1085
		// _ = "end of CoverTab[92782]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1086
		_go_fuzz_dep_.CoverTab[92783]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1086
		// _ = "end of CoverTab[92783]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1086
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1086
	// _ = "end of CoverTab[92543]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1086
	_go_fuzz_dep_.CoverTab[92544]++
													blk.recentOffsets[0] = uint32(offset1)
													blk.recentOffsets[1] = uint32(offset2)
													if debugEncoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1089
		_go_fuzz_dep_.CoverTab[92784]++
														println("returning, recent offsets:", blk.recentOffsets, "extra literals:", blk.extraLits)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1090
		// _ = "end of CoverTab[92784]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1091
		_go_fuzz_dep_.CoverTab[92785]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1091
		// _ = "end of CoverTab[92785]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1091
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1091
	// _ = "end of CoverTab[92544]"
}

// ResetDict will reset and set a dictionary if not nil
func (e *betterFastEncoder) Reset(d *dict, singleBlock bool) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1095
	_go_fuzz_dep_.CoverTab[92786]++
													e.resetBase(d, singleBlock)
													if d != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1097
		_go_fuzz_dep_.CoverTab[92787]++
														panic("betterFastEncoder: Reset with dict")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1098
		// _ = "end of CoverTab[92787]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1099
		_go_fuzz_dep_.CoverTab[92788]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1099
		// _ = "end of CoverTab[92788]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1099
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1099
	// _ = "end of CoverTab[92786]"
}

// ResetDict will reset and set a dictionary if not nil
func (e *betterFastEncoderDict) Reset(d *dict, singleBlock bool) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1103
	_go_fuzz_dep_.CoverTab[92789]++
													e.resetBase(d, singleBlock)
													if d == nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1105
		_go_fuzz_dep_.CoverTab[92793]++
														return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1106
		// _ = "end of CoverTab[92793]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1107
		_go_fuzz_dep_.CoverTab[92794]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1107
		// _ = "end of CoverTab[92794]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1107
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1107
	// _ = "end of CoverTab[92789]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1107
	_go_fuzz_dep_.CoverTab[92790]++

													if len(e.dictTable) != len(e.table) || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1109
		_go_fuzz_dep_.CoverTab[92795]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1109
		return d.id != e.lastDictID
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1109
		// _ = "end of CoverTab[92795]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1109
	}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1109
		_go_fuzz_dep_.CoverTab[92796]++
														if len(e.dictTable) != len(e.table) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1110
			_go_fuzz_dep_.CoverTab[92799]++
															e.dictTable = make([]tableEntry, len(e.table))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1111
			// _ = "end of CoverTab[92799]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1112
			_go_fuzz_dep_.CoverTab[92800]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1112
			// _ = "end of CoverTab[92800]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1112
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1112
		// _ = "end of CoverTab[92796]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1112
		_go_fuzz_dep_.CoverTab[92797]++
														end := int32(len(d.content)) - 8 + e.maxMatchOff
														for i := e.maxMatchOff; i < end; i += 4 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1114
			_go_fuzz_dep_.CoverTab[92801]++
															const hashLog = betterShortTableBits

															cv := load6432(d.content, i-e.maxMatchOff)
															nextHash := hashLen(cv, hashLog, betterShortLen)
															nextHash1 := hashLen(cv>>8, hashLog, betterShortLen)
															nextHash2 := hashLen(cv>>16, hashLog, betterShortLen)
															nextHash3 := hashLen(cv>>24, hashLog, betterShortLen)
															e.dictTable[nextHash] = tableEntry{
				val:	uint32(cv),
				offset:	i,
			}
			e.dictTable[nextHash1] = tableEntry{
				val:	uint32(cv >> 8),
				offset:	i + 1,
			}
			e.dictTable[nextHash2] = tableEntry{
				val:	uint32(cv >> 16),
				offset:	i + 2,
			}
			e.dictTable[nextHash3] = tableEntry{
				val:	uint32(cv >> 24),
				offset:	i + 3,
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1137
			// _ = "end of CoverTab[92801]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1138
		// _ = "end of CoverTab[92797]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1138
		_go_fuzz_dep_.CoverTab[92798]++
														e.lastDictID = d.id
														e.allDirty = true
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1140
		// _ = "end of CoverTab[92798]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1141
		_go_fuzz_dep_.CoverTab[92802]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1141
		// _ = "end of CoverTab[92802]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1141
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1141
	// _ = "end of CoverTab[92790]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1141
	_go_fuzz_dep_.CoverTab[92791]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1144
	if len(e.dictLongTable) != len(e.longTable) || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1144
		_go_fuzz_dep_.CoverTab[92803]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1144
		return d.id != e.lastDictID
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1144
		// _ = "end of CoverTab[92803]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1144
	}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1144
		_go_fuzz_dep_.CoverTab[92804]++
														if len(e.dictLongTable) != len(e.longTable) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1145
			_go_fuzz_dep_.CoverTab[92807]++
															e.dictLongTable = make([]prevEntry, len(e.longTable))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1146
			// _ = "end of CoverTab[92807]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1147
			_go_fuzz_dep_.CoverTab[92808]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1147
			// _ = "end of CoverTab[92808]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1147
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1147
		// _ = "end of CoverTab[92804]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1147
		_go_fuzz_dep_.CoverTab[92805]++
														if len(d.content) >= 8 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1148
			_go_fuzz_dep_.CoverTab[92809]++
															cv := load6432(d.content, 0)
															h := hashLen(cv, betterLongTableBits, betterLongLen)
															e.dictLongTable[h] = prevEntry{
				offset:	e.maxMatchOff,
				prev:	e.dictLongTable[h].offset,
			}

			end := int32(len(d.content)) - 8 + e.maxMatchOff
			off := 8
			for i := e.maxMatchOff + 1; i < end; i++ {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1158
				_go_fuzz_dep_.CoverTab[92810]++
																cv = cv>>8 | (uint64(d.content[off]) << 56)
																h := hashLen(cv, betterLongTableBits, betterLongLen)
																e.dictLongTable[h] = prevEntry{
					offset:	i,
					prev:	e.dictLongTable[h].offset,
				}
																off++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1165
				// _ = "end of CoverTab[92810]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1166
			// _ = "end of CoverTab[92809]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1167
			_go_fuzz_dep_.CoverTab[92811]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1167
			// _ = "end of CoverTab[92811]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1167
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1167
		// _ = "end of CoverTab[92805]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1167
		_go_fuzz_dep_.CoverTab[92806]++
														e.lastDictID = d.id
														e.allDirty = true
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1169
		// _ = "end of CoverTab[92806]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1170
		_go_fuzz_dep_.CoverTab[92812]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1170
		// _ = "end of CoverTab[92812]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1170
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1170
	// _ = "end of CoverTab[92791]"

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1173
	{
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1173
		_go_fuzz_dep_.CoverTab[92813]++
														dirtyShardCnt := 0
														if !e.allDirty {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1175
			_go_fuzz_dep_.CoverTab[92815]++
															for i := range e.shortTableShardDirty {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1176
				_go_fuzz_dep_.CoverTab[92816]++
																if e.shortTableShardDirty[i] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1177
					_go_fuzz_dep_.CoverTab[92817]++
																	dirtyShardCnt++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1178
					// _ = "end of CoverTab[92817]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1179
					_go_fuzz_dep_.CoverTab[92818]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1179
					// _ = "end of CoverTab[92818]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1179
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1179
				// _ = "end of CoverTab[92816]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1180
			// _ = "end of CoverTab[92815]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1181
			_go_fuzz_dep_.CoverTab[92819]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1181
			// _ = "end of CoverTab[92819]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1181
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1181
		// _ = "end of CoverTab[92813]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1181
		_go_fuzz_dep_.CoverTab[92814]++
														const shardCnt = betterShortTableShardCnt
														const shardSize = betterShortTableShardSize
														if e.allDirty || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1184
			_go_fuzz_dep_.CoverTab[92820]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1184
			return dirtyShardCnt > shardCnt*4/6
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1184
			// _ = "end of CoverTab[92820]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1184
		}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1184
			_go_fuzz_dep_.CoverTab[92821]++
															copy(e.table[:], e.dictTable)
															for i := range e.shortTableShardDirty {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1186
				_go_fuzz_dep_.CoverTab[92822]++
																e.shortTableShardDirty[i] = false
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1187
				// _ = "end of CoverTab[92822]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1188
			// _ = "end of CoverTab[92821]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1189
			_go_fuzz_dep_.CoverTab[92823]++
															for i := range e.shortTableShardDirty {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1190
				_go_fuzz_dep_.CoverTab[92824]++
																if !e.shortTableShardDirty[i] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1191
					_go_fuzz_dep_.CoverTab[92826]++
																	continue
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1192
					// _ = "end of CoverTab[92826]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1193
					_go_fuzz_dep_.CoverTab[92827]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1193
					// _ = "end of CoverTab[92827]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1193
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1193
				// _ = "end of CoverTab[92824]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1193
				_go_fuzz_dep_.CoverTab[92825]++

																copy(e.table[i*shardSize:(i+1)*shardSize], e.dictTable[i*shardSize:(i+1)*shardSize])
																e.shortTableShardDirty[i] = false
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1196
				// _ = "end of CoverTab[92825]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1197
			// _ = "end of CoverTab[92823]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1198
		// _ = "end of CoverTab[92814]"
	}
	{
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1200
		_go_fuzz_dep_.CoverTab[92828]++
														dirtyShardCnt := 0
														if !e.allDirty {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1202
			_go_fuzz_dep_.CoverTab[92830]++
															for i := range e.shortTableShardDirty {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1203
				_go_fuzz_dep_.CoverTab[92831]++
																if e.shortTableShardDirty[i] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1204
					_go_fuzz_dep_.CoverTab[92832]++
																	dirtyShardCnt++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1205
					// _ = "end of CoverTab[92832]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1206
					_go_fuzz_dep_.CoverTab[92833]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1206
					// _ = "end of CoverTab[92833]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1206
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1206
				// _ = "end of CoverTab[92831]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1207
			// _ = "end of CoverTab[92830]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1208
			_go_fuzz_dep_.CoverTab[92834]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1208
			// _ = "end of CoverTab[92834]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1208
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1208
		// _ = "end of CoverTab[92828]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1208
		_go_fuzz_dep_.CoverTab[92829]++
														const shardCnt = betterLongTableShardCnt
														const shardSize = betterLongTableShardSize
														if e.allDirty || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1211
			_go_fuzz_dep_.CoverTab[92835]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1211
			return dirtyShardCnt > shardCnt*4/6
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1211
			// _ = "end of CoverTab[92835]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1211
		}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1211
			_go_fuzz_dep_.CoverTab[92836]++
															copy(e.longTable[:], e.dictLongTable)
															for i := range e.longTableShardDirty {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1213
				_go_fuzz_dep_.CoverTab[92837]++
																e.longTableShardDirty[i] = false
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1214
				// _ = "end of CoverTab[92837]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1215
			// _ = "end of CoverTab[92836]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1216
			_go_fuzz_dep_.CoverTab[92838]++
															for i := range e.longTableShardDirty {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1217
				_go_fuzz_dep_.CoverTab[92839]++
																if !e.longTableShardDirty[i] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1218
					_go_fuzz_dep_.CoverTab[92841]++
																	continue
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1219
					// _ = "end of CoverTab[92841]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1220
					_go_fuzz_dep_.CoverTab[92842]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1220
					// _ = "end of CoverTab[92842]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1220
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1220
				// _ = "end of CoverTab[92839]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1220
				_go_fuzz_dep_.CoverTab[92840]++

																copy(e.longTable[i*shardSize:(i+1)*shardSize], e.dictLongTable[i*shardSize:(i+1)*shardSize])
																e.longTableShardDirty[i] = false
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1223
				// _ = "end of CoverTab[92840]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1224
			// _ = "end of CoverTab[92838]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1225
		// _ = "end of CoverTab[92829]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1226
	_go_fuzz_dep_.CoverTab[92792]++
													e.cur = e.maxMatchOff
													e.allDirty = false
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1228
	// _ = "end of CoverTab[92792]"
}

func (e *betterFastEncoderDict) markLongShardDirty(entryNum uint32) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1231
	_go_fuzz_dep_.CoverTab[92843]++
													e.longTableShardDirty[entryNum/betterLongTableShardSize] = true
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1232
	// _ = "end of CoverTab[92843]"
}

func (e *betterFastEncoderDict) markShortShardDirty(entryNum uint32) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1235
	_go_fuzz_dep_.CoverTab[92844]++
													e.shortTableShardDirty[entryNum/betterShortTableShardSize] = true
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1236
	// _ = "end of CoverTab[92844]"
}

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1237
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_better.go:1237
var _ = _go_fuzz_dep_.CoverTab
