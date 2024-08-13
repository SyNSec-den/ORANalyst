// Copyright 2019+ Klaus Post. All rights reserved.
// License information can be found in the LICENSE file.
// Based on work by Yann Collet, released under BSD License.

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:5
package zstd

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:5
)

import (
	"encoding/binary"
	"errors"
	"io"
	"math/bits"
)

// bitReader reads a bitstream in reverse.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:14
// The last set bit indicates the start of the stream and is used
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:14
// for aligning the input.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:17
type bitReader struct {
	in		[]byte
	off		uint	// next byte to read is at in[off - 1]
	value		uint64	// Maybe use [16]byte, but shifting is awkward.
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
func (b *bitReader) getBits(n uint8) int {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:49
	_go_fuzz_dep_.CoverTab[90841]++
												if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:50
		_go_fuzz_dep_.CoverTab[90843]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:51
		// _ = "end of CoverTab[90843]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:52
		_go_fuzz_dep_.CoverTab[90844]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:52
		// _ = "end of CoverTab[90844]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:52
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:52
	// _ = "end of CoverTab[90841]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:52
	_go_fuzz_dep_.CoverTab[90842]++
												return int(b.get32BitsFast(n))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:53
	// _ = "end of CoverTab[90842]"
}

// get32BitsFast requires that at least one bit is requested every time.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:56
// There are no checks if the buffer is filled.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:58
func (b *bitReader) get32BitsFast(n uint8) uint32 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:58
	_go_fuzz_dep_.CoverTab[90845]++
												const regMask = 64 - 1
												v := uint32((b.value << (b.bitsRead & regMask)) >> ((regMask + 1 - n) & regMask))
												b.bitsRead += n
												return v
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:62
	// _ = "end of CoverTab[90845]"
}

func (b *bitReader) get16BitsFast(n uint8) uint16 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:65
	_go_fuzz_dep_.CoverTab[90846]++
												const regMask = 64 - 1
												v := uint16((b.value << (b.bitsRead & regMask)) >> ((regMask + 1 - n) & regMask))
												b.bitsRead += n
												return v
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:69
	// _ = "end of CoverTab[90846]"
}

// fillFast() will make sure at least 32 bits are available.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:72
// There must be at least 4 bytes available.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:74
func (b *bitReader) fillFast() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:74
	_go_fuzz_dep_.CoverTab[90847]++
												if b.bitsRead < 32 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:75
		_go_fuzz_dep_.CoverTab[90849]++
													return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:76
		// _ = "end of CoverTab[90849]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:77
		_go_fuzz_dep_.CoverTab[90850]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:77
		// _ = "end of CoverTab[90850]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:77
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:77
	// _ = "end of CoverTab[90847]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:77
	_go_fuzz_dep_.CoverTab[90848]++

												v := b.in[b.off-4:]
												v = v[:4]
												low := (uint32(v[0])) | (uint32(v[1]) << 8) | (uint32(v[2]) << 16) | (uint32(v[3]) << 24)
												b.value = (b.value << 32) | uint64(low)
												b.bitsRead -= 32
												b.off -= 4
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:84
	// _ = "end of CoverTab[90848]"
}

// fillFastStart() assumes the bitreader is empty and there is at least 8 bytes to read.
func (b *bitReader) fillFastStart() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:88
	_go_fuzz_dep_.CoverTab[90851]++

												b.value = binary.LittleEndian.Uint64(b.in[b.off-8:])
												b.bitsRead = 0
												b.off -= 8
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:92
	// _ = "end of CoverTab[90851]"
}

// fill() will make sure at least 32 bits are available.
func (b *bitReader) fill() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:96
	_go_fuzz_dep_.CoverTab[90852]++
												if b.bitsRead < 32 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:97
		_go_fuzz_dep_.CoverTab[90855]++
													return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:98
		// _ = "end of CoverTab[90855]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:99
		_go_fuzz_dep_.CoverTab[90856]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:99
		// _ = "end of CoverTab[90856]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:99
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:99
	// _ = "end of CoverTab[90852]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:99
	_go_fuzz_dep_.CoverTab[90853]++
												if b.off >= 4 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:100
		_go_fuzz_dep_.CoverTab[90857]++
													v := b.in[b.off-4:]
													v = v[:4]
													low := (uint32(v[0])) | (uint32(v[1]) << 8) | (uint32(v[2]) << 16) | (uint32(v[3]) << 24)
													b.value = (b.value << 32) | uint64(low)
													b.bitsRead -= 32
													b.off -= 4
													return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:107
		// _ = "end of CoverTab[90857]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:108
		_go_fuzz_dep_.CoverTab[90858]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:108
		// _ = "end of CoverTab[90858]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:108
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:108
	// _ = "end of CoverTab[90853]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:108
	_go_fuzz_dep_.CoverTab[90854]++
												for b.off > 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:109
		_go_fuzz_dep_.CoverTab[90859]++
													b.value = (b.value << 8) | uint64(b.in[b.off-1])
													b.bitsRead -= 8
													b.off--
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:112
		// _ = "end of CoverTab[90859]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:113
	// _ = "end of CoverTab[90854]"
}

// finished returns true if all bits have been read from the bit stream.
func (b *bitReader) finished() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:117
	_go_fuzz_dep_.CoverTab[90860]++
												return b.off == 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:118
		_go_fuzz_dep_.CoverTab[90861]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:118
		return b.bitsRead >= 64
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:118
		// _ = "end of CoverTab[90861]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:118
	}()
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:118
	// _ = "end of CoverTab[90860]"
}

// overread returns true if more bits have been requested than is on the stream.
func (b *bitReader) overread() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:122
	_go_fuzz_dep_.CoverTab[90862]++
												return b.bitsRead > 64
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:123
	// _ = "end of CoverTab[90862]"
}

// remain returns the number of bits remaining.
func (b *bitReader) remain() uint {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:127
	_go_fuzz_dep_.CoverTab[90863]++
												return b.off*8 + 64 - uint(b.bitsRead)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:128
	// _ = "end of CoverTab[90863]"
}

// close the bitstream and returns an error if out-of-buffer reads occurred.
func (b *bitReader) close() error {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:132
	_go_fuzz_dep_.CoverTab[90864]++

												b.in = nil
												if b.bitsRead > 64 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:135
		_go_fuzz_dep_.CoverTab[90866]++
													return io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:136
		// _ = "end of CoverTab[90866]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:137
		_go_fuzz_dep_.CoverTab[90867]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:137
		// _ = "end of CoverTab[90867]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:137
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:137
	// _ = "end of CoverTab[90864]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:137
	_go_fuzz_dep_.CoverTab[90865]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:138
	// _ = "end of CoverTab[90865]"
}

func highBits(val uint32) (n uint32) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:141
	_go_fuzz_dep_.CoverTab[90868]++
												return uint32(bits.Len32(val) - 1)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:142
	// _ = "end of CoverTab[90868]"
}

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:143
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bitreader.go:143
var _ = _go_fuzz_dep_.CoverTab
