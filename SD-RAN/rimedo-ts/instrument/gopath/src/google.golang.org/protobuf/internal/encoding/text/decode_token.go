// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:5
package text

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:5
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:5
)

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
	"strings"

	"google.golang.org/protobuf/internal/flags"
)

// Kind represents a token kind expressible in the textproto format.
type Kind uint8

// Kind values.
const (
	Invalid	Kind	= iota
	EOF
	Name	// Name indicates the field name.
	Scalar	// Scalar are scalar values, e.g. "string", 47, ENUM_LITERAL, true.
	MessageOpen
	MessageClose
	ListOpen
	ListClose

	// comma and semi-colon are only for parsing in between values and should not be exposed.
	comma
	semicolon

	// bof indicates beginning of file, which is the default token
	// kind at the beginning of parsing.
	bof	= Invalid
)

func (t Kind) String() string {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:40
	_go_fuzz_dep_.CoverTab[49875]++
														switch t {
	case Invalid:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:42
		_go_fuzz_dep_.CoverTab[49876]++
															return "<invalid>"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:43
		// _ = "end of CoverTab[49876]"
	case EOF:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:44
		_go_fuzz_dep_.CoverTab[49877]++
															return "eof"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:45
		// _ = "end of CoverTab[49877]"
	case Scalar:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:46
		_go_fuzz_dep_.CoverTab[49878]++
															return "scalar"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:47
		// _ = "end of CoverTab[49878]"
	case Name:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:48
		_go_fuzz_dep_.CoverTab[49879]++
															return "name"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:49
		// _ = "end of CoverTab[49879]"
	case MessageOpen:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:50
		_go_fuzz_dep_.CoverTab[49880]++
															return "{"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:51
		// _ = "end of CoverTab[49880]"
	case MessageClose:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:52
		_go_fuzz_dep_.CoverTab[49881]++
															return "}"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:53
		// _ = "end of CoverTab[49881]"
	case ListOpen:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:54
		_go_fuzz_dep_.CoverTab[49882]++
															return "["
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:55
		// _ = "end of CoverTab[49882]"
	case ListClose:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:56
		_go_fuzz_dep_.CoverTab[49883]++
															return "]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:57
		// _ = "end of CoverTab[49883]"
	case comma:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:58
		_go_fuzz_dep_.CoverTab[49884]++
															return ","
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:59
		// _ = "end of CoverTab[49884]"
	case semicolon:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:60
		_go_fuzz_dep_.CoverTab[49885]++
															return ";"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:61
		// _ = "end of CoverTab[49885]"
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:62
		_go_fuzz_dep_.CoverTab[49886]++
															return fmt.Sprintf("<invalid:%v>", uint8(t))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:63
		// _ = "end of CoverTab[49886]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:64
	// _ = "end of CoverTab[49875]"
}

// NameKind represents different types of field names.
type NameKind uint8

// NameKind values.
const (
	IdentName	NameKind	= iota + 1
	TypeName
	FieldNumber
)

func (t NameKind) String() string {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:77
	_go_fuzz_dep_.CoverTab[49887]++
														switch t {
	case IdentName:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:79
		_go_fuzz_dep_.CoverTab[49888]++
															return "IdentName"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:80
		// _ = "end of CoverTab[49888]"
	case TypeName:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:81
		_go_fuzz_dep_.CoverTab[49889]++
															return "TypeName"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:82
		// _ = "end of CoverTab[49889]"
	case FieldNumber:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:83
		_go_fuzz_dep_.CoverTab[49890]++
															return "FieldNumber"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:84
		// _ = "end of CoverTab[49890]"
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:85
		_go_fuzz_dep_.CoverTab[49891]++
															return fmt.Sprintf("<invalid:%v>", uint8(t))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:86
		// _ = "end of CoverTab[49891]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:87
	// _ = "end of CoverTab[49887]"
}

// Bit mask in Token.attrs to indicate if a Name token is followed by the
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:90
// separator char ':'. The field name separator char is optional for message
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:90
// field or repeated message field, but required for all other types. Decoder
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:90
// simply indicates whether a Name token is followed by separator or not.  It is
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:90
// up to the prototext package to validate.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:95
const hasSeparator = 1 << 7

// Scalar value types.
const (
	numberValue	= iota + 1
	stringValue
	literalValue
)

// Bit mask in Token.numAttrs to indicate that the number is a negative.
const isNegative = 1 << 7

// Token provides a parsed token kind and value. Values are provided by the
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:107
// different accessor methods.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:109
type Token struct {
	// Kind of the Token object.
	kind	Kind
	// attrs contains metadata for the following Kinds:
	// Name: hasSeparator bit and one of NameKind.
	// Scalar: one of numberValue, stringValue, literalValue.
	attrs	uint8
	// numAttrs contains metadata for numberValue:
	// - highest bit is whether negative or positive.
	// - lower bits indicate one of numDec, numHex, numOct, numFloat.
	numAttrs	uint8
	// pos provides the position of the token in the original input.
	pos	int
	// raw bytes of the serialized token.
	// This is a subslice into the original input.
	raw	[]byte
	// str contains parsed string for the following:
	// - stringValue of Scalar kind
	// - numberValue of Scalar kind
	// - TypeName of Name kind
	str	string
}

// Kind returns the token kind.
func (t Token) Kind() Kind {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:133
	_go_fuzz_dep_.CoverTab[49892]++
															return t.kind
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:134
	// _ = "end of CoverTab[49892]"
}

// RawString returns the read value in string.
func (t Token) RawString() string {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:138
	_go_fuzz_dep_.CoverTab[49893]++
															return string(t.raw)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:139
	// _ = "end of CoverTab[49893]"
}

// Pos returns the token position from the input.
func (t Token) Pos() int {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:143
	_go_fuzz_dep_.CoverTab[49894]++
															return t.pos
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:144
	// _ = "end of CoverTab[49894]"
}

// NameKind returns IdentName, TypeName or FieldNumber.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:147
// It panics if type is not Name.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:149
func (t Token) NameKind() NameKind {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:149
	_go_fuzz_dep_.CoverTab[49895]++
															if t.kind == Name {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:150
		_go_fuzz_dep_.CoverTab[49897]++
																return NameKind(t.attrs &^ hasSeparator)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:151
		// _ = "end of CoverTab[49897]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:152
		_go_fuzz_dep_.CoverTab[49898]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:152
		// _ = "end of CoverTab[49898]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:152
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:152
	// _ = "end of CoverTab[49895]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:152
	_go_fuzz_dep_.CoverTab[49896]++
															panic(fmt.Sprintf("Token is not a Name type: %s", t.kind))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:153
	// _ = "end of CoverTab[49896]"
}

// HasSeparator returns true if the field name is followed by the separator char
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:156
// ':', else false. It panics if type is not Name.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:158
func (t Token) HasSeparator() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:158
	_go_fuzz_dep_.CoverTab[49899]++
															if t.kind == Name {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:159
		_go_fuzz_dep_.CoverTab[49901]++
																return t.attrs&hasSeparator != 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:160
		// _ = "end of CoverTab[49901]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:161
		_go_fuzz_dep_.CoverTab[49902]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:161
		// _ = "end of CoverTab[49902]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:161
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:161
	// _ = "end of CoverTab[49899]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:161
	_go_fuzz_dep_.CoverTab[49900]++
															panic(fmt.Sprintf("Token is not a Name type: %s", t.kind))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:162
	// _ = "end of CoverTab[49900]"
}

// IdentName returns the value for IdentName type.
func (t Token) IdentName() string {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:166
	_go_fuzz_dep_.CoverTab[49903]++
															if t.kind == Name && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:167
		_go_fuzz_dep_.CoverTab[49905]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:167
		return t.attrs&uint8(IdentName) != 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:167
		// _ = "end of CoverTab[49905]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:167
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:167
		_go_fuzz_dep_.CoverTab[49906]++
																return string(t.raw)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:168
		// _ = "end of CoverTab[49906]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:169
		_go_fuzz_dep_.CoverTab[49907]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:169
		// _ = "end of CoverTab[49907]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:169
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:169
	// _ = "end of CoverTab[49903]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:169
	_go_fuzz_dep_.CoverTab[49904]++
															panic(fmt.Sprintf("Token is not an IdentName: %s:%s", t.kind, NameKind(t.attrs&^hasSeparator)))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:170
	// _ = "end of CoverTab[49904]"
}

// TypeName returns the value for TypeName type.
func (t Token) TypeName() string {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:174
	_go_fuzz_dep_.CoverTab[49908]++
															if t.kind == Name && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:175
		_go_fuzz_dep_.CoverTab[49910]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:175
		return t.attrs&uint8(TypeName) != 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:175
		// _ = "end of CoverTab[49910]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:175
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:175
		_go_fuzz_dep_.CoverTab[49911]++
																return t.str
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:176
		// _ = "end of CoverTab[49911]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:177
		_go_fuzz_dep_.CoverTab[49912]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:177
		// _ = "end of CoverTab[49912]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:177
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:177
	// _ = "end of CoverTab[49908]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:177
	_go_fuzz_dep_.CoverTab[49909]++
															panic(fmt.Sprintf("Token is not a TypeName: %s:%s", t.kind, NameKind(t.attrs&^hasSeparator)))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:178
	// _ = "end of CoverTab[49909]"
}

// FieldNumber returns the value for FieldNumber type. It returns a
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:181
// non-negative int32 value. Caller will still need to validate for the correct
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:181
// field number range.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:184
func (t Token) FieldNumber() int32 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:184
	_go_fuzz_dep_.CoverTab[49913]++
															if t.kind != Name || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:185
		_go_fuzz_dep_.CoverTab[49915]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:185
		return t.attrs&uint8(FieldNumber) == 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:185
		// _ = "end of CoverTab[49915]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:185
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:185
		_go_fuzz_dep_.CoverTab[49916]++
																panic(fmt.Sprintf("Token is not a FieldNumber: %s:%s", t.kind, NameKind(t.attrs&^hasSeparator)))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:186
		// _ = "end of CoverTab[49916]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:187
		_go_fuzz_dep_.CoverTab[49917]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:187
		// _ = "end of CoverTab[49917]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:187
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:187
	// _ = "end of CoverTab[49913]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:187
	_go_fuzz_dep_.CoverTab[49914]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:190
	num, _ := strconv.ParseInt(string(t.raw), 10, 32)
															return int32(num)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:191
	// _ = "end of CoverTab[49914]"
}

// String returns the string value for a Scalar type.
func (t Token) String() (string, bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:195
	_go_fuzz_dep_.CoverTab[49918]++
															if t.kind != Scalar || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:196
		_go_fuzz_dep_.CoverTab[49920]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:196
		return t.attrs != stringValue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:196
		// _ = "end of CoverTab[49920]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:196
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:196
		_go_fuzz_dep_.CoverTab[49921]++
																return "", false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:197
		// _ = "end of CoverTab[49921]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:198
		_go_fuzz_dep_.CoverTab[49922]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:198
		// _ = "end of CoverTab[49922]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:198
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:198
	// _ = "end of CoverTab[49918]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:198
	_go_fuzz_dep_.CoverTab[49919]++
															return t.str, true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:199
	// _ = "end of CoverTab[49919]"
}

// Enum returns the literal value for a Scalar type for use as enum literals.
func (t Token) Enum() (string, bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:203
	_go_fuzz_dep_.CoverTab[49923]++
															if t.kind != Scalar || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:204
		_go_fuzz_dep_.CoverTab[49925]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:204
		return t.attrs != literalValue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:204
		// _ = "end of CoverTab[49925]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:204
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:204
		_go_fuzz_dep_.CoverTab[49926]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:204
		return (len(t.raw) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:204
			_go_fuzz_dep_.CoverTab[49927]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:204
			return t.raw[0] == '-'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:204
			// _ = "end of CoverTab[49927]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:204
		}())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:204
		// _ = "end of CoverTab[49926]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:204
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:204
		_go_fuzz_dep_.CoverTab[49928]++
																return "", false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:205
		// _ = "end of CoverTab[49928]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:206
		_go_fuzz_dep_.CoverTab[49929]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:206
		// _ = "end of CoverTab[49929]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:206
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:206
	// _ = "end of CoverTab[49923]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:206
	_go_fuzz_dep_.CoverTab[49924]++
															return string(t.raw), true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:207
	// _ = "end of CoverTab[49924]"
}

// Bool returns the bool value for a Scalar type.
func (t Token) Bool() (bool, bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:211
	_go_fuzz_dep_.CoverTab[49930]++
															if t.kind != Scalar {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:212
		_go_fuzz_dep_.CoverTab[49933]++
																return false, false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:213
		// _ = "end of CoverTab[49933]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:214
		_go_fuzz_dep_.CoverTab[49934]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:214
		// _ = "end of CoverTab[49934]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:214
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:214
	// _ = "end of CoverTab[49930]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:214
	_go_fuzz_dep_.CoverTab[49931]++
															switch t.attrs {
	case literalValue:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:216
		_go_fuzz_dep_.CoverTab[49935]++
																if b, ok := boolLits[string(t.raw)]; ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:217
			_go_fuzz_dep_.CoverTab[49938]++
																	return b, true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:218
			// _ = "end of CoverTab[49938]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:219
			_go_fuzz_dep_.CoverTab[49939]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:219
			// _ = "end of CoverTab[49939]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:219
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:219
		// _ = "end of CoverTab[49935]"
	case numberValue:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:220
		_go_fuzz_dep_.CoverTab[49936]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:223
		n, err := strconv.ParseUint(t.str, 0, 64)
		if err == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:224
			_go_fuzz_dep_.CoverTab[49940]++
																	switch n {
			case 0:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:226
				_go_fuzz_dep_.CoverTab[49941]++
																		return false, true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:227
				// _ = "end of CoverTab[49941]"
			case 1:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:228
				_go_fuzz_dep_.CoverTab[49942]++
																		return true, true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:229
				// _ = "end of CoverTab[49942]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:229
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:229
				_go_fuzz_dep_.CoverTab[49943]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:229
				// _ = "end of CoverTab[49943]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:230
			// _ = "end of CoverTab[49940]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:231
			_go_fuzz_dep_.CoverTab[49944]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:231
			// _ = "end of CoverTab[49944]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:231
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:231
		// _ = "end of CoverTab[49936]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:231
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:231
		_go_fuzz_dep_.CoverTab[49937]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:231
		// _ = "end of CoverTab[49937]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:232
	// _ = "end of CoverTab[49931]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:232
	_go_fuzz_dep_.CoverTab[49932]++
															return false, false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:233
	// _ = "end of CoverTab[49932]"
}

// These exact boolean literals are the ones supported in C++.
var boolLits = map[string]bool{
	"t":		true,
	"true":		true,
	"True":		true,
	"f":		false,
	"false":	false,
	"False":	false,
}

// Uint64 returns the uint64 value for a Scalar type.
func (t Token) Uint64() (uint64, bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:247
	_go_fuzz_dep_.CoverTab[49945]++
															if t.kind != Scalar || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:248
		_go_fuzz_dep_.CoverTab[49948]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:248
		return t.attrs != numberValue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:248
		// _ = "end of CoverTab[49948]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:248
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:248
		_go_fuzz_dep_.CoverTab[49949]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:248
		return t.numAttrs&isNegative > 0
																// _ = "end of CoverTab[49949]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:249
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:249
		_go_fuzz_dep_.CoverTab[49950]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:249
		return t.numAttrs&numFloat > 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:249
		// _ = "end of CoverTab[49950]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:249
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:249
		_go_fuzz_dep_.CoverTab[49951]++
																return 0, false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:250
		// _ = "end of CoverTab[49951]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:251
		_go_fuzz_dep_.CoverTab[49952]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:251
		// _ = "end of CoverTab[49952]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:251
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:251
	// _ = "end of CoverTab[49945]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:251
	_go_fuzz_dep_.CoverTab[49946]++
															n, err := strconv.ParseUint(t.str, 0, 64)
															if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:253
		_go_fuzz_dep_.CoverTab[49953]++
																return 0, false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:254
		// _ = "end of CoverTab[49953]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:255
		_go_fuzz_dep_.CoverTab[49954]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:255
		// _ = "end of CoverTab[49954]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:255
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:255
	// _ = "end of CoverTab[49946]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:255
	_go_fuzz_dep_.CoverTab[49947]++
															return n, true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:256
	// _ = "end of CoverTab[49947]"
}

// Uint32 returns the uint32 value for a Scalar type.
func (t Token) Uint32() (uint32, bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:260
	_go_fuzz_dep_.CoverTab[49955]++
															if t.kind != Scalar || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:261
		_go_fuzz_dep_.CoverTab[49958]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:261
		return t.attrs != numberValue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:261
		// _ = "end of CoverTab[49958]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:261
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:261
		_go_fuzz_dep_.CoverTab[49959]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:261
		return t.numAttrs&isNegative > 0
																// _ = "end of CoverTab[49959]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:262
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:262
		_go_fuzz_dep_.CoverTab[49960]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:262
		return t.numAttrs&numFloat > 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:262
		// _ = "end of CoverTab[49960]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:262
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:262
		_go_fuzz_dep_.CoverTab[49961]++
																return 0, false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:263
		// _ = "end of CoverTab[49961]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:264
		_go_fuzz_dep_.CoverTab[49962]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:264
		// _ = "end of CoverTab[49962]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:264
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:264
	// _ = "end of CoverTab[49955]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:264
	_go_fuzz_dep_.CoverTab[49956]++
															n, err := strconv.ParseUint(t.str, 0, 32)
															if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:266
		_go_fuzz_dep_.CoverTab[49963]++
																return 0, false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:267
		// _ = "end of CoverTab[49963]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:268
		_go_fuzz_dep_.CoverTab[49964]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:268
		// _ = "end of CoverTab[49964]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:268
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:268
	// _ = "end of CoverTab[49956]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:268
	_go_fuzz_dep_.CoverTab[49957]++
															return uint32(n), true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:269
	// _ = "end of CoverTab[49957]"
}

// Int64 returns the int64 value for a Scalar type.
func (t Token) Int64() (int64, bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:273
	_go_fuzz_dep_.CoverTab[49965]++
															if t.kind != Scalar || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:274
		_go_fuzz_dep_.CoverTab[49969]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:274
		return t.attrs != numberValue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:274
		// _ = "end of CoverTab[49969]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:274
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:274
		_go_fuzz_dep_.CoverTab[49970]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:274
		return t.numAttrs&numFloat > 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:274
		// _ = "end of CoverTab[49970]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:274
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:274
		_go_fuzz_dep_.CoverTab[49971]++
																return 0, false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:275
		// _ = "end of CoverTab[49971]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:276
		_go_fuzz_dep_.CoverTab[49972]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:276
		// _ = "end of CoverTab[49972]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:276
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:276
	// _ = "end of CoverTab[49965]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:276
	_go_fuzz_dep_.CoverTab[49966]++
															if n, err := strconv.ParseInt(t.str, 0, 64); err == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:277
		_go_fuzz_dep_.CoverTab[49973]++
																return n, true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:278
		// _ = "end of CoverTab[49973]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:279
		_go_fuzz_dep_.CoverTab[49974]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:279
		// _ = "end of CoverTab[49974]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:279
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:279
	// _ = "end of CoverTab[49966]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:279
	_go_fuzz_dep_.CoverTab[49967]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:282
	if flags.ProtoLegacy && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:282
		_go_fuzz_dep_.CoverTab[49975]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:282
		return (t.numAttrs == numHex)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:282
		// _ = "end of CoverTab[49975]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:282
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:282
		_go_fuzz_dep_.CoverTab[49976]++
																if n, err := strconv.ParseUint(t.str, 0, 64); err == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:283
			_go_fuzz_dep_.CoverTab[49977]++
																	return int64(n), true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:284
			// _ = "end of CoverTab[49977]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:285
			_go_fuzz_dep_.CoverTab[49978]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:285
			// _ = "end of CoverTab[49978]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:285
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:285
		// _ = "end of CoverTab[49976]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:286
		_go_fuzz_dep_.CoverTab[49979]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:286
		// _ = "end of CoverTab[49979]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:286
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:286
	// _ = "end of CoverTab[49967]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:286
	_go_fuzz_dep_.CoverTab[49968]++
															return 0, false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:287
	// _ = "end of CoverTab[49968]"
}

// Int32 returns the int32 value for a Scalar type.
func (t Token) Int32() (int32, bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:291
	_go_fuzz_dep_.CoverTab[49980]++
															if t.kind != Scalar || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:292
		_go_fuzz_dep_.CoverTab[49984]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:292
		return t.attrs != numberValue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:292
		// _ = "end of CoverTab[49984]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:292
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:292
		_go_fuzz_dep_.CoverTab[49985]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:292
		return t.numAttrs&numFloat > 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:292
		// _ = "end of CoverTab[49985]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:292
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:292
		_go_fuzz_dep_.CoverTab[49986]++
																return 0, false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:293
		// _ = "end of CoverTab[49986]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:294
		_go_fuzz_dep_.CoverTab[49987]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:294
		// _ = "end of CoverTab[49987]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:294
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:294
	// _ = "end of CoverTab[49980]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:294
	_go_fuzz_dep_.CoverTab[49981]++
															if n, err := strconv.ParseInt(t.str, 0, 32); err == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:295
		_go_fuzz_dep_.CoverTab[49988]++
																return int32(n), true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:296
		// _ = "end of CoverTab[49988]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:297
		_go_fuzz_dep_.CoverTab[49989]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:297
		// _ = "end of CoverTab[49989]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:297
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:297
	// _ = "end of CoverTab[49981]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:297
	_go_fuzz_dep_.CoverTab[49982]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:300
	if flags.ProtoLegacy && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:300
		_go_fuzz_dep_.CoverTab[49990]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:300
		return (t.numAttrs == numHex)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:300
		// _ = "end of CoverTab[49990]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:300
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:300
		_go_fuzz_dep_.CoverTab[49991]++
																if n, err := strconv.ParseUint(t.str, 0, 32); err == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:301
			_go_fuzz_dep_.CoverTab[49992]++
																	return int32(n), true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:302
			// _ = "end of CoverTab[49992]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:303
			_go_fuzz_dep_.CoverTab[49993]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:303
			// _ = "end of CoverTab[49993]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:303
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:303
		// _ = "end of CoverTab[49991]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:304
		_go_fuzz_dep_.CoverTab[49994]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:304
		// _ = "end of CoverTab[49994]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:304
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:304
	// _ = "end of CoverTab[49982]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:304
	_go_fuzz_dep_.CoverTab[49983]++
															return 0, false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:305
	// _ = "end of CoverTab[49983]"
}

// Float64 returns the float64 value for a Scalar type.
func (t Token) Float64() (float64, bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:309
	_go_fuzz_dep_.CoverTab[49995]++
															if t.kind != Scalar {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:310
		_go_fuzz_dep_.CoverTab[49998]++
																return 0, false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:311
		// _ = "end of CoverTab[49998]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:312
		_go_fuzz_dep_.CoverTab[49999]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:312
		// _ = "end of CoverTab[49999]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:312
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:312
	// _ = "end of CoverTab[49995]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:312
	_go_fuzz_dep_.CoverTab[49996]++
															switch t.attrs {
	case literalValue:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:314
		_go_fuzz_dep_.CoverTab[50000]++
																if f, ok := floatLits[strings.ToLower(string(t.raw))]; ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:315
			_go_fuzz_dep_.CoverTab[50004]++
																	return f, true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:316
			// _ = "end of CoverTab[50004]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:317
			_go_fuzz_dep_.CoverTab[50005]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:317
			// _ = "end of CoverTab[50005]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:317
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:317
		// _ = "end of CoverTab[50000]"
	case numberValue:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:318
		_go_fuzz_dep_.CoverTab[50001]++
																n, err := strconv.ParseFloat(t.str, 64)
																if err == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:320
			_go_fuzz_dep_.CoverTab[50006]++
																	return n, true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:321
			// _ = "end of CoverTab[50006]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:322
			_go_fuzz_dep_.CoverTab[50007]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:322
			// _ = "end of CoverTab[50007]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:322
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:322
		// _ = "end of CoverTab[50001]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:322
		_go_fuzz_dep_.CoverTab[50002]++
																nerr := err.(*strconv.NumError)
																if nerr.Err == strconv.ErrRange {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:324
			_go_fuzz_dep_.CoverTab[50008]++
																	return n, true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:325
			// _ = "end of CoverTab[50008]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:326
			_go_fuzz_dep_.CoverTab[50009]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:326
			// _ = "end of CoverTab[50009]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:326
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:326
		// _ = "end of CoverTab[50002]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:326
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:326
		_go_fuzz_dep_.CoverTab[50003]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:326
		// _ = "end of CoverTab[50003]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:327
	// _ = "end of CoverTab[49996]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:327
	_go_fuzz_dep_.CoverTab[49997]++
															return 0, false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:328
	// _ = "end of CoverTab[49997]"
}

// Float32 returns the float32 value for a Scalar type.
func (t Token) Float32() (float32, bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:332
	_go_fuzz_dep_.CoverTab[50010]++
															if t.kind != Scalar {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:333
		_go_fuzz_dep_.CoverTab[50013]++
																return 0, false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:334
		// _ = "end of CoverTab[50013]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:335
		_go_fuzz_dep_.CoverTab[50014]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:335
		// _ = "end of CoverTab[50014]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:335
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:335
	// _ = "end of CoverTab[50010]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:335
	_go_fuzz_dep_.CoverTab[50011]++
															switch t.attrs {
	case literalValue:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:337
		_go_fuzz_dep_.CoverTab[50015]++
																if f, ok := floatLits[strings.ToLower(string(t.raw))]; ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:338
			_go_fuzz_dep_.CoverTab[50019]++
																	return float32(f), true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:339
			// _ = "end of CoverTab[50019]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:340
			_go_fuzz_dep_.CoverTab[50020]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:340
			// _ = "end of CoverTab[50020]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:340
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:340
		// _ = "end of CoverTab[50015]"
	case numberValue:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:341
		_go_fuzz_dep_.CoverTab[50016]++
																n, err := strconv.ParseFloat(t.str, 64)
																if err == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:343
			_go_fuzz_dep_.CoverTab[50021]++

																	return float32(n), true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:345
			// _ = "end of CoverTab[50021]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:346
			_go_fuzz_dep_.CoverTab[50022]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:346
			// _ = "end of CoverTab[50022]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:346
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:346
		// _ = "end of CoverTab[50016]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:346
		_go_fuzz_dep_.CoverTab[50017]++
																nerr := err.(*strconv.NumError)
																if nerr.Err == strconv.ErrRange {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:348
			_go_fuzz_dep_.CoverTab[50023]++
																	return float32(n), true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:349
			// _ = "end of CoverTab[50023]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:350
			_go_fuzz_dep_.CoverTab[50024]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:350
			// _ = "end of CoverTab[50024]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:350
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:350
		// _ = "end of CoverTab[50017]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:350
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:350
		_go_fuzz_dep_.CoverTab[50018]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:350
		// _ = "end of CoverTab[50018]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:351
	// _ = "end of CoverTab[50011]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:351
	_go_fuzz_dep_.CoverTab[50012]++
															return 0, false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:352
	// _ = "end of CoverTab[50012]"
}

// These are the supported float literals which C++ permits case-insensitive
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:355
// variants of these.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:357
var floatLits = map[string]float64{
	"nan":		math.NaN(),
	"inf":		math.Inf(1),
	"infinity":	math.Inf(1),
	"-inf":		math.Inf(-1),
	"-infinity":	math.Inf(-1),
}

// TokenEquals returns true if given Tokens are equal, else false.
func TokenEquals(x, y Token) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:366
	_go_fuzz_dep_.CoverTab[50025]++
															return x.kind == y.kind && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:367
		_go_fuzz_dep_.CoverTab[50026]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:367
		return x.attrs == y.attrs
																// _ = "end of CoverTab[50026]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:368
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:368
		_go_fuzz_dep_.CoverTab[50027]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:368
		return x.numAttrs == y.numAttrs
																// _ = "end of CoverTab[50027]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:369
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:369
		_go_fuzz_dep_.CoverTab[50028]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:369
		return x.pos == y.pos
																// _ = "end of CoverTab[50028]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:370
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:370
		_go_fuzz_dep_.CoverTab[50029]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:370
		return bytes.Equal(x.raw, y.raw)
																// _ = "end of CoverTab[50029]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:371
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:371
		_go_fuzz_dep_.CoverTab[50030]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:371
		return x.str == y.str
																// _ = "end of CoverTab[50030]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:372
	}()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:372
	// _ = "end of CoverTab[50025]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:373
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_token.go:373
var _ = _go_fuzz_dep_.CoverTab
