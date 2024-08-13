// Copyright 2018 Klaus Post. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Based on work Copyright (c) 2013, Yann Collet, released under BSD License.

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:6
// Package fse provides Finite State Entropy encoding and decoding.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:6
//
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:6
// Finite State Entropy encoding provides a fast near-optimal symbol encoding/decoding
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:6
// for byte blocks as implemented in zstd.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:6
//
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:6
// See https://github.com/klauspost/compress/tree/master/fse for more information.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:12
package fse

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:12
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:12
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:12
)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:12
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:12
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:12
)

import (
	"errors"
	"fmt"
	"math/bits"
)

const (
	/*!MEMORY_USAGE :
	 *  Memory usage formula : N->2^N Bytes (examples : 10 -> 1KB; 12 -> 4KB ; 16 -> 64KB; 20 -> 1MB; etc.)
	 *  Increasing memory usage improves compression ratio
	 *  Reduced memory usage can improve speed, due to cache effect
	 *  Recommended max value is 14, for 16KB, which nicely fits into Intel x86 L1 cache */
	maxMemoryUsage		= 14
	defaultMemoryUsage	= 13

	maxTableLog	= maxMemoryUsage - 2
	maxTablesize	= 1 << maxTableLog
	defaultTablelog	= defaultMemoryUsage - 2
	minTablelog	= 5
	maxSymbolValue	= 255
)

var (
	// ErrIncompressible is returned when input is judged to be too hard to compress.
	ErrIncompressible	= errors.New("input is not compressible")

	// ErrUseRLE is returned from the compressor when the input is a single byte value repeated.
	ErrUseRLE	= errors.New("input is single value repeated")
)

// Scratch provides temporary storage for compression and decompression.
type Scratch struct {
	// Private
	count		[maxSymbolValue + 1]uint32
	norm		[maxSymbolValue + 1]int16
	br		byteReader
	bits		bitReader
	bw		bitWriter
	ct		cTable		// Compression tables.
	decTable	[]decSymbol	// Decompression table.
	maxCount	int		// count of the most probable symbol

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:60
	// Out is output buffer.
	// If the scratch is re-used before the caller is done processing the output,
	// set this field to nil.
	// Otherwise the output buffer will be re-used for next Compression/Decompression step
	// and allocation will be avoided.
	Out	[]byte

	// DecompressLimit limits the maximum decoded size acceptable.
	// If > 0 decompression will stop when approximately this many bytes
	// has been decoded.
	// If 0, maximum size will be 2GB.
	DecompressLimit	int

	symbolLen	uint16	// Length of active part of the symbol table.
	actualTableLog	uint8	// Selected tablelog.
	zeroBits	bool	// no bits has prob > 50%.
	clearCount	bool	// clear count

	// MaxSymbolValue will override the maximum symbol value of the next block.
	MaxSymbolValue	uint8

	// TableLog will attempt to override the tablelog for the next block.
	TableLog	uint8
}

// Histogram allows to populate the histogram and skip that step in the compression,
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:85
// It otherwise allows to inspect the histogram when compression is done.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:85
// To indicate that you have populated the histogram call HistogramFinished
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:85
// with the value of the highest populated symbol, as well as the number of entries
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:85
// in the most populated entry. These are accepted at face value.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:85
// The returned slice will always be length 256.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:91
func (s *Scratch) Histogram() []uint32 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:91
	_go_fuzz_dep_.CoverTab[89396]++
											return s.count[:]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:92
	// _ = "end of CoverTab[89396]"
}

// HistogramFinished can be called to indicate that the histogram has been populated.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:95
// maxSymbol is the index of the highest set symbol of the next data segment.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:95
// maxCount is the number of entries in the most populated entry.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:95
// These are accepted at face value.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:99
func (s *Scratch) HistogramFinished(maxSymbol uint8, maxCount int) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:99
	_go_fuzz_dep_.CoverTab[89397]++
											s.maxCount = maxCount
											s.symbolLen = uint16(maxSymbol) + 1
											s.clearCount = maxCount != 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:102
	// _ = "end of CoverTab[89397]"
}

// prepare will prepare and allocate scratch tables used for both compression and decompression.
func (s *Scratch) prepare(in []byte) (*Scratch, error) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:106
	_go_fuzz_dep_.CoverTab[89398]++
											if s == nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:107
		_go_fuzz_dep_.CoverTab[89406]++
												s = &Scratch{}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:108
		// _ = "end of CoverTab[89406]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:109
		_go_fuzz_dep_.CoverTab[89407]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:109
		// _ = "end of CoverTab[89407]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:109
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:109
	// _ = "end of CoverTab[89398]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:109
	_go_fuzz_dep_.CoverTab[89399]++
											if s.MaxSymbolValue == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:110
		_go_fuzz_dep_.CoverTab[89408]++
												s.MaxSymbolValue = 255
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:111
		// _ = "end of CoverTab[89408]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:112
		_go_fuzz_dep_.CoverTab[89409]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:112
		// _ = "end of CoverTab[89409]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:112
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:112
	// _ = "end of CoverTab[89399]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:112
	_go_fuzz_dep_.CoverTab[89400]++
											if s.TableLog == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:113
		_go_fuzz_dep_.CoverTab[89410]++
												s.TableLog = defaultTablelog
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:114
		// _ = "end of CoverTab[89410]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:115
		_go_fuzz_dep_.CoverTab[89411]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:115
		// _ = "end of CoverTab[89411]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:115
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:115
	// _ = "end of CoverTab[89400]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:115
	_go_fuzz_dep_.CoverTab[89401]++
											if s.TableLog > maxTableLog {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:116
		_go_fuzz_dep_.CoverTab[89412]++
												return nil, fmt.Errorf("tableLog (%d) > maxTableLog (%d)", s.TableLog, maxTableLog)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:117
		// _ = "end of CoverTab[89412]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:118
		_go_fuzz_dep_.CoverTab[89413]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:118
		// _ = "end of CoverTab[89413]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:118
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:118
	// _ = "end of CoverTab[89401]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:118
	_go_fuzz_dep_.CoverTab[89402]++
											if cap(s.Out) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:119
		_go_fuzz_dep_.CoverTab[89414]++
												s.Out = make([]byte, 0, len(in))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:120
		// _ = "end of CoverTab[89414]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:121
		_go_fuzz_dep_.CoverTab[89415]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:121
		// _ = "end of CoverTab[89415]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:121
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:121
	// _ = "end of CoverTab[89402]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:121
	_go_fuzz_dep_.CoverTab[89403]++
											if s.clearCount && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:122
		_go_fuzz_dep_.CoverTab[89416]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:122
		return s.maxCount == 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:122
		// _ = "end of CoverTab[89416]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:122
	}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:122
		_go_fuzz_dep_.CoverTab[89417]++
												for i := range s.count {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:123
			_go_fuzz_dep_.CoverTab[89419]++
													s.count[i] = 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:124
			// _ = "end of CoverTab[89419]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:125
		// _ = "end of CoverTab[89417]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:125
		_go_fuzz_dep_.CoverTab[89418]++
												s.clearCount = false
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:126
		// _ = "end of CoverTab[89418]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:127
		_go_fuzz_dep_.CoverTab[89420]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:127
		// _ = "end of CoverTab[89420]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:127
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:127
	// _ = "end of CoverTab[89403]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:127
	_go_fuzz_dep_.CoverTab[89404]++
											s.br.init(in)
											if s.DecompressLimit == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:129
		_go_fuzz_dep_.CoverTab[89421]++

												s.DecompressLimit = (2 << 30) - 1
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:131
		// _ = "end of CoverTab[89421]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:132
		_go_fuzz_dep_.CoverTab[89422]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:132
		// _ = "end of CoverTab[89422]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:132
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:132
	// _ = "end of CoverTab[89404]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:132
	_go_fuzz_dep_.CoverTab[89405]++

											return s, nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:134
	// _ = "end of CoverTab[89405]"
}

// tableStep returns the next table index.
func tableStep(tableSize uint32) uint32 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:138
	_go_fuzz_dep_.CoverTab[89423]++
											return (tableSize >> 1) + (tableSize >> 3) + 3
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:139
	// _ = "end of CoverTab[89423]"
}

func highBits(val uint32) (n uint32) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:142
	_go_fuzz_dep_.CoverTab[89424]++
											return uint32(bits.Len32(val) - 1)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:143
	// _ = "end of CoverTab[89424]"
}

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:144
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/fse.go:144
var _ = _go_fuzz_dep_.CoverTab
