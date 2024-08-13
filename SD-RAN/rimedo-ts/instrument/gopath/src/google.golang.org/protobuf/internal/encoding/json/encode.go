// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:5
package json

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:5
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:5
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

// kind represents an encoding type.
type kind uint8

const (
	_	kind	= (1 << iota) / 2
	name
	scalar
	objectOpen
	objectClose
	arrayOpen
	arrayClose
)

// Encoder provides methods to write out JSON constructs and values. The user is
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:31
// responsible for producing valid sequences of JSON constructs and values.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:33
type Encoder struct {
	indent		string
	lastKind	kind
	indents		[]byte
	out		[]byte
}

// NewEncoder returns an Encoder.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:40
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:40
// If indent is a non-empty string, it causes every entry for an Array or Object
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:40
// to be preceded by the indent and trailed by a newline.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:44
func NewEncoder(indent string) (*Encoder, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:44
	_go_fuzz_dep_.CoverTab[65813]++
														e := &Encoder{}
														if len(indent) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:46
		_go_fuzz_dep_.CoverTab[65815]++
															if strings.Trim(indent, " \t") != "" {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:47
			_go_fuzz_dep_.CoverTab[65817]++
																return nil, errors.New("indent may only be composed of space or tab characters")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:48
			// _ = "end of CoverTab[65817]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:49
			_go_fuzz_dep_.CoverTab[65818]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:49
			// _ = "end of CoverTab[65818]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:49
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:49
		// _ = "end of CoverTab[65815]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:49
		_go_fuzz_dep_.CoverTab[65816]++
															e.indent = indent
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:50
		// _ = "end of CoverTab[65816]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:51
		_go_fuzz_dep_.CoverTab[65819]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:51
		// _ = "end of CoverTab[65819]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:51
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:51
	// _ = "end of CoverTab[65813]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:51
	_go_fuzz_dep_.CoverTab[65814]++
														return e, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:52
	// _ = "end of CoverTab[65814]"
}

// Bytes returns the content of the written bytes.
func (e *Encoder) Bytes() []byte {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:56
	_go_fuzz_dep_.CoverTab[65820]++
														return e.out
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:57
	// _ = "end of CoverTab[65820]"
}

// WriteNull writes out the null value.
func (e *Encoder) WriteNull() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:61
	_go_fuzz_dep_.CoverTab[65821]++
														e.prepareNext(scalar)
														e.out = append(e.out, "null"...)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:63
	// _ = "end of CoverTab[65821]"
}

// WriteBool writes out the given boolean value.
func (e *Encoder) WriteBool(b bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:67
	_go_fuzz_dep_.CoverTab[65822]++
														e.prepareNext(scalar)
														if b {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:69
		_go_fuzz_dep_.CoverTab[65823]++
															e.out = append(e.out, "true"...)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:70
		// _ = "end of CoverTab[65823]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:71
		_go_fuzz_dep_.CoverTab[65824]++
															e.out = append(e.out, "false"...)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:72
		// _ = "end of CoverTab[65824]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:73
	// _ = "end of CoverTab[65822]"
}

// WriteString writes out the given string in JSON string value. Returns error
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:76
// if input string contains invalid UTF-8.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:78
func (e *Encoder) WriteString(s string) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:78
	_go_fuzz_dep_.CoverTab[65825]++
														e.prepareNext(scalar)
														var err error
														if e.out, err = appendString(e.out, s); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:81
		_go_fuzz_dep_.CoverTab[65827]++
															return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:82
		// _ = "end of CoverTab[65827]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:83
		_go_fuzz_dep_.CoverTab[65828]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:83
		// _ = "end of CoverTab[65828]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:83
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:83
	// _ = "end of CoverTab[65825]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:83
	_go_fuzz_dep_.CoverTab[65826]++
														return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:84
	// _ = "end of CoverTab[65826]"
}

// Sentinel error used for indicating invalid UTF-8.
var errInvalidUTF8 = errors.New("invalid UTF-8")

func appendString(out []byte, in string) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:90
	_go_fuzz_dep_.CoverTab[65829]++
														out = append(out, '"')
														i := indexNeedEscapeInString(in)
														in, out = in[i:], append(out, in[:i]...)
														for len(in) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:94
		_go_fuzz_dep_.CoverTab[65831]++
															switch r, n := utf8.DecodeRuneInString(in); {
		case r == utf8.RuneError && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:96
			_go_fuzz_dep_.CoverTab[65836]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:96
			return n == 1
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:96
			// _ = "end of CoverTab[65836]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:96
		}():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:96
			_go_fuzz_dep_.CoverTab[65832]++
																return out, errInvalidUTF8
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:97
			// _ = "end of CoverTab[65832]"
		case r < ' ' || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:98
			_go_fuzz_dep_.CoverTab[65837]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:98
			return r == '"'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:98
			// _ = "end of CoverTab[65837]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:98
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:98
			_go_fuzz_dep_.CoverTab[65838]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:98
			return r == '\\'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:98
			// _ = "end of CoverTab[65838]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:98
		}():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:98
			_go_fuzz_dep_.CoverTab[65833]++
																out = append(out, '\\')
																switch r {
			case '"', '\\':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:101
				_go_fuzz_dep_.CoverTab[65839]++
																	out = append(out, byte(r))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:102
				// _ = "end of CoverTab[65839]"
			case '\b':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:103
				_go_fuzz_dep_.CoverTab[65840]++
																	out = append(out, 'b')
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:104
				// _ = "end of CoverTab[65840]"
			case '\f':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:105
				_go_fuzz_dep_.CoverTab[65841]++
																	out = append(out, 'f')
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:106
				// _ = "end of CoverTab[65841]"
			case '\n':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:107
				_go_fuzz_dep_.CoverTab[65842]++
																	out = append(out, 'n')
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:108
				// _ = "end of CoverTab[65842]"
			case '\r':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:109
				_go_fuzz_dep_.CoverTab[65843]++
																	out = append(out, 'r')
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:110
				// _ = "end of CoverTab[65843]"
			case '\t':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:111
				_go_fuzz_dep_.CoverTab[65844]++
																	out = append(out, 't')
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:112
				// _ = "end of CoverTab[65844]"
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:113
				_go_fuzz_dep_.CoverTab[65845]++
																	out = append(out, 'u')
																	out = append(out, "0000"[1+(bits.Len32(uint32(r))-1)/4:]...)
																	out = strconv.AppendUint(out, uint64(r), 16)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:116
				// _ = "end of CoverTab[65845]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:117
			// _ = "end of CoverTab[65833]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:117
			_go_fuzz_dep_.CoverTab[65834]++
																in = in[n:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:118
			// _ = "end of CoverTab[65834]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:119
			_go_fuzz_dep_.CoverTab[65835]++
																i := indexNeedEscapeInString(in[n:])
																in, out = in[n+i:], append(out, in[:n+i]...)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:121
			// _ = "end of CoverTab[65835]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:122
		// _ = "end of CoverTab[65831]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:123
	// _ = "end of CoverTab[65829]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:123
	_go_fuzz_dep_.CoverTab[65830]++
														out = append(out, '"')
														return out, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:125
	// _ = "end of CoverTab[65830]"
}

// indexNeedEscapeInString returns the index of the character that needs
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:128
// escaping. If no characters need escaping, this returns the input length.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:130
func indexNeedEscapeInString(s string) int {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:130
	_go_fuzz_dep_.CoverTab[65846]++
														for i, r := range s {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:131
		_go_fuzz_dep_.CoverTab[65848]++
															if r < ' ' || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:132
			_go_fuzz_dep_.CoverTab[65849]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:132
			return r == '\\'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:132
			// _ = "end of CoverTab[65849]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:132
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:132
			_go_fuzz_dep_.CoverTab[65850]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:132
			return r == '"'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:132
			// _ = "end of CoverTab[65850]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:132
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:132
			_go_fuzz_dep_.CoverTab[65851]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:132
			return r == utf8.RuneError
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:132
			// _ = "end of CoverTab[65851]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:132
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:132
			_go_fuzz_dep_.CoverTab[65852]++
																return i
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:133
			// _ = "end of CoverTab[65852]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:134
			_go_fuzz_dep_.CoverTab[65853]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:134
			// _ = "end of CoverTab[65853]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:134
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:134
		// _ = "end of CoverTab[65848]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:135
	// _ = "end of CoverTab[65846]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:135
	_go_fuzz_dep_.CoverTab[65847]++
														return len(s)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:136
	// _ = "end of CoverTab[65847]"
}

// WriteFloat writes out the given float and bitSize in JSON number value.
func (e *Encoder) WriteFloat(n float64, bitSize int) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:140
	_go_fuzz_dep_.CoverTab[65854]++
														e.prepareNext(scalar)
														e.out = appendFloat(e.out, n, bitSize)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:142
	// _ = "end of CoverTab[65854]"
}

// appendFloat formats given float in bitSize, and appends to the given []byte.
func appendFloat(out []byte, n float64, bitSize int) []byte {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:146
	_go_fuzz_dep_.CoverTab[65855]++
														switch {
	case math.IsNaN(n):
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:148
		_go_fuzz_dep_.CoverTab[65859]++
															return append(out, `"NaN"`...)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:149
		// _ = "end of CoverTab[65859]"
	case math.IsInf(n, +1):
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:150
		_go_fuzz_dep_.CoverTab[65860]++
															return append(out, `"Infinity"`...)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:151
		// _ = "end of CoverTab[65860]"
	case math.IsInf(n, -1):
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:152
		_go_fuzz_dep_.CoverTab[65861]++
															return append(out, `"-Infinity"`...)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:153
		// _ = "end of CoverTab[65861]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:153
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:153
		_go_fuzz_dep_.CoverTab[65862]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:153
		// _ = "end of CoverTab[65862]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:154
	// _ = "end of CoverTab[65855]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:154
	_go_fuzz_dep_.CoverTab[65856]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:158
	fmt := byte('f')
	if abs := math.Abs(n); abs != 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:159
		_go_fuzz_dep_.CoverTab[65863]++
															if bitSize == 64 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:160
			_go_fuzz_dep_.CoverTab[65864]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:160
			return (abs < 1e-6 || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:160
				_go_fuzz_dep_.CoverTab[65865]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:160
				return abs >= 1e21
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:160
				// _ = "end of CoverTab[65865]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:160
			}())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:160
			// _ = "end of CoverTab[65864]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:160
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:160
			_go_fuzz_dep_.CoverTab[65866]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:160
			return bitSize == 32 && func() bool {
																	_go_fuzz_dep_.CoverTab[65867]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:161
				return (float32(abs) < 1e-6 || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:161
					_go_fuzz_dep_.CoverTab[65868]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:161
					return float32(abs) >= 1e21
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:161
					// _ = "end of CoverTab[65868]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:161
				}())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:161
				// _ = "end of CoverTab[65867]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:161
			}()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:161
			// _ = "end of CoverTab[65866]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:161
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:161
			_go_fuzz_dep_.CoverTab[65869]++
																fmt = 'e'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:162
			// _ = "end of CoverTab[65869]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:163
			_go_fuzz_dep_.CoverTab[65870]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:163
			// _ = "end of CoverTab[65870]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:163
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:163
		// _ = "end of CoverTab[65863]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:164
		_go_fuzz_dep_.CoverTab[65871]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:164
		// _ = "end of CoverTab[65871]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:164
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:164
	// _ = "end of CoverTab[65856]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:164
	_go_fuzz_dep_.CoverTab[65857]++
														out = strconv.AppendFloat(out, n, fmt, -1, bitSize)
														if fmt == 'e' {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:166
		_go_fuzz_dep_.CoverTab[65872]++
															n := len(out)
															if n >= 4 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:168
			_go_fuzz_dep_.CoverTab[65873]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:168
			return out[n-4] == 'e'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:168
			// _ = "end of CoverTab[65873]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:168
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:168
			_go_fuzz_dep_.CoverTab[65874]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:168
			return out[n-3] == '-'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:168
			// _ = "end of CoverTab[65874]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:168
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:168
			_go_fuzz_dep_.CoverTab[65875]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:168
			return out[n-2] == '0'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:168
			// _ = "end of CoverTab[65875]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:168
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:168
			_go_fuzz_dep_.CoverTab[65876]++
																out[n-2] = out[n-1]
																out = out[:n-1]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:170
			// _ = "end of CoverTab[65876]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:171
			_go_fuzz_dep_.CoverTab[65877]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:171
			// _ = "end of CoverTab[65877]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:171
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:171
		// _ = "end of CoverTab[65872]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:172
		_go_fuzz_dep_.CoverTab[65878]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:172
		// _ = "end of CoverTab[65878]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:172
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:172
	// _ = "end of CoverTab[65857]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:172
	_go_fuzz_dep_.CoverTab[65858]++
														return out
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:173
	// _ = "end of CoverTab[65858]"
}

// WriteInt writes out the given signed integer in JSON number value.
func (e *Encoder) WriteInt(n int64) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:177
	_go_fuzz_dep_.CoverTab[65879]++
														e.prepareNext(scalar)
														e.out = append(e.out, strconv.FormatInt(n, 10)...)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:179
	// _ = "end of CoverTab[65879]"
}

// WriteUint writes out the given unsigned integer in JSON number value.
func (e *Encoder) WriteUint(n uint64) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:183
	_go_fuzz_dep_.CoverTab[65880]++
														e.prepareNext(scalar)
														e.out = append(e.out, strconv.FormatUint(n, 10)...)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:185
	// _ = "end of CoverTab[65880]"
}

// StartObject writes out the '{' symbol.
func (e *Encoder) StartObject() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:189
	_go_fuzz_dep_.CoverTab[65881]++
														e.prepareNext(objectOpen)
														e.out = append(e.out, '{')
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:191
	// _ = "end of CoverTab[65881]"
}

// EndObject writes out the '}' symbol.
func (e *Encoder) EndObject() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:195
	_go_fuzz_dep_.CoverTab[65882]++
														e.prepareNext(objectClose)
														e.out = append(e.out, '}')
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:197
	// _ = "end of CoverTab[65882]"
}

// WriteName writes out the given string in JSON string value and the name
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:200
// separator ':'. Returns error if input string contains invalid UTF-8, which
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:200
// should not be likely as protobuf field names should be valid.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:203
func (e *Encoder) WriteName(s string) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:203
	_go_fuzz_dep_.CoverTab[65883]++
														e.prepareNext(name)
														var err error

														e.out, err = appendString(e.out, s)
														e.out = append(e.out, ':')
														return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:209
	// _ = "end of CoverTab[65883]"
}

// StartArray writes out the '[' symbol.
func (e *Encoder) StartArray() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:213
	_go_fuzz_dep_.CoverTab[65884]++
														e.prepareNext(arrayOpen)
														e.out = append(e.out, '[')
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:215
	// _ = "end of CoverTab[65884]"
}

// EndArray writes out the ']' symbol.
func (e *Encoder) EndArray() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:219
	_go_fuzz_dep_.CoverTab[65885]++
														e.prepareNext(arrayClose)
														e.out = append(e.out, ']')
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:221
	// _ = "end of CoverTab[65885]"
}

// prepareNext adds possible comma and indentation for the next value based
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:224
// on last type and indent option. It also updates lastKind to next.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:226
func (e *Encoder) prepareNext(next kind) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:226
	_go_fuzz_dep_.CoverTab[65886]++
														defer func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:227
		_go_fuzz_dep_.CoverTab[65889]++

															e.lastKind = next
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:229
		// _ = "end of CoverTab[65889]"
	}()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:230
	// _ = "end of CoverTab[65886]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:230
	_go_fuzz_dep_.CoverTab[65887]++

														if len(e.indent) == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:232
		_go_fuzz_dep_.CoverTab[65890]++

															if e.lastKind&(scalar|objectClose|arrayClose) != 0 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:234
			_go_fuzz_dep_.CoverTab[65892]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:234
			return next&(name|scalar|objectOpen|arrayOpen) != 0
																// _ = "end of CoverTab[65892]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:235
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:235
			_go_fuzz_dep_.CoverTab[65893]++
																e.out = append(e.out, ',')

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:239
			if detrand.Bool() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:239
				_go_fuzz_dep_.CoverTab[65894]++
																	e.out = append(e.out, ' ')
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:240
				// _ = "end of CoverTab[65894]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:241
				_go_fuzz_dep_.CoverTab[65895]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:241
				// _ = "end of CoverTab[65895]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:241
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:241
			// _ = "end of CoverTab[65893]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:242
			_go_fuzz_dep_.CoverTab[65896]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:242
			// _ = "end of CoverTab[65896]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:242
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:242
		// _ = "end of CoverTab[65890]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:242
		_go_fuzz_dep_.CoverTab[65891]++
															return
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:243
		// _ = "end of CoverTab[65891]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:244
		_go_fuzz_dep_.CoverTab[65897]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:244
		// _ = "end of CoverTab[65897]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:244
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:244
	// _ = "end of CoverTab[65887]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:244
	_go_fuzz_dep_.CoverTab[65888]++

														switch {
	case e.lastKind&(objectOpen|arrayOpen) != 0:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:247
		_go_fuzz_dep_.CoverTab[65898]++

															if next&(objectClose|arrayClose) == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:249
			_go_fuzz_dep_.CoverTab[65903]++
																e.indents = append(e.indents, e.indent...)
																e.out = append(e.out, '\n')
																e.out = append(e.out, e.indents...)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:252
			// _ = "end of CoverTab[65903]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:253
			_go_fuzz_dep_.CoverTab[65904]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:253
			// _ = "end of CoverTab[65904]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:253
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:253
		// _ = "end of CoverTab[65898]"

	case e.lastKind&(scalar|objectClose|arrayClose) != 0:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:255
		_go_fuzz_dep_.CoverTab[65899]++
															switch {

		case next&(name|scalar|objectOpen|arrayOpen) != 0:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:258
			_go_fuzz_dep_.CoverTab[65905]++
																e.out = append(e.out, ',', '\n')
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:259
			// _ = "end of CoverTab[65905]"

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:262
		case next&(objectClose|arrayClose) != 0:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:262
			_go_fuzz_dep_.CoverTab[65906]++
																e.indents = e.indents[:len(e.indents)-len(e.indent)]
																e.out = append(e.out, '\n')
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:264
			// _ = "end of CoverTab[65906]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:264
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:264
			_go_fuzz_dep_.CoverTab[65907]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:264
			// _ = "end of CoverTab[65907]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:265
		// _ = "end of CoverTab[65899]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:265
		_go_fuzz_dep_.CoverTab[65900]++
															e.out = append(e.out, e.indents...)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:266
		// _ = "end of CoverTab[65900]"

	case e.lastKind&name != 0:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:268
		_go_fuzz_dep_.CoverTab[65901]++
															e.out = append(e.out, ' ')

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:272
		if detrand.Bool() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:272
			_go_fuzz_dep_.CoverTab[65908]++
																e.out = append(e.out, ' ')
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:273
			// _ = "end of CoverTab[65908]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:274
			_go_fuzz_dep_.CoverTab[65909]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:274
			// _ = "end of CoverTab[65909]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:274
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:274
		// _ = "end of CoverTab[65901]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:274
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:274
		_go_fuzz_dep_.CoverTab[65902]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:274
		// _ = "end of CoverTab[65902]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:275
	// _ = "end of CoverTab[65888]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:276
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/encode.go:276
var _ = _go_fuzz_dep_.CoverTab
