//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:1
package fse

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:1
)

import (
	"errors"
	"fmt"
)

const (
	tablelogAbsoluteMax = 15
)

// Decompress a block of data.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:12
// You can provide a scratch buffer to avoid allocations.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:12
// If nil is provided a temporary one will be allocated.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:12
// It is possible, but by no way guaranteed that corrupt data will
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:12
// return an error.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:12
// It is up to the caller to verify integrity of the returned data.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:12
// Use a predefined Scrach to set maximum acceptable output size.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:19
func Decompress(b []byte, s *Scratch) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:19
	_go_fuzz_dep_.CoverTab[89263]++
												s, err := s.prepare(b)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:21
		_go_fuzz_dep_.CoverTab[89268]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:22
		// _ = "end of CoverTab[89268]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:23
		_go_fuzz_dep_.CoverTab[89269]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:23
		// _ = "end of CoverTab[89269]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:23
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:23
	// _ = "end of CoverTab[89263]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:23
	_go_fuzz_dep_.CoverTab[89264]++
												s.Out = s.Out[:0]
												err = s.readNCount()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:26
		_go_fuzz_dep_.CoverTab[89270]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:27
		// _ = "end of CoverTab[89270]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:28
		_go_fuzz_dep_.CoverTab[89271]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:28
		// _ = "end of CoverTab[89271]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:28
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:28
	// _ = "end of CoverTab[89264]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:28
	_go_fuzz_dep_.CoverTab[89265]++
												err = s.buildDtable()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:30
		_go_fuzz_dep_.CoverTab[89272]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:31
		// _ = "end of CoverTab[89272]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:32
		_go_fuzz_dep_.CoverTab[89273]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:32
		// _ = "end of CoverTab[89273]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:32
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:32
	// _ = "end of CoverTab[89265]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:32
	_go_fuzz_dep_.CoverTab[89266]++
												err = s.decompress()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:34
		_go_fuzz_dep_.CoverTab[89274]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:35
		// _ = "end of CoverTab[89274]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:36
		_go_fuzz_dep_.CoverTab[89275]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:36
		// _ = "end of CoverTab[89275]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:36
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:36
	// _ = "end of CoverTab[89266]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:36
	_go_fuzz_dep_.CoverTab[89267]++

												return s.Out, nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:38
	// _ = "end of CoverTab[89267]"
}

// readNCount will read the symbol distribution so decoding tables can be constructed.
func (s *Scratch) readNCount() error {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:42
	_go_fuzz_dep_.CoverTab[89276]++
												var (
		charnum		uint16
		previous0	bool
		b		= &s.br
	)
	iend := b.remain()
	if iend < 4 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:49
		_go_fuzz_dep_.CoverTab[89285]++
													return errors.New("input too small")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:50
		// _ = "end of CoverTab[89285]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:51
		_go_fuzz_dep_.CoverTab[89286]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:51
		// _ = "end of CoverTab[89286]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:51
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:51
	// _ = "end of CoverTab[89276]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:51
	_go_fuzz_dep_.CoverTab[89277]++
												bitStream := b.Uint32()
												nbBits := uint((bitStream & 0xF) + minTablelog)
												if nbBits > tablelogAbsoluteMax {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:54
		_go_fuzz_dep_.CoverTab[89287]++
													return errors.New("tableLog too large")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:55
		// _ = "end of CoverTab[89287]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:56
		_go_fuzz_dep_.CoverTab[89288]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:56
		// _ = "end of CoverTab[89288]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:56
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:56
	// _ = "end of CoverTab[89277]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:56
	_go_fuzz_dep_.CoverTab[89278]++
												bitStream >>= 4
												bitCount := uint(4)

												s.actualTableLog = uint8(nbBits)
												remaining := int32((1 << nbBits) + 1)
												threshold := int32(1 << nbBits)
												gotTotal := int32(0)
												nbBits++

												for remaining > 1 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:66
		_go_fuzz_dep_.CoverTab[89289]++
													if previous0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:67
			_go_fuzz_dep_.CoverTab[89295]++
														n0 := charnum
														for (bitStream & 0xFFFF) == 0xFFFF {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:69
				_go_fuzz_dep_.CoverTab[89300]++
															n0 += 24
															if b.off < iend-5 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:71
					_go_fuzz_dep_.CoverTab[89301]++
																b.advance(2)
																bitStream = b.Uint32() >> bitCount
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:73
					// _ = "end of CoverTab[89301]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:74
					_go_fuzz_dep_.CoverTab[89302]++
																bitStream >>= 16
																bitCount += 16
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:76
					// _ = "end of CoverTab[89302]"
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:77
				// _ = "end of CoverTab[89300]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:78
			// _ = "end of CoverTab[89295]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:78
			_go_fuzz_dep_.CoverTab[89296]++
														for (bitStream & 3) == 3 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:79
				_go_fuzz_dep_.CoverTab[89303]++
															n0 += 3
															bitStream >>= 2
															bitCount += 2
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:82
				// _ = "end of CoverTab[89303]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:83
			// _ = "end of CoverTab[89296]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:83
			_go_fuzz_dep_.CoverTab[89297]++
														n0 += uint16(bitStream & 3)
														bitCount += 2
														if n0 > maxSymbolValue {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:86
				_go_fuzz_dep_.CoverTab[89304]++
															return errors.New("maxSymbolValue too small")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:87
				// _ = "end of CoverTab[89304]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:88
				_go_fuzz_dep_.CoverTab[89305]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:88
				// _ = "end of CoverTab[89305]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:88
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:88
			// _ = "end of CoverTab[89297]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:88
			_go_fuzz_dep_.CoverTab[89298]++
														for charnum < n0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:89
				_go_fuzz_dep_.CoverTab[89306]++
															s.norm[charnum&0xff] = 0
															charnum++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:91
				// _ = "end of CoverTab[89306]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:92
			// _ = "end of CoverTab[89298]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:92
			_go_fuzz_dep_.CoverTab[89299]++

														if b.off <= iend-7 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:94
				_go_fuzz_dep_.CoverTab[89307]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:94
				return b.off+int(bitCount>>3) <= iend-4
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:94
				// _ = "end of CoverTab[89307]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:94
			}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:94
				_go_fuzz_dep_.CoverTab[89308]++
															b.advance(bitCount >> 3)
															bitCount &= 7
															bitStream = b.Uint32() >> bitCount
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:97
				// _ = "end of CoverTab[89308]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:98
				_go_fuzz_dep_.CoverTab[89309]++
															bitStream >>= 2
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:99
				// _ = "end of CoverTab[89309]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:100
			// _ = "end of CoverTab[89299]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:101
			_go_fuzz_dep_.CoverTab[89310]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:101
			// _ = "end of CoverTab[89310]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:101
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:101
		// _ = "end of CoverTab[89289]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:101
		_go_fuzz_dep_.CoverTab[89290]++

													max := (2*(threshold) - 1) - (remaining)
													var count int32

													if (int32(bitStream) & (threshold - 1)) < max {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:106
			_go_fuzz_dep_.CoverTab[89311]++
														count = int32(bitStream) & (threshold - 1)
														bitCount += nbBits - 1
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:108
			// _ = "end of CoverTab[89311]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:109
			_go_fuzz_dep_.CoverTab[89312]++
														count = int32(bitStream) & (2*threshold - 1)
														if count >= threshold {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:111
				_go_fuzz_dep_.CoverTab[89314]++
															count -= max
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:112
				// _ = "end of CoverTab[89314]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:113
				_go_fuzz_dep_.CoverTab[89315]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:113
				// _ = "end of CoverTab[89315]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:113
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:113
			// _ = "end of CoverTab[89312]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:113
			_go_fuzz_dep_.CoverTab[89313]++
														bitCount += nbBits
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:114
			// _ = "end of CoverTab[89313]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:115
		// _ = "end of CoverTab[89290]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:115
		_go_fuzz_dep_.CoverTab[89291]++

													count--
													if count < 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:118
			_go_fuzz_dep_.CoverTab[89316]++

														remaining += count
														gotTotal -= count
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:121
			// _ = "end of CoverTab[89316]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:122
			_go_fuzz_dep_.CoverTab[89317]++
														remaining -= count
														gotTotal += count
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:124
			// _ = "end of CoverTab[89317]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:125
		// _ = "end of CoverTab[89291]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:125
		_go_fuzz_dep_.CoverTab[89292]++
													s.norm[charnum&0xff] = int16(count)
													charnum++
													previous0 = count == 0
													for remaining < threshold {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:129
			_go_fuzz_dep_.CoverTab[89318]++
														nbBits--
														threshold >>= 1
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:131
			// _ = "end of CoverTab[89318]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:132
		// _ = "end of CoverTab[89292]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:132
		_go_fuzz_dep_.CoverTab[89293]++
													if b.off <= iend-7 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:133
			_go_fuzz_dep_.CoverTab[89319]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:133
			return b.off+int(bitCount>>3) <= iend-4
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:133
			// _ = "end of CoverTab[89319]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:133
		}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:133
			_go_fuzz_dep_.CoverTab[89320]++
														b.advance(bitCount >> 3)
														bitCount &= 7
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:135
			// _ = "end of CoverTab[89320]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:136
			_go_fuzz_dep_.CoverTab[89321]++
														bitCount -= (uint)(8 * (len(b.b) - 4 - b.off))
														b.off = len(b.b) - 4
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:138
			// _ = "end of CoverTab[89321]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:139
		// _ = "end of CoverTab[89293]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:139
		_go_fuzz_dep_.CoverTab[89294]++
													bitStream = b.Uint32() >> (bitCount & 31)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:140
		// _ = "end of CoverTab[89294]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:141
	// _ = "end of CoverTab[89278]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:141
	_go_fuzz_dep_.CoverTab[89279]++
												s.symbolLen = charnum

												if s.symbolLen <= 1 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:144
		_go_fuzz_dep_.CoverTab[89322]++
													return fmt.Errorf("symbolLen (%d) too small", s.symbolLen)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:145
		// _ = "end of CoverTab[89322]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:146
		_go_fuzz_dep_.CoverTab[89323]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:146
		// _ = "end of CoverTab[89323]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:146
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:146
	// _ = "end of CoverTab[89279]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:146
	_go_fuzz_dep_.CoverTab[89280]++
												if s.symbolLen > maxSymbolValue+1 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:147
		_go_fuzz_dep_.CoverTab[89324]++
													return fmt.Errorf("symbolLen (%d) too big", s.symbolLen)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:148
		// _ = "end of CoverTab[89324]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:149
		_go_fuzz_dep_.CoverTab[89325]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:149
		// _ = "end of CoverTab[89325]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:149
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:149
	// _ = "end of CoverTab[89280]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:149
	_go_fuzz_dep_.CoverTab[89281]++
												if remaining != 1 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:150
		_go_fuzz_dep_.CoverTab[89326]++
													return fmt.Errorf("corruption detected (remaining %d != 1)", remaining)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:151
		// _ = "end of CoverTab[89326]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:152
		_go_fuzz_dep_.CoverTab[89327]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:152
		// _ = "end of CoverTab[89327]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:152
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:152
	// _ = "end of CoverTab[89281]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:152
	_go_fuzz_dep_.CoverTab[89282]++
												if bitCount > 32 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:153
		_go_fuzz_dep_.CoverTab[89328]++
													return fmt.Errorf("corruption detected (bitCount %d > 32)", bitCount)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:154
		// _ = "end of CoverTab[89328]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:155
		_go_fuzz_dep_.CoverTab[89329]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:155
		// _ = "end of CoverTab[89329]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:155
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:155
	// _ = "end of CoverTab[89282]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:155
	_go_fuzz_dep_.CoverTab[89283]++
												if gotTotal != 1<<s.actualTableLog {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:156
		_go_fuzz_dep_.CoverTab[89330]++
													return fmt.Errorf("corruption detected (total %d != %d)", gotTotal, 1<<s.actualTableLog)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:157
		// _ = "end of CoverTab[89330]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:158
		_go_fuzz_dep_.CoverTab[89331]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:158
		// _ = "end of CoverTab[89331]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:158
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:158
	// _ = "end of CoverTab[89283]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:158
	_go_fuzz_dep_.CoverTab[89284]++
												b.advance((bitCount + 7) >> 3)
												return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:160
	// _ = "end of CoverTab[89284]"
}

// decSymbol contains information about a state entry,
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:163
// Including the state offset base, the output symbol and
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:163
// the number of bits to read for the low part of the destination state.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:166
type decSymbol struct {
	newState	uint16
	symbol		uint8
	nbBits		uint8
}

// allocDtable will allocate decoding tables if they are not big enough.
func (s *Scratch) allocDtable() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:173
	_go_fuzz_dep_.CoverTab[89332]++
												tableSize := 1 << s.actualTableLog
												if cap(s.decTable) < tableSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:175
		_go_fuzz_dep_.CoverTab[89336]++
													s.decTable = make([]decSymbol, tableSize)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:176
		// _ = "end of CoverTab[89336]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:177
		_go_fuzz_dep_.CoverTab[89337]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:177
		// _ = "end of CoverTab[89337]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:177
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:177
	// _ = "end of CoverTab[89332]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:177
	_go_fuzz_dep_.CoverTab[89333]++
												s.decTable = s.decTable[:tableSize]

												if cap(s.ct.tableSymbol) < 256 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:180
		_go_fuzz_dep_.CoverTab[89338]++
													s.ct.tableSymbol = make([]byte, 256)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:181
		// _ = "end of CoverTab[89338]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:182
		_go_fuzz_dep_.CoverTab[89339]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:182
		// _ = "end of CoverTab[89339]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:182
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:182
	// _ = "end of CoverTab[89333]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:182
	_go_fuzz_dep_.CoverTab[89334]++
												s.ct.tableSymbol = s.ct.tableSymbol[:256]

												if cap(s.ct.stateTable) < 256 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:185
		_go_fuzz_dep_.CoverTab[89340]++
													s.ct.stateTable = make([]uint16, 256)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:186
		// _ = "end of CoverTab[89340]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:187
		_go_fuzz_dep_.CoverTab[89341]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:187
		// _ = "end of CoverTab[89341]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:187
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:187
	// _ = "end of CoverTab[89334]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:187
	_go_fuzz_dep_.CoverTab[89335]++
												s.ct.stateTable = s.ct.stateTable[:256]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:188
	// _ = "end of CoverTab[89335]"
}

// buildDtable will build the decoding table.
func (s *Scratch) buildDtable() error {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:192
	_go_fuzz_dep_.CoverTab[89342]++
												tableSize := uint32(1 << s.actualTableLog)
												highThreshold := tableSize - 1
												s.allocDtable()
												symbolNext := s.ct.stateTable[:256]

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:199
	s.zeroBits = false
	{
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:200
		_go_fuzz_dep_.CoverTab[89344]++
													largeLimit := int16(1 << (s.actualTableLog - 1))
													for i, v := range s.norm[:s.symbolLen] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:202
			_go_fuzz_dep_.CoverTab[89345]++
														if v == -1 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:203
				_go_fuzz_dep_.CoverTab[89346]++
															s.decTable[highThreshold].symbol = uint8(i)
															highThreshold--
															symbolNext[i] = 1
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:206
				// _ = "end of CoverTab[89346]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:207
				_go_fuzz_dep_.CoverTab[89347]++
															if v >= largeLimit {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:208
					_go_fuzz_dep_.CoverTab[89349]++
																s.zeroBits = true
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:209
					// _ = "end of CoverTab[89349]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:210
					_go_fuzz_dep_.CoverTab[89350]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:210
					// _ = "end of CoverTab[89350]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:210
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:210
				// _ = "end of CoverTab[89347]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:210
				_go_fuzz_dep_.CoverTab[89348]++
															symbolNext[i] = uint16(v)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:211
				// _ = "end of CoverTab[89348]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:212
			// _ = "end of CoverTab[89345]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:213
		// _ = "end of CoverTab[89344]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:214
	// _ = "end of CoverTab[89342]"

												{
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:216
		_go_fuzz_dep_.CoverTab[89351]++
													tableMask := tableSize - 1
													step := tableStep(tableSize)
													position := uint32(0)
													for ss, v := range s.norm[:s.symbolLen] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:220
			_go_fuzz_dep_.CoverTab[89353]++
														for i := 0; i < int(v); i++ {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:221
				_go_fuzz_dep_.CoverTab[89354]++
															s.decTable[position].symbol = uint8(ss)
															position = (position + step) & tableMask
															for position > highThreshold {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:224
					_go_fuzz_dep_.CoverTab[89355]++

																position = (position + step) & tableMask
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:226
					// _ = "end of CoverTab[89355]"
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:227
				// _ = "end of CoverTab[89354]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:228
			// _ = "end of CoverTab[89353]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:229
		// _ = "end of CoverTab[89351]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:229
		_go_fuzz_dep_.CoverTab[89352]++
													if position != 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:230
			_go_fuzz_dep_.CoverTab[89356]++

														return errors.New("corrupted input (position != 0)")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:232
			// _ = "end of CoverTab[89356]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:233
			_go_fuzz_dep_.CoverTab[89357]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:233
			// _ = "end of CoverTab[89357]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:233
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:233
		// _ = "end of CoverTab[89352]"
	}

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:237
	{
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:237
		_go_fuzz_dep_.CoverTab[89358]++
													tableSize := uint16(1 << s.actualTableLog)
													for u, v := range s.decTable {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:239
			_go_fuzz_dep_.CoverTab[89359]++
														symbol := v.symbol
														nextState := symbolNext[symbol]
														symbolNext[symbol] = nextState + 1
														nBits := s.actualTableLog - byte(highBits(uint32(nextState)))
														s.decTable[u].nbBits = nBits
														newState := (nextState << nBits) - tableSize
														if newState >= tableSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:246
				_go_fuzz_dep_.CoverTab[89362]++
															return fmt.Errorf("newState (%d) outside table size (%d)", newState, tableSize)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:247
				// _ = "end of CoverTab[89362]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:248
				_go_fuzz_dep_.CoverTab[89363]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:248
				// _ = "end of CoverTab[89363]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:248
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:248
			// _ = "end of CoverTab[89359]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:248
			_go_fuzz_dep_.CoverTab[89360]++
														if newState == uint16(u) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:249
				_go_fuzz_dep_.CoverTab[89364]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:249
				return nBits == 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:249
				// _ = "end of CoverTab[89364]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:249
			}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:249
				_go_fuzz_dep_.CoverTab[89365]++

															return fmt.Errorf("newState (%d) == oldState (%d) and no bits", newState, u)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:251
				// _ = "end of CoverTab[89365]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:252
				_go_fuzz_dep_.CoverTab[89366]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:252
				// _ = "end of CoverTab[89366]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:252
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:252
			// _ = "end of CoverTab[89360]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:252
			_go_fuzz_dep_.CoverTab[89361]++
														s.decTable[u].newState = newState
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:253
			// _ = "end of CoverTab[89361]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:254
		// _ = "end of CoverTab[89358]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:255
	_go_fuzz_dep_.CoverTab[89343]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:256
	// _ = "end of CoverTab[89343]"
}

// decompress will decompress the bitstream.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:259
// If the buffer is over-read an error is returned.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:261
func (s *Scratch) decompress() error {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:261
	_go_fuzz_dep_.CoverTab[89367]++
												br := &s.bits
												br.init(s.br.unread())

												var s1, s2 decoder

												s1.init(br, s.decTable, s.actualTableLog)
												s2.init(br, s.decTable, s.actualTableLog)

												// Use temp table to avoid bound checks/append penalty.
												var tmp = s.ct.tableSymbol[:256]
												var off uint8

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:275
	if !s.zeroBits {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:275
		_go_fuzz_dep_.CoverTab[89370]++
													for br.off >= 8 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:276
			_go_fuzz_dep_.CoverTab[89371]++
														br.fillFast()
														tmp[off+0] = s1.nextFast()
														tmp[off+1] = s2.nextFast()
														br.fillFast()
														tmp[off+2] = s1.nextFast()
														tmp[off+3] = s2.nextFast()
														off += 4

														if off == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:285
				_go_fuzz_dep_.CoverTab[89372]++
															s.Out = append(s.Out, tmp...)
															if len(s.Out) >= s.DecompressLimit {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:287
					_go_fuzz_dep_.CoverTab[89373]++
																return fmt.Errorf("output size (%d) > DecompressLimit (%d)", len(s.Out), s.DecompressLimit)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:288
					// _ = "end of CoverTab[89373]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:289
					_go_fuzz_dep_.CoverTab[89374]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:289
					// _ = "end of CoverTab[89374]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:289
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:289
				// _ = "end of CoverTab[89372]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:290
				_go_fuzz_dep_.CoverTab[89375]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:290
				// _ = "end of CoverTab[89375]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:290
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:290
			// _ = "end of CoverTab[89371]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:291
		// _ = "end of CoverTab[89370]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:292
		_go_fuzz_dep_.CoverTab[89376]++
													for br.off >= 8 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:293
			_go_fuzz_dep_.CoverTab[89377]++
														br.fillFast()
														tmp[off+0] = s1.next()
														tmp[off+1] = s2.next()
														br.fillFast()
														tmp[off+2] = s1.next()
														tmp[off+3] = s2.next()
														off += 4
														if off == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:301
				_go_fuzz_dep_.CoverTab[89378]++
															s.Out = append(s.Out, tmp...)

															if len(s.Out) >= s.DecompressLimit {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:304
					_go_fuzz_dep_.CoverTab[89379]++
																return fmt.Errorf("output size (%d) > DecompressLimit (%d)", len(s.Out), s.DecompressLimit)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:305
					// _ = "end of CoverTab[89379]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:306
					_go_fuzz_dep_.CoverTab[89380]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:306
					// _ = "end of CoverTab[89380]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:306
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:306
				// _ = "end of CoverTab[89378]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:307
				_go_fuzz_dep_.CoverTab[89381]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:307
				// _ = "end of CoverTab[89381]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:307
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:307
			// _ = "end of CoverTab[89377]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:308
		// _ = "end of CoverTab[89376]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:309
	// _ = "end of CoverTab[89367]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:309
	_go_fuzz_dep_.CoverTab[89368]++
												s.Out = append(s.Out, tmp[:off]...)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:313
	for {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:313
		_go_fuzz_dep_.CoverTab[89382]++
													if s1.finished() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:314
			_go_fuzz_dep_.CoverTab[89385]++
														s.Out = append(s.Out, s1.final(), s2.final())
														break
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:316
			// _ = "end of CoverTab[89385]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:317
			_go_fuzz_dep_.CoverTab[89386]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:317
			// _ = "end of CoverTab[89386]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:317
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:317
		// _ = "end of CoverTab[89382]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:317
		_go_fuzz_dep_.CoverTab[89383]++
													br.fill()
													s.Out = append(s.Out, s1.next())
													if s2.finished() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:320
			_go_fuzz_dep_.CoverTab[89387]++
														s.Out = append(s.Out, s2.final(), s1.final())
														break
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:322
			// _ = "end of CoverTab[89387]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:323
			_go_fuzz_dep_.CoverTab[89388]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:323
			// _ = "end of CoverTab[89388]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:323
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:323
		// _ = "end of CoverTab[89383]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:323
		_go_fuzz_dep_.CoverTab[89384]++
													s.Out = append(s.Out, s2.next())
													if len(s.Out) >= s.DecompressLimit {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:325
			_go_fuzz_dep_.CoverTab[89389]++
														return fmt.Errorf("output size (%d) > DecompressLimit (%d)", len(s.Out), s.DecompressLimit)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:326
			// _ = "end of CoverTab[89389]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:327
			_go_fuzz_dep_.CoverTab[89390]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:327
			// _ = "end of CoverTab[89390]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:327
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:327
		// _ = "end of CoverTab[89384]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:328
	// _ = "end of CoverTab[89368]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:328
	_go_fuzz_dep_.CoverTab[89369]++
												return br.close()
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:329
	// _ = "end of CoverTab[89369]"
}

// decoder keeps track of the current state and updates it from the bitstream.
type decoder struct {
	state	uint16
	br	*bitReader
	dt	[]decSymbol
}

// init will initialize the decoder and read the first state from the stream.
func (d *decoder) init(in *bitReader, dt []decSymbol, tableLog uint8) {
	d.dt = dt
	d.br = in
	d.state = in.getBits(tableLog)
}

// next returns the next symbol and sets the next state.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:346
// At least tablelog bits must be available in the bit reader.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:348
func (d *decoder) next() uint8 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:348
	_go_fuzz_dep_.CoverTab[89391]++
												n := &d.dt[d.state]
												lowBits := d.br.getBits(n.nbBits)
												d.state = n.newState + lowBits
												return n.symbol
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:352
	// _ = "end of CoverTab[89391]"
}

// finished returns true if all bits have been read from the bitstream
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:355
// and the next state would require reading bits from the input.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:357
func (d *decoder) finished() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:357
	_go_fuzz_dep_.CoverTab[89392]++
												return d.br.finished() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:358
		_go_fuzz_dep_.CoverTab[89393]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:358
		return d.dt[d.state].nbBits > 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:358
		// _ = "end of CoverTab[89393]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:358
	}()
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:358
	// _ = "end of CoverTab[89392]"
}

// final returns the current state symbol without decoding the next.
func (d *decoder) final() uint8 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:362
	_go_fuzz_dep_.CoverTab[89394]++
												return d.dt[d.state].symbol
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:363
	// _ = "end of CoverTab[89394]"
}

// nextFast returns the next symbol and sets the next state.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:366
// This can only be used if no symbols are 0 bits.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:366
// At least tablelog bits must be available in the bit reader.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:369
func (d *decoder) nextFast() uint8 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:369
	_go_fuzz_dep_.CoverTab[89395]++
												n := d.dt[d.state]
												lowBits := d.br.getBitsFast(n.nbBits)
												d.state = n.newState + lowBits
												return n.symbol
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:373
	// _ = "end of CoverTab[89395]"
}

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:374
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/decompress.go:374
var _ = _go_fuzz_dep_.CoverTab
