// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /snap/go/10455/src/net/netip/netip.go:5
// Package netip defines an IP address type that's a small value type.
//line /snap/go/10455/src/net/netip/netip.go:5
// Building on that [Addr] type, the package also defines [AddrPort] (an
//line /snap/go/10455/src/net/netip/netip.go:5
// IP address and a port) and [Prefix] (an IP address and a bit length
//line /snap/go/10455/src/net/netip/netip.go:5
// prefix).
//line /snap/go/10455/src/net/netip/netip.go:5
//
//line /snap/go/10455/src/net/netip/netip.go:5
// Compared to the [net.IP] type, [Addr] type takes less memory, is immutable,
//line /snap/go/10455/src/net/netip/netip.go:5
// and is comparable (supports == and being a map key).
//line /snap/go/10455/src/net/netip/netip.go:12
package netip

//line /snap/go/10455/src/net/netip/netip.go:12
import (
//line /snap/go/10455/src/net/netip/netip.go:12
	_go_fuzz_dep_ "go-fuzz-dep"
//line /snap/go/10455/src/net/netip/netip.go:12
)
//line /snap/go/10455/src/net/netip/netip.go:12
import (
//line /snap/go/10455/src/net/netip/netip.go:12
	_atomic_ "sync/atomic"
//line /snap/go/10455/src/net/netip/netip.go:12
)

import (
	"errors"
	"math"
	"strconv"

	"internal/bytealg"
	"internal/intern"
	"internal/itoa"
)

//line /snap/go/10455/src/net/netip/netip.go:29
// Addr represents an IPv4 or IPv6 address (with or without a scoped
//line /snap/go/10455/src/net/netip/netip.go:29
// addressing zone), similar to [net.IP] or [net.IPAddr].
//line /snap/go/10455/src/net/netip/netip.go:29
//
//line /snap/go/10455/src/net/netip/netip.go:29
// Unlike [net.IP] or [net.IPAddr], Addr is a comparable value
//line /snap/go/10455/src/net/netip/netip.go:29
// type (it supports == and can be a map key) and is immutable.
//line /snap/go/10455/src/net/netip/netip.go:29
//
//line /snap/go/10455/src/net/netip/netip.go:29
// The zero Addr is not a valid IP address.
//line /snap/go/10455/src/net/netip/netip.go:29
// Addr{} is distinct from both 0.0.0.0 and ::.
//line /snap/go/10455/src/net/netip/netip.go:37
type Addr struct {
	// addr is the hi and lo bits of an IPv6 address. If z==z4,
	// hi and lo contain the IPv4-mapped IPv6 address.
	//
	// hi and lo are constructed by interpreting a 16-byte IPv6
	// address as a big-endian 128-bit number. The most significant
	// bits of that number go into hi, the rest into lo.
	//
	// For example, 0011:2233:4455:6677:8899:aabb:ccdd:eeff is stored as:
	//  addr.hi = 0x0011223344556677
	//  addr.lo = 0x8899aabbccddeeff
	//
	// We store IPs like this, rather than as [16]byte, because it
	// turns most operations on IPs into arithmetic and bit-twiddling
	// operations on 64-bit registers, which is much faster than
	// bytewise processing.
	addr	uint128

	// z is a combination of the address family and the IPv6 zone.
	//
	// nil means invalid IP address (for a zero Addr).
	// z4 means an IPv4 address.
	// z6noz means an IPv6 address without a zone.
	//
	// Otherwise it's the interned zone name string.
	z	*intern.Value
}

// z0, z4, and z6noz are sentinel Addr.z values.
//line /snap/go/10455/src/net/netip/netip.go:65
// See the Addr type's field docs.
//line /snap/go/10455/src/net/netip/netip.go:67
var (
	z0	= (*intern.Value)(nil)
	z4	= new(intern.Value)
	z6noz	= new(intern.Value)
)

// IPv6LinkLocalAllNodes returns the IPv6 link-local all nodes multicast
//line /snap/go/10455/src/net/netip/netip.go:73
// address ff02::1.
//line /snap/go/10455/src/net/netip/netip.go:75
func IPv6LinkLocalAllNodes() Addr {
//line /snap/go/10455/src/net/netip/netip.go:75
	_go_fuzz_dep_.CoverTab[3892]++
//line /snap/go/10455/src/net/netip/netip.go:75
	return AddrFrom16([16]byte{0: 0xff, 1: 0x02, 15: 0x01})
//line /snap/go/10455/src/net/netip/netip.go:75
	// _ = "end of CoverTab[3892]"
//line /snap/go/10455/src/net/netip/netip.go:75
}

// IPv6LinkLocalAllRouters returns the IPv6 link-local all routers multicast
//line /snap/go/10455/src/net/netip/netip.go:77
// address ff02::2.
//line /snap/go/10455/src/net/netip/netip.go:79
func IPv6LinkLocalAllRouters() Addr {
//line /snap/go/10455/src/net/netip/netip.go:79
	_go_fuzz_dep_.CoverTab[3893]++
//line /snap/go/10455/src/net/netip/netip.go:79
	return AddrFrom16([16]byte{0: 0xff, 1: 0x02, 15: 0x02})
//line /snap/go/10455/src/net/netip/netip.go:79
	// _ = "end of CoverTab[3893]"
//line /snap/go/10455/src/net/netip/netip.go:79
}

// IPv6Loopback returns the IPv6 loopback address ::1.
func IPv6Loopback() Addr {
//line /snap/go/10455/src/net/netip/netip.go:82
	_go_fuzz_dep_.CoverTab[3894]++
//line /snap/go/10455/src/net/netip/netip.go:82
	return AddrFrom16([16]byte{15: 0x01})
//line /snap/go/10455/src/net/netip/netip.go:82
	// _ = "end of CoverTab[3894]"
//line /snap/go/10455/src/net/netip/netip.go:82
}

// IPv6Unspecified returns the IPv6 unspecified address "::".
func IPv6Unspecified() Addr {
//line /snap/go/10455/src/net/netip/netip.go:85
	_go_fuzz_dep_.CoverTab[3895]++
//line /snap/go/10455/src/net/netip/netip.go:85
	return Addr{z: z6noz}
//line /snap/go/10455/src/net/netip/netip.go:85
	// _ = "end of CoverTab[3895]"
//line /snap/go/10455/src/net/netip/netip.go:85
}

// IPv4Unspecified returns the IPv4 unspecified address "0.0.0.0".
func IPv4Unspecified() Addr {
//line /snap/go/10455/src/net/netip/netip.go:88
	_go_fuzz_dep_.CoverTab[3896]++
//line /snap/go/10455/src/net/netip/netip.go:88
	return AddrFrom4([4]byte{})
//line /snap/go/10455/src/net/netip/netip.go:88
	// _ = "end of CoverTab[3896]"
//line /snap/go/10455/src/net/netip/netip.go:88
}

// AddrFrom4 returns the address of the IPv4 address given by the bytes in addr.
func AddrFrom4(addr [4]byte) Addr {
//line /snap/go/10455/src/net/netip/netip.go:91
	_go_fuzz_dep_.CoverTab[3897]++
							return Addr{
		addr:	uint128{0, 0xffff00000000 | uint64(addr[0])<<24 | uint64(addr[1])<<16 | uint64(addr[2])<<8 | uint64(addr[3])},
		z:	z4,
	}
//line /snap/go/10455/src/net/netip/netip.go:95
	// _ = "end of CoverTab[3897]"
}

// AddrFrom16 returns the IPv6 address given by the bytes in addr.
//line /snap/go/10455/src/net/netip/netip.go:98
// An IPv4-mapped IPv6 address is left as an IPv6 address.
//line /snap/go/10455/src/net/netip/netip.go:98
// (Use Unmap to convert them if needed.)
//line /snap/go/10455/src/net/netip/netip.go:101
func AddrFrom16(addr [16]byte) Addr {
//line /snap/go/10455/src/net/netip/netip.go:101
	_go_fuzz_dep_.CoverTab[3898]++
							return Addr{
		addr: uint128{
			beUint64(addr[:8]),
			beUint64(addr[8:]),
		},
		z:	z6noz,
	}
//line /snap/go/10455/src/net/netip/netip.go:108
	// _ = "end of CoverTab[3898]"
}

// ParseAddr parses s as an IP address, returning the result. The string
//line /snap/go/10455/src/net/netip/netip.go:111
// s can be in dotted decimal ("192.0.2.1"), IPv6 ("2001:db8::68"),
//line /snap/go/10455/src/net/netip/netip.go:111
// or IPv6 with a scoped addressing zone ("fe80::1cc0:3e8c:119f:c2e1%ens18").
//line /snap/go/10455/src/net/netip/netip.go:114
func ParseAddr(s string) (Addr, error) {
//line /snap/go/10455/src/net/netip/netip.go:114
	_go_fuzz_dep_.CoverTab[3899]++
//line /snap/go/10455/src/net/netip/netip.go:114
	_go_fuzz_dep_.CoverTab[786625] = 0
							for i := 0; i < len(s); i++ {
//line /snap/go/10455/src/net/netip/netip.go:115
		if _go_fuzz_dep_.CoverTab[786625] == 0 {
//line /snap/go/10455/src/net/netip/netip.go:115
			_go_fuzz_dep_.CoverTab[527329]++
//line /snap/go/10455/src/net/netip/netip.go:115
		} else {
//line /snap/go/10455/src/net/netip/netip.go:115
			_go_fuzz_dep_.CoverTab[527330]++
//line /snap/go/10455/src/net/netip/netip.go:115
		}
//line /snap/go/10455/src/net/netip/netip.go:115
		_go_fuzz_dep_.CoverTab[786625] = 1
//line /snap/go/10455/src/net/netip/netip.go:115
		_go_fuzz_dep_.CoverTab[3901]++
								switch s[i] {
		case '.':
//line /snap/go/10455/src/net/netip/netip.go:117
			_go_fuzz_dep_.CoverTab[527018]++
//line /snap/go/10455/src/net/netip/netip.go:117
			_go_fuzz_dep_.CoverTab[3902]++
									return parseIPv4(s)
//line /snap/go/10455/src/net/netip/netip.go:118
			// _ = "end of CoverTab[3902]"
		case ':':
//line /snap/go/10455/src/net/netip/netip.go:119
			_go_fuzz_dep_.CoverTab[527019]++
//line /snap/go/10455/src/net/netip/netip.go:119
			_go_fuzz_dep_.CoverTab[3903]++
									return parseIPv6(s)
//line /snap/go/10455/src/net/netip/netip.go:120
			// _ = "end of CoverTab[3903]"
		case '%':
//line /snap/go/10455/src/net/netip/netip.go:121
			_go_fuzz_dep_.CoverTab[527020]++
//line /snap/go/10455/src/net/netip/netip.go:121
			_go_fuzz_dep_.CoverTab[3904]++

//line /snap/go/10455/src/net/netip/netip.go:124
			return Addr{}, parseAddrError{in: s, msg: "missing IPv6 address"}
//line /snap/go/10455/src/net/netip/netip.go:124
			// _ = "end of CoverTab[3904]"
//line /snap/go/10455/src/net/netip/netip.go:124
		default:
//line /snap/go/10455/src/net/netip/netip.go:124
			_go_fuzz_dep_.CoverTab[527021]++
//line /snap/go/10455/src/net/netip/netip.go:124
			_go_fuzz_dep_.CoverTab[3905]++
//line /snap/go/10455/src/net/netip/netip.go:124
			// _ = "end of CoverTab[3905]"
		}
//line /snap/go/10455/src/net/netip/netip.go:125
		// _ = "end of CoverTab[3901]"
	}
//line /snap/go/10455/src/net/netip/netip.go:126
	if _go_fuzz_dep_.CoverTab[786625] == 0 {
//line /snap/go/10455/src/net/netip/netip.go:126
		_go_fuzz_dep_.CoverTab[527331]++
//line /snap/go/10455/src/net/netip/netip.go:126
	} else {
//line /snap/go/10455/src/net/netip/netip.go:126
		_go_fuzz_dep_.CoverTab[527332]++
//line /snap/go/10455/src/net/netip/netip.go:126
	}
//line /snap/go/10455/src/net/netip/netip.go:126
	// _ = "end of CoverTab[3899]"
//line /snap/go/10455/src/net/netip/netip.go:126
	_go_fuzz_dep_.CoverTab[3900]++
							return Addr{}, parseAddrError{in: s, msg: "unable to parse IP"}
//line /snap/go/10455/src/net/netip/netip.go:127
	// _ = "end of CoverTab[3900]"
}

// MustParseAddr calls ParseAddr(s) and panics on error.
//line /snap/go/10455/src/net/netip/netip.go:130
// It is intended for use in tests with hard-coded strings.
//line /snap/go/10455/src/net/netip/netip.go:132
func MustParseAddr(s string) Addr {
//line /snap/go/10455/src/net/netip/netip.go:132
	_go_fuzz_dep_.CoverTab[3906]++
							ip, err := ParseAddr(s)
							if err != nil {
//line /snap/go/10455/src/net/netip/netip.go:134
		_go_fuzz_dep_.CoverTab[527022]++
//line /snap/go/10455/src/net/netip/netip.go:134
		_go_fuzz_dep_.CoverTab[3908]++
								panic(err)
//line /snap/go/10455/src/net/netip/netip.go:135
		// _ = "end of CoverTab[3908]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:136
		_go_fuzz_dep_.CoverTab[527023]++
//line /snap/go/10455/src/net/netip/netip.go:136
		_go_fuzz_dep_.CoverTab[3909]++
//line /snap/go/10455/src/net/netip/netip.go:136
		// _ = "end of CoverTab[3909]"
//line /snap/go/10455/src/net/netip/netip.go:136
	}
//line /snap/go/10455/src/net/netip/netip.go:136
	// _ = "end of CoverTab[3906]"
//line /snap/go/10455/src/net/netip/netip.go:136
	_go_fuzz_dep_.CoverTab[3907]++
							return ip
//line /snap/go/10455/src/net/netip/netip.go:137
	// _ = "end of CoverTab[3907]"
}

type parseAddrError struct {
	in	string	// the string given to ParseAddr
	msg	string	// an explanation of the parse failure
	at	string	// optionally, the unparsed portion of in at which the error occurred.
}

func (err parseAddrError) Error() string {
//line /snap/go/10455/src/net/netip/netip.go:146
	_go_fuzz_dep_.CoverTab[3910]++
							q := strconv.Quote
							if err.at != "" {
//line /snap/go/10455/src/net/netip/netip.go:148
		_go_fuzz_dep_.CoverTab[527024]++
//line /snap/go/10455/src/net/netip/netip.go:148
		_go_fuzz_dep_.CoverTab[3912]++
								return "ParseAddr(" + q(err.in) + "): " + err.msg + " (at " + q(err.at) + ")"
//line /snap/go/10455/src/net/netip/netip.go:149
		// _ = "end of CoverTab[3912]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:150
		_go_fuzz_dep_.CoverTab[527025]++
//line /snap/go/10455/src/net/netip/netip.go:150
		_go_fuzz_dep_.CoverTab[3913]++
//line /snap/go/10455/src/net/netip/netip.go:150
		// _ = "end of CoverTab[3913]"
//line /snap/go/10455/src/net/netip/netip.go:150
	}
//line /snap/go/10455/src/net/netip/netip.go:150
	// _ = "end of CoverTab[3910]"
//line /snap/go/10455/src/net/netip/netip.go:150
	_go_fuzz_dep_.CoverTab[3911]++
							return "ParseAddr(" + q(err.in) + "): " + err.msg
//line /snap/go/10455/src/net/netip/netip.go:151
	// _ = "end of CoverTab[3911]"
}

// parseIPv4 parses s as an IPv4 address (in form "192.168.0.1").
func parseIPv4(s string) (ip Addr, err error) {
//line /snap/go/10455/src/net/netip/netip.go:155
	_go_fuzz_dep_.CoverTab[3914]++
							var fields [4]uint8
							var val, pos int
							var digLen int
//line /snap/go/10455/src/net/netip/netip.go:158
	_go_fuzz_dep_. // number of digits in current octet
//line /snap/go/10455/src/net/netip/netip.go:158
	CoverTab[786626] = 0
							for i := 0; i < len(s); i++ {
//line /snap/go/10455/src/net/netip/netip.go:159
		if _go_fuzz_dep_.CoverTab[786626] == 0 {
//line /snap/go/10455/src/net/netip/netip.go:159
			_go_fuzz_dep_.CoverTab[527333]++
//line /snap/go/10455/src/net/netip/netip.go:159
		} else {
//line /snap/go/10455/src/net/netip/netip.go:159
			_go_fuzz_dep_.CoverTab[527334]++
//line /snap/go/10455/src/net/netip/netip.go:159
		}
//line /snap/go/10455/src/net/netip/netip.go:159
		_go_fuzz_dep_.CoverTab[786626] = 1
//line /snap/go/10455/src/net/netip/netip.go:159
		_go_fuzz_dep_.CoverTab[3917]++
								if s[i] >= '0' && func() bool {
//line /snap/go/10455/src/net/netip/netip.go:160
			_go_fuzz_dep_.CoverTab[3918]++
//line /snap/go/10455/src/net/netip/netip.go:160
			return s[i] <= '9'
//line /snap/go/10455/src/net/netip/netip.go:160
			// _ = "end of CoverTab[3918]"
//line /snap/go/10455/src/net/netip/netip.go:160
		}() {
//line /snap/go/10455/src/net/netip/netip.go:160
			_go_fuzz_dep_.CoverTab[527026]++
//line /snap/go/10455/src/net/netip/netip.go:160
			_go_fuzz_dep_.CoverTab[3919]++
									if digLen == 1 && func() bool {
//line /snap/go/10455/src/net/netip/netip.go:161
				_go_fuzz_dep_.CoverTab[3921]++
//line /snap/go/10455/src/net/netip/netip.go:161
				return val == 0
//line /snap/go/10455/src/net/netip/netip.go:161
				// _ = "end of CoverTab[3921]"
//line /snap/go/10455/src/net/netip/netip.go:161
			}() {
//line /snap/go/10455/src/net/netip/netip.go:161
				_go_fuzz_dep_.CoverTab[527028]++
//line /snap/go/10455/src/net/netip/netip.go:161
				_go_fuzz_dep_.CoverTab[3922]++
										return Addr{}, parseAddrError{in: s, msg: "IPv4 field has octet with leading zero"}
//line /snap/go/10455/src/net/netip/netip.go:162
				// _ = "end of CoverTab[3922]"
			} else {
//line /snap/go/10455/src/net/netip/netip.go:163
				_go_fuzz_dep_.CoverTab[527029]++
//line /snap/go/10455/src/net/netip/netip.go:163
				_go_fuzz_dep_.CoverTab[3923]++
//line /snap/go/10455/src/net/netip/netip.go:163
				// _ = "end of CoverTab[3923]"
//line /snap/go/10455/src/net/netip/netip.go:163
			}
//line /snap/go/10455/src/net/netip/netip.go:163
			// _ = "end of CoverTab[3919]"
//line /snap/go/10455/src/net/netip/netip.go:163
			_go_fuzz_dep_.CoverTab[3920]++
									val = val*10 + int(s[i]) - '0'
									digLen++
									if val > 255 {
//line /snap/go/10455/src/net/netip/netip.go:166
				_go_fuzz_dep_.CoverTab[527030]++
//line /snap/go/10455/src/net/netip/netip.go:166
				_go_fuzz_dep_.CoverTab[3924]++
										return Addr{}, parseAddrError{in: s, msg: "IPv4 field has value >255"}
//line /snap/go/10455/src/net/netip/netip.go:167
				// _ = "end of CoverTab[3924]"
			} else {
//line /snap/go/10455/src/net/netip/netip.go:168
				_go_fuzz_dep_.CoverTab[527031]++
//line /snap/go/10455/src/net/netip/netip.go:168
				_go_fuzz_dep_.CoverTab[3925]++
//line /snap/go/10455/src/net/netip/netip.go:168
				// _ = "end of CoverTab[3925]"
//line /snap/go/10455/src/net/netip/netip.go:168
			}
//line /snap/go/10455/src/net/netip/netip.go:168
			// _ = "end of CoverTab[3920]"
		} else {
//line /snap/go/10455/src/net/netip/netip.go:169
			_go_fuzz_dep_.CoverTab[527027]++
//line /snap/go/10455/src/net/netip/netip.go:169
			_go_fuzz_dep_.CoverTab[3926]++
//line /snap/go/10455/src/net/netip/netip.go:169
			if s[i] == '.' {
//line /snap/go/10455/src/net/netip/netip.go:169
				_go_fuzz_dep_.CoverTab[527032]++
//line /snap/go/10455/src/net/netip/netip.go:169
				_go_fuzz_dep_.CoverTab[3927]++

//line /snap/go/10455/src/net/netip/netip.go:173
				if i == 0 || func() bool {
//line /snap/go/10455/src/net/netip/netip.go:173
					_go_fuzz_dep_.CoverTab[3930]++
//line /snap/go/10455/src/net/netip/netip.go:173
					return i == len(s)-1
//line /snap/go/10455/src/net/netip/netip.go:173
					// _ = "end of CoverTab[3930]"
//line /snap/go/10455/src/net/netip/netip.go:173
				}() || func() bool {
//line /snap/go/10455/src/net/netip/netip.go:173
					_go_fuzz_dep_.CoverTab[3931]++
//line /snap/go/10455/src/net/netip/netip.go:173
					return s[i-1] == '.'
//line /snap/go/10455/src/net/netip/netip.go:173
					// _ = "end of CoverTab[3931]"
//line /snap/go/10455/src/net/netip/netip.go:173
				}() {
//line /snap/go/10455/src/net/netip/netip.go:173
					_go_fuzz_dep_.CoverTab[527034]++
//line /snap/go/10455/src/net/netip/netip.go:173
					_go_fuzz_dep_.CoverTab[3932]++
											return Addr{}, parseAddrError{in: s, msg: "IPv4 field must have at least one digit", at: s[i:]}
//line /snap/go/10455/src/net/netip/netip.go:174
					// _ = "end of CoverTab[3932]"
				} else {
//line /snap/go/10455/src/net/netip/netip.go:175
					_go_fuzz_dep_.CoverTab[527035]++
//line /snap/go/10455/src/net/netip/netip.go:175
					_go_fuzz_dep_.CoverTab[3933]++
//line /snap/go/10455/src/net/netip/netip.go:175
					// _ = "end of CoverTab[3933]"
//line /snap/go/10455/src/net/netip/netip.go:175
				}
//line /snap/go/10455/src/net/netip/netip.go:175
				// _ = "end of CoverTab[3927]"
//line /snap/go/10455/src/net/netip/netip.go:175
				_go_fuzz_dep_.CoverTab[3928]++

										if pos == 3 {
//line /snap/go/10455/src/net/netip/netip.go:177
					_go_fuzz_dep_.CoverTab[527036]++
//line /snap/go/10455/src/net/netip/netip.go:177
					_go_fuzz_dep_.CoverTab[3934]++
											return Addr{}, parseAddrError{in: s, msg: "IPv4 address too long"}
//line /snap/go/10455/src/net/netip/netip.go:178
					// _ = "end of CoverTab[3934]"
				} else {
//line /snap/go/10455/src/net/netip/netip.go:179
					_go_fuzz_dep_.CoverTab[527037]++
//line /snap/go/10455/src/net/netip/netip.go:179
					_go_fuzz_dep_.CoverTab[3935]++
//line /snap/go/10455/src/net/netip/netip.go:179
					// _ = "end of CoverTab[3935]"
//line /snap/go/10455/src/net/netip/netip.go:179
				}
//line /snap/go/10455/src/net/netip/netip.go:179
				// _ = "end of CoverTab[3928]"
//line /snap/go/10455/src/net/netip/netip.go:179
				_go_fuzz_dep_.CoverTab[3929]++
										fields[pos] = uint8(val)
										pos++
										val = 0
										digLen = 0
//line /snap/go/10455/src/net/netip/netip.go:183
				// _ = "end of CoverTab[3929]"
			} else {
//line /snap/go/10455/src/net/netip/netip.go:184
				_go_fuzz_dep_.CoverTab[527033]++
//line /snap/go/10455/src/net/netip/netip.go:184
				_go_fuzz_dep_.CoverTab[3936]++
										return Addr{}, parseAddrError{in: s, msg: "unexpected character", at: s[i:]}
//line /snap/go/10455/src/net/netip/netip.go:185
				// _ = "end of CoverTab[3936]"
			}
//line /snap/go/10455/src/net/netip/netip.go:186
			// _ = "end of CoverTab[3926]"
//line /snap/go/10455/src/net/netip/netip.go:186
		}
//line /snap/go/10455/src/net/netip/netip.go:186
		// _ = "end of CoverTab[3917]"
	}
//line /snap/go/10455/src/net/netip/netip.go:187
	if _go_fuzz_dep_.CoverTab[786626] == 0 {
//line /snap/go/10455/src/net/netip/netip.go:187
		_go_fuzz_dep_.CoverTab[527335]++
//line /snap/go/10455/src/net/netip/netip.go:187
	} else {
//line /snap/go/10455/src/net/netip/netip.go:187
		_go_fuzz_dep_.CoverTab[527336]++
//line /snap/go/10455/src/net/netip/netip.go:187
	}
//line /snap/go/10455/src/net/netip/netip.go:187
	// _ = "end of CoverTab[3914]"
//line /snap/go/10455/src/net/netip/netip.go:187
	_go_fuzz_dep_.CoverTab[3915]++
							if pos < 3 {
//line /snap/go/10455/src/net/netip/netip.go:188
		_go_fuzz_dep_.CoverTab[527038]++
//line /snap/go/10455/src/net/netip/netip.go:188
		_go_fuzz_dep_.CoverTab[3937]++
								return Addr{}, parseAddrError{in: s, msg: "IPv4 address too short"}
//line /snap/go/10455/src/net/netip/netip.go:189
		// _ = "end of CoverTab[3937]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:190
		_go_fuzz_dep_.CoverTab[527039]++
//line /snap/go/10455/src/net/netip/netip.go:190
		_go_fuzz_dep_.CoverTab[3938]++
//line /snap/go/10455/src/net/netip/netip.go:190
		// _ = "end of CoverTab[3938]"
//line /snap/go/10455/src/net/netip/netip.go:190
	}
//line /snap/go/10455/src/net/netip/netip.go:190
	// _ = "end of CoverTab[3915]"
//line /snap/go/10455/src/net/netip/netip.go:190
	_go_fuzz_dep_.CoverTab[3916]++
							fields[3] = uint8(val)
							return AddrFrom4(fields), nil
//line /snap/go/10455/src/net/netip/netip.go:192
	// _ = "end of CoverTab[3916]"
}

// parseIPv6 parses s as an IPv6 address (in form "2001:db8::68").
func parseIPv6(in string) (Addr, error) {
//line /snap/go/10455/src/net/netip/netip.go:196
	_go_fuzz_dep_.CoverTab[3939]++
							s := in

//line /snap/go/10455/src/net/netip/netip.go:203
	zone := ""
	i := bytealg.IndexByteString(s, '%')
	if i != -1 {
//line /snap/go/10455/src/net/netip/netip.go:205
		_go_fuzz_dep_.CoverTab[527040]++
//line /snap/go/10455/src/net/netip/netip.go:205
		_go_fuzz_dep_.CoverTab[3945]++
								s, zone = s[:i], s[i+1:]
								if zone == "" {
//line /snap/go/10455/src/net/netip/netip.go:207
			_go_fuzz_dep_.CoverTab[527042]++
//line /snap/go/10455/src/net/netip/netip.go:207
			_go_fuzz_dep_.CoverTab[3946]++

									return Addr{}, parseAddrError{in: in, msg: "zone must be a non-empty string"}
//line /snap/go/10455/src/net/netip/netip.go:209
			// _ = "end of CoverTab[3946]"
		} else {
//line /snap/go/10455/src/net/netip/netip.go:210
			_go_fuzz_dep_.CoverTab[527043]++
//line /snap/go/10455/src/net/netip/netip.go:210
			_go_fuzz_dep_.CoverTab[3947]++
//line /snap/go/10455/src/net/netip/netip.go:210
			// _ = "end of CoverTab[3947]"
//line /snap/go/10455/src/net/netip/netip.go:210
		}
//line /snap/go/10455/src/net/netip/netip.go:210
		// _ = "end of CoverTab[3945]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:211
		_go_fuzz_dep_.CoverTab[527041]++
//line /snap/go/10455/src/net/netip/netip.go:211
		_go_fuzz_dep_.CoverTab[3948]++
//line /snap/go/10455/src/net/netip/netip.go:211
		// _ = "end of CoverTab[3948]"
//line /snap/go/10455/src/net/netip/netip.go:211
	}
//line /snap/go/10455/src/net/netip/netip.go:211
	// _ = "end of CoverTab[3939]"
//line /snap/go/10455/src/net/netip/netip.go:211
	_go_fuzz_dep_.CoverTab[3940]++

							var ip [16]byte
							ellipsis := -1

//line /snap/go/10455/src/net/netip/netip.go:217
	if len(s) >= 2 && func() bool {
//line /snap/go/10455/src/net/netip/netip.go:217
		_go_fuzz_dep_.CoverTab[3949]++
//line /snap/go/10455/src/net/netip/netip.go:217
		return s[0] == ':'
//line /snap/go/10455/src/net/netip/netip.go:217
		// _ = "end of CoverTab[3949]"
//line /snap/go/10455/src/net/netip/netip.go:217
	}() && func() bool {
//line /snap/go/10455/src/net/netip/netip.go:217
		_go_fuzz_dep_.CoverTab[3950]++
//line /snap/go/10455/src/net/netip/netip.go:217
		return s[1] == ':'
//line /snap/go/10455/src/net/netip/netip.go:217
		// _ = "end of CoverTab[3950]"
//line /snap/go/10455/src/net/netip/netip.go:217
	}() {
//line /snap/go/10455/src/net/netip/netip.go:217
		_go_fuzz_dep_.CoverTab[527044]++
//line /snap/go/10455/src/net/netip/netip.go:217
		_go_fuzz_dep_.CoverTab[3951]++
								ellipsis = 0
								s = s[2:]

								if len(s) == 0 {
//line /snap/go/10455/src/net/netip/netip.go:221
			_go_fuzz_dep_.CoverTab[527046]++
//line /snap/go/10455/src/net/netip/netip.go:221
			_go_fuzz_dep_.CoverTab[3952]++
									return IPv6Unspecified().WithZone(zone), nil
//line /snap/go/10455/src/net/netip/netip.go:222
			// _ = "end of CoverTab[3952]"
		} else {
//line /snap/go/10455/src/net/netip/netip.go:223
			_go_fuzz_dep_.CoverTab[527047]++
//line /snap/go/10455/src/net/netip/netip.go:223
			_go_fuzz_dep_.CoverTab[3953]++
//line /snap/go/10455/src/net/netip/netip.go:223
			// _ = "end of CoverTab[3953]"
//line /snap/go/10455/src/net/netip/netip.go:223
		}
//line /snap/go/10455/src/net/netip/netip.go:223
		// _ = "end of CoverTab[3951]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:224
		_go_fuzz_dep_.CoverTab[527045]++
//line /snap/go/10455/src/net/netip/netip.go:224
		_go_fuzz_dep_.CoverTab[3954]++
//line /snap/go/10455/src/net/netip/netip.go:224
		// _ = "end of CoverTab[3954]"
//line /snap/go/10455/src/net/netip/netip.go:224
	}
//line /snap/go/10455/src/net/netip/netip.go:224
	// _ = "end of CoverTab[3940]"
//line /snap/go/10455/src/net/netip/netip.go:224
	_go_fuzz_dep_.CoverTab[3941]++

//line /snap/go/10455/src/net/netip/netip.go:227
	i = 0
//line /snap/go/10455/src/net/netip/netip.go:227
	_go_fuzz_dep_.CoverTab[786627] = 0
							for i < 16 {
//line /snap/go/10455/src/net/netip/netip.go:228
		if _go_fuzz_dep_.CoverTab[786627] == 0 {
//line /snap/go/10455/src/net/netip/netip.go:228
			_go_fuzz_dep_.CoverTab[527337]++
//line /snap/go/10455/src/net/netip/netip.go:228
		} else {
//line /snap/go/10455/src/net/netip/netip.go:228
			_go_fuzz_dep_.CoverTab[527338]++
//line /snap/go/10455/src/net/netip/netip.go:228
		}
//line /snap/go/10455/src/net/netip/netip.go:228
		_go_fuzz_dep_.CoverTab[786627] = 1
//line /snap/go/10455/src/net/netip/netip.go:228
		_go_fuzz_dep_.CoverTab[3955]++

//line /snap/go/10455/src/net/netip/netip.go:231
		off := 0
								acc := uint32(0)
//line /snap/go/10455/src/net/netip/netip.go:232
		_go_fuzz_dep_.CoverTab[786628] = 0
								for ; off < len(s); off++ {
//line /snap/go/10455/src/net/netip/netip.go:233
			if _go_fuzz_dep_.CoverTab[786628] == 0 {
//line /snap/go/10455/src/net/netip/netip.go:233
				_go_fuzz_dep_.CoverTab[527341]++
//line /snap/go/10455/src/net/netip/netip.go:233
			} else {
//line /snap/go/10455/src/net/netip/netip.go:233
				_go_fuzz_dep_.CoverTab[527342]++
//line /snap/go/10455/src/net/netip/netip.go:233
			}
//line /snap/go/10455/src/net/netip/netip.go:233
			_go_fuzz_dep_.CoverTab[786628] = 1
//line /snap/go/10455/src/net/netip/netip.go:233
			_go_fuzz_dep_.CoverTab[3961]++
									c := s[off]
									if c >= '0' && func() bool {
//line /snap/go/10455/src/net/netip/netip.go:235
				_go_fuzz_dep_.CoverTab[3963]++
//line /snap/go/10455/src/net/netip/netip.go:235
				return c <= '9'
//line /snap/go/10455/src/net/netip/netip.go:235
				// _ = "end of CoverTab[3963]"
//line /snap/go/10455/src/net/netip/netip.go:235
			}() {
//line /snap/go/10455/src/net/netip/netip.go:235
				_go_fuzz_dep_.CoverTab[527048]++
//line /snap/go/10455/src/net/netip/netip.go:235
				_go_fuzz_dep_.CoverTab[3964]++
										acc = (acc << 4) + uint32(c-'0')
//line /snap/go/10455/src/net/netip/netip.go:236
				// _ = "end of CoverTab[3964]"
			} else {
//line /snap/go/10455/src/net/netip/netip.go:237
				_go_fuzz_dep_.CoverTab[527049]++
//line /snap/go/10455/src/net/netip/netip.go:237
				_go_fuzz_dep_.CoverTab[3965]++
//line /snap/go/10455/src/net/netip/netip.go:237
				if c >= 'a' && func() bool {
//line /snap/go/10455/src/net/netip/netip.go:237
					_go_fuzz_dep_.CoverTab[3966]++
//line /snap/go/10455/src/net/netip/netip.go:237
					return c <= 'f'
//line /snap/go/10455/src/net/netip/netip.go:237
					// _ = "end of CoverTab[3966]"
//line /snap/go/10455/src/net/netip/netip.go:237
				}() {
//line /snap/go/10455/src/net/netip/netip.go:237
					_go_fuzz_dep_.CoverTab[527050]++
//line /snap/go/10455/src/net/netip/netip.go:237
					_go_fuzz_dep_.CoverTab[3967]++
											acc = (acc << 4) + uint32(c-'a'+10)
//line /snap/go/10455/src/net/netip/netip.go:238
					// _ = "end of CoverTab[3967]"
				} else {
//line /snap/go/10455/src/net/netip/netip.go:239
					_go_fuzz_dep_.CoverTab[527051]++
//line /snap/go/10455/src/net/netip/netip.go:239
					_go_fuzz_dep_.CoverTab[3968]++
//line /snap/go/10455/src/net/netip/netip.go:239
					if c >= 'A' && func() bool {
//line /snap/go/10455/src/net/netip/netip.go:239
						_go_fuzz_dep_.CoverTab[3969]++
//line /snap/go/10455/src/net/netip/netip.go:239
						return c <= 'F'
//line /snap/go/10455/src/net/netip/netip.go:239
						// _ = "end of CoverTab[3969]"
//line /snap/go/10455/src/net/netip/netip.go:239
					}() {
//line /snap/go/10455/src/net/netip/netip.go:239
						_go_fuzz_dep_.CoverTab[527052]++
//line /snap/go/10455/src/net/netip/netip.go:239
						_go_fuzz_dep_.CoverTab[3970]++
												acc = (acc << 4) + uint32(c-'A'+10)
//line /snap/go/10455/src/net/netip/netip.go:240
						// _ = "end of CoverTab[3970]"
					} else {
//line /snap/go/10455/src/net/netip/netip.go:241
						_go_fuzz_dep_.CoverTab[527053]++
//line /snap/go/10455/src/net/netip/netip.go:241
						_go_fuzz_dep_.CoverTab[3971]++
												break
//line /snap/go/10455/src/net/netip/netip.go:242
						// _ = "end of CoverTab[3971]"
					}
//line /snap/go/10455/src/net/netip/netip.go:243
					// _ = "end of CoverTab[3968]"
//line /snap/go/10455/src/net/netip/netip.go:243
				}
//line /snap/go/10455/src/net/netip/netip.go:243
				// _ = "end of CoverTab[3965]"
//line /snap/go/10455/src/net/netip/netip.go:243
			}
//line /snap/go/10455/src/net/netip/netip.go:243
			// _ = "end of CoverTab[3961]"
//line /snap/go/10455/src/net/netip/netip.go:243
			_go_fuzz_dep_.CoverTab[3962]++
									if acc > math.MaxUint16 {
//line /snap/go/10455/src/net/netip/netip.go:244
				_go_fuzz_dep_.CoverTab[527054]++
//line /snap/go/10455/src/net/netip/netip.go:244
				_go_fuzz_dep_.CoverTab[3972]++

										return Addr{}, parseAddrError{in: in, msg: "IPv6 field has value >=2^16", at: s}
//line /snap/go/10455/src/net/netip/netip.go:246
				// _ = "end of CoverTab[3972]"
			} else {
//line /snap/go/10455/src/net/netip/netip.go:247
				_go_fuzz_dep_.CoverTab[527055]++
//line /snap/go/10455/src/net/netip/netip.go:247
				_go_fuzz_dep_.CoverTab[3973]++
//line /snap/go/10455/src/net/netip/netip.go:247
				// _ = "end of CoverTab[3973]"
//line /snap/go/10455/src/net/netip/netip.go:247
			}
//line /snap/go/10455/src/net/netip/netip.go:247
			// _ = "end of CoverTab[3962]"
		}
//line /snap/go/10455/src/net/netip/netip.go:248
		if _go_fuzz_dep_.CoverTab[786628] == 0 {
//line /snap/go/10455/src/net/netip/netip.go:248
			_go_fuzz_dep_.CoverTab[527343]++
//line /snap/go/10455/src/net/netip/netip.go:248
		} else {
//line /snap/go/10455/src/net/netip/netip.go:248
			_go_fuzz_dep_.CoverTab[527344]++
//line /snap/go/10455/src/net/netip/netip.go:248
		}
//line /snap/go/10455/src/net/netip/netip.go:248
		// _ = "end of CoverTab[3955]"
//line /snap/go/10455/src/net/netip/netip.go:248
		_go_fuzz_dep_.CoverTab[3956]++
								if off == 0 {
//line /snap/go/10455/src/net/netip/netip.go:249
			_go_fuzz_dep_.CoverTab[527056]++
//line /snap/go/10455/src/net/netip/netip.go:249
			_go_fuzz_dep_.CoverTab[3974]++

									return Addr{}, parseAddrError{in: in, msg: "each colon-separated field must have at least one digit", at: s}
//line /snap/go/10455/src/net/netip/netip.go:251
			// _ = "end of CoverTab[3974]"
		} else {
//line /snap/go/10455/src/net/netip/netip.go:252
			_go_fuzz_dep_.CoverTab[527057]++
//line /snap/go/10455/src/net/netip/netip.go:252
			_go_fuzz_dep_.CoverTab[3975]++
//line /snap/go/10455/src/net/netip/netip.go:252
			// _ = "end of CoverTab[3975]"
//line /snap/go/10455/src/net/netip/netip.go:252
		}
//line /snap/go/10455/src/net/netip/netip.go:252
		// _ = "end of CoverTab[3956]"
//line /snap/go/10455/src/net/netip/netip.go:252
		_go_fuzz_dep_.CoverTab[3957]++

//line /snap/go/10455/src/net/netip/netip.go:255
		if off < len(s) && func() bool {
//line /snap/go/10455/src/net/netip/netip.go:255
			_go_fuzz_dep_.CoverTab[3976]++
//line /snap/go/10455/src/net/netip/netip.go:255
			return s[off] == '.'
//line /snap/go/10455/src/net/netip/netip.go:255
			// _ = "end of CoverTab[3976]"
//line /snap/go/10455/src/net/netip/netip.go:255
		}() {
//line /snap/go/10455/src/net/netip/netip.go:255
			_go_fuzz_dep_.CoverTab[527058]++
//line /snap/go/10455/src/net/netip/netip.go:255
			_go_fuzz_dep_.CoverTab[3977]++
									if ellipsis < 0 && func() bool {
//line /snap/go/10455/src/net/netip/netip.go:256
				_go_fuzz_dep_.CoverTab[3981]++
//line /snap/go/10455/src/net/netip/netip.go:256
				return i != 12
//line /snap/go/10455/src/net/netip/netip.go:256
				// _ = "end of CoverTab[3981]"
//line /snap/go/10455/src/net/netip/netip.go:256
			}() {
//line /snap/go/10455/src/net/netip/netip.go:256
				_go_fuzz_dep_.CoverTab[527060]++
//line /snap/go/10455/src/net/netip/netip.go:256
				_go_fuzz_dep_.CoverTab[3982]++

										return Addr{}, parseAddrError{in: in, msg: "embedded IPv4 address must replace the final 2 fields of the address", at: s}
//line /snap/go/10455/src/net/netip/netip.go:258
				// _ = "end of CoverTab[3982]"
			} else {
//line /snap/go/10455/src/net/netip/netip.go:259
				_go_fuzz_dep_.CoverTab[527061]++
//line /snap/go/10455/src/net/netip/netip.go:259
				_go_fuzz_dep_.CoverTab[3983]++
//line /snap/go/10455/src/net/netip/netip.go:259
				// _ = "end of CoverTab[3983]"
//line /snap/go/10455/src/net/netip/netip.go:259
			}
//line /snap/go/10455/src/net/netip/netip.go:259
			// _ = "end of CoverTab[3977]"
//line /snap/go/10455/src/net/netip/netip.go:259
			_go_fuzz_dep_.CoverTab[3978]++
									if i+4 > 16 {
//line /snap/go/10455/src/net/netip/netip.go:260
				_go_fuzz_dep_.CoverTab[527062]++
//line /snap/go/10455/src/net/netip/netip.go:260
				_go_fuzz_dep_.CoverTab[3984]++

										return Addr{}, parseAddrError{in: in, msg: "too many hex fields to fit an embedded IPv4 at the end of the address", at: s}
//line /snap/go/10455/src/net/netip/netip.go:262
				// _ = "end of CoverTab[3984]"
			} else {
//line /snap/go/10455/src/net/netip/netip.go:263
				_go_fuzz_dep_.CoverTab[527063]++
//line /snap/go/10455/src/net/netip/netip.go:263
				_go_fuzz_dep_.CoverTab[3985]++
//line /snap/go/10455/src/net/netip/netip.go:263
				// _ = "end of CoverTab[3985]"
//line /snap/go/10455/src/net/netip/netip.go:263
			}
//line /snap/go/10455/src/net/netip/netip.go:263
			// _ = "end of CoverTab[3978]"
//line /snap/go/10455/src/net/netip/netip.go:263
			_go_fuzz_dep_.CoverTab[3979]++

//line /snap/go/10455/src/net/netip/netip.go:267
			ip4, err := parseIPv4(s)
			if err != nil {
//line /snap/go/10455/src/net/netip/netip.go:268
				_go_fuzz_dep_.CoverTab[527064]++
//line /snap/go/10455/src/net/netip/netip.go:268
				_go_fuzz_dep_.CoverTab[3986]++
										return Addr{}, parseAddrError{in: in, msg: err.Error(), at: s}
//line /snap/go/10455/src/net/netip/netip.go:269
				// _ = "end of CoverTab[3986]"
			} else {
//line /snap/go/10455/src/net/netip/netip.go:270
				_go_fuzz_dep_.CoverTab[527065]++
//line /snap/go/10455/src/net/netip/netip.go:270
				_go_fuzz_dep_.CoverTab[3987]++
//line /snap/go/10455/src/net/netip/netip.go:270
				// _ = "end of CoverTab[3987]"
//line /snap/go/10455/src/net/netip/netip.go:270
			}
//line /snap/go/10455/src/net/netip/netip.go:270
			// _ = "end of CoverTab[3979]"
//line /snap/go/10455/src/net/netip/netip.go:270
			_go_fuzz_dep_.CoverTab[3980]++
									ip[i] = ip4.v4(0)
									ip[i+1] = ip4.v4(1)
									ip[i+2] = ip4.v4(2)
									ip[i+3] = ip4.v4(3)
									s = ""
									i += 4
									break
//line /snap/go/10455/src/net/netip/netip.go:277
			// _ = "end of CoverTab[3980]"
		} else {
//line /snap/go/10455/src/net/netip/netip.go:278
			_go_fuzz_dep_.CoverTab[527059]++
//line /snap/go/10455/src/net/netip/netip.go:278
			_go_fuzz_dep_.CoverTab[3988]++
//line /snap/go/10455/src/net/netip/netip.go:278
			// _ = "end of CoverTab[3988]"
//line /snap/go/10455/src/net/netip/netip.go:278
		}
//line /snap/go/10455/src/net/netip/netip.go:278
		// _ = "end of CoverTab[3957]"
//line /snap/go/10455/src/net/netip/netip.go:278
		_go_fuzz_dep_.CoverTab[3958]++

//line /snap/go/10455/src/net/netip/netip.go:281
		ip[i] = byte(acc >> 8)
								ip[i+1] = byte(acc)
								i += 2

//line /snap/go/10455/src/net/netip/netip.go:286
		s = s[off:]
		if len(s) == 0 {
//line /snap/go/10455/src/net/netip/netip.go:287
			_go_fuzz_dep_.CoverTab[527066]++
//line /snap/go/10455/src/net/netip/netip.go:287
			_go_fuzz_dep_.CoverTab[3989]++
									break
//line /snap/go/10455/src/net/netip/netip.go:288
			// _ = "end of CoverTab[3989]"
		} else {
//line /snap/go/10455/src/net/netip/netip.go:289
			_go_fuzz_dep_.CoverTab[527067]++
//line /snap/go/10455/src/net/netip/netip.go:289
			_go_fuzz_dep_.CoverTab[3990]++
//line /snap/go/10455/src/net/netip/netip.go:289
			// _ = "end of CoverTab[3990]"
//line /snap/go/10455/src/net/netip/netip.go:289
		}
//line /snap/go/10455/src/net/netip/netip.go:289
		// _ = "end of CoverTab[3958]"
//line /snap/go/10455/src/net/netip/netip.go:289
		_go_fuzz_dep_.CoverTab[3959]++

//line /snap/go/10455/src/net/netip/netip.go:292
		if s[0] != ':' {
//line /snap/go/10455/src/net/netip/netip.go:292
			_go_fuzz_dep_.CoverTab[527068]++
//line /snap/go/10455/src/net/netip/netip.go:292
			_go_fuzz_dep_.CoverTab[3991]++
									return Addr{}, parseAddrError{in: in, msg: "unexpected character, want colon", at: s}
//line /snap/go/10455/src/net/netip/netip.go:293
			// _ = "end of CoverTab[3991]"
		} else {
//line /snap/go/10455/src/net/netip/netip.go:294
			_go_fuzz_dep_.CoverTab[527069]++
//line /snap/go/10455/src/net/netip/netip.go:294
			_go_fuzz_dep_.CoverTab[3992]++
//line /snap/go/10455/src/net/netip/netip.go:294
			if len(s) == 1 {
//line /snap/go/10455/src/net/netip/netip.go:294
				_go_fuzz_dep_.CoverTab[527070]++
//line /snap/go/10455/src/net/netip/netip.go:294
				_go_fuzz_dep_.CoverTab[3993]++
										return Addr{}, parseAddrError{in: in, msg: "colon must be followed by more characters", at: s}
//line /snap/go/10455/src/net/netip/netip.go:295
				// _ = "end of CoverTab[3993]"
			} else {
//line /snap/go/10455/src/net/netip/netip.go:296
				_go_fuzz_dep_.CoverTab[527071]++
//line /snap/go/10455/src/net/netip/netip.go:296
				_go_fuzz_dep_.CoverTab[3994]++
//line /snap/go/10455/src/net/netip/netip.go:296
				// _ = "end of CoverTab[3994]"
//line /snap/go/10455/src/net/netip/netip.go:296
			}
//line /snap/go/10455/src/net/netip/netip.go:296
			// _ = "end of CoverTab[3992]"
//line /snap/go/10455/src/net/netip/netip.go:296
		}
//line /snap/go/10455/src/net/netip/netip.go:296
		// _ = "end of CoverTab[3959]"
//line /snap/go/10455/src/net/netip/netip.go:296
		_go_fuzz_dep_.CoverTab[3960]++
								s = s[1:]

//line /snap/go/10455/src/net/netip/netip.go:300
		if s[0] == ':' {
//line /snap/go/10455/src/net/netip/netip.go:300
			_go_fuzz_dep_.CoverTab[527072]++
//line /snap/go/10455/src/net/netip/netip.go:300
			_go_fuzz_dep_.CoverTab[3995]++
									if ellipsis >= 0 {
//line /snap/go/10455/src/net/netip/netip.go:301
				_go_fuzz_dep_.CoverTab[527074]++
//line /snap/go/10455/src/net/netip/netip.go:301
				_go_fuzz_dep_.CoverTab[3997]++
										return Addr{}, parseAddrError{in: in, msg: "multiple :: in address", at: s}
//line /snap/go/10455/src/net/netip/netip.go:302
				// _ = "end of CoverTab[3997]"
			} else {
//line /snap/go/10455/src/net/netip/netip.go:303
				_go_fuzz_dep_.CoverTab[527075]++
//line /snap/go/10455/src/net/netip/netip.go:303
				_go_fuzz_dep_.CoverTab[3998]++
//line /snap/go/10455/src/net/netip/netip.go:303
				// _ = "end of CoverTab[3998]"
//line /snap/go/10455/src/net/netip/netip.go:303
			}
//line /snap/go/10455/src/net/netip/netip.go:303
			// _ = "end of CoverTab[3995]"
//line /snap/go/10455/src/net/netip/netip.go:303
			_go_fuzz_dep_.CoverTab[3996]++
									ellipsis = i
									s = s[1:]
									if len(s) == 0 {
//line /snap/go/10455/src/net/netip/netip.go:306
				_go_fuzz_dep_.CoverTab[527076]++
//line /snap/go/10455/src/net/netip/netip.go:306
				_go_fuzz_dep_.CoverTab[3999]++
										break
//line /snap/go/10455/src/net/netip/netip.go:307
				// _ = "end of CoverTab[3999]"
			} else {
//line /snap/go/10455/src/net/netip/netip.go:308
				_go_fuzz_dep_.CoverTab[527077]++
//line /snap/go/10455/src/net/netip/netip.go:308
				_go_fuzz_dep_.CoverTab[4000]++
//line /snap/go/10455/src/net/netip/netip.go:308
				// _ = "end of CoverTab[4000]"
//line /snap/go/10455/src/net/netip/netip.go:308
			}
//line /snap/go/10455/src/net/netip/netip.go:308
			// _ = "end of CoverTab[3996]"
		} else {
//line /snap/go/10455/src/net/netip/netip.go:309
			_go_fuzz_dep_.CoverTab[527073]++
//line /snap/go/10455/src/net/netip/netip.go:309
			_go_fuzz_dep_.CoverTab[4001]++
//line /snap/go/10455/src/net/netip/netip.go:309
			// _ = "end of CoverTab[4001]"
//line /snap/go/10455/src/net/netip/netip.go:309
		}
//line /snap/go/10455/src/net/netip/netip.go:309
		// _ = "end of CoverTab[3960]"
	}
//line /snap/go/10455/src/net/netip/netip.go:310
	if _go_fuzz_dep_.CoverTab[786627] == 0 {
//line /snap/go/10455/src/net/netip/netip.go:310
		_go_fuzz_dep_.CoverTab[527339]++
//line /snap/go/10455/src/net/netip/netip.go:310
	} else {
//line /snap/go/10455/src/net/netip/netip.go:310
		_go_fuzz_dep_.CoverTab[527340]++
//line /snap/go/10455/src/net/netip/netip.go:310
	}
//line /snap/go/10455/src/net/netip/netip.go:310
	// _ = "end of CoverTab[3941]"
//line /snap/go/10455/src/net/netip/netip.go:310
	_go_fuzz_dep_.CoverTab[3942]++

//line /snap/go/10455/src/net/netip/netip.go:313
	if len(s) != 0 {
//line /snap/go/10455/src/net/netip/netip.go:313
		_go_fuzz_dep_.CoverTab[527078]++
//line /snap/go/10455/src/net/netip/netip.go:313
		_go_fuzz_dep_.CoverTab[4002]++
								return Addr{}, parseAddrError{in: in, msg: "trailing garbage after address", at: s}
//line /snap/go/10455/src/net/netip/netip.go:314
		// _ = "end of CoverTab[4002]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:315
		_go_fuzz_dep_.CoverTab[527079]++
//line /snap/go/10455/src/net/netip/netip.go:315
		_go_fuzz_dep_.CoverTab[4003]++
//line /snap/go/10455/src/net/netip/netip.go:315
		// _ = "end of CoverTab[4003]"
//line /snap/go/10455/src/net/netip/netip.go:315
	}
//line /snap/go/10455/src/net/netip/netip.go:315
	// _ = "end of CoverTab[3942]"
//line /snap/go/10455/src/net/netip/netip.go:315
	_go_fuzz_dep_.CoverTab[3943]++

//line /snap/go/10455/src/net/netip/netip.go:318
	if i < 16 {
//line /snap/go/10455/src/net/netip/netip.go:318
		_go_fuzz_dep_.CoverTab[527080]++
//line /snap/go/10455/src/net/netip/netip.go:318
		_go_fuzz_dep_.CoverTab[4004]++
								if ellipsis < 0 {
//line /snap/go/10455/src/net/netip/netip.go:319
			_go_fuzz_dep_.CoverTab[527082]++
//line /snap/go/10455/src/net/netip/netip.go:319
			_go_fuzz_dep_.CoverTab[4007]++
									return Addr{}, parseAddrError{in: in, msg: "address string too short"}
//line /snap/go/10455/src/net/netip/netip.go:320
			// _ = "end of CoverTab[4007]"
		} else {
//line /snap/go/10455/src/net/netip/netip.go:321
			_go_fuzz_dep_.CoverTab[527083]++
//line /snap/go/10455/src/net/netip/netip.go:321
			_go_fuzz_dep_.CoverTab[4008]++
//line /snap/go/10455/src/net/netip/netip.go:321
			// _ = "end of CoverTab[4008]"
//line /snap/go/10455/src/net/netip/netip.go:321
		}
//line /snap/go/10455/src/net/netip/netip.go:321
		// _ = "end of CoverTab[4004]"
//line /snap/go/10455/src/net/netip/netip.go:321
		_go_fuzz_dep_.CoverTab[4005]++
								n := 16 - i
//line /snap/go/10455/src/net/netip/netip.go:322
		_go_fuzz_dep_.CoverTab[786629] = 0
								for j := i - 1; j >= ellipsis; j-- {
//line /snap/go/10455/src/net/netip/netip.go:323
			if _go_fuzz_dep_.CoverTab[786629] == 0 {
//line /snap/go/10455/src/net/netip/netip.go:323
				_go_fuzz_dep_.CoverTab[527345]++
//line /snap/go/10455/src/net/netip/netip.go:323
			} else {
//line /snap/go/10455/src/net/netip/netip.go:323
				_go_fuzz_dep_.CoverTab[527346]++
//line /snap/go/10455/src/net/netip/netip.go:323
			}
//line /snap/go/10455/src/net/netip/netip.go:323
			_go_fuzz_dep_.CoverTab[786629] = 1
//line /snap/go/10455/src/net/netip/netip.go:323
			_go_fuzz_dep_.CoverTab[4009]++
									ip[j+n] = ip[j]
//line /snap/go/10455/src/net/netip/netip.go:324
			// _ = "end of CoverTab[4009]"
		}
//line /snap/go/10455/src/net/netip/netip.go:325
		if _go_fuzz_dep_.CoverTab[786629] == 0 {
//line /snap/go/10455/src/net/netip/netip.go:325
			_go_fuzz_dep_.CoverTab[527347]++
//line /snap/go/10455/src/net/netip/netip.go:325
		} else {
//line /snap/go/10455/src/net/netip/netip.go:325
			_go_fuzz_dep_.CoverTab[527348]++
//line /snap/go/10455/src/net/netip/netip.go:325
		}
//line /snap/go/10455/src/net/netip/netip.go:325
		// _ = "end of CoverTab[4005]"
//line /snap/go/10455/src/net/netip/netip.go:325
		_go_fuzz_dep_.CoverTab[4006]++
//line /snap/go/10455/src/net/netip/netip.go:325
		_go_fuzz_dep_.CoverTab[786630] = 0
								for j := ellipsis + n - 1; j >= ellipsis; j-- {
//line /snap/go/10455/src/net/netip/netip.go:326
			if _go_fuzz_dep_.CoverTab[786630] == 0 {
//line /snap/go/10455/src/net/netip/netip.go:326
				_go_fuzz_dep_.CoverTab[527349]++
//line /snap/go/10455/src/net/netip/netip.go:326
			} else {
//line /snap/go/10455/src/net/netip/netip.go:326
				_go_fuzz_dep_.CoverTab[527350]++
//line /snap/go/10455/src/net/netip/netip.go:326
			}
//line /snap/go/10455/src/net/netip/netip.go:326
			_go_fuzz_dep_.CoverTab[786630] = 1
//line /snap/go/10455/src/net/netip/netip.go:326
			_go_fuzz_dep_.CoverTab[4010]++
									ip[j] = 0
//line /snap/go/10455/src/net/netip/netip.go:327
			// _ = "end of CoverTab[4010]"
		}
//line /snap/go/10455/src/net/netip/netip.go:328
		if _go_fuzz_dep_.CoverTab[786630] == 0 {
//line /snap/go/10455/src/net/netip/netip.go:328
			_go_fuzz_dep_.CoverTab[527351]++
//line /snap/go/10455/src/net/netip/netip.go:328
		} else {
//line /snap/go/10455/src/net/netip/netip.go:328
			_go_fuzz_dep_.CoverTab[527352]++
//line /snap/go/10455/src/net/netip/netip.go:328
		}
//line /snap/go/10455/src/net/netip/netip.go:328
		// _ = "end of CoverTab[4006]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:329
		_go_fuzz_dep_.CoverTab[527081]++
//line /snap/go/10455/src/net/netip/netip.go:329
		_go_fuzz_dep_.CoverTab[4011]++
//line /snap/go/10455/src/net/netip/netip.go:329
		if ellipsis >= 0 {
//line /snap/go/10455/src/net/netip/netip.go:329
			_go_fuzz_dep_.CoverTab[527084]++
//line /snap/go/10455/src/net/netip/netip.go:329
			_go_fuzz_dep_.CoverTab[4012]++

									return Addr{}, parseAddrError{in: in, msg: "the :: must expand to at least one field of zeros"}
//line /snap/go/10455/src/net/netip/netip.go:331
			// _ = "end of CoverTab[4012]"
		} else {
//line /snap/go/10455/src/net/netip/netip.go:332
			_go_fuzz_dep_.CoverTab[527085]++
//line /snap/go/10455/src/net/netip/netip.go:332
			_go_fuzz_dep_.CoverTab[4013]++
//line /snap/go/10455/src/net/netip/netip.go:332
			// _ = "end of CoverTab[4013]"
//line /snap/go/10455/src/net/netip/netip.go:332
		}
//line /snap/go/10455/src/net/netip/netip.go:332
		// _ = "end of CoverTab[4011]"
//line /snap/go/10455/src/net/netip/netip.go:332
	}
//line /snap/go/10455/src/net/netip/netip.go:332
	// _ = "end of CoverTab[3943]"
//line /snap/go/10455/src/net/netip/netip.go:332
	_go_fuzz_dep_.CoverTab[3944]++
							return AddrFrom16(ip).WithZone(zone), nil
//line /snap/go/10455/src/net/netip/netip.go:333
	// _ = "end of CoverTab[3944]"
}

// AddrFromSlice parses the 4- or 16-byte byte slice as an IPv4 or IPv6 address.
//line /snap/go/10455/src/net/netip/netip.go:336
// Note that a net.IP can be passed directly as the []byte argument.
//line /snap/go/10455/src/net/netip/netip.go:336
// If slice's length is not 4 or 16, AddrFromSlice returns Addr{}, false.
//line /snap/go/10455/src/net/netip/netip.go:339
func AddrFromSlice(slice []byte) (ip Addr, ok bool) {
//line /snap/go/10455/src/net/netip/netip.go:339
	_go_fuzz_dep_.CoverTab[4014]++
							switch len(slice) {
	case 4:
//line /snap/go/10455/src/net/netip/netip.go:341
		_go_fuzz_dep_.CoverTab[527086]++
//line /snap/go/10455/src/net/netip/netip.go:341
		_go_fuzz_dep_.CoverTab[4016]++
								return AddrFrom4([4]byte(slice)), true
//line /snap/go/10455/src/net/netip/netip.go:342
		// _ = "end of CoverTab[4016]"
	case 16:
//line /snap/go/10455/src/net/netip/netip.go:343
		_go_fuzz_dep_.CoverTab[527087]++
//line /snap/go/10455/src/net/netip/netip.go:343
		_go_fuzz_dep_.CoverTab[4017]++
								return AddrFrom16([16]byte(slice)), true
//line /snap/go/10455/src/net/netip/netip.go:344
		// _ = "end of CoverTab[4017]"
//line /snap/go/10455/src/net/netip/netip.go:344
	default:
//line /snap/go/10455/src/net/netip/netip.go:344
		_go_fuzz_dep_.CoverTab[527088]++
//line /snap/go/10455/src/net/netip/netip.go:344
		_go_fuzz_dep_.CoverTab[4018]++
//line /snap/go/10455/src/net/netip/netip.go:344
		// _ = "end of CoverTab[4018]"
	}
//line /snap/go/10455/src/net/netip/netip.go:345
	// _ = "end of CoverTab[4014]"
//line /snap/go/10455/src/net/netip/netip.go:345
	_go_fuzz_dep_.CoverTab[4015]++
							return Addr{}, false
//line /snap/go/10455/src/net/netip/netip.go:346
	// _ = "end of CoverTab[4015]"
}

// v4 returns the i'th byte of ip. If ip is not an IPv4, v4 returns
//line /snap/go/10455/src/net/netip/netip.go:349
// unspecified garbage.
//line /snap/go/10455/src/net/netip/netip.go:351
func (ip Addr) v4(i uint8) uint8 {
//line /snap/go/10455/src/net/netip/netip.go:351
	_go_fuzz_dep_.CoverTab[4019]++
							return uint8(ip.addr.lo >> ((3 - i) * 8))
//line /snap/go/10455/src/net/netip/netip.go:352
	// _ = "end of CoverTab[4019]"
}

// v6 returns the i'th byte of ip. If ip is an IPv4 address, this
//line /snap/go/10455/src/net/netip/netip.go:355
// accesses the IPv4-mapped IPv6 address form of the IP.
//line /snap/go/10455/src/net/netip/netip.go:357
func (ip Addr) v6(i uint8) uint8 {
//line /snap/go/10455/src/net/netip/netip.go:357
	_go_fuzz_dep_.CoverTab[4020]++
							return uint8(*(ip.addr.halves()[(i/8)%2]) >> ((7 - i%8) * 8))
//line /snap/go/10455/src/net/netip/netip.go:358
	// _ = "end of CoverTab[4020]"
}

// v6u16 returns the i'th 16-bit word of ip. If ip is an IPv4 address,
//line /snap/go/10455/src/net/netip/netip.go:361
// this accesses the IPv4-mapped IPv6 address form of the IP.
//line /snap/go/10455/src/net/netip/netip.go:363
func (ip Addr) v6u16(i uint8) uint16 {
//line /snap/go/10455/src/net/netip/netip.go:363
	_go_fuzz_dep_.CoverTab[4021]++
							return uint16(*(ip.addr.halves()[(i/4)%2]) >> ((3 - i%4) * 16))
//line /snap/go/10455/src/net/netip/netip.go:364
	// _ = "end of CoverTab[4021]"
}

// isZero reports whether ip is the zero value of the IP type.
//line /snap/go/10455/src/net/netip/netip.go:367
// The zero value is not a valid IP address of any type.
//line /snap/go/10455/src/net/netip/netip.go:367
//
//line /snap/go/10455/src/net/netip/netip.go:367
// Note that "0.0.0.0" and "::" are not the zero value. Use IsUnspecified to
//line /snap/go/10455/src/net/netip/netip.go:367
// check for these values instead.
//line /snap/go/10455/src/net/netip/netip.go:372
func (ip Addr) isZero() bool {
//line /snap/go/10455/src/net/netip/netip.go:372
	_go_fuzz_dep_.CoverTab[4022]++

//line /snap/go/10455/src/net/netip/netip.go:375
	return ip.z == z0
//line /snap/go/10455/src/net/netip/netip.go:375
	// _ = "end of CoverTab[4022]"
}

// IsValid reports whether the Addr is an initialized address (not the zero Addr).
//line /snap/go/10455/src/net/netip/netip.go:378
//
//line /snap/go/10455/src/net/netip/netip.go:378
// Note that "0.0.0.0" and "::" are both valid values.
//line /snap/go/10455/src/net/netip/netip.go:381
func (ip Addr) IsValid() bool {
//line /snap/go/10455/src/net/netip/netip.go:381
	_go_fuzz_dep_.CoverTab[4023]++
//line /snap/go/10455/src/net/netip/netip.go:381
	return ip.z != z0
//line /snap/go/10455/src/net/netip/netip.go:381
	// _ = "end of CoverTab[4023]"
//line /snap/go/10455/src/net/netip/netip.go:381
}

// BitLen returns the number of bits in the IP address:
//line /snap/go/10455/src/net/netip/netip.go:383
// 128 for IPv6, 32 for IPv4, and 0 for the zero Addr.
//line /snap/go/10455/src/net/netip/netip.go:383
//
//line /snap/go/10455/src/net/netip/netip.go:383
// Note that IPv4-mapped IPv6 addresses are considered IPv6 addresses
//line /snap/go/10455/src/net/netip/netip.go:383
// and therefore have bit length 128.
//line /snap/go/10455/src/net/netip/netip.go:388
func (ip Addr) BitLen() int {
//line /snap/go/10455/src/net/netip/netip.go:388
	_go_fuzz_dep_.CoverTab[4024]++
							switch ip.z {
	case z0:
//line /snap/go/10455/src/net/netip/netip.go:390
		_go_fuzz_dep_.CoverTab[527089]++
//line /snap/go/10455/src/net/netip/netip.go:390
		_go_fuzz_dep_.CoverTab[4026]++
								return 0
//line /snap/go/10455/src/net/netip/netip.go:391
		// _ = "end of CoverTab[4026]"
	case z4:
//line /snap/go/10455/src/net/netip/netip.go:392
		_go_fuzz_dep_.CoverTab[527090]++
//line /snap/go/10455/src/net/netip/netip.go:392
		_go_fuzz_dep_.CoverTab[4027]++
								return 32
//line /snap/go/10455/src/net/netip/netip.go:393
		// _ = "end of CoverTab[4027]"
//line /snap/go/10455/src/net/netip/netip.go:393
	default:
//line /snap/go/10455/src/net/netip/netip.go:393
		_go_fuzz_dep_.CoverTab[527091]++
//line /snap/go/10455/src/net/netip/netip.go:393
		_go_fuzz_dep_.CoverTab[4028]++
//line /snap/go/10455/src/net/netip/netip.go:393
		// _ = "end of CoverTab[4028]"
	}
//line /snap/go/10455/src/net/netip/netip.go:394
	// _ = "end of CoverTab[4024]"
//line /snap/go/10455/src/net/netip/netip.go:394
	_go_fuzz_dep_.CoverTab[4025]++
							return 128
//line /snap/go/10455/src/net/netip/netip.go:395
	// _ = "end of CoverTab[4025]"
}

// Zone returns ip's IPv6 scoped addressing zone, if any.
func (ip Addr) Zone() string {
//line /snap/go/10455/src/net/netip/netip.go:399
	_go_fuzz_dep_.CoverTab[4029]++
							if ip.z == nil {
//line /snap/go/10455/src/net/netip/netip.go:400
		_go_fuzz_dep_.CoverTab[527092]++
//line /snap/go/10455/src/net/netip/netip.go:400
		_go_fuzz_dep_.CoverTab[4031]++
								return ""
//line /snap/go/10455/src/net/netip/netip.go:401
		// _ = "end of CoverTab[4031]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:402
		_go_fuzz_dep_.CoverTab[527093]++
//line /snap/go/10455/src/net/netip/netip.go:402
		_go_fuzz_dep_.CoverTab[4032]++
//line /snap/go/10455/src/net/netip/netip.go:402
		// _ = "end of CoverTab[4032]"
//line /snap/go/10455/src/net/netip/netip.go:402
	}
//line /snap/go/10455/src/net/netip/netip.go:402
	// _ = "end of CoverTab[4029]"
//line /snap/go/10455/src/net/netip/netip.go:402
	_go_fuzz_dep_.CoverTab[4030]++
							zone, _ := ip.z.Get().(string)
							return zone
//line /snap/go/10455/src/net/netip/netip.go:404
	// _ = "end of CoverTab[4030]"
}

// Compare returns an integer comparing two IPs.
//line /snap/go/10455/src/net/netip/netip.go:407
// The result will be 0 if ip == ip2, -1 if ip < ip2, and +1 if ip > ip2.
//line /snap/go/10455/src/net/netip/netip.go:407
// The definition of "less than" is the same as the Less method.
//line /snap/go/10455/src/net/netip/netip.go:410
func (ip Addr) Compare(ip2 Addr) int {
//line /snap/go/10455/src/net/netip/netip.go:410
	_go_fuzz_dep_.CoverTab[4033]++
							f1, f2 := ip.BitLen(), ip2.BitLen()
							if f1 < f2 {
//line /snap/go/10455/src/net/netip/netip.go:412
		_go_fuzz_dep_.CoverTab[527094]++
//line /snap/go/10455/src/net/netip/netip.go:412
		_go_fuzz_dep_.CoverTab[4041]++
								return -1
//line /snap/go/10455/src/net/netip/netip.go:413
		// _ = "end of CoverTab[4041]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:414
		_go_fuzz_dep_.CoverTab[527095]++
//line /snap/go/10455/src/net/netip/netip.go:414
		_go_fuzz_dep_.CoverTab[4042]++
//line /snap/go/10455/src/net/netip/netip.go:414
		// _ = "end of CoverTab[4042]"
//line /snap/go/10455/src/net/netip/netip.go:414
	}
//line /snap/go/10455/src/net/netip/netip.go:414
	// _ = "end of CoverTab[4033]"
//line /snap/go/10455/src/net/netip/netip.go:414
	_go_fuzz_dep_.CoverTab[4034]++
							if f1 > f2 {
//line /snap/go/10455/src/net/netip/netip.go:415
		_go_fuzz_dep_.CoverTab[527096]++
//line /snap/go/10455/src/net/netip/netip.go:415
		_go_fuzz_dep_.CoverTab[4043]++
								return 1
//line /snap/go/10455/src/net/netip/netip.go:416
		// _ = "end of CoverTab[4043]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:417
		_go_fuzz_dep_.CoverTab[527097]++
//line /snap/go/10455/src/net/netip/netip.go:417
		_go_fuzz_dep_.CoverTab[4044]++
//line /snap/go/10455/src/net/netip/netip.go:417
		// _ = "end of CoverTab[4044]"
//line /snap/go/10455/src/net/netip/netip.go:417
	}
//line /snap/go/10455/src/net/netip/netip.go:417
	// _ = "end of CoverTab[4034]"
//line /snap/go/10455/src/net/netip/netip.go:417
	_go_fuzz_dep_.CoverTab[4035]++
							hi1, hi2 := ip.addr.hi, ip2.addr.hi
							if hi1 < hi2 {
//line /snap/go/10455/src/net/netip/netip.go:419
		_go_fuzz_dep_.CoverTab[527098]++
//line /snap/go/10455/src/net/netip/netip.go:419
		_go_fuzz_dep_.CoverTab[4045]++
								return -1
//line /snap/go/10455/src/net/netip/netip.go:420
		// _ = "end of CoverTab[4045]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:421
		_go_fuzz_dep_.CoverTab[527099]++
//line /snap/go/10455/src/net/netip/netip.go:421
		_go_fuzz_dep_.CoverTab[4046]++
//line /snap/go/10455/src/net/netip/netip.go:421
		// _ = "end of CoverTab[4046]"
//line /snap/go/10455/src/net/netip/netip.go:421
	}
//line /snap/go/10455/src/net/netip/netip.go:421
	// _ = "end of CoverTab[4035]"
//line /snap/go/10455/src/net/netip/netip.go:421
	_go_fuzz_dep_.CoverTab[4036]++
							if hi1 > hi2 {
//line /snap/go/10455/src/net/netip/netip.go:422
		_go_fuzz_dep_.CoverTab[527100]++
//line /snap/go/10455/src/net/netip/netip.go:422
		_go_fuzz_dep_.CoverTab[4047]++
								return 1
//line /snap/go/10455/src/net/netip/netip.go:423
		// _ = "end of CoverTab[4047]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:424
		_go_fuzz_dep_.CoverTab[527101]++
//line /snap/go/10455/src/net/netip/netip.go:424
		_go_fuzz_dep_.CoverTab[4048]++
//line /snap/go/10455/src/net/netip/netip.go:424
		// _ = "end of CoverTab[4048]"
//line /snap/go/10455/src/net/netip/netip.go:424
	}
//line /snap/go/10455/src/net/netip/netip.go:424
	// _ = "end of CoverTab[4036]"
//line /snap/go/10455/src/net/netip/netip.go:424
	_go_fuzz_dep_.CoverTab[4037]++
							lo1, lo2 := ip.addr.lo, ip2.addr.lo
							if lo1 < lo2 {
//line /snap/go/10455/src/net/netip/netip.go:426
		_go_fuzz_dep_.CoverTab[527102]++
//line /snap/go/10455/src/net/netip/netip.go:426
		_go_fuzz_dep_.CoverTab[4049]++
								return -1
//line /snap/go/10455/src/net/netip/netip.go:427
		// _ = "end of CoverTab[4049]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:428
		_go_fuzz_dep_.CoverTab[527103]++
//line /snap/go/10455/src/net/netip/netip.go:428
		_go_fuzz_dep_.CoverTab[4050]++
//line /snap/go/10455/src/net/netip/netip.go:428
		// _ = "end of CoverTab[4050]"
//line /snap/go/10455/src/net/netip/netip.go:428
	}
//line /snap/go/10455/src/net/netip/netip.go:428
	// _ = "end of CoverTab[4037]"
//line /snap/go/10455/src/net/netip/netip.go:428
	_go_fuzz_dep_.CoverTab[4038]++
							if lo1 > lo2 {
//line /snap/go/10455/src/net/netip/netip.go:429
		_go_fuzz_dep_.CoverTab[527104]++
//line /snap/go/10455/src/net/netip/netip.go:429
		_go_fuzz_dep_.CoverTab[4051]++
								return 1
//line /snap/go/10455/src/net/netip/netip.go:430
		// _ = "end of CoverTab[4051]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:431
		_go_fuzz_dep_.CoverTab[527105]++
//line /snap/go/10455/src/net/netip/netip.go:431
		_go_fuzz_dep_.CoverTab[4052]++
//line /snap/go/10455/src/net/netip/netip.go:431
		// _ = "end of CoverTab[4052]"
//line /snap/go/10455/src/net/netip/netip.go:431
	}
//line /snap/go/10455/src/net/netip/netip.go:431
	// _ = "end of CoverTab[4038]"
//line /snap/go/10455/src/net/netip/netip.go:431
	_go_fuzz_dep_.CoverTab[4039]++
							if ip.Is6() {
//line /snap/go/10455/src/net/netip/netip.go:432
		_go_fuzz_dep_.CoverTab[527106]++
//line /snap/go/10455/src/net/netip/netip.go:432
		_go_fuzz_dep_.CoverTab[4053]++
								za, zb := ip.Zone(), ip2.Zone()
								if za < zb {
//line /snap/go/10455/src/net/netip/netip.go:434
			_go_fuzz_dep_.CoverTab[527108]++
//line /snap/go/10455/src/net/netip/netip.go:434
			_go_fuzz_dep_.CoverTab[4055]++
									return -1
//line /snap/go/10455/src/net/netip/netip.go:435
			// _ = "end of CoverTab[4055]"
		} else {
//line /snap/go/10455/src/net/netip/netip.go:436
			_go_fuzz_dep_.CoverTab[527109]++
//line /snap/go/10455/src/net/netip/netip.go:436
			_go_fuzz_dep_.CoverTab[4056]++
//line /snap/go/10455/src/net/netip/netip.go:436
			// _ = "end of CoverTab[4056]"
//line /snap/go/10455/src/net/netip/netip.go:436
		}
//line /snap/go/10455/src/net/netip/netip.go:436
		// _ = "end of CoverTab[4053]"
//line /snap/go/10455/src/net/netip/netip.go:436
		_go_fuzz_dep_.CoverTab[4054]++
								if za > zb {
//line /snap/go/10455/src/net/netip/netip.go:437
			_go_fuzz_dep_.CoverTab[527110]++
//line /snap/go/10455/src/net/netip/netip.go:437
			_go_fuzz_dep_.CoverTab[4057]++
									return 1
//line /snap/go/10455/src/net/netip/netip.go:438
			// _ = "end of CoverTab[4057]"
		} else {
//line /snap/go/10455/src/net/netip/netip.go:439
			_go_fuzz_dep_.CoverTab[527111]++
//line /snap/go/10455/src/net/netip/netip.go:439
			_go_fuzz_dep_.CoverTab[4058]++
//line /snap/go/10455/src/net/netip/netip.go:439
			// _ = "end of CoverTab[4058]"
//line /snap/go/10455/src/net/netip/netip.go:439
		}
//line /snap/go/10455/src/net/netip/netip.go:439
		// _ = "end of CoverTab[4054]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:440
		_go_fuzz_dep_.CoverTab[527107]++
//line /snap/go/10455/src/net/netip/netip.go:440
		_go_fuzz_dep_.CoverTab[4059]++
//line /snap/go/10455/src/net/netip/netip.go:440
		// _ = "end of CoverTab[4059]"
//line /snap/go/10455/src/net/netip/netip.go:440
	}
//line /snap/go/10455/src/net/netip/netip.go:440
	// _ = "end of CoverTab[4039]"
//line /snap/go/10455/src/net/netip/netip.go:440
	_go_fuzz_dep_.CoverTab[4040]++
							return 0
//line /snap/go/10455/src/net/netip/netip.go:441
	// _ = "end of CoverTab[4040]"
}

// Less reports whether ip sorts before ip2.
//line /snap/go/10455/src/net/netip/netip.go:444
// IP addresses sort first by length, then their address.
//line /snap/go/10455/src/net/netip/netip.go:444
// IPv6 addresses with zones sort just after the same address without a zone.
//line /snap/go/10455/src/net/netip/netip.go:447
func (ip Addr) Less(ip2 Addr) bool {
//line /snap/go/10455/src/net/netip/netip.go:447
	_go_fuzz_dep_.CoverTab[4060]++
//line /snap/go/10455/src/net/netip/netip.go:447
	return ip.Compare(ip2) == -1
//line /snap/go/10455/src/net/netip/netip.go:447
	// _ = "end of CoverTab[4060]"
//line /snap/go/10455/src/net/netip/netip.go:447
}

// Is4 reports whether ip is an IPv4 address.
//line /snap/go/10455/src/net/netip/netip.go:449
//
//line /snap/go/10455/src/net/netip/netip.go:449
// It returns false for IPv4-mapped IPv6 addresses. See Addr.Unmap.
//line /snap/go/10455/src/net/netip/netip.go:452
func (ip Addr) Is4() bool {
//line /snap/go/10455/src/net/netip/netip.go:452
	_go_fuzz_dep_.CoverTab[4061]++
							return ip.z == z4
//line /snap/go/10455/src/net/netip/netip.go:453
	// _ = "end of CoverTab[4061]"
}

// Is4In6 reports whether ip is an IPv4-mapped IPv6 address.
func (ip Addr) Is4In6() bool {
//line /snap/go/10455/src/net/netip/netip.go:457
	_go_fuzz_dep_.CoverTab[4062]++
							return ip.Is6() && func() bool {
//line /snap/go/10455/src/net/netip/netip.go:458
		_go_fuzz_dep_.CoverTab[4063]++
//line /snap/go/10455/src/net/netip/netip.go:458
		return ip.addr.hi == 0
//line /snap/go/10455/src/net/netip/netip.go:458
		// _ = "end of CoverTab[4063]"
//line /snap/go/10455/src/net/netip/netip.go:458
	}() && func() bool {
//line /snap/go/10455/src/net/netip/netip.go:458
		_go_fuzz_dep_.CoverTab[4064]++
//line /snap/go/10455/src/net/netip/netip.go:458
		return ip.addr.lo>>32 == 0xffff
//line /snap/go/10455/src/net/netip/netip.go:458
		// _ = "end of CoverTab[4064]"
//line /snap/go/10455/src/net/netip/netip.go:458
	}()
//line /snap/go/10455/src/net/netip/netip.go:458
	// _ = "end of CoverTab[4062]"
}

// Is6 reports whether ip is an IPv6 address, including IPv4-mapped
//line /snap/go/10455/src/net/netip/netip.go:461
// IPv6 addresses.
//line /snap/go/10455/src/net/netip/netip.go:463
func (ip Addr) Is6() bool {
//line /snap/go/10455/src/net/netip/netip.go:463
	_go_fuzz_dep_.CoverTab[4065]++
							return ip.z != z0 && func() bool {
//line /snap/go/10455/src/net/netip/netip.go:464
		_go_fuzz_dep_.CoverTab[4066]++
//line /snap/go/10455/src/net/netip/netip.go:464
		return ip.z != z4
//line /snap/go/10455/src/net/netip/netip.go:464
		// _ = "end of CoverTab[4066]"
//line /snap/go/10455/src/net/netip/netip.go:464
	}()
//line /snap/go/10455/src/net/netip/netip.go:464
	// _ = "end of CoverTab[4065]"
}

// Unmap returns ip with any IPv4-mapped IPv6 address prefix removed.
//line /snap/go/10455/src/net/netip/netip.go:467
//
//line /snap/go/10455/src/net/netip/netip.go:467
// That is, if ip is an IPv6 address wrapping an IPv4 address, it
//line /snap/go/10455/src/net/netip/netip.go:467
// returns the wrapped IPv4 address. Otherwise it returns ip unmodified.
//line /snap/go/10455/src/net/netip/netip.go:471
func (ip Addr) Unmap() Addr {
//line /snap/go/10455/src/net/netip/netip.go:471
	_go_fuzz_dep_.CoverTab[4067]++
							if ip.Is4In6() {
//line /snap/go/10455/src/net/netip/netip.go:472
		_go_fuzz_dep_.CoverTab[527112]++
//line /snap/go/10455/src/net/netip/netip.go:472
		_go_fuzz_dep_.CoverTab[4069]++
								ip.z = z4
//line /snap/go/10455/src/net/netip/netip.go:473
		// _ = "end of CoverTab[4069]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:474
		_go_fuzz_dep_.CoverTab[527113]++
//line /snap/go/10455/src/net/netip/netip.go:474
		_go_fuzz_dep_.CoverTab[4070]++
//line /snap/go/10455/src/net/netip/netip.go:474
		// _ = "end of CoverTab[4070]"
//line /snap/go/10455/src/net/netip/netip.go:474
	}
//line /snap/go/10455/src/net/netip/netip.go:474
	// _ = "end of CoverTab[4067]"
//line /snap/go/10455/src/net/netip/netip.go:474
	_go_fuzz_dep_.CoverTab[4068]++
							return ip
//line /snap/go/10455/src/net/netip/netip.go:475
	// _ = "end of CoverTab[4068]"
}

// WithZone returns an IP that's the same as ip but with the provided
//line /snap/go/10455/src/net/netip/netip.go:478
// zone. If zone is empty, the zone is removed. If ip is an IPv4
//line /snap/go/10455/src/net/netip/netip.go:478
// address, WithZone is a no-op and returns ip unchanged.
//line /snap/go/10455/src/net/netip/netip.go:481
func (ip Addr) WithZone(zone string) Addr {
//line /snap/go/10455/src/net/netip/netip.go:481
	_go_fuzz_dep_.CoverTab[4071]++
							if !ip.Is6() {
//line /snap/go/10455/src/net/netip/netip.go:482
		_go_fuzz_dep_.CoverTab[527114]++
//line /snap/go/10455/src/net/netip/netip.go:482
		_go_fuzz_dep_.CoverTab[4074]++
								return ip
//line /snap/go/10455/src/net/netip/netip.go:483
		// _ = "end of CoverTab[4074]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:484
		_go_fuzz_dep_.CoverTab[527115]++
//line /snap/go/10455/src/net/netip/netip.go:484
		_go_fuzz_dep_.CoverTab[4075]++
//line /snap/go/10455/src/net/netip/netip.go:484
		// _ = "end of CoverTab[4075]"
//line /snap/go/10455/src/net/netip/netip.go:484
	}
//line /snap/go/10455/src/net/netip/netip.go:484
	// _ = "end of CoverTab[4071]"
//line /snap/go/10455/src/net/netip/netip.go:484
	_go_fuzz_dep_.CoverTab[4072]++
							if zone == "" {
//line /snap/go/10455/src/net/netip/netip.go:485
		_go_fuzz_dep_.CoverTab[527116]++
//line /snap/go/10455/src/net/netip/netip.go:485
		_go_fuzz_dep_.CoverTab[4076]++
								ip.z = z6noz
								return ip
//line /snap/go/10455/src/net/netip/netip.go:487
		// _ = "end of CoverTab[4076]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:488
		_go_fuzz_dep_.CoverTab[527117]++
//line /snap/go/10455/src/net/netip/netip.go:488
		_go_fuzz_dep_.CoverTab[4077]++
//line /snap/go/10455/src/net/netip/netip.go:488
		// _ = "end of CoverTab[4077]"
//line /snap/go/10455/src/net/netip/netip.go:488
	}
//line /snap/go/10455/src/net/netip/netip.go:488
	// _ = "end of CoverTab[4072]"
//line /snap/go/10455/src/net/netip/netip.go:488
	_go_fuzz_dep_.CoverTab[4073]++
							ip.z = intern.GetByString(zone)
							return ip
//line /snap/go/10455/src/net/netip/netip.go:490
	// _ = "end of CoverTab[4073]"
}

// withoutZone unconditionally strips the zone from ip.
//line /snap/go/10455/src/net/netip/netip.go:493
// It's similar to WithZone, but small enough to be inlinable.
//line /snap/go/10455/src/net/netip/netip.go:495
func (ip Addr) withoutZone() Addr {
//line /snap/go/10455/src/net/netip/netip.go:495
	_go_fuzz_dep_.CoverTab[4078]++
							if !ip.Is6() {
//line /snap/go/10455/src/net/netip/netip.go:496
		_go_fuzz_dep_.CoverTab[527118]++
//line /snap/go/10455/src/net/netip/netip.go:496
		_go_fuzz_dep_.CoverTab[4080]++
								return ip
//line /snap/go/10455/src/net/netip/netip.go:497
		// _ = "end of CoverTab[4080]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:498
		_go_fuzz_dep_.CoverTab[527119]++
//line /snap/go/10455/src/net/netip/netip.go:498
		_go_fuzz_dep_.CoverTab[4081]++
//line /snap/go/10455/src/net/netip/netip.go:498
		// _ = "end of CoverTab[4081]"
//line /snap/go/10455/src/net/netip/netip.go:498
	}
//line /snap/go/10455/src/net/netip/netip.go:498
	// _ = "end of CoverTab[4078]"
//line /snap/go/10455/src/net/netip/netip.go:498
	_go_fuzz_dep_.CoverTab[4079]++
							ip.z = z6noz
							return ip
//line /snap/go/10455/src/net/netip/netip.go:500
	// _ = "end of CoverTab[4079]"
}

// hasZone reports whether ip has an IPv6 zone.
func (ip Addr) hasZone() bool {
//line /snap/go/10455/src/net/netip/netip.go:504
	_go_fuzz_dep_.CoverTab[4082]++
							return ip.z != z0 && func() bool {
//line /snap/go/10455/src/net/netip/netip.go:505
		_go_fuzz_dep_.CoverTab[4083]++
//line /snap/go/10455/src/net/netip/netip.go:505
		return ip.z != z4
//line /snap/go/10455/src/net/netip/netip.go:505
		// _ = "end of CoverTab[4083]"
//line /snap/go/10455/src/net/netip/netip.go:505
	}() && func() bool {
//line /snap/go/10455/src/net/netip/netip.go:505
		_go_fuzz_dep_.CoverTab[4084]++
//line /snap/go/10455/src/net/netip/netip.go:505
		return ip.z != z6noz
//line /snap/go/10455/src/net/netip/netip.go:505
		// _ = "end of CoverTab[4084]"
//line /snap/go/10455/src/net/netip/netip.go:505
	}()
//line /snap/go/10455/src/net/netip/netip.go:505
	// _ = "end of CoverTab[4082]"
}

// IsLinkLocalUnicast reports whether ip is a link-local unicast address.
func (ip Addr) IsLinkLocalUnicast() bool {
//line /snap/go/10455/src/net/netip/netip.go:509
	_go_fuzz_dep_.CoverTab[4085]++

//line /snap/go/10455/src/net/netip/netip.go:512
	if ip.Is4() {
//line /snap/go/10455/src/net/netip/netip.go:512
		_go_fuzz_dep_.CoverTab[527120]++
//line /snap/go/10455/src/net/netip/netip.go:512
		_go_fuzz_dep_.CoverTab[4088]++
								return ip.v4(0) == 169 && func() bool {
//line /snap/go/10455/src/net/netip/netip.go:513
			_go_fuzz_dep_.CoverTab[4089]++
//line /snap/go/10455/src/net/netip/netip.go:513
			return ip.v4(1) == 254
//line /snap/go/10455/src/net/netip/netip.go:513
			// _ = "end of CoverTab[4089]"
//line /snap/go/10455/src/net/netip/netip.go:513
		}()
//line /snap/go/10455/src/net/netip/netip.go:513
		// _ = "end of CoverTab[4088]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:514
		_go_fuzz_dep_.CoverTab[527121]++
//line /snap/go/10455/src/net/netip/netip.go:514
		_go_fuzz_dep_.CoverTab[4090]++
//line /snap/go/10455/src/net/netip/netip.go:514
		// _ = "end of CoverTab[4090]"
//line /snap/go/10455/src/net/netip/netip.go:514
	}
//line /snap/go/10455/src/net/netip/netip.go:514
	// _ = "end of CoverTab[4085]"
//line /snap/go/10455/src/net/netip/netip.go:514
	_go_fuzz_dep_.CoverTab[4086]++

//line /snap/go/10455/src/net/netip/netip.go:517
	if ip.Is6() {
//line /snap/go/10455/src/net/netip/netip.go:517
		_go_fuzz_dep_.CoverTab[527122]++
//line /snap/go/10455/src/net/netip/netip.go:517
		_go_fuzz_dep_.CoverTab[4091]++
								return ip.v6u16(0)&0xffc0 == 0xfe80
//line /snap/go/10455/src/net/netip/netip.go:518
		// _ = "end of CoverTab[4091]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:519
		_go_fuzz_dep_.CoverTab[527123]++
//line /snap/go/10455/src/net/netip/netip.go:519
		_go_fuzz_dep_.CoverTab[4092]++
//line /snap/go/10455/src/net/netip/netip.go:519
		// _ = "end of CoverTab[4092]"
//line /snap/go/10455/src/net/netip/netip.go:519
	}
//line /snap/go/10455/src/net/netip/netip.go:519
	// _ = "end of CoverTab[4086]"
//line /snap/go/10455/src/net/netip/netip.go:519
	_go_fuzz_dep_.CoverTab[4087]++
							return false
//line /snap/go/10455/src/net/netip/netip.go:520
	// _ = "end of CoverTab[4087]"
}

// IsLoopback reports whether ip is a loopback address.
func (ip Addr) IsLoopback() bool {
//line /snap/go/10455/src/net/netip/netip.go:524
	_go_fuzz_dep_.CoverTab[4093]++

//line /snap/go/10455/src/net/netip/netip.go:527
	if ip.Is4() {
//line /snap/go/10455/src/net/netip/netip.go:527
		_go_fuzz_dep_.CoverTab[527124]++
//line /snap/go/10455/src/net/netip/netip.go:527
		_go_fuzz_dep_.CoverTab[4096]++
								return ip.v4(0) == 127
//line /snap/go/10455/src/net/netip/netip.go:528
		// _ = "end of CoverTab[4096]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:529
		_go_fuzz_dep_.CoverTab[527125]++
//line /snap/go/10455/src/net/netip/netip.go:529
		_go_fuzz_dep_.CoverTab[4097]++
//line /snap/go/10455/src/net/netip/netip.go:529
		// _ = "end of CoverTab[4097]"
//line /snap/go/10455/src/net/netip/netip.go:529
	}
//line /snap/go/10455/src/net/netip/netip.go:529
	// _ = "end of CoverTab[4093]"
//line /snap/go/10455/src/net/netip/netip.go:529
	_go_fuzz_dep_.CoverTab[4094]++

//line /snap/go/10455/src/net/netip/netip.go:532
	if ip.Is6() {
//line /snap/go/10455/src/net/netip/netip.go:532
		_go_fuzz_dep_.CoverTab[527126]++
//line /snap/go/10455/src/net/netip/netip.go:532
		_go_fuzz_dep_.CoverTab[4098]++
								return ip.addr.hi == 0 && func() bool {
//line /snap/go/10455/src/net/netip/netip.go:533
			_go_fuzz_dep_.CoverTab[4099]++
//line /snap/go/10455/src/net/netip/netip.go:533
			return ip.addr.lo == 1
//line /snap/go/10455/src/net/netip/netip.go:533
			// _ = "end of CoverTab[4099]"
//line /snap/go/10455/src/net/netip/netip.go:533
		}()
//line /snap/go/10455/src/net/netip/netip.go:533
		// _ = "end of CoverTab[4098]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:534
		_go_fuzz_dep_.CoverTab[527127]++
//line /snap/go/10455/src/net/netip/netip.go:534
		_go_fuzz_dep_.CoverTab[4100]++
//line /snap/go/10455/src/net/netip/netip.go:534
		// _ = "end of CoverTab[4100]"
//line /snap/go/10455/src/net/netip/netip.go:534
	}
//line /snap/go/10455/src/net/netip/netip.go:534
	// _ = "end of CoverTab[4094]"
//line /snap/go/10455/src/net/netip/netip.go:534
	_go_fuzz_dep_.CoverTab[4095]++
							return false
//line /snap/go/10455/src/net/netip/netip.go:535
	// _ = "end of CoverTab[4095]"
}

// IsMulticast reports whether ip is a multicast address.
func (ip Addr) IsMulticast() bool {
//line /snap/go/10455/src/net/netip/netip.go:539
	_go_fuzz_dep_.CoverTab[4101]++

//line /snap/go/10455/src/net/netip/netip.go:542
	if ip.Is4() {
//line /snap/go/10455/src/net/netip/netip.go:542
		_go_fuzz_dep_.CoverTab[527128]++
//line /snap/go/10455/src/net/netip/netip.go:542
		_go_fuzz_dep_.CoverTab[4104]++
								return ip.v4(0)&0xf0 == 0xe0
//line /snap/go/10455/src/net/netip/netip.go:543
		// _ = "end of CoverTab[4104]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:544
		_go_fuzz_dep_.CoverTab[527129]++
//line /snap/go/10455/src/net/netip/netip.go:544
		_go_fuzz_dep_.CoverTab[4105]++
//line /snap/go/10455/src/net/netip/netip.go:544
		// _ = "end of CoverTab[4105]"
//line /snap/go/10455/src/net/netip/netip.go:544
	}
//line /snap/go/10455/src/net/netip/netip.go:544
	// _ = "end of CoverTab[4101]"
//line /snap/go/10455/src/net/netip/netip.go:544
	_go_fuzz_dep_.CoverTab[4102]++

//line /snap/go/10455/src/net/netip/netip.go:547
	if ip.Is6() {
//line /snap/go/10455/src/net/netip/netip.go:547
		_go_fuzz_dep_.CoverTab[527130]++
//line /snap/go/10455/src/net/netip/netip.go:547
		_go_fuzz_dep_.CoverTab[4106]++
								return ip.addr.hi>>(64-8) == 0xff
//line /snap/go/10455/src/net/netip/netip.go:548
		// _ = "end of CoverTab[4106]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:549
		_go_fuzz_dep_.CoverTab[527131]++
//line /snap/go/10455/src/net/netip/netip.go:549
		_go_fuzz_dep_.CoverTab[4107]++
//line /snap/go/10455/src/net/netip/netip.go:549
		// _ = "end of CoverTab[4107]"
//line /snap/go/10455/src/net/netip/netip.go:549
	}
//line /snap/go/10455/src/net/netip/netip.go:549
	// _ = "end of CoverTab[4102]"
//line /snap/go/10455/src/net/netip/netip.go:549
	_go_fuzz_dep_.CoverTab[4103]++
							return false
//line /snap/go/10455/src/net/netip/netip.go:550
	// _ = "end of CoverTab[4103]"
}

// IsInterfaceLocalMulticast reports whether ip is an IPv6 interface-local
//line /snap/go/10455/src/net/netip/netip.go:553
// multicast address.
//line /snap/go/10455/src/net/netip/netip.go:555
func (ip Addr) IsInterfaceLocalMulticast() bool {
//line /snap/go/10455/src/net/netip/netip.go:555
	_go_fuzz_dep_.CoverTab[4108]++

//line /snap/go/10455/src/net/netip/netip.go:558
	if ip.Is6() {
//line /snap/go/10455/src/net/netip/netip.go:558
		_go_fuzz_dep_.CoverTab[527132]++
//line /snap/go/10455/src/net/netip/netip.go:558
		_go_fuzz_dep_.CoverTab[4110]++
								return ip.v6u16(0)&0xff0f == 0xff01
//line /snap/go/10455/src/net/netip/netip.go:559
		// _ = "end of CoverTab[4110]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:560
		_go_fuzz_dep_.CoverTab[527133]++
//line /snap/go/10455/src/net/netip/netip.go:560
		_go_fuzz_dep_.CoverTab[4111]++
//line /snap/go/10455/src/net/netip/netip.go:560
		// _ = "end of CoverTab[4111]"
//line /snap/go/10455/src/net/netip/netip.go:560
	}
//line /snap/go/10455/src/net/netip/netip.go:560
	// _ = "end of CoverTab[4108]"
//line /snap/go/10455/src/net/netip/netip.go:560
	_go_fuzz_dep_.CoverTab[4109]++
							return false
//line /snap/go/10455/src/net/netip/netip.go:561
	// _ = "end of CoverTab[4109]"
}

// IsLinkLocalMulticast reports whether ip is a link-local multicast address.
func (ip Addr) IsLinkLocalMulticast() bool {
//line /snap/go/10455/src/net/netip/netip.go:565
	_go_fuzz_dep_.CoverTab[4112]++

//line /snap/go/10455/src/net/netip/netip.go:568
	if ip.Is4() {
//line /snap/go/10455/src/net/netip/netip.go:568
		_go_fuzz_dep_.CoverTab[527134]++
//line /snap/go/10455/src/net/netip/netip.go:568
		_go_fuzz_dep_.CoverTab[4115]++
								return ip.v4(0) == 224 && func() bool {
//line /snap/go/10455/src/net/netip/netip.go:569
			_go_fuzz_dep_.CoverTab[4116]++
//line /snap/go/10455/src/net/netip/netip.go:569
			return ip.v4(1) == 0
//line /snap/go/10455/src/net/netip/netip.go:569
			// _ = "end of CoverTab[4116]"
//line /snap/go/10455/src/net/netip/netip.go:569
		}() && func() bool {
//line /snap/go/10455/src/net/netip/netip.go:569
			_go_fuzz_dep_.CoverTab[4117]++
//line /snap/go/10455/src/net/netip/netip.go:569
			return ip.v4(2) == 0
//line /snap/go/10455/src/net/netip/netip.go:569
			// _ = "end of CoverTab[4117]"
//line /snap/go/10455/src/net/netip/netip.go:569
		}()
//line /snap/go/10455/src/net/netip/netip.go:569
		// _ = "end of CoverTab[4115]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:570
		_go_fuzz_dep_.CoverTab[527135]++
//line /snap/go/10455/src/net/netip/netip.go:570
		_go_fuzz_dep_.CoverTab[4118]++
//line /snap/go/10455/src/net/netip/netip.go:570
		// _ = "end of CoverTab[4118]"
//line /snap/go/10455/src/net/netip/netip.go:570
	}
//line /snap/go/10455/src/net/netip/netip.go:570
	// _ = "end of CoverTab[4112]"
//line /snap/go/10455/src/net/netip/netip.go:570
	_go_fuzz_dep_.CoverTab[4113]++

//line /snap/go/10455/src/net/netip/netip.go:573
	if ip.Is6() {
//line /snap/go/10455/src/net/netip/netip.go:573
		_go_fuzz_dep_.CoverTab[527136]++
//line /snap/go/10455/src/net/netip/netip.go:573
		_go_fuzz_dep_.CoverTab[4119]++
								return ip.v6u16(0)&0xff0f == 0xff02
//line /snap/go/10455/src/net/netip/netip.go:574
		// _ = "end of CoverTab[4119]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:575
		_go_fuzz_dep_.CoverTab[527137]++
//line /snap/go/10455/src/net/netip/netip.go:575
		_go_fuzz_dep_.CoverTab[4120]++
//line /snap/go/10455/src/net/netip/netip.go:575
		// _ = "end of CoverTab[4120]"
//line /snap/go/10455/src/net/netip/netip.go:575
	}
//line /snap/go/10455/src/net/netip/netip.go:575
	// _ = "end of CoverTab[4113]"
//line /snap/go/10455/src/net/netip/netip.go:575
	_go_fuzz_dep_.CoverTab[4114]++
							return false
//line /snap/go/10455/src/net/netip/netip.go:576
	// _ = "end of CoverTab[4114]"
}

// IsGlobalUnicast reports whether ip is a global unicast address.
//line /snap/go/10455/src/net/netip/netip.go:579
//
//line /snap/go/10455/src/net/netip/netip.go:579
// It returns true for IPv6 addresses which fall outside of the current
//line /snap/go/10455/src/net/netip/netip.go:579
// IANA-allocated 2000::/3 global unicast space, with the exception of the
//line /snap/go/10455/src/net/netip/netip.go:579
// link-local address space. It also returns true even if ip is in the IPv4
//line /snap/go/10455/src/net/netip/netip.go:579
// private address space or IPv6 unique local address space.
//line /snap/go/10455/src/net/netip/netip.go:579
// It returns false for the zero Addr.
//line /snap/go/10455/src/net/netip/netip.go:579
//
//line /snap/go/10455/src/net/netip/netip.go:579
// For reference, see RFC 1122, RFC 4291, and RFC 4632.
//line /snap/go/10455/src/net/netip/netip.go:588
func (ip Addr) IsGlobalUnicast() bool {
//line /snap/go/10455/src/net/netip/netip.go:588
	_go_fuzz_dep_.CoverTab[4121]++
							if ip.z == z0 {
//line /snap/go/10455/src/net/netip/netip.go:589
		_go_fuzz_dep_.CoverTab[527138]++
//line /snap/go/10455/src/net/netip/netip.go:589
		_go_fuzz_dep_.CoverTab[4124]++

								return false
//line /snap/go/10455/src/net/netip/netip.go:591
		// _ = "end of CoverTab[4124]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:592
		_go_fuzz_dep_.CoverTab[527139]++
//line /snap/go/10455/src/net/netip/netip.go:592
		_go_fuzz_dep_.CoverTab[4125]++
//line /snap/go/10455/src/net/netip/netip.go:592
		// _ = "end of CoverTab[4125]"
//line /snap/go/10455/src/net/netip/netip.go:592
	}
//line /snap/go/10455/src/net/netip/netip.go:592
	// _ = "end of CoverTab[4121]"
//line /snap/go/10455/src/net/netip/netip.go:592
	_go_fuzz_dep_.CoverTab[4122]++

//line /snap/go/10455/src/net/netip/netip.go:596
	if ip.Is4() && func() bool {
//line /snap/go/10455/src/net/netip/netip.go:596
		_go_fuzz_dep_.CoverTab[4126]++
//line /snap/go/10455/src/net/netip/netip.go:596
		return (ip == IPv4Unspecified() || func() bool {
//line /snap/go/10455/src/net/netip/netip.go:596
			_go_fuzz_dep_.CoverTab[4127]++
//line /snap/go/10455/src/net/netip/netip.go:596
			return ip == AddrFrom4([4]byte{255, 255, 255, 255})
//line /snap/go/10455/src/net/netip/netip.go:596
			// _ = "end of CoverTab[4127]"
//line /snap/go/10455/src/net/netip/netip.go:596
		}())
//line /snap/go/10455/src/net/netip/netip.go:596
		// _ = "end of CoverTab[4126]"
//line /snap/go/10455/src/net/netip/netip.go:596
	}() {
//line /snap/go/10455/src/net/netip/netip.go:596
		_go_fuzz_dep_.CoverTab[527140]++
//line /snap/go/10455/src/net/netip/netip.go:596
		_go_fuzz_dep_.CoverTab[4128]++
								return false
//line /snap/go/10455/src/net/netip/netip.go:597
		// _ = "end of CoverTab[4128]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:598
		_go_fuzz_dep_.CoverTab[527141]++
//line /snap/go/10455/src/net/netip/netip.go:598
		_go_fuzz_dep_.CoverTab[4129]++
//line /snap/go/10455/src/net/netip/netip.go:598
		// _ = "end of CoverTab[4129]"
//line /snap/go/10455/src/net/netip/netip.go:598
	}
//line /snap/go/10455/src/net/netip/netip.go:598
	// _ = "end of CoverTab[4122]"
//line /snap/go/10455/src/net/netip/netip.go:598
	_go_fuzz_dep_.CoverTab[4123]++

							return ip != IPv6Unspecified() && func() bool {
//line /snap/go/10455/src/net/netip/netip.go:600
		_go_fuzz_dep_.CoverTab[4130]++
//line /snap/go/10455/src/net/netip/netip.go:600
		return !ip.IsLoopback()
								// _ = "end of CoverTab[4130]"
//line /snap/go/10455/src/net/netip/netip.go:601
	}() && func() bool {
//line /snap/go/10455/src/net/netip/netip.go:601
		_go_fuzz_dep_.CoverTab[4131]++
//line /snap/go/10455/src/net/netip/netip.go:601
		return !ip.IsMulticast()
								// _ = "end of CoverTab[4131]"
//line /snap/go/10455/src/net/netip/netip.go:602
	}() && func() bool {
//line /snap/go/10455/src/net/netip/netip.go:602
		_go_fuzz_dep_.CoverTab[4132]++
//line /snap/go/10455/src/net/netip/netip.go:602
		return !ip.IsLinkLocalUnicast()
								// _ = "end of CoverTab[4132]"
//line /snap/go/10455/src/net/netip/netip.go:603
	}()
//line /snap/go/10455/src/net/netip/netip.go:603
	// _ = "end of CoverTab[4123]"
}

// IsPrivate reports whether ip is a private address, according to RFC 1918
//line /snap/go/10455/src/net/netip/netip.go:606
// (IPv4 addresses) and RFC 4193 (IPv6 addresses). That is, it reports whether
//line /snap/go/10455/src/net/netip/netip.go:606
// ip is in 10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16, or fc00::/7. This is the
//line /snap/go/10455/src/net/netip/netip.go:606
// same as net.IP.IsPrivate.
//line /snap/go/10455/src/net/netip/netip.go:610
func (ip Addr) IsPrivate() bool {
//line /snap/go/10455/src/net/netip/netip.go:610
	_go_fuzz_dep_.CoverTab[4133]++

							if ip.Is4() {
//line /snap/go/10455/src/net/netip/netip.go:612
		_go_fuzz_dep_.CoverTab[527142]++
//line /snap/go/10455/src/net/netip/netip.go:612
		_go_fuzz_dep_.CoverTab[4136]++

//line /snap/go/10455/src/net/netip/netip.go:615
		return ip.v4(0) == 10 || func() bool {
//line /snap/go/10455/src/net/netip/netip.go:615
			_go_fuzz_dep_.CoverTab[4137]++
//line /snap/go/10455/src/net/netip/netip.go:615
			return (ip.v4(0) == 172 && func() bool {
										_go_fuzz_dep_.CoverTab[4138]++
//line /snap/go/10455/src/net/netip/netip.go:616
				return ip.v4(1)&0xf0 == 16
//line /snap/go/10455/src/net/netip/netip.go:616
				// _ = "end of CoverTab[4138]"
//line /snap/go/10455/src/net/netip/netip.go:616
			}())
//line /snap/go/10455/src/net/netip/netip.go:616
			// _ = "end of CoverTab[4137]"
//line /snap/go/10455/src/net/netip/netip.go:616
		}() || func() bool {
//line /snap/go/10455/src/net/netip/netip.go:616
			_go_fuzz_dep_.CoverTab[4139]++
//line /snap/go/10455/src/net/netip/netip.go:616
			return (ip.v4(0) == 192 && func() bool {
										_go_fuzz_dep_.CoverTab[4140]++
//line /snap/go/10455/src/net/netip/netip.go:617
				return ip.v4(1) == 168
//line /snap/go/10455/src/net/netip/netip.go:617
				// _ = "end of CoverTab[4140]"
//line /snap/go/10455/src/net/netip/netip.go:617
			}())
//line /snap/go/10455/src/net/netip/netip.go:617
			// _ = "end of CoverTab[4139]"
//line /snap/go/10455/src/net/netip/netip.go:617
		}()
//line /snap/go/10455/src/net/netip/netip.go:617
		// _ = "end of CoverTab[4136]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:618
		_go_fuzz_dep_.CoverTab[527143]++
//line /snap/go/10455/src/net/netip/netip.go:618
		_go_fuzz_dep_.CoverTab[4141]++
//line /snap/go/10455/src/net/netip/netip.go:618
		// _ = "end of CoverTab[4141]"
//line /snap/go/10455/src/net/netip/netip.go:618
	}
//line /snap/go/10455/src/net/netip/netip.go:618
	// _ = "end of CoverTab[4133]"
//line /snap/go/10455/src/net/netip/netip.go:618
	_go_fuzz_dep_.CoverTab[4134]++

							if ip.Is6() {
//line /snap/go/10455/src/net/netip/netip.go:620
		_go_fuzz_dep_.CoverTab[527144]++
//line /snap/go/10455/src/net/netip/netip.go:620
		_go_fuzz_dep_.CoverTab[4142]++

//line /snap/go/10455/src/net/netip/netip.go:623
		return ip.v6(0)&0xfe == 0xfc
//line /snap/go/10455/src/net/netip/netip.go:623
		// _ = "end of CoverTab[4142]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:624
		_go_fuzz_dep_.CoverTab[527145]++
//line /snap/go/10455/src/net/netip/netip.go:624
		_go_fuzz_dep_.CoverTab[4143]++
//line /snap/go/10455/src/net/netip/netip.go:624
		// _ = "end of CoverTab[4143]"
//line /snap/go/10455/src/net/netip/netip.go:624
	}
//line /snap/go/10455/src/net/netip/netip.go:624
	// _ = "end of CoverTab[4134]"
//line /snap/go/10455/src/net/netip/netip.go:624
	_go_fuzz_dep_.CoverTab[4135]++

							return false
//line /snap/go/10455/src/net/netip/netip.go:626
	// _ = "end of CoverTab[4135]"
}

// IsUnspecified reports whether ip is an unspecified address, either the IPv4
//line /snap/go/10455/src/net/netip/netip.go:629
// address "0.0.0.0" or the IPv6 address "::".
//line /snap/go/10455/src/net/netip/netip.go:629
//
//line /snap/go/10455/src/net/netip/netip.go:629
// Note that the zero Addr is not an unspecified address.
//line /snap/go/10455/src/net/netip/netip.go:633
func (ip Addr) IsUnspecified() bool {
//line /snap/go/10455/src/net/netip/netip.go:633
	_go_fuzz_dep_.CoverTab[4144]++
							return ip == IPv4Unspecified() || func() bool {
//line /snap/go/10455/src/net/netip/netip.go:634
		_go_fuzz_dep_.CoverTab[4145]++
//line /snap/go/10455/src/net/netip/netip.go:634
		return ip == IPv6Unspecified()
//line /snap/go/10455/src/net/netip/netip.go:634
		// _ = "end of CoverTab[4145]"
//line /snap/go/10455/src/net/netip/netip.go:634
	}()
//line /snap/go/10455/src/net/netip/netip.go:634
	// _ = "end of CoverTab[4144]"
}

// Prefix keeps only the top b bits of IP, producing a Prefix
//line /snap/go/10455/src/net/netip/netip.go:637
// of the specified length.
//line /snap/go/10455/src/net/netip/netip.go:637
// If ip is a zero Addr, Prefix always returns a zero Prefix and a nil error.
//line /snap/go/10455/src/net/netip/netip.go:637
// Otherwise, if bits is less than zero or greater than ip.BitLen(),
//line /snap/go/10455/src/net/netip/netip.go:637
// Prefix returns an error.
//line /snap/go/10455/src/net/netip/netip.go:642
func (ip Addr) Prefix(b int) (Prefix, error) {
//line /snap/go/10455/src/net/netip/netip.go:642
	_go_fuzz_dep_.CoverTab[4146]++
							if b < 0 {
//line /snap/go/10455/src/net/netip/netip.go:643
		_go_fuzz_dep_.CoverTab[527146]++
//line /snap/go/10455/src/net/netip/netip.go:643
		_go_fuzz_dep_.CoverTab[4149]++
								return Prefix{}, errors.New("negative Prefix bits")
//line /snap/go/10455/src/net/netip/netip.go:644
		// _ = "end of CoverTab[4149]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:645
		_go_fuzz_dep_.CoverTab[527147]++
//line /snap/go/10455/src/net/netip/netip.go:645
		_go_fuzz_dep_.CoverTab[4150]++
//line /snap/go/10455/src/net/netip/netip.go:645
		// _ = "end of CoverTab[4150]"
//line /snap/go/10455/src/net/netip/netip.go:645
	}
//line /snap/go/10455/src/net/netip/netip.go:645
	// _ = "end of CoverTab[4146]"
//line /snap/go/10455/src/net/netip/netip.go:645
	_go_fuzz_dep_.CoverTab[4147]++
							effectiveBits := b
							switch ip.z {
	case z0:
//line /snap/go/10455/src/net/netip/netip.go:648
		_go_fuzz_dep_.CoverTab[527148]++
//line /snap/go/10455/src/net/netip/netip.go:648
		_go_fuzz_dep_.CoverTab[4151]++
								return Prefix{}, nil
//line /snap/go/10455/src/net/netip/netip.go:649
		// _ = "end of CoverTab[4151]"
	case z4:
//line /snap/go/10455/src/net/netip/netip.go:650
		_go_fuzz_dep_.CoverTab[527149]++
//line /snap/go/10455/src/net/netip/netip.go:650
		_go_fuzz_dep_.CoverTab[4152]++
								if b > 32 {
//line /snap/go/10455/src/net/netip/netip.go:651
			_go_fuzz_dep_.CoverTab[527151]++
//line /snap/go/10455/src/net/netip/netip.go:651
			_go_fuzz_dep_.CoverTab[4155]++
									return Prefix{}, errors.New("prefix length " + itoa.Itoa(b) + " too large for IPv4")
//line /snap/go/10455/src/net/netip/netip.go:652
			// _ = "end of CoverTab[4155]"
		} else {
//line /snap/go/10455/src/net/netip/netip.go:653
			_go_fuzz_dep_.CoverTab[527152]++
//line /snap/go/10455/src/net/netip/netip.go:653
			_go_fuzz_dep_.CoverTab[4156]++
//line /snap/go/10455/src/net/netip/netip.go:653
			// _ = "end of CoverTab[4156]"
//line /snap/go/10455/src/net/netip/netip.go:653
		}
//line /snap/go/10455/src/net/netip/netip.go:653
		// _ = "end of CoverTab[4152]"
//line /snap/go/10455/src/net/netip/netip.go:653
		_go_fuzz_dep_.CoverTab[4153]++
								effectiveBits += 96
//line /snap/go/10455/src/net/netip/netip.go:654
		// _ = "end of CoverTab[4153]"
	default:
//line /snap/go/10455/src/net/netip/netip.go:655
		_go_fuzz_dep_.CoverTab[527150]++
//line /snap/go/10455/src/net/netip/netip.go:655
		_go_fuzz_dep_.CoverTab[4154]++
								if b > 128 {
//line /snap/go/10455/src/net/netip/netip.go:656
			_go_fuzz_dep_.CoverTab[527153]++
//line /snap/go/10455/src/net/netip/netip.go:656
			_go_fuzz_dep_.CoverTab[4157]++
									return Prefix{}, errors.New("prefix length " + itoa.Itoa(b) + " too large for IPv6")
//line /snap/go/10455/src/net/netip/netip.go:657
			// _ = "end of CoverTab[4157]"
		} else {
//line /snap/go/10455/src/net/netip/netip.go:658
			_go_fuzz_dep_.CoverTab[527154]++
//line /snap/go/10455/src/net/netip/netip.go:658
			_go_fuzz_dep_.CoverTab[4158]++
//line /snap/go/10455/src/net/netip/netip.go:658
			// _ = "end of CoverTab[4158]"
//line /snap/go/10455/src/net/netip/netip.go:658
		}
//line /snap/go/10455/src/net/netip/netip.go:658
		// _ = "end of CoverTab[4154]"
	}
//line /snap/go/10455/src/net/netip/netip.go:659
	// _ = "end of CoverTab[4147]"
//line /snap/go/10455/src/net/netip/netip.go:659
	_go_fuzz_dep_.CoverTab[4148]++
							ip.addr = ip.addr.and(mask6(effectiveBits))
							return PrefixFrom(ip, b), nil
//line /snap/go/10455/src/net/netip/netip.go:661
	// _ = "end of CoverTab[4148]"
}

const (
	netIPv4len	= 4
	netIPv6len	= 16
)

// As16 returns the IP address in its 16-byte representation.
//line /snap/go/10455/src/net/netip/netip.go:669
// IPv4 addresses are returned as IPv4-mapped IPv6 addresses.
//line /snap/go/10455/src/net/netip/netip.go:669
// IPv6 addresses with zones are returned without their zone (use the
//line /snap/go/10455/src/net/netip/netip.go:669
// Zone method to get it).
//line /snap/go/10455/src/net/netip/netip.go:669
// The ip zero value returns all zeroes.
//line /snap/go/10455/src/net/netip/netip.go:674
func (ip Addr) As16() (a16 [16]byte) {
//line /snap/go/10455/src/net/netip/netip.go:674
	_go_fuzz_dep_.CoverTab[4159]++
							bePutUint64(a16[:8], ip.addr.hi)
							bePutUint64(a16[8:], ip.addr.lo)
							return a16
//line /snap/go/10455/src/net/netip/netip.go:677
	// _ = "end of CoverTab[4159]"
}

// As4 returns an IPv4 or IPv4-in-IPv6 address in its 4-byte representation.
//line /snap/go/10455/src/net/netip/netip.go:680
// If ip is the zero Addr or an IPv6 address, As4 panics.
//line /snap/go/10455/src/net/netip/netip.go:680
// Note that 0.0.0.0 is not the zero Addr.
//line /snap/go/10455/src/net/netip/netip.go:683
func (ip Addr) As4() (a4 [4]byte) {
//line /snap/go/10455/src/net/netip/netip.go:683
	_go_fuzz_dep_.CoverTab[4160]++
							if ip.z == z4 || func() bool {
//line /snap/go/10455/src/net/netip/netip.go:684
		_go_fuzz_dep_.CoverTab[4163]++
//line /snap/go/10455/src/net/netip/netip.go:684
		return ip.Is4In6()
//line /snap/go/10455/src/net/netip/netip.go:684
		// _ = "end of CoverTab[4163]"
//line /snap/go/10455/src/net/netip/netip.go:684
	}() {
//line /snap/go/10455/src/net/netip/netip.go:684
		_go_fuzz_dep_.CoverTab[527155]++
//line /snap/go/10455/src/net/netip/netip.go:684
		_go_fuzz_dep_.CoverTab[4164]++
								bePutUint32(a4[:], uint32(ip.addr.lo))
								return a4
//line /snap/go/10455/src/net/netip/netip.go:686
		// _ = "end of CoverTab[4164]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:687
		_go_fuzz_dep_.CoverTab[527156]++
//line /snap/go/10455/src/net/netip/netip.go:687
		_go_fuzz_dep_.CoverTab[4165]++
//line /snap/go/10455/src/net/netip/netip.go:687
		// _ = "end of CoverTab[4165]"
//line /snap/go/10455/src/net/netip/netip.go:687
	}
//line /snap/go/10455/src/net/netip/netip.go:687
	// _ = "end of CoverTab[4160]"
//line /snap/go/10455/src/net/netip/netip.go:687
	_go_fuzz_dep_.CoverTab[4161]++
							if ip.z == z0 {
//line /snap/go/10455/src/net/netip/netip.go:688
		_go_fuzz_dep_.CoverTab[527157]++
//line /snap/go/10455/src/net/netip/netip.go:688
		_go_fuzz_dep_.CoverTab[4166]++
								panic("As4 called on IP zero value")
//line /snap/go/10455/src/net/netip/netip.go:689
		// _ = "end of CoverTab[4166]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:690
		_go_fuzz_dep_.CoverTab[527158]++
//line /snap/go/10455/src/net/netip/netip.go:690
		_go_fuzz_dep_.CoverTab[4167]++
//line /snap/go/10455/src/net/netip/netip.go:690
		// _ = "end of CoverTab[4167]"
//line /snap/go/10455/src/net/netip/netip.go:690
	}
//line /snap/go/10455/src/net/netip/netip.go:690
	// _ = "end of CoverTab[4161]"
//line /snap/go/10455/src/net/netip/netip.go:690
	_go_fuzz_dep_.CoverTab[4162]++
							panic("As4 called on IPv6 address")
//line /snap/go/10455/src/net/netip/netip.go:691
	// _ = "end of CoverTab[4162]"
}

// AsSlice returns an IPv4 or IPv6 address in its respective 4-byte or 16-byte representation.
func (ip Addr) AsSlice() []byte {
//line /snap/go/10455/src/net/netip/netip.go:695
	_go_fuzz_dep_.CoverTab[4168]++
							switch ip.z {
	case z0:
//line /snap/go/10455/src/net/netip/netip.go:697
		_go_fuzz_dep_.CoverTab[527159]++
//line /snap/go/10455/src/net/netip/netip.go:697
		_go_fuzz_dep_.CoverTab[4169]++
								return nil
//line /snap/go/10455/src/net/netip/netip.go:698
		// _ = "end of CoverTab[4169]"
	case z4:
//line /snap/go/10455/src/net/netip/netip.go:699
		_go_fuzz_dep_.CoverTab[527160]++
//line /snap/go/10455/src/net/netip/netip.go:699
		_go_fuzz_dep_.CoverTab[4170]++
								var ret [4]byte
								bePutUint32(ret[:], uint32(ip.addr.lo))
								return ret[:]
//line /snap/go/10455/src/net/netip/netip.go:702
		// _ = "end of CoverTab[4170]"
	default:
//line /snap/go/10455/src/net/netip/netip.go:703
		_go_fuzz_dep_.CoverTab[527161]++
//line /snap/go/10455/src/net/netip/netip.go:703
		_go_fuzz_dep_.CoverTab[4171]++
								var ret [16]byte
								bePutUint64(ret[:8], ip.addr.hi)
								bePutUint64(ret[8:], ip.addr.lo)
								return ret[:]
//line /snap/go/10455/src/net/netip/netip.go:707
		// _ = "end of CoverTab[4171]"
	}
//line /snap/go/10455/src/net/netip/netip.go:708
	// _ = "end of CoverTab[4168]"
}

// Next returns the address following ip.
//line /snap/go/10455/src/net/netip/netip.go:711
// If there is none, it returns the zero Addr.
//line /snap/go/10455/src/net/netip/netip.go:713
func (ip Addr) Next() Addr {
//line /snap/go/10455/src/net/netip/netip.go:713
	_go_fuzz_dep_.CoverTab[4172]++
							ip.addr = ip.addr.addOne()
							if ip.Is4() {
//line /snap/go/10455/src/net/netip/netip.go:715
		_go_fuzz_dep_.CoverTab[527162]++
//line /snap/go/10455/src/net/netip/netip.go:715
		_go_fuzz_dep_.CoverTab[4174]++
								if uint32(ip.addr.lo) == 0 {
//line /snap/go/10455/src/net/netip/netip.go:716
			_go_fuzz_dep_.CoverTab[527164]++
//line /snap/go/10455/src/net/netip/netip.go:716
			_go_fuzz_dep_.CoverTab[4175]++

									return Addr{}
//line /snap/go/10455/src/net/netip/netip.go:718
			// _ = "end of CoverTab[4175]"
		} else {
//line /snap/go/10455/src/net/netip/netip.go:719
			_go_fuzz_dep_.CoverTab[527165]++
//line /snap/go/10455/src/net/netip/netip.go:719
			_go_fuzz_dep_.CoverTab[4176]++
//line /snap/go/10455/src/net/netip/netip.go:719
			// _ = "end of CoverTab[4176]"
//line /snap/go/10455/src/net/netip/netip.go:719
		}
//line /snap/go/10455/src/net/netip/netip.go:719
		// _ = "end of CoverTab[4174]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:720
		_go_fuzz_dep_.CoverTab[527163]++
//line /snap/go/10455/src/net/netip/netip.go:720
		_go_fuzz_dep_.CoverTab[4177]++
								if ip.addr.isZero() {
//line /snap/go/10455/src/net/netip/netip.go:721
			_go_fuzz_dep_.CoverTab[527166]++
//line /snap/go/10455/src/net/netip/netip.go:721
			_go_fuzz_dep_.CoverTab[4178]++

									return Addr{}
//line /snap/go/10455/src/net/netip/netip.go:723
			// _ = "end of CoverTab[4178]"
		} else {
//line /snap/go/10455/src/net/netip/netip.go:724
			_go_fuzz_dep_.CoverTab[527167]++
//line /snap/go/10455/src/net/netip/netip.go:724
			_go_fuzz_dep_.CoverTab[4179]++
//line /snap/go/10455/src/net/netip/netip.go:724
			// _ = "end of CoverTab[4179]"
//line /snap/go/10455/src/net/netip/netip.go:724
		}
//line /snap/go/10455/src/net/netip/netip.go:724
		// _ = "end of CoverTab[4177]"
	}
//line /snap/go/10455/src/net/netip/netip.go:725
	// _ = "end of CoverTab[4172]"
//line /snap/go/10455/src/net/netip/netip.go:725
	_go_fuzz_dep_.CoverTab[4173]++
							return ip
//line /snap/go/10455/src/net/netip/netip.go:726
	// _ = "end of CoverTab[4173]"
}

// Prev returns the IP before ip.
//line /snap/go/10455/src/net/netip/netip.go:729
// If there is none, it returns the IP zero value.
//line /snap/go/10455/src/net/netip/netip.go:731
func (ip Addr) Prev() Addr {
//line /snap/go/10455/src/net/netip/netip.go:731
	_go_fuzz_dep_.CoverTab[4180]++
							if ip.Is4() {
//line /snap/go/10455/src/net/netip/netip.go:732
		_go_fuzz_dep_.CoverTab[527168]++
//line /snap/go/10455/src/net/netip/netip.go:732
		_go_fuzz_dep_.CoverTab[4182]++
								if uint32(ip.addr.lo) == 0 {
//line /snap/go/10455/src/net/netip/netip.go:733
			_go_fuzz_dep_.CoverTab[527170]++
//line /snap/go/10455/src/net/netip/netip.go:733
			_go_fuzz_dep_.CoverTab[4183]++
									return Addr{}
//line /snap/go/10455/src/net/netip/netip.go:734
			// _ = "end of CoverTab[4183]"
		} else {
//line /snap/go/10455/src/net/netip/netip.go:735
			_go_fuzz_dep_.CoverTab[527171]++
//line /snap/go/10455/src/net/netip/netip.go:735
			_go_fuzz_dep_.CoverTab[4184]++
//line /snap/go/10455/src/net/netip/netip.go:735
			// _ = "end of CoverTab[4184]"
//line /snap/go/10455/src/net/netip/netip.go:735
		}
//line /snap/go/10455/src/net/netip/netip.go:735
		// _ = "end of CoverTab[4182]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:736
		_go_fuzz_dep_.CoverTab[527169]++
//line /snap/go/10455/src/net/netip/netip.go:736
		_go_fuzz_dep_.CoverTab[4185]++
//line /snap/go/10455/src/net/netip/netip.go:736
		if ip.addr.isZero() {
//line /snap/go/10455/src/net/netip/netip.go:736
			_go_fuzz_dep_.CoverTab[527172]++
//line /snap/go/10455/src/net/netip/netip.go:736
			_go_fuzz_dep_.CoverTab[4186]++
									return Addr{}
//line /snap/go/10455/src/net/netip/netip.go:737
			// _ = "end of CoverTab[4186]"
		} else {
//line /snap/go/10455/src/net/netip/netip.go:738
			_go_fuzz_dep_.CoverTab[527173]++
//line /snap/go/10455/src/net/netip/netip.go:738
			_go_fuzz_dep_.CoverTab[4187]++
//line /snap/go/10455/src/net/netip/netip.go:738
			// _ = "end of CoverTab[4187]"
//line /snap/go/10455/src/net/netip/netip.go:738
		}
//line /snap/go/10455/src/net/netip/netip.go:738
		// _ = "end of CoverTab[4185]"
//line /snap/go/10455/src/net/netip/netip.go:738
	}
//line /snap/go/10455/src/net/netip/netip.go:738
	// _ = "end of CoverTab[4180]"
//line /snap/go/10455/src/net/netip/netip.go:738
	_go_fuzz_dep_.CoverTab[4181]++
							ip.addr = ip.addr.subOne()
							return ip
//line /snap/go/10455/src/net/netip/netip.go:740
	// _ = "end of CoverTab[4181]"
}

// String returns the string form of the IP address ip.
//line /snap/go/10455/src/net/netip/netip.go:743
// It returns one of 5 forms:
//line /snap/go/10455/src/net/netip/netip.go:743
//
//line /snap/go/10455/src/net/netip/netip.go:743
//   - "invalid IP", if ip is the zero Addr
//line /snap/go/10455/src/net/netip/netip.go:743
//   - IPv4 dotted decimal ("192.0.2.1")
//line /snap/go/10455/src/net/netip/netip.go:743
//   - IPv6 ("2001:db8::1")
//line /snap/go/10455/src/net/netip/netip.go:743
//   - "::ffff:1.2.3.4" (if Is4In6)
//line /snap/go/10455/src/net/netip/netip.go:743
//   - IPv6 with zone ("fe80:db8::1%eth0")
//line /snap/go/10455/src/net/netip/netip.go:743
//
//line /snap/go/10455/src/net/netip/netip.go:743
// Note that unlike package net's IP.String method,
//line /snap/go/10455/src/net/netip/netip.go:743
// IPv4-mapped IPv6 addresses format with a "::ffff:"
//line /snap/go/10455/src/net/netip/netip.go:743
// prefix before the dotted quad.
//line /snap/go/10455/src/net/netip/netip.go:755
func (ip Addr) String() string {
//line /snap/go/10455/src/net/netip/netip.go:755
	_go_fuzz_dep_.CoverTab[4188]++
							switch ip.z {
	case z0:
//line /snap/go/10455/src/net/netip/netip.go:757
		_go_fuzz_dep_.CoverTab[527174]++
//line /snap/go/10455/src/net/netip/netip.go:757
		_go_fuzz_dep_.CoverTab[4189]++
								return "invalid IP"
//line /snap/go/10455/src/net/netip/netip.go:758
		// _ = "end of CoverTab[4189]"
	case z4:
//line /snap/go/10455/src/net/netip/netip.go:759
		_go_fuzz_dep_.CoverTab[527175]++
//line /snap/go/10455/src/net/netip/netip.go:759
		_go_fuzz_dep_.CoverTab[4190]++
								return ip.string4()
//line /snap/go/10455/src/net/netip/netip.go:760
		// _ = "end of CoverTab[4190]"
	default:
//line /snap/go/10455/src/net/netip/netip.go:761
		_go_fuzz_dep_.CoverTab[527176]++
//line /snap/go/10455/src/net/netip/netip.go:761
		_go_fuzz_dep_.CoverTab[4191]++
								if ip.Is4In6() {
//line /snap/go/10455/src/net/netip/netip.go:762
			_go_fuzz_dep_.CoverTab[527177]++
//line /snap/go/10455/src/net/netip/netip.go:762
			_go_fuzz_dep_.CoverTab[4193]++
									if z := ip.Zone(); z != "" {
//line /snap/go/10455/src/net/netip/netip.go:763
				_go_fuzz_dep_.CoverTab[527179]++
//line /snap/go/10455/src/net/netip/netip.go:763
				_go_fuzz_dep_.CoverTab[4194]++
										return "::ffff:" + ip.Unmap().string4() + "%" + z
//line /snap/go/10455/src/net/netip/netip.go:764
				// _ = "end of CoverTab[4194]"
			} else {
//line /snap/go/10455/src/net/netip/netip.go:765
				_go_fuzz_dep_.CoverTab[527180]++
//line /snap/go/10455/src/net/netip/netip.go:765
				_go_fuzz_dep_.CoverTab[4195]++
										return "::ffff:" + ip.Unmap().string4()
//line /snap/go/10455/src/net/netip/netip.go:766
				// _ = "end of CoverTab[4195]"
			}
//line /snap/go/10455/src/net/netip/netip.go:767
			// _ = "end of CoverTab[4193]"
		} else {
//line /snap/go/10455/src/net/netip/netip.go:768
			_go_fuzz_dep_.CoverTab[527178]++
//line /snap/go/10455/src/net/netip/netip.go:768
			_go_fuzz_dep_.CoverTab[4196]++
//line /snap/go/10455/src/net/netip/netip.go:768
			// _ = "end of CoverTab[4196]"
//line /snap/go/10455/src/net/netip/netip.go:768
		}
//line /snap/go/10455/src/net/netip/netip.go:768
		// _ = "end of CoverTab[4191]"
//line /snap/go/10455/src/net/netip/netip.go:768
		_go_fuzz_dep_.CoverTab[4192]++
								return ip.string6()
//line /snap/go/10455/src/net/netip/netip.go:769
		// _ = "end of CoverTab[4192]"
	}
//line /snap/go/10455/src/net/netip/netip.go:770
	// _ = "end of CoverTab[4188]"
}

// AppendTo appends a text encoding of ip,
//line /snap/go/10455/src/net/netip/netip.go:773
// as generated by MarshalText,
//line /snap/go/10455/src/net/netip/netip.go:773
// to b and returns the extended buffer.
//line /snap/go/10455/src/net/netip/netip.go:776
func (ip Addr) AppendTo(b []byte) []byte {
//line /snap/go/10455/src/net/netip/netip.go:776
	_go_fuzz_dep_.CoverTab[4197]++
							switch ip.z {
	case z0:
//line /snap/go/10455/src/net/netip/netip.go:778
		_go_fuzz_dep_.CoverTab[527181]++
//line /snap/go/10455/src/net/netip/netip.go:778
		_go_fuzz_dep_.CoverTab[4198]++
								return b
//line /snap/go/10455/src/net/netip/netip.go:779
		// _ = "end of CoverTab[4198]"
	case z4:
//line /snap/go/10455/src/net/netip/netip.go:780
		_go_fuzz_dep_.CoverTab[527182]++
//line /snap/go/10455/src/net/netip/netip.go:780
		_go_fuzz_dep_.CoverTab[4199]++
								return ip.appendTo4(b)
//line /snap/go/10455/src/net/netip/netip.go:781
		// _ = "end of CoverTab[4199]"
	default:
//line /snap/go/10455/src/net/netip/netip.go:782
		_go_fuzz_dep_.CoverTab[527183]++
//line /snap/go/10455/src/net/netip/netip.go:782
		_go_fuzz_dep_.CoverTab[4200]++
								if ip.Is4In6() {
//line /snap/go/10455/src/net/netip/netip.go:783
			_go_fuzz_dep_.CoverTab[527184]++
//line /snap/go/10455/src/net/netip/netip.go:783
			_go_fuzz_dep_.CoverTab[4202]++
									b = append(b, "::ffff:"...)
									b = ip.Unmap().appendTo4(b)
									if z := ip.Zone(); z != "" {
//line /snap/go/10455/src/net/netip/netip.go:786
				_go_fuzz_dep_.CoverTab[527186]++
//line /snap/go/10455/src/net/netip/netip.go:786
				_go_fuzz_dep_.CoverTab[4204]++
										b = append(b, '%')
										b = append(b, z...)
//line /snap/go/10455/src/net/netip/netip.go:788
				// _ = "end of CoverTab[4204]"
			} else {
//line /snap/go/10455/src/net/netip/netip.go:789
				_go_fuzz_dep_.CoverTab[527187]++
//line /snap/go/10455/src/net/netip/netip.go:789
				_go_fuzz_dep_.CoverTab[4205]++
//line /snap/go/10455/src/net/netip/netip.go:789
				// _ = "end of CoverTab[4205]"
//line /snap/go/10455/src/net/netip/netip.go:789
			}
//line /snap/go/10455/src/net/netip/netip.go:789
			// _ = "end of CoverTab[4202]"
//line /snap/go/10455/src/net/netip/netip.go:789
			_go_fuzz_dep_.CoverTab[4203]++
									return b
//line /snap/go/10455/src/net/netip/netip.go:790
			// _ = "end of CoverTab[4203]"
		} else {
//line /snap/go/10455/src/net/netip/netip.go:791
			_go_fuzz_dep_.CoverTab[527185]++
//line /snap/go/10455/src/net/netip/netip.go:791
			_go_fuzz_dep_.CoverTab[4206]++
//line /snap/go/10455/src/net/netip/netip.go:791
			// _ = "end of CoverTab[4206]"
//line /snap/go/10455/src/net/netip/netip.go:791
		}
//line /snap/go/10455/src/net/netip/netip.go:791
		// _ = "end of CoverTab[4200]"
//line /snap/go/10455/src/net/netip/netip.go:791
		_go_fuzz_dep_.CoverTab[4201]++
								return ip.appendTo6(b)
//line /snap/go/10455/src/net/netip/netip.go:792
		// _ = "end of CoverTab[4201]"
	}
//line /snap/go/10455/src/net/netip/netip.go:793
	// _ = "end of CoverTab[4197]"
}

// digits is a string of the hex digits from 0 to f. It's used in
//line /snap/go/10455/src/net/netip/netip.go:796
// appendDecimal and appendHex to format IP addresses.
//line /snap/go/10455/src/net/netip/netip.go:798
const digits = "0123456789abcdef"

// appendDecimal appends the decimal string representation of x to b.
func appendDecimal(b []byte, x uint8) []byte {
//line /snap/go/10455/src/net/netip/netip.go:801
	_go_fuzz_dep_.CoverTab[4207]++

//line /snap/go/10455/src/net/netip/netip.go:805
	if x >= 100 {
//line /snap/go/10455/src/net/netip/netip.go:805
		_go_fuzz_dep_.CoverTab[527188]++
//line /snap/go/10455/src/net/netip/netip.go:805
		_go_fuzz_dep_.CoverTab[4210]++
								b = append(b, digits[x/100])
//line /snap/go/10455/src/net/netip/netip.go:806
		// _ = "end of CoverTab[4210]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:807
		_go_fuzz_dep_.CoverTab[527189]++
//line /snap/go/10455/src/net/netip/netip.go:807
		_go_fuzz_dep_.CoverTab[4211]++
//line /snap/go/10455/src/net/netip/netip.go:807
		// _ = "end of CoverTab[4211]"
//line /snap/go/10455/src/net/netip/netip.go:807
	}
//line /snap/go/10455/src/net/netip/netip.go:807
	// _ = "end of CoverTab[4207]"
//line /snap/go/10455/src/net/netip/netip.go:807
	_go_fuzz_dep_.CoverTab[4208]++
							if x >= 10 {
//line /snap/go/10455/src/net/netip/netip.go:808
		_go_fuzz_dep_.CoverTab[527190]++
//line /snap/go/10455/src/net/netip/netip.go:808
		_go_fuzz_dep_.CoverTab[4212]++
								b = append(b, digits[x/10%10])
//line /snap/go/10455/src/net/netip/netip.go:809
		// _ = "end of CoverTab[4212]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:810
		_go_fuzz_dep_.CoverTab[527191]++
//line /snap/go/10455/src/net/netip/netip.go:810
		_go_fuzz_dep_.CoverTab[4213]++
//line /snap/go/10455/src/net/netip/netip.go:810
		// _ = "end of CoverTab[4213]"
//line /snap/go/10455/src/net/netip/netip.go:810
	}
//line /snap/go/10455/src/net/netip/netip.go:810
	// _ = "end of CoverTab[4208]"
//line /snap/go/10455/src/net/netip/netip.go:810
	_go_fuzz_dep_.CoverTab[4209]++
							return append(b, digits[x%10])
//line /snap/go/10455/src/net/netip/netip.go:811
	// _ = "end of CoverTab[4209]"
}

// appendHex appends the hex string representation of x to b.
func appendHex(b []byte, x uint16) []byte {
//line /snap/go/10455/src/net/netip/netip.go:815
	_go_fuzz_dep_.CoverTab[4214]++

//line /snap/go/10455/src/net/netip/netip.go:819
	if x >= 0x1000 {
//line /snap/go/10455/src/net/netip/netip.go:819
		_go_fuzz_dep_.CoverTab[527192]++
//line /snap/go/10455/src/net/netip/netip.go:819
		_go_fuzz_dep_.CoverTab[4218]++
								b = append(b, digits[x>>12])
//line /snap/go/10455/src/net/netip/netip.go:820
		// _ = "end of CoverTab[4218]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:821
		_go_fuzz_dep_.CoverTab[527193]++
//line /snap/go/10455/src/net/netip/netip.go:821
		_go_fuzz_dep_.CoverTab[4219]++
//line /snap/go/10455/src/net/netip/netip.go:821
		// _ = "end of CoverTab[4219]"
//line /snap/go/10455/src/net/netip/netip.go:821
	}
//line /snap/go/10455/src/net/netip/netip.go:821
	// _ = "end of CoverTab[4214]"
//line /snap/go/10455/src/net/netip/netip.go:821
	_go_fuzz_dep_.CoverTab[4215]++
							if x >= 0x100 {
//line /snap/go/10455/src/net/netip/netip.go:822
		_go_fuzz_dep_.CoverTab[527194]++
//line /snap/go/10455/src/net/netip/netip.go:822
		_go_fuzz_dep_.CoverTab[4220]++
								b = append(b, digits[x>>8&0xf])
//line /snap/go/10455/src/net/netip/netip.go:823
		// _ = "end of CoverTab[4220]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:824
		_go_fuzz_dep_.CoverTab[527195]++
//line /snap/go/10455/src/net/netip/netip.go:824
		_go_fuzz_dep_.CoverTab[4221]++
//line /snap/go/10455/src/net/netip/netip.go:824
		// _ = "end of CoverTab[4221]"
//line /snap/go/10455/src/net/netip/netip.go:824
	}
//line /snap/go/10455/src/net/netip/netip.go:824
	// _ = "end of CoverTab[4215]"
//line /snap/go/10455/src/net/netip/netip.go:824
	_go_fuzz_dep_.CoverTab[4216]++
							if x >= 0x10 {
//line /snap/go/10455/src/net/netip/netip.go:825
		_go_fuzz_dep_.CoverTab[527196]++
//line /snap/go/10455/src/net/netip/netip.go:825
		_go_fuzz_dep_.CoverTab[4222]++
								b = append(b, digits[x>>4&0xf])
//line /snap/go/10455/src/net/netip/netip.go:826
		// _ = "end of CoverTab[4222]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:827
		_go_fuzz_dep_.CoverTab[527197]++
//line /snap/go/10455/src/net/netip/netip.go:827
		_go_fuzz_dep_.CoverTab[4223]++
//line /snap/go/10455/src/net/netip/netip.go:827
		// _ = "end of CoverTab[4223]"
//line /snap/go/10455/src/net/netip/netip.go:827
	}
//line /snap/go/10455/src/net/netip/netip.go:827
	// _ = "end of CoverTab[4216]"
//line /snap/go/10455/src/net/netip/netip.go:827
	_go_fuzz_dep_.CoverTab[4217]++
							return append(b, digits[x&0xf])
//line /snap/go/10455/src/net/netip/netip.go:828
	// _ = "end of CoverTab[4217]"
}

// appendHexPad appends the fully padded hex string representation of x to b.
func appendHexPad(b []byte, x uint16) []byte {
//line /snap/go/10455/src/net/netip/netip.go:832
	_go_fuzz_dep_.CoverTab[4224]++
							return append(b, digits[x>>12], digits[x>>8&0xf], digits[x>>4&0xf], digits[x&0xf])
//line /snap/go/10455/src/net/netip/netip.go:833
	// _ = "end of CoverTab[4224]"
}

func (ip Addr) string4() string {
//line /snap/go/10455/src/net/netip/netip.go:836
	_go_fuzz_dep_.CoverTab[4225]++
							const max = len("255.255.255.255")
							ret := make([]byte, 0, max)
							ret = ip.appendTo4(ret)
							return string(ret)
//line /snap/go/10455/src/net/netip/netip.go:840
	// _ = "end of CoverTab[4225]"
}

func (ip Addr) appendTo4(ret []byte) []byte {
//line /snap/go/10455/src/net/netip/netip.go:843
	_go_fuzz_dep_.CoverTab[4226]++
							ret = appendDecimal(ret, ip.v4(0))
							ret = append(ret, '.')
							ret = appendDecimal(ret, ip.v4(1))
							ret = append(ret, '.')
							ret = appendDecimal(ret, ip.v4(2))
							ret = append(ret, '.')
							ret = appendDecimal(ret, ip.v4(3))
							return ret
//line /snap/go/10455/src/net/netip/netip.go:851
	// _ = "end of CoverTab[4226]"
}

// string6 formats ip in IPv6 textual representation. It follows the
//line /snap/go/10455/src/net/netip/netip.go:854
// guidelines in section 4 of RFC 5952
//line /snap/go/10455/src/net/netip/netip.go:854
// (https://tools.ietf.org/html/rfc5952#section-4): no unnecessary
//line /snap/go/10455/src/net/netip/netip.go:854
// zeros, use :: to elide the longest run of zeros, and don't use ::
//line /snap/go/10455/src/net/netip/netip.go:854
// to compact a single zero field.
//line /snap/go/10455/src/net/netip/netip.go:859
func (ip Addr) string6() string {
//line /snap/go/10455/src/net/netip/netip.go:859
	_go_fuzz_dep_.CoverTab[4227]++
	// Use a zone with a "plausibly long" name, so that most zone-ful
	// IP addresses won't require additional allocation.
	//
	// The compiler does a cool optimization here, where ret ends up
	// stack-allocated and so the only allocation this function does
	// is to construct the returned string. As such, it's okay to be a
							// bit greedy here, size-wise.
							const max = len("ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff%enp5s0")
							ret := make([]byte, 0, max)
							ret = ip.appendTo6(ret)
							return string(ret)
//line /snap/go/10455/src/net/netip/netip.go:870
	// _ = "end of CoverTab[4227]"
}

func (ip Addr) appendTo6(ret []byte) []byte {
//line /snap/go/10455/src/net/netip/netip.go:873
	_go_fuzz_dep_.CoverTab[4228]++
							zeroStart, zeroEnd := uint8(255), uint8(255)
//line /snap/go/10455/src/net/netip/netip.go:874
	_go_fuzz_dep_.CoverTab[786631] = 0
							for i := uint8(0); i < 8; i++ {
//line /snap/go/10455/src/net/netip/netip.go:875
		if _go_fuzz_dep_.CoverTab[786631] == 0 {
//line /snap/go/10455/src/net/netip/netip.go:875
			_go_fuzz_dep_.CoverTab[527353]++
//line /snap/go/10455/src/net/netip/netip.go:875
		} else {
//line /snap/go/10455/src/net/netip/netip.go:875
			_go_fuzz_dep_.CoverTab[527354]++
//line /snap/go/10455/src/net/netip/netip.go:875
		}
//line /snap/go/10455/src/net/netip/netip.go:875
		_go_fuzz_dep_.CoverTab[786631] = 1
//line /snap/go/10455/src/net/netip/netip.go:875
		_go_fuzz_dep_.CoverTab[4232]++
								j := i
//line /snap/go/10455/src/net/netip/netip.go:876
		_go_fuzz_dep_.CoverTab[786633] = 0
								for j < 8 && func() bool {
//line /snap/go/10455/src/net/netip/netip.go:877
			_go_fuzz_dep_.CoverTab[4234]++
//line /snap/go/10455/src/net/netip/netip.go:877
			return ip.v6u16(j) == 0
//line /snap/go/10455/src/net/netip/netip.go:877
			// _ = "end of CoverTab[4234]"
//line /snap/go/10455/src/net/netip/netip.go:877
		}() {
//line /snap/go/10455/src/net/netip/netip.go:877
			if _go_fuzz_dep_.CoverTab[786633] == 0 {
//line /snap/go/10455/src/net/netip/netip.go:877
				_go_fuzz_dep_.CoverTab[527361]++
//line /snap/go/10455/src/net/netip/netip.go:877
			} else {
//line /snap/go/10455/src/net/netip/netip.go:877
				_go_fuzz_dep_.CoverTab[527362]++
//line /snap/go/10455/src/net/netip/netip.go:877
			}
//line /snap/go/10455/src/net/netip/netip.go:877
			_go_fuzz_dep_.CoverTab[786633] = 1
//line /snap/go/10455/src/net/netip/netip.go:877
			_go_fuzz_dep_.CoverTab[4235]++
									j++
//line /snap/go/10455/src/net/netip/netip.go:878
			// _ = "end of CoverTab[4235]"
		}
//line /snap/go/10455/src/net/netip/netip.go:879
		if _go_fuzz_dep_.CoverTab[786633] == 0 {
//line /snap/go/10455/src/net/netip/netip.go:879
			_go_fuzz_dep_.CoverTab[527363]++
//line /snap/go/10455/src/net/netip/netip.go:879
		} else {
//line /snap/go/10455/src/net/netip/netip.go:879
			_go_fuzz_dep_.CoverTab[527364]++
//line /snap/go/10455/src/net/netip/netip.go:879
		}
//line /snap/go/10455/src/net/netip/netip.go:879
		// _ = "end of CoverTab[4232]"
//line /snap/go/10455/src/net/netip/netip.go:879
		_go_fuzz_dep_.CoverTab[4233]++
								if l := j - i; l >= 2 && func() bool {
//line /snap/go/10455/src/net/netip/netip.go:880
			_go_fuzz_dep_.CoverTab[4236]++
//line /snap/go/10455/src/net/netip/netip.go:880
			return l > zeroEnd-zeroStart
//line /snap/go/10455/src/net/netip/netip.go:880
			// _ = "end of CoverTab[4236]"
//line /snap/go/10455/src/net/netip/netip.go:880
		}() {
//line /snap/go/10455/src/net/netip/netip.go:880
			_go_fuzz_dep_.CoverTab[527198]++
//line /snap/go/10455/src/net/netip/netip.go:880
			_go_fuzz_dep_.CoverTab[4237]++
									zeroStart, zeroEnd = i, j
//line /snap/go/10455/src/net/netip/netip.go:881
			// _ = "end of CoverTab[4237]"
		} else {
//line /snap/go/10455/src/net/netip/netip.go:882
			_go_fuzz_dep_.CoverTab[527199]++
//line /snap/go/10455/src/net/netip/netip.go:882
			_go_fuzz_dep_.CoverTab[4238]++
//line /snap/go/10455/src/net/netip/netip.go:882
			// _ = "end of CoverTab[4238]"
//line /snap/go/10455/src/net/netip/netip.go:882
		}
//line /snap/go/10455/src/net/netip/netip.go:882
		// _ = "end of CoverTab[4233]"
	}
//line /snap/go/10455/src/net/netip/netip.go:883
	if _go_fuzz_dep_.CoverTab[786631] == 0 {
//line /snap/go/10455/src/net/netip/netip.go:883
		_go_fuzz_dep_.CoverTab[527355]++
//line /snap/go/10455/src/net/netip/netip.go:883
	} else {
//line /snap/go/10455/src/net/netip/netip.go:883
		_go_fuzz_dep_.CoverTab[527356]++
//line /snap/go/10455/src/net/netip/netip.go:883
	}
//line /snap/go/10455/src/net/netip/netip.go:883
	// _ = "end of CoverTab[4228]"
//line /snap/go/10455/src/net/netip/netip.go:883
	_go_fuzz_dep_.CoverTab[4229]++
//line /snap/go/10455/src/net/netip/netip.go:883
	_go_fuzz_dep_.CoverTab[786632] = 0

							for i := uint8(0); i < 8; i++ {
//line /snap/go/10455/src/net/netip/netip.go:885
		if _go_fuzz_dep_.CoverTab[786632] == 0 {
//line /snap/go/10455/src/net/netip/netip.go:885
			_go_fuzz_dep_.CoverTab[527357]++
//line /snap/go/10455/src/net/netip/netip.go:885
		} else {
//line /snap/go/10455/src/net/netip/netip.go:885
			_go_fuzz_dep_.CoverTab[527358]++
//line /snap/go/10455/src/net/netip/netip.go:885
		}
//line /snap/go/10455/src/net/netip/netip.go:885
		_go_fuzz_dep_.CoverTab[786632] = 1
//line /snap/go/10455/src/net/netip/netip.go:885
		_go_fuzz_dep_.CoverTab[4239]++
								if i == zeroStart {
//line /snap/go/10455/src/net/netip/netip.go:886
			_go_fuzz_dep_.CoverTab[527200]++
//line /snap/go/10455/src/net/netip/netip.go:886
			_go_fuzz_dep_.CoverTab[4241]++
									ret = append(ret, ':', ':')
									i = zeroEnd
									if i >= 8 {
//line /snap/go/10455/src/net/netip/netip.go:889
				_go_fuzz_dep_.CoverTab[527202]++
//line /snap/go/10455/src/net/netip/netip.go:889
				_go_fuzz_dep_.CoverTab[4242]++
										break
//line /snap/go/10455/src/net/netip/netip.go:890
				// _ = "end of CoverTab[4242]"
			} else {
//line /snap/go/10455/src/net/netip/netip.go:891
				_go_fuzz_dep_.CoverTab[527203]++
//line /snap/go/10455/src/net/netip/netip.go:891
				_go_fuzz_dep_.CoverTab[4243]++
//line /snap/go/10455/src/net/netip/netip.go:891
				// _ = "end of CoverTab[4243]"
//line /snap/go/10455/src/net/netip/netip.go:891
			}
//line /snap/go/10455/src/net/netip/netip.go:891
			// _ = "end of CoverTab[4241]"
		} else {
//line /snap/go/10455/src/net/netip/netip.go:892
			_go_fuzz_dep_.CoverTab[527201]++
//line /snap/go/10455/src/net/netip/netip.go:892
			_go_fuzz_dep_.CoverTab[4244]++
//line /snap/go/10455/src/net/netip/netip.go:892
			if i > 0 {
//line /snap/go/10455/src/net/netip/netip.go:892
				_go_fuzz_dep_.CoverTab[527204]++
//line /snap/go/10455/src/net/netip/netip.go:892
				_go_fuzz_dep_.CoverTab[4245]++
										ret = append(ret, ':')
//line /snap/go/10455/src/net/netip/netip.go:893
				// _ = "end of CoverTab[4245]"
			} else {
//line /snap/go/10455/src/net/netip/netip.go:894
				_go_fuzz_dep_.CoverTab[527205]++
//line /snap/go/10455/src/net/netip/netip.go:894
				_go_fuzz_dep_.CoverTab[4246]++
//line /snap/go/10455/src/net/netip/netip.go:894
				// _ = "end of CoverTab[4246]"
//line /snap/go/10455/src/net/netip/netip.go:894
			}
//line /snap/go/10455/src/net/netip/netip.go:894
			// _ = "end of CoverTab[4244]"
//line /snap/go/10455/src/net/netip/netip.go:894
		}
//line /snap/go/10455/src/net/netip/netip.go:894
		// _ = "end of CoverTab[4239]"
//line /snap/go/10455/src/net/netip/netip.go:894
		_go_fuzz_dep_.CoverTab[4240]++

								ret = appendHex(ret, ip.v6u16(i))
//line /snap/go/10455/src/net/netip/netip.go:896
		// _ = "end of CoverTab[4240]"
	}
//line /snap/go/10455/src/net/netip/netip.go:897
	if _go_fuzz_dep_.CoverTab[786632] == 0 {
//line /snap/go/10455/src/net/netip/netip.go:897
		_go_fuzz_dep_.CoverTab[527359]++
//line /snap/go/10455/src/net/netip/netip.go:897
	} else {
//line /snap/go/10455/src/net/netip/netip.go:897
		_go_fuzz_dep_.CoverTab[527360]++
//line /snap/go/10455/src/net/netip/netip.go:897
	}
//line /snap/go/10455/src/net/netip/netip.go:897
	// _ = "end of CoverTab[4229]"
//line /snap/go/10455/src/net/netip/netip.go:897
	_go_fuzz_dep_.CoverTab[4230]++

							if ip.z != z6noz {
//line /snap/go/10455/src/net/netip/netip.go:899
		_go_fuzz_dep_.CoverTab[527206]++
//line /snap/go/10455/src/net/netip/netip.go:899
		_go_fuzz_dep_.CoverTab[4247]++
								ret = append(ret, '%')
								ret = append(ret, ip.Zone()...)
//line /snap/go/10455/src/net/netip/netip.go:901
		// _ = "end of CoverTab[4247]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:902
		_go_fuzz_dep_.CoverTab[527207]++
//line /snap/go/10455/src/net/netip/netip.go:902
		_go_fuzz_dep_.CoverTab[4248]++
//line /snap/go/10455/src/net/netip/netip.go:902
		// _ = "end of CoverTab[4248]"
//line /snap/go/10455/src/net/netip/netip.go:902
	}
//line /snap/go/10455/src/net/netip/netip.go:902
	// _ = "end of CoverTab[4230]"
//line /snap/go/10455/src/net/netip/netip.go:902
	_go_fuzz_dep_.CoverTab[4231]++
							return ret
//line /snap/go/10455/src/net/netip/netip.go:903
	// _ = "end of CoverTab[4231]"
}

// StringExpanded is like String but IPv6 addresses are expanded with leading
//line /snap/go/10455/src/net/netip/netip.go:906
// zeroes and no "::" compression. For example, "2001:db8::1" becomes
//line /snap/go/10455/src/net/netip/netip.go:906
// "2001:0db8:0000:0000:0000:0000:0000:0001".
//line /snap/go/10455/src/net/netip/netip.go:909
func (ip Addr) StringExpanded() string {
//line /snap/go/10455/src/net/netip/netip.go:909
	_go_fuzz_dep_.CoverTab[4249]++
							switch ip.z {
	case z0, z4:
//line /snap/go/10455/src/net/netip/netip.go:911
		_go_fuzz_dep_.CoverTab[527208]++
//line /snap/go/10455/src/net/netip/netip.go:911
		_go_fuzz_dep_.CoverTab[4253]++
								return ip.String()
//line /snap/go/10455/src/net/netip/netip.go:912
		// _ = "end of CoverTab[4253]"
//line /snap/go/10455/src/net/netip/netip.go:912
	default:
//line /snap/go/10455/src/net/netip/netip.go:912
		_go_fuzz_dep_.CoverTab[527209]++
//line /snap/go/10455/src/net/netip/netip.go:912
		_go_fuzz_dep_.CoverTab[4254]++
//line /snap/go/10455/src/net/netip/netip.go:912
		// _ = "end of CoverTab[4254]"
	}
//line /snap/go/10455/src/net/netip/netip.go:913
	// _ = "end of CoverTab[4249]"
//line /snap/go/10455/src/net/netip/netip.go:913
	_go_fuzz_dep_.CoverTab[4250]++

							const size = len("ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff")
							ret := make([]byte, 0, size)
//line /snap/go/10455/src/net/netip/netip.go:916
	_go_fuzz_dep_.CoverTab[786634] = 0
							for i := uint8(0); i < 8; i++ {
//line /snap/go/10455/src/net/netip/netip.go:917
		if _go_fuzz_dep_.CoverTab[786634] == 0 {
//line /snap/go/10455/src/net/netip/netip.go:917
			_go_fuzz_dep_.CoverTab[527365]++
//line /snap/go/10455/src/net/netip/netip.go:917
		} else {
//line /snap/go/10455/src/net/netip/netip.go:917
			_go_fuzz_dep_.CoverTab[527366]++
//line /snap/go/10455/src/net/netip/netip.go:917
		}
//line /snap/go/10455/src/net/netip/netip.go:917
		_go_fuzz_dep_.CoverTab[786634] = 1
//line /snap/go/10455/src/net/netip/netip.go:917
		_go_fuzz_dep_.CoverTab[4255]++
								if i > 0 {
//line /snap/go/10455/src/net/netip/netip.go:918
			_go_fuzz_dep_.CoverTab[527210]++
//line /snap/go/10455/src/net/netip/netip.go:918
			_go_fuzz_dep_.CoverTab[4257]++
									ret = append(ret, ':')
//line /snap/go/10455/src/net/netip/netip.go:919
			// _ = "end of CoverTab[4257]"
		} else {
//line /snap/go/10455/src/net/netip/netip.go:920
			_go_fuzz_dep_.CoverTab[527211]++
//line /snap/go/10455/src/net/netip/netip.go:920
			_go_fuzz_dep_.CoverTab[4258]++
//line /snap/go/10455/src/net/netip/netip.go:920
			// _ = "end of CoverTab[4258]"
//line /snap/go/10455/src/net/netip/netip.go:920
		}
//line /snap/go/10455/src/net/netip/netip.go:920
		// _ = "end of CoverTab[4255]"
//line /snap/go/10455/src/net/netip/netip.go:920
		_go_fuzz_dep_.CoverTab[4256]++

								ret = appendHexPad(ret, ip.v6u16(i))
//line /snap/go/10455/src/net/netip/netip.go:922
		// _ = "end of CoverTab[4256]"
	}
//line /snap/go/10455/src/net/netip/netip.go:923
	if _go_fuzz_dep_.CoverTab[786634] == 0 {
//line /snap/go/10455/src/net/netip/netip.go:923
		_go_fuzz_dep_.CoverTab[527367]++
//line /snap/go/10455/src/net/netip/netip.go:923
	} else {
//line /snap/go/10455/src/net/netip/netip.go:923
		_go_fuzz_dep_.CoverTab[527368]++
//line /snap/go/10455/src/net/netip/netip.go:923
	}
//line /snap/go/10455/src/net/netip/netip.go:923
	// _ = "end of CoverTab[4250]"
//line /snap/go/10455/src/net/netip/netip.go:923
	_go_fuzz_dep_.CoverTab[4251]++

							if ip.z != z6noz {
//line /snap/go/10455/src/net/netip/netip.go:925
		_go_fuzz_dep_.CoverTab[527212]++
//line /snap/go/10455/src/net/netip/netip.go:925
		_go_fuzz_dep_.CoverTab[4259]++

//line /snap/go/10455/src/net/netip/netip.go:928
		ret = append(ret, '%')
								ret = append(ret, ip.Zone()...)
//line /snap/go/10455/src/net/netip/netip.go:929
		// _ = "end of CoverTab[4259]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:930
		_go_fuzz_dep_.CoverTab[527213]++
//line /snap/go/10455/src/net/netip/netip.go:930
		_go_fuzz_dep_.CoverTab[4260]++
//line /snap/go/10455/src/net/netip/netip.go:930
		// _ = "end of CoverTab[4260]"
//line /snap/go/10455/src/net/netip/netip.go:930
	}
//line /snap/go/10455/src/net/netip/netip.go:930
	// _ = "end of CoverTab[4251]"
//line /snap/go/10455/src/net/netip/netip.go:930
	_go_fuzz_dep_.CoverTab[4252]++
							return string(ret)
//line /snap/go/10455/src/net/netip/netip.go:931
	// _ = "end of CoverTab[4252]"
}

// MarshalText implements the encoding.TextMarshaler interface,
//line /snap/go/10455/src/net/netip/netip.go:934
// The encoding is the same as returned by String, with one exception:
//line /snap/go/10455/src/net/netip/netip.go:934
// If ip is the zero Addr, the encoding is the empty string.
//line /snap/go/10455/src/net/netip/netip.go:937
func (ip Addr) MarshalText() ([]byte, error) {
//line /snap/go/10455/src/net/netip/netip.go:937
	_go_fuzz_dep_.CoverTab[4261]++
							switch ip.z {
	case z0:
//line /snap/go/10455/src/net/netip/netip.go:939
		_go_fuzz_dep_.CoverTab[527214]++
//line /snap/go/10455/src/net/netip/netip.go:939
		_go_fuzz_dep_.CoverTab[4262]++
								return []byte(""), nil
//line /snap/go/10455/src/net/netip/netip.go:940
		// _ = "end of CoverTab[4262]"
	case z4:
//line /snap/go/10455/src/net/netip/netip.go:941
		_go_fuzz_dep_.CoverTab[527215]++
//line /snap/go/10455/src/net/netip/netip.go:941
		_go_fuzz_dep_.CoverTab[4263]++
								max := len("255.255.255.255")
								b := make([]byte, 0, max)
								return ip.appendTo4(b), nil
//line /snap/go/10455/src/net/netip/netip.go:944
		// _ = "end of CoverTab[4263]"
	default:
//line /snap/go/10455/src/net/netip/netip.go:945
		_go_fuzz_dep_.CoverTab[527216]++
//line /snap/go/10455/src/net/netip/netip.go:945
		_go_fuzz_dep_.CoverTab[4264]++
								max := len("ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff%enp5s0")
								b := make([]byte, 0, max)
								if ip.Is4In6() {
//line /snap/go/10455/src/net/netip/netip.go:948
			_go_fuzz_dep_.CoverTab[527217]++
//line /snap/go/10455/src/net/netip/netip.go:948
			_go_fuzz_dep_.CoverTab[4266]++
									b = append(b, "::ffff:"...)
									b = ip.Unmap().appendTo4(b)
									if z := ip.Zone(); z != "" {
//line /snap/go/10455/src/net/netip/netip.go:951
				_go_fuzz_dep_.CoverTab[527219]++
//line /snap/go/10455/src/net/netip/netip.go:951
				_go_fuzz_dep_.CoverTab[4268]++
										b = append(b, '%')
										b = append(b, z...)
//line /snap/go/10455/src/net/netip/netip.go:953
				// _ = "end of CoverTab[4268]"
			} else {
//line /snap/go/10455/src/net/netip/netip.go:954
				_go_fuzz_dep_.CoverTab[527220]++
//line /snap/go/10455/src/net/netip/netip.go:954
				_go_fuzz_dep_.CoverTab[4269]++
//line /snap/go/10455/src/net/netip/netip.go:954
				// _ = "end of CoverTab[4269]"
//line /snap/go/10455/src/net/netip/netip.go:954
			}
//line /snap/go/10455/src/net/netip/netip.go:954
			// _ = "end of CoverTab[4266]"
//line /snap/go/10455/src/net/netip/netip.go:954
			_go_fuzz_dep_.CoverTab[4267]++
									return b, nil
//line /snap/go/10455/src/net/netip/netip.go:955
			// _ = "end of CoverTab[4267]"
		} else {
//line /snap/go/10455/src/net/netip/netip.go:956
			_go_fuzz_dep_.CoverTab[527218]++
//line /snap/go/10455/src/net/netip/netip.go:956
			_go_fuzz_dep_.CoverTab[4270]++
//line /snap/go/10455/src/net/netip/netip.go:956
			// _ = "end of CoverTab[4270]"
//line /snap/go/10455/src/net/netip/netip.go:956
		}
//line /snap/go/10455/src/net/netip/netip.go:956
		// _ = "end of CoverTab[4264]"
//line /snap/go/10455/src/net/netip/netip.go:956
		_go_fuzz_dep_.CoverTab[4265]++
								return ip.appendTo6(b), nil
//line /snap/go/10455/src/net/netip/netip.go:957
		// _ = "end of CoverTab[4265]"
	}
//line /snap/go/10455/src/net/netip/netip.go:958
	// _ = "end of CoverTab[4261]"

}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
//line /snap/go/10455/src/net/netip/netip.go:962
// The IP address is expected in a form accepted by ParseAddr.
//line /snap/go/10455/src/net/netip/netip.go:962
//
//line /snap/go/10455/src/net/netip/netip.go:962
// If text is empty, UnmarshalText sets *ip to the zero Addr and
//line /snap/go/10455/src/net/netip/netip.go:962
// returns no error.
//line /snap/go/10455/src/net/netip/netip.go:967
func (ip *Addr) UnmarshalText(text []byte) error {
//line /snap/go/10455/src/net/netip/netip.go:967
	_go_fuzz_dep_.CoverTab[4271]++
							if len(text) == 0 {
//line /snap/go/10455/src/net/netip/netip.go:968
		_go_fuzz_dep_.CoverTab[527221]++
//line /snap/go/10455/src/net/netip/netip.go:968
		_go_fuzz_dep_.CoverTab[4273]++
								*ip = Addr{}
								return nil
//line /snap/go/10455/src/net/netip/netip.go:970
		// _ = "end of CoverTab[4273]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:971
		_go_fuzz_dep_.CoverTab[527222]++
//line /snap/go/10455/src/net/netip/netip.go:971
		_go_fuzz_dep_.CoverTab[4274]++
//line /snap/go/10455/src/net/netip/netip.go:971
		// _ = "end of CoverTab[4274]"
//line /snap/go/10455/src/net/netip/netip.go:971
	}
//line /snap/go/10455/src/net/netip/netip.go:971
	// _ = "end of CoverTab[4271]"
//line /snap/go/10455/src/net/netip/netip.go:971
	_go_fuzz_dep_.CoverTab[4272]++
							var err error
							*ip, err = ParseAddr(string(text))
							return err
//line /snap/go/10455/src/net/netip/netip.go:974
	// _ = "end of CoverTab[4272]"
}

func (ip Addr) marshalBinaryWithTrailingBytes(trailingBytes int) []byte {
//line /snap/go/10455/src/net/netip/netip.go:977
	_go_fuzz_dep_.CoverTab[4275]++
							var b []byte
							switch ip.z {
	case z0:
//line /snap/go/10455/src/net/netip/netip.go:980
		_go_fuzz_dep_.CoverTab[527223]++
//line /snap/go/10455/src/net/netip/netip.go:980
		_go_fuzz_dep_.CoverTab[4277]++
								b = make([]byte, trailingBytes)
//line /snap/go/10455/src/net/netip/netip.go:981
		// _ = "end of CoverTab[4277]"
	case z4:
//line /snap/go/10455/src/net/netip/netip.go:982
		_go_fuzz_dep_.CoverTab[527224]++
//line /snap/go/10455/src/net/netip/netip.go:982
		_go_fuzz_dep_.CoverTab[4278]++
								b = make([]byte, 4+trailingBytes)
								bePutUint32(b, uint32(ip.addr.lo))
//line /snap/go/10455/src/net/netip/netip.go:984
		// _ = "end of CoverTab[4278]"
	default:
//line /snap/go/10455/src/net/netip/netip.go:985
		_go_fuzz_dep_.CoverTab[527225]++
//line /snap/go/10455/src/net/netip/netip.go:985
		_go_fuzz_dep_.CoverTab[4279]++
								z := ip.Zone()
								b = make([]byte, 16+len(z)+trailingBytes)
								bePutUint64(b[:8], ip.addr.hi)
								bePutUint64(b[8:], ip.addr.lo)
								copy(b[16:], z)
//line /snap/go/10455/src/net/netip/netip.go:990
		// _ = "end of CoverTab[4279]"
	}
//line /snap/go/10455/src/net/netip/netip.go:991
	// _ = "end of CoverTab[4275]"
//line /snap/go/10455/src/net/netip/netip.go:991
	_go_fuzz_dep_.CoverTab[4276]++
							return b
//line /snap/go/10455/src/net/netip/netip.go:992
	// _ = "end of CoverTab[4276]"
}

// MarshalBinary implements the encoding.BinaryMarshaler interface.
//line /snap/go/10455/src/net/netip/netip.go:995
// It returns a zero-length slice for the zero Addr,
//line /snap/go/10455/src/net/netip/netip.go:995
// the 4-byte form for an IPv4 address,
//line /snap/go/10455/src/net/netip/netip.go:995
// and the 16-byte form with zone appended for an IPv6 address.
//line /snap/go/10455/src/net/netip/netip.go:999
func (ip Addr) MarshalBinary() ([]byte, error) {
//line /snap/go/10455/src/net/netip/netip.go:999
	_go_fuzz_dep_.CoverTab[4280]++
							return ip.marshalBinaryWithTrailingBytes(0), nil
//line /snap/go/10455/src/net/netip/netip.go:1000
	// _ = "end of CoverTab[4280]"
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface.
//line /snap/go/10455/src/net/netip/netip.go:1003
// It expects data in the form generated by MarshalBinary.
//line /snap/go/10455/src/net/netip/netip.go:1005
func (ip *Addr) UnmarshalBinary(b []byte) error {
//line /snap/go/10455/src/net/netip/netip.go:1005
	_go_fuzz_dep_.CoverTab[4281]++
							n := len(b)
							switch {
	case n == 0:
//line /snap/go/10455/src/net/netip/netip.go:1008
		_go_fuzz_dep_.CoverTab[527226]++
//line /snap/go/10455/src/net/netip/netip.go:1008
		_go_fuzz_dep_.CoverTab[4283]++
								*ip = Addr{}
								return nil
//line /snap/go/10455/src/net/netip/netip.go:1010
		// _ = "end of CoverTab[4283]"
	case n == 4:
//line /snap/go/10455/src/net/netip/netip.go:1011
		_go_fuzz_dep_.CoverTab[527227]++
//line /snap/go/10455/src/net/netip/netip.go:1011
		_go_fuzz_dep_.CoverTab[4284]++
								*ip = AddrFrom4([4]byte(b))
								return nil
//line /snap/go/10455/src/net/netip/netip.go:1013
		// _ = "end of CoverTab[4284]"
	case n == 16:
//line /snap/go/10455/src/net/netip/netip.go:1014
		_go_fuzz_dep_.CoverTab[527228]++
//line /snap/go/10455/src/net/netip/netip.go:1014
		_go_fuzz_dep_.CoverTab[4285]++
								*ip = AddrFrom16([16]byte(b))
								return nil
//line /snap/go/10455/src/net/netip/netip.go:1016
		// _ = "end of CoverTab[4285]"
	case n > 16:
//line /snap/go/10455/src/net/netip/netip.go:1017
		_go_fuzz_dep_.CoverTab[527229]++
//line /snap/go/10455/src/net/netip/netip.go:1017
		_go_fuzz_dep_.CoverTab[4286]++
								*ip = AddrFrom16([16]byte(b[:16])).WithZone(string(b[16:]))
								return nil
//line /snap/go/10455/src/net/netip/netip.go:1019
		// _ = "end of CoverTab[4286]"
//line /snap/go/10455/src/net/netip/netip.go:1019
	default:
//line /snap/go/10455/src/net/netip/netip.go:1019
		_go_fuzz_dep_.CoverTab[527230]++
//line /snap/go/10455/src/net/netip/netip.go:1019
		_go_fuzz_dep_.CoverTab[4287]++
//line /snap/go/10455/src/net/netip/netip.go:1019
		// _ = "end of CoverTab[4287]"
	}
//line /snap/go/10455/src/net/netip/netip.go:1020
	// _ = "end of CoverTab[4281]"
//line /snap/go/10455/src/net/netip/netip.go:1020
	_go_fuzz_dep_.CoverTab[4282]++
							return errors.New("unexpected slice size")
//line /snap/go/10455/src/net/netip/netip.go:1021
	// _ = "end of CoverTab[4282]"
}

// AddrPort is an IP and a port number.
type AddrPort struct {
	ip	Addr
	port	uint16
}

// AddrPortFrom returns an AddrPort with the provided IP and port.
//line /snap/go/10455/src/net/netip/netip.go:1030
// It does not allocate.
//line /snap/go/10455/src/net/netip/netip.go:1032
func AddrPortFrom(ip Addr, port uint16) AddrPort {
//line /snap/go/10455/src/net/netip/netip.go:1032
	_go_fuzz_dep_.CoverTab[4288]++
//line /snap/go/10455/src/net/netip/netip.go:1032
	return AddrPort{ip: ip, port: port}
//line /snap/go/10455/src/net/netip/netip.go:1032
	// _ = "end of CoverTab[4288]"
//line /snap/go/10455/src/net/netip/netip.go:1032
}

// Addr returns p's IP address.
func (p AddrPort) Addr() Addr {
//line /snap/go/10455/src/net/netip/netip.go:1035
	_go_fuzz_dep_.CoverTab[4289]++
//line /snap/go/10455/src/net/netip/netip.go:1035
	return p.ip
//line /snap/go/10455/src/net/netip/netip.go:1035
	// _ = "end of CoverTab[4289]"
//line /snap/go/10455/src/net/netip/netip.go:1035
}

// Port returns p's port.
func (p AddrPort) Port() uint16 {
//line /snap/go/10455/src/net/netip/netip.go:1038
	_go_fuzz_dep_.CoverTab[4290]++
//line /snap/go/10455/src/net/netip/netip.go:1038
	return p.port
//line /snap/go/10455/src/net/netip/netip.go:1038
	// _ = "end of CoverTab[4290]"
//line /snap/go/10455/src/net/netip/netip.go:1038
}

// splitAddrPort splits s into an IP address string and a port
//line /snap/go/10455/src/net/netip/netip.go:1040
// string. It splits strings shaped like "foo:bar" or "[foo]:bar",
//line /snap/go/10455/src/net/netip/netip.go:1040
// without further validating the substrings. v6 indicates whether the
//line /snap/go/10455/src/net/netip/netip.go:1040
// ip string should parse as an IPv6 address or an IPv4 address, in
//line /snap/go/10455/src/net/netip/netip.go:1040
// order for s to be a valid ip:port string.
//line /snap/go/10455/src/net/netip/netip.go:1045
func splitAddrPort(s string) (ip, port string, v6 bool, err error) {
//line /snap/go/10455/src/net/netip/netip.go:1045
	_go_fuzz_dep_.CoverTab[4291]++
							i := stringsLastIndexByte(s, ':')
							if i == -1 {
//line /snap/go/10455/src/net/netip/netip.go:1047
		_go_fuzz_dep_.CoverTab[527231]++
//line /snap/go/10455/src/net/netip/netip.go:1047
		_go_fuzz_dep_.CoverTab[4296]++
								return "", "", false, errors.New("not an ip:port")
//line /snap/go/10455/src/net/netip/netip.go:1048
		// _ = "end of CoverTab[4296]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:1049
		_go_fuzz_dep_.CoverTab[527232]++
//line /snap/go/10455/src/net/netip/netip.go:1049
		_go_fuzz_dep_.CoverTab[4297]++
//line /snap/go/10455/src/net/netip/netip.go:1049
		// _ = "end of CoverTab[4297]"
//line /snap/go/10455/src/net/netip/netip.go:1049
	}
//line /snap/go/10455/src/net/netip/netip.go:1049
	// _ = "end of CoverTab[4291]"
//line /snap/go/10455/src/net/netip/netip.go:1049
	_go_fuzz_dep_.CoverTab[4292]++

							ip, port = s[:i], s[i+1:]
							if len(ip) == 0 {
//line /snap/go/10455/src/net/netip/netip.go:1052
		_go_fuzz_dep_.CoverTab[527233]++
//line /snap/go/10455/src/net/netip/netip.go:1052
		_go_fuzz_dep_.CoverTab[4298]++
								return "", "", false, errors.New("no IP")
//line /snap/go/10455/src/net/netip/netip.go:1053
		// _ = "end of CoverTab[4298]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:1054
		_go_fuzz_dep_.CoverTab[527234]++
//line /snap/go/10455/src/net/netip/netip.go:1054
		_go_fuzz_dep_.CoverTab[4299]++
//line /snap/go/10455/src/net/netip/netip.go:1054
		// _ = "end of CoverTab[4299]"
//line /snap/go/10455/src/net/netip/netip.go:1054
	}
//line /snap/go/10455/src/net/netip/netip.go:1054
	// _ = "end of CoverTab[4292]"
//line /snap/go/10455/src/net/netip/netip.go:1054
	_go_fuzz_dep_.CoverTab[4293]++
							if len(port) == 0 {
//line /snap/go/10455/src/net/netip/netip.go:1055
		_go_fuzz_dep_.CoverTab[527235]++
//line /snap/go/10455/src/net/netip/netip.go:1055
		_go_fuzz_dep_.CoverTab[4300]++
								return "", "", false, errors.New("no port")
//line /snap/go/10455/src/net/netip/netip.go:1056
		// _ = "end of CoverTab[4300]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:1057
		_go_fuzz_dep_.CoverTab[527236]++
//line /snap/go/10455/src/net/netip/netip.go:1057
		_go_fuzz_dep_.CoverTab[4301]++
//line /snap/go/10455/src/net/netip/netip.go:1057
		// _ = "end of CoverTab[4301]"
//line /snap/go/10455/src/net/netip/netip.go:1057
	}
//line /snap/go/10455/src/net/netip/netip.go:1057
	// _ = "end of CoverTab[4293]"
//line /snap/go/10455/src/net/netip/netip.go:1057
	_go_fuzz_dep_.CoverTab[4294]++
							if ip[0] == '[' {
//line /snap/go/10455/src/net/netip/netip.go:1058
		_go_fuzz_dep_.CoverTab[527237]++
//line /snap/go/10455/src/net/netip/netip.go:1058
		_go_fuzz_dep_.CoverTab[4302]++
								if len(ip) < 2 || func() bool {
//line /snap/go/10455/src/net/netip/netip.go:1059
			_go_fuzz_dep_.CoverTab[4304]++
//line /snap/go/10455/src/net/netip/netip.go:1059
			return ip[len(ip)-1] != ']'
//line /snap/go/10455/src/net/netip/netip.go:1059
			// _ = "end of CoverTab[4304]"
//line /snap/go/10455/src/net/netip/netip.go:1059
		}() {
//line /snap/go/10455/src/net/netip/netip.go:1059
			_go_fuzz_dep_.CoverTab[527239]++
//line /snap/go/10455/src/net/netip/netip.go:1059
			_go_fuzz_dep_.CoverTab[4305]++
									return "", "", false, errors.New("missing ]")
//line /snap/go/10455/src/net/netip/netip.go:1060
			// _ = "end of CoverTab[4305]"
		} else {
//line /snap/go/10455/src/net/netip/netip.go:1061
			_go_fuzz_dep_.CoverTab[527240]++
//line /snap/go/10455/src/net/netip/netip.go:1061
			_go_fuzz_dep_.CoverTab[4306]++
//line /snap/go/10455/src/net/netip/netip.go:1061
			// _ = "end of CoverTab[4306]"
//line /snap/go/10455/src/net/netip/netip.go:1061
		}
//line /snap/go/10455/src/net/netip/netip.go:1061
		// _ = "end of CoverTab[4302]"
//line /snap/go/10455/src/net/netip/netip.go:1061
		_go_fuzz_dep_.CoverTab[4303]++
								ip = ip[1 : len(ip)-1]
								v6 = true
//line /snap/go/10455/src/net/netip/netip.go:1063
		// _ = "end of CoverTab[4303]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:1064
		_go_fuzz_dep_.CoverTab[527238]++
//line /snap/go/10455/src/net/netip/netip.go:1064
		_go_fuzz_dep_.CoverTab[4307]++
//line /snap/go/10455/src/net/netip/netip.go:1064
		// _ = "end of CoverTab[4307]"
//line /snap/go/10455/src/net/netip/netip.go:1064
	}
//line /snap/go/10455/src/net/netip/netip.go:1064
	// _ = "end of CoverTab[4294]"
//line /snap/go/10455/src/net/netip/netip.go:1064
	_go_fuzz_dep_.CoverTab[4295]++

							return ip, port, v6, nil
//line /snap/go/10455/src/net/netip/netip.go:1066
	// _ = "end of CoverTab[4295]"
}

// ParseAddrPort parses s as an AddrPort.
//line /snap/go/10455/src/net/netip/netip.go:1069
//
//line /snap/go/10455/src/net/netip/netip.go:1069
// It doesn't do any name resolution: both the address and the port
//line /snap/go/10455/src/net/netip/netip.go:1069
// must be numeric.
//line /snap/go/10455/src/net/netip/netip.go:1073
func ParseAddrPort(s string) (AddrPort, error) {
//line /snap/go/10455/src/net/netip/netip.go:1073
	_go_fuzz_dep_.CoverTab[4308]++
							var ipp AddrPort
							ip, port, v6, err := splitAddrPort(s)
							if err != nil {
//line /snap/go/10455/src/net/netip/netip.go:1076
		_go_fuzz_dep_.CoverTab[527241]++
//line /snap/go/10455/src/net/netip/netip.go:1076
		_go_fuzz_dep_.CoverTab[4313]++
								return ipp, err
//line /snap/go/10455/src/net/netip/netip.go:1077
		// _ = "end of CoverTab[4313]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:1078
		_go_fuzz_dep_.CoverTab[527242]++
//line /snap/go/10455/src/net/netip/netip.go:1078
		_go_fuzz_dep_.CoverTab[4314]++
//line /snap/go/10455/src/net/netip/netip.go:1078
		// _ = "end of CoverTab[4314]"
//line /snap/go/10455/src/net/netip/netip.go:1078
	}
//line /snap/go/10455/src/net/netip/netip.go:1078
	// _ = "end of CoverTab[4308]"
//line /snap/go/10455/src/net/netip/netip.go:1078
	_go_fuzz_dep_.CoverTab[4309]++
							port16, err := strconv.ParseUint(port, 10, 16)
							if err != nil {
//line /snap/go/10455/src/net/netip/netip.go:1080
		_go_fuzz_dep_.CoverTab[527243]++
//line /snap/go/10455/src/net/netip/netip.go:1080
		_go_fuzz_dep_.CoverTab[4315]++
								return ipp, errors.New("invalid port " + strconv.Quote(port) + " parsing " + strconv.Quote(s))
//line /snap/go/10455/src/net/netip/netip.go:1081
		// _ = "end of CoverTab[4315]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:1082
		_go_fuzz_dep_.CoverTab[527244]++
//line /snap/go/10455/src/net/netip/netip.go:1082
		_go_fuzz_dep_.CoverTab[4316]++
//line /snap/go/10455/src/net/netip/netip.go:1082
		// _ = "end of CoverTab[4316]"
//line /snap/go/10455/src/net/netip/netip.go:1082
	}
//line /snap/go/10455/src/net/netip/netip.go:1082
	// _ = "end of CoverTab[4309]"
//line /snap/go/10455/src/net/netip/netip.go:1082
	_go_fuzz_dep_.CoverTab[4310]++
							ipp.port = uint16(port16)
							ipp.ip, err = ParseAddr(ip)
							if err != nil {
//line /snap/go/10455/src/net/netip/netip.go:1085
		_go_fuzz_dep_.CoverTab[527245]++
//line /snap/go/10455/src/net/netip/netip.go:1085
		_go_fuzz_dep_.CoverTab[4317]++
								return AddrPort{}, err
//line /snap/go/10455/src/net/netip/netip.go:1086
		// _ = "end of CoverTab[4317]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:1087
		_go_fuzz_dep_.CoverTab[527246]++
//line /snap/go/10455/src/net/netip/netip.go:1087
		_go_fuzz_dep_.CoverTab[4318]++
//line /snap/go/10455/src/net/netip/netip.go:1087
		// _ = "end of CoverTab[4318]"
//line /snap/go/10455/src/net/netip/netip.go:1087
	}
//line /snap/go/10455/src/net/netip/netip.go:1087
	// _ = "end of CoverTab[4310]"
//line /snap/go/10455/src/net/netip/netip.go:1087
	_go_fuzz_dep_.CoverTab[4311]++
							if v6 && func() bool {
//line /snap/go/10455/src/net/netip/netip.go:1088
		_go_fuzz_dep_.CoverTab[4319]++
//line /snap/go/10455/src/net/netip/netip.go:1088
		return ipp.ip.Is4()
//line /snap/go/10455/src/net/netip/netip.go:1088
		// _ = "end of CoverTab[4319]"
//line /snap/go/10455/src/net/netip/netip.go:1088
	}() {
//line /snap/go/10455/src/net/netip/netip.go:1088
		_go_fuzz_dep_.CoverTab[527247]++
//line /snap/go/10455/src/net/netip/netip.go:1088
		_go_fuzz_dep_.CoverTab[4320]++
								return AddrPort{}, errors.New("invalid ip:port " + strconv.Quote(s) + ", square brackets can only be used with IPv6 addresses")
//line /snap/go/10455/src/net/netip/netip.go:1089
		// _ = "end of CoverTab[4320]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:1090
		_go_fuzz_dep_.CoverTab[527248]++
//line /snap/go/10455/src/net/netip/netip.go:1090
		_go_fuzz_dep_.CoverTab[4321]++
//line /snap/go/10455/src/net/netip/netip.go:1090
		if !v6 && func() bool {
//line /snap/go/10455/src/net/netip/netip.go:1090
			_go_fuzz_dep_.CoverTab[4322]++
//line /snap/go/10455/src/net/netip/netip.go:1090
			return ipp.ip.Is6()
//line /snap/go/10455/src/net/netip/netip.go:1090
			// _ = "end of CoverTab[4322]"
//line /snap/go/10455/src/net/netip/netip.go:1090
		}() {
//line /snap/go/10455/src/net/netip/netip.go:1090
			_go_fuzz_dep_.CoverTab[527249]++
//line /snap/go/10455/src/net/netip/netip.go:1090
			_go_fuzz_dep_.CoverTab[4323]++
									return AddrPort{}, errors.New("invalid ip:port " + strconv.Quote(s) + ", IPv6 addresses must be surrounded by square brackets")
//line /snap/go/10455/src/net/netip/netip.go:1091
			// _ = "end of CoverTab[4323]"
		} else {
//line /snap/go/10455/src/net/netip/netip.go:1092
			_go_fuzz_dep_.CoverTab[527250]++
//line /snap/go/10455/src/net/netip/netip.go:1092
			_go_fuzz_dep_.CoverTab[4324]++
//line /snap/go/10455/src/net/netip/netip.go:1092
			// _ = "end of CoverTab[4324]"
//line /snap/go/10455/src/net/netip/netip.go:1092
		}
//line /snap/go/10455/src/net/netip/netip.go:1092
		// _ = "end of CoverTab[4321]"
//line /snap/go/10455/src/net/netip/netip.go:1092
	}
//line /snap/go/10455/src/net/netip/netip.go:1092
	// _ = "end of CoverTab[4311]"
//line /snap/go/10455/src/net/netip/netip.go:1092
	_go_fuzz_dep_.CoverTab[4312]++
							return ipp, nil
//line /snap/go/10455/src/net/netip/netip.go:1093
	// _ = "end of CoverTab[4312]"
}

// MustParseAddrPort calls ParseAddrPort(s) and panics on error.
//line /snap/go/10455/src/net/netip/netip.go:1096
// It is intended for use in tests with hard-coded strings.
//line /snap/go/10455/src/net/netip/netip.go:1098
func MustParseAddrPort(s string) AddrPort {
//line /snap/go/10455/src/net/netip/netip.go:1098
	_go_fuzz_dep_.CoverTab[4325]++
							ip, err := ParseAddrPort(s)
							if err != nil {
//line /snap/go/10455/src/net/netip/netip.go:1100
		_go_fuzz_dep_.CoverTab[527251]++
//line /snap/go/10455/src/net/netip/netip.go:1100
		_go_fuzz_dep_.CoverTab[4327]++
								panic(err)
//line /snap/go/10455/src/net/netip/netip.go:1101
		// _ = "end of CoverTab[4327]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:1102
		_go_fuzz_dep_.CoverTab[527252]++
//line /snap/go/10455/src/net/netip/netip.go:1102
		_go_fuzz_dep_.CoverTab[4328]++
//line /snap/go/10455/src/net/netip/netip.go:1102
		// _ = "end of CoverTab[4328]"
//line /snap/go/10455/src/net/netip/netip.go:1102
	}
//line /snap/go/10455/src/net/netip/netip.go:1102
	// _ = "end of CoverTab[4325]"
//line /snap/go/10455/src/net/netip/netip.go:1102
	_go_fuzz_dep_.CoverTab[4326]++
							return ip
//line /snap/go/10455/src/net/netip/netip.go:1103
	// _ = "end of CoverTab[4326]"
}

// IsValid reports whether p.Addr() is valid.
//line /snap/go/10455/src/net/netip/netip.go:1106
// All ports are valid, including zero.
//line /snap/go/10455/src/net/netip/netip.go:1108
func (p AddrPort) IsValid() bool {
//line /snap/go/10455/src/net/netip/netip.go:1108
	_go_fuzz_dep_.CoverTab[4329]++
//line /snap/go/10455/src/net/netip/netip.go:1108
	return p.ip.IsValid()
//line /snap/go/10455/src/net/netip/netip.go:1108
	// _ = "end of CoverTab[4329]"
//line /snap/go/10455/src/net/netip/netip.go:1108
}

func (p AddrPort) String() string {
//line /snap/go/10455/src/net/netip/netip.go:1110
	_go_fuzz_dep_.CoverTab[4330]++
							switch p.ip.z {
	case z0:
//line /snap/go/10455/src/net/netip/netip.go:1112
		_go_fuzz_dep_.CoverTab[527253]++
//line /snap/go/10455/src/net/netip/netip.go:1112
		_go_fuzz_dep_.CoverTab[4331]++
								return "invalid AddrPort"
//line /snap/go/10455/src/net/netip/netip.go:1113
		// _ = "end of CoverTab[4331]"
	case z4:
//line /snap/go/10455/src/net/netip/netip.go:1114
		_go_fuzz_dep_.CoverTab[527254]++
//line /snap/go/10455/src/net/netip/netip.go:1114
		_go_fuzz_dep_.CoverTab[4332]++
								a := p.ip.As4()
								buf := make([]byte, 0, 21)
								for i := range a {
//line /snap/go/10455/src/net/netip/netip.go:1117
			_go_fuzz_dep_.CoverTab[4335]++
									buf = strconv.AppendUint(buf, uint64(a[i]), 10)
									buf = append(buf, "...:"[i])
//line /snap/go/10455/src/net/netip/netip.go:1119
			// _ = "end of CoverTab[4335]"
		}
//line /snap/go/10455/src/net/netip/netip.go:1120
		// _ = "end of CoverTab[4332]"
//line /snap/go/10455/src/net/netip/netip.go:1120
		_go_fuzz_dep_.CoverTab[4333]++
								buf = strconv.AppendUint(buf, uint64(p.port), 10)
								return string(buf)
//line /snap/go/10455/src/net/netip/netip.go:1122
		// _ = "end of CoverTab[4333]"
	default:
//line /snap/go/10455/src/net/netip/netip.go:1123
		_go_fuzz_dep_.CoverTab[527255]++
//line /snap/go/10455/src/net/netip/netip.go:1123
		_go_fuzz_dep_.CoverTab[4334]++

								return joinHostPort(p.ip.String(), itoa.Itoa(int(p.port)))
//line /snap/go/10455/src/net/netip/netip.go:1125
		// _ = "end of CoverTab[4334]"
	}
//line /snap/go/10455/src/net/netip/netip.go:1126
	// _ = "end of CoverTab[4330]"
}

func joinHostPort(host, port string) string {
//line /snap/go/10455/src/net/netip/netip.go:1129
	_go_fuzz_dep_.CoverTab[4336]++

//line /snap/go/10455/src/net/netip/netip.go:1132
	if bytealg.IndexByteString(host, ':') >= 0 {
//line /snap/go/10455/src/net/netip/netip.go:1132
		_go_fuzz_dep_.CoverTab[527256]++
//line /snap/go/10455/src/net/netip/netip.go:1132
		_go_fuzz_dep_.CoverTab[4338]++
								return "[" + host + "]:" + port
//line /snap/go/10455/src/net/netip/netip.go:1133
		// _ = "end of CoverTab[4338]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:1134
		_go_fuzz_dep_.CoverTab[527257]++
//line /snap/go/10455/src/net/netip/netip.go:1134
		_go_fuzz_dep_.CoverTab[4339]++
//line /snap/go/10455/src/net/netip/netip.go:1134
		// _ = "end of CoverTab[4339]"
//line /snap/go/10455/src/net/netip/netip.go:1134
	}
//line /snap/go/10455/src/net/netip/netip.go:1134
	// _ = "end of CoverTab[4336]"
//line /snap/go/10455/src/net/netip/netip.go:1134
	_go_fuzz_dep_.CoverTab[4337]++
							return host + ":" + port
//line /snap/go/10455/src/net/netip/netip.go:1135
	// _ = "end of CoverTab[4337]"
}

// AppendTo appends a text encoding of p,
//line /snap/go/10455/src/net/netip/netip.go:1138
// as generated by MarshalText,
//line /snap/go/10455/src/net/netip/netip.go:1138
// to b and returns the extended buffer.
//line /snap/go/10455/src/net/netip/netip.go:1141
func (p AddrPort) AppendTo(b []byte) []byte {
//line /snap/go/10455/src/net/netip/netip.go:1141
	_go_fuzz_dep_.CoverTab[4340]++
							switch p.ip.z {
	case z0:
//line /snap/go/10455/src/net/netip/netip.go:1143
		_go_fuzz_dep_.CoverTab[527258]++
//line /snap/go/10455/src/net/netip/netip.go:1143
		_go_fuzz_dep_.CoverTab[4342]++
								return b
//line /snap/go/10455/src/net/netip/netip.go:1144
		// _ = "end of CoverTab[4342]"
	case z4:
//line /snap/go/10455/src/net/netip/netip.go:1145
		_go_fuzz_dep_.CoverTab[527259]++
//line /snap/go/10455/src/net/netip/netip.go:1145
		_go_fuzz_dep_.CoverTab[4343]++
								b = p.ip.appendTo4(b)
//line /snap/go/10455/src/net/netip/netip.go:1146
		// _ = "end of CoverTab[4343]"
	default:
//line /snap/go/10455/src/net/netip/netip.go:1147
		_go_fuzz_dep_.CoverTab[527260]++
//line /snap/go/10455/src/net/netip/netip.go:1147
		_go_fuzz_dep_.CoverTab[4344]++
								if p.ip.Is4In6() {
//line /snap/go/10455/src/net/netip/netip.go:1148
			_go_fuzz_dep_.CoverTab[527261]++
//line /snap/go/10455/src/net/netip/netip.go:1148
			_go_fuzz_dep_.CoverTab[4346]++
									b = append(b, "[::ffff:"...)
									b = p.ip.Unmap().appendTo4(b)
									if z := p.ip.Zone(); z != "" {
//line /snap/go/10455/src/net/netip/netip.go:1151
				_go_fuzz_dep_.CoverTab[527263]++
//line /snap/go/10455/src/net/netip/netip.go:1151
				_go_fuzz_dep_.CoverTab[4347]++
										b = append(b, '%')
										b = append(b, z...)
//line /snap/go/10455/src/net/netip/netip.go:1153
				// _ = "end of CoverTab[4347]"
			} else {
//line /snap/go/10455/src/net/netip/netip.go:1154
				_go_fuzz_dep_.CoverTab[527264]++
//line /snap/go/10455/src/net/netip/netip.go:1154
				_go_fuzz_dep_.CoverTab[4348]++
//line /snap/go/10455/src/net/netip/netip.go:1154
				// _ = "end of CoverTab[4348]"
//line /snap/go/10455/src/net/netip/netip.go:1154
			}
//line /snap/go/10455/src/net/netip/netip.go:1154
			// _ = "end of CoverTab[4346]"
		} else {
//line /snap/go/10455/src/net/netip/netip.go:1155
			_go_fuzz_dep_.CoverTab[527262]++
//line /snap/go/10455/src/net/netip/netip.go:1155
			_go_fuzz_dep_.CoverTab[4349]++
									b = append(b, '[')
									b = p.ip.appendTo6(b)
//line /snap/go/10455/src/net/netip/netip.go:1157
			// _ = "end of CoverTab[4349]"
		}
//line /snap/go/10455/src/net/netip/netip.go:1158
		// _ = "end of CoverTab[4344]"
//line /snap/go/10455/src/net/netip/netip.go:1158
		_go_fuzz_dep_.CoverTab[4345]++
								b = append(b, ']')
//line /snap/go/10455/src/net/netip/netip.go:1159
		// _ = "end of CoverTab[4345]"
	}
//line /snap/go/10455/src/net/netip/netip.go:1160
	// _ = "end of CoverTab[4340]"
//line /snap/go/10455/src/net/netip/netip.go:1160
	_go_fuzz_dep_.CoverTab[4341]++
							b = append(b, ':')
							b = strconv.AppendUint(b, uint64(p.port), 10)
							return b
//line /snap/go/10455/src/net/netip/netip.go:1163
	// _ = "end of CoverTab[4341]"
}

// MarshalText implements the encoding.TextMarshaler interface. The
//line /snap/go/10455/src/net/netip/netip.go:1166
// encoding is the same as returned by String, with one exception: if
//line /snap/go/10455/src/net/netip/netip.go:1166
// p.Addr() is the zero Addr, the encoding is the empty string.
//line /snap/go/10455/src/net/netip/netip.go:1169
func (p AddrPort) MarshalText() ([]byte, error) {
//line /snap/go/10455/src/net/netip/netip.go:1169
	_go_fuzz_dep_.CoverTab[4350]++
							var max int
							switch p.ip.z {
	case z0:
//line /snap/go/10455/src/net/netip/netip.go:1172
		_go_fuzz_dep_.CoverTab[527265]++
//line /snap/go/10455/src/net/netip/netip.go:1172
		_go_fuzz_dep_.CoverTab[4352]++
//line /snap/go/10455/src/net/netip/netip.go:1172
		// _ = "end of CoverTab[4352]"
	case z4:
//line /snap/go/10455/src/net/netip/netip.go:1173
		_go_fuzz_dep_.CoverTab[527266]++
//line /snap/go/10455/src/net/netip/netip.go:1173
		_go_fuzz_dep_.CoverTab[4353]++
								max = len("255.255.255.255:65535")
//line /snap/go/10455/src/net/netip/netip.go:1174
		// _ = "end of CoverTab[4353]"
	default:
//line /snap/go/10455/src/net/netip/netip.go:1175
		_go_fuzz_dep_.CoverTab[527267]++
//line /snap/go/10455/src/net/netip/netip.go:1175
		_go_fuzz_dep_.CoverTab[4354]++
								max = len("[ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff%enp5s0]:65535")
//line /snap/go/10455/src/net/netip/netip.go:1176
		// _ = "end of CoverTab[4354]"
	}
//line /snap/go/10455/src/net/netip/netip.go:1177
	// _ = "end of CoverTab[4350]"
//line /snap/go/10455/src/net/netip/netip.go:1177
	_go_fuzz_dep_.CoverTab[4351]++
							b := make([]byte, 0, max)
							b = p.AppendTo(b)
							return b, nil
//line /snap/go/10455/src/net/netip/netip.go:1180
	// _ = "end of CoverTab[4351]"
}

// UnmarshalText implements the encoding.TextUnmarshaler
//line /snap/go/10455/src/net/netip/netip.go:1183
// interface. The AddrPort is expected in a form
//line /snap/go/10455/src/net/netip/netip.go:1183
// generated by MarshalText or accepted by ParseAddrPort.
//line /snap/go/10455/src/net/netip/netip.go:1186
func (p *AddrPort) UnmarshalText(text []byte) error {
//line /snap/go/10455/src/net/netip/netip.go:1186
	_go_fuzz_dep_.CoverTab[4355]++
							if len(text) == 0 {
//line /snap/go/10455/src/net/netip/netip.go:1187
		_go_fuzz_dep_.CoverTab[527268]++
//line /snap/go/10455/src/net/netip/netip.go:1187
		_go_fuzz_dep_.CoverTab[4357]++
								*p = AddrPort{}
								return nil
//line /snap/go/10455/src/net/netip/netip.go:1189
		// _ = "end of CoverTab[4357]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:1190
		_go_fuzz_dep_.CoverTab[527269]++
//line /snap/go/10455/src/net/netip/netip.go:1190
		_go_fuzz_dep_.CoverTab[4358]++
//line /snap/go/10455/src/net/netip/netip.go:1190
		// _ = "end of CoverTab[4358]"
//line /snap/go/10455/src/net/netip/netip.go:1190
	}
//line /snap/go/10455/src/net/netip/netip.go:1190
	// _ = "end of CoverTab[4355]"
//line /snap/go/10455/src/net/netip/netip.go:1190
	_go_fuzz_dep_.CoverTab[4356]++
							var err error
							*p, err = ParseAddrPort(string(text))
							return err
//line /snap/go/10455/src/net/netip/netip.go:1193
	// _ = "end of CoverTab[4356]"
}

// MarshalBinary implements the encoding.BinaryMarshaler interface.
//line /snap/go/10455/src/net/netip/netip.go:1196
// It returns Addr.MarshalBinary with an additional two bytes appended
//line /snap/go/10455/src/net/netip/netip.go:1196
// containing the port in little-endian.
//line /snap/go/10455/src/net/netip/netip.go:1199
func (p AddrPort) MarshalBinary() ([]byte, error) {
//line /snap/go/10455/src/net/netip/netip.go:1199
	_go_fuzz_dep_.CoverTab[4359]++
							b := p.Addr().marshalBinaryWithTrailingBytes(2)
							lePutUint16(b[len(b)-2:], p.Port())
							return b, nil
//line /snap/go/10455/src/net/netip/netip.go:1202
	// _ = "end of CoverTab[4359]"
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface.
//line /snap/go/10455/src/net/netip/netip.go:1205
// It expects data in the form generated by MarshalBinary.
//line /snap/go/10455/src/net/netip/netip.go:1207
func (p *AddrPort) UnmarshalBinary(b []byte) error {
//line /snap/go/10455/src/net/netip/netip.go:1207
	_go_fuzz_dep_.CoverTab[4360]++
							if len(b) < 2 {
//line /snap/go/10455/src/net/netip/netip.go:1208
		_go_fuzz_dep_.CoverTab[527270]++
//line /snap/go/10455/src/net/netip/netip.go:1208
		_go_fuzz_dep_.CoverTab[4363]++
								return errors.New("unexpected slice size")
//line /snap/go/10455/src/net/netip/netip.go:1209
		// _ = "end of CoverTab[4363]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:1210
		_go_fuzz_dep_.CoverTab[527271]++
//line /snap/go/10455/src/net/netip/netip.go:1210
		_go_fuzz_dep_.CoverTab[4364]++
//line /snap/go/10455/src/net/netip/netip.go:1210
		// _ = "end of CoverTab[4364]"
//line /snap/go/10455/src/net/netip/netip.go:1210
	}
//line /snap/go/10455/src/net/netip/netip.go:1210
	// _ = "end of CoverTab[4360]"
//line /snap/go/10455/src/net/netip/netip.go:1210
	_go_fuzz_dep_.CoverTab[4361]++
							var addr Addr
							err := addr.UnmarshalBinary(b[:len(b)-2])
							if err != nil {
//line /snap/go/10455/src/net/netip/netip.go:1213
		_go_fuzz_dep_.CoverTab[527272]++
//line /snap/go/10455/src/net/netip/netip.go:1213
		_go_fuzz_dep_.CoverTab[4365]++
								return err
//line /snap/go/10455/src/net/netip/netip.go:1214
		// _ = "end of CoverTab[4365]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:1215
		_go_fuzz_dep_.CoverTab[527273]++
//line /snap/go/10455/src/net/netip/netip.go:1215
		_go_fuzz_dep_.CoverTab[4366]++
//line /snap/go/10455/src/net/netip/netip.go:1215
		// _ = "end of CoverTab[4366]"
//line /snap/go/10455/src/net/netip/netip.go:1215
	}
//line /snap/go/10455/src/net/netip/netip.go:1215
	// _ = "end of CoverTab[4361]"
//line /snap/go/10455/src/net/netip/netip.go:1215
	_go_fuzz_dep_.CoverTab[4362]++
							*p = AddrPortFrom(addr, leUint16(b[len(b)-2:]))
							return nil
//line /snap/go/10455/src/net/netip/netip.go:1217
	// _ = "end of CoverTab[4362]"
}

// Prefix is an IP address prefix (CIDR) representing an IP network.
//line /snap/go/10455/src/net/netip/netip.go:1220
//
//line /snap/go/10455/src/net/netip/netip.go:1220
// The first Bits() of Addr() are specified. The remaining bits match any address.
//line /snap/go/10455/src/net/netip/netip.go:1220
// The range of Bits() is [0,32] for IPv4 or [0,128] for IPv6.
//line /snap/go/10455/src/net/netip/netip.go:1224
type Prefix struct {
	ip	Addr

	// bitsPlusOne stores the prefix bit length plus one.
	// A Prefix is valid if and only if bitsPlusOne is non-zero.
	bitsPlusOne	uint8
}

// PrefixFrom returns a Prefix with the provided IP address and bit
//line /snap/go/10455/src/net/netip/netip.go:1232
// prefix length.
//line /snap/go/10455/src/net/netip/netip.go:1232
//
//line /snap/go/10455/src/net/netip/netip.go:1232
// It does not allocate. Unlike Addr.Prefix, PrefixFrom does not mask
//line /snap/go/10455/src/net/netip/netip.go:1232
// off the host bits of ip.
//line /snap/go/10455/src/net/netip/netip.go:1232
//
//line /snap/go/10455/src/net/netip/netip.go:1232
// If bits is less than zero or greater than ip.BitLen, Prefix.Bits
//line /snap/go/10455/src/net/netip/netip.go:1232
// will return an invalid value -1.
//line /snap/go/10455/src/net/netip/netip.go:1240
func PrefixFrom(ip Addr, bits int) Prefix {
//line /snap/go/10455/src/net/netip/netip.go:1240
	_go_fuzz_dep_.CoverTab[4367]++
							var bitsPlusOne uint8
							if !ip.isZero() && func() bool {
//line /snap/go/10455/src/net/netip/netip.go:1242
		_go_fuzz_dep_.CoverTab[4369]++
//line /snap/go/10455/src/net/netip/netip.go:1242
		return bits >= 0
//line /snap/go/10455/src/net/netip/netip.go:1242
		// _ = "end of CoverTab[4369]"
//line /snap/go/10455/src/net/netip/netip.go:1242
	}() && func() bool {
//line /snap/go/10455/src/net/netip/netip.go:1242
		_go_fuzz_dep_.CoverTab[4370]++
//line /snap/go/10455/src/net/netip/netip.go:1242
		return bits <= ip.BitLen()
//line /snap/go/10455/src/net/netip/netip.go:1242
		// _ = "end of CoverTab[4370]"
//line /snap/go/10455/src/net/netip/netip.go:1242
	}() {
//line /snap/go/10455/src/net/netip/netip.go:1242
		_go_fuzz_dep_.CoverTab[527274]++
//line /snap/go/10455/src/net/netip/netip.go:1242
		_go_fuzz_dep_.CoverTab[4371]++
								bitsPlusOne = uint8(bits) + 1
//line /snap/go/10455/src/net/netip/netip.go:1243
		// _ = "end of CoverTab[4371]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:1244
		_go_fuzz_dep_.CoverTab[527275]++
//line /snap/go/10455/src/net/netip/netip.go:1244
		_go_fuzz_dep_.CoverTab[4372]++
//line /snap/go/10455/src/net/netip/netip.go:1244
		// _ = "end of CoverTab[4372]"
//line /snap/go/10455/src/net/netip/netip.go:1244
	}
//line /snap/go/10455/src/net/netip/netip.go:1244
	// _ = "end of CoverTab[4367]"
//line /snap/go/10455/src/net/netip/netip.go:1244
	_go_fuzz_dep_.CoverTab[4368]++
							return Prefix{
		ip:		ip.withoutZone(),
		bitsPlusOne:	bitsPlusOne,
	}
//line /snap/go/10455/src/net/netip/netip.go:1248
	// _ = "end of CoverTab[4368]"
}

// Addr returns p's IP address.
func (p Prefix) Addr() Addr {
//line /snap/go/10455/src/net/netip/netip.go:1252
	_go_fuzz_dep_.CoverTab[4373]++
//line /snap/go/10455/src/net/netip/netip.go:1252
	return p.ip
//line /snap/go/10455/src/net/netip/netip.go:1252
	// _ = "end of CoverTab[4373]"
//line /snap/go/10455/src/net/netip/netip.go:1252
}

// Bits returns p's prefix length.
//line /snap/go/10455/src/net/netip/netip.go:1254
//
//line /snap/go/10455/src/net/netip/netip.go:1254
// It reports -1 if invalid.
//line /snap/go/10455/src/net/netip/netip.go:1257
func (p Prefix) Bits() int {
//line /snap/go/10455/src/net/netip/netip.go:1257
	_go_fuzz_dep_.CoverTab[4374]++
//line /snap/go/10455/src/net/netip/netip.go:1257
	return int(p.bitsPlusOne) - 1
//line /snap/go/10455/src/net/netip/netip.go:1257
	// _ = "end of CoverTab[4374]"
//line /snap/go/10455/src/net/netip/netip.go:1257
}

// IsValid reports whether p.Bits() has a valid range for p.Addr().
//line /snap/go/10455/src/net/netip/netip.go:1259
// If p.Addr() is the zero Addr, IsValid returns false.
//line /snap/go/10455/src/net/netip/netip.go:1259
// Note that if p is the zero Prefix, then p.IsValid() == false.
//line /snap/go/10455/src/net/netip/netip.go:1262
func (p Prefix) IsValid() bool {
//line /snap/go/10455/src/net/netip/netip.go:1262
	_go_fuzz_dep_.CoverTab[4375]++
//line /snap/go/10455/src/net/netip/netip.go:1262
	return p.bitsPlusOne > 0
//line /snap/go/10455/src/net/netip/netip.go:1262
	// _ = "end of CoverTab[4375]"
//line /snap/go/10455/src/net/netip/netip.go:1262
}

func (p Prefix) isZero() bool {
//line /snap/go/10455/src/net/netip/netip.go:1264
	_go_fuzz_dep_.CoverTab[4376]++
//line /snap/go/10455/src/net/netip/netip.go:1264
	return p == Prefix{}
//line /snap/go/10455/src/net/netip/netip.go:1264
	// _ = "end of CoverTab[4376]"
//line /snap/go/10455/src/net/netip/netip.go:1264
}

// IsSingleIP reports whether p contains exactly one IP.
func (p Prefix) IsSingleIP() bool {
//line /snap/go/10455/src/net/netip/netip.go:1267
	_go_fuzz_dep_.CoverTab[4377]++
//line /snap/go/10455/src/net/netip/netip.go:1267
	return p.IsValid() && func() bool {
//line /snap/go/10455/src/net/netip/netip.go:1267
		_go_fuzz_dep_.CoverTab[4378]++
//line /snap/go/10455/src/net/netip/netip.go:1267
		return p.Bits() == p.ip.BitLen()
//line /snap/go/10455/src/net/netip/netip.go:1267
		// _ = "end of CoverTab[4378]"
//line /snap/go/10455/src/net/netip/netip.go:1267
	}()
//line /snap/go/10455/src/net/netip/netip.go:1267
	// _ = "end of CoverTab[4377]"
//line /snap/go/10455/src/net/netip/netip.go:1267
}

// ParsePrefix parses s as an IP address prefix.
//line /snap/go/10455/src/net/netip/netip.go:1269
// The string can be in the form "192.168.1.0/24" or "2001:db8::/32",
//line /snap/go/10455/src/net/netip/netip.go:1269
// the CIDR notation defined in RFC 4632 and RFC 4291.
//line /snap/go/10455/src/net/netip/netip.go:1269
// IPv6 zones are not permitted in prefixes, and an error will be returned if a
//line /snap/go/10455/src/net/netip/netip.go:1269
// zone is present.
//line /snap/go/10455/src/net/netip/netip.go:1269
//
//line /snap/go/10455/src/net/netip/netip.go:1269
// Note that masked address bits are not zeroed. Use Masked for that.
//line /snap/go/10455/src/net/netip/netip.go:1276
func ParsePrefix(s string) (Prefix, error) {
//line /snap/go/10455/src/net/netip/netip.go:1276
	_go_fuzz_dep_.CoverTab[4379]++
							i := stringsLastIndexByte(s, '/')
							if i < 0 {
//line /snap/go/10455/src/net/netip/netip.go:1278
		_go_fuzz_dep_.CoverTab[527276]++
//line /snap/go/10455/src/net/netip/netip.go:1278
		_go_fuzz_dep_.CoverTab[4386]++
								return Prefix{}, errors.New("netip.ParsePrefix(" + strconv.Quote(s) + "): no '/'")
//line /snap/go/10455/src/net/netip/netip.go:1279
		// _ = "end of CoverTab[4386]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:1280
		_go_fuzz_dep_.CoverTab[527277]++
//line /snap/go/10455/src/net/netip/netip.go:1280
		_go_fuzz_dep_.CoverTab[4387]++
//line /snap/go/10455/src/net/netip/netip.go:1280
		// _ = "end of CoverTab[4387]"
//line /snap/go/10455/src/net/netip/netip.go:1280
	}
//line /snap/go/10455/src/net/netip/netip.go:1280
	// _ = "end of CoverTab[4379]"
//line /snap/go/10455/src/net/netip/netip.go:1280
	_go_fuzz_dep_.CoverTab[4380]++
							ip, err := ParseAddr(s[:i])
							if err != nil {
//line /snap/go/10455/src/net/netip/netip.go:1282
		_go_fuzz_dep_.CoverTab[527278]++
//line /snap/go/10455/src/net/netip/netip.go:1282
		_go_fuzz_dep_.CoverTab[4388]++
								return Prefix{}, errors.New("netip.ParsePrefix(" + strconv.Quote(s) + "): " + err.Error())
//line /snap/go/10455/src/net/netip/netip.go:1283
		// _ = "end of CoverTab[4388]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:1284
		_go_fuzz_dep_.CoverTab[527279]++
//line /snap/go/10455/src/net/netip/netip.go:1284
		_go_fuzz_dep_.CoverTab[4389]++
//line /snap/go/10455/src/net/netip/netip.go:1284
		// _ = "end of CoverTab[4389]"
//line /snap/go/10455/src/net/netip/netip.go:1284
	}
//line /snap/go/10455/src/net/netip/netip.go:1284
	// _ = "end of CoverTab[4380]"
//line /snap/go/10455/src/net/netip/netip.go:1284
	_go_fuzz_dep_.CoverTab[4381]++

							if ip.Is6() && func() bool {
//line /snap/go/10455/src/net/netip/netip.go:1286
		_go_fuzz_dep_.CoverTab[4390]++
//line /snap/go/10455/src/net/netip/netip.go:1286
		return ip.z != z6noz
//line /snap/go/10455/src/net/netip/netip.go:1286
		// _ = "end of CoverTab[4390]"
//line /snap/go/10455/src/net/netip/netip.go:1286
	}() {
//line /snap/go/10455/src/net/netip/netip.go:1286
		_go_fuzz_dep_.CoverTab[527280]++
//line /snap/go/10455/src/net/netip/netip.go:1286
		_go_fuzz_dep_.CoverTab[4391]++
								return Prefix{}, errors.New("netip.ParsePrefix(" + strconv.Quote(s) + "): IPv6 zones cannot be present in a prefix")
//line /snap/go/10455/src/net/netip/netip.go:1287
		// _ = "end of CoverTab[4391]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:1288
		_go_fuzz_dep_.CoverTab[527281]++
//line /snap/go/10455/src/net/netip/netip.go:1288
		_go_fuzz_dep_.CoverTab[4392]++
//line /snap/go/10455/src/net/netip/netip.go:1288
		// _ = "end of CoverTab[4392]"
//line /snap/go/10455/src/net/netip/netip.go:1288
	}
//line /snap/go/10455/src/net/netip/netip.go:1288
	// _ = "end of CoverTab[4381]"
//line /snap/go/10455/src/net/netip/netip.go:1288
	_go_fuzz_dep_.CoverTab[4382]++

							bitsStr := s[i+1:]
							bits, err := strconv.Atoi(bitsStr)
							if err != nil {
//line /snap/go/10455/src/net/netip/netip.go:1292
		_go_fuzz_dep_.CoverTab[527282]++
//line /snap/go/10455/src/net/netip/netip.go:1292
		_go_fuzz_dep_.CoverTab[4393]++
								return Prefix{}, errors.New("netip.ParsePrefix(" + strconv.Quote(s) + "): bad bits after slash: " + strconv.Quote(bitsStr))
//line /snap/go/10455/src/net/netip/netip.go:1293
		// _ = "end of CoverTab[4393]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:1294
		_go_fuzz_dep_.CoverTab[527283]++
//line /snap/go/10455/src/net/netip/netip.go:1294
		_go_fuzz_dep_.CoverTab[4394]++
//line /snap/go/10455/src/net/netip/netip.go:1294
		// _ = "end of CoverTab[4394]"
//line /snap/go/10455/src/net/netip/netip.go:1294
	}
//line /snap/go/10455/src/net/netip/netip.go:1294
	// _ = "end of CoverTab[4382]"
//line /snap/go/10455/src/net/netip/netip.go:1294
	_go_fuzz_dep_.CoverTab[4383]++
							maxBits := 32
							if ip.Is6() {
//line /snap/go/10455/src/net/netip/netip.go:1296
		_go_fuzz_dep_.CoverTab[527284]++
//line /snap/go/10455/src/net/netip/netip.go:1296
		_go_fuzz_dep_.CoverTab[4395]++
								maxBits = 128
//line /snap/go/10455/src/net/netip/netip.go:1297
		// _ = "end of CoverTab[4395]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:1298
		_go_fuzz_dep_.CoverTab[527285]++
//line /snap/go/10455/src/net/netip/netip.go:1298
		_go_fuzz_dep_.CoverTab[4396]++
//line /snap/go/10455/src/net/netip/netip.go:1298
		// _ = "end of CoverTab[4396]"
//line /snap/go/10455/src/net/netip/netip.go:1298
	}
//line /snap/go/10455/src/net/netip/netip.go:1298
	// _ = "end of CoverTab[4383]"
//line /snap/go/10455/src/net/netip/netip.go:1298
	_go_fuzz_dep_.CoverTab[4384]++
							if bits < 0 || func() bool {
//line /snap/go/10455/src/net/netip/netip.go:1299
		_go_fuzz_dep_.CoverTab[4397]++
//line /snap/go/10455/src/net/netip/netip.go:1299
		return bits > maxBits
//line /snap/go/10455/src/net/netip/netip.go:1299
		// _ = "end of CoverTab[4397]"
//line /snap/go/10455/src/net/netip/netip.go:1299
	}() {
//line /snap/go/10455/src/net/netip/netip.go:1299
		_go_fuzz_dep_.CoverTab[527286]++
//line /snap/go/10455/src/net/netip/netip.go:1299
		_go_fuzz_dep_.CoverTab[4398]++
								return Prefix{}, errors.New("netip.ParsePrefix(" + strconv.Quote(s) + "): prefix length out of range")
//line /snap/go/10455/src/net/netip/netip.go:1300
		// _ = "end of CoverTab[4398]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:1301
		_go_fuzz_dep_.CoverTab[527287]++
//line /snap/go/10455/src/net/netip/netip.go:1301
		_go_fuzz_dep_.CoverTab[4399]++
//line /snap/go/10455/src/net/netip/netip.go:1301
		// _ = "end of CoverTab[4399]"
//line /snap/go/10455/src/net/netip/netip.go:1301
	}
//line /snap/go/10455/src/net/netip/netip.go:1301
	// _ = "end of CoverTab[4384]"
//line /snap/go/10455/src/net/netip/netip.go:1301
	_go_fuzz_dep_.CoverTab[4385]++
							return PrefixFrom(ip, bits), nil
//line /snap/go/10455/src/net/netip/netip.go:1302
	// _ = "end of CoverTab[4385]"
}

// MustParsePrefix calls ParsePrefix(s) and panics on error.
//line /snap/go/10455/src/net/netip/netip.go:1305
// It is intended for use in tests with hard-coded strings.
//line /snap/go/10455/src/net/netip/netip.go:1307
func MustParsePrefix(s string) Prefix {
//line /snap/go/10455/src/net/netip/netip.go:1307
	_go_fuzz_dep_.CoverTab[4400]++
							ip, err := ParsePrefix(s)
							if err != nil {
//line /snap/go/10455/src/net/netip/netip.go:1309
		_go_fuzz_dep_.CoverTab[527288]++
//line /snap/go/10455/src/net/netip/netip.go:1309
		_go_fuzz_dep_.CoverTab[4402]++
								panic(err)
//line /snap/go/10455/src/net/netip/netip.go:1310
		// _ = "end of CoverTab[4402]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:1311
		_go_fuzz_dep_.CoverTab[527289]++
//line /snap/go/10455/src/net/netip/netip.go:1311
		_go_fuzz_dep_.CoverTab[4403]++
//line /snap/go/10455/src/net/netip/netip.go:1311
		// _ = "end of CoverTab[4403]"
//line /snap/go/10455/src/net/netip/netip.go:1311
	}
//line /snap/go/10455/src/net/netip/netip.go:1311
	// _ = "end of CoverTab[4400]"
//line /snap/go/10455/src/net/netip/netip.go:1311
	_go_fuzz_dep_.CoverTab[4401]++
							return ip
//line /snap/go/10455/src/net/netip/netip.go:1312
	// _ = "end of CoverTab[4401]"
}

// Masked returns p in its canonical form, with all but the high
//line /snap/go/10455/src/net/netip/netip.go:1315
// p.Bits() bits of p.Addr() masked off.
//line /snap/go/10455/src/net/netip/netip.go:1315
//
//line /snap/go/10455/src/net/netip/netip.go:1315
// If p is zero or otherwise invalid, Masked returns the zero Prefix.
//line /snap/go/10455/src/net/netip/netip.go:1319
func (p Prefix) Masked() Prefix {
//line /snap/go/10455/src/net/netip/netip.go:1319
	_go_fuzz_dep_.CoverTab[4404]++
							m, _ := p.ip.Prefix(p.Bits())
							return m
//line /snap/go/10455/src/net/netip/netip.go:1321
	// _ = "end of CoverTab[4404]"
}

// Contains reports whether the network p includes ip.
//line /snap/go/10455/src/net/netip/netip.go:1324
//
//line /snap/go/10455/src/net/netip/netip.go:1324
// An IPv4 address will not match an IPv6 prefix.
//line /snap/go/10455/src/net/netip/netip.go:1324
// An IPv4-mapped IPv6 address will not match an IPv4 prefix.
//line /snap/go/10455/src/net/netip/netip.go:1324
// A zero-value IP will not match any prefix.
//line /snap/go/10455/src/net/netip/netip.go:1324
// If ip has an IPv6 zone, Contains returns false,
//line /snap/go/10455/src/net/netip/netip.go:1324
// because Prefixes strip zones.
//line /snap/go/10455/src/net/netip/netip.go:1331
func (p Prefix) Contains(ip Addr) bool {
//line /snap/go/10455/src/net/netip/netip.go:1331
	_go_fuzz_dep_.CoverTab[4405]++
							if !p.IsValid() || func() bool {
//line /snap/go/10455/src/net/netip/netip.go:1332
		_go_fuzz_dep_.CoverTab[4408]++
//line /snap/go/10455/src/net/netip/netip.go:1332
		return ip.hasZone()
//line /snap/go/10455/src/net/netip/netip.go:1332
		// _ = "end of CoverTab[4408]"
//line /snap/go/10455/src/net/netip/netip.go:1332
	}() {
//line /snap/go/10455/src/net/netip/netip.go:1332
		_go_fuzz_dep_.CoverTab[527290]++
//line /snap/go/10455/src/net/netip/netip.go:1332
		_go_fuzz_dep_.CoverTab[4409]++
								return false
//line /snap/go/10455/src/net/netip/netip.go:1333
		// _ = "end of CoverTab[4409]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:1334
		_go_fuzz_dep_.CoverTab[527291]++
//line /snap/go/10455/src/net/netip/netip.go:1334
		_go_fuzz_dep_.CoverTab[4410]++
//line /snap/go/10455/src/net/netip/netip.go:1334
		// _ = "end of CoverTab[4410]"
//line /snap/go/10455/src/net/netip/netip.go:1334
	}
//line /snap/go/10455/src/net/netip/netip.go:1334
	// _ = "end of CoverTab[4405]"
//line /snap/go/10455/src/net/netip/netip.go:1334
	_go_fuzz_dep_.CoverTab[4406]++
							if f1, f2 := p.ip.BitLen(), ip.BitLen(); f1 == 0 || func() bool {
//line /snap/go/10455/src/net/netip/netip.go:1335
		_go_fuzz_dep_.CoverTab[4411]++
//line /snap/go/10455/src/net/netip/netip.go:1335
		return f2 == 0
//line /snap/go/10455/src/net/netip/netip.go:1335
		// _ = "end of CoverTab[4411]"
//line /snap/go/10455/src/net/netip/netip.go:1335
	}() || func() bool {
//line /snap/go/10455/src/net/netip/netip.go:1335
		_go_fuzz_dep_.CoverTab[4412]++
//line /snap/go/10455/src/net/netip/netip.go:1335
		return f1 != f2
//line /snap/go/10455/src/net/netip/netip.go:1335
		// _ = "end of CoverTab[4412]"
//line /snap/go/10455/src/net/netip/netip.go:1335
	}() {
//line /snap/go/10455/src/net/netip/netip.go:1335
		_go_fuzz_dep_.CoverTab[527292]++
//line /snap/go/10455/src/net/netip/netip.go:1335
		_go_fuzz_dep_.CoverTab[4413]++
								return false
//line /snap/go/10455/src/net/netip/netip.go:1336
		// _ = "end of CoverTab[4413]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:1337
		_go_fuzz_dep_.CoverTab[527293]++
//line /snap/go/10455/src/net/netip/netip.go:1337
		_go_fuzz_dep_.CoverTab[4414]++
//line /snap/go/10455/src/net/netip/netip.go:1337
		// _ = "end of CoverTab[4414]"
//line /snap/go/10455/src/net/netip/netip.go:1337
	}
//line /snap/go/10455/src/net/netip/netip.go:1337
	// _ = "end of CoverTab[4406]"
//line /snap/go/10455/src/net/netip/netip.go:1337
	_go_fuzz_dep_.CoverTab[4407]++
							if ip.Is4() {
//line /snap/go/10455/src/net/netip/netip.go:1338
		_go_fuzz_dep_.CoverTab[527294]++
//line /snap/go/10455/src/net/netip/netip.go:1338
		_go_fuzz_dep_.CoverTab[4415]++

//line /snap/go/10455/src/net/netip/netip.go:1347
		return uint32((ip.addr.lo^p.ip.addr.lo)>>((32-p.Bits())&63)) == 0
//line /snap/go/10455/src/net/netip/netip.go:1347
		// _ = "end of CoverTab[4415]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:1348
		_go_fuzz_dep_.CoverTab[527295]++
//line /snap/go/10455/src/net/netip/netip.go:1348
		_go_fuzz_dep_.CoverTab[4416]++

//line /snap/go/10455/src/net/netip/netip.go:1352
		return ip.addr.xor(p.ip.addr).and(mask6(p.Bits())).isZero()
//line /snap/go/10455/src/net/netip/netip.go:1352
		// _ = "end of CoverTab[4416]"
	}
//line /snap/go/10455/src/net/netip/netip.go:1353
	// _ = "end of CoverTab[4407]"
}

// Overlaps reports whether p and o contain any IP addresses in common.
//line /snap/go/10455/src/net/netip/netip.go:1356
//
//line /snap/go/10455/src/net/netip/netip.go:1356
// If p and o are of different address families or either have a zero
//line /snap/go/10455/src/net/netip/netip.go:1356
// IP, it reports false. Like the Contains method, a prefix with an
//line /snap/go/10455/src/net/netip/netip.go:1356
// IPv4-mapped IPv6 address is still treated as an IPv6 mask.
//line /snap/go/10455/src/net/netip/netip.go:1361
func (p Prefix) Overlaps(o Prefix) bool {
//line /snap/go/10455/src/net/netip/netip.go:1361
	_go_fuzz_dep_.CoverTab[4417]++
							if !p.IsValid() || func() bool {
//line /snap/go/10455/src/net/netip/netip.go:1362
		_go_fuzz_dep_.CoverTab[4425]++
//line /snap/go/10455/src/net/netip/netip.go:1362
		return !o.IsValid()
//line /snap/go/10455/src/net/netip/netip.go:1362
		// _ = "end of CoverTab[4425]"
//line /snap/go/10455/src/net/netip/netip.go:1362
	}() {
//line /snap/go/10455/src/net/netip/netip.go:1362
		_go_fuzz_dep_.CoverTab[527296]++
//line /snap/go/10455/src/net/netip/netip.go:1362
		_go_fuzz_dep_.CoverTab[4426]++
								return false
//line /snap/go/10455/src/net/netip/netip.go:1363
		// _ = "end of CoverTab[4426]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:1364
		_go_fuzz_dep_.CoverTab[527297]++
//line /snap/go/10455/src/net/netip/netip.go:1364
		_go_fuzz_dep_.CoverTab[4427]++
//line /snap/go/10455/src/net/netip/netip.go:1364
		// _ = "end of CoverTab[4427]"
//line /snap/go/10455/src/net/netip/netip.go:1364
	}
//line /snap/go/10455/src/net/netip/netip.go:1364
	// _ = "end of CoverTab[4417]"
//line /snap/go/10455/src/net/netip/netip.go:1364
	_go_fuzz_dep_.CoverTab[4418]++
							if p == o {
//line /snap/go/10455/src/net/netip/netip.go:1365
		_go_fuzz_dep_.CoverTab[527298]++
//line /snap/go/10455/src/net/netip/netip.go:1365
		_go_fuzz_dep_.CoverTab[4428]++
								return true
//line /snap/go/10455/src/net/netip/netip.go:1366
		// _ = "end of CoverTab[4428]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:1367
		_go_fuzz_dep_.CoverTab[527299]++
//line /snap/go/10455/src/net/netip/netip.go:1367
		_go_fuzz_dep_.CoverTab[4429]++
//line /snap/go/10455/src/net/netip/netip.go:1367
		// _ = "end of CoverTab[4429]"
//line /snap/go/10455/src/net/netip/netip.go:1367
	}
//line /snap/go/10455/src/net/netip/netip.go:1367
	// _ = "end of CoverTab[4418]"
//line /snap/go/10455/src/net/netip/netip.go:1367
	_go_fuzz_dep_.CoverTab[4419]++
							if p.ip.Is4() != o.ip.Is4() {
//line /snap/go/10455/src/net/netip/netip.go:1368
		_go_fuzz_dep_.CoverTab[527300]++
//line /snap/go/10455/src/net/netip/netip.go:1368
		_go_fuzz_dep_.CoverTab[4430]++
								return false
//line /snap/go/10455/src/net/netip/netip.go:1369
		// _ = "end of CoverTab[4430]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:1370
		_go_fuzz_dep_.CoverTab[527301]++
//line /snap/go/10455/src/net/netip/netip.go:1370
		_go_fuzz_dep_.CoverTab[4431]++
//line /snap/go/10455/src/net/netip/netip.go:1370
		// _ = "end of CoverTab[4431]"
//line /snap/go/10455/src/net/netip/netip.go:1370
	}
//line /snap/go/10455/src/net/netip/netip.go:1370
	// _ = "end of CoverTab[4419]"
//line /snap/go/10455/src/net/netip/netip.go:1370
	_go_fuzz_dep_.CoverTab[4420]++
							var minBits int
							if pb, ob := p.Bits(), o.Bits(); pb < ob {
//line /snap/go/10455/src/net/netip/netip.go:1372
		_go_fuzz_dep_.CoverTab[527302]++
//line /snap/go/10455/src/net/netip/netip.go:1372
		_go_fuzz_dep_.CoverTab[4432]++
								minBits = pb
//line /snap/go/10455/src/net/netip/netip.go:1373
		// _ = "end of CoverTab[4432]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:1374
		_go_fuzz_dep_.CoverTab[527303]++
//line /snap/go/10455/src/net/netip/netip.go:1374
		_go_fuzz_dep_.CoverTab[4433]++
								minBits = ob
//line /snap/go/10455/src/net/netip/netip.go:1375
		// _ = "end of CoverTab[4433]"
	}
//line /snap/go/10455/src/net/netip/netip.go:1376
	// _ = "end of CoverTab[4420]"
//line /snap/go/10455/src/net/netip/netip.go:1376
	_go_fuzz_dep_.CoverTab[4421]++
							if minBits == 0 {
//line /snap/go/10455/src/net/netip/netip.go:1377
		_go_fuzz_dep_.CoverTab[527304]++
//line /snap/go/10455/src/net/netip/netip.go:1377
		_go_fuzz_dep_.CoverTab[4434]++
								return true
//line /snap/go/10455/src/net/netip/netip.go:1378
		// _ = "end of CoverTab[4434]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:1379
		_go_fuzz_dep_.CoverTab[527305]++
//line /snap/go/10455/src/net/netip/netip.go:1379
		_go_fuzz_dep_.CoverTab[4435]++
//line /snap/go/10455/src/net/netip/netip.go:1379
		// _ = "end of CoverTab[4435]"
//line /snap/go/10455/src/net/netip/netip.go:1379
	}
//line /snap/go/10455/src/net/netip/netip.go:1379
	// _ = "end of CoverTab[4421]"
//line /snap/go/10455/src/net/netip/netip.go:1379
	_go_fuzz_dep_.CoverTab[4422]++
	// One of these Prefix calls might look redundant, but we don't require
	// that p and o values are normalized (via Prefix.Masked) first,
	// so the Prefix call on the one that's already minBits serves to zero
	// out any remaining bits in IP.
	var err error
	if p, err = p.ip.Prefix(minBits); err != nil {
//line /snap/go/10455/src/net/netip/netip.go:1385
		_go_fuzz_dep_.CoverTab[527306]++
//line /snap/go/10455/src/net/netip/netip.go:1385
		_go_fuzz_dep_.CoverTab[4436]++
								return false
//line /snap/go/10455/src/net/netip/netip.go:1386
		// _ = "end of CoverTab[4436]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:1387
		_go_fuzz_dep_.CoverTab[527307]++
//line /snap/go/10455/src/net/netip/netip.go:1387
		_go_fuzz_dep_.CoverTab[4437]++
//line /snap/go/10455/src/net/netip/netip.go:1387
		// _ = "end of CoverTab[4437]"
//line /snap/go/10455/src/net/netip/netip.go:1387
	}
//line /snap/go/10455/src/net/netip/netip.go:1387
	// _ = "end of CoverTab[4422]"
//line /snap/go/10455/src/net/netip/netip.go:1387
	_go_fuzz_dep_.CoverTab[4423]++
							if o, err = o.ip.Prefix(minBits); err != nil {
//line /snap/go/10455/src/net/netip/netip.go:1388
		_go_fuzz_dep_.CoverTab[527308]++
//line /snap/go/10455/src/net/netip/netip.go:1388
		_go_fuzz_dep_.CoverTab[4438]++
								return false
//line /snap/go/10455/src/net/netip/netip.go:1389
		// _ = "end of CoverTab[4438]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:1390
		_go_fuzz_dep_.CoverTab[527309]++
//line /snap/go/10455/src/net/netip/netip.go:1390
		_go_fuzz_dep_.CoverTab[4439]++
//line /snap/go/10455/src/net/netip/netip.go:1390
		// _ = "end of CoverTab[4439]"
//line /snap/go/10455/src/net/netip/netip.go:1390
	}
//line /snap/go/10455/src/net/netip/netip.go:1390
	// _ = "end of CoverTab[4423]"
//line /snap/go/10455/src/net/netip/netip.go:1390
	_go_fuzz_dep_.CoverTab[4424]++
							return p.ip == o.ip
//line /snap/go/10455/src/net/netip/netip.go:1391
	// _ = "end of CoverTab[4424]"
}

// AppendTo appends a text encoding of p,
//line /snap/go/10455/src/net/netip/netip.go:1394
// as generated by MarshalText,
//line /snap/go/10455/src/net/netip/netip.go:1394
// to b and returns the extended buffer.
//line /snap/go/10455/src/net/netip/netip.go:1397
func (p Prefix) AppendTo(b []byte) []byte {
//line /snap/go/10455/src/net/netip/netip.go:1397
	_go_fuzz_dep_.CoverTab[4440]++
							if p.isZero() {
//line /snap/go/10455/src/net/netip/netip.go:1398
		_go_fuzz_dep_.CoverTab[527310]++
//line /snap/go/10455/src/net/netip/netip.go:1398
		_go_fuzz_dep_.CoverTab[4444]++
								return b
//line /snap/go/10455/src/net/netip/netip.go:1399
		// _ = "end of CoverTab[4444]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:1400
		_go_fuzz_dep_.CoverTab[527311]++
//line /snap/go/10455/src/net/netip/netip.go:1400
		_go_fuzz_dep_.CoverTab[4445]++
//line /snap/go/10455/src/net/netip/netip.go:1400
		// _ = "end of CoverTab[4445]"
//line /snap/go/10455/src/net/netip/netip.go:1400
	}
//line /snap/go/10455/src/net/netip/netip.go:1400
	// _ = "end of CoverTab[4440]"
//line /snap/go/10455/src/net/netip/netip.go:1400
	_go_fuzz_dep_.CoverTab[4441]++
							if !p.IsValid() {
//line /snap/go/10455/src/net/netip/netip.go:1401
		_go_fuzz_dep_.CoverTab[527312]++
//line /snap/go/10455/src/net/netip/netip.go:1401
		_go_fuzz_dep_.CoverTab[4446]++
								return append(b, "invalid Prefix"...)
//line /snap/go/10455/src/net/netip/netip.go:1402
		// _ = "end of CoverTab[4446]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:1403
		_go_fuzz_dep_.CoverTab[527313]++
//line /snap/go/10455/src/net/netip/netip.go:1403
		_go_fuzz_dep_.CoverTab[4447]++
//line /snap/go/10455/src/net/netip/netip.go:1403
		// _ = "end of CoverTab[4447]"
//line /snap/go/10455/src/net/netip/netip.go:1403
	}
//line /snap/go/10455/src/net/netip/netip.go:1403
	// _ = "end of CoverTab[4441]"
//line /snap/go/10455/src/net/netip/netip.go:1403
	_go_fuzz_dep_.CoverTab[4442]++

//line /snap/go/10455/src/net/netip/netip.go:1406
	if p.ip.z == z4 {
//line /snap/go/10455/src/net/netip/netip.go:1406
		_go_fuzz_dep_.CoverTab[527314]++
//line /snap/go/10455/src/net/netip/netip.go:1406
		_go_fuzz_dep_.CoverTab[4448]++
								b = p.ip.appendTo4(b)
//line /snap/go/10455/src/net/netip/netip.go:1407
		// _ = "end of CoverTab[4448]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:1408
		_go_fuzz_dep_.CoverTab[527315]++
//line /snap/go/10455/src/net/netip/netip.go:1408
		_go_fuzz_dep_.CoverTab[4449]++
								if p.ip.Is4In6() {
//line /snap/go/10455/src/net/netip/netip.go:1409
			_go_fuzz_dep_.CoverTab[527316]++
//line /snap/go/10455/src/net/netip/netip.go:1409
			_go_fuzz_dep_.CoverTab[4450]++
									b = append(b, "::ffff:"...)
									b = p.ip.Unmap().appendTo4(b)
//line /snap/go/10455/src/net/netip/netip.go:1411
			// _ = "end of CoverTab[4450]"
		} else {
//line /snap/go/10455/src/net/netip/netip.go:1412
			_go_fuzz_dep_.CoverTab[527317]++
//line /snap/go/10455/src/net/netip/netip.go:1412
			_go_fuzz_dep_.CoverTab[4451]++
									b = p.ip.appendTo6(b)
//line /snap/go/10455/src/net/netip/netip.go:1413
			// _ = "end of CoverTab[4451]"
		}
//line /snap/go/10455/src/net/netip/netip.go:1414
		// _ = "end of CoverTab[4449]"
	}
//line /snap/go/10455/src/net/netip/netip.go:1415
	// _ = "end of CoverTab[4442]"
//line /snap/go/10455/src/net/netip/netip.go:1415
	_go_fuzz_dep_.CoverTab[4443]++

							b = append(b, '/')
							b = appendDecimal(b, uint8(p.Bits()))
							return b
//line /snap/go/10455/src/net/netip/netip.go:1419
	// _ = "end of CoverTab[4443]"
}

// MarshalText implements the encoding.TextMarshaler interface,
//line /snap/go/10455/src/net/netip/netip.go:1422
// The encoding is the same as returned by String, with one exception:
//line /snap/go/10455/src/net/netip/netip.go:1422
// If p is the zero value, the encoding is the empty string.
//line /snap/go/10455/src/net/netip/netip.go:1425
func (p Prefix) MarshalText() ([]byte, error) {
//line /snap/go/10455/src/net/netip/netip.go:1425
	_go_fuzz_dep_.CoverTab[4452]++
							var max int
							switch p.ip.z {
	case z0:
//line /snap/go/10455/src/net/netip/netip.go:1428
		_go_fuzz_dep_.CoverTab[527318]++
//line /snap/go/10455/src/net/netip/netip.go:1428
		_go_fuzz_dep_.CoverTab[4454]++
//line /snap/go/10455/src/net/netip/netip.go:1428
		// _ = "end of CoverTab[4454]"
	case z4:
//line /snap/go/10455/src/net/netip/netip.go:1429
		_go_fuzz_dep_.CoverTab[527319]++
//line /snap/go/10455/src/net/netip/netip.go:1429
		_go_fuzz_dep_.CoverTab[4455]++
								max = len("255.255.255.255/32")
//line /snap/go/10455/src/net/netip/netip.go:1430
		// _ = "end of CoverTab[4455]"
	default:
//line /snap/go/10455/src/net/netip/netip.go:1431
		_go_fuzz_dep_.CoverTab[527320]++
//line /snap/go/10455/src/net/netip/netip.go:1431
		_go_fuzz_dep_.CoverTab[4456]++
								max = len("ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff%enp5s0/128")
//line /snap/go/10455/src/net/netip/netip.go:1432
		// _ = "end of CoverTab[4456]"
	}
//line /snap/go/10455/src/net/netip/netip.go:1433
	// _ = "end of CoverTab[4452]"
//line /snap/go/10455/src/net/netip/netip.go:1433
	_go_fuzz_dep_.CoverTab[4453]++
							b := make([]byte, 0, max)
							b = p.AppendTo(b)
							return b, nil
//line /snap/go/10455/src/net/netip/netip.go:1436
	// _ = "end of CoverTab[4453]"
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
//line /snap/go/10455/src/net/netip/netip.go:1439
// The IP address is expected in a form accepted by ParsePrefix
//line /snap/go/10455/src/net/netip/netip.go:1439
// or generated by MarshalText.
//line /snap/go/10455/src/net/netip/netip.go:1442
func (p *Prefix) UnmarshalText(text []byte) error {
//line /snap/go/10455/src/net/netip/netip.go:1442
	_go_fuzz_dep_.CoverTab[4457]++
							if len(text) == 0 {
//line /snap/go/10455/src/net/netip/netip.go:1443
		_go_fuzz_dep_.CoverTab[527321]++
//line /snap/go/10455/src/net/netip/netip.go:1443
		_go_fuzz_dep_.CoverTab[4459]++
								*p = Prefix{}
								return nil
//line /snap/go/10455/src/net/netip/netip.go:1445
		// _ = "end of CoverTab[4459]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:1446
		_go_fuzz_dep_.CoverTab[527322]++
//line /snap/go/10455/src/net/netip/netip.go:1446
		_go_fuzz_dep_.CoverTab[4460]++
//line /snap/go/10455/src/net/netip/netip.go:1446
		// _ = "end of CoverTab[4460]"
//line /snap/go/10455/src/net/netip/netip.go:1446
	}
//line /snap/go/10455/src/net/netip/netip.go:1446
	// _ = "end of CoverTab[4457]"
//line /snap/go/10455/src/net/netip/netip.go:1446
	_go_fuzz_dep_.CoverTab[4458]++
							var err error
							*p, err = ParsePrefix(string(text))
							return err
//line /snap/go/10455/src/net/netip/netip.go:1449
	// _ = "end of CoverTab[4458]"
}

// MarshalBinary implements the encoding.BinaryMarshaler interface.
//line /snap/go/10455/src/net/netip/netip.go:1452
// It returns Addr.MarshalBinary with an additional byte appended
//line /snap/go/10455/src/net/netip/netip.go:1452
// containing the prefix bits.
//line /snap/go/10455/src/net/netip/netip.go:1455
func (p Prefix) MarshalBinary() ([]byte, error) {
//line /snap/go/10455/src/net/netip/netip.go:1455
	_go_fuzz_dep_.CoverTab[4461]++
							b := p.Addr().withoutZone().marshalBinaryWithTrailingBytes(1)
							b[len(b)-1] = uint8(p.Bits())
							return b, nil
//line /snap/go/10455/src/net/netip/netip.go:1458
	// _ = "end of CoverTab[4461]"
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface.
//line /snap/go/10455/src/net/netip/netip.go:1461
// It expects data in the form generated by MarshalBinary.
//line /snap/go/10455/src/net/netip/netip.go:1463
func (p *Prefix) UnmarshalBinary(b []byte) error {
//line /snap/go/10455/src/net/netip/netip.go:1463
	_go_fuzz_dep_.CoverTab[4462]++
							if len(b) < 1 {
//line /snap/go/10455/src/net/netip/netip.go:1464
		_go_fuzz_dep_.CoverTab[527323]++
//line /snap/go/10455/src/net/netip/netip.go:1464
		_go_fuzz_dep_.CoverTab[4465]++
								return errors.New("unexpected slice size")
//line /snap/go/10455/src/net/netip/netip.go:1465
		// _ = "end of CoverTab[4465]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:1466
		_go_fuzz_dep_.CoverTab[527324]++
//line /snap/go/10455/src/net/netip/netip.go:1466
		_go_fuzz_dep_.CoverTab[4466]++
//line /snap/go/10455/src/net/netip/netip.go:1466
		// _ = "end of CoverTab[4466]"
//line /snap/go/10455/src/net/netip/netip.go:1466
	}
//line /snap/go/10455/src/net/netip/netip.go:1466
	// _ = "end of CoverTab[4462]"
//line /snap/go/10455/src/net/netip/netip.go:1466
	_go_fuzz_dep_.CoverTab[4463]++
							var addr Addr
							err := addr.UnmarshalBinary(b[:len(b)-1])
							if err != nil {
//line /snap/go/10455/src/net/netip/netip.go:1469
		_go_fuzz_dep_.CoverTab[527325]++
//line /snap/go/10455/src/net/netip/netip.go:1469
		_go_fuzz_dep_.CoverTab[4467]++
								return err
//line /snap/go/10455/src/net/netip/netip.go:1470
		// _ = "end of CoverTab[4467]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:1471
		_go_fuzz_dep_.CoverTab[527326]++
//line /snap/go/10455/src/net/netip/netip.go:1471
		_go_fuzz_dep_.CoverTab[4468]++
//line /snap/go/10455/src/net/netip/netip.go:1471
		// _ = "end of CoverTab[4468]"
//line /snap/go/10455/src/net/netip/netip.go:1471
	}
//line /snap/go/10455/src/net/netip/netip.go:1471
	// _ = "end of CoverTab[4463]"
//line /snap/go/10455/src/net/netip/netip.go:1471
	_go_fuzz_dep_.CoverTab[4464]++
							*p = PrefixFrom(addr, int(b[len(b)-1]))
							return nil
//line /snap/go/10455/src/net/netip/netip.go:1473
	// _ = "end of CoverTab[4464]"
}

// String returns the CIDR notation of p: "<ip>/<bits>".
func (p Prefix) String() string {
//line /snap/go/10455/src/net/netip/netip.go:1477
	_go_fuzz_dep_.CoverTab[4469]++
							if !p.IsValid() {
//line /snap/go/10455/src/net/netip/netip.go:1478
		_go_fuzz_dep_.CoverTab[527327]++
//line /snap/go/10455/src/net/netip/netip.go:1478
		_go_fuzz_dep_.CoverTab[4471]++
								return "invalid Prefix"
//line /snap/go/10455/src/net/netip/netip.go:1479
		// _ = "end of CoverTab[4471]"
	} else {
//line /snap/go/10455/src/net/netip/netip.go:1480
		_go_fuzz_dep_.CoverTab[527328]++
//line /snap/go/10455/src/net/netip/netip.go:1480
		_go_fuzz_dep_.CoverTab[4472]++
//line /snap/go/10455/src/net/netip/netip.go:1480
		// _ = "end of CoverTab[4472]"
//line /snap/go/10455/src/net/netip/netip.go:1480
	}
//line /snap/go/10455/src/net/netip/netip.go:1480
	// _ = "end of CoverTab[4469]"
//line /snap/go/10455/src/net/netip/netip.go:1480
	_go_fuzz_dep_.CoverTab[4470]++
							return p.ip.String() + "/" + itoa.Itoa(p.Bits())
//line /snap/go/10455/src/net/netip/netip.go:1481
	// _ = "end of CoverTab[4470]"
}

//line /snap/go/10455/src/net/netip/netip.go:1482
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /snap/go/10455/src/net/netip/netip.go:1482
var _ = _go_fuzz_dep_.CoverTab
