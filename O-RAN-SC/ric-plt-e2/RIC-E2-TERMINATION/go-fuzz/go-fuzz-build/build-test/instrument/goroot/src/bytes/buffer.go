// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/bytes/buffer.go:5
package bytes

//line /usr/local/go/src/bytes/buffer.go:5
import (
//line /usr/local/go/src/bytes/buffer.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/bytes/buffer.go:5
)
//line /usr/local/go/src/bytes/buffer.go:5
import (
//line /usr/local/go/src/bytes/buffer.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/bytes/buffer.go:5
)

//line /usr/local/go/src/bytes/buffer.go:9
import (
	"errors"
	"io"
	"unicode/utf8"
)

// smallBufferSize is an initial allocation minimal capacity.
const smallBufferSize = 64

// A Buffer is a variable-sized buffer of bytes with Read and Write methods.
//line /usr/local/go/src/bytes/buffer.go:18
// The zero value for Buffer is an empty buffer ready to use.
//line /usr/local/go/src/bytes/buffer.go:20
type Buffer struct {
	buf		[]byte	// contents are the bytes buf[off : len(buf)]
	off		int	// read at &buf[off], write at &buf[len(buf)]
	lastRead	readOp	// last read operation, so that Unread* can work correctly.
}

// The readOp constants describe the last action performed on
//line /usr/local/go/src/bytes/buffer.go:26
// the buffer, so that UnreadRune and UnreadByte can check for
//line /usr/local/go/src/bytes/buffer.go:26
// invalid usage. opReadRuneX constants are chosen such that
//line /usr/local/go/src/bytes/buffer.go:26
// converted to int they correspond to the rune size that was read.
//line /usr/local/go/src/bytes/buffer.go:30
type readOp int8

// Don't use iota for these, as the values need to correspond with the
//line /usr/local/go/src/bytes/buffer.go:32
// names and comments, which is easier to see when being explicit.
//line /usr/local/go/src/bytes/buffer.go:34
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
//line /usr/local/go/src/bytes/buffer.go:49
// The slice is valid for use only until the next buffer modification (that is,
//line /usr/local/go/src/bytes/buffer.go:49
// only until the next call to a method like Read, Write, Reset, or Truncate).
//line /usr/local/go/src/bytes/buffer.go:49
// The slice aliases the buffer content at least until the next buffer modification,
//line /usr/local/go/src/bytes/buffer.go:49
// so immediate changes to the slice will affect the result of future reads.
//line /usr/local/go/src/bytes/buffer.go:54
func (b *Buffer) Bytes() []byte {
//line /usr/local/go/src/bytes/buffer.go:54
	_go_fuzz_dep_.CoverTab[1]++
//line /usr/local/go/src/bytes/buffer.go:54
	return b.buf[b.off:]
//line /usr/local/go/src/bytes/buffer.go:54
	// _ = "end of CoverTab[1]"
//line /usr/local/go/src/bytes/buffer.go:54
}

// String returns the contents of the unread portion of the buffer
//line /usr/local/go/src/bytes/buffer.go:56
// as a string. If the Buffer is a nil pointer, it returns "<nil>".
//line /usr/local/go/src/bytes/buffer.go:56
//
//line /usr/local/go/src/bytes/buffer.go:56
// To build strings more efficiently, see the strings.Builder type.
//line /usr/local/go/src/bytes/buffer.go:60
func (b *Buffer) String() string {
//line /usr/local/go/src/bytes/buffer.go:60
	_go_fuzz_dep_.CoverTab[2]++
						if b == nil {
//line /usr/local/go/src/bytes/buffer.go:61
		_go_fuzz_dep_.CoverTab[4]++

							return "<nil>"
//line /usr/local/go/src/bytes/buffer.go:63
		// _ = "end of CoverTab[4]"
	} else {
//line /usr/local/go/src/bytes/buffer.go:64
		_go_fuzz_dep_.CoverTab[5]++
//line /usr/local/go/src/bytes/buffer.go:64
		// _ = "end of CoverTab[5]"
//line /usr/local/go/src/bytes/buffer.go:64
	}
//line /usr/local/go/src/bytes/buffer.go:64
	// _ = "end of CoverTab[2]"
//line /usr/local/go/src/bytes/buffer.go:64
	_go_fuzz_dep_.CoverTab[3]++
						return string(b.buf[b.off:])
//line /usr/local/go/src/bytes/buffer.go:65
	// _ = "end of CoverTab[3]"
}

// empty reports whether the unread portion of the buffer is empty.
func (b *Buffer) empty() bool {
//line /usr/local/go/src/bytes/buffer.go:69
	_go_fuzz_dep_.CoverTab[6]++
//line /usr/local/go/src/bytes/buffer.go:69
	return len(b.buf) <= b.off
//line /usr/local/go/src/bytes/buffer.go:69
	// _ = "end of CoverTab[6]"
//line /usr/local/go/src/bytes/buffer.go:69
}

// Len returns the number of bytes of the unread portion of the buffer;
//line /usr/local/go/src/bytes/buffer.go:71
// b.Len() == len(b.Bytes()).
//line /usr/local/go/src/bytes/buffer.go:73
func (b *Buffer) Len() int {
//line /usr/local/go/src/bytes/buffer.go:73
	_go_fuzz_dep_.CoverTab[7]++
//line /usr/local/go/src/bytes/buffer.go:73
	return len(b.buf) - b.off
//line /usr/local/go/src/bytes/buffer.go:73
	// _ = "end of CoverTab[7]"
//line /usr/local/go/src/bytes/buffer.go:73
}

// Cap returns the capacity of the buffer's underlying byte slice, that is, the
//line /usr/local/go/src/bytes/buffer.go:75
// total space allocated for the buffer's data.
//line /usr/local/go/src/bytes/buffer.go:77
func (b *Buffer) Cap() int {
//line /usr/local/go/src/bytes/buffer.go:77
	_go_fuzz_dep_.CoverTab[8]++
//line /usr/local/go/src/bytes/buffer.go:77
	return cap(b.buf)
//line /usr/local/go/src/bytes/buffer.go:77
	// _ = "end of CoverTab[8]"
//line /usr/local/go/src/bytes/buffer.go:77
}

// Truncate discards all but the first n unread bytes from the buffer
//line /usr/local/go/src/bytes/buffer.go:79
// but continues to use the same allocated storage.
//line /usr/local/go/src/bytes/buffer.go:79
// It panics if n is negative or greater than the length of the buffer.
//line /usr/local/go/src/bytes/buffer.go:82
func (b *Buffer) Truncate(n int) {
//line /usr/local/go/src/bytes/buffer.go:82
	_go_fuzz_dep_.CoverTab[9]++
						if n == 0 {
//line /usr/local/go/src/bytes/buffer.go:83
		_go_fuzz_dep_.CoverTab[12]++
							b.Reset()
							return
//line /usr/local/go/src/bytes/buffer.go:85
		// _ = "end of CoverTab[12]"
	} else {
//line /usr/local/go/src/bytes/buffer.go:86
		_go_fuzz_dep_.CoverTab[13]++
//line /usr/local/go/src/bytes/buffer.go:86
		// _ = "end of CoverTab[13]"
//line /usr/local/go/src/bytes/buffer.go:86
	}
//line /usr/local/go/src/bytes/buffer.go:86
	// _ = "end of CoverTab[9]"
//line /usr/local/go/src/bytes/buffer.go:86
	_go_fuzz_dep_.CoverTab[10]++
						b.lastRead = opInvalid
						if n < 0 || func() bool {
//line /usr/local/go/src/bytes/buffer.go:88
		_go_fuzz_dep_.CoverTab[14]++
//line /usr/local/go/src/bytes/buffer.go:88
		return n > b.Len()
//line /usr/local/go/src/bytes/buffer.go:88
		// _ = "end of CoverTab[14]"
//line /usr/local/go/src/bytes/buffer.go:88
	}() {
//line /usr/local/go/src/bytes/buffer.go:88
		_go_fuzz_dep_.CoverTab[15]++
							panic("bytes.Buffer: truncation out of range")
//line /usr/local/go/src/bytes/buffer.go:89
		// _ = "end of CoverTab[15]"
	} else {
//line /usr/local/go/src/bytes/buffer.go:90
		_go_fuzz_dep_.CoverTab[16]++
//line /usr/local/go/src/bytes/buffer.go:90
		// _ = "end of CoverTab[16]"
//line /usr/local/go/src/bytes/buffer.go:90
	}
//line /usr/local/go/src/bytes/buffer.go:90
	// _ = "end of CoverTab[10]"
//line /usr/local/go/src/bytes/buffer.go:90
	_go_fuzz_dep_.CoverTab[11]++
						b.buf = b.buf[:b.off+n]
//line /usr/local/go/src/bytes/buffer.go:91
	// _ = "end of CoverTab[11]"
}

// Reset resets the buffer to be empty,
//line /usr/local/go/src/bytes/buffer.go:94
// but it retains the underlying storage for use by future writes.
//line /usr/local/go/src/bytes/buffer.go:94
// Reset is the same as Truncate(0).
//line /usr/local/go/src/bytes/buffer.go:97
func (b *Buffer) Reset() {
//line /usr/local/go/src/bytes/buffer.go:97
	_go_fuzz_dep_.CoverTab[17]++
						b.buf = b.buf[:0]
						b.off = 0
						b.lastRead = opInvalid
//line /usr/local/go/src/bytes/buffer.go:100
	// _ = "end of CoverTab[17]"
}

// tryGrowByReslice is a inlineable version of grow for the fast-case where the
//line /usr/local/go/src/bytes/buffer.go:103
// internal buffer only needs to be resliced.
//line /usr/local/go/src/bytes/buffer.go:103
// It returns the index where bytes should be written and whether it succeeded.
//line /usr/local/go/src/bytes/buffer.go:106
func (b *Buffer) tryGrowByReslice(n int) (int, bool) {
//line /usr/local/go/src/bytes/buffer.go:106
	_go_fuzz_dep_.CoverTab[18]++
						if l := len(b.buf); n <= cap(b.buf)-l {
//line /usr/local/go/src/bytes/buffer.go:107
		_go_fuzz_dep_.CoverTab[20]++
							b.buf = b.buf[:l+n]
							return l, true
//line /usr/local/go/src/bytes/buffer.go:109
		// _ = "end of CoverTab[20]"
	} else {
//line /usr/local/go/src/bytes/buffer.go:110
		_go_fuzz_dep_.CoverTab[21]++
//line /usr/local/go/src/bytes/buffer.go:110
		// _ = "end of CoverTab[21]"
//line /usr/local/go/src/bytes/buffer.go:110
	}
//line /usr/local/go/src/bytes/buffer.go:110
	// _ = "end of CoverTab[18]"
//line /usr/local/go/src/bytes/buffer.go:110
	_go_fuzz_dep_.CoverTab[19]++
						return 0, false
//line /usr/local/go/src/bytes/buffer.go:111
	// _ = "end of CoverTab[19]"
}

// grow grows the buffer to guarantee space for n more bytes.
//line /usr/local/go/src/bytes/buffer.go:114
// It returns the index where bytes should be written.
//line /usr/local/go/src/bytes/buffer.go:114
// If the buffer can't grow it will panic with ErrTooLarge.
//line /usr/local/go/src/bytes/buffer.go:117
func (b *Buffer) grow(n int) int {
//line /usr/local/go/src/bytes/buffer.go:117
	_go_fuzz_dep_.CoverTab[22]++
						m := b.Len()

						if m == 0 && func() bool {
//line /usr/local/go/src/bytes/buffer.go:120
		_go_fuzz_dep_.CoverTab[27]++
//line /usr/local/go/src/bytes/buffer.go:120
		return b.off != 0
//line /usr/local/go/src/bytes/buffer.go:120
		// _ = "end of CoverTab[27]"
//line /usr/local/go/src/bytes/buffer.go:120
	}() {
//line /usr/local/go/src/bytes/buffer.go:120
		_go_fuzz_dep_.CoverTab[28]++
							b.Reset()
//line /usr/local/go/src/bytes/buffer.go:121
		// _ = "end of CoverTab[28]"
	} else {
//line /usr/local/go/src/bytes/buffer.go:122
		_go_fuzz_dep_.CoverTab[29]++
//line /usr/local/go/src/bytes/buffer.go:122
		// _ = "end of CoverTab[29]"
//line /usr/local/go/src/bytes/buffer.go:122
	}
//line /usr/local/go/src/bytes/buffer.go:122
	// _ = "end of CoverTab[22]"
//line /usr/local/go/src/bytes/buffer.go:122
	_go_fuzz_dep_.CoverTab[23]++

						if i, ok := b.tryGrowByReslice(n); ok {
//line /usr/local/go/src/bytes/buffer.go:124
		_go_fuzz_dep_.CoverTab[30]++
							return i
//line /usr/local/go/src/bytes/buffer.go:125
		// _ = "end of CoverTab[30]"
	} else {
//line /usr/local/go/src/bytes/buffer.go:126
		_go_fuzz_dep_.CoverTab[31]++
//line /usr/local/go/src/bytes/buffer.go:126
		// _ = "end of CoverTab[31]"
//line /usr/local/go/src/bytes/buffer.go:126
	}
//line /usr/local/go/src/bytes/buffer.go:126
	// _ = "end of CoverTab[23]"
//line /usr/local/go/src/bytes/buffer.go:126
	_go_fuzz_dep_.CoverTab[24]++
						if b.buf == nil && func() bool {
//line /usr/local/go/src/bytes/buffer.go:127
		_go_fuzz_dep_.CoverTab[32]++
//line /usr/local/go/src/bytes/buffer.go:127
		return n <= smallBufferSize
//line /usr/local/go/src/bytes/buffer.go:127
		// _ = "end of CoverTab[32]"
//line /usr/local/go/src/bytes/buffer.go:127
	}() {
//line /usr/local/go/src/bytes/buffer.go:127
		_go_fuzz_dep_.CoverTab[33]++
							b.buf = make([]byte, n, smallBufferSize)
							return 0
//line /usr/local/go/src/bytes/buffer.go:129
		// _ = "end of CoverTab[33]"
	} else {
//line /usr/local/go/src/bytes/buffer.go:130
		_go_fuzz_dep_.CoverTab[34]++
//line /usr/local/go/src/bytes/buffer.go:130
		// _ = "end of CoverTab[34]"
//line /usr/local/go/src/bytes/buffer.go:130
	}
//line /usr/local/go/src/bytes/buffer.go:130
	// _ = "end of CoverTab[24]"
//line /usr/local/go/src/bytes/buffer.go:130
	_go_fuzz_dep_.CoverTab[25]++
						c := cap(b.buf)
						if n <= c/2-m {
//line /usr/local/go/src/bytes/buffer.go:132
		_go_fuzz_dep_.CoverTab[35]++

//line /usr/local/go/src/bytes/buffer.go:137
		copy(b.buf, b.buf[b.off:])
//line /usr/local/go/src/bytes/buffer.go:137
		// _ = "end of CoverTab[35]"
	} else {
//line /usr/local/go/src/bytes/buffer.go:138
		_go_fuzz_dep_.CoverTab[36]++
//line /usr/local/go/src/bytes/buffer.go:138
		if c > maxInt-c-n {
//line /usr/local/go/src/bytes/buffer.go:138
			_go_fuzz_dep_.CoverTab[37]++
								panic(ErrTooLarge)
//line /usr/local/go/src/bytes/buffer.go:139
			// _ = "end of CoverTab[37]"
		} else {
//line /usr/local/go/src/bytes/buffer.go:140
			_go_fuzz_dep_.CoverTab[38]++

								b.buf = growSlice(b.buf[b.off:], b.off+n)
//line /usr/local/go/src/bytes/buffer.go:142
			// _ = "end of CoverTab[38]"
		}
//line /usr/local/go/src/bytes/buffer.go:143
		// _ = "end of CoverTab[36]"
//line /usr/local/go/src/bytes/buffer.go:143
	}
//line /usr/local/go/src/bytes/buffer.go:143
	// _ = "end of CoverTab[25]"
//line /usr/local/go/src/bytes/buffer.go:143
	_go_fuzz_dep_.CoverTab[26]++

						b.off = 0
						b.buf = b.buf[:m+n]
						return m
//line /usr/local/go/src/bytes/buffer.go:147
	// _ = "end of CoverTab[26]"
}

// Grow grows the buffer's capacity, if necessary, to guarantee space for
//line /usr/local/go/src/bytes/buffer.go:150
// another n bytes. After Grow(n), at least n bytes can be written to the
//line /usr/local/go/src/bytes/buffer.go:150
// buffer without another allocation.
//line /usr/local/go/src/bytes/buffer.go:150
// If n is negative, Grow will panic.
//line /usr/local/go/src/bytes/buffer.go:150
// If the buffer can't grow it will panic with ErrTooLarge.
//line /usr/local/go/src/bytes/buffer.go:155
func (b *Buffer) Grow(n int) {
//line /usr/local/go/src/bytes/buffer.go:155
	_go_fuzz_dep_.CoverTab[39]++
						if n < 0 {
//line /usr/local/go/src/bytes/buffer.go:156
		_go_fuzz_dep_.CoverTab[41]++
							panic("bytes.Buffer.Grow: negative count")
//line /usr/local/go/src/bytes/buffer.go:157
		// _ = "end of CoverTab[41]"
	} else {
//line /usr/local/go/src/bytes/buffer.go:158
		_go_fuzz_dep_.CoverTab[42]++
//line /usr/local/go/src/bytes/buffer.go:158
		// _ = "end of CoverTab[42]"
//line /usr/local/go/src/bytes/buffer.go:158
	}
//line /usr/local/go/src/bytes/buffer.go:158
	// _ = "end of CoverTab[39]"
//line /usr/local/go/src/bytes/buffer.go:158
	_go_fuzz_dep_.CoverTab[40]++
						m := b.grow(n)
						b.buf = b.buf[:m]
//line /usr/local/go/src/bytes/buffer.go:160
	// _ = "end of CoverTab[40]"
}

// Write appends the contents of p to the buffer, growing the buffer as
//line /usr/local/go/src/bytes/buffer.go:163
// needed. The return value n is the length of p; err is always nil. If the
//line /usr/local/go/src/bytes/buffer.go:163
// buffer becomes too large, Write will panic with ErrTooLarge.
//line /usr/local/go/src/bytes/buffer.go:166
func (b *Buffer) Write(p []byte) (n int, err error) {
//line /usr/local/go/src/bytes/buffer.go:166
	_go_fuzz_dep_.CoverTab[43]++
						b.lastRead = opInvalid
						m, ok := b.tryGrowByReslice(len(p))
						if !ok {
//line /usr/local/go/src/bytes/buffer.go:169
		_go_fuzz_dep_.CoverTab[45]++
							m = b.grow(len(p))
//line /usr/local/go/src/bytes/buffer.go:170
		// _ = "end of CoverTab[45]"
	} else {
//line /usr/local/go/src/bytes/buffer.go:171
		_go_fuzz_dep_.CoverTab[46]++
//line /usr/local/go/src/bytes/buffer.go:171
		// _ = "end of CoverTab[46]"
//line /usr/local/go/src/bytes/buffer.go:171
	}
//line /usr/local/go/src/bytes/buffer.go:171
	// _ = "end of CoverTab[43]"
//line /usr/local/go/src/bytes/buffer.go:171
	_go_fuzz_dep_.CoverTab[44]++
						return copy(b.buf[m:], p), nil
//line /usr/local/go/src/bytes/buffer.go:172
	// _ = "end of CoverTab[44]"
}

// WriteString appends the contents of s to the buffer, growing the buffer as
//line /usr/local/go/src/bytes/buffer.go:175
// needed. The return value n is the length of s; err is always nil. If the
//line /usr/local/go/src/bytes/buffer.go:175
// buffer becomes too large, WriteString will panic with ErrTooLarge.
//line /usr/local/go/src/bytes/buffer.go:178
func (b *Buffer) WriteString(s string) (n int, err error) {
//line /usr/local/go/src/bytes/buffer.go:178
	_go_fuzz_dep_.CoverTab[47]++
						b.lastRead = opInvalid
						m, ok := b.tryGrowByReslice(len(s))
						if !ok {
//line /usr/local/go/src/bytes/buffer.go:181
		_go_fuzz_dep_.CoverTab[49]++
							m = b.grow(len(s))
//line /usr/local/go/src/bytes/buffer.go:182
		// _ = "end of CoverTab[49]"
	} else {
//line /usr/local/go/src/bytes/buffer.go:183
		_go_fuzz_dep_.CoverTab[50]++
//line /usr/local/go/src/bytes/buffer.go:183
		// _ = "end of CoverTab[50]"
//line /usr/local/go/src/bytes/buffer.go:183
	}
//line /usr/local/go/src/bytes/buffer.go:183
	// _ = "end of CoverTab[47]"
//line /usr/local/go/src/bytes/buffer.go:183
	_go_fuzz_dep_.CoverTab[48]++
						return copy(b.buf[m:], s), nil
//line /usr/local/go/src/bytes/buffer.go:184
	// _ = "end of CoverTab[48]"
}

// MinRead is the minimum slice size passed to a Read call by
//line /usr/local/go/src/bytes/buffer.go:187
// Buffer.ReadFrom. As long as the Buffer has at least MinRead bytes beyond
//line /usr/local/go/src/bytes/buffer.go:187
// what is required to hold the contents of r, ReadFrom will not grow the
//line /usr/local/go/src/bytes/buffer.go:187
// underlying buffer.
//line /usr/local/go/src/bytes/buffer.go:191
const MinRead = 512

// ReadFrom reads data from r until EOF and appends it to the buffer, growing
//line /usr/local/go/src/bytes/buffer.go:193
// the buffer as needed. The return value n is the number of bytes read. Any
//line /usr/local/go/src/bytes/buffer.go:193
// error except io.EOF encountered during the read is also returned. If the
//line /usr/local/go/src/bytes/buffer.go:193
// buffer becomes too large, ReadFrom will panic with ErrTooLarge.
//line /usr/local/go/src/bytes/buffer.go:197
func (b *Buffer) ReadFrom(r io.Reader) (n int64, err error) {
//line /usr/local/go/src/bytes/buffer.go:197
	_go_fuzz_dep_.CoverTab[51]++
						b.lastRead = opInvalid
						for {
//line /usr/local/go/src/bytes/buffer.go:199
		_go_fuzz_dep_.CoverTab[52]++
							i := b.grow(MinRead)
							b.buf = b.buf[:i]
							m, e := r.Read(b.buf[i:cap(b.buf)])
							if m < 0 {
//line /usr/local/go/src/bytes/buffer.go:203
			_go_fuzz_dep_.CoverTab[55]++
								panic(errNegativeRead)
//line /usr/local/go/src/bytes/buffer.go:204
			// _ = "end of CoverTab[55]"
		} else {
//line /usr/local/go/src/bytes/buffer.go:205
			_go_fuzz_dep_.CoverTab[56]++
//line /usr/local/go/src/bytes/buffer.go:205
			// _ = "end of CoverTab[56]"
//line /usr/local/go/src/bytes/buffer.go:205
		}
//line /usr/local/go/src/bytes/buffer.go:205
		// _ = "end of CoverTab[52]"
//line /usr/local/go/src/bytes/buffer.go:205
		_go_fuzz_dep_.CoverTab[53]++

							b.buf = b.buf[:i+m]
							n += int64(m)
							if e == io.EOF {
//line /usr/local/go/src/bytes/buffer.go:209
			_go_fuzz_dep_.CoverTab[57]++
								return n, nil
//line /usr/local/go/src/bytes/buffer.go:210
			// _ = "end of CoverTab[57]"
		} else {
//line /usr/local/go/src/bytes/buffer.go:211
			_go_fuzz_dep_.CoverTab[58]++
//line /usr/local/go/src/bytes/buffer.go:211
			// _ = "end of CoverTab[58]"
//line /usr/local/go/src/bytes/buffer.go:211
		}
//line /usr/local/go/src/bytes/buffer.go:211
		// _ = "end of CoverTab[53]"
//line /usr/local/go/src/bytes/buffer.go:211
		_go_fuzz_dep_.CoverTab[54]++
							if e != nil {
//line /usr/local/go/src/bytes/buffer.go:212
			_go_fuzz_dep_.CoverTab[59]++
								return n, e
//line /usr/local/go/src/bytes/buffer.go:213
			// _ = "end of CoverTab[59]"
		} else {
//line /usr/local/go/src/bytes/buffer.go:214
			_go_fuzz_dep_.CoverTab[60]++
//line /usr/local/go/src/bytes/buffer.go:214
			// _ = "end of CoverTab[60]"
//line /usr/local/go/src/bytes/buffer.go:214
		}
//line /usr/local/go/src/bytes/buffer.go:214
		// _ = "end of CoverTab[54]"
	}
//line /usr/local/go/src/bytes/buffer.go:215
	// _ = "end of CoverTab[51]"
}

// growSlice grows b by n, preserving the original content of b.
//line /usr/local/go/src/bytes/buffer.go:218
// If the allocation fails, it panics with ErrTooLarge.
//line /usr/local/go/src/bytes/buffer.go:220
func growSlice(b []byte, n int) []byte {
//line /usr/local/go/src/bytes/buffer.go:220
	_go_fuzz_dep_.CoverTab[61]++
						defer func() {
//line /usr/local/go/src/bytes/buffer.go:221
		_go_fuzz_dep_.CoverTab[64]++
							if recover() != nil {
//line /usr/local/go/src/bytes/buffer.go:222
			_go_fuzz_dep_.CoverTab[65]++
								panic(ErrTooLarge)
//line /usr/local/go/src/bytes/buffer.go:223
			// _ = "end of CoverTab[65]"
		} else {
//line /usr/local/go/src/bytes/buffer.go:224
			_go_fuzz_dep_.CoverTab[66]++
//line /usr/local/go/src/bytes/buffer.go:224
			// _ = "end of CoverTab[66]"
//line /usr/local/go/src/bytes/buffer.go:224
		}
//line /usr/local/go/src/bytes/buffer.go:224
		// _ = "end of CoverTab[64]"
	}()
//line /usr/local/go/src/bytes/buffer.go:225
	// _ = "end of CoverTab[61]"
//line /usr/local/go/src/bytes/buffer.go:225
	_go_fuzz_dep_.CoverTab[62]++

//line /usr/local/go/src/bytes/buffer.go:234
	c := len(b) + n
	if c < 2*cap(b) {
//line /usr/local/go/src/bytes/buffer.go:235
		_go_fuzz_dep_.CoverTab[67]++

//line /usr/local/go/src/bytes/buffer.go:238
		c = 2 * cap(b)
//line /usr/local/go/src/bytes/buffer.go:238
		// _ = "end of CoverTab[67]"
	} else {
//line /usr/local/go/src/bytes/buffer.go:239
		_go_fuzz_dep_.CoverTab[68]++
//line /usr/local/go/src/bytes/buffer.go:239
		// _ = "end of CoverTab[68]"
//line /usr/local/go/src/bytes/buffer.go:239
	}
//line /usr/local/go/src/bytes/buffer.go:239
	// _ = "end of CoverTab[62]"
//line /usr/local/go/src/bytes/buffer.go:239
	_go_fuzz_dep_.CoverTab[63]++
						b2 := append([]byte(nil), make([]byte, c)...)
						copy(b2, b)
						return b2[:len(b)]
//line /usr/local/go/src/bytes/buffer.go:242
	// _ = "end of CoverTab[63]"
}

// WriteTo writes data to w until the buffer is drained or an error occurs.
//line /usr/local/go/src/bytes/buffer.go:245
// The return value n is the number of bytes written; it always fits into an
//line /usr/local/go/src/bytes/buffer.go:245
// int, but it is int64 to match the io.WriterTo interface. Any error
//line /usr/local/go/src/bytes/buffer.go:245
// encountered during the write is also returned.
//line /usr/local/go/src/bytes/buffer.go:249
func (b *Buffer) WriteTo(w io.Writer) (n int64, err error) {
//line /usr/local/go/src/bytes/buffer.go:249
	_go_fuzz_dep_.CoverTab[69]++
						b.lastRead = opInvalid
						if nBytes := b.Len(); nBytes > 0 {
//line /usr/local/go/src/bytes/buffer.go:251
		_go_fuzz_dep_.CoverTab[71]++
							m, e := w.Write(b.buf[b.off:])
							if m > nBytes {
//line /usr/local/go/src/bytes/buffer.go:253
			_go_fuzz_dep_.CoverTab[74]++
								panic("bytes.Buffer.WriteTo: invalid Write count")
//line /usr/local/go/src/bytes/buffer.go:254
			// _ = "end of CoverTab[74]"
		} else {
//line /usr/local/go/src/bytes/buffer.go:255
			_go_fuzz_dep_.CoverTab[75]++
//line /usr/local/go/src/bytes/buffer.go:255
			// _ = "end of CoverTab[75]"
//line /usr/local/go/src/bytes/buffer.go:255
		}
//line /usr/local/go/src/bytes/buffer.go:255
		// _ = "end of CoverTab[71]"
//line /usr/local/go/src/bytes/buffer.go:255
		_go_fuzz_dep_.CoverTab[72]++
							b.off += m
							n = int64(m)
							if e != nil {
//line /usr/local/go/src/bytes/buffer.go:258
			_go_fuzz_dep_.CoverTab[76]++
								return n, e
//line /usr/local/go/src/bytes/buffer.go:259
			// _ = "end of CoverTab[76]"
		} else {
//line /usr/local/go/src/bytes/buffer.go:260
			_go_fuzz_dep_.CoverTab[77]++
//line /usr/local/go/src/bytes/buffer.go:260
			// _ = "end of CoverTab[77]"
//line /usr/local/go/src/bytes/buffer.go:260
		}
//line /usr/local/go/src/bytes/buffer.go:260
		// _ = "end of CoverTab[72]"
//line /usr/local/go/src/bytes/buffer.go:260
		_go_fuzz_dep_.CoverTab[73]++

//line /usr/local/go/src/bytes/buffer.go:263
		if m != nBytes {
//line /usr/local/go/src/bytes/buffer.go:263
			_go_fuzz_dep_.CoverTab[78]++
								return n, io.ErrShortWrite
//line /usr/local/go/src/bytes/buffer.go:264
			// _ = "end of CoverTab[78]"
		} else {
//line /usr/local/go/src/bytes/buffer.go:265
			_go_fuzz_dep_.CoverTab[79]++
//line /usr/local/go/src/bytes/buffer.go:265
			// _ = "end of CoverTab[79]"
//line /usr/local/go/src/bytes/buffer.go:265
		}
//line /usr/local/go/src/bytes/buffer.go:265
		// _ = "end of CoverTab[73]"
	} else {
//line /usr/local/go/src/bytes/buffer.go:266
		_go_fuzz_dep_.CoverTab[80]++
//line /usr/local/go/src/bytes/buffer.go:266
		// _ = "end of CoverTab[80]"
//line /usr/local/go/src/bytes/buffer.go:266
	}
//line /usr/local/go/src/bytes/buffer.go:266
	// _ = "end of CoverTab[69]"
//line /usr/local/go/src/bytes/buffer.go:266
	_go_fuzz_dep_.CoverTab[70]++

						b.Reset()
						return n, nil
//line /usr/local/go/src/bytes/buffer.go:269
	// _ = "end of CoverTab[70]"
}

// WriteByte appends the byte c to the buffer, growing the buffer as needed.
//line /usr/local/go/src/bytes/buffer.go:272
// The returned error is always nil, but is included to match bufio.Writer's
//line /usr/local/go/src/bytes/buffer.go:272
// WriteByte. If the buffer becomes too large, WriteByte will panic with
//line /usr/local/go/src/bytes/buffer.go:272
// ErrTooLarge.
//line /usr/local/go/src/bytes/buffer.go:276
func (b *Buffer) WriteByte(c byte) error {
//line /usr/local/go/src/bytes/buffer.go:276
	_go_fuzz_dep_.CoverTab[81]++
						b.lastRead = opInvalid
						m, ok := b.tryGrowByReslice(1)
						if !ok {
//line /usr/local/go/src/bytes/buffer.go:279
		_go_fuzz_dep_.CoverTab[83]++
							m = b.grow(1)
//line /usr/local/go/src/bytes/buffer.go:280
		// _ = "end of CoverTab[83]"
	} else {
//line /usr/local/go/src/bytes/buffer.go:281
		_go_fuzz_dep_.CoverTab[84]++
//line /usr/local/go/src/bytes/buffer.go:281
		// _ = "end of CoverTab[84]"
//line /usr/local/go/src/bytes/buffer.go:281
	}
//line /usr/local/go/src/bytes/buffer.go:281
	// _ = "end of CoverTab[81]"
//line /usr/local/go/src/bytes/buffer.go:281
	_go_fuzz_dep_.CoverTab[82]++
						b.buf[m] = c
						return nil
//line /usr/local/go/src/bytes/buffer.go:283
	// _ = "end of CoverTab[82]"
}

// WriteRune appends the UTF-8 encoding of Unicode code point r to the
//line /usr/local/go/src/bytes/buffer.go:286
// buffer, returning its length and an error, which is always nil but is
//line /usr/local/go/src/bytes/buffer.go:286
// included to match bufio.Writer's WriteRune. The buffer is grown as needed;
//line /usr/local/go/src/bytes/buffer.go:286
// if it becomes too large, WriteRune will panic with ErrTooLarge.
//line /usr/local/go/src/bytes/buffer.go:290
func (b *Buffer) WriteRune(r rune) (n int, err error) {
//line /usr/local/go/src/bytes/buffer.go:290
	_go_fuzz_dep_.CoverTab[85]++

						if uint32(r) < utf8.RuneSelf {
//line /usr/local/go/src/bytes/buffer.go:292
		_go_fuzz_dep_.CoverTab[88]++
							b.WriteByte(byte(r))
							return 1, nil
//line /usr/local/go/src/bytes/buffer.go:294
		// _ = "end of CoverTab[88]"
	} else {
//line /usr/local/go/src/bytes/buffer.go:295
		_go_fuzz_dep_.CoverTab[89]++
//line /usr/local/go/src/bytes/buffer.go:295
		// _ = "end of CoverTab[89]"
//line /usr/local/go/src/bytes/buffer.go:295
	}
//line /usr/local/go/src/bytes/buffer.go:295
	// _ = "end of CoverTab[85]"
//line /usr/local/go/src/bytes/buffer.go:295
	_go_fuzz_dep_.CoverTab[86]++
						b.lastRead = opInvalid
						m, ok := b.tryGrowByReslice(utf8.UTFMax)
						if !ok {
//line /usr/local/go/src/bytes/buffer.go:298
		_go_fuzz_dep_.CoverTab[90]++
							m = b.grow(utf8.UTFMax)
//line /usr/local/go/src/bytes/buffer.go:299
		// _ = "end of CoverTab[90]"
	} else {
//line /usr/local/go/src/bytes/buffer.go:300
		_go_fuzz_dep_.CoverTab[91]++
//line /usr/local/go/src/bytes/buffer.go:300
		// _ = "end of CoverTab[91]"
//line /usr/local/go/src/bytes/buffer.go:300
	}
//line /usr/local/go/src/bytes/buffer.go:300
	// _ = "end of CoverTab[86]"
//line /usr/local/go/src/bytes/buffer.go:300
	_go_fuzz_dep_.CoverTab[87]++
						b.buf = utf8.AppendRune(b.buf[:m], r)
						return len(b.buf) - m, nil
//line /usr/local/go/src/bytes/buffer.go:302
	// _ = "end of CoverTab[87]"
}

// Read reads the next len(p) bytes from the buffer or until the buffer
//line /usr/local/go/src/bytes/buffer.go:305
// is drained. The return value n is the number of bytes read. If the
//line /usr/local/go/src/bytes/buffer.go:305
// buffer has no data to return, err is io.EOF (unless len(p) is zero);
//line /usr/local/go/src/bytes/buffer.go:305
// otherwise it is nil.
//line /usr/local/go/src/bytes/buffer.go:309
func (b *Buffer) Read(p []byte) (n int, err error) {
//line /usr/local/go/src/bytes/buffer.go:309
	_go_fuzz_dep_.CoverTab[92]++
						b.lastRead = opInvalid
						if b.empty() {
//line /usr/local/go/src/bytes/buffer.go:311
		_go_fuzz_dep_.CoverTab[95]++

							b.Reset()
							if len(p) == 0 {
//line /usr/local/go/src/bytes/buffer.go:314
			_go_fuzz_dep_.CoverTab[97]++
								return 0, nil
//line /usr/local/go/src/bytes/buffer.go:315
			// _ = "end of CoverTab[97]"
		} else {
//line /usr/local/go/src/bytes/buffer.go:316
			_go_fuzz_dep_.CoverTab[98]++
//line /usr/local/go/src/bytes/buffer.go:316
			// _ = "end of CoverTab[98]"
//line /usr/local/go/src/bytes/buffer.go:316
		}
//line /usr/local/go/src/bytes/buffer.go:316
		// _ = "end of CoverTab[95]"
//line /usr/local/go/src/bytes/buffer.go:316
		_go_fuzz_dep_.CoverTab[96]++
							return 0, io.EOF
//line /usr/local/go/src/bytes/buffer.go:317
		// _ = "end of CoverTab[96]"
	} else {
//line /usr/local/go/src/bytes/buffer.go:318
		_go_fuzz_dep_.CoverTab[99]++
//line /usr/local/go/src/bytes/buffer.go:318
		// _ = "end of CoverTab[99]"
//line /usr/local/go/src/bytes/buffer.go:318
	}
//line /usr/local/go/src/bytes/buffer.go:318
	// _ = "end of CoverTab[92]"
//line /usr/local/go/src/bytes/buffer.go:318
	_go_fuzz_dep_.CoverTab[93]++
						n = copy(p, b.buf[b.off:])
						b.off += n
						if n > 0 {
//line /usr/local/go/src/bytes/buffer.go:321
		_go_fuzz_dep_.CoverTab[100]++
							b.lastRead = opRead
//line /usr/local/go/src/bytes/buffer.go:322
		// _ = "end of CoverTab[100]"
	} else {
//line /usr/local/go/src/bytes/buffer.go:323
		_go_fuzz_dep_.CoverTab[101]++
//line /usr/local/go/src/bytes/buffer.go:323
		// _ = "end of CoverTab[101]"
//line /usr/local/go/src/bytes/buffer.go:323
	}
//line /usr/local/go/src/bytes/buffer.go:323
	// _ = "end of CoverTab[93]"
//line /usr/local/go/src/bytes/buffer.go:323
	_go_fuzz_dep_.CoverTab[94]++
						return n, nil
//line /usr/local/go/src/bytes/buffer.go:324
	// _ = "end of CoverTab[94]"
}

// Next returns a slice containing the next n bytes from the buffer,
//line /usr/local/go/src/bytes/buffer.go:327
// advancing the buffer as if the bytes had been returned by Read.
//line /usr/local/go/src/bytes/buffer.go:327
// If there are fewer than n bytes in the buffer, Next returns the entire buffer.
//line /usr/local/go/src/bytes/buffer.go:327
// The slice is only valid until the next call to a read or write method.
//line /usr/local/go/src/bytes/buffer.go:331
func (b *Buffer) Next(n int) []byte {
//line /usr/local/go/src/bytes/buffer.go:331
	_go_fuzz_dep_.CoverTab[102]++
						b.lastRead = opInvalid
						m := b.Len()
						if n > m {
//line /usr/local/go/src/bytes/buffer.go:334
		_go_fuzz_dep_.CoverTab[105]++
							n = m
//line /usr/local/go/src/bytes/buffer.go:335
		// _ = "end of CoverTab[105]"
	} else {
//line /usr/local/go/src/bytes/buffer.go:336
		_go_fuzz_dep_.CoverTab[106]++
//line /usr/local/go/src/bytes/buffer.go:336
		// _ = "end of CoverTab[106]"
//line /usr/local/go/src/bytes/buffer.go:336
	}
//line /usr/local/go/src/bytes/buffer.go:336
	// _ = "end of CoverTab[102]"
//line /usr/local/go/src/bytes/buffer.go:336
	_go_fuzz_dep_.CoverTab[103]++
						data := b.buf[b.off : b.off+n]
						b.off += n
						if n > 0 {
//line /usr/local/go/src/bytes/buffer.go:339
		_go_fuzz_dep_.CoverTab[107]++
							b.lastRead = opRead
//line /usr/local/go/src/bytes/buffer.go:340
		// _ = "end of CoverTab[107]"
	} else {
//line /usr/local/go/src/bytes/buffer.go:341
		_go_fuzz_dep_.CoverTab[108]++
//line /usr/local/go/src/bytes/buffer.go:341
		// _ = "end of CoverTab[108]"
//line /usr/local/go/src/bytes/buffer.go:341
	}
//line /usr/local/go/src/bytes/buffer.go:341
	// _ = "end of CoverTab[103]"
//line /usr/local/go/src/bytes/buffer.go:341
	_go_fuzz_dep_.CoverTab[104]++
						return data
//line /usr/local/go/src/bytes/buffer.go:342
	// _ = "end of CoverTab[104]"
}

// ReadByte reads and returns the next byte from the buffer.
//line /usr/local/go/src/bytes/buffer.go:345
// If no byte is available, it returns error io.EOF.
//line /usr/local/go/src/bytes/buffer.go:347
func (b *Buffer) ReadByte() (byte, error) {
//line /usr/local/go/src/bytes/buffer.go:347
	_go_fuzz_dep_.CoverTab[109]++
						if b.empty() {
//line /usr/local/go/src/bytes/buffer.go:348
		_go_fuzz_dep_.CoverTab[111]++

							b.Reset()
							return 0, io.EOF
//line /usr/local/go/src/bytes/buffer.go:351
		// _ = "end of CoverTab[111]"
	} else {
//line /usr/local/go/src/bytes/buffer.go:352
		_go_fuzz_dep_.CoverTab[112]++
//line /usr/local/go/src/bytes/buffer.go:352
		// _ = "end of CoverTab[112]"
//line /usr/local/go/src/bytes/buffer.go:352
	}
//line /usr/local/go/src/bytes/buffer.go:352
	// _ = "end of CoverTab[109]"
//line /usr/local/go/src/bytes/buffer.go:352
	_go_fuzz_dep_.CoverTab[110]++
						c := b.buf[b.off]
						b.off++
						b.lastRead = opRead
						return c, nil
//line /usr/local/go/src/bytes/buffer.go:356
	// _ = "end of CoverTab[110]"
}

// ReadRune reads and returns the next UTF-8-encoded
//line /usr/local/go/src/bytes/buffer.go:359
// Unicode code point from the buffer.
//line /usr/local/go/src/bytes/buffer.go:359
// If no bytes are available, the error returned is io.EOF.
//line /usr/local/go/src/bytes/buffer.go:359
// If the bytes are an erroneous UTF-8 encoding, it
//line /usr/local/go/src/bytes/buffer.go:359
// consumes one byte and returns U+FFFD, 1.
//line /usr/local/go/src/bytes/buffer.go:364
func (b *Buffer) ReadRune() (r rune, size int, err error) {
//line /usr/local/go/src/bytes/buffer.go:364
	_go_fuzz_dep_.CoverTab[113]++
						if b.empty() {
//line /usr/local/go/src/bytes/buffer.go:365
		_go_fuzz_dep_.CoverTab[116]++

							b.Reset()
							return 0, 0, io.EOF
//line /usr/local/go/src/bytes/buffer.go:368
		// _ = "end of CoverTab[116]"
	} else {
//line /usr/local/go/src/bytes/buffer.go:369
		_go_fuzz_dep_.CoverTab[117]++
//line /usr/local/go/src/bytes/buffer.go:369
		// _ = "end of CoverTab[117]"
//line /usr/local/go/src/bytes/buffer.go:369
	}
//line /usr/local/go/src/bytes/buffer.go:369
	// _ = "end of CoverTab[113]"
//line /usr/local/go/src/bytes/buffer.go:369
	_go_fuzz_dep_.CoverTab[114]++
						c := b.buf[b.off]
						if c < utf8.RuneSelf {
//line /usr/local/go/src/bytes/buffer.go:371
		_go_fuzz_dep_.CoverTab[118]++
							b.off++
							b.lastRead = opReadRune1
							return rune(c), 1, nil
//line /usr/local/go/src/bytes/buffer.go:374
		// _ = "end of CoverTab[118]"
	} else {
//line /usr/local/go/src/bytes/buffer.go:375
		_go_fuzz_dep_.CoverTab[119]++
//line /usr/local/go/src/bytes/buffer.go:375
		// _ = "end of CoverTab[119]"
//line /usr/local/go/src/bytes/buffer.go:375
	}
//line /usr/local/go/src/bytes/buffer.go:375
	// _ = "end of CoverTab[114]"
//line /usr/local/go/src/bytes/buffer.go:375
	_go_fuzz_dep_.CoverTab[115]++
						r, n := utf8.DecodeRune(b.buf[b.off:])
						b.off += n
						b.lastRead = readOp(n)
						return r, n, nil
//line /usr/local/go/src/bytes/buffer.go:379
	// _ = "end of CoverTab[115]"
}

// UnreadRune unreads the last rune returned by ReadRune.
//line /usr/local/go/src/bytes/buffer.go:382
// If the most recent read or write operation on the buffer was
//line /usr/local/go/src/bytes/buffer.go:382
// not a successful ReadRune, UnreadRune returns an error.  (In this regard
//line /usr/local/go/src/bytes/buffer.go:382
// it is stricter than UnreadByte, which will unread the last byte
//line /usr/local/go/src/bytes/buffer.go:382
// from any read operation.)
//line /usr/local/go/src/bytes/buffer.go:387
func (b *Buffer) UnreadRune() error {
//line /usr/local/go/src/bytes/buffer.go:387
	_go_fuzz_dep_.CoverTab[120]++
						if b.lastRead <= opInvalid {
//line /usr/local/go/src/bytes/buffer.go:388
		_go_fuzz_dep_.CoverTab[123]++
							return errors.New("bytes.Buffer: UnreadRune: previous operation was not a successful ReadRune")
//line /usr/local/go/src/bytes/buffer.go:389
		// _ = "end of CoverTab[123]"
	} else {
//line /usr/local/go/src/bytes/buffer.go:390
		_go_fuzz_dep_.CoverTab[124]++
//line /usr/local/go/src/bytes/buffer.go:390
		// _ = "end of CoverTab[124]"
//line /usr/local/go/src/bytes/buffer.go:390
	}
//line /usr/local/go/src/bytes/buffer.go:390
	// _ = "end of CoverTab[120]"
//line /usr/local/go/src/bytes/buffer.go:390
	_go_fuzz_dep_.CoverTab[121]++
						if b.off >= int(b.lastRead) {
//line /usr/local/go/src/bytes/buffer.go:391
		_go_fuzz_dep_.CoverTab[125]++
							b.off -= int(b.lastRead)
//line /usr/local/go/src/bytes/buffer.go:392
		// _ = "end of CoverTab[125]"
	} else {
//line /usr/local/go/src/bytes/buffer.go:393
		_go_fuzz_dep_.CoverTab[126]++
//line /usr/local/go/src/bytes/buffer.go:393
		// _ = "end of CoverTab[126]"
//line /usr/local/go/src/bytes/buffer.go:393
	}
//line /usr/local/go/src/bytes/buffer.go:393
	// _ = "end of CoverTab[121]"
//line /usr/local/go/src/bytes/buffer.go:393
	_go_fuzz_dep_.CoverTab[122]++
						b.lastRead = opInvalid
						return nil
//line /usr/local/go/src/bytes/buffer.go:395
	// _ = "end of CoverTab[122]"
}

var errUnreadByte = errors.New("bytes.Buffer: UnreadByte: previous operation was not a successful read")

// UnreadByte unreads the last byte returned by the most recent successful
//line /usr/local/go/src/bytes/buffer.go:400
// read operation that read at least one byte. If a write has happened since
//line /usr/local/go/src/bytes/buffer.go:400
// the last read, if the last read returned an error, or if the read read zero
//line /usr/local/go/src/bytes/buffer.go:400
// bytes, UnreadByte returns an error.
//line /usr/local/go/src/bytes/buffer.go:404
func (b *Buffer) UnreadByte() error {
//line /usr/local/go/src/bytes/buffer.go:404
	_go_fuzz_dep_.CoverTab[127]++
						if b.lastRead == opInvalid {
//line /usr/local/go/src/bytes/buffer.go:405
		_go_fuzz_dep_.CoverTab[130]++
							return errUnreadByte
//line /usr/local/go/src/bytes/buffer.go:406
		// _ = "end of CoverTab[130]"
	} else {
//line /usr/local/go/src/bytes/buffer.go:407
		_go_fuzz_dep_.CoverTab[131]++
//line /usr/local/go/src/bytes/buffer.go:407
		// _ = "end of CoverTab[131]"
//line /usr/local/go/src/bytes/buffer.go:407
	}
//line /usr/local/go/src/bytes/buffer.go:407
	// _ = "end of CoverTab[127]"
//line /usr/local/go/src/bytes/buffer.go:407
	_go_fuzz_dep_.CoverTab[128]++
						b.lastRead = opInvalid
						if b.off > 0 {
//line /usr/local/go/src/bytes/buffer.go:409
		_go_fuzz_dep_.CoverTab[132]++
							b.off--
//line /usr/local/go/src/bytes/buffer.go:410
		// _ = "end of CoverTab[132]"
	} else {
//line /usr/local/go/src/bytes/buffer.go:411
		_go_fuzz_dep_.CoverTab[133]++
//line /usr/local/go/src/bytes/buffer.go:411
		// _ = "end of CoverTab[133]"
//line /usr/local/go/src/bytes/buffer.go:411
	}
//line /usr/local/go/src/bytes/buffer.go:411
	// _ = "end of CoverTab[128]"
//line /usr/local/go/src/bytes/buffer.go:411
	_go_fuzz_dep_.CoverTab[129]++
						return nil
//line /usr/local/go/src/bytes/buffer.go:412
	// _ = "end of CoverTab[129]"
}

// ReadBytes reads until the first occurrence of delim in the input,
//line /usr/local/go/src/bytes/buffer.go:415
// returning a slice containing the data up to and including the delimiter.
//line /usr/local/go/src/bytes/buffer.go:415
// If ReadBytes encounters an error before finding a delimiter,
//line /usr/local/go/src/bytes/buffer.go:415
// it returns the data read before the error and the error itself (often io.EOF).
//line /usr/local/go/src/bytes/buffer.go:415
// ReadBytes returns err != nil if and only if the returned data does not end in
//line /usr/local/go/src/bytes/buffer.go:415
// delim.
//line /usr/local/go/src/bytes/buffer.go:421
func (b *Buffer) ReadBytes(delim byte) (line []byte, err error) {
//line /usr/local/go/src/bytes/buffer.go:421
	_go_fuzz_dep_.CoverTab[134]++
						slice, err := b.readSlice(delim)

//line /usr/local/go/src/bytes/buffer.go:425
	line = append(line, slice...)
						return line, err
//line /usr/local/go/src/bytes/buffer.go:426
	// _ = "end of CoverTab[134]"
}

// readSlice is like ReadBytes but returns a reference to internal buffer data.
func (b *Buffer) readSlice(delim byte) (line []byte, err error) {
//line /usr/local/go/src/bytes/buffer.go:430
	_go_fuzz_dep_.CoverTab[135]++
						i := IndexByte(b.buf[b.off:], delim)
						end := b.off + i + 1
						if i < 0 {
//line /usr/local/go/src/bytes/buffer.go:433
		_go_fuzz_dep_.CoverTab[137]++
							end = len(b.buf)
							err = io.EOF
//line /usr/local/go/src/bytes/buffer.go:435
		// _ = "end of CoverTab[137]"
	} else {
//line /usr/local/go/src/bytes/buffer.go:436
		_go_fuzz_dep_.CoverTab[138]++
//line /usr/local/go/src/bytes/buffer.go:436
		// _ = "end of CoverTab[138]"
//line /usr/local/go/src/bytes/buffer.go:436
	}
//line /usr/local/go/src/bytes/buffer.go:436
	// _ = "end of CoverTab[135]"
//line /usr/local/go/src/bytes/buffer.go:436
	_go_fuzz_dep_.CoverTab[136]++
						line = b.buf[b.off:end]
						b.off = end
						b.lastRead = opRead
						return line, err
//line /usr/local/go/src/bytes/buffer.go:440
	// _ = "end of CoverTab[136]"
}

// ReadString reads until the first occurrence of delim in the input,
//line /usr/local/go/src/bytes/buffer.go:443
// returning a string containing the data up to and including the delimiter.
//line /usr/local/go/src/bytes/buffer.go:443
// If ReadString encounters an error before finding a delimiter,
//line /usr/local/go/src/bytes/buffer.go:443
// it returns the data read before the error and the error itself (often io.EOF).
//line /usr/local/go/src/bytes/buffer.go:443
// ReadString returns err != nil if and only if the returned data does not end
//line /usr/local/go/src/bytes/buffer.go:443
// in delim.
//line /usr/local/go/src/bytes/buffer.go:449
func (b *Buffer) ReadString(delim byte) (line string, err error) {
//line /usr/local/go/src/bytes/buffer.go:449
	_go_fuzz_dep_.CoverTab[139]++
						slice, err := b.readSlice(delim)
						return string(slice), err
//line /usr/local/go/src/bytes/buffer.go:451
	// _ = "end of CoverTab[139]"
}

// NewBuffer creates and initializes a new Buffer using buf as its
//line /usr/local/go/src/bytes/buffer.go:454
// initial contents. The new Buffer takes ownership of buf, and the
//line /usr/local/go/src/bytes/buffer.go:454
// caller should not use buf after this call. NewBuffer is intended to
//line /usr/local/go/src/bytes/buffer.go:454
// prepare a Buffer to read existing data. It can also be used to set
//line /usr/local/go/src/bytes/buffer.go:454
// the initial size of the internal buffer for writing. To do that,
//line /usr/local/go/src/bytes/buffer.go:454
// buf should have the desired capacity but a length of zero.
//line /usr/local/go/src/bytes/buffer.go:454
//
//line /usr/local/go/src/bytes/buffer.go:454
// In most cases, new(Buffer) (or just declaring a Buffer variable) is
//line /usr/local/go/src/bytes/buffer.go:454
// sufficient to initialize a Buffer.
//line /usr/local/go/src/bytes/buffer.go:463
func NewBuffer(buf []byte) *Buffer {
//line /usr/local/go/src/bytes/buffer.go:463
	_go_fuzz_dep_.CoverTab[140]++
//line /usr/local/go/src/bytes/buffer.go:463
	return &Buffer{buf: buf}
//line /usr/local/go/src/bytes/buffer.go:463
	// _ = "end of CoverTab[140]"
//line /usr/local/go/src/bytes/buffer.go:463
}

// NewBufferString creates and initializes a new Buffer using string s as its
//line /usr/local/go/src/bytes/buffer.go:465
// initial contents. It is intended to prepare a buffer to read an existing
//line /usr/local/go/src/bytes/buffer.go:465
// string.
//line /usr/local/go/src/bytes/buffer.go:465
//
//line /usr/local/go/src/bytes/buffer.go:465
// In most cases, new(Buffer) (or just declaring a Buffer variable) is
//line /usr/local/go/src/bytes/buffer.go:465
// sufficient to initialize a Buffer.
//line /usr/local/go/src/bytes/buffer.go:471
func NewBufferString(s string) *Buffer {
//line /usr/local/go/src/bytes/buffer.go:471
	_go_fuzz_dep_.CoverTab[141]++
						return &Buffer{buf: []byte(s)}
//line /usr/local/go/src/bytes/buffer.go:472
	// _ = "end of CoverTab[141]"
}

//line /usr/local/go/src/bytes/buffer.go:473
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/bytes/buffer.go:473
var _ = _go_fuzz_dep_.CoverTab
