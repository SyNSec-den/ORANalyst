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
	_go_fuzz_dep_.CoverTab[3556]++
//line /usr/local/go/src/net/netip/netip.go:76
	return AddrFrom16([16]byte{0: 0xff, 1: 0x02, 15: 0x01})
//line /usr/local/go/src/net/netip/netip.go:76
	// _ = "end of CoverTab[3556]"
//line /usr/local/go/src/net/netip/netip.go:76
}

// IPv6LinkLocalAllRouters returns the IPv6 link-local all routers multicast
//line /usr/local/go/src/net/netip/netip.go:78
// address ff02::2.
//line /usr/local/go/src/net/netip/netip.go:80
func IPv6LinkLocalAllRouters() Addr {
//line /usr/local/go/src/net/netip/netip.go:80
	_go_fuzz_dep_.CoverTab[3557]++
//line /usr/local/go/src/net/netip/netip.go:80
	return AddrFrom16([16]byte{0: 0xff, 1: 0x02, 15: 0x02})
//line /usr/local/go/src/net/netip/netip.go:80
	// _ = "end of CoverTab[3557]"
//line /usr/local/go/src/net/netip/netip.go:80
}

// IPv6Loopback returns the IPv6 loopback address ::1.
func IPv6Loopback() Addr {
//line /usr/local/go/src/net/netip/netip.go:83
	_go_fuzz_dep_.CoverTab[3558]++
//line /usr/local/go/src/net/netip/netip.go:83
	return AddrFrom16([16]byte{15: 0x01})
//line /usr/local/go/src/net/netip/netip.go:83
	// _ = "end of CoverTab[3558]"
//line /usr/local/go/src/net/netip/netip.go:83
}

// IPv6Unspecified returns the IPv6 unspecified address "::".
func IPv6Unspecified() Addr {
//line /usr/local/go/src/net/netip/netip.go:86
	_go_fuzz_dep_.CoverTab[3559]++
//line /usr/local/go/src/net/netip/netip.go:86
	return Addr{z: z6noz}
//line /usr/local/go/src/net/netip/netip.go:86
	// _ = "end of CoverTab[3559]"
//line /usr/local/go/src/net/netip/netip.go:86
}

// IPv4Unspecified returns the IPv4 unspecified address "0.0.0.0".
func IPv4Unspecified() Addr {
//line /usr/local/go/src/net/netip/netip.go:89
	_go_fuzz_dep_.CoverTab[3560]++
//line /usr/local/go/src/net/netip/netip.go:89
	return AddrFrom4([4]byte{})
//line /usr/local/go/src/net/netip/netip.go:89
	// _ = "end of CoverTab[3560]"
//line /usr/local/go/src/net/netip/netip.go:89
}

// AddrFrom4 returns the address of the IPv4 address given by the bytes in addr.
func AddrFrom4(addr [4]byte) Addr {
//line /usr/local/go/src/net/netip/netip.go:92
	_go_fuzz_dep_.CoverTab[3561]++
						return Addr{
		addr:	uint128{0, 0xffff00000000 | uint64(addr[0])<<24 | uint64(addr[1])<<16 | uint64(addr[2])<<8 | uint64(addr[3])},
		z:	z4,
	}
//line /usr/local/go/src/net/netip/netip.go:96
	// _ = "end of CoverTab[3561]"
}

// AddrFrom16 returns the IPv6 address given by the bytes in addr.
//line /usr/local/go/src/net/netip/netip.go:99
// An IPv4-mapped IPv6 address is left as an IPv6 address.
//line /usr/local/go/src/net/netip/netip.go:99
// (Use Unmap to convert them if needed.)
//line /usr/local/go/src/net/netip/netip.go:102
func AddrFrom16(addr [16]byte) Addr {
//line /usr/local/go/src/net/netip/netip.go:102
	_go_fuzz_dep_.CoverTab[3562]++
							return Addr{
		addr: uint128{
			beUint64(addr[:8]),
			beUint64(addr[8:]),
		},
		z:	z6noz,
	}
//line /usr/local/go/src/net/netip/netip.go:109
	// _ = "end of CoverTab[3562]"
}

// ParseAddr parses s as an IP address, returning the result. The string
//line /usr/local/go/src/net/netip/netip.go:112
// s can be in dotted decimal ("192.0.2.1"), IPv6 ("2001:db8::68"),
//line /usr/local/go/src/net/netip/netip.go:112
// or IPv6 with a scoped addressing zone ("fe80::1cc0:3e8c:119f:c2e1%ens18").
//line /usr/local/go/src/net/netip/netip.go:115
func ParseAddr(s string) (Addr, error) {
//line /usr/local/go/src/net/netip/netip.go:115
	_go_fuzz_dep_.CoverTab[3563]++
							for i := 0; i < len(s); i++ {
//line /usr/local/go/src/net/netip/netip.go:116
		_go_fuzz_dep_.CoverTab[3565]++
								switch s[i] {
		case '.':
//line /usr/local/go/src/net/netip/netip.go:118
			_go_fuzz_dep_.CoverTab[3566]++
									return parseIPv4(s)
//line /usr/local/go/src/net/netip/netip.go:119
			// _ = "end of CoverTab[3566]"
		case ':':
//line /usr/local/go/src/net/netip/netip.go:120
			_go_fuzz_dep_.CoverTab[3567]++
									return parseIPv6(s)
//line /usr/local/go/src/net/netip/netip.go:121
			// _ = "end of CoverTab[3567]"
		case '%':
//line /usr/local/go/src/net/netip/netip.go:122
			_go_fuzz_dep_.CoverTab[3568]++

//line /usr/local/go/src/net/netip/netip.go:125
			return Addr{}, parseAddrError{in: s, msg: "missing IPv6 address"}
//line /usr/local/go/src/net/netip/netip.go:125
			// _ = "end of CoverTab[3568]"
//line /usr/local/go/src/net/netip/netip.go:125
		default:
//line /usr/local/go/src/net/netip/netip.go:125
			_go_fuzz_dep_.CoverTab[3569]++
//line /usr/local/go/src/net/netip/netip.go:125
			// _ = "end of CoverTab[3569]"
		}
//line /usr/local/go/src/net/netip/netip.go:126
		// _ = "end of CoverTab[3565]"
	}
//line /usr/local/go/src/net/netip/netip.go:127
	// _ = "end of CoverTab[3563]"
//line /usr/local/go/src/net/netip/netip.go:127
	_go_fuzz_dep_.CoverTab[3564]++
							return Addr{}, parseAddrError{in: s, msg: "unable to parse IP"}
//line /usr/local/go/src/net/netip/netip.go:128
	// _ = "end of CoverTab[3564]"
}

// MustParseAddr calls ParseAddr(s) and panics on error.
//line /usr/local/go/src/net/netip/netip.go:131
// It is intended for use in tests with hard-coded strings.
//line /usr/local/go/src/net/netip/netip.go:133
func MustParseAddr(s string) Addr {
//line /usr/local/go/src/net/netip/netip.go:133
	_go_fuzz_dep_.CoverTab[3570]++
							ip, err := ParseAddr(s)
							if err != nil {
//line /usr/local/go/src/net/netip/netip.go:135
		_go_fuzz_dep_.CoverTab[3572]++
								panic(err)
//line /usr/local/go/src/net/netip/netip.go:136
		// _ = "end of CoverTab[3572]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:137
		_go_fuzz_dep_.CoverTab[3573]++
//line /usr/local/go/src/net/netip/netip.go:137
		// _ = "end of CoverTab[3573]"
//line /usr/local/go/src/net/netip/netip.go:137
	}
//line /usr/local/go/src/net/netip/netip.go:137
	// _ = "end of CoverTab[3570]"
//line /usr/local/go/src/net/netip/netip.go:137
	_go_fuzz_dep_.CoverTab[3571]++
							return ip
//line /usr/local/go/src/net/netip/netip.go:138
	// _ = "end of CoverTab[3571]"
}

type parseAddrError struct {
	in	string	// the string given to ParseAddr
	msg	string	// an explanation of the parse failure
	at	string	// optionally, the unparsed portion of in at which the error occurred.
}

func (err parseAddrError) Error() string {
//line /usr/local/go/src/net/netip/netip.go:147
	_go_fuzz_dep_.CoverTab[3574]++
							q := strconv.Quote
							if err.at != "" {
//line /usr/local/go/src/net/netip/netip.go:149
		_go_fuzz_dep_.CoverTab[3576]++
								return "ParseAddr(" + q(err.in) + "): " + err.msg + " (at " + q(err.at) + ")"
//line /usr/local/go/src/net/netip/netip.go:150
		// _ = "end of CoverTab[3576]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:151
		_go_fuzz_dep_.CoverTab[3577]++
//line /usr/local/go/src/net/netip/netip.go:151
		// _ = "end of CoverTab[3577]"
//line /usr/local/go/src/net/netip/netip.go:151
	}
//line /usr/local/go/src/net/netip/netip.go:151
	// _ = "end of CoverTab[3574]"
//line /usr/local/go/src/net/netip/netip.go:151
	_go_fuzz_dep_.CoverTab[3575]++
							return "ParseAddr(" + q(err.in) + "): " + err.msg
//line /usr/local/go/src/net/netip/netip.go:152
	// _ = "end of CoverTab[3575]"
}

// parseIPv4 parses s as an IPv4 address (in form "192.168.0.1").
func parseIPv4(s string) (ip Addr, err error) {
//line /usr/local/go/src/net/netip/netip.go:156
	_go_fuzz_dep_.CoverTab[3578]++
							var fields [4]uint8
							var val, pos int
							var digLen int	// number of digits in current octet
							for i := 0; i < len(s); i++ {
//line /usr/local/go/src/net/netip/netip.go:160
		_go_fuzz_dep_.CoverTab[3581]++
								if s[i] >= '0' && func() bool {
//line /usr/local/go/src/net/netip/netip.go:161
			_go_fuzz_dep_.CoverTab[3582]++
//line /usr/local/go/src/net/netip/netip.go:161
			return s[i] <= '9'
//line /usr/local/go/src/net/netip/netip.go:161
			// _ = "end of CoverTab[3582]"
//line /usr/local/go/src/net/netip/netip.go:161
		}() {
//line /usr/local/go/src/net/netip/netip.go:161
			_go_fuzz_dep_.CoverTab[3583]++
									if digLen == 1 && func() bool {
//line /usr/local/go/src/net/netip/netip.go:162
				_go_fuzz_dep_.CoverTab[3585]++
//line /usr/local/go/src/net/netip/netip.go:162
				return val == 0
//line /usr/local/go/src/net/netip/netip.go:162
				// _ = "end of CoverTab[3585]"
//line /usr/local/go/src/net/netip/netip.go:162
			}() {
//line /usr/local/go/src/net/netip/netip.go:162
				_go_fuzz_dep_.CoverTab[3586]++
										return Addr{}, parseAddrError{in: s, msg: "IPv4 field has octet with leading zero"}
//line /usr/local/go/src/net/netip/netip.go:163
				// _ = "end of CoverTab[3586]"
			} else {
//line /usr/local/go/src/net/netip/netip.go:164
				_go_fuzz_dep_.CoverTab[3587]++
//line /usr/local/go/src/net/netip/netip.go:164
				// _ = "end of CoverTab[3587]"
//line /usr/local/go/src/net/netip/netip.go:164
			}
//line /usr/local/go/src/net/netip/netip.go:164
			// _ = "end of CoverTab[3583]"
//line /usr/local/go/src/net/netip/netip.go:164
			_go_fuzz_dep_.CoverTab[3584]++
									val = val*10 + int(s[i]) - '0'
									digLen++
									if val > 255 {
//line /usr/local/go/src/net/netip/netip.go:167
				_go_fuzz_dep_.CoverTab[3588]++
										return Addr{}, parseAddrError{in: s, msg: "IPv4 field has value >255"}
//line /usr/local/go/src/net/netip/netip.go:168
				// _ = "end of CoverTab[3588]"
			} else {
//line /usr/local/go/src/net/netip/netip.go:169
				_go_fuzz_dep_.CoverTab[3589]++
//line /usr/local/go/src/net/netip/netip.go:169
				// _ = "end of CoverTab[3589]"
//line /usr/local/go/src/net/netip/netip.go:169
			}
//line /usr/local/go/src/net/netip/netip.go:169
			// _ = "end of CoverTab[3584]"
		} else {
//line /usr/local/go/src/net/netip/netip.go:170
			_go_fuzz_dep_.CoverTab[3590]++
//line /usr/local/go/src/net/netip/netip.go:170
			if s[i] == '.' {
//line /usr/local/go/src/net/netip/netip.go:170
				_go_fuzz_dep_.CoverTab[3591]++

//line /usr/local/go/src/net/netip/netip.go:174
				if i == 0 || func() bool {
//line /usr/local/go/src/net/netip/netip.go:174
					_go_fuzz_dep_.CoverTab[3594]++
//line /usr/local/go/src/net/netip/netip.go:174
					return i == len(s)-1
//line /usr/local/go/src/net/netip/netip.go:174
					// _ = "end of CoverTab[3594]"
//line /usr/local/go/src/net/netip/netip.go:174
				}() || func() bool {
//line /usr/local/go/src/net/netip/netip.go:174
					_go_fuzz_dep_.CoverTab[3595]++
//line /usr/local/go/src/net/netip/netip.go:174
					return s[i-1] == '.'
//line /usr/local/go/src/net/netip/netip.go:174
					// _ = "end of CoverTab[3595]"
//line /usr/local/go/src/net/netip/netip.go:174
				}() {
//line /usr/local/go/src/net/netip/netip.go:174
					_go_fuzz_dep_.CoverTab[3596]++
											return Addr{}, parseAddrError{in: s, msg: "IPv4 field must have at least one digit", at: s[i:]}
//line /usr/local/go/src/net/netip/netip.go:175
					// _ = "end of CoverTab[3596]"
				} else {
//line /usr/local/go/src/net/netip/netip.go:176
					_go_fuzz_dep_.CoverTab[3597]++
//line /usr/local/go/src/net/netip/netip.go:176
					// _ = "end of CoverTab[3597]"
//line /usr/local/go/src/net/netip/netip.go:176
				}
//line /usr/local/go/src/net/netip/netip.go:176
				// _ = "end of CoverTab[3591]"
//line /usr/local/go/src/net/netip/netip.go:176
				_go_fuzz_dep_.CoverTab[3592]++

										if pos == 3 {
//line /usr/local/go/src/net/netip/netip.go:178
					_go_fuzz_dep_.CoverTab[3598]++
											return Addr{}, parseAddrError{in: s, msg: "IPv4 address too long"}
//line /usr/local/go/src/net/netip/netip.go:179
					// _ = "end of CoverTab[3598]"
				} else {
//line /usr/local/go/src/net/netip/netip.go:180
					_go_fuzz_dep_.CoverTab[3599]++
//line /usr/local/go/src/net/netip/netip.go:180
					// _ = "end of CoverTab[3599]"
//line /usr/local/go/src/net/netip/netip.go:180
				}
//line /usr/local/go/src/net/netip/netip.go:180
				// _ = "end of CoverTab[3592]"
//line /usr/local/go/src/net/netip/netip.go:180
				_go_fuzz_dep_.CoverTab[3593]++
										fields[pos] = uint8(val)
										pos++
										val = 0
										digLen = 0
//line /usr/local/go/src/net/netip/netip.go:184
				// _ = "end of CoverTab[3593]"
			} else {
//line /usr/local/go/src/net/netip/netip.go:185
				_go_fuzz_dep_.CoverTab[3600]++
										return Addr{}, parseAddrError{in: s, msg: "unexpected character", at: s[i:]}
//line /usr/local/go/src/net/netip/netip.go:186
				// _ = "end of CoverTab[3600]"
			}
//line /usr/local/go/src/net/netip/netip.go:187
			// _ = "end of CoverTab[3590]"
//line /usr/local/go/src/net/netip/netip.go:187
		}
//line /usr/local/go/src/net/netip/netip.go:187
		// _ = "end of CoverTab[3581]"
	}
//line /usr/local/go/src/net/netip/netip.go:188
	// _ = "end of CoverTab[3578]"
//line /usr/local/go/src/net/netip/netip.go:188
	_go_fuzz_dep_.CoverTab[3579]++
							if pos < 3 {
//line /usr/local/go/src/net/netip/netip.go:189
		_go_fuzz_dep_.CoverTab[3601]++
								return Addr{}, parseAddrError{in: s, msg: "IPv4 address too short"}
//line /usr/local/go/src/net/netip/netip.go:190
		// _ = "end of CoverTab[3601]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:191
		_go_fuzz_dep_.CoverTab[3602]++
//line /usr/local/go/src/net/netip/netip.go:191
		// _ = "end of CoverTab[3602]"
//line /usr/local/go/src/net/netip/netip.go:191
	}
//line /usr/local/go/src/net/netip/netip.go:191
	// _ = "end of CoverTab[3579]"
//line /usr/local/go/src/net/netip/netip.go:191
	_go_fuzz_dep_.CoverTab[3580]++
							fields[3] = uint8(val)
							return AddrFrom4(fields), nil
//line /usr/local/go/src/net/netip/netip.go:193
	// _ = "end of CoverTab[3580]"
}

// parseIPv6 parses s as an IPv6 address (in form "2001:db8::68").
func parseIPv6(in string) (Addr, error) {
//line /usr/local/go/src/net/netip/netip.go:197
	_go_fuzz_dep_.CoverTab[3603]++
							s := in

//line /usr/local/go/src/net/netip/netip.go:204
	zone := ""
	i := bytealg.IndexByteString(s, '%')
	if i != -1 {
//line /usr/local/go/src/net/netip/netip.go:206
		_go_fuzz_dep_.CoverTab[3609]++
								s, zone = s[:i], s[i+1:]
								if zone == "" {
//line /usr/local/go/src/net/netip/netip.go:208
			_go_fuzz_dep_.CoverTab[3610]++

									return Addr{}, parseAddrError{in: in, msg: "zone must be a non-empty string"}
//line /usr/local/go/src/net/netip/netip.go:210
			// _ = "end of CoverTab[3610]"
		} else {
//line /usr/local/go/src/net/netip/netip.go:211
			_go_fuzz_dep_.CoverTab[3611]++
//line /usr/local/go/src/net/netip/netip.go:211
			// _ = "end of CoverTab[3611]"
//line /usr/local/go/src/net/netip/netip.go:211
		}
//line /usr/local/go/src/net/netip/netip.go:211
		// _ = "end of CoverTab[3609]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:212
		_go_fuzz_dep_.CoverTab[3612]++
//line /usr/local/go/src/net/netip/netip.go:212
		// _ = "end of CoverTab[3612]"
//line /usr/local/go/src/net/netip/netip.go:212
	}
//line /usr/local/go/src/net/netip/netip.go:212
	// _ = "end of CoverTab[3603]"
//line /usr/local/go/src/net/netip/netip.go:212
	_go_fuzz_dep_.CoverTab[3604]++

							var ip [16]byte
							ellipsis := -1

//line /usr/local/go/src/net/netip/netip.go:218
	if len(s) >= 2 && func() bool {
//line /usr/local/go/src/net/netip/netip.go:218
		_go_fuzz_dep_.CoverTab[3613]++
//line /usr/local/go/src/net/netip/netip.go:218
		return s[0] == ':'
//line /usr/local/go/src/net/netip/netip.go:218
		// _ = "end of CoverTab[3613]"
//line /usr/local/go/src/net/netip/netip.go:218
	}() && func() bool {
//line /usr/local/go/src/net/netip/netip.go:218
		_go_fuzz_dep_.CoverTab[3614]++
//line /usr/local/go/src/net/netip/netip.go:218
		return s[1] == ':'
//line /usr/local/go/src/net/netip/netip.go:218
		// _ = "end of CoverTab[3614]"
//line /usr/local/go/src/net/netip/netip.go:218
	}() {
//line /usr/local/go/src/net/netip/netip.go:218
		_go_fuzz_dep_.CoverTab[3615]++
								ellipsis = 0
								s = s[2:]

								if len(s) == 0 {
//line /usr/local/go/src/net/netip/netip.go:222
			_go_fuzz_dep_.CoverTab[3616]++
									return IPv6Unspecified().WithZone(zone), nil
//line /usr/local/go/src/net/netip/netip.go:223
			// _ = "end of CoverTab[3616]"
		} else {
//line /usr/local/go/src/net/netip/netip.go:224
			_go_fuzz_dep_.CoverTab[3617]++
//line /usr/local/go/src/net/netip/netip.go:224
			// _ = "end of CoverTab[3617]"
//line /usr/local/go/src/net/netip/netip.go:224
		}
//line /usr/local/go/src/net/netip/netip.go:224
		// _ = "end of CoverTab[3615]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:225
		_go_fuzz_dep_.CoverTab[3618]++
//line /usr/local/go/src/net/netip/netip.go:225
		// _ = "end of CoverTab[3618]"
//line /usr/local/go/src/net/netip/netip.go:225
	}
//line /usr/local/go/src/net/netip/netip.go:225
	// _ = "end of CoverTab[3604]"
//line /usr/local/go/src/net/netip/netip.go:225
	_go_fuzz_dep_.CoverTab[3605]++

//line /usr/local/go/src/net/netip/netip.go:228
	i = 0
	for i < 16 {
//line /usr/local/go/src/net/netip/netip.go:229
		_go_fuzz_dep_.CoverTab[3619]++

//line /usr/local/go/src/net/netip/netip.go:232
		off := 0
		acc := uint32(0)
		for ; off < len(s); off++ {
//line /usr/local/go/src/net/netip/netip.go:234
			_go_fuzz_dep_.CoverTab[3625]++
									c := s[off]
									if c >= '0' && func() bool {
//line /usr/local/go/src/net/netip/netip.go:236
				_go_fuzz_dep_.CoverTab[3627]++
//line /usr/local/go/src/net/netip/netip.go:236
				return c <= '9'
//line /usr/local/go/src/net/netip/netip.go:236
				// _ = "end of CoverTab[3627]"
//line /usr/local/go/src/net/netip/netip.go:236
			}() {
//line /usr/local/go/src/net/netip/netip.go:236
				_go_fuzz_dep_.CoverTab[3628]++
										acc = (acc << 4) + uint32(c-'0')
//line /usr/local/go/src/net/netip/netip.go:237
				// _ = "end of CoverTab[3628]"
			} else {
//line /usr/local/go/src/net/netip/netip.go:238
				_go_fuzz_dep_.CoverTab[3629]++
//line /usr/local/go/src/net/netip/netip.go:238
				if c >= 'a' && func() bool {
//line /usr/local/go/src/net/netip/netip.go:238
					_go_fuzz_dep_.CoverTab[3630]++
//line /usr/local/go/src/net/netip/netip.go:238
					return c <= 'f'
//line /usr/local/go/src/net/netip/netip.go:238
					// _ = "end of CoverTab[3630]"
//line /usr/local/go/src/net/netip/netip.go:238
				}() {
//line /usr/local/go/src/net/netip/netip.go:238
					_go_fuzz_dep_.CoverTab[3631]++
											acc = (acc << 4) + uint32(c-'a'+10)
//line /usr/local/go/src/net/netip/netip.go:239
					// _ = "end of CoverTab[3631]"
				} else {
//line /usr/local/go/src/net/netip/netip.go:240
					_go_fuzz_dep_.CoverTab[3632]++
//line /usr/local/go/src/net/netip/netip.go:240
					if c >= 'A' && func() bool {
//line /usr/local/go/src/net/netip/netip.go:240
						_go_fuzz_dep_.CoverTab[3633]++
//line /usr/local/go/src/net/netip/netip.go:240
						return c <= 'F'
//line /usr/local/go/src/net/netip/netip.go:240
						// _ = "end of CoverTab[3633]"
//line /usr/local/go/src/net/netip/netip.go:240
					}() {
//line /usr/local/go/src/net/netip/netip.go:240
						_go_fuzz_dep_.CoverTab[3634]++
												acc = (acc << 4) + uint32(c-'A'+10)
//line /usr/local/go/src/net/netip/netip.go:241
						// _ = "end of CoverTab[3634]"
					} else {
//line /usr/local/go/src/net/netip/netip.go:242
						_go_fuzz_dep_.CoverTab[3635]++
												break
//line /usr/local/go/src/net/netip/netip.go:243
						// _ = "end of CoverTab[3635]"
					}
//line /usr/local/go/src/net/netip/netip.go:244
					// _ = "end of CoverTab[3632]"
//line /usr/local/go/src/net/netip/netip.go:244
				}
//line /usr/local/go/src/net/netip/netip.go:244
				// _ = "end of CoverTab[3629]"
//line /usr/local/go/src/net/netip/netip.go:244
			}
//line /usr/local/go/src/net/netip/netip.go:244
			// _ = "end of CoverTab[3625]"
//line /usr/local/go/src/net/netip/netip.go:244
			_go_fuzz_dep_.CoverTab[3626]++
									if acc > math.MaxUint16 {
//line /usr/local/go/src/net/netip/netip.go:245
				_go_fuzz_dep_.CoverTab[3636]++

										return Addr{}, parseAddrError{in: in, msg: "IPv6 field has value >=2^16", at: s}
//line /usr/local/go/src/net/netip/netip.go:247
				// _ = "end of CoverTab[3636]"
			} else {
//line /usr/local/go/src/net/netip/netip.go:248
				_go_fuzz_dep_.CoverTab[3637]++
//line /usr/local/go/src/net/netip/netip.go:248
				// _ = "end of CoverTab[3637]"
//line /usr/local/go/src/net/netip/netip.go:248
			}
//line /usr/local/go/src/net/netip/netip.go:248
			// _ = "end of CoverTab[3626]"
		}
//line /usr/local/go/src/net/netip/netip.go:249
		// _ = "end of CoverTab[3619]"
//line /usr/local/go/src/net/netip/netip.go:249
		_go_fuzz_dep_.CoverTab[3620]++
								if off == 0 {
//line /usr/local/go/src/net/netip/netip.go:250
			_go_fuzz_dep_.CoverTab[3638]++

									return Addr{}, parseAddrError{in: in, msg: "each colon-separated field must have at least one digit", at: s}
//line /usr/local/go/src/net/netip/netip.go:252
			// _ = "end of CoverTab[3638]"
		} else {
//line /usr/local/go/src/net/netip/netip.go:253
			_go_fuzz_dep_.CoverTab[3639]++
//line /usr/local/go/src/net/netip/netip.go:253
			// _ = "end of CoverTab[3639]"
//line /usr/local/go/src/net/netip/netip.go:253
		}
//line /usr/local/go/src/net/netip/netip.go:253
		// _ = "end of CoverTab[3620]"
//line /usr/local/go/src/net/netip/netip.go:253
		_go_fuzz_dep_.CoverTab[3621]++

//line /usr/local/go/src/net/netip/netip.go:256
		if off < len(s) && func() bool {
//line /usr/local/go/src/net/netip/netip.go:256
			_go_fuzz_dep_.CoverTab[3640]++
//line /usr/local/go/src/net/netip/netip.go:256
			return s[off] == '.'
//line /usr/local/go/src/net/netip/netip.go:256
			// _ = "end of CoverTab[3640]"
//line /usr/local/go/src/net/netip/netip.go:256
		}() {
//line /usr/local/go/src/net/netip/netip.go:256
			_go_fuzz_dep_.CoverTab[3641]++
									if ellipsis < 0 && func() bool {
//line /usr/local/go/src/net/netip/netip.go:257
				_go_fuzz_dep_.CoverTab[3645]++
//line /usr/local/go/src/net/netip/netip.go:257
				return i != 12
//line /usr/local/go/src/net/netip/netip.go:257
				// _ = "end of CoverTab[3645]"
//line /usr/local/go/src/net/netip/netip.go:257
			}() {
//line /usr/local/go/src/net/netip/netip.go:257
				_go_fuzz_dep_.CoverTab[3646]++

										return Addr{}, parseAddrError{in: in, msg: "embedded IPv4 address must replace the final 2 fields of the address", at: s}
//line /usr/local/go/src/net/netip/netip.go:259
				// _ = "end of CoverTab[3646]"
			} else {
//line /usr/local/go/src/net/netip/netip.go:260
				_go_fuzz_dep_.CoverTab[3647]++
//line /usr/local/go/src/net/netip/netip.go:260
				// _ = "end of CoverTab[3647]"
//line /usr/local/go/src/net/netip/netip.go:260
			}
//line /usr/local/go/src/net/netip/netip.go:260
			// _ = "end of CoverTab[3641]"
//line /usr/local/go/src/net/netip/netip.go:260
			_go_fuzz_dep_.CoverTab[3642]++
									if i+4 > 16 {
//line /usr/local/go/src/net/netip/netip.go:261
				_go_fuzz_dep_.CoverTab[3648]++

										return Addr{}, parseAddrError{in: in, msg: "too many hex fields to fit an embedded IPv4 at the end of the address", at: s}
//line /usr/local/go/src/net/netip/netip.go:263
				// _ = "end of CoverTab[3648]"
			} else {
//line /usr/local/go/src/net/netip/netip.go:264
				_go_fuzz_dep_.CoverTab[3649]++
//line /usr/local/go/src/net/netip/netip.go:264
				// _ = "end of CoverTab[3649]"
//line /usr/local/go/src/net/netip/netip.go:264
			}
//line /usr/local/go/src/net/netip/netip.go:264
			// _ = "end of CoverTab[3642]"
//line /usr/local/go/src/net/netip/netip.go:264
			_go_fuzz_dep_.CoverTab[3643]++

//line /usr/local/go/src/net/netip/netip.go:268
			ip4, err := parseIPv4(s)
			if err != nil {
//line /usr/local/go/src/net/netip/netip.go:269
				_go_fuzz_dep_.CoverTab[3650]++
										return Addr{}, parseAddrError{in: in, msg: err.Error(), at: s}
//line /usr/local/go/src/net/netip/netip.go:270
				// _ = "end of CoverTab[3650]"
			} else {
//line /usr/local/go/src/net/netip/netip.go:271
				_go_fuzz_dep_.CoverTab[3651]++
//line /usr/local/go/src/net/netip/netip.go:271
				// _ = "end of CoverTab[3651]"
//line /usr/local/go/src/net/netip/netip.go:271
			}
//line /usr/local/go/src/net/netip/netip.go:271
			// _ = "end of CoverTab[3643]"
//line /usr/local/go/src/net/netip/netip.go:271
			_go_fuzz_dep_.CoverTab[3644]++
									ip[i] = ip4.v4(0)
									ip[i+1] = ip4.v4(1)
									ip[i+2] = ip4.v4(2)
									ip[i+3] = ip4.v4(3)
									s = ""
									i += 4
									break
//line /usr/local/go/src/net/netip/netip.go:278
			// _ = "end of CoverTab[3644]"
		} else {
//line /usr/local/go/src/net/netip/netip.go:279
			_go_fuzz_dep_.CoverTab[3652]++
//line /usr/local/go/src/net/netip/netip.go:279
			// _ = "end of CoverTab[3652]"
//line /usr/local/go/src/net/netip/netip.go:279
		}
//line /usr/local/go/src/net/netip/netip.go:279
		// _ = "end of CoverTab[3621]"
//line /usr/local/go/src/net/netip/netip.go:279
		_go_fuzz_dep_.CoverTab[3622]++

//line /usr/local/go/src/net/netip/netip.go:282
		ip[i] = byte(acc >> 8)
								ip[i+1] = byte(acc)
								i += 2

//line /usr/local/go/src/net/netip/netip.go:287
		s = s[off:]
		if len(s) == 0 {
//line /usr/local/go/src/net/netip/netip.go:288
			_go_fuzz_dep_.CoverTab[3653]++
									break
//line /usr/local/go/src/net/netip/netip.go:289
			// _ = "end of CoverTab[3653]"
		} else {
//line /usr/local/go/src/net/netip/netip.go:290
			_go_fuzz_dep_.CoverTab[3654]++
//line /usr/local/go/src/net/netip/netip.go:290
			// _ = "end of CoverTab[3654]"
//line /usr/local/go/src/net/netip/netip.go:290
		}
//line /usr/local/go/src/net/netip/netip.go:290
		// _ = "end of CoverTab[3622]"
//line /usr/local/go/src/net/netip/netip.go:290
		_go_fuzz_dep_.CoverTab[3623]++

//line /usr/local/go/src/net/netip/netip.go:293
		if s[0] != ':' {
//line /usr/local/go/src/net/netip/netip.go:293
			_go_fuzz_dep_.CoverTab[3655]++
									return Addr{}, parseAddrError{in: in, msg: "unexpected character, want colon", at: s}
//line /usr/local/go/src/net/netip/netip.go:294
			// _ = "end of CoverTab[3655]"
		} else {
//line /usr/local/go/src/net/netip/netip.go:295
			_go_fuzz_dep_.CoverTab[3656]++
//line /usr/local/go/src/net/netip/netip.go:295
			if len(s) == 1 {
//line /usr/local/go/src/net/netip/netip.go:295
				_go_fuzz_dep_.CoverTab[3657]++
										return Addr{}, parseAddrError{in: in, msg: "colon must be followed by more characters", at: s}
//line /usr/local/go/src/net/netip/netip.go:296
				// _ = "end of CoverTab[3657]"
			} else {
//line /usr/local/go/src/net/netip/netip.go:297
				_go_fuzz_dep_.CoverTab[3658]++
//line /usr/local/go/src/net/netip/netip.go:297
				// _ = "end of CoverTab[3658]"
//line /usr/local/go/src/net/netip/netip.go:297
			}
//line /usr/local/go/src/net/netip/netip.go:297
			// _ = "end of CoverTab[3656]"
//line /usr/local/go/src/net/netip/netip.go:297
		}
//line /usr/local/go/src/net/netip/netip.go:297
		// _ = "end of CoverTab[3623]"
//line /usr/local/go/src/net/netip/netip.go:297
		_go_fuzz_dep_.CoverTab[3624]++
								s = s[1:]

//line /usr/local/go/src/net/netip/netip.go:301
		if s[0] == ':' {
//line /usr/local/go/src/net/netip/netip.go:301
			_go_fuzz_dep_.CoverTab[3659]++
									if ellipsis >= 0 {
//line /usr/local/go/src/net/netip/netip.go:302
				_go_fuzz_dep_.CoverTab[3661]++
										return Addr{}, parseAddrError{in: in, msg: "multiple :: in address", at: s}
//line /usr/local/go/src/net/netip/netip.go:303
				// _ = "end of CoverTab[3661]"
			} else {
//line /usr/local/go/src/net/netip/netip.go:304
				_go_fuzz_dep_.CoverTab[3662]++
//line /usr/local/go/src/net/netip/netip.go:304
				// _ = "end of CoverTab[3662]"
//line /usr/local/go/src/net/netip/netip.go:304
			}
//line /usr/local/go/src/net/netip/netip.go:304
			// _ = "end of CoverTab[3659]"
//line /usr/local/go/src/net/netip/netip.go:304
			_go_fuzz_dep_.CoverTab[3660]++
									ellipsis = i
									s = s[1:]
									if len(s) == 0 {
//line /usr/local/go/src/net/netip/netip.go:307
				_go_fuzz_dep_.CoverTab[3663]++
										break
//line /usr/local/go/src/net/netip/netip.go:308
				// _ = "end of CoverTab[3663]"
			} else {
//line /usr/local/go/src/net/netip/netip.go:309
				_go_fuzz_dep_.CoverTab[3664]++
//line /usr/local/go/src/net/netip/netip.go:309
				// _ = "end of CoverTab[3664]"
//line /usr/local/go/src/net/netip/netip.go:309
			}
//line /usr/local/go/src/net/netip/netip.go:309
			// _ = "end of CoverTab[3660]"
		} else {
//line /usr/local/go/src/net/netip/netip.go:310
			_go_fuzz_dep_.CoverTab[3665]++
//line /usr/local/go/src/net/netip/netip.go:310
			// _ = "end of CoverTab[3665]"
//line /usr/local/go/src/net/netip/netip.go:310
		}
//line /usr/local/go/src/net/netip/netip.go:310
		// _ = "end of CoverTab[3624]"
	}
//line /usr/local/go/src/net/netip/netip.go:311
	// _ = "end of CoverTab[3605]"
//line /usr/local/go/src/net/netip/netip.go:311
	_go_fuzz_dep_.CoverTab[3606]++

//line /usr/local/go/src/net/netip/netip.go:314
	if len(s) != 0 {
//line /usr/local/go/src/net/netip/netip.go:314
		_go_fuzz_dep_.CoverTab[3666]++
								return Addr{}, parseAddrError{in: in, msg: "trailing garbage after address", at: s}
//line /usr/local/go/src/net/netip/netip.go:315
		// _ = "end of CoverTab[3666]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:316
		_go_fuzz_dep_.CoverTab[3667]++
//line /usr/local/go/src/net/netip/netip.go:316
		// _ = "end of CoverTab[3667]"
//line /usr/local/go/src/net/netip/netip.go:316
	}
//line /usr/local/go/src/net/netip/netip.go:316
	// _ = "end of CoverTab[3606]"
//line /usr/local/go/src/net/netip/netip.go:316
	_go_fuzz_dep_.CoverTab[3607]++

//line /usr/local/go/src/net/netip/netip.go:319
	if i < 16 {
//line /usr/local/go/src/net/netip/netip.go:319
		_go_fuzz_dep_.CoverTab[3668]++
								if ellipsis < 0 {
//line /usr/local/go/src/net/netip/netip.go:320
			_go_fuzz_dep_.CoverTab[3671]++
									return Addr{}, parseAddrError{in: in, msg: "address string too short"}
//line /usr/local/go/src/net/netip/netip.go:321
			// _ = "end of CoverTab[3671]"
		} else {
//line /usr/local/go/src/net/netip/netip.go:322
			_go_fuzz_dep_.CoverTab[3672]++
//line /usr/local/go/src/net/netip/netip.go:322
			// _ = "end of CoverTab[3672]"
//line /usr/local/go/src/net/netip/netip.go:322
		}
//line /usr/local/go/src/net/netip/netip.go:322
		// _ = "end of CoverTab[3668]"
//line /usr/local/go/src/net/netip/netip.go:322
		_go_fuzz_dep_.CoverTab[3669]++
								n := 16 - i
								for j := i - 1; j >= ellipsis; j-- {
//line /usr/local/go/src/net/netip/netip.go:324
			_go_fuzz_dep_.CoverTab[3673]++
									ip[j+n] = ip[j]
//line /usr/local/go/src/net/netip/netip.go:325
			// _ = "end of CoverTab[3673]"
		}
//line /usr/local/go/src/net/netip/netip.go:326
		// _ = "end of CoverTab[3669]"
//line /usr/local/go/src/net/netip/netip.go:326
		_go_fuzz_dep_.CoverTab[3670]++
								for j := ellipsis + n - 1; j >= ellipsis; j-- {
//line /usr/local/go/src/net/netip/netip.go:327
			_go_fuzz_dep_.CoverTab[3674]++
									ip[j] = 0
//line /usr/local/go/src/net/netip/netip.go:328
			// _ = "end of CoverTab[3674]"
		}
//line /usr/local/go/src/net/netip/netip.go:329
		// _ = "end of CoverTab[3670]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:330
		_go_fuzz_dep_.CoverTab[3675]++
//line /usr/local/go/src/net/netip/netip.go:330
		if ellipsis >= 0 {
//line /usr/local/go/src/net/netip/netip.go:330
			_go_fuzz_dep_.CoverTab[3676]++

									return Addr{}, parseAddrError{in: in, msg: "the :: must expand to at least one field of zeros"}
//line /usr/local/go/src/net/netip/netip.go:332
			// _ = "end of CoverTab[3676]"
		} else {
//line /usr/local/go/src/net/netip/netip.go:333
			_go_fuzz_dep_.CoverTab[3677]++
//line /usr/local/go/src/net/netip/netip.go:333
			// _ = "end of CoverTab[3677]"
//line /usr/local/go/src/net/netip/netip.go:333
		}
//line /usr/local/go/src/net/netip/netip.go:333
		// _ = "end of CoverTab[3675]"
//line /usr/local/go/src/net/netip/netip.go:333
	}
//line /usr/local/go/src/net/netip/netip.go:333
	// _ = "end of CoverTab[3607]"
//line /usr/local/go/src/net/netip/netip.go:333
	_go_fuzz_dep_.CoverTab[3608]++
							return AddrFrom16(ip).WithZone(zone), nil
//line /usr/local/go/src/net/netip/netip.go:334
	// _ = "end of CoverTab[3608]"
}

// AddrFromSlice parses the 4- or 16-byte byte slice as an IPv4 or IPv6 address.
//line /usr/local/go/src/net/netip/netip.go:337
// Note that a net.IP can be passed directly as the []byte argument.
//line /usr/local/go/src/net/netip/netip.go:337
// If slice's length is not 4 or 16, AddrFromSlice returns Addr{}, false.
//line /usr/local/go/src/net/netip/netip.go:340
func AddrFromSlice(slice []byte) (ip Addr, ok bool) {
//line /usr/local/go/src/net/netip/netip.go:340
	_go_fuzz_dep_.CoverTab[3678]++
							switch len(slice) {
	case 4:
//line /usr/local/go/src/net/netip/netip.go:342
		_go_fuzz_dep_.CoverTab[3680]++
								return AddrFrom4([4]byte(slice)), true
//line /usr/local/go/src/net/netip/netip.go:343
		// _ = "end of CoverTab[3680]"
	case 16:
//line /usr/local/go/src/net/netip/netip.go:344
		_go_fuzz_dep_.CoverTab[3681]++
								return AddrFrom16([16]byte(slice)), true
//line /usr/local/go/src/net/netip/netip.go:345
		// _ = "end of CoverTab[3681]"
//line /usr/local/go/src/net/netip/netip.go:345
	default:
//line /usr/local/go/src/net/netip/netip.go:345
		_go_fuzz_dep_.CoverTab[3682]++
//line /usr/local/go/src/net/netip/netip.go:345
		// _ = "end of CoverTab[3682]"
	}
//line /usr/local/go/src/net/netip/netip.go:346
	// _ = "end of CoverTab[3678]"
//line /usr/local/go/src/net/netip/netip.go:346
	_go_fuzz_dep_.CoverTab[3679]++
							return Addr{}, false
//line /usr/local/go/src/net/netip/netip.go:347
	// _ = "end of CoverTab[3679]"
}

// v4 returns the i'th byte of ip. If ip is not an IPv4, v4 returns
//line /usr/local/go/src/net/netip/netip.go:350
// unspecified garbage.
//line /usr/local/go/src/net/netip/netip.go:352
func (ip Addr) v4(i uint8) uint8 {
//line /usr/local/go/src/net/netip/netip.go:352
	_go_fuzz_dep_.CoverTab[3683]++
							return uint8(ip.addr.lo >> ((3 - i) * 8))
//line /usr/local/go/src/net/netip/netip.go:353
	// _ = "end of CoverTab[3683]"
}

// v6 returns the i'th byte of ip. If ip is an IPv4 address, this
//line /usr/local/go/src/net/netip/netip.go:356
// accesses the IPv4-mapped IPv6 address form of the IP.
//line /usr/local/go/src/net/netip/netip.go:358
func (ip Addr) v6(i uint8) uint8 {
//line /usr/local/go/src/net/netip/netip.go:358
	_go_fuzz_dep_.CoverTab[3684]++
							return uint8(*(ip.addr.halves()[(i/8)%2]) >> ((7 - i%8) * 8))
//line /usr/local/go/src/net/netip/netip.go:359
	// _ = "end of CoverTab[3684]"
}

// v6u16 returns the i'th 16-bit word of ip. If ip is an IPv4 address,
//line /usr/local/go/src/net/netip/netip.go:362
// this accesses the IPv4-mapped IPv6 address form of the IP.
//line /usr/local/go/src/net/netip/netip.go:364
func (ip Addr) v6u16(i uint8) uint16 {
//line /usr/local/go/src/net/netip/netip.go:364
	_go_fuzz_dep_.CoverTab[3685]++
							return uint16(*(ip.addr.halves()[(i/4)%2]) >> ((3 - i%4) * 16))
//line /usr/local/go/src/net/netip/netip.go:365
	// _ = "end of CoverTab[3685]"
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
	_go_fuzz_dep_.CoverTab[3686]++

//line /usr/local/go/src/net/netip/netip.go:376
	return ip.z == z0
//line /usr/local/go/src/net/netip/netip.go:376
	// _ = "end of CoverTab[3686]"
}

// IsValid reports whether the Addr is an initialized address (not the zero Addr).
//line /usr/local/go/src/net/netip/netip.go:379
//
//line /usr/local/go/src/net/netip/netip.go:379
// Note that "0.0.0.0" and "::" are both valid values.
//line /usr/local/go/src/net/netip/netip.go:382
func (ip Addr) IsValid() bool {
//line /usr/local/go/src/net/netip/netip.go:382
	_go_fuzz_dep_.CoverTab[3687]++
//line /usr/local/go/src/net/netip/netip.go:382
	return ip.z != z0
//line /usr/local/go/src/net/netip/netip.go:382
	// _ = "end of CoverTab[3687]"
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
	_go_fuzz_dep_.CoverTab[3688]++
							switch ip.z {
	case z0:
//line /usr/local/go/src/net/netip/netip.go:391
		_go_fuzz_dep_.CoverTab[3690]++
								return 0
//line /usr/local/go/src/net/netip/netip.go:392
		// _ = "end of CoverTab[3690]"
	case z4:
//line /usr/local/go/src/net/netip/netip.go:393
		_go_fuzz_dep_.CoverTab[3691]++
								return 32
//line /usr/local/go/src/net/netip/netip.go:394
		// _ = "end of CoverTab[3691]"
//line /usr/local/go/src/net/netip/netip.go:394
	default:
//line /usr/local/go/src/net/netip/netip.go:394
		_go_fuzz_dep_.CoverTab[3692]++
//line /usr/local/go/src/net/netip/netip.go:394
		// _ = "end of CoverTab[3692]"
	}
//line /usr/local/go/src/net/netip/netip.go:395
	// _ = "end of CoverTab[3688]"
//line /usr/local/go/src/net/netip/netip.go:395
	_go_fuzz_dep_.CoverTab[3689]++
							return 128
//line /usr/local/go/src/net/netip/netip.go:396
	// _ = "end of CoverTab[3689]"
}

// Zone returns ip's IPv6 scoped addressing zone, if any.
func (ip Addr) Zone() string {
//line /usr/local/go/src/net/netip/netip.go:400
	_go_fuzz_dep_.CoverTab[3693]++
							if ip.z == nil {
//line /usr/local/go/src/net/netip/netip.go:401
		_go_fuzz_dep_.CoverTab[3695]++
								return ""
//line /usr/local/go/src/net/netip/netip.go:402
		// _ = "end of CoverTab[3695]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:403
		_go_fuzz_dep_.CoverTab[3696]++
//line /usr/local/go/src/net/netip/netip.go:403
		// _ = "end of CoverTab[3696]"
//line /usr/local/go/src/net/netip/netip.go:403
	}
//line /usr/local/go/src/net/netip/netip.go:403
	// _ = "end of CoverTab[3693]"
//line /usr/local/go/src/net/netip/netip.go:403
	_go_fuzz_dep_.CoverTab[3694]++
							zone, _ := ip.z.Get().(string)
							return zone
//line /usr/local/go/src/net/netip/netip.go:405
	// _ = "end of CoverTab[3694]"
}

// Compare returns an integer comparing two IPs.
//line /usr/local/go/src/net/netip/netip.go:408
// The result will be 0 if ip == ip2, -1 if ip < ip2, and +1 if ip > ip2.
//line /usr/local/go/src/net/netip/netip.go:408
// The definition of "less than" is the same as the Less method.
//line /usr/local/go/src/net/netip/netip.go:411
func (ip Addr) Compare(ip2 Addr) int {
//line /usr/local/go/src/net/netip/netip.go:411
	_go_fuzz_dep_.CoverTab[3697]++
							f1, f2 := ip.BitLen(), ip2.BitLen()
							if f1 < f2 {
//line /usr/local/go/src/net/netip/netip.go:413
		_go_fuzz_dep_.CoverTab[3705]++
								return -1
//line /usr/local/go/src/net/netip/netip.go:414
		// _ = "end of CoverTab[3705]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:415
		_go_fuzz_dep_.CoverTab[3706]++
//line /usr/local/go/src/net/netip/netip.go:415
		// _ = "end of CoverTab[3706]"
//line /usr/local/go/src/net/netip/netip.go:415
	}
//line /usr/local/go/src/net/netip/netip.go:415
	// _ = "end of CoverTab[3697]"
//line /usr/local/go/src/net/netip/netip.go:415
	_go_fuzz_dep_.CoverTab[3698]++
							if f1 > f2 {
//line /usr/local/go/src/net/netip/netip.go:416
		_go_fuzz_dep_.CoverTab[3707]++
								return 1
//line /usr/local/go/src/net/netip/netip.go:417
		// _ = "end of CoverTab[3707]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:418
		_go_fuzz_dep_.CoverTab[3708]++
//line /usr/local/go/src/net/netip/netip.go:418
		// _ = "end of CoverTab[3708]"
//line /usr/local/go/src/net/netip/netip.go:418
	}
//line /usr/local/go/src/net/netip/netip.go:418
	// _ = "end of CoverTab[3698]"
//line /usr/local/go/src/net/netip/netip.go:418
	_go_fuzz_dep_.CoverTab[3699]++
							hi1, hi2 := ip.addr.hi, ip2.addr.hi
							if hi1 < hi2 {
//line /usr/local/go/src/net/netip/netip.go:420
		_go_fuzz_dep_.CoverTab[3709]++
								return -1
//line /usr/local/go/src/net/netip/netip.go:421
		// _ = "end of CoverTab[3709]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:422
		_go_fuzz_dep_.CoverTab[3710]++
//line /usr/local/go/src/net/netip/netip.go:422
		// _ = "end of CoverTab[3710]"
//line /usr/local/go/src/net/netip/netip.go:422
	}
//line /usr/local/go/src/net/netip/netip.go:422
	// _ = "end of CoverTab[3699]"
//line /usr/local/go/src/net/netip/netip.go:422
	_go_fuzz_dep_.CoverTab[3700]++
							if hi1 > hi2 {
//line /usr/local/go/src/net/netip/netip.go:423
		_go_fuzz_dep_.CoverTab[3711]++
								return 1
//line /usr/local/go/src/net/netip/netip.go:424
		// _ = "end of CoverTab[3711]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:425
		_go_fuzz_dep_.CoverTab[3712]++
//line /usr/local/go/src/net/netip/netip.go:425
		// _ = "end of CoverTab[3712]"
//line /usr/local/go/src/net/netip/netip.go:425
	}
//line /usr/local/go/src/net/netip/netip.go:425
	// _ = "end of CoverTab[3700]"
//line /usr/local/go/src/net/netip/netip.go:425
	_go_fuzz_dep_.CoverTab[3701]++
							lo1, lo2 := ip.addr.lo, ip2.addr.lo
							if lo1 < lo2 {
//line /usr/local/go/src/net/netip/netip.go:427
		_go_fuzz_dep_.CoverTab[3713]++
								return -1
//line /usr/local/go/src/net/netip/netip.go:428
		// _ = "end of CoverTab[3713]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:429
		_go_fuzz_dep_.CoverTab[3714]++
//line /usr/local/go/src/net/netip/netip.go:429
		// _ = "end of CoverTab[3714]"
//line /usr/local/go/src/net/netip/netip.go:429
	}
//line /usr/local/go/src/net/netip/netip.go:429
	// _ = "end of CoverTab[3701]"
//line /usr/local/go/src/net/netip/netip.go:429
	_go_fuzz_dep_.CoverTab[3702]++
							if lo1 > lo2 {
//line /usr/local/go/src/net/netip/netip.go:430
		_go_fuzz_dep_.CoverTab[3715]++
								return 1
//line /usr/local/go/src/net/netip/netip.go:431
		// _ = "end of CoverTab[3715]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:432
		_go_fuzz_dep_.CoverTab[3716]++
//line /usr/local/go/src/net/netip/netip.go:432
		// _ = "end of CoverTab[3716]"
//line /usr/local/go/src/net/netip/netip.go:432
	}
//line /usr/local/go/src/net/netip/netip.go:432
	// _ = "end of CoverTab[3702]"
//line /usr/local/go/src/net/netip/netip.go:432
	_go_fuzz_dep_.CoverTab[3703]++
							if ip.Is6() {
//line /usr/local/go/src/net/netip/netip.go:433
		_go_fuzz_dep_.CoverTab[3717]++
								za, zb := ip.Zone(), ip2.Zone()
								if za < zb {
//line /usr/local/go/src/net/netip/netip.go:435
			_go_fuzz_dep_.CoverTab[3719]++
									return -1
//line /usr/local/go/src/net/netip/netip.go:436
			// _ = "end of CoverTab[3719]"
		} else {
//line /usr/local/go/src/net/netip/netip.go:437
			_go_fuzz_dep_.CoverTab[3720]++
//line /usr/local/go/src/net/netip/netip.go:437
			// _ = "end of CoverTab[3720]"
//line /usr/local/go/src/net/netip/netip.go:437
		}
//line /usr/local/go/src/net/netip/netip.go:437
		// _ = "end of CoverTab[3717]"
//line /usr/local/go/src/net/netip/netip.go:437
		_go_fuzz_dep_.CoverTab[3718]++
								if za > zb {
//line /usr/local/go/src/net/netip/netip.go:438
			_go_fuzz_dep_.CoverTab[3721]++
									return 1
//line /usr/local/go/src/net/netip/netip.go:439
			// _ = "end of CoverTab[3721]"
		} else {
//line /usr/local/go/src/net/netip/netip.go:440
			_go_fuzz_dep_.CoverTab[3722]++
//line /usr/local/go/src/net/netip/netip.go:440
			// _ = "end of CoverTab[3722]"
//line /usr/local/go/src/net/netip/netip.go:440
		}
//line /usr/local/go/src/net/netip/netip.go:440
		// _ = "end of CoverTab[3718]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:441
		_go_fuzz_dep_.CoverTab[3723]++
//line /usr/local/go/src/net/netip/netip.go:441
		// _ = "end of CoverTab[3723]"
//line /usr/local/go/src/net/netip/netip.go:441
	}
//line /usr/local/go/src/net/netip/netip.go:441
	// _ = "end of CoverTab[3703]"
//line /usr/local/go/src/net/netip/netip.go:441
	_go_fuzz_dep_.CoverTab[3704]++
							return 0
//line /usr/local/go/src/net/netip/netip.go:442
	// _ = "end of CoverTab[3704]"
}

// Less reports whether ip sorts before ip2.
//line /usr/local/go/src/net/netip/netip.go:445
// IP addresses sort first by length, then their address.
//line /usr/local/go/src/net/netip/netip.go:445
// IPv6 addresses with zones sort just after the same address without a zone.
//line /usr/local/go/src/net/netip/netip.go:448
func (ip Addr) Less(ip2 Addr) bool {
//line /usr/local/go/src/net/netip/netip.go:448
	_go_fuzz_dep_.CoverTab[3724]++
//line /usr/local/go/src/net/netip/netip.go:448
	return ip.Compare(ip2) == -1
//line /usr/local/go/src/net/netip/netip.go:448
	// _ = "end of CoverTab[3724]"
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
	_go_fuzz_dep_.CoverTab[3725]++
							return ip.z == z4
//line /usr/local/go/src/net/netip/netip.go:454
	// _ = "end of CoverTab[3725]"
}

// Is4In6 reports whether ip is an IPv4-mapped IPv6 address.
func (ip Addr) Is4In6() bool {
//line /usr/local/go/src/net/netip/netip.go:458
	_go_fuzz_dep_.CoverTab[3726]++
							return ip.Is6() && func() bool {
//line /usr/local/go/src/net/netip/netip.go:459
		_go_fuzz_dep_.CoverTab[3727]++
//line /usr/local/go/src/net/netip/netip.go:459
		return ip.addr.hi == 0
//line /usr/local/go/src/net/netip/netip.go:459
		// _ = "end of CoverTab[3727]"
//line /usr/local/go/src/net/netip/netip.go:459
	}() && func() bool {
//line /usr/local/go/src/net/netip/netip.go:459
		_go_fuzz_dep_.CoverTab[3728]++
//line /usr/local/go/src/net/netip/netip.go:459
		return ip.addr.lo>>32 == 0xffff
//line /usr/local/go/src/net/netip/netip.go:459
		// _ = "end of CoverTab[3728]"
//line /usr/local/go/src/net/netip/netip.go:459
	}()
//line /usr/local/go/src/net/netip/netip.go:459
	// _ = "end of CoverTab[3726]"
}

// Is6 reports whether ip is an IPv6 address, including IPv4-mapped
//line /usr/local/go/src/net/netip/netip.go:462
// IPv6 addresses.
//line /usr/local/go/src/net/netip/netip.go:464
func (ip Addr) Is6() bool {
//line /usr/local/go/src/net/netip/netip.go:464
	_go_fuzz_dep_.CoverTab[3729]++
							return ip.z != z0 && func() bool {
//line /usr/local/go/src/net/netip/netip.go:465
		_go_fuzz_dep_.CoverTab[3730]++
//line /usr/local/go/src/net/netip/netip.go:465
		return ip.z != z4
//line /usr/local/go/src/net/netip/netip.go:465
		// _ = "end of CoverTab[3730]"
//line /usr/local/go/src/net/netip/netip.go:465
	}()
//line /usr/local/go/src/net/netip/netip.go:465
	// _ = "end of CoverTab[3729]"
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
	_go_fuzz_dep_.CoverTab[3731]++
							if ip.Is4In6() {
//line /usr/local/go/src/net/netip/netip.go:473
		_go_fuzz_dep_.CoverTab[3733]++
								ip.z = z4
//line /usr/local/go/src/net/netip/netip.go:474
		// _ = "end of CoverTab[3733]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:475
		_go_fuzz_dep_.CoverTab[3734]++
//line /usr/local/go/src/net/netip/netip.go:475
		// _ = "end of CoverTab[3734]"
//line /usr/local/go/src/net/netip/netip.go:475
	}
//line /usr/local/go/src/net/netip/netip.go:475
	// _ = "end of CoverTab[3731]"
//line /usr/local/go/src/net/netip/netip.go:475
	_go_fuzz_dep_.CoverTab[3732]++
							return ip
//line /usr/local/go/src/net/netip/netip.go:476
	// _ = "end of CoverTab[3732]"
}

// WithZone returns an IP that's the same as ip but with the provided
//line /usr/local/go/src/net/netip/netip.go:479
// zone. If zone is empty, the zone is removed. If ip is an IPv4
//line /usr/local/go/src/net/netip/netip.go:479
// address, WithZone is a no-op and returns ip unchanged.
//line /usr/local/go/src/net/netip/netip.go:482
func (ip Addr) WithZone(zone string) Addr {
//line /usr/local/go/src/net/netip/netip.go:482
	_go_fuzz_dep_.CoverTab[3735]++
							if !ip.Is6() {
//line /usr/local/go/src/net/netip/netip.go:483
		_go_fuzz_dep_.CoverTab[3738]++
								return ip
//line /usr/local/go/src/net/netip/netip.go:484
		// _ = "end of CoverTab[3738]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:485
		_go_fuzz_dep_.CoverTab[3739]++
//line /usr/local/go/src/net/netip/netip.go:485
		// _ = "end of CoverTab[3739]"
//line /usr/local/go/src/net/netip/netip.go:485
	}
//line /usr/local/go/src/net/netip/netip.go:485
	// _ = "end of CoverTab[3735]"
//line /usr/local/go/src/net/netip/netip.go:485
	_go_fuzz_dep_.CoverTab[3736]++
							if zone == "" {
//line /usr/local/go/src/net/netip/netip.go:486
		_go_fuzz_dep_.CoverTab[3740]++
								ip.z = z6noz
								return ip
//line /usr/local/go/src/net/netip/netip.go:488
		// _ = "end of CoverTab[3740]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:489
		_go_fuzz_dep_.CoverTab[3741]++
//line /usr/local/go/src/net/netip/netip.go:489
		// _ = "end of CoverTab[3741]"
//line /usr/local/go/src/net/netip/netip.go:489
	}
//line /usr/local/go/src/net/netip/netip.go:489
	// _ = "end of CoverTab[3736]"
//line /usr/local/go/src/net/netip/netip.go:489
	_go_fuzz_dep_.CoverTab[3737]++
							ip.z = intern.GetByString(zone)
							return ip
//line /usr/local/go/src/net/netip/netip.go:491
	// _ = "end of CoverTab[3737]"
}

// withoutZone unconditionally strips the zone from ip.
//line /usr/local/go/src/net/netip/netip.go:494
// It's similar to WithZone, but small enough to be inlinable.
//line /usr/local/go/src/net/netip/netip.go:496
func (ip Addr) withoutZone() Addr {
//line /usr/local/go/src/net/netip/netip.go:496
	_go_fuzz_dep_.CoverTab[3742]++
							if !ip.Is6() {
//line /usr/local/go/src/net/netip/netip.go:497
		_go_fuzz_dep_.CoverTab[3744]++
								return ip
//line /usr/local/go/src/net/netip/netip.go:498
		// _ = "end of CoverTab[3744]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:499
		_go_fuzz_dep_.CoverTab[3745]++
//line /usr/local/go/src/net/netip/netip.go:499
		// _ = "end of CoverTab[3745]"
//line /usr/local/go/src/net/netip/netip.go:499
	}
//line /usr/local/go/src/net/netip/netip.go:499
	// _ = "end of CoverTab[3742]"
//line /usr/local/go/src/net/netip/netip.go:499
	_go_fuzz_dep_.CoverTab[3743]++
							ip.z = z6noz
							return ip
//line /usr/local/go/src/net/netip/netip.go:501
	// _ = "end of CoverTab[3743]"
}

// hasZone reports whether ip has an IPv6 zone.
func (ip Addr) hasZone() bool {
//line /usr/local/go/src/net/netip/netip.go:505
	_go_fuzz_dep_.CoverTab[3746]++
							return ip.z != z0 && func() bool {
//line /usr/local/go/src/net/netip/netip.go:506
		_go_fuzz_dep_.CoverTab[3747]++
//line /usr/local/go/src/net/netip/netip.go:506
		return ip.z != z4
//line /usr/local/go/src/net/netip/netip.go:506
		// _ = "end of CoverTab[3747]"
//line /usr/local/go/src/net/netip/netip.go:506
	}() && func() bool {
//line /usr/local/go/src/net/netip/netip.go:506
		_go_fuzz_dep_.CoverTab[3748]++
//line /usr/local/go/src/net/netip/netip.go:506
		return ip.z != z6noz
//line /usr/local/go/src/net/netip/netip.go:506
		// _ = "end of CoverTab[3748]"
//line /usr/local/go/src/net/netip/netip.go:506
	}()
//line /usr/local/go/src/net/netip/netip.go:506
	// _ = "end of CoverTab[3746]"
}

// IsLinkLocalUnicast reports whether ip is a link-local unicast address.
func (ip Addr) IsLinkLocalUnicast() bool {
//line /usr/local/go/src/net/netip/netip.go:510
	_go_fuzz_dep_.CoverTab[3749]++

//line /usr/local/go/src/net/netip/netip.go:513
	if ip.Is4() {
//line /usr/local/go/src/net/netip/netip.go:513
		_go_fuzz_dep_.CoverTab[3752]++
								return ip.v4(0) == 169 && func() bool {
//line /usr/local/go/src/net/netip/netip.go:514
			_go_fuzz_dep_.CoverTab[3753]++
//line /usr/local/go/src/net/netip/netip.go:514
			return ip.v4(1) == 254
//line /usr/local/go/src/net/netip/netip.go:514
			// _ = "end of CoverTab[3753]"
//line /usr/local/go/src/net/netip/netip.go:514
		}()
//line /usr/local/go/src/net/netip/netip.go:514
		// _ = "end of CoverTab[3752]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:515
		_go_fuzz_dep_.CoverTab[3754]++
//line /usr/local/go/src/net/netip/netip.go:515
		// _ = "end of CoverTab[3754]"
//line /usr/local/go/src/net/netip/netip.go:515
	}
//line /usr/local/go/src/net/netip/netip.go:515
	// _ = "end of CoverTab[3749]"
//line /usr/local/go/src/net/netip/netip.go:515
	_go_fuzz_dep_.CoverTab[3750]++

//line /usr/local/go/src/net/netip/netip.go:518
	if ip.Is6() {
//line /usr/local/go/src/net/netip/netip.go:518
		_go_fuzz_dep_.CoverTab[3755]++
								return ip.v6u16(0)&0xffc0 == 0xfe80
//line /usr/local/go/src/net/netip/netip.go:519
		// _ = "end of CoverTab[3755]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:520
		_go_fuzz_dep_.CoverTab[3756]++
//line /usr/local/go/src/net/netip/netip.go:520
		// _ = "end of CoverTab[3756]"
//line /usr/local/go/src/net/netip/netip.go:520
	}
//line /usr/local/go/src/net/netip/netip.go:520
	// _ = "end of CoverTab[3750]"
//line /usr/local/go/src/net/netip/netip.go:520
	_go_fuzz_dep_.CoverTab[3751]++
							return false
//line /usr/local/go/src/net/netip/netip.go:521
	// _ = "end of CoverTab[3751]"
}

// IsLoopback reports whether ip is a loopback address.
func (ip Addr) IsLoopback() bool {
//line /usr/local/go/src/net/netip/netip.go:525
	_go_fuzz_dep_.CoverTab[3757]++

//line /usr/local/go/src/net/netip/netip.go:528
	if ip.Is4() {
//line /usr/local/go/src/net/netip/netip.go:528
		_go_fuzz_dep_.CoverTab[3760]++
								return ip.v4(0) == 127
//line /usr/local/go/src/net/netip/netip.go:529
		// _ = "end of CoverTab[3760]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:530
		_go_fuzz_dep_.CoverTab[3761]++
//line /usr/local/go/src/net/netip/netip.go:530
		// _ = "end of CoverTab[3761]"
//line /usr/local/go/src/net/netip/netip.go:530
	}
//line /usr/local/go/src/net/netip/netip.go:530
	// _ = "end of CoverTab[3757]"
//line /usr/local/go/src/net/netip/netip.go:530
	_go_fuzz_dep_.CoverTab[3758]++

//line /usr/local/go/src/net/netip/netip.go:533
	if ip.Is6() {
//line /usr/local/go/src/net/netip/netip.go:533
		_go_fuzz_dep_.CoverTab[3762]++
								return ip.addr.hi == 0 && func() bool {
//line /usr/local/go/src/net/netip/netip.go:534
			_go_fuzz_dep_.CoverTab[3763]++
//line /usr/local/go/src/net/netip/netip.go:534
			return ip.addr.lo == 1
//line /usr/local/go/src/net/netip/netip.go:534
			// _ = "end of CoverTab[3763]"
//line /usr/local/go/src/net/netip/netip.go:534
		}()
//line /usr/local/go/src/net/netip/netip.go:534
		// _ = "end of CoverTab[3762]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:535
		_go_fuzz_dep_.CoverTab[3764]++
//line /usr/local/go/src/net/netip/netip.go:535
		// _ = "end of CoverTab[3764]"
//line /usr/local/go/src/net/netip/netip.go:535
	}
//line /usr/local/go/src/net/netip/netip.go:535
	// _ = "end of CoverTab[3758]"
//line /usr/local/go/src/net/netip/netip.go:535
	_go_fuzz_dep_.CoverTab[3759]++
							return false
//line /usr/local/go/src/net/netip/netip.go:536
	// _ = "end of CoverTab[3759]"
}

// IsMulticast reports whether ip is a multicast address.
func (ip Addr) IsMulticast() bool {
//line /usr/local/go/src/net/netip/netip.go:540
	_go_fuzz_dep_.CoverTab[3765]++

//line /usr/local/go/src/net/netip/netip.go:543
	if ip.Is4() {
//line /usr/local/go/src/net/netip/netip.go:543
		_go_fuzz_dep_.CoverTab[3768]++
								return ip.v4(0)&0xf0 == 0xe0
//line /usr/local/go/src/net/netip/netip.go:544
		// _ = "end of CoverTab[3768]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:545
		_go_fuzz_dep_.CoverTab[3769]++
//line /usr/local/go/src/net/netip/netip.go:545
		// _ = "end of CoverTab[3769]"
//line /usr/local/go/src/net/netip/netip.go:545
	}
//line /usr/local/go/src/net/netip/netip.go:545
	// _ = "end of CoverTab[3765]"
//line /usr/local/go/src/net/netip/netip.go:545
	_go_fuzz_dep_.CoverTab[3766]++

//line /usr/local/go/src/net/netip/netip.go:548
	if ip.Is6() {
//line /usr/local/go/src/net/netip/netip.go:548
		_go_fuzz_dep_.CoverTab[3770]++
								return ip.addr.hi>>(64-8) == 0xff
//line /usr/local/go/src/net/netip/netip.go:549
		// _ = "end of CoverTab[3770]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:550
		_go_fuzz_dep_.CoverTab[3771]++
//line /usr/local/go/src/net/netip/netip.go:550
		// _ = "end of CoverTab[3771]"
//line /usr/local/go/src/net/netip/netip.go:550
	}
//line /usr/local/go/src/net/netip/netip.go:550
	// _ = "end of CoverTab[3766]"
//line /usr/local/go/src/net/netip/netip.go:550
	_go_fuzz_dep_.CoverTab[3767]++
							return false
//line /usr/local/go/src/net/netip/netip.go:551
	// _ = "end of CoverTab[3767]"
}

// IsInterfaceLocalMulticast reports whether ip is an IPv6 interface-local
//line /usr/local/go/src/net/netip/netip.go:554
// multicast address.
//line /usr/local/go/src/net/netip/netip.go:556
func (ip Addr) IsInterfaceLocalMulticast() bool {
//line /usr/local/go/src/net/netip/netip.go:556
	_go_fuzz_dep_.CoverTab[3772]++

//line /usr/local/go/src/net/netip/netip.go:559
	if ip.Is6() {
//line /usr/local/go/src/net/netip/netip.go:559
		_go_fuzz_dep_.CoverTab[3774]++
								return ip.v6u16(0)&0xff0f == 0xff01
//line /usr/local/go/src/net/netip/netip.go:560
		// _ = "end of CoverTab[3774]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:561
		_go_fuzz_dep_.CoverTab[3775]++
//line /usr/local/go/src/net/netip/netip.go:561
		// _ = "end of CoverTab[3775]"
//line /usr/local/go/src/net/netip/netip.go:561
	}
//line /usr/local/go/src/net/netip/netip.go:561
	// _ = "end of CoverTab[3772]"
//line /usr/local/go/src/net/netip/netip.go:561
	_go_fuzz_dep_.CoverTab[3773]++
							return false
//line /usr/local/go/src/net/netip/netip.go:562
	// _ = "end of CoverTab[3773]"
}

// IsLinkLocalMulticast reports whether ip is a link-local multicast address.
func (ip Addr) IsLinkLocalMulticast() bool {
//line /usr/local/go/src/net/netip/netip.go:566
	_go_fuzz_dep_.CoverTab[3776]++

//line /usr/local/go/src/net/netip/netip.go:569
	if ip.Is4() {
//line /usr/local/go/src/net/netip/netip.go:569
		_go_fuzz_dep_.CoverTab[3779]++
								return ip.v4(0) == 224 && func() bool {
//line /usr/local/go/src/net/netip/netip.go:570
			_go_fuzz_dep_.CoverTab[3780]++
//line /usr/local/go/src/net/netip/netip.go:570
			return ip.v4(1) == 0
//line /usr/local/go/src/net/netip/netip.go:570
			// _ = "end of CoverTab[3780]"
//line /usr/local/go/src/net/netip/netip.go:570
		}() && func() bool {
//line /usr/local/go/src/net/netip/netip.go:570
			_go_fuzz_dep_.CoverTab[3781]++
//line /usr/local/go/src/net/netip/netip.go:570
			return ip.v4(2) == 0
//line /usr/local/go/src/net/netip/netip.go:570
			// _ = "end of CoverTab[3781]"
//line /usr/local/go/src/net/netip/netip.go:570
		}()
//line /usr/local/go/src/net/netip/netip.go:570
		// _ = "end of CoverTab[3779]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:571
		_go_fuzz_dep_.CoverTab[3782]++
//line /usr/local/go/src/net/netip/netip.go:571
		// _ = "end of CoverTab[3782]"
//line /usr/local/go/src/net/netip/netip.go:571
	}
//line /usr/local/go/src/net/netip/netip.go:571
	// _ = "end of CoverTab[3776]"
//line /usr/local/go/src/net/netip/netip.go:571
	_go_fuzz_dep_.CoverTab[3777]++

//line /usr/local/go/src/net/netip/netip.go:574
	if ip.Is6() {
//line /usr/local/go/src/net/netip/netip.go:574
		_go_fuzz_dep_.CoverTab[3783]++
								return ip.v6u16(0)&0xff0f == 0xff02
//line /usr/local/go/src/net/netip/netip.go:575
		// _ = "end of CoverTab[3783]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:576
		_go_fuzz_dep_.CoverTab[3784]++
//line /usr/local/go/src/net/netip/netip.go:576
		// _ = "end of CoverTab[3784]"
//line /usr/local/go/src/net/netip/netip.go:576
	}
//line /usr/local/go/src/net/netip/netip.go:576
	// _ = "end of CoverTab[3777]"
//line /usr/local/go/src/net/netip/netip.go:576
	_go_fuzz_dep_.CoverTab[3778]++
							return false
//line /usr/local/go/src/net/netip/netip.go:577
	// _ = "end of CoverTab[3778]"
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
	_go_fuzz_dep_.CoverTab[3785]++
							if ip.z == z0 {
//line /usr/local/go/src/net/netip/netip.go:590
		_go_fuzz_dep_.CoverTab[3788]++

								return false
//line /usr/local/go/src/net/netip/netip.go:592
		// _ = "end of CoverTab[3788]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:593
		_go_fuzz_dep_.CoverTab[3789]++
//line /usr/local/go/src/net/netip/netip.go:593
		// _ = "end of CoverTab[3789]"
//line /usr/local/go/src/net/netip/netip.go:593
	}
//line /usr/local/go/src/net/netip/netip.go:593
	// _ = "end of CoverTab[3785]"
//line /usr/local/go/src/net/netip/netip.go:593
	_go_fuzz_dep_.CoverTab[3786]++

//line /usr/local/go/src/net/netip/netip.go:597
	if ip.Is4() && func() bool {
//line /usr/local/go/src/net/netip/netip.go:597
		_go_fuzz_dep_.CoverTab[3790]++
//line /usr/local/go/src/net/netip/netip.go:597
		return (ip == IPv4Unspecified() || func() bool {
//line /usr/local/go/src/net/netip/netip.go:597
			_go_fuzz_dep_.CoverTab[3791]++
//line /usr/local/go/src/net/netip/netip.go:597
			return ip == AddrFrom4([4]byte{255, 255, 255, 255})
//line /usr/local/go/src/net/netip/netip.go:597
			// _ = "end of CoverTab[3791]"
//line /usr/local/go/src/net/netip/netip.go:597
		}())
//line /usr/local/go/src/net/netip/netip.go:597
		// _ = "end of CoverTab[3790]"
//line /usr/local/go/src/net/netip/netip.go:597
	}() {
//line /usr/local/go/src/net/netip/netip.go:597
		_go_fuzz_dep_.CoverTab[3792]++
								return false
//line /usr/local/go/src/net/netip/netip.go:598
		// _ = "end of CoverTab[3792]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:599
		_go_fuzz_dep_.CoverTab[3793]++
//line /usr/local/go/src/net/netip/netip.go:599
		// _ = "end of CoverTab[3793]"
//line /usr/local/go/src/net/netip/netip.go:599
	}
//line /usr/local/go/src/net/netip/netip.go:599
	// _ = "end of CoverTab[3786]"
//line /usr/local/go/src/net/netip/netip.go:599
	_go_fuzz_dep_.CoverTab[3787]++

							return ip != IPv6Unspecified() && func() bool {
//line /usr/local/go/src/net/netip/netip.go:601
		_go_fuzz_dep_.CoverTab[3794]++
//line /usr/local/go/src/net/netip/netip.go:601
		return !ip.IsLoopback()
								// _ = "end of CoverTab[3794]"
//line /usr/local/go/src/net/netip/netip.go:602
	}() && func() bool {
//line /usr/local/go/src/net/netip/netip.go:602
		_go_fuzz_dep_.CoverTab[3795]++
//line /usr/local/go/src/net/netip/netip.go:602
		return !ip.IsMulticast()
								// _ = "end of CoverTab[3795]"
//line /usr/local/go/src/net/netip/netip.go:603
	}() && func() bool {
//line /usr/local/go/src/net/netip/netip.go:603
		_go_fuzz_dep_.CoverTab[3796]++
//line /usr/local/go/src/net/netip/netip.go:603
		return !ip.IsLinkLocalUnicast()
								// _ = "end of CoverTab[3796]"
//line /usr/local/go/src/net/netip/netip.go:604
	}()
//line /usr/local/go/src/net/netip/netip.go:604
	// _ = "end of CoverTab[3787]"
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
	_go_fuzz_dep_.CoverTab[3797]++

							if ip.Is4() {
//line /usr/local/go/src/net/netip/netip.go:613
		_go_fuzz_dep_.CoverTab[3800]++

//line /usr/local/go/src/net/netip/netip.go:616
		return ip.v4(0) == 10 || func() bool {
//line /usr/local/go/src/net/netip/netip.go:616
			_go_fuzz_dep_.CoverTab[3801]++
//line /usr/local/go/src/net/netip/netip.go:616
			return (ip.v4(0) == 172 && func() bool {
										_go_fuzz_dep_.CoverTab[3802]++
//line /usr/local/go/src/net/netip/netip.go:617
				return ip.v4(1)&0xf0 == 16
//line /usr/local/go/src/net/netip/netip.go:617
				// _ = "end of CoverTab[3802]"
//line /usr/local/go/src/net/netip/netip.go:617
			}())
//line /usr/local/go/src/net/netip/netip.go:617
			// _ = "end of CoverTab[3801]"
//line /usr/local/go/src/net/netip/netip.go:617
		}() || func() bool {
//line /usr/local/go/src/net/netip/netip.go:617
			_go_fuzz_dep_.CoverTab[3803]++
//line /usr/local/go/src/net/netip/netip.go:617
			return (ip.v4(0) == 192 && func() bool {
										_go_fuzz_dep_.CoverTab[3804]++
//line /usr/local/go/src/net/netip/netip.go:618
				return ip.v4(1) == 168
//line /usr/local/go/src/net/netip/netip.go:618
				// _ = "end of CoverTab[3804]"
//line /usr/local/go/src/net/netip/netip.go:618
			}())
//line /usr/local/go/src/net/netip/netip.go:618
			// _ = "end of CoverTab[3803]"
//line /usr/local/go/src/net/netip/netip.go:618
		}()
//line /usr/local/go/src/net/netip/netip.go:618
		// _ = "end of CoverTab[3800]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:619
		_go_fuzz_dep_.CoverTab[3805]++
//line /usr/local/go/src/net/netip/netip.go:619
		// _ = "end of CoverTab[3805]"
//line /usr/local/go/src/net/netip/netip.go:619
	}
//line /usr/local/go/src/net/netip/netip.go:619
	// _ = "end of CoverTab[3797]"
//line /usr/local/go/src/net/netip/netip.go:619
	_go_fuzz_dep_.CoverTab[3798]++

							if ip.Is6() {
//line /usr/local/go/src/net/netip/netip.go:621
		_go_fuzz_dep_.CoverTab[3806]++

//line /usr/local/go/src/net/netip/netip.go:624
		return ip.v6(0)&0xfe == 0xfc
//line /usr/local/go/src/net/netip/netip.go:624
		// _ = "end of CoverTab[3806]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:625
		_go_fuzz_dep_.CoverTab[3807]++
//line /usr/local/go/src/net/netip/netip.go:625
		// _ = "end of CoverTab[3807]"
//line /usr/local/go/src/net/netip/netip.go:625
	}
//line /usr/local/go/src/net/netip/netip.go:625
	// _ = "end of CoverTab[3798]"
//line /usr/local/go/src/net/netip/netip.go:625
	_go_fuzz_dep_.CoverTab[3799]++

							return false
//line /usr/local/go/src/net/netip/netip.go:627
	// _ = "end of CoverTab[3799]"
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
	_go_fuzz_dep_.CoverTab[3808]++
							return ip == IPv4Unspecified() || func() bool {
//line /usr/local/go/src/net/netip/netip.go:635
		_go_fuzz_dep_.CoverTab[3809]++
//line /usr/local/go/src/net/netip/netip.go:635
		return ip == IPv6Unspecified()
//line /usr/local/go/src/net/netip/netip.go:635
		// _ = "end of CoverTab[3809]"
//line /usr/local/go/src/net/netip/netip.go:635
	}()
//line /usr/local/go/src/net/netip/netip.go:635
	// _ = "end of CoverTab[3808]"
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
	_go_fuzz_dep_.CoverTab[3810]++
							if b < 0 {
//line /usr/local/go/src/net/netip/netip.go:644
		_go_fuzz_dep_.CoverTab[3813]++
								return Prefix{}, errors.New("negative Prefix bits")
//line /usr/local/go/src/net/netip/netip.go:645
		// _ = "end of CoverTab[3813]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:646
		_go_fuzz_dep_.CoverTab[3814]++
//line /usr/local/go/src/net/netip/netip.go:646
		// _ = "end of CoverTab[3814]"
//line /usr/local/go/src/net/netip/netip.go:646
	}
//line /usr/local/go/src/net/netip/netip.go:646
	// _ = "end of CoverTab[3810]"
//line /usr/local/go/src/net/netip/netip.go:646
	_go_fuzz_dep_.CoverTab[3811]++
							effectiveBits := b
							switch ip.z {
	case z0:
//line /usr/local/go/src/net/netip/netip.go:649
		_go_fuzz_dep_.CoverTab[3815]++
								return Prefix{}, nil
//line /usr/local/go/src/net/netip/netip.go:650
		// _ = "end of CoverTab[3815]"
	case z4:
//line /usr/local/go/src/net/netip/netip.go:651
		_go_fuzz_dep_.CoverTab[3816]++
								if b > 32 {
//line /usr/local/go/src/net/netip/netip.go:652
			_go_fuzz_dep_.CoverTab[3819]++
									return Prefix{}, errors.New("prefix length " + itoa.Itoa(b) + " too large for IPv4")
//line /usr/local/go/src/net/netip/netip.go:653
			// _ = "end of CoverTab[3819]"
		} else {
//line /usr/local/go/src/net/netip/netip.go:654
			_go_fuzz_dep_.CoverTab[3820]++
//line /usr/local/go/src/net/netip/netip.go:654
			// _ = "end of CoverTab[3820]"
//line /usr/local/go/src/net/netip/netip.go:654
		}
//line /usr/local/go/src/net/netip/netip.go:654
		// _ = "end of CoverTab[3816]"
//line /usr/local/go/src/net/netip/netip.go:654
		_go_fuzz_dep_.CoverTab[3817]++
								effectiveBits += 96
//line /usr/local/go/src/net/netip/netip.go:655
		// _ = "end of CoverTab[3817]"
	default:
//line /usr/local/go/src/net/netip/netip.go:656
		_go_fuzz_dep_.CoverTab[3818]++
								if b > 128 {
//line /usr/local/go/src/net/netip/netip.go:657
			_go_fuzz_dep_.CoverTab[3821]++
									return Prefix{}, errors.New("prefix length " + itoa.Itoa(b) + " too large for IPv6")
//line /usr/local/go/src/net/netip/netip.go:658
			// _ = "end of CoverTab[3821]"
		} else {
//line /usr/local/go/src/net/netip/netip.go:659
			_go_fuzz_dep_.CoverTab[3822]++
//line /usr/local/go/src/net/netip/netip.go:659
			// _ = "end of CoverTab[3822]"
//line /usr/local/go/src/net/netip/netip.go:659
		}
//line /usr/local/go/src/net/netip/netip.go:659
		// _ = "end of CoverTab[3818]"
	}
//line /usr/local/go/src/net/netip/netip.go:660
	// _ = "end of CoverTab[3811]"
//line /usr/local/go/src/net/netip/netip.go:660
	_go_fuzz_dep_.CoverTab[3812]++
							ip.addr = ip.addr.and(mask6(effectiveBits))
							return PrefixFrom(ip, b), nil
//line /usr/local/go/src/net/netip/netip.go:662
	// _ = "end of CoverTab[3812]"
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
	_go_fuzz_dep_.CoverTab[3823]++
							bePutUint64(a16[:8], ip.addr.hi)
							bePutUint64(a16[8:], ip.addr.lo)
							return a16
//line /usr/local/go/src/net/netip/netip.go:678
	// _ = "end of CoverTab[3823]"
}

// As4 returns an IPv4 or IPv4-in-IPv6 address in its 4-byte representation.
//line /usr/local/go/src/net/netip/netip.go:681
// If ip is the zero Addr or an IPv6 address, As4 panics.
//line /usr/local/go/src/net/netip/netip.go:681
// Note that 0.0.0.0 is not the zero Addr.
//line /usr/local/go/src/net/netip/netip.go:684
func (ip Addr) As4() (a4 [4]byte) {
//line /usr/local/go/src/net/netip/netip.go:684
	_go_fuzz_dep_.CoverTab[3824]++
							if ip.z == z4 || func() bool {
//line /usr/local/go/src/net/netip/netip.go:685
		_go_fuzz_dep_.CoverTab[3827]++
//line /usr/local/go/src/net/netip/netip.go:685
		return ip.Is4In6()
//line /usr/local/go/src/net/netip/netip.go:685
		// _ = "end of CoverTab[3827]"
//line /usr/local/go/src/net/netip/netip.go:685
	}() {
//line /usr/local/go/src/net/netip/netip.go:685
		_go_fuzz_dep_.CoverTab[3828]++
								bePutUint32(a4[:], uint32(ip.addr.lo))
								return a4
//line /usr/local/go/src/net/netip/netip.go:687
		// _ = "end of CoverTab[3828]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:688
		_go_fuzz_dep_.CoverTab[3829]++
//line /usr/local/go/src/net/netip/netip.go:688
		// _ = "end of CoverTab[3829]"
//line /usr/local/go/src/net/netip/netip.go:688
	}
//line /usr/local/go/src/net/netip/netip.go:688
	// _ = "end of CoverTab[3824]"
//line /usr/local/go/src/net/netip/netip.go:688
	_go_fuzz_dep_.CoverTab[3825]++
							if ip.z == z0 {
//line /usr/local/go/src/net/netip/netip.go:689
		_go_fuzz_dep_.CoverTab[3830]++
								panic("As4 called on IP zero value")
//line /usr/local/go/src/net/netip/netip.go:690
		// _ = "end of CoverTab[3830]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:691
		_go_fuzz_dep_.CoverTab[3831]++
//line /usr/local/go/src/net/netip/netip.go:691
		// _ = "end of CoverTab[3831]"
//line /usr/local/go/src/net/netip/netip.go:691
	}
//line /usr/local/go/src/net/netip/netip.go:691
	// _ = "end of CoverTab[3825]"
//line /usr/local/go/src/net/netip/netip.go:691
	_go_fuzz_dep_.CoverTab[3826]++
							panic("As4 called on IPv6 address")
//line /usr/local/go/src/net/netip/netip.go:692
	// _ = "end of CoverTab[3826]"
}

// AsSlice returns an IPv4 or IPv6 address in its respective 4-byte or 16-byte representation.
func (ip Addr) AsSlice() []byte {
//line /usr/local/go/src/net/netip/netip.go:696
	_go_fuzz_dep_.CoverTab[3832]++
							switch ip.z {
	case z0:
//line /usr/local/go/src/net/netip/netip.go:698
		_go_fuzz_dep_.CoverTab[3833]++
								return nil
//line /usr/local/go/src/net/netip/netip.go:699
		// _ = "end of CoverTab[3833]"
	case z4:
//line /usr/local/go/src/net/netip/netip.go:700
		_go_fuzz_dep_.CoverTab[3834]++
								var ret [4]byte
								bePutUint32(ret[:], uint32(ip.addr.lo))
								return ret[:]
//line /usr/local/go/src/net/netip/netip.go:703
		// _ = "end of CoverTab[3834]"
	default:
//line /usr/local/go/src/net/netip/netip.go:704
		_go_fuzz_dep_.CoverTab[3835]++
								var ret [16]byte
								bePutUint64(ret[:8], ip.addr.hi)
								bePutUint64(ret[8:], ip.addr.lo)
								return ret[:]
//line /usr/local/go/src/net/netip/netip.go:708
		// _ = "end of CoverTab[3835]"
	}
//line /usr/local/go/src/net/netip/netip.go:709
	// _ = "end of CoverTab[3832]"
}

// Next returns the address following ip.
//line /usr/local/go/src/net/netip/netip.go:712
// If there is none, it returns the zero Addr.
//line /usr/local/go/src/net/netip/netip.go:714
func (ip Addr) Next() Addr {
//line /usr/local/go/src/net/netip/netip.go:714
	_go_fuzz_dep_.CoverTab[3836]++
							ip.addr = ip.addr.addOne()
							if ip.Is4() {
//line /usr/local/go/src/net/netip/netip.go:716
		_go_fuzz_dep_.CoverTab[3838]++
								if uint32(ip.addr.lo) == 0 {
//line /usr/local/go/src/net/netip/netip.go:717
			_go_fuzz_dep_.CoverTab[3839]++

									return Addr{}
//line /usr/local/go/src/net/netip/netip.go:719
			// _ = "end of CoverTab[3839]"
		} else {
//line /usr/local/go/src/net/netip/netip.go:720
			_go_fuzz_dep_.CoverTab[3840]++
//line /usr/local/go/src/net/netip/netip.go:720
			// _ = "end of CoverTab[3840]"
//line /usr/local/go/src/net/netip/netip.go:720
		}
//line /usr/local/go/src/net/netip/netip.go:720
		// _ = "end of CoverTab[3838]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:721
		_go_fuzz_dep_.CoverTab[3841]++
								if ip.addr.isZero() {
//line /usr/local/go/src/net/netip/netip.go:722
			_go_fuzz_dep_.CoverTab[3842]++

									return Addr{}
//line /usr/local/go/src/net/netip/netip.go:724
			// _ = "end of CoverTab[3842]"
		} else {
//line /usr/local/go/src/net/netip/netip.go:725
			_go_fuzz_dep_.CoverTab[3843]++
//line /usr/local/go/src/net/netip/netip.go:725
			// _ = "end of CoverTab[3843]"
//line /usr/local/go/src/net/netip/netip.go:725
		}
//line /usr/local/go/src/net/netip/netip.go:725
		// _ = "end of CoverTab[3841]"
	}
//line /usr/local/go/src/net/netip/netip.go:726
	// _ = "end of CoverTab[3836]"
//line /usr/local/go/src/net/netip/netip.go:726
	_go_fuzz_dep_.CoverTab[3837]++
							return ip
//line /usr/local/go/src/net/netip/netip.go:727
	// _ = "end of CoverTab[3837]"
}

// Prev returns the IP before ip.
//line /usr/local/go/src/net/netip/netip.go:730
// If there is none, it returns the IP zero value.
//line /usr/local/go/src/net/netip/netip.go:732
func (ip Addr) Prev() Addr {
//line /usr/local/go/src/net/netip/netip.go:732
	_go_fuzz_dep_.CoverTab[3844]++
							if ip.Is4() {
//line /usr/local/go/src/net/netip/netip.go:733
		_go_fuzz_dep_.CoverTab[3846]++
								if uint32(ip.addr.lo) == 0 {
//line /usr/local/go/src/net/netip/netip.go:734
			_go_fuzz_dep_.CoverTab[3847]++
									return Addr{}
//line /usr/local/go/src/net/netip/netip.go:735
			// _ = "end of CoverTab[3847]"
		} else {
//line /usr/local/go/src/net/netip/netip.go:736
			_go_fuzz_dep_.CoverTab[3848]++
//line /usr/local/go/src/net/netip/netip.go:736
			// _ = "end of CoverTab[3848]"
//line /usr/local/go/src/net/netip/netip.go:736
		}
//line /usr/local/go/src/net/netip/netip.go:736
		// _ = "end of CoverTab[3846]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:737
		_go_fuzz_dep_.CoverTab[3849]++
//line /usr/local/go/src/net/netip/netip.go:737
		if ip.addr.isZero() {
//line /usr/local/go/src/net/netip/netip.go:737
			_go_fuzz_dep_.CoverTab[3850]++
									return Addr{}
//line /usr/local/go/src/net/netip/netip.go:738
			// _ = "end of CoverTab[3850]"
		} else {
//line /usr/local/go/src/net/netip/netip.go:739
			_go_fuzz_dep_.CoverTab[3851]++
//line /usr/local/go/src/net/netip/netip.go:739
			// _ = "end of CoverTab[3851]"
//line /usr/local/go/src/net/netip/netip.go:739
		}
//line /usr/local/go/src/net/netip/netip.go:739
		// _ = "end of CoverTab[3849]"
//line /usr/local/go/src/net/netip/netip.go:739
	}
//line /usr/local/go/src/net/netip/netip.go:739
	// _ = "end of CoverTab[3844]"
//line /usr/local/go/src/net/netip/netip.go:739
	_go_fuzz_dep_.CoverTab[3845]++
							ip.addr = ip.addr.subOne()
							return ip
//line /usr/local/go/src/net/netip/netip.go:741
	// _ = "end of CoverTab[3845]"
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
	_go_fuzz_dep_.CoverTab[3852]++
							switch ip.z {
	case z0:
//line /usr/local/go/src/net/netip/netip.go:758
		_go_fuzz_dep_.CoverTab[3853]++
								return "invalid IP"
//line /usr/local/go/src/net/netip/netip.go:759
		// _ = "end of CoverTab[3853]"
	case z4:
//line /usr/local/go/src/net/netip/netip.go:760
		_go_fuzz_dep_.CoverTab[3854]++
								return ip.string4()
//line /usr/local/go/src/net/netip/netip.go:761
		// _ = "end of CoverTab[3854]"
	default:
//line /usr/local/go/src/net/netip/netip.go:762
		_go_fuzz_dep_.CoverTab[3855]++
								if ip.Is4In6() {
//line /usr/local/go/src/net/netip/netip.go:763
			_go_fuzz_dep_.CoverTab[3857]++
									if z := ip.Zone(); z != "" {
//line /usr/local/go/src/net/netip/netip.go:764
				_go_fuzz_dep_.CoverTab[3858]++
										return "::ffff:" + ip.Unmap().string4() + "%" + z
//line /usr/local/go/src/net/netip/netip.go:765
				// _ = "end of CoverTab[3858]"
			} else {
//line /usr/local/go/src/net/netip/netip.go:766
				_go_fuzz_dep_.CoverTab[3859]++
										return "::ffff:" + ip.Unmap().string4()
//line /usr/local/go/src/net/netip/netip.go:767
				// _ = "end of CoverTab[3859]"
			}
//line /usr/local/go/src/net/netip/netip.go:768
			// _ = "end of CoverTab[3857]"
		} else {
//line /usr/local/go/src/net/netip/netip.go:769
			_go_fuzz_dep_.CoverTab[3860]++
//line /usr/local/go/src/net/netip/netip.go:769
			// _ = "end of CoverTab[3860]"
//line /usr/local/go/src/net/netip/netip.go:769
		}
//line /usr/local/go/src/net/netip/netip.go:769
		// _ = "end of CoverTab[3855]"
//line /usr/local/go/src/net/netip/netip.go:769
		_go_fuzz_dep_.CoverTab[3856]++
								return ip.string6()
//line /usr/local/go/src/net/netip/netip.go:770
		// _ = "end of CoverTab[3856]"
	}
//line /usr/local/go/src/net/netip/netip.go:771
	// _ = "end of CoverTab[3852]"
}

// AppendTo appends a text encoding of ip,
//line /usr/local/go/src/net/netip/netip.go:774
// as generated by MarshalText,
//line /usr/local/go/src/net/netip/netip.go:774
// to b and returns the extended buffer.
//line /usr/local/go/src/net/netip/netip.go:777
func (ip Addr) AppendTo(b []byte) []byte {
//line /usr/local/go/src/net/netip/netip.go:777
	_go_fuzz_dep_.CoverTab[3861]++
							switch ip.z {
	case z0:
//line /usr/local/go/src/net/netip/netip.go:779
		_go_fuzz_dep_.CoverTab[3862]++
								return b
//line /usr/local/go/src/net/netip/netip.go:780
		// _ = "end of CoverTab[3862]"
	case z4:
//line /usr/local/go/src/net/netip/netip.go:781
		_go_fuzz_dep_.CoverTab[3863]++
								return ip.appendTo4(b)
//line /usr/local/go/src/net/netip/netip.go:782
		// _ = "end of CoverTab[3863]"
	default:
//line /usr/local/go/src/net/netip/netip.go:783
		_go_fuzz_dep_.CoverTab[3864]++
								if ip.Is4In6() {
//line /usr/local/go/src/net/netip/netip.go:784
			_go_fuzz_dep_.CoverTab[3866]++
									b = append(b, "::ffff:"...)
									b = ip.Unmap().appendTo4(b)
									if z := ip.Zone(); z != "" {
//line /usr/local/go/src/net/netip/netip.go:787
				_go_fuzz_dep_.CoverTab[3868]++
										b = append(b, '%')
										b = append(b, z...)
//line /usr/local/go/src/net/netip/netip.go:789
				// _ = "end of CoverTab[3868]"
			} else {
//line /usr/local/go/src/net/netip/netip.go:790
				_go_fuzz_dep_.CoverTab[3869]++
//line /usr/local/go/src/net/netip/netip.go:790
				// _ = "end of CoverTab[3869]"
//line /usr/local/go/src/net/netip/netip.go:790
			}
//line /usr/local/go/src/net/netip/netip.go:790
			// _ = "end of CoverTab[3866]"
//line /usr/local/go/src/net/netip/netip.go:790
			_go_fuzz_dep_.CoverTab[3867]++
									return b
//line /usr/local/go/src/net/netip/netip.go:791
			// _ = "end of CoverTab[3867]"
		} else {
//line /usr/local/go/src/net/netip/netip.go:792
			_go_fuzz_dep_.CoverTab[3870]++
//line /usr/local/go/src/net/netip/netip.go:792
			// _ = "end of CoverTab[3870]"
//line /usr/local/go/src/net/netip/netip.go:792
		}
//line /usr/local/go/src/net/netip/netip.go:792
		// _ = "end of CoverTab[3864]"
//line /usr/local/go/src/net/netip/netip.go:792
		_go_fuzz_dep_.CoverTab[3865]++
								return ip.appendTo6(b)
//line /usr/local/go/src/net/netip/netip.go:793
		// _ = "end of CoverTab[3865]"
	}
//line /usr/local/go/src/net/netip/netip.go:794
	// _ = "end of CoverTab[3861]"
}

// digits is a string of the hex digits from 0 to f. It's used in
//line /usr/local/go/src/net/netip/netip.go:797
// appendDecimal and appendHex to format IP addresses.
//line /usr/local/go/src/net/netip/netip.go:799
const digits = "0123456789abcdef"

// appendDecimal appends the decimal string representation of x to b.
func appendDecimal(b []byte, x uint8) []byte {
//line /usr/local/go/src/net/netip/netip.go:802
	_go_fuzz_dep_.CoverTab[3871]++

//line /usr/local/go/src/net/netip/netip.go:806
	if x >= 100 {
//line /usr/local/go/src/net/netip/netip.go:806
		_go_fuzz_dep_.CoverTab[3874]++
								b = append(b, digits[x/100])
//line /usr/local/go/src/net/netip/netip.go:807
		// _ = "end of CoverTab[3874]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:808
		_go_fuzz_dep_.CoverTab[3875]++
//line /usr/local/go/src/net/netip/netip.go:808
		// _ = "end of CoverTab[3875]"
//line /usr/local/go/src/net/netip/netip.go:808
	}
//line /usr/local/go/src/net/netip/netip.go:808
	// _ = "end of CoverTab[3871]"
//line /usr/local/go/src/net/netip/netip.go:808
	_go_fuzz_dep_.CoverTab[3872]++
							if x >= 10 {
//line /usr/local/go/src/net/netip/netip.go:809
		_go_fuzz_dep_.CoverTab[3876]++
								b = append(b, digits[x/10%10])
//line /usr/local/go/src/net/netip/netip.go:810
		// _ = "end of CoverTab[3876]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:811
		_go_fuzz_dep_.CoverTab[3877]++
//line /usr/local/go/src/net/netip/netip.go:811
		// _ = "end of CoverTab[3877]"
//line /usr/local/go/src/net/netip/netip.go:811
	}
//line /usr/local/go/src/net/netip/netip.go:811
	// _ = "end of CoverTab[3872]"
//line /usr/local/go/src/net/netip/netip.go:811
	_go_fuzz_dep_.CoverTab[3873]++
							return append(b, digits[x%10])
//line /usr/local/go/src/net/netip/netip.go:812
	// _ = "end of CoverTab[3873]"
}

// appendHex appends the hex string representation of x to b.
func appendHex(b []byte, x uint16) []byte {
//line /usr/local/go/src/net/netip/netip.go:816
	_go_fuzz_dep_.CoverTab[3878]++

//line /usr/local/go/src/net/netip/netip.go:820
	if x >= 0x1000 {
//line /usr/local/go/src/net/netip/netip.go:820
		_go_fuzz_dep_.CoverTab[3882]++
								b = append(b, digits[x>>12])
//line /usr/local/go/src/net/netip/netip.go:821
		// _ = "end of CoverTab[3882]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:822
		_go_fuzz_dep_.CoverTab[3883]++
//line /usr/local/go/src/net/netip/netip.go:822
		// _ = "end of CoverTab[3883]"
//line /usr/local/go/src/net/netip/netip.go:822
	}
//line /usr/local/go/src/net/netip/netip.go:822
	// _ = "end of CoverTab[3878]"
//line /usr/local/go/src/net/netip/netip.go:822
	_go_fuzz_dep_.CoverTab[3879]++
							if x >= 0x100 {
//line /usr/local/go/src/net/netip/netip.go:823
		_go_fuzz_dep_.CoverTab[3884]++
								b = append(b, digits[x>>8&0xf])
//line /usr/local/go/src/net/netip/netip.go:824
		// _ = "end of CoverTab[3884]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:825
		_go_fuzz_dep_.CoverTab[3885]++
//line /usr/local/go/src/net/netip/netip.go:825
		// _ = "end of CoverTab[3885]"
//line /usr/local/go/src/net/netip/netip.go:825
	}
//line /usr/local/go/src/net/netip/netip.go:825
	// _ = "end of CoverTab[3879]"
//line /usr/local/go/src/net/netip/netip.go:825
	_go_fuzz_dep_.CoverTab[3880]++
							if x >= 0x10 {
//line /usr/local/go/src/net/netip/netip.go:826
		_go_fuzz_dep_.CoverTab[3886]++
								b = append(b, digits[x>>4&0xf])
//line /usr/local/go/src/net/netip/netip.go:827
		// _ = "end of CoverTab[3886]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:828
		_go_fuzz_dep_.CoverTab[3887]++
//line /usr/local/go/src/net/netip/netip.go:828
		// _ = "end of CoverTab[3887]"
//line /usr/local/go/src/net/netip/netip.go:828
	}
//line /usr/local/go/src/net/netip/netip.go:828
	// _ = "end of CoverTab[3880]"
//line /usr/local/go/src/net/netip/netip.go:828
	_go_fuzz_dep_.CoverTab[3881]++
							return append(b, digits[x&0xf])
//line /usr/local/go/src/net/netip/netip.go:829
	// _ = "end of CoverTab[3881]"
}

// appendHexPad appends the fully padded hex string representation of x to b.
func appendHexPad(b []byte, x uint16) []byte {
//line /usr/local/go/src/net/netip/netip.go:833
	_go_fuzz_dep_.CoverTab[3888]++
							return append(b, digits[x>>12], digits[x>>8&0xf], digits[x>>4&0xf], digits[x&0xf])
//line /usr/local/go/src/net/netip/netip.go:834
	// _ = "end of CoverTab[3888]"
}

func (ip Addr) string4() string {
//line /usr/local/go/src/net/netip/netip.go:837
	_go_fuzz_dep_.CoverTab[3889]++
							const max = len("255.255.255.255")
							ret := make([]byte, 0, max)
							ret = ip.appendTo4(ret)
							return string(ret)
//line /usr/local/go/src/net/netip/netip.go:841
	// _ = "end of CoverTab[3889]"
}

func (ip Addr) appendTo4(ret []byte) []byte {
//line /usr/local/go/src/net/netip/netip.go:844
	_go_fuzz_dep_.CoverTab[3890]++
							ret = appendDecimal(ret, ip.v4(0))
							ret = append(ret, '.')
							ret = appendDecimal(ret, ip.v4(1))
							ret = append(ret, '.')
							ret = appendDecimal(ret, ip.v4(2))
							ret = append(ret, '.')
							ret = appendDecimal(ret, ip.v4(3))
							return ret
//line /usr/local/go/src/net/netip/netip.go:852
	// _ = "end of CoverTab[3890]"
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
	_go_fuzz_dep_.CoverTab[3891]++
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
	// _ = "end of CoverTab[3891]"
}

func (ip Addr) appendTo6(ret []byte) []byte {
//line /usr/local/go/src/net/netip/netip.go:874
	_go_fuzz_dep_.CoverTab[3892]++
							zeroStart, zeroEnd := uint8(255), uint8(255)
							for i := uint8(0); i < 8; i++ {
//line /usr/local/go/src/net/netip/netip.go:876
		_go_fuzz_dep_.CoverTab[3896]++
								j := i
								for j < 8 && func() bool {
//line /usr/local/go/src/net/netip/netip.go:878
			_go_fuzz_dep_.CoverTab[3898]++
//line /usr/local/go/src/net/netip/netip.go:878
			return ip.v6u16(j) == 0
//line /usr/local/go/src/net/netip/netip.go:878
			// _ = "end of CoverTab[3898]"
//line /usr/local/go/src/net/netip/netip.go:878
		}() {
//line /usr/local/go/src/net/netip/netip.go:878
			_go_fuzz_dep_.CoverTab[3899]++
									j++
//line /usr/local/go/src/net/netip/netip.go:879
			// _ = "end of CoverTab[3899]"
		}
//line /usr/local/go/src/net/netip/netip.go:880
		// _ = "end of CoverTab[3896]"
//line /usr/local/go/src/net/netip/netip.go:880
		_go_fuzz_dep_.CoverTab[3897]++
								if l := j - i; l >= 2 && func() bool {
//line /usr/local/go/src/net/netip/netip.go:881
			_go_fuzz_dep_.CoverTab[3900]++
//line /usr/local/go/src/net/netip/netip.go:881
			return l > zeroEnd-zeroStart
//line /usr/local/go/src/net/netip/netip.go:881
			// _ = "end of CoverTab[3900]"
//line /usr/local/go/src/net/netip/netip.go:881
		}() {
//line /usr/local/go/src/net/netip/netip.go:881
			_go_fuzz_dep_.CoverTab[3901]++
									zeroStart, zeroEnd = i, j
//line /usr/local/go/src/net/netip/netip.go:882
			// _ = "end of CoverTab[3901]"
		} else {
//line /usr/local/go/src/net/netip/netip.go:883
			_go_fuzz_dep_.CoverTab[3902]++
//line /usr/local/go/src/net/netip/netip.go:883
			// _ = "end of CoverTab[3902]"
//line /usr/local/go/src/net/netip/netip.go:883
		}
//line /usr/local/go/src/net/netip/netip.go:883
		// _ = "end of CoverTab[3897]"
	}
//line /usr/local/go/src/net/netip/netip.go:884
	// _ = "end of CoverTab[3892]"
//line /usr/local/go/src/net/netip/netip.go:884
	_go_fuzz_dep_.CoverTab[3893]++

							for i := uint8(0); i < 8; i++ {
//line /usr/local/go/src/net/netip/netip.go:886
		_go_fuzz_dep_.CoverTab[3903]++
								if i == zeroStart {
//line /usr/local/go/src/net/netip/netip.go:887
			_go_fuzz_dep_.CoverTab[3905]++
									ret = append(ret, ':', ':')
									i = zeroEnd
									if i >= 8 {
//line /usr/local/go/src/net/netip/netip.go:890
				_go_fuzz_dep_.CoverTab[3906]++
										break
//line /usr/local/go/src/net/netip/netip.go:891
				// _ = "end of CoverTab[3906]"
			} else {
//line /usr/local/go/src/net/netip/netip.go:892
				_go_fuzz_dep_.CoverTab[3907]++
//line /usr/local/go/src/net/netip/netip.go:892
				// _ = "end of CoverTab[3907]"
//line /usr/local/go/src/net/netip/netip.go:892
			}
//line /usr/local/go/src/net/netip/netip.go:892
			// _ = "end of CoverTab[3905]"
		} else {
//line /usr/local/go/src/net/netip/netip.go:893
			_go_fuzz_dep_.CoverTab[3908]++
//line /usr/local/go/src/net/netip/netip.go:893
			if i > 0 {
//line /usr/local/go/src/net/netip/netip.go:893
				_go_fuzz_dep_.CoverTab[3909]++
										ret = append(ret, ':')
//line /usr/local/go/src/net/netip/netip.go:894
				// _ = "end of CoverTab[3909]"
			} else {
//line /usr/local/go/src/net/netip/netip.go:895
				_go_fuzz_dep_.CoverTab[3910]++
//line /usr/local/go/src/net/netip/netip.go:895
				// _ = "end of CoverTab[3910]"
//line /usr/local/go/src/net/netip/netip.go:895
			}
//line /usr/local/go/src/net/netip/netip.go:895
			// _ = "end of CoverTab[3908]"
//line /usr/local/go/src/net/netip/netip.go:895
		}
//line /usr/local/go/src/net/netip/netip.go:895
		// _ = "end of CoverTab[3903]"
//line /usr/local/go/src/net/netip/netip.go:895
		_go_fuzz_dep_.CoverTab[3904]++

								ret = appendHex(ret, ip.v6u16(i))
//line /usr/local/go/src/net/netip/netip.go:897
		// _ = "end of CoverTab[3904]"
	}
//line /usr/local/go/src/net/netip/netip.go:898
	// _ = "end of CoverTab[3893]"
//line /usr/local/go/src/net/netip/netip.go:898
	_go_fuzz_dep_.CoverTab[3894]++

							if ip.z != z6noz {
//line /usr/local/go/src/net/netip/netip.go:900
		_go_fuzz_dep_.CoverTab[3911]++
								ret = append(ret, '%')
								ret = append(ret, ip.Zone()...)
//line /usr/local/go/src/net/netip/netip.go:902
		// _ = "end of CoverTab[3911]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:903
		_go_fuzz_dep_.CoverTab[3912]++
//line /usr/local/go/src/net/netip/netip.go:903
		// _ = "end of CoverTab[3912]"
//line /usr/local/go/src/net/netip/netip.go:903
	}
//line /usr/local/go/src/net/netip/netip.go:903
	// _ = "end of CoverTab[3894]"
//line /usr/local/go/src/net/netip/netip.go:903
	_go_fuzz_dep_.CoverTab[3895]++
							return ret
//line /usr/local/go/src/net/netip/netip.go:904
	// _ = "end of CoverTab[3895]"
}

// StringExpanded is like String but IPv6 addresses are expanded with leading
//line /usr/local/go/src/net/netip/netip.go:907
// zeroes and no "::" compression. For example, "2001:db8::1" becomes
//line /usr/local/go/src/net/netip/netip.go:907
// "2001:0db8:0000:0000:0000:0000:0000:0001".
//line /usr/local/go/src/net/netip/netip.go:910
func (ip Addr) StringExpanded() string {
//line /usr/local/go/src/net/netip/netip.go:910
	_go_fuzz_dep_.CoverTab[3913]++
							switch ip.z {
	case z0, z4:
//line /usr/local/go/src/net/netip/netip.go:912
		_go_fuzz_dep_.CoverTab[3917]++
								return ip.String()
//line /usr/local/go/src/net/netip/netip.go:913
		// _ = "end of CoverTab[3917]"
//line /usr/local/go/src/net/netip/netip.go:913
	default:
//line /usr/local/go/src/net/netip/netip.go:913
		_go_fuzz_dep_.CoverTab[3918]++
//line /usr/local/go/src/net/netip/netip.go:913
		// _ = "end of CoverTab[3918]"
	}
//line /usr/local/go/src/net/netip/netip.go:914
	// _ = "end of CoverTab[3913]"
//line /usr/local/go/src/net/netip/netip.go:914
	_go_fuzz_dep_.CoverTab[3914]++

							const size = len("ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff")
							ret := make([]byte, 0, size)
							for i := uint8(0); i < 8; i++ {
//line /usr/local/go/src/net/netip/netip.go:918
		_go_fuzz_dep_.CoverTab[3919]++
								if i > 0 {
//line /usr/local/go/src/net/netip/netip.go:919
			_go_fuzz_dep_.CoverTab[3921]++
									ret = append(ret, ':')
//line /usr/local/go/src/net/netip/netip.go:920
			// _ = "end of CoverTab[3921]"
		} else {
//line /usr/local/go/src/net/netip/netip.go:921
			_go_fuzz_dep_.CoverTab[3922]++
//line /usr/local/go/src/net/netip/netip.go:921
			// _ = "end of CoverTab[3922]"
//line /usr/local/go/src/net/netip/netip.go:921
		}
//line /usr/local/go/src/net/netip/netip.go:921
		// _ = "end of CoverTab[3919]"
//line /usr/local/go/src/net/netip/netip.go:921
		_go_fuzz_dep_.CoverTab[3920]++

								ret = appendHexPad(ret, ip.v6u16(i))
//line /usr/local/go/src/net/netip/netip.go:923
		// _ = "end of CoverTab[3920]"
	}
//line /usr/local/go/src/net/netip/netip.go:924
	// _ = "end of CoverTab[3914]"
//line /usr/local/go/src/net/netip/netip.go:924
	_go_fuzz_dep_.CoverTab[3915]++

							if ip.z != z6noz {
//line /usr/local/go/src/net/netip/netip.go:926
		_go_fuzz_dep_.CoverTab[3923]++

//line /usr/local/go/src/net/netip/netip.go:929
		ret = append(ret, '%')
								ret = append(ret, ip.Zone()...)
//line /usr/local/go/src/net/netip/netip.go:930
		// _ = "end of CoverTab[3923]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:931
		_go_fuzz_dep_.CoverTab[3924]++
//line /usr/local/go/src/net/netip/netip.go:931
		// _ = "end of CoverTab[3924]"
//line /usr/local/go/src/net/netip/netip.go:931
	}
//line /usr/local/go/src/net/netip/netip.go:931
	// _ = "end of CoverTab[3915]"
//line /usr/local/go/src/net/netip/netip.go:931
	_go_fuzz_dep_.CoverTab[3916]++
							return string(ret)
//line /usr/local/go/src/net/netip/netip.go:932
	// _ = "end of CoverTab[3916]"
}

// MarshalText implements the encoding.TextMarshaler interface,
//line /usr/local/go/src/net/netip/netip.go:935
// The encoding is the same as returned by String, with one exception:
//line /usr/local/go/src/net/netip/netip.go:935
// If ip is the zero Addr, the encoding is the empty string.
//line /usr/local/go/src/net/netip/netip.go:938
func (ip Addr) MarshalText() ([]byte, error) {
//line /usr/local/go/src/net/netip/netip.go:938
	_go_fuzz_dep_.CoverTab[3925]++
							switch ip.z {
	case z0:
//line /usr/local/go/src/net/netip/netip.go:940
		_go_fuzz_dep_.CoverTab[3926]++
								return []byte(""), nil
//line /usr/local/go/src/net/netip/netip.go:941
		// _ = "end of CoverTab[3926]"
	case z4:
//line /usr/local/go/src/net/netip/netip.go:942
		_go_fuzz_dep_.CoverTab[3927]++
								max := len("255.255.255.255")
								b := make([]byte, 0, max)
								return ip.appendTo4(b), nil
//line /usr/local/go/src/net/netip/netip.go:945
		// _ = "end of CoverTab[3927]"
	default:
//line /usr/local/go/src/net/netip/netip.go:946
		_go_fuzz_dep_.CoverTab[3928]++
								max := len("ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff%enp5s0")
								b := make([]byte, 0, max)
								if ip.Is4In6() {
//line /usr/local/go/src/net/netip/netip.go:949
			_go_fuzz_dep_.CoverTab[3930]++
									b = append(b, "::ffff:"...)
									b = ip.Unmap().appendTo4(b)
									if z := ip.Zone(); z != "" {
//line /usr/local/go/src/net/netip/netip.go:952
				_go_fuzz_dep_.CoverTab[3932]++
										b = append(b, '%')
										b = append(b, z...)
//line /usr/local/go/src/net/netip/netip.go:954
				// _ = "end of CoverTab[3932]"
			} else {
//line /usr/local/go/src/net/netip/netip.go:955
				_go_fuzz_dep_.CoverTab[3933]++
//line /usr/local/go/src/net/netip/netip.go:955
				// _ = "end of CoverTab[3933]"
//line /usr/local/go/src/net/netip/netip.go:955
			}
//line /usr/local/go/src/net/netip/netip.go:955
			// _ = "end of CoverTab[3930]"
//line /usr/local/go/src/net/netip/netip.go:955
			_go_fuzz_dep_.CoverTab[3931]++
									return b, nil
//line /usr/local/go/src/net/netip/netip.go:956
			// _ = "end of CoverTab[3931]"
		} else {
//line /usr/local/go/src/net/netip/netip.go:957
			_go_fuzz_dep_.CoverTab[3934]++
//line /usr/local/go/src/net/netip/netip.go:957
			// _ = "end of CoverTab[3934]"
//line /usr/local/go/src/net/netip/netip.go:957
		}
//line /usr/local/go/src/net/netip/netip.go:957
		// _ = "end of CoverTab[3928]"
//line /usr/local/go/src/net/netip/netip.go:957
		_go_fuzz_dep_.CoverTab[3929]++
								return ip.appendTo6(b), nil
//line /usr/local/go/src/net/netip/netip.go:958
		// _ = "end of CoverTab[3929]"
	}
//line /usr/local/go/src/net/netip/netip.go:959
	// _ = "end of CoverTab[3925]"

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
	_go_fuzz_dep_.CoverTab[3935]++
							if len(text) == 0 {
//line /usr/local/go/src/net/netip/netip.go:969
		_go_fuzz_dep_.CoverTab[3937]++
								*ip = Addr{}
								return nil
//line /usr/local/go/src/net/netip/netip.go:971
		// _ = "end of CoverTab[3937]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:972
		_go_fuzz_dep_.CoverTab[3938]++
//line /usr/local/go/src/net/netip/netip.go:972
		// _ = "end of CoverTab[3938]"
//line /usr/local/go/src/net/netip/netip.go:972
	}
//line /usr/local/go/src/net/netip/netip.go:972
	// _ = "end of CoverTab[3935]"
//line /usr/local/go/src/net/netip/netip.go:972
	_go_fuzz_dep_.CoverTab[3936]++
							var err error
							*ip, err = ParseAddr(string(text))
							return err
//line /usr/local/go/src/net/netip/netip.go:975
	// _ = "end of CoverTab[3936]"
}

func (ip Addr) marshalBinaryWithTrailingBytes(trailingBytes int) []byte {
//line /usr/local/go/src/net/netip/netip.go:978
	_go_fuzz_dep_.CoverTab[3939]++
							var b []byte
							switch ip.z {
	case z0:
//line /usr/local/go/src/net/netip/netip.go:981
		_go_fuzz_dep_.CoverTab[3941]++
								b = make([]byte, trailingBytes)
//line /usr/local/go/src/net/netip/netip.go:982
		// _ = "end of CoverTab[3941]"
	case z4:
//line /usr/local/go/src/net/netip/netip.go:983
		_go_fuzz_dep_.CoverTab[3942]++
								b = make([]byte, 4+trailingBytes)
								bePutUint32(b, uint32(ip.addr.lo))
//line /usr/local/go/src/net/netip/netip.go:985
		// _ = "end of CoverTab[3942]"
	default:
//line /usr/local/go/src/net/netip/netip.go:986
		_go_fuzz_dep_.CoverTab[3943]++
								z := ip.Zone()
								b = make([]byte, 16+len(z)+trailingBytes)
								bePutUint64(b[:8], ip.addr.hi)
								bePutUint64(b[8:], ip.addr.lo)
								copy(b[16:], z)
//line /usr/local/go/src/net/netip/netip.go:991
		// _ = "end of CoverTab[3943]"
	}
//line /usr/local/go/src/net/netip/netip.go:992
	// _ = "end of CoverTab[3939]"
//line /usr/local/go/src/net/netip/netip.go:992
	_go_fuzz_dep_.CoverTab[3940]++
							return b
//line /usr/local/go/src/net/netip/netip.go:993
	// _ = "end of CoverTab[3940]"
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
	_go_fuzz_dep_.CoverTab[3944]++
							return ip.marshalBinaryWithTrailingBytes(0), nil
//line /usr/local/go/src/net/netip/netip.go:1001
	// _ = "end of CoverTab[3944]"
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface.
//line /usr/local/go/src/net/netip/netip.go:1004
// It expects data in the form generated by MarshalBinary.
//line /usr/local/go/src/net/netip/netip.go:1006
func (ip *Addr) UnmarshalBinary(b []byte) error {
//line /usr/local/go/src/net/netip/netip.go:1006
	_go_fuzz_dep_.CoverTab[3945]++
							n := len(b)
							switch {
	case n == 0:
//line /usr/local/go/src/net/netip/netip.go:1009
		_go_fuzz_dep_.CoverTab[3947]++
								*ip = Addr{}
								return nil
//line /usr/local/go/src/net/netip/netip.go:1011
		// _ = "end of CoverTab[3947]"
	case n == 4:
//line /usr/local/go/src/net/netip/netip.go:1012
		_go_fuzz_dep_.CoverTab[3948]++
								*ip = AddrFrom4([4]byte(b))
								return nil
//line /usr/local/go/src/net/netip/netip.go:1014
		// _ = "end of CoverTab[3948]"
	case n == 16:
//line /usr/local/go/src/net/netip/netip.go:1015
		_go_fuzz_dep_.CoverTab[3949]++
								*ip = AddrFrom16([16]byte(b))
								return nil
//line /usr/local/go/src/net/netip/netip.go:1017
		// _ = "end of CoverTab[3949]"
	case n > 16:
//line /usr/local/go/src/net/netip/netip.go:1018
		_go_fuzz_dep_.CoverTab[3950]++
								*ip = AddrFrom16([16]byte(b[:16])).WithZone(string(b[16:]))
								return nil
//line /usr/local/go/src/net/netip/netip.go:1020
		// _ = "end of CoverTab[3950]"
//line /usr/local/go/src/net/netip/netip.go:1020
	default:
//line /usr/local/go/src/net/netip/netip.go:1020
		_go_fuzz_dep_.CoverTab[3951]++
//line /usr/local/go/src/net/netip/netip.go:1020
		// _ = "end of CoverTab[3951]"
	}
//line /usr/local/go/src/net/netip/netip.go:1021
	// _ = "end of CoverTab[3945]"
//line /usr/local/go/src/net/netip/netip.go:1021
	_go_fuzz_dep_.CoverTab[3946]++
							return errors.New("unexpected slice size")
//line /usr/local/go/src/net/netip/netip.go:1022
	// _ = "end of CoverTab[3946]"
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
	_go_fuzz_dep_.CoverTab[3952]++
//line /usr/local/go/src/net/netip/netip.go:1033
	return AddrPort{ip: ip, port: port}
//line /usr/local/go/src/net/netip/netip.go:1033
	// _ = "end of CoverTab[3952]"
//line /usr/local/go/src/net/netip/netip.go:1033
}

// Addr returns p's IP address.
func (p AddrPort) Addr() Addr {
//line /usr/local/go/src/net/netip/netip.go:1036
	_go_fuzz_dep_.CoverTab[3953]++
//line /usr/local/go/src/net/netip/netip.go:1036
	return p.ip
//line /usr/local/go/src/net/netip/netip.go:1036
	// _ = "end of CoverTab[3953]"
//line /usr/local/go/src/net/netip/netip.go:1036
}

// Port returns p's port.
func (p AddrPort) Port() uint16 {
//line /usr/local/go/src/net/netip/netip.go:1039
	_go_fuzz_dep_.CoverTab[3954]++
//line /usr/local/go/src/net/netip/netip.go:1039
	return p.port
//line /usr/local/go/src/net/netip/netip.go:1039
	// _ = "end of CoverTab[3954]"
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
	_go_fuzz_dep_.CoverTab[3955]++
							i := stringsLastIndexByte(s, ':')
							if i == -1 {
//line /usr/local/go/src/net/netip/netip.go:1048
		_go_fuzz_dep_.CoverTab[3960]++
								return "", "", false, errors.New("not an ip:port")
//line /usr/local/go/src/net/netip/netip.go:1049
		// _ = "end of CoverTab[3960]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1050
		_go_fuzz_dep_.CoverTab[3961]++
//line /usr/local/go/src/net/netip/netip.go:1050
		// _ = "end of CoverTab[3961]"
//line /usr/local/go/src/net/netip/netip.go:1050
	}
//line /usr/local/go/src/net/netip/netip.go:1050
	// _ = "end of CoverTab[3955]"
//line /usr/local/go/src/net/netip/netip.go:1050
	_go_fuzz_dep_.CoverTab[3956]++

							ip, port = s[:i], s[i+1:]
							if len(ip) == 0 {
//line /usr/local/go/src/net/netip/netip.go:1053
		_go_fuzz_dep_.CoverTab[3962]++
								return "", "", false, errors.New("no IP")
//line /usr/local/go/src/net/netip/netip.go:1054
		// _ = "end of CoverTab[3962]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1055
		_go_fuzz_dep_.CoverTab[3963]++
//line /usr/local/go/src/net/netip/netip.go:1055
		// _ = "end of CoverTab[3963]"
//line /usr/local/go/src/net/netip/netip.go:1055
	}
//line /usr/local/go/src/net/netip/netip.go:1055
	// _ = "end of CoverTab[3956]"
//line /usr/local/go/src/net/netip/netip.go:1055
	_go_fuzz_dep_.CoverTab[3957]++
							if len(port) == 0 {
//line /usr/local/go/src/net/netip/netip.go:1056
		_go_fuzz_dep_.CoverTab[3964]++
								return "", "", false, errors.New("no port")
//line /usr/local/go/src/net/netip/netip.go:1057
		// _ = "end of CoverTab[3964]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1058
		_go_fuzz_dep_.CoverTab[3965]++
//line /usr/local/go/src/net/netip/netip.go:1058
		// _ = "end of CoverTab[3965]"
//line /usr/local/go/src/net/netip/netip.go:1058
	}
//line /usr/local/go/src/net/netip/netip.go:1058
	// _ = "end of CoverTab[3957]"
//line /usr/local/go/src/net/netip/netip.go:1058
	_go_fuzz_dep_.CoverTab[3958]++
							if ip[0] == '[' {
//line /usr/local/go/src/net/netip/netip.go:1059
		_go_fuzz_dep_.CoverTab[3966]++
								if len(ip) < 2 || func() bool {
//line /usr/local/go/src/net/netip/netip.go:1060
			_go_fuzz_dep_.CoverTab[3968]++
//line /usr/local/go/src/net/netip/netip.go:1060
			return ip[len(ip)-1] != ']'
//line /usr/local/go/src/net/netip/netip.go:1060
			// _ = "end of CoverTab[3968]"
//line /usr/local/go/src/net/netip/netip.go:1060
		}() {
//line /usr/local/go/src/net/netip/netip.go:1060
			_go_fuzz_dep_.CoverTab[3969]++
									return "", "", false, errors.New("missing ]")
//line /usr/local/go/src/net/netip/netip.go:1061
			// _ = "end of CoverTab[3969]"
		} else {
//line /usr/local/go/src/net/netip/netip.go:1062
			_go_fuzz_dep_.CoverTab[3970]++
//line /usr/local/go/src/net/netip/netip.go:1062
			// _ = "end of CoverTab[3970]"
//line /usr/local/go/src/net/netip/netip.go:1062
		}
//line /usr/local/go/src/net/netip/netip.go:1062
		// _ = "end of CoverTab[3966]"
//line /usr/local/go/src/net/netip/netip.go:1062
		_go_fuzz_dep_.CoverTab[3967]++
								ip = ip[1 : len(ip)-1]
								v6 = true
//line /usr/local/go/src/net/netip/netip.go:1064
		// _ = "end of CoverTab[3967]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1065
		_go_fuzz_dep_.CoverTab[3971]++
//line /usr/local/go/src/net/netip/netip.go:1065
		// _ = "end of CoverTab[3971]"
//line /usr/local/go/src/net/netip/netip.go:1065
	}
//line /usr/local/go/src/net/netip/netip.go:1065
	// _ = "end of CoverTab[3958]"
//line /usr/local/go/src/net/netip/netip.go:1065
	_go_fuzz_dep_.CoverTab[3959]++

							return ip, port, v6, nil
//line /usr/local/go/src/net/netip/netip.go:1067
	// _ = "end of CoverTab[3959]"
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
	_go_fuzz_dep_.CoverTab[3972]++
							var ipp AddrPort
							ip, port, v6, err := splitAddrPort(s)
							if err != nil {
//line /usr/local/go/src/net/netip/netip.go:1077
		_go_fuzz_dep_.CoverTab[3977]++
								return ipp, err
//line /usr/local/go/src/net/netip/netip.go:1078
		// _ = "end of CoverTab[3977]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1079
		_go_fuzz_dep_.CoverTab[3978]++
//line /usr/local/go/src/net/netip/netip.go:1079
		// _ = "end of CoverTab[3978]"
//line /usr/local/go/src/net/netip/netip.go:1079
	}
//line /usr/local/go/src/net/netip/netip.go:1079
	// _ = "end of CoverTab[3972]"
//line /usr/local/go/src/net/netip/netip.go:1079
	_go_fuzz_dep_.CoverTab[3973]++
							port16, err := strconv.ParseUint(port, 10, 16)
							if err != nil {
//line /usr/local/go/src/net/netip/netip.go:1081
		_go_fuzz_dep_.CoverTab[3979]++
								return ipp, errors.New("invalid port " + strconv.Quote(port) + " parsing " + strconv.Quote(s))
//line /usr/local/go/src/net/netip/netip.go:1082
		// _ = "end of CoverTab[3979]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1083
		_go_fuzz_dep_.CoverTab[3980]++
//line /usr/local/go/src/net/netip/netip.go:1083
		// _ = "end of CoverTab[3980]"
//line /usr/local/go/src/net/netip/netip.go:1083
	}
//line /usr/local/go/src/net/netip/netip.go:1083
	// _ = "end of CoverTab[3973]"
//line /usr/local/go/src/net/netip/netip.go:1083
	_go_fuzz_dep_.CoverTab[3974]++
							ipp.port = uint16(port16)
							ipp.ip, err = ParseAddr(ip)
							if err != nil {
//line /usr/local/go/src/net/netip/netip.go:1086
		_go_fuzz_dep_.CoverTab[3981]++
								return AddrPort{}, err
//line /usr/local/go/src/net/netip/netip.go:1087
		// _ = "end of CoverTab[3981]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1088
		_go_fuzz_dep_.CoverTab[3982]++
//line /usr/local/go/src/net/netip/netip.go:1088
		// _ = "end of CoverTab[3982]"
//line /usr/local/go/src/net/netip/netip.go:1088
	}
//line /usr/local/go/src/net/netip/netip.go:1088
	// _ = "end of CoverTab[3974]"
//line /usr/local/go/src/net/netip/netip.go:1088
	_go_fuzz_dep_.CoverTab[3975]++
							if v6 && func() bool {
//line /usr/local/go/src/net/netip/netip.go:1089
		_go_fuzz_dep_.CoverTab[3983]++
//line /usr/local/go/src/net/netip/netip.go:1089
		return ipp.ip.Is4()
//line /usr/local/go/src/net/netip/netip.go:1089
		// _ = "end of CoverTab[3983]"
//line /usr/local/go/src/net/netip/netip.go:1089
	}() {
//line /usr/local/go/src/net/netip/netip.go:1089
		_go_fuzz_dep_.CoverTab[3984]++
								return AddrPort{}, errors.New("invalid ip:port " + strconv.Quote(s) + ", square brackets can only be used with IPv6 addresses")
//line /usr/local/go/src/net/netip/netip.go:1090
		// _ = "end of CoverTab[3984]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1091
		_go_fuzz_dep_.CoverTab[3985]++
//line /usr/local/go/src/net/netip/netip.go:1091
		if !v6 && func() bool {
//line /usr/local/go/src/net/netip/netip.go:1091
			_go_fuzz_dep_.CoverTab[3986]++
//line /usr/local/go/src/net/netip/netip.go:1091
			return ipp.ip.Is6()
//line /usr/local/go/src/net/netip/netip.go:1091
			// _ = "end of CoverTab[3986]"
//line /usr/local/go/src/net/netip/netip.go:1091
		}() {
//line /usr/local/go/src/net/netip/netip.go:1091
			_go_fuzz_dep_.CoverTab[3987]++
									return AddrPort{}, errors.New("invalid ip:port " + strconv.Quote(s) + ", IPv6 addresses must be surrounded by square brackets")
//line /usr/local/go/src/net/netip/netip.go:1092
			// _ = "end of CoverTab[3987]"
		} else {
//line /usr/local/go/src/net/netip/netip.go:1093
			_go_fuzz_dep_.CoverTab[3988]++
//line /usr/local/go/src/net/netip/netip.go:1093
			// _ = "end of CoverTab[3988]"
//line /usr/local/go/src/net/netip/netip.go:1093
		}
//line /usr/local/go/src/net/netip/netip.go:1093
		// _ = "end of CoverTab[3985]"
//line /usr/local/go/src/net/netip/netip.go:1093
	}
//line /usr/local/go/src/net/netip/netip.go:1093
	// _ = "end of CoverTab[3975]"
//line /usr/local/go/src/net/netip/netip.go:1093
	_go_fuzz_dep_.CoverTab[3976]++
							return ipp, nil
//line /usr/local/go/src/net/netip/netip.go:1094
	// _ = "end of CoverTab[3976]"
}

// MustParseAddrPort calls ParseAddrPort(s) and panics on error.
//line /usr/local/go/src/net/netip/netip.go:1097
// It is intended for use in tests with hard-coded strings.
//line /usr/local/go/src/net/netip/netip.go:1099
func MustParseAddrPort(s string) AddrPort {
//line /usr/local/go/src/net/netip/netip.go:1099
	_go_fuzz_dep_.CoverTab[3989]++
							ip, err := ParseAddrPort(s)
							if err != nil {
//line /usr/local/go/src/net/netip/netip.go:1101
		_go_fuzz_dep_.CoverTab[3991]++
								panic(err)
//line /usr/local/go/src/net/netip/netip.go:1102
		// _ = "end of CoverTab[3991]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1103
		_go_fuzz_dep_.CoverTab[3992]++
//line /usr/local/go/src/net/netip/netip.go:1103
		// _ = "end of CoverTab[3992]"
//line /usr/local/go/src/net/netip/netip.go:1103
	}
//line /usr/local/go/src/net/netip/netip.go:1103
	// _ = "end of CoverTab[3989]"
//line /usr/local/go/src/net/netip/netip.go:1103
	_go_fuzz_dep_.CoverTab[3990]++
							return ip
//line /usr/local/go/src/net/netip/netip.go:1104
	// _ = "end of CoverTab[3990]"
}

// IsValid reports whether p.Addr() is valid.
//line /usr/local/go/src/net/netip/netip.go:1107
// All ports are valid, including zero.
//line /usr/local/go/src/net/netip/netip.go:1109
func (p AddrPort) IsValid() bool {
//line /usr/local/go/src/net/netip/netip.go:1109
	_go_fuzz_dep_.CoverTab[3993]++
//line /usr/local/go/src/net/netip/netip.go:1109
	return p.ip.IsValid()
//line /usr/local/go/src/net/netip/netip.go:1109
	// _ = "end of CoverTab[3993]"
//line /usr/local/go/src/net/netip/netip.go:1109
}

func (p AddrPort) String() string {
//line /usr/local/go/src/net/netip/netip.go:1111
	_go_fuzz_dep_.CoverTab[3994]++
							switch p.ip.z {
	case z0:
//line /usr/local/go/src/net/netip/netip.go:1113
		_go_fuzz_dep_.CoverTab[3995]++
								return "invalid AddrPort"
//line /usr/local/go/src/net/netip/netip.go:1114
		// _ = "end of CoverTab[3995]"
	case z4:
//line /usr/local/go/src/net/netip/netip.go:1115
		_go_fuzz_dep_.CoverTab[3996]++
								a := p.ip.As4()
								buf := make([]byte, 0, 21)
								for i := range a {
//line /usr/local/go/src/net/netip/netip.go:1118
			_go_fuzz_dep_.CoverTab[3999]++
									buf = strconv.AppendUint(buf, uint64(a[i]), 10)
									buf = append(buf, "...:"[i])
//line /usr/local/go/src/net/netip/netip.go:1120
			// _ = "end of CoverTab[3999]"
		}
//line /usr/local/go/src/net/netip/netip.go:1121
		// _ = "end of CoverTab[3996]"
//line /usr/local/go/src/net/netip/netip.go:1121
		_go_fuzz_dep_.CoverTab[3997]++
								buf = strconv.AppendUint(buf, uint64(p.port), 10)
								return string(buf)
//line /usr/local/go/src/net/netip/netip.go:1123
		// _ = "end of CoverTab[3997]"
	default:
//line /usr/local/go/src/net/netip/netip.go:1124
		_go_fuzz_dep_.CoverTab[3998]++

								return joinHostPort(p.ip.String(), itoa.Itoa(int(p.port)))
//line /usr/local/go/src/net/netip/netip.go:1126
		// _ = "end of CoverTab[3998]"
	}
//line /usr/local/go/src/net/netip/netip.go:1127
	// _ = "end of CoverTab[3994]"
}

func joinHostPort(host, port string) string {
//line /usr/local/go/src/net/netip/netip.go:1130
	_go_fuzz_dep_.CoverTab[4000]++

//line /usr/local/go/src/net/netip/netip.go:1133
	if bytealg.IndexByteString(host, ':') >= 0 {
//line /usr/local/go/src/net/netip/netip.go:1133
		_go_fuzz_dep_.CoverTab[4002]++
								return "[" + host + "]:" + port
//line /usr/local/go/src/net/netip/netip.go:1134
		// _ = "end of CoverTab[4002]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1135
		_go_fuzz_dep_.CoverTab[4003]++
//line /usr/local/go/src/net/netip/netip.go:1135
		// _ = "end of CoverTab[4003]"
//line /usr/local/go/src/net/netip/netip.go:1135
	}
//line /usr/local/go/src/net/netip/netip.go:1135
	// _ = "end of CoverTab[4000]"
//line /usr/local/go/src/net/netip/netip.go:1135
	_go_fuzz_dep_.CoverTab[4001]++
							return host + ":" + port
//line /usr/local/go/src/net/netip/netip.go:1136
	// _ = "end of CoverTab[4001]"
}

// AppendTo appends a text encoding of p,
//line /usr/local/go/src/net/netip/netip.go:1139
// as generated by MarshalText,
//line /usr/local/go/src/net/netip/netip.go:1139
// to b and returns the extended buffer.
//line /usr/local/go/src/net/netip/netip.go:1142
func (p AddrPort) AppendTo(b []byte) []byte {
//line /usr/local/go/src/net/netip/netip.go:1142
	_go_fuzz_dep_.CoverTab[4004]++
							switch p.ip.z {
	case z0:
//line /usr/local/go/src/net/netip/netip.go:1144
		_go_fuzz_dep_.CoverTab[4006]++
								return b
//line /usr/local/go/src/net/netip/netip.go:1145
		// _ = "end of CoverTab[4006]"
	case z4:
//line /usr/local/go/src/net/netip/netip.go:1146
		_go_fuzz_dep_.CoverTab[4007]++
								b = p.ip.appendTo4(b)
//line /usr/local/go/src/net/netip/netip.go:1147
		// _ = "end of CoverTab[4007]"
	default:
//line /usr/local/go/src/net/netip/netip.go:1148
		_go_fuzz_dep_.CoverTab[4008]++
								if p.ip.Is4In6() {
//line /usr/local/go/src/net/netip/netip.go:1149
			_go_fuzz_dep_.CoverTab[4010]++
									b = append(b, "[::ffff:"...)
									b = p.ip.Unmap().appendTo4(b)
									if z := p.ip.Zone(); z != "" {
//line /usr/local/go/src/net/netip/netip.go:1152
				_go_fuzz_dep_.CoverTab[4011]++
										b = append(b, '%')
										b = append(b, z...)
//line /usr/local/go/src/net/netip/netip.go:1154
				// _ = "end of CoverTab[4011]"
			} else {
//line /usr/local/go/src/net/netip/netip.go:1155
				_go_fuzz_dep_.CoverTab[4012]++
//line /usr/local/go/src/net/netip/netip.go:1155
				// _ = "end of CoverTab[4012]"
//line /usr/local/go/src/net/netip/netip.go:1155
			}
//line /usr/local/go/src/net/netip/netip.go:1155
			// _ = "end of CoverTab[4010]"
		} else {
//line /usr/local/go/src/net/netip/netip.go:1156
			_go_fuzz_dep_.CoverTab[4013]++
									b = append(b, '[')
									b = p.ip.appendTo6(b)
//line /usr/local/go/src/net/netip/netip.go:1158
			// _ = "end of CoverTab[4013]"
		}
//line /usr/local/go/src/net/netip/netip.go:1159
		// _ = "end of CoverTab[4008]"
//line /usr/local/go/src/net/netip/netip.go:1159
		_go_fuzz_dep_.CoverTab[4009]++
								b = append(b, ']')
//line /usr/local/go/src/net/netip/netip.go:1160
		// _ = "end of CoverTab[4009]"
	}
//line /usr/local/go/src/net/netip/netip.go:1161
	// _ = "end of CoverTab[4004]"
//line /usr/local/go/src/net/netip/netip.go:1161
	_go_fuzz_dep_.CoverTab[4005]++
							b = append(b, ':')
							b = strconv.AppendUint(b, uint64(p.port), 10)
							return b
//line /usr/local/go/src/net/netip/netip.go:1164
	// _ = "end of CoverTab[4005]"
}

// MarshalText implements the encoding.TextMarshaler interface. The
//line /usr/local/go/src/net/netip/netip.go:1167
// encoding is the same as returned by String, with one exception: if
//line /usr/local/go/src/net/netip/netip.go:1167
// p.Addr() is the zero Addr, the encoding is the empty string.
//line /usr/local/go/src/net/netip/netip.go:1170
func (p AddrPort) MarshalText() ([]byte, error) {
//line /usr/local/go/src/net/netip/netip.go:1170
	_go_fuzz_dep_.CoverTab[4014]++
							var max int
							switch p.ip.z {
	case z0:
//line /usr/local/go/src/net/netip/netip.go:1173
		_go_fuzz_dep_.CoverTab[4016]++
//line /usr/local/go/src/net/netip/netip.go:1173
		// _ = "end of CoverTab[4016]"
	case z4:
//line /usr/local/go/src/net/netip/netip.go:1174
		_go_fuzz_dep_.CoverTab[4017]++
								max = len("255.255.255.255:65535")
//line /usr/local/go/src/net/netip/netip.go:1175
		// _ = "end of CoverTab[4017]"
	default:
//line /usr/local/go/src/net/netip/netip.go:1176
		_go_fuzz_dep_.CoverTab[4018]++
								max = len("[ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff%enp5s0]:65535")
//line /usr/local/go/src/net/netip/netip.go:1177
		// _ = "end of CoverTab[4018]"
	}
//line /usr/local/go/src/net/netip/netip.go:1178
	// _ = "end of CoverTab[4014]"
//line /usr/local/go/src/net/netip/netip.go:1178
	_go_fuzz_dep_.CoverTab[4015]++
							b := make([]byte, 0, max)
							b = p.AppendTo(b)
							return b, nil
//line /usr/local/go/src/net/netip/netip.go:1181
	// _ = "end of CoverTab[4015]"
}

// UnmarshalText implements the encoding.TextUnmarshaler
//line /usr/local/go/src/net/netip/netip.go:1184
// interface. The AddrPort is expected in a form
//line /usr/local/go/src/net/netip/netip.go:1184
// generated by MarshalText or accepted by ParseAddrPort.
//line /usr/local/go/src/net/netip/netip.go:1187
func (p *AddrPort) UnmarshalText(text []byte) error {
//line /usr/local/go/src/net/netip/netip.go:1187
	_go_fuzz_dep_.CoverTab[4019]++
							if len(text) == 0 {
//line /usr/local/go/src/net/netip/netip.go:1188
		_go_fuzz_dep_.CoverTab[4021]++
								*p = AddrPort{}
								return nil
//line /usr/local/go/src/net/netip/netip.go:1190
		// _ = "end of CoverTab[4021]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1191
		_go_fuzz_dep_.CoverTab[4022]++
//line /usr/local/go/src/net/netip/netip.go:1191
		// _ = "end of CoverTab[4022]"
//line /usr/local/go/src/net/netip/netip.go:1191
	}
//line /usr/local/go/src/net/netip/netip.go:1191
	// _ = "end of CoverTab[4019]"
//line /usr/local/go/src/net/netip/netip.go:1191
	_go_fuzz_dep_.CoverTab[4020]++
							var err error
							*p, err = ParseAddrPort(string(text))
							return err
//line /usr/local/go/src/net/netip/netip.go:1194
	// _ = "end of CoverTab[4020]"
}

// MarshalBinary implements the encoding.BinaryMarshaler interface.
//line /usr/local/go/src/net/netip/netip.go:1197
// It returns Addr.MarshalBinary with an additional two bytes appended
//line /usr/local/go/src/net/netip/netip.go:1197
// containing the port in little-endian.
//line /usr/local/go/src/net/netip/netip.go:1200
func (p AddrPort) MarshalBinary() ([]byte, error) {
//line /usr/local/go/src/net/netip/netip.go:1200
	_go_fuzz_dep_.CoverTab[4023]++
							b := p.Addr().marshalBinaryWithTrailingBytes(2)
							lePutUint16(b[len(b)-2:], p.Port())
							return b, nil
//line /usr/local/go/src/net/netip/netip.go:1203
	// _ = "end of CoverTab[4023]"
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface.
//line /usr/local/go/src/net/netip/netip.go:1206
// It expects data in the form generated by MarshalBinary.
//line /usr/local/go/src/net/netip/netip.go:1208
func (p *AddrPort) UnmarshalBinary(b []byte) error {
//line /usr/local/go/src/net/netip/netip.go:1208
	_go_fuzz_dep_.CoverTab[4024]++
							if len(b) < 2 {
//line /usr/local/go/src/net/netip/netip.go:1209
		_go_fuzz_dep_.CoverTab[4027]++
								return errors.New("unexpected slice size")
//line /usr/local/go/src/net/netip/netip.go:1210
		// _ = "end of CoverTab[4027]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1211
		_go_fuzz_dep_.CoverTab[4028]++
//line /usr/local/go/src/net/netip/netip.go:1211
		// _ = "end of CoverTab[4028]"
//line /usr/local/go/src/net/netip/netip.go:1211
	}
//line /usr/local/go/src/net/netip/netip.go:1211
	// _ = "end of CoverTab[4024]"
//line /usr/local/go/src/net/netip/netip.go:1211
	_go_fuzz_dep_.CoverTab[4025]++
							var addr Addr
							err := addr.UnmarshalBinary(b[:len(b)-2])
							if err != nil {
//line /usr/local/go/src/net/netip/netip.go:1214
		_go_fuzz_dep_.CoverTab[4029]++
								return err
//line /usr/local/go/src/net/netip/netip.go:1215
		// _ = "end of CoverTab[4029]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1216
		_go_fuzz_dep_.CoverTab[4030]++
//line /usr/local/go/src/net/netip/netip.go:1216
		// _ = "end of CoverTab[4030]"
//line /usr/local/go/src/net/netip/netip.go:1216
	}
//line /usr/local/go/src/net/netip/netip.go:1216
	// _ = "end of CoverTab[4025]"
//line /usr/local/go/src/net/netip/netip.go:1216
	_go_fuzz_dep_.CoverTab[4026]++
							*p = AddrPortFrom(addr, leUint16(b[len(b)-2:]))
							return nil
//line /usr/local/go/src/net/netip/netip.go:1218
	// _ = "end of CoverTab[4026]"
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
	_go_fuzz_dep_.CoverTab[4031]++
							if bits < 0 || func() bool {
//line /usr/local/go/src/net/netip/netip.go:1251
		_go_fuzz_dep_.CoverTab[4033]++
//line /usr/local/go/src/net/netip/netip.go:1251
		return bits > ip.BitLen()
//line /usr/local/go/src/net/netip/netip.go:1251
		// _ = "end of CoverTab[4033]"
//line /usr/local/go/src/net/netip/netip.go:1251
	}() {
//line /usr/local/go/src/net/netip/netip.go:1251
		_go_fuzz_dep_.CoverTab[4034]++
								bits = invalidPrefixBits
//line /usr/local/go/src/net/netip/netip.go:1252
		// _ = "end of CoverTab[4034]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1253
		_go_fuzz_dep_.CoverTab[4035]++
//line /usr/local/go/src/net/netip/netip.go:1253
		// _ = "end of CoverTab[4035]"
//line /usr/local/go/src/net/netip/netip.go:1253
	}
//line /usr/local/go/src/net/netip/netip.go:1253
	// _ = "end of CoverTab[4031]"
//line /usr/local/go/src/net/netip/netip.go:1253
	_go_fuzz_dep_.CoverTab[4032]++
							b16 := int16(bits)
							return Prefix{
		ip:	ip.withoutZone(),
		bits:	b16,
	}
//line /usr/local/go/src/net/netip/netip.go:1258
	// _ = "end of CoverTab[4032]"
}

// Addr returns p's IP address.
func (p Prefix) Addr() Addr {
//line /usr/local/go/src/net/netip/netip.go:1262
	_go_fuzz_dep_.CoverTab[4036]++
//line /usr/local/go/src/net/netip/netip.go:1262
	return p.ip
//line /usr/local/go/src/net/netip/netip.go:1262
	// _ = "end of CoverTab[4036]"
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
	_go_fuzz_dep_.CoverTab[4037]++
//line /usr/local/go/src/net/netip/netip.go:1267
	return int(p.bits)
//line /usr/local/go/src/net/netip/netip.go:1267
	// _ = "end of CoverTab[4037]"
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
	_go_fuzz_dep_.CoverTab[4038]++
//line /usr/local/go/src/net/netip/netip.go:1272
	return !p.ip.isZero() && func() bool {
//line /usr/local/go/src/net/netip/netip.go:1272
		_go_fuzz_dep_.CoverTab[4039]++
//line /usr/local/go/src/net/netip/netip.go:1272
		return p.bits >= 0
//line /usr/local/go/src/net/netip/netip.go:1272
		// _ = "end of CoverTab[4039]"
//line /usr/local/go/src/net/netip/netip.go:1272
	}() && func() bool {
//line /usr/local/go/src/net/netip/netip.go:1272
		_go_fuzz_dep_.CoverTab[4040]++
//line /usr/local/go/src/net/netip/netip.go:1272
		return int(p.bits) <= p.ip.BitLen()
//line /usr/local/go/src/net/netip/netip.go:1272
		// _ = "end of CoverTab[4040]"
//line /usr/local/go/src/net/netip/netip.go:1272
	}()
//line /usr/local/go/src/net/netip/netip.go:1272
	// _ = "end of CoverTab[4038]"
//line /usr/local/go/src/net/netip/netip.go:1272
}

func (p Prefix) isZero() bool {
//line /usr/local/go/src/net/netip/netip.go:1274
	_go_fuzz_dep_.CoverTab[4041]++
//line /usr/local/go/src/net/netip/netip.go:1274
	return p == Prefix{}
//line /usr/local/go/src/net/netip/netip.go:1274
	// _ = "end of CoverTab[4041]"
//line /usr/local/go/src/net/netip/netip.go:1274
}

// IsSingleIP reports whether p contains exactly one IP.
func (p Prefix) IsSingleIP() bool {
//line /usr/local/go/src/net/netip/netip.go:1277
	_go_fuzz_dep_.CoverTab[4042]++
//line /usr/local/go/src/net/netip/netip.go:1277
	return p.bits != 0 && func() bool {
//line /usr/local/go/src/net/netip/netip.go:1277
		_go_fuzz_dep_.CoverTab[4043]++
//line /usr/local/go/src/net/netip/netip.go:1277
		return int(p.bits) == p.ip.BitLen()
//line /usr/local/go/src/net/netip/netip.go:1277
		// _ = "end of CoverTab[4043]"
//line /usr/local/go/src/net/netip/netip.go:1277
	}()
//line /usr/local/go/src/net/netip/netip.go:1277
	// _ = "end of CoverTab[4042]"
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
	_go_fuzz_dep_.CoverTab[4044]++
							i := stringsLastIndexByte(s, '/')
							if i < 0 {
//line /usr/local/go/src/net/netip/netip.go:1288
		_go_fuzz_dep_.CoverTab[4051]++
								return Prefix{}, errors.New("netip.ParsePrefix(" + strconv.Quote(s) + "): no '/'")
//line /usr/local/go/src/net/netip/netip.go:1289
		// _ = "end of CoverTab[4051]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1290
		_go_fuzz_dep_.CoverTab[4052]++
//line /usr/local/go/src/net/netip/netip.go:1290
		// _ = "end of CoverTab[4052]"
//line /usr/local/go/src/net/netip/netip.go:1290
	}
//line /usr/local/go/src/net/netip/netip.go:1290
	// _ = "end of CoverTab[4044]"
//line /usr/local/go/src/net/netip/netip.go:1290
	_go_fuzz_dep_.CoverTab[4045]++
							ip, err := ParseAddr(s[:i])
							if err != nil {
//line /usr/local/go/src/net/netip/netip.go:1292
		_go_fuzz_dep_.CoverTab[4053]++
								return Prefix{}, errors.New("netip.ParsePrefix(" + strconv.Quote(s) + "): " + err.Error())
//line /usr/local/go/src/net/netip/netip.go:1293
		// _ = "end of CoverTab[4053]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1294
		_go_fuzz_dep_.CoverTab[4054]++
//line /usr/local/go/src/net/netip/netip.go:1294
		// _ = "end of CoverTab[4054]"
//line /usr/local/go/src/net/netip/netip.go:1294
	}
//line /usr/local/go/src/net/netip/netip.go:1294
	// _ = "end of CoverTab[4045]"
//line /usr/local/go/src/net/netip/netip.go:1294
	_go_fuzz_dep_.CoverTab[4046]++

							if ip.Is6() && func() bool {
//line /usr/local/go/src/net/netip/netip.go:1296
		_go_fuzz_dep_.CoverTab[4055]++
//line /usr/local/go/src/net/netip/netip.go:1296
		return ip.z != z6noz
//line /usr/local/go/src/net/netip/netip.go:1296
		// _ = "end of CoverTab[4055]"
//line /usr/local/go/src/net/netip/netip.go:1296
	}() {
//line /usr/local/go/src/net/netip/netip.go:1296
		_go_fuzz_dep_.CoverTab[4056]++
								return Prefix{}, errors.New("netip.ParsePrefix(" + strconv.Quote(s) + "): IPv6 zones cannot be present in a prefix")
//line /usr/local/go/src/net/netip/netip.go:1297
		// _ = "end of CoverTab[4056]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1298
		_go_fuzz_dep_.CoverTab[4057]++
//line /usr/local/go/src/net/netip/netip.go:1298
		// _ = "end of CoverTab[4057]"
//line /usr/local/go/src/net/netip/netip.go:1298
	}
//line /usr/local/go/src/net/netip/netip.go:1298
	// _ = "end of CoverTab[4046]"
//line /usr/local/go/src/net/netip/netip.go:1298
	_go_fuzz_dep_.CoverTab[4047]++

							bitsStr := s[i+1:]
							bits, err := strconv.Atoi(bitsStr)
							if err != nil {
//line /usr/local/go/src/net/netip/netip.go:1302
		_go_fuzz_dep_.CoverTab[4058]++
								return Prefix{}, errors.New("netip.ParsePrefix(" + strconv.Quote(s) + "): bad bits after slash: " + strconv.Quote(bitsStr))
//line /usr/local/go/src/net/netip/netip.go:1303
		// _ = "end of CoverTab[4058]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1304
		_go_fuzz_dep_.CoverTab[4059]++
//line /usr/local/go/src/net/netip/netip.go:1304
		// _ = "end of CoverTab[4059]"
//line /usr/local/go/src/net/netip/netip.go:1304
	}
//line /usr/local/go/src/net/netip/netip.go:1304
	// _ = "end of CoverTab[4047]"
//line /usr/local/go/src/net/netip/netip.go:1304
	_go_fuzz_dep_.CoverTab[4048]++
							maxBits := 32
							if ip.Is6() {
//line /usr/local/go/src/net/netip/netip.go:1306
		_go_fuzz_dep_.CoverTab[4060]++
								maxBits = 128
//line /usr/local/go/src/net/netip/netip.go:1307
		// _ = "end of CoverTab[4060]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1308
		_go_fuzz_dep_.CoverTab[4061]++
//line /usr/local/go/src/net/netip/netip.go:1308
		// _ = "end of CoverTab[4061]"
//line /usr/local/go/src/net/netip/netip.go:1308
	}
//line /usr/local/go/src/net/netip/netip.go:1308
	// _ = "end of CoverTab[4048]"
//line /usr/local/go/src/net/netip/netip.go:1308
	_go_fuzz_dep_.CoverTab[4049]++
							if bits < 0 || func() bool {
//line /usr/local/go/src/net/netip/netip.go:1309
		_go_fuzz_dep_.CoverTab[4062]++
//line /usr/local/go/src/net/netip/netip.go:1309
		return bits > maxBits
//line /usr/local/go/src/net/netip/netip.go:1309
		// _ = "end of CoverTab[4062]"
//line /usr/local/go/src/net/netip/netip.go:1309
	}() {
//line /usr/local/go/src/net/netip/netip.go:1309
		_go_fuzz_dep_.CoverTab[4063]++
								return Prefix{}, errors.New("netip.ParsePrefix(" + strconv.Quote(s) + "): prefix length out of range")
//line /usr/local/go/src/net/netip/netip.go:1310
		// _ = "end of CoverTab[4063]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1311
		_go_fuzz_dep_.CoverTab[4064]++
//line /usr/local/go/src/net/netip/netip.go:1311
		// _ = "end of CoverTab[4064]"
//line /usr/local/go/src/net/netip/netip.go:1311
	}
//line /usr/local/go/src/net/netip/netip.go:1311
	// _ = "end of CoverTab[4049]"
//line /usr/local/go/src/net/netip/netip.go:1311
	_go_fuzz_dep_.CoverTab[4050]++
							return PrefixFrom(ip, bits), nil
//line /usr/local/go/src/net/netip/netip.go:1312
	// _ = "end of CoverTab[4050]"
}

// MustParsePrefix calls ParsePrefix(s) and panics on error.
//line /usr/local/go/src/net/netip/netip.go:1315
// It is intended for use in tests with hard-coded strings.
//line /usr/local/go/src/net/netip/netip.go:1317
func MustParsePrefix(s string) Prefix {
//line /usr/local/go/src/net/netip/netip.go:1317
	_go_fuzz_dep_.CoverTab[4065]++
							ip, err := ParsePrefix(s)
							if err != nil {
//line /usr/local/go/src/net/netip/netip.go:1319
		_go_fuzz_dep_.CoverTab[4067]++
								panic(err)
//line /usr/local/go/src/net/netip/netip.go:1320
		// _ = "end of CoverTab[4067]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1321
		_go_fuzz_dep_.CoverTab[4068]++
//line /usr/local/go/src/net/netip/netip.go:1321
		// _ = "end of CoverTab[4068]"
//line /usr/local/go/src/net/netip/netip.go:1321
	}
//line /usr/local/go/src/net/netip/netip.go:1321
	// _ = "end of CoverTab[4065]"
//line /usr/local/go/src/net/netip/netip.go:1321
	_go_fuzz_dep_.CoverTab[4066]++
							return ip
//line /usr/local/go/src/net/netip/netip.go:1322
	// _ = "end of CoverTab[4066]"
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
	_go_fuzz_dep_.CoverTab[4069]++
							if m, err := p.ip.Prefix(int(p.bits)); err == nil {
//line /usr/local/go/src/net/netip/netip.go:1330
		_go_fuzz_dep_.CoverTab[4071]++
								return m
//line /usr/local/go/src/net/netip/netip.go:1331
		// _ = "end of CoverTab[4071]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1332
		_go_fuzz_dep_.CoverTab[4072]++
//line /usr/local/go/src/net/netip/netip.go:1332
		// _ = "end of CoverTab[4072]"
//line /usr/local/go/src/net/netip/netip.go:1332
	}
//line /usr/local/go/src/net/netip/netip.go:1332
	// _ = "end of CoverTab[4069]"
//line /usr/local/go/src/net/netip/netip.go:1332
	_go_fuzz_dep_.CoverTab[4070]++
							return Prefix{}
//line /usr/local/go/src/net/netip/netip.go:1333
	// _ = "end of CoverTab[4070]"
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
	_go_fuzz_dep_.CoverTab[4073]++
							if !p.IsValid() || func() bool {
//line /usr/local/go/src/net/netip/netip.go:1344
		_go_fuzz_dep_.CoverTab[4076]++
//line /usr/local/go/src/net/netip/netip.go:1344
		return ip.hasZone()
//line /usr/local/go/src/net/netip/netip.go:1344
		// _ = "end of CoverTab[4076]"
//line /usr/local/go/src/net/netip/netip.go:1344
	}() {
//line /usr/local/go/src/net/netip/netip.go:1344
		_go_fuzz_dep_.CoverTab[4077]++
								return false
//line /usr/local/go/src/net/netip/netip.go:1345
		// _ = "end of CoverTab[4077]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1346
		_go_fuzz_dep_.CoverTab[4078]++
//line /usr/local/go/src/net/netip/netip.go:1346
		// _ = "end of CoverTab[4078]"
//line /usr/local/go/src/net/netip/netip.go:1346
	}
//line /usr/local/go/src/net/netip/netip.go:1346
	// _ = "end of CoverTab[4073]"
//line /usr/local/go/src/net/netip/netip.go:1346
	_go_fuzz_dep_.CoverTab[4074]++
							if f1, f2 := p.ip.BitLen(), ip.BitLen(); f1 == 0 || func() bool {
//line /usr/local/go/src/net/netip/netip.go:1347
		_go_fuzz_dep_.CoverTab[4079]++
//line /usr/local/go/src/net/netip/netip.go:1347
		return f2 == 0
//line /usr/local/go/src/net/netip/netip.go:1347
		// _ = "end of CoverTab[4079]"
//line /usr/local/go/src/net/netip/netip.go:1347
	}() || func() bool {
//line /usr/local/go/src/net/netip/netip.go:1347
		_go_fuzz_dep_.CoverTab[4080]++
//line /usr/local/go/src/net/netip/netip.go:1347
		return f1 != f2
//line /usr/local/go/src/net/netip/netip.go:1347
		// _ = "end of CoverTab[4080]"
//line /usr/local/go/src/net/netip/netip.go:1347
	}() {
//line /usr/local/go/src/net/netip/netip.go:1347
		_go_fuzz_dep_.CoverTab[4081]++
								return false
//line /usr/local/go/src/net/netip/netip.go:1348
		// _ = "end of CoverTab[4081]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1349
		_go_fuzz_dep_.CoverTab[4082]++
//line /usr/local/go/src/net/netip/netip.go:1349
		// _ = "end of CoverTab[4082]"
//line /usr/local/go/src/net/netip/netip.go:1349
	}
//line /usr/local/go/src/net/netip/netip.go:1349
	// _ = "end of CoverTab[4074]"
//line /usr/local/go/src/net/netip/netip.go:1349
	_go_fuzz_dep_.CoverTab[4075]++
							if ip.Is4() {
//line /usr/local/go/src/net/netip/netip.go:1350
		_go_fuzz_dep_.CoverTab[4083]++

//line /usr/local/go/src/net/netip/netip.go:1359
		return uint32((ip.addr.lo^p.ip.addr.lo)>>((32-p.bits)&63)) == 0
//line /usr/local/go/src/net/netip/netip.go:1359
		// _ = "end of CoverTab[4083]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1360
		_go_fuzz_dep_.CoverTab[4084]++

//line /usr/local/go/src/net/netip/netip.go:1364
		return ip.addr.xor(p.ip.addr).and(mask6(int(p.bits))).isZero()
//line /usr/local/go/src/net/netip/netip.go:1364
		// _ = "end of CoverTab[4084]"
	}
//line /usr/local/go/src/net/netip/netip.go:1365
	// _ = "end of CoverTab[4075]"
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
	_go_fuzz_dep_.CoverTab[4085]++
							if !p.IsValid() || func() bool {
//line /usr/local/go/src/net/netip/netip.go:1374
		_go_fuzz_dep_.CoverTab[4093]++
//line /usr/local/go/src/net/netip/netip.go:1374
		return !o.IsValid()
//line /usr/local/go/src/net/netip/netip.go:1374
		// _ = "end of CoverTab[4093]"
//line /usr/local/go/src/net/netip/netip.go:1374
	}() {
//line /usr/local/go/src/net/netip/netip.go:1374
		_go_fuzz_dep_.CoverTab[4094]++
								return false
//line /usr/local/go/src/net/netip/netip.go:1375
		// _ = "end of CoverTab[4094]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1376
		_go_fuzz_dep_.CoverTab[4095]++
//line /usr/local/go/src/net/netip/netip.go:1376
		// _ = "end of CoverTab[4095]"
//line /usr/local/go/src/net/netip/netip.go:1376
	}
//line /usr/local/go/src/net/netip/netip.go:1376
	// _ = "end of CoverTab[4085]"
//line /usr/local/go/src/net/netip/netip.go:1376
	_go_fuzz_dep_.CoverTab[4086]++
							if p == o {
//line /usr/local/go/src/net/netip/netip.go:1377
		_go_fuzz_dep_.CoverTab[4096]++
								return true
//line /usr/local/go/src/net/netip/netip.go:1378
		// _ = "end of CoverTab[4096]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1379
		_go_fuzz_dep_.CoverTab[4097]++
//line /usr/local/go/src/net/netip/netip.go:1379
		// _ = "end of CoverTab[4097]"
//line /usr/local/go/src/net/netip/netip.go:1379
	}
//line /usr/local/go/src/net/netip/netip.go:1379
	// _ = "end of CoverTab[4086]"
//line /usr/local/go/src/net/netip/netip.go:1379
	_go_fuzz_dep_.CoverTab[4087]++
							if p.ip.Is4() != o.ip.Is4() {
//line /usr/local/go/src/net/netip/netip.go:1380
		_go_fuzz_dep_.CoverTab[4098]++
								return false
//line /usr/local/go/src/net/netip/netip.go:1381
		// _ = "end of CoverTab[4098]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1382
		_go_fuzz_dep_.CoverTab[4099]++
//line /usr/local/go/src/net/netip/netip.go:1382
		// _ = "end of CoverTab[4099]"
//line /usr/local/go/src/net/netip/netip.go:1382
	}
//line /usr/local/go/src/net/netip/netip.go:1382
	// _ = "end of CoverTab[4087]"
//line /usr/local/go/src/net/netip/netip.go:1382
	_go_fuzz_dep_.CoverTab[4088]++
							var minBits int16
							if p.bits < o.bits {
//line /usr/local/go/src/net/netip/netip.go:1384
		_go_fuzz_dep_.CoverTab[4100]++
								minBits = p.bits
//line /usr/local/go/src/net/netip/netip.go:1385
		// _ = "end of CoverTab[4100]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1386
		_go_fuzz_dep_.CoverTab[4101]++
								minBits = o.bits
//line /usr/local/go/src/net/netip/netip.go:1387
		// _ = "end of CoverTab[4101]"
	}
//line /usr/local/go/src/net/netip/netip.go:1388
	// _ = "end of CoverTab[4088]"
//line /usr/local/go/src/net/netip/netip.go:1388
	_go_fuzz_dep_.CoverTab[4089]++
							if minBits == 0 {
//line /usr/local/go/src/net/netip/netip.go:1389
		_go_fuzz_dep_.CoverTab[4102]++
								return true
//line /usr/local/go/src/net/netip/netip.go:1390
		// _ = "end of CoverTab[4102]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1391
		_go_fuzz_dep_.CoverTab[4103]++
//line /usr/local/go/src/net/netip/netip.go:1391
		// _ = "end of CoverTab[4103]"
//line /usr/local/go/src/net/netip/netip.go:1391
	}
//line /usr/local/go/src/net/netip/netip.go:1391
	// _ = "end of CoverTab[4089]"
//line /usr/local/go/src/net/netip/netip.go:1391
	_go_fuzz_dep_.CoverTab[4090]++
	// One of these Prefix calls might look redundant, but we don't require
	// that p and o values are normalized (via Prefix.Masked) first,
	// so the Prefix call on the one that's already minBits serves to zero
	// out any remaining bits in IP.
	var err error
	if p, err = p.ip.Prefix(int(minBits)); err != nil {
//line /usr/local/go/src/net/netip/netip.go:1397
		_go_fuzz_dep_.CoverTab[4104]++
								return false
//line /usr/local/go/src/net/netip/netip.go:1398
		// _ = "end of CoverTab[4104]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1399
		_go_fuzz_dep_.CoverTab[4105]++
//line /usr/local/go/src/net/netip/netip.go:1399
		// _ = "end of CoverTab[4105]"
//line /usr/local/go/src/net/netip/netip.go:1399
	}
//line /usr/local/go/src/net/netip/netip.go:1399
	// _ = "end of CoverTab[4090]"
//line /usr/local/go/src/net/netip/netip.go:1399
	_go_fuzz_dep_.CoverTab[4091]++
							if o, err = o.ip.Prefix(int(minBits)); err != nil {
//line /usr/local/go/src/net/netip/netip.go:1400
		_go_fuzz_dep_.CoverTab[4106]++
								return false
//line /usr/local/go/src/net/netip/netip.go:1401
		// _ = "end of CoverTab[4106]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1402
		_go_fuzz_dep_.CoverTab[4107]++
//line /usr/local/go/src/net/netip/netip.go:1402
		// _ = "end of CoverTab[4107]"
//line /usr/local/go/src/net/netip/netip.go:1402
	}
//line /usr/local/go/src/net/netip/netip.go:1402
	// _ = "end of CoverTab[4091]"
//line /usr/local/go/src/net/netip/netip.go:1402
	_go_fuzz_dep_.CoverTab[4092]++
							return p.ip == o.ip
//line /usr/local/go/src/net/netip/netip.go:1403
	// _ = "end of CoverTab[4092]"
}

// AppendTo appends a text encoding of p,
//line /usr/local/go/src/net/netip/netip.go:1406
// as generated by MarshalText,
//line /usr/local/go/src/net/netip/netip.go:1406
// to b and returns the extended buffer.
//line /usr/local/go/src/net/netip/netip.go:1409
func (p Prefix) AppendTo(b []byte) []byte {
//line /usr/local/go/src/net/netip/netip.go:1409
	_go_fuzz_dep_.CoverTab[4108]++
							if p.isZero() {
//line /usr/local/go/src/net/netip/netip.go:1410
		_go_fuzz_dep_.CoverTab[4112]++
								return b
//line /usr/local/go/src/net/netip/netip.go:1411
		// _ = "end of CoverTab[4112]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1412
		_go_fuzz_dep_.CoverTab[4113]++
//line /usr/local/go/src/net/netip/netip.go:1412
		// _ = "end of CoverTab[4113]"
//line /usr/local/go/src/net/netip/netip.go:1412
	}
//line /usr/local/go/src/net/netip/netip.go:1412
	// _ = "end of CoverTab[4108]"
//line /usr/local/go/src/net/netip/netip.go:1412
	_go_fuzz_dep_.CoverTab[4109]++
							if !p.IsValid() {
//line /usr/local/go/src/net/netip/netip.go:1413
		_go_fuzz_dep_.CoverTab[4114]++
								return append(b, "invalid Prefix"...)
//line /usr/local/go/src/net/netip/netip.go:1414
		// _ = "end of CoverTab[4114]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1415
		_go_fuzz_dep_.CoverTab[4115]++
//line /usr/local/go/src/net/netip/netip.go:1415
		// _ = "end of CoverTab[4115]"
//line /usr/local/go/src/net/netip/netip.go:1415
	}
//line /usr/local/go/src/net/netip/netip.go:1415
	// _ = "end of CoverTab[4109]"
//line /usr/local/go/src/net/netip/netip.go:1415
	_go_fuzz_dep_.CoverTab[4110]++

//line /usr/local/go/src/net/netip/netip.go:1418
	if p.ip.z == z4 {
//line /usr/local/go/src/net/netip/netip.go:1418
		_go_fuzz_dep_.CoverTab[4116]++
								b = p.ip.appendTo4(b)
//line /usr/local/go/src/net/netip/netip.go:1419
		// _ = "end of CoverTab[4116]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1420
		_go_fuzz_dep_.CoverTab[4117]++
								if p.ip.Is4In6() {
//line /usr/local/go/src/net/netip/netip.go:1421
			_go_fuzz_dep_.CoverTab[4118]++
									b = append(b, "::ffff:"...)
									b = p.ip.Unmap().appendTo4(b)
//line /usr/local/go/src/net/netip/netip.go:1423
			// _ = "end of CoverTab[4118]"
		} else {
//line /usr/local/go/src/net/netip/netip.go:1424
			_go_fuzz_dep_.CoverTab[4119]++
									b = p.ip.appendTo6(b)
//line /usr/local/go/src/net/netip/netip.go:1425
			// _ = "end of CoverTab[4119]"
		}
//line /usr/local/go/src/net/netip/netip.go:1426
		// _ = "end of CoverTab[4117]"
	}
//line /usr/local/go/src/net/netip/netip.go:1427
	// _ = "end of CoverTab[4110]"
//line /usr/local/go/src/net/netip/netip.go:1427
	_go_fuzz_dep_.CoverTab[4111]++

							b = append(b, '/')
							b = appendDecimal(b, uint8(p.bits))
							return b
//line /usr/local/go/src/net/netip/netip.go:1431
	// _ = "end of CoverTab[4111]"
}

// MarshalText implements the encoding.TextMarshaler interface,
//line /usr/local/go/src/net/netip/netip.go:1434
// The encoding is the same as returned by String, with one exception:
//line /usr/local/go/src/net/netip/netip.go:1434
// If p is the zero value, the encoding is the empty string.
//line /usr/local/go/src/net/netip/netip.go:1437
func (p Prefix) MarshalText() ([]byte, error) {
//line /usr/local/go/src/net/netip/netip.go:1437
	_go_fuzz_dep_.CoverTab[4120]++
							var max int
							switch p.ip.z {
	case z0:
//line /usr/local/go/src/net/netip/netip.go:1440
		_go_fuzz_dep_.CoverTab[4122]++
//line /usr/local/go/src/net/netip/netip.go:1440
		// _ = "end of CoverTab[4122]"
	case z4:
//line /usr/local/go/src/net/netip/netip.go:1441
		_go_fuzz_dep_.CoverTab[4123]++
								max = len("255.255.255.255/32")
//line /usr/local/go/src/net/netip/netip.go:1442
		// _ = "end of CoverTab[4123]"
	default:
//line /usr/local/go/src/net/netip/netip.go:1443
		_go_fuzz_dep_.CoverTab[4124]++
								max = len("ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff%enp5s0/128")
//line /usr/local/go/src/net/netip/netip.go:1444
		// _ = "end of CoverTab[4124]"
	}
//line /usr/local/go/src/net/netip/netip.go:1445
	// _ = "end of CoverTab[4120]"
//line /usr/local/go/src/net/netip/netip.go:1445
	_go_fuzz_dep_.CoverTab[4121]++
							b := make([]byte, 0, max)
							b = p.AppendTo(b)
							return b, nil
//line /usr/local/go/src/net/netip/netip.go:1448
	// _ = "end of CoverTab[4121]"
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
//line /usr/local/go/src/net/netip/netip.go:1451
// The IP address is expected in a form accepted by ParsePrefix
//line /usr/local/go/src/net/netip/netip.go:1451
// or generated by MarshalText.
//line /usr/local/go/src/net/netip/netip.go:1454
func (p *Prefix) UnmarshalText(text []byte) error {
//line /usr/local/go/src/net/netip/netip.go:1454
	_go_fuzz_dep_.CoverTab[4125]++
							if len(text) == 0 {
//line /usr/local/go/src/net/netip/netip.go:1455
		_go_fuzz_dep_.CoverTab[4127]++
								*p = Prefix{}
								return nil
//line /usr/local/go/src/net/netip/netip.go:1457
		// _ = "end of CoverTab[4127]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1458
		_go_fuzz_dep_.CoverTab[4128]++
//line /usr/local/go/src/net/netip/netip.go:1458
		// _ = "end of CoverTab[4128]"
//line /usr/local/go/src/net/netip/netip.go:1458
	}
//line /usr/local/go/src/net/netip/netip.go:1458
	// _ = "end of CoverTab[4125]"
//line /usr/local/go/src/net/netip/netip.go:1458
	_go_fuzz_dep_.CoverTab[4126]++
							var err error
							*p, err = ParsePrefix(string(text))
							return err
//line /usr/local/go/src/net/netip/netip.go:1461
	// _ = "end of CoverTab[4126]"
}

// MarshalBinary implements the encoding.BinaryMarshaler interface.
//line /usr/local/go/src/net/netip/netip.go:1464
// It returns Addr.MarshalBinary with an additional byte appended
//line /usr/local/go/src/net/netip/netip.go:1464
// containing the prefix bits.
//line /usr/local/go/src/net/netip/netip.go:1467
func (p Prefix) MarshalBinary() ([]byte, error) {
//line /usr/local/go/src/net/netip/netip.go:1467
	_go_fuzz_dep_.CoverTab[4129]++
							b := p.Addr().withoutZone().marshalBinaryWithTrailingBytes(1)
							b[len(b)-1] = uint8(p.Bits())
							return b, nil
//line /usr/local/go/src/net/netip/netip.go:1470
	// _ = "end of CoverTab[4129]"
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface.
//line /usr/local/go/src/net/netip/netip.go:1473
// It expects data in the form generated by MarshalBinary.
//line /usr/local/go/src/net/netip/netip.go:1475
func (p *Prefix) UnmarshalBinary(b []byte) error {
//line /usr/local/go/src/net/netip/netip.go:1475
	_go_fuzz_dep_.CoverTab[4130]++
							if len(b) < 1 {
//line /usr/local/go/src/net/netip/netip.go:1476
		_go_fuzz_dep_.CoverTab[4133]++
								return errors.New("unexpected slice size")
//line /usr/local/go/src/net/netip/netip.go:1477
		// _ = "end of CoverTab[4133]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1478
		_go_fuzz_dep_.CoverTab[4134]++
//line /usr/local/go/src/net/netip/netip.go:1478
		// _ = "end of CoverTab[4134]"
//line /usr/local/go/src/net/netip/netip.go:1478
	}
//line /usr/local/go/src/net/netip/netip.go:1478
	// _ = "end of CoverTab[4130]"
//line /usr/local/go/src/net/netip/netip.go:1478
	_go_fuzz_dep_.CoverTab[4131]++
							var addr Addr
							err := addr.UnmarshalBinary(b[:len(b)-1])
							if err != nil {
//line /usr/local/go/src/net/netip/netip.go:1481
		_go_fuzz_dep_.CoverTab[4135]++
								return err
//line /usr/local/go/src/net/netip/netip.go:1482
		// _ = "end of CoverTab[4135]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1483
		_go_fuzz_dep_.CoverTab[4136]++
//line /usr/local/go/src/net/netip/netip.go:1483
		// _ = "end of CoverTab[4136]"
//line /usr/local/go/src/net/netip/netip.go:1483
	}
//line /usr/local/go/src/net/netip/netip.go:1483
	// _ = "end of CoverTab[4131]"
//line /usr/local/go/src/net/netip/netip.go:1483
	_go_fuzz_dep_.CoverTab[4132]++
							*p = PrefixFrom(addr, int(b[len(b)-1]))
							return nil
//line /usr/local/go/src/net/netip/netip.go:1485
	// _ = "end of CoverTab[4132]"
}

// String returns the CIDR notation of p: "<ip>/<bits>".
func (p Prefix) String() string {
//line /usr/local/go/src/net/netip/netip.go:1489
	_go_fuzz_dep_.CoverTab[4137]++
							if !p.IsValid() {
//line /usr/local/go/src/net/netip/netip.go:1490
		_go_fuzz_dep_.CoverTab[4139]++
								return "invalid Prefix"
//line /usr/local/go/src/net/netip/netip.go:1491
		// _ = "end of CoverTab[4139]"
	} else {
//line /usr/local/go/src/net/netip/netip.go:1492
		_go_fuzz_dep_.CoverTab[4140]++
//line /usr/local/go/src/net/netip/netip.go:1492
		// _ = "end of CoverTab[4140]"
//line /usr/local/go/src/net/netip/netip.go:1492
	}
//line /usr/local/go/src/net/netip/netip.go:1492
	// _ = "end of CoverTab[4137]"
//line /usr/local/go/src/net/netip/netip.go:1492
	_go_fuzz_dep_.CoverTab[4138]++
							return p.ip.String() + "/" + itoa.Itoa(int(p.bits))
//line /usr/local/go/src/net/netip/netip.go:1493
	// _ = "end of CoverTab[4138]"
}

//line /usr/local/go/src/net/netip/netip.go:1494
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/netip/netip.go:1494
var _ = _go_fuzz_dep_.CoverTab
