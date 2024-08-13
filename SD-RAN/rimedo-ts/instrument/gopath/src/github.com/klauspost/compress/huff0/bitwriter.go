// Copyright 2018 Klaus Post. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Based on work Copyright (c) 2013, Yann Collet, released under BSD License.

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:6
package huff0

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:6
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:6
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:6
)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:6
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:6
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:6
)

import "fmt"

// bitWriter will write bits.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:10
// First bit will be LSB of the first byte of output.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:12
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
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:26
// It will not check if there is space for them,
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:26
// so the caller must ensure that it has flushed recently.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:29
func (b *bitWriter) addBits16NC(value uint16, bits uint8) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:29
	_go_fuzz_dep_.CoverTab[89488]++
												b.bitContainer |= uint64(value&bitMask16[bits&31]) << (b.nBits & 63)
												b.nBits += bits
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:31
	// _ = "end of CoverTab[89488]"
}

// addBits16Clean will add up to 16 bits. value may not contain more set bits than indicated.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:34
// It will not check if there is space for them, so the caller must ensure that it has flushed recently.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:36
func (b *bitWriter) addBits16Clean(value uint16, bits uint8) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:36
	_go_fuzz_dep_.CoverTab[89489]++
												b.bitContainer |= uint64(value) << (b.nBits & 63)
												b.nBits += bits
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:38
	// _ = "end of CoverTab[89489]"
}

// encSymbol will add up to 16 bits. value may not contain more set bits than indicated.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:41
// It will not check if there is space for them, so the caller must ensure that it has flushed recently.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:43
func (b *bitWriter) encSymbol(ct cTable, symbol byte) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:43
	_go_fuzz_dep_.CoverTab[89490]++
												enc := ct[symbol]
												b.bitContainer |= uint64(enc.val) << (b.nBits & 63)
												if false {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:46
		_go_fuzz_dep_.CoverTab[89492]++
													if enc.nBits == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:47
			_go_fuzz_dep_.CoverTab[89493]++
														panic("nbits 0")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:48
			// _ = "end of CoverTab[89493]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:49
			_go_fuzz_dep_.CoverTab[89494]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:49
			// _ = "end of CoverTab[89494]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:49
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:49
		// _ = "end of CoverTab[89492]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:50
		_go_fuzz_dep_.CoverTab[89495]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:50
		// _ = "end of CoverTab[89495]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:50
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:50
	// _ = "end of CoverTab[89490]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:50
	_go_fuzz_dep_.CoverTab[89491]++
												b.nBits += enc.nBits
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:51
	// _ = "end of CoverTab[89491]"
}

// encTwoSymbols will add up to 32 bits. value may not contain more set bits than indicated.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:54
// It will not check if there is space for them, so the caller must ensure that it has flushed recently.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:56
func (b *bitWriter) encTwoSymbols(ct cTable, av, bv byte) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:56
	_go_fuzz_dep_.CoverTab[89496]++
												encA := ct[av]
												encB := ct[bv]
												sh := b.nBits & 63
												combined := uint64(encA.val) | (uint64(encB.val) << (encA.nBits & 63))
												b.bitContainer |= combined << sh
												if false {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:62
		_go_fuzz_dep_.CoverTab[89498]++
													if encA.nBits == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:63
			_go_fuzz_dep_.CoverTab[89500]++
														panic("nbitsA 0")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:64
			// _ = "end of CoverTab[89500]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:65
			_go_fuzz_dep_.CoverTab[89501]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:65
			// _ = "end of CoverTab[89501]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:65
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:65
		// _ = "end of CoverTab[89498]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:65
		_go_fuzz_dep_.CoverTab[89499]++
													if encB.nBits == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:66
			_go_fuzz_dep_.CoverTab[89502]++
														panic("nbitsB 0")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:67
			// _ = "end of CoverTab[89502]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:68
			_go_fuzz_dep_.CoverTab[89503]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:68
			// _ = "end of CoverTab[89503]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:68
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:68
		// _ = "end of CoverTab[89499]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:69
		_go_fuzz_dep_.CoverTab[89504]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:69
		// _ = "end of CoverTab[89504]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:69
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:69
	// _ = "end of CoverTab[89496]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:69
	_go_fuzz_dep_.CoverTab[89497]++
												b.nBits += encA.nBits + encB.nBits
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:70
	// _ = "end of CoverTab[89497]"
}

// addBits16ZeroNC will add up to 16 bits.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:73
// It will not check if there is space for them,
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:73
// so the caller must ensure that it has flushed recently.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:73
// This is fastest if bits can be zero.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:77
func (b *bitWriter) addBits16ZeroNC(value uint16, bits uint8) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:77
	_go_fuzz_dep_.CoverTab[89505]++
												if bits == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:78
		_go_fuzz_dep_.CoverTab[89507]++
													return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:79
		// _ = "end of CoverTab[89507]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:80
		_go_fuzz_dep_.CoverTab[89508]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:80
		// _ = "end of CoverTab[89508]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:80
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:80
	// _ = "end of CoverTab[89505]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:80
	_go_fuzz_dep_.CoverTab[89506]++
												value <<= (16 - bits) & 15
												value >>= (16 - bits) & 15
												b.bitContainer |= uint64(value) << (b.nBits & 63)
												b.nBits += bits
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:84
	// _ = "end of CoverTab[89506]"
}

// flush will flush all pending full bytes.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:87
// There will be at least 56 bits available for writing when this has been called.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:87
// Using flush32 is faster, but leaves less space for writing.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:90
func (b *bitWriter) flush() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:90
	_go_fuzz_dep_.CoverTab[89509]++
												v := b.nBits >> 3
												switch v {
	case 0:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:93
		_go_fuzz_dep_.CoverTab[89511]++
													return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:94
		// _ = "end of CoverTab[89511]"
	case 1:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:95
		_go_fuzz_dep_.CoverTab[89512]++
													b.out = append(b.out,
			byte(b.bitContainer),
		)
													b.bitContainer >>= 1 << 3
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:99
		// _ = "end of CoverTab[89512]"
	case 2:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:100
		_go_fuzz_dep_.CoverTab[89513]++
													b.out = append(b.out,
			byte(b.bitContainer),
			byte(b.bitContainer>>8),
		)
													b.bitContainer >>= 2 << 3
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:105
		// _ = "end of CoverTab[89513]"
	case 3:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:106
		_go_fuzz_dep_.CoverTab[89514]++
													b.out = append(b.out,
			byte(b.bitContainer),
			byte(b.bitContainer>>8),
			byte(b.bitContainer>>16),
		)
													b.bitContainer >>= 3 << 3
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:112
		// _ = "end of CoverTab[89514]"
	case 4:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:113
		_go_fuzz_dep_.CoverTab[89515]++
													b.out = append(b.out,
			byte(b.bitContainer),
			byte(b.bitContainer>>8),
			byte(b.bitContainer>>16),
			byte(b.bitContainer>>24),
		)
													b.bitContainer >>= 4 << 3
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:120
		// _ = "end of CoverTab[89515]"
	case 5:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:121
		_go_fuzz_dep_.CoverTab[89516]++
													b.out = append(b.out,
			byte(b.bitContainer),
			byte(b.bitContainer>>8),
			byte(b.bitContainer>>16),
			byte(b.bitContainer>>24),
			byte(b.bitContainer>>32),
		)
													b.bitContainer >>= 5 << 3
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:129
		// _ = "end of CoverTab[89516]"
	case 6:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:130
		_go_fuzz_dep_.CoverTab[89517]++
													b.out = append(b.out,
			byte(b.bitContainer),
			byte(b.bitContainer>>8),
			byte(b.bitContainer>>16),
			byte(b.bitContainer>>24),
			byte(b.bitContainer>>32),
			byte(b.bitContainer>>40),
		)
													b.bitContainer >>= 6 << 3
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:139
		// _ = "end of CoverTab[89517]"
	case 7:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:140
		_go_fuzz_dep_.CoverTab[89518]++
													b.out = append(b.out,
			byte(b.bitContainer),
			byte(b.bitContainer>>8),
			byte(b.bitContainer>>16),
			byte(b.bitContainer>>24),
			byte(b.bitContainer>>32),
			byte(b.bitContainer>>40),
			byte(b.bitContainer>>48),
		)
													b.bitContainer >>= 7 << 3
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:150
		// _ = "end of CoverTab[89518]"
	case 8:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:151
		_go_fuzz_dep_.CoverTab[89519]++
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
													b.bitContainer = 0
													b.nBits = 0
													return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:164
		// _ = "end of CoverTab[89519]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:165
		_go_fuzz_dep_.CoverTab[89520]++
													panic(fmt.Errorf("bits (%d) > 64", b.nBits))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:166
		// _ = "end of CoverTab[89520]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:167
	// _ = "end of CoverTab[89509]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:167
	_go_fuzz_dep_.CoverTab[89510]++
												b.nBits &= 7
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:168
	// _ = "end of CoverTab[89510]"
}

// flush32 will flush out, so there are at least 32 bits available for writing.
func (b *bitWriter) flush32() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:172
	_go_fuzz_dep_.CoverTab[89521]++
												if b.nBits < 32 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:173
		_go_fuzz_dep_.CoverTab[89523]++
													return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:174
		// _ = "end of CoverTab[89523]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:175
		_go_fuzz_dep_.CoverTab[89524]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:175
		// _ = "end of CoverTab[89524]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:175
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:175
	// _ = "end of CoverTab[89521]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:175
	_go_fuzz_dep_.CoverTab[89522]++
												b.out = append(b.out,
		byte(b.bitContainer),
		byte(b.bitContainer>>8),
		byte(b.bitContainer>>16),
		byte(b.bitContainer>>24))
												b.nBits -= 32
												b.bitContainer >>= 32
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:182
	// _ = "end of CoverTab[89522]"
}

// flushAlign will flush remaining full bytes and align to next byte boundary.
func (b *bitWriter) flushAlign() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:186
	_go_fuzz_dep_.CoverTab[89525]++
												nbBytes := (b.nBits + 7) >> 3
												for i := uint8(0); i < nbBytes; i++ {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:188
		_go_fuzz_dep_.CoverTab[89527]++
													b.out = append(b.out, byte(b.bitContainer>>(i*8)))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:189
		// _ = "end of CoverTab[89527]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:190
	// _ = "end of CoverTab[89525]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:190
	_go_fuzz_dep_.CoverTab[89526]++
												b.nBits = 0
												b.bitContainer = 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:192
	// _ = "end of CoverTab[89526]"
}

// close will write the alignment bit and write the final byte(s)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:195
// to the output.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:197
func (b *bitWriter) close() error {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:197
	_go_fuzz_dep_.CoverTab[89528]++

												b.addBits16Clean(1, 1)

												b.flushAlign()
												return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:202
	// _ = "end of CoverTab[89528]"
}

// reset and continue writing by appending to out.
func (b *bitWriter) reset(out []byte) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:206
	_go_fuzz_dep_.CoverTab[89529]++
												b.bitContainer = 0
												b.nBits = 0
												b.out = out
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:209
	// _ = "end of CoverTab[89529]"
}

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:210
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitwriter.go:210
var _ = _go_fuzz_dep_.CoverTab
