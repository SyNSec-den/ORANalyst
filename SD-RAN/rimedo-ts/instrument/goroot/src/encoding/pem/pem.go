// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/encoding/pem/pem.go:5
// Package pem implements the PEM data encoding, which originated in Privacy
//line /usr/local/go/src/encoding/pem/pem.go:5
// Enhanced Mail. The most common use of PEM encoding today is in TLS keys and
//line /usr/local/go/src/encoding/pem/pem.go:5
// certificates. See RFC 1421.
//line /usr/local/go/src/encoding/pem/pem.go:8
package pem

//line /usr/local/go/src/encoding/pem/pem.go:8
import (
//line /usr/local/go/src/encoding/pem/pem.go:8
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/encoding/pem/pem.go:8
)
//line /usr/local/go/src/encoding/pem/pem.go:8
import (
//line /usr/local/go/src/encoding/pem/pem.go:8
	_atomic_ "sync/atomic"
//line /usr/local/go/src/encoding/pem/pem.go:8
)

import (
	"bytes"
	"encoding/base64"
	"errors"
	"io"
	"sort"
	"strings"
)

// A Block represents a PEM encoded structure.
//line /usr/local/go/src/encoding/pem/pem.go:19
//
//line /usr/local/go/src/encoding/pem/pem.go:19
// The encoded form is:
//line /usr/local/go/src/encoding/pem/pem.go:19
//
//line /usr/local/go/src/encoding/pem/pem.go:19
//	-----BEGIN Type-----
//line /usr/local/go/src/encoding/pem/pem.go:19
//	Headers
//line /usr/local/go/src/encoding/pem/pem.go:19
//	base64-encoded Bytes
//line /usr/local/go/src/encoding/pem/pem.go:19
//	-----END Type-----
//line /usr/local/go/src/encoding/pem/pem.go:19
//
//line /usr/local/go/src/encoding/pem/pem.go:19
// where Headers is a possibly empty sequence of Key: Value lines.
//line /usr/local/go/src/encoding/pem/pem.go:29
type Block struct {
	Type	string			// The type, taken from the preamble (i.e. "RSA PRIVATE KEY").
	Headers	map[string]string	// Optional headers.
	Bytes	[]byte			// The decoded bytes of the contents. Typically a DER encoded ASN.1 structure.
}

// getLine results the first \r\n or \n delineated line from the given byte
//line /usr/local/go/src/encoding/pem/pem.go:35
// array. The line does not include trailing whitespace or the trailing new
//line /usr/local/go/src/encoding/pem/pem.go:35
// line bytes. The remainder of the byte array (also not including the new line
//line /usr/local/go/src/encoding/pem/pem.go:35
// bytes) is also returned and this will always be smaller than the original
//line /usr/local/go/src/encoding/pem/pem.go:35
// argument.
//line /usr/local/go/src/encoding/pem/pem.go:40
func getLine(data []byte) (line, rest []byte) {
//line /usr/local/go/src/encoding/pem/pem.go:40
	_go_fuzz_dep_.CoverTab[10740]++
							i := bytes.IndexByte(data, '\n')
							var j int
							if i < 0 {
//line /usr/local/go/src/encoding/pem/pem.go:43
		_go_fuzz_dep_.CoverTab[10742]++
								i = len(data)
								j = i
//line /usr/local/go/src/encoding/pem/pem.go:45
		// _ = "end of CoverTab[10742]"
	} else {
//line /usr/local/go/src/encoding/pem/pem.go:46
		_go_fuzz_dep_.CoverTab[10743]++
								j = i + 1
								if i > 0 && func() bool {
//line /usr/local/go/src/encoding/pem/pem.go:48
			_go_fuzz_dep_.CoverTab[10744]++
//line /usr/local/go/src/encoding/pem/pem.go:48
			return data[i-1] == '\r'
//line /usr/local/go/src/encoding/pem/pem.go:48
			// _ = "end of CoverTab[10744]"
//line /usr/local/go/src/encoding/pem/pem.go:48
		}() {
//line /usr/local/go/src/encoding/pem/pem.go:48
			_go_fuzz_dep_.CoverTab[10745]++
									i--
//line /usr/local/go/src/encoding/pem/pem.go:49
			// _ = "end of CoverTab[10745]"
		} else {
//line /usr/local/go/src/encoding/pem/pem.go:50
			_go_fuzz_dep_.CoverTab[10746]++
//line /usr/local/go/src/encoding/pem/pem.go:50
			// _ = "end of CoverTab[10746]"
//line /usr/local/go/src/encoding/pem/pem.go:50
		}
//line /usr/local/go/src/encoding/pem/pem.go:50
		// _ = "end of CoverTab[10743]"
	}
//line /usr/local/go/src/encoding/pem/pem.go:51
	// _ = "end of CoverTab[10740]"
//line /usr/local/go/src/encoding/pem/pem.go:51
	_go_fuzz_dep_.CoverTab[10741]++
							return bytes.TrimRight(data[0:i], " \t"), data[j:]
//line /usr/local/go/src/encoding/pem/pem.go:52
	// _ = "end of CoverTab[10741]"
}

// removeSpacesAndTabs returns a copy of its input with all spaces and tabs
//line /usr/local/go/src/encoding/pem/pem.go:55
// removed, if there were any. Otherwise, the input is returned unchanged.
//line /usr/local/go/src/encoding/pem/pem.go:55
//
//line /usr/local/go/src/encoding/pem/pem.go:55
// The base64 decoder already skips newline characters, so we don't need to
//line /usr/local/go/src/encoding/pem/pem.go:55
// filter them out here.
//line /usr/local/go/src/encoding/pem/pem.go:60
func removeSpacesAndTabs(data []byte) []byte {
//line /usr/local/go/src/encoding/pem/pem.go:60
	_go_fuzz_dep_.CoverTab[10747]++
							if !bytes.ContainsAny(data, " \t") {
//line /usr/local/go/src/encoding/pem/pem.go:61
		_go_fuzz_dep_.CoverTab[10750]++

//line /usr/local/go/src/encoding/pem/pem.go:64
		return data
//line /usr/local/go/src/encoding/pem/pem.go:64
		// _ = "end of CoverTab[10750]"
	} else {
//line /usr/local/go/src/encoding/pem/pem.go:65
		_go_fuzz_dep_.CoverTab[10751]++
//line /usr/local/go/src/encoding/pem/pem.go:65
		// _ = "end of CoverTab[10751]"
//line /usr/local/go/src/encoding/pem/pem.go:65
	}
//line /usr/local/go/src/encoding/pem/pem.go:65
	// _ = "end of CoverTab[10747]"
//line /usr/local/go/src/encoding/pem/pem.go:65
	_go_fuzz_dep_.CoverTab[10748]++
							result := make([]byte, len(data))
							n := 0

							for _, b := range data {
//line /usr/local/go/src/encoding/pem/pem.go:69
		_go_fuzz_dep_.CoverTab[10752]++
								if b == ' ' || func() bool {
//line /usr/local/go/src/encoding/pem/pem.go:70
			_go_fuzz_dep_.CoverTab[10754]++
//line /usr/local/go/src/encoding/pem/pem.go:70
			return b == '\t'
//line /usr/local/go/src/encoding/pem/pem.go:70
			// _ = "end of CoverTab[10754]"
//line /usr/local/go/src/encoding/pem/pem.go:70
		}() {
//line /usr/local/go/src/encoding/pem/pem.go:70
			_go_fuzz_dep_.CoverTab[10755]++
									continue
//line /usr/local/go/src/encoding/pem/pem.go:71
			// _ = "end of CoverTab[10755]"
		} else {
//line /usr/local/go/src/encoding/pem/pem.go:72
			_go_fuzz_dep_.CoverTab[10756]++
//line /usr/local/go/src/encoding/pem/pem.go:72
			// _ = "end of CoverTab[10756]"
//line /usr/local/go/src/encoding/pem/pem.go:72
		}
//line /usr/local/go/src/encoding/pem/pem.go:72
		// _ = "end of CoverTab[10752]"
//line /usr/local/go/src/encoding/pem/pem.go:72
		_go_fuzz_dep_.CoverTab[10753]++
								result[n] = b
								n++
//line /usr/local/go/src/encoding/pem/pem.go:74
		// _ = "end of CoverTab[10753]"
	}
//line /usr/local/go/src/encoding/pem/pem.go:75
	// _ = "end of CoverTab[10748]"
//line /usr/local/go/src/encoding/pem/pem.go:75
	_go_fuzz_dep_.CoverTab[10749]++

							return result[0:n]
//line /usr/local/go/src/encoding/pem/pem.go:77
	// _ = "end of CoverTab[10749]"
}

var pemStart = []byte("\n-----BEGIN ")
var pemEnd = []byte("\n-----END ")
var pemEndOfLine = []byte("-----")
var colon = []byte(":")

// Decode will find the next PEM formatted block (certificate, private key
//line /usr/local/go/src/encoding/pem/pem.go:85
// etc) in the input. It returns that block and the remainder of the input. If
//line /usr/local/go/src/encoding/pem/pem.go:85
// no PEM data is found, p is nil and the whole of the input is returned in
//line /usr/local/go/src/encoding/pem/pem.go:85
// rest.
//line /usr/local/go/src/encoding/pem/pem.go:89
func Decode(data []byte) (p *Block, rest []byte) {
//line /usr/local/go/src/encoding/pem/pem.go:89
	_go_fuzz_dep_.CoverTab[10757]++

//line /usr/local/go/src/encoding/pem/pem.go:92
	rest = data
	for {
//line /usr/local/go/src/encoding/pem/pem.go:93
		_go_fuzz_dep_.CoverTab[10758]++
								if bytes.HasPrefix(rest, pemStart[1:]) {
//line /usr/local/go/src/encoding/pem/pem.go:94
			_go_fuzz_dep_.CoverTab[10768]++
									rest = rest[len(pemStart)-1:]
//line /usr/local/go/src/encoding/pem/pem.go:95
			// _ = "end of CoverTab[10768]"
		} else {
//line /usr/local/go/src/encoding/pem/pem.go:96
			_go_fuzz_dep_.CoverTab[10769]++
//line /usr/local/go/src/encoding/pem/pem.go:96
			if _, after, ok := bytes.Cut(rest, pemStart); ok {
//line /usr/local/go/src/encoding/pem/pem.go:96
				_go_fuzz_dep_.CoverTab[10770]++
										rest = after
//line /usr/local/go/src/encoding/pem/pem.go:97
				// _ = "end of CoverTab[10770]"
			} else {
//line /usr/local/go/src/encoding/pem/pem.go:98
				_go_fuzz_dep_.CoverTab[10771]++
										return nil, data
//line /usr/local/go/src/encoding/pem/pem.go:99
				// _ = "end of CoverTab[10771]"
			}
//line /usr/local/go/src/encoding/pem/pem.go:100
			// _ = "end of CoverTab[10769]"
//line /usr/local/go/src/encoding/pem/pem.go:100
		}
//line /usr/local/go/src/encoding/pem/pem.go:100
		// _ = "end of CoverTab[10758]"
//line /usr/local/go/src/encoding/pem/pem.go:100
		_go_fuzz_dep_.CoverTab[10759]++

								var typeLine []byte
								typeLine, rest = getLine(rest)
								if !bytes.HasSuffix(typeLine, pemEndOfLine) {
//line /usr/local/go/src/encoding/pem/pem.go:104
			_go_fuzz_dep_.CoverTab[10772]++
									continue
//line /usr/local/go/src/encoding/pem/pem.go:105
			// _ = "end of CoverTab[10772]"
		} else {
//line /usr/local/go/src/encoding/pem/pem.go:106
			_go_fuzz_dep_.CoverTab[10773]++
//line /usr/local/go/src/encoding/pem/pem.go:106
			// _ = "end of CoverTab[10773]"
//line /usr/local/go/src/encoding/pem/pem.go:106
		}
//line /usr/local/go/src/encoding/pem/pem.go:106
		// _ = "end of CoverTab[10759]"
//line /usr/local/go/src/encoding/pem/pem.go:106
		_go_fuzz_dep_.CoverTab[10760]++
								typeLine = typeLine[0 : len(typeLine)-len(pemEndOfLine)]

								p = &Block{
			Headers:	make(map[string]string),
			Type:		string(typeLine),
		}

		for {
//line /usr/local/go/src/encoding/pem/pem.go:114
			_go_fuzz_dep_.CoverTab[10774]++

//line /usr/local/go/src/encoding/pem/pem.go:117
			if len(rest) == 0 {
//line /usr/local/go/src/encoding/pem/pem.go:117
				_go_fuzz_dep_.CoverTab[10777]++
										return nil, data
//line /usr/local/go/src/encoding/pem/pem.go:118
				// _ = "end of CoverTab[10777]"
			} else {
//line /usr/local/go/src/encoding/pem/pem.go:119
				_go_fuzz_dep_.CoverTab[10778]++
//line /usr/local/go/src/encoding/pem/pem.go:119
				// _ = "end of CoverTab[10778]"
//line /usr/local/go/src/encoding/pem/pem.go:119
			}
//line /usr/local/go/src/encoding/pem/pem.go:119
			// _ = "end of CoverTab[10774]"
//line /usr/local/go/src/encoding/pem/pem.go:119
			_go_fuzz_dep_.CoverTab[10775]++
									line, next := getLine(rest)

									key, val, ok := bytes.Cut(line, colon)
									if !ok {
//line /usr/local/go/src/encoding/pem/pem.go:123
				_go_fuzz_dep_.CoverTab[10779]++
										break
//line /usr/local/go/src/encoding/pem/pem.go:124
				// _ = "end of CoverTab[10779]"
			} else {
//line /usr/local/go/src/encoding/pem/pem.go:125
				_go_fuzz_dep_.CoverTab[10780]++
//line /usr/local/go/src/encoding/pem/pem.go:125
				// _ = "end of CoverTab[10780]"
//line /usr/local/go/src/encoding/pem/pem.go:125
			}
//line /usr/local/go/src/encoding/pem/pem.go:125
			// _ = "end of CoverTab[10775]"
//line /usr/local/go/src/encoding/pem/pem.go:125
			_go_fuzz_dep_.CoverTab[10776]++

//line /usr/local/go/src/encoding/pem/pem.go:128
			key = bytes.TrimSpace(key)
									val = bytes.TrimSpace(val)
									p.Headers[string(key)] = string(val)
									rest = next
//line /usr/local/go/src/encoding/pem/pem.go:131
			// _ = "end of CoverTab[10776]"
		}
//line /usr/local/go/src/encoding/pem/pem.go:132
		// _ = "end of CoverTab[10760]"
//line /usr/local/go/src/encoding/pem/pem.go:132
		_go_fuzz_dep_.CoverTab[10761]++

								var endIndex, endTrailerIndex int

//line /usr/local/go/src/encoding/pem/pem.go:138
		if len(p.Headers) == 0 && func() bool {
//line /usr/local/go/src/encoding/pem/pem.go:138
			_go_fuzz_dep_.CoverTab[10781]++
//line /usr/local/go/src/encoding/pem/pem.go:138
			return bytes.HasPrefix(rest, pemEnd[1:])
//line /usr/local/go/src/encoding/pem/pem.go:138
			// _ = "end of CoverTab[10781]"
//line /usr/local/go/src/encoding/pem/pem.go:138
		}() {
//line /usr/local/go/src/encoding/pem/pem.go:138
			_go_fuzz_dep_.CoverTab[10782]++
									endIndex = 0
									endTrailerIndex = len(pemEnd) - 1
//line /usr/local/go/src/encoding/pem/pem.go:140
			// _ = "end of CoverTab[10782]"
		} else {
//line /usr/local/go/src/encoding/pem/pem.go:141
			_go_fuzz_dep_.CoverTab[10783]++
									endIndex = bytes.Index(rest, pemEnd)
									endTrailerIndex = endIndex + len(pemEnd)
//line /usr/local/go/src/encoding/pem/pem.go:143
			// _ = "end of CoverTab[10783]"
		}
//line /usr/local/go/src/encoding/pem/pem.go:144
		// _ = "end of CoverTab[10761]"
//line /usr/local/go/src/encoding/pem/pem.go:144
		_go_fuzz_dep_.CoverTab[10762]++

								if endIndex < 0 {
//line /usr/local/go/src/encoding/pem/pem.go:146
			_go_fuzz_dep_.CoverTab[10784]++
									continue
//line /usr/local/go/src/encoding/pem/pem.go:147
			// _ = "end of CoverTab[10784]"
		} else {
//line /usr/local/go/src/encoding/pem/pem.go:148
			_go_fuzz_dep_.CoverTab[10785]++
//line /usr/local/go/src/encoding/pem/pem.go:148
			// _ = "end of CoverTab[10785]"
//line /usr/local/go/src/encoding/pem/pem.go:148
		}
//line /usr/local/go/src/encoding/pem/pem.go:148
		// _ = "end of CoverTab[10762]"
//line /usr/local/go/src/encoding/pem/pem.go:148
		_go_fuzz_dep_.CoverTab[10763]++

//line /usr/local/go/src/encoding/pem/pem.go:152
		endTrailer := rest[endTrailerIndex:]
		endTrailerLen := len(typeLine) + len(pemEndOfLine)
		if len(endTrailer) < endTrailerLen {
//line /usr/local/go/src/encoding/pem/pem.go:154
			_go_fuzz_dep_.CoverTab[10786]++
									continue
//line /usr/local/go/src/encoding/pem/pem.go:155
			// _ = "end of CoverTab[10786]"
		} else {
//line /usr/local/go/src/encoding/pem/pem.go:156
			_go_fuzz_dep_.CoverTab[10787]++
//line /usr/local/go/src/encoding/pem/pem.go:156
			// _ = "end of CoverTab[10787]"
//line /usr/local/go/src/encoding/pem/pem.go:156
		}
//line /usr/local/go/src/encoding/pem/pem.go:156
		// _ = "end of CoverTab[10763]"
//line /usr/local/go/src/encoding/pem/pem.go:156
		_go_fuzz_dep_.CoverTab[10764]++

								restOfEndLine := endTrailer[endTrailerLen:]
								endTrailer = endTrailer[:endTrailerLen]
								if !bytes.HasPrefix(endTrailer, typeLine) || func() bool {
//line /usr/local/go/src/encoding/pem/pem.go:160
			_go_fuzz_dep_.CoverTab[10788]++
//line /usr/local/go/src/encoding/pem/pem.go:160
			return !bytes.HasSuffix(endTrailer, pemEndOfLine)
									// _ = "end of CoverTab[10788]"
//line /usr/local/go/src/encoding/pem/pem.go:161
		}() {
//line /usr/local/go/src/encoding/pem/pem.go:161
			_go_fuzz_dep_.CoverTab[10789]++
									continue
//line /usr/local/go/src/encoding/pem/pem.go:162
			// _ = "end of CoverTab[10789]"
		} else {
//line /usr/local/go/src/encoding/pem/pem.go:163
			_go_fuzz_dep_.CoverTab[10790]++
//line /usr/local/go/src/encoding/pem/pem.go:163
			// _ = "end of CoverTab[10790]"
//line /usr/local/go/src/encoding/pem/pem.go:163
		}
//line /usr/local/go/src/encoding/pem/pem.go:163
		// _ = "end of CoverTab[10764]"
//line /usr/local/go/src/encoding/pem/pem.go:163
		_go_fuzz_dep_.CoverTab[10765]++

//line /usr/local/go/src/encoding/pem/pem.go:166
		if s, _ := getLine(restOfEndLine); len(s) != 0 {
//line /usr/local/go/src/encoding/pem/pem.go:166
			_go_fuzz_dep_.CoverTab[10791]++
									continue
//line /usr/local/go/src/encoding/pem/pem.go:167
			// _ = "end of CoverTab[10791]"
		} else {
//line /usr/local/go/src/encoding/pem/pem.go:168
			_go_fuzz_dep_.CoverTab[10792]++
//line /usr/local/go/src/encoding/pem/pem.go:168
			// _ = "end of CoverTab[10792]"
//line /usr/local/go/src/encoding/pem/pem.go:168
		}
//line /usr/local/go/src/encoding/pem/pem.go:168
		// _ = "end of CoverTab[10765]"
//line /usr/local/go/src/encoding/pem/pem.go:168
		_go_fuzz_dep_.CoverTab[10766]++

								base64Data := removeSpacesAndTabs(rest[:endIndex])
								p.Bytes = make([]byte, base64.StdEncoding.DecodedLen(len(base64Data)))
								n, err := base64.StdEncoding.Decode(p.Bytes, base64Data)
								if err != nil {
//line /usr/local/go/src/encoding/pem/pem.go:173
			_go_fuzz_dep_.CoverTab[10793]++
									continue
//line /usr/local/go/src/encoding/pem/pem.go:174
			// _ = "end of CoverTab[10793]"
		} else {
//line /usr/local/go/src/encoding/pem/pem.go:175
			_go_fuzz_dep_.CoverTab[10794]++
//line /usr/local/go/src/encoding/pem/pem.go:175
			// _ = "end of CoverTab[10794]"
//line /usr/local/go/src/encoding/pem/pem.go:175
		}
//line /usr/local/go/src/encoding/pem/pem.go:175
		// _ = "end of CoverTab[10766]"
//line /usr/local/go/src/encoding/pem/pem.go:175
		_go_fuzz_dep_.CoverTab[10767]++
								p.Bytes = p.Bytes[:n]

//line /usr/local/go/src/encoding/pem/pem.go:180
		_, rest = getLine(rest[endIndex+len(pemEnd)-1:])
								return p, rest
//line /usr/local/go/src/encoding/pem/pem.go:181
		// _ = "end of CoverTab[10767]"
	}
//line /usr/local/go/src/encoding/pem/pem.go:182
	// _ = "end of CoverTab[10757]"
}

const pemLineLength = 64

type lineBreaker struct {
	line	[pemLineLength]byte
	used	int
	out	io.Writer
}

var nl = []byte{'\n'}

func (l *lineBreaker) Write(b []byte) (n int, err error) {
//line /usr/local/go/src/encoding/pem/pem.go:195
	_go_fuzz_dep_.CoverTab[10795]++
							if l.used+len(b) < pemLineLength {
//line /usr/local/go/src/encoding/pem/pem.go:196
		_go_fuzz_dep_.CoverTab[10800]++
								copy(l.line[l.used:], b)
								l.used += len(b)
								return len(b), nil
//line /usr/local/go/src/encoding/pem/pem.go:199
		// _ = "end of CoverTab[10800]"
	} else {
//line /usr/local/go/src/encoding/pem/pem.go:200
		_go_fuzz_dep_.CoverTab[10801]++
//line /usr/local/go/src/encoding/pem/pem.go:200
		// _ = "end of CoverTab[10801]"
//line /usr/local/go/src/encoding/pem/pem.go:200
	}
//line /usr/local/go/src/encoding/pem/pem.go:200
	// _ = "end of CoverTab[10795]"
//line /usr/local/go/src/encoding/pem/pem.go:200
	_go_fuzz_dep_.CoverTab[10796]++

							n, err = l.out.Write(l.line[0:l.used])
							if err != nil {
//line /usr/local/go/src/encoding/pem/pem.go:203
		_go_fuzz_dep_.CoverTab[10802]++
								return
//line /usr/local/go/src/encoding/pem/pem.go:204
		// _ = "end of CoverTab[10802]"
	} else {
//line /usr/local/go/src/encoding/pem/pem.go:205
		_go_fuzz_dep_.CoverTab[10803]++
//line /usr/local/go/src/encoding/pem/pem.go:205
		// _ = "end of CoverTab[10803]"
//line /usr/local/go/src/encoding/pem/pem.go:205
	}
//line /usr/local/go/src/encoding/pem/pem.go:205
	// _ = "end of CoverTab[10796]"
//line /usr/local/go/src/encoding/pem/pem.go:205
	_go_fuzz_dep_.CoverTab[10797]++
							excess := pemLineLength - l.used
							l.used = 0

							n, err = l.out.Write(b[0:excess])
							if err != nil {
//line /usr/local/go/src/encoding/pem/pem.go:210
		_go_fuzz_dep_.CoverTab[10804]++
								return
//line /usr/local/go/src/encoding/pem/pem.go:211
		// _ = "end of CoverTab[10804]"
	} else {
//line /usr/local/go/src/encoding/pem/pem.go:212
		_go_fuzz_dep_.CoverTab[10805]++
//line /usr/local/go/src/encoding/pem/pem.go:212
		// _ = "end of CoverTab[10805]"
//line /usr/local/go/src/encoding/pem/pem.go:212
	}
//line /usr/local/go/src/encoding/pem/pem.go:212
	// _ = "end of CoverTab[10797]"
//line /usr/local/go/src/encoding/pem/pem.go:212
	_go_fuzz_dep_.CoverTab[10798]++

							n, err = l.out.Write(nl)
							if err != nil {
//line /usr/local/go/src/encoding/pem/pem.go:215
		_go_fuzz_dep_.CoverTab[10806]++
								return
//line /usr/local/go/src/encoding/pem/pem.go:216
		// _ = "end of CoverTab[10806]"
	} else {
//line /usr/local/go/src/encoding/pem/pem.go:217
		_go_fuzz_dep_.CoverTab[10807]++
//line /usr/local/go/src/encoding/pem/pem.go:217
		// _ = "end of CoverTab[10807]"
//line /usr/local/go/src/encoding/pem/pem.go:217
	}
//line /usr/local/go/src/encoding/pem/pem.go:217
	// _ = "end of CoverTab[10798]"
//line /usr/local/go/src/encoding/pem/pem.go:217
	_go_fuzz_dep_.CoverTab[10799]++

							return l.Write(b[excess:])
//line /usr/local/go/src/encoding/pem/pem.go:219
	// _ = "end of CoverTab[10799]"
}

func (l *lineBreaker) Close() (err error) {
//line /usr/local/go/src/encoding/pem/pem.go:222
	_go_fuzz_dep_.CoverTab[10808]++
							if l.used > 0 {
//line /usr/local/go/src/encoding/pem/pem.go:223
		_go_fuzz_dep_.CoverTab[10810]++
								_, err = l.out.Write(l.line[0:l.used])
								if err != nil {
//line /usr/local/go/src/encoding/pem/pem.go:225
			_go_fuzz_dep_.CoverTab[10812]++
									return
//line /usr/local/go/src/encoding/pem/pem.go:226
			// _ = "end of CoverTab[10812]"
		} else {
//line /usr/local/go/src/encoding/pem/pem.go:227
			_go_fuzz_dep_.CoverTab[10813]++
//line /usr/local/go/src/encoding/pem/pem.go:227
			// _ = "end of CoverTab[10813]"
//line /usr/local/go/src/encoding/pem/pem.go:227
		}
//line /usr/local/go/src/encoding/pem/pem.go:227
		// _ = "end of CoverTab[10810]"
//line /usr/local/go/src/encoding/pem/pem.go:227
		_go_fuzz_dep_.CoverTab[10811]++
								_, err = l.out.Write(nl)
//line /usr/local/go/src/encoding/pem/pem.go:228
		// _ = "end of CoverTab[10811]"
	} else {
//line /usr/local/go/src/encoding/pem/pem.go:229
		_go_fuzz_dep_.CoverTab[10814]++
//line /usr/local/go/src/encoding/pem/pem.go:229
		// _ = "end of CoverTab[10814]"
//line /usr/local/go/src/encoding/pem/pem.go:229
	}
//line /usr/local/go/src/encoding/pem/pem.go:229
	// _ = "end of CoverTab[10808]"
//line /usr/local/go/src/encoding/pem/pem.go:229
	_go_fuzz_dep_.CoverTab[10809]++

							return
//line /usr/local/go/src/encoding/pem/pem.go:231
	// _ = "end of CoverTab[10809]"
}

func writeHeader(out io.Writer, k, v string) error {
//line /usr/local/go/src/encoding/pem/pem.go:234
	_go_fuzz_dep_.CoverTab[10815]++
							_, err := out.Write([]byte(k + ": " + v + "\n"))
							return err
//line /usr/local/go/src/encoding/pem/pem.go:236
	// _ = "end of CoverTab[10815]"
}

// Encode writes the PEM encoding of b to out.
func Encode(out io.Writer, b *Block) error {
//line /usr/local/go/src/encoding/pem/pem.go:240
	_go_fuzz_dep_.CoverTab[10816]++

							for k := range b.Headers {
//line /usr/local/go/src/encoding/pem/pem.go:242
		_go_fuzz_dep_.CoverTab[10823]++
								if strings.Contains(k, ":") {
//line /usr/local/go/src/encoding/pem/pem.go:243
			_go_fuzz_dep_.CoverTab[10824]++
									return errors.New("pem: cannot encode a header key that contains a colon")
//line /usr/local/go/src/encoding/pem/pem.go:244
			// _ = "end of CoverTab[10824]"
		} else {
//line /usr/local/go/src/encoding/pem/pem.go:245
			_go_fuzz_dep_.CoverTab[10825]++
//line /usr/local/go/src/encoding/pem/pem.go:245
			// _ = "end of CoverTab[10825]"
//line /usr/local/go/src/encoding/pem/pem.go:245
		}
//line /usr/local/go/src/encoding/pem/pem.go:245
		// _ = "end of CoverTab[10823]"
	}
//line /usr/local/go/src/encoding/pem/pem.go:246
	// _ = "end of CoverTab[10816]"
//line /usr/local/go/src/encoding/pem/pem.go:246
	_go_fuzz_dep_.CoverTab[10817]++

//line /usr/local/go/src/encoding/pem/pem.go:251
	if _, err := out.Write(pemStart[1:]); err != nil {
//line /usr/local/go/src/encoding/pem/pem.go:251
		_go_fuzz_dep_.CoverTab[10826]++
								return err
//line /usr/local/go/src/encoding/pem/pem.go:252
		// _ = "end of CoverTab[10826]"
	} else {
//line /usr/local/go/src/encoding/pem/pem.go:253
		_go_fuzz_dep_.CoverTab[10827]++
//line /usr/local/go/src/encoding/pem/pem.go:253
		// _ = "end of CoverTab[10827]"
//line /usr/local/go/src/encoding/pem/pem.go:253
	}
//line /usr/local/go/src/encoding/pem/pem.go:253
	// _ = "end of CoverTab[10817]"
//line /usr/local/go/src/encoding/pem/pem.go:253
	_go_fuzz_dep_.CoverTab[10818]++
							if _, err := out.Write([]byte(b.Type + "-----\n")); err != nil {
//line /usr/local/go/src/encoding/pem/pem.go:254
		_go_fuzz_dep_.CoverTab[10828]++
								return err
//line /usr/local/go/src/encoding/pem/pem.go:255
		// _ = "end of CoverTab[10828]"
	} else {
//line /usr/local/go/src/encoding/pem/pem.go:256
		_go_fuzz_dep_.CoverTab[10829]++
//line /usr/local/go/src/encoding/pem/pem.go:256
		// _ = "end of CoverTab[10829]"
//line /usr/local/go/src/encoding/pem/pem.go:256
	}
//line /usr/local/go/src/encoding/pem/pem.go:256
	// _ = "end of CoverTab[10818]"
//line /usr/local/go/src/encoding/pem/pem.go:256
	_go_fuzz_dep_.CoverTab[10819]++

							if len(b.Headers) > 0 {
//line /usr/local/go/src/encoding/pem/pem.go:258
		_go_fuzz_dep_.CoverTab[10830]++
								const procType = "Proc-Type"
								h := make([]string, 0, len(b.Headers))
								hasProcType := false
								for k := range b.Headers {
//line /usr/local/go/src/encoding/pem/pem.go:262
			_go_fuzz_dep_.CoverTab[10834]++
									if k == procType {
//line /usr/local/go/src/encoding/pem/pem.go:263
				_go_fuzz_dep_.CoverTab[10836]++
										hasProcType = true
										continue
//line /usr/local/go/src/encoding/pem/pem.go:265
				// _ = "end of CoverTab[10836]"
			} else {
//line /usr/local/go/src/encoding/pem/pem.go:266
				_go_fuzz_dep_.CoverTab[10837]++
//line /usr/local/go/src/encoding/pem/pem.go:266
				// _ = "end of CoverTab[10837]"
//line /usr/local/go/src/encoding/pem/pem.go:266
			}
//line /usr/local/go/src/encoding/pem/pem.go:266
			// _ = "end of CoverTab[10834]"
//line /usr/local/go/src/encoding/pem/pem.go:266
			_go_fuzz_dep_.CoverTab[10835]++
									h = append(h, k)
//line /usr/local/go/src/encoding/pem/pem.go:267
			// _ = "end of CoverTab[10835]"
		}
//line /usr/local/go/src/encoding/pem/pem.go:268
		// _ = "end of CoverTab[10830]"
//line /usr/local/go/src/encoding/pem/pem.go:268
		_go_fuzz_dep_.CoverTab[10831]++

//line /usr/local/go/src/encoding/pem/pem.go:271
		if hasProcType {
//line /usr/local/go/src/encoding/pem/pem.go:271
			_go_fuzz_dep_.CoverTab[10838]++
									if err := writeHeader(out, procType, b.Headers[procType]); err != nil {
//line /usr/local/go/src/encoding/pem/pem.go:272
				_go_fuzz_dep_.CoverTab[10839]++
										return err
//line /usr/local/go/src/encoding/pem/pem.go:273
				// _ = "end of CoverTab[10839]"
			} else {
//line /usr/local/go/src/encoding/pem/pem.go:274
				_go_fuzz_dep_.CoverTab[10840]++
//line /usr/local/go/src/encoding/pem/pem.go:274
				// _ = "end of CoverTab[10840]"
//line /usr/local/go/src/encoding/pem/pem.go:274
			}
//line /usr/local/go/src/encoding/pem/pem.go:274
			// _ = "end of CoverTab[10838]"
		} else {
//line /usr/local/go/src/encoding/pem/pem.go:275
			_go_fuzz_dep_.CoverTab[10841]++
//line /usr/local/go/src/encoding/pem/pem.go:275
			// _ = "end of CoverTab[10841]"
//line /usr/local/go/src/encoding/pem/pem.go:275
		}
//line /usr/local/go/src/encoding/pem/pem.go:275
		// _ = "end of CoverTab[10831]"
//line /usr/local/go/src/encoding/pem/pem.go:275
		_go_fuzz_dep_.CoverTab[10832]++

								sort.Strings(h)
								for _, k := range h {
//line /usr/local/go/src/encoding/pem/pem.go:278
			_go_fuzz_dep_.CoverTab[10842]++
									if err := writeHeader(out, k, b.Headers[k]); err != nil {
//line /usr/local/go/src/encoding/pem/pem.go:279
				_go_fuzz_dep_.CoverTab[10843]++
										return err
//line /usr/local/go/src/encoding/pem/pem.go:280
				// _ = "end of CoverTab[10843]"
			} else {
//line /usr/local/go/src/encoding/pem/pem.go:281
				_go_fuzz_dep_.CoverTab[10844]++
//line /usr/local/go/src/encoding/pem/pem.go:281
				// _ = "end of CoverTab[10844]"
//line /usr/local/go/src/encoding/pem/pem.go:281
			}
//line /usr/local/go/src/encoding/pem/pem.go:281
			// _ = "end of CoverTab[10842]"
		}
//line /usr/local/go/src/encoding/pem/pem.go:282
		// _ = "end of CoverTab[10832]"
//line /usr/local/go/src/encoding/pem/pem.go:282
		_go_fuzz_dep_.CoverTab[10833]++
								if _, err := out.Write(nl); err != nil {
//line /usr/local/go/src/encoding/pem/pem.go:283
			_go_fuzz_dep_.CoverTab[10845]++
									return err
//line /usr/local/go/src/encoding/pem/pem.go:284
			// _ = "end of CoverTab[10845]"
		} else {
//line /usr/local/go/src/encoding/pem/pem.go:285
			_go_fuzz_dep_.CoverTab[10846]++
//line /usr/local/go/src/encoding/pem/pem.go:285
			// _ = "end of CoverTab[10846]"
//line /usr/local/go/src/encoding/pem/pem.go:285
		}
//line /usr/local/go/src/encoding/pem/pem.go:285
		// _ = "end of CoverTab[10833]"
	} else {
//line /usr/local/go/src/encoding/pem/pem.go:286
		_go_fuzz_dep_.CoverTab[10847]++
//line /usr/local/go/src/encoding/pem/pem.go:286
		// _ = "end of CoverTab[10847]"
//line /usr/local/go/src/encoding/pem/pem.go:286
	}
//line /usr/local/go/src/encoding/pem/pem.go:286
	// _ = "end of CoverTab[10819]"
//line /usr/local/go/src/encoding/pem/pem.go:286
	_go_fuzz_dep_.CoverTab[10820]++

							var breaker lineBreaker
							breaker.out = out

							b64 := base64.NewEncoder(base64.StdEncoding, &breaker)
							if _, err := b64.Write(b.Bytes); err != nil {
//line /usr/local/go/src/encoding/pem/pem.go:292
		_go_fuzz_dep_.CoverTab[10848]++
								return err
//line /usr/local/go/src/encoding/pem/pem.go:293
		// _ = "end of CoverTab[10848]"
	} else {
//line /usr/local/go/src/encoding/pem/pem.go:294
		_go_fuzz_dep_.CoverTab[10849]++
//line /usr/local/go/src/encoding/pem/pem.go:294
		// _ = "end of CoverTab[10849]"
//line /usr/local/go/src/encoding/pem/pem.go:294
	}
//line /usr/local/go/src/encoding/pem/pem.go:294
	// _ = "end of CoverTab[10820]"
//line /usr/local/go/src/encoding/pem/pem.go:294
	_go_fuzz_dep_.CoverTab[10821]++
							b64.Close()
							breaker.Close()

							if _, err := out.Write(pemEnd[1:]); err != nil {
//line /usr/local/go/src/encoding/pem/pem.go:298
		_go_fuzz_dep_.CoverTab[10850]++
								return err
//line /usr/local/go/src/encoding/pem/pem.go:299
		// _ = "end of CoverTab[10850]"
	} else {
//line /usr/local/go/src/encoding/pem/pem.go:300
		_go_fuzz_dep_.CoverTab[10851]++
//line /usr/local/go/src/encoding/pem/pem.go:300
		// _ = "end of CoverTab[10851]"
//line /usr/local/go/src/encoding/pem/pem.go:300
	}
//line /usr/local/go/src/encoding/pem/pem.go:300
	// _ = "end of CoverTab[10821]"
//line /usr/local/go/src/encoding/pem/pem.go:300
	_go_fuzz_dep_.CoverTab[10822]++
							_, err := out.Write([]byte(b.Type + "-----\n"))
							return err
//line /usr/local/go/src/encoding/pem/pem.go:302
	// _ = "end of CoverTab[10822]"
}

// EncodeToMemory returns the PEM encoding of b.
//line /usr/local/go/src/encoding/pem/pem.go:305
//
//line /usr/local/go/src/encoding/pem/pem.go:305
// If b has invalid headers and cannot be encoded,
//line /usr/local/go/src/encoding/pem/pem.go:305
// EncodeToMemory returns nil. If it is important to
//line /usr/local/go/src/encoding/pem/pem.go:305
// report details about this error case, use Encode instead.
//line /usr/local/go/src/encoding/pem/pem.go:310
func EncodeToMemory(b *Block) []byte {
//line /usr/local/go/src/encoding/pem/pem.go:310
	_go_fuzz_dep_.CoverTab[10852]++
							var buf bytes.Buffer
							if err := Encode(&buf, b); err != nil {
//line /usr/local/go/src/encoding/pem/pem.go:312
		_go_fuzz_dep_.CoverTab[10854]++
								return nil
//line /usr/local/go/src/encoding/pem/pem.go:313
		// _ = "end of CoverTab[10854]"
	} else {
//line /usr/local/go/src/encoding/pem/pem.go:314
		_go_fuzz_dep_.CoverTab[10855]++
//line /usr/local/go/src/encoding/pem/pem.go:314
		// _ = "end of CoverTab[10855]"
//line /usr/local/go/src/encoding/pem/pem.go:314
	}
//line /usr/local/go/src/encoding/pem/pem.go:314
	// _ = "end of CoverTab[10852]"
//line /usr/local/go/src/encoding/pem/pem.go:314
	_go_fuzz_dep_.CoverTab[10853]++
							return buf.Bytes()
//line /usr/local/go/src/encoding/pem/pem.go:315
	// _ = "end of CoverTab[10853]"
}

//line /usr/local/go/src/encoding/pem/pem.go:316
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/encoding/pem/pem.go:316
var _ = _go_fuzz_dep_.CoverTab
