// Copyright 2018 Klaus Post. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Based on work Copyright (c) 2013, Yann Collet, released under BSD License.

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:6
package fse

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:6
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:6
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:6
)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:6
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:6
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:6
)

import "fmt"

// bitWriter will write bits.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:10
// First bit will be LSB of the first byte of output.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:12
type bitWriter struct {
	bitContainer	uint64
	nBits		uint8
	out		[]byte
}

// bitMask16 is bitmasks. Has extra to avoid bounds check.
var bitMask16 = [32]uint16{
	0, 1, 3, 7, 0xF, 0x1F,
	0x3F, 0x7F, 0xFF, 0x1FF, 0x3FF, 0x7FF,
	0xFFF, 0x1FFF, 0x3FFF, 0x7FFF, 0xFFFF, 0xFFFF,
	0xFFFF, 0xFFFF, 0xFFFF, 0xFFFF, 0xFFFF, 0xFFFF,
	0xFFFF, 0xFFFF}	/* up to 16 bits */

// addBits16NC will add up to 16 bits.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:26
// It will not check if there is space for them,
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:26
// so the caller must ensure that it has flushed recently.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:29
func (b *bitWriter) addBits16NC(value uint16, bits uint8) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:29
	_go_fuzz_dep_.CoverTab[88982]++
												b.bitContainer |= uint64(value&bitMask16[bits&31]) << (b.nBits & 63)
												b.nBits += bits
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:31
	// _ = "end of CoverTab[88982]"
}

// addBits16Clean will add up to 16 bits. value may not contain more set bits than indicated.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:34
// It will not check if there is space for them, so the caller must ensure that it has flushed recently.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:36
func (b *bitWriter) addBits16Clean(value uint16, bits uint8) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:36
	_go_fuzz_dep_.CoverTab[88983]++
												b.bitContainer |= uint64(value) << (b.nBits & 63)
												b.nBits += bits
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:38
	// _ = "end of CoverTab[88983]"
}

// addBits16ZeroNC will add up to 16 bits.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:41
// It will not check if there is space for them,
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:41
// so the caller must ensure that it has flushed recently.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:41
// This is fastest if bits can be zero.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:45
func (b *bitWriter) addBits16ZeroNC(value uint16, bits uint8) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:45
	_go_fuzz_dep_.CoverTab[88984]++
												if bits == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:46
		_go_fuzz_dep_.CoverTab[88986]++
													return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:47
		// _ = "end of CoverTab[88986]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:48
		_go_fuzz_dep_.CoverTab[88987]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:48
		// _ = "end of CoverTab[88987]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:48
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:48
	// _ = "end of CoverTab[88984]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:48
	_go_fuzz_dep_.CoverTab[88985]++
												value <<= (16 - bits) & 15
												value >>= (16 - bits) & 15
												b.bitContainer |= uint64(value) << (b.nBits & 63)
												b.nBits += bits
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:52
	// _ = "end of CoverTab[88985]"
}

// flush will flush all pending full bytes.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:55
// There will be at least 56 bits available for writing when this has been called.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:55
// Using flush32 is faster, but leaves less space for writing.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:58
func (b *bitWriter) flush() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:58
	_go_fuzz_dep_.CoverTab[88988]++
												v := b.nBits >> 3
												switch v {
	case 0:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:61
		_go_fuzz_dep_.CoverTab[88990]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:61
		// _ = "end of CoverTab[88990]"
	case 1:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:62
		_go_fuzz_dep_.CoverTab[88991]++
													b.out = append(b.out,
			byte(b.bitContainer),
		)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:65
		// _ = "end of CoverTab[88991]"
	case 2:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:66
		_go_fuzz_dep_.CoverTab[88992]++
													b.out = append(b.out,
			byte(b.bitContainer),
			byte(b.bitContainer>>8),
		)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:70
		// _ = "end of CoverTab[88992]"
	case 3:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:71
		_go_fuzz_dep_.CoverTab[88993]++
													b.out = append(b.out,
			byte(b.bitContainer),
			byte(b.bitContainer>>8),
			byte(b.bitContainer>>16),
		)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:76
		// _ = "end of CoverTab[88993]"
	case 4:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:77
		_go_fuzz_dep_.CoverTab[88994]++
													b.out = append(b.out,
			byte(b.bitContainer),
			byte(b.bitContainer>>8),
			byte(b.bitContainer>>16),
			byte(b.bitContainer>>24),
		)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:83
		// _ = "end of CoverTab[88994]"
	case 5:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:84
		_go_fuzz_dep_.CoverTab[88995]++
													b.out = append(b.out,
			byte(b.bitContainer),
			byte(b.bitContainer>>8),
			byte(b.bitContainer>>16),
			byte(b.bitContainer>>24),
			byte(b.bitContainer>>32),
		)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:91
		// _ = "end of CoverTab[88995]"
	case 6:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:92
		_go_fuzz_dep_.CoverTab[88996]++
													b.out = append(b.out,
			byte(b.bitContainer),
			byte(b.bitContainer>>8),
			byte(b.bitContainer>>16),
			byte(b.bitContainer>>24),
			byte(b.bitContainer>>32),
			byte(b.bitContainer>>40),
		)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:100
		// _ = "end of CoverTab[88996]"
	case 7:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:101
		_go_fuzz_dep_.CoverTab[88997]++
													b.out = append(b.out,
			byte(b.bitContainer),
			byte(b.bitContainer>>8),
			byte(b.bitContainer>>16),
			byte(b.bitContainer>>24),
			byte(b.bitContainer>>32),
			byte(b.bitContainer>>40),
			byte(b.bitContainer>>48),
		)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:110
		// _ = "end of CoverTab[88997]"
	case 8:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:111
		_go_fuzz_dep_.CoverTab[88998]++
													b.out = append(b.out,
			byte(b.bitContainer),
			byte(b.bitContainer>>8),
			byte(b.bitContainer>>16),
			byte(b.bitContainer>>24),
			byte(b.bitContainer>>32),
			byte(b.bitContainer>>40),
			byte(b.bitContainer>>48),
			byte(b.bitContainer>>56),
		)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:121
		// _ = "end of CoverTab[88998]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:122
		_go_fuzz_dep_.CoverTab[88999]++
													panic(fmt.Errorf("bits (%d) > 64", b.nBits))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:123
		// _ = "end of CoverTab[88999]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:124
	// _ = "end of CoverTab[88988]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:124
	_go_fuzz_dep_.CoverTab[88989]++
												b.bitContainer >>= v << 3
												b.nBits &= 7
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:126
	// _ = "end of CoverTab[88989]"
}

// flush32 will flush out, so there are at least 32 bits available for writing.
func (b *bitWriter) flush32() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:130
	_go_fuzz_dep_.CoverTab[89000]++
												if b.nBits < 32 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:131
		_go_fuzz_dep_.CoverTab[89002]++
													return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:132
		// _ = "end of CoverTab[89002]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:133
		_go_fuzz_dep_.CoverTab[89003]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:133
		// _ = "end of CoverTab[89003]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:133
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:133
	// _ = "end of CoverTab[89000]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:133
	_go_fuzz_dep_.CoverTab[89001]++
												b.out = append(b.out,
		byte(b.bitContainer),
		byte(b.bitContainer>>8),
		byte(b.bitContainer>>16),
		byte(b.bitContainer>>24))
												b.nBits -= 32
												b.bitContainer >>= 32
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:140
	// _ = "end of CoverTab[89001]"
}

// flushAlign will flush remaining full bytes and align to next byte boundary.
func (b *bitWriter) flushAlign() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:144
	_go_fuzz_dep_.CoverTab[89004]++
												nbBytes := (b.nBits + 7) >> 3
												for i := uint8(0); i < nbBytes; i++ {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:146
		_go_fuzz_dep_.CoverTab[89006]++
													b.out = append(b.out, byte(b.bitContainer>>(i*8)))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:147
		// _ = "end of CoverTab[89006]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:148
	// _ = "end of CoverTab[89004]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:148
	_go_fuzz_dep_.CoverTab[89005]++
												b.nBits = 0
												b.bitContainer = 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:150
	// _ = "end of CoverTab[89005]"
}

// close will write the alignment bit and write the final byte(s)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:153
// to the output.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:155
func (b *bitWriter) close() error {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:155
	_go_fuzz_dep_.CoverTab[89007]++

												b.addBits16Clean(1, 1)

												b.flushAlign()
												return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:160
	// _ = "end of CoverTab[89007]"
}

// reset and continue writing by appending to out.
func (b *bitWriter) reset(out []byte) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:164
	_go_fuzz_dep_.CoverTab[89008]++
												b.bitContainer = 0
												b.nBits = 0
												b.out = out
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:167
	// _ = "end of CoverTab[89008]"
}

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:168
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitwriter.go:168
var _ = _go_fuzz_dep_.CoverTab
