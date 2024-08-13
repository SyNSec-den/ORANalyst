// Copyright 2019+ Klaus Post. All rights reserved.
// License information can be found in the LICENSE file.
// Based on work by Yann Collet, released under BSD License.

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:5
package zstd

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:5
)

import (
	"errors"
	"fmt"
)

const (
	tablelogAbsoluteMax = 9
)

const (
	/*!MEMORY_USAGE :
	 *  Memory usage formula : N->2^N Bytes (examples : 10 -> 1KB; 12 -> 4KB ; 16 -> 64KB; 20 -> 1MB; etc.)
	 *  Increasing memory usage improves compression ratio
	 *  Reduced memory usage can improve speed, due to cache effect
	 *  Recommended max value is 14, for 16KB, which nicely fits into Intel x86 L1 cache */
	maxMemoryUsage	= tablelogAbsoluteMax + 2

	maxTableLog	= maxMemoryUsage - 2
	maxTablesize	= 1 << maxTableLog
	maxTableMask	= (1 << maxTableLog) - 1
	minTablelog	= 5
	maxSymbolValue	= 255
)

// fseDecoder provides temporary storage for compression and decompression.
type fseDecoder struct {
	dt		[maxTablesize]decSymbol	// Decompression table.
	symbolLen	uint16			// Length of active part of the symbol table.
	actualTableLog	uint8			// Selected tablelog.
	maxBits		uint8			// Maximum number of additional bits

	// used for table creation to avoid allocations.
	stateTable	[256]uint16
	norm		[maxSymbolValue + 1]int16
	preDefined	bool
}

// tableStep returns the next table index.
func tableStep(tableSize uint32) uint32 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:45
	_go_fuzz_dep_.CoverTab[94405]++
												return (tableSize >> 1) + (tableSize >> 3) + 3
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:46
	// _ = "end of CoverTab[94405]"
}

// readNCount will read the symbol distribution so decoding tables can be constructed.
func (s *fseDecoder) readNCount(b *byteReader, maxSymbol uint16) error {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:50
	_go_fuzz_dep_.CoverTab[94406]++
												var (
		charnum		uint16
		previous0	bool
	)
	if b.remain() < 4 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:55
		_go_fuzz_dep_.CoverTab[94415]++
													return errors.New("input too small")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:56
		// _ = "end of CoverTab[94415]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:57
		_go_fuzz_dep_.CoverTab[94416]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:57
		// _ = "end of CoverTab[94416]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:57
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:57
	// _ = "end of CoverTab[94406]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:57
	_go_fuzz_dep_.CoverTab[94407]++
												bitStream := b.Uint32NC()
												nbBits := uint((bitStream & 0xF) + minTablelog)
												if nbBits > tablelogAbsoluteMax {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:60
		_go_fuzz_dep_.CoverTab[94417]++
													println("Invalid tablelog:", nbBits)
													return errors.New("tableLog too large")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:62
		// _ = "end of CoverTab[94417]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:63
		_go_fuzz_dep_.CoverTab[94418]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:63
		// _ = "end of CoverTab[94418]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:63
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:63
	// _ = "end of CoverTab[94407]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:63
	_go_fuzz_dep_.CoverTab[94408]++
												bitStream >>= 4
												bitCount := uint(4)

												s.actualTableLog = uint8(nbBits)
												remaining := int32((1 << nbBits) + 1)
												threshold := int32(1 << nbBits)
												gotTotal := int32(0)
												nbBits++

												for remaining > 1 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:73
		_go_fuzz_dep_.CoverTab[94419]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:73
		return charnum <= maxSymbol
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:73
		// _ = "end of CoverTab[94419]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:73
	}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:73
		_go_fuzz_dep_.CoverTab[94420]++
													if previous0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:74
			_go_fuzz_dep_.CoverTab[94425]++

														n0 := charnum
														for (bitStream & 0xFFFF) == 0xFFFF {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:77
				_go_fuzz_dep_.CoverTab[94430]++

															n0 += 24
															if r := b.remain(); r > 5 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:80
					_go_fuzz_dep_.CoverTab[94431]++
																b.advance(2)

																bitStream = b.Uint32NC() >> bitCount
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:83
					// _ = "end of CoverTab[94431]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:84
					_go_fuzz_dep_.CoverTab[94432]++

																bitStream >>= 16
																bitCount += 16
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:87
					// _ = "end of CoverTab[94432]"
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:88
				// _ = "end of CoverTab[94430]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:89
			// _ = "end of CoverTab[94425]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:89
			_go_fuzz_dep_.CoverTab[94426]++

														for (bitStream & 3) == 3 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:91
				_go_fuzz_dep_.CoverTab[94433]++
															n0 += 3
															bitStream >>= 2
															bitCount += 2
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:94
				// _ = "end of CoverTab[94433]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:95
			// _ = "end of CoverTab[94426]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:95
			_go_fuzz_dep_.CoverTab[94427]++
														n0 += uint16(bitStream & 3)
														bitCount += 2

														if n0 > maxSymbolValue {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:99
					_go_fuzz_dep_.CoverTab[94434]++
																return errors.New("maxSymbolValue too small")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:100
				// _ = "end of CoverTab[94434]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:101
				_go_fuzz_dep_.CoverTab[94435]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:101
				// _ = "end of CoverTab[94435]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:101
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:101
			// _ = "end of CoverTab[94427]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:101
			_go_fuzz_dep_.CoverTab[94428]++

															for charnum < n0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:103
				_go_fuzz_dep_.CoverTab[94436]++
																s.norm[uint8(charnum)] = 0
																charnum++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:105
				// _ = "end of CoverTab[94436]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:106
			// _ = "end of CoverTab[94428]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:106
			_go_fuzz_dep_.CoverTab[94429]++

															if r := b.remain(); r >= 7 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:108
				_go_fuzz_dep_.CoverTab[94437]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:108
				return r-int(bitCount>>3) >= 4
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:108
				// _ = "end of CoverTab[94437]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:108
			}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:108
				_go_fuzz_dep_.CoverTab[94438]++
																b.advance(bitCount >> 3)
																bitCount &= 7

																bitStream = b.Uint32NC() >> bitCount
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:112
				// _ = "end of CoverTab[94438]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:113
				_go_fuzz_dep_.CoverTab[94439]++
																bitStream >>= 2
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:114
				// _ = "end of CoverTab[94439]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:115
			// _ = "end of CoverTab[94429]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:116
			_go_fuzz_dep_.CoverTab[94440]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:116
			// _ = "end of CoverTab[94440]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:116
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:116
		// _ = "end of CoverTab[94420]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:116
		_go_fuzz_dep_.CoverTab[94421]++

														max := (2*threshold - 1) - remaining
														var count int32

														if int32(bitStream)&(threshold-1) < max {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:121
			_go_fuzz_dep_.CoverTab[94441]++
															count = int32(bitStream) & (threshold - 1)
															if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:123
				_go_fuzz_dep_.CoverTab[94443]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:123
				return nbBits < 1
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:123
				// _ = "end of CoverTab[94443]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:123
			}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:123
				_go_fuzz_dep_.CoverTab[94444]++
																panic("nbBits underflow")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:124
				// _ = "end of CoverTab[94444]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:125
				_go_fuzz_dep_.CoverTab[94445]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:125
				// _ = "end of CoverTab[94445]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:125
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:125
			// _ = "end of CoverTab[94441]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:125
			_go_fuzz_dep_.CoverTab[94442]++
															bitCount += nbBits - 1
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:126
			// _ = "end of CoverTab[94442]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:127
			_go_fuzz_dep_.CoverTab[94446]++
															count = int32(bitStream) & (2*threshold - 1)
															if count >= threshold {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:129
				_go_fuzz_dep_.CoverTab[94448]++
																count -= max
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:130
				// _ = "end of CoverTab[94448]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:131
				_go_fuzz_dep_.CoverTab[94449]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:131
				// _ = "end of CoverTab[94449]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:131
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:131
			// _ = "end of CoverTab[94446]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:131
			_go_fuzz_dep_.CoverTab[94447]++
															bitCount += nbBits
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:132
			// _ = "end of CoverTab[94447]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:133
		// _ = "end of CoverTab[94421]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:133
		_go_fuzz_dep_.CoverTab[94422]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:136
		count--
		if count < 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:137
			_go_fuzz_dep_.CoverTab[94450]++

															remaining += count
															gotTotal -= count
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:140
			// _ = "end of CoverTab[94450]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:141
			_go_fuzz_dep_.CoverTab[94451]++
															remaining -= count
															gotTotal += count
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:143
			// _ = "end of CoverTab[94451]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:144
		// _ = "end of CoverTab[94422]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:144
		_go_fuzz_dep_.CoverTab[94423]++
														s.norm[charnum&0xff] = int16(count)
														charnum++
														previous0 = count == 0
														for remaining < threshold {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:148
			_go_fuzz_dep_.CoverTab[94452]++
															nbBits--
															threshold >>= 1
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:150
			// _ = "end of CoverTab[94452]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:151
		// _ = "end of CoverTab[94423]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:151
		_go_fuzz_dep_.CoverTab[94424]++

														if r := b.remain(); r >= 7 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:153
			_go_fuzz_dep_.CoverTab[94453]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:153
			return r-int(bitCount>>3) >= 4
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:153
			// _ = "end of CoverTab[94453]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:153
		}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:153
			_go_fuzz_dep_.CoverTab[94454]++
															b.advance(bitCount >> 3)
															bitCount &= 7

															bitStream = b.Uint32NC() >> (bitCount & 31)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:157
			// _ = "end of CoverTab[94454]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:158
			_go_fuzz_dep_.CoverTab[94455]++
															bitCount -= (uint)(8 * (len(b.b) - 4 - b.off))
															b.off = len(b.b) - 4
															bitStream = b.Uint32() >> (bitCount & 31)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:161
			// _ = "end of CoverTab[94455]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:162
		// _ = "end of CoverTab[94424]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:163
	// _ = "end of CoverTab[94408]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:163
	_go_fuzz_dep_.CoverTab[94409]++
													s.symbolLen = charnum
													if s.symbolLen <= 1 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:165
		_go_fuzz_dep_.CoverTab[94456]++
														return fmt.Errorf("symbolLen (%d) too small", s.symbolLen)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:166
		// _ = "end of CoverTab[94456]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:167
		_go_fuzz_dep_.CoverTab[94457]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:167
		// _ = "end of CoverTab[94457]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:167
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:167
	// _ = "end of CoverTab[94409]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:167
	_go_fuzz_dep_.CoverTab[94410]++
													if s.symbolLen > maxSymbolValue+1 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:168
		_go_fuzz_dep_.CoverTab[94458]++
														return fmt.Errorf("symbolLen (%d) too big", s.symbolLen)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:169
		// _ = "end of CoverTab[94458]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:170
		_go_fuzz_dep_.CoverTab[94459]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:170
		// _ = "end of CoverTab[94459]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:170
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:170
	// _ = "end of CoverTab[94410]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:170
	_go_fuzz_dep_.CoverTab[94411]++
													if remaining != 1 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:171
		_go_fuzz_dep_.CoverTab[94460]++
														return fmt.Errorf("corruption detected (remaining %d != 1)", remaining)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:172
		// _ = "end of CoverTab[94460]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:173
		_go_fuzz_dep_.CoverTab[94461]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:173
		// _ = "end of CoverTab[94461]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:173
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:173
	// _ = "end of CoverTab[94411]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:173
	_go_fuzz_dep_.CoverTab[94412]++
													if bitCount > 32 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:174
		_go_fuzz_dep_.CoverTab[94462]++
														return fmt.Errorf("corruption detected (bitCount %d > 32)", bitCount)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:175
		// _ = "end of CoverTab[94462]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:176
		_go_fuzz_dep_.CoverTab[94463]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:176
		// _ = "end of CoverTab[94463]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:176
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:176
	// _ = "end of CoverTab[94412]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:176
	_go_fuzz_dep_.CoverTab[94413]++
													if gotTotal != 1<<s.actualTableLog {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:177
		_go_fuzz_dep_.CoverTab[94464]++
														return fmt.Errorf("corruption detected (total %d != %d)", gotTotal, 1<<s.actualTableLog)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:178
		// _ = "end of CoverTab[94464]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:179
		_go_fuzz_dep_.CoverTab[94465]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:179
		// _ = "end of CoverTab[94465]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:179
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:179
	// _ = "end of CoverTab[94413]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:179
	_go_fuzz_dep_.CoverTab[94414]++
													b.advance((bitCount + 7) >> 3)

													return s.buildDtable()
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:182
	// _ = "end of CoverTab[94414]"
}

// decSymbol contains information about a state entry,
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:185
// Including the state offset base, the output symbol and
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:185
// the number of bits to read for the low part of the destination state.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:185
// Using a composite uint64 is faster than a struct with separate members.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:189
type decSymbol uint64

func newDecSymbol(nbits, addBits uint8, newState uint16, baseline uint32) decSymbol {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:191
	_go_fuzz_dep_.CoverTab[94466]++
													return decSymbol(nbits) | (decSymbol(addBits) << 8) | (decSymbol(newState) << 16) | (decSymbol(baseline) << 32)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:192
	// _ = "end of CoverTab[94466]"
}

func (d decSymbol) nbBits() uint8 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:195
	_go_fuzz_dep_.CoverTab[94467]++
													return uint8(d)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:196
	// _ = "end of CoverTab[94467]"
}

func (d decSymbol) addBits() uint8 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:199
	_go_fuzz_dep_.CoverTab[94468]++
													return uint8(d >> 8)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:200
	// _ = "end of CoverTab[94468]"
}

func (d decSymbol) newState() uint16 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:203
	_go_fuzz_dep_.CoverTab[94469]++
													return uint16(d >> 16)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:204
	// _ = "end of CoverTab[94469]"
}

func (d decSymbol) baseline() uint32 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:207
	_go_fuzz_dep_.CoverTab[94470]++
													return uint32(d >> 32)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:208
	// _ = "end of CoverTab[94470]"
}

func (d decSymbol) baselineInt() int {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:211
	_go_fuzz_dep_.CoverTab[94471]++
													return int(d >> 32)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:212
	// _ = "end of CoverTab[94471]"
}

func (d *decSymbol) set(nbits, addBits uint8, newState uint16, baseline uint32) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:215
	_go_fuzz_dep_.CoverTab[94472]++
													*d = decSymbol(nbits) | (decSymbol(addBits) << 8) | (decSymbol(newState) << 16) | (decSymbol(baseline) << 32)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:216
	// _ = "end of CoverTab[94472]"
}

func (d *decSymbol) setNBits(nBits uint8) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:219
	_go_fuzz_dep_.CoverTab[94473]++
													const mask = 0xffffffffffffff00
													*d = (*d & mask) | decSymbol(nBits)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:221
	// _ = "end of CoverTab[94473]"
}

func (d *decSymbol) setAddBits(addBits uint8) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:224
	_go_fuzz_dep_.CoverTab[94474]++
													const mask = 0xffffffffffff00ff
													*d = (*d & mask) | (decSymbol(addBits) << 8)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:226
	// _ = "end of CoverTab[94474]"
}

func (d *decSymbol) setNewState(state uint16) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:229
	_go_fuzz_dep_.CoverTab[94475]++
													const mask = 0xffffffff0000ffff
													*d = (*d & mask) | decSymbol(state)<<16
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:231
	// _ = "end of CoverTab[94475]"
}

func (d *decSymbol) setBaseline(baseline uint32) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:234
	_go_fuzz_dep_.CoverTab[94476]++
													const mask = 0xffffffff
													*d = (*d & mask) | decSymbol(baseline)<<32
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:236
	// _ = "end of CoverTab[94476]"
}

func (d *decSymbol) setExt(addBits uint8, baseline uint32) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:239
	_go_fuzz_dep_.CoverTab[94477]++
													const mask = 0xffff00ff
													*d = (*d & mask) | (decSymbol(addBits) << 8) | (decSymbol(baseline) << 32)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:241
	// _ = "end of CoverTab[94477]"
}

// decSymbolValue returns the transformed decSymbol for the given symbol.
func decSymbolValue(symb uint8, t []baseOffset) (decSymbol, error) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:245
	_go_fuzz_dep_.CoverTab[94478]++
													if int(symb) >= len(t) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:246
		_go_fuzz_dep_.CoverTab[94480]++
														return 0, fmt.Errorf("rle symbol %d >= max %d", symb, len(t))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:247
		// _ = "end of CoverTab[94480]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:248
		_go_fuzz_dep_.CoverTab[94481]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:248
		// _ = "end of CoverTab[94481]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:248
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:248
	// _ = "end of CoverTab[94478]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:248
	_go_fuzz_dep_.CoverTab[94479]++
													lu := t[symb]
													return newDecSymbol(0, lu.addBits, 0, lu.baseLine), nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:250
	// _ = "end of CoverTab[94479]"
}

// setRLE will set the decoder til RLE mode.
func (s *fseDecoder) setRLE(symbol decSymbol) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:254
	_go_fuzz_dep_.CoverTab[94482]++
													s.actualTableLog = 0
													s.maxBits = symbol.addBits()
													s.dt[0] = symbol
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:257
	// _ = "end of CoverTab[94482]"
}

// buildDtable will build the decoding table.
func (s *fseDecoder) buildDtable() error {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:261
	_go_fuzz_dep_.CoverTab[94483]++
													tableSize := uint32(1 << s.actualTableLog)
													highThreshold := tableSize - 1
													symbolNext := s.stateTable[:256]

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:267
	{
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:267
		_go_fuzz_dep_.CoverTab[94485]++
														for i, v := range s.norm[:s.symbolLen] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:268
			_go_fuzz_dep_.CoverTab[94486]++
															if v == -1 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:269
				_go_fuzz_dep_.CoverTab[94487]++
																s.dt[highThreshold].setAddBits(uint8(i))
																highThreshold--
																symbolNext[i] = 1
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:272
				// _ = "end of CoverTab[94487]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:273
				_go_fuzz_dep_.CoverTab[94488]++
																symbolNext[i] = uint16(v)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:274
				// _ = "end of CoverTab[94488]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:275
			// _ = "end of CoverTab[94486]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:276
		// _ = "end of CoverTab[94485]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:277
	// _ = "end of CoverTab[94483]"

													{
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:279
		_go_fuzz_dep_.CoverTab[94489]++
														tableMask := tableSize - 1
														step := tableStep(tableSize)
														position := uint32(0)
														for ss, v := range s.norm[:s.symbolLen] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:283
			_go_fuzz_dep_.CoverTab[94491]++
															for i := 0; i < int(v); i++ {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:284
				_go_fuzz_dep_.CoverTab[94492]++
																s.dt[position].setAddBits(uint8(ss))
																position = (position + step) & tableMask
																for position > highThreshold {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:287
					_go_fuzz_dep_.CoverTab[94493]++

																	position = (position + step) & tableMask
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:289
					// _ = "end of CoverTab[94493]"
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:290
				// _ = "end of CoverTab[94492]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:291
			// _ = "end of CoverTab[94491]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:292
		// _ = "end of CoverTab[94489]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:292
		_go_fuzz_dep_.CoverTab[94490]++
														if position != 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:293
			_go_fuzz_dep_.CoverTab[94494]++

															return errors.New("corrupted input (position != 0)")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:295
			// _ = "end of CoverTab[94494]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:296
			_go_fuzz_dep_.CoverTab[94495]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:296
			// _ = "end of CoverTab[94495]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:296
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:296
		// _ = "end of CoverTab[94490]"
	}

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:300
	{
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:300
		_go_fuzz_dep_.CoverTab[94496]++
														tableSize := uint16(1 << s.actualTableLog)
														for u, v := range s.dt[:tableSize] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:302
			_go_fuzz_dep_.CoverTab[94497]++
															symbol := v.addBits()
															nextState := symbolNext[symbol]
															symbolNext[symbol] = nextState + 1
															nBits := s.actualTableLog - byte(highBits(uint32(nextState)))
															s.dt[u&maxTableMask].setNBits(nBits)
															newState := (nextState << nBits) - tableSize
															if newState > tableSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:309
				_go_fuzz_dep_.CoverTab[94500]++
																return fmt.Errorf("newState (%d) outside table size (%d)", newState, tableSize)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:310
				// _ = "end of CoverTab[94500]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:311
				_go_fuzz_dep_.CoverTab[94501]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:311
				// _ = "end of CoverTab[94501]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:311
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:311
			// _ = "end of CoverTab[94497]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:311
			_go_fuzz_dep_.CoverTab[94498]++
															if newState == uint16(u) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:312
				_go_fuzz_dep_.CoverTab[94502]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:312
				return nBits == 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:312
				// _ = "end of CoverTab[94502]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:312
			}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:312
				_go_fuzz_dep_.CoverTab[94503]++

																return fmt.Errorf("newState (%d) == oldState (%d) and no bits", newState, u)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:314
				// _ = "end of CoverTab[94503]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:315
				_go_fuzz_dep_.CoverTab[94504]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:315
				// _ = "end of CoverTab[94504]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:315
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:315
			// _ = "end of CoverTab[94498]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:315
			_go_fuzz_dep_.CoverTab[94499]++
															s.dt[u&maxTableMask].setNewState(newState)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:316
			// _ = "end of CoverTab[94499]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:317
		// _ = "end of CoverTab[94496]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:318
	_go_fuzz_dep_.CoverTab[94484]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:319
	// _ = "end of CoverTab[94484]"
}

// transform will transform the decoder table into a table usable for
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:322
// decoding without having to apply the transformation while decoding.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:322
// The state will contain the base value and the number of bits to read.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:325
func (s *fseDecoder) transform(t []baseOffset) error {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:325
	_go_fuzz_dep_.CoverTab[94505]++
													tableSize := uint16(1 << s.actualTableLog)
													s.maxBits = 0
													for i, v := range s.dt[:tableSize] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:328
		_go_fuzz_dep_.CoverTab[94507]++
														add := v.addBits()
														if int(add) >= len(t) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:330
			_go_fuzz_dep_.CoverTab[94510]++
															return fmt.Errorf("invalid decoding table entry %d, symbol %d >= max (%d)", i, v.addBits(), len(t))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:331
			// _ = "end of CoverTab[94510]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:332
			_go_fuzz_dep_.CoverTab[94511]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:332
			// _ = "end of CoverTab[94511]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:332
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:332
		// _ = "end of CoverTab[94507]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:332
		_go_fuzz_dep_.CoverTab[94508]++
														lu := t[add]
														if lu.addBits > s.maxBits {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:334
			_go_fuzz_dep_.CoverTab[94512]++
															s.maxBits = lu.addBits
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:335
			// _ = "end of CoverTab[94512]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:336
			_go_fuzz_dep_.CoverTab[94513]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:336
			// _ = "end of CoverTab[94513]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:336
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:336
		// _ = "end of CoverTab[94508]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:336
		_go_fuzz_dep_.CoverTab[94509]++
														v.setExt(lu.addBits, lu.baseLine)
														s.dt[i] = v
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:338
		// _ = "end of CoverTab[94509]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:339
	// _ = "end of CoverTab[94505]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:339
	_go_fuzz_dep_.CoverTab[94506]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:340
	// _ = "end of CoverTab[94506]"
}

type fseState struct {
	dt	[]decSymbol
	state	decSymbol
}

// Initialize and decodeAsync first state and symbol.
func (s *fseState) init(br *bitReader, tableLog uint8, dt []decSymbol) {
	s.dt = dt
	br.fill()
	s.state = dt[br.getBits(tableLog)]
}

// next returns the current symbol and sets the next state.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:355
// At least tablelog bits must be available in the bit reader.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:357
func (s *fseState) next(br *bitReader) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:357
	_go_fuzz_dep_.CoverTab[94514]++
													lowBits := uint16(br.getBits(s.state.nbBits()))
													s.state = s.dt[s.state.newState()+lowBits]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:359
	// _ = "end of CoverTab[94514]"
}

// finished returns true if all bits have been read from the bitstream
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:362
// and the next state would require reading bits from the input.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:364
func (s *fseState) finished(br *bitReader) bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:364
	_go_fuzz_dep_.CoverTab[94515]++
													return br.finished() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:365
		_go_fuzz_dep_.CoverTab[94516]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:365
		return s.state.nbBits() > 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:365
		// _ = "end of CoverTab[94516]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:365
	}()
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:365
	// _ = "end of CoverTab[94515]"
}

// final returns the current state symbol without decoding the next.
func (s *fseState) final() (int, uint8) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:369
	_go_fuzz_dep_.CoverTab[94517]++
													return s.state.baselineInt(), s.state.addBits()
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:370
	// _ = "end of CoverTab[94517]"
}

// final returns the current state symbol without decoding the next.
func (s decSymbol) final() (int, uint8) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:374
	_go_fuzz_dep_.CoverTab[94518]++
													return s.baselineInt(), s.addBits()
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:375
	// _ = "end of CoverTab[94518]"
}

// nextFast returns the next symbol and sets the next state.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:378
// This can only be used if no symbols are 0 bits.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:378
// At least tablelog bits must be available in the bit reader.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:381
func (s *fseState) nextFast(br *bitReader) (uint32, uint8) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:381
	_go_fuzz_dep_.CoverTab[94519]++
													lowBits := br.get16BitsFast(s.state.nbBits())
													s.state = s.dt[s.state.newState()+lowBits]
													return s.state.baseline(), s.state.addBits()
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:384
	// _ = "end of CoverTab[94519]"
}

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:385
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_decoder.go:385
var _ = _go_fuzz_dep_.CoverTab
