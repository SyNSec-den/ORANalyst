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

//line /usr/local/go/src/net/ip.go:13
package net

//line /usr/local/go/src/net/ip.go:13
import (
//line /usr/local/go/src/net/ip.go:13
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/ip.go:13
)
//line /usr/local/go/src/net/ip.go:13
import (
//line /usr/local/go/src/net/ip.go:13
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/ip.go:13
)

import (
	"internal/bytealg"
	"internal/itoa"
)

// IP address lengths (bytes).
const (
	IPv4len	= 4
	IPv6len	= 16
)

// An IP is a single IP address, a slice of bytes.
//line /usr/local/go/src/net/ip.go:26
// Functions in this package accept either 4-byte (IPv4)
//line /usr/local/go/src/net/ip.go:26
// or 16-byte (IPv6) slices as input.
//line /usr/local/go/src/net/ip.go:26
//
//line /usr/local/go/src/net/ip.go:26
// Note that in this documentation, referring to an
//line /usr/local/go/src/net/ip.go:26
// IP address as an IPv4 address or an IPv6 address
//line /usr/local/go/src/net/ip.go:26
// is a semantic property of the address, not just the
//line /usr/local/go/src/net/ip.go:26
// length of the byte slice: a 16-byte slice can still
//line /usr/local/go/src/net/ip.go:26
// be an IPv4 address.
//line /usr/local/go/src/net/ip.go:35
type IP []byte

// An IPMask is a bitmask that can be used to manipulate
//line /usr/local/go/src/net/ip.go:37
// IP addresses for IP addressing and routing.
//line /usr/local/go/src/net/ip.go:37
//
//line /usr/local/go/src/net/ip.go:37
// See type IPNet and func ParseCIDR for details.
//line /usr/local/go/src/net/ip.go:41
type IPMask []byte

// An IPNet represents an IP network.
type IPNet struct {
	IP	IP	// network number
	Mask	IPMask	// network mask
}

// IPv4 returns the IP address (in 16-byte form) of the
//line /usr/local/go/src/net/ip.go:49
// IPv4 address a.b.c.d.
//line /usr/local/go/src/net/ip.go:51
func IPv4(a, b, c, d byte) IP {
//line /usr/local/go/src/net/ip.go:51
	_go_fuzz_dep_.CoverTab[5912]++
					p := make(IP, IPv6len)
					copy(p, v4InV6Prefix)
					p[12] = a
					p[13] = b
					p[14] = c
					p[15] = d
					return p
//line /usr/local/go/src/net/ip.go:58
	// _ = "end of CoverTab[5912]"
}

var v4InV6Prefix = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xff, 0xff}

// IPv4Mask returns the IP mask (in 4-byte form) of the
//line /usr/local/go/src/net/ip.go:63
// IPv4 mask a.b.c.d.
//line /usr/local/go/src/net/ip.go:65
func IPv4Mask(a, b, c, d byte) IPMask {
//line /usr/local/go/src/net/ip.go:65
	_go_fuzz_dep_.CoverTab[5913]++
					p := make(IPMask, IPv4len)
					p[0] = a
					p[1] = b
					p[2] = c
					p[3] = d
					return p
//line /usr/local/go/src/net/ip.go:71
	// _ = "end of CoverTab[5913]"
}

// CIDRMask returns an IPMask consisting of 'ones' 1 bits
//line /usr/local/go/src/net/ip.go:74
// followed by 0s up to a total length of 'bits' bits.
//line /usr/local/go/src/net/ip.go:74
// For a mask of this form, CIDRMask is the inverse of IPMask.Size.
//line /usr/local/go/src/net/ip.go:77
func CIDRMask(ones, bits int) IPMask {
//line /usr/local/go/src/net/ip.go:77
	_go_fuzz_dep_.CoverTab[5914]++
					if bits != 8*IPv4len && func() bool {
//line /usr/local/go/src/net/ip.go:78
		_go_fuzz_dep_.CoverTab[5918]++
//line /usr/local/go/src/net/ip.go:78
		return bits != 8*IPv6len
//line /usr/local/go/src/net/ip.go:78
		// _ = "end of CoverTab[5918]"
//line /usr/local/go/src/net/ip.go:78
	}() {
//line /usr/local/go/src/net/ip.go:78
		_go_fuzz_dep_.CoverTab[5919]++
						return nil
//line /usr/local/go/src/net/ip.go:79
		// _ = "end of CoverTab[5919]"
	} else {
//line /usr/local/go/src/net/ip.go:80
		_go_fuzz_dep_.CoverTab[5920]++
//line /usr/local/go/src/net/ip.go:80
		// _ = "end of CoverTab[5920]"
//line /usr/local/go/src/net/ip.go:80
	}
//line /usr/local/go/src/net/ip.go:80
	// _ = "end of CoverTab[5914]"
//line /usr/local/go/src/net/ip.go:80
	_go_fuzz_dep_.CoverTab[5915]++
					if ones < 0 || func() bool {
//line /usr/local/go/src/net/ip.go:81
		_go_fuzz_dep_.CoverTab[5921]++
//line /usr/local/go/src/net/ip.go:81
		return ones > bits
//line /usr/local/go/src/net/ip.go:81
		// _ = "end of CoverTab[5921]"
//line /usr/local/go/src/net/ip.go:81
	}() {
//line /usr/local/go/src/net/ip.go:81
		_go_fuzz_dep_.CoverTab[5922]++
						return nil
//line /usr/local/go/src/net/ip.go:82
		// _ = "end of CoverTab[5922]"
	} else {
//line /usr/local/go/src/net/ip.go:83
		_go_fuzz_dep_.CoverTab[5923]++
//line /usr/local/go/src/net/ip.go:83
		// _ = "end of CoverTab[5923]"
//line /usr/local/go/src/net/ip.go:83
	}
//line /usr/local/go/src/net/ip.go:83
	// _ = "end of CoverTab[5915]"
//line /usr/local/go/src/net/ip.go:83
	_go_fuzz_dep_.CoverTab[5916]++
					l := bits / 8
					m := make(IPMask, l)
					n := uint(ones)
					for i := 0; i < l; i++ {
//line /usr/local/go/src/net/ip.go:87
		_go_fuzz_dep_.CoverTab[5924]++
						if n >= 8 {
//line /usr/local/go/src/net/ip.go:88
			_go_fuzz_dep_.CoverTab[5926]++
							m[i] = 0xff
							n -= 8
							continue
//line /usr/local/go/src/net/ip.go:91
			// _ = "end of CoverTab[5926]"
		} else {
//line /usr/local/go/src/net/ip.go:92
			_go_fuzz_dep_.CoverTab[5927]++
//line /usr/local/go/src/net/ip.go:92
			// _ = "end of CoverTab[5927]"
//line /usr/local/go/src/net/ip.go:92
		}
//line /usr/local/go/src/net/ip.go:92
		// _ = "end of CoverTab[5924]"
//line /usr/local/go/src/net/ip.go:92
		_go_fuzz_dep_.CoverTab[5925]++
						m[i] = ^byte(0xff >> n)
						n = 0
//line /usr/local/go/src/net/ip.go:94
		// _ = "end of CoverTab[5925]"
	}
//line /usr/local/go/src/net/ip.go:95
	// _ = "end of CoverTab[5916]"
//line /usr/local/go/src/net/ip.go:95
	_go_fuzz_dep_.CoverTab[5917]++
					return m
//line /usr/local/go/src/net/ip.go:96
	// _ = "end of CoverTab[5917]"
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
//line /usr/local/go/src/net/ip.go:117
// the IPv4 address "0.0.0.0" or the IPv6 address "::".
//line /usr/local/go/src/net/ip.go:119
func (ip IP) IsUnspecified() bool {
//line /usr/local/go/src/net/ip.go:119
	_go_fuzz_dep_.CoverTab[5928]++
					return ip.Equal(IPv4zero) || func() bool {
//line /usr/local/go/src/net/ip.go:120
		_go_fuzz_dep_.CoverTab[5929]++
//line /usr/local/go/src/net/ip.go:120
		return ip.Equal(IPv6unspecified)
//line /usr/local/go/src/net/ip.go:120
		// _ = "end of CoverTab[5929]"
//line /usr/local/go/src/net/ip.go:120
	}()
//line /usr/local/go/src/net/ip.go:120
	// _ = "end of CoverTab[5928]"
}

// IsLoopback reports whether ip is a loopback address.
func (ip IP) IsLoopback() bool {
//line /usr/local/go/src/net/ip.go:124
	_go_fuzz_dep_.CoverTab[5930]++
					if ip4 := ip.To4(); ip4 != nil {
//line /usr/local/go/src/net/ip.go:125
		_go_fuzz_dep_.CoverTab[5932]++
						return ip4[0] == 127
//line /usr/local/go/src/net/ip.go:126
		// _ = "end of CoverTab[5932]"
	} else {
//line /usr/local/go/src/net/ip.go:127
		_go_fuzz_dep_.CoverTab[5933]++
//line /usr/local/go/src/net/ip.go:127
		// _ = "end of CoverTab[5933]"
//line /usr/local/go/src/net/ip.go:127
	}
//line /usr/local/go/src/net/ip.go:127
	// _ = "end of CoverTab[5930]"
//line /usr/local/go/src/net/ip.go:127
	_go_fuzz_dep_.CoverTab[5931]++
					return ip.Equal(IPv6loopback)
//line /usr/local/go/src/net/ip.go:128
	// _ = "end of CoverTab[5931]"
}

// IsPrivate reports whether ip is a private address, according to
//line /usr/local/go/src/net/ip.go:131
// RFC 1918 (IPv4 addresses) and RFC 4193 (IPv6 addresses).
//line /usr/local/go/src/net/ip.go:133
func (ip IP) IsPrivate() bool {
//line /usr/local/go/src/net/ip.go:133
	_go_fuzz_dep_.CoverTab[5934]++
					if ip4 := ip.To4(); ip4 != nil {
//line /usr/local/go/src/net/ip.go:134
		_go_fuzz_dep_.CoverTab[5936]++

//line /usr/local/go/src/net/ip.go:141
		return ip4[0] == 10 || func() bool {
//line /usr/local/go/src/net/ip.go:141
			_go_fuzz_dep_.CoverTab[5937]++
//line /usr/local/go/src/net/ip.go:141
			return (ip4[0] == 172 && func() bool {
								_go_fuzz_dep_.CoverTab[5938]++
//line /usr/local/go/src/net/ip.go:142
				return ip4[1]&0xf0 == 16
//line /usr/local/go/src/net/ip.go:142
				// _ = "end of CoverTab[5938]"
//line /usr/local/go/src/net/ip.go:142
			}())
//line /usr/local/go/src/net/ip.go:142
			// _ = "end of CoverTab[5937]"
//line /usr/local/go/src/net/ip.go:142
		}() || func() bool {
//line /usr/local/go/src/net/ip.go:142
			_go_fuzz_dep_.CoverTab[5939]++
//line /usr/local/go/src/net/ip.go:142
			return (ip4[0] == 192 && func() bool {
								_go_fuzz_dep_.CoverTab[5940]++
//line /usr/local/go/src/net/ip.go:143
				return ip4[1] == 168
//line /usr/local/go/src/net/ip.go:143
				// _ = "end of CoverTab[5940]"
//line /usr/local/go/src/net/ip.go:143
			}())
//line /usr/local/go/src/net/ip.go:143
			// _ = "end of CoverTab[5939]"
//line /usr/local/go/src/net/ip.go:143
		}()
//line /usr/local/go/src/net/ip.go:143
		// _ = "end of CoverTab[5936]"
	} else {
//line /usr/local/go/src/net/ip.go:144
		_go_fuzz_dep_.CoverTab[5941]++
//line /usr/local/go/src/net/ip.go:144
		// _ = "end of CoverTab[5941]"
//line /usr/local/go/src/net/ip.go:144
	}
//line /usr/local/go/src/net/ip.go:144
	// _ = "end of CoverTab[5934]"
//line /usr/local/go/src/net/ip.go:144
	_go_fuzz_dep_.CoverTab[5935]++

//line /usr/local/go/src/net/ip.go:147
	return len(ip) == IPv6len && func() bool {
//line /usr/local/go/src/net/ip.go:147
		_go_fuzz_dep_.CoverTab[5942]++
//line /usr/local/go/src/net/ip.go:147
		return ip[0]&0xfe == 0xfc
//line /usr/local/go/src/net/ip.go:147
		// _ = "end of CoverTab[5942]"
//line /usr/local/go/src/net/ip.go:147
	}()
//line /usr/local/go/src/net/ip.go:147
	// _ = "end of CoverTab[5935]"
}

// IsMulticast reports whether ip is a multicast address.
func (ip IP) IsMulticast() bool {
//line /usr/local/go/src/net/ip.go:151
	_go_fuzz_dep_.CoverTab[5943]++
					if ip4 := ip.To4(); ip4 != nil {
//line /usr/local/go/src/net/ip.go:152
		_go_fuzz_dep_.CoverTab[5945]++
						return ip4[0]&0xf0 == 0xe0
//line /usr/local/go/src/net/ip.go:153
		// _ = "end of CoverTab[5945]"
	} else {
//line /usr/local/go/src/net/ip.go:154
		_go_fuzz_dep_.CoverTab[5946]++
//line /usr/local/go/src/net/ip.go:154
		// _ = "end of CoverTab[5946]"
//line /usr/local/go/src/net/ip.go:154
	}
//line /usr/local/go/src/net/ip.go:154
	// _ = "end of CoverTab[5943]"
//line /usr/local/go/src/net/ip.go:154
	_go_fuzz_dep_.CoverTab[5944]++
					return len(ip) == IPv6len && func() bool {
//line /usr/local/go/src/net/ip.go:155
		_go_fuzz_dep_.CoverTab[5947]++
//line /usr/local/go/src/net/ip.go:155
		return ip[0] == 0xff
//line /usr/local/go/src/net/ip.go:155
		// _ = "end of CoverTab[5947]"
//line /usr/local/go/src/net/ip.go:155
	}()
//line /usr/local/go/src/net/ip.go:155
	// _ = "end of CoverTab[5944]"
}

// IsInterfaceLocalMulticast reports whether ip is
//line /usr/local/go/src/net/ip.go:158
// an interface-local multicast address.
//line /usr/local/go/src/net/ip.go:160
func (ip IP) IsInterfaceLocalMulticast() bool {
//line /usr/local/go/src/net/ip.go:160
	_go_fuzz_dep_.CoverTab[5948]++
					return len(ip) == IPv6len && func() bool {
//line /usr/local/go/src/net/ip.go:161
		_go_fuzz_dep_.CoverTab[5949]++
//line /usr/local/go/src/net/ip.go:161
		return ip[0] == 0xff
//line /usr/local/go/src/net/ip.go:161
		// _ = "end of CoverTab[5949]"
//line /usr/local/go/src/net/ip.go:161
	}() && func() bool {
//line /usr/local/go/src/net/ip.go:161
		_go_fuzz_dep_.CoverTab[5950]++
//line /usr/local/go/src/net/ip.go:161
		return ip[1]&0x0f == 0x01
//line /usr/local/go/src/net/ip.go:161
		// _ = "end of CoverTab[5950]"
//line /usr/local/go/src/net/ip.go:161
	}()
//line /usr/local/go/src/net/ip.go:161
	// _ = "end of CoverTab[5948]"
}

// IsLinkLocalMulticast reports whether ip is a link-local
//line /usr/local/go/src/net/ip.go:164
// multicast address.
//line /usr/local/go/src/net/ip.go:166
func (ip IP) IsLinkLocalMulticast() bool {
//line /usr/local/go/src/net/ip.go:166
	_go_fuzz_dep_.CoverTab[5951]++
					if ip4 := ip.To4(); ip4 != nil {
//line /usr/local/go/src/net/ip.go:167
		_go_fuzz_dep_.CoverTab[5953]++
						return ip4[0] == 224 && func() bool {
//line /usr/local/go/src/net/ip.go:168
			_go_fuzz_dep_.CoverTab[5954]++
//line /usr/local/go/src/net/ip.go:168
			return ip4[1] == 0
//line /usr/local/go/src/net/ip.go:168
			// _ = "end of CoverTab[5954]"
//line /usr/local/go/src/net/ip.go:168
		}() && func() bool {
//line /usr/local/go/src/net/ip.go:168
			_go_fuzz_dep_.CoverTab[5955]++
//line /usr/local/go/src/net/ip.go:168
			return ip4[2] == 0
//line /usr/local/go/src/net/ip.go:168
			// _ = "end of CoverTab[5955]"
//line /usr/local/go/src/net/ip.go:168
		}()
//line /usr/local/go/src/net/ip.go:168
		// _ = "end of CoverTab[5953]"
	} else {
//line /usr/local/go/src/net/ip.go:169
		_go_fuzz_dep_.CoverTab[5956]++
//line /usr/local/go/src/net/ip.go:169
		// _ = "end of CoverTab[5956]"
//line /usr/local/go/src/net/ip.go:169
	}
//line /usr/local/go/src/net/ip.go:169
	// _ = "end of CoverTab[5951]"
//line /usr/local/go/src/net/ip.go:169
	_go_fuzz_dep_.CoverTab[5952]++
					return len(ip) == IPv6len && func() bool {
//line /usr/local/go/src/net/ip.go:170
		_go_fuzz_dep_.CoverTab[5957]++
//line /usr/local/go/src/net/ip.go:170
		return ip[0] == 0xff
//line /usr/local/go/src/net/ip.go:170
		// _ = "end of CoverTab[5957]"
//line /usr/local/go/src/net/ip.go:170
	}() && func() bool {
//line /usr/local/go/src/net/ip.go:170
		_go_fuzz_dep_.CoverTab[5958]++
//line /usr/local/go/src/net/ip.go:170
		return ip[1]&0x0f == 0x02
//line /usr/local/go/src/net/ip.go:170
		// _ = "end of CoverTab[5958]"
//line /usr/local/go/src/net/ip.go:170
	}()
//line /usr/local/go/src/net/ip.go:170
	// _ = "end of CoverTab[5952]"
}

// IsLinkLocalUnicast reports whether ip is a link-local
//line /usr/local/go/src/net/ip.go:173
// unicast address.
//line /usr/local/go/src/net/ip.go:175
func (ip IP) IsLinkLocalUnicast() bool {
//line /usr/local/go/src/net/ip.go:175
	_go_fuzz_dep_.CoverTab[5959]++
					if ip4 := ip.To4(); ip4 != nil {
//line /usr/local/go/src/net/ip.go:176
		_go_fuzz_dep_.CoverTab[5961]++
						return ip4[0] == 169 && func() bool {
//line /usr/local/go/src/net/ip.go:177
			_go_fuzz_dep_.CoverTab[5962]++
//line /usr/local/go/src/net/ip.go:177
			return ip4[1] == 254
//line /usr/local/go/src/net/ip.go:177
			// _ = "end of CoverTab[5962]"
//line /usr/local/go/src/net/ip.go:177
		}()
//line /usr/local/go/src/net/ip.go:177
		// _ = "end of CoverTab[5961]"
	} else {
//line /usr/local/go/src/net/ip.go:178
		_go_fuzz_dep_.CoverTab[5963]++
//line /usr/local/go/src/net/ip.go:178
		// _ = "end of CoverTab[5963]"
//line /usr/local/go/src/net/ip.go:178
	}
//line /usr/local/go/src/net/ip.go:178
	// _ = "end of CoverTab[5959]"
//line /usr/local/go/src/net/ip.go:178
	_go_fuzz_dep_.CoverTab[5960]++
					return len(ip) == IPv6len && func() bool {
//line /usr/local/go/src/net/ip.go:179
		_go_fuzz_dep_.CoverTab[5964]++
//line /usr/local/go/src/net/ip.go:179
		return ip[0] == 0xfe
//line /usr/local/go/src/net/ip.go:179
		// _ = "end of CoverTab[5964]"
//line /usr/local/go/src/net/ip.go:179
	}() && func() bool {
//line /usr/local/go/src/net/ip.go:179
		_go_fuzz_dep_.CoverTab[5965]++
//line /usr/local/go/src/net/ip.go:179
		return ip[1]&0xc0 == 0x80
//line /usr/local/go/src/net/ip.go:179
		// _ = "end of CoverTab[5965]"
//line /usr/local/go/src/net/ip.go:179
	}()
//line /usr/local/go/src/net/ip.go:179
	// _ = "end of CoverTab[5960]"
}

// IsGlobalUnicast reports whether ip is a global unicast
//line /usr/local/go/src/net/ip.go:182
// address.
//line /usr/local/go/src/net/ip.go:182
//
//line /usr/local/go/src/net/ip.go:182
// The identification of global unicast addresses uses address type
//line /usr/local/go/src/net/ip.go:182
// identification as defined in RFC 1122, RFC 4632 and RFC 4291 with
//line /usr/local/go/src/net/ip.go:182
// the exception of IPv4 directed broadcast addresses.
//line /usr/local/go/src/net/ip.go:182
// It returns true even if ip is in IPv4 private address space or
//line /usr/local/go/src/net/ip.go:182
// local IPv6 unicast address space.
//line /usr/local/go/src/net/ip.go:190
func (ip IP) IsGlobalUnicast() bool {
//line /usr/local/go/src/net/ip.go:190
	_go_fuzz_dep_.CoverTab[5966]++
					return (len(ip) == IPv4len || func() bool {
//line /usr/local/go/src/net/ip.go:191
		_go_fuzz_dep_.CoverTab[5967]++
//line /usr/local/go/src/net/ip.go:191
		return len(ip) == IPv6len
//line /usr/local/go/src/net/ip.go:191
		// _ = "end of CoverTab[5967]"
//line /usr/local/go/src/net/ip.go:191
	}()) && func() bool {
//line /usr/local/go/src/net/ip.go:191
		_go_fuzz_dep_.CoverTab[5968]++
//line /usr/local/go/src/net/ip.go:191
		return !ip.Equal(IPv4bcast)
						// _ = "end of CoverTab[5968]"
//line /usr/local/go/src/net/ip.go:192
	}() && func() bool {
//line /usr/local/go/src/net/ip.go:192
		_go_fuzz_dep_.CoverTab[5969]++
//line /usr/local/go/src/net/ip.go:192
		return !ip.IsUnspecified()
						// _ = "end of CoverTab[5969]"
//line /usr/local/go/src/net/ip.go:193
	}() && func() bool {
//line /usr/local/go/src/net/ip.go:193
		_go_fuzz_dep_.CoverTab[5970]++
//line /usr/local/go/src/net/ip.go:193
		return !ip.IsLoopback()
						// _ = "end of CoverTab[5970]"
//line /usr/local/go/src/net/ip.go:194
	}() && func() bool {
//line /usr/local/go/src/net/ip.go:194
		_go_fuzz_dep_.CoverTab[5971]++
//line /usr/local/go/src/net/ip.go:194
		return !ip.IsMulticast()
						// _ = "end of CoverTab[5971]"
//line /usr/local/go/src/net/ip.go:195
	}() && func() bool {
//line /usr/local/go/src/net/ip.go:195
		_go_fuzz_dep_.CoverTab[5972]++
//line /usr/local/go/src/net/ip.go:195
		return !ip.IsLinkLocalUnicast()
						// _ = "end of CoverTab[5972]"
//line /usr/local/go/src/net/ip.go:196
	}()
//line /usr/local/go/src/net/ip.go:196
	// _ = "end of CoverTab[5966]"
}

// Is p all zeros?
func isZeros(p IP) bool {
//line /usr/local/go/src/net/ip.go:200
	_go_fuzz_dep_.CoverTab[5973]++
					for i := 0; i < len(p); i++ {
//line /usr/local/go/src/net/ip.go:201
		_go_fuzz_dep_.CoverTab[5975]++
						if p[i] != 0 {
//line /usr/local/go/src/net/ip.go:202
			_go_fuzz_dep_.CoverTab[5976]++
							return false
//line /usr/local/go/src/net/ip.go:203
			// _ = "end of CoverTab[5976]"
		} else {
//line /usr/local/go/src/net/ip.go:204
			_go_fuzz_dep_.CoverTab[5977]++
//line /usr/local/go/src/net/ip.go:204
			// _ = "end of CoverTab[5977]"
//line /usr/local/go/src/net/ip.go:204
		}
//line /usr/local/go/src/net/ip.go:204
		// _ = "end of CoverTab[5975]"
	}
//line /usr/local/go/src/net/ip.go:205
	// _ = "end of CoverTab[5973]"
//line /usr/local/go/src/net/ip.go:205
	_go_fuzz_dep_.CoverTab[5974]++
					return true
//line /usr/local/go/src/net/ip.go:206
	// _ = "end of CoverTab[5974]"
}

// To4 converts the IPv4 address ip to a 4-byte representation.
//line /usr/local/go/src/net/ip.go:209
// If ip is not an IPv4 address, To4 returns nil.
//line /usr/local/go/src/net/ip.go:211
func (ip IP) To4() IP {
//line /usr/local/go/src/net/ip.go:211
	_go_fuzz_dep_.CoverTab[5978]++
					if len(ip) == IPv4len {
//line /usr/local/go/src/net/ip.go:212
		_go_fuzz_dep_.CoverTab[5981]++
						return ip
//line /usr/local/go/src/net/ip.go:213
		// _ = "end of CoverTab[5981]"
	} else {
//line /usr/local/go/src/net/ip.go:214
		_go_fuzz_dep_.CoverTab[5982]++
//line /usr/local/go/src/net/ip.go:214
		// _ = "end of CoverTab[5982]"
//line /usr/local/go/src/net/ip.go:214
	}
//line /usr/local/go/src/net/ip.go:214
	// _ = "end of CoverTab[5978]"
//line /usr/local/go/src/net/ip.go:214
	_go_fuzz_dep_.CoverTab[5979]++
					if len(ip) == IPv6len && func() bool {
//line /usr/local/go/src/net/ip.go:215
		_go_fuzz_dep_.CoverTab[5983]++
//line /usr/local/go/src/net/ip.go:215
		return isZeros(ip[0:10])
						// _ = "end of CoverTab[5983]"
//line /usr/local/go/src/net/ip.go:216
	}() && func() bool {
//line /usr/local/go/src/net/ip.go:216
		_go_fuzz_dep_.CoverTab[5984]++
//line /usr/local/go/src/net/ip.go:216
		return ip[10] == 0xff
						// _ = "end of CoverTab[5984]"
//line /usr/local/go/src/net/ip.go:217
	}() && func() bool {
//line /usr/local/go/src/net/ip.go:217
		_go_fuzz_dep_.CoverTab[5985]++
//line /usr/local/go/src/net/ip.go:217
		return ip[11] == 0xff
						// _ = "end of CoverTab[5985]"
//line /usr/local/go/src/net/ip.go:218
	}() {
//line /usr/local/go/src/net/ip.go:218
		_go_fuzz_dep_.CoverTab[5986]++
						return ip[12:16]
//line /usr/local/go/src/net/ip.go:219
		// _ = "end of CoverTab[5986]"
	} else {
//line /usr/local/go/src/net/ip.go:220
		_go_fuzz_dep_.CoverTab[5987]++
//line /usr/local/go/src/net/ip.go:220
		// _ = "end of CoverTab[5987]"
//line /usr/local/go/src/net/ip.go:220
	}
//line /usr/local/go/src/net/ip.go:220
	// _ = "end of CoverTab[5979]"
//line /usr/local/go/src/net/ip.go:220
	_go_fuzz_dep_.CoverTab[5980]++
					return nil
//line /usr/local/go/src/net/ip.go:221
	// _ = "end of CoverTab[5980]"
}

// To16 converts the IP address ip to a 16-byte representation.
//line /usr/local/go/src/net/ip.go:224
// If ip is not an IP address (it is the wrong length), To16 returns nil.
//line /usr/local/go/src/net/ip.go:226
func (ip IP) To16() IP {
//line /usr/local/go/src/net/ip.go:226
	_go_fuzz_dep_.CoverTab[5988]++
					if len(ip) == IPv4len {
//line /usr/local/go/src/net/ip.go:227
		_go_fuzz_dep_.CoverTab[5991]++
						return IPv4(ip[0], ip[1], ip[2], ip[3])
//line /usr/local/go/src/net/ip.go:228
		// _ = "end of CoverTab[5991]"
	} else {
//line /usr/local/go/src/net/ip.go:229
		_go_fuzz_dep_.CoverTab[5992]++
//line /usr/local/go/src/net/ip.go:229
		// _ = "end of CoverTab[5992]"
//line /usr/local/go/src/net/ip.go:229
	}
//line /usr/local/go/src/net/ip.go:229
	// _ = "end of CoverTab[5988]"
//line /usr/local/go/src/net/ip.go:229
	_go_fuzz_dep_.CoverTab[5989]++
					if len(ip) == IPv6len {
//line /usr/local/go/src/net/ip.go:230
		_go_fuzz_dep_.CoverTab[5993]++
						return ip
//line /usr/local/go/src/net/ip.go:231
		// _ = "end of CoverTab[5993]"
	} else {
//line /usr/local/go/src/net/ip.go:232
		_go_fuzz_dep_.CoverTab[5994]++
//line /usr/local/go/src/net/ip.go:232
		// _ = "end of CoverTab[5994]"
//line /usr/local/go/src/net/ip.go:232
	}
//line /usr/local/go/src/net/ip.go:232
	// _ = "end of CoverTab[5989]"
//line /usr/local/go/src/net/ip.go:232
	_go_fuzz_dep_.CoverTab[5990]++
					return nil
//line /usr/local/go/src/net/ip.go:233
	// _ = "end of CoverTab[5990]"
}

// Default route masks for IPv4.
var (
	classAMask	= IPv4Mask(0xff, 0, 0, 0)
	classBMask	= IPv4Mask(0xff, 0xff, 0, 0)
	classCMask	= IPv4Mask(0xff, 0xff, 0xff, 0)
)

// DefaultMask returns the default IP mask for the IP address ip.
//line /usr/local/go/src/net/ip.go:243
// Only IPv4 addresses have default masks; DefaultMask returns
//line /usr/local/go/src/net/ip.go:243
// nil if ip is not a valid IPv4 address.
//line /usr/local/go/src/net/ip.go:246
func (ip IP) DefaultMask() IPMask {
//line /usr/local/go/src/net/ip.go:246
	_go_fuzz_dep_.CoverTab[5995]++
					if ip = ip.To4(); ip == nil {
//line /usr/local/go/src/net/ip.go:247
		_go_fuzz_dep_.CoverTab[5997]++
						return nil
//line /usr/local/go/src/net/ip.go:248
		// _ = "end of CoverTab[5997]"
	} else {
//line /usr/local/go/src/net/ip.go:249
		_go_fuzz_dep_.CoverTab[5998]++
//line /usr/local/go/src/net/ip.go:249
		// _ = "end of CoverTab[5998]"
//line /usr/local/go/src/net/ip.go:249
	}
//line /usr/local/go/src/net/ip.go:249
	// _ = "end of CoverTab[5995]"
//line /usr/local/go/src/net/ip.go:249
	_go_fuzz_dep_.CoverTab[5996]++
					switch {
	case ip[0] < 0x80:
//line /usr/local/go/src/net/ip.go:251
		_go_fuzz_dep_.CoverTab[5999]++
						return classAMask
//line /usr/local/go/src/net/ip.go:252
		// _ = "end of CoverTab[5999]"
	case ip[0] < 0xC0:
//line /usr/local/go/src/net/ip.go:253
		_go_fuzz_dep_.CoverTab[6000]++
						return classBMask
//line /usr/local/go/src/net/ip.go:254
		// _ = "end of CoverTab[6000]"
	default:
//line /usr/local/go/src/net/ip.go:255
		_go_fuzz_dep_.CoverTab[6001]++
						return classCMask
//line /usr/local/go/src/net/ip.go:256
		// _ = "end of CoverTab[6001]"
	}
//line /usr/local/go/src/net/ip.go:257
	// _ = "end of CoverTab[5996]"
}

func allFF(b []byte) bool {
//line /usr/local/go/src/net/ip.go:260
	_go_fuzz_dep_.CoverTab[6002]++
					for _, c := range b {
//line /usr/local/go/src/net/ip.go:261
		_go_fuzz_dep_.CoverTab[6004]++
						if c != 0xff {
//line /usr/local/go/src/net/ip.go:262
			_go_fuzz_dep_.CoverTab[6005]++
							return false
//line /usr/local/go/src/net/ip.go:263
			// _ = "end of CoverTab[6005]"
		} else {
//line /usr/local/go/src/net/ip.go:264
			_go_fuzz_dep_.CoverTab[6006]++
//line /usr/local/go/src/net/ip.go:264
			// _ = "end of CoverTab[6006]"
//line /usr/local/go/src/net/ip.go:264
		}
//line /usr/local/go/src/net/ip.go:264
		// _ = "end of CoverTab[6004]"
	}
//line /usr/local/go/src/net/ip.go:265
	// _ = "end of CoverTab[6002]"
//line /usr/local/go/src/net/ip.go:265
	_go_fuzz_dep_.CoverTab[6003]++
					return true
//line /usr/local/go/src/net/ip.go:266
	// _ = "end of CoverTab[6003]"
}

// Mask returns the result of masking the IP address ip with mask.
func (ip IP) Mask(mask IPMask) IP {
//line /usr/local/go/src/net/ip.go:270
	_go_fuzz_dep_.CoverTab[6007]++
					if len(mask) == IPv6len && func() bool {
//line /usr/local/go/src/net/ip.go:271
		_go_fuzz_dep_.CoverTab[6012]++
//line /usr/local/go/src/net/ip.go:271
		return len(ip) == IPv4len
//line /usr/local/go/src/net/ip.go:271
		// _ = "end of CoverTab[6012]"
//line /usr/local/go/src/net/ip.go:271
	}() && func() bool {
//line /usr/local/go/src/net/ip.go:271
		_go_fuzz_dep_.CoverTab[6013]++
//line /usr/local/go/src/net/ip.go:271
		return allFF(mask[:12])
//line /usr/local/go/src/net/ip.go:271
		// _ = "end of CoverTab[6013]"
//line /usr/local/go/src/net/ip.go:271
	}() {
//line /usr/local/go/src/net/ip.go:271
		_go_fuzz_dep_.CoverTab[6014]++
						mask = mask[12:]
//line /usr/local/go/src/net/ip.go:272
		// _ = "end of CoverTab[6014]"
	} else {
//line /usr/local/go/src/net/ip.go:273
		_go_fuzz_dep_.CoverTab[6015]++
//line /usr/local/go/src/net/ip.go:273
		// _ = "end of CoverTab[6015]"
//line /usr/local/go/src/net/ip.go:273
	}
//line /usr/local/go/src/net/ip.go:273
	// _ = "end of CoverTab[6007]"
//line /usr/local/go/src/net/ip.go:273
	_go_fuzz_dep_.CoverTab[6008]++
					if len(mask) == IPv4len && func() bool {
//line /usr/local/go/src/net/ip.go:274
		_go_fuzz_dep_.CoverTab[6016]++
//line /usr/local/go/src/net/ip.go:274
		return len(ip) == IPv6len
//line /usr/local/go/src/net/ip.go:274
		// _ = "end of CoverTab[6016]"
//line /usr/local/go/src/net/ip.go:274
	}() && func() bool {
//line /usr/local/go/src/net/ip.go:274
		_go_fuzz_dep_.CoverTab[6017]++
//line /usr/local/go/src/net/ip.go:274
		return bytealg.Equal(ip[:12], v4InV6Prefix)
//line /usr/local/go/src/net/ip.go:274
		// _ = "end of CoverTab[6017]"
//line /usr/local/go/src/net/ip.go:274
	}() {
//line /usr/local/go/src/net/ip.go:274
		_go_fuzz_dep_.CoverTab[6018]++
						ip = ip[12:]
//line /usr/local/go/src/net/ip.go:275
		// _ = "end of CoverTab[6018]"
	} else {
//line /usr/local/go/src/net/ip.go:276
		_go_fuzz_dep_.CoverTab[6019]++
//line /usr/local/go/src/net/ip.go:276
		// _ = "end of CoverTab[6019]"
//line /usr/local/go/src/net/ip.go:276
	}
//line /usr/local/go/src/net/ip.go:276
	// _ = "end of CoverTab[6008]"
//line /usr/local/go/src/net/ip.go:276
	_go_fuzz_dep_.CoverTab[6009]++
					n := len(ip)
					if n != len(mask) {
//line /usr/local/go/src/net/ip.go:278
		_go_fuzz_dep_.CoverTab[6020]++
						return nil
//line /usr/local/go/src/net/ip.go:279
		// _ = "end of CoverTab[6020]"
	} else {
//line /usr/local/go/src/net/ip.go:280
		_go_fuzz_dep_.CoverTab[6021]++
//line /usr/local/go/src/net/ip.go:280
		// _ = "end of CoverTab[6021]"
//line /usr/local/go/src/net/ip.go:280
	}
//line /usr/local/go/src/net/ip.go:280
	// _ = "end of CoverTab[6009]"
//line /usr/local/go/src/net/ip.go:280
	_go_fuzz_dep_.CoverTab[6010]++
					out := make(IP, n)
					for i := 0; i < n; i++ {
//line /usr/local/go/src/net/ip.go:282
		_go_fuzz_dep_.CoverTab[6022]++
						out[i] = ip[i] & mask[i]
//line /usr/local/go/src/net/ip.go:283
		// _ = "end of CoverTab[6022]"
	}
//line /usr/local/go/src/net/ip.go:284
	// _ = "end of CoverTab[6010]"
//line /usr/local/go/src/net/ip.go:284
	_go_fuzz_dep_.CoverTab[6011]++
					return out
//line /usr/local/go/src/net/ip.go:285
	// _ = "end of CoverTab[6011]"
}

// ubtoa encodes the string form of the integer v to dst[start:] and
//line /usr/local/go/src/net/ip.go:288
// returns the number of bytes written to dst. The caller must ensure
//line /usr/local/go/src/net/ip.go:288
// that dst has sufficient length.
//line /usr/local/go/src/net/ip.go:291
func ubtoa(dst []byte, start int, v byte) int {
//line /usr/local/go/src/net/ip.go:291
	_go_fuzz_dep_.CoverTab[6023]++
					if v < 10 {
//line /usr/local/go/src/net/ip.go:292
		_go_fuzz_dep_.CoverTab[6025]++
						dst[start] = v + '0'
						return 1
//line /usr/local/go/src/net/ip.go:294
		// _ = "end of CoverTab[6025]"
	} else {
//line /usr/local/go/src/net/ip.go:295
		_go_fuzz_dep_.CoverTab[6026]++
//line /usr/local/go/src/net/ip.go:295
		if v < 100 {
//line /usr/local/go/src/net/ip.go:295
			_go_fuzz_dep_.CoverTab[6027]++
							dst[start+1] = v%10 + '0'
							dst[start] = v/10 + '0'
							return 2
//line /usr/local/go/src/net/ip.go:298
			// _ = "end of CoverTab[6027]"
		} else {
//line /usr/local/go/src/net/ip.go:299
			_go_fuzz_dep_.CoverTab[6028]++
//line /usr/local/go/src/net/ip.go:299
			// _ = "end of CoverTab[6028]"
//line /usr/local/go/src/net/ip.go:299
		}
//line /usr/local/go/src/net/ip.go:299
		// _ = "end of CoverTab[6026]"
//line /usr/local/go/src/net/ip.go:299
	}
//line /usr/local/go/src/net/ip.go:299
	// _ = "end of CoverTab[6023]"
//line /usr/local/go/src/net/ip.go:299
	_go_fuzz_dep_.CoverTab[6024]++

					dst[start+2] = v%10 + '0'
					dst[start+1] = (v/10)%10 + '0'
					dst[start] = v/100 + '0'
					return 3
//line /usr/local/go/src/net/ip.go:304
	// _ = "end of CoverTab[6024]"
}

// String returns the string form of the IP address ip.
//line /usr/local/go/src/net/ip.go:307
// It returns one of 4 forms:
//line /usr/local/go/src/net/ip.go:307
//   - "<nil>", if ip has length 0
//line /usr/local/go/src/net/ip.go:307
//   - dotted decimal ("192.0.2.1"), if ip is an IPv4 or IP4-mapped IPv6 address
//line /usr/local/go/src/net/ip.go:307
//   - IPv6 conforming to RFC 5952 ("2001:db8::1"), if ip is a valid IPv6 address
//line /usr/local/go/src/net/ip.go:307
//   - the hexadecimal form of ip, without punctuation, if no other cases apply
//line /usr/local/go/src/net/ip.go:313
func (ip IP) String() string {
//line /usr/local/go/src/net/ip.go:313
	_go_fuzz_dep_.CoverTab[6029]++
					p := ip

					if len(ip) == 0 {
//line /usr/local/go/src/net/ip.go:316
		_go_fuzz_dep_.CoverTab[6036]++
						return "<nil>"
//line /usr/local/go/src/net/ip.go:317
		// _ = "end of CoverTab[6036]"
	} else {
//line /usr/local/go/src/net/ip.go:318
		_go_fuzz_dep_.CoverTab[6037]++
//line /usr/local/go/src/net/ip.go:318
		// _ = "end of CoverTab[6037]"
//line /usr/local/go/src/net/ip.go:318
	}
//line /usr/local/go/src/net/ip.go:318
	// _ = "end of CoverTab[6029]"
//line /usr/local/go/src/net/ip.go:318
	_go_fuzz_dep_.CoverTab[6030]++

//line /usr/local/go/src/net/ip.go:321
	if p4 := p.To4(); len(p4) == IPv4len {
//line /usr/local/go/src/net/ip.go:321
		_go_fuzz_dep_.CoverTab[6038]++
						const maxIPv4StringLen = len("255.255.255.255")
						b := make([]byte, maxIPv4StringLen)

						n := ubtoa(b, 0, p4[0])
						b[n] = '.'
						n++

						n += ubtoa(b, n, p4[1])
						b[n] = '.'
						n++

						n += ubtoa(b, n, p4[2])
						b[n] = '.'
						n++

						n += ubtoa(b, n, p4[3])
						return string(b[:n])
//line /usr/local/go/src/net/ip.go:338
		// _ = "end of CoverTab[6038]"
	} else {
//line /usr/local/go/src/net/ip.go:339
		_go_fuzz_dep_.CoverTab[6039]++
//line /usr/local/go/src/net/ip.go:339
		// _ = "end of CoverTab[6039]"
//line /usr/local/go/src/net/ip.go:339
	}
//line /usr/local/go/src/net/ip.go:339
	// _ = "end of CoverTab[6030]"
//line /usr/local/go/src/net/ip.go:339
	_go_fuzz_dep_.CoverTab[6031]++
					if len(p) != IPv6len {
//line /usr/local/go/src/net/ip.go:340
		_go_fuzz_dep_.CoverTab[6040]++
						return "?" + hexString(ip)
//line /usr/local/go/src/net/ip.go:341
		// _ = "end of CoverTab[6040]"
	} else {
//line /usr/local/go/src/net/ip.go:342
		_go_fuzz_dep_.CoverTab[6041]++
//line /usr/local/go/src/net/ip.go:342
		// _ = "end of CoverTab[6041]"
//line /usr/local/go/src/net/ip.go:342
	}
//line /usr/local/go/src/net/ip.go:342
	// _ = "end of CoverTab[6031]"
//line /usr/local/go/src/net/ip.go:342
	_go_fuzz_dep_.CoverTab[6032]++

//line /usr/local/go/src/net/ip.go:345
	e0 := -1
	e1 := -1
	for i := 0; i < IPv6len; i += 2 {
//line /usr/local/go/src/net/ip.go:347
		_go_fuzz_dep_.CoverTab[6042]++
						j := i
						for j < IPv6len && func() bool {
//line /usr/local/go/src/net/ip.go:349
			_go_fuzz_dep_.CoverTab[6044]++
//line /usr/local/go/src/net/ip.go:349
			return p[j] == 0
//line /usr/local/go/src/net/ip.go:349
			// _ = "end of CoverTab[6044]"
//line /usr/local/go/src/net/ip.go:349
		}() && func() bool {
//line /usr/local/go/src/net/ip.go:349
			_go_fuzz_dep_.CoverTab[6045]++
//line /usr/local/go/src/net/ip.go:349
			return p[j+1] == 0
//line /usr/local/go/src/net/ip.go:349
			// _ = "end of CoverTab[6045]"
//line /usr/local/go/src/net/ip.go:349
		}() {
//line /usr/local/go/src/net/ip.go:349
			_go_fuzz_dep_.CoverTab[6046]++
							j += 2
//line /usr/local/go/src/net/ip.go:350
			// _ = "end of CoverTab[6046]"
		}
//line /usr/local/go/src/net/ip.go:351
		// _ = "end of CoverTab[6042]"
//line /usr/local/go/src/net/ip.go:351
		_go_fuzz_dep_.CoverTab[6043]++
						if j > i && func() bool {
//line /usr/local/go/src/net/ip.go:352
			_go_fuzz_dep_.CoverTab[6047]++
//line /usr/local/go/src/net/ip.go:352
			return j-i > e1-e0
//line /usr/local/go/src/net/ip.go:352
			// _ = "end of CoverTab[6047]"
//line /usr/local/go/src/net/ip.go:352
		}() {
//line /usr/local/go/src/net/ip.go:352
			_go_fuzz_dep_.CoverTab[6048]++
							e0 = i
							e1 = j
							i = j
//line /usr/local/go/src/net/ip.go:355
			// _ = "end of CoverTab[6048]"
		} else {
//line /usr/local/go/src/net/ip.go:356
			_go_fuzz_dep_.CoverTab[6049]++
//line /usr/local/go/src/net/ip.go:356
			// _ = "end of CoverTab[6049]"
//line /usr/local/go/src/net/ip.go:356
		}
//line /usr/local/go/src/net/ip.go:356
		// _ = "end of CoverTab[6043]"
	}
//line /usr/local/go/src/net/ip.go:357
	// _ = "end of CoverTab[6032]"
//line /usr/local/go/src/net/ip.go:357
	_go_fuzz_dep_.CoverTab[6033]++

					if e1-e0 <= 2 {
//line /usr/local/go/src/net/ip.go:359
		_go_fuzz_dep_.CoverTab[6050]++
						e0 = -1
						e1 = -1
//line /usr/local/go/src/net/ip.go:361
		// _ = "end of CoverTab[6050]"
	} else {
//line /usr/local/go/src/net/ip.go:362
		_go_fuzz_dep_.CoverTab[6051]++
//line /usr/local/go/src/net/ip.go:362
		// _ = "end of CoverTab[6051]"
//line /usr/local/go/src/net/ip.go:362
	}
//line /usr/local/go/src/net/ip.go:362
	// _ = "end of CoverTab[6033]"
//line /usr/local/go/src/net/ip.go:362
	_go_fuzz_dep_.CoverTab[6034]++

					const maxLen = len("ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff")
					b := make([]byte, 0, maxLen)

//line /usr/local/go/src/net/ip.go:368
	for i := 0; i < IPv6len; i += 2 {
//line /usr/local/go/src/net/ip.go:368
		_go_fuzz_dep_.CoverTab[6052]++
						if i == e0 {
//line /usr/local/go/src/net/ip.go:369
			_go_fuzz_dep_.CoverTab[6054]++
							b = append(b, ':', ':')
							i = e1
							if i >= IPv6len {
//line /usr/local/go/src/net/ip.go:372
				_go_fuzz_dep_.CoverTab[6055]++
								break
//line /usr/local/go/src/net/ip.go:373
				// _ = "end of CoverTab[6055]"
			} else {
//line /usr/local/go/src/net/ip.go:374
				_go_fuzz_dep_.CoverTab[6056]++
//line /usr/local/go/src/net/ip.go:374
				// _ = "end of CoverTab[6056]"
//line /usr/local/go/src/net/ip.go:374
			}
//line /usr/local/go/src/net/ip.go:374
			// _ = "end of CoverTab[6054]"
		} else {
//line /usr/local/go/src/net/ip.go:375
			_go_fuzz_dep_.CoverTab[6057]++
//line /usr/local/go/src/net/ip.go:375
			if i > 0 {
//line /usr/local/go/src/net/ip.go:375
				_go_fuzz_dep_.CoverTab[6058]++
								b = append(b, ':')
//line /usr/local/go/src/net/ip.go:376
				// _ = "end of CoverTab[6058]"
			} else {
//line /usr/local/go/src/net/ip.go:377
				_go_fuzz_dep_.CoverTab[6059]++
//line /usr/local/go/src/net/ip.go:377
				// _ = "end of CoverTab[6059]"
//line /usr/local/go/src/net/ip.go:377
			}
//line /usr/local/go/src/net/ip.go:377
			// _ = "end of CoverTab[6057]"
//line /usr/local/go/src/net/ip.go:377
		}
//line /usr/local/go/src/net/ip.go:377
		// _ = "end of CoverTab[6052]"
//line /usr/local/go/src/net/ip.go:377
		_go_fuzz_dep_.CoverTab[6053]++
						b = appendHex(b, (uint32(p[i])<<8)|uint32(p[i+1]))
//line /usr/local/go/src/net/ip.go:378
		// _ = "end of CoverTab[6053]"
	}
//line /usr/local/go/src/net/ip.go:379
	// _ = "end of CoverTab[6034]"
//line /usr/local/go/src/net/ip.go:379
	_go_fuzz_dep_.CoverTab[6035]++
					return string(b)
//line /usr/local/go/src/net/ip.go:380
	// _ = "end of CoverTab[6035]"
}

func hexString(b []byte) string {
//line /usr/local/go/src/net/ip.go:383
	_go_fuzz_dep_.CoverTab[6060]++
					s := make([]byte, len(b)*2)
					for i, tn := range b {
//line /usr/local/go/src/net/ip.go:385
		_go_fuzz_dep_.CoverTab[6062]++
						s[i*2], s[i*2+1] = hexDigit[tn>>4], hexDigit[tn&0xf]
//line /usr/local/go/src/net/ip.go:386
		// _ = "end of CoverTab[6062]"
	}
//line /usr/local/go/src/net/ip.go:387
	// _ = "end of CoverTab[6060]"
//line /usr/local/go/src/net/ip.go:387
	_go_fuzz_dep_.CoverTab[6061]++
					return string(s)
//line /usr/local/go/src/net/ip.go:388
	// _ = "end of CoverTab[6061]"
}

// ipEmptyString is like ip.String except that it returns
//line /usr/local/go/src/net/ip.go:391
// an empty string when ip is unset.
//line /usr/local/go/src/net/ip.go:393
func ipEmptyString(ip IP) string {
//line /usr/local/go/src/net/ip.go:393
	_go_fuzz_dep_.CoverTab[6063]++
					if len(ip) == 0 {
//line /usr/local/go/src/net/ip.go:394
		_go_fuzz_dep_.CoverTab[6065]++
						return ""
//line /usr/local/go/src/net/ip.go:395
		// _ = "end of CoverTab[6065]"
	} else {
//line /usr/local/go/src/net/ip.go:396
		_go_fuzz_dep_.CoverTab[6066]++
//line /usr/local/go/src/net/ip.go:396
		// _ = "end of CoverTab[6066]"
//line /usr/local/go/src/net/ip.go:396
	}
//line /usr/local/go/src/net/ip.go:396
	// _ = "end of CoverTab[6063]"
//line /usr/local/go/src/net/ip.go:396
	_go_fuzz_dep_.CoverTab[6064]++
					return ip.String()
//line /usr/local/go/src/net/ip.go:397
	// _ = "end of CoverTab[6064]"
}

// MarshalText implements the encoding.TextMarshaler interface.
//line /usr/local/go/src/net/ip.go:400
// The encoding is the same as returned by String, with one exception:
//line /usr/local/go/src/net/ip.go:400
// When len(ip) is zero, it returns an empty slice.
//line /usr/local/go/src/net/ip.go:403
func (ip IP) MarshalText() ([]byte, error) {
//line /usr/local/go/src/net/ip.go:403
	_go_fuzz_dep_.CoverTab[6067]++
					if len(ip) == 0 {
//line /usr/local/go/src/net/ip.go:404
		_go_fuzz_dep_.CoverTab[6070]++
						return []byte(""), nil
//line /usr/local/go/src/net/ip.go:405
		// _ = "end of CoverTab[6070]"
	} else {
//line /usr/local/go/src/net/ip.go:406
		_go_fuzz_dep_.CoverTab[6071]++
//line /usr/local/go/src/net/ip.go:406
		// _ = "end of CoverTab[6071]"
//line /usr/local/go/src/net/ip.go:406
	}
//line /usr/local/go/src/net/ip.go:406
	// _ = "end of CoverTab[6067]"
//line /usr/local/go/src/net/ip.go:406
	_go_fuzz_dep_.CoverTab[6068]++
					if len(ip) != IPv4len && func() bool {
//line /usr/local/go/src/net/ip.go:407
		_go_fuzz_dep_.CoverTab[6072]++
//line /usr/local/go/src/net/ip.go:407
		return len(ip) != IPv6len
//line /usr/local/go/src/net/ip.go:407
		// _ = "end of CoverTab[6072]"
//line /usr/local/go/src/net/ip.go:407
	}() {
//line /usr/local/go/src/net/ip.go:407
		_go_fuzz_dep_.CoverTab[6073]++
						return nil, &AddrError{Err: "invalid IP address", Addr: hexString(ip)}
//line /usr/local/go/src/net/ip.go:408
		// _ = "end of CoverTab[6073]"
	} else {
//line /usr/local/go/src/net/ip.go:409
		_go_fuzz_dep_.CoverTab[6074]++
//line /usr/local/go/src/net/ip.go:409
		// _ = "end of CoverTab[6074]"
//line /usr/local/go/src/net/ip.go:409
	}
//line /usr/local/go/src/net/ip.go:409
	// _ = "end of CoverTab[6068]"
//line /usr/local/go/src/net/ip.go:409
	_go_fuzz_dep_.CoverTab[6069]++
					return []byte(ip.String()), nil
//line /usr/local/go/src/net/ip.go:410
	// _ = "end of CoverTab[6069]"
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
//line /usr/local/go/src/net/ip.go:413
// The IP address is expected in a form accepted by ParseIP.
//line /usr/local/go/src/net/ip.go:415
func (ip *IP) UnmarshalText(text []byte) error {
//line /usr/local/go/src/net/ip.go:415
	_go_fuzz_dep_.CoverTab[6075]++
					if len(text) == 0 {
//line /usr/local/go/src/net/ip.go:416
		_go_fuzz_dep_.CoverTab[6078]++
						*ip = nil
						return nil
//line /usr/local/go/src/net/ip.go:418
		// _ = "end of CoverTab[6078]"
	} else {
//line /usr/local/go/src/net/ip.go:419
		_go_fuzz_dep_.CoverTab[6079]++
//line /usr/local/go/src/net/ip.go:419
		// _ = "end of CoverTab[6079]"
//line /usr/local/go/src/net/ip.go:419
	}
//line /usr/local/go/src/net/ip.go:419
	// _ = "end of CoverTab[6075]"
//line /usr/local/go/src/net/ip.go:419
	_go_fuzz_dep_.CoverTab[6076]++
					s := string(text)
					x := ParseIP(s)
					if x == nil {
//line /usr/local/go/src/net/ip.go:422
		_go_fuzz_dep_.CoverTab[6080]++
						return &ParseError{Type: "IP address", Text: s}
//line /usr/local/go/src/net/ip.go:423
		// _ = "end of CoverTab[6080]"
	} else {
//line /usr/local/go/src/net/ip.go:424
		_go_fuzz_dep_.CoverTab[6081]++
//line /usr/local/go/src/net/ip.go:424
		// _ = "end of CoverTab[6081]"
//line /usr/local/go/src/net/ip.go:424
	}
//line /usr/local/go/src/net/ip.go:424
	// _ = "end of CoverTab[6076]"
//line /usr/local/go/src/net/ip.go:424
	_go_fuzz_dep_.CoverTab[6077]++
					*ip = x
					return nil
//line /usr/local/go/src/net/ip.go:426
	// _ = "end of CoverTab[6077]"
}

// Equal reports whether ip and x are the same IP address.
//line /usr/local/go/src/net/ip.go:429
// An IPv4 address and that same address in IPv6 form are
//line /usr/local/go/src/net/ip.go:429
// considered to be equal.
//line /usr/local/go/src/net/ip.go:432
func (ip IP) Equal(x IP) bool {
//line /usr/local/go/src/net/ip.go:432
	_go_fuzz_dep_.CoverTab[6082]++
					if len(ip) == len(x) {
//line /usr/local/go/src/net/ip.go:433
		_go_fuzz_dep_.CoverTab[6086]++
						return bytealg.Equal(ip, x)
//line /usr/local/go/src/net/ip.go:434
		// _ = "end of CoverTab[6086]"
	} else {
//line /usr/local/go/src/net/ip.go:435
		_go_fuzz_dep_.CoverTab[6087]++
//line /usr/local/go/src/net/ip.go:435
		// _ = "end of CoverTab[6087]"
//line /usr/local/go/src/net/ip.go:435
	}
//line /usr/local/go/src/net/ip.go:435
	// _ = "end of CoverTab[6082]"
//line /usr/local/go/src/net/ip.go:435
	_go_fuzz_dep_.CoverTab[6083]++
					if len(ip) == IPv4len && func() bool {
//line /usr/local/go/src/net/ip.go:436
		_go_fuzz_dep_.CoverTab[6088]++
//line /usr/local/go/src/net/ip.go:436
		return len(x) == IPv6len
//line /usr/local/go/src/net/ip.go:436
		// _ = "end of CoverTab[6088]"
//line /usr/local/go/src/net/ip.go:436
	}() {
//line /usr/local/go/src/net/ip.go:436
		_go_fuzz_dep_.CoverTab[6089]++
						return bytealg.Equal(x[0:12], v4InV6Prefix) && func() bool {
//line /usr/local/go/src/net/ip.go:437
			_go_fuzz_dep_.CoverTab[6090]++
//line /usr/local/go/src/net/ip.go:437
			return bytealg.Equal(ip, x[12:])
//line /usr/local/go/src/net/ip.go:437
			// _ = "end of CoverTab[6090]"
//line /usr/local/go/src/net/ip.go:437
		}()
//line /usr/local/go/src/net/ip.go:437
		// _ = "end of CoverTab[6089]"
	} else {
//line /usr/local/go/src/net/ip.go:438
		_go_fuzz_dep_.CoverTab[6091]++
//line /usr/local/go/src/net/ip.go:438
		// _ = "end of CoverTab[6091]"
//line /usr/local/go/src/net/ip.go:438
	}
//line /usr/local/go/src/net/ip.go:438
	// _ = "end of CoverTab[6083]"
//line /usr/local/go/src/net/ip.go:438
	_go_fuzz_dep_.CoverTab[6084]++
					if len(ip) == IPv6len && func() bool {
//line /usr/local/go/src/net/ip.go:439
		_go_fuzz_dep_.CoverTab[6092]++
//line /usr/local/go/src/net/ip.go:439
		return len(x) == IPv4len
//line /usr/local/go/src/net/ip.go:439
		// _ = "end of CoverTab[6092]"
//line /usr/local/go/src/net/ip.go:439
	}() {
//line /usr/local/go/src/net/ip.go:439
		_go_fuzz_dep_.CoverTab[6093]++
						return bytealg.Equal(ip[0:12], v4InV6Prefix) && func() bool {
//line /usr/local/go/src/net/ip.go:440
			_go_fuzz_dep_.CoverTab[6094]++
//line /usr/local/go/src/net/ip.go:440
			return bytealg.Equal(ip[12:], x)
//line /usr/local/go/src/net/ip.go:440
			// _ = "end of CoverTab[6094]"
//line /usr/local/go/src/net/ip.go:440
		}()
//line /usr/local/go/src/net/ip.go:440
		// _ = "end of CoverTab[6093]"
	} else {
//line /usr/local/go/src/net/ip.go:441
		_go_fuzz_dep_.CoverTab[6095]++
//line /usr/local/go/src/net/ip.go:441
		// _ = "end of CoverTab[6095]"
//line /usr/local/go/src/net/ip.go:441
	}
//line /usr/local/go/src/net/ip.go:441
	// _ = "end of CoverTab[6084]"
//line /usr/local/go/src/net/ip.go:441
	_go_fuzz_dep_.CoverTab[6085]++
					return false
//line /usr/local/go/src/net/ip.go:442
	// _ = "end of CoverTab[6085]"
}

func (ip IP) matchAddrFamily(x IP) bool {
//line /usr/local/go/src/net/ip.go:445
	_go_fuzz_dep_.CoverTab[6096]++
					return ip.To4() != nil && func() bool {
//line /usr/local/go/src/net/ip.go:446
		_go_fuzz_dep_.CoverTab[6097]++
//line /usr/local/go/src/net/ip.go:446
		return x.To4() != nil
//line /usr/local/go/src/net/ip.go:446
		// _ = "end of CoverTab[6097]"
//line /usr/local/go/src/net/ip.go:446
	}() || func() bool {
//line /usr/local/go/src/net/ip.go:446
		_go_fuzz_dep_.CoverTab[6098]++
//line /usr/local/go/src/net/ip.go:446
		return ip.To16() != nil && func() bool {
//line /usr/local/go/src/net/ip.go:446
			_go_fuzz_dep_.CoverTab[6099]++
//line /usr/local/go/src/net/ip.go:446
			return ip.To4() == nil
//line /usr/local/go/src/net/ip.go:446
			// _ = "end of CoverTab[6099]"
//line /usr/local/go/src/net/ip.go:446
		}() && func() bool {
//line /usr/local/go/src/net/ip.go:446
			_go_fuzz_dep_.CoverTab[6100]++
//line /usr/local/go/src/net/ip.go:446
			return x.To16() != nil
//line /usr/local/go/src/net/ip.go:446
			// _ = "end of CoverTab[6100]"
//line /usr/local/go/src/net/ip.go:446
		}() && func() bool {
//line /usr/local/go/src/net/ip.go:446
			_go_fuzz_dep_.CoverTab[6101]++
//line /usr/local/go/src/net/ip.go:446
			return x.To4() == nil
//line /usr/local/go/src/net/ip.go:446
			// _ = "end of CoverTab[6101]"
//line /usr/local/go/src/net/ip.go:446
		}()
//line /usr/local/go/src/net/ip.go:446
		// _ = "end of CoverTab[6098]"
//line /usr/local/go/src/net/ip.go:446
	}()
//line /usr/local/go/src/net/ip.go:446
	// _ = "end of CoverTab[6096]"
}

// If mask is a sequence of 1 bits followed by 0 bits,
//line /usr/local/go/src/net/ip.go:449
// return the number of 1 bits.
//line /usr/local/go/src/net/ip.go:451
func simpleMaskLength(mask IPMask) int {
//line /usr/local/go/src/net/ip.go:451
	_go_fuzz_dep_.CoverTab[6102]++
					var n int
					for i, v := range mask {
//line /usr/local/go/src/net/ip.go:453
		_go_fuzz_dep_.CoverTab[6104]++
						if v == 0xff {
//line /usr/local/go/src/net/ip.go:454
			_go_fuzz_dep_.CoverTab[6109]++
							n += 8
							continue
//line /usr/local/go/src/net/ip.go:456
			// _ = "end of CoverTab[6109]"
		} else {
//line /usr/local/go/src/net/ip.go:457
			_go_fuzz_dep_.CoverTab[6110]++
//line /usr/local/go/src/net/ip.go:457
			// _ = "end of CoverTab[6110]"
//line /usr/local/go/src/net/ip.go:457
		}
//line /usr/local/go/src/net/ip.go:457
		// _ = "end of CoverTab[6104]"
//line /usr/local/go/src/net/ip.go:457
		_go_fuzz_dep_.CoverTab[6105]++

//line /usr/local/go/src/net/ip.go:460
		for v&0x80 != 0 {
//line /usr/local/go/src/net/ip.go:460
			_go_fuzz_dep_.CoverTab[6111]++
							n++
							v <<= 1
//line /usr/local/go/src/net/ip.go:462
			// _ = "end of CoverTab[6111]"
		}
//line /usr/local/go/src/net/ip.go:463
		// _ = "end of CoverTab[6105]"
//line /usr/local/go/src/net/ip.go:463
		_go_fuzz_dep_.CoverTab[6106]++

						if v != 0 {
//line /usr/local/go/src/net/ip.go:465
			_go_fuzz_dep_.CoverTab[6112]++
							return -1
//line /usr/local/go/src/net/ip.go:466
			// _ = "end of CoverTab[6112]"
		} else {
//line /usr/local/go/src/net/ip.go:467
			_go_fuzz_dep_.CoverTab[6113]++
//line /usr/local/go/src/net/ip.go:467
			// _ = "end of CoverTab[6113]"
//line /usr/local/go/src/net/ip.go:467
		}
//line /usr/local/go/src/net/ip.go:467
		// _ = "end of CoverTab[6106]"
//line /usr/local/go/src/net/ip.go:467
		_go_fuzz_dep_.CoverTab[6107]++
						for i++; i < len(mask); i++ {
//line /usr/local/go/src/net/ip.go:468
			_go_fuzz_dep_.CoverTab[6114]++
							if mask[i] != 0 {
//line /usr/local/go/src/net/ip.go:469
				_go_fuzz_dep_.CoverTab[6115]++
								return -1
//line /usr/local/go/src/net/ip.go:470
				// _ = "end of CoverTab[6115]"
			} else {
//line /usr/local/go/src/net/ip.go:471
				_go_fuzz_dep_.CoverTab[6116]++
//line /usr/local/go/src/net/ip.go:471
				// _ = "end of CoverTab[6116]"
//line /usr/local/go/src/net/ip.go:471
			}
//line /usr/local/go/src/net/ip.go:471
			// _ = "end of CoverTab[6114]"
		}
//line /usr/local/go/src/net/ip.go:472
		// _ = "end of CoverTab[6107]"
//line /usr/local/go/src/net/ip.go:472
		_go_fuzz_dep_.CoverTab[6108]++
						break
//line /usr/local/go/src/net/ip.go:473
		// _ = "end of CoverTab[6108]"
	}
//line /usr/local/go/src/net/ip.go:474
	// _ = "end of CoverTab[6102]"
//line /usr/local/go/src/net/ip.go:474
	_go_fuzz_dep_.CoverTab[6103]++
					return n
//line /usr/local/go/src/net/ip.go:475
	// _ = "end of CoverTab[6103]"
}

// Size returns the number of leading ones and total bits in the mask.
//line /usr/local/go/src/net/ip.go:478
// If the mask is not in the canonical form--ones followed by zeros--then
//line /usr/local/go/src/net/ip.go:478
// Size returns 0, 0.
//line /usr/local/go/src/net/ip.go:481
func (m IPMask) Size() (ones, bits int) {
//line /usr/local/go/src/net/ip.go:481
	_go_fuzz_dep_.CoverTab[6117]++
					ones, bits = simpleMaskLength(m), len(m)*8
					if ones == -1 {
//line /usr/local/go/src/net/ip.go:483
		_go_fuzz_dep_.CoverTab[6119]++
						return 0, 0
//line /usr/local/go/src/net/ip.go:484
		// _ = "end of CoverTab[6119]"
	} else {
//line /usr/local/go/src/net/ip.go:485
		_go_fuzz_dep_.CoverTab[6120]++
//line /usr/local/go/src/net/ip.go:485
		// _ = "end of CoverTab[6120]"
//line /usr/local/go/src/net/ip.go:485
	}
//line /usr/local/go/src/net/ip.go:485
	// _ = "end of CoverTab[6117]"
//line /usr/local/go/src/net/ip.go:485
	_go_fuzz_dep_.CoverTab[6118]++
					return
//line /usr/local/go/src/net/ip.go:486
	// _ = "end of CoverTab[6118]"
}

// String returns the hexadecimal form of m, with no punctuation.
func (m IPMask) String() string {
//line /usr/local/go/src/net/ip.go:490
	_go_fuzz_dep_.CoverTab[6121]++
					if len(m) == 0 {
//line /usr/local/go/src/net/ip.go:491
		_go_fuzz_dep_.CoverTab[6123]++
						return "<nil>"
//line /usr/local/go/src/net/ip.go:492
		// _ = "end of CoverTab[6123]"
	} else {
//line /usr/local/go/src/net/ip.go:493
		_go_fuzz_dep_.CoverTab[6124]++
//line /usr/local/go/src/net/ip.go:493
		// _ = "end of CoverTab[6124]"
//line /usr/local/go/src/net/ip.go:493
	}
//line /usr/local/go/src/net/ip.go:493
	// _ = "end of CoverTab[6121]"
//line /usr/local/go/src/net/ip.go:493
	_go_fuzz_dep_.CoverTab[6122]++
					return hexString(m)
//line /usr/local/go/src/net/ip.go:494
	// _ = "end of CoverTab[6122]"
}

func networkNumberAndMask(n *IPNet) (ip IP, m IPMask) {
//line /usr/local/go/src/net/ip.go:497
	_go_fuzz_dep_.CoverTab[6125]++
					if ip = n.IP.To4(); ip == nil {
//line /usr/local/go/src/net/ip.go:498
		_go_fuzz_dep_.CoverTab[6128]++
						ip = n.IP
						if len(ip) != IPv6len {
//line /usr/local/go/src/net/ip.go:500
			_go_fuzz_dep_.CoverTab[6129]++
							return nil, nil
//line /usr/local/go/src/net/ip.go:501
			// _ = "end of CoverTab[6129]"
		} else {
//line /usr/local/go/src/net/ip.go:502
			_go_fuzz_dep_.CoverTab[6130]++
//line /usr/local/go/src/net/ip.go:502
			// _ = "end of CoverTab[6130]"
//line /usr/local/go/src/net/ip.go:502
		}
//line /usr/local/go/src/net/ip.go:502
		// _ = "end of CoverTab[6128]"
	} else {
//line /usr/local/go/src/net/ip.go:503
		_go_fuzz_dep_.CoverTab[6131]++
//line /usr/local/go/src/net/ip.go:503
		// _ = "end of CoverTab[6131]"
//line /usr/local/go/src/net/ip.go:503
	}
//line /usr/local/go/src/net/ip.go:503
	// _ = "end of CoverTab[6125]"
//line /usr/local/go/src/net/ip.go:503
	_go_fuzz_dep_.CoverTab[6126]++
					m = n.Mask
					switch len(m) {
	case IPv4len:
//line /usr/local/go/src/net/ip.go:506
		_go_fuzz_dep_.CoverTab[6132]++
						if len(ip) != IPv4len {
//line /usr/local/go/src/net/ip.go:507
			_go_fuzz_dep_.CoverTab[6135]++
							return nil, nil
//line /usr/local/go/src/net/ip.go:508
			// _ = "end of CoverTab[6135]"
		} else {
//line /usr/local/go/src/net/ip.go:509
			_go_fuzz_dep_.CoverTab[6136]++
//line /usr/local/go/src/net/ip.go:509
			// _ = "end of CoverTab[6136]"
//line /usr/local/go/src/net/ip.go:509
		}
//line /usr/local/go/src/net/ip.go:509
		// _ = "end of CoverTab[6132]"
	case IPv6len:
//line /usr/local/go/src/net/ip.go:510
		_go_fuzz_dep_.CoverTab[6133]++
						if len(ip) == IPv4len {
//line /usr/local/go/src/net/ip.go:511
			_go_fuzz_dep_.CoverTab[6137]++
							m = m[12:]
//line /usr/local/go/src/net/ip.go:512
			// _ = "end of CoverTab[6137]"
		} else {
//line /usr/local/go/src/net/ip.go:513
			_go_fuzz_dep_.CoverTab[6138]++
//line /usr/local/go/src/net/ip.go:513
			// _ = "end of CoverTab[6138]"
//line /usr/local/go/src/net/ip.go:513
		}
//line /usr/local/go/src/net/ip.go:513
		// _ = "end of CoverTab[6133]"
	default:
//line /usr/local/go/src/net/ip.go:514
		_go_fuzz_dep_.CoverTab[6134]++
						return nil, nil
//line /usr/local/go/src/net/ip.go:515
		// _ = "end of CoverTab[6134]"
	}
//line /usr/local/go/src/net/ip.go:516
	// _ = "end of CoverTab[6126]"
//line /usr/local/go/src/net/ip.go:516
	_go_fuzz_dep_.CoverTab[6127]++
					return
//line /usr/local/go/src/net/ip.go:517
	// _ = "end of CoverTab[6127]"
}

// Contains reports whether the network includes ip.
func (n *IPNet) Contains(ip IP) bool {
//line /usr/local/go/src/net/ip.go:521
	_go_fuzz_dep_.CoverTab[6139]++
					nn, m := networkNumberAndMask(n)
					if x := ip.To4(); x != nil {
//line /usr/local/go/src/net/ip.go:523
		_go_fuzz_dep_.CoverTab[6143]++
						ip = x
//line /usr/local/go/src/net/ip.go:524
		// _ = "end of CoverTab[6143]"
	} else {
//line /usr/local/go/src/net/ip.go:525
		_go_fuzz_dep_.CoverTab[6144]++
//line /usr/local/go/src/net/ip.go:525
		// _ = "end of CoverTab[6144]"
//line /usr/local/go/src/net/ip.go:525
	}
//line /usr/local/go/src/net/ip.go:525
	// _ = "end of CoverTab[6139]"
//line /usr/local/go/src/net/ip.go:525
	_go_fuzz_dep_.CoverTab[6140]++
					l := len(ip)
					if l != len(nn) {
//line /usr/local/go/src/net/ip.go:527
		_go_fuzz_dep_.CoverTab[6145]++
						return false
//line /usr/local/go/src/net/ip.go:528
		// _ = "end of CoverTab[6145]"
	} else {
//line /usr/local/go/src/net/ip.go:529
		_go_fuzz_dep_.CoverTab[6146]++
//line /usr/local/go/src/net/ip.go:529
		// _ = "end of CoverTab[6146]"
//line /usr/local/go/src/net/ip.go:529
	}
//line /usr/local/go/src/net/ip.go:529
	// _ = "end of CoverTab[6140]"
//line /usr/local/go/src/net/ip.go:529
	_go_fuzz_dep_.CoverTab[6141]++
					for i := 0; i < l; i++ {
//line /usr/local/go/src/net/ip.go:530
		_go_fuzz_dep_.CoverTab[6147]++
						if nn[i]&m[i] != ip[i]&m[i] {
//line /usr/local/go/src/net/ip.go:531
			_go_fuzz_dep_.CoverTab[6148]++
							return false
//line /usr/local/go/src/net/ip.go:532
			// _ = "end of CoverTab[6148]"
		} else {
//line /usr/local/go/src/net/ip.go:533
			_go_fuzz_dep_.CoverTab[6149]++
//line /usr/local/go/src/net/ip.go:533
			// _ = "end of CoverTab[6149]"
//line /usr/local/go/src/net/ip.go:533
		}
//line /usr/local/go/src/net/ip.go:533
		// _ = "end of CoverTab[6147]"
	}
//line /usr/local/go/src/net/ip.go:534
	// _ = "end of CoverTab[6141]"
//line /usr/local/go/src/net/ip.go:534
	_go_fuzz_dep_.CoverTab[6142]++
					return true
//line /usr/local/go/src/net/ip.go:535
	// _ = "end of CoverTab[6142]"
}

// Network returns the address's network name, "ip+net".
func (n *IPNet) Network() string {
//line /usr/local/go/src/net/ip.go:539
	_go_fuzz_dep_.CoverTab[6150]++
//line /usr/local/go/src/net/ip.go:539
	return "ip+net"
//line /usr/local/go/src/net/ip.go:539
	// _ = "end of CoverTab[6150]"
//line /usr/local/go/src/net/ip.go:539
}

// String returns the CIDR notation of n like "192.0.2.0/24"
//line /usr/local/go/src/net/ip.go:541
// or "2001:db8::/48" as defined in RFC 4632 and RFC 4291.
//line /usr/local/go/src/net/ip.go:541
// If the mask is not in the canonical form, it returns the
//line /usr/local/go/src/net/ip.go:541
// string which consists of an IP address, followed by a slash
//line /usr/local/go/src/net/ip.go:541
// character and a mask expressed as hexadecimal form with no
//line /usr/local/go/src/net/ip.go:541
// punctuation like "198.51.100.0/c000ff00".
//line /usr/local/go/src/net/ip.go:547
func (n *IPNet) String() string {
//line /usr/local/go/src/net/ip.go:547
	_go_fuzz_dep_.CoverTab[6151]++
					if n == nil {
//line /usr/local/go/src/net/ip.go:548
		_go_fuzz_dep_.CoverTab[6155]++
						return "<nil>"
//line /usr/local/go/src/net/ip.go:549
		// _ = "end of CoverTab[6155]"
	} else {
//line /usr/local/go/src/net/ip.go:550
		_go_fuzz_dep_.CoverTab[6156]++
//line /usr/local/go/src/net/ip.go:550
		// _ = "end of CoverTab[6156]"
//line /usr/local/go/src/net/ip.go:550
	}
//line /usr/local/go/src/net/ip.go:550
	// _ = "end of CoverTab[6151]"
//line /usr/local/go/src/net/ip.go:550
	_go_fuzz_dep_.CoverTab[6152]++
					nn, m := networkNumberAndMask(n)
					if nn == nil || func() bool {
//line /usr/local/go/src/net/ip.go:552
		_go_fuzz_dep_.CoverTab[6157]++
//line /usr/local/go/src/net/ip.go:552
		return m == nil
//line /usr/local/go/src/net/ip.go:552
		// _ = "end of CoverTab[6157]"
//line /usr/local/go/src/net/ip.go:552
	}() {
//line /usr/local/go/src/net/ip.go:552
		_go_fuzz_dep_.CoverTab[6158]++
						return "<nil>"
//line /usr/local/go/src/net/ip.go:553
		// _ = "end of CoverTab[6158]"
	} else {
//line /usr/local/go/src/net/ip.go:554
		_go_fuzz_dep_.CoverTab[6159]++
//line /usr/local/go/src/net/ip.go:554
		// _ = "end of CoverTab[6159]"
//line /usr/local/go/src/net/ip.go:554
	}
//line /usr/local/go/src/net/ip.go:554
	// _ = "end of CoverTab[6152]"
//line /usr/local/go/src/net/ip.go:554
	_go_fuzz_dep_.CoverTab[6153]++
					l := simpleMaskLength(m)
					if l == -1 {
//line /usr/local/go/src/net/ip.go:556
		_go_fuzz_dep_.CoverTab[6160]++
						return nn.String() + "/" + m.String()
//line /usr/local/go/src/net/ip.go:557
		// _ = "end of CoverTab[6160]"
	} else {
//line /usr/local/go/src/net/ip.go:558
		_go_fuzz_dep_.CoverTab[6161]++
//line /usr/local/go/src/net/ip.go:558
		// _ = "end of CoverTab[6161]"
//line /usr/local/go/src/net/ip.go:558
	}
//line /usr/local/go/src/net/ip.go:558
	// _ = "end of CoverTab[6153]"
//line /usr/local/go/src/net/ip.go:558
	_go_fuzz_dep_.CoverTab[6154]++
					return nn.String() + "/" + itoa.Uitoa(uint(l))
//line /usr/local/go/src/net/ip.go:559
	// _ = "end of CoverTab[6154]"
}

// Parse IPv4 address (d.d.d.d).
func parseIPv4(s string) IP {
//line /usr/local/go/src/net/ip.go:563
	_go_fuzz_dep_.CoverTab[6162]++
					var p [IPv4len]byte
					for i := 0; i < IPv4len; i++ {
//line /usr/local/go/src/net/ip.go:565
		_go_fuzz_dep_.CoverTab[6165]++
						if len(s) == 0 {
//line /usr/local/go/src/net/ip.go:566
			_go_fuzz_dep_.CoverTab[6170]++

							return nil
//line /usr/local/go/src/net/ip.go:568
			// _ = "end of CoverTab[6170]"
		} else {
//line /usr/local/go/src/net/ip.go:569
			_go_fuzz_dep_.CoverTab[6171]++
//line /usr/local/go/src/net/ip.go:569
			// _ = "end of CoverTab[6171]"
//line /usr/local/go/src/net/ip.go:569
		}
//line /usr/local/go/src/net/ip.go:569
		// _ = "end of CoverTab[6165]"
//line /usr/local/go/src/net/ip.go:569
		_go_fuzz_dep_.CoverTab[6166]++
						if i > 0 {
//line /usr/local/go/src/net/ip.go:570
			_go_fuzz_dep_.CoverTab[6172]++
							if s[0] != '.' {
//line /usr/local/go/src/net/ip.go:571
				_go_fuzz_dep_.CoverTab[6174]++
								return nil
//line /usr/local/go/src/net/ip.go:572
				// _ = "end of CoverTab[6174]"
			} else {
//line /usr/local/go/src/net/ip.go:573
				_go_fuzz_dep_.CoverTab[6175]++
//line /usr/local/go/src/net/ip.go:573
				// _ = "end of CoverTab[6175]"
//line /usr/local/go/src/net/ip.go:573
			}
//line /usr/local/go/src/net/ip.go:573
			// _ = "end of CoverTab[6172]"
//line /usr/local/go/src/net/ip.go:573
			_go_fuzz_dep_.CoverTab[6173]++
							s = s[1:]
//line /usr/local/go/src/net/ip.go:574
			// _ = "end of CoverTab[6173]"
		} else {
//line /usr/local/go/src/net/ip.go:575
			_go_fuzz_dep_.CoverTab[6176]++
//line /usr/local/go/src/net/ip.go:575
			// _ = "end of CoverTab[6176]"
//line /usr/local/go/src/net/ip.go:575
		}
//line /usr/local/go/src/net/ip.go:575
		// _ = "end of CoverTab[6166]"
//line /usr/local/go/src/net/ip.go:575
		_go_fuzz_dep_.CoverTab[6167]++
						n, c, ok := dtoi(s)
						if !ok || func() bool {
//line /usr/local/go/src/net/ip.go:577
			_go_fuzz_dep_.CoverTab[6177]++
//line /usr/local/go/src/net/ip.go:577
			return n > 0xFF
//line /usr/local/go/src/net/ip.go:577
			// _ = "end of CoverTab[6177]"
//line /usr/local/go/src/net/ip.go:577
		}() {
//line /usr/local/go/src/net/ip.go:577
			_go_fuzz_dep_.CoverTab[6178]++
							return nil
//line /usr/local/go/src/net/ip.go:578
			// _ = "end of CoverTab[6178]"
		} else {
//line /usr/local/go/src/net/ip.go:579
			_go_fuzz_dep_.CoverTab[6179]++
//line /usr/local/go/src/net/ip.go:579
			// _ = "end of CoverTab[6179]"
//line /usr/local/go/src/net/ip.go:579
		}
//line /usr/local/go/src/net/ip.go:579
		// _ = "end of CoverTab[6167]"
//line /usr/local/go/src/net/ip.go:579
		_go_fuzz_dep_.CoverTab[6168]++
						if c > 1 && func() bool {
//line /usr/local/go/src/net/ip.go:580
			_go_fuzz_dep_.CoverTab[6180]++
//line /usr/local/go/src/net/ip.go:580
			return s[0] == '0'
//line /usr/local/go/src/net/ip.go:580
			// _ = "end of CoverTab[6180]"
//line /usr/local/go/src/net/ip.go:580
		}() {
//line /usr/local/go/src/net/ip.go:580
			_go_fuzz_dep_.CoverTab[6181]++

							return nil
//line /usr/local/go/src/net/ip.go:582
			// _ = "end of CoverTab[6181]"
		} else {
//line /usr/local/go/src/net/ip.go:583
			_go_fuzz_dep_.CoverTab[6182]++
//line /usr/local/go/src/net/ip.go:583
			// _ = "end of CoverTab[6182]"
//line /usr/local/go/src/net/ip.go:583
		}
//line /usr/local/go/src/net/ip.go:583
		// _ = "end of CoverTab[6168]"
//line /usr/local/go/src/net/ip.go:583
		_go_fuzz_dep_.CoverTab[6169]++
						s = s[c:]
						p[i] = byte(n)
//line /usr/local/go/src/net/ip.go:585
		// _ = "end of CoverTab[6169]"
	}
//line /usr/local/go/src/net/ip.go:586
	// _ = "end of CoverTab[6162]"
//line /usr/local/go/src/net/ip.go:586
	_go_fuzz_dep_.CoverTab[6163]++
					if len(s) != 0 {
//line /usr/local/go/src/net/ip.go:587
		_go_fuzz_dep_.CoverTab[6183]++
						return nil
//line /usr/local/go/src/net/ip.go:588
		// _ = "end of CoverTab[6183]"
	} else {
//line /usr/local/go/src/net/ip.go:589
		_go_fuzz_dep_.CoverTab[6184]++
//line /usr/local/go/src/net/ip.go:589
		// _ = "end of CoverTab[6184]"
//line /usr/local/go/src/net/ip.go:589
	}
//line /usr/local/go/src/net/ip.go:589
	// _ = "end of CoverTab[6163]"
//line /usr/local/go/src/net/ip.go:589
	_go_fuzz_dep_.CoverTab[6164]++
					return IPv4(p[0], p[1], p[2], p[3])
//line /usr/local/go/src/net/ip.go:590
	// _ = "end of CoverTab[6164]"
}

// parseIPv6Zone parses s as a literal IPv6 address and its associated zone
//line /usr/local/go/src/net/ip.go:593
// identifier which is described in RFC 4007.
//line /usr/local/go/src/net/ip.go:595
func parseIPv6Zone(s string) (IP, string) {
//line /usr/local/go/src/net/ip.go:595
	_go_fuzz_dep_.CoverTab[6185]++
					s, zone := splitHostZone(s)
					return parseIPv6(s), zone
//line /usr/local/go/src/net/ip.go:597
	// _ = "end of CoverTab[6185]"
}

// parseIPv6 parses s as a literal IPv6 address described in RFC 4291
//line /usr/local/go/src/net/ip.go:600
// and RFC 5952.
//line /usr/local/go/src/net/ip.go:602
func parseIPv6(s string) (ip IP) {
//line /usr/local/go/src/net/ip.go:602
	_go_fuzz_dep_.CoverTab[6186]++
					ip = make(IP, IPv6len)
					ellipsis := -1

//line /usr/local/go/src/net/ip.go:607
	if len(s) >= 2 && func() bool {
//line /usr/local/go/src/net/ip.go:607
		_go_fuzz_dep_.CoverTab[6191]++
//line /usr/local/go/src/net/ip.go:607
		return s[0] == ':'
//line /usr/local/go/src/net/ip.go:607
		// _ = "end of CoverTab[6191]"
//line /usr/local/go/src/net/ip.go:607
	}() && func() bool {
//line /usr/local/go/src/net/ip.go:607
		_go_fuzz_dep_.CoverTab[6192]++
//line /usr/local/go/src/net/ip.go:607
		return s[1] == ':'
//line /usr/local/go/src/net/ip.go:607
		// _ = "end of CoverTab[6192]"
//line /usr/local/go/src/net/ip.go:607
	}() {
//line /usr/local/go/src/net/ip.go:607
		_go_fuzz_dep_.CoverTab[6193]++
						ellipsis = 0
						s = s[2:]

						if len(s) == 0 {
//line /usr/local/go/src/net/ip.go:611
			_go_fuzz_dep_.CoverTab[6194]++
							return ip
//line /usr/local/go/src/net/ip.go:612
			// _ = "end of CoverTab[6194]"
		} else {
//line /usr/local/go/src/net/ip.go:613
			_go_fuzz_dep_.CoverTab[6195]++
//line /usr/local/go/src/net/ip.go:613
			// _ = "end of CoverTab[6195]"
//line /usr/local/go/src/net/ip.go:613
		}
//line /usr/local/go/src/net/ip.go:613
		// _ = "end of CoverTab[6193]"
	} else {
//line /usr/local/go/src/net/ip.go:614
		_go_fuzz_dep_.CoverTab[6196]++
//line /usr/local/go/src/net/ip.go:614
		// _ = "end of CoverTab[6196]"
//line /usr/local/go/src/net/ip.go:614
	}
//line /usr/local/go/src/net/ip.go:614
	// _ = "end of CoverTab[6186]"
//line /usr/local/go/src/net/ip.go:614
	_go_fuzz_dep_.CoverTab[6187]++

//line /usr/local/go/src/net/ip.go:617
	i := 0
	for i < IPv6len {
//line /usr/local/go/src/net/ip.go:618
		_go_fuzz_dep_.CoverTab[6197]++

						n, c, ok := xtoi(s)
						if !ok || func() bool {
//line /usr/local/go/src/net/ip.go:621
			_go_fuzz_dep_.CoverTab[6202]++
//line /usr/local/go/src/net/ip.go:621
			return n > 0xFFFF
//line /usr/local/go/src/net/ip.go:621
			// _ = "end of CoverTab[6202]"
//line /usr/local/go/src/net/ip.go:621
		}() {
//line /usr/local/go/src/net/ip.go:621
			_go_fuzz_dep_.CoverTab[6203]++
							return nil
//line /usr/local/go/src/net/ip.go:622
			// _ = "end of CoverTab[6203]"
		} else {
//line /usr/local/go/src/net/ip.go:623
			_go_fuzz_dep_.CoverTab[6204]++
//line /usr/local/go/src/net/ip.go:623
			// _ = "end of CoverTab[6204]"
//line /usr/local/go/src/net/ip.go:623
		}
//line /usr/local/go/src/net/ip.go:623
		// _ = "end of CoverTab[6197]"
//line /usr/local/go/src/net/ip.go:623
		_go_fuzz_dep_.CoverTab[6198]++

//line /usr/local/go/src/net/ip.go:626
		if c < len(s) && func() bool {
//line /usr/local/go/src/net/ip.go:626
			_go_fuzz_dep_.CoverTab[6205]++
//line /usr/local/go/src/net/ip.go:626
			return s[c] == '.'
//line /usr/local/go/src/net/ip.go:626
			// _ = "end of CoverTab[6205]"
//line /usr/local/go/src/net/ip.go:626
		}() {
//line /usr/local/go/src/net/ip.go:626
			_go_fuzz_dep_.CoverTab[6206]++
							if ellipsis < 0 && func() bool {
//line /usr/local/go/src/net/ip.go:627
				_go_fuzz_dep_.CoverTab[6210]++
//line /usr/local/go/src/net/ip.go:627
				return i != IPv6len-IPv4len
//line /usr/local/go/src/net/ip.go:627
				// _ = "end of CoverTab[6210]"
//line /usr/local/go/src/net/ip.go:627
			}() {
//line /usr/local/go/src/net/ip.go:627
				_go_fuzz_dep_.CoverTab[6211]++

								return nil
//line /usr/local/go/src/net/ip.go:629
				// _ = "end of CoverTab[6211]"
			} else {
//line /usr/local/go/src/net/ip.go:630
				_go_fuzz_dep_.CoverTab[6212]++
//line /usr/local/go/src/net/ip.go:630
				// _ = "end of CoverTab[6212]"
//line /usr/local/go/src/net/ip.go:630
			}
//line /usr/local/go/src/net/ip.go:630
			// _ = "end of CoverTab[6206]"
//line /usr/local/go/src/net/ip.go:630
			_go_fuzz_dep_.CoverTab[6207]++
							if i+IPv4len > IPv6len {
//line /usr/local/go/src/net/ip.go:631
				_go_fuzz_dep_.CoverTab[6213]++

								return nil
//line /usr/local/go/src/net/ip.go:633
				// _ = "end of CoverTab[6213]"
			} else {
//line /usr/local/go/src/net/ip.go:634
				_go_fuzz_dep_.CoverTab[6214]++
//line /usr/local/go/src/net/ip.go:634
				// _ = "end of CoverTab[6214]"
//line /usr/local/go/src/net/ip.go:634
			}
//line /usr/local/go/src/net/ip.go:634
			// _ = "end of CoverTab[6207]"
//line /usr/local/go/src/net/ip.go:634
			_go_fuzz_dep_.CoverTab[6208]++
							ip4 := parseIPv4(s)
							if ip4 == nil {
//line /usr/local/go/src/net/ip.go:636
				_go_fuzz_dep_.CoverTab[6215]++
								return nil
//line /usr/local/go/src/net/ip.go:637
				// _ = "end of CoverTab[6215]"
			} else {
//line /usr/local/go/src/net/ip.go:638
				_go_fuzz_dep_.CoverTab[6216]++
//line /usr/local/go/src/net/ip.go:638
				// _ = "end of CoverTab[6216]"
//line /usr/local/go/src/net/ip.go:638
			}
//line /usr/local/go/src/net/ip.go:638
			// _ = "end of CoverTab[6208]"
//line /usr/local/go/src/net/ip.go:638
			_go_fuzz_dep_.CoverTab[6209]++
							ip[i] = ip4[12]
							ip[i+1] = ip4[13]
							ip[i+2] = ip4[14]
							ip[i+3] = ip4[15]
							s = ""
							i += IPv4len
							break
//line /usr/local/go/src/net/ip.go:645
			// _ = "end of CoverTab[6209]"
		} else {
//line /usr/local/go/src/net/ip.go:646
			_go_fuzz_dep_.CoverTab[6217]++
//line /usr/local/go/src/net/ip.go:646
			// _ = "end of CoverTab[6217]"
//line /usr/local/go/src/net/ip.go:646
		}
//line /usr/local/go/src/net/ip.go:646
		// _ = "end of CoverTab[6198]"
//line /usr/local/go/src/net/ip.go:646
		_go_fuzz_dep_.CoverTab[6199]++

//line /usr/local/go/src/net/ip.go:649
		ip[i] = byte(n >> 8)
						ip[i+1] = byte(n)
						i += 2

//line /usr/local/go/src/net/ip.go:654
		s = s[c:]
		if len(s) == 0 {
//line /usr/local/go/src/net/ip.go:655
			_go_fuzz_dep_.CoverTab[6218]++
							break
//line /usr/local/go/src/net/ip.go:656
			// _ = "end of CoverTab[6218]"
		} else {
//line /usr/local/go/src/net/ip.go:657
			_go_fuzz_dep_.CoverTab[6219]++
//line /usr/local/go/src/net/ip.go:657
			// _ = "end of CoverTab[6219]"
//line /usr/local/go/src/net/ip.go:657
		}
//line /usr/local/go/src/net/ip.go:657
		// _ = "end of CoverTab[6199]"
//line /usr/local/go/src/net/ip.go:657
		_go_fuzz_dep_.CoverTab[6200]++

//line /usr/local/go/src/net/ip.go:660
		if s[0] != ':' || func() bool {
//line /usr/local/go/src/net/ip.go:660
			_go_fuzz_dep_.CoverTab[6220]++
//line /usr/local/go/src/net/ip.go:660
			return len(s) == 1
//line /usr/local/go/src/net/ip.go:660
			// _ = "end of CoverTab[6220]"
//line /usr/local/go/src/net/ip.go:660
		}() {
//line /usr/local/go/src/net/ip.go:660
			_go_fuzz_dep_.CoverTab[6221]++
							return nil
//line /usr/local/go/src/net/ip.go:661
			// _ = "end of CoverTab[6221]"
		} else {
//line /usr/local/go/src/net/ip.go:662
			_go_fuzz_dep_.CoverTab[6222]++
//line /usr/local/go/src/net/ip.go:662
			// _ = "end of CoverTab[6222]"
//line /usr/local/go/src/net/ip.go:662
		}
//line /usr/local/go/src/net/ip.go:662
		// _ = "end of CoverTab[6200]"
//line /usr/local/go/src/net/ip.go:662
		_go_fuzz_dep_.CoverTab[6201]++
						s = s[1:]

//line /usr/local/go/src/net/ip.go:666
		if s[0] == ':' {
//line /usr/local/go/src/net/ip.go:666
			_go_fuzz_dep_.CoverTab[6223]++
							if ellipsis >= 0 {
//line /usr/local/go/src/net/ip.go:667
				_go_fuzz_dep_.CoverTab[6225]++
								return nil
//line /usr/local/go/src/net/ip.go:668
				// _ = "end of CoverTab[6225]"
			} else {
//line /usr/local/go/src/net/ip.go:669
				_go_fuzz_dep_.CoverTab[6226]++
//line /usr/local/go/src/net/ip.go:669
				// _ = "end of CoverTab[6226]"
//line /usr/local/go/src/net/ip.go:669
			}
//line /usr/local/go/src/net/ip.go:669
			// _ = "end of CoverTab[6223]"
//line /usr/local/go/src/net/ip.go:669
			_go_fuzz_dep_.CoverTab[6224]++
							ellipsis = i
							s = s[1:]
							if len(s) == 0 {
//line /usr/local/go/src/net/ip.go:672
				_go_fuzz_dep_.CoverTab[6227]++
								break
//line /usr/local/go/src/net/ip.go:673
				// _ = "end of CoverTab[6227]"
			} else {
//line /usr/local/go/src/net/ip.go:674
				_go_fuzz_dep_.CoverTab[6228]++
//line /usr/local/go/src/net/ip.go:674
				// _ = "end of CoverTab[6228]"
//line /usr/local/go/src/net/ip.go:674
			}
//line /usr/local/go/src/net/ip.go:674
			// _ = "end of CoverTab[6224]"
		} else {
//line /usr/local/go/src/net/ip.go:675
			_go_fuzz_dep_.CoverTab[6229]++
//line /usr/local/go/src/net/ip.go:675
			// _ = "end of CoverTab[6229]"
//line /usr/local/go/src/net/ip.go:675
		}
//line /usr/local/go/src/net/ip.go:675
		// _ = "end of CoverTab[6201]"
	}
//line /usr/local/go/src/net/ip.go:676
	// _ = "end of CoverTab[6187]"
//line /usr/local/go/src/net/ip.go:676
	_go_fuzz_dep_.CoverTab[6188]++

//line /usr/local/go/src/net/ip.go:679
	if len(s) != 0 {
//line /usr/local/go/src/net/ip.go:679
		_go_fuzz_dep_.CoverTab[6230]++
						return nil
//line /usr/local/go/src/net/ip.go:680
		// _ = "end of CoverTab[6230]"
	} else {
//line /usr/local/go/src/net/ip.go:681
		_go_fuzz_dep_.CoverTab[6231]++
//line /usr/local/go/src/net/ip.go:681
		// _ = "end of CoverTab[6231]"
//line /usr/local/go/src/net/ip.go:681
	}
//line /usr/local/go/src/net/ip.go:681
	// _ = "end of CoverTab[6188]"
//line /usr/local/go/src/net/ip.go:681
	_go_fuzz_dep_.CoverTab[6189]++

//line /usr/local/go/src/net/ip.go:684
	if i < IPv6len {
//line /usr/local/go/src/net/ip.go:684
		_go_fuzz_dep_.CoverTab[6232]++
						if ellipsis < 0 {
//line /usr/local/go/src/net/ip.go:685
			_go_fuzz_dep_.CoverTab[6235]++
							return nil
//line /usr/local/go/src/net/ip.go:686
			// _ = "end of CoverTab[6235]"
		} else {
//line /usr/local/go/src/net/ip.go:687
			_go_fuzz_dep_.CoverTab[6236]++
//line /usr/local/go/src/net/ip.go:687
			// _ = "end of CoverTab[6236]"
//line /usr/local/go/src/net/ip.go:687
		}
//line /usr/local/go/src/net/ip.go:687
		// _ = "end of CoverTab[6232]"
//line /usr/local/go/src/net/ip.go:687
		_go_fuzz_dep_.CoverTab[6233]++
						n := IPv6len - i
						for j := i - 1; j >= ellipsis; j-- {
//line /usr/local/go/src/net/ip.go:689
			_go_fuzz_dep_.CoverTab[6237]++
							ip[j+n] = ip[j]
//line /usr/local/go/src/net/ip.go:690
			// _ = "end of CoverTab[6237]"
		}
//line /usr/local/go/src/net/ip.go:691
		// _ = "end of CoverTab[6233]"
//line /usr/local/go/src/net/ip.go:691
		_go_fuzz_dep_.CoverTab[6234]++
						for j := ellipsis + n - 1; j >= ellipsis; j-- {
//line /usr/local/go/src/net/ip.go:692
			_go_fuzz_dep_.CoverTab[6238]++
							ip[j] = 0
//line /usr/local/go/src/net/ip.go:693
			// _ = "end of CoverTab[6238]"
		}
//line /usr/local/go/src/net/ip.go:694
		// _ = "end of CoverTab[6234]"
	} else {
//line /usr/local/go/src/net/ip.go:695
		_go_fuzz_dep_.CoverTab[6239]++
//line /usr/local/go/src/net/ip.go:695
		if ellipsis >= 0 {
//line /usr/local/go/src/net/ip.go:695
			_go_fuzz_dep_.CoverTab[6240]++

							return nil
//line /usr/local/go/src/net/ip.go:697
			// _ = "end of CoverTab[6240]"
		} else {
//line /usr/local/go/src/net/ip.go:698
			_go_fuzz_dep_.CoverTab[6241]++
//line /usr/local/go/src/net/ip.go:698
			// _ = "end of CoverTab[6241]"
//line /usr/local/go/src/net/ip.go:698
		}
//line /usr/local/go/src/net/ip.go:698
		// _ = "end of CoverTab[6239]"
//line /usr/local/go/src/net/ip.go:698
	}
//line /usr/local/go/src/net/ip.go:698
	// _ = "end of CoverTab[6189]"
//line /usr/local/go/src/net/ip.go:698
	_go_fuzz_dep_.CoverTab[6190]++
					return ip
//line /usr/local/go/src/net/ip.go:699
	// _ = "end of CoverTab[6190]"
}

// ParseIP parses s as an IP address, returning the result.
//line /usr/local/go/src/net/ip.go:702
// The string s can be in IPv4 dotted decimal ("192.0.2.1"), IPv6
//line /usr/local/go/src/net/ip.go:702
// ("2001:db8::68"), or IPv4-mapped IPv6 ("::ffff:192.0.2.1") form.
//line /usr/local/go/src/net/ip.go:702
// If s is not a valid textual representation of an IP address,
//line /usr/local/go/src/net/ip.go:702
// ParseIP returns nil.
//line /usr/local/go/src/net/ip.go:707
func ParseIP(s string) IP {
//line /usr/local/go/src/net/ip.go:707
	_go_fuzz_dep_.CoverTab[6242]++
					for i := 0; i < len(s); i++ {
//line /usr/local/go/src/net/ip.go:708
		_go_fuzz_dep_.CoverTab[6244]++
						switch s[i] {
		case '.':
//line /usr/local/go/src/net/ip.go:710
			_go_fuzz_dep_.CoverTab[6245]++
							return parseIPv4(s)
//line /usr/local/go/src/net/ip.go:711
			// _ = "end of CoverTab[6245]"
		case ':':
//line /usr/local/go/src/net/ip.go:712
			_go_fuzz_dep_.CoverTab[6246]++
							return parseIPv6(s)
//line /usr/local/go/src/net/ip.go:713
			// _ = "end of CoverTab[6246]"
//line /usr/local/go/src/net/ip.go:713
		default:
//line /usr/local/go/src/net/ip.go:713
			_go_fuzz_dep_.CoverTab[6247]++
//line /usr/local/go/src/net/ip.go:713
			// _ = "end of CoverTab[6247]"
		}
//line /usr/local/go/src/net/ip.go:714
		// _ = "end of CoverTab[6244]"
	}
//line /usr/local/go/src/net/ip.go:715
	// _ = "end of CoverTab[6242]"
//line /usr/local/go/src/net/ip.go:715
	_go_fuzz_dep_.CoverTab[6243]++
					return nil
//line /usr/local/go/src/net/ip.go:716
	// _ = "end of CoverTab[6243]"
}

// parseIPZone parses s as an IP address, return it and its associated zone
//line /usr/local/go/src/net/ip.go:719
// identifier (IPv6 only).
//line /usr/local/go/src/net/ip.go:721
func parseIPZone(s string) (IP, string) {
//line /usr/local/go/src/net/ip.go:721
	_go_fuzz_dep_.CoverTab[6248]++
					for i := 0; i < len(s); i++ {
//line /usr/local/go/src/net/ip.go:722
		_go_fuzz_dep_.CoverTab[6250]++
						switch s[i] {
		case '.':
//line /usr/local/go/src/net/ip.go:724
			_go_fuzz_dep_.CoverTab[6251]++
							return parseIPv4(s), ""
//line /usr/local/go/src/net/ip.go:725
			// _ = "end of CoverTab[6251]"
		case ':':
//line /usr/local/go/src/net/ip.go:726
			_go_fuzz_dep_.CoverTab[6252]++
							return parseIPv6Zone(s)
//line /usr/local/go/src/net/ip.go:727
			// _ = "end of CoverTab[6252]"
//line /usr/local/go/src/net/ip.go:727
		default:
//line /usr/local/go/src/net/ip.go:727
			_go_fuzz_dep_.CoverTab[6253]++
//line /usr/local/go/src/net/ip.go:727
			// _ = "end of CoverTab[6253]"
		}
//line /usr/local/go/src/net/ip.go:728
		// _ = "end of CoverTab[6250]"
	}
//line /usr/local/go/src/net/ip.go:729
	// _ = "end of CoverTab[6248]"
//line /usr/local/go/src/net/ip.go:729
	_go_fuzz_dep_.CoverTab[6249]++
					return nil, ""
//line /usr/local/go/src/net/ip.go:730
	// _ = "end of CoverTab[6249]"
}

// ParseCIDR parses s as a CIDR notation IP address and prefix length,
//line /usr/local/go/src/net/ip.go:733
// like "192.0.2.0/24" or "2001:db8::/32", as defined in
//line /usr/local/go/src/net/ip.go:733
// RFC 4632 and RFC 4291.
//line /usr/local/go/src/net/ip.go:733
//
//line /usr/local/go/src/net/ip.go:733
// It returns the IP address and the network implied by the IP and
//line /usr/local/go/src/net/ip.go:733
// prefix length.
//line /usr/local/go/src/net/ip.go:733
// For example, ParseCIDR("192.0.2.1/24") returns the IP address
//line /usr/local/go/src/net/ip.go:733
// 192.0.2.1 and the network 192.0.2.0/24.
//line /usr/local/go/src/net/ip.go:741
func ParseCIDR(s string) (IP, *IPNet, error) {
//line /usr/local/go/src/net/ip.go:741
	_go_fuzz_dep_.CoverTab[6254]++
					i := bytealg.IndexByteString(s, '/')
					if i < 0 {
//line /usr/local/go/src/net/ip.go:743
		_go_fuzz_dep_.CoverTab[6258]++
						return nil, nil, &ParseError{Type: "CIDR address", Text: s}
//line /usr/local/go/src/net/ip.go:744
		// _ = "end of CoverTab[6258]"
	} else {
//line /usr/local/go/src/net/ip.go:745
		_go_fuzz_dep_.CoverTab[6259]++
//line /usr/local/go/src/net/ip.go:745
		// _ = "end of CoverTab[6259]"
//line /usr/local/go/src/net/ip.go:745
	}
//line /usr/local/go/src/net/ip.go:745
	// _ = "end of CoverTab[6254]"
//line /usr/local/go/src/net/ip.go:745
	_go_fuzz_dep_.CoverTab[6255]++
					addr, mask := s[:i], s[i+1:]
					iplen := IPv4len
					ip := parseIPv4(addr)
					if ip == nil {
//line /usr/local/go/src/net/ip.go:749
		_go_fuzz_dep_.CoverTab[6260]++
						iplen = IPv6len
						ip = parseIPv6(addr)
//line /usr/local/go/src/net/ip.go:751
		// _ = "end of CoverTab[6260]"
	} else {
//line /usr/local/go/src/net/ip.go:752
		_go_fuzz_dep_.CoverTab[6261]++
//line /usr/local/go/src/net/ip.go:752
		// _ = "end of CoverTab[6261]"
//line /usr/local/go/src/net/ip.go:752
	}
//line /usr/local/go/src/net/ip.go:752
	// _ = "end of CoverTab[6255]"
//line /usr/local/go/src/net/ip.go:752
	_go_fuzz_dep_.CoverTab[6256]++
					n, i, ok := dtoi(mask)
					if ip == nil || func() bool {
//line /usr/local/go/src/net/ip.go:754
		_go_fuzz_dep_.CoverTab[6262]++
//line /usr/local/go/src/net/ip.go:754
		return !ok
//line /usr/local/go/src/net/ip.go:754
		// _ = "end of CoverTab[6262]"
//line /usr/local/go/src/net/ip.go:754
	}() || func() bool {
//line /usr/local/go/src/net/ip.go:754
		_go_fuzz_dep_.CoverTab[6263]++
//line /usr/local/go/src/net/ip.go:754
		return i != len(mask)
//line /usr/local/go/src/net/ip.go:754
		// _ = "end of CoverTab[6263]"
//line /usr/local/go/src/net/ip.go:754
	}() || func() bool {
//line /usr/local/go/src/net/ip.go:754
		_go_fuzz_dep_.CoverTab[6264]++
//line /usr/local/go/src/net/ip.go:754
		return n < 0
//line /usr/local/go/src/net/ip.go:754
		// _ = "end of CoverTab[6264]"
//line /usr/local/go/src/net/ip.go:754
	}() || func() bool {
//line /usr/local/go/src/net/ip.go:754
		_go_fuzz_dep_.CoverTab[6265]++
//line /usr/local/go/src/net/ip.go:754
		return n > 8*iplen
//line /usr/local/go/src/net/ip.go:754
		// _ = "end of CoverTab[6265]"
//line /usr/local/go/src/net/ip.go:754
	}() {
//line /usr/local/go/src/net/ip.go:754
		_go_fuzz_dep_.CoverTab[6266]++
						return nil, nil, &ParseError{Type: "CIDR address", Text: s}
//line /usr/local/go/src/net/ip.go:755
		// _ = "end of CoverTab[6266]"
	} else {
//line /usr/local/go/src/net/ip.go:756
		_go_fuzz_dep_.CoverTab[6267]++
//line /usr/local/go/src/net/ip.go:756
		// _ = "end of CoverTab[6267]"
//line /usr/local/go/src/net/ip.go:756
	}
//line /usr/local/go/src/net/ip.go:756
	// _ = "end of CoverTab[6256]"
//line /usr/local/go/src/net/ip.go:756
	_go_fuzz_dep_.CoverTab[6257]++
					m := CIDRMask(n, 8*iplen)
					return ip, &IPNet{IP: ip.Mask(m), Mask: m}, nil
//line /usr/local/go/src/net/ip.go:758
	// _ = "end of CoverTab[6257]"
}

func copyIP(x IP) IP {
//line /usr/local/go/src/net/ip.go:761
	_go_fuzz_dep_.CoverTab[6268]++
					y := make(IP, len(x))
					copy(y, x)
					return y
//line /usr/local/go/src/net/ip.go:764
	// _ = "end of CoverTab[6268]"
}

//line /usr/local/go/src/net/ip.go:765
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/ip.go:765
var _ = _go_fuzz_dep_.CoverTab
