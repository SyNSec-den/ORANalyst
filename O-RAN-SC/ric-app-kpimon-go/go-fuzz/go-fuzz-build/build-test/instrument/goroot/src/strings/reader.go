// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /snap/go/10455/src/strings/reader.go:5
package strings

//line /snap/go/10455/src/strings/reader.go:5
import (
//line /snap/go/10455/src/strings/reader.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /snap/go/10455/src/strings/reader.go:5
)
//line /snap/go/10455/src/strings/reader.go:5
import (
//line /snap/go/10455/src/strings/reader.go:5
	_atomic_ "sync/atomic"
//line /snap/go/10455/src/strings/reader.go:5
)

import (
	"errors"
	"io"
	"unicode/utf8"
)

// A Reader implements the io.Reader, io.ReaderAt, io.ByteReader, io.ByteScanner,
//line /snap/go/10455/src/strings/reader.go:13
// io.RuneReader, io.RuneScanner, io.Seeker, and io.WriterTo interfaces by reading
//line /snap/go/10455/src/strings/reader.go:13
// from a string.
//line /snap/go/10455/src/strings/reader.go:13
// The zero value for Reader operates like a Reader of an empty string.
//line /snap/go/10455/src/strings/reader.go:17
type Reader struct {
	s		string
	i		int64	// current reading index
	prevRune	int	// index of previous rune; or < 0
}

// Len returns the number of bytes of the unread portion of the
//line /snap/go/10455/src/strings/reader.go:23
// string.
//line /snap/go/10455/src/strings/reader.go:25
func (r *Reader) Len() int {
//line /snap/go/10455/src/strings/reader.go:25
	_go_fuzz_dep_.CoverTab[909]++
						if r.i >= int64(len(r.s)) {
//line /snap/go/10455/src/strings/reader.go:26
		_go_fuzz_dep_.CoverTab[524943]++
//line /snap/go/10455/src/strings/reader.go:26
		_go_fuzz_dep_.CoverTab[911]++
							return 0
//line /snap/go/10455/src/strings/reader.go:27
		// _ = "end of CoverTab[911]"
	} else {
//line /snap/go/10455/src/strings/reader.go:28
		_go_fuzz_dep_.CoverTab[524944]++
//line /snap/go/10455/src/strings/reader.go:28
		_go_fuzz_dep_.CoverTab[912]++
//line /snap/go/10455/src/strings/reader.go:28
		// _ = "end of CoverTab[912]"
//line /snap/go/10455/src/strings/reader.go:28
	}
//line /snap/go/10455/src/strings/reader.go:28
	// _ = "end of CoverTab[909]"
//line /snap/go/10455/src/strings/reader.go:28
	_go_fuzz_dep_.CoverTab[910]++
						return int(int64(len(r.s)) - r.i)
//line /snap/go/10455/src/strings/reader.go:29
	// _ = "end of CoverTab[910]"
}

// Size returns the original length of the underlying string.
//line /snap/go/10455/src/strings/reader.go:32
// Size is the number of bytes available for reading via ReadAt.
//line /snap/go/10455/src/strings/reader.go:32
// The returned value is always the same and is not affected by calls
//line /snap/go/10455/src/strings/reader.go:32
// to any other method.
//line /snap/go/10455/src/strings/reader.go:36
func (r *Reader) Size() int64 {
//line /snap/go/10455/src/strings/reader.go:36
	_go_fuzz_dep_.CoverTab[913]++
//line /snap/go/10455/src/strings/reader.go:36
	return int64(len(r.s))
//line /snap/go/10455/src/strings/reader.go:36
	// _ = "end of CoverTab[913]"
//line /snap/go/10455/src/strings/reader.go:36
}

// Read implements the io.Reader interface.
func (r *Reader) Read(b []byte) (n int, err error) {
//line /snap/go/10455/src/strings/reader.go:39
	_go_fuzz_dep_.CoverTab[914]++
						if r.i >= int64(len(r.s)) {
//line /snap/go/10455/src/strings/reader.go:40
		_go_fuzz_dep_.CoverTab[524945]++
//line /snap/go/10455/src/strings/reader.go:40
		_go_fuzz_dep_.CoverTab[916]++
							return 0, io.EOF
//line /snap/go/10455/src/strings/reader.go:41
		// _ = "end of CoverTab[916]"
	} else {
//line /snap/go/10455/src/strings/reader.go:42
		_go_fuzz_dep_.CoverTab[524946]++
//line /snap/go/10455/src/strings/reader.go:42
		_go_fuzz_dep_.CoverTab[917]++
//line /snap/go/10455/src/strings/reader.go:42
		// _ = "end of CoverTab[917]"
//line /snap/go/10455/src/strings/reader.go:42
	}
//line /snap/go/10455/src/strings/reader.go:42
	// _ = "end of CoverTab[914]"
//line /snap/go/10455/src/strings/reader.go:42
	_go_fuzz_dep_.CoverTab[915]++
						r.prevRune = -1
						n = copy(b, r.s[r.i:])
						r.i += int64(n)
						return
//line /snap/go/10455/src/strings/reader.go:46
	// _ = "end of CoverTab[915]"
}

// ReadAt implements the io.ReaderAt interface.
func (r *Reader) ReadAt(b []byte, off int64) (n int, err error) {
//line /snap/go/10455/src/strings/reader.go:50
	_go_fuzz_dep_.CoverTab[918]++

						if off < 0 {
//line /snap/go/10455/src/strings/reader.go:52
		_go_fuzz_dep_.CoverTab[524947]++
//line /snap/go/10455/src/strings/reader.go:52
		_go_fuzz_dep_.CoverTab[922]++
							return 0, errors.New("strings.Reader.ReadAt: negative offset")
//line /snap/go/10455/src/strings/reader.go:53
		// _ = "end of CoverTab[922]"
	} else {
//line /snap/go/10455/src/strings/reader.go:54
		_go_fuzz_dep_.CoverTab[524948]++
//line /snap/go/10455/src/strings/reader.go:54
		_go_fuzz_dep_.CoverTab[923]++
//line /snap/go/10455/src/strings/reader.go:54
		// _ = "end of CoverTab[923]"
//line /snap/go/10455/src/strings/reader.go:54
	}
//line /snap/go/10455/src/strings/reader.go:54
	// _ = "end of CoverTab[918]"
//line /snap/go/10455/src/strings/reader.go:54
	_go_fuzz_dep_.CoverTab[919]++
						if off >= int64(len(r.s)) {
//line /snap/go/10455/src/strings/reader.go:55
		_go_fuzz_dep_.CoverTab[524949]++
//line /snap/go/10455/src/strings/reader.go:55
		_go_fuzz_dep_.CoverTab[924]++
							return 0, io.EOF
//line /snap/go/10455/src/strings/reader.go:56
		// _ = "end of CoverTab[924]"
	} else {
//line /snap/go/10455/src/strings/reader.go:57
		_go_fuzz_dep_.CoverTab[524950]++
//line /snap/go/10455/src/strings/reader.go:57
		_go_fuzz_dep_.CoverTab[925]++
//line /snap/go/10455/src/strings/reader.go:57
		// _ = "end of CoverTab[925]"
//line /snap/go/10455/src/strings/reader.go:57
	}
//line /snap/go/10455/src/strings/reader.go:57
	// _ = "end of CoverTab[919]"
//line /snap/go/10455/src/strings/reader.go:57
	_go_fuzz_dep_.CoverTab[920]++
						n = copy(b, r.s[off:])
						if n < len(b) {
//line /snap/go/10455/src/strings/reader.go:59
		_go_fuzz_dep_.CoverTab[524951]++
//line /snap/go/10455/src/strings/reader.go:59
		_go_fuzz_dep_.CoverTab[926]++
							err = io.EOF
//line /snap/go/10455/src/strings/reader.go:60
		// _ = "end of CoverTab[926]"
	} else {
//line /snap/go/10455/src/strings/reader.go:61
		_go_fuzz_dep_.CoverTab[524952]++
//line /snap/go/10455/src/strings/reader.go:61
		_go_fuzz_dep_.CoverTab[927]++
//line /snap/go/10455/src/strings/reader.go:61
		// _ = "end of CoverTab[927]"
//line /snap/go/10455/src/strings/reader.go:61
	}
//line /snap/go/10455/src/strings/reader.go:61
	// _ = "end of CoverTab[920]"
//line /snap/go/10455/src/strings/reader.go:61
	_go_fuzz_dep_.CoverTab[921]++
						return
//line /snap/go/10455/src/strings/reader.go:62
	// _ = "end of CoverTab[921]"
}

// ReadByte implements the io.ByteReader interface.
func (r *Reader) ReadByte() (byte, error) {
//line /snap/go/10455/src/strings/reader.go:66
	_go_fuzz_dep_.CoverTab[928]++
						r.prevRune = -1
						if r.i >= int64(len(r.s)) {
//line /snap/go/10455/src/strings/reader.go:68
		_go_fuzz_dep_.CoverTab[524953]++
//line /snap/go/10455/src/strings/reader.go:68
		_go_fuzz_dep_.CoverTab[930]++
							return 0, io.EOF
//line /snap/go/10455/src/strings/reader.go:69
		// _ = "end of CoverTab[930]"
	} else {
//line /snap/go/10455/src/strings/reader.go:70
		_go_fuzz_dep_.CoverTab[524954]++
//line /snap/go/10455/src/strings/reader.go:70
		_go_fuzz_dep_.CoverTab[931]++
//line /snap/go/10455/src/strings/reader.go:70
		// _ = "end of CoverTab[931]"
//line /snap/go/10455/src/strings/reader.go:70
	}
//line /snap/go/10455/src/strings/reader.go:70
	// _ = "end of CoverTab[928]"
//line /snap/go/10455/src/strings/reader.go:70
	_go_fuzz_dep_.CoverTab[929]++
						b := r.s[r.i]
						r.i++
						return b, nil
//line /snap/go/10455/src/strings/reader.go:73
	// _ = "end of CoverTab[929]"
}

// UnreadByte implements the io.ByteScanner interface.
func (r *Reader) UnreadByte() error {
//line /snap/go/10455/src/strings/reader.go:77
	_go_fuzz_dep_.CoverTab[932]++
						if r.i <= 0 {
//line /snap/go/10455/src/strings/reader.go:78
		_go_fuzz_dep_.CoverTab[524955]++
//line /snap/go/10455/src/strings/reader.go:78
		_go_fuzz_dep_.CoverTab[934]++
							return errors.New("strings.Reader.UnreadByte: at beginning of string")
//line /snap/go/10455/src/strings/reader.go:79
		// _ = "end of CoverTab[934]"
	} else {
//line /snap/go/10455/src/strings/reader.go:80
		_go_fuzz_dep_.CoverTab[524956]++
//line /snap/go/10455/src/strings/reader.go:80
		_go_fuzz_dep_.CoverTab[935]++
//line /snap/go/10455/src/strings/reader.go:80
		// _ = "end of CoverTab[935]"
//line /snap/go/10455/src/strings/reader.go:80
	}
//line /snap/go/10455/src/strings/reader.go:80
	// _ = "end of CoverTab[932]"
//line /snap/go/10455/src/strings/reader.go:80
	_go_fuzz_dep_.CoverTab[933]++
						r.prevRune = -1
						r.i--
						return nil
//line /snap/go/10455/src/strings/reader.go:83
	// _ = "end of CoverTab[933]"
}

// ReadRune implements the io.RuneReader interface.
func (r *Reader) ReadRune() (ch rune, size int, err error) {
//line /snap/go/10455/src/strings/reader.go:87
	_go_fuzz_dep_.CoverTab[936]++
						if r.i >= int64(len(r.s)) {
//line /snap/go/10455/src/strings/reader.go:88
		_go_fuzz_dep_.CoverTab[524957]++
//line /snap/go/10455/src/strings/reader.go:88
		_go_fuzz_dep_.CoverTab[939]++
							r.prevRune = -1
							return 0, 0, io.EOF
//line /snap/go/10455/src/strings/reader.go:90
		// _ = "end of CoverTab[939]"
	} else {
//line /snap/go/10455/src/strings/reader.go:91
		_go_fuzz_dep_.CoverTab[524958]++
//line /snap/go/10455/src/strings/reader.go:91
		_go_fuzz_dep_.CoverTab[940]++
//line /snap/go/10455/src/strings/reader.go:91
		// _ = "end of CoverTab[940]"
//line /snap/go/10455/src/strings/reader.go:91
	}
//line /snap/go/10455/src/strings/reader.go:91
	// _ = "end of CoverTab[936]"
//line /snap/go/10455/src/strings/reader.go:91
	_go_fuzz_dep_.CoverTab[937]++
						r.prevRune = int(r.i)
						if c := r.s[r.i]; c < utf8.RuneSelf {
//line /snap/go/10455/src/strings/reader.go:93
		_go_fuzz_dep_.CoverTab[524959]++
//line /snap/go/10455/src/strings/reader.go:93
		_go_fuzz_dep_.CoverTab[941]++
							r.i++
							return rune(c), 1, nil
//line /snap/go/10455/src/strings/reader.go:95
		// _ = "end of CoverTab[941]"
	} else {
//line /snap/go/10455/src/strings/reader.go:96
		_go_fuzz_dep_.CoverTab[524960]++
//line /snap/go/10455/src/strings/reader.go:96
		_go_fuzz_dep_.CoverTab[942]++
//line /snap/go/10455/src/strings/reader.go:96
		// _ = "end of CoverTab[942]"
//line /snap/go/10455/src/strings/reader.go:96
	}
//line /snap/go/10455/src/strings/reader.go:96
	// _ = "end of CoverTab[937]"
//line /snap/go/10455/src/strings/reader.go:96
	_go_fuzz_dep_.CoverTab[938]++
						ch, size = utf8.DecodeRuneInString(r.s[r.i:])
						r.i += int64(size)
						return
//line /snap/go/10455/src/strings/reader.go:99
	// _ = "end of CoverTab[938]"
}

// UnreadRune implements the io.RuneScanner interface.
func (r *Reader) UnreadRune() error {
//line /snap/go/10455/src/strings/reader.go:103
	_go_fuzz_dep_.CoverTab[943]++
							if r.i <= 0 {
//line /snap/go/10455/src/strings/reader.go:104
		_go_fuzz_dep_.CoverTab[524961]++
//line /snap/go/10455/src/strings/reader.go:104
		_go_fuzz_dep_.CoverTab[946]++
								return errors.New("strings.Reader.UnreadRune: at beginning of string")
//line /snap/go/10455/src/strings/reader.go:105
		// _ = "end of CoverTab[946]"
	} else {
//line /snap/go/10455/src/strings/reader.go:106
		_go_fuzz_dep_.CoverTab[524962]++
//line /snap/go/10455/src/strings/reader.go:106
		_go_fuzz_dep_.CoverTab[947]++
//line /snap/go/10455/src/strings/reader.go:106
		// _ = "end of CoverTab[947]"
//line /snap/go/10455/src/strings/reader.go:106
	}
//line /snap/go/10455/src/strings/reader.go:106
	// _ = "end of CoverTab[943]"
//line /snap/go/10455/src/strings/reader.go:106
	_go_fuzz_dep_.CoverTab[944]++
							if r.prevRune < 0 {
//line /snap/go/10455/src/strings/reader.go:107
		_go_fuzz_dep_.CoverTab[524963]++
//line /snap/go/10455/src/strings/reader.go:107
		_go_fuzz_dep_.CoverTab[948]++
								return errors.New("strings.Reader.UnreadRune: previous operation was not ReadRune")
//line /snap/go/10455/src/strings/reader.go:108
		// _ = "end of CoverTab[948]"
	} else {
//line /snap/go/10455/src/strings/reader.go:109
		_go_fuzz_dep_.CoverTab[524964]++
//line /snap/go/10455/src/strings/reader.go:109
		_go_fuzz_dep_.CoverTab[949]++
//line /snap/go/10455/src/strings/reader.go:109
		// _ = "end of CoverTab[949]"
//line /snap/go/10455/src/strings/reader.go:109
	}
//line /snap/go/10455/src/strings/reader.go:109
	// _ = "end of CoverTab[944]"
//line /snap/go/10455/src/strings/reader.go:109
	_go_fuzz_dep_.CoverTab[945]++
							r.i = int64(r.prevRune)
							r.prevRune = -1
							return nil
//line /snap/go/10455/src/strings/reader.go:112
	// _ = "end of CoverTab[945]"
}

// Seek implements the io.Seeker interface.
func (r *Reader) Seek(offset int64, whence int) (int64, error) {
//line /snap/go/10455/src/strings/reader.go:116
	_go_fuzz_dep_.CoverTab[950]++
							r.prevRune = -1
							var abs int64
							switch whence {
	case io.SeekStart:
//line /snap/go/10455/src/strings/reader.go:120
		_go_fuzz_dep_.CoverTab[524965]++
//line /snap/go/10455/src/strings/reader.go:120
		_go_fuzz_dep_.CoverTab[953]++
								abs = offset
//line /snap/go/10455/src/strings/reader.go:121
		// _ = "end of CoverTab[953]"
	case io.SeekCurrent:
//line /snap/go/10455/src/strings/reader.go:122
		_go_fuzz_dep_.CoverTab[524966]++
//line /snap/go/10455/src/strings/reader.go:122
		_go_fuzz_dep_.CoverTab[954]++
								abs = r.i + offset
//line /snap/go/10455/src/strings/reader.go:123
		// _ = "end of CoverTab[954]"
	case io.SeekEnd:
//line /snap/go/10455/src/strings/reader.go:124
		_go_fuzz_dep_.CoverTab[524967]++
//line /snap/go/10455/src/strings/reader.go:124
		_go_fuzz_dep_.CoverTab[955]++
								abs = int64(len(r.s)) + offset
//line /snap/go/10455/src/strings/reader.go:125
		// _ = "end of CoverTab[955]"
	default:
//line /snap/go/10455/src/strings/reader.go:126
		_go_fuzz_dep_.CoverTab[524968]++
//line /snap/go/10455/src/strings/reader.go:126
		_go_fuzz_dep_.CoverTab[956]++
								return 0, errors.New("strings.Reader.Seek: invalid whence")
//line /snap/go/10455/src/strings/reader.go:127
		// _ = "end of CoverTab[956]"
	}
//line /snap/go/10455/src/strings/reader.go:128
	// _ = "end of CoverTab[950]"
//line /snap/go/10455/src/strings/reader.go:128
	_go_fuzz_dep_.CoverTab[951]++
							if abs < 0 {
//line /snap/go/10455/src/strings/reader.go:129
		_go_fuzz_dep_.CoverTab[524969]++
//line /snap/go/10455/src/strings/reader.go:129
		_go_fuzz_dep_.CoverTab[957]++
								return 0, errors.New("strings.Reader.Seek: negative position")
//line /snap/go/10455/src/strings/reader.go:130
		// _ = "end of CoverTab[957]"
	} else {
//line /snap/go/10455/src/strings/reader.go:131
		_go_fuzz_dep_.CoverTab[524970]++
//line /snap/go/10455/src/strings/reader.go:131
		_go_fuzz_dep_.CoverTab[958]++
//line /snap/go/10455/src/strings/reader.go:131
		// _ = "end of CoverTab[958]"
//line /snap/go/10455/src/strings/reader.go:131
	}
//line /snap/go/10455/src/strings/reader.go:131
	// _ = "end of CoverTab[951]"
//line /snap/go/10455/src/strings/reader.go:131
	_go_fuzz_dep_.CoverTab[952]++
							r.i = abs
							return abs, nil
//line /snap/go/10455/src/strings/reader.go:133
	// _ = "end of CoverTab[952]"
}

// WriteTo implements the io.WriterTo interface.
func (r *Reader) WriteTo(w io.Writer) (n int64, err error) {
//line /snap/go/10455/src/strings/reader.go:137
	_go_fuzz_dep_.CoverTab[959]++
							r.prevRune = -1
							if r.i >= int64(len(r.s)) {
//line /snap/go/10455/src/strings/reader.go:139
		_go_fuzz_dep_.CoverTab[524971]++
//line /snap/go/10455/src/strings/reader.go:139
		_go_fuzz_dep_.CoverTab[963]++
								return 0, nil
//line /snap/go/10455/src/strings/reader.go:140
		// _ = "end of CoverTab[963]"
	} else {
//line /snap/go/10455/src/strings/reader.go:141
		_go_fuzz_dep_.CoverTab[524972]++
//line /snap/go/10455/src/strings/reader.go:141
		_go_fuzz_dep_.CoverTab[964]++
//line /snap/go/10455/src/strings/reader.go:141
		// _ = "end of CoverTab[964]"
//line /snap/go/10455/src/strings/reader.go:141
	}
//line /snap/go/10455/src/strings/reader.go:141
	// _ = "end of CoverTab[959]"
//line /snap/go/10455/src/strings/reader.go:141
	_go_fuzz_dep_.CoverTab[960]++
							s := r.s[r.i:]
							m, err := io.WriteString(w, s)
							if m > len(s) {
//line /snap/go/10455/src/strings/reader.go:144
		_go_fuzz_dep_.CoverTab[524973]++
//line /snap/go/10455/src/strings/reader.go:144
		_go_fuzz_dep_.CoverTab[965]++
								panic("strings.Reader.WriteTo: invalid WriteString count")
//line /snap/go/10455/src/strings/reader.go:145
		// _ = "end of CoverTab[965]"
	} else {
//line /snap/go/10455/src/strings/reader.go:146
		_go_fuzz_dep_.CoverTab[524974]++
//line /snap/go/10455/src/strings/reader.go:146
		_go_fuzz_dep_.CoverTab[966]++
//line /snap/go/10455/src/strings/reader.go:146
		// _ = "end of CoverTab[966]"
//line /snap/go/10455/src/strings/reader.go:146
	}
//line /snap/go/10455/src/strings/reader.go:146
	// _ = "end of CoverTab[960]"
//line /snap/go/10455/src/strings/reader.go:146
	_go_fuzz_dep_.CoverTab[961]++
							r.i += int64(m)
							n = int64(m)
							if m != len(s) && func() bool {
//line /snap/go/10455/src/strings/reader.go:149
		_go_fuzz_dep_.CoverTab[967]++
//line /snap/go/10455/src/strings/reader.go:149
		return err == nil
//line /snap/go/10455/src/strings/reader.go:149
		// _ = "end of CoverTab[967]"
//line /snap/go/10455/src/strings/reader.go:149
	}() {
//line /snap/go/10455/src/strings/reader.go:149
		_go_fuzz_dep_.CoverTab[524975]++
//line /snap/go/10455/src/strings/reader.go:149
		_go_fuzz_dep_.CoverTab[968]++
								err = io.ErrShortWrite
//line /snap/go/10455/src/strings/reader.go:150
		// _ = "end of CoverTab[968]"
	} else {
//line /snap/go/10455/src/strings/reader.go:151
		_go_fuzz_dep_.CoverTab[524976]++
//line /snap/go/10455/src/strings/reader.go:151
		_go_fuzz_dep_.CoverTab[969]++
//line /snap/go/10455/src/strings/reader.go:151
		// _ = "end of CoverTab[969]"
//line /snap/go/10455/src/strings/reader.go:151
	}
//line /snap/go/10455/src/strings/reader.go:151
	// _ = "end of CoverTab[961]"
//line /snap/go/10455/src/strings/reader.go:151
	_go_fuzz_dep_.CoverTab[962]++
							return
//line /snap/go/10455/src/strings/reader.go:152
	// _ = "end of CoverTab[962]"
}

// Reset resets the Reader to be reading from s.
func (r *Reader) Reset(s string) {
//line /snap/go/10455/src/strings/reader.go:156
	_go_fuzz_dep_.CoverTab[970]++
//line /snap/go/10455/src/strings/reader.go:156
	*r = Reader{s, 0, -1}
//line /snap/go/10455/src/strings/reader.go:156
	// _ = "end of CoverTab[970]"
//line /snap/go/10455/src/strings/reader.go:156
}

// NewReader returns a new Reader reading from s.
//line /snap/go/10455/src/strings/reader.go:158
// It is similar to bytes.NewBufferString but more efficient and non-writable.
//line /snap/go/10455/src/strings/reader.go:160
func NewReader(s string) *Reader {
//line /snap/go/10455/src/strings/reader.go:160
	_go_fuzz_dep_.CoverTab[971]++
//line /snap/go/10455/src/strings/reader.go:160
	return &Reader{s, 0, -1}
//line /snap/go/10455/src/strings/reader.go:160
	// _ = "end of CoverTab[971]"
//line /snap/go/10455/src/strings/reader.go:160
}

//line /snap/go/10455/src/strings/reader.go:160
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /snap/go/10455/src/strings/reader.go:160
var _ = _go_fuzz_dep_.CoverTab
