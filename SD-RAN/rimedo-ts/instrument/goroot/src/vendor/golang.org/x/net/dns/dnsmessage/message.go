// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:5
// Package dnsmessage provides a mostly RFC 1035 compliant implementation of
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:5
// DNS message packing and unpacking.
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:5
//
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:5
// The package also supports messages with Extension Mechanisms for DNS
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:5
// (EDNS(0)) as defined in RFC 6891.
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:5
//
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:5
// This implementation is designed to minimize heap allocations and avoid
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:5
// unnecessary packing and unpacking as much as possible.
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:13
package dnsmessage

//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:13
import (
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:13
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:13
)
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:13
import (
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:13
	_atomic_ "sync/atomic"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:13
)

import (
	"errors"
)

//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:21
// A Type is a type of DNS request and response.
type Type uint16

const (
	// ResourceHeader.Type and Question.Type
	TypeA		Type	= 1
	TypeNS		Type	= 2
	TypeCNAME	Type	= 5
	TypeSOA		Type	= 6
	TypePTR		Type	= 12
	TypeMX		Type	= 15
	TypeTXT		Type	= 16
	TypeAAAA	Type	= 28
	TypeSRV		Type	= 33
	TypeOPT		Type	= 41

	// Question.Type
	TypeWKS		Type	= 11
	TypeHINFO	Type	= 13
	TypeMINFO	Type	= 14
	TypeAXFR	Type	= 252
	TypeALL		Type	= 255
)

var typeNames = map[Type]string{
	TypeA:		"TypeA",
	TypeNS:		"TypeNS",
	TypeCNAME:	"TypeCNAME",
	TypeSOA:	"TypeSOA",
	TypePTR:	"TypePTR",
	TypeMX:		"TypeMX",
	TypeTXT:	"TypeTXT",
	TypeAAAA:	"TypeAAAA",
	TypeSRV:	"TypeSRV",
	TypeOPT:	"TypeOPT",
	TypeWKS:	"TypeWKS",
	TypeHINFO:	"TypeHINFO",
	TypeMINFO:	"TypeMINFO",
	TypeAXFR:	"TypeAXFR",
	TypeALL:	"TypeALL",
}

// String implements fmt.Stringer.String.
func (t Type) String() string {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:64
	_go_fuzz_dep_.CoverTab[10856]++
										if n, ok := typeNames[t]; ok {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:65
		_go_fuzz_dep_.CoverTab[10858]++
											return n
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:66
		// _ = "end of CoverTab[10858]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:67
		_go_fuzz_dep_.CoverTab[10859]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:67
		// _ = "end of CoverTab[10859]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:67
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:67
	// _ = "end of CoverTab[10856]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:67
	_go_fuzz_dep_.CoverTab[10857]++
										return printUint16(uint16(t))
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:68
	// _ = "end of CoverTab[10857]"
}

// GoString implements fmt.GoStringer.GoString.
func (t Type) GoString() string {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:72
	_go_fuzz_dep_.CoverTab[10860]++
										if n, ok := typeNames[t]; ok {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:73
		_go_fuzz_dep_.CoverTab[10862]++
											return "dnsmessage." + n
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:74
		// _ = "end of CoverTab[10862]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:75
		_go_fuzz_dep_.CoverTab[10863]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:75
		// _ = "end of CoverTab[10863]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:75
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:75
	// _ = "end of CoverTab[10860]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:75
	_go_fuzz_dep_.CoverTab[10861]++
										return printUint16(uint16(t))
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:76
	// _ = "end of CoverTab[10861]"
}

// A Class is a type of network.
type Class uint16

const (
	// ResourceHeader.Class and Question.Class
	ClassINET	Class	= 1
	ClassCSNET	Class	= 2
	ClassCHAOS	Class	= 3
	ClassHESIOD	Class	= 4

	// Question.Class
	ClassANY	Class	= 255
)

var classNames = map[Class]string{
	ClassINET:	"ClassINET",
	ClassCSNET:	"ClassCSNET",
	ClassCHAOS:	"ClassCHAOS",
	ClassHESIOD:	"ClassHESIOD",
	ClassANY:	"ClassANY",
}

// String implements fmt.Stringer.String.
func (c Class) String() string {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:102
	_go_fuzz_dep_.CoverTab[10864]++
										if n, ok := classNames[c]; ok {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:103
		_go_fuzz_dep_.CoverTab[10866]++
											return n
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:104
		// _ = "end of CoverTab[10866]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:105
		_go_fuzz_dep_.CoverTab[10867]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:105
		// _ = "end of CoverTab[10867]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:105
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:105
	// _ = "end of CoverTab[10864]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:105
	_go_fuzz_dep_.CoverTab[10865]++
										return printUint16(uint16(c))
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:106
	// _ = "end of CoverTab[10865]"
}

// GoString implements fmt.GoStringer.GoString.
func (c Class) GoString() string {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:110
	_go_fuzz_dep_.CoverTab[10868]++
										if n, ok := classNames[c]; ok {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:111
		_go_fuzz_dep_.CoverTab[10870]++
											return "dnsmessage." + n
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:112
		// _ = "end of CoverTab[10870]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:113
		_go_fuzz_dep_.CoverTab[10871]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:113
		// _ = "end of CoverTab[10871]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:113
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:113
	// _ = "end of CoverTab[10868]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:113
	_go_fuzz_dep_.CoverTab[10869]++
										return printUint16(uint16(c))
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:114
	// _ = "end of CoverTab[10869]"
}

// An OpCode is a DNS operation code.
type OpCode uint16

// GoString implements fmt.GoStringer.GoString.
func (o OpCode) GoString() string {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:121
	_go_fuzz_dep_.CoverTab[10872]++
										return printUint16(uint16(o))
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:122
	// _ = "end of CoverTab[10872]"
}

// An RCode is a DNS response status code.
type RCode uint16

// Header.RCode values.
const (
	RCodeSuccess		RCode	= 0	// NoError
	RCodeFormatError	RCode	= 1	// FormErr
	RCodeServerFailure	RCode	= 2	// ServFail
	RCodeNameError		RCode	= 3	// NXDomain
	RCodeNotImplemented	RCode	= 4	// NotImp
	RCodeRefused		RCode	= 5	// Refused
)

var rCodeNames = map[RCode]string{
	RCodeSuccess:		"RCodeSuccess",
	RCodeFormatError:	"RCodeFormatError",
	RCodeServerFailure:	"RCodeServerFailure",
	RCodeNameError:		"RCodeNameError",
	RCodeNotImplemented:	"RCodeNotImplemented",
	RCodeRefused:		"RCodeRefused",
}

// String implements fmt.Stringer.String.
func (r RCode) String() string {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:148
	_go_fuzz_dep_.CoverTab[10873]++
										if n, ok := rCodeNames[r]; ok {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:149
		_go_fuzz_dep_.CoverTab[10875]++
											return n
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:150
		// _ = "end of CoverTab[10875]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:151
		_go_fuzz_dep_.CoverTab[10876]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:151
		// _ = "end of CoverTab[10876]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:151
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:151
	// _ = "end of CoverTab[10873]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:151
	_go_fuzz_dep_.CoverTab[10874]++
										return printUint16(uint16(r))
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:152
	// _ = "end of CoverTab[10874]"
}

// GoString implements fmt.GoStringer.GoString.
func (r RCode) GoString() string {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:156
	_go_fuzz_dep_.CoverTab[10877]++
										if n, ok := rCodeNames[r]; ok {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:157
		_go_fuzz_dep_.CoverTab[10879]++
											return "dnsmessage." + n
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:158
		// _ = "end of CoverTab[10879]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:159
		_go_fuzz_dep_.CoverTab[10880]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:159
		// _ = "end of CoverTab[10880]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:159
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:159
	// _ = "end of CoverTab[10877]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:159
	_go_fuzz_dep_.CoverTab[10878]++
										return printUint16(uint16(r))
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:160
	// _ = "end of CoverTab[10878]"
}

func printPaddedUint8(i uint8) string {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:163
	_go_fuzz_dep_.CoverTab[10881]++
										b := byte(i)
										return string([]byte{
		b/100 + '0',
		b/10%10 + '0',
		b%10 + '0',
	})
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:169
	// _ = "end of CoverTab[10881]"
}

func printUint8Bytes(buf []byte, i uint8) []byte {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:172
	_go_fuzz_dep_.CoverTab[10882]++
										b := byte(i)
										if i >= 100 {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:174
		_go_fuzz_dep_.CoverTab[10885]++
											buf = append(buf, b/100+'0')
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:175
		// _ = "end of CoverTab[10885]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:176
		_go_fuzz_dep_.CoverTab[10886]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:176
		// _ = "end of CoverTab[10886]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:176
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:176
	// _ = "end of CoverTab[10882]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:176
	_go_fuzz_dep_.CoverTab[10883]++
										if i >= 10 {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:177
		_go_fuzz_dep_.CoverTab[10887]++
											buf = append(buf, b/10%10+'0')
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:178
		// _ = "end of CoverTab[10887]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:179
		_go_fuzz_dep_.CoverTab[10888]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:179
		// _ = "end of CoverTab[10888]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:179
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:179
	// _ = "end of CoverTab[10883]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:179
	_go_fuzz_dep_.CoverTab[10884]++
										return append(buf, b%10+'0')
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:180
	// _ = "end of CoverTab[10884]"
}

func printByteSlice(b []byte) string {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:183
	_go_fuzz_dep_.CoverTab[10889]++
										if len(b) == 0 {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:184
		_go_fuzz_dep_.CoverTab[10892]++
											return ""
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:185
		// _ = "end of CoverTab[10892]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:186
		_go_fuzz_dep_.CoverTab[10893]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:186
		// _ = "end of CoverTab[10893]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:186
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:186
	// _ = "end of CoverTab[10889]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:186
	_go_fuzz_dep_.CoverTab[10890]++
										buf := make([]byte, 0, 5*len(b))
										buf = printUint8Bytes(buf, uint8(b[0]))
										for _, n := range b[1:] {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:189
		_go_fuzz_dep_.CoverTab[10894]++
											buf = append(buf, ',', ' ')
											buf = printUint8Bytes(buf, uint8(n))
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:191
		// _ = "end of CoverTab[10894]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:192
	// _ = "end of CoverTab[10890]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:192
	_go_fuzz_dep_.CoverTab[10891]++
										return string(buf)
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:193
	// _ = "end of CoverTab[10891]"
}

const hexDigits = "0123456789abcdef"

func printString(str []byte) string {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:198
	_go_fuzz_dep_.CoverTab[10895]++
										buf := make([]byte, 0, len(str))
										for i := 0; i < len(str); i++ {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:200
		_go_fuzz_dep_.CoverTab[10897]++
											c := str[i]
											if c == '.' || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:202
			_go_fuzz_dep_.CoverTab[10899]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:202
			return c == '-'
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:202
			// _ = "end of CoverTab[10899]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:202
		}() || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:202
			_go_fuzz_dep_.CoverTab[10900]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:202
			return c == ' '
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:202
			// _ = "end of CoverTab[10900]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:202
		}() || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:202
			_go_fuzz_dep_.CoverTab[10901]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:202
			return 'A' <= c && func() bool {
													_go_fuzz_dep_.CoverTab[10902]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:203
				return c <= 'Z'
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:203
				// _ = "end of CoverTab[10902]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:203
			}()
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:203
			// _ = "end of CoverTab[10901]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:203
		}() || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:203
			_go_fuzz_dep_.CoverTab[10903]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:203
			return 'a' <= c && func() bool {
													_go_fuzz_dep_.CoverTab[10904]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:204
				return c <= 'z'
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:204
				// _ = "end of CoverTab[10904]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:204
			}()
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:204
			// _ = "end of CoverTab[10903]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:204
		}() || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:204
			_go_fuzz_dep_.CoverTab[10905]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:204
			return '0' <= c && func() bool {
													_go_fuzz_dep_.CoverTab[10906]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:205
				return c <= '9'
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:205
				// _ = "end of CoverTab[10906]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:205
			}()
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:205
			// _ = "end of CoverTab[10905]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:205
		}() {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:205
			_go_fuzz_dep_.CoverTab[10907]++
												buf = append(buf, c)
												continue
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:207
			// _ = "end of CoverTab[10907]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:208
			_go_fuzz_dep_.CoverTab[10908]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:208
			// _ = "end of CoverTab[10908]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:208
		}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:208
		// _ = "end of CoverTab[10897]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:208
		_go_fuzz_dep_.CoverTab[10898]++

											upper := c >> 4
											lower := (c << 4) >> 4
											buf = append(
			buf,
			'\\',
			'x',
			hexDigits[upper],
			hexDigits[lower],
		)
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:218
		// _ = "end of CoverTab[10898]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:219
	// _ = "end of CoverTab[10895]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:219
	_go_fuzz_dep_.CoverTab[10896]++
										return string(buf)
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:220
	// _ = "end of CoverTab[10896]"
}

func printUint16(i uint16) string {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:223
	_go_fuzz_dep_.CoverTab[10909]++
										return printUint32(uint32(i))
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:224
	// _ = "end of CoverTab[10909]"
}

func printUint32(i uint32) string {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:227
	_go_fuzz_dep_.CoverTab[10910]++

										buf := make([]byte, 10)
										for b, d := buf, uint32(1000000000); d > 0; d /= 10 {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:230
		_go_fuzz_dep_.CoverTab[10912]++
											b[0] = byte(i/d%10 + '0')
											if b[0] == '0' && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:232
			_go_fuzz_dep_.CoverTab[10914]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:232
			return len(b) == len(buf)
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:232
			// _ = "end of CoverTab[10914]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:232
		}() && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:232
			_go_fuzz_dep_.CoverTab[10915]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:232
			return len(buf) > 1
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:232
			// _ = "end of CoverTab[10915]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:232
		}() {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:232
			_go_fuzz_dep_.CoverTab[10916]++
												buf = buf[1:]
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:233
			// _ = "end of CoverTab[10916]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:234
			_go_fuzz_dep_.CoverTab[10917]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:234
			// _ = "end of CoverTab[10917]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:234
		}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:234
		// _ = "end of CoverTab[10912]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:234
		_go_fuzz_dep_.CoverTab[10913]++
											b = b[1:]
											i %= d
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:236
		// _ = "end of CoverTab[10913]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:237
	// _ = "end of CoverTab[10910]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:237
	_go_fuzz_dep_.CoverTab[10911]++
										return string(buf)
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:238
	// _ = "end of CoverTab[10911]"
}

func printBool(b bool) string {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:241
	_go_fuzz_dep_.CoverTab[10918]++
										if b {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:242
		_go_fuzz_dep_.CoverTab[10920]++
											return "true"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:243
		// _ = "end of CoverTab[10920]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:244
		_go_fuzz_dep_.CoverTab[10921]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:244
		// _ = "end of CoverTab[10921]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:244
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:244
	// _ = "end of CoverTab[10918]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:244
	_go_fuzz_dep_.CoverTab[10919]++
										return "false"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:245
	// _ = "end of CoverTab[10919]"
}

var (
	// ErrNotStarted indicates that the prerequisite information isn't
	// available yet because the previous records haven't been appropriately
	// parsed, skipped or finished.
	ErrNotStarted	= errors.New("parsing/packing of this type isn't available yet")

	// ErrSectionDone indicated that all records in the section have been
	// parsed or finished.
	ErrSectionDone	= errors.New("parsing/packing of this section has completed")

	errBaseLen		= errors.New("insufficient data for base length type")
	errCalcLen		= errors.New("insufficient data for calculated length type")
	errReserved		= errors.New("segment prefix is reserved")
	errTooManyPtr		= errors.New("too many pointers (>10)")
	errInvalidPtr		= errors.New("invalid pointer")
	errNilResouceBody	= errors.New("nil resource body")
	errResourceLen		= errors.New("insufficient data for resource body length")
	errSegTooLong		= errors.New("segment length too long")
	errZeroSegLen		= errors.New("zero length segment")
	errResTooLong		= errors.New("resource length too long")
	errTooManyQuestions	= errors.New("too many Questions to pack (>65535)")
	errTooManyAnswers	= errors.New("too many Answers to pack (>65535)")
	errTooManyAuthorities	= errors.New("too many Authorities to pack (>65535)")
	errTooManyAdditionals	= errors.New("too many Additionals to pack (>65535)")
	errNonCanonicalName	= errors.New("name is not in canonical format (it must end with a .)")
	errStringTooLong	= errors.New("character string exceeds maximum length (255)")
	errCompressedSRV	= errors.New("compressed name in SRV resource data")
)

// Internal constants.
const (
	// packStartingCap is the default initial buffer size allocated during
	// packing.
	//
	// The starting capacity doesn't matter too much, but most DNS responses
	// Will be <= 512 bytes as it is the limit for DNS over UDP.
	packStartingCap	= 512

	// uint16Len is the length (in bytes) of a uint16.
	uint16Len	= 2

	// uint32Len is the length (in bytes) of a uint32.
	uint32Len	= 4

	// headerLen is the length (in bytes) of a DNS header.
	//
	// A header is comprised of 6 uint16s and no padding.
	headerLen	= 6 * uint16Len
)

type nestedError struct {
	// s is the current level's error message.
	s	string

	// err is the nested error.
	err	error
}

// nestedError implements error.Error.
func (e *nestedError) Error() string {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:307
	_go_fuzz_dep_.CoverTab[10922]++
										return e.s + ": " + e.err.Error()
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:308
	// _ = "end of CoverTab[10922]"
}

// Header is a representation of a DNS message header.
type Header struct {
	ID			uint16
	Response		bool
	OpCode			OpCode
	Authoritative		bool
	Truncated		bool
	RecursionDesired	bool
	RecursionAvailable	bool
	AuthenticData		bool
	CheckingDisabled	bool
	RCode			RCode
}

func (m *Header) pack() (id uint16, bits uint16) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:325
	_go_fuzz_dep_.CoverTab[10923]++
										id = m.ID
										bits = uint16(m.OpCode)<<11 | uint16(m.RCode)
										if m.RecursionAvailable {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:328
		_go_fuzz_dep_.CoverTab[10931]++
											bits |= headerBitRA
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:329
		// _ = "end of CoverTab[10931]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:330
		_go_fuzz_dep_.CoverTab[10932]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:330
		// _ = "end of CoverTab[10932]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:330
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:330
	// _ = "end of CoverTab[10923]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:330
	_go_fuzz_dep_.CoverTab[10924]++
										if m.RecursionDesired {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:331
		_go_fuzz_dep_.CoverTab[10933]++
											bits |= headerBitRD
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:332
		// _ = "end of CoverTab[10933]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:333
		_go_fuzz_dep_.CoverTab[10934]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:333
		// _ = "end of CoverTab[10934]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:333
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:333
	// _ = "end of CoverTab[10924]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:333
	_go_fuzz_dep_.CoverTab[10925]++
										if m.Truncated {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:334
		_go_fuzz_dep_.CoverTab[10935]++
											bits |= headerBitTC
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:335
		// _ = "end of CoverTab[10935]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:336
		_go_fuzz_dep_.CoverTab[10936]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:336
		// _ = "end of CoverTab[10936]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:336
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:336
	// _ = "end of CoverTab[10925]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:336
	_go_fuzz_dep_.CoverTab[10926]++
										if m.Authoritative {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:337
		_go_fuzz_dep_.CoverTab[10937]++
											bits |= headerBitAA
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:338
		// _ = "end of CoverTab[10937]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:339
		_go_fuzz_dep_.CoverTab[10938]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:339
		// _ = "end of CoverTab[10938]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:339
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:339
	// _ = "end of CoverTab[10926]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:339
	_go_fuzz_dep_.CoverTab[10927]++
										if m.Response {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:340
		_go_fuzz_dep_.CoverTab[10939]++
											bits |= headerBitQR
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:341
		// _ = "end of CoverTab[10939]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:342
		_go_fuzz_dep_.CoverTab[10940]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:342
		// _ = "end of CoverTab[10940]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:342
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:342
	// _ = "end of CoverTab[10927]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:342
	_go_fuzz_dep_.CoverTab[10928]++
										if m.AuthenticData {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:343
		_go_fuzz_dep_.CoverTab[10941]++
											bits |= headerBitAD
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:344
		// _ = "end of CoverTab[10941]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:345
		_go_fuzz_dep_.CoverTab[10942]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:345
		// _ = "end of CoverTab[10942]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:345
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:345
	// _ = "end of CoverTab[10928]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:345
	_go_fuzz_dep_.CoverTab[10929]++
										if m.CheckingDisabled {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:346
		_go_fuzz_dep_.CoverTab[10943]++
											bits |= headerBitCD
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:347
		// _ = "end of CoverTab[10943]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:348
		_go_fuzz_dep_.CoverTab[10944]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:348
		// _ = "end of CoverTab[10944]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:348
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:348
	// _ = "end of CoverTab[10929]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:348
	_go_fuzz_dep_.CoverTab[10930]++
										return
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:349
	// _ = "end of CoverTab[10930]"
}

// GoString implements fmt.GoStringer.GoString.
func (m *Header) GoString() string {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:353
	_go_fuzz_dep_.CoverTab[10945]++
										return "dnsmessage.Header{" +
		"ID: " + printUint16(m.ID) + ", " +
		"Response: " + printBool(m.Response) + ", " +
		"OpCode: " + m.OpCode.GoString() + ", " +
		"Authoritative: " + printBool(m.Authoritative) + ", " +
		"Truncated: " + printBool(m.Truncated) + ", " +
		"RecursionDesired: " + printBool(m.RecursionDesired) + ", " +
		"RecursionAvailable: " + printBool(m.RecursionAvailable) + ", " +
		"RCode: " + m.RCode.GoString() + "}"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:362
	// _ = "end of CoverTab[10945]"
}

// Message is a representation of a DNS message.
type Message struct {
	Header
	Questions	[]Question
	Answers		[]Resource
	Authorities	[]Resource
	Additionals	[]Resource
}

type section uint8

const (
	sectionNotStarted	section	= iota
	sectionHeader
	sectionQuestions
	sectionAnswers
	sectionAuthorities
	sectionAdditionals
	sectionDone

	headerBitQR	= 1 << 15	// query/response (response=1)
	headerBitAA	= 1 << 10	// authoritative
	headerBitTC	= 1 << 9	// truncated
	headerBitRD	= 1 << 8	// recursion desired
	headerBitRA	= 1 << 7	// recursion available
	headerBitAD	= 1 << 5	// authentic data
	headerBitCD	= 1 << 4	// checking disabled
)

var sectionNames = map[section]string{
	sectionHeader:		"header",
	sectionQuestions:	"Question",
	sectionAnswers:		"Answer",
	sectionAuthorities:	"Authority",
	sectionAdditionals:	"Additional",
}

// header is the wire format for a DNS message header.
type header struct {
	id		uint16
	bits		uint16
	questions	uint16
	answers		uint16
	authorities	uint16
	additionals	uint16
}

func (h *header) count(sec section) uint16 {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:412
	_go_fuzz_dep_.CoverTab[10946]++
										switch sec {
	case sectionQuestions:
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:414
		_go_fuzz_dep_.CoverTab[10948]++
											return h.questions
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:415
		// _ = "end of CoverTab[10948]"
	case sectionAnswers:
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:416
		_go_fuzz_dep_.CoverTab[10949]++
											return h.answers
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:417
		// _ = "end of CoverTab[10949]"
	case sectionAuthorities:
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:418
		_go_fuzz_dep_.CoverTab[10950]++
											return h.authorities
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:419
		// _ = "end of CoverTab[10950]"
	case sectionAdditionals:
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:420
		_go_fuzz_dep_.CoverTab[10951]++
											return h.additionals
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:421
		// _ = "end of CoverTab[10951]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:421
	default:
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:421
		_go_fuzz_dep_.CoverTab[10952]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:421
		// _ = "end of CoverTab[10952]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:422
	// _ = "end of CoverTab[10946]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:422
	_go_fuzz_dep_.CoverTab[10947]++
										return 0
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:423
	// _ = "end of CoverTab[10947]"
}

// pack appends the wire format of the header to msg.
func (h *header) pack(msg []byte) []byte {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:427
	_go_fuzz_dep_.CoverTab[10953]++
										msg = packUint16(msg, h.id)
										msg = packUint16(msg, h.bits)
										msg = packUint16(msg, h.questions)
										msg = packUint16(msg, h.answers)
										msg = packUint16(msg, h.authorities)
										return packUint16(msg, h.additionals)
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:433
	// _ = "end of CoverTab[10953]"
}

func (h *header) unpack(msg []byte, off int) (int, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:436
	_go_fuzz_dep_.CoverTab[10954]++
										newOff := off
										var err error
										if h.id, newOff, err = unpackUint16(msg, newOff); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:439
		_go_fuzz_dep_.CoverTab[10961]++
											return off, &nestedError{"id", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:440
		// _ = "end of CoverTab[10961]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:441
		_go_fuzz_dep_.CoverTab[10962]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:441
		// _ = "end of CoverTab[10962]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:441
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:441
	// _ = "end of CoverTab[10954]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:441
	_go_fuzz_dep_.CoverTab[10955]++
										if h.bits, newOff, err = unpackUint16(msg, newOff); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:442
		_go_fuzz_dep_.CoverTab[10963]++
											return off, &nestedError{"bits", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:443
		// _ = "end of CoverTab[10963]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:444
		_go_fuzz_dep_.CoverTab[10964]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:444
		// _ = "end of CoverTab[10964]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:444
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:444
	// _ = "end of CoverTab[10955]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:444
	_go_fuzz_dep_.CoverTab[10956]++
										if h.questions, newOff, err = unpackUint16(msg, newOff); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:445
		_go_fuzz_dep_.CoverTab[10965]++
											return off, &nestedError{"questions", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:446
		// _ = "end of CoverTab[10965]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:447
		_go_fuzz_dep_.CoverTab[10966]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:447
		// _ = "end of CoverTab[10966]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:447
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:447
	// _ = "end of CoverTab[10956]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:447
	_go_fuzz_dep_.CoverTab[10957]++
										if h.answers, newOff, err = unpackUint16(msg, newOff); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:448
		_go_fuzz_dep_.CoverTab[10967]++
											return off, &nestedError{"answers", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:449
		// _ = "end of CoverTab[10967]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:450
		_go_fuzz_dep_.CoverTab[10968]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:450
		// _ = "end of CoverTab[10968]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:450
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:450
	// _ = "end of CoverTab[10957]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:450
	_go_fuzz_dep_.CoverTab[10958]++
										if h.authorities, newOff, err = unpackUint16(msg, newOff); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:451
		_go_fuzz_dep_.CoverTab[10969]++
											return off, &nestedError{"authorities", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:452
		// _ = "end of CoverTab[10969]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:453
		_go_fuzz_dep_.CoverTab[10970]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:453
		// _ = "end of CoverTab[10970]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:453
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:453
	// _ = "end of CoverTab[10958]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:453
	_go_fuzz_dep_.CoverTab[10959]++
										if h.additionals, newOff, err = unpackUint16(msg, newOff); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:454
		_go_fuzz_dep_.CoverTab[10971]++
											return off, &nestedError{"additionals", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:455
		// _ = "end of CoverTab[10971]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:456
		_go_fuzz_dep_.CoverTab[10972]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:456
		// _ = "end of CoverTab[10972]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:456
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:456
	// _ = "end of CoverTab[10959]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:456
	_go_fuzz_dep_.CoverTab[10960]++
										return newOff, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:457
	// _ = "end of CoverTab[10960]"
}

func (h *header) header() Header {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:460
	_go_fuzz_dep_.CoverTab[10973]++
										return Header{
		ID:			h.id,
		Response:		(h.bits & headerBitQR) != 0,
		OpCode:			OpCode(h.bits>>11) & 0xF,
		Authoritative:		(h.bits & headerBitAA) != 0,
		Truncated:		(h.bits & headerBitTC) != 0,
		RecursionDesired:	(h.bits & headerBitRD) != 0,
		RecursionAvailable:	(h.bits & headerBitRA) != 0,
		AuthenticData:		(h.bits & headerBitAD) != 0,
		CheckingDisabled:	(h.bits & headerBitCD) != 0,
		RCode:			RCode(h.bits & 0xF),
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:472
	// _ = "end of CoverTab[10973]"
}

// A Resource is a DNS resource record.
type Resource struct {
	Header	ResourceHeader
	Body	ResourceBody
}

func (r *Resource) GoString() string {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:481
	_go_fuzz_dep_.CoverTab[10974]++
										return "dnsmessage.Resource{" +
		"Header: " + r.Header.GoString() +
		", Body: &" + r.Body.GoString() +
		"}"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:485
	// _ = "end of CoverTab[10974]"
}

// A ResourceBody is a DNS resource record minus the header.
type ResourceBody interface {
	// pack packs a Resource except for its header.
	pack(msg []byte, compression map[string]int, compressionOff int) ([]byte, error)

	// realType returns the actual type of the Resource. This is used to
	// fill in the header Type field.
	realType() Type

	// GoString implements fmt.GoStringer.GoString.
	GoString() string
}

// pack appends the wire format of the Resource to msg.
func (r *Resource) pack(msg []byte, compression map[string]int, compressionOff int) ([]byte, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:502
	_go_fuzz_dep_.CoverTab[10975]++
										if r.Body == nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:503
		_go_fuzz_dep_.CoverTab[10980]++
											return msg, errNilResouceBody
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:504
		// _ = "end of CoverTab[10980]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:505
		_go_fuzz_dep_.CoverTab[10981]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:505
		// _ = "end of CoverTab[10981]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:505
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:505
	// _ = "end of CoverTab[10975]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:505
	_go_fuzz_dep_.CoverTab[10976]++
										oldMsg := msg
										r.Header.Type = r.Body.realType()
										msg, lenOff, err := r.Header.pack(msg, compression, compressionOff)
										if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:509
		_go_fuzz_dep_.CoverTab[10982]++
											return msg, &nestedError{"ResourceHeader", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:510
		// _ = "end of CoverTab[10982]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:511
		_go_fuzz_dep_.CoverTab[10983]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:511
		// _ = "end of CoverTab[10983]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:511
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:511
	// _ = "end of CoverTab[10976]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:511
	_go_fuzz_dep_.CoverTab[10977]++
										preLen := len(msg)
										msg, err = r.Body.pack(msg, compression, compressionOff)
										if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:514
		_go_fuzz_dep_.CoverTab[10984]++
											return msg, &nestedError{"content", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:515
		// _ = "end of CoverTab[10984]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:516
		_go_fuzz_dep_.CoverTab[10985]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:516
		// _ = "end of CoverTab[10985]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:516
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:516
	// _ = "end of CoverTab[10977]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:516
	_go_fuzz_dep_.CoverTab[10978]++
										if err := r.Header.fixLen(msg, lenOff, preLen); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:517
		_go_fuzz_dep_.CoverTab[10986]++
											return oldMsg, err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:518
		// _ = "end of CoverTab[10986]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:519
		_go_fuzz_dep_.CoverTab[10987]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:519
		// _ = "end of CoverTab[10987]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:519
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:519
	// _ = "end of CoverTab[10978]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:519
	_go_fuzz_dep_.CoverTab[10979]++
										return msg, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:520
	// _ = "end of CoverTab[10979]"
}

// A Parser allows incrementally parsing a DNS message.
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:523
//
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:523
// When parsing is started, the Header is parsed. Next, each Question can be
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:523
// either parsed or skipped. Alternatively, all Questions can be skipped at
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:523
// once. When all Questions have been parsed, attempting to parse Questions
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:523
// will return (nil, nil) and attempting to skip Questions will return
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:523
// (true, nil). After all Questions have been either parsed or skipped, all
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:523
// Answers, Authorities and Additionals can be either parsed or skipped in the
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:523
// same way, and each type of Resource must be fully parsed or skipped before
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:523
// proceeding to the next type of Resource.
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:523
//
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:523
// Note that there is no requirement to fully skip or parse the message.
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:535
type Parser struct {
	msg	[]byte
	header	header

	section		section
	off		int
	index		int
	resHeaderValid	bool
	resHeader	ResourceHeader
}

// Start parses the header and enables the parsing of Questions.
func (p *Parser) Start(msg []byte) (Header, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:547
	_go_fuzz_dep_.CoverTab[10988]++
										if p.msg != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:548
		_go_fuzz_dep_.CoverTab[10991]++
											*p = Parser{}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:549
		// _ = "end of CoverTab[10991]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:550
		_go_fuzz_dep_.CoverTab[10992]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:550
		// _ = "end of CoverTab[10992]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:550
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:550
	// _ = "end of CoverTab[10988]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:550
	_go_fuzz_dep_.CoverTab[10989]++
										p.msg = msg
										var err error
										if p.off, err = p.header.unpack(msg, 0); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:553
		_go_fuzz_dep_.CoverTab[10993]++
											return Header{}, &nestedError{"unpacking header", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:554
		// _ = "end of CoverTab[10993]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:555
		_go_fuzz_dep_.CoverTab[10994]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:555
		// _ = "end of CoverTab[10994]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:555
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:555
	// _ = "end of CoverTab[10989]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:555
	_go_fuzz_dep_.CoverTab[10990]++
										p.section = sectionQuestions
										return p.header.header(), nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:557
	// _ = "end of CoverTab[10990]"
}

func (p *Parser) checkAdvance(sec section) error {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:560
	_go_fuzz_dep_.CoverTab[10995]++
										if p.section < sec {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:561
		_go_fuzz_dep_.CoverTab[10999]++
											return ErrNotStarted
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:562
		// _ = "end of CoverTab[10999]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:563
		_go_fuzz_dep_.CoverTab[11000]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:563
		// _ = "end of CoverTab[11000]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:563
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:563
	// _ = "end of CoverTab[10995]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:563
	_go_fuzz_dep_.CoverTab[10996]++
										if p.section > sec {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:564
		_go_fuzz_dep_.CoverTab[11001]++
											return ErrSectionDone
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:565
		// _ = "end of CoverTab[11001]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:566
		_go_fuzz_dep_.CoverTab[11002]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:566
		// _ = "end of CoverTab[11002]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:566
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:566
	// _ = "end of CoverTab[10996]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:566
	_go_fuzz_dep_.CoverTab[10997]++
										p.resHeaderValid = false
										if p.index == int(p.header.count(sec)) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:568
		_go_fuzz_dep_.CoverTab[11003]++
											p.index = 0
											p.section++
											return ErrSectionDone
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:571
		// _ = "end of CoverTab[11003]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:572
		_go_fuzz_dep_.CoverTab[11004]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:572
		// _ = "end of CoverTab[11004]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:572
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:572
	// _ = "end of CoverTab[10997]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:572
	_go_fuzz_dep_.CoverTab[10998]++
										return nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:573
	// _ = "end of CoverTab[10998]"
}

func (p *Parser) resource(sec section) (Resource, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:576
	_go_fuzz_dep_.CoverTab[11005]++
										var r Resource
										var err error
										r.Header, err = p.resourceHeader(sec)
										if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:580
		_go_fuzz_dep_.CoverTab[11008]++
											return r, err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:581
		// _ = "end of CoverTab[11008]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:582
		_go_fuzz_dep_.CoverTab[11009]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:582
		// _ = "end of CoverTab[11009]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:582
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:582
	// _ = "end of CoverTab[11005]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:582
	_go_fuzz_dep_.CoverTab[11006]++
										p.resHeaderValid = false
										r.Body, p.off, err = unpackResourceBody(p.msg, p.off, r.Header)
										if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:585
		_go_fuzz_dep_.CoverTab[11010]++
											return Resource{}, &nestedError{"unpacking " + sectionNames[sec], err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:586
		// _ = "end of CoverTab[11010]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:587
		_go_fuzz_dep_.CoverTab[11011]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:587
		// _ = "end of CoverTab[11011]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:587
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:587
	// _ = "end of CoverTab[11006]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:587
	_go_fuzz_dep_.CoverTab[11007]++
										p.index++
										return r, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:589
	// _ = "end of CoverTab[11007]"
}

func (p *Parser) resourceHeader(sec section) (ResourceHeader, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:592
	_go_fuzz_dep_.CoverTab[11012]++
										if p.resHeaderValid {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:593
		_go_fuzz_dep_.CoverTab[11016]++
											return p.resHeader, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:594
		// _ = "end of CoverTab[11016]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:595
		_go_fuzz_dep_.CoverTab[11017]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:595
		// _ = "end of CoverTab[11017]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:595
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:595
	// _ = "end of CoverTab[11012]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:595
	_go_fuzz_dep_.CoverTab[11013]++
										if err := p.checkAdvance(sec); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:596
		_go_fuzz_dep_.CoverTab[11018]++
											return ResourceHeader{}, err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:597
		// _ = "end of CoverTab[11018]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:598
		_go_fuzz_dep_.CoverTab[11019]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:598
		// _ = "end of CoverTab[11019]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:598
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:598
	// _ = "end of CoverTab[11013]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:598
	_go_fuzz_dep_.CoverTab[11014]++
										var hdr ResourceHeader
										off, err := hdr.unpack(p.msg, p.off)
										if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:601
		_go_fuzz_dep_.CoverTab[11020]++
											return ResourceHeader{}, err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:602
		// _ = "end of CoverTab[11020]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:603
		_go_fuzz_dep_.CoverTab[11021]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:603
		// _ = "end of CoverTab[11021]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:603
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:603
	// _ = "end of CoverTab[11014]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:603
	_go_fuzz_dep_.CoverTab[11015]++
										p.resHeaderValid = true
										p.resHeader = hdr
										p.off = off
										return hdr, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:607
	// _ = "end of CoverTab[11015]"
}

func (p *Parser) skipResource(sec section) error {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:610
	_go_fuzz_dep_.CoverTab[11022]++
										if p.resHeaderValid {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:611
		_go_fuzz_dep_.CoverTab[11026]++
											newOff := p.off + int(p.resHeader.Length)
											if newOff > len(p.msg) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:613
			_go_fuzz_dep_.CoverTab[11028]++
												return errResourceLen
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:614
			// _ = "end of CoverTab[11028]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:615
			_go_fuzz_dep_.CoverTab[11029]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:615
			// _ = "end of CoverTab[11029]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:615
		}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:615
		// _ = "end of CoverTab[11026]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:615
		_go_fuzz_dep_.CoverTab[11027]++
											p.off = newOff
											p.resHeaderValid = false
											p.index++
											return nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:619
		// _ = "end of CoverTab[11027]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:620
		_go_fuzz_dep_.CoverTab[11030]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:620
		// _ = "end of CoverTab[11030]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:620
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:620
	// _ = "end of CoverTab[11022]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:620
	_go_fuzz_dep_.CoverTab[11023]++
										if err := p.checkAdvance(sec); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:621
		_go_fuzz_dep_.CoverTab[11031]++
											return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:622
		// _ = "end of CoverTab[11031]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:623
		_go_fuzz_dep_.CoverTab[11032]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:623
		// _ = "end of CoverTab[11032]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:623
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:623
	// _ = "end of CoverTab[11023]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:623
	_go_fuzz_dep_.CoverTab[11024]++
										var err error
										p.off, err = skipResource(p.msg, p.off)
										if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:626
		_go_fuzz_dep_.CoverTab[11033]++
											return &nestedError{"skipping: " + sectionNames[sec], err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:627
		// _ = "end of CoverTab[11033]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:628
		_go_fuzz_dep_.CoverTab[11034]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:628
		// _ = "end of CoverTab[11034]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:628
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:628
	// _ = "end of CoverTab[11024]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:628
	_go_fuzz_dep_.CoverTab[11025]++
										p.index++
										return nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:630
	// _ = "end of CoverTab[11025]"
}

// Question parses a single Question.
func (p *Parser) Question() (Question, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:634
	_go_fuzz_dep_.CoverTab[11035]++
										if err := p.checkAdvance(sectionQuestions); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:635
		_go_fuzz_dep_.CoverTab[11040]++
											return Question{}, err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:636
		// _ = "end of CoverTab[11040]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:637
		_go_fuzz_dep_.CoverTab[11041]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:637
		// _ = "end of CoverTab[11041]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:637
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:637
	// _ = "end of CoverTab[11035]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:637
	_go_fuzz_dep_.CoverTab[11036]++
										var name Name
										off, err := name.unpack(p.msg, p.off)
										if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:640
		_go_fuzz_dep_.CoverTab[11042]++
											return Question{}, &nestedError{"unpacking Question.Name", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:641
		// _ = "end of CoverTab[11042]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:642
		_go_fuzz_dep_.CoverTab[11043]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:642
		// _ = "end of CoverTab[11043]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:642
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:642
	// _ = "end of CoverTab[11036]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:642
	_go_fuzz_dep_.CoverTab[11037]++
										typ, off, err := unpackType(p.msg, off)
										if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:644
		_go_fuzz_dep_.CoverTab[11044]++
											return Question{}, &nestedError{"unpacking Question.Type", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:645
		// _ = "end of CoverTab[11044]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:646
		_go_fuzz_dep_.CoverTab[11045]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:646
		// _ = "end of CoverTab[11045]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:646
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:646
	// _ = "end of CoverTab[11037]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:646
	_go_fuzz_dep_.CoverTab[11038]++
										class, off, err := unpackClass(p.msg, off)
										if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:648
		_go_fuzz_dep_.CoverTab[11046]++
											return Question{}, &nestedError{"unpacking Question.Class", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:649
		// _ = "end of CoverTab[11046]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:650
		_go_fuzz_dep_.CoverTab[11047]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:650
		// _ = "end of CoverTab[11047]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:650
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:650
	// _ = "end of CoverTab[11038]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:650
	_go_fuzz_dep_.CoverTab[11039]++
										p.off = off
										p.index++
										return Question{name, typ, class}, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:653
	// _ = "end of CoverTab[11039]"
}

// AllQuestions parses all Questions.
func (p *Parser) AllQuestions() ([]Question, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:657
	_go_fuzz_dep_.CoverTab[11048]++

//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:664
	qs := []Question{}
	for {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:665
		_go_fuzz_dep_.CoverTab[11049]++
											q, err := p.Question()
											if err == ErrSectionDone {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:667
			_go_fuzz_dep_.CoverTab[11052]++
												return qs, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:668
			// _ = "end of CoverTab[11052]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:669
			_go_fuzz_dep_.CoverTab[11053]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:669
			// _ = "end of CoverTab[11053]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:669
		}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:669
		// _ = "end of CoverTab[11049]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:669
		_go_fuzz_dep_.CoverTab[11050]++
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:670
			_go_fuzz_dep_.CoverTab[11054]++
												return nil, err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:671
			// _ = "end of CoverTab[11054]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:672
			_go_fuzz_dep_.CoverTab[11055]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:672
			// _ = "end of CoverTab[11055]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:672
		}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:672
		// _ = "end of CoverTab[11050]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:672
		_go_fuzz_dep_.CoverTab[11051]++
											qs = append(qs, q)
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:673
		// _ = "end of CoverTab[11051]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:674
	// _ = "end of CoverTab[11048]"
}

// SkipQuestion skips a single Question.
func (p *Parser) SkipQuestion() error {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:678
	_go_fuzz_dep_.CoverTab[11056]++
										if err := p.checkAdvance(sectionQuestions); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:679
		_go_fuzz_dep_.CoverTab[11061]++
											return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:680
		// _ = "end of CoverTab[11061]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:681
		_go_fuzz_dep_.CoverTab[11062]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:681
		// _ = "end of CoverTab[11062]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:681
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:681
	// _ = "end of CoverTab[11056]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:681
	_go_fuzz_dep_.CoverTab[11057]++
										off, err := skipName(p.msg, p.off)
										if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:683
		_go_fuzz_dep_.CoverTab[11063]++
											return &nestedError{"skipping Question Name", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:684
		// _ = "end of CoverTab[11063]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:685
		_go_fuzz_dep_.CoverTab[11064]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:685
		// _ = "end of CoverTab[11064]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:685
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:685
	// _ = "end of CoverTab[11057]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:685
	_go_fuzz_dep_.CoverTab[11058]++
										if off, err = skipType(p.msg, off); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:686
		_go_fuzz_dep_.CoverTab[11065]++
											return &nestedError{"skipping Question Type", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:687
		// _ = "end of CoverTab[11065]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:688
		_go_fuzz_dep_.CoverTab[11066]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:688
		// _ = "end of CoverTab[11066]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:688
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:688
	// _ = "end of CoverTab[11058]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:688
	_go_fuzz_dep_.CoverTab[11059]++
										if off, err = skipClass(p.msg, off); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:689
		_go_fuzz_dep_.CoverTab[11067]++
											return &nestedError{"skipping Question Class", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:690
		// _ = "end of CoverTab[11067]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:691
		_go_fuzz_dep_.CoverTab[11068]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:691
		// _ = "end of CoverTab[11068]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:691
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:691
	// _ = "end of CoverTab[11059]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:691
	_go_fuzz_dep_.CoverTab[11060]++
										p.off = off
										p.index++
										return nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:694
	// _ = "end of CoverTab[11060]"
}

// SkipAllQuestions skips all Questions.
func (p *Parser) SkipAllQuestions() error {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:698
	_go_fuzz_dep_.CoverTab[11069]++
										for {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:699
		_go_fuzz_dep_.CoverTab[11070]++
											if err := p.SkipQuestion(); err == ErrSectionDone {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:700
			_go_fuzz_dep_.CoverTab[11071]++
												return nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:701
			// _ = "end of CoverTab[11071]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:702
			_go_fuzz_dep_.CoverTab[11072]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:702
			if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:702
				_go_fuzz_dep_.CoverTab[11073]++
													return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:703
				// _ = "end of CoverTab[11073]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:704
				_go_fuzz_dep_.CoverTab[11074]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:704
				// _ = "end of CoverTab[11074]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:704
			}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:704
			// _ = "end of CoverTab[11072]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:704
		}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:704
		// _ = "end of CoverTab[11070]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:705
	// _ = "end of CoverTab[11069]"
}

// AnswerHeader parses a single Answer ResourceHeader.
func (p *Parser) AnswerHeader() (ResourceHeader, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:709
	_go_fuzz_dep_.CoverTab[11075]++
										return p.resourceHeader(sectionAnswers)
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:710
	// _ = "end of CoverTab[11075]"
}

// Answer parses a single Answer Resource.
func (p *Parser) Answer() (Resource, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:714
	_go_fuzz_dep_.CoverTab[11076]++
										return p.resource(sectionAnswers)
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:715
	// _ = "end of CoverTab[11076]"
}

// AllAnswers parses all Answer Resources.
func (p *Parser) AllAnswers() ([]Resource, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:719
	_go_fuzz_dep_.CoverTab[11077]++

//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:725
	n := int(p.header.answers)
	if n > 20 {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:726
		_go_fuzz_dep_.CoverTab[11079]++
											n = 20
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:727
		// _ = "end of CoverTab[11079]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:728
		_go_fuzz_dep_.CoverTab[11080]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:728
		// _ = "end of CoverTab[11080]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:728
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:728
	// _ = "end of CoverTab[11077]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:728
	_go_fuzz_dep_.CoverTab[11078]++
										as := make([]Resource, 0, n)
										for {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:730
		_go_fuzz_dep_.CoverTab[11081]++
											a, err := p.Answer()
											if err == ErrSectionDone {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:732
			_go_fuzz_dep_.CoverTab[11084]++
												return as, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:733
			// _ = "end of CoverTab[11084]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:734
			_go_fuzz_dep_.CoverTab[11085]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:734
			// _ = "end of CoverTab[11085]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:734
		}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:734
		// _ = "end of CoverTab[11081]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:734
		_go_fuzz_dep_.CoverTab[11082]++
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:735
			_go_fuzz_dep_.CoverTab[11086]++
												return nil, err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:736
			// _ = "end of CoverTab[11086]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:737
			_go_fuzz_dep_.CoverTab[11087]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:737
			// _ = "end of CoverTab[11087]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:737
		}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:737
		// _ = "end of CoverTab[11082]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:737
		_go_fuzz_dep_.CoverTab[11083]++
											as = append(as, a)
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:738
		// _ = "end of CoverTab[11083]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:739
	// _ = "end of CoverTab[11078]"
}

// SkipAnswer skips a single Answer Resource.
func (p *Parser) SkipAnswer() error {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:743
	_go_fuzz_dep_.CoverTab[11088]++
										return p.skipResource(sectionAnswers)
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:744
	// _ = "end of CoverTab[11088]"
}

// SkipAllAnswers skips all Answer Resources.
func (p *Parser) SkipAllAnswers() error {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:748
	_go_fuzz_dep_.CoverTab[11089]++
										for {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:749
		_go_fuzz_dep_.CoverTab[11090]++
											if err := p.SkipAnswer(); err == ErrSectionDone {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:750
			_go_fuzz_dep_.CoverTab[11091]++
												return nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:751
			// _ = "end of CoverTab[11091]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:752
			_go_fuzz_dep_.CoverTab[11092]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:752
			if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:752
				_go_fuzz_dep_.CoverTab[11093]++
													return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:753
				// _ = "end of CoverTab[11093]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:754
				_go_fuzz_dep_.CoverTab[11094]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:754
				// _ = "end of CoverTab[11094]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:754
			}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:754
			// _ = "end of CoverTab[11092]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:754
		}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:754
		// _ = "end of CoverTab[11090]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:755
	// _ = "end of CoverTab[11089]"
}

// AuthorityHeader parses a single Authority ResourceHeader.
func (p *Parser) AuthorityHeader() (ResourceHeader, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:759
	_go_fuzz_dep_.CoverTab[11095]++
										return p.resourceHeader(sectionAuthorities)
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:760
	// _ = "end of CoverTab[11095]"
}

// Authority parses a single Authority Resource.
func (p *Parser) Authority() (Resource, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:764
	_go_fuzz_dep_.CoverTab[11096]++
										return p.resource(sectionAuthorities)
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:765
	// _ = "end of CoverTab[11096]"
}

// AllAuthorities parses all Authority Resources.
func (p *Parser) AllAuthorities() ([]Resource, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:769
	_go_fuzz_dep_.CoverTab[11097]++

//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:775
	n := int(p.header.authorities)
	if n > 10 {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:776
		_go_fuzz_dep_.CoverTab[11099]++
											n = 10
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:777
		// _ = "end of CoverTab[11099]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:778
		_go_fuzz_dep_.CoverTab[11100]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:778
		// _ = "end of CoverTab[11100]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:778
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:778
	// _ = "end of CoverTab[11097]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:778
	_go_fuzz_dep_.CoverTab[11098]++
										as := make([]Resource, 0, n)
										for {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:780
		_go_fuzz_dep_.CoverTab[11101]++
											a, err := p.Authority()
											if err == ErrSectionDone {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:782
			_go_fuzz_dep_.CoverTab[11104]++
												return as, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:783
			// _ = "end of CoverTab[11104]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:784
			_go_fuzz_dep_.CoverTab[11105]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:784
			// _ = "end of CoverTab[11105]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:784
		}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:784
		// _ = "end of CoverTab[11101]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:784
		_go_fuzz_dep_.CoverTab[11102]++
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:785
			_go_fuzz_dep_.CoverTab[11106]++
												return nil, err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:786
			// _ = "end of CoverTab[11106]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:787
			_go_fuzz_dep_.CoverTab[11107]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:787
			// _ = "end of CoverTab[11107]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:787
		}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:787
		// _ = "end of CoverTab[11102]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:787
		_go_fuzz_dep_.CoverTab[11103]++
											as = append(as, a)
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:788
		// _ = "end of CoverTab[11103]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:789
	// _ = "end of CoverTab[11098]"
}

// SkipAuthority skips a single Authority Resource.
func (p *Parser) SkipAuthority() error {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:793
	_go_fuzz_dep_.CoverTab[11108]++
										return p.skipResource(sectionAuthorities)
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:794
	// _ = "end of CoverTab[11108]"
}

// SkipAllAuthorities skips all Authority Resources.
func (p *Parser) SkipAllAuthorities() error {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:798
	_go_fuzz_dep_.CoverTab[11109]++
										for {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:799
		_go_fuzz_dep_.CoverTab[11110]++
											if err := p.SkipAuthority(); err == ErrSectionDone {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:800
			_go_fuzz_dep_.CoverTab[11111]++
												return nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:801
			// _ = "end of CoverTab[11111]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:802
			_go_fuzz_dep_.CoverTab[11112]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:802
			if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:802
				_go_fuzz_dep_.CoverTab[11113]++
													return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:803
				// _ = "end of CoverTab[11113]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:804
				_go_fuzz_dep_.CoverTab[11114]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:804
				// _ = "end of CoverTab[11114]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:804
			}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:804
			// _ = "end of CoverTab[11112]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:804
		}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:804
		// _ = "end of CoverTab[11110]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:805
	// _ = "end of CoverTab[11109]"
}

// AdditionalHeader parses a single Additional ResourceHeader.
func (p *Parser) AdditionalHeader() (ResourceHeader, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:809
	_go_fuzz_dep_.CoverTab[11115]++
										return p.resourceHeader(sectionAdditionals)
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:810
	// _ = "end of CoverTab[11115]"
}

// Additional parses a single Additional Resource.
func (p *Parser) Additional() (Resource, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:814
	_go_fuzz_dep_.CoverTab[11116]++
										return p.resource(sectionAdditionals)
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:815
	// _ = "end of CoverTab[11116]"
}

// AllAdditionals parses all Additional Resources.
func (p *Parser) AllAdditionals() ([]Resource, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:819
	_go_fuzz_dep_.CoverTab[11117]++

//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:825
	n := int(p.header.additionals)
	if n > 10 {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:826
		_go_fuzz_dep_.CoverTab[11119]++
											n = 10
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:827
		// _ = "end of CoverTab[11119]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:828
		_go_fuzz_dep_.CoverTab[11120]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:828
		// _ = "end of CoverTab[11120]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:828
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:828
	// _ = "end of CoverTab[11117]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:828
	_go_fuzz_dep_.CoverTab[11118]++
										as := make([]Resource, 0, n)
										for {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:830
		_go_fuzz_dep_.CoverTab[11121]++
											a, err := p.Additional()
											if err == ErrSectionDone {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:832
			_go_fuzz_dep_.CoverTab[11124]++
												return as, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:833
			// _ = "end of CoverTab[11124]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:834
			_go_fuzz_dep_.CoverTab[11125]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:834
			// _ = "end of CoverTab[11125]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:834
		}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:834
		// _ = "end of CoverTab[11121]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:834
		_go_fuzz_dep_.CoverTab[11122]++
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:835
			_go_fuzz_dep_.CoverTab[11126]++
												return nil, err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:836
			// _ = "end of CoverTab[11126]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:837
			_go_fuzz_dep_.CoverTab[11127]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:837
			// _ = "end of CoverTab[11127]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:837
		}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:837
		// _ = "end of CoverTab[11122]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:837
		_go_fuzz_dep_.CoverTab[11123]++
											as = append(as, a)
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:838
		// _ = "end of CoverTab[11123]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:839
	// _ = "end of CoverTab[11118]"
}

// SkipAdditional skips a single Additional Resource.
func (p *Parser) SkipAdditional() error {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:843
	_go_fuzz_dep_.CoverTab[11128]++
										return p.skipResource(sectionAdditionals)
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:844
	// _ = "end of CoverTab[11128]"
}

// SkipAllAdditionals skips all Additional Resources.
func (p *Parser) SkipAllAdditionals() error {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:848
	_go_fuzz_dep_.CoverTab[11129]++
										for {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:849
		_go_fuzz_dep_.CoverTab[11130]++
											if err := p.SkipAdditional(); err == ErrSectionDone {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:850
			_go_fuzz_dep_.CoverTab[11131]++
												return nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:851
			// _ = "end of CoverTab[11131]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:852
			_go_fuzz_dep_.CoverTab[11132]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:852
			if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:852
				_go_fuzz_dep_.CoverTab[11133]++
													return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:853
				// _ = "end of CoverTab[11133]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:854
				_go_fuzz_dep_.CoverTab[11134]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:854
				// _ = "end of CoverTab[11134]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:854
			}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:854
			// _ = "end of CoverTab[11132]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:854
		}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:854
		// _ = "end of CoverTab[11130]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:855
	// _ = "end of CoverTab[11129]"
}

// CNAMEResource parses a single CNAMEResource.
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:858
//
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:858
// One of the XXXHeader methods must have been called before calling this
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:858
// method.
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:862
func (p *Parser) CNAMEResource() (CNAMEResource, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:862
	_go_fuzz_dep_.CoverTab[11135]++
										if !p.resHeaderValid || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:863
		_go_fuzz_dep_.CoverTab[11138]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:863
		return p.resHeader.Type != TypeCNAME
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:863
		// _ = "end of CoverTab[11138]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:863
	}() {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:863
		_go_fuzz_dep_.CoverTab[11139]++
											return CNAMEResource{}, ErrNotStarted
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:864
		// _ = "end of CoverTab[11139]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:865
		_go_fuzz_dep_.CoverTab[11140]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:865
		// _ = "end of CoverTab[11140]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:865
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:865
	// _ = "end of CoverTab[11135]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:865
	_go_fuzz_dep_.CoverTab[11136]++
										r, err := unpackCNAMEResource(p.msg, p.off)
										if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:867
		_go_fuzz_dep_.CoverTab[11141]++
											return CNAMEResource{}, err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:868
		// _ = "end of CoverTab[11141]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:869
		_go_fuzz_dep_.CoverTab[11142]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:869
		// _ = "end of CoverTab[11142]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:869
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:869
	// _ = "end of CoverTab[11136]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:869
	_go_fuzz_dep_.CoverTab[11137]++
										p.off += int(p.resHeader.Length)
										p.resHeaderValid = false
										p.index++
										return r, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:873
	// _ = "end of CoverTab[11137]"
}

// MXResource parses a single MXResource.
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:876
//
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:876
// One of the XXXHeader methods must have been called before calling this
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:876
// method.
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:880
func (p *Parser) MXResource() (MXResource, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:880
	_go_fuzz_dep_.CoverTab[11143]++
										if !p.resHeaderValid || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:881
		_go_fuzz_dep_.CoverTab[11146]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:881
		return p.resHeader.Type != TypeMX
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:881
		// _ = "end of CoverTab[11146]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:881
	}() {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:881
		_go_fuzz_dep_.CoverTab[11147]++
											return MXResource{}, ErrNotStarted
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:882
		// _ = "end of CoverTab[11147]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:883
		_go_fuzz_dep_.CoverTab[11148]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:883
		// _ = "end of CoverTab[11148]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:883
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:883
	// _ = "end of CoverTab[11143]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:883
	_go_fuzz_dep_.CoverTab[11144]++
										r, err := unpackMXResource(p.msg, p.off)
										if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:885
		_go_fuzz_dep_.CoverTab[11149]++
											return MXResource{}, err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:886
		// _ = "end of CoverTab[11149]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:887
		_go_fuzz_dep_.CoverTab[11150]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:887
		// _ = "end of CoverTab[11150]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:887
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:887
	// _ = "end of CoverTab[11144]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:887
	_go_fuzz_dep_.CoverTab[11145]++
										p.off += int(p.resHeader.Length)
										p.resHeaderValid = false
										p.index++
										return r, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:891
	// _ = "end of CoverTab[11145]"
}

// NSResource parses a single NSResource.
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:894
//
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:894
// One of the XXXHeader methods must have been called before calling this
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:894
// method.
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:898
func (p *Parser) NSResource() (NSResource, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:898
	_go_fuzz_dep_.CoverTab[11151]++
										if !p.resHeaderValid || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:899
		_go_fuzz_dep_.CoverTab[11154]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:899
		return p.resHeader.Type != TypeNS
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:899
		// _ = "end of CoverTab[11154]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:899
	}() {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:899
		_go_fuzz_dep_.CoverTab[11155]++
											return NSResource{}, ErrNotStarted
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:900
		// _ = "end of CoverTab[11155]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:901
		_go_fuzz_dep_.CoverTab[11156]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:901
		// _ = "end of CoverTab[11156]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:901
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:901
	// _ = "end of CoverTab[11151]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:901
	_go_fuzz_dep_.CoverTab[11152]++
										r, err := unpackNSResource(p.msg, p.off)
										if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:903
		_go_fuzz_dep_.CoverTab[11157]++
											return NSResource{}, err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:904
		// _ = "end of CoverTab[11157]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:905
		_go_fuzz_dep_.CoverTab[11158]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:905
		// _ = "end of CoverTab[11158]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:905
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:905
	// _ = "end of CoverTab[11152]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:905
	_go_fuzz_dep_.CoverTab[11153]++
										p.off += int(p.resHeader.Length)
										p.resHeaderValid = false
										p.index++
										return r, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:909
	// _ = "end of CoverTab[11153]"
}

// PTRResource parses a single PTRResource.
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:912
//
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:912
// One of the XXXHeader methods must have been called before calling this
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:912
// method.
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:916
func (p *Parser) PTRResource() (PTRResource, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:916
	_go_fuzz_dep_.CoverTab[11159]++
										if !p.resHeaderValid || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:917
		_go_fuzz_dep_.CoverTab[11162]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:917
		return p.resHeader.Type != TypePTR
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:917
		// _ = "end of CoverTab[11162]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:917
	}() {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:917
		_go_fuzz_dep_.CoverTab[11163]++
											return PTRResource{}, ErrNotStarted
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:918
		// _ = "end of CoverTab[11163]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:919
		_go_fuzz_dep_.CoverTab[11164]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:919
		// _ = "end of CoverTab[11164]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:919
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:919
	// _ = "end of CoverTab[11159]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:919
	_go_fuzz_dep_.CoverTab[11160]++
										r, err := unpackPTRResource(p.msg, p.off)
										if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:921
		_go_fuzz_dep_.CoverTab[11165]++
											return PTRResource{}, err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:922
		// _ = "end of CoverTab[11165]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:923
		_go_fuzz_dep_.CoverTab[11166]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:923
		// _ = "end of CoverTab[11166]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:923
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:923
	// _ = "end of CoverTab[11160]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:923
	_go_fuzz_dep_.CoverTab[11161]++
										p.off += int(p.resHeader.Length)
										p.resHeaderValid = false
										p.index++
										return r, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:927
	// _ = "end of CoverTab[11161]"
}

// SOAResource parses a single SOAResource.
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:930
//
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:930
// One of the XXXHeader methods must have been called before calling this
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:930
// method.
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:934
func (p *Parser) SOAResource() (SOAResource, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:934
	_go_fuzz_dep_.CoverTab[11167]++
										if !p.resHeaderValid || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:935
		_go_fuzz_dep_.CoverTab[11170]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:935
		return p.resHeader.Type != TypeSOA
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:935
		// _ = "end of CoverTab[11170]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:935
	}() {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:935
		_go_fuzz_dep_.CoverTab[11171]++
											return SOAResource{}, ErrNotStarted
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:936
		// _ = "end of CoverTab[11171]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:937
		_go_fuzz_dep_.CoverTab[11172]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:937
		// _ = "end of CoverTab[11172]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:937
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:937
	// _ = "end of CoverTab[11167]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:937
	_go_fuzz_dep_.CoverTab[11168]++
										r, err := unpackSOAResource(p.msg, p.off)
										if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:939
		_go_fuzz_dep_.CoverTab[11173]++
											return SOAResource{}, err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:940
		// _ = "end of CoverTab[11173]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:941
		_go_fuzz_dep_.CoverTab[11174]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:941
		// _ = "end of CoverTab[11174]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:941
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:941
	// _ = "end of CoverTab[11168]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:941
	_go_fuzz_dep_.CoverTab[11169]++
										p.off += int(p.resHeader.Length)
										p.resHeaderValid = false
										p.index++
										return r, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:945
	// _ = "end of CoverTab[11169]"
}

// TXTResource parses a single TXTResource.
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:948
//
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:948
// One of the XXXHeader methods must have been called before calling this
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:948
// method.
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:952
func (p *Parser) TXTResource() (TXTResource, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:952
	_go_fuzz_dep_.CoverTab[11175]++
										if !p.resHeaderValid || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:953
		_go_fuzz_dep_.CoverTab[11178]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:953
		return p.resHeader.Type != TypeTXT
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:953
		// _ = "end of CoverTab[11178]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:953
	}() {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:953
		_go_fuzz_dep_.CoverTab[11179]++
											return TXTResource{}, ErrNotStarted
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:954
		// _ = "end of CoverTab[11179]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:955
		_go_fuzz_dep_.CoverTab[11180]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:955
		// _ = "end of CoverTab[11180]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:955
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:955
	// _ = "end of CoverTab[11175]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:955
	_go_fuzz_dep_.CoverTab[11176]++
										r, err := unpackTXTResource(p.msg, p.off, p.resHeader.Length)
										if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:957
		_go_fuzz_dep_.CoverTab[11181]++
											return TXTResource{}, err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:958
		// _ = "end of CoverTab[11181]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:959
		_go_fuzz_dep_.CoverTab[11182]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:959
		// _ = "end of CoverTab[11182]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:959
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:959
	// _ = "end of CoverTab[11176]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:959
	_go_fuzz_dep_.CoverTab[11177]++
										p.off += int(p.resHeader.Length)
										p.resHeaderValid = false
										p.index++
										return r, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:963
	// _ = "end of CoverTab[11177]"
}

// SRVResource parses a single SRVResource.
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:966
//
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:966
// One of the XXXHeader methods must have been called before calling this
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:966
// method.
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:970
func (p *Parser) SRVResource() (SRVResource, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:970
	_go_fuzz_dep_.CoverTab[11183]++
										if !p.resHeaderValid || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:971
		_go_fuzz_dep_.CoverTab[11186]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:971
		return p.resHeader.Type != TypeSRV
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:971
		// _ = "end of CoverTab[11186]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:971
	}() {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:971
		_go_fuzz_dep_.CoverTab[11187]++
											return SRVResource{}, ErrNotStarted
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:972
		// _ = "end of CoverTab[11187]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:973
		_go_fuzz_dep_.CoverTab[11188]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:973
		// _ = "end of CoverTab[11188]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:973
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:973
	// _ = "end of CoverTab[11183]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:973
	_go_fuzz_dep_.CoverTab[11184]++
										r, err := unpackSRVResource(p.msg, p.off)
										if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:975
		_go_fuzz_dep_.CoverTab[11189]++
											return SRVResource{}, err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:976
		// _ = "end of CoverTab[11189]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:977
		_go_fuzz_dep_.CoverTab[11190]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:977
		// _ = "end of CoverTab[11190]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:977
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:977
	// _ = "end of CoverTab[11184]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:977
	_go_fuzz_dep_.CoverTab[11185]++
										p.off += int(p.resHeader.Length)
										p.resHeaderValid = false
										p.index++
										return r, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:981
	// _ = "end of CoverTab[11185]"
}

// AResource parses a single AResource.
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:984
//
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:984
// One of the XXXHeader methods must have been called before calling this
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:984
// method.
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:988
func (p *Parser) AResource() (AResource, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:988
	_go_fuzz_dep_.CoverTab[11191]++
										if !p.resHeaderValid || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:989
		_go_fuzz_dep_.CoverTab[11194]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:989
		return p.resHeader.Type != TypeA
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:989
		// _ = "end of CoverTab[11194]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:989
	}() {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:989
		_go_fuzz_dep_.CoverTab[11195]++
											return AResource{}, ErrNotStarted
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:990
		// _ = "end of CoverTab[11195]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:991
		_go_fuzz_dep_.CoverTab[11196]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:991
		// _ = "end of CoverTab[11196]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:991
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:991
	// _ = "end of CoverTab[11191]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:991
	_go_fuzz_dep_.CoverTab[11192]++
										r, err := unpackAResource(p.msg, p.off)
										if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:993
		_go_fuzz_dep_.CoverTab[11197]++
											return AResource{}, err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:994
		// _ = "end of CoverTab[11197]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:995
		_go_fuzz_dep_.CoverTab[11198]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:995
		// _ = "end of CoverTab[11198]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:995
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:995
	// _ = "end of CoverTab[11192]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:995
	_go_fuzz_dep_.CoverTab[11193]++
										p.off += int(p.resHeader.Length)
										p.resHeaderValid = false
										p.index++
										return r, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:999
	// _ = "end of CoverTab[11193]"
}

// AAAAResource parses a single AAAAResource.
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1002
//
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1002
// One of the XXXHeader methods must have been called before calling this
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1002
// method.
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1006
func (p *Parser) AAAAResource() (AAAAResource, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1006
	_go_fuzz_dep_.CoverTab[11199]++
											if !p.resHeaderValid || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1007
		_go_fuzz_dep_.CoverTab[11202]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1007
		return p.resHeader.Type != TypeAAAA
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1007
		// _ = "end of CoverTab[11202]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1007
	}() {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1007
		_go_fuzz_dep_.CoverTab[11203]++
												return AAAAResource{}, ErrNotStarted
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1008
		// _ = "end of CoverTab[11203]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1009
		_go_fuzz_dep_.CoverTab[11204]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1009
		// _ = "end of CoverTab[11204]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1009
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1009
	// _ = "end of CoverTab[11199]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1009
	_go_fuzz_dep_.CoverTab[11200]++
											r, err := unpackAAAAResource(p.msg, p.off)
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1011
		_go_fuzz_dep_.CoverTab[11205]++
												return AAAAResource{}, err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1012
		// _ = "end of CoverTab[11205]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1013
		_go_fuzz_dep_.CoverTab[11206]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1013
		// _ = "end of CoverTab[11206]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1013
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1013
	// _ = "end of CoverTab[11200]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1013
	_go_fuzz_dep_.CoverTab[11201]++
											p.off += int(p.resHeader.Length)
											p.resHeaderValid = false
											p.index++
											return r, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1017
	// _ = "end of CoverTab[11201]"
}

// OPTResource parses a single OPTResource.
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1020
//
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1020
// One of the XXXHeader methods must have been called before calling this
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1020
// method.
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1024
func (p *Parser) OPTResource() (OPTResource, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1024
	_go_fuzz_dep_.CoverTab[11207]++
											if !p.resHeaderValid || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1025
		_go_fuzz_dep_.CoverTab[11210]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1025
		return p.resHeader.Type != TypeOPT
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1025
		// _ = "end of CoverTab[11210]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1025
	}() {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1025
		_go_fuzz_dep_.CoverTab[11211]++
												return OPTResource{}, ErrNotStarted
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1026
		// _ = "end of CoverTab[11211]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1027
		_go_fuzz_dep_.CoverTab[11212]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1027
		// _ = "end of CoverTab[11212]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1027
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1027
	// _ = "end of CoverTab[11207]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1027
	_go_fuzz_dep_.CoverTab[11208]++
											r, err := unpackOPTResource(p.msg, p.off, p.resHeader.Length)
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1029
		_go_fuzz_dep_.CoverTab[11213]++
												return OPTResource{}, err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1030
		// _ = "end of CoverTab[11213]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1031
		_go_fuzz_dep_.CoverTab[11214]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1031
		// _ = "end of CoverTab[11214]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1031
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1031
	// _ = "end of CoverTab[11208]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1031
	_go_fuzz_dep_.CoverTab[11209]++
											p.off += int(p.resHeader.Length)
											p.resHeaderValid = false
											p.index++
											return r, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1035
	// _ = "end of CoverTab[11209]"
}

// UnknownResource parses a single UnknownResource.
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1038
//
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1038
// One of the XXXHeader methods must have been called before calling this
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1038
// method.
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1042
func (p *Parser) UnknownResource() (UnknownResource, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1042
	_go_fuzz_dep_.CoverTab[11215]++
											if !p.resHeaderValid {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1043
		_go_fuzz_dep_.CoverTab[11218]++
												return UnknownResource{}, ErrNotStarted
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1044
		// _ = "end of CoverTab[11218]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1045
		_go_fuzz_dep_.CoverTab[11219]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1045
		// _ = "end of CoverTab[11219]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1045
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1045
	// _ = "end of CoverTab[11215]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1045
	_go_fuzz_dep_.CoverTab[11216]++
											r, err := unpackUnknownResource(p.resHeader.Type, p.msg, p.off, p.resHeader.Length)
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1047
		_go_fuzz_dep_.CoverTab[11220]++
												return UnknownResource{}, err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1048
		// _ = "end of CoverTab[11220]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1049
		_go_fuzz_dep_.CoverTab[11221]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1049
		// _ = "end of CoverTab[11221]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1049
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1049
	// _ = "end of CoverTab[11216]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1049
	_go_fuzz_dep_.CoverTab[11217]++
											p.off += int(p.resHeader.Length)
											p.resHeaderValid = false
											p.index++
											return r, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1053
	// _ = "end of CoverTab[11217]"
}

// Unpack parses a full Message.
func (m *Message) Unpack(msg []byte) error {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1057
	_go_fuzz_dep_.CoverTab[11222]++
											var p Parser
											var err error
											if m.Header, err = p.Start(msg); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1060
		_go_fuzz_dep_.CoverTab[11228]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1061
		// _ = "end of CoverTab[11228]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1062
		_go_fuzz_dep_.CoverTab[11229]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1062
		// _ = "end of CoverTab[11229]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1062
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1062
	// _ = "end of CoverTab[11222]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1062
	_go_fuzz_dep_.CoverTab[11223]++
											if m.Questions, err = p.AllQuestions(); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1063
		_go_fuzz_dep_.CoverTab[11230]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1064
		// _ = "end of CoverTab[11230]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1065
		_go_fuzz_dep_.CoverTab[11231]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1065
		// _ = "end of CoverTab[11231]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1065
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1065
	// _ = "end of CoverTab[11223]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1065
	_go_fuzz_dep_.CoverTab[11224]++
											if m.Answers, err = p.AllAnswers(); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1066
		_go_fuzz_dep_.CoverTab[11232]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1067
		// _ = "end of CoverTab[11232]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1068
		_go_fuzz_dep_.CoverTab[11233]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1068
		// _ = "end of CoverTab[11233]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1068
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1068
	// _ = "end of CoverTab[11224]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1068
	_go_fuzz_dep_.CoverTab[11225]++
											if m.Authorities, err = p.AllAuthorities(); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1069
		_go_fuzz_dep_.CoverTab[11234]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1070
		// _ = "end of CoverTab[11234]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1071
		_go_fuzz_dep_.CoverTab[11235]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1071
		// _ = "end of CoverTab[11235]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1071
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1071
	// _ = "end of CoverTab[11225]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1071
	_go_fuzz_dep_.CoverTab[11226]++
											if m.Additionals, err = p.AllAdditionals(); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1072
		_go_fuzz_dep_.CoverTab[11236]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1073
		// _ = "end of CoverTab[11236]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1074
		_go_fuzz_dep_.CoverTab[11237]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1074
		// _ = "end of CoverTab[11237]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1074
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1074
	// _ = "end of CoverTab[11226]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1074
	_go_fuzz_dep_.CoverTab[11227]++
											return nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1075
	// _ = "end of CoverTab[11227]"
}

// Pack packs a full Message.
func (m *Message) Pack() ([]byte, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1079
	_go_fuzz_dep_.CoverTab[11238]++
											return m.AppendPack(make([]byte, 0, packStartingCap))
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1080
	// _ = "end of CoverTab[11238]"
}

// AppendPack is like Pack but appends the full Message to b and returns the
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1083
// extended buffer.
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1085
func (m *Message) AppendPack(b []byte) ([]byte, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1085
	_go_fuzz_dep_.CoverTab[11239]++

//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1089
	if len(m.Questions) > int(^uint16(0)) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1089
		_go_fuzz_dep_.CoverTab[11248]++
												return nil, errTooManyQuestions
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1090
		// _ = "end of CoverTab[11248]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1091
		_go_fuzz_dep_.CoverTab[11249]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1091
		// _ = "end of CoverTab[11249]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1091
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1091
	// _ = "end of CoverTab[11239]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1091
	_go_fuzz_dep_.CoverTab[11240]++
											if len(m.Answers) > int(^uint16(0)) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1092
		_go_fuzz_dep_.CoverTab[11250]++
												return nil, errTooManyAnswers
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1093
		// _ = "end of CoverTab[11250]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1094
		_go_fuzz_dep_.CoverTab[11251]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1094
		// _ = "end of CoverTab[11251]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1094
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1094
	// _ = "end of CoverTab[11240]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1094
	_go_fuzz_dep_.CoverTab[11241]++
											if len(m.Authorities) > int(^uint16(0)) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1095
		_go_fuzz_dep_.CoverTab[11252]++
												return nil, errTooManyAuthorities
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1096
		// _ = "end of CoverTab[11252]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1097
		_go_fuzz_dep_.CoverTab[11253]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1097
		// _ = "end of CoverTab[11253]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1097
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1097
	// _ = "end of CoverTab[11241]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1097
	_go_fuzz_dep_.CoverTab[11242]++
											if len(m.Additionals) > int(^uint16(0)) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1098
		_go_fuzz_dep_.CoverTab[11254]++
												return nil, errTooManyAdditionals
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1099
		// _ = "end of CoverTab[11254]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1100
		_go_fuzz_dep_.CoverTab[11255]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1100
		// _ = "end of CoverTab[11255]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1100
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1100
	// _ = "end of CoverTab[11242]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1100
	_go_fuzz_dep_.CoverTab[11243]++

											var h header
											h.id, h.bits = m.Header.pack()

											h.questions = uint16(len(m.Questions))
											h.answers = uint16(len(m.Answers))
											h.authorities = uint16(len(m.Authorities))
											h.additionals = uint16(len(m.Additionals))

											compressionOff := len(b)
											msg := h.pack(b)

//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1121
	compression := map[string]int{}

	for i := range m.Questions {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1123
		_go_fuzz_dep_.CoverTab[11256]++
												var err error
												if msg, err = m.Questions[i].pack(msg, compression, compressionOff); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1125
			_go_fuzz_dep_.CoverTab[11257]++
													return nil, &nestedError{"packing Question", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1126
			// _ = "end of CoverTab[11257]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1127
			_go_fuzz_dep_.CoverTab[11258]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1127
			// _ = "end of CoverTab[11258]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1127
		}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1127
		// _ = "end of CoverTab[11256]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1128
	// _ = "end of CoverTab[11243]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1128
	_go_fuzz_dep_.CoverTab[11244]++
											for i := range m.Answers {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1129
		_go_fuzz_dep_.CoverTab[11259]++
												var err error
												if msg, err = m.Answers[i].pack(msg, compression, compressionOff); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1131
			_go_fuzz_dep_.CoverTab[11260]++
													return nil, &nestedError{"packing Answer", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1132
			// _ = "end of CoverTab[11260]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1133
			_go_fuzz_dep_.CoverTab[11261]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1133
			// _ = "end of CoverTab[11261]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1133
		}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1133
		// _ = "end of CoverTab[11259]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1134
	// _ = "end of CoverTab[11244]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1134
	_go_fuzz_dep_.CoverTab[11245]++
											for i := range m.Authorities {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1135
		_go_fuzz_dep_.CoverTab[11262]++
												var err error
												if msg, err = m.Authorities[i].pack(msg, compression, compressionOff); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1137
			_go_fuzz_dep_.CoverTab[11263]++
													return nil, &nestedError{"packing Authority", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1138
			// _ = "end of CoverTab[11263]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1139
			_go_fuzz_dep_.CoverTab[11264]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1139
			// _ = "end of CoverTab[11264]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1139
		}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1139
		// _ = "end of CoverTab[11262]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1140
	// _ = "end of CoverTab[11245]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1140
	_go_fuzz_dep_.CoverTab[11246]++
											for i := range m.Additionals {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1141
		_go_fuzz_dep_.CoverTab[11265]++
												var err error
												if msg, err = m.Additionals[i].pack(msg, compression, compressionOff); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1143
			_go_fuzz_dep_.CoverTab[11266]++
													return nil, &nestedError{"packing Additional", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1144
			// _ = "end of CoverTab[11266]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1145
			_go_fuzz_dep_.CoverTab[11267]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1145
			// _ = "end of CoverTab[11267]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1145
		}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1145
		// _ = "end of CoverTab[11265]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1146
	// _ = "end of CoverTab[11246]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1146
	_go_fuzz_dep_.CoverTab[11247]++

											return msg, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1148
	// _ = "end of CoverTab[11247]"
}

// GoString implements fmt.GoStringer.GoString.
func (m *Message) GoString() string {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1152
	_go_fuzz_dep_.CoverTab[11268]++
											s := "dnsmessage.Message{Header: " + m.Header.GoString() + ", " +
		"Questions: []dnsmessage.Question{"
	if len(m.Questions) > 0 {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1155
		_go_fuzz_dep_.CoverTab[11273]++
												s += m.Questions[0].GoString()
												for _, q := range m.Questions[1:] {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1157
			_go_fuzz_dep_.CoverTab[11274]++
													s += ", " + q.GoString()
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1158
			// _ = "end of CoverTab[11274]"
		}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1159
		// _ = "end of CoverTab[11273]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1160
		_go_fuzz_dep_.CoverTab[11275]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1160
		// _ = "end of CoverTab[11275]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1160
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1160
	// _ = "end of CoverTab[11268]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1160
	_go_fuzz_dep_.CoverTab[11269]++
											s += "}, Answers: []dnsmessage.Resource{"
											if len(m.Answers) > 0 {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1162
		_go_fuzz_dep_.CoverTab[11276]++
												s += m.Answers[0].GoString()
												for _, a := range m.Answers[1:] {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1164
			_go_fuzz_dep_.CoverTab[11277]++
													s += ", " + a.GoString()
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1165
			// _ = "end of CoverTab[11277]"
		}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1166
		// _ = "end of CoverTab[11276]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1167
		_go_fuzz_dep_.CoverTab[11278]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1167
		// _ = "end of CoverTab[11278]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1167
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1167
	// _ = "end of CoverTab[11269]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1167
	_go_fuzz_dep_.CoverTab[11270]++
											s += "}, Authorities: []dnsmessage.Resource{"
											if len(m.Authorities) > 0 {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1169
		_go_fuzz_dep_.CoverTab[11279]++
												s += m.Authorities[0].GoString()
												for _, a := range m.Authorities[1:] {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1171
			_go_fuzz_dep_.CoverTab[11280]++
													s += ", " + a.GoString()
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1172
			// _ = "end of CoverTab[11280]"
		}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1173
		// _ = "end of CoverTab[11279]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1174
		_go_fuzz_dep_.CoverTab[11281]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1174
		// _ = "end of CoverTab[11281]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1174
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1174
	// _ = "end of CoverTab[11270]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1174
	_go_fuzz_dep_.CoverTab[11271]++
											s += "}, Additionals: []dnsmessage.Resource{"
											if len(m.Additionals) > 0 {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1176
		_go_fuzz_dep_.CoverTab[11282]++
												s += m.Additionals[0].GoString()
												for _, a := range m.Additionals[1:] {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1178
			_go_fuzz_dep_.CoverTab[11283]++
													s += ", " + a.GoString()
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1179
			// _ = "end of CoverTab[11283]"
		}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1180
		// _ = "end of CoverTab[11282]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1181
		_go_fuzz_dep_.CoverTab[11284]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1181
		// _ = "end of CoverTab[11284]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1181
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1181
	// _ = "end of CoverTab[11271]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1181
	_go_fuzz_dep_.CoverTab[11272]++
											return s + "}}"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1182
	// _ = "end of CoverTab[11272]"
}

// A Builder allows incrementally packing a DNS message.
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1185
//
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1185
// Example usage:
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1185
//
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1185
//	buf := make([]byte, 2, 514)
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1185
//	b := NewBuilder(buf, Header{...})
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1185
//	b.EnableCompression()
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1185
//	// Optionally start a section and add things to that section.
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1185
//	// Repeat adding sections as necessary.
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1185
//	buf, err := b.Finish()
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1185
//	// If err is nil, buf[2:] will contain the built bytes.
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1196
type Builder struct {
	// msg is the storage for the message being built.
	msg	[]byte

	// section keeps track of the current section being built.
	section	section

	// header keeps track of what should go in the header when Finish is
	// called.
	header	header

	// start is the starting index of the bytes allocated in msg for header.
	start	int

	// compression is a mapping from name suffixes to their starting index
	// in msg.
	compression	map[string]int
}

// NewBuilder creates a new builder with compression disabled.
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1215
//
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1215
// Note: Most users will want to immediately enable compression with the
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1215
// EnableCompression method. See that method's comment for why you may or may
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1215
// not want to enable compression.
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1215
//
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1215
// The DNS message is appended to the provided initial buffer buf (which may be
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1215
// nil) as it is built. The final message is returned by the (*Builder).Finish
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1215
// method, which includes buf[:len(buf)] and may return the same underlying
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1215
// array if there was sufficient capacity in the slice.
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1225
func NewBuilder(buf []byte, h Header) Builder {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1225
	_go_fuzz_dep_.CoverTab[11285]++
											if buf == nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1226
		_go_fuzz_dep_.CoverTab[11287]++
												buf = make([]byte, 0, packStartingCap)
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1227
		// _ = "end of CoverTab[11287]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1228
		_go_fuzz_dep_.CoverTab[11288]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1228
		// _ = "end of CoverTab[11288]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1228
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1228
	// _ = "end of CoverTab[11285]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1228
	_go_fuzz_dep_.CoverTab[11286]++
											b := Builder{msg: buf, start: len(buf)}
											b.header.id, b.header.bits = h.pack()
											var hb [headerLen]byte
											b.msg = append(b.msg, hb[:]...)
											b.section = sectionHeader
											return b
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1234
	// _ = "end of CoverTab[11286]"
}

// EnableCompression enables compression in the Builder.
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1237
//
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1237
// Leaving compression disabled avoids compression related allocations, but can
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1237
// result in larger message sizes. Be careful with this mode as it can cause
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1237
// messages to exceed the UDP size limit.
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1237
//
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1237
// According to RFC 1035, section 4.1.4, the use of compression is optional, but
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1237
// all implementations must accept both compressed and uncompressed DNS
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1237
// messages.
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1237
//
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1237
// Compression should be enabled before any sections are added for best results.
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1248
func (b *Builder) EnableCompression() {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1248
	_go_fuzz_dep_.CoverTab[11289]++
											b.compression = map[string]int{}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1249
	// _ = "end of CoverTab[11289]"
}

func (b *Builder) startCheck(s section) error {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1252
	_go_fuzz_dep_.CoverTab[11290]++
											if b.section <= sectionNotStarted {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1253
		_go_fuzz_dep_.CoverTab[11293]++
												return ErrNotStarted
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1254
		// _ = "end of CoverTab[11293]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1255
		_go_fuzz_dep_.CoverTab[11294]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1255
		// _ = "end of CoverTab[11294]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1255
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1255
	// _ = "end of CoverTab[11290]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1255
	_go_fuzz_dep_.CoverTab[11291]++
											if b.section > s {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1256
		_go_fuzz_dep_.CoverTab[11295]++
												return ErrSectionDone
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1257
		// _ = "end of CoverTab[11295]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1258
		_go_fuzz_dep_.CoverTab[11296]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1258
		// _ = "end of CoverTab[11296]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1258
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1258
	// _ = "end of CoverTab[11291]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1258
	_go_fuzz_dep_.CoverTab[11292]++
											return nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1259
	// _ = "end of CoverTab[11292]"
}

// StartQuestions prepares the builder for packing Questions.
func (b *Builder) StartQuestions() error {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1263
	_go_fuzz_dep_.CoverTab[11297]++
											if err := b.startCheck(sectionQuestions); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1264
		_go_fuzz_dep_.CoverTab[11299]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1265
		// _ = "end of CoverTab[11299]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1266
		_go_fuzz_dep_.CoverTab[11300]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1266
		// _ = "end of CoverTab[11300]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1266
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1266
	// _ = "end of CoverTab[11297]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1266
	_go_fuzz_dep_.CoverTab[11298]++
											b.section = sectionQuestions
											return nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1268
	// _ = "end of CoverTab[11298]"
}

// StartAnswers prepares the builder for packing Answers.
func (b *Builder) StartAnswers() error {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1272
	_go_fuzz_dep_.CoverTab[11301]++
											if err := b.startCheck(sectionAnswers); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1273
		_go_fuzz_dep_.CoverTab[11303]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1274
		// _ = "end of CoverTab[11303]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1275
		_go_fuzz_dep_.CoverTab[11304]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1275
		// _ = "end of CoverTab[11304]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1275
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1275
	// _ = "end of CoverTab[11301]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1275
	_go_fuzz_dep_.CoverTab[11302]++
											b.section = sectionAnswers
											return nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1277
	// _ = "end of CoverTab[11302]"
}

// StartAuthorities prepares the builder for packing Authorities.
func (b *Builder) StartAuthorities() error {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1281
	_go_fuzz_dep_.CoverTab[11305]++
											if err := b.startCheck(sectionAuthorities); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1282
		_go_fuzz_dep_.CoverTab[11307]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1283
		// _ = "end of CoverTab[11307]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1284
		_go_fuzz_dep_.CoverTab[11308]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1284
		// _ = "end of CoverTab[11308]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1284
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1284
	// _ = "end of CoverTab[11305]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1284
	_go_fuzz_dep_.CoverTab[11306]++
											b.section = sectionAuthorities
											return nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1286
	// _ = "end of CoverTab[11306]"
}

// StartAdditionals prepares the builder for packing Additionals.
func (b *Builder) StartAdditionals() error {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1290
	_go_fuzz_dep_.CoverTab[11309]++
											if err := b.startCheck(sectionAdditionals); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1291
		_go_fuzz_dep_.CoverTab[11311]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1292
		// _ = "end of CoverTab[11311]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1293
		_go_fuzz_dep_.CoverTab[11312]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1293
		// _ = "end of CoverTab[11312]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1293
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1293
	// _ = "end of CoverTab[11309]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1293
	_go_fuzz_dep_.CoverTab[11310]++
											b.section = sectionAdditionals
											return nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1295
	// _ = "end of CoverTab[11310]"
}

func (b *Builder) incrementSectionCount() error {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1298
	_go_fuzz_dep_.CoverTab[11313]++
											var count *uint16
											var err error
											switch b.section {
	case sectionQuestions:
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1302
		_go_fuzz_dep_.CoverTab[11316]++
												count = &b.header.questions
												err = errTooManyQuestions
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1304
		// _ = "end of CoverTab[11316]"
	case sectionAnswers:
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1305
		_go_fuzz_dep_.CoverTab[11317]++
												count = &b.header.answers
												err = errTooManyAnswers
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1307
		// _ = "end of CoverTab[11317]"
	case sectionAuthorities:
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1308
		_go_fuzz_dep_.CoverTab[11318]++
												count = &b.header.authorities
												err = errTooManyAuthorities
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1310
		// _ = "end of CoverTab[11318]"
	case sectionAdditionals:
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1311
		_go_fuzz_dep_.CoverTab[11319]++
												count = &b.header.additionals
												err = errTooManyAdditionals
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1313
		// _ = "end of CoverTab[11319]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1313
	default:
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1313
		_go_fuzz_dep_.CoverTab[11320]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1313
		// _ = "end of CoverTab[11320]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1314
	// _ = "end of CoverTab[11313]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1314
	_go_fuzz_dep_.CoverTab[11314]++
											if *count == ^uint16(0) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1315
		_go_fuzz_dep_.CoverTab[11321]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1316
		// _ = "end of CoverTab[11321]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1317
		_go_fuzz_dep_.CoverTab[11322]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1317
		// _ = "end of CoverTab[11322]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1317
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1317
	// _ = "end of CoverTab[11314]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1317
	_go_fuzz_dep_.CoverTab[11315]++
											*count++
											return nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1319
	// _ = "end of CoverTab[11315]"
}

// Question adds a single Question.
func (b *Builder) Question(q Question) error {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1323
	_go_fuzz_dep_.CoverTab[11323]++
											if b.section < sectionQuestions {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1324
		_go_fuzz_dep_.CoverTab[11328]++
												return ErrNotStarted
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1325
		// _ = "end of CoverTab[11328]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1326
		_go_fuzz_dep_.CoverTab[11329]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1326
		// _ = "end of CoverTab[11329]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1326
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1326
	// _ = "end of CoverTab[11323]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1326
	_go_fuzz_dep_.CoverTab[11324]++
											if b.section > sectionQuestions {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1327
		_go_fuzz_dep_.CoverTab[11330]++
												return ErrSectionDone
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1328
		// _ = "end of CoverTab[11330]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1329
		_go_fuzz_dep_.CoverTab[11331]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1329
		// _ = "end of CoverTab[11331]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1329
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1329
	// _ = "end of CoverTab[11324]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1329
	_go_fuzz_dep_.CoverTab[11325]++
											msg, err := q.pack(b.msg, b.compression, b.start)
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1331
		_go_fuzz_dep_.CoverTab[11332]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1332
		// _ = "end of CoverTab[11332]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1333
		_go_fuzz_dep_.CoverTab[11333]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1333
		// _ = "end of CoverTab[11333]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1333
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1333
	// _ = "end of CoverTab[11325]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1333
	_go_fuzz_dep_.CoverTab[11326]++
											if err := b.incrementSectionCount(); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1334
		_go_fuzz_dep_.CoverTab[11334]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1335
		// _ = "end of CoverTab[11334]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1336
		_go_fuzz_dep_.CoverTab[11335]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1336
		// _ = "end of CoverTab[11335]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1336
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1336
	// _ = "end of CoverTab[11326]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1336
	_go_fuzz_dep_.CoverTab[11327]++
											b.msg = msg
											return nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1338
	// _ = "end of CoverTab[11327]"
}

func (b *Builder) checkResourceSection() error {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1341
	_go_fuzz_dep_.CoverTab[11336]++
											if b.section < sectionAnswers {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1342
		_go_fuzz_dep_.CoverTab[11339]++
												return ErrNotStarted
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1343
		// _ = "end of CoverTab[11339]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1344
		_go_fuzz_dep_.CoverTab[11340]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1344
		// _ = "end of CoverTab[11340]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1344
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1344
	// _ = "end of CoverTab[11336]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1344
	_go_fuzz_dep_.CoverTab[11337]++
											if b.section > sectionAdditionals {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1345
		_go_fuzz_dep_.CoverTab[11341]++
												return ErrSectionDone
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1346
		// _ = "end of CoverTab[11341]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1347
		_go_fuzz_dep_.CoverTab[11342]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1347
		// _ = "end of CoverTab[11342]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1347
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1347
	// _ = "end of CoverTab[11337]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1347
	_go_fuzz_dep_.CoverTab[11338]++
											return nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1348
	// _ = "end of CoverTab[11338]"
}

// CNAMEResource adds a single CNAMEResource.
func (b *Builder) CNAMEResource(h ResourceHeader, r CNAMEResource) error {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1352
	_go_fuzz_dep_.CoverTab[11343]++
											if err := b.checkResourceSection(); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1353
		_go_fuzz_dep_.CoverTab[11349]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1354
		// _ = "end of CoverTab[11349]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1355
		_go_fuzz_dep_.CoverTab[11350]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1355
		// _ = "end of CoverTab[11350]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1355
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1355
	// _ = "end of CoverTab[11343]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1355
	_go_fuzz_dep_.CoverTab[11344]++
											h.Type = r.realType()
											msg, lenOff, err := h.pack(b.msg, b.compression, b.start)
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1358
		_go_fuzz_dep_.CoverTab[11351]++
												return &nestedError{"ResourceHeader", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1359
		// _ = "end of CoverTab[11351]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1360
		_go_fuzz_dep_.CoverTab[11352]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1360
		// _ = "end of CoverTab[11352]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1360
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1360
	// _ = "end of CoverTab[11344]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1360
	_go_fuzz_dep_.CoverTab[11345]++
											preLen := len(msg)
											if msg, err = r.pack(msg, b.compression, b.start); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1362
		_go_fuzz_dep_.CoverTab[11353]++
												return &nestedError{"CNAMEResource body", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1363
		// _ = "end of CoverTab[11353]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1364
		_go_fuzz_dep_.CoverTab[11354]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1364
		// _ = "end of CoverTab[11354]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1364
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1364
	// _ = "end of CoverTab[11345]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1364
	_go_fuzz_dep_.CoverTab[11346]++
											if err := h.fixLen(msg, lenOff, preLen); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1365
		_go_fuzz_dep_.CoverTab[11355]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1366
		// _ = "end of CoverTab[11355]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1367
		_go_fuzz_dep_.CoverTab[11356]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1367
		// _ = "end of CoverTab[11356]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1367
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1367
	// _ = "end of CoverTab[11346]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1367
	_go_fuzz_dep_.CoverTab[11347]++
											if err := b.incrementSectionCount(); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1368
		_go_fuzz_dep_.CoverTab[11357]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1369
		// _ = "end of CoverTab[11357]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1370
		_go_fuzz_dep_.CoverTab[11358]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1370
		// _ = "end of CoverTab[11358]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1370
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1370
	// _ = "end of CoverTab[11347]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1370
	_go_fuzz_dep_.CoverTab[11348]++
											b.msg = msg
											return nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1372
	// _ = "end of CoverTab[11348]"
}

// MXResource adds a single MXResource.
func (b *Builder) MXResource(h ResourceHeader, r MXResource) error {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1376
	_go_fuzz_dep_.CoverTab[11359]++
											if err := b.checkResourceSection(); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1377
		_go_fuzz_dep_.CoverTab[11365]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1378
		// _ = "end of CoverTab[11365]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1379
		_go_fuzz_dep_.CoverTab[11366]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1379
		// _ = "end of CoverTab[11366]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1379
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1379
	// _ = "end of CoverTab[11359]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1379
	_go_fuzz_dep_.CoverTab[11360]++
											h.Type = r.realType()
											msg, lenOff, err := h.pack(b.msg, b.compression, b.start)
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1382
		_go_fuzz_dep_.CoverTab[11367]++
												return &nestedError{"ResourceHeader", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1383
		// _ = "end of CoverTab[11367]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1384
		_go_fuzz_dep_.CoverTab[11368]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1384
		// _ = "end of CoverTab[11368]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1384
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1384
	// _ = "end of CoverTab[11360]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1384
	_go_fuzz_dep_.CoverTab[11361]++
											preLen := len(msg)
											if msg, err = r.pack(msg, b.compression, b.start); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1386
		_go_fuzz_dep_.CoverTab[11369]++
												return &nestedError{"MXResource body", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1387
		// _ = "end of CoverTab[11369]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1388
		_go_fuzz_dep_.CoverTab[11370]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1388
		// _ = "end of CoverTab[11370]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1388
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1388
	// _ = "end of CoverTab[11361]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1388
	_go_fuzz_dep_.CoverTab[11362]++
											if err := h.fixLen(msg, lenOff, preLen); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1389
		_go_fuzz_dep_.CoverTab[11371]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1390
		// _ = "end of CoverTab[11371]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1391
		_go_fuzz_dep_.CoverTab[11372]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1391
		// _ = "end of CoverTab[11372]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1391
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1391
	// _ = "end of CoverTab[11362]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1391
	_go_fuzz_dep_.CoverTab[11363]++
											if err := b.incrementSectionCount(); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1392
		_go_fuzz_dep_.CoverTab[11373]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1393
		// _ = "end of CoverTab[11373]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1394
		_go_fuzz_dep_.CoverTab[11374]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1394
		// _ = "end of CoverTab[11374]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1394
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1394
	// _ = "end of CoverTab[11363]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1394
	_go_fuzz_dep_.CoverTab[11364]++
											b.msg = msg
											return nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1396
	// _ = "end of CoverTab[11364]"
}

// NSResource adds a single NSResource.
func (b *Builder) NSResource(h ResourceHeader, r NSResource) error {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1400
	_go_fuzz_dep_.CoverTab[11375]++
											if err := b.checkResourceSection(); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1401
		_go_fuzz_dep_.CoverTab[11381]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1402
		// _ = "end of CoverTab[11381]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1403
		_go_fuzz_dep_.CoverTab[11382]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1403
		// _ = "end of CoverTab[11382]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1403
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1403
	// _ = "end of CoverTab[11375]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1403
	_go_fuzz_dep_.CoverTab[11376]++
											h.Type = r.realType()
											msg, lenOff, err := h.pack(b.msg, b.compression, b.start)
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1406
		_go_fuzz_dep_.CoverTab[11383]++
												return &nestedError{"ResourceHeader", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1407
		// _ = "end of CoverTab[11383]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1408
		_go_fuzz_dep_.CoverTab[11384]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1408
		// _ = "end of CoverTab[11384]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1408
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1408
	// _ = "end of CoverTab[11376]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1408
	_go_fuzz_dep_.CoverTab[11377]++
											preLen := len(msg)
											if msg, err = r.pack(msg, b.compression, b.start); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1410
		_go_fuzz_dep_.CoverTab[11385]++
												return &nestedError{"NSResource body", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1411
		// _ = "end of CoverTab[11385]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1412
		_go_fuzz_dep_.CoverTab[11386]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1412
		// _ = "end of CoverTab[11386]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1412
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1412
	// _ = "end of CoverTab[11377]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1412
	_go_fuzz_dep_.CoverTab[11378]++
											if err := h.fixLen(msg, lenOff, preLen); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1413
		_go_fuzz_dep_.CoverTab[11387]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1414
		// _ = "end of CoverTab[11387]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1415
		_go_fuzz_dep_.CoverTab[11388]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1415
		// _ = "end of CoverTab[11388]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1415
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1415
	// _ = "end of CoverTab[11378]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1415
	_go_fuzz_dep_.CoverTab[11379]++
											if err := b.incrementSectionCount(); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1416
		_go_fuzz_dep_.CoverTab[11389]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1417
		// _ = "end of CoverTab[11389]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1418
		_go_fuzz_dep_.CoverTab[11390]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1418
		// _ = "end of CoverTab[11390]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1418
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1418
	// _ = "end of CoverTab[11379]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1418
	_go_fuzz_dep_.CoverTab[11380]++
											b.msg = msg
											return nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1420
	// _ = "end of CoverTab[11380]"
}

// PTRResource adds a single PTRResource.
func (b *Builder) PTRResource(h ResourceHeader, r PTRResource) error {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1424
	_go_fuzz_dep_.CoverTab[11391]++
											if err := b.checkResourceSection(); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1425
		_go_fuzz_dep_.CoverTab[11397]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1426
		// _ = "end of CoverTab[11397]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1427
		_go_fuzz_dep_.CoverTab[11398]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1427
		// _ = "end of CoverTab[11398]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1427
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1427
	// _ = "end of CoverTab[11391]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1427
	_go_fuzz_dep_.CoverTab[11392]++
											h.Type = r.realType()
											msg, lenOff, err := h.pack(b.msg, b.compression, b.start)
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1430
		_go_fuzz_dep_.CoverTab[11399]++
												return &nestedError{"ResourceHeader", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1431
		// _ = "end of CoverTab[11399]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1432
		_go_fuzz_dep_.CoverTab[11400]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1432
		// _ = "end of CoverTab[11400]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1432
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1432
	// _ = "end of CoverTab[11392]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1432
	_go_fuzz_dep_.CoverTab[11393]++
											preLen := len(msg)
											if msg, err = r.pack(msg, b.compression, b.start); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1434
		_go_fuzz_dep_.CoverTab[11401]++
												return &nestedError{"PTRResource body", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1435
		// _ = "end of CoverTab[11401]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1436
		_go_fuzz_dep_.CoverTab[11402]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1436
		// _ = "end of CoverTab[11402]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1436
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1436
	// _ = "end of CoverTab[11393]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1436
	_go_fuzz_dep_.CoverTab[11394]++
											if err := h.fixLen(msg, lenOff, preLen); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1437
		_go_fuzz_dep_.CoverTab[11403]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1438
		// _ = "end of CoverTab[11403]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1439
		_go_fuzz_dep_.CoverTab[11404]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1439
		// _ = "end of CoverTab[11404]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1439
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1439
	// _ = "end of CoverTab[11394]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1439
	_go_fuzz_dep_.CoverTab[11395]++
											if err := b.incrementSectionCount(); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1440
		_go_fuzz_dep_.CoverTab[11405]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1441
		// _ = "end of CoverTab[11405]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1442
		_go_fuzz_dep_.CoverTab[11406]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1442
		// _ = "end of CoverTab[11406]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1442
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1442
	// _ = "end of CoverTab[11395]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1442
	_go_fuzz_dep_.CoverTab[11396]++
											b.msg = msg
											return nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1444
	// _ = "end of CoverTab[11396]"
}

// SOAResource adds a single SOAResource.
func (b *Builder) SOAResource(h ResourceHeader, r SOAResource) error {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1448
	_go_fuzz_dep_.CoverTab[11407]++
											if err := b.checkResourceSection(); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1449
		_go_fuzz_dep_.CoverTab[11413]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1450
		// _ = "end of CoverTab[11413]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1451
		_go_fuzz_dep_.CoverTab[11414]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1451
		// _ = "end of CoverTab[11414]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1451
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1451
	// _ = "end of CoverTab[11407]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1451
	_go_fuzz_dep_.CoverTab[11408]++
											h.Type = r.realType()
											msg, lenOff, err := h.pack(b.msg, b.compression, b.start)
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1454
		_go_fuzz_dep_.CoverTab[11415]++
												return &nestedError{"ResourceHeader", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1455
		// _ = "end of CoverTab[11415]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1456
		_go_fuzz_dep_.CoverTab[11416]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1456
		// _ = "end of CoverTab[11416]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1456
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1456
	// _ = "end of CoverTab[11408]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1456
	_go_fuzz_dep_.CoverTab[11409]++
											preLen := len(msg)
											if msg, err = r.pack(msg, b.compression, b.start); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1458
		_go_fuzz_dep_.CoverTab[11417]++
												return &nestedError{"SOAResource body", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1459
		// _ = "end of CoverTab[11417]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1460
		_go_fuzz_dep_.CoverTab[11418]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1460
		// _ = "end of CoverTab[11418]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1460
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1460
	// _ = "end of CoverTab[11409]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1460
	_go_fuzz_dep_.CoverTab[11410]++
											if err := h.fixLen(msg, lenOff, preLen); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1461
		_go_fuzz_dep_.CoverTab[11419]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1462
		// _ = "end of CoverTab[11419]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1463
		_go_fuzz_dep_.CoverTab[11420]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1463
		// _ = "end of CoverTab[11420]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1463
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1463
	// _ = "end of CoverTab[11410]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1463
	_go_fuzz_dep_.CoverTab[11411]++
											if err := b.incrementSectionCount(); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1464
		_go_fuzz_dep_.CoverTab[11421]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1465
		// _ = "end of CoverTab[11421]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1466
		_go_fuzz_dep_.CoverTab[11422]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1466
		// _ = "end of CoverTab[11422]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1466
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1466
	// _ = "end of CoverTab[11411]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1466
	_go_fuzz_dep_.CoverTab[11412]++
											b.msg = msg
											return nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1468
	// _ = "end of CoverTab[11412]"
}

// TXTResource adds a single TXTResource.
func (b *Builder) TXTResource(h ResourceHeader, r TXTResource) error {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1472
	_go_fuzz_dep_.CoverTab[11423]++
											if err := b.checkResourceSection(); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1473
		_go_fuzz_dep_.CoverTab[11429]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1474
		// _ = "end of CoverTab[11429]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1475
		_go_fuzz_dep_.CoverTab[11430]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1475
		// _ = "end of CoverTab[11430]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1475
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1475
	// _ = "end of CoverTab[11423]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1475
	_go_fuzz_dep_.CoverTab[11424]++
											h.Type = r.realType()
											msg, lenOff, err := h.pack(b.msg, b.compression, b.start)
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1478
		_go_fuzz_dep_.CoverTab[11431]++
												return &nestedError{"ResourceHeader", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1479
		// _ = "end of CoverTab[11431]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1480
		_go_fuzz_dep_.CoverTab[11432]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1480
		// _ = "end of CoverTab[11432]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1480
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1480
	// _ = "end of CoverTab[11424]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1480
	_go_fuzz_dep_.CoverTab[11425]++
											preLen := len(msg)
											if msg, err = r.pack(msg, b.compression, b.start); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1482
		_go_fuzz_dep_.CoverTab[11433]++
												return &nestedError{"TXTResource body", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1483
		// _ = "end of CoverTab[11433]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1484
		_go_fuzz_dep_.CoverTab[11434]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1484
		// _ = "end of CoverTab[11434]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1484
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1484
	// _ = "end of CoverTab[11425]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1484
	_go_fuzz_dep_.CoverTab[11426]++
											if err := h.fixLen(msg, lenOff, preLen); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1485
		_go_fuzz_dep_.CoverTab[11435]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1486
		// _ = "end of CoverTab[11435]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1487
		_go_fuzz_dep_.CoverTab[11436]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1487
		// _ = "end of CoverTab[11436]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1487
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1487
	// _ = "end of CoverTab[11426]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1487
	_go_fuzz_dep_.CoverTab[11427]++
											if err := b.incrementSectionCount(); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1488
		_go_fuzz_dep_.CoverTab[11437]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1489
		// _ = "end of CoverTab[11437]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1490
		_go_fuzz_dep_.CoverTab[11438]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1490
		// _ = "end of CoverTab[11438]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1490
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1490
	// _ = "end of CoverTab[11427]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1490
	_go_fuzz_dep_.CoverTab[11428]++
											b.msg = msg
											return nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1492
	// _ = "end of CoverTab[11428]"
}

// SRVResource adds a single SRVResource.
func (b *Builder) SRVResource(h ResourceHeader, r SRVResource) error {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1496
	_go_fuzz_dep_.CoverTab[11439]++
											if err := b.checkResourceSection(); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1497
		_go_fuzz_dep_.CoverTab[11445]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1498
		// _ = "end of CoverTab[11445]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1499
		_go_fuzz_dep_.CoverTab[11446]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1499
		// _ = "end of CoverTab[11446]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1499
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1499
	// _ = "end of CoverTab[11439]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1499
	_go_fuzz_dep_.CoverTab[11440]++
											h.Type = r.realType()
											msg, lenOff, err := h.pack(b.msg, b.compression, b.start)
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1502
		_go_fuzz_dep_.CoverTab[11447]++
												return &nestedError{"ResourceHeader", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1503
		// _ = "end of CoverTab[11447]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1504
		_go_fuzz_dep_.CoverTab[11448]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1504
		// _ = "end of CoverTab[11448]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1504
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1504
	// _ = "end of CoverTab[11440]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1504
	_go_fuzz_dep_.CoverTab[11441]++
											preLen := len(msg)
											if msg, err = r.pack(msg, b.compression, b.start); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1506
		_go_fuzz_dep_.CoverTab[11449]++
												return &nestedError{"SRVResource body", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1507
		// _ = "end of CoverTab[11449]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1508
		_go_fuzz_dep_.CoverTab[11450]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1508
		// _ = "end of CoverTab[11450]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1508
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1508
	// _ = "end of CoverTab[11441]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1508
	_go_fuzz_dep_.CoverTab[11442]++
											if err := h.fixLen(msg, lenOff, preLen); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1509
		_go_fuzz_dep_.CoverTab[11451]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1510
		// _ = "end of CoverTab[11451]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1511
		_go_fuzz_dep_.CoverTab[11452]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1511
		// _ = "end of CoverTab[11452]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1511
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1511
	// _ = "end of CoverTab[11442]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1511
	_go_fuzz_dep_.CoverTab[11443]++
											if err := b.incrementSectionCount(); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1512
		_go_fuzz_dep_.CoverTab[11453]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1513
		// _ = "end of CoverTab[11453]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1514
		_go_fuzz_dep_.CoverTab[11454]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1514
		// _ = "end of CoverTab[11454]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1514
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1514
	// _ = "end of CoverTab[11443]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1514
	_go_fuzz_dep_.CoverTab[11444]++
											b.msg = msg
											return nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1516
	// _ = "end of CoverTab[11444]"
}

// AResource adds a single AResource.
func (b *Builder) AResource(h ResourceHeader, r AResource) error {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1520
	_go_fuzz_dep_.CoverTab[11455]++
											if err := b.checkResourceSection(); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1521
		_go_fuzz_dep_.CoverTab[11461]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1522
		// _ = "end of CoverTab[11461]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1523
		_go_fuzz_dep_.CoverTab[11462]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1523
		// _ = "end of CoverTab[11462]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1523
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1523
	// _ = "end of CoverTab[11455]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1523
	_go_fuzz_dep_.CoverTab[11456]++
											h.Type = r.realType()
											msg, lenOff, err := h.pack(b.msg, b.compression, b.start)
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1526
		_go_fuzz_dep_.CoverTab[11463]++
												return &nestedError{"ResourceHeader", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1527
		// _ = "end of CoverTab[11463]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1528
		_go_fuzz_dep_.CoverTab[11464]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1528
		// _ = "end of CoverTab[11464]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1528
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1528
	// _ = "end of CoverTab[11456]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1528
	_go_fuzz_dep_.CoverTab[11457]++
											preLen := len(msg)
											if msg, err = r.pack(msg, b.compression, b.start); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1530
		_go_fuzz_dep_.CoverTab[11465]++
												return &nestedError{"AResource body", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1531
		// _ = "end of CoverTab[11465]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1532
		_go_fuzz_dep_.CoverTab[11466]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1532
		// _ = "end of CoverTab[11466]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1532
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1532
	// _ = "end of CoverTab[11457]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1532
	_go_fuzz_dep_.CoverTab[11458]++
											if err := h.fixLen(msg, lenOff, preLen); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1533
		_go_fuzz_dep_.CoverTab[11467]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1534
		// _ = "end of CoverTab[11467]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1535
		_go_fuzz_dep_.CoverTab[11468]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1535
		// _ = "end of CoverTab[11468]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1535
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1535
	// _ = "end of CoverTab[11458]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1535
	_go_fuzz_dep_.CoverTab[11459]++
											if err := b.incrementSectionCount(); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1536
		_go_fuzz_dep_.CoverTab[11469]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1537
		// _ = "end of CoverTab[11469]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1538
		_go_fuzz_dep_.CoverTab[11470]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1538
		// _ = "end of CoverTab[11470]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1538
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1538
	// _ = "end of CoverTab[11459]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1538
	_go_fuzz_dep_.CoverTab[11460]++
											b.msg = msg
											return nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1540
	// _ = "end of CoverTab[11460]"
}

// AAAAResource adds a single AAAAResource.
func (b *Builder) AAAAResource(h ResourceHeader, r AAAAResource) error {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1544
	_go_fuzz_dep_.CoverTab[11471]++
											if err := b.checkResourceSection(); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1545
		_go_fuzz_dep_.CoverTab[11477]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1546
		// _ = "end of CoverTab[11477]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1547
		_go_fuzz_dep_.CoverTab[11478]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1547
		// _ = "end of CoverTab[11478]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1547
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1547
	// _ = "end of CoverTab[11471]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1547
	_go_fuzz_dep_.CoverTab[11472]++
											h.Type = r.realType()
											msg, lenOff, err := h.pack(b.msg, b.compression, b.start)
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1550
		_go_fuzz_dep_.CoverTab[11479]++
												return &nestedError{"ResourceHeader", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1551
		// _ = "end of CoverTab[11479]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1552
		_go_fuzz_dep_.CoverTab[11480]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1552
		// _ = "end of CoverTab[11480]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1552
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1552
	// _ = "end of CoverTab[11472]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1552
	_go_fuzz_dep_.CoverTab[11473]++
											preLen := len(msg)
											if msg, err = r.pack(msg, b.compression, b.start); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1554
		_go_fuzz_dep_.CoverTab[11481]++
												return &nestedError{"AAAAResource body", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1555
		// _ = "end of CoverTab[11481]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1556
		_go_fuzz_dep_.CoverTab[11482]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1556
		// _ = "end of CoverTab[11482]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1556
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1556
	// _ = "end of CoverTab[11473]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1556
	_go_fuzz_dep_.CoverTab[11474]++
											if err := h.fixLen(msg, lenOff, preLen); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1557
		_go_fuzz_dep_.CoverTab[11483]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1558
		// _ = "end of CoverTab[11483]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1559
		_go_fuzz_dep_.CoverTab[11484]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1559
		// _ = "end of CoverTab[11484]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1559
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1559
	// _ = "end of CoverTab[11474]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1559
	_go_fuzz_dep_.CoverTab[11475]++
											if err := b.incrementSectionCount(); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1560
		_go_fuzz_dep_.CoverTab[11485]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1561
		// _ = "end of CoverTab[11485]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1562
		_go_fuzz_dep_.CoverTab[11486]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1562
		// _ = "end of CoverTab[11486]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1562
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1562
	// _ = "end of CoverTab[11475]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1562
	_go_fuzz_dep_.CoverTab[11476]++
											b.msg = msg
											return nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1564
	// _ = "end of CoverTab[11476]"
}

// OPTResource adds a single OPTResource.
func (b *Builder) OPTResource(h ResourceHeader, r OPTResource) error {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1568
	_go_fuzz_dep_.CoverTab[11487]++
											if err := b.checkResourceSection(); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1569
		_go_fuzz_dep_.CoverTab[11493]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1570
		// _ = "end of CoverTab[11493]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1571
		_go_fuzz_dep_.CoverTab[11494]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1571
		// _ = "end of CoverTab[11494]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1571
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1571
	// _ = "end of CoverTab[11487]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1571
	_go_fuzz_dep_.CoverTab[11488]++
											h.Type = r.realType()
											msg, lenOff, err := h.pack(b.msg, b.compression, b.start)
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1574
		_go_fuzz_dep_.CoverTab[11495]++
												return &nestedError{"ResourceHeader", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1575
		// _ = "end of CoverTab[11495]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1576
		_go_fuzz_dep_.CoverTab[11496]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1576
		// _ = "end of CoverTab[11496]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1576
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1576
	// _ = "end of CoverTab[11488]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1576
	_go_fuzz_dep_.CoverTab[11489]++
											preLen := len(msg)
											if msg, err = r.pack(msg, b.compression, b.start); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1578
		_go_fuzz_dep_.CoverTab[11497]++
												return &nestedError{"OPTResource body", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1579
		// _ = "end of CoverTab[11497]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1580
		_go_fuzz_dep_.CoverTab[11498]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1580
		// _ = "end of CoverTab[11498]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1580
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1580
	// _ = "end of CoverTab[11489]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1580
	_go_fuzz_dep_.CoverTab[11490]++
											if err := h.fixLen(msg, lenOff, preLen); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1581
		_go_fuzz_dep_.CoverTab[11499]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1582
		// _ = "end of CoverTab[11499]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1583
		_go_fuzz_dep_.CoverTab[11500]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1583
		// _ = "end of CoverTab[11500]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1583
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1583
	// _ = "end of CoverTab[11490]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1583
	_go_fuzz_dep_.CoverTab[11491]++
											if err := b.incrementSectionCount(); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1584
		_go_fuzz_dep_.CoverTab[11501]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1585
		// _ = "end of CoverTab[11501]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1586
		_go_fuzz_dep_.CoverTab[11502]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1586
		// _ = "end of CoverTab[11502]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1586
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1586
	// _ = "end of CoverTab[11491]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1586
	_go_fuzz_dep_.CoverTab[11492]++
											b.msg = msg
											return nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1588
	// _ = "end of CoverTab[11492]"
}

// UnknownResource adds a single UnknownResource.
func (b *Builder) UnknownResource(h ResourceHeader, r UnknownResource) error {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1592
	_go_fuzz_dep_.CoverTab[11503]++
											if err := b.checkResourceSection(); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1593
		_go_fuzz_dep_.CoverTab[11509]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1594
		// _ = "end of CoverTab[11509]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1595
		_go_fuzz_dep_.CoverTab[11510]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1595
		// _ = "end of CoverTab[11510]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1595
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1595
	// _ = "end of CoverTab[11503]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1595
	_go_fuzz_dep_.CoverTab[11504]++
											h.Type = r.realType()
											msg, lenOff, err := h.pack(b.msg, b.compression, b.start)
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1598
		_go_fuzz_dep_.CoverTab[11511]++
												return &nestedError{"ResourceHeader", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1599
		// _ = "end of CoverTab[11511]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1600
		_go_fuzz_dep_.CoverTab[11512]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1600
		// _ = "end of CoverTab[11512]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1600
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1600
	// _ = "end of CoverTab[11504]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1600
	_go_fuzz_dep_.CoverTab[11505]++
											preLen := len(msg)
											if msg, err = r.pack(msg, b.compression, b.start); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1602
		_go_fuzz_dep_.CoverTab[11513]++
												return &nestedError{"UnknownResource body", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1603
		// _ = "end of CoverTab[11513]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1604
		_go_fuzz_dep_.CoverTab[11514]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1604
		// _ = "end of CoverTab[11514]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1604
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1604
	// _ = "end of CoverTab[11505]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1604
	_go_fuzz_dep_.CoverTab[11506]++
											if err := h.fixLen(msg, lenOff, preLen); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1605
		_go_fuzz_dep_.CoverTab[11515]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1606
		// _ = "end of CoverTab[11515]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1607
		_go_fuzz_dep_.CoverTab[11516]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1607
		// _ = "end of CoverTab[11516]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1607
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1607
	// _ = "end of CoverTab[11506]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1607
	_go_fuzz_dep_.CoverTab[11507]++
											if err := b.incrementSectionCount(); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1608
		_go_fuzz_dep_.CoverTab[11517]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1609
		// _ = "end of CoverTab[11517]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1610
		_go_fuzz_dep_.CoverTab[11518]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1610
		// _ = "end of CoverTab[11518]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1610
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1610
	// _ = "end of CoverTab[11507]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1610
	_go_fuzz_dep_.CoverTab[11508]++
											b.msg = msg
											return nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1612
	// _ = "end of CoverTab[11508]"
}

// Finish ends message building and generates a binary message.
func (b *Builder) Finish() ([]byte, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1616
	_go_fuzz_dep_.CoverTab[11519]++
											if b.section < sectionHeader {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1617
		_go_fuzz_dep_.CoverTab[11521]++
												return nil, ErrNotStarted
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1618
		// _ = "end of CoverTab[11521]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1619
		_go_fuzz_dep_.CoverTab[11522]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1619
		// _ = "end of CoverTab[11522]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1619
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1619
	// _ = "end of CoverTab[11519]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1619
	_go_fuzz_dep_.CoverTab[11520]++
											b.section = sectionDone

											b.header.pack(b.msg[b.start:b.start])
											return b.msg, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1623
	// _ = "end of CoverTab[11520]"
}

// A ResourceHeader is the header of a DNS resource record. There are
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1626
// many types of DNS resource records, but they all share the same header.
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1628
type ResourceHeader struct {
	// Name is the domain name for which this resource record pertains.
	Name	Name

	// Type is the type of DNS resource record.
	//
	// This field will be set automatically during packing.
	Type	Type

	// Class is the class of network to which this DNS resource record
	// pertains.
	Class	Class

	// TTL is the length of time (measured in seconds) which this resource
	// record is valid for (time to live). All Resources in a set should
	// have the same TTL (RFC 2181 Section 5.2).
	TTL	uint32

	// Length is the length of data in the resource record after the header.
	//
	// This field will be set automatically during packing.
	Length	uint16
}

// GoString implements fmt.GoStringer.GoString.
func (h *ResourceHeader) GoString() string {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1653
	_go_fuzz_dep_.CoverTab[11523]++
											return "dnsmessage.ResourceHeader{" +
		"Name: " + h.Name.GoString() + ", " +
		"Type: " + h.Type.GoString() + ", " +
		"Class: " + h.Class.GoString() + ", " +
		"TTL: " + printUint32(h.TTL) + ", " +
		"Length: " + printUint16(h.Length) + "}"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1659
	// _ = "end of CoverTab[11523]"
}

// pack appends the wire format of the ResourceHeader to oldMsg.
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1662
//
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1662
// lenOff is the offset in msg where the Length field was packed.
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1665
func (h *ResourceHeader) pack(oldMsg []byte, compression map[string]int, compressionOff int) (msg []byte, lenOff int, err error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1665
	_go_fuzz_dep_.CoverTab[11524]++
											msg = oldMsg
											if msg, err = h.Name.pack(msg, compression, compressionOff); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1667
		_go_fuzz_dep_.CoverTab[11526]++
												return oldMsg, 0, &nestedError{"Name", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1668
		// _ = "end of CoverTab[11526]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1669
		_go_fuzz_dep_.CoverTab[11527]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1669
		// _ = "end of CoverTab[11527]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1669
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1669
	// _ = "end of CoverTab[11524]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1669
	_go_fuzz_dep_.CoverTab[11525]++
											msg = packType(msg, h.Type)
											msg = packClass(msg, h.Class)
											msg = packUint32(msg, h.TTL)
											lenOff = len(msg)
											msg = packUint16(msg, h.Length)
											return msg, lenOff, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1675
	// _ = "end of CoverTab[11525]"
}

func (h *ResourceHeader) unpack(msg []byte, off int) (int, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1678
	_go_fuzz_dep_.CoverTab[11528]++
											newOff := off
											var err error
											if newOff, err = h.Name.unpack(msg, newOff); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1681
		_go_fuzz_dep_.CoverTab[11534]++
												return off, &nestedError{"Name", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1682
		// _ = "end of CoverTab[11534]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1683
		_go_fuzz_dep_.CoverTab[11535]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1683
		// _ = "end of CoverTab[11535]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1683
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1683
	// _ = "end of CoverTab[11528]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1683
	_go_fuzz_dep_.CoverTab[11529]++
											if h.Type, newOff, err = unpackType(msg, newOff); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1684
		_go_fuzz_dep_.CoverTab[11536]++
												return off, &nestedError{"Type", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1685
		// _ = "end of CoverTab[11536]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1686
		_go_fuzz_dep_.CoverTab[11537]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1686
		// _ = "end of CoverTab[11537]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1686
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1686
	// _ = "end of CoverTab[11529]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1686
	_go_fuzz_dep_.CoverTab[11530]++
											if h.Class, newOff, err = unpackClass(msg, newOff); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1687
		_go_fuzz_dep_.CoverTab[11538]++
												return off, &nestedError{"Class", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1688
		// _ = "end of CoverTab[11538]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1689
		_go_fuzz_dep_.CoverTab[11539]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1689
		// _ = "end of CoverTab[11539]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1689
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1689
	// _ = "end of CoverTab[11530]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1689
	_go_fuzz_dep_.CoverTab[11531]++
											if h.TTL, newOff, err = unpackUint32(msg, newOff); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1690
		_go_fuzz_dep_.CoverTab[11540]++
												return off, &nestedError{"TTL", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1691
		// _ = "end of CoverTab[11540]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1692
		_go_fuzz_dep_.CoverTab[11541]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1692
		// _ = "end of CoverTab[11541]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1692
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1692
	// _ = "end of CoverTab[11531]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1692
	_go_fuzz_dep_.CoverTab[11532]++
											if h.Length, newOff, err = unpackUint16(msg, newOff); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1693
		_go_fuzz_dep_.CoverTab[11542]++
												return off, &nestedError{"Length", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1694
		// _ = "end of CoverTab[11542]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1695
		_go_fuzz_dep_.CoverTab[11543]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1695
		// _ = "end of CoverTab[11543]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1695
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1695
	// _ = "end of CoverTab[11532]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1695
	_go_fuzz_dep_.CoverTab[11533]++
											return newOff, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1696
	// _ = "end of CoverTab[11533]"
}

// fixLen updates a packed ResourceHeader to include the length of the
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1699
// ResourceBody.
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1699
//
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1699
// lenOff is the offset of the ResourceHeader.Length field in msg.
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1699
//
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1699
// preLen is the length that msg was before the ResourceBody was packed.
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1705
func (h *ResourceHeader) fixLen(msg []byte, lenOff int, preLen int) error {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1705
	_go_fuzz_dep_.CoverTab[11544]++
											conLen := len(msg) - preLen
											if conLen > int(^uint16(0)) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1707
		_go_fuzz_dep_.CoverTab[11546]++
												return errResTooLong
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1708
		// _ = "end of CoverTab[11546]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1709
		_go_fuzz_dep_.CoverTab[11547]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1709
		// _ = "end of CoverTab[11547]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1709
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1709
	// _ = "end of CoverTab[11544]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1709
	_go_fuzz_dep_.CoverTab[11545]++

//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1712
	packUint16(msg[lenOff:lenOff], uint16(conLen))
											h.Length = uint16(conLen)

											return nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1715
	// _ = "end of CoverTab[11545]"
}

// EDNS(0) wire constants.
const (
	edns0Version	= 0

	edns0DNSSECOK		= 0x00008000
	ednsVersionMask		= 0x00ff0000
	edns0DNSSECOKMask	= 0x00ff8000
)

// SetEDNS0 configures h for EDNS(0).
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1727
//
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1727
// The provided extRCode must be an extended RCode.
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1730
func (h *ResourceHeader) SetEDNS0(udpPayloadLen int, extRCode RCode, dnssecOK bool) error {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1730
	_go_fuzz_dep_.CoverTab[11548]++
											h.Name = Name{Data: [nameLen]byte{'.'}, Length: 1}
											h.Type = TypeOPT
											h.Class = Class(udpPayloadLen)
											h.TTL = uint32(extRCode) >> 4 << 24
											if dnssecOK {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1735
		_go_fuzz_dep_.CoverTab[11550]++
												h.TTL |= edns0DNSSECOK
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1736
		// _ = "end of CoverTab[11550]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1737
		_go_fuzz_dep_.CoverTab[11551]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1737
		// _ = "end of CoverTab[11551]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1737
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1737
	// _ = "end of CoverTab[11548]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1737
	_go_fuzz_dep_.CoverTab[11549]++
											return nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1738
	// _ = "end of CoverTab[11549]"
}

// DNSSECAllowed reports whether the DNSSEC OK bit is set.
func (h *ResourceHeader) DNSSECAllowed() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1742
	_go_fuzz_dep_.CoverTab[11552]++
											return h.TTL&edns0DNSSECOKMask == edns0DNSSECOK
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1743
	// _ = "end of CoverTab[11552]"
}

// ExtendedRCode returns an extended RCode.
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1746
//
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1746
// The provided rcode must be the RCode in DNS message header.
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1749
func (h *ResourceHeader) ExtendedRCode(rcode RCode) RCode {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1749
	_go_fuzz_dep_.CoverTab[11553]++
											if h.TTL&ednsVersionMask == edns0Version {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1750
		_go_fuzz_dep_.CoverTab[11555]++
												return RCode(h.TTL>>24<<4) | rcode
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1751
		// _ = "end of CoverTab[11555]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1752
		_go_fuzz_dep_.CoverTab[11556]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1752
		// _ = "end of CoverTab[11556]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1752
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1752
	// _ = "end of CoverTab[11553]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1752
	_go_fuzz_dep_.CoverTab[11554]++
											return rcode
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1753
	// _ = "end of CoverTab[11554]"
}

func skipResource(msg []byte, off int) (int, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1756
	_go_fuzz_dep_.CoverTab[11557]++
											newOff, err := skipName(msg, off)
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1758
		_go_fuzz_dep_.CoverTab[11564]++
												return off, &nestedError{"Name", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1759
		// _ = "end of CoverTab[11564]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1760
		_go_fuzz_dep_.CoverTab[11565]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1760
		// _ = "end of CoverTab[11565]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1760
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1760
	// _ = "end of CoverTab[11557]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1760
	_go_fuzz_dep_.CoverTab[11558]++
											if newOff, err = skipType(msg, newOff); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1761
		_go_fuzz_dep_.CoverTab[11566]++
												return off, &nestedError{"Type", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1762
		// _ = "end of CoverTab[11566]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1763
		_go_fuzz_dep_.CoverTab[11567]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1763
		// _ = "end of CoverTab[11567]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1763
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1763
	// _ = "end of CoverTab[11558]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1763
	_go_fuzz_dep_.CoverTab[11559]++
											if newOff, err = skipClass(msg, newOff); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1764
		_go_fuzz_dep_.CoverTab[11568]++
												return off, &nestedError{"Class", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1765
		// _ = "end of CoverTab[11568]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1766
		_go_fuzz_dep_.CoverTab[11569]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1766
		// _ = "end of CoverTab[11569]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1766
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1766
	// _ = "end of CoverTab[11559]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1766
	_go_fuzz_dep_.CoverTab[11560]++
											if newOff, err = skipUint32(msg, newOff); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1767
		_go_fuzz_dep_.CoverTab[11570]++
												return off, &nestedError{"TTL", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1768
		// _ = "end of CoverTab[11570]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1769
		_go_fuzz_dep_.CoverTab[11571]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1769
		// _ = "end of CoverTab[11571]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1769
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1769
	// _ = "end of CoverTab[11560]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1769
	_go_fuzz_dep_.CoverTab[11561]++
											length, newOff, err := unpackUint16(msg, newOff)
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1771
		_go_fuzz_dep_.CoverTab[11572]++
												return off, &nestedError{"Length", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1772
		// _ = "end of CoverTab[11572]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1773
		_go_fuzz_dep_.CoverTab[11573]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1773
		// _ = "end of CoverTab[11573]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1773
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1773
	// _ = "end of CoverTab[11561]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1773
	_go_fuzz_dep_.CoverTab[11562]++
											if newOff += int(length); newOff > len(msg) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1774
		_go_fuzz_dep_.CoverTab[11574]++
												return off, errResourceLen
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1775
		// _ = "end of CoverTab[11574]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1776
		_go_fuzz_dep_.CoverTab[11575]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1776
		// _ = "end of CoverTab[11575]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1776
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1776
	// _ = "end of CoverTab[11562]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1776
	_go_fuzz_dep_.CoverTab[11563]++
											return newOff, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1777
	// _ = "end of CoverTab[11563]"
}

// packUint16 appends the wire format of field to msg.
func packUint16(msg []byte, field uint16) []byte {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1781
	_go_fuzz_dep_.CoverTab[11576]++
											return append(msg, byte(field>>8), byte(field))
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1782
	// _ = "end of CoverTab[11576]"
}

func unpackUint16(msg []byte, off int) (uint16, int, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1785
	_go_fuzz_dep_.CoverTab[11577]++
											if off+uint16Len > len(msg) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1786
		_go_fuzz_dep_.CoverTab[11579]++
												return 0, off, errBaseLen
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1787
		// _ = "end of CoverTab[11579]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1788
		_go_fuzz_dep_.CoverTab[11580]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1788
		// _ = "end of CoverTab[11580]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1788
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1788
	// _ = "end of CoverTab[11577]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1788
	_go_fuzz_dep_.CoverTab[11578]++
											return uint16(msg[off])<<8 | uint16(msg[off+1]), off + uint16Len, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1789
	// _ = "end of CoverTab[11578]"
}

func skipUint16(msg []byte, off int) (int, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1792
	_go_fuzz_dep_.CoverTab[11581]++
											if off+uint16Len > len(msg) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1793
		_go_fuzz_dep_.CoverTab[11583]++
												return off, errBaseLen
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1794
		// _ = "end of CoverTab[11583]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1795
		_go_fuzz_dep_.CoverTab[11584]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1795
		// _ = "end of CoverTab[11584]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1795
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1795
	// _ = "end of CoverTab[11581]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1795
	_go_fuzz_dep_.CoverTab[11582]++
											return off + uint16Len, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1796
	// _ = "end of CoverTab[11582]"
}

// packType appends the wire format of field to msg.
func packType(msg []byte, field Type) []byte {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1800
	_go_fuzz_dep_.CoverTab[11585]++
											return packUint16(msg, uint16(field))
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1801
	// _ = "end of CoverTab[11585]"
}

func unpackType(msg []byte, off int) (Type, int, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1804
	_go_fuzz_dep_.CoverTab[11586]++
											t, o, err := unpackUint16(msg, off)
											return Type(t), o, err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1806
	// _ = "end of CoverTab[11586]"
}

func skipType(msg []byte, off int) (int, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1809
	_go_fuzz_dep_.CoverTab[11587]++
											return skipUint16(msg, off)
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1810
	// _ = "end of CoverTab[11587]"
}

// packClass appends the wire format of field to msg.
func packClass(msg []byte, field Class) []byte {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1814
	_go_fuzz_dep_.CoverTab[11588]++
											return packUint16(msg, uint16(field))
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1815
	// _ = "end of CoverTab[11588]"
}

func unpackClass(msg []byte, off int) (Class, int, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1818
	_go_fuzz_dep_.CoverTab[11589]++
											c, o, err := unpackUint16(msg, off)
											return Class(c), o, err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1820
	// _ = "end of CoverTab[11589]"
}

func skipClass(msg []byte, off int) (int, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1823
	_go_fuzz_dep_.CoverTab[11590]++
											return skipUint16(msg, off)
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1824
	// _ = "end of CoverTab[11590]"
}

// packUint32 appends the wire format of field to msg.
func packUint32(msg []byte, field uint32) []byte {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1828
	_go_fuzz_dep_.CoverTab[11591]++
											return append(
		msg,
		byte(field>>24),
		byte(field>>16),
		byte(field>>8),
		byte(field),
	)
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1835
	// _ = "end of CoverTab[11591]"
}

func unpackUint32(msg []byte, off int) (uint32, int, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1838
	_go_fuzz_dep_.CoverTab[11592]++
											if off+uint32Len > len(msg) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1839
		_go_fuzz_dep_.CoverTab[11594]++
												return 0, off, errBaseLen
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1840
		// _ = "end of CoverTab[11594]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1841
		_go_fuzz_dep_.CoverTab[11595]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1841
		// _ = "end of CoverTab[11595]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1841
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1841
	// _ = "end of CoverTab[11592]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1841
	_go_fuzz_dep_.CoverTab[11593]++
											v := uint32(msg[off])<<24 | uint32(msg[off+1])<<16 | uint32(msg[off+2])<<8 | uint32(msg[off+3])
											return v, off + uint32Len, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1843
	// _ = "end of CoverTab[11593]"
}

func skipUint32(msg []byte, off int) (int, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1846
	_go_fuzz_dep_.CoverTab[11596]++
											if off+uint32Len > len(msg) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1847
		_go_fuzz_dep_.CoverTab[11598]++
												return off, errBaseLen
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1848
		// _ = "end of CoverTab[11598]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1849
		_go_fuzz_dep_.CoverTab[11599]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1849
		// _ = "end of CoverTab[11599]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1849
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1849
	// _ = "end of CoverTab[11596]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1849
	_go_fuzz_dep_.CoverTab[11597]++
											return off + uint32Len, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1850
	// _ = "end of CoverTab[11597]"
}

// packText appends the wire format of field to msg.
func packText(msg []byte, field string) ([]byte, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1854
	_go_fuzz_dep_.CoverTab[11600]++
											l := len(field)
											if l > 255 {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1856
		_go_fuzz_dep_.CoverTab[11602]++
												return nil, errStringTooLong
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1857
		// _ = "end of CoverTab[11602]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1858
		_go_fuzz_dep_.CoverTab[11603]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1858
		// _ = "end of CoverTab[11603]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1858
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1858
	// _ = "end of CoverTab[11600]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1858
	_go_fuzz_dep_.CoverTab[11601]++
											msg = append(msg, byte(l))
											msg = append(msg, field...)

											return msg, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1862
	// _ = "end of CoverTab[11601]"
}

func unpackText(msg []byte, off int) (string, int, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1865
	_go_fuzz_dep_.CoverTab[11604]++
											if off >= len(msg) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1866
		_go_fuzz_dep_.CoverTab[11607]++
												return "", off, errBaseLen
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1867
		// _ = "end of CoverTab[11607]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1868
		_go_fuzz_dep_.CoverTab[11608]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1868
		// _ = "end of CoverTab[11608]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1868
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1868
	// _ = "end of CoverTab[11604]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1868
	_go_fuzz_dep_.CoverTab[11605]++
											beginOff := off + 1
											endOff := beginOff + int(msg[off])
											if endOff > len(msg) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1871
		_go_fuzz_dep_.CoverTab[11609]++
												return "", off, errCalcLen
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1872
		// _ = "end of CoverTab[11609]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1873
		_go_fuzz_dep_.CoverTab[11610]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1873
		// _ = "end of CoverTab[11610]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1873
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1873
	// _ = "end of CoverTab[11605]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1873
	_go_fuzz_dep_.CoverTab[11606]++
											return string(msg[beginOff:endOff]), endOff, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1874
	// _ = "end of CoverTab[11606]"
}

// packBytes appends the wire format of field to msg.
func packBytes(msg []byte, field []byte) []byte {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1878
	_go_fuzz_dep_.CoverTab[11611]++
											return append(msg, field...)
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1879
	// _ = "end of CoverTab[11611]"
}

func unpackBytes(msg []byte, off int, field []byte) (int, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1882
	_go_fuzz_dep_.CoverTab[11612]++
											newOff := off + len(field)
											if newOff > len(msg) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1884
		_go_fuzz_dep_.CoverTab[11614]++
												return off, errBaseLen
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1885
		// _ = "end of CoverTab[11614]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1886
		_go_fuzz_dep_.CoverTab[11615]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1886
		// _ = "end of CoverTab[11615]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1886
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1886
	// _ = "end of CoverTab[11612]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1886
	_go_fuzz_dep_.CoverTab[11613]++
											copy(field, msg[off:newOff])
											return newOff, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1888
	// _ = "end of CoverTab[11613]"
}

const nameLen = 255

// A Name is a non-encoded domain name. It is used instead of strings to avoid
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1893
// allocations.
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1895
type Name struct {
	Data	[nameLen]byte	// 255 bytes
	Length	uint8
}

// NewName creates a new Name from a string.
func NewName(name string) (Name, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1901
	_go_fuzz_dep_.CoverTab[11616]++
											if len(name) > nameLen {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1902
		_go_fuzz_dep_.CoverTab[11618]++
												return Name{}, errCalcLen
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1903
		// _ = "end of CoverTab[11618]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1904
		_go_fuzz_dep_.CoverTab[11619]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1904
		// _ = "end of CoverTab[11619]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1904
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1904
	// _ = "end of CoverTab[11616]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1904
	_go_fuzz_dep_.CoverTab[11617]++
											n := Name{Length: uint8(len(name))}
											copy(n.Data[:], name)
											return n, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1907
	// _ = "end of CoverTab[11617]"
}

// MustNewName creates a new Name from a string and panics on error.
func MustNewName(name string) Name {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1911
	_go_fuzz_dep_.CoverTab[11620]++
											n, err := NewName(name)
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1913
		_go_fuzz_dep_.CoverTab[11622]++
												panic("creating name: " + err.Error())
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1914
		// _ = "end of CoverTab[11622]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1915
		_go_fuzz_dep_.CoverTab[11623]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1915
		// _ = "end of CoverTab[11623]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1915
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1915
	// _ = "end of CoverTab[11620]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1915
	_go_fuzz_dep_.CoverTab[11621]++
											return n
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1916
	// _ = "end of CoverTab[11621]"
}

// String implements fmt.Stringer.String.
func (n Name) String() string {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1920
	_go_fuzz_dep_.CoverTab[11624]++
											return string(n.Data[:n.Length])
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1921
	// _ = "end of CoverTab[11624]"
}

// GoString implements fmt.GoStringer.GoString.
func (n *Name) GoString() string {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1925
	_go_fuzz_dep_.CoverTab[11625]++
											return `dnsmessage.MustNewName("` + printString(n.Data[:n.Length]) + `")`
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1926
	// _ = "end of CoverTab[11625]"
}

// pack appends the wire format of the Name to msg.
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1929
//
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1929
// Domain names are a sequence of counted strings split at the dots. They end
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1929
// with a zero-length string. Compression can be used to reuse domain suffixes.
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1929
//
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1929
// The compression map will be updated with new domain suffixes. If compression
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1929
// is nil, compression will not be used.
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1936
func (n *Name) pack(msg []byte, compression map[string]int, compressionOff int) ([]byte, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1936
	_go_fuzz_dep_.CoverTab[11626]++
											oldMsg := msg

//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1940
	if n.Length == 0 || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1940
		_go_fuzz_dep_.CoverTab[11630]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1940
		return n.Data[n.Length-1] != '.'
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1940
		// _ = "end of CoverTab[11630]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1940
	}() {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1940
		_go_fuzz_dep_.CoverTab[11631]++
												return oldMsg, errNonCanonicalName
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1941
		// _ = "end of CoverTab[11631]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1942
		_go_fuzz_dep_.CoverTab[11632]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1942
		// _ = "end of CoverTab[11632]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1942
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1942
	// _ = "end of CoverTab[11626]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1942
	_go_fuzz_dep_.CoverTab[11627]++

//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1945
	if n.Data[0] == '.' && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1945
		_go_fuzz_dep_.CoverTab[11633]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1945
		return n.Length == 1
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1945
		// _ = "end of CoverTab[11633]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1945
	}() {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1945
		_go_fuzz_dep_.CoverTab[11634]++
												return append(msg, 0), nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1946
		// _ = "end of CoverTab[11634]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1947
		_go_fuzz_dep_.CoverTab[11635]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1947
		// _ = "end of CoverTab[11635]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1947
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1947
	// _ = "end of CoverTab[11627]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1947
	_go_fuzz_dep_.CoverTab[11628]++

//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1950
	for i, begin := 0, 0; i < int(n.Length); i++ {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1950
		_go_fuzz_dep_.CoverTab[11636]++

												if n.Data[i] == '.' {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1952
			_go_fuzz_dep_.CoverTab[11638]++

//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1956
			if i-begin >= 1<<6 {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1956
				_go_fuzz_dep_.CoverTab[11642]++
														return oldMsg, errSegTooLong
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1957
				// _ = "end of CoverTab[11642]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1958
				_go_fuzz_dep_.CoverTab[11643]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1958
				// _ = "end of CoverTab[11643]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1958
			}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1958
			// _ = "end of CoverTab[11638]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1958
			_go_fuzz_dep_.CoverTab[11639]++

//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1961
			if i-begin == 0 {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1961
				_go_fuzz_dep_.CoverTab[11644]++
														return oldMsg, errZeroSegLen
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1962
				// _ = "end of CoverTab[11644]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1963
				_go_fuzz_dep_.CoverTab[11645]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1963
				// _ = "end of CoverTab[11645]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1963
			}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1963
			// _ = "end of CoverTab[11639]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1963
			_go_fuzz_dep_.CoverTab[11640]++

													msg = append(msg, byte(i-begin))

													for j := begin; j < i; j++ {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1967
				_go_fuzz_dep_.CoverTab[11646]++
														msg = append(msg, n.Data[j])
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1968
				// _ = "end of CoverTab[11646]"
			}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1969
			// _ = "end of CoverTab[11640]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1969
			_go_fuzz_dep_.CoverTab[11641]++

													begin = i + 1
													continue
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1972
			// _ = "end of CoverTab[11641]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1973
			_go_fuzz_dep_.CoverTab[11647]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1973
			// _ = "end of CoverTab[11647]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1973
		}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1973
		// _ = "end of CoverTab[11636]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1973
		_go_fuzz_dep_.CoverTab[11637]++

//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1978
		if (i == 0 || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1978
			_go_fuzz_dep_.CoverTab[11648]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1978
			return n.Data[i-1] == '.'
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1978
			// _ = "end of CoverTab[11648]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1978
		}()) && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1978
			_go_fuzz_dep_.CoverTab[11649]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1978
			return compression != nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1978
			// _ = "end of CoverTab[11649]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1978
		}() {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1978
			_go_fuzz_dep_.CoverTab[11650]++
													if ptr, ok := compression[string(n.Data[i:])]; ok {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1979
				_go_fuzz_dep_.CoverTab[11652]++

//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1982
				return append(msg, byte(ptr>>8|0xC0), byte(ptr)), nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1982
				// _ = "end of CoverTab[11652]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1983
				_go_fuzz_dep_.CoverTab[11653]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1983
				// _ = "end of CoverTab[11653]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1983
			}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1983
			// _ = "end of CoverTab[11650]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1983
			_go_fuzz_dep_.CoverTab[11651]++

//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1987
			if len(msg) <= int(^uint16(0)>>2) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1987
				_go_fuzz_dep_.CoverTab[11654]++
														compression[string(n.Data[i:])] = len(msg) - compressionOff
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1988
				// _ = "end of CoverTab[11654]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1989
				_go_fuzz_dep_.CoverTab[11655]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1989
				// _ = "end of CoverTab[11655]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1989
			}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1989
			// _ = "end of CoverTab[11651]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1990
			_go_fuzz_dep_.CoverTab[11656]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1990
			// _ = "end of CoverTab[11656]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1990
		}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1990
		// _ = "end of CoverTab[11637]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1991
	// _ = "end of CoverTab[11628]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1991
	_go_fuzz_dep_.CoverTab[11629]++
											return append(msg, 0), nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1992
	// _ = "end of CoverTab[11629]"
}

// unpack unpacks a domain name.
func (n *Name) unpack(msg []byte, off int) (int, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1996
	_go_fuzz_dep_.CoverTab[11657]++
											return n.unpackCompressed(msg, off, true)
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1997
	// _ = "end of CoverTab[11657]"
}

func (n *Name) unpackCompressed(msg []byte, off int, allowCompression bool) (int, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2000
	_go_fuzz_dep_.CoverTab[11658]++

											currOff := off

//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2007
	newOff := off

											// ptr is the number of pointers followed.
											var ptr int

//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2013
	name := n.Data[:0]

Loop:
	for {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2016
		_go_fuzz_dep_.CoverTab[11663]++
												if currOff >= len(msg) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2017
			_go_fuzz_dep_.CoverTab[11665]++
													return off, errBaseLen
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2018
			// _ = "end of CoverTab[11665]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2019
			_go_fuzz_dep_.CoverTab[11666]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2019
			// _ = "end of CoverTab[11666]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2019
		}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2019
		// _ = "end of CoverTab[11663]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2019
		_go_fuzz_dep_.CoverTab[11664]++
												c := int(msg[currOff])
												currOff++
												switch c & 0xC0 {
		case 0x00:
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2023
			_go_fuzz_dep_.CoverTab[11667]++
													if c == 0x00 {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2024
				_go_fuzz_dep_.CoverTab[11676]++

														break Loop
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2026
				// _ = "end of CoverTab[11676]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2027
				_go_fuzz_dep_.CoverTab[11677]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2027
				// _ = "end of CoverTab[11677]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2027
			}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2027
			// _ = "end of CoverTab[11667]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2027
			_go_fuzz_dep_.CoverTab[11668]++
													endOff := currOff + c
													if endOff > len(msg) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2029
				_go_fuzz_dep_.CoverTab[11678]++
														return off, errCalcLen
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2030
				// _ = "end of CoverTab[11678]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2031
				_go_fuzz_dep_.CoverTab[11679]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2031
				// _ = "end of CoverTab[11679]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2031
			}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2031
			// _ = "end of CoverTab[11668]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2031
			_go_fuzz_dep_.CoverTab[11669]++
													name = append(name, msg[currOff:endOff]...)
													name = append(name, '.')
													currOff = endOff
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2034
			// _ = "end of CoverTab[11669]"
		case 0xC0:
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2035
			_go_fuzz_dep_.CoverTab[11670]++
													if !allowCompression {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2036
				_go_fuzz_dep_.CoverTab[11680]++
														return off, errCompressedSRV
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2037
				// _ = "end of CoverTab[11680]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2038
				_go_fuzz_dep_.CoverTab[11681]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2038
				// _ = "end of CoverTab[11681]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2038
			}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2038
			// _ = "end of CoverTab[11670]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2038
			_go_fuzz_dep_.CoverTab[11671]++
													if currOff >= len(msg) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2039
				_go_fuzz_dep_.CoverTab[11682]++
														return off, errInvalidPtr
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2040
				// _ = "end of CoverTab[11682]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2041
				_go_fuzz_dep_.CoverTab[11683]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2041
				// _ = "end of CoverTab[11683]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2041
			}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2041
			// _ = "end of CoverTab[11671]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2041
			_go_fuzz_dep_.CoverTab[11672]++
													c1 := msg[currOff]
													currOff++
													if ptr == 0 {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2044
				_go_fuzz_dep_.CoverTab[11684]++
														newOff = currOff
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2045
				// _ = "end of CoverTab[11684]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2046
				_go_fuzz_dep_.CoverTab[11685]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2046
				// _ = "end of CoverTab[11685]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2046
			}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2046
			// _ = "end of CoverTab[11672]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2046
			_go_fuzz_dep_.CoverTab[11673]++

													if ptr++; ptr > 10 {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2048
				_go_fuzz_dep_.CoverTab[11686]++
														return off, errTooManyPtr
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2049
				// _ = "end of CoverTab[11686]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2050
				_go_fuzz_dep_.CoverTab[11687]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2050
				// _ = "end of CoverTab[11687]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2050
			}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2050
			// _ = "end of CoverTab[11673]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2050
			_go_fuzz_dep_.CoverTab[11674]++
													currOff = (c^0xC0)<<8 | int(c1)
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2051
			// _ = "end of CoverTab[11674]"
		default:
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2052
			_go_fuzz_dep_.CoverTab[11675]++

													return off, errReserved
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2054
			// _ = "end of CoverTab[11675]"
		}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2055
		// _ = "end of CoverTab[11664]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2056
	// _ = "end of CoverTab[11658]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2056
	_go_fuzz_dep_.CoverTab[11659]++
											if len(name) == 0 {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2057
		_go_fuzz_dep_.CoverTab[11688]++
												name = append(name, '.')
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2058
		// _ = "end of CoverTab[11688]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2059
		_go_fuzz_dep_.CoverTab[11689]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2059
		// _ = "end of CoverTab[11689]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2059
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2059
	// _ = "end of CoverTab[11659]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2059
	_go_fuzz_dep_.CoverTab[11660]++
											if len(name) > len(n.Data) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2060
		_go_fuzz_dep_.CoverTab[11690]++
												return off, errCalcLen
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2061
		// _ = "end of CoverTab[11690]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2062
		_go_fuzz_dep_.CoverTab[11691]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2062
		// _ = "end of CoverTab[11691]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2062
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2062
	// _ = "end of CoverTab[11660]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2062
	_go_fuzz_dep_.CoverTab[11661]++
											n.Length = uint8(len(name))
											if ptr == 0 {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2064
		_go_fuzz_dep_.CoverTab[11692]++
												newOff = currOff
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2065
		// _ = "end of CoverTab[11692]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2066
		_go_fuzz_dep_.CoverTab[11693]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2066
		// _ = "end of CoverTab[11693]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2066
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2066
	// _ = "end of CoverTab[11661]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2066
	_go_fuzz_dep_.CoverTab[11662]++
											return newOff, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2067
	// _ = "end of CoverTab[11662]"
}

func skipName(msg []byte, off int) (int, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2070
	_go_fuzz_dep_.CoverTab[11694]++

//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2074
	newOff := off

Loop:
	for {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2077
		_go_fuzz_dep_.CoverTab[11696]++
												if newOff >= len(msg) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2078
			_go_fuzz_dep_.CoverTab[11698]++
													return off, errBaseLen
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2079
			// _ = "end of CoverTab[11698]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2080
			_go_fuzz_dep_.CoverTab[11699]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2080
			// _ = "end of CoverTab[11699]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2080
		}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2080
		// _ = "end of CoverTab[11696]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2080
		_go_fuzz_dep_.CoverTab[11697]++
												c := int(msg[newOff])
												newOff++
												switch c & 0xC0 {
		case 0x00:
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2084
			_go_fuzz_dep_.CoverTab[11700]++
													if c == 0x00 {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2085
				_go_fuzz_dep_.CoverTab[11704]++

														break Loop
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2087
				// _ = "end of CoverTab[11704]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2088
				_go_fuzz_dep_.CoverTab[11705]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2088
				// _ = "end of CoverTab[11705]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2088
			}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2088
			// _ = "end of CoverTab[11700]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2088
			_go_fuzz_dep_.CoverTab[11701]++

													newOff += c
													if newOff > len(msg) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2091
				_go_fuzz_dep_.CoverTab[11706]++
														return off, errCalcLen
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2092
				// _ = "end of CoverTab[11706]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2093
				_go_fuzz_dep_.CoverTab[11707]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2093
				// _ = "end of CoverTab[11707]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2093
			}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2093
			// _ = "end of CoverTab[11701]"
		case 0xC0:
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2094
			_go_fuzz_dep_.CoverTab[11702]++

//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2098
			newOff++

//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2101
			break Loop
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2101
			// _ = "end of CoverTab[11702]"
		default:
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2102
			_go_fuzz_dep_.CoverTab[11703]++

													return off, errReserved
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2104
			// _ = "end of CoverTab[11703]"
		}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2105
		// _ = "end of CoverTab[11697]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2106
	// _ = "end of CoverTab[11694]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2106
	_go_fuzz_dep_.CoverTab[11695]++

											return newOff, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2108
	// _ = "end of CoverTab[11695]"
}

// A Question is a DNS query.
type Question struct {
	Name	Name
	Type	Type
	Class	Class
}

// pack appends the wire format of the Question to msg.
func (q *Question) pack(msg []byte, compression map[string]int, compressionOff int) ([]byte, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2119
	_go_fuzz_dep_.CoverTab[11708]++
											msg, err := q.Name.pack(msg, compression, compressionOff)
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2121
		_go_fuzz_dep_.CoverTab[11710]++
												return msg, &nestedError{"Name", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2122
		// _ = "end of CoverTab[11710]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2123
		_go_fuzz_dep_.CoverTab[11711]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2123
		// _ = "end of CoverTab[11711]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2123
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2123
	// _ = "end of CoverTab[11708]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2123
	_go_fuzz_dep_.CoverTab[11709]++
											msg = packType(msg, q.Type)
											return packClass(msg, q.Class), nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2125
	// _ = "end of CoverTab[11709]"
}

// GoString implements fmt.GoStringer.GoString.
func (q *Question) GoString() string {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2129
	_go_fuzz_dep_.CoverTab[11712]++
											return "dnsmessage.Question{" +
		"Name: " + q.Name.GoString() + ", " +
		"Type: " + q.Type.GoString() + ", " +
		"Class: " + q.Class.GoString() + "}"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2133
	// _ = "end of CoverTab[11712]"
}

func unpackResourceBody(msg []byte, off int, hdr ResourceHeader) (ResourceBody, int, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2136
	_go_fuzz_dep_.CoverTab[11713]++
											var (
		r	ResourceBody
		err	error
		name	string
	)
	switch hdr.Type {
	case TypeA:
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2143
		_go_fuzz_dep_.CoverTab[11716]++
												var rb AResource
												rb, err = unpackAResource(msg, off)
												r = &rb
												name = "A"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2147
		// _ = "end of CoverTab[11716]"
	case TypeNS:
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2148
		_go_fuzz_dep_.CoverTab[11717]++
												var rb NSResource
												rb, err = unpackNSResource(msg, off)
												r = &rb
												name = "NS"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2152
		// _ = "end of CoverTab[11717]"
	case TypeCNAME:
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2153
		_go_fuzz_dep_.CoverTab[11718]++
												var rb CNAMEResource
												rb, err = unpackCNAMEResource(msg, off)
												r = &rb
												name = "CNAME"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2157
		// _ = "end of CoverTab[11718]"
	case TypeSOA:
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2158
		_go_fuzz_dep_.CoverTab[11719]++
												var rb SOAResource
												rb, err = unpackSOAResource(msg, off)
												r = &rb
												name = "SOA"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2162
		// _ = "end of CoverTab[11719]"
	case TypePTR:
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2163
		_go_fuzz_dep_.CoverTab[11720]++
												var rb PTRResource
												rb, err = unpackPTRResource(msg, off)
												r = &rb
												name = "PTR"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2167
		// _ = "end of CoverTab[11720]"
	case TypeMX:
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2168
		_go_fuzz_dep_.CoverTab[11721]++
												var rb MXResource
												rb, err = unpackMXResource(msg, off)
												r = &rb
												name = "MX"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2172
		// _ = "end of CoverTab[11721]"
	case TypeTXT:
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2173
		_go_fuzz_dep_.CoverTab[11722]++
												var rb TXTResource
												rb, err = unpackTXTResource(msg, off, hdr.Length)
												r = &rb
												name = "TXT"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2177
		// _ = "end of CoverTab[11722]"
	case TypeAAAA:
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2178
		_go_fuzz_dep_.CoverTab[11723]++
												var rb AAAAResource
												rb, err = unpackAAAAResource(msg, off)
												r = &rb
												name = "AAAA"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2182
		// _ = "end of CoverTab[11723]"
	case TypeSRV:
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2183
		_go_fuzz_dep_.CoverTab[11724]++
												var rb SRVResource
												rb, err = unpackSRVResource(msg, off)
												r = &rb
												name = "SRV"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2187
		// _ = "end of CoverTab[11724]"
	case TypeOPT:
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2188
		_go_fuzz_dep_.CoverTab[11725]++
												var rb OPTResource
												rb, err = unpackOPTResource(msg, off, hdr.Length)
												r = &rb
												name = "OPT"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2192
		// _ = "end of CoverTab[11725]"
	default:
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2193
		_go_fuzz_dep_.CoverTab[11726]++
												var rb UnknownResource
												rb, err = unpackUnknownResource(hdr.Type, msg, off, hdr.Length)
												r = &rb
												name = "Unknown"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2197
		// _ = "end of CoverTab[11726]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2198
	// _ = "end of CoverTab[11713]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2198
	_go_fuzz_dep_.CoverTab[11714]++
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2199
		_go_fuzz_dep_.CoverTab[11727]++
												return nil, off, &nestedError{name + " record", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2200
		// _ = "end of CoverTab[11727]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2201
		_go_fuzz_dep_.CoverTab[11728]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2201
		// _ = "end of CoverTab[11728]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2201
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2201
	// _ = "end of CoverTab[11714]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2201
	_go_fuzz_dep_.CoverTab[11715]++
											return r, off + int(hdr.Length), nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2202
	// _ = "end of CoverTab[11715]"
}

// A CNAMEResource is a CNAME Resource record.
type CNAMEResource struct {
	CNAME Name
}

func (r *CNAMEResource) realType() Type {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2210
	_go_fuzz_dep_.CoverTab[11729]++
											return TypeCNAME
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2211
	// _ = "end of CoverTab[11729]"
}

// pack appends the wire format of the CNAMEResource to msg.
func (r *CNAMEResource) pack(msg []byte, compression map[string]int, compressionOff int) ([]byte, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2215
	_go_fuzz_dep_.CoverTab[11730]++
											return r.CNAME.pack(msg, compression, compressionOff)
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2216
	// _ = "end of CoverTab[11730]"
}

// GoString implements fmt.GoStringer.GoString.
func (r *CNAMEResource) GoString() string {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2220
	_go_fuzz_dep_.CoverTab[11731]++
											return "dnsmessage.CNAMEResource{CNAME: " + r.CNAME.GoString() + "}"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2221
	// _ = "end of CoverTab[11731]"
}

func unpackCNAMEResource(msg []byte, off int) (CNAMEResource, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2224
	_go_fuzz_dep_.CoverTab[11732]++
											var cname Name
											if _, err := cname.unpack(msg, off); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2226
		_go_fuzz_dep_.CoverTab[11734]++
												return CNAMEResource{}, err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2227
		// _ = "end of CoverTab[11734]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2228
		_go_fuzz_dep_.CoverTab[11735]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2228
		// _ = "end of CoverTab[11735]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2228
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2228
	// _ = "end of CoverTab[11732]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2228
	_go_fuzz_dep_.CoverTab[11733]++
											return CNAMEResource{cname}, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2229
	// _ = "end of CoverTab[11733]"
}

// An MXResource is an MX Resource record.
type MXResource struct {
	Pref	uint16
	MX	Name
}

func (r *MXResource) realType() Type {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2238
	_go_fuzz_dep_.CoverTab[11736]++
											return TypeMX
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2239
	// _ = "end of CoverTab[11736]"
}

// pack appends the wire format of the MXResource to msg.
func (r *MXResource) pack(msg []byte, compression map[string]int, compressionOff int) ([]byte, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2243
	_go_fuzz_dep_.CoverTab[11737]++
											oldMsg := msg
											msg = packUint16(msg, r.Pref)
											msg, err := r.MX.pack(msg, compression, compressionOff)
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2247
		_go_fuzz_dep_.CoverTab[11739]++
												return oldMsg, &nestedError{"MXResource.MX", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2248
		// _ = "end of CoverTab[11739]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2249
		_go_fuzz_dep_.CoverTab[11740]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2249
		// _ = "end of CoverTab[11740]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2249
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2249
	// _ = "end of CoverTab[11737]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2249
	_go_fuzz_dep_.CoverTab[11738]++
											return msg, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2250
	// _ = "end of CoverTab[11738]"
}

// GoString implements fmt.GoStringer.GoString.
func (r *MXResource) GoString() string {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2254
	_go_fuzz_dep_.CoverTab[11741]++
											return "dnsmessage.MXResource{" +
		"Pref: " + printUint16(r.Pref) + ", " +
		"MX: " + r.MX.GoString() + "}"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2257
	// _ = "end of CoverTab[11741]"
}

func unpackMXResource(msg []byte, off int) (MXResource, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2260
	_go_fuzz_dep_.CoverTab[11742]++
											pref, off, err := unpackUint16(msg, off)
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2262
		_go_fuzz_dep_.CoverTab[11745]++
												return MXResource{}, &nestedError{"Pref", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2263
		// _ = "end of CoverTab[11745]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2264
		_go_fuzz_dep_.CoverTab[11746]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2264
		// _ = "end of CoverTab[11746]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2264
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2264
	// _ = "end of CoverTab[11742]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2264
	_go_fuzz_dep_.CoverTab[11743]++
											var mx Name
											if _, err := mx.unpack(msg, off); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2266
		_go_fuzz_dep_.CoverTab[11747]++
												return MXResource{}, &nestedError{"MX", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2267
		// _ = "end of CoverTab[11747]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2268
		_go_fuzz_dep_.CoverTab[11748]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2268
		// _ = "end of CoverTab[11748]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2268
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2268
	// _ = "end of CoverTab[11743]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2268
	_go_fuzz_dep_.CoverTab[11744]++
											return MXResource{pref, mx}, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2269
	// _ = "end of CoverTab[11744]"
}

// An NSResource is an NS Resource record.
type NSResource struct {
	NS Name
}

func (r *NSResource) realType() Type {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2277
	_go_fuzz_dep_.CoverTab[11749]++
											return TypeNS
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2278
	// _ = "end of CoverTab[11749]"
}

// pack appends the wire format of the NSResource to msg.
func (r *NSResource) pack(msg []byte, compression map[string]int, compressionOff int) ([]byte, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2282
	_go_fuzz_dep_.CoverTab[11750]++
											return r.NS.pack(msg, compression, compressionOff)
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2283
	// _ = "end of CoverTab[11750]"
}

// GoString implements fmt.GoStringer.GoString.
func (r *NSResource) GoString() string {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2287
	_go_fuzz_dep_.CoverTab[11751]++
											return "dnsmessage.NSResource{NS: " + r.NS.GoString() + "}"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2288
	// _ = "end of CoverTab[11751]"
}

func unpackNSResource(msg []byte, off int) (NSResource, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2291
	_go_fuzz_dep_.CoverTab[11752]++
											var ns Name
											if _, err := ns.unpack(msg, off); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2293
		_go_fuzz_dep_.CoverTab[11754]++
												return NSResource{}, err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2294
		// _ = "end of CoverTab[11754]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2295
		_go_fuzz_dep_.CoverTab[11755]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2295
		// _ = "end of CoverTab[11755]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2295
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2295
	// _ = "end of CoverTab[11752]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2295
	_go_fuzz_dep_.CoverTab[11753]++
											return NSResource{ns}, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2296
	// _ = "end of CoverTab[11753]"
}

// A PTRResource is a PTR Resource record.
type PTRResource struct {
	PTR Name
}

func (r *PTRResource) realType() Type {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2304
	_go_fuzz_dep_.CoverTab[11756]++
											return TypePTR
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2305
	// _ = "end of CoverTab[11756]"
}

// pack appends the wire format of the PTRResource to msg.
func (r *PTRResource) pack(msg []byte, compression map[string]int, compressionOff int) ([]byte, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2309
	_go_fuzz_dep_.CoverTab[11757]++
											return r.PTR.pack(msg, compression, compressionOff)
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2310
	// _ = "end of CoverTab[11757]"
}

// GoString implements fmt.GoStringer.GoString.
func (r *PTRResource) GoString() string {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2314
	_go_fuzz_dep_.CoverTab[11758]++
											return "dnsmessage.PTRResource{PTR: " + r.PTR.GoString() + "}"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2315
	// _ = "end of CoverTab[11758]"
}

func unpackPTRResource(msg []byte, off int) (PTRResource, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2318
	_go_fuzz_dep_.CoverTab[11759]++
											var ptr Name
											if _, err := ptr.unpack(msg, off); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2320
		_go_fuzz_dep_.CoverTab[11761]++
												return PTRResource{}, err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2321
		// _ = "end of CoverTab[11761]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2322
		_go_fuzz_dep_.CoverTab[11762]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2322
		// _ = "end of CoverTab[11762]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2322
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2322
	// _ = "end of CoverTab[11759]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2322
	_go_fuzz_dep_.CoverTab[11760]++
											return PTRResource{ptr}, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2323
	// _ = "end of CoverTab[11760]"
}

// An SOAResource is an SOA Resource record.
type SOAResource struct {
	NS	Name
	MBox	Name
	Serial	uint32
	Refresh	uint32
	Retry	uint32
	Expire	uint32

	// MinTTL the is the default TTL of Resources records which did not
	// contain a TTL value and the TTL of negative responses. (RFC 2308
	// Section 4)
	MinTTL	uint32
}

func (r *SOAResource) realType() Type {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2341
	_go_fuzz_dep_.CoverTab[11763]++
											return TypeSOA
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2342
	// _ = "end of CoverTab[11763]"
}

// pack appends the wire format of the SOAResource to msg.
func (r *SOAResource) pack(msg []byte, compression map[string]int, compressionOff int) ([]byte, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2346
	_go_fuzz_dep_.CoverTab[11764]++
											oldMsg := msg
											msg, err := r.NS.pack(msg, compression, compressionOff)
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2349
		_go_fuzz_dep_.CoverTab[11767]++
												return oldMsg, &nestedError{"SOAResource.NS", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2350
		// _ = "end of CoverTab[11767]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2351
		_go_fuzz_dep_.CoverTab[11768]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2351
		// _ = "end of CoverTab[11768]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2351
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2351
	// _ = "end of CoverTab[11764]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2351
	_go_fuzz_dep_.CoverTab[11765]++
											msg, err = r.MBox.pack(msg, compression, compressionOff)
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2353
		_go_fuzz_dep_.CoverTab[11769]++
												return oldMsg, &nestedError{"SOAResource.MBox", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2354
		// _ = "end of CoverTab[11769]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2355
		_go_fuzz_dep_.CoverTab[11770]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2355
		// _ = "end of CoverTab[11770]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2355
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2355
	// _ = "end of CoverTab[11765]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2355
	_go_fuzz_dep_.CoverTab[11766]++
											msg = packUint32(msg, r.Serial)
											msg = packUint32(msg, r.Refresh)
											msg = packUint32(msg, r.Retry)
											msg = packUint32(msg, r.Expire)
											return packUint32(msg, r.MinTTL), nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2360
	// _ = "end of CoverTab[11766]"
}

// GoString implements fmt.GoStringer.GoString.
func (r *SOAResource) GoString() string {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2364
	_go_fuzz_dep_.CoverTab[11771]++
											return "dnsmessage.SOAResource{" +
		"NS: " + r.NS.GoString() + ", " +
		"MBox: " + r.MBox.GoString() + ", " +
		"Serial: " + printUint32(r.Serial) + ", " +
		"Refresh: " + printUint32(r.Refresh) + ", " +
		"Retry: " + printUint32(r.Retry) + ", " +
		"Expire: " + printUint32(r.Expire) + ", " +
		"MinTTL: " + printUint32(r.MinTTL) + "}"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2372
	// _ = "end of CoverTab[11771]"
}

func unpackSOAResource(msg []byte, off int) (SOAResource, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2375
	_go_fuzz_dep_.CoverTab[11772]++
											var ns Name
											off, err := ns.unpack(msg, off)
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2378
		_go_fuzz_dep_.CoverTab[11780]++
												return SOAResource{}, &nestedError{"NS", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2379
		// _ = "end of CoverTab[11780]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2380
		_go_fuzz_dep_.CoverTab[11781]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2380
		// _ = "end of CoverTab[11781]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2380
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2380
	// _ = "end of CoverTab[11772]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2380
	_go_fuzz_dep_.CoverTab[11773]++
											var mbox Name
											if off, err = mbox.unpack(msg, off); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2382
		_go_fuzz_dep_.CoverTab[11782]++
												return SOAResource{}, &nestedError{"MBox", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2383
		// _ = "end of CoverTab[11782]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2384
		_go_fuzz_dep_.CoverTab[11783]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2384
		// _ = "end of CoverTab[11783]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2384
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2384
	// _ = "end of CoverTab[11773]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2384
	_go_fuzz_dep_.CoverTab[11774]++
											serial, off, err := unpackUint32(msg, off)
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2386
		_go_fuzz_dep_.CoverTab[11784]++
												return SOAResource{}, &nestedError{"Serial", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2387
		// _ = "end of CoverTab[11784]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2388
		_go_fuzz_dep_.CoverTab[11785]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2388
		// _ = "end of CoverTab[11785]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2388
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2388
	// _ = "end of CoverTab[11774]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2388
	_go_fuzz_dep_.CoverTab[11775]++
											refresh, off, err := unpackUint32(msg, off)
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2390
		_go_fuzz_dep_.CoverTab[11786]++
												return SOAResource{}, &nestedError{"Refresh", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2391
		// _ = "end of CoverTab[11786]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2392
		_go_fuzz_dep_.CoverTab[11787]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2392
		// _ = "end of CoverTab[11787]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2392
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2392
	// _ = "end of CoverTab[11775]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2392
	_go_fuzz_dep_.CoverTab[11776]++
											retry, off, err := unpackUint32(msg, off)
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2394
		_go_fuzz_dep_.CoverTab[11788]++
												return SOAResource{}, &nestedError{"Retry", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2395
		// _ = "end of CoverTab[11788]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2396
		_go_fuzz_dep_.CoverTab[11789]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2396
		// _ = "end of CoverTab[11789]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2396
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2396
	// _ = "end of CoverTab[11776]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2396
	_go_fuzz_dep_.CoverTab[11777]++
											expire, off, err := unpackUint32(msg, off)
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2398
		_go_fuzz_dep_.CoverTab[11790]++
												return SOAResource{}, &nestedError{"Expire", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2399
		// _ = "end of CoverTab[11790]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2400
		_go_fuzz_dep_.CoverTab[11791]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2400
		// _ = "end of CoverTab[11791]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2400
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2400
	// _ = "end of CoverTab[11777]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2400
	_go_fuzz_dep_.CoverTab[11778]++
											minTTL, _, err := unpackUint32(msg, off)
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2402
		_go_fuzz_dep_.CoverTab[11792]++
												return SOAResource{}, &nestedError{"MinTTL", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2403
		// _ = "end of CoverTab[11792]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2404
		_go_fuzz_dep_.CoverTab[11793]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2404
		// _ = "end of CoverTab[11793]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2404
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2404
	// _ = "end of CoverTab[11778]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2404
	_go_fuzz_dep_.CoverTab[11779]++
											return SOAResource{ns, mbox, serial, refresh, retry, expire, minTTL}, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2405
	// _ = "end of CoverTab[11779]"
}

// A TXTResource is a TXT Resource record.
type TXTResource struct {
	TXT []string
}

func (r *TXTResource) realType() Type {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2413
	_go_fuzz_dep_.CoverTab[11794]++
											return TypeTXT
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2414
	// _ = "end of CoverTab[11794]"
}

// pack appends the wire format of the TXTResource to msg.
func (r *TXTResource) pack(msg []byte, compression map[string]int, compressionOff int) ([]byte, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2418
	_go_fuzz_dep_.CoverTab[11795]++
											oldMsg := msg
											for _, s := range r.TXT {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2420
		_go_fuzz_dep_.CoverTab[11797]++
												var err error
												msg, err = packText(msg, s)
												if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2423
			_go_fuzz_dep_.CoverTab[11798]++
													return oldMsg, err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2424
			// _ = "end of CoverTab[11798]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2425
			_go_fuzz_dep_.CoverTab[11799]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2425
			// _ = "end of CoverTab[11799]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2425
		}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2425
		// _ = "end of CoverTab[11797]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2426
	// _ = "end of CoverTab[11795]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2426
	_go_fuzz_dep_.CoverTab[11796]++
											return msg, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2427
	// _ = "end of CoverTab[11796]"
}

// GoString implements fmt.GoStringer.GoString.
func (r *TXTResource) GoString() string {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2431
	_go_fuzz_dep_.CoverTab[11800]++
											s := "dnsmessage.TXTResource{TXT: []string{"
											if len(r.TXT) == 0 {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2433
		_go_fuzz_dep_.CoverTab[11803]++
												return s + "}}"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2434
		// _ = "end of CoverTab[11803]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2435
		_go_fuzz_dep_.CoverTab[11804]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2435
		// _ = "end of CoverTab[11804]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2435
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2435
	// _ = "end of CoverTab[11800]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2435
	_go_fuzz_dep_.CoverTab[11801]++
											s += `"` + printString([]byte(r.TXT[0]))
											for _, t := range r.TXT[1:] {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2437
		_go_fuzz_dep_.CoverTab[11805]++
												s += `", "` + printString([]byte(t))
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2438
		// _ = "end of CoverTab[11805]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2439
	// _ = "end of CoverTab[11801]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2439
	_go_fuzz_dep_.CoverTab[11802]++
											return s + `"}}`
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2440
	// _ = "end of CoverTab[11802]"
}

func unpackTXTResource(msg []byte, off int, length uint16) (TXTResource, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2443
	_go_fuzz_dep_.CoverTab[11806]++
											txts := make([]string, 0, 1)
											for n := uint16(0); n < length; {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2445
		_go_fuzz_dep_.CoverTab[11808]++
												var t string
												var err error
												if t, off, err = unpackText(msg, off); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2448
			_go_fuzz_dep_.CoverTab[11811]++
													return TXTResource{}, &nestedError{"text", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2449
			// _ = "end of CoverTab[11811]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2450
			_go_fuzz_dep_.CoverTab[11812]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2450
			// _ = "end of CoverTab[11812]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2450
		}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2450
		// _ = "end of CoverTab[11808]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2450
		_go_fuzz_dep_.CoverTab[11809]++

												if length-n < uint16(len(t))+1 {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2452
			_go_fuzz_dep_.CoverTab[11813]++
													return TXTResource{}, errCalcLen
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2453
			// _ = "end of CoverTab[11813]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2454
			_go_fuzz_dep_.CoverTab[11814]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2454
			// _ = "end of CoverTab[11814]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2454
		}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2454
		// _ = "end of CoverTab[11809]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2454
		_go_fuzz_dep_.CoverTab[11810]++
												n += uint16(len(t)) + 1
												txts = append(txts, t)
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2456
		// _ = "end of CoverTab[11810]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2457
	// _ = "end of CoverTab[11806]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2457
	_go_fuzz_dep_.CoverTab[11807]++
											return TXTResource{txts}, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2458
	// _ = "end of CoverTab[11807]"
}

// An SRVResource is an SRV Resource record.
type SRVResource struct {
	Priority	uint16
	Weight		uint16
	Port		uint16
	Target		Name	// Not compressed as per RFC 2782.
}

func (r *SRVResource) realType() Type {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2469
	_go_fuzz_dep_.CoverTab[11815]++
											return TypeSRV
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2470
	// _ = "end of CoverTab[11815]"
}

// pack appends the wire format of the SRVResource to msg.
func (r *SRVResource) pack(msg []byte, compression map[string]int, compressionOff int) ([]byte, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2474
	_go_fuzz_dep_.CoverTab[11816]++
											oldMsg := msg
											msg = packUint16(msg, r.Priority)
											msg = packUint16(msg, r.Weight)
											msg = packUint16(msg, r.Port)
											msg, err := r.Target.pack(msg, nil, compressionOff)
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2480
		_go_fuzz_dep_.CoverTab[11818]++
												return oldMsg, &nestedError{"SRVResource.Target", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2481
		// _ = "end of CoverTab[11818]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2482
		_go_fuzz_dep_.CoverTab[11819]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2482
		// _ = "end of CoverTab[11819]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2482
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2482
	// _ = "end of CoverTab[11816]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2482
	_go_fuzz_dep_.CoverTab[11817]++
											return msg, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2483
	// _ = "end of CoverTab[11817]"
}

// GoString implements fmt.GoStringer.GoString.
func (r *SRVResource) GoString() string {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2487
	_go_fuzz_dep_.CoverTab[11820]++
											return "dnsmessage.SRVResource{" +
		"Priority: " + printUint16(r.Priority) + ", " +
		"Weight: " + printUint16(r.Weight) + ", " +
		"Port: " + printUint16(r.Port) + ", " +
		"Target: " + r.Target.GoString() + "}"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2492
	// _ = "end of CoverTab[11820]"
}

func unpackSRVResource(msg []byte, off int) (SRVResource, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2495
	_go_fuzz_dep_.CoverTab[11821]++
											priority, off, err := unpackUint16(msg, off)
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2497
		_go_fuzz_dep_.CoverTab[11826]++
												return SRVResource{}, &nestedError{"Priority", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2498
		// _ = "end of CoverTab[11826]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2499
		_go_fuzz_dep_.CoverTab[11827]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2499
		// _ = "end of CoverTab[11827]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2499
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2499
	// _ = "end of CoverTab[11821]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2499
	_go_fuzz_dep_.CoverTab[11822]++
											weight, off, err := unpackUint16(msg, off)
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2501
		_go_fuzz_dep_.CoverTab[11828]++
												return SRVResource{}, &nestedError{"Weight", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2502
		// _ = "end of CoverTab[11828]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2503
		_go_fuzz_dep_.CoverTab[11829]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2503
		// _ = "end of CoverTab[11829]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2503
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2503
	// _ = "end of CoverTab[11822]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2503
	_go_fuzz_dep_.CoverTab[11823]++
											port, off, err := unpackUint16(msg, off)
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2505
		_go_fuzz_dep_.CoverTab[11830]++
												return SRVResource{}, &nestedError{"Port", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2506
		// _ = "end of CoverTab[11830]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2507
		_go_fuzz_dep_.CoverTab[11831]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2507
		// _ = "end of CoverTab[11831]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2507
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2507
	// _ = "end of CoverTab[11823]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2507
	_go_fuzz_dep_.CoverTab[11824]++
											var target Name
											if _, err := target.unpackCompressed(msg, off, false); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2509
		_go_fuzz_dep_.CoverTab[11832]++
												return SRVResource{}, &nestedError{"Target", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2510
		// _ = "end of CoverTab[11832]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2511
		_go_fuzz_dep_.CoverTab[11833]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2511
		// _ = "end of CoverTab[11833]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2511
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2511
	// _ = "end of CoverTab[11824]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2511
	_go_fuzz_dep_.CoverTab[11825]++
											return SRVResource{priority, weight, port, target}, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2512
	// _ = "end of CoverTab[11825]"
}

// An AResource is an A Resource record.
type AResource struct {
	A [4]byte
}

func (r *AResource) realType() Type {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2520
	_go_fuzz_dep_.CoverTab[11834]++
											return TypeA
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2521
	// _ = "end of CoverTab[11834]"
}

// pack appends the wire format of the AResource to msg.
func (r *AResource) pack(msg []byte, compression map[string]int, compressionOff int) ([]byte, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2525
	_go_fuzz_dep_.CoverTab[11835]++
											return packBytes(msg, r.A[:]), nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2526
	// _ = "end of CoverTab[11835]"
}

// GoString implements fmt.GoStringer.GoString.
func (r *AResource) GoString() string {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2530
	_go_fuzz_dep_.CoverTab[11836]++
											return "dnsmessage.AResource{" +
		"A: [4]byte{" + printByteSlice(r.A[:]) + "}}"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2532
	// _ = "end of CoverTab[11836]"
}

func unpackAResource(msg []byte, off int) (AResource, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2535
	_go_fuzz_dep_.CoverTab[11837]++
											var a [4]byte
											if _, err := unpackBytes(msg, off, a[:]); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2537
		_go_fuzz_dep_.CoverTab[11839]++
												return AResource{}, err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2538
		// _ = "end of CoverTab[11839]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2539
		_go_fuzz_dep_.CoverTab[11840]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2539
		// _ = "end of CoverTab[11840]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2539
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2539
	// _ = "end of CoverTab[11837]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2539
	_go_fuzz_dep_.CoverTab[11838]++
											return AResource{a}, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2540
	// _ = "end of CoverTab[11838]"
}

// An AAAAResource is an AAAA Resource record.
type AAAAResource struct {
	AAAA [16]byte
}

func (r *AAAAResource) realType() Type {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2548
	_go_fuzz_dep_.CoverTab[11841]++
											return TypeAAAA
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2549
	// _ = "end of CoverTab[11841]"
}

// GoString implements fmt.GoStringer.GoString.
func (r *AAAAResource) GoString() string {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2553
	_go_fuzz_dep_.CoverTab[11842]++
											return "dnsmessage.AAAAResource{" +
		"AAAA: [16]byte{" + printByteSlice(r.AAAA[:]) + "}}"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2555
	// _ = "end of CoverTab[11842]"
}

// pack appends the wire format of the AAAAResource to msg.
func (r *AAAAResource) pack(msg []byte, compression map[string]int, compressionOff int) ([]byte, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2559
	_go_fuzz_dep_.CoverTab[11843]++
											return packBytes(msg, r.AAAA[:]), nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2560
	// _ = "end of CoverTab[11843]"
}

func unpackAAAAResource(msg []byte, off int) (AAAAResource, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2563
	_go_fuzz_dep_.CoverTab[11844]++
											var aaaa [16]byte
											if _, err := unpackBytes(msg, off, aaaa[:]); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2565
		_go_fuzz_dep_.CoverTab[11846]++
												return AAAAResource{}, err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2566
		// _ = "end of CoverTab[11846]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2567
		_go_fuzz_dep_.CoverTab[11847]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2567
		// _ = "end of CoverTab[11847]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2567
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2567
	// _ = "end of CoverTab[11844]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2567
	_go_fuzz_dep_.CoverTab[11845]++
											return AAAAResource{aaaa}, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2568
	// _ = "end of CoverTab[11845]"
}

// An OPTResource is an OPT pseudo Resource record.
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2571
//
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2571
// The pseudo resource record is part of the extension mechanisms for DNS
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2571
// as defined in RFC 6891.
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2575
type OPTResource struct {
	Options []Option
}

// An Option represents a DNS message option within OPTResource.
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2579
//
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2579
// The message option is part of the extension mechanisms for DNS as
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2579
// defined in RFC 6891.
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2583
type Option struct {
	Code	uint16	// option code
	Data	[]byte
}

// GoString implements fmt.GoStringer.GoString.
func (o *Option) GoString() string {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2589
	_go_fuzz_dep_.CoverTab[11848]++
											return "dnsmessage.Option{" +
		"Code: " + printUint16(o.Code) + ", " +
		"Data: []byte{" + printByteSlice(o.Data) + "}}"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2592
	// _ = "end of CoverTab[11848]"
}

func (r *OPTResource) realType() Type {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2595
	_go_fuzz_dep_.CoverTab[11849]++
											return TypeOPT
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2596
	// _ = "end of CoverTab[11849]"
}

func (r *OPTResource) pack(msg []byte, compression map[string]int, compressionOff int) ([]byte, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2599
	_go_fuzz_dep_.CoverTab[11850]++
											for _, opt := range r.Options {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2600
		_go_fuzz_dep_.CoverTab[11852]++
												msg = packUint16(msg, opt.Code)
												l := uint16(len(opt.Data))
												msg = packUint16(msg, l)
												msg = packBytes(msg, opt.Data)
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2604
		// _ = "end of CoverTab[11852]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2605
	// _ = "end of CoverTab[11850]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2605
	_go_fuzz_dep_.CoverTab[11851]++
											return msg, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2606
	// _ = "end of CoverTab[11851]"
}

// GoString implements fmt.GoStringer.GoString.
func (r *OPTResource) GoString() string {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2610
	_go_fuzz_dep_.CoverTab[11853]++
											s := "dnsmessage.OPTResource{Options: []dnsmessage.Option{"
											if len(r.Options) == 0 {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2612
		_go_fuzz_dep_.CoverTab[11856]++
												return s + "}}"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2613
		// _ = "end of CoverTab[11856]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2614
		_go_fuzz_dep_.CoverTab[11857]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2614
		// _ = "end of CoverTab[11857]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2614
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2614
	// _ = "end of CoverTab[11853]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2614
	_go_fuzz_dep_.CoverTab[11854]++
											s += r.Options[0].GoString()
											for _, o := range r.Options[1:] {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2616
		_go_fuzz_dep_.CoverTab[11858]++
												s += ", " + o.GoString()
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2617
		// _ = "end of CoverTab[11858]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2618
	// _ = "end of CoverTab[11854]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2618
	_go_fuzz_dep_.CoverTab[11855]++
											return s + "}}"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2619
	// _ = "end of CoverTab[11855]"
}

func unpackOPTResource(msg []byte, off int, length uint16) (OPTResource, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2622
	_go_fuzz_dep_.CoverTab[11859]++
											var opts []Option
											for oldOff := off; off < oldOff+int(length); {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2624
		_go_fuzz_dep_.CoverTab[11861]++
												var err error
												var o Option
												o.Code, off, err = unpackUint16(msg, off)
												if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2628
			_go_fuzz_dep_.CoverTab[11865]++
													return OPTResource{}, &nestedError{"Code", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2629
			// _ = "end of CoverTab[11865]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2630
			_go_fuzz_dep_.CoverTab[11866]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2630
			// _ = "end of CoverTab[11866]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2630
		}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2630
		// _ = "end of CoverTab[11861]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2630
		_go_fuzz_dep_.CoverTab[11862]++
												var l uint16
												l, off, err = unpackUint16(msg, off)
												if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2633
			_go_fuzz_dep_.CoverTab[11867]++
													return OPTResource{}, &nestedError{"Data", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2634
			// _ = "end of CoverTab[11867]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2635
			_go_fuzz_dep_.CoverTab[11868]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2635
			// _ = "end of CoverTab[11868]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2635
		}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2635
		// _ = "end of CoverTab[11862]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2635
		_go_fuzz_dep_.CoverTab[11863]++
												o.Data = make([]byte, l)
												if copy(o.Data, msg[off:]) != int(l) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2637
			_go_fuzz_dep_.CoverTab[11869]++
													return OPTResource{}, &nestedError{"Data", errCalcLen}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2638
			// _ = "end of CoverTab[11869]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2639
			_go_fuzz_dep_.CoverTab[11870]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2639
			// _ = "end of CoverTab[11870]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2639
		}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2639
		// _ = "end of CoverTab[11863]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2639
		_go_fuzz_dep_.CoverTab[11864]++
												off += int(l)
												opts = append(opts, o)
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2641
		// _ = "end of CoverTab[11864]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2642
	// _ = "end of CoverTab[11859]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2642
	_go_fuzz_dep_.CoverTab[11860]++
											return OPTResource{opts}, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2643
	// _ = "end of CoverTab[11860]"
}

// An UnknownResource is a catch-all container for unknown record types.
type UnknownResource struct {
	Type	Type
	Data	[]byte
}

func (r *UnknownResource) realType() Type {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2652
	_go_fuzz_dep_.CoverTab[11871]++
											return r.Type
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2653
	// _ = "end of CoverTab[11871]"
}

// pack appends the wire format of the UnknownResource to msg.
func (r *UnknownResource) pack(msg []byte, compression map[string]int, compressionOff int) ([]byte, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2657
	_go_fuzz_dep_.CoverTab[11872]++
											return packBytes(msg, r.Data[:]), nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2658
	// _ = "end of CoverTab[11872]"
}

// GoString implements fmt.GoStringer.GoString.
func (r *UnknownResource) GoString() string {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2662
	_go_fuzz_dep_.CoverTab[11873]++
											return "dnsmessage.UnknownResource{" +
		"Type: " + r.Type.GoString() + ", " +
		"Data: []byte{" + printByteSlice(r.Data) + "}}"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2665
	// _ = "end of CoverTab[11873]"
}

func unpackUnknownResource(recordType Type, msg []byte, off int, length uint16) (UnknownResource, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2668
	_go_fuzz_dep_.CoverTab[11874]++
											parsed := UnknownResource{
		Type:	recordType,
		Data:	make([]byte, length),
	}
	if _, err := unpackBytes(msg, off, parsed.Data); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2673
		_go_fuzz_dep_.CoverTab[11876]++
												return UnknownResource{}, err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2674
		// _ = "end of CoverTab[11876]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2675
		_go_fuzz_dep_.CoverTab[11877]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2675
		// _ = "end of CoverTab[11877]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2675
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2675
	// _ = "end of CoverTab[11874]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2675
	_go_fuzz_dep_.CoverTab[11875]++
											return parsed, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2676
	// _ = "end of CoverTab[11875]"
}

//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2677
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2677
var _ = _go_fuzz_dep_.CoverTab
