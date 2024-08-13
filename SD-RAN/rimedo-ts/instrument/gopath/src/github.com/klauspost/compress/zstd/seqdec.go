// Copyright 2019+ Klaus Post. All rights reserved.
// License information can be found in the LICENSE file.
// Based on work by Yann Collet, released under BSD License.

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:5
package zstd

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:5
)

import (
	"errors"
	"fmt"
	"io"
)

type seq struct {
	litLen		uint32
	matchLen	uint32
	offset		uint32

	// Codes are stored here for the encoder
	// so they only have to be looked up once.
	llCode, mlCode, ofCode	uint8
}

func (s seq) String() string {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:23
	_go_fuzz_dep_.CoverTab[94872]++
												if s.offset <= 3 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:24
		_go_fuzz_dep_.CoverTab[94874]++
													if s.offset == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:25
			_go_fuzz_dep_.CoverTab[94876]++
														return fmt.Sprint("litLen:", s.litLen, ", matchLen:", s.matchLen+zstdMinMatch, ", offset: INVALID (0)")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:26
			// _ = "end of CoverTab[94876]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:27
			_go_fuzz_dep_.CoverTab[94877]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:27
			// _ = "end of CoverTab[94877]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:27
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:27
		// _ = "end of CoverTab[94874]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:27
		_go_fuzz_dep_.CoverTab[94875]++
													return fmt.Sprint("litLen:", s.litLen, ", matchLen:", s.matchLen+zstdMinMatch, ", offset:", s.offset, " (repeat)")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:28
		// _ = "end of CoverTab[94875]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:29
		_go_fuzz_dep_.CoverTab[94878]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:29
		// _ = "end of CoverTab[94878]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:29
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:29
	// _ = "end of CoverTab[94872]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:29
	_go_fuzz_dep_.CoverTab[94873]++
												return fmt.Sprint("litLen:", s.litLen, ", matchLen:", s.matchLen+zstdMinMatch, ", offset:", s.offset-3, " (new)")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:30
	// _ = "end of CoverTab[94873]"
}

type seqCompMode uint8

const (
	compModePredefined	seqCompMode	= iota
	compModeRLE
	compModeFSE
	compModeRepeat
)

type sequenceDec struct {
	// decoder keeps track of the current state and updates it from the bitstream.
	fse	*fseDecoder
	state	fseState
	repeat	bool
}

// init the state of the decoder with input from stream.
func (s *sequenceDec) init(br *bitReader) error {
	if s.fse == nil {
		return errors.New("sequence decoder not defined")
	}
	s.state.init(br, s.fse.actualTableLog, s.fse.dt[:1<<s.fse.actualTableLog])
	return nil
}

// sequenceDecs contains all 3 sequence decoders and their state.
type sequenceDecs struct {
	litLengths	sequenceDec
	offsets		sequenceDec
	matchLengths	sequenceDec
	prevOffset	[3]int
	hist		[]byte
	dict		[]byte
	literals	[]byte
	out		[]byte
	windowSize	int
	maxBits		uint8
}

// initialize all 3 decoders from the stream input.
func (s *sequenceDecs) initialize(br *bitReader, hist *history, literals, out []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:73
	_go_fuzz_dep_.CoverTab[94879]++
												if err := s.litLengths.init(br); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:74
		_go_fuzz_dep_.CoverTab[94884]++
													return errors.New("litLengths:" + err.Error())
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:75
		// _ = "end of CoverTab[94884]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:76
		_go_fuzz_dep_.CoverTab[94885]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:76
		// _ = "end of CoverTab[94885]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:76
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:76
	// _ = "end of CoverTab[94879]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:76
	_go_fuzz_dep_.CoverTab[94880]++
												if err := s.offsets.init(br); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:77
		_go_fuzz_dep_.CoverTab[94886]++
													return errors.New("offsets:" + err.Error())
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:78
		// _ = "end of CoverTab[94886]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:79
		_go_fuzz_dep_.CoverTab[94887]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:79
		// _ = "end of CoverTab[94887]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:79
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:79
	// _ = "end of CoverTab[94880]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:79
	_go_fuzz_dep_.CoverTab[94881]++
												if err := s.matchLengths.init(br); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:80
		_go_fuzz_dep_.CoverTab[94888]++
													return errors.New("matchLengths:" + err.Error())
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:81
		// _ = "end of CoverTab[94888]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:82
		_go_fuzz_dep_.CoverTab[94889]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:82
		// _ = "end of CoverTab[94889]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:82
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:82
	// _ = "end of CoverTab[94881]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:82
	_go_fuzz_dep_.CoverTab[94882]++
												s.literals = literals
												s.hist = hist.b
												s.prevOffset = hist.recentOffsets
												s.maxBits = s.litLengths.fse.maxBits + s.offsets.fse.maxBits + s.matchLengths.fse.maxBits
												s.windowSize = hist.windowSize
												s.out = out
												s.dict = nil
												if hist.dict != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:90
		_go_fuzz_dep_.CoverTab[94890]++
													s.dict = hist.dict.content
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:91
		// _ = "end of CoverTab[94890]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:92
		_go_fuzz_dep_.CoverTab[94891]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:92
		// _ = "end of CoverTab[94891]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:92
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:92
	// _ = "end of CoverTab[94882]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:92
	_go_fuzz_dep_.CoverTab[94883]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:93
	// _ = "end of CoverTab[94883]"
}

// decode sequences from the stream with the provided history.
func (s *sequenceDecs) decode(seqs int, br *bitReader, hist []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:97
	_go_fuzz_dep_.CoverTab[94892]++
												startSize := len(s.out)

												llTable, mlTable, ofTable := s.litLengths.fse.dt[:maxTablesize], s.matchLengths.fse.dt[:maxTablesize], s.offsets.fse.dt[:maxTablesize]
												llState, mlState, ofState := s.litLengths.state.state, s.matchLengths.state.state, s.offsets.state.state

												for i := seqs - 1; i >= 0; i-- {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:103
		_go_fuzz_dep_.CoverTab[94894]++
													if br.overread() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:104
			_go_fuzz_dep_.CoverTab[94907]++
														printf("reading sequence %d, exceeded available data\n", seqs-i)
														return io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:106
			// _ = "end of CoverTab[94907]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:107
			_go_fuzz_dep_.CoverTab[94908]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:107
			// _ = "end of CoverTab[94908]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:107
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:107
		// _ = "end of CoverTab[94894]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:107
		_go_fuzz_dep_.CoverTab[94895]++
													var ll, mo, ml int
													if br.off > 4+((maxOffsetBits+16+16)>>3) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:109
			_go_fuzz_dep_.CoverTab[94909]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:113
			// Final will not read from stream.
														var llB, mlB, moB uint8
														ll, llB = llState.final()
														ml, mlB = mlState.final()
														mo, moB = ofState.final()

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:120
			br.fillFast()
			mo += br.getBits(moB)
			if s.maxBits > 32 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:122
				_go_fuzz_dep_.CoverTab[94912]++
															br.fillFast()
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:123
				// _ = "end of CoverTab[94912]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:124
				_go_fuzz_dep_.CoverTab[94913]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:124
				// _ = "end of CoverTab[94913]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:124
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:124
			// _ = "end of CoverTab[94909]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:124
			_go_fuzz_dep_.CoverTab[94910]++
														ml += br.getBits(mlB)
														ll += br.getBits(llB)

														if moB > 1 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:128
				_go_fuzz_dep_.CoverTab[94914]++
															s.prevOffset[2] = s.prevOffset[1]
															s.prevOffset[1] = s.prevOffset[0]
															s.prevOffset[0] = mo
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:131
				// _ = "end of CoverTab[94914]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:132
				_go_fuzz_dep_.CoverTab[94915]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:135
				if ll == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:135
					_go_fuzz_dep_.CoverTab[94917]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:139
					mo++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:139
					// _ = "end of CoverTab[94917]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:140
					_go_fuzz_dep_.CoverTab[94918]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:140
					// _ = "end of CoverTab[94918]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:140
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:140
				// _ = "end of CoverTab[94915]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:140
				_go_fuzz_dep_.CoverTab[94916]++

															if mo == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:142
					_go_fuzz_dep_.CoverTab[94919]++
																mo = s.prevOffset[0]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:143
					// _ = "end of CoverTab[94919]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:144
					_go_fuzz_dep_.CoverTab[94920]++
																var temp int
																if mo == 3 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:146
						_go_fuzz_dep_.CoverTab[94924]++
																	temp = s.prevOffset[0] - 1
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:147
						// _ = "end of CoverTab[94924]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:148
						_go_fuzz_dep_.CoverTab[94925]++
																	temp = s.prevOffset[mo]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:149
						// _ = "end of CoverTab[94925]"
					}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:150
					// _ = "end of CoverTab[94920]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:150
					_go_fuzz_dep_.CoverTab[94921]++

																if temp == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:152
						_go_fuzz_dep_.CoverTab[94926]++

																	println("temp was 0")
																	temp = 1
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:155
						// _ = "end of CoverTab[94926]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:156
						_go_fuzz_dep_.CoverTab[94927]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:156
						// _ = "end of CoverTab[94927]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:156
					}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:156
					// _ = "end of CoverTab[94921]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:156
					_go_fuzz_dep_.CoverTab[94922]++

																if mo != 1 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:158
						_go_fuzz_dep_.CoverTab[94928]++
																	s.prevOffset[2] = s.prevOffset[1]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:159
						// _ = "end of CoverTab[94928]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:160
						_go_fuzz_dep_.CoverTab[94929]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:160
						// _ = "end of CoverTab[94929]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:160
					}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:160
					// _ = "end of CoverTab[94922]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:160
					_go_fuzz_dep_.CoverTab[94923]++
																s.prevOffset[1] = s.prevOffset[0]
																s.prevOffset[0] = temp
																mo = temp
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:163
					// _ = "end of CoverTab[94923]"
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:164
				// _ = "end of CoverTab[94916]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:165
			// _ = "end of CoverTab[94910]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:165
			_go_fuzz_dep_.CoverTab[94911]++
														br.fillFast()
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:166
			// _ = "end of CoverTab[94911]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:167
			_go_fuzz_dep_.CoverTab[94930]++
														ll, mo, ml = s.next(br, llState, mlState, ofState)
														br.fill()
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:169
			// _ = "end of CoverTab[94930]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:170
		// _ = "end of CoverTab[94895]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:170
		_go_fuzz_dep_.CoverTab[94896]++

													if debugSequences {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:172
			_go_fuzz_dep_.CoverTab[94931]++
														println("Seq", seqs-i-1, "Litlen:", ll, "mo:", mo, "(abs) ml:", ml)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:173
			// _ = "end of CoverTab[94931]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:174
			_go_fuzz_dep_.CoverTab[94932]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:174
			// _ = "end of CoverTab[94932]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:174
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:174
		// _ = "end of CoverTab[94896]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:174
		_go_fuzz_dep_.CoverTab[94897]++

													if ll > len(s.literals) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:176
			_go_fuzz_dep_.CoverTab[94933]++
														return fmt.Errorf("unexpected literal count, want %d bytes, but only %d is available", ll, len(s.literals))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:177
			// _ = "end of CoverTab[94933]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:178
			_go_fuzz_dep_.CoverTab[94934]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:178
			// _ = "end of CoverTab[94934]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:178
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:178
		// _ = "end of CoverTab[94897]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:178
		_go_fuzz_dep_.CoverTab[94898]++
													size := ll + ml + len(s.out)
													if size-startSize > maxBlockSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:180
			_go_fuzz_dep_.CoverTab[94935]++
														return fmt.Errorf("output (%d) bigger than max block size", size)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:181
			// _ = "end of CoverTab[94935]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:182
			_go_fuzz_dep_.CoverTab[94936]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:182
			// _ = "end of CoverTab[94936]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:182
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:182
		// _ = "end of CoverTab[94898]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:182
		_go_fuzz_dep_.CoverTab[94899]++
													if size > cap(s.out) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:183
			_go_fuzz_dep_.CoverTab[94937]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:188
			used := len(s.out) - startSize
			addBytes := 256 + ll + ml + used>>2

			if used+addBytes > maxBlockSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:191
				_go_fuzz_dep_.CoverTab[94939]++
															addBytes = maxBlockSize - used
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:192
				// _ = "end of CoverTab[94939]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:193
				_go_fuzz_dep_.CoverTab[94940]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:193
				// _ = "end of CoverTab[94940]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:193
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:193
			// _ = "end of CoverTab[94937]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:193
			_go_fuzz_dep_.CoverTab[94938]++
														s.out = append(s.out, make([]byte, addBytes)...)
														s.out = s.out[:len(s.out)-addBytes]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:195
			// _ = "end of CoverTab[94938]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:196
			_go_fuzz_dep_.CoverTab[94941]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:196
			// _ = "end of CoverTab[94941]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:196
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:196
		// _ = "end of CoverTab[94899]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:196
		_go_fuzz_dep_.CoverTab[94900]++
													if ml > maxMatchLen {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:197
			_go_fuzz_dep_.CoverTab[94942]++
														return fmt.Errorf("match len (%d) bigger than max allowed length", ml)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:198
			// _ = "end of CoverTab[94942]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:199
			_go_fuzz_dep_.CoverTab[94943]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:199
			// _ = "end of CoverTab[94943]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:199
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:199
		// _ = "end of CoverTab[94900]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:199
		_go_fuzz_dep_.CoverTab[94901]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:202
		s.out = append(s.out, s.literals[:ll]...)
		s.literals = s.literals[ll:]
		out := s.out

		if mo == 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:206
			_go_fuzz_dep_.CoverTab[94944]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:206
			return ml > 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:206
			// _ = "end of CoverTab[94944]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:206
		}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:206
			_go_fuzz_dep_.CoverTab[94945]++
														return fmt.Errorf("zero matchoff and matchlen (%d) > 0", ml)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:207
			// _ = "end of CoverTab[94945]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:208
			_go_fuzz_dep_.CoverTab[94946]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:208
			// _ = "end of CoverTab[94946]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:208
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:208
		// _ = "end of CoverTab[94901]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:208
		_go_fuzz_dep_.CoverTab[94902]++

													if mo > len(s.out)+len(hist) || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:210
			_go_fuzz_dep_.CoverTab[94947]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:210
			return mo > s.windowSize
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:210
			// _ = "end of CoverTab[94947]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:210
		}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:210
			_go_fuzz_dep_.CoverTab[94948]++
														if len(s.dict) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:211
				_go_fuzz_dep_.CoverTab[94951]++
															return fmt.Errorf("match offset (%d) bigger than current history (%d)", mo, len(s.out)+len(hist))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:212
				// _ = "end of CoverTab[94951]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:213
				_go_fuzz_dep_.CoverTab[94952]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:213
				// _ = "end of CoverTab[94952]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:213
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:213
			// _ = "end of CoverTab[94948]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:213
			_go_fuzz_dep_.CoverTab[94949]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:216
			dictO := len(s.dict) - (mo - (len(s.out) + len(hist)))
			if dictO < 0 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:217
				_go_fuzz_dep_.CoverTab[94953]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:217
				return dictO >= len(s.dict)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:217
				// _ = "end of CoverTab[94953]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:217
			}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:217
				_go_fuzz_dep_.CoverTab[94954]++
															return fmt.Errorf("match offset (%d) bigger than current history (%d)", mo, len(s.out)+len(hist))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:218
				// _ = "end of CoverTab[94954]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:219
				_go_fuzz_dep_.CoverTab[94955]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:219
				// _ = "end of CoverTab[94955]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:219
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:219
			// _ = "end of CoverTab[94949]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:219
			_go_fuzz_dep_.CoverTab[94950]++
														end := dictO + ml
														if end > len(s.dict) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:221
				_go_fuzz_dep_.CoverTab[94956]++
															out = append(out, s.dict[dictO:]...)
															mo -= len(s.dict) - dictO
															ml -= len(s.dict) - dictO
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:224
				// _ = "end of CoverTab[94956]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:225
				_go_fuzz_dep_.CoverTab[94957]++
															out = append(out, s.dict[dictO:end]...)
															mo = 0
															ml = 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:228
				// _ = "end of CoverTab[94957]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:229
			// _ = "end of CoverTab[94950]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:230
			_go_fuzz_dep_.CoverTab[94958]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:230
			// _ = "end of CoverTab[94958]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:230
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:230
		// _ = "end of CoverTab[94902]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:230
		_go_fuzz_dep_.CoverTab[94903]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:234
		if v := mo - len(s.out); v > 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:234
			_go_fuzz_dep_.CoverTab[94959]++

														start := len(s.hist) - v
														if ml > v {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:237
				_go_fuzz_dep_.CoverTab[94960]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:240
				out = append(out, s.hist[start:]...)
															mo -= v
															ml -= v
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:242
				// _ = "end of CoverTab[94960]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:243
				_go_fuzz_dep_.CoverTab[94961]++
															out = append(out, s.hist[start:start+ml]...)
															ml = 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:245
				// _ = "end of CoverTab[94961]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:246
			// _ = "end of CoverTab[94959]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:247
			_go_fuzz_dep_.CoverTab[94962]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:247
			// _ = "end of CoverTab[94962]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:247
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:247
		// _ = "end of CoverTab[94903]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:247
		_go_fuzz_dep_.CoverTab[94904]++

													if ml > 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:249
			_go_fuzz_dep_.CoverTab[94963]++
														start := len(s.out) - mo
														if ml <= len(s.out)-start {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:251
				_go_fuzz_dep_.CoverTab[94964]++

															out = append(out, s.out[start:start+ml]...)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:253
				// _ = "end of CoverTab[94964]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:254
				_go_fuzz_dep_.CoverTab[94965]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:257
				out = out[:len(out)+ml]
				src := out[start : start+ml]

				dst := out[len(out)-ml:]
				dst = dst[:len(src)]
				for i := range src {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:262
					_go_fuzz_dep_.CoverTab[94966]++
																dst[i] = src[i]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:263
					// _ = "end of CoverTab[94966]"
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:264
				// _ = "end of CoverTab[94965]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:265
			// _ = "end of CoverTab[94963]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:266
			_go_fuzz_dep_.CoverTab[94967]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:266
			// _ = "end of CoverTab[94967]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:266
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:266
		// _ = "end of CoverTab[94904]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:266
		_go_fuzz_dep_.CoverTab[94905]++
													s.out = out
													if i == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:268
			_go_fuzz_dep_.CoverTab[94968]++

														break
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:270
			// _ = "end of CoverTab[94968]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:271
			_go_fuzz_dep_.CoverTab[94969]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:271
			// _ = "end of CoverTab[94969]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:271
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:271
		// _ = "end of CoverTab[94905]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:271
		_go_fuzz_dep_.CoverTab[94906]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:275
		nBits := llState.nbBits() + mlState.nbBits() + ofState.nbBits()
		if nBits == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:276
			_go_fuzz_dep_.CoverTab[94970]++
														llState = llTable[llState.newState()&maxTableMask]
														mlState = mlTable[mlState.newState()&maxTableMask]
														ofState = ofTable[ofState.newState()&maxTableMask]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:279
			// _ = "end of CoverTab[94970]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:280
			_go_fuzz_dep_.CoverTab[94971]++
														bits := br.get32BitsFast(nBits)
														lowBits := uint16(bits >> ((ofState.nbBits() + mlState.nbBits()) & 31))
														llState = llTable[(llState.newState()+lowBits)&maxTableMask]

														lowBits = uint16(bits >> (ofState.nbBits() & 31))
														lowBits &= bitMask[mlState.nbBits()&15]
														mlState = mlTable[(mlState.newState()+lowBits)&maxTableMask]

														lowBits = uint16(bits) & bitMask[ofState.nbBits()&15]
														ofState = ofTable[(ofState.newState()+lowBits)&maxTableMask]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:290
			// _ = "end of CoverTab[94971]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:291
		// _ = "end of CoverTab[94906]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:292
	// _ = "end of CoverTab[94892]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:292
	_go_fuzz_dep_.CoverTab[94893]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:295
	s.out = append(s.out, s.literals...)
												return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:296
	// _ = "end of CoverTab[94893]"
}

// update states, at least 27 bits must be available.
func (s *sequenceDecs) update(br *bitReader) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:300
	_go_fuzz_dep_.CoverTab[94972]++

												s.litLengths.state.next(br)

												s.matchLengths.state.next(br)

												s.offsets.state.next(br)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:306
	// _ = "end of CoverTab[94972]"
}

var bitMask [16]uint16

func init() {
	for i := range bitMask[:] {
		bitMask[i] = uint16((1 << uint(i)) - 1)
	}
}

// update states, at least 27 bits must be available.
func (s *sequenceDecs) updateAlt(br *bitReader) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:318
	_go_fuzz_dep_.CoverTab[94973]++

												a, b, c := s.litLengths.state.state, s.matchLengths.state.state, s.offsets.state.state

												nBits := a.nbBits() + b.nbBits() + c.nbBits()
												if nBits == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:323
		_go_fuzz_dep_.CoverTab[94975]++
													s.litLengths.state.state = s.litLengths.state.dt[a.newState()]
													s.matchLengths.state.state = s.matchLengths.state.dt[b.newState()]
													s.offsets.state.state = s.offsets.state.dt[c.newState()]
													return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:327
		// _ = "end of CoverTab[94975]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:328
		_go_fuzz_dep_.CoverTab[94976]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:328
		// _ = "end of CoverTab[94976]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:328
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:328
	// _ = "end of CoverTab[94973]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:328
	_go_fuzz_dep_.CoverTab[94974]++
												bits := br.get32BitsFast(nBits)
												lowBits := uint16(bits >> ((c.nbBits() + b.nbBits()) & 31))
												s.litLengths.state.state = s.litLengths.state.dt[a.newState()+lowBits]

												lowBits = uint16(bits >> (c.nbBits() & 31))
												lowBits &= bitMask[b.nbBits()&15]
												s.matchLengths.state.state = s.matchLengths.state.dt[b.newState()+lowBits]

												lowBits = uint16(bits) & bitMask[c.nbBits()&15]
												s.offsets.state.state = s.offsets.state.dt[c.newState()+lowBits]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:338
	// _ = "end of CoverTab[94974]"
}

// nextFast will return new states when there are at least 4 unused bytes left on the stream when done.
func (s *sequenceDecs) nextFast(br *bitReader, llState, mlState, ofState decSymbol) (ll, mo, ml int) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:342
	_go_fuzz_dep_.CoverTab[94977]++

												ll, llB := llState.final()
												ml, mlB := mlState.final()
												mo, moB := ofState.final()

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:349
	br.fillFast()
	mo += br.getBits(moB)
	if s.maxBits > 32 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:351
		_go_fuzz_dep_.CoverTab[94985]++
													br.fillFast()
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:352
		// _ = "end of CoverTab[94985]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:353
		_go_fuzz_dep_.CoverTab[94986]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:353
		// _ = "end of CoverTab[94986]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:353
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:353
	// _ = "end of CoverTab[94977]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:353
	_go_fuzz_dep_.CoverTab[94978]++
												ml += br.getBits(mlB)
												ll += br.getBits(llB)

												if moB > 1 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:357
		_go_fuzz_dep_.CoverTab[94987]++
													s.prevOffset[2] = s.prevOffset[1]
													s.prevOffset[1] = s.prevOffset[0]
													s.prevOffset[0] = mo
													return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:361
		// _ = "end of CoverTab[94987]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:362
		_go_fuzz_dep_.CoverTab[94988]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:362
		// _ = "end of CoverTab[94988]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:362
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:362
	// _ = "end of CoverTab[94978]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:362
	_go_fuzz_dep_.CoverTab[94979]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:365
	if ll == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:365
		_go_fuzz_dep_.CoverTab[94989]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:369
		mo++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:369
		// _ = "end of CoverTab[94989]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:370
		_go_fuzz_dep_.CoverTab[94990]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:370
		// _ = "end of CoverTab[94990]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:370
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:370
	// _ = "end of CoverTab[94979]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:370
	_go_fuzz_dep_.CoverTab[94980]++

												if mo == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:372
		_go_fuzz_dep_.CoverTab[94991]++
													mo = s.prevOffset[0]
													return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:374
		// _ = "end of CoverTab[94991]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:375
		_go_fuzz_dep_.CoverTab[94992]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:375
		// _ = "end of CoverTab[94992]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:375
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:375
	// _ = "end of CoverTab[94980]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:375
	_go_fuzz_dep_.CoverTab[94981]++
												var temp int
												if mo == 3 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:377
		_go_fuzz_dep_.CoverTab[94993]++
													temp = s.prevOffset[0] - 1
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:378
		// _ = "end of CoverTab[94993]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:379
		_go_fuzz_dep_.CoverTab[94994]++
													temp = s.prevOffset[mo]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:380
		// _ = "end of CoverTab[94994]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:381
	// _ = "end of CoverTab[94981]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:381
	_go_fuzz_dep_.CoverTab[94982]++

												if temp == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:383
		_go_fuzz_dep_.CoverTab[94995]++

													println("temp was 0")
													temp = 1
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:386
		// _ = "end of CoverTab[94995]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:387
		_go_fuzz_dep_.CoverTab[94996]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:387
		// _ = "end of CoverTab[94996]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:387
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:387
	// _ = "end of CoverTab[94982]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:387
	_go_fuzz_dep_.CoverTab[94983]++

												if mo != 1 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:389
		_go_fuzz_dep_.CoverTab[94997]++
													s.prevOffset[2] = s.prevOffset[1]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:390
		// _ = "end of CoverTab[94997]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:391
		_go_fuzz_dep_.CoverTab[94998]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:391
		// _ = "end of CoverTab[94998]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:391
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:391
	// _ = "end of CoverTab[94983]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:391
	_go_fuzz_dep_.CoverTab[94984]++
												s.prevOffset[1] = s.prevOffset[0]
												s.prevOffset[0] = temp
												mo = temp
												return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:395
	// _ = "end of CoverTab[94984]"
}

func (s *sequenceDecs) next(br *bitReader, llState, mlState, ofState decSymbol) (ll, mo, ml int) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:398
	_go_fuzz_dep_.CoverTab[94999]++

												ll, llB := llState.final()
												ml, mlB := mlState.final()
												mo, moB := ofState.final()

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:405
	br.fill()
	if s.maxBits <= 32 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:406
		_go_fuzz_dep_.CoverTab[95001]++
													mo += br.getBits(moB)
													ml += br.getBits(mlB)
													ll += br.getBits(llB)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:409
		// _ = "end of CoverTab[95001]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:410
		_go_fuzz_dep_.CoverTab[95002]++
													mo += br.getBits(moB)
													br.fill()

													ml += br.getBits(mlB)
													ll += br.getBits(llB)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:415
		// _ = "end of CoverTab[95002]"

	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:417
	// _ = "end of CoverTab[94999]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:417
	_go_fuzz_dep_.CoverTab[95000]++
												mo = s.adjustOffset(mo, ll, moB)
												return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:419
	// _ = "end of CoverTab[95000]"
}

func (s *sequenceDecs) adjustOffset(offset, litLen int, offsetB uint8) int {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:422
	_go_fuzz_dep_.CoverTab[95003]++
												if offsetB > 1 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:423
		_go_fuzz_dep_.CoverTab[95010]++
													s.prevOffset[2] = s.prevOffset[1]
													s.prevOffset[1] = s.prevOffset[0]
													s.prevOffset[0] = offset
													return offset
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:427
		// _ = "end of CoverTab[95010]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:428
		_go_fuzz_dep_.CoverTab[95011]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:428
		// _ = "end of CoverTab[95011]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:428
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:428
	// _ = "end of CoverTab[95003]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:428
	_go_fuzz_dep_.CoverTab[95004]++

												if litLen == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:430
		_go_fuzz_dep_.CoverTab[95012]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:434
		offset++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:434
		// _ = "end of CoverTab[95012]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:435
		_go_fuzz_dep_.CoverTab[95013]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:435
		// _ = "end of CoverTab[95013]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:435
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:435
	// _ = "end of CoverTab[95004]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:435
	_go_fuzz_dep_.CoverTab[95005]++

												if offset == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:437
		_go_fuzz_dep_.CoverTab[95014]++
													return s.prevOffset[0]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:438
		// _ = "end of CoverTab[95014]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:439
		_go_fuzz_dep_.CoverTab[95015]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:439
		// _ = "end of CoverTab[95015]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:439
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:439
	// _ = "end of CoverTab[95005]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:439
	_go_fuzz_dep_.CoverTab[95006]++
												var temp int
												if offset == 3 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:441
		_go_fuzz_dep_.CoverTab[95016]++
													temp = s.prevOffset[0] - 1
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:442
		// _ = "end of CoverTab[95016]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:443
		_go_fuzz_dep_.CoverTab[95017]++
													temp = s.prevOffset[offset]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:444
		// _ = "end of CoverTab[95017]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:445
	// _ = "end of CoverTab[95006]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:445
	_go_fuzz_dep_.CoverTab[95007]++

												if temp == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:447
		_go_fuzz_dep_.CoverTab[95018]++

													println("temp was 0")
													temp = 1
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:450
		// _ = "end of CoverTab[95018]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:451
		_go_fuzz_dep_.CoverTab[95019]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:451
		// _ = "end of CoverTab[95019]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:451
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:451
	// _ = "end of CoverTab[95007]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:451
	_go_fuzz_dep_.CoverTab[95008]++

												if offset != 1 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:453
		_go_fuzz_dep_.CoverTab[95020]++
													s.prevOffset[2] = s.prevOffset[1]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:454
		// _ = "end of CoverTab[95020]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:455
		_go_fuzz_dep_.CoverTab[95021]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:455
		// _ = "end of CoverTab[95021]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:455
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:455
	// _ = "end of CoverTab[95008]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:455
	_go_fuzz_dep_.CoverTab[95009]++
												s.prevOffset[1] = s.prevOffset[0]
												s.prevOffset[0] = temp
												return temp
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:458
	// _ = "end of CoverTab[95009]"
}

// mergeHistory will merge history.
func (s *sequenceDecs) mergeHistory(hist *sequenceDecs) (*sequenceDecs, error) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:462
	_go_fuzz_dep_.CoverTab[95022]++
												for i := uint(0); i < 3; i++ {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:463
		_go_fuzz_dep_.CoverTab[95024]++
													var sNew, sHist *sequenceDec
													switch i {
		default:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:466
			_go_fuzz_dep_.CoverTab[95029]++

														sNew = &s.litLengths
														sHist = &hist.litLengths
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:469
			// _ = "end of CoverTab[95029]"
		case 1:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:470
			_go_fuzz_dep_.CoverTab[95030]++
														sNew = &s.offsets
														sHist = &hist.offsets
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:472
			// _ = "end of CoverTab[95030]"
		case 2:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:473
			_go_fuzz_dep_.CoverTab[95031]++
														sNew = &s.matchLengths
														sHist = &hist.matchLengths
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:475
			// _ = "end of CoverTab[95031]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:476
		// _ = "end of CoverTab[95024]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:476
		_go_fuzz_dep_.CoverTab[95025]++
													if sNew.repeat {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:477
			_go_fuzz_dep_.CoverTab[95032]++
														if sHist.fse == nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:478
				_go_fuzz_dep_.CoverTab[95034]++
															return nil, fmt.Errorf("sequence stream %d, repeat requested, but no history", i)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:479
				// _ = "end of CoverTab[95034]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:480
				_go_fuzz_dep_.CoverTab[95035]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:480
				// _ = "end of CoverTab[95035]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:480
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:480
			// _ = "end of CoverTab[95032]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:480
			_go_fuzz_dep_.CoverTab[95033]++
														continue
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:481
			// _ = "end of CoverTab[95033]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:482
			_go_fuzz_dep_.CoverTab[95036]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:482
			// _ = "end of CoverTab[95036]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:482
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:482
		// _ = "end of CoverTab[95025]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:482
		_go_fuzz_dep_.CoverTab[95026]++
													if sNew.fse == nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:483
			_go_fuzz_dep_.CoverTab[95037]++
														return nil, fmt.Errorf("sequence stream %d, no fse found", i)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:484
			// _ = "end of CoverTab[95037]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:485
			_go_fuzz_dep_.CoverTab[95038]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:485
			// _ = "end of CoverTab[95038]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:485
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:485
		// _ = "end of CoverTab[95026]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:485
		_go_fuzz_dep_.CoverTab[95027]++
													if sHist.fse != nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:486
			_go_fuzz_dep_.CoverTab[95039]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:486
			return !sHist.fse.preDefined
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:486
			// _ = "end of CoverTab[95039]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:486
		}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:486
			_go_fuzz_dep_.CoverTab[95040]++
														fseDecoderPool.Put(sHist.fse)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:487
			// _ = "end of CoverTab[95040]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:488
			_go_fuzz_dep_.CoverTab[95041]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:488
			// _ = "end of CoverTab[95041]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:488
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:488
		// _ = "end of CoverTab[95027]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:488
		_go_fuzz_dep_.CoverTab[95028]++
													sHist.fse = sNew.fse
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:489
		// _ = "end of CoverTab[95028]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:490
	// _ = "end of CoverTab[95022]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:490
	_go_fuzz_dep_.CoverTab[95023]++
												return hist, nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:491
	// _ = "end of CoverTab[95023]"
}

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:492
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/seqdec.go:492
var _ = _go_fuzz_dep_.CoverTab
