// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:5
package text

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:5
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:5
)

import (
	"math"
	"math/bits"
	"strconv"
	"strings"
	"unicode/utf8"

	"google.golang.org/protobuf/internal/detrand"
	"google.golang.org/protobuf/internal/errors"
)

// encType represents an encoding type.
type encType uint8

const (
	_	encType	= (1 << iota) / 2
	name
	scalar
	messageOpen
	messageClose
)

// Encoder provides methods to write out textproto constructs and values. The user is
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:29
// responsible for producing valid sequences of constructs and values.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:31
type Encoder struct {
	encoderState

	indent		string
	delims		[2]byte
	outputASCII	bool
}

type encoderState struct {
	lastType	encType
	indents		[]byte
	out		[]byte
}

// NewEncoder returns an Encoder.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:45
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:45
// If indent is a non-empty string, it causes every entry in a List or Message
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:45
// to be preceded by the indent and trailed by a newline.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:45
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:45
// If delims is not the zero value, it controls the delimiter characters used
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:45
// for messages (e.g., "{}" vs "<>").
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:45
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:45
// If outputASCII is true, strings will be serialized in such a way that
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:45
// multi-byte UTF-8 sequences are escaped. This property ensures that the
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:45
// overall output is ASCII (as opposed to UTF-8).
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:56
func NewEncoder(indent string, delims [2]byte, outputASCII bool) (*Encoder, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:56
	_go_fuzz_dep_.CoverTab[50031]++
														e := &Encoder{}
														if len(indent) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:58
		_go_fuzz_dep_.CoverTab[50034]++
															if strings.Trim(indent, " \t") != "" {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:59
			_go_fuzz_dep_.CoverTab[50036]++
																return nil, errors.New("indent may only be composed of space and tab characters")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:60
			// _ = "end of CoverTab[50036]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:61
			_go_fuzz_dep_.CoverTab[50037]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:61
			// _ = "end of CoverTab[50037]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:61
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:61
		// _ = "end of CoverTab[50034]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:61
		_go_fuzz_dep_.CoverTab[50035]++
															e.indent = indent
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:62
		// _ = "end of CoverTab[50035]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:63
		_go_fuzz_dep_.CoverTab[50038]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:63
		// _ = "end of CoverTab[50038]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:63
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:63
	// _ = "end of CoverTab[50031]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:63
	_go_fuzz_dep_.CoverTab[50032]++
														switch delims {
	case [2]byte{0, 0}:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:65
		_go_fuzz_dep_.CoverTab[50039]++
															e.delims = [2]byte{'{', '}'}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:66
		// _ = "end of CoverTab[50039]"
	case [2]byte{'{', '}'}, [2]byte{'<', '>'}:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:67
		_go_fuzz_dep_.CoverTab[50040]++
															e.delims = delims
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:68
		// _ = "end of CoverTab[50040]"
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:69
		_go_fuzz_dep_.CoverTab[50041]++
															return nil, errors.New("delimiters may only be \"{}\" or \"<>\"")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:70
		// _ = "end of CoverTab[50041]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:71
	// _ = "end of CoverTab[50032]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:71
	_go_fuzz_dep_.CoverTab[50033]++
														e.outputASCII = outputASCII

														return e, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:74
	// _ = "end of CoverTab[50033]"
}

// Bytes returns the content of the written bytes.
func (e *Encoder) Bytes() []byte {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:78
	_go_fuzz_dep_.CoverTab[50042]++
														return e.out
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:79
	// _ = "end of CoverTab[50042]"
}

// StartMessage writes out the '{' or '<' symbol.
func (e *Encoder) StartMessage() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:83
	_go_fuzz_dep_.CoverTab[50043]++
														e.prepareNext(messageOpen)
														e.out = append(e.out, e.delims[0])
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:85
	// _ = "end of CoverTab[50043]"
}

// EndMessage writes out the '}' or '>' symbol.
func (e *Encoder) EndMessage() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:89
	_go_fuzz_dep_.CoverTab[50044]++
														e.prepareNext(messageClose)
														e.out = append(e.out, e.delims[1])
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:91
	// _ = "end of CoverTab[50044]"
}

// WriteName writes out the field name and the separator ':'.
func (e *Encoder) WriteName(s string) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:95
	_go_fuzz_dep_.CoverTab[50045]++
														e.prepareNext(name)
														e.out = append(e.out, s...)
														e.out = append(e.out, ':')
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:98
	// _ = "end of CoverTab[50045]"
}

// WriteBool writes out the given boolean value.
func (e *Encoder) WriteBool(b bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:102
	_go_fuzz_dep_.CoverTab[50046]++
														if b {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:103
		_go_fuzz_dep_.CoverTab[50047]++
															e.WriteLiteral("true")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:104
		// _ = "end of CoverTab[50047]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:105
		_go_fuzz_dep_.CoverTab[50048]++
															e.WriteLiteral("false")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:106
		// _ = "end of CoverTab[50048]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:107
	// _ = "end of CoverTab[50046]"
}

// WriteString writes out the given string value.
func (e *Encoder) WriteString(s string) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:111
	_go_fuzz_dep_.CoverTab[50049]++
														e.prepareNext(scalar)
														e.out = appendString(e.out, s, e.outputASCII)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:113
	// _ = "end of CoverTab[50049]"
}

func appendString(out []byte, in string, outputASCII bool) []byte {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:116
	_go_fuzz_dep_.CoverTab[50050]++
														out = append(out, '"')
														i := indexNeedEscapeInString(in)
														in, out = in[i:], append(out, in[:i]...)
														for len(in) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:120
		_go_fuzz_dep_.CoverTab[50052]++
															switch r, n := utf8.DecodeRuneInString(in); {
		case r == utf8.RuneError && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:122
			_go_fuzz_dep_.CoverTab[50059]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:122
			return n == 1
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:122
			// _ = "end of CoverTab[50059]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:122
		}():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:122
			_go_fuzz_dep_.CoverTab[50053]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:125
			r = rune(in[0])
																fallthrough
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:126
			// _ = "end of CoverTab[50053]"
		case r < ' ' || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:127
			_go_fuzz_dep_.CoverTab[50060]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:127
			return r == '"'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:127
			// _ = "end of CoverTab[50060]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:127
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:127
			_go_fuzz_dep_.CoverTab[50061]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:127
			return r == '\\'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:127
			// _ = "end of CoverTab[50061]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:127
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:127
			_go_fuzz_dep_.CoverTab[50062]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:127
			return r == 0x7f
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:127
			// _ = "end of CoverTab[50062]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:127
		}():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:127
			_go_fuzz_dep_.CoverTab[50054]++
																out = append(out, '\\')
																switch r {
			case '"', '\\':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:130
				_go_fuzz_dep_.CoverTab[50063]++
																	out = append(out, byte(r))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:131
				// _ = "end of CoverTab[50063]"
			case '\n':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:132
				_go_fuzz_dep_.CoverTab[50064]++
																	out = append(out, 'n')
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:133
				// _ = "end of CoverTab[50064]"
			case '\r':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:134
				_go_fuzz_dep_.CoverTab[50065]++
																	out = append(out, 'r')
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:135
				// _ = "end of CoverTab[50065]"
			case '\t':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:136
				_go_fuzz_dep_.CoverTab[50066]++
																	out = append(out, 't')
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:137
				// _ = "end of CoverTab[50066]"
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:138
				_go_fuzz_dep_.CoverTab[50067]++
																	out = append(out, 'x')
																	out = append(out, "00"[1+(bits.Len32(uint32(r))-1)/4:]...)
																	out = strconv.AppendUint(out, uint64(r), 16)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:141
				// _ = "end of CoverTab[50067]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:142
			// _ = "end of CoverTab[50054]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:142
			_go_fuzz_dep_.CoverTab[50055]++
																in = in[n:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:143
			// _ = "end of CoverTab[50055]"
		case r >= utf8.RuneSelf && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:144
			_go_fuzz_dep_.CoverTab[50068]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:144
			return (outputASCII || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:144
				_go_fuzz_dep_.CoverTab[50069]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:144
				return r <= 0x009f
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:144
				// _ = "end of CoverTab[50069]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:144
			}())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:144
			// _ = "end of CoverTab[50068]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:144
		}():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:144
			_go_fuzz_dep_.CoverTab[50056]++
																out = append(out, '\\')
																if r <= math.MaxUint16 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:146
				_go_fuzz_dep_.CoverTab[50070]++
																	out = append(out, 'u')
																	out = append(out, "0000"[1+(bits.Len32(uint32(r))-1)/4:]...)
																	out = strconv.AppendUint(out, uint64(r), 16)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:149
				// _ = "end of CoverTab[50070]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:150
				_go_fuzz_dep_.CoverTab[50071]++
																	out = append(out, 'U')
																	out = append(out, "00000000"[1+(bits.Len32(uint32(r))-1)/4:]...)
																	out = strconv.AppendUint(out, uint64(r), 16)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:153
				// _ = "end of CoverTab[50071]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:154
			// _ = "end of CoverTab[50056]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:154
			_go_fuzz_dep_.CoverTab[50057]++
																in = in[n:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:155
			// _ = "end of CoverTab[50057]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:156
			_go_fuzz_dep_.CoverTab[50058]++
																i := indexNeedEscapeInString(in[n:])
																in, out = in[n+i:], append(out, in[:n+i]...)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:158
			// _ = "end of CoverTab[50058]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:159
		// _ = "end of CoverTab[50052]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:160
	// _ = "end of CoverTab[50050]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:160
	_go_fuzz_dep_.CoverTab[50051]++
														out = append(out, '"')
														return out
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:162
	// _ = "end of CoverTab[50051]"
}

// indexNeedEscapeInString returns the index of the character that needs
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:165
// escaping. If no characters need escaping, this returns the input length.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:167
func indexNeedEscapeInString(s string) int {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:167
	_go_fuzz_dep_.CoverTab[50072]++
														for i := 0; i < len(s); i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:168
		_go_fuzz_dep_.CoverTab[50074]++
															if c := s[i]; c < ' ' || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:169
			_go_fuzz_dep_.CoverTab[50075]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:169
			return c == '"'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:169
			// _ = "end of CoverTab[50075]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:169
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:169
			_go_fuzz_dep_.CoverTab[50076]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:169
			return c == '\''
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:169
			// _ = "end of CoverTab[50076]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:169
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:169
			_go_fuzz_dep_.CoverTab[50077]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:169
			return c == '\\'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:169
			// _ = "end of CoverTab[50077]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:169
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:169
			_go_fuzz_dep_.CoverTab[50078]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:169
			return c >= 0x7f
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:169
			// _ = "end of CoverTab[50078]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:169
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:169
			_go_fuzz_dep_.CoverTab[50079]++
																return i
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:170
			// _ = "end of CoverTab[50079]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:171
			_go_fuzz_dep_.CoverTab[50080]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:171
			// _ = "end of CoverTab[50080]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:171
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:171
		// _ = "end of CoverTab[50074]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:172
	// _ = "end of CoverTab[50072]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:172
	_go_fuzz_dep_.CoverTab[50073]++
														return len(s)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:173
	// _ = "end of CoverTab[50073]"
}

// WriteFloat writes out the given float value for given bitSize.
func (e *Encoder) WriteFloat(n float64, bitSize int) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:177
	_go_fuzz_dep_.CoverTab[50081]++
														e.prepareNext(scalar)
														e.out = appendFloat(e.out, n, bitSize)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:179
	// _ = "end of CoverTab[50081]"
}

func appendFloat(out []byte, n float64, bitSize int) []byte {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:182
	_go_fuzz_dep_.CoverTab[50082]++
														switch {
	case math.IsNaN(n):
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:184
		_go_fuzz_dep_.CoverTab[50083]++
															return append(out, "nan"...)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:185
		// _ = "end of CoverTab[50083]"
	case math.IsInf(n, +1):
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:186
		_go_fuzz_dep_.CoverTab[50084]++
															return append(out, "inf"...)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:187
		// _ = "end of CoverTab[50084]"
	case math.IsInf(n, -1):
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:188
		_go_fuzz_dep_.CoverTab[50085]++
															return append(out, "-inf"...)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:189
		// _ = "end of CoverTab[50085]"
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:190
		_go_fuzz_dep_.CoverTab[50086]++
															return strconv.AppendFloat(out, n, 'g', -1, bitSize)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:191
		// _ = "end of CoverTab[50086]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:192
	// _ = "end of CoverTab[50082]"
}

// WriteInt writes out the given signed integer value.
func (e *Encoder) WriteInt(n int64) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:196
	_go_fuzz_dep_.CoverTab[50087]++
														e.prepareNext(scalar)
														e.out = append(e.out, strconv.FormatInt(n, 10)...)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:198
	// _ = "end of CoverTab[50087]"
}

// WriteUint writes out the given unsigned integer value.
func (e *Encoder) WriteUint(n uint64) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:202
	_go_fuzz_dep_.CoverTab[50088]++
														e.prepareNext(scalar)
														e.out = append(e.out, strconv.FormatUint(n, 10)...)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:204
	// _ = "end of CoverTab[50088]"
}

// WriteLiteral writes out the given string as a literal value without quotes.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:207
// This is used for writing enum literal strings.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:209
func (e *Encoder) WriteLiteral(s string) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:209
	_go_fuzz_dep_.CoverTab[50089]++
														e.prepareNext(scalar)
														e.out = append(e.out, s...)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:211
	// _ = "end of CoverTab[50089]"
}

// prepareNext adds possible space and indentation for the next value based
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:214
// on last encType and indent option. It also updates e.lastType to next.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:216
func (e *Encoder) prepareNext(next encType) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:216
	_go_fuzz_dep_.CoverTab[50090]++
														defer func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:217
		_go_fuzz_dep_.CoverTab[50093]++
															e.lastType = next
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:218
		// _ = "end of CoverTab[50093]"
	}()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:219
	// _ = "end of CoverTab[50090]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:219
	_go_fuzz_dep_.CoverTab[50091]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:222
	if len(e.indent) == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:222
		_go_fuzz_dep_.CoverTab[50094]++

															if e.lastType&(scalar|messageClose) != 0 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:224
			_go_fuzz_dep_.CoverTab[50096]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:224
			return next == name
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:224
			// _ = "end of CoverTab[50096]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:224
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:224
			_go_fuzz_dep_.CoverTab[50097]++
																e.out = append(e.out, ' ')

																if detrand.Bool() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:227
				_go_fuzz_dep_.CoverTab[50098]++
																	e.out = append(e.out, ' ')
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:228
				// _ = "end of CoverTab[50098]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:229
				_go_fuzz_dep_.CoverTab[50099]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:229
				// _ = "end of CoverTab[50099]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:229
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:229
			// _ = "end of CoverTab[50097]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:230
			_go_fuzz_dep_.CoverTab[50100]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:230
			// _ = "end of CoverTab[50100]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:230
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:230
		// _ = "end of CoverTab[50094]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:230
		_go_fuzz_dep_.CoverTab[50095]++
															return
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:231
		// _ = "end of CoverTab[50095]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:232
		_go_fuzz_dep_.CoverTab[50101]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:232
		// _ = "end of CoverTab[50101]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:232
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:232
	// _ = "end of CoverTab[50091]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:232
	_go_fuzz_dep_.CoverTab[50092]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:235
	switch {
	case e.lastType == name:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:236
		_go_fuzz_dep_.CoverTab[50102]++
															e.out = append(e.out, ' ')

															if detrand.Bool() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:239
			_go_fuzz_dep_.CoverTab[50107]++
																e.out = append(e.out, ' ')
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:240
			// _ = "end of CoverTab[50107]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:241
			_go_fuzz_dep_.CoverTab[50108]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:241
			// _ = "end of CoverTab[50108]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:241
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:241
		// _ = "end of CoverTab[50102]"

	case e.lastType == messageOpen && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:243
		_go_fuzz_dep_.CoverTab[50109]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:243
		return next != messageClose
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:243
		// _ = "end of CoverTab[50109]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:243
	}():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:243
		_go_fuzz_dep_.CoverTab[50103]++
															e.indents = append(e.indents, e.indent...)
															e.out = append(e.out, '\n')
															e.out = append(e.out, e.indents...)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:246
		// _ = "end of CoverTab[50103]"

	case e.lastType&(scalar|messageClose) != 0:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:248
		_go_fuzz_dep_.CoverTab[50104]++
															if next == messageClose {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:249
			_go_fuzz_dep_.CoverTab[50110]++
																e.indents = e.indents[:len(e.indents)-len(e.indent)]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:250
			// _ = "end of CoverTab[50110]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:251
			_go_fuzz_dep_.CoverTab[50111]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:251
			// _ = "end of CoverTab[50111]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:251
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:251
		// _ = "end of CoverTab[50104]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:251
		_go_fuzz_dep_.CoverTab[50105]++
															e.out = append(e.out, '\n')
															e.out = append(e.out, e.indents...)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:253
		// _ = "end of CoverTab[50105]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:253
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:253
		_go_fuzz_dep_.CoverTab[50106]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:253
		// _ = "end of CoverTab[50106]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:254
	// _ = "end of CoverTab[50092]"
}

// Snapshot returns the current snapshot for use in Reset.
func (e *Encoder) Snapshot() encoderState {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:258
	_go_fuzz_dep_.CoverTab[50112]++
														return e.encoderState
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:259
	// _ = "end of CoverTab[50112]"
}

// Reset resets the Encoder to the given encoderState from a Snapshot.
func (e *Encoder) Reset(es encoderState) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:263
	_go_fuzz_dep_.CoverTab[50113]++
														e.encoderState = es
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:264
	// _ = "end of CoverTab[50113]"
}

// AppendString appends the escaped form of the input string to b.
func AppendString(b []byte, s string) []byte {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:268
	_go_fuzz_dep_.CoverTab[50114]++
														return appendString(b, s, false)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:269
	// _ = "end of CoverTab[50114]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:270
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/encode.go:270
var _ = _go_fuzz_dep_.CoverTab
