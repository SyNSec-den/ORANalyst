// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/encoding/asn1/asn1.go:5
// Package asn1 implements parsing of DER-encoded ASN.1 data structures,
//line /usr/local/go/src/encoding/asn1/asn1.go:5
// as defined in ITU-T Rec X.690.
//line /usr/local/go/src/encoding/asn1/asn1.go:5
//
//line /usr/local/go/src/encoding/asn1/asn1.go:5
// See also “A Layman's Guide to a Subset of ASN.1, BER, and DER,”
//line /usr/local/go/src/encoding/asn1/asn1.go:5
// http://luca.ntop.org/Teaching/Appunti/asn1.html.
//line /usr/local/go/src/encoding/asn1/asn1.go:10
package asn1

//line /usr/local/go/src/encoding/asn1/asn1.go:10
import (
//line /usr/local/go/src/encoding/asn1/asn1.go:10
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/encoding/asn1/asn1.go:10
)
//line /usr/local/go/src/encoding/asn1/asn1.go:10
import (
//line /usr/local/go/src/encoding/asn1/asn1.go:10
	_atomic_ "sync/atomic"
//line /usr/local/go/src/encoding/asn1/asn1.go:10
)

//line /usr/local/go/src/encoding/asn1/asn1.go:22
import (
	"errors"
	"fmt"
	"math"
	"math/big"
	"reflect"
	"strconv"
	"time"
	"unicode/utf16"
	"unicode/utf8"
)

// A StructuralError suggests that the ASN.1 data is valid, but the Go type
//line /usr/local/go/src/encoding/asn1/asn1.go:34
// which is receiving it doesn't match.
//line /usr/local/go/src/encoding/asn1/asn1.go:36
type StructuralError struct {
	Msg string
}

func (e StructuralError) Error() string {
//line /usr/local/go/src/encoding/asn1/asn1.go:40
	_go_fuzz_dep_.CoverTab[7486]++
//line /usr/local/go/src/encoding/asn1/asn1.go:40
	return "asn1: structure error: " + e.Msg
//line /usr/local/go/src/encoding/asn1/asn1.go:40
	// _ = "end of CoverTab[7486]"
//line /usr/local/go/src/encoding/asn1/asn1.go:40
}

// A SyntaxError suggests that the ASN.1 data is invalid.
type SyntaxError struct {
	Msg string
}

func (e SyntaxError) Error() string {
//line /usr/local/go/src/encoding/asn1/asn1.go:47
	_go_fuzz_dep_.CoverTab[7487]++
//line /usr/local/go/src/encoding/asn1/asn1.go:47
	return "asn1: syntax error: " + e.Msg
//line /usr/local/go/src/encoding/asn1/asn1.go:47
	// _ = "end of CoverTab[7487]"
//line /usr/local/go/src/encoding/asn1/asn1.go:47
}

//line /usr/local/go/src/encoding/asn1/asn1.go:53
func parseBool(bytes []byte) (ret bool, err error) {
//line /usr/local/go/src/encoding/asn1/asn1.go:53
	_go_fuzz_dep_.CoverTab[7488]++
							if len(bytes) != 1 {
//line /usr/local/go/src/encoding/asn1/asn1.go:54
		_go_fuzz_dep_.CoverTab[7491]++
								err = SyntaxError{"invalid boolean"}
								return
//line /usr/local/go/src/encoding/asn1/asn1.go:56
		// _ = "end of CoverTab[7491]"
	} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:57
		_go_fuzz_dep_.CoverTab[7492]++
//line /usr/local/go/src/encoding/asn1/asn1.go:57
		// _ = "end of CoverTab[7492]"
//line /usr/local/go/src/encoding/asn1/asn1.go:57
	}
//line /usr/local/go/src/encoding/asn1/asn1.go:57
	// _ = "end of CoverTab[7488]"
//line /usr/local/go/src/encoding/asn1/asn1.go:57
	_go_fuzz_dep_.CoverTab[7489]++

//line /usr/local/go/src/encoding/asn1/asn1.go:62
	switch bytes[0] {
	case 0:
//line /usr/local/go/src/encoding/asn1/asn1.go:63
		_go_fuzz_dep_.CoverTab[7493]++
								ret = false
//line /usr/local/go/src/encoding/asn1/asn1.go:64
		// _ = "end of CoverTab[7493]"
	case 0xff:
//line /usr/local/go/src/encoding/asn1/asn1.go:65
		_go_fuzz_dep_.CoverTab[7494]++
								ret = true
//line /usr/local/go/src/encoding/asn1/asn1.go:66
		// _ = "end of CoverTab[7494]"
	default:
//line /usr/local/go/src/encoding/asn1/asn1.go:67
		_go_fuzz_dep_.CoverTab[7495]++
								err = SyntaxError{"invalid boolean"}
//line /usr/local/go/src/encoding/asn1/asn1.go:68
		// _ = "end of CoverTab[7495]"
	}
//line /usr/local/go/src/encoding/asn1/asn1.go:69
	// _ = "end of CoverTab[7489]"
//line /usr/local/go/src/encoding/asn1/asn1.go:69
	_go_fuzz_dep_.CoverTab[7490]++

							return
//line /usr/local/go/src/encoding/asn1/asn1.go:71
	// _ = "end of CoverTab[7490]"
}

//line /usr/local/go/src/encoding/asn1/asn1.go:76
// checkInteger returns nil if the given bytes are a valid DER-encoded
//line /usr/local/go/src/encoding/asn1/asn1.go:76
// INTEGER and an error otherwise.
//line /usr/local/go/src/encoding/asn1/asn1.go:78
func checkInteger(bytes []byte) error {
//line /usr/local/go/src/encoding/asn1/asn1.go:78
	_go_fuzz_dep_.CoverTab[7496]++
							if len(bytes) == 0 {
//line /usr/local/go/src/encoding/asn1/asn1.go:79
		_go_fuzz_dep_.CoverTab[7500]++
								return StructuralError{"empty integer"}
//line /usr/local/go/src/encoding/asn1/asn1.go:80
		// _ = "end of CoverTab[7500]"
	} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:81
		_go_fuzz_dep_.CoverTab[7501]++
//line /usr/local/go/src/encoding/asn1/asn1.go:81
		// _ = "end of CoverTab[7501]"
//line /usr/local/go/src/encoding/asn1/asn1.go:81
	}
//line /usr/local/go/src/encoding/asn1/asn1.go:81
	// _ = "end of CoverTab[7496]"
//line /usr/local/go/src/encoding/asn1/asn1.go:81
	_go_fuzz_dep_.CoverTab[7497]++
							if len(bytes) == 1 {
//line /usr/local/go/src/encoding/asn1/asn1.go:82
		_go_fuzz_dep_.CoverTab[7502]++
								return nil
//line /usr/local/go/src/encoding/asn1/asn1.go:83
		// _ = "end of CoverTab[7502]"
	} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:84
		_go_fuzz_dep_.CoverTab[7503]++
//line /usr/local/go/src/encoding/asn1/asn1.go:84
		// _ = "end of CoverTab[7503]"
//line /usr/local/go/src/encoding/asn1/asn1.go:84
	}
//line /usr/local/go/src/encoding/asn1/asn1.go:84
	// _ = "end of CoverTab[7497]"
//line /usr/local/go/src/encoding/asn1/asn1.go:84
	_go_fuzz_dep_.CoverTab[7498]++
							if (bytes[0] == 0 && func() bool {
//line /usr/local/go/src/encoding/asn1/asn1.go:85
		_go_fuzz_dep_.CoverTab[7504]++
//line /usr/local/go/src/encoding/asn1/asn1.go:85
		return bytes[1]&0x80 == 0
//line /usr/local/go/src/encoding/asn1/asn1.go:85
		// _ = "end of CoverTab[7504]"
//line /usr/local/go/src/encoding/asn1/asn1.go:85
	}()) || func() bool {
//line /usr/local/go/src/encoding/asn1/asn1.go:85
		_go_fuzz_dep_.CoverTab[7505]++
//line /usr/local/go/src/encoding/asn1/asn1.go:85
		return (bytes[0] == 0xff && func() bool {
//line /usr/local/go/src/encoding/asn1/asn1.go:85
			_go_fuzz_dep_.CoverTab[7506]++
//line /usr/local/go/src/encoding/asn1/asn1.go:85
			return bytes[1]&0x80 == 0x80
//line /usr/local/go/src/encoding/asn1/asn1.go:85
			// _ = "end of CoverTab[7506]"
//line /usr/local/go/src/encoding/asn1/asn1.go:85
		}())
//line /usr/local/go/src/encoding/asn1/asn1.go:85
		// _ = "end of CoverTab[7505]"
//line /usr/local/go/src/encoding/asn1/asn1.go:85
	}() {
//line /usr/local/go/src/encoding/asn1/asn1.go:85
		_go_fuzz_dep_.CoverTab[7507]++
								return StructuralError{"integer not minimally-encoded"}
//line /usr/local/go/src/encoding/asn1/asn1.go:86
		// _ = "end of CoverTab[7507]"
	} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:87
		_go_fuzz_dep_.CoverTab[7508]++
//line /usr/local/go/src/encoding/asn1/asn1.go:87
		// _ = "end of CoverTab[7508]"
//line /usr/local/go/src/encoding/asn1/asn1.go:87
	}
//line /usr/local/go/src/encoding/asn1/asn1.go:87
	// _ = "end of CoverTab[7498]"
//line /usr/local/go/src/encoding/asn1/asn1.go:87
	_go_fuzz_dep_.CoverTab[7499]++
							return nil
//line /usr/local/go/src/encoding/asn1/asn1.go:88
	// _ = "end of CoverTab[7499]"
}

// parseInt64 treats the given bytes as a big-endian, signed integer and
//line /usr/local/go/src/encoding/asn1/asn1.go:91
// returns the result.
//line /usr/local/go/src/encoding/asn1/asn1.go:93
func parseInt64(bytes []byte) (ret int64, err error) {
//line /usr/local/go/src/encoding/asn1/asn1.go:93
	_go_fuzz_dep_.CoverTab[7509]++
							err = checkInteger(bytes)
							if err != nil {
//line /usr/local/go/src/encoding/asn1/asn1.go:95
		_go_fuzz_dep_.CoverTab[7513]++
								return
//line /usr/local/go/src/encoding/asn1/asn1.go:96
		// _ = "end of CoverTab[7513]"
	} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:97
		_go_fuzz_dep_.CoverTab[7514]++
//line /usr/local/go/src/encoding/asn1/asn1.go:97
		// _ = "end of CoverTab[7514]"
//line /usr/local/go/src/encoding/asn1/asn1.go:97
	}
//line /usr/local/go/src/encoding/asn1/asn1.go:97
	// _ = "end of CoverTab[7509]"
//line /usr/local/go/src/encoding/asn1/asn1.go:97
	_go_fuzz_dep_.CoverTab[7510]++
							if len(bytes) > 8 {
//line /usr/local/go/src/encoding/asn1/asn1.go:98
		_go_fuzz_dep_.CoverTab[7515]++

								err = StructuralError{"integer too large"}
								return
//line /usr/local/go/src/encoding/asn1/asn1.go:101
		// _ = "end of CoverTab[7515]"
	} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:102
		_go_fuzz_dep_.CoverTab[7516]++
//line /usr/local/go/src/encoding/asn1/asn1.go:102
		// _ = "end of CoverTab[7516]"
//line /usr/local/go/src/encoding/asn1/asn1.go:102
	}
//line /usr/local/go/src/encoding/asn1/asn1.go:102
	// _ = "end of CoverTab[7510]"
//line /usr/local/go/src/encoding/asn1/asn1.go:102
	_go_fuzz_dep_.CoverTab[7511]++
							for bytesRead := 0; bytesRead < len(bytes); bytesRead++ {
//line /usr/local/go/src/encoding/asn1/asn1.go:103
		_go_fuzz_dep_.CoverTab[7517]++
								ret <<= 8
								ret |= int64(bytes[bytesRead])
//line /usr/local/go/src/encoding/asn1/asn1.go:105
		// _ = "end of CoverTab[7517]"
	}
//line /usr/local/go/src/encoding/asn1/asn1.go:106
	// _ = "end of CoverTab[7511]"
//line /usr/local/go/src/encoding/asn1/asn1.go:106
	_go_fuzz_dep_.CoverTab[7512]++

//line /usr/local/go/src/encoding/asn1/asn1.go:109
	ret <<= 64 - uint8(len(bytes))*8
							ret >>= 64 - uint8(len(bytes))*8
							return
//line /usr/local/go/src/encoding/asn1/asn1.go:111
	// _ = "end of CoverTab[7512]"
}

// parseInt32 treats the given bytes as a big-endian, signed integer and returns
//line /usr/local/go/src/encoding/asn1/asn1.go:114
// the result.
//line /usr/local/go/src/encoding/asn1/asn1.go:116
func parseInt32(bytes []byte) (int32, error) {
//line /usr/local/go/src/encoding/asn1/asn1.go:116
	_go_fuzz_dep_.CoverTab[7518]++
							if err := checkInteger(bytes); err != nil {
//line /usr/local/go/src/encoding/asn1/asn1.go:117
		_go_fuzz_dep_.CoverTab[7522]++
								return 0, err
//line /usr/local/go/src/encoding/asn1/asn1.go:118
		// _ = "end of CoverTab[7522]"
	} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:119
		_go_fuzz_dep_.CoverTab[7523]++
//line /usr/local/go/src/encoding/asn1/asn1.go:119
		// _ = "end of CoverTab[7523]"
//line /usr/local/go/src/encoding/asn1/asn1.go:119
	}
//line /usr/local/go/src/encoding/asn1/asn1.go:119
	// _ = "end of CoverTab[7518]"
//line /usr/local/go/src/encoding/asn1/asn1.go:119
	_go_fuzz_dep_.CoverTab[7519]++
							ret64, err := parseInt64(bytes)
							if err != nil {
//line /usr/local/go/src/encoding/asn1/asn1.go:121
		_go_fuzz_dep_.CoverTab[7524]++
								return 0, err
//line /usr/local/go/src/encoding/asn1/asn1.go:122
		// _ = "end of CoverTab[7524]"
	} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:123
		_go_fuzz_dep_.CoverTab[7525]++
//line /usr/local/go/src/encoding/asn1/asn1.go:123
		// _ = "end of CoverTab[7525]"
//line /usr/local/go/src/encoding/asn1/asn1.go:123
	}
//line /usr/local/go/src/encoding/asn1/asn1.go:123
	// _ = "end of CoverTab[7519]"
//line /usr/local/go/src/encoding/asn1/asn1.go:123
	_go_fuzz_dep_.CoverTab[7520]++
							if ret64 != int64(int32(ret64)) {
//line /usr/local/go/src/encoding/asn1/asn1.go:124
		_go_fuzz_dep_.CoverTab[7526]++
								return 0, StructuralError{"integer too large"}
//line /usr/local/go/src/encoding/asn1/asn1.go:125
		// _ = "end of CoverTab[7526]"
	} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:126
		_go_fuzz_dep_.CoverTab[7527]++
//line /usr/local/go/src/encoding/asn1/asn1.go:126
		// _ = "end of CoverTab[7527]"
//line /usr/local/go/src/encoding/asn1/asn1.go:126
	}
//line /usr/local/go/src/encoding/asn1/asn1.go:126
	// _ = "end of CoverTab[7520]"
//line /usr/local/go/src/encoding/asn1/asn1.go:126
	_go_fuzz_dep_.CoverTab[7521]++
							return int32(ret64), nil
//line /usr/local/go/src/encoding/asn1/asn1.go:127
	// _ = "end of CoverTab[7521]"
}

var bigOne = big.NewInt(1)

// parseBigInt treats the given bytes as a big-endian, signed integer and returns
//line /usr/local/go/src/encoding/asn1/asn1.go:132
// the result.
//line /usr/local/go/src/encoding/asn1/asn1.go:134
func parseBigInt(bytes []byte) (*big.Int, error) {
//line /usr/local/go/src/encoding/asn1/asn1.go:134
	_go_fuzz_dep_.CoverTab[7528]++
							if err := checkInteger(bytes); err != nil {
//line /usr/local/go/src/encoding/asn1/asn1.go:135
		_go_fuzz_dep_.CoverTab[7531]++
								return nil, err
//line /usr/local/go/src/encoding/asn1/asn1.go:136
		// _ = "end of CoverTab[7531]"
	} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:137
		_go_fuzz_dep_.CoverTab[7532]++
//line /usr/local/go/src/encoding/asn1/asn1.go:137
		// _ = "end of CoverTab[7532]"
//line /usr/local/go/src/encoding/asn1/asn1.go:137
	}
//line /usr/local/go/src/encoding/asn1/asn1.go:137
	// _ = "end of CoverTab[7528]"
//line /usr/local/go/src/encoding/asn1/asn1.go:137
	_go_fuzz_dep_.CoverTab[7529]++
							ret := new(big.Int)
							if len(bytes) > 0 && func() bool {
//line /usr/local/go/src/encoding/asn1/asn1.go:139
		_go_fuzz_dep_.CoverTab[7533]++
//line /usr/local/go/src/encoding/asn1/asn1.go:139
		return bytes[0]&0x80 == 0x80
//line /usr/local/go/src/encoding/asn1/asn1.go:139
		// _ = "end of CoverTab[7533]"
//line /usr/local/go/src/encoding/asn1/asn1.go:139
	}() {
//line /usr/local/go/src/encoding/asn1/asn1.go:139
		_go_fuzz_dep_.CoverTab[7534]++

								notBytes := make([]byte, len(bytes))
								for i := range notBytes {
//line /usr/local/go/src/encoding/asn1/asn1.go:142
			_go_fuzz_dep_.CoverTab[7536]++
									notBytes[i] = ^bytes[i]
//line /usr/local/go/src/encoding/asn1/asn1.go:143
			// _ = "end of CoverTab[7536]"
		}
//line /usr/local/go/src/encoding/asn1/asn1.go:144
		// _ = "end of CoverTab[7534]"
//line /usr/local/go/src/encoding/asn1/asn1.go:144
		_go_fuzz_dep_.CoverTab[7535]++
								ret.SetBytes(notBytes)
								ret.Add(ret, bigOne)
								ret.Neg(ret)
								return ret, nil
//line /usr/local/go/src/encoding/asn1/asn1.go:148
		// _ = "end of CoverTab[7535]"
	} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:149
		_go_fuzz_dep_.CoverTab[7537]++
//line /usr/local/go/src/encoding/asn1/asn1.go:149
		// _ = "end of CoverTab[7537]"
//line /usr/local/go/src/encoding/asn1/asn1.go:149
	}
//line /usr/local/go/src/encoding/asn1/asn1.go:149
	// _ = "end of CoverTab[7529]"
//line /usr/local/go/src/encoding/asn1/asn1.go:149
	_go_fuzz_dep_.CoverTab[7530]++
							ret.SetBytes(bytes)
							return ret, nil
//line /usr/local/go/src/encoding/asn1/asn1.go:151
	// _ = "end of CoverTab[7530]"
}

//line /usr/local/go/src/encoding/asn1/asn1.go:156
// BitString is the structure to use when you want an ASN.1 BIT STRING type. A
//line /usr/local/go/src/encoding/asn1/asn1.go:156
// bit string is padded up to the nearest byte in memory and the number of
//line /usr/local/go/src/encoding/asn1/asn1.go:156
// valid bits is recorded. Padding bits will be zero.
//line /usr/local/go/src/encoding/asn1/asn1.go:159
type BitString struct {
	Bytes		[]byte	// bits packed into bytes.
	BitLength	int	// length in bits.
}

// At returns the bit at the given index. If the index is out of range it
//line /usr/local/go/src/encoding/asn1/asn1.go:164
// returns 0.
//line /usr/local/go/src/encoding/asn1/asn1.go:166
func (b BitString) At(i int) int {
//line /usr/local/go/src/encoding/asn1/asn1.go:166
	_go_fuzz_dep_.CoverTab[7538]++
							if i < 0 || func() bool {
//line /usr/local/go/src/encoding/asn1/asn1.go:167
		_go_fuzz_dep_.CoverTab[7540]++
//line /usr/local/go/src/encoding/asn1/asn1.go:167
		return i >= b.BitLength
//line /usr/local/go/src/encoding/asn1/asn1.go:167
		// _ = "end of CoverTab[7540]"
//line /usr/local/go/src/encoding/asn1/asn1.go:167
	}() {
//line /usr/local/go/src/encoding/asn1/asn1.go:167
		_go_fuzz_dep_.CoverTab[7541]++
								return 0
//line /usr/local/go/src/encoding/asn1/asn1.go:168
		// _ = "end of CoverTab[7541]"
	} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:169
		_go_fuzz_dep_.CoverTab[7542]++
//line /usr/local/go/src/encoding/asn1/asn1.go:169
		// _ = "end of CoverTab[7542]"
//line /usr/local/go/src/encoding/asn1/asn1.go:169
	}
//line /usr/local/go/src/encoding/asn1/asn1.go:169
	// _ = "end of CoverTab[7538]"
//line /usr/local/go/src/encoding/asn1/asn1.go:169
	_go_fuzz_dep_.CoverTab[7539]++
							x := i / 8
							y := 7 - uint(i%8)
							return int(b.Bytes[x]>>y) & 1
//line /usr/local/go/src/encoding/asn1/asn1.go:172
	// _ = "end of CoverTab[7539]"
}

// RightAlign returns a slice where the padding bits are at the beginning. The
//line /usr/local/go/src/encoding/asn1/asn1.go:175
// slice may share memory with the BitString.
//line /usr/local/go/src/encoding/asn1/asn1.go:177
func (b BitString) RightAlign() []byte {
//line /usr/local/go/src/encoding/asn1/asn1.go:177
	_go_fuzz_dep_.CoverTab[7543]++
							shift := uint(8 - (b.BitLength % 8))
							if shift == 8 || func() bool {
//line /usr/local/go/src/encoding/asn1/asn1.go:179
		_go_fuzz_dep_.CoverTab[7546]++
//line /usr/local/go/src/encoding/asn1/asn1.go:179
		return len(b.Bytes) == 0
//line /usr/local/go/src/encoding/asn1/asn1.go:179
		// _ = "end of CoverTab[7546]"
//line /usr/local/go/src/encoding/asn1/asn1.go:179
	}() {
//line /usr/local/go/src/encoding/asn1/asn1.go:179
		_go_fuzz_dep_.CoverTab[7547]++
								return b.Bytes
//line /usr/local/go/src/encoding/asn1/asn1.go:180
		// _ = "end of CoverTab[7547]"
	} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:181
		_go_fuzz_dep_.CoverTab[7548]++
//line /usr/local/go/src/encoding/asn1/asn1.go:181
		// _ = "end of CoverTab[7548]"
//line /usr/local/go/src/encoding/asn1/asn1.go:181
	}
//line /usr/local/go/src/encoding/asn1/asn1.go:181
	// _ = "end of CoverTab[7543]"
//line /usr/local/go/src/encoding/asn1/asn1.go:181
	_go_fuzz_dep_.CoverTab[7544]++

							a := make([]byte, len(b.Bytes))
							a[0] = b.Bytes[0] >> shift
							for i := 1; i < len(b.Bytes); i++ {
//line /usr/local/go/src/encoding/asn1/asn1.go:185
		_go_fuzz_dep_.CoverTab[7549]++
								a[i] = b.Bytes[i-1] << (8 - shift)
								a[i] |= b.Bytes[i] >> shift
//line /usr/local/go/src/encoding/asn1/asn1.go:187
		// _ = "end of CoverTab[7549]"
	}
//line /usr/local/go/src/encoding/asn1/asn1.go:188
	// _ = "end of CoverTab[7544]"
//line /usr/local/go/src/encoding/asn1/asn1.go:188
	_go_fuzz_dep_.CoverTab[7545]++

							return a
//line /usr/local/go/src/encoding/asn1/asn1.go:190
	// _ = "end of CoverTab[7545]"
}

// parseBitString parses an ASN.1 bit string from the given byte slice and returns it.
func parseBitString(bytes []byte) (ret BitString, err error) {
//line /usr/local/go/src/encoding/asn1/asn1.go:194
	_go_fuzz_dep_.CoverTab[7550]++
							if len(bytes) == 0 {
//line /usr/local/go/src/encoding/asn1/asn1.go:195
		_go_fuzz_dep_.CoverTab[7553]++
								err = SyntaxError{"zero length BIT STRING"}
								return
//line /usr/local/go/src/encoding/asn1/asn1.go:197
		// _ = "end of CoverTab[7553]"
	} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:198
		_go_fuzz_dep_.CoverTab[7554]++
//line /usr/local/go/src/encoding/asn1/asn1.go:198
		// _ = "end of CoverTab[7554]"
//line /usr/local/go/src/encoding/asn1/asn1.go:198
	}
//line /usr/local/go/src/encoding/asn1/asn1.go:198
	// _ = "end of CoverTab[7550]"
//line /usr/local/go/src/encoding/asn1/asn1.go:198
	_go_fuzz_dep_.CoverTab[7551]++
							paddingBits := int(bytes[0])
							if paddingBits > 7 || func() bool {
//line /usr/local/go/src/encoding/asn1/asn1.go:200
		_go_fuzz_dep_.CoverTab[7555]++
//line /usr/local/go/src/encoding/asn1/asn1.go:200
		return len(bytes) == 1 && func() bool {
									_go_fuzz_dep_.CoverTab[7556]++
//line /usr/local/go/src/encoding/asn1/asn1.go:201
			return paddingBits > 0
//line /usr/local/go/src/encoding/asn1/asn1.go:201
			// _ = "end of CoverTab[7556]"
//line /usr/local/go/src/encoding/asn1/asn1.go:201
		}()
//line /usr/local/go/src/encoding/asn1/asn1.go:201
		// _ = "end of CoverTab[7555]"
//line /usr/local/go/src/encoding/asn1/asn1.go:201
	}() || func() bool {
//line /usr/local/go/src/encoding/asn1/asn1.go:201
		_go_fuzz_dep_.CoverTab[7557]++
//line /usr/local/go/src/encoding/asn1/asn1.go:201
		return bytes[len(bytes)-1]&((1<<bytes[0])-1) != 0
								// _ = "end of CoverTab[7557]"
//line /usr/local/go/src/encoding/asn1/asn1.go:202
	}() {
//line /usr/local/go/src/encoding/asn1/asn1.go:202
		_go_fuzz_dep_.CoverTab[7558]++
								err = SyntaxError{"invalid padding bits in BIT STRING"}
								return
//line /usr/local/go/src/encoding/asn1/asn1.go:204
		// _ = "end of CoverTab[7558]"
	} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:205
		_go_fuzz_dep_.CoverTab[7559]++
//line /usr/local/go/src/encoding/asn1/asn1.go:205
		// _ = "end of CoverTab[7559]"
//line /usr/local/go/src/encoding/asn1/asn1.go:205
	}
//line /usr/local/go/src/encoding/asn1/asn1.go:205
	// _ = "end of CoverTab[7551]"
//line /usr/local/go/src/encoding/asn1/asn1.go:205
	_go_fuzz_dep_.CoverTab[7552]++
							ret.BitLength = (len(bytes)-1)*8 - paddingBits
							ret.Bytes = bytes[1:]
							return
//line /usr/local/go/src/encoding/asn1/asn1.go:208
	// _ = "end of CoverTab[7552]"
}

//line /usr/local/go/src/encoding/asn1/asn1.go:213
// NullRawValue is a RawValue with its Tag set to the ASN.1 NULL type tag (5).
var NullRawValue = RawValue{Tag: TagNull}

// NullBytes contains bytes representing the DER-encoded ASN.1 NULL type.
var NullBytes = []byte{TagNull, 0}

//line /usr/local/go/src/encoding/asn1/asn1.go:221
// An ObjectIdentifier represents an ASN.1 OBJECT IDENTIFIER.
type ObjectIdentifier []int

// Equal reports whether oi and other represent the same identifier.
func (oi ObjectIdentifier) Equal(other ObjectIdentifier) bool {
//line /usr/local/go/src/encoding/asn1/asn1.go:225
	_go_fuzz_dep_.CoverTab[7560]++
							if len(oi) != len(other) {
//line /usr/local/go/src/encoding/asn1/asn1.go:226
		_go_fuzz_dep_.CoverTab[7563]++
								return false
//line /usr/local/go/src/encoding/asn1/asn1.go:227
		// _ = "end of CoverTab[7563]"
	} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:228
		_go_fuzz_dep_.CoverTab[7564]++
//line /usr/local/go/src/encoding/asn1/asn1.go:228
		// _ = "end of CoverTab[7564]"
//line /usr/local/go/src/encoding/asn1/asn1.go:228
	}
//line /usr/local/go/src/encoding/asn1/asn1.go:228
	// _ = "end of CoverTab[7560]"
//line /usr/local/go/src/encoding/asn1/asn1.go:228
	_go_fuzz_dep_.CoverTab[7561]++
							for i := 0; i < len(oi); i++ {
//line /usr/local/go/src/encoding/asn1/asn1.go:229
		_go_fuzz_dep_.CoverTab[7565]++
								if oi[i] != other[i] {
//line /usr/local/go/src/encoding/asn1/asn1.go:230
			_go_fuzz_dep_.CoverTab[7566]++
									return false
//line /usr/local/go/src/encoding/asn1/asn1.go:231
			// _ = "end of CoverTab[7566]"
		} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:232
			_go_fuzz_dep_.CoverTab[7567]++
//line /usr/local/go/src/encoding/asn1/asn1.go:232
			// _ = "end of CoverTab[7567]"
//line /usr/local/go/src/encoding/asn1/asn1.go:232
		}
//line /usr/local/go/src/encoding/asn1/asn1.go:232
		// _ = "end of CoverTab[7565]"
	}
//line /usr/local/go/src/encoding/asn1/asn1.go:233
	// _ = "end of CoverTab[7561]"
//line /usr/local/go/src/encoding/asn1/asn1.go:233
	_go_fuzz_dep_.CoverTab[7562]++

							return true
//line /usr/local/go/src/encoding/asn1/asn1.go:235
	// _ = "end of CoverTab[7562]"
}

func (oi ObjectIdentifier) String() string {
//line /usr/local/go/src/encoding/asn1/asn1.go:238
	_go_fuzz_dep_.CoverTab[7568]++
							var s string

							for i, v := range oi {
//line /usr/local/go/src/encoding/asn1/asn1.go:241
		_go_fuzz_dep_.CoverTab[7570]++
								if i > 0 {
//line /usr/local/go/src/encoding/asn1/asn1.go:242
			_go_fuzz_dep_.CoverTab[7572]++
									s += "."
//line /usr/local/go/src/encoding/asn1/asn1.go:243
			// _ = "end of CoverTab[7572]"
		} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:244
			_go_fuzz_dep_.CoverTab[7573]++
//line /usr/local/go/src/encoding/asn1/asn1.go:244
			// _ = "end of CoverTab[7573]"
//line /usr/local/go/src/encoding/asn1/asn1.go:244
		}
//line /usr/local/go/src/encoding/asn1/asn1.go:244
		// _ = "end of CoverTab[7570]"
//line /usr/local/go/src/encoding/asn1/asn1.go:244
		_go_fuzz_dep_.CoverTab[7571]++
								s += strconv.Itoa(v)
//line /usr/local/go/src/encoding/asn1/asn1.go:245
		// _ = "end of CoverTab[7571]"
	}
//line /usr/local/go/src/encoding/asn1/asn1.go:246
	// _ = "end of CoverTab[7568]"
//line /usr/local/go/src/encoding/asn1/asn1.go:246
	_go_fuzz_dep_.CoverTab[7569]++

							return s
//line /usr/local/go/src/encoding/asn1/asn1.go:248
	// _ = "end of CoverTab[7569]"
}

// parseObjectIdentifier parses an OBJECT IDENTIFIER from the given bytes and
//line /usr/local/go/src/encoding/asn1/asn1.go:251
// returns it. An object identifier is a sequence of variable length integers
//line /usr/local/go/src/encoding/asn1/asn1.go:251
// that are assigned in a hierarchy.
//line /usr/local/go/src/encoding/asn1/asn1.go:254
func parseObjectIdentifier(bytes []byte) (s ObjectIdentifier, err error) {
//line /usr/local/go/src/encoding/asn1/asn1.go:254
	_go_fuzz_dep_.CoverTab[7574]++
							if len(bytes) == 0 {
//line /usr/local/go/src/encoding/asn1/asn1.go:255
		_go_fuzz_dep_.CoverTab[7579]++
								err = SyntaxError{"zero length OBJECT IDENTIFIER"}
								return
//line /usr/local/go/src/encoding/asn1/asn1.go:257
		// _ = "end of CoverTab[7579]"
	} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:258
		_go_fuzz_dep_.CoverTab[7580]++
//line /usr/local/go/src/encoding/asn1/asn1.go:258
		// _ = "end of CoverTab[7580]"
//line /usr/local/go/src/encoding/asn1/asn1.go:258
	}
//line /usr/local/go/src/encoding/asn1/asn1.go:258
	// _ = "end of CoverTab[7574]"
//line /usr/local/go/src/encoding/asn1/asn1.go:258
	_go_fuzz_dep_.CoverTab[7575]++

//line /usr/local/go/src/encoding/asn1/asn1.go:262
	s = make([]int, len(bytes)+1)

//line /usr/local/go/src/encoding/asn1/asn1.go:268
	v, offset, err := parseBase128Int(bytes, 0)
	if err != nil {
//line /usr/local/go/src/encoding/asn1/asn1.go:269
		_go_fuzz_dep_.CoverTab[7581]++
								return
//line /usr/local/go/src/encoding/asn1/asn1.go:270
		// _ = "end of CoverTab[7581]"
	} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:271
		_go_fuzz_dep_.CoverTab[7582]++
//line /usr/local/go/src/encoding/asn1/asn1.go:271
		// _ = "end of CoverTab[7582]"
//line /usr/local/go/src/encoding/asn1/asn1.go:271
	}
//line /usr/local/go/src/encoding/asn1/asn1.go:271
	// _ = "end of CoverTab[7575]"
//line /usr/local/go/src/encoding/asn1/asn1.go:271
	_go_fuzz_dep_.CoverTab[7576]++
							if v < 80 {
//line /usr/local/go/src/encoding/asn1/asn1.go:272
		_go_fuzz_dep_.CoverTab[7583]++
								s[0] = v / 40
								s[1] = v % 40
//line /usr/local/go/src/encoding/asn1/asn1.go:274
		// _ = "end of CoverTab[7583]"
	} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:275
		_go_fuzz_dep_.CoverTab[7584]++
								s[0] = 2
								s[1] = v - 80
//line /usr/local/go/src/encoding/asn1/asn1.go:277
		// _ = "end of CoverTab[7584]"
	}
//line /usr/local/go/src/encoding/asn1/asn1.go:278
	// _ = "end of CoverTab[7576]"
//line /usr/local/go/src/encoding/asn1/asn1.go:278
	_go_fuzz_dep_.CoverTab[7577]++

							i := 2
							for ; offset < len(bytes); i++ {
//line /usr/local/go/src/encoding/asn1/asn1.go:281
		_go_fuzz_dep_.CoverTab[7585]++
								v, offset, err = parseBase128Int(bytes, offset)
								if err != nil {
//line /usr/local/go/src/encoding/asn1/asn1.go:283
			_go_fuzz_dep_.CoverTab[7587]++
									return
//line /usr/local/go/src/encoding/asn1/asn1.go:284
			// _ = "end of CoverTab[7587]"
		} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:285
			_go_fuzz_dep_.CoverTab[7588]++
//line /usr/local/go/src/encoding/asn1/asn1.go:285
			// _ = "end of CoverTab[7588]"
//line /usr/local/go/src/encoding/asn1/asn1.go:285
		}
//line /usr/local/go/src/encoding/asn1/asn1.go:285
		// _ = "end of CoverTab[7585]"
//line /usr/local/go/src/encoding/asn1/asn1.go:285
		_go_fuzz_dep_.CoverTab[7586]++
								s[i] = v
//line /usr/local/go/src/encoding/asn1/asn1.go:286
		// _ = "end of CoverTab[7586]"
	}
//line /usr/local/go/src/encoding/asn1/asn1.go:287
	// _ = "end of CoverTab[7577]"
//line /usr/local/go/src/encoding/asn1/asn1.go:287
	_go_fuzz_dep_.CoverTab[7578]++
							s = s[0:i]
							return
//line /usr/local/go/src/encoding/asn1/asn1.go:289
	// _ = "end of CoverTab[7578]"
}

//line /usr/local/go/src/encoding/asn1/asn1.go:294
// An Enumerated is represented as a plain int.
type Enumerated int

//line /usr/local/go/src/encoding/asn1/asn1.go:299
// A Flag accepts any data and is set to true if present.
type Flag bool

// parseBase128Int parses a base-128 encoded int from the given offset in the
//line /usr/local/go/src/encoding/asn1/asn1.go:302
// given byte slice. It returns the value and the new offset.
//line /usr/local/go/src/encoding/asn1/asn1.go:304
func parseBase128Int(bytes []byte, initOffset int) (ret, offset int, err error) {
//line /usr/local/go/src/encoding/asn1/asn1.go:304
	_go_fuzz_dep_.CoverTab[7589]++
							offset = initOffset
							var ret64 int64
							for shifted := 0; offset < len(bytes); shifted++ {
//line /usr/local/go/src/encoding/asn1/asn1.go:307
		_go_fuzz_dep_.CoverTab[7591]++

//line /usr/local/go/src/encoding/asn1/asn1.go:310
		if shifted == 5 {
//line /usr/local/go/src/encoding/asn1/asn1.go:310
			_go_fuzz_dep_.CoverTab[7594]++
									err = StructuralError{"base 128 integer too large"}
									return
//line /usr/local/go/src/encoding/asn1/asn1.go:312
			// _ = "end of CoverTab[7594]"
		} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:313
			_go_fuzz_dep_.CoverTab[7595]++
//line /usr/local/go/src/encoding/asn1/asn1.go:313
			// _ = "end of CoverTab[7595]"
//line /usr/local/go/src/encoding/asn1/asn1.go:313
		}
//line /usr/local/go/src/encoding/asn1/asn1.go:313
		// _ = "end of CoverTab[7591]"
//line /usr/local/go/src/encoding/asn1/asn1.go:313
		_go_fuzz_dep_.CoverTab[7592]++
								ret64 <<= 7
								b := bytes[offset]

//line /usr/local/go/src/encoding/asn1/asn1.go:318
		if shifted == 0 && func() bool {
//line /usr/local/go/src/encoding/asn1/asn1.go:318
			_go_fuzz_dep_.CoverTab[7596]++
//line /usr/local/go/src/encoding/asn1/asn1.go:318
			return b == 0x80
//line /usr/local/go/src/encoding/asn1/asn1.go:318
			// _ = "end of CoverTab[7596]"
//line /usr/local/go/src/encoding/asn1/asn1.go:318
		}() {
//line /usr/local/go/src/encoding/asn1/asn1.go:318
			_go_fuzz_dep_.CoverTab[7597]++
									err = SyntaxError{"integer is not minimally encoded"}
									return
//line /usr/local/go/src/encoding/asn1/asn1.go:320
			// _ = "end of CoverTab[7597]"
		} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:321
			_go_fuzz_dep_.CoverTab[7598]++
//line /usr/local/go/src/encoding/asn1/asn1.go:321
			// _ = "end of CoverTab[7598]"
//line /usr/local/go/src/encoding/asn1/asn1.go:321
		}
//line /usr/local/go/src/encoding/asn1/asn1.go:321
		// _ = "end of CoverTab[7592]"
//line /usr/local/go/src/encoding/asn1/asn1.go:321
		_go_fuzz_dep_.CoverTab[7593]++
								ret64 |= int64(b & 0x7f)
								offset++
								if b&0x80 == 0 {
//line /usr/local/go/src/encoding/asn1/asn1.go:324
			_go_fuzz_dep_.CoverTab[7599]++
									ret = int(ret64)

									if ret64 > math.MaxInt32 {
//line /usr/local/go/src/encoding/asn1/asn1.go:327
				_go_fuzz_dep_.CoverTab[7601]++
										err = StructuralError{"base 128 integer too large"}
//line /usr/local/go/src/encoding/asn1/asn1.go:328
				// _ = "end of CoverTab[7601]"
			} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:329
				_go_fuzz_dep_.CoverTab[7602]++
//line /usr/local/go/src/encoding/asn1/asn1.go:329
				// _ = "end of CoverTab[7602]"
//line /usr/local/go/src/encoding/asn1/asn1.go:329
			}
//line /usr/local/go/src/encoding/asn1/asn1.go:329
			// _ = "end of CoverTab[7599]"
//line /usr/local/go/src/encoding/asn1/asn1.go:329
			_go_fuzz_dep_.CoverTab[7600]++
									return
//line /usr/local/go/src/encoding/asn1/asn1.go:330
			// _ = "end of CoverTab[7600]"
		} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:331
			_go_fuzz_dep_.CoverTab[7603]++
//line /usr/local/go/src/encoding/asn1/asn1.go:331
			// _ = "end of CoverTab[7603]"
//line /usr/local/go/src/encoding/asn1/asn1.go:331
		}
//line /usr/local/go/src/encoding/asn1/asn1.go:331
		// _ = "end of CoverTab[7593]"
	}
//line /usr/local/go/src/encoding/asn1/asn1.go:332
	// _ = "end of CoverTab[7589]"
//line /usr/local/go/src/encoding/asn1/asn1.go:332
	_go_fuzz_dep_.CoverTab[7590]++
							err = SyntaxError{"truncated base 128 integer"}
							return
//line /usr/local/go/src/encoding/asn1/asn1.go:334
	// _ = "end of CoverTab[7590]"
}

//line /usr/local/go/src/encoding/asn1/asn1.go:339
func parseUTCTime(bytes []byte) (ret time.Time, err error) {
//line /usr/local/go/src/encoding/asn1/asn1.go:339
	_go_fuzz_dep_.CoverTab[7604]++
							s := string(bytes)

							formatStr := "0601021504Z0700"
							ret, err = time.Parse(formatStr, s)
							if err != nil {
//line /usr/local/go/src/encoding/asn1/asn1.go:344
		_go_fuzz_dep_.CoverTab[7609]++
								formatStr = "060102150405Z0700"
								ret, err = time.Parse(formatStr, s)
//line /usr/local/go/src/encoding/asn1/asn1.go:346
		// _ = "end of CoverTab[7609]"
	} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:347
		_go_fuzz_dep_.CoverTab[7610]++
//line /usr/local/go/src/encoding/asn1/asn1.go:347
		// _ = "end of CoverTab[7610]"
//line /usr/local/go/src/encoding/asn1/asn1.go:347
	}
//line /usr/local/go/src/encoding/asn1/asn1.go:347
	// _ = "end of CoverTab[7604]"
//line /usr/local/go/src/encoding/asn1/asn1.go:347
	_go_fuzz_dep_.CoverTab[7605]++
							if err != nil {
//line /usr/local/go/src/encoding/asn1/asn1.go:348
		_go_fuzz_dep_.CoverTab[7611]++
								return
//line /usr/local/go/src/encoding/asn1/asn1.go:349
		// _ = "end of CoverTab[7611]"
	} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:350
		_go_fuzz_dep_.CoverTab[7612]++
//line /usr/local/go/src/encoding/asn1/asn1.go:350
		// _ = "end of CoverTab[7612]"
//line /usr/local/go/src/encoding/asn1/asn1.go:350
	}
//line /usr/local/go/src/encoding/asn1/asn1.go:350
	// _ = "end of CoverTab[7605]"
//line /usr/local/go/src/encoding/asn1/asn1.go:350
	_go_fuzz_dep_.CoverTab[7606]++

							if serialized := ret.Format(formatStr); serialized != s {
//line /usr/local/go/src/encoding/asn1/asn1.go:352
		_go_fuzz_dep_.CoverTab[7613]++
								err = fmt.Errorf("asn1: time did not serialize back to the original value and may be invalid: given %q, but serialized as %q", s, serialized)
								return
//line /usr/local/go/src/encoding/asn1/asn1.go:354
		// _ = "end of CoverTab[7613]"
	} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:355
		_go_fuzz_dep_.CoverTab[7614]++
//line /usr/local/go/src/encoding/asn1/asn1.go:355
		// _ = "end of CoverTab[7614]"
//line /usr/local/go/src/encoding/asn1/asn1.go:355
	}
//line /usr/local/go/src/encoding/asn1/asn1.go:355
	// _ = "end of CoverTab[7606]"
//line /usr/local/go/src/encoding/asn1/asn1.go:355
	_go_fuzz_dep_.CoverTab[7607]++

							if ret.Year() >= 2050 {
//line /usr/local/go/src/encoding/asn1/asn1.go:357
		_go_fuzz_dep_.CoverTab[7615]++

								ret = ret.AddDate(-100, 0, 0)
//line /usr/local/go/src/encoding/asn1/asn1.go:359
		// _ = "end of CoverTab[7615]"
	} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:360
		_go_fuzz_dep_.CoverTab[7616]++
//line /usr/local/go/src/encoding/asn1/asn1.go:360
		// _ = "end of CoverTab[7616]"
//line /usr/local/go/src/encoding/asn1/asn1.go:360
	}
//line /usr/local/go/src/encoding/asn1/asn1.go:360
	// _ = "end of CoverTab[7607]"
//line /usr/local/go/src/encoding/asn1/asn1.go:360
	_go_fuzz_dep_.CoverTab[7608]++

							return
//line /usr/local/go/src/encoding/asn1/asn1.go:362
	// _ = "end of CoverTab[7608]"
}

// parseGeneralizedTime parses the GeneralizedTime from the given byte slice
//line /usr/local/go/src/encoding/asn1/asn1.go:365
// and returns the resulting time.
//line /usr/local/go/src/encoding/asn1/asn1.go:367
func parseGeneralizedTime(bytes []byte) (ret time.Time, err error) {
//line /usr/local/go/src/encoding/asn1/asn1.go:367
	_go_fuzz_dep_.CoverTab[7617]++
							const formatStr = "20060102150405Z0700"
							s := string(bytes)

							if ret, err = time.Parse(formatStr, s); err != nil {
//line /usr/local/go/src/encoding/asn1/asn1.go:371
		_go_fuzz_dep_.CoverTab[7620]++
								return
//line /usr/local/go/src/encoding/asn1/asn1.go:372
		// _ = "end of CoverTab[7620]"
	} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:373
		_go_fuzz_dep_.CoverTab[7621]++
//line /usr/local/go/src/encoding/asn1/asn1.go:373
		// _ = "end of CoverTab[7621]"
//line /usr/local/go/src/encoding/asn1/asn1.go:373
	}
//line /usr/local/go/src/encoding/asn1/asn1.go:373
	// _ = "end of CoverTab[7617]"
//line /usr/local/go/src/encoding/asn1/asn1.go:373
	_go_fuzz_dep_.CoverTab[7618]++

							if serialized := ret.Format(formatStr); serialized != s {
//line /usr/local/go/src/encoding/asn1/asn1.go:375
		_go_fuzz_dep_.CoverTab[7622]++
								err = fmt.Errorf("asn1: time did not serialize back to the original value and may be invalid: given %q, but serialized as %q", s, serialized)
//line /usr/local/go/src/encoding/asn1/asn1.go:376
		// _ = "end of CoverTab[7622]"
	} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:377
		_go_fuzz_dep_.CoverTab[7623]++
//line /usr/local/go/src/encoding/asn1/asn1.go:377
		// _ = "end of CoverTab[7623]"
//line /usr/local/go/src/encoding/asn1/asn1.go:377
	}
//line /usr/local/go/src/encoding/asn1/asn1.go:377
	// _ = "end of CoverTab[7618]"
//line /usr/local/go/src/encoding/asn1/asn1.go:377
	_go_fuzz_dep_.CoverTab[7619]++

							return
//line /usr/local/go/src/encoding/asn1/asn1.go:379
	// _ = "end of CoverTab[7619]"
}

//line /usr/local/go/src/encoding/asn1/asn1.go:384
// parseNumericString parses an ASN.1 NumericString from the given byte array
//line /usr/local/go/src/encoding/asn1/asn1.go:384
// and returns it.
//line /usr/local/go/src/encoding/asn1/asn1.go:386
func parseNumericString(bytes []byte) (ret string, err error) {
//line /usr/local/go/src/encoding/asn1/asn1.go:386
	_go_fuzz_dep_.CoverTab[7624]++
							for _, b := range bytes {
//line /usr/local/go/src/encoding/asn1/asn1.go:387
		_go_fuzz_dep_.CoverTab[7626]++
								if !isNumeric(b) {
//line /usr/local/go/src/encoding/asn1/asn1.go:388
			_go_fuzz_dep_.CoverTab[7627]++
									return "", SyntaxError{"NumericString contains invalid character"}
//line /usr/local/go/src/encoding/asn1/asn1.go:389
			// _ = "end of CoverTab[7627]"
		} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:390
			_go_fuzz_dep_.CoverTab[7628]++
//line /usr/local/go/src/encoding/asn1/asn1.go:390
			// _ = "end of CoverTab[7628]"
//line /usr/local/go/src/encoding/asn1/asn1.go:390
		}
//line /usr/local/go/src/encoding/asn1/asn1.go:390
		// _ = "end of CoverTab[7626]"
	}
//line /usr/local/go/src/encoding/asn1/asn1.go:391
	// _ = "end of CoverTab[7624]"
//line /usr/local/go/src/encoding/asn1/asn1.go:391
	_go_fuzz_dep_.CoverTab[7625]++
							return string(bytes), nil
//line /usr/local/go/src/encoding/asn1/asn1.go:392
	// _ = "end of CoverTab[7625]"
}

// isNumeric reports whether the given b is in the ASN.1 NumericString set.
func isNumeric(b byte) bool {
//line /usr/local/go/src/encoding/asn1/asn1.go:396
	_go_fuzz_dep_.CoverTab[7629]++
							return '0' <= b && func() bool {
//line /usr/local/go/src/encoding/asn1/asn1.go:397
		_go_fuzz_dep_.CoverTab[7630]++
//line /usr/local/go/src/encoding/asn1/asn1.go:397
		return b <= '9'
//line /usr/local/go/src/encoding/asn1/asn1.go:397
		// _ = "end of CoverTab[7630]"
//line /usr/local/go/src/encoding/asn1/asn1.go:397
	}() || func() bool {
//line /usr/local/go/src/encoding/asn1/asn1.go:397
		_go_fuzz_dep_.CoverTab[7631]++
//line /usr/local/go/src/encoding/asn1/asn1.go:397
		return b == ' '
								// _ = "end of CoverTab[7631]"
//line /usr/local/go/src/encoding/asn1/asn1.go:398
	}()
//line /usr/local/go/src/encoding/asn1/asn1.go:398
	// _ = "end of CoverTab[7629]"
}

//line /usr/local/go/src/encoding/asn1/asn1.go:403
// parsePrintableString parses an ASN.1 PrintableString from the given byte
//line /usr/local/go/src/encoding/asn1/asn1.go:403
// array and returns it.
//line /usr/local/go/src/encoding/asn1/asn1.go:405
func parsePrintableString(bytes []byte) (ret string, err error) {
//line /usr/local/go/src/encoding/asn1/asn1.go:405
	_go_fuzz_dep_.CoverTab[7632]++
							for _, b := range bytes {
//line /usr/local/go/src/encoding/asn1/asn1.go:406
		_go_fuzz_dep_.CoverTab[7634]++
								if !isPrintable(b, allowAsterisk, allowAmpersand) {
//line /usr/local/go/src/encoding/asn1/asn1.go:407
			_go_fuzz_dep_.CoverTab[7635]++
									err = SyntaxError{"PrintableString contains invalid character"}
									return
//line /usr/local/go/src/encoding/asn1/asn1.go:409
			// _ = "end of CoverTab[7635]"
		} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:410
			_go_fuzz_dep_.CoverTab[7636]++
//line /usr/local/go/src/encoding/asn1/asn1.go:410
			// _ = "end of CoverTab[7636]"
//line /usr/local/go/src/encoding/asn1/asn1.go:410
		}
//line /usr/local/go/src/encoding/asn1/asn1.go:410
		// _ = "end of CoverTab[7634]"
	}
//line /usr/local/go/src/encoding/asn1/asn1.go:411
	// _ = "end of CoverTab[7632]"
//line /usr/local/go/src/encoding/asn1/asn1.go:411
	_go_fuzz_dep_.CoverTab[7633]++
							ret = string(bytes)
							return
//line /usr/local/go/src/encoding/asn1/asn1.go:413
	// _ = "end of CoverTab[7633]"
}

type asteriskFlag bool
type ampersandFlag bool

const (
	allowAsterisk	asteriskFlag	= true
	rejectAsterisk	asteriskFlag	= false

	allowAmpersand	ampersandFlag	= true
	rejectAmpersand	ampersandFlag	= false
)

// isPrintable reports whether the given b is in the ASN.1 PrintableString set.
//line /usr/local/go/src/encoding/asn1/asn1.go:427
// If asterisk is allowAsterisk then '*' is also allowed, reflecting existing
//line /usr/local/go/src/encoding/asn1/asn1.go:427
// practice. If ampersand is allowAmpersand then '&' is allowed as well.
//line /usr/local/go/src/encoding/asn1/asn1.go:430
func isPrintable(b byte, asterisk asteriskFlag, ampersand ampersandFlag) bool {
//line /usr/local/go/src/encoding/asn1/asn1.go:430
	_go_fuzz_dep_.CoverTab[7637]++
							return 'a' <= b && func() bool {
//line /usr/local/go/src/encoding/asn1/asn1.go:431
		_go_fuzz_dep_.CoverTab[7638]++
//line /usr/local/go/src/encoding/asn1/asn1.go:431
		return b <= 'z'
//line /usr/local/go/src/encoding/asn1/asn1.go:431
		// _ = "end of CoverTab[7638]"
//line /usr/local/go/src/encoding/asn1/asn1.go:431
	}() || func() bool {
//line /usr/local/go/src/encoding/asn1/asn1.go:431
		_go_fuzz_dep_.CoverTab[7639]++
//line /usr/local/go/src/encoding/asn1/asn1.go:431
		return 'A' <= b && func() bool {
									_go_fuzz_dep_.CoverTab[7640]++
//line /usr/local/go/src/encoding/asn1/asn1.go:432
			return b <= 'Z'
//line /usr/local/go/src/encoding/asn1/asn1.go:432
			// _ = "end of CoverTab[7640]"
//line /usr/local/go/src/encoding/asn1/asn1.go:432
		}()
//line /usr/local/go/src/encoding/asn1/asn1.go:432
		// _ = "end of CoverTab[7639]"
//line /usr/local/go/src/encoding/asn1/asn1.go:432
	}() || func() bool {
//line /usr/local/go/src/encoding/asn1/asn1.go:432
		_go_fuzz_dep_.CoverTab[7641]++
//line /usr/local/go/src/encoding/asn1/asn1.go:432
		return '0' <= b && func() bool {
									_go_fuzz_dep_.CoverTab[7642]++
//line /usr/local/go/src/encoding/asn1/asn1.go:433
			return b <= '9'
//line /usr/local/go/src/encoding/asn1/asn1.go:433
			// _ = "end of CoverTab[7642]"
//line /usr/local/go/src/encoding/asn1/asn1.go:433
		}()
//line /usr/local/go/src/encoding/asn1/asn1.go:433
		// _ = "end of CoverTab[7641]"
//line /usr/local/go/src/encoding/asn1/asn1.go:433
	}() || func() bool {
//line /usr/local/go/src/encoding/asn1/asn1.go:433
		_go_fuzz_dep_.CoverTab[7643]++
//line /usr/local/go/src/encoding/asn1/asn1.go:433
		return '\'' <= b && func() bool {
									_go_fuzz_dep_.CoverTab[7644]++
//line /usr/local/go/src/encoding/asn1/asn1.go:434
			return b <= ')'
//line /usr/local/go/src/encoding/asn1/asn1.go:434
			// _ = "end of CoverTab[7644]"
//line /usr/local/go/src/encoding/asn1/asn1.go:434
		}()
//line /usr/local/go/src/encoding/asn1/asn1.go:434
		// _ = "end of CoverTab[7643]"
//line /usr/local/go/src/encoding/asn1/asn1.go:434
	}() || func() bool {
//line /usr/local/go/src/encoding/asn1/asn1.go:434
		_go_fuzz_dep_.CoverTab[7645]++
//line /usr/local/go/src/encoding/asn1/asn1.go:434
		return '+' <= b && func() bool {
									_go_fuzz_dep_.CoverTab[7646]++
//line /usr/local/go/src/encoding/asn1/asn1.go:435
			return b <= '/'
//line /usr/local/go/src/encoding/asn1/asn1.go:435
			// _ = "end of CoverTab[7646]"
//line /usr/local/go/src/encoding/asn1/asn1.go:435
		}()
//line /usr/local/go/src/encoding/asn1/asn1.go:435
		// _ = "end of CoverTab[7645]"
//line /usr/local/go/src/encoding/asn1/asn1.go:435
	}() || func() bool {
//line /usr/local/go/src/encoding/asn1/asn1.go:435
		_go_fuzz_dep_.CoverTab[7647]++
//line /usr/local/go/src/encoding/asn1/asn1.go:435
		return b == ' '
								// _ = "end of CoverTab[7647]"
//line /usr/local/go/src/encoding/asn1/asn1.go:436
	}() || func() bool {
//line /usr/local/go/src/encoding/asn1/asn1.go:436
		_go_fuzz_dep_.CoverTab[7648]++
//line /usr/local/go/src/encoding/asn1/asn1.go:436
		return b == ':'
								// _ = "end of CoverTab[7648]"
//line /usr/local/go/src/encoding/asn1/asn1.go:437
	}() || func() bool {
//line /usr/local/go/src/encoding/asn1/asn1.go:437
		_go_fuzz_dep_.CoverTab[7649]++
//line /usr/local/go/src/encoding/asn1/asn1.go:437
		return b == '='
								// _ = "end of CoverTab[7649]"
//line /usr/local/go/src/encoding/asn1/asn1.go:438
	}() || func() bool {
//line /usr/local/go/src/encoding/asn1/asn1.go:438
		_go_fuzz_dep_.CoverTab[7650]++
//line /usr/local/go/src/encoding/asn1/asn1.go:438
		return b == '?'
								// _ = "end of CoverTab[7650]"
//line /usr/local/go/src/encoding/asn1/asn1.go:439
	}() || func() bool {
//line /usr/local/go/src/encoding/asn1/asn1.go:439
		_go_fuzz_dep_.CoverTab[7651]++
//line /usr/local/go/src/encoding/asn1/asn1.go:439
		return (bool(asterisk) && func() bool {
//line /usr/local/go/src/encoding/asn1/asn1.go:443
			_go_fuzz_dep_.CoverTab[7652]++
//line /usr/local/go/src/encoding/asn1/asn1.go:443
			return b == '*'
//line /usr/local/go/src/encoding/asn1/asn1.go:443
			// _ = "end of CoverTab[7652]"
//line /usr/local/go/src/encoding/asn1/asn1.go:443
		}())
//line /usr/local/go/src/encoding/asn1/asn1.go:443
		// _ = "end of CoverTab[7651]"
//line /usr/local/go/src/encoding/asn1/asn1.go:443
	}() || func() bool {
//line /usr/local/go/src/encoding/asn1/asn1.go:443
		_go_fuzz_dep_.CoverTab[7653]++
//line /usr/local/go/src/encoding/asn1/asn1.go:443
		return (bool(ampersand) && func() bool {
//line /usr/local/go/src/encoding/asn1/asn1.go:448
			_go_fuzz_dep_.CoverTab[7654]++
//line /usr/local/go/src/encoding/asn1/asn1.go:448
			return b == '&'
//line /usr/local/go/src/encoding/asn1/asn1.go:448
			// _ = "end of CoverTab[7654]"
//line /usr/local/go/src/encoding/asn1/asn1.go:448
		}())
//line /usr/local/go/src/encoding/asn1/asn1.go:448
		// _ = "end of CoverTab[7653]"
//line /usr/local/go/src/encoding/asn1/asn1.go:448
	}()
//line /usr/local/go/src/encoding/asn1/asn1.go:448
	// _ = "end of CoverTab[7637]"
}

//line /usr/local/go/src/encoding/asn1/asn1.go:453
// parseIA5String parses an ASN.1 IA5String (ASCII string) from the given
//line /usr/local/go/src/encoding/asn1/asn1.go:453
// byte slice and returns it.
//line /usr/local/go/src/encoding/asn1/asn1.go:455
func parseIA5String(bytes []byte) (ret string, err error) {
//line /usr/local/go/src/encoding/asn1/asn1.go:455
	_go_fuzz_dep_.CoverTab[7655]++
							for _, b := range bytes {
//line /usr/local/go/src/encoding/asn1/asn1.go:456
		_go_fuzz_dep_.CoverTab[7657]++
								if b >= utf8.RuneSelf {
//line /usr/local/go/src/encoding/asn1/asn1.go:457
			_go_fuzz_dep_.CoverTab[7658]++
									err = SyntaxError{"IA5String contains invalid character"}
									return
//line /usr/local/go/src/encoding/asn1/asn1.go:459
			// _ = "end of CoverTab[7658]"
		} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:460
			_go_fuzz_dep_.CoverTab[7659]++
//line /usr/local/go/src/encoding/asn1/asn1.go:460
			// _ = "end of CoverTab[7659]"
//line /usr/local/go/src/encoding/asn1/asn1.go:460
		}
//line /usr/local/go/src/encoding/asn1/asn1.go:460
		// _ = "end of CoverTab[7657]"
	}
//line /usr/local/go/src/encoding/asn1/asn1.go:461
	// _ = "end of CoverTab[7655]"
//line /usr/local/go/src/encoding/asn1/asn1.go:461
	_go_fuzz_dep_.CoverTab[7656]++
							ret = string(bytes)
							return
//line /usr/local/go/src/encoding/asn1/asn1.go:463
	// _ = "end of CoverTab[7656]"
}

//line /usr/local/go/src/encoding/asn1/asn1.go:468
// parseT61String parses an ASN.1 T61String (8-bit clean string) from the given
//line /usr/local/go/src/encoding/asn1/asn1.go:468
// byte slice and returns it.
//line /usr/local/go/src/encoding/asn1/asn1.go:470
func parseT61String(bytes []byte) (ret string, err error) {
//line /usr/local/go/src/encoding/asn1/asn1.go:470
	_go_fuzz_dep_.CoverTab[7660]++
							return string(bytes), nil
//line /usr/local/go/src/encoding/asn1/asn1.go:471
	// _ = "end of CoverTab[7660]"
}

//line /usr/local/go/src/encoding/asn1/asn1.go:476
// parseUTF8String parses an ASN.1 UTF8String (raw UTF-8) from the given byte
//line /usr/local/go/src/encoding/asn1/asn1.go:476
// array and returns it.
//line /usr/local/go/src/encoding/asn1/asn1.go:478
func parseUTF8String(bytes []byte) (ret string, err error) {
//line /usr/local/go/src/encoding/asn1/asn1.go:478
	_go_fuzz_dep_.CoverTab[7661]++
							if !utf8.Valid(bytes) {
//line /usr/local/go/src/encoding/asn1/asn1.go:479
		_go_fuzz_dep_.CoverTab[7663]++
								return "", errors.New("asn1: invalid UTF-8 string")
//line /usr/local/go/src/encoding/asn1/asn1.go:480
		// _ = "end of CoverTab[7663]"
	} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:481
		_go_fuzz_dep_.CoverTab[7664]++
//line /usr/local/go/src/encoding/asn1/asn1.go:481
		// _ = "end of CoverTab[7664]"
//line /usr/local/go/src/encoding/asn1/asn1.go:481
	}
//line /usr/local/go/src/encoding/asn1/asn1.go:481
	// _ = "end of CoverTab[7661]"
//line /usr/local/go/src/encoding/asn1/asn1.go:481
	_go_fuzz_dep_.CoverTab[7662]++
							return string(bytes), nil
//line /usr/local/go/src/encoding/asn1/asn1.go:482
	// _ = "end of CoverTab[7662]"
}

//line /usr/local/go/src/encoding/asn1/asn1.go:487
// parseBMPString parses an ASN.1 BMPString (Basic Multilingual Plane of
//line /usr/local/go/src/encoding/asn1/asn1.go:487
// ISO/IEC/ITU 10646-1) from the given byte slice and returns it.
//line /usr/local/go/src/encoding/asn1/asn1.go:489
func parseBMPString(bmpString []byte) (string, error) {
//line /usr/local/go/src/encoding/asn1/asn1.go:489
	_go_fuzz_dep_.CoverTab[7665]++
							if len(bmpString)%2 != 0 {
//line /usr/local/go/src/encoding/asn1/asn1.go:490
		_go_fuzz_dep_.CoverTab[7669]++
								return "", errors.New("pkcs12: odd-length BMP string")
//line /usr/local/go/src/encoding/asn1/asn1.go:491
		// _ = "end of CoverTab[7669]"
	} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:492
		_go_fuzz_dep_.CoverTab[7670]++
//line /usr/local/go/src/encoding/asn1/asn1.go:492
		// _ = "end of CoverTab[7670]"
//line /usr/local/go/src/encoding/asn1/asn1.go:492
	}
//line /usr/local/go/src/encoding/asn1/asn1.go:492
	// _ = "end of CoverTab[7665]"
//line /usr/local/go/src/encoding/asn1/asn1.go:492
	_go_fuzz_dep_.CoverTab[7666]++

//line /usr/local/go/src/encoding/asn1/asn1.go:495
	if l := len(bmpString); l >= 2 && func() bool {
//line /usr/local/go/src/encoding/asn1/asn1.go:495
		_go_fuzz_dep_.CoverTab[7671]++
//line /usr/local/go/src/encoding/asn1/asn1.go:495
		return bmpString[l-1] == 0
//line /usr/local/go/src/encoding/asn1/asn1.go:495
		// _ = "end of CoverTab[7671]"
//line /usr/local/go/src/encoding/asn1/asn1.go:495
	}() && func() bool {
//line /usr/local/go/src/encoding/asn1/asn1.go:495
		_go_fuzz_dep_.CoverTab[7672]++
//line /usr/local/go/src/encoding/asn1/asn1.go:495
		return bmpString[l-2] == 0
//line /usr/local/go/src/encoding/asn1/asn1.go:495
		// _ = "end of CoverTab[7672]"
//line /usr/local/go/src/encoding/asn1/asn1.go:495
	}() {
//line /usr/local/go/src/encoding/asn1/asn1.go:495
		_go_fuzz_dep_.CoverTab[7673]++
								bmpString = bmpString[:l-2]
//line /usr/local/go/src/encoding/asn1/asn1.go:496
		// _ = "end of CoverTab[7673]"
	} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:497
		_go_fuzz_dep_.CoverTab[7674]++
//line /usr/local/go/src/encoding/asn1/asn1.go:497
		// _ = "end of CoverTab[7674]"
//line /usr/local/go/src/encoding/asn1/asn1.go:497
	}
//line /usr/local/go/src/encoding/asn1/asn1.go:497
	// _ = "end of CoverTab[7666]"
//line /usr/local/go/src/encoding/asn1/asn1.go:497
	_go_fuzz_dep_.CoverTab[7667]++

							s := make([]uint16, 0, len(bmpString)/2)
							for len(bmpString) > 0 {
//line /usr/local/go/src/encoding/asn1/asn1.go:500
		_go_fuzz_dep_.CoverTab[7675]++
								s = append(s, uint16(bmpString[0])<<8+uint16(bmpString[1]))
								bmpString = bmpString[2:]
//line /usr/local/go/src/encoding/asn1/asn1.go:502
		// _ = "end of CoverTab[7675]"
	}
//line /usr/local/go/src/encoding/asn1/asn1.go:503
	// _ = "end of CoverTab[7667]"
//line /usr/local/go/src/encoding/asn1/asn1.go:503
	_go_fuzz_dep_.CoverTab[7668]++

							return string(utf16.Decode(s)), nil
//line /usr/local/go/src/encoding/asn1/asn1.go:505
	// _ = "end of CoverTab[7668]"
}

// A RawValue represents an undecoded ASN.1 object.
type RawValue struct {
	Class, Tag	int
	IsCompound	bool
	Bytes		[]byte
	FullBytes	[]byte	// includes the tag and length
}

// RawContent is used to signal that the undecoded, DER data needs to be
//line /usr/local/go/src/encoding/asn1/asn1.go:516
// preserved for a struct. To use it, the first field of the struct must have
//line /usr/local/go/src/encoding/asn1/asn1.go:516
// this type. It's an error for any of the other fields to have this type.
//line /usr/local/go/src/encoding/asn1/asn1.go:519
type RawContent []byte

//line /usr/local/go/src/encoding/asn1/asn1.go:523
// parseTagAndLength parses an ASN.1 tag and length pair from the given offset
//line /usr/local/go/src/encoding/asn1/asn1.go:523
// into a byte slice. It returns the parsed data and the new offset. SET and
//line /usr/local/go/src/encoding/asn1/asn1.go:523
// SET OF (tag 17) are mapped to SEQUENCE and SEQUENCE OF (tag 16) since we
//line /usr/local/go/src/encoding/asn1/asn1.go:523
// don't distinguish between ordered and unordered objects in this code.
//line /usr/local/go/src/encoding/asn1/asn1.go:527
func parseTagAndLength(bytes []byte, initOffset int) (ret tagAndLength, offset int, err error) {
//line /usr/local/go/src/encoding/asn1/asn1.go:527
	_go_fuzz_dep_.CoverTab[7676]++
							offset = initOffset

//line /usr/local/go/src/encoding/asn1/asn1.go:531
	if offset >= len(bytes) {
//line /usr/local/go/src/encoding/asn1/asn1.go:531
		_go_fuzz_dep_.CoverTab[7681]++
								err = errors.New("asn1: internal error in parseTagAndLength")
								return
//line /usr/local/go/src/encoding/asn1/asn1.go:533
		// _ = "end of CoverTab[7681]"
	} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:534
		_go_fuzz_dep_.CoverTab[7682]++
//line /usr/local/go/src/encoding/asn1/asn1.go:534
		// _ = "end of CoverTab[7682]"
//line /usr/local/go/src/encoding/asn1/asn1.go:534
	}
//line /usr/local/go/src/encoding/asn1/asn1.go:534
	// _ = "end of CoverTab[7676]"
//line /usr/local/go/src/encoding/asn1/asn1.go:534
	_go_fuzz_dep_.CoverTab[7677]++
							b := bytes[offset]
							offset++
							ret.class = int(b >> 6)
							ret.isCompound = b&0x20 == 0x20
							ret.tag = int(b & 0x1f)

//line /usr/local/go/src/encoding/asn1/asn1.go:543
	if ret.tag == 0x1f {
//line /usr/local/go/src/encoding/asn1/asn1.go:543
		_go_fuzz_dep_.CoverTab[7683]++
								ret.tag, offset, err = parseBase128Int(bytes, offset)
								if err != nil {
//line /usr/local/go/src/encoding/asn1/asn1.go:545
			_go_fuzz_dep_.CoverTab[7685]++
									return
//line /usr/local/go/src/encoding/asn1/asn1.go:546
			// _ = "end of CoverTab[7685]"
		} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:547
			_go_fuzz_dep_.CoverTab[7686]++
//line /usr/local/go/src/encoding/asn1/asn1.go:547
			// _ = "end of CoverTab[7686]"
//line /usr/local/go/src/encoding/asn1/asn1.go:547
		}
//line /usr/local/go/src/encoding/asn1/asn1.go:547
		// _ = "end of CoverTab[7683]"
//line /usr/local/go/src/encoding/asn1/asn1.go:547
		_go_fuzz_dep_.CoverTab[7684]++

								if ret.tag < 0x1f {
//line /usr/local/go/src/encoding/asn1/asn1.go:549
			_go_fuzz_dep_.CoverTab[7687]++
									err = SyntaxError{"non-minimal tag"}
									return
//line /usr/local/go/src/encoding/asn1/asn1.go:551
			// _ = "end of CoverTab[7687]"
		} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:552
			_go_fuzz_dep_.CoverTab[7688]++
//line /usr/local/go/src/encoding/asn1/asn1.go:552
			// _ = "end of CoverTab[7688]"
//line /usr/local/go/src/encoding/asn1/asn1.go:552
		}
//line /usr/local/go/src/encoding/asn1/asn1.go:552
		// _ = "end of CoverTab[7684]"
	} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:553
		_go_fuzz_dep_.CoverTab[7689]++
//line /usr/local/go/src/encoding/asn1/asn1.go:553
		// _ = "end of CoverTab[7689]"
//line /usr/local/go/src/encoding/asn1/asn1.go:553
	}
//line /usr/local/go/src/encoding/asn1/asn1.go:553
	// _ = "end of CoverTab[7677]"
//line /usr/local/go/src/encoding/asn1/asn1.go:553
	_go_fuzz_dep_.CoverTab[7678]++
							if offset >= len(bytes) {
//line /usr/local/go/src/encoding/asn1/asn1.go:554
		_go_fuzz_dep_.CoverTab[7690]++
								err = SyntaxError{"truncated tag or length"}
								return
//line /usr/local/go/src/encoding/asn1/asn1.go:556
		// _ = "end of CoverTab[7690]"
	} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:557
		_go_fuzz_dep_.CoverTab[7691]++
//line /usr/local/go/src/encoding/asn1/asn1.go:557
		// _ = "end of CoverTab[7691]"
//line /usr/local/go/src/encoding/asn1/asn1.go:557
	}
//line /usr/local/go/src/encoding/asn1/asn1.go:557
	// _ = "end of CoverTab[7678]"
//line /usr/local/go/src/encoding/asn1/asn1.go:557
	_go_fuzz_dep_.CoverTab[7679]++
							b = bytes[offset]
							offset++
							if b&0x80 == 0 {
//line /usr/local/go/src/encoding/asn1/asn1.go:560
		_go_fuzz_dep_.CoverTab[7692]++

								ret.length = int(b & 0x7f)
//line /usr/local/go/src/encoding/asn1/asn1.go:562
		// _ = "end of CoverTab[7692]"
	} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:563
		_go_fuzz_dep_.CoverTab[7693]++

								numBytes := int(b & 0x7f)
								if numBytes == 0 {
//line /usr/local/go/src/encoding/asn1/asn1.go:566
			_go_fuzz_dep_.CoverTab[7696]++
									err = SyntaxError{"indefinite length found (not DER)"}
									return
//line /usr/local/go/src/encoding/asn1/asn1.go:568
			// _ = "end of CoverTab[7696]"
		} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:569
			_go_fuzz_dep_.CoverTab[7697]++
//line /usr/local/go/src/encoding/asn1/asn1.go:569
			// _ = "end of CoverTab[7697]"
//line /usr/local/go/src/encoding/asn1/asn1.go:569
		}
//line /usr/local/go/src/encoding/asn1/asn1.go:569
		// _ = "end of CoverTab[7693]"
//line /usr/local/go/src/encoding/asn1/asn1.go:569
		_go_fuzz_dep_.CoverTab[7694]++
								ret.length = 0
								for i := 0; i < numBytes; i++ {
//line /usr/local/go/src/encoding/asn1/asn1.go:571
			_go_fuzz_dep_.CoverTab[7698]++
									if offset >= len(bytes) {
//line /usr/local/go/src/encoding/asn1/asn1.go:572
				_go_fuzz_dep_.CoverTab[7701]++
										err = SyntaxError{"truncated tag or length"}
										return
//line /usr/local/go/src/encoding/asn1/asn1.go:574
				// _ = "end of CoverTab[7701]"
			} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:575
				_go_fuzz_dep_.CoverTab[7702]++
//line /usr/local/go/src/encoding/asn1/asn1.go:575
				// _ = "end of CoverTab[7702]"
//line /usr/local/go/src/encoding/asn1/asn1.go:575
			}
//line /usr/local/go/src/encoding/asn1/asn1.go:575
			// _ = "end of CoverTab[7698]"
//line /usr/local/go/src/encoding/asn1/asn1.go:575
			_go_fuzz_dep_.CoverTab[7699]++
									b = bytes[offset]
									offset++
									if ret.length >= 1<<23 {
//line /usr/local/go/src/encoding/asn1/asn1.go:578
				_go_fuzz_dep_.CoverTab[7703]++

//line /usr/local/go/src/encoding/asn1/asn1.go:581
				err = StructuralError{"length too large"}
										return
//line /usr/local/go/src/encoding/asn1/asn1.go:582
				// _ = "end of CoverTab[7703]"
			} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:583
				_go_fuzz_dep_.CoverTab[7704]++
//line /usr/local/go/src/encoding/asn1/asn1.go:583
				// _ = "end of CoverTab[7704]"
//line /usr/local/go/src/encoding/asn1/asn1.go:583
			}
//line /usr/local/go/src/encoding/asn1/asn1.go:583
			// _ = "end of CoverTab[7699]"
//line /usr/local/go/src/encoding/asn1/asn1.go:583
			_go_fuzz_dep_.CoverTab[7700]++
									ret.length <<= 8
									ret.length |= int(b)
									if ret.length == 0 {
//line /usr/local/go/src/encoding/asn1/asn1.go:586
				_go_fuzz_dep_.CoverTab[7705]++

										err = StructuralError{"superfluous leading zeros in length"}
										return
//line /usr/local/go/src/encoding/asn1/asn1.go:589
				// _ = "end of CoverTab[7705]"
			} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:590
				_go_fuzz_dep_.CoverTab[7706]++
//line /usr/local/go/src/encoding/asn1/asn1.go:590
				// _ = "end of CoverTab[7706]"
//line /usr/local/go/src/encoding/asn1/asn1.go:590
			}
//line /usr/local/go/src/encoding/asn1/asn1.go:590
			// _ = "end of CoverTab[7700]"
		}
//line /usr/local/go/src/encoding/asn1/asn1.go:591
		// _ = "end of CoverTab[7694]"
//line /usr/local/go/src/encoding/asn1/asn1.go:591
		_go_fuzz_dep_.CoverTab[7695]++

								if ret.length < 0x80 {
//line /usr/local/go/src/encoding/asn1/asn1.go:593
			_go_fuzz_dep_.CoverTab[7707]++
									err = StructuralError{"non-minimal length"}
									return
//line /usr/local/go/src/encoding/asn1/asn1.go:595
			// _ = "end of CoverTab[7707]"
		} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:596
			_go_fuzz_dep_.CoverTab[7708]++
//line /usr/local/go/src/encoding/asn1/asn1.go:596
			// _ = "end of CoverTab[7708]"
//line /usr/local/go/src/encoding/asn1/asn1.go:596
		}
//line /usr/local/go/src/encoding/asn1/asn1.go:596
		// _ = "end of CoverTab[7695]"
	}
//line /usr/local/go/src/encoding/asn1/asn1.go:597
	// _ = "end of CoverTab[7679]"
//line /usr/local/go/src/encoding/asn1/asn1.go:597
	_go_fuzz_dep_.CoverTab[7680]++

							return
//line /usr/local/go/src/encoding/asn1/asn1.go:599
	// _ = "end of CoverTab[7680]"
}

// parseSequenceOf is used for SEQUENCE OF and SET OF values. It tries to parse
//line /usr/local/go/src/encoding/asn1/asn1.go:602
// a number of ASN.1 values from the given byte slice and returns them as a
//line /usr/local/go/src/encoding/asn1/asn1.go:602
// slice of Go values of the given type.
//line /usr/local/go/src/encoding/asn1/asn1.go:605
func parseSequenceOf(bytes []byte, sliceType reflect.Type, elemType reflect.Type) (ret reflect.Value, err error) {
//line /usr/local/go/src/encoding/asn1/asn1.go:605
	_go_fuzz_dep_.CoverTab[7709]++
							matchAny, expectedTag, compoundType, ok := getUniversalType(elemType)
							if !ok {
//line /usr/local/go/src/encoding/asn1/asn1.go:607
		_go_fuzz_dep_.CoverTab[7713]++
								err = StructuralError{"unknown Go type for slice"}
								return
//line /usr/local/go/src/encoding/asn1/asn1.go:609
		// _ = "end of CoverTab[7713]"
	} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:610
		_go_fuzz_dep_.CoverTab[7714]++
//line /usr/local/go/src/encoding/asn1/asn1.go:610
		// _ = "end of CoverTab[7714]"
//line /usr/local/go/src/encoding/asn1/asn1.go:610
	}
//line /usr/local/go/src/encoding/asn1/asn1.go:610
	// _ = "end of CoverTab[7709]"
//line /usr/local/go/src/encoding/asn1/asn1.go:610
	_go_fuzz_dep_.CoverTab[7710]++

//line /usr/local/go/src/encoding/asn1/asn1.go:614
	numElements := 0
	for offset := 0; offset < len(bytes); {
//line /usr/local/go/src/encoding/asn1/asn1.go:615
		_go_fuzz_dep_.CoverTab[7715]++
								var t tagAndLength
								t, offset, err = parseTagAndLength(bytes, offset)
								if err != nil {
//line /usr/local/go/src/encoding/asn1/asn1.go:618
			_go_fuzz_dep_.CoverTab[7720]++
									return
//line /usr/local/go/src/encoding/asn1/asn1.go:619
			// _ = "end of CoverTab[7720]"
		} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:620
			_go_fuzz_dep_.CoverTab[7721]++
//line /usr/local/go/src/encoding/asn1/asn1.go:620
			// _ = "end of CoverTab[7721]"
//line /usr/local/go/src/encoding/asn1/asn1.go:620
		}
//line /usr/local/go/src/encoding/asn1/asn1.go:620
		// _ = "end of CoverTab[7715]"
//line /usr/local/go/src/encoding/asn1/asn1.go:620
		_go_fuzz_dep_.CoverTab[7716]++
								switch t.tag {
		case TagIA5String, TagGeneralString, TagT61String, TagUTF8String, TagNumericString, TagBMPString:
//line /usr/local/go/src/encoding/asn1/asn1.go:622
			_go_fuzz_dep_.CoverTab[7722]++

//line /usr/local/go/src/encoding/asn1/asn1.go:626
			t.tag = TagPrintableString
//line /usr/local/go/src/encoding/asn1/asn1.go:626
			// _ = "end of CoverTab[7722]"
		case TagGeneralizedTime, TagUTCTime:
//line /usr/local/go/src/encoding/asn1/asn1.go:627
			_go_fuzz_dep_.CoverTab[7723]++

									t.tag = TagUTCTime
//line /usr/local/go/src/encoding/asn1/asn1.go:629
			// _ = "end of CoverTab[7723]"
//line /usr/local/go/src/encoding/asn1/asn1.go:629
		default:
//line /usr/local/go/src/encoding/asn1/asn1.go:629
			_go_fuzz_dep_.CoverTab[7724]++
//line /usr/local/go/src/encoding/asn1/asn1.go:629
			// _ = "end of CoverTab[7724]"
		}
//line /usr/local/go/src/encoding/asn1/asn1.go:630
		// _ = "end of CoverTab[7716]"
//line /usr/local/go/src/encoding/asn1/asn1.go:630
		_go_fuzz_dep_.CoverTab[7717]++

								if !matchAny && func() bool {
//line /usr/local/go/src/encoding/asn1/asn1.go:632
			_go_fuzz_dep_.CoverTab[7725]++
//line /usr/local/go/src/encoding/asn1/asn1.go:632
			return (t.class != ClassUniversal || func() bool {
//line /usr/local/go/src/encoding/asn1/asn1.go:632
				_go_fuzz_dep_.CoverTab[7726]++
//line /usr/local/go/src/encoding/asn1/asn1.go:632
				return t.isCompound != compoundType
//line /usr/local/go/src/encoding/asn1/asn1.go:632
				// _ = "end of CoverTab[7726]"
//line /usr/local/go/src/encoding/asn1/asn1.go:632
			}() || func() bool {
//line /usr/local/go/src/encoding/asn1/asn1.go:632
				_go_fuzz_dep_.CoverTab[7727]++
//line /usr/local/go/src/encoding/asn1/asn1.go:632
				return t.tag != expectedTag
//line /usr/local/go/src/encoding/asn1/asn1.go:632
				// _ = "end of CoverTab[7727]"
//line /usr/local/go/src/encoding/asn1/asn1.go:632
			}())
//line /usr/local/go/src/encoding/asn1/asn1.go:632
			// _ = "end of CoverTab[7725]"
//line /usr/local/go/src/encoding/asn1/asn1.go:632
		}() {
//line /usr/local/go/src/encoding/asn1/asn1.go:632
			_go_fuzz_dep_.CoverTab[7728]++
									err = StructuralError{"sequence tag mismatch"}
									return
//line /usr/local/go/src/encoding/asn1/asn1.go:634
			// _ = "end of CoverTab[7728]"
		} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:635
			_go_fuzz_dep_.CoverTab[7729]++
//line /usr/local/go/src/encoding/asn1/asn1.go:635
			// _ = "end of CoverTab[7729]"
//line /usr/local/go/src/encoding/asn1/asn1.go:635
		}
//line /usr/local/go/src/encoding/asn1/asn1.go:635
		// _ = "end of CoverTab[7717]"
//line /usr/local/go/src/encoding/asn1/asn1.go:635
		_go_fuzz_dep_.CoverTab[7718]++
								if invalidLength(offset, t.length, len(bytes)) {
//line /usr/local/go/src/encoding/asn1/asn1.go:636
			_go_fuzz_dep_.CoverTab[7730]++
									err = SyntaxError{"truncated sequence"}
									return
//line /usr/local/go/src/encoding/asn1/asn1.go:638
			// _ = "end of CoverTab[7730]"
		} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:639
			_go_fuzz_dep_.CoverTab[7731]++
//line /usr/local/go/src/encoding/asn1/asn1.go:639
			// _ = "end of CoverTab[7731]"
//line /usr/local/go/src/encoding/asn1/asn1.go:639
		}
//line /usr/local/go/src/encoding/asn1/asn1.go:639
		// _ = "end of CoverTab[7718]"
//line /usr/local/go/src/encoding/asn1/asn1.go:639
		_go_fuzz_dep_.CoverTab[7719]++
								offset += t.length
								numElements++
//line /usr/local/go/src/encoding/asn1/asn1.go:641
		// _ = "end of CoverTab[7719]"
	}
//line /usr/local/go/src/encoding/asn1/asn1.go:642
	// _ = "end of CoverTab[7710]"
//line /usr/local/go/src/encoding/asn1/asn1.go:642
	_go_fuzz_dep_.CoverTab[7711]++
							ret = reflect.MakeSlice(sliceType, numElements, numElements)
							params := fieldParameters{}
							offset := 0
							for i := 0; i < numElements; i++ {
//line /usr/local/go/src/encoding/asn1/asn1.go:646
		_go_fuzz_dep_.CoverTab[7732]++
								offset, err = parseField(ret.Index(i), bytes, offset, params)
								if err != nil {
//line /usr/local/go/src/encoding/asn1/asn1.go:648
			_go_fuzz_dep_.CoverTab[7733]++
									return
//line /usr/local/go/src/encoding/asn1/asn1.go:649
			// _ = "end of CoverTab[7733]"
		} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:650
			_go_fuzz_dep_.CoverTab[7734]++
//line /usr/local/go/src/encoding/asn1/asn1.go:650
			// _ = "end of CoverTab[7734]"
//line /usr/local/go/src/encoding/asn1/asn1.go:650
		}
//line /usr/local/go/src/encoding/asn1/asn1.go:650
		// _ = "end of CoverTab[7732]"
	}
//line /usr/local/go/src/encoding/asn1/asn1.go:651
	// _ = "end of CoverTab[7711]"
//line /usr/local/go/src/encoding/asn1/asn1.go:651
	_go_fuzz_dep_.CoverTab[7712]++
							return
//line /usr/local/go/src/encoding/asn1/asn1.go:652
	// _ = "end of CoverTab[7712]"
}

var (
	bitStringType		= reflect.TypeOf(BitString{})
	objectIdentifierType	= reflect.TypeOf(ObjectIdentifier{})
	enumeratedType		= reflect.TypeOf(Enumerated(0))
	flagType		= reflect.TypeOf(Flag(false))
	timeType		= reflect.TypeOf(time.Time{})
	rawValueType		= reflect.TypeOf(RawValue{})
	rawContentsType		= reflect.TypeOf(RawContent(nil))
	bigIntType		= reflect.TypeOf((*big.Int)(nil))
)

// invalidLength reports whether offset + length > sliceLength, or if the
//line /usr/local/go/src/encoding/asn1/asn1.go:666
// addition would overflow.
//line /usr/local/go/src/encoding/asn1/asn1.go:668
func invalidLength(offset, length, sliceLength int) bool {
//line /usr/local/go/src/encoding/asn1/asn1.go:668
	_go_fuzz_dep_.CoverTab[7735]++
							return offset+length < offset || func() bool {
//line /usr/local/go/src/encoding/asn1/asn1.go:669
		_go_fuzz_dep_.CoverTab[7736]++
//line /usr/local/go/src/encoding/asn1/asn1.go:669
		return offset+length > sliceLength
//line /usr/local/go/src/encoding/asn1/asn1.go:669
		// _ = "end of CoverTab[7736]"
//line /usr/local/go/src/encoding/asn1/asn1.go:669
	}()
//line /usr/local/go/src/encoding/asn1/asn1.go:669
	// _ = "end of CoverTab[7735]"
}

// parseField is the main parsing function. Given a byte slice and an offset
//line /usr/local/go/src/encoding/asn1/asn1.go:672
// into the array, it will try to parse a suitable ASN.1 value out and store it
//line /usr/local/go/src/encoding/asn1/asn1.go:672
// in the given Value.
//line /usr/local/go/src/encoding/asn1/asn1.go:675
func parseField(v reflect.Value, bytes []byte, initOffset int, params fieldParameters) (offset int, err error) {
//line /usr/local/go/src/encoding/asn1/asn1.go:675
	_go_fuzz_dep_.CoverTab[7737]++
							offset = initOffset
							fieldType := v.Type()

//line /usr/local/go/src/encoding/asn1/asn1.go:680
	if offset == len(bytes) {
//line /usr/local/go/src/encoding/asn1/asn1.go:680
		_go_fuzz_dep_.CoverTab[7753]++
								if !setDefaultValue(v, params) {
//line /usr/local/go/src/encoding/asn1/asn1.go:681
			_go_fuzz_dep_.CoverTab[7755]++
									err = SyntaxError{"sequence truncated"}
//line /usr/local/go/src/encoding/asn1/asn1.go:682
			// _ = "end of CoverTab[7755]"
		} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:683
			_go_fuzz_dep_.CoverTab[7756]++
//line /usr/local/go/src/encoding/asn1/asn1.go:683
			// _ = "end of CoverTab[7756]"
//line /usr/local/go/src/encoding/asn1/asn1.go:683
		}
//line /usr/local/go/src/encoding/asn1/asn1.go:683
		// _ = "end of CoverTab[7753]"
//line /usr/local/go/src/encoding/asn1/asn1.go:683
		_go_fuzz_dep_.CoverTab[7754]++
								return
//line /usr/local/go/src/encoding/asn1/asn1.go:684
		// _ = "end of CoverTab[7754]"
	} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:685
		_go_fuzz_dep_.CoverTab[7757]++
//line /usr/local/go/src/encoding/asn1/asn1.go:685
		// _ = "end of CoverTab[7757]"
//line /usr/local/go/src/encoding/asn1/asn1.go:685
	}
//line /usr/local/go/src/encoding/asn1/asn1.go:685
	// _ = "end of CoverTab[7737]"
//line /usr/local/go/src/encoding/asn1/asn1.go:685
	_go_fuzz_dep_.CoverTab[7738]++

//line /usr/local/go/src/encoding/asn1/asn1.go:688
	if ifaceType := fieldType; ifaceType.Kind() == reflect.Interface && func() bool {
//line /usr/local/go/src/encoding/asn1/asn1.go:688
		_go_fuzz_dep_.CoverTab[7758]++
//line /usr/local/go/src/encoding/asn1/asn1.go:688
		return ifaceType.NumMethod() == 0
//line /usr/local/go/src/encoding/asn1/asn1.go:688
		// _ = "end of CoverTab[7758]"
//line /usr/local/go/src/encoding/asn1/asn1.go:688
	}() {
//line /usr/local/go/src/encoding/asn1/asn1.go:688
		_go_fuzz_dep_.CoverTab[7759]++
								var t tagAndLength
								t, offset, err = parseTagAndLength(bytes, offset)
								if err != nil {
//line /usr/local/go/src/encoding/asn1/asn1.go:691
			_go_fuzz_dep_.CoverTab[7765]++
									return
//line /usr/local/go/src/encoding/asn1/asn1.go:692
			// _ = "end of CoverTab[7765]"
		} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:693
			_go_fuzz_dep_.CoverTab[7766]++
//line /usr/local/go/src/encoding/asn1/asn1.go:693
			// _ = "end of CoverTab[7766]"
//line /usr/local/go/src/encoding/asn1/asn1.go:693
		}
//line /usr/local/go/src/encoding/asn1/asn1.go:693
		// _ = "end of CoverTab[7759]"
//line /usr/local/go/src/encoding/asn1/asn1.go:693
		_go_fuzz_dep_.CoverTab[7760]++
								if invalidLength(offset, t.length, len(bytes)) {
//line /usr/local/go/src/encoding/asn1/asn1.go:694
			_go_fuzz_dep_.CoverTab[7767]++
									err = SyntaxError{"data truncated"}
									return
//line /usr/local/go/src/encoding/asn1/asn1.go:696
			// _ = "end of CoverTab[7767]"
		} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:697
			_go_fuzz_dep_.CoverTab[7768]++
//line /usr/local/go/src/encoding/asn1/asn1.go:697
			// _ = "end of CoverTab[7768]"
//line /usr/local/go/src/encoding/asn1/asn1.go:697
		}
//line /usr/local/go/src/encoding/asn1/asn1.go:697
		// _ = "end of CoverTab[7760]"
//line /usr/local/go/src/encoding/asn1/asn1.go:697
		_go_fuzz_dep_.CoverTab[7761]++
								var result any
								if !t.isCompound && func() bool {
//line /usr/local/go/src/encoding/asn1/asn1.go:699
			_go_fuzz_dep_.CoverTab[7769]++
//line /usr/local/go/src/encoding/asn1/asn1.go:699
			return t.class == ClassUniversal
//line /usr/local/go/src/encoding/asn1/asn1.go:699
			// _ = "end of CoverTab[7769]"
//line /usr/local/go/src/encoding/asn1/asn1.go:699
		}() {
//line /usr/local/go/src/encoding/asn1/asn1.go:699
			_go_fuzz_dep_.CoverTab[7770]++
									innerBytes := bytes[offset : offset+t.length]
									switch t.tag {
			case TagPrintableString:
//line /usr/local/go/src/encoding/asn1/asn1.go:702
				_go_fuzz_dep_.CoverTab[7771]++
										result, err = parsePrintableString(innerBytes)
//line /usr/local/go/src/encoding/asn1/asn1.go:703
				// _ = "end of CoverTab[7771]"
			case TagNumericString:
//line /usr/local/go/src/encoding/asn1/asn1.go:704
				_go_fuzz_dep_.CoverTab[7772]++
										result, err = parseNumericString(innerBytes)
//line /usr/local/go/src/encoding/asn1/asn1.go:705
				// _ = "end of CoverTab[7772]"
			case TagIA5String:
//line /usr/local/go/src/encoding/asn1/asn1.go:706
				_go_fuzz_dep_.CoverTab[7773]++
										result, err = parseIA5String(innerBytes)
//line /usr/local/go/src/encoding/asn1/asn1.go:707
				// _ = "end of CoverTab[7773]"
			case TagT61String:
//line /usr/local/go/src/encoding/asn1/asn1.go:708
				_go_fuzz_dep_.CoverTab[7774]++
										result, err = parseT61String(innerBytes)
//line /usr/local/go/src/encoding/asn1/asn1.go:709
				// _ = "end of CoverTab[7774]"
			case TagUTF8String:
//line /usr/local/go/src/encoding/asn1/asn1.go:710
				_go_fuzz_dep_.CoverTab[7775]++
										result, err = parseUTF8String(innerBytes)
//line /usr/local/go/src/encoding/asn1/asn1.go:711
				// _ = "end of CoverTab[7775]"
			case TagInteger:
//line /usr/local/go/src/encoding/asn1/asn1.go:712
				_go_fuzz_dep_.CoverTab[7776]++
										result, err = parseInt64(innerBytes)
//line /usr/local/go/src/encoding/asn1/asn1.go:713
				// _ = "end of CoverTab[7776]"
			case TagBitString:
//line /usr/local/go/src/encoding/asn1/asn1.go:714
				_go_fuzz_dep_.CoverTab[7777]++
										result, err = parseBitString(innerBytes)
//line /usr/local/go/src/encoding/asn1/asn1.go:715
				// _ = "end of CoverTab[7777]"
			case TagOID:
//line /usr/local/go/src/encoding/asn1/asn1.go:716
				_go_fuzz_dep_.CoverTab[7778]++
										result, err = parseObjectIdentifier(innerBytes)
//line /usr/local/go/src/encoding/asn1/asn1.go:717
				// _ = "end of CoverTab[7778]"
			case TagUTCTime:
//line /usr/local/go/src/encoding/asn1/asn1.go:718
				_go_fuzz_dep_.CoverTab[7779]++
										result, err = parseUTCTime(innerBytes)
//line /usr/local/go/src/encoding/asn1/asn1.go:719
				// _ = "end of CoverTab[7779]"
			case TagGeneralizedTime:
//line /usr/local/go/src/encoding/asn1/asn1.go:720
				_go_fuzz_dep_.CoverTab[7780]++
										result, err = parseGeneralizedTime(innerBytes)
//line /usr/local/go/src/encoding/asn1/asn1.go:721
				// _ = "end of CoverTab[7780]"
			case TagOctetString:
//line /usr/local/go/src/encoding/asn1/asn1.go:722
				_go_fuzz_dep_.CoverTab[7781]++
										result = innerBytes
//line /usr/local/go/src/encoding/asn1/asn1.go:723
				// _ = "end of CoverTab[7781]"
			case TagBMPString:
//line /usr/local/go/src/encoding/asn1/asn1.go:724
				_go_fuzz_dep_.CoverTab[7782]++
										result, err = parseBMPString(innerBytes)
//line /usr/local/go/src/encoding/asn1/asn1.go:725
				// _ = "end of CoverTab[7782]"
			default:
//line /usr/local/go/src/encoding/asn1/asn1.go:726
				_go_fuzz_dep_.CoverTab[7783]++
//line /usr/local/go/src/encoding/asn1/asn1.go:726
				// _ = "end of CoverTab[7783]"

			}
//line /usr/local/go/src/encoding/asn1/asn1.go:728
			// _ = "end of CoverTab[7770]"
		} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:729
			_go_fuzz_dep_.CoverTab[7784]++
//line /usr/local/go/src/encoding/asn1/asn1.go:729
			// _ = "end of CoverTab[7784]"
//line /usr/local/go/src/encoding/asn1/asn1.go:729
		}
//line /usr/local/go/src/encoding/asn1/asn1.go:729
		// _ = "end of CoverTab[7761]"
//line /usr/local/go/src/encoding/asn1/asn1.go:729
		_go_fuzz_dep_.CoverTab[7762]++
								offset += t.length
								if err != nil {
//line /usr/local/go/src/encoding/asn1/asn1.go:731
			_go_fuzz_dep_.CoverTab[7785]++
									return
//line /usr/local/go/src/encoding/asn1/asn1.go:732
			// _ = "end of CoverTab[7785]"
		} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:733
			_go_fuzz_dep_.CoverTab[7786]++
//line /usr/local/go/src/encoding/asn1/asn1.go:733
			// _ = "end of CoverTab[7786]"
//line /usr/local/go/src/encoding/asn1/asn1.go:733
		}
//line /usr/local/go/src/encoding/asn1/asn1.go:733
		// _ = "end of CoverTab[7762]"
//line /usr/local/go/src/encoding/asn1/asn1.go:733
		_go_fuzz_dep_.CoverTab[7763]++
								if result != nil {
//line /usr/local/go/src/encoding/asn1/asn1.go:734
			_go_fuzz_dep_.CoverTab[7787]++
									v.Set(reflect.ValueOf(result))
//line /usr/local/go/src/encoding/asn1/asn1.go:735
			// _ = "end of CoverTab[7787]"
		} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:736
			_go_fuzz_dep_.CoverTab[7788]++
//line /usr/local/go/src/encoding/asn1/asn1.go:736
			// _ = "end of CoverTab[7788]"
//line /usr/local/go/src/encoding/asn1/asn1.go:736
		}
//line /usr/local/go/src/encoding/asn1/asn1.go:736
		// _ = "end of CoverTab[7763]"
//line /usr/local/go/src/encoding/asn1/asn1.go:736
		_go_fuzz_dep_.CoverTab[7764]++
								return
//line /usr/local/go/src/encoding/asn1/asn1.go:737
		// _ = "end of CoverTab[7764]"
	} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:738
		_go_fuzz_dep_.CoverTab[7789]++
//line /usr/local/go/src/encoding/asn1/asn1.go:738
		// _ = "end of CoverTab[7789]"
//line /usr/local/go/src/encoding/asn1/asn1.go:738
	}
//line /usr/local/go/src/encoding/asn1/asn1.go:738
	// _ = "end of CoverTab[7738]"
//line /usr/local/go/src/encoding/asn1/asn1.go:738
	_go_fuzz_dep_.CoverTab[7739]++

							t, offset, err := parseTagAndLength(bytes, offset)
							if err != nil {
//line /usr/local/go/src/encoding/asn1/asn1.go:741
		_go_fuzz_dep_.CoverTab[7790]++
								return
//line /usr/local/go/src/encoding/asn1/asn1.go:742
		// _ = "end of CoverTab[7790]"
	} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:743
		_go_fuzz_dep_.CoverTab[7791]++
//line /usr/local/go/src/encoding/asn1/asn1.go:743
		// _ = "end of CoverTab[7791]"
//line /usr/local/go/src/encoding/asn1/asn1.go:743
	}
//line /usr/local/go/src/encoding/asn1/asn1.go:743
	// _ = "end of CoverTab[7739]"
//line /usr/local/go/src/encoding/asn1/asn1.go:743
	_go_fuzz_dep_.CoverTab[7740]++
							if params.explicit {
//line /usr/local/go/src/encoding/asn1/asn1.go:744
		_go_fuzz_dep_.CoverTab[7792]++
								expectedClass := ClassContextSpecific
								if params.application {
//line /usr/local/go/src/encoding/asn1/asn1.go:746
			_go_fuzz_dep_.CoverTab[7795]++
									expectedClass = ClassApplication
//line /usr/local/go/src/encoding/asn1/asn1.go:747
			// _ = "end of CoverTab[7795]"
		} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:748
			_go_fuzz_dep_.CoverTab[7796]++
//line /usr/local/go/src/encoding/asn1/asn1.go:748
			// _ = "end of CoverTab[7796]"
//line /usr/local/go/src/encoding/asn1/asn1.go:748
		}
//line /usr/local/go/src/encoding/asn1/asn1.go:748
		// _ = "end of CoverTab[7792]"
//line /usr/local/go/src/encoding/asn1/asn1.go:748
		_go_fuzz_dep_.CoverTab[7793]++
								if offset == len(bytes) {
//line /usr/local/go/src/encoding/asn1/asn1.go:749
			_go_fuzz_dep_.CoverTab[7797]++
									err = StructuralError{"explicit tag has no child"}
									return
//line /usr/local/go/src/encoding/asn1/asn1.go:751
			// _ = "end of CoverTab[7797]"
		} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:752
			_go_fuzz_dep_.CoverTab[7798]++
//line /usr/local/go/src/encoding/asn1/asn1.go:752
			// _ = "end of CoverTab[7798]"
//line /usr/local/go/src/encoding/asn1/asn1.go:752
		}
//line /usr/local/go/src/encoding/asn1/asn1.go:752
		// _ = "end of CoverTab[7793]"
//line /usr/local/go/src/encoding/asn1/asn1.go:752
		_go_fuzz_dep_.CoverTab[7794]++
								if t.class == expectedClass && func() bool {
//line /usr/local/go/src/encoding/asn1/asn1.go:753
			_go_fuzz_dep_.CoverTab[7799]++
//line /usr/local/go/src/encoding/asn1/asn1.go:753
			return t.tag == *params.tag
//line /usr/local/go/src/encoding/asn1/asn1.go:753
			// _ = "end of CoverTab[7799]"
//line /usr/local/go/src/encoding/asn1/asn1.go:753
		}() && func() bool {
//line /usr/local/go/src/encoding/asn1/asn1.go:753
			_go_fuzz_dep_.CoverTab[7800]++
//line /usr/local/go/src/encoding/asn1/asn1.go:753
			return (t.length == 0 || func() bool {
//line /usr/local/go/src/encoding/asn1/asn1.go:753
				_go_fuzz_dep_.CoverTab[7801]++
//line /usr/local/go/src/encoding/asn1/asn1.go:753
				return t.isCompound
//line /usr/local/go/src/encoding/asn1/asn1.go:753
				// _ = "end of CoverTab[7801]"
//line /usr/local/go/src/encoding/asn1/asn1.go:753
			}())
//line /usr/local/go/src/encoding/asn1/asn1.go:753
			// _ = "end of CoverTab[7800]"
//line /usr/local/go/src/encoding/asn1/asn1.go:753
		}() {
//line /usr/local/go/src/encoding/asn1/asn1.go:753
			_go_fuzz_dep_.CoverTab[7802]++
									if fieldType == rawValueType {
//line /usr/local/go/src/encoding/asn1/asn1.go:754
				_go_fuzz_dep_.CoverTab[7803]++
//line /usr/local/go/src/encoding/asn1/asn1.go:754
				// _ = "end of CoverTab[7803]"

			} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:756
				_go_fuzz_dep_.CoverTab[7804]++
//line /usr/local/go/src/encoding/asn1/asn1.go:756
				if t.length > 0 {
//line /usr/local/go/src/encoding/asn1/asn1.go:756
					_go_fuzz_dep_.CoverTab[7805]++
											t, offset, err = parseTagAndLength(bytes, offset)
											if err != nil {
//line /usr/local/go/src/encoding/asn1/asn1.go:758
						_go_fuzz_dep_.CoverTab[7806]++
												return
//line /usr/local/go/src/encoding/asn1/asn1.go:759
						// _ = "end of CoverTab[7806]"
					} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:760
						_go_fuzz_dep_.CoverTab[7807]++
//line /usr/local/go/src/encoding/asn1/asn1.go:760
						// _ = "end of CoverTab[7807]"
//line /usr/local/go/src/encoding/asn1/asn1.go:760
					}
//line /usr/local/go/src/encoding/asn1/asn1.go:760
					// _ = "end of CoverTab[7805]"
				} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:761
					_go_fuzz_dep_.CoverTab[7808]++
											if fieldType != flagType {
//line /usr/local/go/src/encoding/asn1/asn1.go:762
						_go_fuzz_dep_.CoverTab[7810]++
												err = StructuralError{"zero length explicit tag was not an asn1.Flag"}
												return
//line /usr/local/go/src/encoding/asn1/asn1.go:764
						// _ = "end of CoverTab[7810]"
					} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:765
						_go_fuzz_dep_.CoverTab[7811]++
//line /usr/local/go/src/encoding/asn1/asn1.go:765
						// _ = "end of CoverTab[7811]"
//line /usr/local/go/src/encoding/asn1/asn1.go:765
					}
//line /usr/local/go/src/encoding/asn1/asn1.go:765
					// _ = "end of CoverTab[7808]"
//line /usr/local/go/src/encoding/asn1/asn1.go:765
					_go_fuzz_dep_.CoverTab[7809]++
											v.SetBool(true)
											return
//line /usr/local/go/src/encoding/asn1/asn1.go:767
					// _ = "end of CoverTab[7809]"
				}
//line /usr/local/go/src/encoding/asn1/asn1.go:768
				// _ = "end of CoverTab[7804]"
//line /usr/local/go/src/encoding/asn1/asn1.go:768
			}
//line /usr/local/go/src/encoding/asn1/asn1.go:768
			// _ = "end of CoverTab[7802]"
		} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:769
			_go_fuzz_dep_.CoverTab[7812]++

									ok := setDefaultValue(v, params)
									if ok {
//line /usr/local/go/src/encoding/asn1/asn1.go:772
				_go_fuzz_dep_.CoverTab[7814]++
										offset = initOffset
//line /usr/local/go/src/encoding/asn1/asn1.go:773
				// _ = "end of CoverTab[7814]"
			} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:774
				_go_fuzz_dep_.CoverTab[7815]++
										err = StructuralError{"explicitly tagged member didn't match"}
//line /usr/local/go/src/encoding/asn1/asn1.go:775
				// _ = "end of CoverTab[7815]"
			}
//line /usr/local/go/src/encoding/asn1/asn1.go:776
			// _ = "end of CoverTab[7812]"
//line /usr/local/go/src/encoding/asn1/asn1.go:776
			_go_fuzz_dep_.CoverTab[7813]++
									return
//line /usr/local/go/src/encoding/asn1/asn1.go:777
			// _ = "end of CoverTab[7813]"
		}
//line /usr/local/go/src/encoding/asn1/asn1.go:778
		// _ = "end of CoverTab[7794]"
	} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:779
		_go_fuzz_dep_.CoverTab[7816]++
//line /usr/local/go/src/encoding/asn1/asn1.go:779
		// _ = "end of CoverTab[7816]"
//line /usr/local/go/src/encoding/asn1/asn1.go:779
	}
//line /usr/local/go/src/encoding/asn1/asn1.go:779
	// _ = "end of CoverTab[7740]"
//line /usr/local/go/src/encoding/asn1/asn1.go:779
	_go_fuzz_dep_.CoverTab[7741]++

							matchAny, universalTag, compoundType, ok1 := getUniversalType(fieldType)
							if !ok1 {
//line /usr/local/go/src/encoding/asn1/asn1.go:782
		_go_fuzz_dep_.CoverTab[7817]++
								err = StructuralError{fmt.Sprintf("unknown Go type: %v", fieldType)}
								return
//line /usr/local/go/src/encoding/asn1/asn1.go:784
		// _ = "end of CoverTab[7817]"
	} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:785
		_go_fuzz_dep_.CoverTab[7818]++
//line /usr/local/go/src/encoding/asn1/asn1.go:785
		// _ = "end of CoverTab[7818]"
//line /usr/local/go/src/encoding/asn1/asn1.go:785
	}
//line /usr/local/go/src/encoding/asn1/asn1.go:785
	// _ = "end of CoverTab[7741]"
//line /usr/local/go/src/encoding/asn1/asn1.go:785
	_go_fuzz_dep_.CoverTab[7742]++

//line /usr/local/go/src/encoding/asn1/asn1.go:791
	if universalTag == TagPrintableString {
//line /usr/local/go/src/encoding/asn1/asn1.go:791
		_go_fuzz_dep_.CoverTab[7819]++
								if t.class == ClassUniversal {
//line /usr/local/go/src/encoding/asn1/asn1.go:792
			_go_fuzz_dep_.CoverTab[7820]++
									switch t.tag {
			case TagIA5String, TagGeneralString, TagT61String, TagUTF8String, TagNumericString, TagBMPString:
//line /usr/local/go/src/encoding/asn1/asn1.go:794
				_go_fuzz_dep_.CoverTab[7821]++
										universalTag = t.tag
//line /usr/local/go/src/encoding/asn1/asn1.go:795
				// _ = "end of CoverTab[7821]"
//line /usr/local/go/src/encoding/asn1/asn1.go:795
			default:
//line /usr/local/go/src/encoding/asn1/asn1.go:795
				_go_fuzz_dep_.CoverTab[7822]++
//line /usr/local/go/src/encoding/asn1/asn1.go:795
				// _ = "end of CoverTab[7822]"
			}
//line /usr/local/go/src/encoding/asn1/asn1.go:796
			// _ = "end of CoverTab[7820]"
		} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:797
			_go_fuzz_dep_.CoverTab[7823]++
//line /usr/local/go/src/encoding/asn1/asn1.go:797
			if params.stringType != 0 {
//line /usr/local/go/src/encoding/asn1/asn1.go:797
				_go_fuzz_dep_.CoverTab[7824]++
										universalTag = params.stringType
//line /usr/local/go/src/encoding/asn1/asn1.go:798
				// _ = "end of CoverTab[7824]"
			} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:799
				_go_fuzz_dep_.CoverTab[7825]++
//line /usr/local/go/src/encoding/asn1/asn1.go:799
				// _ = "end of CoverTab[7825]"
//line /usr/local/go/src/encoding/asn1/asn1.go:799
			}
//line /usr/local/go/src/encoding/asn1/asn1.go:799
			// _ = "end of CoverTab[7823]"
//line /usr/local/go/src/encoding/asn1/asn1.go:799
		}
//line /usr/local/go/src/encoding/asn1/asn1.go:799
		// _ = "end of CoverTab[7819]"
	} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:800
		_go_fuzz_dep_.CoverTab[7826]++
//line /usr/local/go/src/encoding/asn1/asn1.go:800
		// _ = "end of CoverTab[7826]"
//line /usr/local/go/src/encoding/asn1/asn1.go:800
	}
//line /usr/local/go/src/encoding/asn1/asn1.go:800
	// _ = "end of CoverTab[7742]"
//line /usr/local/go/src/encoding/asn1/asn1.go:800
	_go_fuzz_dep_.CoverTab[7743]++

//line /usr/local/go/src/encoding/asn1/asn1.go:804
	if universalTag == TagUTCTime && func() bool {
//line /usr/local/go/src/encoding/asn1/asn1.go:804
		_go_fuzz_dep_.CoverTab[7827]++
//line /usr/local/go/src/encoding/asn1/asn1.go:804
		return t.tag == TagGeneralizedTime
//line /usr/local/go/src/encoding/asn1/asn1.go:804
		// _ = "end of CoverTab[7827]"
//line /usr/local/go/src/encoding/asn1/asn1.go:804
	}() && func() bool {
//line /usr/local/go/src/encoding/asn1/asn1.go:804
		_go_fuzz_dep_.CoverTab[7828]++
//line /usr/local/go/src/encoding/asn1/asn1.go:804
		return t.class == ClassUniversal
//line /usr/local/go/src/encoding/asn1/asn1.go:804
		// _ = "end of CoverTab[7828]"
//line /usr/local/go/src/encoding/asn1/asn1.go:804
	}() {
//line /usr/local/go/src/encoding/asn1/asn1.go:804
		_go_fuzz_dep_.CoverTab[7829]++
								universalTag = TagGeneralizedTime
//line /usr/local/go/src/encoding/asn1/asn1.go:805
		// _ = "end of CoverTab[7829]"
	} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:806
		_go_fuzz_dep_.CoverTab[7830]++
//line /usr/local/go/src/encoding/asn1/asn1.go:806
		// _ = "end of CoverTab[7830]"
//line /usr/local/go/src/encoding/asn1/asn1.go:806
	}
//line /usr/local/go/src/encoding/asn1/asn1.go:806
	// _ = "end of CoverTab[7743]"
//line /usr/local/go/src/encoding/asn1/asn1.go:806
	_go_fuzz_dep_.CoverTab[7744]++

							if params.set {
//line /usr/local/go/src/encoding/asn1/asn1.go:808
		_go_fuzz_dep_.CoverTab[7831]++
								universalTag = TagSet
//line /usr/local/go/src/encoding/asn1/asn1.go:809
		// _ = "end of CoverTab[7831]"
	} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:810
		_go_fuzz_dep_.CoverTab[7832]++
//line /usr/local/go/src/encoding/asn1/asn1.go:810
		// _ = "end of CoverTab[7832]"
//line /usr/local/go/src/encoding/asn1/asn1.go:810
	}
//line /usr/local/go/src/encoding/asn1/asn1.go:810
	// _ = "end of CoverTab[7744]"
//line /usr/local/go/src/encoding/asn1/asn1.go:810
	_go_fuzz_dep_.CoverTab[7745]++

							matchAnyClassAndTag := matchAny
							expectedClass := ClassUniversal
							expectedTag := universalTag

							if !params.explicit && func() bool {
//line /usr/local/go/src/encoding/asn1/asn1.go:816
		_go_fuzz_dep_.CoverTab[7833]++
//line /usr/local/go/src/encoding/asn1/asn1.go:816
		return params.tag != nil
//line /usr/local/go/src/encoding/asn1/asn1.go:816
		// _ = "end of CoverTab[7833]"
//line /usr/local/go/src/encoding/asn1/asn1.go:816
	}() {
//line /usr/local/go/src/encoding/asn1/asn1.go:816
		_go_fuzz_dep_.CoverTab[7834]++
								expectedClass = ClassContextSpecific
								expectedTag = *params.tag
								matchAnyClassAndTag = false
//line /usr/local/go/src/encoding/asn1/asn1.go:819
		// _ = "end of CoverTab[7834]"
	} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:820
		_go_fuzz_dep_.CoverTab[7835]++
//line /usr/local/go/src/encoding/asn1/asn1.go:820
		// _ = "end of CoverTab[7835]"
//line /usr/local/go/src/encoding/asn1/asn1.go:820
	}
//line /usr/local/go/src/encoding/asn1/asn1.go:820
	// _ = "end of CoverTab[7745]"
//line /usr/local/go/src/encoding/asn1/asn1.go:820
	_go_fuzz_dep_.CoverTab[7746]++

							if !params.explicit && func() bool {
//line /usr/local/go/src/encoding/asn1/asn1.go:822
		_go_fuzz_dep_.CoverTab[7836]++
//line /usr/local/go/src/encoding/asn1/asn1.go:822
		return params.application
//line /usr/local/go/src/encoding/asn1/asn1.go:822
		// _ = "end of CoverTab[7836]"
//line /usr/local/go/src/encoding/asn1/asn1.go:822
	}() && func() bool {
//line /usr/local/go/src/encoding/asn1/asn1.go:822
		_go_fuzz_dep_.CoverTab[7837]++
//line /usr/local/go/src/encoding/asn1/asn1.go:822
		return params.tag != nil
//line /usr/local/go/src/encoding/asn1/asn1.go:822
		// _ = "end of CoverTab[7837]"
//line /usr/local/go/src/encoding/asn1/asn1.go:822
	}() {
//line /usr/local/go/src/encoding/asn1/asn1.go:822
		_go_fuzz_dep_.CoverTab[7838]++
								expectedClass = ClassApplication
								expectedTag = *params.tag
								matchAnyClassAndTag = false
//line /usr/local/go/src/encoding/asn1/asn1.go:825
		// _ = "end of CoverTab[7838]"
	} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:826
		_go_fuzz_dep_.CoverTab[7839]++
//line /usr/local/go/src/encoding/asn1/asn1.go:826
		// _ = "end of CoverTab[7839]"
//line /usr/local/go/src/encoding/asn1/asn1.go:826
	}
//line /usr/local/go/src/encoding/asn1/asn1.go:826
	// _ = "end of CoverTab[7746]"
//line /usr/local/go/src/encoding/asn1/asn1.go:826
	_go_fuzz_dep_.CoverTab[7747]++

							if !params.explicit && func() bool {
//line /usr/local/go/src/encoding/asn1/asn1.go:828
		_go_fuzz_dep_.CoverTab[7840]++
//line /usr/local/go/src/encoding/asn1/asn1.go:828
		return params.private
//line /usr/local/go/src/encoding/asn1/asn1.go:828
		// _ = "end of CoverTab[7840]"
//line /usr/local/go/src/encoding/asn1/asn1.go:828
	}() && func() bool {
//line /usr/local/go/src/encoding/asn1/asn1.go:828
		_go_fuzz_dep_.CoverTab[7841]++
//line /usr/local/go/src/encoding/asn1/asn1.go:828
		return params.tag != nil
//line /usr/local/go/src/encoding/asn1/asn1.go:828
		// _ = "end of CoverTab[7841]"
//line /usr/local/go/src/encoding/asn1/asn1.go:828
	}() {
//line /usr/local/go/src/encoding/asn1/asn1.go:828
		_go_fuzz_dep_.CoverTab[7842]++
								expectedClass = ClassPrivate
								expectedTag = *params.tag
								matchAnyClassAndTag = false
//line /usr/local/go/src/encoding/asn1/asn1.go:831
		// _ = "end of CoverTab[7842]"
	} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:832
		_go_fuzz_dep_.CoverTab[7843]++
//line /usr/local/go/src/encoding/asn1/asn1.go:832
		// _ = "end of CoverTab[7843]"
//line /usr/local/go/src/encoding/asn1/asn1.go:832
	}
//line /usr/local/go/src/encoding/asn1/asn1.go:832
	// _ = "end of CoverTab[7747]"
//line /usr/local/go/src/encoding/asn1/asn1.go:832
	_go_fuzz_dep_.CoverTab[7748]++

//line /usr/local/go/src/encoding/asn1/asn1.go:835
	if !matchAnyClassAndTag && func() bool {
//line /usr/local/go/src/encoding/asn1/asn1.go:835
		_go_fuzz_dep_.CoverTab[7844]++
//line /usr/local/go/src/encoding/asn1/asn1.go:835
		return (t.class != expectedClass || func() bool {
//line /usr/local/go/src/encoding/asn1/asn1.go:835
			_go_fuzz_dep_.CoverTab[7845]++
//line /usr/local/go/src/encoding/asn1/asn1.go:835
			return t.tag != expectedTag
//line /usr/local/go/src/encoding/asn1/asn1.go:835
			// _ = "end of CoverTab[7845]"
//line /usr/local/go/src/encoding/asn1/asn1.go:835
		}())
//line /usr/local/go/src/encoding/asn1/asn1.go:835
		// _ = "end of CoverTab[7844]"
//line /usr/local/go/src/encoding/asn1/asn1.go:835
	}() || func() bool {
//line /usr/local/go/src/encoding/asn1/asn1.go:835
		_go_fuzz_dep_.CoverTab[7846]++
//line /usr/local/go/src/encoding/asn1/asn1.go:835
		return (!matchAny && func() bool {
									_go_fuzz_dep_.CoverTab[7847]++
//line /usr/local/go/src/encoding/asn1/asn1.go:836
			return t.isCompound != compoundType
//line /usr/local/go/src/encoding/asn1/asn1.go:836
			// _ = "end of CoverTab[7847]"
//line /usr/local/go/src/encoding/asn1/asn1.go:836
		}())
//line /usr/local/go/src/encoding/asn1/asn1.go:836
		// _ = "end of CoverTab[7846]"
//line /usr/local/go/src/encoding/asn1/asn1.go:836
	}() {
//line /usr/local/go/src/encoding/asn1/asn1.go:836
		_go_fuzz_dep_.CoverTab[7848]++

								ok := setDefaultValue(v, params)
								if ok {
//line /usr/local/go/src/encoding/asn1/asn1.go:839
			_go_fuzz_dep_.CoverTab[7850]++
									offset = initOffset
//line /usr/local/go/src/encoding/asn1/asn1.go:840
			// _ = "end of CoverTab[7850]"
		} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:841
			_go_fuzz_dep_.CoverTab[7851]++
									err = StructuralError{fmt.Sprintf("tags don't match (%d vs %+v) %+v %s @%d", expectedTag, t, params, fieldType.Name(), offset)}
//line /usr/local/go/src/encoding/asn1/asn1.go:842
			// _ = "end of CoverTab[7851]"
		}
//line /usr/local/go/src/encoding/asn1/asn1.go:843
		// _ = "end of CoverTab[7848]"
//line /usr/local/go/src/encoding/asn1/asn1.go:843
		_go_fuzz_dep_.CoverTab[7849]++
								return
//line /usr/local/go/src/encoding/asn1/asn1.go:844
		// _ = "end of CoverTab[7849]"
	} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:845
		_go_fuzz_dep_.CoverTab[7852]++
//line /usr/local/go/src/encoding/asn1/asn1.go:845
		// _ = "end of CoverTab[7852]"
//line /usr/local/go/src/encoding/asn1/asn1.go:845
	}
//line /usr/local/go/src/encoding/asn1/asn1.go:845
	// _ = "end of CoverTab[7748]"
//line /usr/local/go/src/encoding/asn1/asn1.go:845
	_go_fuzz_dep_.CoverTab[7749]++
							if invalidLength(offset, t.length, len(bytes)) {
//line /usr/local/go/src/encoding/asn1/asn1.go:846
		_go_fuzz_dep_.CoverTab[7853]++
								err = SyntaxError{"data truncated"}
								return
//line /usr/local/go/src/encoding/asn1/asn1.go:848
		// _ = "end of CoverTab[7853]"
	} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:849
		_go_fuzz_dep_.CoverTab[7854]++
//line /usr/local/go/src/encoding/asn1/asn1.go:849
		// _ = "end of CoverTab[7854]"
//line /usr/local/go/src/encoding/asn1/asn1.go:849
	}
//line /usr/local/go/src/encoding/asn1/asn1.go:849
	// _ = "end of CoverTab[7749]"
//line /usr/local/go/src/encoding/asn1/asn1.go:849
	_go_fuzz_dep_.CoverTab[7750]++
							innerBytes := bytes[offset : offset+t.length]
							offset += t.length

//line /usr/local/go/src/encoding/asn1/asn1.go:854
	switch v := v.Addr().Interface().(type) {
	case *RawValue:
//line /usr/local/go/src/encoding/asn1/asn1.go:855
		_go_fuzz_dep_.CoverTab[7855]++
								*v = RawValue{t.class, t.tag, t.isCompound, innerBytes, bytes[initOffset:offset]}
								return
//line /usr/local/go/src/encoding/asn1/asn1.go:857
		// _ = "end of CoverTab[7855]"
	case *ObjectIdentifier:
//line /usr/local/go/src/encoding/asn1/asn1.go:858
		_go_fuzz_dep_.CoverTab[7856]++
								*v, err = parseObjectIdentifier(innerBytes)
								return
//line /usr/local/go/src/encoding/asn1/asn1.go:860
		// _ = "end of CoverTab[7856]"
	case *BitString:
//line /usr/local/go/src/encoding/asn1/asn1.go:861
		_go_fuzz_dep_.CoverTab[7857]++
								*v, err = parseBitString(innerBytes)
								return
//line /usr/local/go/src/encoding/asn1/asn1.go:863
		// _ = "end of CoverTab[7857]"
	case *time.Time:
//line /usr/local/go/src/encoding/asn1/asn1.go:864
		_go_fuzz_dep_.CoverTab[7858]++
								if universalTag == TagUTCTime {
//line /usr/local/go/src/encoding/asn1/asn1.go:865
			_go_fuzz_dep_.CoverTab[7865]++
									*v, err = parseUTCTime(innerBytes)
									return
//line /usr/local/go/src/encoding/asn1/asn1.go:867
			// _ = "end of CoverTab[7865]"
		} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:868
			_go_fuzz_dep_.CoverTab[7866]++
//line /usr/local/go/src/encoding/asn1/asn1.go:868
			// _ = "end of CoverTab[7866]"
//line /usr/local/go/src/encoding/asn1/asn1.go:868
		}
//line /usr/local/go/src/encoding/asn1/asn1.go:868
		// _ = "end of CoverTab[7858]"
//line /usr/local/go/src/encoding/asn1/asn1.go:868
		_go_fuzz_dep_.CoverTab[7859]++
								*v, err = parseGeneralizedTime(innerBytes)
								return
//line /usr/local/go/src/encoding/asn1/asn1.go:870
		// _ = "end of CoverTab[7859]"
	case *Enumerated:
//line /usr/local/go/src/encoding/asn1/asn1.go:871
		_go_fuzz_dep_.CoverTab[7860]++
								parsedInt, err1 := parseInt32(innerBytes)
								if err1 == nil {
//line /usr/local/go/src/encoding/asn1/asn1.go:873
			_go_fuzz_dep_.CoverTab[7867]++
									*v = Enumerated(parsedInt)
//line /usr/local/go/src/encoding/asn1/asn1.go:874
			// _ = "end of CoverTab[7867]"
		} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:875
			_go_fuzz_dep_.CoverTab[7868]++
//line /usr/local/go/src/encoding/asn1/asn1.go:875
			// _ = "end of CoverTab[7868]"
//line /usr/local/go/src/encoding/asn1/asn1.go:875
		}
//line /usr/local/go/src/encoding/asn1/asn1.go:875
		// _ = "end of CoverTab[7860]"
//line /usr/local/go/src/encoding/asn1/asn1.go:875
		_go_fuzz_dep_.CoverTab[7861]++
								err = err1
								return
//line /usr/local/go/src/encoding/asn1/asn1.go:877
		// _ = "end of CoverTab[7861]"
	case *Flag:
//line /usr/local/go/src/encoding/asn1/asn1.go:878
		_go_fuzz_dep_.CoverTab[7862]++
								*v = true
								return
//line /usr/local/go/src/encoding/asn1/asn1.go:880
		// _ = "end of CoverTab[7862]"
	case **big.Int:
//line /usr/local/go/src/encoding/asn1/asn1.go:881
		_go_fuzz_dep_.CoverTab[7863]++
								parsedInt, err1 := parseBigInt(innerBytes)
								if err1 == nil {
//line /usr/local/go/src/encoding/asn1/asn1.go:883
			_go_fuzz_dep_.CoverTab[7869]++
									*v = parsedInt
//line /usr/local/go/src/encoding/asn1/asn1.go:884
			// _ = "end of CoverTab[7869]"
		} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:885
			_go_fuzz_dep_.CoverTab[7870]++
//line /usr/local/go/src/encoding/asn1/asn1.go:885
			// _ = "end of CoverTab[7870]"
//line /usr/local/go/src/encoding/asn1/asn1.go:885
		}
//line /usr/local/go/src/encoding/asn1/asn1.go:885
		// _ = "end of CoverTab[7863]"
//line /usr/local/go/src/encoding/asn1/asn1.go:885
		_go_fuzz_dep_.CoverTab[7864]++
								err = err1
								return
//line /usr/local/go/src/encoding/asn1/asn1.go:887
		// _ = "end of CoverTab[7864]"
	}
//line /usr/local/go/src/encoding/asn1/asn1.go:888
	// _ = "end of CoverTab[7750]"
//line /usr/local/go/src/encoding/asn1/asn1.go:888
	_go_fuzz_dep_.CoverTab[7751]++
							switch val := v; val.Kind() {
	case reflect.Bool:
//line /usr/local/go/src/encoding/asn1/asn1.go:890
		_go_fuzz_dep_.CoverTab[7871]++
								parsedBool, err1 := parseBool(innerBytes)
								if err1 == nil {
//line /usr/local/go/src/encoding/asn1/asn1.go:892
			_go_fuzz_dep_.CoverTab[7886]++
									val.SetBool(parsedBool)
//line /usr/local/go/src/encoding/asn1/asn1.go:893
			// _ = "end of CoverTab[7886]"
		} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:894
			_go_fuzz_dep_.CoverTab[7887]++
//line /usr/local/go/src/encoding/asn1/asn1.go:894
			// _ = "end of CoverTab[7887]"
//line /usr/local/go/src/encoding/asn1/asn1.go:894
		}
//line /usr/local/go/src/encoding/asn1/asn1.go:894
		// _ = "end of CoverTab[7871]"
//line /usr/local/go/src/encoding/asn1/asn1.go:894
		_go_fuzz_dep_.CoverTab[7872]++
								err = err1
								return
//line /usr/local/go/src/encoding/asn1/asn1.go:896
		// _ = "end of CoverTab[7872]"
	case reflect.Int, reflect.Int32, reflect.Int64:
//line /usr/local/go/src/encoding/asn1/asn1.go:897
		_go_fuzz_dep_.CoverTab[7873]++
								if val.Type().Size() == 4 {
//line /usr/local/go/src/encoding/asn1/asn1.go:898
			_go_fuzz_dep_.CoverTab[7888]++
									parsedInt, err1 := parseInt32(innerBytes)
									if err1 == nil {
//line /usr/local/go/src/encoding/asn1/asn1.go:900
				_go_fuzz_dep_.CoverTab[7890]++
										val.SetInt(int64(parsedInt))
//line /usr/local/go/src/encoding/asn1/asn1.go:901
				// _ = "end of CoverTab[7890]"
			} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:902
				_go_fuzz_dep_.CoverTab[7891]++
//line /usr/local/go/src/encoding/asn1/asn1.go:902
				// _ = "end of CoverTab[7891]"
//line /usr/local/go/src/encoding/asn1/asn1.go:902
			}
//line /usr/local/go/src/encoding/asn1/asn1.go:902
			// _ = "end of CoverTab[7888]"
//line /usr/local/go/src/encoding/asn1/asn1.go:902
			_go_fuzz_dep_.CoverTab[7889]++
									err = err1
//line /usr/local/go/src/encoding/asn1/asn1.go:903
			// _ = "end of CoverTab[7889]"
		} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:904
			_go_fuzz_dep_.CoverTab[7892]++
									parsedInt, err1 := parseInt64(innerBytes)
									if err1 == nil {
//line /usr/local/go/src/encoding/asn1/asn1.go:906
				_go_fuzz_dep_.CoverTab[7894]++
										val.SetInt(parsedInt)
//line /usr/local/go/src/encoding/asn1/asn1.go:907
				// _ = "end of CoverTab[7894]"
			} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:908
				_go_fuzz_dep_.CoverTab[7895]++
//line /usr/local/go/src/encoding/asn1/asn1.go:908
				// _ = "end of CoverTab[7895]"
//line /usr/local/go/src/encoding/asn1/asn1.go:908
			}
//line /usr/local/go/src/encoding/asn1/asn1.go:908
			// _ = "end of CoverTab[7892]"
//line /usr/local/go/src/encoding/asn1/asn1.go:908
			_go_fuzz_dep_.CoverTab[7893]++
									err = err1
//line /usr/local/go/src/encoding/asn1/asn1.go:909
			// _ = "end of CoverTab[7893]"
		}
//line /usr/local/go/src/encoding/asn1/asn1.go:910
		// _ = "end of CoverTab[7873]"
//line /usr/local/go/src/encoding/asn1/asn1.go:910
		_go_fuzz_dep_.CoverTab[7874]++
								return
//line /usr/local/go/src/encoding/asn1/asn1.go:911
		// _ = "end of CoverTab[7874]"

	case reflect.Struct:
//line /usr/local/go/src/encoding/asn1/asn1.go:913
		_go_fuzz_dep_.CoverTab[7875]++
								structType := fieldType

								for i := 0; i < structType.NumField(); i++ {
//line /usr/local/go/src/encoding/asn1/asn1.go:916
			_go_fuzz_dep_.CoverTab[7896]++
									if !structType.Field(i).IsExported() {
//line /usr/local/go/src/encoding/asn1/asn1.go:917
				_go_fuzz_dep_.CoverTab[7897]++
										err = StructuralError{"struct contains unexported fields"}
										return
//line /usr/local/go/src/encoding/asn1/asn1.go:919
				// _ = "end of CoverTab[7897]"
			} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:920
				_go_fuzz_dep_.CoverTab[7898]++
//line /usr/local/go/src/encoding/asn1/asn1.go:920
				// _ = "end of CoverTab[7898]"
//line /usr/local/go/src/encoding/asn1/asn1.go:920
			}
//line /usr/local/go/src/encoding/asn1/asn1.go:920
			// _ = "end of CoverTab[7896]"
		}
//line /usr/local/go/src/encoding/asn1/asn1.go:921
		// _ = "end of CoverTab[7875]"
//line /usr/local/go/src/encoding/asn1/asn1.go:921
		_go_fuzz_dep_.CoverTab[7876]++

								if structType.NumField() > 0 && func() bool {
//line /usr/local/go/src/encoding/asn1/asn1.go:923
			_go_fuzz_dep_.CoverTab[7899]++
//line /usr/local/go/src/encoding/asn1/asn1.go:923
			return structType.Field(0).Type == rawContentsType
									// _ = "end of CoverTab[7899]"
//line /usr/local/go/src/encoding/asn1/asn1.go:924
		}() {
//line /usr/local/go/src/encoding/asn1/asn1.go:924
			_go_fuzz_dep_.CoverTab[7900]++
									bytes := bytes[initOffset:offset]
									val.Field(0).Set(reflect.ValueOf(RawContent(bytes)))
//line /usr/local/go/src/encoding/asn1/asn1.go:926
			// _ = "end of CoverTab[7900]"
		} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:927
			_go_fuzz_dep_.CoverTab[7901]++
//line /usr/local/go/src/encoding/asn1/asn1.go:927
			// _ = "end of CoverTab[7901]"
//line /usr/local/go/src/encoding/asn1/asn1.go:927
		}
//line /usr/local/go/src/encoding/asn1/asn1.go:927
		// _ = "end of CoverTab[7876]"
//line /usr/local/go/src/encoding/asn1/asn1.go:927
		_go_fuzz_dep_.CoverTab[7877]++

								innerOffset := 0
								for i := 0; i < structType.NumField(); i++ {
//line /usr/local/go/src/encoding/asn1/asn1.go:930
			_go_fuzz_dep_.CoverTab[7902]++
									field := structType.Field(i)
									if i == 0 && func() bool {
//line /usr/local/go/src/encoding/asn1/asn1.go:932
				_go_fuzz_dep_.CoverTab[7904]++
//line /usr/local/go/src/encoding/asn1/asn1.go:932
				return field.Type == rawContentsType
//line /usr/local/go/src/encoding/asn1/asn1.go:932
				// _ = "end of CoverTab[7904]"
//line /usr/local/go/src/encoding/asn1/asn1.go:932
			}() {
//line /usr/local/go/src/encoding/asn1/asn1.go:932
				_go_fuzz_dep_.CoverTab[7905]++
										continue
//line /usr/local/go/src/encoding/asn1/asn1.go:933
				// _ = "end of CoverTab[7905]"
			} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:934
				_go_fuzz_dep_.CoverTab[7906]++
//line /usr/local/go/src/encoding/asn1/asn1.go:934
				// _ = "end of CoverTab[7906]"
//line /usr/local/go/src/encoding/asn1/asn1.go:934
			}
//line /usr/local/go/src/encoding/asn1/asn1.go:934
			// _ = "end of CoverTab[7902]"
//line /usr/local/go/src/encoding/asn1/asn1.go:934
			_go_fuzz_dep_.CoverTab[7903]++
									innerOffset, err = parseField(val.Field(i), innerBytes, innerOffset, parseFieldParameters(field.Tag.Get("asn1")))
									if err != nil {
//line /usr/local/go/src/encoding/asn1/asn1.go:936
				_go_fuzz_dep_.CoverTab[7907]++
										return
//line /usr/local/go/src/encoding/asn1/asn1.go:937
				// _ = "end of CoverTab[7907]"
			} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:938
				_go_fuzz_dep_.CoverTab[7908]++
//line /usr/local/go/src/encoding/asn1/asn1.go:938
				// _ = "end of CoverTab[7908]"
//line /usr/local/go/src/encoding/asn1/asn1.go:938
			}
//line /usr/local/go/src/encoding/asn1/asn1.go:938
			// _ = "end of CoverTab[7903]"
		}
//line /usr/local/go/src/encoding/asn1/asn1.go:939
		// _ = "end of CoverTab[7877]"
//line /usr/local/go/src/encoding/asn1/asn1.go:939
		_go_fuzz_dep_.CoverTab[7878]++

//line /usr/local/go/src/encoding/asn1/asn1.go:943
		return
//line /usr/local/go/src/encoding/asn1/asn1.go:943
		// _ = "end of CoverTab[7878]"
	case reflect.Slice:
//line /usr/local/go/src/encoding/asn1/asn1.go:944
		_go_fuzz_dep_.CoverTab[7879]++
								sliceType := fieldType
								if sliceType.Elem().Kind() == reflect.Uint8 {
//line /usr/local/go/src/encoding/asn1/asn1.go:946
			_go_fuzz_dep_.CoverTab[7909]++
									val.Set(reflect.MakeSlice(sliceType, len(innerBytes), len(innerBytes)))
									reflect.Copy(val, reflect.ValueOf(innerBytes))
									return
//line /usr/local/go/src/encoding/asn1/asn1.go:949
			// _ = "end of CoverTab[7909]"
		} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:950
			_go_fuzz_dep_.CoverTab[7910]++
//line /usr/local/go/src/encoding/asn1/asn1.go:950
			// _ = "end of CoverTab[7910]"
//line /usr/local/go/src/encoding/asn1/asn1.go:950
		}
//line /usr/local/go/src/encoding/asn1/asn1.go:950
		// _ = "end of CoverTab[7879]"
//line /usr/local/go/src/encoding/asn1/asn1.go:950
		_go_fuzz_dep_.CoverTab[7880]++
								newSlice, err1 := parseSequenceOf(innerBytes, sliceType, sliceType.Elem())
								if err1 == nil {
//line /usr/local/go/src/encoding/asn1/asn1.go:952
			_go_fuzz_dep_.CoverTab[7911]++
									val.Set(newSlice)
//line /usr/local/go/src/encoding/asn1/asn1.go:953
			// _ = "end of CoverTab[7911]"
		} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:954
			_go_fuzz_dep_.CoverTab[7912]++
//line /usr/local/go/src/encoding/asn1/asn1.go:954
			// _ = "end of CoverTab[7912]"
//line /usr/local/go/src/encoding/asn1/asn1.go:954
		}
//line /usr/local/go/src/encoding/asn1/asn1.go:954
		// _ = "end of CoverTab[7880]"
//line /usr/local/go/src/encoding/asn1/asn1.go:954
		_go_fuzz_dep_.CoverTab[7881]++
								err = err1
								return
//line /usr/local/go/src/encoding/asn1/asn1.go:956
		// _ = "end of CoverTab[7881]"
	case reflect.String:
//line /usr/local/go/src/encoding/asn1/asn1.go:957
		_go_fuzz_dep_.CoverTab[7882]++
								var v string
								switch universalTag {
		case TagPrintableString:
//line /usr/local/go/src/encoding/asn1/asn1.go:960
			_go_fuzz_dep_.CoverTab[7913]++
									v, err = parsePrintableString(innerBytes)
//line /usr/local/go/src/encoding/asn1/asn1.go:961
			// _ = "end of CoverTab[7913]"
		case TagNumericString:
//line /usr/local/go/src/encoding/asn1/asn1.go:962
			_go_fuzz_dep_.CoverTab[7914]++
									v, err = parseNumericString(innerBytes)
//line /usr/local/go/src/encoding/asn1/asn1.go:963
			// _ = "end of CoverTab[7914]"
		case TagIA5String:
//line /usr/local/go/src/encoding/asn1/asn1.go:964
			_go_fuzz_dep_.CoverTab[7915]++
									v, err = parseIA5String(innerBytes)
//line /usr/local/go/src/encoding/asn1/asn1.go:965
			// _ = "end of CoverTab[7915]"
		case TagT61String:
//line /usr/local/go/src/encoding/asn1/asn1.go:966
			_go_fuzz_dep_.CoverTab[7916]++
									v, err = parseT61String(innerBytes)
//line /usr/local/go/src/encoding/asn1/asn1.go:967
			// _ = "end of CoverTab[7916]"
		case TagUTF8String:
//line /usr/local/go/src/encoding/asn1/asn1.go:968
			_go_fuzz_dep_.CoverTab[7917]++
									v, err = parseUTF8String(innerBytes)
//line /usr/local/go/src/encoding/asn1/asn1.go:969
			// _ = "end of CoverTab[7917]"
		case TagGeneralString:
//line /usr/local/go/src/encoding/asn1/asn1.go:970
			_go_fuzz_dep_.CoverTab[7918]++

//line /usr/local/go/src/encoding/asn1/asn1.go:975
			v, err = parseT61String(innerBytes)
//line /usr/local/go/src/encoding/asn1/asn1.go:975
			// _ = "end of CoverTab[7918]"
		case TagBMPString:
//line /usr/local/go/src/encoding/asn1/asn1.go:976
			_go_fuzz_dep_.CoverTab[7919]++
									v, err = parseBMPString(innerBytes)
//line /usr/local/go/src/encoding/asn1/asn1.go:977
			// _ = "end of CoverTab[7919]"

		default:
//line /usr/local/go/src/encoding/asn1/asn1.go:979
			_go_fuzz_dep_.CoverTab[7920]++
									err = SyntaxError{fmt.Sprintf("internal error: unknown string type %d", universalTag)}
//line /usr/local/go/src/encoding/asn1/asn1.go:980
			// _ = "end of CoverTab[7920]"
		}
//line /usr/local/go/src/encoding/asn1/asn1.go:981
		// _ = "end of CoverTab[7882]"
//line /usr/local/go/src/encoding/asn1/asn1.go:981
		_go_fuzz_dep_.CoverTab[7883]++
								if err == nil {
//line /usr/local/go/src/encoding/asn1/asn1.go:982
			_go_fuzz_dep_.CoverTab[7921]++
									val.SetString(v)
//line /usr/local/go/src/encoding/asn1/asn1.go:983
			// _ = "end of CoverTab[7921]"
		} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:984
			_go_fuzz_dep_.CoverTab[7922]++
//line /usr/local/go/src/encoding/asn1/asn1.go:984
			// _ = "end of CoverTab[7922]"
//line /usr/local/go/src/encoding/asn1/asn1.go:984
		}
//line /usr/local/go/src/encoding/asn1/asn1.go:984
		// _ = "end of CoverTab[7883]"
//line /usr/local/go/src/encoding/asn1/asn1.go:984
		_go_fuzz_dep_.CoverTab[7884]++
								return
//line /usr/local/go/src/encoding/asn1/asn1.go:985
		// _ = "end of CoverTab[7884]"
//line /usr/local/go/src/encoding/asn1/asn1.go:985
	default:
//line /usr/local/go/src/encoding/asn1/asn1.go:985
		_go_fuzz_dep_.CoverTab[7885]++
//line /usr/local/go/src/encoding/asn1/asn1.go:985
		// _ = "end of CoverTab[7885]"
	}
//line /usr/local/go/src/encoding/asn1/asn1.go:986
	// _ = "end of CoverTab[7751]"
//line /usr/local/go/src/encoding/asn1/asn1.go:986
	_go_fuzz_dep_.CoverTab[7752]++
							err = StructuralError{"unsupported: " + v.Type().String()}
							return
//line /usr/local/go/src/encoding/asn1/asn1.go:988
	// _ = "end of CoverTab[7752]"
}

// canHaveDefaultValue reports whether k is a Kind that we will set a default
//line /usr/local/go/src/encoding/asn1/asn1.go:991
// value for. (A signed integer, essentially.)
//line /usr/local/go/src/encoding/asn1/asn1.go:993
func canHaveDefaultValue(k reflect.Kind) bool {
//line /usr/local/go/src/encoding/asn1/asn1.go:993
	_go_fuzz_dep_.CoverTab[7923]++
							switch k {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
//line /usr/local/go/src/encoding/asn1/asn1.go:995
		_go_fuzz_dep_.CoverTab[7925]++
								return true
//line /usr/local/go/src/encoding/asn1/asn1.go:996
		// _ = "end of CoverTab[7925]"
//line /usr/local/go/src/encoding/asn1/asn1.go:996
	default:
//line /usr/local/go/src/encoding/asn1/asn1.go:996
		_go_fuzz_dep_.CoverTab[7926]++
//line /usr/local/go/src/encoding/asn1/asn1.go:996
		// _ = "end of CoverTab[7926]"
	}
//line /usr/local/go/src/encoding/asn1/asn1.go:997
	// _ = "end of CoverTab[7923]"
//line /usr/local/go/src/encoding/asn1/asn1.go:997
	_go_fuzz_dep_.CoverTab[7924]++

							return false
//line /usr/local/go/src/encoding/asn1/asn1.go:999
	// _ = "end of CoverTab[7924]"
}

// setDefaultValue is used to install a default value, from a tag string, into
//line /usr/local/go/src/encoding/asn1/asn1.go:1002
// a Value. It is successful if the field was optional, even if a default value
//line /usr/local/go/src/encoding/asn1/asn1.go:1002
// wasn't provided or it failed to install it into the Value.
//line /usr/local/go/src/encoding/asn1/asn1.go:1005
func setDefaultValue(v reflect.Value, params fieldParameters) (ok bool) {
//line /usr/local/go/src/encoding/asn1/asn1.go:1005
	_go_fuzz_dep_.CoverTab[7927]++
							if !params.optional {
//line /usr/local/go/src/encoding/asn1/asn1.go:1006
		_go_fuzz_dep_.CoverTab[7931]++
								return
//line /usr/local/go/src/encoding/asn1/asn1.go:1007
		// _ = "end of CoverTab[7931]"
	} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:1008
		_go_fuzz_dep_.CoverTab[7932]++
//line /usr/local/go/src/encoding/asn1/asn1.go:1008
		// _ = "end of CoverTab[7932]"
//line /usr/local/go/src/encoding/asn1/asn1.go:1008
	}
//line /usr/local/go/src/encoding/asn1/asn1.go:1008
	// _ = "end of CoverTab[7927]"
//line /usr/local/go/src/encoding/asn1/asn1.go:1008
	_go_fuzz_dep_.CoverTab[7928]++
							ok = true
							if params.defaultValue == nil {
//line /usr/local/go/src/encoding/asn1/asn1.go:1010
		_go_fuzz_dep_.CoverTab[7933]++
								return
//line /usr/local/go/src/encoding/asn1/asn1.go:1011
		// _ = "end of CoverTab[7933]"
	} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:1012
		_go_fuzz_dep_.CoverTab[7934]++
//line /usr/local/go/src/encoding/asn1/asn1.go:1012
		// _ = "end of CoverTab[7934]"
//line /usr/local/go/src/encoding/asn1/asn1.go:1012
	}
//line /usr/local/go/src/encoding/asn1/asn1.go:1012
	// _ = "end of CoverTab[7928]"
//line /usr/local/go/src/encoding/asn1/asn1.go:1012
	_go_fuzz_dep_.CoverTab[7929]++
							if canHaveDefaultValue(v.Kind()) {
//line /usr/local/go/src/encoding/asn1/asn1.go:1013
		_go_fuzz_dep_.CoverTab[7935]++
								v.SetInt(*params.defaultValue)
//line /usr/local/go/src/encoding/asn1/asn1.go:1014
		// _ = "end of CoverTab[7935]"
	} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:1015
		_go_fuzz_dep_.CoverTab[7936]++
//line /usr/local/go/src/encoding/asn1/asn1.go:1015
		// _ = "end of CoverTab[7936]"
//line /usr/local/go/src/encoding/asn1/asn1.go:1015
	}
//line /usr/local/go/src/encoding/asn1/asn1.go:1015
	// _ = "end of CoverTab[7929]"
//line /usr/local/go/src/encoding/asn1/asn1.go:1015
	_go_fuzz_dep_.CoverTab[7930]++
							return
//line /usr/local/go/src/encoding/asn1/asn1.go:1016
	// _ = "end of CoverTab[7930]"
}

// Unmarshal parses the DER-encoded ASN.1 data structure b
//line /usr/local/go/src/encoding/asn1/asn1.go:1019
// and uses the reflect package to fill in an arbitrary value pointed at by val.
//line /usr/local/go/src/encoding/asn1/asn1.go:1019
// Because Unmarshal uses the reflect package, the structs
//line /usr/local/go/src/encoding/asn1/asn1.go:1019
// being written to must use upper case field names. If val
//line /usr/local/go/src/encoding/asn1/asn1.go:1019
// is nil or not a pointer, Unmarshal returns an error.
//line /usr/local/go/src/encoding/asn1/asn1.go:1019
//
//line /usr/local/go/src/encoding/asn1/asn1.go:1019
// After parsing b, any bytes that were leftover and not used to fill
//line /usr/local/go/src/encoding/asn1/asn1.go:1019
// val will be returned in rest. When parsing a SEQUENCE into a struct,
//line /usr/local/go/src/encoding/asn1/asn1.go:1019
// any trailing elements of the SEQUENCE that do not have matching
//line /usr/local/go/src/encoding/asn1/asn1.go:1019
// fields in val will not be included in rest, as these are considered
//line /usr/local/go/src/encoding/asn1/asn1.go:1019
// valid elements of the SEQUENCE and not trailing data.
//line /usr/local/go/src/encoding/asn1/asn1.go:1019
//
//line /usr/local/go/src/encoding/asn1/asn1.go:1019
// An ASN.1 INTEGER can be written to an int, int32, int64,
//line /usr/local/go/src/encoding/asn1/asn1.go:1019
// or *big.Int (from the math/big package).
//line /usr/local/go/src/encoding/asn1/asn1.go:1019
// If the encoded value does not fit in the Go type,
//line /usr/local/go/src/encoding/asn1/asn1.go:1019
// Unmarshal returns a parse error.
//line /usr/local/go/src/encoding/asn1/asn1.go:1019
//
//line /usr/local/go/src/encoding/asn1/asn1.go:1019
// An ASN.1 BIT STRING can be written to a BitString.
//line /usr/local/go/src/encoding/asn1/asn1.go:1019
//
//line /usr/local/go/src/encoding/asn1/asn1.go:1019
// An ASN.1 OCTET STRING can be written to a []byte.
//line /usr/local/go/src/encoding/asn1/asn1.go:1019
//
//line /usr/local/go/src/encoding/asn1/asn1.go:1019
// An ASN.1 OBJECT IDENTIFIER can be written to an
//line /usr/local/go/src/encoding/asn1/asn1.go:1019
// ObjectIdentifier.
//line /usr/local/go/src/encoding/asn1/asn1.go:1019
//
//line /usr/local/go/src/encoding/asn1/asn1.go:1019
// An ASN.1 ENUMERATED can be written to an Enumerated.
//line /usr/local/go/src/encoding/asn1/asn1.go:1019
//
//line /usr/local/go/src/encoding/asn1/asn1.go:1019
// An ASN.1 UTCTIME or GENERALIZEDTIME can be written to a time.Time.
//line /usr/local/go/src/encoding/asn1/asn1.go:1019
//
//line /usr/local/go/src/encoding/asn1/asn1.go:1019
// An ASN.1 PrintableString, IA5String, or NumericString can be written to a string.
//line /usr/local/go/src/encoding/asn1/asn1.go:1019
//
//line /usr/local/go/src/encoding/asn1/asn1.go:1019
// Any of the above ASN.1 values can be written to an interface{}.
//line /usr/local/go/src/encoding/asn1/asn1.go:1019
// The value stored in the interface has the corresponding Go type.
//line /usr/local/go/src/encoding/asn1/asn1.go:1019
// For integers, that type is int64.
//line /usr/local/go/src/encoding/asn1/asn1.go:1019
//
//line /usr/local/go/src/encoding/asn1/asn1.go:1019
// An ASN.1 SEQUENCE OF x or SET OF x can be written
//line /usr/local/go/src/encoding/asn1/asn1.go:1019
// to a slice if an x can be written to the slice's element type.
//line /usr/local/go/src/encoding/asn1/asn1.go:1019
//
//line /usr/local/go/src/encoding/asn1/asn1.go:1019
// An ASN.1 SEQUENCE or SET can be written to a struct
//line /usr/local/go/src/encoding/asn1/asn1.go:1019
// if each of the elements in the sequence can be
//line /usr/local/go/src/encoding/asn1/asn1.go:1019
// written to the corresponding element in the struct.
//line /usr/local/go/src/encoding/asn1/asn1.go:1019
//
//line /usr/local/go/src/encoding/asn1/asn1.go:1019
// The following tags on struct fields have special meaning to Unmarshal:
//line /usr/local/go/src/encoding/asn1/asn1.go:1019
//
//line /usr/local/go/src/encoding/asn1/asn1.go:1019
//	application specifies that an APPLICATION tag is used
//line /usr/local/go/src/encoding/asn1/asn1.go:1019
//	private     specifies that a PRIVATE tag is used
//line /usr/local/go/src/encoding/asn1/asn1.go:1019
//	default:x   sets the default value for optional integer fields (only used if optional is also present)
//line /usr/local/go/src/encoding/asn1/asn1.go:1019
//	explicit    specifies that an additional, explicit tag wraps the implicit one
//line /usr/local/go/src/encoding/asn1/asn1.go:1019
//	optional    marks the field as ASN.1 OPTIONAL
//line /usr/local/go/src/encoding/asn1/asn1.go:1019
//	set         causes a SET, rather than a SEQUENCE type to be expected
//line /usr/local/go/src/encoding/asn1/asn1.go:1019
//	tag:x       specifies the ASN.1 tag number; implies ASN.1 CONTEXT SPECIFIC
//line /usr/local/go/src/encoding/asn1/asn1.go:1019
//
//line /usr/local/go/src/encoding/asn1/asn1.go:1019
// When decoding an ASN.1 value with an IMPLICIT tag into a string field,
//line /usr/local/go/src/encoding/asn1/asn1.go:1019
// Unmarshal will default to a PrintableString, which doesn't support
//line /usr/local/go/src/encoding/asn1/asn1.go:1019
// characters such as '@' and '&'. To force other encodings, use the following
//line /usr/local/go/src/encoding/asn1/asn1.go:1019
// tags:
//line /usr/local/go/src/encoding/asn1/asn1.go:1019
//
//line /usr/local/go/src/encoding/asn1/asn1.go:1019
//	ia5     causes strings to be unmarshaled as ASN.1 IA5String values
//line /usr/local/go/src/encoding/asn1/asn1.go:1019
//	numeric causes strings to be unmarshaled as ASN.1 NumericString values
//line /usr/local/go/src/encoding/asn1/asn1.go:1019
//	utf8    causes strings to be unmarshaled as ASN.1 UTF8String values
//line /usr/local/go/src/encoding/asn1/asn1.go:1019
//
//line /usr/local/go/src/encoding/asn1/asn1.go:1019
// If the type of the first field of a structure is RawContent then the raw
//line /usr/local/go/src/encoding/asn1/asn1.go:1019
// ASN1 contents of the struct will be stored in it.
//line /usr/local/go/src/encoding/asn1/asn1.go:1019
//
//line /usr/local/go/src/encoding/asn1/asn1.go:1019
// If the name of a slice type ends with "SET" then it's treated as if
//line /usr/local/go/src/encoding/asn1/asn1.go:1019
// the "set" tag was set on it. This results in interpreting the type as a
//line /usr/local/go/src/encoding/asn1/asn1.go:1019
// SET OF x rather than a SEQUENCE OF x. This can be used with nested slices
//line /usr/local/go/src/encoding/asn1/asn1.go:1019
// where a struct tag cannot be given.
//line /usr/local/go/src/encoding/asn1/asn1.go:1019
//
//line /usr/local/go/src/encoding/asn1/asn1.go:1019
// Other ASN.1 types are not supported; if it encounters them,
//line /usr/local/go/src/encoding/asn1/asn1.go:1019
// Unmarshal returns a parse error.
//line /usr/local/go/src/encoding/asn1/asn1.go:1089
func Unmarshal(b []byte, val any) (rest []byte, err error) {
//line /usr/local/go/src/encoding/asn1/asn1.go:1089
	_go_fuzz_dep_.CoverTab[7937]++
							return UnmarshalWithParams(b, val, "")
//line /usr/local/go/src/encoding/asn1/asn1.go:1090
	// _ = "end of CoverTab[7937]"
}

// An invalidUnmarshalError describes an invalid argument passed to Unmarshal.
//line /usr/local/go/src/encoding/asn1/asn1.go:1093
// (The argument to Unmarshal must be a non-nil pointer.)
//line /usr/local/go/src/encoding/asn1/asn1.go:1095
type invalidUnmarshalError struct {
	Type reflect.Type
}

func (e *invalidUnmarshalError) Error() string {
//line /usr/local/go/src/encoding/asn1/asn1.go:1099
	_go_fuzz_dep_.CoverTab[7938]++
							if e.Type == nil {
//line /usr/local/go/src/encoding/asn1/asn1.go:1100
		_go_fuzz_dep_.CoverTab[7941]++
								return "asn1: Unmarshal recipient value is nil"
//line /usr/local/go/src/encoding/asn1/asn1.go:1101
		// _ = "end of CoverTab[7941]"
	} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:1102
		_go_fuzz_dep_.CoverTab[7942]++
//line /usr/local/go/src/encoding/asn1/asn1.go:1102
		// _ = "end of CoverTab[7942]"
//line /usr/local/go/src/encoding/asn1/asn1.go:1102
	}
//line /usr/local/go/src/encoding/asn1/asn1.go:1102
	// _ = "end of CoverTab[7938]"
//line /usr/local/go/src/encoding/asn1/asn1.go:1102
	_go_fuzz_dep_.CoverTab[7939]++

							if e.Type.Kind() != reflect.Pointer {
//line /usr/local/go/src/encoding/asn1/asn1.go:1104
		_go_fuzz_dep_.CoverTab[7943]++
								return "asn1: Unmarshal recipient value is non-pointer " + e.Type.String()
//line /usr/local/go/src/encoding/asn1/asn1.go:1105
		// _ = "end of CoverTab[7943]"
	} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:1106
		_go_fuzz_dep_.CoverTab[7944]++
//line /usr/local/go/src/encoding/asn1/asn1.go:1106
		// _ = "end of CoverTab[7944]"
//line /usr/local/go/src/encoding/asn1/asn1.go:1106
	}
//line /usr/local/go/src/encoding/asn1/asn1.go:1106
	// _ = "end of CoverTab[7939]"
//line /usr/local/go/src/encoding/asn1/asn1.go:1106
	_go_fuzz_dep_.CoverTab[7940]++
							return "asn1: Unmarshal recipient value is nil " + e.Type.String()
//line /usr/local/go/src/encoding/asn1/asn1.go:1107
	// _ = "end of CoverTab[7940]"
}

// UnmarshalWithParams allows field parameters to be specified for the
//line /usr/local/go/src/encoding/asn1/asn1.go:1110
// top-level element. The form of the params is the same as the field tags.
//line /usr/local/go/src/encoding/asn1/asn1.go:1112
func UnmarshalWithParams(b []byte, val any, params string) (rest []byte, err error) {
//line /usr/local/go/src/encoding/asn1/asn1.go:1112
	_go_fuzz_dep_.CoverTab[7945]++
							v := reflect.ValueOf(val)
							if v.Kind() != reflect.Pointer || func() bool {
//line /usr/local/go/src/encoding/asn1/asn1.go:1114
		_go_fuzz_dep_.CoverTab[7948]++
//line /usr/local/go/src/encoding/asn1/asn1.go:1114
		return v.IsNil()
//line /usr/local/go/src/encoding/asn1/asn1.go:1114
		// _ = "end of CoverTab[7948]"
//line /usr/local/go/src/encoding/asn1/asn1.go:1114
	}() {
//line /usr/local/go/src/encoding/asn1/asn1.go:1114
		_go_fuzz_dep_.CoverTab[7949]++
								return nil, &invalidUnmarshalError{reflect.TypeOf(val)}
//line /usr/local/go/src/encoding/asn1/asn1.go:1115
		// _ = "end of CoverTab[7949]"
	} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:1116
		_go_fuzz_dep_.CoverTab[7950]++
//line /usr/local/go/src/encoding/asn1/asn1.go:1116
		// _ = "end of CoverTab[7950]"
//line /usr/local/go/src/encoding/asn1/asn1.go:1116
	}
//line /usr/local/go/src/encoding/asn1/asn1.go:1116
	// _ = "end of CoverTab[7945]"
//line /usr/local/go/src/encoding/asn1/asn1.go:1116
	_go_fuzz_dep_.CoverTab[7946]++
							offset, err := parseField(v.Elem(), b, 0, parseFieldParameters(params))
							if err != nil {
//line /usr/local/go/src/encoding/asn1/asn1.go:1118
		_go_fuzz_dep_.CoverTab[7951]++
								return nil, err
//line /usr/local/go/src/encoding/asn1/asn1.go:1119
		// _ = "end of CoverTab[7951]"
	} else {
//line /usr/local/go/src/encoding/asn1/asn1.go:1120
		_go_fuzz_dep_.CoverTab[7952]++
//line /usr/local/go/src/encoding/asn1/asn1.go:1120
		// _ = "end of CoverTab[7952]"
//line /usr/local/go/src/encoding/asn1/asn1.go:1120
	}
//line /usr/local/go/src/encoding/asn1/asn1.go:1120
	// _ = "end of CoverTab[7946]"
//line /usr/local/go/src/encoding/asn1/asn1.go:1120
	_go_fuzz_dep_.CoverTab[7947]++
							return b[offset:], nil
//line /usr/local/go/src/encoding/asn1/asn1.go:1121
	// _ = "end of CoverTab[7947]"
}

//line /usr/local/go/src/encoding/asn1/asn1.go:1122
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/encoding/asn1/asn1.go:1122
var _ = _go_fuzz_dep_.CoverTab
