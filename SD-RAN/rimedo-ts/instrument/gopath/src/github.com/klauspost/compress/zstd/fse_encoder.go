// Copyright 2019+ Klaus Post. All rights reserved.
// License information can be found in the LICENSE file.
// Based on work by Yann Collet, released under BSD License.

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:5
package zstd

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:5
)

import (
	"errors"
	"fmt"
	"math"
)

const (
	// For encoding we only support up to
	maxEncTableLog		= 8
	maxEncTablesize		= 1 << maxTableLog
	maxEncTableMask		= (1 << maxTableLog) - 1
	minEncTablelog		= 5
	maxEncSymbolValue	= maxMatchLengthSymbol
)

// Scratch provides temporary storage for compression and decompression.
type fseEncoder struct {
	symbolLen	uint16	// Length of active part of the symbol table.
	actualTableLog	uint8	// Selected tablelog.
	ct		cTable	// Compression tables.
	maxCount	int	// count of the most probable symbol
	zeroBits	bool	// no bits has prob > 50%.
	clearCount	bool	// clear count
	useRLE		bool	// This encoder is for RLE
	preDefined	bool	// This encoder is predefined.
	reUsed		bool	// Set to know when the encoder has been reused.
	rleVal		uint8	// RLE Symbol
	maxBits		uint8	// Maximum output bits after transform.

	// TODO: Technically zstd should be fine with 64 bytes.
	count	[256]uint32
	norm	[256]int16
}

// cTable contains tables used for compression.
type cTable struct {
	tableSymbol	[]byte
	stateTable	[]uint16
	symbolTT	[]symbolTransform
}

// symbolTransform contains the state transform for a symbol.
type symbolTransform struct {
	deltaNbBits	uint32
	deltaFindState	int16
	outBits		uint8
}

// String prints values as a human readable string.
func (s symbolTransform) String() string {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:56
	_go_fuzz_dep_.CoverTab[94520]++
												return fmt.Sprintf("{deltabits: %08x, findstate:%d outbits:%d}", s.deltaNbBits, s.deltaFindState, s.outBits)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:57
	// _ = "end of CoverTab[94520]"
}

// Histogram allows to populate the histogram and skip that step in the compression,
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:60
// It otherwise allows to inspect the histogram when compression is done.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:60
// To indicate that you have populated the histogram call HistogramFinished
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:60
// with the value of the highest populated symbol, as well as the number of entries
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:60
// in the most populated entry. These are accepted at face value.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:65
func (s *fseEncoder) Histogram() *[256]uint32 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:65
	_go_fuzz_dep_.CoverTab[94521]++
												return &s.count
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:66
	// _ = "end of CoverTab[94521]"
}

// HistogramFinished can be called to indicate that the histogram has been populated.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:69
// maxSymbol is the index of the highest set symbol of the next data segment.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:69
// maxCount is the number of entries in the most populated entry.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:69
// These are accepted at face value.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:73
func (s *fseEncoder) HistogramFinished(maxSymbol uint8, maxCount int) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:73
	_go_fuzz_dep_.CoverTab[94522]++
												s.maxCount = maxCount
												s.symbolLen = uint16(maxSymbol) + 1
												s.clearCount = maxCount != 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:76
	// _ = "end of CoverTab[94522]"
}

// prepare will prepare and allocate scratch tables used for both compression and decompression.
func (s *fseEncoder) prepare() (*fseEncoder, error) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:80
	_go_fuzz_dep_.CoverTab[94523]++
												if s == nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:81
		_go_fuzz_dep_.CoverTab[94526]++
													s = &fseEncoder{}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:82
		// _ = "end of CoverTab[94526]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:83
		_go_fuzz_dep_.CoverTab[94527]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:83
		// _ = "end of CoverTab[94527]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:83
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:83
	// _ = "end of CoverTab[94523]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:83
	_go_fuzz_dep_.CoverTab[94524]++
												s.useRLE = false
												if s.clearCount && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:85
		_go_fuzz_dep_.CoverTab[94528]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:85
		return s.maxCount == 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:85
		// _ = "end of CoverTab[94528]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:85
	}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:85
		_go_fuzz_dep_.CoverTab[94529]++
													for i := range s.count {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:86
			_go_fuzz_dep_.CoverTab[94531]++
														s.count[i] = 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:87
			// _ = "end of CoverTab[94531]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:88
		// _ = "end of CoverTab[94529]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:88
		_go_fuzz_dep_.CoverTab[94530]++
													s.clearCount = false
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:89
		// _ = "end of CoverTab[94530]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:90
		_go_fuzz_dep_.CoverTab[94532]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:90
		// _ = "end of CoverTab[94532]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:90
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:90
	// _ = "end of CoverTab[94524]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:90
	_go_fuzz_dep_.CoverTab[94525]++
												return s, nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:91
	// _ = "end of CoverTab[94525]"
}

// allocCtable will allocate tables needed for compression.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:94
// If existing tables a re big enough, they are simply re-used.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:96
func (s *fseEncoder) allocCtable() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:96
	_go_fuzz_dep_.CoverTab[94533]++
												tableSize := 1 << s.actualTableLog

												if cap(s.ct.tableSymbol) < tableSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:99
			_go_fuzz_dep_.CoverTab[94537]++
														s.ct.tableSymbol = make([]byte, tableSize)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:100
		// _ = "end of CoverTab[94537]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:101
		_go_fuzz_dep_.CoverTab[94538]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:101
		// _ = "end of CoverTab[94538]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:101
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:101
	// _ = "end of CoverTab[94533]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:101
	_go_fuzz_dep_.CoverTab[94534]++
													s.ct.tableSymbol = s.ct.tableSymbol[:tableSize]

													ctSize := tableSize
													if cap(s.ct.stateTable) < ctSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:105
		_go_fuzz_dep_.CoverTab[94539]++
														s.ct.stateTable = make([]uint16, ctSize)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:106
		// _ = "end of CoverTab[94539]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:107
		_go_fuzz_dep_.CoverTab[94540]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:107
		// _ = "end of CoverTab[94540]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:107
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:107
	// _ = "end of CoverTab[94534]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:107
	_go_fuzz_dep_.CoverTab[94535]++
													s.ct.stateTable = s.ct.stateTable[:ctSize]

													if cap(s.ct.symbolTT) < 256 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:110
		_go_fuzz_dep_.CoverTab[94541]++
														s.ct.symbolTT = make([]symbolTransform, 256)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:111
		// _ = "end of CoverTab[94541]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:112
		_go_fuzz_dep_.CoverTab[94542]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:112
		// _ = "end of CoverTab[94542]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:112
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:112
	// _ = "end of CoverTab[94535]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:112
	_go_fuzz_dep_.CoverTab[94536]++
													s.ct.symbolTT = s.ct.symbolTT[:256]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:113
	// _ = "end of CoverTab[94536]"
}

// buildCTable will populate the compression table so it is ready to be used.
func (s *fseEncoder) buildCTable() error {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:117
	_go_fuzz_dep_.CoverTab[94543]++
													tableSize := uint32(1 << s.actualTableLog)
													highThreshold := tableSize - 1
													var cumul [256]int16

													s.allocCtable()
													tableSymbol := s.ct.tableSymbol[:tableSize]

													{
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:125
		_go_fuzz_dep_.CoverTab[94547]++
														cumul[0] = 0
														for ui, v := range s.norm[:s.symbolLen-1] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:127
			_go_fuzz_dep_.CoverTab[94551]++
															u := byte(ui)
															if v == -1 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:129
				_go_fuzz_dep_.CoverTab[94552]++

																cumul[u+1] = cumul[u] + 1
																tableSymbol[highThreshold] = u
																highThreshold--
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:133
				// _ = "end of CoverTab[94552]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:134
				_go_fuzz_dep_.CoverTab[94553]++
																cumul[u+1] = cumul[u] + v
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:135
				// _ = "end of CoverTab[94553]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:136
			// _ = "end of CoverTab[94551]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:137
		// _ = "end of CoverTab[94547]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:137
		_go_fuzz_dep_.CoverTab[94548]++

														u := int(s.symbolLen - 1)
														v := s.norm[s.symbolLen-1]
														if v == -1 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:141
			_go_fuzz_dep_.CoverTab[94554]++

															cumul[u+1] = cumul[u] + 1
															tableSymbol[highThreshold] = byte(u)
															highThreshold--
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:145
			// _ = "end of CoverTab[94554]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:146
			_go_fuzz_dep_.CoverTab[94555]++
															cumul[u+1] = cumul[u] + v
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:147
			// _ = "end of CoverTab[94555]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:148
		// _ = "end of CoverTab[94548]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:148
		_go_fuzz_dep_.CoverTab[94549]++
														if uint32(cumul[s.symbolLen]) != tableSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:149
			_go_fuzz_dep_.CoverTab[94556]++
															return fmt.Errorf("internal error: expected cumul[s.symbolLen] (%d) == tableSize (%d)", cumul[s.symbolLen], tableSize)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:150
			// _ = "end of CoverTab[94556]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:151
			_go_fuzz_dep_.CoverTab[94557]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:151
			// _ = "end of CoverTab[94557]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:151
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:151
		// _ = "end of CoverTab[94549]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:151
		_go_fuzz_dep_.CoverTab[94550]++
														cumul[s.symbolLen] = int16(tableSize) + 1
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:152
		// _ = "end of CoverTab[94550]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:153
	// _ = "end of CoverTab[94543]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:153
	_go_fuzz_dep_.CoverTab[94544]++

													s.zeroBits = false
													{
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:156
		_go_fuzz_dep_.CoverTab[94558]++
														step := tableStep(tableSize)
														tableMask := tableSize - 1
														var position uint32

														largeLimit := int16(1 << (s.actualTableLog - 1))
														for ui, v := range s.norm[:s.symbolLen] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:162
			_go_fuzz_dep_.CoverTab[94560]++
															symbol := byte(ui)
															if v > largeLimit {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:164
				_go_fuzz_dep_.CoverTab[94562]++
																s.zeroBits = true
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:165
				// _ = "end of CoverTab[94562]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:166
				_go_fuzz_dep_.CoverTab[94563]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:166
				// _ = "end of CoverTab[94563]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:166
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:166
			// _ = "end of CoverTab[94560]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:166
			_go_fuzz_dep_.CoverTab[94561]++
															for nbOccurrences := int16(0); nbOccurrences < v; nbOccurrences++ {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:167
				_go_fuzz_dep_.CoverTab[94564]++
																tableSymbol[position] = symbol
																position = (position + step) & tableMask
																for position > highThreshold {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:170
					_go_fuzz_dep_.CoverTab[94565]++
																	position = (position + step) & tableMask
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:171
					// _ = "end of CoverTab[94565]"
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:172
				// _ = "end of CoverTab[94564]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:173
			// _ = "end of CoverTab[94561]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:174
		// _ = "end of CoverTab[94558]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:174
		_go_fuzz_dep_.CoverTab[94559]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:177
		if position != 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:177
			_go_fuzz_dep_.CoverTab[94566]++
															return errors.New("position!=0")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:178
			// _ = "end of CoverTab[94566]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:179
			_go_fuzz_dep_.CoverTab[94567]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:179
			// _ = "end of CoverTab[94567]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:179
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:179
		// _ = "end of CoverTab[94559]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:180
	// _ = "end of CoverTab[94544]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:180
	_go_fuzz_dep_.CoverTab[94545]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:183
	table := s.ct.stateTable
	{
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:184
		_go_fuzz_dep_.CoverTab[94568]++
														tsi := int(tableSize)
														for u, v := range tableSymbol {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:186
			_go_fuzz_dep_.CoverTab[94569]++

															table[cumul[v]] = uint16(tsi + u)
															cumul[v]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:189
			// _ = "end of CoverTab[94569]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:190
		// _ = "end of CoverTab[94568]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:191
	// _ = "end of CoverTab[94545]"

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:194
	{
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:194
		_go_fuzz_dep_.CoverTab[94570]++
														total := int16(0)
														symbolTT := s.ct.symbolTT[:s.symbolLen]
														tableLog := s.actualTableLog
														tl := (uint32(tableLog) << 16) - (1 << tableLog)
														for i, v := range s.norm[:s.symbolLen] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:199
			_go_fuzz_dep_.CoverTab[94572]++
															switch v {
			case 0:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:201
				_go_fuzz_dep_.CoverTab[94573]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:201
				// _ = "end of CoverTab[94573]"
			case -1, 1:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:202
				_go_fuzz_dep_.CoverTab[94574]++
																symbolTT[i].deltaNbBits = tl
																symbolTT[i].deltaFindState = total - 1
																total++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:205
				// _ = "end of CoverTab[94574]"
			default:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:206
				_go_fuzz_dep_.CoverTab[94575]++
																maxBitsOut := uint32(tableLog) - highBit(uint32(v-1))
																minStatePlus := uint32(v) << maxBitsOut
																symbolTT[i].deltaNbBits = (maxBitsOut << 16) - minStatePlus
																symbolTT[i].deltaFindState = total - v
																total += v
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:211
				// _ = "end of CoverTab[94575]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:212
			// _ = "end of CoverTab[94572]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:213
		// _ = "end of CoverTab[94570]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:213
		_go_fuzz_dep_.CoverTab[94571]++
														if total != int16(tableSize) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:214
			_go_fuzz_dep_.CoverTab[94576]++
															return fmt.Errorf("total mismatch %d (got) != %d (want)", total, tableSize)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:215
			// _ = "end of CoverTab[94576]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:216
			_go_fuzz_dep_.CoverTab[94577]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:216
			// _ = "end of CoverTab[94577]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:216
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:216
		// _ = "end of CoverTab[94571]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:217
	_go_fuzz_dep_.CoverTab[94546]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:218
	// _ = "end of CoverTab[94546]"
}

var rtbTable = [...]uint32{0, 473195, 504333, 520860, 550000, 700000, 750000, 830000}

func (s *fseEncoder) setRLE(val byte) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:223
	_go_fuzz_dep_.CoverTab[94578]++
													s.allocCtable()
													s.actualTableLog = 0
													s.ct.stateTable = s.ct.stateTable[:1]
													s.ct.symbolTT[val] = symbolTransform{
		deltaFindState:	0,
		deltaNbBits:	0,
	}
	if debugEncoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:231
		_go_fuzz_dep_.CoverTab[94580]++
														println("setRLE: val", val, "symbolTT", s.ct.symbolTT[val])
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:232
		// _ = "end of CoverTab[94580]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:233
		_go_fuzz_dep_.CoverTab[94581]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:233
		// _ = "end of CoverTab[94581]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:233
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:233
	// _ = "end of CoverTab[94578]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:233
	_go_fuzz_dep_.CoverTab[94579]++
													s.rleVal = val
													s.useRLE = true
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:235
	// _ = "end of CoverTab[94579]"
}

// setBits will set output bits for the transform.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:238
// if nil is provided, the number of bits is equal to the index.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:240
func (s *fseEncoder) setBits(transform []byte) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:240
	_go_fuzz_dep_.CoverTab[94582]++
													if s.reUsed || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:241
		_go_fuzz_dep_.CoverTab[94586]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:241
		return s.preDefined
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:241
		// _ = "end of CoverTab[94586]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:241
	}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:241
		_go_fuzz_dep_.CoverTab[94587]++
														return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:242
		// _ = "end of CoverTab[94587]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:243
		_go_fuzz_dep_.CoverTab[94588]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:243
		// _ = "end of CoverTab[94588]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:243
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:243
	// _ = "end of CoverTab[94582]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:243
	_go_fuzz_dep_.CoverTab[94583]++
													if s.useRLE {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:244
		_go_fuzz_dep_.CoverTab[94589]++
														if transform == nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:245
			_go_fuzz_dep_.CoverTab[94591]++
															s.ct.symbolTT[s.rleVal].outBits = s.rleVal
															s.maxBits = s.rleVal
															return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:248
			// _ = "end of CoverTab[94591]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:249
			_go_fuzz_dep_.CoverTab[94592]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:249
			// _ = "end of CoverTab[94592]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:249
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:249
		// _ = "end of CoverTab[94589]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:249
		_go_fuzz_dep_.CoverTab[94590]++
														s.maxBits = transform[s.rleVal]
														s.ct.symbolTT[s.rleVal].outBits = s.maxBits
														return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:252
		// _ = "end of CoverTab[94590]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:253
		_go_fuzz_dep_.CoverTab[94593]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:253
		// _ = "end of CoverTab[94593]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:253
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:253
	// _ = "end of CoverTab[94583]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:253
	_go_fuzz_dep_.CoverTab[94584]++
													if transform == nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:254
		_go_fuzz_dep_.CoverTab[94594]++
														for i := range s.ct.symbolTT[:s.symbolLen] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:255
			_go_fuzz_dep_.CoverTab[94596]++
															s.ct.symbolTT[i].outBits = uint8(i)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:256
			// _ = "end of CoverTab[94596]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:257
		// _ = "end of CoverTab[94594]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:257
		_go_fuzz_dep_.CoverTab[94595]++
														s.maxBits = uint8(s.symbolLen - 1)
														return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:259
		// _ = "end of CoverTab[94595]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:260
		_go_fuzz_dep_.CoverTab[94597]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:260
		// _ = "end of CoverTab[94597]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:260
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:260
	// _ = "end of CoverTab[94584]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:260
	_go_fuzz_dep_.CoverTab[94585]++
													s.maxBits = 0
													for i, v := range transform[:s.symbolLen] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:262
		_go_fuzz_dep_.CoverTab[94598]++
														s.ct.symbolTT[i].outBits = v
														if v > s.maxBits {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:264
			_go_fuzz_dep_.CoverTab[94599]++

															s.maxBits = v
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:266
			// _ = "end of CoverTab[94599]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:267
			_go_fuzz_dep_.CoverTab[94600]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:267
			// _ = "end of CoverTab[94600]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:267
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:267
		// _ = "end of CoverTab[94598]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:268
	// _ = "end of CoverTab[94585]"
}

// normalizeCount will normalize the count of the symbols so
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:271
// the total is equal to the table size.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:271
// If successful, compression tables will also be made ready.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:274
func (s *fseEncoder) normalizeCount(length int) error {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:274
	_go_fuzz_dep_.CoverTab[94601]++
													if s.reUsed {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:275
		_go_fuzz_dep_.CoverTab[94607]++
														return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:276
		// _ = "end of CoverTab[94607]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:277
		_go_fuzz_dep_.CoverTab[94608]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:277
		// _ = "end of CoverTab[94608]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:277
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:277
	// _ = "end of CoverTab[94601]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:277
	_go_fuzz_dep_.CoverTab[94602]++
													s.optimalTableLog(length)
													var (
		tableLog		= s.actualTableLog
		scale			= 62 - uint64(tableLog)
		step			= (1 << 62) / uint64(length)
		vStep			= uint64(1) << (scale - 20)
		stillToDistribute	= int16(1 << tableLog)
		largest			int
		largestP		int16
		lowThreshold		= (uint32)(length >> tableLog)
	)
	if s.maxCount == length {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:289
		_go_fuzz_dep_.CoverTab[94609]++
														s.useRLE = true
														return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:291
		// _ = "end of CoverTab[94609]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:292
		_go_fuzz_dep_.CoverTab[94610]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:292
		// _ = "end of CoverTab[94610]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:292
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:292
	// _ = "end of CoverTab[94602]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:292
	_go_fuzz_dep_.CoverTab[94603]++
													s.useRLE = false
													for i, cnt := range s.count[:s.symbolLen] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:294
		_go_fuzz_dep_.CoverTab[94611]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:298
		if cnt == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:298
			_go_fuzz_dep_.CoverTab[94613]++
															s.norm[i] = 0
															continue
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:300
			// _ = "end of CoverTab[94613]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:301
			_go_fuzz_dep_.CoverTab[94614]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:301
			// _ = "end of CoverTab[94614]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:301
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:301
		// _ = "end of CoverTab[94611]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:301
		_go_fuzz_dep_.CoverTab[94612]++
														if cnt <= lowThreshold {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:302
			_go_fuzz_dep_.CoverTab[94615]++
															s.norm[i] = -1
															stillToDistribute--
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:304
			// _ = "end of CoverTab[94615]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:305
			_go_fuzz_dep_.CoverTab[94616]++
															proba := (int16)((uint64(cnt) * step) >> scale)
															if proba < 8 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:307
				_go_fuzz_dep_.CoverTab[94619]++
																restToBeat := vStep * uint64(rtbTable[proba])
																v := uint64(cnt)*step - (uint64(proba) << scale)
																if v > restToBeat {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:310
					_go_fuzz_dep_.CoverTab[94620]++
																	proba++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:311
					// _ = "end of CoverTab[94620]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:312
					_go_fuzz_dep_.CoverTab[94621]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:312
					// _ = "end of CoverTab[94621]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:312
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:312
				// _ = "end of CoverTab[94619]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:313
				_go_fuzz_dep_.CoverTab[94622]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:313
				// _ = "end of CoverTab[94622]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:313
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:313
			// _ = "end of CoverTab[94616]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:313
			_go_fuzz_dep_.CoverTab[94617]++
															if proba > largestP {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:314
				_go_fuzz_dep_.CoverTab[94623]++
																largestP = proba
																largest = i
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:316
				// _ = "end of CoverTab[94623]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:317
				_go_fuzz_dep_.CoverTab[94624]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:317
				// _ = "end of CoverTab[94624]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:317
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:317
			// _ = "end of CoverTab[94617]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:317
			_go_fuzz_dep_.CoverTab[94618]++
															s.norm[i] = proba
															stillToDistribute -= proba
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:319
			// _ = "end of CoverTab[94618]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:320
		// _ = "end of CoverTab[94612]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:321
	// _ = "end of CoverTab[94603]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:321
	_go_fuzz_dep_.CoverTab[94604]++

													if -stillToDistribute >= (s.norm[largest] >> 1) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:323
		_go_fuzz_dep_.CoverTab[94625]++

														err := s.normalizeCount2(length)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:326
			_go_fuzz_dep_.CoverTab[94628]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:327
			// _ = "end of CoverTab[94628]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:328
			_go_fuzz_dep_.CoverTab[94629]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:328
			// _ = "end of CoverTab[94629]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:328
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:328
		// _ = "end of CoverTab[94625]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:328
		_go_fuzz_dep_.CoverTab[94626]++
														if debugAsserts {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:329
			_go_fuzz_dep_.CoverTab[94630]++
															err = s.validateNorm()
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:331
				_go_fuzz_dep_.CoverTab[94631]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:332
				// _ = "end of CoverTab[94631]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:333
				_go_fuzz_dep_.CoverTab[94632]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:333
				// _ = "end of CoverTab[94632]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:333
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:333
			// _ = "end of CoverTab[94630]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:334
			_go_fuzz_dep_.CoverTab[94633]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:334
			// _ = "end of CoverTab[94633]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:334
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:334
		// _ = "end of CoverTab[94626]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:334
		_go_fuzz_dep_.CoverTab[94627]++
														return s.buildCTable()
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:335
		// _ = "end of CoverTab[94627]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:336
		_go_fuzz_dep_.CoverTab[94634]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:336
		// _ = "end of CoverTab[94634]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:336
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:336
	// _ = "end of CoverTab[94604]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:336
	_go_fuzz_dep_.CoverTab[94605]++
													s.norm[largest] += stillToDistribute
													if debugAsserts {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:338
		_go_fuzz_dep_.CoverTab[94635]++
														err := s.validateNorm()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:340
			_go_fuzz_dep_.CoverTab[94636]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:341
			// _ = "end of CoverTab[94636]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:342
			_go_fuzz_dep_.CoverTab[94637]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:342
			// _ = "end of CoverTab[94637]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:342
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:342
		// _ = "end of CoverTab[94635]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:343
		_go_fuzz_dep_.CoverTab[94638]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:343
		// _ = "end of CoverTab[94638]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:343
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:343
	// _ = "end of CoverTab[94605]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:343
	_go_fuzz_dep_.CoverTab[94606]++
													return s.buildCTable()
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:344
	// _ = "end of CoverTab[94606]"
}

// Secondary normalization method.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:347
// To be used when primary method fails.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:349
func (s *fseEncoder) normalizeCount2(length int) error {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:349
	_go_fuzz_dep_.CoverTab[94639]++
													const notYetAssigned = -2
													var (
		distributed	uint32
		total		= uint32(length)
		tableLog	= s.actualTableLog
		lowThreshold	= total >> tableLog
		lowOne		= (total * 3) >> (tableLog + 1)
	)
	for i, cnt := range s.count[:s.symbolLen] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:358
		_go_fuzz_dep_.CoverTab[94645]++
														if cnt == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:359
			_go_fuzz_dep_.CoverTab[94649]++
															s.norm[i] = 0
															continue
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:361
			// _ = "end of CoverTab[94649]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:362
			_go_fuzz_dep_.CoverTab[94650]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:362
			// _ = "end of CoverTab[94650]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:362
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:362
		// _ = "end of CoverTab[94645]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:362
		_go_fuzz_dep_.CoverTab[94646]++
														if cnt <= lowThreshold {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:363
			_go_fuzz_dep_.CoverTab[94651]++
															s.norm[i] = -1
															distributed++
															total -= cnt
															continue
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:367
			// _ = "end of CoverTab[94651]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:368
			_go_fuzz_dep_.CoverTab[94652]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:368
			// _ = "end of CoverTab[94652]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:368
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:368
		// _ = "end of CoverTab[94646]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:368
		_go_fuzz_dep_.CoverTab[94647]++
														if cnt <= lowOne {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:369
			_go_fuzz_dep_.CoverTab[94653]++
															s.norm[i] = 1
															distributed++
															total -= cnt
															continue
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:373
			// _ = "end of CoverTab[94653]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:374
			_go_fuzz_dep_.CoverTab[94654]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:374
			// _ = "end of CoverTab[94654]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:374
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:374
		// _ = "end of CoverTab[94647]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:374
		_go_fuzz_dep_.CoverTab[94648]++
														s.norm[i] = notYetAssigned
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:375
		// _ = "end of CoverTab[94648]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:376
	// _ = "end of CoverTab[94639]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:376
	_go_fuzz_dep_.CoverTab[94640]++
													toDistribute := (1 << tableLog) - distributed

													if (total / toDistribute) > lowOne {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:379
		_go_fuzz_dep_.CoverTab[94655]++

														lowOne = (total * 3) / (toDistribute * 2)
														for i, cnt := range s.count[:s.symbolLen] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:382
			_go_fuzz_dep_.CoverTab[94657]++
															if (s.norm[i] == notYetAssigned) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:383
				_go_fuzz_dep_.CoverTab[94658]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:383
				return (cnt <= lowOne)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:383
				// _ = "end of CoverTab[94658]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:383
			}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:383
				_go_fuzz_dep_.CoverTab[94659]++
																s.norm[i] = 1
																distributed++
																total -= cnt
																continue
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:387
				// _ = "end of CoverTab[94659]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:388
				_go_fuzz_dep_.CoverTab[94660]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:388
				// _ = "end of CoverTab[94660]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:388
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:388
			// _ = "end of CoverTab[94657]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:389
		// _ = "end of CoverTab[94655]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:389
		_go_fuzz_dep_.CoverTab[94656]++
														toDistribute = (1 << tableLog) - distributed
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:390
		// _ = "end of CoverTab[94656]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:391
		_go_fuzz_dep_.CoverTab[94661]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:391
		// _ = "end of CoverTab[94661]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:391
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:391
	// _ = "end of CoverTab[94640]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:391
	_go_fuzz_dep_.CoverTab[94641]++
													if distributed == uint32(s.symbolLen)+1 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:392
		_go_fuzz_dep_.CoverTab[94662]++
		// all values are pretty poor;
		//   probably incompressible data (should have already been detected);
		//   find max, then give all remaining points to max
		var maxV int
		var maxC uint32
		for i, cnt := range s.count[:s.symbolLen] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:398
			_go_fuzz_dep_.CoverTab[94664]++
															if cnt > maxC {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:399
				_go_fuzz_dep_.CoverTab[94665]++
																maxV = i
																maxC = cnt
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:401
				// _ = "end of CoverTab[94665]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:402
				_go_fuzz_dep_.CoverTab[94666]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:402
				// _ = "end of CoverTab[94666]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:402
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:402
			// _ = "end of CoverTab[94664]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:403
		// _ = "end of CoverTab[94662]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:403
		_go_fuzz_dep_.CoverTab[94663]++
														s.norm[maxV] += int16(toDistribute)
														return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:405
		// _ = "end of CoverTab[94663]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:406
		_go_fuzz_dep_.CoverTab[94667]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:406
		// _ = "end of CoverTab[94667]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:406
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:406
	// _ = "end of CoverTab[94641]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:406
	_go_fuzz_dep_.CoverTab[94642]++

													if total == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:408
		_go_fuzz_dep_.CoverTab[94668]++

														for i := uint32(0); toDistribute > 0; i = (i + 1) % (uint32(s.symbolLen)) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:410
			_go_fuzz_dep_.CoverTab[94670]++
															if s.norm[i] > 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:411
				_go_fuzz_dep_.CoverTab[94671]++
																toDistribute--
																s.norm[i]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:413
				// _ = "end of CoverTab[94671]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:414
				_go_fuzz_dep_.CoverTab[94672]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:414
				// _ = "end of CoverTab[94672]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:414
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:414
			// _ = "end of CoverTab[94670]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:415
		// _ = "end of CoverTab[94668]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:415
		_go_fuzz_dep_.CoverTab[94669]++
														return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:416
		// _ = "end of CoverTab[94669]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:417
		_go_fuzz_dep_.CoverTab[94673]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:417
		// _ = "end of CoverTab[94673]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:417
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:417
	// _ = "end of CoverTab[94642]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:417
	_go_fuzz_dep_.CoverTab[94643]++

													var (
		vStepLog	= 62 - uint64(tableLog)
		mid		= uint64((1 << (vStepLog - 1)) - 1)
		rStep		= (((1 << vStepLog) * uint64(toDistribute)) + mid) / uint64(total)	// scale on remaining
		tmpTotal	= mid
	)
	for i, cnt := range s.count[:s.symbolLen] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:425
		_go_fuzz_dep_.CoverTab[94674]++
														if s.norm[i] == notYetAssigned {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:426
			_go_fuzz_dep_.CoverTab[94675]++
															var (
				end	= tmpTotal + uint64(cnt)*rStep
				sStart	= uint32(tmpTotal >> vStepLog)
				sEnd	= uint32(end >> vStepLog)
				weight	= sEnd - sStart
			)
			if weight < 1 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:433
				_go_fuzz_dep_.CoverTab[94677]++
																return errors.New("weight < 1")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:434
				// _ = "end of CoverTab[94677]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:435
				_go_fuzz_dep_.CoverTab[94678]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:435
				// _ = "end of CoverTab[94678]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:435
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:435
			// _ = "end of CoverTab[94675]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:435
			_go_fuzz_dep_.CoverTab[94676]++
															s.norm[i] = int16(weight)
															tmpTotal = end
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:437
			// _ = "end of CoverTab[94676]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:438
			_go_fuzz_dep_.CoverTab[94679]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:438
			// _ = "end of CoverTab[94679]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:438
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:438
		// _ = "end of CoverTab[94674]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:439
	// _ = "end of CoverTab[94643]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:439
	_go_fuzz_dep_.CoverTab[94644]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:440
	// _ = "end of CoverTab[94644]"
}

// optimalTableLog calculates and sets the optimal tableLog in s.actualTableLog
func (s *fseEncoder) optimalTableLog(length int) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:444
	_go_fuzz_dep_.CoverTab[94680]++
													tableLog := uint8(maxEncTableLog)
													minBitsSrc := highBit(uint32(length)) + 1
													minBitsSymbols := highBit(uint32(s.symbolLen-1)) + 2
													minBits := uint8(minBitsSymbols)
													if minBitsSrc < minBitsSymbols {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:449
		_go_fuzz_dep_.CoverTab[94686]++
														minBits = uint8(minBitsSrc)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:450
		// _ = "end of CoverTab[94686]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:451
		_go_fuzz_dep_.CoverTab[94687]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:451
		// _ = "end of CoverTab[94687]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:451
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:451
	// _ = "end of CoverTab[94680]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:451
	_go_fuzz_dep_.CoverTab[94681]++

													maxBitsSrc := uint8(highBit(uint32(length-1))) - 2
													if maxBitsSrc < tableLog {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:454
		_go_fuzz_dep_.CoverTab[94688]++

														tableLog = maxBitsSrc
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:456
		// _ = "end of CoverTab[94688]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:457
		_go_fuzz_dep_.CoverTab[94689]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:457
		// _ = "end of CoverTab[94689]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:457
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:457
	// _ = "end of CoverTab[94681]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:457
	_go_fuzz_dep_.CoverTab[94682]++
													if minBits > tableLog {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:458
		_go_fuzz_dep_.CoverTab[94690]++
														tableLog = minBits
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:459
		// _ = "end of CoverTab[94690]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:460
		_go_fuzz_dep_.CoverTab[94691]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:460
		// _ = "end of CoverTab[94691]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:460
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:460
	// _ = "end of CoverTab[94682]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:460
	_go_fuzz_dep_.CoverTab[94683]++

													if tableLog < minEncTablelog {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:462
		_go_fuzz_dep_.CoverTab[94692]++
														tableLog = minEncTablelog
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:463
		// _ = "end of CoverTab[94692]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:464
		_go_fuzz_dep_.CoverTab[94693]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:464
		// _ = "end of CoverTab[94693]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:464
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:464
	// _ = "end of CoverTab[94683]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:464
	_go_fuzz_dep_.CoverTab[94684]++
													if tableLog > maxEncTableLog {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:465
		_go_fuzz_dep_.CoverTab[94694]++
														tableLog = maxEncTableLog
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:466
		// _ = "end of CoverTab[94694]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:467
		_go_fuzz_dep_.CoverTab[94695]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:467
		// _ = "end of CoverTab[94695]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:467
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:467
	// _ = "end of CoverTab[94684]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:467
	_go_fuzz_dep_.CoverTab[94685]++
													s.actualTableLog = tableLog
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:468
	// _ = "end of CoverTab[94685]"
}

// validateNorm validates the normalized histogram table.
func (s *fseEncoder) validateNorm() (err error) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:472
	_go_fuzz_dep_.CoverTab[94696]++
													var total int
													for _, v := range s.norm[:s.symbolLen] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:474
		_go_fuzz_dep_.CoverTab[94701]++
														if v >= 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:475
			_go_fuzz_dep_.CoverTab[94702]++
															total += int(v)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:476
			// _ = "end of CoverTab[94702]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:477
			_go_fuzz_dep_.CoverTab[94703]++
															total -= int(v)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:478
			// _ = "end of CoverTab[94703]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:479
		// _ = "end of CoverTab[94701]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:480
	// _ = "end of CoverTab[94696]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:480
	_go_fuzz_dep_.CoverTab[94697]++
													defer func() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:481
		_go_fuzz_dep_.CoverTab[94704]++
														if err == nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:482
			_go_fuzz_dep_.CoverTab[94706]++
															return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:483
			// _ = "end of CoverTab[94706]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:484
			_go_fuzz_dep_.CoverTab[94707]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:484
			// _ = "end of CoverTab[94707]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:484
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:484
		// _ = "end of CoverTab[94704]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:484
		_go_fuzz_dep_.CoverTab[94705]++
														fmt.Printf("selected TableLog: %d, Symbol length: %d\n", s.actualTableLog, s.symbolLen)
														for i, v := range s.norm[:s.symbolLen] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:486
			_go_fuzz_dep_.CoverTab[94708]++
															fmt.Printf("%3d: %5d -> %4d \n", i, s.count[i], v)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:487
			// _ = "end of CoverTab[94708]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:488
		// _ = "end of CoverTab[94705]"
	}()
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:489
	// _ = "end of CoverTab[94697]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:489
	_go_fuzz_dep_.CoverTab[94698]++
													if total != (1 << s.actualTableLog) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:490
		_go_fuzz_dep_.CoverTab[94709]++
														return fmt.Errorf("warning: Total == %d != %d", total, 1<<s.actualTableLog)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:491
		// _ = "end of CoverTab[94709]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:492
		_go_fuzz_dep_.CoverTab[94710]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:492
		// _ = "end of CoverTab[94710]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:492
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:492
	// _ = "end of CoverTab[94698]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:492
	_go_fuzz_dep_.CoverTab[94699]++
													for i, v := range s.count[s.symbolLen:] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:493
		_go_fuzz_dep_.CoverTab[94711]++
														if v != 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:494
			_go_fuzz_dep_.CoverTab[94712]++
															return fmt.Errorf("warning: Found symbol out of range, %d after cut", i)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:495
			// _ = "end of CoverTab[94712]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:496
			_go_fuzz_dep_.CoverTab[94713]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:496
			// _ = "end of CoverTab[94713]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:496
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:496
		// _ = "end of CoverTab[94711]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:497
	// _ = "end of CoverTab[94699]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:497
	_go_fuzz_dep_.CoverTab[94700]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:498
	// _ = "end of CoverTab[94700]"
}

// writeCount will write the normalized histogram count to header.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:501
// This is read back by readNCount.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:503
func (s *fseEncoder) writeCount(out []byte) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:503
	_go_fuzz_dep_.CoverTab[94714]++
													if s.useRLE {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:504
		_go_fuzz_dep_.CoverTab[94721]++
														return append(out, s.rleVal), nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:505
		// _ = "end of CoverTab[94721]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:506
		_go_fuzz_dep_.CoverTab[94722]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:506
		// _ = "end of CoverTab[94722]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:506
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:506
	// _ = "end of CoverTab[94714]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:506
	_go_fuzz_dep_.CoverTab[94715]++
													if s.preDefined || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:507
		_go_fuzz_dep_.CoverTab[94723]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:507
		return s.reUsed
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:507
		// _ = "end of CoverTab[94723]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:507
	}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:507
		_go_fuzz_dep_.CoverTab[94724]++

														return out, nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:509
		// _ = "end of CoverTab[94724]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:510
		_go_fuzz_dep_.CoverTab[94725]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:510
		// _ = "end of CoverTab[94725]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:510
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:510
	// _ = "end of CoverTab[94715]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:510
	_go_fuzz_dep_.CoverTab[94716]++

													var (
		tableLog	= s.actualTableLog
		tableSize	= 1 << tableLog
		previous0	bool
		charnum		uint16

		// maximum header size plus 2 extra bytes for final output if bitCount == 0.
		maxHeaderSize	= ((int(s.symbolLen) * int(tableLog)) >> 3) + 3 + 2

		// Write Table Size
		bitStream	= uint32(tableLog - minEncTablelog)
		bitCount	= uint(4)
		remaining	= int16(tableSize + 1)	/* +1 for extra accuracy */
		threshold	= int16(tableSize)
		nbBits		= uint(tableLog + 1)
		outP		= len(out)
	)
	if cap(out) < outP+maxHeaderSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:529
		_go_fuzz_dep_.CoverTab[94726]++
														out = append(out, make([]byte, maxHeaderSize*3)...)
														out = out[:len(out)-maxHeaderSize*3]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:531
		// _ = "end of CoverTab[94726]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:532
		_go_fuzz_dep_.CoverTab[94727]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:532
		// _ = "end of CoverTab[94727]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:532
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:532
	// _ = "end of CoverTab[94716]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:532
	_go_fuzz_dep_.CoverTab[94717]++
													out = out[:outP+maxHeaderSize]

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:536
	for remaining > 1 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:536
		_go_fuzz_dep_.CoverTab[94728]++
														if previous0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:537
			_go_fuzz_dep_.CoverTab[94735]++
															start := charnum
															for s.norm[charnum] == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:539
				_go_fuzz_dep_.CoverTab[94739]++
																charnum++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:540
				// _ = "end of CoverTab[94739]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:541
			// _ = "end of CoverTab[94735]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:541
			_go_fuzz_dep_.CoverTab[94736]++
															for charnum >= start+24 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:542
				_go_fuzz_dep_.CoverTab[94740]++
																start += 24
																bitStream += uint32(0xFFFF) << bitCount
																out[outP] = byte(bitStream)
																out[outP+1] = byte(bitStream >> 8)
																outP += 2
																bitStream >>= 16
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:548
				// _ = "end of CoverTab[94740]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:549
			// _ = "end of CoverTab[94736]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:549
			_go_fuzz_dep_.CoverTab[94737]++
															for charnum >= start+3 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:550
				_go_fuzz_dep_.CoverTab[94741]++
																start += 3
																bitStream += 3 << bitCount
																bitCount += 2
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:553
				// _ = "end of CoverTab[94741]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:554
			// _ = "end of CoverTab[94737]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:554
			_go_fuzz_dep_.CoverTab[94738]++
															bitStream += uint32(charnum-start) << bitCount
															bitCount += 2
															if bitCount > 16 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:557
				_go_fuzz_dep_.CoverTab[94742]++
																out[outP] = byte(bitStream)
																out[outP+1] = byte(bitStream >> 8)
																outP += 2
																bitStream >>= 16
																bitCount -= 16
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:562
				// _ = "end of CoverTab[94742]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:563
				_go_fuzz_dep_.CoverTab[94743]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:563
				// _ = "end of CoverTab[94743]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:563
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:563
			// _ = "end of CoverTab[94738]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:564
			_go_fuzz_dep_.CoverTab[94744]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:564
			// _ = "end of CoverTab[94744]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:564
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:564
		// _ = "end of CoverTab[94728]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:564
		_go_fuzz_dep_.CoverTab[94729]++

														count := s.norm[charnum]
														charnum++
														max := (2*threshold - 1) - remaining
														if count < 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:569
			_go_fuzz_dep_.CoverTab[94745]++
															remaining += count
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:570
			// _ = "end of CoverTab[94745]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:571
			_go_fuzz_dep_.CoverTab[94746]++
															remaining -= count
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:572
			// _ = "end of CoverTab[94746]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:573
		// _ = "end of CoverTab[94729]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:573
		_go_fuzz_dep_.CoverTab[94730]++
														count++
														if count >= threshold {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:575
			_go_fuzz_dep_.CoverTab[94747]++
															count += max
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:576
			// _ = "end of CoverTab[94747]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:577
			_go_fuzz_dep_.CoverTab[94748]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:577
			// _ = "end of CoverTab[94748]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:577
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:577
		// _ = "end of CoverTab[94730]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:577
		_go_fuzz_dep_.CoverTab[94731]++
														bitStream += uint32(count) << bitCount
														bitCount += nbBits
														if count < max {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:580
			_go_fuzz_dep_.CoverTab[94749]++
															bitCount--
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:581
			// _ = "end of CoverTab[94749]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:582
			_go_fuzz_dep_.CoverTab[94750]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:582
			// _ = "end of CoverTab[94750]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:582
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:582
		// _ = "end of CoverTab[94731]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:582
		_go_fuzz_dep_.CoverTab[94732]++

														previous0 = count == 1
														if remaining < 1 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:585
			_go_fuzz_dep_.CoverTab[94751]++
															return nil, errors.New("internal error: remaining < 1")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:586
			// _ = "end of CoverTab[94751]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:587
			_go_fuzz_dep_.CoverTab[94752]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:587
			// _ = "end of CoverTab[94752]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:587
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:587
		// _ = "end of CoverTab[94732]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:587
		_go_fuzz_dep_.CoverTab[94733]++
														for remaining < threshold {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:588
			_go_fuzz_dep_.CoverTab[94753]++
															nbBits--
															threshold >>= 1
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:590
			// _ = "end of CoverTab[94753]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:591
		// _ = "end of CoverTab[94733]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:591
		_go_fuzz_dep_.CoverTab[94734]++

														if bitCount > 16 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:593
			_go_fuzz_dep_.CoverTab[94754]++
															out[outP] = byte(bitStream)
															out[outP+1] = byte(bitStream >> 8)
															outP += 2
															bitStream >>= 16
															bitCount -= 16
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:598
			// _ = "end of CoverTab[94754]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:599
			_go_fuzz_dep_.CoverTab[94755]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:599
			// _ = "end of CoverTab[94755]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:599
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:599
		// _ = "end of CoverTab[94734]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:600
	// _ = "end of CoverTab[94717]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:600
	_go_fuzz_dep_.CoverTab[94718]++

													if outP+2 > len(out) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:602
		_go_fuzz_dep_.CoverTab[94756]++
														return nil, fmt.Errorf("internal error: %d > %d, maxheader: %d, sl: %d, tl: %d, normcount: %v", outP+2, len(out), maxHeaderSize, s.symbolLen, int(tableLog), s.norm[:s.symbolLen])
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:603
		// _ = "end of CoverTab[94756]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:604
		_go_fuzz_dep_.CoverTab[94757]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:604
		// _ = "end of CoverTab[94757]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:604
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:604
	// _ = "end of CoverTab[94718]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:604
	_go_fuzz_dep_.CoverTab[94719]++
													out[outP] = byte(bitStream)
													out[outP+1] = byte(bitStream >> 8)
													outP += int((bitCount + 7) / 8)

													if charnum > s.symbolLen {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:609
		_go_fuzz_dep_.CoverTab[94758]++
														return nil, errors.New("internal error: charnum > s.symbolLen")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:610
		// _ = "end of CoverTab[94758]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:611
		_go_fuzz_dep_.CoverTab[94759]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:611
		// _ = "end of CoverTab[94759]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:611
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:611
	// _ = "end of CoverTab[94719]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:611
	_go_fuzz_dep_.CoverTab[94720]++
													return out[:outP], nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:612
	// _ = "end of CoverTab[94720]"
}

// Approximate symbol cost, as fractional value, using fixed-point format (accuracyLog fractional bits)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:615
// note 1 : assume symbolValue is valid (<= maxSymbolValue)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:615
// note 2 : if freq[symbolValue]==0, @return a fake cost of tableLog+1 bits *
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:618
func (s *fseEncoder) bitCost(symbolValue uint8, accuracyLog uint32) uint32 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:618
	_go_fuzz_dep_.CoverTab[94760]++
													minNbBits := s.ct.symbolTT[symbolValue].deltaNbBits >> 16
													threshold := (minNbBits + 1) << 16
													if debugAsserts {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:621
		_go_fuzz_dep_.CoverTab[94763]++
														if !(s.actualTableLog < 16) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:622
			_go_fuzz_dep_.CoverTab[94765]++
															panic("!s.actualTableLog < 16")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:623
			// _ = "end of CoverTab[94765]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:624
			_go_fuzz_dep_.CoverTab[94766]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:624
			// _ = "end of CoverTab[94766]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:624
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:624
		// _ = "end of CoverTab[94763]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:624
		_go_fuzz_dep_.CoverTab[94764]++

														if !(uint8(accuracyLog) < 31-s.actualTableLog) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:626
			_go_fuzz_dep_.CoverTab[94767]++
															panic("!uint8(accuracyLog) < 31-s.actualTableLog")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:627
			// _ = "end of CoverTab[94767]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:628
			_go_fuzz_dep_.CoverTab[94768]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:628
			// _ = "end of CoverTab[94768]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:628
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:628
		// _ = "end of CoverTab[94764]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:629
		_go_fuzz_dep_.CoverTab[94769]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:629
		// _ = "end of CoverTab[94769]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:629
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:629
	// _ = "end of CoverTab[94760]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:629
	_go_fuzz_dep_.CoverTab[94761]++
													tableSize := uint32(1) << s.actualTableLog
													deltaFromThreshold := threshold - (s.ct.symbolTT[symbolValue].deltaNbBits + tableSize)

													normalizedDeltaFromThreshold := (deltaFromThreshold << accuracyLog) >> s.actualTableLog
													bitMultiplier := uint32(1) << accuracyLog
													if debugAsserts {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:635
		_go_fuzz_dep_.CoverTab[94770]++
														if s.ct.symbolTT[symbolValue].deltaNbBits+tableSize > threshold {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:636
			_go_fuzz_dep_.CoverTab[94772]++
															panic("s.ct.symbolTT[symbolValue].deltaNbBits+tableSize > threshold")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:637
			// _ = "end of CoverTab[94772]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:638
			_go_fuzz_dep_.CoverTab[94773]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:638
			// _ = "end of CoverTab[94773]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:638
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:638
		// _ = "end of CoverTab[94770]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:638
		_go_fuzz_dep_.CoverTab[94771]++
														if normalizedDeltaFromThreshold > bitMultiplier {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:639
			_go_fuzz_dep_.CoverTab[94774]++
															panic("normalizedDeltaFromThreshold > bitMultiplier")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:640
			// _ = "end of CoverTab[94774]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:641
			_go_fuzz_dep_.CoverTab[94775]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:641
			// _ = "end of CoverTab[94775]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:641
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:641
		// _ = "end of CoverTab[94771]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:642
		_go_fuzz_dep_.CoverTab[94776]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:642
		// _ = "end of CoverTab[94776]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:642
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:642
	// _ = "end of CoverTab[94761]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:642
	_go_fuzz_dep_.CoverTab[94762]++
													return (minNbBits+1)*bitMultiplier - normalizedDeltaFromThreshold
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:643
	// _ = "end of CoverTab[94762]"
}

// Returns the cost in bits of encoding the distribution in count using ctable.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:646
// Histogram should only be up to the last non-zero symbol.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:646
// Returns an -1 if ctable cannot represent all the symbols in count.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:649
func (s *fseEncoder) approxSize(hist []uint32) uint32 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:649
	_go_fuzz_dep_.CoverTab[94777]++
													if int(s.symbolLen) < len(hist) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:650
		_go_fuzz_dep_.CoverTab[94781]++

														return math.MaxUint32
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:652
		// _ = "end of CoverTab[94781]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:653
		_go_fuzz_dep_.CoverTab[94782]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:653
		// _ = "end of CoverTab[94782]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:653
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:653
	// _ = "end of CoverTab[94777]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:653
	_go_fuzz_dep_.CoverTab[94778]++
													if s.useRLE {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:654
		_go_fuzz_dep_.CoverTab[94783]++

														return math.MaxUint32
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:656
		// _ = "end of CoverTab[94783]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:657
		_go_fuzz_dep_.CoverTab[94784]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:657
		// _ = "end of CoverTab[94784]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:657
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:657
	// _ = "end of CoverTab[94778]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:657
	_go_fuzz_dep_.CoverTab[94779]++
													const kAccuracyLog = 8
													badCost := (uint32(s.actualTableLog) + 1) << kAccuracyLog
													var cost uint32
													for i, v := range hist {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:661
		_go_fuzz_dep_.CoverTab[94785]++
														if v == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:662
			_go_fuzz_dep_.CoverTab[94789]++
															continue
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:663
			// _ = "end of CoverTab[94789]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:664
			_go_fuzz_dep_.CoverTab[94790]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:664
			// _ = "end of CoverTab[94790]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:664
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:664
		// _ = "end of CoverTab[94785]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:664
		_go_fuzz_dep_.CoverTab[94786]++
														if s.norm[i] == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:665
			_go_fuzz_dep_.CoverTab[94791]++
															return math.MaxUint32
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:666
			// _ = "end of CoverTab[94791]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:667
			_go_fuzz_dep_.CoverTab[94792]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:667
			// _ = "end of CoverTab[94792]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:667
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:667
		// _ = "end of CoverTab[94786]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:667
		_go_fuzz_dep_.CoverTab[94787]++
														bitCost := s.bitCost(uint8(i), kAccuracyLog)
														if bitCost > badCost {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:669
			_go_fuzz_dep_.CoverTab[94793]++
															return math.MaxUint32
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:670
			// _ = "end of CoverTab[94793]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:671
			_go_fuzz_dep_.CoverTab[94794]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:671
			// _ = "end of CoverTab[94794]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:671
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:671
		// _ = "end of CoverTab[94787]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:671
		_go_fuzz_dep_.CoverTab[94788]++
														cost += v * bitCost
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:672
		// _ = "end of CoverTab[94788]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:673
	// _ = "end of CoverTab[94779]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:673
	_go_fuzz_dep_.CoverTab[94780]++
													return cost >> kAccuracyLog
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:674
	// _ = "end of CoverTab[94780]"
}

// maxHeaderSize returns the maximum header size in bits.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:677
// This is not exact size, but we want a penalty for new tables anyway.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:679
func (s *fseEncoder) maxHeaderSize() uint32 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:679
	_go_fuzz_dep_.CoverTab[94795]++
													if s.preDefined {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:680
		_go_fuzz_dep_.CoverTab[94798]++
														return 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:681
		// _ = "end of CoverTab[94798]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:682
		_go_fuzz_dep_.CoverTab[94799]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:682
		// _ = "end of CoverTab[94799]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:682
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:682
	// _ = "end of CoverTab[94795]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:682
	_go_fuzz_dep_.CoverTab[94796]++
													if s.useRLE {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:683
		_go_fuzz_dep_.CoverTab[94800]++
														return 8
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:684
		// _ = "end of CoverTab[94800]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:685
		_go_fuzz_dep_.CoverTab[94801]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:685
		// _ = "end of CoverTab[94801]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:685
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:685
	// _ = "end of CoverTab[94796]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:685
	_go_fuzz_dep_.CoverTab[94797]++
													return (((uint32(s.symbolLen) * uint32(s.actualTableLog)) >> 3) + 3) * 8
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:686
	// _ = "end of CoverTab[94797]"
}

// cState contains the compression state of a stream.
type cState struct {
	bw		*bitWriter
	stateTable	[]uint16
	state		uint16
}

// init will initialize the compression state to the first symbol of the stream.
func (c *cState) init(bw *bitWriter, ct *cTable, first symbolTransform) {
	c.bw = bw
	c.stateTable = ct.stateTable
	if len(c.stateTable) == 1 {

		c.stateTable[0] = uint16(0)
		c.state = 0
		return
	}
	nbBitsOut := (first.deltaNbBits + (1 << 15)) >> 16
	im := int32((nbBitsOut << 16) - first.deltaNbBits)
	lu := (im >> nbBitsOut) + int32(first.deltaFindState)
	c.state = c.stateTable[lu]
}

// encode the output symbol provided and write it to the bitstream.
func (c *cState) encode(symbolTT symbolTransform) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:713
	_go_fuzz_dep_.CoverTab[94802]++
													nbBitsOut := (uint32(c.state) + symbolTT.deltaNbBits) >> 16
													dstState := int32(c.state>>(nbBitsOut&15)) + int32(symbolTT.deltaFindState)
													c.bw.addBits16NC(c.state, uint8(nbBitsOut))
													c.state = c.stateTable[dstState]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:717
	// _ = "end of CoverTab[94802]"
}

// flush will write the tablelog to the output and flush the remaining full bytes.
func (c *cState) flush(tableLog uint8) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:721
	_go_fuzz_dep_.CoverTab[94803]++
													c.bw.flush32()
													c.bw.addBits16NC(c.state, tableLog)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:723
	// _ = "end of CoverTab[94803]"
}

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:724
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/fse_encoder.go:724
var _ = _go_fuzz_dep_.CoverTab
