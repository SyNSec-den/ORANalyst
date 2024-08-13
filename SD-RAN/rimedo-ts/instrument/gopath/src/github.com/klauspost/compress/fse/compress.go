// Copyright 2018 Klaus Post. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Based on work Copyright (c) 2013, Yann Collet, released under BSD License.

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:6
package fse

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:6
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:6
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:6
)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:6
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:6
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:6
)

import (
	"errors"
	"fmt"
)

// Compress the input bytes. Input must be < 2GB.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:13
// Provide a Scratch buffer to avoid memory allocations.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:13
// Note that the output is also kept in the scratch buffer.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:13
// If input is too hard to compress, ErrIncompressible is returned.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:13
// If input is a single byte value repeated ErrUseRLE is returned.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:18
func Compress(in []byte, s *Scratch) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:18
	_go_fuzz_dep_.CoverTab[89013]++
												if len(in) <= 1 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:19
		_go_fuzz_dep_.CoverTab[89026]++
													return nil, ErrIncompressible
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:20
		// _ = "end of CoverTab[89026]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:21
		_go_fuzz_dep_.CoverTab[89027]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:21
		// _ = "end of CoverTab[89027]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:21
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:21
	// _ = "end of CoverTab[89013]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:21
	_go_fuzz_dep_.CoverTab[89014]++
												if len(in) > (2<<30)-1 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:22
		_go_fuzz_dep_.CoverTab[89028]++
													return nil, errors.New("input too big, must be < 2GB")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:23
		// _ = "end of CoverTab[89028]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:24
		_go_fuzz_dep_.CoverTab[89029]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:24
		// _ = "end of CoverTab[89029]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:24
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:24
	// _ = "end of CoverTab[89014]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:24
	_go_fuzz_dep_.CoverTab[89015]++
												s, err := s.prepare(in)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:26
		_go_fuzz_dep_.CoverTab[89030]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:27
		// _ = "end of CoverTab[89030]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:28
		_go_fuzz_dep_.CoverTab[89031]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:28
		// _ = "end of CoverTab[89031]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:28
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:28
	// _ = "end of CoverTab[89015]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:28
	_go_fuzz_dep_.CoverTab[89016]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:31
	maxCount := s.maxCount
	if maxCount == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:32
		_go_fuzz_dep_.CoverTab[89032]++
													maxCount = s.countSimple(in)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:33
		// _ = "end of CoverTab[89032]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:34
		_go_fuzz_dep_.CoverTab[89033]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:34
		// _ = "end of CoverTab[89033]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:34
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:34
	// _ = "end of CoverTab[89016]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:34
	_go_fuzz_dep_.CoverTab[89017]++

												s.clearCount = true
												s.maxCount = 0
												if maxCount == len(in) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:38
		_go_fuzz_dep_.CoverTab[89034]++

													return nil, ErrUseRLE
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:40
		// _ = "end of CoverTab[89034]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:41
		_go_fuzz_dep_.CoverTab[89035]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:41
		// _ = "end of CoverTab[89035]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:41
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:41
	// _ = "end of CoverTab[89017]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:41
	_go_fuzz_dep_.CoverTab[89018]++
												if maxCount == 1 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:42
		_go_fuzz_dep_.CoverTab[89036]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:42
		return maxCount < (len(in) >> 7)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:42
		// _ = "end of CoverTab[89036]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:42
	}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:42
		_go_fuzz_dep_.CoverTab[89037]++

													return nil, ErrIncompressible
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:44
		// _ = "end of CoverTab[89037]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:45
		_go_fuzz_dep_.CoverTab[89038]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:45
		// _ = "end of CoverTab[89038]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:45
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:45
	// _ = "end of CoverTab[89018]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:45
	_go_fuzz_dep_.CoverTab[89019]++
												s.optimalTableLog()
												err = s.normalizeCount()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:48
		_go_fuzz_dep_.CoverTab[89039]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:49
		// _ = "end of CoverTab[89039]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:50
		_go_fuzz_dep_.CoverTab[89040]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:50
		// _ = "end of CoverTab[89040]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:50
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:50
	// _ = "end of CoverTab[89019]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:50
	_go_fuzz_dep_.CoverTab[89020]++
												err = s.writeCount()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:52
		_go_fuzz_dep_.CoverTab[89041]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:53
		// _ = "end of CoverTab[89041]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:54
		_go_fuzz_dep_.CoverTab[89042]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:54
		// _ = "end of CoverTab[89042]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:54
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:54
	// _ = "end of CoverTab[89020]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:54
	_go_fuzz_dep_.CoverTab[89021]++

												if false {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:56
		_go_fuzz_dep_.CoverTab[89043]++
													err = s.validateNorm()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:58
			_go_fuzz_dep_.CoverTab[89044]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:59
			// _ = "end of CoverTab[89044]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:60
			_go_fuzz_dep_.CoverTab[89045]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:60
			// _ = "end of CoverTab[89045]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:60
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:60
		// _ = "end of CoverTab[89043]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:61
		_go_fuzz_dep_.CoverTab[89046]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:61
		// _ = "end of CoverTab[89046]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:61
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:61
	// _ = "end of CoverTab[89021]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:61
	_go_fuzz_dep_.CoverTab[89022]++

												err = s.buildCTable()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:64
		_go_fuzz_dep_.CoverTab[89047]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:65
		// _ = "end of CoverTab[89047]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:66
		_go_fuzz_dep_.CoverTab[89048]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:66
		// _ = "end of CoverTab[89048]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:66
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:66
	// _ = "end of CoverTab[89022]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:66
	_go_fuzz_dep_.CoverTab[89023]++
												err = s.compress(in)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:68
		_go_fuzz_dep_.CoverTab[89049]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:69
		// _ = "end of CoverTab[89049]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:70
		_go_fuzz_dep_.CoverTab[89050]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:70
		// _ = "end of CoverTab[89050]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:70
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:70
	// _ = "end of CoverTab[89023]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:70
	_go_fuzz_dep_.CoverTab[89024]++
												s.Out = s.bw.out

												if len(s.Out) >= len(in) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:73
		_go_fuzz_dep_.CoverTab[89051]++
													return nil, ErrIncompressible
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:74
		// _ = "end of CoverTab[89051]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:75
		_go_fuzz_dep_.CoverTab[89052]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:75
		// _ = "end of CoverTab[89052]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:75
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:75
	// _ = "end of CoverTab[89024]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:75
	_go_fuzz_dep_.CoverTab[89025]++
												return s.Out, nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:76
	// _ = "end of CoverTab[89025]"
}

// cState contains the compression state of a stream.
type cState struct {
	bw		*bitWriter
	stateTable	[]uint16
	state		uint16
}

// init will initialize the compression state to the first symbol of the stream.
func (c *cState) init(bw *bitWriter, ct *cTable, tableLog uint8, first symbolTransform) {
	c.bw = bw
	c.stateTable = ct.stateTable

	nbBitsOut := (first.deltaNbBits + (1 << 15)) >> 16
	im := int32((nbBitsOut << 16) - first.deltaNbBits)
	lu := (im >> nbBitsOut) + first.deltaFindState
	c.state = c.stateTable[lu]
}

// encode the output symbol provided and write it to the bitstream.
func (c *cState) encode(symbolTT symbolTransform) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:98
	_go_fuzz_dep_.CoverTab[89053]++
												nbBitsOut := (uint32(c.state) + symbolTT.deltaNbBits) >> 16
												dstState := int32(c.state>>(nbBitsOut&15)) + symbolTT.deltaFindState
												c.bw.addBits16NC(c.state, uint8(nbBitsOut))
												c.state = c.stateTable[dstState]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:102
	// _ = "end of CoverTab[89053]"
}

// encode the output symbol provided and write it to the bitstream.
func (c *cState) encodeZero(symbolTT symbolTransform) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:106
	_go_fuzz_dep_.CoverTab[89054]++
												nbBitsOut := (uint32(c.state) + symbolTT.deltaNbBits) >> 16
												dstState := int32(c.state>>(nbBitsOut&15)) + symbolTT.deltaFindState
												c.bw.addBits16ZeroNC(c.state, uint8(nbBitsOut))
												c.state = c.stateTable[dstState]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:110
	// _ = "end of CoverTab[89054]"
}

// flush will write the tablelog to the output and flush the remaining full bytes.
func (c *cState) flush(tableLog uint8) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:114
	_go_fuzz_dep_.CoverTab[89055]++
												c.bw.flush32()
												c.bw.addBits16NC(c.state, tableLog)
												c.bw.flush()
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:117
	// _ = "end of CoverTab[89055]"
}

// compress is the main compression loop that will encode the input from the last byte to the first.
func (s *Scratch) compress(src []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:121
	_go_fuzz_dep_.CoverTab[89056]++
												if len(src) <= 2 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:122
		_go_fuzz_dep_.CoverTab[89061]++
													return errors.New("compress: src too small")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:123
		// _ = "end of CoverTab[89061]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:124
		_go_fuzz_dep_.CoverTab[89062]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:124
		// _ = "end of CoverTab[89062]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:124
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:124
	// _ = "end of CoverTab[89056]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:124
	_go_fuzz_dep_.CoverTab[89057]++
												tt := s.ct.symbolTT[:256]
												s.bw.reset(s.Out)

	// Our two states each encodes every second byte.
												// Last byte encoded (first byte decoded) will always be encoded by c1.
												var c1, c2 cState

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:133
	ip := len(src)
	if ip&1 == 1 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:134
		_go_fuzz_dep_.CoverTab[89063]++
													c1.init(&s.bw, &s.ct, s.actualTableLog, tt[src[ip-1]])
													c2.init(&s.bw, &s.ct, s.actualTableLog, tt[src[ip-2]])
													c1.encodeZero(tt[src[ip-3]])
													ip -= 3
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:138
		// _ = "end of CoverTab[89063]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:139
		_go_fuzz_dep_.CoverTab[89064]++
													c2.init(&s.bw, &s.ct, s.actualTableLog, tt[src[ip-1]])
													c1.init(&s.bw, &s.ct, s.actualTableLog, tt[src[ip-2]])
													ip -= 2
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:142
		// _ = "end of CoverTab[89064]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:143
	// _ = "end of CoverTab[89057]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:143
	_go_fuzz_dep_.CoverTab[89058]++
												if ip&2 != 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:144
		_go_fuzz_dep_.CoverTab[89065]++
													c2.encodeZero(tt[src[ip-1]])
													c1.encodeZero(tt[src[ip-2]])
													ip -= 2
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:147
		// _ = "end of CoverTab[89065]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:148
		_go_fuzz_dep_.CoverTab[89066]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:148
		// _ = "end of CoverTab[89066]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:148
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:148
	// _ = "end of CoverTab[89058]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:148
	_go_fuzz_dep_.CoverTab[89059]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:151
	switch {
	case !s.zeroBits && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:152
		_go_fuzz_dep_.CoverTab[89071]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:152
		return s.actualTableLog <= 8
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:152
		// _ = "end of CoverTab[89071]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:152
	}():
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:152
		_go_fuzz_dep_.CoverTab[89067]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:155
		for ip >= 4 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:155
			_go_fuzz_dep_.CoverTab[89072]++
														s.bw.flush32()
														v3, v2, v1, v0 := src[ip-4], src[ip-3], src[ip-2], src[ip-1]
														c2.encode(tt[v0])
														c1.encode(tt[v1])
														c2.encode(tt[v2])
														c1.encode(tt[v3])
														ip -= 4
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:162
			// _ = "end of CoverTab[89072]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:163
		// _ = "end of CoverTab[89067]"
	case !s.zeroBits:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:164
		_go_fuzz_dep_.CoverTab[89068]++

													for ip >= 4 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:166
			_go_fuzz_dep_.CoverTab[89073]++
														s.bw.flush32()
														v3, v2, v1, v0 := src[ip-4], src[ip-3], src[ip-2], src[ip-1]
														c2.encode(tt[v0])
														c1.encode(tt[v1])
														s.bw.flush32()
														c2.encode(tt[v2])
														c1.encode(tt[v3])
														ip -= 4
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:174
			// _ = "end of CoverTab[89073]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:175
		// _ = "end of CoverTab[89068]"
	case s.actualTableLog <= 8:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:176
		_go_fuzz_dep_.CoverTab[89069]++

													for ip >= 4 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:178
			_go_fuzz_dep_.CoverTab[89074]++
														s.bw.flush32()
														v3, v2, v1, v0 := src[ip-4], src[ip-3], src[ip-2], src[ip-1]
														c2.encodeZero(tt[v0])
														c1.encodeZero(tt[v1])
														c2.encodeZero(tt[v2])
														c1.encodeZero(tt[v3])
														ip -= 4
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:185
			// _ = "end of CoverTab[89074]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:186
		// _ = "end of CoverTab[89069]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:187
		_go_fuzz_dep_.CoverTab[89070]++
													for ip >= 4 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:188
			_go_fuzz_dep_.CoverTab[89075]++
														s.bw.flush32()
														v3, v2, v1, v0 := src[ip-4], src[ip-3], src[ip-2], src[ip-1]
														c2.encodeZero(tt[v0])
														c1.encodeZero(tt[v1])
														s.bw.flush32()
														c2.encodeZero(tt[v2])
														c1.encodeZero(tt[v3])
														ip -= 4
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:196
			// _ = "end of CoverTab[89075]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:197
		// _ = "end of CoverTab[89070]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:198
	// _ = "end of CoverTab[89059]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:198
	_go_fuzz_dep_.CoverTab[89060]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:202
	c2.flush(s.actualTableLog)
												c1.flush(s.actualTableLog)

												return s.bw.close()
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:205
	// _ = "end of CoverTab[89060]"
}

// writeCount will write the normalized histogram count to header.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:208
// This is read back by readNCount.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:210
func (s *Scratch) writeCount() error {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:210
	_go_fuzz_dep_.CoverTab[89076]++
												var (
		tableLog	= s.actualTableLog
		tableSize	= 1 << tableLog
		previous0	bool
		charnum		uint16

		maxHeaderSize	= ((int(s.symbolLen) * int(tableLog)) >> 3) + 3

		// Write Table Size
		bitStream	= uint32(tableLog - minTablelog)
		bitCount	= uint(4)
		remaining	= int16(tableSize + 1)	/* +1 for extra accuracy */
		threshold	= int16(tableSize)
		nbBits		= uint(tableLog + 1)
	)
	if cap(s.Out) < maxHeaderSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:226
		_go_fuzz_dep_.CoverTab[89080]++
													s.Out = make([]byte, 0, s.br.remain()+maxHeaderSize)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:227
		// _ = "end of CoverTab[89080]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:228
		_go_fuzz_dep_.CoverTab[89081]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:228
		// _ = "end of CoverTab[89081]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:228
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:228
	// _ = "end of CoverTab[89076]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:228
	_go_fuzz_dep_.CoverTab[89077]++
												outP := uint(0)
												out := s.Out[:maxHeaderSize]

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:233
	for remaining > 1 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:233
		_go_fuzz_dep_.CoverTab[89082]++
													if previous0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:234
			_go_fuzz_dep_.CoverTab[89089]++
														start := charnum
														for s.norm[charnum] == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:236
				_go_fuzz_dep_.CoverTab[89093]++
															charnum++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:237
				// _ = "end of CoverTab[89093]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:238
			// _ = "end of CoverTab[89089]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:238
			_go_fuzz_dep_.CoverTab[89090]++
														for charnum >= start+24 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:239
				_go_fuzz_dep_.CoverTab[89094]++
															start += 24
															bitStream += uint32(0xFFFF) << bitCount
															out[outP] = byte(bitStream)
															out[outP+1] = byte(bitStream >> 8)
															outP += 2
															bitStream >>= 16
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:245
				// _ = "end of CoverTab[89094]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:246
			// _ = "end of CoverTab[89090]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:246
			_go_fuzz_dep_.CoverTab[89091]++
														for charnum >= start+3 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:247
				_go_fuzz_dep_.CoverTab[89095]++
															start += 3
															bitStream += 3 << bitCount
															bitCount += 2
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:250
				// _ = "end of CoverTab[89095]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:251
			// _ = "end of CoverTab[89091]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:251
			_go_fuzz_dep_.CoverTab[89092]++
														bitStream += uint32(charnum-start) << bitCount
														bitCount += 2
														if bitCount > 16 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:254
				_go_fuzz_dep_.CoverTab[89096]++
															out[outP] = byte(bitStream)
															out[outP+1] = byte(bitStream >> 8)
															outP += 2
															bitStream >>= 16
															bitCount -= 16
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:259
				// _ = "end of CoverTab[89096]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:260
				_go_fuzz_dep_.CoverTab[89097]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:260
				// _ = "end of CoverTab[89097]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:260
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:260
			// _ = "end of CoverTab[89092]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:261
			_go_fuzz_dep_.CoverTab[89098]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:261
			// _ = "end of CoverTab[89098]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:261
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:261
		// _ = "end of CoverTab[89082]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:261
		_go_fuzz_dep_.CoverTab[89083]++

													count := s.norm[charnum]
													charnum++
													max := (2*threshold - 1) - remaining
													if count < 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:266
			_go_fuzz_dep_.CoverTab[89099]++
														remaining += count
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:267
			// _ = "end of CoverTab[89099]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:268
			_go_fuzz_dep_.CoverTab[89100]++
														remaining -= count
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:269
			// _ = "end of CoverTab[89100]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:270
		// _ = "end of CoverTab[89083]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:270
		_go_fuzz_dep_.CoverTab[89084]++
													count++
													if count >= threshold {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:272
			_go_fuzz_dep_.CoverTab[89101]++
														count += max
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:273
			// _ = "end of CoverTab[89101]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:274
			_go_fuzz_dep_.CoverTab[89102]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:274
			// _ = "end of CoverTab[89102]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:274
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:274
		// _ = "end of CoverTab[89084]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:274
		_go_fuzz_dep_.CoverTab[89085]++
													bitStream += uint32(count) << bitCount
													bitCount += nbBits
													if count < max {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:277
			_go_fuzz_dep_.CoverTab[89103]++
														bitCount--
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:278
			// _ = "end of CoverTab[89103]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:279
			_go_fuzz_dep_.CoverTab[89104]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:279
			// _ = "end of CoverTab[89104]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:279
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:279
		// _ = "end of CoverTab[89085]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:279
		_go_fuzz_dep_.CoverTab[89086]++

													previous0 = count == 1
													if remaining < 1 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:282
			_go_fuzz_dep_.CoverTab[89105]++
														return errors.New("internal error: remaining<1")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:283
			// _ = "end of CoverTab[89105]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:284
			_go_fuzz_dep_.CoverTab[89106]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:284
			// _ = "end of CoverTab[89106]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:284
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:284
		// _ = "end of CoverTab[89086]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:284
		_go_fuzz_dep_.CoverTab[89087]++
													for remaining < threshold {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:285
			_go_fuzz_dep_.CoverTab[89107]++
														nbBits--
														threshold >>= 1
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:287
			// _ = "end of CoverTab[89107]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:288
		// _ = "end of CoverTab[89087]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:288
		_go_fuzz_dep_.CoverTab[89088]++

													if bitCount > 16 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:290
			_go_fuzz_dep_.CoverTab[89108]++
														out[outP] = byte(bitStream)
														out[outP+1] = byte(bitStream >> 8)
														outP += 2
														bitStream >>= 16
														bitCount -= 16
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:295
			// _ = "end of CoverTab[89108]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:296
			_go_fuzz_dep_.CoverTab[89109]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:296
			// _ = "end of CoverTab[89109]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:296
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:296
		// _ = "end of CoverTab[89088]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:297
	// _ = "end of CoverTab[89077]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:297
	_go_fuzz_dep_.CoverTab[89078]++

												out[outP] = byte(bitStream)
												out[outP+1] = byte(bitStream >> 8)
												outP += (bitCount + 7) / 8

												if charnum > s.symbolLen {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:303
		_go_fuzz_dep_.CoverTab[89110]++
													return errors.New("internal error: charnum > s.symbolLen")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:304
		// _ = "end of CoverTab[89110]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:305
		_go_fuzz_dep_.CoverTab[89111]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:305
		// _ = "end of CoverTab[89111]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:305
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:305
	// _ = "end of CoverTab[89078]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:305
	_go_fuzz_dep_.CoverTab[89079]++
												s.Out = out[:outP]
												return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:307
	// _ = "end of CoverTab[89079]"
}

// symbolTransform contains the state transform for a symbol.
type symbolTransform struct {
	deltaFindState	int32
	deltaNbBits	uint32
}

// String prints values as a human readable string.
func (s symbolTransform) String() string {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:317
	_go_fuzz_dep_.CoverTab[89112]++
												return fmt.Sprintf("dnbits: %08x, fs:%d", s.deltaNbBits, s.deltaFindState)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:318
	// _ = "end of CoverTab[89112]"
}

// cTable contains tables used for compression.
type cTable struct {
	tableSymbol	[]byte
	stateTable	[]uint16
	symbolTT	[]symbolTransform
}

// allocCtable will allocate tables needed for compression.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:328
// If existing tables a re big enough, they are simply re-used.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:330
func (s *Scratch) allocCtable() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:330
	_go_fuzz_dep_.CoverTab[89113]++
												tableSize := 1 << s.actualTableLog

												if cap(s.ct.tableSymbol) < tableSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:333
		_go_fuzz_dep_.CoverTab[89117]++
													s.ct.tableSymbol = make([]byte, tableSize)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:334
		// _ = "end of CoverTab[89117]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:335
		_go_fuzz_dep_.CoverTab[89118]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:335
		// _ = "end of CoverTab[89118]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:335
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:335
	// _ = "end of CoverTab[89113]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:335
	_go_fuzz_dep_.CoverTab[89114]++
												s.ct.tableSymbol = s.ct.tableSymbol[:tableSize]

												ctSize := tableSize
												if cap(s.ct.stateTable) < ctSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:339
		_go_fuzz_dep_.CoverTab[89119]++
													s.ct.stateTable = make([]uint16, ctSize)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:340
		// _ = "end of CoverTab[89119]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:341
		_go_fuzz_dep_.CoverTab[89120]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:341
		// _ = "end of CoverTab[89120]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:341
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:341
	// _ = "end of CoverTab[89114]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:341
	_go_fuzz_dep_.CoverTab[89115]++
												s.ct.stateTable = s.ct.stateTable[:ctSize]

												if cap(s.ct.symbolTT) < 256 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:344
		_go_fuzz_dep_.CoverTab[89121]++
													s.ct.symbolTT = make([]symbolTransform, 256)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:345
		// _ = "end of CoverTab[89121]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:346
		_go_fuzz_dep_.CoverTab[89122]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:346
		// _ = "end of CoverTab[89122]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:346
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:346
	// _ = "end of CoverTab[89115]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:346
	_go_fuzz_dep_.CoverTab[89116]++
												s.ct.symbolTT = s.ct.symbolTT[:256]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:347
	// _ = "end of CoverTab[89116]"
}

// buildCTable will populate the compression table so it is ready to be used.
func (s *Scratch) buildCTable() error {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:351
	_go_fuzz_dep_.CoverTab[89123]++
												tableSize := uint32(1 << s.actualTableLog)
												highThreshold := tableSize - 1
												var cumul [maxSymbolValue + 2]int16

												s.allocCtable()
												tableSymbol := s.ct.tableSymbol[:tableSize]

												{
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:359
		_go_fuzz_dep_.CoverTab[89127]++
													cumul[0] = 0
													for ui, v := range s.norm[:s.symbolLen-1] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:361
			_go_fuzz_dep_.CoverTab[89131]++
														u := byte(ui)
														if v == -1 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:363
				_go_fuzz_dep_.CoverTab[89132]++

															cumul[u+1] = cumul[u] + 1
															tableSymbol[highThreshold] = u
															highThreshold--
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:367
				// _ = "end of CoverTab[89132]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:368
				_go_fuzz_dep_.CoverTab[89133]++
															cumul[u+1] = cumul[u] + v
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:369
				// _ = "end of CoverTab[89133]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:370
			// _ = "end of CoverTab[89131]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:371
		// _ = "end of CoverTab[89127]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:371
		_go_fuzz_dep_.CoverTab[89128]++

													u := int(s.symbolLen - 1)
													v := s.norm[s.symbolLen-1]
													if v == -1 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:375
			_go_fuzz_dep_.CoverTab[89134]++

														cumul[u+1] = cumul[u] + 1
														tableSymbol[highThreshold] = byte(u)
														highThreshold--
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:379
			// _ = "end of CoverTab[89134]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:380
			_go_fuzz_dep_.CoverTab[89135]++
														cumul[u+1] = cumul[u] + v
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:381
			// _ = "end of CoverTab[89135]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:382
		// _ = "end of CoverTab[89128]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:382
		_go_fuzz_dep_.CoverTab[89129]++
													if uint32(cumul[s.symbolLen]) != tableSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:383
			_go_fuzz_dep_.CoverTab[89136]++
														return fmt.Errorf("internal error: expected cumul[s.symbolLen] (%d) == tableSize (%d)", cumul[s.symbolLen], tableSize)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:384
			// _ = "end of CoverTab[89136]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:385
			_go_fuzz_dep_.CoverTab[89137]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:385
			// _ = "end of CoverTab[89137]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:385
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:385
		// _ = "end of CoverTab[89129]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:385
		_go_fuzz_dep_.CoverTab[89130]++
													cumul[s.symbolLen] = int16(tableSize) + 1
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:386
		// _ = "end of CoverTab[89130]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:387
	// _ = "end of CoverTab[89123]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:387
	_go_fuzz_dep_.CoverTab[89124]++

												s.zeroBits = false
												{
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:390
		_go_fuzz_dep_.CoverTab[89138]++
													step := tableStep(tableSize)
													tableMask := tableSize - 1
													var position uint32

													largeLimit := int16(1 << (s.actualTableLog - 1))
													for ui, v := range s.norm[:s.symbolLen] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:396
			_go_fuzz_dep_.CoverTab[89140]++
														symbol := byte(ui)
														if v > largeLimit {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:398
				_go_fuzz_dep_.CoverTab[89142]++
															s.zeroBits = true
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:399
				// _ = "end of CoverTab[89142]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:400
				_go_fuzz_dep_.CoverTab[89143]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:400
				// _ = "end of CoverTab[89143]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:400
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:400
			// _ = "end of CoverTab[89140]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:400
			_go_fuzz_dep_.CoverTab[89141]++
														for nbOccurrences := int16(0); nbOccurrences < v; nbOccurrences++ {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:401
				_go_fuzz_dep_.CoverTab[89144]++
															tableSymbol[position] = symbol
															position = (position + step) & tableMask
															for position > highThreshold {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:404
					_go_fuzz_dep_.CoverTab[89145]++
																position = (position + step) & tableMask
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:405
					// _ = "end of CoverTab[89145]"
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:406
				// _ = "end of CoverTab[89144]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:407
			// _ = "end of CoverTab[89141]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:408
		// _ = "end of CoverTab[89138]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:408
		_go_fuzz_dep_.CoverTab[89139]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:411
		if position != 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:411
			_go_fuzz_dep_.CoverTab[89146]++
														return errors.New("position!=0")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:412
			// _ = "end of CoverTab[89146]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:413
			_go_fuzz_dep_.CoverTab[89147]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:413
			// _ = "end of CoverTab[89147]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:413
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:413
		// _ = "end of CoverTab[89139]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:414
	// _ = "end of CoverTab[89124]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:414
	_go_fuzz_dep_.CoverTab[89125]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:417
	table := s.ct.stateTable
	{
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:418
		_go_fuzz_dep_.CoverTab[89148]++
													tsi := int(tableSize)
													for u, v := range tableSymbol {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:420
			_go_fuzz_dep_.CoverTab[89149]++

														table[cumul[v]] = uint16(tsi + u)
														cumul[v]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:423
			// _ = "end of CoverTab[89149]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:424
		// _ = "end of CoverTab[89148]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:425
	// _ = "end of CoverTab[89125]"

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:428
	{
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:428
		_go_fuzz_dep_.CoverTab[89150]++
													total := int16(0)
													symbolTT := s.ct.symbolTT[:s.symbolLen]
													tableLog := s.actualTableLog
													tl := (uint32(tableLog) << 16) - (1 << tableLog)
													for i, v := range s.norm[:s.symbolLen] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:433
			_go_fuzz_dep_.CoverTab[89152]++
														switch v {
			case 0:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:435
				_go_fuzz_dep_.CoverTab[89153]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:435
				// _ = "end of CoverTab[89153]"
			case -1, 1:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:436
				_go_fuzz_dep_.CoverTab[89154]++
															symbolTT[i].deltaNbBits = tl
															symbolTT[i].deltaFindState = int32(total - 1)
															total++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:439
				// _ = "end of CoverTab[89154]"
			default:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:440
				_go_fuzz_dep_.CoverTab[89155]++
															maxBitsOut := uint32(tableLog) - highBits(uint32(v-1))
															minStatePlus := uint32(v) << maxBitsOut
															symbolTT[i].deltaNbBits = (maxBitsOut << 16) - minStatePlus
															symbolTT[i].deltaFindState = int32(total - v)
															total += v
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:445
				// _ = "end of CoverTab[89155]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:446
			// _ = "end of CoverTab[89152]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:447
		// _ = "end of CoverTab[89150]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:447
		_go_fuzz_dep_.CoverTab[89151]++
													if total != int16(tableSize) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:448
			_go_fuzz_dep_.CoverTab[89156]++
														return fmt.Errorf("total mismatch %d (got) != %d (want)", total, tableSize)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:449
			// _ = "end of CoverTab[89156]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:450
			_go_fuzz_dep_.CoverTab[89157]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:450
			// _ = "end of CoverTab[89157]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:450
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:450
		// _ = "end of CoverTab[89151]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:451
	_go_fuzz_dep_.CoverTab[89126]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:452
	// _ = "end of CoverTab[89126]"
}

// countSimple will create a simple histogram in s.count.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:455
// Returns the biggest count.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:455
// Does not update s.clearCount.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:458
func (s *Scratch) countSimple(in []byte) (max int) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:458
	_go_fuzz_dep_.CoverTab[89158]++
												for _, v := range in {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:459
		_go_fuzz_dep_.CoverTab[89161]++
													s.count[v]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:460
		// _ = "end of CoverTab[89161]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:461
	// _ = "end of CoverTab[89158]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:461
	_go_fuzz_dep_.CoverTab[89159]++
												m := uint32(0)
												for i, v := range s.count[:] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:463
		_go_fuzz_dep_.CoverTab[89162]++
													if v > m {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:464
			_go_fuzz_dep_.CoverTab[89164]++
														m = v
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:465
			// _ = "end of CoverTab[89164]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:466
			_go_fuzz_dep_.CoverTab[89165]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:466
			// _ = "end of CoverTab[89165]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:466
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:466
		// _ = "end of CoverTab[89162]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:466
		_go_fuzz_dep_.CoverTab[89163]++
													if v > 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:467
			_go_fuzz_dep_.CoverTab[89166]++
														s.symbolLen = uint16(i) + 1
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:468
			// _ = "end of CoverTab[89166]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:469
			_go_fuzz_dep_.CoverTab[89167]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:469
			// _ = "end of CoverTab[89167]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:469
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:469
		// _ = "end of CoverTab[89163]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:470
	// _ = "end of CoverTab[89159]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:470
	_go_fuzz_dep_.CoverTab[89160]++
												return int(m)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:471
	// _ = "end of CoverTab[89160]"
}

// minTableLog provides the minimum logSize to safely represent a distribution.
func (s *Scratch) minTableLog() uint8 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:475
	_go_fuzz_dep_.CoverTab[89168]++
												minBitsSrc := highBits(uint32(s.br.remain()-1)) + 1
												minBitsSymbols := highBits(uint32(s.symbolLen-1)) + 2
												if minBitsSrc < minBitsSymbols {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:478
		_go_fuzz_dep_.CoverTab[89170]++
													return uint8(minBitsSrc)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:479
		// _ = "end of CoverTab[89170]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:480
		_go_fuzz_dep_.CoverTab[89171]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:480
		// _ = "end of CoverTab[89171]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:480
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:480
	// _ = "end of CoverTab[89168]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:480
	_go_fuzz_dep_.CoverTab[89169]++
												return uint8(minBitsSymbols)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:481
	// _ = "end of CoverTab[89169]"
}

// optimalTableLog calculates and sets the optimal tableLog in s.actualTableLog
func (s *Scratch) optimalTableLog() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:485
	_go_fuzz_dep_.CoverTab[89172]++
												tableLog := s.TableLog
												minBits := s.minTableLog()
												maxBitsSrc := uint8(highBits(uint32(s.br.remain()-1))) - 2
												if maxBitsSrc < tableLog {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:489
		_go_fuzz_dep_.CoverTab[89177]++

													tableLog = maxBitsSrc
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:491
		// _ = "end of CoverTab[89177]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:492
		_go_fuzz_dep_.CoverTab[89178]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:492
		// _ = "end of CoverTab[89178]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:492
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:492
	// _ = "end of CoverTab[89172]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:492
	_go_fuzz_dep_.CoverTab[89173]++
												if minBits > tableLog {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:493
		_go_fuzz_dep_.CoverTab[89179]++
													tableLog = minBits
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:494
		// _ = "end of CoverTab[89179]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:495
		_go_fuzz_dep_.CoverTab[89180]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:495
		// _ = "end of CoverTab[89180]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:495
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:495
	// _ = "end of CoverTab[89173]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:495
	_go_fuzz_dep_.CoverTab[89174]++

												if tableLog < minTablelog {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:497
		_go_fuzz_dep_.CoverTab[89181]++
													tableLog = minTablelog
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:498
		// _ = "end of CoverTab[89181]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:499
		_go_fuzz_dep_.CoverTab[89182]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:499
		// _ = "end of CoverTab[89182]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:499
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:499
	// _ = "end of CoverTab[89174]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:499
	_go_fuzz_dep_.CoverTab[89175]++
												if tableLog > maxTableLog {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:500
		_go_fuzz_dep_.CoverTab[89183]++
													tableLog = maxTableLog
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:501
		// _ = "end of CoverTab[89183]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:502
		_go_fuzz_dep_.CoverTab[89184]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:502
		// _ = "end of CoverTab[89184]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:502
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:502
	// _ = "end of CoverTab[89175]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:502
	_go_fuzz_dep_.CoverTab[89176]++
												s.actualTableLog = tableLog
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:503
	// _ = "end of CoverTab[89176]"
}

var rtbTable = [...]uint32{0, 473195, 504333, 520860, 550000, 700000, 750000, 830000}

// normalizeCount will normalize the count of the symbols so
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:508
// the total is equal to the table size.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:510
func (s *Scratch) normalizeCount() error {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:510
	_go_fuzz_dep_.CoverTab[89185]++
												var (
		tableLog		= s.actualTableLog
		scale			= 62 - uint64(tableLog)
		step			= (1 << 62) / uint64(s.br.remain())
		vStep			= uint64(1) << (scale - 20)
		stillToDistribute	= int16(1 << tableLog)
		largest			int
		largestP		int16
		lowThreshold		= (uint32)(s.br.remain() >> tableLog)
	)

	for i, cnt := range s.count[:s.symbolLen] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:522
		_go_fuzz_dep_.CoverTab[89188]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:526
		if cnt == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:526
			_go_fuzz_dep_.CoverTab[89190]++
														s.norm[i] = 0
														continue
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:528
			// _ = "end of CoverTab[89190]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:529
			_go_fuzz_dep_.CoverTab[89191]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:529
			// _ = "end of CoverTab[89191]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:529
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:529
		// _ = "end of CoverTab[89188]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:529
		_go_fuzz_dep_.CoverTab[89189]++
													if cnt <= lowThreshold {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:530
			_go_fuzz_dep_.CoverTab[89192]++
														s.norm[i] = -1
														stillToDistribute--
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:532
			// _ = "end of CoverTab[89192]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:533
			_go_fuzz_dep_.CoverTab[89193]++
														proba := (int16)((uint64(cnt) * step) >> scale)
														if proba < 8 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:535
				_go_fuzz_dep_.CoverTab[89196]++
															restToBeat := vStep * uint64(rtbTable[proba])
															v := uint64(cnt)*step - (uint64(proba) << scale)
															if v > restToBeat {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:538
					_go_fuzz_dep_.CoverTab[89197]++
																proba++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:539
					// _ = "end of CoverTab[89197]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:540
					_go_fuzz_dep_.CoverTab[89198]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:540
					// _ = "end of CoverTab[89198]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:540
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:540
				// _ = "end of CoverTab[89196]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:541
				_go_fuzz_dep_.CoverTab[89199]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:541
				// _ = "end of CoverTab[89199]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:541
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:541
			// _ = "end of CoverTab[89193]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:541
			_go_fuzz_dep_.CoverTab[89194]++
														if proba > largestP {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:542
				_go_fuzz_dep_.CoverTab[89200]++
															largestP = proba
															largest = i
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:544
				// _ = "end of CoverTab[89200]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:545
				_go_fuzz_dep_.CoverTab[89201]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:545
				// _ = "end of CoverTab[89201]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:545
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:545
			// _ = "end of CoverTab[89194]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:545
			_go_fuzz_dep_.CoverTab[89195]++
														s.norm[i] = proba
														stillToDistribute -= proba
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:547
			// _ = "end of CoverTab[89195]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:548
		// _ = "end of CoverTab[89189]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:549
	// _ = "end of CoverTab[89185]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:549
	_go_fuzz_dep_.CoverTab[89186]++

												if -stillToDistribute >= (s.norm[largest] >> 1) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:551
		_go_fuzz_dep_.CoverTab[89202]++

													return s.normalizeCount2()
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:553
		// _ = "end of CoverTab[89202]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:554
		_go_fuzz_dep_.CoverTab[89203]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:554
		// _ = "end of CoverTab[89203]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:554
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:554
	// _ = "end of CoverTab[89186]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:554
	_go_fuzz_dep_.CoverTab[89187]++
												s.norm[largest] += stillToDistribute
												return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:556
	// _ = "end of CoverTab[89187]"
}

// Secondary normalization method.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:559
// To be used when primary method fails.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:561
func (s *Scratch) normalizeCount2() error {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:561
	_go_fuzz_dep_.CoverTab[89204]++
												const notYetAssigned = -2
												var (
		distributed	uint32
		total		= uint32(s.br.remain())
		tableLog	= s.actualTableLog
		lowThreshold	= total >> tableLog
		lowOne		= (total * 3) >> (tableLog + 1)
	)
	for i, cnt := range s.count[:s.symbolLen] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:570
		_go_fuzz_dep_.CoverTab[89210]++
													if cnt == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:571
			_go_fuzz_dep_.CoverTab[89214]++
														s.norm[i] = 0
														continue
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:573
			// _ = "end of CoverTab[89214]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:574
			_go_fuzz_dep_.CoverTab[89215]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:574
			// _ = "end of CoverTab[89215]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:574
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:574
		// _ = "end of CoverTab[89210]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:574
		_go_fuzz_dep_.CoverTab[89211]++
													if cnt <= lowThreshold {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:575
			_go_fuzz_dep_.CoverTab[89216]++
														s.norm[i] = -1
														distributed++
														total -= cnt
														continue
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:579
			// _ = "end of CoverTab[89216]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:580
			_go_fuzz_dep_.CoverTab[89217]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:580
			// _ = "end of CoverTab[89217]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:580
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:580
		// _ = "end of CoverTab[89211]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:580
		_go_fuzz_dep_.CoverTab[89212]++
													if cnt <= lowOne {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:581
			_go_fuzz_dep_.CoverTab[89218]++
														s.norm[i] = 1
														distributed++
														total -= cnt
														continue
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:585
			// _ = "end of CoverTab[89218]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:586
			_go_fuzz_dep_.CoverTab[89219]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:586
			// _ = "end of CoverTab[89219]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:586
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:586
		// _ = "end of CoverTab[89212]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:586
		_go_fuzz_dep_.CoverTab[89213]++
													s.norm[i] = notYetAssigned
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:587
		// _ = "end of CoverTab[89213]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:588
	// _ = "end of CoverTab[89204]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:588
	_go_fuzz_dep_.CoverTab[89205]++
												toDistribute := (1 << tableLog) - distributed

												if (total / toDistribute) > lowOne {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:591
		_go_fuzz_dep_.CoverTab[89220]++

													lowOne = (total * 3) / (toDistribute * 2)
													for i, cnt := range s.count[:s.symbolLen] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:594
			_go_fuzz_dep_.CoverTab[89222]++
														if (s.norm[i] == notYetAssigned) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:595
				_go_fuzz_dep_.CoverTab[89223]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:595
				return (cnt <= lowOne)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:595
				// _ = "end of CoverTab[89223]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:595
			}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:595
				_go_fuzz_dep_.CoverTab[89224]++
															s.norm[i] = 1
															distributed++
															total -= cnt
															continue
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:599
				// _ = "end of CoverTab[89224]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:600
				_go_fuzz_dep_.CoverTab[89225]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:600
				// _ = "end of CoverTab[89225]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:600
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:600
			// _ = "end of CoverTab[89222]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:601
		// _ = "end of CoverTab[89220]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:601
		_go_fuzz_dep_.CoverTab[89221]++
													toDistribute = (1 << tableLog) - distributed
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:602
		// _ = "end of CoverTab[89221]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:603
		_go_fuzz_dep_.CoverTab[89226]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:603
		// _ = "end of CoverTab[89226]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:603
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:603
	// _ = "end of CoverTab[89205]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:603
	_go_fuzz_dep_.CoverTab[89206]++
												if distributed == uint32(s.symbolLen)+1 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:604
		_go_fuzz_dep_.CoverTab[89227]++
		// all values are pretty poor;
		//   probably incompressible data (should have already been detected);
		//   find max, then give all remaining points to max
		var maxV int
		var maxC uint32
		for i, cnt := range s.count[:s.symbolLen] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:610
			_go_fuzz_dep_.CoverTab[89229]++
														if cnt > maxC {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:611
				_go_fuzz_dep_.CoverTab[89230]++
															maxV = i
															maxC = cnt
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:613
				// _ = "end of CoverTab[89230]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:614
				_go_fuzz_dep_.CoverTab[89231]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:614
				// _ = "end of CoverTab[89231]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:614
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:614
			// _ = "end of CoverTab[89229]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:615
		// _ = "end of CoverTab[89227]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:615
		_go_fuzz_dep_.CoverTab[89228]++
													s.norm[maxV] += int16(toDistribute)
													return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:617
		// _ = "end of CoverTab[89228]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:618
		_go_fuzz_dep_.CoverTab[89232]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:618
		// _ = "end of CoverTab[89232]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:618
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:618
	// _ = "end of CoverTab[89206]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:618
	_go_fuzz_dep_.CoverTab[89207]++

												if total == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:620
		_go_fuzz_dep_.CoverTab[89233]++

													for i := uint32(0); toDistribute > 0; i = (i + 1) % (uint32(s.symbolLen)) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:622
			_go_fuzz_dep_.CoverTab[89235]++
														if s.norm[i] > 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:623
				_go_fuzz_dep_.CoverTab[89236]++
															toDistribute--
															s.norm[i]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:625
				// _ = "end of CoverTab[89236]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:626
				_go_fuzz_dep_.CoverTab[89237]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:626
				// _ = "end of CoverTab[89237]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:626
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:626
			// _ = "end of CoverTab[89235]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:627
		// _ = "end of CoverTab[89233]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:627
		_go_fuzz_dep_.CoverTab[89234]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:628
		// _ = "end of CoverTab[89234]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:629
		_go_fuzz_dep_.CoverTab[89238]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:629
		// _ = "end of CoverTab[89238]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:629
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:629
	// _ = "end of CoverTab[89207]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:629
	_go_fuzz_dep_.CoverTab[89208]++

												var (
		vStepLog	= 62 - uint64(tableLog)
		mid		= uint64((1 << (vStepLog - 1)) - 1)
		rStep		= (((1 << vStepLog) * uint64(toDistribute)) + mid) / uint64(total)	// scale on remaining
		tmpTotal	= mid
	)
	for i, cnt := range s.count[:s.symbolLen] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:637
		_go_fuzz_dep_.CoverTab[89239]++
													if s.norm[i] == notYetAssigned {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:638
			_go_fuzz_dep_.CoverTab[89240]++
														var (
				end	= tmpTotal + uint64(cnt)*rStep
				sStart	= uint32(tmpTotal >> vStepLog)
				sEnd	= uint32(end >> vStepLog)
				weight	= sEnd - sStart
			)
			if weight < 1 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:645
				_go_fuzz_dep_.CoverTab[89242]++
															return errors.New("weight < 1")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:646
				// _ = "end of CoverTab[89242]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:647
				_go_fuzz_dep_.CoverTab[89243]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:647
				// _ = "end of CoverTab[89243]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:647
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:647
			// _ = "end of CoverTab[89240]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:647
			_go_fuzz_dep_.CoverTab[89241]++
														s.norm[i] = int16(weight)
														tmpTotal = end
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:649
			// _ = "end of CoverTab[89241]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:650
			_go_fuzz_dep_.CoverTab[89244]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:650
			// _ = "end of CoverTab[89244]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:650
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:650
		// _ = "end of CoverTab[89239]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:651
	// _ = "end of CoverTab[89208]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:651
	_go_fuzz_dep_.CoverTab[89209]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:652
	// _ = "end of CoverTab[89209]"
}

// validateNorm validates the normalized histogram table.
func (s *Scratch) validateNorm() (err error) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:656
	_go_fuzz_dep_.CoverTab[89245]++
												var total int
												for _, v := range s.norm[:s.symbolLen] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:658
		_go_fuzz_dep_.CoverTab[89250]++
													if v >= 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:659
			_go_fuzz_dep_.CoverTab[89251]++
														total += int(v)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:660
			// _ = "end of CoverTab[89251]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:661
			_go_fuzz_dep_.CoverTab[89252]++
														total -= int(v)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:662
			// _ = "end of CoverTab[89252]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:663
		// _ = "end of CoverTab[89250]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:664
	// _ = "end of CoverTab[89245]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:664
	_go_fuzz_dep_.CoverTab[89246]++
												defer func() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:665
		_go_fuzz_dep_.CoverTab[89253]++
													if err == nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:666
			_go_fuzz_dep_.CoverTab[89255]++
														return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:667
			// _ = "end of CoverTab[89255]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:668
			_go_fuzz_dep_.CoverTab[89256]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:668
			// _ = "end of CoverTab[89256]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:668
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:668
		// _ = "end of CoverTab[89253]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:668
		_go_fuzz_dep_.CoverTab[89254]++
													fmt.Printf("selected TableLog: %d, Symbol length: %d\n", s.actualTableLog, s.symbolLen)
													for i, v := range s.norm[:s.symbolLen] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:670
			_go_fuzz_dep_.CoverTab[89257]++
														fmt.Printf("%3d: %5d -> %4d \n", i, s.count[i], v)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:671
			// _ = "end of CoverTab[89257]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:672
		// _ = "end of CoverTab[89254]"
	}()
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:673
	// _ = "end of CoverTab[89246]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:673
	_go_fuzz_dep_.CoverTab[89247]++
												if total != (1 << s.actualTableLog) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:674
		_go_fuzz_dep_.CoverTab[89258]++
													return fmt.Errorf("warning: Total == %d != %d", total, 1<<s.actualTableLog)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:675
		// _ = "end of CoverTab[89258]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:676
		_go_fuzz_dep_.CoverTab[89259]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:676
		// _ = "end of CoverTab[89259]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:676
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:676
	// _ = "end of CoverTab[89247]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:676
	_go_fuzz_dep_.CoverTab[89248]++
												for i, v := range s.count[s.symbolLen:] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:677
		_go_fuzz_dep_.CoverTab[89260]++
													if v != 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:678
			_go_fuzz_dep_.CoverTab[89261]++
														return fmt.Errorf("warning: Found symbol out of range, %d after cut", i)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:679
			// _ = "end of CoverTab[89261]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:680
			_go_fuzz_dep_.CoverTab[89262]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:680
			// _ = "end of CoverTab[89262]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:680
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:680
		// _ = "end of CoverTab[89260]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:681
	// _ = "end of CoverTab[89248]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:681
	_go_fuzz_dep_.CoverTab[89249]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:682
	// _ = "end of CoverTab[89249]"
}

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:683
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/compress.go:683
var _ = _go_fuzz_dep_.CoverTab
