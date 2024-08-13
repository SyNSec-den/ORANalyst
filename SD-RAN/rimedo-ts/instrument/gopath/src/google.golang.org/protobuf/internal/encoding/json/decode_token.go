// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:5
package json

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:5
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:5
)

import (
	"bytes"
	"fmt"
	"strconv"
)

// Kind represents a token kind expressible in the JSON format.
type Kind uint16

const (
	Invalid	Kind	= (1 << iota) / 2
	EOF
	Null
	Bool
	Number
	String
	Name
	ObjectOpen
	ObjectClose
	ArrayOpen
	ArrayClose

	// comma is only for parsing in between tokens and
	// does not need to be exported.
	comma
)

func (k Kind) String() string {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:34
	_go_fuzz_dep_.CoverTab[65751]++
														switch k {
	case EOF:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:36
		_go_fuzz_dep_.CoverTab[65753]++
															return "eof"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:37
		// _ = "end of CoverTab[65753]"
	case Null:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:38
		_go_fuzz_dep_.CoverTab[65754]++
															return "null"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:39
		// _ = "end of CoverTab[65754]"
	case Bool:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:40
		_go_fuzz_dep_.CoverTab[65755]++
															return "bool"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:41
		// _ = "end of CoverTab[65755]"
	case Number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:42
		_go_fuzz_dep_.CoverTab[65756]++
															return "number"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:43
		// _ = "end of CoverTab[65756]"
	case String:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:44
		_go_fuzz_dep_.CoverTab[65757]++
															return "string"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:45
		// _ = "end of CoverTab[65757]"
	case ObjectOpen:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:46
		_go_fuzz_dep_.CoverTab[65758]++
															return "{"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:47
		// _ = "end of CoverTab[65758]"
	case ObjectClose:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:48
		_go_fuzz_dep_.CoverTab[65759]++
															return "}"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:49
		// _ = "end of CoverTab[65759]"
	case Name:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:50
		_go_fuzz_dep_.CoverTab[65760]++
															return "name"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:51
		// _ = "end of CoverTab[65760]"
	case ArrayOpen:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:52
		_go_fuzz_dep_.CoverTab[65761]++
															return "["
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:53
		// _ = "end of CoverTab[65761]"
	case ArrayClose:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:54
		_go_fuzz_dep_.CoverTab[65762]++
															return "]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:55
		// _ = "end of CoverTab[65762]"
	case comma:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:56
		_go_fuzz_dep_.CoverTab[65763]++
															return ","
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:57
		// _ = "end of CoverTab[65763]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:57
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:57
		_go_fuzz_dep_.CoverTab[65764]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:57
		// _ = "end of CoverTab[65764]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:58
	// _ = "end of CoverTab[65751]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:58
	_go_fuzz_dep_.CoverTab[65752]++
														return "<invalid>"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:59
	// _ = "end of CoverTab[65752]"
}

// Token provides a parsed token kind and value.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:62
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:62
// Values are provided by the difference accessor methods. The accessor methods
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:62
// Name, Bool, and ParsedString will panic if called on the wrong kind. There
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:62
// are different accessor methods for the Number kind for converting to the
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:62
// appropriate Go numeric type and those methods have the ok return value.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:68
type Token struct {
	// Token kind.
	kind	Kind
	// pos provides the position of the token in the original input.
	pos	int
	// raw bytes of the serialized token.
	// This is a subslice into the original input.
	raw	[]byte
	// boo is parsed boolean value.
	boo	bool
	// str is parsed string value.
	str	string
}

// Kind returns the token kind.
func (t Token) Kind() Kind {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:83
	_go_fuzz_dep_.CoverTab[65765]++
														return t.kind
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:84
	// _ = "end of CoverTab[65765]"
}

// RawString returns the read value in string.
func (t Token) RawString() string {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:88
	_go_fuzz_dep_.CoverTab[65766]++
														return string(t.raw)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:89
	// _ = "end of CoverTab[65766]"
}

// Pos returns the token position from the input.
func (t Token) Pos() int {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:93
	_go_fuzz_dep_.CoverTab[65767]++
														return t.pos
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:94
	// _ = "end of CoverTab[65767]"
}

// Name returns the object name if token is Name, else it panics.
func (t Token) Name() string {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:98
	_go_fuzz_dep_.CoverTab[65768]++
														if t.kind == Name {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:99
			_go_fuzz_dep_.CoverTab[65770]++
																return t.str
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:100
		// _ = "end of CoverTab[65770]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:101
		_go_fuzz_dep_.CoverTab[65771]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:101
		// _ = "end of CoverTab[65771]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:101
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:101
	// _ = "end of CoverTab[65768]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:101
	_go_fuzz_dep_.CoverTab[65769]++
															panic(fmt.Sprintf("Token is not a Name: %v", t.RawString()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:102
	// _ = "end of CoverTab[65769]"
}

// Bool returns the bool value if token kind is Bool, else it panics.
func (t Token) Bool() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:106
	_go_fuzz_dep_.CoverTab[65772]++
															if t.kind == Bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:107
		_go_fuzz_dep_.CoverTab[65774]++
																return t.boo
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:108
		// _ = "end of CoverTab[65774]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:109
		_go_fuzz_dep_.CoverTab[65775]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:109
		// _ = "end of CoverTab[65775]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:109
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:109
	// _ = "end of CoverTab[65772]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:109
	_go_fuzz_dep_.CoverTab[65773]++
															panic(fmt.Sprintf("Token is not a Bool: %v", t.RawString()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:110
	// _ = "end of CoverTab[65773]"
}

// ParsedString returns the string value for a JSON string token or the read
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:113
// value in string if token is not a string.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:115
func (t Token) ParsedString() string {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:115
	_go_fuzz_dep_.CoverTab[65776]++
															if t.kind == String {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:116
		_go_fuzz_dep_.CoverTab[65778]++
																return t.str
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:117
		// _ = "end of CoverTab[65778]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:118
		_go_fuzz_dep_.CoverTab[65779]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:118
		// _ = "end of CoverTab[65779]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:118
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:118
	// _ = "end of CoverTab[65776]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:118
	_go_fuzz_dep_.CoverTab[65777]++
															panic(fmt.Sprintf("Token is not a String: %v", t.RawString()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:119
	// _ = "end of CoverTab[65777]"
}

// Float returns the floating-point number if token kind is Number.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:122
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:122
// The floating-point precision is specified by the bitSize parameter: 32 for
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:122
// float32 or 64 for float64. If bitSize=32, the result still has type float64,
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:122
// but it will be convertible to float32 without changing its value. It will
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:122
// return false if the number exceeds the floating point limits for given
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:122
// bitSize.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:129
func (t Token) Float(bitSize int) (float64, bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:129
	_go_fuzz_dep_.CoverTab[65780]++
															if t.kind != Number {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:130
		_go_fuzz_dep_.CoverTab[65783]++
																return 0, false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:131
		// _ = "end of CoverTab[65783]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:132
		_go_fuzz_dep_.CoverTab[65784]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:132
		// _ = "end of CoverTab[65784]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:132
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:132
	// _ = "end of CoverTab[65780]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:132
	_go_fuzz_dep_.CoverTab[65781]++
															f, err := strconv.ParseFloat(t.RawString(), bitSize)
															if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:134
		_go_fuzz_dep_.CoverTab[65785]++
																return 0, false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:135
		// _ = "end of CoverTab[65785]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:136
		_go_fuzz_dep_.CoverTab[65786]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:136
		// _ = "end of CoverTab[65786]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:136
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:136
	// _ = "end of CoverTab[65781]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:136
	_go_fuzz_dep_.CoverTab[65782]++
															return f, true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:137
	// _ = "end of CoverTab[65782]"
}

// Int returns the signed integer number if token is Number.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:140
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:140
// The given bitSize specifies the integer type that the result must fit into.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:140
// It returns false if the number is not an integer value or if the result
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:140
// exceeds the limits for given bitSize.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:145
func (t Token) Int(bitSize int) (int64, bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:145
	_go_fuzz_dep_.CoverTab[65787]++
															s, ok := t.getIntStr()
															if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:147
		_go_fuzz_dep_.CoverTab[65790]++
																return 0, false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:148
		// _ = "end of CoverTab[65790]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:149
		_go_fuzz_dep_.CoverTab[65791]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:149
		// _ = "end of CoverTab[65791]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:149
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:149
	// _ = "end of CoverTab[65787]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:149
	_go_fuzz_dep_.CoverTab[65788]++
															n, err := strconv.ParseInt(s, 10, bitSize)
															if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:151
		_go_fuzz_dep_.CoverTab[65792]++
																return 0, false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:152
		// _ = "end of CoverTab[65792]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:153
		_go_fuzz_dep_.CoverTab[65793]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:153
		// _ = "end of CoverTab[65793]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:153
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:153
	// _ = "end of CoverTab[65788]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:153
	_go_fuzz_dep_.CoverTab[65789]++
															return n, true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:154
	// _ = "end of CoverTab[65789]"
}

// Uint returns the signed integer number if token is Number.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:157
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:157
// The given bitSize specifies the unsigned integer type that the result must
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:157
// fit into. It returns false if the number is not an unsigned integer value
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:157
// or if the result exceeds the limits for given bitSize.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:162
func (t Token) Uint(bitSize int) (uint64, bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:162
	_go_fuzz_dep_.CoverTab[65794]++
															s, ok := t.getIntStr()
															if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:164
		_go_fuzz_dep_.CoverTab[65797]++
																return 0, false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:165
		// _ = "end of CoverTab[65797]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:166
		_go_fuzz_dep_.CoverTab[65798]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:166
		// _ = "end of CoverTab[65798]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:166
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:166
	// _ = "end of CoverTab[65794]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:166
	_go_fuzz_dep_.CoverTab[65795]++
															n, err := strconv.ParseUint(s, 10, bitSize)
															if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:168
		_go_fuzz_dep_.CoverTab[65799]++
																return 0, false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:169
		// _ = "end of CoverTab[65799]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:170
		_go_fuzz_dep_.CoverTab[65800]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:170
		// _ = "end of CoverTab[65800]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:170
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:170
	// _ = "end of CoverTab[65795]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:170
	_go_fuzz_dep_.CoverTab[65796]++
															return n, true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:171
	// _ = "end of CoverTab[65796]"
}

func (t Token) getIntStr() (string, bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:174
	_go_fuzz_dep_.CoverTab[65801]++
															if t.kind != Number {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:175
		_go_fuzz_dep_.CoverTab[65804]++
																return "", false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:176
		// _ = "end of CoverTab[65804]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:177
		_go_fuzz_dep_.CoverTab[65805]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:177
		// _ = "end of CoverTab[65805]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:177
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:177
	// _ = "end of CoverTab[65801]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:177
	_go_fuzz_dep_.CoverTab[65802]++
															parts, ok := parseNumberParts(t.raw)
															if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:179
		_go_fuzz_dep_.CoverTab[65806]++
																return "", false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:180
		// _ = "end of CoverTab[65806]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:181
		_go_fuzz_dep_.CoverTab[65807]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:181
		// _ = "end of CoverTab[65807]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:181
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:181
	// _ = "end of CoverTab[65802]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:181
	_go_fuzz_dep_.CoverTab[65803]++
															return normalizeToIntString(parts)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:182
	// _ = "end of CoverTab[65803]"
}

// TokenEquals returns true if given Tokens are equal, else false.
func TokenEquals(x, y Token) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:186
	_go_fuzz_dep_.CoverTab[65808]++
															return x.kind == y.kind && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:187
		_go_fuzz_dep_.CoverTab[65809]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:187
		return x.pos == y.pos
																// _ = "end of CoverTab[65809]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:188
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:188
		_go_fuzz_dep_.CoverTab[65810]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:188
		return bytes.Equal(x.raw, y.raw)
																// _ = "end of CoverTab[65810]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:189
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:189
		_go_fuzz_dep_.CoverTab[65811]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:189
		return x.boo == y.boo
																// _ = "end of CoverTab[65811]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:190
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:190
		_go_fuzz_dep_.CoverTab[65812]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:190
		return x.str == y.str
																// _ = "end of CoverTab[65812]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:191
	}()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:191
	// _ = "end of CoverTab[65808]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:192
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_token.go:192
var _ = _go_fuzz_dep_.CoverTab
