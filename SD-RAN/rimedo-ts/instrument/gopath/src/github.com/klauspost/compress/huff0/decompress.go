//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1
package huff0

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1
)

import (
	"errors"
	"fmt"
	"io"

	"github.com/klauspost/compress/fse"
)

type dTable struct {
	single	[]dEntrySingle
	double	[]dEntryDouble
}

// single-symbols decoding
type dEntrySingle struct {
	entry uint16
}

// double-symbols decoding
type dEntryDouble struct {
	seq	[4]byte
	nBits	uint8
	len	uint8
}

// Uses special code for all tables that are < 8 bits.
const use8BitTables = true

// ReadTable will read a table from the input.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:31
// The size of the input may be larger than the table definition.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:31
// Any content remaining after the table definition will be returned.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:31
// If no Scratch is provided a new one is allocated.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:31
// The returned Scratch can be used for encoding or decoding input using this table.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:36
func ReadTable(in []byte, s *Scratch) (s2 *Scratch, remain []byte, err error) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:36
	_go_fuzz_dep_.CoverTab[89847]++
												s, err = s.prepare(in)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:38
		_go_fuzz_dep_.CoverTab[89857]++
													return s, nil, err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:39
		// _ = "end of CoverTab[89857]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:40
		_go_fuzz_dep_.CoverTab[89858]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:40
		// _ = "end of CoverTab[89858]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:40
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:40
	// _ = "end of CoverTab[89847]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:40
	_go_fuzz_dep_.CoverTab[89848]++
												if len(in) <= 1 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:41
		_go_fuzz_dep_.CoverTab[89859]++
													return s, nil, errors.New("input too small for table")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:42
		// _ = "end of CoverTab[89859]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:43
		_go_fuzz_dep_.CoverTab[89860]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:43
		// _ = "end of CoverTab[89860]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:43
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:43
	// _ = "end of CoverTab[89848]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:43
	_go_fuzz_dep_.CoverTab[89849]++
												iSize := in[0]
												in = in[1:]
												if iSize >= 128 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:46
		_go_fuzz_dep_.CoverTab[89861]++

													oSize := iSize - 127
													iSize = (oSize + 1) / 2
													if int(iSize) > len(in) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:50
			_go_fuzz_dep_.CoverTab[89864]++
														return s, nil, errors.New("input too small for table")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:51
			// _ = "end of CoverTab[89864]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:52
			_go_fuzz_dep_.CoverTab[89865]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:52
			// _ = "end of CoverTab[89865]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:52
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:52
		// _ = "end of CoverTab[89861]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:52
		_go_fuzz_dep_.CoverTab[89862]++
													for n := uint8(0); n < oSize; n += 2 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:53
			_go_fuzz_dep_.CoverTab[89866]++
														v := in[n/2]
														s.huffWeight[n] = v >> 4
														s.huffWeight[n+1] = v & 15
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:56
			// _ = "end of CoverTab[89866]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:57
		// _ = "end of CoverTab[89862]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:57
		_go_fuzz_dep_.CoverTab[89863]++
													s.symbolLen = uint16(oSize)
													in = in[iSize:]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:59
		// _ = "end of CoverTab[89863]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:60
		_go_fuzz_dep_.CoverTab[89867]++
													if len(in) < int(iSize) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:61
			_go_fuzz_dep_.CoverTab[89871]++
														return s, nil, fmt.Errorf("input too small for table, want %d bytes, have %d", iSize, len(in))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:62
			// _ = "end of CoverTab[89871]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:63
			_go_fuzz_dep_.CoverTab[89872]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:63
			// _ = "end of CoverTab[89872]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:63
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:63
		// _ = "end of CoverTab[89867]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:63
		_go_fuzz_dep_.CoverTab[89868]++

													s.fse.DecompressLimit = 255
													hw := s.huffWeight[:]
													s.fse.Out = hw
													b, err := fse.Decompress(in[:iSize], s.fse)
													s.fse.Out = nil
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:70
			_go_fuzz_dep_.CoverTab[89873]++
														return s, nil, err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:71
			// _ = "end of CoverTab[89873]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:72
			_go_fuzz_dep_.CoverTab[89874]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:72
			// _ = "end of CoverTab[89874]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:72
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:72
		// _ = "end of CoverTab[89868]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:72
		_go_fuzz_dep_.CoverTab[89869]++
													if len(b) > 255 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:73
			_go_fuzz_dep_.CoverTab[89875]++
														return s, nil, errors.New("corrupt input: output table too large")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:74
			// _ = "end of CoverTab[89875]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:75
			_go_fuzz_dep_.CoverTab[89876]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:75
			// _ = "end of CoverTab[89876]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:75
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:75
		// _ = "end of CoverTab[89869]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:75
		_go_fuzz_dep_.CoverTab[89870]++
													s.symbolLen = uint16(len(b))
													in = in[iSize:]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:77
		// _ = "end of CoverTab[89870]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:78
	// _ = "end of CoverTab[89849]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:78
	_go_fuzz_dep_.CoverTab[89850]++

	// collect weight stats
	var rankStats [16]uint32
	weightTotal := uint32(0)
	for _, v := range s.huffWeight[:s.symbolLen] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:83
		_go_fuzz_dep_.CoverTab[89877]++
													if v > tableLogMax {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:84
			_go_fuzz_dep_.CoverTab[89879]++
														return s, nil, errors.New("corrupt input: weight too large")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:85
			// _ = "end of CoverTab[89879]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:86
			_go_fuzz_dep_.CoverTab[89880]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:86
			// _ = "end of CoverTab[89880]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:86
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:86
		// _ = "end of CoverTab[89877]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:86
		_go_fuzz_dep_.CoverTab[89878]++
													v2 := v & 15
													rankStats[v2]++

													weightTotal += (1 << v2) >> 1
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:90
		// _ = "end of CoverTab[89878]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:91
	// _ = "end of CoverTab[89850]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:91
	_go_fuzz_dep_.CoverTab[89851]++
												if weightTotal == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:92
		_go_fuzz_dep_.CoverTab[89881]++
													return s, nil, errors.New("corrupt input: weights zero")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:93
		// _ = "end of CoverTab[89881]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:94
		_go_fuzz_dep_.CoverTab[89882]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:94
		// _ = "end of CoverTab[89882]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:94
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:94
	// _ = "end of CoverTab[89851]"

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:97
	{
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:97
		_go_fuzz_dep_.CoverTab[89883]++
													tableLog := highBit32(weightTotal) + 1
													if tableLog > tableLogMax {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:99
				_go_fuzz_dep_.CoverTab[89885]++
															return s, nil, errors.New("corrupt input: tableLog too big")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:100
			// _ = "end of CoverTab[89885]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:101
			_go_fuzz_dep_.CoverTab[89886]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:101
			// _ = "end of CoverTab[89886]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:101
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:101
		// _ = "end of CoverTab[89883]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:101
		_go_fuzz_dep_.CoverTab[89884]++
														s.actualTableLog = uint8(tableLog)

														{
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:104
			_go_fuzz_dep_.CoverTab[89887]++
															total := uint32(1) << tableLog
															rest := total - weightTotal
															verif := uint32(1) << highBit32(rest)
															lastWeight := highBit32(rest) + 1
															if verif != rest {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:109
				_go_fuzz_dep_.CoverTab[89889]++

																return s, nil, errors.New("corrupt input: last value not power of two")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:111
				// _ = "end of CoverTab[89889]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:112
				_go_fuzz_dep_.CoverTab[89890]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:112
				// _ = "end of CoverTab[89890]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:112
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:112
			// _ = "end of CoverTab[89887]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:112
			_go_fuzz_dep_.CoverTab[89888]++
															s.huffWeight[s.symbolLen] = uint8(lastWeight)
															s.symbolLen++
															rankStats[lastWeight]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:115
			// _ = "end of CoverTab[89888]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:116
		// _ = "end of CoverTab[89884]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:117
	_go_fuzz_dep_.CoverTab[89852]++

													if (rankStats[1] < 2) || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:119
		_go_fuzz_dep_.CoverTab[89891]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:119
		return (rankStats[1]&1 != 0)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:119
		// _ = "end of CoverTab[89891]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:119
	}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:119
		_go_fuzz_dep_.CoverTab[89892]++

														return s, nil, errors.New("corrupt input: min elt size, even check failed ")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:121
		// _ = "end of CoverTab[89892]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:122
		_go_fuzz_dep_.CoverTab[89893]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:122
		// _ = "end of CoverTab[89893]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:122
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:122
	// _ = "end of CoverTab[89852]"

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:127
	{
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:127
		_go_fuzz_dep_.CoverTab[89894]++
														var nextRankStart uint32
														for n := uint8(1); n < s.actualTableLog+1; n++ {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:129
			_go_fuzz_dep_.CoverTab[89895]++
															current := nextRankStart
															nextRankStart += rankStats[n] << (n - 1)
															rankStats[n] = current
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:132
			// _ = "end of CoverTab[89895]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:133
		// _ = "end of CoverTab[89894]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:134
	_go_fuzz_dep_.CoverTab[89853]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:137
	tSize := 1 << tableLogMax
	if len(s.dt.single) != tSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:138
		_go_fuzz_dep_.CoverTab[89896]++
														s.dt.single = make([]dEntrySingle, tSize)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:139
		// _ = "end of CoverTab[89896]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:140
		_go_fuzz_dep_.CoverTab[89897]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:140
		// _ = "end of CoverTab[89897]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:140
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:140
	// _ = "end of CoverTab[89853]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:140
	_go_fuzz_dep_.CoverTab[89854]++
													cTable := s.prevTable
													if cap(cTable) < maxSymbolValue+1 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:142
		_go_fuzz_dep_.CoverTab[89898]++
														cTable = make([]cTableEntry, 0, maxSymbolValue+1)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:143
		// _ = "end of CoverTab[89898]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:144
		_go_fuzz_dep_.CoverTab[89899]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:144
		// _ = "end of CoverTab[89899]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:144
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:144
	// _ = "end of CoverTab[89854]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:144
	_go_fuzz_dep_.CoverTab[89855]++
													cTable = cTable[:maxSymbolValue+1]
													s.prevTable = cTable[:s.symbolLen]
													s.prevTableLog = s.actualTableLog

													for n, w := range s.huffWeight[:s.symbolLen] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:149
		_go_fuzz_dep_.CoverTab[89900]++
														if w == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:150
			_go_fuzz_dep_.CoverTab[89903]++
															cTable[n] = cTableEntry{
				val:	0,
				nBits:	0,
			}
															continue
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:155
			// _ = "end of CoverTab[89903]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:156
			_go_fuzz_dep_.CoverTab[89904]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:156
			// _ = "end of CoverTab[89904]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:156
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:156
		// _ = "end of CoverTab[89900]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:156
		_go_fuzz_dep_.CoverTab[89901]++
														length := (uint32(1) << w) >> 1
														d := dEntrySingle{
			entry: uint16(s.actualTableLog+1-w) | (uint16(n) << 8),
		}

		rank := &rankStats[w]
		cTable[n] = cTableEntry{
			val:	uint16(*rank >> (w - 1)),
			nBits:	uint8(d.entry),
		}

		single := s.dt.single[*rank : *rank+length]
		for i := range single {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:169
			_go_fuzz_dep_.CoverTab[89905]++
															single[i] = d
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:170
			// _ = "end of CoverTab[89905]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:171
		// _ = "end of CoverTab[89901]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:171
		_go_fuzz_dep_.CoverTab[89902]++
														*rank += length
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:172
		// _ = "end of CoverTab[89902]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:173
	// _ = "end of CoverTab[89855]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:173
	_go_fuzz_dep_.CoverTab[89856]++

													return s, in, nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:175
	// _ = "end of CoverTab[89856]"
}

// Decompress1X will decompress a 1X encoded stream.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:178
// The length of the supplied input must match the end of a block exactly.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:178
// Before this is called, the table must be initialized with ReadTable unless
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:178
// the encoder re-used the table.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:178
// deprecated: Use the stateless Decoder() to get a concurrent version.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:183
func (s *Scratch) Decompress1X(in []byte) (out []byte, err error) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:183
	_go_fuzz_dep_.CoverTab[89906]++
													if cap(s.Out) < s.MaxDecodedSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:184
		_go_fuzz_dep_.CoverTab[89908]++
														s.Out = make([]byte, s.MaxDecodedSize)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:185
		// _ = "end of CoverTab[89908]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:186
		_go_fuzz_dep_.CoverTab[89909]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:186
		// _ = "end of CoverTab[89909]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:186
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:186
	// _ = "end of CoverTab[89906]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:186
	_go_fuzz_dep_.CoverTab[89907]++
													s.Out = s.Out[:0:s.MaxDecodedSize]
													s.Out, err = s.Decoder().Decompress1X(s.Out, in)
													return s.Out, err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:189
	// _ = "end of CoverTab[89907]"
}

// Decompress4X will decompress a 4X encoded stream.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:192
// Before this is called, the table must be initialized with ReadTable unless
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:192
// the encoder re-used the table.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:192
// The length of the supplied input must match the end of a block exactly.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:192
// The destination size of the uncompressed data must be known and provided.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:192
// deprecated: Use the stateless Decoder() to get a concurrent version.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:198
func (s *Scratch) Decompress4X(in []byte, dstSize int) (out []byte, err error) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:198
	_go_fuzz_dep_.CoverTab[89910]++
													if dstSize > s.MaxDecodedSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:199
		_go_fuzz_dep_.CoverTab[89913]++
														return nil, ErrMaxDecodedSizeExceeded
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:200
		// _ = "end of CoverTab[89913]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:201
		_go_fuzz_dep_.CoverTab[89914]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:201
		// _ = "end of CoverTab[89914]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:201
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:201
	// _ = "end of CoverTab[89910]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:201
	_go_fuzz_dep_.CoverTab[89911]++
													if cap(s.Out) < dstSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:202
		_go_fuzz_dep_.CoverTab[89915]++
														s.Out = make([]byte, s.MaxDecodedSize)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:203
		// _ = "end of CoverTab[89915]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:204
		_go_fuzz_dep_.CoverTab[89916]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:204
		// _ = "end of CoverTab[89916]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:204
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:204
	// _ = "end of CoverTab[89911]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:204
	_go_fuzz_dep_.CoverTab[89912]++
													s.Out = s.Out[:0:dstSize]
													s.Out, err = s.Decoder().Decompress4X(s.Out, in)
													return s.Out, err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:207
	// _ = "end of CoverTab[89912]"
}

// Decoder will return a stateless decoder that can be used by multiple
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:210
// decompressors concurrently.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:210
// Before this is called, the table must be initialized with ReadTable.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:210
// The Decoder is still linked to the scratch buffer so that cannot be reused.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:210
// However, it is safe to discard the scratch.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:215
func (s *Scratch) Decoder() *Decoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:215
	_go_fuzz_dep_.CoverTab[89917]++
													return &Decoder{
		dt:		s.dt,
		actualTableLog:	s.actualTableLog,
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:219
	// _ = "end of CoverTab[89917]"
}

// Decoder provides stateless decoding.
type Decoder struct {
	dt		dTable
	actualTableLog	uint8
}

// Decompress1X will decompress a 1X encoded stream.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:228
// The cap of the output buffer will be the maximum decompressed size.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:228
// The length of the supplied input must match the end of a block exactly.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:231
func (d *Decoder) Decompress1X(dst, src []byte) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:231
	_go_fuzz_dep_.CoverTab[89918]++
													if len(d.dt.single) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:232
		_go_fuzz_dep_.CoverTab[89925]++
														return nil, errors.New("no table loaded")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:233
		// _ = "end of CoverTab[89925]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:234
		_go_fuzz_dep_.CoverTab[89926]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:234
		// _ = "end of CoverTab[89926]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:234
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:234
	// _ = "end of CoverTab[89918]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:234
	_go_fuzz_dep_.CoverTab[89919]++
													if use8BitTables && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:235
		_go_fuzz_dep_.CoverTab[89927]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:235
		return d.actualTableLog <= 8
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:235
		// _ = "end of CoverTab[89927]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:235
	}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:235
		_go_fuzz_dep_.CoverTab[89928]++
														return d.decompress1X8Bit(dst, src)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:236
		// _ = "end of CoverTab[89928]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:237
		_go_fuzz_dep_.CoverTab[89929]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:237
		// _ = "end of CoverTab[89929]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:237
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:237
	// _ = "end of CoverTab[89919]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:237
	_go_fuzz_dep_.CoverTab[89920]++
													var br bitReaderShifted
													err := br.init(src)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:240
		_go_fuzz_dep_.CoverTab[89930]++
														return dst, err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:241
		// _ = "end of CoverTab[89930]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:242
		_go_fuzz_dep_.CoverTab[89931]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:242
		// _ = "end of CoverTab[89931]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:242
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:242
	// _ = "end of CoverTab[89920]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:242
	_go_fuzz_dep_.CoverTab[89921]++
													maxDecodedSize := cap(dst)
													dst = dst[:0]

	// Avoid bounds check by always having full sized table.
	const tlSize = 1 << tableLogMax
	const tlMask = tlSize - 1
	dt := d.dt.single[:tlSize]

	// Use temp table to avoid bound checks/append penalty.
	var buf [256]byte
	var off uint8

	for br.off >= 8 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:255
		_go_fuzz_dep_.CoverTab[89932]++
														br.fillFast()
														v := dt[br.peekBitsFast(d.actualTableLog)&tlMask]
														br.advance(uint8(v.entry))
														buf[off+0] = uint8(v.entry >> 8)

														v = dt[br.peekBitsFast(d.actualTableLog)&tlMask]
														br.advance(uint8(v.entry))
														buf[off+1] = uint8(v.entry >> 8)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:266
		br.fillFast()

		v = dt[br.peekBitsFast(d.actualTableLog)&tlMask]
		br.advance(uint8(v.entry))
		buf[off+2] = uint8(v.entry >> 8)

		v = dt[br.peekBitsFast(d.actualTableLog)&tlMask]
		br.advance(uint8(v.entry))
		buf[off+3] = uint8(v.entry >> 8)

		off += 4
		if off == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:277
			_go_fuzz_dep_.CoverTab[89933]++
															if len(dst)+256 > maxDecodedSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:278
				_go_fuzz_dep_.CoverTab[89935]++
																br.close()
																return nil, ErrMaxDecodedSizeExceeded
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:280
				// _ = "end of CoverTab[89935]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:281
				_go_fuzz_dep_.CoverTab[89936]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:281
				// _ = "end of CoverTab[89936]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:281
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:281
			// _ = "end of CoverTab[89933]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:281
			_go_fuzz_dep_.CoverTab[89934]++
															dst = append(dst, buf[:]...)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:282
			// _ = "end of CoverTab[89934]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:283
			_go_fuzz_dep_.CoverTab[89937]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:283
			// _ = "end of CoverTab[89937]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:283
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:283
		// _ = "end of CoverTab[89932]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:284
	// _ = "end of CoverTab[89921]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:284
	_go_fuzz_dep_.CoverTab[89922]++

													if len(dst)+int(off) > maxDecodedSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:286
		_go_fuzz_dep_.CoverTab[89938]++
														br.close()
														return nil, ErrMaxDecodedSizeExceeded
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:288
		// _ = "end of CoverTab[89938]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:289
		_go_fuzz_dep_.CoverTab[89939]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:289
		// _ = "end of CoverTab[89939]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:289
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:289
	// _ = "end of CoverTab[89922]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:289
	_go_fuzz_dep_.CoverTab[89923]++
													dst = append(dst, buf[:off]...)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:293
	bitsLeft := uint8(br.off)*8 + 64 - br.bitsRead
	for bitsLeft > 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:294
		_go_fuzz_dep_.CoverTab[89940]++
														br.fill()
														if false && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:296
			_go_fuzz_dep_.CoverTab[89943]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:296
			return br.bitsRead >= 32
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:296
			// _ = "end of CoverTab[89943]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:296
		}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:296
			_go_fuzz_dep_.CoverTab[89944]++
															if br.off >= 4 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:297
				_go_fuzz_dep_.CoverTab[89945]++
																v := br.in[br.off-4:]
																v = v[:4]
																low := (uint32(v[0])) | (uint32(v[1]) << 8) | (uint32(v[2]) << 16) | (uint32(v[3]) << 24)
																br.value = (br.value << 32) | uint64(low)
																br.bitsRead -= 32
																br.off -= 4
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:303
				// _ = "end of CoverTab[89945]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:304
				_go_fuzz_dep_.CoverTab[89946]++
																for br.off > 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:305
					_go_fuzz_dep_.CoverTab[89947]++
																	br.value = (br.value << 8) | uint64(br.in[br.off-1])
																	br.bitsRead -= 8
																	br.off--
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:308
					// _ = "end of CoverTab[89947]"
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:309
				// _ = "end of CoverTab[89946]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:310
			// _ = "end of CoverTab[89944]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:311
			_go_fuzz_dep_.CoverTab[89948]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:311
			// _ = "end of CoverTab[89948]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:311
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:311
		// _ = "end of CoverTab[89940]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:311
		_go_fuzz_dep_.CoverTab[89941]++
														if len(dst) >= maxDecodedSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:312
			_go_fuzz_dep_.CoverTab[89949]++
															br.close()
															return nil, ErrMaxDecodedSizeExceeded
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:314
			// _ = "end of CoverTab[89949]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:315
			_go_fuzz_dep_.CoverTab[89950]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:315
			// _ = "end of CoverTab[89950]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:315
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:315
		// _ = "end of CoverTab[89941]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:315
		_go_fuzz_dep_.CoverTab[89942]++
														v := d.dt.single[br.peekBitsFast(d.actualTableLog)&tlMask]
														nBits := uint8(v.entry)
														br.advance(nBits)
														bitsLeft -= nBits
														dst = append(dst, uint8(v.entry>>8))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:320
		// _ = "end of CoverTab[89942]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:321
	// _ = "end of CoverTab[89923]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:321
	_go_fuzz_dep_.CoverTab[89924]++
													return dst, br.close()
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:322
	// _ = "end of CoverTab[89924]"
}

// decompress1X8Bit will decompress a 1X encoded stream with tablelog <= 8.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:325
// The cap of the output buffer will be the maximum decompressed size.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:325
// The length of the supplied input must match the end of a block exactly.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:328
func (d *Decoder) decompress1X8Bit(dst, src []byte) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:328
	_go_fuzz_dep_.CoverTab[89951]++
													if d.actualTableLog == 8 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:329
		_go_fuzz_dep_.CoverTab[89957]++
														return d.decompress1X8BitExactly(dst, src)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:330
		// _ = "end of CoverTab[89957]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:331
		_go_fuzz_dep_.CoverTab[89958]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:331
		// _ = "end of CoverTab[89958]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:331
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:331
	// _ = "end of CoverTab[89951]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:331
	_go_fuzz_dep_.CoverTab[89952]++
													var br bitReaderBytes
													err := br.init(src)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:334
		_go_fuzz_dep_.CoverTab[89959]++
														return dst, err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:335
		// _ = "end of CoverTab[89959]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:336
		_go_fuzz_dep_.CoverTab[89960]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:336
		// _ = "end of CoverTab[89960]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:336
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:336
	// _ = "end of CoverTab[89952]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:336
	_go_fuzz_dep_.CoverTab[89953]++
													maxDecodedSize := cap(dst)
													dst = dst[:0]

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:341
	dt := d.dt.single[:256]

	// Use temp table to avoid bound checks/append penalty.
	var buf [256]byte
	var off uint8

	switch d.actualTableLog {
	case 8:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:348
		_go_fuzz_dep_.CoverTab[89961]++
														const shift = 8 - 8
														for br.off >= 4 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:350
			_go_fuzz_dep_.CoverTab[89970]++
															br.fillFast()
															v := dt[uint8(br.value>>(56+shift))]
															br.advance(uint8(v.entry))
															buf[off+0] = uint8(v.entry >> 8)

															v = dt[uint8(br.value>>(56+shift))]
															br.advance(uint8(v.entry))
															buf[off+1] = uint8(v.entry >> 8)

															v = dt[uint8(br.value>>(56+shift))]
															br.advance(uint8(v.entry))
															buf[off+2] = uint8(v.entry >> 8)

															v = dt[uint8(br.value>>(56+shift))]
															br.advance(uint8(v.entry))
															buf[off+3] = uint8(v.entry >> 8)

															off += 4
															if off == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:369
				_go_fuzz_dep_.CoverTab[89971]++
																if len(dst)+256 > maxDecodedSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:370
					_go_fuzz_dep_.CoverTab[89973]++
																	br.close()
																	return nil, ErrMaxDecodedSizeExceeded
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:372
					// _ = "end of CoverTab[89973]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:373
					_go_fuzz_dep_.CoverTab[89974]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:373
					// _ = "end of CoverTab[89974]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:373
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:373
				// _ = "end of CoverTab[89971]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:373
				_go_fuzz_dep_.CoverTab[89972]++
																dst = append(dst, buf[:]...)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:374
				// _ = "end of CoverTab[89972]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:375
				_go_fuzz_dep_.CoverTab[89975]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:375
				// _ = "end of CoverTab[89975]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:375
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:375
			// _ = "end of CoverTab[89970]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:376
		// _ = "end of CoverTab[89961]"
	case 7:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:377
		_go_fuzz_dep_.CoverTab[89962]++
														const shift = 8 - 7
														for br.off >= 4 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:379
			_go_fuzz_dep_.CoverTab[89976]++
															br.fillFast()
															v := dt[uint8(br.value>>(56+shift))]
															br.advance(uint8(v.entry))
															buf[off+0] = uint8(v.entry >> 8)

															v = dt[uint8(br.value>>(56+shift))]
															br.advance(uint8(v.entry))
															buf[off+1] = uint8(v.entry >> 8)

															v = dt[uint8(br.value>>(56+shift))]
															br.advance(uint8(v.entry))
															buf[off+2] = uint8(v.entry >> 8)

															v = dt[uint8(br.value>>(56+shift))]
															br.advance(uint8(v.entry))
															buf[off+3] = uint8(v.entry >> 8)

															off += 4
															if off == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:398
				_go_fuzz_dep_.CoverTab[89977]++
																if len(dst)+256 > maxDecodedSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:399
					_go_fuzz_dep_.CoverTab[89979]++
																	br.close()
																	return nil, ErrMaxDecodedSizeExceeded
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:401
					// _ = "end of CoverTab[89979]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:402
					_go_fuzz_dep_.CoverTab[89980]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:402
					// _ = "end of CoverTab[89980]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:402
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:402
				// _ = "end of CoverTab[89977]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:402
				_go_fuzz_dep_.CoverTab[89978]++
																dst = append(dst, buf[:]...)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:403
				// _ = "end of CoverTab[89978]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:404
				_go_fuzz_dep_.CoverTab[89981]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:404
				// _ = "end of CoverTab[89981]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:404
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:404
			// _ = "end of CoverTab[89976]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:405
		// _ = "end of CoverTab[89962]"
	case 6:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:406
		_go_fuzz_dep_.CoverTab[89963]++
														const shift = 8 - 6
														for br.off >= 4 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:408
			_go_fuzz_dep_.CoverTab[89982]++
															br.fillFast()
															v := dt[uint8(br.value>>(56+shift))]
															br.advance(uint8(v.entry))
															buf[off+0] = uint8(v.entry >> 8)

															v = dt[uint8(br.value>>(56+shift))]
															br.advance(uint8(v.entry))
															buf[off+1] = uint8(v.entry >> 8)

															v = dt[uint8(br.value>>(56+shift))]
															br.advance(uint8(v.entry))
															buf[off+2] = uint8(v.entry >> 8)

															v = dt[uint8(br.value>>(56+shift))]
															br.advance(uint8(v.entry))
															buf[off+3] = uint8(v.entry >> 8)

															off += 4
															if off == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:427
				_go_fuzz_dep_.CoverTab[89983]++
																if len(dst)+256 > maxDecodedSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:428
					_go_fuzz_dep_.CoverTab[89985]++
																	br.close()
																	return nil, ErrMaxDecodedSizeExceeded
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:430
					// _ = "end of CoverTab[89985]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:431
					_go_fuzz_dep_.CoverTab[89986]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:431
					// _ = "end of CoverTab[89986]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:431
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:431
				// _ = "end of CoverTab[89983]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:431
				_go_fuzz_dep_.CoverTab[89984]++
																dst = append(dst, buf[:]...)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:432
				// _ = "end of CoverTab[89984]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:433
				_go_fuzz_dep_.CoverTab[89987]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:433
				// _ = "end of CoverTab[89987]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:433
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:433
			// _ = "end of CoverTab[89982]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:434
		// _ = "end of CoverTab[89963]"
	case 5:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:435
		_go_fuzz_dep_.CoverTab[89964]++
														const shift = 8 - 5
														for br.off >= 4 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:437
			_go_fuzz_dep_.CoverTab[89988]++
															br.fillFast()
															v := dt[uint8(br.value>>(56+shift))]
															br.advance(uint8(v.entry))
															buf[off+0] = uint8(v.entry >> 8)

															v = dt[uint8(br.value>>(56+shift))]
															br.advance(uint8(v.entry))
															buf[off+1] = uint8(v.entry >> 8)

															v = dt[uint8(br.value>>(56+shift))]
															br.advance(uint8(v.entry))
															buf[off+2] = uint8(v.entry >> 8)

															v = dt[uint8(br.value>>(56+shift))]
															br.advance(uint8(v.entry))
															buf[off+3] = uint8(v.entry >> 8)

															off += 4
															if off == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:456
				_go_fuzz_dep_.CoverTab[89989]++
																if len(dst)+256 > maxDecodedSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:457
					_go_fuzz_dep_.CoverTab[89991]++
																	br.close()
																	return nil, ErrMaxDecodedSizeExceeded
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:459
					// _ = "end of CoverTab[89991]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:460
					_go_fuzz_dep_.CoverTab[89992]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:460
					// _ = "end of CoverTab[89992]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:460
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:460
				// _ = "end of CoverTab[89989]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:460
				_go_fuzz_dep_.CoverTab[89990]++
																dst = append(dst, buf[:]...)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:461
				// _ = "end of CoverTab[89990]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:462
				_go_fuzz_dep_.CoverTab[89993]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:462
				// _ = "end of CoverTab[89993]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:462
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:462
			// _ = "end of CoverTab[89988]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:463
		// _ = "end of CoverTab[89964]"
	case 4:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:464
		_go_fuzz_dep_.CoverTab[89965]++
														const shift = 8 - 4
														for br.off >= 4 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:466
			_go_fuzz_dep_.CoverTab[89994]++
															br.fillFast()
															v := dt[uint8(br.value>>(56+shift))]
															br.advance(uint8(v.entry))
															buf[off+0] = uint8(v.entry >> 8)

															v = dt[uint8(br.value>>(56+shift))]
															br.advance(uint8(v.entry))
															buf[off+1] = uint8(v.entry >> 8)

															v = dt[uint8(br.value>>(56+shift))]
															br.advance(uint8(v.entry))
															buf[off+2] = uint8(v.entry >> 8)

															v = dt[uint8(br.value>>(56+shift))]
															br.advance(uint8(v.entry))
															buf[off+3] = uint8(v.entry >> 8)

															off += 4
															if off == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:485
				_go_fuzz_dep_.CoverTab[89995]++
																if len(dst)+256 > maxDecodedSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:486
					_go_fuzz_dep_.CoverTab[89997]++
																	br.close()
																	return nil, ErrMaxDecodedSizeExceeded
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:488
					// _ = "end of CoverTab[89997]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:489
					_go_fuzz_dep_.CoverTab[89998]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:489
					// _ = "end of CoverTab[89998]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:489
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:489
				// _ = "end of CoverTab[89995]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:489
				_go_fuzz_dep_.CoverTab[89996]++
																dst = append(dst, buf[:]...)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:490
				// _ = "end of CoverTab[89996]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:491
				_go_fuzz_dep_.CoverTab[89999]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:491
				// _ = "end of CoverTab[89999]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:491
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:491
			// _ = "end of CoverTab[89994]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:492
		// _ = "end of CoverTab[89965]"
	case 3:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:493
		_go_fuzz_dep_.CoverTab[89966]++
														const shift = 8 - 3
														for br.off >= 4 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:495
			_go_fuzz_dep_.CoverTab[90000]++
															br.fillFast()
															v := dt[uint8(br.value>>(56+shift))]
															br.advance(uint8(v.entry))
															buf[off+0] = uint8(v.entry >> 8)

															v = dt[uint8(br.value>>(56+shift))]
															br.advance(uint8(v.entry))
															buf[off+1] = uint8(v.entry >> 8)

															v = dt[uint8(br.value>>(56+shift))]
															br.advance(uint8(v.entry))
															buf[off+2] = uint8(v.entry >> 8)

															v = dt[uint8(br.value>>(56+shift))]
															br.advance(uint8(v.entry))
															buf[off+3] = uint8(v.entry >> 8)

															off += 4
															if off == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:514
				_go_fuzz_dep_.CoverTab[90001]++
																if len(dst)+256 > maxDecodedSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:515
					_go_fuzz_dep_.CoverTab[90003]++
																	br.close()
																	return nil, ErrMaxDecodedSizeExceeded
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:517
					// _ = "end of CoverTab[90003]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:518
					_go_fuzz_dep_.CoverTab[90004]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:518
					// _ = "end of CoverTab[90004]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:518
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:518
				// _ = "end of CoverTab[90001]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:518
				_go_fuzz_dep_.CoverTab[90002]++
																dst = append(dst, buf[:]...)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:519
				// _ = "end of CoverTab[90002]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:520
				_go_fuzz_dep_.CoverTab[90005]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:520
				// _ = "end of CoverTab[90005]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:520
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:520
			// _ = "end of CoverTab[90000]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:521
		// _ = "end of CoverTab[89966]"
	case 2:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:522
		_go_fuzz_dep_.CoverTab[89967]++
														const shift = 8 - 2
														for br.off >= 4 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:524
			_go_fuzz_dep_.CoverTab[90006]++
															br.fillFast()
															v := dt[uint8(br.value>>(56+shift))]
															br.advance(uint8(v.entry))
															buf[off+0] = uint8(v.entry >> 8)

															v = dt[uint8(br.value>>(56+shift))]
															br.advance(uint8(v.entry))
															buf[off+1] = uint8(v.entry >> 8)

															v = dt[uint8(br.value>>(56+shift))]
															br.advance(uint8(v.entry))
															buf[off+2] = uint8(v.entry >> 8)

															v = dt[uint8(br.value>>(56+shift))]
															br.advance(uint8(v.entry))
															buf[off+3] = uint8(v.entry >> 8)

															off += 4
															if off == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:543
				_go_fuzz_dep_.CoverTab[90007]++
																if len(dst)+256 > maxDecodedSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:544
					_go_fuzz_dep_.CoverTab[90009]++
																	br.close()
																	return nil, ErrMaxDecodedSizeExceeded
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:546
					// _ = "end of CoverTab[90009]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:547
					_go_fuzz_dep_.CoverTab[90010]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:547
					// _ = "end of CoverTab[90010]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:547
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:547
				// _ = "end of CoverTab[90007]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:547
				_go_fuzz_dep_.CoverTab[90008]++
																dst = append(dst, buf[:]...)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:548
				// _ = "end of CoverTab[90008]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:549
				_go_fuzz_dep_.CoverTab[90011]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:549
				// _ = "end of CoverTab[90011]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:549
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:549
			// _ = "end of CoverTab[90006]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:550
		// _ = "end of CoverTab[89967]"
	case 1:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:551
		_go_fuzz_dep_.CoverTab[89968]++
														const shift = 8 - 1
														for br.off >= 4 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:553
			_go_fuzz_dep_.CoverTab[90012]++
															br.fillFast()
															v := dt[uint8(br.value>>(56+shift))]
															br.advance(uint8(v.entry))
															buf[off+0] = uint8(v.entry >> 8)

															v = dt[uint8(br.value>>(56+shift))]
															br.advance(uint8(v.entry))
															buf[off+1] = uint8(v.entry >> 8)

															v = dt[uint8(br.value>>(56+shift))]
															br.advance(uint8(v.entry))
															buf[off+2] = uint8(v.entry >> 8)

															v = dt[uint8(br.value>>(56+shift))]
															br.advance(uint8(v.entry))
															buf[off+3] = uint8(v.entry >> 8)

															off += 4
															if off == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:572
				_go_fuzz_dep_.CoverTab[90013]++
																if len(dst)+256 > maxDecodedSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:573
					_go_fuzz_dep_.CoverTab[90015]++
																	br.close()
																	return nil, ErrMaxDecodedSizeExceeded
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:575
					// _ = "end of CoverTab[90015]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:576
					_go_fuzz_dep_.CoverTab[90016]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:576
					// _ = "end of CoverTab[90016]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:576
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:576
				// _ = "end of CoverTab[90013]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:576
				_go_fuzz_dep_.CoverTab[90014]++
																dst = append(dst, buf[:]...)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:577
				// _ = "end of CoverTab[90014]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:578
				_go_fuzz_dep_.CoverTab[90017]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:578
				// _ = "end of CoverTab[90017]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:578
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:578
			// _ = "end of CoverTab[90012]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:579
		// _ = "end of CoverTab[89968]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:580
		_go_fuzz_dep_.CoverTab[89969]++
														return nil, fmt.Errorf("invalid tablelog: %d", d.actualTableLog)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:581
		// _ = "end of CoverTab[89969]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:582
	// _ = "end of CoverTab[89953]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:582
	_go_fuzz_dep_.CoverTab[89954]++

													if len(dst)+int(off) > maxDecodedSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:584
		_go_fuzz_dep_.CoverTab[90018]++
														br.close()
														return nil, ErrMaxDecodedSizeExceeded
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:586
		// _ = "end of CoverTab[90018]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:587
		_go_fuzz_dep_.CoverTab[90019]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:587
		// _ = "end of CoverTab[90019]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:587
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:587
	// _ = "end of CoverTab[89954]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:587
	_go_fuzz_dep_.CoverTab[89955]++
													dst = append(dst, buf[:off]...)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:591
	bitsLeft := int8(uint8(br.off)*8 + (64 - br.bitsRead))
	shift := (8 - d.actualTableLog) & 7

	for bitsLeft > 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:594
		_go_fuzz_dep_.CoverTab[90020]++
														if br.bitsRead >= 64-8 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:595
			_go_fuzz_dep_.CoverTab[90023]++
															for br.off > 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:596
				_go_fuzz_dep_.CoverTab[90024]++
																br.value |= uint64(br.in[br.off-1]) << (br.bitsRead - 8)
																br.bitsRead -= 8
																br.off--
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:599
				// _ = "end of CoverTab[90024]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:600
			// _ = "end of CoverTab[90023]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:601
			_go_fuzz_dep_.CoverTab[90025]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:601
			// _ = "end of CoverTab[90025]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:601
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:601
		// _ = "end of CoverTab[90020]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:601
		_go_fuzz_dep_.CoverTab[90021]++
														if len(dst) >= maxDecodedSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:602
			_go_fuzz_dep_.CoverTab[90026]++
															br.close()
															return nil, ErrMaxDecodedSizeExceeded
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:604
			// _ = "end of CoverTab[90026]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:605
			_go_fuzz_dep_.CoverTab[90027]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:605
			// _ = "end of CoverTab[90027]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:605
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:605
		// _ = "end of CoverTab[90021]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:605
		_go_fuzz_dep_.CoverTab[90022]++
														v := dt[br.peekByteFast()>>shift]
														nBits := uint8(v.entry)
														br.advance(nBits)
														bitsLeft -= int8(nBits)
														dst = append(dst, uint8(v.entry>>8))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:610
		// _ = "end of CoverTab[90022]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:611
	// _ = "end of CoverTab[89955]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:611
	_go_fuzz_dep_.CoverTab[89956]++
													return dst, br.close()
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:612
	// _ = "end of CoverTab[89956]"
}

// decompress1X8Bit will decompress a 1X encoded stream with tablelog <= 8.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:615
// The cap of the output buffer will be the maximum decompressed size.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:615
// The length of the supplied input must match the end of a block exactly.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:618
func (d *Decoder) decompress1X8BitExactly(dst, src []byte) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:618
	_go_fuzz_dep_.CoverTab[90028]++
													var br bitReaderBytes
													err := br.init(src)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:621
		_go_fuzz_dep_.CoverTab[90033]++
														return dst, err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:622
		// _ = "end of CoverTab[90033]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:623
		_go_fuzz_dep_.CoverTab[90034]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:623
		// _ = "end of CoverTab[90034]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:623
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:623
	// _ = "end of CoverTab[90028]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:623
	_go_fuzz_dep_.CoverTab[90029]++
													maxDecodedSize := cap(dst)
													dst = dst[:0]

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:628
	dt := d.dt.single[:256]

													// Use temp table to avoid bound checks/append penalty.
													var buf [256]byte
													var off uint8

													const shift = 56

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:637
	for br.off >= 4 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:637
		_go_fuzz_dep_.CoverTab[90035]++
														br.fillFast()
														v := dt[uint8(br.value>>shift)]
														br.advance(uint8(v.entry))
														buf[off+0] = uint8(v.entry >> 8)

														v = dt[uint8(br.value>>shift)]
														br.advance(uint8(v.entry))
														buf[off+1] = uint8(v.entry >> 8)

														v = dt[uint8(br.value>>shift)]
														br.advance(uint8(v.entry))
														buf[off+2] = uint8(v.entry >> 8)

														v = dt[uint8(br.value>>shift)]
														br.advance(uint8(v.entry))
														buf[off+3] = uint8(v.entry >> 8)

														off += 4
														if off == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:656
			_go_fuzz_dep_.CoverTab[90036]++
															if len(dst)+256 > maxDecodedSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:657
				_go_fuzz_dep_.CoverTab[90038]++
																br.close()
																return nil, ErrMaxDecodedSizeExceeded
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:659
				// _ = "end of CoverTab[90038]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:660
				_go_fuzz_dep_.CoverTab[90039]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:660
				// _ = "end of CoverTab[90039]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:660
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:660
			// _ = "end of CoverTab[90036]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:660
			_go_fuzz_dep_.CoverTab[90037]++
															dst = append(dst, buf[:]...)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:661
			// _ = "end of CoverTab[90037]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:662
			_go_fuzz_dep_.CoverTab[90040]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:662
			// _ = "end of CoverTab[90040]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:662
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:662
		// _ = "end of CoverTab[90035]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:663
	// _ = "end of CoverTab[90029]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:663
	_go_fuzz_dep_.CoverTab[90030]++

													if len(dst)+int(off) > maxDecodedSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:665
		_go_fuzz_dep_.CoverTab[90041]++
														br.close()
														return nil, ErrMaxDecodedSizeExceeded
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:667
		// _ = "end of CoverTab[90041]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:668
		_go_fuzz_dep_.CoverTab[90042]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:668
		// _ = "end of CoverTab[90042]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:668
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:668
	// _ = "end of CoverTab[90030]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:668
	_go_fuzz_dep_.CoverTab[90031]++
													dst = append(dst, buf[:off]...)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:672
	bitsLeft := int8(uint8(br.off)*8 + (64 - br.bitsRead))
	for bitsLeft > 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:673
		_go_fuzz_dep_.CoverTab[90043]++
														if br.bitsRead >= 64-8 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:674
			_go_fuzz_dep_.CoverTab[90046]++
															for br.off > 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:675
				_go_fuzz_dep_.CoverTab[90047]++
																br.value |= uint64(br.in[br.off-1]) << (br.bitsRead - 8)
																br.bitsRead -= 8
																br.off--
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:678
				// _ = "end of CoverTab[90047]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:679
			// _ = "end of CoverTab[90046]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:680
			_go_fuzz_dep_.CoverTab[90048]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:680
			// _ = "end of CoverTab[90048]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:680
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:680
		// _ = "end of CoverTab[90043]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:680
		_go_fuzz_dep_.CoverTab[90044]++
														if len(dst) >= maxDecodedSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:681
			_go_fuzz_dep_.CoverTab[90049]++
															br.close()
															return nil, ErrMaxDecodedSizeExceeded
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:683
			// _ = "end of CoverTab[90049]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:684
			_go_fuzz_dep_.CoverTab[90050]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:684
			// _ = "end of CoverTab[90050]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:684
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:684
		// _ = "end of CoverTab[90044]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:684
		_go_fuzz_dep_.CoverTab[90045]++
														v := dt[br.peekByteFast()]
														nBits := uint8(v.entry)
														br.advance(nBits)
														bitsLeft -= int8(nBits)
														dst = append(dst, uint8(v.entry>>8))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:689
		// _ = "end of CoverTab[90045]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:690
	// _ = "end of CoverTab[90031]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:690
	_go_fuzz_dep_.CoverTab[90032]++
													return dst, br.close()
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:691
	// _ = "end of CoverTab[90032]"
}

// Decompress4X will decompress a 4X encoded stream.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:694
// The length of the supplied input must match the end of a block exactly.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:694
// The *capacity* of the dst slice must match the destination size of
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:694
// the uncompressed data exactly.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:698
func (d *Decoder) Decompress4X(dst, src []byte) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:698
	_go_fuzz_dep_.CoverTab[90051]++
													if len(d.dt.single) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:699
		_go_fuzz_dep_.CoverTab[90061]++
														return nil, errors.New("no table loaded")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:700
		// _ = "end of CoverTab[90061]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:701
		_go_fuzz_dep_.CoverTab[90062]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:701
		// _ = "end of CoverTab[90062]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:701
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:701
	// _ = "end of CoverTab[90051]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:701
	_go_fuzz_dep_.CoverTab[90052]++
													if len(src) < 6+(4*1) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:702
		_go_fuzz_dep_.CoverTab[90063]++
														return nil, errors.New("input too small")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:703
		// _ = "end of CoverTab[90063]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:704
		_go_fuzz_dep_.CoverTab[90064]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:704
		// _ = "end of CoverTab[90064]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:704
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:704
	// _ = "end of CoverTab[90052]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:704
	_go_fuzz_dep_.CoverTab[90053]++
													if use8BitTables && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:705
		_go_fuzz_dep_.CoverTab[90065]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:705
		return d.actualTableLog <= 8
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:705
		// _ = "end of CoverTab[90065]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:705
	}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:705
		_go_fuzz_dep_.CoverTab[90066]++
														return d.decompress4X8bit(dst, src)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:706
		// _ = "end of CoverTab[90066]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:707
		_go_fuzz_dep_.CoverTab[90067]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:707
		// _ = "end of CoverTab[90067]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:707
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:707
	// _ = "end of CoverTab[90053]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:707
	_go_fuzz_dep_.CoverTab[90054]++

													var br [4]bitReaderShifted
													start := 6
													for i := 0; i < 3; i++ {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:711
		_go_fuzz_dep_.CoverTab[90068]++
														length := int(src[i*2]) | (int(src[i*2+1]) << 8)
														if start+length >= len(src) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:713
			_go_fuzz_dep_.CoverTab[90071]++
															return nil, errors.New("truncated input (or invalid offset)")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:714
			// _ = "end of CoverTab[90071]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:715
			_go_fuzz_dep_.CoverTab[90072]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:715
			// _ = "end of CoverTab[90072]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:715
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:715
		// _ = "end of CoverTab[90068]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:715
		_go_fuzz_dep_.CoverTab[90069]++
														err := br[i].init(src[start : start+length])
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:717
			_go_fuzz_dep_.CoverTab[90073]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:718
			// _ = "end of CoverTab[90073]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:719
			_go_fuzz_dep_.CoverTab[90074]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:719
			// _ = "end of CoverTab[90074]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:719
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:719
		// _ = "end of CoverTab[90069]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:719
		_go_fuzz_dep_.CoverTab[90070]++
														start += length
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:720
		// _ = "end of CoverTab[90070]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:721
	// _ = "end of CoverTab[90054]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:721
	_go_fuzz_dep_.CoverTab[90055]++
													err := br[3].init(src[start:])
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:723
		_go_fuzz_dep_.CoverTab[90075]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:724
		// _ = "end of CoverTab[90075]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:725
		_go_fuzz_dep_.CoverTab[90076]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:725
		// _ = "end of CoverTab[90076]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:725
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:725
	// _ = "end of CoverTab[90055]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:725
	_go_fuzz_dep_.CoverTab[90056]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:728
	dstSize := cap(dst)
	dst = dst[:dstSize]
	out := dst
	dstEvery := (dstSize + 3) / 4

	const tlSize = 1 << tableLogMax
	const tlMask = tlSize - 1
	single := d.dt.single[:tlSize]

	// Use temp table to avoid bound checks/append penalty.
	var buf [256]byte
	var off uint8
	var decoded int

	// Decode 2 values from each decoder/loop.
	const bufoff = 256 / 4
	for {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:744
		_go_fuzz_dep_.CoverTab[90077]++
														if br[0].off < 4 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:745
			_go_fuzz_dep_.CoverTab[90079]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:745
			return br[1].off < 4
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:745
			// _ = "end of CoverTab[90079]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:745
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:745
			_go_fuzz_dep_.CoverTab[90080]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:745
			return br[2].off < 4
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:745
			// _ = "end of CoverTab[90080]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:745
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:745
			_go_fuzz_dep_.CoverTab[90081]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:745
			return br[3].off < 4
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:745
			// _ = "end of CoverTab[90081]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:745
		}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:745
			_go_fuzz_dep_.CoverTab[90082]++
															break
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:746
			// _ = "end of CoverTab[90082]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:747
			_go_fuzz_dep_.CoverTab[90083]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:747
			// _ = "end of CoverTab[90083]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:747
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:747
		// _ = "end of CoverTab[90077]"

														{
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:749
			_go_fuzz_dep_.CoverTab[90084]++
															const stream = 0
															const stream2 = 1
															br[stream].fillFast()
															br[stream2].fillFast()

															val := br[stream].peekBitsFast(d.actualTableLog)
															val2 := br[stream2].peekBitsFast(d.actualTableLog)
															v := single[val&tlMask]
															v2 := single[val2&tlMask]
															br[stream].advance(uint8(v.entry))
															br[stream2].advance(uint8(v2.entry))
															buf[off+bufoff*stream] = uint8(v.entry >> 8)
															buf[off+bufoff*stream2] = uint8(v2.entry >> 8)

															val = br[stream].peekBitsFast(d.actualTableLog)
															val2 = br[stream2].peekBitsFast(d.actualTableLog)
															v = single[val&tlMask]
															v2 = single[val2&tlMask]
															br[stream].advance(uint8(v.entry))
															br[stream2].advance(uint8(v2.entry))
															buf[off+bufoff*stream+1] = uint8(v.entry >> 8)
															buf[off+bufoff*stream2+1] = uint8(v2.entry >> 8)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:771
			// _ = "end of CoverTab[90084]"
		}

		{
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:774
			_go_fuzz_dep_.CoverTab[90085]++
															const stream = 2
															const stream2 = 3
															br[stream].fillFast()
															br[stream2].fillFast()

															val := br[stream].peekBitsFast(d.actualTableLog)
															val2 := br[stream2].peekBitsFast(d.actualTableLog)
															v := single[val&tlMask]
															v2 := single[val2&tlMask]
															br[stream].advance(uint8(v.entry))
															br[stream2].advance(uint8(v2.entry))
															buf[off+bufoff*stream] = uint8(v.entry >> 8)
															buf[off+bufoff*stream2] = uint8(v2.entry >> 8)

															val = br[stream].peekBitsFast(d.actualTableLog)
															val2 = br[stream2].peekBitsFast(d.actualTableLog)
															v = single[val&tlMask]
															v2 = single[val2&tlMask]
															br[stream].advance(uint8(v.entry))
															br[stream2].advance(uint8(v2.entry))
															buf[off+bufoff*stream+1] = uint8(v.entry >> 8)
															buf[off+bufoff*stream2+1] = uint8(v2.entry >> 8)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:796
			// _ = "end of CoverTab[90085]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:797
		_go_fuzz_dep_.CoverTab[90078]++

														off += 2

														if off == bufoff {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:801
			_go_fuzz_dep_.CoverTab[90086]++
															if bufoff > dstEvery {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:802
				_go_fuzz_dep_.CoverTab[90088]++
																return nil, errors.New("corruption detected: stream overrun 1")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:803
				// _ = "end of CoverTab[90088]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:804
				_go_fuzz_dep_.CoverTab[90089]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:804
				// _ = "end of CoverTab[90089]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:804
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:804
			// _ = "end of CoverTab[90086]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:804
			_go_fuzz_dep_.CoverTab[90087]++
															copy(out, buf[:bufoff])
															copy(out[dstEvery:], buf[bufoff:bufoff*2])
															copy(out[dstEvery*2:], buf[bufoff*2:bufoff*3])
															copy(out[dstEvery*3:], buf[bufoff*3:bufoff*4])
															off = 0
															out = out[bufoff:]
															decoded += 256

															if len(out) < dstEvery*3 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:813
				_go_fuzz_dep_.CoverTab[90090]++
																return nil, errors.New("corruption detected: stream overrun 2")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:814
				// _ = "end of CoverTab[90090]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:815
				_go_fuzz_dep_.CoverTab[90091]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:815
				// _ = "end of CoverTab[90091]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:815
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:815
			// _ = "end of CoverTab[90087]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:816
			_go_fuzz_dep_.CoverTab[90092]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:816
			// _ = "end of CoverTab[90092]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:816
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:816
		// _ = "end of CoverTab[90078]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:817
	// _ = "end of CoverTab[90056]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:817
	_go_fuzz_dep_.CoverTab[90057]++
													if off > 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:818
		_go_fuzz_dep_.CoverTab[90093]++
														ioff := int(off)
														if len(out) < dstEvery*3+ioff {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:820
			_go_fuzz_dep_.CoverTab[90095]++
															return nil, errors.New("corruption detected: stream overrun 3")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:821
			// _ = "end of CoverTab[90095]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:822
			_go_fuzz_dep_.CoverTab[90096]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:822
			// _ = "end of CoverTab[90096]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:822
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:822
		// _ = "end of CoverTab[90093]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:822
		_go_fuzz_dep_.CoverTab[90094]++
														copy(out, buf[:off])
														copy(out[dstEvery:dstEvery+ioff], buf[bufoff:bufoff*2])
														copy(out[dstEvery*2:dstEvery*2+ioff], buf[bufoff*2:bufoff*3])
														copy(out[dstEvery*3:dstEvery*3+ioff], buf[bufoff*3:bufoff*4])
														decoded += int(off) * 4
														out = out[off:]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:828
		// _ = "end of CoverTab[90094]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:829
		_go_fuzz_dep_.CoverTab[90097]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:829
		// _ = "end of CoverTab[90097]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:829
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:829
	// _ = "end of CoverTab[90057]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:829
	_go_fuzz_dep_.CoverTab[90058]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:832
	for i := range br {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:832
		_go_fuzz_dep_.CoverTab[90098]++
														offset := dstEvery * i
														br := &br[i]
														bitsLeft := br.off*8 + uint(64-br.bitsRead)
														for bitsLeft > 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:836
			_go_fuzz_dep_.CoverTab[90100]++
															br.fill()
															if false && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:838
				_go_fuzz_dep_.CoverTab[90103]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:838
				return br.bitsRead >= 32
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:838
				// _ = "end of CoverTab[90103]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:838
			}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:838
				_go_fuzz_dep_.CoverTab[90104]++
																if br.off >= 4 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:839
					_go_fuzz_dep_.CoverTab[90105]++
																	v := br.in[br.off-4:]
																	v = v[:4]
																	low := (uint32(v[0])) | (uint32(v[1]) << 8) | (uint32(v[2]) << 16) | (uint32(v[3]) << 24)
																	br.value = (br.value << 32) | uint64(low)
																	br.bitsRead -= 32
																	br.off -= 4
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:845
					// _ = "end of CoverTab[90105]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:846
					_go_fuzz_dep_.CoverTab[90106]++
																	for br.off > 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:847
						_go_fuzz_dep_.CoverTab[90107]++
																		br.value = (br.value << 8) | uint64(br.in[br.off-1])
																		br.bitsRead -= 8
																		br.off--
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:850
						// _ = "end of CoverTab[90107]"
					}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:851
					// _ = "end of CoverTab[90106]"
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:852
				// _ = "end of CoverTab[90104]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:853
				_go_fuzz_dep_.CoverTab[90108]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:853
				// _ = "end of CoverTab[90108]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:853
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:853
			// _ = "end of CoverTab[90100]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:853
			_go_fuzz_dep_.CoverTab[90101]++

															if offset >= len(out) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:855
				_go_fuzz_dep_.CoverTab[90109]++
																return nil, errors.New("corruption detected: stream overrun 4")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:856
				// _ = "end of CoverTab[90109]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:857
				_go_fuzz_dep_.CoverTab[90110]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:857
				// _ = "end of CoverTab[90110]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:857
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:857
			// _ = "end of CoverTab[90101]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:857
			_go_fuzz_dep_.CoverTab[90102]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:860
			val := br.peekBitsFast(d.actualTableLog)
															v := single[val&tlMask].entry
															nBits := uint8(v)
															br.advance(nBits)
															bitsLeft -= uint(nBits)
															out[offset] = uint8(v >> 8)
															offset++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:866
			// _ = "end of CoverTab[90102]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:867
		// _ = "end of CoverTab[90098]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:867
		_go_fuzz_dep_.CoverTab[90099]++
														decoded += offset - dstEvery*i
														err = br.close()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:870
			_go_fuzz_dep_.CoverTab[90111]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:871
			// _ = "end of CoverTab[90111]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:872
			_go_fuzz_dep_.CoverTab[90112]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:872
			// _ = "end of CoverTab[90112]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:872
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:872
		// _ = "end of CoverTab[90099]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:873
	// _ = "end of CoverTab[90058]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:873
	_go_fuzz_dep_.CoverTab[90059]++
													if dstSize != decoded {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:874
		_go_fuzz_dep_.CoverTab[90113]++
														return nil, errors.New("corruption detected: short output block")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:875
		// _ = "end of CoverTab[90113]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:876
		_go_fuzz_dep_.CoverTab[90114]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:876
		// _ = "end of CoverTab[90114]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:876
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:876
	// _ = "end of CoverTab[90059]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:876
	_go_fuzz_dep_.CoverTab[90060]++
													return dst, nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:877
	// _ = "end of CoverTab[90060]"
}

// Decompress4X will decompress a 4X encoded stream.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:880
// The length of the supplied input must match the end of a block exactly.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:880
// The *capacity* of the dst slice must match the destination size of
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:880
// the uncompressed data exactly.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:884
func (d *Decoder) decompress4X8bit(dst, src []byte) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:884
	_go_fuzz_dep_.CoverTab[90115]++
													if d.actualTableLog == 8 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:885
		_go_fuzz_dep_.CoverTab[90123]++
														return d.decompress4X8bitExactly(dst, src)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:886
		// _ = "end of CoverTab[90123]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:887
		_go_fuzz_dep_.CoverTab[90124]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:887
		// _ = "end of CoverTab[90124]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:887
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:887
	// _ = "end of CoverTab[90115]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:887
	_go_fuzz_dep_.CoverTab[90116]++

													var br [4]bitReaderBytes
													start := 6
													for i := 0; i < 3; i++ {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:891
		_go_fuzz_dep_.CoverTab[90125]++
														length := int(src[i*2]) | (int(src[i*2+1]) << 8)
														if start+length >= len(src) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:893
			_go_fuzz_dep_.CoverTab[90128]++
															return nil, errors.New("truncated input (or invalid offset)")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:894
			// _ = "end of CoverTab[90128]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:895
			_go_fuzz_dep_.CoverTab[90129]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:895
			// _ = "end of CoverTab[90129]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:895
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:895
		// _ = "end of CoverTab[90125]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:895
		_go_fuzz_dep_.CoverTab[90126]++
														err := br[i].init(src[start : start+length])
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:897
			_go_fuzz_dep_.CoverTab[90130]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:898
			// _ = "end of CoverTab[90130]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:899
			_go_fuzz_dep_.CoverTab[90131]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:899
			// _ = "end of CoverTab[90131]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:899
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:899
		// _ = "end of CoverTab[90126]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:899
		_go_fuzz_dep_.CoverTab[90127]++
														start += length
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:900
		// _ = "end of CoverTab[90127]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:901
	// _ = "end of CoverTab[90116]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:901
	_go_fuzz_dep_.CoverTab[90117]++
													err := br[3].init(src[start:])
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:903
		_go_fuzz_dep_.CoverTab[90132]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:904
		// _ = "end of CoverTab[90132]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:905
		_go_fuzz_dep_.CoverTab[90133]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:905
		// _ = "end of CoverTab[90133]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:905
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:905
	// _ = "end of CoverTab[90117]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:905
	_go_fuzz_dep_.CoverTab[90118]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:908
	dstSize := cap(dst)
	dst = dst[:dstSize]
	out := dst
	dstEvery := (dstSize + 3) / 4

	shift := (56 + (8 - d.actualTableLog)) & 63

	const tlSize = 1 << 8
	single := d.dt.single[:tlSize]

	// Use temp table to avoid bound checks/append penalty.
	var buf [256]byte
	var off uint8
	var decoded int

	// Decode 4 values from each decoder/loop.
	const bufoff = 256 / 4
	for {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:925
		_go_fuzz_dep_.CoverTab[90134]++
														if br[0].off < 4 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:926
			_go_fuzz_dep_.CoverTab[90136]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:926
			return br[1].off < 4
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:926
			// _ = "end of CoverTab[90136]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:926
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:926
			_go_fuzz_dep_.CoverTab[90137]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:926
			return br[2].off < 4
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:926
			// _ = "end of CoverTab[90137]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:926
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:926
			_go_fuzz_dep_.CoverTab[90138]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:926
			return br[3].off < 4
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:926
			// _ = "end of CoverTab[90138]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:926
		}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:926
			_go_fuzz_dep_.CoverTab[90139]++
															break
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:927
			// _ = "end of CoverTab[90139]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:928
			_go_fuzz_dep_.CoverTab[90140]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:928
			// _ = "end of CoverTab[90140]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:928
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:928
		// _ = "end of CoverTab[90134]"

														{
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:930
			_go_fuzz_dep_.CoverTab[90141]++
															// Interleave 2 decodes.
															const stream = 0
															const stream2 = 1
															br1 := &br[stream]
															br2 := &br[stream2]
															br1.fillFast()
															br2.fillFast()

															v := single[uint8(br1.value>>shift)].entry
															v2 := single[uint8(br2.value>>shift)].entry
															br1.bitsRead += uint8(v)
															br1.value <<= v & 63
															br2.bitsRead += uint8(v2)
															br2.value <<= v2 & 63
															buf[off+bufoff*stream] = uint8(v >> 8)
															buf[off+bufoff*stream2] = uint8(v2 >> 8)

															v = single[uint8(br1.value>>shift)].entry
															v2 = single[uint8(br2.value>>shift)].entry
															br1.bitsRead += uint8(v)
															br1.value <<= v & 63
															br2.bitsRead += uint8(v2)
															br2.value <<= v2 & 63
															buf[off+bufoff*stream+1] = uint8(v >> 8)
															buf[off+bufoff*stream2+1] = uint8(v2 >> 8)

															v = single[uint8(br1.value>>shift)].entry
															v2 = single[uint8(br2.value>>shift)].entry
															br1.bitsRead += uint8(v)
															br1.value <<= v & 63
															br2.bitsRead += uint8(v2)
															br2.value <<= v2 & 63
															buf[off+bufoff*stream+2] = uint8(v >> 8)
															buf[off+bufoff*stream2+2] = uint8(v2 >> 8)

															v = single[uint8(br1.value>>shift)].entry
															v2 = single[uint8(br2.value>>shift)].entry
															br1.bitsRead += uint8(v)
															br1.value <<= v & 63
															br2.bitsRead += uint8(v2)
															br2.value <<= v2 & 63
															buf[off+bufoff*stream2+3] = uint8(v2 >> 8)
															buf[off+bufoff*stream+3] = uint8(v >> 8)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:973
			// _ = "end of CoverTab[90141]"
		}

		{
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:976
			_go_fuzz_dep_.CoverTab[90142]++
															const stream = 2
															const stream2 = 3
															br1 := &br[stream]
															br2 := &br[stream2]
															br1.fillFast()
															br2.fillFast()

															v := single[uint8(br1.value>>shift)].entry
															v2 := single[uint8(br2.value>>shift)].entry
															br1.bitsRead += uint8(v)
															br1.value <<= v & 63
															br2.bitsRead += uint8(v2)
															br2.value <<= v2 & 63
															buf[off+bufoff*stream] = uint8(v >> 8)
															buf[off+bufoff*stream2] = uint8(v2 >> 8)

															v = single[uint8(br1.value>>shift)].entry
															v2 = single[uint8(br2.value>>shift)].entry
															br1.bitsRead += uint8(v)
															br1.value <<= v & 63
															br2.bitsRead += uint8(v2)
															br2.value <<= v2 & 63
															buf[off+bufoff*stream+1] = uint8(v >> 8)
															buf[off+bufoff*stream2+1] = uint8(v2 >> 8)

															v = single[uint8(br1.value>>shift)].entry
															v2 = single[uint8(br2.value>>shift)].entry
															br1.bitsRead += uint8(v)
															br1.value <<= v & 63
															br2.bitsRead += uint8(v2)
															br2.value <<= v2 & 63
															buf[off+bufoff*stream+2] = uint8(v >> 8)
															buf[off+bufoff*stream2+2] = uint8(v2 >> 8)

															v = single[uint8(br1.value>>shift)].entry
															v2 = single[uint8(br2.value>>shift)].entry
															br1.bitsRead += uint8(v)
															br1.value <<= v & 63
															br2.bitsRead += uint8(v2)
															br2.value <<= v2 & 63
															buf[off+bufoff*stream2+3] = uint8(v2 >> 8)
															buf[off+bufoff*stream+3] = uint8(v >> 8)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1018
			// _ = "end of CoverTab[90142]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1019
		_go_fuzz_dep_.CoverTab[90135]++

														off += 4

														if off == bufoff {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1023
			_go_fuzz_dep_.CoverTab[90143]++
															if bufoff > dstEvery {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1024
				_go_fuzz_dep_.CoverTab[90145]++
																return nil, errors.New("corruption detected: stream overrun 1")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1025
				// _ = "end of CoverTab[90145]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1026
				_go_fuzz_dep_.CoverTab[90146]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1026
				// _ = "end of CoverTab[90146]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1026
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1026
			// _ = "end of CoverTab[90143]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1026
			_go_fuzz_dep_.CoverTab[90144]++
															copy(out, buf[:bufoff])
															copy(out[dstEvery:], buf[bufoff:bufoff*2])
															copy(out[dstEvery*2:], buf[bufoff*2:bufoff*3])
															copy(out[dstEvery*3:], buf[bufoff*3:bufoff*4])
															off = 0
															out = out[bufoff:]
															decoded += 256

															if len(out) < dstEvery*3 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1035
				_go_fuzz_dep_.CoverTab[90147]++
																return nil, errors.New("corruption detected: stream overrun 2")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1036
				// _ = "end of CoverTab[90147]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1037
				_go_fuzz_dep_.CoverTab[90148]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1037
				// _ = "end of CoverTab[90148]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1037
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1037
			// _ = "end of CoverTab[90144]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1038
			_go_fuzz_dep_.CoverTab[90149]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1038
			// _ = "end of CoverTab[90149]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1038
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1038
		// _ = "end of CoverTab[90135]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1039
	// _ = "end of CoverTab[90118]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1039
	_go_fuzz_dep_.CoverTab[90119]++
													if off > 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1040
		_go_fuzz_dep_.CoverTab[90150]++
														ioff := int(off)
														if len(out) < dstEvery*3+ioff {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1042
			_go_fuzz_dep_.CoverTab[90152]++
															return nil, errors.New("corruption detected: stream overrun 3")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1043
			// _ = "end of CoverTab[90152]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1044
			_go_fuzz_dep_.CoverTab[90153]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1044
			// _ = "end of CoverTab[90153]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1044
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1044
		// _ = "end of CoverTab[90150]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1044
		_go_fuzz_dep_.CoverTab[90151]++
														copy(out, buf[:off])
														copy(out[dstEvery:dstEvery+ioff], buf[bufoff:bufoff*2])
														copy(out[dstEvery*2:dstEvery*2+ioff], buf[bufoff*2:bufoff*3])
														copy(out[dstEvery*3:dstEvery*3+ioff], buf[bufoff*3:bufoff*4])
														decoded += int(off) * 4
														out = out[off:]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1050
		// _ = "end of CoverTab[90151]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1051
		_go_fuzz_dep_.CoverTab[90154]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1051
		// _ = "end of CoverTab[90154]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1051
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1051
	// _ = "end of CoverTab[90119]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1051
	_go_fuzz_dep_.CoverTab[90120]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1054
	for i := range br {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1054
		_go_fuzz_dep_.CoverTab[90155]++
														offset := dstEvery * i
														br := &br[i]
														bitsLeft := int(br.off*8) + int(64-br.bitsRead)
														for bitsLeft > 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1058
			_go_fuzz_dep_.CoverTab[90157]++
															if br.finished() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1059
				_go_fuzz_dep_.CoverTab[90161]++
																return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1060
				// _ = "end of CoverTab[90161]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1061
				_go_fuzz_dep_.CoverTab[90162]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1061
				// _ = "end of CoverTab[90162]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1061
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1061
			// _ = "end of CoverTab[90157]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1061
			_go_fuzz_dep_.CoverTab[90158]++
															if br.bitsRead >= 56 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1062
				_go_fuzz_dep_.CoverTab[90163]++
																if br.off >= 4 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1063
					_go_fuzz_dep_.CoverTab[90164]++
																	v := br.in[br.off-4:]
																	v = v[:4]
																	low := (uint32(v[0])) | (uint32(v[1]) << 8) | (uint32(v[2]) << 16) | (uint32(v[3]) << 24)
																	br.value |= uint64(low) << (br.bitsRead - 32)
																	br.bitsRead -= 32
																	br.off -= 4
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1069
					// _ = "end of CoverTab[90164]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1070
					_go_fuzz_dep_.CoverTab[90165]++
																	for br.off > 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1071
						_go_fuzz_dep_.CoverTab[90166]++
																		br.value |= uint64(br.in[br.off-1]) << (br.bitsRead - 8)
																		br.bitsRead -= 8
																		br.off--
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1074
						// _ = "end of CoverTab[90166]"
					}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1075
					// _ = "end of CoverTab[90165]"
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1076
				// _ = "end of CoverTab[90163]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1077
				_go_fuzz_dep_.CoverTab[90167]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1077
				// _ = "end of CoverTab[90167]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1077
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1077
			// _ = "end of CoverTab[90158]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1077
			_go_fuzz_dep_.CoverTab[90159]++

															if offset >= len(out) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1079
				_go_fuzz_dep_.CoverTab[90168]++
																return nil, errors.New("corruption detected: stream overrun 4")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1080
				// _ = "end of CoverTab[90168]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1081
				_go_fuzz_dep_.CoverTab[90169]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1081
				// _ = "end of CoverTab[90169]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1081
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1081
			// _ = "end of CoverTab[90159]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1081
			_go_fuzz_dep_.CoverTab[90160]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1084
			v := single[uint8(br.value>>shift)].entry
															nBits := uint8(v)
															br.advance(nBits)
															bitsLeft -= int(nBits)
															out[offset] = uint8(v >> 8)
															offset++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1089
			// _ = "end of CoverTab[90160]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1090
		// _ = "end of CoverTab[90155]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1090
		_go_fuzz_dep_.CoverTab[90156]++
														decoded += offset - dstEvery*i
														err = br.close()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1093
			_go_fuzz_dep_.CoverTab[90170]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1094
			// _ = "end of CoverTab[90170]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1095
			_go_fuzz_dep_.CoverTab[90171]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1095
			// _ = "end of CoverTab[90171]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1095
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1095
		// _ = "end of CoverTab[90156]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1096
	// _ = "end of CoverTab[90120]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1096
	_go_fuzz_dep_.CoverTab[90121]++
													if dstSize != decoded {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1097
		_go_fuzz_dep_.CoverTab[90172]++
														return nil, errors.New("corruption detected: short output block")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1098
		// _ = "end of CoverTab[90172]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1099
		_go_fuzz_dep_.CoverTab[90173]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1099
		// _ = "end of CoverTab[90173]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1099
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1099
	// _ = "end of CoverTab[90121]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1099
	_go_fuzz_dep_.CoverTab[90122]++
													return dst, nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1100
	// _ = "end of CoverTab[90122]"
}

// Decompress4X will decompress a 4X encoded stream.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1103
// The length of the supplied input must match the end of a block exactly.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1103
// The *capacity* of the dst slice must match the destination size of
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1103
// the uncompressed data exactly.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1107
func (d *Decoder) decompress4X8bitExactly(dst, src []byte) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1107
	_go_fuzz_dep_.CoverTab[90174]++
													var br [4]bitReaderBytes
													start := 6
													for i := 0; i < 3; i++ {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1110
		_go_fuzz_dep_.CoverTab[90181]++
														length := int(src[i*2]) | (int(src[i*2+1]) << 8)
														if start+length >= len(src) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1112
			_go_fuzz_dep_.CoverTab[90184]++
															return nil, errors.New("truncated input (or invalid offset)")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1113
			// _ = "end of CoverTab[90184]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1114
			_go_fuzz_dep_.CoverTab[90185]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1114
			// _ = "end of CoverTab[90185]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1114
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1114
		// _ = "end of CoverTab[90181]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1114
		_go_fuzz_dep_.CoverTab[90182]++
														err := br[i].init(src[start : start+length])
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1116
			_go_fuzz_dep_.CoverTab[90186]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1117
			// _ = "end of CoverTab[90186]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1118
			_go_fuzz_dep_.CoverTab[90187]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1118
			// _ = "end of CoverTab[90187]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1118
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1118
		// _ = "end of CoverTab[90182]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1118
		_go_fuzz_dep_.CoverTab[90183]++
														start += length
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1119
		// _ = "end of CoverTab[90183]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1120
	// _ = "end of CoverTab[90174]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1120
	_go_fuzz_dep_.CoverTab[90175]++
													err := br[3].init(src[start:])
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1122
		_go_fuzz_dep_.CoverTab[90188]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1123
		// _ = "end of CoverTab[90188]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1124
		_go_fuzz_dep_.CoverTab[90189]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1124
		// _ = "end of CoverTab[90189]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1124
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1124
	// _ = "end of CoverTab[90175]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1124
	_go_fuzz_dep_.CoverTab[90176]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1127
	dstSize := cap(dst)
	dst = dst[:dstSize]
	out := dst
	dstEvery := (dstSize + 3) / 4

	const shift = 56
	const tlSize = 1 << 8
	const tlMask = tlSize - 1
	single := d.dt.single[:tlSize]

	// Use temp table to avoid bound checks/append penalty.
	var buf [256]byte
	var off uint8
	var decoded int

	// Decode 4 values from each decoder/loop.
	const bufoff = 256 / 4
	for {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1144
		_go_fuzz_dep_.CoverTab[90190]++
														if br[0].off < 4 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1145
			_go_fuzz_dep_.CoverTab[90192]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1145
			return br[1].off < 4
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1145
			// _ = "end of CoverTab[90192]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1145
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1145
			_go_fuzz_dep_.CoverTab[90193]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1145
			return br[2].off < 4
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1145
			// _ = "end of CoverTab[90193]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1145
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1145
			_go_fuzz_dep_.CoverTab[90194]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1145
			return br[3].off < 4
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1145
			// _ = "end of CoverTab[90194]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1145
		}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1145
			_go_fuzz_dep_.CoverTab[90195]++
															break
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1146
			// _ = "end of CoverTab[90195]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1147
			_go_fuzz_dep_.CoverTab[90196]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1147
			// _ = "end of CoverTab[90196]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1147
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1147
		// _ = "end of CoverTab[90190]"

														{
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1149
			_go_fuzz_dep_.CoverTab[90197]++
															// Interleave 2 decodes.
															const stream = 0
															const stream2 = 1
															br[stream].fillFast()
															br[stream2].fillFast()

															v := single[uint8(br[stream].value>>shift)].entry
															v2 := single[uint8(br[stream2].value>>shift)].entry
															br[stream].bitsRead += uint8(v)
															br[stream].value <<= v & 63
															br[stream2].bitsRead += uint8(v2)
															br[stream2].value <<= v2 & 63
															buf[off+bufoff*stream] = uint8(v >> 8)
															buf[off+bufoff*stream2] = uint8(v2 >> 8)

															v = single[uint8(br[stream].value>>shift)].entry
															v2 = single[uint8(br[stream2].value>>shift)].entry
															br[stream].bitsRead += uint8(v)
															br[stream].value <<= v & 63
															br[stream2].bitsRead += uint8(v2)
															br[stream2].value <<= v2 & 63
															buf[off+bufoff*stream+1] = uint8(v >> 8)
															buf[off+bufoff*stream2+1] = uint8(v2 >> 8)

															v = single[uint8(br[stream].value>>shift)].entry
															v2 = single[uint8(br[stream2].value>>shift)].entry
															br[stream].bitsRead += uint8(v)
															br[stream].value <<= v & 63
															br[stream2].bitsRead += uint8(v2)
															br[stream2].value <<= v2 & 63
															buf[off+bufoff*stream+2] = uint8(v >> 8)
															buf[off+bufoff*stream2+2] = uint8(v2 >> 8)

															v = single[uint8(br[stream].value>>shift)].entry
															v2 = single[uint8(br[stream2].value>>shift)].entry
															br[stream].bitsRead += uint8(v)
															br[stream].value <<= v & 63
															br[stream2].bitsRead += uint8(v2)
															br[stream2].value <<= v2 & 63
															buf[off+bufoff*stream+3] = uint8(v >> 8)
															buf[off+bufoff*stream2+3] = uint8(v2 >> 8)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1190
			// _ = "end of CoverTab[90197]"
		}

		{
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1193
			_go_fuzz_dep_.CoverTab[90198]++
															const stream = 2
															const stream2 = 3
															br[stream].fillFast()
															br[stream2].fillFast()

															v := single[uint8(br[stream].value>>shift)].entry
															v2 := single[uint8(br[stream2].value>>shift)].entry
															br[stream].bitsRead += uint8(v)
															br[stream].value <<= v & 63
															br[stream2].bitsRead += uint8(v2)
															br[stream2].value <<= v2 & 63
															buf[off+bufoff*stream] = uint8(v >> 8)
															buf[off+bufoff*stream2] = uint8(v2 >> 8)

															v = single[uint8(br[stream].value>>shift)].entry
															v2 = single[uint8(br[stream2].value>>shift)].entry
															br[stream].bitsRead += uint8(v)
															br[stream].value <<= v & 63
															br[stream2].bitsRead += uint8(v2)
															br[stream2].value <<= v2 & 63
															buf[off+bufoff*stream+1] = uint8(v >> 8)
															buf[off+bufoff*stream2+1] = uint8(v2 >> 8)

															v = single[uint8(br[stream].value>>shift)].entry
															v2 = single[uint8(br[stream2].value>>shift)].entry
															br[stream].bitsRead += uint8(v)
															br[stream].value <<= v & 63
															br[stream2].bitsRead += uint8(v2)
															br[stream2].value <<= v2 & 63
															buf[off+bufoff*stream+2] = uint8(v >> 8)
															buf[off+bufoff*stream2+2] = uint8(v2 >> 8)

															v = single[uint8(br[stream].value>>shift)].entry
															v2 = single[uint8(br[stream2].value>>shift)].entry
															br[stream].bitsRead += uint8(v)
															br[stream].value <<= v & 63
															br[stream2].bitsRead += uint8(v2)
															br[stream2].value <<= v2 & 63
															buf[off+bufoff*stream+3] = uint8(v >> 8)
															buf[off+bufoff*stream2+3] = uint8(v2 >> 8)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1233
			// _ = "end of CoverTab[90198]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1234
		_go_fuzz_dep_.CoverTab[90191]++

														off += 4

														if off == bufoff {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1238
			_go_fuzz_dep_.CoverTab[90199]++
															if bufoff > dstEvery {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1239
				_go_fuzz_dep_.CoverTab[90201]++
																return nil, errors.New("corruption detected: stream overrun 1")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1240
				// _ = "end of CoverTab[90201]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1241
				_go_fuzz_dep_.CoverTab[90202]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1241
				// _ = "end of CoverTab[90202]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1241
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1241
			// _ = "end of CoverTab[90199]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1241
			_go_fuzz_dep_.CoverTab[90200]++
															copy(out, buf[:bufoff])
															copy(out[dstEvery:], buf[bufoff:bufoff*2])
															copy(out[dstEvery*2:], buf[bufoff*2:bufoff*3])
															copy(out[dstEvery*3:], buf[bufoff*3:bufoff*4])
															off = 0
															out = out[bufoff:]
															decoded += 256

															if len(out) < dstEvery*3 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1250
				_go_fuzz_dep_.CoverTab[90203]++
																return nil, errors.New("corruption detected: stream overrun 2")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1251
				// _ = "end of CoverTab[90203]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1252
				_go_fuzz_dep_.CoverTab[90204]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1252
				// _ = "end of CoverTab[90204]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1252
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1252
			// _ = "end of CoverTab[90200]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1253
			_go_fuzz_dep_.CoverTab[90205]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1253
			// _ = "end of CoverTab[90205]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1253
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1253
		// _ = "end of CoverTab[90191]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1254
	// _ = "end of CoverTab[90176]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1254
	_go_fuzz_dep_.CoverTab[90177]++
													if off > 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1255
		_go_fuzz_dep_.CoverTab[90206]++
														ioff := int(off)
														if len(out) < dstEvery*3+ioff {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1257
			_go_fuzz_dep_.CoverTab[90208]++
															return nil, errors.New("corruption detected: stream overrun 3")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1258
			// _ = "end of CoverTab[90208]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1259
			_go_fuzz_dep_.CoverTab[90209]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1259
			// _ = "end of CoverTab[90209]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1259
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1259
		// _ = "end of CoverTab[90206]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1259
		_go_fuzz_dep_.CoverTab[90207]++
														copy(out, buf[:off])
														copy(out[dstEvery:dstEvery+ioff], buf[bufoff:bufoff*2])
														copy(out[dstEvery*2:dstEvery*2+ioff], buf[bufoff*2:bufoff*3])
														copy(out[dstEvery*3:dstEvery*3+ioff], buf[bufoff*3:bufoff*4])
														decoded += int(off) * 4
														out = out[off:]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1265
		// _ = "end of CoverTab[90207]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1266
		_go_fuzz_dep_.CoverTab[90210]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1266
		// _ = "end of CoverTab[90210]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1266
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1266
	// _ = "end of CoverTab[90177]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1266
	_go_fuzz_dep_.CoverTab[90178]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1269
	for i := range br {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1269
		_go_fuzz_dep_.CoverTab[90211]++
														offset := dstEvery * i
														br := &br[i]
														bitsLeft := int(br.off*8) + int(64-br.bitsRead)
														for bitsLeft > 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1273
			_go_fuzz_dep_.CoverTab[90213]++
															if br.finished() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1274
				_go_fuzz_dep_.CoverTab[90217]++
																return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1275
				// _ = "end of CoverTab[90217]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1276
				_go_fuzz_dep_.CoverTab[90218]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1276
				// _ = "end of CoverTab[90218]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1276
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1276
			// _ = "end of CoverTab[90213]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1276
			_go_fuzz_dep_.CoverTab[90214]++
															if br.bitsRead >= 56 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1277
				_go_fuzz_dep_.CoverTab[90219]++
																if br.off >= 4 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1278
					_go_fuzz_dep_.CoverTab[90220]++
																	v := br.in[br.off-4:]
																	v = v[:4]
																	low := (uint32(v[0])) | (uint32(v[1]) << 8) | (uint32(v[2]) << 16) | (uint32(v[3]) << 24)
																	br.value |= uint64(low) << (br.bitsRead - 32)
																	br.bitsRead -= 32
																	br.off -= 4
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1284
					// _ = "end of CoverTab[90220]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1285
					_go_fuzz_dep_.CoverTab[90221]++
																	for br.off > 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1286
						_go_fuzz_dep_.CoverTab[90222]++
																		br.value |= uint64(br.in[br.off-1]) << (br.bitsRead - 8)
																		br.bitsRead -= 8
																		br.off--
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1289
						// _ = "end of CoverTab[90222]"
					}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1290
					// _ = "end of CoverTab[90221]"
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1291
				// _ = "end of CoverTab[90219]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1292
				_go_fuzz_dep_.CoverTab[90223]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1292
				// _ = "end of CoverTab[90223]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1292
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1292
			// _ = "end of CoverTab[90214]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1292
			_go_fuzz_dep_.CoverTab[90215]++

															if offset >= len(out) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1294
				_go_fuzz_dep_.CoverTab[90224]++
																return nil, errors.New("corruption detected: stream overrun 4")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1295
				// _ = "end of CoverTab[90224]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1296
				_go_fuzz_dep_.CoverTab[90225]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1296
				// _ = "end of CoverTab[90225]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1296
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1296
			// _ = "end of CoverTab[90215]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1296
			_go_fuzz_dep_.CoverTab[90216]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1299
			v := single[br.peekByteFast()].entry
															nBits := uint8(v)
															br.advance(nBits)
															bitsLeft -= int(nBits)
															out[offset] = uint8(v >> 8)
															offset++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1304
			// _ = "end of CoverTab[90216]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1305
		// _ = "end of CoverTab[90211]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1305
		_go_fuzz_dep_.CoverTab[90212]++
														decoded += offset - dstEvery*i
														err = br.close()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1308
			_go_fuzz_dep_.CoverTab[90226]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1309
			// _ = "end of CoverTab[90226]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1310
			_go_fuzz_dep_.CoverTab[90227]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1310
			// _ = "end of CoverTab[90227]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1310
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1310
		// _ = "end of CoverTab[90212]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1311
	// _ = "end of CoverTab[90178]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1311
	_go_fuzz_dep_.CoverTab[90179]++
													if dstSize != decoded {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1312
		_go_fuzz_dep_.CoverTab[90228]++
														return nil, errors.New("corruption detected: short output block")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1313
		// _ = "end of CoverTab[90228]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1314
		_go_fuzz_dep_.CoverTab[90229]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1314
		// _ = "end of CoverTab[90229]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1314
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1314
	// _ = "end of CoverTab[90179]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1314
	_go_fuzz_dep_.CoverTab[90180]++
													return dst, nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1315
	// _ = "end of CoverTab[90180]"
}

// matches will compare a decoding table to a coding table.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1318
// Errors are written to the writer.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1318
// Nothing will be written if table is ok.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1321
func (s *Scratch) matches(ct cTable, w io.Writer) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1321
	_go_fuzz_dep_.CoverTab[90230]++
													if s == nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1322
		_go_fuzz_dep_.CoverTab[90233]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1322
		return len(s.dt.single) == 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1322
		// _ = "end of CoverTab[90233]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1322
	}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1322
		_go_fuzz_dep_.CoverTab[90234]++
														return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1323
		// _ = "end of CoverTab[90234]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1324
		_go_fuzz_dep_.CoverTab[90235]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1324
		// _ = "end of CoverTab[90235]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1324
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1324
	// _ = "end of CoverTab[90230]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1324
	_go_fuzz_dep_.CoverTab[90231]++
													dt := s.dt.single[:1<<s.actualTableLog]
													tablelog := s.actualTableLog
													ok := 0
													broken := 0
													for sym, enc := range ct {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1329
		_go_fuzz_dep_.CoverTab[90236]++
														errs := 0
														broken++
														if enc.nBits == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1332
			_go_fuzz_dep_.CoverTab[90242]++
															for _, dec := range dt {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1333
				_go_fuzz_dep_.CoverTab[90245]++
																if uint8(dec.entry>>8) == byte(sym) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1334
					_go_fuzz_dep_.CoverTab[90246]++
																	fmt.Fprintf(w, "symbol %x has decoder, but no encoder\n", sym)
																	errs++
																	break
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1337
					// _ = "end of CoverTab[90246]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1338
					_go_fuzz_dep_.CoverTab[90247]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1338
					// _ = "end of CoverTab[90247]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1338
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1338
				// _ = "end of CoverTab[90245]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1339
			// _ = "end of CoverTab[90242]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1339
			_go_fuzz_dep_.CoverTab[90243]++
															if errs == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1340
				_go_fuzz_dep_.CoverTab[90248]++
																broken--
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1341
				// _ = "end of CoverTab[90248]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1342
				_go_fuzz_dep_.CoverTab[90249]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1342
				// _ = "end of CoverTab[90249]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1342
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1342
			// _ = "end of CoverTab[90243]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1342
			_go_fuzz_dep_.CoverTab[90244]++
															continue
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1343
			// _ = "end of CoverTab[90244]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1344
			_go_fuzz_dep_.CoverTab[90250]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1344
			// _ = "end of CoverTab[90250]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1344
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1344
		// _ = "end of CoverTab[90236]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1344
		_go_fuzz_dep_.CoverTab[90237]++

														ub := tablelog - enc.nBits
														top := enc.val << ub

														dec := dt[top]
														if uint8(dec.entry) != enc.nBits {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1350
			_go_fuzz_dep_.CoverTab[90251]++
															fmt.Fprintf(w, "symbol 0x%x bit size mismatch (enc: %d, dec:%d).\n", sym, enc.nBits, uint8(dec.entry))
															errs++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1352
			// _ = "end of CoverTab[90251]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1353
			_go_fuzz_dep_.CoverTab[90252]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1353
			// _ = "end of CoverTab[90252]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1353
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1353
		// _ = "end of CoverTab[90237]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1353
		_go_fuzz_dep_.CoverTab[90238]++
														if uint8(dec.entry>>8) != uint8(sym) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1354
			_go_fuzz_dep_.CoverTab[90253]++
															fmt.Fprintf(w, "symbol 0x%x decoder output mismatch (enc: %d, dec:%d).\n", sym, sym, uint8(dec.entry>>8))
															errs++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1356
			// _ = "end of CoverTab[90253]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1357
			_go_fuzz_dep_.CoverTab[90254]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1357
			// _ = "end of CoverTab[90254]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1357
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1357
		// _ = "end of CoverTab[90238]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1357
		_go_fuzz_dep_.CoverTab[90239]++
														if errs > 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1358
			_go_fuzz_dep_.CoverTab[90255]++
															fmt.Fprintf(w, "%d errros in base, stopping\n", errs)
															continue
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1360
			// _ = "end of CoverTab[90255]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1361
			_go_fuzz_dep_.CoverTab[90256]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1361
			// _ = "end of CoverTab[90256]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1361
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1361
		// _ = "end of CoverTab[90239]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1361
		_go_fuzz_dep_.CoverTab[90240]++

														for i := uint16(0); i < (1 << ub); i++ {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1363
			_go_fuzz_dep_.CoverTab[90257]++
															vval := top | i
															dec := dt[vval]
															if uint8(dec.entry) != enc.nBits {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1366
				_go_fuzz_dep_.CoverTab[90260]++
																fmt.Fprintf(w, "symbol 0x%x bit size mismatch (enc: %d, dec:%d).\n", vval, enc.nBits, uint8(dec.entry))
																errs++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1368
				// _ = "end of CoverTab[90260]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1369
				_go_fuzz_dep_.CoverTab[90261]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1369
				// _ = "end of CoverTab[90261]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1369
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1369
			// _ = "end of CoverTab[90257]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1369
			_go_fuzz_dep_.CoverTab[90258]++
															if uint8(dec.entry>>8) != uint8(sym) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1370
				_go_fuzz_dep_.CoverTab[90262]++
																fmt.Fprintf(w, "symbol 0x%x decoder output mismatch (enc: %d, dec:%d).\n", vval, sym, uint8(dec.entry>>8))
																errs++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1372
				// _ = "end of CoverTab[90262]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1373
				_go_fuzz_dep_.CoverTab[90263]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1373
				// _ = "end of CoverTab[90263]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1373
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1373
			// _ = "end of CoverTab[90258]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1373
			_go_fuzz_dep_.CoverTab[90259]++
															if errs > 20 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1374
				_go_fuzz_dep_.CoverTab[90264]++
																fmt.Fprintf(w, "%d errros, stopping\n", errs)
																break
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1376
				// _ = "end of CoverTab[90264]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1377
				_go_fuzz_dep_.CoverTab[90265]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1377
				// _ = "end of CoverTab[90265]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1377
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1377
			// _ = "end of CoverTab[90259]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1378
		// _ = "end of CoverTab[90240]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1378
		_go_fuzz_dep_.CoverTab[90241]++
														if errs == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1379
			_go_fuzz_dep_.CoverTab[90266]++
															ok++
															broken--
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1381
			// _ = "end of CoverTab[90266]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1382
			_go_fuzz_dep_.CoverTab[90267]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1382
			// _ = "end of CoverTab[90267]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1382
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1382
		// _ = "end of CoverTab[90241]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1383
	// _ = "end of CoverTab[90231]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1383
	_go_fuzz_dep_.CoverTab[90232]++
													if broken > 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1384
		_go_fuzz_dep_.CoverTab[90268]++
														fmt.Fprintf(w, "%d broken, %d ok\n", broken, ok)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1385
		// _ = "end of CoverTab[90268]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1386
		_go_fuzz_dep_.CoverTab[90269]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1386
		// _ = "end of CoverTab[90269]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1386
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1386
	// _ = "end of CoverTab[90232]"
}

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1387
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/decompress.go:1387
var _ = _go_fuzz_dep_.CoverTab
