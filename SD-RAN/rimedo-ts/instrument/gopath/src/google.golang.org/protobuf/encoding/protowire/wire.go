// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:5
// Package protowire parses and formats the raw wire encoding.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:5
// See https://developers.google.com/protocol-buffers/docs/encoding.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:5
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:5
// For marshaling and unmarshaling entire protobuf messages,
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:5
// use the "google.golang.org/protobuf/proto" package instead.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:10
package protowire

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:10
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:10
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:10
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:10
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:10
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:10
)

import (
	"io"
	"math"
	"math/bits"

	"google.golang.org/protobuf/internal/errors"
)

// Number represents the field number.
type Number int32

const (
	MinValidNumber		Number	= 1
	FirstReservedNumber	Number	= 19000
	LastReservedNumber	Number	= 19999
	MaxValidNumber		Number	= 1<<29 - 1
	DefaultRecursionLimit		= 10000
)

// IsValid reports whether the field number is semantically valid.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:31
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:31
// Note that while numbers within the reserved range are semantically invalid,
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:31
// they are syntactically valid in the wire format.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:31
// Implementations may treat records with reserved field numbers as unknown.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:36
func (n Number) IsValid() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:36
	_go_fuzz_dep_.CoverTab[48327]++
													return MinValidNumber <= n && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:37
		_go_fuzz_dep_.CoverTab[48328]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:37
		return n < FirstReservedNumber
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:37
		// _ = "end of CoverTab[48328]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:37
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:37
		_go_fuzz_dep_.CoverTab[48329]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:37
		return LastReservedNumber < n && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:37
			_go_fuzz_dep_.CoverTab[48330]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:37
			return n <= MaxValidNumber
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:37
			// _ = "end of CoverTab[48330]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:37
		}()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:37
		// _ = "end of CoverTab[48329]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:37
	}()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:37
	// _ = "end of CoverTab[48327]"
}

// Type represents the wire type.
type Type int8

const (
	VarintType	Type	= 0
	Fixed32Type	Type	= 5
	Fixed64Type	Type	= 1
	BytesType	Type	= 2
	StartGroupType	Type	= 3
	EndGroupType	Type	= 4
)

const (
	_	= -iota
	errCodeTruncated
	errCodeFieldNumber
	errCodeOverflow
	errCodeReserved
	errCodeEndGroup
	errCodeRecursionDepth
)

var (
	errFieldNumber	= errors.New("invalid field number")
	errOverflow	= errors.New("variable length integer overflow")
	errReserved	= errors.New("cannot parse reserved wire type")
	errEndGroup	= errors.New("mismatching end group marker")
	errParse	= errors.New("parse error")
)

// ParseError converts an error code into an error value.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:70
// This returns nil if n is a non-negative number.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:72
func ParseError(n int) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:72
	_go_fuzz_dep_.CoverTab[48331]++
													if n >= 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:73
		_go_fuzz_dep_.CoverTab[48333]++
														return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:74
		// _ = "end of CoverTab[48333]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:75
		_go_fuzz_dep_.CoverTab[48334]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:75
		// _ = "end of CoverTab[48334]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:75
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:75
	// _ = "end of CoverTab[48331]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:75
	_go_fuzz_dep_.CoverTab[48332]++
													switch n {
	case errCodeTruncated:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:77
		_go_fuzz_dep_.CoverTab[48335]++
														return io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:78
		// _ = "end of CoverTab[48335]"
	case errCodeFieldNumber:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:79
		_go_fuzz_dep_.CoverTab[48336]++
														return errFieldNumber
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:80
		// _ = "end of CoverTab[48336]"
	case errCodeOverflow:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:81
		_go_fuzz_dep_.CoverTab[48337]++
														return errOverflow
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:82
		// _ = "end of CoverTab[48337]"
	case errCodeReserved:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:83
		_go_fuzz_dep_.CoverTab[48338]++
														return errReserved
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:84
		// _ = "end of CoverTab[48338]"
	case errCodeEndGroup:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:85
		_go_fuzz_dep_.CoverTab[48339]++
														return errEndGroup
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:86
		// _ = "end of CoverTab[48339]"
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:87
		_go_fuzz_dep_.CoverTab[48340]++
														return errParse
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:88
		// _ = "end of CoverTab[48340]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:89
	// _ = "end of CoverTab[48332]"
}

// ConsumeField parses an entire field record (both tag and value) and returns
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:92
// the field number, the wire type, and the total length.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:92
// This returns a negative length upon an error (see ParseError).
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:92
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:92
// The total length includes the tag header and the end group marker (if the
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:92
// field is a group).
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:98
func ConsumeField(b []byte) (Number, Type, int) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:98
	_go_fuzz_dep_.CoverTab[48341]++
													num, typ, n := ConsumeTag(b)
													if n < 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:100
		_go_fuzz_dep_.CoverTab[48344]++
														return 0, 0, n
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:101
		// _ = "end of CoverTab[48344]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:102
		_go_fuzz_dep_.CoverTab[48345]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:102
		// _ = "end of CoverTab[48345]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:102
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:102
	// _ = "end of CoverTab[48341]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:102
	_go_fuzz_dep_.CoverTab[48342]++
													m := ConsumeFieldValue(num, typ, b[n:])
													if m < 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:104
		_go_fuzz_dep_.CoverTab[48346]++
														return 0, 0, m
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:105
		// _ = "end of CoverTab[48346]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:106
		_go_fuzz_dep_.CoverTab[48347]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:106
		// _ = "end of CoverTab[48347]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:106
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:106
	// _ = "end of CoverTab[48342]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:106
	_go_fuzz_dep_.CoverTab[48343]++
													return num, typ, n + m
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:107
	// _ = "end of CoverTab[48343]"
}

// ConsumeFieldValue parses a field value and returns its length.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:110
// This assumes that the field Number and wire Type have already been parsed.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:110
// This returns a negative length upon an error (see ParseError).
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:110
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:110
// When parsing a group, the length includes the end group marker and
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:110
// the end group is verified to match the starting field number.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:116
func ConsumeFieldValue(num Number, typ Type, b []byte) (n int) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:116
	_go_fuzz_dep_.CoverTab[48348]++
													return consumeFieldValueD(num, typ, b, DefaultRecursionLimit)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:117
	// _ = "end of CoverTab[48348]"
}

func consumeFieldValueD(num Number, typ Type, b []byte, depth int) (n int) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:120
	_go_fuzz_dep_.CoverTab[48349]++
													switch typ {
	case VarintType:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:122
		_go_fuzz_dep_.CoverTab[48350]++
														_, n = ConsumeVarint(b)
														return n
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:124
		// _ = "end of CoverTab[48350]"
	case Fixed32Type:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:125
		_go_fuzz_dep_.CoverTab[48351]++
														_, n = ConsumeFixed32(b)
														return n
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:127
		// _ = "end of CoverTab[48351]"
	case Fixed64Type:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:128
		_go_fuzz_dep_.CoverTab[48352]++
														_, n = ConsumeFixed64(b)
														return n
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:130
		// _ = "end of CoverTab[48352]"
	case BytesType:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:131
		_go_fuzz_dep_.CoverTab[48353]++
														_, n = ConsumeBytes(b)
														return n
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:133
		// _ = "end of CoverTab[48353]"
	case StartGroupType:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:134
		_go_fuzz_dep_.CoverTab[48354]++
														if depth < 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:135
			_go_fuzz_dep_.CoverTab[48358]++
															return errCodeRecursionDepth
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:136
			// _ = "end of CoverTab[48358]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:137
			_go_fuzz_dep_.CoverTab[48359]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:137
			// _ = "end of CoverTab[48359]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:137
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:137
		// _ = "end of CoverTab[48354]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:137
		_go_fuzz_dep_.CoverTab[48355]++
														n0 := len(b)
														for {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:139
			_go_fuzz_dep_.CoverTab[48360]++
															num2, typ2, n := ConsumeTag(b)
															if n < 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:141
				_go_fuzz_dep_.CoverTab[48364]++
																return n
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:142
				// _ = "end of CoverTab[48364]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:143
				_go_fuzz_dep_.CoverTab[48365]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:143
				// _ = "end of CoverTab[48365]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:143
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:143
			// _ = "end of CoverTab[48360]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:143
			_go_fuzz_dep_.CoverTab[48361]++
															b = b[n:]
															if typ2 == EndGroupType {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:145
				_go_fuzz_dep_.CoverTab[48366]++
																if num != num2 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:146
					_go_fuzz_dep_.CoverTab[48368]++
																	return errCodeEndGroup
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:147
					// _ = "end of CoverTab[48368]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:148
					_go_fuzz_dep_.CoverTab[48369]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:148
					// _ = "end of CoverTab[48369]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:148
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:148
				// _ = "end of CoverTab[48366]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:148
				_go_fuzz_dep_.CoverTab[48367]++
																return n0 - len(b)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:149
				// _ = "end of CoverTab[48367]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:150
				_go_fuzz_dep_.CoverTab[48370]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:150
				// _ = "end of CoverTab[48370]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:150
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:150
			// _ = "end of CoverTab[48361]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:150
			_go_fuzz_dep_.CoverTab[48362]++

															n = consumeFieldValueD(num2, typ2, b, depth-1)
															if n < 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:153
				_go_fuzz_dep_.CoverTab[48371]++
																return n
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:154
				// _ = "end of CoverTab[48371]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:155
				_go_fuzz_dep_.CoverTab[48372]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:155
				// _ = "end of CoverTab[48372]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:155
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:155
			// _ = "end of CoverTab[48362]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:155
			_go_fuzz_dep_.CoverTab[48363]++
															b = b[n:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:156
			// _ = "end of CoverTab[48363]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:157
		// _ = "end of CoverTab[48355]"
	case EndGroupType:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:158
		_go_fuzz_dep_.CoverTab[48356]++
														return errCodeEndGroup
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:159
		// _ = "end of CoverTab[48356]"
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:160
		_go_fuzz_dep_.CoverTab[48357]++
														return errCodeReserved
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:161
		// _ = "end of CoverTab[48357]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:162
	// _ = "end of CoverTab[48349]"
}

// AppendTag encodes num and typ as a varint-encoded tag and appends it to b.
func AppendTag(b []byte, num Number, typ Type) []byte {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:166
	_go_fuzz_dep_.CoverTab[48373]++
													return AppendVarint(b, EncodeTag(num, typ))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:167
	// _ = "end of CoverTab[48373]"
}

// ConsumeTag parses b as a varint-encoded tag, reporting its length.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:170
// This returns a negative length upon an error (see ParseError).
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:172
func ConsumeTag(b []byte) (Number, Type, int) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:172
	_go_fuzz_dep_.CoverTab[48374]++
													v, n := ConsumeVarint(b)
													if n < 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:174
		_go_fuzz_dep_.CoverTab[48377]++
														return 0, 0, n
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:175
		// _ = "end of CoverTab[48377]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:176
		_go_fuzz_dep_.CoverTab[48378]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:176
		// _ = "end of CoverTab[48378]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:176
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:176
	// _ = "end of CoverTab[48374]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:176
	_go_fuzz_dep_.CoverTab[48375]++
													num, typ := DecodeTag(v)
													if num < MinValidNumber {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:178
		_go_fuzz_dep_.CoverTab[48379]++
														return 0, 0, errCodeFieldNumber
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:179
		// _ = "end of CoverTab[48379]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:180
		_go_fuzz_dep_.CoverTab[48380]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:180
		// _ = "end of CoverTab[48380]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:180
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:180
	// _ = "end of CoverTab[48375]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:180
	_go_fuzz_dep_.CoverTab[48376]++
													return num, typ, n
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:181
	// _ = "end of CoverTab[48376]"
}

func SizeTag(num Number) int {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:184
	_go_fuzz_dep_.CoverTab[48381]++
													return SizeVarint(EncodeTag(num, 0))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:185
	// _ = "end of CoverTab[48381]"
}

// AppendVarint appends v to b as a varint-encoded uint64.
func AppendVarint(b []byte, v uint64) []byte {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:189
	_go_fuzz_dep_.CoverTab[48382]++
													switch {
	case v < 1<<7:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:191
		_go_fuzz_dep_.CoverTab[48384]++
														b = append(b, byte(v))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:192
		// _ = "end of CoverTab[48384]"
	case v < 1<<14:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:193
		_go_fuzz_dep_.CoverTab[48385]++
														b = append(b,
			byte((v>>0)&0x7f|0x80),
			byte(v>>7))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:196
		// _ = "end of CoverTab[48385]"
	case v < 1<<21:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:197
		_go_fuzz_dep_.CoverTab[48386]++
														b = append(b,
			byte((v>>0)&0x7f|0x80),
			byte((v>>7)&0x7f|0x80),
			byte(v>>14))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:201
		// _ = "end of CoverTab[48386]"
	case v < 1<<28:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:202
		_go_fuzz_dep_.CoverTab[48387]++
														b = append(b,
			byte((v>>0)&0x7f|0x80),
			byte((v>>7)&0x7f|0x80),
			byte((v>>14)&0x7f|0x80),
			byte(v>>21))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:207
		// _ = "end of CoverTab[48387]"
	case v < 1<<35:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:208
		_go_fuzz_dep_.CoverTab[48388]++
														b = append(b,
			byte((v>>0)&0x7f|0x80),
			byte((v>>7)&0x7f|0x80),
			byte((v>>14)&0x7f|0x80),
			byte((v>>21)&0x7f|0x80),
			byte(v>>28))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:214
		// _ = "end of CoverTab[48388]"
	case v < 1<<42:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:215
		_go_fuzz_dep_.CoverTab[48389]++
														b = append(b,
			byte((v>>0)&0x7f|0x80),
			byte((v>>7)&0x7f|0x80),
			byte((v>>14)&0x7f|0x80),
			byte((v>>21)&0x7f|0x80),
			byte((v>>28)&0x7f|0x80),
			byte(v>>35))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:222
		// _ = "end of CoverTab[48389]"
	case v < 1<<49:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:223
		_go_fuzz_dep_.CoverTab[48390]++
														b = append(b,
			byte((v>>0)&0x7f|0x80),
			byte((v>>7)&0x7f|0x80),
			byte((v>>14)&0x7f|0x80),
			byte((v>>21)&0x7f|0x80),
			byte((v>>28)&0x7f|0x80),
			byte((v>>35)&0x7f|0x80),
			byte(v>>42))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:231
		// _ = "end of CoverTab[48390]"
	case v < 1<<56:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:232
		_go_fuzz_dep_.CoverTab[48391]++
														b = append(b,
			byte((v>>0)&0x7f|0x80),
			byte((v>>7)&0x7f|0x80),
			byte((v>>14)&0x7f|0x80),
			byte((v>>21)&0x7f|0x80),
			byte((v>>28)&0x7f|0x80),
			byte((v>>35)&0x7f|0x80),
			byte((v>>42)&0x7f|0x80),
			byte(v>>49))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:241
		// _ = "end of CoverTab[48391]"
	case v < 1<<63:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:242
		_go_fuzz_dep_.CoverTab[48392]++
														b = append(b,
			byte((v>>0)&0x7f|0x80),
			byte((v>>7)&0x7f|0x80),
			byte((v>>14)&0x7f|0x80),
			byte((v>>21)&0x7f|0x80),
			byte((v>>28)&0x7f|0x80),
			byte((v>>35)&0x7f|0x80),
			byte((v>>42)&0x7f|0x80),
			byte((v>>49)&0x7f|0x80),
			byte(v>>56))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:252
		// _ = "end of CoverTab[48392]"
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:253
		_go_fuzz_dep_.CoverTab[48393]++
														b = append(b,
			byte((v>>0)&0x7f|0x80),
			byte((v>>7)&0x7f|0x80),
			byte((v>>14)&0x7f|0x80),
			byte((v>>21)&0x7f|0x80),
			byte((v>>28)&0x7f|0x80),
			byte((v>>35)&0x7f|0x80),
			byte((v>>42)&0x7f|0x80),
			byte((v>>49)&0x7f|0x80),
			byte((v>>56)&0x7f|0x80),
			1)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:264
		// _ = "end of CoverTab[48393]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:265
	// _ = "end of CoverTab[48382]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:265
	_go_fuzz_dep_.CoverTab[48383]++
													return b
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:266
	// _ = "end of CoverTab[48383]"
}

// ConsumeVarint parses b as a varint-encoded uint64, reporting its length.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:269
// This returns a negative length upon an error (see ParseError).
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:271
func ConsumeVarint(b []byte) (v uint64, n int) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:271
	_go_fuzz_dep_.CoverTab[48394]++
													var y uint64
													if len(b) <= 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:273
		_go_fuzz_dep_.CoverTab[48415]++
														return 0, errCodeTruncated
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:274
		// _ = "end of CoverTab[48415]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:275
		_go_fuzz_dep_.CoverTab[48416]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:275
		// _ = "end of CoverTab[48416]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:275
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:275
	// _ = "end of CoverTab[48394]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:275
	_go_fuzz_dep_.CoverTab[48395]++
													v = uint64(b[0])
													if v < 0x80 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:277
		_go_fuzz_dep_.CoverTab[48417]++
														return v, 1
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:278
		// _ = "end of CoverTab[48417]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:279
		_go_fuzz_dep_.CoverTab[48418]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:279
		// _ = "end of CoverTab[48418]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:279
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:279
	// _ = "end of CoverTab[48395]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:279
	_go_fuzz_dep_.CoverTab[48396]++
													v -= 0x80

													if len(b) <= 1 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:282
		_go_fuzz_dep_.CoverTab[48419]++
														return 0, errCodeTruncated
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:283
		// _ = "end of CoverTab[48419]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:284
		_go_fuzz_dep_.CoverTab[48420]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:284
		// _ = "end of CoverTab[48420]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:284
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:284
	// _ = "end of CoverTab[48396]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:284
	_go_fuzz_dep_.CoverTab[48397]++
													y = uint64(b[1])
													v += y << 7
													if y < 0x80 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:287
		_go_fuzz_dep_.CoverTab[48421]++
														return v, 2
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:288
		// _ = "end of CoverTab[48421]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:289
		_go_fuzz_dep_.CoverTab[48422]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:289
		// _ = "end of CoverTab[48422]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:289
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:289
	// _ = "end of CoverTab[48397]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:289
	_go_fuzz_dep_.CoverTab[48398]++
													v -= 0x80 << 7

													if len(b) <= 2 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:292
		_go_fuzz_dep_.CoverTab[48423]++
														return 0, errCodeTruncated
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:293
		// _ = "end of CoverTab[48423]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:294
		_go_fuzz_dep_.CoverTab[48424]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:294
		// _ = "end of CoverTab[48424]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:294
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:294
	// _ = "end of CoverTab[48398]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:294
	_go_fuzz_dep_.CoverTab[48399]++
													y = uint64(b[2])
													v += y << 14
													if y < 0x80 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:297
		_go_fuzz_dep_.CoverTab[48425]++
														return v, 3
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:298
		// _ = "end of CoverTab[48425]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:299
		_go_fuzz_dep_.CoverTab[48426]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:299
		// _ = "end of CoverTab[48426]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:299
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:299
	// _ = "end of CoverTab[48399]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:299
	_go_fuzz_dep_.CoverTab[48400]++
													v -= 0x80 << 14

													if len(b) <= 3 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:302
		_go_fuzz_dep_.CoverTab[48427]++
														return 0, errCodeTruncated
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:303
		// _ = "end of CoverTab[48427]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:304
		_go_fuzz_dep_.CoverTab[48428]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:304
		// _ = "end of CoverTab[48428]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:304
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:304
	// _ = "end of CoverTab[48400]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:304
	_go_fuzz_dep_.CoverTab[48401]++
													y = uint64(b[3])
													v += y << 21
													if y < 0x80 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:307
		_go_fuzz_dep_.CoverTab[48429]++
														return v, 4
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:308
		// _ = "end of CoverTab[48429]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:309
		_go_fuzz_dep_.CoverTab[48430]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:309
		// _ = "end of CoverTab[48430]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:309
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:309
	// _ = "end of CoverTab[48401]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:309
	_go_fuzz_dep_.CoverTab[48402]++
													v -= 0x80 << 21

													if len(b) <= 4 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:312
		_go_fuzz_dep_.CoverTab[48431]++
														return 0, errCodeTruncated
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:313
		// _ = "end of CoverTab[48431]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:314
		_go_fuzz_dep_.CoverTab[48432]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:314
		// _ = "end of CoverTab[48432]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:314
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:314
	// _ = "end of CoverTab[48402]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:314
	_go_fuzz_dep_.CoverTab[48403]++
													y = uint64(b[4])
													v += y << 28
													if y < 0x80 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:317
		_go_fuzz_dep_.CoverTab[48433]++
														return v, 5
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:318
		// _ = "end of CoverTab[48433]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:319
		_go_fuzz_dep_.CoverTab[48434]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:319
		// _ = "end of CoverTab[48434]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:319
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:319
	// _ = "end of CoverTab[48403]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:319
	_go_fuzz_dep_.CoverTab[48404]++
													v -= 0x80 << 28

													if len(b) <= 5 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:322
		_go_fuzz_dep_.CoverTab[48435]++
														return 0, errCodeTruncated
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:323
		// _ = "end of CoverTab[48435]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:324
		_go_fuzz_dep_.CoverTab[48436]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:324
		// _ = "end of CoverTab[48436]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:324
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:324
	// _ = "end of CoverTab[48404]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:324
	_go_fuzz_dep_.CoverTab[48405]++
													y = uint64(b[5])
													v += y << 35
													if y < 0x80 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:327
		_go_fuzz_dep_.CoverTab[48437]++
														return v, 6
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:328
		// _ = "end of CoverTab[48437]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:329
		_go_fuzz_dep_.CoverTab[48438]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:329
		// _ = "end of CoverTab[48438]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:329
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:329
	// _ = "end of CoverTab[48405]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:329
	_go_fuzz_dep_.CoverTab[48406]++
													v -= 0x80 << 35

													if len(b) <= 6 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:332
		_go_fuzz_dep_.CoverTab[48439]++
														return 0, errCodeTruncated
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:333
		// _ = "end of CoverTab[48439]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:334
		_go_fuzz_dep_.CoverTab[48440]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:334
		// _ = "end of CoverTab[48440]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:334
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:334
	// _ = "end of CoverTab[48406]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:334
	_go_fuzz_dep_.CoverTab[48407]++
													y = uint64(b[6])
													v += y << 42
													if y < 0x80 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:337
		_go_fuzz_dep_.CoverTab[48441]++
														return v, 7
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:338
		// _ = "end of CoverTab[48441]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:339
		_go_fuzz_dep_.CoverTab[48442]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:339
		// _ = "end of CoverTab[48442]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:339
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:339
	// _ = "end of CoverTab[48407]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:339
	_go_fuzz_dep_.CoverTab[48408]++
													v -= 0x80 << 42

													if len(b) <= 7 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:342
		_go_fuzz_dep_.CoverTab[48443]++
														return 0, errCodeTruncated
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:343
		// _ = "end of CoverTab[48443]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:344
		_go_fuzz_dep_.CoverTab[48444]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:344
		// _ = "end of CoverTab[48444]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:344
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:344
	// _ = "end of CoverTab[48408]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:344
	_go_fuzz_dep_.CoverTab[48409]++
													y = uint64(b[7])
													v += y << 49
													if y < 0x80 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:347
		_go_fuzz_dep_.CoverTab[48445]++
														return v, 8
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:348
		// _ = "end of CoverTab[48445]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:349
		_go_fuzz_dep_.CoverTab[48446]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:349
		// _ = "end of CoverTab[48446]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:349
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:349
	// _ = "end of CoverTab[48409]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:349
	_go_fuzz_dep_.CoverTab[48410]++
													v -= 0x80 << 49

													if len(b) <= 8 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:352
		_go_fuzz_dep_.CoverTab[48447]++
														return 0, errCodeTruncated
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:353
		// _ = "end of CoverTab[48447]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:354
		_go_fuzz_dep_.CoverTab[48448]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:354
		// _ = "end of CoverTab[48448]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:354
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:354
	// _ = "end of CoverTab[48410]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:354
	_go_fuzz_dep_.CoverTab[48411]++
													y = uint64(b[8])
													v += y << 56
													if y < 0x80 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:357
		_go_fuzz_dep_.CoverTab[48449]++
														return v, 9
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:358
		// _ = "end of CoverTab[48449]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:359
		_go_fuzz_dep_.CoverTab[48450]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:359
		// _ = "end of CoverTab[48450]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:359
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:359
	// _ = "end of CoverTab[48411]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:359
	_go_fuzz_dep_.CoverTab[48412]++
													v -= 0x80 << 56

													if len(b) <= 9 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:362
		_go_fuzz_dep_.CoverTab[48451]++
														return 0, errCodeTruncated
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:363
		// _ = "end of CoverTab[48451]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:364
		_go_fuzz_dep_.CoverTab[48452]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:364
		// _ = "end of CoverTab[48452]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:364
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:364
	// _ = "end of CoverTab[48412]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:364
	_go_fuzz_dep_.CoverTab[48413]++
													y = uint64(b[9])
													v += y << 63
													if y < 2 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:367
		_go_fuzz_dep_.CoverTab[48453]++
														return v, 10
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:368
		// _ = "end of CoverTab[48453]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:369
		_go_fuzz_dep_.CoverTab[48454]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:369
		// _ = "end of CoverTab[48454]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:369
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:369
	// _ = "end of CoverTab[48413]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:369
	_go_fuzz_dep_.CoverTab[48414]++
													return 0, errCodeOverflow
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:370
	// _ = "end of CoverTab[48414]"
}

// SizeVarint returns the encoded size of a varint.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:373
// The size is guaranteed to be within 1 and 10, inclusive.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:375
func SizeVarint(v uint64) int {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:375
	_go_fuzz_dep_.CoverTab[48455]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:378
	return int(9*uint32(bits.Len64(v))+64) / 64
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:378
	// _ = "end of CoverTab[48455]"
}

// AppendFixed32 appends v to b as a little-endian uint32.
func AppendFixed32(b []byte, v uint32) []byte {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:382
	_go_fuzz_dep_.CoverTab[48456]++
													return append(b,
		byte(v>>0),
		byte(v>>8),
		byte(v>>16),
		byte(v>>24))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:387
	// _ = "end of CoverTab[48456]"
}

// ConsumeFixed32 parses b as a little-endian uint32, reporting its length.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:390
// This returns a negative length upon an error (see ParseError).
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:392
func ConsumeFixed32(b []byte) (v uint32, n int) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:392
	_go_fuzz_dep_.CoverTab[48457]++
													if len(b) < 4 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:393
		_go_fuzz_dep_.CoverTab[48459]++
														return 0, errCodeTruncated
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:394
		// _ = "end of CoverTab[48459]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:395
		_go_fuzz_dep_.CoverTab[48460]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:395
		// _ = "end of CoverTab[48460]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:395
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:395
	// _ = "end of CoverTab[48457]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:395
	_go_fuzz_dep_.CoverTab[48458]++
													v = uint32(b[0])<<0 | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24
													return v, 4
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:397
	// _ = "end of CoverTab[48458]"
}

// SizeFixed32 returns the encoded size of a fixed32; which is always 4.
func SizeFixed32() int {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:401
	_go_fuzz_dep_.CoverTab[48461]++
													return 4
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:402
	// _ = "end of CoverTab[48461]"
}

// AppendFixed64 appends v to b as a little-endian uint64.
func AppendFixed64(b []byte, v uint64) []byte {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:406
	_go_fuzz_dep_.CoverTab[48462]++
													return append(b,
		byte(v>>0),
		byte(v>>8),
		byte(v>>16),
		byte(v>>24),
		byte(v>>32),
		byte(v>>40),
		byte(v>>48),
		byte(v>>56))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:415
	// _ = "end of CoverTab[48462]"
}

// ConsumeFixed64 parses b as a little-endian uint64, reporting its length.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:418
// This returns a negative length upon an error (see ParseError).
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:420
func ConsumeFixed64(b []byte) (v uint64, n int) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:420
	_go_fuzz_dep_.CoverTab[48463]++
													if len(b) < 8 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:421
		_go_fuzz_dep_.CoverTab[48465]++
														return 0, errCodeTruncated
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:422
		// _ = "end of CoverTab[48465]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:423
		_go_fuzz_dep_.CoverTab[48466]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:423
		// _ = "end of CoverTab[48466]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:423
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:423
	// _ = "end of CoverTab[48463]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:423
	_go_fuzz_dep_.CoverTab[48464]++
													v = uint64(b[0])<<0 | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 | uint64(b[4])<<32 | uint64(b[5])<<40 | uint64(b[6])<<48 | uint64(b[7])<<56
													return v, 8
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:425
	// _ = "end of CoverTab[48464]"
}

// SizeFixed64 returns the encoded size of a fixed64; which is always 8.
func SizeFixed64() int {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:429
	_go_fuzz_dep_.CoverTab[48467]++
													return 8
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:430
	// _ = "end of CoverTab[48467]"
}

// AppendBytes appends v to b as a length-prefixed bytes value.
func AppendBytes(b []byte, v []byte) []byte {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:434
	_go_fuzz_dep_.CoverTab[48468]++
													return append(AppendVarint(b, uint64(len(v))), v...)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:435
	// _ = "end of CoverTab[48468]"
}

// ConsumeBytes parses b as a length-prefixed bytes value, reporting its length.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:438
// This returns a negative length upon an error (see ParseError).
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:440
func ConsumeBytes(b []byte) (v []byte, n int) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:440
	_go_fuzz_dep_.CoverTab[48469]++
													m, n := ConsumeVarint(b)
													if n < 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:442
		_go_fuzz_dep_.CoverTab[48472]++
														return nil, n
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:443
		// _ = "end of CoverTab[48472]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:444
		_go_fuzz_dep_.CoverTab[48473]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:444
		// _ = "end of CoverTab[48473]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:444
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:444
	// _ = "end of CoverTab[48469]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:444
	_go_fuzz_dep_.CoverTab[48470]++
													if m > uint64(len(b[n:])) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:445
		_go_fuzz_dep_.CoverTab[48474]++
														return nil, errCodeTruncated
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:446
		// _ = "end of CoverTab[48474]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:447
		_go_fuzz_dep_.CoverTab[48475]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:447
		// _ = "end of CoverTab[48475]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:447
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:447
	// _ = "end of CoverTab[48470]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:447
	_go_fuzz_dep_.CoverTab[48471]++
													return b[n:][:m], n + int(m)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:448
	// _ = "end of CoverTab[48471]"
}

// SizeBytes returns the encoded size of a length-prefixed bytes value,
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:451
// given only the length.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:453
func SizeBytes(n int) int {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:453
	_go_fuzz_dep_.CoverTab[48476]++
													return SizeVarint(uint64(n)) + n
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:454
	// _ = "end of CoverTab[48476]"
}

// AppendString appends v to b as a length-prefixed bytes value.
func AppendString(b []byte, v string) []byte {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:458
	_go_fuzz_dep_.CoverTab[48477]++
													return append(AppendVarint(b, uint64(len(v))), v...)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:459
	// _ = "end of CoverTab[48477]"
}

// ConsumeString parses b as a length-prefixed bytes value, reporting its length.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:462
// This returns a negative length upon an error (see ParseError).
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:464
func ConsumeString(b []byte) (v string, n int) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:464
	_go_fuzz_dep_.CoverTab[48478]++
													bb, n := ConsumeBytes(b)
													return string(bb), n
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:466
	// _ = "end of CoverTab[48478]"
}

// AppendGroup appends v to b as group value, with a trailing end group marker.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:469
// The value v must not contain the end marker.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:471
func AppendGroup(b []byte, num Number, v []byte) []byte {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:471
	_go_fuzz_dep_.CoverTab[48479]++
													return AppendVarint(append(b, v...), EncodeTag(num, EndGroupType))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:472
	// _ = "end of CoverTab[48479]"
}

// ConsumeGroup parses b as a group value until the trailing end group marker,
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:475
// and verifies that the end marker matches the provided num. The value v
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:475
// does not contain the end marker, while the length does contain the end marker.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:475
// This returns a negative length upon an error (see ParseError).
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:479
func ConsumeGroup(num Number, b []byte) (v []byte, n int) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:479
	_go_fuzz_dep_.CoverTab[48480]++
													n = ConsumeFieldValue(num, StartGroupType, b)
													if n < 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:481
		_go_fuzz_dep_.CoverTab[48483]++
														return nil, n
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:482
		// _ = "end of CoverTab[48483]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:483
		_go_fuzz_dep_.CoverTab[48484]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:483
		// _ = "end of CoverTab[48484]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:483
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:483
	// _ = "end of CoverTab[48480]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:483
	_go_fuzz_dep_.CoverTab[48481]++
													b = b[:n]

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:490
	for len(b) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:490
		_go_fuzz_dep_.CoverTab[48485]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:490
		return b[len(b)-1]&0x7f == 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:490
		// _ = "end of CoverTab[48485]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:490
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:490
		_go_fuzz_dep_.CoverTab[48486]++
														b = b[:len(b)-1]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:491
		// _ = "end of CoverTab[48486]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:492
	// _ = "end of CoverTab[48481]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:492
	_go_fuzz_dep_.CoverTab[48482]++
													b = b[:len(b)-SizeTag(num)]
													return b, n
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:494
	// _ = "end of CoverTab[48482]"
}

// SizeGroup returns the encoded size of a group, given only the length.
func SizeGroup(num Number, n int) int {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:498
	_go_fuzz_dep_.CoverTab[48487]++
													return n + SizeTag(num)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:499
	// _ = "end of CoverTab[48487]"
}

// DecodeTag decodes the field Number and wire Type from its unified form.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:502
// The Number is -1 if the decoded field number overflows int32.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:502
// Other than overflow, this does not check for field number validity.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:505
func DecodeTag(x uint64) (Number, Type) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:505
	_go_fuzz_dep_.CoverTab[48488]++

													if x>>3 > uint64(math.MaxInt32) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:507
		_go_fuzz_dep_.CoverTab[48490]++
														return -1, 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:508
		// _ = "end of CoverTab[48490]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:509
		_go_fuzz_dep_.CoverTab[48491]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:509
		// _ = "end of CoverTab[48491]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:509
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:509
	// _ = "end of CoverTab[48488]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:509
	_go_fuzz_dep_.CoverTab[48489]++
													return Number(x >> 3), Type(x & 7)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:510
	// _ = "end of CoverTab[48489]"
}

// EncodeTag encodes the field Number and wire Type into its unified form.
func EncodeTag(num Number, typ Type) uint64 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:514
	_go_fuzz_dep_.CoverTab[48492]++
													return uint64(num)<<3 | uint64(typ&7)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:515
	// _ = "end of CoverTab[48492]"
}

// DecodeZigZag decodes a zig-zag-encoded uint64 as an int64.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:518
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:518
//	Input:  {…,  5,  3,  1,  0,  2,  4,  6, …}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:518
//	Output: {…, -3, -2, -1,  0, +1, +2, +3, …}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:522
func DecodeZigZag(x uint64) int64 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:522
	_go_fuzz_dep_.CoverTab[48493]++
													return int64(x>>1) ^ int64(x)<<63>>63
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:523
	// _ = "end of CoverTab[48493]"
}

// EncodeZigZag encodes an int64 as a zig-zag-encoded uint64.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:526
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:526
//	Input:  {…, -3, -2, -1,  0, +1, +2, +3, …}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:526
//	Output: {…,  5,  3,  1,  0,  2,  4,  6, …}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:530
func EncodeZigZag(x int64) uint64 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:530
	_go_fuzz_dep_.CoverTab[48494]++
													return uint64(x<<1) ^ uint64(x>>63)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:531
	// _ = "end of CoverTab[48494]"
}

// DecodeBool decodes a uint64 as a bool.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:534
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:534
//	Input:  {    0,    1,    2, …}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:534
//	Output: {false, true, true, …}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:538
func DecodeBool(x uint64) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:538
	_go_fuzz_dep_.CoverTab[48495]++
													return x != 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:539
	// _ = "end of CoverTab[48495]"
}

// EncodeBool encodes a bool as a uint64.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:542
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:542
//	Input:  {false, true}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:542
//	Output: {    0,    1}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:546
func EncodeBool(x bool) uint64 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:546
	_go_fuzz_dep_.CoverTab[48496]++
													if x {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:547
		_go_fuzz_dep_.CoverTab[48498]++
														return 1
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:548
		// _ = "end of CoverTab[48498]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:549
		_go_fuzz_dep_.CoverTab[48499]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:549
		// _ = "end of CoverTab[48499]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:549
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:549
	// _ = "end of CoverTab[48496]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:549
	_go_fuzz_dep_.CoverTab[48497]++
													return 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:550
	// _ = "end of CoverTab[48497]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:551
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protowire/wire.go:551
var _ = _go_fuzz_dep_.CoverTab
