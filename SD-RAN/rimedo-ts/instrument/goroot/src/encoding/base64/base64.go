// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/encoding/base64/base64.go:5
// Package base64 implements base64 encoding as specified by RFC 4648.
package base64

//line /usr/local/go/src/encoding/base64/base64.go:6
import (
//line /usr/local/go/src/encoding/base64/base64.go:6
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/encoding/base64/base64.go:6
)
//line /usr/local/go/src/encoding/base64/base64.go:6
import (
//line /usr/local/go/src/encoding/base64/base64.go:6
	_atomic_ "sync/atomic"
//line /usr/local/go/src/encoding/base64/base64.go:6
)

import (
	"encoding/binary"
	"io"
	"strconv"
)

//line /usr/local/go/src/encoding/base64/base64.go:18
// An Encoding is a radix 64 encoding/decoding scheme, defined by a
//line /usr/local/go/src/encoding/base64/base64.go:18
// 64-character alphabet. The most common encoding is the "base64"
//line /usr/local/go/src/encoding/base64/base64.go:18
// encoding defined in RFC 4648 and used in MIME (RFC 2045) and PEM
//line /usr/local/go/src/encoding/base64/base64.go:18
// (RFC 1421).  RFC 4648 also defines an alternate encoding, which is
//line /usr/local/go/src/encoding/base64/base64.go:18
// the standard encoding with - and _ substituted for + and /.
//line /usr/local/go/src/encoding/base64/base64.go:23
type Encoding struct {
	encode		[64]byte
	decodeMap	[256]byte
	padChar		rune
	strict		bool
}

const (
	StdPadding		rune	= '='	// Standard padding character
	NoPadding		rune	= -1	// No padding
	decodeMapInitialize		= "" +
		"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" +
		"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" +
		"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" +
		"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" +
		"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" +
		"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" +
		"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" +
		"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" +
		"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" +
		"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" +
		"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" +
		"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" +
		"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" +
		"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" +
		"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" +
		"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff"
)

const encodeStd = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
const encodeURL = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"

// NewEncoding returns a new padded Encoding defined by the given alphabet,
//line /usr/local/go/src/encoding/base64/base64.go:55
// which must be a 64-byte string that does not contain the padding character
//line /usr/local/go/src/encoding/base64/base64.go:55
// or CR / LF ('\r', '\n').
//line /usr/local/go/src/encoding/base64/base64.go:55
// The resulting Encoding uses the default padding character ('='),
//line /usr/local/go/src/encoding/base64/base64.go:55
// which may be changed or disabled via WithPadding.
//line /usr/local/go/src/encoding/base64/base64.go:60
func NewEncoding(encoder string) *Encoding {
//line /usr/local/go/src/encoding/base64/base64.go:60
	_go_fuzz_dep_.CoverTab[10521]++
							if len(encoder) != 64 {
//line /usr/local/go/src/encoding/base64/base64.go:61
		_go_fuzz_dep_.CoverTab[10525]++
								panic("encoding alphabet is not 64-bytes long")
//line /usr/local/go/src/encoding/base64/base64.go:62
		// _ = "end of CoverTab[10525]"
	} else {
//line /usr/local/go/src/encoding/base64/base64.go:63
		_go_fuzz_dep_.CoverTab[10526]++
//line /usr/local/go/src/encoding/base64/base64.go:63
		// _ = "end of CoverTab[10526]"
//line /usr/local/go/src/encoding/base64/base64.go:63
	}
//line /usr/local/go/src/encoding/base64/base64.go:63
	// _ = "end of CoverTab[10521]"
//line /usr/local/go/src/encoding/base64/base64.go:63
	_go_fuzz_dep_.CoverTab[10522]++
							for i := 0; i < len(encoder); i++ {
//line /usr/local/go/src/encoding/base64/base64.go:64
		_go_fuzz_dep_.CoverTab[10527]++
								if encoder[i] == '\n' || func() bool {
//line /usr/local/go/src/encoding/base64/base64.go:65
			_go_fuzz_dep_.CoverTab[10528]++
//line /usr/local/go/src/encoding/base64/base64.go:65
			return encoder[i] == '\r'
//line /usr/local/go/src/encoding/base64/base64.go:65
			// _ = "end of CoverTab[10528]"
//line /usr/local/go/src/encoding/base64/base64.go:65
		}() {
//line /usr/local/go/src/encoding/base64/base64.go:65
			_go_fuzz_dep_.CoverTab[10529]++
									panic("encoding alphabet contains newline character")
//line /usr/local/go/src/encoding/base64/base64.go:66
			// _ = "end of CoverTab[10529]"
		} else {
//line /usr/local/go/src/encoding/base64/base64.go:67
			_go_fuzz_dep_.CoverTab[10530]++
//line /usr/local/go/src/encoding/base64/base64.go:67
			// _ = "end of CoverTab[10530]"
//line /usr/local/go/src/encoding/base64/base64.go:67
		}
//line /usr/local/go/src/encoding/base64/base64.go:67
		// _ = "end of CoverTab[10527]"
	}
//line /usr/local/go/src/encoding/base64/base64.go:68
	// _ = "end of CoverTab[10522]"
//line /usr/local/go/src/encoding/base64/base64.go:68
	_go_fuzz_dep_.CoverTab[10523]++

							e := new(Encoding)
							e.padChar = StdPadding
							copy(e.encode[:], encoder)
							copy(e.decodeMap[:], decodeMapInitialize)

							for i := 0; i < len(encoder); i++ {
//line /usr/local/go/src/encoding/base64/base64.go:75
		_go_fuzz_dep_.CoverTab[10531]++
								e.decodeMap[encoder[i]] = byte(i)
//line /usr/local/go/src/encoding/base64/base64.go:76
		// _ = "end of CoverTab[10531]"
	}
//line /usr/local/go/src/encoding/base64/base64.go:77
	// _ = "end of CoverTab[10523]"
//line /usr/local/go/src/encoding/base64/base64.go:77
	_go_fuzz_dep_.CoverTab[10524]++
							return e
//line /usr/local/go/src/encoding/base64/base64.go:78
	// _ = "end of CoverTab[10524]"
}

// WithPadding creates a new encoding identical to enc except
//line /usr/local/go/src/encoding/base64/base64.go:81
// with a specified padding character, or NoPadding to disable padding.
//line /usr/local/go/src/encoding/base64/base64.go:81
// The padding character must not be '\r' or '\n', must not
//line /usr/local/go/src/encoding/base64/base64.go:81
// be contained in the encoding's alphabet and must be a rune equal or
//line /usr/local/go/src/encoding/base64/base64.go:81
// below '\xff'.
//line /usr/local/go/src/encoding/base64/base64.go:86
func (enc Encoding) WithPadding(padding rune) *Encoding {
//line /usr/local/go/src/encoding/base64/base64.go:86
	_go_fuzz_dep_.CoverTab[10532]++
							if padding == '\r' || func() bool {
//line /usr/local/go/src/encoding/base64/base64.go:87
		_go_fuzz_dep_.CoverTab[10535]++
//line /usr/local/go/src/encoding/base64/base64.go:87
		return padding == '\n'
//line /usr/local/go/src/encoding/base64/base64.go:87
		// _ = "end of CoverTab[10535]"
//line /usr/local/go/src/encoding/base64/base64.go:87
	}() || func() bool {
//line /usr/local/go/src/encoding/base64/base64.go:87
		_go_fuzz_dep_.CoverTab[10536]++
//line /usr/local/go/src/encoding/base64/base64.go:87
		return padding > 0xff
//line /usr/local/go/src/encoding/base64/base64.go:87
		// _ = "end of CoverTab[10536]"
//line /usr/local/go/src/encoding/base64/base64.go:87
	}() {
//line /usr/local/go/src/encoding/base64/base64.go:87
		_go_fuzz_dep_.CoverTab[10537]++
								panic("invalid padding")
//line /usr/local/go/src/encoding/base64/base64.go:88
		// _ = "end of CoverTab[10537]"
	} else {
//line /usr/local/go/src/encoding/base64/base64.go:89
		_go_fuzz_dep_.CoverTab[10538]++
//line /usr/local/go/src/encoding/base64/base64.go:89
		// _ = "end of CoverTab[10538]"
//line /usr/local/go/src/encoding/base64/base64.go:89
	}
//line /usr/local/go/src/encoding/base64/base64.go:89
	// _ = "end of CoverTab[10532]"
//line /usr/local/go/src/encoding/base64/base64.go:89
	_go_fuzz_dep_.CoverTab[10533]++

							for i := 0; i < len(enc.encode); i++ {
//line /usr/local/go/src/encoding/base64/base64.go:91
		_go_fuzz_dep_.CoverTab[10539]++
								if rune(enc.encode[i]) == padding {
//line /usr/local/go/src/encoding/base64/base64.go:92
			_go_fuzz_dep_.CoverTab[10540]++
									panic("padding contained in alphabet")
//line /usr/local/go/src/encoding/base64/base64.go:93
			// _ = "end of CoverTab[10540]"
		} else {
//line /usr/local/go/src/encoding/base64/base64.go:94
			_go_fuzz_dep_.CoverTab[10541]++
//line /usr/local/go/src/encoding/base64/base64.go:94
			// _ = "end of CoverTab[10541]"
//line /usr/local/go/src/encoding/base64/base64.go:94
		}
//line /usr/local/go/src/encoding/base64/base64.go:94
		// _ = "end of CoverTab[10539]"
	}
//line /usr/local/go/src/encoding/base64/base64.go:95
	// _ = "end of CoverTab[10533]"
//line /usr/local/go/src/encoding/base64/base64.go:95
	_go_fuzz_dep_.CoverTab[10534]++

							enc.padChar = padding
							return &enc
//line /usr/local/go/src/encoding/base64/base64.go:98
	// _ = "end of CoverTab[10534]"
}

// Strict creates a new encoding identical to enc except with
//line /usr/local/go/src/encoding/base64/base64.go:101
// strict decoding enabled. In this mode, the decoder requires that
//line /usr/local/go/src/encoding/base64/base64.go:101
// trailing padding bits are zero, as described in RFC 4648 section 3.5.
//line /usr/local/go/src/encoding/base64/base64.go:101
//
//line /usr/local/go/src/encoding/base64/base64.go:101
// Note that the input is still malleable, as new line characters
//line /usr/local/go/src/encoding/base64/base64.go:101
// (CR and LF) are still ignored.
//line /usr/local/go/src/encoding/base64/base64.go:107
func (enc Encoding) Strict() *Encoding {
//line /usr/local/go/src/encoding/base64/base64.go:107
	_go_fuzz_dep_.CoverTab[10542]++
							enc.strict = true
							return &enc
//line /usr/local/go/src/encoding/base64/base64.go:109
	// _ = "end of CoverTab[10542]"
}

// StdEncoding is the standard base64 encoding, as defined in
//line /usr/local/go/src/encoding/base64/base64.go:112
// RFC 4648.
//line /usr/local/go/src/encoding/base64/base64.go:114
var StdEncoding = NewEncoding(encodeStd)

// URLEncoding is the alternate base64 encoding defined in RFC 4648.
//line /usr/local/go/src/encoding/base64/base64.go:116
// It is typically used in URLs and file names.
//line /usr/local/go/src/encoding/base64/base64.go:118
var URLEncoding = NewEncoding(encodeURL)

// RawStdEncoding is the standard raw, unpadded base64 encoding,
//line /usr/local/go/src/encoding/base64/base64.go:120
// as defined in RFC 4648 section 3.2.
//line /usr/local/go/src/encoding/base64/base64.go:120
// This is the same as StdEncoding but omits padding characters.
//line /usr/local/go/src/encoding/base64/base64.go:123
var RawStdEncoding = StdEncoding.WithPadding(NoPadding)

// RawURLEncoding is the unpadded alternate base64 encoding defined in RFC 4648.
//line /usr/local/go/src/encoding/base64/base64.go:125
// It is typically used in URLs and file names.
//line /usr/local/go/src/encoding/base64/base64.go:125
// This is the same as URLEncoding but omits padding characters.
//line /usr/local/go/src/encoding/base64/base64.go:128
var RawURLEncoding = URLEncoding.WithPadding(NoPadding)

//line /usr/local/go/src/encoding/base64/base64.go:134
// Encode encodes src using the encoding enc, writing
//line /usr/local/go/src/encoding/base64/base64.go:134
// EncodedLen(len(src)) bytes to dst.
//line /usr/local/go/src/encoding/base64/base64.go:134
//
//line /usr/local/go/src/encoding/base64/base64.go:134
// The encoding pads the output to a multiple of 4 bytes,
//line /usr/local/go/src/encoding/base64/base64.go:134
// so Encode is not appropriate for use on individual blocks
//line /usr/local/go/src/encoding/base64/base64.go:134
// of a large data stream. Use NewEncoder() instead.
//line /usr/local/go/src/encoding/base64/base64.go:140
func (enc *Encoding) Encode(dst, src []byte) {
//line /usr/local/go/src/encoding/base64/base64.go:140
	_go_fuzz_dep_.CoverTab[10543]++
							if len(src) == 0 {
//line /usr/local/go/src/encoding/base64/base64.go:141
		_go_fuzz_dep_.CoverTab[10548]++
								return
//line /usr/local/go/src/encoding/base64/base64.go:142
		// _ = "end of CoverTab[10548]"
	} else {
//line /usr/local/go/src/encoding/base64/base64.go:143
		_go_fuzz_dep_.CoverTab[10549]++
//line /usr/local/go/src/encoding/base64/base64.go:143
		// _ = "end of CoverTab[10549]"
//line /usr/local/go/src/encoding/base64/base64.go:143
	}
//line /usr/local/go/src/encoding/base64/base64.go:143
	// _ = "end of CoverTab[10543]"
//line /usr/local/go/src/encoding/base64/base64.go:143
	_go_fuzz_dep_.CoverTab[10544]++

//line /usr/local/go/src/encoding/base64/base64.go:147
	_ = enc.encode

	di, si := 0, 0
	n := (len(src) / 3) * 3
	for si < n {
//line /usr/local/go/src/encoding/base64/base64.go:151
		_go_fuzz_dep_.CoverTab[10550]++

								val := uint(src[si+0])<<16 | uint(src[si+1])<<8 | uint(src[si+2])

								dst[di+0] = enc.encode[val>>18&0x3F]
								dst[di+1] = enc.encode[val>>12&0x3F]
								dst[di+2] = enc.encode[val>>6&0x3F]
								dst[di+3] = enc.encode[val&0x3F]

								si += 3
								di += 4
//line /usr/local/go/src/encoding/base64/base64.go:161
		// _ = "end of CoverTab[10550]"
	}
//line /usr/local/go/src/encoding/base64/base64.go:162
	// _ = "end of CoverTab[10544]"
//line /usr/local/go/src/encoding/base64/base64.go:162
	_go_fuzz_dep_.CoverTab[10545]++

							remain := len(src) - si
							if remain == 0 {
//line /usr/local/go/src/encoding/base64/base64.go:165
		_go_fuzz_dep_.CoverTab[10551]++
								return
//line /usr/local/go/src/encoding/base64/base64.go:166
		// _ = "end of CoverTab[10551]"
	} else {
//line /usr/local/go/src/encoding/base64/base64.go:167
		_go_fuzz_dep_.CoverTab[10552]++
//line /usr/local/go/src/encoding/base64/base64.go:167
		// _ = "end of CoverTab[10552]"
//line /usr/local/go/src/encoding/base64/base64.go:167
	}
//line /usr/local/go/src/encoding/base64/base64.go:167
	// _ = "end of CoverTab[10545]"
//line /usr/local/go/src/encoding/base64/base64.go:167
	_go_fuzz_dep_.CoverTab[10546]++

							val := uint(src[si+0]) << 16
							if remain == 2 {
//line /usr/local/go/src/encoding/base64/base64.go:170
		_go_fuzz_dep_.CoverTab[10553]++
								val |= uint(src[si+1]) << 8
//line /usr/local/go/src/encoding/base64/base64.go:171
		// _ = "end of CoverTab[10553]"
	} else {
//line /usr/local/go/src/encoding/base64/base64.go:172
		_go_fuzz_dep_.CoverTab[10554]++
//line /usr/local/go/src/encoding/base64/base64.go:172
		// _ = "end of CoverTab[10554]"
//line /usr/local/go/src/encoding/base64/base64.go:172
	}
//line /usr/local/go/src/encoding/base64/base64.go:172
	// _ = "end of CoverTab[10546]"
//line /usr/local/go/src/encoding/base64/base64.go:172
	_go_fuzz_dep_.CoverTab[10547]++

							dst[di+0] = enc.encode[val>>18&0x3F]
							dst[di+1] = enc.encode[val>>12&0x3F]

							switch remain {
	case 2:
//line /usr/local/go/src/encoding/base64/base64.go:178
		_go_fuzz_dep_.CoverTab[10555]++
								dst[di+2] = enc.encode[val>>6&0x3F]
								if enc.padChar != NoPadding {
//line /usr/local/go/src/encoding/base64/base64.go:180
			_go_fuzz_dep_.CoverTab[10558]++
									dst[di+3] = byte(enc.padChar)
//line /usr/local/go/src/encoding/base64/base64.go:181
			// _ = "end of CoverTab[10558]"
		} else {
//line /usr/local/go/src/encoding/base64/base64.go:182
			_go_fuzz_dep_.CoverTab[10559]++
//line /usr/local/go/src/encoding/base64/base64.go:182
			// _ = "end of CoverTab[10559]"
//line /usr/local/go/src/encoding/base64/base64.go:182
		}
//line /usr/local/go/src/encoding/base64/base64.go:182
		// _ = "end of CoverTab[10555]"
	case 1:
//line /usr/local/go/src/encoding/base64/base64.go:183
		_go_fuzz_dep_.CoverTab[10556]++
								if enc.padChar != NoPadding {
//line /usr/local/go/src/encoding/base64/base64.go:184
			_go_fuzz_dep_.CoverTab[10560]++
									dst[di+2] = byte(enc.padChar)
									dst[di+3] = byte(enc.padChar)
//line /usr/local/go/src/encoding/base64/base64.go:186
			// _ = "end of CoverTab[10560]"
		} else {
//line /usr/local/go/src/encoding/base64/base64.go:187
			_go_fuzz_dep_.CoverTab[10561]++
//line /usr/local/go/src/encoding/base64/base64.go:187
			// _ = "end of CoverTab[10561]"
//line /usr/local/go/src/encoding/base64/base64.go:187
		}
//line /usr/local/go/src/encoding/base64/base64.go:187
		// _ = "end of CoverTab[10556]"
//line /usr/local/go/src/encoding/base64/base64.go:187
	default:
//line /usr/local/go/src/encoding/base64/base64.go:187
		_go_fuzz_dep_.CoverTab[10557]++
//line /usr/local/go/src/encoding/base64/base64.go:187
		// _ = "end of CoverTab[10557]"
	}
//line /usr/local/go/src/encoding/base64/base64.go:188
	// _ = "end of CoverTab[10547]"
}

// EncodeToString returns the base64 encoding of src.
func (enc *Encoding) EncodeToString(src []byte) string {
//line /usr/local/go/src/encoding/base64/base64.go:192
	_go_fuzz_dep_.CoverTab[10562]++
							buf := make([]byte, enc.EncodedLen(len(src)))
							enc.Encode(buf, src)
							return string(buf)
//line /usr/local/go/src/encoding/base64/base64.go:195
	// _ = "end of CoverTab[10562]"
}

type encoder struct {
	err	error
	enc	*Encoding
	w	io.Writer
	buf	[3]byte		// buffered data waiting to be encoded
	nbuf	int		// number of bytes in buf
	out	[1024]byte	// output buffer
}

func (e *encoder) Write(p []byte) (n int, err error) {
//line /usr/local/go/src/encoding/base64/base64.go:207
	_go_fuzz_dep_.CoverTab[10563]++
							if e.err != nil {
//line /usr/local/go/src/encoding/base64/base64.go:208
		_go_fuzz_dep_.CoverTab[10567]++
								return 0, e.err
//line /usr/local/go/src/encoding/base64/base64.go:209
		// _ = "end of CoverTab[10567]"
	} else {
//line /usr/local/go/src/encoding/base64/base64.go:210
		_go_fuzz_dep_.CoverTab[10568]++
//line /usr/local/go/src/encoding/base64/base64.go:210
		// _ = "end of CoverTab[10568]"
//line /usr/local/go/src/encoding/base64/base64.go:210
	}
//line /usr/local/go/src/encoding/base64/base64.go:210
	// _ = "end of CoverTab[10563]"
//line /usr/local/go/src/encoding/base64/base64.go:210
	_go_fuzz_dep_.CoverTab[10564]++

//line /usr/local/go/src/encoding/base64/base64.go:213
	if e.nbuf > 0 {
//line /usr/local/go/src/encoding/base64/base64.go:213
		_go_fuzz_dep_.CoverTab[10569]++
								var i int
								for i = 0; i < len(p) && func() bool {
//line /usr/local/go/src/encoding/base64/base64.go:215
			_go_fuzz_dep_.CoverTab[10573]++
//line /usr/local/go/src/encoding/base64/base64.go:215
			return e.nbuf < 3
//line /usr/local/go/src/encoding/base64/base64.go:215
			// _ = "end of CoverTab[10573]"
//line /usr/local/go/src/encoding/base64/base64.go:215
		}(); i++ {
//line /usr/local/go/src/encoding/base64/base64.go:215
			_go_fuzz_dep_.CoverTab[10574]++
									e.buf[e.nbuf] = p[i]
									e.nbuf++
//line /usr/local/go/src/encoding/base64/base64.go:217
			// _ = "end of CoverTab[10574]"
		}
//line /usr/local/go/src/encoding/base64/base64.go:218
		// _ = "end of CoverTab[10569]"
//line /usr/local/go/src/encoding/base64/base64.go:218
		_go_fuzz_dep_.CoverTab[10570]++
								n += i
								p = p[i:]
								if e.nbuf < 3 {
//line /usr/local/go/src/encoding/base64/base64.go:221
			_go_fuzz_dep_.CoverTab[10575]++
									return
//line /usr/local/go/src/encoding/base64/base64.go:222
			// _ = "end of CoverTab[10575]"
		} else {
//line /usr/local/go/src/encoding/base64/base64.go:223
			_go_fuzz_dep_.CoverTab[10576]++
//line /usr/local/go/src/encoding/base64/base64.go:223
			// _ = "end of CoverTab[10576]"
//line /usr/local/go/src/encoding/base64/base64.go:223
		}
//line /usr/local/go/src/encoding/base64/base64.go:223
		// _ = "end of CoverTab[10570]"
//line /usr/local/go/src/encoding/base64/base64.go:223
		_go_fuzz_dep_.CoverTab[10571]++
								e.enc.Encode(e.out[:], e.buf[:])
								if _, e.err = e.w.Write(e.out[:4]); e.err != nil {
//line /usr/local/go/src/encoding/base64/base64.go:225
			_go_fuzz_dep_.CoverTab[10577]++
									return n, e.err
//line /usr/local/go/src/encoding/base64/base64.go:226
			// _ = "end of CoverTab[10577]"
		} else {
//line /usr/local/go/src/encoding/base64/base64.go:227
			_go_fuzz_dep_.CoverTab[10578]++
//line /usr/local/go/src/encoding/base64/base64.go:227
			// _ = "end of CoverTab[10578]"
//line /usr/local/go/src/encoding/base64/base64.go:227
		}
//line /usr/local/go/src/encoding/base64/base64.go:227
		// _ = "end of CoverTab[10571]"
//line /usr/local/go/src/encoding/base64/base64.go:227
		_go_fuzz_dep_.CoverTab[10572]++
								e.nbuf = 0
//line /usr/local/go/src/encoding/base64/base64.go:228
		// _ = "end of CoverTab[10572]"
	} else {
//line /usr/local/go/src/encoding/base64/base64.go:229
		_go_fuzz_dep_.CoverTab[10579]++
//line /usr/local/go/src/encoding/base64/base64.go:229
		// _ = "end of CoverTab[10579]"
//line /usr/local/go/src/encoding/base64/base64.go:229
	}
//line /usr/local/go/src/encoding/base64/base64.go:229
	// _ = "end of CoverTab[10564]"
//line /usr/local/go/src/encoding/base64/base64.go:229
	_go_fuzz_dep_.CoverTab[10565]++

//line /usr/local/go/src/encoding/base64/base64.go:232
	for len(p) >= 3 {
//line /usr/local/go/src/encoding/base64/base64.go:232
		_go_fuzz_dep_.CoverTab[10580]++
								nn := len(e.out) / 4 * 3
								if nn > len(p) {
//line /usr/local/go/src/encoding/base64/base64.go:234
			_go_fuzz_dep_.CoverTab[10583]++
									nn = len(p)
									nn -= nn % 3
//line /usr/local/go/src/encoding/base64/base64.go:236
			// _ = "end of CoverTab[10583]"
		} else {
//line /usr/local/go/src/encoding/base64/base64.go:237
			_go_fuzz_dep_.CoverTab[10584]++
//line /usr/local/go/src/encoding/base64/base64.go:237
			// _ = "end of CoverTab[10584]"
//line /usr/local/go/src/encoding/base64/base64.go:237
		}
//line /usr/local/go/src/encoding/base64/base64.go:237
		// _ = "end of CoverTab[10580]"
//line /usr/local/go/src/encoding/base64/base64.go:237
		_go_fuzz_dep_.CoverTab[10581]++
								e.enc.Encode(e.out[:], p[:nn])
								if _, e.err = e.w.Write(e.out[0 : nn/3*4]); e.err != nil {
//line /usr/local/go/src/encoding/base64/base64.go:239
			_go_fuzz_dep_.CoverTab[10585]++
									return n, e.err
//line /usr/local/go/src/encoding/base64/base64.go:240
			// _ = "end of CoverTab[10585]"
		} else {
//line /usr/local/go/src/encoding/base64/base64.go:241
			_go_fuzz_dep_.CoverTab[10586]++
//line /usr/local/go/src/encoding/base64/base64.go:241
			// _ = "end of CoverTab[10586]"
//line /usr/local/go/src/encoding/base64/base64.go:241
		}
//line /usr/local/go/src/encoding/base64/base64.go:241
		// _ = "end of CoverTab[10581]"
//line /usr/local/go/src/encoding/base64/base64.go:241
		_go_fuzz_dep_.CoverTab[10582]++
								n += nn
								p = p[nn:]
//line /usr/local/go/src/encoding/base64/base64.go:243
		// _ = "end of CoverTab[10582]"
	}
//line /usr/local/go/src/encoding/base64/base64.go:244
	// _ = "end of CoverTab[10565]"
//line /usr/local/go/src/encoding/base64/base64.go:244
	_go_fuzz_dep_.CoverTab[10566]++

//line /usr/local/go/src/encoding/base64/base64.go:247
	copy(e.buf[:], p)
							e.nbuf = len(p)
							n += len(p)
							return
//line /usr/local/go/src/encoding/base64/base64.go:250
	// _ = "end of CoverTab[10566]"
}

// Close flushes any pending output from the encoder.
//line /usr/local/go/src/encoding/base64/base64.go:253
// It is an error to call Write after calling Close.
//line /usr/local/go/src/encoding/base64/base64.go:255
func (e *encoder) Close() error {
//line /usr/local/go/src/encoding/base64/base64.go:255
	_go_fuzz_dep_.CoverTab[10587]++

							if e.err == nil && func() bool {
//line /usr/local/go/src/encoding/base64/base64.go:257
		_go_fuzz_dep_.CoverTab[10589]++
//line /usr/local/go/src/encoding/base64/base64.go:257
		return e.nbuf > 0
//line /usr/local/go/src/encoding/base64/base64.go:257
		// _ = "end of CoverTab[10589]"
//line /usr/local/go/src/encoding/base64/base64.go:257
	}() {
//line /usr/local/go/src/encoding/base64/base64.go:257
		_go_fuzz_dep_.CoverTab[10590]++
								e.enc.Encode(e.out[:], e.buf[:e.nbuf])
								_, e.err = e.w.Write(e.out[:e.enc.EncodedLen(e.nbuf)])
								e.nbuf = 0
//line /usr/local/go/src/encoding/base64/base64.go:260
		// _ = "end of CoverTab[10590]"
	} else {
//line /usr/local/go/src/encoding/base64/base64.go:261
		_go_fuzz_dep_.CoverTab[10591]++
//line /usr/local/go/src/encoding/base64/base64.go:261
		// _ = "end of CoverTab[10591]"
//line /usr/local/go/src/encoding/base64/base64.go:261
	}
//line /usr/local/go/src/encoding/base64/base64.go:261
	// _ = "end of CoverTab[10587]"
//line /usr/local/go/src/encoding/base64/base64.go:261
	_go_fuzz_dep_.CoverTab[10588]++
							return e.err
//line /usr/local/go/src/encoding/base64/base64.go:262
	// _ = "end of CoverTab[10588]"
}

// NewEncoder returns a new base64 stream encoder. Data written to
//line /usr/local/go/src/encoding/base64/base64.go:265
// the returned writer will be encoded using enc and then written to w.
//line /usr/local/go/src/encoding/base64/base64.go:265
// Base64 encodings operate in 4-byte blocks; when finished
//line /usr/local/go/src/encoding/base64/base64.go:265
// writing, the caller must Close the returned encoder to flush any
//line /usr/local/go/src/encoding/base64/base64.go:265
// partially written blocks.
//line /usr/local/go/src/encoding/base64/base64.go:270
func NewEncoder(enc *Encoding, w io.Writer) io.WriteCloser {
//line /usr/local/go/src/encoding/base64/base64.go:270
	_go_fuzz_dep_.CoverTab[10592]++
							return &encoder{enc: enc, w: w}
//line /usr/local/go/src/encoding/base64/base64.go:271
	// _ = "end of CoverTab[10592]"
}

// EncodedLen returns the length in bytes of the base64 encoding
//line /usr/local/go/src/encoding/base64/base64.go:274
// of an input buffer of length n.
//line /usr/local/go/src/encoding/base64/base64.go:276
func (enc *Encoding) EncodedLen(n int) int {
//line /usr/local/go/src/encoding/base64/base64.go:276
	_go_fuzz_dep_.CoverTab[10593]++
							if enc.padChar == NoPadding {
//line /usr/local/go/src/encoding/base64/base64.go:277
		_go_fuzz_dep_.CoverTab[10595]++
								return (n*8 + 5) / 6
//line /usr/local/go/src/encoding/base64/base64.go:278
		// _ = "end of CoverTab[10595]"
	} else {
//line /usr/local/go/src/encoding/base64/base64.go:279
		_go_fuzz_dep_.CoverTab[10596]++
//line /usr/local/go/src/encoding/base64/base64.go:279
		// _ = "end of CoverTab[10596]"
//line /usr/local/go/src/encoding/base64/base64.go:279
	}
//line /usr/local/go/src/encoding/base64/base64.go:279
	// _ = "end of CoverTab[10593]"
//line /usr/local/go/src/encoding/base64/base64.go:279
	_go_fuzz_dep_.CoverTab[10594]++
							return (n + 2) / 3 * 4
//line /usr/local/go/src/encoding/base64/base64.go:280
	// _ = "end of CoverTab[10594]"
}

//line /usr/local/go/src/encoding/base64/base64.go:287
type CorruptInputError int64

func (e CorruptInputError) Error() string {
//line /usr/local/go/src/encoding/base64/base64.go:289
	_go_fuzz_dep_.CoverTab[10597]++
							return "illegal base64 data at input byte " + strconv.FormatInt(int64(e), 10)
//line /usr/local/go/src/encoding/base64/base64.go:290
	// _ = "end of CoverTab[10597]"
}

// decodeQuantum decodes up to 4 base64 bytes. The received parameters are
//line /usr/local/go/src/encoding/base64/base64.go:293
// the destination buffer dst, the source buffer src and an index in the
//line /usr/local/go/src/encoding/base64/base64.go:293
// source buffer si.
//line /usr/local/go/src/encoding/base64/base64.go:293
// It returns the number of bytes read from src, the number of bytes written
//line /usr/local/go/src/encoding/base64/base64.go:293
// to dst, and an error, if any.
//line /usr/local/go/src/encoding/base64/base64.go:298
func (enc *Encoding) decodeQuantum(dst, src []byte, si int) (nsi, n int, err error) {
//line /usr/local/go/src/encoding/base64/base64.go:298
	_go_fuzz_dep_.CoverTab[10598]++
							// Decode quantum using the base64 alphabet
							var dbuf [4]byte
							dlen := 4

//line /usr/local/go/src/encoding/base64/base64.go:304
	_ = enc.decodeMap

	for j := 0; j < len(dbuf); j++ {
//line /usr/local/go/src/encoding/base64/base64.go:306
		_go_fuzz_dep_.CoverTab[10601]++
								if len(src) == si {
//line /usr/local/go/src/encoding/base64/base64.go:307
			_go_fuzz_dep_.CoverTab[10609]++
									switch {
			case j == 0:
//line /usr/local/go/src/encoding/base64/base64.go:309
				_go_fuzz_dep_.CoverTab[10611]++
										return si, 0, nil
//line /usr/local/go/src/encoding/base64/base64.go:310
				// _ = "end of CoverTab[10611]"
			case j == 1, enc.padChar != NoPadding:
//line /usr/local/go/src/encoding/base64/base64.go:311
				_go_fuzz_dep_.CoverTab[10612]++
										return si, 0, CorruptInputError(si - j)
//line /usr/local/go/src/encoding/base64/base64.go:312
				// _ = "end of CoverTab[10612]"
//line /usr/local/go/src/encoding/base64/base64.go:312
			default:
//line /usr/local/go/src/encoding/base64/base64.go:312
				_go_fuzz_dep_.CoverTab[10613]++
//line /usr/local/go/src/encoding/base64/base64.go:312
				// _ = "end of CoverTab[10613]"
			}
//line /usr/local/go/src/encoding/base64/base64.go:313
			// _ = "end of CoverTab[10609]"
//line /usr/local/go/src/encoding/base64/base64.go:313
			_go_fuzz_dep_.CoverTab[10610]++
									dlen = j
									break
//line /usr/local/go/src/encoding/base64/base64.go:315
			// _ = "end of CoverTab[10610]"
		} else {
//line /usr/local/go/src/encoding/base64/base64.go:316
			_go_fuzz_dep_.CoverTab[10614]++
//line /usr/local/go/src/encoding/base64/base64.go:316
			// _ = "end of CoverTab[10614]"
//line /usr/local/go/src/encoding/base64/base64.go:316
		}
//line /usr/local/go/src/encoding/base64/base64.go:316
		// _ = "end of CoverTab[10601]"
//line /usr/local/go/src/encoding/base64/base64.go:316
		_go_fuzz_dep_.CoverTab[10602]++
								in := src[si]
								si++

								out := enc.decodeMap[in]
								if out != 0xff {
//line /usr/local/go/src/encoding/base64/base64.go:321
			_go_fuzz_dep_.CoverTab[10615]++
									dbuf[j] = out
									continue
//line /usr/local/go/src/encoding/base64/base64.go:323
			// _ = "end of CoverTab[10615]"
		} else {
//line /usr/local/go/src/encoding/base64/base64.go:324
			_go_fuzz_dep_.CoverTab[10616]++
//line /usr/local/go/src/encoding/base64/base64.go:324
			// _ = "end of CoverTab[10616]"
//line /usr/local/go/src/encoding/base64/base64.go:324
		}
//line /usr/local/go/src/encoding/base64/base64.go:324
		// _ = "end of CoverTab[10602]"
//line /usr/local/go/src/encoding/base64/base64.go:324
		_go_fuzz_dep_.CoverTab[10603]++

								if in == '\n' || func() bool {
//line /usr/local/go/src/encoding/base64/base64.go:326
			_go_fuzz_dep_.CoverTab[10617]++
//line /usr/local/go/src/encoding/base64/base64.go:326
			return in == '\r'
//line /usr/local/go/src/encoding/base64/base64.go:326
			// _ = "end of CoverTab[10617]"
//line /usr/local/go/src/encoding/base64/base64.go:326
		}() {
//line /usr/local/go/src/encoding/base64/base64.go:326
			_go_fuzz_dep_.CoverTab[10618]++
									j--
									continue
//line /usr/local/go/src/encoding/base64/base64.go:328
			// _ = "end of CoverTab[10618]"
		} else {
//line /usr/local/go/src/encoding/base64/base64.go:329
			_go_fuzz_dep_.CoverTab[10619]++
//line /usr/local/go/src/encoding/base64/base64.go:329
			// _ = "end of CoverTab[10619]"
//line /usr/local/go/src/encoding/base64/base64.go:329
		}
//line /usr/local/go/src/encoding/base64/base64.go:329
		// _ = "end of CoverTab[10603]"
//line /usr/local/go/src/encoding/base64/base64.go:329
		_go_fuzz_dep_.CoverTab[10604]++

								if rune(in) != enc.padChar {
//line /usr/local/go/src/encoding/base64/base64.go:331
			_go_fuzz_dep_.CoverTab[10620]++
									return si, 0, CorruptInputError(si - 1)
//line /usr/local/go/src/encoding/base64/base64.go:332
			// _ = "end of CoverTab[10620]"
		} else {
//line /usr/local/go/src/encoding/base64/base64.go:333
			_go_fuzz_dep_.CoverTab[10621]++
//line /usr/local/go/src/encoding/base64/base64.go:333
			// _ = "end of CoverTab[10621]"
//line /usr/local/go/src/encoding/base64/base64.go:333
		}
//line /usr/local/go/src/encoding/base64/base64.go:333
		// _ = "end of CoverTab[10604]"
//line /usr/local/go/src/encoding/base64/base64.go:333
		_go_fuzz_dep_.CoverTab[10605]++

//line /usr/local/go/src/encoding/base64/base64.go:336
		switch j {
		case 0, 1:
//line /usr/local/go/src/encoding/base64/base64.go:337
			_go_fuzz_dep_.CoverTab[10622]++

									return si, 0, CorruptInputError(si - 1)
//line /usr/local/go/src/encoding/base64/base64.go:339
			// _ = "end of CoverTab[10622]"
		case 2:
//line /usr/local/go/src/encoding/base64/base64.go:340
			_go_fuzz_dep_.CoverTab[10623]++

//line /usr/local/go/src/encoding/base64/base64.go:343
			for si < len(src) && func() bool {
//line /usr/local/go/src/encoding/base64/base64.go:343
				_go_fuzz_dep_.CoverTab[10628]++
//line /usr/local/go/src/encoding/base64/base64.go:343
				return (src[si] == '\n' || func() bool {
//line /usr/local/go/src/encoding/base64/base64.go:343
					_go_fuzz_dep_.CoverTab[10629]++
//line /usr/local/go/src/encoding/base64/base64.go:343
					return src[si] == '\r'
//line /usr/local/go/src/encoding/base64/base64.go:343
					// _ = "end of CoverTab[10629]"
//line /usr/local/go/src/encoding/base64/base64.go:343
				}())
//line /usr/local/go/src/encoding/base64/base64.go:343
				// _ = "end of CoverTab[10628]"
//line /usr/local/go/src/encoding/base64/base64.go:343
			}() {
//line /usr/local/go/src/encoding/base64/base64.go:343
				_go_fuzz_dep_.CoverTab[10630]++
										si++
//line /usr/local/go/src/encoding/base64/base64.go:344
				// _ = "end of CoverTab[10630]"
			}
//line /usr/local/go/src/encoding/base64/base64.go:345
			// _ = "end of CoverTab[10623]"
//line /usr/local/go/src/encoding/base64/base64.go:345
			_go_fuzz_dep_.CoverTab[10624]++
									if si == len(src) {
//line /usr/local/go/src/encoding/base64/base64.go:346
				_go_fuzz_dep_.CoverTab[10631]++

										return si, 0, CorruptInputError(len(src))
//line /usr/local/go/src/encoding/base64/base64.go:348
				// _ = "end of CoverTab[10631]"
			} else {
//line /usr/local/go/src/encoding/base64/base64.go:349
				_go_fuzz_dep_.CoverTab[10632]++
//line /usr/local/go/src/encoding/base64/base64.go:349
				// _ = "end of CoverTab[10632]"
//line /usr/local/go/src/encoding/base64/base64.go:349
			}
//line /usr/local/go/src/encoding/base64/base64.go:349
			// _ = "end of CoverTab[10624]"
//line /usr/local/go/src/encoding/base64/base64.go:349
			_go_fuzz_dep_.CoverTab[10625]++
									if rune(src[si]) != enc.padChar {
//line /usr/local/go/src/encoding/base64/base64.go:350
				_go_fuzz_dep_.CoverTab[10633]++

										return si, 0, CorruptInputError(si - 1)
//line /usr/local/go/src/encoding/base64/base64.go:352
				// _ = "end of CoverTab[10633]"
			} else {
//line /usr/local/go/src/encoding/base64/base64.go:353
				_go_fuzz_dep_.CoverTab[10634]++
//line /usr/local/go/src/encoding/base64/base64.go:353
				// _ = "end of CoverTab[10634]"
//line /usr/local/go/src/encoding/base64/base64.go:353
			}
//line /usr/local/go/src/encoding/base64/base64.go:353
			// _ = "end of CoverTab[10625]"
//line /usr/local/go/src/encoding/base64/base64.go:353
			_go_fuzz_dep_.CoverTab[10626]++

									si++
//line /usr/local/go/src/encoding/base64/base64.go:355
			// _ = "end of CoverTab[10626]"
//line /usr/local/go/src/encoding/base64/base64.go:355
		default:
//line /usr/local/go/src/encoding/base64/base64.go:355
			_go_fuzz_dep_.CoverTab[10627]++
//line /usr/local/go/src/encoding/base64/base64.go:355
			// _ = "end of CoverTab[10627]"
		}
//line /usr/local/go/src/encoding/base64/base64.go:356
		// _ = "end of CoverTab[10605]"
//line /usr/local/go/src/encoding/base64/base64.go:356
		_go_fuzz_dep_.CoverTab[10606]++

//line /usr/local/go/src/encoding/base64/base64.go:359
		for si < len(src) && func() bool {
//line /usr/local/go/src/encoding/base64/base64.go:359
			_go_fuzz_dep_.CoverTab[10635]++
//line /usr/local/go/src/encoding/base64/base64.go:359
			return (src[si] == '\n' || func() bool {
//line /usr/local/go/src/encoding/base64/base64.go:359
				_go_fuzz_dep_.CoverTab[10636]++
//line /usr/local/go/src/encoding/base64/base64.go:359
				return src[si] == '\r'
//line /usr/local/go/src/encoding/base64/base64.go:359
				// _ = "end of CoverTab[10636]"
//line /usr/local/go/src/encoding/base64/base64.go:359
			}())
//line /usr/local/go/src/encoding/base64/base64.go:359
			// _ = "end of CoverTab[10635]"
//line /usr/local/go/src/encoding/base64/base64.go:359
		}() {
//line /usr/local/go/src/encoding/base64/base64.go:359
			_go_fuzz_dep_.CoverTab[10637]++
									si++
//line /usr/local/go/src/encoding/base64/base64.go:360
			// _ = "end of CoverTab[10637]"
		}
//line /usr/local/go/src/encoding/base64/base64.go:361
		// _ = "end of CoverTab[10606]"
//line /usr/local/go/src/encoding/base64/base64.go:361
		_go_fuzz_dep_.CoverTab[10607]++
								if si < len(src) {
//line /usr/local/go/src/encoding/base64/base64.go:362
			_go_fuzz_dep_.CoverTab[10638]++

									err = CorruptInputError(si)
//line /usr/local/go/src/encoding/base64/base64.go:364
			// _ = "end of CoverTab[10638]"
		} else {
//line /usr/local/go/src/encoding/base64/base64.go:365
			_go_fuzz_dep_.CoverTab[10639]++
//line /usr/local/go/src/encoding/base64/base64.go:365
			// _ = "end of CoverTab[10639]"
//line /usr/local/go/src/encoding/base64/base64.go:365
		}
//line /usr/local/go/src/encoding/base64/base64.go:365
		// _ = "end of CoverTab[10607]"
//line /usr/local/go/src/encoding/base64/base64.go:365
		_go_fuzz_dep_.CoverTab[10608]++
								dlen = j
								break
//line /usr/local/go/src/encoding/base64/base64.go:367
		// _ = "end of CoverTab[10608]"
	}
//line /usr/local/go/src/encoding/base64/base64.go:368
	// _ = "end of CoverTab[10598]"
//line /usr/local/go/src/encoding/base64/base64.go:368
	_go_fuzz_dep_.CoverTab[10599]++

//line /usr/local/go/src/encoding/base64/base64.go:371
	val := uint(dbuf[0])<<18 | uint(dbuf[1])<<12 | uint(dbuf[2])<<6 | uint(dbuf[3])
	dbuf[2], dbuf[1], dbuf[0] = byte(val>>0), byte(val>>8), byte(val>>16)
	switch dlen {
	case 4:
//line /usr/local/go/src/encoding/base64/base64.go:374
		_go_fuzz_dep_.CoverTab[10640]++
								dst[2] = dbuf[2]
								dbuf[2] = 0
								fallthrough
//line /usr/local/go/src/encoding/base64/base64.go:377
		// _ = "end of CoverTab[10640]"
	case 3:
//line /usr/local/go/src/encoding/base64/base64.go:378
		_go_fuzz_dep_.CoverTab[10641]++
								dst[1] = dbuf[1]
								if enc.strict && func() bool {
//line /usr/local/go/src/encoding/base64/base64.go:380
			_go_fuzz_dep_.CoverTab[10645]++
//line /usr/local/go/src/encoding/base64/base64.go:380
			return dbuf[2] != 0
//line /usr/local/go/src/encoding/base64/base64.go:380
			// _ = "end of CoverTab[10645]"
//line /usr/local/go/src/encoding/base64/base64.go:380
		}() {
//line /usr/local/go/src/encoding/base64/base64.go:380
			_go_fuzz_dep_.CoverTab[10646]++
									return si, 0, CorruptInputError(si - 1)
//line /usr/local/go/src/encoding/base64/base64.go:381
			// _ = "end of CoverTab[10646]"
		} else {
//line /usr/local/go/src/encoding/base64/base64.go:382
			_go_fuzz_dep_.CoverTab[10647]++
//line /usr/local/go/src/encoding/base64/base64.go:382
			// _ = "end of CoverTab[10647]"
//line /usr/local/go/src/encoding/base64/base64.go:382
		}
//line /usr/local/go/src/encoding/base64/base64.go:382
		// _ = "end of CoverTab[10641]"
//line /usr/local/go/src/encoding/base64/base64.go:382
		_go_fuzz_dep_.CoverTab[10642]++
								dbuf[1] = 0
								fallthrough
//line /usr/local/go/src/encoding/base64/base64.go:384
		// _ = "end of CoverTab[10642]"
	case 2:
//line /usr/local/go/src/encoding/base64/base64.go:385
		_go_fuzz_dep_.CoverTab[10643]++
								dst[0] = dbuf[0]
								if enc.strict && func() bool {
//line /usr/local/go/src/encoding/base64/base64.go:387
			_go_fuzz_dep_.CoverTab[10648]++
//line /usr/local/go/src/encoding/base64/base64.go:387
			return (dbuf[1] != 0 || func() bool {
//line /usr/local/go/src/encoding/base64/base64.go:387
				_go_fuzz_dep_.CoverTab[10649]++
//line /usr/local/go/src/encoding/base64/base64.go:387
				return dbuf[2] != 0
//line /usr/local/go/src/encoding/base64/base64.go:387
				// _ = "end of CoverTab[10649]"
//line /usr/local/go/src/encoding/base64/base64.go:387
			}())
//line /usr/local/go/src/encoding/base64/base64.go:387
			// _ = "end of CoverTab[10648]"
//line /usr/local/go/src/encoding/base64/base64.go:387
		}() {
//line /usr/local/go/src/encoding/base64/base64.go:387
			_go_fuzz_dep_.CoverTab[10650]++
									return si, 0, CorruptInputError(si - 2)
//line /usr/local/go/src/encoding/base64/base64.go:388
			// _ = "end of CoverTab[10650]"
		} else {
//line /usr/local/go/src/encoding/base64/base64.go:389
			_go_fuzz_dep_.CoverTab[10651]++
//line /usr/local/go/src/encoding/base64/base64.go:389
			// _ = "end of CoverTab[10651]"
//line /usr/local/go/src/encoding/base64/base64.go:389
		}
//line /usr/local/go/src/encoding/base64/base64.go:389
		// _ = "end of CoverTab[10643]"
//line /usr/local/go/src/encoding/base64/base64.go:389
	default:
//line /usr/local/go/src/encoding/base64/base64.go:389
		_go_fuzz_dep_.CoverTab[10644]++
//line /usr/local/go/src/encoding/base64/base64.go:389
		// _ = "end of CoverTab[10644]"
	}
//line /usr/local/go/src/encoding/base64/base64.go:390
	// _ = "end of CoverTab[10599]"
//line /usr/local/go/src/encoding/base64/base64.go:390
	_go_fuzz_dep_.CoverTab[10600]++

							return si, dlen - 1, err
//line /usr/local/go/src/encoding/base64/base64.go:392
	// _ = "end of CoverTab[10600]"
}

// DecodeString returns the bytes represented by the base64 string s.
func (enc *Encoding) DecodeString(s string) ([]byte, error) {
//line /usr/local/go/src/encoding/base64/base64.go:396
	_go_fuzz_dep_.CoverTab[10652]++
							dbuf := make([]byte, enc.DecodedLen(len(s)))
							n, err := enc.Decode(dbuf, []byte(s))
							return dbuf[:n], err
//line /usr/local/go/src/encoding/base64/base64.go:399
	// _ = "end of CoverTab[10652]"
}

type decoder struct {
	err	error
	readErr	error	// error from r.Read
	enc	*Encoding
	r	io.Reader
	buf	[1024]byte	// leftover input
	nbuf	int
	out	[]byte	// leftover decoded output
	outbuf	[1024 / 4 * 3]byte
}

func (d *decoder) Read(p []byte) (n int, err error) {
//line /usr/local/go/src/encoding/base64/base64.go:413
	_go_fuzz_dep_.CoverTab[10653]++

							if len(d.out) > 0 {
//line /usr/local/go/src/encoding/base64/base64.go:415
		_go_fuzz_dep_.CoverTab[10659]++
								n = copy(p, d.out)
								d.out = d.out[n:]
								return n, nil
//line /usr/local/go/src/encoding/base64/base64.go:418
		// _ = "end of CoverTab[10659]"
	} else {
//line /usr/local/go/src/encoding/base64/base64.go:419
		_go_fuzz_dep_.CoverTab[10660]++
//line /usr/local/go/src/encoding/base64/base64.go:419
		// _ = "end of CoverTab[10660]"
//line /usr/local/go/src/encoding/base64/base64.go:419
	}
//line /usr/local/go/src/encoding/base64/base64.go:419
	// _ = "end of CoverTab[10653]"
//line /usr/local/go/src/encoding/base64/base64.go:419
	_go_fuzz_dep_.CoverTab[10654]++

							if d.err != nil {
//line /usr/local/go/src/encoding/base64/base64.go:421
		_go_fuzz_dep_.CoverTab[10661]++
								return 0, d.err
//line /usr/local/go/src/encoding/base64/base64.go:422
		// _ = "end of CoverTab[10661]"
	} else {
//line /usr/local/go/src/encoding/base64/base64.go:423
		_go_fuzz_dep_.CoverTab[10662]++
//line /usr/local/go/src/encoding/base64/base64.go:423
		// _ = "end of CoverTab[10662]"
//line /usr/local/go/src/encoding/base64/base64.go:423
	}
//line /usr/local/go/src/encoding/base64/base64.go:423
	// _ = "end of CoverTab[10654]"
//line /usr/local/go/src/encoding/base64/base64.go:423
	_go_fuzz_dep_.CoverTab[10655]++

//line /usr/local/go/src/encoding/base64/base64.go:428
	for d.nbuf < 4 && func() bool {
//line /usr/local/go/src/encoding/base64/base64.go:428
		_go_fuzz_dep_.CoverTab[10663]++
//line /usr/local/go/src/encoding/base64/base64.go:428
		return d.readErr == nil
//line /usr/local/go/src/encoding/base64/base64.go:428
		// _ = "end of CoverTab[10663]"
//line /usr/local/go/src/encoding/base64/base64.go:428
	}() {
//line /usr/local/go/src/encoding/base64/base64.go:428
		_go_fuzz_dep_.CoverTab[10664]++
								nn := len(p) / 3 * 4
								if nn < 4 {
//line /usr/local/go/src/encoding/base64/base64.go:430
			_go_fuzz_dep_.CoverTab[10667]++
									nn = 4
//line /usr/local/go/src/encoding/base64/base64.go:431
			// _ = "end of CoverTab[10667]"
		} else {
//line /usr/local/go/src/encoding/base64/base64.go:432
			_go_fuzz_dep_.CoverTab[10668]++
//line /usr/local/go/src/encoding/base64/base64.go:432
			// _ = "end of CoverTab[10668]"
//line /usr/local/go/src/encoding/base64/base64.go:432
		}
//line /usr/local/go/src/encoding/base64/base64.go:432
		// _ = "end of CoverTab[10664]"
//line /usr/local/go/src/encoding/base64/base64.go:432
		_go_fuzz_dep_.CoverTab[10665]++
								if nn > len(d.buf) {
//line /usr/local/go/src/encoding/base64/base64.go:433
			_go_fuzz_dep_.CoverTab[10669]++
									nn = len(d.buf)
//line /usr/local/go/src/encoding/base64/base64.go:434
			// _ = "end of CoverTab[10669]"
		} else {
//line /usr/local/go/src/encoding/base64/base64.go:435
			_go_fuzz_dep_.CoverTab[10670]++
//line /usr/local/go/src/encoding/base64/base64.go:435
			// _ = "end of CoverTab[10670]"
//line /usr/local/go/src/encoding/base64/base64.go:435
		}
//line /usr/local/go/src/encoding/base64/base64.go:435
		// _ = "end of CoverTab[10665]"
//line /usr/local/go/src/encoding/base64/base64.go:435
		_go_fuzz_dep_.CoverTab[10666]++
								nn, d.readErr = d.r.Read(d.buf[d.nbuf:nn])
								d.nbuf += nn
//line /usr/local/go/src/encoding/base64/base64.go:437
		// _ = "end of CoverTab[10666]"
	}
//line /usr/local/go/src/encoding/base64/base64.go:438
	// _ = "end of CoverTab[10655]"
//line /usr/local/go/src/encoding/base64/base64.go:438
	_go_fuzz_dep_.CoverTab[10656]++

							if d.nbuf < 4 {
//line /usr/local/go/src/encoding/base64/base64.go:440
		_go_fuzz_dep_.CoverTab[10671]++
								if d.enc.padChar == NoPadding && func() bool {
//line /usr/local/go/src/encoding/base64/base64.go:441
			_go_fuzz_dep_.CoverTab[10674]++
//line /usr/local/go/src/encoding/base64/base64.go:441
			return d.nbuf > 0
//line /usr/local/go/src/encoding/base64/base64.go:441
			// _ = "end of CoverTab[10674]"
//line /usr/local/go/src/encoding/base64/base64.go:441
		}() {
//line /usr/local/go/src/encoding/base64/base64.go:441
			_go_fuzz_dep_.CoverTab[10675]++
			// Decode final fragment, without padding.
			var nw int
			nw, d.err = d.enc.Decode(d.outbuf[:], d.buf[:d.nbuf])
			d.nbuf = 0
			d.out = d.outbuf[:nw]
			n = copy(p, d.out)
			d.out = d.out[n:]
			if n > 0 || func() bool {
//line /usr/local/go/src/encoding/base64/base64.go:449
				_go_fuzz_dep_.CoverTab[10677]++
//line /usr/local/go/src/encoding/base64/base64.go:449
				return len(p) == 0 && func() bool {
//line /usr/local/go/src/encoding/base64/base64.go:449
					_go_fuzz_dep_.CoverTab[10678]++
//line /usr/local/go/src/encoding/base64/base64.go:449
					return len(d.out) > 0
//line /usr/local/go/src/encoding/base64/base64.go:449
					// _ = "end of CoverTab[10678]"
//line /usr/local/go/src/encoding/base64/base64.go:449
				}()
//line /usr/local/go/src/encoding/base64/base64.go:449
				// _ = "end of CoverTab[10677]"
//line /usr/local/go/src/encoding/base64/base64.go:449
			}() {
//line /usr/local/go/src/encoding/base64/base64.go:449
				_go_fuzz_dep_.CoverTab[10679]++
										return n, nil
//line /usr/local/go/src/encoding/base64/base64.go:450
				// _ = "end of CoverTab[10679]"
			} else {
//line /usr/local/go/src/encoding/base64/base64.go:451
				_go_fuzz_dep_.CoverTab[10680]++
//line /usr/local/go/src/encoding/base64/base64.go:451
				// _ = "end of CoverTab[10680]"
//line /usr/local/go/src/encoding/base64/base64.go:451
			}
//line /usr/local/go/src/encoding/base64/base64.go:451
			// _ = "end of CoverTab[10675]"
//line /usr/local/go/src/encoding/base64/base64.go:451
			_go_fuzz_dep_.CoverTab[10676]++
									if d.err != nil {
//line /usr/local/go/src/encoding/base64/base64.go:452
				_go_fuzz_dep_.CoverTab[10681]++
										return 0, d.err
//line /usr/local/go/src/encoding/base64/base64.go:453
				// _ = "end of CoverTab[10681]"
			} else {
//line /usr/local/go/src/encoding/base64/base64.go:454
				_go_fuzz_dep_.CoverTab[10682]++
//line /usr/local/go/src/encoding/base64/base64.go:454
				// _ = "end of CoverTab[10682]"
//line /usr/local/go/src/encoding/base64/base64.go:454
			}
//line /usr/local/go/src/encoding/base64/base64.go:454
			// _ = "end of CoverTab[10676]"
		} else {
//line /usr/local/go/src/encoding/base64/base64.go:455
			_go_fuzz_dep_.CoverTab[10683]++
//line /usr/local/go/src/encoding/base64/base64.go:455
			// _ = "end of CoverTab[10683]"
//line /usr/local/go/src/encoding/base64/base64.go:455
		}
//line /usr/local/go/src/encoding/base64/base64.go:455
		// _ = "end of CoverTab[10671]"
//line /usr/local/go/src/encoding/base64/base64.go:455
		_go_fuzz_dep_.CoverTab[10672]++
								d.err = d.readErr
								if d.err == io.EOF && func() bool {
//line /usr/local/go/src/encoding/base64/base64.go:457
			_go_fuzz_dep_.CoverTab[10684]++
//line /usr/local/go/src/encoding/base64/base64.go:457
			return d.nbuf > 0
//line /usr/local/go/src/encoding/base64/base64.go:457
			// _ = "end of CoverTab[10684]"
//line /usr/local/go/src/encoding/base64/base64.go:457
		}() {
//line /usr/local/go/src/encoding/base64/base64.go:457
			_go_fuzz_dep_.CoverTab[10685]++
									d.err = io.ErrUnexpectedEOF
//line /usr/local/go/src/encoding/base64/base64.go:458
			// _ = "end of CoverTab[10685]"
		} else {
//line /usr/local/go/src/encoding/base64/base64.go:459
			_go_fuzz_dep_.CoverTab[10686]++
//line /usr/local/go/src/encoding/base64/base64.go:459
			// _ = "end of CoverTab[10686]"
//line /usr/local/go/src/encoding/base64/base64.go:459
		}
//line /usr/local/go/src/encoding/base64/base64.go:459
		// _ = "end of CoverTab[10672]"
//line /usr/local/go/src/encoding/base64/base64.go:459
		_go_fuzz_dep_.CoverTab[10673]++
								return 0, d.err
//line /usr/local/go/src/encoding/base64/base64.go:460
		// _ = "end of CoverTab[10673]"
	} else {
//line /usr/local/go/src/encoding/base64/base64.go:461
		_go_fuzz_dep_.CoverTab[10687]++
//line /usr/local/go/src/encoding/base64/base64.go:461
		// _ = "end of CoverTab[10687]"
//line /usr/local/go/src/encoding/base64/base64.go:461
	}
//line /usr/local/go/src/encoding/base64/base64.go:461
	// _ = "end of CoverTab[10656]"
//line /usr/local/go/src/encoding/base64/base64.go:461
	_go_fuzz_dep_.CoverTab[10657]++

//line /usr/local/go/src/encoding/base64/base64.go:464
	nr := d.nbuf / 4 * 4
	nw := d.nbuf / 4 * 3
	if nw > len(p) {
//line /usr/local/go/src/encoding/base64/base64.go:466
		_go_fuzz_dep_.CoverTab[10688]++
								nw, d.err = d.enc.Decode(d.outbuf[:], d.buf[:nr])
								d.out = d.outbuf[:nw]
								n = copy(p, d.out)
								d.out = d.out[n:]
//line /usr/local/go/src/encoding/base64/base64.go:470
		// _ = "end of CoverTab[10688]"
	} else {
//line /usr/local/go/src/encoding/base64/base64.go:471
		_go_fuzz_dep_.CoverTab[10689]++
								n, d.err = d.enc.Decode(p, d.buf[:nr])
//line /usr/local/go/src/encoding/base64/base64.go:472
		// _ = "end of CoverTab[10689]"
	}
//line /usr/local/go/src/encoding/base64/base64.go:473
	// _ = "end of CoverTab[10657]"
//line /usr/local/go/src/encoding/base64/base64.go:473
	_go_fuzz_dep_.CoverTab[10658]++
							d.nbuf -= nr
							copy(d.buf[:d.nbuf], d.buf[nr:])
							return n, d.err
//line /usr/local/go/src/encoding/base64/base64.go:476
	// _ = "end of CoverTab[10658]"
}

// Decode decodes src using the encoding enc. It writes at most
//line /usr/local/go/src/encoding/base64/base64.go:479
// DecodedLen(len(src)) bytes to dst and returns the number of bytes
//line /usr/local/go/src/encoding/base64/base64.go:479
// written. If src contains invalid base64 data, it will return the
//line /usr/local/go/src/encoding/base64/base64.go:479
// number of bytes successfully written and CorruptInputError.
//line /usr/local/go/src/encoding/base64/base64.go:479
// New line characters (\r and \n) are ignored.
//line /usr/local/go/src/encoding/base64/base64.go:484
func (enc *Encoding) Decode(dst, src []byte) (n int, err error) {
//line /usr/local/go/src/encoding/base64/base64.go:484
	_go_fuzz_dep_.CoverTab[10690]++
							if len(src) == 0 {
//line /usr/local/go/src/encoding/base64/base64.go:485
		_go_fuzz_dep_.CoverTab[10695]++
								return 0, nil
//line /usr/local/go/src/encoding/base64/base64.go:486
		// _ = "end of CoverTab[10695]"
	} else {
//line /usr/local/go/src/encoding/base64/base64.go:487
		_go_fuzz_dep_.CoverTab[10696]++
//line /usr/local/go/src/encoding/base64/base64.go:487
		// _ = "end of CoverTab[10696]"
//line /usr/local/go/src/encoding/base64/base64.go:487
	}
//line /usr/local/go/src/encoding/base64/base64.go:487
	// _ = "end of CoverTab[10690]"
//line /usr/local/go/src/encoding/base64/base64.go:487
	_go_fuzz_dep_.CoverTab[10691]++

//line /usr/local/go/src/encoding/base64/base64.go:492
	_ = enc.decodeMap

	si := 0
	for strconv.IntSize >= 64 && func() bool {
//line /usr/local/go/src/encoding/base64/base64.go:495
		_go_fuzz_dep_.CoverTab[10697]++
//line /usr/local/go/src/encoding/base64/base64.go:495
		return len(src)-si >= 8
//line /usr/local/go/src/encoding/base64/base64.go:495
		// _ = "end of CoverTab[10697]"
//line /usr/local/go/src/encoding/base64/base64.go:495
	}() && func() bool {
//line /usr/local/go/src/encoding/base64/base64.go:495
		_go_fuzz_dep_.CoverTab[10698]++
//line /usr/local/go/src/encoding/base64/base64.go:495
		return len(dst)-n >= 8
//line /usr/local/go/src/encoding/base64/base64.go:495
		// _ = "end of CoverTab[10698]"
//line /usr/local/go/src/encoding/base64/base64.go:495
	}() {
//line /usr/local/go/src/encoding/base64/base64.go:495
		_go_fuzz_dep_.CoverTab[10699]++
								src2 := src[si : si+8]
								if dn, ok := assemble64(
			enc.decodeMap[src2[0]],
			enc.decodeMap[src2[1]],
			enc.decodeMap[src2[2]],
			enc.decodeMap[src2[3]],
			enc.decodeMap[src2[4]],
			enc.decodeMap[src2[5]],
			enc.decodeMap[src2[6]],
			enc.decodeMap[src2[7]],
		); ok {
//line /usr/local/go/src/encoding/base64/base64.go:506
			_go_fuzz_dep_.CoverTab[10700]++
									binary.BigEndian.PutUint64(dst[n:], dn)
									n += 6
									si += 8
//line /usr/local/go/src/encoding/base64/base64.go:509
			// _ = "end of CoverTab[10700]"
		} else {
//line /usr/local/go/src/encoding/base64/base64.go:510
			_go_fuzz_dep_.CoverTab[10701]++
									var ninc int
									si, ninc, err = enc.decodeQuantum(dst[n:], src, si)
									n += ninc
									if err != nil {
//line /usr/local/go/src/encoding/base64/base64.go:514
				_go_fuzz_dep_.CoverTab[10702]++
										return n, err
//line /usr/local/go/src/encoding/base64/base64.go:515
				// _ = "end of CoverTab[10702]"
			} else {
//line /usr/local/go/src/encoding/base64/base64.go:516
				_go_fuzz_dep_.CoverTab[10703]++
//line /usr/local/go/src/encoding/base64/base64.go:516
				// _ = "end of CoverTab[10703]"
//line /usr/local/go/src/encoding/base64/base64.go:516
			}
//line /usr/local/go/src/encoding/base64/base64.go:516
			// _ = "end of CoverTab[10701]"
		}
//line /usr/local/go/src/encoding/base64/base64.go:517
		// _ = "end of CoverTab[10699]"
	}
//line /usr/local/go/src/encoding/base64/base64.go:518
	// _ = "end of CoverTab[10691]"
//line /usr/local/go/src/encoding/base64/base64.go:518
	_go_fuzz_dep_.CoverTab[10692]++

							for len(src)-si >= 4 && func() bool {
//line /usr/local/go/src/encoding/base64/base64.go:520
		_go_fuzz_dep_.CoverTab[10704]++
//line /usr/local/go/src/encoding/base64/base64.go:520
		return len(dst)-n >= 4
//line /usr/local/go/src/encoding/base64/base64.go:520
		// _ = "end of CoverTab[10704]"
//line /usr/local/go/src/encoding/base64/base64.go:520
	}() {
//line /usr/local/go/src/encoding/base64/base64.go:520
		_go_fuzz_dep_.CoverTab[10705]++
								src2 := src[si : si+4]
								if dn, ok := assemble32(
			enc.decodeMap[src2[0]],
			enc.decodeMap[src2[1]],
			enc.decodeMap[src2[2]],
			enc.decodeMap[src2[3]],
		); ok {
//line /usr/local/go/src/encoding/base64/base64.go:527
			_go_fuzz_dep_.CoverTab[10706]++
									binary.BigEndian.PutUint32(dst[n:], dn)
									n += 3
									si += 4
//line /usr/local/go/src/encoding/base64/base64.go:530
			// _ = "end of CoverTab[10706]"
		} else {
//line /usr/local/go/src/encoding/base64/base64.go:531
			_go_fuzz_dep_.CoverTab[10707]++
									var ninc int
									si, ninc, err = enc.decodeQuantum(dst[n:], src, si)
									n += ninc
									if err != nil {
//line /usr/local/go/src/encoding/base64/base64.go:535
				_go_fuzz_dep_.CoverTab[10708]++
										return n, err
//line /usr/local/go/src/encoding/base64/base64.go:536
				// _ = "end of CoverTab[10708]"
			} else {
//line /usr/local/go/src/encoding/base64/base64.go:537
				_go_fuzz_dep_.CoverTab[10709]++
//line /usr/local/go/src/encoding/base64/base64.go:537
				// _ = "end of CoverTab[10709]"
//line /usr/local/go/src/encoding/base64/base64.go:537
			}
//line /usr/local/go/src/encoding/base64/base64.go:537
			// _ = "end of CoverTab[10707]"
		}
//line /usr/local/go/src/encoding/base64/base64.go:538
		// _ = "end of CoverTab[10705]"
	}
//line /usr/local/go/src/encoding/base64/base64.go:539
	// _ = "end of CoverTab[10692]"
//line /usr/local/go/src/encoding/base64/base64.go:539
	_go_fuzz_dep_.CoverTab[10693]++

							for si < len(src) {
//line /usr/local/go/src/encoding/base64/base64.go:541
		_go_fuzz_dep_.CoverTab[10710]++
								var ninc int
								si, ninc, err = enc.decodeQuantum(dst[n:], src, si)
								n += ninc
								if err != nil {
//line /usr/local/go/src/encoding/base64/base64.go:545
			_go_fuzz_dep_.CoverTab[10711]++
									return n, err
//line /usr/local/go/src/encoding/base64/base64.go:546
			// _ = "end of CoverTab[10711]"
		} else {
//line /usr/local/go/src/encoding/base64/base64.go:547
			_go_fuzz_dep_.CoverTab[10712]++
//line /usr/local/go/src/encoding/base64/base64.go:547
			// _ = "end of CoverTab[10712]"
//line /usr/local/go/src/encoding/base64/base64.go:547
		}
//line /usr/local/go/src/encoding/base64/base64.go:547
		// _ = "end of CoverTab[10710]"
	}
//line /usr/local/go/src/encoding/base64/base64.go:548
	// _ = "end of CoverTab[10693]"
//line /usr/local/go/src/encoding/base64/base64.go:548
	_go_fuzz_dep_.CoverTab[10694]++
							return n, err
//line /usr/local/go/src/encoding/base64/base64.go:549
	// _ = "end of CoverTab[10694]"
}

// assemble32 assembles 4 base64 digits into 3 bytes.
//line /usr/local/go/src/encoding/base64/base64.go:552
// Each digit comes from the decode map, and will be 0xff
//line /usr/local/go/src/encoding/base64/base64.go:552
// if it came from an invalid character.
//line /usr/local/go/src/encoding/base64/base64.go:555
func assemble32(n1, n2, n3, n4 byte) (dn uint32, ok bool) {
//line /usr/local/go/src/encoding/base64/base64.go:555
	_go_fuzz_dep_.CoverTab[10713]++

//line /usr/local/go/src/encoding/base64/base64.go:558
	if n1|n2|n3|n4 == 0xff {
//line /usr/local/go/src/encoding/base64/base64.go:558
		_go_fuzz_dep_.CoverTab[10715]++
								return 0, false
//line /usr/local/go/src/encoding/base64/base64.go:559
		// _ = "end of CoverTab[10715]"
	} else {
//line /usr/local/go/src/encoding/base64/base64.go:560
		_go_fuzz_dep_.CoverTab[10716]++
//line /usr/local/go/src/encoding/base64/base64.go:560
		// _ = "end of CoverTab[10716]"
//line /usr/local/go/src/encoding/base64/base64.go:560
	}
//line /usr/local/go/src/encoding/base64/base64.go:560
	// _ = "end of CoverTab[10713]"
//line /usr/local/go/src/encoding/base64/base64.go:560
	_go_fuzz_dep_.CoverTab[10714]++
							return uint32(n1)<<26 |
			uint32(n2)<<20 |
			uint32(n3)<<14 |
			uint32(n4)<<8,
		true
//line /usr/local/go/src/encoding/base64/base64.go:565
	// _ = "end of CoverTab[10714]"
}

// assemble64 assembles 8 base64 digits into 6 bytes.
//line /usr/local/go/src/encoding/base64/base64.go:568
// Each digit comes from the decode map, and will be 0xff
//line /usr/local/go/src/encoding/base64/base64.go:568
// if it came from an invalid character.
//line /usr/local/go/src/encoding/base64/base64.go:571
func assemble64(n1, n2, n3, n4, n5, n6, n7, n8 byte) (dn uint64, ok bool) {
//line /usr/local/go/src/encoding/base64/base64.go:571
	_go_fuzz_dep_.CoverTab[10717]++

//line /usr/local/go/src/encoding/base64/base64.go:574
	if n1|n2|n3|n4|n5|n6|n7|n8 == 0xff {
//line /usr/local/go/src/encoding/base64/base64.go:574
		_go_fuzz_dep_.CoverTab[10719]++
								return 0, false
//line /usr/local/go/src/encoding/base64/base64.go:575
		// _ = "end of CoverTab[10719]"
	} else {
//line /usr/local/go/src/encoding/base64/base64.go:576
		_go_fuzz_dep_.CoverTab[10720]++
//line /usr/local/go/src/encoding/base64/base64.go:576
		// _ = "end of CoverTab[10720]"
//line /usr/local/go/src/encoding/base64/base64.go:576
	}
//line /usr/local/go/src/encoding/base64/base64.go:576
	// _ = "end of CoverTab[10717]"
//line /usr/local/go/src/encoding/base64/base64.go:576
	_go_fuzz_dep_.CoverTab[10718]++
							return uint64(n1)<<58 |
			uint64(n2)<<52 |
			uint64(n3)<<46 |
			uint64(n4)<<40 |
			uint64(n5)<<34 |
			uint64(n6)<<28 |
			uint64(n7)<<22 |
			uint64(n8)<<16,
		true
//line /usr/local/go/src/encoding/base64/base64.go:585
	// _ = "end of CoverTab[10718]"
}

type newlineFilteringReader struct {
	wrapped io.Reader
}

func (r *newlineFilteringReader) Read(p []byte) (int, error) {
//line /usr/local/go/src/encoding/base64/base64.go:592
	_go_fuzz_dep_.CoverTab[10721]++
							n, err := r.wrapped.Read(p)
							for n > 0 {
//line /usr/local/go/src/encoding/base64/base64.go:594
		_go_fuzz_dep_.CoverTab[10723]++
								offset := 0
								for i, b := range p[:n] {
//line /usr/local/go/src/encoding/base64/base64.go:596
			_go_fuzz_dep_.CoverTab[10726]++
									if b != '\r' && func() bool {
//line /usr/local/go/src/encoding/base64/base64.go:597
				_go_fuzz_dep_.CoverTab[10727]++
//line /usr/local/go/src/encoding/base64/base64.go:597
				return b != '\n'
//line /usr/local/go/src/encoding/base64/base64.go:597
				// _ = "end of CoverTab[10727]"
//line /usr/local/go/src/encoding/base64/base64.go:597
			}() {
//line /usr/local/go/src/encoding/base64/base64.go:597
				_go_fuzz_dep_.CoverTab[10728]++
										if i != offset {
//line /usr/local/go/src/encoding/base64/base64.go:598
					_go_fuzz_dep_.CoverTab[10730]++
											p[offset] = b
//line /usr/local/go/src/encoding/base64/base64.go:599
					// _ = "end of CoverTab[10730]"
				} else {
//line /usr/local/go/src/encoding/base64/base64.go:600
					_go_fuzz_dep_.CoverTab[10731]++
//line /usr/local/go/src/encoding/base64/base64.go:600
					// _ = "end of CoverTab[10731]"
//line /usr/local/go/src/encoding/base64/base64.go:600
				}
//line /usr/local/go/src/encoding/base64/base64.go:600
				// _ = "end of CoverTab[10728]"
//line /usr/local/go/src/encoding/base64/base64.go:600
				_go_fuzz_dep_.CoverTab[10729]++
										offset++
//line /usr/local/go/src/encoding/base64/base64.go:601
				// _ = "end of CoverTab[10729]"
			} else {
//line /usr/local/go/src/encoding/base64/base64.go:602
				_go_fuzz_dep_.CoverTab[10732]++
//line /usr/local/go/src/encoding/base64/base64.go:602
				// _ = "end of CoverTab[10732]"
//line /usr/local/go/src/encoding/base64/base64.go:602
			}
//line /usr/local/go/src/encoding/base64/base64.go:602
			// _ = "end of CoverTab[10726]"
		}
//line /usr/local/go/src/encoding/base64/base64.go:603
		// _ = "end of CoverTab[10723]"
//line /usr/local/go/src/encoding/base64/base64.go:603
		_go_fuzz_dep_.CoverTab[10724]++
								if offset > 0 {
//line /usr/local/go/src/encoding/base64/base64.go:604
			_go_fuzz_dep_.CoverTab[10733]++
									return offset, err
//line /usr/local/go/src/encoding/base64/base64.go:605
			// _ = "end of CoverTab[10733]"
		} else {
//line /usr/local/go/src/encoding/base64/base64.go:606
			_go_fuzz_dep_.CoverTab[10734]++
//line /usr/local/go/src/encoding/base64/base64.go:606
			// _ = "end of CoverTab[10734]"
//line /usr/local/go/src/encoding/base64/base64.go:606
		}
//line /usr/local/go/src/encoding/base64/base64.go:606
		// _ = "end of CoverTab[10724]"
//line /usr/local/go/src/encoding/base64/base64.go:606
		_go_fuzz_dep_.CoverTab[10725]++

								n, err = r.wrapped.Read(p)
//line /usr/local/go/src/encoding/base64/base64.go:608
		// _ = "end of CoverTab[10725]"
	}
//line /usr/local/go/src/encoding/base64/base64.go:609
	// _ = "end of CoverTab[10721]"
//line /usr/local/go/src/encoding/base64/base64.go:609
	_go_fuzz_dep_.CoverTab[10722]++
							return n, err
//line /usr/local/go/src/encoding/base64/base64.go:610
	// _ = "end of CoverTab[10722]"
}

// NewDecoder constructs a new base64 stream decoder.
func NewDecoder(enc *Encoding, r io.Reader) io.Reader {
//line /usr/local/go/src/encoding/base64/base64.go:614
	_go_fuzz_dep_.CoverTab[10735]++
							return &decoder{enc: enc, r: &newlineFilteringReader{r}}
//line /usr/local/go/src/encoding/base64/base64.go:615
	// _ = "end of CoverTab[10735]"
}

// DecodedLen returns the maximum length in bytes of the decoded data
//line /usr/local/go/src/encoding/base64/base64.go:618
// corresponding to n bytes of base64-encoded data.
//line /usr/local/go/src/encoding/base64/base64.go:620
func (enc *Encoding) DecodedLen(n int) int {
//line /usr/local/go/src/encoding/base64/base64.go:620
	_go_fuzz_dep_.CoverTab[10736]++
							if enc.padChar == NoPadding {
//line /usr/local/go/src/encoding/base64/base64.go:621
		_go_fuzz_dep_.CoverTab[10738]++

								return n * 6 / 8
//line /usr/local/go/src/encoding/base64/base64.go:623
		// _ = "end of CoverTab[10738]"
	} else {
//line /usr/local/go/src/encoding/base64/base64.go:624
		_go_fuzz_dep_.CoverTab[10739]++
//line /usr/local/go/src/encoding/base64/base64.go:624
		// _ = "end of CoverTab[10739]"
//line /usr/local/go/src/encoding/base64/base64.go:624
	}
//line /usr/local/go/src/encoding/base64/base64.go:624
	// _ = "end of CoverTab[10736]"
//line /usr/local/go/src/encoding/base64/base64.go:624
	_go_fuzz_dep_.CoverTab[10737]++

							return n / 4 * 3
//line /usr/local/go/src/encoding/base64/base64.go:626
	// _ = "end of CoverTab[10737]"
}

//line /usr/local/go/src/encoding/base64/base64.go:627
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/encoding/base64/base64.go:627
var _ = _go_fuzz_dep_.CoverTab
