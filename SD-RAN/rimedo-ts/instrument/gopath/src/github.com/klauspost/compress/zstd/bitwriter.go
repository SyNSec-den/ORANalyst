// Copyright 2018 Klaus Post. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Based on work Copyright (c) 2013, Yann Collet, released under BSD License.

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:6
package zstd

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:6
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:6
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:6
)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:6
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:6
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:6
)

import "fmt"

// bitWriter will write bits.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:10
// First bit will be LSB of the first byte of output.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:12
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

var bitMask32 = [32]uint32{
	0, 1, 3, 7, 0xF, 0x1F, 0x3F, 0x7F, 0xFF,
	0x1FF, 0x3FF, 0x7FF, 0xFFF, 0x1FFF, 0x3FFF, 0x7FFF, 0xFFFF,
	0x1ffff, 0x3ffff, 0x7FFFF, 0xfFFFF, 0x1fFFFF, 0x3fFFFF, 0x7fFFFF, 0xffFFFF,
	0x1ffFFFF, 0x3ffFFFF, 0x7ffFFFF, 0xfffFFFF, 0x1fffFFFF, 0x3fffFFFF, 0x7fffFFFF,
}	// up to 32 bits

// addBits16NC will add up to 16 bits.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:33
// It will not check if there is space for them,
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:33
// so the caller must ensure that it has flushed recently.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:36
func (b *bitWriter) addBits16NC(value uint16, bits uint8) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:36
	_go_fuzz_dep_.CoverTab[90869]++
												b.bitContainer |= uint64(value&bitMask16[bits&31]) << (b.nBits & 63)
												b.nBits += bits
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:38
	// _ = "end of CoverTab[90869]"
}

// addBits32NC will add up to 31 bits.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:41
// It will not check if there is space for them,
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:41
// so the caller must ensure that it has flushed recently.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:44
func (b *bitWriter) addBits32NC(value uint32, bits uint8) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:44
	_go_fuzz_dep_.CoverTab[90870]++
												b.bitContainer |= uint64(value&bitMask32[bits&31]) << (b.nBits & 63)
												b.nBits += bits
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:46
	// _ = "end of CoverTab[90870]"
}

// addBits64NC will add up to 64 bits.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:49
// There must be space for 32 bits.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:51
func (b *bitWriter) addBits64NC(value uint64, bits uint8) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:51
	_go_fuzz_dep_.CoverTab[90871]++
												if bits <= 31 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:52
		_go_fuzz_dep_.CoverTab[90873]++
													b.addBits32Clean(uint32(value), bits)
													return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:54
		// _ = "end of CoverTab[90873]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:55
		_go_fuzz_dep_.CoverTab[90874]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:55
		// _ = "end of CoverTab[90874]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:55
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:55
	// _ = "end of CoverTab[90871]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:55
	_go_fuzz_dep_.CoverTab[90872]++
												b.addBits32Clean(uint32(value), 32)
												b.flush32()
												b.addBits32Clean(uint32(value>>32), bits-32)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:58
	// _ = "end of CoverTab[90872]"
}

// addBits32Clean will add up to 32 bits.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:61
// It will not check if there is space for them.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:61
// The input must not contain more bits than specified.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:64
func (b *bitWriter) addBits32Clean(value uint32, bits uint8) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:64
	_go_fuzz_dep_.CoverTab[90875]++
												b.bitContainer |= uint64(value) << (b.nBits & 63)
												b.nBits += bits
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:66
	// _ = "end of CoverTab[90875]"
}

// addBits16Clean will add up to 16 bits. value may not contain more set bits than indicated.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:69
// It will not check if there is space for them, so the caller must ensure that it has flushed recently.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:71
func (b *bitWriter) addBits16Clean(value uint16, bits uint8) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:71
	_go_fuzz_dep_.CoverTab[90876]++
												b.bitContainer |= uint64(value) << (b.nBits & 63)
												b.nBits += bits
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:73
	// _ = "end of CoverTab[90876]"
}

// flush will flush all pending full bytes.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:76
// There will be at least 56 bits available for writing when this has been called.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:76
// Using flush32 is faster, but leaves less space for writing.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:79
func (b *bitWriter) flush() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:79
	_go_fuzz_dep_.CoverTab[90877]++
												v := b.nBits >> 3
												switch v {
	case 0:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:82
		_go_fuzz_dep_.CoverTab[90879]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:82
		// _ = "end of CoverTab[90879]"
	case 1:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:83
		_go_fuzz_dep_.CoverTab[90880]++
													b.out = append(b.out,
			byte(b.bitContainer),
		)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:86
		// _ = "end of CoverTab[90880]"
	case 2:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:87
		_go_fuzz_dep_.CoverTab[90881]++
													b.out = append(b.out,
			byte(b.bitContainer),
			byte(b.bitContainer>>8),
		)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:91
		// _ = "end of CoverTab[90881]"
	case 3:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:92
		_go_fuzz_dep_.CoverTab[90882]++
													b.out = append(b.out,
			byte(b.bitContainer),
			byte(b.bitContainer>>8),
			byte(b.bitContainer>>16),
		)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:97
		// _ = "end of CoverTab[90882]"
	case 4:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:98
		_go_fuzz_dep_.CoverTab[90883]++
													b.out = append(b.out,
			byte(b.bitContainer),
			byte(b.bitContainer>>8),
			byte(b.bitContainer>>16),
			byte(b.bitContainer>>24),
		)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:104
		// _ = "end of CoverTab[90883]"
	case 5:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:105
		_go_fuzz_dep_.CoverTab[90884]++
													b.out = append(b.out,
			byte(b.bitContainer),
			byte(b.bitContainer>>8),
			byte(b.bitContainer>>16),
			byte(b.bitContainer>>24),
			byte(b.bitContainer>>32),
		)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:112
		// _ = "end of CoverTab[90884]"
	case 6:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:113
		_go_fuzz_dep_.CoverTab[90885]++
													b.out = append(b.out,
			byte(b.bitContainer),
			byte(b.bitContainer>>8),
			byte(b.bitContainer>>16),
			byte(b.bitContainer>>24),
			byte(b.bitContainer>>32),
			byte(b.bitContainer>>40),
		)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:121
		// _ = "end of CoverTab[90885]"
	case 7:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:122
		_go_fuzz_dep_.CoverTab[90886]++
													b.out = append(b.out,
			byte(b.bitContainer),
			byte(b.bitContainer>>8),
			byte(b.bitContainer>>16),
			byte(b.bitContainer>>24),
			byte(b.bitContainer>>32),
			byte(b.bitContainer>>40),
			byte(b.bitContainer>>48),
		)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:131
		// _ = "end of CoverTab[90886]"
	case 8:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:132
		_go_fuzz_dep_.CoverTab[90887]++
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
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:142
		// _ = "end of CoverTab[90887]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:143
		_go_fuzz_dep_.CoverTab[90888]++
													panic(fmt.Errorf("bits (%d) > 64", b.nBits))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:144
		// _ = "end of CoverTab[90888]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:145
	// _ = "end of CoverTab[90877]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:145
	_go_fuzz_dep_.CoverTab[90878]++
												b.bitContainer >>= v << 3
												b.nBits &= 7
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:147
	// _ = "end of CoverTab[90878]"
}

// flush32 will flush out, so there are at least 32 bits available for writing.
func (b *bitWriter) flush32() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:151
	_go_fuzz_dep_.CoverTab[90889]++
												if b.nBits < 32 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:152
		_go_fuzz_dep_.CoverTab[90891]++
													return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:153
		// _ = "end of CoverTab[90891]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:154
		_go_fuzz_dep_.CoverTab[90892]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:154
		// _ = "end of CoverTab[90892]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:154
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:154
	// _ = "end of CoverTab[90889]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:154
	_go_fuzz_dep_.CoverTab[90890]++
												b.out = append(b.out,
		byte(b.bitContainer),
		byte(b.bitContainer>>8),
		byte(b.bitContainer>>16),
		byte(b.bitContainer>>24))
												b.nBits -= 32
												b.bitContainer >>= 32
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:161
	// _ = "end of CoverTab[90890]"
}

// flushAlign will flush remaining full bytes and align to next byte boundary.
func (b *bitWriter) flushAlign() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:165
	_go_fuzz_dep_.CoverTab[90893]++
												nbBytes := (b.nBits + 7) >> 3
												for i := uint8(0); i < nbBytes; i++ {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:167
		_go_fuzz_dep_.CoverTab[90895]++
													b.out = append(b.out, byte(b.bitContainer>>(i*8)))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:168
		// _ = "end of CoverTab[90895]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:169
	// _ = "end of CoverTab[90893]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:169
	_go_fuzz_dep_.CoverTab[90894]++
												b.nBits = 0
												b.bitContainer = 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:171
	// _ = "end of CoverTab[90894]"
}

// close will write the alignment bit and write the final byte(s)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:174
// to the output.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:176
func (b *bitWriter) close() error {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:176
	_go_fuzz_dep_.CoverTab[90896]++

												b.addBits16Clean(1, 1)

												b.flushAlign()
												return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:181
	// _ = "end of CoverTab[90896]"
}

// reset and continue writing by appending to out.
func (b *bitWriter) reset(out []byte) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:185
	_go_fuzz_dep_.CoverTab[90897]++
												b.bitContainer = 0
												b.nBits = 0
												b.out = out
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:188
	// _ = "end of CoverTab[90897]"
}

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:189
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitwriter.go:189
var _ = _go_fuzz_dep_.CoverTab
