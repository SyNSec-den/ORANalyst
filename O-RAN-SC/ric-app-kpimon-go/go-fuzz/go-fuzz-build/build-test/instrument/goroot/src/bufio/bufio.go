// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /snap/go/10455/src/bufio/bufio.go:5
// Package bufio implements buffered I/O. It wraps an io.Reader or io.Writer
//line /snap/go/10455/src/bufio/bufio.go:5
// object, creating another object (Reader or Writer) that also implements
//line /snap/go/10455/src/bufio/bufio.go:5
// the interface but provides buffering and some help for textual I/O.
//line /snap/go/10455/src/bufio/bufio.go:8
package bufio

//line /snap/go/10455/src/bufio/bufio.go:8
import (
//line /snap/go/10455/src/bufio/bufio.go:8
	_go_fuzz_dep_ "go-fuzz-dep"
//line /snap/go/10455/src/bufio/bufio.go:8
)
//line /snap/go/10455/src/bufio/bufio.go:8
import (
//line /snap/go/10455/src/bufio/bufio.go:8
	_atomic_ "sync/atomic"
//line /snap/go/10455/src/bufio/bufio.go:8
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

//line /snap/go/10455/src/bufio/bufio.go:31
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
//line /snap/go/10455/src/bufio/bufio.go:44
// size. If the argument io.Reader is already a Reader with large enough
//line /snap/go/10455/src/bufio/bufio.go:44
// size, it returns the underlying Reader.
//line /snap/go/10455/src/bufio/bufio.go:47
func NewReaderSize(rd io.Reader, size int) *Reader {
//line /snap/go/10455/src/bufio/bufio.go:47
	_go_fuzz_dep_.CoverTab[1809]++

						b, ok := rd.(*Reader)
						if ok && func() bool {
//line /snap/go/10455/src/bufio/bufio.go:50
		_go_fuzz_dep_.CoverTab[1812]++
//line /snap/go/10455/src/bufio/bufio.go:50
		return len(b.buf) >= size
//line /snap/go/10455/src/bufio/bufio.go:50
		// _ = "end of CoverTab[1812]"
//line /snap/go/10455/src/bufio/bufio.go:50
	}() {
//line /snap/go/10455/src/bufio/bufio.go:50
		_go_fuzz_dep_.CoverTab[525671]++
//line /snap/go/10455/src/bufio/bufio.go:50
		_go_fuzz_dep_.CoverTab[1813]++
							return b
//line /snap/go/10455/src/bufio/bufio.go:51
		// _ = "end of CoverTab[1813]"
	} else {
//line /snap/go/10455/src/bufio/bufio.go:52
		_go_fuzz_dep_.CoverTab[525672]++
//line /snap/go/10455/src/bufio/bufio.go:52
		_go_fuzz_dep_.CoverTab[1814]++
//line /snap/go/10455/src/bufio/bufio.go:52
		// _ = "end of CoverTab[1814]"
//line /snap/go/10455/src/bufio/bufio.go:52
	}
//line /snap/go/10455/src/bufio/bufio.go:52
	// _ = "end of CoverTab[1809]"
//line /snap/go/10455/src/bufio/bufio.go:52
	_go_fuzz_dep_.CoverTab[1810]++
						if size < minReadBufferSize {
//line /snap/go/10455/src/bufio/bufio.go:53
		_go_fuzz_dep_.CoverTab[525673]++
//line /snap/go/10455/src/bufio/bufio.go:53
		_go_fuzz_dep_.CoverTab[1815]++
							size = minReadBufferSize
//line /snap/go/10455/src/bufio/bufio.go:54
		// _ = "end of CoverTab[1815]"
	} else {
//line /snap/go/10455/src/bufio/bufio.go:55
		_go_fuzz_dep_.CoverTab[525674]++
//line /snap/go/10455/src/bufio/bufio.go:55
		_go_fuzz_dep_.CoverTab[1816]++
//line /snap/go/10455/src/bufio/bufio.go:55
		// _ = "end of CoverTab[1816]"
//line /snap/go/10455/src/bufio/bufio.go:55
	}
//line /snap/go/10455/src/bufio/bufio.go:55
	// _ = "end of CoverTab[1810]"
//line /snap/go/10455/src/bufio/bufio.go:55
	_go_fuzz_dep_.CoverTab[1811]++
						r := new(Reader)
						r.reset(make([]byte, size), rd)
						return r
//line /snap/go/10455/src/bufio/bufio.go:58
	// _ = "end of CoverTab[1811]"
}

// NewReader returns a new Reader whose buffer has the default size.
func NewReader(rd io.Reader) *Reader {
//line /snap/go/10455/src/bufio/bufio.go:62
	_go_fuzz_dep_.CoverTab[1817]++
						return NewReaderSize(rd, defaultBufSize)
//line /snap/go/10455/src/bufio/bufio.go:63
	// _ = "end of CoverTab[1817]"
}

// Size returns the size of the underlying buffer in bytes.
func (b *Reader) Size() int {
//line /snap/go/10455/src/bufio/bufio.go:67
	_go_fuzz_dep_.CoverTab[1818]++
//line /snap/go/10455/src/bufio/bufio.go:67
	return len(b.buf)
//line /snap/go/10455/src/bufio/bufio.go:67
	// _ = "end of CoverTab[1818]"
//line /snap/go/10455/src/bufio/bufio.go:67
}

// Reset discards any buffered data, resets all state, and switches
//line /snap/go/10455/src/bufio/bufio.go:69
// the buffered reader to read from r.
//line /snap/go/10455/src/bufio/bufio.go:69
// Calling Reset on the zero value of Reader initializes the internal buffer
//line /snap/go/10455/src/bufio/bufio.go:69
// to the default size.
//line /snap/go/10455/src/bufio/bufio.go:69
// Calling b.Reset(b) (that is, resetting a Reader to itself) does nothing.
//line /snap/go/10455/src/bufio/bufio.go:74
func (b *Reader) Reset(r io.Reader) {
//line /snap/go/10455/src/bufio/bufio.go:74
	_go_fuzz_dep_.CoverTab[1819]++

//line /snap/go/10455/src/bufio/bufio.go:78
	if b == r {
//line /snap/go/10455/src/bufio/bufio.go:78
		_go_fuzz_dep_.CoverTab[525675]++
//line /snap/go/10455/src/bufio/bufio.go:78
		_go_fuzz_dep_.CoverTab[1822]++
							return
//line /snap/go/10455/src/bufio/bufio.go:79
		// _ = "end of CoverTab[1822]"
	} else {
//line /snap/go/10455/src/bufio/bufio.go:80
		_go_fuzz_dep_.CoverTab[525676]++
//line /snap/go/10455/src/bufio/bufio.go:80
		_go_fuzz_dep_.CoverTab[1823]++
//line /snap/go/10455/src/bufio/bufio.go:80
		// _ = "end of CoverTab[1823]"
//line /snap/go/10455/src/bufio/bufio.go:80
	}
//line /snap/go/10455/src/bufio/bufio.go:80
	// _ = "end of CoverTab[1819]"
//line /snap/go/10455/src/bufio/bufio.go:80
	_go_fuzz_dep_.CoverTab[1820]++
						if b.buf == nil {
//line /snap/go/10455/src/bufio/bufio.go:81
		_go_fuzz_dep_.CoverTab[525677]++
//line /snap/go/10455/src/bufio/bufio.go:81
		_go_fuzz_dep_.CoverTab[1824]++
							b.buf = make([]byte, defaultBufSize)
//line /snap/go/10455/src/bufio/bufio.go:82
		// _ = "end of CoverTab[1824]"
	} else {
//line /snap/go/10455/src/bufio/bufio.go:83
		_go_fuzz_dep_.CoverTab[525678]++
//line /snap/go/10455/src/bufio/bufio.go:83
		_go_fuzz_dep_.CoverTab[1825]++
//line /snap/go/10455/src/bufio/bufio.go:83
		// _ = "end of CoverTab[1825]"
//line /snap/go/10455/src/bufio/bufio.go:83
	}
//line /snap/go/10455/src/bufio/bufio.go:83
	// _ = "end of CoverTab[1820]"
//line /snap/go/10455/src/bufio/bufio.go:83
	_go_fuzz_dep_.CoverTab[1821]++
						b.reset(b.buf, r)
//line /snap/go/10455/src/bufio/bufio.go:84
	// _ = "end of CoverTab[1821]"
}

func (b *Reader) reset(buf []byte, r io.Reader) {
//line /snap/go/10455/src/bufio/bufio.go:87
	_go_fuzz_dep_.CoverTab[1826]++
						*b = Reader{
		buf:		buf,
		rd:		r,
		lastByte:	-1,
		lastRuneSize:	-1,
	}
//line /snap/go/10455/src/bufio/bufio.go:93
	// _ = "end of CoverTab[1826]"
}

var errNegativeRead = errors.New("bufio: reader returned negative count from Read")

// fill reads a new chunk into the buffer.
func (b *Reader) fill() {
//line /snap/go/10455/src/bufio/bufio.go:99
	_go_fuzz_dep_.CoverTab[1827]++

						if b.r > 0 {
//line /snap/go/10455/src/bufio/bufio.go:101
		_go_fuzz_dep_.CoverTab[525679]++
//line /snap/go/10455/src/bufio/bufio.go:101
		_go_fuzz_dep_.CoverTab[1831]++
							copy(b.buf, b.buf[b.r:b.w])
							b.w -= b.r
							b.r = 0
//line /snap/go/10455/src/bufio/bufio.go:104
		// _ = "end of CoverTab[1831]"
	} else {
//line /snap/go/10455/src/bufio/bufio.go:105
		_go_fuzz_dep_.CoverTab[525680]++
//line /snap/go/10455/src/bufio/bufio.go:105
		_go_fuzz_dep_.CoverTab[1832]++
//line /snap/go/10455/src/bufio/bufio.go:105
		// _ = "end of CoverTab[1832]"
//line /snap/go/10455/src/bufio/bufio.go:105
	}
//line /snap/go/10455/src/bufio/bufio.go:105
	// _ = "end of CoverTab[1827]"
//line /snap/go/10455/src/bufio/bufio.go:105
	_go_fuzz_dep_.CoverTab[1828]++

						if b.w >= len(b.buf) {
//line /snap/go/10455/src/bufio/bufio.go:107
		_go_fuzz_dep_.CoverTab[525681]++
//line /snap/go/10455/src/bufio/bufio.go:107
		_go_fuzz_dep_.CoverTab[1833]++
							panic("bufio: tried to fill full buffer")
//line /snap/go/10455/src/bufio/bufio.go:108
		// _ = "end of CoverTab[1833]"
	} else {
//line /snap/go/10455/src/bufio/bufio.go:109
		_go_fuzz_dep_.CoverTab[525682]++
//line /snap/go/10455/src/bufio/bufio.go:109
		_go_fuzz_dep_.CoverTab[1834]++
//line /snap/go/10455/src/bufio/bufio.go:109
		// _ = "end of CoverTab[1834]"
//line /snap/go/10455/src/bufio/bufio.go:109
	}
//line /snap/go/10455/src/bufio/bufio.go:109
	// _ = "end of CoverTab[1828]"
//line /snap/go/10455/src/bufio/bufio.go:109
	_go_fuzz_dep_.CoverTab[1829]++
//line /snap/go/10455/src/bufio/bufio.go:109
	_go_fuzz_dep_.CoverTab[786554] = 0

//line /snap/go/10455/src/bufio/bufio.go:112
	for i := maxConsecutiveEmptyReads; i > 0; i-- {
//line /snap/go/10455/src/bufio/bufio.go:112
		if _go_fuzz_dep_.CoverTab[786554] == 0 {
//line /snap/go/10455/src/bufio/bufio.go:112
			_go_fuzz_dep_.CoverTab[525841]++
//line /snap/go/10455/src/bufio/bufio.go:112
		} else {
//line /snap/go/10455/src/bufio/bufio.go:112
			_go_fuzz_dep_.CoverTab[525842]++
//line /snap/go/10455/src/bufio/bufio.go:112
		}
//line /snap/go/10455/src/bufio/bufio.go:112
		_go_fuzz_dep_.CoverTab[786554] = 1
//line /snap/go/10455/src/bufio/bufio.go:112
		_go_fuzz_dep_.CoverTab[1835]++
							n, err := b.rd.Read(b.buf[b.w:])
							if n < 0 {
//line /snap/go/10455/src/bufio/bufio.go:114
			_go_fuzz_dep_.CoverTab[525683]++
//line /snap/go/10455/src/bufio/bufio.go:114
			_go_fuzz_dep_.CoverTab[1838]++
								panic(errNegativeRead)
//line /snap/go/10455/src/bufio/bufio.go:115
			// _ = "end of CoverTab[1838]"
		} else {
//line /snap/go/10455/src/bufio/bufio.go:116
			_go_fuzz_dep_.CoverTab[525684]++
//line /snap/go/10455/src/bufio/bufio.go:116
			_go_fuzz_dep_.CoverTab[1839]++
//line /snap/go/10455/src/bufio/bufio.go:116
			// _ = "end of CoverTab[1839]"
//line /snap/go/10455/src/bufio/bufio.go:116
		}
//line /snap/go/10455/src/bufio/bufio.go:116
		// _ = "end of CoverTab[1835]"
//line /snap/go/10455/src/bufio/bufio.go:116
		_go_fuzz_dep_.CoverTab[1836]++
							b.w += n
							if err != nil {
//line /snap/go/10455/src/bufio/bufio.go:118
			_go_fuzz_dep_.CoverTab[525685]++
//line /snap/go/10455/src/bufio/bufio.go:118
			_go_fuzz_dep_.CoverTab[1840]++
								b.err = err
								return
//line /snap/go/10455/src/bufio/bufio.go:120
			// _ = "end of CoverTab[1840]"
		} else {
//line /snap/go/10455/src/bufio/bufio.go:121
			_go_fuzz_dep_.CoverTab[525686]++
//line /snap/go/10455/src/bufio/bufio.go:121
			_go_fuzz_dep_.CoverTab[1841]++
//line /snap/go/10455/src/bufio/bufio.go:121
			// _ = "end of CoverTab[1841]"
//line /snap/go/10455/src/bufio/bufio.go:121
		}
//line /snap/go/10455/src/bufio/bufio.go:121
		// _ = "end of CoverTab[1836]"
//line /snap/go/10455/src/bufio/bufio.go:121
		_go_fuzz_dep_.CoverTab[1837]++
							if n > 0 {
//line /snap/go/10455/src/bufio/bufio.go:122
			_go_fuzz_dep_.CoverTab[525687]++
//line /snap/go/10455/src/bufio/bufio.go:122
			_go_fuzz_dep_.CoverTab[1842]++
								return
//line /snap/go/10455/src/bufio/bufio.go:123
			// _ = "end of CoverTab[1842]"
		} else {
//line /snap/go/10455/src/bufio/bufio.go:124
			_go_fuzz_dep_.CoverTab[525688]++
//line /snap/go/10455/src/bufio/bufio.go:124
			_go_fuzz_dep_.CoverTab[1843]++
//line /snap/go/10455/src/bufio/bufio.go:124
			// _ = "end of CoverTab[1843]"
//line /snap/go/10455/src/bufio/bufio.go:124
		}
//line /snap/go/10455/src/bufio/bufio.go:124
		// _ = "end of CoverTab[1837]"
	}
//line /snap/go/10455/src/bufio/bufio.go:125
	if _go_fuzz_dep_.CoverTab[786554] == 0 {
//line /snap/go/10455/src/bufio/bufio.go:125
		_go_fuzz_dep_.CoverTab[525843]++
//line /snap/go/10455/src/bufio/bufio.go:125
	} else {
//line /snap/go/10455/src/bufio/bufio.go:125
		_go_fuzz_dep_.CoverTab[525844]++
//line /snap/go/10455/src/bufio/bufio.go:125
	}
//line /snap/go/10455/src/bufio/bufio.go:125
	// _ = "end of CoverTab[1829]"
//line /snap/go/10455/src/bufio/bufio.go:125
	_go_fuzz_dep_.CoverTab[1830]++
						b.err = io.ErrNoProgress
//line /snap/go/10455/src/bufio/bufio.go:126
	// _ = "end of CoverTab[1830]"
}

func (b *Reader) readErr() error {
//line /snap/go/10455/src/bufio/bufio.go:129
	_go_fuzz_dep_.CoverTab[1844]++
						err := b.err
						b.err = nil
						return err
//line /snap/go/10455/src/bufio/bufio.go:132
	// _ = "end of CoverTab[1844]"
}

// Peek returns the next n bytes without advancing the reader. The bytes stop
//line /snap/go/10455/src/bufio/bufio.go:135
// being valid at the next read call. If Peek returns fewer than n bytes, it
//line /snap/go/10455/src/bufio/bufio.go:135
// also returns an error explaining why the read is short. The error is
//line /snap/go/10455/src/bufio/bufio.go:135
// ErrBufferFull if n is larger than b's buffer size.
//line /snap/go/10455/src/bufio/bufio.go:135
//
//line /snap/go/10455/src/bufio/bufio.go:135
// Calling Peek prevents a UnreadByte or UnreadRune call from succeeding
//line /snap/go/10455/src/bufio/bufio.go:135
// until the next read operation.
//line /snap/go/10455/src/bufio/bufio.go:142
func (b *Reader) Peek(n int) ([]byte, error) {
//line /snap/go/10455/src/bufio/bufio.go:142
	_go_fuzz_dep_.CoverTab[1845]++
						if n < 0 {
//line /snap/go/10455/src/bufio/bufio.go:143
		_go_fuzz_dep_.CoverTab[525689]++
//line /snap/go/10455/src/bufio/bufio.go:143
		_go_fuzz_dep_.CoverTab[1850]++
							return nil, ErrNegativeCount
//line /snap/go/10455/src/bufio/bufio.go:144
		// _ = "end of CoverTab[1850]"
	} else {
//line /snap/go/10455/src/bufio/bufio.go:145
		_go_fuzz_dep_.CoverTab[525690]++
//line /snap/go/10455/src/bufio/bufio.go:145
		_go_fuzz_dep_.CoverTab[1851]++
//line /snap/go/10455/src/bufio/bufio.go:145
		// _ = "end of CoverTab[1851]"
//line /snap/go/10455/src/bufio/bufio.go:145
	}
//line /snap/go/10455/src/bufio/bufio.go:145
	// _ = "end of CoverTab[1845]"
//line /snap/go/10455/src/bufio/bufio.go:145
	_go_fuzz_dep_.CoverTab[1846]++

						b.lastByte = -1
						b.lastRuneSize = -1
//line /snap/go/10455/src/bufio/bufio.go:148
	_go_fuzz_dep_.CoverTab[786555] = 0

						for b.w-b.r < n && func() bool {
//line /snap/go/10455/src/bufio/bufio.go:150
		_go_fuzz_dep_.CoverTab[1852]++
//line /snap/go/10455/src/bufio/bufio.go:150
		return b.w-b.r < len(b.buf)
//line /snap/go/10455/src/bufio/bufio.go:150
		// _ = "end of CoverTab[1852]"
//line /snap/go/10455/src/bufio/bufio.go:150
	}() && func() bool {
//line /snap/go/10455/src/bufio/bufio.go:150
		_go_fuzz_dep_.CoverTab[1853]++
//line /snap/go/10455/src/bufio/bufio.go:150
		return b.err == nil
//line /snap/go/10455/src/bufio/bufio.go:150
		// _ = "end of CoverTab[1853]"
//line /snap/go/10455/src/bufio/bufio.go:150
	}() {
//line /snap/go/10455/src/bufio/bufio.go:150
		if _go_fuzz_dep_.CoverTab[786555] == 0 {
//line /snap/go/10455/src/bufio/bufio.go:150
			_go_fuzz_dep_.CoverTab[525845]++
//line /snap/go/10455/src/bufio/bufio.go:150
		} else {
//line /snap/go/10455/src/bufio/bufio.go:150
			_go_fuzz_dep_.CoverTab[525846]++
//line /snap/go/10455/src/bufio/bufio.go:150
		}
//line /snap/go/10455/src/bufio/bufio.go:150
		_go_fuzz_dep_.CoverTab[786555] = 1
//line /snap/go/10455/src/bufio/bufio.go:150
		_go_fuzz_dep_.CoverTab[1854]++
							b.fill()
//line /snap/go/10455/src/bufio/bufio.go:151
		// _ = "end of CoverTab[1854]"
	}
//line /snap/go/10455/src/bufio/bufio.go:152
	if _go_fuzz_dep_.CoverTab[786555] == 0 {
//line /snap/go/10455/src/bufio/bufio.go:152
		_go_fuzz_dep_.CoverTab[525847]++
//line /snap/go/10455/src/bufio/bufio.go:152
	} else {
//line /snap/go/10455/src/bufio/bufio.go:152
		_go_fuzz_dep_.CoverTab[525848]++
//line /snap/go/10455/src/bufio/bufio.go:152
	}
//line /snap/go/10455/src/bufio/bufio.go:152
	// _ = "end of CoverTab[1846]"
//line /snap/go/10455/src/bufio/bufio.go:152
	_go_fuzz_dep_.CoverTab[1847]++

						if n > len(b.buf) {
//line /snap/go/10455/src/bufio/bufio.go:154
		_go_fuzz_dep_.CoverTab[525691]++
//line /snap/go/10455/src/bufio/bufio.go:154
		_go_fuzz_dep_.CoverTab[1855]++
							return b.buf[b.r:b.w], ErrBufferFull
//line /snap/go/10455/src/bufio/bufio.go:155
		// _ = "end of CoverTab[1855]"
	} else {
//line /snap/go/10455/src/bufio/bufio.go:156
		_go_fuzz_dep_.CoverTab[525692]++
//line /snap/go/10455/src/bufio/bufio.go:156
		_go_fuzz_dep_.CoverTab[1856]++
//line /snap/go/10455/src/bufio/bufio.go:156
		// _ = "end of CoverTab[1856]"
//line /snap/go/10455/src/bufio/bufio.go:156
	}
//line /snap/go/10455/src/bufio/bufio.go:156
	// _ = "end of CoverTab[1847]"
//line /snap/go/10455/src/bufio/bufio.go:156
	_go_fuzz_dep_.CoverTab[1848]++

	// 0 <= n <= len(b.buf)
	var err error
	if avail := b.w - b.r; avail < n {
//line /snap/go/10455/src/bufio/bufio.go:160
		_go_fuzz_dep_.CoverTab[525693]++
//line /snap/go/10455/src/bufio/bufio.go:160
		_go_fuzz_dep_.CoverTab[1857]++

							n = avail
							err = b.readErr()
							if err == nil {
//line /snap/go/10455/src/bufio/bufio.go:164
			_go_fuzz_dep_.CoverTab[525695]++
//line /snap/go/10455/src/bufio/bufio.go:164
			_go_fuzz_dep_.CoverTab[1858]++
								err = ErrBufferFull
//line /snap/go/10455/src/bufio/bufio.go:165
			// _ = "end of CoverTab[1858]"
		} else {
//line /snap/go/10455/src/bufio/bufio.go:166
			_go_fuzz_dep_.CoverTab[525696]++
//line /snap/go/10455/src/bufio/bufio.go:166
			_go_fuzz_dep_.CoverTab[1859]++
//line /snap/go/10455/src/bufio/bufio.go:166
			// _ = "end of CoverTab[1859]"
//line /snap/go/10455/src/bufio/bufio.go:166
		}
//line /snap/go/10455/src/bufio/bufio.go:166
		// _ = "end of CoverTab[1857]"
	} else {
//line /snap/go/10455/src/bufio/bufio.go:167
		_go_fuzz_dep_.CoverTab[525694]++
//line /snap/go/10455/src/bufio/bufio.go:167
		_go_fuzz_dep_.CoverTab[1860]++
//line /snap/go/10455/src/bufio/bufio.go:167
		// _ = "end of CoverTab[1860]"
//line /snap/go/10455/src/bufio/bufio.go:167
	}
//line /snap/go/10455/src/bufio/bufio.go:167
	// _ = "end of CoverTab[1848]"
//line /snap/go/10455/src/bufio/bufio.go:167
	_go_fuzz_dep_.CoverTab[1849]++
						return b.buf[b.r : b.r+n], err
//line /snap/go/10455/src/bufio/bufio.go:168
	// _ = "end of CoverTab[1849]"
}

// Discard skips the next n bytes, returning the number of bytes discarded.
//line /snap/go/10455/src/bufio/bufio.go:171
//
//line /snap/go/10455/src/bufio/bufio.go:171
// If Discard skips fewer than n bytes, it also returns an error.
//line /snap/go/10455/src/bufio/bufio.go:171
// If 0 <= n <= b.Buffered(), Discard is guaranteed to succeed without
//line /snap/go/10455/src/bufio/bufio.go:171
// reading from the underlying io.Reader.
//line /snap/go/10455/src/bufio/bufio.go:176
func (b *Reader) Discard(n int) (discarded int, err error) {
//line /snap/go/10455/src/bufio/bufio.go:176
	_go_fuzz_dep_.CoverTab[1861]++
						if n < 0 {
//line /snap/go/10455/src/bufio/bufio.go:177
		_go_fuzz_dep_.CoverTab[525697]++
//line /snap/go/10455/src/bufio/bufio.go:177
		_go_fuzz_dep_.CoverTab[1864]++
							return 0, ErrNegativeCount
//line /snap/go/10455/src/bufio/bufio.go:178
		// _ = "end of CoverTab[1864]"
	} else {
//line /snap/go/10455/src/bufio/bufio.go:179
		_go_fuzz_dep_.CoverTab[525698]++
//line /snap/go/10455/src/bufio/bufio.go:179
		_go_fuzz_dep_.CoverTab[1865]++
//line /snap/go/10455/src/bufio/bufio.go:179
		// _ = "end of CoverTab[1865]"
//line /snap/go/10455/src/bufio/bufio.go:179
	}
//line /snap/go/10455/src/bufio/bufio.go:179
	// _ = "end of CoverTab[1861]"
//line /snap/go/10455/src/bufio/bufio.go:179
	_go_fuzz_dep_.CoverTab[1862]++
						if n == 0 {
//line /snap/go/10455/src/bufio/bufio.go:180
		_go_fuzz_dep_.CoverTab[525699]++
//line /snap/go/10455/src/bufio/bufio.go:180
		_go_fuzz_dep_.CoverTab[1866]++
							return
//line /snap/go/10455/src/bufio/bufio.go:181
		// _ = "end of CoverTab[1866]"
	} else {
//line /snap/go/10455/src/bufio/bufio.go:182
		_go_fuzz_dep_.CoverTab[525700]++
//line /snap/go/10455/src/bufio/bufio.go:182
		_go_fuzz_dep_.CoverTab[1867]++
//line /snap/go/10455/src/bufio/bufio.go:182
		// _ = "end of CoverTab[1867]"
//line /snap/go/10455/src/bufio/bufio.go:182
	}
//line /snap/go/10455/src/bufio/bufio.go:182
	// _ = "end of CoverTab[1862]"
//line /snap/go/10455/src/bufio/bufio.go:182
	_go_fuzz_dep_.CoverTab[1863]++

						b.lastByte = -1
						b.lastRuneSize = -1

						remain := n
//line /snap/go/10455/src/bufio/bufio.go:187
	_go_fuzz_dep_.CoverTab[786556] = 0
						for {
//line /snap/go/10455/src/bufio/bufio.go:188
		if _go_fuzz_dep_.CoverTab[786556] == 0 {
//line /snap/go/10455/src/bufio/bufio.go:188
			_go_fuzz_dep_.CoverTab[525849]++
//line /snap/go/10455/src/bufio/bufio.go:188
		} else {
//line /snap/go/10455/src/bufio/bufio.go:188
			_go_fuzz_dep_.CoverTab[525850]++
//line /snap/go/10455/src/bufio/bufio.go:188
		}
//line /snap/go/10455/src/bufio/bufio.go:188
		_go_fuzz_dep_.CoverTab[786556] = 1
//line /snap/go/10455/src/bufio/bufio.go:188
		_go_fuzz_dep_.CoverTab[1868]++
							skip := b.Buffered()
							if skip == 0 {
//line /snap/go/10455/src/bufio/bufio.go:190
			_go_fuzz_dep_.CoverTab[525701]++
//line /snap/go/10455/src/bufio/bufio.go:190
			_go_fuzz_dep_.CoverTab[1872]++
								b.fill()
								skip = b.Buffered()
//line /snap/go/10455/src/bufio/bufio.go:192
			// _ = "end of CoverTab[1872]"
		} else {
//line /snap/go/10455/src/bufio/bufio.go:193
			_go_fuzz_dep_.CoverTab[525702]++
//line /snap/go/10455/src/bufio/bufio.go:193
			_go_fuzz_dep_.CoverTab[1873]++
//line /snap/go/10455/src/bufio/bufio.go:193
			// _ = "end of CoverTab[1873]"
//line /snap/go/10455/src/bufio/bufio.go:193
		}
//line /snap/go/10455/src/bufio/bufio.go:193
		// _ = "end of CoverTab[1868]"
//line /snap/go/10455/src/bufio/bufio.go:193
		_go_fuzz_dep_.CoverTab[1869]++
							if skip > remain {
//line /snap/go/10455/src/bufio/bufio.go:194
			_go_fuzz_dep_.CoverTab[525703]++
//line /snap/go/10455/src/bufio/bufio.go:194
			_go_fuzz_dep_.CoverTab[1874]++
								skip = remain
//line /snap/go/10455/src/bufio/bufio.go:195
			// _ = "end of CoverTab[1874]"
		} else {
//line /snap/go/10455/src/bufio/bufio.go:196
			_go_fuzz_dep_.CoverTab[525704]++
//line /snap/go/10455/src/bufio/bufio.go:196
			_go_fuzz_dep_.CoverTab[1875]++
//line /snap/go/10455/src/bufio/bufio.go:196
			// _ = "end of CoverTab[1875]"
//line /snap/go/10455/src/bufio/bufio.go:196
		}
//line /snap/go/10455/src/bufio/bufio.go:196
		// _ = "end of CoverTab[1869]"
//line /snap/go/10455/src/bufio/bufio.go:196
		_go_fuzz_dep_.CoverTab[1870]++
							b.r += skip
							remain -= skip
							if remain == 0 {
//line /snap/go/10455/src/bufio/bufio.go:199
			_go_fuzz_dep_.CoverTab[525705]++
//line /snap/go/10455/src/bufio/bufio.go:199
			_go_fuzz_dep_.CoverTab[1876]++
								return n, nil
//line /snap/go/10455/src/bufio/bufio.go:200
			// _ = "end of CoverTab[1876]"
		} else {
//line /snap/go/10455/src/bufio/bufio.go:201
			_go_fuzz_dep_.CoverTab[525706]++
//line /snap/go/10455/src/bufio/bufio.go:201
			_go_fuzz_dep_.CoverTab[1877]++
//line /snap/go/10455/src/bufio/bufio.go:201
			// _ = "end of CoverTab[1877]"
//line /snap/go/10455/src/bufio/bufio.go:201
		}
//line /snap/go/10455/src/bufio/bufio.go:201
		// _ = "end of CoverTab[1870]"
//line /snap/go/10455/src/bufio/bufio.go:201
		_go_fuzz_dep_.CoverTab[1871]++
							if b.err != nil {
//line /snap/go/10455/src/bufio/bufio.go:202
			_go_fuzz_dep_.CoverTab[525707]++
//line /snap/go/10455/src/bufio/bufio.go:202
			_go_fuzz_dep_.CoverTab[1878]++
								return n - remain, b.readErr()
//line /snap/go/10455/src/bufio/bufio.go:203
			// _ = "end of CoverTab[1878]"
		} else {
//line /snap/go/10455/src/bufio/bufio.go:204
			_go_fuzz_dep_.CoverTab[525708]++
//line /snap/go/10455/src/bufio/bufio.go:204
			_go_fuzz_dep_.CoverTab[1879]++
//line /snap/go/10455/src/bufio/bufio.go:204
			// _ = "end of CoverTab[1879]"
//line /snap/go/10455/src/bufio/bufio.go:204
		}
//line /snap/go/10455/src/bufio/bufio.go:204
		// _ = "end of CoverTab[1871]"
	}
//line /snap/go/10455/src/bufio/bufio.go:205
	// _ = "end of CoverTab[1863]"
}

// Read reads data into p.
//line /snap/go/10455/src/bufio/bufio.go:208
// It returns the number of bytes read into p.
//line /snap/go/10455/src/bufio/bufio.go:208
// The bytes are taken from at most one Read on the underlying Reader,
//line /snap/go/10455/src/bufio/bufio.go:208
// hence n may be less than len(p).
//line /snap/go/10455/src/bufio/bufio.go:208
// To read exactly len(p) bytes, use io.ReadFull(b, p).
//line /snap/go/10455/src/bufio/bufio.go:208
// If the underlying Reader can return a non-zero count with io.EOF,
//line /snap/go/10455/src/bufio/bufio.go:208
// then this Read method can do so as well; see the [io.Reader] docs.
//line /snap/go/10455/src/bufio/bufio.go:215
func (b *Reader) Read(p []byte) (n int, err error) {
//line /snap/go/10455/src/bufio/bufio.go:215
	_go_fuzz_dep_.CoverTab[1880]++
						n = len(p)
						if n == 0 {
//line /snap/go/10455/src/bufio/bufio.go:217
		_go_fuzz_dep_.CoverTab[525709]++
//line /snap/go/10455/src/bufio/bufio.go:217
		_go_fuzz_dep_.CoverTab[1883]++
							if b.Buffered() > 0 {
//line /snap/go/10455/src/bufio/bufio.go:218
			_go_fuzz_dep_.CoverTab[525711]++
//line /snap/go/10455/src/bufio/bufio.go:218
			_go_fuzz_dep_.CoverTab[1885]++
								return 0, nil
//line /snap/go/10455/src/bufio/bufio.go:219
			// _ = "end of CoverTab[1885]"
		} else {
//line /snap/go/10455/src/bufio/bufio.go:220
			_go_fuzz_dep_.CoverTab[525712]++
//line /snap/go/10455/src/bufio/bufio.go:220
			_go_fuzz_dep_.CoverTab[1886]++
//line /snap/go/10455/src/bufio/bufio.go:220
			// _ = "end of CoverTab[1886]"
//line /snap/go/10455/src/bufio/bufio.go:220
		}
//line /snap/go/10455/src/bufio/bufio.go:220
		// _ = "end of CoverTab[1883]"
//line /snap/go/10455/src/bufio/bufio.go:220
		_go_fuzz_dep_.CoverTab[1884]++
							return 0, b.readErr()
//line /snap/go/10455/src/bufio/bufio.go:221
		// _ = "end of CoverTab[1884]"
	} else {
//line /snap/go/10455/src/bufio/bufio.go:222
		_go_fuzz_dep_.CoverTab[525710]++
//line /snap/go/10455/src/bufio/bufio.go:222
		_go_fuzz_dep_.CoverTab[1887]++
//line /snap/go/10455/src/bufio/bufio.go:222
		// _ = "end of CoverTab[1887]"
//line /snap/go/10455/src/bufio/bufio.go:222
	}
//line /snap/go/10455/src/bufio/bufio.go:222
	// _ = "end of CoverTab[1880]"
//line /snap/go/10455/src/bufio/bufio.go:222
	_go_fuzz_dep_.CoverTab[1881]++
						if b.r == b.w {
//line /snap/go/10455/src/bufio/bufio.go:223
		_go_fuzz_dep_.CoverTab[525713]++
//line /snap/go/10455/src/bufio/bufio.go:223
		_go_fuzz_dep_.CoverTab[1888]++
							if b.err != nil {
//line /snap/go/10455/src/bufio/bufio.go:224
			_go_fuzz_dep_.CoverTab[525715]++
//line /snap/go/10455/src/bufio/bufio.go:224
			_go_fuzz_dep_.CoverTab[1893]++
								return 0, b.readErr()
//line /snap/go/10455/src/bufio/bufio.go:225
			// _ = "end of CoverTab[1893]"
		} else {
//line /snap/go/10455/src/bufio/bufio.go:226
			_go_fuzz_dep_.CoverTab[525716]++
//line /snap/go/10455/src/bufio/bufio.go:226
			_go_fuzz_dep_.CoverTab[1894]++
//line /snap/go/10455/src/bufio/bufio.go:226
			// _ = "end of CoverTab[1894]"
//line /snap/go/10455/src/bufio/bufio.go:226
		}
//line /snap/go/10455/src/bufio/bufio.go:226
		// _ = "end of CoverTab[1888]"
//line /snap/go/10455/src/bufio/bufio.go:226
		_go_fuzz_dep_.CoverTab[1889]++
							if len(p) >= len(b.buf) {
//line /snap/go/10455/src/bufio/bufio.go:227
			_go_fuzz_dep_.CoverTab[525717]++
//line /snap/go/10455/src/bufio/bufio.go:227
			_go_fuzz_dep_.CoverTab[1895]++

//line /snap/go/10455/src/bufio/bufio.go:230
			n, b.err = b.rd.Read(p)
			if n < 0 {
//line /snap/go/10455/src/bufio/bufio.go:231
				_go_fuzz_dep_.CoverTab[525719]++
//line /snap/go/10455/src/bufio/bufio.go:231
				_go_fuzz_dep_.CoverTab[1898]++
									panic(errNegativeRead)
//line /snap/go/10455/src/bufio/bufio.go:232
				// _ = "end of CoverTab[1898]"
			} else {
//line /snap/go/10455/src/bufio/bufio.go:233
				_go_fuzz_dep_.CoverTab[525720]++
//line /snap/go/10455/src/bufio/bufio.go:233
				_go_fuzz_dep_.CoverTab[1899]++
//line /snap/go/10455/src/bufio/bufio.go:233
				// _ = "end of CoverTab[1899]"
//line /snap/go/10455/src/bufio/bufio.go:233
			}
//line /snap/go/10455/src/bufio/bufio.go:233
			// _ = "end of CoverTab[1895]"
//line /snap/go/10455/src/bufio/bufio.go:233
			_go_fuzz_dep_.CoverTab[1896]++
								if n > 0 {
//line /snap/go/10455/src/bufio/bufio.go:234
				_go_fuzz_dep_.CoverTab[525721]++
//line /snap/go/10455/src/bufio/bufio.go:234
				_go_fuzz_dep_.CoverTab[1900]++
									b.lastByte = int(p[n-1])
									b.lastRuneSize = -1
//line /snap/go/10455/src/bufio/bufio.go:236
				// _ = "end of CoverTab[1900]"
			} else {
//line /snap/go/10455/src/bufio/bufio.go:237
				_go_fuzz_dep_.CoverTab[525722]++
//line /snap/go/10455/src/bufio/bufio.go:237
				_go_fuzz_dep_.CoverTab[1901]++
//line /snap/go/10455/src/bufio/bufio.go:237
				// _ = "end of CoverTab[1901]"
//line /snap/go/10455/src/bufio/bufio.go:237
			}
//line /snap/go/10455/src/bufio/bufio.go:237
			// _ = "end of CoverTab[1896]"
//line /snap/go/10455/src/bufio/bufio.go:237
			_go_fuzz_dep_.CoverTab[1897]++
								return n, b.readErr()
//line /snap/go/10455/src/bufio/bufio.go:238
			// _ = "end of CoverTab[1897]"
		} else {
//line /snap/go/10455/src/bufio/bufio.go:239
			_go_fuzz_dep_.CoverTab[525718]++
//line /snap/go/10455/src/bufio/bufio.go:239
			_go_fuzz_dep_.CoverTab[1902]++
//line /snap/go/10455/src/bufio/bufio.go:239
			// _ = "end of CoverTab[1902]"
//line /snap/go/10455/src/bufio/bufio.go:239
		}
//line /snap/go/10455/src/bufio/bufio.go:239
		// _ = "end of CoverTab[1889]"
//line /snap/go/10455/src/bufio/bufio.go:239
		_go_fuzz_dep_.CoverTab[1890]++

//line /snap/go/10455/src/bufio/bufio.go:242
		b.r = 0
		b.w = 0
		n, b.err = b.rd.Read(b.buf)
		if n < 0 {
//line /snap/go/10455/src/bufio/bufio.go:245
			_go_fuzz_dep_.CoverTab[525723]++
//line /snap/go/10455/src/bufio/bufio.go:245
			_go_fuzz_dep_.CoverTab[1903]++
								panic(errNegativeRead)
//line /snap/go/10455/src/bufio/bufio.go:246
			// _ = "end of CoverTab[1903]"
		} else {
//line /snap/go/10455/src/bufio/bufio.go:247
			_go_fuzz_dep_.CoverTab[525724]++
//line /snap/go/10455/src/bufio/bufio.go:247
			_go_fuzz_dep_.CoverTab[1904]++
//line /snap/go/10455/src/bufio/bufio.go:247
			// _ = "end of CoverTab[1904]"
//line /snap/go/10455/src/bufio/bufio.go:247
		}
//line /snap/go/10455/src/bufio/bufio.go:247
		// _ = "end of CoverTab[1890]"
//line /snap/go/10455/src/bufio/bufio.go:247
		_go_fuzz_dep_.CoverTab[1891]++
							if n == 0 {
//line /snap/go/10455/src/bufio/bufio.go:248
			_go_fuzz_dep_.CoverTab[525725]++
//line /snap/go/10455/src/bufio/bufio.go:248
			_go_fuzz_dep_.CoverTab[1905]++
								return 0, b.readErr()
//line /snap/go/10455/src/bufio/bufio.go:249
			// _ = "end of CoverTab[1905]"
		} else {
//line /snap/go/10455/src/bufio/bufio.go:250
			_go_fuzz_dep_.CoverTab[525726]++
//line /snap/go/10455/src/bufio/bufio.go:250
			_go_fuzz_dep_.CoverTab[1906]++
//line /snap/go/10455/src/bufio/bufio.go:250
			// _ = "end of CoverTab[1906]"
//line /snap/go/10455/src/bufio/bufio.go:250
		}
//line /snap/go/10455/src/bufio/bufio.go:250
		// _ = "end of CoverTab[1891]"
//line /snap/go/10455/src/bufio/bufio.go:250
		_go_fuzz_dep_.CoverTab[1892]++
							b.w += n
//line /snap/go/10455/src/bufio/bufio.go:251
		// _ = "end of CoverTab[1892]"
	} else {
//line /snap/go/10455/src/bufio/bufio.go:252
		_go_fuzz_dep_.CoverTab[525714]++
//line /snap/go/10455/src/bufio/bufio.go:252
		_go_fuzz_dep_.CoverTab[1907]++
//line /snap/go/10455/src/bufio/bufio.go:252
		// _ = "end of CoverTab[1907]"
//line /snap/go/10455/src/bufio/bufio.go:252
	}
//line /snap/go/10455/src/bufio/bufio.go:252
	// _ = "end of CoverTab[1881]"
//line /snap/go/10455/src/bufio/bufio.go:252
	_go_fuzz_dep_.CoverTab[1882]++

//line /snap/go/10455/src/bufio/bufio.go:257
	n = copy(p, b.buf[b.r:b.w])
						b.r += n
						b.lastByte = int(b.buf[b.r-1])
						b.lastRuneSize = -1
						return n, nil
//line /snap/go/10455/src/bufio/bufio.go:261
	// _ = "end of CoverTab[1882]"
}

// ReadByte reads and returns a single byte.
//line /snap/go/10455/src/bufio/bufio.go:264
// If no byte is available, returns an error.
//line /snap/go/10455/src/bufio/bufio.go:266
func (b *Reader) ReadByte() (byte, error) {
//line /snap/go/10455/src/bufio/bufio.go:266
	_go_fuzz_dep_.CoverTab[1908]++
						b.lastRuneSize = -1
//line /snap/go/10455/src/bufio/bufio.go:267
	_go_fuzz_dep_.CoverTab[786557] = 0
						for b.r == b.w {
//line /snap/go/10455/src/bufio/bufio.go:268
		if _go_fuzz_dep_.CoverTab[786557] == 0 {
//line /snap/go/10455/src/bufio/bufio.go:268
			_go_fuzz_dep_.CoverTab[525853]++
//line /snap/go/10455/src/bufio/bufio.go:268
		} else {
//line /snap/go/10455/src/bufio/bufio.go:268
			_go_fuzz_dep_.CoverTab[525854]++
//line /snap/go/10455/src/bufio/bufio.go:268
		}
//line /snap/go/10455/src/bufio/bufio.go:268
		_go_fuzz_dep_.CoverTab[786557] = 1
//line /snap/go/10455/src/bufio/bufio.go:268
		_go_fuzz_dep_.CoverTab[1910]++
							if b.err != nil {
//line /snap/go/10455/src/bufio/bufio.go:269
			_go_fuzz_dep_.CoverTab[525727]++
//line /snap/go/10455/src/bufio/bufio.go:269
			_go_fuzz_dep_.CoverTab[1912]++
								return 0, b.readErr()
//line /snap/go/10455/src/bufio/bufio.go:270
			// _ = "end of CoverTab[1912]"
		} else {
//line /snap/go/10455/src/bufio/bufio.go:271
			_go_fuzz_dep_.CoverTab[525728]++
//line /snap/go/10455/src/bufio/bufio.go:271
			_go_fuzz_dep_.CoverTab[1913]++
//line /snap/go/10455/src/bufio/bufio.go:271
			// _ = "end of CoverTab[1913]"
//line /snap/go/10455/src/bufio/bufio.go:271
		}
//line /snap/go/10455/src/bufio/bufio.go:271
		// _ = "end of CoverTab[1910]"
//line /snap/go/10455/src/bufio/bufio.go:271
		_go_fuzz_dep_.CoverTab[1911]++
							b.fill()
//line /snap/go/10455/src/bufio/bufio.go:272
		// _ = "end of CoverTab[1911]"
	}
//line /snap/go/10455/src/bufio/bufio.go:273
	if _go_fuzz_dep_.CoverTab[786557] == 0 {
//line /snap/go/10455/src/bufio/bufio.go:273
		_go_fuzz_dep_.CoverTab[525855]++
//line /snap/go/10455/src/bufio/bufio.go:273
	} else {
//line /snap/go/10455/src/bufio/bufio.go:273
		_go_fuzz_dep_.CoverTab[525856]++
//line /snap/go/10455/src/bufio/bufio.go:273
	}
//line /snap/go/10455/src/bufio/bufio.go:273
	// _ = "end of CoverTab[1908]"
//line /snap/go/10455/src/bufio/bufio.go:273
	_go_fuzz_dep_.CoverTab[1909]++
						c := b.buf[b.r]
						b.r++
						b.lastByte = int(c)
						return c, nil
//line /snap/go/10455/src/bufio/bufio.go:277
	// _ = "end of CoverTab[1909]"
}

// UnreadByte unreads the last byte. Only the most recently read byte can be unread.
//line /snap/go/10455/src/bufio/bufio.go:280
//
//line /snap/go/10455/src/bufio/bufio.go:280
// UnreadByte returns an error if the most recent method called on the
//line /snap/go/10455/src/bufio/bufio.go:280
// Reader was not a read operation. Notably, Peek, Discard, and WriteTo are not
//line /snap/go/10455/src/bufio/bufio.go:280
// considered read operations.
//line /snap/go/10455/src/bufio/bufio.go:285
func (b *Reader) UnreadByte() error {
//line /snap/go/10455/src/bufio/bufio.go:285
	_go_fuzz_dep_.CoverTab[1914]++
						if b.lastByte < 0 || func() bool {
//line /snap/go/10455/src/bufio/bufio.go:286
		_go_fuzz_dep_.CoverTab[1917]++
//line /snap/go/10455/src/bufio/bufio.go:286
		return b.r == 0 && func() bool {
//line /snap/go/10455/src/bufio/bufio.go:286
			_go_fuzz_dep_.CoverTab[1918]++
//line /snap/go/10455/src/bufio/bufio.go:286
			return b.w > 0
//line /snap/go/10455/src/bufio/bufio.go:286
			// _ = "end of CoverTab[1918]"
//line /snap/go/10455/src/bufio/bufio.go:286
		}()
//line /snap/go/10455/src/bufio/bufio.go:286
		// _ = "end of CoverTab[1917]"
//line /snap/go/10455/src/bufio/bufio.go:286
	}() {
//line /snap/go/10455/src/bufio/bufio.go:286
		_go_fuzz_dep_.CoverTab[525729]++
//line /snap/go/10455/src/bufio/bufio.go:286
		_go_fuzz_dep_.CoverTab[1919]++
							return ErrInvalidUnreadByte
//line /snap/go/10455/src/bufio/bufio.go:287
		// _ = "end of CoverTab[1919]"
	} else {
//line /snap/go/10455/src/bufio/bufio.go:288
		_go_fuzz_dep_.CoverTab[525730]++
//line /snap/go/10455/src/bufio/bufio.go:288
		_go_fuzz_dep_.CoverTab[1920]++
//line /snap/go/10455/src/bufio/bufio.go:288
		// _ = "end of CoverTab[1920]"
//line /snap/go/10455/src/bufio/bufio.go:288
	}
//line /snap/go/10455/src/bufio/bufio.go:288
	// _ = "end of CoverTab[1914]"
//line /snap/go/10455/src/bufio/bufio.go:288
	_go_fuzz_dep_.CoverTab[1915]++

						if b.r > 0 {
//line /snap/go/10455/src/bufio/bufio.go:290
		_go_fuzz_dep_.CoverTab[525731]++
//line /snap/go/10455/src/bufio/bufio.go:290
		_go_fuzz_dep_.CoverTab[1921]++
							b.r--
//line /snap/go/10455/src/bufio/bufio.go:291
		// _ = "end of CoverTab[1921]"
	} else {
//line /snap/go/10455/src/bufio/bufio.go:292
		_go_fuzz_dep_.CoverTab[525732]++
//line /snap/go/10455/src/bufio/bufio.go:292
		_go_fuzz_dep_.CoverTab[1922]++

							b.w = 1
//line /snap/go/10455/src/bufio/bufio.go:294
		// _ = "end of CoverTab[1922]"
	}
//line /snap/go/10455/src/bufio/bufio.go:295
	// _ = "end of CoverTab[1915]"
//line /snap/go/10455/src/bufio/bufio.go:295
	_go_fuzz_dep_.CoverTab[1916]++
						b.buf[b.r] = byte(b.lastByte)
						b.lastByte = -1
						b.lastRuneSize = -1
						return nil
//line /snap/go/10455/src/bufio/bufio.go:299
	// _ = "end of CoverTab[1916]"
}

// ReadRune reads a single UTF-8 encoded Unicode character and returns the
//line /snap/go/10455/src/bufio/bufio.go:302
// rune and its size in bytes. If the encoded rune is invalid, it consumes one byte
//line /snap/go/10455/src/bufio/bufio.go:302
// and returns unicode.ReplacementChar (U+FFFD) with a size of 1.
//line /snap/go/10455/src/bufio/bufio.go:305
func (b *Reader) ReadRune() (r rune, size int, err error) {
//line /snap/go/10455/src/bufio/bufio.go:305
	_go_fuzz_dep_.CoverTab[1923]++
//line /snap/go/10455/src/bufio/bufio.go:305
	_go_fuzz_dep_.CoverTab[786558] = 0
						for b.r+utf8.UTFMax > b.w && func() bool {
//line /snap/go/10455/src/bufio/bufio.go:306
		_go_fuzz_dep_.CoverTab[1927]++
//line /snap/go/10455/src/bufio/bufio.go:306
		return !utf8.FullRune(b.buf[b.r:b.w])
//line /snap/go/10455/src/bufio/bufio.go:306
		// _ = "end of CoverTab[1927]"
//line /snap/go/10455/src/bufio/bufio.go:306
	}() && func() bool {
//line /snap/go/10455/src/bufio/bufio.go:306
		_go_fuzz_dep_.CoverTab[1928]++
//line /snap/go/10455/src/bufio/bufio.go:306
		return b.err == nil
//line /snap/go/10455/src/bufio/bufio.go:306
		// _ = "end of CoverTab[1928]"
//line /snap/go/10455/src/bufio/bufio.go:306
	}() && func() bool {
//line /snap/go/10455/src/bufio/bufio.go:306
		_go_fuzz_dep_.CoverTab[1929]++
//line /snap/go/10455/src/bufio/bufio.go:306
		return b.w-b.r < len(b.buf)
//line /snap/go/10455/src/bufio/bufio.go:306
		// _ = "end of CoverTab[1929]"
//line /snap/go/10455/src/bufio/bufio.go:306
	}() {
//line /snap/go/10455/src/bufio/bufio.go:306
		if _go_fuzz_dep_.CoverTab[786558] == 0 {
//line /snap/go/10455/src/bufio/bufio.go:306
			_go_fuzz_dep_.CoverTab[525857]++
//line /snap/go/10455/src/bufio/bufio.go:306
		} else {
//line /snap/go/10455/src/bufio/bufio.go:306
			_go_fuzz_dep_.CoverTab[525858]++
//line /snap/go/10455/src/bufio/bufio.go:306
		}
//line /snap/go/10455/src/bufio/bufio.go:306
		_go_fuzz_dep_.CoverTab[786558] = 1
//line /snap/go/10455/src/bufio/bufio.go:306
		_go_fuzz_dep_.CoverTab[1930]++
							b.fill()
//line /snap/go/10455/src/bufio/bufio.go:307
		// _ = "end of CoverTab[1930]"
	}
//line /snap/go/10455/src/bufio/bufio.go:308
	if _go_fuzz_dep_.CoverTab[786558] == 0 {
//line /snap/go/10455/src/bufio/bufio.go:308
		_go_fuzz_dep_.CoverTab[525859]++
//line /snap/go/10455/src/bufio/bufio.go:308
	} else {
//line /snap/go/10455/src/bufio/bufio.go:308
		_go_fuzz_dep_.CoverTab[525860]++
//line /snap/go/10455/src/bufio/bufio.go:308
	}
//line /snap/go/10455/src/bufio/bufio.go:308
	// _ = "end of CoverTab[1923]"
//line /snap/go/10455/src/bufio/bufio.go:308
	_go_fuzz_dep_.CoverTab[1924]++
						b.lastRuneSize = -1
						if b.r == b.w {
//line /snap/go/10455/src/bufio/bufio.go:310
		_go_fuzz_dep_.CoverTab[525733]++
//line /snap/go/10455/src/bufio/bufio.go:310
		_go_fuzz_dep_.CoverTab[1931]++
							return 0, 0, b.readErr()
//line /snap/go/10455/src/bufio/bufio.go:311
		// _ = "end of CoverTab[1931]"
	} else {
//line /snap/go/10455/src/bufio/bufio.go:312
		_go_fuzz_dep_.CoverTab[525734]++
//line /snap/go/10455/src/bufio/bufio.go:312
		_go_fuzz_dep_.CoverTab[1932]++
//line /snap/go/10455/src/bufio/bufio.go:312
		// _ = "end of CoverTab[1932]"
//line /snap/go/10455/src/bufio/bufio.go:312
	}
//line /snap/go/10455/src/bufio/bufio.go:312
	// _ = "end of CoverTab[1924]"
//line /snap/go/10455/src/bufio/bufio.go:312
	_go_fuzz_dep_.CoverTab[1925]++
						r, size = rune(b.buf[b.r]), 1
						if r >= utf8.RuneSelf {
//line /snap/go/10455/src/bufio/bufio.go:314
		_go_fuzz_dep_.CoverTab[525735]++
//line /snap/go/10455/src/bufio/bufio.go:314
		_go_fuzz_dep_.CoverTab[1933]++
							r, size = utf8.DecodeRune(b.buf[b.r:b.w])
//line /snap/go/10455/src/bufio/bufio.go:315
		// _ = "end of CoverTab[1933]"
	} else {
//line /snap/go/10455/src/bufio/bufio.go:316
		_go_fuzz_dep_.CoverTab[525736]++
//line /snap/go/10455/src/bufio/bufio.go:316
		_go_fuzz_dep_.CoverTab[1934]++
//line /snap/go/10455/src/bufio/bufio.go:316
		// _ = "end of CoverTab[1934]"
//line /snap/go/10455/src/bufio/bufio.go:316
	}
//line /snap/go/10455/src/bufio/bufio.go:316
	// _ = "end of CoverTab[1925]"
//line /snap/go/10455/src/bufio/bufio.go:316
	_go_fuzz_dep_.CoverTab[1926]++
						b.r += size
						b.lastByte = int(b.buf[b.r-1])
						b.lastRuneSize = size
						return r, size, nil
//line /snap/go/10455/src/bufio/bufio.go:320
	// _ = "end of CoverTab[1926]"
}

// UnreadRune unreads the last rune. If the most recent method called on
//line /snap/go/10455/src/bufio/bufio.go:323
// the Reader was not a ReadRune, UnreadRune returns an error. (In this
//line /snap/go/10455/src/bufio/bufio.go:323
// regard it is stricter than UnreadByte, which will unread the last byte
//line /snap/go/10455/src/bufio/bufio.go:323
// from any read operation.)
//line /snap/go/10455/src/bufio/bufio.go:327
func (b *Reader) UnreadRune() error {
//line /snap/go/10455/src/bufio/bufio.go:327
	_go_fuzz_dep_.CoverTab[1935]++
						if b.lastRuneSize < 0 || func() bool {
//line /snap/go/10455/src/bufio/bufio.go:328
		_go_fuzz_dep_.CoverTab[1937]++
//line /snap/go/10455/src/bufio/bufio.go:328
		return b.r < b.lastRuneSize
//line /snap/go/10455/src/bufio/bufio.go:328
		// _ = "end of CoverTab[1937]"
//line /snap/go/10455/src/bufio/bufio.go:328
	}() {
//line /snap/go/10455/src/bufio/bufio.go:328
		_go_fuzz_dep_.CoverTab[525737]++
//line /snap/go/10455/src/bufio/bufio.go:328
		_go_fuzz_dep_.CoverTab[1938]++
							return ErrInvalidUnreadRune
//line /snap/go/10455/src/bufio/bufio.go:329
		// _ = "end of CoverTab[1938]"
	} else {
//line /snap/go/10455/src/bufio/bufio.go:330
		_go_fuzz_dep_.CoverTab[525738]++
//line /snap/go/10455/src/bufio/bufio.go:330
		_go_fuzz_dep_.CoverTab[1939]++
//line /snap/go/10455/src/bufio/bufio.go:330
		// _ = "end of CoverTab[1939]"
//line /snap/go/10455/src/bufio/bufio.go:330
	}
//line /snap/go/10455/src/bufio/bufio.go:330
	// _ = "end of CoverTab[1935]"
//line /snap/go/10455/src/bufio/bufio.go:330
	_go_fuzz_dep_.CoverTab[1936]++
						b.r -= b.lastRuneSize
						b.lastByte = -1
						b.lastRuneSize = -1
						return nil
//line /snap/go/10455/src/bufio/bufio.go:334
	// _ = "end of CoverTab[1936]"
}

// Buffered returns the number of bytes that can be read from the current buffer.
func (b *Reader) Buffered() int {
//line /snap/go/10455/src/bufio/bufio.go:338
	_go_fuzz_dep_.CoverTab[1940]++
//line /snap/go/10455/src/bufio/bufio.go:338
	return b.w - b.r
//line /snap/go/10455/src/bufio/bufio.go:338
	// _ = "end of CoverTab[1940]"
//line /snap/go/10455/src/bufio/bufio.go:338
}

// ReadSlice reads until the first occurrence of delim in the input,
//line /snap/go/10455/src/bufio/bufio.go:340
// returning a slice pointing at the bytes in the buffer.
//line /snap/go/10455/src/bufio/bufio.go:340
// The bytes stop being valid at the next read.
//line /snap/go/10455/src/bufio/bufio.go:340
// If ReadSlice encounters an error before finding a delimiter,
//line /snap/go/10455/src/bufio/bufio.go:340
// it returns all the data in the buffer and the error itself (often io.EOF).
//line /snap/go/10455/src/bufio/bufio.go:340
// ReadSlice fails with error ErrBufferFull if the buffer fills without a delim.
//line /snap/go/10455/src/bufio/bufio.go:340
// Because the data returned from ReadSlice will be overwritten
//line /snap/go/10455/src/bufio/bufio.go:340
// by the next I/O operation, most clients should use
//line /snap/go/10455/src/bufio/bufio.go:340
// ReadBytes or ReadString instead.
//line /snap/go/10455/src/bufio/bufio.go:340
// ReadSlice returns err != nil if and only if line does not end in delim.
//line /snap/go/10455/src/bufio/bufio.go:350
func (b *Reader) ReadSlice(delim byte) (line []byte, err error) {
//line /snap/go/10455/src/bufio/bufio.go:350
	_go_fuzz_dep_.CoverTab[1941]++
						s := 0
//line /snap/go/10455/src/bufio/bufio.go:351
	_go_fuzz_dep_.CoverTab[786559] = 0
						for {
//line /snap/go/10455/src/bufio/bufio.go:352
		if _go_fuzz_dep_.CoverTab[786559] == 0 {
//line /snap/go/10455/src/bufio/bufio.go:352
			_go_fuzz_dep_.CoverTab[525861]++
//line /snap/go/10455/src/bufio/bufio.go:352
		} else {
//line /snap/go/10455/src/bufio/bufio.go:352
			_go_fuzz_dep_.CoverTab[525862]++
//line /snap/go/10455/src/bufio/bufio.go:352
		}
//line /snap/go/10455/src/bufio/bufio.go:352
		_go_fuzz_dep_.CoverTab[786559] = 1
//line /snap/go/10455/src/bufio/bufio.go:352
		_go_fuzz_dep_.CoverTab[1944]++

							if i := bytes.IndexByte(b.buf[b.r+s:b.w], delim); i >= 0 {
//line /snap/go/10455/src/bufio/bufio.go:354
			_go_fuzz_dep_.CoverTab[525739]++
//line /snap/go/10455/src/bufio/bufio.go:354
			_go_fuzz_dep_.CoverTab[1948]++
								i += s
								line = b.buf[b.r : b.r+i+1]
								b.r += i + 1
								break
//line /snap/go/10455/src/bufio/bufio.go:358
			// _ = "end of CoverTab[1948]"
		} else {
//line /snap/go/10455/src/bufio/bufio.go:359
			_go_fuzz_dep_.CoverTab[525740]++
//line /snap/go/10455/src/bufio/bufio.go:359
			_go_fuzz_dep_.CoverTab[1949]++
//line /snap/go/10455/src/bufio/bufio.go:359
			// _ = "end of CoverTab[1949]"
//line /snap/go/10455/src/bufio/bufio.go:359
		}
//line /snap/go/10455/src/bufio/bufio.go:359
		// _ = "end of CoverTab[1944]"
//line /snap/go/10455/src/bufio/bufio.go:359
		_go_fuzz_dep_.CoverTab[1945]++

//line /snap/go/10455/src/bufio/bufio.go:362
		if b.err != nil {
//line /snap/go/10455/src/bufio/bufio.go:362
			_go_fuzz_dep_.CoverTab[525741]++
//line /snap/go/10455/src/bufio/bufio.go:362
			_go_fuzz_dep_.CoverTab[1950]++
								line = b.buf[b.r:b.w]
								b.r = b.w
								err = b.readErr()
								break
//line /snap/go/10455/src/bufio/bufio.go:366
			// _ = "end of CoverTab[1950]"
		} else {
//line /snap/go/10455/src/bufio/bufio.go:367
			_go_fuzz_dep_.CoverTab[525742]++
//line /snap/go/10455/src/bufio/bufio.go:367
			_go_fuzz_dep_.CoverTab[1951]++
//line /snap/go/10455/src/bufio/bufio.go:367
			// _ = "end of CoverTab[1951]"
//line /snap/go/10455/src/bufio/bufio.go:367
		}
//line /snap/go/10455/src/bufio/bufio.go:367
		// _ = "end of CoverTab[1945]"
//line /snap/go/10455/src/bufio/bufio.go:367
		_go_fuzz_dep_.CoverTab[1946]++

//line /snap/go/10455/src/bufio/bufio.go:370
		if b.Buffered() >= len(b.buf) {
//line /snap/go/10455/src/bufio/bufio.go:370
			_go_fuzz_dep_.CoverTab[525743]++
//line /snap/go/10455/src/bufio/bufio.go:370
			_go_fuzz_dep_.CoverTab[1952]++
								b.r = b.w
								line = b.buf
								err = ErrBufferFull
								break
//line /snap/go/10455/src/bufio/bufio.go:374
			// _ = "end of CoverTab[1952]"
		} else {
//line /snap/go/10455/src/bufio/bufio.go:375
			_go_fuzz_dep_.CoverTab[525744]++
//line /snap/go/10455/src/bufio/bufio.go:375
			_go_fuzz_dep_.CoverTab[1953]++
//line /snap/go/10455/src/bufio/bufio.go:375
			// _ = "end of CoverTab[1953]"
//line /snap/go/10455/src/bufio/bufio.go:375
		}
//line /snap/go/10455/src/bufio/bufio.go:375
		// _ = "end of CoverTab[1946]"
//line /snap/go/10455/src/bufio/bufio.go:375
		_go_fuzz_dep_.CoverTab[1947]++

							s = b.w - b.r

							b.fill()
//line /snap/go/10455/src/bufio/bufio.go:379
		// _ = "end of CoverTab[1947]"
	}
//line /snap/go/10455/src/bufio/bufio.go:380
	// _ = "end of CoverTab[1941]"
//line /snap/go/10455/src/bufio/bufio.go:380
	_go_fuzz_dep_.CoverTab[1942]++

//line /snap/go/10455/src/bufio/bufio.go:383
	if i := len(line) - 1; i >= 0 {
//line /snap/go/10455/src/bufio/bufio.go:383
		_go_fuzz_dep_.CoverTab[525745]++
//line /snap/go/10455/src/bufio/bufio.go:383
		_go_fuzz_dep_.CoverTab[1954]++
							b.lastByte = int(line[i])
							b.lastRuneSize = -1
//line /snap/go/10455/src/bufio/bufio.go:385
		// _ = "end of CoverTab[1954]"
	} else {
//line /snap/go/10455/src/bufio/bufio.go:386
		_go_fuzz_dep_.CoverTab[525746]++
//line /snap/go/10455/src/bufio/bufio.go:386
		_go_fuzz_dep_.CoverTab[1955]++
//line /snap/go/10455/src/bufio/bufio.go:386
		// _ = "end of CoverTab[1955]"
//line /snap/go/10455/src/bufio/bufio.go:386
	}
//line /snap/go/10455/src/bufio/bufio.go:386
	// _ = "end of CoverTab[1942]"
//line /snap/go/10455/src/bufio/bufio.go:386
	_go_fuzz_dep_.CoverTab[1943]++

						return
//line /snap/go/10455/src/bufio/bufio.go:388
	// _ = "end of CoverTab[1943]"
}

// ReadLine is a low-level line-reading primitive. Most callers should use
//line /snap/go/10455/src/bufio/bufio.go:391
// ReadBytes('\n') or ReadString('\n') instead or use a Scanner.
//line /snap/go/10455/src/bufio/bufio.go:391
//
//line /snap/go/10455/src/bufio/bufio.go:391
// ReadLine tries to return a single line, not including the end-of-line bytes.
//line /snap/go/10455/src/bufio/bufio.go:391
// If the line was too long for the buffer then isPrefix is set and the
//line /snap/go/10455/src/bufio/bufio.go:391
// beginning of the line is returned. The rest of the line will be returned
//line /snap/go/10455/src/bufio/bufio.go:391
// from future calls. isPrefix will be false when returning the last fragment
//line /snap/go/10455/src/bufio/bufio.go:391
// of the line. The returned buffer is only valid until the next call to
//line /snap/go/10455/src/bufio/bufio.go:391
// ReadLine. ReadLine either returns a non-nil line or it returns an error,
//line /snap/go/10455/src/bufio/bufio.go:391
// never both.
//line /snap/go/10455/src/bufio/bufio.go:391
//
//line /snap/go/10455/src/bufio/bufio.go:391
// The text returned from ReadLine does not include the line end ("\r\n" or "\n").
//line /snap/go/10455/src/bufio/bufio.go:391
// No indication or error is given if the input ends without a final line end.
//line /snap/go/10455/src/bufio/bufio.go:391
// Calling UnreadByte after ReadLine will always unread the last byte read
//line /snap/go/10455/src/bufio/bufio.go:391
// (possibly a character belonging to the line end) even if that byte is not
//line /snap/go/10455/src/bufio/bufio.go:391
// part of the line returned by ReadLine.
//line /snap/go/10455/src/bufio/bufio.go:407
func (b *Reader) ReadLine() (line []byte, isPrefix bool, err error) {
//line /snap/go/10455/src/bufio/bufio.go:407
	_go_fuzz_dep_.CoverTab[1956]++
						line, err = b.ReadSlice('\n')
						if err == ErrBufferFull {
//line /snap/go/10455/src/bufio/bufio.go:409
		_go_fuzz_dep_.CoverTab[525747]++
//line /snap/go/10455/src/bufio/bufio.go:409
		_go_fuzz_dep_.CoverTab[1960]++

							if len(line) > 0 && func() bool {
//line /snap/go/10455/src/bufio/bufio.go:411
			_go_fuzz_dep_.CoverTab[1962]++
//line /snap/go/10455/src/bufio/bufio.go:411
			return line[len(line)-1] == '\r'
//line /snap/go/10455/src/bufio/bufio.go:411
			// _ = "end of CoverTab[1962]"
//line /snap/go/10455/src/bufio/bufio.go:411
		}() {
//line /snap/go/10455/src/bufio/bufio.go:411
			_go_fuzz_dep_.CoverTab[525749]++
//line /snap/go/10455/src/bufio/bufio.go:411
			_go_fuzz_dep_.CoverTab[1963]++

//line /snap/go/10455/src/bufio/bufio.go:414
			if b.r == 0 {
//line /snap/go/10455/src/bufio/bufio.go:414
				_go_fuzz_dep_.CoverTab[525751]++
//line /snap/go/10455/src/bufio/bufio.go:414
				_go_fuzz_dep_.CoverTab[1965]++

									panic("bufio: tried to rewind past start of buffer")
//line /snap/go/10455/src/bufio/bufio.go:416
				// _ = "end of CoverTab[1965]"
			} else {
//line /snap/go/10455/src/bufio/bufio.go:417
				_go_fuzz_dep_.CoverTab[525752]++
//line /snap/go/10455/src/bufio/bufio.go:417
				_go_fuzz_dep_.CoverTab[1966]++
//line /snap/go/10455/src/bufio/bufio.go:417
				// _ = "end of CoverTab[1966]"
//line /snap/go/10455/src/bufio/bufio.go:417
			}
//line /snap/go/10455/src/bufio/bufio.go:417
			// _ = "end of CoverTab[1963]"
//line /snap/go/10455/src/bufio/bufio.go:417
			_go_fuzz_dep_.CoverTab[1964]++
								b.r--
								line = line[:len(line)-1]
//line /snap/go/10455/src/bufio/bufio.go:419
			// _ = "end of CoverTab[1964]"
		} else {
//line /snap/go/10455/src/bufio/bufio.go:420
			_go_fuzz_dep_.CoverTab[525750]++
//line /snap/go/10455/src/bufio/bufio.go:420
			_go_fuzz_dep_.CoverTab[1967]++
//line /snap/go/10455/src/bufio/bufio.go:420
			// _ = "end of CoverTab[1967]"
//line /snap/go/10455/src/bufio/bufio.go:420
		}
//line /snap/go/10455/src/bufio/bufio.go:420
		// _ = "end of CoverTab[1960]"
//line /snap/go/10455/src/bufio/bufio.go:420
		_go_fuzz_dep_.CoverTab[1961]++
							return line, true, nil
//line /snap/go/10455/src/bufio/bufio.go:421
		// _ = "end of CoverTab[1961]"
	} else {
//line /snap/go/10455/src/bufio/bufio.go:422
		_go_fuzz_dep_.CoverTab[525748]++
//line /snap/go/10455/src/bufio/bufio.go:422
		_go_fuzz_dep_.CoverTab[1968]++
//line /snap/go/10455/src/bufio/bufio.go:422
		// _ = "end of CoverTab[1968]"
//line /snap/go/10455/src/bufio/bufio.go:422
	}
//line /snap/go/10455/src/bufio/bufio.go:422
	// _ = "end of CoverTab[1956]"
//line /snap/go/10455/src/bufio/bufio.go:422
	_go_fuzz_dep_.CoverTab[1957]++

						if len(line) == 0 {
//line /snap/go/10455/src/bufio/bufio.go:424
		_go_fuzz_dep_.CoverTab[525753]++
//line /snap/go/10455/src/bufio/bufio.go:424
		_go_fuzz_dep_.CoverTab[1969]++
							if err != nil {
//line /snap/go/10455/src/bufio/bufio.go:425
			_go_fuzz_dep_.CoverTab[525755]++
//line /snap/go/10455/src/bufio/bufio.go:425
			_go_fuzz_dep_.CoverTab[1971]++
								line = nil
//line /snap/go/10455/src/bufio/bufio.go:426
			// _ = "end of CoverTab[1971]"
		} else {
//line /snap/go/10455/src/bufio/bufio.go:427
			_go_fuzz_dep_.CoverTab[525756]++
//line /snap/go/10455/src/bufio/bufio.go:427
			_go_fuzz_dep_.CoverTab[1972]++
//line /snap/go/10455/src/bufio/bufio.go:427
			// _ = "end of CoverTab[1972]"
//line /snap/go/10455/src/bufio/bufio.go:427
		}
//line /snap/go/10455/src/bufio/bufio.go:427
		// _ = "end of CoverTab[1969]"
//line /snap/go/10455/src/bufio/bufio.go:427
		_go_fuzz_dep_.CoverTab[1970]++
							return
//line /snap/go/10455/src/bufio/bufio.go:428
		// _ = "end of CoverTab[1970]"
	} else {
//line /snap/go/10455/src/bufio/bufio.go:429
		_go_fuzz_dep_.CoverTab[525754]++
//line /snap/go/10455/src/bufio/bufio.go:429
		_go_fuzz_dep_.CoverTab[1973]++
//line /snap/go/10455/src/bufio/bufio.go:429
		// _ = "end of CoverTab[1973]"
//line /snap/go/10455/src/bufio/bufio.go:429
	}
//line /snap/go/10455/src/bufio/bufio.go:429
	// _ = "end of CoverTab[1957]"
//line /snap/go/10455/src/bufio/bufio.go:429
	_go_fuzz_dep_.CoverTab[1958]++
						err = nil

						if line[len(line)-1] == '\n' {
//line /snap/go/10455/src/bufio/bufio.go:432
		_go_fuzz_dep_.CoverTab[525757]++
//line /snap/go/10455/src/bufio/bufio.go:432
		_go_fuzz_dep_.CoverTab[1974]++
							drop := 1
							if len(line) > 1 && func() bool {
//line /snap/go/10455/src/bufio/bufio.go:434
			_go_fuzz_dep_.CoverTab[1976]++
//line /snap/go/10455/src/bufio/bufio.go:434
			return line[len(line)-2] == '\r'
//line /snap/go/10455/src/bufio/bufio.go:434
			// _ = "end of CoverTab[1976]"
//line /snap/go/10455/src/bufio/bufio.go:434
		}() {
//line /snap/go/10455/src/bufio/bufio.go:434
			_go_fuzz_dep_.CoverTab[525759]++
//line /snap/go/10455/src/bufio/bufio.go:434
			_go_fuzz_dep_.CoverTab[1977]++
								drop = 2
//line /snap/go/10455/src/bufio/bufio.go:435
			// _ = "end of CoverTab[1977]"
		} else {
//line /snap/go/10455/src/bufio/bufio.go:436
			_go_fuzz_dep_.CoverTab[525760]++
//line /snap/go/10455/src/bufio/bufio.go:436
			_go_fuzz_dep_.CoverTab[1978]++
//line /snap/go/10455/src/bufio/bufio.go:436
			// _ = "end of CoverTab[1978]"
//line /snap/go/10455/src/bufio/bufio.go:436
		}
//line /snap/go/10455/src/bufio/bufio.go:436
		// _ = "end of CoverTab[1974]"
//line /snap/go/10455/src/bufio/bufio.go:436
		_go_fuzz_dep_.CoverTab[1975]++
							line = line[:len(line)-drop]
//line /snap/go/10455/src/bufio/bufio.go:437
		// _ = "end of CoverTab[1975]"
	} else {
//line /snap/go/10455/src/bufio/bufio.go:438
		_go_fuzz_dep_.CoverTab[525758]++
//line /snap/go/10455/src/bufio/bufio.go:438
		_go_fuzz_dep_.CoverTab[1979]++
//line /snap/go/10455/src/bufio/bufio.go:438
		// _ = "end of CoverTab[1979]"
//line /snap/go/10455/src/bufio/bufio.go:438
	}
//line /snap/go/10455/src/bufio/bufio.go:438
	// _ = "end of CoverTab[1958]"
//line /snap/go/10455/src/bufio/bufio.go:438
	_go_fuzz_dep_.CoverTab[1959]++
						return
//line /snap/go/10455/src/bufio/bufio.go:439
	// _ = "end of CoverTab[1959]"
}

// collectFragments reads until the first occurrence of delim in the input. It
//line /snap/go/10455/src/bufio/bufio.go:442
// returns (slice of full buffers, remaining bytes before delim, total number
//line /snap/go/10455/src/bufio/bufio.go:442
// of bytes in the combined first two elements, error).
//line /snap/go/10455/src/bufio/bufio.go:442
// The complete result is equal to
//line /snap/go/10455/src/bufio/bufio.go:442
// `bytes.Join(append(fullBuffers, finalFragment), nil)`, which has a
//line /snap/go/10455/src/bufio/bufio.go:442
// length of `totalLen`. The result is structured in this way to allow callers
//line /snap/go/10455/src/bufio/bufio.go:442
// to minimize allocations and copies.
//line /snap/go/10455/src/bufio/bufio.go:449
func (b *Reader) collectFragments(delim byte) (fullBuffers [][]byte, finalFragment []byte, totalLen int, err error) {
//line /snap/go/10455/src/bufio/bufio.go:449
	_go_fuzz_dep_.CoverTab[1980]++
						var frag []byte
//line /snap/go/10455/src/bufio/bufio.go:450
	_go_fuzz_dep_.CoverTab[786560] = 0

						for {
//line /snap/go/10455/src/bufio/bufio.go:452
		if _go_fuzz_dep_.CoverTab[786560] == 0 {
//line /snap/go/10455/src/bufio/bufio.go:452
			_go_fuzz_dep_.CoverTab[525865]++
//line /snap/go/10455/src/bufio/bufio.go:452
		} else {
//line /snap/go/10455/src/bufio/bufio.go:452
			_go_fuzz_dep_.CoverTab[525866]++
//line /snap/go/10455/src/bufio/bufio.go:452
		}
//line /snap/go/10455/src/bufio/bufio.go:452
		_go_fuzz_dep_.CoverTab[786560] = 1
//line /snap/go/10455/src/bufio/bufio.go:452
		_go_fuzz_dep_.CoverTab[1982]++
							var e error
							frag, e = b.ReadSlice(delim)
							if e == nil {
//line /snap/go/10455/src/bufio/bufio.go:455
			_go_fuzz_dep_.CoverTab[525761]++
//line /snap/go/10455/src/bufio/bufio.go:455
			_go_fuzz_dep_.CoverTab[1985]++
								break
//line /snap/go/10455/src/bufio/bufio.go:456
			// _ = "end of CoverTab[1985]"
		} else {
//line /snap/go/10455/src/bufio/bufio.go:457
			_go_fuzz_dep_.CoverTab[525762]++
//line /snap/go/10455/src/bufio/bufio.go:457
			_go_fuzz_dep_.CoverTab[1986]++
//line /snap/go/10455/src/bufio/bufio.go:457
			// _ = "end of CoverTab[1986]"
//line /snap/go/10455/src/bufio/bufio.go:457
		}
//line /snap/go/10455/src/bufio/bufio.go:457
		// _ = "end of CoverTab[1982]"
//line /snap/go/10455/src/bufio/bufio.go:457
		_go_fuzz_dep_.CoverTab[1983]++
							if e != ErrBufferFull {
//line /snap/go/10455/src/bufio/bufio.go:458
			_go_fuzz_dep_.CoverTab[525763]++
//line /snap/go/10455/src/bufio/bufio.go:458
			_go_fuzz_dep_.CoverTab[1987]++
								err = e
								break
//line /snap/go/10455/src/bufio/bufio.go:460
			// _ = "end of CoverTab[1987]"
		} else {
//line /snap/go/10455/src/bufio/bufio.go:461
			_go_fuzz_dep_.CoverTab[525764]++
//line /snap/go/10455/src/bufio/bufio.go:461
			_go_fuzz_dep_.CoverTab[1988]++
//line /snap/go/10455/src/bufio/bufio.go:461
			// _ = "end of CoverTab[1988]"
//line /snap/go/10455/src/bufio/bufio.go:461
		}
//line /snap/go/10455/src/bufio/bufio.go:461
		// _ = "end of CoverTab[1983]"
//line /snap/go/10455/src/bufio/bufio.go:461
		_go_fuzz_dep_.CoverTab[1984]++

//line /snap/go/10455/src/bufio/bufio.go:464
		buf := bytes.Clone(frag)
							fullBuffers = append(fullBuffers, buf)
							totalLen += len(buf)
//line /snap/go/10455/src/bufio/bufio.go:466
		// _ = "end of CoverTab[1984]"
	}
//line /snap/go/10455/src/bufio/bufio.go:467
	// _ = "end of CoverTab[1980]"
//line /snap/go/10455/src/bufio/bufio.go:467
	_go_fuzz_dep_.CoverTab[1981]++

						totalLen += len(frag)
						return fullBuffers, frag, totalLen, err
//line /snap/go/10455/src/bufio/bufio.go:470
	// _ = "end of CoverTab[1981]"
}

// ReadBytes reads until the first occurrence of delim in the input,
//line /snap/go/10455/src/bufio/bufio.go:473
// returning a slice containing the data up to and including the delimiter.
//line /snap/go/10455/src/bufio/bufio.go:473
// If ReadBytes encounters an error before finding a delimiter,
//line /snap/go/10455/src/bufio/bufio.go:473
// it returns the data read before the error and the error itself (often io.EOF).
//line /snap/go/10455/src/bufio/bufio.go:473
// ReadBytes returns err != nil if and only if the returned data does not end in
//line /snap/go/10455/src/bufio/bufio.go:473
// delim.
//line /snap/go/10455/src/bufio/bufio.go:473
// For simple uses, a Scanner may be more convenient.
//line /snap/go/10455/src/bufio/bufio.go:480
func (b *Reader) ReadBytes(delim byte) ([]byte, error) {
//line /snap/go/10455/src/bufio/bufio.go:480
	_go_fuzz_dep_.CoverTab[1989]++
						full, frag, n, err := b.collectFragments(delim)

						buf := make([]byte, n)
						n = 0
//line /snap/go/10455/src/bufio/bufio.go:484
	_go_fuzz_dep_.CoverTab[786561] = 0

						for i := range full {
//line /snap/go/10455/src/bufio/bufio.go:486
		if _go_fuzz_dep_.CoverTab[786561] == 0 {
//line /snap/go/10455/src/bufio/bufio.go:486
			_go_fuzz_dep_.CoverTab[525869]++
//line /snap/go/10455/src/bufio/bufio.go:486
		} else {
//line /snap/go/10455/src/bufio/bufio.go:486
			_go_fuzz_dep_.CoverTab[525870]++
//line /snap/go/10455/src/bufio/bufio.go:486
		}
//line /snap/go/10455/src/bufio/bufio.go:486
		_go_fuzz_dep_.CoverTab[786561] = 1
//line /snap/go/10455/src/bufio/bufio.go:486
		_go_fuzz_dep_.CoverTab[1991]++
							n += copy(buf[n:], full[i])
//line /snap/go/10455/src/bufio/bufio.go:487
		// _ = "end of CoverTab[1991]"
	}
//line /snap/go/10455/src/bufio/bufio.go:488
	if _go_fuzz_dep_.CoverTab[786561] == 0 {
//line /snap/go/10455/src/bufio/bufio.go:488
		_go_fuzz_dep_.CoverTab[525871]++
//line /snap/go/10455/src/bufio/bufio.go:488
	} else {
//line /snap/go/10455/src/bufio/bufio.go:488
		_go_fuzz_dep_.CoverTab[525872]++
//line /snap/go/10455/src/bufio/bufio.go:488
	}
//line /snap/go/10455/src/bufio/bufio.go:488
	// _ = "end of CoverTab[1989]"
//line /snap/go/10455/src/bufio/bufio.go:488
	_go_fuzz_dep_.CoverTab[1990]++
						copy(buf[n:], frag)
						return buf, err
//line /snap/go/10455/src/bufio/bufio.go:490
	// _ = "end of CoverTab[1990]"
}

// ReadString reads until the first occurrence of delim in the input,
//line /snap/go/10455/src/bufio/bufio.go:493
// returning a string containing the data up to and including the delimiter.
//line /snap/go/10455/src/bufio/bufio.go:493
// If ReadString encounters an error before finding a delimiter,
//line /snap/go/10455/src/bufio/bufio.go:493
// it returns the data read before the error and the error itself (often io.EOF).
//line /snap/go/10455/src/bufio/bufio.go:493
// ReadString returns err != nil if and only if the returned data does not end in
//line /snap/go/10455/src/bufio/bufio.go:493
// delim.
//line /snap/go/10455/src/bufio/bufio.go:493
// For simple uses, a Scanner may be more convenient.
//line /snap/go/10455/src/bufio/bufio.go:500
func (b *Reader) ReadString(delim byte) (string, error) {
//line /snap/go/10455/src/bufio/bufio.go:500
	_go_fuzz_dep_.CoverTab[1992]++
						full, frag, n, err := b.collectFragments(delim)
						// Allocate new buffer to hold the full pieces and the fragment.
						var buf strings.Builder
						buf.Grow(n)
//line /snap/go/10455/src/bufio/bufio.go:504
	_go_fuzz_dep_.CoverTab[786562] = 0

						for _, fb := range full {
//line /snap/go/10455/src/bufio/bufio.go:506
		if _go_fuzz_dep_.CoverTab[786562] == 0 {
//line /snap/go/10455/src/bufio/bufio.go:506
			_go_fuzz_dep_.CoverTab[525873]++
//line /snap/go/10455/src/bufio/bufio.go:506
		} else {
//line /snap/go/10455/src/bufio/bufio.go:506
			_go_fuzz_dep_.CoverTab[525874]++
//line /snap/go/10455/src/bufio/bufio.go:506
		}
//line /snap/go/10455/src/bufio/bufio.go:506
		_go_fuzz_dep_.CoverTab[786562] = 1
//line /snap/go/10455/src/bufio/bufio.go:506
		_go_fuzz_dep_.CoverTab[1994]++
							buf.Write(fb)
//line /snap/go/10455/src/bufio/bufio.go:507
		// _ = "end of CoverTab[1994]"
	}
//line /snap/go/10455/src/bufio/bufio.go:508
	if _go_fuzz_dep_.CoverTab[786562] == 0 {
//line /snap/go/10455/src/bufio/bufio.go:508
		_go_fuzz_dep_.CoverTab[525875]++
//line /snap/go/10455/src/bufio/bufio.go:508
	} else {
//line /snap/go/10455/src/bufio/bufio.go:508
		_go_fuzz_dep_.CoverTab[525876]++
//line /snap/go/10455/src/bufio/bufio.go:508
	}
//line /snap/go/10455/src/bufio/bufio.go:508
	// _ = "end of CoverTab[1992]"
//line /snap/go/10455/src/bufio/bufio.go:508
	_go_fuzz_dep_.CoverTab[1993]++
						buf.Write(frag)
						return buf.String(), err
//line /snap/go/10455/src/bufio/bufio.go:510
	// _ = "end of CoverTab[1993]"
}

// WriteTo implements io.WriterTo.
//line /snap/go/10455/src/bufio/bufio.go:513
// This may make multiple calls to the Read method of the underlying Reader.
//line /snap/go/10455/src/bufio/bufio.go:513
// If the underlying reader supports the WriteTo method,
//line /snap/go/10455/src/bufio/bufio.go:513
// this calls the underlying WriteTo without buffering.
//line /snap/go/10455/src/bufio/bufio.go:517
func (b *Reader) WriteTo(w io.Writer) (n int64, err error) {
//line /snap/go/10455/src/bufio/bufio.go:517
	_go_fuzz_dep_.CoverTab[1995]++
						b.lastByte = -1
						b.lastRuneSize = -1

						n, err = b.writeBuf(w)
						if err != nil {
//line /snap/go/10455/src/bufio/bufio.go:522
		_go_fuzz_dep_.CoverTab[525765]++
//line /snap/go/10455/src/bufio/bufio.go:522
		_go_fuzz_dep_.CoverTab[2002]++
							return
//line /snap/go/10455/src/bufio/bufio.go:523
		// _ = "end of CoverTab[2002]"
	} else {
//line /snap/go/10455/src/bufio/bufio.go:524
		_go_fuzz_dep_.CoverTab[525766]++
//line /snap/go/10455/src/bufio/bufio.go:524
		_go_fuzz_dep_.CoverTab[2003]++
//line /snap/go/10455/src/bufio/bufio.go:524
		// _ = "end of CoverTab[2003]"
//line /snap/go/10455/src/bufio/bufio.go:524
	}
//line /snap/go/10455/src/bufio/bufio.go:524
	// _ = "end of CoverTab[1995]"
//line /snap/go/10455/src/bufio/bufio.go:524
	_go_fuzz_dep_.CoverTab[1996]++

						if r, ok := b.rd.(io.WriterTo); ok {
//line /snap/go/10455/src/bufio/bufio.go:526
		_go_fuzz_dep_.CoverTab[525767]++
//line /snap/go/10455/src/bufio/bufio.go:526
		_go_fuzz_dep_.CoverTab[2004]++
							m, err := r.WriteTo(w)
							n += m
							return n, err
//line /snap/go/10455/src/bufio/bufio.go:529
		// _ = "end of CoverTab[2004]"
	} else {
//line /snap/go/10455/src/bufio/bufio.go:530
		_go_fuzz_dep_.CoverTab[525768]++
//line /snap/go/10455/src/bufio/bufio.go:530
		_go_fuzz_dep_.CoverTab[2005]++
//line /snap/go/10455/src/bufio/bufio.go:530
		// _ = "end of CoverTab[2005]"
//line /snap/go/10455/src/bufio/bufio.go:530
	}
//line /snap/go/10455/src/bufio/bufio.go:530
	// _ = "end of CoverTab[1996]"
//line /snap/go/10455/src/bufio/bufio.go:530
	_go_fuzz_dep_.CoverTab[1997]++

						if w, ok := w.(io.ReaderFrom); ok {
//line /snap/go/10455/src/bufio/bufio.go:532
		_go_fuzz_dep_.CoverTab[525769]++
//line /snap/go/10455/src/bufio/bufio.go:532
		_go_fuzz_dep_.CoverTab[2006]++
							m, err := w.ReadFrom(b.rd)
							n += m
							return n, err
//line /snap/go/10455/src/bufio/bufio.go:535
		// _ = "end of CoverTab[2006]"
	} else {
//line /snap/go/10455/src/bufio/bufio.go:536
		_go_fuzz_dep_.CoverTab[525770]++
//line /snap/go/10455/src/bufio/bufio.go:536
		_go_fuzz_dep_.CoverTab[2007]++
//line /snap/go/10455/src/bufio/bufio.go:536
		// _ = "end of CoverTab[2007]"
//line /snap/go/10455/src/bufio/bufio.go:536
	}
//line /snap/go/10455/src/bufio/bufio.go:536
	// _ = "end of CoverTab[1997]"
//line /snap/go/10455/src/bufio/bufio.go:536
	_go_fuzz_dep_.CoverTab[1998]++

						if b.w-b.r < len(b.buf) {
//line /snap/go/10455/src/bufio/bufio.go:538
		_go_fuzz_dep_.CoverTab[525771]++
//line /snap/go/10455/src/bufio/bufio.go:538
		_go_fuzz_dep_.CoverTab[2008]++
							b.fill()
//line /snap/go/10455/src/bufio/bufio.go:539
		// _ = "end of CoverTab[2008]"
	} else {
//line /snap/go/10455/src/bufio/bufio.go:540
		_go_fuzz_dep_.CoverTab[525772]++
//line /snap/go/10455/src/bufio/bufio.go:540
		_go_fuzz_dep_.CoverTab[2009]++
//line /snap/go/10455/src/bufio/bufio.go:540
		// _ = "end of CoverTab[2009]"
//line /snap/go/10455/src/bufio/bufio.go:540
	}
//line /snap/go/10455/src/bufio/bufio.go:540
	// _ = "end of CoverTab[1998]"
//line /snap/go/10455/src/bufio/bufio.go:540
	_go_fuzz_dep_.CoverTab[1999]++
//line /snap/go/10455/src/bufio/bufio.go:540
	_go_fuzz_dep_.CoverTab[786563] = 0

						for b.r < b.w {
//line /snap/go/10455/src/bufio/bufio.go:542
		if _go_fuzz_dep_.CoverTab[786563] == 0 {
//line /snap/go/10455/src/bufio/bufio.go:542
			_go_fuzz_dep_.CoverTab[525877]++
//line /snap/go/10455/src/bufio/bufio.go:542
		} else {
//line /snap/go/10455/src/bufio/bufio.go:542
			_go_fuzz_dep_.CoverTab[525878]++
//line /snap/go/10455/src/bufio/bufio.go:542
		}
//line /snap/go/10455/src/bufio/bufio.go:542
		_go_fuzz_dep_.CoverTab[786563] = 1
//line /snap/go/10455/src/bufio/bufio.go:542
		_go_fuzz_dep_.CoverTab[2010]++

							m, err := b.writeBuf(w)
							n += m
							if err != nil {
//line /snap/go/10455/src/bufio/bufio.go:546
			_go_fuzz_dep_.CoverTab[525773]++
//line /snap/go/10455/src/bufio/bufio.go:546
			_go_fuzz_dep_.CoverTab[2012]++
								return n, err
//line /snap/go/10455/src/bufio/bufio.go:547
			// _ = "end of CoverTab[2012]"
		} else {
//line /snap/go/10455/src/bufio/bufio.go:548
			_go_fuzz_dep_.CoverTab[525774]++
//line /snap/go/10455/src/bufio/bufio.go:548
			_go_fuzz_dep_.CoverTab[2013]++
//line /snap/go/10455/src/bufio/bufio.go:548
			// _ = "end of CoverTab[2013]"
//line /snap/go/10455/src/bufio/bufio.go:548
		}
//line /snap/go/10455/src/bufio/bufio.go:548
		// _ = "end of CoverTab[2010]"
//line /snap/go/10455/src/bufio/bufio.go:548
		_go_fuzz_dep_.CoverTab[2011]++
							b.fill()
//line /snap/go/10455/src/bufio/bufio.go:549
		// _ = "end of CoverTab[2011]"
	}
//line /snap/go/10455/src/bufio/bufio.go:550
	if _go_fuzz_dep_.CoverTab[786563] == 0 {
//line /snap/go/10455/src/bufio/bufio.go:550
		_go_fuzz_dep_.CoverTab[525879]++
//line /snap/go/10455/src/bufio/bufio.go:550
	} else {
//line /snap/go/10455/src/bufio/bufio.go:550
		_go_fuzz_dep_.CoverTab[525880]++
//line /snap/go/10455/src/bufio/bufio.go:550
	}
//line /snap/go/10455/src/bufio/bufio.go:550
	// _ = "end of CoverTab[1999]"
//line /snap/go/10455/src/bufio/bufio.go:550
	_go_fuzz_dep_.CoverTab[2000]++

						if b.err == io.EOF {
//line /snap/go/10455/src/bufio/bufio.go:552
		_go_fuzz_dep_.CoverTab[525775]++
//line /snap/go/10455/src/bufio/bufio.go:552
		_go_fuzz_dep_.CoverTab[2014]++
							b.err = nil
//line /snap/go/10455/src/bufio/bufio.go:553
		// _ = "end of CoverTab[2014]"
	} else {
//line /snap/go/10455/src/bufio/bufio.go:554
		_go_fuzz_dep_.CoverTab[525776]++
//line /snap/go/10455/src/bufio/bufio.go:554
		_go_fuzz_dep_.CoverTab[2015]++
//line /snap/go/10455/src/bufio/bufio.go:554
		// _ = "end of CoverTab[2015]"
//line /snap/go/10455/src/bufio/bufio.go:554
	}
//line /snap/go/10455/src/bufio/bufio.go:554
	// _ = "end of CoverTab[2000]"
//line /snap/go/10455/src/bufio/bufio.go:554
	_go_fuzz_dep_.CoverTab[2001]++

						return n, b.readErr()
//line /snap/go/10455/src/bufio/bufio.go:556
	// _ = "end of CoverTab[2001]"
}

var errNegativeWrite = errors.New("bufio: writer returned negative count from Write")

// writeBuf writes the Reader's buffer to the writer.
func (b *Reader) writeBuf(w io.Writer) (int64, error) {
//line /snap/go/10455/src/bufio/bufio.go:562
	_go_fuzz_dep_.CoverTab[2016]++
						n, err := w.Write(b.buf[b.r:b.w])
						if n < 0 {
//line /snap/go/10455/src/bufio/bufio.go:564
		_go_fuzz_dep_.CoverTab[525777]++
//line /snap/go/10455/src/bufio/bufio.go:564
		_go_fuzz_dep_.CoverTab[2018]++
							panic(errNegativeWrite)
//line /snap/go/10455/src/bufio/bufio.go:565
		// _ = "end of CoverTab[2018]"
	} else {
//line /snap/go/10455/src/bufio/bufio.go:566
		_go_fuzz_dep_.CoverTab[525778]++
//line /snap/go/10455/src/bufio/bufio.go:566
		_go_fuzz_dep_.CoverTab[2019]++
//line /snap/go/10455/src/bufio/bufio.go:566
		// _ = "end of CoverTab[2019]"
//line /snap/go/10455/src/bufio/bufio.go:566
	}
//line /snap/go/10455/src/bufio/bufio.go:566
	// _ = "end of CoverTab[2016]"
//line /snap/go/10455/src/bufio/bufio.go:566
	_go_fuzz_dep_.CoverTab[2017]++
						b.r += n
						return int64(n), err
//line /snap/go/10455/src/bufio/bufio.go:568
	// _ = "end of CoverTab[2017]"
}

//line /snap/go/10455/src/bufio/bufio.go:573
// Writer implements buffering for an io.Writer object.
//line /snap/go/10455/src/bufio/bufio.go:573
// If an error occurs writing to a Writer, no more data will be
//line /snap/go/10455/src/bufio/bufio.go:573
// accepted and all subsequent writes, and Flush, will return the error.
//line /snap/go/10455/src/bufio/bufio.go:573
// After all data has been written, the client should call the
//line /snap/go/10455/src/bufio/bufio.go:573
// Flush method to guarantee all data has been forwarded to
//line /snap/go/10455/src/bufio/bufio.go:573
// the underlying io.Writer.
//line /snap/go/10455/src/bufio/bufio.go:579
type Writer struct {
	err	error
	buf	[]byte
	n	int
	wr	io.Writer
}

// NewWriterSize returns a new Writer whose buffer has at least the specified
//line /snap/go/10455/src/bufio/bufio.go:586
// size. If the argument io.Writer is already a Writer with large enough
//line /snap/go/10455/src/bufio/bufio.go:586
// size, it returns the underlying Writer.
//line /snap/go/10455/src/bufio/bufio.go:589
func NewWriterSize(w io.Writer, size int) *Writer {
//line /snap/go/10455/src/bufio/bufio.go:589
	_go_fuzz_dep_.CoverTab[2020]++

						b, ok := w.(*Writer)
						if ok && func() bool {
//line /snap/go/10455/src/bufio/bufio.go:592
		_go_fuzz_dep_.CoverTab[2023]++
//line /snap/go/10455/src/bufio/bufio.go:592
		return len(b.buf) >= size
//line /snap/go/10455/src/bufio/bufio.go:592
		// _ = "end of CoverTab[2023]"
//line /snap/go/10455/src/bufio/bufio.go:592
	}() {
//line /snap/go/10455/src/bufio/bufio.go:592
		_go_fuzz_dep_.CoverTab[525779]++
//line /snap/go/10455/src/bufio/bufio.go:592
		_go_fuzz_dep_.CoverTab[2024]++
							return b
//line /snap/go/10455/src/bufio/bufio.go:593
		// _ = "end of CoverTab[2024]"
	} else {
//line /snap/go/10455/src/bufio/bufio.go:594
		_go_fuzz_dep_.CoverTab[525780]++
//line /snap/go/10455/src/bufio/bufio.go:594
		_go_fuzz_dep_.CoverTab[2025]++
//line /snap/go/10455/src/bufio/bufio.go:594
		// _ = "end of CoverTab[2025]"
//line /snap/go/10455/src/bufio/bufio.go:594
	}
//line /snap/go/10455/src/bufio/bufio.go:594
	// _ = "end of CoverTab[2020]"
//line /snap/go/10455/src/bufio/bufio.go:594
	_go_fuzz_dep_.CoverTab[2021]++
						if size <= 0 {
//line /snap/go/10455/src/bufio/bufio.go:595
		_go_fuzz_dep_.CoverTab[525781]++
//line /snap/go/10455/src/bufio/bufio.go:595
		_go_fuzz_dep_.CoverTab[2026]++
							size = defaultBufSize
//line /snap/go/10455/src/bufio/bufio.go:596
		// _ = "end of CoverTab[2026]"
	} else {
//line /snap/go/10455/src/bufio/bufio.go:597
		_go_fuzz_dep_.CoverTab[525782]++
//line /snap/go/10455/src/bufio/bufio.go:597
		_go_fuzz_dep_.CoverTab[2027]++
//line /snap/go/10455/src/bufio/bufio.go:597
		// _ = "end of CoverTab[2027]"
//line /snap/go/10455/src/bufio/bufio.go:597
	}
//line /snap/go/10455/src/bufio/bufio.go:597
	// _ = "end of CoverTab[2021]"
//line /snap/go/10455/src/bufio/bufio.go:597
	_go_fuzz_dep_.CoverTab[2022]++
						return &Writer{
		buf:	make([]byte, size),
		wr:	w,
	}
//line /snap/go/10455/src/bufio/bufio.go:601
	// _ = "end of CoverTab[2022]"
}

// NewWriter returns a new Writer whose buffer has the default size.
//line /snap/go/10455/src/bufio/bufio.go:604
// If the argument io.Writer is already a Writer with large enough buffer size,
//line /snap/go/10455/src/bufio/bufio.go:604
// it returns the underlying Writer.
//line /snap/go/10455/src/bufio/bufio.go:607
func NewWriter(w io.Writer) *Writer {
//line /snap/go/10455/src/bufio/bufio.go:607
	_go_fuzz_dep_.CoverTab[2028]++
						return NewWriterSize(w, defaultBufSize)
//line /snap/go/10455/src/bufio/bufio.go:608
	// _ = "end of CoverTab[2028]"
}

// Size returns the size of the underlying buffer in bytes.
func (b *Writer) Size() int {
//line /snap/go/10455/src/bufio/bufio.go:612
	_go_fuzz_dep_.CoverTab[2029]++
//line /snap/go/10455/src/bufio/bufio.go:612
	return len(b.buf)
//line /snap/go/10455/src/bufio/bufio.go:612
	// _ = "end of CoverTab[2029]"
//line /snap/go/10455/src/bufio/bufio.go:612
}

// Reset discards any unflushed buffered data, clears any error, and
//line /snap/go/10455/src/bufio/bufio.go:614
// resets b to write its output to w.
//line /snap/go/10455/src/bufio/bufio.go:614
// Calling Reset on the zero value of Writer initializes the internal buffer
//line /snap/go/10455/src/bufio/bufio.go:614
// to the default size.
//line /snap/go/10455/src/bufio/bufio.go:614
// Calling w.Reset(w) (that is, resetting a Writer to itself) does nothing.
//line /snap/go/10455/src/bufio/bufio.go:619
func (b *Writer) Reset(w io.Writer) {
//line /snap/go/10455/src/bufio/bufio.go:619
	_go_fuzz_dep_.CoverTab[2030]++

//line /snap/go/10455/src/bufio/bufio.go:623
	if b == w {
//line /snap/go/10455/src/bufio/bufio.go:623
		_go_fuzz_dep_.CoverTab[525783]++
//line /snap/go/10455/src/bufio/bufio.go:623
		_go_fuzz_dep_.CoverTab[2033]++
							return
//line /snap/go/10455/src/bufio/bufio.go:624
		// _ = "end of CoverTab[2033]"
	} else {
//line /snap/go/10455/src/bufio/bufio.go:625
		_go_fuzz_dep_.CoverTab[525784]++
//line /snap/go/10455/src/bufio/bufio.go:625
		_go_fuzz_dep_.CoverTab[2034]++
//line /snap/go/10455/src/bufio/bufio.go:625
		// _ = "end of CoverTab[2034]"
//line /snap/go/10455/src/bufio/bufio.go:625
	}
//line /snap/go/10455/src/bufio/bufio.go:625
	// _ = "end of CoverTab[2030]"
//line /snap/go/10455/src/bufio/bufio.go:625
	_go_fuzz_dep_.CoverTab[2031]++
						if b.buf == nil {
//line /snap/go/10455/src/bufio/bufio.go:626
		_go_fuzz_dep_.CoverTab[525785]++
//line /snap/go/10455/src/bufio/bufio.go:626
		_go_fuzz_dep_.CoverTab[2035]++
							b.buf = make([]byte, defaultBufSize)
//line /snap/go/10455/src/bufio/bufio.go:627
		// _ = "end of CoverTab[2035]"
	} else {
//line /snap/go/10455/src/bufio/bufio.go:628
		_go_fuzz_dep_.CoverTab[525786]++
//line /snap/go/10455/src/bufio/bufio.go:628
		_go_fuzz_dep_.CoverTab[2036]++
//line /snap/go/10455/src/bufio/bufio.go:628
		// _ = "end of CoverTab[2036]"
//line /snap/go/10455/src/bufio/bufio.go:628
	}
//line /snap/go/10455/src/bufio/bufio.go:628
	// _ = "end of CoverTab[2031]"
//line /snap/go/10455/src/bufio/bufio.go:628
	_go_fuzz_dep_.CoverTab[2032]++
						b.err = nil
						b.n = 0
						b.wr = w
//line /snap/go/10455/src/bufio/bufio.go:631
	// _ = "end of CoverTab[2032]"
}

// Flush writes any buffered data to the underlying io.Writer.
func (b *Writer) Flush() error {
//line /snap/go/10455/src/bufio/bufio.go:635
	_go_fuzz_dep_.CoverTab[2037]++
						if b.err != nil {
//line /snap/go/10455/src/bufio/bufio.go:636
		_go_fuzz_dep_.CoverTab[525787]++
//line /snap/go/10455/src/bufio/bufio.go:636
		_go_fuzz_dep_.CoverTab[2042]++
							return b.err
//line /snap/go/10455/src/bufio/bufio.go:637
		// _ = "end of CoverTab[2042]"
	} else {
//line /snap/go/10455/src/bufio/bufio.go:638
		_go_fuzz_dep_.CoverTab[525788]++
//line /snap/go/10455/src/bufio/bufio.go:638
		_go_fuzz_dep_.CoverTab[2043]++
//line /snap/go/10455/src/bufio/bufio.go:638
		// _ = "end of CoverTab[2043]"
//line /snap/go/10455/src/bufio/bufio.go:638
	}
//line /snap/go/10455/src/bufio/bufio.go:638
	// _ = "end of CoverTab[2037]"
//line /snap/go/10455/src/bufio/bufio.go:638
	_go_fuzz_dep_.CoverTab[2038]++
						if b.n == 0 {
//line /snap/go/10455/src/bufio/bufio.go:639
		_go_fuzz_dep_.CoverTab[525789]++
//line /snap/go/10455/src/bufio/bufio.go:639
		_go_fuzz_dep_.CoverTab[2044]++
							return nil
//line /snap/go/10455/src/bufio/bufio.go:640
		// _ = "end of CoverTab[2044]"
	} else {
//line /snap/go/10455/src/bufio/bufio.go:641
		_go_fuzz_dep_.CoverTab[525790]++
//line /snap/go/10455/src/bufio/bufio.go:641
		_go_fuzz_dep_.CoverTab[2045]++
//line /snap/go/10455/src/bufio/bufio.go:641
		// _ = "end of CoverTab[2045]"
//line /snap/go/10455/src/bufio/bufio.go:641
	}
//line /snap/go/10455/src/bufio/bufio.go:641
	// _ = "end of CoverTab[2038]"
//line /snap/go/10455/src/bufio/bufio.go:641
	_go_fuzz_dep_.CoverTab[2039]++
						n, err := b.wr.Write(b.buf[0:b.n])
						if n < b.n && func() bool {
//line /snap/go/10455/src/bufio/bufio.go:643
		_go_fuzz_dep_.CoverTab[2046]++
//line /snap/go/10455/src/bufio/bufio.go:643
		return err == nil
//line /snap/go/10455/src/bufio/bufio.go:643
		// _ = "end of CoverTab[2046]"
//line /snap/go/10455/src/bufio/bufio.go:643
	}() {
//line /snap/go/10455/src/bufio/bufio.go:643
		_go_fuzz_dep_.CoverTab[525791]++
//line /snap/go/10455/src/bufio/bufio.go:643
		_go_fuzz_dep_.CoverTab[2047]++
							err = io.ErrShortWrite
//line /snap/go/10455/src/bufio/bufio.go:644
		// _ = "end of CoverTab[2047]"
	} else {
//line /snap/go/10455/src/bufio/bufio.go:645
		_go_fuzz_dep_.CoverTab[525792]++
//line /snap/go/10455/src/bufio/bufio.go:645
		_go_fuzz_dep_.CoverTab[2048]++
//line /snap/go/10455/src/bufio/bufio.go:645
		// _ = "end of CoverTab[2048]"
//line /snap/go/10455/src/bufio/bufio.go:645
	}
//line /snap/go/10455/src/bufio/bufio.go:645
	// _ = "end of CoverTab[2039]"
//line /snap/go/10455/src/bufio/bufio.go:645
	_go_fuzz_dep_.CoverTab[2040]++
						if err != nil {
//line /snap/go/10455/src/bufio/bufio.go:646
		_go_fuzz_dep_.CoverTab[525793]++
//line /snap/go/10455/src/bufio/bufio.go:646
		_go_fuzz_dep_.CoverTab[2049]++
							if n > 0 && func() bool {
//line /snap/go/10455/src/bufio/bufio.go:647
			_go_fuzz_dep_.CoverTab[2051]++
//line /snap/go/10455/src/bufio/bufio.go:647
			return n < b.n
//line /snap/go/10455/src/bufio/bufio.go:647
			// _ = "end of CoverTab[2051]"
//line /snap/go/10455/src/bufio/bufio.go:647
		}() {
//line /snap/go/10455/src/bufio/bufio.go:647
			_go_fuzz_dep_.CoverTab[525795]++
//line /snap/go/10455/src/bufio/bufio.go:647
			_go_fuzz_dep_.CoverTab[2052]++
								copy(b.buf[0:b.n-n], b.buf[n:b.n])
//line /snap/go/10455/src/bufio/bufio.go:648
			// _ = "end of CoverTab[2052]"
		} else {
//line /snap/go/10455/src/bufio/bufio.go:649
			_go_fuzz_dep_.CoverTab[525796]++
//line /snap/go/10455/src/bufio/bufio.go:649
			_go_fuzz_dep_.CoverTab[2053]++
//line /snap/go/10455/src/bufio/bufio.go:649
			// _ = "end of CoverTab[2053]"
//line /snap/go/10455/src/bufio/bufio.go:649
		}
//line /snap/go/10455/src/bufio/bufio.go:649
		// _ = "end of CoverTab[2049]"
//line /snap/go/10455/src/bufio/bufio.go:649
		_go_fuzz_dep_.CoverTab[2050]++
							b.n -= n
							b.err = err
							return err
//line /snap/go/10455/src/bufio/bufio.go:652
		// _ = "end of CoverTab[2050]"
	} else {
//line /snap/go/10455/src/bufio/bufio.go:653
		_go_fuzz_dep_.CoverTab[525794]++
//line /snap/go/10455/src/bufio/bufio.go:653
		_go_fuzz_dep_.CoverTab[2054]++
//line /snap/go/10455/src/bufio/bufio.go:653
		// _ = "end of CoverTab[2054]"
//line /snap/go/10455/src/bufio/bufio.go:653
	}
//line /snap/go/10455/src/bufio/bufio.go:653
	// _ = "end of CoverTab[2040]"
//line /snap/go/10455/src/bufio/bufio.go:653
	_go_fuzz_dep_.CoverTab[2041]++
						b.n = 0
						return nil
//line /snap/go/10455/src/bufio/bufio.go:655
	// _ = "end of CoverTab[2041]"
}

// Available returns how many bytes are unused in the buffer.
func (b *Writer) Available() int {
//line /snap/go/10455/src/bufio/bufio.go:659
	_go_fuzz_dep_.CoverTab[2055]++
//line /snap/go/10455/src/bufio/bufio.go:659
	return len(b.buf) - b.n
//line /snap/go/10455/src/bufio/bufio.go:659
	// _ = "end of CoverTab[2055]"
//line /snap/go/10455/src/bufio/bufio.go:659
}

// AvailableBuffer returns an empty buffer with b.Available() capacity.
//line /snap/go/10455/src/bufio/bufio.go:661
// This buffer is intended to be appended to and
//line /snap/go/10455/src/bufio/bufio.go:661
// passed to an immediately succeeding Write call.
//line /snap/go/10455/src/bufio/bufio.go:661
// The buffer is only valid until the next write operation on b.
//line /snap/go/10455/src/bufio/bufio.go:665
func (b *Writer) AvailableBuffer() []byte {
//line /snap/go/10455/src/bufio/bufio.go:665
	_go_fuzz_dep_.CoverTab[2056]++
						return b.buf[b.n:][:0]
//line /snap/go/10455/src/bufio/bufio.go:666
	// _ = "end of CoverTab[2056]"
}

// Buffered returns the number of bytes that have been written into the current buffer.
func (b *Writer) Buffered() int {
//line /snap/go/10455/src/bufio/bufio.go:670
	_go_fuzz_dep_.CoverTab[2057]++
//line /snap/go/10455/src/bufio/bufio.go:670
	return b.n
//line /snap/go/10455/src/bufio/bufio.go:670
	// _ = "end of CoverTab[2057]"
//line /snap/go/10455/src/bufio/bufio.go:670
}

// Write writes the contents of p into the buffer.
//line /snap/go/10455/src/bufio/bufio.go:672
// It returns the number of bytes written.
//line /snap/go/10455/src/bufio/bufio.go:672
// If nn < len(p), it also returns an error explaining
//line /snap/go/10455/src/bufio/bufio.go:672
// why the write is short.
//line /snap/go/10455/src/bufio/bufio.go:676
func (b *Writer) Write(p []byte) (nn int, err error) {
//line /snap/go/10455/src/bufio/bufio.go:676
	_go_fuzz_dep_.CoverTab[2058]++
//line /snap/go/10455/src/bufio/bufio.go:676
	_go_fuzz_dep_.CoverTab[786564] = 0
						for len(p) > b.Available() && func() bool {
//line /snap/go/10455/src/bufio/bufio.go:677
		_go_fuzz_dep_.CoverTab[2061]++
//line /snap/go/10455/src/bufio/bufio.go:677
		return b.err == nil
//line /snap/go/10455/src/bufio/bufio.go:677
		// _ = "end of CoverTab[2061]"
//line /snap/go/10455/src/bufio/bufio.go:677
	}() {
//line /snap/go/10455/src/bufio/bufio.go:677
		if _go_fuzz_dep_.CoverTab[786564] == 0 {
//line /snap/go/10455/src/bufio/bufio.go:677
			_go_fuzz_dep_.CoverTab[525881]++
//line /snap/go/10455/src/bufio/bufio.go:677
		} else {
//line /snap/go/10455/src/bufio/bufio.go:677
			_go_fuzz_dep_.CoverTab[525882]++
//line /snap/go/10455/src/bufio/bufio.go:677
		}
//line /snap/go/10455/src/bufio/bufio.go:677
		_go_fuzz_dep_.CoverTab[786564] = 1
//line /snap/go/10455/src/bufio/bufio.go:677
		_go_fuzz_dep_.CoverTab[2062]++
							var n int
							if b.Buffered() == 0 {
//line /snap/go/10455/src/bufio/bufio.go:679
			_go_fuzz_dep_.CoverTab[525797]++
//line /snap/go/10455/src/bufio/bufio.go:679
			_go_fuzz_dep_.CoverTab[2064]++

//line /snap/go/10455/src/bufio/bufio.go:682
			n, b.err = b.wr.Write(p)
//line /snap/go/10455/src/bufio/bufio.go:682
			// _ = "end of CoverTab[2064]"
		} else {
//line /snap/go/10455/src/bufio/bufio.go:683
			_go_fuzz_dep_.CoverTab[525798]++
//line /snap/go/10455/src/bufio/bufio.go:683
			_go_fuzz_dep_.CoverTab[2065]++
								n = copy(b.buf[b.n:], p)
								b.n += n
								b.Flush()
//line /snap/go/10455/src/bufio/bufio.go:686
			// _ = "end of CoverTab[2065]"
		}
//line /snap/go/10455/src/bufio/bufio.go:687
		// _ = "end of CoverTab[2062]"
//line /snap/go/10455/src/bufio/bufio.go:687
		_go_fuzz_dep_.CoverTab[2063]++
							nn += n
							p = p[n:]
//line /snap/go/10455/src/bufio/bufio.go:689
		// _ = "end of CoverTab[2063]"
	}
//line /snap/go/10455/src/bufio/bufio.go:690
	if _go_fuzz_dep_.CoverTab[786564] == 0 {
//line /snap/go/10455/src/bufio/bufio.go:690
		_go_fuzz_dep_.CoverTab[525883]++
//line /snap/go/10455/src/bufio/bufio.go:690
	} else {
//line /snap/go/10455/src/bufio/bufio.go:690
		_go_fuzz_dep_.CoverTab[525884]++
//line /snap/go/10455/src/bufio/bufio.go:690
	}
//line /snap/go/10455/src/bufio/bufio.go:690
	// _ = "end of CoverTab[2058]"
//line /snap/go/10455/src/bufio/bufio.go:690
	_go_fuzz_dep_.CoverTab[2059]++
						if b.err != nil {
//line /snap/go/10455/src/bufio/bufio.go:691
		_go_fuzz_dep_.CoverTab[525799]++
//line /snap/go/10455/src/bufio/bufio.go:691
		_go_fuzz_dep_.CoverTab[2066]++
							return nn, b.err
//line /snap/go/10455/src/bufio/bufio.go:692
		// _ = "end of CoverTab[2066]"
	} else {
//line /snap/go/10455/src/bufio/bufio.go:693
		_go_fuzz_dep_.CoverTab[525800]++
//line /snap/go/10455/src/bufio/bufio.go:693
		_go_fuzz_dep_.CoverTab[2067]++
//line /snap/go/10455/src/bufio/bufio.go:693
		// _ = "end of CoverTab[2067]"
//line /snap/go/10455/src/bufio/bufio.go:693
	}
//line /snap/go/10455/src/bufio/bufio.go:693
	// _ = "end of CoverTab[2059]"
//line /snap/go/10455/src/bufio/bufio.go:693
	_go_fuzz_dep_.CoverTab[2060]++
						n := copy(b.buf[b.n:], p)
						b.n += n
						nn += n
						return nn, nil
//line /snap/go/10455/src/bufio/bufio.go:697
	// _ = "end of CoverTab[2060]"
}

// WriteByte writes a single byte.
func (b *Writer) WriteByte(c byte) error {
//line /snap/go/10455/src/bufio/bufio.go:701
	_go_fuzz_dep_.CoverTab[2068]++
						if b.err != nil {
//line /snap/go/10455/src/bufio/bufio.go:702
		_go_fuzz_dep_.CoverTab[525801]++
//line /snap/go/10455/src/bufio/bufio.go:702
		_go_fuzz_dep_.CoverTab[2071]++
							return b.err
//line /snap/go/10455/src/bufio/bufio.go:703
		// _ = "end of CoverTab[2071]"
	} else {
//line /snap/go/10455/src/bufio/bufio.go:704
		_go_fuzz_dep_.CoverTab[525802]++
//line /snap/go/10455/src/bufio/bufio.go:704
		_go_fuzz_dep_.CoverTab[2072]++
//line /snap/go/10455/src/bufio/bufio.go:704
		// _ = "end of CoverTab[2072]"
//line /snap/go/10455/src/bufio/bufio.go:704
	}
//line /snap/go/10455/src/bufio/bufio.go:704
	// _ = "end of CoverTab[2068]"
//line /snap/go/10455/src/bufio/bufio.go:704
	_go_fuzz_dep_.CoverTab[2069]++
						if b.Available() <= 0 && func() bool {
//line /snap/go/10455/src/bufio/bufio.go:705
		_go_fuzz_dep_.CoverTab[2073]++
//line /snap/go/10455/src/bufio/bufio.go:705
		return b.Flush() != nil
//line /snap/go/10455/src/bufio/bufio.go:705
		// _ = "end of CoverTab[2073]"
//line /snap/go/10455/src/bufio/bufio.go:705
	}() {
//line /snap/go/10455/src/bufio/bufio.go:705
		_go_fuzz_dep_.CoverTab[525803]++
//line /snap/go/10455/src/bufio/bufio.go:705
		_go_fuzz_dep_.CoverTab[2074]++
							return b.err
//line /snap/go/10455/src/bufio/bufio.go:706
		// _ = "end of CoverTab[2074]"
	} else {
//line /snap/go/10455/src/bufio/bufio.go:707
		_go_fuzz_dep_.CoverTab[525804]++
//line /snap/go/10455/src/bufio/bufio.go:707
		_go_fuzz_dep_.CoverTab[2075]++
//line /snap/go/10455/src/bufio/bufio.go:707
		// _ = "end of CoverTab[2075]"
//line /snap/go/10455/src/bufio/bufio.go:707
	}
//line /snap/go/10455/src/bufio/bufio.go:707
	// _ = "end of CoverTab[2069]"
//line /snap/go/10455/src/bufio/bufio.go:707
	_go_fuzz_dep_.CoverTab[2070]++
						b.buf[b.n] = c
						b.n++
						return nil
//line /snap/go/10455/src/bufio/bufio.go:710
	// _ = "end of CoverTab[2070]"
}

// WriteRune writes a single Unicode code point, returning
//line /snap/go/10455/src/bufio/bufio.go:713
// the number of bytes written and any error.
//line /snap/go/10455/src/bufio/bufio.go:715
func (b *Writer) WriteRune(r rune) (size int, err error) {
//line /snap/go/10455/src/bufio/bufio.go:715
	_go_fuzz_dep_.CoverTab[2076]++

						if uint32(r) < utf8.RuneSelf {
//line /snap/go/10455/src/bufio/bufio.go:717
		_go_fuzz_dep_.CoverTab[525805]++
//line /snap/go/10455/src/bufio/bufio.go:717
		_go_fuzz_dep_.CoverTab[2080]++
							err = b.WriteByte(byte(r))
							if err != nil {
//line /snap/go/10455/src/bufio/bufio.go:719
			_go_fuzz_dep_.CoverTab[525807]++
//line /snap/go/10455/src/bufio/bufio.go:719
			_go_fuzz_dep_.CoverTab[2082]++
								return 0, err
//line /snap/go/10455/src/bufio/bufio.go:720
			// _ = "end of CoverTab[2082]"
		} else {
//line /snap/go/10455/src/bufio/bufio.go:721
			_go_fuzz_dep_.CoverTab[525808]++
//line /snap/go/10455/src/bufio/bufio.go:721
			_go_fuzz_dep_.CoverTab[2083]++
//line /snap/go/10455/src/bufio/bufio.go:721
			// _ = "end of CoverTab[2083]"
//line /snap/go/10455/src/bufio/bufio.go:721
		}
//line /snap/go/10455/src/bufio/bufio.go:721
		// _ = "end of CoverTab[2080]"
//line /snap/go/10455/src/bufio/bufio.go:721
		_go_fuzz_dep_.CoverTab[2081]++
							return 1, nil
//line /snap/go/10455/src/bufio/bufio.go:722
		// _ = "end of CoverTab[2081]"
	} else {
//line /snap/go/10455/src/bufio/bufio.go:723
		_go_fuzz_dep_.CoverTab[525806]++
//line /snap/go/10455/src/bufio/bufio.go:723
		_go_fuzz_dep_.CoverTab[2084]++
//line /snap/go/10455/src/bufio/bufio.go:723
		// _ = "end of CoverTab[2084]"
//line /snap/go/10455/src/bufio/bufio.go:723
	}
//line /snap/go/10455/src/bufio/bufio.go:723
	// _ = "end of CoverTab[2076]"
//line /snap/go/10455/src/bufio/bufio.go:723
	_go_fuzz_dep_.CoverTab[2077]++
						if b.err != nil {
//line /snap/go/10455/src/bufio/bufio.go:724
		_go_fuzz_dep_.CoverTab[525809]++
//line /snap/go/10455/src/bufio/bufio.go:724
		_go_fuzz_dep_.CoverTab[2085]++
							return 0, b.err
//line /snap/go/10455/src/bufio/bufio.go:725
		// _ = "end of CoverTab[2085]"
	} else {
//line /snap/go/10455/src/bufio/bufio.go:726
		_go_fuzz_dep_.CoverTab[525810]++
//line /snap/go/10455/src/bufio/bufio.go:726
		_go_fuzz_dep_.CoverTab[2086]++
//line /snap/go/10455/src/bufio/bufio.go:726
		// _ = "end of CoverTab[2086]"
//line /snap/go/10455/src/bufio/bufio.go:726
	}
//line /snap/go/10455/src/bufio/bufio.go:726
	// _ = "end of CoverTab[2077]"
//line /snap/go/10455/src/bufio/bufio.go:726
	_go_fuzz_dep_.CoverTab[2078]++
						n := b.Available()
						if n < utf8.UTFMax {
//line /snap/go/10455/src/bufio/bufio.go:728
		_go_fuzz_dep_.CoverTab[525811]++
//line /snap/go/10455/src/bufio/bufio.go:728
		_go_fuzz_dep_.CoverTab[2087]++
							if b.Flush(); b.err != nil {
//line /snap/go/10455/src/bufio/bufio.go:729
			_go_fuzz_dep_.CoverTab[525813]++
//line /snap/go/10455/src/bufio/bufio.go:729
			_go_fuzz_dep_.CoverTab[2089]++
								return 0, b.err
//line /snap/go/10455/src/bufio/bufio.go:730
			// _ = "end of CoverTab[2089]"
		} else {
//line /snap/go/10455/src/bufio/bufio.go:731
			_go_fuzz_dep_.CoverTab[525814]++
//line /snap/go/10455/src/bufio/bufio.go:731
			_go_fuzz_dep_.CoverTab[2090]++
//line /snap/go/10455/src/bufio/bufio.go:731
			// _ = "end of CoverTab[2090]"
//line /snap/go/10455/src/bufio/bufio.go:731
		}
//line /snap/go/10455/src/bufio/bufio.go:731
		// _ = "end of CoverTab[2087]"
//line /snap/go/10455/src/bufio/bufio.go:731
		_go_fuzz_dep_.CoverTab[2088]++
							n = b.Available()
							if n < utf8.UTFMax {
//line /snap/go/10455/src/bufio/bufio.go:733
			_go_fuzz_dep_.CoverTab[525815]++
//line /snap/go/10455/src/bufio/bufio.go:733
			_go_fuzz_dep_.CoverTab[2091]++

								return b.WriteString(string(r))
//line /snap/go/10455/src/bufio/bufio.go:735
			// _ = "end of CoverTab[2091]"
		} else {
//line /snap/go/10455/src/bufio/bufio.go:736
			_go_fuzz_dep_.CoverTab[525816]++
//line /snap/go/10455/src/bufio/bufio.go:736
			_go_fuzz_dep_.CoverTab[2092]++
//line /snap/go/10455/src/bufio/bufio.go:736
			// _ = "end of CoverTab[2092]"
//line /snap/go/10455/src/bufio/bufio.go:736
		}
//line /snap/go/10455/src/bufio/bufio.go:736
		// _ = "end of CoverTab[2088]"
	} else {
//line /snap/go/10455/src/bufio/bufio.go:737
		_go_fuzz_dep_.CoverTab[525812]++
//line /snap/go/10455/src/bufio/bufio.go:737
		_go_fuzz_dep_.CoverTab[2093]++
//line /snap/go/10455/src/bufio/bufio.go:737
		// _ = "end of CoverTab[2093]"
//line /snap/go/10455/src/bufio/bufio.go:737
	}
//line /snap/go/10455/src/bufio/bufio.go:737
	// _ = "end of CoverTab[2078]"
//line /snap/go/10455/src/bufio/bufio.go:737
	_go_fuzz_dep_.CoverTab[2079]++
						size = utf8.EncodeRune(b.buf[b.n:], r)
						b.n += size
						return size, nil
//line /snap/go/10455/src/bufio/bufio.go:740
	// _ = "end of CoverTab[2079]"
}

// WriteString writes a string.
//line /snap/go/10455/src/bufio/bufio.go:743
// It returns the number of bytes written.
//line /snap/go/10455/src/bufio/bufio.go:743
// If the count is less than len(s), it also returns an error explaining
//line /snap/go/10455/src/bufio/bufio.go:743
// why the write is short.
//line /snap/go/10455/src/bufio/bufio.go:747
func (b *Writer) WriteString(s string) (int, error) {
//line /snap/go/10455/src/bufio/bufio.go:747
	_go_fuzz_dep_.CoverTab[2094]++
						var sw io.StringWriter
						tryStringWriter := true

						nn := 0
//line /snap/go/10455/src/bufio/bufio.go:751
	_go_fuzz_dep_.CoverTab[786565] = 0
						for len(s) > b.Available() && func() bool {
//line /snap/go/10455/src/bufio/bufio.go:752
		_go_fuzz_dep_.CoverTab[2097]++
//line /snap/go/10455/src/bufio/bufio.go:752
		return b.err == nil
//line /snap/go/10455/src/bufio/bufio.go:752
		// _ = "end of CoverTab[2097]"
//line /snap/go/10455/src/bufio/bufio.go:752
	}() {
//line /snap/go/10455/src/bufio/bufio.go:752
		if _go_fuzz_dep_.CoverTab[786565] == 0 {
//line /snap/go/10455/src/bufio/bufio.go:752
			_go_fuzz_dep_.CoverTab[525885]++
//line /snap/go/10455/src/bufio/bufio.go:752
		} else {
//line /snap/go/10455/src/bufio/bufio.go:752
			_go_fuzz_dep_.CoverTab[525886]++
//line /snap/go/10455/src/bufio/bufio.go:752
		}
//line /snap/go/10455/src/bufio/bufio.go:752
		_go_fuzz_dep_.CoverTab[786565] = 1
//line /snap/go/10455/src/bufio/bufio.go:752
		_go_fuzz_dep_.CoverTab[2098]++
							var n int
							if b.Buffered() == 0 && func() bool {
//line /snap/go/10455/src/bufio/bufio.go:754
			_go_fuzz_dep_.CoverTab[2101]++
//line /snap/go/10455/src/bufio/bufio.go:754
			return sw == nil
//line /snap/go/10455/src/bufio/bufio.go:754
			// _ = "end of CoverTab[2101]"
//line /snap/go/10455/src/bufio/bufio.go:754
		}() && func() bool {
//line /snap/go/10455/src/bufio/bufio.go:754
			_go_fuzz_dep_.CoverTab[2102]++
//line /snap/go/10455/src/bufio/bufio.go:754
			return tryStringWriter
//line /snap/go/10455/src/bufio/bufio.go:754
			// _ = "end of CoverTab[2102]"
//line /snap/go/10455/src/bufio/bufio.go:754
		}() {
//line /snap/go/10455/src/bufio/bufio.go:754
			_go_fuzz_dep_.CoverTab[525817]++
//line /snap/go/10455/src/bufio/bufio.go:754
			_go_fuzz_dep_.CoverTab[2103]++

								sw, tryStringWriter = b.wr.(io.StringWriter)
//line /snap/go/10455/src/bufio/bufio.go:756
			// _ = "end of CoverTab[2103]"
		} else {
//line /snap/go/10455/src/bufio/bufio.go:757
			_go_fuzz_dep_.CoverTab[525818]++
//line /snap/go/10455/src/bufio/bufio.go:757
			_go_fuzz_dep_.CoverTab[2104]++
//line /snap/go/10455/src/bufio/bufio.go:757
			// _ = "end of CoverTab[2104]"
//line /snap/go/10455/src/bufio/bufio.go:757
		}
//line /snap/go/10455/src/bufio/bufio.go:757
		// _ = "end of CoverTab[2098]"
//line /snap/go/10455/src/bufio/bufio.go:757
		_go_fuzz_dep_.CoverTab[2099]++
							if b.Buffered() == 0 && func() bool {
//line /snap/go/10455/src/bufio/bufio.go:758
			_go_fuzz_dep_.CoverTab[2105]++
//line /snap/go/10455/src/bufio/bufio.go:758
			return tryStringWriter
//line /snap/go/10455/src/bufio/bufio.go:758
			// _ = "end of CoverTab[2105]"
//line /snap/go/10455/src/bufio/bufio.go:758
		}() {
//line /snap/go/10455/src/bufio/bufio.go:758
			_go_fuzz_dep_.CoverTab[525819]++
//line /snap/go/10455/src/bufio/bufio.go:758
			_go_fuzz_dep_.CoverTab[2106]++

//line /snap/go/10455/src/bufio/bufio.go:762
			n, b.err = sw.WriteString(s)
//line /snap/go/10455/src/bufio/bufio.go:762
			// _ = "end of CoverTab[2106]"
		} else {
//line /snap/go/10455/src/bufio/bufio.go:763
			_go_fuzz_dep_.CoverTab[525820]++
//line /snap/go/10455/src/bufio/bufio.go:763
			_go_fuzz_dep_.CoverTab[2107]++
								n = copy(b.buf[b.n:], s)
								b.n += n
								b.Flush()
//line /snap/go/10455/src/bufio/bufio.go:766
			// _ = "end of CoverTab[2107]"
		}
//line /snap/go/10455/src/bufio/bufio.go:767
		// _ = "end of CoverTab[2099]"
//line /snap/go/10455/src/bufio/bufio.go:767
		_go_fuzz_dep_.CoverTab[2100]++
							nn += n
							s = s[n:]
//line /snap/go/10455/src/bufio/bufio.go:769
		// _ = "end of CoverTab[2100]"
	}
//line /snap/go/10455/src/bufio/bufio.go:770
	if _go_fuzz_dep_.CoverTab[786565] == 0 {
//line /snap/go/10455/src/bufio/bufio.go:770
		_go_fuzz_dep_.CoverTab[525887]++
//line /snap/go/10455/src/bufio/bufio.go:770
	} else {
//line /snap/go/10455/src/bufio/bufio.go:770
		_go_fuzz_dep_.CoverTab[525888]++
//line /snap/go/10455/src/bufio/bufio.go:770
	}
//line /snap/go/10455/src/bufio/bufio.go:770
	// _ = "end of CoverTab[2094]"
//line /snap/go/10455/src/bufio/bufio.go:770
	_go_fuzz_dep_.CoverTab[2095]++
						if b.err != nil {
//line /snap/go/10455/src/bufio/bufio.go:771
		_go_fuzz_dep_.CoverTab[525821]++
//line /snap/go/10455/src/bufio/bufio.go:771
		_go_fuzz_dep_.CoverTab[2108]++
							return nn, b.err
//line /snap/go/10455/src/bufio/bufio.go:772
		// _ = "end of CoverTab[2108]"
	} else {
//line /snap/go/10455/src/bufio/bufio.go:773
		_go_fuzz_dep_.CoverTab[525822]++
//line /snap/go/10455/src/bufio/bufio.go:773
		_go_fuzz_dep_.CoverTab[2109]++
//line /snap/go/10455/src/bufio/bufio.go:773
		// _ = "end of CoverTab[2109]"
//line /snap/go/10455/src/bufio/bufio.go:773
	}
//line /snap/go/10455/src/bufio/bufio.go:773
	// _ = "end of CoverTab[2095]"
//line /snap/go/10455/src/bufio/bufio.go:773
	_go_fuzz_dep_.CoverTab[2096]++
						n := copy(b.buf[b.n:], s)
						b.n += n
						nn += n
						return nn, nil
//line /snap/go/10455/src/bufio/bufio.go:777
	// _ = "end of CoverTab[2096]"
}

// ReadFrom implements io.ReaderFrom. If the underlying writer
//line /snap/go/10455/src/bufio/bufio.go:780
// supports the ReadFrom method, this calls the underlying ReadFrom.
//line /snap/go/10455/src/bufio/bufio.go:780
// If there is buffered data and an underlying ReadFrom, this fills
//line /snap/go/10455/src/bufio/bufio.go:780
// the buffer and writes it before calling ReadFrom.
//line /snap/go/10455/src/bufio/bufio.go:784
func (b *Writer) ReadFrom(r io.Reader) (n int64, err error) {
//line /snap/go/10455/src/bufio/bufio.go:784
	_go_fuzz_dep_.CoverTab[2110]++
						if b.err != nil {
//line /snap/go/10455/src/bufio/bufio.go:785
		_go_fuzz_dep_.CoverTab[525823]++
//line /snap/go/10455/src/bufio/bufio.go:785
		_go_fuzz_dep_.CoverTab[2114]++
							return 0, b.err
//line /snap/go/10455/src/bufio/bufio.go:786
		// _ = "end of CoverTab[2114]"
	} else {
//line /snap/go/10455/src/bufio/bufio.go:787
		_go_fuzz_dep_.CoverTab[525824]++
//line /snap/go/10455/src/bufio/bufio.go:787
		_go_fuzz_dep_.CoverTab[2115]++
//line /snap/go/10455/src/bufio/bufio.go:787
		// _ = "end of CoverTab[2115]"
//line /snap/go/10455/src/bufio/bufio.go:787
	}
//line /snap/go/10455/src/bufio/bufio.go:787
	// _ = "end of CoverTab[2110]"
//line /snap/go/10455/src/bufio/bufio.go:787
	_go_fuzz_dep_.CoverTab[2111]++
						readerFrom, readerFromOK := b.wr.(io.ReaderFrom)
						var m int
//line /snap/go/10455/src/bufio/bufio.go:789
	_go_fuzz_dep_.CoverTab[786566] = 0
						for {
//line /snap/go/10455/src/bufio/bufio.go:790
		if _go_fuzz_dep_.CoverTab[786566] == 0 {
//line /snap/go/10455/src/bufio/bufio.go:790
			_go_fuzz_dep_.CoverTab[525889]++
//line /snap/go/10455/src/bufio/bufio.go:790
		} else {
//line /snap/go/10455/src/bufio/bufio.go:790
			_go_fuzz_dep_.CoverTab[525890]++
//line /snap/go/10455/src/bufio/bufio.go:790
		}
//line /snap/go/10455/src/bufio/bufio.go:790
		_go_fuzz_dep_.CoverTab[786566] = 1
//line /snap/go/10455/src/bufio/bufio.go:790
		_go_fuzz_dep_.CoverTab[2116]++
							if b.Available() == 0 {
//line /snap/go/10455/src/bufio/bufio.go:791
			_go_fuzz_dep_.CoverTab[525825]++
//line /snap/go/10455/src/bufio/bufio.go:791
			_go_fuzz_dep_.CoverTab[2121]++
								if err1 := b.Flush(); err1 != nil {
//line /snap/go/10455/src/bufio/bufio.go:792
				_go_fuzz_dep_.CoverTab[525827]++
//line /snap/go/10455/src/bufio/bufio.go:792
				_go_fuzz_dep_.CoverTab[2122]++
									return n, err1
//line /snap/go/10455/src/bufio/bufio.go:793
				// _ = "end of CoverTab[2122]"
			} else {
//line /snap/go/10455/src/bufio/bufio.go:794
				_go_fuzz_dep_.CoverTab[525828]++
//line /snap/go/10455/src/bufio/bufio.go:794
				_go_fuzz_dep_.CoverTab[2123]++
//line /snap/go/10455/src/bufio/bufio.go:794
				// _ = "end of CoverTab[2123]"
//line /snap/go/10455/src/bufio/bufio.go:794
			}
//line /snap/go/10455/src/bufio/bufio.go:794
			// _ = "end of CoverTab[2121]"
		} else {
//line /snap/go/10455/src/bufio/bufio.go:795
			_go_fuzz_dep_.CoverTab[525826]++
//line /snap/go/10455/src/bufio/bufio.go:795
			_go_fuzz_dep_.CoverTab[2124]++
//line /snap/go/10455/src/bufio/bufio.go:795
			// _ = "end of CoverTab[2124]"
//line /snap/go/10455/src/bufio/bufio.go:795
		}
//line /snap/go/10455/src/bufio/bufio.go:795
		// _ = "end of CoverTab[2116]"
//line /snap/go/10455/src/bufio/bufio.go:795
		_go_fuzz_dep_.CoverTab[2117]++
							if readerFromOK && func() bool {
//line /snap/go/10455/src/bufio/bufio.go:796
			_go_fuzz_dep_.CoverTab[2125]++
//line /snap/go/10455/src/bufio/bufio.go:796
			return b.Buffered() == 0
//line /snap/go/10455/src/bufio/bufio.go:796
			// _ = "end of CoverTab[2125]"
//line /snap/go/10455/src/bufio/bufio.go:796
		}() {
//line /snap/go/10455/src/bufio/bufio.go:796
			_go_fuzz_dep_.CoverTab[525829]++
//line /snap/go/10455/src/bufio/bufio.go:796
			_go_fuzz_dep_.CoverTab[2126]++
								nn, err := readerFrom.ReadFrom(r)
								b.err = err
								n += nn
								return n, err
//line /snap/go/10455/src/bufio/bufio.go:800
			// _ = "end of CoverTab[2126]"
		} else {
//line /snap/go/10455/src/bufio/bufio.go:801
			_go_fuzz_dep_.CoverTab[525830]++
//line /snap/go/10455/src/bufio/bufio.go:801
			_go_fuzz_dep_.CoverTab[2127]++
//line /snap/go/10455/src/bufio/bufio.go:801
			// _ = "end of CoverTab[2127]"
//line /snap/go/10455/src/bufio/bufio.go:801
		}
//line /snap/go/10455/src/bufio/bufio.go:801
		// _ = "end of CoverTab[2117]"
//line /snap/go/10455/src/bufio/bufio.go:801
		_go_fuzz_dep_.CoverTab[2118]++
							nr := 0
//line /snap/go/10455/src/bufio/bufio.go:802
		_go_fuzz_dep_.CoverTab[786567] = 0
							for nr < maxConsecutiveEmptyReads {
//line /snap/go/10455/src/bufio/bufio.go:803
			if _go_fuzz_dep_.CoverTab[786567] == 0 {
//line /snap/go/10455/src/bufio/bufio.go:803
				_go_fuzz_dep_.CoverTab[525893]++
//line /snap/go/10455/src/bufio/bufio.go:803
			} else {
//line /snap/go/10455/src/bufio/bufio.go:803
				_go_fuzz_dep_.CoverTab[525894]++
//line /snap/go/10455/src/bufio/bufio.go:803
			}
//line /snap/go/10455/src/bufio/bufio.go:803
			_go_fuzz_dep_.CoverTab[786567] = 1
//line /snap/go/10455/src/bufio/bufio.go:803
			_go_fuzz_dep_.CoverTab[2128]++
								m, err = r.Read(b.buf[b.n:])
								if m != 0 || func() bool {
//line /snap/go/10455/src/bufio/bufio.go:805
				_go_fuzz_dep_.CoverTab[2130]++
//line /snap/go/10455/src/bufio/bufio.go:805
				return err != nil
//line /snap/go/10455/src/bufio/bufio.go:805
				// _ = "end of CoverTab[2130]"
//line /snap/go/10455/src/bufio/bufio.go:805
			}() {
//line /snap/go/10455/src/bufio/bufio.go:805
				_go_fuzz_dep_.CoverTab[525831]++
//line /snap/go/10455/src/bufio/bufio.go:805
				_go_fuzz_dep_.CoverTab[2131]++
									break
//line /snap/go/10455/src/bufio/bufio.go:806
				// _ = "end of CoverTab[2131]"
			} else {
//line /snap/go/10455/src/bufio/bufio.go:807
				_go_fuzz_dep_.CoverTab[525832]++
//line /snap/go/10455/src/bufio/bufio.go:807
				_go_fuzz_dep_.CoverTab[2132]++
//line /snap/go/10455/src/bufio/bufio.go:807
				// _ = "end of CoverTab[2132]"
//line /snap/go/10455/src/bufio/bufio.go:807
			}
//line /snap/go/10455/src/bufio/bufio.go:807
			// _ = "end of CoverTab[2128]"
//line /snap/go/10455/src/bufio/bufio.go:807
			_go_fuzz_dep_.CoverTab[2129]++
								nr++
//line /snap/go/10455/src/bufio/bufio.go:808
			// _ = "end of CoverTab[2129]"
		}
//line /snap/go/10455/src/bufio/bufio.go:809
		if _go_fuzz_dep_.CoverTab[786567] == 0 {
//line /snap/go/10455/src/bufio/bufio.go:809
			_go_fuzz_dep_.CoverTab[525895]++
//line /snap/go/10455/src/bufio/bufio.go:809
		} else {
//line /snap/go/10455/src/bufio/bufio.go:809
			_go_fuzz_dep_.CoverTab[525896]++
//line /snap/go/10455/src/bufio/bufio.go:809
		}
//line /snap/go/10455/src/bufio/bufio.go:809
		// _ = "end of CoverTab[2118]"
//line /snap/go/10455/src/bufio/bufio.go:809
		_go_fuzz_dep_.CoverTab[2119]++
							if nr == maxConsecutiveEmptyReads {
//line /snap/go/10455/src/bufio/bufio.go:810
			_go_fuzz_dep_.CoverTab[525833]++
//line /snap/go/10455/src/bufio/bufio.go:810
			_go_fuzz_dep_.CoverTab[2133]++
								return n, io.ErrNoProgress
//line /snap/go/10455/src/bufio/bufio.go:811
			// _ = "end of CoverTab[2133]"
		} else {
//line /snap/go/10455/src/bufio/bufio.go:812
			_go_fuzz_dep_.CoverTab[525834]++
//line /snap/go/10455/src/bufio/bufio.go:812
			_go_fuzz_dep_.CoverTab[2134]++
//line /snap/go/10455/src/bufio/bufio.go:812
			// _ = "end of CoverTab[2134]"
//line /snap/go/10455/src/bufio/bufio.go:812
		}
//line /snap/go/10455/src/bufio/bufio.go:812
		// _ = "end of CoverTab[2119]"
//line /snap/go/10455/src/bufio/bufio.go:812
		_go_fuzz_dep_.CoverTab[2120]++
							b.n += m
							n += int64(m)
							if err != nil {
//line /snap/go/10455/src/bufio/bufio.go:815
			_go_fuzz_dep_.CoverTab[525835]++
//line /snap/go/10455/src/bufio/bufio.go:815
			_go_fuzz_dep_.CoverTab[2135]++
								break
//line /snap/go/10455/src/bufio/bufio.go:816
			// _ = "end of CoverTab[2135]"
		} else {
//line /snap/go/10455/src/bufio/bufio.go:817
			_go_fuzz_dep_.CoverTab[525836]++
//line /snap/go/10455/src/bufio/bufio.go:817
			_go_fuzz_dep_.CoverTab[2136]++
//line /snap/go/10455/src/bufio/bufio.go:817
			// _ = "end of CoverTab[2136]"
//line /snap/go/10455/src/bufio/bufio.go:817
		}
//line /snap/go/10455/src/bufio/bufio.go:817
		// _ = "end of CoverTab[2120]"
	}
//line /snap/go/10455/src/bufio/bufio.go:818
	// _ = "end of CoverTab[2111]"
//line /snap/go/10455/src/bufio/bufio.go:818
	_go_fuzz_dep_.CoverTab[2112]++
						if err == io.EOF {
//line /snap/go/10455/src/bufio/bufio.go:819
		_go_fuzz_dep_.CoverTab[525837]++
//line /snap/go/10455/src/bufio/bufio.go:819
		_go_fuzz_dep_.CoverTab[2137]++

							if b.Available() == 0 {
//line /snap/go/10455/src/bufio/bufio.go:821
			_go_fuzz_dep_.CoverTab[525839]++
//line /snap/go/10455/src/bufio/bufio.go:821
			_go_fuzz_dep_.CoverTab[2138]++
								err = b.Flush()
//line /snap/go/10455/src/bufio/bufio.go:822
			// _ = "end of CoverTab[2138]"
		} else {
//line /snap/go/10455/src/bufio/bufio.go:823
			_go_fuzz_dep_.CoverTab[525840]++
//line /snap/go/10455/src/bufio/bufio.go:823
			_go_fuzz_dep_.CoverTab[2139]++
								err = nil
//line /snap/go/10455/src/bufio/bufio.go:824
			// _ = "end of CoverTab[2139]"
		}
//line /snap/go/10455/src/bufio/bufio.go:825
		// _ = "end of CoverTab[2137]"
	} else {
//line /snap/go/10455/src/bufio/bufio.go:826
		_go_fuzz_dep_.CoverTab[525838]++
//line /snap/go/10455/src/bufio/bufio.go:826
		_go_fuzz_dep_.CoverTab[2140]++
//line /snap/go/10455/src/bufio/bufio.go:826
		// _ = "end of CoverTab[2140]"
//line /snap/go/10455/src/bufio/bufio.go:826
	}
//line /snap/go/10455/src/bufio/bufio.go:826
	// _ = "end of CoverTab[2112]"
//line /snap/go/10455/src/bufio/bufio.go:826
	_go_fuzz_dep_.CoverTab[2113]++
						return n, err
//line /snap/go/10455/src/bufio/bufio.go:827
	// _ = "end of CoverTab[2113]"
}

//line /snap/go/10455/src/bufio/bufio.go:832
// ReadWriter stores pointers to a Reader and a Writer.
//line /snap/go/10455/src/bufio/bufio.go:832
// It implements io.ReadWriter.
//line /snap/go/10455/src/bufio/bufio.go:834
type ReadWriter struct {
	*Reader
	*Writer
}

// NewReadWriter allocates a new ReadWriter that dispatches to r and w.
func NewReadWriter(r *Reader, w *Writer) *ReadWriter {
//line /snap/go/10455/src/bufio/bufio.go:840
	_go_fuzz_dep_.CoverTab[2141]++
						return &ReadWriter{r, w}
//line /snap/go/10455/src/bufio/bufio.go:841
	// _ = "end of CoverTab[2141]"
}

//line /snap/go/10455/src/bufio/bufio.go:842
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /snap/go/10455/src/bufio/bufio.go:842
var _ = _go_fuzz_dep_.CoverTab
