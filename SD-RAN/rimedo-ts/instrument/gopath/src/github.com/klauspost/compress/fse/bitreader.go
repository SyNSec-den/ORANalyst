// Copyright 2018 Klaus Post. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Based on work Copyright (c) 2013, Yann Collet, released under BSD License.

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitreader.go:6
package fse

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitreader.go:6
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitreader.go:6
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitreader.go:6
)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitreader.go:6
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitreader.go:6
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitreader.go:6
)

import (
	"encoding/binary"
	"errors"
	"io"
)

// bitReader reads a bitstream in reverse.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitreader.go:14
// The last set bit indicates the start of the stream and is used
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitreader.go:14
// for aligning the input.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitreader.go:17
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
	b.bitsRead += 8 - uint8(highBits(uint32(v)))
	return nil
}

// getBits will return n bits. n can be 0.
func (b *bitReader) getBits(n uint8) uint16 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitreader.go:49
	_go_fuzz_dep_.CoverTab[88957]++
												if n == 0 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitreader.go:50
		_go_fuzz_dep_.CoverTab[88959]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitreader.go:50
		return b.bitsRead >= 64
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitreader.go:50
		// _ = "end of CoverTab[88959]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitreader.go:50
	}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitreader.go:50
		_go_fuzz_dep_.CoverTab[88960]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitreader.go:51
		// _ = "end of CoverTab[88960]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitreader.go:52
		_go_fuzz_dep_.CoverTab[88961]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitreader.go:52
		// _ = "end of CoverTab[88961]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitreader.go:52
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitreader.go:52
	// _ = "end of CoverTab[88957]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitreader.go:52
	_go_fuzz_dep_.CoverTab[88958]++
												return b.getBitsFast(n)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitreader.go:53
	// _ = "end of CoverTab[88958]"
}

// getBitsFast requires that at least one bit is requested every time.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitreader.go:56
// There are no checks if the buffer is filled.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitreader.go:58
func (b *bitReader) getBitsFast(n uint8) uint16 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitreader.go:58
	_go_fuzz_dep_.CoverTab[88962]++
												const regMask = 64 - 1
												v := uint16((b.value << (b.bitsRead & regMask)) >> ((regMask + 1 - n) & regMask))
												b.bitsRead += n
												return v
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitreader.go:62
	// _ = "end of CoverTab[88962]"
}

// fillFast() will make sure at least 32 bits are available.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitreader.go:65
// There must be at least 4 bytes available.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitreader.go:67
func (b *bitReader) fillFast() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitreader.go:67
	_go_fuzz_dep_.CoverTab[88963]++
												if b.bitsRead < 32 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitreader.go:68
		_go_fuzz_dep_.CoverTab[88965]++
													return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitreader.go:69
		// _ = "end of CoverTab[88965]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitreader.go:70
		_go_fuzz_dep_.CoverTab[88966]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitreader.go:70
		// _ = "end of CoverTab[88966]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitreader.go:70
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitreader.go:70
	// _ = "end of CoverTab[88963]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitreader.go:70
	_go_fuzz_dep_.CoverTab[88964]++

												v := b.in[b.off-4:]
												v = v[:4]
												low := (uint32(v[0])) | (uint32(v[1]) << 8) | (uint32(v[2]) << 16) | (uint32(v[3]) << 24)
												b.value = (b.value << 32) | uint64(low)
												b.bitsRead -= 32
												b.off -= 4
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitreader.go:77
	// _ = "end of CoverTab[88964]"
}

// fill() will make sure at least 32 bits are available.
func (b *bitReader) fill() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitreader.go:81
	_go_fuzz_dep_.CoverTab[88967]++
												if b.bitsRead < 32 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitreader.go:82
		_go_fuzz_dep_.CoverTab[88970]++
													return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitreader.go:83
		// _ = "end of CoverTab[88970]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitreader.go:84
		_go_fuzz_dep_.CoverTab[88971]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitreader.go:84
		// _ = "end of CoverTab[88971]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitreader.go:84
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitreader.go:84
	// _ = "end of CoverTab[88967]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitreader.go:84
	_go_fuzz_dep_.CoverTab[88968]++
												if b.off > 4 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitreader.go:85
		_go_fuzz_dep_.CoverTab[88972]++
													v := b.in[b.off-4:]
													v = v[:4]
													low := (uint32(v[0])) | (uint32(v[1]) << 8) | (uint32(v[2]) << 16) | (uint32(v[3]) << 24)
													b.value = (b.value << 32) | uint64(low)
													b.bitsRead -= 32
													b.off -= 4
													return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitreader.go:92
		// _ = "end of CoverTab[88972]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitreader.go:93
		_go_fuzz_dep_.CoverTab[88973]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitreader.go:93
		// _ = "end of CoverTab[88973]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitreader.go:93
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitreader.go:93
	// _ = "end of CoverTab[88968]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitreader.go:93
	_go_fuzz_dep_.CoverTab[88969]++
												for b.off > 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitreader.go:94
		_go_fuzz_dep_.CoverTab[88974]++
													b.value = (b.value << 8) | uint64(b.in[b.off-1])
													b.bitsRead -= 8
													b.off--
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitreader.go:97
		// _ = "end of CoverTab[88974]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitreader.go:98
	// _ = "end of CoverTab[88969]"
}

// fillFastStart() assumes the bitreader is empty and there is at least 8 bytes to read.
func (b *bitReader) fillFastStart() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitreader.go:102
	_go_fuzz_dep_.CoverTab[88975]++

												b.value = binary.LittleEndian.Uint64(b.in[b.off-8:])
												b.bitsRead = 0
												b.off -= 8
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitreader.go:106
	// _ = "end of CoverTab[88975]"
}

// finished returns true if all bits have been read from the bit stream.
func (b *bitReader) finished() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitreader.go:110
	_go_fuzz_dep_.CoverTab[88976]++
												return b.bitsRead >= 64 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitreader.go:111
		_go_fuzz_dep_.CoverTab[88977]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitreader.go:111
		return b.off == 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitreader.go:111
		// _ = "end of CoverTab[88977]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitreader.go:111
	}()
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitreader.go:111
	// _ = "end of CoverTab[88976]"
}

// close the bitstream and returns an error if out-of-buffer reads occurred.
func (b *bitReader) close() error {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitreader.go:115
	_go_fuzz_dep_.CoverTab[88978]++

												b.in = nil
												if b.bitsRead > 64 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitreader.go:118
		_go_fuzz_dep_.CoverTab[88980]++
													return io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitreader.go:119
		// _ = "end of CoverTab[88980]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitreader.go:120
		_go_fuzz_dep_.CoverTab[88981]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitreader.go:120
		// _ = "end of CoverTab[88981]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitreader.go:120
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitreader.go:120
	// _ = "end of CoverTab[88978]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitreader.go:120
	_go_fuzz_dep_.CoverTab[88979]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitreader.go:121
	// _ = "end of CoverTab[88979]"
}

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitreader.go:122
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/fse/bitreader.go:122
var _ = _go_fuzz_dep_.CoverTab
