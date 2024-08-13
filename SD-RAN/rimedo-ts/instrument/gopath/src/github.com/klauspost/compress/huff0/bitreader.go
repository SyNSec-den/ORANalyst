// Copyright 2018 Klaus Post. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Based on work Copyright (c) 2013, Yann Collet, released under BSD License.

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:6
package huff0

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:6
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:6
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:6
)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:6
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:6
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:6
)

import (
	"encoding/binary"
	"errors"
	"io"
)

// bitReader reads a bitstream in reverse.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:14
// The last set bit indicates the start of the stream and is used
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:14
// for aligning the input.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:17
type bitReader struct {
	in		[]byte
	off		uint	// next byte to read is at in[off - 1]
	value		uint64
	bitsRead	uint8
}

// init initializes and resets the bit reader.
func (b *bitReader) init(in []byte) error {
	if len(in) < 1 {
		return errors.New("corrupt stream: too short")
	}
	b.in = in
	b.off = uint(len(in))

	v := in[len(in)-1]
	if v == 0 {
		return errors.New("corrupt stream, did not find end of stream")
	}
	b.bitsRead = 64
	b.value = 0
	if len(in) >= 8 {
		b.fillFastStart()
	} else {
		b.fill()
		b.fill()
	}
	b.bitsRead += 8 - uint8(highBit32(uint32(v)))
	return nil
}

// peekBitsFast requires that at least one bit is requested every time.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:48
// There are no checks if the buffer is filled.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:50
func (b *bitReader) peekBitsFast(n uint8) uint16 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:50
	_go_fuzz_dep_.CoverTab[89425]++
												const regMask = 64 - 1
												v := uint16((b.value << (b.bitsRead & regMask)) >> ((regMask + 1 - n) & regMask))
												return v
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:53
	// _ = "end of CoverTab[89425]"
}

// fillFast() will make sure at least 32 bits are available.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:56
// There must be at least 4 bytes available.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:58
func (b *bitReader) fillFast() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:58
	_go_fuzz_dep_.CoverTab[89426]++
												if b.bitsRead < 32 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:59
		_go_fuzz_dep_.CoverTab[89428]++
													return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:60
		// _ = "end of CoverTab[89428]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:61
		_go_fuzz_dep_.CoverTab[89429]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:61
		// _ = "end of CoverTab[89429]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:61
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:61
	// _ = "end of CoverTab[89426]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:61
	_go_fuzz_dep_.CoverTab[89427]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:64
	v := b.in[b.off-4 : b.off]
												v = v[:4]
												low := (uint32(v[0])) | (uint32(v[1]) << 8) | (uint32(v[2]) << 16) | (uint32(v[3]) << 24)
												b.value = (b.value << 32) | uint64(low)
												b.bitsRead -= 32
												b.off -= 4
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:69
	// _ = "end of CoverTab[89427]"
}

func (b *bitReader) advance(n uint8) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:72
	_go_fuzz_dep_.CoverTab[89430]++
												b.bitsRead += n
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:73
	// _ = "end of CoverTab[89430]"
}

// fillFastStart() assumes the bitreader is empty and there is at least 8 bytes to read.
func (b *bitReader) fillFastStart() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:77
	_go_fuzz_dep_.CoverTab[89431]++

												b.value = binary.LittleEndian.Uint64(b.in[b.off-8:])
												b.bitsRead = 0
												b.off -= 8
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:81
	// _ = "end of CoverTab[89431]"
}

// fill() will make sure at least 32 bits are available.
func (b *bitReader) fill() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:85
	_go_fuzz_dep_.CoverTab[89432]++
												if b.bitsRead < 32 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:86
		_go_fuzz_dep_.CoverTab[89435]++
													return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:87
		// _ = "end of CoverTab[89435]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:88
		_go_fuzz_dep_.CoverTab[89436]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:88
		// _ = "end of CoverTab[89436]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:88
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:88
	// _ = "end of CoverTab[89432]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:88
	_go_fuzz_dep_.CoverTab[89433]++
												if b.off > 4 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:89
		_go_fuzz_dep_.CoverTab[89437]++
													v := b.in[b.off-4:]
													v = v[:4]
													low := (uint32(v[0])) | (uint32(v[1]) << 8) | (uint32(v[2]) << 16) | (uint32(v[3]) << 24)
													b.value = (b.value << 32) | uint64(low)
													b.bitsRead -= 32
													b.off -= 4
													return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:96
		// _ = "end of CoverTab[89437]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:97
		_go_fuzz_dep_.CoverTab[89438]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:97
		// _ = "end of CoverTab[89438]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:97
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:97
	// _ = "end of CoverTab[89433]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:97
	_go_fuzz_dep_.CoverTab[89434]++
												for b.off > 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:98
		_go_fuzz_dep_.CoverTab[89439]++
													b.value = (b.value << 8) | uint64(b.in[b.off-1])
													b.bitsRead -= 8
													b.off--
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:101
		// _ = "end of CoverTab[89439]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:102
	// _ = "end of CoverTab[89434]"
}

// finished returns true if all bits have been read from the bit stream.
func (b *bitReader) finished() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:106
	_go_fuzz_dep_.CoverTab[89440]++
												return b.off == 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:107
		_go_fuzz_dep_.CoverTab[89441]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:107
		return b.bitsRead >= 64
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:107
		// _ = "end of CoverTab[89441]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:107
	}()
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:107
	// _ = "end of CoverTab[89440]"
}

// close the bitstream and returns an error if out-of-buffer reads occurred.
func (b *bitReader) close() error {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:111
	_go_fuzz_dep_.CoverTab[89442]++

												b.in = nil
												if b.bitsRead > 64 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:114
		_go_fuzz_dep_.CoverTab[89444]++
													return io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:115
		// _ = "end of CoverTab[89444]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:116
		_go_fuzz_dep_.CoverTab[89445]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:116
		// _ = "end of CoverTab[89445]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:116
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:116
	// _ = "end of CoverTab[89442]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:116
	_go_fuzz_dep_.CoverTab[89443]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:117
	// _ = "end of CoverTab[89443]"
}

// bitReader reads a bitstream in reverse.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:120
// The last set bit indicates the start of the stream and is used
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:120
// for aligning the input.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:123
type bitReaderBytes struct {
	in		[]byte
	off		uint	// next byte to read is at in[off - 1]
	value		uint64
	bitsRead	uint8
}

// init initializes and resets the bit reader.
func (b *bitReaderBytes) init(in []byte) error {
	if len(in) < 1 {
		return errors.New("corrupt stream: too short")
	}
	b.in = in
	b.off = uint(len(in))

	v := in[len(in)-1]
	if v == 0 {
		return errors.New("corrupt stream, did not find end of stream")
	}
	b.bitsRead = 64
	b.value = 0
	if len(in) >= 8 {
		b.fillFastStart()
	} else {
		b.fill()
		b.fill()
	}
	b.advance(8 - uint8(highBit32(uint32(v))))
	return nil
}

// peekBitsFast requires that at least one bit is requested every time.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:154
// There are no checks if the buffer is filled.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:156
func (b *bitReaderBytes) peekByteFast() uint8 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:156
	_go_fuzz_dep_.CoverTab[89446]++
												got := uint8(b.value >> 56)
												return got
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:158
	// _ = "end of CoverTab[89446]"
}

func (b *bitReaderBytes) advance(n uint8) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:161
	_go_fuzz_dep_.CoverTab[89447]++
												b.bitsRead += n
												b.value <<= n & 63
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:163
	// _ = "end of CoverTab[89447]"
}

// fillFast() will make sure at least 32 bits are available.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:166
// There must be at least 4 bytes available.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:168
func (b *bitReaderBytes) fillFast() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:168
	_go_fuzz_dep_.CoverTab[89448]++
												if b.bitsRead < 32 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:169
		_go_fuzz_dep_.CoverTab[89450]++
													return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:170
		// _ = "end of CoverTab[89450]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:171
		_go_fuzz_dep_.CoverTab[89451]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:171
		// _ = "end of CoverTab[89451]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:171
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:171
	// _ = "end of CoverTab[89448]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:171
	_go_fuzz_dep_.CoverTab[89449]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:174
	v := b.in[b.off-4 : b.off]
												v = v[:4]
												low := (uint32(v[0])) | (uint32(v[1]) << 8) | (uint32(v[2]) << 16) | (uint32(v[3]) << 24)
												b.value |= uint64(low) << (b.bitsRead - 32)
												b.bitsRead -= 32
												b.off -= 4
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:179
	// _ = "end of CoverTab[89449]"
}

// fillFastStart() assumes the bitReaderBytes is empty and there is at least 8 bytes to read.
func (b *bitReaderBytes) fillFastStart() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:183
	_go_fuzz_dep_.CoverTab[89452]++

												b.value = binary.LittleEndian.Uint64(b.in[b.off-8:])
												b.bitsRead = 0
												b.off -= 8
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:187
	// _ = "end of CoverTab[89452]"
}

// fill() will make sure at least 32 bits are available.
func (b *bitReaderBytes) fill() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:191
	_go_fuzz_dep_.CoverTab[89453]++
												if b.bitsRead < 32 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:192
		_go_fuzz_dep_.CoverTab[89456]++
													return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:193
		// _ = "end of CoverTab[89456]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:194
		_go_fuzz_dep_.CoverTab[89457]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:194
		// _ = "end of CoverTab[89457]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:194
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:194
	// _ = "end of CoverTab[89453]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:194
	_go_fuzz_dep_.CoverTab[89454]++
												if b.off > 4 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:195
		_go_fuzz_dep_.CoverTab[89458]++
													v := b.in[b.off-4:]
													v = v[:4]
													low := (uint32(v[0])) | (uint32(v[1]) << 8) | (uint32(v[2]) << 16) | (uint32(v[3]) << 24)
													b.value |= uint64(low) << (b.bitsRead - 32)
													b.bitsRead -= 32
													b.off -= 4
													return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:202
		// _ = "end of CoverTab[89458]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:203
		_go_fuzz_dep_.CoverTab[89459]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:203
		// _ = "end of CoverTab[89459]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:203
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:203
	// _ = "end of CoverTab[89454]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:203
	_go_fuzz_dep_.CoverTab[89455]++
												for b.off > 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:204
		_go_fuzz_dep_.CoverTab[89460]++
													b.value |= uint64(b.in[b.off-1]) << (b.bitsRead - 8)
													b.bitsRead -= 8
													b.off--
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:207
		// _ = "end of CoverTab[89460]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:208
	// _ = "end of CoverTab[89455]"
}

// finished returns true if all bits have been read from the bit stream.
func (b *bitReaderBytes) finished() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:212
	_go_fuzz_dep_.CoverTab[89461]++
												return b.off == 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:213
		_go_fuzz_dep_.CoverTab[89462]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:213
		return b.bitsRead >= 64
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:213
		// _ = "end of CoverTab[89462]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:213
	}()
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:213
	// _ = "end of CoverTab[89461]"
}

// close the bitstream and returns an error if out-of-buffer reads occurred.
func (b *bitReaderBytes) close() error {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:217
	_go_fuzz_dep_.CoverTab[89463]++

												b.in = nil
												if b.bitsRead > 64 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:220
		_go_fuzz_dep_.CoverTab[89465]++
													return io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:221
		// _ = "end of CoverTab[89465]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:222
		_go_fuzz_dep_.CoverTab[89466]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:222
		// _ = "end of CoverTab[89466]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:222
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:222
	// _ = "end of CoverTab[89463]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:222
	_go_fuzz_dep_.CoverTab[89464]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:223
	// _ = "end of CoverTab[89464]"
}

// bitReaderShifted reads a bitstream in reverse.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:226
// The last set bit indicates the start of the stream and is used
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:226
// for aligning the input.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:229
type bitReaderShifted struct {
	in		[]byte
	off		uint	// next byte to read is at in[off - 1]
	value		uint64
	bitsRead	uint8
}

// init initializes and resets the bit reader.
func (b *bitReaderShifted) init(in []byte) error {
	if len(in) < 1 {
		return errors.New("corrupt stream: too short")
	}
	b.in = in
	b.off = uint(len(in))

	v := in[len(in)-1]
	if v == 0 {
		return errors.New("corrupt stream, did not find end of stream")
	}
	b.bitsRead = 64
	b.value = 0
	if len(in) >= 8 {
		b.fillFastStart()
	} else {
		b.fill()
		b.fill()
	}
	b.advance(8 - uint8(highBit32(uint32(v))))
	return nil
}

// peekBitsFast requires that at least one bit is requested every time.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:260
// There are no checks if the buffer is filled.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:262
func (b *bitReaderShifted) peekBitsFast(n uint8) uint16 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:262
	_go_fuzz_dep_.CoverTab[89467]++
												return uint16(b.value >> ((64 - n) & 63))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:263
	// _ = "end of CoverTab[89467]"
}

func (b *bitReaderShifted) advance(n uint8) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:266
	_go_fuzz_dep_.CoverTab[89468]++
												b.bitsRead += n
												b.value <<= n & 63
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:268
	// _ = "end of CoverTab[89468]"
}

// fillFast() will make sure at least 32 bits are available.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:271
// There must be at least 4 bytes available.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:273
func (b *bitReaderShifted) fillFast() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:273
	_go_fuzz_dep_.CoverTab[89469]++
												if b.bitsRead < 32 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:274
		_go_fuzz_dep_.CoverTab[89471]++
													return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:275
		// _ = "end of CoverTab[89471]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:276
		_go_fuzz_dep_.CoverTab[89472]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:276
		// _ = "end of CoverTab[89472]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:276
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:276
	// _ = "end of CoverTab[89469]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:276
	_go_fuzz_dep_.CoverTab[89470]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:279
	v := b.in[b.off-4 : b.off]
												v = v[:4]
												low := (uint32(v[0])) | (uint32(v[1]) << 8) | (uint32(v[2]) << 16) | (uint32(v[3]) << 24)
												b.value |= uint64(low) << ((b.bitsRead - 32) & 63)
												b.bitsRead -= 32
												b.off -= 4
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:284
	// _ = "end of CoverTab[89470]"
}

// fillFastStart() assumes the bitReaderShifted is empty and there is at least 8 bytes to read.
func (b *bitReaderShifted) fillFastStart() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:288
	_go_fuzz_dep_.CoverTab[89473]++

												b.value = binary.LittleEndian.Uint64(b.in[b.off-8:])
												b.bitsRead = 0
												b.off -= 8
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:292
	// _ = "end of CoverTab[89473]"
}

// fill() will make sure at least 32 bits are available.
func (b *bitReaderShifted) fill() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:296
	_go_fuzz_dep_.CoverTab[89474]++
												if b.bitsRead < 32 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:297
		_go_fuzz_dep_.CoverTab[89477]++
													return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:298
		// _ = "end of CoverTab[89477]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:299
		_go_fuzz_dep_.CoverTab[89478]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:299
		// _ = "end of CoverTab[89478]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:299
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:299
	// _ = "end of CoverTab[89474]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:299
	_go_fuzz_dep_.CoverTab[89475]++
												if b.off > 4 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:300
		_go_fuzz_dep_.CoverTab[89479]++
													v := b.in[b.off-4:]
													v = v[:4]
													low := (uint32(v[0])) | (uint32(v[1]) << 8) | (uint32(v[2]) << 16) | (uint32(v[3]) << 24)
													b.value |= uint64(low) << ((b.bitsRead - 32) & 63)
													b.bitsRead -= 32
													b.off -= 4
													return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:307
		// _ = "end of CoverTab[89479]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:308
		_go_fuzz_dep_.CoverTab[89480]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:308
		// _ = "end of CoverTab[89480]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:308
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:308
	// _ = "end of CoverTab[89475]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:308
	_go_fuzz_dep_.CoverTab[89476]++
												for b.off > 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:309
		_go_fuzz_dep_.CoverTab[89481]++
													b.value |= uint64(b.in[b.off-1]) << ((b.bitsRead - 8) & 63)
													b.bitsRead -= 8
													b.off--
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:312
		// _ = "end of CoverTab[89481]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:313
	// _ = "end of CoverTab[89476]"
}

// finished returns true if all bits have been read from the bit stream.
func (b *bitReaderShifted) finished() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:317
	_go_fuzz_dep_.CoverTab[89482]++
												return b.off == 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:318
		_go_fuzz_dep_.CoverTab[89483]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:318
		return b.bitsRead >= 64
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:318
		// _ = "end of CoverTab[89483]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:318
	}()
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:318
	// _ = "end of CoverTab[89482]"
}

// close the bitstream and returns an error if out-of-buffer reads occurred.
func (b *bitReaderShifted) close() error {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:322
	_go_fuzz_dep_.CoverTab[89484]++

												b.in = nil
												if b.bitsRead > 64 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:325
		_go_fuzz_dep_.CoverTab[89486]++
													return io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:326
		// _ = "end of CoverTab[89486]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:327
		_go_fuzz_dep_.CoverTab[89487]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:327
		// _ = "end of CoverTab[89487]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:327
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:327
	// _ = "end of CoverTab[89484]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:327
	_go_fuzz_dep_.CoverTab[89485]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:328
	// _ = "end of CoverTab[89485]"
}

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:329
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bitreader.go:329
var _ = _go_fuzz_dep_.CoverTab
