// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /snap/go/10455/src/bytes/reader.go:5
package bytes

//line /snap/go/10455/src/bytes/reader.go:5
import (
//line /snap/go/10455/src/bytes/reader.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /snap/go/10455/src/bytes/reader.go:5
)
//line /snap/go/10455/src/bytes/reader.go:5
import (
//line /snap/go/10455/src/bytes/reader.go:5
	_atomic_ "sync/atomic"
//line /snap/go/10455/src/bytes/reader.go:5
)

import (
	"errors"
	"io"
	"unicode/utf8"
)

// A Reader implements the io.Reader, io.ReaderAt, io.WriterTo, io.Seeker,
//line /snap/go/10455/src/bytes/reader.go:13
// io.ByteScanner, and io.RuneScanner interfaces by reading from
//line /snap/go/10455/src/bytes/reader.go:13
// a byte slice.
//line /snap/go/10455/src/bytes/reader.go:13
// Unlike a Buffer, a Reader is read-only and supports seeking.
//line /snap/go/10455/src/bytes/reader.go:13
// The zero value for Reader operates like a Reader of an empty slice.
//line /snap/go/10455/src/bytes/reader.go:18
type Reader struct {
	s		[]byte
	i		int64	// current reading index
	prevRune	int	// index of previous rune; or < 0
}

// Len returns the number of bytes of the unread portion of the
//line /snap/go/10455/src/bytes/reader.go:24
// slice.
//line /snap/go/10455/src/bytes/reader.go:26
func (r *Reader) Len() int {
//line /snap/go/10455/src/bytes/reader.go:26
	_go_fuzz_dep_.CoverTab[814]++
						if r.i >= int64(len(r.s)) {
//line /snap/go/10455/src/bytes/reader.go:27
		_go_fuzz_dep_.CoverTab[524895]++
//line /snap/go/10455/src/bytes/reader.go:27
		_go_fuzz_dep_.CoverTab[816]++
							return 0
//line /snap/go/10455/src/bytes/reader.go:28
		// _ = "end of CoverTab[816]"
	} else {
//line /snap/go/10455/src/bytes/reader.go:29
		_go_fuzz_dep_.CoverTab[524896]++
//line /snap/go/10455/src/bytes/reader.go:29
		_go_fuzz_dep_.CoverTab[817]++
//line /snap/go/10455/src/bytes/reader.go:29
		// _ = "end of CoverTab[817]"
//line /snap/go/10455/src/bytes/reader.go:29
	}
//line /snap/go/10455/src/bytes/reader.go:29
	// _ = "end of CoverTab[814]"
//line /snap/go/10455/src/bytes/reader.go:29
	_go_fuzz_dep_.CoverTab[815]++
						return int(int64(len(r.s)) - r.i)
//line /snap/go/10455/src/bytes/reader.go:30
	// _ = "end of CoverTab[815]"
}

// Size returns the original length of the underlying byte slice.
//line /snap/go/10455/src/bytes/reader.go:33
// Size is the number of bytes available for reading via ReadAt.
//line /snap/go/10455/src/bytes/reader.go:33
// The result is unaffected by any method calls except Reset.
//line /snap/go/10455/src/bytes/reader.go:36
func (r *Reader) Size() int64 {
//line /snap/go/10455/src/bytes/reader.go:36
	_go_fuzz_dep_.CoverTab[818]++
//line /snap/go/10455/src/bytes/reader.go:36
	return int64(len(r.s))
//line /snap/go/10455/src/bytes/reader.go:36
	// _ = "end of CoverTab[818]"
//line /snap/go/10455/src/bytes/reader.go:36
}

// Read implements the io.Reader interface.
func (r *Reader) Read(b []byte) (n int, err error) {
//line /snap/go/10455/src/bytes/reader.go:39
	_go_fuzz_dep_.CoverTab[819]++
						if r.i >= int64(len(r.s)) {
//line /snap/go/10455/src/bytes/reader.go:40
		_go_fuzz_dep_.CoverTab[524897]++
//line /snap/go/10455/src/bytes/reader.go:40
		_go_fuzz_dep_.CoverTab[821]++
							return 0, io.EOF
//line /snap/go/10455/src/bytes/reader.go:41
		// _ = "end of CoverTab[821]"
	} else {
//line /snap/go/10455/src/bytes/reader.go:42
		_go_fuzz_dep_.CoverTab[524898]++
//line /snap/go/10455/src/bytes/reader.go:42
		_go_fuzz_dep_.CoverTab[822]++
//line /snap/go/10455/src/bytes/reader.go:42
		// _ = "end of CoverTab[822]"
//line /snap/go/10455/src/bytes/reader.go:42
	}
//line /snap/go/10455/src/bytes/reader.go:42
	// _ = "end of CoverTab[819]"
//line /snap/go/10455/src/bytes/reader.go:42
	_go_fuzz_dep_.CoverTab[820]++
						r.prevRune = -1
						n = copy(b, r.s[r.i:])
						r.i += int64(n)
						return
//line /snap/go/10455/src/bytes/reader.go:46
	// _ = "end of CoverTab[820]"
}

// ReadAt implements the io.ReaderAt interface.
func (r *Reader) ReadAt(b []byte, off int64) (n int, err error) {
//line /snap/go/10455/src/bytes/reader.go:50
	_go_fuzz_dep_.CoverTab[823]++

						if off < 0 {
//line /snap/go/10455/src/bytes/reader.go:52
		_go_fuzz_dep_.CoverTab[524899]++
//line /snap/go/10455/src/bytes/reader.go:52
		_go_fuzz_dep_.CoverTab[827]++
							return 0, errors.New("bytes.Reader.ReadAt: negative offset")
//line /snap/go/10455/src/bytes/reader.go:53
		// _ = "end of CoverTab[827]"
	} else {
//line /snap/go/10455/src/bytes/reader.go:54
		_go_fuzz_dep_.CoverTab[524900]++
//line /snap/go/10455/src/bytes/reader.go:54
		_go_fuzz_dep_.CoverTab[828]++
//line /snap/go/10455/src/bytes/reader.go:54
		// _ = "end of CoverTab[828]"
//line /snap/go/10455/src/bytes/reader.go:54
	}
//line /snap/go/10455/src/bytes/reader.go:54
	// _ = "end of CoverTab[823]"
//line /snap/go/10455/src/bytes/reader.go:54
	_go_fuzz_dep_.CoverTab[824]++
						if off >= int64(len(r.s)) {
//line /snap/go/10455/src/bytes/reader.go:55
		_go_fuzz_dep_.CoverTab[524901]++
//line /snap/go/10455/src/bytes/reader.go:55
		_go_fuzz_dep_.CoverTab[829]++
							return 0, io.EOF
//line /snap/go/10455/src/bytes/reader.go:56
		// _ = "end of CoverTab[829]"
	} else {
//line /snap/go/10455/src/bytes/reader.go:57
		_go_fuzz_dep_.CoverTab[524902]++
//line /snap/go/10455/src/bytes/reader.go:57
		_go_fuzz_dep_.CoverTab[830]++
//line /snap/go/10455/src/bytes/reader.go:57
		// _ = "end of CoverTab[830]"
//line /snap/go/10455/src/bytes/reader.go:57
	}
//line /snap/go/10455/src/bytes/reader.go:57
	// _ = "end of CoverTab[824]"
//line /snap/go/10455/src/bytes/reader.go:57
	_go_fuzz_dep_.CoverTab[825]++
						n = copy(b, r.s[off:])
						if n < len(b) {
//line /snap/go/10455/src/bytes/reader.go:59
		_go_fuzz_dep_.CoverTab[524903]++
//line /snap/go/10455/src/bytes/reader.go:59
		_go_fuzz_dep_.CoverTab[831]++
							err = io.EOF
//line /snap/go/10455/src/bytes/reader.go:60
		// _ = "end of CoverTab[831]"
	} else {
//line /snap/go/10455/src/bytes/reader.go:61
		_go_fuzz_dep_.CoverTab[524904]++
//line /snap/go/10455/src/bytes/reader.go:61
		_go_fuzz_dep_.CoverTab[832]++
//line /snap/go/10455/src/bytes/reader.go:61
		// _ = "end of CoverTab[832]"
//line /snap/go/10455/src/bytes/reader.go:61
	}
//line /snap/go/10455/src/bytes/reader.go:61
	// _ = "end of CoverTab[825]"
//line /snap/go/10455/src/bytes/reader.go:61
	_go_fuzz_dep_.CoverTab[826]++
						return
//line /snap/go/10455/src/bytes/reader.go:62
	// _ = "end of CoverTab[826]"
}

// ReadByte implements the io.ByteReader interface.
func (r *Reader) ReadByte() (byte, error) {
//line /snap/go/10455/src/bytes/reader.go:66
	_go_fuzz_dep_.CoverTab[833]++
						r.prevRune = -1
						if r.i >= int64(len(r.s)) {
//line /snap/go/10455/src/bytes/reader.go:68
		_go_fuzz_dep_.CoverTab[524905]++
//line /snap/go/10455/src/bytes/reader.go:68
		_go_fuzz_dep_.CoverTab[835]++
							return 0, io.EOF
//line /snap/go/10455/src/bytes/reader.go:69
		// _ = "end of CoverTab[835]"
	} else {
//line /snap/go/10455/src/bytes/reader.go:70
		_go_fuzz_dep_.CoverTab[524906]++
//line /snap/go/10455/src/bytes/reader.go:70
		_go_fuzz_dep_.CoverTab[836]++
//line /snap/go/10455/src/bytes/reader.go:70
		// _ = "end of CoverTab[836]"
//line /snap/go/10455/src/bytes/reader.go:70
	}
//line /snap/go/10455/src/bytes/reader.go:70
	// _ = "end of CoverTab[833]"
//line /snap/go/10455/src/bytes/reader.go:70
	_go_fuzz_dep_.CoverTab[834]++
						b := r.s[r.i]
						r.i++
						return b, nil
//line /snap/go/10455/src/bytes/reader.go:73
	// _ = "end of CoverTab[834]"
}

// UnreadByte complements ReadByte in implementing the io.ByteScanner interface.
func (r *Reader) UnreadByte() error {
//line /snap/go/10455/src/bytes/reader.go:77
	_go_fuzz_dep_.CoverTab[837]++
						if r.i <= 0 {
//line /snap/go/10455/src/bytes/reader.go:78
		_go_fuzz_dep_.CoverTab[524907]++
//line /snap/go/10455/src/bytes/reader.go:78
		_go_fuzz_dep_.CoverTab[839]++
							return errors.New("bytes.Reader.UnreadByte: at beginning of slice")
//line /snap/go/10455/src/bytes/reader.go:79
		// _ = "end of CoverTab[839]"
	} else {
//line /snap/go/10455/src/bytes/reader.go:80
		_go_fuzz_dep_.CoverTab[524908]++
//line /snap/go/10455/src/bytes/reader.go:80
		_go_fuzz_dep_.CoverTab[840]++
//line /snap/go/10455/src/bytes/reader.go:80
		// _ = "end of CoverTab[840]"
//line /snap/go/10455/src/bytes/reader.go:80
	}
//line /snap/go/10455/src/bytes/reader.go:80
	// _ = "end of CoverTab[837]"
//line /snap/go/10455/src/bytes/reader.go:80
	_go_fuzz_dep_.CoverTab[838]++
						r.prevRune = -1
						r.i--
						return nil
//line /snap/go/10455/src/bytes/reader.go:83
	// _ = "end of CoverTab[838]"
}

// ReadRune implements the io.RuneReader interface.
func (r *Reader) ReadRune() (ch rune, size int, err error) {
//line /snap/go/10455/src/bytes/reader.go:87
	_go_fuzz_dep_.CoverTab[841]++
						if r.i >= int64(len(r.s)) {
//line /snap/go/10455/src/bytes/reader.go:88
		_go_fuzz_dep_.CoverTab[524909]++
//line /snap/go/10455/src/bytes/reader.go:88
		_go_fuzz_dep_.CoverTab[844]++
							r.prevRune = -1
							return 0, 0, io.EOF
//line /snap/go/10455/src/bytes/reader.go:90
		// _ = "end of CoverTab[844]"
	} else {
//line /snap/go/10455/src/bytes/reader.go:91
		_go_fuzz_dep_.CoverTab[524910]++
//line /snap/go/10455/src/bytes/reader.go:91
		_go_fuzz_dep_.CoverTab[845]++
//line /snap/go/10455/src/bytes/reader.go:91
		// _ = "end of CoverTab[845]"
//line /snap/go/10455/src/bytes/reader.go:91
	}
//line /snap/go/10455/src/bytes/reader.go:91
	// _ = "end of CoverTab[841]"
//line /snap/go/10455/src/bytes/reader.go:91
	_go_fuzz_dep_.CoverTab[842]++
						r.prevRune = int(r.i)
						if c := r.s[r.i]; c < utf8.RuneSelf {
//line /snap/go/10455/src/bytes/reader.go:93
		_go_fuzz_dep_.CoverTab[524911]++
//line /snap/go/10455/src/bytes/reader.go:93
		_go_fuzz_dep_.CoverTab[846]++
							r.i++
							return rune(c), 1, nil
//line /snap/go/10455/src/bytes/reader.go:95
		// _ = "end of CoverTab[846]"
	} else {
//line /snap/go/10455/src/bytes/reader.go:96
		_go_fuzz_dep_.CoverTab[524912]++
//line /snap/go/10455/src/bytes/reader.go:96
		_go_fuzz_dep_.CoverTab[847]++
//line /snap/go/10455/src/bytes/reader.go:96
		// _ = "end of CoverTab[847]"
//line /snap/go/10455/src/bytes/reader.go:96
	}
//line /snap/go/10455/src/bytes/reader.go:96
	// _ = "end of CoverTab[842]"
//line /snap/go/10455/src/bytes/reader.go:96
	_go_fuzz_dep_.CoverTab[843]++
						ch, size = utf8.DecodeRune(r.s[r.i:])
						r.i += int64(size)
						return
//line /snap/go/10455/src/bytes/reader.go:99
	// _ = "end of CoverTab[843]"
}

// UnreadRune complements ReadRune in implementing the io.RuneScanner interface.
func (r *Reader) UnreadRune() error {
//line /snap/go/10455/src/bytes/reader.go:103
	_go_fuzz_dep_.CoverTab[848]++
						if r.i <= 0 {
//line /snap/go/10455/src/bytes/reader.go:104
		_go_fuzz_dep_.CoverTab[524913]++
//line /snap/go/10455/src/bytes/reader.go:104
		_go_fuzz_dep_.CoverTab[851]++
							return errors.New("bytes.Reader.UnreadRune: at beginning of slice")
//line /snap/go/10455/src/bytes/reader.go:105
		// _ = "end of CoverTab[851]"
	} else {
//line /snap/go/10455/src/bytes/reader.go:106
		_go_fuzz_dep_.CoverTab[524914]++
//line /snap/go/10455/src/bytes/reader.go:106
		_go_fuzz_dep_.CoverTab[852]++
//line /snap/go/10455/src/bytes/reader.go:106
		// _ = "end of CoverTab[852]"
//line /snap/go/10455/src/bytes/reader.go:106
	}
//line /snap/go/10455/src/bytes/reader.go:106
	// _ = "end of CoverTab[848]"
//line /snap/go/10455/src/bytes/reader.go:106
	_go_fuzz_dep_.CoverTab[849]++
						if r.prevRune < 0 {
//line /snap/go/10455/src/bytes/reader.go:107
		_go_fuzz_dep_.CoverTab[524915]++
//line /snap/go/10455/src/bytes/reader.go:107
		_go_fuzz_dep_.CoverTab[853]++
							return errors.New("bytes.Reader.UnreadRune: previous operation was not ReadRune")
//line /snap/go/10455/src/bytes/reader.go:108
		// _ = "end of CoverTab[853]"
	} else {
//line /snap/go/10455/src/bytes/reader.go:109
		_go_fuzz_dep_.CoverTab[524916]++
//line /snap/go/10455/src/bytes/reader.go:109
		_go_fuzz_dep_.CoverTab[854]++
//line /snap/go/10455/src/bytes/reader.go:109
		// _ = "end of CoverTab[854]"
//line /snap/go/10455/src/bytes/reader.go:109
	}
//line /snap/go/10455/src/bytes/reader.go:109
	// _ = "end of CoverTab[849]"
//line /snap/go/10455/src/bytes/reader.go:109
	_go_fuzz_dep_.CoverTab[850]++
						r.i = int64(r.prevRune)
						r.prevRune = -1
						return nil
//line /snap/go/10455/src/bytes/reader.go:112
	// _ = "end of CoverTab[850]"
}

// Seek implements the io.Seeker interface.
func (r *Reader) Seek(offset int64, whence int) (int64, error) {
//line /snap/go/10455/src/bytes/reader.go:116
	_go_fuzz_dep_.CoverTab[855]++
						r.prevRune = -1
						var abs int64
						switch whence {
	case io.SeekStart:
//line /snap/go/10455/src/bytes/reader.go:120
		_go_fuzz_dep_.CoverTab[524917]++
//line /snap/go/10455/src/bytes/reader.go:120
		_go_fuzz_dep_.CoverTab[858]++
							abs = offset
//line /snap/go/10455/src/bytes/reader.go:121
		// _ = "end of CoverTab[858]"
	case io.SeekCurrent:
//line /snap/go/10455/src/bytes/reader.go:122
		_go_fuzz_dep_.CoverTab[524918]++
//line /snap/go/10455/src/bytes/reader.go:122
		_go_fuzz_dep_.CoverTab[859]++
							abs = r.i + offset
//line /snap/go/10455/src/bytes/reader.go:123
		// _ = "end of CoverTab[859]"
	case io.SeekEnd:
//line /snap/go/10455/src/bytes/reader.go:124
		_go_fuzz_dep_.CoverTab[524919]++
//line /snap/go/10455/src/bytes/reader.go:124
		_go_fuzz_dep_.CoverTab[860]++
							abs = int64(len(r.s)) + offset
//line /snap/go/10455/src/bytes/reader.go:125
		// _ = "end of CoverTab[860]"
	default:
//line /snap/go/10455/src/bytes/reader.go:126
		_go_fuzz_dep_.CoverTab[524920]++
//line /snap/go/10455/src/bytes/reader.go:126
		_go_fuzz_dep_.CoverTab[861]++
							return 0, errors.New("bytes.Reader.Seek: invalid whence")
//line /snap/go/10455/src/bytes/reader.go:127
		// _ = "end of CoverTab[861]"
	}
//line /snap/go/10455/src/bytes/reader.go:128
	// _ = "end of CoverTab[855]"
//line /snap/go/10455/src/bytes/reader.go:128
	_go_fuzz_dep_.CoverTab[856]++
						if abs < 0 {
//line /snap/go/10455/src/bytes/reader.go:129
		_go_fuzz_dep_.CoverTab[524921]++
//line /snap/go/10455/src/bytes/reader.go:129
		_go_fuzz_dep_.CoverTab[862]++
							return 0, errors.New("bytes.Reader.Seek: negative position")
//line /snap/go/10455/src/bytes/reader.go:130
		// _ = "end of CoverTab[862]"
	} else {
//line /snap/go/10455/src/bytes/reader.go:131
		_go_fuzz_dep_.CoverTab[524922]++
//line /snap/go/10455/src/bytes/reader.go:131
		_go_fuzz_dep_.CoverTab[863]++
//line /snap/go/10455/src/bytes/reader.go:131
		// _ = "end of CoverTab[863]"
//line /snap/go/10455/src/bytes/reader.go:131
	}
//line /snap/go/10455/src/bytes/reader.go:131
	// _ = "end of CoverTab[856]"
//line /snap/go/10455/src/bytes/reader.go:131
	_go_fuzz_dep_.CoverTab[857]++
						r.i = abs
						return abs, nil
//line /snap/go/10455/src/bytes/reader.go:133
	// _ = "end of CoverTab[857]"
}

// WriteTo implements the io.WriterTo interface.
func (r *Reader) WriteTo(w io.Writer) (n int64, err error) {
//line /snap/go/10455/src/bytes/reader.go:137
	_go_fuzz_dep_.CoverTab[864]++
						r.prevRune = -1
						if r.i >= int64(len(r.s)) {
//line /snap/go/10455/src/bytes/reader.go:139
		_go_fuzz_dep_.CoverTab[524923]++
//line /snap/go/10455/src/bytes/reader.go:139
		_go_fuzz_dep_.CoverTab[868]++
							return 0, nil
//line /snap/go/10455/src/bytes/reader.go:140
		// _ = "end of CoverTab[868]"
	} else {
//line /snap/go/10455/src/bytes/reader.go:141
		_go_fuzz_dep_.CoverTab[524924]++
//line /snap/go/10455/src/bytes/reader.go:141
		_go_fuzz_dep_.CoverTab[869]++
//line /snap/go/10455/src/bytes/reader.go:141
		// _ = "end of CoverTab[869]"
//line /snap/go/10455/src/bytes/reader.go:141
	}
//line /snap/go/10455/src/bytes/reader.go:141
	// _ = "end of CoverTab[864]"
//line /snap/go/10455/src/bytes/reader.go:141
	_go_fuzz_dep_.CoverTab[865]++
						b := r.s[r.i:]
						m, err := w.Write(b)
						if m > len(b) {
//line /snap/go/10455/src/bytes/reader.go:144
		_go_fuzz_dep_.CoverTab[524925]++
//line /snap/go/10455/src/bytes/reader.go:144
		_go_fuzz_dep_.CoverTab[870]++
							panic("bytes.Reader.WriteTo: invalid Write count")
//line /snap/go/10455/src/bytes/reader.go:145
		// _ = "end of CoverTab[870]"
	} else {
//line /snap/go/10455/src/bytes/reader.go:146
		_go_fuzz_dep_.CoverTab[524926]++
//line /snap/go/10455/src/bytes/reader.go:146
		_go_fuzz_dep_.CoverTab[871]++
//line /snap/go/10455/src/bytes/reader.go:146
		// _ = "end of CoverTab[871]"
//line /snap/go/10455/src/bytes/reader.go:146
	}
//line /snap/go/10455/src/bytes/reader.go:146
	// _ = "end of CoverTab[865]"
//line /snap/go/10455/src/bytes/reader.go:146
	_go_fuzz_dep_.CoverTab[866]++
						r.i += int64(m)
						n = int64(m)
						if m != len(b) && func() bool {
//line /snap/go/10455/src/bytes/reader.go:149
		_go_fuzz_dep_.CoverTab[872]++
//line /snap/go/10455/src/bytes/reader.go:149
		return err == nil
//line /snap/go/10455/src/bytes/reader.go:149
		// _ = "end of CoverTab[872]"
//line /snap/go/10455/src/bytes/reader.go:149
	}() {
//line /snap/go/10455/src/bytes/reader.go:149
		_go_fuzz_dep_.CoverTab[524927]++
//line /snap/go/10455/src/bytes/reader.go:149
		_go_fuzz_dep_.CoverTab[873]++
							err = io.ErrShortWrite
//line /snap/go/10455/src/bytes/reader.go:150
		// _ = "end of CoverTab[873]"
	} else {
//line /snap/go/10455/src/bytes/reader.go:151
		_go_fuzz_dep_.CoverTab[524928]++
//line /snap/go/10455/src/bytes/reader.go:151
		_go_fuzz_dep_.CoverTab[874]++
//line /snap/go/10455/src/bytes/reader.go:151
		// _ = "end of CoverTab[874]"
//line /snap/go/10455/src/bytes/reader.go:151
	}
//line /snap/go/10455/src/bytes/reader.go:151
	// _ = "end of CoverTab[866]"
//line /snap/go/10455/src/bytes/reader.go:151
	_go_fuzz_dep_.CoverTab[867]++
						return
//line /snap/go/10455/src/bytes/reader.go:152
	// _ = "end of CoverTab[867]"
}

// Reset resets the Reader to be reading from b.
func (r *Reader) Reset(b []byte) {
//line /snap/go/10455/src/bytes/reader.go:156
	_go_fuzz_dep_.CoverTab[875]++
//line /snap/go/10455/src/bytes/reader.go:156
	*r = Reader{b, 0, -1}
//line /snap/go/10455/src/bytes/reader.go:156
	// _ = "end of CoverTab[875]"
//line /snap/go/10455/src/bytes/reader.go:156
}

// NewReader returns a new Reader reading from b.
func NewReader(b []byte) *Reader {
//line /snap/go/10455/src/bytes/reader.go:159
	_go_fuzz_dep_.CoverTab[876]++
//line /snap/go/10455/src/bytes/reader.go:159
	return &Reader{b, 0, -1}
//line /snap/go/10455/src/bytes/reader.go:159
	// _ = "end of CoverTab[876]"
//line /snap/go/10455/src/bytes/reader.go:159
}

//line /snap/go/10455/src/bytes/reader.go:159
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /snap/go/10455/src/bytes/reader.go:159
var _ = _go_fuzz_dep_.CoverTab
