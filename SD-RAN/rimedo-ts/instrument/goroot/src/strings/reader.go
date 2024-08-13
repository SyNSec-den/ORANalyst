// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/strings/reader.go:5
package strings

//line /usr/local/go/src/strings/reader.go:5
import (
//line /usr/local/go/src/strings/reader.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/strings/reader.go:5
)
//line /usr/local/go/src/strings/reader.go:5
import (
//line /usr/local/go/src/strings/reader.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/strings/reader.go:5
)

import (
	"errors"
	"io"
	"unicode/utf8"
)

// A Reader implements the io.Reader, io.ReaderAt, io.ByteReader, io.ByteScanner,
//line /usr/local/go/src/strings/reader.go:13
// io.RuneReader, io.RuneScanner, io.Seeker, and io.WriterTo interfaces by reading
//line /usr/local/go/src/strings/reader.go:13
// from a string.
//line /usr/local/go/src/strings/reader.go:13
// The zero value for Reader operates like a Reader of an empty string.
//line /usr/local/go/src/strings/reader.go:17
type Reader struct {
	s		string
	i		int64	// current reading index
	prevRune	int	// index of previous rune; or < 0
}

// Len returns the number of bytes of the unread portion of the
//line /usr/local/go/src/strings/reader.go:23
// string.
//line /usr/local/go/src/strings/reader.go:25
func (r *Reader) Len() int {
//line /usr/local/go/src/strings/reader.go:25
	_go_fuzz_dep_.CoverTab[3151]++
						if r.i >= int64(len(r.s)) {
//line /usr/local/go/src/strings/reader.go:26
		_go_fuzz_dep_.CoverTab[3153]++
							return 0
//line /usr/local/go/src/strings/reader.go:27
		// _ = "end of CoverTab[3153]"
	} else {
//line /usr/local/go/src/strings/reader.go:28
		_go_fuzz_dep_.CoverTab[3154]++
//line /usr/local/go/src/strings/reader.go:28
		// _ = "end of CoverTab[3154]"
//line /usr/local/go/src/strings/reader.go:28
	}
//line /usr/local/go/src/strings/reader.go:28
	// _ = "end of CoverTab[3151]"
//line /usr/local/go/src/strings/reader.go:28
	_go_fuzz_dep_.CoverTab[3152]++
						return int(int64(len(r.s)) - r.i)
//line /usr/local/go/src/strings/reader.go:29
	// _ = "end of CoverTab[3152]"
}

// Size returns the original length of the underlying string.
//line /usr/local/go/src/strings/reader.go:32
// Size is the number of bytes available for reading via ReadAt.
//line /usr/local/go/src/strings/reader.go:32
// The returned value is always the same and is not affected by calls
//line /usr/local/go/src/strings/reader.go:32
// to any other method.
//line /usr/local/go/src/strings/reader.go:36
func (r *Reader) Size() int64 {
//line /usr/local/go/src/strings/reader.go:36
	_go_fuzz_dep_.CoverTab[3155]++
//line /usr/local/go/src/strings/reader.go:36
	return int64(len(r.s))
//line /usr/local/go/src/strings/reader.go:36
	// _ = "end of CoverTab[3155]"
//line /usr/local/go/src/strings/reader.go:36
}

// Read implements the io.Reader interface.
func (r *Reader) Read(b []byte) (n int, err error) {
//line /usr/local/go/src/strings/reader.go:39
	_go_fuzz_dep_.CoverTab[3156]++
						if r.i >= int64(len(r.s)) {
//line /usr/local/go/src/strings/reader.go:40
		_go_fuzz_dep_.CoverTab[3158]++
							return 0, io.EOF
//line /usr/local/go/src/strings/reader.go:41
		// _ = "end of CoverTab[3158]"
	} else {
//line /usr/local/go/src/strings/reader.go:42
		_go_fuzz_dep_.CoverTab[3159]++
//line /usr/local/go/src/strings/reader.go:42
		// _ = "end of CoverTab[3159]"
//line /usr/local/go/src/strings/reader.go:42
	}
//line /usr/local/go/src/strings/reader.go:42
	// _ = "end of CoverTab[3156]"
//line /usr/local/go/src/strings/reader.go:42
	_go_fuzz_dep_.CoverTab[3157]++
						r.prevRune = -1
						n = copy(b, r.s[r.i:])
						r.i += int64(n)
						return
//line /usr/local/go/src/strings/reader.go:46
	// _ = "end of CoverTab[3157]"
}

// ReadAt implements the io.ReaderAt interface.
func (r *Reader) ReadAt(b []byte, off int64) (n int, err error) {
//line /usr/local/go/src/strings/reader.go:50
	_go_fuzz_dep_.CoverTab[3160]++

						if off < 0 {
//line /usr/local/go/src/strings/reader.go:52
		_go_fuzz_dep_.CoverTab[3164]++
							return 0, errors.New("strings.Reader.ReadAt: negative offset")
//line /usr/local/go/src/strings/reader.go:53
		// _ = "end of CoverTab[3164]"
	} else {
//line /usr/local/go/src/strings/reader.go:54
		_go_fuzz_dep_.CoverTab[3165]++
//line /usr/local/go/src/strings/reader.go:54
		// _ = "end of CoverTab[3165]"
//line /usr/local/go/src/strings/reader.go:54
	}
//line /usr/local/go/src/strings/reader.go:54
	// _ = "end of CoverTab[3160]"
//line /usr/local/go/src/strings/reader.go:54
	_go_fuzz_dep_.CoverTab[3161]++
						if off >= int64(len(r.s)) {
//line /usr/local/go/src/strings/reader.go:55
		_go_fuzz_dep_.CoverTab[3166]++
							return 0, io.EOF
//line /usr/local/go/src/strings/reader.go:56
		// _ = "end of CoverTab[3166]"
	} else {
//line /usr/local/go/src/strings/reader.go:57
		_go_fuzz_dep_.CoverTab[3167]++
//line /usr/local/go/src/strings/reader.go:57
		// _ = "end of CoverTab[3167]"
//line /usr/local/go/src/strings/reader.go:57
	}
//line /usr/local/go/src/strings/reader.go:57
	// _ = "end of CoverTab[3161]"
//line /usr/local/go/src/strings/reader.go:57
	_go_fuzz_dep_.CoverTab[3162]++
						n = copy(b, r.s[off:])
						if n < len(b) {
//line /usr/local/go/src/strings/reader.go:59
		_go_fuzz_dep_.CoverTab[3168]++
							err = io.EOF
//line /usr/local/go/src/strings/reader.go:60
		// _ = "end of CoverTab[3168]"
	} else {
//line /usr/local/go/src/strings/reader.go:61
		_go_fuzz_dep_.CoverTab[3169]++
//line /usr/local/go/src/strings/reader.go:61
		// _ = "end of CoverTab[3169]"
//line /usr/local/go/src/strings/reader.go:61
	}
//line /usr/local/go/src/strings/reader.go:61
	// _ = "end of CoverTab[3162]"
//line /usr/local/go/src/strings/reader.go:61
	_go_fuzz_dep_.CoverTab[3163]++
						return
//line /usr/local/go/src/strings/reader.go:62
	// _ = "end of CoverTab[3163]"
}

// ReadByte implements the io.ByteReader interface.
func (r *Reader) ReadByte() (byte, error) {
//line /usr/local/go/src/strings/reader.go:66
	_go_fuzz_dep_.CoverTab[3170]++
						r.prevRune = -1
						if r.i >= int64(len(r.s)) {
//line /usr/local/go/src/strings/reader.go:68
		_go_fuzz_dep_.CoverTab[3172]++
							return 0, io.EOF
//line /usr/local/go/src/strings/reader.go:69
		// _ = "end of CoverTab[3172]"
	} else {
//line /usr/local/go/src/strings/reader.go:70
		_go_fuzz_dep_.CoverTab[3173]++
//line /usr/local/go/src/strings/reader.go:70
		// _ = "end of CoverTab[3173]"
//line /usr/local/go/src/strings/reader.go:70
	}
//line /usr/local/go/src/strings/reader.go:70
	// _ = "end of CoverTab[3170]"
//line /usr/local/go/src/strings/reader.go:70
	_go_fuzz_dep_.CoverTab[3171]++
						b := r.s[r.i]
						r.i++
						return b, nil
//line /usr/local/go/src/strings/reader.go:73
	// _ = "end of CoverTab[3171]"
}

// UnreadByte implements the io.ByteScanner interface.
func (r *Reader) UnreadByte() error {
//line /usr/local/go/src/strings/reader.go:77
	_go_fuzz_dep_.CoverTab[3174]++
						if r.i <= 0 {
//line /usr/local/go/src/strings/reader.go:78
		_go_fuzz_dep_.CoverTab[3176]++
							return errors.New("strings.Reader.UnreadByte: at beginning of string")
//line /usr/local/go/src/strings/reader.go:79
		// _ = "end of CoverTab[3176]"
	} else {
//line /usr/local/go/src/strings/reader.go:80
		_go_fuzz_dep_.CoverTab[3177]++
//line /usr/local/go/src/strings/reader.go:80
		// _ = "end of CoverTab[3177]"
//line /usr/local/go/src/strings/reader.go:80
	}
//line /usr/local/go/src/strings/reader.go:80
	// _ = "end of CoverTab[3174]"
//line /usr/local/go/src/strings/reader.go:80
	_go_fuzz_dep_.CoverTab[3175]++
						r.prevRune = -1
						r.i--
						return nil
//line /usr/local/go/src/strings/reader.go:83
	// _ = "end of CoverTab[3175]"
}

// ReadRune implements the io.RuneReader interface.
func (r *Reader) ReadRune() (ch rune, size int, err error) {
//line /usr/local/go/src/strings/reader.go:87
	_go_fuzz_dep_.CoverTab[3178]++
						if r.i >= int64(len(r.s)) {
//line /usr/local/go/src/strings/reader.go:88
		_go_fuzz_dep_.CoverTab[3181]++
							r.prevRune = -1
							return 0, 0, io.EOF
//line /usr/local/go/src/strings/reader.go:90
		// _ = "end of CoverTab[3181]"
	} else {
//line /usr/local/go/src/strings/reader.go:91
		_go_fuzz_dep_.CoverTab[3182]++
//line /usr/local/go/src/strings/reader.go:91
		// _ = "end of CoverTab[3182]"
//line /usr/local/go/src/strings/reader.go:91
	}
//line /usr/local/go/src/strings/reader.go:91
	// _ = "end of CoverTab[3178]"
//line /usr/local/go/src/strings/reader.go:91
	_go_fuzz_dep_.CoverTab[3179]++
						r.prevRune = int(r.i)
						if c := r.s[r.i]; c < utf8.RuneSelf {
//line /usr/local/go/src/strings/reader.go:93
		_go_fuzz_dep_.CoverTab[3183]++
							r.i++
							return rune(c), 1, nil
//line /usr/local/go/src/strings/reader.go:95
		// _ = "end of CoverTab[3183]"
	} else {
//line /usr/local/go/src/strings/reader.go:96
		_go_fuzz_dep_.CoverTab[3184]++
//line /usr/local/go/src/strings/reader.go:96
		// _ = "end of CoverTab[3184]"
//line /usr/local/go/src/strings/reader.go:96
	}
//line /usr/local/go/src/strings/reader.go:96
	// _ = "end of CoverTab[3179]"
//line /usr/local/go/src/strings/reader.go:96
	_go_fuzz_dep_.CoverTab[3180]++
						ch, size = utf8.DecodeRuneInString(r.s[r.i:])
						r.i += int64(size)
						return
//line /usr/local/go/src/strings/reader.go:99
	// _ = "end of CoverTab[3180]"
}

// UnreadRune implements the io.RuneScanner interface.
func (r *Reader) UnreadRune() error {
//line /usr/local/go/src/strings/reader.go:103
	_go_fuzz_dep_.CoverTab[3185]++
						if r.i <= 0 {
//line /usr/local/go/src/strings/reader.go:104
		_go_fuzz_dep_.CoverTab[3188]++
							return errors.New("strings.Reader.UnreadRune: at beginning of string")
//line /usr/local/go/src/strings/reader.go:105
		// _ = "end of CoverTab[3188]"
	} else {
//line /usr/local/go/src/strings/reader.go:106
		_go_fuzz_dep_.CoverTab[3189]++
//line /usr/local/go/src/strings/reader.go:106
		// _ = "end of CoverTab[3189]"
//line /usr/local/go/src/strings/reader.go:106
	}
//line /usr/local/go/src/strings/reader.go:106
	// _ = "end of CoverTab[3185]"
//line /usr/local/go/src/strings/reader.go:106
	_go_fuzz_dep_.CoverTab[3186]++
						if r.prevRune < 0 {
//line /usr/local/go/src/strings/reader.go:107
		_go_fuzz_dep_.CoverTab[3190]++
							return errors.New("strings.Reader.UnreadRune: previous operation was not ReadRune")
//line /usr/local/go/src/strings/reader.go:108
		// _ = "end of CoverTab[3190]"
	} else {
//line /usr/local/go/src/strings/reader.go:109
		_go_fuzz_dep_.CoverTab[3191]++
//line /usr/local/go/src/strings/reader.go:109
		// _ = "end of CoverTab[3191]"
//line /usr/local/go/src/strings/reader.go:109
	}
//line /usr/local/go/src/strings/reader.go:109
	// _ = "end of CoverTab[3186]"
//line /usr/local/go/src/strings/reader.go:109
	_go_fuzz_dep_.CoverTab[3187]++
						r.i = int64(r.prevRune)
						r.prevRune = -1
						return nil
//line /usr/local/go/src/strings/reader.go:112
	// _ = "end of CoverTab[3187]"
}

// Seek implements the io.Seeker interface.
func (r *Reader) Seek(offset int64, whence int) (int64, error) {
//line /usr/local/go/src/strings/reader.go:116
	_go_fuzz_dep_.CoverTab[3192]++
						r.prevRune = -1
						var abs int64
						switch whence {
	case io.SeekStart:
//line /usr/local/go/src/strings/reader.go:120
		_go_fuzz_dep_.CoverTab[3195]++
							abs = offset
//line /usr/local/go/src/strings/reader.go:121
		// _ = "end of CoverTab[3195]"
	case io.SeekCurrent:
//line /usr/local/go/src/strings/reader.go:122
		_go_fuzz_dep_.CoverTab[3196]++
							abs = r.i + offset
//line /usr/local/go/src/strings/reader.go:123
		// _ = "end of CoverTab[3196]"
	case io.SeekEnd:
//line /usr/local/go/src/strings/reader.go:124
		_go_fuzz_dep_.CoverTab[3197]++
							abs = int64(len(r.s)) + offset
//line /usr/local/go/src/strings/reader.go:125
		// _ = "end of CoverTab[3197]"
	default:
//line /usr/local/go/src/strings/reader.go:126
		_go_fuzz_dep_.CoverTab[3198]++
							return 0, errors.New("strings.Reader.Seek: invalid whence")
//line /usr/local/go/src/strings/reader.go:127
		// _ = "end of CoverTab[3198]"
	}
//line /usr/local/go/src/strings/reader.go:128
	// _ = "end of CoverTab[3192]"
//line /usr/local/go/src/strings/reader.go:128
	_go_fuzz_dep_.CoverTab[3193]++
						if abs < 0 {
//line /usr/local/go/src/strings/reader.go:129
		_go_fuzz_dep_.CoverTab[3199]++
							return 0, errors.New("strings.Reader.Seek: negative position")
//line /usr/local/go/src/strings/reader.go:130
		// _ = "end of CoverTab[3199]"
	} else {
//line /usr/local/go/src/strings/reader.go:131
		_go_fuzz_dep_.CoverTab[3200]++
//line /usr/local/go/src/strings/reader.go:131
		// _ = "end of CoverTab[3200]"
//line /usr/local/go/src/strings/reader.go:131
	}
//line /usr/local/go/src/strings/reader.go:131
	// _ = "end of CoverTab[3193]"
//line /usr/local/go/src/strings/reader.go:131
	_go_fuzz_dep_.CoverTab[3194]++
						r.i = abs
						return abs, nil
//line /usr/local/go/src/strings/reader.go:133
	// _ = "end of CoverTab[3194]"
}

// WriteTo implements the io.WriterTo interface.
func (r *Reader) WriteTo(w io.Writer) (n int64, err error) {
//line /usr/local/go/src/strings/reader.go:137
	_go_fuzz_dep_.CoverTab[3201]++
						r.prevRune = -1
						if r.i >= int64(len(r.s)) {
//line /usr/local/go/src/strings/reader.go:139
		_go_fuzz_dep_.CoverTab[3205]++
							return 0, nil
//line /usr/local/go/src/strings/reader.go:140
		// _ = "end of CoverTab[3205]"
	} else {
//line /usr/local/go/src/strings/reader.go:141
		_go_fuzz_dep_.CoverTab[3206]++
//line /usr/local/go/src/strings/reader.go:141
		// _ = "end of CoverTab[3206]"
//line /usr/local/go/src/strings/reader.go:141
	}
//line /usr/local/go/src/strings/reader.go:141
	// _ = "end of CoverTab[3201]"
//line /usr/local/go/src/strings/reader.go:141
	_go_fuzz_dep_.CoverTab[3202]++
						s := r.s[r.i:]
						m, err := io.WriteString(w, s)
						if m > len(s) {
//line /usr/local/go/src/strings/reader.go:144
		_go_fuzz_dep_.CoverTab[3207]++
							panic("strings.Reader.WriteTo: invalid WriteString count")
//line /usr/local/go/src/strings/reader.go:145
		// _ = "end of CoverTab[3207]"
	} else {
//line /usr/local/go/src/strings/reader.go:146
		_go_fuzz_dep_.CoverTab[3208]++
//line /usr/local/go/src/strings/reader.go:146
		// _ = "end of CoverTab[3208]"
//line /usr/local/go/src/strings/reader.go:146
	}
//line /usr/local/go/src/strings/reader.go:146
	// _ = "end of CoverTab[3202]"
//line /usr/local/go/src/strings/reader.go:146
	_go_fuzz_dep_.CoverTab[3203]++
						r.i += int64(m)
						n = int64(m)
						if m != len(s) && func() bool {
//line /usr/local/go/src/strings/reader.go:149
		_go_fuzz_dep_.CoverTab[3209]++
//line /usr/local/go/src/strings/reader.go:149
		return err == nil
//line /usr/local/go/src/strings/reader.go:149
		// _ = "end of CoverTab[3209]"
//line /usr/local/go/src/strings/reader.go:149
	}() {
//line /usr/local/go/src/strings/reader.go:149
		_go_fuzz_dep_.CoverTab[3210]++
							err = io.ErrShortWrite
//line /usr/local/go/src/strings/reader.go:150
		// _ = "end of CoverTab[3210]"
	} else {
//line /usr/local/go/src/strings/reader.go:151
		_go_fuzz_dep_.CoverTab[3211]++
//line /usr/local/go/src/strings/reader.go:151
		// _ = "end of CoverTab[3211]"
//line /usr/local/go/src/strings/reader.go:151
	}
//line /usr/local/go/src/strings/reader.go:151
	// _ = "end of CoverTab[3203]"
//line /usr/local/go/src/strings/reader.go:151
	_go_fuzz_dep_.CoverTab[3204]++
						return
//line /usr/local/go/src/strings/reader.go:152
	// _ = "end of CoverTab[3204]"
}

// Reset resets the Reader to be reading from s.
func (r *Reader) Reset(s string) {
//line /usr/local/go/src/strings/reader.go:156
	_go_fuzz_dep_.CoverTab[3212]++
//line /usr/local/go/src/strings/reader.go:156
	*r = Reader{s, 0, -1}
//line /usr/local/go/src/strings/reader.go:156
	// _ = "end of CoverTab[3212]"
//line /usr/local/go/src/strings/reader.go:156
}

// NewReader returns a new Reader reading from s.
//line /usr/local/go/src/strings/reader.go:158
// It is similar to bytes.NewBufferString but more efficient and read-only.
//line /usr/local/go/src/strings/reader.go:160
func NewReader(s string) *Reader {
//line /usr/local/go/src/strings/reader.go:160
	_go_fuzz_dep_.CoverTab[3213]++
//line /usr/local/go/src/strings/reader.go:160
	return &Reader{s, 0, -1}
//line /usr/local/go/src/strings/reader.go:160
	// _ = "end of CoverTab[3213]"
//line /usr/local/go/src/strings/reader.go:160
}

//line /usr/local/go/src/strings/reader.go:160
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/strings/reader.go:160
var _ = _go_fuzz_dep_.CoverTab
