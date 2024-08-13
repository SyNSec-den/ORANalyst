// Copyright 2019+ Klaus Post. All rights reserved.
// License information can be found in the LICENSE file.
// Based on work by Yann Collet, released under BSD License.

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:5
package zstd

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:5
)

import (
	"fmt"
)

const (
	tableBits		= 15					// Bits used in the table
	tableSize		= 1 << tableBits			// Size of the table
	tableShardCnt		= 1 << (tableBits - dictShardBits)	// Number of shards in the table
	tableShardSize		= tableSize / tableShardCnt		// Size of an individual shard
	tableFastHashLen	= 6
	tableMask		= tableSize - 1	// Mask for table indices. Redundant, but can eliminate bounds checks.
	maxMatchLength		= 131074
)

type tableEntry struct {
	val	uint32
	offset	int32
}

type fastEncoder struct {
	fastBase
	table	[tableSize]tableEntry
}

type fastEncoderDict struct {
	fastEncoder
	dictTable	[]tableEntry
	tableShardDirty	[tableShardCnt]bool
	allDirty	bool
}

// Encode mimmics functionality in zstd_fast.c
func (e *fastEncoder) Encode(blk *blockEnc, src []byte) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:39
	_go_fuzz_dep_.CoverTab[93314]++
												const (
		inputMargin		= 8
		minNonLiteralBlockSize	= 1 + 1 + inputMargin
	)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:46
	for e.cur >= bufferReset {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:46
		_go_fuzz_dep_.CoverTab[93321]++
													if len(e.hist) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:47
			_go_fuzz_dep_.CoverTab[93324]++
														for i := range e.table[:] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:48
				_go_fuzz_dep_.CoverTab[93326]++
															e.table[i] = tableEntry{}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:49
				// _ = "end of CoverTab[93326]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:50
			// _ = "end of CoverTab[93324]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:50
			_go_fuzz_dep_.CoverTab[93325]++
														e.cur = e.maxMatchOff
														break
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:52
			// _ = "end of CoverTab[93325]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:53
			_go_fuzz_dep_.CoverTab[93327]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:53
			// _ = "end of CoverTab[93327]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:53
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:53
		// _ = "end of CoverTab[93321]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:53
		_go_fuzz_dep_.CoverTab[93322]++

													minOff := e.cur + int32(len(e.hist)) - e.maxMatchOff
													for i := range e.table[:] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:56
			_go_fuzz_dep_.CoverTab[93328]++
														v := e.table[i].offset
														if v < minOff {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:58
				_go_fuzz_dep_.CoverTab[93330]++
															v = 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:59
				// _ = "end of CoverTab[93330]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:60
				_go_fuzz_dep_.CoverTab[93331]++
															v = v - e.cur + e.maxMatchOff
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:61
				// _ = "end of CoverTab[93331]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:62
			// _ = "end of CoverTab[93328]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:62
			_go_fuzz_dep_.CoverTab[93329]++
														e.table[i].offset = v
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:63
			// _ = "end of CoverTab[93329]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:64
		// _ = "end of CoverTab[93322]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:64
		_go_fuzz_dep_.CoverTab[93323]++
													e.cur = e.maxMatchOff
													break
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:66
		// _ = "end of CoverTab[93323]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:67
	// _ = "end of CoverTab[93314]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:67
	_go_fuzz_dep_.CoverTab[93315]++

												s := e.addBlock(src)
												blk.size = len(src)
												if len(src) < minNonLiteralBlockSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:71
		_go_fuzz_dep_.CoverTab[93332]++
													blk.extraLits = len(src)
													blk.literals = blk.literals[:len(src)]
													copy(blk.literals, src)
													return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:75
		// _ = "end of CoverTab[93332]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:76
		_go_fuzz_dep_.CoverTab[93333]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:76
		// _ = "end of CoverTab[93333]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:76
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:76
	// _ = "end of CoverTab[93315]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:76
	_go_fuzz_dep_.CoverTab[93316]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:79
	src = e.hist
	sLimit := int32(len(src)) - inputMargin
	// stepSize is the number of bytes to skip on every main loop iteration.
	// It should be >= 2.
	const stepSize = 2

	// TEMPLATE
	const hashLog = tableBits
												// seems global, but would be nice to tweak.
												const kSearchStrength = 7

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:91
	nextEmit := s
												cv := load6432(src, s)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:95
	offset1 := int32(blk.recentOffsets[0])
	offset2 := int32(blk.recentOffsets[1])

	addLiterals := func(s *seq, until int32) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:98
		_go_fuzz_dep_.CoverTab[93334]++
													if until == nextEmit {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:99
			_go_fuzz_dep_.CoverTab[93336]++
														return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:100
			// _ = "end of CoverTab[93336]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:101
			_go_fuzz_dep_.CoverTab[93337]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:101
			// _ = "end of CoverTab[93337]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:101
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:101
		// _ = "end of CoverTab[93334]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:101
		_go_fuzz_dep_.CoverTab[93335]++
													blk.literals = append(blk.literals, src[nextEmit:until]...)
													s.litLen = uint32(until - nextEmit)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:103
		// _ = "end of CoverTab[93335]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:104
	// _ = "end of CoverTab[93316]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:104
	_go_fuzz_dep_.CoverTab[93317]++
												if debugEncoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:105
		_go_fuzz_dep_.CoverTab[93338]++
													println("recent offsets:", blk.recentOffsets)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:106
		// _ = "end of CoverTab[93338]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:107
		_go_fuzz_dep_.CoverTab[93339]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:107
		// _ = "end of CoverTab[93339]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:107
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:107
	// _ = "end of CoverTab[93317]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:107
	_go_fuzz_dep_.CoverTab[93318]++

encodeLoop:
	for {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:110
		_go_fuzz_dep_.CoverTab[93340]++
		// t will contain the match offset when we find one.
													// When existing the search loop, we have already checked 4 bytes.
													var t int32

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:117
		canRepeat := len(blk.sequences) > 2

		for {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:119
			_go_fuzz_dep_.CoverTab[93349]++
														if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:120
				_go_fuzz_dep_.CoverTab[93355]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:120
				return canRepeat
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:120
				// _ = "end of CoverTab[93355]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:120
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:120
				_go_fuzz_dep_.CoverTab[93356]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:120
				return offset1 == 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:120
				// _ = "end of CoverTab[93356]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:120
			}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:120
				_go_fuzz_dep_.CoverTab[93357]++
															panic("offset0 was 0")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:121
				// _ = "end of CoverTab[93357]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:122
				_go_fuzz_dep_.CoverTab[93358]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:122
				// _ = "end of CoverTab[93358]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:122
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:122
			// _ = "end of CoverTab[93349]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:122
			_go_fuzz_dep_.CoverTab[93350]++

														nextHash := hashLen(cv, hashLog, tableFastHashLen)
														nextHash2 := hashLen(cv>>8, hashLog, tableFastHashLen)
														candidate := e.table[nextHash]
														candidate2 := e.table[nextHash2]
														repIndex := s - offset1 + 2

														e.table[nextHash] = tableEntry{offset: s + e.cur, val: uint32(cv)}
														e.table[nextHash2] = tableEntry{offset: s + e.cur + 1, val: uint32(cv >> 8)}

														if canRepeat && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:133
				_go_fuzz_dep_.CoverTab[93359]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:133
				return repIndex >= 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:133
				// _ = "end of CoverTab[93359]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:133
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:133
				_go_fuzz_dep_.CoverTab[93360]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:133
				return load3232(src, repIndex) == uint32(cv>>16)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:133
				// _ = "end of CoverTab[93360]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:133
			}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:133
				_go_fuzz_dep_.CoverTab[93361]++
															// Consider history as well.
															var seq seq
															var length int32
															length = 4 + e.matchlen(s+6, repIndex+4, src)
															seq.matchLen = uint32(length - zstdMinMatch)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:142
				start := s + 2

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:145
				startLimit := nextEmit + 1

				sMin := s - e.maxMatchOff
				if sMin < 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:148
					_go_fuzz_dep_.CoverTab[93366]++
																sMin = 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:149
					// _ = "end of CoverTab[93366]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:150
					_go_fuzz_dep_.CoverTab[93367]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:150
					// _ = "end of CoverTab[93367]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:150
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:150
				// _ = "end of CoverTab[93361]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:150
				_go_fuzz_dep_.CoverTab[93362]++
															for repIndex > sMin && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:151
					_go_fuzz_dep_.CoverTab[93368]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:151
					return start > startLimit
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:151
					// _ = "end of CoverTab[93368]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:151
				}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:151
					_go_fuzz_dep_.CoverTab[93369]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:151
					return src[repIndex-1] == src[start-1]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:151
					// _ = "end of CoverTab[93369]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:151
				}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:151
					_go_fuzz_dep_.CoverTab[93370]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:151
					return seq.matchLen < maxMatchLength-zstdMinMatch
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:151
					// _ = "end of CoverTab[93370]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:151
				}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:151
					_go_fuzz_dep_.CoverTab[93371]++
																repIndex--
																start--
																seq.matchLen++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:154
					// _ = "end of CoverTab[93371]"
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:155
				// _ = "end of CoverTab[93362]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:155
				_go_fuzz_dep_.CoverTab[93363]++
															addLiterals(&seq, start)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:159
				seq.offset = 1
				if debugSequences {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:160
					_go_fuzz_dep_.CoverTab[93372]++
																println("repeat sequence", seq, "next s:", s)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:161
					// _ = "end of CoverTab[93372]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:162
					_go_fuzz_dep_.CoverTab[93373]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:162
					// _ = "end of CoverTab[93373]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:162
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:162
				// _ = "end of CoverTab[93363]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:162
				_go_fuzz_dep_.CoverTab[93364]++
															blk.sequences = append(blk.sequences, seq)
															s += length + 2
															nextEmit = s
															if s >= sLimit {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:166
					_go_fuzz_dep_.CoverTab[93374]++
																if debugEncoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:167
						_go_fuzz_dep_.CoverTab[93376]++
																	println("repeat ended", s, length)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:168
						// _ = "end of CoverTab[93376]"

					} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:170
						_go_fuzz_dep_.CoverTab[93377]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:170
						// _ = "end of CoverTab[93377]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:170
					}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:170
					// _ = "end of CoverTab[93374]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:170
					_go_fuzz_dep_.CoverTab[93375]++
																break encodeLoop
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:171
					// _ = "end of CoverTab[93375]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:172
					_go_fuzz_dep_.CoverTab[93378]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:172
					// _ = "end of CoverTab[93378]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:172
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:172
				// _ = "end of CoverTab[93364]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:172
				_go_fuzz_dep_.CoverTab[93365]++
															cv = load6432(src, s)
															continue
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:174
				// _ = "end of CoverTab[93365]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:175
				_go_fuzz_dep_.CoverTab[93379]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:175
				// _ = "end of CoverTab[93379]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:175
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:175
			// _ = "end of CoverTab[93350]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:175
			_go_fuzz_dep_.CoverTab[93351]++
														coffset0 := s - (candidate.offset - e.cur)
														coffset1 := s - (candidate2.offset - e.cur) + 1
														if coffset0 < e.maxMatchOff && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:178
				_go_fuzz_dep_.CoverTab[93380]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:178
				return uint32(cv) == candidate.val
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:178
				// _ = "end of CoverTab[93380]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:178
			}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:178
				_go_fuzz_dep_.CoverTab[93381]++

															t = candidate.offset - e.cur
															if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:181
					_go_fuzz_dep_.CoverTab[93384]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:181
					return s <= t
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:181
					// _ = "end of CoverTab[93384]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:181
				}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:181
					_go_fuzz_dep_.CoverTab[93385]++
																panic(fmt.Sprintf("s (%d) <= t (%d)", s, t))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:182
					// _ = "end of CoverTab[93385]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:183
					_go_fuzz_dep_.CoverTab[93386]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:183
					// _ = "end of CoverTab[93386]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:183
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:183
				// _ = "end of CoverTab[93381]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:183
				_go_fuzz_dep_.CoverTab[93382]++
															if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:184
					_go_fuzz_dep_.CoverTab[93387]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:184
					return s-t > e.maxMatchOff
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:184
					// _ = "end of CoverTab[93387]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:184
				}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:184
					_go_fuzz_dep_.CoverTab[93388]++
																panic("s - t >e.maxMatchOff")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:185
					// _ = "end of CoverTab[93388]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:186
					_go_fuzz_dep_.CoverTab[93389]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:186
					// _ = "end of CoverTab[93389]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:186
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:186
				// _ = "end of CoverTab[93382]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:186
				_go_fuzz_dep_.CoverTab[93383]++
															break
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:187
				// _ = "end of CoverTab[93383]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:188
				_go_fuzz_dep_.CoverTab[93390]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:188
				// _ = "end of CoverTab[93390]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:188
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:188
			// _ = "end of CoverTab[93351]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:188
			_go_fuzz_dep_.CoverTab[93352]++

														if coffset1 < e.maxMatchOff && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:190
				_go_fuzz_dep_.CoverTab[93391]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:190
				return uint32(cv>>8) == candidate2.val
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:190
				// _ = "end of CoverTab[93391]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:190
			}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:190
				_go_fuzz_dep_.CoverTab[93392]++

															t = candidate2.offset - e.cur
															s++
															if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:194
					_go_fuzz_dep_.CoverTab[93396]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:194
					return s <= t
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:194
					// _ = "end of CoverTab[93396]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:194
				}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:194
					_go_fuzz_dep_.CoverTab[93397]++
																panic(fmt.Sprintf("s (%d) <= t (%d)", s, t))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:195
					// _ = "end of CoverTab[93397]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:196
					_go_fuzz_dep_.CoverTab[93398]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:196
					// _ = "end of CoverTab[93398]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:196
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:196
				// _ = "end of CoverTab[93392]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:196
				_go_fuzz_dep_.CoverTab[93393]++
															if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:197
					_go_fuzz_dep_.CoverTab[93399]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:197
					return s-t > e.maxMatchOff
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:197
					// _ = "end of CoverTab[93399]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:197
				}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:197
					_go_fuzz_dep_.CoverTab[93400]++
																panic("s - t >e.maxMatchOff")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:198
					// _ = "end of CoverTab[93400]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:199
					_go_fuzz_dep_.CoverTab[93401]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:199
					// _ = "end of CoverTab[93401]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:199
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:199
				// _ = "end of CoverTab[93393]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:199
				_go_fuzz_dep_.CoverTab[93394]++
															if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:200
					_go_fuzz_dep_.CoverTab[93402]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:200
					return t < 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:200
					// _ = "end of CoverTab[93402]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:200
				}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:200
					_go_fuzz_dep_.CoverTab[93403]++
																panic("t<0")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:201
					// _ = "end of CoverTab[93403]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:202
					_go_fuzz_dep_.CoverTab[93404]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:202
					// _ = "end of CoverTab[93404]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:202
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:202
				// _ = "end of CoverTab[93394]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:202
				_go_fuzz_dep_.CoverTab[93395]++
															break
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:203
				// _ = "end of CoverTab[93395]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:204
				_go_fuzz_dep_.CoverTab[93405]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:204
				// _ = "end of CoverTab[93405]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:204
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:204
			// _ = "end of CoverTab[93352]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:204
			_go_fuzz_dep_.CoverTab[93353]++
														s += stepSize + ((s - nextEmit) >> (kSearchStrength - 1))
														if s >= sLimit {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:206
				_go_fuzz_dep_.CoverTab[93406]++
															break encodeLoop
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:207
				// _ = "end of CoverTab[93406]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:208
				_go_fuzz_dep_.CoverTab[93407]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:208
				// _ = "end of CoverTab[93407]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:208
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:208
			// _ = "end of CoverTab[93353]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:208
			_go_fuzz_dep_.CoverTab[93354]++
														cv = load6432(src, s)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:209
			// _ = "end of CoverTab[93354]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:210
		// _ = "end of CoverTab[93340]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:210
		_go_fuzz_dep_.CoverTab[93341]++

													offset2 = offset1
													offset1 = s - t

													if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:215
			_go_fuzz_dep_.CoverTab[93408]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:215
			return s <= t
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:215
			// _ = "end of CoverTab[93408]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:215
		}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:215
			_go_fuzz_dep_.CoverTab[93409]++
														panic(fmt.Sprintf("s (%d) <= t (%d)", s, t))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:216
			// _ = "end of CoverTab[93409]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:217
			_go_fuzz_dep_.CoverTab[93410]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:217
			// _ = "end of CoverTab[93410]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:217
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:217
		// _ = "end of CoverTab[93341]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:217
		_go_fuzz_dep_.CoverTab[93342]++

													if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:219
			_go_fuzz_dep_.CoverTab[93411]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:219
			return canRepeat
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:219
			// _ = "end of CoverTab[93411]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:219
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:219
			_go_fuzz_dep_.CoverTab[93412]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:219
			return int(offset1) > len(src)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:219
			// _ = "end of CoverTab[93412]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:219
		}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:219
			_go_fuzz_dep_.CoverTab[93413]++
														panic("invalid offset")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:220
			// _ = "end of CoverTab[93413]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:221
			_go_fuzz_dep_.CoverTab[93414]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:221
			// _ = "end of CoverTab[93414]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:221
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:221
		// _ = "end of CoverTab[93342]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:221
		_go_fuzz_dep_.CoverTab[93343]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:224
		l := e.matchlen(s+4, t+4, src) + 4

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:227
		tMin := s - e.maxMatchOff
		if tMin < 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:228
			_go_fuzz_dep_.CoverTab[93415]++
														tMin = 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:229
			// _ = "end of CoverTab[93415]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:230
			_go_fuzz_dep_.CoverTab[93416]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:230
			// _ = "end of CoverTab[93416]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:230
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:230
		// _ = "end of CoverTab[93343]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:230
		_go_fuzz_dep_.CoverTab[93344]++
													for t > tMin && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:231
			_go_fuzz_dep_.CoverTab[93417]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:231
			return s > nextEmit
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:231
			// _ = "end of CoverTab[93417]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:231
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:231
			_go_fuzz_dep_.CoverTab[93418]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:231
			return src[t-1] == src[s-1]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:231
			// _ = "end of CoverTab[93418]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:231
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:231
			_go_fuzz_dep_.CoverTab[93419]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:231
			return l < maxMatchLength
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:231
			// _ = "end of CoverTab[93419]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:231
		}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:231
			_go_fuzz_dep_.CoverTab[93420]++
														s--
														t--
														l++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:234
			// _ = "end of CoverTab[93420]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:235
		// _ = "end of CoverTab[93344]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:235
		_go_fuzz_dep_.CoverTab[93345]++

		// Write our sequence.
		var seq seq
		seq.litLen = uint32(s - nextEmit)
		seq.matchLen = uint32(l - zstdMinMatch)
		if seq.litLen > 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:241
			_go_fuzz_dep_.CoverTab[93421]++
														blk.literals = append(blk.literals, src[nextEmit:s]...)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:242
			// _ = "end of CoverTab[93421]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:243
			_go_fuzz_dep_.CoverTab[93422]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:243
			// _ = "end of CoverTab[93422]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:243
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:243
		// _ = "end of CoverTab[93345]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:243
		_go_fuzz_dep_.CoverTab[93346]++

													seq.offset = uint32(s-t) + 3
													s += l
													if debugSequences {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:247
			_go_fuzz_dep_.CoverTab[93423]++
														println("sequence", seq, "next s:", s)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:248
			// _ = "end of CoverTab[93423]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:249
			_go_fuzz_dep_.CoverTab[93424]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:249
			// _ = "end of CoverTab[93424]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:249
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:249
		// _ = "end of CoverTab[93346]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:249
		_go_fuzz_dep_.CoverTab[93347]++
													blk.sequences = append(blk.sequences, seq)
													nextEmit = s
													if s >= sLimit {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:252
			_go_fuzz_dep_.CoverTab[93425]++
														break encodeLoop
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:253
			// _ = "end of CoverTab[93425]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:254
			_go_fuzz_dep_.CoverTab[93426]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:254
			// _ = "end of CoverTab[93426]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:254
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:254
		// _ = "end of CoverTab[93347]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:254
		_go_fuzz_dep_.CoverTab[93348]++
													cv = load6432(src, s)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:258
		if o2 := s - offset2; canRepeat && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:258
			_go_fuzz_dep_.CoverTab[93427]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:258
			return load3232(src, o2) == uint32(cv)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:258
			// _ = "end of CoverTab[93427]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:258
		}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:258
			_go_fuzz_dep_.CoverTab[93428]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:261
			l := 4 + e.matchlen(s+4, o2+4, src)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:264
			nextHash := hashLen(cv, hashLog, tableFastHashLen)
			e.table[nextHash] = tableEntry{offset: s + e.cur, val: uint32(cv)}
			seq.matchLen = uint32(l) - zstdMinMatch
			seq.litLen = 0

			seq.offset = 1
			s += l
			nextEmit = s
			if debugSequences {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:272
				_go_fuzz_dep_.CoverTab[93431]++
															println("sequence", seq, "next s:", s)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:273
				// _ = "end of CoverTab[93431]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:274
				_go_fuzz_dep_.CoverTab[93432]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:274
				// _ = "end of CoverTab[93432]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:274
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:274
			// _ = "end of CoverTab[93428]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:274
			_go_fuzz_dep_.CoverTab[93429]++
														blk.sequences = append(blk.sequences, seq)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:278
			offset1, offset2 = offset2, offset1
			if s >= sLimit {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:279
				_go_fuzz_dep_.CoverTab[93433]++
															break encodeLoop
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:280
				// _ = "end of CoverTab[93433]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:281
				_go_fuzz_dep_.CoverTab[93434]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:281
				// _ = "end of CoverTab[93434]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:281
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:281
			// _ = "end of CoverTab[93429]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:281
			_go_fuzz_dep_.CoverTab[93430]++

														cv = load6432(src, s)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:283
			// _ = "end of CoverTab[93430]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:284
			_go_fuzz_dep_.CoverTab[93435]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:284
			// _ = "end of CoverTab[93435]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:284
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:284
		// _ = "end of CoverTab[93348]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:285
	// _ = "end of CoverTab[93318]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:285
	_go_fuzz_dep_.CoverTab[93319]++

												if int(nextEmit) < len(src) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:287
		_go_fuzz_dep_.CoverTab[93436]++
													blk.literals = append(blk.literals, src[nextEmit:]...)
													blk.extraLits = len(src) - int(nextEmit)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:289
		// _ = "end of CoverTab[93436]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:290
		_go_fuzz_dep_.CoverTab[93437]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:290
		// _ = "end of CoverTab[93437]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:290
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:290
	// _ = "end of CoverTab[93319]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:290
	_go_fuzz_dep_.CoverTab[93320]++
												blk.recentOffsets[0] = uint32(offset1)
												blk.recentOffsets[1] = uint32(offset2)
												if debugEncoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:293
		_go_fuzz_dep_.CoverTab[93438]++
													println("returning, recent offsets:", blk.recentOffsets, "extra literals:", blk.extraLits)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:294
		// _ = "end of CoverTab[93438]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:295
		_go_fuzz_dep_.CoverTab[93439]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:295
		// _ = "end of CoverTab[93439]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:295
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:295
	// _ = "end of CoverTab[93320]"
}

// EncodeNoHist will encode a block with no history and no following blocks.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:298
// Most notable difference is that src will not be copied for history and
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:298
// we do not need to check for max match length.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:301
func (e *fastEncoder) EncodeNoHist(blk *blockEnc, src []byte) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:301
	_go_fuzz_dep_.CoverTab[93440]++
												const (
		inputMargin		= 8
		minNonLiteralBlockSize	= 1 + 1 + inputMargin
	)
	if debugEncoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:306
		_go_fuzz_dep_.CoverTab[93449]++
													if len(src) > maxBlockSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:307
			_go_fuzz_dep_.CoverTab[93450]++
														panic("src too big")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:308
			// _ = "end of CoverTab[93450]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:309
			_go_fuzz_dep_.CoverTab[93451]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:309
			// _ = "end of CoverTab[93451]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:309
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:309
		// _ = "end of CoverTab[93449]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:310
		_go_fuzz_dep_.CoverTab[93452]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:310
		// _ = "end of CoverTab[93452]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:310
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:310
	// _ = "end of CoverTab[93440]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:310
	_go_fuzz_dep_.CoverTab[93441]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:313
	if e.cur >= bufferReset {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:313
		_go_fuzz_dep_.CoverTab[93453]++
													for i := range e.table[:] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:314
			_go_fuzz_dep_.CoverTab[93455]++
														e.table[i] = tableEntry{}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:315
			// _ = "end of CoverTab[93455]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:316
		// _ = "end of CoverTab[93453]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:316
		_go_fuzz_dep_.CoverTab[93454]++
													e.cur = e.maxMatchOff
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:317
		// _ = "end of CoverTab[93454]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:318
		_go_fuzz_dep_.CoverTab[93456]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:318
		// _ = "end of CoverTab[93456]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:318
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:318
	// _ = "end of CoverTab[93441]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:318
	_go_fuzz_dep_.CoverTab[93442]++

												s := int32(0)
												blk.size = len(src)
												if len(src) < minNonLiteralBlockSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:322
		_go_fuzz_dep_.CoverTab[93457]++
													blk.extraLits = len(src)
													blk.literals = blk.literals[:len(src)]
													copy(blk.literals, src)
													return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:326
		// _ = "end of CoverTab[93457]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:327
		_go_fuzz_dep_.CoverTab[93458]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:327
		// _ = "end of CoverTab[93458]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:327
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:327
	// _ = "end of CoverTab[93442]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:327
	_go_fuzz_dep_.CoverTab[93443]++

												sLimit := int32(len(src)) - inputMargin
	// stepSize is the number of bytes to skip on every main loop iteration.
	// It should be >= 2.
	const stepSize = 2

	// TEMPLATE
	const hashLog = tableBits
												// seems global, but would be nice to tweak.
												const kSearchStrength = 8

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:340
	nextEmit := s
												cv := load6432(src, s)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:344
	offset1 := int32(blk.recentOffsets[0])
	offset2 := int32(blk.recentOffsets[1])

	addLiterals := func(s *seq, until int32) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:347
		_go_fuzz_dep_.CoverTab[93459]++
													if until == nextEmit {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:348
			_go_fuzz_dep_.CoverTab[93461]++
														return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:349
			// _ = "end of CoverTab[93461]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:350
			_go_fuzz_dep_.CoverTab[93462]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:350
			// _ = "end of CoverTab[93462]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:350
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:350
		// _ = "end of CoverTab[93459]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:350
		_go_fuzz_dep_.CoverTab[93460]++
													blk.literals = append(blk.literals, src[nextEmit:until]...)
													s.litLen = uint32(until - nextEmit)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:352
		// _ = "end of CoverTab[93460]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:353
	// _ = "end of CoverTab[93443]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:353
	_go_fuzz_dep_.CoverTab[93444]++
												if debugEncoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:354
		_go_fuzz_dep_.CoverTab[93463]++
													println("recent offsets:", blk.recentOffsets)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:355
		// _ = "end of CoverTab[93463]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:356
		_go_fuzz_dep_.CoverTab[93464]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:356
		// _ = "end of CoverTab[93464]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:356
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:356
	// _ = "end of CoverTab[93444]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:356
	_go_fuzz_dep_.CoverTab[93445]++

encodeLoop:
	for {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:359
		_go_fuzz_dep_.CoverTab[93465]++
		// t will contain the match offset when we find one.
													// When existing the search loop, we have already checked 4 bytes.
													var t int32

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:367
		for {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:367
			_go_fuzz_dep_.CoverTab[93474]++
														nextHash := hashLen(cv, hashLog, tableFastHashLen)
														nextHash2 := hashLen(cv>>8, hashLog, tableFastHashLen)
														candidate := e.table[nextHash]
														candidate2 := e.table[nextHash2]
														repIndex := s - offset1 + 2

														e.table[nextHash] = tableEntry{offset: s + e.cur, val: uint32(cv)}
														e.table[nextHash2] = tableEntry{offset: s + e.cur + 1, val: uint32(cv >> 8)}

														if len(blk.sequences) > 2 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:377
				_go_fuzz_dep_.CoverTab[93479]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:377
				return load3232(src, repIndex) == uint32(cv>>16)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:377
				// _ = "end of CoverTab[93479]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:377
			}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:377
				_go_fuzz_dep_.CoverTab[93480]++
															// Consider history as well.
															var seq seq
															length := 4 + e.matchlen(s+6, repIndex+4, src)

															seq.matchLen = uint32(length - zstdMinMatch)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:386
				start := s + 2

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:389
				startLimit := nextEmit + 1

				sMin := s - e.maxMatchOff
				if sMin < 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:392
					_go_fuzz_dep_.CoverTab[93485]++
																sMin = 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:393
					// _ = "end of CoverTab[93485]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:394
					_go_fuzz_dep_.CoverTab[93486]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:394
					// _ = "end of CoverTab[93486]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:394
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:394
				// _ = "end of CoverTab[93480]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:394
				_go_fuzz_dep_.CoverTab[93481]++
															for repIndex > sMin && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:395
					_go_fuzz_dep_.CoverTab[93487]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:395
					return start > startLimit
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:395
					// _ = "end of CoverTab[93487]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:395
				}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:395
					_go_fuzz_dep_.CoverTab[93488]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:395
					return src[repIndex-1] == src[start-1]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:395
					// _ = "end of CoverTab[93488]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:395
				}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:395
					_go_fuzz_dep_.CoverTab[93489]++
																repIndex--
																start--
																seq.matchLen++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:398
					// _ = "end of CoverTab[93489]"
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:399
				// _ = "end of CoverTab[93481]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:399
				_go_fuzz_dep_.CoverTab[93482]++
															addLiterals(&seq, start)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:403
				seq.offset = 1
				if debugSequences {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:404
					_go_fuzz_dep_.CoverTab[93490]++
																println("repeat sequence", seq, "next s:", s)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:405
					// _ = "end of CoverTab[93490]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:406
					_go_fuzz_dep_.CoverTab[93491]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:406
					// _ = "end of CoverTab[93491]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:406
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:406
				// _ = "end of CoverTab[93482]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:406
				_go_fuzz_dep_.CoverTab[93483]++
															blk.sequences = append(blk.sequences, seq)
															s += length + 2
															nextEmit = s
															if s >= sLimit {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:410
					_go_fuzz_dep_.CoverTab[93492]++
																if debugEncoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:411
						_go_fuzz_dep_.CoverTab[93494]++
																	println("repeat ended", s, length)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:412
						// _ = "end of CoverTab[93494]"

					} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:414
						_go_fuzz_dep_.CoverTab[93495]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:414
						// _ = "end of CoverTab[93495]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:414
					}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:414
					// _ = "end of CoverTab[93492]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:414
					_go_fuzz_dep_.CoverTab[93493]++
																break encodeLoop
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:415
					// _ = "end of CoverTab[93493]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:416
					_go_fuzz_dep_.CoverTab[93496]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:416
					// _ = "end of CoverTab[93496]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:416
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:416
				// _ = "end of CoverTab[93483]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:416
				_go_fuzz_dep_.CoverTab[93484]++
															cv = load6432(src, s)
															continue
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:418
				// _ = "end of CoverTab[93484]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:419
				_go_fuzz_dep_.CoverTab[93497]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:419
				// _ = "end of CoverTab[93497]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:419
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:419
			// _ = "end of CoverTab[93474]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:419
			_go_fuzz_dep_.CoverTab[93475]++
														coffset0 := s - (candidate.offset - e.cur)
														coffset1 := s - (candidate2.offset - e.cur) + 1
														if coffset0 < e.maxMatchOff && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:422
				_go_fuzz_dep_.CoverTab[93498]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:422
				return uint32(cv) == candidate.val
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:422
				// _ = "end of CoverTab[93498]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:422
			}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:422
				_go_fuzz_dep_.CoverTab[93499]++

															t = candidate.offset - e.cur
															if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:425
					_go_fuzz_dep_.CoverTab[93503]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:425
					return s <= t
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:425
					// _ = "end of CoverTab[93503]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:425
				}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:425
					_go_fuzz_dep_.CoverTab[93504]++
																panic(fmt.Sprintf("s (%d) <= t (%d)", s, t))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:426
					// _ = "end of CoverTab[93504]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:427
					_go_fuzz_dep_.CoverTab[93505]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:427
					// _ = "end of CoverTab[93505]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:427
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:427
				// _ = "end of CoverTab[93499]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:427
				_go_fuzz_dep_.CoverTab[93500]++
															if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:428
					_go_fuzz_dep_.CoverTab[93506]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:428
					return s-t > e.maxMatchOff
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:428
					// _ = "end of CoverTab[93506]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:428
				}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:428
					_go_fuzz_dep_.CoverTab[93507]++
																panic("s - t >e.maxMatchOff")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:429
					// _ = "end of CoverTab[93507]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:430
					_go_fuzz_dep_.CoverTab[93508]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:430
					// _ = "end of CoverTab[93508]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:430
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:430
				// _ = "end of CoverTab[93500]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:430
				_go_fuzz_dep_.CoverTab[93501]++
															if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:431
					_go_fuzz_dep_.CoverTab[93509]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:431
					return t < 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:431
					// _ = "end of CoverTab[93509]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:431
				}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:431
					_go_fuzz_dep_.CoverTab[93510]++
																panic(fmt.Sprintf("t (%d) < 0, candidate.offset: %d, e.cur: %d, coffset0: %d, e.maxMatchOff: %d", t, candidate.offset, e.cur, coffset0, e.maxMatchOff))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:432
					// _ = "end of CoverTab[93510]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:433
					_go_fuzz_dep_.CoverTab[93511]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:433
					// _ = "end of CoverTab[93511]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:433
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:433
				// _ = "end of CoverTab[93501]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:433
				_go_fuzz_dep_.CoverTab[93502]++
															break
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:434
				// _ = "end of CoverTab[93502]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:435
				_go_fuzz_dep_.CoverTab[93512]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:435
				// _ = "end of CoverTab[93512]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:435
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:435
			// _ = "end of CoverTab[93475]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:435
			_go_fuzz_dep_.CoverTab[93476]++

														if coffset1 < e.maxMatchOff && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:437
				_go_fuzz_dep_.CoverTab[93513]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:437
				return uint32(cv>>8) == candidate2.val
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:437
				// _ = "end of CoverTab[93513]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:437
			}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:437
				_go_fuzz_dep_.CoverTab[93514]++

															t = candidate2.offset - e.cur
															s++
															if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:441
					_go_fuzz_dep_.CoverTab[93518]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:441
					return s <= t
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:441
					// _ = "end of CoverTab[93518]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:441
				}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:441
					_go_fuzz_dep_.CoverTab[93519]++
																panic(fmt.Sprintf("s (%d) <= t (%d)", s, t))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:442
					// _ = "end of CoverTab[93519]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:443
					_go_fuzz_dep_.CoverTab[93520]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:443
					// _ = "end of CoverTab[93520]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:443
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:443
				// _ = "end of CoverTab[93514]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:443
				_go_fuzz_dep_.CoverTab[93515]++
															if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:444
					_go_fuzz_dep_.CoverTab[93521]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:444
					return s-t > e.maxMatchOff
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:444
					// _ = "end of CoverTab[93521]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:444
				}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:444
					_go_fuzz_dep_.CoverTab[93522]++
																panic("s - t >e.maxMatchOff")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:445
					// _ = "end of CoverTab[93522]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:446
					_go_fuzz_dep_.CoverTab[93523]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:446
					// _ = "end of CoverTab[93523]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:446
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:446
				// _ = "end of CoverTab[93515]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:446
				_go_fuzz_dep_.CoverTab[93516]++
															if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:447
					_go_fuzz_dep_.CoverTab[93524]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:447
					return t < 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:447
					// _ = "end of CoverTab[93524]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:447
				}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:447
					_go_fuzz_dep_.CoverTab[93525]++
																panic("t<0")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:448
					// _ = "end of CoverTab[93525]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:449
					_go_fuzz_dep_.CoverTab[93526]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:449
					// _ = "end of CoverTab[93526]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:449
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:449
				// _ = "end of CoverTab[93516]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:449
				_go_fuzz_dep_.CoverTab[93517]++
															break
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:450
				// _ = "end of CoverTab[93517]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:451
				_go_fuzz_dep_.CoverTab[93527]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:451
				// _ = "end of CoverTab[93527]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:451
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:451
			// _ = "end of CoverTab[93476]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:451
			_go_fuzz_dep_.CoverTab[93477]++
														s += stepSize + ((s - nextEmit) >> (kSearchStrength - 1))
														if s >= sLimit {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:453
				_go_fuzz_dep_.CoverTab[93528]++
															break encodeLoop
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:454
				// _ = "end of CoverTab[93528]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:455
				_go_fuzz_dep_.CoverTab[93529]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:455
				// _ = "end of CoverTab[93529]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:455
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:455
			// _ = "end of CoverTab[93477]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:455
			_go_fuzz_dep_.CoverTab[93478]++
														cv = load6432(src, s)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:456
			// _ = "end of CoverTab[93478]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:457
		// _ = "end of CoverTab[93465]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:457
		_go_fuzz_dep_.CoverTab[93466]++

													offset2 = offset1
													offset1 = s - t

													if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:462
			_go_fuzz_dep_.CoverTab[93530]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:462
			return s <= t
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:462
			// _ = "end of CoverTab[93530]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:462
		}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:462
			_go_fuzz_dep_.CoverTab[93531]++
														panic(fmt.Sprintf("s (%d) <= t (%d)", s, t))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:463
			// _ = "end of CoverTab[93531]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:464
			_go_fuzz_dep_.CoverTab[93532]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:464
			// _ = "end of CoverTab[93532]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:464
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:464
		// _ = "end of CoverTab[93466]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:464
		_go_fuzz_dep_.CoverTab[93467]++

													if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:466
			_go_fuzz_dep_.CoverTab[93533]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:466
			return t < 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:466
			// _ = "end of CoverTab[93533]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:466
		}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:466
			_go_fuzz_dep_.CoverTab[93534]++
														panic(fmt.Sprintf("t (%d) < 0 ", t))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:467
			// _ = "end of CoverTab[93534]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:468
			_go_fuzz_dep_.CoverTab[93535]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:468
			// _ = "end of CoverTab[93535]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:468
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:468
		// _ = "end of CoverTab[93467]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:468
		_go_fuzz_dep_.CoverTab[93468]++

													l := e.matchlen(s+4, t+4, src) + 4

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:473
		tMin := s - e.maxMatchOff
		if tMin < 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:474
			_go_fuzz_dep_.CoverTab[93536]++
														tMin = 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:475
			// _ = "end of CoverTab[93536]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:476
			_go_fuzz_dep_.CoverTab[93537]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:476
			// _ = "end of CoverTab[93537]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:476
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:476
		// _ = "end of CoverTab[93468]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:476
		_go_fuzz_dep_.CoverTab[93469]++
													for t > tMin && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:477
			_go_fuzz_dep_.CoverTab[93538]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:477
			return s > nextEmit
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:477
			// _ = "end of CoverTab[93538]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:477
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:477
			_go_fuzz_dep_.CoverTab[93539]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:477
			return src[t-1] == src[s-1]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:477
			// _ = "end of CoverTab[93539]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:477
		}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:477
			_go_fuzz_dep_.CoverTab[93540]++
														s--
														t--
														l++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:480
			// _ = "end of CoverTab[93540]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:481
		// _ = "end of CoverTab[93469]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:481
		_go_fuzz_dep_.CoverTab[93470]++

		// Write our sequence.
		var seq seq
		seq.litLen = uint32(s - nextEmit)
		seq.matchLen = uint32(l - zstdMinMatch)
		if seq.litLen > 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:487
			_go_fuzz_dep_.CoverTab[93541]++
														blk.literals = append(blk.literals, src[nextEmit:s]...)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:488
			// _ = "end of CoverTab[93541]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:489
			_go_fuzz_dep_.CoverTab[93542]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:489
			// _ = "end of CoverTab[93542]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:489
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:489
		// _ = "end of CoverTab[93470]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:489
		_go_fuzz_dep_.CoverTab[93471]++

													seq.offset = uint32(s-t) + 3
													s += l
													if debugSequences {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:493
			_go_fuzz_dep_.CoverTab[93543]++
														println("sequence", seq, "next s:", s)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:494
			// _ = "end of CoverTab[93543]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:495
			_go_fuzz_dep_.CoverTab[93544]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:495
			// _ = "end of CoverTab[93544]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:495
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:495
		// _ = "end of CoverTab[93471]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:495
		_go_fuzz_dep_.CoverTab[93472]++
													blk.sequences = append(blk.sequences, seq)
													nextEmit = s
													if s >= sLimit {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:498
			_go_fuzz_dep_.CoverTab[93545]++
														break encodeLoop
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:499
			// _ = "end of CoverTab[93545]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:500
			_go_fuzz_dep_.CoverTab[93546]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:500
			// _ = "end of CoverTab[93546]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:500
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:500
		// _ = "end of CoverTab[93472]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:500
		_go_fuzz_dep_.CoverTab[93473]++
													cv = load6432(src, s)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:504
		if o2 := s - offset2; len(blk.sequences) > 2 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:504
			_go_fuzz_dep_.CoverTab[93547]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:504
			return load3232(src, o2) == uint32(cv)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:504
			// _ = "end of CoverTab[93547]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:504
		}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:504
			_go_fuzz_dep_.CoverTab[93548]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:507
			l := 4 + e.matchlen(s+4, o2+4, src)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:510
			nextHash := hashLen(cv, hashLog, tableFastHashLen)
			e.table[nextHash] = tableEntry{offset: s + e.cur, val: uint32(cv)}
			seq.matchLen = uint32(l) - zstdMinMatch
			seq.litLen = 0

			seq.offset = 1
			s += l
			nextEmit = s
			if debugSequences {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:518
				_go_fuzz_dep_.CoverTab[93551]++
															println("sequence", seq, "next s:", s)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:519
				// _ = "end of CoverTab[93551]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:520
				_go_fuzz_dep_.CoverTab[93552]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:520
				// _ = "end of CoverTab[93552]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:520
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:520
			// _ = "end of CoverTab[93548]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:520
			_go_fuzz_dep_.CoverTab[93549]++
														blk.sequences = append(blk.sequences, seq)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:524
			offset1, offset2 = offset2, offset1
			if s >= sLimit {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:525
				_go_fuzz_dep_.CoverTab[93553]++
															break encodeLoop
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:526
				// _ = "end of CoverTab[93553]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:527
				_go_fuzz_dep_.CoverTab[93554]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:527
				// _ = "end of CoverTab[93554]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:527
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:527
			// _ = "end of CoverTab[93549]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:527
			_go_fuzz_dep_.CoverTab[93550]++

														cv = load6432(src, s)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:529
			// _ = "end of CoverTab[93550]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:530
			_go_fuzz_dep_.CoverTab[93555]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:530
			// _ = "end of CoverTab[93555]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:530
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:530
		// _ = "end of CoverTab[93473]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:531
	// _ = "end of CoverTab[93445]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:531
	_go_fuzz_dep_.CoverTab[93446]++

												if int(nextEmit) < len(src) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:533
		_go_fuzz_dep_.CoverTab[93556]++
													blk.literals = append(blk.literals, src[nextEmit:]...)
													blk.extraLits = len(src) - int(nextEmit)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:535
		// _ = "end of CoverTab[93556]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:536
		_go_fuzz_dep_.CoverTab[93557]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:536
		// _ = "end of CoverTab[93557]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:536
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:536
	// _ = "end of CoverTab[93446]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:536
	_go_fuzz_dep_.CoverTab[93447]++
												if debugEncoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:537
		_go_fuzz_dep_.CoverTab[93558]++
													println("returning, recent offsets:", blk.recentOffsets, "extra literals:", blk.extraLits)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:538
		// _ = "end of CoverTab[93558]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:539
		_go_fuzz_dep_.CoverTab[93559]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:539
		// _ = "end of CoverTab[93559]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:539
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:539
	// _ = "end of CoverTab[93447]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:539
	_go_fuzz_dep_.CoverTab[93448]++

												if e.cur < bufferReset {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:541
		_go_fuzz_dep_.CoverTab[93560]++
													e.cur += int32(len(src))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:542
		// _ = "end of CoverTab[93560]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:543
		_go_fuzz_dep_.CoverTab[93561]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:543
		// _ = "end of CoverTab[93561]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:543
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:543
	// _ = "end of CoverTab[93448]"
}

// Encode will encode the content, with a dictionary if initialized for it.
func (e *fastEncoderDict) Encode(blk *blockEnc, src []byte) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:547
	_go_fuzz_dep_.CoverTab[93562]++
												const (
		inputMargin		= 8
		minNonLiteralBlockSize	= 1 + 1 + inputMargin
	)
	if e.allDirty || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:552
		_go_fuzz_dep_.CoverTab[93570]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:552
		return len(src) > 32<<10
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:552
		// _ = "end of CoverTab[93570]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:552
	}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:552
		_go_fuzz_dep_.CoverTab[93571]++
													e.fastEncoder.Encode(blk, src)
													e.allDirty = true
													return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:555
		// _ = "end of CoverTab[93571]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:556
		_go_fuzz_dep_.CoverTab[93572]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:556
		// _ = "end of CoverTab[93572]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:556
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:556
	// _ = "end of CoverTab[93562]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:556
	_go_fuzz_dep_.CoverTab[93563]++

												for e.cur >= bufferReset {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:558
		_go_fuzz_dep_.CoverTab[93573]++
													if len(e.hist) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:559
			_go_fuzz_dep_.CoverTab[93576]++
														for i := range e.table[:] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:560
				_go_fuzz_dep_.CoverTab[93578]++
															e.table[i] = tableEntry{}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:561
				// _ = "end of CoverTab[93578]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:562
			// _ = "end of CoverTab[93576]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:562
			_go_fuzz_dep_.CoverTab[93577]++
														e.cur = e.maxMatchOff
														break
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:564
			// _ = "end of CoverTab[93577]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:565
			_go_fuzz_dep_.CoverTab[93579]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:565
			// _ = "end of CoverTab[93579]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:565
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:565
		// _ = "end of CoverTab[93573]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:565
		_go_fuzz_dep_.CoverTab[93574]++

													minOff := e.cur + int32(len(e.hist)) - e.maxMatchOff
													for i := range e.table[:] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:568
			_go_fuzz_dep_.CoverTab[93580]++
														v := e.table[i].offset
														if v < minOff {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:570
				_go_fuzz_dep_.CoverTab[93582]++
															v = 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:571
				// _ = "end of CoverTab[93582]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:572
				_go_fuzz_dep_.CoverTab[93583]++
															v = v - e.cur + e.maxMatchOff
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:573
				// _ = "end of CoverTab[93583]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:574
			// _ = "end of CoverTab[93580]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:574
			_go_fuzz_dep_.CoverTab[93581]++
														e.table[i].offset = v
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:575
			// _ = "end of CoverTab[93581]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:576
		// _ = "end of CoverTab[93574]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:576
		_go_fuzz_dep_.CoverTab[93575]++
													e.cur = e.maxMatchOff
													break
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:578
		// _ = "end of CoverTab[93575]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:579
	// _ = "end of CoverTab[93563]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:579
	_go_fuzz_dep_.CoverTab[93564]++

												s := e.addBlock(src)
												blk.size = len(src)
												if len(src) < minNonLiteralBlockSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:583
		_go_fuzz_dep_.CoverTab[93584]++
													blk.extraLits = len(src)
													blk.literals = blk.literals[:len(src)]
													copy(blk.literals, src)
													return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:587
		// _ = "end of CoverTab[93584]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:588
		_go_fuzz_dep_.CoverTab[93585]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:588
		// _ = "end of CoverTab[93585]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:588
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:588
	// _ = "end of CoverTab[93564]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:588
	_go_fuzz_dep_.CoverTab[93565]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:591
	src = e.hist
	sLimit := int32(len(src)) - inputMargin
	// stepSize is the number of bytes to skip on every main loop iteration.
	// It should be >= 2.
	const stepSize = 2

	// TEMPLATE
	const hashLog = tableBits
												// seems global, but would be nice to tweak.
												const kSearchStrength = 7

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:603
	nextEmit := s
												cv := load6432(src, s)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:607
	offset1 := int32(blk.recentOffsets[0])
	offset2 := int32(blk.recentOffsets[1])

	addLiterals := func(s *seq, until int32) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:610
		_go_fuzz_dep_.CoverTab[93586]++
													if until == nextEmit {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:611
			_go_fuzz_dep_.CoverTab[93588]++
														return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:612
			// _ = "end of CoverTab[93588]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:613
			_go_fuzz_dep_.CoverTab[93589]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:613
			// _ = "end of CoverTab[93589]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:613
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:613
		// _ = "end of CoverTab[93586]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:613
		_go_fuzz_dep_.CoverTab[93587]++
													blk.literals = append(blk.literals, src[nextEmit:until]...)
													s.litLen = uint32(until - nextEmit)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:615
		// _ = "end of CoverTab[93587]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:616
	// _ = "end of CoverTab[93565]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:616
	_go_fuzz_dep_.CoverTab[93566]++
												if debugEncoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:617
		_go_fuzz_dep_.CoverTab[93590]++
													println("recent offsets:", blk.recentOffsets)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:618
		// _ = "end of CoverTab[93590]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:619
		_go_fuzz_dep_.CoverTab[93591]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:619
		// _ = "end of CoverTab[93591]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:619
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:619
	// _ = "end of CoverTab[93566]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:619
	_go_fuzz_dep_.CoverTab[93567]++

encodeLoop:
	for {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:622
		_go_fuzz_dep_.CoverTab[93592]++
		// t will contain the match offset when we find one.
													// When existing the search loop, we have already checked 4 bytes.
													var t int32

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:629
		canRepeat := len(blk.sequences) > 2

		for {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:631
			_go_fuzz_dep_.CoverTab[93601]++
														if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:632
				_go_fuzz_dep_.CoverTab[93607]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:632
				return canRepeat
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:632
				// _ = "end of CoverTab[93607]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:632
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:632
				_go_fuzz_dep_.CoverTab[93608]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:632
				return offset1 == 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:632
				// _ = "end of CoverTab[93608]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:632
			}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:632
				_go_fuzz_dep_.CoverTab[93609]++
															panic("offset0 was 0")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:633
				// _ = "end of CoverTab[93609]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:634
				_go_fuzz_dep_.CoverTab[93610]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:634
				// _ = "end of CoverTab[93610]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:634
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:634
			// _ = "end of CoverTab[93601]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:634
			_go_fuzz_dep_.CoverTab[93602]++

														nextHash := hashLen(cv, hashLog, tableFastHashLen)
														nextHash2 := hashLen(cv>>8, hashLog, tableFastHashLen)
														candidate := e.table[nextHash]
														candidate2 := e.table[nextHash2]
														repIndex := s - offset1 + 2

														e.table[nextHash] = tableEntry{offset: s + e.cur, val: uint32(cv)}
														e.markShardDirty(nextHash)
														e.table[nextHash2] = tableEntry{offset: s + e.cur + 1, val: uint32(cv >> 8)}
														e.markShardDirty(nextHash2)

														if canRepeat && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:647
				_go_fuzz_dep_.CoverTab[93611]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:647
				return repIndex >= 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:647
				// _ = "end of CoverTab[93611]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:647
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:647
				_go_fuzz_dep_.CoverTab[93612]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:647
				return load3232(src, repIndex) == uint32(cv>>16)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:647
				// _ = "end of CoverTab[93612]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:647
			}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:647
				_go_fuzz_dep_.CoverTab[93613]++
															// Consider history as well.
															var seq seq
															var length int32
															length = 4 + e.matchlen(s+6, repIndex+4, src)

															seq.matchLen = uint32(length - zstdMinMatch)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:657
				start := s + 2

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:660
				startLimit := nextEmit + 1

				sMin := s - e.maxMatchOff
				if sMin < 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:663
					_go_fuzz_dep_.CoverTab[93618]++
																sMin = 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:664
					// _ = "end of CoverTab[93618]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:665
					_go_fuzz_dep_.CoverTab[93619]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:665
					// _ = "end of CoverTab[93619]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:665
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:665
				// _ = "end of CoverTab[93613]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:665
				_go_fuzz_dep_.CoverTab[93614]++
															for repIndex > sMin && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:666
					_go_fuzz_dep_.CoverTab[93620]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:666
					return start > startLimit
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:666
					// _ = "end of CoverTab[93620]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:666
				}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:666
					_go_fuzz_dep_.CoverTab[93621]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:666
					return src[repIndex-1] == src[start-1]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:666
					// _ = "end of CoverTab[93621]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:666
				}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:666
					_go_fuzz_dep_.CoverTab[93622]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:666
					return seq.matchLen < maxMatchLength-zstdMinMatch
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:666
					// _ = "end of CoverTab[93622]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:666
				}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:666
					_go_fuzz_dep_.CoverTab[93623]++
																repIndex--
																start--
																seq.matchLen++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:669
					// _ = "end of CoverTab[93623]"
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:670
				// _ = "end of CoverTab[93614]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:670
				_go_fuzz_dep_.CoverTab[93615]++
															addLiterals(&seq, start)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:674
				seq.offset = 1
				if debugSequences {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:675
					_go_fuzz_dep_.CoverTab[93624]++
																println("repeat sequence", seq, "next s:", s)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:676
					// _ = "end of CoverTab[93624]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:677
					_go_fuzz_dep_.CoverTab[93625]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:677
					// _ = "end of CoverTab[93625]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:677
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:677
				// _ = "end of CoverTab[93615]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:677
				_go_fuzz_dep_.CoverTab[93616]++
															blk.sequences = append(blk.sequences, seq)
															s += length + 2
															nextEmit = s
															if s >= sLimit {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:681
					_go_fuzz_dep_.CoverTab[93626]++
																if debugEncoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:682
						_go_fuzz_dep_.CoverTab[93628]++
																	println("repeat ended", s, length)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:683
						// _ = "end of CoverTab[93628]"

					} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:685
						_go_fuzz_dep_.CoverTab[93629]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:685
						// _ = "end of CoverTab[93629]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:685
					}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:685
					// _ = "end of CoverTab[93626]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:685
					_go_fuzz_dep_.CoverTab[93627]++
																break encodeLoop
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:686
					// _ = "end of CoverTab[93627]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:687
					_go_fuzz_dep_.CoverTab[93630]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:687
					// _ = "end of CoverTab[93630]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:687
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:687
				// _ = "end of CoverTab[93616]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:687
				_go_fuzz_dep_.CoverTab[93617]++
															cv = load6432(src, s)
															continue
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:689
				// _ = "end of CoverTab[93617]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:690
				_go_fuzz_dep_.CoverTab[93631]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:690
				// _ = "end of CoverTab[93631]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:690
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:690
			// _ = "end of CoverTab[93602]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:690
			_go_fuzz_dep_.CoverTab[93603]++
														coffset0 := s - (candidate.offset - e.cur)
														coffset1 := s - (candidate2.offset - e.cur) + 1
														if coffset0 < e.maxMatchOff && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:693
				_go_fuzz_dep_.CoverTab[93632]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:693
				return uint32(cv) == candidate.val
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:693
				// _ = "end of CoverTab[93632]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:693
			}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:693
				_go_fuzz_dep_.CoverTab[93633]++

															t = candidate.offset - e.cur
															if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:696
					_go_fuzz_dep_.CoverTab[93636]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:696
					return s <= t
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:696
					// _ = "end of CoverTab[93636]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:696
				}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:696
					_go_fuzz_dep_.CoverTab[93637]++
																panic(fmt.Sprintf("s (%d) <= t (%d)", s, t))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:697
					// _ = "end of CoverTab[93637]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:698
					_go_fuzz_dep_.CoverTab[93638]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:698
					// _ = "end of CoverTab[93638]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:698
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:698
				// _ = "end of CoverTab[93633]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:698
				_go_fuzz_dep_.CoverTab[93634]++
															if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:699
					_go_fuzz_dep_.CoverTab[93639]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:699
					return s-t > e.maxMatchOff
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:699
					// _ = "end of CoverTab[93639]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:699
				}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:699
					_go_fuzz_dep_.CoverTab[93640]++
																panic("s - t >e.maxMatchOff")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:700
					// _ = "end of CoverTab[93640]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:701
					_go_fuzz_dep_.CoverTab[93641]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:701
					// _ = "end of CoverTab[93641]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:701
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:701
				// _ = "end of CoverTab[93634]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:701
				_go_fuzz_dep_.CoverTab[93635]++
															break
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:702
				// _ = "end of CoverTab[93635]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:703
				_go_fuzz_dep_.CoverTab[93642]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:703
				// _ = "end of CoverTab[93642]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:703
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:703
			// _ = "end of CoverTab[93603]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:703
			_go_fuzz_dep_.CoverTab[93604]++

														if coffset1 < e.maxMatchOff && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:705
				_go_fuzz_dep_.CoverTab[93643]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:705
				return uint32(cv>>8) == candidate2.val
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:705
				// _ = "end of CoverTab[93643]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:705
			}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:705
				_go_fuzz_dep_.CoverTab[93644]++

															t = candidate2.offset - e.cur
															s++
															if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:709
					_go_fuzz_dep_.CoverTab[93648]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:709
					return s <= t
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:709
					// _ = "end of CoverTab[93648]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:709
				}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:709
					_go_fuzz_dep_.CoverTab[93649]++
																panic(fmt.Sprintf("s (%d) <= t (%d)", s, t))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:710
					// _ = "end of CoverTab[93649]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:711
					_go_fuzz_dep_.CoverTab[93650]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:711
					// _ = "end of CoverTab[93650]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:711
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:711
				// _ = "end of CoverTab[93644]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:711
				_go_fuzz_dep_.CoverTab[93645]++
															if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:712
					_go_fuzz_dep_.CoverTab[93651]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:712
					return s-t > e.maxMatchOff
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:712
					// _ = "end of CoverTab[93651]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:712
				}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:712
					_go_fuzz_dep_.CoverTab[93652]++
																panic("s - t >e.maxMatchOff")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:713
					// _ = "end of CoverTab[93652]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:714
					_go_fuzz_dep_.CoverTab[93653]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:714
					// _ = "end of CoverTab[93653]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:714
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:714
				// _ = "end of CoverTab[93645]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:714
				_go_fuzz_dep_.CoverTab[93646]++
															if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:715
					_go_fuzz_dep_.CoverTab[93654]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:715
					return t < 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:715
					// _ = "end of CoverTab[93654]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:715
				}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:715
					_go_fuzz_dep_.CoverTab[93655]++
																panic("t<0")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:716
					// _ = "end of CoverTab[93655]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:717
					_go_fuzz_dep_.CoverTab[93656]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:717
					// _ = "end of CoverTab[93656]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:717
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:717
				// _ = "end of CoverTab[93646]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:717
				_go_fuzz_dep_.CoverTab[93647]++
															break
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:718
				// _ = "end of CoverTab[93647]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:719
				_go_fuzz_dep_.CoverTab[93657]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:719
				// _ = "end of CoverTab[93657]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:719
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:719
			// _ = "end of CoverTab[93604]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:719
			_go_fuzz_dep_.CoverTab[93605]++
														s += stepSize + ((s - nextEmit) >> (kSearchStrength - 1))
														if s >= sLimit {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:721
				_go_fuzz_dep_.CoverTab[93658]++
															break encodeLoop
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:722
				// _ = "end of CoverTab[93658]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:723
				_go_fuzz_dep_.CoverTab[93659]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:723
				// _ = "end of CoverTab[93659]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:723
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:723
			// _ = "end of CoverTab[93605]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:723
			_go_fuzz_dep_.CoverTab[93606]++
														cv = load6432(src, s)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:724
			// _ = "end of CoverTab[93606]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:725
		// _ = "end of CoverTab[93592]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:725
		_go_fuzz_dep_.CoverTab[93593]++

													offset2 = offset1
													offset1 = s - t

													if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:730
			_go_fuzz_dep_.CoverTab[93660]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:730
			return s <= t
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:730
			// _ = "end of CoverTab[93660]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:730
		}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:730
			_go_fuzz_dep_.CoverTab[93661]++
														panic(fmt.Sprintf("s (%d) <= t (%d)", s, t))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:731
			// _ = "end of CoverTab[93661]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:732
			_go_fuzz_dep_.CoverTab[93662]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:732
			// _ = "end of CoverTab[93662]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:732
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:732
		// _ = "end of CoverTab[93593]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:732
		_go_fuzz_dep_.CoverTab[93594]++

													if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:734
			_go_fuzz_dep_.CoverTab[93663]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:734
			return canRepeat
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:734
			// _ = "end of CoverTab[93663]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:734
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:734
			_go_fuzz_dep_.CoverTab[93664]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:734
			return int(offset1) > len(src)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:734
			// _ = "end of CoverTab[93664]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:734
		}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:734
			_go_fuzz_dep_.CoverTab[93665]++
														panic("invalid offset")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:735
			// _ = "end of CoverTab[93665]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:736
			_go_fuzz_dep_.CoverTab[93666]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:736
			// _ = "end of CoverTab[93666]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:736
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:736
		// _ = "end of CoverTab[93594]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:736
		_go_fuzz_dep_.CoverTab[93595]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:739
		l := e.matchlen(s+4, t+4, src) + 4

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:742
		tMin := s - e.maxMatchOff
		if tMin < 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:743
			_go_fuzz_dep_.CoverTab[93667]++
														tMin = 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:744
			// _ = "end of CoverTab[93667]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:745
			_go_fuzz_dep_.CoverTab[93668]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:745
			// _ = "end of CoverTab[93668]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:745
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:745
		// _ = "end of CoverTab[93595]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:745
		_go_fuzz_dep_.CoverTab[93596]++
													for t > tMin && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:746
			_go_fuzz_dep_.CoverTab[93669]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:746
			return s > nextEmit
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:746
			// _ = "end of CoverTab[93669]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:746
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:746
			_go_fuzz_dep_.CoverTab[93670]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:746
			return src[t-1] == src[s-1]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:746
			// _ = "end of CoverTab[93670]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:746
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:746
			_go_fuzz_dep_.CoverTab[93671]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:746
			return l < maxMatchLength
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:746
			// _ = "end of CoverTab[93671]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:746
		}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:746
			_go_fuzz_dep_.CoverTab[93672]++
														s--
														t--
														l++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:749
			// _ = "end of CoverTab[93672]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:750
		// _ = "end of CoverTab[93596]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:750
		_go_fuzz_dep_.CoverTab[93597]++

		// Write our sequence.
		var seq seq
		seq.litLen = uint32(s - nextEmit)
		seq.matchLen = uint32(l - zstdMinMatch)
		if seq.litLen > 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:756
			_go_fuzz_dep_.CoverTab[93673]++
														blk.literals = append(blk.literals, src[nextEmit:s]...)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:757
			// _ = "end of CoverTab[93673]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:758
			_go_fuzz_dep_.CoverTab[93674]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:758
			// _ = "end of CoverTab[93674]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:758
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:758
		// _ = "end of CoverTab[93597]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:758
		_go_fuzz_dep_.CoverTab[93598]++

													seq.offset = uint32(s-t) + 3
													s += l
													if debugSequences {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:762
			_go_fuzz_dep_.CoverTab[93675]++
														println("sequence", seq, "next s:", s)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:763
			// _ = "end of CoverTab[93675]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:764
			_go_fuzz_dep_.CoverTab[93676]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:764
			// _ = "end of CoverTab[93676]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:764
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:764
		// _ = "end of CoverTab[93598]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:764
		_go_fuzz_dep_.CoverTab[93599]++
													blk.sequences = append(blk.sequences, seq)
													nextEmit = s
													if s >= sLimit {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:767
			_go_fuzz_dep_.CoverTab[93677]++
														break encodeLoop
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:768
			// _ = "end of CoverTab[93677]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:769
			_go_fuzz_dep_.CoverTab[93678]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:769
			// _ = "end of CoverTab[93678]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:769
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:769
		// _ = "end of CoverTab[93599]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:769
		_go_fuzz_dep_.CoverTab[93600]++
													cv = load6432(src, s)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:773
		if o2 := s - offset2; canRepeat && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:773
			_go_fuzz_dep_.CoverTab[93679]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:773
			return load3232(src, o2) == uint32(cv)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:773
			// _ = "end of CoverTab[93679]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:773
		}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:773
			_go_fuzz_dep_.CoverTab[93680]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:776
			l := 4 + e.matchlen(s+4, o2+4, src)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:779
			nextHash := hashLen(cv, hashLog, tableFastHashLen)
			e.table[nextHash] = tableEntry{offset: s + e.cur, val: uint32(cv)}
			e.markShardDirty(nextHash)
			seq.matchLen = uint32(l) - zstdMinMatch
			seq.litLen = 0

			seq.offset = 1
			s += l
			nextEmit = s
			if debugSequences {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:788
				_go_fuzz_dep_.CoverTab[93683]++
															println("sequence", seq, "next s:", s)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:789
				// _ = "end of CoverTab[93683]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:790
				_go_fuzz_dep_.CoverTab[93684]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:790
				// _ = "end of CoverTab[93684]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:790
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:790
			// _ = "end of CoverTab[93680]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:790
			_go_fuzz_dep_.CoverTab[93681]++
														blk.sequences = append(blk.sequences, seq)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:794
			offset1, offset2 = offset2, offset1
			if s >= sLimit {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:795
				_go_fuzz_dep_.CoverTab[93685]++
															break encodeLoop
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:796
				// _ = "end of CoverTab[93685]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:797
				_go_fuzz_dep_.CoverTab[93686]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:797
				// _ = "end of CoverTab[93686]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:797
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:797
			// _ = "end of CoverTab[93681]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:797
			_go_fuzz_dep_.CoverTab[93682]++

														cv = load6432(src, s)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:799
			// _ = "end of CoverTab[93682]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:800
			_go_fuzz_dep_.CoverTab[93687]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:800
			// _ = "end of CoverTab[93687]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:800
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:800
		// _ = "end of CoverTab[93600]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:801
	// _ = "end of CoverTab[93567]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:801
	_go_fuzz_dep_.CoverTab[93568]++

												if int(nextEmit) < len(src) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:803
		_go_fuzz_dep_.CoverTab[93688]++
													blk.literals = append(blk.literals, src[nextEmit:]...)
													blk.extraLits = len(src) - int(nextEmit)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:805
		// _ = "end of CoverTab[93688]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:806
		_go_fuzz_dep_.CoverTab[93689]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:806
		// _ = "end of CoverTab[93689]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:806
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:806
	// _ = "end of CoverTab[93568]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:806
	_go_fuzz_dep_.CoverTab[93569]++
												blk.recentOffsets[0] = uint32(offset1)
												blk.recentOffsets[1] = uint32(offset2)
												if debugEncoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:809
		_go_fuzz_dep_.CoverTab[93690]++
													println("returning, recent offsets:", blk.recentOffsets, "extra literals:", blk.extraLits)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:810
		// _ = "end of CoverTab[93690]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:811
		_go_fuzz_dep_.CoverTab[93691]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:811
		// _ = "end of CoverTab[93691]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:811
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:811
	// _ = "end of CoverTab[93569]"
}

// ResetDict will reset and set a dictionary if not nil
func (e *fastEncoder) Reset(d *dict, singleBlock bool) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:815
	_go_fuzz_dep_.CoverTab[93692]++
												e.resetBase(d, singleBlock)
												if d != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:817
		_go_fuzz_dep_.CoverTab[93693]++
													panic("fastEncoder: Reset with dict")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:818
		// _ = "end of CoverTab[93693]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:819
		_go_fuzz_dep_.CoverTab[93694]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:819
		// _ = "end of CoverTab[93694]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:819
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:819
	// _ = "end of CoverTab[93692]"
}

// ResetDict will reset and set a dictionary if not nil
func (e *fastEncoderDict) Reset(d *dict, singleBlock bool) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:823
	_go_fuzz_dep_.CoverTab[93695]++
												e.resetBase(d, singleBlock)
												if d == nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:825
		_go_fuzz_dep_.CoverTab[93701]++
													return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:826
		// _ = "end of CoverTab[93701]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:827
		_go_fuzz_dep_.CoverTab[93702]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:827
		// _ = "end of CoverTab[93702]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:827
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:827
	// _ = "end of CoverTab[93695]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:827
	_go_fuzz_dep_.CoverTab[93696]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:830
	if len(e.dictTable) != len(e.table) || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:830
		_go_fuzz_dep_.CoverTab[93703]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:830
		return d.id != e.lastDictID
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:830
		// _ = "end of CoverTab[93703]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:830
	}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:830
		_go_fuzz_dep_.CoverTab[93704]++
													if len(e.dictTable) != len(e.table) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:831
			_go_fuzz_dep_.CoverTab[93707]++
														e.dictTable = make([]tableEntry, len(e.table))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:832
			// _ = "end of CoverTab[93707]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:833
			_go_fuzz_dep_.CoverTab[93708]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:833
			// _ = "end of CoverTab[93708]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:833
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:833
		// _ = "end of CoverTab[93704]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:833
		_go_fuzz_dep_.CoverTab[93705]++
													if true {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:834
			_go_fuzz_dep_.CoverTab[93709]++
														end := e.maxMatchOff + int32(len(d.content)) - 8
														for i := e.maxMatchOff; i < end; i += 3 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:836
				_go_fuzz_dep_.CoverTab[93710]++
															const hashLog = tableBits

															cv := load6432(d.content, i-e.maxMatchOff)
															nextHash := hashLen(cv, hashLog, tableFastHashLen)
															nextHash1 := hashLen(cv>>8, hashLog, tableFastHashLen)
															nextHash2 := hashLen(cv>>16, hashLog, tableFastHashLen)
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
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:854
				// _ = "end of CoverTab[93710]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:855
			// _ = "end of CoverTab[93709]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:856
			_go_fuzz_dep_.CoverTab[93711]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:856
			// _ = "end of CoverTab[93711]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:856
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:856
		// _ = "end of CoverTab[93705]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:856
		_go_fuzz_dep_.CoverTab[93706]++
													e.lastDictID = d.id
													e.allDirty = true
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:858
		// _ = "end of CoverTab[93706]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:859
		_go_fuzz_dep_.CoverTab[93712]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:859
		// _ = "end of CoverTab[93712]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:859
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:859
	// _ = "end of CoverTab[93696]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:859
	_go_fuzz_dep_.CoverTab[93697]++

												e.cur = e.maxMatchOff
												dirtyShardCnt := 0
												if !e.allDirty {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:863
		_go_fuzz_dep_.CoverTab[93713]++
													for i := range e.tableShardDirty {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:864
			_go_fuzz_dep_.CoverTab[93714]++
														if e.tableShardDirty[i] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:865
				_go_fuzz_dep_.CoverTab[93715]++
															dirtyShardCnt++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:866
				// _ = "end of CoverTab[93715]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:867
				_go_fuzz_dep_.CoverTab[93716]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:867
				// _ = "end of CoverTab[93716]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:867
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:867
			// _ = "end of CoverTab[93714]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:868
		// _ = "end of CoverTab[93713]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:869
		_go_fuzz_dep_.CoverTab[93717]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:869
		// _ = "end of CoverTab[93717]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:869
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:869
	// _ = "end of CoverTab[93697]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:869
	_go_fuzz_dep_.CoverTab[93698]++

												const shardCnt = tableShardCnt
												const shardSize = tableShardSize
												if e.allDirty || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:873
		_go_fuzz_dep_.CoverTab[93718]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:873
		return dirtyShardCnt > shardCnt*4/6
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:873
		// _ = "end of CoverTab[93718]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:873
	}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:873
		_go_fuzz_dep_.CoverTab[93719]++
													copy(e.table[:], e.dictTable)
													for i := range e.tableShardDirty {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:875
			_go_fuzz_dep_.CoverTab[93721]++
														e.tableShardDirty[i] = false
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:876
			// _ = "end of CoverTab[93721]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:877
		// _ = "end of CoverTab[93719]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:877
		_go_fuzz_dep_.CoverTab[93720]++
													e.allDirty = false
													return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:879
		// _ = "end of CoverTab[93720]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:880
		_go_fuzz_dep_.CoverTab[93722]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:880
		// _ = "end of CoverTab[93722]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:880
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:880
	// _ = "end of CoverTab[93698]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:880
	_go_fuzz_dep_.CoverTab[93699]++
												for i := range e.tableShardDirty {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:881
		_go_fuzz_dep_.CoverTab[93723]++
													if !e.tableShardDirty[i] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:882
			_go_fuzz_dep_.CoverTab[93725]++
														continue
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:883
			// _ = "end of CoverTab[93725]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:884
			_go_fuzz_dep_.CoverTab[93726]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:884
			// _ = "end of CoverTab[93726]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:884
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:884
		// _ = "end of CoverTab[93723]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:884
		_go_fuzz_dep_.CoverTab[93724]++

													copy(e.table[i*shardSize:(i+1)*shardSize], e.dictTable[i*shardSize:(i+1)*shardSize])
													e.tableShardDirty[i] = false
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:887
		// _ = "end of CoverTab[93724]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:888
	// _ = "end of CoverTab[93699]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:888
	_go_fuzz_dep_.CoverTab[93700]++
												e.allDirty = false
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:889
	// _ = "end of CoverTab[93700]"
}

func (e *fastEncoderDict) markAllShardsDirty() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:892
	_go_fuzz_dep_.CoverTab[93727]++
												e.allDirty = true
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:893
	// _ = "end of CoverTab[93727]"
}

func (e *fastEncoderDict) markShardDirty(entryNum uint32) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:896
	_go_fuzz_dep_.CoverTab[93728]++
												e.tableShardDirty[entryNum/tableShardSize] = true
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:897
	// _ = "end of CoverTab[93728]"
}

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:898
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_fast.go:898
var _ = _go_fuzz_dep_.CoverTab
