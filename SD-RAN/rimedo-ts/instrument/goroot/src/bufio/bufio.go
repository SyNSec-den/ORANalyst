// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/bufio/bufio.go:5
// Package bufio implements buffered I/O. It wraps an io.Reader or io.Writer
//line /usr/local/go/src/bufio/bufio.go:5
// object, creating another object (Reader or Writer) that also implements
//line /usr/local/go/src/bufio/bufio.go:5
// the interface but provides buffering and some help for textual I/O.
//line /usr/local/go/src/bufio/bufio.go:8
package bufio

//line /usr/local/go/src/bufio/bufio.go:8
import (
//line /usr/local/go/src/bufio/bufio.go:8
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/bufio/bufio.go:8
)
//line /usr/local/go/src/bufio/bufio.go:8
import (
//line /usr/local/go/src/bufio/bufio.go:8
	_atomic_ "sync/atomic"
//line /usr/local/go/src/bufio/bufio.go:8
)

import (
	"bytes"
	"errors"
	"io"
	"strings"
	"unicode/utf8"
)

const (
	defaultBufSize = 4096
)

var (
	ErrInvalidUnreadByte	= errors.New("bufio: invalid use of UnreadByte")
	ErrInvalidUnreadRune	= errors.New("bufio: invalid use of UnreadRune")
	ErrBufferFull		= errors.New("bufio: buffer full")
	ErrNegativeCount	= errors.New("bufio: negative count")
)

//line /usr/local/go/src/bufio/bufio.go:31
// Reader implements buffering for an io.Reader object.
type Reader struct {
	buf		[]byte
	rd		io.Reader	// reader provided by the client
	r, w		int		// buf read and write positions
	err		error
	lastByte	int	// last byte read for UnreadByte; -1 means invalid
	lastRuneSize	int	// size of last rune read for UnreadRune; -1 means invalid
}

const minReadBufferSize = 16
const maxConsecutiveEmptyReads = 100

// NewReaderSize returns a new Reader whose buffer has at least the specified
//line /usr/local/go/src/bufio/bufio.go:44
// size. If the argument io.Reader is already a Reader with large enough
//line /usr/local/go/src/bufio/bufio.go:44
// size, it returns the underlying Reader.
//line /usr/local/go/src/bufio/bufio.go:47
func NewReaderSize(rd io.Reader, size int) *Reader {
//line /usr/local/go/src/bufio/bufio.go:47
	_go_fuzz_dep_.CoverTab[25194]++

						b, ok := rd.(*Reader)
						if ok && func() bool {
//line /usr/local/go/src/bufio/bufio.go:50
		_go_fuzz_dep_.CoverTab[25197]++
//line /usr/local/go/src/bufio/bufio.go:50
		return len(b.buf) >= size
//line /usr/local/go/src/bufio/bufio.go:50
		// _ = "end of CoverTab[25197]"
//line /usr/local/go/src/bufio/bufio.go:50
	}() {
//line /usr/local/go/src/bufio/bufio.go:50
		_go_fuzz_dep_.CoverTab[25198]++
							return b
//line /usr/local/go/src/bufio/bufio.go:51
		// _ = "end of CoverTab[25198]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:52
		_go_fuzz_dep_.CoverTab[25199]++
//line /usr/local/go/src/bufio/bufio.go:52
		// _ = "end of CoverTab[25199]"
//line /usr/local/go/src/bufio/bufio.go:52
	}
//line /usr/local/go/src/bufio/bufio.go:52
	// _ = "end of CoverTab[25194]"
//line /usr/local/go/src/bufio/bufio.go:52
	_go_fuzz_dep_.CoverTab[25195]++
						if size < minReadBufferSize {
//line /usr/local/go/src/bufio/bufio.go:53
		_go_fuzz_dep_.CoverTab[25200]++
							size = minReadBufferSize
//line /usr/local/go/src/bufio/bufio.go:54
		// _ = "end of CoverTab[25200]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:55
		_go_fuzz_dep_.CoverTab[25201]++
//line /usr/local/go/src/bufio/bufio.go:55
		// _ = "end of CoverTab[25201]"
//line /usr/local/go/src/bufio/bufio.go:55
	}
//line /usr/local/go/src/bufio/bufio.go:55
	// _ = "end of CoverTab[25195]"
//line /usr/local/go/src/bufio/bufio.go:55
	_go_fuzz_dep_.CoverTab[25196]++
						r := new(Reader)
						r.reset(make([]byte, size), rd)
						return r
//line /usr/local/go/src/bufio/bufio.go:58
	// _ = "end of CoverTab[25196]"
}

// NewReader returns a new Reader whose buffer has the default size.
func NewReader(rd io.Reader) *Reader {
//line /usr/local/go/src/bufio/bufio.go:62
	_go_fuzz_dep_.CoverTab[25202]++
						return NewReaderSize(rd, defaultBufSize)
//line /usr/local/go/src/bufio/bufio.go:63
	// _ = "end of CoverTab[25202]"
}

// Size returns the size of the underlying buffer in bytes.
func (b *Reader) Size() int {
//line /usr/local/go/src/bufio/bufio.go:67
	_go_fuzz_dep_.CoverTab[25203]++
//line /usr/local/go/src/bufio/bufio.go:67
	return len(b.buf)
//line /usr/local/go/src/bufio/bufio.go:67
	// _ = "end of CoverTab[25203]"
//line /usr/local/go/src/bufio/bufio.go:67
}

// Reset discards any buffered data, resets all state, and switches
//line /usr/local/go/src/bufio/bufio.go:69
// the buffered reader to read from r.
//line /usr/local/go/src/bufio/bufio.go:69
// Calling Reset on the zero value of Reader initializes the internal buffer
//line /usr/local/go/src/bufio/bufio.go:69
// to the default size.
//line /usr/local/go/src/bufio/bufio.go:73
func (b *Reader) Reset(r io.Reader) {
//line /usr/local/go/src/bufio/bufio.go:73
	_go_fuzz_dep_.CoverTab[25204]++
						if b.buf == nil {
//line /usr/local/go/src/bufio/bufio.go:74
		_go_fuzz_dep_.CoverTab[25206]++
							b.buf = make([]byte, defaultBufSize)
//line /usr/local/go/src/bufio/bufio.go:75
		// _ = "end of CoverTab[25206]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:76
		_go_fuzz_dep_.CoverTab[25207]++
//line /usr/local/go/src/bufio/bufio.go:76
		// _ = "end of CoverTab[25207]"
//line /usr/local/go/src/bufio/bufio.go:76
	}
//line /usr/local/go/src/bufio/bufio.go:76
	// _ = "end of CoverTab[25204]"
//line /usr/local/go/src/bufio/bufio.go:76
	_go_fuzz_dep_.CoverTab[25205]++
						b.reset(b.buf, r)
//line /usr/local/go/src/bufio/bufio.go:77
	// _ = "end of CoverTab[25205]"
}

func (b *Reader) reset(buf []byte, r io.Reader) {
//line /usr/local/go/src/bufio/bufio.go:80
	_go_fuzz_dep_.CoverTab[25208]++
						*b = Reader{
		buf:		buf,
		rd:		r,
		lastByte:	-1,
		lastRuneSize:	-1,
	}
//line /usr/local/go/src/bufio/bufio.go:86
	// _ = "end of CoverTab[25208]"
}

var errNegativeRead = errors.New("bufio: reader returned negative count from Read")

// fill reads a new chunk into the buffer.
func (b *Reader) fill() {
//line /usr/local/go/src/bufio/bufio.go:92
	_go_fuzz_dep_.CoverTab[25209]++

						if b.r > 0 {
//line /usr/local/go/src/bufio/bufio.go:94
		_go_fuzz_dep_.CoverTab[25213]++
							copy(b.buf, b.buf[b.r:b.w])
							b.w -= b.r
							b.r = 0
//line /usr/local/go/src/bufio/bufio.go:97
		// _ = "end of CoverTab[25213]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:98
		_go_fuzz_dep_.CoverTab[25214]++
//line /usr/local/go/src/bufio/bufio.go:98
		// _ = "end of CoverTab[25214]"
//line /usr/local/go/src/bufio/bufio.go:98
	}
//line /usr/local/go/src/bufio/bufio.go:98
	// _ = "end of CoverTab[25209]"
//line /usr/local/go/src/bufio/bufio.go:98
	_go_fuzz_dep_.CoverTab[25210]++

						if b.w >= len(b.buf) {
//line /usr/local/go/src/bufio/bufio.go:100
		_go_fuzz_dep_.CoverTab[25215]++
							panic("bufio: tried to fill full buffer")
//line /usr/local/go/src/bufio/bufio.go:101
		// _ = "end of CoverTab[25215]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:102
		_go_fuzz_dep_.CoverTab[25216]++
//line /usr/local/go/src/bufio/bufio.go:102
		// _ = "end of CoverTab[25216]"
//line /usr/local/go/src/bufio/bufio.go:102
	}
//line /usr/local/go/src/bufio/bufio.go:102
	// _ = "end of CoverTab[25210]"
//line /usr/local/go/src/bufio/bufio.go:102
	_go_fuzz_dep_.CoverTab[25211]++

//line /usr/local/go/src/bufio/bufio.go:105
	for i := maxConsecutiveEmptyReads; i > 0; i-- {
//line /usr/local/go/src/bufio/bufio.go:105
		_go_fuzz_dep_.CoverTab[25217]++
							n, err := b.rd.Read(b.buf[b.w:])
							if n < 0 {
//line /usr/local/go/src/bufio/bufio.go:107
			_go_fuzz_dep_.CoverTab[25220]++
								panic(errNegativeRead)
//line /usr/local/go/src/bufio/bufio.go:108
			// _ = "end of CoverTab[25220]"
		} else {
//line /usr/local/go/src/bufio/bufio.go:109
			_go_fuzz_dep_.CoverTab[25221]++
//line /usr/local/go/src/bufio/bufio.go:109
			// _ = "end of CoverTab[25221]"
//line /usr/local/go/src/bufio/bufio.go:109
		}
//line /usr/local/go/src/bufio/bufio.go:109
		// _ = "end of CoverTab[25217]"
//line /usr/local/go/src/bufio/bufio.go:109
		_go_fuzz_dep_.CoverTab[25218]++
							b.w += n
							if err != nil {
//line /usr/local/go/src/bufio/bufio.go:111
			_go_fuzz_dep_.CoverTab[25222]++
								b.err = err
								return
//line /usr/local/go/src/bufio/bufio.go:113
			// _ = "end of CoverTab[25222]"
		} else {
//line /usr/local/go/src/bufio/bufio.go:114
			_go_fuzz_dep_.CoverTab[25223]++
//line /usr/local/go/src/bufio/bufio.go:114
			// _ = "end of CoverTab[25223]"
//line /usr/local/go/src/bufio/bufio.go:114
		}
//line /usr/local/go/src/bufio/bufio.go:114
		// _ = "end of CoverTab[25218]"
//line /usr/local/go/src/bufio/bufio.go:114
		_go_fuzz_dep_.CoverTab[25219]++
							if n > 0 {
//line /usr/local/go/src/bufio/bufio.go:115
			_go_fuzz_dep_.CoverTab[25224]++
								return
//line /usr/local/go/src/bufio/bufio.go:116
			// _ = "end of CoverTab[25224]"
		} else {
//line /usr/local/go/src/bufio/bufio.go:117
			_go_fuzz_dep_.CoverTab[25225]++
//line /usr/local/go/src/bufio/bufio.go:117
			// _ = "end of CoverTab[25225]"
//line /usr/local/go/src/bufio/bufio.go:117
		}
//line /usr/local/go/src/bufio/bufio.go:117
		// _ = "end of CoverTab[25219]"
	}
//line /usr/local/go/src/bufio/bufio.go:118
	// _ = "end of CoverTab[25211]"
//line /usr/local/go/src/bufio/bufio.go:118
	_go_fuzz_dep_.CoverTab[25212]++
						b.err = io.ErrNoProgress
//line /usr/local/go/src/bufio/bufio.go:119
	// _ = "end of CoverTab[25212]"
}

func (b *Reader) readErr() error {
//line /usr/local/go/src/bufio/bufio.go:122
	_go_fuzz_dep_.CoverTab[25226]++
						err := b.err
						b.err = nil
						return err
//line /usr/local/go/src/bufio/bufio.go:125
	// _ = "end of CoverTab[25226]"
}

// Peek returns the next n bytes without advancing the reader. The bytes stop
//line /usr/local/go/src/bufio/bufio.go:128
// being valid at the next read call. If Peek returns fewer than n bytes, it
//line /usr/local/go/src/bufio/bufio.go:128
// also returns an error explaining why the read is short. The error is
//line /usr/local/go/src/bufio/bufio.go:128
// ErrBufferFull if n is larger than b's buffer size.
//line /usr/local/go/src/bufio/bufio.go:128
//
//line /usr/local/go/src/bufio/bufio.go:128
// Calling Peek prevents a UnreadByte or UnreadRune call from succeeding
//line /usr/local/go/src/bufio/bufio.go:128
// until the next read operation.
//line /usr/local/go/src/bufio/bufio.go:135
func (b *Reader) Peek(n int) ([]byte, error) {
//line /usr/local/go/src/bufio/bufio.go:135
	_go_fuzz_dep_.CoverTab[25227]++
						if n < 0 {
//line /usr/local/go/src/bufio/bufio.go:136
		_go_fuzz_dep_.CoverTab[25232]++
							return nil, ErrNegativeCount
//line /usr/local/go/src/bufio/bufio.go:137
		// _ = "end of CoverTab[25232]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:138
		_go_fuzz_dep_.CoverTab[25233]++
//line /usr/local/go/src/bufio/bufio.go:138
		// _ = "end of CoverTab[25233]"
//line /usr/local/go/src/bufio/bufio.go:138
	}
//line /usr/local/go/src/bufio/bufio.go:138
	// _ = "end of CoverTab[25227]"
//line /usr/local/go/src/bufio/bufio.go:138
	_go_fuzz_dep_.CoverTab[25228]++

						b.lastByte = -1
						b.lastRuneSize = -1

						for b.w-b.r < n && func() bool {
//line /usr/local/go/src/bufio/bufio.go:143
		_go_fuzz_dep_.CoverTab[25234]++
//line /usr/local/go/src/bufio/bufio.go:143
		return b.w-b.r < len(b.buf)
//line /usr/local/go/src/bufio/bufio.go:143
		// _ = "end of CoverTab[25234]"
//line /usr/local/go/src/bufio/bufio.go:143
	}() && func() bool {
//line /usr/local/go/src/bufio/bufio.go:143
		_go_fuzz_dep_.CoverTab[25235]++
//line /usr/local/go/src/bufio/bufio.go:143
		return b.err == nil
//line /usr/local/go/src/bufio/bufio.go:143
		// _ = "end of CoverTab[25235]"
//line /usr/local/go/src/bufio/bufio.go:143
	}() {
//line /usr/local/go/src/bufio/bufio.go:143
		_go_fuzz_dep_.CoverTab[25236]++
							b.fill()
//line /usr/local/go/src/bufio/bufio.go:144
		// _ = "end of CoverTab[25236]"
	}
//line /usr/local/go/src/bufio/bufio.go:145
	// _ = "end of CoverTab[25228]"
//line /usr/local/go/src/bufio/bufio.go:145
	_go_fuzz_dep_.CoverTab[25229]++

						if n > len(b.buf) {
//line /usr/local/go/src/bufio/bufio.go:147
		_go_fuzz_dep_.CoverTab[25237]++
							return b.buf[b.r:b.w], ErrBufferFull
//line /usr/local/go/src/bufio/bufio.go:148
		// _ = "end of CoverTab[25237]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:149
		_go_fuzz_dep_.CoverTab[25238]++
//line /usr/local/go/src/bufio/bufio.go:149
		// _ = "end of CoverTab[25238]"
//line /usr/local/go/src/bufio/bufio.go:149
	}
//line /usr/local/go/src/bufio/bufio.go:149
	// _ = "end of CoverTab[25229]"
//line /usr/local/go/src/bufio/bufio.go:149
	_go_fuzz_dep_.CoverTab[25230]++

	// 0 <= n <= len(b.buf)
	var err error
	if avail := b.w - b.r; avail < n {
//line /usr/local/go/src/bufio/bufio.go:153
		_go_fuzz_dep_.CoverTab[25239]++

							n = avail
							err = b.readErr()
							if err == nil {
//line /usr/local/go/src/bufio/bufio.go:157
			_go_fuzz_dep_.CoverTab[25240]++
								err = ErrBufferFull
//line /usr/local/go/src/bufio/bufio.go:158
			// _ = "end of CoverTab[25240]"
		} else {
//line /usr/local/go/src/bufio/bufio.go:159
			_go_fuzz_dep_.CoverTab[25241]++
//line /usr/local/go/src/bufio/bufio.go:159
			// _ = "end of CoverTab[25241]"
//line /usr/local/go/src/bufio/bufio.go:159
		}
//line /usr/local/go/src/bufio/bufio.go:159
		// _ = "end of CoverTab[25239]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:160
		_go_fuzz_dep_.CoverTab[25242]++
//line /usr/local/go/src/bufio/bufio.go:160
		// _ = "end of CoverTab[25242]"
//line /usr/local/go/src/bufio/bufio.go:160
	}
//line /usr/local/go/src/bufio/bufio.go:160
	// _ = "end of CoverTab[25230]"
//line /usr/local/go/src/bufio/bufio.go:160
	_go_fuzz_dep_.CoverTab[25231]++
						return b.buf[b.r : b.r+n], err
//line /usr/local/go/src/bufio/bufio.go:161
	// _ = "end of CoverTab[25231]"
}

// Discard skips the next n bytes, returning the number of bytes discarded.
//line /usr/local/go/src/bufio/bufio.go:164
//
//line /usr/local/go/src/bufio/bufio.go:164
// If Discard skips fewer than n bytes, it also returns an error.
//line /usr/local/go/src/bufio/bufio.go:164
// If 0 <= n <= b.Buffered(), Discard is guaranteed to succeed without
//line /usr/local/go/src/bufio/bufio.go:164
// reading from the underlying io.Reader.
//line /usr/local/go/src/bufio/bufio.go:169
func (b *Reader) Discard(n int) (discarded int, err error) {
//line /usr/local/go/src/bufio/bufio.go:169
	_go_fuzz_dep_.CoverTab[25243]++
						if n < 0 {
//line /usr/local/go/src/bufio/bufio.go:170
		_go_fuzz_dep_.CoverTab[25246]++
							return 0, ErrNegativeCount
//line /usr/local/go/src/bufio/bufio.go:171
		// _ = "end of CoverTab[25246]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:172
		_go_fuzz_dep_.CoverTab[25247]++
//line /usr/local/go/src/bufio/bufio.go:172
		// _ = "end of CoverTab[25247]"
//line /usr/local/go/src/bufio/bufio.go:172
	}
//line /usr/local/go/src/bufio/bufio.go:172
	// _ = "end of CoverTab[25243]"
//line /usr/local/go/src/bufio/bufio.go:172
	_go_fuzz_dep_.CoverTab[25244]++
						if n == 0 {
//line /usr/local/go/src/bufio/bufio.go:173
		_go_fuzz_dep_.CoverTab[25248]++
							return
//line /usr/local/go/src/bufio/bufio.go:174
		// _ = "end of CoverTab[25248]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:175
		_go_fuzz_dep_.CoverTab[25249]++
//line /usr/local/go/src/bufio/bufio.go:175
		// _ = "end of CoverTab[25249]"
//line /usr/local/go/src/bufio/bufio.go:175
	}
//line /usr/local/go/src/bufio/bufio.go:175
	// _ = "end of CoverTab[25244]"
//line /usr/local/go/src/bufio/bufio.go:175
	_go_fuzz_dep_.CoverTab[25245]++

						b.lastByte = -1
						b.lastRuneSize = -1

						remain := n
						for {
//line /usr/local/go/src/bufio/bufio.go:181
		_go_fuzz_dep_.CoverTab[25250]++
							skip := b.Buffered()
							if skip == 0 {
//line /usr/local/go/src/bufio/bufio.go:183
			_go_fuzz_dep_.CoverTab[25254]++
								b.fill()
								skip = b.Buffered()
//line /usr/local/go/src/bufio/bufio.go:185
			// _ = "end of CoverTab[25254]"
		} else {
//line /usr/local/go/src/bufio/bufio.go:186
			_go_fuzz_dep_.CoverTab[25255]++
//line /usr/local/go/src/bufio/bufio.go:186
			// _ = "end of CoverTab[25255]"
//line /usr/local/go/src/bufio/bufio.go:186
		}
//line /usr/local/go/src/bufio/bufio.go:186
		// _ = "end of CoverTab[25250]"
//line /usr/local/go/src/bufio/bufio.go:186
		_go_fuzz_dep_.CoverTab[25251]++
							if skip > remain {
//line /usr/local/go/src/bufio/bufio.go:187
			_go_fuzz_dep_.CoverTab[25256]++
								skip = remain
//line /usr/local/go/src/bufio/bufio.go:188
			// _ = "end of CoverTab[25256]"
		} else {
//line /usr/local/go/src/bufio/bufio.go:189
			_go_fuzz_dep_.CoverTab[25257]++
//line /usr/local/go/src/bufio/bufio.go:189
			// _ = "end of CoverTab[25257]"
//line /usr/local/go/src/bufio/bufio.go:189
		}
//line /usr/local/go/src/bufio/bufio.go:189
		// _ = "end of CoverTab[25251]"
//line /usr/local/go/src/bufio/bufio.go:189
		_go_fuzz_dep_.CoverTab[25252]++
							b.r += skip
							remain -= skip
							if remain == 0 {
//line /usr/local/go/src/bufio/bufio.go:192
			_go_fuzz_dep_.CoverTab[25258]++
								return n, nil
//line /usr/local/go/src/bufio/bufio.go:193
			// _ = "end of CoverTab[25258]"
		} else {
//line /usr/local/go/src/bufio/bufio.go:194
			_go_fuzz_dep_.CoverTab[25259]++
//line /usr/local/go/src/bufio/bufio.go:194
			// _ = "end of CoverTab[25259]"
//line /usr/local/go/src/bufio/bufio.go:194
		}
//line /usr/local/go/src/bufio/bufio.go:194
		// _ = "end of CoverTab[25252]"
//line /usr/local/go/src/bufio/bufio.go:194
		_go_fuzz_dep_.CoverTab[25253]++
							if b.err != nil {
//line /usr/local/go/src/bufio/bufio.go:195
			_go_fuzz_dep_.CoverTab[25260]++
								return n - remain, b.readErr()
//line /usr/local/go/src/bufio/bufio.go:196
			// _ = "end of CoverTab[25260]"
		} else {
//line /usr/local/go/src/bufio/bufio.go:197
			_go_fuzz_dep_.CoverTab[25261]++
//line /usr/local/go/src/bufio/bufio.go:197
			// _ = "end of CoverTab[25261]"
//line /usr/local/go/src/bufio/bufio.go:197
		}
//line /usr/local/go/src/bufio/bufio.go:197
		// _ = "end of CoverTab[25253]"
	}
//line /usr/local/go/src/bufio/bufio.go:198
	// _ = "end of CoverTab[25245]"
}

// Read reads data into p.
//line /usr/local/go/src/bufio/bufio.go:201
// It returns the number of bytes read into p.
//line /usr/local/go/src/bufio/bufio.go:201
// The bytes are taken from at most one Read on the underlying Reader,
//line /usr/local/go/src/bufio/bufio.go:201
// hence n may be less than len(p).
//line /usr/local/go/src/bufio/bufio.go:201
// To read exactly len(p) bytes, use io.ReadFull(b, p).
//line /usr/local/go/src/bufio/bufio.go:201
// If the underlying Reader can return a non-zero count with io.EOF,
//line /usr/local/go/src/bufio/bufio.go:201
// then this Read method can do so as well; see the [io.Reader] docs.
//line /usr/local/go/src/bufio/bufio.go:208
func (b *Reader) Read(p []byte) (n int, err error) {
//line /usr/local/go/src/bufio/bufio.go:208
	_go_fuzz_dep_.CoverTab[25262]++
						n = len(p)
						if n == 0 {
//line /usr/local/go/src/bufio/bufio.go:210
		_go_fuzz_dep_.CoverTab[25265]++
							if b.Buffered() > 0 {
//line /usr/local/go/src/bufio/bufio.go:211
			_go_fuzz_dep_.CoverTab[25267]++
								return 0, nil
//line /usr/local/go/src/bufio/bufio.go:212
			// _ = "end of CoverTab[25267]"
		} else {
//line /usr/local/go/src/bufio/bufio.go:213
			_go_fuzz_dep_.CoverTab[25268]++
//line /usr/local/go/src/bufio/bufio.go:213
			// _ = "end of CoverTab[25268]"
//line /usr/local/go/src/bufio/bufio.go:213
		}
//line /usr/local/go/src/bufio/bufio.go:213
		// _ = "end of CoverTab[25265]"
//line /usr/local/go/src/bufio/bufio.go:213
		_go_fuzz_dep_.CoverTab[25266]++
							return 0, b.readErr()
//line /usr/local/go/src/bufio/bufio.go:214
		// _ = "end of CoverTab[25266]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:215
		_go_fuzz_dep_.CoverTab[25269]++
//line /usr/local/go/src/bufio/bufio.go:215
		// _ = "end of CoverTab[25269]"
//line /usr/local/go/src/bufio/bufio.go:215
	}
//line /usr/local/go/src/bufio/bufio.go:215
	// _ = "end of CoverTab[25262]"
//line /usr/local/go/src/bufio/bufio.go:215
	_go_fuzz_dep_.CoverTab[25263]++
						if b.r == b.w {
//line /usr/local/go/src/bufio/bufio.go:216
		_go_fuzz_dep_.CoverTab[25270]++
							if b.err != nil {
//line /usr/local/go/src/bufio/bufio.go:217
			_go_fuzz_dep_.CoverTab[25275]++
								return 0, b.readErr()
//line /usr/local/go/src/bufio/bufio.go:218
			// _ = "end of CoverTab[25275]"
		} else {
//line /usr/local/go/src/bufio/bufio.go:219
			_go_fuzz_dep_.CoverTab[25276]++
//line /usr/local/go/src/bufio/bufio.go:219
			// _ = "end of CoverTab[25276]"
//line /usr/local/go/src/bufio/bufio.go:219
		}
//line /usr/local/go/src/bufio/bufio.go:219
		// _ = "end of CoverTab[25270]"
//line /usr/local/go/src/bufio/bufio.go:219
		_go_fuzz_dep_.CoverTab[25271]++
							if len(p) >= len(b.buf) {
//line /usr/local/go/src/bufio/bufio.go:220
			_go_fuzz_dep_.CoverTab[25277]++

//line /usr/local/go/src/bufio/bufio.go:223
			n, b.err = b.rd.Read(p)
			if n < 0 {
//line /usr/local/go/src/bufio/bufio.go:224
				_go_fuzz_dep_.CoverTab[25280]++
									panic(errNegativeRead)
//line /usr/local/go/src/bufio/bufio.go:225
				// _ = "end of CoverTab[25280]"
			} else {
//line /usr/local/go/src/bufio/bufio.go:226
				_go_fuzz_dep_.CoverTab[25281]++
//line /usr/local/go/src/bufio/bufio.go:226
				// _ = "end of CoverTab[25281]"
//line /usr/local/go/src/bufio/bufio.go:226
			}
//line /usr/local/go/src/bufio/bufio.go:226
			// _ = "end of CoverTab[25277]"
//line /usr/local/go/src/bufio/bufio.go:226
			_go_fuzz_dep_.CoverTab[25278]++
								if n > 0 {
//line /usr/local/go/src/bufio/bufio.go:227
				_go_fuzz_dep_.CoverTab[25282]++
									b.lastByte = int(p[n-1])
									b.lastRuneSize = -1
//line /usr/local/go/src/bufio/bufio.go:229
				// _ = "end of CoverTab[25282]"
			} else {
//line /usr/local/go/src/bufio/bufio.go:230
				_go_fuzz_dep_.CoverTab[25283]++
//line /usr/local/go/src/bufio/bufio.go:230
				// _ = "end of CoverTab[25283]"
//line /usr/local/go/src/bufio/bufio.go:230
			}
//line /usr/local/go/src/bufio/bufio.go:230
			// _ = "end of CoverTab[25278]"
//line /usr/local/go/src/bufio/bufio.go:230
			_go_fuzz_dep_.CoverTab[25279]++
								return n, b.readErr()
//line /usr/local/go/src/bufio/bufio.go:231
			// _ = "end of CoverTab[25279]"
		} else {
//line /usr/local/go/src/bufio/bufio.go:232
			_go_fuzz_dep_.CoverTab[25284]++
//line /usr/local/go/src/bufio/bufio.go:232
			// _ = "end of CoverTab[25284]"
//line /usr/local/go/src/bufio/bufio.go:232
		}
//line /usr/local/go/src/bufio/bufio.go:232
		// _ = "end of CoverTab[25271]"
//line /usr/local/go/src/bufio/bufio.go:232
		_go_fuzz_dep_.CoverTab[25272]++

//line /usr/local/go/src/bufio/bufio.go:235
		b.r = 0
		b.w = 0
		n, b.err = b.rd.Read(b.buf)
		if n < 0 {
//line /usr/local/go/src/bufio/bufio.go:238
			_go_fuzz_dep_.CoverTab[25285]++
								panic(errNegativeRead)
//line /usr/local/go/src/bufio/bufio.go:239
			// _ = "end of CoverTab[25285]"
		} else {
//line /usr/local/go/src/bufio/bufio.go:240
			_go_fuzz_dep_.CoverTab[25286]++
//line /usr/local/go/src/bufio/bufio.go:240
			// _ = "end of CoverTab[25286]"
//line /usr/local/go/src/bufio/bufio.go:240
		}
//line /usr/local/go/src/bufio/bufio.go:240
		// _ = "end of CoverTab[25272]"
//line /usr/local/go/src/bufio/bufio.go:240
		_go_fuzz_dep_.CoverTab[25273]++
							if n == 0 {
//line /usr/local/go/src/bufio/bufio.go:241
			_go_fuzz_dep_.CoverTab[25287]++
								return 0, b.readErr()
//line /usr/local/go/src/bufio/bufio.go:242
			// _ = "end of CoverTab[25287]"
		} else {
//line /usr/local/go/src/bufio/bufio.go:243
			_go_fuzz_dep_.CoverTab[25288]++
//line /usr/local/go/src/bufio/bufio.go:243
			// _ = "end of CoverTab[25288]"
//line /usr/local/go/src/bufio/bufio.go:243
		}
//line /usr/local/go/src/bufio/bufio.go:243
		// _ = "end of CoverTab[25273]"
//line /usr/local/go/src/bufio/bufio.go:243
		_go_fuzz_dep_.CoverTab[25274]++
							b.w += n
//line /usr/local/go/src/bufio/bufio.go:244
		// _ = "end of CoverTab[25274]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:245
		_go_fuzz_dep_.CoverTab[25289]++
//line /usr/local/go/src/bufio/bufio.go:245
		// _ = "end of CoverTab[25289]"
//line /usr/local/go/src/bufio/bufio.go:245
	}
//line /usr/local/go/src/bufio/bufio.go:245
	// _ = "end of CoverTab[25263]"
//line /usr/local/go/src/bufio/bufio.go:245
	_go_fuzz_dep_.CoverTab[25264]++

//line /usr/local/go/src/bufio/bufio.go:250
	n = copy(p, b.buf[b.r:b.w])
						b.r += n
						b.lastByte = int(b.buf[b.r-1])
						b.lastRuneSize = -1
						return n, nil
//line /usr/local/go/src/bufio/bufio.go:254
	// _ = "end of CoverTab[25264]"
}

// ReadByte reads and returns a single byte.
//line /usr/local/go/src/bufio/bufio.go:257
// If no byte is available, returns an error.
//line /usr/local/go/src/bufio/bufio.go:259
func (b *Reader) ReadByte() (byte, error) {
//line /usr/local/go/src/bufio/bufio.go:259
	_go_fuzz_dep_.CoverTab[25290]++
						b.lastRuneSize = -1
						for b.r == b.w {
//line /usr/local/go/src/bufio/bufio.go:261
		_go_fuzz_dep_.CoverTab[25292]++
							if b.err != nil {
//line /usr/local/go/src/bufio/bufio.go:262
			_go_fuzz_dep_.CoverTab[25294]++
								return 0, b.readErr()
//line /usr/local/go/src/bufio/bufio.go:263
			// _ = "end of CoverTab[25294]"
		} else {
//line /usr/local/go/src/bufio/bufio.go:264
			_go_fuzz_dep_.CoverTab[25295]++
//line /usr/local/go/src/bufio/bufio.go:264
			// _ = "end of CoverTab[25295]"
//line /usr/local/go/src/bufio/bufio.go:264
		}
//line /usr/local/go/src/bufio/bufio.go:264
		// _ = "end of CoverTab[25292]"
//line /usr/local/go/src/bufio/bufio.go:264
		_go_fuzz_dep_.CoverTab[25293]++
							b.fill()
//line /usr/local/go/src/bufio/bufio.go:265
		// _ = "end of CoverTab[25293]"
	}
//line /usr/local/go/src/bufio/bufio.go:266
	// _ = "end of CoverTab[25290]"
//line /usr/local/go/src/bufio/bufio.go:266
	_go_fuzz_dep_.CoverTab[25291]++
						c := b.buf[b.r]
						b.r++
						b.lastByte = int(c)
						return c, nil
//line /usr/local/go/src/bufio/bufio.go:270
	// _ = "end of CoverTab[25291]"
}

// UnreadByte unreads the last byte. Only the most recently read byte can be unread.
//line /usr/local/go/src/bufio/bufio.go:273
//
//line /usr/local/go/src/bufio/bufio.go:273
// UnreadByte returns an error if the most recent method called on the
//line /usr/local/go/src/bufio/bufio.go:273
// Reader was not a read operation. Notably, Peek, Discard, and WriteTo are not
//line /usr/local/go/src/bufio/bufio.go:273
// considered read operations.
//line /usr/local/go/src/bufio/bufio.go:278
func (b *Reader) UnreadByte() error {
//line /usr/local/go/src/bufio/bufio.go:278
	_go_fuzz_dep_.CoverTab[25296]++
						if b.lastByte < 0 || func() bool {
//line /usr/local/go/src/bufio/bufio.go:279
		_go_fuzz_dep_.CoverTab[25299]++
//line /usr/local/go/src/bufio/bufio.go:279
		return b.r == 0 && func() bool {
//line /usr/local/go/src/bufio/bufio.go:279
			_go_fuzz_dep_.CoverTab[25300]++
//line /usr/local/go/src/bufio/bufio.go:279
			return b.w > 0
//line /usr/local/go/src/bufio/bufio.go:279
			// _ = "end of CoverTab[25300]"
//line /usr/local/go/src/bufio/bufio.go:279
		}()
//line /usr/local/go/src/bufio/bufio.go:279
		// _ = "end of CoverTab[25299]"
//line /usr/local/go/src/bufio/bufio.go:279
	}() {
//line /usr/local/go/src/bufio/bufio.go:279
		_go_fuzz_dep_.CoverTab[25301]++
							return ErrInvalidUnreadByte
//line /usr/local/go/src/bufio/bufio.go:280
		// _ = "end of CoverTab[25301]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:281
		_go_fuzz_dep_.CoverTab[25302]++
//line /usr/local/go/src/bufio/bufio.go:281
		// _ = "end of CoverTab[25302]"
//line /usr/local/go/src/bufio/bufio.go:281
	}
//line /usr/local/go/src/bufio/bufio.go:281
	// _ = "end of CoverTab[25296]"
//line /usr/local/go/src/bufio/bufio.go:281
	_go_fuzz_dep_.CoverTab[25297]++

						if b.r > 0 {
//line /usr/local/go/src/bufio/bufio.go:283
		_go_fuzz_dep_.CoverTab[25303]++
							b.r--
//line /usr/local/go/src/bufio/bufio.go:284
		// _ = "end of CoverTab[25303]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:285
		_go_fuzz_dep_.CoverTab[25304]++

							b.w = 1
//line /usr/local/go/src/bufio/bufio.go:287
		// _ = "end of CoverTab[25304]"
	}
//line /usr/local/go/src/bufio/bufio.go:288
	// _ = "end of CoverTab[25297]"
//line /usr/local/go/src/bufio/bufio.go:288
	_go_fuzz_dep_.CoverTab[25298]++
						b.buf[b.r] = byte(b.lastByte)
						b.lastByte = -1
						b.lastRuneSize = -1
						return nil
//line /usr/local/go/src/bufio/bufio.go:292
	// _ = "end of CoverTab[25298]"
}

// ReadRune reads a single UTF-8 encoded Unicode character and returns the
//line /usr/local/go/src/bufio/bufio.go:295
// rune and its size in bytes. If the encoded rune is invalid, it consumes one byte
//line /usr/local/go/src/bufio/bufio.go:295
// and returns unicode.ReplacementChar (U+FFFD) with a size of 1.
//line /usr/local/go/src/bufio/bufio.go:298
func (b *Reader) ReadRune() (r rune, size int, err error) {
//line /usr/local/go/src/bufio/bufio.go:298
	_go_fuzz_dep_.CoverTab[25305]++
						for b.r+utf8.UTFMax > b.w && func() bool {
//line /usr/local/go/src/bufio/bufio.go:299
		_go_fuzz_dep_.CoverTab[25309]++
//line /usr/local/go/src/bufio/bufio.go:299
		return !utf8.FullRune(b.buf[b.r:b.w])
//line /usr/local/go/src/bufio/bufio.go:299
		// _ = "end of CoverTab[25309]"
//line /usr/local/go/src/bufio/bufio.go:299
	}() && func() bool {
//line /usr/local/go/src/bufio/bufio.go:299
		_go_fuzz_dep_.CoverTab[25310]++
//line /usr/local/go/src/bufio/bufio.go:299
		return b.err == nil
//line /usr/local/go/src/bufio/bufio.go:299
		// _ = "end of CoverTab[25310]"
//line /usr/local/go/src/bufio/bufio.go:299
	}() && func() bool {
//line /usr/local/go/src/bufio/bufio.go:299
		_go_fuzz_dep_.CoverTab[25311]++
//line /usr/local/go/src/bufio/bufio.go:299
		return b.w-b.r < len(b.buf)
//line /usr/local/go/src/bufio/bufio.go:299
		// _ = "end of CoverTab[25311]"
//line /usr/local/go/src/bufio/bufio.go:299
	}() {
//line /usr/local/go/src/bufio/bufio.go:299
		_go_fuzz_dep_.CoverTab[25312]++
							b.fill()
//line /usr/local/go/src/bufio/bufio.go:300
		// _ = "end of CoverTab[25312]"
	}
//line /usr/local/go/src/bufio/bufio.go:301
	// _ = "end of CoverTab[25305]"
//line /usr/local/go/src/bufio/bufio.go:301
	_go_fuzz_dep_.CoverTab[25306]++
						b.lastRuneSize = -1
						if b.r == b.w {
//line /usr/local/go/src/bufio/bufio.go:303
		_go_fuzz_dep_.CoverTab[25313]++
							return 0, 0, b.readErr()
//line /usr/local/go/src/bufio/bufio.go:304
		// _ = "end of CoverTab[25313]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:305
		_go_fuzz_dep_.CoverTab[25314]++
//line /usr/local/go/src/bufio/bufio.go:305
		// _ = "end of CoverTab[25314]"
//line /usr/local/go/src/bufio/bufio.go:305
	}
//line /usr/local/go/src/bufio/bufio.go:305
	// _ = "end of CoverTab[25306]"
//line /usr/local/go/src/bufio/bufio.go:305
	_go_fuzz_dep_.CoverTab[25307]++
						r, size = rune(b.buf[b.r]), 1
						if r >= utf8.RuneSelf {
//line /usr/local/go/src/bufio/bufio.go:307
		_go_fuzz_dep_.CoverTab[25315]++
							r, size = utf8.DecodeRune(b.buf[b.r:b.w])
//line /usr/local/go/src/bufio/bufio.go:308
		// _ = "end of CoverTab[25315]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:309
		_go_fuzz_dep_.CoverTab[25316]++
//line /usr/local/go/src/bufio/bufio.go:309
		// _ = "end of CoverTab[25316]"
//line /usr/local/go/src/bufio/bufio.go:309
	}
//line /usr/local/go/src/bufio/bufio.go:309
	// _ = "end of CoverTab[25307]"
//line /usr/local/go/src/bufio/bufio.go:309
	_go_fuzz_dep_.CoverTab[25308]++
						b.r += size
						b.lastByte = int(b.buf[b.r-1])
						b.lastRuneSize = size
						return r, size, nil
//line /usr/local/go/src/bufio/bufio.go:313
	// _ = "end of CoverTab[25308]"
}

// UnreadRune unreads the last rune. If the most recent method called on
//line /usr/local/go/src/bufio/bufio.go:316
// the Reader was not a ReadRune, UnreadRune returns an error. (In this
//line /usr/local/go/src/bufio/bufio.go:316
// regard it is stricter than UnreadByte, which will unread the last byte
//line /usr/local/go/src/bufio/bufio.go:316
// from any read operation.)
//line /usr/local/go/src/bufio/bufio.go:320
func (b *Reader) UnreadRune() error {
//line /usr/local/go/src/bufio/bufio.go:320
	_go_fuzz_dep_.CoverTab[25317]++
						if b.lastRuneSize < 0 || func() bool {
//line /usr/local/go/src/bufio/bufio.go:321
		_go_fuzz_dep_.CoverTab[25319]++
//line /usr/local/go/src/bufio/bufio.go:321
		return b.r < b.lastRuneSize
//line /usr/local/go/src/bufio/bufio.go:321
		// _ = "end of CoverTab[25319]"
//line /usr/local/go/src/bufio/bufio.go:321
	}() {
//line /usr/local/go/src/bufio/bufio.go:321
		_go_fuzz_dep_.CoverTab[25320]++
							return ErrInvalidUnreadRune
//line /usr/local/go/src/bufio/bufio.go:322
		// _ = "end of CoverTab[25320]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:323
		_go_fuzz_dep_.CoverTab[25321]++
//line /usr/local/go/src/bufio/bufio.go:323
		// _ = "end of CoverTab[25321]"
//line /usr/local/go/src/bufio/bufio.go:323
	}
//line /usr/local/go/src/bufio/bufio.go:323
	// _ = "end of CoverTab[25317]"
//line /usr/local/go/src/bufio/bufio.go:323
	_go_fuzz_dep_.CoverTab[25318]++
						b.r -= b.lastRuneSize
						b.lastByte = -1
						b.lastRuneSize = -1
						return nil
//line /usr/local/go/src/bufio/bufio.go:327
	// _ = "end of CoverTab[25318]"
}

// Buffered returns the number of bytes that can be read from the current buffer.
func (b *Reader) Buffered() int {
//line /usr/local/go/src/bufio/bufio.go:331
	_go_fuzz_dep_.CoverTab[25322]++
//line /usr/local/go/src/bufio/bufio.go:331
	return b.w - b.r
//line /usr/local/go/src/bufio/bufio.go:331
	// _ = "end of CoverTab[25322]"
//line /usr/local/go/src/bufio/bufio.go:331
}

// ReadSlice reads until the first occurrence of delim in the input,
//line /usr/local/go/src/bufio/bufio.go:333
// returning a slice pointing at the bytes in the buffer.
//line /usr/local/go/src/bufio/bufio.go:333
// The bytes stop being valid at the next read.
//line /usr/local/go/src/bufio/bufio.go:333
// If ReadSlice encounters an error before finding a delimiter,
//line /usr/local/go/src/bufio/bufio.go:333
// it returns all the data in the buffer and the error itself (often io.EOF).
//line /usr/local/go/src/bufio/bufio.go:333
// ReadSlice fails with error ErrBufferFull if the buffer fills without a delim.
//line /usr/local/go/src/bufio/bufio.go:333
// Because the data returned from ReadSlice will be overwritten
//line /usr/local/go/src/bufio/bufio.go:333
// by the next I/O operation, most clients should use
//line /usr/local/go/src/bufio/bufio.go:333
// ReadBytes or ReadString instead.
//line /usr/local/go/src/bufio/bufio.go:333
// ReadSlice returns err != nil if and only if line does not end in delim.
//line /usr/local/go/src/bufio/bufio.go:343
func (b *Reader) ReadSlice(delim byte) (line []byte, err error) {
//line /usr/local/go/src/bufio/bufio.go:343
	_go_fuzz_dep_.CoverTab[25323]++
						s := 0
						for {
//line /usr/local/go/src/bufio/bufio.go:345
		_go_fuzz_dep_.CoverTab[25326]++

							if i := bytes.IndexByte(b.buf[b.r+s:b.w], delim); i >= 0 {
//line /usr/local/go/src/bufio/bufio.go:347
			_go_fuzz_dep_.CoverTab[25330]++
								i += s
								line = b.buf[b.r : b.r+i+1]
								b.r += i + 1
								break
//line /usr/local/go/src/bufio/bufio.go:351
			// _ = "end of CoverTab[25330]"
		} else {
//line /usr/local/go/src/bufio/bufio.go:352
			_go_fuzz_dep_.CoverTab[25331]++
//line /usr/local/go/src/bufio/bufio.go:352
			// _ = "end of CoverTab[25331]"
//line /usr/local/go/src/bufio/bufio.go:352
		}
//line /usr/local/go/src/bufio/bufio.go:352
		// _ = "end of CoverTab[25326]"
//line /usr/local/go/src/bufio/bufio.go:352
		_go_fuzz_dep_.CoverTab[25327]++

//line /usr/local/go/src/bufio/bufio.go:355
		if b.err != nil {
//line /usr/local/go/src/bufio/bufio.go:355
			_go_fuzz_dep_.CoverTab[25332]++
								line = b.buf[b.r:b.w]
								b.r = b.w
								err = b.readErr()
								break
//line /usr/local/go/src/bufio/bufio.go:359
			// _ = "end of CoverTab[25332]"
		} else {
//line /usr/local/go/src/bufio/bufio.go:360
			_go_fuzz_dep_.CoverTab[25333]++
//line /usr/local/go/src/bufio/bufio.go:360
			// _ = "end of CoverTab[25333]"
//line /usr/local/go/src/bufio/bufio.go:360
		}
//line /usr/local/go/src/bufio/bufio.go:360
		// _ = "end of CoverTab[25327]"
//line /usr/local/go/src/bufio/bufio.go:360
		_go_fuzz_dep_.CoverTab[25328]++

//line /usr/local/go/src/bufio/bufio.go:363
		if b.Buffered() >= len(b.buf) {
//line /usr/local/go/src/bufio/bufio.go:363
			_go_fuzz_dep_.CoverTab[25334]++
								b.r = b.w
								line = b.buf
								err = ErrBufferFull
								break
//line /usr/local/go/src/bufio/bufio.go:367
			// _ = "end of CoverTab[25334]"
		} else {
//line /usr/local/go/src/bufio/bufio.go:368
			_go_fuzz_dep_.CoverTab[25335]++
//line /usr/local/go/src/bufio/bufio.go:368
			// _ = "end of CoverTab[25335]"
//line /usr/local/go/src/bufio/bufio.go:368
		}
//line /usr/local/go/src/bufio/bufio.go:368
		// _ = "end of CoverTab[25328]"
//line /usr/local/go/src/bufio/bufio.go:368
		_go_fuzz_dep_.CoverTab[25329]++

							s = b.w - b.r

							b.fill()
//line /usr/local/go/src/bufio/bufio.go:372
		// _ = "end of CoverTab[25329]"
	}
//line /usr/local/go/src/bufio/bufio.go:373
	// _ = "end of CoverTab[25323]"
//line /usr/local/go/src/bufio/bufio.go:373
	_go_fuzz_dep_.CoverTab[25324]++

//line /usr/local/go/src/bufio/bufio.go:376
	if i := len(line) - 1; i >= 0 {
//line /usr/local/go/src/bufio/bufio.go:376
		_go_fuzz_dep_.CoverTab[25336]++
							b.lastByte = int(line[i])
							b.lastRuneSize = -1
//line /usr/local/go/src/bufio/bufio.go:378
		// _ = "end of CoverTab[25336]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:379
		_go_fuzz_dep_.CoverTab[25337]++
//line /usr/local/go/src/bufio/bufio.go:379
		// _ = "end of CoverTab[25337]"
//line /usr/local/go/src/bufio/bufio.go:379
	}
//line /usr/local/go/src/bufio/bufio.go:379
	// _ = "end of CoverTab[25324]"
//line /usr/local/go/src/bufio/bufio.go:379
	_go_fuzz_dep_.CoverTab[25325]++

						return
//line /usr/local/go/src/bufio/bufio.go:381
	// _ = "end of CoverTab[25325]"
}

// ReadLine is a low-level line-reading primitive. Most callers should use
//line /usr/local/go/src/bufio/bufio.go:384
// ReadBytes('\n') or ReadString('\n') instead or use a Scanner.
//line /usr/local/go/src/bufio/bufio.go:384
//
//line /usr/local/go/src/bufio/bufio.go:384
// ReadLine tries to return a single line, not including the end-of-line bytes.
//line /usr/local/go/src/bufio/bufio.go:384
// If the line was too long for the buffer then isPrefix is set and the
//line /usr/local/go/src/bufio/bufio.go:384
// beginning of the line is returned. The rest of the line will be returned
//line /usr/local/go/src/bufio/bufio.go:384
// from future calls. isPrefix will be false when returning the last fragment
//line /usr/local/go/src/bufio/bufio.go:384
// of the line. The returned buffer is only valid until the next call to
//line /usr/local/go/src/bufio/bufio.go:384
// ReadLine. ReadLine either returns a non-nil line or it returns an error,
//line /usr/local/go/src/bufio/bufio.go:384
// never both.
//line /usr/local/go/src/bufio/bufio.go:384
//
//line /usr/local/go/src/bufio/bufio.go:384
// The text returned from ReadLine does not include the line end ("\r\n" or "\n").
//line /usr/local/go/src/bufio/bufio.go:384
// No indication or error is given if the input ends without a final line end.
//line /usr/local/go/src/bufio/bufio.go:384
// Calling UnreadByte after ReadLine will always unread the last byte read
//line /usr/local/go/src/bufio/bufio.go:384
// (possibly a character belonging to the line end) even if that byte is not
//line /usr/local/go/src/bufio/bufio.go:384
// part of the line returned by ReadLine.
//line /usr/local/go/src/bufio/bufio.go:400
func (b *Reader) ReadLine() (line []byte, isPrefix bool, err error) {
//line /usr/local/go/src/bufio/bufio.go:400
	_go_fuzz_dep_.CoverTab[25338]++
						line, err = b.ReadSlice('\n')
						if err == ErrBufferFull {
//line /usr/local/go/src/bufio/bufio.go:402
		_go_fuzz_dep_.CoverTab[25342]++

							if len(line) > 0 && func() bool {
//line /usr/local/go/src/bufio/bufio.go:404
			_go_fuzz_dep_.CoverTab[25344]++
//line /usr/local/go/src/bufio/bufio.go:404
			return line[len(line)-1] == '\r'
//line /usr/local/go/src/bufio/bufio.go:404
			// _ = "end of CoverTab[25344]"
//line /usr/local/go/src/bufio/bufio.go:404
		}() {
//line /usr/local/go/src/bufio/bufio.go:404
			_go_fuzz_dep_.CoverTab[25345]++

//line /usr/local/go/src/bufio/bufio.go:407
			if b.r == 0 {
//line /usr/local/go/src/bufio/bufio.go:407
				_go_fuzz_dep_.CoverTab[25347]++

									panic("bufio: tried to rewind past start of buffer")
//line /usr/local/go/src/bufio/bufio.go:409
				// _ = "end of CoverTab[25347]"
			} else {
//line /usr/local/go/src/bufio/bufio.go:410
				_go_fuzz_dep_.CoverTab[25348]++
//line /usr/local/go/src/bufio/bufio.go:410
				// _ = "end of CoverTab[25348]"
//line /usr/local/go/src/bufio/bufio.go:410
			}
//line /usr/local/go/src/bufio/bufio.go:410
			// _ = "end of CoverTab[25345]"
//line /usr/local/go/src/bufio/bufio.go:410
			_go_fuzz_dep_.CoverTab[25346]++
								b.r--
								line = line[:len(line)-1]
//line /usr/local/go/src/bufio/bufio.go:412
			// _ = "end of CoverTab[25346]"
		} else {
//line /usr/local/go/src/bufio/bufio.go:413
			_go_fuzz_dep_.CoverTab[25349]++
//line /usr/local/go/src/bufio/bufio.go:413
			// _ = "end of CoverTab[25349]"
//line /usr/local/go/src/bufio/bufio.go:413
		}
//line /usr/local/go/src/bufio/bufio.go:413
		// _ = "end of CoverTab[25342]"
//line /usr/local/go/src/bufio/bufio.go:413
		_go_fuzz_dep_.CoverTab[25343]++
							return line, true, nil
//line /usr/local/go/src/bufio/bufio.go:414
		// _ = "end of CoverTab[25343]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:415
		_go_fuzz_dep_.CoverTab[25350]++
//line /usr/local/go/src/bufio/bufio.go:415
		// _ = "end of CoverTab[25350]"
//line /usr/local/go/src/bufio/bufio.go:415
	}
//line /usr/local/go/src/bufio/bufio.go:415
	// _ = "end of CoverTab[25338]"
//line /usr/local/go/src/bufio/bufio.go:415
	_go_fuzz_dep_.CoverTab[25339]++

						if len(line) == 0 {
//line /usr/local/go/src/bufio/bufio.go:417
		_go_fuzz_dep_.CoverTab[25351]++
							if err != nil {
//line /usr/local/go/src/bufio/bufio.go:418
			_go_fuzz_dep_.CoverTab[25353]++
								line = nil
//line /usr/local/go/src/bufio/bufio.go:419
			// _ = "end of CoverTab[25353]"
		} else {
//line /usr/local/go/src/bufio/bufio.go:420
			_go_fuzz_dep_.CoverTab[25354]++
//line /usr/local/go/src/bufio/bufio.go:420
			// _ = "end of CoverTab[25354]"
//line /usr/local/go/src/bufio/bufio.go:420
		}
//line /usr/local/go/src/bufio/bufio.go:420
		// _ = "end of CoverTab[25351]"
//line /usr/local/go/src/bufio/bufio.go:420
		_go_fuzz_dep_.CoverTab[25352]++
							return
//line /usr/local/go/src/bufio/bufio.go:421
		// _ = "end of CoverTab[25352]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:422
		_go_fuzz_dep_.CoverTab[25355]++
//line /usr/local/go/src/bufio/bufio.go:422
		// _ = "end of CoverTab[25355]"
//line /usr/local/go/src/bufio/bufio.go:422
	}
//line /usr/local/go/src/bufio/bufio.go:422
	// _ = "end of CoverTab[25339]"
//line /usr/local/go/src/bufio/bufio.go:422
	_go_fuzz_dep_.CoverTab[25340]++
						err = nil

						if line[len(line)-1] == '\n' {
//line /usr/local/go/src/bufio/bufio.go:425
		_go_fuzz_dep_.CoverTab[25356]++
							drop := 1
							if len(line) > 1 && func() bool {
//line /usr/local/go/src/bufio/bufio.go:427
			_go_fuzz_dep_.CoverTab[25358]++
//line /usr/local/go/src/bufio/bufio.go:427
			return line[len(line)-2] == '\r'
//line /usr/local/go/src/bufio/bufio.go:427
			// _ = "end of CoverTab[25358]"
//line /usr/local/go/src/bufio/bufio.go:427
		}() {
//line /usr/local/go/src/bufio/bufio.go:427
			_go_fuzz_dep_.CoverTab[25359]++
								drop = 2
//line /usr/local/go/src/bufio/bufio.go:428
			// _ = "end of CoverTab[25359]"
		} else {
//line /usr/local/go/src/bufio/bufio.go:429
			_go_fuzz_dep_.CoverTab[25360]++
//line /usr/local/go/src/bufio/bufio.go:429
			// _ = "end of CoverTab[25360]"
//line /usr/local/go/src/bufio/bufio.go:429
		}
//line /usr/local/go/src/bufio/bufio.go:429
		// _ = "end of CoverTab[25356]"
//line /usr/local/go/src/bufio/bufio.go:429
		_go_fuzz_dep_.CoverTab[25357]++
							line = line[:len(line)-drop]
//line /usr/local/go/src/bufio/bufio.go:430
		// _ = "end of CoverTab[25357]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:431
		_go_fuzz_dep_.CoverTab[25361]++
//line /usr/local/go/src/bufio/bufio.go:431
		// _ = "end of CoverTab[25361]"
//line /usr/local/go/src/bufio/bufio.go:431
	}
//line /usr/local/go/src/bufio/bufio.go:431
	// _ = "end of CoverTab[25340]"
//line /usr/local/go/src/bufio/bufio.go:431
	_go_fuzz_dep_.CoverTab[25341]++
						return
//line /usr/local/go/src/bufio/bufio.go:432
	// _ = "end of CoverTab[25341]"
}

// collectFragments reads until the first occurrence of delim in the input. It
//line /usr/local/go/src/bufio/bufio.go:435
// returns (slice of full buffers, remaining bytes before delim, total number
//line /usr/local/go/src/bufio/bufio.go:435
// of bytes in the combined first two elements, error).
//line /usr/local/go/src/bufio/bufio.go:435
// The complete result is equal to
//line /usr/local/go/src/bufio/bufio.go:435
// `bytes.Join(append(fullBuffers, finalFragment), nil)`, which has a
//line /usr/local/go/src/bufio/bufio.go:435
// length of `totalLen`. The result is structured in this way to allow callers
//line /usr/local/go/src/bufio/bufio.go:435
// to minimize allocations and copies.
//line /usr/local/go/src/bufio/bufio.go:442
func (b *Reader) collectFragments(delim byte) (fullBuffers [][]byte, finalFragment []byte, totalLen int, err error) {
//line /usr/local/go/src/bufio/bufio.go:442
	_go_fuzz_dep_.CoverTab[25362]++
						var frag []byte

						for {
//line /usr/local/go/src/bufio/bufio.go:445
		_go_fuzz_dep_.CoverTab[25364]++
							var e error
							frag, e = b.ReadSlice(delim)
							if e == nil {
//line /usr/local/go/src/bufio/bufio.go:448
			_go_fuzz_dep_.CoverTab[25367]++
								break
//line /usr/local/go/src/bufio/bufio.go:449
			// _ = "end of CoverTab[25367]"
		} else {
//line /usr/local/go/src/bufio/bufio.go:450
			_go_fuzz_dep_.CoverTab[25368]++
//line /usr/local/go/src/bufio/bufio.go:450
			// _ = "end of CoverTab[25368]"
//line /usr/local/go/src/bufio/bufio.go:450
		}
//line /usr/local/go/src/bufio/bufio.go:450
		// _ = "end of CoverTab[25364]"
//line /usr/local/go/src/bufio/bufio.go:450
		_go_fuzz_dep_.CoverTab[25365]++
							if e != ErrBufferFull {
//line /usr/local/go/src/bufio/bufio.go:451
			_go_fuzz_dep_.CoverTab[25369]++
								err = e
								break
//line /usr/local/go/src/bufio/bufio.go:453
			// _ = "end of CoverTab[25369]"
		} else {
//line /usr/local/go/src/bufio/bufio.go:454
			_go_fuzz_dep_.CoverTab[25370]++
//line /usr/local/go/src/bufio/bufio.go:454
			// _ = "end of CoverTab[25370]"
//line /usr/local/go/src/bufio/bufio.go:454
		}
//line /usr/local/go/src/bufio/bufio.go:454
		// _ = "end of CoverTab[25365]"
//line /usr/local/go/src/bufio/bufio.go:454
		_go_fuzz_dep_.CoverTab[25366]++

//line /usr/local/go/src/bufio/bufio.go:457
		buf := bytes.Clone(frag)
							fullBuffers = append(fullBuffers, buf)
							totalLen += len(buf)
//line /usr/local/go/src/bufio/bufio.go:459
		// _ = "end of CoverTab[25366]"
	}
//line /usr/local/go/src/bufio/bufio.go:460
	// _ = "end of CoverTab[25362]"
//line /usr/local/go/src/bufio/bufio.go:460
	_go_fuzz_dep_.CoverTab[25363]++

						totalLen += len(frag)
						return fullBuffers, frag, totalLen, err
//line /usr/local/go/src/bufio/bufio.go:463
	// _ = "end of CoverTab[25363]"
}

// ReadBytes reads until the first occurrence of delim in the input,
//line /usr/local/go/src/bufio/bufio.go:466
// returning a slice containing the data up to and including the delimiter.
//line /usr/local/go/src/bufio/bufio.go:466
// If ReadBytes encounters an error before finding a delimiter,
//line /usr/local/go/src/bufio/bufio.go:466
// it returns the data read before the error and the error itself (often io.EOF).
//line /usr/local/go/src/bufio/bufio.go:466
// ReadBytes returns err != nil if and only if the returned data does not end in
//line /usr/local/go/src/bufio/bufio.go:466
// delim.
//line /usr/local/go/src/bufio/bufio.go:466
// For simple uses, a Scanner may be more convenient.
//line /usr/local/go/src/bufio/bufio.go:473
func (b *Reader) ReadBytes(delim byte) ([]byte, error) {
//line /usr/local/go/src/bufio/bufio.go:473
	_go_fuzz_dep_.CoverTab[25371]++
						full, frag, n, err := b.collectFragments(delim)

						buf := make([]byte, n)
						n = 0

						for i := range full {
//line /usr/local/go/src/bufio/bufio.go:479
		_go_fuzz_dep_.CoverTab[25373]++
							n += copy(buf[n:], full[i])
//line /usr/local/go/src/bufio/bufio.go:480
		// _ = "end of CoverTab[25373]"
	}
//line /usr/local/go/src/bufio/bufio.go:481
	// _ = "end of CoverTab[25371]"
//line /usr/local/go/src/bufio/bufio.go:481
	_go_fuzz_dep_.CoverTab[25372]++
						copy(buf[n:], frag)
						return buf, err
//line /usr/local/go/src/bufio/bufio.go:483
	// _ = "end of CoverTab[25372]"
}

// ReadString reads until the first occurrence of delim in the input,
//line /usr/local/go/src/bufio/bufio.go:486
// returning a string containing the data up to and including the delimiter.
//line /usr/local/go/src/bufio/bufio.go:486
// If ReadString encounters an error before finding a delimiter,
//line /usr/local/go/src/bufio/bufio.go:486
// it returns the data read before the error and the error itself (often io.EOF).
//line /usr/local/go/src/bufio/bufio.go:486
// ReadString returns err != nil if and only if the returned data does not end in
//line /usr/local/go/src/bufio/bufio.go:486
// delim.
//line /usr/local/go/src/bufio/bufio.go:486
// For simple uses, a Scanner may be more convenient.
//line /usr/local/go/src/bufio/bufio.go:493
func (b *Reader) ReadString(delim byte) (string, error) {
//line /usr/local/go/src/bufio/bufio.go:493
	_go_fuzz_dep_.CoverTab[25374]++
						full, frag, n, err := b.collectFragments(delim)
	// Allocate new buffer to hold the full pieces and the fragment.
	var buf strings.Builder
	buf.Grow(n)

	for _, fb := range full {
//line /usr/local/go/src/bufio/bufio.go:499
		_go_fuzz_dep_.CoverTab[25376]++
							buf.Write(fb)
//line /usr/local/go/src/bufio/bufio.go:500
		// _ = "end of CoverTab[25376]"
	}
//line /usr/local/go/src/bufio/bufio.go:501
	// _ = "end of CoverTab[25374]"
//line /usr/local/go/src/bufio/bufio.go:501
	_go_fuzz_dep_.CoverTab[25375]++
						buf.Write(frag)
						return buf.String(), err
//line /usr/local/go/src/bufio/bufio.go:503
	// _ = "end of CoverTab[25375]"
}

// WriteTo implements io.WriterTo.
//line /usr/local/go/src/bufio/bufio.go:506
// This may make multiple calls to the Read method of the underlying Reader.
//line /usr/local/go/src/bufio/bufio.go:506
// If the underlying reader supports the WriteTo method,
//line /usr/local/go/src/bufio/bufio.go:506
// this calls the underlying WriteTo without buffering.
//line /usr/local/go/src/bufio/bufio.go:510
func (b *Reader) WriteTo(w io.Writer) (n int64, err error) {
//line /usr/local/go/src/bufio/bufio.go:510
	_go_fuzz_dep_.CoverTab[25377]++
						b.lastByte = -1
						b.lastRuneSize = -1

						n, err = b.writeBuf(w)
						if err != nil {
//line /usr/local/go/src/bufio/bufio.go:515
		_go_fuzz_dep_.CoverTab[25384]++
							return
//line /usr/local/go/src/bufio/bufio.go:516
		// _ = "end of CoverTab[25384]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:517
		_go_fuzz_dep_.CoverTab[25385]++
//line /usr/local/go/src/bufio/bufio.go:517
		// _ = "end of CoverTab[25385]"
//line /usr/local/go/src/bufio/bufio.go:517
	}
//line /usr/local/go/src/bufio/bufio.go:517
	// _ = "end of CoverTab[25377]"
//line /usr/local/go/src/bufio/bufio.go:517
	_go_fuzz_dep_.CoverTab[25378]++

						if r, ok := b.rd.(io.WriterTo); ok {
//line /usr/local/go/src/bufio/bufio.go:519
		_go_fuzz_dep_.CoverTab[25386]++
							m, err := r.WriteTo(w)
							n += m
							return n, err
//line /usr/local/go/src/bufio/bufio.go:522
		// _ = "end of CoverTab[25386]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:523
		_go_fuzz_dep_.CoverTab[25387]++
//line /usr/local/go/src/bufio/bufio.go:523
		// _ = "end of CoverTab[25387]"
//line /usr/local/go/src/bufio/bufio.go:523
	}
//line /usr/local/go/src/bufio/bufio.go:523
	// _ = "end of CoverTab[25378]"
//line /usr/local/go/src/bufio/bufio.go:523
	_go_fuzz_dep_.CoverTab[25379]++

						if w, ok := w.(io.ReaderFrom); ok {
//line /usr/local/go/src/bufio/bufio.go:525
		_go_fuzz_dep_.CoverTab[25388]++
							m, err := w.ReadFrom(b.rd)
							n += m
							return n, err
//line /usr/local/go/src/bufio/bufio.go:528
		// _ = "end of CoverTab[25388]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:529
		_go_fuzz_dep_.CoverTab[25389]++
//line /usr/local/go/src/bufio/bufio.go:529
		// _ = "end of CoverTab[25389]"
//line /usr/local/go/src/bufio/bufio.go:529
	}
//line /usr/local/go/src/bufio/bufio.go:529
	// _ = "end of CoverTab[25379]"
//line /usr/local/go/src/bufio/bufio.go:529
	_go_fuzz_dep_.CoverTab[25380]++

						if b.w-b.r < len(b.buf) {
//line /usr/local/go/src/bufio/bufio.go:531
		_go_fuzz_dep_.CoverTab[25390]++
							b.fill()
//line /usr/local/go/src/bufio/bufio.go:532
		// _ = "end of CoverTab[25390]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:533
		_go_fuzz_dep_.CoverTab[25391]++
//line /usr/local/go/src/bufio/bufio.go:533
		// _ = "end of CoverTab[25391]"
//line /usr/local/go/src/bufio/bufio.go:533
	}
//line /usr/local/go/src/bufio/bufio.go:533
	// _ = "end of CoverTab[25380]"
//line /usr/local/go/src/bufio/bufio.go:533
	_go_fuzz_dep_.CoverTab[25381]++

						for b.r < b.w {
//line /usr/local/go/src/bufio/bufio.go:535
		_go_fuzz_dep_.CoverTab[25392]++

							m, err := b.writeBuf(w)
							n += m
							if err != nil {
//line /usr/local/go/src/bufio/bufio.go:539
			_go_fuzz_dep_.CoverTab[25394]++
								return n, err
//line /usr/local/go/src/bufio/bufio.go:540
			// _ = "end of CoverTab[25394]"
		} else {
//line /usr/local/go/src/bufio/bufio.go:541
			_go_fuzz_dep_.CoverTab[25395]++
//line /usr/local/go/src/bufio/bufio.go:541
			// _ = "end of CoverTab[25395]"
//line /usr/local/go/src/bufio/bufio.go:541
		}
//line /usr/local/go/src/bufio/bufio.go:541
		// _ = "end of CoverTab[25392]"
//line /usr/local/go/src/bufio/bufio.go:541
		_go_fuzz_dep_.CoverTab[25393]++
							b.fill()
//line /usr/local/go/src/bufio/bufio.go:542
		// _ = "end of CoverTab[25393]"
	}
//line /usr/local/go/src/bufio/bufio.go:543
	// _ = "end of CoverTab[25381]"
//line /usr/local/go/src/bufio/bufio.go:543
	_go_fuzz_dep_.CoverTab[25382]++

						if b.err == io.EOF {
//line /usr/local/go/src/bufio/bufio.go:545
		_go_fuzz_dep_.CoverTab[25396]++
							b.err = nil
//line /usr/local/go/src/bufio/bufio.go:546
		// _ = "end of CoverTab[25396]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:547
		_go_fuzz_dep_.CoverTab[25397]++
//line /usr/local/go/src/bufio/bufio.go:547
		// _ = "end of CoverTab[25397]"
//line /usr/local/go/src/bufio/bufio.go:547
	}
//line /usr/local/go/src/bufio/bufio.go:547
	// _ = "end of CoverTab[25382]"
//line /usr/local/go/src/bufio/bufio.go:547
	_go_fuzz_dep_.CoverTab[25383]++

						return n, b.readErr()
//line /usr/local/go/src/bufio/bufio.go:549
	// _ = "end of CoverTab[25383]"
}

var errNegativeWrite = errors.New("bufio: writer returned negative count from Write")

// writeBuf writes the Reader's buffer to the writer.
func (b *Reader) writeBuf(w io.Writer) (int64, error) {
//line /usr/local/go/src/bufio/bufio.go:555
	_go_fuzz_dep_.CoverTab[25398]++
						n, err := w.Write(b.buf[b.r:b.w])
						if n < 0 {
//line /usr/local/go/src/bufio/bufio.go:557
		_go_fuzz_dep_.CoverTab[25400]++
							panic(errNegativeWrite)
//line /usr/local/go/src/bufio/bufio.go:558
		// _ = "end of CoverTab[25400]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:559
		_go_fuzz_dep_.CoverTab[25401]++
//line /usr/local/go/src/bufio/bufio.go:559
		// _ = "end of CoverTab[25401]"
//line /usr/local/go/src/bufio/bufio.go:559
	}
//line /usr/local/go/src/bufio/bufio.go:559
	// _ = "end of CoverTab[25398]"
//line /usr/local/go/src/bufio/bufio.go:559
	_go_fuzz_dep_.CoverTab[25399]++
						b.r += n
						return int64(n), err
//line /usr/local/go/src/bufio/bufio.go:561
	// _ = "end of CoverTab[25399]"
}

//line /usr/local/go/src/bufio/bufio.go:566
// Writer implements buffering for an io.Writer object.
//line /usr/local/go/src/bufio/bufio.go:566
// If an error occurs writing to a Writer, no more data will be
//line /usr/local/go/src/bufio/bufio.go:566
// accepted and all subsequent writes, and Flush, will return the error.
//line /usr/local/go/src/bufio/bufio.go:566
// After all data has been written, the client should call the
//line /usr/local/go/src/bufio/bufio.go:566
// Flush method to guarantee all data has been forwarded to
//line /usr/local/go/src/bufio/bufio.go:566
// the underlying io.Writer.
//line /usr/local/go/src/bufio/bufio.go:572
type Writer struct {
	err	error
	buf	[]byte
	n	int
	wr	io.Writer
}

// NewWriterSize returns a new Writer whose buffer has at least the specified
//line /usr/local/go/src/bufio/bufio.go:579
// size. If the argument io.Writer is already a Writer with large enough
//line /usr/local/go/src/bufio/bufio.go:579
// size, it returns the underlying Writer.
//line /usr/local/go/src/bufio/bufio.go:582
func NewWriterSize(w io.Writer, size int) *Writer {
//line /usr/local/go/src/bufio/bufio.go:582
	_go_fuzz_dep_.CoverTab[25402]++

						b, ok := w.(*Writer)
						if ok && func() bool {
//line /usr/local/go/src/bufio/bufio.go:585
		_go_fuzz_dep_.CoverTab[25405]++
//line /usr/local/go/src/bufio/bufio.go:585
		return len(b.buf) >= size
//line /usr/local/go/src/bufio/bufio.go:585
		// _ = "end of CoverTab[25405]"
//line /usr/local/go/src/bufio/bufio.go:585
	}() {
//line /usr/local/go/src/bufio/bufio.go:585
		_go_fuzz_dep_.CoverTab[25406]++
							return b
//line /usr/local/go/src/bufio/bufio.go:586
		// _ = "end of CoverTab[25406]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:587
		_go_fuzz_dep_.CoverTab[25407]++
//line /usr/local/go/src/bufio/bufio.go:587
		// _ = "end of CoverTab[25407]"
//line /usr/local/go/src/bufio/bufio.go:587
	}
//line /usr/local/go/src/bufio/bufio.go:587
	// _ = "end of CoverTab[25402]"
//line /usr/local/go/src/bufio/bufio.go:587
	_go_fuzz_dep_.CoverTab[25403]++
						if size <= 0 {
//line /usr/local/go/src/bufio/bufio.go:588
		_go_fuzz_dep_.CoverTab[25408]++
							size = defaultBufSize
//line /usr/local/go/src/bufio/bufio.go:589
		// _ = "end of CoverTab[25408]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:590
		_go_fuzz_dep_.CoverTab[25409]++
//line /usr/local/go/src/bufio/bufio.go:590
		// _ = "end of CoverTab[25409]"
//line /usr/local/go/src/bufio/bufio.go:590
	}
//line /usr/local/go/src/bufio/bufio.go:590
	// _ = "end of CoverTab[25403]"
//line /usr/local/go/src/bufio/bufio.go:590
	_go_fuzz_dep_.CoverTab[25404]++
						return &Writer{
		buf:	make([]byte, size),
		wr:	w,
	}
//line /usr/local/go/src/bufio/bufio.go:594
	// _ = "end of CoverTab[25404]"
}

// NewWriter returns a new Writer whose buffer has the default size.
//line /usr/local/go/src/bufio/bufio.go:597
// If the argument io.Writer is already a Writer with large enough buffer size,
//line /usr/local/go/src/bufio/bufio.go:597
// it returns the underlying Writer.
//line /usr/local/go/src/bufio/bufio.go:600
func NewWriter(w io.Writer) *Writer {
//line /usr/local/go/src/bufio/bufio.go:600
	_go_fuzz_dep_.CoverTab[25410]++
						return NewWriterSize(w, defaultBufSize)
//line /usr/local/go/src/bufio/bufio.go:601
	// _ = "end of CoverTab[25410]"
}

// Size returns the size of the underlying buffer in bytes.
func (b *Writer) Size() int {
//line /usr/local/go/src/bufio/bufio.go:605
	_go_fuzz_dep_.CoverTab[25411]++
//line /usr/local/go/src/bufio/bufio.go:605
	return len(b.buf)
//line /usr/local/go/src/bufio/bufio.go:605
	// _ = "end of CoverTab[25411]"
//line /usr/local/go/src/bufio/bufio.go:605
}

// Reset discards any unflushed buffered data, clears any error, and
//line /usr/local/go/src/bufio/bufio.go:607
// resets b to write its output to w.
//line /usr/local/go/src/bufio/bufio.go:607
// Calling Reset on the zero value of Writer initializes the internal buffer
//line /usr/local/go/src/bufio/bufio.go:607
// to the default size.
//line /usr/local/go/src/bufio/bufio.go:611
func (b *Writer) Reset(w io.Writer) {
//line /usr/local/go/src/bufio/bufio.go:611
	_go_fuzz_dep_.CoverTab[25412]++
						if b.buf == nil {
//line /usr/local/go/src/bufio/bufio.go:612
		_go_fuzz_dep_.CoverTab[25414]++
							b.buf = make([]byte, defaultBufSize)
//line /usr/local/go/src/bufio/bufio.go:613
		// _ = "end of CoverTab[25414]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:614
		_go_fuzz_dep_.CoverTab[25415]++
//line /usr/local/go/src/bufio/bufio.go:614
		// _ = "end of CoverTab[25415]"
//line /usr/local/go/src/bufio/bufio.go:614
	}
//line /usr/local/go/src/bufio/bufio.go:614
	// _ = "end of CoverTab[25412]"
//line /usr/local/go/src/bufio/bufio.go:614
	_go_fuzz_dep_.CoverTab[25413]++
						b.err = nil
						b.n = 0
						b.wr = w
//line /usr/local/go/src/bufio/bufio.go:617
	// _ = "end of CoverTab[25413]"
}

// Flush writes any buffered data to the underlying io.Writer.
func (b *Writer) Flush() error {
//line /usr/local/go/src/bufio/bufio.go:621
	_go_fuzz_dep_.CoverTab[25416]++
						if b.err != nil {
//line /usr/local/go/src/bufio/bufio.go:622
		_go_fuzz_dep_.CoverTab[25421]++
							return b.err
//line /usr/local/go/src/bufio/bufio.go:623
		// _ = "end of CoverTab[25421]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:624
		_go_fuzz_dep_.CoverTab[25422]++
//line /usr/local/go/src/bufio/bufio.go:624
		// _ = "end of CoverTab[25422]"
//line /usr/local/go/src/bufio/bufio.go:624
	}
//line /usr/local/go/src/bufio/bufio.go:624
	// _ = "end of CoverTab[25416]"
//line /usr/local/go/src/bufio/bufio.go:624
	_go_fuzz_dep_.CoverTab[25417]++
						if b.n == 0 {
//line /usr/local/go/src/bufio/bufio.go:625
		_go_fuzz_dep_.CoverTab[25423]++
							return nil
//line /usr/local/go/src/bufio/bufio.go:626
		// _ = "end of CoverTab[25423]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:627
		_go_fuzz_dep_.CoverTab[25424]++
//line /usr/local/go/src/bufio/bufio.go:627
		// _ = "end of CoverTab[25424]"
//line /usr/local/go/src/bufio/bufio.go:627
	}
//line /usr/local/go/src/bufio/bufio.go:627
	// _ = "end of CoverTab[25417]"
//line /usr/local/go/src/bufio/bufio.go:627
	_go_fuzz_dep_.CoverTab[25418]++
						n, err := b.wr.Write(b.buf[0:b.n])
						if n < b.n && func() bool {
//line /usr/local/go/src/bufio/bufio.go:629
		_go_fuzz_dep_.CoverTab[25425]++
//line /usr/local/go/src/bufio/bufio.go:629
		return err == nil
//line /usr/local/go/src/bufio/bufio.go:629
		// _ = "end of CoverTab[25425]"
//line /usr/local/go/src/bufio/bufio.go:629
	}() {
//line /usr/local/go/src/bufio/bufio.go:629
		_go_fuzz_dep_.CoverTab[25426]++
							err = io.ErrShortWrite
//line /usr/local/go/src/bufio/bufio.go:630
		// _ = "end of CoverTab[25426]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:631
		_go_fuzz_dep_.CoverTab[25427]++
//line /usr/local/go/src/bufio/bufio.go:631
		// _ = "end of CoverTab[25427]"
//line /usr/local/go/src/bufio/bufio.go:631
	}
//line /usr/local/go/src/bufio/bufio.go:631
	// _ = "end of CoverTab[25418]"
//line /usr/local/go/src/bufio/bufio.go:631
	_go_fuzz_dep_.CoverTab[25419]++
						if err != nil {
//line /usr/local/go/src/bufio/bufio.go:632
		_go_fuzz_dep_.CoverTab[25428]++
							if n > 0 && func() bool {
//line /usr/local/go/src/bufio/bufio.go:633
			_go_fuzz_dep_.CoverTab[25430]++
//line /usr/local/go/src/bufio/bufio.go:633
			return n < b.n
//line /usr/local/go/src/bufio/bufio.go:633
			// _ = "end of CoverTab[25430]"
//line /usr/local/go/src/bufio/bufio.go:633
		}() {
//line /usr/local/go/src/bufio/bufio.go:633
			_go_fuzz_dep_.CoverTab[25431]++
								copy(b.buf[0:b.n-n], b.buf[n:b.n])
//line /usr/local/go/src/bufio/bufio.go:634
			// _ = "end of CoverTab[25431]"
		} else {
//line /usr/local/go/src/bufio/bufio.go:635
			_go_fuzz_dep_.CoverTab[25432]++
//line /usr/local/go/src/bufio/bufio.go:635
			// _ = "end of CoverTab[25432]"
//line /usr/local/go/src/bufio/bufio.go:635
		}
//line /usr/local/go/src/bufio/bufio.go:635
		// _ = "end of CoverTab[25428]"
//line /usr/local/go/src/bufio/bufio.go:635
		_go_fuzz_dep_.CoverTab[25429]++
							b.n -= n
							b.err = err
							return err
//line /usr/local/go/src/bufio/bufio.go:638
		// _ = "end of CoverTab[25429]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:639
		_go_fuzz_dep_.CoverTab[25433]++
//line /usr/local/go/src/bufio/bufio.go:639
		// _ = "end of CoverTab[25433]"
//line /usr/local/go/src/bufio/bufio.go:639
	}
//line /usr/local/go/src/bufio/bufio.go:639
	// _ = "end of CoverTab[25419]"
//line /usr/local/go/src/bufio/bufio.go:639
	_go_fuzz_dep_.CoverTab[25420]++
						b.n = 0
						return nil
//line /usr/local/go/src/bufio/bufio.go:641
	// _ = "end of CoverTab[25420]"
}

// Available returns how many bytes are unused in the buffer.
func (b *Writer) Available() int {
//line /usr/local/go/src/bufio/bufio.go:645
	_go_fuzz_dep_.CoverTab[25434]++
//line /usr/local/go/src/bufio/bufio.go:645
	return len(b.buf) - b.n
//line /usr/local/go/src/bufio/bufio.go:645
	// _ = "end of CoverTab[25434]"
//line /usr/local/go/src/bufio/bufio.go:645
}

// AvailableBuffer returns an empty buffer with b.Available() capacity.
//line /usr/local/go/src/bufio/bufio.go:647
// This buffer is intended to be appended to and
//line /usr/local/go/src/bufio/bufio.go:647
// passed to an immediately succeeding Write call.
//line /usr/local/go/src/bufio/bufio.go:647
// The buffer is only valid until the next write operation on b.
//line /usr/local/go/src/bufio/bufio.go:651
func (b *Writer) AvailableBuffer() []byte {
//line /usr/local/go/src/bufio/bufio.go:651
	_go_fuzz_dep_.CoverTab[25435]++
						return b.buf[b.n:][:0]
//line /usr/local/go/src/bufio/bufio.go:652
	// _ = "end of CoverTab[25435]"
}

// Buffered returns the number of bytes that have been written into the current buffer.
func (b *Writer) Buffered() int {
//line /usr/local/go/src/bufio/bufio.go:656
	_go_fuzz_dep_.CoverTab[25436]++
//line /usr/local/go/src/bufio/bufio.go:656
	return b.n
//line /usr/local/go/src/bufio/bufio.go:656
	// _ = "end of CoverTab[25436]"
//line /usr/local/go/src/bufio/bufio.go:656
}

// Write writes the contents of p into the buffer.
//line /usr/local/go/src/bufio/bufio.go:658
// It returns the number of bytes written.
//line /usr/local/go/src/bufio/bufio.go:658
// If nn < len(p), it also returns an error explaining
//line /usr/local/go/src/bufio/bufio.go:658
// why the write is short.
//line /usr/local/go/src/bufio/bufio.go:662
func (b *Writer) Write(p []byte) (nn int, err error) {
//line /usr/local/go/src/bufio/bufio.go:662
	_go_fuzz_dep_.CoverTab[25437]++
						for len(p) > b.Available() && func() bool {
//line /usr/local/go/src/bufio/bufio.go:663
		_go_fuzz_dep_.CoverTab[25440]++
//line /usr/local/go/src/bufio/bufio.go:663
		return b.err == nil
//line /usr/local/go/src/bufio/bufio.go:663
		// _ = "end of CoverTab[25440]"
//line /usr/local/go/src/bufio/bufio.go:663
	}() {
//line /usr/local/go/src/bufio/bufio.go:663
		_go_fuzz_dep_.CoverTab[25441]++
							var n int
							if b.Buffered() == 0 {
//line /usr/local/go/src/bufio/bufio.go:665
			_go_fuzz_dep_.CoverTab[25443]++

//line /usr/local/go/src/bufio/bufio.go:668
			n, b.err = b.wr.Write(p)
//line /usr/local/go/src/bufio/bufio.go:668
			// _ = "end of CoverTab[25443]"
		} else {
//line /usr/local/go/src/bufio/bufio.go:669
			_go_fuzz_dep_.CoverTab[25444]++
								n = copy(b.buf[b.n:], p)
								b.n += n
								b.Flush()
//line /usr/local/go/src/bufio/bufio.go:672
			// _ = "end of CoverTab[25444]"
		}
//line /usr/local/go/src/bufio/bufio.go:673
		// _ = "end of CoverTab[25441]"
//line /usr/local/go/src/bufio/bufio.go:673
		_go_fuzz_dep_.CoverTab[25442]++
							nn += n
							p = p[n:]
//line /usr/local/go/src/bufio/bufio.go:675
		// _ = "end of CoverTab[25442]"
	}
//line /usr/local/go/src/bufio/bufio.go:676
	// _ = "end of CoverTab[25437]"
//line /usr/local/go/src/bufio/bufio.go:676
	_go_fuzz_dep_.CoverTab[25438]++
						if b.err != nil {
//line /usr/local/go/src/bufio/bufio.go:677
		_go_fuzz_dep_.CoverTab[25445]++
							return nn, b.err
//line /usr/local/go/src/bufio/bufio.go:678
		// _ = "end of CoverTab[25445]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:679
		_go_fuzz_dep_.CoverTab[25446]++
//line /usr/local/go/src/bufio/bufio.go:679
		// _ = "end of CoverTab[25446]"
//line /usr/local/go/src/bufio/bufio.go:679
	}
//line /usr/local/go/src/bufio/bufio.go:679
	// _ = "end of CoverTab[25438]"
//line /usr/local/go/src/bufio/bufio.go:679
	_go_fuzz_dep_.CoverTab[25439]++
						n := copy(b.buf[b.n:], p)
						b.n += n
						nn += n
						return nn, nil
//line /usr/local/go/src/bufio/bufio.go:683
	// _ = "end of CoverTab[25439]"
}

// WriteByte writes a single byte.
func (b *Writer) WriteByte(c byte) error {
//line /usr/local/go/src/bufio/bufio.go:687
	_go_fuzz_dep_.CoverTab[25447]++
						if b.err != nil {
//line /usr/local/go/src/bufio/bufio.go:688
		_go_fuzz_dep_.CoverTab[25450]++
							return b.err
//line /usr/local/go/src/bufio/bufio.go:689
		// _ = "end of CoverTab[25450]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:690
		_go_fuzz_dep_.CoverTab[25451]++
//line /usr/local/go/src/bufio/bufio.go:690
		// _ = "end of CoverTab[25451]"
//line /usr/local/go/src/bufio/bufio.go:690
	}
//line /usr/local/go/src/bufio/bufio.go:690
	// _ = "end of CoverTab[25447]"
//line /usr/local/go/src/bufio/bufio.go:690
	_go_fuzz_dep_.CoverTab[25448]++
						if b.Available() <= 0 && func() bool {
//line /usr/local/go/src/bufio/bufio.go:691
		_go_fuzz_dep_.CoverTab[25452]++
//line /usr/local/go/src/bufio/bufio.go:691
		return b.Flush() != nil
//line /usr/local/go/src/bufio/bufio.go:691
		// _ = "end of CoverTab[25452]"
//line /usr/local/go/src/bufio/bufio.go:691
	}() {
//line /usr/local/go/src/bufio/bufio.go:691
		_go_fuzz_dep_.CoverTab[25453]++
							return b.err
//line /usr/local/go/src/bufio/bufio.go:692
		// _ = "end of CoverTab[25453]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:693
		_go_fuzz_dep_.CoverTab[25454]++
//line /usr/local/go/src/bufio/bufio.go:693
		// _ = "end of CoverTab[25454]"
//line /usr/local/go/src/bufio/bufio.go:693
	}
//line /usr/local/go/src/bufio/bufio.go:693
	// _ = "end of CoverTab[25448]"
//line /usr/local/go/src/bufio/bufio.go:693
	_go_fuzz_dep_.CoverTab[25449]++
						b.buf[b.n] = c
						b.n++
						return nil
//line /usr/local/go/src/bufio/bufio.go:696
	// _ = "end of CoverTab[25449]"
}

// WriteRune writes a single Unicode code point, returning
//line /usr/local/go/src/bufio/bufio.go:699
// the number of bytes written and any error.
//line /usr/local/go/src/bufio/bufio.go:701
func (b *Writer) WriteRune(r rune) (size int, err error) {
//line /usr/local/go/src/bufio/bufio.go:701
	_go_fuzz_dep_.CoverTab[25455]++

						if uint32(r) < utf8.RuneSelf {
//line /usr/local/go/src/bufio/bufio.go:703
		_go_fuzz_dep_.CoverTab[25459]++
							err = b.WriteByte(byte(r))
							if err != nil {
//line /usr/local/go/src/bufio/bufio.go:705
			_go_fuzz_dep_.CoverTab[25461]++
								return 0, err
//line /usr/local/go/src/bufio/bufio.go:706
			// _ = "end of CoverTab[25461]"
		} else {
//line /usr/local/go/src/bufio/bufio.go:707
			_go_fuzz_dep_.CoverTab[25462]++
//line /usr/local/go/src/bufio/bufio.go:707
			// _ = "end of CoverTab[25462]"
//line /usr/local/go/src/bufio/bufio.go:707
		}
//line /usr/local/go/src/bufio/bufio.go:707
		// _ = "end of CoverTab[25459]"
//line /usr/local/go/src/bufio/bufio.go:707
		_go_fuzz_dep_.CoverTab[25460]++
							return 1, nil
//line /usr/local/go/src/bufio/bufio.go:708
		// _ = "end of CoverTab[25460]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:709
		_go_fuzz_dep_.CoverTab[25463]++
//line /usr/local/go/src/bufio/bufio.go:709
		// _ = "end of CoverTab[25463]"
//line /usr/local/go/src/bufio/bufio.go:709
	}
//line /usr/local/go/src/bufio/bufio.go:709
	// _ = "end of CoverTab[25455]"
//line /usr/local/go/src/bufio/bufio.go:709
	_go_fuzz_dep_.CoverTab[25456]++
						if b.err != nil {
//line /usr/local/go/src/bufio/bufio.go:710
		_go_fuzz_dep_.CoverTab[25464]++
							return 0, b.err
//line /usr/local/go/src/bufio/bufio.go:711
		// _ = "end of CoverTab[25464]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:712
		_go_fuzz_dep_.CoverTab[25465]++
//line /usr/local/go/src/bufio/bufio.go:712
		// _ = "end of CoverTab[25465]"
//line /usr/local/go/src/bufio/bufio.go:712
	}
//line /usr/local/go/src/bufio/bufio.go:712
	// _ = "end of CoverTab[25456]"
//line /usr/local/go/src/bufio/bufio.go:712
	_go_fuzz_dep_.CoverTab[25457]++
						n := b.Available()
						if n < utf8.UTFMax {
//line /usr/local/go/src/bufio/bufio.go:714
		_go_fuzz_dep_.CoverTab[25466]++
							if b.Flush(); b.err != nil {
//line /usr/local/go/src/bufio/bufio.go:715
			_go_fuzz_dep_.CoverTab[25468]++
								return 0, b.err
//line /usr/local/go/src/bufio/bufio.go:716
			// _ = "end of CoverTab[25468]"
		} else {
//line /usr/local/go/src/bufio/bufio.go:717
			_go_fuzz_dep_.CoverTab[25469]++
//line /usr/local/go/src/bufio/bufio.go:717
			// _ = "end of CoverTab[25469]"
//line /usr/local/go/src/bufio/bufio.go:717
		}
//line /usr/local/go/src/bufio/bufio.go:717
		// _ = "end of CoverTab[25466]"
//line /usr/local/go/src/bufio/bufio.go:717
		_go_fuzz_dep_.CoverTab[25467]++
							n = b.Available()
							if n < utf8.UTFMax {
//line /usr/local/go/src/bufio/bufio.go:719
			_go_fuzz_dep_.CoverTab[25470]++

								return b.WriteString(string(r))
//line /usr/local/go/src/bufio/bufio.go:721
			// _ = "end of CoverTab[25470]"
		} else {
//line /usr/local/go/src/bufio/bufio.go:722
			_go_fuzz_dep_.CoverTab[25471]++
//line /usr/local/go/src/bufio/bufio.go:722
			// _ = "end of CoverTab[25471]"
//line /usr/local/go/src/bufio/bufio.go:722
		}
//line /usr/local/go/src/bufio/bufio.go:722
		// _ = "end of CoverTab[25467]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:723
		_go_fuzz_dep_.CoverTab[25472]++
//line /usr/local/go/src/bufio/bufio.go:723
		// _ = "end of CoverTab[25472]"
//line /usr/local/go/src/bufio/bufio.go:723
	}
//line /usr/local/go/src/bufio/bufio.go:723
	// _ = "end of CoverTab[25457]"
//line /usr/local/go/src/bufio/bufio.go:723
	_go_fuzz_dep_.CoverTab[25458]++
						size = utf8.EncodeRune(b.buf[b.n:], r)
						b.n += size
						return size, nil
//line /usr/local/go/src/bufio/bufio.go:726
	// _ = "end of CoverTab[25458]"
}

// WriteString writes a string.
//line /usr/local/go/src/bufio/bufio.go:729
// It returns the number of bytes written.
//line /usr/local/go/src/bufio/bufio.go:729
// If the count is less than len(s), it also returns an error explaining
//line /usr/local/go/src/bufio/bufio.go:729
// why the write is short.
//line /usr/local/go/src/bufio/bufio.go:733
func (b *Writer) WriteString(s string) (int, error) {
//line /usr/local/go/src/bufio/bufio.go:733
	_go_fuzz_dep_.CoverTab[25473]++
						var sw io.StringWriter
						tryStringWriter := true

						nn := 0
						for len(s) > b.Available() && func() bool {
//line /usr/local/go/src/bufio/bufio.go:738
		_go_fuzz_dep_.CoverTab[25476]++
//line /usr/local/go/src/bufio/bufio.go:738
		return b.err == nil
//line /usr/local/go/src/bufio/bufio.go:738
		// _ = "end of CoverTab[25476]"
//line /usr/local/go/src/bufio/bufio.go:738
	}() {
//line /usr/local/go/src/bufio/bufio.go:738
		_go_fuzz_dep_.CoverTab[25477]++
							var n int
							if b.Buffered() == 0 && func() bool {
//line /usr/local/go/src/bufio/bufio.go:740
			_go_fuzz_dep_.CoverTab[25480]++
//line /usr/local/go/src/bufio/bufio.go:740
			return sw == nil
//line /usr/local/go/src/bufio/bufio.go:740
			// _ = "end of CoverTab[25480]"
//line /usr/local/go/src/bufio/bufio.go:740
		}() && func() bool {
//line /usr/local/go/src/bufio/bufio.go:740
			_go_fuzz_dep_.CoverTab[25481]++
//line /usr/local/go/src/bufio/bufio.go:740
			return tryStringWriter
//line /usr/local/go/src/bufio/bufio.go:740
			// _ = "end of CoverTab[25481]"
//line /usr/local/go/src/bufio/bufio.go:740
		}() {
//line /usr/local/go/src/bufio/bufio.go:740
			_go_fuzz_dep_.CoverTab[25482]++

								sw, tryStringWriter = b.wr.(io.StringWriter)
//line /usr/local/go/src/bufio/bufio.go:742
			// _ = "end of CoverTab[25482]"
		} else {
//line /usr/local/go/src/bufio/bufio.go:743
			_go_fuzz_dep_.CoverTab[25483]++
//line /usr/local/go/src/bufio/bufio.go:743
			// _ = "end of CoverTab[25483]"
//line /usr/local/go/src/bufio/bufio.go:743
		}
//line /usr/local/go/src/bufio/bufio.go:743
		// _ = "end of CoverTab[25477]"
//line /usr/local/go/src/bufio/bufio.go:743
		_go_fuzz_dep_.CoverTab[25478]++
							if b.Buffered() == 0 && func() bool {
//line /usr/local/go/src/bufio/bufio.go:744
			_go_fuzz_dep_.CoverTab[25484]++
//line /usr/local/go/src/bufio/bufio.go:744
			return tryStringWriter
//line /usr/local/go/src/bufio/bufio.go:744
			// _ = "end of CoverTab[25484]"
//line /usr/local/go/src/bufio/bufio.go:744
		}() {
//line /usr/local/go/src/bufio/bufio.go:744
			_go_fuzz_dep_.CoverTab[25485]++

//line /usr/local/go/src/bufio/bufio.go:748
			n, b.err = sw.WriteString(s)
//line /usr/local/go/src/bufio/bufio.go:748
			// _ = "end of CoverTab[25485]"
		} else {
//line /usr/local/go/src/bufio/bufio.go:749
			_go_fuzz_dep_.CoverTab[25486]++
								n = copy(b.buf[b.n:], s)
								b.n += n
								b.Flush()
//line /usr/local/go/src/bufio/bufio.go:752
			// _ = "end of CoverTab[25486]"
		}
//line /usr/local/go/src/bufio/bufio.go:753
		// _ = "end of CoverTab[25478]"
//line /usr/local/go/src/bufio/bufio.go:753
		_go_fuzz_dep_.CoverTab[25479]++
							nn += n
							s = s[n:]
//line /usr/local/go/src/bufio/bufio.go:755
		// _ = "end of CoverTab[25479]"
	}
//line /usr/local/go/src/bufio/bufio.go:756
	// _ = "end of CoverTab[25473]"
//line /usr/local/go/src/bufio/bufio.go:756
	_go_fuzz_dep_.CoverTab[25474]++
						if b.err != nil {
//line /usr/local/go/src/bufio/bufio.go:757
		_go_fuzz_dep_.CoverTab[25487]++
							return nn, b.err
//line /usr/local/go/src/bufio/bufio.go:758
		// _ = "end of CoverTab[25487]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:759
		_go_fuzz_dep_.CoverTab[25488]++
//line /usr/local/go/src/bufio/bufio.go:759
		// _ = "end of CoverTab[25488]"
//line /usr/local/go/src/bufio/bufio.go:759
	}
//line /usr/local/go/src/bufio/bufio.go:759
	// _ = "end of CoverTab[25474]"
//line /usr/local/go/src/bufio/bufio.go:759
	_go_fuzz_dep_.CoverTab[25475]++
						n := copy(b.buf[b.n:], s)
						b.n += n
						nn += n
						return nn, nil
//line /usr/local/go/src/bufio/bufio.go:763
	// _ = "end of CoverTab[25475]"
}

// ReadFrom implements io.ReaderFrom. If the underlying writer
//line /usr/local/go/src/bufio/bufio.go:766
// supports the ReadFrom method, this calls the underlying ReadFrom.
//line /usr/local/go/src/bufio/bufio.go:766
// If there is buffered data and an underlying ReadFrom, this fills
//line /usr/local/go/src/bufio/bufio.go:766
// the buffer and writes it before calling ReadFrom.
//line /usr/local/go/src/bufio/bufio.go:770
func (b *Writer) ReadFrom(r io.Reader) (n int64, err error) {
//line /usr/local/go/src/bufio/bufio.go:770
	_go_fuzz_dep_.CoverTab[25489]++
						if b.err != nil {
//line /usr/local/go/src/bufio/bufio.go:771
		_go_fuzz_dep_.CoverTab[25493]++
							return 0, b.err
//line /usr/local/go/src/bufio/bufio.go:772
		// _ = "end of CoverTab[25493]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:773
		_go_fuzz_dep_.CoverTab[25494]++
//line /usr/local/go/src/bufio/bufio.go:773
		// _ = "end of CoverTab[25494]"
//line /usr/local/go/src/bufio/bufio.go:773
	}
//line /usr/local/go/src/bufio/bufio.go:773
	// _ = "end of CoverTab[25489]"
//line /usr/local/go/src/bufio/bufio.go:773
	_go_fuzz_dep_.CoverTab[25490]++
						readerFrom, readerFromOK := b.wr.(io.ReaderFrom)
						var m int
						for {
//line /usr/local/go/src/bufio/bufio.go:776
		_go_fuzz_dep_.CoverTab[25495]++
							if b.Available() == 0 {
//line /usr/local/go/src/bufio/bufio.go:777
			_go_fuzz_dep_.CoverTab[25500]++
								if err1 := b.Flush(); err1 != nil {
//line /usr/local/go/src/bufio/bufio.go:778
				_go_fuzz_dep_.CoverTab[25501]++
									return n, err1
//line /usr/local/go/src/bufio/bufio.go:779
				// _ = "end of CoverTab[25501]"
			} else {
//line /usr/local/go/src/bufio/bufio.go:780
				_go_fuzz_dep_.CoverTab[25502]++
//line /usr/local/go/src/bufio/bufio.go:780
				// _ = "end of CoverTab[25502]"
//line /usr/local/go/src/bufio/bufio.go:780
			}
//line /usr/local/go/src/bufio/bufio.go:780
			// _ = "end of CoverTab[25500]"
		} else {
//line /usr/local/go/src/bufio/bufio.go:781
			_go_fuzz_dep_.CoverTab[25503]++
//line /usr/local/go/src/bufio/bufio.go:781
			// _ = "end of CoverTab[25503]"
//line /usr/local/go/src/bufio/bufio.go:781
		}
//line /usr/local/go/src/bufio/bufio.go:781
		// _ = "end of CoverTab[25495]"
//line /usr/local/go/src/bufio/bufio.go:781
		_go_fuzz_dep_.CoverTab[25496]++
							if readerFromOK && func() bool {
//line /usr/local/go/src/bufio/bufio.go:782
			_go_fuzz_dep_.CoverTab[25504]++
//line /usr/local/go/src/bufio/bufio.go:782
			return b.Buffered() == 0
//line /usr/local/go/src/bufio/bufio.go:782
			// _ = "end of CoverTab[25504]"
//line /usr/local/go/src/bufio/bufio.go:782
		}() {
//line /usr/local/go/src/bufio/bufio.go:782
			_go_fuzz_dep_.CoverTab[25505]++
								nn, err := readerFrom.ReadFrom(r)
								b.err = err
								n += nn
								return n, err
//line /usr/local/go/src/bufio/bufio.go:786
			// _ = "end of CoverTab[25505]"
		} else {
//line /usr/local/go/src/bufio/bufio.go:787
			_go_fuzz_dep_.CoverTab[25506]++
//line /usr/local/go/src/bufio/bufio.go:787
			// _ = "end of CoverTab[25506]"
//line /usr/local/go/src/bufio/bufio.go:787
		}
//line /usr/local/go/src/bufio/bufio.go:787
		// _ = "end of CoverTab[25496]"
//line /usr/local/go/src/bufio/bufio.go:787
		_go_fuzz_dep_.CoverTab[25497]++
							nr := 0
							for nr < maxConsecutiveEmptyReads {
//line /usr/local/go/src/bufio/bufio.go:789
			_go_fuzz_dep_.CoverTab[25507]++
								m, err = r.Read(b.buf[b.n:])
								if m != 0 || func() bool {
//line /usr/local/go/src/bufio/bufio.go:791
				_go_fuzz_dep_.CoverTab[25509]++
//line /usr/local/go/src/bufio/bufio.go:791
				return err != nil
//line /usr/local/go/src/bufio/bufio.go:791
				// _ = "end of CoverTab[25509]"
//line /usr/local/go/src/bufio/bufio.go:791
			}() {
//line /usr/local/go/src/bufio/bufio.go:791
				_go_fuzz_dep_.CoverTab[25510]++
									break
//line /usr/local/go/src/bufio/bufio.go:792
				// _ = "end of CoverTab[25510]"
			} else {
//line /usr/local/go/src/bufio/bufio.go:793
				_go_fuzz_dep_.CoverTab[25511]++
//line /usr/local/go/src/bufio/bufio.go:793
				// _ = "end of CoverTab[25511]"
//line /usr/local/go/src/bufio/bufio.go:793
			}
//line /usr/local/go/src/bufio/bufio.go:793
			// _ = "end of CoverTab[25507]"
//line /usr/local/go/src/bufio/bufio.go:793
			_go_fuzz_dep_.CoverTab[25508]++
								nr++
//line /usr/local/go/src/bufio/bufio.go:794
			// _ = "end of CoverTab[25508]"
		}
//line /usr/local/go/src/bufio/bufio.go:795
		// _ = "end of CoverTab[25497]"
//line /usr/local/go/src/bufio/bufio.go:795
		_go_fuzz_dep_.CoverTab[25498]++
							if nr == maxConsecutiveEmptyReads {
//line /usr/local/go/src/bufio/bufio.go:796
			_go_fuzz_dep_.CoverTab[25512]++
								return n, io.ErrNoProgress
//line /usr/local/go/src/bufio/bufio.go:797
			// _ = "end of CoverTab[25512]"
		} else {
//line /usr/local/go/src/bufio/bufio.go:798
			_go_fuzz_dep_.CoverTab[25513]++
//line /usr/local/go/src/bufio/bufio.go:798
			// _ = "end of CoverTab[25513]"
//line /usr/local/go/src/bufio/bufio.go:798
		}
//line /usr/local/go/src/bufio/bufio.go:798
		// _ = "end of CoverTab[25498]"
//line /usr/local/go/src/bufio/bufio.go:798
		_go_fuzz_dep_.CoverTab[25499]++
							b.n += m
							n += int64(m)
							if err != nil {
//line /usr/local/go/src/bufio/bufio.go:801
			_go_fuzz_dep_.CoverTab[25514]++
								break
//line /usr/local/go/src/bufio/bufio.go:802
			// _ = "end of CoverTab[25514]"
		} else {
//line /usr/local/go/src/bufio/bufio.go:803
			_go_fuzz_dep_.CoverTab[25515]++
//line /usr/local/go/src/bufio/bufio.go:803
			// _ = "end of CoverTab[25515]"
//line /usr/local/go/src/bufio/bufio.go:803
		}
//line /usr/local/go/src/bufio/bufio.go:803
		// _ = "end of CoverTab[25499]"
	}
//line /usr/local/go/src/bufio/bufio.go:804
	// _ = "end of CoverTab[25490]"
//line /usr/local/go/src/bufio/bufio.go:804
	_go_fuzz_dep_.CoverTab[25491]++
						if err == io.EOF {
//line /usr/local/go/src/bufio/bufio.go:805
		_go_fuzz_dep_.CoverTab[25516]++

							if b.Available() == 0 {
//line /usr/local/go/src/bufio/bufio.go:807
			_go_fuzz_dep_.CoverTab[25517]++
								err = b.Flush()
//line /usr/local/go/src/bufio/bufio.go:808
			// _ = "end of CoverTab[25517]"
		} else {
//line /usr/local/go/src/bufio/bufio.go:809
			_go_fuzz_dep_.CoverTab[25518]++
								err = nil
//line /usr/local/go/src/bufio/bufio.go:810
			// _ = "end of CoverTab[25518]"
		}
//line /usr/local/go/src/bufio/bufio.go:811
		// _ = "end of CoverTab[25516]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:812
		_go_fuzz_dep_.CoverTab[25519]++
//line /usr/local/go/src/bufio/bufio.go:812
		// _ = "end of CoverTab[25519]"
//line /usr/local/go/src/bufio/bufio.go:812
	}
//line /usr/local/go/src/bufio/bufio.go:812
	// _ = "end of CoverTab[25491]"
//line /usr/local/go/src/bufio/bufio.go:812
	_go_fuzz_dep_.CoverTab[25492]++
						return n, err
//line /usr/local/go/src/bufio/bufio.go:813
	// _ = "end of CoverTab[25492]"
}

//line /usr/local/go/src/bufio/bufio.go:818
// ReadWriter stores pointers to a Reader and a Writer.
//line /usr/local/go/src/bufio/bufio.go:818
// It implements io.ReadWriter.
//line /usr/local/go/src/bufio/bufio.go:820
type ReadWriter struct {
	*Reader
	*Writer
}

// NewReadWriter allocates a new ReadWriter that dispatches to r and w.
func NewReadWriter(r *Reader, w *Writer) *ReadWriter {
//line /usr/local/go/src/bufio/bufio.go:826
	_go_fuzz_dep_.CoverTab[25520]++
						return &ReadWriter{r, w}
//line /usr/local/go/src/bufio/bufio.go:827
	// _ = "end of CoverTab[25520]"
}

//line /usr/local/go/src/bufio/bufio.go:828
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/bufio/bufio.go:828
var _ = _go_fuzz_dep_.CoverTab
