// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/mime/encodedword.go:5
package mime

//line /usr/local/go/src/mime/encodedword.go:5
import (
//line /usr/local/go/src/mime/encodedword.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/mime/encodedword.go:5
)
//line /usr/local/go/src/mime/encodedword.go:5
import (
//line /usr/local/go/src/mime/encodedword.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/mime/encodedword.go:5
)

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"strings"
	"unicode"
	"unicode/utf8"
)

// A WordEncoder is an RFC 2047 encoded-word encoder.
type WordEncoder byte

const (
	// BEncoding represents Base64 encoding scheme as defined by RFC 2045.
	BEncoding	= WordEncoder('b')
	// QEncoding represents the Q-encoding scheme as defined by RFC 2047.
	QEncoding	= WordEncoder('q')
)

var (
	errInvalidWord = errors.New("mime: invalid RFC 2047 encoded-word")
)

// Encode returns the encoded-word form of s. If s is ASCII without special
//line /usr/local/go/src/mime/encodedword.go:32
// characters, it is returned unchanged. The provided charset is the IANA
//line /usr/local/go/src/mime/encodedword.go:32
// charset name of s. It is case insensitive.
//line /usr/local/go/src/mime/encodedword.go:35
func (e WordEncoder) Encode(charset, s string) string {
//line /usr/local/go/src/mime/encodedword.go:35
	_go_fuzz_dep_.CoverTab[35433]++
							if !needsEncoding(s) {
//line /usr/local/go/src/mime/encodedword.go:36
		_go_fuzz_dep_.CoverTab[35435]++
								return s
//line /usr/local/go/src/mime/encodedword.go:37
		// _ = "end of CoverTab[35435]"
	} else {
//line /usr/local/go/src/mime/encodedword.go:38
		_go_fuzz_dep_.CoverTab[35436]++
//line /usr/local/go/src/mime/encodedword.go:38
		// _ = "end of CoverTab[35436]"
//line /usr/local/go/src/mime/encodedword.go:38
	}
//line /usr/local/go/src/mime/encodedword.go:38
	// _ = "end of CoverTab[35433]"
//line /usr/local/go/src/mime/encodedword.go:38
	_go_fuzz_dep_.CoverTab[35434]++
							return e.encodeWord(charset, s)
//line /usr/local/go/src/mime/encodedword.go:39
	// _ = "end of CoverTab[35434]"
}

func needsEncoding(s string) bool {
//line /usr/local/go/src/mime/encodedword.go:42
	_go_fuzz_dep_.CoverTab[35437]++
							for _, b := range s {
//line /usr/local/go/src/mime/encodedword.go:43
		_go_fuzz_dep_.CoverTab[35439]++
								if (b < ' ' || func() bool {
//line /usr/local/go/src/mime/encodedword.go:44
			_go_fuzz_dep_.CoverTab[35440]++
//line /usr/local/go/src/mime/encodedword.go:44
			return b > '~'
//line /usr/local/go/src/mime/encodedword.go:44
			// _ = "end of CoverTab[35440]"
//line /usr/local/go/src/mime/encodedword.go:44
		}()) && func() bool {
//line /usr/local/go/src/mime/encodedword.go:44
			_go_fuzz_dep_.CoverTab[35441]++
//line /usr/local/go/src/mime/encodedword.go:44
			return b != '\t'
//line /usr/local/go/src/mime/encodedword.go:44
			// _ = "end of CoverTab[35441]"
//line /usr/local/go/src/mime/encodedword.go:44
		}() {
//line /usr/local/go/src/mime/encodedword.go:44
			_go_fuzz_dep_.CoverTab[35442]++
									return true
//line /usr/local/go/src/mime/encodedword.go:45
			// _ = "end of CoverTab[35442]"
		} else {
//line /usr/local/go/src/mime/encodedword.go:46
			_go_fuzz_dep_.CoverTab[35443]++
//line /usr/local/go/src/mime/encodedword.go:46
			// _ = "end of CoverTab[35443]"
//line /usr/local/go/src/mime/encodedword.go:46
		}
//line /usr/local/go/src/mime/encodedword.go:46
		// _ = "end of CoverTab[35439]"
	}
//line /usr/local/go/src/mime/encodedword.go:47
	// _ = "end of CoverTab[35437]"
//line /usr/local/go/src/mime/encodedword.go:47
	_go_fuzz_dep_.CoverTab[35438]++
							return false
//line /usr/local/go/src/mime/encodedword.go:48
	// _ = "end of CoverTab[35438]"
}

// encodeWord encodes a string into an encoded-word.
func (e WordEncoder) encodeWord(charset, s string) string {
//line /usr/local/go/src/mime/encodedword.go:52
	_go_fuzz_dep_.CoverTab[35444]++
							var buf strings.Builder

//line /usr/local/go/src/mime/encodedword.go:57
	buf.Grow(48)

	e.openWord(&buf, charset)
	if e == BEncoding {
//line /usr/local/go/src/mime/encodedword.go:60
		_go_fuzz_dep_.CoverTab[35446]++
								e.bEncode(&buf, charset, s)
//line /usr/local/go/src/mime/encodedword.go:61
		// _ = "end of CoverTab[35446]"
	} else {
//line /usr/local/go/src/mime/encodedword.go:62
		_go_fuzz_dep_.CoverTab[35447]++
								e.qEncode(&buf, charset, s)
//line /usr/local/go/src/mime/encodedword.go:63
		// _ = "end of CoverTab[35447]"
	}
//line /usr/local/go/src/mime/encodedword.go:64
	// _ = "end of CoverTab[35444]"
//line /usr/local/go/src/mime/encodedword.go:64
	_go_fuzz_dep_.CoverTab[35445]++
							closeWord(&buf)

							return buf.String()
//line /usr/local/go/src/mime/encodedword.go:67
	// _ = "end of CoverTab[35445]"
}

const (
	// The maximum length of an encoded-word is 75 characters.
	// See RFC 2047, section 2.
	maxEncodedWordLen	= 75
	// maxContentLen is how much content can be encoded, ignoring the header and
	// 2-byte footer.
	maxContentLen	= maxEncodedWordLen - len("=?UTF-8?q?") - len("?=")
)

var maxBase64Len = base64.StdEncoding.DecodedLen(maxContentLen)

// bEncode encodes s using base64 encoding and writes it to buf.
func (e WordEncoder) bEncode(buf *strings.Builder, charset, s string) {
//line /usr/local/go/src/mime/encodedword.go:82
	_go_fuzz_dep_.CoverTab[35448]++
							w := base64.NewEncoder(base64.StdEncoding, buf)

//line /usr/local/go/src/mime/encodedword.go:86
	if !isUTF8(charset) || func() bool {
//line /usr/local/go/src/mime/encodedword.go:86
		_go_fuzz_dep_.CoverTab[35451]++
//line /usr/local/go/src/mime/encodedword.go:86
		return base64.StdEncoding.EncodedLen(len(s)) <= maxContentLen
//line /usr/local/go/src/mime/encodedword.go:86
		// _ = "end of CoverTab[35451]"
//line /usr/local/go/src/mime/encodedword.go:86
	}() {
//line /usr/local/go/src/mime/encodedword.go:86
		_go_fuzz_dep_.CoverTab[35452]++
								io.WriteString(w, s)
								w.Close()
								return
//line /usr/local/go/src/mime/encodedword.go:89
		// _ = "end of CoverTab[35452]"
	} else {
//line /usr/local/go/src/mime/encodedword.go:90
		_go_fuzz_dep_.CoverTab[35453]++
//line /usr/local/go/src/mime/encodedword.go:90
		// _ = "end of CoverTab[35453]"
//line /usr/local/go/src/mime/encodedword.go:90
	}
//line /usr/local/go/src/mime/encodedword.go:90
	// _ = "end of CoverTab[35448]"
//line /usr/local/go/src/mime/encodedword.go:90
	_go_fuzz_dep_.CoverTab[35449]++

							var currentLen, last, runeLen int
							for i := 0; i < len(s); i += runeLen {
//line /usr/local/go/src/mime/encodedword.go:93
		_go_fuzz_dep_.CoverTab[35454]++

//line /usr/local/go/src/mime/encodedword.go:96
		_, runeLen = utf8.DecodeRuneInString(s[i:])

		if currentLen+runeLen <= maxBase64Len {
//line /usr/local/go/src/mime/encodedword.go:98
			_go_fuzz_dep_.CoverTab[35455]++
									currentLen += runeLen
//line /usr/local/go/src/mime/encodedword.go:99
			// _ = "end of CoverTab[35455]"
		} else {
//line /usr/local/go/src/mime/encodedword.go:100
			_go_fuzz_dep_.CoverTab[35456]++
									io.WriteString(w, s[last:i])
									w.Close()
									e.splitWord(buf, charset)
									last = i
									currentLen = runeLen
//line /usr/local/go/src/mime/encodedword.go:105
			// _ = "end of CoverTab[35456]"
		}
//line /usr/local/go/src/mime/encodedword.go:106
		// _ = "end of CoverTab[35454]"
	}
//line /usr/local/go/src/mime/encodedword.go:107
	// _ = "end of CoverTab[35449]"
//line /usr/local/go/src/mime/encodedword.go:107
	_go_fuzz_dep_.CoverTab[35450]++
							io.WriteString(w, s[last:])
							w.Close()
//line /usr/local/go/src/mime/encodedword.go:109
	// _ = "end of CoverTab[35450]"
}

// qEncode encodes s using Q encoding and writes it to buf. It splits the
//line /usr/local/go/src/mime/encodedword.go:112
// encoded-words when necessary.
//line /usr/local/go/src/mime/encodedword.go:114
func (e WordEncoder) qEncode(buf *strings.Builder, charset, s string) {
//line /usr/local/go/src/mime/encodedword.go:114
	_go_fuzz_dep_.CoverTab[35457]++

							if !isUTF8(charset) {
//line /usr/local/go/src/mime/encodedword.go:116
		_go_fuzz_dep_.CoverTab[35459]++
								writeQString(buf, s)
								return
//line /usr/local/go/src/mime/encodedword.go:118
		// _ = "end of CoverTab[35459]"
	} else {
//line /usr/local/go/src/mime/encodedword.go:119
		_go_fuzz_dep_.CoverTab[35460]++
//line /usr/local/go/src/mime/encodedword.go:119
		// _ = "end of CoverTab[35460]"
//line /usr/local/go/src/mime/encodedword.go:119
	}
//line /usr/local/go/src/mime/encodedword.go:119
	// _ = "end of CoverTab[35457]"
//line /usr/local/go/src/mime/encodedword.go:119
	_go_fuzz_dep_.CoverTab[35458]++

							var currentLen, runeLen int
							for i := 0; i < len(s); i += runeLen {
//line /usr/local/go/src/mime/encodedword.go:122
		_go_fuzz_dep_.CoverTab[35461]++
								b := s[i]
		// Multi-byte characters must not be split across encoded-words.
		// See RFC 2047, section 5.3.
		var encLen int
		if b >= ' ' && func() bool {
//line /usr/local/go/src/mime/encodedword.go:127
			_go_fuzz_dep_.CoverTab[35464]++
//line /usr/local/go/src/mime/encodedword.go:127
			return b <= '~'
//line /usr/local/go/src/mime/encodedword.go:127
			// _ = "end of CoverTab[35464]"
//line /usr/local/go/src/mime/encodedword.go:127
		}() && func() bool {
//line /usr/local/go/src/mime/encodedword.go:127
			_go_fuzz_dep_.CoverTab[35465]++
//line /usr/local/go/src/mime/encodedword.go:127
			return b != '='
//line /usr/local/go/src/mime/encodedword.go:127
			// _ = "end of CoverTab[35465]"
//line /usr/local/go/src/mime/encodedword.go:127
		}() && func() bool {
//line /usr/local/go/src/mime/encodedword.go:127
			_go_fuzz_dep_.CoverTab[35466]++
//line /usr/local/go/src/mime/encodedword.go:127
			return b != '?'
//line /usr/local/go/src/mime/encodedword.go:127
			// _ = "end of CoverTab[35466]"
//line /usr/local/go/src/mime/encodedword.go:127
		}() && func() bool {
//line /usr/local/go/src/mime/encodedword.go:127
			_go_fuzz_dep_.CoverTab[35467]++
//line /usr/local/go/src/mime/encodedword.go:127
			return b != '_'
//line /usr/local/go/src/mime/encodedword.go:127
			// _ = "end of CoverTab[35467]"
//line /usr/local/go/src/mime/encodedword.go:127
		}() {
//line /usr/local/go/src/mime/encodedword.go:127
			_go_fuzz_dep_.CoverTab[35468]++
									runeLen, encLen = 1, 1
//line /usr/local/go/src/mime/encodedword.go:128
			// _ = "end of CoverTab[35468]"
		} else {
//line /usr/local/go/src/mime/encodedword.go:129
			_go_fuzz_dep_.CoverTab[35469]++
									_, runeLen = utf8.DecodeRuneInString(s[i:])
									encLen = 3 * runeLen
//line /usr/local/go/src/mime/encodedword.go:131
			// _ = "end of CoverTab[35469]"
		}
//line /usr/local/go/src/mime/encodedword.go:132
		// _ = "end of CoverTab[35461]"
//line /usr/local/go/src/mime/encodedword.go:132
		_go_fuzz_dep_.CoverTab[35462]++

								if currentLen+encLen > maxContentLen {
//line /usr/local/go/src/mime/encodedword.go:134
			_go_fuzz_dep_.CoverTab[35470]++
									e.splitWord(buf, charset)
									currentLen = 0
//line /usr/local/go/src/mime/encodedword.go:136
			// _ = "end of CoverTab[35470]"
		} else {
//line /usr/local/go/src/mime/encodedword.go:137
			_go_fuzz_dep_.CoverTab[35471]++
//line /usr/local/go/src/mime/encodedword.go:137
			// _ = "end of CoverTab[35471]"
//line /usr/local/go/src/mime/encodedword.go:137
		}
//line /usr/local/go/src/mime/encodedword.go:137
		// _ = "end of CoverTab[35462]"
//line /usr/local/go/src/mime/encodedword.go:137
		_go_fuzz_dep_.CoverTab[35463]++
								writeQString(buf, s[i:i+runeLen])
								currentLen += encLen
//line /usr/local/go/src/mime/encodedword.go:139
		// _ = "end of CoverTab[35463]"
	}
//line /usr/local/go/src/mime/encodedword.go:140
	// _ = "end of CoverTab[35458]"
}

// writeQString encodes s using Q encoding and writes it to buf.
func writeQString(buf *strings.Builder, s string) {
//line /usr/local/go/src/mime/encodedword.go:144
	_go_fuzz_dep_.CoverTab[35472]++
							for i := 0; i < len(s); i++ {
//line /usr/local/go/src/mime/encodedword.go:145
		_go_fuzz_dep_.CoverTab[35473]++
								switch b := s[i]; {
		case b == ' ':
//line /usr/local/go/src/mime/encodedword.go:147
			_go_fuzz_dep_.CoverTab[35474]++
									buf.WriteByte('_')
//line /usr/local/go/src/mime/encodedword.go:148
			// _ = "end of CoverTab[35474]"
		case b >= '!' && func() bool {
//line /usr/local/go/src/mime/encodedword.go:149
			_go_fuzz_dep_.CoverTab[35477]++
//line /usr/local/go/src/mime/encodedword.go:149
			return b <= '~'
//line /usr/local/go/src/mime/encodedword.go:149
			// _ = "end of CoverTab[35477]"
//line /usr/local/go/src/mime/encodedword.go:149
		}() && func() bool {
//line /usr/local/go/src/mime/encodedword.go:149
			_go_fuzz_dep_.CoverTab[35478]++
//line /usr/local/go/src/mime/encodedword.go:149
			return b != '='
//line /usr/local/go/src/mime/encodedword.go:149
			// _ = "end of CoverTab[35478]"
//line /usr/local/go/src/mime/encodedword.go:149
		}() && func() bool {
//line /usr/local/go/src/mime/encodedword.go:149
			_go_fuzz_dep_.CoverTab[35479]++
//line /usr/local/go/src/mime/encodedword.go:149
			return b != '?'
//line /usr/local/go/src/mime/encodedword.go:149
			// _ = "end of CoverTab[35479]"
//line /usr/local/go/src/mime/encodedword.go:149
		}() && func() bool {
//line /usr/local/go/src/mime/encodedword.go:149
			_go_fuzz_dep_.CoverTab[35480]++
//line /usr/local/go/src/mime/encodedword.go:149
			return b != '_'
//line /usr/local/go/src/mime/encodedword.go:149
			// _ = "end of CoverTab[35480]"
//line /usr/local/go/src/mime/encodedword.go:149
		}():
//line /usr/local/go/src/mime/encodedword.go:149
			_go_fuzz_dep_.CoverTab[35475]++
									buf.WriteByte(b)
//line /usr/local/go/src/mime/encodedword.go:150
			// _ = "end of CoverTab[35475]"
		default:
//line /usr/local/go/src/mime/encodedword.go:151
			_go_fuzz_dep_.CoverTab[35476]++
									buf.WriteByte('=')
									buf.WriteByte(upperhex[b>>4])
									buf.WriteByte(upperhex[b&0x0f])
//line /usr/local/go/src/mime/encodedword.go:154
			// _ = "end of CoverTab[35476]"
		}
//line /usr/local/go/src/mime/encodedword.go:155
		// _ = "end of CoverTab[35473]"
	}
//line /usr/local/go/src/mime/encodedword.go:156
	// _ = "end of CoverTab[35472]"
}

// openWord writes the beginning of an encoded-word into buf.
func (e WordEncoder) openWord(buf *strings.Builder, charset string) {
//line /usr/local/go/src/mime/encodedword.go:160
	_go_fuzz_dep_.CoverTab[35481]++
							buf.WriteString("=?")
							buf.WriteString(charset)
							buf.WriteByte('?')
							buf.WriteByte(byte(e))
							buf.WriteByte('?')
//line /usr/local/go/src/mime/encodedword.go:165
	// _ = "end of CoverTab[35481]"
}

// closeWord writes the end of an encoded-word into buf.
func closeWord(buf *strings.Builder) {
//line /usr/local/go/src/mime/encodedword.go:169
	_go_fuzz_dep_.CoverTab[35482]++
							buf.WriteString("?=")
//line /usr/local/go/src/mime/encodedword.go:170
	// _ = "end of CoverTab[35482]"
}

// splitWord closes the current encoded-word and opens a new one.
func (e WordEncoder) splitWord(buf *strings.Builder, charset string) {
//line /usr/local/go/src/mime/encodedword.go:174
	_go_fuzz_dep_.CoverTab[35483]++
							closeWord(buf)
							buf.WriteByte(' ')
							e.openWord(buf, charset)
//line /usr/local/go/src/mime/encodedword.go:177
	// _ = "end of CoverTab[35483]"
}

func isUTF8(charset string) bool {
//line /usr/local/go/src/mime/encodedword.go:180
	_go_fuzz_dep_.CoverTab[35484]++
							return strings.EqualFold(charset, "UTF-8")
//line /usr/local/go/src/mime/encodedword.go:181
	// _ = "end of CoverTab[35484]"
}

const upperhex = "0123456789ABCDEF"

// A WordDecoder decodes MIME headers containing RFC 2047 encoded-words.
type WordDecoder struct {
	// CharsetReader, if non-nil, defines a function to generate
	// charset-conversion readers, converting from the provided
	// charset into UTF-8.
	// Charsets are always lower-case. utf-8, iso-8859-1 and us-ascii charsets
	// are handled by default.
	// One of the CharsetReader's result values must be non-nil.
	CharsetReader func(charset string, input io.Reader) (io.Reader, error)
}

// Decode decodes an RFC 2047 encoded-word.
func (d *WordDecoder) Decode(word string) (string, error) {
//line /usr/local/go/src/mime/encodedword.go:198
	_go_fuzz_dep_.CoverTab[35485]++

//line /usr/local/go/src/mime/encodedword.go:201
	if len(word) < 8 || func() bool {
//line /usr/local/go/src/mime/encodedword.go:201
		_go_fuzz_dep_.CoverTab[35491]++
//line /usr/local/go/src/mime/encodedword.go:201
		return !strings.HasPrefix(word, "=?")
//line /usr/local/go/src/mime/encodedword.go:201
		// _ = "end of CoverTab[35491]"
//line /usr/local/go/src/mime/encodedword.go:201
	}() || func() bool {
//line /usr/local/go/src/mime/encodedword.go:201
		_go_fuzz_dep_.CoverTab[35492]++
//line /usr/local/go/src/mime/encodedword.go:201
		return !strings.HasSuffix(word, "?=")
//line /usr/local/go/src/mime/encodedword.go:201
		// _ = "end of CoverTab[35492]"
//line /usr/local/go/src/mime/encodedword.go:201
	}() || func() bool {
//line /usr/local/go/src/mime/encodedword.go:201
		_go_fuzz_dep_.CoverTab[35493]++
//line /usr/local/go/src/mime/encodedword.go:201
		return strings.Count(word, "?") != 4
//line /usr/local/go/src/mime/encodedword.go:201
		// _ = "end of CoverTab[35493]"
//line /usr/local/go/src/mime/encodedword.go:201
	}() {
//line /usr/local/go/src/mime/encodedword.go:201
		_go_fuzz_dep_.CoverTab[35494]++
								return "", errInvalidWord
//line /usr/local/go/src/mime/encodedword.go:202
		// _ = "end of CoverTab[35494]"
	} else {
//line /usr/local/go/src/mime/encodedword.go:203
		_go_fuzz_dep_.CoverTab[35495]++
//line /usr/local/go/src/mime/encodedword.go:203
		// _ = "end of CoverTab[35495]"
//line /usr/local/go/src/mime/encodedword.go:203
	}
//line /usr/local/go/src/mime/encodedword.go:203
	// _ = "end of CoverTab[35485]"
//line /usr/local/go/src/mime/encodedword.go:203
	_go_fuzz_dep_.CoverTab[35486]++
							word = word[2 : len(word)-2]

//line /usr/local/go/src/mime/encodedword.go:207
	charset, text, _ := strings.Cut(word, "?")
	if charset == "" {
//line /usr/local/go/src/mime/encodedword.go:208
		_go_fuzz_dep_.CoverTab[35496]++
								return "", errInvalidWord
//line /usr/local/go/src/mime/encodedword.go:209
		// _ = "end of CoverTab[35496]"
	} else {
//line /usr/local/go/src/mime/encodedword.go:210
		_go_fuzz_dep_.CoverTab[35497]++
//line /usr/local/go/src/mime/encodedword.go:210
		// _ = "end of CoverTab[35497]"
//line /usr/local/go/src/mime/encodedword.go:210
	}
//line /usr/local/go/src/mime/encodedword.go:210
	// _ = "end of CoverTab[35486]"
//line /usr/local/go/src/mime/encodedword.go:210
	_go_fuzz_dep_.CoverTab[35487]++
							encoding, text, _ := strings.Cut(text, "?")
							if len(encoding) != 1 {
//line /usr/local/go/src/mime/encodedword.go:212
		_go_fuzz_dep_.CoverTab[35498]++
								return "", errInvalidWord
//line /usr/local/go/src/mime/encodedword.go:213
		// _ = "end of CoverTab[35498]"
	} else {
//line /usr/local/go/src/mime/encodedword.go:214
		_go_fuzz_dep_.CoverTab[35499]++
//line /usr/local/go/src/mime/encodedword.go:214
		// _ = "end of CoverTab[35499]"
//line /usr/local/go/src/mime/encodedword.go:214
	}
//line /usr/local/go/src/mime/encodedword.go:214
	// _ = "end of CoverTab[35487]"
//line /usr/local/go/src/mime/encodedword.go:214
	_go_fuzz_dep_.CoverTab[35488]++

							content, err := decode(encoding[0], text)
							if err != nil {
//line /usr/local/go/src/mime/encodedword.go:217
		_go_fuzz_dep_.CoverTab[35500]++
								return "", err
//line /usr/local/go/src/mime/encodedword.go:218
		// _ = "end of CoverTab[35500]"
	} else {
//line /usr/local/go/src/mime/encodedword.go:219
		_go_fuzz_dep_.CoverTab[35501]++
//line /usr/local/go/src/mime/encodedword.go:219
		// _ = "end of CoverTab[35501]"
//line /usr/local/go/src/mime/encodedword.go:219
	}
//line /usr/local/go/src/mime/encodedword.go:219
	// _ = "end of CoverTab[35488]"
//line /usr/local/go/src/mime/encodedword.go:219
	_go_fuzz_dep_.CoverTab[35489]++

							var buf strings.Builder
							if err := d.convert(&buf, charset, content); err != nil {
//line /usr/local/go/src/mime/encodedword.go:222
		_go_fuzz_dep_.CoverTab[35502]++
								return "", err
//line /usr/local/go/src/mime/encodedword.go:223
		// _ = "end of CoverTab[35502]"
	} else {
//line /usr/local/go/src/mime/encodedword.go:224
		_go_fuzz_dep_.CoverTab[35503]++
//line /usr/local/go/src/mime/encodedword.go:224
		// _ = "end of CoverTab[35503]"
//line /usr/local/go/src/mime/encodedword.go:224
	}
//line /usr/local/go/src/mime/encodedword.go:224
	// _ = "end of CoverTab[35489]"
//line /usr/local/go/src/mime/encodedword.go:224
	_go_fuzz_dep_.CoverTab[35490]++
							return buf.String(), nil
//line /usr/local/go/src/mime/encodedword.go:225
	// _ = "end of CoverTab[35490]"
}

// DecodeHeader decodes all encoded-words of the given string. It returns an
//line /usr/local/go/src/mime/encodedword.go:228
// error if and only if CharsetReader of d returns an error.
//line /usr/local/go/src/mime/encodedword.go:230
func (d *WordDecoder) DecodeHeader(header string) (string, error) {
//line /usr/local/go/src/mime/encodedword.go:230
	_go_fuzz_dep_.CoverTab[35504]++

							i := strings.Index(header, "=?")
							if i == -1 {
//line /usr/local/go/src/mime/encodedword.go:233
		_go_fuzz_dep_.CoverTab[35508]++
								return header, nil
//line /usr/local/go/src/mime/encodedword.go:234
		// _ = "end of CoverTab[35508]"
	} else {
//line /usr/local/go/src/mime/encodedword.go:235
		_go_fuzz_dep_.CoverTab[35509]++
//line /usr/local/go/src/mime/encodedword.go:235
		// _ = "end of CoverTab[35509]"
//line /usr/local/go/src/mime/encodedword.go:235
	}
//line /usr/local/go/src/mime/encodedword.go:235
	// _ = "end of CoverTab[35504]"
//line /usr/local/go/src/mime/encodedword.go:235
	_go_fuzz_dep_.CoverTab[35505]++

							var buf strings.Builder

							buf.WriteString(header[:i])
							header = header[i:]

							betweenWords := false
							for {
//line /usr/local/go/src/mime/encodedword.go:243
		_go_fuzz_dep_.CoverTab[35510]++
								start := strings.Index(header, "=?")
								if start == -1 {
//line /usr/local/go/src/mime/encodedword.go:245
			_go_fuzz_dep_.CoverTab[35519]++
									break
//line /usr/local/go/src/mime/encodedword.go:246
			// _ = "end of CoverTab[35519]"
		} else {
//line /usr/local/go/src/mime/encodedword.go:247
			_go_fuzz_dep_.CoverTab[35520]++
//line /usr/local/go/src/mime/encodedword.go:247
			// _ = "end of CoverTab[35520]"
//line /usr/local/go/src/mime/encodedword.go:247
		}
//line /usr/local/go/src/mime/encodedword.go:247
		// _ = "end of CoverTab[35510]"
//line /usr/local/go/src/mime/encodedword.go:247
		_go_fuzz_dep_.CoverTab[35511]++
								cur := start + len("=?")

								i := strings.Index(header[cur:], "?")
								if i == -1 {
//line /usr/local/go/src/mime/encodedword.go:251
			_go_fuzz_dep_.CoverTab[35521]++
									break
//line /usr/local/go/src/mime/encodedword.go:252
			// _ = "end of CoverTab[35521]"
		} else {
//line /usr/local/go/src/mime/encodedword.go:253
			_go_fuzz_dep_.CoverTab[35522]++
//line /usr/local/go/src/mime/encodedword.go:253
			// _ = "end of CoverTab[35522]"
//line /usr/local/go/src/mime/encodedword.go:253
		}
//line /usr/local/go/src/mime/encodedword.go:253
		// _ = "end of CoverTab[35511]"
//line /usr/local/go/src/mime/encodedword.go:253
		_go_fuzz_dep_.CoverTab[35512]++
								charset := header[cur : cur+i]
								cur += i + len("?")

								if len(header) < cur+len("Q??=") {
//line /usr/local/go/src/mime/encodedword.go:257
			_go_fuzz_dep_.CoverTab[35523]++
									break
//line /usr/local/go/src/mime/encodedword.go:258
			// _ = "end of CoverTab[35523]"
		} else {
//line /usr/local/go/src/mime/encodedword.go:259
			_go_fuzz_dep_.CoverTab[35524]++
//line /usr/local/go/src/mime/encodedword.go:259
			// _ = "end of CoverTab[35524]"
//line /usr/local/go/src/mime/encodedword.go:259
		}
//line /usr/local/go/src/mime/encodedword.go:259
		// _ = "end of CoverTab[35512]"
//line /usr/local/go/src/mime/encodedword.go:259
		_go_fuzz_dep_.CoverTab[35513]++
								encoding := header[cur]
								cur++

								if header[cur] != '?' {
//line /usr/local/go/src/mime/encodedword.go:263
			_go_fuzz_dep_.CoverTab[35525]++
									break
//line /usr/local/go/src/mime/encodedword.go:264
			// _ = "end of CoverTab[35525]"
		} else {
//line /usr/local/go/src/mime/encodedword.go:265
			_go_fuzz_dep_.CoverTab[35526]++
//line /usr/local/go/src/mime/encodedword.go:265
			// _ = "end of CoverTab[35526]"
//line /usr/local/go/src/mime/encodedword.go:265
		}
//line /usr/local/go/src/mime/encodedword.go:265
		// _ = "end of CoverTab[35513]"
//line /usr/local/go/src/mime/encodedword.go:265
		_go_fuzz_dep_.CoverTab[35514]++
								cur++

								j := strings.Index(header[cur:], "?=")
								if j == -1 {
//line /usr/local/go/src/mime/encodedword.go:269
			_go_fuzz_dep_.CoverTab[35527]++
									break
//line /usr/local/go/src/mime/encodedword.go:270
			// _ = "end of CoverTab[35527]"
		} else {
//line /usr/local/go/src/mime/encodedword.go:271
			_go_fuzz_dep_.CoverTab[35528]++
//line /usr/local/go/src/mime/encodedword.go:271
			// _ = "end of CoverTab[35528]"
//line /usr/local/go/src/mime/encodedword.go:271
		}
//line /usr/local/go/src/mime/encodedword.go:271
		// _ = "end of CoverTab[35514]"
//line /usr/local/go/src/mime/encodedword.go:271
		_go_fuzz_dep_.CoverTab[35515]++
								text := header[cur : cur+j]
								end := cur + j + len("?=")

								content, err := decode(encoding, text)
								if err != nil {
//line /usr/local/go/src/mime/encodedword.go:276
			_go_fuzz_dep_.CoverTab[35529]++
									betweenWords = false
									buf.WriteString(header[:start+2])
									header = header[start+2:]
									continue
//line /usr/local/go/src/mime/encodedword.go:280
			// _ = "end of CoverTab[35529]"
		} else {
//line /usr/local/go/src/mime/encodedword.go:281
			_go_fuzz_dep_.CoverTab[35530]++
//line /usr/local/go/src/mime/encodedword.go:281
			// _ = "end of CoverTab[35530]"
//line /usr/local/go/src/mime/encodedword.go:281
		}
//line /usr/local/go/src/mime/encodedword.go:281
		// _ = "end of CoverTab[35515]"
//line /usr/local/go/src/mime/encodedword.go:281
		_go_fuzz_dep_.CoverTab[35516]++

//line /usr/local/go/src/mime/encodedword.go:285
		if start > 0 && func() bool {
//line /usr/local/go/src/mime/encodedword.go:285
			_go_fuzz_dep_.CoverTab[35531]++
//line /usr/local/go/src/mime/encodedword.go:285
			return (!betweenWords || func() bool {
//line /usr/local/go/src/mime/encodedword.go:285
				_go_fuzz_dep_.CoverTab[35532]++
//line /usr/local/go/src/mime/encodedword.go:285
				return hasNonWhitespace(header[:start])
//line /usr/local/go/src/mime/encodedword.go:285
				// _ = "end of CoverTab[35532]"
//line /usr/local/go/src/mime/encodedword.go:285
			}())
//line /usr/local/go/src/mime/encodedword.go:285
			// _ = "end of CoverTab[35531]"
//line /usr/local/go/src/mime/encodedword.go:285
		}() {
//line /usr/local/go/src/mime/encodedword.go:285
			_go_fuzz_dep_.CoverTab[35533]++
									buf.WriteString(header[:start])
//line /usr/local/go/src/mime/encodedword.go:286
			// _ = "end of CoverTab[35533]"
		} else {
//line /usr/local/go/src/mime/encodedword.go:287
			_go_fuzz_dep_.CoverTab[35534]++
//line /usr/local/go/src/mime/encodedword.go:287
			// _ = "end of CoverTab[35534]"
//line /usr/local/go/src/mime/encodedword.go:287
		}
//line /usr/local/go/src/mime/encodedword.go:287
		// _ = "end of CoverTab[35516]"
//line /usr/local/go/src/mime/encodedword.go:287
		_go_fuzz_dep_.CoverTab[35517]++

								if err := d.convert(&buf, charset, content); err != nil {
//line /usr/local/go/src/mime/encodedword.go:289
			_go_fuzz_dep_.CoverTab[35535]++
									return "", err
//line /usr/local/go/src/mime/encodedword.go:290
			// _ = "end of CoverTab[35535]"
		} else {
//line /usr/local/go/src/mime/encodedword.go:291
			_go_fuzz_dep_.CoverTab[35536]++
//line /usr/local/go/src/mime/encodedword.go:291
			// _ = "end of CoverTab[35536]"
//line /usr/local/go/src/mime/encodedword.go:291
		}
//line /usr/local/go/src/mime/encodedword.go:291
		// _ = "end of CoverTab[35517]"
//line /usr/local/go/src/mime/encodedword.go:291
		_go_fuzz_dep_.CoverTab[35518]++

								header = header[end:]
								betweenWords = true
//line /usr/local/go/src/mime/encodedword.go:294
		// _ = "end of CoverTab[35518]"
	}
//line /usr/local/go/src/mime/encodedword.go:295
	// _ = "end of CoverTab[35505]"
//line /usr/local/go/src/mime/encodedword.go:295
	_go_fuzz_dep_.CoverTab[35506]++

							if len(header) > 0 {
//line /usr/local/go/src/mime/encodedword.go:297
		_go_fuzz_dep_.CoverTab[35537]++
								buf.WriteString(header)
//line /usr/local/go/src/mime/encodedword.go:298
		// _ = "end of CoverTab[35537]"
	} else {
//line /usr/local/go/src/mime/encodedword.go:299
		_go_fuzz_dep_.CoverTab[35538]++
//line /usr/local/go/src/mime/encodedword.go:299
		// _ = "end of CoverTab[35538]"
//line /usr/local/go/src/mime/encodedword.go:299
	}
//line /usr/local/go/src/mime/encodedword.go:299
	// _ = "end of CoverTab[35506]"
//line /usr/local/go/src/mime/encodedword.go:299
	_go_fuzz_dep_.CoverTab[35507]++

							return buf.String(), nil
//line /usr/local/go/src/mime/encodedword.go:301
	// _ = "end of CoverTab[35507]"
}

func decode(encoding byte, text string) ([]byte, error) {
//line /usr/local/go/src/mime/encodedword.go:304
	_go_fuzz_dep_.CoverTab[35539]++
							switch encoding {
	case 'B', 'b':
//line /usr/local/go/src/mime/encodedword.go:306
		_go_fuzz_dep_.CoverTab[35540]++
								return base64.StdEncoding.DecodeString(text)
//line /usr/local/go/src/mime/encodedword.go:307
		// _ = "end of CoverTab[35540]"
	case 'Q', 'q':
//line /usr/local/go/src/mime/encodedword.go:308
		_go_fuzz_dep_.CoverTab[35541]++
								return qDecode(text)
//line /usr/local/go/src/mime/encodedword.go:309
		// _ = "end of CoverTab[35541]"
	default:
//line /usr/local/go/src/mime/encodedword.go:310
		_go_fuzz_dep_.CoverTab[35542]++
								return nil, errInvalidWord
//line /usr/local/go/src/mime/encodedword.go:311
		// _ = "end of CoverTab[35542]"
	}
//line /usr/local/go/src/mime/encodedword.go:312
	// _ = "end of CoverTab[35539]"
}

func (d *WordDecoder) convert(buf *strings.Builder, charset string, content []byte) error {
//line /usr/local/go/src/mime/encodedword.go:315
	_go_fuzz_dep_.CoverTab[35543]++
							switch {
	case strings.EqualFold("utf-8", charset):
//line /usr/local/go/src/mime/encodedword.go:317
		_go_fuzz_dep_.CoverTab[35545]++
								buf.Write(content)
//line /usr/local/go/src/mime/encodedword.go:318
		// _ = "end of CoverTab[35545]"
	case strings.EqualFold("iso-8859-1", charset):
//line /usr/local/go/src/mime/encodedword.go:319
		_go_fuzz_dep_.CoverTab[35546]++
								for _, c := range content {
//line /usr/local/go/src/mime/encodedword.go:320
			_go_fuzz_dep_.CoverTab[35551]++
									buf.WriteRune(rune(c))
//line /usr/local/go/src/mime/encodedword.go:321
			// _ = "end of CoverTab[35551]"
		}
//line /usr/local/go/src/mime/encodedword.go:322
		// _ = "end of CoverTab[35546]"
	case strings.EqualFold("us-ascii", charset):
//line /usr/local/go/src/mime/encodedword.go:323
		_go_fuzz_dep_.CoverTab[35547]++
								for _, c := range content {
//line /usr/local/go/src/mime/encodedword.go:324
			_go_fuzz_dep_.CoverTab[35552]++
									if c >= utf8.RuneSelf {
//line /usr/local/go/src/mime/encodedword.go:325
				_go_fuzz_dep_.CoverTab[35553]++
										buf.WriteRune(unicode.ReplacementChar)
//line /usr/local/go/src/mime/encodedword.go:326
				// _ = "end of CoverTab[35553]"
			} else {
//line /usr/local/go/src/mime/encodedword.go:327
				_go_fuzz_dep_.CoverTab[35554]++
										buf.WriteByte(c)
//line /usr/local/go/src/mime/encodedword.go:328
				// _ = "end of CoverTab[35554]"
			}
//line /usr/local/go/src/mime/encodedword.go:329
			// _ = "end of CoverTab[35552]"
		}
//line /usr/local/go/src/mime/encodedword.go:330
		// _ = "end of CoverTab[35547]"
	default:
//line /usr/local/go/src/mime/encodedword.go:331
		_go_fuzz_dep_.CoverTab[35548]++
								if d.CharsetReader == nil {
//line /usr/local/go/src/mime/encodedword.go:332
			_go_fuzz_dep_.CoverTab[35555]++
									return fmt.Errorf("mime: unhandled charset %q", charset)
//line /usr/local/go/src/mime/encodedword.go:333
			// _ = "end of CoverTab[35555]"
		} else {
//line /usr/local/go/src/mime/encodedword.go:334
			_go_fuzz_dep_.CoverTab[35556]++
//line /usr/local/go/src/mime/encodedword.go:334
			// _ = "end of CoverTab[35556]"
//line /usr/local/go/src/mime/encodedword.go:334
		}
//line /usr/local/go/src/mime/encodedword.go:334
		// _ = "end of CoverTab[35548]"
//line /usr/local/go/src/mime/encodedword.go:334
		_go_fuzz_dep_.CoverTab[35549]++
								r, err := d.CharsetReader(strings.ToLower(charset), bytes.NewReader(content))
								if err != nil {
//line /usr/local/go/src/mime/encodedword.go:336
			_go_fuzz_dep_.CoverTab[35557]++
									return err
//line /usr/local/go/src/mime/encodedword.go:337
			// _ = "end of CoverTab[35557]"
		} else {
//line /usr/local/go/src/mime/encodedword.go:338
			_go_fuzz_dep_.CoverTab[35558]++
//line /usr/local/go/src/mime/encodedword.go:338
			// _ = "end of CoverTab[35558]"
//line /usr/local/go/src/mime/encodedword.go:338
		}
//line /usr/local/go/src/mime/encodedword.go:338
		// _ = "end of CoverTab[35549]"
//line /usr/local/go/src/mime/encodedword.go:338
		_go_fuzz_dep_.CoverTab[35550]++
								if _, err = io.Copy(buf, r); err != nil {
//line /usr/local/go/src/mime/encodedword.go:339
			_go_fuzz_dep_.CoverTab[35559]++
									return err
//line /usr/local/go/src/mime/encodedword.go:340
			// _ = "end of CoverTab[35559]"
		} else {
//line /usr/local/go/src/mime/encodedword.go:341
			_go_fuzz_dep_.CoverTab[35560]++
//line /usr/local/go/src/mime/encodedword.go:341
			// _ = "end of CoverTab[35560]"
//line /usr/local/go/src/mime/encodedword.go:341
		}
//line /usr/local/go/src/mime/encodedword.go:341
		// _ = "end of CoverTab[35550]"
	}
//line /usr/local/go/src/mime/encodedword.go:342
	// _ = "end of CoverTab[35543]"
//line /usr/local/go/src/mime/encodedword.go:342
	_go_fuzz_dep_.CoverTab[35544]++
							return nil
//line /usr/local/go/src/mime/encodedword.go:343
	// _ = "end of CoverTab[35544]"
}

// hasNonWhitespace reports whether s (assumed to be ASCII) contains at least
//line /usr/local/go/src/mime/encodedword.go:346
// one byte of non-whitespace.
//line /usr/local/go/src/mime/encodedword.go:348
func hasNonWhitespace(s string) bool {
//line /usr/local/go/src/mime/encodedword.go:348
	_go_fuzz_dep_.CoverTab[35561]++
							for _, b := range s {
//line /usr/local/go/src/mime/encodedword.go:349
		_go_fuzz_dep_.CoverTab[35563]++
								switch b {

//line /usr/local/go/src/mime/encodedword.go:353
		case ' ', '\t', '\n', '\r':
//line /usr/local/go/src/mime/encodedword.go:353
			_go_fuzz_dep_.CoverTab[35564]++
//line /usr/local/go/src/mime/encodedword.go:353
			// _ = "end of CoverTab[35564]"
		default:
//line /usr/local/go/src/mime/encodedword.go:354
			_go_fuzz_dep_.CoverTab[35565]++
									return true
//line /usr/local/go/src/mime/encodedword.go:355
			// _ = "end of CoverTab[35565]"
		}
//line /usr/local/go/src/mime/encodedword.go:356
		// _ = "end of CoverTab[35563]"
	}
//line /usr/local/go/src/mime/encodedword.go:357
	// _ = "end of CoverTab[35561]"
//line /usr/local/go/src/mime/encodedword.go:357
	_go_fuzz_dep_.CoverTab[35562]++
							return false
//line /usr/local/go/src/mime/encodedword.go:358
	// _ = "end of CoverTab[35562]"
}

// qDecode decodes a Q encoded string.
func qDecode(s string) ([]byte, error) {
//line /usr/local/go/src/mime/encodedword.go:362
	_go_fuzz_dep_.CoverTab[35566]++
							dec := make([]byte, len(s))
							n := 0
							for i := 0; i < len(s); i++ {
//line /usr/local/go/src/mime/encodedword.go:365
		_go_fuzz_dep_.CoverTab[35568]++
								switch c := s[i]; {
		case c == '_':
//line /usr/local/go/src/mime/encodedword.go:367
			_go_fuzz_dep_.CoverTab[35570]++
									dec[n] = ' '
//line /usr/local/go/src/mime/encodedword.go:368
			// _ = "end of CoverTab[35570]"
		case c == '=':
//line /usr/local/go/src/mime/encodedword.go:369
			_go_fuzz_dep_.CoverTab[35571]++
									if i+2 >= len(s) {
//line /usr/local/go/src/mime/encodedword.go:370
				_go_fuzz_dep_.CoverTab[35576]++
										return nil, errInvalidWord
//line /usr/local/go/src/mime/encodedword.go:371
				// _ = "end of CoverTab[35576]"
			} else {
//line /usr/local/go/src/mime/encodedword.go:372
				_go_fuzz_dep_.CoverTab[35577]++
//line /usr/local/go/src/mime/encodedword.go:372
				// _ = "end of CoverTab[35577]"
//line /usr/local/go/src/mime/encodedword.go:372
			}
//line /usr/local/go/src/mime/encodedword.go:372
			// _ = "end of CoverTab[35571]"
//line /usr/local/go/src/mime/encodedword.go:372
			_go_fuzz_dep_.CoverTab[35572]++
									b, err := readHexByte(s[i+1], s[i+2])
									if err != nil {
//line /usr/local/go/src/mime/encodedword.go:374
				_go_fuzz_dep_.CoverTab[35578]++
										return nil, err
//line /usr/local/go/src/mime/encodedword.go:375
				// _ = "end of CoverTab[35578]"
			} else {
//line /usr/local/go/src/mime/encodedword.go:376
				_go_fuzz_dep_.CoverTab[35579]++
//line /usr/local/go/src/mime/encodedword.go:376
				// _ = "end of CoverTab[35579]"
//line /usr/local/go/src/mime/encodedword.go:376
			}
//line /usr/local/go/src/mime/encodedword.go:376
			// _ = "end of CoverTab[35572]"
//line /usr/local/go/src/mime/encodedword.go:376
			_go_fuzz_dep_.CoverTab[35573]++
									dec[n] = b
									i += 2
//line /usr/local/go/src/mime/encodedword.go:378
			// _ = "end of CoverTab[35573]"
		case (c <= '~' && func() bool {
//line /usr/local/go/src/mime/encodedword.go:379
			_go_fuzz_dep_.CoverTab[35580]++
//line /usr/local/go/src/mime/encodedword.go:379
			return c >= ' '
//line /usr/local/go/src/mime/encodedword.go:379
			// _ = "end of CoverTab[35580]"
//line /usr/local/go/src/mime/encodedword.go:379
		}()) || func() bool {
//line /usr/local/go/src/mime/encodedword.go:379
			_go_fuzz_dep_.CoverTab[35581]++
//line /usr/local/go/src/mime/encodedword.go:379
			return c == '\n'
//line /usr/local/go/src/mime/encodedword.go:379
			// _ = "end of CoverTab[35581]"
//line /usr/local/go/src/mime/encodedword.go:379
		}() || func() bool {
//line /usr/local/go/src/mime/encodedword.go:379
			_go_fuzz_dep_.CoverTab[35582]++
//line /usr/local/go/src/mime/encodedword.go:379
			return c == '\r'
//line /usr/local/go/src/mime/encodedword.go:379
			// _ = "end of CoverTab[35582]"
//line /usr/local/go/src/mime/encodedword.go:379
		}() || func() bool {
//line /usr/local/go/src/mime/encodedword.go:379
			_go_fuzz_dep_.CoverTab[35583]++
//line /usr/local/go/src/mime/encodedword.go:379
			return c == '\t'
//line /usr/local/go/src/mime/encodedword.go:379
			// _ = "end of CoverTab[35583]"
//line /usr/local/go/src/mime/encodedword.go:379
		}():
//line /usr/local/go/src/mime/encodedword.go:379
			_go_fuzz_dep_.CoverTab[35574]++
									dec[n] = c
//line /usr/local/go/src/mime/encodedword.go:380
			// _ = "end of CoverTab[35574]"
		default:
//line /usr/local/go/src/mime/encodedword.go:381
			_go_fuzz_dep_.CoverTab[35575]++
									return nil, errInvalidWord
//line /usr/local/go/src/mime/encodedword.go:382
			// _ = "end of CoverTab[35575]"
		}
//line /usr/local/go/src/mime/encodedword.go:383
		// _ = "end of CoverTab[35568]"
//line /usr/local/go/src/mime/encodedword.go:383
		_go_fuzz_dep_.CoverTab[35569]++
								n++
//line /usr/local/go/src/mime/encodedword.go:384
		// _ = "end of CoverTab[35569]"
	}
//line /usr/local/go/src/mime/encodedword.go:385
	// _ = "end of CoverTab[35566]"
//line /usr/local/go/src/mime/encodedword.go:385
	_go_fuzz_dep_.CoverTab[35567]++

							return dec[:n], nil
//line /usr/local/go/src/mime/encodedword.go:387
	// _ = "end of CoverTab[35567]"
}

// readHexByte returns the byte from its quoted-printable representation.
func readHexByte(a, b byte) (byte, error) {
//line /usr/local/go/src/mime/encodedword.go:391
	_go_fuzz_dep_.CoverTab[35584]++
							var hb, lb byte
							var err error
							if hb, err = fromHex(a); err != nil {
//line /usr/local/go/src/mime/encodedword.go:394
		_go_fuzz_dep_.CoverTab[35587]++
								return 0, err
//line /usr/local/go/src/mime/encodedword.go:395
		// _ = "end of CoverTab[35587]"
	} else {
//line /usr/local/go/src/mime/encodedword.go:396
		_go_fuzz_dep_.CoverTab[35588]++
//line /usr/local/go/src/mime/encodedword.go:396
		// _ = "end of CoverTab[35588]"
//line /usr/local/go/src/mime/encodedword.go:396
	}
//line /usr/local/go/src/mime/encodedword.go:396
	// _ = "end of CoverTab[35584]"
//line /usr/local/go/src/mime/encodedword.go:396
	_go_fuzz_dep_.CoverTab[35585]++
							if lb, err = fromHex(b); err != nil {
//line /usr/local/go/src/mime/encodedword.go:397
		_go_fuzz_dep_.CoverTab[35589]++
								return 0, err
//line /usr/local/go/src/mime/encodedword.go:398
		// _ = "end of CoverTab[35589]"
	} else {
//line /usr/local/go/src/mime/encodedword.go:399
		_go_fuzz_dep_.CoverTab[35590]++
//line /usr/local/go/src/mime/encodedword.go:399
		// _ = "end of CoverTab[35590]"
//line /usr/local/go/src/mime/encodedword.go:399
	}
//line /usr/local/go/src/mime/encodedword.go:399
	// _ = "end of CoverTab[35585]"
//line /usr/local/go/src/mime/encodedword.go:399
	_go_fuzz_dep_.CoverTab[35586]++
							return hb<<4 | lb, nil
//line /usr/local/go/src/mime/encodedword.go:400
	// _ = "end of CoverTab[35586]"
}

func fromHex(b byte) (byte, error) {
//line /usr/local/go/src/mime/encodedword.go:403
	_go_fuzz_dep_.CoverTab[35591]++
							switch {
	case b >= '0' && func() bool {
//line /usr/local/go/src/mime/encodedword.go:405
		_go_fuzz_dep_.CoverTab[35597]++
//line /usr/local/go/src/mime/encodedword.go:405
		return b <= '9'
//line /usr/local/go/src/mime/encodedword.go:405
		// _ = "end of CoverTab[35597]"
//line /usr/local/go/src/mime/encodedword.go:405
	}():
//line /usr/local/go/src/mime/encodedword.go:405
		_go_fuzz_dep_.CoverTab[35593]++
								return b - '0', nil
//line /usr/local/go/src/mime/encodedword.go:406
		// _ = "end of CoverTab[35593]"
	case b >= 'A' && func() bool {
//line /usr/local/go/src/mime/encodedword.go:407
		_go_fuzz_dep_.CoverTab[35598]++
//line /usr/local/go/src/mime/encodedword.go:407
		return b <= 'F'
//line /usr/local/go/src/mime/encodedword.go:407
		// _ = "end of CoverTab[35598]"
//line /usr/local/go/src/mime/encodedword.go:407
	}():
//line /usr/local/go/src/mime/encodedword.go:407
		_go_fuzz_dep_.CoverTab[35594]++
								return b - 'A' + 10, nil
//line /usr/local/go/src/mime/encodedword.go:408
		// _ = "end of CoverTab[35594]"

	case b >= 'a' && func() bool {
//line /usr/local/go/src/mime/encodedword.go:410
		_go_fuzz_dep_.CoverTab[35599]++
//line /usr/local/go/src/mime/encodedword.go:410
		return b <= 'f'
//line /usr/local/go/src/mime/encodedword.go:410
		// _ = "end of CoverTab[35599]"
//line /usr/local/go/src/mime/encodedword.go:410
	}():
//line /usr/local/go/src/mime/encodedword.go:410
		_go_fuzz_dep_.CoverTab[35595]++
								return b - 'a' + 10, nil
//line /usr/local/go/src/mime/encodedword.go:411
		// _ = "end of CoverTab[35595]"
//line /usr/local/go/src/mime/encodedword.go:411
	default:
//line /usr/local/go/src/mime/encodedword.go:411
		_go_fuzz_dep_.CoverTab[35596]++
//line /usr/local/go/src/mime/encodedword.go:411
		// _ = "end of CoverTab[35596]"
	}
//line /usr/local/go/src/mime/encodedword.go:412
	// _ = "end of CoverTab[35591]"
//line /usr/local/go/src/mime/encodedword.go:412
	_go_fuzz_dep_.CoverTab[35592]++
							return 0, fmt.Errorf("mime: invalid hex byte %#02x", b)
//line /usr/local/go/src/mime/encodedword.go:413
	// _ = "end of CoverTab[35592]"
}

//line /usr/local/go/src/mime/encodedword.go:414
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/mime/encodedword.go:414
var _ = _go_fuzz_dep_.CoverTab
