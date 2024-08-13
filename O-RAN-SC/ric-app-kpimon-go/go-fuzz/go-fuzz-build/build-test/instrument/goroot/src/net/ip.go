// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// IP address manipulations
//
// IPv4 addresses are 4 bytes; IPv6 addresses are 16 bytes.
// An IPv4 address can be converted to an IPv6 address by
// adding a canonical prefix (10 zeros, 2 0xFFs).
// This library accepts either size of byte slice but always
// returns 16-byte addresses.

//line /snap/go/10455/src/net/ip.go:13
package net

//line /snap/go/10455/src/net/ip.go:13
import (
//line /snap/go/10455/src/net/ip.go:13
	_go_fuzz_dep_ "go-fuzz-dep"
//line /snap/go/10455/src/net/ip.go:13
)
//line /snap/go/10455/src/net/ip.go:13
import (
//line /snap/go/10455/src/net/ip.go:13
	_atomic_ "sync/atomic"
//line /snap/go/10455/src/net/ip.go:13
)

import (
	"internal/bytealg"
	"internal/itoa"
	"net/netip"
)

// IP address lengths (bytes).
const (
	IPv4len	= 4
	IPv6len	= 16
)

// An IP is a single IP address, a slice of bytes.
//line /snap/go/10455/src/net/ip.go:27
// Functions in this package accept either 4-byte (IPv4)
//line /snap/go/10455/src/net/ip.go:27
// or 16-byte (IPv6) slices as input.
//line /snap/go/10455/src/net/ip.go:27
//
//line /snap/go/10455/src/net/ip.go:27
// Note that in this documentation, referring to an
//line /snap/go/10455/src/net/ip.go:27
// IP address as an IPv4 address or an IPv6 address
//line /snap/go/10455/src/net/ip.go:27
// is a semantic property of the address, not just the
//line /snap/go/10455/src/net/ip.go:27
// length of the byte slice: a 16-byte slice can still
//line /snap/go/10455/src/net/ip.go:27
// be an IPv4 address.
//line /snap/go/10455/src/net/ip.go:36
type IP []byte

// An IPMask is a bitmask that can be used to manipulate
//line /snap/go/10455/src/net/ip.go:38
// IP addresses for IP addressing and routing.
//line /snap/go/10455/src/net/ip.go:38
//
//line /snap/go/10455/src/net/ip.go:38
// See type IPNet and func ParseCIDR for details.
//line /snap/go/10455/src/net/ip.go:42
type IPMask []byte

// An IPNet represents an IP network.
type IPNet struct {
	IP	IP	// network number
	Mask	IPMask	// network mask
}

// IPv4 returns the IP address (in 16-byte form) of the
//line /snap/go/10455/src/net/ip.go:50
// IPv4 address a.b.c.d.
//line /snap/go/10455/src/net/ip.go:52
func IPv4(a, b, c, d byte) IP {
//line /snap/go/10455/src/net/ip.go:52
	_go_fuzz_dep_.CoverTab[6287]++
					p := make(IP, IPv6len)
					copy(p, v4InV6Prefix)
					p[12] = a
					p[13] = b
					p[14] = c
					p[15] = d
					return p
//line /snap/go/10455/src/net/ip.go:59
	// _ = "end of CoverTab[6287]"
}

var v4InV6Prefix = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xff, 0xff}

// IPv4Mask returns the IP mask (in 4-byte form) of the
//line /snap/go/10455/src/net/ip.go:64
// IPv4 mask a.b.c.d.
//line /snap/go/10455/src/net/ip.go:66
func IPv4Mask(a, b, c, d byte) IPMask {
//line /snap/go/10455/src/net/ip.go:66
	_go_fuzz_dep_.CoverTab[6288]++
					p := make(IPMask, IPv4len)
					p[0] = a
					p[1] = b
					p[2] = c
					p[3] = d
					return p
//line /snap/go/10455/src/net/ip.go:72
	// _ = "end of CoverTab[6288]"
}

// CIDRMask returns an IPMask consisting of 'ones' 1 bits
//line /snap/go/10455/src/net/ip.go:75
// followed by 0s up to a total length of 'bits' bits.
//line /snap/go/10455/src/net/ip.go:75
// For a mask of this form, CIDRMask is the inverse of IPMask.Size.
//line /snap/go/10455/src/net/ip.go:78
func CIDRMask(ones, bits int) IPMask {
//line /snap/go/10455/src/net/ip.go:78
	_go_fuzz_dep_.CoverTab[6289]++
					if bits != 8*IPv4len && func() bool {
//line /snap/go/10455/src/net/ip.go:79
		_go_fuzz_dep_.CoverTab[6293]++
//line /snap/go/10455/src/net/ip.go:79
		return bits != 8*IPv6len
//line /snap/go/10455/src/net/ip.go:79
		// _ = "end of CoverTab[6293]"
//line /snap/go/10455/src/net/ip.go:79
	}() {
//line /snap/go/10455/src/net/ip.go:79
		_go_fuzz_dep_.CoverTab[528557]++
//line /snap/go/10455/src/net/ip.go:79
		_go_fuzz_dep_.CoverTab[6294]++
						return nil
//line /snap/go/10455/src/net/ip.go:80
		// _ = "end of CoverTab[6294]"
	} else {
//line /snap/go/10455/src/net/ip.go:81
		_go_fuzz_dep_.CoverTab[528558]++
//line /snap/go/10455/src/net/ip.go:81
		_go_fuzz_dep_.CoverTab[6295]++
//line /snap/go/10455/src/net/ip.go:81
		// _ = "end of CoverTab[6295]"
//line /snap/go/10455/src/net/ip.go:81
	}
//line /snap/go/10455/src/net/ip.go:81
	// _ = "end of CoverTab[6289]"
//line /snap/go/10455/src/net/ip.go:81
	_go_fuzz_dep_.CoverTab[6290]++
					if ones < 0 || func() bool {
//line /snap/go/10455/src/net/ip.go:82
		_go_fuzz_dep_.CoverTab[6296]++
//line /snap/go/10455/src/net/ip.go:82
		return ones > bits
//line /snap/go/10455/src/net/ip.go:82
		// _ = "end of CoverTab[6296]"
//line /snap/go/10455/src/net/ip.go:82
	}() {
//line /snap/go/10455/src/net/ip.go:82
		_go_fuzz_dep_.CoverTab[528559]++
//line /snap/go/10455/src/net/ip.go:82
		_go_fuzz_dep_.CoverTab[6297]++
						return nil
//line /snap/go/10455/src/net/ip.go:83
		// _ = "end of CoverTab[6297]"
	} else {
//line /snap/go/10455/src/net/ip.go:84
		_go_fuzz_dep_.CoverTab[528560]++
//line /snap/go/10455/src/net/ip.go:84
		_go_fuzz_dep_.CoverTab[6298]++
//line /snap/go/10455/src/net/ip.go:84
		// _ = "end of CoverTab[6298]"
//line /snap/go/10455/src/net/ip.go:84
	}
//line /snap/go/10455/src/net/ip.go:84
	// _ = "end of CoverTab[6290]"
//line /snap/go/10455/src/net/ip.go:84
	_go_fuzz_dep_.CoverTab[6291]++
					l := bits / 8
					m := make(IPMask, l)
					n := uint(ones)
//line /snap/go/10455/src/net/ip.go:87
	_go_fuzz_dep_.CoverTab[786688] = 0
					for i := 0; i < l; i++ {
//line /snap/go/10455/src/net/ip.go:88
		if _go_fuzz_dep_.CoverTab[786688] == 0 {
//line /snap/go/10455/src/net/ip.go:88
			_go_fuzz_dep_.CoverTab[528661]++
//line /snap/go/10455/src/net/ip.go:88
		} else {
//line /snap/go/10455/src/net/ip.go:88
			_go_fuzz_dep_.CoverTab[528662]++
//line /snap/go/10455/src/net/ip.go:88
		}
//line /snap/go/10455/src/net/ip.go:88
		_go_fuzz_dep_.CoverTab[786688] = 1
//line /snap/go/10455/src/net/ip.go:88
		_go_fuzz_dep_.CoverTab[6299]++
						if n >= 8 {
//line /snap/go/10455/src/net/ip.go:89
			_go_fuzz_dep_.CoverTab[528561]++
//line /snap/go/10455/src/net/ip.go:89
			_go_fuzz_dep_.CoverTab[6301]++
							m[i] = 0xff
							n -= 8
							continue
//line /snap/go/10455/src/net/ip.go:92
			// _ = "end of CoverTab[6301]"
		} else {
//line /snap/go/10455/src/net/ip.go:93
			_go_fuzz_dep_.CoverTab[528562]++
//line /snap/go/10455/src/net/ip.go:93
			_go_fuzz_dep_.CoverTab[6302]++
//line /snap/go/10455/src/net/ip.go:93
			// _ = "end of CoverTab[6302]"
//line /snap/go/10455/src/net/ip.go:93
		}
//line /snap/go/10455/src/net/ip.go:93
		// _ = "end of CoverTab[6299]"
//line /snap/go/10455/src/net/ip.go:93
		_go_fuzz_dep_.CoverTab[6300]++
						m[i] = ^byte(0xff >> n)
						n = 0
//line /snap/go/10455/src/net/ip.go:95
		// _ = "end of CoverTab[6300]"
	}
//line /snap/go/10455/src/net/ip.go:96
	if _go_fuzz_dep_.CoverTab[786688] == 0 {
//line /snap/go/10455/src/net/ip.go:96
		_go_fuzz_dep_.CoverTab[528663]++
//line /snap/go/10455/src/net/ip.go:96
	} else {
//line /snap/go/10455/src/net/ip.go:96
		_go_fuzz_dep_.CoverTab[528664]++
//line /snap/go/10455/src/net/ip.go:96
	}
//line /snap/go/10455/src/net/ip.go:96
	// _ = "end of CoverTab[6291]"
//line /snap/go/10455/src/net/ip.go:96
	_go_fuzz_dep_.CoverTab[6292]++
					return m
//line /snap/go/10455/src/net/ip.go:97
	// _ = "end of CoverTab[6292]"
}

// Well-known IPv4 addresses
var (
	IPv4bcast	= IPv4(255, 255, 255, 255)	// limited broadcast
	IPv4allsys	= IPv4(224, 0, 0, 1)		// all systems
	IPv4allrouter	= IPv4(224, 0, 0, 2)		// all routers
	IPv4zero	= IPv4(0, 0, 0, 0)		// all zeros
)

// Well-known IPv6 addresses
var (
	IPv6zero			= IP{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	IPv6unspecified			= IP{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	IPv6loopback			= IP{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	IPv6interfacelocalallnodes	= IP{0xff, 0x01, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x01}
	IPv6linklocalallnodes		= IP{0xff, 0x02, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x01}
	IPv6linklocalallrouters		= IP{0xff, 0x02, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x02}
)

// IsUnspecified reports whether ip is an unspecified address, either
//line /snap/go/10455/src/net/ip.go:118
// the IPv4 address "0.0.0.0" or the IPv6 address "::".
//line /snap/go/10455/src/net/ip.go:120
func (ip IP) IsUnspecified() bool {
//line /snap/go/10455/src/net/ip.go:120
	_go_fuzz_dep_.CoverTab[6303]++
						return ip.Equal(IPv4zero) || func() bool {
//line /snap/go/10455/src/net/ip.go:121
		_go_fuzz_dep_.CoverTab[6304]++
//line /snap/go/10455/src/net/ip.go:121
		return ip.Equal(IPv6unspecified)
//line /snap/go/10455/src/net/ip.go:121
		// _ = "end of CoverTab[6304]"
//line /snap/go/10455/src/net/ip.go:121
	}()
//line /snap/go/10455/src/net/ip.go:121
	// _ = "end of CoverTab[6303]"
}

// IsLoopback reports whether ip is a loopback address.
func (ip IP) IsLoopback() bool {
//line /snap/go/10455/src/net/ip.go:125
	_go_fuzz_dep_.CoverTab[6305]++
						if ip4 := ip.To4(); ip4 != nil {
//line /snap/go/10455/src/net/ip.go:126
		_go_fuzz_dep_.CoverTab[528563]++
//line /snap/go/10455/src/net/ip.go:126
		_go_fuzz_dep_.CoverTab[6307]++
							return ip4[0] == 127
//line /snap/go/10455/src/net/ip.go:127
		// _ = "end of CoverTab[6307]"
	} else {
//line /snap/go/10455/src/net/ip.go:128
		_go_fuzz_dep_.CoverTab[528564]++
//line /snap/go/10455/src/net/ip.go:128
		_go_fuzz_dep_.CoverTab[6308]++
//line /snap/go/10455/src/net/ip.go:128
		// _ = "end of CoverTab[6308]"
//line /snap/go/10455/src/net/ip.go:128
	}
//line /snap/go/10455/src/net/ip.go:128
	// _ = "end of CoverTab[6305]"
//line /snap/go/10455/src/net/ip.go:128
	_go_fuzz_dep_.CoverTab[6306]++
						return ip.Equal(IPv6loopback)
//line /snap/go/10455/src/net/ip.go:129
	// _ = "end of CoverTab[6306]"
}

// IsPrivate reports whether ip is a private address, according to
//line /snap/go/10455/src/net/ip.go:132
// RFC 1918 (IPv4 addresses) and RFC 4193 (IPv6 addresses).
//line /snap/go/10455/src/net/ip.go:134
func (ip IP) IsPrivate() bool {
//line /snap/go/10455/src/net/ip.go:134
	_go_fuzz_dep_.CoverTab[6309]++
						if ip4 := ip.To4(); ip4 != nil {
//line /snap/go/10455/src/net/ip.go:135
		_go_fuzz_dep_.CoverTab[528565]++
//line /snap/go/10455/src/net/ip.go:135
		_go_fuzz_dep_.CoverTab[6311]++

//line /snap/go/10455/src/net/ip.go:142
		return ip4[0] == 10 || func() bool {
//line /snap/go/10455/src/net/ip.go:142
			_go_fuzz_dep_.CoverTab[6312]++
//line /snap/go/10455/src/net/ip.go:142
			return (ip4[0] == 172 && func() bool {
									_go_fuzz_dep_.CoverTab[6313]++
//line /snap/go/10455/src/net/ip.go:143
				return ip4[1]&0xf0 == 16
//line /snap/go/10455/src/net/ip.go:143
				// _ = "end of CoverTab[6313]"
//line /snap/go/10455/src/net/ip.go:143
			}())
//line /snap/go/10455/src/net/ip.go:143
			// _ = "end of CoverTab[6312]"
//line /snap/go/10455/src/net/ip.go:143
		}() || func() bool {
//line /snap/go/10455/src/net/ip.go:143
			_go_fuzz_dep_.CoverTab[6314]++
//line /snap/go/10455/src/net/ip.go:143
			return (ip4[0] == 192 && func() bool {
									_go_fuzz_dep_.CoverTab[6315]++
//line /snap/go/10455/src/net/ip.go:144
				return ip4[1] == 168
//line /snap/go/10455/src/net/ip.go:144
				// _ = "end of CoverTab[6315]"
//line /snap/go/10455/src/net/ip.go:144
			}())
//line /snap/go/10455/src/net/ip.go:144
			// _ = "end of CoverTab[6314]"
//line /snap/go/10455/src/net/ip.go:144
		}()
//line /snap/go/10455/src/net/ip.go:144
		// _ = "end of CoverTab[6311]"
	} else {
//line /snap/go/10455/src/net/ip.go:145
		_go_fuzz_dep_.CoverTab[528566]++
//line /snap/go/10455/src/net/ip.go:145
		_go_fuzz_dep_.CoverTab[6316]++
//line /snap/go/10455/src/net/ip.go:145
		// _ = "end of CoverTab[6316]"
//line /snap/go/10455/src/net/ip.go:145
	}
//line /snap/go/10455/src/net/ip.go:145
	// _ = "end of CoverTab[6309]"
//line /snap/go/10455/src/net/ip.go:145
	_go_fuzz_dep_.CoverTab[6310]++

//line /snap/go/10455/src/net/ip.go:148
	return len(ip) == IPv6len && func() bool {
//line /snap/go/10455/src/net/ip.go:148
		_go_fuzz_dep_.CoverTab[6317]++
//line /snap/go/10455/src/net/ip.go:148
		return ip[0]&0xfe == 0xfc
//line /snap/go/10455/src/net/ip.go:148
		// _ = "end of CoverTab[6317]"
//line /snap/go/10455/src/net/ip.go:148
	}()
//line /snap/go/10455/src/net/ip.go:148
	// _ = "end of CoverTab[6310]"
}

// IsMulticast reports whether ip is a multicast address.
func (ip IP) IsMulticast() bool {
//line /snap/go/10455/src/net/ip.go:152
	_go_fuzz_dep_.CoverTab[6318]++
						if ip4 := ip.To4(); ip4 != nil {
//line /snap/go/10455/src/net/ip.go:153
		_go_fuzz_dep_.CoverTab[528567]++
//line /snap/go/10455/src/net/ip.go:153
		_go_fuzz_dep_.CoverTab[6320]++
							return ip4[0]&0xf0 == 0xe0
//line /snap/go/10455/src/net/ip.go:154
		// _ = "end of CoverTab[6320]"
	} else {
//line /snap/go/10455/src/net/ip.go:155
		_go_fuzz_dep_.CoverTab[528568]++
//line /snap/go/10455/src/net/ip.go:155
		_go_fuzz_dep_.CoverTab[6321]++
//line /snap/go/10455/src/net/ip.go:155
		// _ = "end of CoverTab[6321]"
//line /snap/go/10455/src/net/ip.go:155
	}
//line /snap/go/10455/src/net/ip.go:155
	// _ = "end of CoverTab[6318]"
//line /snap/go/10455/src/net/ip.go:155
	_go_fuzz_dep_.CoverTab[6319]++
						return len(ip) == IPv6len && func() bool {
//line /snap/go/10455/src/net/ip.go:156
		_go_fuzz_dep_.CoverTab[6322]++
//line /snap/go/10455/src/net/ip.go:156
		return ip[0] == 0xff
//line /snap/go/10455/src/net/ip.go:156
		// _ = "end of CoverTab[6322]"
//line /snap/go/10455/src/net/ip.go:156
	}()
//line /snap/go/10455/src/net/ip.go:156
	// _ = "end of CoverTab[6319]"
}

// IsInterfaceLocalMulticast reports whether ip is
//line /snap/go/10455/src/net/ip.go:159
// an interface-local multicast address.
//line /snap/go/10455/src/net/ip.go:161
func (ip IP) IsInterfaceLocalMulticast() bool {
//line /snap/go/10455/src/net/ip.go:161
	_go_fuzz_dep_.CoverTab[6323]++
						return len(ip) == IPv6len && func() bool {
//line /snap/go/10455/src/net/ip.go:162
		_go_fuzz_dep_.CoverTab[6324]++
//line /snap/go/10455/src/net/ip.go:162
		return ip[0] == 0xff
//line /snap/go/10455/src/net/ip.go:162
		// _ = "end of CoverTab[6324]"
//line /snap/go/10455/src/net/ip.go:162
	}() && func() bool {
//line /snap/go/10455/src/net/ip.go:162
		_go_fuzz_dep_.CoverTab[6325]++
//line /snap/go/10455/src/net/ip.go:162
		return ip[1]&0x0f == 0x01
//line /snap/go/10455/src/net/ip.go:162
		// _ = "end of CoverTab[6325]"
//line /snap/go/10455/src/net/ip.go:162
	}()
//line /snap/go/10455/src/net/ip.go:162
	// _ = "end of CoverTab[6323]"
}

// IsLinkLocalMulticast reports whether ip is a link-local
//line /snap/go/10455/src/net/ip.go:165
// multicast address.
//line /snap/go/10455/src/net/ip.go:167
func (ip IP) IsLinkLocalMulticast() bool {
//line /snap/go/10455/src/net/ip.go:167
	_go_fuzz_dep_.CoverTab[6326]++
						if ip4 := ip.To4(); ip4 != nil {
//line /snap/go/10455/src/net/ip.go:168
		_go_fuzz_dep_.CoverTab[528569]++
//line /snap/go/10455/src/net/ip.go:168
		_go_fuzz_dep_.CoverTab[6328]++
							return ip4[0] == 224 && func() bool {
//line /snap/go/10455/src/net/ip.go:169
			_go_fuzz_dep_.CoverTab[6329]++
//line /snap/go/10455/src/net/ip.go:169
			return ip4[1] == 0
//line /snap/go/10455/src/net/ip.go:169
			// _ = "end of CoverTab[6329]"
//line /snap/go/10455/src/net/ip.go:169
		}() && func() bool {
//line /snap/go/10455/src/net/ip.go:169
			_go_fuzz_dep_.CoverTab[6330]++
//line /snap/go/10455/src/net/ip.go:169
			return ip4[2] == 0
//line /snap/go/10455/src/net/ip.go:169
			// _ = "end of CoverTab[6330]"
//line /snap/go/10455/src/net/ip.go:169
		}()
//line /snap/go/10455/src/net/ip.go:169
		// _ = "end of CoverTab[6328]"
	} else {
//line /snap/go/10455/src/net/ip.go:170
		_go_fuzz_dep_.CoverTab[528570]++
//line /snap/go/10455/src/net/ip.go:170
		_go_fuzz_dep_.CoverTab[6331]++
//line /snap/go/10455/src/net/ip.go:170
		// _ = "end of CoverTab[6331]"
//line /snap/go/10455/src/net/ip.go:170
	}
//line /snap/go/10455/src/net/ip.go:170
	// _ = "end of CoverTab[6326]"
//line /snap/go/10455/src/net/ip.go:170
	_go_fuzz_dep_.CoverTab[6327]++
						return len(ip) == IPv6len && func() bool {
//line /snap/go/10455/src/net/ip.go:171
		_go_fuzz_dep_.CoverTab[6332]++
//line /snap/go/10455/src/net/ip.go:171
		return ip[0] == 0xff
//line /snap/go/10455/src/net/ip.go:171
		// _ = "end of CoverTab[6332]"
//line /snap/go/10455/src/net/ip.go:171
	}() && func() bool {
//line /snap/go/10455/src/net/ip.go:171
		_go_fuzz_dep_.CoverTab[6333]++
//line /snap/go/10455/src/net/ip.go:171
		return ip[1]&0x0f == 0x02
//line /snap/go/10455/src/net/ip.go:171
		// _ = "end of CoverTab[6333]"
//line /snap/go/10455/src/net/ip.go:171
	}()
//line /snap/go/10455/src/net/ip.go:171
	// _ = "end of CoverTab[6327]"
}

// IsLinkLocalUnicast reports whether ip is a link-local
//line /snap/go/10455/src/net/ip.go:174
// unicast address.
//line /snap/go/10455/src/net/ip.go:176
func (ip IP) IsLinkLocalUnicast() bool {
//line /snap/go/10455/src/net/ip.go:176
	_go_fuzz_dep_.CoverTab[6334]++
						if ip4 := ip.To4(); ip4 != nil {
//line /snap/go/10455/src/net/ip.go:177
		_go_fuzz_dep_.CoverTab[528571]++
//line /snap/go/10455/src/net/ip.go:177
		_go_fuzz_dep_.CoverTab[6336]++
							return ip4[0] == 169 && func() bool {
//line /snap/go/10455/src/net/ip.go:178
			_go_fuzz_dep_.CoverTab[6337]++
//line /snap/go/10455/src/net/ip.go:178
			return ip4[1] == 254
//line /snap/go/10455/src/net/ip.go:178
			// _ = "end of CoverTab[6337]"
//line /snap/go/10455/src/net/ip.go:178
		}()
//line /snap/go/10455/src/net/ip.go:178
		// _ = "end of CoverTab[6336]"
	} else {
//line /snap/go/10455/src/net/ip.go:179
		_go_fuzz_dep_.CoverTab[528572]++
//line /snap/go/10455/src/net/ip.go:179
		_go_fuzz_dep_.CoverTab[6338]++
//line /snap/go/10455/src/net/ip.go:179
		// _ = "end of CoverTab[6338]"
//line /snap/go/10455/src/net/ip.go:179
	}
//line /snap/go/10455/src/net/ip.go:179
	// _ = "end of CoverTab[6334]"
//line /snap/go/10455/src/net/ip.go:179
	_go_fuzz_dep_.CoverTab[6335]++
						return len(ip) == IPv6len && func() bool {
//line /snap/go/10455/src/net/ip.go:180
		_go_fuzz_dep_.CoverTab[6339]++
//line /snap/go/10455/src/net/ip.go:180
		return ip[0] == 0xfe
//line /snap/go/10455/src/net/ip.go:180
		// _ = "end of CoverTab[6339]"
//line /snap/go/10455/src/net/ip.go:180
	}() && func() bool {
//line /snap/go/10455/src/net/ip.go:180
		_go_fuzz_dep_.CoverTab[6340]++
//line /snap/go/10455/src/net/ip.go:180
		return ip[1]&0xc0 == 0x80
//line /snap/go/10455/src/net/ip.go:180
		// _ = "end of CoverTab[6340]"
//line /snap/go/10455/src/net/ip.go:180
	}()
//line /snap/go/10455/src/net/ip.go:180
	// _ = "end of CoverTab[6335]"
}

// IsGlobalUnicast reports whether ip is a global unicast
//line /snap/go/10455/src/net/ip.go:183
// address.
//line /snap/go/10455/src/net/ip.go:183
//
//line /snap/go/10455/src/net/ip.go:183
// The identification of global unicast addresses uses address type
//line /snap/go/10455/src/net/ip.go:183
// identification as defined in RFC 1122, RFC 4632 and RFC 4291 with
//line /snap/go/10455/src/net/ip.go:183
// the exception of IPv4 directed broadcast addresses.
//line /snap/go/10455/src/net/ip.go:183
// It returns true even if ip is in IPv4 private address space or
//line /snap/go/10455/src/net/ip.go:183
// local IPv6 unicast address space.
//line /snap/go/10455/src/net/ip.go:191
func (ip IP) IsGlobalUnicast() bool {
//line /snap/go/10455/src/net/ip.go:191
	_go_fuzz_dep_.CoverTab[6341]++
						return (len(ip) == IPv4len || func() bool {
//line /snap/go/10455/src/net/ip.go:192
		_go_fuzz_dep_.CoverTab[6342]++
//line /snap/go/10455/src/net/ip.go:192
		return len(ip) == IPv6len
//line /snap/go/10455/src/net/ip.go:192
		// _ = "end of CoverTab[6342]"
//line /snap/go/10455/src/net/ip.go:192
	}()) && func() bool {
//line /snap/go/10455/src/net/ip.go:192
		_go_fuzz_dep_.CoverTab[6343]++
//line /snap/go/10455/src/net/ip.go:192
		return !ip.Equal(IPv4bcast)
							// _ = "end of CoverTab[6343]"
//line /snap/go/10455/src/net/ip.go:193
	}() && func() bool {
//line /snap/go/10455/src/net/ip.go:193
		_go_fuzz_dep_.CoverTab[6344]++
//line /snap/go/10455/src/net/ip.go:193
		return !ip.IsUnspecified()
							// _ = "end of CoverTab[6344]"
//line /snap/go/10455/src/net/ip.go:194
	}() && func() bool {
//line /snap/go/10455/src/net/ip.go:194
		_go_fuzz_dep_.CoverTab[6345]++
//line /snap/go/10455/src/net/ip.go:194
		return !ip.IsLoopback()
							// _ = "end of CoverTab[6345]"
//line /snap/go/10455/src/net/ip.go:195
	}() && func() bool {
//line /snap/go/10455/src/net/ip.go:195
		_go_fuzz_dep_.CoverTab[6346]++
//line /snap/go/10455/src/net/ip.go:195
		return !ip.IsMulticast()
							// _ = "end of CoverTab[6346]"
//line /snap/go/10455/src/net/ip.go:196
	}() && func() bool {
//line /snap/go/10455/src/net/ip.go:196
		_go_fuzz_dep_.CoverTab[6347]++
//line /snap/go/10455/src/net/ip.go:196
		return !ip.IsLinkLocalUnicast()
							// _ = "end of CoverTab[6347]"
//line /snap/go/10455/src/net/ip.go:197
	}()
//line /snap/go/10455/src/net/ip.go:197
	// _ = "end of CoverTab[6341]"
}

// Is p all zeros?
func isZeros(p IP) bool {
//line /snap/go/10455/src/net/ip.go:201
	_go_fuzz_dep_.CoverTab[6348]++
//line /snap/go/10455/src/net/ip.go:201
	_go_fuzz_dep_.CoverTab[786689] = 0
						for i := 0; i < len(p); i++ {
//line /snap/go/10455/src/net/ip.go:202
		if _go_fuzz_dep_.CoverTab[786689] == 0 {
//line /snap/go/10455/src/net/ip.go:202
			_go_fuzz_dep_.CoverTab[528665]++
//line /snap/go/10455/src/net/ip.go:202
		} else {
//line /snap/go/10455/src/net/ip.go:202
			_go_fuzz_dep_.CoverTab[528666]++
//line /snap/go/10455/src/net/ip.go:202
		}
//line /snap/go/10455/src/net/ip.go:202
		_go_fuzz_dep_.CoverTab[786689] = 1
//line /snap/go/10455/src/net/ip.go:202
		_go_fuzz_dep_.CoverTab[6350]++
							if p[i] != 0 {
//line /snap/go/10455/src/net/ip.go:203
			_go_fuzz_dep_.CoverTab[528573]++
//line /snap/go/10455/src/net/ip.go:203
			_go_fuzz_dep_.CoverTab[6351]++
								return false
//line /snap/go/10455/src/net/ip.go:204
			// _ = "end of CoverTab[6351]"
		} else {
//line /snap/go/10455/src/net/ip.go:205
			_go_fuzz_dep_.CoverTab[528574]++
//line /snap/go/10455/src/net/ip.go:205
			_go_fuzz_dep_.CoverTab[6352]++
//line /snap/go/10455/src/net/ip.go:205
			// _ = "end of CoverTab[6352]"
//line /snap/go/10455/src/net/ip.go:205
		}
//line /snap/go/10455/src/net/ip.go:205
		// _ = "end of CoverTab[6350]"
	}
//line /snap/go/10455/src/net/ip.go:206
	if _go_fuzz_dep_.CoverTab[786689] == 0 {
//line /snap/go/10455/src/net/ip.go:206
		_go_fuzz_dep_.CoverTab[528667]++
//line /snap/go/10455/src/net/ip.go:206
	} else {
//line /snap/go/10455/src/net/ip.go:206
		_go_fuzz_dep_.CoverTab[528668]++
//line /snap/go/10455/src/net/ip.go:206
	}
//line /snap/go/10455/src/net/ip.go:206
	// _ = "end of CoverTab[6348]"
//line /snap/go/10455/src/net/ip.go:206
	_go_fuzz_dep_.CoverTab[6349]++
						return true
//line /snap/go/10455/src/net/ip.go:207
	// _ = "end of CoverTab[6349]"
}

// To4 converts the IPv4 address ip to a 4-byte representation.
//line /snap/go/10455/src/net/ip.go:210
// If ip is not an IPv4 address, To4 returns nil.
//line /snap/go/10455/src/net/ip.go:212
func (ip IP) To4() IP {
//line /snap/go/10455/src/net/ip.go:212
	_go_fuzz_dep_.CoverTab[6353]++
						if len(ip) == IPv4len {
//line /snap/go/10455/src/net/ip.go:213
		_go_fuzz_dep_.CoverTab[528575]++
//line /snap/go/10455/src/net/ip.go:213
		_go_fuzz_dep_.CoverTab[6356]++
							return ip
//line /snap/go/10455/src/net/ip.go:214
		// _ = "end of CoverTab[6356]"
	} else {
//line /snap/go/10455/src/net/ip.go:215
		_go_fuzz_dep_.CoverTab[528576]++
//line /snap/go/10455/src/net/ip.go:215
		_go_fuzz_dep_.CoverTab[6357]++
//line /snap/go/10455/src/net/ip.go:215
		// _ = "end of CoverTab[6357]"
//line /snap/go/10455/src/net/ip.go:215
	}
//line /snap/go/10455/src/net/ip.go:215
	// _ = "end of CoverTab[6353]"
//line /snap/go/10455/src/net/ip.go:215
	_go_fuzz_dep_.CoverTab[6354]++
						if len(ip) == IPv6len && func() bool {
//line /snap/go/10455/src/net/ip.go:216
		_go_fuzz_dep_.CoverTab[6358]++
//line /snap/go/10455/src/net/ip.go:216
		return isZeros(ip[0:10])
							// _ = "end of CoverTab[6358]"
//line /snap/go/10455/src/net/ip.go:217
	}() && func() bool {
//line /snap/go/10455/src/net/ip.go:217
		_go_fuzz_dep_.CoverTab[6359]++
//line /snap/go/10455/src/net/ip.go:217
		return ip[10] == 0xff
							// _ = "end of CoverTab[6359]"
//line /snap/go/10455/src/net/ip.go:218
	}() && func() bool {
//line /snap/go/10455/src/net/ip.go:218
		_go_fuzz_dep_.CoverTab[6360]++
//line /snap/go/10455/src/net/ip.go:218
		return ip[11] == 0xff
							// _ = "end of CoverTab[6360]"
//line /snap/go/10455/src/net/ip.go:219
	}() {
//line /snap/go/10455/src/net/ip.go:219
		_go_fuzz_dep_.CoverTab[528577]++
//line /snap/go/10455/src/net/ip.go:219
		_go_fuzz_dep_.CoverTab[6361]++
							return ip[12:16]
//line /snap/go/10455/src/net/ip.go:220
		// _ = "end of CoverTab[6361]"
	} else {
//line /snap/go/10455/src/net/ip.go:221
		_go_fuzz_dep_.CoverTab[528578]++
//line /snap/go/10455/src/net/ip.go:221
		_go_fuzz_dep_.CoverTab[6362]++
//line /snap/go/10455/src/net/ip.go:221
		// _ = "end of CoverTab[6362]"
//line /snap/go/10455/src/net/ip.go:221
	}
//line /snap/go/10455/src/net/ip.go:221
	// _ = "end of CoverTab[6354]"
//line /snap/go/10455/src/net/ip.go:221
	_go_fuzz_dep_.CoverTab[6355]++
						return nil
//line /snap/go/10455/src/net/ip.go:222
	// _ = "end of CoverTab[6355]"
}

// To16 converts the IP address ip to a 16-byte representation.
//line /snap/go/10455/src/net/ip.go:225
// If ip is not an IP address (it is the wrong length), To16 returns nil.
//line /snap/go/10455/src/net/ip.go:227
func (ip IP) To16() IP {
//line /snap/go/10455/src/net/ip.go:227
	_go_fuzz_dep_.CoverTab[6363]++
						if len(ip) == IPv4len {
//line /snap/go/10455/src/net/ip.go:228
		_go_fuzz_dep_.CoverTab[528579]++
//line /snap/go/10455/src/net/ip.go:228
		_go_fuzz_dep_.CoverTab[6366]++
							return IPv4(ip[0], ip[1], ip[2], ip[3])
//line /snap/go/10455/src/net/ip.go:229
		// _ = "end of CoverTab[6366]"
	} else {
//line /snap/go/10455/src/net/ip.go:230
		_go_fuzz_dep_.CoverTab[528580]++
//line /snap/go/10455/src/net/ip.go:230
		_go_fuzz_dep_.CoverTab[6367]++
//line /snap/go/10455/src/net/ip.go:230
		// _ = "end of CoverTab[6367]"
//line /snap/go/10455/src/net/ip.go:230
	}
//line /snap/go/10455/src/net/ip.go:230
	// _ = "end of CoverTab[6363]"
//line /snap/go/10455/src/net/ip.go:230
	_go_fuzz_dep_.CoverTab[6364]++
						if len(ip) == IPv6len {
//line /snap/go/10455/src/net/ip.go:231
		_go_fuzz_dep_.CoverTab[528581]++
//line /snap/go/10455/src/net/ip.go:231
		_go_fuzz_dep_.CoverTab[6368]++
							return ip
//line /snap/go/10455/src/net/ip.go:232
		// _ = "end of CoverTab[6368]"
	} else {
//line /snap/go/10455/src/net/ip.go:233
		_go_fuzz_dep_.CoverTab[528582]++
//line /snap/go/10455/src/net/ip.go:233
		_go_fuzz_dep_.CoverTab[6369]++
//line /snap/go/10455/src/net/ip.go:233
		// _ = "end of CoverTab[6369]"
//line /snap/go/10455/src/net/ip.go:233
	}
//line /snap/go/10455/src/net/ip.go:233
	// _ = "end of CoverTab[6364]"
//line /snap/go/10455/src/net/ip.go:233
	_go_fuzz_dep_.CoverTab[6365]++
						return nil
//line /snap/go/10455/src/net/ip.go:234
	// _ = "end of CoverTab[6365]"
}

// Default route masks for IPv4.
var (
	classAMask	= IPv4Mask(0xff, 0, 0, 0)
	classBMask	= IPv4Mask(0xff, 0xff, 0, 0)
	classCMask	= IPv4Mask(0xff, 0xff, 0xff, 0)
)

// DefaultMask returns the default IP mask for the IP address ip.
//line /snap/go/10455/src/net/ip.go:244
// Only IPv4 addresses have default masks; DefaultMask returns
//line /snap/go/10455/src/net/ip.go:244
// nil if ip is not a valid IPv4 address.
//line /snap/go/10455/src/net/ip.go:247
func (ip IP) DefaultMask() IPMask {
//line /snap/go/10455/src/net/ip.go:247
	_go_fuzz_dep_.CoverTab[6370]++
						if ip = ip.To4(); ip == nil {
//line /snap/go/10455/src/net/ip.go:248
		_go_fuzz_dep_.CoverTab[528583]++
//line /snap/go/10455/src/net/ip.go:248
		_go_fuzz_dep_.CoverTab[6372]++
							return nil
//line /snap/go/10455/src/net/ip.go:249
		// _ = "end of CoverTab[6372]"
	} else {
//line /snap/go/10455/src/net/ip.go:250
		_go_fuzz_dep_.CoverTab[528584]++
//line /snap/go/10455/src/net/ip.go:250
		_go_fuzz_dep_.CoverTab[6373]++
//line /snap/go/10455/src/net/ip.go:250
		// _ = "end of CoverTab[6373]"
//line /snap/go/10455/src/net/ip.go:250
	}
//line /snap/go/10455/src/net/ip.go:250
	// _ = "end of CoverTab[6370]"
//line /snap/go/10455/src/net/ip.go:250
	_go_fuzz_dep_.CoverTab[6371]++
						switch {
	case ip[0] < 0x80:
//line /snap/go/10455/src/net/ip.go:252
		_go_fuzz_dep_.CoverTab[528585]++
//line /snap/go/10455/src/net/ip.go:252
		_go_fuzz_dep_.CoverTab[6374]++
							return classAMask
//line /snap/go/10455/src/net/ip.go:253
		// _ = "end of CoverTab[6374]"
	case ip[0] < 0xC0:
//line /snap/go/10455/src/net/ip.go:254
		_go_fuzz_dep_.CoverTab[528586]++
//line /snap/go/10455/src/net/ip.go:254
		_go_fuzz_dep_.CoverTab[6375]++
							return classBMask
//line /snap/go/10455/src/net/ip.go:255
		// _ = "end of CoverTab[6375]"
	default:
//line /snap/go/10455/src/net/ip.go:256
		_go_fuzz_dep_.CoverTab[528587]++
//line /snap/go/10455/src/net/ip.go:256
		_go_fuzz_dep_.CoverTab[6376]++
							return classCMask
//line /snap/go/10455/src/net/ip.go:257
		// _ = "end of CoverTab[6376]"
	}
//line /snap/go/10455/src/net/ip.go:258
	// _ = "end of CoverTab[6371]"
}

func allFF(b []byte) bool {
//line /snap/go/10455/src/net/ip.go:261
	_go_fuzz_dep_.CoverTab[6377]++
//line /snap/go/10455/src/net/ip.go:261
	_go_fuzz_dep_.CoverTab[786690] = 0
						for _, c := range b {
//line /snap/go/10455/src/net/ip.go:262
		if _go_fuzz_dep_.CoverTab[786690] == 0 {
//line /snap/go/10455/src/net/ip.go:262
			_go_fuzz_dep_.CoverTab[528669]++
//line /snap/go/10455/src/net/ip.go:262
		} else {
//line /snap/go/10455/src/net/ip.go:262
			_go_fuzz_dep_.CoverTab[528670]++
//line /snap/go/10455/src/net/ip.go:262
		}
//line /snap/go/10455/src/net/ip.go:262
		_go_fuzz_dep_.CoverTab[786690] = 1
//line /snap/go/10455/src/net/ip.go:262
		_go_fuzz_dep_.CoverTab[6379]++
							if c != 0xff {
//line /snap/go/10455/src/net/ip.go:263
			_go_fuzz_dep_.CoverTab[528588]++
//line /snap/go/10455/src/net/ip.go:263
			_go_fuzz_dep_.CoverTab[6380]++
								return false
//line /snap/go/10455/src/net/ip.go:264
			// _ = "end of CoverTab[6380]"
		} else {
//line /snap/go/10455/src/net/ip.go:265
			_go_fuzz_dep_.CoverTab[528589]++
//line /snap/go/10455/src/net/ip.go:265
			_go_fuzz_dep_.CoverTab[6381]++
//line /snap/go/10455/src/net/ip.go:265
			// _ = "end of CoverTab[6381]"
//line /snap/go/10455/src/net/ip.go:265
		}
//line /snap/go/10455/src/net/ip.go:265
		// _ = "end of CoverTab[6379]"
	}
//line /snap/go/10455/src/net/ip.go:266
	if _go_fuzz_dep_.CoverTab[786690] == 0 {
//line /snap/go/10455/src/net/ip.go:266
		_go_fuzz_dep_.CoverTab[528671]++
//line /snap/go/10455/src/net/ip.go:266
	} else {
//line /snap/go/10455/src/net/ip.go:266
		_go_fuzz_dep_.CoverTab[528672]++
//line /snap/go/10455/src/net/ip.go:266
	}
//line /snap/go/10455/src/net/ip.go:266
	// _ = "end of CoverTab[6377]"
//line /snap/go/10455/src/net/ip.go:266
	_go_fuzz_dep_.CoverTab[6378]++
						return true
//line /snap/go/10455/src/net/ip.go:267
	// _ = "end of CoverTab[6378]"
}

// Mask returns the result of masking the IP address ip with mask.
func (ip IP) Mask(mask IPMask) IP {
//line /snap/go/10455/src/net/ip.go:271
	_go_fuzz_dep_.CoverTab[6382]++
						if len(mask) == IPv6len && func() bool {
//line /snap/go/10455/src/net/ip.go:272
		_go_fuzz_dep_.CoverTab[6387]++
//line /snap/go/10455/src/net/ip.go:272
		return len(ip) == IPv4len
//line /snap/go/10455/src/net/ip.go:272
		// _ = "end of CoverTab[6387]"
//line /snap/go/10455/src/net/ip.go:272
	}() && func() bool {
//line /snap/go/10455/src/net/ip.go:272
		_go_fuzz_dep_.CoverTab[6388]++
//line /snap/go/10455/src/net/ip.go:272
		return allFF(mask[:12])
//line /snap/go/10455/src/net/ip.go:272
		// _ = "end of CoverTab[6388]"
//line /snap/go/10455/src/net/ip.go:272
	}() {
//line /snap/go/10455/src/net/ip.go:272
		_go_fuzz_dep_.CoverTab[528590]++
//line /snap/go/10455/src/net/ip.go:272
		_go_fuzz_dep_.CoverTab[6389]++
							mask = mask[12:]
//line /snap/go/10455/src/net/ip.go:273
		// _ = "end of CoverTab[6389]"
	} else {
//line /snap/go/10455/src/net/ip.go:274
		_go_fuzz_dep_.CoverTab[528591]++
//line /snap/go/10455/src/net/ip.go:274
		_go_fuzz_dep_.CoverTab[6390]++
//line /snap/go/10455/src/net/ip.go:274
		// _ = "end of CoverTab[6390]"
//line /snap/go/10455/src/net/ip.go:274
	}
//line /snap/go/10455/src/net/ip.go:274
	// _ = "end of CoverTab[6382]"
//line /snap/go/10455/src/net/ip.go:274
	_go_fuzz_dep_.CoverTab[6383]++
						if len(mask) == IPv4len && func() bool {
//line /snap/go/10455/src/net/ip.go:275
		_go_fuzz_dep_.CoverTab[6391]++
//line /snap/go/10455/src/net/ip.go:275
		return len(ip) == IPv6len
//line /snap/go/10455/src/net/ip.go:275
		// _ = "end of CoverTab[6391]"
//line /snap/go/10455/src/net/ip.go:275
	}() && func() bool {
//line /snap/go/10455/src/net/ip.go:275
		_go_fuzz_dep_.CoverTab[6392]++
//line /snap/go/10455/src/net/ip.go:275
		return bytealg.Equal(ip[:12], v4InV6Prefix)
//line /snap/go/10455/src/net/ip.go:275
		// _ = "end of CoverTab[6392]"
//line /snap/go/10455/src/net/ip.go:275
	}() {
//line /snap/go/10455/src/net/ip.go:275
		_go_fuzz_dep_.CoverTab[528592]++
//line /snap/go/10455/src/net/ip.go:275
		_go_fuzz_dep_.CoverTab[6393]++
							ip = ip[12:]
//line /snap/go/10455/src/net/ip.go:276
		// _ = "end of CoverTab[6393]"
	} else {
//line /snap/go/10455/src/net/ip.go:277
		_go_fuzz_dep_.CoverTab[528593]++
//line /snap/go/10455/src/net/ip.go:277
		_go_fuzz_dep_.CoverTab[6394]++
//line /snap/go/10455/src/net/ip.go:277
		// _ = "end of CoverTab[6394]"
//line /snap/go/10455/src/net/ip.go:277
	}
//line /snap/go/10455/src/net/ip.go:277
	// _ = "end of CoverTab[6383]"
//line /snap/go/10455/src/net/ip.go:277
	_go_fuzz_dep_.CoverTab[6384]++
						n := len(ip)
						if n != len(mask) {
//line /snap/go/10455/src/net/ip.go:279
		_go_fuzz_dep_.CoverTab[528594]++
//line /snap/go/10455/src/net/ip.go:279
		_go_fuzz_dep_.CoverTab[6395]++
							return nil
//line /snap/go/10455/src/net/ip.go:280
		// _ = "end of CoverTab[6395]"
	} else {
//line /snap/go/10455/src/net/ip.go:281
		_go_fuzz_dep_.CoverTab[528595]++
//line /snap/go/10455/src/net/ip.go:281
		_go_fuzz_dep_.CoverTab[6396]++
//line /snap/go/10455/src/net/ip.go:281
		// _ = "end of CoverTab[6396]"
//line /snap/go/10455/src/net/ip.go:281
	}
//line /snap/go/10455/src/net/ip.go:281
	// _ = "end of CoverTab[6384]"
//line /snap/go/10455/src/net/ip.go:281
	_go_fuzz_dep_.CoverTab[6385]++
						out := make(IP, n)
//line /snap/go/10455/src/net/ip.go:282
	_go_fuzz_dep_.CoverTab[786691] = 0
						for i := 0; i < n; i++ {
//line /snap/go/10455/src/net/ip.go:283
		if _go_fuzz_dep_.CoverTab[786691] == 0 {
//line /snap/go/10455/src/net/ip.go:283
			_go_fuzz_dep_.CoverTab[528673]++
//line /snap/go/10455/src/net/ip.go:283
		} else {
//line /snap/go/10455/src/net/ip.go:283
			_go_fuzz_dep_.CoverTab[528674]++
//line /snap/go/10455/src/net/ip.go:283
		}
//line /snap/go/10455/src/net/ip.go:283
		_go_fuzz_dep_.CoverTab[786691] = 1
//line /snap/go/10455/src/net/ip.go:283
		_go_fuzz_dep_.CoverTab[6397]++
							out[i] = ip[i] & mask[i]
//line /snap/go/10455/src/net/ip.go:284
		// _ = "end of CoverTab[6397]"
	}
//line /snap/go/10455/src/net/ip.go:285
	if _go_fuzz_dep_.CoverTab[786691] == 0 {
//line /snap/go/10455/src/net/ip.go:285
		_go_fuzz_dep_.CoverTab[528675]++
//line /snap/go/10455/src/net/ip.go:285
	} else {
//line /snap/go/10455/src/net/ip.go:285
		_go_fuzz_dep_.CoverTab[528676]++
//line /snap/go/10455/src/net/ip.go:285
	}
//line /snap/go/10455/src/net/ip.go:285
	// _ = "end of CoverTab[6385]"
//line /snap/go/10455/src/net/ip.go:285
	_go_fuzz_dep_.CoverTab[6386]++
						return out
//line /snap/go/10455/src/net/ip.go:286
	// _ = "end of CoverTab[6386]"
}

// String returns the string form of the IP address ip.
//line /snap/go/10455/src/net/ip.go:289
// It returns one of 4 forms:
//line /snap/go/10455/src/net/ip.go:289
//   - "<nil>", if ip has length 0
//line /snap/go/10455/src/net/ip.go:289
//   - dotted decimal ("192.0.2.1"), if ip is an IPv4 or IP4-mapped IPv6 address
//line /snap/go/10455/src/net/ip.go:289
//   - IPv6 conforming to RFC 5952 ("2001:db8::1"), if ip is a valid IPv6 address
//line /snap/go/10455/src/net/ip.go:289
//   - the hexadecimal form of ip, without punctuation, if no other cases apply
//line /snap/go/10455/src/net/ip.go:295
func (ip IP) String() string {
//line /snap/go/10455/src/net/ip.go:295
	_go_fuzz_dep_.CoverTab[6398]++
						if len(ip) == 0 {
//line /snap/go/10455/src/net/ip.go:296
		_go_fuzz_dep_.CoverTab[528596]++
//line /snap/go/10455/src/net/ip.go:296
		_go_fuzz_dep_.CoverTab[6402]++
							return "<nil>"
//line /snap/go/10455/src/net/ip.go:297
		// _ = "end of CoverTab[6402]"
	} else {
//line /snap/go/10455/src/net/ip.go:298
		_go_fuzz_dep_.CoverTab[528597]++
//line /snap/go/10455/src/net/ip.go:298
		_go_fuzz_dep_.CoverTab[6403]++
//line /snap/go/10455/src/net/ip.go:298
		// _ = "end of CoverTab[6403]"
//line /snap/go/10455/src/net/ip.go:298
	}
//line /snap/go/10455/src/net/ip.go:298
	// _ = "end of CoverTab[6398]"
//line /snap/go/10455/src/net/ip.go:298
	_go_fuzz_dep_.CoverTab[6399]++

						if len(ip) != IPv4len && func() bool {
//line /snap/go/10455/src/net/ip.go:300
		_go_fuzz_dep_.CoverTab[6404]++
//line /snap/go/10455/src/net/ip.go:300
		return len(ip) != IPv6len
//line /snap/go/10455/src/net/ip.go:300
		// _ = "end of CoverTab[6404]"
//line /snap/go/10455/src/net/ip.go:300
	}() {
//line /snap/go/10455/src/net/ip.go:300
		_go_fuzz_dep_.CoverTab[528598]++
//line /snap/go/10455/src/net/ip.go:300
		_go_fuzz_dep_.CoverTab[6405]++
							return "?" + hexString(ip)
//line /snap/go/10455/src/net/ip.go:301
		// _ = "end of CoverTab[6405]"
	} else {
//line /snap/go/10455/src/net/ip.go:302
		_go_fuzz_dep_.CoverTab[528599]++
//line /snap/go/10455/src/net/ip.go:302
		_go_fuzz_dep_.CoverTab[6406]++
//line /snap/go/10455/src/net/ip.go:302
		// _ = "end of CoverTab[6406]"
//line /snap/go/10455/src/net/ip.go:302
	}
//line /snap/go/10455/src/net/ip.go:302
	// _ = "end of CoverTab[6399]"
//line /snap/go/10455/src/net/ip.go:302
	_go_fuzz_dep_.CoverTab[6400]++

						if p4 := ip.To4(); len(p4) == IPv4len {
//line /snap/go/10455/src/net/ip.go:304
		_go_fuzz_dep_.CoverTab[528600]++
//line /snap/go/10455/src/net/ip.go:304
		_go_fuzz_dep_.CoverTab[6407]++
							return netip.AddrFrom4([4]byte(p4)).String()
//line /snap/go/10455/src/net/ip.go:305
		// _ = "end of CoverTab[6407]"
	} else {
//line /snap/go/10455/src/net/ip.go:306
		_go_fuzz_dep_.CoverTab[528601]++
//line /snap/go/10455/src/net/ip.go:306
		_go_fuzz_dep_.CoverTab[6408]++
//line /snap/go/10455/src/net/ip.go:306
		// _ = "end of CoverTab[6408]"
//line /snap/go/10455/src/net/ip.go:306
	}
//line /snap/go/10455/src/net/ip.go:306
	// _ = "end of CoverTab[6400]"
//line /snap/go/10455/src/net/ip.go:306
	_go_fuzz_dep_.CoverTab[6401]++
						return netip.AddrFrom16([16]byte(ip)).String()
//line /snap/go/10455/src/net/ip.go:307
	// _ = "end of CoverTab[6401]"
}

func hexString(b []byte) string {
//line /snap/go/10455/src/net/ip.go:310
	_go_fuzz_dep_.CoverTab[6409]++
						s := make([]byte, len(b)*2)
//line /snap/go/10455/src/net/ip.go:311
	_go_fuzz_dep_.CoverTab[786692] = 0
						for i, tn := range b {
//line /snap/go/10455/src/net/ip.go:312
		if _go_fuzz_dep_.CoverTab[786692] == 0 {
//line /snap/go/10455/src/net/ip.go:312
			_go_fuzz_dep_.CoverTab[528677]++
//line /snap/go/10455/src/net/ip.go:312
		} else {
//line /snap/go/10455/src/net/ip.go:312
			_go_fuzz_dep_.CoverTab[528678]++
//line /snap/go/10455/src/net/ip.go:312
		}
//line /snap/go/10455/src/net/ip.go:312
		_go_fuzz_dep_.CoverTab[786692] = 1
//line /snap/go/10455/src/net/ip.go:312
		_go_fuzz_dep_.CoverTab[6411]++
							s[i*2], s[i*2+1] = hexDigit[tn>>4], hexDigit[tn&0xf]
//line /snap/go/10455/src/net/ip.go:313
		// _ = "end of CoverTab[6411]"
	}
//line /snap/go/10455/src/net/ip.go:314
	if _go_fuzz_dep_.CoverTab[786692] == 0 {
//line /snap/go/10455/src/net/ip.go:314
		_go_fuzz_dep_.CoverTab[528679]++
//line /snap/go/10455/src/net/ip.go:314
	} else {
//line /snap/go/10455/src/net/ip.go:314
		_go_fuzz_dep_.CoverTab[528680]++
//line /snap/go/10455/src/net/ip.go:314
	}
//line /snap/go/10455/src/net/ip.go:314
	// _ = "end of CoverTab[6409]"
//line /snap/go/10455/src/net/ip.go:314
	_go_fuzz_dep_.CoverTab[6410]++
						return string(s)
//line /snap/go/10455/src/net/ip.go:315
	// _ = "end of CoverTab[6410]"
}

// ipEmptyString is like ip.String except that it returns
//line /snap/go/10455/src/net/ip.go:318
// an empty string when ip is unset.
//line /snap/go/10455/src/net/ip.go:320
func ipEmptyString(ip IP) string {
//line /snap/go/10455/src/net/ip.go:320
	_go_fuzz_dep_.CoverTab[6412]++
						if len(ip) == 0 {
//line /snap/go/10455/src/net/ip.go:321
		_go_fuzz_dep_.CoverTab[528602]++
//line /snap/go/10455/src/net/ip.go:321
		_go_fuzz_dep_.CoverTab[6414]++
							return ""
//line /snap/go/10455/src/net/ip.go:322
		// _ = "end of CoverTab[6414]"
	} else {
//line /snap/go/10455/src/net/ip.go:323
		_go_fuzz_dep_.CoverTab[528603]++
//line /snap/go/10455/src/net/ip.go:323
		_go_fuzz_dep_.CoverTab[6415]++
//line /snap/go/10455/src/net/ip.go:323
		// _ = "end of CoverTab[6415]"
//line /snap/go/10455/src/net/ip.go:323
	}
//line /snap/go/10455/src/net/ip.go:323
	// _ = "end of CoverTab[6412]"
//line /snap/go/10455/src/net/ip.go:323
	_go_fuzz_dep_.CoverTab[6413]++
						return ip.String()
//line /snap/go/10455/src/net/ip.go:324
	// _ = "end of CoverTab[6413]"
}

// MarshalText implements the encoding.TextMarshaler interface.
//line /snap/go/10455/src/net/ip.go:327
// The encoding is the same as returned by String, with one exception:
//line /snap/go/10455/src/net/ip.go:327
// When len(ip) is zero, it returns an empty slice.
//line /snap/go/10455/src/net/ip.go:330
func (ip IP) MarshalText() ([]byte, error) {
//line /snap/go/10455/src/net/ip.go:330
	_go_fuzz_dep_.CoverTab[6416]++
						if len(ip) == 0 {
//line /snap/go/10455/src/net/ip.go:331
		_go_fuzz_dep_.CoverTab[528604]++
//line /snap/go/10455/src/net/ip.go:331
		_go_fuzz_dep_.CoverTab[6419]++
							return []byte(""), nil
//line /snap/go/10455/src/net/ip.go:332
		// _ = "end of CoverTab[6419]"
	} else {
//line /snap/go/10455/src/net/ip.go:333
		_go_fuzz_dep_.CoverTab[528605]++
//line /snap/go/10455/src/net/ip.go:333
		_go_fuzz_dep_.CoverTab[6420]++
//line /snap/go/10455/src/net/ip.go:333
		// _ = "end of CoverTab[6420]"
//line /snap/go/10455/src/net/ip.go:333
	}
//line /snap/go/10455/src/net/ip.go:333
	// _ = "end of CoverTab[6416]"
//line /snap/go/10455/src/net/ip.go:333
	_go_fuzz_dep_.CoverTab[6417]++
						if len(ip) != IPv4len && func() bool {
//line /snap/go/10455/src/net/ip.go:334
		_go_fuzz_dep_.CoverTab[6421]++
//line /snap/go/10455/src/net/ip.go:334
		return len(ip) != IPv6len
//line /snap/go/10455/src/net/ip.go:334
		// _ = "end of CoverTab[6421]"
//line /snap/go/10455/src/net/ip.go:334
	}() {
//line /snap/go/10455/src/net/ip.go:334
		_go_fuzz_dep_.CoverTab[528606]++
//line /snap/go/10455/src/net/ip.go:334
		_go_fuzz_dep_.CoverTab[6422]++
							return nil, &AddrError{Err: "invalid IP address", Addr: hexString(ip)}
//line /snap/go/10455/src/net/ip.go:335
		// _ = "end of CoverTab[6422]"
	} else {
//line /snap/go/10455/src/net/ip.go:336
		_go_fuzz_dep_.CoverTab[528607]++
//line /snap/go/10455/src/net/ip.go:336
		_go_fuzz_dep_.CoverTab[6423]++
//line /snap/go/10455/src/net/ip.go:336
		// _ = "end of CoverTab[6423]"
//line /snap/go/10455/src/net/ip.go:336
	}
//line /snap/go/10455/src/net/ip.go:336
	// _ = "end of CoverTab[6417]"
//line /snap/go/10455/src/net/ip.go:336
	_go_fuzz_dep_.CoverTab[6418]++
						return []byte(ip.String()), nil
//line /snap/go/10455/src/net/ip.go:337
	// _ = "end of CoverTab[6418]"
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
//line /snap/go/10455/src/net/ip.go:340
// The IP address is expected in a form accepted by ParseIP.
//line /snap/go/10455/src/net/ip.go:342
func (ip *IP) UnmarshalText(text []byte) error {
//line /snap/go/10455/src/net/ip.go:342
	_go_fuzz_dep_.CoverTab[6424]++
						if len(text) == 0 {
//line /snap/go/10455/src/net/ip.go:343
		_go_fuzz_dep_.CoverTab[528608]++
//line /snap/go/10455/src/net/ip.go:343
		_go_fuzz_dep_.CoverTab[6427]++
							*ip = nil
							return nil
//line /snap/go/10455/src/net/ip.go:345
		// _ = "end of CoverTab[6427]"
	} else {
//line /snap/go/10455/src/net/ip.go:346
		_go_fuzz_dep_.CoverTab[528609]++
//line /snap/go/10455/src/net/ip.go:346
		_go_fuzz_dep_.CoverTab[6428]++
//line /snap/go/10455/src/net/ip.go:346
		// _ = "end of CoverTab[6428]"
//line /snap/go/10455/src/net/ip.go:346
	}
//line /snap/go/10455/src/net/ip.go:346
	// _ = "end of CoverTab[6424]"
//line /snap/go/10455/src/net/ip.go:346
	_go_fuzz_dep_.CoverTab[6425]++
						s := string(text)
						x := ParseIP(s)
						if x == nil {
//line /snap/go/10455/src/net/ip.go:349
		_go_fuzz_dep_.CoverTab[528610]++
//line /snap/go/10455/src/net/ip.go:349
		_go_fuzz_dep_.CoverTab[6429]++
							return &ParseError{Type: "IP address", Text: s}
//line /snap/go/10455/src/net/ip.go:350
		// _ = "end of CoverTab[6429]"
	} else {
//line /snap/go/10455/src/net/ip.go:351
		_go_fuzz_dep_.CoverTab[528611]++
//line /snap/go/10455/src/net/ip.go:351
		_go_fuzz_dep_.CoverTab[6430]++
//line /snap/go/10455/src/net/ip.go:351
		// _ = "end of CoverTab[6430]"
//line /snap/go/10455/src/net/ip.go:351
	}
//line /snap/go/10455/src/net/ip.go:351
	// _ = "end of CoverTab[6425]"
//line /snap/go/10455/src/net/ip.go:351
	_go_fuzz_dep_.CoverTab[6426]++
						*ip = x
						return nil
//line /snap/go/10455/src/net/ip.go:353
	// _ = "end of CoverTab[6426]"
}

// Equal reports whether ip and x are the same IP address.
//line /snap/go/10455/src/net/ip.go:356
// An IPv4 address and that same address in IPv6 form are
//line /snap/go/10455/src/net/ip.go:356
// considered to be equal.
//line /snap/go/10455/src/net/ip.go:359
func (ip IP) Equal(x IP) bool {
//line /snap/go/10455/src/net/ip.go:359
	_go_fuzz_dep_.CoverTab[6431]++
						if len(ip) == len(x) {
//line /snap/go/10455/src/net/ip.go:360
		_go_fuzz_dep_.CoverTab[528612]++
//line /snap/go/10455/src/net/ip.go:360
		_go_fuzz_dep_.CoverTab[6435]++
							return bytealg.Equal(ip, x)
//line /snap/go/10455/src/net/ip.go:361
		// _ = "end of CoverTab[6435]"
	} else {
//line /snap/go/10455/src/net/ip.go:362
		_go_fuzz_dep_.CoverTab[528613]++
//line /snap/go/10455/src/net/ip.go:362
		_go_fuzz_dep_.CoverTab[6436]++
//line /snap/go/10455/src/net/ip.go:362
		// _ = "end of CoverTab[6436]"
//line /snap/go/10455/src/net/ip.go:362
	}
//line /snap/go/10455/src/net/ip.go:362
	// _ = "end of CoverTab[6431]"
//line /snap/go/10455/src/net/ip.go:362
	_go_fuzz_dep_.CoverTab[6432]++
						if len(ip) == IPv4len && func() bool {
//line /snap/go/10455/src/net/ip.go:363
		_go_fuzz_dep_.CoverTab[6437]++
//line /snap/go/10455/src/net/ip.go:363
		return len(x) == IPv6len
//line /snap/go/10455/src/net/ip.go:363
		// _ = "end of CoverTab[6437]"
//line /snap/go/10455/src/net/ip.go:363
	}() {
//line /snap/go/10455/src/net/ip.go:363
		_go_fuzz_dep_.CoverTab[528614]++
//line /snap/go/10455/src/net/ip.go:363
		_go_fuzz_dep_.CoverTab[6438]++
							return bytealg.Equal(x[0:12], v4InV6Prefix) && func() bool {
//line /snap/go/10455/src/net/ip.go:364
			_go_fuzz_dep_.CoverTab[6439]++
//line /snap/go/10455/src/net/ip.go:364
			return bytealg.Equal(ip, x[12:])
//line /snap/go/10455/src/net/ip.go:364
			// _ = "end of CoverTab[6439]"
//line /snap/go/10455/src/net/ip.go:364
		}()
//line /snap/go/10455/src/net/ip.go:364
		// _ = "end of CoverTab[6438]"
	} else {
//line /snap/go/10455/src/net/ip.go:365
		_go_fuzz_dep_.CoverTab[528615]++
//line /snap/go/10455/src/net/ip.go:365
		_go_fuzz_dep_.CoverTab[6440]++
//line /snap/go/10455/src/net/ip.go:365
		// _ = "end of CoverTab[6440]"
//line /snap/go/10455/src/net/ip.go:365
	}
//line /snap/go/10455/src/net/ip.go:365
	// _ = "end of CoverTab[6432]"
//line /snap/go/10455/src/net/ip.go:365
	_go_fuzz_dep_.CoverTab[6433]++
						if len(ip) == IPv6len && func() bool {
//line /snap/go/10455/src/net/ip.go:366
		_go_fuzz_dep_.CoverTab[6441]++
//line /snap/go/10455/src/net/ip.go:366
		return len(x) == IPv4len
//line /snap/go/10455/src/net/ip.go:366
		// _ = "end of CoverTab[6441]"
//line /snap/go/10455/src/net/ip.go:366
	}() {
//line /snap/go/10455/src/net/ip.go:366
		_go_fuzz_dep_.CoverTab[528616]++
//line /snap/go/10455/src/net/ip.go:366
		_go_fuzz_dep_.CoverTab[6442]++
							return bytealg.Equal(ip[0:12], v4InV6Prefix) && func() bool {
//line /snap/go/10455/src/net/ip.go:367
			_go_fuzz_dep_.CoverTab[6443]++
//line /snap/go/10455/src/net/ip.go:367
			return bytealg.Equal(ip[12:], x)
//line /snap/go/10455/src/net/ip.go:367
			// _ = "end of CoverTab[6443]"
//line /snap/go/10455/src/net/ip.go:367
		}()
//line /snap/go/10455/src/net/ip.go:367
		// _ = "end of CoverTab[6442]"
	} else {
//line /snap/go/10455/src/net/ip.go:368
		_go_fuzz_dep_.CoverTab[528617]++
//line /snap/go/10455/src/net/ip.go:368
		_go_fuzz_dep_.CoverTab[6444]++
//line /snap/go/10455/src/net/ip.go:368
		// _ = "end of CoverTab[6444]"
//line /snap/go/10455/src/net/ip.go:368
	}
//line /snap/go/10455/src/net/ip.go:368
	// _ = "end of CoverTab[6433]"
//line /snap/go/10455/src/net/ip.go:368
	_go_fuzz_dep_.CoverTab[6434]++
						return false
//line /snap/go/10455/src/net/ip.go:369
	// _ = "end of CoverTab[6434]"
}

func (ip IP) matchAddrFamily(x IP) bool {
//line /snap/go/10455/src/net/ip.go:372
	_go_fuzz_dep_.CoverTab[6445]++
						return ip.To4() != nil && func() bool {
//line /snap/go/10455/src/net/ip.go:373
		_go_fuzz_dep_.CoverTab[6446]++
//line /snap/go/10455/src/net/ip.go:373
		return x.To4() != nil
//line /snap/go/10455/src/net/ip.go:373
		// _ = "end of CoverTab[6446]"
//line /snap/go/10455/src/net/ip.go:373
	}() || func() bool {
//line /snap/go/10455/src/net/ip.go:373
		_go_fuzz_dep_.CoverTab[6447]++
//line /snap/go/10455/src/net/ip.go:373
		return ip.To16() != nil && func() bool {
//line /snap/go/10455/src/net/ip.go:373
			_go_fuzz_dep_.CoverTab[6448]++
//line /snap/go/10455/src/net/ip.go:373
			return ip.To4() == nil
//line /snap/go/10455/src/net/ip.go:373
			// _ = "end of CoverTab[6448]"
//line /snap/go/10455/src/net/ip.go:373
		}() && func() bool {
//line /snap/go/10455/src/net/ip.go:373
			_go_fuzz_dep_.CoverTab[6449]++
//line /snap/go/10455/src/net/ip.go:373
			return x.To16() != nil
//line /snap/go/10455/src/net/ip.go:373
			// _ = "end of CoverTab[6449]"
//line /snap/go/10455/src/net/ip.go:373
		}() && func() bool {
//line /snap/go/10455/src/net/ip.go:373
			_go_fuzz_dep_.CoverTab[6450]++
//line /snap/go/10455/src/net/ip.go:373
			return x.To4() == nil
//line /snap/go/10455/src/net/ip.go:373
			// _ = "end of CoverTab[6450]"
//line /snap/go/10455/src/net/ip.go:373
		}()
//line /snap/go/10455/src/net/ip.go:373
		// _ = "end of CoverTab[6447]"
//line /snap/go/10455/src/net/ip.go:373
	}()
//line /snap/go/10455/src/net/ip.go:373
	// _ = "end of CoverTab[6445]"
}

// If mask is a sequence of 1 bits followed by 0 bits,
//line /snap/go/10455/src/net/ip.go:376
// return the number of 1 bits.
//line /snap/go/10455/src/net/ip.go:378
func simpleMaskLength(mask IPMask) int {
//line /snap/go/10455/src/net/ip.go:378
	_go_fuzz_dep_.CoverTab[6451]++
						var n int
//line /snap/go/10455/src/net/ip.go:379
	_go_fuzz_dep_.CoverTab[786693] = 0
						for i, v := range mask {
//line /snap/go/10455/src/net/ip.go:380
		if _go_fuzz_dep_.CoverTab[786693] == 0 {
//line /snap/go/10455/src/net/ip.go:380
			_go_fuzz_dep_.CoverTab[528681]++
//line /snap/go/10455/src/net/ip.go:380
		} else {
//line /snap/go/10455/src/net/ip.go:380
			_go_fuzz_dep_.CoverTab[528682]++
//line /snap/go/10455/src/net/ip.go:380
		}
//line /snap/go/10455/src/net/ip.go:380
		_go_fuzz_dep_.CoverTab[786693] = 1
//line /snap/go/10455/src/net/ip.go:380
		_go_fuzz_dep_.CoverTab[6453]++
							if v == 0xff {
//line /snap/go/10455/src/net/ip.go:381
			_go_fuzz_dep_.CoverTab[528618]++
//line /snap/go/10455/src/net/ip.go:381
			_go_fuzz_dep_.CoverTab[6458]++
								n += 8
								continue
//line /snap/go/10455/src/net/ip.go:383
			// _ = "end of CoverTab[6458]"
		} else {
//line /snap/go/10455/src/net/ip.go:384
			_go_fuzz_dep_.CoverTab[528619]++
//line /snap/go/10455/src/net/ip.go:384
			_go_fuzz_dep_.CoverTab[6459]++
//line /snap/go/10455/src/net/ip.go:384
			// _ = "end of CoverTab[6459]"
//line /snap/go/10455/src/net/ip.go:384
		}
//line /snap/go/10455/src/net/ip.go:384
		// _ = "end of CoverTab[6453]"
//line /snap/go/10455/src/net/ip.go:384
		_go_fuzz_dep_.CoverTab[6454]++
//line /snap/go/10455/src/net/ip.go:384
		_go_fuzz_dep_.CoverTab[786694] = 0

//line /snap/go/10455/src/net/ip.go:387
		for v&0x80 != 0 {
//line /snap/go/10455/src/net/ip.go:387
			if _go_fuzz_dep_.CoverTab[786694] == 0 {
//line /snap/go/10455/src/net/ip.go:387
				_go_fuzz_dep_.CoverTab[528685]++
//line /snap/go/10455/src/net/ip.go:387
			} else {
//line /snap/go/10455/src/net/ip.go:387
				_go_fuzz_dep_.CoverTab[528686]++
//line /snap/go/10455/src/net/ip.go:387
			}
//line /snap/go/10455/src/net/ip.go:387
			_go_fuzz_dep_.CoverTab[786694] = 1
//line /snap/go/10455/src/net/ip.go:387
			_go_fuzz_dep_.CoverTab[6460]++
								n++
								v <<= 1
//line /snap/go/10455/src/net/ip.go:389
			// _ = "end of CoverTab[6460]"
		}
//line /snap/go/10455/src/net/ip.go:390
		if _go_fuzz_dep_.CoverTab[786694] == 0 {
//line /snap/go/10455/src/net/ip.go:390
			_go_fuzz_dep_.CoverTab[528687]++
//line /snap/go/10455/src/net/ip.go:390
		} else {
//line /snap/go/10455/src/net/ip.go:390
			_go_fuzz_dep_.CoverTab[528688]++
//line /snap/go/10455/src/net/ip.go:390
		}
//line /snap/go/10455/src/net/ip.go:390
		// _ = "end of CoverTab[6454]"
//line /snap/go/10455/src/net/ip.go:390
		_go_fuzz_dep_.CoverTab[6455]++

							if v != 0 {
//line /snap/go/10455/src/net/ip.go:392
			_go_fuzz_dep_.CoverTab[528620]++
//line /snap/go/10455/src/net/ip.go:392
			_go_fuzz_dep_.CoverTab[6461]++
								return -1
//line /snap/go/10455/src/net/ip.go:393
			// _ = "end of CoverTab[6461]"
		} else {
//line /snap/go/10455/src/net/ip.go:394
			_go_fuzz_dep_.CoverTab[528621]++
//line /snap/go/10455/src/net/ip.go:394
			_go_fuzz_dep_.CoverTab[6462]++
//line /snap/go/10455/src/net/ip.go:394
			// _ = "end of CoverTab[6462]"
//line /snap/go/10455/src/net/ip.go:394
		}
//line /snap/go/10455/src/net/ip.go:394
		// _ = "end of CoverTab[6455]"
//line /snap/go/10455/src/net/ip.go:394
		_go_fuzz_dep_.CoverTab[6456]++
//line /snap/go/10455/src/net/ip.go:394
		_go_fuzz_dep_.CoverTab[786695] = 0
							for i++; i < len(mask); i++ {
//line /snap/go/10455/src/net/ip.go:395
			if _go_fuzz_dep_.CoverTab[786695] == 0 {
//line /snap/go/10455/src/net/ip.go:395
				_go_fuzz_dep_.CoverTab[528689]++
//line /snap/go/10455/src/net/ip.go:395
			} else {
//line /snap/go/10455/src/net/ip.go:395
				_go_fuzz_dep_.CoverTab[528690]++
//line /snap/go/10455/src/net/ip.go:395
			}
//line /snap/go/10455/src/net/ip.go:395
			_go_fuzz_dep_.CoverTab[786695] = 1
//line /snap/go/10455/src/net/ip.go:395
			_go_fuzz_dep_.CoverTab[6463]++
								if mask[i] != 0 {
//line /snap/go/10455/src/net/ip.go:396
				_go_fuzz_dep_.CoverTab[528622]++
//line /snap/go/10455/src/net/ip.go:396
				_go_fuzz_dep_.CoverTab[6464]++
									return -1
//line /snap/go/10455/src/net/ip.go:397
				// _ = "end of CoverTab[6464]"
			} else {
//line /snap/go/10455/src/net/ip.go:398
				_go_fuzz_dep_.CoverTab[528623]++
//line /snap/go/10455/src/net/ip.go:398
				_go_fuzz_dep_.CoverTab[6465]++
//line /snap/go/10455/src/net/ip.go:398
				// _ = "end of CoverTab[6465]"
//line /snap/go/10455/src/net/ip.go:398
			}
//line /snap/go/10455/src/net/ip.go:398
			// _ = "end of CoverTab[6463]"
		}
//line /snap/go/10455/src/net/ip.go:399
		if _go_fuzz_dep_.CoverTab[786695] == 0 {
//line /snap/go/10455/src/net/ip.go:399
			_go_fuzz_dep_.CoverTab[528691]++
//line /snap/go/10455/src/net/ip.go:399
		} else {
//line /snap/go/10455/src/net/ip.go:399
			_go_fuzz_dep_.CoverTab[528692]++
//line /snap/go/10455/src/net/ip.go:399
		}
//line /snap/go/10455/src/net/ip.go:399
		// _ = "end of CoverTab[6456]"
//line /snap/go/10455/src/net/ip.go:399
		_go_fuzz_dep_.CoverTab[6457]++
							break
//line /snap/go/10455/src/net/ip.go:400
		// _ = "end of CoverTab[6457]"
	}
//line /snap/go/10455/src/net/ip.go:401
	if _go_fuzz_dep_.CoverTab[786693] == 0 {
//line /snap/go/10455/src/net/ip.go:401
		_go_fuzz_dep_.CoverTab[528683]++
//line /snap/go/10455/src/net/ip.go:401
	} else {
//line /snap/go/10455/src/net/ip.go:401
		_go_fuzz_dep_.CoverTab[528684]++
//line /snap/go/10455/src/net/ip.go:401
	}
//line /snap/go/10455/src/net/ip.go:401
	// _ = "end of CoverTab[6451]"
//line /snap/go/10455/src/net/ip.go:401
	_go_fuzz_dep_.CoverTab[6452]++
						return n
//line /snap/go/10455/src/net/ip.go:402
	// _ = "end of CoverTab[6452]"
}

// Size returns the number of leading ones and total bits in the mask.
//line /snap/go/10455/src/net/ip.go:405
// If the mask is not in the canonical form--ones followed by zeros--then
//line /snap/go/10455/src/net/ip.go:405
// Size returns 0, 0.
//line /snap/go/10455/src/net/ip.go:408
func (m IPMask) Size() (ones, bits int) {
//line /snap/go/10455/src/net/ip.go:408
	_go_fuzz_dep_.CoverTab[6466]++
						ones, bits = simpleMaskLength(m), len(m)*8
						if ones == -1 {
//line /snap/go/10455/src/net/ip.go:410
		_go_fuzz_dep_.CoverTab[528624]++
//line /snap/go/10455/src/net/ip.go:410
		_go_fuzz_dep_.CoverTab[6468]++
							return 0, 0
//line /snap/go/10455/src/net/ip.go:411
		// _ = "end of CoverTab[6468]"
	} else {
//line /snap/go/10455/src/net/ip.go:412
		_go_fuzz_dep_.CoverTab[528625]++
//line /snap/go/10455/src/net/ip.go:412
		_go_fuzz_dep_.CoverTab[6469]++
//line /snap/go/10455/src/net/ip.go:412
		// _ = "end of CoverTab[6469]"
//line /snap/go/10455/src/net/ip.go:412
	}
//line /snap/go/10455/src/net/ip.go:412
	// _ = "end of CoverTab[6466]"
//line /snap/go/10455/src/net/ip.go:412
	_go_fuzz_dep_.CoverTab[6467]++
						return
//line /snap/go/10455/src/net/ip.go:413
	// _ = "end of CoverTab[6467]"
}

// String returns the hexadecimal form of m, with no punctuation.
func (m IPMask) String() string {
//line /snap/go/10455/src/net/ip.go:417
	_go_fuzz_dep_.CoverTab[6470]++
						if len(m) == 0 {
//line /snap/go/10455/src/net/ip.go:418
		_go_fuzz_dep_.CoverTab[528626]++
//line /snap/go/10455/src/net/ip.go:418
		_go_fuzz_dep_.CoverTab[6472]++
							return "<nil>"
//line /snap/go/10455/src/net/ip.go:419
		// _ = "end of CoverTab[6472]"
	} else {
//line /snap/go/10455/src/net/ip.go:420
		_go_fuzz_dep_.CoverTab[528627]++
//line /snap/go/10455/src/net/ip.go:420
		_go_fuzz_dep_.CoverTab[6473]++
//line /snap/go/10455/src/net/ip.go:420
		// _ = "end of CoverTab[6473]"
//line /snap/go/10455/src/net/ip.go:420
	}
//line /snap/go/10455/src/net/ip.go:420
	// _ = "end of CoverTab[6470]"
//line /snap/go/10455/src/net/ip.go:420
	_go_fuzz_dep_.CoverTab[6471]++
						return hexString(m)
//line /snap/go/10455/src/net/ip.go:421
	// _ = "end of CoverTab[6471]"
}

func networkNumberAndMask(n *IPNet) (ip IP, m IPMask) {
//line /snap/go/10455/src/net/ip.go:424
	_go_fuzz_dep_.CoverTab[6474]++
						if ip = n.IP.To4(); ip == nil {
//line /snap/go/10455/src/net/ip.go:425
		_go_fuzz_dep_.CoverTab[528628]++
//line /snap/go/10455/src/net/ip.go:425
		_go_fuzz_dep_.CoverTab[6477]++
							ip = n.IP
							if len(ip) != IPv6len {
//line /snap/go/10455/src/net/ip.go:427
			_go_fuzz_dep_.CoverTab[528630]++
//line /snap/go/10455/src/net/ip.go:427
			_go_fuzz_dep_.CoverTab[6478]++
								return nil, nil
//line /snap/go/10455/src/net/ip.go:428
			// _ = "end of CoverTab[6478]"
		} else {
//line /snap/go/10455/src/net/ip.go:429
			_go_fuzz_dep_.CoverTab[528631]++
//line /snap/go/10455/src/net/ip.go:429
			_go_fuzz_dep_.CoverTab[6479]++
//line /snap/go/10455/src/net/ip.go:429
			// _ = "end of CoverTab[6479]"
//line /snap/go/10455/src/net/ip.go:429
		}
//line /snap/go/10455/src/net/ip.go:429
		// _ = "end of CoverTab[6477]"
	} else {
//line /snap/go/10455/src/net/ip.go:430
		_go_fuzz_dep_.CoverTab[528629]++
//line /snap/go/10455/src/net/ip.go:430
		_go_fuzz_dep_.CoverTab[6480]++
//line /snap/go/10455/src/net/ip.go:430
		// _ = "end of CoverTab[6480]"
//line /snap/go/10455/src/net/ip.go:430
	}
//line /snap/go/10455/src/net/ip.go:430
	// _ = "end of CoverTab[6474]"
//line /snap/go/10455/src/net/ip.go:430
	_go_fuzz_dep_.CoverTab[6475]++
						m = n.Mask
						switch len(m) {
	case IPv4len:
//line /snap/go/10455/src/net/ip.go:433
		_go_fuzz_dep_.CoverTab[528632]++
//line /snap/go/10455/src/net/ip.go:433
		_go_fuzz_dep_.CoverTab[6481]++
							if len(ip) != IPv4len {
//line /snap/go/10455/src/net/ip.go:434
			_go_fuzz_dep_.CoverTab[528635]++
//line /snap/go/10455/src/net/ip.go:434
			_go_fuzz_dep_.CoverTab[6484]++
								return nil, nil
//line /snap/go/10455/src/net/ip.go:435
			// _ = "end of CoverTab[6484]"
		} else {
//line /snap/go/10455/src/net/ip.go:436
			_go_fuzz_dep_.CoverTab[528636]++
//line /snap/go/10455/src/net/ip.go:436
			_go_fuzz_dep_.CoverTab[6485]++
//line /snap/go/10455/src/net/ip.go:436
			// _ = "end of CoverTab[6485]"
//line /snap/go/10455/src/net/ip.go:436
		}
//line /snap/go/10455/src/net/ip.go:436
		// _ = "end of CoverTab[6481]"
	case IPv6len:
//line /snap/go/10455/src/net/ip.go:437
		_go_fuzz_dep_.CoverTab[528633]++
//line /snap/go/10455/src/net/ip.go:437
		_go_fuzz_dep_.CoverTab[6482]++
							if len(ip) == IPv4len {
//line /snap/go/10455/src/net/ip.go:438
			_go_fuzz_dep_.CoverTab[528637]++
//line /snap/go/10455/src/net/ip.go:438
			_go_fuzz_dep_.CoverTab[6486]++
								m = m[12:]
//line /snap/go/10455/src/net/ip.go:439
			// _ = "end of CoverTab[6486]"
		} else {
//line /snap/go/10455/src/net/ip.go:440
			_go_fuzz_dep_.CoverTab[528638]++
//line /snap/go/10455/src/net/ip.go:440
			_go_fuzz_dep_.CoverTab[6487]++
//line /snap/go/10455/src/net/ip.go:440
			// _ = "end of CoverTab[6487]"
//line /snap/go/10455/src/net/ip.go:440
		}
//line /snap/go/10455/src/net/ip.go:440
		// _ = "end of CoverTab[6482]"
	default:
//line /snap/go/10455/src/net/ip.go:441
		_go_fuzz_dep_.CoverTab[528634]++
//line /snap/go/10455/src/net/ip.go:441
		_go_fuzz_dep_.CoverTab[6483]++
							return nil, nil
//line /snap/go/10455/src/net/ip.go:442
		// _ = "end of CoverTab[6483]"
	}
//line /snap/go/10455/src/net/ip.go:443
	// _ = "end of CoverTab[6475]"
//line /snap/go/10455/src/net/ip.go:443
	_go_fuzz_dep_.CoverTab[6476]++
						return
//line /snap/go/10455/src/net/ip.go:444
	// _ = "end of CoverTab[6476]"
}

// Contains reports whether the network includes ip.
func (n *IPNet) Contains(ip IP) bool {
//line /snap/go/10455/src/net/ip.go:448
	_go_fuzz_dep_.CoverTab[6488]++
						nn, m := networkNumberAndMask(n)
						if x := ip.To4(); x != nil {
//line /snap/go/10455/src/net/ip.go:450
		_go_fuzz_dep_.CoverTab[528639]++
//line /snap/go/10455/src/net/ip.go:450
		_go_fuzz_dep_.CoverTab[6492]++
							ip = x
//line /snap/go/10455/src/net/ip.go:451
		// _ = "end of CoverTab[6492]"
	} else {
//line /snap/go/10455/src/net/ip.go:452
		_go_fuzz_dep_.CoverTab[528640]++
//line /snap/go/10455/src/net/ip.go:452
		_go_fuzz_dep_.CoverTab[6493]++
//line /snap/go/10455/src/net/ip.go:452
		// _ = "end of CoverTab[6493]"
//line /snap/go/10455/src/net/ip.go:452
	}
//line /snap/go/10455/src/net/ip.go:452
	// _ = "end of CoverTab[6488]"
//line /snap/go/10455/src/net/ip.go:452
	_go_fuzz_dep_.CoverTab[6489]++
						l := len(ip)
						if l != len(nn) {
//line /snap/go/10455/src/net/ip.go:454
		_go_fuzz_dep_.CoverTab[528641]++
//line /snap/go/10455/src/net/ip.go:454
		_go_fuzz_dep_.CoverTab[6494]++
							return false
//line /snap/go/10455/src/net/ip.go:455
		// _ = "end of CoverTab[6494]"
	} else {
//line /snap/go/10455/src/net/ip.go:456
		_go_fuzz_dep_.CoverTab[528642]++
//line /snap/go/10455/src/net/ip.go:456
		_go_fuzz_dep_.CoverTab[6495]++
//line /snap/go/10455/src/net/ip.go:456
		// _ = "end of CoverTab[6495]"
//line /snap/go/10455/src/net/ip.go:456
	}
//line /snap/go/10455/src/net/ip.go:456
	// _ = "end of CoverTab[6489]"
//line /snap/go/10455/src/net/ip.go:456
	_go_fuzz_dep_.CoverTab[6490]++
//line /snap/go/10455/src/net/ip.go:456
	_go_fuzz_dep_.CoverTab[786696] = 0
						for i := 0; i < l; i++ {
//line /snap/go/10455/src/net/ip.go:457
		if _go_fuzz_dep_.CoverTab[786696] == 0 {
//line /snap/go/10455/src/net/ip.go:457
			_go_fuzz_dep_.CoverTab[528693]++
//line /snap/go/10455/src/net/ip.go:457
		} else {
//line /snap/go/10455/src/net/ip.go:457
			_go_fuzz_dep_.CoverTab[528694]++
//line /snap/go/10455/src/net/ip.go:457
		}
//line /snap/go/10455/src/net/ip.go:457
		_go_fuzz_dep_.CoverTab[786696] = 1
//line /snap/go/10455/src/net/ip.go:457
		_go_fuzz_dep_.CoverTab[6496]++
							if nn[i]&m[i] != ip[i]&m[i] {
//line /snap/go/10455/src/net/ip.go:458
			_go_fuzz_dep_.CoverTab[528643]++
//line /snap/go/10455/src/net/ip.go:458
			_go_fuzz_dep_.CoverTab[6497]++
								return false
//line /snap/go/10455/src/net/ip.go:459
			// _ = "end of CoverTab[6497]"
		} else {
//line /snap/go/10455/src/net/ip.go:460
			_go_fuzz_dep_.CoverTab[528644]++
//line /snap/go/10455/src/net/ip.go:460
			_go_fuzz_dep_.CoverTab[6498]++
//line /snap/go/10455/src/net/ip.go:460
			// _ = "end of CoverTab[6498]"
//line /snap/go/10455/src/net/ip.go:460
		}
//line /snap/go/10455/src/net/ip.go:460
		// _ = "end of CoverTab[6496]"
	}
//line /snap/go/10455/src/net/ip.go:461
	if _go_fuzz_dep_.CoverTab[786696] == 0 {
//line /snap/go/10455/src/net/ip.go:461
		_go_fuzz_dep_.CoverTab[528695]++
//line /snap/go/10455/src/net/ip.go:461
	} else {
//line /snap/go/10455/src/net/ip.go:461
		_go_fuzz_dep_.CoverTab[528696]++
//line /snap/go/10455/src/net/ip.go:461
	}
//line /snap/go/10455/src/net/ip.go:461
	// _ = "end of CoverTab[6490]"
//line /snap/go/10455/src/net/ip.go:461
	_go_fuzz_dep_.CoverTab[6491]++
						return true
//line /snap/go/10455/src/net/ip.go:462
	// _ = "end of CoverTab[6491]"
}

// Network returns the address's network name, "ip+net".
func (n *IPNet) Network() string {
//line /snap/go/10455/src/net/ip.go:466
	_go_fuzz_dep_.CoverTab[6499]++
//line /snap/go/10455/src/net/ip.go:466
	return "ip+net"
//line /snap/go/10455/src/net/ip.go:466
	// _ = "end of CoverTab[6499]"
//line /snap/go/10455/src/net/ip.go:466
}

// String returns the CIDR notation of n like "192.0.2.0/24"
//line /snap/go/10455/src/net/ip.go:468
// or "2001:db8::/48" as defined in RFC 4632 and RFC 4291.
//line /snap/go/10455/src/net/ip.go:468
// If the mask is not in the canonical form, it returns the
//line /snap/go/10455/src/net/ip.go:468
// string which consists of an IP address, followed by a slash
//line /snap/go/10455/src/net/ip.go:468
// character and a mask expressed as hexadecimal form with no
//line /snap/go/10455/src/net/ip.go:468
// punctuation like "198.51.100.0/c000ff00".
//line /snap/go/10455/src/net/ip.go:474
func (n *IPNet) String() string {
//line /snap/go/10455/src/net/ip.go:474
	_go_fuzz_dep_.CoverTab[6500]++
						if n == nil {
//line /snap/go/10455/src/net/ip.go:475
		_go_fuzz_dep_.CoverTab[528645]++
//line /snap/go/10455/src/net/ip.go:475
		_go_fuzz_dep_.CoverTab[6504]++
							return "<nil>"
//line /snap/go/10455/src/net/ip.go:476
		// _ = "end of CoverTab[6504]"
	} else {
//line /snap/go/10455/src/net/ip.go:477
		_go_fuzz_dep_.CoverTab[528646]++
//line /snap/go/10455/src/net/ip.go:477
		_go_fuzz_dep_.CoverTab[6505]++
//line /snap/go/10455/src/net/ip.go:477
		// _ = "end of CoverTab[6505]"
//line /snap/go/10455/src/net/ip.go:477
	}
//line /snap/go/10455/src/net/ip.go:477
	// _ = "end of CoverTab[6500]"
//line /snap/go/10455/src/net/ip.go:477
	_go_fuzz_dep_.CoverTab[6501]++
						nn, m := networkNumberAndMask(n)
						if nn == nil || func() bool {
//line /snap/go/10455/src/net/ip.go:479
		_go_fuzz_dep_.CoverTab[6506]++
//line /snap/go/10455/src/net/ip.go:479
		return m == nil
//line /snap/go/10455/src/net/ip.go:479
		// _ = "end of CoverTab[6506]"
//line /snap/go/10455/src/net/ip.go:479
	}() {
//line /snap/go/10455/src/net/ip.go:479
		_go_fuzz_dep_.CoverTab[528647]++
//line /snap/go/10455/src/net/ip.go:479
		_go_fuzz_dep_.CoverTab[6507]++
							return "<nil>"
//line /snap/go/10455/src/net/ip.go:480
		// _ = "end of CoverTab[6507]"
	} else {
//line /snap/go/10455/src/net/ip.go:481
		_go_fuzz_dep_.CoverTab[528648]++
//line /snap/go/10455/src/net/ip.go:481
		_go_fuzz_dep_.CoverTab[6508]++
//line /snap/go/10455/src/net/ip.go:481
		// _ = "end of CoverTab[6508]"
//line /snap/go/10455/src/net/ip.go:481
	}
//line /snap/go/10455/src/net/ip.go:481
	// _ = "end of CoverTab[6501]"
//line /snap/go/10455/src/net/ip.go:481
	_go_fuzz_dep_.CoverTab[6502]++
						l := simpleMaskLength(m)
						if l == -1 {
//line /snap/go/10455/src/net/ip.go:483
		_go_fuzz_dep_.CoverTab[528649]++
//line /snap/go/10455/src/net/ip.go:483
		_go_fuzz_dep_.CoverTab[6509]++
							return nn.String() + "/" + m.String()
//line /snap/go/10455/src/net/ip.go:484
		// _ = "end of CoverTab[6509]"
	} else {
//line /snap/go/10455/src/net/ip.go:485
		_go_fuzz_dep_.CoverTab[528650]++
//line /snap/go/10455/src/net/ip.go:485
		_go_fuzz_dep_.CoverTab[6510]++
//line /snap/go/10455/src/net/ip.go:485
		// _ = "end of CoverTab[6510]"
//line /snap/go/10455/src/net/ip.go:485
	}
//line /snap/go/10455/src/net/ip.go:485
	// _ = "end of CoverTab[6502]"
//line /snap/go/10455/src/net/ip.go:485
	_go_fuzz_dep_.CoverTab[6503]++
						return nn.String() + "/" + itoa.Uitoa(uint(l))
//line /snap/go/10455/src/net/ip.go:486
	// _ = "end of CoverTab[6503]"
}

// ParseIP parses s as an IP address, returning the result.
//line /snap/go/10455/src/net/ip.go:489
// The string s can be in IPv4 dotted decimal ("192.0.2.1"), IPv6
//line /snap/go/10455/src/net/ip.go:489
// ("2001:db8::68"), or IPv4-mapped IPv6 ("::ffff:192.0.2.1") form.
//line /snap/go/10455/src/net/ip.go:489
// If s is not a valid textual representation of an IP address,
//line /snap/go/10455/src/net/ip.go:489
// ParseIP returns nil.
//line /snap/go/10455/src/net/ip.go:494
func ParseIP(s string) IP {
//line /snap/go/10455/src/net/ip.go:494
	_go_fuzz_dep_.CoverTab[6511]++
						if addr, valid := parseIP(s); valid {
//line /snap/go/10455/src/net/ip.go:495
		_go_fuzz_dep_.CoverTab[528651]++
//line /snap/go/10455/src/net/ip.go:495
		_go_fuzz_dep_.CoverTab[6513]++
							return IP(addr[:])
//line /snap/go/10455/src/net/ip.go:496
		// _ = "end of CoverTab[6513]"
	} else {
//line /snap/go/10455/src/net/ip.go:497
		_go_fuzz_dep_.CoverTab[528652]++
//line /snap/go/10455/src/net/ip.go:497
		_go_fuzz_dep_.CoverTab[6514]++
//line /snap/go/10455/src/net/ip.go:497
		// _ = "end of CoverTab[6514]"
//line /snap/go/10455/src/net/ip.go:497
	}
//line /snap/go/10455/src/net/ip.go:497
	// _ = "end of CoverTab[6511]"
//line /snap/go/10455/src/net/ip.go:497
	_go_fuzz_dep_.CoverTab[6512]++
						return nil
//line /snap/go/10455/src/net/ip.go:498
	// _ = "end of CoverTab[6512]"
}

func parseIP(s string) ([16]byte, bool) {
//line /snap/go/10455/src/net/ip.go:501
	_go_fuzz_dep_.CoverTab[6515]++
						ip, err := netip.ParseAddr(s)
						if err != nil || func() bool {
//line /snap/go/10455/src/net/ip.go:503
		_go_fuzz_dep_.CoverTab[6517]++
//line /snap/go/10455/src/net/ip.go:503
		return ip.Zone() != ""
//line /snap/go/10455/src/net/ip.go:503
		// _ = "end of CoverTab[6517]"
//line /snap/go/10455/src/net/ip.go:503
	}() {
//line /snap/go/10455/src/net/ip.go:503
		_go_fuzz_dep_.CoverTab[528653]++
//line /snap/go/10455/src/net/ip.go:503
		_go_fuzz_dep_.CoverTab[6518]++
							return [16]byte{}, false
//line /snap/go/10455/src/net/ip.go:504
		// _ = "end of CoverTab[6518]"
	} else {
//line /snap/go/10455/src/net/ip.go:505
		_go_fuzz_dep_.CoverTab[528654]++
//line /snap/go/10455/src/net/ip.go:505
		_go_fuzz_dep_.CoverTab[6519]++
//line /snap/go/10455/src/net/ip.go:505
		// _ = "end of CoverTab[6519]"
//line /snap/go/10455/src/net/ip.go:505
	}
//line /snap/go/10455/src/net/ip.go:505
	// _ = "end of CoverTab[6515]"
//line /snap/go/10455/src/net/ip.go:505
	_go_fuzz_dep_.CoverTab[6516]++
						return ip.As16(), true
//line /snap/go/10455/src/net/ip.go:506
	// _ = "end of CoverTab[6516]"
}

// ParseCIDR parses s as a CIDR notation IP address and prefix length,
//line /snap/go/10455/src/net/ip.go:509
// like "192.0.2.0/24" or "2001:db8::/32", as defined in
//line /snap/go/10455/src/net/ip.go:509
// RFC 4632 and RFC 4291.
//line /snap/go/10455/src/net/ip.go:509
//
//line /snap/go/10455/src/net/ip.go:509
// It returns the IP address and the network implied by the IP and
//line /snap/go/10455/src/net/ip.go:509
// prefix length.
//line /snap/go/10455/src/net/ip.go:509
// For example, ParseCIDR("192.0.2.1/24") returns the IP address
//line /snap/go/10455/src/net/ip.go:509
// 192.0.2.1 and the network 192.0.2.0/24.
//line /snap/go/10455/src/net/ip.go:517
func ParseCIDR(s string) (IP, *IPNet, error) {
//line /snap/go/10455/src/net/ip.go:517
	_go_fuzz_dep_.CoverTab[6520]++
						i := bytealg.IndexByteString(s, '/')
						if i < 0 {
//line /snap/go/10455/src/net/ip.go:519
		_go_fuzz_dep_.CoverTab[528655]++
//line /snap/go/10455/src/net/ip.go:519
		_go_fuzz_dep_.CoverTab[6524]++
							return nil, nil, &ParseError{Type: "CIDR address", Text: s}
//line /snap/go/10455/src/net/ip.go:520
		// _ = "end of CoverTab[6524]"
	} else {
//line /snap/go/10455/src/net/ip.go:521
		_go_fuzz_dep_.CoverTab[528656]++
//line /snap/go/10455/src/net/ip.go:521
		_go_fuzz_dep_.CoverTab[6525]++
//line /snap/go/10455/src/net/ip.go:521
		// _ = "end of CoverTab[6525]"
//line /snap/go/10455/src/net/ip.go:521
	}
//line /snap/go/10455/src/net/ip.go:521
	// _ = "end of CoverTab[6520]"
//line /snap/go/10455/src/net/ip.go:521
	_go_fuzz_dep_.CoverTab[6521]++
						addr, mask := s[:i], s[i+1:]

						ipAddr, err := netip.ParseAddr(addr)
						if err != nil || func() bool {
//line /snap/go/10455/src/net/ip.go:525
		_go_fuzz_dep_.CoverTab[6526]++
//line /snap/go/10455/src/net/ip.go:525
		return ipAddr.Zone() != ""
//line /snap/go/10455/src/net/ip.go:525
		// _ = "end of CoverTab[6526]"
//line /snap/go/10455/src/net/ip.go:525
	}() {
//line /snap/go/10455/src/net/ip.go:525
		_go_fuzz_dep_.CoverTab[528657]++
//line /snap/go/10455/src/net/ip.go:525
		_go_fuzz_dep_.CoverTab[6527]++
							return nil, nil, &ParseError{Type: "CIDR address", Text: s}
//line /snap/go/10455/src/net/ip.go:526
		// _ = "end of CoverTab[6527]"
	} else {
//line /snap/go/10455/src/net/ip.go:527
		_go_fuzz_dep_.CoverTab[528658]++
//line /snap/go/10455/src/net/ip.go:527
		_go_fuzz_dep_.CoverTab[6528]++
//line /snap/go/10455/src/net/ip.go:527
		// _ = "end of CoverTab[6528]"
//line /snap/go/10455/src/net/ip.go:527
	}
//line /snap/go/10455/src/net/ip.go:527
	// _ = "end of CoverTab[6521]"
//line /snap/go/10455/src/net/ip.go:527
	_go_fuzz_dep_.CoverTab[6522]++

						n, i, ok := dtoi(mask)
						if !ok || func() bool {
//line /snap/go/10455/src/net/ip.go:530
		_go_fuzz_dep_.CoverTab[6529]++
//line /snap/go/10455/src/net/ip.go:530
		return i != len(mask)
//line /snap/go/10455/src/net/ip.go:530
		// _ = "end of CoverTab[6529]"
//line /snap/go/10455/src/net/ip.go:530
	}() || func() bool {
//line /snap/go/10455/src/net/ip.go:530
		_go_fuzz_dep_.CoverTab[6530]++
//line /snap/go/10455/src/net/ip.go:530
		return n < 0
//line /snap/go/10455/src/net/ip.go:530
		// _ = "end of CoverTab[6530]"
//line /snap/go/10455/src/net/ip.go:530
	}() || func() bool {
//line /snap/go/10455/src/net/ip.go:530
		_go_fuzz_dep_.CoverTab[6531]++
//line /snap/go/10455/src/net/ip.go:530
		return n > ipAddr.BitLen()
//line /snap/go/10455/src/net/ip.go:530
		// _ = "end of CoverTab[6531]"
//line /snap/go/10455/src/net/ip.go:530
	}() {
//line /snap/go/10455/src/net/ip.go:530
		_go_fuzz_dep_.CoverTab[528659]++
//line /snap/go/10455/src/net/ip.go:530
		_go_fuzz_dep_.CoverTab[6532]++
							return nil, nil, &ParseError{Type: "CIDR address", Text: s}
//line /snap/go/10455/src/net/ip.go:531
		// _ = "end of CoverTab[6532]"
	} else {
//line /snap/go/10455/src/net/ip.go:532
		_go_fuzz_dep_.CoverTab[528660]++
//line /snap/go/10455/src/net/ip.go:532
		_go_fuzz_dep_.CoverTab[6533]++
//line /snap/go/10455/src/net/ip.go:532
		// _ = "end of CoverTab[6533]"
//line /snap/go/10455/src/net/ip.go:532
	}
//line /snap/go/10455/src/net/ip.go:532
	// _ = "end of CoverTab[6522]"
//line /snap/go/10455/src/net/ip.go:532
	_go_fuzz_dep_.CoverTab[6523]++
						m := CIDRMask(n, ipAddr.BitLen())
						addr16 := ipAddr.As16()
						return IP(addr16[:]), &IPNet{IP: IP(addr16[:]).Mask(m), Mask: m}, nil
//line /snap/go/10455/src/net/ip.go:535
	// _ = "end of CoverTab[6523]"
}

func copyIP(x IP) IP {
//line /snap/go/10455/src/net/ip.go:538
	_go_fuzz_dep_.CoverTab[6534]++
						y := make(IP, len(x))
						copy(y, x)
						return y
//line /snap/go/10455/src/net/ip.go:541
	// _ = "end of CoverTab[6534]"
}

//line /snap/go/10455/src/net/ip.go:542
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /snap/go/10455/src/net/ip.go:542
var _ = _go_fuzz_dep_.CoverTab
