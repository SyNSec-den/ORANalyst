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
	_go_fuzz_dep_.CoverTab[1785]++

						b, ok := rd.(*Reader)
						if ok && func() bool {
//line /usr/local/go/src/bufio/bufio.go:50
		_go_fuzz_dep_.CoverTab[1788]++
//line /usr/local/go/src/bufio/bufio.go:50
		return len(b.buf) >= size
//line /usr/local/go/src/bufio/bufio.go:50
		// _ = "end of CoverTab[1788]"
//line /usr/local/go/src/bufio/bufio.go:50
	}() {
//line /usr/local/go/src/bufio/bufio.go:50
		_go_fuzz_dep_.CoverTab[1789]++
							return b
//line /usr/local/go/src/bufio/bufio.go:51
		// _ = "end of CoverTab[1789]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:52
		_go_fuzz_dep_.CoverTab[1790]++
//line /usr/local/go/src/bufio/bufio.go:52
		// _ = "end of CoverTab[1790]"
//line /usr/local/go/src/bufio/bufio.go:52
	}
//line /usr/local/go/src/bufio/bufio.go:52
	// _ = "end of CoverTab[1785]"
//line /usr/local/go/src/bufio/bufio.go:52
	_go_fuzz_dep_.CoverTab[1786]++
						if size < minReadBufferSize {
//line /usr/local/go/src/bufio/bufio.go:53
		_go_fuzz_dep_.CoverTab[1791]++
							size = minReadBufferSize
//line /usr/local/go/src/bufio/bufio.go:54
		// _ = "end of CoverTab[1791]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:55
		_go_fuzz_dep_.CoverTab[1792]++
//line /usr/local/go/src/bufio/bufio.go:55
		// _ = "end of CoverTab[1792]"
//line /usr/local/go/src/bufio/bufio.go:55
	}
//line /usr/local/go/src/bufio/bufio.go:55
	// _ = "end of CoverTab[1786]"
//line /usr/local/go/src/bufio/bufio.go:55
	_go_fuzz_dep_.CoverTab[1787]++
						r := new(Reader)
						r.reset(make([]byte, size), rd)
						return r
//line /usr/local/go/src/bufio/bufio.go:58
	// _ = "end of CoverTab[1787]"
}

// NewReader returns a new Reader whose buffer has the default size.
func NewReader(rd io.Reader) *Reader {
//line /usr/local/go/src/bufio/bufio.go:62
	_go_fuzz_dep_.CoverTab[1793]++
						return NewReaderSize(rd, defaultBufSize)
//line /usr/local/go/src/bufio/bufio.go:63
	// _ = "end of CoverTab[1793]"
}

// Size returns the size of the underlying buffer in bytes.
func (b *Reader) Size() int {
//line /usr/local/go/src/bufio/bufio.go:67
	_go_fuzz_dep_.CoverTab[1794]++
//line /usr/local/go/src/bufio/bufio.go:67
	return len(b.buf)
//line /usr/local/go/src/bufio/bufio.go:67
	// _ = "end of CoverTab[1794]"
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
	_go_fuzz_dep_.CoverTab[1795]++
						if b.buf == nil {
//line /usr/local/go/src/bufio/bufio.go:74
		_go_fuzz_dep_.CoverTab[1797]++
							b.buf = make([]byte, defaultBufSize)
//line /usr/local/go/src/bufio/bufio.go:75
		// _ = "end of CoverTab[1797]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:76
		_go_fuzz_dep_.CoverTab[1798]++
//line /usr/local/go/src/bufio/bufio.go:76
		// _ = "end of CoverTab[1798]"
//line /usr/local/go/src/bufio/bufio.go:76
	}
//line /usr/local/go/src/bufio/bufio.go:76
	// _ = "end of CoverTab[1795]"
//line /usr/local/go/src/bufio/bufio.go:76
	_go_fuzz_dep_.CoverTab[1796]++
						b.reset(b.buf, r)
//line /usr/local/go/src/bufio/bufio.go:77
	// _ = "end of CoverTab[1796]"
}

func (b *Reader) reset(buf []byte, r io.Reader) {
//line /usr/local/go/src/bufio/bufio.go:80
	_go_fuzz_dep_.CoverTab[1799]++
						*b = Reader{
		buf:		buf,
		rd:		r,
		lastByte:	-1,
		lastRuneSize:	-1,
	}
//line /usr/local/go/src/bufio/bufio.go:86
	// _ = "end of CoverTab[1799]"
}

var errNegativeRead = errors.New("bufio: reader returned negative count from Read")

// fill reads a new chunk into the buffer.
func (b *Reader) fill() {
//line /usr/local/go/src/bufio/bufio.go:92
	_go_fuzz_dep_.CoverTab[1800]++

						if b.r > 0 {
//line /usr/local/go/src/bufio/bufio.go:94
		_go_fuzz_dep_.CoverTab[1804]++
							copy(b.buf, b.buf[b.r:b.w])
							b.w -= b.r
							b.r = 0
//line /usr/local/go/src/bufio/bufio.go:97
		// _ = "end of CoverTab[1804]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:98
		_go_fuzz_dep_.CoverTab[1805]++
//line /usr/local/go/src/bufio/bufio.go:98
		// _ = "end of CoverTab[1805]"
//line /usr/local/go/src/bufio/bufio.go:98
	}
//line /usr/local/go/src/bufio/bufio.go:98
	// _ = "end of CoverTab[1800]"
//line /usr/local/go/src/bufio/bufio.go:98
	_go_fuzz_dep_.CoverTab[1801]++

						if b.w >= len(b.buf) {
//line /usr/local/go/src/bufio/bufio.go:100
		_go_fuzz_dep_.CoverTab[1806]++
							panic("bufio: tried to fill full buffer")
//line /usr/local/go/src/bufio/bufio.go:101
		// _ = "end of CoverTab[1806]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:102
		_go_fuzz_dep_.CoverTab[1807]++
//line /usr/local/go/src/bufio/bufio.go:102
		// _ = "end of CoverTab[1807]"
//line /usr/local/go/src/bufio/bufio.go:102
	}
//line /usr/local/go/src/bufio/bufio.go:102
	// _ = "end of CoverTab[1801]"
//line /usr/local/go/src/bufio/bufio.go:102
	_go_fuzz_dep_.CoverTab[1802]++

//line /usr/local/go/src/bufio/bufio.go:105
	for i := maxConsecutiveEmptyReads; i > 0; i-- {
//line /usr/local/go/src/bufio/bufio.go:105
		_go_fuzz_dep_.CoverTab[1808]++
							n, err := b.rd.Read(b.buf[b.w:])
							if n < 0 {
//line /usr/local/go/src/bufio/bufio.go:107
			_go_fuzz_dep_.CoverTab[1811]++
								panic(errNegativeRead)
//line /usr/local/go/src/bufio/bufio.go:108
			// _ = "end of CoverTab[1811]"
		} else {
//line /usr/local/go/src/bufio/bufio.go:109
			_go_fuzz_dep_.CoverTab[1812]++
//line /usr/local/go/src/bufio/bufio.go:109
			// _ = "end of CoverTab[1812]"
//line /usr/local/go/src/bufio/bufio.go:109
		}
//line /usr/local/go/src/bufio/bufio.go:109
		// _ = "end of CoverTab[1808]"
//line /usr/local/go/src/bufio/bufio.go:109
		_go_fuzz_dep_.CoverTab[1809]++
							b.w += n
							if err != nil {
//line /usr/local/go/src/bufio/bufio.go:111
			_go_fuzz_dep_.CoverTab[1813]++
								b.err = err
								return
//line /usr/local/go/src/bufio/bufio.go:113
			// _ = "end of CoverTab[1813]"
		} else {
//line /usr/local/go/src/bufio/bufio.go:114
			_go_fuzz_dep_.CoverTab[1814]++
//line /usr/local/go/src/bufio/bufio.go:114
			// _ = "end of CoverTab[1814]"
//line /usr/local/go/src/bufio/bufio.go:114
		}
//line /usr/local/go/src/bufio/bufio.go:114
		// _ = "end of CoverTab[1809]"
//line /usr/local/go/src/bufio/bufio.go:114
		_go_fuzz_dep_.CoverTab[1810]++
							if n > 0 {
//line /usr/local/go/src/bufio/bufio.go:115
			_go_fuzz_dep_.CoverTab[1815]++
								return
//line /usr/local/go/src/bufio/bufio.go:116
			// _ = "end of CoverTab[1815]"
		} else {
//line /usr/local/go/src/bufio/bufio.go:117
			_go_fuzz_dep_.CoverTab[1816]++
//line /usr/local/go/src/bufio/bufio.go:117
			// _ = "end of CoverTab[1816]"
//line /usr/local/go/src/bufio/bufio.go:117
		}
//line /usr/local/go/src/bufio/bufio.go:117
		// _ = "end of CoverTab[1810]"
	}
//line /usr/local/go/src/bufio/bufio.go:118
	// _ = "end of CoverTab[1802]"
//line /usr/local/go/src/bufio/bufio.go:118
	_go_fuzz_dep_.CoverTab[1803]++
						b.err = io.ErrNoProgress
//line /usr/local/go/src/bufio/bufio.go:119
	// _ = "end of CoverTab[1803]"
}

func (b *Reader) readErr() error {
//line /usr/local/go/src/bufio/bufio.go:122
	_go_fuzz_dep_.CoverTab[1817]++
						err := b.err
						b.err = nil
						return err
//line /usr/local/go/src/bufio/bufio.go:125
	// _ = "end of CoverTab[1817]"
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
	_go_fuzz_dep_.CoverTab[1818]++
						if n < 0 {
//line /usr/local/go/src/bufio/bufio.go:136
		_go_fuzz_dep_.CoverTab[1823]++
							return nil, ErrNegativeCount
//line /usr/local/go/src/bufio/bufio.go:137
		// _ = "end of CoverTab[1823]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:138
		_go_fuzz_dep_.CoverTab[1824]++
//line /usr/local/go/src/bufio/bufio.go:138
		// _ = "end of CoverTab[1824]"
//line /usr/local/go/src/bufio/bufio.go:138
	}
//line /usr/local/go/src/bufio/bufio.go:138
	// _ = "end of CoverTab[1818]"
//line /usr/local/go/src/bufio/bufio.go:138
	_go_fuzz_dep_.CoverTab[1819]++

						b.lastByte = -1
						b.lastRuneSize = -1

						for b.w-b.r < n && func() bool {
//line /usr/local/go/src/bufio/bufio.go:143
		_go_fuzz_dep_.CoverTab[1825]++
//line /usr/local/go/src/bufio/bufio.go:143
		return b.w-b.r < len(b.buf)
//line /usr/local/go/src/bufio/bufio.go:143
		// _ = "end of CoverTab[1825]"
//line /usr/local/go/src/bufio/bufio.go:143
	}() && func() bool {
//line /usr/local/go/src/bufio/bufio.go:143
		_go_fuzz_dep_.CoverTab[1826]++
//line /usr/local/go/src/bufio/bufio.go:143
		return b.err == nil
//line /usr/local/go/src/bufio/bufio.go:143
		// _ = "end of CoverTab[1826]"
//line /usr/local/go/src/bufio/bufio.go:143
	}() {
//line /usr/local/go/src/bufio/bufio.go:143
		_go_fuzz_dep_.CoverTab[1827]++
							b.fill()
//line /usr/local/go/src/bufio/bufio.go:144
		// _ = "end of CoverTab[1827]"
	}
//line /usr/local/go/src/bufio/bufio.go:145
	// _ = "end of CoverTab[1819]"
//line /usr/local/go/src/bufio/bufio.go:145
	_go_fuzz_dep_.CoverTab[1820]++

						if n > len(b.buf) {
//line /usr/local/go/src/bufio/bufio.go:147
		_go_fuzz_dep_.CoverTab[1828]++
							return b.buf[b.r:b.w], ErrBufferFull
//line /usr/local/go/src/bufio/bufio.go:148
		// _ = "end of CoverTab[1828]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:149
		_go_fuzz_dep_.CoverTab[1829]++
//line /usr/local/go/src/bufio/bufio.go:149
		// _ = "end of CoverTab[1829]"
//line /usr/local/go/src/bufio/bufio.go:149
	}
//line /usr/local/go/src/bufio/bufio.go:149
	// _ = "end of CoverTab[1820]"
//line /usr/local/go/src/bufio/bufio.go:149
	_go_fuzz_dep_.CoverTab[1821]++

	// 0 <= n <= len(b.buf)
	var err error
	if avail := b.w - b.r; avail < n {
//line /usr/local/go/src/bufio/bufio.go:153
		_go_fuzz_dep_.CoverTab[1830]++

							n = avail
							err = b.readErr()
							if err == nil {
//line /usr/local/go/src/bufio/bufio.go:157
			_go_fuzz_dep_.CoverTab[1831]++
								err = ErrBufferFull
//line /usr/local/go/src/bufio/bufio.go:158
			// _ = "end of CoverTab[1831]"
		} else {
//line /usr/local/go/src/bufio/bufio.go:159
			_go_fuzz_dep_.CoverTab[1832]++
//line /usr/local/go/src/bufio/bufio.go:159
			// _ = "end of CoverTab[1832]"
//line /usr/local/go/src/bufio/bufio.go:159
		}
//line /usr/local/go/src/bufio/bufio.go:159
		// _ = "end of CoverTab[1830]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:160
		_go_fuzz_dep_.CoverTab[1833]++
//line /usr/local/go/src/bufio/bufio.go:160
		// _ = "end of CoverTab[1833]"
//line /usr/local/go/src/bufio/bufio.go:160
	}
//line /usr/local/go/src/bufio/bufio.go:160
	// _ = "end of CoverTab[1821]"
//line /usr/local/go/src/bufio/bufio.go:160
	_go_fuzz_dep_.CoverTab[1822]++
						return b.buf[b.r : b.r+n], err
//line /usr/local/go/src/bufio/bufio.go:161
	// _ = "end of CoverTab[1822]"
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
	_go_fuzz_dep_.CoverTab[1834]++
						if n < 0 {
//line /usr/local/go/src/bufio/bufio.go:170
		_go_fuzz_dep_.CoverTab[1837]++
							return 0, ErrNegativeCount
//line /usr/local/go/src/bufio/bufio.go:171
		// _ = "end of CoverTab[1837]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:172
		_go_fuzz_dep_.CoverTab[1838]++
//line /usr/local/go/src/bufio/bufio.go:172
		// _ = "end of CoverTab[1838]"
//line /usr/local/go/src/bufio/bufio.go:172
	}
//line /usr/local/go/src/bufio/bufio.go:172
	// _ = "end of CoverTab[1834]"
//line /usr/local/go/src/bufio/bufio.go:172
	_go_fuzz_dep_.CoverTab[1835]++
						if n == 0 {
//line /usr/local/go/src/bufio/bufio.go:173
		_go_fuzz_dep_.CoverTab[1839]++
							return
//line /usr/local/go/src/bufio/bufio.go:174
		// _ = "end of CoverTab[1839]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:175
		_go_fuzz_dep_.CoverTab[1840]++
//line /usr/local/go/src/bufio/bufio.go:175
		// _ = "end of CoverTab[1840]"
//line /usr/local/go/src/bufio/bufio.go:175
	}
//line /usr/local/go/src/bufio/bufio.go:175
	// _ = "end of CoverTab[1835]"
//line /usr/local/go/src/bufio/bufio.go:175
	_go_fuzz_dep_.CoverTab[1836]++

						b.lastByte = -1
						b.lastRuneSize = -1

						remain := n
						for {
//line /usr/local/go/src/bufio/bufio.go:181
		_go_fuzz_dep_.CoverTab[1841]++
							skip := b.Buffered()
							if skip == 0 {
//line /usr/local/go/src/bufio/bufio.go:183
			_go_fuzz_dep_.CoverTab[1845]++
								b.fill()
								skip = b.Buffered()
//line /usr/local/go/src/bufio/bufio.go:185
			// _ = "end of CoverTab[1845]"
		} else {
//line /usr/local/go/src/bufio/bufio.go:186
			_go_fuzz_dep_.CoverTab[1846]++
//line /usr/local/go/src/bufio/bufio.go:186
			// _ = "end of CoverTab[1846]"
//line /usr/local/go/src/bufio/bufio.go:186
		}
//line /usr/local/go/src/bufio/bufio.go:186
		// _ = "end of CoverTab[1841]"
//line /usr/local/go/src/bufio/bufio.go:186
		_go_fuzz_dep_.CoverTab[1842]++
							if skip > remain {
//line /usr/local/go/src/bufio/bufio.go:187
			_go_fuzz_dep_.CoverTab[1847]++
								skip = remain
//line /usr/local/go/src/bufio/bufio.go:188
			// _ = "end of CoverTab[1847]"
		} else {
//line /usr/local/go/src/bufio/bufio.go:189
			_go_fuzz_dep_.CoverTab[1848]++
//line /usr/local/go/src/bufio/bufio.go:189
			// _ = "end of CoverTab[1848]"
//line /usr/local/go/src/bufio/bufio.go:189
		}
//line /usr/local/go/src/bufio/bufio.go:189
		// _ = "end of CoverTab[1842]"
//line /usr/local/go/src/bufio/bufio.go:189
		_go_fuzz_dep_.CoverTab[1843]++
							b.r += skip
							remain -= skip
							if remain == 0 {
//line /usr/local/go/src/bufio/bufio.go:192
			_go_fuzz_dep_.CoverTab[1849]++
								return n, nil
//line /usr/local/go/src/bufio/bufio.go:193
			// _ = "end of CoverTab[1849]"
		} else {
//line /usr/local/go/src/bufio/bufio.go:194
			_go_fuzz_dep_.CoverTab[1850]++
//line /usr/local/go/src/bufio/bufio.go:194
			// _ = "end of CoverTab[1850]"
//line /usr/local/go/src/bufio/bufio.go:194
		}
//line /usr/local/go/src/bufio/bufio.go:194
		// _ = "end of CoverTab[1843]"
//line /usr/local/go/src/bufio/bufio.go:194
		_go_fuzz_dep_.CoverTab[1844]++
							if b.err != nil {
//line /usr/local/go/src/bufio/bufio.go:195
			_go_fuzz_dep_.CoverTab[1851]++
								return n - remain, b.readErr()
//line /usr/local/go/src/bufio/bufio.go:196
			// _ = "end of CoverTab[1851]"
		} else {
//line /usr/local/go/src/bufio/bufio.go:197
			_go_fuzz_dep_.CoverTab[1852]++
//line /usr/local/go/src/bufio/bufio.go:197
			// _ = "end of CoverTab[1852]"
//line /usr/local/go/src/bufio/bufio.go:197
		}
//line /usr/local/go/src/bufio/bufio.go:197
		// _ = "end of CoverTab[1844]"
	}
//line /usr/local/go/src/bufio/bufio.go:198
	// _ = "end of CoverTab[1836]"
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
	_go_fuzz_dep_.CoverTab[1853]++
						n = len(p)
						if n == 0 {
//line /usr/local/go/src/bufio/bufio.go:210
		_go_fuzz_dep_.CoverTab[1856]++
							if b.Buffered() > 0 {
//line /usr/local/go/src/bufio/bufio.go:211
			_go_fuzz_dep_.CoverTab[1858]++
								return 0, nil
//line /usr/local/go/src/bufio/bufio.go:212
			// _ = "end of CoverTab[1858]"
		} else {
//line /usr/local/go/src/bufio/bufio.go:213
			_go_fuzz_dep_.CoverTab[1859]++
//line /usr/local/go/src/bufio/bufio.go:213
			// _ = "end of CoverTab[1859]"
//line /usr/local/go/src/bufio/bufio.go:213
		}
//line /usr/local/go/src/bufio/bufio.go:213
		// _ = "end of CoverTab[1856]"
//line /usr/local/go/src/bufio/bufio.go:213
		_go_fuzz_dep_.CoverTab[1857]++
							return 0, b.readErr()
//line /usr/local/go/src/bufio/bufio.go:214
		// _ = "end of CoverTab[1857]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:215
		_go_fuzz_dep_.CoverTab[1860]++
//line /usr/local/go/src/bufio/bufio.go:215
		// _ = "end of CoverTab[1860]"
//line /usr/local/go/src/bufio/bufio.go:215
	}
//line /usr/local/go/src/bufio/bufio.go:215
	// _ = "end of CoverTab[1853]"
//line /usr/local/go/src/bufio/bufio.go:215
	_go_fuzz_dep_.CoverTab[1854]++
						if b.r == b.w {
//line /usr/local/go/src/bufio/bufio.go:216
		_go_fuzz_dep_.CoverTab[1861]++
							if b.err != nil {
//line /usr/local/go/src/bufio/bufio.go:217
			_go_fuzz_dep_.CoverTab[1866]++
								return 0, b.readErr()
//line /usr/local/go/src/bufio/bufio.go:218
			// _ = "end of CoverTab[1866]"
		} else {
//line /usr/local/go/src/bufio/bufio.go:219
			_go_fuzz_dep_.CoverTab[1867]++
//line /usr/local/go/src/bufio/bufio.go:219
			// _ = "end of CoverTab[1867]"
//line /usr/local/go/src/bufio/bufio.go:219
		}
//line /usr/local/go/src/bufio/bufio.go:219
		// _ = "end of CoverTab[1861]"
//line /usr/local/go/src/bufio/bufio.go:219
		_go_fuzz_dep_.CoverTab[1862]++
							if len(p) >= len(b.buf) {
//line /usr/local/go/src/bufio/bufio.go:220
			_go_fuzz_dep_.CoverTab[1868]++

//line /usr/local/go/src/bufio/bufio.go:223
			n, b.err = b.rd.Read(p)
			if n < 0 {
//line /usr/local/go/src/bufio/bufio.go:224
				_go_fuzz_dep_.CoverTab[1871]++
									panic(errNegativeRead)
//line /usr/local/go/src/bufio/bufio.go:225
				// _ = "end of CoverTab[1871]"
			} else {
//line /usr/local/go/src/bufio/bufio.go:226
				_go_fuzz_dep_.CoverTab[1872]++
//line /usr/local/go/src/bufio/bufio.go:226
				// _ = "end of CoverTab[1872]"
//line /usr/local/go/src/bufio/bufio.go:226
			}
//line /usr/local/go/src/bufio/bufio.go:226
			// _ = "end of CoverTab[1868]"
//line /usr/local/go/src/bufio/bufio.go:226
			_go_fuzz_dep_.CoverTab[1869]++
								if n > 0 {
//line /usr/local/go/src/bufio/bufio.go:227
				_go_fuzz_dep_.CoverTab[1873]++
									b.lastByte = int(p[n-1])
									b.lastRuneSize = -1
//line /usr/local/go/src/bufio/bufio.go:229
				// _ = "end of CoverTab[1873]"
			} else {
//line /usr/local/go/src/bufio/bufio.go:230
				_go_fuzz_dep_.CoverTab[1874]++
//line /usr/local/go/src/bufio/bufio.go:230
				// _ = "end of CoverTab[1874]"
//line /usr/local/go/src/bufio/bufio.go:230
			}
//line /usr/local/go/src/bufio/bufio.go:230
			// _ = "end of CoverTab[1869]"
//line /usr/local/go/src/bufio/bufio.go:230
			_go_fuzz_dep_.CoverTab[1870]++
								return n, b.readErr()
//line /usr/local/go/src/bufio/bufio.go:231
			// _ = "end of CoverTab[1870]"
		} else {
//line /usr/local/go/src/bufio/bufio.go:232
			_go_fuzz_dep_.CoverTab[1875]++
//line /usr/local/go/src/bufio/bufio.go:232
			// _ = "end of CoverTab[1875]"
//line /usr/local/go/src/bufio/bufio.go:232
		}
//line /usr/local/go/src/bufio/bufio.go:232
		// _ = "end of CoverTab[1862]"
//line /usr/local/go/src/bufio/bufio.go:232
		_go_fuzz_dep_.CoverTab[1863]++

//line /usr/local/go/src/bufio/bufio.go:235
		b.r = 0
		b.w = 0
		n, b.err = b.rd.Read(b.buf)
		if n < 0 {
//line /usr/local/go/src/bufio/bufio.go:238
			_go_fuzz_dep_.CoverTab[1876]++
								panic(errNegativeRead)
//line /usr/local/go/src/bufio/bufio.go:239
			// _ = "end of CoverTab[1876]"
		} else {
//line /usr/local/go/src/bufio/bufio.go:240
			_go_fuzz_dep_.CoverTab[1877]++
//line /usr/local/go/src/bufio/bufio.go:240
			// _ = "end of CoverTab[1877]"
//line /usr/local/go/src/bufio/bufio.go:240
		}
//line /usr/local/go/src/bufio/bufio.go:240
		// _ = "end of CoverTab[1863]"
//line /usr/local/go/src/bufio/bufio.go:240
		_go_fuzz_dep_.CoverTab[1864]++
							if n == 0 {
//line /usr/local/go/src/bufio/bufio.go:241
			_go_fuzz_dep_.CoverTab[1878]++
								return 0, b.readErr()
//line /usr/local/go/src/bufio/bufio.go:242
			// _ = "end of CoverTab[1878]"
		} else {
//line /usr/local/go/src/bufio/bufio.go:243
			_go_fuzz_dep_.CoverTab[1879]++
//line /usr/local/go/src/bufio/bufio.go:243
			// _ = "end of CoverTab[1879]"
//line /usr/local/go/src/bufio/bufio.go:243
		}
//line /usr/local/go/src/bufio/bufio.go:243
		// _ = "end of CoverTab[1864]"
//line /usr/local/go/src/bufio/bufio.go:243
		_go_fuzz_dep_.CoverTab[1865]++
							b.w += n
//line /usr/local/go/src/bufio/bufio.go:244
		// _ = "end of CoverTab[1865]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:245
		_go_fuzz_dep_.CoverTab[1880]++
//line /usr/local/go/src/bufio/bufio.go:245
		// _ = "end of CoverTab[1880]"
//line /usr/local/go/src/bufio/bufio.go:245
	}
//line /usr/local/go/src/bufio/bufio.go:245
	// _ = "end of CoverTab[1854]"
//line /usr/local/go/src/bufio/bufio.go:245
	_go_fuzz_dep_.CoverTab[1855]++

//line /usr/local/go/src/bufio/bufio.go:250
	n = copy(p, b.buf[b.r:b.w])
						b.r += n
						b.lastByte = int(b.buf[b.r-1])
						b.lastRuneSize = -1
						return n, nil
//line /usr/local/go/src/bufio/bufio.go:254
	// _ = "end of CoverTab[1855]"
}

// ReadByte reads and returns a single byte.
//line /usr/local/go/src/bufio/bufio.go:257
// If no byte is available, returns an error.
//line /usr/local/go/src/bufio/bufio.go:259
func (b *Reader) ReadByte() (byte, error) {
//line /usr/local/go/src/bufio/bufio.go:259
	_go_fuzz_dep_.CoverTab[1881]++
						b.lastRuneSize = -1
						for b.r == b.w {
//line /usr/local/go/src/bufio/bufio.go:261
		_go_fuzz_dep_.CoverTab[1883]++
							if b.err != nil {
//line /usr/local/go/src/bufio/bufio.go:262
			_go_fuzz_dep_.CoverTab[1885]++
								return 0, b.readErr()
//line /usr/local/go/src/bufio/bufio.go:263
			// _ = "end of CoverTab[1885]"
		} else {
//line /usr/local/go/src/bufio/bufio.go:264
			_go_fuzz_dep_.CoverTab[1886]++
//line /usr/local/go/src/bufio/bufio.go:264
			// _ = "end of CoverTab[1886]"
//line /usr/local/go/src/bufio/bufio.go:264
		}
//line /usr/local/go/src/bufio/bufio.go:264
		// _ = "end of CoverTab[1883]"
//line /usr/local/go/src/bufio/bufio.go:264
		_go_fuzz_dep_.CoverTab[1884]++
							b.fill()
//line /usr/local/go/src/bufio/bufio.go:265
		// _ = "end of CoverTab[1884]"
	}
//line /usr/local/go/src/bufio/bufio.go:266
	// _ = "end of CoverTab[1881]"
//line /usr/local/go/src/bufio/bufio.go:266
	_go_fuzz_dep_.CoverTab[1882]++
						c := b.buf[b.r]
						b.r++
						b.lastByte = int(c)
						return c, nil
//line /usr/local/go/src/bufio/bufio.go:270
	// _ = "end of CoverTab[1882]"
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
	_go_fuzz_dep_.CoverTab[1887]++
						if b.lastByte < 0 || func() bool {
//line /usr/local/go/src/bufio/bufio.go:279
		_go_fuzz_dep_.CoverTab[1890]++
//line /usr/local/go/src/bufio/bufio.go:279
		return b.r == 0 && func() bool {
//line /usr/local/go/src/bufio/bufio.go:279
			_go_fuzz_dep_.CoverTab[1891]++
//line /usr/local/go/src/bufio/bufio.go:279
			return b.w > 0
//line /usr/local/go/src/bufio/bufio.go:279
			// _ = "end of CoverTab[1891]"
//line /usr/local/go/src/bufio/bufio.go:279
		}()
//line /usr/local/go/src/bufio/bufio.go:279
		// _ = "end of CoverTab[1890]"
//line /usr/local/go/src/bufio/bufio.go:279
	}() {
//line /usr/local/go/src/bufio/bufio.go:279
		_go_fuzz_dep_.CoverTab[1892]++
							return ErrInvalidUnreadByte
//line /usr/local/go/src/bufio/bufio.go:280
		// _ = "end of CoverTab[1892]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:281
		_go_fuzz_dep_.CoverTab[1893]++
//line /usr/local/go/src/bufio/bufio.go:281
		// _ = "end of CoverTab[1893]"
//line /usr/local/go/src/bufio/bufio.go:281
	}
//line /usr/local/go/src/bufio/bufio.go:281
	// _ = "end of CoverTab[1887]"
//line /usr/local/go/src/bufio/bufio.go:281
	_go_fuzz_dep_.CoverTab[1888]++

						if b.r > 0 {
//line /usr/local/go/src/bufio/bufio.go:283
		_go_fuzz_dep_.CoverTab[1894]++
							b.r--
//line /usr/local/go/src/bufio/bufio.go:284
		// _ = "end of CoverTab[1894]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:285
		_go_fuzz_dep_.CoverTab[1895]++

							b.w = 1
//line /usr/local/go/src/bufio/bufio.go:287
		// _ = "end of CoverTab[1895]"
	}
//line /usr/local/go/src/bufio/bufio.go:288
	// _ = "end of CoverTab[1888]"
//line /usr/local/go/src/bufio/bufio.go:288
	_go_fuzz_dep_.CoverTab[1889]++
						b.buf[b.r] = byte(b.lastByte)
						b.lastByte = -1
						b.lastRuneSize = -1
						return nil
//line /usr/local/go/src/bufio/bufio.go:292
	// _ = "end of CoverTab[1889]"
}

// ReadRune reads a single UTF-8 encoded Unicode character and returns the
//line /usr/local/go/src/bufio/bufio.go:295
// rune and its size in bytes. If the encoded rune is invalid, it consumes one byte
//line /usr/local/go/src/bufio/bufio.go:295
// and returns unicode.ReplacementChar (U+FFFD) with a size of 1.
//line /usr/local/go/src/bufio/bufio.go:298
func (b *Reader) ReadRune() (r rune, size int, err error) {
//line /usr/local/go/src/bufio/bufio.go:298
	_go_fuzz_dep_.CoverTab[1896]++
						for b.r+utf8.UTFMax > b.w && func() bool {
//line /usr/local/go/src/bufio/bufio.go:299
		_go_fuzz_dep_.CoverTab[1900]++
//line /usr/local/go/src/bufio/bufio.go:299
		return !utf8.FullRune(b.buf[b.r:b.w])
//line /usr/local/go/src/bufio/bufio.go:299
		// _ = "end of CoverTab[1900]"
//line /usr/local/go/src/bufio/bufio.go:299
	}() && func() bool {
//line /usr/local/go/src/bufio/bufio.go:299
		_go_fuzz_dep_.CoverTab[1901]++
//line /usr/local/go/src/bufio/bufio.go:299
		return b.err == nil
//line /usr/local/go/src/bufio/bufio.go:299
		// _ = "end of CoverTab[1901]"
//line /usr/local/go/src/bufio/bufio.go:299
	}() && func() bool {
//line /usr/local/go/src/bufio/bufio.go:299
		_go_fuzz_dep_.CoverTab[1902]++
//line /usr/local/go/src/bufio/bufio.go:299
		return b.w-b.r < len(b.buf)
//line /usr/local/go/src/bufio/bufio.go:299
		// _ = "end of CoverTab[1902]"
//line /usr/local/go/src/bufio/bufio.go:299
	}() {
//line /usr/local/go/src/bufio/bufio.go:299
		_go_fuzz_dep_.CoverTab[1903]++
							b.fill()
//line /usr/local/go/src/bufio/bufio.go:300
		// _ = "end of CoverTab[1903]"
	}
//line /usr/local/go/src/bufio/bufio.go:301
	// _ = "end of CoverTab[1896]"
//line /usr/local/go/src/bufio/bufio.go:301
	_go_fuzz_dep_.CoverTab[1897]++
						b.lastRuneSize = -1
						if b.r == b.w {
//line /usr/local/go/src/bufio/bufio.go:303
		_go_fuzz_dep_.CoverTab[1904]++
							return 0, 0, b.readErr()
//line /usr/local/go/src/bufio/bufio.go:304
		// _ = "end of CoverTab[1904]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:305
		_go_fuzz_dep_.CoverTab[1905]++
//line /usr/local/go/src/bufio/bufio.go:305
		// _ = "end of CoverTab[1905]"
//line /usr/local/go/src/bufio/bufio.go:305
	}
//line /usr/local/go/src/bufio/bufio.go:305
	// _ = "end of CoverTab[1897]"
//line /usr/local/go/src/bufio/bufio.go:305
	_go_fuzz_dep_.CoverTab[1898]++
						r, size = rune(b.buf[b.r]), 1
						if r >= utf8.RuneSelf {
//line /usr/local/go/src/bufio/bufio.go:307
		_go_fuzz_dep_.CoverTab[1906]++
							r, size = utf8.DecodeRune(b.buf[b.r:b.w])
//line /usr/local/go/src/bufio/bufio.go:308
		// _ = "end of CoverTab[1906]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:309
		_go_fuzz_dep_.CoverTab[1907]++
//line /usr/local/go/src/bufio/bufio.go:309
		// _ = "end of CoverTab[1907]"
//line /usr/local/go/src/bufio/bufio.go:309
	}
//line /usr/local/go/src/bufio/bufio.go:309
	// _ = "end of CoverTab[1898]"
//line /usr/local/go/src/bufio/bufio.go:309
	_go_fuzz_dep_.CoverTab[1899]++
						b.r += size
						b.lastByte = int(b.buf[b.r-1])
						b.lastRuneSize = size
						return r, size, nil
//line /usr/local/go/src/bufio/bufio.go:313
	// _ = "end of CoverTab[1899]"
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
	_go_fuzz_dep_.CoverTab[1908]++
						if b.lastRuneSize < 0 || func() bool {
//line /usr/local/go/src/bufio/bufio.go:321
		_go_fuzz_dep_.CoverTab[1910]++
//line /usr/local/go/src/bufio/bufio.go:321
		return b.r < b.lastRuneSize
//line /usr/local/go/src/bufio/bufio.go:321
		// _ = "end of CoverTab[1910]"
//line /usr/local/go/src/bufio/bufio.go:321
	}() {
//line /usr/local/go/src/bufio/bufio.go:321
		_go_fuzz_dep_.CoverTab[1911]++
							return ErrInvalidUnreadRune
//line /usr/local/go/src/bufio/bufio.go:322
		// _ = "end of CoverTab[1911]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:323
		_go_fuzz_dep_.CoverTab[1912]++
//line /usr/local/go/src/bufio/bufio.go:323
		// _ = "end of CoverTab[1912]"
//line /usr/local/go/src/bufio/bufio.go:323
	}
//line /usr/local/go/src/bufio/bufio.go:323
	// _ = "end of CoverTab[1908]"
//line /usr/local/go/src/bufio/bufio.go:323
	_go_fuzz_dep_.CoverTab[1909]++
						b.r -= b.lastRuneSize
						b.lastByte = -1
						b.lastRuneSize = -1
						return nil
//line /usr/local/go/src/bufio/bufio.go:327
	// _ = "end of CoverTab[1909]"
}

// Buffered returns the number of bytes that can be read from the current buffer.
func (b *Reader) Buffered() int {
//line /usr/local/go/src/bufio/bufio.go:331
	_go_fuzz_dep_.CoverTab[1913]++
//line /usr/local/go/src/bufio/bufio.go:331
	return b.w - b.r
//line /usr/local/go/src/bufio/bufio.go:331
	// _ = "end of CoverTab[1913]"
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
	_go_fuzz_dep_.CoverTab[1914]++
						s := 0
						for {
//line /usr/local/go/src/bufio/bufio.go:345
		_go_fuzz_dep_.CoverTab[1917]++

							if i := bytes.IndexByte(b.buf[b.r+s:b.w], delim); i >= 0 {
//line /usr/local/go/src/bufio/bufio.go:347
			_go_fuzz_dep_.CoverTab[1921]++
								i += s
								line = b.buf[b.r : b.r+i+1]
								b.r += i + 1
								break
//line /usr/local/go/src/bufio/bufio.go:351
			// _ = "end of CoverTab[1921]"
		} else {
//line /usr/local/go/src/bufio/bufio.go:352
			_go_fuzz_dep_.CoverTab[1922]++
//line /usr/local/go/src/bufio/bufio.go:352
			// _ = "end of CoverTab[1922]"
//line /usr/local/go/src/bufio/bufio.go:352
		}
//line /usr/local/go/src/bufio/bufio.go:352
		// _ = "end of CoverTab[1917]"
//line /usr/local/go/src/bufio/bufio.go:352
		_go_fuzz_dep_.CoverTab[1918]++

//line /usr/local/go/src/bufio/bufio.go:355
		if b.err != nil {
//line /usr/local/go/src/bufio/bufio.go:355
			_go_fuzz_dep_.CoverTab[1923]++
								line = b.buf[b.r:b.w]
								b.r = b.w
								err = b.readErr()
								break
//line /usr/local/go/src/bufio/bufio.go:359
			// _ = "end of CoverTab[1923]"
		} else {
//line /usr/local/go/src/bufio/bufio.go:360
			_go_fuzz_dep_.CoverTab[1924]++
//line /usr/local/go/src/bufio/bufio.go:360
			// _ = "end of CoverTab[1924]"
//line /usr/local/go/src/bufio/bufio.go:360
		}
//line /usr/local/go/src/bufio/bufio.go:360
		// _ = "end of CoverTab[1918]"
//line /usr/local/go/src/bufio/bufio.go:360
		_go_fuzz_dep_.CoverTab[1919]++

//line /usr/local/go/src/bufio/bufio.go:363
		if b.Buffered() >= len(b.buf) {
//line /usr/local/go/src/bufio/bufio.go:363
			_go_fuzz_dep_.CoverTab[1925]++
								b.r = b.w
								line = b.buf
								err = ErrBufferFull
								break
//line /usr/local/go/src/bufio/bufio.go:367
			// _ = "end of CoverTab[1925]"
		} else {
//line /usr/local/go/src/bufio/bufio.go:368
			_go_fuzz_dep_.CoverTab[1926]++
//line /usr/local/go/src/bufio/bufio.go:368
			// _ = "end of CoverTab[1926]"
//line /usr/local/go/src/bufio/bufio.go:368
		}
//line /usr/local/go/src/bufio/bufio.go:368
		// _ = "end of CoverTab[1919]"
//line /usr/local/go/src/bufio/bufio.go:368
		_go_fuzz_dep_.CoverTab[1920]++

							s = b.w - b.r

							b.fill()
//line /usr/local/go/src/bufio/bufio.go:372
		// _ = "end of CoverTab[1920]"
	}
//line /usr/local/go/src/bufio/bufio.go:373
	// _ = "end of CoverTab[1914]"
//line /usr/local/go/src/bufio/bufio.go:373
	_go_fuzz_dep_.CoverTab[1915]++

//line /usr/local/go/src/bufio/bufio.go:376
	if i := len(line) - 1; i >= 0 {
//line /usr/local/go/src/bufio/bufio.go:376
		_go_fuzz_dep_.CoverTab[1927]++
							b.lastByte = int(line[i])
							b.lastRuneSize = -1
//line /usr/local/go/src/bufio/bufio.go:378
		// _ = "end of CoverTab[1927]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:379
		_go_fuzz_dep_.CoverTab[1928]++
//line /usr/local/go/src/bufio/bufio.go:379
		// _ = "end of CoverTab[1928]"
//line /usr/local/go/src/bufio/bufio.go:379
	}
//line /usr/local/go/src/bufio/bufio.go:379
	// _ = "end of CoverTab[1915]"
//line /usr/local/go/src/bufio/bufio.go:379
	_go_fuzz_dep_.CoverTab[1916]++

						return
//line /usr/local/go/src/bufio/bufio.go:381
	// _ = "end of CoverTab[1916]"
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
	_go_fuzz_dep_.CoverTab[1929]++
						line, err = b.ReadSlice('\n')
						if err == ErrBufferFull {
//line /usr/local/go/src/bufio/bufio.go:402
		_go_fuzz_dep_.CoverTab[1933]++

							if len(line) > 0 && func() bool {
//line /usr/local/go/src/bufio/bufio.go:404
			_go_fuzz_dep_.CoverTab[1935]++
//line /usr/local/go/src/bufio/bufio.go:404
			return line[len(line)-1] == '\r'
//line /usr/local/go/src/bufio/bufio.go:404
			// _ = "end of CoverTab[1935]"
//line /usr/local/go/src/bufio/bufio.go:404
		}() {
//line /usr/local/go/src/bufio/bufio.go:404
			_go_fuzz_dep_.CoverTab[1936]++

//line /usr/local/go/src/bufio/bufio.go:407
			if b.r == 0 {
//line /usr/local/go/src/bufio/bufio.go:407
				_go_fuzz_dep_.CoverTab[1938]++

									panic("bufio: tried to rewind past start of buffer")
//line /usr/local/go/src/bufio/bufio.go:409
				// _ = "end of CoverTab[1938]"
			} else {
//line /usr/local/go/src/bufio/bufio.go:410
				_go_fuzz_dep_.CoverTab[1939]++
//line /usr/local/go/src/bufio/bufio.go:410
				// _ = "end of CoverTab[1939]"
//line /usr/local/go/src/bufio/bufio.go:410
			}
//line /usr/local/go/src/bufio/bufio.go:410
			// _ = "end of CoverTab[1936]"
//line /usr/local/go/src/bufio/bufio.go:410
			_go_fuzz_dep_.CoverTab[1937]++
								b.r--
								line = line[:len(line)-1]
//line /usr/local/go/src/bufio/bufio.go:412
			// _ = "end of CoverTab[1937]"
		} else {
//line /usr/local/go/src/bufio/bufio.go:413
			_go_fuzz_dep_.CoverTab[1940]++
//line /usr/local/go/src/bufio/bufio.go:413
			// _ = "end of CoverTab[1940]"
//line /usr/local/go/src/bufio/bufio.go:413
		}
//line /usr/local/go/src/bufio/bufio.go:413
		// _ = "end of CoverTab[1933]"
//line /usr/local/go/src/bufio/bufio.go:413
		_go_fuzz_dep_.CoverTab[1934]++
							return line, true, nil
//line /usr/local/go/src/bufio/bufio.go:414
		// _ = "end of CoverTab[1934]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:415
		_go_fuzz_dep_.CoverTab[1941]++
//line /usr/local/go/src/bufio/bufio.go:415
		// _ = "end of CoverTab[1941]"
//line /usr/local/go/src/bufio/bufio.go:415
	}
//line /usr/local/go/src/bufio/bufio.go:415
	// _ = "end of CoverTab[1929]"
//line /usr/local/go/src/bufio/bufio.go:415
	_go_fuzz_dep_.CoverTab[1930]++

						if len(line) == 0 {
//line /usr/local/go/src/bufio/bufio.go:417
		_go_fuzz_dep_.CoverTab[1942]++
							if err != nil {
//line /usr/local/go/src/bufio/bufio.go:418
			_go_fuzz_dep_.CoverTab[1944]++
								line = nil
//line /usr/local/go/src/bufio/bufio.go:419
			// _ = "end of CoverTab[1944]"
		} else {
//line /usr/local/go/src/bufio/bufio.go:420
			_go_fuzz_dep_.CoverTab[1945]++
//line /usr/local/go/src/bufio/bufio.go:420
			// _ = "end of CoverTab[1945]"
//line /usr/local/go/src/bufio/bufio.go:420
		}
//line /usr/local/go/src/bufio/bufio.go:420
		// _ = "end of CoverTab[1942]"
//line /usr/local/go/src/bufio/bufio.go:420
		_go_fuzz_dep_.CoverTab[1943]++
							return
//line /usr/local/go/src/bufio/bufio.go:421
		// _ = "end of CoverTab[1943]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:422
		_go_fuzz_dep_.CoverTab[1946]++
//line /usr/local/go/src/bufio/bufio.go:422
		// _ = "end of CoverTab[1946]"
//line /usr/local/go/src/bufio/bufio.go:422
	}
//line /usr/local/go/src/bufio/bufio.go:422
	// _ = "end of CoverTab[1930]"
//line /usr/local/go/src/bufio/bufio.go:422
	_go_fuzz_dep_.CoverTab[1931]++
						err = nil

						if line[len(line)-1] == '\n' {
//line /usr/local/go/src/bufio/bufio.go:425
		_go_fuzz_dep_.CoverTab[1947]++
							drop := 1
							if len(line) > 1 && func() bool {
//line /usr/local/go/src/bufio/bufio.go:427
			_go_fuzz_dep_.CoverTab[1949]++
//line /usr/local/go/src/bufio/bufio.go:427
			return line[len(line)-2] == '\r'
//line /usr/local/go/src/bufio/bufio.go:427
			// _ = "end of CoverTab[1949]"
//line /usr/local/go/src/bufio/bufio.go:427
		}() {
//line /usr/local/go/src/bufio/bufio.go:427
			_go_fuzz_dep_.CoverTab[1950]++
								drop = 2
//line /usr/local/go/src/bufio/bufio.go:428
			// _ = "end of CoverTab[1950]"
		} else {
//line /usr/local/go/src/bufio/bufio.go:429
			_go_fuzz_dep_.CoverTab[1951]++
//line /usr/local/go/src/bufio/bufio.go:429
			// _ = "end of CoverTab[1951]"
//line /usr/local/go/src/bufio/bufio.go:429
		}
//line /usr/local/go/src/bufio/bufio.go:429
		// _ = "end of CoverTab[1947]"
//line /usr/local/go/src/bufio/bufio.go:429
		_go_fuzz_dep_.CoverTab[1948]++
							line = line[:len(line)-drop]
//line /usr/local/go/src/bufio/bufio.go:430
		// _ = "end of CoverTab[1948]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:431
		_go_fuzz_dep_.CoverTab[1952]++
//line /usr/local/go/src/bufio/bufio.go:431
		// _ = "end of CoverTab[1952]"
//line /usr/local/go/src/bufio/bufio.go:431
	}
//line /usr/local/go/src/bufio/bufio.go:431
	// _ = "end of CoverTab[1931]"
//line /usr/local/go/src/bufio/bufio.go:431
	_go_fuzz_dep_.CoverTab[1932]++
						return
//line /usr/local/go/src/bufio/bufio.go:432
	// _ = "end of CoverTab[1932]"
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
	_go_fuzz_dep_.CoverTab[1953]++
						var frag []byte

						for {
//line /usr/local/go/src/bufio/bufio.go:445
		_go_fuzz_dep_.CoverTab[1955]++
							var e error
							frag, e = b.ReadSlice(delim)
							if e == nil {
//line /usr/local/go/src/bufio/bufio.go:448
			_go_fuzz_dep_.CoverTab[1958]++
								break
//line /usr/local/go/src/bufio/bufio.go:449
			// _ = "end of CoverTab[1958]"
		} else {
//line /usr/local/go/src/bufio/bufio.go:450
			_go_fuzz_dep_.CoverTab[1959]++
//line /usr/local/go/src/bufio/bufio.go:450
			// _ = "end of CoverTab[1959]"
//line /usr/local/go/src/bufio/bufio.go:450
		}
//line /usr/local/go/src/bufio/bufio.go:450
		// _ = "end of CoverTab[1955]"
//line /usr/local/go/src/bufio/bufio.go:450
		_go_fuzz_dep_.CoverTab[1956]++
							if e != ErrBufferFull {
//line /usr/local/go/src/bufio/bufio.go:451
			_go_fuzz_dep_.CoverTab[1960]++
								err = e
								break
//line /usr/local/go/src/bufio/bufio.go:453
			// _ = "end of CoverTab[1960]"
		} else {
//line /usr/local/go/src/bufio/bufio.go:454
			_go_fuzz_dep_.CoverTab[1961]++
//line /usr/local/go/src/bufio/bufio.go:454
			// _ = "end of CoverTab[1961]"
//line /usr/local/go/src/bufio/bufio.go:454
		}
//line /usr/local/go/src/bufio/bufio.go:454
		// _ = "end of CoverTab[1956]"
//line /usr/local/go/src/bufio/bufio.go:454
		_go_fuzz_dep_.CoverTab[1957]++

//line /usr/local/go/src/bufio/bufio.go:457
		buf := bytes.Clone(frag)
							fullBuffers = append(fullBuffers, buf)
							totalLen += len(buf)
//line /usr/local/go/src/bufio/bufio.go:459
		// _ = "end of CoverTab[1957]"
	}
//line /usr/local/go/src/bufio/bufio.go:460
	// _ = "end of CoverTab[1953]"
//line /usr/local/go/src/bufio/bufio.go:460
	_go_fuzz_dep_.CoverTab[1954]++

						totalLen += len(frag)
						return fullBuffers, frag, totalLen, err
//line /usr/local/go/src/bufio/bufio.go:463
	// _ = "end of CoverTab[1954]"
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
	_go_fuzz_dep_.CoverTab[1962]++
						full, frag, n, err := b.collectFragments(delim)

						buf := make([]byte, n)
						n = 0

						for i := range full {
//line /usr/local/go/src/bufio/bufio.go:479
		_go_fuzz_dep_.CoverTab[1964]++
							n += copy(buf[n:], full[i])
//line /usr/local/go/src/bufio/bufio.go:480
		// _ = "end of CoverTab[1964]"
	}
//line /usr/local/go/src/bufio/bufio.go:481
	// _ = "end of CoverTab[1962]"
//line /usr/local/go/src/bufio/bufio.go:481
	_go_fuzz_dep_.CoverTab[1963]++
						copy(buf[n:], frag)
						return buf, err
//line /usr/local/go/src/bufio/bufio.go:483
	// _ = "end of CoverTab[1963]"
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
	_go_fuzz_dep_.CoverTab[1965]++
						full, frag, n, err := b.collectFragments(delim)
	// Allocate new buffer to hold the full pieces and the fragment.
	var buf strings.Builder
	buf.Grow(n)

	for _, fb := range full {
//line /usr/local/go/src/bufio/bufio.go:499
		_go_fuzz_dep_.CoverTab[1967]++
							buf.Write(fb)
//line /usr/local/go/src/bufio/bufio.go:500
		// _ = "end of CoverTab[1967]"
	}
//line /usr/local/go/src/bufio/bufio.go:501
	// _ = "end of CoverTab[1965]"
//line /usr/local/go/src/bufio/bufio.go:501
	_go_fuzz_dep_.CoverTab[1966]++
						buf.Write(frag)
						return buf.String(), err
//line /usr/local/go/src/bufio/bufio.go:503
	// _ = "end of CoverTab[1966]"
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
	_go_fuzz_dep_.CoverTab[1968]++
						b.lastByte = -1
						b.lastRuneSize = -1

						n, err = b.writeBuf(w)
						if err != nil {
//line /usr/local/go/src/bufio/bufio.go:515
		_go_fuzz_dep_.CoverTab[1975]++
							return
//line /usr/local/go/src/bufio/bufio.go:516
		// _ = "end of CoverTab[1975]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:517
		_go_fuzz_dep_.CoverTab[1976]++
//line /usr/local/go/src/bufio/bufio.go:517
		// _ = "end of CoverTab[1976]"
//line /usr/local/go/src/bufio/bufio.go:517
	}
//line /usr/local/go/src/bufio/bufio.go:517
	// _ = "end of CoverTab[1968]"
//line /usr/local/go/src/bufio/bufio.go:517
	_go_fuzz_dep_.CoverTab[1969]++

						if r, ok := b.rd.(io.WriterTo); ok {
//line /usr/local/go/src/bufio/bufio.go:519
		_go_fuzz_dep_.CoverTab[1977]++
							m, err := r.WriteTo(w)
							n += m
							return n, err
//line /usr/local/go/src/bufio/bufio.go:522
		// _ = "end of CoverTab[1977]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:523
		_go_fuzz_dep_.CoverTab[1978]++
//line /usr/local/go/src/bufio/bufio.go:523
		// _ = "end of CoverTab[1978]"
//line /usr/local/go/src/bufio/bufio.go:523
	}
//line /usr/local/go/src/bufio/bufio.go:523
	// _ = "end of CoverTab[1969]"
//line /usr/local/go/src/bufio/bufio.go:523
	_go_fuzz_dep_.CoverTab[1970]++

						if w, ok := w.(io.ReaderFrom); ok {
//line /usr/local/go/src/bufio/bufio.go:525
		_go_fuzz_dep_.CoverTab[1979]++
							m, err := w.ReadFrom(b.rd)
							n += m
							return n, err
//line /usr/local/go/src/bufio/bufio.go:528
		// _ = "end of CoverTab[1979]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:529
		_go_fuzz_dep_.CoverTab[1980]++
//line /usr/local/go/src/bufio/bufio.go:529
		// _ = "end of CoverTab[1980]"
//line /usr/local/go/src/bufio/bufio.go:529
	}
//line /usr/local/go/src/bufio/bufio.go:529
	// _ = "end of CoverTab[1970]"
//line /usr/local/go/src/bufio/bufio.go:529
	_go_fuzz_dep_.CoverTab[1971]++

						if b.w-b.r < len(b.buf) {
//line /usr/local/go/src/bufio/bufio.go:531
		_go_fuzz_dep_.CoverTab[1981]++
							b.fill()
//line /usr/local/go/src/bufio/bufio.go:532
		// _ = "end of CoverTab[1981]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:533
		_go_fuzz_dep_.CoverTab[1982]++
//line /usr/local/go/src/bufio/bufio.go:533
		// _ = "end of CoverTab[1982]"
//line /usr/local/go/src/bufio/bufio.go:533
	}
//line /usr/local/go/src/bufio/bufio.go:533
	// _ = "end of CoverTab[1971]"
//line /usr/local/go/src/bufio/bufio.go:533
	_go_fuzz_dep_.CoverTab[1972]++

						for b.r < b.w {
//line /usr/local/go/src/bufio/bufio.go:535
		_go_fuzz_dep_.CoverTab[1983]++

							m, err := b.writeBuf(w)
							n += m
							if err != nil {
//line /usr/local/go/src/bufio/bufio.go:539
			_go_fuzz_dep_.CoverTab[1985]++
								return n, err
//line /usr/local/go/src/bufio/bufio.go:540
			// _ = "end of CoverTab[1985]"
		} else {
//line /usr/local/go/src/bufio/bufio.go:541
			_go_fuzz_dep_.CoverTab[1986]++
//line /usr/local/go/src/bufio/bufio.go:541
			// _ = "end of CoverTab[1986]"
//line /usr/local/go/src/bufio/bufio.go:541
		}
//line /usr/local/go/src/bufio/bufio.go:541
		// _ = "end of CoverTab[1983]"
//line /usr/local/go/src/bufio/bufio.go:541
		_go_fuzz_dep_.CoverTab[1984]++
							b.fill()
//line /usr/local/go/src/bufio/bufio.go:542
		// _ = "end of CoverTab[1984]"
	}
//line /usr/local/go/src/bufio/bufio.go:543
	// _ = "end of CoverTab[1972]"
//line /usr/local/go/src/bufio/bufio.go:543
	_go_fuzz_dep_.CoverTab[1973]++

						if b.err == io.EOF {
//line /usr/local/go/src/bufio/bufio.go:545
		_go_fuzz_dep_.CoverTab[1987]++
							b.err = nil
//line /usr/local/go/src/bufio/bufio.go:546
		// _ = "end of CoverTab[1987]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:547
		_go_fuzz_dep_.CoverTab[1988]++
//line /usr/local/go/src/bufio/bufio.go:547
		// _ = "end of CoverTab[1988]"
//line /usr/local/go/src/bufio/bufio.go:547
	}
//line /usr/local/go/src/bufio/bufio.go:547
	// _ = "end of CoverTab[1973]"
//line /usr/local/go/src/bufio/bufio.go:547
	_go_fuzz_dep_.CoverTab[1974]++

						return n, b.readErr()
//line /usr/local/go/src/bufio/bufio.go:549
	// _ = "end of CoverTab[1974]"
}

var errNegativeWrite = errors.New("bufio: writer returned negative count from Write")

// writeBuf writes the Reader's buffer to the writer.
func (b *Reader) writeBuf(w io.Writer) (int64, error) {
//line /usr/local/go/src/bufio/bufio.go:555
	_go_fuzz_dep_.CoverTab[1989]++
						n, err := w.Write(b.buf[b.r:b.w])
						if n < 0 {
//line /usr/local/go/src/bufio/bufio.go:557
		_go_fuzz_dep_.CoverTab[1991]++
							panic(errNegativeWrite)
//line /usr/local/go/src/bufio/bufio.go:558
		// _ = "end of CoverTab[1991]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:559
		_go_fuzz_dep_.CoverTab[1992]++
//line /usr/local/go/src/bufio/bufio.go:559
		// _ = "end of CoverTab[1992]"
//line /usr/local/go/src/bufio/bufio.go:559
	}
//line /usr/local/go/src/bufio/bufio.go:559
	// _ = "end of CoverTab[1989]"
//line /usr/local/go/src/bufio/bufio.go:559
	_go_fuzz_dep_.CoverTab[1990]++
						b.r += n
						return int64(n), err
//line /usr/local/go/src/bufio/bufio.go:561
	// _ = "end of CoverTab[1990]"
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
	_go_fuzz_dep_.CoverTab[1993]++

						b, ok := w.(*Writer)
						if ok && func() bool {
//line /usr/local/go/src/bufio/bufio.go:585
		_go_fuzz_dep_.CoverTab[1996]++
//line /usr/local/go/src/bufio/bufio.go:585
		return len(b.buf) >= size
//line /usr/local/go/src/bufio/bufio.go:585
		// _ = "end of CoverTab[1996]"
//line /usr/local/go/src/bufio/bufio.go:585
	}() {
//line /usr/local/go/src/bufio/bufio.go:585
		_go_fuzz_dep_.CoverTab[1997]++
							return b
//line /usr/local/go/src/bufio/bufio.go:586
		// _ = "end of CoverTab[1997]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:587
		_go_fuzz_dep_.CoverTab[1998]++
//line /usr/local/go/src/bufio/bufio.go:587
		// _ = "end of CoverTab[1998]"
//line /usr/local/go/src/bufio/bufio.go:587
	}
//line /usr/local/go/src/bufio/bufio.go:587
	// _ = "end of CoverTab[1993]"
//line /usr/local/go/src/bufio/bufio.go:587
	_go_fuzz_dep_.CoverTab[1994]++
						if size <= 0 {
//line /usr/local/go/src/bufio/bufio.go:588
		_go_fuzz_dep_.CoverTab[1999]++
							size = defaultBufSize
//line /usr/local/go/src/bufio/bufio.go:589
		// _ = "end of CoverTab[1999]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:590
		_go_fuzz_dep_.CoverTab[2000]++
//line /usr/local/go/src/bufio/bufio.go:590
		// _ = "end of CoverTab[2000]"
//line /usr/local/go/src/bufio/bufio.go:590
	}
//line /usr/local/go/src/bufio/bufio.go:590
	// _ = "end of CoverTab[1994]"
//line /usr/local/go/src/bufio/bufio.go:590
	_go_fuzz_dep_.CoverTab[1995]++
						return &Writer{
		buf:	make([]byte, size),
		wr:	w,
	}
//line /usr/local/go/src/bufio/bufio.go:594
	// _ = "end of CoverTab[1995]"
}

// NewWriter returns a new Writer whose buffer has the default size.
//line /usr/local/go/src/bufio/bufio.go:597
// If the argument io.Writer is already a Writer with large enough buffer size,
//line /usr/local/go/src/bufio/bufio.go:597
// it returns the underlying Writer.
//line /usr/local/go/src/bufio/bufio.go:600
func NewWriter(w io.Writer) *Writer {
//line /usr/local/go/src/bufio/bufio.go:600
	_go_fuzz_dep_.CoverTab[2001]++
						return NewWriterSize(w, defaultBufSize)
//line /usr/local/go/src/bufio/bufio.go:601
	// _ = "end of CoverTab[2001]"
}

// Size returns the size of the underlying buffer in bytes.
func (b *Writer) Size() int {
//line /usr/local/go/src/bufio/bufio.go:605
	_go_fuzz_dep_.CoverTab[2002]++
//line /usr/local/go/src/bufio/bufio.go:605
	return len(b.buf)
//line /usr/local/go/src/bufio/bufio.go:605
	// _ = "end of CoverTab[2002]"
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
	_go_fuzz_dep_.CoverTab[2003]++
						if b.buf == nil {
//line /usr/local/go/src/bufio/bufio.go:612
		_go_fuzz_dep_.CoverTab[2005]++
							b.buf = make([]byte, defaultBufSize)
//line /usr/local/go/src/bufio/bufio.go:613
		// _ = "end of CoverTab[2005]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:614
		_go_fuzz_dep_.CoverTab[2006]++
//line /usr/local/go/src/bufio/bufio.go:614
		// _ = "end of CoverTab[2006]"
//line /usr/local/go/src/bufio/bufio.go:614
	}
//line /usr/local/go/src/bufio/bufio.go:614
	// _ = "end of CoverTab[2003]"
//line /usr/local/go/src/bufio/bufio.go:614
	_go_fuzz_dep_.CoverTab[2004]++
						b.err = nil
						b.n = 0
						b.wr = w
//line /usr/local/go/src/bufio/bufio.go:617
	// _ = "end of CoverTab[2004]"
}

// Flush writes any buffered data to the underlying io.Writer.
func (b *Writer) Flush() error {
//line /usr/local/go/src/bufio/bufio.go:621
	_go_fuzz_dep_.CoverTab[2007]++
						if b.err != nil {
//line /usr/local/go/src/bufio/bufio.go:622
		_go_fuzz_dep_.CoverTab[2012]++
							return b.err
//line /usr/local/go/src/bufio/bufio.go:623
		// _ = "end of CoverTab[2012]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:624
		_go_fuzz_dep_.CoverTab[2013]++
//line /usr/local/go/src/bufio/bufio.go:624
		// _ = "end of CoverTab[2013]"
//line /usr/local/go/src/bufio/bufio.go:624
	}
//line /usr/local/go/src/bufio/bufio.go:624
	// _ = "end of CoverTab[2007]"
//line /usr/local/go/src/bufio/bufio.go:624
	_go_fuzz_dep_.CoverTab[2008]++
						if b.n == 0 {
//line /usr/local/go/src/bufio/bufio.go:625
		_go_fuzz_dep_.CoverTab[2014]++
							return nil
//line /usr/local/go/src/bufio/bufio.go:626
		// _ = "end of CoverTab[2014]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:627
		_go_fuzz_dep_.CoverTab[2015]++
//line /usr/local/go/src/bufio/bufio.go:627
		// _ = "end of CoverTab[2015]"
//line /usr/local/go/src/bufio/bufio.go:627
	}
//line /usr/local/go/src/bufio/bufio.go:627
	// _ = "end of CoverTab[2008]"
//line /usr/local/go/src/bufio/bufio.go:627
	_go_fuzz_dep_.CoverTab[2009]++
						n, err := b.wr.Write(b.buf[0:b.n])
						if n < b.n && func() bool {
//line /usr/local/go/src/bufio/bufio.go:629
		_go_fuzz_dep_.CoverTab[2016]++
//line /usr/local/go/src/bufio/bufio.go:629
		return err == nil
//line /usr/local/go/src/bufio/bufio.go:629
		// _ = "end of CoverTab[2016]"
//line /usr/local/go/src/bufio/bufio.go:629
	}() {
//line /usr/local/go/src/bufio/bufio.go:629
		_go_fuzz_dep_.CoverTab[2017]++
							err = io.ErrShortWrite
//line /usr/local/go/src/bufio/bufio.go:630
		// _ = "end of CoverTab[2017]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:631
		_go_fuzz_dep_.CoverTab[2018]++
//line /usr/local/go/src/bufio/bufio.go:631
		// _ = "end of CoverTab[2018]"
//line /usr/local/go/src/bufio/bufio.go:631
	}
//line /usr/local/go/src/bufio/bufio.go:631
	// _ = "end of CoverTab[2009]"
//line /usr/local/go/src/bufio/bufio.go:631
	_go_fuzz_dep_.CoverTab[2010]++
						if err != nil {
//line /usr/local/go/src/bufio/bufio.go:632
		_go_fuzz_dep_.CoverTab[2019]++
							if n > 0 && func() bool {
//line /usr/local/go/src/bufio/bufio.go:633
			_go_fuzz_dep_.CoverTab[2021]++
//line /usr/local/go/src/bufio/bufio.go:633
			return n < b.n
//line /usr/local/go/src/bufio/bufio.go:633
			// _ = "end of CoverTab[2021]"
//line /usr/local/go/src/bufio/bufio.go:633
		}() {
//line /usr/local/go/src/bufio/bufio.go:633
			_go_fuzz_dep_.CoverTab[2022]++
								copy(b.buf[0:b.n-n], b.buf[n:b.n])
//line /usr/local/go/src/bufio/bufio.go:634
			// _ = "end of CoverTab[2022]"
		} else {
//line /usr/local/go/src/bufio/bufio.go:635
			_go_fuzz_dep_.CoverTab[2023]++
//line /usr/local/go/src/bufio/bufio.go:635
			// _ = "end of CoverTab[2023]"
//line /usr/local/go/src/bufio/bufio.go:635
		}
//line /usr/local/go/src/bufio/bufio.go:635
		// _ = "end of CoverTab[2019]"
//line /usr/local/go/src/bufio/bufio.go:635
		_go_fuzz_dep_.CoverTab[2020]++
							b.n -= n
							b.err = err
							return err
//line /usr/local/go/src/bufio/bufio.go:638
		// _ = "end of CoverTab[2020]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:639
		_go_fuzz_dep_.CoverTab[2024]++
//line /usr/local/go/src/bufio/bufio.go:639
		// _ = "end of CoverTab[2024]"
//line /usr/local/go/src/bufio/bufio.go:639
	}
//line /usr/local/go/src/bufio/bufio.go:639
	// _ = "end of CoverTab[2010]"
//line /usr/local/go/src/bufio/bufio.go:639
	_go_fuzz_dep_.CoverTab[2011]++
						b.n = 0
						return nil
//line /usr/local/go/src/bufio/bufio.go:641
	// _ = "end of CoverTab[2011]"
}

// Available returns how many bytes are unused in the buffer.
func (b *Writer) Available() int {
//line /usr/local/go/src/bufio/bufio.go:645
	_go_fuzz_dep_.CoverTab[2025]++
//line /usr/local/go/src/bufio/bufio.go:645
	return len(b.buf) - b.n
//line /usr/local/go/src/bufio/bufio.go:645
	// _ = "end of CoverTab[2025]"
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
	_go_fuzz_dep_.CoverTab[2026]++
						return b.buf[b.n:][:0]
//line /usr/local/go/src/bufio/bufio.go:652
	// _ = "end of CoverTab[2026]"
}

// Buffered returns the number of bytes that have been written into the current buffer.
func (b *Writer) Buffered() int {
//line /usr/local/go/src/bufio/bufio.go:656
	_go_fuzz_dep_.CoverTab[2027]++
//line /usr/local/go/src/bufio/bufio.go:656
	return b.n
//line /usr/local/go/src/bufio/bufio.go:656
	// _ = "end of CoverTab[2027]"
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
	_go_fuzz_dep_.CoverTab[2028]++
						for len(p) > b.Available() && func() bool {
//line /usr/local/go/src/bufio/bufio.go:663
		_go_fuzz_dep_.CoverTab[2031]++
//line /usr/local/go/src/bufio/bufio.go:663
		return b.err == nil
//line /usr/local/go/src/bufio/bufio.go:663
		// _ = "end of CoverTab[2031]"
//line /usr/local/go/src/bufio/bufio.go:663
	}() {
//line /usr/local/go/src/bufio/bufio.go:663
		_go_fuzz_dep_.CoverTab[2032]++
							var n int
							if b.Buffered() == 0 {
//line /usr/local/go/src/bufio/bufio.go:665
			_go_fuzz_dep_.CoverTab[2034]++

//line /usr/local/go/src/bufio/bufio.go:668
			n, b.err = b.wr.Write(p)
//line /usr/local/go/src/bufio/bufio.go:668
			// _ = "end of CoverTab[2034]"
		} else {
//line /usr/local/go/src/bufio/bufio.go:669
			_go_fuzz_dep_.CoverTab[2035]++
								n = copy(b.buf[b.n:], p)
								b.n += n
								b.Flush()
//line /usr/local/go/src/bufio/bufio.go:672
			// _ = "end of CoverTab[2035]"
		}
//line /usr/local/go/src/bufio/bufio.go:673
		// _ = "end of CoverTab[2032]"
//line /usr/local/go/src/bufio/bufio.go:673
		_go_fuzz_dep_.CoverTab[2033]++
							nn += n
							p = p[n:]
//line /usr/local/go/src/bufio/bufio.go:675
		// _ = "end of CoverTab[2033]"
	}
//line /usr/local/go/src/bufio/bufio.go:676
	// _ = "end of CoverTab[2028]"
//line /usr/local/go/src/bufio/bufio.go:676
	_go_fuzz_dep_.CoverTab[2029]++
						if b.err != nil {
//line /usr/local/go/src/bufio/bufio.go:677
		_go_fuzz_dep_.CoverTab[2036]++
							return nn, b.err
//line /usr/local/go/src/bufio/bufio.go:678
		// _ = "end of CoverTab[2036]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:679
		_go_fuzz_dep_.CoverTab[2037]++
//line /usr/local/go/src/bufio/bufio.go:679
		// _ = "end of CoverTab[2037]"
//line /usr/local/go/src/bufio/bufio.go:679
	}
//line /usr/local/go/src/bufio/bufio.go:679
	// _ = "end of CoverTab[2029]"
//line /usr/local/go/src/bufio/bufio.go:679
	_go_fuzz_dep_.CoverTab[2030]++
						n := copy(b.buf[b.n:], p)
						b.n += n
						nn += n
						return nn, nil
//line /usr/local/go/src/bufio/bufio.go:683
	// _ = "end of CoverTab[2030]"
}

// WriteByte writes a single byte.
func (b *Writer) WriteByte(c byte) error {
//line /usr/local/go/src/bufio/bufio.go:687
	_go_fuzz_dep_.CoverTab[2038]++
						if b.err != nil {
//line /usr/local/go/src/bufio/bufio.go:688
		_go_fuzz_dep_.CoverTab[2041]++
							return b.err
//line /usr/local/go/src/bufio/bufio.go:689
		// _ = "end of CoverTab[2041]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:690
		_go_fuzz_dep_.CoverTab[2042]++
//line /usr/local/go/src/bufio/bufio.go:690
		// _ = "end of CoverTab[2042]"
//line /usr/local/go/src/bufio/bufio.go:690
	}
//line /usr/local/go/src/bufio/bufio.go:690
	// _ = "end of CoverTab[2038]"
//line /usr/local/go/src/bufio/bufio.go:690
	_go_fuzz_dep_.CoverTab[2039]++
						if b.Available() <= 0 && func() bool {
//line /usr/local/go/src/bufio/bufio.go:691
		_go_fuzz_dep_.CoverTab[2043]++
//line /usr/local/go/src/bufio/bufio.go:691
		return b.Flush() != nil
//line /usr/local/go/src/bufio/bufio.go:691
		// _ = "end of CoverTab[2043]"
//line /usr/local/go/src/bufio/bufio.go:691
	}() {
//line /usr/local/go/src/bufio/bufio.go:691
		_go_fuzz_dep_.CoverTab[2044]++
							return b.err
//line /usr/local/go/src/bufio/bufio.go:692
		// _ = "end of CoverTab[2044]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:693
		_go_fuzz_dep_.CoverTab[2045]++
//line /usr/local/go/src/bufio/bufio.go:693
		// _ = "end of CoverTab[2045]"
//line /usr/local/go/src/bufio/bufio.go:693
	}
//line /usr/local/go/src/bufio/bufio.go:693
	// _ = "end of CoverTab[2039]"
//line /usr/local/go/src/bufio/bufio.go:693
	_go_fuzz_dep_.CoverTab[2040]++
						b.buf[b.n] = c
						b.n++
						return nil
//line /usr/local/go/src/bufio/bufio.go:696
	// _ = "end of CoverTab[2040]"
}

// WriteRune writes a single Unicode code point, returning
//line /usr/local/go/src/bufio/bufio.go:699
// the number of bytes written and any error.
//line /usr/local/go/src/bufio/bufio.go:701
func (b *Writer) WriteRune(r rune) (size int, err error) {
//line /usr/local/go/src/bufio/bufio.go:701
	_go_fuzz_dep_.CoverTab[2046]++

						if uint32(r) < utf8.RuneSelf {
//line /usr/local/go/src/bufio/bufio.go:703
		_go_fuzz_dep_.CoverTab[2050]++
							err = b.WriteByte(byte(r))
							if err != nil {
//line /usr/local/go/src/bufio/bufio.go:705
			_go_fuzz_dep_.CoverTab[2052]++
								return 0, err
//line /usr/local/go/src/bufio/bufio.go:706
			// _ = "end of CoverTab[2052]"
		} else {
//line /usr/local/go/src/bufio/bufio.go:707
			_go_fuzz_dep_.CoverTab[2053]++
//line /usr/local/go/src/bufio/bufio.go:707
			// _ = "end of CoverTab[2053]"
//line /usr/local/go/src/bufio/bufio.go:707
		}
//line /usr/local/go/src/bufio/bufio.go:707
		// _ = "end of CoverTab[2050]"
//line /usr/local/go/src/bufio/bufio.go:707
		_go_fuzz_dep_.CoverTab[2051]++
							return 1, nil
//line /usr/local/go/src/bufio/bufio.go:708
		// _ = "end of CoverTab[2051]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:709
		_go_fuzz_dep_.CoverTab[2054]++
//line /usr/local/go/src/bufio/bufio.go:709
		// _ = "end of CoverTab[2054]"
//line /usr/local/go/src/bufio/bufio.go:709
	}
//line /usr/local/go/src/bufio/bufio.go:709
	// _ = "end of CoverTab[2046]"
//line /usr/local/go/src/bufio/bufio.go:709
	_go_fuzz_dep_.CoverTab[2047]++
						if b.err != nil {
//line /usr/local/go/src/bufio/bufio.go:710
		_go_fuzz_dep_.CoverTab[2055]++
							return 0, b.err
//line /usr/local/go/src/bufio/bufio.go:711
		// _ = "end of CoverTab[2055]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:712
		_go_fuzz_dep_.CoverTab[2056]++
//line /usr/local/go/src/bufio/bufio.go:712
		// _ = "end of CoverTab[2056]"
//line /usr/local/go/src/bufio/bufio.go:712
	}
//line /usr/local/go/src/bufio/bufio.go:712
	// _ = "end of CoverTab[2047]"
//line /usr/local/go/src/bufio/bufio.go:712
	_go_fuzz_dep_.CoverTab[2048]++
						n := b.Available()
						if n < utf8.UTFMax {
//line /usr/local/go/src/bufio/bufio.go:714
		_go_fuzz_dep_.CoverTab[2057]++
							if b.Flush(); b.err != nil {
//line /usr/local/go/src/bufio/bufio.go:715
			_go_fuzz_dep_.CoverTab[2059]++
								return 0, b.err
//line /usr/local/go/src/bufio/bufio.go:716
			// _ = "end of CoverTab[2059]"
		} else {
//line /usr/local/go/src/bufio/bufio.go:717
			_go_fuzz_dep_.CoverTab[2060]++
//line /usr/local/go/src/bufio/bufio.go:717
			// _ = "end of CoverTab[2060]"
//line /usr/local/go/src/bufio/bufio.go:717
		}
//line /usr/local/go/src/bufio/bufio.go:717
		// _ = "end of CoverTab[2057]"
//line /usr/local/go/src/bufio/bufio.go:717
		_go_fuzz_dep_.CoverTab[2058]++
							n = b.Available()
							if n < utf8.UTFMax {
//line /usr/local/go/src/bufio/bufio.go:719
			_go_fuzz_dep_.CoverTab[2061]++

								return b.WriteString(string(r))
//line /usr/local/go/src/bufio/bufio.go:721
			// _ = "end of CoverTab[2061]"
		} else {
//line /usr/local/go/src/bufio/bufio.go:722
			_go_fuzz_dep_.CoverTab[2062]++
//line /usr/local/go/src/bufio/bufio.go:722
			// _ = "end of CoverTab[2062]"
//line /usr/local/go/src/bufio/bufio.go:722
		}
//line /usr/local/go/src/bufio/bufio.go:722
		// _ = "end of CoverTab[2058]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:723
		_go_fuzz_dep_.CoverTab[2063]++
//line /usr/local/go/src/bufio/bufio.go:723
		// _ = "end of CoverTab[2063]"
//line /usr/local/go/src/bufio/bufio.go:723
	}
//line /usr/local/go/src/bufio/bufio.go:723
	// _ = "end of CoverTab[2048]"
//line /usr/local/go/src/bufio/bufio.go:723
	_go_fuzz_dep_.CoverTab[2049]++
						size = utf8.EncodeRune(b.buf[b.n:], r)
						b.n += size
						return size, nil
//line /usr/local/go/src/bufio/bufio.go:726
	// _ = "end of CoverTab[2049]"
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
	_go_fuzz_dep_.CoverTab[2064]++
						var sw io.StringWriter
						tryStringWriter := true

						nn := 0
						for len(s) > b.Available() && func() bool {
//line /usr/local/go/src/bufio/bufio.go:738
		_go_fuzz_dep_.CoverTab[2067]++
//line /usr/local/go/src/bufio/bufio.go:738
		return b.err == nil
//line /usr/local/go/src/bufio/bufio.go:738
		// _ = "end of CoverTab[2067]"
//line /usr/local/go/src/bufio/bufio.go:738
	}() {
//line /usr/local/go/src/bufio/bufio.go:738
		_go_fuzz_dep_.CoverTab[2068]++
							var n int
							if b.Buffered() == 0 && func() bool {
//line /usr/local/go/src/bufio/bufio.go:740
			_go_fuzz_dep_.CoverTab[2071]++
//line /usr/local/go/src/bufio/bufio.go:740
			return sw == nil
//line /usr/local/go/src/bufio/bufio.go:740
			// _ = "end of CoverTab[2071]"
//line /usr/local/go/src/bufio/bufio.go:740
		}() && func() bool {
//line /usr/local/go/src/bufio/bufio.go:740
			_go_fuzz_dep_.CoverTab[2072]++
//line /usr/local/go/src/bufio/bufio.go:740
			return tryStringWriter
//line /usr/local/go/src/bufio/bufio.go:740
			// _ = "end of CoverTab[2072]"
//line /usr/local/go/src/bufio/bufio.go:740
		}() {
//line /usr/local/go/src/bufio/bufio.go:740
			_go_fuzz_dep_.CoverTab[2073]++

								sw, tryStringWriter = b.wr.(io.StringWriter)
//line /usr/local/go/src/bufio/bufio.go:742
			// _ = "end of CoverTab[2073]"
		} else {
//line /usr/local/go/src/bufio/bufio.go:743
			_go_fuzz_dep_.CoverTab[2074]++
//line /usr/local/go/src/bufio/bufio.go:743
			// _ = "end of CoverTab[2074]"
//line /usr/local/go/src/bufio/bufio.go:743
		}
//line /usr/local/go/src/bufio/bufio.go:743
		// _ = "end of CoverTab[2068]"
//line /usr/local/go/src/bufio/bufio.go:743
		_go_fuzz_dep_.CoverTab[2069]++
							if b.Buffered() == 0 && func() bool {
//line /usr/local/go/src/bufio/bufio.go:744
			_go_fuzz_dep_.CoverTab[2075]++
//line /usr/local/go/src/bufio/bufio.go:744
			return tryStringWriter
//line /usr/local/go/src/bufio/bufio.go:744
			// _ = "end of CoverTab[2075]"
//line /usr/local/go/src/bufio/bufio.go:744
		}() {
//line /usr/local/go/src/bufio/bufio.go:744
			_go_fuzz_dep_.CoverTab[2076]++

//line /usr/local/go/src/bufio/bufio.go:748
			n, b.err = sw.WriteString(s)
//line /usr/local/go/src/bufio/bufio.go:748
			// _ = "end of CoverTab[2076]"
		} else {
//line /usr/local/go/src/bufio/bufio.go:749
			_go_fuzz_dep_.CoverTab[2077]++
								n = copy(b.buf[b.n:], s)
								b.n += n
								b.Flush()
//line /usr/local/go/src/bufio/bufio.go:752
			// _ = "end of CoverTab[2077]"
		}
//line /usr/local/go/src/bufio/bufio.go:753
		// _ = "end of CoverTab[2069]"
//line /usr/local/go/src/bufio/bufio.go:753
		_go_fuzz_dep_.CoverTab[2070]++
							nn += n
							s = s[n:]
//line /usr/local/go/src/bufio/bufio.go:755
		// _ = "end of CoverTab[2070]"
	}
//line /usr/local/go/src/bufio/bufio.go:756
	// _ = "end of CoverTab[2064]"
//line /usr/local/go/src/bufio/bufio.go:756
	_go_fuzz_dep_.CoverTab[2065]++
						if b.err != nil {
//line /usr/local/go/src/bufio/bufio.go:757
		_go_fuzz_dep_.CoverTab[2078]++
							return nn, b.err
//line /usr/local/go/src/bufio/bufio.go:758
		// _ = "end of CoverTab[2078]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:759
		_go_fuzz_dep_.CoverTab[2079]++
//line /usr/local/go/src/bufio/bufio.go:759
		// _ = "end of CoverTab[2079]"
//line /usr/local/go/src/bufio/bufio.go:759
	}
//line /usr/local/go/src/bufio/bufio.go:759
	// _ = "end of CoverTab[2065]"
//line /usr/local/go/src/bufio/bufio.go:759
	_go_fuzz_dep_.CoverTab[2066]++
						n := copy(b.buf[b.n:], s)
						b.n += n
						nn += n
						return nn, nil
//line /usr/local/go/src/bufio/bufio.go:763
	// _ = "end of CoverTab[2066]"
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
	_go_fuzz_dep_.CoverTab[2080]++
						if b.err != nil {
//line /usr/local/go/src/bufio/bufio.go:771
		_go_fuzz_dep_.CoverTab[2084]++
							return 0, b.err
//line /usr/local/go/src/bufio/bufio.go:772
		// _ = "end of CoverTab[2084]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:773
		_go_fuzz_dep_.CoverTab[2085]++
//line /usr/local/go/src/bufio/bufio.go:773
		// _ = "end of CoverTab[2085]"
//line /usr/local/go/src/bufio/bufio.go:773
	}
//line /usr/local/go/src/bufio/bufio.go:773
	// _ = "end of CoverTab[2080]"
//line /usr/local/go/src/bufio/bufio.go:773
	_go_fuzz_dep_.CoverTab[2081]++
						readerFrom, readerFromOK := b.wr.(io.ReaderFrom)
						var m int
						for {
//line /usr/local/go/src/bufio/bufio.go:776
		_go_fuzz_dep_.CoverTab[2086]++
							if b.Available() == 0 {
//line /usr/local/go/src/bufio/bufio.go:777
			_go_fuzz_dep_.CoverTab[2091]++
								if err1 := b.Flush(); err1 != nil {
//line /usr/local/go/src/bufio/bufio.go:778
				_go_fuzz_dep_.CoverTab[2092]++
									return n, err1
//line /usr/local/go/src/bufio/bufio.go:779
				// _ = "end of CoverTab[2092]"
			} else {
//line /usr/local/go/src/bufio/bufio.go:780
				_go_fuzz_dep_.CoverTab[2093]++
//line /usr/local/go/src/bufio/bufio.go:780
				// _ = "end of CoverTab[2093]"
//line /usr/local/go/src/bufio/bufio.go:780
			}
//line /usr/local/go/src/bufio/bufio.go:780
			// _ = "end of CoverTab[2091]"
		} else {
//line /usr/local/go/src/bufio/bufio.go:781
			_go_fuzz_dep_.CoverTab[2094]++
//line /usr/local/go/src/bufio/bufio.go:781
			// _ = "end of CoverTab[2094]"
//line /usr/local/go/src/bufio/bufio.go:781
		}
//line /usr/local/go/src/bufio/bufio.go:781
		// _ = "end of CoverTab[2086]"
//line /usr/local/go/src/bufio/bufio.go:781
		_go_fuzz_dep_.CoverTab[2087]++
							if readerFromOK && func() bool {
//line /usr/local/go/src/bufio/bufio.go:782
			_go_fuzz_dep_.CoverTab[2095]++
//line /usr/local/go/src/bufio/bufio.go:782
			return b.Buffered() == 0
//line /usr/local/go/src/bufio/bufio.go:782
			// _ = "end of CoverTab[2095]"
//line /usr/local/go/src/bufio/bufio.go:782
		}() {
//line /usr/local/go/src/bufio/bufio.go:782
			_go_fuzz_dep_.CoverTab[2096]++
								nn, err := readerFrom.ReadFrom(r)
								b.err = err
								n += nn
								return n, err
//line /usr/local/go/src/bufio/bufio.go:786
			// _ = "end of CoverTab[2096]"
		} else {
//line /usr/local/go/src/bufio/bufio.go:787
			_go_fuzz_dep_.CoverTab[2097]++
//line /usr/local/go/src/bufio/bufio.go:787
			// _ = "end of CoverTab[2097]"
//line /usr/local/go/src/bufio/bufio.go:787
		}
//line /usr/local/go/src/bufio/bufio.go:787
		// _ = "end of CoverTab[2087]"
//line /usr/local/go/src/bufio/bufio.go:787
		_go_fuzz_dep_.CoverTab[2088]++
							nr := 0
							for nr < maxConsecutiveEmptyReads {
//line /usr/local/go/src/bufio/bufio.go:789
			_go_fuzz_dep_.CoverTab[2098]++
								m, err = r.Read(b.buf[b.n:])
								if m != 0 || func() bool {
//line /usr/local/go/src/bufio/bufio.go:791
				_go_fuzz_dep_.CoverTab[2100]++
//line /usr/local/go/src/bufio/bufio.go:791
				return err != nil
//line /usr/local/go/src/bufio/bufio.go:791
				// _ = "end of CoverTab[2100]"
//line /usr/local/go/src/bufio/bufio.go:791
			}() {
//line /usr/local/go/src/bufio/bufio.go:791
				_go_fuzz_dep_.CoverTab[2101]++
									break
//line /usr/local/go/src/bufio/bufio.go:792
				// _ = "end of CoverTab[2101]"
			} else {
//line /usr/local/go/src/bufio/bufio.go:793
				_go_fuzz_dep_.CoverTab[2102]++
//line /usr/local/go/src/bufio/bufio.go:793
				// _ = "end of CoverTab[2102]"
//line /usr/local/go/src/bufio/bufio.go:793
			}
//line /usr/local/go/src/bufio/bufio.go:793
			// _ = "end of CoverTab[2098]"
//line /usr/local/go/src/bufio/bufio.go:793
			_go_fuzz_dep_.CoverTab[2099]++
								nr++
//line /usr/local/go/src/bufio/bufio.go:794
			// _ = "end of CoverTab[2099]"
		}
//line /usr/local/go/src/bufio/bufio.go:795
		// _ = "end of CoverTab[2088]"
//line /usr/local/go/src/bufio/bufio.go:795
		_go_fuzz_dep_.CoverTab[2089]++
							if nr == maxConsecutiveEmptyReads {
//line /usr/local/go/src/bufio/bufio.go:796
			_go_fuzz_dep_.CoverTab[2103]++
								return n, io.ErrNoProgress
//line /usr/local/go/src/bufio/bufio.go:797
			// _ = "end of CoverTab[2103]"
		} else {
//line /usr/local/go/src/bufio/bufio.go:798
			_go_fuzz_dep_.CoverTab[2104]++
//line /usr/local/go/src/bufio/bufio.go:798
			// _ = "end of CoverTab[2104]"
//line /usr/local/go/src/bufio/bufio.go:798
		}
//line /usr/local/go/src/bufio/bufio.go:798
		// _ = "end of CoverTab[2089]"
//line /usr/local/go/src/bufio/bufio.go:798
		_go_fuzz_dep_.CoverTab[2090]++
							b.n += m
							n += int64(m)
							if err != nil {
//line /usr/local/go/src/bufio/bufio.go:801
			_go_fuzz_dep_.CoverTab[2105]++
								break
//line /usr/local/go/src/bufio/bufio.go:802
			// _ = "end of CoverTab[2105]"
		} else {
//line /usr/local/go/src/bufio/bufio.go:803
			_go_fuzz_dep_.CoverTab[2106]++
//line /usr/local/go/src/bufio/bufio.go:803
			// _ = "end of CoverTab[2106]"
//line /usr/local/go/src/bufio/bufio.go:803
		}
//line /usr/local/go/src/bufio/bufio.go:803
		// _ = "end of CoverTab[2090]"
	}
//line /usr/local/go/src/bufio/bufio.go:804
	// _ = "end of CoverTab[2081]"
//line /usr/local/go/src/bufio/bufio.go:804
	_go_fuzz_dep_.CoverTab[2082]++
						if err == io.EOF {
//line /usr/local/go/src/bufio/bufio.go:805
		_go_fuzz_dep_.CoverTab[2107]++

							if b.Available() == 0 {
//line /usr/local/go/src/bufio/bufio.go:807
			_go_fuzz_dep_.CoverTab[2108]++
								err = b.Flush()
//line /usr/local/go/src/bufio/bufio.go:808
			// _ = "end of CoverTab[2108]"
		} else {
//line /usr/local/go/src/bufio/bufio.go:809
			_go_fuzz_dep_.CoverTab[2109]++
								err = nil
//line /usr/local/go/src/bufio/bufio.go:810
			// _ = "end of CoverTab[2109]"
		}
//line /usr/local/go/src/bufio/bufio.go:811
		// _ = "end of CoverTab[2107]"
	} else {
//line /usr/local/go/src/bufio/bufio.go:812
		_go_fuzz_dep_.CoverTab[2110]++
//line /usr/local/go/src/bufio/bufio.go:812
		// _ = "end of CoverTab[2110]"
//line /usr/local/go/src/bufio/bufio.go:812
	}
//line /usr/local/go/src/bufio/bufio.go:812
	// _ = "end of CoverTab[2082]"
//line /usr/local/go/src/bufio/bufio.go:812
	_go_fuzz_dep_.CoverTab[2083]++
						return n, err
//line /usr/local/go/src/bufio/bufio.go:813
	// _ = "end of CoverTab[2083]"
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
	_go_fuzz_dep_.CoverTab[2111]++
						return &ReadWriter{r, w}
//line /usr/local/go/src/bufio/bufio.go:827
	// _ = "end of CoverTab[2111]"
}

//line /usr/local/go/src/bufio/bufio.go:828
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/bufio/bufio.go:828
var _ = _go_fuzz_dep_.CoverTab
