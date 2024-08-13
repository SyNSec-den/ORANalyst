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
	_go_fuzz_dep_.CoverTab[2437]++
										if n, ok := typeNames[t]; ok {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:65
		_go_fuzz_dep_.CoverTab[2439]++
											return n
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:66
		// _ = "end of CoverTab[2439]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:67
		_go_fuzz_dep_.CoverTab[2440]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:67
		// _ = "end of CoverTab[2440]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:67
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:67
	// _ = "end of CoverTab[2437]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:67
	_go_fuzz_dep_.CoverTab[2438]++
										return printUint16(uint16(t))
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:68
	// _ = "end of CoverTab[2438]"
}

// GoString implements fmt.GoStringer.GoString.
func (t Type) GoString() string {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:72
	_go_fuzz_dep_.CoverTab[2441]++
										if n, ok := typeNames[t]; ok {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:73
		_go_fuzz_dep_.CoverTab[2443]++
											return "dnsmessage." + n
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:74
		// _ = "end of CoverTab[2443]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:75
		_go_fuzz_dep_.CoverTab[2444]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:75
		// _ = "end of CoverTab[2444]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:75
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:75
	// _ = "end of CoverTab[2441]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:75
	_go_fuzz_dep_.CoverTab[2442]++
										return printUint16(uint16(t))
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:76
	// _ = "end of CoverTab[2442]"
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
	_go_fuzz_dep_.CoverTab[2445]++
										if n, ok := classNames[c]; ok {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:103
		_go_fuzz_dep_.CoverTab[2447]++
											return n
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:104
		// _ = "end of CoverTab[2447]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:105
		_go_fuzz_dep_.CoverTab[2448]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:105
		// _ = "end of CoverTab[2448]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:105
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:105
	// _ = "end of CoverTab[2445]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:105
	_go_fuzz_dep_.CoverTab[2446]++
										return printUint16(uint16(c))
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:106
	// _ = "end of CoverTab[2446]"
}

// GoString implements fmt.GoStringer.GoString.
func (c Class) GoString() string {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:110
	_go_fuzz_dep_.CoverTab[2449]++
										if n, ok := classNames[c]; ok {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:111
		_go_fuzz_dep_.CoverTab[2451]++
											return "dnsmessage." + n
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:112
		// _ = "end of CoverTab[2451]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:113
		_go_fuzz_dep_.CoverTab[2452]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:113
		// _ = "end of CoverTab[2452]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:113
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:113
	// _ = "end of CoverTab[2449]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:113
	_go_fuzz_dep_.CoverTab[2450]++
										return printUint16(uint16(c))
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:114
	// _ = "end of CoverTab[2450]"
}

// An OpCode is a DNS operation code.
type OpCode uint16

// GoString implements fmt.GoStringer.GoString.
func (o OpCode) GoString() string {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:121
	_go_fuzz_dep_.CoverTab[2453]++
										return printUint16(uint16(o))
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:122
	// _ = "end of CoverTab[2453]"
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
	_go_fuzz_dep_.CoverTab[2454]++
										if n, ok := rCodeNames[r]; ok {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:149
		_go_fuzz_dep_.CoverTab[2456]++
											return n
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:150
		// _ = "end of CoverTab[2456]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:151
		_go_fuzz_dep_.CoverTab[2457]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:151
		// _ = "end of CoverTab[2457]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:151
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:151
	// _ = "end of CoverTab[2454]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:151
	_go_fuzz_dep_.CoverTab[2455]++
										return printUint16(uint16(r))
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:152
	// _ = "end of CoverTab[2455]"
}

// GoString implements fmt.GoStringer.GoString.
func (r RCode) GoString() string {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:156
	_go_fuzz_dep_.CoverTab[2458]++
										if n, ok := rCodeNames[r]; ok {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:157
		_go_fuzz_dep_.CoverTab[2460]++
											return "dnsmessage." + n
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:158
		// _ = "end of CoverTab[2460]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:159
		_go_fuzz_dep_.CoverTab[2461]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:159
		// _ = "end of CoverTab[2461]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:159
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:159
	// _ = "end of CoverTab[2458]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:159
	_go_fuzz_dep_.CoverTab[2459]++
										return printUint16(uint16(r))
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:160
	// _ = "end of CoverTab[2459]"
}

func printPaddedUint8(i uint8) string {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:163
	_go_fuzz_dep_.CoverTab[2462]++
										b := byte(i)
										return string([]byte{
		b/100 + '0',
		b/10%10 + '0',
		b%10 + '0',
	})
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:169
	// _ = "end of CoverTab[2462]"
}

func printUint8Bytes(buf []byte, i uint8) []byte {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:172
	_go_fuzz_dep_.CoverTab[2463]++
										b := byte(i)
										if i >= 100 {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:174
		_go_fuzz_dep_.CoverTab[2466]++
											buf = append(buf, b/100+'0')
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:175
		// _ = "end of CoverTab[2466]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:176
		_go_fuzz_dep_.CoverTab[2467]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:176
		// _ = "end of CoverTab[2467]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:176
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:176
	// _ = "end of CoverTab[2463]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:176
	_go_fuzz_dep_.CoverTab[2464]++
										if i >= 10 {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:177
		_go_fuzz_dep_.CoverTab[2468]++
											buf = append(buf, b/10%10+'0')
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:178
		// _ = "end of CoverTab[2468]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:179
		_go_fuzz_dep_.CoverTab[2469]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:179
		// _ = "end of CoverTab[2469]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:179
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:179
	// _ = "end of CoverTab[2464]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:179
	_go_fuzz_dep_.CoverTab[2465]++
										return append(buf, b%10+'0')
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:180
	// _ = "end of CoverTab[2465]"
}

func printByteSlice(b []byte) string {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:183
	_go_fuzz_dep_.CoverTab[2470]++
										if len(b) == 0 {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:184
		_go_fuzz_dep_.CoverTab[2473]++
											return ""
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:185
		// _ = "end of CoverTab[2473]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:186
		_go_fuzz_dep_.CoverTab[2474]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:186
		// _ = "end of CoverTab[2474]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:186
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:186
	// _ = "end of CoverTab[2470]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:186
	_go_fuzz_dep_.CoverTab[2471]++
										buf := make([]byte, 0, 5*len(b))
										buf = printUint8Bytes(buf, uint8(b[0]))
										for _, n := range b[1:] {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:189
		_go_fuzz_dep_.CoverTab[2475]++
											buf = append(buf, ',', ' ')
											buf = printUint8Bytes(buf, uint8(n))
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:191
		// _ = "end of CoverTab[2475]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:192
	// _ = "end of CoverTab[2471]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:192
	_go_fuzz_dep_.CoverTab[2472]++
										return string(buf)
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:193
	// _ = "end of CoverTab[2472]"
}

const hexDigits = "0123456789abcdef"

func printString(str []byte) string {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:198
	_go_fuzz_dep_.CoverTab[2476]++
										buf := make([]byte, 0, len(str))
										for i := 0; i < len(str); i++ {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:200
		_go_fuzz_dep_.CoverTab[2478]++
											c := str[i]
											if c == '.' || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:202
			_go_fuzz_dep_.CoverTab[2480]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:202
			return c == '-'
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:202
			// _ = "end of CoverTab[2480]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:202
		}() || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:202
			_go_fuzz_dep_.CoverTab[2481]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:202
			return c == ' '
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:202
			// _ = "end of CoverTab[2481]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:202
		}() || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:202
			_go_fuzz_dep_.CoverTab[2482]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:202
			return 'A' <= c && func() bool {
													_go_fuzz_dep_.CoverTab[2483]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:203
				return c <= 'Z'
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:203
				// _ = "end of CoverTab[2483]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:203
			}()
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:203
			// _ = "end of CoverTab[2482]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:203
		}() || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:203
			_go_fuzz_dep_.CoverTab[2484]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:203
			return 'a' <= c && func() bool {
													_go_fuzz_dep_.CoverTab[2485]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:204
				return c <= 'z'
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:204
				// _ = "end of CoverTab[2485]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:204
			}()
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:204
			// _ = "end of CoverTab[2484]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:204
		}() || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:204
			_go_fuzz_dep_.CoverTab[2486]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:204
			return '0' <= c && func() bool {
													_go_fuzz_dep_.CoverTab[2487]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:205
				return c <= '9'
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:205
				// _ = "end of CoverTab[2487]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:205
			}()
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:205
			// _ = "end of CoverTab[2486]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:205
		}() {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:205
			_go_fuzz_dep_.CoverTab[2488]++
												buf = append(buf, c)
												continue
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:207
			// _ = "end of CoverTab[2488]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:208
			_go_fuzz_dep_.CoverTab[2489]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:208
			// _ = "end of CoverTab[2489]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:208
		}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:208
		// _ = "end of CoverTab[2478]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:208
		_go_fuzz_dep_.CoverTab[2479]++

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
		// _ = "end of CoverTab[2479]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:219
	// _ = "end of CoverTab[2476]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:219
	_go_fuzz_dep_.CoverTab[2477]++
										return string(buf)
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:220
	// _ = "end of CoverTab[2477]"
}

func printUint16(i uint16) string {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:223
	_go_fuzz_dep_.CoverTab[2490]++
										return printUint32(uint32(i))
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:224
	// _ = "end of CoverTab[2490]"
}

func printUint32(i uint32) string {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:227
	_go_fuzz_dep_.CoverTab[2491]++

										buf := make([]byte, 10)
										for b, d := buf, uint32(1000000000); d > 0; d /= 10 {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:230
		_go_fuzz_dep_.CoverTab[2493]++
											b[0] = byte(i/d%10 + '0')
											if b[0] == '0' && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:232
			_go_fuzz_dep_.CoverTab[2495]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:232
			return len(b) == len(buf)
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:232
			// _ = "end of CoverTab[2495]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:232
		}() && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:232
			_go_fuzz_dep_.CoverTab[2496]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:232
			return len(buf) > 1
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:232
			// _ = "end of CoverTab[2496]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:232
		}() {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:232
			_go_fuzz_dep_.CoverTab[2497]++
												buf = buf[1:]
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:233
			// _ = "end of CoverTab[2497]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:234
			_go_fuzz_dep_.CoverTab[2498]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:234
			// _ = "end of CoverTab[2498]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:234
		}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:234
		// _ = "end of CoverTab[2493]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:234
		_go_fuzz_dep_.CoverTab[2494]++
											b = b[1:]
											i %= d
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:236
		// _ = "end of CoverTab[2494]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:237
	// _ = "end of CoverTab[2491]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:237
	_go_fuzz_dep_.CoverTab[2492]++
										return string(buf)
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:238
	// _ = "end of CoverTab[2492]"
}

func printBool(b bool) string {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:241
	_go_fuzz_dep_.CoverTab[2499]++
										if b {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:242
		_go_fuzz_dep_.CoverTab[2501]++
											return "true"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:243
		// _ = "end of CoverTab[2501]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:244
		_go_fuzz_dep_.CoverTab[2502]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:244
		// _ = "end of CoverTab[2502]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:244
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:244
	// _ = "end of CoverTab[2499]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:244
	_go_fuzz_dep_.CoverTab[2500]++
										return "false"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:245
	// _ = "end of CoverTab[2500]"
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
	_go_fuzz_dep_.CoverTab[2503]++
										return e.s + ": " + e.err.Error()
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:308
	// _ = "end of CoverTab[2503]"
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
	_go_fuzz_dep_.CoverTab[2504]++
										id = m.ID
										bits = uint16(m.OpCode)<<11 | uint16(m.RCode)
										if m.RecursionAvailable {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:328
		_go_fuzz_dep_.CoverTab[2512]++
											bits |= headerBitRA
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:329
		// _ = "end of CoverTab[2512]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:330
		_go_fuzz_dep_.CoverTab[2513]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:330
		// _ = "end of CoverTab[2513]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:330
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:330
	// _ = "end of CoverTab[2504]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:330
	_go_fuzz_dep_.CoverTab[2505]++
										if m.RecursionDesired {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:331
		_go_fuzz_dep_.CoverTab[2514]++
											bits |= headerBitRD
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:332
		// _ = "end of CoverTab[2514]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:333
		_go_fuzz_dep_.CoverTab[2515]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:333
		// _ = "end of CoverTab[2515]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:333
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:333
	// _ = "end of CoverTab[2505]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:333
	_go_fuzz_dep_.CoverTab[2506]++
										if m.Truncated {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:334
		_go_fuzz_dep_.CoverTab[2516]++
											bits |= headerBitTC
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:335
		// _ = "end of CoverTab[2516]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:336
		_go_fuzz_dep_.CoverTab[2517]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:336
		// _ = "end of CoverTab[2517]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:336
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:336
	// _ = "end of CoverTab[2506]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:336
	_go_fuzz_dep_.CoverTab[2507]++
										if m.Authoritative {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:337
		_go_fuzz_dep_.CoverTab[2518]++
											bits |= headerBitAA
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:338
		// _ = "end of CoverTab[2518]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:339
		_go_fuzz_dep_.CoverTab[2519]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:339
		// _ = "end of CoverTab[2519]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:339
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:339
	// _ = "end of CoverTab[2507]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:339
	_go_fuzz_dep_.CoverTab[2508]++
										if m.Response {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:340
		_go_fuzz_dep_.CoverTab[2520]++
											bits |= headerBitQR
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:341
		// _ = "end of CoverTab[2520]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:342
		_go_fuzz_dep_.CoverTab[2521]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:342
		// _ = "end of CoverTab[2521]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:342
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:342
	// _ = "end of CoverTab[2508]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:342
	_go_fuzz_dep_.CoverTab[2509]++
										if m.AuthenticData {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:343
		_go_fuzz_dep_.CoverTab[2522]++
											bits |= headerBitAD
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:344
		// _ = "end of CoverTab[2522]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:345
		_go_fuzz_dep_.CoverTab[2523]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:345
		// _ = "end of CoverTab[2523]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:345
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:345
	// _ = "end of CoverTab[2509]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:345
	_go_fuzz_dep_.CoverTab[2510]++
										if m.CheckingDisabled {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:346
		_go_fuzz_dep_.CoverTab[2524]++
											bits |= headerBitCD
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:347
		// _ = "end of CoverTab[2524]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:348
		_go_fuzz_dep_.CoverTab[2525]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:348
		// _ = "end of CoverTab[2525]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:348
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:348
	// _ = "end of CoverTab[2510]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:348
	_go_fuzz_dep_.CoverTab[2511]++
										return
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:349
	// _ = "end of CoverTab[2511]"
}

// GoString implements fmt.GoStringer.GoString.
func (m *Header) GoString() string {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:353
	_go_fuzz_dep_.CoverTab[2526]++
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
	// _ = "end of CoverTab[2526]"
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
	_go_fuzz_dep_.CoverTab[2527]++
										switch sec {
	case sectionQuestions:
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:414
		_go_fuzz_dep_.CoverTab[2529]++
											return h.questions
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:415
		// _ = "end of CoverTab[2529]"
	case sectionAnswers:
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:416
		_go_fuzz_dep_.CoverTab[2530]++
											return h.answers
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:417
		// _ = "end of CoverTab[2530]"
	case sectionAuthorities:
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:418
		_go_fuzz_dep_.CoverTab[2531]++
											return h.authorities
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:419
		// _ = "end of CoverTab[2531]"
	case sectionAdditionals:
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:420
		_go_fuzz_dep_.CoverTab[2532]++
											return h.additionals
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:421
		// _ = "end of CoverTab[2532]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:421
	default:
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:421
		_go_fuzz_dep_.CoverTab[2533]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:421
		// _ = "end of CoverTab[2533]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:422
	// _ = "end of CoverTab[2527]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:422
	_go_fuzz_dep_.CoverTab[2528]++
										return 0
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:423
	// _ = "end of CoverTab[2528]"
}

// pack appends the wire format of the header to msg.
func (h *header) pack(msg []byte) []byte {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:427
	_go_fuzz_dep_.CoverTab[2534]++
										msg = packUint16(msg, h.id)
										msg = packUint16(msg, h.bits)
										msg = packUint16(msg, h.questions)
										msg = packUint16(msg, h.answers)
										msg = packUint16(msg, h.authorities)
										return packUint16(msg, h.additionals)
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:433
	// _ = "end of CoverTab[2534]"
}

func (h *header) unpack(msg []byte, off int) (int, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:436
	_go_fuzz_dep_.CoverTab[2535]++
										newOff := off
										var err error
										if h.id, newOff, err = unpackUint16(msg, newOff); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:439
		_go_fuzz_dep_.CoverTab[2542]++
											return off, &nestedError{"id", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:440
		// _ = "end of CoverTab[2542]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:441
		_go_fuzz_dep_.CoverTab[2543]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:441
		// _ = "end of CoverTab[2543]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:441
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:441
	// _ = "end of CoverTab[2535]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:441
	_go_fuzz_dep_.CoverTab[2536]++
										if h.bits, newOff, err = unpackUint16(msg, newOff); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:442
		_go_fuzz_dep_.CoverTab[2544]++
											return off, &nestedError{"bits", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:443
		// _ = "end of CoverTab[2544]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:444
		_go_fuzz_dep_.CoverTab[2545]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:444
		// _ = "end of CoverTab[2545]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:444
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:444
	// _ = "end of CoverTab[2536]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:444
	_go_fuzz_dep_.CoverTab[2537]++
										if h.questions, newOff, err = unpackUint16(msg, newOff); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:445
		_go_fuzz_dep_.CoverTab[2546]++
											return off, &nestedError{"questions", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:446
		// _ = "end of CoverTab[2546]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:447
		_go_fuzz_dep_.CoverTab[2547]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:447
		// _ = "end of CoverTab[2547]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:447
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:447
	// _ = "end of CoverTab[2537]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:447
	_go_fuzz_dep_.CoverTab[2538]++
										if h.answers, newOff, err = unpackUint16(msg, newOff); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:448
		_go_fuzz_dep_.CoverTab[2548]++
											return off, &nestedError{"answers", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:449
		// _ = "end of CoverTab[2548]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:450
		_go_fuzz_dep_.CoverTab[2549]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:450
		// _ = "end of CoverTab[2549]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:450
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:450
	// _ = "end of CoverTab[2538]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:450
	_go_fuzz_dep_.CoverTab[2539]++
										if h.authorities, newOff, err = unpackUint16(msg, newOff); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:451
		_go_fuzz_dep_.CoverTab[2550]++
											return off, &nestedError{"authorities", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:452
		// _ = "end of CoverTab[2550]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:453
		_go_fuzz_dep_.CoverTab[2551]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:453
		// _ = "end of CoverTab[2551]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:453
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:453
	// _ = "end of CoverTab[2539]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:453
	_go_fuzz_dep_.CoverTab[2540]++
										if h.additionals, newOff, err = unpackUint16(msg, newOff); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:454
		_go_fuzz_dep_.CoverTab[2552]++
											return off, &nestedError{"additionals", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:455
		// _ = "end of CoverTab[2552]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:456
		_go_fuzz_dep_.CoverTab[2553]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:456
		// _ = "end of CoverTab[2553]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:456
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:456
	// _ = "end of CoverTab[2540]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:456
	_go_fuzz_dep_.CoverTab[2541]++
										return newOff, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:457
	// _ = "end of CoverTab[2541]"
}

func (h *header) header() Header {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:460
	_go_fuzz_dep_.CoverTab[2554]++
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
	// _ = "end of CoverTab[2554]"
}

// A Resource is a DNS resource record.
type Resource struct {
	Header	ResourceHeader
	Body	ResourceBody
}

func (r *Resource) GoString() string {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:481
	_go_fuzz_dep_.CoverTab[2555]++
										return "dnsmessage.Resource{" +
		"Header: " + r.Header.GoString() +
		", Body: &" + r.Body.GoString() +
		"}"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:485
	// _ = "end of CoverTab[2555]"
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
	_go_fuzz_dep_.CoverTab[2556]++
										if r.Body == nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:503
		_go_fuzz_dep_.CoverTab[2561]++
											return msg, errNilResouceBody
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:504
		// _ = "end of CoverTab[2561]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:505
		_go_fuzz_dep_.CoverTab[2562]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:505
		// _ = "end of CoverTab[2562]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:505
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:505
	// _ = "end of CoverTab[2556]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:505
	_go_fuzz_dep_.CoverTab[2557]++
										oldMsg := msg
										r.Header.Type = r.Body.realType()
										msg, lenOff, err := r.Header.pack(msg, compression, compressionOff)
										if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:509
		_go_fuzz_dep_.CoverTab[2563]++
											return msg, &nestedError{"ResourceHeader", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:510
		// _ = "end of CoverTab[2563]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:511
		_go_fuzz_dep_.CoverTab[2564]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:511
		// _ = "end of CoverTab[2564]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:511
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:511
	// _ = "end of CoverTab[2557]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:511
	_go_fuzz_dep_.CoverTab[2558]++
										preLen := len(msg)
										msg, err = r.Body.pack(msg, compression, compressionOff)
										if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:514
		_go_fuzz_dep_.CoverTab[2565]++
											return msg, &nestedError{"content", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:515
		// _ = "end of CoverTab[2565]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:516
		_go_fuzz_dep_.CoverTab[2566]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:516
		// _ = "end of CoverTab[2566]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:516
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:516
	// _ = "end of CoverTab[2558]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:516
	_go_fuzz_dep_.CoverTab[2559]++
										if err := r.Header.fixLen(msg, lenOff, preLen); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:517
		_go_fuzz_dep_.CoverTab[2567]++
											return oldMsg, err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:518
		// _ = "end of CoverTab[2567]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:519
		_go_fuzz_dep_.CoverTab[2568]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:519
		// _ = "end of CoverTab[2568]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:519
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:519
	// _ = "end of CoverTab[2559]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:519
	_go_fuzz_dep_.CoverTab[2560]++
										return msg, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:520
	// _ = "end of CoverTab[2560]"
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
	_go_fuzz_dep_.CoverTab[2569]++
										if p.msg != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:548
		_go_fuzz_dep_.CoverTab[2572]++
											*p = Parser{}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:549
		// _ = "end of CoverTab[2572]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:550
		_go_fuzz_dep_.CoverTab[2573]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:550
		// _ = "end of CoverTab[2573]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:550
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:550
	// _ = "end of CoverTab[2569]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:550
	_go_fuzz_dep_.CoverTab[2570]++
										p.msg = msg
										var err error
										if p.off, err = p.header.unpack(msg, 0); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:553
		_go_fuzz_dep_.CoverTab[2574]++
											return Header{}, &nestedError{"unpacking header", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:554
		// _ = "end of CoverTab[2574]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:555
		_go_fuzz_dep_.CoverTab[2575]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:555
		// _ = "end of CoverTab[2575]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:555
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:555
	// _ = "end of CoverTab[2570]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:555
	_go_fuzz_dep_.CoverTab[2571]++
										p.section = sectionQuestions
										return p.header.header(), nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:557
	// _ = "end of CoverTab[2571]"
}

func (p *Parser) checkAdvance(sec section) error {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:560
	_go_fuzz_dep_.CoverTab[2576]++
										if p.section < sec {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:561
		_go_fuzz_dep_.CoverTab[2580]++
											return ErrNotStarted
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:562
		// _ = "end of CoverTab[2580]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:563
		_go_fuzz_dep_.CoverTab[2581]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:563
		// _ = "end of CoverTab[2581]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:563
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:563
	// _ = "end of CoverTab[2576]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:563
	_go_fuzz_dep_.CoverTab[2577]++
										if p.section > sec {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:564
		_go_fuzz_dep_.CoverTab[2582]++
											return ErrSectionDone
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:565
		// _ = "end of CoverTab[2582]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:566
		_go_fuzz_dep_.CoverTab[2583]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:566
		// _ = "end of CoverTab[2583]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:566
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:566
	// _ = "end of CoverTab[2577]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:566
	_go_fuzz_dep_.CoverTab[2578]++
										p.resHeaderValid = false
										if p.index == int(p.header.count(sec)) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:568
		_go_fuzz_dep_.CoverTab[2584]++
											p.index = 0
											p.section++
											return ErrSectionDone
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:571
		// _ = "end of CoverTab[2584]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:572
		_go_fuzz_dep_.CoverTab[2585]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:572
		// _ = "end of CoverTab[2585]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:572
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:572
	// _ = "end of CoverTab[2578]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:572
	_go_fuzz_dep_.CoverTab[2579]++
										return nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:573
	// _ = "end of CoverTab[2579]"
}

func (p *Parser) resource(sec section) (Resource, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:576
	_go_fuzz_dep_.CoverTab[2586]++
										var r Resource
										var err error
										r.Header, err = p.resourceHeader(sec)
										if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:580
		_go_fuzz_dep_.CoverTab[2589]++
											return r, err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:581
		// _ = "end of CoverTab[2589]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:582
		_go_fuzz_dep_.CoverTab[2590]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:582
		// _ = "end of CoverTab[2590]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:582
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:582
	// _ = "end of CoverTab[2586]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:582
	_go_fuzz_dep_.CoverTab[2587]++
										p.resHeaderValid = false
										r.Body, p.off, err = unpackResourceBody(p.msg, p.off, r.Header)
										if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:585
		_go_fuzz_dep_.CoverTab[2591]++
											return Resource{}, &nestedError{"unpacking " + sectionNames[sec], err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:586
		// _ = "end of CoverTab[2591]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:587
		_go_fuzz_dep_.CoverTab[2592]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:587
		// _ = "end of CoverTab[2592]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:587
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:587
	// _ = "end of CoverTab[2587]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:587
	_go_fuzz_dep_.CoverTab[2588]++
										p.index++
										return r, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:589
	// _ = "end of CoverTab[2588]"
}

func (p *Parser) resourceHeader(sec section) (ResourceHeader, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:592
	_go_fuzz_dep_.CoverTab[2593]++
										if p.resHeaderValid {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:593
		_go_fuzz_dep_.CoverTab[2597]++
											return p.resHeader, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:594
		// _ = "end of CoverTab[2597]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:595
		_go_fuzz_dep_.CoverTab[2598]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:595
		// _ = "end of CoverTab[2598]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:595
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:595
	// _ = "end of CoverTab[2593]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:595
	_go_fuzz_dep_.CoverTab[2594]++
										if err := p.checkAdvance(sec); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:596
		_go_fuzz_dep_.CoverTab[2599]++
											return ResourceHeader{}, err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:597
		// _ = "end of CoverTab[2599]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:598
		_go_fuzz_dep_.CoverTab[2600]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:598
		// _ = "end of CoverTab[2600]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:598
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:598
	// _ = "end of CoverTab[2594]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:598
	_go_fuzz_dep_.CoverTab[2595]++
										var hdr ResourceHeader
										off, err := hdr.unpack(p.msg, p.off)
										if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:601
		_go_fuzz_dep_.CoverTab[2601]++
											return ResourceHeader{}, err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:602
		// _ = "end of CoverTab[2601]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:603
		_go_fuzz_dep_.CoverTab[2602]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:603
		// _ = "end of CoverTab[2602]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:603
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:603
	// _ = "end of CoverTab[2595]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:603
	_go_fuzz_dep_.CoverTab[2596]++
										p.resHeaderValid = true
										p.resHeader = hdr
										p.off = off
										return hdr, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:607
	// _ = "end of CoverTab[2596]"
}

func (p *Parser) skipResource(sec section) error {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:610
	_go_fuzz_dep_.CoverTab[2603]++
										if p.resHeaderValid {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:611
		_go_fuzz_dep_.CoverTab[2607]++
											newOff := p.off + int(p.resHeader.Length)
											if newOff > len(p.msg) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:613
			_go_fuzz_dep_.CoverTab[2609]++
												return errResourceLen
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:614
			// _ = "end of CoverTab[2609]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:615
			_go_fuzz_dep_.CoverTab[2610]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:615
			// _ = "end of CoverTab[2610]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:615
		}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:615
		// _ = "end of CoverTab[2607]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:615
		_go_fuzz_dep_.CoverTab[2608]++
											p.off = newOff
											p.resHeaderValid = false
											p.index++
											return nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:619
		// _ = "end of CoverTab[2608]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:620
		_go_fuzz_dep_.CoverTab[2611]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:620
		// _ = "end of CoverTab[2611]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:620
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:620
	// _ = "end of CoverTab[2603]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:620
	_go_fuzz_dep_.CoverTab[2604]++
										if err := p.checkAdvance(sec); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:621
		_go_fuzz_dep_.CoverTab[2612]++
											return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:622
		// _ = "end of CoverTab[2612]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:623
		_go_fuzz_dep_.CoverTab[2613]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:623
		// _ = "end of CoverTab[2613]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:623
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:623
	// _ = "end of CoverTab[2604]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:623
	_go_fuzz_dep_.CoverTab[2605]++
										var err error
										p.off, err = skipResource(p.msg, p.off)
										if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:626
		_go_fuzz_dep_.CoverTab[2614]++
											return &nestedError{"skipping: " + sectionNames[sec], err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:627
		// _ = "end of CoverTab[2614]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:628
		_go_fuzz_dep_.CoverTab[2615]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:628
		// _ = "end of CoverTab[2615]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:628
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:628
	// _ = "end of CoverTab[2605]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:628
	_go_fuzz_dep_.CoverTab[2606]++
										p.index++
										return nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:630
	// _ = "end of CoverTab[2606]"
}

// Question parses a single Question.
func (p *Parser) Question() (Question, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:634
	_go_fuzz_dep_.CoverTab[2616]++
										if err := p.checkAdvance(sectionQuestions); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:635
		_go_fuzz_dep_.CoverTab[2621]++
											return Question{}, err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:636
		// _ = "end of CoverTab[2621]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:637
		_go_fuzz_dep_.CoverTab[2622]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:637
		// _ = "end of CoverTab[2622]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:637
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:637
	// _ = "end of CoverTab[2616]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:637
	_go_fuzz_dep_.CoverTab[2617]++
										var name Name
										off, err := name.unpack(p.msg, p.off)
										if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:640
		_go_fuzz_dep_.CoverTab[2623]++
											return Question{}, &nestedError{"unpacking Question.Name", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:641
		// _ = "end of CoverTab[2623]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:642
		_go_fuzz_dep_.CoverTab[2624]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:642
		// _ = "end of CoverTab[2624]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:642
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:642
	// _ = "end of CoverTab[2617]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:642
	_go_fuzz_dep_.CoverTab[2618]++
										typ, off, err := unpackType(p.msg, off)
										if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:644
		_go_fuzz_dep_.CoverTab[2625]++
											return Question{}, &nestedError{"unpacking Question.Type", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:645
		// _ = "end of CoverTab[2625]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:646
		_go_fuzz_dep_.CoverTab[2626]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:646
		// _ = "end of CoverTab[2626]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:646
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:646
	// _ = "end of CoverTab[2618]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:646
	_go_fuzz_dep_.CoverTab[2619]++
										class, off, err := unpackClass(p.msg, off)
										if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:648
		_go_fuzz_dep_.CoverTab[2627]++
											return Question{}, &nestedError{"unpacking Question.Class", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:649
		// _ = "end of CoverTab[2627]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:650
		_go_fuzz_dep_.CoverTab[2628]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:650
		// _ = "end of CoverTab[2628]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:650
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:650
	// _ = "end of CoverTab[2619]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:650
	_go_fuzz_dep_.CoverTab[2620]++
										p.off = off
										p.index++
										return Question{name, typ, class}, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:653
	// _ = "end of CoverTab[2620]"
}

// AllQuestions parses all Questions.
func (p *Parser) AllQuestions() ([]Question, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:657
	_go_fuzz_dep_.CoverTab[2629]++

//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:664
	qs := []Question{}
	for {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:665
		_go_fuzz_dep_.CoverTab[2630]++
											q, err := p.Question()
											if err == ErrSectionDone {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:667
			_go_fuzz_dep_.CoverTab[2633]++
												return qs, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:668
			// _ = "end of CoverTab[2633]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:669
			_go_fuzz_dep_.CoverTab[2634]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:669
			// _ = "end of CoverTab[2634]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:669
		}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:669
		// _ = "end of CoverTab[2630]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:669
		_go_fuzz_dep_.CoverTab[2631]++
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:670
			_go_fuzz_dep_.CoverTab[2635]++
												return nil, err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:671
			// _ = "end of CoverTab[2635]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:672
			_go_fuzz_dep_.CoverTab[2636]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:672
			// _ = "end of CoverTab[2636]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:672
		}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:672
		// _ = "end of CoverTab[2631]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:672
		_go_fuzz_dep_.CoverTab[2632]++
											qs = append(qs, q)
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:673
		// _ = "end of CoverTab[2632]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:674
	// _ = "end of CoverTab[2629]"
}

// SkipQuestion skips a single Question.
func (p *Parser) SkipQuestion() error {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:678
	_go_fuzz_dep_.CoverTab[2637]++
										if err := p.checkAdvance(sectionQuestions); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:679
		_go_fuzz_dep_.CoverTab[2642]++
											return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:680
		// _ = "end of CoverTab[2642]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:681
		_go_fuzz_dep_.CoverTab[2643]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:681
		// _ = "end of CoverTab[2643]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:681
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:681
	// _ = "end of CoverTab[2637]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:681
	_go_fuzz_dep_.CoverTab[2638]++
										off, err := skipName(p.msg, p.off)
										if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:683
		_go_fuzz_dep_.CoverTab[2644]++
											return &nestedError{"skipping Question Name", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:684
		// _ = "end of CoverTab[2644]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:685
		_go_fuzz_dep_.CoverTab[2645]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:685
		// _ = "end of CoverTab[2645]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:685
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:685
	// _ = "end of CoverTab[2638]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:685
	_go_fuzz_dep_.CoverTab[2639]++
										if off, err = skipType(p.msg, off); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:686
		_go_fuzz_dep_.CoverTab[2646]++
											return &nestedError{"skipping Question Type", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:687
		// _ = "end of CoverTab[2646]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:688
		_go_fuzz_dep_.CoverTab[2647]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:688
		// _ = "end of CoverTab[2647]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:688
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:688
	// _ = "end of CoverTab[2639]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:688
	_go_fuzz_dep_.CoverTab[2640]++
										if off, err = skipClass(p.msg, off); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:689
		_go_fuzz_dep_.CoverTab[2648]++
											return &nestedError{"skipping Question Class", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:690
		// _ = "end of CoverTab[2648]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:691
		_go_fuzz_dep_.CoverTab[2649]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:691
		// _ = "end of CoverTab[2649]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:691
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:691
	// _ = "end of CoverTab[2640]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:691
	_go_fuzz_dep_.CoverTab[2641]++
										p.off = off
										p.index++
										return nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:694
	// _ = "end of CoverTab[2641]"
}

// SkipAllQuestions skips all Questions.
func (p *Parser) SkipAllQuestions() error {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:698
	_go_fuzz_dep_.CoverTab[2650]++
										for {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:699
		_go_fuzz_dep_.CoverTab[2651]++
											if err := p.SkipQuestion(); err == ErrSectionDone {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:700
			_go_fuzz_dep_.CoverTab[2652]++
												return nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:701
			// _ = "end of CoverTab[2652]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:702
			_go_fuzz_dep_.CoverTab[2653]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:702
			if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:702
				_go_fuzz_dep_.CoverTab[2654]++
													return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:703
				// _ = "end of CoverTab[2654]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:704
				_go_fuzz_dep_.CoverTab[2655]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:704
				// _ = "end of CoverTab[2655]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:704
			}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:704
			// _ = "end of CoverTab[2653]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:704
		}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:704
		// _ = "end of CoverTab[2651]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:705
	// _ = "end of CoverTab[2650]"
}

// AnswerHeader parses a single Answer ResourceHeader.
func (p *Parser) AnswerHeader() (ResourceHeader, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:709
	_go_fuzz_dep_.CoverTab[2656]++
										return p.resourceHeader(sectionAnswers)
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:710
	// _ = "end of CoverTab[2656]"
}

// Answer parses a single Answer Resource.
func (p *Parser) Answer() (Resource, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:714
	_go_fuzz_dep_.CoverTab[2657]++
										return p.resource(sectionAnswers)
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:715
	// _ = "end of CoverTab[2657]"
}

// AllAnswers parses all Answer Resources.
func (p *Parser) AllAnswers() ([]Resource, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:719
	_go_fuzz_dep_.CoverTab[2658]++

//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:725
	n := int(p.header.answers)
	if n > 20 {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:726
		_go_fuzz_dep_.CoverTab[2660]++
											n = 20
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:727
		// _ = "end of CoverTab[2660]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:728
		_go_fuzz_dep_.CoverTab[2661]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:728
		// _ = "end of CoverTab[2661]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:728
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:728
	// _ = "end of CoverTab[2658]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:728
	_go_fuzz_dep_.CoverTab[2659]++
										as := make([]Resource, 0, n)
										for {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:730
		_go_fuzz_dep_.CoverTab[2662]++
											a, err := p.Answer()
											if err == ErrSectionDone {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:732
			_go_fuzz_dep_.CoverTab[2665]++
												return as, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:733
			// _ = "end of CoverTab[2665]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:734
			_go_fuzz_dep_.CoverTab[2666]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:734
			// _ = "end of CoverTab[2666]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:734
		}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:734
		// _ = "end of CoverTab[2662]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:734
		_go_fuzz_dep_.CoverTab[2663]++
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:735
			_go_fuzz_dep_.CoverTab[2667]++
												return nil, err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:736
			// _ = "end of CoverTab[2667]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:737
			_go_fuzz_dep_.CoverTab[2668]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:737
			// _ = "end of CoverTab[2668]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:737
		}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:737
		// _ = "end of CoverTab[2663]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:737
		_go_fuzz_dep_.CoverTab[2664]++
											as = append(as, a)
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:738
		// _ = "end of CoverTab[2664]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:739
	// _ = "end of CoverTab[2659]"
}

// SkipAnswer skips a single Answer Resource.
func (p *Parser) SkipAnswer() error {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:743
	_go_fuzz_dep_.CoverTab[2669]++
										return p.skipResource(sectionAnswers)
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:744
	// _ = "end of CoverTab[2669]"
}

// SkipAllAnswers skips all Answer Resources.
func (p *Parser) SkipAllAnswers() error {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:748
	_go_fuzz_dep_.CoverTab[2670]++
										for {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:749
		_go_fuzz_dep_.CoverTab[2671]++
											if err := p.SkipAnswer(); err == ErrSectionDone {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:750
			_go_fuzz_dep_.CoverTab[2672]++
												return nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:751
			// _ = "end of CoverTab[2672]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:752
			_go_fuzz_dep_.CoverTab[2673]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:752
			if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:752
				_go_fuzz_dep_.CoverTab[2674]++
													return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:753
				// _ = "end of CoverTab[2674]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:754
				_go_fuzz_dep_.CoverTab[2675]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:754
				// _ = "end of CoverTab[2675]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:754
			}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:754
			// _ = "end of CoverTab[2673]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:754
		}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:754
		// _ = "end of CoverTab[2671]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:755
	// _ = "end of CoverTab[2670]"
}

// AuthorityHeader parses a single Authority ResourceHeader.
func (p *Parser) AuthorityHeader() (ResourceHeader, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:759
	_go_fuzz_dep_.CoverTab[2676]++
										return p.resourceHeader(sectionAuthorities)
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:760
	// _ = "end of CoverTab[2676]"
}

// Authority parses a single Authority Resource.
func (p *Parser) Authority() (Resource, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:764
	_go_fuzz_dep_.CoverTab[2677]++
										return p.resource(sectionAuthorities)
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:765
	// _ = "end of CoverTab[2677]"
}

// AllAuthorities parses all Authority Resources.
func (p *Parser) AllAuthorities() ([]Resource, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:769
	_go_fuzz_dep_.CoverTab[2678]++

//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:775
	n := int(p.header.authorities)
	if n > 10 {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:776
		_go_fuzz_dep_.CoverTab[2680]++
											n = 10
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:777
		// _ = "end of CoverTab[2680]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:778
		_go_fuzz_dep_.CoverTab[2681]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:778
		// _ = "end of CoverTab[2681]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:778
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:778
	// _ = "end of CoverTab[2678]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:778
	_go_fuzz_dep_.CoverTab[2679]++
										as := make([]Resource, 0, n)
										for {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:780
		_go_fuzz_dep_.CoverTab[2682]++
											a, err := p.Authority()
											if err == ErrSectionDone {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:782
			_go_fuzz_dep_.CoverTab[2685]++
												return as, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:783
			// _ = "end of CoverTab[2685]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:784
			_go_fuzz_dep_.CoverTab[2686]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:784
			// _ = "end of CoverTab[2686]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:784
		}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:784
		// _ = "end of CoverTab[2682]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:784
		_go_fuzz_dep_.CoverTab[2683]++
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:785
			_go_fuzz_dep_.CoverTab[2687]++
												return nil, err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:786
			// _ = "end of CoverTab[2687]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:787
			_go_fuzz_dep_.CoverTab[2688]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:787
			// _ = "end of CoverTab[2688]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:787
		}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:787
		// _ = "end of CoverTab[2683]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:787
		_go_fuzz_dep_.CoverTab[2684]++
											as = append(as, a)
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:788
		// _ = "end of CoverTab[2684]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:789
	// _ = "end of CoverTab[2679]"
}

// SkipAuthority skips a single Authority Resource.
func (p *Parser) SkipAuthority() error {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:793
	_go_fuzz_dep_.CoverTab[2689]++
										return p.skipResource(sectionAuthorities)
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:794
	// _ = "end of CoverTab[2689]"
}

// SkipAllAuthorities skips all Authority Resources.
func (p *Parser) SkipAllAuthorities() error {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:798
	_go_fuzz_dep_.CoverTab[2690]++
										for {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:799
		_go_fuzz_dep_.CoverTab[2691]++
											if err := p.SkipAuthority(); err == ErrSectionDone {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:800
			_go_fuzz_dep_.CoverTab[2692]++
												return nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:801
			// _ = "end of CoverTab[2692]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:802
			_go_fuzz_dep_.CoverTab[2693]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:802
			if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:802
				_go_fuzz_dep_.CoverTab[2694]++
													return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:803
				// _ = "end of CoverTab[2694]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:804
				_go_fuzz_dep_.CoverTab[2695]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:804
				// _ = "end of CoverTab[2695]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:804
			}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:804
			// _ = "end of CoverTab[2693]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:804
		}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:804
		// _ = "end of CoverTab[2691]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:805
	// _ = "end of CoverTab[2690]"
}

// AdditionalHeader parses a single Additional ResourceHeader.
func (p *Parser) AdditionalHeader() (ResourceHeader, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:809
	_go_fuzz_dep_.CoverTab[2696]++
										return p.resourceHeader(sectionAdditionals)
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:810
	// _ = "end of CoverTab[2696]"
}

// Additional parses a single Additional Resource.
func (p *Parser) Additional() (Resource, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:814
	_go_fuzz_dep_.CoverTab[2697]++
										return p.resource(sectionAdditionals)
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:815
	// _ = "end of CoverTab[2697]"
}

// AllAdditionals parses all Additional Resources.
func (p *Parser) AllAdditionals() ([]Resource, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:819
	_go_fuzz_dep_.CoverTab[2698]++

//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:825
	n := int(p.header.additionals)
	if n > 10 {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:826
		_go_fuzz_dep_.CoverTab[2700]++
											n = 10
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:827
		// _ = "end of CoverTab[2700]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:828
		_go_fuzz_dep_.CoverTab[2701]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:828
		// _ = "end of CoverTab[2701]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:828
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:828
	// _ = "end of CoverTab[2698]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:828
	_go_fuzz_dep_.CoverTab[2699]++
										as := make([]Resource, 0, n)
										for {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:830
		_go_fuzz_dep_.CoverTab[2702]++
											a, err := p.Additional()
											if err == ErrSectionDone {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:832
			_go_fuzz_dep_.CoverTab[2705]++
												return as, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:833
			// _ = "end of CoverTab[2705]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:834
			_go_fuzz_dep_.CoverTab[2706]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:834
			// _ = "end of CoverTab[2706]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:834
		}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:834
		// _ = "end of CoverTab[2702]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:834
		_go_fuzz_dep_.CoverTab[2703]++
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:835
			_go_fuzz_dep_.CoverTab[2707]++
												return nil, err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:836
			// _ = "end of CoverTab[2707]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:837
			_go_fuzz_dep_.CoverTab[2708]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:837
			// _ = "end of CoverTab[2708]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:837
		}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:837
		// _ = "end of CoverTab[2703]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:837
		_go_fuzz_dep_.CoverTab[2704]++
											as = append(as, a)
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:838
		// _ = "end of CoverTab[2704]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:839
	// _ = "end of CoverTab[2699]"
}

// SkipAdditional skips a single Additional Resource.
func (p *Parser) SkipAdditional() error {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:843
	_go_fuzz_dep_.CoverTab[2709]++
										return p.skipResource(sectionAdditionals)
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:844
	// _ = "end of CoverTab[2709]"
}

// SkipAllAdditionals skips all Additional Resources.
func (p *Parser) SkipAllAdditionals() error {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:848
	_go_fuzz_dep_.CoverTab[2710]++
										for {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:849
		_go_fuzz_dep_.CoverTab[2711]++
											if err := p.SkipAdditional(); err == ErrSectionDone {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:850
			_go_fuzz_dep_.CoverTab[2712]++
												return nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:851
			// _ = "end of CoverTab[2712]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:852
			_go_fuzz_dep_.CoverTab[2713]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:852
			if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:852
				_go_fuzz_dep_.CoverTab[2714]++
													return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:853
				// _ = "end of CoverTab[2714]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:854
				_go_fuzz_dep_.CoverTab[2715]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:854
				// _ = "end of CoverTab[2715]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:854
			}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:854
			// _ = "end of CoverTab[2713]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:854
		}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:854
		// _ = "end of CoverTab[2711]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:855
	// _ = "end of CoverTab[2710]"
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
	_go_fuzz_dep_.CoverTab[2716]++
										if !p.resHeaderValid || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:863
		_go_fuzz_dep_.CoverTab[2719]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:863
		return p.resHeader.Type != TypeCNAME
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:863
		// _ = "end of CoverTab[2719]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:863
	}() {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:863
		_go_fuzz_dep_.CoverTab[2720]++
											return CNAMEResource{}, ErrNotStarted
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:864
		// _ = "end of CoverTab[2720]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:865
		_go_fuzz_dep_.CoverTab[2721]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:865
		// _ = "end of CoverTab[2721]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:865
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:865
	// _ = "end of CoverTab[2716]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:865
	_go_fuzz_dep_.CoverTab[2717]++
										r, err := unpackCNAMEResource(p.msg, p.off)
										if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:867
		_go_fuzz_dep_.CoverTab[2722]++
											return CNAMEResource{}, err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:868
		// _ = "end of CoverTab[2722]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:869
		_go_fuzz_dep_.CoverTab[2723]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:869
		// _ = "end of CoverTab[2723]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:869
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:869
	// _ = "end of CoverTab[2717]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:869
	_go_fuzz_dep_.CoverTab[2718]++
										p.off += int(p.resHeader.Length)
										p.resHeaderValid = false
										p.index++
										return r, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:873
	// _ = "end of CoverTab[2718]"
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
	_go_fuzz_dep_.CoverTab[2724]++
										if !p.resHeaderValid || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:881
		_go_fuzz_dep_.CoverTab[2727]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:881
		return p.resHeader.Type != TypeMX
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:881
		// _ = "end of CoverTab[2727]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:881
	}() {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:881
		_go_fuzz_dep_.CoverTab[2728]++
											return MXResource{}, ErrNotStarted
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:882
		// _ = "end of CoverTab[2728]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:883
		_go_fuzz_dep_.CoverTab[2729]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:883
		// _ = "end of CoverTab[2729]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:883
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:883
	// _ = "end of CoverTab[2724]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:883
	_go_fuzz_dep_.CoverTab[2725]++
										r, err := unpackMXResource(p.msg, p.off)
										if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:885
		_go_fuzz_dep_.CoverTab[2730]++
											return MXResource{}, err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:886
		// _ = "end of CoverTab[2730]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:887
		_go_fuzz_dep_.CoverTab[2731]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:887
		// _ = "end of CoverTab[2731]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:887
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:887
	// _ = "end of CoverTab[2725]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:887
	_go_fuzz_dep_.CoverTab[2726]++
										p.off += int(p.resHeader.Length)
										p.resHeaderValid = false
										p.index++
										return r, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:891
	// _ = "end of CoverTab[2726]"
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
	_go_fuzz_dep_.CoverTab[2732]++
										if !p.resHeaderValid || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:899
		_go_fuzz_dep_.CoverTab[2735]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:899
		return p.resHeader.Type != TypeNS
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:899
		// _ = "end of CoverTab[2735]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:899
	}() {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:899
		_go_fuzz_dep_.CoverTab[2736]++
											return NSResource{}, ErrNotStarted
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:900
		// _ = "end of CoverTab[2736]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:901
		_go_fuzz_dep_.CoverTab[2737]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:901
		// _ = "end of CoverTab[2737]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:901
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:901
	// _ = "end of CoverTab[2732]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:901
	_go_fuzz_dep_.CoverTab[2733]++
										r, err := unpackNSResource(p.msg, p.off)
										if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:903
		_go_fuzz_dep_.CoverTab[2738]++
											return NSResource{}, err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:904
		// _ = "end of CoverTab[2738]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:905
		_go_fuzz_dep_.CoverTab[2739]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:905
		// _ = "end of CoverTab[2739]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:905
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:905
	// _ = "end of CoverTab[2733]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:905
	_go_fuzz_dep_.CoverTab[2734]++
										p.off += int(p.resHeader.Length)
										p.resHeaderValid = false
										p.index++
										return r, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:909
	// _ = "end of CoverTab[2734]"
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
	_go_fuzz_dep_.CoverTab[2740]++
										if !p.resHeaderValid || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:917
		_go_fuzz_dep_.CoverTab[2743]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:917
		return p.resHeader.Type != TypePTR
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:917
		// _ = "end of CoverTab[2743]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:917
	}() {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:917
		_go_fuzz_dep_.CoverTab[2744]++
											return PTRResource{}, ErrNotStarted
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:918
		// _ = "end of CoverTab[2744]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:919
		_go_fuzz_dep_.CoverTab[2745]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:919
		// _ = "end of CoverTab[2745]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:919
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:919
	// _ = "end of CoverTab[2740]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:919
	_go_fuzz_dep_.CoverTab[2741]++
										r, err := unpackPTRResource(p.msg, p.off)
										if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:921
		_go_fuzz_dep_.CoverTab[2746]++
											return PTRResource{}, err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:922
		// _ = "end of CoverTab[2746]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:923
		_go_fuzz_dep_.CoverTab[2747]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:923
		// _ = "end of CoverTab[2747]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:923
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:923
	// _ = "end of CoverTab[2741]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:923
	_go_fuzz_dep_.CoverTab[2742]++
										p.off += int(p.resHeader.Length)
										p.resHeaderValid = false
										p.index++
										return r, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:927
	// _ = "end of CoverTab[2742]"
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
	_go_fuzz_dep_.CoverTab[2748]++
										if !p.resHeaderValid || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:935
		_go_fuzz_dep_.CoverTab[2751]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:935
		return p.resHeader.Type != TypeSOA
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:935
		// _ = "end of CoverTab[2751]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:935
	}() {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:935
		_go_fuzz_dep_.CoverTab[2752]++
											return SOAResource{}, ErrNotStarted
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:936
		// _ = "end of CoverTab[2752]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:937
		_go_fuzz_dep_.CoverTab[2753]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:937
		// _ = "end of CoverTab[2753]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:937
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:937
	// _ = "end of CoverTab[2748]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:937
	_go_fuzz_dep_.CoverTab[2749]++
										r, err := unpackSOAResource(p.msg, p.off)
										if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:939
		_go_fuzz_dep_.CoverTab[2754]++
											return SOAResource{}, err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:940
		// _ = "end of CoverTab[2754]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:941
		_go_fuzz_dep_.CoverTab[2755]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:941
		// _ = "end of CoverTab[2755]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:941
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:941
	// _ = "end of CoverTab[2749]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:941
	_go_fuzz_dep_.CoverTab[2750]++
										p.off += int(p.resHeader.Length)
										p.resHeaderValid = false
										p.index++
										return r, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:945
	// _ = "end of CoverTab[2750]"
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
	_go_fuzz_dep_.CoverTab[2756]++
										if !p.resHeaderValid || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:953
		_go_fuzz_dep_.CoverTab[2759]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:953
		return p.resHeader.Type != TypeTXT
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:953
		// _ = "end of CoverTab[2759]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:953
	}() {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:953
		_go_fuzz_dep_.CoverTab[2760]++
											return TXTResource{}, ErrNotStarted
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:954
		// _ = "end of CoverTab[2760]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:955
		_go_fuzz_dep_.CoverTab[2761]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:955
		// _ = "end of CoverTab[2761]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:955
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:955
	// _ = "end of CoverTab[2756]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:955
	_go_fuzz_dep_.CoverTab[2757]++
										r, err := unpackTXTResource(p.msg, p.off, p.resHeader.Length)
										if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:957
		_go_fuzz_dep_.CoverTab[2762]++
											return TXTResource{}, err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:958
		// _ = "end of CoverTab[2762]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:959
		_go_fuzz_dep_.CoverTab[2763]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:959
		// _ = "end of CoverTab[2763]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:959
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:959
	// _ = "end of CoverTab[2757]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:959
	_go_fuzz_dep_.CoverTab[2758]++
										p.off += int(p.resHeader.Length)
										p.resHeaderValid = false
										p.index++
										return r, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:963
	// _ = "end of CoverTab[2758]"
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
	_go_fuzz_dep_.CoverTab[2764]++
										if !p.resHeaderValid || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:971
		_go_fuzz_dep_.CoverTab[2767]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:971
		return p.resHeader.Type != TypeSRV
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:971
		// _ = "end of CoverTab[2767]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:971
	}() {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:971
		_go_fuzz_dep_.CoverTab[2768]++
											return SRVResource{}, ErrNotStarted
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:972
		// _ = "end of CoverTab[2768]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:973
		_go_fuzz_dep_.CoverTab[2769]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:973
		// _ = "end of CoverTab[2769]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:973
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:973
	// _ = "end of CoverTab[2764]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:973
	_go_fuzz_dep_.CoverTab[2765]++
										r, err := unpackSRVResource(p.msg, p.off)
										if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:975
		_go_fuzz_dep_.CoverTab[2770]++
											return SRVResource{}, err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:976
		// _ = "end of CoverTab[2770]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:977
		_go_fuzz_dep_.CoverTab[2771]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:977
		// _ = "end of CoverTab[2771]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:977
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:977
	// _ = "end of CoverTab[2765]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:977
	_go_fuzz_dep_.CoverTab[2766]++
										p.off += int(p.resHeader.Length)
										p.resHeaderValid = false
										p.index++
										return r, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:981
	// _ = "end of CoverTab[2766]"
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
	_go_fuzz_dep_.CoverTab[2772]++
										if !p.resHeaderValid || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:989
		_go_fuzz_dep_.CoverTab[2775]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:989
		return p.resHeader.Type != TypeA
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:989
		// _ = "end of CoverTab[2775]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:989
	}() {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:989
		_go_fuzz_dep_.CoverTab[2776]++
											return AResource{}, ErrNotStarted
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:990
		// _ = "end of CoverTab[2776]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:991
		_go_fuzz_dep_.CoverTab[2777]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:991
		// _ = "end of CoverTab[2777]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:991
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:991
	// _ = "end of CoverTab[2772]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:991
	_go_fuzz_dep_.CoverTab[2773]++
										r, err := unpackAResource(p.msg, p.off)
										if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:993
		_go_fuzz_dep_.CoverTab[2778]++
											return AResource{}, err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:994
		// _ = "end of CoverTab[2778]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:995
		_go_fuzz_dep_.CoverTab[2779]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:995
		// _ = "end of CoverTab[2779]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:995
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:995
	// _ = "end of CoverTab[2773]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:995
	_go_fuzz_dep_.CoverTab[2774]++
										p.off += int(p.resHeader.Length)
										p.resHeaderValid = false
										p.index++
										return r, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:999
	// _ = "end of CoverTab[2774]"
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
	_go_fuzz_dep_.CoverTab[2780]++
											if !p.resHeaderValid || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1007
		_go_fuzz_dep_.CoverTab[2783]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1007
		return p.resHeader.Type != TypeAAAA
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1007
		// _ = "end of CoverTab[2783]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1007
	}() {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1007
		_go_fuzz_dep_.CoverTab[2784]++
												return AAAAResource{}, ErrNotStarted
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1008
		// _ = "end of CoverTab[2784]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1009
		_go_fuzz_dep_.CoverTab[2785]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1009
		// _ = "end of CoverTab[2785]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1009
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1009
	// _ = "end of CoverTab[2780]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1009
	_go_fuzz_dep_.CoverTab[2781]++
											r, err := unpackAAAAResource(p.msg, p.off)
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1011
		_go_fuzz_dep_.CoverTab[2786]++
												return AAAAResource{}, err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1012
		// _ = "end of CoverTab[2786]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1013
		_go_fuzz_dep_.CoverTab[2787]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1013
		// _ = "end of CoverTab[2787]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1013
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1013
	// _ = "end of CoverTab[2781]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1013
	_go_fuzz_dep_.CoverTab[2782]++
											p.off += int(p.resHeader.Length)
											p.resHeaderValid = false
											p.index++
											return r, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1017
	// _ = "end of CoverTab[2782]"
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
	_go_fuzz_dep_.CoverTab[2788]++
											if !p.resHeaderValid || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1025
		_go_fuzz_dep_.CoverTab[2791]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1025
		return p.resHeader.Type != TypeOPT
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1025
		// _ = "end of CoverTab[2791]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1025
	}() {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1025
		_go_fuzz_dep_.CoverTab[2792]++
												return OPTResource{}, ErrNotStarted
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1026
		// _ = "end of CoverTab[2792]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1027
		_go_fuzz_dep_.CoverTab[2793]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1027
		// _ = "end of CoverTab[2793]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1027
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1027
	// _ = "end of CoverTab[2788]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1027
	_go_fuzz_dep_.CoverTab[2789]++
											r, err := unpackOPTResource(p.msg, p.off, p.resHeader.Length)
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1029
		_go_fuzz_dep_.CoverTab[2794]++
												return OPTResource{}, err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1030
		// _ = "end of CoverTab[2794]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1031
		_go_fuzz_dep_.CoverTab[2795]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1031
		// _ = "end of CoverTab[2795]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1031
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1031
	// _ = "end of CoverTab[2789]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1031
	_go_fuzz_dep_.CoverTab[2790]++
											p.off += int(p.resHeader.Length)
											p.resHeaderValid = false
											p.index++
											return r, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1035
	// _ = "end of CoverTab[2790]"
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
	_go_fuzz_dep_.CoverTab[2796]++
											if !p.resHeaderValid {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1043
		_go_fuzz_dep_.CoverTab[2799]++
												return UnknownResource{}, ErrNotStarted
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1044
		// _ = "end of CoverTab[2799]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1045
		_go_fuzz_dep_.CoverTab[2800]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1045
		// _ = "end of CoverTab[2800]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1045
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1045
	// _ = "end of CoverTab[2796]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1045
	_go_fuzz_dep_.CoverTab[2797]++
											r, err := unpackUnknownResource(p.resHeader.Type, p.msg, p.off, p.resHeader.Length)
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1047
		_go_fuzz_dep_.CoverTab[2801]++
												return UnknownResource{}, err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1048
		// _ = "end of CoverTab[2801]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1049
		_go_fuzz_dep_.CoverTab[2802]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1049
		// _ = "end of CoverTab[2802]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1049
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1049
	// _ = "end of CoverTab[2797]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1049
	_go_fuzz_dep_.CoverTab[2798]++
											p.off += int(p.resHeader.Length)
											p.resHeaderValid = false
											p.index++
											return r, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1053
	// _ = "end of CoverTab[2798]"
}

// Unpack parses a full Message.
func (m *Message) Unpack(msg []byte) error {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1057
	_go_fuzz_dep_.CoverTab[2803]++
											var p Parser
											var err error
											if m.Header, err = p.Start(msg); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1060
		_go_fuzz_dep_.CoverTab[2809]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1061
		// _ = "end of CoverTab[2809]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1062
		_go_fuzz_dep_.CoverTab[2810]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1062
		// _ = "end of CoverTab[2810]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1062
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1062
	// _ = "end of CoverTab[2803]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1062
	_go_fuzz_dep_.CoverTab[2804]++
											if m.Questions, err = p.AllQuestions(); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1063
		_go_fuzz_dep_.CoverTab[2811]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1064
		// _ = "end of CoverTab[2811]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1065
		_go_fuzz_dep_.CoverTab[2812]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1065
		// _ = "end of CoverTab[2812]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1065
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1065
	// _ = "end of CoverTab[2804]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1065
	_go_fuzz_dep_.CoverTab[2805]++
											if m.Answers, err = p.AllAnswers(); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1066
		_go_fuzz_dep_.CoverTab[2813]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1067
		// _ = "end of CoverTab[2813]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1068
		_go_fuzz_dep_.CoverTab[2814]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1068
		// _ = "end of CoverTab[2814]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1068
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1068
	// _ = "end of CoverTab[2805]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1068
	_go_fuzz_dep_.CoverTab[2806]++
											if m.Authorities, err = p.AllAuthorities(); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1069
		_go_fuzz_dep_.CoverTab[2815]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1070
		// _ = "end of CoverTab[2815]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1071
		_go_fuzz_dep_.CoverTab[2816]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1071
		// _ = "end of CoverTab[2816]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1071
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1071
	// _ = "end of CoverTab[2806]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1071
	_go_fuzz_dep_.CoverTab[2807]++
											if m.Additionals, err = p.AllAdditionals(); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1072
		_go_fuzz_dep_.CoverTab[2817]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1073
		// _ = "end of CoverTab[2817]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1074
		_go_fuzz_dep_.CoverTab[2818]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1074
		// _ = "end of CoverTab[2818]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1074
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1074
	// _ = "end of CoverTab[2807]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1074
	_go_fuzz_dep_.CoverTab[2808]++
											return nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1075
	// _ = "end of CoverTab[2808]"
}

// Pack packs a full Message.
func (m *Message) Pack() ([]byte, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1079
	_go_fuzz_dep_.CoverTab[2819]++
											return m.AppendPack(make([]byte, 0, packStartingCap))
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1080
	// _ = "end of CoverTab[2819]"
}

// AppendPack is like Pack but appends the full Message to b and returns the
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1083
// extended buffer.
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1085
func (m *Message) AppendPack(b []byte) ([]byte, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1085
	_go_fuzz_dep_.CoverTab[2820]++

//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1089
	if len(m.Questions) > int(^uint16(0)) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1089
		_go_fuzz_dep_.CoverTab[2829]++
												return nil, errTooManyQuestions
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1090
		// _ = "end of CoverTab[2829]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1091
		_go_fuzz_dep_.CoverTab[2830]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1091
		// _ = "end of CoverTab[2830]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1091
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1091
	// _ = "end of CoverTab[2820]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1091
	_go_fuzz_dep_.CoverTab[2821]++
											if len(m.Answers) > int(^uint16(0)) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1092
		_go_fuzz_dep_.CoverTab[2831]++
												return nil, errTooManyAnswers
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1093
		// _ = "end of CoverTab[2831]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1094
		_go_fuzz_dep_.CoverTab[2832]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1094
		// _ = "end of CoverTab[2832]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1094
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1094
	// _ = "end of CoverTab[2821]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1094
	_go_fuzz_dep_.CoverTab[2822]++
											if len(m.Authorities) > int(^uint16(0)) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1095
		_go_fuzz_dep_.CoverTab[2833]++
												return nil, errTooManyAuthorities
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1096
		// _ = "end of CoverTab[2833]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1097
		_go_fuzz_dep_.CoverTab[2834]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1097
		// _ = "end of CoverTab[2834]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1097
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1097
	// _ = "end of CoverTab[2822]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1097
	_go_fuzz_dep_.CoverTab[2823]++
											if len(m.Additionals) > int(^uint16(0)) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1098
		_go_fuzz_dep_.CoverTab[2835]++
												return nil, errTooManyAdditionals
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1099
		// _ = "end of CoverTab[2835]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1100
		_go_fuzz_dep_.CoverTab[2836]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1100
		// _ = "end of CoverTab[2836]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1100
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1100
	// _ = "end of CoverTab[2823]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1100
	_go_fuzz_dep_.CoverTab[2824]++

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
		_go_fuzz_dep_.CoverTab[2837]++
												var err error
												if msg, err = m.Questions[i].pack(msg, compression, compressionOff); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1125
			_go_fuzz_dep_.CoverTab[2838]++
													return nil, &nestedError{"packing Question", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1126
			// _ = "end of CoverTab[2838]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1127
			_go_fuzz_dep_.CoverTab[2839]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1127
			// _ = "end of CoverTab[2839]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1127
		}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1127
		// _ = "end of CoverTab[2837]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1128
	// _ = "end of CoverTab[2824]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1128
	_go_fuzz_dep_.CoverTab[2825]++
											for i := range m.Answers {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1129
		_go_fuzz_dep_.CoverTab[2840]++
												var err error
												if msg, err = m.Answers[i].pack(msg, compression, compressionOff); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1131
			_go_fuzz_dep_.CoverTab[2841]++
													return nil, &nestedError{"packing Answer", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1132
			// _ = "end of CoverTab[2841]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1133
			_go_fuzz_dep_.CoverTab[2842]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1133
			// _ = "end of CoverTab[2842]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1133
		}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1133
		// _ = "end of CoverTab[2840]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1134
	// _ = "end of CoverTab[2825]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1134
	_go_fuzz_dep_.CoverTab[2826]++
											for i := range m.Authorities {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1135
		_go_fuzz_dep_.CoverTab[2843]++
												var err error
												if msg, err = m.Authorities[i].pack(msg, compression, compressionOff); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1137
			_go_fuzz_dep_.CoverTab[2844]++
													return nil, &nestedError{"packing Authority", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1138
			// _ = "end of CoverTab[2844]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1139
			_go_fuzz_dep_.CoverTab[2845]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1139
			// _ = "end of CoverTab[2845]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1139
		}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1139
		// _ = "end of CoverTab[2843]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1140
	// _ = "end of CoverTab[2826]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1140
	_go_fuzz_dep_.CoverTab[2827]++
											for i := range m.Additionals {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1141
		_go_fuzz_dep_.CoverTab[2846]++
												var err error
												if msg, err = m.Additionals[i].pack(msg, compression, compressionOff); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1143
			_go_fuzz_dep_.CoverTab[2847]++
													return nil, &nestedError{"packing Additional", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1144
			// _ = "end of CoverTab[2847]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1145
			_go_fuzz_dep_.CoverTab[2848]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1145
			// _ = "end of CoverTab[2848]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1145
		}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1145
		// _ = "end of CoverTab[2846]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1146
	// _ = "end of CoverTab[2827]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1146
	_go_fuzz_dep_.CoverTab[2828]++

											return msg, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1148
	// _ = "end of CoverTab[2828]"
}

// GoString implements fmt.GoStringer.GoString.
func (m *Message) GoString() string {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1152
	_go_fuzz_dep_.CoverTab[2849]++
											s := "dnsmessage.Message{Header: " + m.Header.GoString() + ", " +
		"Questions: []dnsmessage.Question{"
	if len(m.Questions) > 0 {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1155
		_go_fuzz_dep_.CoverTab[2854]++
												s += m.Questions[0].GoString()
												for _, q := range m.Questions[1:] {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1157
			_go_fuzz_dep_.CoverTab[2855]++
													s += ", " + q.GoString()
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1158
			// _ = "end of CoverTab[2855]"
		}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1159
		// _ = "end of CoverTab[2854]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1160
		_go_fuzz_dep_.CoverTab[2856]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1160
		// _ = "end of CoverTab[2856]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1160
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1160
	// _ = "end of CoverTab[2849]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1160
	_go_fuzz_dep_.CoverTab[2850]++
											s += "}, Answers: []dnsmessage.Resource{"
											if len(m.Answers) > 0 {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1162
		_go_fuzz_dep_.CoverTab[2857]++
												s += m.Answers[0].GoString()
												for _, a := range m.Answers[1:] {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1164
			_go_fuzz_dep_.CoverTab[2858]++
													s += ", " + a.GoString()
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1165
			// _ = "end of CoverTab[2858]"
		}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1166
		// _ = "end of CoverTab[2857]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1167
		_go_fuzz_dep_.CoverTab[2859]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1167
		// _ = "end of CoverTab[2859]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1167
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1167
	// _ = "end of CoverTab[2850]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1167
	_go_fuzz_dep_.CoverTab[2851]++
											s += "}, Authorities: []dnsmessage.Resource{"
											if len(m.Authorities) > 0 {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1169
		_go_fuzz_dep_.CoverTab[2860]++
												s += m.Authorities[0].GoString()
												for _, a := range m.Authorities[1:] {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1171
			_go_fuzz_dep_.CoverTab[2861]++
													s += ", " + a.GoString()
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1172
			// _ = "end of CoverTab[2861]"
		}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1173
		// _ = "end of CoverTab[2860]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1174
		_go_fuzz_dep_.CoverTab[2862]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1174
		// _ = "end of CoverTab[2862]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1174
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1174
	// _ = "end of CoverTab[2851]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1174
	_go_fuzz_dep_.CoverTab[2852]++
											s += "}, Additionals: []dnsmessage.Resource{"
											if len(m.Additionals) > 0 {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1176
		_go_fuzz_dep_.CoverTab[2863]++
												s += m.Additionals[0].GoString()
												for _, a := range m.Additionals[1:] {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1178
			_go_fuzz_dep_.CoverTab[2864]++
													s += ", " + a.GoString()
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1179
			// _ = "end of CoverTab[2864]"
		}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1180
		// _ = "end of CoverTab[2863]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1181
		_go_fuzz_dep_.CoverTab[2865]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1181
		// _ = "end of CoverTab[2865]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1181
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1181
	// _ = "end of CoverTab[2852]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1181
	_go_fuzz_dep_.CoverTab[2853]++
											return s + "}}"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1182
	// _ = "end of CoverTab[2853]"
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
	_go_fuzz_dep_.CoverTab[2866]++
											if buf == nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1226
		_go_fuzz_dep_.CoverTab[2868]++
												buf = make([]byte, 0, packStartingCap)
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1227
		// _ = "end of CoverTab[2868]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1228
		_go_fuzz_dep_.CoverTab[2869]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1228
		// _ = "end of CoverTab[2869]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1228
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1228
	// _ = "end of CoverTab[2866]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1228
	_go_fuzz_dep_.CoverTab[2867]++
											b := Builder{msg: buf, start: len(buf)}
											b.header.id, b.header.bits = h.pack()
											var hb [headerLen]byte
											b.msg = append(b.msg, hb[:]...)
											b.section = sectionHeader
											return b
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1234
	// _ = "end of CoverTab[2867]"
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
	_go_fuzz_dep_.CoverTab[2870]++
											b.compression = map[string]int{}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1249
	// _ = "end of CoverTab[2870]"
}

func (b *Builder) startCheck(s section) error {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1252
	_go_fuzz_dep_.CoverTab[2871]++
											if b.section <= sectionNotStarted {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1253
		_go_fuzz_dep_.CoverTab[2874]++
												return ErrNotStarted
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1254
		// _ = "end of CoverTab[2874]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1255
		_go_fuzz_dep_.CoverTab[2875]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1255
		// _ = "end of CoverTab[2875]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1255
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1255
	// _ = "end of CoverTab[2871]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1255
	_go_fuzz_dep_.CoverTab[2872]++
											if b.section > s {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1256
		_go_fuzz_dep_.CoverTab[2876]++
												return ErrSectionDone
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1257
		// _ = "end of CoverTab[2876]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1258
		_go_fuzz_dep_.CoverTab[2877]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1258
		// _ = "end of CoverTab[2877]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1258
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1258
	// _ = "end of CoverTab[2872]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1258
	_go_fuzz_dep_.CoverTab[2873]++
											return nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1259
	// _ = "end of CoverTab[2873]"
}

// StartQuestions prepares the builder for packing Questions.
func (b *Builder) StartQuestions() error {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1263
	_go_fuzz_dep_.CoverTab[2878]++
											if err := b.startCheck(sectionQuestions); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1264
		_go_fuzz_dep_.CoverTab[2880]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1265
		// _ = "end of CoverTab[2880]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1266
		_go_fuzz_dep_.CoverTab[2881]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1266
		// _ = "end of CoverTab[2881]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1266
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1266
	// _ = "end of CoverTab[2878]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1266
	_go_fuzz_dep_.CoverTab[2879]++
											b.section = sectionQuestions
											return nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1268
	// _ = "end of CoverTab[2879]"
}

// StartAnswers prepares the builder for packing Answers.
func (b *Builder) StartAnswers() error {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1272
	_go_fuzz_dep_.CoverTab[2882]++
											if err := b.startCheck(sectionAnswers); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1273
		_go_fuzz_dep_.CoverTab[2884]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1274
		// _ = "end of CoverTab[2884]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1275
		_go_fuzz_dep_.CoverTab[2885]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1275
		// _ = "end of CoverTab[2885]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1275
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1275
	// _ = "end of CoverTab[2882]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1275
	_go_fuzz_dep_.CoverTab[2883]++
											b.section = sectionAnswers
											return nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1277
	// _ = "end of CoverTab[2883]"
}

// StartAuthorities prepares the builder for packing Authorities.
func (b *Builder) StartAuthorities() error {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1281
	_go_fuzz_dep_.CoverTab[2886]++
											if err := b.startCheck(sectionAuthorities); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1282
		_go_fuzz_dep_.CoverTab[2888]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1283
		// _ = "end of CoverTab[2888]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1284
		_go_fuzz_dep_.CoverTab[2889]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1284
		// _ = "end of CoverTab[2889]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1284
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1284
	// _ = "end of CoverTab[2886]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1284
	_go_fuzz_dep_.CoverTab[2887]++
											b.section = sectionAuthorities
											return nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1286
	// _ = "end of CoverTab[2887]"
}

// StartAdditionals prepares the builder for packing Additionals.
func (b *Builder) StartAdditionals() error {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1290
	_go_fuzz_dep_.CoverTab[2890]++
											if err := b.startCheck(sectionAdditionals); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1291
		_go_fuzz_dep_.CoverTab[2892]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1292
		// _ = "end of CoverTab[2892]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1293
		_go_fuzz_dep_.CoverTab[2893]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1293
		// _ = "end of CoverTab[2893]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1293
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1293
	// _ = "end of CoverTab[2890]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1293
	_go_fuzz_dep_.CoverTab[2891]++
											b.section = sectionAdditionals
											return nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1295
	// _ = "end of CoverTab[2891]"
}

func (b *Builder) incrementSectionCount() error {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1298
	_go_fuzz_dep_.CoverTab[2894]++
											var count *uint16
											var err error
											switch b.section {
	case sectionQuestions:
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1302
		_go_fuzz_dep_.CoverTab[2897]++
												count = &b.header.questions
												err = errTooManyQuestions
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1304
		// _ = "end of CoverTab[2897]"
	case sectionAnswers:
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1305
		_go_fuzz_dep_.CoverTab[2898]++
												count = &b.header.answers
												err = errTooManyAnswers
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1307
		// _ = "end of CoverTab[2898]"
	case sectionAuthorities:
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1308
		_go_fuzz_dep_.CoverTab[2899]++
												count = &b.header.authorities
												err = errTooManyAuthorities
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1310
		// _ = "end of CoverTab[2899]"
	case sectionAdditionals:
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1311
		_go_fuzz_dep_.CoverTab[2900]++
												count = &b.header.additionals
												err = errTooManyAdditionals
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1313
		// _ = "end of CoverTab[2900]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1313
	default:
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1313
		_go_fuzz_dep_.CoverTab[2901]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1313
		// _ = "end of CoverTab[2901]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1314
	// _ = "end of CoverTab[2894]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1314
	_go_fuzz_dep_.CoverTab[2895]++
											if *count == ^uint16(0) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1315
		_go_fuzz_dep_.CoverTab[2902]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1316
		// _ = "end of CoverTab[2902]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1317
		_go_fuzz_dep_.CoverTab[2903]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1317
		// _ = "end of CoverTab[2903]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1317
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1317
	// _ = "end of CoverTab[2895]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1317
	_go_fuzz_dep_.CoverTab[2896]++
											*count++
											return nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1319
	// _ = "end of CoverTab[2896]"
}

// Question adds a single Question.
func (b *Builder) Question(q Question) error {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1323
	_go_fuzz_dep_.CoverTab[2904]++
											if b.section < sectionQuestions {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1324
		_go_fuzz_dep_.CoverTab[2909]++
												return ErrNotStarted
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1325
		// _ = "end of CoverTab[2909]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1326
		_go_fuzz_dep_.CoverTab[2910]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1326
		// _ = "end of CoverTab[2910]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1326
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1326
	// _ = "end of CoverTab[2904]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1326
	_go_fuzz_dep_.CoverTab[2905]++
											if b.section > sectionQuestions {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1327
		_go_fuzz_dep_.CoverTab[2911]++
												return ErrSectionDone
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1328
		// _ = "end of CoverTab[2911]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1329
		_go_fuzz_dep_.CoverTab[2912]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1329
		// _ = "end of CoverTab[2912]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1329
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1329
	// _ = "end of CoverTab[2905]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1329
	_go_fuzz_dep_.CoverTab[2906]++
											msg, err := q.pack(b.msg, b.compression, b.start)
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1331
		_go_fuzz_dep_.CoverTab[2913]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1332
		// _ = "end of CoverTab[2913]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1333
		_go_fuzz_dep_.CoverTab[2914]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1333
		// _ = "end of CoverTab[2914]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1333
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1333
	// _ = "end of CoverTab[2906]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1333
	_go_fuzz_dep_.CoverTab[2907]++
											if err := b.incrementSectionCount(); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1334
		_go_fuzz_dep_.CoverTab[2915]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1335
		// _ = "end of CoverTab[2915]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1336
		_go_fuzz_dep_.CoverTab[2916]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1336
		// _ = "end of CoverTab[2916]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1336
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1336
	// _ = "end of CoverTab[2907]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1336
	_go_fuzz_dep_.CoverTab[2908]++
											b.msg = msg
											return nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1338
	// _ = "end of CoverTab[2908]"
}

func (b *Builder) checkResourceSection() error {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1341
	_go_fuzz_dep_.CoverTab[2917]++
											if b.section < sectionAnswers {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1342
		_go_fuzz_dep_.CoverTab[2920]++
												return ErrNotStarted
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1343
		// _ = "end of CoverTab[2920]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1344
		_go_fuzz_dep_.CoverTab[2921]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1344
		// _ = "end of CoverTab[2921]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1344
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1344
	// _ = "end of CoverTab[2917]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1344
	_go_fuzz_dep_.CoverTab[2918]++
											if b.section > sectionAdditionals {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1345
		_go_fuzz_dep_.CoverTab[2922]++
												return ErrSectionDone
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1346
		// _ = "end of CoverTab[2922]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1347
		_go_fuzz_dep_.CoverTab[2923]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1347
		// _ = "end of CoverTab[2923]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1347
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1347
	// _ = "end of CoverTab[2918]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1347
	_go_fuzz_dep_.CoverTab[2919]++
											return nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1348
	// _ = "end of CoverTab[2919]"
}

// CNAMEResource adds a single CNAMEResource.
func (b *Builder) CNAMEResource(h ResourceHeader, r CNAMEResource) error {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1352
	_go_fuzz_dep_.CoverTab[2924]++
											if err := b.checkResourceSection(); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1353
		_go_fuzz_dep_.CoverTab[2930]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1354
		// _ = "end of CoverTab[2930]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1355
		_go_fuzz_dep_.CoverTab[2931]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1355
		// _ = "end of CoverTab[2931]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1355
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1355
	// _ = "end of CoverTab[2924]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1355
	_go_fuzz_dep_.CoverTab[2925]++
											h.Type = r.realType()
											msg, lenOff, err := h.pack(b.msg, b.compression, b.start)
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1358
		_go_fuzz_dep_.CoverTab[2932]++
												return &nestedError{"ResourceHeader", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1359
		// _ = "end of CoverTab[2932]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1360
		_go_fuzz_dep_.CoverTab[2933]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1360
		// _ = "end of CoverTab[2933]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1360
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1360
	// _ = "end of CoverTab[2925]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1360
	_go_fuzz_dep_.CoverTab[2926]++
											preLen := len(msg)
											if msg, err = r.pack(msg, b.compression, b.start); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1362
		_go_fuzz_dep_.CoverTab[2934]++
												return &nestedError{"CNAMEResource body", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1363
		// _ = "end of CoverTab[2934]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1364
		_go_fuzz_dep_.CoverTab[2935]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1364
		// _ = "end of CoverTab[2935]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1364
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1364
	// _ = "end of CoverTab[2926]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1364
	_go_fuzz_dep_.CoverTab[2927]++
											if err := h.fixLen(msg, lenOff, preLen); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1365
		_go_fuzz_dep_.CoverTab[2936]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1366
		// _ = "end of CoverTab[2936]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1367
		_go_fuzz_dep_.CoverTab[2937]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1367
		// _ = "end of CoverTab[2937]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1367
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1367
	// _ = "end of CoverTab[2927]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1367
	_go_fuzz_dep_.CoverTab[2928]++
											if err := b.incrementSectionCount(); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1368
		_go_fuzz_dep_.CoverTab[2938]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1369
		// _ = "end of CoverTab[2938]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1370
		_go_fuzz_dep_.CoverTab[2939]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1370
		// _ = "end of CoverTab[2939]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1370
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1370
	// _ = "end of CoverTab[2928]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1370
	_go_fuzz_dep_.CoverTab[2929]++
											b.msg = msg
											return nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1372
	// _ = "end of CoverTab[2929]"
}

// MXResource adds a single MXResource.
func (b *Builder) MXResource(h ResourceHeader, r MXResource) error {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1376
	_go_fuzz_dep_.CoverTab[2940]++
											if err := b.checkResourceSection(); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1377
		_go_fuzz_dep_.CoverTab[2946]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1378
		// _ = "end of CoverTab[2946]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1379
		_go_fuzz_dep_.CoverTab[2947]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1379
		// _ = "end of CoverTab[2947]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1379
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1379
	// _ = "end of CoverTab[2940]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1379
	_go_fuzz_dep_.CoverTab[2941]++
											h.Type = r.realType()
											msg, lenOff, err := h.pack(b.msg, b.compression, b.start)
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1382
		_go_fuzz_dep_.CoverTab[2948]++
												return &nestedError{"ResourceHeader", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1383
		// _ = "end of CoverTab[2948]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1384
		_go_fuzz_dep_.CoverTab[2949]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1384
		// _ = "end of CoverTab[2949]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1384
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1384
	// _ = "end of CoverTab[2941]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1384
	_go_fuzz_dep_.CoverTab[2942]++
											preLen := len(msg)
											if msg, err = r.pack(msg, b.compression, b.start); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1386
		_go_fuzz_dep_.CoverTab[2950]++
												return &nestedError{"MXResource body", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1387
		// _ = "end of CoverTab[2950]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1388
		_go_fuzz_dep_.CoverTab[2951]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1388
		// _ = "end of CoverTab[2951]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1388
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1388
	// _ = "end of CoverTab[2942]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1388
	_go_fuzz_dep_.CoverTab[2943]++
											if err := h.fixLen(msg, lenOff, preLen); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1389
		_go_fuzz_dep_.CoverTab[2952]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1390
		// _ = "end of CoverTab[2952]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1391
		_go_fuzz_dep_.CoverTab[2953]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1391
		// _ = "end of CoverTab[2953]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1391
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1391
	// _ = "end of CoverTab[2943]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1391
	_go_fuzz_dep_.CoverTab[2944]++
											if err := b.incrementSectionCount(); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1392
		_go_fuzz_dep_.CoverTab[2954]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1393
		// _ = "end of CoverTab[2954]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1394
		_go_fuzz_dep_.CoverTab[2955]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1394
		// _ = "end of CoverTab[2955]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1394
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1394
	// _ = "end of CoverTab[2944]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1394
	_go_fuzz_dep_.CoverTab[2945]++
											b.msg = msg
											return nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1396
	// _ = "end of CoverTab[2945]"
}

// NSResource adds a single NSResource.
func (b *Builder) NSResource(h ResourceHeader, r NSResource) error {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1400
	_go_fuzz_dep_.CoverTab[2956]++
											if err := b.checkResourceSection(); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1401
		_go_fuzz_dep_.CoverTab[2962]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1402
		// _ = "end of CoverTab[2962]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1403
		_go_fuzz_dep_.CoverTab[2963]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1403
		// _ = "end of CoverTab[2963]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1403
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1403
	// _ = "end of CoverTab[2956]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1403
	_go_fuzz_dep_.CoverTab[2957]++
											h.Type = r.realType()
											msg, lenOff, err := h.pack(b.msg, b.compression, b.start)
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1406
		_go_fuzz_dep_.CoverTab[2964]++
												return &nestedError{"ResourceHeader", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1407
		// _ = "end of CoverTab[2964]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1408
		_go_fuzz_dep_.CoverTab[2965]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1408
		// _ = "end of CoverTab[2965]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1408
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1408
	// _ = "end of CoverTab[2957]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1408
	_go_fuzz_dep_.CoverTab[2958]++
											preLen := len(msg)
											if msg, err = r.pack(msg, b.compression, b.start); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1410
		_go_fuzz_dep_.CoverTab[2966]++
												return &nestedError{"NSResource body", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1411
		// _ = "end of CoverTab[2966]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1412
		_go_fuzz_dep_.CoverTab[2967]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1412
		// _ = "end of CoverTab[2967]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1412
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1412
	// _ = "end of CoverTab[2958]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1412
	_go_fuzz_dep_.CoverTab[2959]++
											if err := h.fixLen(msg, lenOff, preLen); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1413
		_go_fuzz_dep_.CoverTab[2968]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1414
		// _ = "end of CoverTab[2968]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1415
		_go_fuzz_dep_.CoverTab[2969]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1415
		// _ = "end of CoverTab[2969]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1415
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1415
	// _ = "end of CoverTab[2959]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1415
	_go_fuzz_dep_.CoverTab[2960]++
											if err := b.incrementSectionCount(); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1416
		_go_fuzz_dep_.CoverTab[2970]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1417
		// _ = "end of CoverTab[2970]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1418
		_go_fuzz_dep_.CoverTab[2971]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1418
		// _ = "end of CoverTab[2971]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1418
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1418
	// _ = "end of CoverTab[2960]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1418
	_go_fuzz_dep_.CoverTab[2961]++
											b.msg = msg
											return nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1420
	// _ = "end of CoverTab[2961]"
}

// PTRResource adds a single PTRResource.
func (b *Builder) PTRResource(h ResourceHeader, r PTRResource) error {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1424
	_go_fuzz_dep_.CoverTab[2972]++
											if err := b.checkResourceSection(); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1425
		_go_fuzz_dep_.CoverTab[2978]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1426
		// _ = "end of CoverTab[2978]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1427
		_go_fuzz_dep_.CoverTab[2979]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1427
		// _ = "end of CoverTab[2979]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1427
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1427
	// _ = "end of CoverTab[2972]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1427
	_go_fuzz_dep_.CoverTab[2973]++
											h.Type = r.realType()
											msg, lenOff, err := h.pack(b.msg, b.compression, b.start)
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1430
		_go_fuzz_dep_.CoverTab[2980]++
												return &nestedError{"ResourceHeader", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1431
		// _ = "end of CoverTab[2980]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1432
		_go_fuzz_dep_.CoverTab[2981]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1432
		// _ = "end of CoverTab[2981]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1432
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1432
	// _ = "end of CoverTab[2973]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1432
	_go_fuzz_dep_.CoverTab[2974]++
											preLen := len(msg)
											if msg, err = r.pack(msg, b.compression, b.start); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1434
		_go_fuzz_dep_.CoverTab[2982]++
												return &nestedError{"PTRResource body", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1435
		// _ = "end of CoverTab[2982]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1436
		_go_fuzz_dep_.CoverTab[2983]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1436
		// _ = "end of CoverTab[2983]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1436
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1436
	// _ = "end of CoverTab[2974]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1436
	_go_fuzz_dep_.CoverTab[2975]++
											if err := h.fixLen(msg, lenOff, preLen); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1437
		_go_fuzz_dep_.CoverTab[2984]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1438
		// _ = "end of CoverTab[2984]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1439
		_go_fuzz_dep_.CoverTab[2985]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1439
		// _ = "end of CoverTab[2985]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1439
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1439
	// _ = "end of CoverTab[2975]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1439
	_go_fuzz_dep_.CoverTab[2976]++
											if err := b.incrementSectionCount(); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1440
		_go_fuzz_dep_.CoverTab[2986]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1441
		// _ = "end of CoverTab[2986]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1442
		_go_fuzz_dep_.CoverTab[2987]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1442
		// _ = "end of CoverTab[2987]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1442
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1442
	// _ = "end of CoverTab[2976]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1442
	_go_fuzz_dep_.CoverTab[2977]++
											b.msg = msg
											return nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1444
	// _ = "end of CoverTab[2977]"
}

// SOAResource adds a single SOAResource.
func (b *Builder) SOAResource(h ResourceHeader, r SOAResource) error {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1448
	_go_fuzz_dep_.CoverTab[2988]++
											if err := b.checkResourceSection(); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1449
		_go_fuzz_dep_.CoverTab[2994]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1450
		// _ = "end of CoverTab[2994]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1451
		_go_fuzz_dep_.CoverTab[2995]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1451
		// _ = "end of CoverTab[2995]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1451
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1451
	// _ = "end of CoverTab[2988]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1451
	_go_fuzz_dep_.CoverTab[2989]++
											h.Type = r.realType()
											msg, lenOff, err := h.pack(b.msg, b.compression, b.start)
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1454
		_go_fuzz_dep_.CoverTab[2996]++
												return &nestedError{"ResourceHeader", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1455
		// _ = "end of CoverTab[2996]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1456
		_go_fuzz_dep_.CoverTab[2997]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1456
		// _ = "end of CoverTab[2997]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1456
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1456
	// _ = "end of CoverTab[2989]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1456
	_go_fuzz_dep_.CoverTab[2990]++
											preLen := len(msg)
											if msg, err = r.pack(msg, b.compression, b.start); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1458
		_go_fuzz_dep_.CoverTab[2998]++
												return &nestedError{"SOAResource body", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1459
		// _ = "end of CoverTab[2998]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1460
		_go_fuzz_dep_.CoverTab[2999]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1460
		// _ = "end of CoverTab[2999]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1460
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1460
	// _ = "end of CoverTab[2990]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1460
	_go_fuzz_dep_.CoverTab[2991]++
											if err := h.fixLen(msg, lenOff, preLen); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1461
		_go_fuzz_dep_.CoverTab[3000]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1462
		// _ = "end of CoverTab[3000]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1463
		_go_fuzz_dep_.CoverTab[3001]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1463
		// _ = "end of CoverTab[3001]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1463
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1463
	// _ = "end of CoverTab[2991]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1463
	_go_fuzz_dep_.CoverTab[2992]++
											if err := b.incrementSectionCount(); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1464
		_go_fuzz_dep_.CoverTab[3002]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1465
		// _ = "end of CoverTab[3002]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1466
		_go_fuzz_dep_.CoverTab[3003]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1466
		// _ = "end of CoverTab[3003]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1466
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1466
	// _ = "end of CoverTab[2992]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1466
	_go_fuzz_dep_.CoverTab[2993]++
											b.msg = msg
											return nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1468
	// _ = "end of CoverTab[2993]"
}

// TXTResource adds a single TXTResource.
func (b *Builder) TXTResource(h ResourceHeader, r TXTResource) error {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1472
	_go_fuzz_dep_.CoverTab[3004]++
											if err := b.checkResourceSection(); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1473
		_go_fuzz_dep_.CoverTab[3010]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1474
		// _ = "end of CoverTab[3010]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1475
		_go_fuzz_dep_.CoverTab[3011]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1475
		// _ = "end of CoverTab[3011]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1475
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1475
	// _ = "end of CoverTab[3004]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1475
	_go_fuzz_dep_.CoverTab[3005]++
											h.Type = r.realType()
											msg, lenOff, err := h.pack(b.msg, b.compression, b.start)
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1478
		_go_fuzz_dep_.CoverTab[3012]++
												return &nestedError{"ResourceHeader", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1479
		// _ = "end of CoverTab[3012]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1480
		_go_fuzz_dep_.CoverTab[3013]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1480
		// _ = "end of CoverTab[3013]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1480
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1480
	// _ = "end of CoverTab[3005]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1480
	_go_fuzz_dep_.CoverTab[3006]++
											preLen := len(msg)
											if msg, err = r.pack(msg, b.compression, b.start); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1482
		_go_fuzz_dep_.CoverTab[3014]++
												return &nestedError{"TXTResource body", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1483
		// _ = "end of CoverTab[3014]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1484
		_go_fuzz_dep_.CoverTab[3015]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1484
		// _ = "end of CoverTab[3015]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1484
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1484
	// _ = "end of CoverTab[3006]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1484
	_go_fuzz_dep_.CoverTab[3007]++
											if err := h.fixLen(msg, lenOff, preLen); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1485
		_go_fuzz_dep_.CoverTab[3016]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1486
		// _ = "end of CoverTab[3016]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1487
		_go_fuzz_dep_.CoverTab[3017]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1487
		// _ = "end of CoverTab[3017]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1487
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1487
	// _ = "end of CoverTab[3007]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1487
	_go_fuzz_dep_.CoverTab[3008]++
											if err := b.incrementSectionCount(); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1488
		_go_fuzz_dep_.CoverTab[3018]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1489
		// _ = "end of CoverTab[3018]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1490
		_go_fuzz_dep_.CoverTab[3019]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1490
		// _ = "end of CoverTab[3019]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1490
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1490
	// _ = "end of CoverTab[3008]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1490
	_go_fuzz_dep_.CoverTab[3009]++
											b.msg = msg
											return nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1492
	// _ = "end of CoverTab[3009]"
}

// SRVResource adds a single SRVResource.
func (b *Builder) SRVResource(h ResourceHeader, r SRVResource) error {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1496
	_go_fuzz_dep_.CoverTab[3020]++
											if err := b.checkResourceSection(); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1497
		_go_fuzz_dep_.CoverTab[3026]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1498
		// _ = "end of CoverTab[3026]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1499
		_go_fuzz_dep_.CoverTab[3027]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1499
		// _ = "end of CoverTab[3027]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1499
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1499
	// _ = "end of CoverTab[3020]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1499
	_go_fuzz_dep_.CoverTab[3021]++
											h.Type = r.realType()
											msg, lenOff, err := h.pack(b.msg, b.compression, b.start)
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1502
		_go_fuzz_dep_.CoverTab[3028]++
												return &nestedError{"ResourceHeader", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1503
		// _ = "end of CoverTab[3028]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1504
		_go_fuzz_dep_.CoverTab[3029]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1504
		// _ = "end of CoverTab[3029]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1504
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1504
	// _ = "end of CoverTab[3021]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1504
	_go_fuzz_dep_.CoverTab[3022]++
											preLen := len(msg)
											if msg, err = r.pack(msg, b.compression, b.start); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1506
		_go_fuzz_dep_.CoverTab[3030]++
												return &nestedError{"SRVResource body", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1507
		// _ = "end of CoverTab[3030]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1508
		_go_fuzz_dep_.CoverTab[3031]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1508
		// _ = "end of CoverTab[3031]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1508
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1508
	// _ = "end of CoverTab[3022]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1508
	_go_fuzz_dep_.CoverTab[3023]++
											if err := h.fixLen(msg, lenOff, preLen); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1509
		_go_fuzz_dep_.CoverTab[3032]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1510
		// _ = "end of CoverTab[3032]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1511
		_go_fuzz_dep_.CoverTab[3033]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1511
		// _ = "end of CoverTab[3033]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1511
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1511
	// _ = "end of CoverTab[3023]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1511
	_go_fuzz_dep_.CoverTab[3024]++
											if err := b.incrementSectionCount(); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1512
		_go_fuzz_dep_.CoverTab[3034]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1513
		// _ = "end of CoverTab[3034]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1514
		_go_fuzz_dep_.CoverTab[3035]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1514
		// _ = "end of CoverTab[3035]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1514
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1514
	// _ = "end of CoverTab[3024]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1514
	_go_fuzz_dep_.CoverTab[3025]++
											b.msg = msg
											return nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1516
	// _ = "end of CoverTab[3025]"
}

// AResource adds a single AResource.
func (b *Builder) AResource(h ResourceHeader, r AResource) error {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1520
	_go_fuzz_dep_.CoverTab[3036]++
											if err := b.checkResourceSection(); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1521
		_go_fuzz_dep_.CoverTab[3042]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1522
		// _ = "end of CoverTab[3042]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1523
		_go_fuzz_dep_.CoverTab[3043]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1523
		// _ = "end of CoverTab[3043]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1523
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1523
	// _ = "end of CoverTab[3036]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1523
	_go_fuzz_dep_.CoverTab[3037]++
											h.Type = r.realType()
											msg, lenOff, err := h.pack(b.msg, b.compression, b.start)
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1526
		_go_fuzz_dep_.CoverTab[3044]++
												return &nestedError{"ResourceHeader", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1527
		// _ = "end of CoverTab[3044]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1528
		_go_fuzz_dep_.CoverTab[3045]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1528
		// _ = "end of CoverTab[3045]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1528
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1528
	// _ = "end of CoverTab[3037]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1528
	_go_fuzz_dep_.CoverTab[3038]++
											preLen := len(msg)
											if msg, err = r.pack(msg, b.compression, b.start); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1530
		_go_fuzz_dep_.CoverTab[3046]++
												return &nestedError{"AResource body", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1531
		// _ = "end of CoverTab[3046]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1532
		_go_fuzz_dep_.CoverTab[3047]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1532
		// _ = "end of CoverTab[3047]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1532
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1532
	// _ = "end of CoverTab[3038]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1532
	_go_fuzz_dep_.CoverTab[3039]++
											if err := h.fixLen(msg, lenOff, preLen); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1533
		_go_fuzz_dep_.CoverTab[3048]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1534
		// _ = "end of CoverTab[3048]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1535
		_go_fuzz_dep_.CoverTab[3049]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1535
		// _ = "end of CoverTab[3049]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1535
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1535
	// _ = "end of CoverTab[3039]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1535
	_go_fuzz_dep_.CoverTab[3040]++
											if err := b.incrementSectionCount(); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1536
		_go_fuzz_dep_.CoverTab[3050]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1537
		// _ = "end of CoverTab[3050]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1538
		_go_fuzz_dep_.CoverTab[3051]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1538
		// _ = "end of CoverTab[3051]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1538
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1538
	// _ = "end of CoverTab[3040]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1538
	_go_fuzz_dep_.CoverTab[3041]++
											b.msg = msg
											return nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1540
	// _ = "end of CoverTab[3041]"
}

// AAAAResource adds a single AAAAResource.
func (b *Builder) AAAAResource(h ResourceHeader, r AAAAResource) error {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1544
	_go_fuzz_dep_.CoverTab[3052]++
											if err := b.checkResourceSection(); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1545
		_go_fuzz_dep_.CoverTab[3058]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1546
		// _ = "end of CoverTab[3058]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1547
		_go_fuzz_dep_.CoverTab[3059]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1547
		// _ = "end of CoverTab[3059]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1547
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1547
	// _ = "end of CoverTab[3052]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1547
	_go_fuzz_dep_.CoverTab[3053]++
											h.Type = r.realType()
											msg, lenOff, err := h.pack(b.msg, b.compression, b.start)
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1550
		_go_fuzz_dep_.CoverTab[3060]++
												return &nestedError{"ResourceHeader", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1551
		// _ = "end of CoverTab[3060]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1552
		_go_fuzz_dep_.CoverTab[3061]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1552
		// _ = "end of CoverTab[3061]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1552
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1552
	// _ = "end of CoverTab[3053]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1552
	_go_fuzz_dep_.CoverTab[3054]++
											preLen := len(msg)
											if msg, err = r.pack(msg, b.compression, b.start); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1554
		_go_fuzz_dep_.CoverTab[3062]++
												return &nestedError{"AAAAResource body", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1555
		// _ = "end of CoverTab[3062]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1556
		_go_fuzz_dep_.CoverTab[3063]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1556
		// _ = "end of CoverTab[3063]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1556
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1556
	// _ = "end of CoverTab[3054]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1556
	_go_fuzz_dep_.CoverTab[3055]++
											if err := h.fixLen(msg, lenOff, preLen); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1557
		_go_fuzz_dep_.CoverTab[3064]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1558
		// _ = "end of CoverTab[3064]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1559
		_go_fuzz_dep_.CoverTab[3065]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1559
		// _ = "end of CoverTab[3065]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1559
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1559
	// _ = "end of CoverTab[3055]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1559
	_go_fuzz_dep_.CoverTab[3056]++
											if err := b.incrementSectionCount(); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1560
		_go_fuzz_dep_.CoverTab[3066]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1561
		// _ = "end of CoverTab[3066]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1562
		_go_fuzz_dep_.CoverTab[3067]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1562
		// _ = "end of CoverTab[3067]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1562
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1562
	// _ = "end of CoverTab[3056]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1562
	_go_fuzz_dep_.CoverTab[3057]++
											b.msg = msg
											return nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1564
	// _ = "end of CoverTab[3057]"
}

// OPTResource adds a single OPTResource.
func (b *Builder) OPTResource(h ResourceHeader, r OPTResource) error {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1568
	_go_fuzz_dep_.CoverTab[3068]++
											if err := b.checkResourceSection(); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1569
		_go_fuzz_dep_.CoverTab[3074]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1570
		// _ = "end of CoverTab[3074]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1571
		_go_fuzz_dep_.CoverTab[3075]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1571
		// _ = "end of CoverTab[3075]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1571
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1571
	// _ = "end of CoverTab[3068]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1571
	_go_fuzz_dep_.CoverTab[3069]++
											h.Type = r.realType()
											msg, lenOff, err := h.pack(b.msg, b.compression, b.start)
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1574
		_go_fuzz_dep_.CoverTab[3076]++
												return &nestedError{"ResourceHeader", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1575
		// _ = "end of CoverTab[3076]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1576
		_go_fuzz_dep_.CoverTab[3077]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1576
		// _ = "end of CoverTab[3077]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1576
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1576
	// _ = "end of CoverTab[3069]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1576
	_go_fuzz_dep_.CoverTab[3070]++
											preLen := len(msg)
											if msg, err = r.pack(msg, b.compression, b.start); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1578
		_go_fuzz_dep_.CoverTab[3078]++
												return &nestedError{"OPTResource body", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1579
		// _ = "end of CoverTab[3078]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1580
		_go_fuzz_dep_.CoverTab[3079]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1580
		// _ = "end of CoverTab[3079]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1580
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1580
	// _ = "end of CoverTab[3070]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1580
	_go_fuzz_dep_.CoverTab[3071]++
											if err := h.fixLen(msg, lenOff, preLen); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1581
		_go_fuzz_dep_.CoverTab[3080]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1582
		// _ = "end of CoverTab[3080]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1583
		_go_fuzz_dep_.CoverTab[3081]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1583
		// _ = "end of CoverTab[3081]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1583
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1583
	// _ = "end of CoverTab[3071]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1583
	_go_fuzz_dep_.CoverTab[3072]++
											if err := b.incrementSectionCount(); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1584
		_go_fuzz_dep_.CoverTab[3082]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1585
		// _ = "end of CoverTab[3082]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1586
		_go_fuzz_dep_.CoverTab[3083]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1586
		// _ = "end of CoverTab[3083]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1586
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1586
	// _ = "end of CoverTab[3072]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1586
	_go_fuzz_dep_.CoverTab[3073]++
											b.msg = msg
											return nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1588
	// _ = "end of CoverTab[3073]"
}

// UnknownResource adds a single UnknownResource.
func (b *Builder) UnknownResource(h ResourceHeader, r UnknownResource) error {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1592
	_go_fuzz_dep_.CoverTab[3084]++
											if err := b.checkResourceSection(); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1593
		_go_fuzz_dep_.CoverTab[3090]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1594
		// _ = "end of CoverTab[3090]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1595
		_go_fuzz_dep_.CoverTab[3091]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1595
		// _ = "end of CoverTab[3091]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1595
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1595
	// _ = "end of CoverTab[3084]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1595
	_go_fuzz_dep_.CoverTab[3085]++
											h.Type = r.realType()
											msg, lenOff, err := h.pack(b.msg, b.compression, b.start)
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1598
		_go_fuzz_dep_.CoverTab[3092]++
												return &nestedError{"ResourceHeader", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1599
		// _ = "end of CoverTab[3092]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1600
		_go_fuzz_dep_.CoverTab[3093]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1600
		// _ = "end of CoverTab[3093]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1600
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1600
	// _ = "end of CoverTab[3085]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1600
	_go_fuzz_dep_.CoverTab[3086]++
											preLen := len(msg)
											if msg, err = r.pack(msg, b.compression, b.start); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1602
		_go_fuzz_dep_.CoverTab[3094]++
												return &nestedError{"UnknownResource body", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1603
		// _ = "end of CoverTab[3094]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1604
		_go_fuzz_dep_.CoverTab[3095]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1604
		// _ = "end of CoverTab[3095]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1604
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1604
	// _ = "end of CoverTab[3086]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1604
	_go_fuzz_dep_.CoverTab[3087]++
											if err := h.fixLen(msg, lenOff, preLen); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1605
		_go_fuzz_dep_.CoverTab[3096]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1606
		// _ = "end of CoverTab[3096]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1607
		_go_fuzz_dep_.CoverTab[3097]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1607
		// _ = "end of CoverTab[3097]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1607
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1607
	// _ = "end of CoverTab[3087]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1607
	_go_fuzz_dep_.CoverTab[3088]++
											if err := b.incrementSectionCount(); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1608
		_go_fuzz_dep_.CoverTab[3098]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1609
		// _ = "end of CoverTab[3098]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1610
		_go_fuzz_dep_.CoverTab[3099]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1610
		// _ = "end of CoverTab[3099]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1610
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1610
	// _ = "end of CoverTab[3088]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1610
	_go_fuzz_dep_.CoverTab[3089]++
											b.msg = msg
											return nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1612
	// _ = "end of CoverTab[3089]"
}

// Finish ends message building and generates a binary message.
func (b *Builder) Finish() ([]byte, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1616
	_go_fuzz_dep_.CoverTab[3100]++
											if b.section < sectionHeader {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1617
		_go_fuzz_dep_.CoverTab[3102]++
												return nil, ErrNotStarted
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1618
		// _ = "end of CoverTab[3102]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1619
		_go_fuzz_dep_.CoverTab[3103]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1619
		// _ = "end of CoverTab[3103]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1619
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1619
	// _ = "end of CoverTab[3100]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1619
	_go_fuzz_dep_.CoverTab[3101]++
											b.section = sectionDone

											b.header.pack(b.msg[b.start:b.start])
											return b.msg, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1623
	// _ = "end of CoverTab[3101]"
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
	_go_fuzz_dep_.CoverTab[3104]++
											return "dnsmessage.ResourceHeader{" +
		"Name: " + h.Name.GoString() + ", " +
		"Type: " + h.Type.GoString() + ", " +
		"Class: " + h.Class.GoString() + ", " +
		"TTL: " + printUint32(h.TTL) + ", " +
		"Length: " + printUint16(h.Length) + "}"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1659
	// _ = "end of CoverTab[3104]"
}

// pack appends the wire format of the ResourceHeader to oldMsg.
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1662
//
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1662
// lenOff is the offset in msg where the Length field was packed.
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1665
func (h *ResourceHeader) pack(oldMsg []byte, compression map[string]int, compressionOff int) (msg []byte, lenOff int, err error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1665
	_go_fuzz_dep_.CoverTab[3105]++
											msg = oldMsg
											if msg, err = h.Name.pack(msg, compression, compressionOff); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1667
		_go_fuzz_dep_.CoverTab[3107]++
												return oldMsg, 0, &nestedError{"Name", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1668
		// _ = "end of CoverTab[3107]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1669
		_go_fuzz_dep_.CoverTab[3108]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1669
		// _ = "end of CoverTab[3108]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1669
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1669
	// _ = "end of CoverTab[3105]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1669
	_go_fuzz_dep_.CoverTab[3106]++
											msg = packType(msg, h.Type)
											msg = packClass(msg, h.Class)
											msg = packUint32(msg, h.TTL)
											lenOff = len(msg)
											msg = packUint16(msg, h.Length)
											return msg, lenOff, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1675
	// _ = "end of CoverTab[3106]"
}

func (h *ResourceHeader) unpack(msg []byte, off int) (int, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1678
	_go_fuzz_dep_.CoverTab[3109]++
											newOff := off
											var err error
											if newOff, err = h.Name.unpack(msg, newOff); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1681
		_go_fuzz_dep_.CoverTab[3115]++
												return off, &nestedError{"Name", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1682
		// _ = "end of CoverTab[3115]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1683
		_go_fuzz_dep_.CoverTab[3116]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1683
		// _ = "end of CoverTab[3116]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1683
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1683
	// _ = "end of CoverTab[3109]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1683
	_go_fuzz_dep_.CoverTab[3110]++
											if h.Type, newOff, err = unpackType(msg, newOff); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1684
		_go_fuzz_dep_.CoverTab[3117]++
												return off, &nestedError{"Type", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1685
		// _ = "end of CoverTab[3117]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1686
		_go_fuzz_dep_.CoverTab[3118]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1686
		// _ = "end of CoverTab[3118]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1686
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1686
	// _ = "end of CoverTab[3110]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1686
	_go_fuzz_dep_.CoverTab[3111]++
											if h.Class, newOff, err = unpackClass(msg, newOff); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1687
		_go_fuzz_dep_.CoverTab[3119]++
												return off, &nestedError{"Class", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1688
		// _ = "end of CoverTab[3119]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1689
		_go_fuzz_dep_.CoverTab[3120]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1689
		// _ = "end of CoverTab[3120]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1689
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1689
	// _ = "end of CoverTab[3111]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1689
	_go_fuzz_dep_.CoverTab[3112]++
											if h.TTL, newOff, err = unpackUint32(msg, newOff); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1690
		_go_fuzz_dep_.CoverTab[3121]++
												return off, &nestedError{"TTL", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1691
		// _ = "end of CoverTab[3121]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1692
		_go_fuzz_dep_.CoverTab[3122]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1692
		// _ = "end of CoverTab[3122]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1692
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1692
	// _ = "end of CoverTab[3112]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1692
	_go_fuzz_dep_.CoverTab[3113]++
											if h.Length, newOff, err = unpackUint16(msg, newOff); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1693
		_go_fuzz_dep_.CoverTab[3123]++
												return off, &nestedError{"Length", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1694
		// _ = "end of CoverTab[3123]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1695
		_go_fuzz_dep_.CoverTab[3124]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1695
		// _ = "end of CoverTab[3124]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1695
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1695
	// _ = "end of CoverTab[3113]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1695
	_go_fuzz_dep_.CoverTab[3114]++
											return newOff, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1696
	// _ = "end of CoverTab[3114]"
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
	_go_fuzz_dep_.CoverTab[3125]++
											conLen := len(msg) - preLen
											if conLen > int(^uint16(0)) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1707
		_go_fuzz_dep_.CoverTab[3127]++
												return errResTooLong
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1708
		// _ = "end of CoverTab[3127]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1709
		_go_fuzz_dep_.CoverTab[3128]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1709
		// _ = "end of CoverTab[3128]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1709
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1709
	// _ = "end of CoverTab[3125]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1709
	_go_fuzz_dep_.CoverTab[3126]++

//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1712
	packUint16(msg[lenOff:lenOff], uint16(conLen))
											h.Length = uint16(conLen)

											return nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1715
	// _ = "end of CoverTab[3126]"
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
	_go_fuzz_dep_.CoverTab[3129]++
											h.Name = Name{Data: [nameLen]byte{'.'}, Length: 1}
											h.Type = TypeOPT
											h.Class = Class(udpPayloadLen)
											h.TTL = uint32(extRCode) >> 4 << 24
											if dnssecOK {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1735
		_go_fuzz_dep_.CoverTab[3131]++
												h.TTL |= edns0DNSSECOK
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1736
		// _ = "end of CoverTab[3131]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1737
		_go_fuzz_dep_.CoverTab[3132]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1737
		// _ = "end of CoverTab[3132]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1737
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1737
	// _ = "end of CoverTab[3129]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1737
	_go_fuzz_dep_.CoverTab[3130]++
											return nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1738
	// _ = "end of CoverTab[3130]"
}

// DNSSECAllowed reports whether the DNSSEC OK bit is set.
func (h *ResourceHeader) DNSSECAllowed() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1742
	_go_fuzz_dep_.CoverTab[3133]++
											return h.TTL&edns0DNSSECOKMask == edns0DNSSECOK
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1743
	// _ = "end of CoverTab[3133]"
}

// ExtendedRCode returns an extended RCode.
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1746
//
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1746
// The provided rcode must be the RCode in DNS message header.
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1749
func (h *ResourceHeader) ExtendedRCode(rcode RCode) RCode {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1749
	_go_fuzz_dep_.CoverTab[3134]++
											if h.TTL&ednsVersionMask == edns0Version {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1750
		_go_fuzz_dep_.CoverTab[3136]++
												return RCode(h.TTL>>24<<4) | rcode
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1751
		// _ = "end of CoverTab[3136]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1752
		_go_fuzz_dep_.CoverTab[3137]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1752
		// _ = "end of CoverTab[3137]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1752
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1752
	// _ = "end of CoverTab[3134]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1752
	_go_fuzz_dep_.CoverTab[3135]++
											return rcode
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1753
	// _ = "end of CoverTab[3135]"
}

func skipResource(msg []byte, off int) (int, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1756
	_go_fuzz_dep_.CoverTab[3138]++
											newOff, err := skipName(msg, off)
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1758
		_go_fuzz_dep_.CoverTab[3145]++
												return off, &nestedError{"Name", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1759
		// _ = "end of CoverTab[3145]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1760
		_go_fuzz_dep_.CoverTab[3146]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1760
		// _ = "end of CoverTab[3146]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1760
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1760
	// _ = "end of CoverTab[3138]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1760
	_go_fuzz_dep_.CoverTab[3139]++
											if newOff, err = skipType(msg, newOff); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1761
		_go_fuzz_dep_.CoverTab[3147]++
												return off, &nestedError{"Type", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1762
		// _ = "end of CoverTab[3147]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1763
		_go_fuzz_dep_.CoverTab[3148]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1763
		// _ = "end of CoverTab[3148]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1763
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1763
	// _ = "end of CoverTab[3139]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1763
	_go_fuzz_dep_.CoverTab[3140]++
											if newOff, err = skipClass(msg, newOff); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1764
		_go_fuzz_dep_.CoverTab[3149]++
												return off, &nestedError{"Class", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1765
		// _ = "end of CoverTab[3149]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1766
		_go_fuzz_dep_.CoverTab[3150]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1766
		// _ = "end of CoverTab[3150]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1766
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1766
	// _ = "end of CoverTab[3140]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1766
	_go_fuzz_dep_.CoverTab[3141]++
											if newOff, err = skipUint32(msg, newOff); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1767
		_go_fuzz_dep_.CoverTab[3151]++
												return off, &nestedError{"TTL", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1768
		// _ = "end of CoverTab[3151]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1769
		_go_fuzz_dep_.CoverTab[3152]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1769
		// _ = "end of CoverTab[3152]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1769
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1769
	// _ = "end of CoverTab[3141]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1769
	_go_fuzz_dep_.CoverTab[3142]++
											length, newOff, err := unpackUint16(msg, newOff)
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1771
		_go_fuzz_dep_.CoverTab[3153]++
												return off, &nestedError{"Length", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1772
		// _ = "end of CoverTab[3153]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1773
		_go_fuzz_dep_.CoverTab[3154]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1773
		// _ = "end of CoverTab[3154]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1773
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1773
	// _ = "end of CoverTab[3142]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1773
	_go_fuzz_dep_.CoverTab[3143]++
											if newOff += int(length); newOff > len(msg) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1774
		_go_fuzz_dep_.CoverTab[3155]++
												return off, errResourceLen
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1775
		// _ = "end of CoverTab[3155]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1776
		_go_fuzz_dep_.CoverTab[3156]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1776
		// _ = "end of CoverTab[3156]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1776
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1776
	// _ = "end of CoverTab[3143]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1776
	_go_fuzz_dep_.CoverTab[3144]++
											return newOff, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1777
	// _ = "end of CoverTab[3144]"
}

// packUint16 appends the wire format of field to msg.
func packUint16(msg []byte, field uint16) []byte {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1781
	_go_fuzz_dep_.CoverTab[3157]++
											return append(msg, byte(field>>8), byte(field))
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1782
	// _ = "end of CoverTab[3157]"
}

func unpackUint16(msg []byte, off int) (uint16, int, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1785
	_go_fuzz_dep_.CoverTab[3158]++
											if off+uint16Len > len(msg) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1786
		_go_fuzz_dep_.CoverTab[3160]++
												return 0, off, errBaseLen
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1787
		// _ = "end of CoverTab[3160]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1788
		_go_fuzz_dep_.CoverTab[3161]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1788
		// _ = "end of CoverTab[3161]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1788
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1788
	// _ = "end of CoverTab[3158]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1788
	_go_fuzz_dep_.CoverTab[3159]++
											return uint16(msg[off])<<8 | uint16(msg[off+1]), off + uint16Len, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1789
	// _ = "end of CoverTab[3159]"
}

func skipUint16(msg []byte, off int) (int, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1792
	_go_fuzz_dep_.CoverTab[3162]++
											if off+uint16Len > len(msg) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1793
		_go_fuzz_dep_.CoverTab[3164]++
												return off, errBaseLen
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1794
		// _ = "end of CoverTab[3164]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1795
		_go_fuzz_dep_.CoverTab[3165]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1795
		// _ = "end of CoverTab[3165]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1795
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1795
	// _ = "end of CoverTab[3162]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1795
	_go_fuzz_dep_.CoverTab[3163]++
											return off + uint16Len, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1796
	// _ = "end of CoverTab[3163]"
}

// packType appends the wire format of field to msg.
func packType(msg []byte, field Type) []byte {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1800
	_go_fuzz_dep_.CoverTab[3166]++
											return packUint16(msg, uint16(field))
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1801
	// _ = "end of CoverTab[3166]"
}

func unpackType(msg []byte, off int) (Type, int, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1804
	_go_fuzz_dep_.CoverTab[3167]++
											t, o, err := unpackUint16(msg, off)
											return Type(t), o, err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1806
	// _ = "end of CoverTab[3167]"
}

func skipType(msg []byte, off int) (int, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1809
	_go_fuzz_dep_.CoverTab[3168]++
											return skipUint16(msg, off)
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1810
	// _ = "end of CoverTab[3168]"
}

// packClass appends the wire format of field to msg.
func packClass(msg []byte, field Class) []byte {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1814
	_go_fuzz_dep_.CoverTab[3169]++
											return packUint16(msg, uint16(field))
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1815
	// _ = "end of CoverTab[3169]"
}

func unpackClass(msg []byte, off int) (Class, int, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1818
	_go_fuzz_dep_.CoverTab[3170]++
											c, o, err := unpackUint16(msg, off)
											return Class(c), o, err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1820
	// _ = "end of CoverTab[3170]"
}

func skipClass(msg []byte, off int) (int, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1823
	_go_fuzz_dep_.CoverTab[3171]++
											return skipUint16(msg, off)
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1824
	// _ = "end of CoverTab[3171]"
}

// packUint32 appends the wire format of field to msg.
func packUint32(msg []byte, field uint32) []byte {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1828
	_go_fuzz_dep_.CoverTab[3172]++
											return append(
		msg,
		byte(field>>24),
		byte(field>>16),
		byte(field>>8),
		byte(field),
	)
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1835
	// _ = "end of CoverTab[3172]"
}

func unpackUint32(msg []byte, off int) (uint32, int, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1838
	_go_fuzz_dep_.CoverTab[3173]++
											if off+uint32Len > len(msg) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1839
		_go_fuzz_dep_.CoverTab[3175]++
												return 0, off, errBaseLen
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1840
		// _ = "end of CoverTab[3175]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1841
		_go_fuzz_dep_.CoverTab[3176]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1841
		// _ = "end of CoverTab[3176]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1841
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1841
	// _ = "end of CoverTab[3173]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1841
	_go_fuzz_dep_.CoverTab[3174]++
											v := uint32(msg[off])<<24 | uint32(msg[off+1])<<16 | uint32(msg[off+2])<<8 | uint32(msg[off+3])
											return v, off + uint32Len, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1843
	// _ = "end of CoverTab[3174]"
}

func skipUint32(msg []byte, off int) (int, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1846
	_go_fuzz_dep_.CoverTab[3177]++
											if off+uint32Len > len(msg) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1847
		_go_fuzz_dep_.CoverTab[3179]++
												return off, errBaseLen
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1848
		// _ = "end of CoverTab[3179]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1849
		_go_fuzz_dep_.CoverTab[3180]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1849
		// _ = "end of CoverTab[3180]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1849
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1849
	// _ = "end of CoverTab[3177]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1849
	_go_fuzz_dep_.CoverTab[3178]++
											return off + uint32Len, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1850
	// _ = "end of CoverTab[3178]"
}

// packText appends the wire format of field to msg.
func packText(msg []byte, field string) ([]byte, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1854
	_go_fuzz_dep_.CoverTab[3181]++
											l := len(field)
											if l > 255 {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1856
		_go_fuzz_dep_.CoverTab[3183]++
												return nil, errStringTooLong
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1857
		// _ = "end of CoverTab[3183]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1858
		_go_fuzz_dep_.CoverTab[3184]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1858
		// _ = "end of CoverTab[3184]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1858
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1858
	// _ = "end of CoverTab[3181]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1858
	_go_fuzz_dep_.CoverTab[3182]++
											msg = append(msg, byte(l))
											msg = append(msg, field...)

											return msg, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1862
	// _ = "end of CoverTab[3182]"
}

func unpackText(msg []byte, off int) (string, int, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1865
	_go_fuzz_dep_.CoverTab[3185]++
											if off >= len(msg) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1866
		_go_fuzz_dep_.CoverTab[3188]++
												return "", off, errBaseLen
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1867
		// _ = "end of CoverTab[3188]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1868
		_go_fuzz_dep_.CoverTab[3189]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1868
		// _ = "end of CoverTab[3189]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1868
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1868
	// _ = "end of CoverTab[3185]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1868
	_go_fuzz_dep_.CoverTab[3186]++
											beginOff := off + 1
											endOff := beginOff + int(msg[off])
											if endOff > len(msg) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1871
		_go_fuzz_dep_.CoverTab[3190]++
												return "", off, errCalcLen
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1872
		// _ = "end of CoverTab[3190]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1873
		_go_fuzz_dep_.CoverTab[3191]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1873
		// _ = "end of CoverTab[3191]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1873
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1873
	// _ = "end of CoverTab[3186]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1873
	_go_fuzz_dep_.CoverTab[3187]++
											return string(msg[beginOff:endOff]), endOff, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1874
	// _ = "end of CoverTab[3187]"
}

// packBytes appends the wire format of field to msg.
func packBytes(msg []byte, field []byte) []byte {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1878
	_go_fuzz_dep_.CoverTab[3192]++
											return append(msg, field...)
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1879
	// _ = "end of CoverTab[3192]"
}

func unpackBytes(msg []byte, off int, field []byte) (int, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1882
	_go_fuzz_dep_.CoverTab[3193]++
											newOff := off + len(field)
											if newOff > len(msg) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1884
		_go_fuzz_dep_.CoverTab[3195]++
												return off, errBaseLen
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1885
		// _ = "end of CoverTab[3195]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1886
		_go_fuzz_dep_.CoverTab[3196]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1886
		// _ = "end of CoverTab[3196]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1886
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1886
	// _ = "end of CoverTab[3193]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1886
	_go_fuzz_dep_.CoverTab[3194]++
											copy(field, msg[off:newOff])
											return newOff, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1888
	// _ = "end of CoverTab[3194]"
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
	_go_fuzz_dep_.CoverTab[3197]++
											if len(name) > nameLen {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1902
		_go_fuzz_dep_.CoverTab[3199]++
												return Name{}, errCalcLen
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1903
		// _ = "end of CoverTab[3199]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1904
		_go_fuzz_dep_.CoverTab[3200]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1904
		// _ = "end of CoverTab[3200]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1904
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1904
	// _ = "end of CoverTab[3197]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1904
	_go_fuzz_dep_.CoverTab[3198]++
											n := Name{Length: uint8(len(name))}
											copy(n.Data[:], name)
											return n, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1907
	// _ = "end of CoverTab[3198]"
}

// MustNewName creates a new Name from a string and panics on error.
func MustNewName(name string) Name {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1911
	_go_fuzz_dep_.CoverTab[3201]++
											n, err := NewName(name)
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1913
		_go_fuzz_dep_.CoverTab[3203]++
												panic("creating name: " + err.Error())
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1914
		// _ = "end of CoverTab[3203]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1915
		_go_fuzz_dep_.CoverTab[3204]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1915
		// _ = "end of CoverTab[3204]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1915
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1915
	// _ = "end of CoverTab[3201]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1915
	_go_fuzz_dep_.CoverTab[3202]++
											return n
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1916
	// _ = "end of CoverTab[3202]"
}

// String implements fmt.Stringer.String.
func (n Name) String() string {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1920
	_go_fuzz_dep_.CoverTab[3205]++
											return string(n.Data[:n.Length])
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1921
	// _ = "end of CoverTab[3205]"
}

// GoString implements fmt.GoStringer.GoString.
func (n *Name) GoString() string {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1925
	_go_fuzz_dep_.CoverTab[3206]++
											return `dnsmessage.MustNewName("` + printString(n.Data[:n.Length]) + `")`
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1926
	// _ = "end of CoverTab[3206]"
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
	_go_fuzz_dep_.CoverTab[3207]++
											oldMsg := msg

//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1940
	if n.Length == 0 || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1940
		_go_fuzz_dep_.CoverTab[3211]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1940
		return n.Data[n.Length-1] != '.'
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1940
		// _ = "end of CoverTab[3211]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1940
	}() {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1940
		_go_fuzz_dep_.CoverTab[3212]++
												return oldMsg, errNonCanonicalName
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1941
		// _ = "end of CoverTab[3212]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1942
		_go_fuzz_dep_.CoverTab[3213]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1942
		// _ = "end of CoverTab[3213]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1942
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1942
	// _ = "end of CoverTab[3207]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1942
	_go_fuzz_dep_.CoverTab[3208]++

//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1945
	if n.Data[0] == '.' && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1945
		_go_fuzz_dep_.CoverTab[3214]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1945
		return n.Length == 1
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1945
		// _ = "end of CoverTab[3214]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1945
	}() {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1945
		_go_fuzz_dep_.CoverTab[3215]++
												return append(msg, 0), nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1946
		// _ = "end of CoverTab[3215]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1947
		_go_fuzz_dep_.CoverTab[3216]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1947
		// _ = "end of CoverTab[3216]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1947
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1947
	// _ = "end of CoverTab[3208]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1947
	_go_fuzz_dep_.CoverTab[3209]++

//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1950
	for i, begin := 0, 0; i < int(n.Length); i++ {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1950
		_go_fuzz_dep_.CoverTab[3217]++

												if n.Data[i] == '.' {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1952
			_go_fuzz_dep_.CoverTab[3219]++

//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1956
			if i-begin >= 1<<6 {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1956
				_go_fuzz_dep_.CoverTab[3223]++
														return oldMsg, errSegTooLong
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1957
				// _ = "end of CoverTab[3223]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1958
				_go_fuzz_dep_.CoverTab[3224]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1958
				// _ = "end of CoverTab[3224]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1958
			}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1958
			// _ = "end of CoverTab[3219]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1958
			_go_fuzz_dep_.CoverTab[3220]++

//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1961
			if i-begin == 0 {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1961
				_go_fuzz_dep_.CoverTab[3225]++
														return oldMsg, errZeroSegLen
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1962
				// _ = "end of CoverTab[3225]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1963
				_go_fuzz_dep_.CoverTab[3226]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1963
				// _ = "end of CoverTab[3226]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1963
			}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1963
			// _ = "end of CoverTab[3220]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1963
			_go_fuzz_dep_.CoverTab[3221]++

													msg = append(msg, byte(i-begin))

													for j := begin; j < i; j++ {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1967
				_go_fuzz_dep_.CoverTab[3227]++
														msg = append(msg, n.Data[j])
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1968
				// _ = "end of CoverTab[3227]"
			}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1969
			// _ = "end of CoverTab[3221]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1969
			_go_fuzz_dep_.CoverTab[3222]++

													begin = i + 1
													continue
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1972
			// _ = "end of CoverTab[3222]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1973
			_go_fuzz_dep_.CoverTab[3228]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1973
			// _ = "end of CoverTab[3228]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1973
		}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1973
		// _ = "end of CoverTab[3217]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1973
		_go_fuzz_dep_.CoverTab[3218]++

//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1978
		if (i == 0 || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1978
			_go_fuzz_dep_.CoverTab[3229]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1978
			return n.Data[i-1] == '.'
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1978
			// _ = "end of CoverTab[3229]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1978
		}()) && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1978
			_go_fuzz_dep_.CoverTab[3230]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1978
			return compression != nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1978
			// _ = "end of CoverTab[3230]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1978
		}() {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1978
			_go_fuzz_dep_.CoverTab[3231]++
													if ptr, ok := compression[string(n.Data[i:])]; ok {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1979
				_go_fuzz_dep_.CoverTab[3233]++

//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1982
				return append(msg, byte(ptr>>8|0xC0), byte(ptr)), nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1982
				// _ = "end of CoverTab[3233]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1983
				_go_fuzz_dep_.CoverTab[3234]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1983
				// _ = "end of CoverTab[3234]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1983
			}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1983
			// _ = "end of CoverTab[3231]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1983
			_go_fuzz_dep_.CoverTab[3232]++

//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1987
			if len(msg) <= int(^uint16(0)>>2) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1987
				_go_fuzz_dep_.CoverTab[3235]++
														compression[string(n.Data[i:])] = len(msg) - compressionOff
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1988
				// _ = "end of CoverTab[3235]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1989
				_go_fuzz_dep_.CoverTab[3236]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1989
				// _ = "end of CoverTab[3236]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1989
			}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1989
			// _ = "end of CoverTab[3232]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1990
			_go_fuzz_dep_.CoverTab[3237]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1990
			// _ = "end of CoverTab[3237]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1990
		}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1990
		// _ = "end of CoverTab[3218]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1991
	// _ = "end of CoverTab[3209]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1991
	_go_fuzz_dep_.CoverTab[3210]++
											return append(msg, 0), nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1992
	// _ = "end of CoverTab[3210]"
}

// unpack unpacks a domain name.
func (n *Name) unpack(msg []byte, off int) (int, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1996
	_go_fuzz_dep_.CoverTab[3238]++
											return n.unpackCompressed(msg, off, true)
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:1997
	// _ = "end of CoverTab[3238]"
}

func (n *Name) unpackCompressed(msg []byte, off int, allowCompression bool) (int, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2000
	_go_fuzz_dep_.CoverTab[3239]++

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
		_go_fuzz_dep_.CoverTab[3244]++
												if currOff >= len(msg) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2017
			_go_fuzz_dep_.CoverTab[3246]++
													return off, errBaseLen
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2018
			// _ = "end of CoverTab[3246]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2019
			_go_fuzz_dep_.CoverTab[3247]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2019
			// _ = "end of CoverTab[3247]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2019
		}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2019
		// _ = "end of CoverTab[3244]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2019
		_go_fuzz_dep_.CoverTab[3245]++
												c := int(msg[currOff])
												currOff++
												switch c & 0xC0 {
		case 0x00:
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2023
			_go_fuzz_dep_.CoverTab[3248]++
													if c == 0x00 {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2024
				_go_fuzz_dep_.CoverTab[3257]++

														break Loop
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2026
				// _ = "end of CoverTab[3257]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2027
				_go_fuzz_dep_.CoverTab[3258]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2027
				// _ = "end of CoverTab[3258]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2027
			}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2027
			// _ = "end of CoverTab[3248]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2027
			_go_fuzz_dep_.CoverTab[3249]++
													endOff := currOff + c
													if endOff > len(msg) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2029
				_go_fuzz_dep_.CoverTab[3259]++
														return off, errCalcLen
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2030
				// _ = "end of CoverTab[3259]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2031
				_go_fuzz_dep_.CoverTab[3260]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2031
				// _ = "end of CoverTab[3260]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2031
			}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2031
			// _ = "end of CoverTab[3249]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2031
			_go_fuzz_dep_.CoverTab[3250]++
													name = append(name, msg[currOff:endOff]...)
													name = append(name, '.')
													currOff = endOff
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2034
			// _ = "end of CoverTab[3250]"
		case 0xC0:
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2035
			_go_fuzz_dep_.CoverTab[3251]++
													if !allowCompression {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2036
				_go_fuzz_dep_.CoverTab[3261]++
														return off, errCompressedSRV
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2037
				// _ = "end of CoverTab[3261]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2038
				_go_fuzz_dep_.CoverTab[3262]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2038
				// _ = "end of CoverTab[3262]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2038
			}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2038
			// _ = "end of CoverTab[3251]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2038
			_go_fuzz_dep_.CoverTab[3252]++
													if currOff >= len(msg) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2039
				_go_fuzz_dep_.CoverTab[3263]++
														return off, errInvalidPtr
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2040
				// _ = "end of CoverTab[3263]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2041
				_go_fuzz_dep_.CoverTab[3264]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2041
				// _ = "end of CoverTab[3264]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2041
			}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2041
			// _ = "end of CoverTab[3252]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2041
			_go_fuzz_dep_.CoverTab[3253]++
													c1 := msg[currOff]
													currOff++
													if ptr == 0 {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2044
				_go_fuzz_dep_.CoverTab[3265]++
														newOff = currOff
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2045
				// _ = "end of CoverTab[3265]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2046
				_go_fuzz_dep_.CoverTab[3266]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2046
				// _ = "end of CoverTab[3266]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2046
			}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2046
			// _ = "end of CoverTab[3253]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2046
			_go_fuzz_dep_.CoverTab[3254]++

													if ptr++; ptr > 10 {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2048
				_go_fuzz_dep_.CoverTab[3267]++
														return off, errTooManyPtr
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2049
				// _ = "end of CoverTab[3267]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2050
				_go_fuzz_dep_.CoverTab[3268]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2050
				// _ = "end of CoverTab[3268]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2050
			}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2050
			// _ = "end of CoverTab[3254]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2050
			_go_fuzz_dep_.CoverTab[3255]++
													currOff = (c^0xC0)<<8 | int(c1)
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2051
			// _ = "end of CoverTab[3255]"
		default:
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2052
			_go_fuzz_dep_.CoverTab[3256]++

													return off, errReserved
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2054
			// _ = "end of CoverTab[3256]"
		}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2055
		// _ = "end of CoverTab[3245]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2056
	// _ = "end of CoverTab[3239]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2056
	_go_fuzz_dep_.CoverTab[3240]++
											if len(name) == 0 {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2057
		_go_fuzz_dep_.CoverTab[3269]++
												name = append(name, '.')
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2058
		// _ = "end of CoverTab[3269]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2059
		_go_fuzz_dep_.CoverTab[3270]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2059
		// _ = "end of CoverTab[3270]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2059
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2059
	// _ = "end of CoverTab[3240]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2059
	_go_fuzz_dep_.CoverTab[3241]++
											if len(name) > len(n.Data) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2060
		_go_fuzz_dep_.CoverTab[3271]++
												return off, errCalcLen
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2061
		// _ = "end of CoverTab[3271]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2062
		_go_fuzz_dep_.CoverTab[3272]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2062
		// _ = "end of CoverTab[3272]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2062
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2062
	// _ = "end of CoverTab[3241]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2062
	_go_fuzz_dep_.CoverTab[3242]++
											n.Length = uint8(len(name))
											if ptr == 0 {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2064
		_go_fuzz_dep_.CoverTab[3273]++
												newOff = currOff
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2065
		// _ = "end of CoverTab[3273]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2066
		_go_fuzz_dep_.CoverTab[3274]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2066
		// _ = "end of CoverTab[3274]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2066
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2066
	// _ = "end of CoverTab[3242]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2066
	_go_fuzz_dep_.CoverTab[3243]++
											return newOff, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2067
	// _ = "end of CoverTab[3243]"
}

func skipName(msg []byte, off int) (int, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2070
	_go_fuzz_dep_.CoverTab[3275]++

//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2074
	newOff := off

Loop:
	for {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2077
		_go_fuzz_dep_.CoverTab[3277]++
												if newOff >= len(msg) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2078
			_go_fuzz_dep_.CoverTab[3279]++
													return off, errBaseLen
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2079
			// _ = "end of CoverTab[3279]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2080
			_go_fuzz_dep_.CoverTab[3280]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2080
			// _ = "end of CoverTab[3280]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2080
		}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2080
		// _ = "end of CoverTab[3277]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2080
		_go_fuzz_dep_.CoverTab[3278]++
												c := int(msg[newOff])
												newOff++
												switch c & 0xC0 {
		case 0x00:
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2084
			_go_fuzz_dep_.CoverTab[3281]++
													if c == 0x00 {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2085
				_go_fuzz_dep_.CoverTab[3285]++

														break Loop
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2087
				// _ = "end of CoverTab[3285]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2088
				_go_fuzz_dep_.CoverTab[3286]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2088
				// _ = "end of CoverTab[3286]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2088
			}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2088
			// _ = "end of CoverTab[3281]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2088
			_go_fuzz_dep_.CoverTab[3282]++

													newOff += c
													if newOff > len(msg) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2091
				_go_fuzz_dep_.CoverTab[3287]++
														return off, errCalcLen
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2092
				// _ = "end of CoverTab[3287]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2093
				_go_fuzz_dep_.CoverTab[3288]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2093
				// _ = "end of CoverTab[3288]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2093
			}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2093
			// _ = "end of CoverTab[3282]"
		case 0xC0:
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2094
			_go_fuzz_dep_.CoverTab[3283]++

//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2098
			newOff++

//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2101
			break Loop
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2101
			// _ = "end of CoverTab[3283]"
		default:
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2102
			_go_fuzz_dep_.CoverTab[3284]++

													return off, errReserved
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2104
			// _ = "end of CoverTab[3284]"
		}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2105
		// _ = "end of CoverTab[3278]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2106
	// _ = "end of CoverTab[3275]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2106
	_go_fuzz_dep_.CoverTab[3276]++

											return newOff, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2108
	// _ = "end of CoverTab[3276]"
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
	_go_fuzz_dep_.CoverTab[3289]++
											msg, err := q.Name.pack(msg, compression, compressionOff)
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2121
		_go_fuzz_dep_.CoverTab[3291]++
												return msg, &nestedError{"Name", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2122
		// _ = "end of CoverTab[3291]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2123
		_go_fuzz_dep_.CoverTab[3292]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2123
		// _ = "end of CoverTab[3292]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2123
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2123
	// _ = "end of CoverTab[3289]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2123
	_go_fuzz_dep_.CoverTab[3290]++
											msg = packType(msg, q.Type)
											return packClass(msg, q.Class), nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2125
	// _ = "end of CoverTab[3290]"
}

// GoString implements fmt.GoStringer.GoString.
func (q *Question) GoString() string {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2129
	_go_fuzz_dep_.CoverTab[3293]++
											return "dnsmessage.Question{" +
		"Name: " + q.Name.GoString() + ", " +
		"Type: " + q.Type.GoString() + ", " +
		"Class: " + q.Class.GoString() + "}"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2133
	// _ = "end of CoverTab[3293]"
}

func unpackResourceBody(msg []byte, off int, hdr ResourceHeader) (ResourceBody, int, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2136
	_go_fuzz_dep_.CoverTab[3294]++
											var (
		r	ResourceBody
		err	error
		name	string
	)
	switch hdr.Type {
	case TypeA:
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2143
		_go_fuzz_dep_.CoverTab[3297]++
												var rb AResource
												rb, err = unpackAResource(msg, off)
												r = &rb
												name = "A"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2147
		// _ = "end of CoverTab[3297]"
	case TypeNS:
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2148
		_go_fuzz_dep_.CoverTab[3298]++
												var rb NSResource
												rb, err = unpackNSResource(msg, off)
												r = &rb
												name = "NS"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2152
		// _ = "end of CoverTab[3298]"
	case TypeCNAME:
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2153
		_go_fuzz_dep_.CoverTab[3299]++
												var rb CNAMEResource
												rb, err = unpackCNAMEResource(msg, off)
												r = &rb
												name = "CNAME"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2157
		// _ = "end of CoverTab[3299]"
	case TypeSOA:
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2158
		_go_fuzz_dep_.CoverTab[3300]++
												var rb SOAResource
												rb, err = unpackSOAResource(msg, off)
												r = &rb
												name = "SOA"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2162
		// _ = "end of CoverTab[3300]"
	case TypePTR:
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2163
		_go_fuzz_dep_.CoverTab[3301]++
												var rb PTRResource
												rb, err = unpackPTRResource(msg, off)
												r = &rb
												name = "PTR"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2167
		// _ = "end of CoverTab[3301]"
	case TypeMX:
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2168
		_go_fuzz_dep_.CoverTab[3302]++
												var rb MXResource
												rb, err = unpackMXResource(msg, off)
												r = &rb
												name = "MX"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2172
		// _ = "end of CoverTab[3302]"
	case TypeTXT:
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2173
		_go_fuzz_dep_.CoverTab[3303]++
												var rb TXTResource
												rb, err = unpackTXTResource(msg, off, hdr.Length)
												r = &rb
												name = "TXT"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2177
		// _ = "end of CoverTab[3303]"
	case TypeAAAA:
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2178
		_go_fuzz_dep_.CoverTab[3304]++
												var rb AAAAResource
												rb, err = unpackAAAAResource(msg, off)
												r = &rb
												name = "AAAA"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2182
		// _ = "end of CoverTab[3304]"
	case TypeSRV:
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2183
		_go_fuzz_dep_.CoverTab[3305]++
												var rb SRVResource
												rb, err = unpackSRVResource(msg, off)
												r = &rb
												name = "SRV"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2187
		// _ = "end of CoverTab[3305]"
	case TypeOPT:
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2188
		_go_fuzz_dep_.CoverTab[3306]++
												var rb OPTResource
												rb, err = unpackOPTResource(msg, off, hdr.Length)
												r = &rb
												name = "OPT"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2192
		// _ = "end of CoverTab[3306]"
	default:
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2193
		_go_fuzz_dep_.CoverTab[3307]++
												var rb UnknownResource
												rb, err = unpackUnknownResource(hdr.Type, msg, off, hdr.Length)
												r = &rb
												name = "Unknown"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2197
		// _ = "end of CoverTab[3307]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2198
	// _ = "end of CoverTab[3294]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2198
	_go_fuzz_dep_.CoverTab[3295]++
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2199
		_go_fuzz_dep_.CoverTab[3308]++
												return nil, off, &nestedError{name + " record", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2200
		// _ = "end of CoverTab[3308]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2201
		_go_fuzz_dep_.CoverTab[3309]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2201
		// _ = "end of CoverTab[3309]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2201
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2201
	// _ = "end of CoverTab[3295]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2201
	_go_fuzz_dep_.CoverTab[3296]++
											return r, off + int(hdr.Length), nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2202
	// _ = "end of CoverTab[3296]"
}

// A CNAMEResource is a CNAME Resource record.
type CNAMEResource struct {
	CNAME Name
}

func (r *CNAMEResource) realType() Type {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2210
	_go_fuzz_dep_.CoverTab[3310]++
											return TypeCNAME
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2211
	// _ = "end of CoverTab[3310]"
}

// pack appends the wire format of the CNAMEResource to msg.
func (r *CNAMEResource) pack(msg []byte, compression map[string]int, compressionOff int) ([]byte, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2215
	_go_fuzz_dep_.CoverTab[3311]++
											return r.CNAME.pack(msg, compression, compressionOff)
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2216
	// _ = "end of CoverTab[3311]"
}

// GoString implements fmt.GoStringer.GoString.
func (r *CNAMEResource) GoString() string {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2220
	_go_fuzz_dep_.CoverTab[3312]++
											return "dnsmessage.CNAMEResource{CNAME: " + r.CNAME.GoString() + "}"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2221
	// _ = "end of CoverTab[3312]"
}

func unpackCNAMEResource(msg []byte, off int) (CNAMEResource, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2224
	_go_fuzz_dep_.CoverTab[3313]++
											var cname Name
											if _, err := cname.unpack(msg, off); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2226
		_go_fuzz_dep_.CoverTab[3315]++
												return CNAMEResource{}, err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2227
		// _ = "end of CoverTab[3315]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2228
		_go_fuzz_dep_.CoverTab[3316]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2228
		// _ = "end of CoverTab[3316]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2228
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2228
	// _ = "end of CoverTab[3313]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2228
	_go_fuzz_dep_.CoverTab[3314]++
											return CNAMEResource{cname}, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2229
	// _ = "end of CoverTab[3314]"
}

// An MXResource is an MX Resource record.
type MXResource struct {
	Pref	uint16
	MX	Name
}

func (r *MXResource) realType() Type {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2238
	_go_fuzz_dep_.CoverTab[3317]++
											return TypeMX
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2239
	// _ = "end of CoverTab[3317]"
}

// pack appends the wire format of the MXResource to msg.
func (r *MXResource) pack(msg []byte, compression map[string]int, compressionOff int) ([]byte, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2243
	_go_fuzz_dep_.CoverTab[3318]++
											oldMsg := msg
											msg = packUint16(msg, r.Pref)
											msg, err := r.MX.pack(msg, compression, compressionOff)
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2247
		_go_fuzz_dep_.CoverTab[3320]++
												return oldMsg, &nestedError{"MXResource.MX", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2248
		// _ = "end of CoverTab[3320]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2249
		_go_fuzz_dep_.CoverTab[3321]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2249
		// _ = "end of CoverTab[3321]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2249
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2249
	// _ = "end of CoverTab[3318]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2249
	_go_fuzz_dep_.CoverTab[3319]++
											return msg, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2250
	// _ = "end of CoverTab[3319]"
}

// GoString implements fmt.GoStringer.GoString.
func (r *MXResource) GoString() string {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2254
	_go_fuzz_dep_.CoverTab[3322]++
											return "dnsmessage.MXResource{" +
		"Pref: " + printUint16(r.Pref) + ", " +
		"MX: " + r.MX.GoString() + "}"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2257
	// _ = "end of CoverTab[3322]"
}

func unpackMXResource(msg []byte, off int) (MXResource, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2260
	_go_fuzz_dep_.CoverTab[3323]++
											pref, off, err := unpackUint16(msg, off)
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2262
		_go_fuzz_dep_.CoverTab[3326]++
												return MXResource{}, &nestedError{"Pref", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2263
		// _ = "end of CoverTab[3326]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2264
		_go_fuzz_dep_.CoverTab[3327]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2264
		// _ = "end of CoverTab[3327]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2264
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2264
	// _ = "end of CoverTab[3323]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2264
	_go_fuzz_dep_.CoverTab[3324]++
											var mx Name
											if _, err := mx.unpack(msg, off); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2266
		_go_fuzz_dep_.CoverTab[3328]++
												return MXResource{}, &nestedError{"MX", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2267
		// _ = "end of CoverTab[3328]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2268
		_go_fuzz_dep_.CoverTab[3329]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2268
		// _ = "end of CoverTab[3329]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2268
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2268
	// _ = "end of CoverTab[3324]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2268
	_go_fuzz_dep_.CoverTab[3325]++
											return MXResource{pref, mx}, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2269
	// _ = "end of CoverTab[3325]"
}

// An NSResource is an NS Resource record.
type NSResource struct {
	NS Name
}

func (r *NSResource) realType() Type {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2277
	_go_fuzz_dep_.CoverTab[3330]++
											return TypeNS
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2278
	// _ = "end of CoverTab[3330]"
}

// pack appends the wire format of the NSResource to msg.
func (r *NSResource) pack(msg []byte, compression map[string]int, compressionOff int) ([]byte, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2282
	_go_fuzz_dep_.CoverTab[3331]++
											return r.NS.pack(msg, compression, compressionOff)
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2283
	// _ = "end of CoverTab[3331]"
}

// GoString implements fmt.GoStringer.GoString.
func (r *NSResource) GoString() string {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2287
	_go_fuzz_dep_.CoverTab[3332]++
											return "dnsmessage.NSResource{NS: " + r.NS.GoString() + "}"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2288
	// _ = "end of CoverTab[3332]"
}

func unpackNSResource(msg []byte, off int) (NSResource, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2291
	_go_fuzz_dep_.CoverTab[3333]++
											var ns Name
											if _, err := ns.unpack(msg, off); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2293
		_go_fuzz_dep_.CoverTab[3335]++
												return NSResource{}, err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2294
		// _ = "end of CoverTab[3335]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2295
		_go_fuzz_dep_.CoverTab[3336]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2295
		// _ = "end of CoverTab[3336]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2295
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2295
	// _ = "end of CoverTab[3333]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2295
	_go_fuzz_dep_.CoverTab[3334]++
											return NSResource{ns}, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2296
	// _ = "end of CoverTab[3334]"
}

// A PTRResource is a PTR Resource record.
type PTRResource struct {
	PTR Name
}

func (r *PTRResource) realType() Type {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2304
	_go_fuzz_dep_.CoverTab[3337]++
											return TypePTR
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2305
	// _ = "end of CoverTab[3337]"
}

// pack appends the wire format of the PTRResource to msg.
func (r *PTRResource) pack(msg []byte, compression map[string]int, compressionOff int) ([]byte, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2309
	_go_fuzz_dep_.CoverTab[3338]++
											return r.PTR.pack(msg, compression, compressionOff)
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2310
	// _ = "end of CoverTab[3338]"
}

// GoString implements fmt.GoStringer.GoString.
func (r *PTRResource) GoString() string {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2314
	_go_fuzz_dep_.CoverTab[3339]++
											return "dnsmessage.PTRResource{PTR: " + r.PTR.GoString() + "}"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2315
	// _ = "end of CoverTab[3339]"
}

func unpackPTRResource(msg []byte, off int) (PTRResource, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2318
	_go_fuzz_dep_.CoverTab[3340]++
											var ptr Name
											if _, err := ptr.unpack(msg, off); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2320
		_go_fuzz_dep_.CoverTab[3342]++
												return PTRResource{}, err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2321
		// _ = "end of CoverTab[3342]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2322
		_go_fuzz_dep_.CoverTab[3343]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2322
		// _ = "end of CoverTab[3343]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2322
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2322
	// _ = "end of CoverTab[3340]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2322
	_go_fuzz_dep_.CoverTab[3341]++
											return PTRResource{ptr}, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2323
	// _ = "end of CoverTab[3341]"
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
	_go_fuzz_dep_.CoverTab[3344]++
											return TypeSOA
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2342
	// _ = "end of CoverTab[3344]"
}

// pack appends the wire format of the SOAResource to msg.
func (r *SOAResource) pack(msg []byte, compression map[string]int, compressionOff int) ([]byte, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2346
	_go_fuzz_dep_.CoverTab[3345]++
											oldMsg := msg
											msg, err := r.NS.pack(msg, compression, compressionOff)
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2349
		_go_fuzz_dep_.CoverTab[3348]++
												return oldMsg, &nestedError{"SOAResource.NS", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2350
		// _ = "end of CoverTab[3348]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2351
		_go_fuzz_dep_.CoverTab[3349]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2351
		// _ = "end of CoverTab[3349]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2351
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2351
	// _ = "end of CoverTab[3345]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2351
	_go_fuzz_dep_.CoverTab[3346]++
											msg, err = r.MBox.pack(msg, compression, compressionOff)
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2353
		_go_fuzz_dep_.CoverTab[3350]++
												return oldMsg, &nestedError{"SOAResource.MBox", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2354
		// _ = "end of CoverTab[3350]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2355
		_go_fuzz_dep_.CoverTab[3351]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2355
		// _ = "end of CoverTab[3351]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2355
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2355
	// _ = "end of CoverTab[3346]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2355
	_go_fuzz_dep_.CoverTab[3347]++
											msg = packUint32(msg, r.Serial)
											msg = packUint32(msg, r.Refresh)
											msg = packUint32(msg, r.Retry)
											msg = packUint32(msg, r.Expire)
											return packUint32(msg, r.MinTTL), nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2360
	// _ = "end of CoverTab[3347]"
}

// GoString implements fmt.GoStringer.GoString.
func (r *SOAResource) GoString() string {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2364
	_go_fuzz_dep_.CoverTab[3352]++
											return "dnsmessage.SOAResource{" +
		"NS: " + r.NS.GoString() + ", " +
		"MBox: " + r.MBox.GoString() + ", " +
		"Serial: " + printUint32(r.Serial) + ", " +
		"Refresh: " + printUint32(r.Refresh) + ", " +
		"Retry: " + printUint32(r.Retry) + ", " +
		"Expire: " + printUint32(r.Expire) + ", " +
		"MinTTL: " + printUint32(r.MinTTL) + "}"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2372
	// _ = "end of CoverTab[3352]"
}

func unpackSOAResource(msg []byte, off int) (SOAResource, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2375
	_go_fuzz_dep_.CoverTab[3353]++
											var ns Name
											off, err := ns.unpack(msg, off)
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2378
		_go_fuzz_dep_.CoverTab[3361]++
												return SOAResource{}, &nestedError{"NS", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2379
		// _ = "end of CoverTab[3361]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2380
		_go_fuzz_dep_.CoverTab[3362]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2380
		// _ = "end of CoverTab[3362]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2380
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2380
	// _ = "end of CoverTab[3353]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2380
	_go_fuzz_dep_.CoverTab[3354]++
											var mbox Name
											if off, err = mbox.unpack(msg, off); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2382
		_go_fuzz_dep_.CoverTab[3363]++
												return SOAResource{}, &nestedError{"MBox", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2383
		// _ = "end of CoverTab[3363]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2384
		_go_fuzz_dep_.CoverTab[3364]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2384
		// _ = "end of CoverTab[3364]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2384
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2384
	// _ = "end of CoverTab[3354]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2384
	_go_fuzz_dep_.CoverTab[3355]++
											serial, off, err := unpackUint32(msg, off)
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2386
		_go_fuzz_dep_.CoverTab[3365]++
												return SOAResource{}, &nestedError{"Serial", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2387
		// _ = "end of CoverTab[3365]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2388
		_go_fuzz_dep_.CoverTab[3366]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2388
		// _ = "end of CoverTab[3366]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2388
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2388
	// _ = "end of CoverTab[3355]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2388
	_go_fuzz_dep_.CoverTab[3356]++
											refresh, off, err := unpackUint32(msg, off)
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2390
		_go_fuzz_dep_.CoverTab[3367]++
												return SOAResource{}, &nestedError{"Refresh", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2391
		// _ = "end of CoverTab[3367]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2392
		_go_fuzz_dep_.CoverTab[3368]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2392
		// _ = "end of CoverTab[3368]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2392
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2392
	// _ = "end of CoverTab[3356]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2392
	_go_fuzz_dep_.CoverTab[3357]++
											retry, off, err := unpackUint32(msg, off)
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2394
		_go_fuzz_dep_.CoverTab[3369]++
												return SOAResource{}, &nestedError{"Retry", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2395
		// _ = "end of CoverTab[3369]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2396
		_go_fuzz_dep_.CoverTab[3370]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2396
		// _ = "end of CoverTab[3370]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2396
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2396
	// _ = "end of CoverTab[3357]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2396
	_go_fuzz_dep_.CoverTab[3358]++
											expire, off, err := unpackUint32(msg, off)
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2398
		_go_fuzz_dep_.CoverTab[3371]++
												return SOAResource{}, &nestedError{"Expire", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2399
		// _ = "end of CoverTab[3371]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2400
		_go_fuzz_dep_.CoverTab[3372]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2400
		// _ = "end of CoverTab[3372]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2400
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2400
	// _ = "end of CoverTab[3358]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2400
	_go_fuzz_dep_.CoverTab[3359]++
											minTTL, _, err := unpackUint32(msg, off)
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2402
		_go_fuzz_dep_.CoverTab[3373]++
												return SOAResource{}, &nestedError{"MinTTL", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2403
		// _ = "end of CoverTab[3373]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2404
		_go_fuzz_dep_.CoverTab[3374]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2404
		// _ = "end of CoverTab[3374]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2404
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2404
	// _ = "end of CoverTab[3359]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2404
	_go_fuzz_dep_.CoverTab[3360]++
											return SOAResource{ns, mbox, serial, refresh, retry, expire, minTTL}, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2405
	// _ = "end of CoverTab[3360]"
}

// A TXTResource is a TXT Resource record.
type TXTResource struct {
	TXT []string
}

func (r *TXTResource) realType() Type {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2413
	_go_fuzz_dep_.CoverTab[3375]++
											return TypeTXT
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2414
	// _ = "end of CoverTab[3375]"
}

// pack appends the wire format of the TXTResource to msg.
func (r *TXTResource) pack(msg []byte, compression map[string]int, compressionOff int) ([]byte, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2418
	_go_fuzz_dep_.CoverTab[3376]++
											oldMsg := msg
											for _, s := range r.TXT {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2420
		_go_fuzz_dep_.CoverTab[3378]++
												var err error
												msg, err = packText(msg, s)
												if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2423
			_go_fuzz_dep_.CoverTab[3379]++
													return oldMsg, err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2424
			// _ = "end of CoverTab[3379]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2425
			_go_fuzz_dep_.CoverTab[3380]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2425
			// _ = "end of CoverTab[3380]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2425
		}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2425
		// _ = "end of CoverTab[3378]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2426
	// _ = "end of CoverTab[3376]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2426
	_go_fuzz_dep_.CoverTab[3377]++
											return msg, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2427
	// _ = "end of CoverTab[3377]"
}

// GoString implements fmt.GoStringer.GoString.
func (r *TXTResource) GoString() string {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2431
	_go_fuzz_dep_.CoverTab[3381]++
											s := "dnsmessage.TXTResource{TXT: []string{"
											if len(r.TXT) == 0 {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2433
		_go_fuzz_dep_.CoverTab[3384]++
												return s + "}}"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2434
		// _ = "end of CoverTab[3384]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2435
		_go_fuzz_dep_.CoverTab[3385]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2435
		// _ = "end of CoverTab[3385]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2435
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2435
	// _ = "end of CoverTab[3381]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2435
	_go_fuzz_dep_.CoverTab[3382]++
											s += `"` + printString([]byte(r.TXT[0]))
											for _, t := range r.TXT[1:] {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2437
		_go_fuzz_dep_.CoverTab[3386]++
												s += `", "` + printString([]byte(t))
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2438
		// _ = "end of CoverTab[3386]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2439
	// _ = "end of CoverTab[3382]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2439
	_go_fuzz_dep_.CoverTab[3383]++
											return s + `"}}`
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2440
	// _ = "end of CoverTab[3383]"
}

func unpackTXTResource(msg []byte, off int, length uint16) (TXTResource, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2443
	_go_fuzz_dep_.CoverTab[3387]++
											txts := make([]string, 0, 1)
											for n := uint16(0); n < length; {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2445
		_go_fuzz_dep_.CoverTab[3389]++
												var t string
												var err error
												if t, off, err = unpackText(msg, off); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2448
			_go_fuzz_dep_.CoverTab[3392]++
													return TXTResource{}, &nestedError{"text", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2449
			// _ = "end of CoverTab[3392]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2450
			_go_fuzz_dep_.CoverTab[3393]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2450
			// _ = "end of CoverTab[3393]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2450
		}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2450
		// _ = "end of CoverTab[3389]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2450
		_go_fuzz_dep_.CoverTab[3390]++

												if length-n < uint16(len(t))+1 {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2452
			_go_fuzz_dep_.CoverTab[3394]++
													return TXTResource{}, errCalcLen
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2453
			// _ = "end of CoverTab[3394]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2454
			_go_fuzz_dep_.CoverTab[3395]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2454
			// _ = "end of CoverTab[3395]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2454
		}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2454
		// _ = "end of CoverTab[3390]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2454
		_go_fuzz_dep_.CoverTab[3391]++
												n += uint16(len(t)) + 1
												txts = append(txts, t)
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2456
		// _ = "end of CoverTab[3391]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2457
	// _ = "end of CoverTab[3387]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2457
	_go_fuzz_dep_.CoverTab[3388]++
											return TXTResource{txts}, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2458
	// _ = "end of CoverTab[3388]"
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
	_go_fuzz_dep_.CoverTab[3396]++
											return TypeSRV
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2470
	// _ = "end of CoverTab[3396]"
}

// pack appends the wire format of the SRVResource to msg.
func (r *SRVResource) pack(msg []byte, compression map[string]int, compressionOff int) ([]byte, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2474
	_go_fuzz_dep_.CoverTab[3397]++
											oldMsg := msg
											msg = packUint16(msg, r.Priority)
											msg = packUint16(msg, r.Weight)
											msg = packUint16(msg, r.Port)
											msg, err := r.Target.pack(msg, nil, compressionOff)
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2480
		_go_fuzz_dep_.CoverTab[3399]++
												return oldMsg, &nestedError{"SRVResource.Target", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2481
		// _ = "end of CoverTab[3399]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2482
		_go_fuzz_dep_.CoverTab[3400]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2482
		// _ = "end of CoverTab[3400]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2482
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2482
	// _ = "end of CoverTab[3397]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2482
	_go_fuzz_dep_.CoverTab[3398]++
											return msg, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2483
	// _ = "end of CoverTab[3398]"
}

// GoString implements fmt.GoStringer.GoString.
func (r *SRVResource) GoString() string {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2487
	_go_fuzz_dep_.CoverTab[3401]++
											return "dnsmessage.SRVResource{" +
		"Priority: " + printUint16(r.Priority) + ", " +
		"Weight: " + printUint16(r.Weight) + ", " +
		"Port: " + printUint16(r.Port) + ", " +
		"Target: " + r.Target.GoString() + "}"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2492
	// _ = "end of CoverTab[3401]"
}

func unpackSRVResource(msg []byte, off int) (SRVResource, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2495
	_go_fuzz_dep_.CoverTab[3402]++
											priority, off, err := unpackUint16(msg, off)
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2497
		_go_fuzz_dep_.CoverTab[3407]++
												return SRVResource{}, &nestedError{"Priority", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2498
		// _ = "end of CoverTab[3407]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2499
		_go_fuzz_dep_.CoverTab[3408]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2499
		// _ = "end of CoverTab[3408]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2499
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2499
	// _ = "end of CoverTab[3402]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2499
	_go_fuzz_dep_.CoverTab[3403]++
											weight, off, err := unpackUint16(msg, off)
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2501
		_go_fuzz_dep_.CoverTab[3409]++
												return SRVResource{}, &nestedError{"Weight", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2502
		// _ = "end of CoverTab[3409]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2503
		_go_fuzz_dep_.CoverTab[3410]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2503
		// _ = "end of CoverTab[3410]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2503
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2503
	// _ = "end of CoverTab[3403]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2503
	_go_fuzz_dep_.CoverTab[3404]++
											port, off, err := unpackUint16(msg, off)
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2505
		_go_fuzz_dep_.CoverTab[3411]++
												return SRVResource{}, &nestedError{"Port", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2506
		// _ = "end of CoverTab[3411]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2507
		_go_fuzz_dep_.CoverTab[3412]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2507
		// _ = "end of CoverTab[3412]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2507
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2507
	// _ = "end of CoverTab[3404]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2507
	_go_fuzz_dep_.CoverTab[3405]++
											var target Name
											if _, err := target.unpackCompressed(msg, off, false); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2509
		_go_fuzz_dep_.CoverTab[3413]++
												return SRVResource{}, &nestedError{"Target", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2510
		// _ = "end of CoverTab[3413]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2511
		_go_fuzz_dep_.CoverTab[3414]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2511
		// _ = "end of CoverTab[3414]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2511
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2511
	// _ = "end of CoverTab[3405]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2511
	_go_fuzz_dep_.CoverTab[3406]++
											return SRVResource{priority, weight, port, target}, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2512
	// _ = "end of CoverTab[3406]"
}

// An AResource is an A Resource record.
type AResource struct {
	A [4]byte
}

func (r *AResource) realType() Type {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2520
	_go_fuzz_dep_.CoverTab[3415]++
											return TypeA
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2521
	// _ = "end of CoverTab[3415]"
}

// pack appends the wire format of the AResource to msg.
func (r *AResource) pack(msg []byte, compression map[string]int, compressionOff int) ([]byte, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2525
	_go_fuzz_dep_.CoverTab[3416]++
											return packBytes(msg, r.A[:]), nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2526
	// _ = "end of CoverTab[3416]"
}

// GoString implements fmt.GoStringer.GoString.
func (r *AResource) GoString() string {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2530
	_go_fuzz_dep_.CoverTab[3417]++
											return "dnsmessage.AResource{" +
		"A: [4]byte{" + printByteSlice(r.A[:]) + "}}"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2532
	// _ = "end of CoverTab[3417]"
}

func unpackAResource(msg []byte, off int) (AResource, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2535
	_go_fuzz_dep_.CoverTab[3418]++
											var a [4]byte
											if _, err := unpackBytes(msg, off, a[:]); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2537
		_go_fuzz_dep_.CoverTab[3420]++
												return AResource{}, err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2538
		// _ = "end of CoverTab[3420]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2539
		_go_fuzz_dep_.CoverTab[3421]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2539
		// _ = "end of CoverTab[3421]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2539
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2539
	// _ = "end of CoverTab[3418]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2539
	_go_fuzz_dep_.CoverTab[3419]++
											return AResource{a}, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2540
	// _ = "end of CoverTab[3419]"
}

// An AAAAResource is an AAAA Resource record.
type AAAAResource struct {
	AAAA [16]byte
}

func (r *AAAAResource) realType() Type {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2548
	_go_fuzz_dep_.CoverTab[3422]++
											return TypeAAAA
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2549
	// _ = "end of CoverTab[3422]"
}

// GoString implements fmt.GoStringer.GoString.
func (r *AAAAResource) GoString() string {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2553
	_go_fuzz_dep_.CoverTab[3423]++
											return "dnsmessage.AAAAResource{" +
		"AAAA: [16]byte{" + printByteSlice(r.AAAA[:]) + "}}"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2555
	// _ = "end of CoverTab[3423]"
}

// pack appends the wire format of the AAAAResource to msg.
func (r *AAAAResource) pack(msg []byte, compression map[string]int, compressionOff int) ([]byte, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2559
	_go_fuzz_dep_.CoverTab[3424]++
											return packBytes(msg, r.AAAA[:]), nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2560
	// _ = "end of CoverTab[3424]"
}

func unpackAAAAResource(msg []byte, off int) (AAAAResource, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2563
	_go_fuzz_dep_.CoverTab[3425]++
											var aaaa [16]byte
											if _, err := unpackBytes(msg, off, aaaa[:]); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2565
		_go_fuzz_dep_.CoverTab[3427]++
												return AAAAResource{}, err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2566
		// _ = "end of CoverTab[3427]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2567
		_go_fuzz_dep_.CoverTab[3428]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2567
		// _ = "end of CoverTab[3428]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2567
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2567
	// _ = "end of CoverTab[3425]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2567
	_go_fuzz_dep_.CoverTab[3426]++
											return AAAAResource{aaaa}, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2568
	// _ = "end of CoverTab[3426]"
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
	_go_fuzz_dep_.CoverTab[3429]++
											return "dnsmessage.Option{" +
		"Code: " + printUint16(o.Code) + ", " +
		"Data: []byte{" + printByteSlice(o.Data) + "}}"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2592
	// _ = "end of CoverTab[3429]"
}

func (r *OPTResource) realType() Type {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2595
	_go_fuzz_dep_.CoverTab[3430]++
											return TypeOPT
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2596
	// _ = "end of CoverTab[3430]"
}

func (r *OPTResource) pack(msg []byte, compression map[string]int, compressionOff int) ([]byte, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2599
	_go_fuzz_dep_.CoverTab[3431]++
											for _, opt := range r.Options {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2600
		_go_fuzz_dep_.CoverTab[3433]++
												msg = packUint16(msg, opt.Code)
												l := uint16(len(opt.Data))
												msg = packUint16(msg, l)
												msg = packBytes(msg, opt.Data)
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2604
		// _ = "end of CoverTab[3433]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2605
	// _ = "end of CoverTab[3431]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2605
	_go_fuzz_dep_.CoverTab[3432]++
											return msg, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2606
	// _ = "end of CoverTab[3432]"
}

// GoString implements fmt.GoStringer.GoString.
func (r *OPTResource) GoString() string {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2610
	_go_fuzz_dep_.CoverTab[3434]++
											s := "dnsmessage.OPTResource{Options: []dnsmessage.Option{"
											if len(r.Options) == 0 {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2612
		_go_fuzz_dep_.CoverTab[3437]++
												return s + "}}"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2613
		// _ = "end of CoverTab[3437]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2614
		_go_fuzz_dep_.CoverTab[3438]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2614
		// _ = "end of CoverTab[3438]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2614
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2614
	// _ = "end of CoverTab[3434]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2614
	_go_fuzz_dep_.CoverTab[3435]++
											s += r.Options[0].GoString()
											for _, o := range r.Options[1:] {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2616
		_go_fuzz_dep_.CoverTab[3439]++
												s += ", " + o.GoString()
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2617
		// _ = "end of CoverTab[3439]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2618
	// _ = "end of CoverTab[3435]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2618
	_go_fuzz_dep_.CoverTab[3436]++
											return s + "}}"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2619
	// _ = "end of CoverTab[3436]"
}

func unpackOPTResource(msg []byte, off int, length uint16) (OPTResource, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2622
	_go_fuzz_dep_.CoverTab[3440]++
											var opts []Option
											for oldOff := off; off < oldOff+int(length); {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2624
		_go_fuzz_dep_.CoverTab[3442]++
												var err error
												var o Option
												o.Code, off, err = unpackUint16(msg, off)
												if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2628
			_go_fuzz_dep_.CoverTab[3446]++
													return OPTResource{}, &nestedError{"Code", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2629
			// _ = "end of CoverTab[3446]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2630
			_go_fuzz_dep_.CoverTab[3447]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2630
			// _ = "end of CoverTab[3447]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2630
		}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2630
		// _ = "end of CoverTab[3442]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2630
		_go_fuzz_dep_.CoverTab[3443]++
												var l uint16
												l, off, err = unpackUint16(msg, off)
												if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2633
			_go_fuzz_dep_.CoverTab[3448]++
													return OPTResource{}, &nestedError{"Data", err}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2634
			// _ = "end of CoverTab[3448]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2635
			_go_fuzz_dep_.CoverTab[3449]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2635
			// _ = "end of CoverTab[3449]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2635
		}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2635
		// _ = "end of CoverTab[3443]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2635
		_go_fuzz_dep_.CoverTab[3444]++
												o.Data = make([]byte, l)
												if copy(o.Data, msg[off:]) != int(l) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2637
			_go_fuzz_dep_.CoverTab[3450]++
													return OPTResource{}, &nestedError{"Data", errCalcLen}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2638
			// _ = "end of CoverTab[3450]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2639
			_go_fuzz_dep_.CoverTab[3451]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2639
			// _ = "end of CoverTab[3451]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2639
		}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2639
		// _ = "end of CoverTab[3444]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2639
		_go_fuzz_dep_.CoverTab[3445]++
												off += int(l)
												opts = append(opts, o)
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2641
		// _ = "end of CoverTab[3445]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2642
	// _ = "end of CoverTab[3440]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2642
	_go_fuzz_dep_.CoverTab[3441]++
											return OPTResource{opts}, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2643
	// _ = "end of CoverTab[3441]"
}

// An UnknownResource is a catch-all container for unknown record types.
type UnknownResource struct {
	Type	Type
	Data	[]byte
}

func (r *UnknownResource) realType() Type {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2652
	_go_fuzz_dep_.CoverTab[3452]++
											return r.Type
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2653
	// _ = "end of CoverTab[3452]"
}

// pack appends the wire format of the UnknownResource to msg.
func (r *UnknownResource) pack(msg []byte, compression map[string]int, compressionOff int) ([]byte, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2657
	_go_fuzz_dep_.CoverTab[3453]++
											return packBytes(msg, r.Data[:]), nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2658
	// _ = "end of CoverTab[3453]"
}

// GoString implements fmt.GoStringer.GoString.
func (r *UnknownResource) GoString() string {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2662
	_go_fuzz_dep_.CoverTab[3454]++
											return "dnsmessage.UnknownResource{" +
		"Type: " + r.Type.GoString() + ", " +
		"Data: []byte{" + printByteSlice(r.Data) + "}}"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2665
	// _ = "end of CoverTab[3454]"
}

func unpackUnknownResource(recordType Type, msg []byte, off int, length uint16) (UnknownResource, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2668
	_go_fuzz_dep_.CoverTab[3455]++
											parsed := UnknownResource{
		Type:	recordType,
		Data:	make([]byte, length),
	}
	if _, err := unpackBytes(msg, off, parsed.Data); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2673
		_go_fuzz_dep_.CoverTab[3457]++
												return UnknownResource{}, err
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2674
		// _ = "end of CoverTab[3457]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2675
		_go_fuzz_dep_.CoverTab[3458]++
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2675
		// _ = "end of CoverTab[3458]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2675
	}
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2675
	// _ = "end of CoverTab[3455]"
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2675
	_go_fuzz_dep_.CoverTab[3456]++
											return parsed, nil
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2676
	// _ = "end of CoverTab[3456]"
}

//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2677
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/vendor/golang.org/x/net/dns/dnsmessage/message.go:2677
var _ = _go_fuzz_dep_.CoverTab
