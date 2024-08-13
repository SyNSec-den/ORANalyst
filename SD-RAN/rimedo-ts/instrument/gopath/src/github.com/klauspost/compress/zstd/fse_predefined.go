// Copyright 2019+ Klaus Post. All rights reserved.
// License information can be found in the LICENSE file.
// Based on work by Yann Collet, released under BSD License.

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:5
package zstd

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:5
)

import (
	"fmt"
	"math"
	"sync"
)

var (
	// fsePredef are the predefined fse tables as defined here:
	// https://github.com/facebook/zstd/blob/dev/doc/zstd_compression_format.md#default-distributions
	// These values are already transformed.
	fsePredef	[3]fseDecoder

	// fsePredefEnc are the predefined encoder based on fse tables as defined here:
	// https://github.com/facebook/zstd/blob/dev/doc/zstd_compression_format.md#default-distributions
	// These values are already transformed.
	fsePredefEnc	[3]fseEncoder

	// symbolTableX contain the transformations needed for each type as defined in
	// https://github.com/facebook/zstd/blob/dev/doc/zstd_compression_format.md#the-codes-for-literals-lengths-match-lengths-and-offsets
	symbolTableX	[3][]baseOffset

	// maxTableSymbol is the biggest supported symbol for each table type
	// https://github.com/facebook/zstd/blob/dev/doc/zstd_compression_format.md#the-codes-for-literals-lengths-match-lengths-and-offsets
	maxTableSymbol	= [3]uint8{tableLiteralLengths: maxLiteralLengthSymbol, tableOffsets: maxOffsetLengthSymbol, tableMatchLengths: maxMatchLengthSymbol}

	// bitTables is the bits table for each table.
	bitTables	= [3][]byte{tableLiteralLengths: llBitsTable[:], tableOffsets: nil, tableMatchLengths: mlBitsTable[:]}
)

type tableIndex uint8

const (
	// indexes for fsePredef and symbolTableX
	tableLiteralLengths	tableIndex	= 0
	tableOffsets		tableIndex	= 1
	tableMatchLengths	tableIndex	= 2

	maxLiteralLengthSymbol	= 35
	maxOffsetLengthSymbol	= 30
	maxMatchLengthSymbol	= 52
)

// baseOffset is used for calculating transformations.
type baseOffset struct {
	baseLine	uint32
	addBits		uint8
}

// fillBase will precalculate base offsets with the given bit distributions.
func fillBase(dst []baseOffset, base uint32, bits ...uint8) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:56
	_go_fuzz_dep_.CoverTab[94804]++
													if len(bits) != len(dst) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:57
		_go_fuzz_dep_.CoverTab[94806]++
														panic(fmt.Sprintf("len(dst) (%d) != len(bits) (%d)", len(dst), len(bits)))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:58
		// _ = "end of CoverTab[94806]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:59
		_go_fuzz_dep_.CoverTab[94807]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:59
		// _ = "end of CoverTab[94807]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:59
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:59
	// _ = "end of CoverTab[94804]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:59
	_go_fuzz_dep_.CoverTab[94805]++
													for i, bit := range bits {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:60
		_go_fuzz_dep_.CoverTab[94808]++
														if base > math.MaxInt32 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:61
			_go_fuzz_dep_.CoverTab[94810]++
															panic("invalid decoding table, base overflows int32")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:62
			// _ = "end of CoverTab[94810]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:63
			_go_fuzz_dep_.CoverTab[94811]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:63
			// _ = "end of CoverTab[94811]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:63
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:63
		// _ = "end of CoverTab[94808]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:63
		_go_fuzz_dep_.CoverTab[94809]++

														dst[i] = baseOffset{
			baseLine:	base,
			addBits:	bit,
		}
														base += 1 << bit
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:69
		// _ = "end of CoverTab[94809]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:70
	// _ = "end of CoverTab[94805]"
}

var predef sync.Once

func initPredefined() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:75
	_go_fuzz_dep_.CoverTab[94812]++
													predef.Do(func() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:76
		_go_fuzz_dep_.CoverTab[94813]++

														tmp := make([]baseOffset, 36)
														for i := range tmp[:16] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:79
			_go_fuzz_dep_.CoverTab[94816]++
															tmp[i] = baseOffset{
				baseLine:	uint32(i),
				addBits:	0,
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:83
			// _ = "end of CoverTab[94816]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:84
		// _ = "end of CoverTab[94813]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:84
		_go_fuzz_dep_.CoverTab[94814]++
														fillBase(tmp[16:], 16, 1, 1, 1, 1, 2, 2, 3, 3, 4, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16)
														symbolTableX[tableLiteralLengths] = tmp

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:89
		tmp = make([]baseOffset, 53)
		for i := range tmp[:32] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:90
			_go_fuzz_dep_.CoverTab[94817]++
															tmp[i] = baseOffset{

				baseLine:	uint32(i) + 3,
				addBits:	0,
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:95
			// _ = "end of CoverTab[94817]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:96
		// _ = "end of CoverTab[94814]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:96
		_go_fuzz_dep_.CoverTab[94815]++
														fillBase(tmp[32:], 35, 1, 1, 1, 1, 2, 2, 3, 3, 4, 4, 5, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16)
														symbolTableX[tableMatchLengths] = tmp

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:101
		tmp = make([]baseOffset, maxOffsetBits+1)
		tmp[1] = baseOffset{
			baseLine:	1,
			addBits:	1,
		}
														fillBase(tmp[2:], 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30)
														symbolTableX[tableOffsets] = tmp

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:111
		for i := range fsePredef[:] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:111
			_go_fuzz_dep_.CoverTab[94818]++
															f := &fsePredef[i]
															switch tableIndex(i) {
			case tableLiteralLengths:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:114
				_go_fuzz_dep_.CoverTab[94823]++

																f.actualTableLog = 6
																copy(f.norm[:], []int16{4, 3, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1, 1, 1,
					2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 2, 1, 1, 1, 1, 1,
					-1, -1, -1, -1})
																f.symbolLen = 36
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:120
				// _ = "end of CoverTab[94823]"
			case tableOffsets:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:121
				_go_fuzz_dep_.CoverTab[94824]++

																f.actualTableLog = 5
																copy(f.norm[:], []int16{
					1, 1, 1, 1, 1, 1, 2, 2, 2, 1, 1, 1, 1, 1, 1, 1,
					1, 1, 1, 1, 1, 1, 1, 1, -1, -1, -1, -1, -1})
																f.symbolLen = 29
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:127
				// _ = "end of CoverTab[94824]"
			case tableMatchLengths:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:128
				_go_fuzz_dep_.CoverTab[94825]++

																f.actualTableLog = 6
																copy(f.norm[:], []int16{
					1, 4, 3, 2, 2, 2, 2, 2, 2, 1, 1, 1, 1, 1, 1, 1,
					1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
					1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, -1, -1,
					-1, -1, -1, -1, -1})
																f.symbolLen = 53
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:136
				// _ = "end of CoverTab[94825]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:136
			default:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:136
				_go_fuzz_dep_.CoverTab[94826]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:136
				// _ = "end of CoverTab[94826]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:137
			// _ = "end of CoverTab[94818]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:137
			_go_fuzz_dep_.CoverTab[94819]++
															if err := f.buildDtable(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:138
				_go_fuzz_dep_.CoverTab[94827]++
																panic(fmt.Errorf("building table %v: %v", tableIndex(i), err))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:139
				// _ = "end of CoverTab[94827]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:140
				_go_fuzz_dep_.CoverTab[94828]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:140
				// _ = "end of CoverTab[94828]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:140
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:140
			// _ = "end of CoverTab[94819]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:140
			_go_fuzz_dep_.CoverTab[94820]++
															if err := f.transform(symbolTableX[i]); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:141
				_go_fuzz_dep_.CoverTab[94829]++
																panic(fmt.Errorf("building table %v: %v", tableIndex(i), err))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:142
				// _ = "end of CoverTab[94829]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:143
				_go_fuzz_dep_.CoverTab[94830]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:143
				// _ = "end of CoverTab[94830]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:143
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:143
			// _ = "end of CoverTab[94820]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:143
			_go_fuzz_dep_.CoverTab[94821]++
															f.preDefined = true

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:147
			enc := &fsePredefEnc[i]
			copy(enc.norm[:], f.norm[:])
			enc.symbolLen = f.symbolLen
			enc.actualTableLog = f.actualTableLog
			if err := enc.buildCTable(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:151
				_go_fuzz_dep_.CoverTab[94831]++
																panic(fmt.Errorf("building encoding table %v: %v", tableIndex(i), err))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:152
				// _ = "end of CoverTab[94831]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:153
				_go_fuzz_dep_.CoverTab[94832]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:153
				// _ = "end of CoverTab[94832]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:153
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:153
			// _ = "end of CoverTab[94821]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:153
			_go_fuzz_dep_.CoverTab[94822]++
															enc.setBits(bitTables[i])
															enc.preDefined = true
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:155
			// _ = "end of CoverTab[94822]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:156
		// _ = "end of CoverTab[94815]"
	})
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:157
	// _ = "end of CoverTab[94812]"
}

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:158
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_predefined.go:158
var _ = _go_fuzz_dep_.CoverTab
