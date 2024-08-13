// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/net/netip/netip.go:5
// Package netip defines an IP address type that's a small value type.
//line /usr/local/go/src/net/netip/netip.go:5
// Building on that Addr type, the package also defines AddrPort (an
//line /usr/local/go/src/net/netip/netip.go:5
// IP address and a port), and Prefix (an IP address and a bit length
//line /usr/local/go/src/net/netip/netip.go:5
// prefix).
//line /usr/local/go/src/net/netip/netip.go:5
//
//line /usr/local/go/src/net/netip/netip.go:5
// Compared to the net.IP type, this package's Addr type takes less
//line /usr/local/go/src/net/netip/netip.go:5
// memory, is immutable, and is comparable (supports == and being a
//line /usr/local/go/src/net/netip/netip.go:5
// map key).
//line /usr/local/go/src/net/netip/netip.go:13
package netip

//line /usr/local/go/src/net/netip/netip.go:13
import (
//line /usr/local/go/src/net/netip/netip.go:13
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/netip/netip.go:13
)
//line /usr/local/go/src/net/netip/netip.go:13
import (
//line /usr/local/go/src/net/netip/netip.go:13
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/netip/netip.go:13
)

import (
	"errors"
	"math"
	"strconv"

	"internal/bytealg"
	"internal/intern"
	"internal/itoa"
)

//line /usr/local/go/src/net/netip/netip.go:30
// Addr represents an IPv4 or IPv6 address (with or without a scoped
//line /usr/local/go/src/net/netip/netip.go:30
// addressing zone), similar to net.IP or net.IPAddr.
//line /usr/local/go/src/net/netip/netip.go:30
//
//line /usr/local/go/src/net/netip/netip.go:30
// Unlike net.IP or net.IPAddr, Addr is a comparable value
//line /usr/local/go/src/net/netip/netip.go:30
// type (it supports == and can be a map key) and is immutable.
//line /usr/local/go/src/net/netip/netip.go:30
//
//line /usr/local/go/src/net/netip/netip.go:30
// The zero Addr is not a valid IP address.
//line /usr/local/go/src/net/netip/netip.go:30
// Addr{} is distinct from both 0.0.0.0 and ::.
//line /usr/local/go/src/net/netip/netip.go:38
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
//line /usr/local/go/src/net/netip/netip.go:66
// See the Addr type's field docs.
//line /usr/local/go/src/net/netip/netip.go:68
var (
	z0	= (*intern.Value)(nil)
	z4	= new(intern.Value)
	z6noz	= new(intern.Value)
)

// IPv6LinkLocalAllNodes returns the IPv6 link-local all nodes multicast
//line /usr/local/go/src/net/netip/netip.go:74
// address ff02::1.
//line /usr/local/go/src/net/netip/netip.go:76
func IPv6LinkLocalAllNodes() Addr {
//line /usr/local/go/src/net/netip/netip.go:76
	_go_fuzz_dep_.CoverTab[11946]++
//line /usr/local/go/src/net/netip/netip.go:76
	return AddrFrom16([16]byte{0: 0xff, 1: 0x02, 15: 0x01})
//line /usr/local/go/src/net/netip/netip.go:76
	// _ = "end of CoverTab[11946]"
//line /usr/local/go/src/net/netip/netip.go:76
}

// IPv6LinkLocalAllRouters returns the IPv6 link-local all routers multicast
//line /usr/local/go/src/net/netip/netip.go:78
// address ff02::2.
//line /usr/local/go/src/net/netip/netip.go:80
func IPv6LinkLocalAllRouters() Addr {
//line /usr/local/go/src/net/netip/netip.go:80
	_go_fuzz_dep_.CoverTab[11947]++
//line /usr/local/go/src/net/netip/netip.go:80
	return AddrFrom16([16]byte{0: 0xff, 1: 0x02, 15: 0x02})
//line /usr/local/go/src/net/netip/netip.go:80
	// _ = "end of CoverTab[11947]"
//line /usr/local/go/src/net/netip/netip.go:80
}

// IPv6Loopback returns the IPv6 loopback address ::1.
func IPv6Loopback() Addr {
//line /usr/local/go/src/net/netip/netip.go:83
	_go_fuzz_dep_.CoverTab[11948]++
//line /usr/local/go/src/net/netip/netip.go:83
	return AddrFrom16([16]byte{15: 0x01})
//line /usr/local/go/src/net/netip/netip.go:83
	// _ = "end of CoverTab[11948]"
//line /usr/local/go/src/net/netip/netip.go:83
}

// IPv6Unspecified returns the IPv6 unspecified address "::".
func IPv6Unspecified() Addr {
//line /usr/local/go/src/net/netip/netip.go:86
	_go_fuzz_dep_.CoverTab[11949]++
//line /usr/local/go/src/net/netip/netip.go:86
	return Addr{z: z6noz}
//line /usr/local/go/src/net/netip/netip.go:86
	// _ = "end of CoverTab[11949]"
//line /usr/local/go/src/net/netip/netip.go:86
}

// IPv4Unspecified returns the IPv4 unspecified address "0.0.0.0".
func IPv4Unspecified() Addr {
//line /usr/local/go/src/net/netip/netip.go:89
	_go_fuzz_dep_.CoverTab[11950]++
//line /usr/local/go/src/net/netip/netip.go:89
	return AddrFrom4([4]byte{})
//line /usr/local/go/src/net/netip/netip.go:89
	// _ = "end of CoverTab[11950]"
//line /usr/local/go/src/net/netip/netip.go:89
}

// AddrFrom4 returns the address of the IPv4 address given by the bytes in addr.
func AddrFrom4(addr [4]byte) Addr {
//line /usr/local/go/src/net/netip/netip.go:92
	_go_fuzz_dep_.CoverTab[11951]++
						return Addr{
		addr:	uint128{0, 0xffff00000000 | uint64(addr[0])<<24 | uint64(addr[1])<<16 | uint64(addr[2])<<8 | uint64(addr[3])},
		z:	z4,
	}
//line /usr/local/go/src/net/netip/netip.go:96
	// _ = "end of CoverTab[11951]"
}

// AddrFrom16 returns the IPv6 address given by the bytes in addr.
//line /usr/local/go/src/net/netip/netip.go:99
// An IPv4-mapped IPv6 address is left as an IPv6 address.
//line /usr/local/go/src/net/netip/netip.go:99
// (Use Unmap to convert them if needed.)
//line /usr/local/go/src/net/netip/netip.go:102
func AddrFrom16(addr [16]byte) Addr {
//line /usr/local/go/src/net/netip/netip.go:102
	_go_fuzz_dep_.CoverTab[11952]++
							return Addr{
		addr: uint128{
			beUint64(addr[:8]),
			beUint64(addr[8:]),
		},
		z:	z6noz,
	}
//line /usr/local/go/src/net/netip/netip.go:109
	// _ = "end of CoverTab[11952]"
}

// ParseAddr parses s as an IP address, returning the result. The string
//line /usr/local/go/src/net/netip/netip.go:112
// s can be in dotted decimal ("192.0.2.1"), IPv6 ("2001:db8::68"),
//line /usr/local/go/src/net/netip/netip.go:112
// or IPv6 with a scoped addressing zone ("fe80::1cc0:3e8c:119f:c2e1%ens18").
//line /usr/local/go/src/net/netip/netip.go:115
func ParseAddr(s string) (Addr, error) {
//line /usr/local/go/src/net/netip/netip.go:115
	_go_fuzz_dep_.CoverTab[11953]++
							for i := 0; i < len(s); i++ {
//line /usr/local/go/src/net/netip/netip.go:116
		_go_fuzz_dep_.CoverTab[11955]++
								switch s[i] {
		case '.':
//line /usr/local/go/src/net/netip/netip.go:118
			_go_fuzz_dep_.CoverTab[11956]++
									return parseIPv4(s)
//line /usr/local/go/src/net/netip/netip.go:119
			// _ = "end of CoverTab[11956]"
		case ':':
//line /usr/local/go/src/net/netip/netip.go:120
			_go_fuzz_dep_.CoverTab[11957]++
									return parseIPv6(s)
//line /usr/local/go/src/net/netip/netip.go:121
			// _ = "end of CoverTab[11957]"
		case '%':
//line /usr/local/go/src/net/netip/netip.go:122
			_go_fuzz_dep_.CoverTab[11958]++

//line /usr/local/go/src/net/netip/netip.go:125
			return Addr{}, parseAddrError{in: s, msg: "missing IPv6 address"}
//line /usr/local/go/src/net/netip/netip.go:125
			// _ = "end of CoverTab[11958]"
//line /usr/local/go/src/net/netip/netip.go:125
		default:
//line /usr/local/go/src/net/netip/netip.go:125
			_go_fuzz_dep_.CoverTab[11959]++
//line /usr/local/go/src/net/netip/netip.go:125
			// _ = "end of CoverTab[11959]"
		}
//line /usr/local/go/src/net/netip/netip.go:126
		// _ = "end of CoverTab[11955]"
	}
//line /usr/local/go/src/net/netip/netip.go:127
	// _ = "end of CoverTab[11953]"
//line /usr/local/go/src/net/netip/netip.go:127
	_go_fuzz_dep_.CoverTab[11954]++
							return Addr{}, parseAddrError{in: s, msg: "unable to parse IP"}
//line /usr/local/go/src/net/netip/netip.go:128
	// _ = "end of CoverTab[11954]"
}

// MustParseAddr calls ParseAddr(s) and panics on error.
//line /usr/local/go/src/net/netip/netip.go:131
// It is intended for use in tests with hard-coded strings.
//line /usr/local/go/src/net/netip/netip.go:133
func MustParseAddr(s string) Addr {
//line /usr/local/go/src/net/netip/netip.go:133
	_go_fuzz_dep_.CoverTab[11960]++
							ip, err := ParseAddr(s)
							if err != nil {
//line /usr/local/go/src/net/netip/netip.go:135
		_go_fuzz_dep_.CoverTab[11962]++
								panic(err)
//line /usr/local/go/src/net/netip/netip.go:136
		// _ = "end of CoverTab[11962]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:137
		_go_fuzz_dep_.CoverTab[11963]++
//line /usr/local/go/src/net/netip/netip.go:137
		// _ = "end of CoverTab[11963]"
//line /usr/local/go/src/net/netip/netip.go:137
	}
//line /usr/local/go/src/net/netip/netip.go:137
	// _ = "end of CoverTab[11960]"
//line /usr/local/go/src/net/netip/netip.go:137
	_go_fuzz_dep_.CoverTab[11961]++
							return ip
//line /usr/local/go/src/net/netip/netip.go:138
	// _ = "end of CoverTab[11961]"
}

type parseAddrError struct {
	in	string	// the string given to ParseAddr
	msg	string	// an explanation of the parse failure
	at	string	// optionally, the unparsed portion of in at which the error occurred.
}

func (err parseAddrError) Error() string {
//line /usr/local/go/src/net/netip/netip.go:147
	_go_fuzz_dep_.CoverTab[11964]++
							q := strconv.Quote
							if err.at != "" {
//line /usr/local/go/src/net/netip/netip.go:149
		_go_fuzz_dep_.CoverTab[11966]++
								return "ParseAddr(" + q(err.in) + "): " + err.msg + " (at " + q(err.at) + ")"
//line /usr/local/go/src/net/netip/netip.go:150
		// _ = "end of CoverTab[11966]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:151
		_go_fuzz_dep_.CoverTab[11967]++
//line /usr/local/go/src/net/netip/netip.go:151
		// _ = "end of CoverTab[11967]"
//line /usr/local/go/src/net/netip/netip.go:151
	}
//line /usr/local/go/src/net/netip/netip.go:151
	// _ = "end of CoverTab[11964]"
//line /usr/local/go/src/net/netip/netip.go:151
	_go_fuzz_dep_.CoverTab[11965]++
							return "ParseAddr(" + q(err.in) + "): " + err.msg
//line /usr/local/go/src/net/netip/netip.go:152
	// _ = "end of CoverTab[11965]"
}

// parseIPv4 parses s as an IPv4 address (in form "192.168.0.1").
func parseIPv4(s string) (ip Addr, err error) {
//line /usr/local/go/src/net/netip/netip.go:156
	_go_fuzz_dep_.CoverTab[11968]++
							var fields [4]uint8
							var val, pos int
							var digLen int	// number of digits in current octet
							for i := 0; i < len(s); i++ {
//line /usr/local/go/src/net/netip/netip.go:160
		_go_fuzz_dep_.CoverTab[11971]++
								if s[i] >= '0' && func() bool {
//line /usr/local/go/src/net/netip/netip.go:161
			_go_fuzz_dep_.CoverTab[11972]++
//line /usr/local/go/src/net/netip/netip.go:161
			return s[i] <= '9'
//line /usr/local/go/src/net/netip/netip.go:161
			// _ = "end of CoverTab[11972]"
//line /usr/local/go/src/net/netip/netip.go:161
		}() {
//line /usr/local/go/src/net/netip/netip.go:161
			_go_fuzz_dep_.CoverTab[11973]++
									if digLen == 1 && func() bool {
//line /usr/local/go/src/net/netip/netip.go:162
				_go_fuzz_dep_.CoverTab[11975]++
//line /usr/local/go/src/net/netip/netip.go:162
				return val == 0
//line /usr/local/go/src/net/netip/netip.go:162
				// _ = "end of CoverTab[11975]"
//line /usr/local/go/src/net/netip/netip.go:162
			}() {
//line /usr/local/go/src/net/netip/netip.go:162
				_go_fuzz_dep_.CoverTab[11976]++
										return Addr{}, parseAddrError{in: s, msg: "IPv4 field has octet with leading zero"}
//line /usr/local/go/src/net/netip/netip.go:163
				// _ = "end of CoverTab[11976]"
			} else {
//line /usr/local/go/src/net/netip/netip.go:164
				_go_fuzz_dep_.CoverTab[11977]++
//line /usr/local/go/src/net/netip/netip.go:164
				// _ = "end of CoverTab[11977]"
//line /usr/local/go/src/net/netip/netip.go:164
			}
//line /usr/local/go/src/net/netip/netip.go:164
			// _ = "end of CoverTab[11973]"
//line /usr/local/go/src/net/netip/netip.go:164
			_go_fuzz_dep_.CoverTab[11974]++
									val = val*10 + int(s[i]) - '0'
									digLen++
									if val > 255 {
//line /usr/local/go/src/net/netip/netip.go:167
				_go_fuzz_dep_.CoverTab[11978]++
										return Addr{}, parseAddrError{in: s, msg: "IPv4 field has value >255"}
//line /usr/local/go/src/net/netip/netip.go:168
				// _ = "end of CoverTab[11978]"
			} else {
//line /usr/local/go/src/net/netip/netip.go:169
				_go_fuzz_dep_.CoverTab[11979]++
//line /usr/local/go/src/net/netip/netip.go:169
				// _ = "end of CoverTab[11979]"
//line /usr/local/go/src/net/netip/netip.go:169
			}
//line /usr/local/go/src/net/netip/netip.go:169
			// _ = "end of CoverTab[11974]"
		} else {
//line /usr/local/go/src/net/netip/netip.go:170
			_go_fuzz_dep_.CoverTab[11980]++
//line /usr/local/go/src/net/netip/netip.go:170
			if s[i] == '.' {
//line /usr/local/go/src/net/netip/netip.go:170
				_go_fuzz_dep_.CoverTab[11981]++

//line /usr/local/go/src/net/netip/netip.go:174
				if i == 0 || func() bool {
//line /usr/local/go/src/net/netip/netip.go:174
					_go_fuzz_dep_.CoverTab[11984]++
//line /usr/local/go/src/net/netip/netip.go:174
					return i == len(s)-1
//line /usr/local/go/src/net/netip/netip.go:174
					// _ = "end of CoverTab[11984]"
//line /usr/local/go/src/net/netip/netip.go:174
				}() || func() bool {
//line /usr/local/go/src/net/netip/netip.go:174
					_go_fuzz_dep_.CoverTab[11985]++
//line /usr/local/go/src/net/netip/netip.go:174
					return s[i-1] == '.'
//line /usr/local/go/src/net/netip/netip.go:174
					// _ = "end of CoverTab[11985]"
//line /usr/local/go/src/net/netip/netip.go:174
				}() {
//line /usr/local/go/src/net/netip/netip.go:174
					_go_fuzz_dep_.CoverTab[11986]++
											return Addr{}, parseAddrError{in: s, msg: "IPv4 field must have at least one digit", at: s[i:]}
//line /usr/local/go/src/net/netip/netip.go:175
					// _ = "end of CoverTab[11986]"
				} else {
//line /usr/local/go/src/net/netip/netip.go:176
					_go_fuzz_dep_.CoverTab[11987]++
//line /usr/local/go/src/net/netip/netip.go:176
					// _ = "end of CoverTab[11987]"
//line /usr/local/go/src/net/netip/netip.go:176
				}
//line /usr/local/go/src/net/netip/netip.go:176
				// _ = "end of CoverTab[11981]"
//line /usr/local/go/src/net/netip/netip.go:176
				_go_fuzz_dep_.CoverTab[11982]++

										if pos == 3 {
//line /usr/local/go/src/net/netip/netip.go:178
					_go_fuzz_dep_.CoverTab[11988]++
											return Addr{}, parseAddrError{in: s, msg: "IPv4 address too long"}
//line /usr/local/go/src/net/netip/netip.go:179
					// _ = "end of CoverTab[11988]"
				} else {
//line /usr/local/go/src/net/netip/netip.go:180
					_go_fuzz_dep_.CoverTab[11989]++
//line /usr/local/go/src/net/netip/netip.go:180
					// _ = "end of CoverTab[11989]"
//line /usr/local/go/src/net/netip/netip.go:180
				}
//line /usr/local/go/src/net/netip/netip.go:180
				// _ = "end of CoverTab[11982]"
//line /usr/local/go/src/net/netip/netip.go:180
				_go_fuzz_dep_.CoverTab[11983]++
										fields[pos] = uint8(val)
										pos++
										val = 0
										digLen = 0
//line /usr/local/go/src/net/netip/netip.go:184
				// _ = "end of CoverTab[11983]"
			} else {
//line /usr/local/go/src/net/netip/netip.go:185
				_go_fuzz_dep_.CoverTab[11990]++
										return Addr{}, parseAddrError{in: s, msg: "unexpected character", at: s[i:]}
//line /usr/local/go/src/net/netip/netip.go:186
				// _ = "end of CoverTab[11990]"
			}
//line /usr/local/go/src/net/netip/netip.go:187
			// _ = "end of CoverTab[11980]"
//line /usr/local/go/src/net/netip/netip.go:187
		}
//line /usr/local/go/src/net/netip/netip.go:187
		// _ = "end of CoverTab[11971]"
	}
//line /usr/local/go/src/net/netip/netip.go:188
	// _ = "end of CoverTab[11968]"
//line /usr/local/go/src/net/netip/netip.go:188
	_go_fuzz_dep_.CoverTab[11969]++
							if pos < 3 {
//line /usr/local/go/src/net/netip/netip.go:189
		_go_fuzz_dep_.CoverTab[11991]++
								return Addr{}, parseAddrError{in: s, msg: "IPv4 address too short"}
//line /usr/local/go/src/net/netip/netip.go:190
		// _ = "end of CoverTab[11991]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:191
		_go_fuzz_dep_.CoverTab[11992]++
//line /usr/local/go/src/net/netip/netip.go:191
		// _ = "end of CoverTab[11992]"
//line /usr/local/go/src/net/netip/netip.go:191
	}
//line /usr/local/go/src/net/netip/netip.go:191
	// _ = "end of CoverTab[11969]"
//line /usr/local/go/src/net/netip/netip.go:191
	_go_fuzz_dep_.CoverTab[11970]++
							fields[3] = uint8(val)
							return AddrFrom4(fields), nil
//line /usr/local/go/src/net/netip/netip.go:193
	// _ = "end of CoverTab[11970]"
}

// parseIPv6 parses s as an IPv6 address (in form "2001:db8::68").
func parseIPv6(in string) (Addr, error) {
//line /usr/local/go/src/net/netip/netip.go:197
	_go_fuzz_dep_.CoverTab[11993]++
							s := in

//line /usr/local/go/src/net/netip/netip.go:204
	zone := ""
	i := bytealg.IndexByteString(s, '%')
	if i != -1 {
//line /usr/local/go/src/net/netip/netip.go:206
		_go_fuzz_dep_.CoverTab[11999]++
								s, zone = s[:i], s[i+1:]
								if zone == "" {
//line /usr/local/go/src/net/netip/netip.go:208
			_go_fuzz_dep_.CoverTab[12000]++

									return Addr{}, parseAddrError{in: in, msg: "zone must be a non-empty string"}
//line /usr/local/go/src/net/netip/netip.go:210
			// _ = "end of CoverTab[12000]"
		} else {
//line /usr/local/go/src/net/netip/netip.go:211
			_go_fuzz_dep_.CoverTab[12001]++
//line /usr/local/go/src/net/netip/netip.go:211
			// _ = "end of CoverTab[12001]"
//line /usr/local/go/src/net/netip/netip.go:211
		}
//line /usr/local/go/src/net/netip/netip.go:211
		// _ = "end of CoverTab[11999]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:212
		_go_fuzz_dep_.CoverTab[12002]++
//line /usr/local/go/src/net/netip/netip.go:212
		// _ = "end of CoverTab[12002]"
//line /usr/local/go/src/net/netip/netip.go:212
	}
//line /usr/local/go/src/net/netip/netip.go:212
	// _ = "end of CoverTab[11993]"
//line /usr/local/go/src/net/netip/netip.go:212
	_go_fuzz_dep_.CoverTab[11994]++

							var ip [16]byte
							ellipsis := -1

//line /usr/local/go/src/net/netip/netip.go:218
	if len(s) >= 2 && func() bool {
//line /usr/local/go/src/net/netip/netip.go:218
		_go_fuzz_dep_.CoverTab[12003]++
//line /usr/local/go/src/net/netip/netip.go:218
		return s[0] == ':'
//line /usr/local/go/src/net/netip/netip.go:218
		// _ = "end of CoverTab[12003]"
//line /usr/local/go/src/net/netip/netip.go:218
	}() && func() bool {
//line /usr/local/go/src/net/netip/netip.go:218
		_go_fuzz_dep_.CoverTab[12004]++
//line /usr/local/go/src/net/netip/netip.go:218
		return s[1] == ':'
//line /usr/local/go/src/net/netip/netip.go:218
		// _ = "end of CoverTab[12004]"
//line /usr/local/go/src/net/netip/netip.go:218
	}() {
//line /usr/local/go/src/net/netip/netip.go:218
		_go_fuzz_dep_.CoverTab[12005]++
								ellipsis = 0
								s = s[2:]

								if len(s) == 0 {
//line /usr/local/go/src/net/netip/netip.go:222
			_go_fuzz_dep_.CoverTab[12006]++
									return IPv6Unspecified().WithZone(zone), nil
//line /usr/local/go/src/net/netip/netip.go:223
			// _ = "end of CoverTab[12006]"
		} else {
//line /usr/local/go/src/net/netip/netip.go:224
			_go_fuzz_dep_.CoverTab[12007]++
//line /usr/local/go/src/net/netip/netip.go:224
			// _ = "end of CoverTab[12007]"
//line /usr/local/go/src/net/netip/netip.go:224
		}
//line /usr/local/go/src/net/netip/netip.go:224
		// _ = "end of CoverTab[12005]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:225
		_go_fuzz_dep_.CoverTab[12008]++
//line /usr/local/go/src/net/netip/netip.go:225
		// _ = "end of CoverTab[12008]"
//line /usr/local/go/src/net/netip/netip.go:225
	}
//line /usr/local/go/src/net/netip/netip.go:225
	// _ = "end of CoverTab[11994]"
//line /usr/local/go/src/net/netip/netip.go:225
	_go_fuzz_dep_.CoverTab[11995]++

//line /usr/local/go/src/net/netip/netip.go:228
	i = 0
	for i < 16 {
//line /usr/local/go/src/net/netip/netip.go:229
		_go_fuzz_dep_.CoverTab[12009]++

//line /usr/local/go/src/net/netip/netip.go:232
		off := 0
		acc := uint32(0)
		for ; off < len(s); off++ {
//line /usr/local/go/src/net/netip/netip.go:234
			_go_fuzz_dep_.CoverTab[12015]++
									c := s[off]
									if c >= '0' && func() bool {
//line /usr/local/go/src/net/netip/netip.go:236
				_go_fuzz_dep_.CoverTab[12017]++
//line /usr/local/go/src/net/netip/netip.go:236
				return c <= '9'
//line /usr/local/go/src/net/netip/netip.go:236
				// _ = "end of CoverTab[12017]"
//line /usr/local/go/src/net/netip/netip.go:236
			}() {
//line /usr/local/go/src/net/netip/netip.go:236
				_go_fuzz_dep_.CoverTab[12018]++
										acc = (acc << 4) + uint32(c-'0')
//line /usr/local/go/src/net/netip/netip.go:237
				// _ = "end of CoverTab[12018]"
			} else {
//line /usr/local/go/src/net/netip/netip.go:238
				_go_fuzz_dep_.CoverTab[12019]++
//line /usr/local/go/src/net/netip/netip.go:238
				if c >= 'a' && func() bool {
//line /usr/local/go/src/net/netip/netip.go:238
					_go_fuzz_dep_.CoverTab[12020]++
//line /usr/local/go/src/net/netip/netip.go:238
					return c <= 'f'
//line /usr/local/go/src/net/netip/netip.go:238
					// _ = "end of CoverTab[12020]"
//line /usr/local/go/src/net/netip/netip.go:238
				}() {
//line /usr/local/go/src/net/netip/netip.go:238
					_go_fuzz_dep_.CoverTab[12021]++
											acc = (acc << 4) + uint32(c-'a'+10)
//line /usr/local/go/src/net/netip/netip.go:239
					// _ = "end of CoverTab[12021]"
				} else {
//line /usr/local/go/src/net/netip/netip.go:240
					_go_fuzz_dep_.CoverTab[12022]++
//line /usr/local/go/src/net/netip/netip.go:240
					if c >= 'A' && func() bool {
//line /usr/local/go/src/net/netip/netip.go:240
						_go_fuzz_dep_.CoverTab[12023]++
//line /usr/local/go/src/net/netip/netip.go:240
						return c <= 'F'
//line /usr/local/go/src/net/netip/netip.go:240
						// _ = "end of CoverTab[12023]"
//line /usr/local/go/src/net/netip/netip.go:240
					}() {
//line /usr/local/go/src/net/netip/netip.go:240
						_go_fuzz_dep_.CoverTab[12024]++
												acc = (acc << 4) + uint32(c-'A'+10)
//line /usr/local/go/src/net/netip/netip.go:241
						// _ = "end of CoverTab[12024]"
					} else {
//line /usr/local/go/src/net/netip/netip.go:242
						_go_fuzz_dep_.CoverTab[12025]++
												break
//line /usr/local/go/src/net/netip/netip.go:243
						// _ = "end of CoverTab[12025]"
					}
//line /usr/local/go/src/net/netip/netip.go:244
					// _ = "end of CoverTab[12022]"
//line /usr/local/go/src/net/netip/netip.go:244
				}
//line /usr/local/go/src/net/netip/netip.go:244
				// _ = "end of CoverTab[12019]"
//line /usr/local/go/src/net/netip/netip.go:244
			}
//line /usr/local/go/src/net/netip/netip.go:244
			// _ = "end of CoverTab[12015]"
//line /usr/local/go/src/net/netip/netip.go:244
			_go_fuzz_dep_.CoverTab[12016]++
									if acc > math.MaxUint16 {
//line /usr/local/go/src/net/netip/netip.go:245
				_go_fuzz_dep_.CoverTab[12026]++

										return Addr{}, parseAddrError{in: in, msg: "IPv6 field has value >=2^16", at: s}
//line /usr/local/go/src/net/netip/netip.go:247
				// _ = "end of CoverTab[12026]"
			} else {
//line /usr/local/go/src/net/netip/netip.go:248
				_go_fuzz_dep_.CoverTab[12027]++
//line /usr/local/go/src/net/netip/netip.go:248
				// _ = "end of CoverTab[12027]"
//line /usr/local/go/src/net/netip/netip.go:248
			}
//line /usr/local/go/src/net/netip/netip.go:248
			// _ = "end of CoverTab[12016]"
		}
//line /usr/local/go/src/net/netip/netip.go:249
		// _ = "end of CoverTab[12009]"
//line /usr/local/go/src/net/netip/netip.go:249
		_go_fuzz_dep_.CoverTab[12010]++
								if off == 0 {
//line /usr/local/go/src/net/netip/netip.go:250
			_go_fuzz_dep_.CoverTab[12028]++

									return Addr{}, parseAddrError{in: in, msg: "each colon-separated field must have at least one digit", at: s}
//line /usr/local/go/src/net/netip/netip.go:252
			// _ = "end of CoverTab[12028]"
		} else {
//line /usr/local/go/src/net/netip/netip.go:253
			_go_fuzz_dep_.CoverTab[12029]++
//line /usr/local/go/src/net/netip/netip.go:253
			// _ = "end of CoverTab[12029]"
//line /usr/local/go/src/net/netip/netip.go:253
		}
//line /usr/local/go/src/net/netip/netip.go:253
		// _ = "end of CoverTab[12010]"
//line /usr/local/go/src/net/netip/netip.go:253
		_go_fuzz_dep_.CoverTab[12011]++

//line /usr/local/go/src/net/netip/netip.go:256
		if off < len(s) && func() bool {
//line /usr/local/go/src/net/netip/netip.go:256
			_go_fuzz_dep_.CoverTab[12030]++
//line /usr/local/go/src/net/netip/netip.go:256
			return s[off] == '.'
//line /usr/local/go/src/net/netip/netip.go:256
			// _ = "end of CoverTab[12030]"
//line /usr/local/go/src/net/netip/netip.go:256
		}() {
//line /usr/local/go/src/net/netip/netip.go:256
			_go_fuzz_dep_.CoverTab[12031]++
									if ellipsis < 0 && func() bool {
//line /usr/local/go/src/net/netip/netip.go:257
				_go_fuzz_dep_.CoverTab[12035]++
//line /usr/local/go/src/net/netip/netip.go:257
				return i != 12
//line /usr/local/go/src/net/netip/netip.go:257
				// _ = "end of CoverTab[12035]"
//line /usr/local/go/src/net/netip/netip.go:257
			}() {
//line /usr/local/go/src/net/netip/netip.go:257
				_go_fuzz_dep_.CoverTab[12036]++

										return Addr{}, parseAddrError{in: in, msg: "embedded IPv4 address must replace the final 2 fields of the address", at: s}
//line /usr/local/go/src/net/netip/netip.go:259
				// _ = "end of CoverTab[12036]"
			} else {
//line /usr/local/go/src/net/netip/netip.go:260
				_go_fuzz_dep_.CoverTab[12037]++
//line /usr/local/go/src/net/netip/netip.go:260
				// _ = "end of CoverTab[12037]"
//line /usr/local/go/src/net/netip/netip.go:260
			}
//line /usr/local/go/src/net/netip/netip.go:260
			// _ = "end of CoverTab[12031]"
//line /usr/local/go/src/net/netip/netip.go:260
			_go_fuzz_dep_.CoverTab[12032]++
									if i+4 > 16 {
//line /usr/local/go/src/net/netip/netip.go:261
				_go_fuzz_dep_.CoverTab[12038]++

										return Addr{}, parseAddrError{in: in, msg: "too many hex fields to fit an embedded IPv4 at the end of the address", at: s}
//line /usr/local/go/src/net/netip/netip.go:263
				// _ = "end of CoverTab[12038]"
			} else {
//line /usr/local/go/src/net/netip/netip.go:264
				_go_fuzz_dep_.CoverTab[12039]++
//line /usr/local/go/src/net/netip/netip.go:264
				// _ = "end of CoverTab[12039]"
//line /usr/local/go/src/net/netip/netip.go:264
			}
//line /usr/local/go/src/net/netip/netip.go:264
			// _ = "end of CoverTab[12032]"
//line /usr/local/go/src/net/netip/netip.go:264
			_go_fuzz_dep_.CoverTab[12033]++

//line /usr/local/go/src/net/netip/netip.go:268
			ip4, err := parseIPv4(s)
			if err != nil {
//line /usr/local/go/src/net/netip/netip.go:269
				_go_fuzz_dep_.CoverTab[12040]++
										return Addr{}, parseAddrError{in: in, msg: err.Error(), at: s}
//line /usr/local/go/src/net/netip/netip.go:270
				// _ = "end of CoverTab[12040]"
			} else {
//line /usr/local/go/src/net/netip/netip.go:271
				_go_fuzz_dep_.CoverTab[12041]++
//line /usr/local/go/src/net/netip/netip.go:271
				// _ = "end of CoverTab[12041]"
//line /usr/local/go/src/net/netip/netip.go:271
			}
//line /usr/local/go/src/net/netip/netip.go:271
			// _ = "end of CoverTab[12033]"
//line /usr/local/go/src/net/netip/netip.go:271
			_go_fuzz_dep_.CoverTab[12034]++
									ip[i] = ip4.v4(0)
									ip[i+1] = ip4.v4(1)
									ip[i+2] = ip4.v4(2)
									ip[i+3] = ip4.v4(3)
									s = ""
									i += 4
									break
//line /usr/local/go/src/net/netip/netip.go:278
			// _ = "end of CoverTab[12034]"
		} else {
//line /usr/local/go/src/net/netip/netip.go:279
			_go_fuzz_dep_.CoverTab[12042]++
//line /usr/local/go/src/net/netip/netip.go:279
			// _ = "end of CoverTab[12042]"
//line /usr/local/go/src/net/netip/netip.go:279
		}
//line /usr/local/go/src/net/netip/netip.go:279
		// _ = "end of CoverTab[12011]"
//line /usr/local/go/src/net/netip/netip.go:279
		_go_fuzz_dep_.CoverTab[12012]++

//line /usr/local/go/src/net/netip/netip.go:282
		ip[i] = byte(acc >> 8)
								ip[i+1] = byte(acc)
								i += 2

//line /usr/local/go/src/net/netip/netip.go:287
		s = s[off:]
		if len(s) == 0 {
//line /usr/local/go/src/net/netip/netip.go:288
			_go_fuzz_dep_.CoverTab[12043]++
									break
//line /usr/local/go/src/net/netip/netip.go:289
			// _ = "end of CoverTab[12043]"
		} else {
//line /usr/local/go/src/net/netip/netip.go:290
			_go_fuzz_dep_.CoverTab[12044]++
//line /usr/local/go/src/net/netip/netip.go:290
			// _ = "end of CoverTab[12044]"
//line /usr/local/go/src/net/netip/netip.go:290
		}
//line /usr/local/go/src/net/netip/netip.go:290
		// _ = "end of CoverTab[12012]"
//line /usr/local/go/src/net/netip/netip.go:290
		_go_fuzz_dep_.CoverTab[12013]++

//line /usr/local/go/src/net/netip/netip.go:293
		if s[0] != ':' {
//line /usr/local/go/src/net/netip/netip.go:293
			_go_fuzz_dep_.CoverTab[12045]++
									return Addr{}, parseAddrError{in: in, msg: "unexpected character, want colon", at: s}
//line /usr/local/go/src/net/netip/netip.go:294
			// _ = "end of CoverTab[12045]"
		} else {
//line /usr/local/go/src/net/netip/netip.go:295
			_go_fuzz_dep_.CoverTab[12046]++
//line /usr/local/go/src/net/netip/netip.go:295
			if len(s) == 1 {
//line /usr/local/go/src/net/netip/netip.go:295
				_go_fuzz_dep_.CoverTab[12047]++
										return Addr{}, parseAddrError{in: in, msg: "colon must be followed by more characters", at: s}
//line /usr/local/go/src/net/netip/netip.go:296
				// _ = "end of CoverTab[12047]"
			} else {
//line /usr/local/go/src/net/netip/netip.go:297
				_go_fuzz_dep_.CoverTab[12048]++
//line /usr/local/go/src/net/netip/netip.go:297
				// _ = "end of CoverTab[12048]"
//line /usr/local/go/src/net/netip/netip.go:297
			}
//line /usr/local/go/src/net/netip/netip.go:297
			// _ = "end of CoverTab[12046]"
//line /usr/local/go/src/net/netip/netip.go:297
		}
//line /usr/local/go/src/net/netip/netip.go:297
		// _ = "end of CoverTab[12013]"
//line /usr/local/go/src/net/netip/netip.go:297
		_go_fuzz_dep_.CoverTab[12014]++
								s = s[1:]

//line /usr/local/go/src/net/netip/netip.go:301
		if s[0] == ':' {
//line /usr/local/go/src/net/netip/netip.go:301
			_go_fuzz_dep_.CoverTab[12049]++
									if ellipsis >= 0 {
//line /usr/local/go/src/net/netip/netip.go:302
				_go_fuzz_dep_.CoverTab[12051]++
										return Addr{}, parseAddrError{in: in, msg: "multiple :: in address", at: s}
//line /usr/local/go/src/net/netip/netip.go:303
				// _ = "end of CoverTab[12051]"
			} else {
//line /usr/local/go/src/net/netip/netip.go:304
				_go_fuzz_dep_.CoverTab[12052]++
//line /usr/local/go/src/net/netip/netip.go:304
				// _ = "end of CoverTab[12052]"
//line /usr/local/go/src/net/netip/netip.go:304
			}
//line /usr/local/go/src/net/netip/netip.go:304
			// _ = "end of CoverTab[12049]"
//line /usr/local/go/src/net/netip/netip.go:304
			_go_fuzz_dep_.CoverTab[12050]++
									ellipsis = i
									s = s[1:]
									if len(s) == 0 {
//line /usr/local/go/src/net/netip/netip.go:307
				_go_fuzz_dep_.CoverTab[12053]++
										break
//line /usr/local/go/src/net/netip/netip.go:308
				// _ = "end of CoverTab[12053]"
			} else {
//line /usr/local/go/src/net/netip/netip.go:309
				_go_fuzz_dep_.CoverTab[12054]++
//line /usr/local/go/src/net/netip/netip.go:309
				// _ = "end of CoverTab[12054]"
//line /usr/local/go/src/net/netip/netip.go:309
			}
//line /usr/local/go/src/net/netip/netip.go:309
			// _ = "end of CoverTab[12050]"
		} else {
//line /usr/local/go/src/net/netip/netip.go:310
			_go_fuzz_dep_.CoverTab[12055]++
//line /usr/local/go/src/net/netip/netip.go:310
			// _ = "end of CoverTab[12055]"
//line /usr/local/go/src/net/netip/netip.go:310
		}
//line /usr/local/go/src/net/netip/netip.go:310
		// _ = "end of CoverTab[12014]"
	}
//line /usr/local/go/src/net/netip/netip.go:311
	// _ = "end of CoverTab[11995]"
//line /usr/local/go/src/net/netip/netip.go:311
	_go_fuzz_dep_.CoverTab[11996]++

//line /usr/local/go/src/net/netip/netip.go:314
	if len(s) != 0 {
//line /usr/local/go/src/net/netip/netip.go:314
		_go_fuzz_dep_.CoverTab[12056]++
								return Addr{}, parseAddrError{in: in, msg: "trailing garbage after address", at: s}
//line /usr/local/go/src/net/netip/netip.go:315
		// _ = "end of CoverTab[12056]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:316
		_go_fuzz_dep_.CoverTab[12057]++
//line /usr/local/go/src/net/netip/netip.go:316
		// _ = "end of CoverTab[12057]"
//line /usr/local/go/src/net/netip/netip.go:316
	}
//line /usr/local/go/src/net/netip/netip.go:316
	// _ = "end of CoverTab[11996]"
//line /usr/local/go/src/net/netip/netip.go:316
	_go_fuzz_dep_.CoverTab[11997]++

//line /usr/local/go/src/net/netip/netip.go:319
	if i < 16 {
//line /usr/local/go/src/net/netip/netip.go:319
		_go_fuzz_dep_.CoverTab[12058]++
								if ellipsis < 0 {
//line /usr/local/go/src/net/netip/netip.go:320
			_go_fuzz_dep_.CoverTab[12061]++
									return Addr{}, parseAddrError{in: in, msg: "address string too short"}
//line /usr/local/go/src/net/netip/netip.go:321
			// _ = "end of CoverTab[12061]"
		} else {
//line /usr/local/go/src/net/netip/netip.go:322
			_go_fuzz_dep_.CoverTab[12062]++
//line /usr/local/go/src/net/netip/netip.go:322
			// _ = "end of CoverTab[12062]"
//line /usr/local/go/src/net/netip/netip.go:322
		}
//line /usr/local/go/src/net/netip/netip.go:322
		// _ = "end of CoverTab[12058]"
//line /usr/local/go/src/net/netip/netip.go:322
		_go_fuzz_dep_.CoverTab[12059]++
								n := 16 - i
								for j := i - 1; j >= ellipsis; j-- {
//line /usr/local/go/src/net/netip/netip.go:324
			_go_fuzz_dep_.CoverTab[12063]++
									ip[j+n] = ip[j]
//line /usr/local/go/src/net/netip/netip.go:325
			// _ = "end of CoverTab[12063]"
		}
//line /usr/local/go/src/net/netip/netip.go:326
		// _ = "end of CoverTab[12059]"
//line /usr/local/go/src/net/netip/netip.go:326
		_go_fuzz_dep_.CoverTab[12060]++
								for j := ellipsis + n - 1; j >= ellipsis; j-- {
//line /usr/local/go/src/net/netip/netip.go:327
			_go_fuzz_dep_.CoverTab[12064]++
									ip[j] = 0
//line /usr/local/go/src/net/netip/netip.go:328
			// _ = "end of CoverTab[12064]"
		}
//line /usr/local/go/src/net/netip/netip.go:329
		// _ = "end of CoverTab[12060]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:330
		_go_fuzz_dep_.CoverTab[12065]++
//line /usr/local/go/src/net/netip/netip.go:330
		if ellipsis >= 0 {
//line /usr/local/go/src/net/netip/netip.go:330
			_go_fuzz_dep_.CoverTab[12066]++

									return Addr{}, parseAddrError{in: in, msg: "the :: must expand to at least one field of zeros"}
//line /usr/local/go/src/net/netip/netip.go:332
			// _ = "end of CoverTab[12066]"
		} else {
//line /usr/local/go/src/net/netip/netip.go:333
			_go_fuzz_dep_.CoverTab[12067]++
//line /usr/local/go/src/net/netip/netip.go:333
			// _ = "end of CoverTab[12067]"
//line /usr/local/go/src/net/netip/netip.go:333
		}
//line /usr/local/go/src/net/netip/netip.go:333
		// _ = "end of CoverTab[12065]"
//line /usr/local/go/src/net/netip/netip.go:333
	}
//line /usr/local/go/src/net/netip/netip.go:333
	// _ = "end of CoverTab[11997]"
//line /usr/local/go/src/net/netip/netip.go:333
	_go_fuzz_dep_.CoverTab[11998]++
							return AddrFrom16(ip).WithZone(zone), nil
//line /usr/local/go/src/net/netip/netip.go:334
	// _ = "end of CoverTab[11998]"
}

// AddrFromSlice parses the 4- or 16-byte byte slice as an IPv4 or IPv6 address.
//line /usr/local/go/src/net/netip/netip.go:337
// Note that a net.IP can be passed directly as the []byte argument.
//line /usr/local/go/src/net/netip/netip.go:337
// If slice's length is not 4 or 16, AddrFromSlice returns Addr{}, false.
//line /usr/local/go/src/net/netip/netip.go:340
func AddrFromSlice(slice []byte) (ip Addr, ok bool) {
//line /usr/local/go/src/net/netip/netip.go:340
	_go_fuzz_dep_.CoverTab[12068]++
							switch len(slice) {
	case 4:
//line /usr/local/go/src/net/netip/netip.go:342
		_go_fuzz_dep_.CoverTab[12070]++
								return AddrFrom4([4]byte(slice)), true
//line /usr/local/go/src/net/netip/netip.go:343
		// _ = "end of CoverTab[12070]"
	case 16:
//line /usr/local/go/src/net/netip/netip.go:344
		_go_fuzz_dep_.CoverTab[12071]++
								return AddrFrom16([16]byte(slice)), true
//line /usr/local/go/src/net/netip/netip.go:345
		// _ = "end of CoverTab[12071]"
//line /usr/local/go/src/net/netip/netip.go:345
	default:
//line /usr/local/go/src/net/netip/netip.go:345
		_go_fuzz_dep_.CoverTab[12072]++
//line /usr/local/go/src/net/netip/netip.go:345
		// _ = "end of CoverTab[12072]"
	}
//line /usr/local/go/src/net/netip/netip.go:346
	// _ = "end of CoverTab[12068]"
//line /usr/local/go/src/net/netip/netip.go:346
	_go_fuzz_dep_.CoverTab[12069]++
							return Addr{}, false
//line /usr/local/go/src/net/netip/netip.go:347
	// _ = "end of CoverTab[12069]"
}

// v4 returns the i'th byte of ip. If ip is not an IPv4, v4 returns
//line /usr/local/go/src/net/netip/netip.go:350
// unspecified garbage.
//line /usr/local/go/src/net/netip/netip.go:352
func (ip Addr) v4(i uint8) uint8 {
//line /usr/local/go/src/net/netip/netip.go:352
	_go_fuzz_dep_.CoverTab[12073]++
							return uint8(ip.addr.lo >> ((3 - i) * 8))
//line /usr/local/go/src/net/netip/netip.go:353
	// _ = "end of CoverTab[12073]"
}

// v6 returns the i'th byte of ip. If ip is an IPv4 address, this
//line /usr/local/go/src/net/netip/netip.go:356
// accesses the IPv4-mapped IPv6 address form of the IP.
//line /usr/local/go/src/net/netip/netip.go:358
func (ip Addr) v6(i uint8) uint8 {
//line /usr/local/go/src/net/netip/netip.go:358
	_go_fuzz_dep_.CoverTab[12074]++
							return uint8(*(ip.addr.halves()[(i/8)%2]) >> ((7 - i%8) * 8))
//line /usr/local/go/src/net/netip/netip.go:359
	// _ = "end of CoverTab[12074]"
}

// v6u16 returns the i'th 16-bit word of ip. If ip is an IPv4 address,
//line /usr/local/go/src/net/netip/netip.go:362
// this accesses the IPv4-mapped IPv6 address form of the IP.
//line /usr/local/go/src/net/netip/netip.go:364
func (ip Addr) v6u16(i uint8) uint16 {
//line /usr/local/go/src/net/netip/netip.go:364
	_go_fuzz_dep_.CoverTab[12075]++
							return uint16(*(ip.addr.halves()[(i/4)%2]) >> ((3 - i%4) * 16))
//line /usr/local/go/src/net/netip/netip.go:365
	// _ = "end of CoverTab[12075]"
}

// isZero reports whether ip is the zero value of the IP type.
//line /usr/local/go/src/net/netip/netip.go:368
// The zero value is not a valid IP address of any type.
//line /usr/local/go/src/net/netip/netip.go:368
//
//line /usr/local/go/src/net/netip/netip.go:368
// Note that "0.0.0.0" and "::" are not the zero value. Use IsUnspecified to
//line /usr/local/go/src/net/netip/netip.go:368
// check for these values instead.
//line /usr/local/go/src/net/netip/netip.go:373
func (ip Addr) isZero() bool {
//line /usr/local/go/src/net/netip/netip.go:373
	_go_fuzz_dep_.CoverTab[12076]++

//line /usr/local/go/src/net/netip/netip.go:376
	return ip.z == z0
//line /usr/local/go/src/net/netip/netip.go:376
	// _ = "end of CoverTab[12076]"
}

// IsValid reports whether the Addr is an initialized address (not the zero Addr).
//line /usr/local/go/src/net/netip/netip.go:379
//
//line /usr/local/go/src/net/netip/netip.go:379
// Note that "0.0.0.0" and "::" are both valid values.
//line /usr/local/go/src/net/netip/netip.go:382
func (ip Addr) IsValid() bool {
//line /usr/local/go/src/net/netip/netip.go:382
	_go_fuzz_dep_.CoverTab[12077]++
//line /usr/local/go/src/net/netip/netip.go:382
	return ip.z != z0
//line /usr/local/go/src/net/netip/netip.go:382
	// _ = "end of CoverTab[12077]"
//line /usr/local/go/src/net/netip/netip.go:382
}

// BitLen returns the number of bits in the IP address:
//line /usr/local/go/src/net/netip/netip.go:384
// 128 for IPv6, 32 for IPv4, and 0 for the zero Addr.
//line /usr/local/go/src/net/netip/netip.go:384
//
//line /usr/local/go/src/net/netip/netip.go:384
// Note that IPv4-mapped IPv6 addresses are considered IPv6 addresses
//line /usr/local/go/src/net/netip/netip.go:384
// and therefore have bit length 128.
//line /usr/local/go/src/net/netip/netip.go:389
func (ip Addr) BitLen() int {
//line /usr/local/go/src/net/netip/netip.go:389
	_go_fuzz_dep_.CoverTab[12078]++
							switch ip.z {
	case z0:
//line /usr/local/go/src/net/netip/netip.go:391
		_go_fuzz_dep_.CoverTab[12080]++
								return 0
//line /usr/local/go/src/net/netip/netip.go:392
		// _ = "end of CoverTab[12080]"
	case z4:
//line /usr/local/go/src/net/netip/netip.go:393
		_go_fuzz_dep_.CoverTab[12081]++
								return 32
//line /usr/local/go/src/net/netip/netip.go:394
		// _ = "end of CoverTab[12081]"
//line /usr/local/go/src/net/netip/netip.go:394
	default:
//line /usr/local/go/src/net/netip/netip.go:394
		_go_fuzz_dep_.CoverTab[12082]++
//line /usr/local/go/src/net/netip/netip.go:394
		// _ = "end of CoverTab[12082]"
	}
//line /usr/local/go/src/net/netip/netip.go:395
	// _ = "end of CoverTab[12078]"
//line /usr/local/go/src/net/netip/netip.go:395
	_go_fuzz_dep_.CoverTab[12079]++
							return 128
//line /usr/local/go/src/net/netip/netip.go:396
	// _ = "end of CoverTab[12079]"
}

// Zone returns ip's IPv6 scoped addressing zone, if any.
func (ip Addr) Zone() string {
//line /usr/local/go/src/net/netip/netip.go:400
	_go_fuzz_dep_.CoverTab[12083]++
							if ip.z == nil {
//line /usr/local/go/src/net/netip/netip.go:401
		_go_fuzz_dep_.CoverTab[12085]++
								return ""
//line /usr/local/go/src/net/netip/netip.go:402
		// _ = "end of CoverTab[12085]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:403
		_go_fuzz_dep_.CoverTab[12086]++
//line /usr/local/go/src/net/netip/netip.go:403
		// _ = "end of CoverTab[12086]"
//line /usr/local/go/src/net/netip/netip.go:403
	}
//line /usr/local/go/src/net/netip/netip.go:403
	// _ = "end of CoverTab[12083]"
//line /usr/local/go/src/net/netip/netip.go:403
	_go_fuzz_dep_.CoverTab[12084]++
							zone, _ := ip.z.Get().(string)
							return zone
//line /usr/local/go/src/net/netip/netip.go:405
	// _ = "end of CoverTab[12084]"
}

// Compare returns an integer comparing two IPs.
//line /usr/local/go/src/net/netip/netip.go:408
// The result will be 0 if ip == ip2, -1 if ip < ip2, and +1 if ip > ip2.
//line /usr/local/go/src/net/netip/netip.go:408
// The definition of "less than" is the same as the Less method.
//line /usr/local/go/src/net/netip/netip.go:411
func (ip Addr) Compare(ip2 Addr) int {
//line /usr/local/go/src/net/netip/netip.go:411
	_go_fuzz_dep_.CoverTab[12087]++
							f1, f2 := ip.BitLen(), ip2.BitLen()
							if f1 < f2 {
//line /usr/local/go/src/net/netip/netip.go:413
		_go_fuzz_dep_.CoverTab[12095]++
								return -1
//line /usr/local/go/src/net/netip/netip.go:414
		// _ = "end of CoverTab[12095]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:415
		_go_fuzz_dep_.CoverTab[12096]++
//line /usr/local/go/src/net/netip/netip.go:415
		// _ = "end of CoverTab[12096]"
//line /usr/local/go/src/net/netip/netip.go:415
	}
//line /usr/local/go/src/net/netip/netip.go:415
	// _ = "end of CoverTab[12087]"
//line /usr/local/go/src/net/netip/netip.go:415
	_go_fuzz_dep_.CoverTab[12088]++
							if f1 > f2 {
//line /usr/local/go/src/net/netip/netip.go:416
		_go_fuzz_dep_.CoverTab[12097]++
								return 1
//line /usr/local/go/src/net/netip/netip.go:417
		// _ = "end of CoverTab[12097]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:418
		_go_fuzz_dep_.CoverTab[12098]++
//line /usr/local/go/src/net/netip/netip.go:418
		// _ = "end of CoverTab[12098]"
//line /usr/local/go/src/net/netip/netip.go:418
	}
//line /usr/local/go/src/net/netip/netip.go:418
	// _ = "end of CoverTab[12088]"
//line /usr/local/go/src/net/netip/netip.go:418
	_go_fuzz_dep_.CoverTab[12089]++
							hi1, hi2 := ip.addr.hi, ip2.addr.hi
							if hi1 < hi2 {
//line /usr/local/go/src/net/netip/netip.go:420
		_go_fuzz_dep_.CoverTab[12099]++
								return -1
//line /usr/local/go/src/net/netip/netip.go:421
		// _ = "end of CoverTab[12099]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:422
		_go_fuzz_dep_.CoverTab[12100]++
//line /usr/local/go/src/net/netip/netip.go:422
		// _ = "end of CoverTab[12100]"
//line /usr/local/go/src/net/netip/netip.go:422
	}
//line /usr/local/go/src/net/netip/netip.go:422
	// _ = "end of CoverTab[12089]"
//line /usr/local/go/src/net/netip/netip.go:422
	_go_fuzz_dep_.CoverTab[12090]++
							if hi1 > hi2 {
//line /usr/local/go/src/net/netip/netip.go:423
		_go_fuzz_dep_.CoverTab[12101]++
								return 1
//line /usr/local/go/src/net/netip/netip.go:424
		// _ = "end of CoverTab[12101]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:425
		_go_fuzz_dep_.CoverTab[12102]++
//line /usr/local/go/src/net/netip/netip.go:425
		// _ = "end of CoverTab[12102]"
//line /usr/local/go/src/net/netip/netip.go:425
	}
//line /usr/local/go/src/net/netip/netip.go:425
	// _ = "end of CoverTab[12090]"
//line /usr/local/go/src/net/netip/netip.go:425
	_go_fuzz_dep_.CoverTab[12091]++
							lo1, lo2 := ip.addr.lo, ip2.addr.lo
							if lo1 < lo2 {
//line /usr/local/go/src/net/netip/netip.go:427
		_go_fuzz_dep_.CoverTab[12103]++
								return -1
//line /usr/local/go/src/net/netip/netip.go:428
		// _ = "end of CoverTab[12103]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:429
		_go_fuzz_dep_.CoverTab[12104]++
//line /usr/local/go/src/net/netip/netip.go:429
		// _ = "end of CoverTab[12104]"
//line /usr/local/go/src/net/netip/netip.go:429
	}
//line /usr/local/go/src/net/netip/netip.go:429
	// _ = "end of CoverTab[12091]"
//line /usr/local/go/src/net/netip/netip.go:429
	_go_fuzz_dep_.CoverTab[12092]++
							if lo1 > lo2 {
//line /usr/local/go/src/net/netip/netip.go:430
		_go_fuzz_dep_.CoverTab[12105]++
								return 1
//line /usr/local/go/src/net/netip/netip.go:431
		// _ = "end of CoverTab[12105]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:432
		_go_fuzz_dep_.CoverTab[12106]++
//line /usr/local/go/src/net/netip/netip.go:432
		// _ = "end of CoverTab[12106]"
//line /usr/local/go/src/net/netip/netip.go:432
	}
//line /usr/local/go/src/net/netip/netip.go:432
	// _ = "end of CoverTab[12092]"
//line /usr/local/go/src/net/netip/netip.go:432
	_go_fuzz_dep_.CoverTab[12093]++
							if ip.Is6() {
//line /usr/local/go/src/net/netip/netip.go:433
		_go_fuzz_dep_.CoverTab[12107]++
								za, zb := ip.Zone(), ip2.Zone()
								if za < zb {
//line /usr/local/go/src/net/netip/netip.go:435
			_go_fuzz_dep_.CoverTab[12109]++
									return -1
//line /usr/local/go/src/net/netip/netip.go:436
			// _ = "end of CoverTab[12109]"
		} else {
//line /usr/local/go/src/net/netip/netip.go:437
			_go_fuzz_dep_.CoverTab[12110]++
//line /usr/local/go/src/net/netip/netip.go:437
			// _ = "end of CoverTab[12110]"
//line /usr/local/go/src/net/netip/netip.go:437
		}
//line /usr/local/go/src/net/netip/netip.go:437
		// _ = "end of CoverTab[12107]"
//line /usr/local/go/src/net/netip/netip.go:437
		_go_fuzz_dep_.CoverTab[12108]++
								if za > zb {
//line /usr/local/go/src/net/netip/netip.go:438
			_go_fuzz_dep_.CoverTab[12111]++
									return 1
//line /usr/local/go/src/net/netip/netip.go:439
			// _ = "end of CoverTab[12111]"
		} else {
//line /usr/local/go/src/net/netip/netip.go:440
			_go_fuzz_dep_.CoverTab[12112]++
//line /usr/local/go/src/net/netip/netip.go:440
			// _ = "end of CoverTab[12112]"
//line /usr/local/go/src/net/netip/netip.go:440
		}
//line /usr/local/go/src/net/netip/netip.go:440
		// _ = "end of CoverTab[12108]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:441
		_go_fuzz_dep_.CoverTab[12113]++
//line /usr/local/go/src/net/netip/netip.go:441
		// _ = "end of CoverTab[12113]"
//line /usr/local/go/src/net/netip/netip.go:441
	}
//line /usr/local/go/src/net/netip/netip.go:441
	// _ = "end of CoverTab[12093]"
//line /usr/local/go/src/net/netip/netip.go:441
	_go_fuzz_dep_.CoverTab[12094]++
							return 0
//line /usr/local/go/src/net/netip/netip.go:442
	// _ = "end of CoverTab[12094]"
}

// Less reports whether ip sorts before ip2.
//line /usr/local/go/src/net/netip/netip.go:445
// IP addresses sort first by length, then their address.
//line /usr/local/go/src/net/netip/netip.go:445
// IPv6 addresses with zones sort just after the same address without a zone.
//line /usr/local/go/src/net/netip/netip.go:448
func (ip Addr) Less(ip2 Addr) bool {
//line /usr/local/go/src/net/netip/netip.go:448
	_go_fuzz_dep_.CoverTab[12114]++
//line /usr/local/go/src/net/netip/netip.go:448
	return ip.Compare(ip2) == -1
//line /usr/local/go/src/net/netip/netip.go:448
	// _ = "end of CoverTab[12114]"
//line /usr/local/go/src/net/netip/netip.go:448
}

// Is4 reports whether ip is an IPv4 address.
//line /usr/local/go/src/net/netip/netip.go:450
//
//line /usr/local/go/src/net/netip/netip.go:450
// It returns false for IPv4-mapped IPv6 addresses. See Addr.Unmap.
//line /usr/local/go/src/net/netip/netip.go:453
func (ip Addr) Is4() bool {
//line /usr/local/go/src/net/netip/netip.go:453
	_go_fuzz_dep_.CoverTab[12115]++
							return ip.z == z4
//line /usr/local/go/src/net/netip/netip.go:454
	// _ = "end of CoverTab[12115]"
}

// Is4In6 reports whether ip is an IPv4-mapped IPv6 address.
func (ip Addr) Is4In6() bool {
//line /usr/local/go/src/net/netip/netip.go:458
	_go_fuzz_dep_.CoverTab[12116]++
							return ip.Is6() && func() bool {
//line /usr/local/go/src/net/netip/netip.go:459
		_go_fuzz_dep_.CoverTab[12117]++
//line /usr/local/go/src/net/netip/netip.go:459
		return ip.addr.hi == 0
//line /usr/local/go/src/net/netip/netip.go:459
		// _ = "end of CoverTab[12117]"
//line /usr/local/go/src/net/netip/netip.go:459
	}() && func() bool {
//line /usr/local/go/src/net/netip/netip.go:459
		_go_fuzz_dep_.CoverTab[12118]++
//line /usr/local/go/src/net/netip/netip.go:459
		return ip.addr.lo>>32 == 0xffff
//line /usr/local/go/src/net/netip/netip.go:459
		// _ = "end of CoverTab[12118]"
//line /usr/local/go/src/net/netip/netip.go:459
	}()
//line /usr/local/go/src/net/netip/netip.go:459
	// _ = "end of CoverTab[12116]"
}

// Is6 reports whether ip is an IPv6 address, including IPv4-mapped
//line /usr/local/go/src/net/netip/netip.go:462
// IPv6 addresses.
//line /usr/local/go/src/net/netip/netip.go:464
func (ip Addr) Is6() bool {
//line /usr/local/go/src/net/netip/netip.go:464
	_go_fuzz_dep_.CoverTab[12119]++
							return ip.z != z0 && func() bool {
//line /usr/local/go/src/net/netip/netip.go:465
		_go_fuzz_dep_.CoverTab[12120]++
//line /usr/local/go/src/net/netip/netip.go:465
		return ip.z != z4
//line /usr/local/go/src/net/netip/netip.go:465
		// _ = "end of CoverTab[12120]"
//line /usr/local/go/src/net/netip/netip.go:465
	}()
//line /usr/local/go/src/net/netip/netip.go:465
	// _ = "end of CoverTab[12119]"
}

// Unmap returns ip with any IPv4-mapped IPv6 address prefix removed.
//line /usr/local/go/src/net/netip/netip.go:468
//
//line /usr/local/go/src/net/netip/netip.go:468
// That is, if ip is an IPv6 address wrapping an IPv4 address, it
//line /usr/local/go/src/net/netip/netip.go:468
// returns the wrapped IPv4 address. Otherwise it returns ip unmodified.
//line /usr/local/go/src/net/netip/netip.go:472
func (ip Addr) Unmap() Addr {
//line /usr/local/go/src/net/netip/netip.go:472
	_go_fuzz_dep_.CoverTab[12121]++
							if ip.Is4In6() {
//line /usr/local/go/src/net/netip/netip.go:473
		_go_fuzz_dep_.CoverTab[12123]++
								ip.z = z4
//line /usr/local/go/src/net/netip/netip.go:474
		// _ = "end of CoverTab[12123]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:475
		_go_fuzz_dep_.CoverTab[12124]++
//line /usr/local/go/src/net/netip/netip.go:475
		// _ = "end of CoverTab[12124]"
//line /usr/local/go/src/net/netip/netip.go:475
	}
//line /usr/local/go/src/net/netip/netip.go:475
	// _ = "end of CoverTab[12121]"
//line /usr/local/go/src/net/netip/netip.go:475
	_go_fuzz_dep_.CoverTab[12122]++
							return ip
//line /usr/local/go/src/net/netip/netip.go:476
	// _ = "end of CoverTab[12122]"
}

// WithZone returns an IP that's the same as ip but with the provided
//line /usr/local/go/src/net/netip/netip.go:479
// zone. If zone is empty, the zone is removed. If ip is an IPv4
//line /usr/local/go/src/net/netip/netip.go:479
// address, WithZone is a no-op and returns ip unchanged.
//line /usr/local/go/src/net/netip/netip.go:482
func (ip Addr) WithZone(zone string) Addr {
//line /usr/local/go/src/net/netip/netip.go:482
	_go_fuzz_dep_.CoverTab[12125]++
							if !ip.Is6() {
//line /usr/local/go/src/net/netip/netip.go:483
		_go_fuzz_dep_.CoverTab[12128]++
								return ip
//line /usr/local/go/src/net/netip/netip.go:484
		// _ = "end of CoverTab[12128]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:485
		_go_fuzz_dep_.CoverTab[12129]++
//line /usr/local/go/src/net/netip/netip.go:485
		// _ = "end of CoverTab[12129]"
//line /usr/local/go/src/net/netip/netip.go:485
	}
//line /usr/local/go/src/net/netip/netip.go:485
	// _ = "end of CoverTab[12125]"
//line /usr/local/go/src/net/netip/netip.go:485
	_go_fuzz_dep_.CoverTab[12126]++
							if zone == "" {
//line /usr/local/go/src/net/netip/netip.go:486
		_go_fuzz_dep_.CoverTab[12130]++
								ip.z = z6noz
								return ip
//line /usr/local/go/src/net/netip/netip.go:488
		// _ = "end of CoverTab[12130]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:489
		_go_fuzz_dep_.CoverTab[12131]++
//line /usr/local/go/src/net/netip/netip.go:489
		// _ = "end of CoverTab[12131]"
//line /usr/local/go/src/net/netip/netip.go:489
	}
//line /usr/local/go/src/net/netip/netip.go:489
	// _ = "end of CoverTab[12126]"
//line /usr/local/go/src/net/netip/netip.go:489
	_go_fuzz_dep_.CoverTab[12127]++
							ip.z = intern.GetByString(zone)
							return ip
//line /usr/local/go/src/net/netip/netip.go:491
	// _ = "end of CoverTab[12127]"
}

// withoutZone unconditionally strips the zone from ip.
//line /usr/local/go/src/net/netip/netip.go:494
// It's similar to WithZone, but small enough to be inlinable.
//line /usr/local/go/src/net/netip/netip.go:496
func (ip Addr) withoutZone() Addr {
//line /usr/local/go/src/net/netip/netip.go:496
	_go_fuzz_dep_.CoverTab[12132]++
							if !ip.Is6() {
//line /usr/local/go/src/net/netip/netip.go:497
		_go_fuzz_dep_.CoverTab[12134]++
								return ip
//line /usr/local/go/src/net/netip/netip.go:498
		// _ = "end of CoverTab[12134]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:499
		_go_fuzz_dep_.CoverTab[12135]++
//line /usr/local/go/src/net/netip/netip.go:499
		// _ = "end of CoverTab[12135]"
//line /usr/local/go/src/net/netip/netip.go:499
	}
//line /usr/local/go/src/net/netip/netip.go:499
	// _ = "end of CoverTab[12132]"
//line /usr/local/go/src/net/netip/netip.go:499
	_go_fuzz_dep_.CoverTab[12133]++
							ip.z = z6noz
							return ip
//line /usr/local/go/src/net/netip/netip.go:501
	// _ = "end of CoverTab[12133]"
}

// hasZone reports whether ip has an IPv6 zone.
func (ip Addr) hasZone() bool {
//line /usr/local/go/src/net/netip/netip.go:505
	_go_fuzz_dep_.CoverTab[12136]++
							return ip.z != z0 && func() bool {
//line /usr/local/go/src/net/netip/netip.go:506
		_go_fuzz_dep_.CoverTab[12137]++
//line /usr/local/go/src/net/netip/netip.go:506
		return ip.z != z4
//line /usr/local/go/src/net/netip/netip.go:506
		// _ = "end of CoverTab[12137]"
//line /usr/local/go/src/net/netip/netip.go:506
	}() && func() bool {
//line /usr/local/go/src/net/netip/netip.go:506
		_go_fuzz_dep_.CoverTab[12138]++
//line /usr/local/go/src/net/netip/netip.go:506
		return ip.z != z6noz
//line /usr/local/go/src/net/netip/netip.go:506
		// _ = "end of CoverTab[12138]"
//line /usr/local/go/src/net/netip/netip.go:506
	}()
//line /usr/local/go/src/net/netip/netip.go:506
	// _ = "end of CoverTab[12136]"
}

// IsLinkLocalUnicast reports whether ip is a link-local unicast address.
func (ip Addr) IsLinkLocalUnicast() bool {
//line /usr/local/go/src/net/netip/netip.go:510
	_go_fuzz_dep_.CoverTab[12139]++

//line /usr/local/go/src/net/netip/netip.go:513
	if ip.Is4() {
//line /usr/local/go/src/net/netip/netip.go:513
		_go_fuzz_dep_.CoverTab[12142]++
								return ip.v4(0) == 169 && func() bool {
//line /usr/local/go/src/net/netip/netip.go:514
			_go_fuzz_dep_.CoverTab[12143]++
//line /usr/local/go/src/net/netip/netip.go:514
			return ip.v4(1) == 254
//line /usr/local/go/src/net/netip/netip.go:514
			// _ = "end of CoverTab[12143]"
//line /usr/local/go/src/net/netip/netip.go:514
		}()
//line /usr/local/go/src/net/netip/netip.go:514
		// _ = "end of CoverTab[12142]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:515
		_go_fuzz_dep_.CoverTab[12144]++
//line /usr/local/go/src/net/netip/netip.go:515
		// _ = "end of CoverTab[12144]"
//line /usr/local/go/src/net/netip/netip.go:515
	}
//line /usr/local/go/src/net/netip/netip.go:515
	// _ = "end of CoverTab[12139]"
//line /usr/local/go/src/net/netip/netip.go:515
	_go_fuzz_dep_.CoverTab[12140]++

//line /usr/local/go/src/net/netip/netip.go:518
	if ip.Is6() {
//line /usr/local/go/src/net/netip/netip.go:518
		_go_fuzz_dep_.CoverTab[12145]++
								return ip.v6u16(0)&0xffc0 == 0xfe80
//line /usr/local/go/src/net/netip/netip.go:519
		// _ = "end of CoverTab[12145]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:520
		_go_fuzz_dep_.CoverTab[12146]++
//line /usr/local/go/src/net/netip/netip.go:520
		// _ = "end of CoverTab[12146]"
//line /usr/local/go/src/net/netip/netip.go:520
	}
//line /usr/local/go/src/net/netip/netip.go:520
	// _ = "end of CoverTab[12140]"
//line /usr/local/go/src/net/netip/netip.go:520
	_go_fuzz_dep_.CoverTab[12141]++
							return false
//line /usr/local/go/src/net/netip/netip.go:521
	// _ = "end of CoverTab[12141]"
}

// IsLoopback reports whether ip is a loopback address.
func (ip Addr) IsLoopback() bool {
//line /usr/local/go/src/net/netip/netip.go:525
	_go_fuzz_dep_.CoverTab[12147]++

//line /usr/local/go/src/net/netip/netip.go:528
	if ip.Is4() {
//line /usr/local/go/src/net/netip/netip.go:528
		_go_fuzz_dep_.CoverTab[12150]++
								return ip.v4(0) == 127
//line /usr/local/go/src/net/netip/netip.go:529
		// _ = "end of CoverTab[12150]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:530
		_go_fuzz_dep_.CoverTab[12151]++
//line /usr/local/go/src/net/netip/netip.go:530
		// _ = "end of CoverTab[12151]"
//line /usr/local/go/src/net/netip/netip.go:530
	}
//line /usr/local/go/src/net/netip/netip.go:530
	// _ = "end of CoverTab[12147]"
//line /usr/local/go/src/net/netip/netip.go:530
	_go_fuzz_dep_.CoverTab[12148]++

//line /usr/local/go/src/net/netip/netip.go:533
	if ip.Is6() {
//line /usr/local/go/src/net/netip/netip.go:533
		_go_fuzz_dep_.CoverTab[12152]++
								return ip.addr.hi == 0 && func() bool {
//line /usr/local/go/src/net/netip/netip.go:534
			_go_fuzz_dep_.CoverTab[12153]++
//line /usr/local/go/src/net/netip/netip.go:534
			return ip.addr.lo == 1
//line /usr/local/go/src/net/netip/netip.go:534
			// _ = "end of CoverTab[12153]"
//line /usr/local/go/src/net/netip/netip.go:534
		}()
//line /usr/local/go/src/net/netip/netip.go:534
		// _ = "end of CoverTab[12152]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:535
		_go_fuzz_dep_.CoverTab[12154]++
//line /usr/local/go/src/net/netip/netip.go:535
		// _ = "end of CoverTab[12154]"
//line /usr/local/go/src/net/netip/netip.go:535
	}
//line /usr/local/go/src/net/netip/netip.go:535
	// _ = "end of CoverTab[12148]"
//line /usr/local/go/src/net/netip/netip.go:535
	_go_fuzz_dep_.CoverTab[12149]++
							return false
//line /usr/local/go/src/net/netip/netip.go:536
	// _ = "end of CoverTab[12149]"
}

// IsMulticast reports whether ip is a multicast address.
func (ip Addr) IsMulticast() bool {
//line /usr/local/go/src/net/netip/netip.go:540
	_go_fuzz_dep_.CoverTab[12155]++

//line /usr/local/go/src/net/netip/netip.go:543
	if ip.Is4() {
//line /usr/local/go/src/net/netip/netip.go:543
		_go_fuzz_dep_.CoverTab[12158]++
								return ip.v4(0)&0xf0 == 0xe0
//line /usr/local/go/src/net/netip/netip.go:544
		// _ = "end of CoverTab[12158]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:545
		_go_fuzz_dep_.CoverTab[12159]++
//line /usr/local/go/src/net/netip/netip.go:545
		// _ = "end of CoverTab[12159]"
//line /usr/local/go/src/net/netip/netip.go:545
	}
//line /usr/local/go/src/net/netip/netip.go:545
	// _ = "end of CoverTab[12155]"
//line /usr/local/go/src/net/netip/netip.go:545
	_go_fuzz_dep_.CoverTab[12156]++

//line /usr/local/go/src/net/netip/netip.go:548
	if ip.Is6() {
//line /usr/local/go/src/net/netip/netip.go:548
		_go_fuzz_dep_.CoverTab[12160]++
								return ip.addr.hi>>(64-8) == 0xff
//line /usr/local/go/src/net/netip/netip.go:549
		// _ = "end of CoverTab[12160]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:550
		_go_fuzz_dep_.CoverTab[12161]++
//line /usr/local/go/src/net/netip/netip.go:550
		// _ = "end of CoverTab[12161]"
//line /usr/local/go/src/net/netip/netip.go:550
	}
//line /usr/local/go/src/net/netip/netip.go:550
	// _ = "end of CoverTab[12156]"
//line /usr/local/go/src/net/netip/netip.go:550
	_go_fuzz_dep_.CoverTab[12157]++
							return false
//line /usr/local/go/src/net/netip/netip.go:551
	// _ = "end of CoverTab[12157]"
}

// IsInterfaceLocalMulticast reports whether ip is an IPv6 interface-local
//line /usr/local/go/src/net/netip/netip.go:554
// multicast address.
//line /usr/local/go/src/net/netip/netip.go:556
func (ip Addr) IsInterfaceLocalMulticast() bool {
//line /usr/local/go/src/net/netip/netip.go:556
	_go_fuzz_dep_.CoverTab[12162]++

//line /usr/local/go/src/net/netip/netip.go:559
	if ip.Is6() {
//line /usr/local/go/src/net/netip/netip.go:559
		_go_fuzz_dep_.CoverTab[12164]++
								return ip.v6u16(0)&0xff0f == 0xff01
//line /usr/local/go/src/net/netip/netip.go:560
		// _ = "end of CoverTab[12164]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:561
		_go_fuzz_dep_.CoverTab[12165]++
//line /usr/local/go/src/net/netip/netip.go:561
		// _ = "end of CoverTab[12165]"
//line /usr/local/go/src/net/netip/netip.go:561
	}
//line /usr/local/go/src/net/netip/netip.go:561
	// _ = "end of CoverTab[12162]"
//line /usr/local/go/src/net/netip/netip.go:561
	_go_fuzz_dep_.CoverTab[12163]++
							return false
//line /usr/local/go/src/net/netip/netip.go:562
	// _ = "end of CoverTab[12163]"
}

// IsLinkLocalMulticast reports whether ip is a link-local multicast address.
func (ip Addr) IsLinkLocalMulticast() bool {
//line /usr/local/go/src/net/netip/netip.go:566
	_go_fuzz_dep_.CoverTab[12166]++

//line /usr/local/go/src/net/netip/netip.go:569
	if ip.Is4() {
//line /usr/local/go/src/net/netip/netip.go:569
		_go_fuzz_dep_.CoverTab[12169]++
								return ip.v4(0) == 224 && func() bool {
//line /usr/local/go/src/net/netip/netip.go:570
			_go_fuzz_dep_.CoverTab[12170]++
//line /usr/local/go/src/net/netip/netip.go:570
			return ip.v4(1) == 0
//line /usr/local/go/src/net/netip/netip.go:570
			// _ = "end of CoverTab[12170]"
//line /usr/local/go/src/net/netip/netip.go:570
		}() && func() bool {
//line /usr/local/go/src/net/netip/netip.go:570
			_go_fuzz_dep_.CoverTab[12171]++
//line /usr/local/go/src/net/netip/netip.go:570
			return ip.v4(2) == 0
//line /usr/local/go/src/net/netip/netip.go:570
			// _ = "end of CoverTab[12171]"
//line /usr/local/go/src/net/netip/netip.go:570
		}()
//line /usr/local/go/src/net/netip/netip.go:570
		// _ = "end of CoverTab[12169]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:571
		_go_fuzz_dep_.CoverTab[12172]++
//line /usr/local/go/src/net/netip/netip.go:571
		// _ = "end of CoverTab[12172]"
//line /usr/local/go/src/net/netip/netip.go:571
	}
//line /usr/local/go/src/net/netip/netip.go:571
	// _ = "end of CoverTab[12166]"
//line /usr/local/go/src/net/netip/netip.go:571
	_go_fuzz_dep_.CoverTab[12167]++

//line /usr/local/go/src/net/netip/netip.go:574
	if ip.Is6() {
//line /usr/local/go/src/net/netip/netip.go:574
		_go_fuzz_dep_.CoverTab[12173]++
								return ip.v6u16(0)&0xff0f == 0xff02
//line /usr/local/go/src/net/netip/netip.go:575
		// _ = "end of CoverTab[12173]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:576
		_go_fuzz_dep_.CoverTab[12174]++
//line /usr/local/go/src/net/netip/netip.go:576
		// _ = "end of CoverTab[12174]"
//line /usr/local/go/src/net/netip/netip.go:576
	}
//line /usr/local/go/src/net/netip/netip.go:576
	// _ = "end of CoverTab[12167]"
//line /usr/local/go/src/net/netip/netip.go:576
	_go_fuzz_dep_.CoverTab[12168]++
							return false
//line /usr/local/go/src/net/netip/netip.go:577
	// _ = "end of CoverTab[12168]"
}

// IsGlobalUnicast reports whether ip is a global unicast address.
//line /usr/local/go/src/net/netip/netip.go:580
//
//line /usr/local/go/src/net/netip/netip.go:580
// It returns true for IPv6 addresses which fall outside of the current
//line /usr/local/go/src/net/netip/netip.go:580
// IANA-allocated 2000::/3 global unicast space, with the exception of the
//line /usr/local/go/src/net/netip/netip.go:580
// link-local address space. It also returns true even if ip is in the IPv4
//line /usr/local/go/src/net/netip/netip.go:580
// private address space or IPv6 unique local address space.
//line /usr/local/go/src/net/netip/netip.go:580
// It returns false for the zero Addr.
//line /usr/local/go/src/net/netip/netip.go:580
//
//line /usr/local/go/src/net/netip/netip.go:580
// For reference, see RFC 1122, RFC 4291, and RFC 4632.
//line /usr/local/go/src/net/netip/netip.go:589
func (ip Addr) IsGlobalUnicast() bool {
//line /usr/local/go/src/net/netip/netip.go:589
	_go_fuzz_dep_.CoverTab[12175]++
							if ip.z == z0 {
//line /usr/local/go/src/net/netip/netip.go:590
		_go_fuzz_dep_.CoverTab[12178]++

								return false
//line /usr/local/go/src/net/netip/netip.go:592
		// _ = "end of CoverTab[12178]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:593
		_go_fuzz_dep_.CoverTab[12179]++
//line /usr/local/go/src/net/netip/netip.go:593
		// _ = "end of CoverTab[12179]"
//line /usr/local/go/src/net/netip/netip.go:593
	}
//line /usr/local/go/src/net/netip/netip.go:593
	// _ = "end of CoverTab[12175]"
//line /usr/local/go/src/net/netip/netip.go:593
	_go_fuzz_dep_.CoverTab[12176]++

//line /usr/local/go/src/net/netip/netip.go:597
	if ip.Is4() && func() bool {
//line /usr/local/go/src/net/netip/netip.go:597
		_go_fuzz_dep_.CoverTab[12180]++
//line /usr/local/go/src/net/netip/netip.go:597
		return (ip == IPv4Unspecified() || func() bool {
//line /usr/local/go/src/net/netip/netip.go:597
			_go_fuzz_dep_.CoverTab[12181]++
//line /usr/local/go/src/net/netip/netip.go:597
			return ip == AddrFrom4([4]byte{255, 255, 255, 255})
//line /usr/local/go/src/net/netip/netip.go:597
			// _ = "end of CoverTab[12181]"
//line /usr/local/go/src/net/netip/netip.go:597
		}())
//line /usr/local/go/src/net/netip/netip.go:597
		// _ = "end of CoverTab[12180]"
//line /usr/local/go/src/net/netip/netip.go:597
	}() {
//line /usr/local/go/src/net/netip/netip.go:597
		_go_fuzz_dep_.CoverTab[12182]++
								return false
//line /usr/local/go/src/net/netip/netip.go:598
		// _ = "end of CoverTab[12182]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:599
		_go_fuzz_dep_.CoverTab[12183]++
//line /usr/local/go/src/net/netip/netip.go:599
		// _ = "end of CoverTab[12183]"
//line /usr/local/go/src/net/netip/netip.go:599
	}
//line /usr/local/go/src/net/netip/netip.go:599
	// _ = "end of CoverTab[12176]"
//line /usr/local/go/src/net/netip/netip.go:599
	_go_fuzz_dep_.CoverTab[12177]++

							return ip != IPv6Unspecified() && func() bool {
//line /usr/local/go/src/net/netip/netip.go:601
		_go_fuzz_dep_.CoverTab[12184]++
//line /usr/local/go/src/net/netip/netip.go:601
		return !ip.IsLoopback()
								// _ = "end of CoverTab[12184]"
//line /usr/local/go/src/net/netip/netip.go:602
	}() && func() bool {
//line /usr/local/go/src/net/netip/netip.go:602
		_go_fuzz_dep_.CoverTab[12185]++
//line /usr/local/go/src/net/netip/netip.go:602
		return !ip.IsMulticast()
								// _ = "end of CoverTab[12185]"
//line /usr/local/go/src/net/netip/netip.go:603
	}() && func() bool {
//line /usr/local/go/src/net/netip/netip.go:603
		_go_fuzz_dep_.CoverTab[12186]++
//line /usr/local/go/src/net/netip/netip.go:603
		return !ip.IsLinkLocalUnicast()
								// _ = "end of CoverTab[12186]"
//line /usr/local/go/src/net/netip/netip.go:604
	}()
//line /usr/local/go/src/net/netip/netip.go:604
	// _ = "end of CoverTab[12177]"
}

// IsPrivate reports whether ip is a private address, according to RFC 1918
//line /usr/local/go/src/net/netip/netip.go:607
// (IPv4 addresses) and RFC 4193 (IPv6 addresses). That is, it reports whether
//line /usr/local/go/src/net/netip/netip.go:607
// ip is in 10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16, or fc00::/7. This is the
//line /usr/local/go/src/net/netip/netip.go:607
// same as net.IP.IsPrivate.
//line /usr/local/go/src/net/netip/netip.go:611
func (ip Addr) IsPrivate() bool {
//line /usr/local/go/src/net/netip/netip.go:611
	_go_fuzz_dep_.CoverTab[12187]++

							if ip.Is4() {
//line /usr/local/go/src/net/netip/netip.go:613
		_go_fuzz_dep_.CoverTab[12190]++

//line /usr/local/go/src/net/netip/netip.go:616
		return ip.v4(0) == 10 || func() bool {
//line /usr/local/go/src/net/netip/netip.go:616
			_go_fuzz_dep_.CoverTab[12191]++
//line /usr/local/go/src/net/netip/netip.go:616
			return (ip.v4(0) == 172 && func() bool {
										_go_fuzz_dep_.CoverTab[12192]++
//line /usr/local/go/src/net/netip/netip.go:617
				return ip.v4(1)&0xf0 == 16
//line /usr/local/go/src/net/netip/netip.go:617
				// _ = "end of CoverTab[12192]"
//line /usr/local/go/src/net/netip/netip.go:617
			}())
//line /usr/local/go/src/net/netip/netip.go:617
			// _ = "end of CoverTab[12191]"
//line /usr/local/go/src/net/netip/netip.go:617
		}() || func() bool {
//line /usr/local/go/src/net/netip/netip.go:617
			_go_fuzz_dep_.CoverTab[12193]++
//line /usr/local/go/src/net/netip/netip.go:617
			return (ip.v4(0) == 192 && func() bool {
										_go_fuzz_dep_.CoverTab[12194]++
//line /usr/local/go/src/net/netip/netip.go:618
				return ip.v4(1) == 168
//line /usr/local/go/src/net/netip/netip.go:618
				// _ = "end of CoverTab[12194]"
//line /usr/local/go/src/net/netip/netip.go:618
			}())
//line /usr/local/go/src/net/netip/netip.go:618
			// _ = "end of CoverTab[12193]"
//line /usr/local/go/src/net/netip/netip.go:618
		}()
//line /usr/local/go/src/net/netip/netip.go:618
		// _ = "end of CoverTab[12190]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:619
		_go_fuzz_dep_.CoverTab[12195]++
//line /usr/local/go/src/net/netip/netip.go:619
		// _ = "end of CoverTab[12195]"
//line /usr/local/go/src/net/netip/netip.go:619
	}
//line /usr/local/go/src/net/netip/netip.go:619
	// _ = "end of CoverTab[12187]"
//line /usr/local/go/src/net/netip/netip.go:619
	_go_fuzz_dep_.CoverTab[12188]++

							if ip.Is6() {
//line /usr/local/go/src/net/netip/netip.go:621
		_go_fuzz_dep_.CoverTab[12196]++

//line /usr/local/go/src/net/netip/netip.go:624
		return ip.v6(0)&0xfe == 0xfc
//line /usr/local/go/src/net/netip/netip.go:624
		// _ = "end of CoverTab[12196]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:625
		_go_fuzz_dep_.CoverTab[12197]++
//line /usr/local/go/src/net/netip/netip.go:625
		// _ = "end of CoverTab[12197]"
//line /usr/local/go/src/net/netip/netip.go:625
	}
//line /usr/local/go/src/net/netip/netip.go:625
	// _ = "end of CoverTab[12188]"
//line /usr/local/go/src/net/netip/netip.go:625
	_go_fuzz_dep_.CoverTab[12189]++

							return false
//line /usr/local/go/src/net/netip/netip.go:627
	// _ = "end of CoverTab[12189]"
}

// IsUnspecified reports whether ip is an unspecified address, either the IPv4
//line /usr/local/go/src/net/netip/netip.go:630
// address "0.0.0.0" or the IPv6 address "::".
//line /usr/local/go/src/net/netip/netip.go:630
//
//line /usr/local/go/src/net/netip/netip.go:630
// Note that the zero Addr is not an unspecified address.
//line /usr/local/go/src/net/netip/netip.go:634
func (ip Addr) IsUnspecified() bool {
//line /usr/local/go/src/net/netip/netip.go:634
	_go_fuzz_dep_.CoverTab[12198]++
							return ip == IPv4Unspecified() || func() bool {
//line /usr/local/go/src/net/netip/netip.go:635
		_go_fuzz_dep_.CoverTab[12199]++
//line /usr/local/go/src/net/netip/netip.go:635
		return ip == IPv6Unspecified()
//line /usr/local/go/src/net/netip/netip.go:635
		// _ = "end of CoverTab[12199]"
//line /usr/local/go/src/net/netip/netip.go:635
	}()
//line /usr/local/go/src/net/netip/netip.go:635
	// _ = "end of CoverTab[12198]"
}

// Prefix keeps only the top b bits of IP, producing a Prefix
//line /usr/local/go/src/net/netip/netip.go:638
// of the specified length.
//line /usr/local/go/src/net/netip/netip.go:638
// If ip is a zero Addr, Prefix always returns a zero Prefix and a nil error.
//line /usr/local/go/src/net/netip/netip.go:638
// Otherwise, if bits is less than zero or greater than ip.BitLen(),
//line /usr/local/go/src/net/netip/netip.go:638
// Prefix returns an error.
//line /usr/local/go/src/net/netip/netip.go:643
func (ip Addr) Prefix(b int) (Prefix, error) {
//line /usr/local/go/src/net/netip/netip.go:643
	_go_fuzz_dep_.CoverTab[12200]++
							if b < 0 {
//line /usr/local/go/src/net/netip/netip.go:644
		_go_fuzz_dep_.CoverTab[12203]++
								return Prefix{}, errors.New("negative Prefix bits")
//line /usr/local/go/src/net/netip/netip.go:645
		// _ = "end of CoverTab[12203]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:646
		_go_fuzz_dep_.CoverTab[12204]++
//line /usr/local/go/src/net/netip/netip.go:646
		// _ = "end of CoverTab[12204]"
//line /usr/local/go/src/net/netip/netip.go:646
	}
//line /usr/local/go/src/net/netip/netip.go:646
	// _ = "end of CoverTab[12200]"
//line /usr/local/go/src/net/netip/netip.go:646
	_go_fuzz_dep_.CoverTab[12201]++
							effectiveBits := b
							switch ip.z {
	case z0:
//line /usr/local/go/src/net/netip/netip.go:649
		_go_fuzz_dep_.CoverTab[12205]++
								return Prefix{}, nil
//line /usr/local/go/src/net/netip/netip.go:650
		// _ = "end of CoverTab[12205]"
	case z4:
//line /usr/local/go/src/net/netip/netip.go:651
		_go_fuzz_dep_.CoverTab[12206]++
								if b > 32 {
//line /usr/local/go/src/net/netip/netip.go:652
			_go_fuzz_dep_.CoverTab[12209]++
									return Prefix{}, errors.New("prefix length " + itoa.Itoa(b) + " too large for IPv4")
//line /usr/local/go/src/net/netip/netip.go:653
			// _ = "end of CoverTab[12209]"
		} else {
//line /usr/local/go/src/net/netip/netip.go:654
			_go_fuzz_dep_.CoverTab[12210]++
//line /usr/local/go/src/net/netip/netip.go:654
			// _ = "end of CoverTab[12210]"
//line /usr/local/go/src/net/netip/netip.go:654
		}
//line /usr/local/go/src/net/netip/netip.go:654
		// _ = "end of CoverTab[12206]"
//line /usr/local/go/src/net/netip/netip.go:654
		_go_fuzz_dep_.CoverTab[12207]++
								effectiveBits += 96
//line /usr/local/go/src/net/netip/netip.go:655
		// _ = "end of CoverTab[12207]"
	default:
//line /usr/local/go/src/net/netip/netip.go:656
		_go_fuzz_dep_.CoverTab[12208]++
								if b > 128 {
//line /usr/local/go/src/net/netip/netip.go:657
			_go_fuzz_dep_.CoverTab[12211]++
									return Prefix{}, errors.New("prefix length " + itoa.Itoa(b) + " too large for IPv6")
//line /usr/local/go/src/net/netip/netip.go:658
			// _ = "end of CoverTab[12211]"
		} else {
//line /usr/local/go/src/net/netip/netip.go:659
			_go_fuzz_dep_.CoverTab[12212]++
//line /usr/local/go/src/net/netip/netip.go:659
			// _ = "end of CoverTab[12212]"
//line /usr/local/go/src/net/netip/netip.go:659
		}
//line /usr/local/go/src/net/netip/netip.go:659
		// _ = "end of CoverTab[12208]"
	}
//line /usr/local/go/src/net/netip/netip.go:660
	// _ = "end of CoverTab[12201]"
//line /usr/local/go/src/net/netip/netip.go:660
	_go_fuzz_dep_.CoverTab[12202]++
							ip.addr = ip.addr.and(mask6(effectiveBits))
							return PrefixFrom(ip, b), nil
//line /usr/local/go/src/net/netip/netip.go:662
	// _ = "end of CoverTab[12202]"
}

const (
	netIPv4len	= 4
	netIPv6len	= 16
)

// As16 returns the IP address in its 16-byte representation.
//line /usr/local/go/src/net/netip/netip.go:670
// IPv4 addresses are returned as IPv4-mapped IPv6 addresses.
//line /usr/local/go/src/net/netip/netip.go:670
// IPv6 addresses with zones are returned without their zone (use the
//line /usr/local/go/src/net/netip/netip.go:670
// Zone method to get it).
//line /usr/local/go/src/net/netip/netip.go:670
// The ip zero value returns all zeroes.
//line /usr/local/go/src/net/netip/netip.go:675
func (ip Addr) As16() (a16 [16]byte) {
//line /usr/local/go/src/net/netip/netip.go:675
	_go_fuzz_dep_.CoverTab[12213]++
							bePutUint64(a16[:8], ip.addr.hi)
							bePutUint64(a16[8:], ip.addr.lo)
							return a16
//line /usr/local/go/src/net/netip/netip.go:678
	// _ = "end of CoverTab[12213]"
}

// As4 returns an IPv4 or IPv4-in-IPv6 address in its 4-byte representation.
//line /usr/local/go/src/net/netip/netip.go:681
// If ip is the zero Addr or an IPv6 address, As4 panics.
//line /usr/local/go/src/net/netip/netip.go:681
// Note that 0.0.0.0 is not the zero Addr.
//line /usr/local/go/src/net/netip/netip.go:684
func (ip Addr) As4() (a4 [4]byte) {
//line /usr/local/go/src/net/netip/netip.go:684
	_go_fuzz_dep_.CoverTab[12214]++
							if ip.z == z4 || func() bool {
//line /usr/local/go/src/net/netip/netip.go:685
		_go_fuzz_dep_.CoverTab[12217]++
//line /usr/local/go/src/net/netip/netip.go:685
		return ip.Is4In6()
//line /usr/local/go/src/net/netip/netip.go:685
		// _ = "end of CoverTab[12217]"
//line /usr/local/go/src/net/netip/netip.go:685
	}() {
//line /usr/local/go/src/net/netip/netip.go:685
		_go_fuzz_dep_.CoverTab[12218]++
								bePutUint32(a4[:], uint32(ip.addr.lo))
								return a4
//line /usr/local/go/src/net/netip/netip.go:687
		// _ = "end of CoverTab[12218]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:688
		_go_fuzz_dep_.CoverTab[12219]++
//line /usr/local/go/src/net/netip/netip.go:688
		// _ = "end of CoverTab[12219]"
//line /usr/local/go/src/net/netip/netip.go:688
	}
//line /usr/local/go/src/net/netip/netip.go:688
	// _ = "end of CoverTab[12214]"
//line /usr/local/go/src/net/netip/netip.go:688
	_go_fuzz_dep_.CoverTab[12215]++
							if ip.z == z0 {
//line /usr/local/go/src/net/netip/netip.go:689
		_go_fuzz_dep_.CoverTab[12220]++
								panic("As4 called on IP zero value")
//line /usr/local/go/src/net/netip/netip.go:690
		// _ = "end of CoverTab[12220]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:691
		_go_fuzz_dep_.CoverTab[12221]++
//line /usr/local/go/src/net/netip/netip.go:691
		// _ = "end of CoverTab[12221]"
//line /usr/local/go/src/net/netip/netip.go:691
	}
//line /usr/local/go/src/net/netip/netip.go:691
	// _ = "end of CoverTab[12215]"
//line /usr/local/go/src/net/netip/netip.go:691
	_go_fuzz_dep_.CoverTab[12216]++
							panic("As4 called on IPv6 address")
//line /usr/local/go/src/net/netip/netip.go:692
	// _ = "end of CoverTab[12216]"
}

// AsSlice returns an IPv4 or IPv6 address in its respective 4-byte or 16-byte representation.
func (ip Addr) AsSlice() []byte {
//line /usr/local/go/src/net/netip/netip.go:696
	_go_fuzz_dep_.CoverTab[12222]++
							switch ip.z {
	case z0:
//line /usr/local/go/src/net/netip/netip.go:698
		_go_fuzz_dep_.CoverTab[12223]++
								return nil
//line /usr/local/go/src/net/netip/netip.go:699
		// _ = "end of CoverTab[12223]"
	case z4:
//line /usr/local/go/src/net/netip/netip.go:700
		_go_fuzz_dep_.CoverTab[12224]++
								var ret [4]byte
								bePutUint32(ret[:], uint32(ip.addr.lo))
								return ret[:]
//line /usr/local/go/src/net/netip/netip.go:703
		// _ = "end of CoverTab[12224]"
	default:
//line /usr/local/go/src/net/netip/netip.go:704
		_go_fuzz_dep_.CoverTab[12225]++
								var ret [16]byte
								bePutUint64(ret[:8], ip.addr.hi)
								bePutUint64(ret[8:], ip.addr.lo)
								return ret[:]
//line /usr/local/go/src/net/netip/netip.go:708
		// _ = "end of CoverTab[12225]"
	}
//line /usr/local/go/src/net/netip/netip.go:709
	// _ = "end of CoverTab[12222]"
}

// Next returns the address following ip.
//line /usr/local/go/src/net/netip/netip.go:712
// If there is none, it returns the zero Addr.
//line /usr/local/go/src/net/netip/netip.go:714
func (ip Addr) Next() Addr {
//line /usr/local/go/src/net/netip/netip.go:714
	_go_fuzz_dep_.CoverTab[12226]++
							ip.addr = ip.addr.addOne()
							if ip.Is4() {
//line /usr/local/go/src/net/netip/netip.go:716
		_go_fuzz_dep_.CoverTab[12228]++
								if uint32(ip.addr.lo) == 0 {
//line /usr/local/go/src/net/netip/netip.go:717
			_go_fuzz_dep_.CoverTab[12229]++

									return Addr{}
//line /usr/local/go/src/net/netip/netip.go:719
			// _ = "end of CoverTab[12229]"
		} else {
//line /usr/local/go/src/net/netip/netip.go:720
			_go_fuzz_dep_.CoverTab[12230]++
//line /usr/local/go/src/net/netip/netip.go:720
			// _ = "end of CoverTab[12230]"
//line /usr/local/go/src/net/netip/netip.go:720
		}
//line /usr/local/go/src/net/netip/netip.go:720
		// _ = "end of CoverTab[12228]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:721
		_go_fuzz_dep_.CoverTab[12231]++
								if ip.addr.isZero() {
//line /usr/local/go/src/net/netip/netip.go:722
			_go_fuzz_dep_.CoverTab[12232]++

									return Addr{}
//line /usr/local/go/src/net/netip/netip.go:724
			// _ = "end of CoverTab[12232]"
		} else {
//line /usr/local/go/src/net/netip/netip.go:725
			_go_fuzz_dep_.CoverTab[12233]++
//line /usr/local/go/src/net/netip/netip.go:725
			// _ = "end of CoverTab[12233]"
//line /usr/local/go/src/net/netip/netip.go:725
		}
//line /usr/local/go/src/net/netip/netip.go:725
		// _ = "end of CoverTab[12231]"
	}
//line /usr/local/go/src/net/netip/netip.go:726
	// _ = "end of CoverTab[12226]"
//line /usr/local/go/src/net/netip/netip.go:726
	_go_fuzz_dep_.CoverTab[12227]++
							return ip
//line /usr/local/go/src/net/netip/netip.go:727
	// _ = "end of CoverTab[12227]"
}

// Prev returns the IP before ip.
//line /usr/local/go/src/net/netip/netip.go:730
// If there is none, it returns the IP zero value.
//line /usr/local/go/src/net/netip/netip.go:732
func (ip Addr) Prev() Addr {
//line /usr/local/go/src/net/netip/netip.go:732
	_go_fuzz_dep_.CoverTab[12234]++
							if ip.Is4() {
//line /usr/local/go/src/net/netip/netip.go:733
		_go_fuzz_dep_.CoverTab[12236]++
								if uint32(ip.addr.lo) == 0 {
//line /usr/local/go/src/net/netip/netip.go:734
			_go_fuzz_dep_.CoverTab[12237]++
									return Addr{}
//line /usr/local/go/src/net/netip/netip.go:735
			// _ = "end of CoverTab[12237]"
		} else {
//line /usr/local/go/src/net/netip/netip.go:736
			_go_fuzz_dep_.CoverTab[12238]++
//line /usr/local/go/src/net/netip/netip.go:736
			// _ = "end of CoverTab[12238]"
//line /usr/local/go/src/net/netip/netip.go:736
		}
//line /usr/local/go/src/net/netip/netip.go:736
		// _ = "end of CoverTab[12236]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:737
		_go_fuzz_dep_.CoverTab[12239]++
//line /usr/local/go/src/net/netip/netip.go:737
		if ip.addr.isZero() {
//line /usr/local/go/src/net/netip/netip.go:737
			_go_fuzz_dep_.CoverTab[12240]++
									return Addr{}
//line /usr/local/go/src/net/netip/netip.go:738
			// _ = "end of CoverTab[12240]"
		} else {
//line /usr/local/go/src/net/netip/netip.go:739
			_go_fuzz_dep_.CoverTab[12241]++
//line /usr/local/go/src/net/netip/netip.go:739
			// _ = "end of CoverTab[12241]"
//line /usr/local/go/src/net/netip/netip.go:739
		}
//line /usr/local/go/src/net/netip/netip.go:739
		// _ = "end of CoverTab[12239]"
//line /usr/local/go/src/net/netip/netip.go:739
	}
//line /usr/local/go/src/net/netip/netip.go:739
	// _ = "end of CoverTab[12234]"
//line /usr/local/go/src/net/netip/netip.go:739
	_go_fuzz_dep_.CoverTab[12235]++
							ip.addr = ip.addr.subOne()
							return ip
//line /usr/local/go/src/net/netip/netip.go:741
	// _ = "end of CoverTab[12235]"
}

// String returns the string form of the IP address ip.
//line /usr/local/go/src/net/netip/netip.go:744
// It returns one of 5 forms:
//line /usr/local/go/src/net/netip/netip.go:744
//
//line /usr/local/go/src/net/netip/netip.go:744
//   - "invalid IP", if ip is the zero Addr
//line /usr/local/go/src/net/netip/netip.go:744
//   - IPv4 dotted decimal ("192.0.2.1")
//line /usr/local/go/src/net/netip/netip.go:744
//   - IPv6 ("2001:db8::1")
//line /usr/local/go/src/net/netip/netip.go:744
//   - "::ffff:1.2.3.4" (if Is4In6)
//line /usr/local/go/src/net/netip/netip.go:744
//   - IPv6 with zone ("fe80:db8::1%eth0")
//line /usr/local/go/src/net/netip/netip.go:744
//
//line /usr/local/go/src/net/netip/netip.go:744
// Note that unlike package net's IP.String method,
//line /usr/local/go/src/net/netip/netip.go:744
// IPv4-mapped IPv6 addresses format with a "::ffff:"
//line /usr/local/go/src/net/netip/netip.go:744
// prefix before the dotted quad.
//line /usr/local/go/src/net/netip/netip.go:756
func (ip Addr) String() string {
//line /usr/local/go/src/net/netip/netip.go:756
	_go_fuzz_dep_.CoverTab[12242]++
							switch ip.z {
	case z0:
//line /usr/local/go/src/net/netip/netip.go:758
		_go_fuzz_dep_.CoverTab[12243]++
								return "invalid IP"
//line /usr/local/go/src/net/netip/netip.go:759
		// _ = "end of CoverTab[12243]"
	case z4:
//line /usr/local/go/src/net/netip/netip.go:760
		_go_fuzz_dep_.CoverTab[12244]++
								return ip.string4()
//line /usr/local/go/src/net/netip/netip.go:761
		// _ = "end of CoverTab[12244]"
	default:
//line /usr/local/go/src/net/netip/netip.go:762
		_go_fuzz_dep_.CoverTab[12245]++
								if ip.Is4In6() {
//line /usr/local/go/src/net/netip/netip.go:763
			_go_fuzz_dep_.CoverTab[12247]++
									if z := ip.Zone(); z != "" {
//line /usr/local/go/src/net/netip/netip.go:764
				_go_fuzz_dep_.CoverTab[12248]++
										return "::ffff:" + ip.Unmap().string4() + "%" + z
//line /usr/local/go/src/net/netip/netip.go:765
				// _ = "end of CoverTab[12248]"
			} else {
//line /usr/local/go/src/net/netip/netip.go:766
				_go_fuzz_dep_.CoverTab[12249]++
										return "::ffff:" + ip.Unmap().string4()
//line /usr/local/go/src/net/netip/netip.go:767
				// _ = "end of CoverTab[12249]"
			}
//line /usr/local/go/src/net/netip/netip.go:768
			// _ = "end of CoverTab[12247]"
		} else {
//line /usr/local/go/src/net/netip/netip.go:769
			_go_fuzz_dep_.CoverTab[12250]++
//line /usr/local/go/src/net/netip/netip.go:769
			// _ = "end of CoverTab[12250]"
//line /usr/local/go/src/net/netip/netip.go:769
		}
//line /usr/local/go/src/net/netip/netip.go:769
		// _ = "end of CoverTab[12245]"
//line /usr/local/go/src/net/netip/netip.go:769
		_go_fuzz_dep_.CoverTab[12246]++
								return ip.string6()
//line /usr/local/go/src/net/netip/netip.go:770
		// _ = "end of CoverTab[12246]"
	}
//line /usr/local/go/src/net/netip/netip.go:771
	// _ = "end of CoverTab[12242]"
}

// AppendTo appends a text encoding of ip,
//line /usr/local/go/src/net/netip/netip.go:774
// as generated by MarshalText,
//line /usr/local/go/src/net/netip/netip.go:774
// to b and returns the extended buffer.
//line /usr/local/go/src/net/netip/netip.go:777
func (ip Addr) AppendTo(b []byte) []byte {
//line /usr/local/go/src/net/netip/netip.go:777
	_go_fuzz_dep_.CoverTab[12251]++
							switch ip.z {
	case z0:
//line /usr/local/go/src/net/netip/netip.go:779
		_go_fuzz_dep_.CoverTab[12252]++
								return b
//line /usr/local/go/src/net/netip/netip.go:780
		// _ = "end of CoverTab[12252]"
	case z4:
//line /usr/local/go/src/net/netip/netip.go:781
		_go_fuzz_dep_.CoverTab[12253]++
								return ip.appendTo4(b)
//line /usr/local/go/src/net/netip/netip.go:782
		// _ = "end of CoverTab[12253]"
	default:
//line /usr/local/go/src/net/netip/netip.go:783
		_go_fuzz_dep_.CoverTab[12254]++
								if ip.Is4In6() {
//line /usr/local/go/src/net/netip/netip.go:784
			_go_fuzz_dep_.CoverTab[12256]++
									b = append(b, "::ffff:"...)
									b = ip.Unmap().appendTo4(b)
									if z := ip.Zone(); z != "" {
//line /usr/local/go/src/net/netip/netip.go:787
				_go_fuzz_dep_.CoverTab[12258]++
										b = append(b, '%')
										b = append(b, z...)
//line /usr/local/go/src/net/netip/netip.go:789
				// _ = "end of CoverTab[12258]"
			} else {
//line /usr/local/go/src/net/netip/netip.go:790
				_go_fuzz_dep_.CoverTab[12259]++
//line /usr/local/go/src/net/netip/netip.go:790
				// _ = "end of CoverTab[12259]"
//line /usr/local/go/src/net/netip/netip.go:790
			}
//line /usr/local/go/src/net/netip/netip.go:790
			// _ = "end of CoverTab[12256]"
//line /usr/local/go/src/net/netip/netip.go:790
			_go_fuzz_dep_.CoverTab[12257]++
									return b
//line /usr/local/go/src/net/netip/netip.go:791
			// _ = "end of CoverTab[12257]"
		} else {
//line /usr/local/go/src/net/netip/netip.go:792
			_go_fuzz_dep_.CoverTab[12260]++
//line /usr/local/go/src/net/netip/netip.go:792
			// _ = "end of CoverTab[12260]"
//line /usr/local/go/src/net/netip/netip.go:792
		}
//line /usr/local/go/src/net/netip/netip.go:792
		// _ = "end of CoverTab[12254]"
//line /usr/local/go/src/net/netip/netip.go:792
		_go_fuzz_dep_.CoverTab[12255]++
								return ip.appendTo6(b)
//line /usr/local/go/src/net/netip/netip.go:793
		// _ = "end of CoverTab[12255]"
	}
//line /usr/local/go/src/net/netip/netip.go:794
	// _ = "end of CoverTab[12251]"
}

// digits is a string of the hex digits from 0 to f. It's used in
//line /usr/local/go/src/net/netip/netip.go:797
// appendDecimal and appendHex to format IP addresses.
//line /usr/local/go/src/net/netip/netip.go:799
const digits = "0123456789abcdef"

// appendDecimal appends the decimal string representation of x to b.
func appendDecimal(b []byte, x uint8) []byte {
//line /usr/local/go/src/net/netip/netip.go:802
	_go_fuzz_dep_.CoverTab[12261]++

//line /usr/local/go/src/net/netip/netip.go:806
	if x >= 100 {
//line /usr/local/go/src/net/netip/netip.go:806
		_go_fuzz_dep_.CoverTab[12264]++
								b = append(b, digits[x/100])
//line /usr/local/go/src/net/netip/netip.go:807
		// _ = "end of CoverTab[12264]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:808
		_go_fuzz_dep_.CoverTab[12265]++
//line /usr/local/go/src/net/netip/netip.go:808
		// _ = "end of CoverTab[12265]"
//line /usr/local/go/src/net/netip/netip.go:808
	}
//line /usr/local/go/src/net/netip/netip.go:808
	// _ = "end of CoverTab[12261]"
//line /usr/local/go/src/net/netip/netip.go:808
	_go_fuzz_dep_.CoverTab[12262]++
							if x >= 10 {
//line /usr/local/go/src/net/netip/netip.go:809
		_go_fuzz_dep_.CoverTab[12266]++
								b = append(b, digits[x/10%10])
//line /usr/local/go/src/net/netip/netip.go:810
		// _ = "end of CoverTab[12266]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:811
		_go_fuzz_dep_.CoverTab[12267]++
//line /usr/local/go/src/net/netip/netip.go:811
		// _ = "end of CoverTab[12267]"
//line /usr/local/go/src/net/netip/netip.go:811
	}
//line /usr/local/go/src/net/netip/netip.go:811
	// _ = "end of CoverTab[12262]"
//line /usr/local/go/src/net/netip/netip.go:811
	_go_fuzz_dep_.CoverTab[12263]++
							return append(b, digits[x%10])
//line /usr/local/go/src/net/netip/netip.go:812
	// _ = "end of CoverTab[12263]"
}

// appendHex appends the hex string representation of x to b.
func appendHex(b []byte, x uint16) []byte {
//line /usr/local/go/src/net/netip/netip.go:816
	_go_fuzz_dep_.CoverTab[12268]++

//line /usr/local/go/src/net/netip/netip.go:820
	if x >= 0x1000 {
//line /usr/local/go/src/net/netip/netip.go:820
		_go_fuzz_dep_.CoverTab[12272]++
								b = append(b, digits[x>>12])
//line /usr/local/go/src/net/netip/netip.go:821
		// _ = "end of CoverTab[12272]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:822
		_go_fuzz_dep_.CoverTab[12273]++
//line /usr/local/go/src/net/netip/netip.go:822
		// _ = "end of CoverTab[12273]"
//line /usr/local/go/src/net/netip/netip.go:822
	}
//line /usr/local/go/src/net/netip/netip.go:822
	// _ = "end of CoverTab[12268]"
//line /usr/local/go/src/net/netip/netip.go:822
	_go_fuzz_dep_.CoverTab[12269]++
							if x >= 0x100 {
//line /usr/local/go/src/net/netip/netip.go:823
		_go_fuzz_dep_.CoverTab[12274]++
								b = append(b, digits[x>>8&0xf])
//line /usr/local/go/src/net/netip/netip.go:824
		// _ = "end of CoverTab[12274]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:825
		_go_fuzz_dep_.CoverTab[12275]++
//line /usr/local/go/src/net/netip/netip.go:825
		// _ = "end of CoverTab[12275]"
//line /usr/local/go/src/net/netip/netip.go:825
	}
//line /usr/local/go/src/net/netip/netip.go:825
	// _ = "end of CoverTab[12269]"
//line /usr/local/go/src/net/netip/netip.go:825
	_go_fuzz_dep_.CoverTab[12270]++
							if x >= 0x10 {
//line /usr/local/go/src/net/netip/netip.go:826
		_go_fuzz_dep_.CoverTab[12276]++
								b = append(b, digits[x>>4&0xf])
//line /usr/local/go/src/net/netip/netip.go:827
		// _ = "end of CoverTab[12276]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:828
		_go_fuzz_dep_.CoverTab[12277]++
//line /usr/local/go/src/net/netip/netip.go:828
		// _ = "end of CoverTab[12277]"
//line /usr/local/go/src/net/netip/netip.go:828
	}
//line /usr/local/go/src/net/netip/netip.go:828
	// _ = "end of CoverTab[12270]"
//line /usr/local/go/src/net/netip/netip.go:828
	_go_fuzz_dep_.CoverTab[12271]++
							return append(b, digits[x&0xf])
//line /usr/local/go/src/net/netip/netip.go:829
	// _ = "end of CoverTab[12271]"
}

// appendHexPad appends the fully padded hex string representation of x to b.
func appendHexPad(b []byte, x uint16) []byte {
//line /usr/local/go/src/net/netip/netip.go:833
	_go_fuzz_dep_.CoverTab[12278]++
							return append(b, digits[x>>12], digits[x>>8&0xf], digits[x>>4&0xf], digits[x&0xf])
//line /usr/local/go/src/net/netip/netip.go:834
	// _ = "end of CoverTab[12278]"
}

func (ip Addr) string4() string {
//line /usr/local/go/src/net/netip/netip.go:837
	_go_fuzz_dep_.CoverTab[12279]++
							const max = len("255.255.255.255")
							ret := make([]byte, 0, max)
							ret = ip.appendTo4(ret)
							return string(ret)
//line /usr/local/go/src/net/netip/netip.go:841
	// _ = "end of CoverTab[12279]"
}

func (ip Addr) appendTo4(ret []byte) []byte {
//line /usr/local/go/src/net/netip/netip.go:844
	_go_fuzz_dep_.CoverTab[12280]++
							ret = appendDecimal(ret, ip.v4(0))
							ret = append(ret, '.')
							ret = appendDecimal(ret, ip.v4(1))
							ret = append(ret, '.')
							ret = appendDecimal(ret, ip.v4(2))
							ret = append(ret, '.')
							ret = appendDecimal(ret, ip.v4(3))
							return ret
//line /usr/local/go/src/net/netip/netip.go:852
	// _ = "end of CoverTab[12280]"
}

// string6 formats ip in IPv6 textual representation. It follows the
//line /usr/local/go/src/net/netip/netip.go:855
// guidelines in section 4 of RFC 5952
//line /usr/local/go/src/net/netip/netip.go:855
// (https://tools.ietf.org/html/rfc5952#section-4): no unnecessary
//line /usr/local/go/src/net/netip/netip.go:855
// zeros, use :: to elide the longest run of zeros, and don't use ::
//line /usr/local/go/src/net/netip/netip.go:855
// to compact a single zero field.
//line /usr/local/go/src/net/netip/netip.go:860
func (ip Addr) string6() string {
//line /usr/local/go/src/net/netip/netip.go:860
	_go_fuzz_dep_.CoverTab[12281]++
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
//line /usr/local/go/src/net/netip/netip.go:871
	// _ = "end of CoverTab[12281]"
}

func (ip Addr) appendTo6(ret []byte) []byte {
//line /usr/local/go/src/net/netip/netip.go:874
	_go_fuzz_dep_.CoverTab[12282]++
							zeroStart, zeroEnd := uint8(255), uint8(255)
							for i := uint8(0); i < 8; i++ {
//line /usr/local/go/src/net/netip/netip.go:876
		_go_fuzz_dep_.CoverTab[12286]++
								j := i
								for j < 8 && func() bool {
//line /usr/local/go/src/net/netip/netip.go:878
			_go_fuzz_dep_.CoverTab[12288]++
//line /usr/local/go/src/net/netip/netip.go:878
			return ip.v6u16(j) == 0
//line /usr/local/go/src/net/netip/netip.go:878
			// _ = "end of CoverTab[12288]"
//line /usr/local/go/src/net/netip/netip.go:878
		}() {
//line /usr/local/go/src/net/netip/netip.go:878
			_go_fuzz_dep_.CoverTab[12289]++
									j++
//line /usr/local/go/src/net/netip/netip.go:879
			// _ = "end of CoverTab[12289]"
		}
//line /usr/local/go/src/net/netip/netip.go:880
		// _ = "end of CoverTab[12286]"
//line /usr/local/go/src/net/netip/netip.go:880
		_go_fuzz_dep_.CoverTab[12287]++
								if l := j - i; l >= 2 && func() bool {
//line /usr/local/go/src/net/netip/netip.go:881
			_go_fuzz_dep_.CoverTab[12290]++
//line /usr/local/go/src/net/netip/netip.go:881
			return l > zeroEnd-zeroStart
//line /usr/local/go/src/net/netip/netip.go:881
			// _ = "end of CoverTab[12290]"
//line /usr/local/go/src/net/netip/netip.go:881
		}() {
//line /usr/local/go/src/net/netip/netip.go:881
			_go_fuzz_dep_.CoverTab[12291]++
									zeroStart, zeroEnd = i, j
//line /usr/local/go/src/net/netip/netip.go:882
			// _ = "end of CoverTab[12291]"
		} else {
//line /usr/local/go/src/net/netip/netip.go:883
			_go_fuzz_dep_.CoverTab[12292]++
//line /usr/local/go/src/net/netip/netip.go:883
			// _ = "end of CoverTab[12292]"
//line /usr/local/go/src/net/netip/netip.go:883
		}
//line /usr/local/go/src/net/netip/netip.go:883
		// _ = "end of CoverTab[12287]"
	}
//line /usr/local/go/src/net/netip/netip.go:884
	// _ = "end of CoverTab[12282]"
//line /usr/local/go/src/net/netip/netip.go:884
	_go_fuzz_dep_.CoverTab[12283]++

							for i := uint8(0); i < 8; i++ {
//line /usr/local/go/src/net/netip/netip.go:886
		_go_fuzz_dep_.CoverTab[12293]++
								if i == zeroStart {
//line /usr/local/go/src/net/netip/netip.go:887
			_go_fuzz_dep_.CoverTab[12295]++
									ret = append(ret, ':', ':')
									i = zeroEnd
									if i >= 8 {
//line /usr/local/go/src/net/netip/netip.go:890
				_go_fuzz_dep_.CoverTab[12296]++
										break
//line /usr/local/go/src/net/netip/netip.go:891
				// _ = "end of CoverTab[12296]"
			} else {
//line /usr/local/go/src/net/netip/netip.go:892
				_go_fuzz_dep_.CoverTab[12297]++
//line /usr/local/go/src/net/netip/netip.go:892
				// _ = "end of CoverTab[12297]"
//line /usr/local/go/src/net/netip/netip.go:892
			}
//line /usr/local/go/src/net/netip/netip.go:892
			// _ = "end of CoverTab[12295]"
		} else {
//line /usr/local/go/src/net/netip/netip.go:893
			_go_fuzz_dep_.CoverTab[12298]++
//line /usr/local/go/src/net/netip/netip.go:893
			if i > 0 {
//line /usr/local/go/src/net/netip/netip.go:893
				_go_fuzz_dep_.CoverTab[12299]++
										ret = append(ret, ':')
//line /usr/local/go/src/net/netip/netip.go:894
				// _ = "end of CoverTab[12299]"
			} else {
//line /usr/local/go/src/net/netip/netip.go:895
				_go_fuzz_dep_.CoverTab[12300]++
//line /usr/local/go/src/net/netip/netip.go:895
				// _ = "end of CoverTab[12300]"
//line /usr/local/go/src/net/netip/netip.go:895
			}
//line /usr/local/go/src/net/netip/netip.go:895
			// _ = "end of CoverTab[12298]"
//line /usr/local/go/src/net/netip/netip.go:895
		}
//line /usr/local/go/src/net/netip/netip.go:895
		// _ = "end of CoverTab[12293]"
//line /usr/local/go/src/net/netip/netip.go:895
		_go_fuzz_dep_.CoverTab[12294]++

								ret = appendHex(ret, ip.v6u16(i))
//line /usr/local/go/src/net/netip/netip.go:897
		// _ = "end of CoverTab[12294]"
	}
//line /usr/local/go/src/net/netip/netip.go:898
	// _ = "end of CoverTab[12283]"
//line /usr/local/go/src/net/netip/netip.go:898
	_go_fuzz_dep_.CoverTab[12284]++

							if ip.z != z6noz {
//line /usr/local/go/src/net/netip/netip.go:900
		_go_fuzz_dep_.CoverTab[12301]++
								ret = append(ret, '%')
								ret = append(ret, ip.Zone()...)
//line /usr/local/go/src/net/netip/netip.go:902
		// _ = "end of CoverTab[12301]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:903
		_go_fuzz_dep_.CoverTab[12302]++
//line /usr/local/go/src/net/netip/netip.go:903
		// _ = "end of CoverTab[12302]"
//line /usr/local/go/src/net/netip/netip.go:903
	}
//line /usr/local/go/src/net/netip/netip.go:903
	// _ = "end of CoverTab[12284]"
//line /usr/local/go/src/net/netip/netip.go:903
	_go_fuzz_dep_.CoverTab[12285]++
							return ret
//line /usr/local/go/src/net/netip/netip.go:904
	// _ = "end of CoverTab[12285]"
}

// StringExpanded is like String but IPv6 addresses are expanded with leading
//line /usr/local/go/src/net/netip/netip.go:907
// zeroes and no "::" compression. For example, "2001:db8::1" becomes
//line /usr/local/go/src/net/netip/netip.go:907
// "2001:0db8:0000:0000:0000:0000:0000:0001".
//line /usr/local/go/src/net/netip/netip.go:910
func (ip Addr) StringExpanded() string {
//line /usr/local/go/src/net/netip/netip.go:910
	_go_fuzz_dep_.CoverTab[12303]++
							switch ip.z {
	case z0, z4:
//line /usr/local/go/src/net/netip/netip.go:912
		_go_fuzz_dep_.CoverTab[12307]++
								return ip.String()
//line /usr/local/go/src/net/netip/netip.go:913
		// _ = "end of CoverTab[12307]"
//line /usr/local/go/src/net/netip/netip.go:913
	default:
//line /usr/local/go/src/net/netip/netip.go:913
		_go_fuzz_dep_.CoverTab[12308]++
//line /usr/local/go/src/net/netip/netip.go:913
		// _ = "end of CoverTab[12308]"
	}
//line /usr/local/go/src/net/netip/netip.go:914
	// _ = "end of CoverTab[12303]"
//line /usr/local/go/src/net/netip/netip.go:914
	_go_fuzz_dep_.CoverTab[12304]++

							const size = len("ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff")
							ret := make([]byte, 0, size)
							for i := uint8(0); i < 8; i++ {
//line /usr/local/go/src/net/netip/netip.go:918
		_go_fuzz_dep_.CoverTab[12309]++
								if i > 0 {
//line /usr/local/go/src/net/netip/netip.go:919
			_go_fuzz_dep_.CoverTab[12311]++
									ret = append(ret, ':')
//line /usr/local/go/src/net/netip/netip.go:920
			// _ = "end of CoverTab[12311]"
		} else {
//line /usr/local/go/src/net/netip/netip.go:921
			_go_fuzz_dep_.CoverTab[12312]++
//line /usr/local/go/src/net/netip/netip.go:921
			// _ = "end of CoverTab[12312]"
//line /usr/local/go/src/net/netip/netip.go:921
		}
//line /usr/local/go/src/net/netip/netip.go:921
		// _ = "end of CoverTab[12309]"
//line /usr/local/go/src/net/netip/netip.go:921
		_go_fuzz_dep_.CoverTab[12310]++

								ret = appendHexPad(ret, ip.v6u16(i))
//line /usr/local/go/src/net/netip/netip.go:923
		// _ = "end of CoverTab[12310]"
	}
//line /usr/local/go/src/net/netip/netip.go:924
	// _ = "end of CoverTab[12304]"
//line /usr/local/go/src/net/netip/netip.go:924
	_go_fuzz_dep_.CoverTab[12305]++

							if ip.z != z6noz {
//line /usr/local/go/src/net/netip/netip.go:926
		_go_fuzz_dep_.CoverTab[12313]++

//line /usr/local/go/src/net/netip/netip.go:929
		ret = append(ret, '%')
								ret = append(ret, ip.Zone()...)
//line /usr/local/go/src/net/netip/netip.go:930
		// _ = "end of CoverTab[12313]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:931
		_go_fuzz_dep_.CoverTab[12314]++
//line /usr/local/go/src/net/netip/netip.go:931
		// _ = "end of CoverTab[12314]"
//line /usr/local/go/src/net/netip/netip.go:931
	}
//line /usr/local/go/src/net/netip/netip.go:931
	// _ = "end of CoverTab[12305]"
//line /usr/local/go/src/net/netip/netip.go:931
	_go_fuzz_dep_.CoverTab[12306]++
							return string(ret)
//line /usr/local/go/src/net/netip/netip.go:932
	// _ = "end of CoverTab[12306]"
}

// MarshalText implements the encoding.TextMarshaler interface,
//line /usr/local/go/src/net/netip/netip.go:935
// The encoding is the same as returned by String, with one exception:
//line /usr/local/go/src/net/netip/netip.go:935
// If ip is the zero Addr, the encoding is the empty string.
//line /usr/local/go/src/net/netip/netip.go:938
func (ip Addr) MarshalText() ([]byte, error) {
//line /usr/local/go/src/net/netip/netip.go:938
	_go_fuzz_dep_.CoverTab[12315]++
							switch ip.z {
	case z0:
//line /usr/local/go/src/net/netip/netip.go:940
		_go_fuzz_dep_.CoverTab[12316]++
								return []byte(""), nil
//line /usr/local/go/src/net/netip/netip.go:941
		// _ = "end of CoverTab[12316]"
	case z4:
//line /usr/local/go/src/net/netip/netip.go:942
		_go_fuzz_dep_.CoverTab[12317]++
								max := len("255.255.255.255")
								b := make([]byte, 0, max)
								return ip.appendTo4(b), nil
//line /usr/local/go/src/net/netip/netip.go:945
		// _ = "end of CoverTab[12317]"
	default:
//line /usr/local/go/src/net/netip/netip.go:946
		_go_fuzz_dep_.CoverTab[12318]++
								max := len("ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff%enp5s0")
								b := make([]byte, 0, max)
								if ip.Is4In6() {
//line /usr/local/go/src/net/netip/netip.go:949
			_go_fuzz_dep_.CoverTab[12320]++
									b = append(b, "::ffff:"...)
									b = ip.Unmap().appendTo4(b)
									if z := ip.Zone(); z != "" {
//line /usr/local/go/src/net/netip/netip.go:952
				_go_fuzz_dep_.CoverTab[12322]++
										b = append(b, '%')
										b = append(b, z...)
//line /usr/local/go/src/net/netip/netip.go:954
				// _ = "end of CoverTab[12322]"
			} else {
//line /usr/local/go/src/net/netip/netip.go:955
				_go_fuzz_dep_.CoverTab[12323]++
//line /usr/local/go/src/net/netip/netip.go:955
				// _ = "end of CoverTab[12323]"
//line /usr/local/go/src/net/netip/netip.go:955
			}
//line /usr/local/go/src/net/netip/netip.go:955
			// _ = "end of CoverTab[12320]"
//line /usr/local/go/src/net/netip/netip.go:955
			_go_fuzz_dep_.CoverTab[12321]++
									return b, nil
//line /usr/local/go/src/net/netip/netip.go:956
			// _ = "end of CoverTab[12321]"
		} else {
//line /usr/local/go/src/net/netip/netip.go:957
			_go_fuzz_dep_.CoverTab[12324]++
//line /usr/local/go/src/net/netip/netip.go:957
			// _ = "end of CoverTab[12324]"
//line /usr/local/go/src/net/netip/netip.go:957
		}
//line /usr/local/go/src/net/netip/netip.go:957
		// _ = "end of CoverTab[12318]"
//line /usr/local/go/src/net/netip/netip.go:957
		_go_fuzz_dep_.CoverTab[12319]++
								return ip.appendTo6(b), nil
//line /usr/local/go/src/net/netip/netip.go:958
		// _ = "end of CoverTab[12319]"
	}
//line /usr/local/go/src/net/netip/netip.go:959
	// _ = "end of CoverTab[12315]"

}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
//line /usr/local/go/src/net/netip/netip.go:963
// The IP address is expected in a form accepted by ParseAddr.
//line /usr/local/go/src/net/netip/netip.go:963
//
//line /usr/local/go/src/net/netip/netip.go:963
// If text is empty, UnmarshalText sets *ip to the zero Addr and
//line /usr/local/go/src/net/netip/netip.go:963
// returns no error.
//line /usr/local/go/src/net/netip/netip.go:968
func (ip *Addr) UnmarshalText(text []byte) error {
//line /usr/local/go/src/net/netip/netip.go:968
	_go_fuzz_dep_.CoverTab[12325]++
							if len(text) == 0 {
//line /usr/local/go/src/net/netip/netip.go:969
		_go_fuzz_dep_.CoverTab[12327]++
								*ip = Addr{}
								return nil
//line /usr/local/go/src/net/netip/netip.go:971
		// _ = "end of CoverTab[12327]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:972
		_go_fuzz_dep_.CoverTab[12328]++
//line /usr/local/go/src/net/netip/netip.go:972
		// _ = "end of CoverTab[12328]"
//line /usr/local/go/src/net/netip/netip.go:972
	}
//line /usr/local/go/src/net/netip/netip.go:972
	// _ = "end of CoverTab[12325]"
//line /usr/local/go/src/net/netip/netip.go:972
	_go_fuzz_dep_.CoverTab[12326]++
							var err error
							*ip, err = ParseAddr(string(text))
							return err
//line /usr/local/go/src/net/netip/netip.go:975
	// _ = "end of CoverTab[12326]"
}

func (ip Addr) marshalBinaryWithTrailingBytes(trailingBytes int) []byte {
//line /usr/local/go/src/net/netip/netip.go:978
	_go_fuzz_dep_.CoverTab[12329]++
							var b []byte
							switch ip.z {
	case z0:
//line /usr/local/go/src/net/netip/netip.go:981
		_go_fuzz_dep_.CoverTab[12331]++
								b = make([]byte, trailingBytes)
//line /usr/local/go/src/net/netip/netip.go:982
		// _ = "end of CoverTab[12331]"
	case z4:
//line /usr/local/go/src/net/netip/netip.go:983
		_go_fuzz_dep_.CoverTab[12332]++
								b = make([]byte, 4+trailingBytes)
								bePutUint32(b, uint32(ip.addr.lo))
//line /usr/local/go/src/net/netip/netip.go:985
		// _ = "end of CoverTab[12332]"
	default:
//line /usr/local/go/src/net/netip/netip.go:986
		_go_fuzz_dep_.CoverTab[12333]++
								z := ip.Zone()
								b = make([]byte, 16+len(z)+trailingBytes)
								bePutUint64(b[:8], ip.addr.hi)
								bePutUint64(b[8:], ip.addr.lo)
								copy(b[16:], z)
//line /usr/local/go/src/net/netip/netip.go:991
		// _ = "end of CoverTab[12333]"
	}
//line /usr/local/go/src/net/netip/netip.go:992
	// _ = "end of CoverTab[12329]"
//line /usr/local/go/src/net/netip/netip.go:992
	_go_fuzz_dep_.CoverTab[12330]++
							return b
//line /usr/local/go/src/net/netip/netip.go:993
	// _ = "end of CoverTab[12330]"
}

// MarshalBinary implements the encoding.BinaryMarshaler interface.
//line /usr/local/go/src/net/netip/netip.go:996
// It returns a zero-length slice for the zero Addr,
//line /usr/local/go/src/net/netip/netip.go:996
// the 4-byte form for an IPv4 address,
//line /usr/local/go/src/net/netip/netip.go:996
// and the 16-byte form with zone appended for an IPv6 address.
//line /usr/local/go/src/net/netip/netip.go:1000
func (ip Addr) MarshalBinary() ([]byte, error) {
//line /usr/local/go/src/net/netip/netip.go:1000
	_go_fuzz_dep_.CoverTab[12334]++
							return ip.marshalBinaryWithTrailingBytes(0), nil
//line /usr/local/go/src/net/netip/netip.go:1001
	// _ = "end of CoverTab[12334]"
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface.
//line /usr/local/go/src/net/netip/netip.go:1004
// It expects data in the form generated by MarshalBinary.
//line /usr/local/go/src/net/netip/netip.go:1006
func (ip *Addr) UnmarshalBinary(b []byte) error {
//line /usr/local/go/src/net/netip/netip.go:1006
	_go_fuzz_dep_.CoverTab[12335]++
							n := len(b)
							switch {
	case n == 0:
//line /usr/local/go/src/net/netip/netip.go:1009
		_go_fuzz_dep_.CoverTab[12337]++
								*ip = Addr{}
								return nil
//line /usr/local/go/src/net/netip/netip.go:1011
		// _ = "end of CoverTab[12337]"
	case n == 4:
//line /usr/local/go/src/net/netip/netip.go:1012
		_go_fuzz_dep_.CoverTab[12338]++
								*ip = AddrFrom4([4]byte(b))
								return nil
//line /usr/local/go/src/net/netip/netip.go:1014
		// _ = "end of CoverTab[12338]"
	case n == 16:
//line /usr/local/go/src/net/netip/netip.go:1015
		_go_fuzz_dep_.CoverTab[12339]++
								*ip = AddrFrom16([16]byte(b))
								return nil
//line /usr/local/go/src/net/netip/netip.go:1017
		// _ = "end of CoverTab[12339]"
	case n > 16:
//line /usr/local/go/src/net/netip/netip.go:1018
		_go_fuzz_dep_.CoverTab[12340]++
								*ip = AddrFrom16([16]byte(b[:16])).WithZone(string(b[16:]))
								return nil
//line /usr/local/go/src/net/netip/netip.go:1020
		// _ = "end of CoverTab[12340]"
//line /usr/local/go/src/net/netip/netip.go:1020
	default:
//line /usr/local/go/src/net/netip/netip.go:1020
		_go_fuzz_dep_.CoverTab[12341]++
//line /usr/local/go/src/net/netip/netip.go:1020
		// _ = "end of CoverTab[12341]"
	}
//line /usr/local/go/src/net/netip/netip.go:1021
	// _ = "end of CoverTab[12335]"
//line /usr/local/go/src/net/netip/netip.go:1021
	_go_fuzz_dep_.CoverTab[12336]++
							return errors.New("unexpected slice size")
//line /usr/local/go/src/net/netip/netip.go:1022
	// _ = "end of CoverTab[12336]"
}

// AddrPort is an IP and a port number.
type AddrPort struct {
	ip	Addr
	port	uint16
}

// AddrPortFrom returns an AddrPort with the provided IP and port.
//line /usr/local/go/src/net/netip/netip.go:1031
// It does not allocate.
//line /usr/local/go/src/net/netip/netip.go:1033
func AddrPortFrom(ip Addr, port uint16) AddrPort {
//line /usr/local/go/src/net/netip/netip.go:1033
	_go_fuzz_dep_.CoverTab[12342]++
//line /usr/local/go/src/net/netip/netip.go:1033
	return AddrPort{ip: ip, port: port}
//line /usr/local/go/src/net/netip/netip.go:1033
	// _ = "end of CoverTab[12342]"
//line /usr/local/go/src/net/netip/netip.go:1033
}

// Addr returns p's IP address.
func (p AddrPort) Addr() Addr {
//line /usr/local/go/src/net/netip/netip.go:1036
	_go_fuzz_dep_.CoverTab[12343]++
//line /usr/local/go/src/net/netip/netip.go:1036
	return p.ip
//line /usr/local/go/src/net/netip/netip.go:1036
	// _ = "end of CoverTab[12343]"
//line /usr/local/go/src/net/netip/netip.go:1036
}

// Port returns p's port.
func (p AddrPort) Port() uint16 {
//line /usr/local/go/src/net/netip/netip.go:1039
	_go_fuzz_dep_.CoverTab[12344]++
//line /usr/local/go/src/net/netip/netip.go:1039
	return p.port
//line /usr/local/go/src/net/netip/netip.go:1039
	// _ = "end of CoverTab[12344]"
//line /usr/local/go/src/net/netip/netip.go:1039
}

// splitAddrPort splits s into an IP address string and a port
//line /usr/local/go/src/net/netip/netip.go:1041
// string. It splits strings shaped like "foo:bar" or "[foo]:bar",
//line /usr/local/go/src/net/netip/netip.go:1041
// without further validating the substrings. v6 indicates whether the
//line /usr/local/go/src/net/netip/netip.go:1041
// ip string should parse as an IPv6 address or an IPv4 address, in
//line /usr/local/go/src/net/netip/netip.go:1041
// order for s to be a valid ip:port string.
//line /usr/local/go/src/net/netip/netip.go:1046
func splitAddrPort(s string) (ip, port string, v6 bool, err error) {
//line /usr/local/go/src/net/netip/netip.go:1046
	_go_fuzz_dep_.CoverTab[12345]++
							i := stringsLastIndexByte(s, ':')
							if i == -1 {
//line /usr/local/go/src/net/netip/netip.go:1048
		_go_fuzz_dep_.CoverTab[12350]++
								return "", "", false, errors.New("not an ip:port")
//line /usr/local/go/src/net/netip/netip.go:1049
		// _ = "end of CoverTab[12350]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1050
		_go_fuzz_dep_.CoverTab[12351]++
//line /usr/local/go/src/net/netip/netip.go:1050
		// _ = "end of CoverTab[12351]"
//line /usr/local/go/src/net/netip/netip.go:1050
	}
//line /usr/local/go/src/net/netip/netip.go:1050
	// _ = "end of CoverTab[12345]"
//line /usr/local/go/src/net/netip/netip.go:1050
	_go_fuzz_dep_.CoverTab[12346]++

							ip, port = s[:i], s[i+1:]
							if len(ip) == 0 {
//line /usr/local/go/src/net/netip/netip.go:1053
		_go_fuzz_dep_.CoverTab[12352]++
								return "", "", false, errors.New("no IP")
//line /usr/local/go/src/net/netip/netip.go:1054
		// _ = "end of CoverTab[12352]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1055
		_go_fuzz_dep_.CoverTab[12353]++
//line /usr/local/go/src/net/netip/netip.go:1055
		// _ = "end of CoverTab[12353]"
//line /usr/local/go/src/net/netip/netip.go:1055
	}
//line /usr/local/go/src/net/netip/netip.go:1055
	// _ = "end of CoverTab[12346]"
//line /usr/local/go/src/net/netip/netip.go:1055
	_go_fuzz_dep_.CoverTab[12347]++
							if len(port) == 0 {
//line /usr/local/go/src/net/netip/netip.go:1056
		_go_fuzz_dep_.CoverTab[12354]++
								return "", "", false, errors.New("no port")
//line /usr/local/go/src/net/netip/netip.go:1057
		// _ = "end of CoverTab[12354]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1058
		_go_fuzz_dep_.CoverTab[12355]++
//line /usr/local/go/src/net/netip/netip.go:1058
		// _ = "end of CoverTab[12355]"
//line /usr/local/go/src/net/netip/netip.go:1058
	}
//line /usr/local/go/src/net/netip/netip.go:1058
	// _ = "end of CoverTab[12347]"
//line /usr/local/go/src/net/netip/netip.go:1058
	_go_fuzz_dep_.CoverTab[12348]++
							if ip[0] == '[' {
//line /usr/local/go/src/net/netip/netip.go:1059
		_go_fuzz_dep_.CoverTab[12356]++
								if len(ip) < 2 || func() bool {
//line /usr/local/go/src/net/netip/netip.go:1060
			_go_fuzz_dep_.CoverTab[12358]++
//line /usr/local/go/src/net/netip/netip.go:1060
			return ip[len(ip)-1] != ']'
//line /usr/local/go/src/net/netip/netip.go:1060
			// _ = "end of CoverTab[12358]"
//line /usr/local/go/src/net/netip/netip.go:1060
		}() {
//line /usr/local/go/src/net/netip/netip.go:1060
			_go_fuzz_dep_.CoverTab[12359]++
									return "", "", false, errors.New("missing ]")
//line /usr/local/go/src/net/netip/netip.go:1061
			// _ = "end of CoverTab[12359]"
		} else {
//line /usr/local/go/src/net/netip/netip.go:1062
			_go_fuzz_dep_.CoverTab[12360]++
//line /usr/local/go/src/net/netip/netip.go:1062
			// _ = "end of CoverTab[12360]"
//line /usr/local/go/src/net/netip/netip.go:1062
		}
//line /usr/local/go/src/net/netip/netip.go:1062
		// _ = "end of CoverTab[12356]"
//line /usr/local/go/src/net/netip/netip.go:1062
		_go_fuzz_dep_.CoverTab[12357]++
								ip = ip[1 : len(ip)-1]
								v6 = true
//line /usr/local/go/src/net/netip/netip.go:1064
		// _ = "end of CoverTab[12357]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1065
		_go_fuzz_dep_.CoverTab[12361]++
//line /usr/local/go/src/net/netip/netip.go:1065
		// _ = "end of CoverTab[12361]"
//line /usr/local/go/src/net/netip/netip.go:1065
	}
//line /usr/local/go/src/net/netip/netip.go:1065
	// _ = "end of CoverTab[12348]"
//line /usr/local/go/src/net/netip/netip.go:1065
	_go_fuzz_dep_.CoverTab[12349]++

							return ip, port, v6, nil
//line /usr/local/go/src/net/netip/netip.go:1067
	// _ = "end of CoverTab[12349]"
}

// ParseAddrPort parses s as an AddrPort.
//line /usr/local/go/src/net/netip/netip.go:1070
//
//line /usr/local/go/src/net/netip/netip.go:1070
// It doesn't do any name resolution: both the address and the port
//line /usr/local/go/src/net/netip/netip.go:1070
// must be numeric.
//line /usr/local/go/src/net/netip/netip.go:1074
func ParseAddrPort(s string) (AddrPort, error) {
//line /usr/local/go/src/net/netip/netip.go:1074
	_go_fuzz_dep_.CoverTab[12362]++
							var ipp AddrPort
							ip, port, v6, err := splitAddrPort(s)
							if err != nil {
//line /usr/local/go/src/net/netip/netip.go:1077
		_go_fuzz_dep_.CoverTab[12367]++
								return ipp, err
//line /usr/local/go/src/net/netip/netip.go:1078
		// _ = "end of CoverTab[12367]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1079
		_go_fuzz_dep_.CoverTab[12368]++
//line /usr/local/go/src/net/netip/netip.go:1079
		// _ = "end of CoverTab[12368]"
//line /usr/local/go/src/net/netip/netip.go:1079
	}
//line /usr/local/go/src/net/netip/netip.go:1079
	// _ = "end of CoverTab[12362]"
//line /usr/local/go/src/net/netip/netip.go:1079
	_go_fuzz_dep_.CoverTab[12363]++
							port16, err := strconv.ParseUint(port, 10, 16)
							if err != nil {
//line /usr/local/go/src/net/netip/netip.go:1081
		_go_fuzz_dep_.CoverTab[12369]++
								return ipp, errors.New("invalid port " + strconv.Quote(port) + " parsing " + strconv.Quote(s))
//line /usr/local/go/src/net/netip/netip.go:1082
		// _ = "end of CoverTab[12369]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1083
		_go_fuzz_dep_.CoverTab[12370]++
//line /usr/local/go/src/net/netip/netip.go:1083
		// _ = "end of CoverTab[12370]"
//line /usr/local/go/src/net/netip/netip.go:1083
	}
//line /usr/local/go/src/net/netip/netip.go:1083
	// _ = "end of CoverTab[12363]"
//line /usr/local/go/src/net/netip/netip.go:1083
	_go_fuzz_dep_.CoverTab[12364]++
							ipp.port = uint16(port16)
							ipp.ip, err = ParseAddr(ip)
							if err != nil {
//line /usr/local/go/src/net/netip/netip.go:1086
		_go_fuzz_dep_.CoverTab[12371]++
								return AddrPort{}, err
//line /usr/local/go/src/net/netip/netip.go:1087
		// _ = "end of CoverTab[12371]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1088
		_go_fuzz_dep_.CoverTab[12372]++
//line /usr/local/go/src/net/netip/netip.go:1088
		// _ = "end of CoverTab[12372]"
//line /usr/local/go/src/net/netip/netip.go:1088
	}
//line /usr/local/go/src/net/netip/netip.go:1088
	// _ = "end of CoverTab[12364]"
//line /usr/local/go/src/net/netip/netip.go:1088
	_go_fuzz_dep_.CoverTab[12365]++
							if v6 && func() bool {
//line /usr/local/go/src/net/netip/netip.go:1089
		_go_fuzz_dep_.CoverTab[12373]++
//line /usr/local/go/src/net/netip/netip.go:1089
		return ipp.ip.Is4()
//line /usr/local/go/src/net/netip/netip.go:1089
		// _ = "end of CoverTab[12373]"
//line /usr/local/go/src/net/netip/netip.go:1089
	}() {
//line /usr/local/go/src/net/netip/netip.go:1089
		_go_fuzz_dep_.CoverTab[12374]++
								return AddrPort{}, errors.New("invalid ip:port " + strconv.Quote(s) + ", square brackets can only be used with IPv6 addresses")
//line /usr/local/go/src/net/netip/netip.go:1090
		// _ = "end of CoverTab[12374]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1091
		_go_fuzz_dep_.CoverTab[12375]++
//line /usr/local/go/src/net/netip/netip.go:1091
		if !v6 && func() bool {
//line /usr/local/go/src/net/netip/netip.go:1091
			_go_fuzz_dep_.CoverTab[12376]++
//line /usr/local/go/src/net/netip/netip.go:1091
			return ipp.ip.Is6()
//line /usr/local/go/src/net/netip/netip.go:1091
			// _ = "end of CoverTab[12376]"
//line /usr/local/go/src/net/netip/netip.go:1091
		}() {
//line /usr/local/go/src/net/netip/netip.go:1091
			_go_fuzz_dep_.CoverTab[12377]++
									return AddrPort{}, errors.New("invalid ip:port " + strconv.Quote(s) + ", IPv6 addresses must be surrounded by square brackets")
//line /usr/local/go/src/net/netip/netip.go:1092
			// _ = "end of CoverTab[12377]"
		} else {
//line /usr/local/go/src/net/netip/netip.go:1093
			_go_fuzz_dep_.CoverTab[12378]++
//line /usr/local/go/src/net/netip/netip.go:1093
			// _ = "end of CoverTab[12378]"
//line /usr/local/go/src/net/netip/netip.go:1093
		}
//line /usr/local/go/src/net/netip/netip.go:1093
		// _ = "end of CoverTab[12375]"
//line /usr/local/go/src/net/netip/netip.go:1093
	}
//line /usr/local/go/src/net/netip/netip.go:1093
	// _ = "end of CoverTab[12365]"
//line /usr/local/go/src/net/netip/netip.go:1093
	_go_fuzz_dep_.CoverTab[12366]++
							return ipp, nil
//line /usr/local/go/src/net/netip/netip.go:1094
	// _ = "end of CoverTab[12366]"
}

// MustParseAddrPort calls ParseAddrPort(s) and panics on error.
//line /usr/local/go/src/net/netip/netip.go:1097
// It is intended for use in tests with hard-coded strings.
//line /usr/local/go/src/net/netip/netip.go:1099
func MustParseAddrPort(s string) AddrPort {
//line /usr/local/go/src/net/netip/netip.go:1099
	_go_fuzz_dep_.CoverTab[12379]++
							ip, err := ParseAddrPort(s)
							if err != nil {
//line /usr/local/go/src/net/netip/netip.go:1101
		_go_fuzz_dep_.CoverTab[12381]++
								panic(err)
//line /usr/local/go/src/net/netip/netip.go:1102
		// _ = "end of CoverTab[12381]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1103
		_go_fuzz_dep_.CoverTab[12382]++
//line /usr/local/go/src/net/netip/netip.go:1103
		// _ = "end of CoverTab[12382]"
//line /usr/local/go/src/net/netip/netip.go:1103
	}
//line /usr/local/go/src/net/netip/netip.go:1103
	// _ = "end of CoverTab[12379]"
//line /usr/local/go/src/net/netip/netip.go:1103
	_go_fuzz_dep_.CoverTab[12380]++
							return ip
//line /usr/local/go/src/net/netip/netip.go:1104
	// _ = "end of CoverTab[12380]"
}

// IsValid reports whether p.Addr() is valid.
//line /usr/local/go/src/net/netip/netip.go:1107
// All ports are valid, including zero.
//line /usr/local/go/src/net/netip/netip.go:1109
func (p AddrPort) IsValid() bool {
//line /usr/local/go/src/net/netip/netip.go:1109
	_go_fuzz_dep_.CoverTab[12383]++
//line /usr/local/go/src/net/netip/netip.go:1109
	return p.ip.IsValid()
//line /usr/local/go/src/net/netip/netip.go:1109
	// _ = "end of CoverTab[12383]"
//line /usr/local/go/src/net/netip/netip.go:1109
}

func (p AddrPort) String() string {
//line /usr/local/go/src/net/netip/netip.go:1111
	_go_fuzz_dep_.CoverTab[12384]++
							switch p.ip.z {
	case z0:
//line /usr/local/go/src/net/netip/netip.go:1113
		_go_fuzz_dep_.CoverTab[12385]++
								return "invalid AddrPort"
//line /usr/local/go/src/net/netip/netip.go:1114
		// _ = "end of CoverTab[12385]"
	case z4:
//line /usr/local/go/src/net/netip/netip.go:1115
		_go_fuzz_dep_.CoverTab[12386]++
								a := p.ip.As4()
								buf := make([]byte, 0, 21)
								for i := range a {
//line /usr/local/go/src/net/netip/netip.go:1118
			_go_fuzz_dep_.CoverTab[12389]++
									buf = strconv.AppendUint(buf, uint64(a[i]), 10)
									buf = append(buf, "...:"[i])
//line /usr/local/go/src/net/netip/netip.go:1120
			// _ = "end of CoverTab[12389]"
		}
//line /usr/local/go/src/net/netip/netip.go:1121
		// _ = "end of CoverTab[12386]"
//line /usr/local/go/src/net/netip/netip.go:1121
		_go_fuzz_dep_.CoverTab[12387]++
								buf = strconv.AppendUint(buf, uint64(p.port), 10)
								return string(buf)
//line /usr/local/go/src/net/netip/netip.go:1123
		// _ = "end of CoverTab[12387]"
	default:
//line /usr/local/go/src/net/netip/netip.go:1124
		_go_fuzz_dep_.CoverTab[12388]++

								return joinHostPort(p.ip.String(), itoa.Itoa(int(p.port)))
//line /usr/local/go/src/net/netip/netip.go:1126
		// _ = "end of CoverTab[12388]"
	}
//line /usr/local/go/src/net/netip/netip.go:1127
	// _ = "end of CoverTab[12384]"
}

func joinHostPort(host, port string) string {
//line /usr/local/go/src/net/netip/netip.go:1130
	_go_fuzz_dep_.CoverTab[12390]++

//line /usr/local/go/src/net/netip/netip.go:1133
	if bytealg.IndexByteString(host, ':') >= 0 {
//line /usr/local/go/src/net/netip/netip.go:1133
		_go_fuzz_dep_.CoverTab[12392]++
								return "[" + host + "]:" + port
//line /usr/local/go/src/net/netip/netip.go:1134
		// _ = "end of CoverTab[12392]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1135
		_go_fuzz_dep_.CoverTab[12393]++
//line /usr/local/go/src/net/netip/netip.go:1135
		// _ = "end of CoverTab[12393]"
//line /usr/local/go/src/net/netip/netip.go:1135
	}
//line /usr/local/go/src/net/netip/netip.go:1135
	// _ = "end of CoverTab[12390]"
//line /usr/local/go/src/net/netip/netip.go:1135
	_go_fuzz_dep_.CoverTab[12391]++
							return host + ":" + port
//line /usr/local/go/src/net/netip/netip.go:1136
	// _ = "end of CoverTab[12391]"
}

// AppendTo appends a text encoding of p,
//line /usr/local/go/src/net/netip/netip.go:1139
// as generated by MarshalText,
//line /usr/local/go/src/net/netip/netip.go:1139
// to b and returns the extended buffer.
//line /usr/local/go/src/net/netip/netip.go:1142
func (p AddrPort) AppendTo(b []byte) []byte {
//line /usr/local/go/src/net/netip/netip.go:1142
	_go_fuzz_dep_.CoverTab[12394]++
							switch p.ip.z {
	case z0:
//line /usr/local/go/src/net/netip/netip.go:1144
		_go_fuzz_dep_.CoverTab[12396]++
								return b
//line /usr/local/go/src/net/netip/netip.go:1145
		// _ = "end of CoverTab[12396]"
	case z4:
//line /usr/local/go/src/net/netip/netip.go:1146
		_go_fuzz_dep_.CoverTab[12397]++
								b = p.ip.appendTo4(b)
//line /usr/local/go/src/net/netip/netip.go:1147
		// _ = "end of CoverTab[12397]"
	default:
//line /usr/local/go/src/net/netip/netip.go:1148
		_go_fuzz_dep_.CoverTab[12398]++
								if p.ip.Is4In6() {
//line /usr/local/go/src/net/netip/netip.go:1149
			_go_fuzz_dep_.CoverTab[12400]++
									b = append(b, "[::ffff:"...)
									b = p.ip.Unmap().appendTo4(b)
									if z := p.ip.Zone(); z != "" {
//line /usr/local/go/src/net/netip/netip.go:1152
				_go_fuzz_dep_.CoverTab[12401]++
										b = append(b, '%')
										b = append(b, z...)
//line /usr/local/go/src/net/netip/netip.go:1154
				// _ = "end of CoverTab[12401]"
			} else {
//line /usr/local/go/src/net/netip/netip.go:1155
				_go_fuzz_dep_.CoverTab[12402]++
//line /usr/local/go/src/net/netip/netip.go:1155
				// _ = "end of CoverTab[12402]"
//line /usr/local/go/src/net/netip/netip.go:1155
			}
//line /usr/local/go/src/net/netip/netip.go:1155
			// _ = "end of CoverTab[12400]"
		} else {
//line /usr/local/go/src/net/netip/netip.go:1156
			_go_fuzz_dep_.CoverTab[12403]++
									b = append(b, '[')
									b = p.ip.appendTo6(b)
//line /usr/local/go/src/net/netip/netip.go:1158
			// _ = "end of CoverTab[12403]"
		}
//line /usr/local/go/src/net/netip/netip.go:1159
		// _ = "end of CoverTab[12398]"
//line /usr/local/go/src/net/netip/netip.go:1159
		_go_fuzz_dep_.CoverTab[12399]++
								b = append(b, ']')
//line /usr/local/go/src/net/netip/netip.go:1160
		// _ = "end of CoverTab[12399]"
	}
//line /usr/local/go/src/net/netip/netip.go:1161
	// _ = "end of CoverTab[12394]"
//line /usr/local/go/src/net/netip/netip.go:1161
	_go_fuzz_dep_.CoverTab[12395]++
							b = append(b, ':')
							b = strconv.AppendUint(b, uint64(p.port), 10)
							return b
//line /usr/local/go/src/net/netip/netip.go:1164
	// _ = "end of CoverTab[12395]"
}

// MarshalText implements the encoding.TextMarshaler interface. The
//line /usr/local/go/src/net/netip/netip.go:1167
// encoding is the same as returned by String, with one exception: if
//line /usr/local/go/src/net/netip/netip.go:1167
// p.Addr() is the zero Addr, the encoding is the empty string.
//line /usr/local/go/src/net/netip/netip.go:1170
func (p AddrPort) MarshalText() ([]byte, error) {
//line /usr/local/go/src/net/netip/netip.go:1170
	_go_fuzz_dep_.CoverTab[12404]++
							var max int
							switch p.ip.z {
	case z0:
//line /usr/local/go/src/net/netip/netip.go:1173
		_go_fuzz_dep_.CoverTab[12406]++
//line /usr/local/go/src/net/netip/netip.go:1173
		// _ = "end of CoverTab[12406]"
	case z4:
//line /usr/local/go/src/net/netip/netip.go:1174
		_go_fuzz_dep_.CoverTab[12407]++
								max = len("255.255.255.255:65535")
//line /usr/local/go/src/net/netip/netip.go:1175
		// _ = "end of CoverTab[12407]"
	default:
//line /usr/local/go/src/net/netip/netip.go:1176
		_go_fuzz_dep_.CoverTab[12408]++
								max = len("[ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff%enp5s0]:65535")
//line /usr/local/go/src/net/netip/netip.go:1177
		// _ = "end of CoverTab[12408]"
	}
//line /usr/local/go/src/net/netip/netip.go:1178
	// _ = "end of CoverTab[12404]"
//line /usr/local/go/src/net/netip/netip.go:1178
	_go_fuzz_dep_.CoverTab[12405]++
							b := make([]byte, 0, max)
							b = p.AppendTo(b)
							return b, nil
//line /usr/local/go/src/net/netip/netip.go:1181
	// _ = "end of CoverTab[12405]"
}

// UnmarshalText implements the encoding.TextUnmarshaler
//line /usr/local/go/src/net/netip/netip.go:1184
// interface. The AddrPort is expected in a form
//line /usr/local/go/src/net/netip/netip.go:1184
// generated by MarshalText or accepted by ParseAddrPort.
//line /usr/local/go/src/net/netip/netip.go:1187
func (p *AddrPort) UnmarshalText(text []byte) error {
//line /usr/local/go/src/net/netip/netip.go:1187
	_go_fuzz_dep_.CoverTab[12409]++
							if len(text) == 0 {
//line /usr/local/go/src/net/netip/netip.go:1188
		_go_fuzz_dep_.CoverTab[12411]++
								*p = AddrPort{}
								return nil
//line /usr/local/go/src/net/netip/netip.go:1190
		// _ = "end of CoverTab[12411]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1191
		_go_fuzz_dep_.CoverTab[12412]++
//line /usr/local/go/src/net/netip/netip.go:1191
		// _ = "end of CoverTab[12412]"
//line /usr/local/go/src/net/netip/netip.go:1191
	}
//line /usr/local/go/src/net/netip/netip.go:1191
	// _ = "end of CoverTab[12409]"
//line /usr/local/go/src/net/netip/netip.go:1191
	_go_fuzz_dep_.CoverTab[12410]++
							var err error
							*p, err = ParseAddrPort(string(text))
							return err
//line /usr/local/go/src/net/netip/netip.go:1194
	// _ = "end of CoverTab[12410]"
}

// MarshalBinary implements the encoding.BinaryMarshaler interface.
//line /usr/local/go/src/net/netip/netip.go:1197
// It returns Addr.MarshalBinary with an additional two bytes appended
//line /usr/local/go/src/net/netip/netip.go:1197
// containing the port in little-endian.
//line /usr/local/go/src/net/netip/netip.go:1200
func (p AddrPort) MarshalBinary() ([]byte, error) {
//line /usr/local/go/src/net/netip/netip.go:1200
	_go_fuzz_dep_.CoverTab[12413]++
							b := p.Addr().marshalBinaryWithTrailingBytes(2)
							lePutUint16(b[len(b)-2:], p.Port())
							return b, nil
//line /usr/local/go/src/net/netip/netip.go:1203
	// _ = "end of CoverTab[12413]"
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface.
//line /usr/local/go/src/net/netip/netip.go:1206
// It expects data in the form generated by MarshalBinary.
//line /usr/local/go/src/net/netip/netip.go:1208
func (p *AddrPort) UnmarshalBinary(b []byte) error {
//line /usr/local/go/src/net/netip/netip.go:1208
	_go_fuzz_dep_.CoverTab[12414]++
							if len(b) < 2 {
//line /usr/local/go/src/net/netip/netip.go:1209
		_go_fuzz_dep_.CoverTab[12417]++
								return errors.New("unexpected slice size")
//line /usr/local/go/src/net/netip/netip.go:1210
		// _ = "end of CoverTab[12417]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1211
		_go_fuzz_dep_.CoverTab[12418]++
//line /usr/local/go/src/net/netip/netip.go:1211
		// _ = "end of CoverTab[12418]"
//line /usr/local/go/src/net/netip/netip.go:1211
	}
//line /usr/local/go/src/net/netip/netip.go:1211
	// _ = "end of CoverTab[12414]"
//line /usr/local/go/src/net/netip/netip.go:1211
	_go_fuzz_dep_.CoverTab[12415]++
							var addr Addr
							err := addr.UnmarshalBinary(b[:len(b)-2])
							if err != nil {
//line /usr/local/go/src/net/netip/netip.go:1214
		_go_fuzz_dep_.CoverTab[12419]++
								return err
//line /usr/local/go/src/net/netip/netip.go:1215
		// _ = "end of CoverTab[12419]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1216
		_go_fuzz_dep_.CoverTab[12420]++
//line /usr/local/go/src/net/netip/netip.go:1216
		// _ = "end of CoverTab[12420]"
//line /usr/local/go/src/net/netip/netip.go:1216
	}
//line /usr/local/go/src/net/netip/netip.go:1216
	// _ = "end of CoverTab[12415]"
//line /usr/local/go/src/net/netip/netip.go:1216
	_go_fuzz_dep_.CoverTab[12416]++
							*p = AddrPortFrom(addr, leUint16(b[len(b)-2:]))
							return nil
//line /usr/local/go/src/net/netip/netip.go:1218
	// _ = "end of CoverTab[12416]"
}

// Prefix is an IP address prefix (CIDR) representing an IP network.
//line /usr/local/go/src/net/netip/netip.go:1221
//
//line /usr/local/go/src/net/netip/netip.go:1221
// The first Bits() of Addr() are specified. The remaining bits match any address.
//line /usr/local/go/src/net/netip/netip.go:1221
// The range of Bits() is [0,32] for IPv4 or [0,128] for IPv6.
//line /usr/local/go/src/net/netip/netip.go:1225
type Prefix struct {
	ip	Addr

	// bits is logically a uint8 (storing [0,128]) but also
	// encodes an "invalid" bit, currently represented by the
	// invalidPrefixBits sentinel value. It could be packed into
	// the uint8 more with more complicated expressions in the
	// accessors, but the extra byte (in padding anyway) doesn't
	// hurt and simplifies code below.
	bits	int16
}

// invalidPrefixBits is the Prefix.bits value used when PrefixFrom is
//line /usr/local/go/src/net/netip/netip.go:1237
// outside the range of a uint8. It's returned as the int -1 in the
//line /usr/local/go/src/net/netip/netip.go:1237
// public API.
//line /usr/local/go/src/net/netip/netip.go:1240
const invalidPrefixBits = -1

// PrefixFrom returns a Prefix with the provided IP address and bit
//line /usr/local/go/src/net/netip/netip.go:1242
// prefix length.
//line /usr/local/go/src/net/netip/netip.go:1242
//
//line /usr/local/go/src/net/netip/netip.go:1242
// It does not allocate. Unlike Addr.Prefix, PrefixFrom does not mask
//line /usr/local/go/src/net/netip/netip.go:1242
// off the host bits of ip.
//line /usr/local/go/src/net/netip/netip.go:1242
//
//line /usr/local/go/src/net/netip/netip.go:1242
// If bits is less than zero or greater than ip.BitLen, Prefix.Bits
//line /usr/local/go/src/net/netip/netip.go:1242
// will return an invalid value -1.
//line /usr/local/go/src/net/netip/netip.go:1250
func PrefixFrom(ip Addr, bits int) Prefix {
//line /usr/local/go/src/net/netip/netip.go:1250
	_go_fuzz_dep_.CoverTab[12421]++
							if bits < 0 || func() bool {
//line /usr/local/go/src/net/netip/netip.go:1251
		_go_fuzz_dep_.CoverTab[12423]++
//line /usr/local/go/src/net/netip/netip.go:1251
		return bits > ip.BitLen()
//line /usr/local/go/src/net/netip/netip.go:1251
		// _ = "end of CoverTab[12423]"
//line /usr/local/go/src/net/netip/netip.go:1251
	}() {
//line /usr/local/go/src/net/netip/netip.go:1251
		_go_fuzz_dep_.CoverTab[12424]++
								bits = invalidPrefixBits
//line /usr/local/go/src/net/netip/netip.go:1252
		// _ = "end of CoverTab[12424]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1253
		_go_fuzz_dep_.CoverTab[12425]++
//line /usr/local/go/src/net/netip/netip.go:1253
		// _ = "end of CoverTab[12425]"
//line /usr/local/go/src/net/netip/netip.go:1253
	}
//line /usr/local/go/src/net/netip/netip.go:1253
	// _ = "end of CoverTab[12421]"
//line /usr/local/go/src/net/netip/netip.go:1253
	_go_fuzz_dep_.CoverTab[12422]++
							b16 := int16(bits)
							return Prefix{
		ip:	ip.withoutZone(),
		bits:	b16,
	}
//line /usr/local/go/src/net/netip/netip.go:1258
	// _ = "end of CoverTab[12422]"
}

// Addr returns p's IP address.
func (p Prefix) Addr() Addr {
//line /usr/local/go/src/net/netip/netip.go:1262
	_go_fuzz_dep_.CoverTab[12426]++
//line /usr/local/go/src/net/netip/netip.go:1262
	return p.ip
//line /usr/local/go/src/net/netip/netip.go:1262
	// _ = "end of CoverTab[12426]"
//line /usr/local/go/src/net/netip/netip.go:1262
}

// Bits returns p's prefix length.
//line /usr/local/go/src/net/netip/netip.go:1264
//
//line /usr/local/go/src/net/netip/netip.go:1264
// It reports -1 if invalid.
//line /usr/local/go/src/net/netip/netip.go:1267
func (p Prefix) Bits() int {
//line /usr/local/go/src/net/netip/netip.go:1267
	_go_fuzz_dep_.CoverTab[12427]++
//line /usr/local/go/src/net/netip/netip.go:1267
	return int(p.bits)
//line /usr/local/go/src/net/netip/netip.go:1267
	// _ = "end of CoverTab[12427]"
//line /usr/local/go/src/net/netip/netip.go:1267
}

// IsValid reports whether p.Bits() has a valid range for p.Addr().
//line /usr/local/go/src/net/netip/netip.go:1269
// If p.Addr() is the zero Addr, IsValid returns false.
//line /usr/local/go/src/net/netip/netip.go:1269
// Note that if p is the zero Prefix, then p.IsValid() == false.
//line /usr/local/go/src/net/netip/netip.go:1272
func (p Prefix) IsValid() bool {
//line /usr/local/go/src/net/netip/netip.go:1272
	_go_fuzz_dep_.CoverTab[12428]++
//line /usr/local/go/src/net/netip/netip.go:1272
	return !p.ip.isZero() && func() bool {
//line /usr/local/go/src/net/netip/netip.go:1272
		_go_fuzz_dep_.CoverTab[12429]++
//line /usr/local/go/src/net/netip/netip.go:1272
		return p.bits >= 0
//line /usr/local/go/src/net/netip/netip.go:1272
		// _ = "end of CoverTab[12429]"
//line /usr/local/go/src/net/netip/netip.go:1272
	}() && func() bool {
//line /usr/local/go/src/net/netip/netip.go:1272
		_go_fuzz_dep_.CoverTab[12430]++
//line /usr/local/go/src/net/netip/netip.go:1272
		return int(p.bits) <= p.ip.BitLen()
//line /usr/local/go/src/net/netip/netip.go:1272
		// _ = "end of CoverTab[12430]"
//line /usr/local/go/src/net/netip/netip.go:1272
	}()
//line /usr/local/go/src/net/netip/netip.go:1272
	// _ = "end of CoverTab[12428]"
//line /usr/local/go/src/net/netip/netip.go:1272
}

func (p Prefix) isZero() bool {
//line /usr/local/go/src/net/netip/netip.go:1274
	_go_fuzz_dep_.CoverTab[12431]++
//line /usr/local/go/src/net/netip/netip.go:1274
	return p == Prefix{}
//line /usr/local/go/src/net/netip/netip.go:1274
	// _ = "end of CoverTab[12431]"
//line /usr/local/go/src/net/netip/netip.go:1274
}

// IsSingleIP reports whether p contains exactly one IP.
func (p Prefix) IsSingleIP() bool {
//line /usr/local/go/src/net/netip/netip.go:1277
	_go_fuzz_dep_.CoverTab[12432]++
//line /usr/local/go/src/net/netip/netip.go:1277
	return p.bits != 0 && func() bool {
//line /usr/local/go/src/net/netip/netip.go:1277
		_go_fuzz_dep_.CoverTab[12433]++
//line /usr/local/go/src/net/netip/netip.go:1277
		return int(p.bits) == p.ip.BitLen()
//line /usr/local/go/src/net/netip/netip.go:1277
		// _ = "end of CoverTab[12433]"
//line /usr/local/go/src/net/netip/netip.go:1277
	}()
//line /usr/local/go/src/net/netip/netip.go:1277
	// _ = "end of CoverTab[12432]"
//line /usr/local/go/src/net/netip/netip.go:1277
}

// ParsePrefix parses s as an IP address prefix.
//line /usr/local/go/src/net/netip/netip.go:1279
// The string can be in the form "192.168.1.0/24" or "2001:db8::/32",
//line /usr/local/go/src/net/netip/netip.go:1279
// the CIDR notation defined in RFC 4632 and RFC 4291.
//line /usr/local/go/src/net/netip/netip.go:1279
// IPv6 zones are not permitted in prefixes, and an error will be returned if a
//line /usr/local/go/src/net/netip/netip.go:1279
// zone is present.
//line /usr/local/go/src/net/netip/netip.go:1279
//
//line /usr/local/go/src/net/netip/netip.go:1279
// Note that masked address bits are not zeroed. Use Masked for that.
//line /usr/local/go/src/net/netip/netip.go:1286
func ParsePrefix(s string) (Prefix, error) {
//line /usr/local/go/src/net/netip/netip.go:1286
	_go_fuzz_dep_.CoverTab[12434]++
							i := stringsLastIndexByte(s, '/')
							if i < 0 {
//line /usr/local/go/src/net/netip/netip.go:1288
		_go_fuzz_dep_.CoverTab[12441]++
								return Prefix{}, errors.New("netip.ParsePrefix(" + strconv.Quote(s) + "): no '/'")
//line /usr/local/go/src/net/netip/netip.go:1289
		// _ = "end of CoverTab[12441]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1290
		_go_fuzz_dep_.CoverTab[12442]++
//line /usr/local/go/src/net/netip/netip.go:1290
		// _ = "end of CoverTab[12442]"
//line /usr/local/go/src/net/netip/netip.go:1290
	}
//line /usr/local/go/src/net/netip/netip.go:1290
	// _ = "end of CoverTab[12434]"
//line /usr/local/go/src/net/netip/netip.go:1290
	_go_fuzz_dep_.CoverTab[12435]++
							ip, err := ParseAddr(s[:i])
							if err != nil {
//line /usr/local/go/src/net/netip/netip.go:1292
		_go_fuzz_dep_.CoverTab[12443]++
								return Prefix{}, errors.New("netip.ParsePrefix(" + strconv.Quote(s) + "): " + err.Error())
//line /usr/local/go/src/net/netip/netip.go:1293
		// _ = "end of CoverTab[12443]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1294
		_go_fuzz_dep_.CoverTab[12444]++
//line /usr/local/go/src/net/netip/netip.go:1294
		// _ = "end of CoverTab[12444]"
//line /usr/local/go/src/net/netip/netip.go:1294
	}
//line /usr/local/go/src/net/netip/netip.go:1294
	// _ = "end of CoverTab[12435]"
//line /usr/local/go/src/net/netip/netip.go:1294
	_go_fuzz_dep_.CoverTab[12436]++

							if ip.Is6() && func() bool {
//line /usr/local/go/src/net/netip/netip.go:1296
		_go_fuzz_dep_.CoverTab[12445]++
//line /usr/local/go/src/net/netip/netip.go:1296
		return ip.z != z6noz
//line /usr/local/go/src/net/netip/netip.go:1296
		// _ = "end of CoverTab[12445]"
//line /usr/local/go/src/net/netip/netip.go:1296
	}() {
//line /usr/local/go/src/net/netip/netip.go:1296
		_go_fuzz_dep_.CoverTab[12446]++
								return Prefix{}, errors.New("netip.ParsePrefix(" + strconv.Quote(s) + "): IPv6 zones cannot be present in a prefix")
//line /usr/local/go/src/net/netip/netip.go:1297
		// _ = "end of CoverTab[12446]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1298
		_go_fuzz_dep_.CoverTab[12447]++
//line /usr/local/go/src/net/netip/netip.go:1298
		// _ = "end of CoverTab[12447]"
//line /usr/local/go/src/net/netip/netip.go:1298
	}
//line /usr/local/go/src/net/netip/netip.go:1298
	// _ = "end of CoverTab[12436]"
//line /usr/local/go/src/net/netip/netip.go:1298
	_go_fuzz_dep_.CoverTab[12437]++

							bitsStr := s[i+1:]
							bits, err := strconv.Atoi(bitsStr)
							if err != nil {
//line /usr/local/go/src/net/netip/netip.go:1302
		_go_fuzz_dep_.CoverTab[12448]++
								return Prefix{}, errors.New("netip.ParsePrefix(" + strconv.Quote(s) + "): bad bits after slash: " + strconv.Quote(bitsStr))
//line /usr/local/go/src/net/netip/netip.go:1303
		// _ = "end of CoverTab[12448]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1304
		_go_fuzz_dep_.CoverTab[12449]++
//line /usr/local/go/src/net/netip/netip.go:1304
		// _ = "end of CoverTab[12449]"
//line /usr/local/go/src/net/netip/netip.go:1304
	}
//line /usr/local/go/src/net/netip/netip.go:1304
	// _ = "end of CoverTab[12437]"
//line /usr/local/go/src/net/netip/netip.go:1304
	_go_fuzz_dep_.CoverTab[12438]++
							maxBits := 32
							if ip.Is6() {
//line /usr/local/go/src/net/netip/netip.go:1306
		_go_fuzz_dep_.CoverTab[12450]++
								maxBits = 128
//line /usr/local/go/src/net/netip/netip.go:1307
		// _ = "end of CoverTab[12450]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1308
		_go_fuzz_dep_.CoverTab[12451]++
//line /usr/local/go/src/net/netip/netip.go:1308
		// _ = "end of CoverTab[12451]"
//line /usr/local/go/src/net/netip/netip.go:1308
	}
//line /usr/local/go/src/net/netip/netip.go:1308
	// _ = "end of CoverTab[12438]"
//line /usr/local/go/src/net/netip/netip.go:1308
	_go_fuzz_dep_.CoverTab[12439]++
							if bits < 0 || func() bool {
//line /usr/local/go/src/net/netip/netip.go:1309
		_go_fuzz_dep_.CoverTab[12452]++
//line /usr/local/go/src/net/netip/netip.go:1309
		return bits > maxBits
//line /usr/local/go/src/net/netip/netip.go:1309
		// _ = "end of CoverTab[12452]"
//line /usr/local/go/src/net/netip/netip.go:1309
	}() {
//line /usr/local/go/src/net/netip/netip.go:1309
		_go_fuzz_dep_.CoverTab[12453]++
								return Prefix{}, errors.New("netip.ParsePrefix(" + strconv.Quote(s) + "): prefix length out of range")
//line /usr/local/go/src/net/netip/netip.go:1310
		// _ = "end of CoverTab[12453]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1311
		_go_fuzz_dep_.CoverTab[12454]++
//line /usr/local/go/src/net/netip/netip.go:1311
		// _ = "end of CoverTab[12454]"
//line /usr/local/go/src/net/netip/netip.go:1311
	}
//line /usr/local/go/src/net/netip/netip.go:1311
	// _ = "end of CoverTab[12439]"
//line /usr/local/go/src/net/netip/netip.go:1311
	_go_fuzz_dep_.CoverTab[12440]++
							return PrefixFrom(ip, bits), nil
//line /usr/local/go/src/net/netip/netip.go:1312
	// _ = "end of CoverTab[12440]"
}

// MustParsePrefix calls ParsePrefix(s) and panics on error.
//line /usr/local/go/src/net/netip/netip.go:1315
// It is intended for use in tests with hard-coded strings.
//line /usr/local/go/src/net/netip/netip.go:1317
func MustParsePrefix(s string) Prefix {
//line /usr/local/go/src/net/netip/netip.go:1317
	_go_fuzz_dep_.CoverTab[12455]++
							ip, err := ParsePrefix(s)
							if err != nil {
//line /usr/local/go/src/net/netip/netip.go:1319
		_go_fuzz_dep_.CoverTab[12457]++
								panic(err)
//line /usr/local/go/src/net/netip/netip.go:1320
		// _ = "end of CoverTab[12457]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1321
		_go_fuzz_dep_.CoverTab[12458]++
//line /usr/local/go/src/net/netip/netip.go:1321
		// _ = "end of CoverTab[12458]"
//line /usr/local/go/src/net/netip/netip.go:1321
	}
//line /usr/local/go/src/net/netip/netip.go:1321
	// _ = "end of CoverTab[12455]"
//line /usr/local/go/src/net/netip/netip.go:1321
	_go_fuzz_dep_.CoverTab[12456]++
							return ip
//line /usr/local/go/src/net/netip/netip.go:1322
	// _ = "end of CoverTab[12456]"
}

// Masked returns p in its canonical form, with all but the high
//line /usr/local/go/src/net/netip/netip.go:1325
// p.Bits() bits of p.Addr() masked off.
//line /usr/local/go/src/net/netip/netip.go:1325
//
//line /usr/local/go/src/net/netip/netip.go:1325
// If p is zero or otherwise invalid, Masked returns the zero Prefix.
//line /usr/local/go/src/net/netip/netip.go:1329
func (p Prefix) Masked() Prefix {
//line /usr/local/go/src/net/netip/netip.go:1329
	_go_fuzz_dep_.CoverTab[12459]++
							if m, err := p.ip.Prefix(int(p.bits)); err == nil {
//line /usr/local/go/src/net/netip/netip.go:1330
		_go_fuzz_dep_.CoverTab[12461]++
								return m
//line /usr/local/go/src/net/netip/netip.go:1331
		// _ = "end of CoverTab[12461]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1332
		_go_fuzz_dep_.CoverTab[12462]++
//line /usr/local/go/src/net/netip/netip.go:1332
		// _ = "end of CoverTab[12462]"
//line /usr/local/go/src/net/netip/netip.go:1332
	}
//line /usr/local/go/src/net/netip/netip.go:1332
	// _ = "end of CoverTab[12459]"
//line /usr/local/go/src/net/netip/netip.go:1332
	_go_fuzz_dep_.CoverTab[12460]++
							return Prefix{}
//line /usr/local/go/src/net/netip/netip.go:1333
	// _ = "end of CoverTab[12460]"
}

// Contains reports whether the network p includes ip.
//line /usr/local/go/src/net/netip/netip.go:1336
//
//line /usr/local/go/src/net/netip/netip.go:1336
// An IPv4 address will not match an IPv6 prefix.
//line /usr/local/go/src/net/netip/netip.go:1336
// An IPv4-mapped IPv6 address will not match an IPv4 prefix.
//line /usr/local/go/src/net/netip/netip.go:1336
// A zero-value IP will not match any prefix.
//line /usr/local/go/src/net/netip/netip.go:1336
// If ip has an IPv6 zone, Contains returns false,
//line /usr/local/go/src/net/netip/netip.go:1336
// because Prefixes strip zones.
//line /usr/local/go/src/net/netip/netip.go:1343
func (p Prefix) Contains(ip Addr) bool {
//line /usr/local/go/src/net/netip/netip.go:1343
	_go_fuzz_dep_.CoverTab[12463]++
							if !p.IsValid() || func() bool {
//line /usr/local/go/src/net/netip/netip.go:1344
		_go_fuzz_dep_.CoverTab[12466]++
//line /usr/local/go/src/net/netip/netip.go:1344
		return ip.hasZone()
//line /usr/local/go/src/net/netip/netip.go:1344
		// _ = "end of CoverTab[12466]"
//line /usr/local/go/src/net/netip/netip.go:1344
	}() {
//line /usr/local/go/src/net/netip/netip.go:1344
		_go_fuzz_dep_.CoverTab[12467]++
								return false
//line /usr/local/go/src/net/netip/netip.go:1345
		// _ = "end of CoverTab[12467]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1346
		_go_fuzz_dep_.CoverTab[12468]++
//line /usr/local/go/src/net/netip/netip.go:1346
		// _ = "end of CoverTab[12468]"
//line /usr/local/go/src/net/netip/netip.go:1346
	}
//line /usr/local/go/src/net/netip/netip.go:1346
	// _ = "end of CoverTab[12463]"
//line /usr/local/go/src/net/netip/netip.go:1346
	_go_fuzz_dep_.CoverTab[12464]++
							if f1, f2 := p.ip.BitLen(), ip.BitLen(); f1 == 0 || func() bool {
//line /usr/local/go/src/net/netip/netip.go:1347
		_go_fuzz_dep_.CoverTab[12469]++
//line /usr/local/go/src/net/netip/netip.go:1347
		return f2 == 0
//line /usr/local/go/src/net/netip/netip.go:1347
		// _ = "end of CoverTab[12469]"
//line /usr/local/go/src/net/netip/netip.go:1347
	}() || func() bool {
//line /usr/local/go/src/net/netip/netip.go:1347
		_go_fuzz_dep_.CoverTab[12470]++
//line /usr/local/go/src/net/netip/netip.go:1347
		return f1 != f2
//line /usr/local/go/src/net/netip/netip.go:1347
		// _ = "end of CoverTab[12470]"
//line /usr/local/go/src/net/netip/netip.go:1347
	}() {
//line /usr/local/go/src/net/netip/netip.go:1347
		_go_fuzz_dep_.CoverTab[12471]++
								return false
//line /usr/local/go/src/net/netip/netip.go:1348
		// _ = "end of CoverTab[12471]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1349
		_go_fuzz_dep_.CoverTab[12472]++
//line /usr/local/go/src/net/netip/netip.go:1349
		// _ = "end of CoverTab[12472]"
//line /usr/local/go/src/net/netip/netip.go:1349
	}
//line /usr/local/go/src/net/netip/netip.go:1349
	// _ = "end of CoverTab[12464]"
//line /usr/local/go/src/net/netip/netip.go:1349
	_go_fuzz_dep_.CoverTab[12465]++
							if ip.Is4() {
//line /usr/local/go/src/net/netip/netip.go:1350
		_go_fuzz_dep_.CoverTab[12473]++

//line /usr/local/go/src/net/netip/netip.go:1359
		return uint32((ip.addr.lo^p.ip.addr.lo)>>((32-p.bits)&63)) == 0
//line /usr/local/go/src/net/netip/netip.go:1359
		// _ = "end of CoverTab[12473]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1360
		_go_fuzz_dep_.CoverTab[12474]++

//line /usr/local/go/src/net/netip/netip.go:1364
		return ip.addr.xor(p.ip.addr).and(mask6(int(p.bits))).isZero()
//line /usr/local/go/src/net/netip/netip.go:1364
		// _ = "end of CoverTab[12474]"
	}
//line /usr/local/go/src/net/netip/netip.go:1365
	// _ = "end of CoverTab[12465]"
}

// Overlaps reports whether p and o contain any IP addresses in common.
//line /usr/local/go/src/net/netip/netip.go:1368
//
//line /usr/local/go/src/net/netip/netip.go:1368
// If p and o are of different address families or either have a zero
//line /usr/local/go/src/net/netip/netip.go:1368
// IP, it reports false. Like the Contains method, a prefix with an
//line /usr/local/go/src/net/netip/netip.go:1368
// IPv4-mapped IPv6 address is still treated as an IPv6 mask.
//line /usr/local/go/src/net/netip/netip.go:1373
func (p Prefix) Overlaps(o Prefix) bool {
//line /usr/local/go/src/net/netip/netip.go:1373
	_go_fuzz_dep_.CoverTab[12475]++
							if !p.IsValid() || func() bool {
//line /usr/local/go/src/net/netip/netip.go:1374
		_go_fuzz_dep_.CoverTab[12483]++
//line /usr/local/go/src/net/netip/netip.go:1374
		return !o.IsValid()
//line /usr/local/go/src/net/netip/netip.go:1374
		// _ = "end of CoverTab[12483]"
//line /usr/local/go/src/net/netip/netip.go:1374
	}() {
//line /usr/local/go/src/net/netip/netip.go:1374
		_go_fuzz_dep_.CoverTab[12484]++
								return false
//line /usr/local/go/src/net/netip/netip.go:1375
		// _ = "end of CoverTab[12484]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1376
		_go_fuzz_dep_.CoverTab[12485]++
//line /usr/local/go/src/net/netip/netip.go:1376
		// _ = "end of CoverTab[12485]"
//line /usr/local/go/src/net/netip/netip.go:1376
	}
//line /usr/local/go/src/net/netip/netip.go:1376
	// _ = "end of CoverTab[12475]"
//line /usr/local/go/src/net/netip/netip.go:1376
	_go_fuzz_dep_.CoverTab[12476]++
							if p == o {
//line /usr/local/go/src/net/netip/netip.go:1377
		_go_fuzz_dep_.CoverTab[12486]++
								return true
//line /usr/local/go/src/net/netip/netip.go:1378
		// _ = "end of CoverTab[12486]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1379
		_go_fuzz_dep_.CoverTab[12487]++
//line /usr/local/go/src/net/netip/netip.go:1379
		// _ = "end of CoverTab[12487]"
//line /usr/local/go/src/net/netip/netip.go:1379
	}
//line /usr/local/go/src/net/netip/netip.go:1379
	// _ = "end of CoverTab[12476]"
//line /usr/local/go/src/net/netip/netip.go:1379
	_go_fuzz_dep_.CoverTab[12477]++
							if p.ip.Is4() != o.ip.Is4() {
//line /usr/local/go/src/net/netip/netip.go:1380
		_go_fuzz_dep_.CoverTab[12488]++
								return false
//line /usr/local/go/src/net/netip/netip.go:1381
		// _ = "end of CoverTab[12488]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1382
		_go_fuzz_dep_.CoverTab[12489]++
//line /usr/local/go/src/net/netip/netip.go:1382
		// _ = "end of CoverTab[12489]"
//line /usr/local/go/src/net/netip/netip.go:1382
	}
//line /usr/local/go/src/net/netip/netip.go:1382
	// _ = "end of CoverTab[12477]"
//line /usr/local/go/src/net/netip/netip.go:1382
	_go_fuzz_dep_.CoverTab[12478]++
							var minBits int16
							if p.bits < o.bits {
//line /usr/local/go/src/net/netip/netip.go:1384
		_go_fuzz_dep_.CoverTab[12490]++
								minBits = p.bits
//line /usr/local/go/src/net/netip/netip.go:1385
		// _ = "end of CoverTab[12490]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1386
		_go_fuzz_dep_.CoverTab[12491]++
								minBits = o.bits
//line /usr/local/go/src/net/netip/netip.go:1387
		// _ = "end of CoverTab[12491]"
	}
//line /usr/local/go/src/net/netip/netip.go:1388
	// _ = "end of CoverTab[12478]"
//line /usr/local/go/src/net/netip/netip.go:1388
	_go_fuzz_dep_.CoverTab[12479]++
							if minBits == 0 {
//line /usr/local/go/src/net/netip/netip.go:1389
		_go_fuzz_dep_.CoverTab[12492]++
								return true
//line /usr/local/go/src/net/netip/netip.go:1390
		// _ = "end of CoverTab[12492]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1391
		_go_fuzz_dep_.CoverTab[12493]++
//line /usr/local/go/src/net/netip/netip.go:1391
		// _ = "end of CoverTab[12493]"
//line /usr/local/go/src/net/netip/netip.go:1391
	}
//line /usr/local/go/src/net/netip/netip.go:1391
	// _ = "end of CoverTab[12479]"
//line /usr/local/go/src/net/netip/netip.go:1391
	_go_fuzz_dep_.CoverTab[12480]++
	// One of these Prefix calls might look redundant, but we don't require
	// that p and o values are normalized (via Prefix.Masked) first,
	// so the Prefix call on the one that's already minBits serves to zero
	// out any remaining bits in IP.
	var err error
	if p, err = p.ip.Prefix(int(minBits)); err != nil {
//line /usr/local/go/src/net/netip/netip.go:1397
		_go_fuzz_dep_.CoverTab[12494]++
								return false
//line /usr/local/go/src/net/netip/netip.go:1398
		// _ = "end of CoverTab[12494]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1399
		_go_fuzz_dep_.CoverTab[12495]++
//line /usr/local/go/src/net/netip/netip.go:1399
		// _ = "end of CoverTab[12495]"
//line /usr/local/go/src/net/netip/netip.go:1399
	}
//line /usr/local/go/src/net/netip/netip.go:1399
	// _ = "end of CoverTab[12480]"
//line /usr/local/go/src/net/netip/netip.go:1399
	_go_fuzz_dep_.CoverTab[12481]++
							if o, err = o.ip.Prefix(int(minBits)); err != nil {
//line /usr/local/go/src/net/netip/netip.go:1400
		_go_fuzz_dep_.CoverTab[12496]++
								return false
//line /usr/local/go/src/net/netip/netip.go:1401
		// _ = "end of CoverTab[12496]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1402
		_go_fuzz_dep_.CoverTab[12497]++
//line /usr/local/go/src/net/netip/netip.go:1402
		// _ = "end of CoverTab[12497]"
//line /usr/local/go/src/net/netip/netip.go:1402
	}
//line /usr/local/go/src/net/netip/netip.go:1402
	// _ = "end of CoverTab[12481]"
//line /usr/local/go/src/net/netip/netip.go:1402
	_go_fuzz_dep_.CoverTab[12482]++
							return p.ip == o.ip
//line /usr/local/go/src/net/netip/netip.go:1403
	// _ = "end of CoverTab[12482]"
}

// AppendTo appends a text encoding of p,
//line /usr/local/go/src/net/netip/netip.go:1406
// as generated by MarshalText,
//line /usr/local/go/src/net/netip/netip.go:1406
// to b and returns the extended buffer.
//line /usr/local/go/src/net/netip/netip.go:1409
func (p Prefix) AppendTo(b []byte) []byte {
//line /usr/local/go/src/net/netip/netip.go:1409
	_go_fuzz_dep_.CoverTab[12498]++
							if p.isZero() {
//line /usr/local/go/src/net/netip/netip.go:1410
		_go_fuzz_dep_.CoverTab[12502]++
								return b
//line /usr/local/go/src/net/netip/netip.go:1411
		// _ = "end of CoverTab[12502]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1412
		_go_fuzz_dep_.CoverTab[12503]++
//line /usr/local/go/src/net/netip/netip.go:1412
		// _ = "end of CoverTab[12503]"
//line /usr/local/go/src/net/netip/netip.go:1412
	}
//line /usr/local/go/src/net/netip/netip.go:1412
	// _ = "end of CoverTab[12498]"
//line /usr/local/go/src/net/netip/netip.go:1412
	_go_fuzz_dep_.CoverTab[12499]++
							if !p.IsValid() {
//line /usr/local/go/src/net/netip/netip.go:1413
		_go_fuzz_dep_.CoverTab[12504]++
								return append(b, "invalid Prefix"...)
//line /usr/local/go/src/net/netip/netip.go:1414
		// _ = "end of CoverTab[12504]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1415
		_go_fuzz_dep_.CoverTab[12505]++
//line /usr/local/go/src/net/netip/netip.go:1415
		// _ = "end of CoverTab[12505]"
//line /usr/local/go/src/net/netip/netip.go:1415
	}
//line /usr/local/go/src/net/netip/netip.go:1415
	// _ = "end of CoverTab[12499]"
//line /usr/local/go/src/net/netip/netip.go:1415
	_go_fuzz_dep_.CoverTab[12500]++

//line /usr/local/go/src/net/netip/netip.go:1418
	if p.ip.z == z4 {
//line /usr/local/go/src/net/netip/netip.go:1418
		_go_fuzz_dep_.CoverTab[12506]++
								b = p.ip.appendTo4(b)
//line /usr/local/go/src/net/netip/netip.go:1419
		// _ = "end of CoverTab[12506]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1420
		_go_fuzz_dep_.CoverTab[12507]++
								if p.ip.Is4In6() {
//line /usr/local/go/src/net/netip/netip.go:1421
			_go_fuzz_dep_.CoverTab[12508]++
									b = append(b, "::ffff:"...)
									b = p.ip.Unmap().appendTo4(b)
//line /usr/local/go/src/net/netip/netip.go:1423
			// _ = "end of CoverTab[12508]"
		} else {
//line /usr/local/go/src/net/netip/netip.go:1424
			_go_fuzz_dep_.CoverTab[12509]++
									b = p.ip.appendTo6(b)
//line /usr/local/go/src/net/netip/netip.go:1425
			// _ = "end of CoverTab[12509]"
		}
//line /usr/local/go/src/net/netip/netip.go:1426
		// _ = "end of CoverTab[12507]"
	}
//line /usr/local/go/src/net/netip/netip.go:1427
	// _ = "end of CoverTab[12500]"
//line /usr/local/go/src/net/netip/netip.go:1427
	_go_fuzz_dep_.CoverTab[12501]++

							b = append(b, '/')
							b = appendDecimal(b, uint8(p.bits))
							return b
//line /usr/local/go/src/net/netip/netip.go:1431
	// _ = "end of CoverTab[12501]"
}

// MarshalText implements the encoding.TextMarshaler interface,
//line /usr/local/go/src/net/netip/netip.go:1434
// The encoding is the same as returned by String, with one exception:
//line /usr/local/go/src/net/netip/netip.go:1434
// If p is the zero value, the encoding is the empty string.
//line /usr/local/go/src/net/netip/netip.go:1437
func (p Prefix) MarshalText() ([]byte, error) {
//line /usr/local/go/src/net/netip/netip.go:1437
	_go_fuzz_dep_.CoverTab[12510]++
							var max int
							switch p.ip.z {
	case z0:
//line /usr/local/go/src/net/netip/netip.go:1440
		_go_fuzz_dep_.CoverTab[12512]++
//line /usr/local/go/src/net/netip/netip.go:1440
		// _ = "end of CoverTab[12512]"
	case z4:
//line /usr/local/go/src/net/netip/netip.go:1441
		_go_fuzz_dep_.CoverTab[12513]++
								max = len("255.255.255.255/32")
//line /usr/local/go/src/net/netip/netip.go:1442
		// _ = "end of CoverTab[12513]"
	default:
//line /usr/local/go/src/net/netip/netip.go:1443
		_go_fuzz_dep_.CoverTab[12514]++
								max = len("ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff%enp5s0/128")
//line /usr/local/go/src/net/netip/netip.go:1444
		// _ = "end of CoverTab[12514]"
	}
//line /usr/local/go/src/net/netip/netip.go:1445
	// _ = "end of CoverTab[12510]"
//line /usr/local/go/src/net/netip/netip.go:1445
	_go_fuzz_dep_.CoverTab[12511]++
							b := make([]byte, 0, max)
							b = p.AppendTo(b)
							return b, nil
//line /usr/local/go/src/net/netip/netip.go:1448
	// _ = "end of CoverTab[12511]"
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
//line /usr/local/go/src/net/netip/netip.go:1451
// The IP address is expected in a form accepted by ParsePrefix
//line /usr/local/go/src/net/netip/netip.go:1451
// or generated by MarshalText.
//line /usr/local/go/src/net/netip/netip.go:1454
func (p *Prefix) UnmarshalText(text []byte) error {
//line /usr/local/go/src/net/netip/netip.go:1454
	_go_fuzz_dep_.CoverTab[12515]++
							if len(text) == 0 {
//line /usr/local/go/src/net/netip/netip.go:1455
		_go_fuzz_dep_.CoverTab[12517]++
								*p = Prefix{}
								return nil
//line /usr/local/go/src/net/netip/netip.go:1457
		// _ = "end of CoverTab[12517]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1458
		_go_fuzz_dep_.CoverTab[12518]++
//line /usr/local/go/src/net/netip/netip.go:1458
		// _ = "end of CoverTab[12518]"
//line /usr/local/go/src/net/netip/netip.go:1458
	}
//line /usr/local/go/src/net/netip/netip.go:1458
	// _ = "end of CoverTab[12515]"
//line /usr/local/go/src/net/netip/netip.go:1458
	_go_fuzz_dep_.CoverTab[12516]++
							var err error
							*p, err = ParsePrefix(string(text))
							return err
//line /usr/local/go/src/net/netip/netip.go:1461
	// _ = "end of CoverTab[12516]"
}

// MarshalBinary implements the encoding.BinaryMarshaler interface.
//line /usr/local/go/src/net/netip/netip.go:1464
// It returns Addr.MarshalBinary with an additional byte appended
//line /usr/local/go/src/net/netip/netip.go:1464
// containing the prefix bits.
//line /usr/local/go/src/net/netip/netip.go:1467
func (p Prefix) MarshalBinary() ([]byte, error) {
//line /usr/local/go/src/net/netip/netip.go:1467
	_go_fuzz_dep_.CoverTab[12519]++
							b := p.Addr().withoutZone().marshalBinaryWithTrailingBytes(1)
							b[len(b)-1] = uint8(p.Bits())
							return b, nil
//line /usr/local/go/src/net/netip/netip.go:1470
	// _ = "end of CoverTab[12519]"
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface.
//line /usr/local/go/src/net/netip/netip.go:1473
// It expects data in the form generated by MarshalBinary.
//line /usr/local/go/src/net/netip/netip.go:1475
func (p *Prefix) UnmarshalBinary(b []byte) error {
//line /usr/local/go/src/net/netip/netip.go:1475
	_go_fuzz_dep_.CoverTab[12520]++
							if len(b) < 1 {
//line /usr/local/go/src/net/netip/netip.go:1476
		_go_fuzz_dep_.CoverTab[12523]++
								return errors.New("unexpected slice size")
//line /usr/local/go/src/net/netip/netip.go:1477
		// _ = "end of CoverTab[12523]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1478
		_go_fuzz_dep_.CoverTab[12524]++
//line /usr/local/go/src/net/netip/netip.go:1478
		// _ = "end of CoverTab[12524]"
//line /usr/local/go/src/net/netip/netip.go:1478
	}
//line /usr/local/go/src/net/netip/netip.go:1478
	// _ = "end of CoverTab[12520]"
//line /usr/local/go/src/net/netip/netip.go:1478
	_go_fuzz_dep_.CoverTab[12521]++
							var addr Addr
							err := addr.UnmarshalBinary(b[:len(b)-1])
							if err != nil {
//line /usr/local/go/src/net/netip/netip.go:1481
		_go_fuzz_dep_.CoverTab[12525]++
								return err
//line /usr/local/go/src/net/netip/netip.go:1482
		// _ = "end of CoverTab[12525]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1483
		_go_fuzz_dep_.CoverTab[12526]++
//line /usr/local/go/src/net/netip/netip.go:1483
		// _ = "end of CoverTab[12526]"
//line /usr/local/go/src/net/netip/netip.go:1483
	}
//line /usr/local/go/src/net/netip/netip.go:1483
	// _ = "end of CoverTab[12521]"
//line /usr/local/go/src/net/netip/netip.go:1483
	_go_fuzz_dep_.CoverTab[12522]++
							*p = PrefixFrom(addr, int(b[len(b)-1]))
							return nil
//line /usr/local/go/src/net/netip/netip.go:1485
	// _ = "end of CoverTab[12522]"
}

// String returns the CIDR notation of p: "<ip>/<bits>".
func (p Prefix) String() string {
//line /usr/local/go/src/net/netip/netip.go:1489
	_go_fuzz_dep_.CoverTab[12527]++
							if !p.IsValid() {
//line /usr/local/go/src/net/netip/netip.go:1490
		_go_fuzz_dep_.CoverTab[12529]++
								return "invalid Prefix"
//line /usr/local/go/src/net/netip/netip.go:1491
		// _ = "end of CoverTab[12529]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1492
		_go_fuzz_dep_.CoverTab[12530]++
//line /usr/local/go/src/net/netip/netip.go:1492
		// _ = "end of CoverTab[12530]"
//line /usr/local/go/src/net/netip/netip.go:1492
	}
//line /usr/local/go/src/net/netip/netip.go:1492
	// _ = "end of CoverTab[12527]"
//line /usr/local/go/src/net/netip/netip.go:1492
	_go_fuzz_dep_.CoverTab[12528]++
							return p.ip.String() + "/" + itoa.Itoa(int(p.bits))
//line /usr/local/go/src/net/netip/netip.go:1493
	// _ = "end of CoverTab[12528]"
}

//line /usr/local/go/src/net/netip/netip.go:1494
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/netip/netip.go:1494
var _ = _go_fuzz_dep_.CoverTab
