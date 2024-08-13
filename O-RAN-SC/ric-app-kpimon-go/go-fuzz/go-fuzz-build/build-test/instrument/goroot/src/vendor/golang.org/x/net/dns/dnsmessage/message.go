// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:5
// Package dnsmessage provides a mostly RFC 1035 compliant implementation of
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:5
// DNS message packing and unpacking.
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:5
//
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:5
// The package also supports messages with Extension Mechanisms for DNS
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:5
// (EDNS(0)) as defined in RFC 6891.
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:5
//
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:5
// This implementation is designed to minimize heap allocations and avoid
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:5
// unnecessary packing and unpacking as much as possible.
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:13
package dnsmessage

//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:13
import (
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:13
	_go_fuzz_dep_ "go-fuzz-dep"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:13
)
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:13
import (
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:13
	_atomic_ "sync/atomic"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:13
)

import (
	"errors"
)

//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:21
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
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:64
	_go_fuzz_dep_.CoverTab[2499]++
										if n, ok := typeNames[t]; ok {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:65
		_go_fuzz_dep_.CoverTab[526092]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:65
		_go_fuzz_dep_.CoverTab[2501]++
											return n
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:66
		// _ = "end of CoverTab[2501]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:67
		_go_fuzz_dep_.CoverTab[526093]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:67
		_go_fuzz_dep_.CoverTab[2502]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:67
		// _ = "end of CoverTab[2502]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:67
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:67
	// _ = "end of CoverTab[2499]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:67
	_go_fuzz_dep_.CoverTab[2500]++
										return printUint16(uint16(t))
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:68
	// _ = "end of CoverTab[2500]"
}

// GoString implements fmt.GoStringer.GoString.
func (t Type) GoString() string {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:72
	_go_fuzz_dep_.CoverTab[2503]++
										if n, ok := typeNames[t]; ok {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:73
		_go_fuzz_dep_.CoverTab[526094]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:73
		_go_fuzz_dep_.CoverTab[2505]++
											return "dnsmessage." + n
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:74
		// _ = "end of CoverTab[2505]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:75
		_go_fuzz_dep_.CoverTab[526095]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:75
		_go_fuzz_dep_.CoverTab[2506]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:75
		// _ = "end of CoverTab[2506]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:75
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:75
	// _ = "end of CoverTab[2503]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:75
	_go_fuzz_dep_.CoverTab[2504]++
										return printUint16(uint16(t))
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:76
	// _ = "end of CoverTab[2504]"
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
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:102
	_go_fuzz_dep_.CoverTab[2507]++
											if n, ok := classNames[c]; ok {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:103
		_go_fuzz_dep_.CoverTab[526096]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:103
		_go_fuzz_dep_.CoverTab[2509]++
												return n
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:104
		// _ = "end of CoverTab[2509]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:105
		_go_fuzz_dep_.CoverTab[526097]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:105
		_go_fuzz_dep_.CoverTab[2510]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:105
		// _ = "end of CoverTab[2510]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:105
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:105
	// _ = "end of CoverTab[2507]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:105
	_go_fuzz_dep_.CoverTab[2508]++
											return printUint16(uint16(c))
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:106
	// _ = "end of CoverTab[2508]"
}

// GoString implements fmt.GoStringer.GoString.
func (c Class) GoString() string {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:110
	_go_fuzz_dep_.CoverTab[2511]++
											if n, ok := classNames[c]; ok {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:111
		_go_fuzz_dep_.CoverTab[526098]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:111
		_go_fuzz_dep_.CoverTab[2513]++
												return "dnsmessage." + n
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:112
		// _ = "end of CoverTab[2513]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:113
		_go_fuzz_dep_.CoverTab[526099]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:113
		_go_fuzz_dep_.CoverTab[2514]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:113
		// _ = "end of CoverTab[2514]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:113
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:113
	// _ = "end of CoverTab[2511]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:113
	_go_fuzz_dep_.CoverTab[2512]++
											return printUint16(uint16(c))
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:114
	// _ = "end of CoverTab[2512]"
}

// An OpCode is a DNS operation code.
type OpCode uint16

// GoString implements fmt.GoStringer.GoString.
func (o OpCode) GoString() string {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:121
	_go_fuzz_dep_.CoverTab[2515]++
											return printUint16(uint16(o))
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:122
	// _ = "end of CoverTab[2515]"
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
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:148
	_go_fuzz_dep_.CoverTab[2516]++
											if n, ok := rCodeNames[r]; ok {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:149
		_go_fuzz_dep_.CoverTab[526100]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:149
		_go_fuzz_dep_.CoverTab[2518]++
												return n
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:150
		// _ = "end of CoverTab[2518]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:151
		_go_fuzz_dep_.CoverTab[526101]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:151
		_go_fuzz_dep_.CoverTab[2519]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:151
		// _ = "end of CoverTab[2519]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:151
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:151
	// _ = "end of CoverTab[2516]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:151
	_go_fuzz_dep_.CoverTab[2517]++
											return printUint16(uint16(r))
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:152
	// _ = "end of CoverTab[2517]"
}

// GoString implements fmt.GoStringer.GoString.
func (r RCode) GoString() string {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:156
	_go_fuzz_dep_.CoverTab[2520]++
											if n, ok := rCodeNames[r]; ok {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:157
		_go_fuzz_dep_.CoverTab[526102]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:157
		_go_fuzz_dep_.CoverTab[2522]++
												return "dnsmessage." + n
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:158
		// _ = "end of CoverTab[2522]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:159
		_go_fuzz_dep_.CoverTab[526103]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:159
		_go_fuzz_dep_.CoverTab[2523]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:159
		// _ = "end of CoverTab[2523]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:159
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:159
	// _ = "end of CoverTab[2520]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:159
	_go_fuzz_dep_.CoverTab[2521]++
											return printUint16(uint16(r))
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:160
	// _ = "end of CoverTab[2521]"
}

func printPaddedUint8(i uint8) string {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:163
	_go_fuzz_dep_.CoverTab[2524]++
											b := byte(i)
											return string([]byte{
		b/100 + '0',
		b/10%10 + '0',
		b%10 + '0',
	})
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:169
	// _ = "end of CoverTab[2524]"
}

func printUint8Bytes(buf []byte, i uint8) []byte {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:172
	_go_fuzz_dep_.CoverTab[2525]++
											b := byte(i)
											if i >= 100 {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:174
		_go_fuzz_dep_.CoverTab[526104]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:174
		_go_fuzz_dep_.CoverTab[2528]++
												buf = append(buf, b/100+'0')
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:175
		// _ = "end of CoverTab[2528]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:176
		_go_fuzz_dep_.CoverTab[526105]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:176
		_go_fuzz_dep_.CoverTab[2529]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:176
		// _ = "end of CoverTab[2529]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:176
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:176
	// _ = "end of CoverTab[2525]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:176
	_go_fuzz_dep_.CoverTab[2526]++
											if i >= 10 {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:177
		_go_fuzz_dep_.CoverTab[526106]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:177
		_go_fuzz_dep_.CoverTab[2530]++
												buf = append(buf, b/10%10+'0')
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:178
		// _ = "end of CoverTab[2530]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:179
		_go_fuzz_dep_.CoverTab[526107]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:179
		_go_fuzz_dep_.CoverTab[2531]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:179
		// _ = "end of CoverTab[2531]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:179
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:179
	// _ = "end of CoverTab[2526]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:179
	_go_fuzz_dep_.CoverTab[2527]++
											return append(buf, b%10+'0')
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:180
	// _ = "end of CoverTab[2527]"
}

func printByteSlice(b []byte) string {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:183
	_go_fuzz_dep_.CoverTab[2532]++
											if len(b) == 0 {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:184
		_go_fuzz_dep_.CoverTab[526108]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:184
		_go_fuzz_dep_.CoverTab[2535]++
												return ""
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:185
		// _ = "end of CoverTab[2535]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:186
		_go_fuzz_dep_.CoverTab[526109]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:186
		_go_fuzz_dep_.CoverTab[2536]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:186
		// _ = "end of CoverTab[2536]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:186
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:186
	// _ = "end of CoverTab[2532]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:186
	_go_fuzz_dep_.CoverTab[2533]++
											buf := make([]byte, 0, 5*len(b))
											buf = printUint8Bytes(buf, uint8(b[0]))
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:188
	_go_fuzz_dep_.CoverTab[786574] = 0
											for _, n := range b[1:] {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:189
		if _go_fuzz_dep_.CoverTab[786574] == 0 {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:189
			_go_fuzz_dep_.CoverTab[526639]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:189
		} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:189
			_go_fuzz_dep_.CoverTab[526640]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:189
		}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:189
		_go_fuzz_dep_.CoverTab[786574] = 1
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:189
		_go_fuzz_dep_.CoverTab[2537]++
												buf = append(buf, ',', ' ')
												buf = printUint8Bytes(buf, uint8(n))
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:191
		// _ = "end of CoverTab[2537]"
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:192
	if _go_fuzz_dep_.CoverTab[786574] == 0 {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:192
		_go_fuzz_dep_.CoverTab[526641]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:192
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:192
		_go_fuzz_dep_.CoverTab[526642]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:192
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:192
	// _ = "end of CoverTab[2533]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:192
	_go_fuzz_dep_.CoverTab[2534]++
											return string(buf)
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:193
	// _ = "end of CoverTab[2534]"
}

const hexDigits = "0123456789abcdef"

func printString(str []byte) string {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:198
	_go_fuzz_dep_.CoverTab[2538]++
											buf := make([]byte, 0, len(str))
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:199
	_go_fuzz_dep_.CoverTab[786575] = 0
											for i := 0; i < len(str); i++ {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:200
		if _go_fuzz_dep_.CoverTab[786575] == 0 {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:200
			_go_fuzz_dep_.CoverTab[526643]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:200
		} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:200
			_go_fuzz_dep_.CoverTab[526644]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:200
		}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:200
		_go_fuzz_dep_.CoverTab[786575] = 1
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:200
		_go_fuzz_dep_.CoverTab[2540]++
												c := str[i]
												if c == '.' || func() bool {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:202
			_go_fuzz_dep_.CoverTab[2542]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:202
			return c == '-'
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:202
			// _ = "end of CoverTab[2542]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:202
		}() || func() bool {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:202
			_go_fuzz_dep_.CoverTab[2543]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:202
			return c == ' '
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:202
			// _ = "end of CoverTab[2543]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:202
		}() || func() bool {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:202
			_go_fuzz_dep_.CoverTab[2544]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:202
			return 'A' <= c && func() bool {
														_go_fuzz_dep_.CoverTab[2545]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:203
				return c <= 'Z'
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:203
				// _ = "end of CoverTab[2545]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:203
			}()
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:203
			// _ = "end of CoverTab[2544]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:203
		}() || func() bool {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:203
			_go_fuzz_dep_.CoverTab[2546]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:203
			return 'a' <= c && func() bool {
														_go_fuzz_dep_.CoverTab[2547]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:204
				return c <= 'z'
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:204
				// _ = "end of CoverTab[2547]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:204
			}()
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:204
			// _ = "end of CoverTab[2546]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:204
		}() || func() bool {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:204
			_go_fuzz_dep_.CoverTab[2548]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:204
			return '0' <= c && func() bool {
														_go_fuzz_dep_.CoverTab[2549]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:205
				return c <= '9'
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:205
				// _ = "end of CoverTab[2549]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:205
			}()
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:205
			// _ = "end of CoverTab[2548]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:205
		}() {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:205
			_go_fuzz_dep_.CoverTab[526110]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:205
			_go_fuzz_dep_.CoverTab[2550]++
													buf = append(buf, c)
													continue
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:207
			// _ = "end of CoverTab[2550]"
		} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:208
			_go_fuzz_dep_.CoverTab[526111]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:208
			_go_fuzz_dep_.CoverTab[2551]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:208
			// _ = "end of CoverTab[2551]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:208
		}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:208
		// _ = "end of CoverTab[2540]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:208
		_go_fuzz_dep_.CoverTab[2541]++

												upper := c >> 4
												lower := (c << 4) >> 4
												buf = append(
			buf,
			'\\',
			'x',
			hexDigits[upper],
			hexDigits[lower],
		)
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:218
		// _ = "end of CoverTab[2541]"
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:219
	if _go_fuzz_dep_.CoverTab[786575] == 0 {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:219
		_go_fuzz_dep_.CoverTab[526645]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:219
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:219
		_go_fuzz_dep_.CoverTab[526646]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:219
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:219
	// _ = "end of CoverTab[2538]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:219
	_go_fuzz_dep_.CoverTab[2539]++
											return string(buf)
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:220
	// _ = "end of CoverTab[2539]"
}

func printUint16(i uint16) string {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:223
	_go_fuzz_dep_.CoverTab[2552]++
											return printUint32(uint32(i))
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:224
	// _ = "end of CoverTab[2552]"
}

func printUint32(i uint32) string {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:227
	_go_fuzz_dep_.CoverTab[2553]++

											buf := make([]byte, 10)
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:229
	_go_fuzz_dep_.CoverTab[786576] = 0
											for b, d := buf, uint32(1000000000); d > 0; d /= 10 {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:230
		if _go_fuzz_dep_.CoverTab[786576] == 0 {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:230
			_go_fuzz_dep_.CoverTab[526647]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:230
		} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:230
			_go_fuzz_dep_.CoverTab[526648]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:230
		}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:230
		_go_fuzz_dep_.CoverTab[786576] = 1
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:230
		_go_fuzz_dep_.CoverTab[2555]++
												b[0] = byte(i/d%10 + '0')
												if b[0] == '0' && func() bool {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:232
			_go_fuzz_dep_.CoverTab[2557]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:232
			return len(b) == len(buf)
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:232
			// _ = "end of CoverTab[2557]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:232
		}() && func() bool {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:232
			_go_fuzz_dep_.CoverTab[2558]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:232
			return len(buf) > 1
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:232
			// _ = "end of CoverTab[2558]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:232
		}() {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:232
			_go_fuzz_dep_.CoverTab[526112]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:232
			_go_fuzz_dep_.CoverTab[2559]++
													buf = buf[1:]
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:233
			// _ = "end of CoverTab[2559]"
		} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:234
			_go_fuzz_dep_.CoverTab[526113]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:234
			_go_fuzz_dep_.CoverTab[2560]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:234
			// _ = "end of CoverTab[2560]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:234
		}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:234
		// _ = "end of CoverTab[2555]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:234
		_go_fuzz_dep_.CoverTab[2556]++
												b = b[1:]
												i %= d
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:236
		// _ = "end of CoverTab[2556]"
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:237
	if _go_fuzz_dep_.CoverTab[786576] == 0 {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:237
		_go_fuzz_dep_.CoverTab[526649]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:237
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:237
		_go_fuzz_dep_.CoverTab[526650]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:237
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:237
	// _ = "end of CoverTab[2553]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:237
	_go_fuzz_dep_.CoverTab[2554]++
											return string(buf)
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:238
	// _ = "end of CoverTab[2554]"
}

func printBool(b bool) string {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:241
	_go_fuzz_dep_.CoverTab[2561]++
											if b {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:242
		_go_fuzz_dep_.CoverTab[526114]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:242
		_go_fuzz_dep_.CoverTab[2563]++
												return "true"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:243
		// _ = "end of CoverTab[2563]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:244
		_go_fuzz_dep_.CoverTab[526115]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:244
		_go_fuzz_dep_.CoverTab[2564]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:244
		// _ = "end of CoverTab[2564]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:244
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:244
	// _ = "end of CoverTab[2561]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:244
	_go_fuzz_dep_.CoverTab[2562]++
											return "false"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:245
	// _ = "end of CoverTab[2562]"
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
	errInvalidName		= errors.New("invalid dns name")
	errNilResouceBody	= errors.New("nil resource body")
	errResourceLen		= errors.New("insufficient data for resource body length")
	errSegTooLong		= errors.New("segment length too long")
	errNameTooLong		= errors.New("name too long")
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
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:309
	_go_fuzz_dep_.CoverTab[2565]++
											return e.s + ": " + e.err.Error()
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:310
	// _ = "end of CoverTab[2565]"
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
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:327
	_go_fuzz_dep_.CoverTab[2566]++
											id = m.ID
											bits = uint16(m.OpCode)<<11 | uint16(m.RCode)
											if m.RecursionAvailable {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:330
		_go_fuzz_dep_.CoverTab[526116]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:330
		_go_fuzz_dep_.CoverTab[2574]++
												bits |= headerBitRA
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:331
		// _ = "end of CoverTab[2574]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:332
		_go_fuzz_dep_.CoverTab[526117]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:332
		_go_fuzz_dep_.CoverTab[2575]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:332
		// _ = "end of CoverTab[2575]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:332
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:332
	// _ = "end of CoverTab[2566]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:332
	_go_fuzz_dep_.CoverTab[2567]++
											if m.RecursionDesired {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:333
		_go_fuzz_dep_.CoverTab[526118]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:333
		_go_fuzz_dep_.CoverTab[2576]++
												bits |= headerBitRD
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:334
		// _ = "end of CoverTab[2576]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:335
		_go_fuzz_dep_.CoverTab[526119]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:335
		_go_fuzz_dep_.CoverTab[2577]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:335
		// _ = "end of CoverTab[2577]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:335
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:335
	// _ = "end of CoverTab[2567]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:335
	_go_fuzz_dep_.CoverTab[2568]++
											if m.Truncated {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:336
		_go_fuzz_dep_.CoverTab[526120]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:336
		_go_fuzz_dep_.CoverTab[2578]++
												bits |= headerBitTC
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:337
		// _ = "end of CoverTab[2578]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:338
		_go_fuzz_dep_.CoverTab[526121]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:338
		_go_fuzz_dep_.CoverTab[2579]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:338
		// _ = "end of CoverTab[2579]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:338
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:338
	// _ = "end of CoverTab[2568]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:338
	_go_fuzz_dep_.CoverTab[2569]++
											if m.Authoritative {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:339
		_go_fuzz_dep_.CoverTab[526122]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:339
		_go_fuzz_dep_.CoverTab[2580]++
												bits |= headerBitAA
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:340
		// _ = "end of CoverTab[2580]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:341
		_go_fuzz_dep_.CoverTab[526123]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:341
		_go_fuzz_dep_.CoverTab[2581]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:341
		// _ = "end of CoverTab[2581]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:341
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:341
	// _ = "end of CoverTab[2569]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:341
	_go_fuzz_dep_.CoverTab[2570]++
											if m.Response {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:342
		_go_fuzz_dep_.CoverTab[526124]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:342
		_go_fuzz_dep_.CoverTab[2582]++
												bits |= headerBitQR
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:343
		// _ = "end of CoverTab[2582]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:344
		_go_fuzz_dep_.CoverTab[526125]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:344
		_go_fuzz_dep_.CoverTab[2583]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:344
		// _ = "end of CoverTab[2583]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:344
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:344
	// _ = "end of CoverTab[2570]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:344
	_go_fuzz_dep_.CoverTab[2571]++
											if m.AuthenticData {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:345
		_go_fuzz_dep_.CoverTab[526126]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:345
		_go_fuzz_dep_.CoverTab[2584]++
												bits |= headerBitAD
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:346
		// _ = "end of CoverTab[2584]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:347
		_go_fuzz_dep_.CoverTab[526127]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:347
		_go_fuzz_dep_.CoverTab[2585]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:347
		// _ = "end of CoverTab[2585]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:347
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:347
	// _ = "end of CoverTab[2571]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:347
	_go_fuzz_dep_.CoverTab[2572]++
											if m.CheckingDisabled {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:348
		_go_fuzz_dep_.CoverTab[526128]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:348
		_go_fuzz_dep_.CoverTab[2586]++
												bits |= headerBitCD
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:349
		// _ = "end of CoverTab[2586]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:350
		_go_fuzz_dep_.CoverTab[526129]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:350
		_go_fuzz_dep_.CoverTab[2587]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:350
		// _ = "end of CoverTab[2587]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:350
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:350
	// _ = "end of CoverTab[2572]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:350
	_go_fuzz_dep_.CoverTab[2573]++
											return
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:351
	// _ = "end of CoverTab[2573]"
}

// GoString implements fmt.GoStringer.GoString.
func (m *Header) GoString() string {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:355
	_go_fuzz_dep_.CoverTab[2588]++
											return "dnsmessage.Header{" +
		"ID: " + printUint16(m.ID) + ", " +
		"Response: " + printBool(m.Response) + ", " +
		"OpCode: " + m.OpCode.GoString() + ", " +
		"Authoritative: " + printBool(m.Authoritative) + ", " +
		"Truncated: " + printBool(m.Truncated) + ", " +
		"RecursionDesired: " + printBool(m.RecursionDesired) + ", " +
		"RecursionAvailable: " + printBool(m.RecursionAvailable) + ", " +
		"RCode: " + m.RCode.GoString() + "}"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:364
	// _ = "end of CoverTab[2588]"
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
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:414
	_go_fuzz_dep_.CoverTab[2589]++
											switch sec {
	case sectionQuestions:
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:416
		_go_fuzz_dep_.CoverTab[526130]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:416
		_go_fuzz_dep_.CoverTab[2591]++
												return h.questions
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:417
		// _ = "end of CoverTab[2591]"
	case sectionAnswers:
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:418
		_go_fuzz_dep_.CoverTab[526131]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:418
		_go_fuzz_dep_.CoverTab[2592]++
												return h.answers
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:419
		// _ = "end of CoverTab[2592]"
	case sectionAuthorities:
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:420
		_go_fuzz_dep_.CoverTab[526132]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:420
		_go_fuzz_dep_.CoverTab[2593]++
												return h.authorities
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:421
		// _ = "end of CoverTab[2593]"
	case sectionAdditionals:
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:422
		_go_fuzz_dep_.CoverTab[526133]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:422
		_go_fuzz_dep_.CoverTab[2594]++
												return h.additionals
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:423
		// _ = "end of CoverTab[2594]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:423
	default:
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:423
		_go_fuzz_dep_.CoverTab[526134]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:423
		_go_fuzz_dep_.CoverTab[2595]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:423
		// _ = "end of CoverTab[2595]"
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:424
	// _ = "end of CoverTab[2589]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:424
	_go_fuzz_dep_.CoverTab[2590]++
											return 0
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:425
	// _ = "end of CoverTab[2590]"
}

// pack appends the wire format of the header to msg.
func (h *header) pack(msg []byte) []byte {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:429
	_go_fuzz_dep_.CoverTab[2596]++
											msg = packUint16(msg, h.id)
											msg = packUint16(msg, h.bits)
											msg = packUint16(msg, h.questions)
											msg = packUint16(msg, h.answers)
											msg = packUint16(msg, h.authorities)
											return packUint16(msg, h.additionals)
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:435
	// _ = "end of CoverTab[2596]"
}

func (h *header) unpack(msg []byte, off int) (int, error) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:438
	_go_fuzz_dep_.CoverTab[2597]++
											newOff := off
											var err error
											if h.id, newOff, err = unpackUint16(msg, newOff); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:441
		_go_fuzz_dep_.CoverTab[526135]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:441
		_go_fuzz_dep_.CoverTab[2604]++
												return off, &nestedError{"id", err}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:442
		// _ = "end of CoverTab[2604]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:443
		_go_fuzz_dep_.CoverTab[526136]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:443
		_go_fuzz_dep_.CoverTab[2605]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:443
		// _ = "end of CoverTab[2605]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:443
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:443
	// _ = "end of CoverTab[2597]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:443
	_go_fuzz_dep_.CoverTab[2598]++
											if h.bits, newOff, err = unpackUint16(msg, newOff); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:444
		_go_fuzz_dep_.CoverTab[526137]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:444
		_go_fuzz_dep_.CoverTab[2606]++
												return off, &nestedError{"bits", err}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:445
		// _ = "end of CoverTab[2606]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:446
		_go_fuzz_dep_.CoverTab[526138]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:446
		_go_fuzz_dep_.CoverTab[2607]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:446
		// _ = "end of CoverTab[2607]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:446
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:446
	// _ = "end of CoverTab[2598]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:446
	_go_fuzz_dep_.CoverTab[2599]++
											if h.questions, newOff, err = unpackUint16(msg, newOff); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:447
		_go_fuzz_dep_.CoverTab[526139]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:447
		_go_fuzz_dep_.CoverTab[2608]++
												return off, &nestedError{"questions", err}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:448
		// _ = "end of CoverTab[2608]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:449
		_go_fuzz_dep_.CoverTab[526140]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:449
		_go_fuzz_dep_.CoverTab[2609]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:449
		// _ = "end of CoverTab[2609]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:449
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:449
	// _ = "end of CoverTab[2599]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:449
	_go_fuzz_dep_.CoverTab[2600]++
											if h.answers, newOff, err = unpackUint16(msg, newOff); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:450
		_go_fuzz_dep_.CoverTab[526141]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:450
		_go_fuzz_dep_.CoverTab[2610]++
												return off, &nestedError{"answers", err}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:451
		// _ = "end of CoverTab[2610]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:452
		_go_fuzz_dep_.CoverTab[526142]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:452
		_go_fuzz_dep_.CoverTab[2611]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:452
		// _ = "end of CoverTab[2611]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:452
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:452
	// _ = "end of CoverTab[2600]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:452
	_go_fuzz_dep_.CoverTab[2601]++
											if h.authorities, newOff, err = unpackUint16(msg, newOff); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:453
		_go_fuzz_dep_.CoverTab[526143]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:453
		_go_fuzz_dep_.CoverTab[2612]++
												return off, &nestedError{"authorities", err}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:454
		// _ = "end of CoverTab[2612]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:455
		_go_fuzz_dep_.CoverTab[526144]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:455
		_go_fuzz_dep_.CoverTab[2613]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:455
		// _ = "end of CoverTab[2613]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:455
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:455
	// _ = "end of CoverTab[2601]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:455
	_go_fuzz_dep_.CoverTab[2602]++
											if h.additionals, newOff, err = unpackUint16(msg, newOff); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:456
		_go_fuzz_dep_.CoverTab[526145]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:456
		_go_fuzz_dep_.CoverTab[2614]++
												return off, &nestedError{"additionals", err}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:457
		// _ = "end of CoverTab[2614]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:458
		_go_fuzz_dep_.CoverTab[526146]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:458
		_go_fuzz_dep_.CoverTab[2615]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:458
		// _ = "end of CoverTab[2615]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:458
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:458
	// _ = "end of CoverTab[2602]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:458
	_go_fuzz_dep_.CoverTab[2603]++
											return newOff, nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:459
	// _ = "end of CoverTab[2603]"
}

func (h *header) header() Header {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:462
	_go_fuzz_dep_.CoverTab[2616]++
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
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:474
	// _ = "end of CoverTab[2616]"
}

// A Resource is a DNS resource record.
type Resource struct {
	Header	ResourceHeader
	Body	ResourceBody
}

func (r *Resource) GoString() string {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:483
	_go_fuzz_dep_.CoverTab[2617]++
											return "dnsmessage.Resource{" +
		"Header: " + r.Header.GoString() +
		", Body: &" + r.Body.GoString() +
		"}"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:487
	// _ = "end of CoverTab[2617]"
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
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:504
	_go_fuzz_dep_.CoverTab[2618]++
											if r.Body == nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:505
		_go_fuzz_dep_.CoverTab[526147]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:505
		_go_fuzz_dep_.CoverTab[2623]++
												return msg, errNilResouceBody
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:506
		// _ = "end of CoverTab[2623]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:507
		_go_fuzz_dep_.CoverTab[526148]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:507
		_go_fuzz_dep_.CoverTab[2624]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:507
		// _ = "end of CoverTab[2624]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:507
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:507
	// _ = "end of CoverTab[2618]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:507
	_go_fuzz_dep_.CoverTab[2619]++
											oldMsg := msg
											r.Header.Type = r.Body.realType()
											msg, lenOff, err := r.Header.pack(msg, compression, compressionOff)
											if err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:511
		_go_fuzz_dep_.CoverTab[526149]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:511
		_go_fuzz_dep_.CoverTab[2625]++
												return msg, &nestedError{"ResourceHeader", err}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:512
		// _ = "end of CoverTab[2625]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:513
		_go_fuzz_dep_.CoverTab[526150]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:513
		_go_fuzz_dep_.CoverTab[2626]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:513
		// _ = "end of CoverTab[2626]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:513
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:513
	// _ = "end of CoverTab[2619]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:513
	_go_fuzz_dep_.CoverTab[2620]++
											preLen := len(msg)
											msg, err = r.Body.pack(msg, compression, compressionOff)
											if err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:516
		_go_fuzz_dep_.CoverTab[526151]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:516
		_go_fuzz_dep_.CoverTab[2627]++
												return msg, &nestedError{"content", err}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:517
		// _ = "end of CoverTab[2627]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:518
		_go_fuzz_dep_.CoverTab[526152]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:518
		_go_fuzz_dep_.CoverTab[2628]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:518
		// _ = "end of CoverTab[2628]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:518
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:518
	// _ = "end of CoverTab[2620]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:518
	_go_fuzz_dep_.CoverTab[2621]++
											if err := r.Header.fixLen(msg, lenOff, preLen); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:519
		_go_fuzz_dep_.CoverTab[526153]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:519
		_go_fuzz_dep_.CoverTab[2629]++
												return oldMsg, err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:520
		// _ = "end of CoverTab[2629]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:521
		_go_fuzz_dep_.CoverTab[526154]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:521
		_go_fuzz_dep_.CoverTab[2630]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:521
		// _ = "end of CoverTab[2630]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:521
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:521
	// _ = "end of CoverTab[2621]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:521
	_go_fuzz_dep_.CoverTab[2622]++
											return msg, nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:522
	// _ = "end of CoverTab[2622]"
}

// A Parser allows incrementally parsing a DNS message.
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:525
//
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:525
// When parsing is started, the Header is parsed. Next, each Question can be
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:525
// either parsed or skipped. Alternatively, all Questions can be skipped at
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:525
// once. When all Questions have been parsed, attempting to parse Questions
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:525
// will return the [ErrSectionDone] error.
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:525
// After all Questions have been either parsed or skipped, all
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:525
// Answers, Authorities and Additionals can be either parsed or skipped in the
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:525
// same way, and each type of Resource must be fully parsed or skipped before
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:525
// proceeding to the next type of Resource.
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:525
//
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:525
// Parser is safe to copy to preserve the parsing state.
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:525
//
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:525
// Note that there is no requirement to fully skip or parse the message.
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:539
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
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:551
	_go_fuzz_dep_.CoverTab[2631]++
											if p.msg != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:552
		_go_fuzz_dep_.CoverTab[526155]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:552
		_go_fuzz_dep_.CoverTab[2634]++
												*p = Parser{}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:553
		// _ = "end of CoverTab[2634]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:554
		_go_fuzz_dep_.CoverTab[526156]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:554
		_go_fuzz_dep_.CoverTab[2635]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:554
		// _ = "end of CoverTab[2635]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:554
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:554
	// _ = "end of CoverTab[2631]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:554
	_go_fuzz_dep_.CoverTab[2632]++
											p.msg = msg
											var err error
											if p.off, err = p.header.unpack(msg, 0); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:557
		_go_fuzz_dep_.CoverTab[526157]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:557
		_go_fuzz_dep_.CoverTab[2636]++
												return Header{}, &nestedError{"unpacking header", err}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:558
		// _ = "end of CoverTab[2636]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:559
		_go_fuzz_dep_.CoverTab[526158]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:559
		_go_fuzz_dep_.CoverTab[2637]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:559
		// _ = "end of CoverTab[2637]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:559
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:559
	// _ = "end of CoverTab[2632]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:559
	_go_fuzz_dep_.CoverTab[2633]++
											p.section = sectionQuestions
											return p.header.header(), nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:561
	// _ = "end of CoverTab[2633]"
}

func (p *Parser) checkAdvance(sec section) error {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:564
	_go_fuzz_dep_.CoverTab[2638]++
											if p.section < sec {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:565
		_go_fuzz_dep_.CoverTab[526159]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:565
		_go_fuzz_dep_.CoverTab[2642]++
												return ErrNotStarted
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:566
		// _ = "end of CoverTab[2642]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:567
		_go_fuzz_dep_.CoverTab[526160]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:567
		_go_fuzz_dep_.CoverTab[2643]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:567
		// _ = "end of CoverTab[2643]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:567
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:567
	// _ = "end of CoverTab[2638]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:567
	_go_fuzz_dep_.CoverTab[2639]++
											if p.section > sec {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:568
		_go_fuzz_dep_.CoverTab[526161]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:568
		_go_fuzz_dep_.CoverTab[2644]++
												return ErrSectionDone
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:569
		// _ = "end of CoverTab[2644]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:570
		_go_fuzz_dep_.CoverTab[526162]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:570
		_go_fuzz_dep_.CoverTab[2645]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:570
		// _ = "end of CoverTab[2645]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:570
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:570
	// _ = "end of CoverTab[2639]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:570
	_go_fuzz_dep_.CoverTab[2640]++
											p.resHeaderValid = false
											if p.index == int(p.header.count(sec)) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:572
		_go_fuzz_dep_.CoverTab[526163]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:572
		_go_fuzz_dep_.CoverTab[2646]++
												p.index = 0
												p.section++
												return ErrSectionDone
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:575
		// _ = "end of CoverTab[2646]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:576
		_go_fuzz_dep_.CoverTab[526164]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:576
		_go_fuzz_dep_.CoverTab[2647]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:576
		// _ = "end of CoverTab[2647]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:576
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:576
	// _ = "end of CoverTab[2640]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:576
	_go_fuzz_dep_.CoverTab[2641]++
											return nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:577
	// _ = "end of CoverTab[2641]"
}

func (p *Parser) resource(sec section) (Resource, error) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:580
	_go_fuzz_dep_.CoverTab[2648]++
											var r Resource
											var err error
											r.Header, err = p.resourceHeader(sec)
											if err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:584
		_go_fuzz_dep_.CoverTab[526165]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:584
		_go_fuzz_dep_.CoverTab[2651]++
												return r, err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:585
		// _ = "end of CoverTab[2651]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:586
		_go_fuzz_dep_.CoverTab[526166]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:586
		_go_fuzz_dep_.CoverTab[2652]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:586
		// _ = "end of CoverTab[2652]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:586
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:586
	// _ = "end of CoverTab[2648]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:586
	_go_fuzz_dep_.CoverTab[2649]++
											p.resHeaderValid = false
											r.Body, p.off, err = unpackResourceBody(p.msg, p.off, r.Header)
											if err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:589
		_go_fuzz_dep_.CoverTab[526167]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:589
		_go_fuzz_dep_.CoverTab[2653]++
												return Resource{}, &nestedError{"unpacking " + sectionNames[sec], err}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:590
		// _ = "end of CoverTab[2653]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:591
		_go_fuzz_dep_.CoverTab[526168]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:591
		_go_fuzz_dep_.CoverTab[2654]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:591
		// _ = "end of CoverTab[2654]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:591
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:591
	// _ = "end of CoverTab[2649]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:591
	_go_fuzz_dep_.CoverTab[2650]++
											p.index++
											return r, nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:593
	// _ = "end of CoverTab[2650]"
}

func (p *Parser) resourceHeader(sec section) (ResourceHeader, error) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:596
	_go_fuzz_dep_.CoverTab[2655]++
											if p.resHeaderValid {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:597
		_go_fuzz_dep_.CoverTab[526169]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:597
		_go_fuzz_dep_.CoverTab[2659]++
												return p.resHeader, nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:598
		// _ = "end of CoverTab[2659]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:599
		_go_fuzz_dep_.CoverTab[526170]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:599
		_go_fuzz_dep_.CoverTab[2660]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:599
		// _ = "end of CoverTab[2660]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:599
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:599
	// _ = "end of CoverTab[2655]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:599
	_go_fuzz_dep_.CoverTab[2656]++
											if err := p.checkAdvance(sec); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:600
		_go_fuzz_dep_.CoverTab[526171]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:600
		_go_fuzz_dep_.CoverTab[2661]++
												return ResourceHeader{}, err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:601
		// _ = "end of CoverTab[2661]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:602
		_go_fuzz_dep_.CoverTab[526172]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:602
		_go_fuzz_dep_.CoverTab[2662]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:602
		// _ = "end of CoverTab[2662]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:602
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:602
	// _ = "end of CoverTab[2656]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:602
	_go_fuzz_dep_.CoverTab[2657]++
											var hdr ResourceHeader
											off, err := hdr.unpack(p.msg, p.off)
											if err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:605
		_go_fuzz_dep_.CoverTab[526173]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:605
		_go_fuzz_dep_.CoverTab[2663]++
												return ResourceHeader{}, err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:606
		// _ = "end of CoverTab[2663]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:607
		_go_fuzz_dep_.CoverTab[526174]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:607
		_go_fuzz_dep_.CoverTab[2664]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:607
		// _ = "end of CoverTab[2664]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:607
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:607
	// _ = "end of CoverTab[2657]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:607
	_go_fuzz_dep_.CoverTab[2658]++
											p.resHeaderValid = true
											p.resHeader = hdr
											p.off = off
											return hdr, nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:611
	// _ = "end of CoverTab[2658]"
}

func (p *Parser) skipResource(sec section) error {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:614
	_go_fuzz_dep_.CoverTab[2665]++
											if p.resHeaderValid {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:615
		_go_fuzz_dep_.CoverTab[526175]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:615
		_go_fuzz_dep_.CoverTab[2669]++
												newOff := p.off + int(p.resHeader.Length)
												if newOff > len(p.msg) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:617
			_go_fuzz_dep_.CoverTab[526177]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:617
			_go_fuzz_dep_.CoverTab[2671]++
													return errResourceLen
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:618
			// _ = "end of CoverTab[2671]"
		} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:619
			_go_fuzz_dep_.CoverTab[526178]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:619
			_go_fuzz_dep_.CoverTab[2672]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:619
			// _ = "end of CoverTab[2672]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:619
		}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:619
		// _ = "end of CoverTab[2669]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:619
		_go_fuzz_dep_.CoverTab[2670]++
												p.off = newOff
												p.resHeaderValid = false
												p.index++
												return nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:623
		// _ = "end of CoverTab[2670]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:624
		_go_fuzz_dep_.CoverTab[526176]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:624
		_go_fuzz_dep_.CoverTab[2673]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:624
		// _ = "end of CoverTab[2673]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:624
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:624
	// _ = "end of CoverTab[2665]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:624
	_go_fuzz_dep_.CoverTab[2666]++
											if err := p.checkAdvance(sec); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:625
		_go_fuzz_dep_.CoverTab[526179]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:625
		_go_fuzz_dep_.CoverTab[2674]++
												return err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:626
		// _ = "end of CoverTab[2674]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:627
		_go_fuzz_dep_.CoverTab[526180]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:627
		_go_fuzz_dep_.CoverTab[2675]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:627
		// _ = "end of CoverTab[2675]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:627
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:627
	// _ = "end of CoverTab[2666]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:627
	_go_fuzz_dep_.CoverTab[2667]++
											var err error
											p.off, err = skipResource(p.msg, p.off)
											if err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:630
		_go_fuzz_dep_.CoverTab[526181]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:630
		_go_fuzz_dep_.CoverTab[2676]++
												return &nestedError{"skipping: " + sectionNames[sec], err}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:631
		// _ = "end of CoverTab[2676]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:632
		_go_fuzz_dep_.CoverTab[526182]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:632
		_go_fuzz_dep_.CoverTab[2677]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:632
		// _ = "end of CoverTab[2677]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:632
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:632
	// _ = "end of CoverTab[2667]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:632
	_go_fuzz_dep_.CoverTab[2668]++
											p.index++
											return nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:634
	// _ = "end of CoverTab[2668]"
}

// Question parses a single Question.
func (p *Parser) Question() (Question, error) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:638
	_go_fuzz_dep_.CoverTab[2678]++
											if err := p.checkAdvance(sectionQuestions); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:639
		_go_fuzz_dep_.CoverTab[526183]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:639
		_go_fuzz_dep_.CoverTab[2683]++
												return Question{}, err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:640
		// _ = "end of CoverTab[2683]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:641
		_go_fuzz_dep_.CoverTab[526184]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:641
		_go_fuzz_dep_.CoverTab[2684]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:641
		// _ = "end of CoverTab[2684]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:641
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:641
	// _ = "end of CoverTab[2678]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:641
	_go_fuzz_dep_.CoverTab[2679]++
											var name Name
											off, err := name.unpack(p.msg, p.off)
											if err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:644
		_go_fuzz_dep_.CoverTab[526185]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:644
		_go_fuzz_dep_.CoverTab[2685]++
												return Question{}, &nestedError{"unpacking Question.Name", err}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:645
		// _ = "end of CoverTab[2685]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:646
		_go_fuzz_dep_.CoverTab[526186]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:646
		_go_fuzz_dep_.CoverTab[2686]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:646
		// _ = "end of CoverTab[2686]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:646
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:646
	// _ = "end of CoverTab[2679]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:646
	_go_fuzz_dep_.CoverTab[2680]++
											typ, off, err := unpackType(p.msg, off)
											if err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:648
		_go_fuzz_dep_.CoverTab[526187]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:648
		_go_fuzz_dep_.CoverTab[2687]++
												return Question{}, &nestedError{"unpacking Question.Type", err}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:649
		// _ = "end of CoverTab[2687]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:650
		_go_fuzz_dep_.CoverTab[526188]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:650
		_go_fuzz_dep_.CoverTab[2688]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:650
		// _ = "end of CoverTab[2688]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:650
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:650
	// _ = "end of CoverTab[2680]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:650
	_go_fuzz_dep_.CoverTab[2681]++
											class, off, err := unpackClass(p.msg, off)
											if err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:652
		_go_fuzz_dep_.CoverTab[526189]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:652
		_go_fuzz_dep_.CoverTab[2689]++
												return Question{}, &nestedError{"unpacking Question.Class", err}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:653
		// _ = "end of CoverTab[2689]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:654
		_go_fuzz_dep_.CoverTab[526190]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:654
		_go_fuzz_dep_.CoverTab[2690]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:654
		// _ = "end of CoverTab[2690]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:654
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:654
	// _ = "end of CoverTab[2681]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:654
	_go_fuzz_dep_.CoverTab[2682]++
											p.off = off
											p.index++
											return Question{name, typ, class}, nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:657
	// _ = "end of CoverTab[2682]"
}

// AllQuestions parses all Questions.
func (p *Parser) AllQuestions() ([]Question, error) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:661
	_go_fuzz_dep_.CoverTab[2691]++

//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:668
	qs := []Question{}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:668
	_go_fuzz_dep_.CoverTab[786577] = 0
											for {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:669
		if _go_fuzz_dep_.CoverTab[786577] == 0 {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:669
			_go_fuzz_dep_.CoverTab[526651]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:669
		} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:669
			_go_fuzz_dep_.CoverTab[526652]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:669
		}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:669
		_go_fuzz_dep_.CoverTab[786577] = 1
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:669
		_go_fuzz_dep_.CoverTab[2692]++
												q, err := p.Question()
												if err == ErrSectionDone {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:671
			_go_fuzz_dep_.CoverTab[526191]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:671
			_go_fuzz_dep_.CoverTab[2695]++
													return qs, nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:672
			// _ = "end of CoverTab[2695]"
		} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:673
			_go_fuzz_dep_.CoverTab[526192]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:673
			_go_fuzz_dep_.CoverTab[2696]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:673
			// _ = "end of CoverTab[2696]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:673
		}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:673
		// _ = "end of CoverTab[2692]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:673
		_go_fuzz_dep_.CoverTab[2693]++
												if err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:674
			_go_fuzz_dep_.CoverTab[526193]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:674
			_go_fuzz_dep_.CoverTab[2697]++
													return nil, err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:675
			// _ = "end of CoverTab[2697]"
		} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:676
			_go_fuzz_dep_.CoverTab[526194]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:676
			_go_fuzz_dep_.CoverTab[2698]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:676
			// _ = "end of CoverTab[2698]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:676
		}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:676
		// _ = "end of CoverTab[2693]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:676
		_go_fuzz_dep_.CoverTab[2694]++
												qs = append(qs, q)
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:677
		// _ = "end of CoverTab[2694]"
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:678
	// _ = "end of CoverTab[2691]"
}

// SkipQuestion skips a single Question.
func (p *Parser) SkipQuestion() error {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:682
	_go_fuzz_dep_.CoverTab[2699]++
											if err := p.checkAdvance(sectionQuestions); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:683
		_go_fuzz_dep_.CoverTab[526195]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:683
		_go_fuzz_dep_.CoverTab[2704]++
												return err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:684
		// _ = "end of CoverTab[2704]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:685
		_go_fuzz_dep_.CoverTab[526196]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:685
		_go_fuzz_dep_.CoverTab[2705]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:685
		// _ = "end of CoverTab[2705]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:685
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:685
	// _ = "end of CoverTab[2699]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:685
	_go_fuzz_dep_.CoverTab[2700]++
											off, err := skipName(p.msg, p.off)
											if err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:687
		_go_fuzz_dep_.CoverTab[526197]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:687
		_go_fuzz_dep_.CoverTab[2706]++
												return &nestedError{"skipping Question Name", err}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:688
		// _ = "end of CoverTab[2706]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:689
		_go_fuzz_dep_.CoverTab[526198]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:689
		_go_fuzz_dep_.CoverTab[2707]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:689
		// _ = "end of CoverTab[2707]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:689
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:689
	// _ = "end of CoverTab[2700]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:689
	_go_fuzz_dep_.CoverTab[2701]++
											if off, err = skipType(p.msg, off); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:690
		_go_fuzz_dep_.CoverTab[526199]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:690
		_go_fuzz_dep_.CoverTab[2708]++
												return &nestedError{"skipping Question Type", err}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:691
		// _ = "end of CoverTab[2708]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:692
		_go_fuzz_dep_.CoverTab[526200]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:692
		_go_fuzz_dep_.CoverTab[2709]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:692
		// _ = "end of CoverTab[2709]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:692
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:692
	// _ = "end of CoverTab[2701]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:692
	_go_fuzz_dep_.CoverTab[2702]++
											if off, err = skipClass(p.msg, off); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:693
		_go_fuzz_dep_.CoverTab[526201]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:693
		_go_fuzz_dep_.CoverTab[2710]++
												return &nestedError{"skipping Question Class", err}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:694
		// _ = "end of CoverTab[2710]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:695
		_go_fuzz_dep_.CoverTab[526202]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:695
		_go_fuzz_dep_.CoverTab[2711]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:695
		// _ = "end of CoverTab[2711]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:695
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:695
	// _ = "end of CoverTab[2702]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:695
	_go_fuzz_dep_.CoverTab[2703]++
											p.off = off
											p.index++
											return nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:698
	// _ = "end of CoverTab[2703]"
}

// SkipAllQuestions skips all Questions.
func (p *Parser) SkipAllQuestions() error {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:702
	_go_fuzz_dep_.CoverTab[2712]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:702
	_go_fuzz_dep_.CoverTab[786578] = 0
											for {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:703
		if _go_fuzz_dep_.CoverTab[786578] == 0 {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:703
			_go_fuzz_dep_.CoverTab[526655]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:703
		} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:703
			_go_fuzz_dep_.CoverTab[526656]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:703
		}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:703
		_go_fuzz_dep_.CoverTab[786578] = 1
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:703
		_go_fuzz_dep_.CoverTab[2713]++
												if err := p.SkipQuestion(); err == ErrSectionDone {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:704
			_go_fuzz_dep_.CoverTab[526203]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:704
			_go_fuzz_dep_.CoverTab[2714]++
													return nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:705
			// _ = "end of CoverTab[2714]"
		} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:706
			_go_fuzz_dep_.CoverTab[526204]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:706
			_go_fuzz_dep_.CoverTab[2715]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:706
			if err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:706
				_go_fuzz_dep_.CoverTab[526205]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:706
				_go_fuzz_dep_.CoverTab[2716]++
														return err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:707
				// _ = "end of CoverTab[2716]"
			} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:708
				_go_fuzz_dep_.CoverTab[526206]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:708
				_go_fuzz_dep_.CoverTab[2717]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:708
				// _ = "end of CoverTab[2717]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:708
			}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:708
			// _ = "end of CoverTab[2715]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:708
		}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:708
		// _ = "end of CoverTab[2713]"
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:709
	// _ = "end of CoverTab[2712]"
}

// AnswerHeader parses a single Answer ResourceHeader.
func (p *Parser) AnswerHeader() (ResourceHeader, error) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:713
	_go_fuzz_dep_.CoverTab[2718]++
											return p.resourceHeader(sectionAnswers)
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:714
	// _ = "end of CoverTab[2718]"
}

// Answer parses a single Answer Resource.
func (p *Parser) Answer() (Resource, error) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:718
	_go_fuzz_dep_.CoverTab[2719]++
											return p.resource(sectionAnswers)
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:719
	// _ = "end of CoverTab[2719]"
}

// AllAnswers parses all Answer Resources.
func (p *Parser) AllAnswers() ([]Resource, error) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:723
	_go_fuzz_dep_.CoverTab[2720]++

//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:729
	n := int(p.header.answers)
	if n > 20 {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:730
		_go_fuzz_dep_.CoverTab[526207]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:730
		_go_fuzz_dep_.CoverTab[2722]++
												n = 20
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:731
		// _ = "end of CoverTab[2722]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:732
		_go_fuzz_dep_.CoverTab[526208]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:732
		_go_fuzz_dep_.CoverTab[2723]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:732
		// _ = "end of CoverTab[2723]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:732
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:732
	// _ = "end of CoverTab[2720]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:732
	_go_fuzz_dep_.CoverTab[2721]++
											as := make([]Resource, 0, n)
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:733
	_go_fuzz_dep_.CoverTab[786579] = 0
											for {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:734
		if _go_fuzz_dep_.CoverTab[786579] == 0 {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:734
			_go_fuzz_dep_.CoverTab[526659]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:734
		} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:734
			_go_fuzz_dep_.CoverTab[526660]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:734
		}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:734
		_go_fuzz_dep_.CoverTab[786579] = 1
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:734
		_go_fuzz_dep_.CoverTab[2724]++
												a, err := p.Answer()
												if err == ErrSectionDone {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:736
			_go_fuzz_dep_.CoverTab[526209]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:736
			_go_fuzz_dep_.CoverTab[2727]++
													return as, nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:737
			// _ = "end of CoverTab[2727]"
		} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:738
			_go_fuzz_dep_.CoverTab[526210]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:738
			_go_fuzz_dep_.CoverTab[2728]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:738
			// _ = "end of CoverTab[2728]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:738
		}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:738
		// _ = "end of CoverTab[2724]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:738
		_go_fuzz_dep_.CoverTab[2725]++
												if err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:739
			_go_fuzz_dep_.CoverTab[526211]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:739
			_go_fuzz_dep_.CoverTab[2729]++
													return nil, err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:740
			// _ = "end of CoverTab[2729]"
		} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:741
			_go_fuzz_dep_.CoverTab[526212]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:741
			_go_fuzz_dep_.CoverTab[2730]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:741
			// _ = "end of CoverTab[2730]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:741
		}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:741
		// _ = "end of CoverTab[2725]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:741
		_go_fuzz_dep_.CoverTab[2726]++
												as = append(as, a)
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:742
		// _ = "end of CoverTab[2726]"
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:743
	// _ = "end of CoverTab[2721]"
}

// SkipAnswer skips a single Answer Resource.
func (p *Parser) SkipAnswer() error {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:747
	_go_fuzz_dep_.CoverTab[2731]++
											return p.skipResource(sectionAnswers)
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:748
	// _ = "end of CoverTab[2731]"
}

// SkipAllAnswers skips all Answer Resources.
func (p *Parser) SkipAllAnswers() error {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:752
	_go_fuzz_dep_.CoverTab[2732]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:752
	_go_fuzz_dep_.CoverTab[786580] = 0
											for {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:753
		if _go_fuzz_dep_.CoverTab[786580] == 0 {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:753
			_go_fuzz_dep_.CoverTab[526663]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:753
		} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:753
			_go_fuzz_dep_.CoverTab[526664]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:753
		}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:753
		_go_fuzz_dep_.CoverTab[786580] = 1
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:753
		_go_fuzz_dep_.CoverTab[2733]++
												if err := p.SkipAnswer(); err == ErrSectionDone {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:754
			_go_fuzz_dep_.CoverTab[526213]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:754
			_go_fuzz_dep_.CoverTab[2734]++
													return nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:755
			// _ = "end of CoverTab[2734]"
		} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:756
			_go_fuzz_dep_.CoverTab[526214]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:756
			_go_fuzz_dep_.CoverTab[2735]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:756
			if err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:756
				_go_fuzz_dep_.CoverTab[526215]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:756
				_go_fuzz_dep_.CoverTab[2736]++
														return err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:757
				// _ = "end of CoverTab[2736]"
			} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:758
				_go_fuzz_dep_.CoverTab[526216]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:758
				_go_fuzz_dep_.CoverTab[2737]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:758
				// _ = "end of CoverTab[2737]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:758
			}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:758
			// _ = "end of CoverTab[2735]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:758
		}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:758
		// _ = "end of CoverTab[2733]"
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:759
	// _ = "end of CoverTab[2732]"
}

// AuthorityHeader parses a single Authority ResourceHeader.
func (p *Parser) AuthorityHeader() (ResourceHeader, error) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:763
	_go_fuzz_dep_.CoverTab[2738]++
											return p.resourceHeader(sectionAuthorities)
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:764
	// _ = "end of CoverTab[2738]"
}

// Authority parses a single Authority Resource.
func (p *Parser) Authority() (Resource, error) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:768
	_go_fuzz_dep_.CoverTab[2739]++
											return p.resource(sectionAuthorities)
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:769
	// _ = "end of CoverTab[2739]"
}

// AllAuthorities parses all Authority Resources.
func (p *Parser) AllAuthorities() ([]Resource, error) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:773
	_go_fuzz_dep_.CoverTab[2740]++

//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:779
	n := int(p.header.authorities)
	if n > 10 {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:780
		_go_fuzz_dep_.CoverTab[526217]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:780
		_go_fuzz_dep_.CoverTab[2742]++
												n = 10
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:781
		// _ = "end of CoverTab[2742]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:782
		_go_fuzz_dep_.CoverTab[526218]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:782
		_go_fuzz_dep_.CoverTab[2743]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:782
		// _ = "end of CoverTab[2743]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:782
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:782
	// _ = "end of CoverTab[2740]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:782
	_go_fuzz_dep_.CoverTab[2741]++
											as := make([]Resource, 0, n)
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:783
	_go_fuzz_dep_.CoverTab[786581] = 0
											for {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:784
		if _go_fuzz_dep_.CoverTab[786581] == 0 {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:784
			_go_fuzz_dep_.CoverTab[526667]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:784
		} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:784
			_go_fuzz_dep_.CoverTab[526668]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:784
		}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:784
		_go_fuzz_dep_.CoverTab[786581] = 1
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:784
		_go_fuzz_dep_.CoverTab[2744]++
												a, err := p.Authority()
												if err == ErrSectionDone {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:786
			_go_fuzz_dep_.CoverTab[526219]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:786
			_go_fuzz_dep_.CoverTab[2747]++
													return as, nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:787
			// _ = "end of CoverTab[2747]"
		} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:788
			_go_fuzz_dep_.CoverTab[526220]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:788
			_go_fuzz_dep_.CoverTab[2748]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:788
			// _ = "end of CoverTab[2748]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:788
		}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:788
		// _ = "end of CoverTab[2744]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:788
		_go_fuzz_dep_.CoverTab[2745]++
												if err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:789
			_go_fuzz_dep_.CoverTab[526221]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:789
			_go_fuzz_dep_.CoverTab[2749]++
													return nil, err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:790
			// _ = "end of CoverTab[2749]"
		} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:791
			_go_fuzz_dep_.CoverTab[526222]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:791
			_go_fuzz_dep_.CoverTab[2750]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:791
			// _ = "end of CoverTab[2750]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:791
		}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:791
		// _ = "end of CoverTab[2745]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:791
		_go_fuzz_dep_.CoverTab[2746]++
												as = append(as, a)
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:792
		// _ = "end of CoverTab[2746]"
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:793
	// _ = "end of CoverTab[2741]"
}

// SkipAuthority skips a single Authority Resource.
func (p *Parser) SkipAuthority() error {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:797
	_go_fuzz_dep_.CoverTab[2751]++
											return p.skipResource(sectionAuthorities)
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:798
	// _ = "end of CoverTab[2751]"
}

// SkipAllAuthorities skips all Authority Resources.
func (p *Parser) SkipAllAuthorities() error {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:802
	_go_fuzz_dep_.CoverTab[2752]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:802
	_go_fuzz_dep_.CoverTab[786582] = 0
											for {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:803
		if _go_fuzz_dep_.CoverTab[786582] == 0 {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:803
			_go_fuzz_dep_.CoverTab[526671]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:803
		} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:803
			_go_fuzz_dep_.CoverTab[526672]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:803
		}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:803
		_go_fuzz_dep_.CoverTab[786582] = 1
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:803
		_go_fuzz_dep_.CoverTab[2753]++
												if err := p.SkipAuthority(); err == ErrSectionDone {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:804
			_go_fuzz_dep_.CoverTab[526223]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:804
			_go_fuzz_dep_.CoverTab[2754]++
													return nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:805
			// _ = "end of CoverTab[2754]"
		} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:806
			_go_fuzz_dep_.CoverTab[526224]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:806
			_go_fuzz_dep_.CoverTab[2755]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:806
			if err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:806
				_go_fuzz_dep_.CoverTab[526225]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:806
				_go_fuzz_dep_.CoverTab[2756]++
														return err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:807
				// _ = "end of CoverTab[2756]"
			} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:808
				_go_fuzz_dep_.CoverTab[526226]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:808
				_go_fuzz_dep_.CoverTab[2757]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:808
				// _ = "end of CoverTab[2757]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:808
			}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:808
			// _ = "end of CoverTab[2755]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:808
		}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:808
		// _ = "end of CoverTab[2753]"
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:809
	// _ = "end of CoverTab[2752]"
}

// AdditionalHeader parses a single Additional ResourceHeader.
func (p *Parser) AdditionalHeader() (ResourceHeader, error) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:813
	_go_fuzz_dep_.CoverTab[2758]++
											return p.resourceHeader(sectionAdditionals)
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:814
	// _ = "end of CoverTab[2758]"
}

// Additional parses a single Additional Resource.
func (p *Parser) Additional() (Resource, error) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:818
	_go_fuzz_dep_.CoverTab[2759]++
											return p.resource(sectionAdditionals)
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:819
	// _ = "end of CoverTab[2759]"
}

// AllAdditionals parses all Additional Resources.
func (p *Parser) AllAdditionals() ([]Resource, error) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:823
	_go_fuzz_dep_.CoverTab[2760]++

//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:829
	n := int(p.header.additionals)
	if n > 10 {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:830
		_go_fuzz_dep_.CoverTab[526227]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:830
		_go_fuzz_dep_.CoverTab[2762]++
												n = 10
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:831
		// _ = "end of CoverTab[2762]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:832
		_go_fuzz_dep_.CoverTab[526228]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:832
		_go_fuzz_dep_.CoverTab[2763]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:832
		// _ = "end of CoverTab[2763]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:832
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:832
	// _ = "end of CoverTab[2760]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:832
	_go_fuzz_dep_.CoverTab[2761]++
											as := make([]Resource, 0, n)
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:833
	_go_fuzz_dep_.CoverTab[786583] = 0
											for {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:834
		if _go_fuzz_dep_.CoverTab[786583] == 0 {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:834
			_go_fuzz_dep_.CoverTab[526675]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:834
		} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:834
			_go_fuzz_dep_.CoverTab[526676]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:834
		}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:834
		_go_fuzz_dep_.CoverTab[786583] = 1
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:834
		_go_fuzz_dep_.CoverTab[2764]++
												a, err := p.Additional()
												if err == ErrSectionDone {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:836
			_go_fuzz_dep_.CoverTab[526229]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:836
			_go_fuzz_dep_.CoverTab[2767]++
													return as, nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:837
			// _ = "end of CoverTab[2767]"
		} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:838
			_go_fuzz_dep_.CoverTab[526230]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:838
			_go_fuzz_dep_.CoverTab[2768]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:838
			// _ = "end of CoverTab[2768]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:838
		}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:838
		// _ = "end of CoverTab[2764]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:838
		_go_fuzz_dep_.CoverTab[2765]++
												if err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:839
			_go_fuzz_dep_.CoverTab[526231]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:839
			_go_fuzz_dep_.CoverTab[2769]++
													return nil, err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:840
			// _ = "end of CoverTab[2769]"
		} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:841
			_go_fuzz_dep_.CoverTab[526232]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:841
			_go_fuzz_dep_.CoverTab[2770]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:841
			// _ = "end of CoverTab[2770]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:841
		}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:841
		// _ = "end of CoverTab[2765]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:841
		_go_fuzz_dep_.CoverTab[2766]++
												as = append(as, a)
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:842
		// _ = "end of CoverTab[2766]"
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:843
	// _ = "end of CoverTab[2761]"
}

// SkipAdditional skips a single Additional Resource.
func (p *Parser) SkipAdditional() error {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:847
	_go_fuzz_dep_.CoverTab[2771]++
											return p.skipResource(sectionAdditionals)
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:848
	// _ = "end of CoverTab[2771]"
}

// SkipAllAdditionals skips all Additional Resources.
func (p *Parser) SkipAllAdditionals() error {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:852
	_go_fuzz_dep_.CoverTab[2772]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:852
	_go_fuzz_dep_.CoverTab[786584] = 0
											for {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:853
		if _go_fuzz_dep_.CoverTab[786584] == 0 {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:853
			_go_fuzz_dep_.CoverTab[526679]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:853
		} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:853
			_go_fuzz_dep_.CoverTab[526680]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:853
		}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:853
		_go_fuzz_dep_.CoverTab[786584] = 1
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:853
		_go_fuzz_dep_.CoverTab[2773]++
												if err := p.SkipAdditional(); err == ErrSectionDone {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:854
			_go_fuzz_dep_.CoverTab[526233]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:854
			_go_fuzz_dep_.CoverTab[2774]++
													return nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:855
			// _ = "end of CoverTab[2774]"
		} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:856
			_go_fuzz_dep_.CoverTab[526234]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:856
			_go_fuzz_dep_.CoverTab[2775]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:856
			if err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:856
				_go_fuzz_dep_.CoverTab[526235]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:856
				_go_fuzz_dep_.CoverTab[2776]++
														return err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:857
				// _ = "end of CoverTab[2776]"
			} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:858
				_go_fuzz_dep_.CoverTab[526236]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:858
				_go_fuzz_dep_.CoverTab[2777]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:858
				// _ = "end of CoverTab[2777]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:858
			}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:858
			// _ = "end of CoverTab[2775]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:858
		}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:858
		// _ = "end of CoverTab[2773]"
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:859
	// _ = "end of CoverTab[2772]"
}

// CNAMEResource parses a single CNAMEResource.
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:862
//
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:862
// One of the XXXHeader methods must have been called before calling this
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:862
// method.
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:866
func (p *Parser) CNAMEResource() (CNAMEResource, error) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:866
	_go_fuzz_dep_.CoverTab[2778]++
											if !p.resHeaderValid || func() bool {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:867
		_go_fuzz_dep_.CoverTab[2781]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:867
		return p.resHeader.Type != TypeCNAME
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:867
		// _ = "end of CoverTab[2781]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:867
	}() {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:867
		_go_fuzz_dep_.CoverTab[526237]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:867
		_go_fuzz_dep_.CoverTab[2782]++
												return CNAMEResource{}, ErrNotStarted
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:868
		// _ = "end of CoverTab[2782]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:869
		_go_fuzz_dep_.CoverTab[526238]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:869
		_go_fuzz_dep_.CoverTab[2783]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:869
		// _ = "end of CoverTab[2783]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:869
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:869
	// _ = "end of CoverTab[2778]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:869
	_go_fuzz_dep_.CoverTab[2779]++
											r, err := unpackCNAMEResource(p.msg, p.off)
											if err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:871
		_go_fuzz_dep_.CoverTab[526239]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:871
		_go_fuzz_dep_.CoverTab[2784]++
												return CNAMEResource{}, err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:872
		// _ = "end of CoverTab[2784]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:873
		_go_fuzz_dep_.CoverTab[526240]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:873
		_go_fuzz_dep_.CoverTab[2785]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:873
		// _ = "end of CoverTab[2785]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:873
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:873
	// _ = "end of CoverTab[2779]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:873
	_go_fuzz_dep_.CoverTab[2780]++
											p.off += int(p.resHeader.Length)
											p.resHeaderValid = false
											p.index++
											return r, nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:877
	// _ = "end of CoverTab[2780]"
}

// MXResource parses a single MXResource.
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:880
//
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:880
// One of the XXXHeader methods must have been called before calling this
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:880
// method.
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:884
func (p *Parser) MXResource() (MXResource, error) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:884
	_go_fuzz_dep_.CoverTab[2786]++
											if !p.resHeaderValid || func() bool {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:885
		_go_fuzz_dep_.CoverTab[2789]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:885
		return p.resHeader.Type != TypeMX
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:885
		// _ = "end of CoverTab[2789]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:885
	}() {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:885
		_go_fuzz_dep_.CoverTab[526241]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:885
		_go_fuzz_dep_.CoverTab[2790]++
												return MXResource{}, ErrNotStarted
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:886
		// _ = "end of CoverTab[2790]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:887
		_go_fuzz_dep_.CoverTab[526242]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:887
		_go_fuzz_dep_.CoverTab[2791]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:887
		// _ = "end of CoverTab[2791]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:887
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:887
	// _ = "end of CoverTab[2786]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:887
	_go_fuzz_dep_.CoverTab[2787]++
											r, err := unpackMXResource(p.msg, p.off)
											if err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:889
		_go_fuzz_dep_.CoverTab[526243]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:889
		_go_fuzz_dep_.CoverTab[2792]++
												return MXResource{}, err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:890
		// _ = "end of CoverTab[2792]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:891
		_go_fuzz_dep_.CoverTab[526244]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:891
		_go_fuzz_dep_.CoverTab[2793]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:891
		// _ = "end of CoverTab[2793]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:891
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:891
	// _ = "end of CoverTab[2787]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:891
	_go_fuzz_dep_.CoverTab[2788]++
											p.off += int(p.resHeader.Length)
											p.resHeaderValid = false
											p.index++
											return r, nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:895
	// _ = "end of CoverTab[2788]"
}

// NSResource parses a single NSResource.
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:898
//
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:898
// One of the XXXHeader methods must have been called before calling this
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:898
// method.
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:902
func (p *Parser) NSResource() (NSResource, error) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:902
	_go_fuzz_dep_.CoverTab[2794]++
											if !p.resHeaderValid || func() bool {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:903
		_go_fuzz_dep_.CoverTab[2797]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:903
		return p.resHeader.Type != TypeNS
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:903
		// _ = "end of CoverTab[2797]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:903
	}() {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:903
		_go_fuzz_dep_.CoverTab[526245]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:903
		_go_fuzz_dep_.CoverTab[2798]++
												return NSResource{}, ErrNotStarted
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:904
		// _ = "end of CoverTab[2798]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:905
		_go_fuzz_dep_.CoverTab[526246]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:905
		_go_fuzz_dep_.CoverTab[2799]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:905
		// _ = "end of CoverTab[2799]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:905
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:905
	// _ = "end of CoverTab[2794]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:905
	_go_fuzz_dep_.CoverTab[2795]++
											r, err := unpackNSResource(p.msg, p.off)
											if err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:907
		_go_fuzz_dep_.CoverTab[526247]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:907
		_go_fuzz_dep_.CoverTab[2800]++
												return NSResource{}, err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:908
		// _ = "end of CoverTab[2800]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:909
		_go_fuzz_dep_.CoverTab[526248]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:909
		_go_fuzz_dep_.CoverTab[2801]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:909
		// _ = "end of CoverTab[2801]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:909
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:909
	// _ = "end of CoverTab[2795]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:909
	_go_fuzz_dep_.CoverTab[2796]++
											p.off += int(p.resHeader.Length)
											p.resHeaderValid = false
											p.index++
											return r, nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:913
	// _ = "end of CoverTab[2796]"
}

// PTRResource parses a single PTRResource.
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:916
//
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:916
// One of the XXXHeader methods must have been called before calling this
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:916
// method.
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:920
func (p *Parser) PTRResource() (PTRResource, error) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:920
	_go_fuzz_dep_.CoverTab[2802]++
											if !p.resHeaderValid || func() bool {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:921
		_go_fuzz_dep_.CoverTab[2805]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:921
		return p.resHeader.Type != TypePTR
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:921
		// _ = "end of CoverTab[2805]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:921
	}() {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:921
		_go_fuzz_dep_.CoverTab[526249]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:921
		_go_fuzz_dep_.CoverTab[2806]++
												return PTRResource{}, ErrNotStarted
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:922
		// _ = "end of CoverTab[2806]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:923
		_go_fuzz_dep_.CoverTab[526250]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:923
		_go_fuzz_dep_.CoverTab[2807]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:923
		// _ = "end of CoverTab[2807]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:923
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:923
	// _ = "end of CoverTab[2802]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:923
	_go_fuzz_dep_.CoverTab[2803]++
											r, err := unpackPTRResource(p.msg, p.off)
											if err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:925
		_go_fuzz_dep_.CoverTab[526251]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:925
		_go_fuzz_dep_.CoverTab[2808]++
												return PTRResource{}, err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:926
		// _ = "end of CoverTab[2808]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:927
		_go_fuzz_dep_.CoverTab[526252]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:927
		_go_fuzz_dep_.CoverTab[2809]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:927
		// _ = "end of CoverTab[2809]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:927
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:927
	// _ = "end of CoverTab[2803]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:927
	_go_fuzz_dep_.CoverTab[2804]++
											p.off += int(p.resHeader.Length)
											p.resHeaderValid = false
											p.index++
											return r, nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:931
	// _ = "end of CoverTab[2804]"
}

// SOAResource parses a single SOAResource.
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:934
//
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:934
// One of the XXXHeader methods must have been called before calling this
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:934
// method.
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:938
func (p *Parser) SOAResource() (SOAResource, error) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:938
	_go_fuzz_dep_.CoverTab[2810]++
											if !p.resHeaderValid || func() bool {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:939
		_go_fuzz_dep_.CoverTab[2813]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:939
		return p.resHeader.Type != TypeSOA
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:939
		// _ = "end of CoverTab[2813]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:939
	}() {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:939
		_go_fuzz_dep_.CoverTab[526253]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:939
		_go_fuzz_dep_.CoverTab[2814]++
												return SOAResource{}, ErrNotStarted
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:940
		// _ = "end of CoverTab[2814]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:941
		_go_fuzz_dep_.CoverTab[526254]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:941
		_go_fuzz_dep_.CoverTab[2815]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:941
		// _ = "end of CoverTab[2815]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:941
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:941
	// _ = "end of CoverTab[2810]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:941
	_go_fuzz_dep_.CoverTab[2811]++
											r, err := unpackSOAResource(p.msg, p.off)
											if err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:943
		_go_fuzz_dep_.CoverTab[526255]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:943
		_go_fuzz_dep_.CoverTab[2816]++
												return SOAResource{}, err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:944
		// _ = "end of CoverTab[2816]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:945
		_go_fuzz_dep_.CoverTab[526256]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:945
		_go_fuzz_dep_.CoverTab[2817]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:945
		// _ = "end of CoverTab[2817]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:945
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:945
	// _ = "end of CoverTab[2811]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:945
	_go_fuzz_dep_.CoverTab[2812]++
											p.off += int(p.resHeader.Length)
											p.resHeaderValid = false
											p.index++
											return r, nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:949
	// _ = "end of CoverTab[2812]"
}

// TXTResource parses a single TXTResource.
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:952
//
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:952
// One of the XXXHeader methods must have been called before calling this
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:952
// method.
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:956
func (p *Parser) TXTResource() (TXTResource, error) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:956
	_go_fuzz_dep_.CoverTab[2818]++
											if !p.resHeaderValid || func() bool {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:957
		_go_fuzz_dep_.CoverTab[2821]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:957
		return p.resHeader.Type != TypeTXT
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:957
		// _ = "end of CoverTab[2821]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:957
	}() {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:957
		_go_fuzz_dep_.CoverTab[526257]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:957
		_go_fuzz_dep_.CoverTab[2822]++
												return TXTResource{}, ErrNotStarted
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:958
		// _ = "end of CoverTab[2822]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:959
		_go_fuzz_dep_.CoverTab[526258]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:959
		_go_fuzz_dep_.CoverTab[2823]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:959
		// _ = "end of CoverTab[2823]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:959
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:959
	// _ = "end of CoverTab[2818]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:959
	_go_fuzz_dep_.CoverTab[2819]++
											r, err := unpackTXTResource(p.msg, p.off, p.resHeader.Length)
											if err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:961
		_go_fuzz_dep_.CoverTab[526259]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:961
		_go_fuzz_dep_.CoverTab[2824]++
												return TXTResource{}, err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:962
		// _ = "end of CoverTab[2824]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:963
		_go_fuzz_dep_.CoverTab[526260]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:963
		_go_fuzz_dep_.CoverTab[2825]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:963
		// _ = "end of CoverTab[2825]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:963
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:963
	// _ = "end of CoverTab[2819]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:963
	_go_fuzz_dep_.CoverTab[2820]++
											p.off += int(p.resHeader.Length)
											p.resHeaderValid = false
											p.index++
											return r, nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:967
	// _ = "end of CoverTab[2820]"
}

// SRVResource parses a single SRVResource.
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:970
//
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:970
// One of the XXXHeader methods must have been called before calling this
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:970
// method.
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:974
func (p *Parser) SRVResource() (SRVResource, error) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:974
	_go_fuzz_dep_.CoverTab[2826]++
											if !p.resHeaderValid || func() bool {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:975
		_go_fuzz_dep_.CoverTab[2829]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:975
		return p.resHeader.Type != TypeSRV
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:975
		// _ = "end of CoverTab[2829]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:975
	}() {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:975
		_go_fuzz_dep_.CoverTab[526261]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:975
		_go_fuzz_dep_.CoverTab[2830]++
												return SRVResource{}, ErrNotStarted
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:976
		// _ = "end of CoverTab[2830]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:977
		_go_fuzz_dep_.CoverTab[526262]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:977
		_go_fuzz_dep_.CoverTab[2831]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:977
		// _ = "end of CoverTab[2831]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:977
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:977
	// _ = "end of CoverTab[2826]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:977
	_go_fuzz_dep_.CoverTab[2827]++
											r, err := unpackSRVResource(p.msg, p.off)
											if err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:979
		_go_fuzz_dep_.CoverTab[526263]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:979
		_go_fuzz_dep_.CoverTab[2832]++
												return SRVResource{}, err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:980
		// _ = "end of CoverTab[2832]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:981
		_go_fuzz_dep_.CoverTab[526264]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:981
		_go_fuzz_dep_.CoverTab[2833]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:981
		// _ = "end of CoverTab[2833]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:981
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:981
	// _ = "end of CoverTab[2827]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:981
	_go_fuzz_dep_.CoverTab[2828]++
											p.off += int(p.resHeader.Length)
											p.resHeaderValid = false
											p.index++
											return r, nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:985
	// _ = "end of CoverTab[2828]"
}

// AResource parses a single AResource.
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:988
//
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:988
// One of the XXXHeader methods must have been called before calling this
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:988
// method.
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:992
func (p *Parser) AResource() (AResource, error) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:992
	_go_fuzz_dep_.CoverTab[2834]++
											if !p.resHeaderValid || func() bool {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:993
		_go_fuzz_dep_.CoverTab[2837]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:993
		return p.resHeader.Type != TypeA
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:993
		// _ = "end of CoverTab[2837]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:993
	}() {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:993
		_go_fuzz_dep_.CoverTab[526265]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:993
		_go_fuzz_dep_.CoverTab[2838]++
												return AResource{}, ErrNotStarted
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:994
		// _ = "end of CoverTab[2838]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:995
		_go_fuzz_dep_.CoverTab[526266]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:995
		_go_fuzz_dep_.CoverTab[2839]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:995
		// _ = "end of CoverTab[2839]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:995
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:995
	// _ = "end of CoverTab[2834]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:995
	_go_fuzz_dep_.CoverTab[2835]++
											r, err := unpackAResource(p.msg, p.off)
											if err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:997
		_go_fuzz_dep_.CoverTab[526267]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:997
		_go_fuzz_dep_.CoverTab[2840]++
												return AResource{}, err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:998
		// _ = "end of CoverTab[2840]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:999
		_go_fuzz_dep_.CoverTab[526268]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:999
		_go_fuzz_dep_.CoverTab[2841]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:999
		// _ = "end of CoverTab[2841]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:999
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:999
	// _ = "end of CoverTab[2835]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:999
	_go_fuzz_dep_.CoverTab[2836]++
											p.off += int(p.resHeader.Length)
											p.resHeaderValid = false
											p.index++
											return r, nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1003
	// _ = "end of CoverTab[2836]"
}

// AAAAResource parses a single AAAAResource.
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1006
//
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1006
// One of the XXXHeader methods must have been called before calling this
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1006
// method.
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1010
func (p *Parser) AAAAResource() (AAAAResource, error) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1010
	_go_fuzz_dep_.CoverTab[2842]++
											if !p.resHeaderValid || func() bool {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1011
		_go_fuzz_dep_.CoverTab[2845]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1011
		return p.resHeader.Type != TypeAAAA
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1011
		// _ = "end of CoverTab[2845]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1011
	}() {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1011
		_go_fuzz_dep_.CoverTab[526269]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1011
		_go_fuzz_dep_.CoverTab[2846]++
												return AAAAResource{}, ErrNotStarted
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1012
		// _ = "end of CoverTab[2846]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1013
		_go_fuzz_dep_.CoverTab[526270]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1013
		_go_fuzz_dep_.CoverTab[2847]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1013
		// _ = "end of CoverTab[2847]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1013
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1013
	// _ = "end of CoverTab[2842]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1013
	_go_fuzz_dep_.CoverTab[2843]++
											r, err := unpackAAAAResource(p.msg, p.off)
											if err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1015
		_go_fuzz_dep_.CoverTab[526271]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1015
		_go_fuzz_dep_.CoverTab[2848]++
												return AAAAResource{}, err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1016
		// _ = "end of CoverTab[2848]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1017
		_go_fuzz_dep_.CoverTab[526272]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1017
		_go_fuzz_dep_.CoverTab[2849]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1017
		// _ = "end of CoverTab[2849]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1017
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1017
	// _ = "end of CoverTab[2843]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1017
	_go_fuzz_dep_.CoverTab[2844]++
											p.off += int(p.resHeader.Length)
											p.resHeaderValid = false
											p.index++
											return r, nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1021
	// _ = "end of CoverTab[2844]"
}

// OPTResource parses a single OPTResource.
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1024
//
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1024
// One of the XXXHeader methods must have been called before calling this
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1024
// method.
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1028
func (p *Parser) OPTResource() (OPTResource, error) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1028
	_go_fuzz_dep_.CoverTab[2850]++
											if !p.resHeaderValid || func() bool {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1029
		_go_fuzz_dep_.CoverTab[2853]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1029
		return p.resHeader.Type != TypeOPT
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1029
		// _ = "end of CoverTab[2853]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1029
	}() {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1029
		_go_fuzz_dep_.CoverTab[526273]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1029
		_go_fuzz_dep_.CoverTab[2854]++
												return OPTResource{}, ErrNotStarted
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1030
		// _ = "end of CoverTab[2854]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1031
		_go_fuzz_dep_.CoverTab[526274]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1031
		_go_fuzz_dep_.CoverTab[2855]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1031
		// _ = "end of CoverTab[2855]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1031
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1031
	// _ = "end of CoverTab[2850]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1031
	_go_fuzz_dep_.CoverTab[2851]++
											r, err := unpackOPTResource(p.msg, p.off, p.resHeader.Length)
											if err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1033
		_go_fuzz_dep_.CoverTab[526275]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1033
		_go_fuzz_dep_.CoverTab[2856]++
												return OPTResource{}, err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1034
		// _ = "end of CoverTab[2856]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1035
		_go_fuzz_dep_.CoverTab[526276]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1035
		_go_fuzz_dep_.CoverTab[2857]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1035
		// _ = "end of CoverTab[2857]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1035
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1035
	// _ = "end of CoverTab[2851]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1035
	_go_fuzz_dep_.CoverTab[2852]++
											p.off += int(p.resHeader.Length)
											p.resHeaderValid = false
											p.index++
											return r, nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1039
	// _ = "end of CoverTab[2852]"
}

// UnknownResource parses a single UnknownResource.
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1042
//
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1042
// One of the XXXHeader methods must have been called before calling this
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1042
// method.
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1046
func (p *Parser) UnknownResource() (UnknownResource, error) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1046
	_go_fuzz_dep_.CoverTab[2858]++
											if !p.resHeaderValid {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1047
		_go_fuzz_dep_.CoverTab[526277]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1047
		_go_fuzz_dep_.CoverTab[2861]++
												return UnknownResource{}, ErrNotStarted
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1048
		// _ = "end of CoverTab[2861]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1049
		_go_fuzz_dep_.CoverTab[526278]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1049
		_go_fuzz_dep_.CoverTab[2862]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1049
		// _ = "end of CoverTab[2862]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1049
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1049
	// _ = "end of CoverTab[2858]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1049
	_go_fuzz_dep_.CoverTab[2859]++
											r, err := unpackUnknownResource(p.resHeader.Type, p.msg, p.off, p.resHeader.Length)
											if err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1051
		_go_fuzz_dep_.CoverTab[526279]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1051
		_go_fuzz_dep_.CoverTab[2863]++
												return UnknownResource{}, err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1052
		// _ = "end of CoverTab[2863]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1053
		_go_fuzz_dep_.CoverTab[526280]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1053
		_go_fuzz_dep_.CoverTab[2864]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1053
		// _ = "end of CoverTab[2864]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1053
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1053
	// _ = "end of CoverTab[2859]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1053
	_go_fuzz_dep_.CoverTab[2860]++
											p.off += int(p.resHeader.Length)
											p.resHeaderValid = false
											p.index++
											return r, nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1057
	// _ = "end of CoverTab[2860]"
}

// Unpack parses a full Message.
func (m *Message) Unpack(msg []byte) error {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1061
	_go_fuzz_dep_.CoverTab[2865]++
											var p Parser
											var err error
											if m.Header, err = p.Start(msg); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1064
		_go_fuzz_dep_.CoverTab[526281]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1064
		_go_fuzz_dep_.CoverTab[2871]++
												return err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1065
		// _ = "end of CoverTab[2871]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1066
		_go_fuzz_dep_.CoverTab[526282]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1066
		_go_fuzz_dep_.CoverTab[2872]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1066
		// _ = "end of CoverTab[2872]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1066
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1066
	// _ = "end of CoverTab[2865]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1066
	_go_fuzz_dep_.CoverTab[2866]++
											if m.Questions, err = p.AllQuestions(); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1067
		_go_fuzz_dep_.CoverTab[526283]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1067
		_go_fuzz_dep_.CoverTab[2873]++
												return err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1068
		// _ = "end of CoverTab[2873]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1069
		_go_fuzz_dep_.CoverTab[526284]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1069
		_go_fuzz_dep_.CoverTab[2874]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1069
		// _ = "end of CoverTab[2874]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1069
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1069
	// _ = "end of CoverTab[2866]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1069
	_go_fuzz_dep_.CoverTab[2867]++
											if m.Answers, err = p.AllAnswers(); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1070
		_go_fuzz_dep_.CoverTab[526285]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1070
		_go_fuzz_dep_.CoverTab[2875]++
												return err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1071
		// _ = "end of CoverTab[2875]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1072
		_go_fuzz_dep_.CoverTab[526286]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1072
		_go_fuzz_dep_.CoverTab[2876]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1072
		// _ = "end of CoverTab[2876]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1072
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1072
	// _ = "end of CoverTab[2867]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1072
	_go_fuzz_dep_.CoverTab[2868]++
											if m.Authorities, err = p.AllAuthorities(); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1073
		_go_fuzz_dep_.CoverTab[526287]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1073
		_go_fuzz_dep_.CoverTab[2877]++
												return err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1074
		// _ = "end of CoverTab[2877]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1075
		_go_fuzz_dep_.CoverTab[526288]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1075
		_go_fuzz_dep_.CoverTab[2878]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1075
		// _ = "end of CoverTab[2878]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1075
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1075
	// _ = "end of CoverTab[2868]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1075
	_go_fuzz_dep_.CoverTab[2869]++
											if m.Additionals, err = p.AllAdditionals(); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1076
		_go_fuzz_dep_.CoverTab[526289]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1076
		_go_fuzz_dep_.CoverTab[2879]++
												return err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1077
		// _ = "end of CoverTab[2879]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1078
		_go_fuzz_dep_.CoverTab[526290]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1078
		_go_fuzz_dep_.CoverTab[2880]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1078
		// _ = "end of CoverTab[2880]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1078
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1078
	// _ = "end of CoverTab[2869]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1078
	_go_fuzz_dep_.CoverTab[2870]++
											return nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1079
	// _ = "end of CoverTab[2870]"
}

// Pack packs a full Message.
func (m *Message) Pack() ([]byte, error) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1083
	_go_fuzz_dep_.CoverTab[2881]++
											return m.AppendPack(make([]byte, 0, packStartingCap))
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1084
	// _ = "end of CoverTab[2881]"
}

// AppendPack is like Pack but appends the full Message to b and returns the
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1087
// extended buffer.
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1089
func (m *Message) AppendPack(b []byte) ([]byte, error) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1089
	_go_fuzz_dep_.CoverTab[2882]++

//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1093
	if len(m.Questions) > int(^uint16(0)) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1093
		_go_fuzz_dep_.CoverTab[526291]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1093
		_go_fuzz_dep_.CoverTab[2891]++
												return nil, errTooManyQuestions
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1094
		// _ = "end of CoverTab[2891]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1095
		_go_fuzz_dep_.CoverTab[526292]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1095
		_go_fuzz_dep_.CoverTab[2892]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1095
		// _ = "end of CoverTab[2892]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1095
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1095
	// _ = "end of CoverTab[2882]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1095
	_go_fuzz_dep_.CoverTab[2883]++
											if len(m.Answers) > int(^uint16(0)) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1096
		_go_fuzz_dep_.CoverTab[526293]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1096
		_go_fuzz_dep_.CoverTab[2893]++
												return nil, errTooManyAnswers
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1097
		// _ = "end of CoverTab[2893]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1098
		_go_fuzz_dep_.CoverTab[526294]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1098
		_go_fuzz_dep_.CoverTab[2894]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1098
		// _ = "end of CoverTab[2894]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1098
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1098
	// _ = "end of CoverTab[2883]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1098
	_go_fuzz_dep_.CoverTab[2884]++
											if len(m.Authorities) > int(^uint16(0)) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1099
		_go_fuzz_dep_.CoverTab[526295]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1099
		_go_fuzz_dep_.CoverTab[2895]++
												return nil, errTooManyAuthorities
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1100
		// _ = "end of CoverTab[2895]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1101
		_go_fuzz_dep_.CoverTab[526296]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1101
		_go_fuzz_dep_.CoverTab[2896]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1101
		// _ = "end of CoverTab[2896]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1101
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1101
	// _ = "end of CoverTab[2884]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1101
	_go_fuzz_dep_.CoverTab[2885]++
											if len(m.Additionals) > int(^uint16(0)) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1102
		_go_fuzz_dep_.CoverTab[526297]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1102
		_go_fuzz_dep_.CoverTab[2897]++
												return nil, errTooManyAdditionals
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1103
		// _ = "end of CoverTab[2897]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1104
		_go_fuzz_dep_.CoverTab[526298]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1104
		_go_fuzz_dep_.CoverTab[2898]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1104
		// _ = "end of CoverTab[2898]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1104
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1104
	// _ = "end of CoverTab[2885]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1104
	_go_fuzz_dep_.CoverTab[2886]++

											var h header
											h.id, h.bits = m.Header.pack()

											h.questions = uint16(len(m.Questions))
											h.answers = uint16(len(m.Answers))
											h.authorities = uint16(len(m.Authorities))
											h.additionals = uint16(len(m.Additionals))

											compressionOff := len(b)
											msg := h.pack(b)

//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1125
	compression := map[string]int{}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1125
	_go_fuzz_dep_.CoverTab[786585] = 0

											for i := range m.Questions {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1127
		if _go_fuzz_dep_.CoverTab[786585] == 0 {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1127
			_go_fuzz_dep_.CoverTab[526683]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1127
		} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1127
			_go_fuzz_dep_.CoverTab[526684]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1127
		}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1127
		_go_fuzz_dep_.CoverTab[786585] = 1
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1127
		_go_fuzz_dep_.CoverTab[2899]++
												var err error
												if msg, err = m.Questions[i].pack(msg, compression, compressionOff); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1129
			_go_fuzz_dep_.CoverTab[526299]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1129
			_go_fuzz_dep_.CoverTab[2900]++
													return nil, &nestedError{"packing Question", err}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1130
			// _ = "end of CoverTab[2900]"
		} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1131
			_go_fuzz_dep_.CoverTab[526300]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1131
			_go_fuzz_dep_.CoverTab[2901]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1131
			// _ = "end of CoverTab[2901]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1131
		}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1131
		// _ = "end of CoverTab[2899]"
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1132
	if _go_fuzz_dep_.CoverTab[786585] == 0 {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1132
		_go_fuzz_dep_.CoverTab[526685]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1132
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1132
		_go_fuzz_dep_.CoverTab[526686]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1132
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1132
	// _ = "end of CoverTab[2886]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1132
	_go_fuzz_dep_.CoverTab[2887]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1132
	_go_fuzz_dep_.CoverTab[786586] = 0
											for i := range m.Answers {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1133
		if _go_fuzz_dep_.CoverTab[786586] == 0 {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1133
			_go_fuzz_dep_.CoverTab[526687]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1133
		} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1133
			_go_fuzz_dep_.CoverTab[526688]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1133
		}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1133
		_go_fuzz_dep_.CoverTab[786586] = 1
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1133
		_go_fuzz_dep_.CoverTab[2902]++
												var err error
												if msg, err = m.Answers[i].pack(msg, compression, compressionOff); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1135
			_go_fuzz_dep_.CoverTab[526301]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1135
			_go_fuzz_dep_.CoverTab[2903]++
													return nil, &nestedError{"packing Answer", err}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1136
			// _ = "end of CoverTab[2903]"
		} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1137
			_go_fuzz_dep_.CoverTab[526302]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1137
			_go_fuzz_dep_.CoverTab[2904]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1137
			// _ = "end of CoverTab[2904]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1137
		}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1137
		// _ = "end of CoverTab[2902]"
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1138
	if _go_fuzz_dep_.CoverTab[786586] == 0 {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1138
		_go_fuzz_dep_.CoverTab[526689]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1138
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1138
		_go_fuzz_dep_.CoverTab[526690]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1138
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1138
	// _ = "end of CoverTab[2887]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1138
	_go_fuzz_dep_.CoverTab[2888]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1138
	_go_fuzz_dep_.CoverTab[786587] = 0
											for i := range m.Authorities {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1139
		if _go_fuzz_dep_.CoverTab[786587] == 0 {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1139
			_go_fuzz_dep_.CoverTab[526691]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1139
		} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1139
			_go_fuzz_dep_.CoverTab[526692]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1139
		}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1139
		_go_fuzz_dep_.CoverTab[786587] = 1
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1139
		_go_fuzz_dep_.CoverTab[2905]++
												var err error
												if msg, err = m.Authorities[i].pack(msg, compression, compressionOff); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1141
			_go_fuzz_dep_.CoverTab[526303]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1141
			_go_fuzz_dep_.CoverTab[2906]++
													return nil, &nestedError{"packing Authority", err}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1142
			// _ = "end of CoverTab[2906]"
		} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1143
			_go_fuzz_dep_.CoverTab[526304]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1143
			_go_fuzz_dep_.CoverTab[2907]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1143
			// _ = "end of CoverTab[2907]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1143
		}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1143
		// _ = "end of CoverTab[2905]"
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1144
	if _go_fuzz_dep_.CoverTab[786587] == 0 {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1144
		_go_fuzz_dep_.CoverTab[526693]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1144
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1144
		_go_fuzz_dep_.CoverTab[526694]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1144
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1144
	// _ = "end of CoverTab[2888]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1144
	_go_fuzz_dep_.CoverTab[2889]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1144
	_go_fuzz_dep_.CoverTab[786588] = 0
											for i := range m.Additionals {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1145
		if _go_fuzz_dep_.CoverTab[786588] == 0 {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1145
			_go_fuzz_dep_.CoverTab[526695]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1145
		} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1145
			_go_fuzz_dep_.CoverTab[526696]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1145
		}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1145
		_go_fuzz_dep_.CoverTab[786588] = 1
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1145
		_go_fuzz_dep_.CoverTab[2908]++
												var err error
												if msg, err = m.Additionals[i].pack(msg, compression, compressionOff); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1147
			_go_fuzz_dep_.CoverTab[526305]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1147
			_go_fuzz_dep_.CoverTab[2909]++
													return nil, &nestedError{"packing Additional", err}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1148
			// _ = "end of CoverTab[2909]"
		} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1149
			_go_fuzz_dep_.CoverTab[526306]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1149
			_go_fuzz_dep_.CoverTab[2910]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1149
			// _ = "end of CoverTab[2910]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1149
		}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1149
		// _ = "end of CoverTab[2908]"
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1150
	if _go_fuzz_dep_.CoverTab[786588] == 0 {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1150
		_go_fuzz_dep_.CoverTab[526697]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1150
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1150
		_go_fuzz_dep_.CoverTab[526698]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1150
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1150
	// _ = "end of CoverTab[2889]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1150
	_go_fuzz_dep_.CoverTab[2890]++

											return msg, nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1152
	// _ = "end of CoverTab[2890]"
}

// GoString implements fmt.GoStringer.GoString.
func (m *Message) GoString() string {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1156
	_go_fuzz_dep_.CoverTab[2911]++
											s := "dnsmessage.Message{Header: " + m.Header.GoString() + ", " +
		"Questions: []dnsmessage.Question{"
	if len(m.Questions) > 0 {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1159
		_go_fuzz_dep_.CoverTab[526307]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1159
		_go_fuzz_dep_.CoverTab[2916]++
												s += m.Questions[0].GoString()
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1160
		_go_fuzz_dep_.CoverTab[786589] = 0
												for _, q := range m.Questions[1:] {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1161
			if _go_fuzz_dep_.CoverTab[786589] == 0 {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1161
				_go_fuzz_dep_.CoverTab[526699]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1161
			} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1161
				_go_fuzz_dep_.CoverTab[526700]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1161
			}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1161
			_go_fuzz_dep_.CoverTab[786589] = 1
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1161
			_go_fuzz_dep_.CoverTab[2917]++
													s += ", " + q.GoString()
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1162
			// _ = "end of CoverTab[2917]"
		}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1163
		if _go_fuzz_dep_.CoverTab[786589] == 0 {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1163
			_go_fuzz_dep_.CoverTab[526701]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1163
		} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1163
			_go_fuzz_dep_.CoverTab[526702]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1163
		}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1163
		// _ = "end of CoverTab[2916]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1164
		_go_fuzz_dep_.CoverTab[526308]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1164
		_go_fuzz_dep_.CoverTab[2918]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1164
		// _ = "end of CoverTab[2918]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1164
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1164
	// _ = "end of CoverTab[2911]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1164
	_go_fuzz_dep_.CoverTab[2912]++
											s += "}, Answers: []dnsmessage.Resource{"
											if len(m.Answers) > 0 {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1166
		_go_fuzz_dep_.CoverTab[526309]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1166
		_go_fuzz_dep_.CoverTab[2919]++
												s += m.Answers[0].GoString()
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1167
		_go_fuzz_dep_.CoverTab[786590] = 0
												for _, a := range m.Answers[1:] {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1168
			if _go_fuzz_dep_.CoverTab[786590] == 0 {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1168
				_go_fuzz_dep_.CoverTab[526703]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1168
			} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1168
				_go_fuzz_dep_.CoverTab[526704]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1168
			}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1168
			_go_fuzz_dep_.CoverTab[786590] = 1
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1168
			_go_fuzz_dep_.CoverTab[2920]++
													s += ", " + a.GoString()
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1169
			// _ = "end of CoverTab[2920]"
		}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1170
		if _go_fuzz_dep_.CoverTab[786590] == 0 {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1170
			_go_fuzz_dep_.CoverTab[526705]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1170
		} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1170
			_go_fuzz_dep_.CoverTab[526706]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1170
		}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1170
		// _ = "end of CoverTab[2919]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1171
		_go_fuzz_dep_.CoverTab[526310]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1171
		_go_fuzz_dep_.CoverTab[2921]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1171
		// _ = "end of CoverTab[2921]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1171
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1171
	// _ = "end of CoverTab[2912]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1171
	_go_fuzz_dep_.CoverTab[2913]++
											s += "}, Authorities: []dnsmessage.Resource{"
											if len(m.Authorities) > 0 {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1173
		_go_fuzz_dep_.CoverTab[526311]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1173
		_go_fuzz_dep_.CoverTab[2922]++
												s += m.Authorities[0].GoString()
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1174
		_go_fuzz_dep_.CoverTab[786591] = 0
												for _, a := range m.Authorities[1:] {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1175
			if _go_fuzz_dep_.CoverTab[786591] == 0 {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1175
				_go_fuzz_dep_.CoverTab[526707]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1175
			} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1175
				_go_fuzz_dep_.CoverTab[526708]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1175
			}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1175
			_go_fuzz_dep_.CoverTab[786591] = 1
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1175
			_go_fuzz_dep_.CoverTab[2923]++
													s += ", " + a.GoString()
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1176
			// _ = "end of CoverTab[2923]"
		}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1177
		if _go_fuzz_dep_.CoverTab[786591] == 0 {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1177
			_go_fuzz_dep_.CoverTab[526709]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1177
		} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1177
			_go_fuzz_dep_.CoverTab[526710]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1177
		}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1177
		// _ = "end of CoverTab[2922]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1178
		_go_fuzz_dep_.CoverTab[526312]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1178
		_go_fuzz_dep_.CoverTab[2924]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1178
		// _ = "end of CoverTab[2924]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1178
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1178
	// _ = "end of CoverTab[2913]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1178
	_go_fuzz_dep_.CoverTab[2914]++
											s += "}, Additionals: []dnsmessage.Resource{"
											if len(m.Additionals) > 0 {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1180
		_go_fuzz_dep_.CoverTab[526313]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1180
		_go_fuzz_dep_.CoverTab[2925]++
												s += m.Additionals[0].GoString()
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1181
		_go_fuzz_dep_.CoverTab[786592] = 0
												for _, a := range m.Additionals[1:] {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1182
			if _go_fuzz_dep_.CoverTab[786592] == 0 {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1182
				_go_fuzz_dep_.CoverTab[526711]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1182
			} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1182
				_go_fuzz_dep_.CoverTab[526712]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1182
			}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1182
			_go_fuzz_dep_.CoverTab[786592] = 1
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1182
			_go_fuzz_dep_.CoverTab[2926]++
													s += ", " + a.GoString()
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1183
			// _ = "end of CoverTab[2926]"
		}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1184
		if _go_fuzz_dep_.CoverTab[786592] == 0 {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1184
			_go_fuzz_dep_.CoverTab[526713]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1184
		} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1184
			_go_fuzz_dep_.CoverTab[526714]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1184
		}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1184
		// _ = "end of CoverTab[2925]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1185
		_go_fuzz_dep_.CoverTab[526314]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1185
		_go_fuzz_dep_.CoverTab[2927]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1185
		// _ = "end of CoverTab[2927]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1185
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1185
	// _ = "end of CoverTab[2914]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1185
	_go_fuzz_dep_.CoverTab[2915]++
											return s + "}}"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1186
	// _ = "end of CoverTab[2915]"
}

// A Builder allows incrementally packing a DNS message.
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1189
//
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1189
// Example usage:
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1189
//
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1189
//	buf := make([]byte, 2, 514)
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1189
//	b := NewBuilder(buf, Header{...})
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1189
//	b.EnableCompression()
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1189
//	// Optionally start a section and add things to that section.
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1189
//	// Repeat adding sections as necessary.
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1189
//	buf, err := b.Finish()
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1189
//	// If err is nil, buf[2:] will contain the built bytes.
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1200
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
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1219
//
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1219
// Note: Most users will want to immediately enable compression with the
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1219
// EnableCompression method. See that method's comment for why you may or may
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1219
// not want to enable compression.
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1219
//
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1219
// The DNS message is appended to the provided initial buffer buf (which may be
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1219
// nil) as it is built. The final message is returned by the (*Builder).Finish
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1219
// method, which includes buf[:len(buf)] and may return the same underlying
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1219
// array if there was sufficient capacity in the slice.
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1229
func NewBuilder(buf []byte, h Header) Builder {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1229
	_go_fuzz_dep_.CoverTab[2928]++
											if buf == nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1230
		_go_fuzz_dep_.CoverTab[526315]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1230
		_go_fuzz_dep_.CoverTab[2930]++
												buf = make([]byte, 0, packStartingCap)
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1231
		// _ = "end of CoverTab[2930]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1232
		_go_fuzz_dep_.CoverTab[526316]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1232
		_go_fuzz_dep_.CoverTab[2931]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1232
		// _ = "end of CoverTab[2931]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1232
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1232
	// _ = "end of CoverTab[2928]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1232
	_go_fuzz_dep_.CoverTab[2929]++
											b := Builder{msg: buf, start: len(buf)}
											b.header.id, b.header.bits = h.pack()
											var hb [headerLen]byte
											b.msg = append(b.msg, hb[:]...)
											b.section = sectionHeader
											return b
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1238
	// _ = "end of CoverTab[2929]"
}

// EnableCompression enables compression in the Builder.
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1241
//
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1241
// Leaving compression disabled avoids compression related allocations, but can
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1241
// result in larger message sizes. Be careful with this mode as it can cause
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1241
// messages to exceed the UDP size limit.
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1241
//
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1241
// According to RFC 1035, section 4.1.4, the use of compression is optional, but
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1241
// all implementations must accept both compressed and uncompressed DNS
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1241
// messages.
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1241
//
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1241
// Compression should be enabled before any sections are added for best results.
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1252
func (b *Builder) EnableCompression() {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1252
	_go_fuzz_dep_.CoverTab[2932]++
											b.compression = map[string]int{}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1253
	// _ = "end of CoverTab[2932]"
}

func (b *Builder) startCheck(s section) error {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1256
	_go_fuzz_dep_.CoverTab[2933]++
											if b.section <= sectionNotStarted {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1257
		_go_fuzz_dep_.CoverTab[526317]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1257
		_go_fuzz_dep_.CoverTab[2936]++
												return ErrNotStarted
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1258
		// _ = "end of CoverTab[2936]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1259
		_go_fuzz_dep_.CoverTab[526318]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1259
		_go_fuzz_dep_.CoverTab[2937]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1259
		// _ = "end of CoverTab[2937]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1259
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1259
	// _ = "end of CoverTab[2933]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1259
	_go_fuzz_dep_.CoverTab[2934]++
											if b.section > s {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1260
		_go_fuzz_dep_.CoverTab[526319]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1260
		_go_fuzz_dep_.CoverTab[2938]++
												return ErrSectionDone
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1261
		// _ = "end of CoverTab[2938]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1262
		_go_fuzz_dep_.CoverTab[526320]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1262
		_go_fuzz_dep_.CoverTab[2939]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1262
		// _ = "end of CoverTab[2939]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1262
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1262
	// _ = "end of CoverTab[2934]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1262
	_go_fuzz_dep_.CoverTab[2935]++
											return nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1263
	// _ = "end of CoverTab[2935]"
}

// StartQuestions prepares the builder for packing Questions.
func (b *Builder) StartQuestions() error {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1267
	_go_fuzz_dep_.CoverTab[2940]++
											if err := b.startCheck(sectionQuestions); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1268
		_go_fuzz_dep_.CoverTab[526321]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1268
		_go_fuzz_dep_.CoverTab[2942]++
												return err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1269
		// _ = "end of CoverTab[2942]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1270
		_go_fuzz_dep_.CoverTab[526322]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1270
		_go_fuzz_dep_.CoverTab[2943]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1270
		// _ = "end of CoverTab[2943]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1270
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1270
	// _ = "end of CoverTab[2940]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1270
	_go_fuzz_dep_.CoverTab[2941]++
											b.section = sectionQuestions
											return nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1272
	// _ = "end of CoverTab[2941]"
}

// StartAnswers prepares the builder for packing Answers.
func (b *Builder) StartAnswers() error {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1276
	_go_fuzz_dep_.CoverTab[2944]++
											if err := b.startCheck(sectionAnswers); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1277
		_go_fuzz_dep_.CoverTab[526323]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1277
		_go_fuzz_dep_.CoverTab[2946]++
												return err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1278
		// _ = "end of CoverTab[2946]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1279
		_go_fuzz_dep_.CoverTab[526324]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1279
		_go_fuzz_dep_.CoverTab[2947]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1279
		// _ = "end of CoverTab[2947]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1279
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1279
	// _ = "end of CoverTab[2944]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1279
	_go_fuzz_dep_.CoverTab[2945]++
											b.section = sectionAnswers
											return nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1281
	// _ = "end of CoverTab[2945]"
}

// StartAuthorities prepares the builder for packing Authorities.
func (b *Builder) StartAuthorities() error {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1285
	_go_fuzz_dep_.CoverTab[2948]++
											if err := b.startCheck(sectionAuthorities); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1286
		_go_fuzz_dep_.CoverTab[526325]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1286
		_go_fuzz_dep_.CoverTab[2950]++
												return err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1287
		// _ = "end of CoverTab[2950]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1288
		_go_fuzz_dep_.CoverTab[526326]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1288
		_go_fuzz_dep_.CoverTab[2951]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1288
		// _ = "end of CoverTab[2951]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1288
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1288
	// _ = "end of CoverTab[2948]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1288
	_go_fuzz_dep_.CoverTab[2949]++
											b.section = sectionAuthorities
											return nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1290
	// _ = "end of CoverTab[2949]"
}

// StartAdditionals prepares the builder for packing Additionals.
func (b *Builder) StartAdditionals() error {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1294
	_go_fuzz_dep_.CoverTab[2952]++
											if err := b.startCheck(sectionAdditionals); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1295
		_go_fuzz_dep_.CoverTab[526327]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1295
		_go_fuzz_dep_.CoverTab[2954]++
												return err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1296
		// _ = "end of CoverTab[2954]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1297
		_go_fuzz_dep_.CoverTab[526328]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1297
		_go_fuzz_dep_.CoverTab[2955]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1297
		// _ = "end of CoverTab[2955]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1297
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1297
	// _ = "end of CoverTab[2952]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1297
	_go_fuzz_dep_.CoverTab[2953]++
											b.section = sectionAdditionals
											return nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1299
	// _ = "end of CoverTab[2953]"
}

func (b *Builder) incrementSectionCount() error {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1302
	_go_fuzz_dep_.CoverTab[2956]++
											var count *uint16
											var err error
											switch b.section {
	case sectionQuestions:
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1306
		_go_fuzz_dep_.CoverTab[526329]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1306
		_go_fuzz_dep_.CoverTab[2959]++
												count = &b.header.questions
												err = errTooManyQuestions
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1308
		// _ = "end of CoverTab[2959]"
	case sectionAnswers:
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1309
		_go_fuzz_dep_.CoverTab[526330]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1309
		_go_fuzz_dep_.CoverTab[2960]++
												count = &b.header.answers
												err = errTooManyAnswers
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1311
		// _ = "end of CoverTab[2960]"
	case sectionAuthorities:
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1312
		_go_fuzz_dep_.CoverTab[526331]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1312
		_go_fuzz_dep_.CoverTab[2961]++
												count = &b.header.authorities
												err = errTooManyAuthorities
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1314
		// _ = "end of CoverTab[2961]"
	case sectionAdditionals:
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1315
		_go_fuzz_dep_.CoverTab[526332]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1315
		_go_fuzz_dep_.CoverTab[2962]++
												count = &b.header.additionals
												err = errTooManyAdditionals
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1317
		// _ = "end of CoverTab[2962]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1317
	default:
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1317
		_go_fuzz_dep_.CoverTab[526333]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1317
		_go_fuzz_dep_.CoverTab[2963]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1317
		// _ = "end of CoverTab[2963]"
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1318
	// _ = "end of CoverTab[2956]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1318
	_go_fuzz_dep_.CoverTab[2957]++
											if *count == ^uint16(0) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1319
		_go_fuzz_dep_.CoverTab[526334]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1319
		_go_fuzz_dep_.CoverTab[2964]++
												return err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1320
		// _ = "end of CoverTab[2964]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1321
		_go_fuzz_dep_.CoverTab[526335]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1321
		_go_fuzz_dep_.CoverTab[2965]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1321
		// _ = "end of CoverTab[2965]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1321
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1321
	// _ = "end of CoverTab[2957]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1321
	_go_fuzz_dep_.CoverTab[2958]++
											*count++
											return nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1323
	// _ = "end of CoverTab[2958]"
}

// Question adds a single Question.
func (b *Builder) Question(q Question) error {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1327
	_go_fuzz_dep_.CoverTab[2966]++
											if b.section < sectionQuestions {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1328
		_go_fuzz_dep_.CoverTab[526336]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1328
		_go_fuzz_dep_.CoverTab[2971]++
												return ErrNotStarted
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1329
		// _ = "end of CoverTab[2971]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1330
		_go_fuzz_dep_.CoverTab[526337]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1330
		_go_fuzz_dep_.CoverTab[2972]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1330
		// _ = "end of CoverTab[2972]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1330
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1330
	// _ = "end of CoverTab[2966]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1330
	_go_fuzz_dep_.CoverTab[2967]++
											if b.section > sectionQuestions {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1331
		_go_fuzz_dep_.CoverTab[526338]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1331
		_go_fuzz_dep_.CoverTab[2973]++
												return ErrSectionDone
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1332
		// _ = "end of CoverTab[2973]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1333
		_go_fuzz_dep_.CoverTab[526339]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1333
		_go_fuzz_dep_.CoverTab[2974]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1333
		// _ = "end of CoverTab[2974]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1333
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1333
	// _ = "end of CoverTab[2967]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1333
	_go_fuzz_dep_.CoverTab[2968]++
											msg, err := q.pack(b.msg, b.compression, b.start)
											if err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1335
		_go_fuzz_dep_.CoverTab[526340]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1335
		_go_fuzz_dep_.CoverTab[2975]++
												return err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1336
		// _ = "end of CoverTab[2975]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1337
		_go_fuzz_dep_.CoverTab[526341]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1337
		_go_fuzz_dep_.CoverTab[2976]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1337
		// _ = "end of CoverTab[2976]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1337
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1337
	// _ = "end of CoverTab[2968]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1337
	_go_fuzz_dep_.CoverTab[2969]++
											if err := b.incrementSectionCount(); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1338
		_go_fuzz_dep_.CoverTab[526342]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1338
		_go_fuzz_dep_.CoverTab[2977]++
												return err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1339
		// _ = "end of CoverTab[2977]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1340
		_go_fuzz_dep_.CoverTab[526343]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1340
		_go_fuzz_dep_.CoverTab[2978]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1340
		// _ = "end of CoverTab[2978]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1340
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1340
	// _ = "end of CoverTab[2969]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1340
	_go_fuzz_dep_.CoverTab[2970]++
											b.msg = msg
											return nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1342
	// _ = "end of CoverTab[2970]"
}

func (b *Builder) checkResourceSection() error {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1345
	_go_fuzz_dep_.CoverTab[2979]++
											if b.section < sectionAnswers {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1346
		_go_fuzz_dep_.CoverTab[526344]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1346
		_go_fuzz_dep_.CoverTab[2982]++
												return ErrNotStarted
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1347
		// _ = "end of CoverTab[2982]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1348
		_go_fuzz_dep_.CoverTab[526345]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1348
		_go_fuzz_dep_.CoverTab[2983]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1348
		// _ = "end of CoverTab[2983]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1348
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1348
	// _ = "end of CoverTab[2979]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1348
	_go_fuzz_dep_.CoverTab[2980]++
											if b.section > sectionAdditionals {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1349
		_go_fuzz_dep_.CoverTab[526346]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1349
		_go_fuzz_dep_.CoverTab[2984]++
												return ErrSectionDone
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1350
		// _ = "end of CoverTab[2984]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1351
		_go_fuzz_dep_.CoverTab[526347]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1351
		_go_fuzz_dep_.CoverTab[2985]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1351
		// _ = "end of CoverTab[2985]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1351
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1351
	// _ = "end of CoverTab[2980]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1351
	_go_fuzz_dep_.CoverTab[2981]++
											return nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1352
	// _ = "end of CoverTab[2981]"
}

// CNAMEResource adds a single CNAMEResource.
func (b *Builder) CNAMEResource(h ResourceHeader, r CNAMEResource) error {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1356
	_go_fuzz_dep_.CoverTab[2986]++
											if err := b.checkResourceSection(); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1357
		_go_fuzz_dep_.CoverTab[526348]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1357
		_go_fuzz_dep_.CoverTab[2992]++
												return err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1358
		// _ = "end of CoverTab[2992]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1359
		_go_fuzz_dep_.CoverTab[526349]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1359
		_go_fuzz_dep_.CoverTab[2993]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1359
		// _ = "end of CoverTab[2993]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1359
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1359
	// _ = "end of CoverTab[2986]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1359
	_go_fuzz_dep_.CoverTab[2987]++
											h.Type = r.realType()
											msg, lenOff, err := h.pack(b.msg, b.compression, b.start)
											if err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1362
		_go_fuzz_dep_.CoverTab[526350]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1362
		_go_fuzz_dep_.CoverTab[2994]++
												return &nestedError{"ResourceHeader", err}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1363
		// _ = "end of CoverTab[2994]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1364
		_go_fuzz_dep_.CoverTab[526351]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1364
		_go_fuzz_dep_.CoverTab[2995]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1364
		// _ = "end of CoverTab[2995]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1364
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1364
	// _ = "end of CoverTab[2987]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1364
	_go_fuzz_dep_.CoverTab[2988]++
											preLen := len(msg)
											if msg, err = r.pack(msg, b.compression, b.start); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1366
		_go_fuzz_dep_.CoverTab[526352]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1366
		_go_fuzz_dep_.CoverTab[2996]++
												return &nestedError{"CNAMEResource body", err}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1367
		// _ = "end of CoverTab[2996]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1368
		_go_fuzz_dep_.CoverTab[526353]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1368
		_go_fuzz_dep_.CoverTab[2997]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1368
		// _ = "end of CoverTab[2997]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1368
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1368
	// _ = "end of CoverTab[2988]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1368
	_go_fuzz_dep_.CoverTab[2989]++
											if err := h.fixLen(msg, lenOff, preLen); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1369
		_go_fuzz_dep_.CoverTab[526354]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1369
		_go_fuzz_dep_.CoverTab[2998]++
												return err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1370
		// _ = "end of CoverTab[2998]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1371
		_go_fuzz_dep_.CoverTab[526355]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1371
		_go_fuzz_dep_.CoverTab[2999]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1371
		// _ = "end of CoverTab[2999]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1371
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1371
	// _ = "end of CoverTab[2989]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1371
	_go_fuzz_dep_.CoverTab[2990]++
											if err := b.incrementSectionCount(); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1372
		_go_fuzz_dep_.CoverTab[526356]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1372
		_go_fuzz_dep_.CoverTab[3000]++
												return err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1373
		// _ = "end of CoverTab[3000]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1374
		_go_fuzz_dep_.CoverTab[526357]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1374
		_go_fuzz_dep_.CoverTab[3001]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1374
		// _ = "end of CoverTab[3001]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1374
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1374
	// _ = "end of CoverTab[2990]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1374
	_go_fuzz_dep_.CoverTab[2991]++
											b.msg = msg
											return nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1376
	// _ = "end of CoverTab[2991]"
}

// MXResource adds a single MXResource.
func (b *Builder) MXResource(h ResourceHeader, r MXResource) error {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1380
	_go_fuzz_dep_.CoverTab[3002]++
											if err := b.checkResourceSection(); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1381
		_go_fuzz_dep_.CoverTab[526358]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1381
		_go_fuzz_dep_.CoverTab[3008]++
												return err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1382
		// _ = "end of CoverTab[3008]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1383
		_go_fuzz_dep_.CoverTab[526359]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1383
		_go_fuzz_dep_.CoverTab[3009]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1383
		// _ = "end of CoverTab[3009]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1383
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1383
	// _ = "end of CoverTab[3002]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1383
	_go_fuzz_dep_.CoverTab[3003]++
											h.Type = r.realType()
											msg, lenOff, err := h.pack(b.msg, b.compression, b.start)
											if err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1386
		_go_fuzz_dep_.CoverTab[526360]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1386
		_go_fuzz_dep_.CoverTab[3010]++
												return &nestedError{"ResourceHeader", err}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1387
		// _ = "end of CoverTab[3010]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1388
		_go_fuzz_dep_.CoverTab[526361]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1388
		_go_fuzz_dep_.CoverTab[3011]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1388
		// _ = "end of CoverTab[3011]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1388
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1388
	// _ = "end of CoverTab[3003]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1388
	_go_fuzz_dep_.CoverTab[3004]++
											preLen := len(msg)
											if msg, err = r.pack(msg, b.compression, b.start); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1390
		_go_fuzz_dep_.CoverTab[526362]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1390
		_go_fuzz_dep_.CoverTab[3012]++
												return &nestedError{"MXResource body", err}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1391
		// _ = "end of CoverTab[3012]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1392
		_go_fuzz_dep_.CoverTab[526363]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1392
		_go_fuzz_dep_.CoverTab[3013]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1392
		// _ = "end of CoverTab[3013]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1392
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1392
	// _ = "end of CoverTab[3004]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1392
	_go_fuzz_dep_.CoverTab[3005]++
											if err := h.fixLen(msg, lenOff, preLen); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1393
		_go_fuzz_dep_.CoverTab[526364]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1393
		_go_fuzz_dep_.CoverTab[3014]++
												return err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1394
		// _ = "end of CoverTab[3014]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1395
		_go_fuzz_dep_.CoverTab[526365]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1395
		_go_fuzz_dep_.CoverTab[3015]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1395
		// _ = "end of CoverTab[3015]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1395
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1395
	// _ = "end of CoverTab[3005]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1395
	_go_fuzz_dep_.CoverTab[3006]++
											if err := b.incrementSectionCount(); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1396
		_go_fuzz_dep_.CoverTab[526366]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1396
		_go_fuzz_dep_.CoverTab[3016]++
												return err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1397
		// _ = "end of CoverTab[3016]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1398
		_go_fuzz_dep_.CoverTab[526367]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1398
		_go_fuzz_dep_.CoverTab[3017]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1398
		// _ = "end of CoverTab[3017]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1398
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1398
	// _ = "end of CoverTab[3006]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1398
	_go_fuzz_dep_.CoverTab[3007]++
											b.msg = msg
											return nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1400
	// _ = "end of CoverTab[3007]"
}

// NSResource adds a single NSResource.
func (b *Builder) NSResource(h ResourceHeader, r NSResource) error {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1404
	_go_fuzz_dep_.CoverTab[3018]++
											if err := b.checkResourceSection(); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1405
		_go_fuzz_dep_.CoverTab[526368]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1405
		_go_fuzz_dep_.CoverTab[3024]++
												return err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1406
		// _ = "end of CoverTab[3024]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1407
		_go_fuzz_dep_.CoverTab[526369]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1407
		_go_fuzz_dep_.CoverTab[3025]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1407
		// _ = "end of CoverTab[3025]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1407
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1407
	// _ = "end of CoverTab[3018]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1407
	_go_fuzz_dep_.CoverTab[3019]++
											h.Type = r.realType()
											msg, lenOff, err := h.pack(b.msg, b.compression, b.start)
											if err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1410
		_go_fuzz_dep_.CoverTab[526370]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1410
		_go_fuzz_dep_.CoverTab[3026]++
												return &nestedError{"ResourceHeader", err}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1411
		// _ = "end of CoverTab[3026]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1412
		_go_fuzz_dep_.CoverTab[526371]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1412
		_go_fuzz_dep_.CoverTab[3027]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1412
		// _ = "end of CoverTab[3027]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1412
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1412
	// _ = "end of CoverTab[3019]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1412
	_go_fuzz_dep_.CoverTab[3020]++
											preLen := len(msg)
											if msg, err = r.pack(msg, b.compression, b.start); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1414
		_go_fuzz_dep_.CoverTab[526372]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1414
		_go_fuzz_dep_.CoverTab[3028]++
												return &nestedError{"NSResource body", err}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1415
		// _ = "end of CoverTab[3028]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1416
		_go_fuzz_dep_.CoverTab[526373]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1416
		_go_fuzz_dep_.CoverTab[3029]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1416
		// _ = "end of CoverTab[3029]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1416
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1416
	// _ = "end of CoverTab[3020]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1416
	_go_fuzz_dep_.CoverTab[3021]++
											if err := h.fixLen(msg, lenOff, preLen); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1417
		_go_fuzz_dep_.CoverTab[526374]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1417
		_go_fuzz_dep_.CoverTab[3030]++
												return err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1418
		// _ = "end of CoverTab[3030]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1419
		_go_fuzz_dep_.CoverTab[526375]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1419
		_go_fuzz_dep_.CoverTab[3031]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1419
		// _ = "end of CoverTab[3031]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1419
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1419
	// _ = "end of CoverTab[3021]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1419
	_go_fuzz_dep_.CoverTab[3022]++
											if err := b.incrementSectionCount(); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1420
		_go_fuzz_dep_.CoverTab[526376]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1420
		_go_fuzz_dep_.CoverTab[3032]++
												return err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1421
		// _ = "end of CoverTab[3032]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1422
		_go_fuzz_dep_.CoverTab[526377]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1422
		_go_fuzz_dep_.CoverTab[3033]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1422
		// _ = "end of CoverTab[3033]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1422
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1422
	// _ = "end of CoverTab[3022]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1422
	_go_fuzz_dep_.CoverTab[3023]++
											b.msg = msg
											return nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1424
	// _ = "end of CoverTab[3023]"
}

// PTRResource adds a single PTRResource.
func (b *Builder) PTRResource(h ResourceHeader, r PTRResource) error {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1428
	_go_fuzz_dep_.CoverTab[3034]++
											if err := b.checkResourceSection(); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1429
		_go_fuzz_dep_.CoverTab[526378]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1429
		_go_fuzz_dep_.CoverTab[3040]++
												return err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1430
		// _ = "end of CoverTab[3040]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1431
		_go_fuzz_dep_.CoverTab[526379]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1431
		_go_fuzz_dep_.CoverTab[3041]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1431
		// _ = "end of CoverTab[3041]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1431
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1431
	// _ = "end of CoverTab[3034]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1431
	_go_fuzz_dep_.CoverTab[3035]++
											h.Type = r.realType()
											msg, lenOff, err := h.pack(b.msg, b.compression, b.start)
											if err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1434
		_go_fuzz_dep_.CoverTab[526380]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1434
		_go_fuzz_dep_.CoverTab[3042]++
												return &nestedError{"ResourceHeader", err}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1435
		// _ = "end of CoverTab[3042]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1436
		_go_fuzz_dep_.CoverTab[526381]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1436
		_go_fuzz_dep_.CoverTab[3043]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1436
		// _ = "end of CoverTab[3043]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1436
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1436
	// _ = "end of CoverTab[3035]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1436
	_go_fuzz_dep_.CoverTab[3036]++
											preLen := len(msg)
											if msg, err = r.pack(msg, b.compression, b.start); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1438
		_go_fuzz_dep_.CoverTab[526382]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1438
		_go_fuzz_dep_.CoverTab[3044]++
												return &nestedError{"PTRResource body", err}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1439
		// _ = "end of CoverTab[3044]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1440
		_go_fuzz_dep_.CoverTab[526383]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1440
		_go_fuzz_dep_.CoverTab[3045]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1440
		// _ = "end of CoverTab[3045]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1440
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1440
	// _ = "end of CoverTab[3036]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1440
	_go_fuzz_dep_.CoverTab[3037]++
											if err := h.fixLen(msg, lenOff, preLen); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1441
		_go_fuzz_dep_.CoverTab[526384]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1441
		_go_fuzz_dep_.CoverTab[3046]++
												return err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1442
		// _ = "end of CoverTab[3046]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1443
		_go_fuzz_dep_.CoverTab[526385]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1443
		_go_fuzz_dep_.CoverTab[3047]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1443
		// _ = "end of CoverTab[3047]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1443
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1443
	// _ = "end of CoverTab[3037]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1443
	_go_fuzz_dep_.CoverTab[3038]++
											if err := b.incrementSectionCount(); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1444
		_go_fuzz_dep_.CoverTab[526386]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1444
		_go_fuzz_dep_.CoverTab[3048]++
												return err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1445
		// _ = "end of CoverTab[3048]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1446
		_go_fuzz_dep_.CoverTab[526387]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1446
		_go_fuzz_dep_.CoverTab[3049]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1446
		// _ = "end of CoverTab[3049]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1446
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1446
	// _ = "end of CoverTab[3038]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1446
	_go_fuzz_dep_.CoverTab[3039]++
											b.msg = msg
											return nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1448
	// _ = "end of CoverTab[3039]"
}

// SOAResource adds a single SOAResource.
func (b *Builder) SOAResource(h ResourceHeader, r SOAResource) error {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1452
	_go_fuzz_dep_.CoverTab[3050]++
											if err := b.checkResourceSection(); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1453
		_go_fuzz_dep_.CoverTab[526388]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1453
		_go_fuzz_dep_.CoverTab[3056]++
												return err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1454
		// _ = "end of CoverTab[3056]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1455
		_go_fuzz_dep_.CoverTab[526389]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1455
		_go_fuzz_dep_.CoverTab[3057]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1455
		// _ = "end of CoverTab[3057]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1455
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1455
	// _ = "end of CoverTab[3050]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1455
	_go_fuzz_dep_.CoverTab[3051]++
											h.Type = r.realType()
											msg, lenOff, err := h.pack(b.msg, b.compression, b.start)
											if err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1458
		_go_fuzz_dep_.CoverTab[526390]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1458
		_go_fuzz_dep_.CoverTab[3058]++
												return &nestedError{"ResourceHeader", err}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1459
		// _ = "end of CoverTab[3058]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1460
		_go_fuzz_dep_.CoverTab[526391]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1460
		_go_fuzz_dep_.CoverTab[3059]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1460
		// _ = "end of CoverTab[3059]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1460
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1460
	// _ = "end of CoverTab[3051]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1460
	_go_fuzz_dep_.CoverTab[3052]++
											preLen := len(msg)
											if msg, err = r.pack(msg, b.compression, b.start); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1462
		_go_fuzz_dep_.CoverTab[526392]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1462
		_go_fuzz_dep_.CoverTab[3060]++
												return &nestedError{"SOAResource body", err}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1463
		// _ = "end of CoverTab[3060]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1464
		_go_fuzz_dep_.CoverTab[526393]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1464
		_go_fuzz_dep_.CoverTab[3061]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1464
		// _ = "end of CoverTab[3061]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1464
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1464
	// _ = "end of CoverTab[3052]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1464
	_go_fuzz_dep_.CoverTab[3053]++
											if err := h.fixLen(msg, lenOff, preLen); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1465
		_go_fuzz_dep_.CoverTab[526394]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1465
		_go_fuzz_dep_.CoverTab[3062]++
												return err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1466
		// _ = "end of CoverTab[3062]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1467
		_go_fuzz_dep_.CoverTab[526395]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1467
		_go_fuzz_dep_.CoverTab[3063]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1467
		// _ = "end of CoverTab[3063]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1467
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1467
	// _ = "end of CoverTab[3053]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1467
	_go_fuzz_dep_.CoverTab[3054]++
											if err := b.incrementSectionCount(); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1468
		_go_fuzz_dep_.CoverTab[526396]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1468
		_go_fuzz_dep_.CoverTab[3064]++
												return err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1469
		// _ = "end of CoverTab[3064]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1470
		_go_fuzz_dep_.CoverTab[526397]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1470
		_go_fuzz_dep_.CoverTab[3065]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1470
		// _ = "end of CoverTab[3065]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1470
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1470
	// _ = "end of CoverTab[3054]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1470
	_go_fuzz_dep_.CoverTab[3055]++
											b.msg = msg
											return nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1472
	// _ = "end of CoverTab[3055]"
}

// TXTResource adds a single TXTResource.
func (b *Builder) TXTResource(h ResourceHeader, r TXTResource) error {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1476
	_go_fuzz_dep_.CoverTab[3066]++
											if err := b.checkResourceSection(); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1477
		_go_fuzz_dep_.CoverTab[526398]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1477
		_go_fuzz_dep_.CoverTab[3072]++
												return err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1478
		// _ = "end of CoverTab[3072]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1479
		_go_fuzz_dep_.CoverTab[526399]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1479
		_go_fuzz_dep_.CoverTab[3073]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1479
		// _ = "end of CoverTab[3073]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1479
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1479
	// _ = "end of CoverTab[3066]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1479
	_go_fuzz_dep_.CoverTab[3067]++
											h.Type = r.realType()
											msg, lenOff, err := h.pack(b.msg, b.compression, b.start)
											if err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1482
		_go_fuzz_dep_.CoverTab[526400]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1482
		_go_fuzz_dep_.CoverTab[3074]++
												return &nestedError{"ResourceHeader", err}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1483
		// _ = "end of CoverTab[3074]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1484
		_go_fuzz_dep_.CoverTab[526401]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1484
		_go_fuzz_dep_.CoverTab[3075]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1484
		// _ = "end of CoverTab[3075]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1484
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1484
	// _ = "end of CoverTab[3067]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1484
	_go_fuzz_dep_.CoverTab[3068]++
											preLen := len(msg)
											if msg, err = r.pack(msg, b.compression, b.start); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1486
		_go_fuzz_dep_.CoverTab[526402]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1486
		_go_fuzz_dep_.CoverTab[3076]++
												return &nestedError{"TXTResource body", err}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1487
		// _ = "end of CoverTab[3076]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1488
		_go_fuzz_dep_.CoverTab[526403]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1488
		_go_fuzz_dep_.CoverTab[3077]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1488
		// _ = "end of CoverTab[3077]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1488
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1488
	// _ = "end of CoverTab[3068]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1488
	_go_fuzz_dep_.CoverTab[3069]++
											if err := h.fixLen(msg, lenOff, preLen); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1489
		_go_fuzz_dep_.CoverTab[526404]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1489
		_go_fuzz_dep_.CoverTab[3078]++
												return err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1490
		// _ = "end of CoverTab[3078]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1491
		_go_fuzz_dep_.CoverTab[526405]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1491
		_go_fuzz_dep_.CoverTab[3079]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1491
		// _ = "end of CoverTab[3079]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1491
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1491
	// _ = "end of CoverTab[3069]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1491
	_go_fuzz_dep_.CoverTab[3070]++
											if err := b.incrementSectionCount(); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1492
		_go_fuzz_dep_.CoverTab[526406]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1492
		_go_fuzz_dep_.CoverTab[3080]++
												return err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1493
		// _ = "end of CoverTab[3080]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1494
		_go_fuzz_dep_.CoverTab[526407]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1494
		_go_fuzz_dep_.CoverTab[3081]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1494
		// _ = "end of CoverTab[3081]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1494
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1494
	// _ = "end of CoverTab[3070]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1494
	_go_fuzz_dep_.CoverTab[3071]++
											b.msg = msg
											return nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1496
	// _ = "end of CoverTab[3071]"
}

// SRVResource adds a single SRVResource.
func (b *Builder) SRVResource(h ResourceHeader, r SRVResource) error {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1500
	_go_fuzz_dep_.CoverTab[3082]++
											if err := b.checkResourceSection(); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1501
		_go_fuzz_dep_.CoverTab[526408]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1501
		_go_fuzz_dep_.CoverTab[3088]++
												return err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1502
		// _ = "end of CoverTab[3088]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1503
		_go_fuzz_dep_.CoverTab[526409]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1503
		_go_fuzz_dep_.CoverTab[3089]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1503
		// _ = "end of CoverTab[3089]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1503
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1503
	// _ = "end of CoverTab[3082]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1503
	_go_fuzz_dep_.CoverTab[3083]++
											h.Type = r.realType()
											msg, lenOff, err := h.pack(b.msg, b.compression, b.start)
											if err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1506
		_go_fuzz_dep_.CoverTab[526410]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1506
		_go_fuzz_dep_.CoverTab[3090]++
												return &nestedError{"ResourceHeader", err}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1507
		// _ = "end of CoverTab[3090]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1508
		_go_fuzz_dep_.CoverTab[526411]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1508
		_go_fuzz_dep_.CoverTab[3091]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1508
		// _ = "end of CoverTab[3091]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1508
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1508
	// _ = "end of CoverTab[3083]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1508
	_go_fuzz_dep_.CoverTab[3084]++
											preLen := len(msg)
											if msg, err = r.pack(msg, b.compression, b.start); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1510
		_go_fuzz_dep_.CoverTab[526412]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1510
		_go_fuzz_dep_.CoverTab[3092]++
												return &nestedError{"SRVResource body", err}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1511
		// _ = "end of CoverTab[3092]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1512
		_go_fuzz_dep_.CoverTab[526413]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1512
		_go_fuzz_dep_.CoverTab[3093]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1512
		// _ = "end of CoverTab[3093]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1512
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1512
	// _ = "end of CoverTab[3084]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1512
	_go_fuzz_dep_.CoverTab[3085]++
											if err := h.fixLen(msg, lenOff, preLen); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1513
		_go_fuzz_dep_.CoverTab[526414]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1513
		_go_fuzz_dep_.CoverTab[3094]++
												return err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1514
		// _ = "end of CoverTab[3094]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1515
		_go_fuzz_dep_.CoverTab[526415]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1515
		_go_fuzz_dep_.CoverTab[3095]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1515
		// _ = "end of CoverTab[3095]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1515
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1515
	// _ = "end of CoverTab[3085]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1515
	_go_fuzz_dep_.CoverTab[3086]++
											if err := b.incrementSectionCount(); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1516
		_go_fuzz_dep_.CoverTab[526416]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1516
		_go_fuzz_dep_.CoverTab[3096]++
												return err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1517
		// _ = "end of CoverTab[3096]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1518
		_go_fuzz_dep_.CoverTab[526417]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1518
		_go_fuzz_dep_.CoverTab[3097]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1518
		// _ = "end of CoverTab[3097]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1518
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1518
	// _ = "end of CoverTab[3086]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1518
	_go_fuzz_dep_.CoverTab[3087]++
											b.msg = msg
											return nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1520
	// _ = "end of CoverTab[3087]"
}

// AResource adds a single AResource.
func (b *Builder) AResource(h ResourceHeader, r AResource) error {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1524
	_go_fuzz_dep_.CoverTab[3098]++
											if err := b.checkResourceSection(); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1525
		_go_fuzz_dep_.CoverTab[526418]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1525
		_go_fuzz_dep_.CoverTab[3104]++
												return err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1526
		// _ = "end of CoverTab[3104]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1527
		_go_fuzz_dep_.CoverTab[526419]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1527
		_go_fuzz_dep_.CoverTab[3105]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1527
		// _ = "end of CoverTab[3105]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1527
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1527
	// _ = "end of CoverTab[3098]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1527
	_go_fuzz_dep_.CoverTab[3099]++
											h.Type = r.realType()
											msg, lenOff, err := h.pack(b.msg, b.compression, b.start)
											if err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1530
		_go_fuzz_dep_.CoverTab[526420]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1530
		_go_fuzz_dep_.CoverTab[3106]++
												return &nestedError{"ResourceHeader", err}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1531
		// _ = "end of CoverTab[3106]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1532
		_go_fuzz_dep_.CoverTab[526421]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1532
		_go_fuzz_dep_.CoverTab[3107]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1532
		// _ = "end of CoverTab[3107]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1532
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1532
	// _ = "end of CoverTab[3099]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1532
	_go_fuzz_dep_.CoverTab[3100]++
											preLen := len(msg)
											if msg, err = r.pack(msg, b.compression, b.start); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1534
		_go_fuzz_dep_.CoverTab[526422]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1534
		_go_fuzz_dep_.CoverTab[3108]++
												return &nestedError{"AResource body", err}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1535
		// _ = "end of CoverTab[3108]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1536
		_go_fuzz_dep_.CoverTab[526423]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1536
		_go_fuzz_dep_.CoverTab[3109]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1536
		// _ = "end of CoverTab[3109]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1536
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1536
	// _ = "end of CoverTab[3100]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1536
	_go_fuzz_dep_.CoverTab[3101]++
											if err := h.fixLen(msg, lenOff, preLen); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1537
		_go_fuzz_dep_.CoverTab[526424]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1537
		_go_fuzz_dep_.CoverTab[3110]++
												return err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1538
		// _ = "end of CoverTab[3110]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1539
		_go_fuzz_dep_.CoverTab[526425]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1539
		_go_fuzz_dep_.CoverTab[3111]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1539
		// _ = "end of CoverTab[3111]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1539
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1539
	// _ = "end of CoverTab[3101]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1539
	_go_fuzz_dep_.CoverTab[3102]++
											if err := b.incrementSectionCount(); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1540
		_go_fuzz_dep_.CoverTab[526426]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1540
		_go_fuzz_dep_.CoverTab[3112]++
												return err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1541
		// _ = "end of CoverTab[3112]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1542
		_go_fuzz_dep_.CoverTab[526427]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1542
		_go_fuzz_dep_.CoverTab[3113]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1542
		// _ = "end of CoverTab[3113]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1542
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1542
	// _ = "end of CoverTab[3102]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1542
	_go_fuzz_dep_.CoverTab[3103]++
											b.msg = msg
											return nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1544
	// _ = "end of CoverTab[3103]"
}

// AAAAResource adds a single AAAAResource.
func (b *Builder) AAAAResource(h ResourceHeader, r AAAAResource) error {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1548
	_go_fuzz_dep_.CoverTab[3114]++
											if err := b.checkResourceSection(); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1549
		_go_fuzz_dep_.CoverTab[526428]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1549
		_go_fuzz_dep_.CoverTab[3120]++
												return err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1550
		// _ = "end of CoverTab[3120]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1551
		_go_fuzz_dep_.CoverTab[526429]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1551
		_go_fuzz_dep_.CoverTab[3121]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1551
		// _ = "end of CoverTab[3121]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1551
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1551
	// _ = "end of CoverTab[3114]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1551
	_go_fuzz_dep_.CoverTab[3115]++
											h.Type = r.realType()
											msg, lenOff, err := h.pack(b.msg, b.compression, b.start)
											if err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1554
		_go_fuzz_dep_.CoverTab[526430]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1554
		_go_fuzz_dep_.CoverTab[3122]++
												return &nestedError{"ResourceHeader", err}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1555
		// _ = "end of CoverTab[3122]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1556
		_go_fuzz_dep_.CoverTab[526431]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1556
		_go_fuzz_dep_.CoverTab[3123]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1556
		// _ = "end of CoverTab[3123]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1556
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1556
	// _ = "end of CoverTab[3115]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1556
	_go_fuzz_dep_.CoverTab[3116]++
											preLen := len(msg)
											if msg, err = r.pack(msg, b.compression, b.start); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1558
		_go_fuzz_dep_.CoverTab[526432]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1558
		_go_fuzz_dep_.CoverTab[3124]++
												return &nestedError{"AAAAResource body", err}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1559
		// _ = "end of CoverTab[3124]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1560
		_go_fuzz_dep_.CoverTab[526433]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1560
		_go_fuzz_dep_.CoverTab[3125]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1560
		// _ = "end of CoverTab[3125]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1560
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1560
	// _ = "end of CoverTab[3116]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1560
	_go_fuzz_dep_.CoverTab[3117]++
											if err := h.fixLen(msg, lenOff, preLen); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1561
		_go_fuzz_dep_.CoverTab[526434]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1561
		_go_fuzz_dep_.CoverTab[3126]++
												return err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1562
		// _ = "end of CoverTab[3126]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1563
		_go_fuzz_dep_.CoverTab[526435]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1563
		_go_fuzz_dep_.CoverTab[3127]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1563
		// _ = "end of CoverTab[3127]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1563
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1563
	// _ = "end of CoverTab[3117]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1563
	_go_fuzz_dep_.CoverTab[3118]++
											if err := b.incrementSectionCount(); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1564
		_go_fuzz_dep_.CoverTab[526436]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1564
		_go_fuzz_dep_.CoverTab[3128]++
												return err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1565
		// _ = "end of CoverTab[3128]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1566
		_go_fuzz_dep_.CoverTab[526437]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1566
		_go_fuzz_dep_.CoverTab[3129]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1566
		// _ = "end of CoverTab[3129]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1566
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1566
	// _ = "end of CoverTab[3118]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1566
	_go_fuzz_dep_.CoverTab[3119]++
											b.msg = msg
											return nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1568
	// _ = "end of CoverTab[3119]"
}

// OPTResource adds a single OPTResource.
func (b *Builder) OPTResource(h ResourceHeader, r OPTResource) error {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1572
	_go_fuzz_dep_.CoverTab[3130]++
											if err := b.checkResourceSection(); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1573
		_go_fuzz_dep_.CoverTab[526438]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1573
		_go_fuzz_dep_.CoverTab[3136]++
												return err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1574
		// _ = "end of CoverTab[3136]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1575
		_go_fuzz_dep_.CoverTab[526439]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1575
		_go_fuzz_dep_.CoverTab[3137]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1575
		// _ = "end of CoverTab[3137]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1575
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1575
	// _ = "end of CoverTab[3130]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1575
	_go_fuzz_dep_.CoverTab[3131]++
											h.Type = r.realType()
											msg, lenOff, err := h.pack(b.msg, b.compression, b.start)
											if err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1578
		_go_fuzz_dep_.CoverTab[526440]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1578
		_go_fuzz_dep_.CoverTab[3138]++
												return &nestedError{"ResourceHeader", err}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1579
		// _ = "end of CoverTab[3138]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1580
		_go_fuzz_dep_.CoverTab[526441]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1580
		_go_fuzz_dep_.CoverTab[3139]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1580
		// _ = "end of CoverTab[3139]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1580
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1580
	// _ = "end of CoverTab[3131]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1580
	_go_fuzz_dep_.CoverTab[3132]++
											preLen := len(msg)
											if msg, err = r.pack(msg, b.compression, b.start); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1582
		_go_fuzz_dep_.CoverTab[526442]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1582
		_go_fuzz_dep_.CoverTab[3140]++
												return &nestedError{"OPTResource body", err}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1583
		// _ = "end of CoverTab[3140]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1584
		_go_fuzz_dep_.CoverTab[526443]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1584
		_go_fuzz_dep_.CoverTab[3141]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1584
		// _ = "end of CoverTab[3141]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1584
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1584
	// _ = "end of CoverTab[3132]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1584
	_go_fuzz_dep_.CoverTab[3133]++
											if err := h.fixLen(msg, lenOff, preLen); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1585
		_go_fuzz_dep_.CoverTab[526444]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1585
		_go_fuzz_dep_.CoverTab[3142]++
												return err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1586
		// _ = "end of CoverTab[3142]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1587
		_go_fuzz_dep_.CoverTab[526445]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1587
		_go_fuzz_dep_.CoverTab[3143]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1587
		// _ = "end of CoverTab[3143]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1587
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1587
	// _ = "end of CoverTab[3133]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1587
	_go_fuzz_dep_.CoverTab[3134]++
											if err := b.incrementSectionCount(); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1588
		_go_fuzz_dep_.CoverTab[526446]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1588
		_go_fuzz_dep_.CoverTab[3144]++
												return err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1589
		// _ = "end of CoverTab[3144]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1590
		_go_fuzz_dep_.CoverTab[526447]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1590
		_go_fuzz_dep_.CoverTab[3145]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1590
		// _ = "end of CoverTab[3145]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1590
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1590
	// _ = "end of CoverTab[3134]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1590
	_go_fuzz_dep_.CoverTab[3135]++
											b.msg = msg
											return nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1592
	// _ = "end of CoverTab[3135]"
}

// UnknownResource adds a single UnknownResource.
func (b *Builder) UnknownResource(h ResourceHeader, r UnknownResource) error {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1596
	_go_fuzz_dep_.CoverTab[3146]++
											if err := b.checkResourceSection(); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1597
		_go_fuzz_dep_.CoverTab[526448]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1597
		_go_fuzz_dep_.CoverTab[3152]++
												return err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1598
		// _ = "end of CoverTab[3152]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1599
		_go_fuzz_dep_.CoverTab[526449]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1599
		_go_fuzz_dep_.CoverTab[3153]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1599
		// _ = "end of CoverTab[3153]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1599
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1599
	// _ = "end of CoverTab[3146]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1599
	_go_fuzz_dep_.CoverTab[3147]++
											h.Type = r.realType()
											msg, lenOff, err := h.pack(b.msg, b.compression, b.start)
											if err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1602
		_go_fuzz_dep_.CoverTab[526450]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1602
		_go_fuzz_dep_.CoverTab[3154]++
												return &nestedError{"ResourceHeader", err}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1603
		// _ = "end of CoverTab[3154]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1604
		_go_fuzz_dep_.CoverTab[526451]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1604
		_go_fuzz_dep_.CoverTab[3155]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1604
		// _ = "end of CoverTab[3155]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1604
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1604
	// _ = "end of CoverTab[3147]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1604
	_go_fuzz_dep_.CoverTab[3148]++
											preLen := len(msg)
											if msg, err = r.pack(msg, b.compression, b.start); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1606
		_go_fuzz_dep_.CoverTab[526452]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1606
		_go_fuzz_dep_.CoverTab[3156]++
												return &nestedError{"UnknownResource body", err}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1607
		// _ = "end of CoverTab[3156]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1608
		_go_fuzz_dep_.CoverTab[526453]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1608
		_go_fuzz_dep_.CoverTab[3157]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1608
		// _ = "end of CoverTab[3157]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1608
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1608
	// _ = "end of CoverTab[3148]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1608
	_go_fuzz_dep_.CoverTab[3149]++
											if err := h.fixLen(msg, lenOff, preLen); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1609
		_go_fuzz_dep_.CoverTab[526454]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1609
		_go_fuzz_dep_.CoverTab[3158]++
												return err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1610
		// _ = "end of CoverTab[3158]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1611
		_go_fuzz_dep_.CoverTab[526455]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1611
		_go_fuzz_dep_.CoverTab[3159]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1611
		// _ = "end of CoverTab[3159]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1611
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1611
	// _ = "end of CoverTab[3149]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1611
	_go_fuzz_dep_.CoverTab[3150]++
											if err := b.incrementSectionCount(); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1612
		_go_fuzz_dep_.CoverTab[526456]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1612
		_go_fuzz_dep_.CoverTab[3160]++
												return err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1613
		// _ = "end of CoverTab[3160]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1614
		_go_fuzz_dep_.CoverTab[526457]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1614
		_go_fuzz_dep_.CoverTab[3161]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1614
		// _ = "end of CoverTab[3161]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1614
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1614
	// _ = "end of CoverTab[3150]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1614
	_go_fuzz_dep_.CoverTab[3151]++
											b.msg = msg
											return nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1616
	// _ = "end of CoverTab[3151]"
}

// Finish ends message building and generates a binary message.
func (b *Builder) Finish() ([]byte, error) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1620
	_go_fuzz_dep_.CoverTab[3162]++
											if b.section < sectionHeader {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1621
		_go_fuzz_dep_.CoverTab[526458]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1621
		_go_fuzz_dep_.CoverTab[3164]++
												return nil, ErrNotStarted
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1622
		// _ = "end of CoverTab[3164]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1623
		_go_fuzz_dep_.CoverTab[526459]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1623
		_go_fuzz_dep_.CoverTab[3165]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1623
		// _ = "end of CoverTab[3165]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1623
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1623
	// _ = "end of CoverTab[3162]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1623
	_go_fuzz_dep_.CoverTab[3163]++
											b.section = sectionDone

											b.header.pack(b.msg[b.start:b.start])
											return b.msg, nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1627
	// _ = "end of CoverTab[3163]"
}

// A ResourceHeader is the header of a DNS resource record. There are
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1630
// many types of DNS resource records, but they all share the same header.
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1632
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
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1657
	_go_fuzz_dep_.CoverTab[3166]++
											return "dnsmessage.ResourceHeader{" +
		"Name: " + h.Name.GoString() + ", " +
		"Type: " + h.Type.GoString() + ", " +
		"Class: " + h.Class.GoString() + ", " +
		"TTL: " + printUint32(h.TTL) + ", " +
		"Length: " + printUint16(h.Length) + "}"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1663
	// _ = "end of CoverTab[3166]"
}

// pack appends the wire format of the ResourceHeader to oldMsg.
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1666
//
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1666
// lenOff is the offset in msg where the Length field was packed.
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1669
func (h *ResourceHeader) pack(oldMsg []byte, compression map[string]int, compressionOff int) (msg []byte, lenOff int, err error) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1669
	_go_fuzz_dep_.CoverTab[3167]++
											msg = oldMsg
											if msg, err = h.Name.pack(msg, compression, compressionOff); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1671
		_go_fuzz_dep_.CoverTab[526460]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1671
		_go_fuzz_dep_.CoverTab[3169]++
												return oldMsg, 0, &nestedError{"Name", err}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1672
		// _ = "end of CoverTab[3169]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1673
		_go_fuzz_dep_.CoverTab[526461]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1673
		_go_fuzz_dep_.CoverTab[3170]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1673
		// _ = "end of CoverTab[3170]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1673
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1673
	// _ = "end of CoverTab[3167]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1673
	_go_fuzz_dep_.CoverTab[3168]++
											msg = packType(msg, h.Type)
											msg = packClass(msg, h.Class)
											msg = packUint32(msg, h.TTL)
											lenOff = len(msg)
											msg = packUint16(msg, h.Length)
											return msg, lenOff, nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1679
	// _ = "end of CoverTab[3168]"
}

func (h *ResourceHeader) unpack(msg []byte, off int) (int, error) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1682
	_go_fuzz_dep_.CoverTab[3171]++
											newOff := off
											var err error
											if newOff, err = h.Name.unpack(msg, newOff); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1685
		_go_fuzz_dep_.CoverTab[526462]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1685
		_go_fuzz_dep_.CoverTab[3177]++
												return off, &nestedError{"Name", err}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1686
		// _ = "end of CoverTab[3177]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1687
		_go_fuzz_dep_.CoverTab[526463]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1687
		_go_fuzz_dep_.CoverTab[3178]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1687
		// _ = "end of CoverTab[3178]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1687
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1687
	// _ = "end of CoverTab[3171]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1687
	_go_fuzz_dep_.CoverTab[3172]++
											if h.Type, newOff, err = unpackType(msg, newOff); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1688
		_go_fuzz_dep_.CoverTab[526464]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1688
		_go_fuzz_dep_.CoverTab[3179]++
												return off, &nestedError{"Type", err}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1689
		// _ = "end of CoverTab[3179]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1690
		_go_fuzz_dep_.CoverTab[526465]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1690
		_go_fuzz_dep_.CoverTab[3180]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1690
		// _ = "end of CoverTab[3180]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1690
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1690
	// _ = "end of CoverTab[3172]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1690
	_go_fuzz_dep_.CoverTab[3173]++
											if h.Class, newOff, err = unpackClass(msg, newOff); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1691
		_go_fuzz_dep_.CoverTab[526466]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1691
		_go_fuzz_dep_.CoverTab[3181]++
												return off, &nestedError{"Class", err}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1692
		// _ = "end of CoverTab[3181]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1693
		_go_fuzz_dep_.CoverTab[526467]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1693
		_go_fuzz_dep_.CoverTab[3182]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1693
		// _ = "end of CoverTab[3182]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1693
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1693
	// _ = "end of CoverTab[3173]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1693
	_go_fuzz_dep_.CoverTab[3174]++
											if h.TTL, newOff, err = unpackUint32(msg, newOff); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1694
		_go_fuzz_dep_.CoverTab[526468]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1694
		_go_fuzz_dep_.CoverTab[3183]++
												return off, &nestedError{"TTL", err}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1695
		// _ = "end of CoverTab[3183]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1696
		_go_fuzz_dep_.CoverTab[526469]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1696
		_go_fuzz_dep_.CoverTab[3184]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1696
		// _ = "end of CoverTab[3184]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1696
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1696
	// _ = "end of CoverTab[3174]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1696
	_go_fuzz_dep_.CoverTab[3175]++
											if h.Length, newOff, err = unpackUint16(msg, newOff); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1697
		_go_fuzz_dep_.CoverTab[526470]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1697
		_go_fuzz_dep_.CoverTab[3185]++
												return off, &nestedError{"Length", err}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1698
		// _ = "end of CoverTab[3185]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1699
		_go_fuzz_dep_.CoverTab[526471]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1699
		_go_fuzz_dep_.CoverTab[3186]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1699
		// _ = "end of CoverTab[3186]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1699
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1699
	// _ = "end of CoverTab[3175]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1699
	_go_fuzz_dep_.CoverTab[3176]++
											return newOff, nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1700
	// _ = "end of CoverTab[3176]"
}

// fixLen updates a packed ResourceHeader to include the length of the
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1703
// ResourceBody.
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1703
//
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1703
// lenOff is the offset of the ResourceHeader.Length field in msg.
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1703
//
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1703
// preLen is the length that msg was before the ResourceBody was packed.
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1709
func (h *ResourceHeader) fixLen(msg []byte, lenOff int, preLen int) error {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1709
	_go_fuzz_dep_.CoverTab[3187]++
											conLen := len(msg) - preLen
											if conLen > int(^uint16(0)) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1711
		_go_fuzz_dep_.CoverTab[526472]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1711
		_go_fuzz_dep_.CoverTab[3189]++
												return errResTooLong
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1712
		// _ = "end of CoverTab[3189]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1713
		_go_fuzz_dep_.CoverTab[526473]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1713
		_go_fuzz_dep_.CoverTab[3190]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1713
		// _ = "end of CoverTab[3190]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1713
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1713
	// _ = "end of CoverTab[3187]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1713
	_go_fuzz_dep_.CoverTab[3188]++

//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1716
	packUint16(msg[lenOff:lenOff], uint16(conLen))
											h.Length = uint16(conLen)

											return nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1719
	// _ = "end of CoverTab[3188]"
}

// EDNS(0) wire constants.
const (
	edns0Version	= 0

	edns0DNSSECOK		= 0x00008000
	ednsVersionMask		= 0x00ff0000
	edns0DNSSECOKMask	= 0x00ff8000
)

// SetEDNS0 configures h for EDNS(0).
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1731
//
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1731
// The provided extRCode must be an extended RCode.
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1734
func (h *ResourceHeader) SetEDNS0(udpPayloadLen int, extRCode RCode, dnssecOK bool) error {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1734
	_go_fuzz_dep_.CoverTab[3191]++
											h.Name = Name{Data: [255]byte{'.'}, Length: 1}
											h.Type = TypeOPT
											h.Class = Class(udpPayloadLen)
											h.TTL = uint32(extRCode) >> 4 << 24
											if dnssecOK {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1739
		_go_fuzz_dep_.CoverTab[526474]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1739
		_go_fuzz_dep_.CoverTab[3193]++
												h.TTL |= edns0DNSSECOK
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1740
		// _ = "end of CoverTab[3193]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1741
		_go_fuzz_dep_.CoverTab[526475]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1741
		_go_fuzz_dep_.CoverTab[3194]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1741
		// _ = "end of CoverTab[3194]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1741
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1741
	// _ = "end of CoverTab[3191]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1741
	_go_fuzz_dep_.CoverTab[3192]++
											return nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1742
	// _ = "end of CoverTab[3192]"
}

// DNSSECAllowed reports whether the DNSSEC OK bit is set.
func (h *ResourceHeader) DNSSECAllowed() bool {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1746
	_go_fuzz_dep_.CoverTab[3195]++
											return h.TTL&edns0DNSSECOKMask == edns0DNSSECOK
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1747
	// _ = "end of CoverTab[3195]"
}

// ExtendedRCode returns an extended RCode.
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1750
//
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1750
// The provided rcode must be the RCode in DNS message header.
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1753
func (h *ResourceHeader) ExtendedRCode(rcode RCode) RCode {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1753
	_go_fuzz_dep_.CoverTab[3196]++
											if h.TTL&ednsVersionMask == edns0Version {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1754
		_go_fuzz_dep_.CoverTab[526476]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1754
		_go_fuzz_dep_.CoverTab[3198]++
												return RCode(h.TTL>>24<<4) | rcode
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1755
		// _ = "end of CoverTab[3198]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1756
		_go_fuzz_dep_.CoverTab[526477]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1756
		_go_fuzz_dep_.CoverTab[3199]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1756
		// _ = "end of CoverTab[3199]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1756
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1756
	// _ = "end of CoverTab[3196]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1756
	_go_fuzz_dep_.CoverTab[3197]++
											return rcode
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1757
	// _ = "end of CoverTab[3197]"
}

func skipResource(msg []byte, off int) (int, error) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1760
	_go_fuzz_dep_.CoverTab[3200]++
											newOff, err := skipName(msg, off)
											if err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1762
		_go_fuzz_dep_.CoverTab[526478]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1762
		_go_fuzz_dep_.CoverTab[3207]++
												return off, &nestedError{"Name", err}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1763
		// _ = "end of CoverTab[3207]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1764
		_go_fuzz_dep_.CoverTab[526479]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1764
		_go_fuzz_dep_.CoverTab[3208]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1764
		// _ = "end of CoverTab[3208]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1764
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1764
	// _ = "end of CoverTab[3200]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1764
	_go_fuzz_dep_.CoverTab[3201]++
											if newOff, err = skipType(msg, newOff); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1765
		_go_fuzz_dep_.CoverTab[526480]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1765
		_go_fuzz_dep_.CoverTab[3209]++
												return off, &nestedError{"Type", err}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1766
		// _ = "end of CoverTab[3209]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1767
		_go_fuzz_dep_.CoverTab[526481]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1767
		_go_fuzz_dep_.CoverTab[3210]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1767
		// _ = "end of CoverTab[3210]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1767
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1767
	// _ = "end of CoverTab[3201]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1767
	_go_fuzz_dep_.CoverTab[3202]++
											if newOff, err = skipClass(msg, newOff); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1768
		_go_fuzz_dep_.CoverTab[526482]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1768
		_go_fuzz_dep_.CoverTab[3211]++
												return off, &nestedError{"Class", err}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1769
		// _ = "end of CoverTab[3211]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1770
		_go_fuzz_dep_.CoverTab[526483]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1770
		_go_fuzz_dep_.CoverTab[3212]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1770
		// _ = "end of CoverTab[3212]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1770
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1770
	// _ = "end of CoverTab[3202]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1770
	_go_fuzz_dep_.CoverTab[3203]++
											if newOff, err = skipUint32(msg, newOff); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1771
		_go_fuzz_dep_.CoverTab[526484]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1771
		_go_fuzz_dep_.CoverTab[3213]++
												return off, &nestedError{"TTL", err}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1772
		// _ = "end of CoverTab[3213]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1773
		_go_fuzz_dep_.CoverTab[526485]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1773
		_go_fuzz_dep_.CoverTab[3214]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1773
		// _ = "end of CoverTab[3214]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1773
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1773
	// _ = "end of CoverTab[3203]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1773
	_go_fuzz_dep_.CoverTab[3204]++
											length, newOff, err := unpackUint16(msg, newOff)
											if err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1775
		_go_fuzz_dep_.CoverTab[526486]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1775
		_go_fuzz_dep_.CoverTab[3215]++
												return off, &nestedError{"Length", err}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1776
		// _ = "end of CoverTab[3215]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1777
		_go_fuzz_dep_.CoverTab[526487]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1777
		_go_fuzz_dep_.CoverTab[3216]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1777
		// _ = "end of CoverTab[3216]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1777
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1777
	// _ = "end of CoverTab[3204]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1777
	_go_fuzz_dep_.CoverTab[3205]++
											if newOff += int(length); newOff > len(msg) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1778
		_go_fuzz_dep_.CoverTab[526488]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1778
		_go_fuzz_dep_.CoverTab[3217]++
												return off, errResourceLen
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1779
		// _ = "end of CoverTab[3217]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1780
		_go_fuzz_dep_.CoverTab[526489]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1780
		_go_fuzz_dep_.CoverTab[3218]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1780
		// _ = "end of CoverTab[3218]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1780
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1780
	// _ = "end of CoverTab[3205]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1780
	_go_fuzz_dep_.CoverTab[3206]++
											return newOff, nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1781
	// _ = "end of CoverTab[3206]"
}

// packUint16 appends the wire format of field to msg.
func packUint16(msg []byte, field uint16) []byte {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1785
	_go_fuzz_dep_.CoverTab[3219]++
											return append(msg, byte(field>>8), byte(field))
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1786
	// _ = "end of CoverTab[3219]"
}

func unpackUint16(msg []byte, off int) (uint16, int, error) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1789
	_go_fuzz_dep_.CoverTab[3220]++
											if off+uint16Len > len(msg) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1790
		_go_fuzz_dep_.CoverTab[526490]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1790
		_go_fuzz_dep_.CoverTab[3222]++
												return 0, off, errBaseLen
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1791
		// _ = "end of CoverTab[3222]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1792
		_go_fuzz_dep_.CoverTab[526491]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1792
		_go_fuzz_dep_.CoverTab[3223]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1792
		// _ = "end of CoverTab[3223]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1792
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1792
	// _ = "end of CoverTab[3220]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1792
	_go_fuzz_dep_.CoverTab[3221]++
											return uint16(msg[off])<<8 | uint16(msg[off+1]), off + uint16Len, nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1793
	// _ = "end of CoverTab[3221]"
}

func skipUint16(msg []byte, off int) (int, error) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1796
	_go_fuzz_dep_.CoverTab[3224]++
											if off+uint16Len > len(msg) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1797
		_go_fuzz_dep_.CoverTab[526492]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1797
		_go_fuzz_dep_.CoverTab[3226]++
												return off, errBaseLen
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1798
		// _ = "end of CoverTab[3226]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1799
		_go_fuzz_dep_.CoverTab[526493]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1799
		_go_fuzz_dep_.CoverTab[3227]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1799
		// _ = "end of CoverTab[3227]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1799
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1799
	// _ = "end of CoverTab[3224]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1799
	_go_fuzz_dep_.CoverTab[3225]++
											return off + uint16Len, nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1800
	// _ = "end of CoverTab[3225]"
}

// packType appends the wire format of field to msg.
func packType(msg []byte, field Type) []byte {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1804
	_go_fuzz_dep_.CoverTab[3228]++
											return packUint16(msg, uint16(field))
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1805
	// _ = "end of CoverTab[3228]"
}

func unpackType(msg []byte, off int) (Type, int, error) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1808
	_go_fuzz_dep_.CoverTab[3229]++
											t, o, err := unpackUint16(msg, off)
											return Type(t), o, err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1810
	// _ = "end of CoverTab[3229]"
}

func skipType(msg []byte, off int) (int, error) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1813
	_go_fuzz_dep_.CoverTab[3230]++
											return skipUint16(msg, off)
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1814
	// _ = "end of CoverTab[3230]"
}

// packClass appends the wire format of field to msg.
func packClass(msg []byte, field Class) []byte {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1818
	_go_fuzz_dep_.CoverTab[3231]++
											return packUint16(msg, uint16(field))
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1819
	// _ = "end of CoverTab[3231]"
}

func unpackClass(msg []byte, off int) (Class, int, error) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1822
	_go_fuzz_dep_.CoverTab[3232]++
											c, o, err := unpackUint16(msg, off)
											return Class(c), o, err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1824
	// _ = "end of CoverTab[3232]"
}

func skipClass(msg []byte, off int) (int, error) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1827
	_go_fuzz_dep_.CoverTab[3233]++
											return skipUint16(msg, off)
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1828
	// _ = "end of CoverTab[3233]"
}

// packUint32 appends the wire format of field to msg.
func packUint32(msg []byte, field uint32) []byte {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1832
	_go_fuzz_dep_.CoverTab[3234]++
											return append(
		msg,
		byte(field>>24),
		byte(field>>16),
		byte(field>>8),
		byte(field),
	)
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1839
	// _ = "end of CoverTab[3234]"
}

func unpackUint32(msg []byte, off int) (uint32, int, error) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1842
	_go_fuzz_dep_.CoverTab[3235]++
											if off+uint32Len > len(msg) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1843
		_go_fuzz_dep_.CoverTab[526494]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1843
		_go_fuzz_dep_.CoverTab[3237]++
												return 0, off, errBaseLen
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1844
		// _ = "end of CoverTab[3237]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1845
		_go_fuzz_dep_.CoverTab[526495]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1845
		_go_fuzz_dep_.CoverTab[3238]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1845
		// _ = "end of CoverTab[3238]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1845
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1845
	// _ = "end of CoverTab[3235]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1845
	_go_fuzz_dep_.CoverTab[3236]++
											v := uint32(msg[off])<<24 | uint32(msg[off+1])<<16 | uint32(msg[off+2])<<8 | uint32(msg[off+3])
											return v, off + uint32Len, nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1847
	// _ = "end of CoverTab[3236]"
}

func skipUint32(msg []byte, off int) (int, error) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1850
	_go_fuzz_dep_.CoverTab[3239]++
											if off+uint32Len > len(msg) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1851
		_go_fuzz_dep_.CoverTab[526496]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1851
		_go_fuzz_dep_.CoverTab[3241]++
												return off, errBaseLen
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1852
		// _ = "end of CoverTab[3241]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1853
		_go_fuzz_dep_.CoverTab[526497]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1853
		_go_fuzz_dep_.CoverTab[3242]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1853
		// _ = "end of CoverTab[3242]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1853
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1853
	// _ = "end of CoverTab[3239]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1853
	_go_fuzz_dep_.CoverTab[3240]++
											return off + uint32Len, nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1854
	// _ = "end of CoverTab[3240]"
}

// packText appends the wire format of field to msg.
func packText(msg []byte, field string) ([]byte, error) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1858
	_go_fuzz_dep_.CoverTab[3243]++
											l := len(field)
											if l > 255 {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1860
		_go_fuzz_dep_.CoverTab[526498]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1860
		_go_fuzz_dep_.CoverTab[3245]++
												return nil, errStringTooLong
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1861
		// _ = "end of CoverTab[3245]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1862
		_go_fuzz_dep_.CoverTab[526499]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1862
		_go_fuzz_dep_.CoverTab[3246]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1862
		// _ = "end of CoverTab[3246]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1862
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1862
	// _ = "end of CoverTab[3243]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1862
	_go_fuzz_dep_.CoverTab[3244]++
											msg = append(msg, byte(l))
											msg = append(msg, field...)

											return msg, nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1866
	// _ = "end of CoverTab[3244]"
}

func unpackText(msg []byte, off int) (string, int, error) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1869
	_go_fuzz_dep_.CoverTab[3247]++
											if off >= len(msg) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1870
		_go_fuzz_dep_.CoverTab[526500]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1870
		_go_fuzz_dep_.CoverTab[3250]++
												return "", off, errBaseLen
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1871
		// _ = "end of CoverTab[3250]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1872
		_go_fuzz_dep_.CoverTab[526501]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1872
		_go_fuzz_dep_.CoverTab[3251]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1872
		// _ = "end of CoverTab[3251]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1872
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1872
	// _ = "end of CoverTab[3247]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1872
	_go_fuzz_dep_.CoverTab[3248]++
											beginOff := off + 1
											endOff := beginOff + int(msg[off])
											if endOff > len(msg) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1875
		_go_fuzz_dep_.CoverTab[526502]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1875
		_go_fuzz_dep_.CoverTab[3252]++
												return "", off, errCalcLen
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1876
		// _ = "end of CoverTab[3252]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1877
		_go_fuzz_dep_.CoverTab[526503]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1877
		_go_fuzz_dep_.CoverTab[3253]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1877
		// _ = "end of CoverTab[3253]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1877
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1877
	// _ = "end of CoverTab[3248]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1877
	_go_fuzz_dep_.CoverTab[3249]++
											return string(msg[beginOff:endOff]), endOff, nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1878
	// _ = "end of CoverTab[3249]"
}

// packBytes appends the wire format of field to msg.
func packBytes(msg []byte, field []byte) []byte {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1882
	_go_fuzz_dep_.CoverTab[3254]++
											return append(msg, field...)
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1883
	// _ = "end of CoverTab[3254]"
}

func unpackBytes(msg []byte, off int, field []byte) (int, error) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1886
	_go_fuzz_dep_.CoverTab[3255]++
											newOff := off + len(field)
											if newOff > len(msg) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1888
		_go_fuzz_dep_.CoverTab[526504]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1888
		_go_fuzz_dep_.CoverTab[3257]++
												return off, errBaseLen
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1889
		// _ = "end of CoverTab[3257]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1890
		_go_fuzz_dep_.CoverTab[526505]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1890
		_go_fuzz_dep_.CoverTab[3258]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1890
		// _ = "end of CoverTab[3258]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1890
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1890
	// _ = "end of CoverTab[3255]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1890
	_go_fuzz_dep_.CoverTab[3256]++
											copy(field, msg[off:newOff])
											return newOff, nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1892
	// _ = "end of CoverTab[3256]"
}

const nonEncodedNameMax = 254

// A Name is a non-encoded domain name. It is used instead of strings to avoid
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1897
// allocations.
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1899
type Name struct {
	Data	[255]byte
	Length	uint8
}

// NewName creates a new Name from a string.
func NewName(name string) (Name, error) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1905
	_go_fuzz_dep_.CoverTab[3259]++
											n := Name{Length: uint8(len(name))}
											if len(name) > len(n.Data) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1907
		_go_fuzz_dep_.CoverTab[526506]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1907
		_go_fuzz_dep_.CoverTab[3261]++
												return Name{}, errCalcLen
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1908
		// _ = "end of CoverTab[3261]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1909
		_go_fuzz_dep_.CoverTab[526507]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1909
		_go_fuzz_dep_.CoverTab[3262]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1909
		// _ = "end of CoverTab[3262]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1909
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1909
	// _ = "end of CoverTab[3259]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1909
	_go_fuzz_dep_.CoverTab[3260]++
											copy(n.Data[:], name)
											return n, nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1911
	// _ = "end of CoverTab[3260]"
}

// MustNewName creates a new Name from a string and panics on error.
func MustNewName(name string) Name {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1915
	_go_fuzz_dep_.CoverTab[3263]++
											n, err := NewName(name)
											if err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1917
		_go_fuzz_dep_.CoverTab[526508]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1917
		_go_fuzz_dep_.CoverTab[3265]++
												panic("creating name: " + err.Error())
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1918
		// _ = "end of CoverTab[3265]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1919
		_go_fuzz_dep_.CoverTab[526509]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1919
		_go_fuzz_dep_.CoverTab[3266]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1919
		// _ = "end of CoverTab[3266]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1919
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1919
	// _ = "end of CoverTab[3263]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1919
	_go_fuzz_dep_.CoverTab[3264]++
											return n
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1920
	// _ = "end of CoverTab[3264]"
}

// String implements fmt.Stringer.String.
func (n Name) String() string {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1924
	_go_fuzz_dep_.CoverTab[3267]++
											return string(n.Data[:n.Length])
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1925
	// _ = "end of CoverTab[3267]"
}

// GoString implements fmt.GoStringer.GoString.
func (n *Name) GoString() string {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1929
	_go_fuzz_dep_.CoverTab[3268]++
											return `dnsmessage.MustNewName("` + printString(n.Data[:n.Length]) + `")`
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1930
	// _ = "end of CoverTab[3268]"
}

// pack appends the wire format of the Name to msg.
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1933
//
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1933
// Domain names are a sequence of counted strings split at the dots. They end
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1933
// with a zero-length string. Compression can be used to reuse domain suffixes.
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1933
//
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1933
// The compression map will be updated with new domain suffixes. If compression
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1933
// is nil, compression will not be used.
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1940
func (n *Name) pack(msg []byte, compression map[string]int, compressionOff int) ([]byte, error) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1940
	_go_fuzz_dep_.CoverTab[3269]++
											oldMsg := msg

											if n.Length > nonEncodedNameMax {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1943
		_go_fuzz_dep_.CoverTab[526510]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1943
		_go_fuzz_dep_.CoverTab[3274]++
												return nil, errNameTooLong
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1944
		// _ = "end of CoverTab[3274]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1945
		_go_fuzz_dep_.CoverTab[526511]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1945
		_go_fuzz_dep_.CoverTab[3275]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1945
		// _ = "end of CoverTab[3275]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1945
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1945
	// _ = "end of CoverTab[3269]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1945
	_go_fuzz_dep_.CoverTab[3270]++

//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1948
	if n.Length == 0 || func() bool {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1948
		_go_fuzz_dep_.CoverTab[3276]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1948
		return n.Data[n.Length-1] != '.'
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1948
		// _ = "end of CoverTab[3276]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1948
	}() {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1948
		_go_fuzz_dep_.CoverTab[526512]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1948
		_go_fuzz_dep_.CoverTab[3277]++
												return oldMsg, errNonCanonicalName
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1949
		// _ = "end of CoverTab[3277]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1950
		_go_fuzz_dep_.CoverTab[526513]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1950
		_go_fuzz_dep_.CoverTab[3278]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1950
		// _ = "end of CoverTab[3278]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1950
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1950
	// _ = "end of CoverTab[3270]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1950
	_go_fuzz_dep_.CoverTab[3271]++

//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1953
	if n.Data[0] == '.' && func() bool {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1953
		_go_fuzz_dep_.CoverTab[3279]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1953
		return n.Length == 1
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1953
		// _ = "end of CoverTab[3279]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1953
	}() {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1953
		_go_fuzz_dep_.CoverTab[526514]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1953
		_go_fuzz_dep_.CoverTab[3280]++
												return append(msg, 0), nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1954
		// _ = "end of CoverTab[3280]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1955
		_go_fuzz_dep_.CoverTab[526515]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1955
		_go_fuzz_dep_.CoverTab[3281]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1955
		// _ = "end of CoverTab[3281]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1955
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1955
	// _ = "end of CoverTab[3271]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1955
	_go_fuzz_dep_.CoverTab[3272]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1955
	_go_fuzz_dep_.CoverTab[786593] = 0

//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1958
	for i, begin := 0, 0; i < int(n.Length); i++ {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1958
		if _go_fuzz_dep_.CoverTab[786593] == 0 {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1958
			_go_fuzz_dep_.CoverTab[526715]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1958
		} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1958
			_go_fuzz_dep_.CoverTab[526716]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1958
		}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1958
		_go_fuzz_dep_.CoverTab[786593] = 1
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1958
		_go_fuzz_dep_.CoverTab[3282]++

												if n.Data[i] == '.' {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1960
			_go_fuzz_dep_.CoverTab[526516]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1960
			_go_fuzz_dep_.CoverTab[3284]++

//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1964
			if i-begin >= 1<<6 {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1964
				_go_fuzz_dep_.CoverTab[526518]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1964
				_go_fuzz_dep_.CoverTab[3288]++
														return oldMsg, errSegTooLong
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1965
				// _ = "end of CoverTab[3288]"
			} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1966
				_go_fuzz_dep_.CoverTab[526519]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1966
				_go_fuzz_dep_.CoverTab[3289]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1966
				// _ = "end of CoverTab[3289]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1966
			}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1966
			// _ = "end of CoverTab[3284]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1966
			_go_fuzz_dep_.CoverTab[3285]++

//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1969
			if i-begin == 0 {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1969
				_go_fuzz_dep_.CoverTab[526520]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1969
				_go_fuzz_dep_.CoverTab[3290]++
														return oldMsg, errZeroSegLen
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1970
				// _ = "end of CoverTab[3290]"
			} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1971
				_go_fuzz_dep_.CoverTab[526521]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1971
				_go_fuzz_dep_.CoverTab[3291]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1971
				// _ = "end of CoverTab[3291]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1971
			}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1971
			// _ = "end of CoverTab[3285]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1971
			_go_fuzz_dep_.CoverTab[3286]++

													msg = append(msg, byte(i-begin))
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1973
			_go_fuzz_dep_.CoverTab[786594] = 0

													for j := begin; j < i; j++ {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1975
				if _go_fuzz_dep_.CoverTab[786594] == 0 {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1975
					_go_fuzz_dep_.CoverTab[526719]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1975
				} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1975
					_go_fuzz_dep_.CoverTab[526720]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1975
				}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1975
				_go_fuzz_dep_.CoverTab[786594] = 1
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1975
				_go_fuzz_dep_.CoverTab[3292]++
														msg = append(msg, n.Data[j])
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1976
				// _ = "end of CoverTab[3292]"
			}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1977
			if _go_fuzz_dep_.CoverTab[786594] == 0 {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1977
				_go_fuzz_dep_.CoverTab[526721]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1977
			} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1977
				_go_fuzz_dep_.CoverTab[526722]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1977
			}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1977
			// _ = "end of CoverTab[3286]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1977
			_go_fuzz_dep_.CoverTab[3287]++

													begin = i + 1
													continue
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1980
			// _ = "end of CoverTab[3287]"
		} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1981
			_go_fuzz_dep_.CoverTab[526517]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1981
			_go_fuzz_dep_.CoverTab[3293]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1981
			// _ = "end of CoverTab[3293]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1981
		}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1981
		// _ = "end of CoverTab[3282]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1981
		_go_fuzz_dep_.CoverTab[3283]++

//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1986
		if (i == 0 || func() bool {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1986
			_go_fuzz_dep_.CoverTab[3294]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1986
			return n.Data[i-1] == '.'
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1986
			// _ = "end of CoverTab[3294]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1986
		}()) && func() bool {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1986
			_go_fuzz_dep_.CoverTab[3295]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1986
			return compression != nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1986
			// _ = "end of CoverTab[3295]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1986
		}() {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1986
			_go_fuzz_dep_.CoverTab[526522]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1986
			_go_fuzz_dep_.CoverTab[3296]++
													if ptr, ok := compression[string(n.Data[i:])]; ok {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1987
				_go_fuzz_dep_.CoverTab[526524]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1987
				_go_fuzz_dep_.CoverTab[3298]++

//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1990
				return append(msg, byte(ptr>>8|0xC0), byte(ptr)), nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1990
				// _ = "end of CoverTab[3298]"
			} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1991
				_go_fuzz_dep_.CoverTab[526525]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1991
				_go_fuzz_dep_.CoverTab[3299]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1991
				// _ = "end of CoverTab[3299]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1991
			}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1991
			// _ = "end of CoverTab[3296]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1991
			_go_fuzz_dep_.CoverTab[3297]++

//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1995
			if len(msg) <= int(^uint16(0)>>2) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1995
				_go_fuzz_dep_.CoverTab[526526]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1995
				_go_fuzz_dep_.CoverTab[3300]++
														compression[string(n.Data[i:])] = len(msg) - compressionOff
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1996
				// _ = "end of CoverTab[3300]"
			} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1997
				_go_fuzz_dep_.CoverTab[526527]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1997
				_go_fuzz_dep_.CoverTab[3301]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1997
				// _ = "end of CoverTab[3301]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1997
			}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1997
			// _ = "end of CoverTab[3297]"
		} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1998
			_go_fuzz_dep_.CoverTab[526523]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1998
			_go_fuzz_dep_.CoverTab[3302]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1998
			// _ = "end of CoverTab[3302]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1998
		}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1998
		// _ = "end of CoverTab[3283]"
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1999
	if _go_fuzz_dep_.CoverTab[786593] == 0 {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1999
		_go_fuzz_dep_.CoverTab[526717]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1999
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1999
		_go_fuzz_dep_.CoverTab[526718]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1999
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1999
	// _ = "end of CoverTab[3272]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1999
	_go_fuzz_dep_.CoverTab[3273]++
											return append(msg, 0), nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2000
	// _ = "end of CoverTab[3273]"
}

// unpack unpacks a domain name.
func (n *Name) unpack(msg []byte, off int) (int, error) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2004
	_go_fuzz_dep_.CoverTab[3303]++
											return n.unpackCompressed(msg, off, true)
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2005
	// _ = "end of CoverTab[3303]"
}

func (n *Name) unpackCompressed(msg []byte, off int, allowCompression bool) (int, error) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2008
	_go_fuzz_dep_.CoverTab[3304]++

											currOff := off

//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2015
	newOff := off

											// ptr is the number of pointers followed.
											var ptr int

//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2021
	name := n.Data[:0]

Loop:
	for {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2024
		_go_fuzz_dep_.CoverTab[3309]++
												if currOff >= len(msg) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2025
			_go_fuzz_dep_.CoverTab[526528]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2025
			_go_fuzz_dep_.CoverTab[3311]++
													return off, errBaseLen
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2026
			// _ = "end of CoverTab[3311]"
		} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2027
			_go_fuzz_dep_.CoverTab[526529]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2027
			_go_fuzz_dep_.CoverTab[3312]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2027
			// _ = "end of CoverTab[3312]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2027
		}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2027
		// _ = "end of CoverTab[3309]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2027
		_go_fuzz_dep_.CoverTab[3310]++
												c := int(msg[currOff])
												currOff++
												switch c & 0xC0 {
		case 0x00:
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2031
			_go_fuzz_dep_.CoverTab[526530]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2031
			_go_fuzz_dep_.CoverTab[3313]++
													if c == 0x00 {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2032
				_go_fuzz_dep_.CoverTab[526533]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2032
				_go_fuzz_dep_.CoverTab[3323]++

														break Loop
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2034
				// _ = "end of CoverTab[3323]"
			} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2035
				_go_fuzz_dep_.CoverTab[526534]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2035
				_go_fuzz_dep_.CoverTab[3324]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2035
				// _ = "end of CoverTab[3324]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2035
			}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2035
			// _ = "end of CoverTab[3313]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2035
			_go_fuzz_dep_.CoverTab[3314]++
													endOff := currOff + c
													if endOff > len(msg) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2037
				_go_fuzz_dep_.CoverTab[526535]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2037
				_go_fuzz_dep_.CoverTab[3325]++
														return off, errCalcLen
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2038
				// _ = "end of CoverTab[3325]"
			} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2039
				_go_fuzz_dep_.CoverTab[526536]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2039
				_go_fuzz_dep_.CoverTab[3326]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2039
				// _ = "end of CoverTab[3326]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2039
			}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2039
			// _ = "end of CoverTab[3314]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2039
			_go_fuzz_dep_.CoverTab[3315]++

//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2043
			for _, v := range msg[currOff:endOff] {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2043
				_go_fuzz_dep_.CoverTab[3327]++
														if v == '.' {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2044
					_go_fuzz_dep_.CoverTab[526537]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2044
					_go_fuzz_dep_.CoverTab[3328]++
															return off, errInvalidName
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2045
					// _ = "end of CoverTab[3328]"
				} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2046
					_go_fuzz_dep_.CoverTab[526538]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2046
					_go_fuzz_dep_.CoverTab[3329]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2046
					// _ = "end of CoverTab[3329]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2046
				}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2046
				// _ = "end of CoverTab[3327]"
			}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2047
			// _ = "end of CoverTab[3315]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2047
			_go_fuzz_dep_.CoverTab[3316]++

													name = append(name, msg[currOff:endOff]...)
													name = append(name, '.')
													currOff = endOff
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2051
			// _ = "end of CoverTab[3316]"
		case 0xC0:
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2052
			_go_fuzz_dep_.CoverTab[526531]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2052
			_go_fuzz_dep_.CoverTab[3317]++
													if !allowCompression {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2053
				_go_fuzz_dep_.CoverTab[526539]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2053
				_go_fuzz_dep_.CoverTab[3330]++
														return off, errCompressedSRV
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2054
				// _ = "end of CoverTab[3330]"
			} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2055
				_go_fuzz_dep_.CoverTab[526540]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2055
				_go_fuzz_dep_.CoverTab[3331]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2055
				// _ = "end of CoverTab[3331]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2055
			}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2055
			// _ = "end of CoverTab[3317]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2055
			_go_fuzz_dep_.CoverTab[3318]++
													if currOff >= len(msg) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2056
				_go_fuzz_dep_.CoverTab[526541]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2056
				_go_fuzz_dep_.CoverTab[3332]++
														return off, errInvalidPtr
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2057
				// _ = "end of CoverTab[3332]"
			} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2058
				_go_fuzz_dep_.CoverTab[526542]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2058
				_go_fuzz_dep_.CoverTab[3333]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2058
				// _ = "end of CoverTab[3333]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2058
			}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2058
			// _ = "end of CoverTab[3318]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2058
			_go_fuzz_dep_.CoverTab[3319]++
													c1 := msg[currOff]
													currOff++
													if ptr == 0 {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2061
				_go_fuzz_dep_.CoverTab[526543]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2061
				_go_fuzz_dep_.CoverTab[3334]++
														newOff = currOff
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2062
				// _ = "end of CoverTab[3334]"
			} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2063
				_go_fuzz_dep_.CoverTab[526544]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2063
				_go_fuzz_dep_.CoverTab[3335]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2063
				// _ = "end of CoverTab[3335]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2063
			}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2063
			// _ = "end of CoverTab[3319]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2063
			_go_fuzz_dep_.CoverTab[3320]++

													if ptr++; ptr > 10 {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2065
				_go_fuzz_dep_.CoverTab[526545]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2065
				_go_fuzz_dep_.CoverTab[3336]++
														return off, errTooManyPtr
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2066
				// _ = "end of CoverTab[3336]"
			} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2067
				_go_fuzz_dep_.CoverTab[526546]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2067
				_go_fuzz_dep_.CoverTab[3337]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2067
				// _ = "end of CoverTab[3337]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2067
			}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2067
			// _ = "end of CoverTab[3320]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2067
			_go_fuzz_dep_.CoverTab[3321]++
													currOff = (c^0xC0)<<8 | int(c1)
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2068
			// _ = "end of CoverTab[3321]"
		default:
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2069
			_go_fuzz_dep_.CoverTab[526532]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2069
			_go_fuzz_dep_.CoverTab[3322]++

													return off, errReserved
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2071
			// _ = "end of CoverTab[3322]"
		}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2072
		// _ = "end of CoverTab[3310]"
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2073
	// _ = "end of CoverTab[3304]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2073
	_go_fuzz_dep_.CoverTab[3305]++
											if len(name) == 0 {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2074
		_go_fuzz_dep_.CoverTab[526547]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2074
		_go_fuzz_dep_.CoverTab[3338]++
												name = append(name, '.')
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2075
		// _ = "end of CoverTab[3338]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2076
		_go_fuzz_dep_.CoverTab[526548]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2076
		_go_fuzz_dep_.CoverTab[3339]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2076
		// _ = "end of CoverTab[3339]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2076
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2076
	// _ = "end of CoverTab[3305]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2076
	_go_fuzz_dep_.CoverTab[3306]++
											if len(name) > nonEncodedNameMax {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2077
		_go_fuzz_dep_.CoverTab[526549]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2077
		_go_fuzz_dep_.CoverTab[3340]++
												return off, errNameTooLong
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2078
		// _ = "end of CoverTab[3340]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2079
		_go_fuzz_dep_.CoverTab[526550]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2079
		_go_fuzz_dep_.CoverTab[3341]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2079
		// _ = "end of CoverTab[3341]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2079
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2079
	// _ = "end of CoverTab[3306]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2079
	_go_fuzz_dep_.CoverTab[3307]++
											n.Length = uint8(len(name))
											if ptr == 0 {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2081
		_go_fuzz_dep_.CoverTab[526551]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2081
		_go_fuzz_dep_.CoverTab[3342]++
												newOff = currOff
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2082
		// _ = "end of CoverTab[3342]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2083
		_go_fuzz_dep_.CoverTab[526552]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2083
		_go_fuzz_dep_.CoverTab[3343]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2083
		// _ = "end of CoverTab[3343]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2083
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2083
	// _ = "end of CoverTab[3307]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2083
	_go_fuzz_dep_.CoverTab[3308]++
											return newOff, nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2084
	// _ = "end of CoverTab[3308]"
}

func skipName(msg []byte, off int) (int, error) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2087
	_go_fuzz_dep_.CoverTab[3344]++

//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2091
	newOff := off

Loop:
	for {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2094
		_go_fuzz_dep_.CoverTab[3346]++
												if newOff >= len(msg) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2095
			_go_fuzz_dep_.CoverTab[526553]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2095
			_go_fuzz_dep_.CoverTab[3348]++
													return off, errBaseLen
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2096
			// _ = "end of CoverTab[3348]"
		} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2097
			_go_fuzz_dep_.CoverTab[526554]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2097
			_go_fuzz_dep_.CoverTab[3349]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2097
			// _ = "end of CoverTab[3349]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2097
		}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2097
		// _ = "end of CoverTab[3346]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2097
		_go_fuzz_dep_.CoverTab[3347]++
												c := int(msg[newOff])
												newOff++
												switch c & 0xC0 {
		case 0x00:
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2101
			_go_fuzz_dep_.CoverTab[526555]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2101
			_go_fuzz_dep_.CoverTab[3350]++
													if c == 0x00 {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2102
				_go_fuzz_dep_.CoverTab[526558]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2102
				_go_fuzz_dep_.CoverTab[3354]++

														break Loop
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2104
				// _ = "end of CoverTab[3354]"
			} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2105
				_go_fuzz_dep_.CoverTab[526559]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2105
				_go_fuzz_dep_.CoverTab[3355]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2105
				// _ = "end of CoverTab[3355]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2105
			}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2105
			// _ = "end of CoverTab[3350]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2105
			_go_fuzz_dep_.CoverTab[3351]++

													newOff += c
													if newOff > len(msg) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2108
				_go_fuzz_dep_.CoverTab[526560]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2108
				_go_fuzz_dep_.CoverTab[3356]++
														return off, errCalcLen
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2109
				// _ = "end of CoverTab[3356]"
			} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2110
				_go_fuzz_dep_.CoverTab[526561]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2110
				_go_fuzz_dep_.CoverTab[3357]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2110
				// _ = "end of CoverTab[3357]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2110
			}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2110
			// _ = "end of CoverTab[3351]"
		case 0xC0:
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2111
			_go_fuzz_dep_.CoverTab[526556]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2111
			_go_fuzz_dep_.CoverTab[3352]++

//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2115
			newOff++

//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2118
			break Loop
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2118
			// _ = "end of CoverTab[3352]"
		default:
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2119
			_go_fuzz_dep_.CoverTab[526557]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2119
			_go_fuzz_dep_.CoverTab[3353]++

													return off, errReserved
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2121
			// _ = "end of CoverTab[3353]"
		}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2122
		// _ = "end of CoverTab[3347]"
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2123
	// _ = "end of CoverTab[3344]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2123
	_go_fuzz_dep_.CoverTab[3345]++

											return newOff, nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2125
	// _ = "end of CoverTab[3345]"
}

// A Question is a DNS query.
type Question struct {
	Name	Name
	Type	Type
	Class	Class
}

// pack appends the wire format of the Question to msg.
func (q *Question) pack(msg []byte, compression map[string]int, compressionOff int) ([]byte, error) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2136
	_go_fuzz_dep_.CoverTab[3358]++
											msg, err := q.Name.pack(msg, compression, compressionOff)
											if err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2138
		_go_fuzz_dep_.CoverTab[526562]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2138
		_go_fuzz_dep_.CoverTab[3360]++
												return msg, &nestedError{"Name", err}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2139
		// _ = "end of CoverTab[3360]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2140
		_go_fuzz_dep_.CoverTab[526563]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2140
		_go_fuzz_dep_.CoverTab[3361]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2140
		// _ = "end of CoverTab[3361]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2140
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2140
	// _ = "end of CoverTab[3358]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2140
	_go_fuzz_dep_.CoverTab[3359]++
											msg = packType(msg, q.Type)
											return packClass(msg, q.Class), nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2142
	// _ = "end of CoverTab[3359]"
}

// GoString implements fmt.GoStringer.GoString.
func (q *Question) GoString() string {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2146
	_go_fuzz_dep_.CoverTab[3362]++
											return "dnsmessage.Question{" +
		"Name: " + q.Name.GoString() + ", " +
		"Type: " + q.Type.GoString() + ", " +
		"Class: " + q.Class.GoString() + "}"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2150
	// _ = "end of CoverTab[3362]"
}

func unpackResourceBody(msg []byte, off int, hdr ResourceHeader) (ResourceBody, int, error) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2153
	_go_fuzz_dep_.CoverTab[3363]++
											var (
		r	ResourceBody
		err	error
		name	string
	)
	switch hdr.Type {
	case TypeA:
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2160
		_go_fuzz_dep_.CoverTab[526564]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2160
		_go_fuzz_dep_.CoverTab[3366]++
												var rb AResource
												rb, err = unpackAResource(msg, off)
												r = &rb
												name = "A"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2164
		// _ = "end of CoverTab[3366]"
	case TypeNS:
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2165
		_go_fuzz_dep_.CoverTab[526565]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2165
		_go_fuzz_dep_.CoverTab[3367]++
												var rb NSResource
												rb, err = unpackNSResource(msg, off)
												r = &rb
												name = "NS"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2169
		// _ = "end of CoverTab[3367]"
	case TypeCNAME:
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2170
		_go_fuzz_dep_.CoverTab[526566]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2170
		_go_fuzz_dep_.CoverTab[3368]++
												var rb CNAMEResource
												rb, err = unpackCNAMEResource(msg, off)
												r = &rb
												name = "CNAME"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2174
		// _ = "end of CoverTab[3368]"
	case TypeSOA:
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2175
		_go_fuzz_dep_.CoverTab[526567]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2175
		_go_fuzz_dep_.CoverTab[3369]++
												var rb SOAResource
												rb, err = unpackSOAResource(msg, off)
												r = &rb
												name = "SOA"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2179
		// _ = "end of CoverTab[3369]"
	case TypePTR:
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2180
		_go_fuzz_dep_.CoverTab[526568]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2180
		_go_fuzz_dep_.CoverTab[3370]++
												var rb PTRResource
												rb, err = unpackPTRResource(msg, off)
												r = &rb
												name = "PTR"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2184
		// _ = "end of CoverTab[3370]"
	case TypeMX:
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2185
		_go_fuzz_dep_.CoverTab[526569]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2185
		_go_fuzz_dep_.CoverTab[3371]++
												var rb MXResource
												rb, err = unpackMXResource(msg, off)
												r = &rb
												name = "MX"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2189
		// _ = "end of CoverTab[3371]"
	case TypeTXT:
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2190
		_go_fuzz_dep_.CoverTab[526570]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2190
		_go_fuzz_dep_.CoverTab[3372]++
												var rb TXTResource
												rb, err = unpackTXTResource(msg, off, hdr.Length)
												r = &rb
												name = "TXT"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2194
		// _ = "end of CoverTab[3372]"
	case TypeAAAA:
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2195
		_go_fuzz_dep_.CoverTab[526571]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2195
		_go_fuzz_dep_.CoverTab[3373]++
												var rb AAAAResource
												rb, err = unpackAAAAResource(msg, off)
												r = &rb
												name = "AAAA"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2199
		// _ = "end of CoverTab[3373]"
	case TypeSRV:
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2200
		_go_fuzz_dep_.CoverTab[526572]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2200
		_go_fuzz_dep_.CoverTab[3374]++
												var rb SRVResource
												rb, err = unpackSRVResource(msg, off)
												r = &rb
												name = "SRV"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2204
		// _ = "end of CoverTab[3374]"
	case TypeOPT:
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2205
		_go_fuzz_dep_.CoverTab[526573]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2205
		_go_fuzz_dep_.CoverTab[3375]++
												var rb OPTResource
												rb, err = unpackOPTResource(msg, off, hdr.Length)
												r = &rb
												name = "OPT"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2209
		// _ = "end of CoverTab[3375]"
	default:
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2210
		_go_fuzz_dep_.CoverTab[526574]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2210
		_go_fuzz_dep_.CoverTab[3376]++
												var rb UnknownResource
												rb, err = unpackUnknownResource(hdr.Type, msg, off, hdr.Length)
												r = &rb
												name = "Unknown"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2214
		// _ = "end of CoverTab[3376]"
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2215
	// _ = "end of CoverTab[3363]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2215
	_go_fuzz_dep_.CoverTab[3364]++
											if err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2216
		_go_fuzz_dep_.CoverTab[526575]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2216
		_go_fuzz_dep_.CoverTab[3377]++
												return nil, off, &nestedError{name + " record", err}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2217
		// _ = "end of CoverTab[3377]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2218
		_go_fuzz_dep_.CoverTab[526576]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2218
		_go_fuzz_dep_.CoverTab[3378]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2218
		// _ = "end of CoverTab[3378]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2218
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2218
	// _ = "end of CoverTab[3364]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2218
	_go_fuzz_dep_.CoverTab[3365]++
											return r, off + int(hdr.Length), nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2219
	// _ = "end of CoverTab[3365]"
}

// A CNAMEResource is a CNAME Resource record.
type CNAMEResource struct {
	CNAME Name
}

func (r *CNAMEResource) realType() Type {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2227
	_go_fuzz_dep_.CoverTab[3379]++
											return TypeCNAME
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2228
	// _ = "end of CoverTab[3379]"
}

// pack appends the wire format of the CNAMEResource to msg.
func (r *CNAMEResource) pack(msg []byte, compression map[string]int, compressionOff int) ([]byte, error) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2232
	_go_fuzz_dep_.CoverTab[3380]++
											return r.CNAME.pack(msg, compression, compressionOff)
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2233
	// _ = "end of CoverTab[3380]"
}

// GoString implements fmt.GoStringer.GoString.
func (r *CNAMEResource) GoString() string {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2237
	_go_fuzz_dep_.CoverTab[3381]++
											return "dnsmessage.CNAMEResource{CNAME: " + r.CNAME.GoString() + "}"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2238
	// _ = "end of CoverTab[3381]"
}

func unpackCNAMEResource(msg []byte, off int) (CNAMEResource, error) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2241
	_go_fuzz_dep_.CoverTab[3382]++
											var cname Name
											if _, err := cname.unpack(msg, off); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2243
		_go_fuzz_dep_.CoverTab[526577]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2243
		_go_fuzz_dep_.CoverTab[3384]++
												return CNAMEResource{}, err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2244
		// _ = "end of CoverTab[3384]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2245
		_go_fuzz_dep_.CoverTab[526578]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2245
		_go_fuzz_dep_.CoverTab[3385]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2245
		// _ = "end of CoverTab[3385]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2245
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2245
	// _ = "end of CoverTab[3382]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2245
	_go_fuzz_dep_.CoverTab[3383]++
											return CNAMEResource{cname}, nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2246
	// _ = "end of CoverTab[3383]"
}

// An MXResource is an MX Resource record.
type MXResource struct {
	Pref	uint16
	MX	Name
}

func (r *MXResource) realType() Type {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2255
	_go_fuzz_dep_.CoverTab[3386]++
											return TypeMX
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2256
	// _ = "end of CoverTab[3386]"
}

// pack appends the wire format of the MXResource to msg.
func (r *MXResource) pack(msg []byte, compression map[string]int, compressionOff int) ([]byte, error) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2260
	_go_fuzz_dep_.CoverTab[3387]++
											oldMsg := msg
											msg = packUint16(msg, r.Pref)
											msg, err := r.MX.pack(msg, compression, compressionOff)
											if err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2264
		_go_fuzz_dep_.CoverTab[526579]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2264
		_go_fuzz_dep_.CoverTab[3389]++
												return oldMsg, &nestedError{"MXResource.MX", err}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2265
		// _ = "end of CoverTab[3389]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2266
		_go_fuzz_dep_.CoverTab[526580]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2266
		_go_fuzz_dep_.CoverTab[3390]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2266
		// _ = "end of CoverTab[3390]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2266
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2266
	// _ = "end of CoverTab[3387]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2266
	_go_fuzz_dep_.CoverTab[3388]++
											return msg, nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2267
	// _ = "end of CoverTab[3388]"
}

// GoString implements fmt.GoStringer.GoString.
func (r *MXResource) GoString() string {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2271
	_go_fuzz_dep_.CoverTab[3391]++
											return "dnsmessage.MXResource{" +
		"Pref: " + printUint16(r.Pref) + ", " +
		"MX: " + r.MX.GoString() + "}"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2274
	// _ = "end of CoverTab[3391]"
}

func unpackMXResource(msg []byte, off int) (MXResource, error) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2277
	_go_fuzz_dep_.CoverTab[3392]++
											pref, off, err := unpackUint16(msg, off)
											if err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2279
		_go_fuzz_dep_.CoverTab[526581]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2279
		_go_fuzz_dep_.CoverTab[3395]++
												return MXResource{}, &nestedError{"Pref", err}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2280
		// _ = "end of CoverTab[3395]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2281
		_go_fuzz_dep_.CoverTab[526582]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2281
		_go_fuzz_dep_.CoverTab[3396]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2281
		// _ = "end of CoverTab[3396]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2281
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2281
	// _ = "end of CoverTab[3392]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2281
	_go_fuzz_dep_.CoverTab[3393]++
											var mx Name
											if _, err := mx.unpack(msg, off); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2283
		_go_fuzz_dep_.CoverTab[526583]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2283
		_go_fuzz_dep_.CoverTab[3397]++
												return MXResource{}, &nestedError{"MX", err}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2284
		// _ = "end of CoverTab[3397]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2285
		_go_fuzz_dep_.CoverTab[526584]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2285
		_go_fuzz_dep_.CoverTab[3398]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2285
		// _ = "end of CoverTab[3398]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2285
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2285
	// _ = "end of CoverTab[3393]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2285
	_go_fuzz_dep_.CoverTab[3394]++
											return MXResource{pref, mx}, nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2286
	// _ = "end of CoverTab[3394]"
}

// An NSResource is an NS Resource record.
type NSResource struct {
	NS Name
}

func (r *NSResource) realType() Type {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2294
	_go_fuzz_dep_.CoverTab[3399]++
											return TypeNS
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2295
	// _ = "end of CoverTab[3399]"
}

// pack appends the wire format of the NSResource to msg.
func (r *NSResource) pack(msg []byte, compression map[string]int, compressionOff int) ([]byte, error) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2299
	_go_fuzz_dep_.CoverTab[3400]++
											return r.NS.pack(msg, compression, compressionOff)
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2300
	// _ = "end of CoverTab[3400]"
}

// GoString implements fmt.GoStringer.GoString.
func (r *NSResource) GoString() string {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2304
	_go_fuzz_dep_.CoverTab[3401]++
											return "dnsmessage.NSResource{NS: " + r.NS.GoString() + "}"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2305
	// _ = "end of CoverTab[3401]"
}

func unpackNSResource(msg []byte, off int) (NSResource, error) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2308
	_go_fuzz_dep_.CoverTab[3402]++
											var ns Name
											if _, err := ns.unpack(msg, off); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2310
		_go_fuzz_dep_.CoverTab[526585]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2310
		_go_fuzz_dep_.CoverTab[3404]++
												return NSResource{}, err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2311
		// _ = "end of CoverTab[3404]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2312
		_go_fuzz_dep_.CoverTab[526586]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2312
		_go_fuzz_dep_.CoverTab[3405]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2312
		// _ = "end of CoverTab[3405]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2312
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2312
	// _ = "end of CoverTab[3402]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2312
	_go_fuzz_dep_.CoverTab[3403]++
											return NSResource{ns}, nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2313
	// _ = "end of CoverTab[3403]"
}

// A PTRResource is a PTR Resource record.
type PTRResource struct {
	PTR Name
}

func (r *PTRResource) realType() Type {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2321
	_go_fuzz_dep_.CoverTab[3406]++
											return TypePTR
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2322
	// _ = "end of CoverTab[3406]"
}

// pack appends the wire format of the PTRResource to msg.
func (r *PTRResource) pack(msg []byte, compression map[string]int, compressionOff int) ([]byte, error) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2326
	_go_fuzz_dep_.CoverTab[3407]++
											return r.PTR.pack(msg, compression, compressionOff)
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2327
	// _ = "end of CoverTab[3407]"
}

// GoString implements fmt.GoStringer.GoString.
func (r *PTRResource) GoString() string {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2331
	_go_fuzz_dep_.CoverTab[3408]++
											return "dnsmessage.PTRResource{PTR: " + r.PTR.GoString() + "}"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2332
	// _ = "end of CoverTab[3408]"
}

func unpackPTRResource(msg []byte, off int) (PTRResource, error) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2335
	_go_fuzz_dep_.CoverTab[3409]++
											var ptr Name
											if _, err := ptr.unpack(msg, off); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2337
		_go_fuzz_dep_.CoverTab[526587]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2337
		_go_fuzz_dep_.CoverTab[3411]++
												return PTRResource{}, err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2338
		// _ = "end of CoverTab[3411]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2339
		_go_fuzz_dep_.CoverTab[526588]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2339
		_go_fuzz_dep_.CoverTab[3412]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2339
		// _ = "end of CoverTab[3412]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2339
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2339
	// _ = "end of CoverTab[3409]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2339
	_go_fuzz_dep_.CoverTab[3410]++
											return PTRResource{ptr}, nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2340
	// _ = "end of CoverTab[3410]"
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
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2358
	_go_fuzz_dep_.CoverTab[3413]++
											return TypeSOA
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2359
	// _ = "end of CoverTab[3413]"
}

// pack appends the wire format of the SOAResource to msg.
func (r *SOAResource) pack(msg []byte, compression map[string]int, compressionOff int) ([]byte, error) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2363
	_go_fuzz_dep_.CoverTab[3414]++
											oldMsg := msg
											msg, err := r.NS.pack(msg, compression, compressionOff)
											if err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2366
		_go_fuzz_dep_.CoverTab[526589]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2366
		_go_fuzz_dep_.CoverTab[3417]++
												return oldMsg, &nestedError{"SOAResource.NS", err}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2367
		// _ = "end of CoverTab[3417]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2368
		_go_fuzz_dep_.CoverTab[526590]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2368
		_go_fuzz_dep_.CoverTab[3418]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2368
		// _ = "end of CoverTab[3418]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2368
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2368
	// _ = "end of CoverTab[3414]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2368
	_go_fuzz_dep_.CoverTab[3415]++
											msg, err = r.MBox.pack(msg, compression, compressionOff)
											if err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2370
		_go_fuzz_dep_.CoverTab[526591]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2370
		_go_fuzz_dep_.CoverTab[3419]++
												return oldMsg, &nestedError{"SOAResource.MBox", err}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2371
		// _ = "end of CoverTab[3419]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2372
		_go_fuzz_dep_.CoverTab[526592]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2372
		_go_fuzz_dep_.CoverTab[3420]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2372
		// _ = "end of CoverTab[3420]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2372
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2372
	// _ = "end of CoverTab[3415]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2372
	_go_fuzz_dep_.CoverTab[3416]++
											msg = packUint32(msg, r.Serial)
											msg = packUint32(msg, r.Refresh)
											msg = packUint32(msg, r.Retry)
											msg = packUint32(msg, r.Expire)
											return packUint32(msg, r.MinTTL), nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2377
	// _ = "end of CoverTab[3416]"
}

// GoString implements fmt.GoStringer.GoString.
func (r *SOAResource) GoString() string {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2381
	_go_fuzz_dep_.CoverTab[3421]++
											return "dnsmessage.SOAResource{" +
		"NS: " + r.NS.GoString() + ", " +
		"MBox: " + r.MBox.GoString() + ", " +
		"Serial: " + printUint32(r.Serial) + ", " +
		"Refresh: " + printUint32(r.Refresh) + ", " +
		"Retry: " + printUint32(r.Retry) + ", " +
		"Expire: " + printUint32(r.Expire) + ", " +
		"MinTTL: " + printUint32(r.MinTTL) + "}"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2389
	// _ = "end of CoverTab[3421]"
}

func unpackSOAResource(msg []byte, off int) (SOAResource, error) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2392
	_go_fuzz_dep_.CoverTab[3422]++
											var ns Name
											off, err := ns.unpack(msg, off)
											if err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2395
		_go_fuzz_dep_.CoverTab[526593]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2395
		_go_fuzz_dep_.CoverTab[3430]++
												return SOAResource{}, &nestedError{"NS", err}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2396
		// _ = "end of CoverTab[3430]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2397
		_go_fuzz_dep_.CoverTab[526594]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2397
		_go_fuzz_dep_.CoverTab[3431]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2397
		// _ = "end of CoverTab[3431]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2397
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2397
	// _ = "end of CoverTab[3422]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2397
	_go_fuzz_dep_.CoverTab[3423]++
											var mbox Name
											if off, err = mbox.unpack(msg, off); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2399
		_go_fuzz_dep_.CoverTab[526595]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2399
		_go_fuzz_dep_.CoverTab[3432]++
												return SOAResource{}, &nestedError{"MBox", err}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2400
		// _ = "end of CoverTab[3432]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2401
		_go_fuzz_dep_.CoverTab[526596]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2401
		_go_fuzz_dep_.CoverTab[3433]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2401
		// _ = "end of CoverTab[3433]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2401
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2401
	// _ = "end of CoverTab[3423]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2401
	_go_fuzz_dep_.CoverTab[3424]++
											serial, off, err := unpackUint32(msg, off)
											if err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2403
		_go_fuzz_dep_.CoverTab[526597]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2403
		_go_fuzz_dep_.CoverTab[3434]++
												return SOAResource{}, &nestedError{"Serial", err}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2404
		// _ = "end of CoverTab[3434]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2405
		_go_fuzz_dep_.CoverTab[526598]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2405
		_go_fuzz_dep_.CoverTab[3435]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2405
		// _ = "end of CoverTab[3435]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2405
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2405
	// _ = "end of CoverTab[3424]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2405
	_go_fuzz_dep_.CoverTab[3425]++
											refresh, off, err := unpackUint32(msg, off)
											if err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2407
		_go_fuzz_dep_.CoverTab[526599]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2407
		_go_fuzz_dep_.CoverTab[3436]++
												return SOAResource{}, &nestedError{"Refresh", err}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2408
		// _ = "end of CoverTab[3436]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2409
		_go_fuzz_dep_.CoverTab[526600]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2409
		_go_fuzz_dep_.CoverTab[3437]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2409
		// _ = "end of CoverTab[3437]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2409
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2409
	// _ = "end of CoverTab[3425]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2409
	_go_fuzz_dep_.CoverTab[3426]++
											retry, off, err := unpackUint32(msg, off)
											if err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2411
		_go_fuzz_dep_.CoverTab[526601]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2411
		_go_fuzz_dep_.CoverTab[3438]++
												return SOAResource{}, &nestedError{"Retry", err}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2412
		// _ = "end of CoverTab[3438]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2413
		_go_fuzz_dep_.CoverTab[526602]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2413
		_go_fuzz_dep_.CoverTab[3439]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2413
		// _ = "end of CoverTab[3439]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2413
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2413
	// _ = "end of CoverTab[3426]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2413
	_go_fuzz_dep_.CoverTab[3427]++
											expire, off, err := unpackUint32(msg, off)
											if err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2415
		_go_fuzz_dep_.CoverTab[526603]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2415
		_go_fuzz_dep_.CoverTab[3440]++
												return SOAResource{}, &nestedError{"Expire", err}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2416
		// _ = "end of CoverTab[3440]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2417
		_go_fuzz_dep_.CoverTab[526604]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2417
		_go_fuzz_dep_.CoverTab[3441]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2417
		// _ = "end of CoverTab[3441]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2417
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2417
	// _ = "end of CoverTab[3427]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2417
	_go_fuzz_dep_.CoverTab[3428]++
											minTTL, _, err := unpackUint32(msg, off)
											if err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2419
		_go_fuzz_dep_.CoverTab[526605]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2419
		_go_fuzz_dep_.CoverTab[3442]++
												return SOAResource{}, &nestedError{"MinTTL", err}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2420
		// _ = "end of CoverTab[3442]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2421
		_go_fuzz_dep_.CoverTab[526606]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2421
		_go_fuzz_dep_.CoverTab[3443]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2421
		// _ = "end of CoverTab[3443]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2421
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2421
	// _ = "end of CoverTab[3428]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2421
	_go_fuzz_dep_.CoverTab[3429]++
											return SOAResource{ns, mbox, serial, refresh, retry, expire, minTTL}, nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2422
	// _ = "end of CoverTab[3429]"
}

// A TXTResource is a TXT Resource record.
type TXTResource struct {
	TXT []string
}

func (r *TXTResource) realType() Type {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2430
	_go_fuzz_dep_.CoverTab[3444]++
											return TypeTXT
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2431
	// _ = "end of CoverTab[3444]"
}

// pack appends the wire format of the TXTResource to msg.
func (r *TXTResource) pack(msg []byte, compression map[string]int, compressionOff int) ([]byte, error) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2435
	_go_fuzz_dep_.CoverTab[3445]++
											oldMsg := msg
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2436
	_go_fuzz_dep_.CoverTab[786595] = 0
											for _, s := range r.TXT {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2437
		if _go_fuzz_dep_.CoverTab[786595] == 0 {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2437
			_go_fuzz_dep_.CoverTab[526723]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2437
		} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2437
			_go_fuzz_dep_.CoverTab[526724]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2437
		}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2437
		_go_fuzz_dep_.CoverTab[786595] = 1
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2437
		_go_fuzz_dep_.CoverTab[3447]++
												var err error
												msg, err = packText(msg, s)
												if err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2440
			_go_fuzz_dep_.CoverTab[526607]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2440
			_go_fuzz_dep_.CoverTab[3448]++
													return oldMsg, err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2441
			// _ = "end of CoverTab[3448]"
		} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2442
			_go_fuzz_dep_.CoverTab[526608]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2442
			_go_fuzz_dep_.CoverTab[3449]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2442
			// _ = "end of CoverTab[3449]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2442
		}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2442
		// _ = "end of CoverTab[3447]"
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2443
	if _go_fuzz_dep_.CoverTab[786595] == 0 {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2443
		_go_fuzz_dep_.CoverTab[526725]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2443
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2443
		_go_fuzz_dep_.CoverTab[526726]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2443
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2443
	// _ = "end of CoverTab[3445]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2443
	_go_fuzz_dep_.CoverTab[3446]++
											return msg, nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2444
	// _ = "end of CoverTab[3446]"
}

// GoString implements fmt.GoStringer.GoString.
func (r *TXTResource) GoString() string {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2448
	_go_fuzz_dep_.CoverTab[3450]++
											s := "dnsmessage.TXTResource{TXT: []string{"
											if len(r.TXT) == 0 {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2450
		_go_fuzz_dep_.CoverTab[526609]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2450
		_go_fuzz_dep_.CoverTab[3453]++
												return s + "}}"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2451
		// _ = "end of CoverTab[3453]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2452
		_go_fuzz_dep_.CoverTab[526610]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2452
		_go_fuzz_dep_.CoverTab[3454]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2452
		// _ = "end of CoverTab[3454]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2452
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2452
	// _ = "end of CoverTab[3450]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2452
	_go_fuzz_dep_.CoverTab[3451]++
											s += `"` + printString([]byte(r.TXT[0]))
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2453
	_go_fuzz_dep_.CoverTab[786596] = 0
											for _, t := range r.TXT[1:] {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2454
		if _go_fuzz_dep_.CoverTab[786596] == 0 {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2454
			_go_fuzz_dep_.CoverTab[526727]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2454
		} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2454
			_go_fuzz_dep_.CoverTab[526728]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2454
		}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2454
		_go_fuzz_dep_.CoverTab[786596] = 1
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2454
		_go_fuzz_dep_.CoverTab[3455]++
												s += `", "` + printString([]byte(t))
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2455
		// _ = "end of CoverTab[3455]"
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2456
	if _go_fuzz_dep_.CoverTab[786596] == 0 {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2456
		_go_fuzz_dep_.CoverTab[526729]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2456
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2456
		_go_fuzz_dep_.CoverTab[526730]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2456
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2456
	// _ = "end of CoverTab[3451]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2456
	_go_fuzz_dep_.CoverTab[3452]++
											return s + `"}}`
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2457
	// _ = "end of CoverTab[3452]"
}

func unpackTXTResource(msg []byte, off int, length uint16) (TXTResource, error) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2460
	_go_fuzz_dep_.CoverTab[3456]++
											txts := make([]string, 0, 1)
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2461
	_go_fuzz_dep_.CoverTab[786597] = 0
											for n := uint16(0); n < length; {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2462
		if _go_fuzz_dep_.CoverTab[786597] == 0 {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2462
			_go_fuzz_dep_.CoverTab[526731]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2462
		} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2462
			_go_fuzz_dep_.CoverTab[526732]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2462
		}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2462
		_go_fuzz_dep_.CoverTab[786597] = 1
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2462
		_go_fuzz_dep_.CoverTab[3458]++
												var t string
												var err error
												if t, off, err = unpackText(msg, off); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2465
			_go_fuzz_dep_.CoverTab[526611]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2465
			_go_fuzz_dep_.CoverTab[3461]++
													return TXTResource{}, &nestedError{"text", err}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2466
			// _ = "end of CoverTab[3461]"
		} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2467
			_go_fuzz_dep_.CoverTab[526612]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2467
			_go_fuzz_dep_.CoverTab[3462]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2467
			// _ = "end of CoverTab[3462]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2467
		}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2467
		// _ = "end of CoverTab[3458]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2467
		_go_fuzz_dep_.CoverTab[3459]++

												if length-n < uint16(len(t))+1 {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2469
			_go_fuzz_dep_.CoverTab[526613]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2469
			_go_fuzz_dep_.CoverTab[3463]++
													return TXTResource{}, errCalcLen
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2470
			// _ = "end of CoverTab[3463]"
		} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2471
			_go_fuzz_dep_.CoverTab[526614]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2471
			_go_fuzz_dep_.CoverTab[3464]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2471
			// _ = "end of CoverTab[3464]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2471
		}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2471
		// _ = "end of CoverTab[3459]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2471
		_go_fuzz_dep_.CoverTab[3460]++
												n += uint16(len(t)) + 1
												txts = append(txts, t)
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2473
		// _ = "end of CoverTab[3460]"
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2474
	if _go_fuzz_dep_.CoverTab[786597] == 0 {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2474
		_go_fuzz_dep_.CoverTab[526733]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2474
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2474
		_go_fuzz_dep_.CoverTab[526734]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2474
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2474
	// _ = "end of CoverTab[3456]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2474
	_go_fuzz_dep_.CoverTab[3457]++
											return TXTResource{txts}, nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2475
	// _ = "end of CoverTab[3457]"
}

// An SRVResource is an SRV Resource record.
type SRVResource struct {
	Priority	uint16
	Weight		uint16
	Port		uint16
	Target		Name	// Not compressed as per RFC 2782.
}

func (r *SRVResource) realType() Type {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2486
	_go_fuzz_dep_.CoverTab[3465]++
											return TypeSRV
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2487
	// _ = "end of CoverTab[3465]"
}

// pack appends the wire format of the SRVResource to msg.
func (r *SRVResource) pack(msg []byte, compression map[string]int, compressionOff int) ([]byte, error) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2491
	_go_fuzz_dep_.CoverTab[3466]++
											oldMsg := msg
											msg = packUint16(msg, r.Priority)
											msg = packUint16(msg, r.Weight)
											msg = packUint16(msg, r.Port)
											msg, err := r.Target.pack(msg, nil, compressionOff)
											if err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2497
		_go_fuzz_dep_.CoverTab[526615]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2497
		_go_fuzz_dep_.CoverTab[3468]++
												return oldMsg, &nestedError{"SRVResource.Target", err}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2498
		// _ = "end of CoverTab[3468]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2499
		_go_fuzz_dep_.CoverTab[526616]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2499
		_go_fuzz_dep_.CoverTab[3469]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2499
		// _ = "end of CoverTab[3469]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2499
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2499
	// _ = "end of CoverTab[3466]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2499
	_go_fuzz_dep_.CoverTab[3467]++
											return msg, nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2500
	// _ = "end of CoverTab[3467]"
}

// GoString implements fmt.GoStringer.GoString.
func (r *SRVResource) GoString() string {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2504
	_go_fuzz_dep_.CoverTab[3470]++
											return "dnsmessage.SRVResource{" +
		"Priority: " + printUint16(r.Priority) + ", " +
		"Weight: " + printUint16(r.Weight) + ", " +
		"Port: " + printUint16(r.Port) + ", " +
		"Target: " + r.Target.GoString() + "}"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2509
	// _ = "end of CoverTab[3470]"
}

func unpackSRVResource(msg []byte, off int) (SRVResource, error) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2512
	_go_fuzz_dep_.CoverTab[3471]++
											priority, off, err := unpackUint16(msg, off)
											if err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2514
		_go_fuzz_dep_.CoverTab[526617]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2514
		_go_fuzz_dep_.CoverTab[3476]++
												return SRVResource{}, &nestedError{"Priority", err}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2515
		// _ = "end of CoverTab[3476]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2516
		_go_fuzz_dep_.CoverTab[526618]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2516
		_go_fuzz_dep_.CoverTab[3477]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2516
		// _ = "end of CoverTab[3477]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2516
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2516
	// _ = "end of CoverTab[3471]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2516
	_go_fuzz_dep_.CoverTab[3472]++
											weight, off, err := unpackUint16(msg, off)
											if err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2518
		_go_fuzz_dep_.CoverTab[526619]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2518
		_go_fuzz_dep_.CoverTab[3478]++
												return SRVResource{}, &nestedError{"Weight", err}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2519
		// _ = "end of CoverTab[3478]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2520
		_go_fuzz_dep_.CoverTab[526620]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2520
		_go_fuzz_dep_.CoverTab[3479]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2520
		// _ = "end of CoverTab[3479]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2520
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2520
	// _ = "end of CoverTab[3472]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2520
	_go_fuzz_dep_.CoverTab[3473]++
											port, off, err := unpackUint16(msg, off)
											if err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2522
		_go_fuzz_dep_.CoverTab[526621]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2522
		_go_fuzz_dep_.CoverTab[3480]++
												return SRVResource{}, &nestedError{"Port", err}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2523
		// _ = "end of CoverTab[3480]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2524
		_go_fuzz_dep_.CoverTab[526622]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2524
		_go_fuzz_dep_.CoverTab[3481]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2524
		// _ = "end of CoverTab[3481]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2524
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2524
	// _ = "end of CoverTab[3473]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2524
	_go_fuzz_dep_.CoverTab[3474]++
											var target Name
											if _, err := target.unpackCompressed(msg, off, false); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2526
		_go_fuzz_dep_.CoverTab[526623]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2526
		_go_fuzz_dep_.CoverTab[3482]++
												return SRVResource{}, &nestedError{"Target", err}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2527
		// _ = "end of CoverTab[3482]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2528
		_go_fuzz_dep_.CoverTab[526624]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2528
		_go_fuzz_dep_.CoverTab[3483]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2528
		// _ = "end of CoverTab[3483]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2528
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2528
	// _ = "end of CoverTab[3474]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2528
	_go_fuzz_dep_.CoverTab[3475]++
											return SRVResource{priority, weight, port, target}, nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2529
	// _ = "end of CoverTab[3475]"
}

// An AResource is an A Resource record.
type AResource struct {
	A [4]byte
}

func (r *AResource) realType() Type {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2537
	_go_fuzz_dep_.CoverTab[3484]++
											return TypeA
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2538
	// _ = "end of CoverTab[3484]"
}

// pack appends the wire format of the AResource to msg.
func (r *AResource) pack(msg []byte, compression map[string]int, compressionOff int) ([]byte, error) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2542
	_go_fuzz_dep_.CoverTab[3485]++
											return packBytes(msg, r.A[:]), nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2543
	// _ = "end of CoverTab[3485]"
}

// GoString implements fmt.GoStringer.GoString.
func (r *AResource) GoString() string {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2547
	_go_fuzz_dep_.CoverTab[3486]++
											return "dnsmessage.AResource{" +
		"A: [4]byte{" + printByteSlice(r.A[:]) + "}}"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2549
	// _ = "end of CoverTab[3486]"
}

func unpackAResource(msg []byte, off int) (AResource, error) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2552
	_go_fuzz_dep_.CoverTab[3487]++
											var a [4]byte
											if _, err := unpackBytes(msg, off, a[:]); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2554
		_go_fuzz_dep_.CoverTab[526625]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2554
		_go_fuzz_dep_.CoverTab[3489]++
												return AResource{}, err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2555
		// _ = "end of CoverTab[3489]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2556
		_go_fuzz_dep_.CoverTab[526626]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2556
		_go_fuzz_dep_.CoverTab[3490]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2556
		// _ = "end of CoverTab[3490]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2556
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2556
	// _ = "end of CoverTab[3487]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2556
	_go_fuzz_dep_.CoverTab[3488]++
											return AResource{a}, nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2557
	// _ = "end of CoverTab[3488]"
}

// An AAAAResource is an AAAA Resource record.
type AAAAResource struct {
	AAAA [16]byte
}

func (r *AAAAResource) realType() Type {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2565
	_go_fuzz_dep_.CoverTab[3491]++
											return TypeAAAA
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2566
	// _ = "end of CoverTab[3491]"
}

// GoString implements fmt.GoStringer.GoString.
func (r *AAAAResource) GoString() string {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2570
	_go_fuzz_dep_.CoverTab[3492]++
											return "dnsmessage.AAAAResource{" +
		"AAAA: [16]byte{" + printByteSlice(r.AAAA[:]) + "}}"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2572
	// _ = "end of CoverTab[3492]"
}

// pack appends the wire format of the AAAAResource to msg.
func (r *AAAAResource) pack(msg []byte, compression map[string]int, compressionOff int) ([]byte, error) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2576
	_go_fuzz_dep_.CoverTab[3493]++
											return packBytes(msg, r.AAAA[:]), nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2577
	// _ = "end of CoverTab[3493]"
}

func unpackAAAAResource(msg []byte, off int) (AAAAResource, error) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2580
	_go_fuzz_dep_.CoverTab[3494]++
											var aaaa [16]byte
											if _, err := unpackBytes(msg, off, aaaa[:]); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2582
		_go_fuzz_dep_.CoverTab[526627]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2582
		_go_fuzz_dep_.CoverTab[3496]++
												return AAAAResource{}, err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2583
		// _ = "end of CoverTab[3496]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2584
		_go_fuzz_dep_.CoverTab[526628]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2584
		_go_fuzz_dep_.CoverTab[3497]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2584
		// _ = "end of CoverTab[3497]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2584
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2584
	// _ = "end of CoverTab[3494]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2584
	_go_fuzz_dep_.CoverTab[3495]++
											return AAAAResource{aaaa}, nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2585
	// _ = "end of CoverTab[3495]"
}

// An OPTResource is an OPT pseudo Resource record.
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2588
//
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2588
// The pseudo resource record is part of the extension mechanisms for DNS
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2588
// as defined in RFC 6891.
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2592
type OPTResource struct {
	Options []Option
}

// An Option represents a DNS message option within OPTResource.
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2596
//
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2596
// The message option is part of the extension mechanisms for DNS as
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2596
// defined in RFC 6891.
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2600
type Option struct {
	Code	uint16	// option code
	Data	[]byte
}

// GoString implements fmt.GoStringer.GoString.
func (o *Option) GoString() string {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2606
	_go_fuzz_dep_.CoverTab[3498]++
											return "dnsmessage.Option{" +
		"Code: " + printUint16(o.Code) + ", " +
		"Data: []byte{" + printByteSlice(o.Data) + "}}"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2609
	// _ = "end of CoverTab[3498]"
}

func (r *OPTResource) realType() Type {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2612
	_go_fuzz_dep_.CoverTab[3499]++
											return TypeOPT
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2613
	// _ = "end of CoverTab[3499]"
}

func (r *OPTResource) pack(msg []byte, compression map[string]int, compressionOff int) ([]byte, error) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2616
	_go_fuzz_dep_.CoverTab[3500]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2616
	_go_fuzz_dep_.CoverTab[786598] = 0
											for _, opt := range r.Options {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2617
		if _go_fuzz_dep_.CoverTab[786598] == 0 {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2617
			_go_fuzz_dep_.CoverTab[526735]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2617
		} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2617
			_go_fuzz_dep_.CoverTab[526736]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2617
		}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2617
		_go_fuzz_dep_.CoverTab[786598] = 1
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2617
		_go_fuzz_dep_.CoverTab[3502]++
												msg = packUint16(msg, opt.Code)
												l := uint16(len(opt.Data))
												msg = packUint16(msg, l)
												msg = packBytes(msg, opt.Data)
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2621
		// _ = "end of CoverTab[3502]"
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2622
	if _go_fuzz_dep_.CoverTab[786598] == 0 {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2622
		_go_fuzz_dep_.CoverTab[526737]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2622
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2622
		_go_fuzz_dep_.CoverTab[526738]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2622
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2622
	// _ = "end of CoverTab[3500]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2622
	_go_fuzz_dep_.CoverTab[3501]++
											return msg, nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2623
	// _ = "end of CoverTab[3501]"
}

// GoString implements fmt.GoStringer.GoString.
func (r *OPTResource) GoString() string {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2627
	_go_fuzz_dep_.CoverTab[3503]++
											s := "dnsmessage.OPTResource{Options: []dnsmessage.Option{"
											if len(r.Options) == 0 {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2629
		_go_fuzz_dep_.CoverTab[526629]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2629
		_go_fuzz_dep_.CoverTab[3506]++
												return s + "}}"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2630
		// _ = "end of CoverTab[3506]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2631
		_go_fuzz_dep_.CoverTab[526630]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2631
		_go_fuzz_dep_.CoverTab[3507]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2631
		// _ = "end of CoverTab[3507]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2631
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2631
	// _ = "end of CoverTab[3503]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2631
	_go_fuzz_dep_.CoverTab[3504]++
											s += r.Options[0].GoString()
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2632
	_go_fuzz_dep_.CoverTab[786599] = 0
											for _, o := range r.Options[1:] {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2633
		if _go_fuzz_dep_.CoverTab[786599] == 0 {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2633
			_go_fuzz_dep_.CoverTab[526739]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2633
		} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2633
			_go_fuzz_dep_.CoverTab[526740]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2633
		}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2633
		_go_fuzz_dep_.CoverTab[786599] = 1
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2633
		_go_fuzz_dep_.CoverTab[3508]++
												s += ", " + o.GoString()
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2634
		// _ = "end of CoverTab[3508]"
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2635
	if _go_fuzz_dep_.CoverTab[786599] == 0 {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2635
		_go_fuzz_dep_.CoverTab[526741]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2635
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2635
		_go_fuzz_dep_.CoverTab[526742]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2635
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2635
	// _ = "end of CoverTab[3504]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2635
	_go_fuzz_dep_.CoverTab[3505]++
											return s + "}}"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2636
	// _ = "end of CoverTab[3505]"
}

func unpackOPTResource(msg []byte, off int, length uint16) (OPTResource, error) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2639
	_go_fuzz_dep_.CoverTab[3509]++
											var opts []Option
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2640
	_go_fuzz_dep_.CoverTab[786600] = 0
											for oldOff := off; off < oldOff+int(length); {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2641
		if _go_fuzz_dep_.CoverTab[786600] == 0 {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2641
			_go_fuzz_dep_.CoverTab[526743]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2641
		} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2641
			_go_fuzz_dep_.CoverTab[526744]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2641
		}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2641
		_go_fuzz_dep_.CoverTab[786600] = 1
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2641
		_go_fuzz_dep_.CoverTab[3511]++
												var err error
												var o Option
												o.Code, off, err = unpackUint16(msg, off)
												if err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2645
			_go_fuzz_dep_.CoverTab[526631]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2645
			_go_fuzz_dep_.CoverTab[3515]++
													return OPTResource{}, &nestedError{"Code", err}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2646
			// _ = "end of CoverTab[3515]"
		} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2647
			_go_fuzz_dep_.CoverTab[526632]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2647
			_go_fuzz_dep_.CoverTab[3516]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2647
			// _ = "end of CoverTab[3516]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2647
		}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2647
		// _ = "end of CoverTab[3511]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2647
		_go_fuzz_dep_.CoverTab[3512]++
												var l uint16
												l, off, err = unpackUint16(msg, off)
												if err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2650
			_go_fuzz_dep_.CoverTab[526633]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2650
			_go_fuzz_dep_.CoverTab[3517]++
													return OPTResource{}, &nestedError{"Data", err}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2651
			// _ = "end of CoverTab[3517]"
		} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2652
			_go_fuzz_dep_.CoverTab[526634]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2652
			_go_fuzz_dep_.CoverTab[3518]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2652
			// _ = "end of CoverTab[3518]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2652
		}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2652
		// _ = "end of CoverTab[3512]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2652
		_go_fuzz_dep_.CoverTab[3513]++
												o.Data = make([]byte, l)
												if copy(o.Data, msg[off:]) != int(l) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2654
			_go_fuzz_dep_.CoverTab[526635]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2654
			_go_fuzz_dep_.CoverTab[3519]++
													return OPTResource{}, &nestedError{"Data", errCalcLen}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2655
			// _ = "end of CoverTab[3519]"
		} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2656
			_go_fuzz_dep_.CoverTab[526636]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2656
			_go_fuzz_dep_.CoverTab[3520]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2656
			// _ = "end of CoverTab[3520]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2656
		}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2656
		// _ = "end of CoverTab[3513]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2656
		_go_fuzz_dep_.CoverTab[3514]++
												off += int(l)
												opts = append(opts, o)
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2658
		// _ = "end of CoverTab[3514]"
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2659
	if _go_fuzz_dep_.CoverTab[786600] == 0 {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2659
		_go_fuzz_dep_.CoverTab[526745]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2659
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2659
		_go_fuzz_dep_.CoverTab[526746]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2659
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2659
	// _ = "end of CoverTab[3509]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2659
	_go_fuzz_dep_.CoverTab[3510]++
											return OPTResource{opts}, nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2660
	// _ = "end of CoverTab[3510]"
}

// An UnknownResource is a catch-all container for unknown record types.
type UnknownResource struct {
	Type	Type
	Data	[]byte
}

func (r *UnknownResource) realType() Type {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2669
	_go_fuzz_dep_.CoverTab[3521]++
											return r.Type
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2670
	// _ = "end of CoverTab[3521]"
}

// pack appends the wire format of the UnknownResource to msg.
func (r *UnknownResource) pack(msg []byte, compression map[string]int, compressionOff int) ([]byte, error) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2674
	_go_fuzz_dep_.CoverTab[3522]++
											return packBytes(msg, r.Data[:]), nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2675
	// _ = "end of CoverTab[3522]"
}

// GoString implements fmt.GoStringer.GoString.
func (r *UnknownResource) GoString() string {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2679
	_go_fuzz_dep_.CoverTab[3523]++
											return "dnsmessage.UnknownResource{" +
		"Type: " + r.Type.GoString() + ", " +
		"Data: []byte{" + printByteSlice(r.Data) + "}}"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2682
	// _ = "end of CoverTab[3523]"
}

func unpackUnknownResource(recordType Type, msg []byte, off int, length uint16) (UnknownResource, error) {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2685
	_go_fuzz_dep_.CoverTab[3524]++
											parsed := UnknownResource{
		Type:	recordType,
		Data:	make([]byte, length),
	}
	if _, err := unpackBytes(msg, off, parsed.Data); err != nil {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2690
		_go_fuzz_dep_.CoverTab[526637]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2690
		_go_fuzz_dep_.CoverTab[3526]++
												return UnknownResource{}, err
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2691
		// _ = "end of CoverTab[3526]"
	} else {
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2692
		_go_fuzz_dep_.CoverTab[526638]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2692
		_go_fuzz_dep_.CoverTab[3527]++
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2692
		// _ = "end of CoverTab[3527]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2692
	}
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2692
	// _ = "end of CoverTab[3524]"
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2692
	_go_fuzz_dep_.CoverTab[3525]++
											return parsed, nil
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2693
	// _ = "end of CoverTab[3525]"
}

//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2694
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /snap/go/10455/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2694
var _ = _go_fuzz_dep_.CoverTab
