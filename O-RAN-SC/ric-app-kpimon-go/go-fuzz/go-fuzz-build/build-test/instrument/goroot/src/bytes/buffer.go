// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /snap/go/10455/src/bytes/buffer.go:5
package bytes

//line /snap/go/10455/src/bytes/buffer.go:5
import (
//line /snap/go/10455/src/bytes/buffer.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /snap/go/10455/src/bytes/buffer.go:5
)
//line /snap/go/10455/src/bytes/buffer.go:5
import (
//line /snap/go/10455/src/bytes/buffer.go:5
	_atomic_ "sync/atomic"
//line /snap/go/10455/src/bytes/buffer.go:5
)

//line /snap/go/10455/src/bytes/buffer.go:9
import (
	"errors"
	"io"
	"unicode/utf8"
)

// smallBufferSize is an initial allocation minimal capacity.
const smallBufferSize = 64

// A Buffer is a variable-sized buffer of bytes with Read and Write methods.
//line /snap/go/10455/src/bytes/buffer.go:18
// The zero value for Buffer is an empty buffer ready to use.
//line /snap/go/10455/src/bytes/buffer.go:20
type Buffer struct {
	buf		[]byte	// contents are the bytes buf[off : len(buf)]
	off		int	// read at &buf[off], write at &buf[len(buf)]
	lastRead	readOp	// last read operation, so that Unread* can work correctly.
}

// The readOp constants describe the last action performed on
//line /snap/go/10455/src/bytes/buffer.go:26
// the buffer, so that UnreadRune and UnreadByte can check for
//line /snap/go/10455/src/bytes/buffer.go:26
// invalid usage. opReadRuneX constants are chosen such that
//line /snap/go/10455/src/bytes/buffer.go:26
// converted to int they correspond to the rune size that was read.
//line /snap/go/10455/src/bytes/buffer.go:30
type readOp int8

// Don't use iota for these, as the values need to correspond with the
//line /snap/go/10455/src/bytes/buffer.go:32
// names and comments, which is easier to see when being explicit.
//line /snap/go/10455/src/bytes/buffer.go:34
const (
	opRead		readOp	= -1	// Any other read operation.
	opInvalid	readOp	= 0	// Non-read operation.
	opReadRune1	readOp	= 1	// Read rune of size 1.
	opReadRune2	readOp	= 2	// Read rune of size 2.
	opReadRune3	readOp	= 3	// Read rune of size 3.
	opReadRune4	readOp	= 4	// Read rune of size 4.
)

// ErrTooLarge is passed to panic if memory cannot be allocated to store data in a buffer.
var ErrTooLarge = errors.New("bytes.Buffer: too large")
var errNegativeRead = errors.New("bytes.Buffer: reader returned negative count from Read")

const maxInt = int(^uint(0) >> 1)

// Bytes returns a slice of length b.Len() holding the unread portion of the buffer.
//line /snap/go/10455/src/bytes/buffer.go:49
// The slice is valid for use only until the next buffer modification (that is,
//line /snap/go/10455/src/bytes/buffer.go:49
// only until the next call to a method like Read, Write, Reset, or Truncate).
//line /snap/go/10455/src/bytes/buffer.go:49
// The slice aliases the buffer content at least until the next buffer modification,
//line /snap/go/10455/src/bytes/buffer.go:49
// so immediate changes to the slice will affect the result of future reads.
//line /snap/go/10455/src/bytes/buffer.go:54
func (b *Buffer) Bytes() []byte {
//line /snap/go/10455/src/bytes/buffer.go:54
	_go_fuzz_dep_.CoverTab[1]++
//line /snap/go/10455/src/bytes/buffer.go:54
	return b.buf[b.off:]
//line /snap/go/10455/src/bytes/buffer.go:54
	// _ = "end of CoverTab[1]"
//line /snap/go/10455/src/bytes/buffer.go:54
}

// AvailableBuffer returns an empty buffer with b.Available() capacity.
//line /snap/go/10455/src/bytes/buffer.go:56
// This buffer is intended to be appended to and
//line /snap/go/10455/src/bytes/buffer.go:56
// passed to an immediately succeeding Write call.
//line /snap/go/10455/src/bytes/buffer.go:56
// The buffer is only valid until the next write operation on b.
//line /snap/go/10455/src/bytes/buffer.go:60
func (b *Buffer) AvailableBuffer() []byte {
//line /snap/go/10455/src/bytes/buffer.go:60
	_go_fuzz_dep_.CoverTab[2]++
//line /snap/go/10455/src/bytes/buffer.go:60
	return b.buf[len(b.buf):]
//line /snap/go/10455/src/bytes/buffer.go:60
	// _ = "end of CoverTab[2]"
//line /snap/go/10455/src/bytes/buffer.go:60
}

// String returns the contents of the unread portion of the buffer
//line /snap/go/10455/src/bytes/buffer.go:62
// as a string. If the Buffer is a nil pointer, it returns "<nil>".
//line /snap/go/10455/src/bytes/buffer.go:62
//
//line /snap/go/10455/src/bytes/buffer.go:62
// To build strings more efficiently, see the strings.Builder type.
//line /snap/go/10455/src/bytes/buffer.go:66
func (b *Buffer) String() string {
//line /snap/go/10455/src/bytes/buffer.go:66
	_go_fuzz_dep_.CoverTab[3]++
						if b == nil {
//line /snap/go/10455/src/bytes/buffer.go:67
		_go_fuzz_dep_.CoverTab[524289]++
//line /snap/go/10455/src/bytes/buffer.go:67
		_go_fuzz_dep_.CoverTab[5]++

							return "<nil>"
//line /snap/go/10455/src/bytes/buffer.go:69
		// _ = "end of CoverTab[5]"
	} else {
//line /snap/go/10455/src/bytes/buffer.go:70
		_go_fuzz_dep_.CoverTab[524290]++
//line /snap/go/10455/src/bytes/buffer.go:70
		_go_fuzz_dep_.CoverTab[6]++
//line /snap/go/10455/src/bytes/buffer.go:70
		// _ = "end of CoverTab[6]"
//line /snap/go/10455/src/bytes/buffer.go:70
	}
//line /snap/go/10455/src/bytes/buffer.go:70
	// _ = "end of CoverTab[3]"
//line /snap/go/10455/src/bytes/buffer.go:70
	_go_fuzz_dep_.CoverTab[4]++
						return string(b.buf[b.off:])
//line /snap/go/10455/src/bytes/buffer.go:71
	// _ = "end of CoverTab[4]"
}

// empty reports whether the unread portion of the buffer is empty.
func (b *Buffer) empty() bool {
//line /snap/go/10455/src/bytes/buffer.go:75
	_go_fuzz_dep_.CoverTab[7]++
//line /snap/go/10455/src/bytes/buffer.go:75
	return len(b.buf) <= b.off
//line /snap/go/10455/src/bytes/buffer.go:75
	// _ = "end of CoverTab[7]"
//line /snap/go/10455/src/bytes/buffer.go:75
}

// Len returns the number of bytes of the unread portion of the buffer;
//line /snap/go/10455/src/bytes/buffer.go:77
// b.Len() == len(b.Bytes()).
//line /snap/go/10455/src/bytes/buffer.go:79
func (b *Buffer) Len() int {
//line /snap/go/10455/src/bytes/buffer.go:79
	_go_fuzz_dep_.CoverTab[8]++
//line /snap/go/10455/src/bytes/buffer.go:79
	return len(b.buf) - b.off
//line /snap/go/10455/src/bytes/buffer.go:79
	// _ = "end of CoverTab[8]"
//line /snap/go/10455/src/bytes/buffer.go:79
}

// Cap returns the capacity of the buffer's underlying byte slice, that is, the
//line /snap/go/10455/src/bytes/buffer.go:81
// total space allocated for the buffer's data.
//line /snap/go/10455/src/bytes/buffer.go:83
func (b *Buffer) Cap() int {
//line /snap/go/10455/src/bytes/buffer.go:83
	_go_fuzz_dep_.CoverTab[9]++
//line /snap/go/10455/src/bytes/buffer.go:83
	return cap(b.buf)
//line /snap/go/10455/src/bytes/buffer.go:83
	// _ = "end of CoverTab[9]"
//line /snap/go/10455/src/bytes/buffer.go:83
}

// Available returns how many bytes are unused in the buffer.
func (b *Buffer) Available() int {
//line /snap/go/10455/src/bytes/buffer.go:86
	_go_fuzz_dep_.CoverTab[10]++
//line /snap/go/10455/src/bytes/buffer.go:86
	return cap(b.buf) - len(b.buf)
//line /snap/go/10455/src/bytes/buffer.go:86
	// _ = "end of CoverTab[10]"
//line /snap/go/10455/src/bytes/buffer.go:86
}

// Truncate discards all but the first n unread bytes from the buffer
//line /snap/go/10455/src/bytes/buffer.go:88
// but continues to use the same allocated storage.
//line /snap/go/10455/src/bytes/buffer.go:88
// It panics if n is negative or greater than the length of the buffer.
//line /snap/go/10455/src/bytes/buffer.go:91
func (b *Buffer) Truncate(n int) {
//line /snap/go/10455/src/bytes/buffer.go:91
	_go_fuzz_dep_.CoverTab[11]++
						if n == 0 {
//line /snap/go/10455/src/bytes/buffer.go:92
		_go_fuzz_dep_.CoverTab[524291]++
//line /snap/go/10455/src/bytes/buffer.go:92
		_go_fuzz_dep_.CoverTab[14]++
							b.Reset()
							return
//line /snap/go/10455/src/bytes/buffer.go:94
		// _ = "end of CoverTab[14]"
	} else {
//line /snap/go/10455/src/bytes/buffer.go:95
		_go_fuzz_dep_.CoverTab[524292]++
//line /snap/go/10455/src/bytes/buffer.go:95
		_go_fuzz_dep_.CoverTab[15]++
//line /snap/go/10455/src/bytes/buffer.go:95
		// _ = "end of CoverTab[15]"
//line /snap/go/10455/src/bytes/buffer.go:95
	}
//line /snap/go/10455/src/bytes/buffer.go:95
	// _ = "end of CoverTab[11]"
//line /snap/go/10455/src/bytes/buffer.go:95
	_go_fuzz_dep_.CoverTab[12]++
						b.lastRead = opInvalid
						if n < 0 || func() bool {
//line /snap/go/10455/src/bytes/buffer.go:97
		_go_fuzz_dep_.CoverTab[16]++
//line /snap/go/10455/src/bytes/buffer.go:97
		return n > b.Len()
//line /snap/go/10455/src/bytes/buffer.go:97
		// _ = "end of CoverTab[16]"
//line /snap/go/10455/src/bytes/buffer.go:97
	}() {
//line /snap/go/10455/src/bytes/buffer.go:97
		_go_fuzz_dep_.CoverTab[524293]++
//line /snap/go/10455/src/bytes/buffer.go:97
		_go_fuzz_dep_.CoverTab[17]++
							panic("bytes.Buffer: truncation out of range")
//line /snap/go/10455/src/bytes/buffer.go:98
		// _ = "end of CoverTab[17]"
	} else {
//line /snap/go/10455/src/bytes/buffer.go:99
		_go_fuzz_dep_.CoverTab[524294]++
//line /snap/go/10455/src/bytes/buffer.go:99
		_go_fuzz_dep_.CoverTab[18]++
//line /snap/go/10455/src/bytes/buffer.go:99
		// _ = "end of CoverTab[18]"
//line /snap/go/10455/src/bytes/buffer.go:99
	}
//line /snap/go/10455/src/bytes/buffer.go:99
	// _ = "end of CoverTab[12]"
//line /snap/go/10455/src/bytes/buffer.go:99
	_go_fuzz_dep_.CoverTab[13]++
						b.buf = b.buf[:b.off+n]
//line /snap/go/10455/src/bytes/buffer.go:100
	// _ = "end of CoverTab[13]"
}

// Reset resets the buffer to be empty,
//line /snap/go/10455/src/bytes/buffer.go:103
// but it retains the underlying storage for use by future writes.
//line /snap/go/10455/src/bytes/buffer.go:103
// Reset is the same as Truncate(0).
//line /snap/go/10455/src/bytes/buffer.go:106
func (b *Buffer) Reset() {
//line /snap/go/10455/src/bytes/buffer.go:106
	_go_fuzz_dep_.CoverTab[19]++
						b.buf = b.buf[:0]
						b.off = 0
						b.lastRead = opInvalid
//line /snap/go/10455/src/bytes/buffer.go:109
	// _ = "end of CoverTab[19]"
}

// tryGrowByReslice is an inlineable version of grow for the fast-case where the
//line /snap/go/10455/src/bytes/buffer.go:112
// internal buffer only needs to be resliced.
//line /snap/go/10455/src/bytes/buffer.go:112
// It returns the index where bytes should be written and whether it succeeded.
//line /snap/go/10455/src/bytes/buffer.go:115
func (b *Buffer) tryGrowByReslice(n int) (int, bool) {
//line /snap/go/10455/src/bytes/buffer.go:115
	_go_fuzz_dep_.CoverTab[20]++
						if l := len(b.buf); n <= cap(b.buf)-l {
//line /snap/go/10455/src/bytes/buffer.go:116
		_go_fuzz_dep_.CoverTab[524295]++
//line /snap/go/10455/src/bytes/buffer.go:116
		_go_fuzz_dep_.CoverTab[22]++
							b.buf = b.buf[:l+n]
							return l, true
//line /snap/go/10455/src/bytes/buffer.go:118
		// _ = "end of CoverTab[22]"
	} else {
//line /snap/go/10455/src/bytes/buffer.go:119
		_go_fuzz_dep_.CoverTab[524296]++
//line /snap/go/10455/src/bytes/buffer.go:119
		_go_fuzz_dep_.CoverTab[23]++
//line /snap/go/10455/src/bytes/buffer.go:119
		// _ = "end of CoverTab[23]"
//line /snap/go/10455/src/bytes/buffer.go:119
	}
//line /snap/go/10455/src/bytes/buffer.go:119
	// _ = "end of CoverTab[20]"
//line /snap/go/10455/src/bytes/buffer.go:119
	_go_fuzz_dep_.CoverTab[21]++
						return 0, false
//line /snap/go/10455/src/bytes/buffer.go:120
	// _ = "end of CoverTab[21]"
}

// grow grows the buffer to guarantee space for n more bytes.
//line /snap/go/10455/src/bytes/buffer.go:123
// It returns the index where bytes should be written.
//line /snap/go/10455/src/bytes/buffer.go:123
// If the buffer can't grow it will panic with ErrTooLarge.
//line /snap/go/10455/src/bytes/buffer.go:126
func (b *Buffer) grow(n int) int {
//line /snap/go/10455/src/bytes/buffer.go:126
	_go_fuzz_dep_.CoverTab[24]++
						m := b.Len()

						if m == 0 && func() bool {
//line /snap/go/10455/src/bytes/buffer.go:129
		_go_fuzz_dep_.CoverTab[29]++
//line /snap/go/10455/src/bytes/buffer.go:129
		return b.off != 0
//line /snap/go/10455/src/bytes/buffer.go:129
		// _ = "end of CoverTab[29]"
//line /snap/go/10455/src/bytes/buffer.go:129
	}() {
//line /snap/go/10455/src/bytes/buffer.go:129
		_go_fuzz_dep_.CoverTab[524297]++
//line /snap/go/10455/src/bytes/buffer.go:129
		_go_fuzz_dep_.CoverTab[30]++
							b.Reset()
//line /snap/go/10455/src/bytes/buffer.go:130
		// _ = "end of CoverTab[30]"
	} else {
//line /snap/go/10455/src/bytes/buffer.go:131
		_go_fuzz_dep_.CoverTab[524298]++
//line /snap/go/10455/src/bytes/buffer.go:131
		_go_fuzz_dep_.CoverTab[31]++
//line /snap/go/10455/src/bytes/buffer.go:131
		// _ = "end of CoverTab[31]"
//line /snap/go/10455/src/bytes/buffer.go:131
	}
//line /snap/go/10455/src/bytes/buffer.go:131
	// _ = "end of CoverTab[24]"
//line /snap/go/10455/src/bytes/buffer.go:131
	_go_fuzz_dep_.CoverTab[25]++

						if i, ok := b.tryGrowByReslice(n); ok {
//line /snap/go/10455/src/bytes/buffer.go:133
		_go_fuzz_dep_.CoverTab[524299]++
//line /snap/go/10455/src/bytes/buffer.go:133
		_go_fuzz_dep_.CoverTab[32]++
							return i
//line /snap/go/10455/src/bytes/buffer.go:134
		// _ = "end of CoverTab[32]"
	} else {
//line /snap/go/10455/src/bytes/buffer.go:135
		_go_fuzz_dep_.CoverTab[524300]++
//line /snap/go/10455/src/bytes/buffer.go:135
		_go_fuzz_dep_.CoverTab[33]++
//line /snap/go/10455/src/bytes/buffer.go:135
		// _ = "end of CoverTab[33]"
//line /snap/go/10455/src/bytes/buffer.go:135
	}
//line /snap/go/10455/src/bytes/buffer.go:135
	// _ = "end of CoverTab[25]"
//line /snap/go/10455/src/bytes/buffer.go:135
	_go_fuzz_dep_.CoverTab[26]++
						if b.buf == nil && func() bool {
//line /snap/go/10455/src/bytes/buffer.go:136
		_go_fuzz_dep_.CoverTab[34]++
//line /snap/go/10455/src/bytes/buffer.go:136
		return n <= smallBufferSize
//line /snap/go/10455/src/bytes/buffer.go:136
		// _ = "end of CoverTab[34]"
//line /snap/go/10455/src/bytes/buffer.go:136
	}() {
//line /snap/go/10455/src/bytes/buffer.go:136
		_go_fuzz_dep_.CoverTab[524301]++
//line /snap/go/10455/src/bytes/buffer.go:136
		_go_fuzz_dep_.CoverTab[35]++
							b.buf = make([]byte, n, smallBufferSize)
							return 0
//line /snap/go/10455/src/bytes/buffer.go:138
		// _ = "end of CoverTab[35]"
	} else {
//line /snap/go/10455/src/bytes/buffer.go:139
		_go_fuzz_dep_.CoverTab[524302]++
//line /snap/go/10455/src/bytes/buffer.go:139
		_go_fuzz_dep_.CoverTab[36]++
//line /snap/go/10455/src/bytes/buffer.go:139
		// _ = "end of CoverTab[36]"
//line /snap/go/10455/src/bytes/buffer.go:139
	}
//line /snap/go/10455/src/bytes/buffer.go:139
	// _ = "end of CoverTab[26]"
//line /snap/go/10455/src/bytes/buffer.go:139
	_go_fuzz_dep_.CoverTab[27]++
						c := cap(b.buf)
						if n <= c/2-m {
//line /snap/go/10455/src/bytes/buffer.go:141
		_go_fuzz_dep_.CoverTab[524303]++
//line /snap/go/10455/src/bytes/buffer.go:141
		_go_fuzz_dep_.CoverTab[37]++

//line /snap/go/10455/src/bytes/buffer.go:146
		copy(b.buf, b.buf[b.off:])
//line /snap/go/10455/src/bytes/buffer.go:146
		// _ = "end of CoverTab[37]"
	} else {
//line /snap/go/10455/src/bytes/buffer.go:147
		_go_fuzz_dep_.CoverTab[524304]++
//line /snap/go/10455/src/bytes/buffer.go:147
		_go_fuzz_dep_.CoverTab[38]++
//line /snap/go/10455/src/bytes/buffer.go:147
		if c > maxInt-c-n {
//line /snap/go/10455/src/bytes/buffer.go:147
			_go_fuzz_dep_.CoverTab[524305]++
//line /snap/go/10455/src/bytes/buffer.go:147
			_go_fuzz_dep_.CoverTab[39]++
								panic(ErrTooLarge)
//line /snap/go/10455/src/bytes/buffer.go:148
			// _ = "end of CoverTab[39]"
		} else {
//line /snap/go/10455/src/bytes/buffer.go:149
			_go_fuzz_dep_.CoverTab[524306]++
//line /snap/go/10455/src/bytes/buffer.go:149
			_go_fuzz_dep_.CoverTab[40]++

								b.buf = growSlice(b.buf[b.off:], b.off+n)
//line /snap/go/10455/src/bytes/buffer.go:151
			// _ = "end of CoverTab[40]"
		}
//line /snap/go/10455/src/bytes/buffer.go:152
		// _ = "end of CoverTab[38]"
//line /snap/go/10455/src/bytes/buffer.go:152
	}
//line /snap/go/10455/src/bytes/buffer.go:152
	// _ = "end of CoverTab[27]"
//line /snap/go/10455/src/bytes/buffer.go:152
	_go_fuzz_dep_.CoverTab[28]++

						b.off = 0
						b.buf = b.buf[:m+n]
						return m
//line /snap/go/10455/src/bytes/buffer.go:156
	// _ = "end of CoverTab[28]"
}

// Grow grows the buffer's capacity, if necessary, to guarantee space for
//line /snap/go/10455/src/bytes/buffer.go:159
// another n bytes. After Grow(n), at least n bytes can be written to the
//line /snap/go/10455/src/bytes/buffer.go:159
// buffer without another allocation.
//line /snap/go/10455/src/bytes/buffer.go:159
// If n is negative, Grow will panic.
//line /snap/go/10455/src/bytes/buffer.go:159
// If the buffer can't grow it will panic with ErrTooLarge.
//line /snap/go/10455/src/bytes/buffer.go:164
func (b *Buffer) Grow(n int) {
//line /snap/go/10455/src/bytes/buffer.go:164
	_go_fuzz_dep_.CoverTab[41]++
						if n < 0 {
//line /snap/go/10455/src/bytes/buffer.go:165
		_go_fuzz_dep_.CoverTab[524307]++
//line /snap/go/10455/src/bytes/buffer.go:165
		_go_fuzz_dep_.CoverTab[43]++
							panic("bytes.Buffer.Grow: negative count")
//line /snap/go/10455/src/bytes/buffer.go:166
		// _ = "end of CoverTab[43]"
	} else {
//line /snap/go/10455/src/bytes/buffer.go:167
		_go_fuzz_dep_.CoverTab[524308]++
//line /snap/go/10455/src/bytes/buffer.go:167
		_go_fuzz_dep_.CoverTab[44]++
//line /snap/go/10455/src/bytes/buffer.go:167
		// _ = "end of CoverTab[44]"
//line /snap/go/10455/src/bytes/buffer.go:167
	}
//line /snap/go/10455/src/bytes/buffer.go:167
	// _ = "end of CoverTab[41]"
//line /snap/go/10455/src/bytes/buffer.go:167
	_go_fuzz_dep_.CoverTab[42]++
						m := b.grow(n)
						b.buf = b.buf[:m]
//line /snap/go/10455/src/bytes/buffer.go:169
	// _ = "end of CoverTab[42]"
}

// Write appends the contents of p to the buffer, growing the buffer as
//line /snap/go/10455/src/bytes/buffer.go:172
// needed. The return value n is the length of p; err is always nil. If the
//line /snap/go/10455/src/bytes/buffer.go:172
// buffer becomes too large, Write will panic with ErrTooLarge.
//line /snap/go/10455/src/bytes/buffer.go:175
func (b *Buffer) Write(p []byte) (n int, err error) {
//line /snap/go/10455/src/bytes/buffer.go:175
	_go_fuzz_dep_.CoverTab[45]++
						b.lastRead = opInvalid
						m, ok := b.tryGrowByReslice(len(p))
						if !ok {
//line /snap/go/10455/src/bytes/buffer.go:178
		_go_fuzz_dep_.CoverTab[524309]++
//line /snap/go/10455/src/bytes/buffer.go:178
		_go_fuzz_dep_.CoverTab[47]++
							m = b.grow(len(p))
//line /snap/go/10455/src/bytes/buffer.go:179
		// _ = "end of CoverTab[47]"
	} else {
//line /snap/go/10455/src/bytes/buffer.go:180
		_go_fuzz_dep_.CoverTab[524310]++
//line /snap/go/10455/src/bytes/buffer.go:180
		_go_fuzz_dep_.CoverTab[48]++
//line /snap/go/10455/src/bytes/buffer.go:180
		// _ = "end of CoverTab[48]"
//line /snap/go/10455/src/bytes/buffer.go:180
	}
//line /snap/go/10455/src/bytes/buffer.go:180
	// _ = "end of CoverTab[45]"
//line /snap/go/10455/src/bytes/buffer.go:180
	_go_fuzz_dep_.CoverTab[46]++
						return copy(b.buf[m:], p), nil
//line /snap/go/10455/src/bytes/buffer.go:181
	// _ = "end of CoverTab[46]"
}

// WriteString appends the contents of s to the buffer, growing the buffer as
//line /snap/go/10455/src/bytes/buffer.go:184
// needed. The return value n is the length of s; err is always nil. If the
//line /snap/go/10455/src/bytes/buffer.go:184
// buffer becomes too large, WriteString will panic with ErrTooLarge.
//line /snap/go/10455/src/bytes/buffer.go:187
func (b *Buffer) WriteString(s string) (n int, err error) {
//line /snap/go/10455/src/bytes/buffer.go:187
	_go_fuzz_dep_.CoverTab[49]++
						b.lastRead = opInvalid
						m, ok := b.tryGrowByReslice(len(s))
						if !ok {
//line /snap/go/10455/src/bytes/buffer.go:190
		_go_fuzz_dep_.CoverTab[524311]++
//line /snap/go/10455/src/bytes/buffer.go:190
		_go_fuzz_dep_.CoverTab[51]++
							m = b.grow(len(s))
//line /snap/go/10455/src/bytes/buffer.go:191
		// _ = "end of CoverTab[51]"
	} else {
//line /snap/go/10455/src/bytes/buffer.go:192
		_go_fuzz_dep_.CoverTab[524312]++
//line /snap/go/10455/src/bytes/buffer.go:192
		_go_fuzz_dep_.CoverTab[52]++
//line /snap/go/10455/src/bytes/buffer.go:192
		// _ = "end of CoverTab[52]"
//line /snap/go/10455/src/bytes/buffer.go:192
	}
//line /snap/go/10455/src/bytes/buffer.go:192
	// _ = "end of CoverTab[49]"
//line /snap/go/10455/src/bytes/buffer.go:192
	_go_fuzz_dep_.CoverTab[50]++
						return copy(b.buf[m:], s), nil
//line /snap/go/10455/src/bytes/buffer.go:193
	// _ = "end of CoverTab[50]"
}

// MinRead is the minimum slice size passed to a Read call by
//line /snap/go/10455/src/bytes/buffer.go:196
// Buffer.ReadFrom. As long as the Buffer has at least MinRead bytes beyond
//line /snap/go/10455/src/bytes/buffer.go:196
// what is required to hold the contents of r, ReadFrom will not grow the
//line /snap/go/10455/src/bytes/buffer.go:196
// underlying buffer.
//line /snap/go/10455/src/bytes/buffer.go:200
const MinRead = 512

// ReadFrom reads data from r until EOF and appends it to the buffer, growing
//line /snap/go/10455/src/bytes/buffer.go:202
// the buffer as needed. The return value n is the number of bytes read. Any
//line /snap/go/10455/src/bytes/buffer.go:202
// error except io.EOF encountered during the read is also returned. If the
//line /snap/go/10455/src/bytes/buffer.go:202
// buffer becomes too large, ReadFrom will panic with ErrTooLarge.
//line /snap/go/10455/src/bytes/buffer.go:206
func (b *Buffer) ReadFrom(r io.Reader) (n int64, err error) {
//line /snap/go/10455/src/bytes/buffer.go:206
	_go_fuzz_dep_.CoverTab[53]++
						b.lastRead = opInvalid
//line /snap/go/10455/src/bytes/buffer.go:207
	_go_fuzz_dep_.CoverTab[786433] = 0
						for {
//line /snap/go/10455/src/bytes/buffer.go:208
		if _go_fuzz_dep_.CoverTab[786433] == 0 {
//line /snap/go/10455/src/bytes/buffer.go:208
			_go_fuzz_dep_.CoverTab[524363]++
//line /snap/go/10455/src/bytes/buffer.go:208
		} else {
//line /snap/go/10455/src/bytes/buffer.go:208
			_go_fuzz_dep_.CoverTab[524364]++
//line /snap/go/10455/src/bytes/buffer.go:208
		}
//line /snap/go/10455/src/bytes/buffer.go:208
		_go_fuzz_dep_.CoverTab[786433] = 1
//line /snap/go/10455/src/bytes/buffer.go:208
		_go_fuzz_dep_.CoverTab[54]++
							i := b.grow(MinRead)
							b.buf = b.buf[:i]
							m, e := r.Read(b.buf[i:cap(b.buf)])
							if m < 0 {
//line /snap/go/10455/src/bytes/buffer.go:212
			_go_fuzz_dep_.CoverTab[524313]++
//line /snap/go/10455/src/bytes/buffer.go:212
			_go_fuzz_dep_.CoverTab[57]++
								panic(errNegativeRead)
//line /snap/go/10455/src/bytes/buffer.go:213
			// _ = "end of CoverTab[57]"
		} else {
//line /snap/go/10455/src/bytes/buffer.go:214
			_go_fuzz_dep_.CoverTab[524314]++
//line /snap/go/10455/src/bytes/buffer.go:214
			_go_fuzz_dep_.CoverTab[58]++
//line /snap/go/10455/src/bytes/buffer.go:214
			// _ = "end of CoverTab[58]"
//line /snap/go/10455/src/bytes/buffer.go:214
		}
//line /snap/go/10455/src/bytes/buffer.go:214
		// _ = "end of CoverTab[54]"
//line /snap/go/10455/src/bytes/buffer.go:214
		_go_fuzz_dep_.CoverTab[55]++

							b.buf = b.buf[:i+m]
							n += int64(m)
							if e == io.EOF {
//line /snap/go/10455/src/bytes/buffer.go:218
			_go_fuzz_dep_.CoverTab[524315]++
//line /snap/go/10455/src/bytes/buffer.go:218
			_go_fuzz_dep_.CoverTab[59]++
								return n, nil
//line /snap/go/10455/src/bytes/buffer.go:219
			// _ = "end of CoverTab[59]"
		} else {
//line /snap/go/10455/src/bytes/buffer.go:220
			_go_fuzz_dep_.CoverTab[524316]++
//line /snap/go/10455/src/bytes/buffer.go:220
			_go_fuzz_dep_.CoverTab[60]++
//line /snap/go/10455/src/bytes/buffer.go:220
			// _ = "end of CoverTab[60]"
//line /snap/go/10455/src/bytes/buffer.go:220
		}
//line /snap/go/10455/src/bytes/buffer.go:220
		// _ = "end of CoverTab[55]"
//line /snap/go/10455/src/bytes/buffer.go:220
		_go_fuzz_dep_.CoverTab[56]++
							if e != nil {
//line /snap/go/10455/src/bytes/buffer.go:221
			_go_fuzz_dep_.CoverTab[524317]++
//line /snap/go/10455/src/bytes/buffer.go:221
			_go_fuzz_dep_.CoverTab[61]++
								return n, e
//line /snap/go/10455/src/bytes/buffer.go:222
			// _ = "end of CoverTab[61]"
		} else {
//line /snap/go/10455/src/bytes/buffer.go:223
			_go_fuzz_dep_.CoverTab[524318]++
//line /snap/go/10455/src/bytes/buffer.go:223
			_go_fuzz_dep_.CoverTab[62]++
//line /snap/go/10455/src/bytes/buffer.go:223
			// _ = "end of CoverTab[62]"
//line /snap/go/10455/src/bytes/buffer.go:223
		}
//line /snap/go/10455/src/bytes/buffer.go:223
		// _ = "end of CoverTab[56]"
	}
//line /snap/go/10455/src/bytes/buffer.go:224
	// _ = "end of CoverTab[53]"
}

// growSlice grows b by n, preserving the original content of b.
//line /snap/go/10455/src/bytes/buffer.go:227
// If the allocation fails, it panics with ErrTooLarge.
//line /snap/go/10455/src/bytes/buffer.go:229
func growSlice(b []byte, n int) []byte {
//line /snap/go/10455/src/bytes/buffer.go:229
	_go_fuzz_dep_.CoverTab[63]++
						defer func() {
//line /snap/go/10455/src/bytes/buffer.go:230
		_go_fuzz_dep_.CoverTab[66]++
							if recover() != nil {
//line /snap/go/10455/src/bytes/buffer.go:231
			_go_fuzz_dep_.CoverTab[524319]++
//line /snap/go/10455/src/bytes/buffer.go:231
			_go_fuzz_dep_.CoverTab[67]++
								panic(ErrTooLarge)
//line /snap/go/10455/src/bytes/buffer.go:232
			// _ = "end of CoverTab[67]"
		} else {
//line /snap/go/10455/src/bytes/buffer.go:233
			_go_fuzz_dep_.CoverTab[524320]++
//line /snap/go/10455/src/bytes/buffer.go:233
			_go_fuzz_dep_.CoverTab[68]++
//line /snap/go/10455/src/bytes/buffer.go:233
			// _ = "end of CoverTab[68]"
//line /snap/go/10455/src/bytes/buffer.go:233
		}
//line /snap/go/10455/src/bytes/buffer.go:233
		// _ = "end of CoverTab[66]"
	}()
//line /snap/go/10455/src/bytes/buffer.go:234
	// _ = "end of CoverTab[63]"
//line /snap/go/10455/src/bytes/buffer.go:234
	_go_fuzz_dep_.CoverTab[64]++

//line /snap/go/10455/src/bytes/buffer.go:243
	c := len(b) + n
	if c < 2*cap(b) {
//line /snap/go/10455/src/bytes/buffer.go:244
		_go_fuzz_dep_.CoverTab[524321]++
//line /snap/go/10455/src/bytes/buffer.go:244
		_go_fuzz_dep_.CoverTab[69]++

//line /snap/go/10455/src/bytes/buffer.go:247
		c = 2 * cap(b)
//line /snap/go/10455/src/bytes/buffer.go:247
		// _ = "end of CoverTab[69]"
	} else {
//line /snap/go/10455/src/bytes/buffer.go:248
		_go_fuzz_dep_.CoverTab[524322]++
//line /snap/go/10455/src/bytes/buffer.go:248
		_go_fuzz_dep_.CoverTab[70]++
//line /snap/go/10455/src/bytes/buffer.go:248
		// _ = "end of CoverTab[70]"
//line /snap/go/10455/src/bytes/buffer.go:248
	}
//line /snap/go/10455/src/bytes/buffer.go:248
	// _ = "end of CoverTab[64]"
//line /snap/go/10455/src/bytes/buffer.go:248
	_go_fuzz_dep_.CoverTab[65]++
						b2 := append([]byte(nil), make([]byte, c)...)
						copy(b2, b)
						return b2[:len(b)]
//line /snap/go/10455/src/bytes/buffer.go:251
	// _ = "end of CoverTab[65]"
}

// WriteTo writes data to w until the buffer is drained or an error occurs.
//line /snap/go/10455/src/bytes/buffer.go:254
// The return value n is the number of bytes written; it always fits into an
//line /snap/go/10455/src/bytes/buffer.go:254
// int, but it is int64 to match the io.WriterTo interface. Any error
//line /snap/go/10455/src/bytes/buffer.go:254
// encountered during the write is also returned.
//line /snap/go/10455/src/bytes/buffer.go:258
func (b *Buffer) WriteTo(w io.Writer) (n int64, err error) {
//line /snap/go/10455/src/bytes/buffer.go:258
	_go_fuzz_dep_.CoverTab[71]++
						b.lastRead = opInvalid
						if nBytes := b.Len(); nBytes > 0 {
//line /snap/go/10455/src/bytes/buffer.go:260
		_go_fuzz_dep_.CoverTab[524323]++
//line /snap/go/10455/src/bytes/buffer.go:260
		_go_fuzz_dep_.CoverTab[73]++
							m, e := w.Write(b.buf[b.off:])
							if m > nBytes {
//line /snap/go/10455/src/bytes/buffer.go:262
			_go_fuzz_dep_.CoverTab[524325]++
//line /snap/go/10455/src/bytes/buffer.go:262
			_go_fuzz_dep_.CoverTab[76]++
								panic("bytes.Buffer.WriteTo: invalid Write count")
//line /snap/go/10455/src/bytes/buffer.go:263
			// _ = "end of CoverTab[76]"
		} else {
//line /snap/go/10455/src/bytes/buffer.go:264
			_go_fuzz_dep_.CoverTab[524326]++
//line /snap/go/10455/src/bytes/buffer.go:264
			_go_fuzz_dep_.CoverTab[77]++
//line /snap/go/10455/src/bytes/buffer.go:264
			// _ = "end of CoverTab[77]"
//line /snap/go/10455/src/bytes/buffer.go:264
		}
//line /snap/go/10455/src/bytes/buffer.go:264
		// _ = "end of CoverTab[73]"
//line /snap/go/10455/src/bytes/buffer.go:264
		_go_fuzz_dep_.CoverTab[74]++
							b.off += m
							n = int64(m)
							if e != nil {
//line /snap/go/10455/src/bytes/buffer.go:267
			_go_fuzz_dep_.CoverTab[524327]++
//line /snap/go/10455/src/bytes/buffer.go:267
			_go_fuzz_dep_.CoverTab[78]++
								return n, e
//line /snap/go/10455/src/bytes/buffer.go:268
			// _ = "end of CoverTab[78]"
		} else {
//line /snap/go/10455/src/bytes/buffer.go:269
			_go_fuzz_dep_.CoverTab[524328]++
//line /snap/go/10455/src/bytes/buffer.go:269
			_go_fuzz_dep_.CoverTab[79]++
//line /snap/go/10455/src/bytes/buffer.go:269
			// _ = "end of CoverTab[79]"
//line /snap/go/10455/src/bytes/buffer.go:269
		}
//line /snap/go/10455/src/bytes/buffer.go:269
		// _ = "end of CoverTab[74]"
//line /snap/go/10455/src/bytes/buffer.go:269
		_go_fuzz_dep_.CoverTab[75]++

//line /snap/go/10455/src/bytes/buffer.go:272
		if m != nBytes {
//line /snap/go/10455/src/bytes/buffer.go:272
			_go_fuzz_dep_.CoverTab[524329]++
//line /snap/go/10455/src/bytes/buffer.go:272
			_go_fuzz_dep_.CoverTab[80]++
								return n, io.ErrShortWrite
//line /snap/go/10455/src/bytes/buffer.go:273
			// _ = "end of CoverTab[80]"
		} else {
//line /snap/go/10455/src/bytes/buffer.go:274
			_go_fuzz_dep_.CoverTab[524330]++
//line /snap/go/10455/src/bytes/buffer.go:274
			_go_fuzz_dep_.CoverTab[81]++
//line /snap/go/10455/src/bytes/buffer.go:274
			// _ = "end of CoverTab[81]"
//line /snap/go/10455/src/bytes/buffer.go:274
		}
//line /snap/go/10455/src/bytes/buffer.go:274
		// _ = "end of CoverTab[75]"
	} else {
//line /snap/go/10455/src/bytes/buffer.go:275
		_go_fuzz_dep_.CoverTab[524324]++
//line /snap/go/10455/src/bytes/buffer.go:275
		_go_fuzz_dep_.CoverTab[82]++
//line /snap/go/10455/src/bytes/buffer.go:275
		// _ = "end of CoverTab[82]"
//line /snap/go/10455/src/bytes/buffer.go:275
	}
//line /snap/go/10455/src/bytes/buffer.go:275
	// _ = "end of CoverTab[71]"
//line /snap/go/10455/src/bytes/buffer.go:275
	_go_fuzz_dep_.CoverTab[72]++

						b.Reset()
						return n, nil
//line /snap/go/10455/src/bytes/buffer.go:278
	// _ = "end of CoverTab[72]"
}

// WriteByte appends the byte c to the buffer, growing the buffer as needed.
//line /snap/go/10455/src/bytes/buffer.go:281
// The returned error is always nil, but is included to match bufio.Writer's
//line /snap/go/10455/src/bytes/buffer.go:281
// WriteByte. If the buffer becomes too large, WriteByte will panic with
//line /snap/go/10455/src/bytes/buffer.go:281
// ErrTooLarge.
//line /snap/go/10455/src/bytes/buffer.go:285
func (b *Buffer) WriteByte(c byte) error {
//line /snap/go/10455/src/bytes/buffer.go:285
	_go_fuzz_dep_.CoverTab[83]++
						b.lastRead = opInvalid
						m, ok := b.tryGrowByReslice(1)
						if !ok {
//line /snap/go/10455/src/bytes/buffer.go:288
		_go_fuzz_dep_.CoverTab[524331]++
//line /snap/go/10455/src/bytes/buffer.go:288
		_go_fuzz_dep_.CoverTab[85]++
							m = b.grow(1)
//line /snap/go/10455/src/bytes/buffer.go:289
		// _ = "end of CoverTab[85]"
	} else {
//line /snap/go/10455/src/bytes/buffer.go:290
		_go_fuzz_dep_.CoverTab[524332]++
//line /snap/go/10455/src/bytes/buffer.go:290
		_go_fuzz_dep_.CoverTab[86]++
//line /snap/go/10455/src/bytes/buffer.go:290
		// _ = "end of CoverTab[86]"
//line /snap/go/10455/src/bytes/buffer.go:290
	}
//line /snap/go/10455/src/bytes/buffer.go:290
	// _ = "end of CoverTab[83]"
//line /snap/go/10455/src/bytes/buffer.go:290
	_go_fuzz_dep_.CoverTab[84]++
						b.buf[m] = c
						return nil
//line /snap/go/10455/src/bytes/buffer.go:292
	// _ = "end of CoverTab[84]"
}

// WriteRune appends the UTF-8 encoding of Unicode code point r to the
//line /snap/go/10455/src/bytes/buffer.go:295
// buffer, returning its length and an error, which is always nil but is
//line /snap/go/10455/src/bytes/buffer.go:295
// included to match bufio.Writer's WriteRune. The buffer is grown as needed;
//line /snap/go/10455/src/bytes/buffer.go:295
// if it becomes too large, WriteRune will panic with ErrTooLarge.
//line /snap/go/10455/src/bytes/buffer.go:299
func (b *Buffer) WriteRune(r rune) (n int, err error) {
//line /snap/go/10455/src/bytes/buffer.go:299
	_go_fuzz_dep_.CoverTab[87]++

						if uint32(r) < utf8.RuneSelf {
//line /snap/go/10455/src/bytes/buffer.go:301
		_go_fuzz_dep_.CoverTab[524333]++
//line /snap/go/10455/src/bytes/buffer.go:301
		_go_fuzz_dep_.CoverTab[90]++
							b.WriteByte(byte(r))
							return 1, nil
//line /snap/go/10455/src/bytes/buffer.go:303
		// _ = "end of CoverTab[90]"
	} else {
//line /snap/go/10455/src/bytes/buffer.go:304
		_go_fuzz_dep_.CoverTab[524334]++
//line /snap/go/10455/src/bytes/buffer.go:304
		_go_fuzz_dep_.CoverTab[91]++
//line /snap/go/10455/src/bytes/buffer.go:304
		// _ = "end of CoverTab[91]"
//line /snap/go/10455/src/bytes/buffer.go:304
	}
//line /snap/go/10455/src/bytes/buffer.go:304
	// _ = "end of CoverTab[87]"
//line /snap/go/10455/src/bytes/buffer.go:304
	_go_fuzz_dep_.CoverTab[88]++
						b.lastRead = opInvalid
						m, ok := b.tryGrowByReslice(utf8.UTFMax)
						if !ok {
//line /snap/go/10455/src/bytes/buffer.go:307
		_go_fuzz_dep_.CoverTab[524335]++
//line /snap/go/10455/src/bytes/buffer.go:307
		_go_fuzz_dep_.CoverTab[92]++
							m = b.grow(utf8.UTFMax)
//line /snap/go/10455/src/bytes/buffer.go:308
		// _ = "end of CoverTab[92]"
	} else {
//line /snap/go/10455/src/bytes/buffer.go:309
		_go_fuzz_dep_.CoverTab[524336]++
//line /snap/go/10455/src/bytes/buffer.go:309
		_go_fuzz_dep_.CoverTab[93]++
//line /snap/go/10455/src/bytes/buffer.go:309
		// _ = "end of CoverTab[93]"
//line /snap/go/10455/src/bytes/buffer.go:309
	}
//line /snap/go/10455/src/bytes/buffer.go:309
	// _ = "end of CoverTab[88]"
//line /snap/go/10455/src/bytes/buffer.go:309
	_go_fuzz_dep_.CoverTab[89]++
						b.buf = utf8.AppendRune(b.buf[:m], r)
						return len(b.buf) - m, nil
//line /snap/go/10455/src/bytes/buffer.go:311
	// _ = "end of CoverTab[89]"
}

// Read reads the next len(p) bytes from the buffer or until the buffer
//line /snap/go/10455/src/bytes/buffer.go:314
// is drained. The return value n is the number of bytes read. If the
//line /snap/go/10455/src/bytes/buffer.go:314
// buffer has no data to return, err is io.EOF (unless len(p) is zero);
//line /snap/go/10455/src/bytes/buffer.go:314
// otherwise it is nil.
//line /snap/go/10455/src/bytes/buffer.go:318
func (b *Buffer) Read(p []byte) (n int, err error) {
//line /snap/go/10455/src/bytes/buffer.go:318
	_go_fuzz_dep_.CoverTab[94]++
						b.lastRead = opInvalid
						if b.empty() {
//line /snap/go/10455/src/bytes/buffer.go:320
		_go_fuzz_dep_.CoverTab[524337]++
//line /snap/go/10455/src/bytes/buffer.go:320
		_go_fuzz_dep_.CoverTab[97]++

							b.Reset()
							if len(p) == 0 {
//line /snap/go/10455/src/bytes/buffer.go:323
			_go_fuzz_dep_.CoverTab[524339]++
//line /snap/go/10455/src/bytes/buffer.go:323
			_go_fuzz_dep_.CoverTab[99]++
								return 0, nil
//line /snap/go/10455/src/bytes/buffer.go:324
			// _ = "end of CoverTab[99]"
		} else {
//line /snap/go/10455/src/bytes/buffer.go:325
			_go_fuzz_dep_.CoverTab[524340]++
//line /snap/go/10455/src/bytes/buffer.go:325
			_go_fuzz_dep_.CoverTab[100]++
//line /snap/go/10455/src/bytes/buffer.go:325
			// _ = "end of CoverTab[100]"
//line /snap/go/10455/src/bytes/buffer.go:325
		}
//line /snap/go/10455/src/bytes/buffer.go:325
		// _ = "end of CoverTab[97]"
//line /snap/go/10455/src/bytes/buffer.go:325
		_go_fuzz_dep_.CoverTab[98]++
							return 0, io.EOF
//line /snap/go/10455/src/bytes/buffer.go:326
		// _ = "end of CoverTab[98]"
	} else {
//line /snap/go/10455/src/bytes/buffer.go:327
		_go_fuzz_dep_.CoverTab[524338]++
//line /snap/go/10455/src/bytes/buffer.go:327
		_go_fuzz_dep_.CoverTab[101]++
//line /snap/go/10455/src/bytes/buffer.go:327
		// _ = "end of CoverTab[101]"
//line /snap/go/10455/src/bytes/buffer.go:327
	}
//line /snap/go/10455/src/bytes/buffer.go:327
	// _ = "end of CoverTab[94]"
//line /snap/go/10455/src/bytes/buffer.go:327
	_go_fuzz_dep_.CoverTab[95]++
						n = copy(p, b.buf[b.off:])
						b.off += n
						if n > 0 {
//line /snap/go/10455/src/bytes/buffer.go:330
		_go_fuzz_dep_.CoverTab[524341]++
//line /snap/go/10455/src/bytes/buffer.go:330
		_go_fuzz_dep_.CoverTab[102]++
							b.lastRead = opRead
//line /snap/go/10455/src/bytes/buffer.go:331
		// _ = "end of CoverTab[102]"
	} else {
//line /snap/go/10455/src/bytes/buffer.go:332
		_go_fuzz_dep_.CoverTab[524342]++
//line /snap/go/10455/src/bytes/buffer.go:332
		_go_fuzz_dep_.CoverTab[103]++
//line /snap/go/10455/src/bytes/buffer.go:332
		// _ = "end of CoverTab[103]"
//line /snap/go/10455/src/bytes/buffer.go:332
	}
//line /snap/go/10455/src/bytes/buffer.go:332
	// _ = "end of CoverTab[95]"
//line /snap/go/10455/src/bytes/buffer.go:332
	_go_fuzz_dep_.CoverTab[96]++
						return n, nil
//line /snap/go/10455/src/bytes/buffer.go:333
	// _ = "end of CoverTab[96]"
}

// Next returns a slice containing the next n bytes from the buffer,
//line /snap/go/10455/src/bytes/buffer.go:336
// advancing the buffer as if the bytes had been returned by Read.
//line /snap/go/10455/src/bytes/buffer.go:336
// If there are fewer than n bytes in the buffer, Next returns the entire buffer.
//line /snap/go/10455/src/bytes/buffer.go:336
// The slice is only valid until the next call to a read or write method.
//line /snap/go/10455/src/bytes/buffer.go:340
func (b *Buffer) Next(n int) []byte {
//line /snap/go/10455/src/bytes/buffer.go:340
	_go_fuzz_dep_.CoverTab[104]++
						b.lastRead = opInvalid
						m := b.Len()
						if n > m {
//line /snap/go/10455/src/bytes/buffer.go:343
		_go_fuzz_dep_.CoverTab[524343]++
//line /snap/go/10455/src/bytes/buffer.go:343
		_go_fuzz_dep_.CoverTab[107]++
							n = m
//line /snap/go/10455/src/bytes/buffer.go:344
		// _ = "end of CoverTab[107]"
	} else {
//line /snap/go/10455/src/bytes/buffer.go:345
		_go_fuzz_dep_.CoverTab[524344]++
//line /snap/go/10455/src/bytes/buffer.go:345
		_go_fuzz_dep_.CoverTab[108]++
//line /snap/go/10455/src/bytes/buffer.go:345
		// _ = "end of CoverTab[108]"
//line /snap/go/10455/src/bytes/buffer.go:345
	}
//line /snap/go/10455/src/bytes/buffer.go:345
	// _ = "end of CoverTab[104]"
//line /snap/go/10455/src/bytes/buffer.go:345
	_go_fuzz_dep_.CoverTab[105]++
						data := b.buf[b.off : b.off+n]
						b.off += n
						if n > 0 {
//line /snap/go/10455/src/bytes/buffer.go:348
		_go_fuzz_dep_.CoverTab[524345]++
//line /snap/go/10455/src/bytes/buffer.go:348
		_go_fuzz_dep_.CoverTab[109]++
							b.lastRead = opRead
//line /snap/go/10455/src/bytes/buffer.go:349
		// _ = "end of CoverTab[109]"
	} else {
//line /snap/go/10455/src/bytes/buffer.go:350
		_go_fuzz_dep_.CoverTab[524346]++
//line /snap/go/10455/src/bytes/buffer.go:350
		_go_fuzz_dep_.CoverTab[110]++
//line /snap/go/10455/src/bytes/buffer.go:350
		// _ = "end of CoverTab[110]"
//line /snap/go/10455/src/bytes/buffer.go:350
	}
//line /snap/go/10455/src/bytes/buffer.go:350
	// _ = "end of CoverTab[105]"
//line /snap/go/10455/src/bytes/buffer.go:350
	_go_fuzz_dep_.CoverTab[106]++
						return data
//line /snap/go/10455/src/bytes/buffer.go:351
	// _ = "end of CoverTab[106]"
}

// ReadByte reads and returns the next byte from the buffer.
//line /snap/go/10455/src/bytes/buffer.go:354
// If no byte is available, it returns error io.EOF.
//line /snap/go/10455/src/bytes/buffer.go:356
func (b *Buffer) ReadByte() (byte, error) {
//line /snap/go/10455/src/bytes/buffer.go:356
	_go_fuzz_dep_.CoverTab[111]++
						if b.empty() {
//line /snap/go/10455/src/bytes/buffer.go:357
		_go_fuzz_dep_.CoverTab[524347]++
//line /snap/go/10455/src/bytes/buffer.go:357
		_go_fuzz_dep_.CoverTab[113]++

							b.Reset()
							return 0, io.EOF
//line /snap/go/10455/src/bytes/buffer.go:360
		// _ = "end of CoverTab[113]"
	} else {
//line /snap/go/10455/src/bytes/buffer.go:361
		_go_fuzz_dep_.CoverTab[524348]++
//line /snap/go/10455/src/bytes/buffer.go:361
		_go_fuzz_dep_.CoverTab[114]++
//line /snap/go/10455/src/bytes/buffer.go:361
		// _ = "end of CoverTab[114]"
//line /snap/go/10455/src/bytes/buffer.go:361
	}
//line /snap/go/10455/src/bytes/buffer.go:361
	// _ = "end of CoverTab[111]"
//line /snap/go/10455/src/bytes/buffer.go:361
	_go_fuzz_dep_.CoverTab[112]++
						c := b.buf[b.off]
						b.off++
						b.lastRead = opRead
						return c, nil
//line /snap/go/10455/src/bytes/buffer.go:365
	// _ = "end of CoverTab[112]"
}

// ReadRune reads and returns the next UTF-8-encoded
//line /snap/go/10455/src/bytes/buffer.go:368
// Unicode code point from the buffer.
//line /snap/go/10455/src/bytes/buffer.go:368
// If no bytes are available, the error returned is io.EOF.
//line /snap/go/10455/src/bytes/buffer.go:368
// If the bytes are an erroneous UTF-8 encoding, it
//line /snap/go/10455/src/bytes/buffer.go:368
// consumes one byte and returns U+FFFD, 1.
//line /snap/go/10455/src/bytes/buffer.go:373
func (b *Buffer) ReadRune() (r rune, size int, err error) {
//line /snap/go/10455/src/bytes/buffer.go:373
	_go_fuzz_dep_.CoverTab[115]++
						if b.empty() {
//line /snap/go/10455/src/bytes/buffer.go:374
		_go_fuzz_dep_.CoverTab[524349]++
//line /snap/go/10455/src/bytes/buffer.go:374
		_go_fuzz_dep_.CoverTab[118]++

							b.Reset()
							return 0, 0, io.EOF
//line /snap/go/10455/src/bytes/buffer.go:377
		// _ = "end of CoverTab[118]"
	} else {
//line /snap/go/10455/src/bytes/buffer.go:378
		_go_fuzz_dep_.CoverTab[524350]++
//line /snap/go/10455/src/bytes/buffer.go:378
		_go_fuzz_dep_.CoverTab[119]++
//line /snap/go/10455/src/bytes/buffer.go:378
		// _ = "end of CoverTab[119]"
//line /snap/go/10455/src/bytes/buffer.go:378
	}
//line /snap/go/10455/src/bytes/buffer.go:378
	// _ = "end of CoverTab[115]"
//line /snap/go/10455/src/bytes/buffer.go:378
	_go_fuzz_dep_.CoverTab[116]++
						c := b.buf[b.off]
						if c < utf8.RuneSelf {
//line /snap/go/10455/src/bytes/buffer.go:380
		_go_fuzz_dep_.CoverTab[524351]++
//line /snap/go/10455/src/bytes/buffer.go:380
		_go_fuzz_dep_.CoverTab[120]++
							b.off++
							b.lastRead = opReadRune1
							return rune(c), 1, nil
//line /snap/go/10455/src/bytes/buffer.go:383
		// _ = "end of CoverTab[120]"
	} else {
//line /snap/go/10455/src/bytes/buffer.go:384
		_go_fuzz_dep_.CoverTab[524352]++
//line /snap/go/10455/src/bytes/buffer.go:384
		_go_fuzz_dep_.CoverTab[121]++
//line /snap/go/10455/src/bytes/buffer.go:384
		// _ = "end of CoverTab[121]"
//line /snap/go/10455/src/bytes/buffer.go:384
	}
//line /snap/go/10455/src/bytes/buffer.go:384
	// _ = "end of CoverTab[116]"
//line /snap/go/10455/src/bytes/buffer.go:384
	_go_fuzz_dep_.CoverTab[117]++
						r, n := utf8.DecodeRune(b.buf[b.off:])
						b.off += n
						b.lastRead = readOp(n)
						return r, n, nil
//line /snap/go/10455/src/bytes/buffer.go:388
	// _ = "end of CoverTab[117]"
}

// UnreadRune unreads the last rune returned by ReadRune.
//line /snap/go/10455/src/bytes/buffer.go:391
// If the most recent read or write operation on the buffer was
//line /snap/go/10455/src/bytes/buffer.go:391
// not a successful ReadRune, UnreadRune returns an error.  (In this regard
//line /snap/go/10455/src/bytes/buffer.go:391
// it is stricter than UnreadByte, which will unread the last byte
//line /snap/go/10455/src/bytes/buffer.go:391
// from any read operation.)
//line /snap/go/10455/src/bytes/buffer.go:396
func (b *Buffer) UnreadRune() error {
//line /snap/go/10455/src/bytes/buffer.go:396
	_go_fuzz_dep_.CoverTab[122]++
						if b.lastRead <= opInvalid {
//line /snap/go/10455/src/bytes/buffer.go:397
		_go_fuzz_dep_.CoverTab[524353]++
//line /snap/go/10455/src/bytes/buffer.go:397
		_go_fuzz_dep_.CoverTab[125]++
							return errors.New("bytes.Buffer: UnreadRune: previous operation was not a successful ReadRune")
//line /snap/go/10455/src/bytes/buffer.go:398
		// _ = "end of CoverTab[125]"
	} else {
//line /snap/go/10455/src/bytes/buffer.go:399
		_go_fuzz_dep_.CoverTab[524354]++
//line /snap/go/10455/src/bytes/buffer.go:399
		_go_fuzz_dep_.CoverTab[126]++
//line /snap/go/10455/src/bytes/buffer.go:399
		// _ = "end of CoverTab[126]"
//line /snap/go/10455/src/bytes/buffer.go:399
	}
//line /snap/go/10455/src/bytes/buffer.go:399
	// _ = "end of CoverTab[122]"
//line /snap/go/10455/src/bytes/buffer.go:399
	_go_fuzz_dep_.CoverTab[123]++
						if b.off >= int(b.lastRead) {
//line /snap/go/10455/src/bytes/buffer.go:400
		_go_fuzz_dep_.CoverTab[524355]++
//line /snap/go/10455/src/bytes/buffer.go:400
		_go_fuzz_dep_.CoverTab[127]++
							b.off -= int(b.lastRead)
//line /snap/go/10455/src/bytes/buffer.go:401
		// _ = "end of CoverTab[127]"
	} else {
//line /snap/go/10455/src/bytes/buffer.go:402
		_go_fuzz_dep_.CoverTab[524356]++
//line /snap/go/10455/src/bytes/buffer.go:402
		_go_fuzz_dep_.CoverTab[128]++
//line /snap/go/10455/src/bytes/buffer.go:402
		// _ = "end of CoverTab[128]"
//line /snap/go/10455/src/bytes/buffer.go:402
	}
//line /snap/go/10455/src/bytes/buffer.go:402
	// _ = "end of CoverTab[123]"
//line /snap/go/10455/src/bytes/buffer.go:402
	_go_fuzz_dep_.CoverTab[124]++
						b.lastRead = opInvalid
						return nil
//line /snap/go/10455/src/bytes/buffer.go:404
	// _ = "end of CoverTab[124]"
}

var errUnreadByte = errors.New("bytes.Buffer: UnreadByte: previous operation was not a successful read")

// UnreadByte unreads the last byte returned by the most recent successful
//line /snap/go/10455/src/bytes/buffer.go:409
// read operation that read at least one byte. If a write has happened since
//line /snap/go/10455/src/bytes/buffer.go:409
// the last read, if the last read returned an error, or if the read read zero
//line /snap/go/10455/src/bytes/buffer.go:409
// bytes, UnreadByte returns an error.
//line /snap/go/10455/src/bytes/buffer.go:413
func (b *Buffer) UnreadByte() error {
//line /snap/go/10455/src/bytes/buffer.go:413
	_go_fuzz_dep_.CoverTab[129]++
						if b.lastRead == opInvalid {
//line /snap/go/10455/src/bytes/buffer.go:414
		_go_fuzz_dep_.CoverTab[524357]++
//line /snap/go/10455/src/bytes/buffer.go:414
		_go_fuzz_dep_.CoverTab[132]++
							return errUnreadByte
//line /snap/go/10455/src/bytes/buffer.go:415
		// _ = "end of CoverTab[132]"
	} else {
//line /snap/go/10455/src/bytes/buffer.go:416
		_go_fuzz_dep_.CoverTab[524358]++
//line /snap/go/10455/src/bytes/buffer.go:416
		_go_fuzz_dep_.CoverTab[133]++
//line /snap/go/10455/src/bytes/buffer.go:416
		// _ = "end of CoverTab[133]"
//line /snap/go/10455/src/bytes/buffer.go:416
	}
//line /snap/go/10455/src/bytes/buffer.go:416
	// _ = "end of CoverTab[129]"
//line /snap/go/10455/src/bytes/buffer.go:416
	_go_fuzz_dep_.CoverTab[130]++
						b.lastRead = opInvalid
						if b.off > 0 {
//line /snap/go/10455/src/bytes/buffer.go:418
		_go_fuzz_dep_.CoverTab[524359]++
//line /snap/go/10455/src/bytes/buffer.go:418
		_go_fuzz_dep_.CoverTab[134]++
							b.off--
//line /snap/go/10455/src/bytes/buffer.go:419
		// _ = "end of CoverTab[134]"
	} else {
//line /snap/go/10455/src/bytes/buffer.go:420
		_go_fuzz_dep_.CoverTab[524360]++
//line /snap/go/10455/src/bytes/buffer.go:420
		_go_fuzz_dep_.CoverTab[135]++
//line /snap/go/10455/src/bytes/buffer.go:420
		// _ = "end of CoverTab[135]"
//line /snap/go/10455/src/bytes/buffer.go:420
	}
//line /snap/go/10455/src/bytes/buffer.go:420
	// _ = "end of CoverTab[130]"
//line /snap/go/10455/src/bytes/buffer.go:420
	_go_fuzz_dep_.CoverTab[131]++
						return nil
//line /snap/go/10455/src/bytes/buffer.go:421
	// _ = "end of CoverTab[131]"
}

// ReadBytes reads until the first occurrence of delim in the input,
//line /snap/go/10455/src/bytes/buffer.go:424
// returning a slice containing the data up to and including the delimiter.
//line /snap/go/10455/src/bytes/buffer.go:424
// If ReadBytes encounters an error before finding a delimiter,
//line /snap/go/10455/src/bytes/buffer.go:424
// it returns the data read before the error and the error itself (often io.EOF).
//line /snap/go/10455/src/bytes/buffer.go:424
// ReadBytes returns err != nil if and only if the returned data does not end in
//line /snap/go/10455/src/bytes/buffer.go:424
// delim.
//line /snap/go/10455/src/bytes/buffer.go:430
func (b *Buffer) ReadBytes(delim byte) (line []byte, err error) {
//line /snap/go/10455/src/bytes/buffer.go:430
	_go_fuzz_dep_.CoverTab[136]++
						slice, err := b.readSlice(delim)

//line /snap/go/10455/src/bytes/buffer.go:434
	line = append(line, slice...)
						return line, err
//line /snap/go/10455/src/bytes/buffer.go:435
	// _ = "end of CoverTab[136]"
}

// readSlice is like ReadBytes but returns a reference to internal buffer data.
func (b *Buffer) readSlice(delim byte) (line []byte, err error) {
//line /snap/go/10455/src/bytes/buffer.go:439
	_go_fuzz_dep_.CoverTab[137]++
						i := IndexByte(b.buf[b.off:], delim)
						end := b.off + i + 1
						if i < 0 {
//line /snap/go/10455/src/bytes/buffer.go:442
		_go_fuzz_dep_.CoverTab[524361]++
//line /snap/go/10455/src/bytes/buffer.go:442
		_go_fuzz_dep_.CoverTab[139]++
							end = len(b.buf)
							err = io.EOF
//line /snap/go/10455/src/bytes/buffer.go:444
		// _ = "end of CoverTab[139]"
	} else {
//line /snap/go/10455/src/bytes/buffer.go:445
		_go_fuzz_dep_.CoverTab[524362]++
//line /snap/go/10455/src/bytes/buffer.go:445
		_go_fuzz_dep_.CoverTab[140]++
//line /snap/go/10455/src/bytes/buffer.go:445
		// _ = "end of CoverTab[140]"
//line /snap/go/10455/src/bytes/buffer.go:445
	}
//line /snap/go/10455/src/bytes/buffer.go:445
	// _ = "end of CoverTab[137]"
//line /snap/go/10455/src/bytes/buffer.go:445
	_go_fuzz_dep_.CoverTab[138]++
						line = b.buf[b.off:end]
						b.off = end
						b.lastRead = opRead
						return line, err
//line /snap/go/10455/src/bytes/buffer.go:449
	// _ = "end of CoverTab[138]"
}

// ReadString reads until the first occurrence of delim in the input,
//line /snap/go/10455/src/bytes/buffer.go:452
// returning a string containing the data up to and including the delimiter.
//line /snap/go/10455/src/bytes/buffer.go:452
// If ReadString encounters an error before finding a delimiter,
//line /snap/go/10455/src/bytes/buffer.go:452
// it returns the data read before the error and the error itself (often io.EOF).
//line /snap/go/10455/src/bytes/buffer.go:452
// ReadString returns err != nil if and only if the returned data does not end
//line /snap/go/10455/src/bytes/buffer.go:452
// in delim.
//line /snap/go/10455/src/bytes/buffer.go:458
func (b *Buffer) ReadString(delim byte) (line string, err error) {
//line /snap/go/10455/src/bytes/buffer.go:458
	_go_fuzz_dep_.CoverTab[141]++
						slice, err := b.readSlice(delim)
						return string(slice), err
//line /snap/go/10455/src/bytes/buffer.go:460
	// _ = "end of CoverTab[141]"
}

// NewBuffer creates and initializes a new Buffer using buf as its
//line /snap/go/10455/src/bytes/buffer.go:463
// initial contents. The new Buffer takes ownership of buf, and the
//line /snap/go/10455/src/bytes/buffer.go:463
// caller should not use buf after this call. NewBuffer is intended to
//line /snap/go/10455/src/bytes/buffer.go:463
// prepare a Buffer to read existing data. It can also be used to set
//line /snap/go/10455/src/bytes/buffer.go:463
// the initial size of the internal buffer for writing. To do that,
//line /snap/go/10455/src/bytes/buffer.go:463
// buf should have the desired capacity but a length of zero.
//line /snap/go/10455/src/bytes/buffer.go:463
//
//line /snap/go/10455/src/bytes/buffer.go:463
// In most cases, new(Buffer) (or just declaring a Buffer variable) is
//line /snap/go/10455/src/bytes/buffer.go:463
// sufficient to initialize a Buffer.
//line /snap/go/10455/src/bytes/buffer.go:472
func NewBuffer(buf []byte) *Buffer {
//line /snap/go/10455/src/bytes/buffer.go:472
	_go_fuzz_dep_.CoverTab[142]++
//line /snap/go/10455/src/bytes/buffer.go:472
	return &Buffer{buf: buf}
//line /snap/go/10455/src/bytes/buffer.go:472
	// _ = "end of CoverTab[142]"
//line /snap/go/10455/src/bytes/buffer.go:472
}

// NewBufferString creates and initializes a new Buffer using string s as its
//line /snap/go/10455/src/bytes/buffer.go:474
// initial contents. It is intended to prepare a buffer to read an existing
//line /snap/go/10455/src/bytes/buffer.go:474
// string.
//line /snap/go/10455/src/bytes/buffer.go:474
//
//line /snap/go/10455/src/bytes/buffer.go:474
// In most cases, new(Buffer) (or just declaring a Buffer variable) is
//line /snap/go/10455/src/bytes/buffer.go:474
// sufficient to initialize a Buffer.
//line /snap/go/10455/src/bytes/buffer.go:480
func NewBufferString(s string) *Buffer {
//line /snap/go/10455/src/bytes/buffer.go:480
	_go_fuzz_dep_.CoverTab[143]++
						return &Buffer{buf: []byte(s)}
//line /snap/go/10455/src/bytes/buffer.go:481
	// _ = "end of CoverTab[143]"
}

//line /snap/go/10455/src/bytes/buffer.go:482
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /snap/go/10455/src/bytes/buffer.go:482
var _ = _go_fuzz_dep_.CoverTab
