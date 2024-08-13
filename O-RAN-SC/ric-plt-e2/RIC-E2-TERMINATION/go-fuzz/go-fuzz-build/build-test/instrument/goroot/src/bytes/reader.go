// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/bytes/reader.go:5
package bytes

//line /usr/local/go/src/bytes/reader.go:5
import (
//line /usr/local/go/src/bytes/reader.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/bytes/reader.go:5
)
//line /usr/local/go/src/bytes/reader.go:5
import (
//line /usr/local/go/src/bytes/reader.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/bytes/reader.go:5
)

import (
	"errors"
	"io"
	"unicode/utf8"
)

// A Reader implements the io.Reader, io.ReaderAt, io.WriterTo, io.Seeker,
//line /usr/local/go/src/bytes/reader.go:13
// io.ByteScanner, and io.RuneScanner interfaces by reading from
//line /usr/local/go/src/bytes/reader.go:13
// a byte slice.
//line /usr/local/go/src/bytes/reader.go:13
// Unlike a Buffer, a Reader is read-only and supports seeking.
//line /usr/local/go/src/bytes/reader.go:13
// The zero value for Reader operates like a Reader of an empty slice.
//line /usr/local/go/src/bytes/reader.go:18
type Reader struct {
	s		[]byte
	i		int64	// current reading index
	prevRune	int	// index of previous rune; or < 0
}

// Len returns the number of bytes of the unread portion of the
//line /usr/local/go/src/bytes/reader.go:24
// slice.
//line /usr/local/go/src/bytes/reader.go:26
func (r *Reader) Len() int {
//line /usr/local/go/src/bytes/reader.go:26
	_go_fuzz_dep_.CoverTab[801]++
						if r.i >= int64(len(r.s)) {
//line /usr/local/go/src/bytes/reader.go:27
		_go_fuzz_dep_.CoverTab[803]++
							return 0
//line /usr/local/go/src/bytes/reader.go:28
		// _ = "end of CoverTab[803]"
	} else {
//line /usr/local/go/src/bytes/reader.go:29
		_go_fuzz_dep_.CoverTab[804]++
//line /usr/local/go/src/bytes/reader.go:29
		// _ = "end of CoverTab[804]"
//line /usr/local/go/src/bytes/reader.go:29
	}
//line /usr/local/go/src/bytes/reader.go:29
	// _ = "end of CoverTab[801]"
//line /usr/local/go/src/bytes/reader.go:29
	_go_fuzz_dep_.CoverTab[802]++
						return int(int64(len(r.s)) - r.i)
//line /usr/local/go/src/bytes/reader.go:30
	// _ = "end of CoverTab[802]"
}

// Size returns the original length of the underlying byte slice.
//line /usr/local/go/src/bytes/reader.go:33
// Size is the number of bytes available for reading via ReadAt.
//line /usr/local/go/src/bytes/reader.go:33
// The result is unaffected by any method calls except Reset.
//line /usr/local/go/src/bytes/reader.go:36
func (r *Reader) Size() int64 {
//line /usr/local/go/src/bytes/reader.go:36
	_go_fuzz_dep_.CoverTab[805]++
//line /usr/local/go/src/bytes/reader.go:36
	return int64(len(r.s))
//line /usr/local/go/src/bytes/reader.go:36
	// _ = "end of CoverTab[805]"
//line /usr/local/go/src/bytes/reader.go:36
}

// Read implements the io.Reader interface.
func (r *Reader) Read(b []byte) (n int, err error) {
//line /usr/local/go/src/bytes/reader.go:39
	_go_fuzz_dep_.CoverTab[806]++
						if r.i >= int64(len(r.s)) {
//line /usr/local/go/src/bytes/reader.go:40
		_go_fuzz_dep_.CoverTab[808]++
							return 0, io.EOF
//line /usr/local/go/src/bytes/reader.go:41
		// _ = "end of CoverTab[808]"
	} else {
//line /usr/local/go/src/bytes/reader.go:42
		_go_fuzz_dep_.CoverTab[809]++
//line /usr/local/go/src/bytes/reader.go:42
		// _ = "end of CoverTab[809]"
//line /usr/local/go/src/bytes/reader.go:42
	}
//line /usr/local/go/src/bytes/reader.go:42
	// _ = "end of CoverTab[806]"
//line /usr/local/go/src/bytes/reader.go:42
	_go_fuzz_dep_.CoverTab[807]++
						r.prevRune = -1
						n = copy(b, r.s[r.i:])
						r.i += int64(n)
						return
//line /usr/local/go/src/bytes/reader.go:46
	// _ = "end of CoverTab[807]"
}

// ReadAt implements the io.ReaderAt interface.
func (r *Reader) ReadAt(b []byte, off int64) (n int, err error) {
//line /usr/local/go/src/bytes/reader.go:50
	_go_fuzz_dep_.CoverTab[810]++

						if off < 0 {
//line /usr/local/go/src/bytes/reader.go:52
		_go_fuzz_dep_.CoverTab[814]++
							return 0, errors.New("bytes.Reader.ReadAt: negative offset")
//line /usr/local/go/src/bytes/reader.go:53
		// _ = "end of CoverTab[814]"
	} else {
//line /usr/local/go/src/bytes/reader.go:54
		_go_fuzz_dep_.CoverTab[815]++
//line /usr/local/go/src/bytes/reader.go:54
		// _ = "end of CoverTab[815]"
//line /usr/local/go/src/bytes/reader.go:54
	}
//line /usr/local/go/src/bytes/reader.go:54
	// _ = "end of CoverTab[810]"
//line /usr/local/go/src/bytes/reader.go:54
	_go_fuzz_dep_.CoverTab[811]++
						if off >= int64(len(r.s)) {
//line /usr/local/go/src/bytes/reader.go:55
		_go_fuzz_dep_.CoverTab[816]++
							return 0, io.EOF
//line /usr/local/go/src/bytes/reader.go:56
		// _ = "end of CoverTab[816]"
	} else {
//line /usr/local/go/src/bytes/reader.go:57
		_go_fuzz_dep_.CoverTab[817]++
//line /usr/local/go/src/bytes/reader.go:57
		// _ = "end of CoverTab[817]"
//line /usr/local/go/src/bytes/reader.go:57
	}
//line /usr/local/go/src/bytes/reader.go:57
	// _ = "end of CoverTab[811]"
//line /usr/local/go/src/bytes/reader.go:57
	_go_fuzz_dep_.CoverTab[812]++
						n = copy(b, r.s[off:])
						if n < len(b) {
//line /usr/local/go/src/bytes/reader.go:59
		_go_fuzz_dep_.CoverTab[818]++
							err = io.EOF
//line /usr/local/go/src/bytes/reader.go:60
		// _ = "end of CoverTab[818]"
	} else {
//line /usr/local/go/src/bytes/reader.go:61
		_go_fuzz_dep_.CoverTab[819]++
//line /usr/local/go/src/bytes/reader.go:61
		// _ = "end of CoverTab[819]"
//line /usr/local/go/src/bytes/reader.go:61
	}
//line /usr/local/go/src/bytes/reader.go:61
	// _ = "end of CoverTab[812]"
//line /usr/local/go/src/bytes/reader.go:61
	_go_fuzz_dep_.CoverTab[813]++
						return
//line /usr/local/go/src/bytes/reader.go:62
	// _ = "end of CoverTab[813]"
}

// ReadByte implements the io.ByteReader interface.
func (r *Reader) ReadByte() (byte, error) {
//line /usr/local/go/src/bytes/reader.go:66
	_go_fuzz_dep_.CoverTab[820]++
						r.prevRune = -1
						if r.i >= int64(len(r.s)) {
//line /usr/local/go/src/bytes/reader.go:68
		_go_fuzz_dep_.CoverTab[822]++
							return 0, io.EOF
//line /usr/local/go/src/bytes/reader.go:69
		// _ = "end of CoverTab[822]"
	} else {
//line /usr/local/go/src/bytes/reader.go:70
		_go_fuzz_dep_.CoverTab[823]++
//line /usr/local/go/src/bytes/reader.go:70
		// _ = "end of CoverTab[823]"
//line /usr/local/go/src/bytes/reader.go:70
	}
//line /usr/local/go/src/bytes/reader.go:70
	// _ = "end of CoverTab[820]"
//line /usr/local/go/src/bytes/reader.go:70
	_go_fuzz_dep_.CoverTab[821]++
						b := r.s[r.i]
						r.i++
						return b, nil
//line /usr/local/go/src/bytes/reader.go:73
	// _ = "end of CoverTab[821]"
}

// UnreadByte complements ReadByte in implementing the io.ByteScanner interface.
func (r *Reader) UnreadByte() error {
//line /usr/local/go/src/bytes/reader.go:77
	_go_fuzz_dep_.CoverTab[824]++
						if r.i <= 0 {
//line /usr/local/go/src/bytes/reader.go:78
		_go_fuzz_dep_.CoverTab[826]++
							return errors.New("bytes.Reader.UnreadByte: at beginning of slice")
//line /usr/local/go/src/bytes/reader.go:79
		// _ = "end of CoverTab[826]"
	} else {
//line /usr/local/go/src/bytes/reader.go:80
		_go_fuzz_dep_.CoverTab[827]++
//line /usr/local/go/src/bytes/reader.go:80
		// _ = "end of CoverTab[827]"
//line /usr/local/go/src/bytes/reader.go:80
	}
//line /usr/local/go/src/bytes/reader.go:80
	// _ = "end of CoverTab[824]"
//line /usr/local/go/src/bytes/reader.go:80
	_go_fuzz_dep_.CoverTab[825]++
						r.prevRune = -1
						r.i--
						return nil
//line /usr/local/go/src/bytes/reader.go:83
	// _ = "end of CoverTab[825]"
}

// ReadRune implements the io.RuneReader interface.
func (r *Reader) ReadRune() (ch rune, size int, err error) {
//line /usr/local/go/src/bytes/reader.go:87
	_go_fuzz_dep_.CoverTab[828]++
						if r.i >= int64(len(r.s)) {
//line /usr/local/go/src/bytes/reader.go:88
		_go_fuzz_dep_.CoverTab[831]++
							r.prevRune = -1
							return 0, 0, io.EOF
//line /usr/local/go/src/bytes/reader.go:90
		// _ = "end of CoverTab[831]"
	} else {
//line /usr/local/go/src/bytes/reader.go:91
		_go_fuzz_dep_.CoverTab[832]++
//line /usr/local/go/src/bytes/reader.go:91
		// _ = "end of CoverTab[832]"
//line /usr/local/go/src/bytes/reader.go:91
	}
//line /usr/local/go/src/bytes/reader.go:91
	// _ = "end of CoverTab[828]"
//line /usr/local/go/src/bytes/reader.go:91
	_go_fuzz_dep_.CoverTab[829]++
						r.prevRune = int(r.i)
						if c := r.s[r.i]; c < utf8.RuneSelf {
//line /usr/local/go/src/bytes/reader.go:93
		_go_fuzz_dep_.CoverTab[833]++
							r.i++
							return rune(c), 1, nil
//line /usr/local/go/src/bytes/reader.go:95
		// _ = "end of CoverTab[833]"
	} else {
//line /usr/local/go/src/bytes/reader.go:96
		_go_fuzz_dep_.CoverTab[834]++
//line /usr/local/go/src/bytes/reader.go:96
		// _ = "end of CoverTab[834]"
//line /usr/local/go/src/bytes/reader.go:96
	}
//line /usr/local/go/src/bytes/reader.go:96
	// _ = "end of CoverTab[829]"
//line /usr/local/go/src/bytes/reader.go:96
	_go_fuzz_dep_.CoverTab[830]++
						ch, size = utf8.DecodeRune(r.s[r.i:])
						r.i += int64(size)
						return
//line /usr/local/go/src/bytes/reader.go:99
	// _ = "end of CoverTab[830]"
}

// UnreadRune complements ReadRune in implementing the io.RuneScanner interface.
func (r *Reader) UnreadRune() error {
//line /usr/local/go/src/bytes/reader.go:103
	_go_fuzz_dep_.CoverTab[835]++
						if r.i <= 0 {
//line /usr/local/go/src/bytes/reader.go:104
		_go_fuzz_dep_.CoverTab[838]++
							return errors.New("bytes.Reader.UnreadRune: at beginning of slice")
//line /usr/local/go/src/bytes/reader.go:105
		// _ = "end of CoverTab[838]"
	} else {
//line /usr/local/go/src/bytes/reader.go:106
		_go_fuzz_dep_.CoverTab[839]++
//line /usr/local/go/src/bytes/reader.go:106
		// _ = "end of CoverTab[839]"
//line /usr/local/go/src/bytes/reader.go:106
	}
//line /usr/local/go/src/bytes/reader.go:106
	// _ = "end of CoverTab[835]"
//line /usr/local/go/src/bytes/reader.go:106
	_go_fuzz_dep_.CoverTab[836]++
						if r.prevRune < 0 {
//line /usr/local/go/src/bytes/reader.go:107
		_go_fuzz_dep_.CoverTab[840]++
							return errors.New("bytes.Reader.UnreadRune: previous operation was not ReadRune")
//line /usr/local/go/src/bytes/reader.go:108
		// _ = "end of CoverTab[840]"
	} else {
//line /usr/local/go/src/bytes/reader.go:109
		_go_fuzz_dep_.CoverTab[841]++
//line /usr/local/go/src/bytes/reader.go:109
		// _ = "end of CoverTab[841]"
//line /usr/local/go/src/bytes/reader.go:109
	}
//line /usr/local/go/src/bytes/reader.go:109
	// _ = "end of CoverTab[836]"
//line /usr/local/go/src/bytes/reader.go:109
	_go_fuzz_dep_.CoverTab[837]++
						r.i = int64(r.prevRune)
						r.prevRune = -1
						return nil
//line /usr/local/go/src/bytes/reader.go:112
	// _ = "end of CoverTab[837]"
}

// Seek implements the io.Seeker interface.
func (r *Reader) Seek(offset int64, whence int) (int64, error) {
//line /usr/local/go/src/bytes/reader.go:116
	_go_fuzz_dep_.CoverTab[842]++
						r.prevRune = -1
						var abs int64
						switch whence {
	case io.SeekStart:
//line /usr/local/go/src/bytes/reader.go:120
		_go_fuzz_dep_.CoverTab[845]++
							abs = offset
//line /usr/local/go/src/bytes/reader.go:121
		// _ = "end of CoverTab[845]"
	case io.SeekCurrent:
//line /usr/local/go/src/bytes/reader.go:122
		_go_fuzz_dep_.CoverTab[846]++
							abs = r.i + offset
//line /usr/local/go/src/bytes/reader.go:123
		// _ = "end of CoverTab[846]"
	case io.SeekEnd:
//line /usr/local/go/src/bytes/reader.go:124
		_go_fuzz_dep_.CoverTab[847]++
							abs = int64(len(r.s)) + offset
//line /usr/local/go/src/bytes/reader.go:125
		// _ = "end of CoverTab[847]"
	default:
//line /usr/local/go/src/bytes/reader.go:126
		_go_fuzz_dep_.CoverTab[848]++
							return 0, errors.New("bytes.Reader.Seek: invalid whence")
//line /usr/local/go/src/bytes/reader.go:127
		// _ = "end of CoverTab[848]"
	}
//line /usr/local/go/src/bytes/reader.go:128
	// _ = "end of CoverTab[842]"
//line /usr/local/go/src/bytes/reader.go:128
	_go_fuzz_dep_.CoverTab[843]++
						if abs < 0 {
//line /usr/local/go/src/bytes/reader.go:129
		_go_fuzz_dep_.CoverTab[849]++
							return 0, errors.New("bytes.Reader.Seek: negative position")
//line /usr/local/go/src/bytes/reader.go:130
		// _ = "end of CoverTab[849]"
	} else {
//line /usr/local/go/src/bytes/reader.go:131
		_go_fuzz_dep_.CoverTab[850]++
//line /usr/local/go/src/bytes/reader.go:131
		// _ = "end of CoverTab[850]"
//line /usr/local/go/src/bytes/reader.go:131
	}
//line /usr/local/go/src/bytes/reader.go:131
	// _ = "end of CoverTab[843]"
//line /usr/local/go/src/bytes/reader.go:131
	_go_fuzz_dep_.CoverTab[844]++
						r.i = abs
						return abs, nil
//line /usr/local/go/src/bytes/reader.go:133
	// _ = "end of CoverTab[844]"
}

// WriteTo implements the io.WriterTo interface.
func (r *Reader) WriteTo(w io.Writer) (n int64, err error) {
//line /usr/local/go/src/bytes/reader.go:137
	_go_fuzz_dep_.CoverTab[851]++
						r.prevRune = -1
						if r.i >= int64(len(r.s)) {
//line /usr/local/go/src/bytes/reader.go:139
		_go_fuzz_dep_.CoverTab[855]++
							return 0, nil
//line /usr/local/go/src/bytes/reader.go:140
		// _ = "end of CoverTab[855]"
	} else {
//line /usr/local/go/src/bytes/reader.go:141
		_go_fuzz_dep_.CoverTab[856]++
//line /usr/local/go/src/bytes/reader.go:141
		// _ = "end of CoverTab[856]"
//line /usr/local/go/src/bytes/reader.go:141
	}
//line /usr/local/go/src/bytes/reader.go:141
	// _ = "end of CoverTab[851]"
//line /usr/local/go/src/bytes/reader.go:141
	_go_fuzz_dep_.CoverTab[852]++
						b := r.s[r.i:]
						m, err := w.Write(b)
						if m > len(b) {
//line /usr/local/go/src/bytes/reader.go:144
		_go_fuzz_dep_.CoverTab[857]++
							panic("bytes.Reader.WriteTo: invalid Write count")
//line /usr/local/go/src/bytes/reader.go:145
		// _ = "end of CoverTab[857]"
	} else {
//line /usr/local/go/src/bytes/reader.go:146
		_go_fuzz_dep_.CoverTab[858]++
//line /usr/local/go/src/bytes/reader.go:146
		// _ = "end of CoverTab[858]"
//line /usr/local/go/src/bytes/reader.go:146
	}
//line /usr/local/go/src/bytes/reader.go:146
	// _ = "end of CoverTab[852]"
//line /usr/local/go/src/bytes/reader.go:146
	_go_fuzz_dep_.CoverTab[853]++
						r.i += int64(m)
						n = int64(m)
						if m != len(b) && func() bool {
//line /usr/local/go/src/bytes/reader.go:149
		_go_fuzz_dep_.CoverTab[859]++
//line /usr/local/go/src/bytes/reader.go:149
		return err == nil
//line /usr/local/go/src/bytes/reader.go:149
		// _ = "end of CoverTab[859]"
//line /usr/local/go/src/bytes/reader.go:149
	}() {
//line /usr/local/go/src/bytes/reader.go:149
		_go_fuzz_dep_.CoverTab[860]++
							err = io.ErrShortWrite
//line /usr/local/go/src/bytes/reader.go:150
		// _ = "end of CoverTab[860]"
	} else {
//line /usr/local/go/src/bytes/reader.go:151
		_go_fuzz_dep_.CoverTab[861]++
//line /usr/local/go/src/bytes/reader.go:151
		// _ = "end of CoverTab[861]"
//line /usr/local/go/src/bytes/reader.go:151
	}
//line /usr/local/go/src/bytes/reader.go:151
	// _ = "end of CoverTab[853]"
//line /usr/local/go/src/bytes/reader.go:151
	_go_fuzz_dep_.CoverTab[854]++
						return
//line /usr/local/go/src/bytes/reader.go:152
	// _ = "end of CoverTab[854]"
}

// Reset resets the Reader to be reading from b.
func (r *Reader) Reset(b []byte) {
//line /usr/local/go/src/bytes/reader.go:156
	_go_fuzz_dep_.CoverTab[862]++
//line /usr/local/go/src/bytes/reader.go:156
	*r = Reader{b, 0, -1}
//line /usr/local/go/src/bytes/reader.go:156
	// _ = "end of CoverTab[862]"
//line /usr/local/go/src/bytes/reader.go:156
}

// NewReader returns a new Reader reading from b.
func NewReader(b []byte) *Reader {
//line /usr/local/go/src/bytes/reader.go:159
	_go_fuzz_dep_.CoverTab[863]++
//line /usr/local/go/src/bytes/reader.go:159
	return &Reader{b, 0, -1}
//line /usr/local/go/src/bytes/reader.go:159
	// _ = "end of CoverTab[863]"
//line /usr/local/go/src/bytes/reader.go:159
}

//line /usr/local/go/src/bytes/reader.go:159
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/bytes/reader.go:159
var _ = _go_fuzz_dep_.CoverTab
