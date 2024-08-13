// Copyright 2019+ Klaus Post. All rights reserved.
// License information can be found in the LICENSE file.
// Based on work by Yann Collet, released under BSD License.

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:5
package zstd

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:5
)

import "fmt"

const (
	dFastLongTableBits	= 17				// Bits used in the long match table
	dFastLongTableSize	= 1 << dFastLongTableBits	// Size of the table
	dFastLongTableMask	= dFastLongTableSize - 1	// Mask for table indices. Redundant, but can eliminate bounds checks.
	dFastLongLen		= 8				// Bytes used for table hash

	dLongTableShardCnt	= 1 << (dFastLongTableBits - dictShardBits)	// Number of shards in the table
	dLongTableShardSize	= dFastLongTableSize / tableShardCnt		// Size of an individual shard

	dFastShortTableBits	= tableBits			// Bits used in the short match table
	dFastShortTableSize	= 1 << dFastShortTableBits	// Size of the table
	dFastShortTableMask	= dFastShortTableSize - 1	// Mask for table indices. Redundant, but can eliminate bounds checks.
	dFastShortLen		= 5				// Bytes used for table hash

)

type doubleFastEncoder struct {
	fastEncoder
	longTable	[dFastLongTableSize]tableEntry
}

type doubleFastEncoderDict struct {
	fastEncoderDict
	longTable		[dFastLongTableSize]tableEntry
	dictLongTable		[]tableEntry
	longTableShardDirty	[dLongTableShardCnt]bool
}

// Encode mimmics functionality in zstd_dfast.c
func (e *doubleFastEncoder) Encode(blk *blockEnc, src []byte) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:38
	_go_fuzz_dep_.CoverTab[92845]++
												const (
		// Input margin is the number of bytes we read (8)
		// and the maximum we will read ahead (2)
		inputMargin		= 8 + 2
		minNonLiteralBlockSize	= 16
	)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:47
	for e.cur >= bufferReset {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:47
		_go_fuzz_dep_.CoverTab[92852]++
													if len(e.hist) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:48
			_go_fuzz_dep_.CoverTab[92856]++
														for i := range e.table[:] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:49
				_go_fuzz_dep_.CoverTab[92859]++
															e.table[i] = tableEntry{}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:50
				// _ = "end of CoverTab[92859]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:51
			// _ = "end of CoverTab[92856]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:51
			_go_fuzz_dep_.CoverTab[92857]++
														for i := range e.longTable[:] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:52
				_go_fuzz_dep_.CoverTab[92860]++
															e.longTable[i] = tableEntry{}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:53
				// _ = "end of CoverTab[92860]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:54
			// _ = "end of CoverTab[92857]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:54
			_go_fuzz_dep_.CoverTab[92858]++
														e.cur = e.maxMatchOff
														break
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:56
			// _ = "end of CoverTab[92858]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:57
			_go_fuzz_dep_.CoverTab[92861]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:57
			// _ = "end of CoverTab[92861]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:57
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:57
		// _ = "end of CoverTab[92852]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:57
		_go_fuzz_dep_.CoverTab[92853]++

													minOff := e.cur + int32(len(e.hist)) - e.maxMatchOff
													for i := range e.table[:] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:60
			_go_fuzz_dep_.CoverTab[92862]++
														v := e.table[i].offset
														if v < minOff {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:62
				_go_fuzz_dep_.CoverTab[92864]++
															v = 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:63
				// _ = "end of CoverTab[92864]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:64
				_go_fuzz_dep_.CoverTab[92865]++
															v = v - e.cur + e.maxMatchOff
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:65
				// _ = "end of CoverTab[92865]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:66
			// _ = "end of CoverTab[92862]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:66
			_go_fuzz_dep_.CoverTab[92863]++
														e.table[i].offset = v
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:67
			// _ = "end of CoverTab[92863]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:68
		// _ = "end of CoverTab[92853]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:68
		_go_fuzz_dep_.CoverTab[92854]++
													for i := range e.longTable[:] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:69
			_go_fuzz_dep_.CoverTab[92866]++
														v := e.longTable[i].offset
														if v < minOff {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:71
				_go_fuzz_dep_.CoverTab[92868]++
															v = 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:72
				// _ = "end of CoverTab[92868]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:73
				_go_fuzz_dep_.CoverTab[92869]++
															v = v - e.cur + e.maxMatchOff
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:74
				// _ = "end of CoverTab[92869]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:75
			// _ = "end of CoverTab[92866]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:75
			_go_fuzz_dep_.CoverTab[92867]++
														e.longTable[i].offset = v
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:76
			// _ = "end of CoverTab[92867]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:77
		// _ = "end of CoverTab[92854]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:77
		_go_fuzz_dep_.CoverTab[92855]++
													e.cur = e.maxMatchOff
													break
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:79
		// _ = "end of CoverTab[92855]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:80
	// _ = "end of CoverTab[92845]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:80
	_go_fuzz_dep_.CoverTab[92846]++

												s := e.addBlock(src)
												blk.size = len(src)
												if len(src) < minNonLiteralBlockSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:84
		_go_fuzz_dep_.CoverTab[92870]++
													blk.extraLits = len(src)
													blk.literals = blk.literals[:len(src)]
													copy(blk.literals, src)
													return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:88
		// _ = "end of CoverTab[92870]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:89
		_go_fuzz_dep_.CoverTab[92871]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:89
		// _ = "end of CoverTab[92871]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:89
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:89
	// _ = "end of CoverTab[92846]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:89
	_go_fuzz_dep_.CoverTab[92847]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:92
	src = e.hist
	sLimit := int32(len(src)) - inputMargin
	// stepSize is the number of bytes to skip on every main loop iteration.
												// It should be >= 1.
												const stepSize = 1

												const kSearchStrength = 8

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:101
	nextEmit := s
												cv := load6432(src, s)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:105
	offset1 := int32(blk.recentOffsets[0])
	offset2 := int32(blk.recentOffsets[1])

	addLiterals := func(s *seq, until int32) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:108
		_go_fuzz_dep_.CoverTab[92872]++
													if until == nextEmit {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:109
			_go_fuzz_dep_.CoverTab[92874]++
														return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:110
			// _ = "end of CoverTab[92874]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:111
			_go_fuzz_dep_.CoverTab[92875]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:111
			// _ = "end of CoverTab[92875]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:111
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:111
		// _ = "end of CoverTab[92872]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:111
		_go_fuzz_dep_.CoverTab[92873]++
													blk.literals = append(blk.literals, src[nextEmit:until]...)
													s.litLen = uint32(until - nextEmit)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:113
		// _ = "end of CoverTab[92873]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:114
	// _ = "end of CoverTab[92847]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:114
	_go_fuzz_dep_.CoverTab[92848]++
												if debugEncoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:115
		_go_fuzz_dep_.CoverTab[92876]++
													println("recent offsets:", blk.recentOffsets)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:116
		// _ = "end of CoverTab[92876]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:117
		_go_fuzz_dep_.CoverTab[92877]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:117
		// _ = "end of CoverTab[92877]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:117
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:117
	// _ = "end of CoverTab[92848]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:117
	_go_fuzz_dep_.CoverTab[92849]++

encodeLoop:
	for {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:120
		_go_fuzz_dep_.CoverTab[92878]++
													var t int32

													canRepeat := len(blk.sequences) > 2

													for {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:125
			_go_fuzz_dep_.CoverTab[92888]++
														if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:126
				_go_fuzz_dep_.CoverTab[92894]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:126
				return canRepeat
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:126
				// _ = "end of CoverTab[92894]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:126
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:126
				_go_fuzz_dep_.CoverTab[92895]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:126
				return offset1 == 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:126
				// _ = "end of CoverTab[92895]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:126
			}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:126
				_go_fuzz_dep_.CoverTab[92896]++
															panic("offset0 was 0")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:127
				// _ = "end of CoverTab[92896]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:128
				_go_fuzz_dep_.CoverTab[92897]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:128
				// _ = "end of CoverTab[92897]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:128
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:128
			// _ = "end of CoverTab[92888]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:128
			_go_fuzz_dep_.CoverTab[92889]++

														nextHashS := hashLen(cv, dFastShortTableBits, dFastShortLen)
														nextHashL := hashLen(cv, dFastLongTableBits, dFastLongLen)
														candidateL := e.longTable[nextHashL]
														candidateS := e.table[nextHashS]

														const repOff = 1
														repIndex := s - offset1 + repOff
														entry := tableEntry{offset: s + e.cur, val: uint32(cv)}
														e.longTable[nextHashL] = entry
														e.table[nextHashS] = entry

														if canRepeat {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:141
				_go_fuzz_dep_.CoverTab[92898]++
															if repIndex >= 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:142
					_go_fuzz_dep_.CoverTab[92899]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:142
					return load3232(src, repIndex) == uint32(cv>>(repOff*8))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:142
					// _ = "end of CoverTab[92899]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:142
				}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:142
					_go_fuzz_dep_.CoverTab[92900]++
																// Consider history as well.
																var seq seq
																lenght := 4 + e.matchlen(s+4+repOff, repIndex+4, src)

																seq.matchLen = uint32(lenght - zstdMinMatch)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:151
					start := s + repOff

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:154
					startLimit := nextEmit + 1

					tMin := s - e.maxMatchOff
					if tMin < 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:157
						_go_fuzz_dep_.CoverTab[92905]++
																	tMin = 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:158
						// _ = "end of CoverTab[92905]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:159
						_go_fuzz_dep_.CoverTab[92906]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:159
						// _ = "end of CoverTab[92906]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:159
					}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:159
					// _ = "end of CoverTab[92900]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:159
					_go_fuzz_dep_.CoverTab[92901]++
																for repIndex > tMin && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:160
						_go_fuzz_dep_.CoverTab[92907]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:160
						return start > startLimit
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:160
						// _ = "end of CoverTab[92907]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:160
					}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:160
						_go_fuzz_dep_.CoverTab[92908]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:160
						return src[repIndex-1] == src[start-1]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:160
						// _ = "end of CoverTab[92908]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:160
					}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:160
						_go_fuzz_dep_.CoverTab[92909]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:160
						return seq.matchLen < maxMatchLength-zstdMinMatch-1
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:160
						// _ = "end of CoverTab[92909]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:160
					}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:160
						_go_fuzz_dep_.CoverTab[92910]++
																	repIndex--
																	start--
																	seq.matchLen++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:163
						// _ = "end of CoverTab[92910]"
					}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:164
					// _ = "end of CoverTab[92901]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:164
					_go_fuzz_dep_.CoverTab[92902]++
																addLiterals(&seq, start)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:168
					seq.offset = 1
					if debugSequences {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:169
						_go_fuzz_dep_.CoverTab[92911]++
																	println("repeat sequence", seq, "next s:", s)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:170
						// _ = "end of CoverTab[92911]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:171
						_go_fuzz_dep_.CoverTab[92912]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:171
						// _ = "end of CoverTab[92912]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:171
					}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:171
					// _ = "end of CoverTab[92902]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:171
					_go_fuzz_dep_.CoverTab[92903]++
																blk.sequences = append(blk.sequences, seq)
																s += lenght + repOff
																nextEmit = s
																if s >= sLimit {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:175
						_go_fuzz_dep_.CoverTab[92913]++
																	if debugEncoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:176
							_go_fuzz_dep_.CoverTab[92915]++
																		println("repeat ended", s, lenght)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:177
							// _ = "end of CoverTab[92915]"

						} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:179
							_go_fuzz_dep_.CoverTab[92916]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:179
							// _ = "end of CoverTab[92916]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:179
						}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:179
						// _ = "end of CoverTab[92913]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:179
						_go_fuzz_dep_.CoverTab[92914]++
																	break encodeLoop
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:180
						// _ = "end of CoverTab[92914]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:181
						_go_fuzz_dep_.CoverTab[92917]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:181
						// _ = "end of CoverTab[92917]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:181
					}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:181
					// _ = "end of CoverTab[92903]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:181
					_go_fuzz_dep_.CoverTab[92904]++
																cv = load6432(src, s)
																continue
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:183
					// _ = "end of CoverTab[92904]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:184
					_go_fuzz_dep_.CoverTab[92918]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:184
					// _ = "end of CoverTab[92918]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:184
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:184
				// _ = "end of CoverTab[92898]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:185
				_go_fuzz_dep_.CoverTab[92919]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:185
				// _ = "end of CoverTab[92919]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:185
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:185
			// _ = "end of CoverTab[92889]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:185
			_go_fuzz_dep_.CoverTab[92890]++

														coffsetL := s - (candidateL.offset - e.cur)
														coffsetS := s - (candidateS.offset - e.cur)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:191
			if coffsetL < e.maxMatchOff && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:191
				_go_fuzz_dep_.CoverTab[92920]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:191
				return uint32(cv) == candidateL.val
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:191
				// _ = "end of CoverTab[92920]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:191
			}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:191
				_go_fuzz_dep_.CoverTab[92921]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:195
				t = candidateL.offset - e.cur
				if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:196
					_go_fuzz_dep_.CoverTab[92925]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:196
					return s <= t
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:196
					// _ = "end of CoverTab[92925]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:196
				}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:196
					_go_fuzz_dep_.CoverTab[92926]++
																panic(fmt.Sprintf("s (%d) <= t (%d)", s, t))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:197
					// _ = "end of CoverTab[92926]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:198
					_go_fuzz_dep_.CoverTab[92927]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:198
					// _ = "end of CoverTab[92927]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:198
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:198
				// _ = "end of CoverTab[92921]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:198
				_go_fuzz_dep_.CoverTab[92922]++
															if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:199
					_go_fuzz_dep_.CoverTab[92928]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:199
					return s-t > e.maxMatchOff
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:199
					// _ = "end of CoverTab[92928]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:199
				}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:199
					_go_fuzz_dep_.CoverTab[92929]++
																panic("s - t >e.maxMatchOff")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:200
					// _ = "end of CoverTab[92929]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:201
					_go_fuzz_dep_.CoverTab[92930]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:201
					// _ = "end of CoverTab[92930]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:201
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:201
				// _ = "end of CoverTab[92922]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:201
				_go_fuzz_dep_.CoverTab[92923]++
															if debugMatches {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:202
					_go_fuzz_dep_.CoverTab[92931]++
																println("long match")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:203
					// _ = "end of CoverTab[92931]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:204
					_go_fuzz_dep_.CoverTab[92932]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:204
					// _ = "end of CoverTab[92932]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:204
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:204
				// _ = "end of CoverTab[92923]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:204
				_go_fuzz_dep_.CoverTab[92924]++
															break
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:205
				// _ = "end of CoverTab[92924]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:206
				_go_fuzz_dep_.CoverTab[92933]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:206
				// _ = "end of CoverTab[92933]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:206
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:206
			// _ = "end of CoverTab[92890]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:206
			_go_fuzz_dep_.CoverTab[92891]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:209
			if coffsetS < e.maxMatchOff && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:209
				_go_fuzz_dep_.CoverTab[92934]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:209
				return uint32(cv) == candidateS.val
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:209
				// _ = "end of CoverTab[92934]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:209
			}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:209
				_go_fuzz_dep_.CoverTab[92935]++
				// found a regular match
															// See if we can find a long match at s+1
															const checkAt = 1
															cv := load6432(src, s+checkAt)
															nextHashL = hashLen(cv, dFastLongTableBits, dFastLongLen)
															candidateL = e.longTable[nextHashL]
															coffsetL = s - (candidateL.offset - e.cur) + checkAt

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:219
				e.longTable[nextHashL] = tableEntry{offset: s + checkAt + e.cur, val: uint32(cv)}
				if coffsetL < e.maxMatchOff && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:220
					_go_fuzz_dep_.CoverTab[92941]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:220
					return uint32(cv) == candidateL.val
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:220
					// _ = "end of CoverTab[92941]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:220
				}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:220
					_go_fuzz_dep_.CoverTab[92942]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:224
					t = candidateL.offset - e.cur
					s += checkAt
					if debugMatches {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:226
						_go_fuzz_dep_.CoverTab[92944]++
																	println("long match (after short)")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:227
						// _ = "end of CoverTab[92944]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:228
						_go_fuzz_dep_.CoverTab[92945]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:228
						// _ = "end of CoverTab[92945]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:228
					}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:228
					// _ = "end of CoverTab[92942]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:228
					_go_fuzz_dep_.CoverTab[92943]++
																break
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:229
					// _ = "end of CoverTab[92943]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:230
					_go_fuzz_dep_.CoverTab[92946]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:230
					// _ = "end of CoverTab[92946]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:230
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:230
				// _ = "end of CoverTab[92935]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:230
				_go_fuzz_dep_.CoverTab[92936]++

															t = candidateS.offset - e.cur
															if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:233
					_go_fuzz_dep_.CoverTab[92947]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:233
					return s <= t
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:233
					// _ = "end of CoverTab[92947]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:233
				}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:233
					_go_fuzz_dep_.CoverTab[92948]++
																panic(fmt.Sprintf("s (%d) <= t (%d)", s, t))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:234
					// _ = "end of CoverTab[92948]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:235
					_go_fuzz_dep_.CoverTab[92949]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:235
					// _ = "end of CoverTab[92949]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:235
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:235
				// _ = "end of CoverTab[92936]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:235
				_go_fuzz_dep_.CoverTab[92937]++
															if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:236
					_go_fuzz_dep_.CoverTab[92950]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:236
					return s-t > e.maxMatchOff
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:236
					// _ = "end of CoverTab[92950]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:236
				}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:236
					_go_fuzz_dep_.CoverTab[92951]++
																panic("s - t >e.maxMatchOff")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:237
					// _ = "end of CoverTab[92951]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:238
					_go_fuzz_dep_.CoverTab[92952]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:238
					// _ = "end of CoverTab[92952]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:238
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:238
				// _ = "end of CoverTab[92937]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:238
				_go_fuzz_dep_.CoverTab[92938]++
															if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:239
					_go_fuzz_dep_.CoverTab[92953]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:239
					return t < 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:239
					// _ = "end of CoverTab[92953]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:239
				}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:239
					_go_fuzz_dep_.CoverTab[92954]++
																panic("t<0")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:240
					// _ = "end of CoverTab[92954]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:241
					_go_fuzz_dep_.CoverTab[92955]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:241
					// _ = "end of CoverTab[92955]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:241
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:241
				// _ = "end of CoverTab[92938]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:241
				_go_fuzz_dep_.CoverTab[92939]++
															if debugMatches {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:242
					_go_fuzz_dep_.CoverTab[92956]++
																println("short match")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:243
					// _ = "end of CoverTab[92956]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:244
					_go_fuzz_dep_.CoverTab[92957]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:244
					// _ = "end of CoverTab[92957]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:244
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:244
				// _ = "end of CoverTab[92939]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:244
				_go_fuzz_dep_.CoverTab[92940]++
															break
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:245
				// _ = "end of CoverTab[92940]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:246
				_go_fuzz_dep_.CoverTab[92958]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:246
				// _ = "end of CoverTab[92958]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:246
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:246
			// _ = "end of CoverTab[92891]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:246
			_go_fuzz_dep_.CoverTab[92892]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:249
			s += stepSize + ((s - nextEmit) >> (kSearchStrength - 1))
			if s >= sLimit {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:250
				_go_fuzz_dep_.CoverTab[92959]++
															break encodeLoop
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:251
				// _ = "end of CoverTab[92959]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:252
				_go_fuzz_dep_.CoverTab[92960]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:252
				// _ = "end of CoverTab[92960]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:252
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:252
			// _ = "end of CoverTab[92892]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:252
			_go_fuzz_dep_.CoverTab[92893]++
														cv = load6432(src, s)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:253
			// _ = "end of CoverTab[92893]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:254
		// _ = "end of CoverTab[92878]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:254
		_go_fuzz_dep_.CoverTab[92879]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:258
		offset2 = offset1
		offset1 = s - t

		if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:261
			_go_fuzz_dep_.CoverTab[92961]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:261
			return s <= t
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:261
			// _ = "end of CoverTab[92961]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:261
		}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:261
			_go_fuzz_dep_.CoverTab[92962]++
														panic(fmt.Sprintf("s (%d) <= t (%d)", s, t))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:262
			// _ = "end of CoverTab[92962]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:263
			_go_fuzz_dep_.CoverTab[92963]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:263
			// _ = "end of CoverTab[92963]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:263
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:263
		// _ = "end of CoverTab[92879]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:263
		_go_fuzz_dep_.CoverTab[92880]++

													if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:265
			_go_fuzz_dep_.CoverTab[92964]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:265
			return canRepeat
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:265
			// _ = "end of CoverTab[92964]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:265
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:265
			_go_fuzz_dep_.CoverTab[92965]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:265
			return int(offset1) > len(src)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:265
			// _ = "end of CoverTab[92965]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:265
		}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:265
			_go_fuzz_dep_.CoverTab[92966]++
														panic("invalid offset")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:266
			// _ = "end of CoverTab[92966]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:267
			_go_fuzz_dep_.CoverTab[92967]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:267
			// _ = "end of CoverTab[92967]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:267
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:267
		// _ = "end of CoverTab[92880]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:267
		_go_fuzz_dep_.CoverTab[92881]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:270
		l := e.matchlen(s+4, t+4, src) + 4

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:273
		tMin := s - e.maxMatchOff
		if tMin < 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:274
			_go_fuzz_dep_.CoverTab[92968]++
														tMin = 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:275
			// _ = "end of CoverTab[92968]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:276
			_go_fuzz_dep_.CoverTab[92969]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:276
			// _ = "end of CoverTab[92969]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:276
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:276
		// _ = "end of CoverTab[92881]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:276
		_go_fuzz_dep_.CoverTab[92882]++
													for t > tMin && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:277
			_go_fuzz_dep_.CoverTab[92970]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:277
			return s > nextEmit
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:277
			// _ = "end of CoverTab[92970]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:277
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:277
			_go_fuzz_dep_.CoverTab[92971]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:277
			return src[t-1] == src[s-1]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:277
			// _ = "end of CoverTab[92971]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:277
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:277
			_go_fuzz_dep_.CoverTab[92972]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:277
			return l < maxMatchLength
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:277
			// _ = "end of CoverTab[92972]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:277
		}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:277
			_go_fuzz_dep_.CoverTab[92973]++
														s--
														t--
														l++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:280
			// _ = "end of CoverTab[92973]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:281
		// _ = "end of CoverTab[92882]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:281
		_go_fuzz_dep_.CoverTab[92883]++

		// Write our sequence
		var seq seq
		seq.litLen = uint32(s - nextEmit)
		seq.matchLen = uint32(l - zstdMinMatch)
		if seq.litLen > 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:287
			_go_fuzz_dep_.CoverTab[92974]++
														blk.literals = append(blk.literals, src[nextEmit:s]...)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:288
			// _ = "end of CoverTab[92974]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:289
			_go_fuzz_dep_.CoverTab[92975]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:289
			// _ = "end of CoverTab[92975]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:289
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:289
		// _ = "end of CoverTab[92883]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:289
		_go_fuzz_dep_.CoverTab[92884]++
													seq.offset = uint32(s-t) + 3
													s += l
													if debugSequences {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:292
			_go_fuzz_dep_.CoverTab[92976]++
														println("sequence", seq, "next s:", s)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:293
			// _ = "end of CoverTab[92976]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:294
			_go_fuzz_dep_.CoverTab[92977]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:294
			// _ = "end of CoverTab[92977]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:294
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:294
		// _ = "end of CoverTab[92884]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:294
		_go_fuzz_dep_.CoverTab[92885]++
													blk.sequences = append(blk.sequences, seq)
													nextEmit = s
													if s >= sLimit {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:297
			_go_fuzz_dep_.CoverTab[92978]++
														break encodeLoop
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:298
			// _ = "end of CoverTab[92978]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:299
			_go_fuzz_dep_.CoverTab[92979]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:299
			// _ = "end of CoverTab[92979]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:299
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:299
		// _ = "end of CoverTab[92885]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:299
		_go_fuzz_dep_.CoverTab[92886]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:302
		index0 := s - l + 1

		index1 := s - 2

		cv0 := load6432(src, index0)
		cv1 := load6432(src, index1)
		te0 := tableEntry{offset: index0 + e.cur, val: uint32(cv0)}
		te1 := tableEntry{offset: index1 + e.cur, val: uint32(cv1)}
		e.longTable[hashLen(cv0, dFastLongTableBits, dFastLongLen)] = te0
		e.longTable[hashLen(cv1, dFastLongTableBits, dFastLongLen)] = te1
		cv0 >>= 8
		cv1 >>= 8
		te0.offset++
		te1.offset++
		te0.val = uint32(cv0)
		te1.val = uint32(cv1)
		e.table[hashLen(cv0, dFastShortTableBits, dFastShortLen)] = te0
		e.table[hashLen(cv1, dFastShortTableBits, dFastShortLen)] = te1

		cv = load6432(src, s)

		if !canRepeat {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:323
			_go_fuzz_dep_.CoverTab[92980]++
														continue
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:324
			// _ = "end of CoverTab[92980]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:325
			_go_fuzz_dep_.CoverTab[92981]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:325
			// _ = "end of CoverTab[92981]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:325
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:325
		// _ = "end of CoverTab[92886]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:325
		_go_fuzz_dep_.CoverTab[92887]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:328
		for {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:328
			_go_fuzz_dep_.CoverTab[92982]++
														o2 := s - offset2
														if load3232(src, o2) != uint32(cv) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:330
				_go_fuzz_dep_.CoverTab[92986]++

															break
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:332
				// _ = "end of CoverTab[92986]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:333
				_go_fuzz_dep_.CoverTab[92987]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:333
				// _ = "end of CoverTab[92987]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:333
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:333
			// _ = "end of CoverTab[92982]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:333
			_go_fuzz_dep_.CoverTab[92983]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:336
			nextHashS := hashLen(cv, dFastShortTableBits, dFastShortLen)
														nextHashL := hashLen(cv, dFastLongTableBits, dFastLongLen)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:341
			l := 4 + e.matchlen(s+4, o2+4, src)

														entry := tableEntry{offset: s + e.cur, val: uint32(cv)}
														e.longTable[nextHashL] = entry
														e.table[nextHashS] = entry
														seq.matchLen = uint32(l) - zstdMinMatch
														seq.litLen = 0

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:350
			seq.offset = 1
			s += l
			nextEmit = s
			if debugSequences {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:353
				_go_fuzz_dep_.CoverTab[92988]++
															println("sequence", seq, "next s:", s)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:354
				// _ = "end of CoverTab[92988]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:355
				_go_fuzz_dep_.CoverTab[92989]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:355
				// _ = "end of CoverTab[92989]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:355
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:355
			// _ = "end of CoverTab[92983]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:355
			_go_fuzz_dep_.CoverTab[92984]++
														blk.sequences = append(blk.sequences, seq)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:359
			offset1, offset2 = offset2, offset1
			if s >= sLimit {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:360
				_go_fuzz_dep_.CoverTab[92990]++

															break encodeLoop
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:362
				// _ = "end of CoverTab[92990]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:363
				_go_fuzz_dep_.CoverTab[92991]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:363
				// _ = "end of CoverTab[92991]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:363
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:363
			// _ = "end of CoverTab[92984]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:363
			_go_fuzz_dep_.CoverTab[92985]++
														cv = load6432(src, s)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:364
			// _ = "end of CoverTab[92985]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:365
		// _ = "end of CoverTab[92887]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:366
	// _ = "end of CoverTab[92849]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:366
	_go_fuzz_dep_.CoverTab[92850]++

												if int(nextEmit) < len(src) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:368
		_go_fuzz_dep_.CoverTab[92992]++
													blk.literals = append(blk.literals, src[nextEmit:]...)
													blk.extraLits = len(src) - int(nextEmit)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:370
		// _ = "end of CoverTab[92992]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:371
		_go_fuzz_dep_.CoverTab[92993]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:371
		// _ = "end of CoverTab[92993]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:371
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:371
	// _ = "end of CoverTab[92850]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:371
	_go_fuzz_dep_.CoverTab[92851]++
												blk.recentOffsets[0] = uint32(offset1)
												blk.recentOffsets[1] = uint32(offset2)
												if debugEncoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:374
		_go_fuzz_dep_.CoverTab[92994]++
													println("returning, recent offsets:", blk.recentOffsets, "extra literals:", blk.extraLits)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:375
		// _ = "end of CoverTab[92994]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:376
		_go_fuzz_dep_.CoverTab[92995]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:376
		// _ = "end of CoverTab[92995]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:376
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:376
	// _ = "end of CoverTab[92851]"
}

// EncodeNoHist will encode a block with no history and no following blocks.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:379
// Most notable difference is that src will not be copied for history and
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:379
// we do not need to check for max match length.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:382
func (e *doubleFastEncoder) EncodeNoHist(blk *blockEnc, src []byte) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:382
	_go_fuzz_dep_.CoverTab[92996]++
												const (
		// Input margin is the number of bytes we read (8)
		// and the maximum we will read ahead (2)
		inputMargin		= 8 + 2
		minNonLiteralBlockSize	= 16
	)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:391
	if e.cur >= bufferReset {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:391
		_go_fuzz_dep_.CoverTab[93004]++
													for i := range e.table[:] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:392
			_go_fuzz_dep_.CoverTab[93007]++
														e.table[i] = tableEntry{}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:393
			// _ = "end of CoverTab[93007]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:394
		// _ = "end of CoverTab[93004]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:394
		_go_fuzz_dep_.CoverTab[93005]++
													for i := range e.longTable[:] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:395
			_go_fuzz_dep_.CoverTab[93008]++
														e.longTable[i] = tableEntry{}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:396
			// _ = "end of CoverTab[93008]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:397
		// _ = "end of CoverTab[93005]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:397
		_go_fuzz_dep_.CoverTab[93006]++
													e.cur = e.maxMatchOff
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:398
		// _ = "end of CoverTab[93006]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:399
		_go_fuzz_dep_.CoverTab[93009]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:399
		// _ = "end of CoverTab[93009]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:399
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:399
	// _ = "end of CoverTab[92996]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:399
	_go_fuzz_dep_.CoverTab[92997]++

												s := int32(0)
												blk.size = len(src)
												if len(src) < minNonLiteralBlockSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:403
		_go_fuzz_dep_.CoverTab[93010]++
													blk.extraLits = len(src)
													blk.literals = blk.literals[:len(src)]
													copy(blk.literals, src)
													return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:407
		// _ = "end of CoverTab[93010]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:408
		_go_fuzz_dep_.CoverTab[93011]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:408
		// _ = "end of CoverTab[93011]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:408
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:408
	// _ = "end of CoverTab[92997]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:408
	_go_fuzz_dep_.CoverTab[92998]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:411
	sLimit := int32(len(src)) - inputMargin
	// stepSize is the number of bytes to skip on every main loop iteration.
												// It should be >= 1.
												const stepSize = 1

												const kSearchStrength = 8

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:419
	nextEmit := s
												cv := load6432(src, s)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:423
	offset1 := int32(blk.recentOffsets[0])
	offset2 := int32(blk.recentOffsets[1])

	addLiterals := func(s *seq, until int32) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:426
		_go_fuzz_dep_.CoverTab[93012]++
													if until == nextEmit {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:427
			_go_fuzz_dep_.CoverTab[93014]++
														return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:428
			// _ = "end of CoverTab[93014]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:429
			_go_fuzz_dep_.CoverTab[93015]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:429
			// _ = "end of CoverTab[93015]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:429
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:429
		// _ = "end of CoverTab[93012]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:429
		_go_fuzz_dep_.CoverTab[93013]++
													blk.literals = append(blk.literals, src[nextEmit:until]...)
													s.litLen = uint32(until - nextEmit)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:431
		// _ = "end of CoverTab[93013]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:432
	// _ = "end of CoverTab[92998]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:432
	_go_fuzz_dep_.CoverTab[92999]++
												if debugEncoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:433
		_go_fuzz_dep_.CoverTab[93016]++
													println("recent offsets:", blk.recentOffsets)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:434
		// _ = "end of CoverTab[93016]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:435
		_go_fuzz_dep_.CoverTab[93017]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:435
		// _ = "end of CoverTab[93017]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:435
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:435
	// _ = "end of CoverTab[92999]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:435
	_go_fuzz_dep_.CoverTab[93000]++

encodeLoop:
	for {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:438
		_go_fuzz_dep_.CoverTab[93018]++
													var t int32
													for {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:440
			_go_fuzz_dep_.CoverTab[93027]++

														nextHashS := hashLen(cv, dFastShortTableBits, dFastShortLen)
														nextHashL := hashLen(cv, dFastLongTableBits, dFastLongLen)
														candidateL := e.longTable[nextHashL]
														candidateS := e.table[nextHashS]

														const repOff = 1
														repIndex := s - offset1 + repOff
														entry := tableEntry{offset: s + e.cur, val: uint32(cv)}
														e.longTable[nextHashL] = entry
														e.table[nextHashS] = entry

														if len(blk.sequences) > 2 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:453
				_go_fuzz_dep_.CoverTab[93032]++
															if load3232(src, repIndex) == uint32(cv>>(repOff*8)) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:454
					_go_fuzz_dep_.CoverTab[93033]++
																// Consider history as well.
																var seq seq

																length := 4 + int32(matchLen(src[s+4+repOff:], src[repIndex+4:]))

																seq.matchLen = uint32(length - zstdMinMatch)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:464
					start := s + repOff

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:467
					startLimit := nextEmit + 1

					tMin := s - e.maxMatchOff
					if tMin < 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:470
						_go_fuzz_dep_.CoverTab[93038]++
																	tMin = 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:471
						// _ = "end of CoverTab[93038]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:472
						_go_fuzz_dep_.CoverTab[93039]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:472
						// _ = "end of CoverTab[93039]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:472
					}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:472
					// _ = "end of CoverTab[93033]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:472
					_go_fuzz_dep_.CoverTab[93034]++
																for repIndex > tMin && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:473
						_go_fuzz_dep_.CoverTab[93040]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:473
						return start > startLimit
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:473
						// _ = "end of CoverTab[93040]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:473
					}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:473
						_go_fuzz_dep_.CoverTab[93041]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:473
						return src[repIndex-1] == src[start-1]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:473
						// _ = "end of CoverTab[93041]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:473
					}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:473
						_go_fuzz_dep_.CoverTab[93042]++
																	repIndex--
																	start--
																	seq.matchLen++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:476
						// _ = "end of CoverTab[93042]"
					}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:477
					// _ = "end of CoverTab[93034]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:477
					_go_fuzz_dep_.CoverTab[93035]++
																addLiterals(&seq, start)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:481
					seq.offset = 1
					if debugSequences {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:482
						_go_fuzz_dep_.CoverTab[93043]++
																	println("repeat sequence", seq, "next s:", s)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:483
						// _ = "end of CoverTab[93043]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:484
						_go_fuzz_dep_.CoverTab[93044]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:484
						// _ = "end of CoverTab[93044]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:484
					}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:484
					// _ = "end of CoverTab[93035]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:484
					_go_fuzz_dep_.CoverTab[93036]++
																blk.sequences = append(blk.sequences, seq)
																s += length + repOff
																nextEmit = s
																if s >= sLimit {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:488
						_go_fuzz_dep_.CoverTab[93045]++
																	if debugEncoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:489
							_go_fuzz_dep_.CoverTab[93047]++
																		println("repeat ended", s, length)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:490
							// _ = "end of CoverTab[93047]"

						} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:492
							_go_fuzz_dep_.CoverTab[93048]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:492
							// _ = "end of CoverTab[93048]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:492
						}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:492
						// _ = "end of CoverTab[93045]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:492
						_go_fuzz_dep_.CoverTab[93046]++
																	break encodeLoop
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:493
						// _ = "end of CoverTab[93046]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:494
						_go_fuzz_dep_.CoverTab[93049]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:494
						// _ = "end of CoverTab[93049]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:494
					}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:494
					// _ = "end of CoverTab[93036]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:494
					_go_fuzz_dep_.CoverTab[93037]++
																cv = load6432(src, s)
																continue
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:496
					// _ = "end of CoverTab[93037]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:497
					_go_fuzz_dep_.CoverTab[93050]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:497
					// _ = "end of CoverTab[93050]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:497
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:497
				// _ = "end of CoverTab[93032]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:498
				_go_fuzz_dep_.CoverTab[93051]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:498
				// _ = "end of CoverTab[93051]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:498
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:498
			// _ = "end of CoverTab[93027]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:498
			_go_fuzz_dep_.CoverTab[93028]++

														coffsetL := s - (candidateL.offset - e.cur)
														coffsetS := s - (candidateS.offset - e.cur)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:504
			if coffsetL < e.maxMatchOff && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:504
				_go_fuzz_dep_.CoverTab[93052]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:504
				return uint32(cv) == candidateL.val
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:504
				// _ = "end of CoverTab[93052]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:504
			}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:504
				_go_fuzz_dep_.CoverTab[93053]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:508
				t = candidateL.offset - e.cur
				if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:509
					_go_fuzz_dep_.CoverTab[93057]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:509
					return s <= t
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:509
					// _ = "end of CoverTab[93057]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:509
				}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:509
					_go_fuzz_dep_.CoverTab[93058]++
																panic(fmt.Sprintf("s (%d) <= t (%d). cur: %d", s, t, e.cur))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:510
					// _ = "end of CoverTab[93058]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:511
					_go_fuzz_dep_.CoverTab[93059]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:511
					// _ = "end of CoverTab[93059]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:511
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:511
				// _ = "end of CoverTab[93053]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:511
				_go_fuzz_dep_.CoverTab[93054]++
															if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:512
					_go_fuzz_dep_.CoverTab[93060]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:512
					return s-t > e.maxMatchOff
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:512
					// _ = "end of CoverTab[93060]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:512
				}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:512
					_go_fuzz_dep_.CoverTab[93061]++
																panic("s - t >e.maxMatchOff")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:513
					// _ = "end of CoverTab[93061]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:514
					_go_fuzz_dep_.CoverTab[93062]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:514
					// _ = "end of CoverTab[93062]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:514
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:514
				// _ = "end of CoverTab[93054]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:514
				_go_fuzz_dep_.CoverTab[93055]++
															if debugMatches {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:515
					_go_fuzz_dep_.CoverTab[93063]++
																println("long match")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:516
					// _ = "end of CoverTab[93063]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:517
					_go_fuzz_dep_.CoverTab[93064]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:517
					// _ = "end of CoverTab[93064]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:517
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:517
				// _ = "end of CoverTab[93055]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:517
				_go_fuzz_dep_.CoverTab[93056]++
															break
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:518
				// _ = "end of CoverTab[93056]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:519
				_go_fuzz_dep_.CoverTab[93065]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:519
				// _ = "end of CoverTab[93065]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:519
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:519
			// _ = "end of CoverTab[93028]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:519
			_go_fuzz_dep_.CoverTab[93029]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:522
			if coffsetS < e.maxMatchOff && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:522
				_go_fuzz_dep_.CoverTab[93066]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:522
				return uint32(cv) == candidateS.val
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:522
				// _ = "end of CoverTab[93066]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:522
			}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:522
				_go_fuzz_dep_.CoverTab[93067]++
				// found a regular match
															// See if we can find a long match at s+1
															const checkAt = 1
															cv := load6432(src, s+checkAt)
															nextHashL = hashLen(cv, dFastLongTableBits, dFastLongLen)
															candidateL = e.longTable[nextHashL]
															coffsetL = s - (candidateL.offset - e.cur) + checkAt

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:532
				e.longTable[nextHashL] = tableEntry{offset: s + checkAt + e.cur, val: uint32(cv)}
				if coffsetL < e.maxMatchOff && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:533
					_go_fuzz_dep_.CoverTab[93073]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:533
					return uint32(cv) == candidateL.val
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:533
					// _ = "end of CoverTab[93073]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:533
				}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:533
					_go_fuzz_dep_.CoverTab[93074]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:537
					t = candidateL.offset - e.cur
					s += checkAt
					if debugMatches {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:539
						_go_fuzz_dep_.CoverTab[93076]++
																	println("long match (after short)")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:540
						// _ = "end of CoverTab[93076]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:541
						_go_fuzz_dep_.CoverTab[93077]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:541
						// _ = "end of CoverTab[93077]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:541
					}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:541
					// _ = "end of CoverTab[93074]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:541
					_go_fuzz_dep_.CoverTab[93075]++
																break
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:542
					// _ = "end of CoverTab[93075]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:543
					_go_fuzz_dep_.CoverTab[93078]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:543
					// _ = "end of CoverTab[93078]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:543
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:543
				// _ = "end of CoverTab[93067]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:543
				_go_fuzz_dep_.CoverTab[93068]++

															t = candidateS.offset - e.cur
															if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:546
					_go_fuzz_dep_.CoverTab[93079]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:546
					return s <= t
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:546
					// _ = "end of CoverTab[93079]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:546
				}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:546
					_go_fuzz_dep_.CoverTab[93080]++
																panic(fmt.Sprintf("s (%d) <= t (%d)", s, t))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:547
					// _ = "end of CoverTab[93080]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:548
					_go_fuzz_dep_.CoverTab[93081]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:548
					// _ = "end of CoverTab[93081]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:548
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:548
				// _ = "end of CoverTab[93068]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:548
				_go_fuzz_dep_.CoverTab[93069]++
															if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:549
					_go_fuzz_dep_.CoverTab[93082]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:549
					return s-t > e.maxMatchOff
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:549
					// _ = "end of CoverTab[93082]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:549
				}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:549
					_go_fuzz_dep_.CoverTab[93083]++
																panic("s - t >e.maxMatchOff")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:550
					// _ = "end of CoverTab[93083]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:551
					_go_fuzz_dep_.CoverTab[93084]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:551
					// _ = "end of CoverTab[93084]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:551
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:551
				// _ = "end of CoverTab[93069]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:551
				_go_fuzz_dep_.CoverTab[93070]++
															if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:552
					_go_fuzz_dep_.CoverTab[93085]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:552
					return t < 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:552
					// _ = "end of CoverTab[93085]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:552
				}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:552
					_go_fuzz_dep_.CoverTab[93086]++
																panic("t<0")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:553
					// _ = "end of CoverTab[93086]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:554
					_go_fuzz_dep_.CoverTab[93087]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:554
					// _ = "end of CoverTab[93087]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:554
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:554
				// _ = "end of CoverTab[93070]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:554
				_go_fuzz_dep_.CoverTab[93071]++
															if debugMatches {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:555
					_go_fuzz_dep_.CoverTab[93088]++
																println("short match")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:556
					// _ = "end of CoverTab[93088]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:557
					_go_fuzz_dep_.CoverTab[93089]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:557
					// _ = "end of CoverTab[93089]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:557
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:557
				// _ = "end of CoverTab[93071]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:557
				_go_fuzz_dep_.CoverTab[93072]++
															break
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:558
				// _ = "end of CoverTab[93072]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:559
				_go_fuzz_dep_.CoverTab[93090]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:559
				// _ = "end of CoverTab[93090]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:559
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:559
			// _ = "end of CoverTab[93029]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:559
			_go_fuzz_dep_.CoverTab[93030]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:562
			s += stepSize + ((s - nextEmit) >> (kSearchStrength - 1))
			if s >= sLimit {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:563
				_go_fuzz_dep_.CoverTab[93091]++
															break encodeLoop
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:564
				// _ = "end of CoverTab[93091]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:565
				_go_fuzz_dep_.CoverTab[93092]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:565
				// _ = "end of CoverTab[93092]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:565
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:565
			// _ = "end of CoverTab[93030]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:565
			_go_fuzz_dep_.CoverTab[93031]++
														cv = load6432(src, s)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:566
			// _ = "end of CoverTab[93031]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:567
		// _ = "end of CoverTab[93018]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:567
		_go_fuzz_dep_.CoverTab[93019]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:571
		offset2 = offset1
		offset1 = s - t

		if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:574
			_go_fuzz_dep_.CoverTab[93093]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:574
			return s <= t
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:574
			// _ = "end of CoverTab[93093]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:574
		}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:574
			_go_fuzz_dep_.CoverTab[93094]++
														panic(fmt.Sprintf("s (%d) <= t (%d)", s, t))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:575
			// _ = "end of CoverTab[93094]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:576
			_go_fuzz_dep_.CoverTab[93095]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:576
			// _ = "end of CoverTab[93095]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:576
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:576
		// _ = "end of CoverTab[93019]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:576
		_go_fuzz_dep_.CoverTab[93020]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:580
		l := int32(matchLen(src[s+4:], src[t+4:])) + 4

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:583
		tMin := s - e.maxMatchOff
		if tMin < 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:584
			_go_fuzz_dep_.CoverTab[93096]++
														tMin = 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:585
			// _ = "end of CoverTab[93096]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:586
			_go_fuzz_dep_.CoverTab[93097]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:586
			// _ = "end of CoverTab[93097]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:586
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:586
		// _ = "end of CoverTab[93020]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:586
		_go_fuzz_dep_.CoverTab[93021]++
													for t > tMin && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:587
			_go_fuzz_dep_.CoverTab[93098]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:587
			return s > nextEmit
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:587
			// _ = "end of CoverTab[93098]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:587
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:587
			_go_fuzz_dep_.CoverTab[93099]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:587
			return src[t-1] == src[s-1]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:587
			// _ = "end of CoverTab[93099]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:587
		}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:587
			_go_fuzz_dep_.CoverTab[93100]++
														s--
														t--
														l++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:590
			// _ = "end of CoverTab[93100]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:591
		// _ = "end of CoverTab[93021]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:591
		_go_fuzz_dep_.CoverTab[93022]++

		// Write our sequence
		var seq seq
		seq.litLen = uint32(s - nextEmit)
		seq.matchLen = uint32(l - zstdMinMatch)
		if seq.litLen > 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:597
			_go_fuzz_dep_.CoverTab[93101]++
														blk.literals = append(blk.literals, src[nextEmit:s]...)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:598
			// _ = "end of CoverTab[93101]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:599
			_go_fuzz_dep_.CoverTab[93102]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:599
			// _ = "end of CoverTab[93102]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:599
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:599
		// _ = "end of CoverTab[93022]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:599
		_go_fuzz_dep_.CoverTab[93023]++
													seq.offset = uint32(s-t) + 3
													s += l
													if debugSequences {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:602
			_go_fuzz_dep_.CoverTab[93103]++
														println("sequence", seq, "next s:", s)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:603
			// _ = "end of CoverTab[93103]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:604
			_go_fuzz_dep_.CoverTab[93104]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:604
			// _ = "end of CoverTab[93104]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:604
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:604
		// _ = "end of CoverTab[93023]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:604
		_go_fuzz_dep_.CoverTab[93024]++
													blk.sequences = append(blk.sequences, seq)
													nextEmit = s
													if s >= sLimit {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:607
			_go_fuzz_dep_.CoverTab[93105]++
														break encodeLoop
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:608
			// _ = "end of CoverTab[93105]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:609
			_go_fuzz_dep_.CoverTab[93106]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:609
			// _ = "end of CoverTab[93106]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:609
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:609
		// _ = "end of CoverTab[93024]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:609
		_go_fuzz_dep_.CoverTab[93025]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:612
		index0 := s - l + 1

		index1 := s - 2

		cv0 := load6432(src, index0)
		cv1 := load6432(src, index1)
		te0 := tableEntry{offset: index0 + e.cur, val: uint32(cv0)}
		te1 := tableEntry{offset: index1 + e.cur, val: uint32(cv1)}
		e.longTable[hashLen(cv0, dFastLongTableBits, dFastLongLen)] = te0
		e.longTable[hashLen(cv1, dFastLongTableBits, dFastLongLen)] = te1
		cv0 >>= 8
		cv1 >>= 8
		te0.offset++
		te1.offset++
		te0.val = uint32(cv0)
		te1.val = uint32(cv1)
		e.table[hashLen(cv0, dFastShortTableBits, dFastShortLen)] = te0
		e.table[hashLen(cv1, dFastShortTableBits, dFastShortLen)] = te1

		cv = load6432(src, s)

		if len(blk.sequences) <= 2 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:633
			_go_fuzz_dep_.CoverTab[93107]++
														continue
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:634
			// _ = "end of CoverTab[93107]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:635
			_go_fuzz_dep_.CoverTab[93108]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:635
			// _ = "end of CoverTab[93108]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:635
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:635
		// _ = "end of CoverTab[93025]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:635
		_go_fuzz_dep_.CoverTab[93026]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:638
		for {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:638
			_go_fuzz_dep_.CoverTab[93109]++
														o2 := s - offset2
														if load3232(src, o2) != uint32(cv) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:640
				_go_fuzz_dep_.CoverTab[93113]++

															break
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:642
				// _ = "end of CoverTab[93113]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:643
				_go_fuzz_dep_.CoverTab[93114]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:643
				// _ = "end of CoverTab[93114]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:643
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:643
			// _ = "end of CoverTab[93109]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:643
			_go_fuzz_dep_.CoverTab[93110]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:646
			nextHashS := hashLen(cv1>>8, dFastShortTableBits, dFastShortLen)
														nextHashL := hashLen(cv, dFastLongTableBits, dFastLongLen)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:652
			l := 4 + int32(matchLen(src[s+4:], src[o2+4:]))

														entry := tableEntry{offset: s + e.cur, val: uint32(cv)}
														e.longTable[nextHashL] = entry
														e.table[nextHashS] = entry
														seq.matchLen = uint32(l) - zstdMinMatch
														seq.litLen = 0

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:661
			seq.offset = 1
			s += l
			nextEmit = s
			if debugSequences {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:664
				_go_fuzz_dep_.CoverTab[93115]++
															println("sequence", seq, "next s:", s)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:665
				// _ = "end of CoverTab[93115]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:666
				_go_fuzz_dep_.CoverTab[93116]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:666
				// _ = "end of CoverTab[93116]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:666
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:666
			// _ = "end of CoverTab[93110]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:666
			_go_fuzz_dep_.CoverTab[93111]++
														blk.sequences = append(blk.sequences, seq)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:670
			offset1, offset2 = offset2, offset1
			if s >= sLimit {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:671
				_go_fuzz_dep_.CoverTab[93117]++

															break encodeLoop
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:673
				// _ = "end of CoverTab[93117]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:674
				_go_fuzz_dep_.CoverTab[93118]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:674
				// _ = "end of CoverTab[93118]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:674
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:674
			// _ = "end of CoverTab[93111]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:674
			_go_fuzz_dep_.CoverTab[93112]++
														cv = load6432(src, s)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:675
			// _ = "end of CoverTab[93112]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:676
		// _ = "end of CoverTab[93026]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:677
	// _ = "end of CoverTab[93000]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:677
	_go_fuzz_dep_.CoverTab[93001]++

												if int(nextEmit) < len(src) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:679
		_go_fuzz_dep_.CoverTab[93119]++
													blk.literals = append(blk.literals, src[nextEmit:]...)
													blk.extraLits = len(src) - int(nextEmit)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:681
		// _ = "end of CoverTab[93119]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:682
		_go_fuzz_dep_.CoverTab[93120]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:682
		// _ = "end of CoverTab[93120]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:682
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:682
	// _ = "end of CoverTab[93001]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:682
	_go_fuzz_dep_.CoverTab[93002]++
												if debugEncoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:683
		_go_fuzz_dep_.CoverTab[93121]++
													println("returning, recent offsets:", blk.recentOffsets, "extra literals:", blk.extraLits)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:684
		// _ = "end of CoverTab[93121]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:685
		_go_fuzz_dep_.CoverTab[93122]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:685
		// _ = "end of CoverTab[93122]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:685
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:685
	// _ = "end of CoverTab[93002]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:685
	_go_fuzz_dep_.CoverTab[93003]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:688
	if e.cur < bufferReset {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:688
		_go_fuzz_dep_.CoverTab[93123]++
													e.cur += int32(len(src))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:689
		// _ = "end of CoverTab[93123]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:690
		_go_fuzz_dep_.CoverTab[93124]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:690
		// _ = "end of CoverTab[93124]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:690
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:690
	// _ = "end of CoverTab[93003]"
}

// Encode will encode the content, with a dictionary if initialized for it.
func (e *doubleFastEncoderDict) Encode(blk *blockEnc, src []byte) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:694
	_go_fuzz_dep_.CoverTab[93125]++
												const (
		// Input margin is the number of bytes we read (8)
		// and the maximum we will read ahead (2)
		inputMargin		= 8 + 2
		minNonLiteralBlockSize	= 16
	)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:703
	for e.cur >= bufferReset {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:703
		_go_fuzz_dep_.CoverTab[93133]++
													if len(e.hist) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:704
			_go_fuzz_dep_.CoverTab[93137]++
														for i := range e.table[:] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:705
				_go_fuzz_dep_.CoverTab[93140]++
															e.table[i] = tableEntry{}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:706
				// _ = "end of CoverTab[93140]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:707
			// _ = "end of CoverTab[93137]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:707
			_go_fuzz_dep_.CoverTab[93138]++
														for i := range e.longTable[:] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:708
				_go_fuzz_dep_.CoverTab[93141]++
															e.longTable[i] = tableEntry{}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:709
				// _ = "end of CoverTab[93141]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:710
			// _ = "end of CoverTab[93138]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:710
			_go_fuzz_dep_.CoverTab[93139]++
														e.markAllShardsDirty()
														e.cur = e.maxMatchOff
														break
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:713
			// _ = "end of CoverTab[93139]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:714
			_go_fuzz_dep_.CoverTab[93142]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:714
			// _ = "end of CoverTab[93142]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:714
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:714
		// _ = "end of CoverTab[93133]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:714
		_go_fuzz_dep_.CoverTab[93134]++

													minOff := e.cur + int32(len(e.hist)) - e.maxMatchOff
													for i := range e.table[:] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:717
			_go_fuzz_dep_.CoverTab[93143]++
														v := e.table[i].offset
														if v < minOff {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:719
				_go_fuzz_dep_.CoverTab[93145]++
															v = 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:720
				// _ = "end of CoverTab[93145]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:721
				_go_fuzz_dep_.CoverTab[93146]++
															v = v - e.cur + e.maxMatchOff
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:722
				// _ = "end of CoverTab[93146]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:723
			// _ = "end of CoverTab[93143]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:723
			_go_fuzz_dep_.CoverTab[93144]++
														e.table[i].offset = v
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:724
			// _ = "end of CoverTab[93144]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:725
		// _ = "end of CoverTab[93134]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:725
		_go_fuzz_dep_.CoverTab[93135]++
													for i := range e.longTable[:] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:726
			_go_fuzz_dep_.CoverTab[93147]++
														v := e.longTable[i].offset
														if v < minOff {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:728
				_go_fuzz_dep_.CoverTab[93149]++
															v = 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:729
				// _ = "end of CoverTab[93149]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:730
				_go_fuzz_dep_.CoverTab[93150]++
															v = v - e.cur + e.maxMatchOff
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:731
				// _ = "end of CoverTab[93150]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:732
			// _ = "end of CoverTab[93147]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:732
			_go_fuzz_dep_.CoverTab[93148]++
														e.longTable[i].offset = v
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:733
			// _ = "end of CoverTab[93148]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:734
		// _ = "end of CoverTab[93135]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:734
		_go_fuzz_dep_.CoverTab[93136]++
													e.markAllShardsDirty()
													e.cur = e.maxMatchOff
													break
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:737
		// _ = "end of CoverTab[93136]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:738
	// _ = "end of CoverTab[93125]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:738
	_go_fuzz_dep_.CoverTab[93126]++

												s := e.addBlock(src)
												blk.size = len(src)
												if len(src) < minNonLiteralBlockSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:742
		_go_fuzz_dep_.CoverTab[93151]++
													blk.extraLits = len(src)
													blk.literals = blk.literals[:len(src)]
													copy(blk.literals, src)
													return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:746
		// _ = "end of CoverTab[93151]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:747
		_go_fuzz_dep_.CoverTab[93152]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:747
		// _ = "end of CoverTab[93152]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:747
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:747
	// _ = "end of CoverTab[93126]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:747
	_go_fuzz_dep_.CoverTab[93127]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:750
	src = e.hist
	sLimit := int32(len(src)) - inputMargin
	// stepSize is the number of bytes to skip on every main loop iteration.
												// It should be >= 1.
												const stepSize = 1

												const kSearchStrength = 8

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:759
	nextEmit := s
												cv := load6432(src, s)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:763
	offset1 := int32(blk.recentOffsets[0])
	offset2 := int32(blk.recentOffsets[1])

	addLiterals := func(s *seq, until int32) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:766
		_go_fuzz_dep_.CoverTab[93153]++
													if until == nextEmit {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:767
			_go_fuzz_dep_.CoverTab[93155]++
														return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:768
			// _ = "end of CoverTab[93155]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:769
			_go_fuzz_dep_.CoverTab[93156]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:769
			// _ = "end of CoverTab[93156]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:769
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:769
		// _ = "end of CoverTab[93153]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:769
		_go_fuzz_dep_.CoverTab[93154]++
													blk.literals = append(blk.literals, src[nextEmit:until]...)
													s.litLen = uint32(until - nextEmit)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:771
		// _ = "end of CoverTab[93154]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:772
	// _ = "end of CoverTab[93127]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:772
	_go_fuzz_dep_.CoverTab[93128]++
												if debugEncoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:773
		_go_fuzz_dep_.CoverTab[93157]++
													println("recent offsets:", blk.recentOffsets)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:774
		// _ = "end of CoverTab[93157]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:775
		_go_fuzz_dep_.CoverTab[93158]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:775
		// _ = "end of CoverTab[93158]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:775
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:775
	// _ = "end of CoverTab[93128]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:775
	_go_fuzz_dep_.CoverTab[93129]++

encodeLoop:
	for {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:778
		_go_fuzz_dep_.CoverTab[93159]++
													var t int32

													canRepeat := len(blk.sequences) > 2

													for {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:783
			_go_fuzz_dep_.CoverTab[93169]++
														if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:784
				_go_fuzz_dep_.CoverTab[93175]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:784
				return canRepeat
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:784
				// _ = "end of CoverTab[93175]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:784
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:784
				_go_fuzz_dep_.CoverTab[93176]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:784
				return offset1 == 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:784
				// _ = "end of CoverTab[93176]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:784
			}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:784
				_go_fuzz_dep_.CoverTab[93177]++
															panic("offset0 was 0")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:785
				// _ = "end of CoverTab[93177]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:786
				_go_fuzz_dep_.CoverTab[93178]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:786
				// _ = "end of CoverTab[93178]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:786
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:786
			// _ = "end of CoverTab[93169]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:786
			_go_fuzz_dep_.CoverTab[93170]++

														nextHashS := hashLen(cv, dFastShortTableBits, dFastShortLen)
														nextHashL := hashLen(cv, dFastLongTableBits, dFastLongLen)
														candidateL := e.longTable[nextHashL]
														candidateS := e.table[nextHashS]

														const repOff = 1
														repIndex := s - offset1 + repOff
														entry := tableEntry{offset: s + e.cur, val: uint32(cv)}
														e.longTable[nextHashL] = entry
														e.markLongShardDirty(nextHashL)
														e.table[nextHashS] = entry
														e.markShardDirty(nextHashS)

														if canRepeat {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:801
				_go_fuzz_dep_.CoverTab[93179]++
															if repIndex >= 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:802
					_go_fuzz_dep_.CoverTab[93180]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:802
					return load3232(src, repIndex) == uint32(cv>>(repOff*8))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:802
					// _ = "end of CoverTab[93180]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:802
				}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:802
					_go_fuzz_dep_.CoverTab[93181]++
																// Consider history as well.
																var seq seq
																lenght := 4 + e.matchlen(s+4+repOff, repIndex+4, src)

																seq.matchLen = uint32(lenght - zstdMinMatch)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:811
					start := s + repOff

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:814
					startLimit := nextEmit + 1

					tMin := s - e.maxMatchOff
					if tMin < 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:817
						_go_fuzz_dep_.CoverTab[93186]++
																	tMin = 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:818
						// _ = "end of CoverTab[93186]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:819
						_go_fuzz_dep_.CoverTab[93187]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:819
						// _ = "end of CoverTab[93187]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:819
					}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:819
					// _ = "end of CoverTab[93181]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:819
					_go_fuzz_dep_.CoverTab[93182]++
																for repIndex > tMin && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:820
						_go_fuzz_dep_.CoverTab[93188]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:820
						return start > startLimit
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:820
						// _ = "end of CoverTab[93188]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:820
					}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:820
						_go_fuzz_dep_.CoverTab[93189]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:820
						return src[repIndex-1] == src[start-1]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:820
						// _ = "end of CoverTab[93189]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:820
					}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:820
						_go_fuzz_dep_.CoverTab[93190]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:820
						return seq.matchLen < maxMatchLength-zstdMinMatch-1
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:820
						// _ = "end of CoverTab[93190]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:820
					}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:820
						_go_fuzz_dep_.CoverTab[93191]++
																	repIndex--
																	start--
																	seq.matchLen++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:823
						// _ = "end of CoverTab[93191]"
					}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:824
					// _ = "end of CoverTab[93182]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:824
					_go_fuzz_dep_.CoverTab[93183]++
																addLiterals(&seq, start)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:828
					seq.offset = 1
					if debugSequences {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:829
						_go_fuzz_dep_.CoverTab[93192]++
																	println("repeat sequence", seq, "next s:", s)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:830
						// _ = "end of CoverTab[93192]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:831
						_go_fuzz_dep_.CoverTab[93193]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:831
						// _ = "end of CoverTab[93193]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:831
					}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:831
					// _ = "end of CoverTab[93183]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:831
					_go_fuzz_dep_.CoverTab[93184]++
																blk.sequences = append(blk.sequences, seq)
																s += lenght + repOff
																nextEmit = s
																if s >= sLimit {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:835
						_go_fuzz_dep_.CoverTab[93194]++
																	if debugEncoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:836
							_go_fuzz_dep_.CoverTab[93196]++
																		println("repeat ended", s, lenght)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:837
							// _ = "end of CoverTab[93196]"

						} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:839
							_go_fuzz_dep_.CoverTab[93197]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:839
							// _ = "end of CoverTab[93197]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:839
						}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:839
						// _ = "end of CoverTab[93194]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:839
						_go_fuzz_dep_.CoverTab[93195]++
																	break encodeLoop
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:840
						// _ = "end of CoverTab[93195]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:841
						_go_fuzz_dep_.CoverTab[93198]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:841
						// _ = "end of CoverTab[93198]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:841
					}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:841
					// _ = "end of CoverTab[93184]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:841
					_go_fuzz_dep_.CoverTab[93185]++
																cv = load6432(src, s)
																continue
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:843
					// _ = "end of CoverTab[93185]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:844
					_go_fuzz_dep_.CoverTab[93199]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:844
					// _ = "end of CoverTab[93199]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:844
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:844
				// _ = "end of CoverTab[93179]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:845
				_go_fuzz_dep_.CoverTab[93200]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:845
				// _ = "end of CoverTab[93200]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:845
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:845
			// _ = "end of CoverTab[93170]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:845
			_go_fuzz_dep_.CoverTab[93171]++

														coffsetL := s - (candidateL.offset - e.cur)
														coffsetS := s - (candidateS.offset - e.cur)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:851
			if coffsetL < e.maxMatchOff && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:851
				_go_fuzz_dep_.CoverTab[93201]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:851
				return uint32(cv) == candidateL.val
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:851
				// _ = "end of CoverTab[93201]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:851
			}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:851
				_go_fuzz_dep_.CoverTab[93202]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:855
				t = candidateL.offset - e.cur
				if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:856
					_go_fuzz_dep_.CoverTab[93206]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:856
					return s <= t
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:856
					// _ = "end of CoverTab[93206]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:856
				}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:856
					_go_fuzz_dep_.CoverTab[93207]++
																panic(fmt.Sprintf("s (%d) <= t (%d)", s, t))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:857
					// _ = "end of CoverTab[93207]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:858
					_go_fuzz_dep_.CoverTab[93208]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:858
					// _ = "end of CoverTab[93208]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:858
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:858
				// _ = "end of CoverTab[93202]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:858
				_go_fuzz_dep_.CoverTab[93203]++
															if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:859
					_go_fuzz_dep_.CoverTab[93209]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:859
					return s-t > e.maxMatchOff
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:859
					// _ = "end of CoverTab[93209]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:859
				}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:859
					_go_fuzz_dep_.CoverTab[93210]++
																panic("s - t >e.maxMatchOff")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:860
					// _ = "end of CoverTab[93210]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:861
					_go_fuzz_dep_.CoverTab[93211]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:861
					// _ = "end of CoverTab[93211]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:861
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:861
				// _ = "end of CoverTab[93203]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:861
				_go_fuzz_dep_.CoverTab[93204]++
															if debugMatches {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:862
					_go_fuzz_dep_.CoverTab[93212]++
																println("long match")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:863
					// _ = "end of CoverTab[93212]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:864
					_go_fuzz_dep_.CoverTab[93213]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:864
					// _ = "end of CoverTab[93213]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:864
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:864
				// _ = "end of CoverTab[93204]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:864
				_go_fuzz_dep_.CoverTab[93205]++
															break
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:865
				// _ = "end of CoverTab[93205]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:866
				_go_fuzz_dep_.CoverTab[93214]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:866
				// _ = "end of CoverTab[93214]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:866
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:866
			// _ = "end of CoverTab[93171]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:866
			_go_fuzz_dep_.CoverTab[93172]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:869
			if coffsetS < e.maxMatchOff && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:869
				_go_fuzz_dep_.CoverTab[93215]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:869
				return uint32(cv) == candidateS.val
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:869
				// _ = "end of CoverTab[93215]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:869
			}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:869
				_go_fuzz_dep_.CoverTab[93216]++
				// found a regular match
															// See if we can find a long match at s+1
															const checkAt = 1
															cv := load6432(src, s+checkAt)
															nextHashL = hashLen(cv, dFastLongTableBits, dFastLongLen)
															candidateL = e.longTable[nextHashL]
															coffsetL = s - (candidateL.offset - e.cur) + checkAt

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:879
				e.longTable[nextHashL] = tableEntry{offset: s + checkAt + e.cur, val: uint32(cv)}
				e.markLongShardDirty(nextHashL)
				if coffsetL < e.maxMatchOff && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:881
					_go_fuzz_dep_.CoverTab[93222]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:881
					return uint32(cv) == candidateL.val
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:881
					// _ = "end of CoverTab[93222]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:881
				}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:881
					_go_fuzz_dep_.CoverTab[93223]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:885
					t = candidateL.offset - e.cur
					s += checkAt
					if debugMatches {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:887
						_go_fuzz_dep_.CoverTab[93225]++
																	println("long match (after short)")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:888
						// _ = "end of CoverTab[93225]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:889
						_go_fuzz_dep_.CoverTab[93226]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:889
						// _ = "end of CoverTab[93226]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:889
					}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:889
					// _ = "end of CoverTab[93223]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:889
					_go_fuzz_dep_.CoverTab[93224]++
																break
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:890
					// _ = "end of CoverTab[93224]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:891
					_go_fuzz_dep_.CoverTab[93227]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:891
					// _ = "end of CoverTab[93227]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:891
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:891
				// _ = "end of CoverTab[93216]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:891
				_go_fuzz_dep_.CoverTab[93217]++

															t = candidateS.offset - e.cur
															if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:894
					_go_fuzz_dep_.CoverTab[93228]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:894
					return s <= t
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:894
					// _ = "end of CoverTab[93228]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:894
				}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:894
					_go_fuzz_dep_.CoverTab[93229]++
																panic(fmt.Sprintf("s (%d) <= t (%d)", s, t))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:895
					// _ = "end of CoverTab[93229]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:896
					_go_fuzz_dep_.CoverTab[93230]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:896
					// _ = "end of CoverTab[93230]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:896
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:896
				// _ = "end of CoverTab[93217]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:896
				_go_fuzz_dep_.CoverTab[93218]++
															if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:897
					_go_fuzz_dep_.CoverTab[93231]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:897
					return s-t > e.maxMatchOff
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:897
					// _ = "end of CoverTab[93231]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:897
				}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:897
					_go_fuzz_dep_.CoverTab[93232]++
																panic("s - t >e.maxMatchOff")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:898
					// _ = "end of CoverTab[93232]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:899
					_go_fuzz_dep_.CoverTab[93233]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:899
					// _ = "end of CoverTab[93233]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:899
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:899
				// _ = "end of CoverTab[93218]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:899
				_go_fuzz_dep_.CoverTab[93219]++
															if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:900
					_go_fuzz_dep_.CoverTab[93234]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:900
					return t < 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:900
					// _ = "end of CoverTab[93234]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:900
				}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:900
					_go_fuzz_dep_.CoverTab[93235]++
																panic("t<0")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:901
					// _ = "end of CoverTab[93235]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:902
					_go_fuzz_dep_.CoverTab[93236]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:902
					// _ = "end of CoverTab[93236]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:902
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:902
				// _ = "end of CoverTab[93219]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:902
				_go_fuzz_dep_.CoverTab[93220]++
															if debugMatches {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:903
					_go_fuzz_dep_.CoverTab[93237]++
																println("short match")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:904
					// _ = "end of CoverTab[93237]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:905
					_go_fuzz_dep_.CoverTab[93238]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:905
					// _ = "end of CoverTab[93238]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:905
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:905
				// _ = "end of CoverTab[93220]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:905
				_go_fuzz_dep_.CoverTab[93221]++
															break
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:906
				// _ = "end of CoverTab[93221]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:907
				_go_fuzz_dep_.CoverTab[93239]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:907
				// _ = "end of CoverTab[93239]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:907
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:907
			// _ = "end of CoverTab[93172]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:907
			_go_fuzz_dep_.CoverTab[93173]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:910
			s += stepSize + ((s - nextEmit) >> (kSearchStrength - 1))
			if s >= sLimit {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:911
				_go_fuzz_dep_.CoverTab[93240]++
															break encodeLoop
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:912
				// _ = "end of CoverTab[93240]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:913
				_go_fuzz_dep_.CoverTab[93241]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:913
				// _ = "end of CoverTab[93241]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:913
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:913
			// _ = "end of CoverTab[93173]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:913
			_go_fuzz_dep_.CoverTab[93174]++
														cv = load6432(src, s)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:914
			// _ = "end of CoverTab[93174]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:915
		// _ = "end of CoverTab[93159]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:915
		_go_fuzz_dep_.CoverTab[93160]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:919
		offset2 = offset1
		offset1 = s - t

		if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:922
			_go_fuzz_dep_.CoverTab[93242]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:922
			return s <= t
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:922
			// _ = "end of CoverTab[93242]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:922
		}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:922
			_go_fuzz_dep_.CoverTab[93243]++
														panic(fmt.Sprintf("s (%d) <= t (%d)", s, t))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:923
			// _ = "end of CoverTab[93243]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:924
			_go_fuzz_dep_.CoverTab[93244]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:924
			// _ = "end of CoverTab[93244]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:924
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:924
		// _ = "end of CoverTab[93160]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:924
		_go_fuzz_dep_.CoverTab[93161]++

													if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:926
			_go_fuzz_dep_.CoverTab[93245]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:926
			return canRepeat
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:926
			// _ = "end of CoverTab[93245]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:926
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:926
			_go_fuzz_dep_.CoverTab[93246]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:926
			return int(offset1) > len(src)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:926
			// _ = "end of CoverTab[93246]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:926
		}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:926
			_go_fuzz_dep_.CoverTab[93247]++
														panic("invalid offset")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:927
			// _ = "end of CoverTab[93247]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:928
			_go_fuzz_dep_.CoverTab[93248]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:928
			// _ = "end of CoverTab[93248]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:928
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:928
		// _ = "end of CoverTab[93161]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:928
		_go_fuzz_dep_.CoverTab[93162]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:931
		l := e.matchlen(s+4, t+4, src) + 4

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:934
		tMin := s - e.maxMatchOff
		if tMin < 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:935
			_go_fuzz_dep_.CoverTab[93249]++
														tMin = 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:936
			// _ = "end of CoverTab[93249]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:937
			_go_fuzz_dep_.CoverTab[93250]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:937
			// _ = "end of CoverTab[93250]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:937
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:937
		// _ = "end of CoverTab[93162]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:937
		_go_fuzz_dep_.CoverTab[93163]++
													for t > tMin && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:938
			_go_fuzz_dep_.CoverTab[93251]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:938
			return s > nextEmit
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:938
			// _ = "end of CoverTab[93251]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:938
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:938
			_go_fuzz_dep_.CoverTab[93252]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:938
			return src[t-1] == src[s-1]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:938
			// _ = "end of CoverTab[93252]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:938
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:938
			_go_fuzz_dep_.CoverTab[93253]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:938
			return l < maxMatchLength
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:938
			// _ = "end of CoverTab[93253]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:938
		}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:938
			_go_fuzz_dep_.CoverTab[93254]++
														s--
														t--
														l++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:941
			// _ = "end of CoverTab[93254]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:942
		// _ = "end of CoverTab[93163]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:942
		_go_fuzz_dep_.CoverTab[93164]++

		// Write our sequence
		var seq seq
		seq.litLen = uint32(s - nextEmit)
		seq.matchLen = uint32(l - zstdMinMatch)
		if seq.litLen > 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:948
			_go_fuzz_dep_.CoverTab[93255]++
														blk.literals = append(blk.literals, src[nextEmit:s]...)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:949
			// _ = "end of CoverTab[93255]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:950
			_go_fuzz_dep_.CoverTab[93256]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:950
			// _ = "end of CoverTab[93256]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:950
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:950
		// _ = "end of CoverTab[93164]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:950
		_go_fuzz_dep_.CoverTab[93165]++
													seq.offset = uint32(s-t) + 3
													s += l
													if debugSequences {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:953
			_go_fuzz_dep_.CoverTab[93257]++
														println("sequence", seq, "next s:", s)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:954
			// _ = "end of CoverTab[93257]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:955
			_go_fuzz_dep_.CoverTab[93258]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:955
			// _ = "end of CoverTab[93258]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:955
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:955
		// _ = "end of CoverTab[93165]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:955
		_go_fuzz_dep_.CoverTab[93166]++
													blk.sequences = append(blk.sequences, seq)
													nextEmit = s
													if s >= sLimit {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:958
			_go_fuzz_dep_.CoverTab[93259]++
														break encodeLoop
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:959
			// _ = "end of CoverTab[93259]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:960
			_go_fuzz_dep_.CoverTab[93260]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:960
			// _ = "end of CoverTab[93260]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:960
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:960
		// _ = "end of CoverTab[93166]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:960
		_go_fuzz_dep_.CoverTab[93167]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:963
		index0 := s - l + 1

		index1 := s - 2

		cv0 := load6432(src, index0)
		cv1 := load6432(src, index1)
		te0 := tableEntry{offset: index0 + e.cur, val: uint32(cv0)}
		te1 := tableEntry{offset: index1 + e.cur, val: uint32(cv1)}
		longHash1 := hashLen(cv0, dFastLongTableBits, dFastLongLen)
		longHash2 := hashLen(cv0, dFastLongTableBits, dFastLongLen)
		e.longTable[longHash1] = te0
		e.longTable[longHash2] = te1
		e.markLongShardDirty(longHash1)
		e.markLongShardDirty(longHash2)
		cv0 >>= 8
		cv1 >>= 8
		te0.offset++
		te1.offset++
		te0.val = uint32(cv0)
		te1.val = uint32(cv1)
		hashVal1 := hashLen(cv0, dFastShortTableBits, dFastShortLen)
		hashVal2 := hashLen(cv1, dFastShortTableBits, dFastShortLen)
		e.table[hashVal1] = te0
		e.markShardDirty(hashVal1)
		e.table[hashVal2] = te1
		e.markShardDirty(hashVal2)

		cv = load6432(src, s)

		if !canRepeat {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:992
			_go_fuzz_dep_.CoverTab[93261]++
														continue
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:993
			// _ = "end of CoverTab[93261]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:994
			_go_fuzz_dep_.CoverTab[93262]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:994
			// _ = "end of CoverTab[93262]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:994
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:994
		// _ = "end of CoverTab[93167]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:994
		_go_fuzz_dep_.CoverTab[93168]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:997
		for {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:997
			_go_fuzz_dep_.CoverTab[93263]++
														o2 := s - offset2
														if load3232(src, o2) != uint32(cv) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:999
				_go_fuzz_dep_.CoverTab[93267]++

															break
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1001
				// _ = "end of CoverTab[93267]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1002
				_go_fuzz_dep_.CoverTab[93268]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1002
				// _ = "end of CoverTab[93268]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1002
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1002
			// _ = "end of CoverTab[93263]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1002
			_go_fuzz_dep_.CoverTab[93264]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1005
			nextHashS := hashLen(cv, dFastShortTableBits, dFastShortLen)
														nextHashL := hashLen(cv, dFastLongTableBits, dFastLongLen)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1010
			l := 4 + e.matchlen(s+4, o2+4, src)

														entry := tableEntry{offset: s + e.cur, val: uint32(cv)}
														e.longTable[nextHashL] = entry
														e.markLongShardDirty(nextHashL)
														e.table[nextHashS] = entry
														e.markShardDirty(nextHashS)
														seq.matchLen = uint32(l) - zstdMinMatch
														seq.litLen = 0

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1021
			seq.offset = 1
			s += l
			nextEmit = s
			if debugSequences {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1024
				_go_fuzz_dep_.CoverTab[93269]++
															println("sequence", seq, "next s:", s)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1025
				// _ = "end of CoverTab[93269]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1026
				_go_fuzz_dep_.CoverTab[93270]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1026
				// _ = "end of CoverTab[93270]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1026
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1026
			// _ = "end of CoverTab[93264]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1026
			_go_fuzz_dep_.CoverTab[93265]++
														blk.sequences = append(blk.sequences, seq)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1030
			offset1, offset2 = offset2, offset1
			if s >= sLimit {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1031
				_go_fuzz_dep_.CoverTab[93271]++

															break encodeLoop
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1033
				// _ = "end of CoverTab[93271]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1034
				_go_fuzz_dep_.CoverTab[93272]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1034
				// _ = "end of CoverTab[93272]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1034
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1034
			// _ = "end of CoverTab[93265]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1034
			_go_fuzz_dep_.CoverTab[93266]++
														cv = load6432(src, s)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1035
			// _ = "end of CoverTab[93266]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1036
		// _ = "end of CoverTab[93168]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1037
	// _ = "end of CoverTab[93129]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1037
	_go_fuzz_dep_.CoverTab[93130]++

												if int(nextEmit) < len(src) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1039
		_go_fuzz_dep_.CoverTab[93273]++
													blk.literals = append(blk.literals, src[nextEmit:]...)
													blk.extraLits = len(src) - int(nextEmit)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1041
		// _ = "end of CoverTab[93273]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1042
		_go_fuzz_dep_.CoverTab[93274]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1042
		// _ = "end of CoverTab[93274]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1042
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1042
	// _ = "end of CoverTab[93130]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1042
	_go_fuzz_dep_.CoverTab[93131]++
												blk.recentOffsets[0] = uint32(offset1)
												blk.recentOffsets[1] = uint32(offset2)
												if debugEncoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1045
		_go_fuzz_dep_.CoverTab[93275]++
													println("returning, recent offsets:", blk.recentOffsets, "extra literals:", blk.extraLits)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1046
		// _ = "end of CoverTab[93275]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1047
		_go_fuzz_dep_.CoverTab[93276]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1047
		// _ = "end of CoverTab[93276]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1047
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1047
	// _ = "end of CoverTab[93131]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1047
	_go_fuzz_dep_.CoverTab[93132]++

												if len(src) > 64<<10 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1049
		_go_fuzz_dep_.CoverTab[93277]++
													e.markAllShardsDirty()
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1050
		// _ = "end of CoverTab[93277]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1051
		_go_fuzz_dep_.CoverTab[93278]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1051
		// _ = "end of CoverTab[93278]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1051
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1051
	// _ = "end of CoverTab[93132]"
}

// ResetDict will reset and set a dictionary if not nil
func (e *doubleFastEncoder) Reset(d *dict, singleBlock bool) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1055
	_go_fuzz_dep_.CoverTab[93279]++
												e.fastEncoder.Reset(d, singleBlock)
												if d != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1057
		_go_fuzz_dep_.CoverTab[93280]++
													panic("doubleFastEncoder: Reset with dict not supported")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1058
		// _ = "end of CoverTab[93280]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1059
		_go_fuzz_dep_.CoverTab[93281]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1059
		// _ = "end of CoverTab[93281]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1059
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1059
	// _ = "end of CoverTab[93279]"
}

// ResetDict will reset and set a dictionary if not nil
func (e *doubleFastEncoderDict) Reset(d *dict, singleBlock bool) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1063
	_go_fuzz_dep_.CoverTab[93282]++
												allDirty := e.allDirty
												e.fastEncoderDict.Reset(d, singleBlock)
												if d == nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1066
		_go_fuzz_dep_.CoverTab[93287]++
													return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1067
		// _ = "end of CoverTab[93287]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1068
		_go_fuzz_dep_.CoverTab[93288]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1068
		// _ = "end of CoverTab[93288]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1068
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1068
	// _ = "end of CoverTab[93282]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1068
	_go_fuzz_dep_.CoverTab[93283]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1071
	if len(e.dictLongTable) != len(e.longTable) || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1071
		_go_fuzz_dep_.CoverTab[93289]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1071
		return d.id != e.lastDictID
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1071
		// _ = "end of CoverTab[93289]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1071
	}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1071
		_go_fuzz_dep_.CoverTab[93290]++
													if len(e.dictLongTable) != len(e.longTable) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1072
			_go_fuzz_dep_.CoverTab[93293]++
														e.dictLongTable = make([]tableEntry, len(e.longTable))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1073
			// _ = "end of CoverTab[93293]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1074
			_go_fuzz_dep_.CoverTab[93294]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1074
			// _ = "end of CoverTab[93294]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1074
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1074
		// _ = "end of CoverTab[93290]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1074
		_go_fuzz_dep_.CoverTab[93291]++
													if len(d.content) >= 8 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1075
			_go_fuzz_dep_.CoverTab[93295]++
														cv := load6432(d.content, 0)
														e.dictLongTable[hashLen(cv, dFastLongTableBits, dFastLongLen)] = tableEntry{
				val:	uint32(cv),
				offset:	e.maxMatchOff,
			}
			end := int32(len(d.content)) - 8 + e.maxMatchOff
			for i := e.maxMatchOff + 1; i < end; i++ {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1082
				_go_fuzz_dep_.CoverTab[93296]++
															cv = cv>>8 | (uint64(d.content[i-e.maxMatchOff+7]) << 56)
															e.dictLongTable[hashLen(cv, dFastLongTableBits, dFastLongLen)] = tableEntry{
					val:	uint32(cv),
					offset:	i,
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1087
				// _ = "end of CoverTab[93296]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1088
			// _ = "end of CoverTab[93295]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1089
			_go_fuzz_dep_.CoverTab[93297]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1089
			// _ = "end of CoverTab[93297]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1089
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1089
		// _ = "end of CoverTab[93291]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1089
		_go_fuzz_dep_.CoverTab[93292]++
													e.lastDictID = d.id
													e.allDirty = true
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1091
		// _ = "end of CoverTab[93292]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1092
		_go_fuzz_dep_.CoverTab[93298]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1092
		// _ = "end of CoverTab[93298]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1092
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1092
	// _ = "end of CoverTab[93283]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1092
	_go_fuzz_dep_.CoverTab[93284]++

												e.cur = e.maxMatchOff

												dirtyShardCnt := 0
												if !allDirty {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1097
		_go_fuzz_dep_.CoverTab[93299]++
													for i := range e.longTableShardDirty {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1098
			_go_fuzz_dep_.CoverTab[93300]++
														if e.longTableShardDirty[i] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1099
				_go_fuzz_dep_.CoverTab[93301]++
															dirtyShardCnt++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1100
				// _ = "end of CoverTab[93301]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1101
				_go_fuzz_dep_.CoverTab[93302]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1101
				// _ = "end of CoverTab[93302]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1101
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1101
			// _ = "end of CoverTab[93300]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1102
		// _ = "end of CoverTab[93299]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1103
		_go_fuzz_dep_.CoverTab[93303]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1103
		// _ = "end of CoverTab[93303]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1103
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1103
	// _ = "end of CoverTab[93284]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1103
	_go_fuzz_dep_.CoverTab[93285]++

												if allDirty || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1105
		_go_fuzz_dep_.CoverTab[93304]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1105
		return dirtyShardCnt > dLongTableShardCnt/2
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1105
		// _ = "end of CoverTab[93304]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1105
	}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1105
		_go_fuzz_dep_.CoverTab[93305]++
													copy(e.longTable[:], e.dictLongTable)
													for i := range e.longTableShardDirty {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1107
			_go_fuzz_dep_.CoverTab[93307]++
														e.longTableShardDirty[i] = false
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1108
			// _ = "end of CoverTab[93307]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1109
		// _ = "end of CoverTab[93305]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1109
		_go_fuzz_dep_.CoverTab[93306]++
													return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1110
		// _ = "end of CoverTab[93306]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1111
		_go_fuzz_dep_.CoverTab[93308]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1111
		// _ = "end of CoverTab[93308]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1111
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1111
	// _ = "end of CoverTab[93285]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1111
	_go_fuzz_dep_.CoverTab[93286]++
												for i := range e.longTableShardDirty {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1112
		_go_fuzz_dep_.CoverTab[93309]++
													if !e.longTableShardDirty[i] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1113
			_go_fuzz_dep_.CoverTab[93311]++
														continue
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1114
			// _ = "end of CoverTab[93311]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1115
			_go_fuzz_dep_.CoverTab[93312]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1115
			// _ = "end of CoverTab[93312]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1115
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1115
		// _ = "end of CoverTab[93309]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1115
		_go_fuzz_dep_.CoverTab[93310]++

													copy(e.longTable[i*dLongTableShardSize:(i+1)*dLongTableShardSize], e.dictLongTable[i*dLongTableShardSize:(i+1)*dLongTableShardSize])
													e.longTableShardDirty[i] = false
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1118
		// _ = "end of CoverTab[93310]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1119
	// _ = "end of CoverTab[93286]"
}

func (e *doubleFastEncoderDict) markLongShardDirty(entryNum uint32) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1122
	_go_fuzz_dep_.CoverTab[93313]++
												e.longTableShardDirty[entryNum/dLongTableShardSize] = true
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1123
	// _ = "end of CoverTab[93313]"
}

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1124
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_dfast.go:1124
var _ = _go_fuzz_dep_.CoverTab
