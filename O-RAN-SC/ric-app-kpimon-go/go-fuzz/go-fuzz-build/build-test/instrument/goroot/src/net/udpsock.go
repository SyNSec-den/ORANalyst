// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /snap/go/10455/src/net/udpsock.go:5
package net

//line /snap/go/10455/src/net/udpsock.go:5
import (
//line /snap/go/10455/src/net/udpsock.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /snap/go/10455/src/net/udpsock.go:5
)
//line /snap/go/10455/src/net/udpsock.go:5
import (
//line /snap/go/10455/src/net/udpsock.go:5
	_atomic_ "sync/atomic"
//line /snap/go/10455/src/net/udpsock.go:5
)

import (
	"context"
	"internal/itoa"
	"net/netip"
	"syscall"
)

//line /snap/go/10455/src/net/udpsock.go:23
// UDPAddr represents the address of a UDP end point.
type UDPAddr struct {
	IP	IP
	Port	int
	Zone	string	// IPv6 scoped addressing zone
}

// AddrPort returns the UDPAddr a as a netip.AddrPort.
//line /snap/go/10455/src/net/udpsock.go:30
//
//line /snap/go/10455/src/net/udpsock.go:30
// If a.Port does not fit in a uint16, it's silently truncated.
//line /snap/go/10455/src/net/udpsock.go:30
//
//line /snap/go/10455/src/net/udpsock.go:30
// If a is nil, a zero value is returned.
//line /snap/go/10455/src/net/udpsock.go:35
func (a *UDPAddr) AddrPort() netip.AddrPort {
//line /snap/go/10455/src/net/udpsock.go:35
	_go_fuzz_dep_.CoverTab[8544]++
						if a == nil {
//line /snap/go/10455/src/net/udpsock.go:36
		_go_fuzz_dep_.CoverTab[529938]++
//line /snap/go/10455/src/net/udpsock.go:36
		_go_fuzz_dep_.CoverTab[8546]++
							return netip.AddrPort{}
//line /snap/go/10455/src/net/udpsock.go:37
		// _ = "end of CoverTab[8546]"
	} else {
//line /snap/go/10455/src/net/udpsock.go:38
		_go_fuzz_dep_.CoverTab[529939]++
//line /snap/go/10455/src/net/udpsock.go:38
		_go_fuzz_dep_.CoverTab[8547]++
//line /snap/go/10455/src/net/udpsock.go:38
		// _ = "end of CoverTab[8547]"
//line /snap/go/10455/src/net/udpsock.go:38
	}
//line /snap/go/10455/src/net/udpsock.go:38
	// _ = "end of CoverTab[8544]"
//line /snap/go/10455/src/net/udpsock.go:38
	_go_fuzz_dep_.CoverTab[8545]++
						na, _ := netip.AddrFromSlice(a.IP)
						na = na.WithZone(a.Zone)
						return netip.AddrPortFrom(na, uint16(a.Port))
//line /snap/go/10455/src/net/udpsock.go:41
	// _ = "end of CoverTab[8545]"
}

// Network returns the address's network name, "udp".
func (a *UDPAddr) Network() string {
//line /snap/go/10455/src/net/udpsock.go:45
	_go_fuzz_dep_.CoverTab[8548]++
//line /snap/go/10455/src/net/udpsock.go:45
	return "udp"
//line /snap/go/10455/src/net/udpsock.go:45
	// _ = "end of CoverTab[8548]"
//line /snap/go/10455/src/net/udpsock.go:45
}

func (a *UDPAddr) String() string {
//line /snap/go/10455/src/net/udpsock.go:47
	_go_fuzz_dep_.CoverTab[8549]++
						if a == nil {
//line /snap/go/10455/src/net/udpsock.go:48
		_go_fuzz_dep_.CoverTab[529940]++
//line /snap/go/10455/src/net/udpsock.go:48
		_go_fuzz_dep_.CoverTab[8552]++
							return "<nil>"
//line /snap/go/10455/src/net/udpsock.go:49
		// _ = "end of CoverTab[8552]"
	} else {
//line /snap/go/10455/src/net/udpsock.go:50
		_go_fuzz_dep_.CoverTab[529941]++
//line /snap/go/10455/src/net/udpsock.go:50
		_go_fuzz_dep_.CoverTab[8553]++
//line /snap/go/10455/src/net/udpsock.go:50
		// _ = "end of CoverTab[8553]"
//line /snap/go/10455/src/net/udpsock.go:50
	}
//line /snap/go/10455/src/net/udpsock.go:50
	// _ = "end of CoverTab[8549]"
//line /snap/go/10455/src/net/udpsock.go:50
	_go_fuzz_dep_.CoverTab[8550]++
						ip := ipEmptyString(a.IP)
						if a.Zone != "" {
//line /snap/go/10455/src/net/udpsock.go:52
		_go_fuzz_dep_.CoverTab[529942]++
//line /snap/go/10455/src/net/udpsock.go:52
		_go_fuzz_dep_.CoverTab[8554]++
							return JoinHostPort(ip+"%"+a.Zone, itoa.Itoa(a.Port))
//line /snap/go/10455/src/net/udpsock.go:53
		// _ = "end of CoverTab[8554]"
	} else {
//line /snap/go/10455/src/net/udpsock.go:54
		_go_fuzz_dep_.CoverTab[529943]++
//line /snap/go/10455/src/net/udpsock.go:54
		_go_fuzz_dep_.CoverTab[8555]++
//line /snap/go/10455/src/net/udpsock.go:54
		// _ = "end of CoverTab[8555]"
//line /snap/go/10455/src/net/udpsock.go:54
	}
//line /snap/go/10455/src/net/udpsock.go:54
	// _ = "end of CoverTab[8550]"
//line /snap/go/10455/src/net/udpsock.go:54
	_go_fuzz_dep_.CoverTab[8551]++
						return JoinHostPort(ip, itoa.Itoa(a.Port))
//line /snap/go/10455/src/net/udpsock.go:55
	// _ = "end of CoverTab[8551]"
}

func (a *UDPAddr) isWildcard() bool {
//line /snap/go/10455/src/net/udpsock.go:58
	_go_fuzz_dep_.CoverTab[8556]++
						if a == nil || func() bool {
//line /snap/go/10455/src/net/udpsock.go:59
		_go_fuzz_dep_.CoverTab[8558]++
//line /snap/go/10455/src/net/udpsock.go:59
		return a.IP == nil
//line /snap/go/10455/src/net/udpsock.go:59
		// _ = "end of CoverTab[8558]"
//line /snap/go/10455/src/net/udpsock.go:59
	}() {
//line /snap/go/10455/src/net/udpsock.go:59
		_go_fuzz_dep_.CoverTab[529944]++
//line /snap/go/10455/src/net/udpsock.go:59
		_go_fuzz_dep_.CoverTab[8559]++
							return true
//line /snap/go/10455/src/net/udpsock.go:60
		// _ = "end of CoverTab[8559]"
	} else {
//line /snap/go/10455/src/net/udpsock.go:61
		_go_fuzz_dep_.CoverTab[529945]++
//line /snap/go/10455/src/net/udpsock.go:61
		_go_fuzz_dep_.CoverTab[8560]++
//line /snap/go/10455/src/net/udpsock.go:61
		// _ = "end of CoverTab[8560]"
//line /snap/go/10455/src/net/udpsock.go:61
	}
//line /snap/go/10455/src/net/udpsock.go:61
	// _ = "end of CoverTab[8556]"
//line /snap/go/10455/src/net/udpsock.go:61
	_go_fuzz_dep_.CoverTab[8557]++
						return a.IP.IsUnspecified()
//line /snap/go/10455/src/net/udpsock.go:62
	// _ = "end of CoverTab[8557]"
}

func (a *UDPAddr) opAddr() Addr {
//line /snap/go/10455/src/net/udpsock.go:65
	_go_fuzz_dep_.CoverTab[8561]++
						if a == nil {
//line /snap/go/10455/src/net/udpsock.go:66
		_go_fuzz_dep_.CoverTab[529946]++
//line /snap/go/10455/src/net/udpsock.go:66
		_go_fuzz_dep_.CoverTab[8563]++
							return nil
//line /snap/go/10455/src/net/udpsock.go:67
		// _ = "end of CoverTab[8563]"
	} else {
//line /snap/go/10455/src/net/udpsock.go:68
		_go_fuzz_dep_.CoverTab[529947]++
//line /snap/go/10455/src/net/udpsock.go:68
		_go_fuzz_dep_.CoverTab[8564]++
//line /snap/go/10455/src/net/udpsock.go:68
		// _ = "end of CoverTab[8564]"
//line /snap/go/10455/src/net/udpsock.go:68
	}
//line /snap/go/10455/src/net/udpsock.go:68
	// _ = "end of CoverTab[8561]"
//line /snap/go/10455/src/net/udpsock.go:68
	_go_fuzz_dep_.CoverTab[8562]++
						return a
//line /snap/go/10455/src/net/udpsock.go:69
	// _ = "end of CoverTab[8562]"
}

// ResolveUDPAddr returns an address of UDP end point.
//line /snap/go/10455/src/net/udpsock.go:72
//
//line /snap/go/10455/src/net/udpsock.go:72
// The network must be a UDP network name.
//line /snap/go/10455/src/net/udpsock.go:72
//
//line /snap/go/10455/src/net/udpsock.go:72
// If the host in the address parameter is not a literal IP address or
//line /snap/go/10455/src/net/udpsock.go:72
// the port is not a literal port number, ResolveUDPAddr resolves the
//line /snap/go/10455/src/net/udpsock.go:72
// address to an address of UDP end point.
//line /snap/go/10455/src/net/udpsock.go:72
// Otherwise, it parses the address as a pair of literal IP address
//line /snap/go/10455/src/net/udpsock.go:72
// and port number.
//line /snap/go/10455/src/net/udpsock.go:72
// The address parameter can use a host name, but this is not
//line /snap/go/10455/src/net/udpsock.go:72
// recommended, because it will return at most one of the host name's
//line /snap/go/10455/src/net/udpsock.go:72
// IP addresses.
//line /snap/go/10455/src/net/udpsock.go:72
//
//line /snap/go/10455/src/net/udpsock.go:72
// See func Dial for a description of the network and address
//line /snap/go/10455/src/net/udpsock.go:72
// parameters.
//line /snap/go/10455/src/net/udpsock.go:87
func ResolveUDPAddr(network, address string) (*UDPAddr, error) {
//line /snap/go/10455/src/net/udpsock.go:87
	_go_fuzz_dep_.CoverTab[8565]++
						switch network {
	case "udp", "udp4", "udp6":
//line /snap/go/10455/src/net/udpsock.go:89
		_go_fuzz_dep_.CoverTab[529948]++
//line /snap/go/10455/src/net/udpsock.go:89
		_go_fuzz_dep_.CoverTab[8568]++
//line /snap/go/10455/src/net/udpsock.go:89
		// _ = "end of CoverTab[8568]"
	case "":
//line /snap/go/10455/src/net/udpsock.go:90
		_go_fuzz_dep_.CoverTab[529949]++
//line /snap/go/10455/src/net/udpsock.go:90
		_go_fuzz_dep_.CoverTab[8569]++
							network = "udp"
//line /snap/go/10455/src/net/udpsock.go:91
		// _ = "end of CoverTab[8569]"
	default:
//line /snap/go/10455/src/net/udpsock.go:92
		_go_fuzz_dep_.CoverTab[529950]++
//line /snap/go/10455/src/net/udpsock.go:92
		_go_fuzz_dep_.CoverTab[8570]++
							return nil, UnknownNetworkError(network)
//line /snap/go/10455/src/net/udpsock.go:93
		// _ = "end of CoverTab[8570]"
	}
//line /snap/go/10455/src/net/udpsock.go:94
	// _ = "end of CoverTab[8565]"
//line /snap/go/10455/src/net/udpsock.go:94
	_go_fuzz_dep_.CoverTab[8566]++
						addrs, err := DefaultResolver.internetAddrList(context.Background(), network, address)
						if err != nil {
//line /snap/go/10455/src/net/udpsock.go:96
		_go_fuzz_dep_.CoverTab[529951]++
//line /snap/go/10455/src/net/udpsock.go:96
		_go_fuzz_dep_.CoverTab[8571]++
							return nil, err
//line /snap/go/10455/src/net/udpsock.go:97
		// _ = "end of CoverTab[8571]"
	} else {
//line /snap/go/10455/src/net/udpsock.go:98
		_go_fuzz_dep_.CoverTab[529952]++
//line /snap/go/10455/src/net/udpsock.go:98
		_go_fuzz_dep_.CoverTab[8572]++
//line /snap/go/10455/src/net/udpsock.go:98
		// _ = "end of CoverTab[8572]"
//line /snap/go/10455/src/net/udpsock.go:98
	}
//line /snap/go/10455/src/net/udpsock.go:98
	// _ = "end of CoverTab[8566]"
//line /snap/go/10455/src/net/udpsock.go:98
	_go_fuzz_dep_.CoverTab[8567]++
						return addrs.forResolve(network, address).(*UDPAddr), nil
//line /snap/go/10455/src/net/udpsock.go:99
	// _ = "end of CoverTab[8567]"
}

// UDPAddrFromAddrPort returns addr as a UDPAddr. If addr.IsValid() is false,
//line /snap/go/10455/src/net/udpsock.go:102
// then the returned UDPAddr will contain a nil IP field, indicating an
//line /snap/go/10455/src/net/udpsock.go:102
// address family-agnostic unspecified address.
//line /snap/go/10455/src/net/udpsock.go:105
func UDPAddrFromAddrPort(addr netip.AddrPort) *UDPAddr {
//line /snap/go/10455/src/net/udpsock.go:105
	_go_fuzz_dep_.CoverTab[8573]++
						return &UDPAddr{
		IP:	addr.Addr().AsSlice(),
		Zone:	addr.Addr().Zone(),
		Port:	int(addr.Port()),
	}
//line /snap/go/10455/src/net/udpsock.go:110
	// _ = "end of CoverTab[8573]"
}

// An addrPortUDPAddr is a netip.AddrPort-based UDP address that satisfies the Addr interface.
type addrPortUDPAddr struct {
	netip.AddrPort
}

func (addrPortUDPAddr) Network() string {
//line /snap/go/10455/src/net/udpsock.go:118
	_go_fuzz_dep_.CoverTab[8574]++
//line /snap/go/10455/src/net/udpsock.go:118
	return "udp"
//line /snap/go/10455/src/net/udpsock.go:118
	// _ = "end of CoverTab[8574]"
//line /snap/go/10455/src/net/udpsock.go:118
}

// UDPConn is the implementation of the Conn and PacketConn interfaces
//line /snap/go/10455/src/net/udpsock.go:120
// for UDP network connections.
//line /snap/go/10455/src/net/udpsock.go:122
type UDPConn struct {
	conn
}

// SyscallConn returns a raw network connection.
//line /snap/go/10455/src/net/udpsock.go:126
// This implements the syscall.Conn interface.
//line /snap/go/10455/src/net/udpsock.go:128
func (c *UDPConn) SyscallConn() (syscall.RawConn, error) {
//line /snap/go/10455/src/net/udpsock.go:128
	_go_fuzz_dep_.CoverTab[8575]++
						if !c.ok() {
//line /snap/go/10455/src/net/udpsock.go:129
		_go_fuzz_dep_.CoverTab[529953]++
//line /snap/go/10455/src/net/udpsock.go:129
		_go_fuzz_dep_.CoverTab[8577]++
							return nil, syscall.EINVAL
//line /snap/go/10455/src/net/udpsock.go:130
		// _ = "end of CoverTab[8577]"
	} else {
//line /snap/go/10455/src/net/udpsock.go:131
		_go_fuzz_dep_.CoverTab[529954]++
//line /snap/go/10455/src/net/udpsock.go:131
		_go_fuzz_dep_.CoverTab[8578]++
//line /snap/go/10455/src/net/udpsock.go:131
		// _ = "end of CoverTab[8578]"
//line /snap/go/10455/src/net/udpsock.go:131
	}
//line /snap/go/10455/src/net/udpsock.go:131
	// _ = "end of CoverTab[8575]"
//line /snap/go/10455/src/net/udpsock.go:131
	_go_fuzz_dep_.CoverTab[8576]++
						return newRawConn(c.fd)
//line /snap/go/10455/src/net/udpsock.go:132
	// _ = "end of CoverTab[8576]"
}

// ReadFromUDP acts like ReadFrom but returns a UDPAddr.
func (c *UDPConn) ReadFromUDP(b []byte) (n int, addr *UDPAddr, err error) {
//line /snap/go/10455/src/net/udpsock.go:136
	_go_fuzz_dep_.CoverTab[8579]++

//line /snap/go/10455/src/net/udpsock.go:141
	return c.readFromUDP(b, &UDPAddr{})
//line /snap/go/10455/src/net/udpsock.go:141
	// _ = "end of CoverTab[8579]"
}

// readFromUDP implements ReadFromUDP.
func (c *UDPConn) readFromUDP(b []byte, addr *UDPAddr) (int, *UDPAddr, error) {
//line /snap/go/10455/src/net/udpsock.go:145
	_go_fuzz_dep_.CoverTab[8580]++
						if !c.ok() {
//line /snap/go/10455/src/net/udpsock.go:146
		_go_fuzz_dep_.CoverTab[529955]++
//line /snap/go/10455/src/net/udpsock.go:146
		_go_fuzz_dep_.CoverTab[8583]++
							return 0, nil, syscall.EINVAL
//line /snap/go/10455/src/net/udpsock.go:147
		// _ = "end of CoverTab[8583]"
	} else {
//line /snap/go/10455/src/net/udpsock.go:148
		_go_fuzz_dep_.CoverTab[529956]++
//line /snap/go/10455/src/net/udpsock.go:148
		_go_fuzz_dep_.CoverTab[8584]++
//line /snap/go/10455/src/net/udpsock.go:148
		// _ = "end of CoverTab[8584]"
//line /snap/go/10455/src/net/udpsock.go:148
	}
//line /snap/go/10455/src/net/udpsock.go:148
	// _ = "end of CoverTab[8580]"
//line /snap/go/10455/src/net/udpsock.go:148
	_go_fuzz_dep_.CoverTab[8581]++
						n, addr, err := c.readFrom(b, addr)
						if err != nil {
//line /snap/go/10455/src/net/udpsock.go:150
		_go_fuzz_dep_.CoverTab[529957]++
//line /snap/go/10455/src/net/udpsock.go:150
		_go_fuzz_dep_.CoverTab[8585]++
							err = &OpError{Op: "read", Net: c.fd.net, Source: c.fd.laddr, Addr: c.fd.raddr, Err: err}
//line /snap/go/10455/src/net/udpsock.go:151
		// _ = "end of CoverTab[8585]"
	} else {
//line /snap/go/10455/src/net/udpsock.go:152
		_go_fuzz_dep_.CoverTab[529958]++
//line /snap/go/10455/src/net/udpsock.go:152
		_go_fuzz_dep_.CoverTab[8586]++
//line /snap/go/10455/src/net/udpsock.go:152
		// _ = "end of CoverTab[8586]"
//line /snap/go/10455/src/net/udpsock.go:152
	}
//line /snap/go/10455/src/net/udpsock.go:152
	// _ = "end of CoverTab[8581]"
//line /snap/go/10455/src/net/udpsock.go:152
	_go_fuzz_dep_.CoverTab[8582]++
						return n, addr, err
//line /snap/go/10455/src/net/udpsock.go:153
	// _ = "end of CoverTab[8582]"
}

// ReadFrom implements the PacketConn ReadFrom method.
func (c *UDPConn) ReadFrom(b []byte) (int, Addr, error) {
//line /snap/go/10455/src/net/udpsock.go:157
	_go_fuzz_dep_.CoverTab[8587]++
						n, addr, err := c.readFromUDP(b, &UDPAddr{})
						if addr == nil {
//line /snap/go/10455/src/net/udpsock.go:159
		_go_fuzz_dep_.CoverTab[529959]++
//line /snap/go/10455/src/net/udpsock.go:159
		_go_fuzz_dep_.CoverTab[8589]++

							return n, nil, err
//line /snap/go/10455/src/net/udpsock.go:161
		// _ = "end of CoverTab[8589]"
	} else {
//line /snap/go/10455/src/net/udpsock.go:162
		_go_fuzz_dep_.CoverTab[529960]++
//line /snap/go/10455/src/net/udpsock.go:162
		_go_fuzz_dep_.CoverTab[8590]++
//line /snap/go/10455/src/net/udpsock.go:162
		// _ = "end of CoverTab[8590]"
//line /snap/go/10455/src/net/udpsock.go:162
	}
//line /snap/go/10455/src/net/udpsock.go:162
	// _ = "end of CoverTab[8587]"
//line /snap/go/10455/src/net/udpsock.go:162
	_go_fuzz_dep_.CoverTab[8588]++
						return n, addr, err
//line /snap/go/10455/src/net/udpsock.go:163
	// _ = "end of CoverTab[8588]"
}

// ReadFromUDPAddrPort acts like ReadFrom but returns a netip.AddrPort.
//line /snap/go/10455/src/net/udpsock.go:166
//
//line /snap/go/10455/src/net/udpsock.go:166
// If c is bound to an unspecified address, the returned
//line /snap/go/10455/src/net/udpsock.go:166
// netip.AddrPort's address might be an IPv4-mapped IPv6 address.
//line /snap/go/10455/src/net/udpsock.go:166
// Use netip.Addr.Unmap to get the address without the IPv6 prefix.
//line /snap/go/10455/src/net/udpsock.go:171
func (c *UDPConn) ReadFromUDPAddrPort(b []byte) (n int, addr netip.AddrPort, err error) {
//line /snap/go/10455/src/net/udpsock.go:171
	_go_fuzz_dep_.CoverTab[8591]++
						if !c.ok() {
//line /snap/go/10455/src/net/udpsock.go:172
		_go_fuzz_dep_.CoverTab[529961]++
//line /snap/go/10455/src/net/udpsock.go:172
		_go_fuzz_dep_.CoverTab[8594]++
							return 0, netip.AddrPort{}, syscall.EINVAL
//line /snap/go/10455/src/net/udpsock.go:173
		// _ = "end of CoverTab[8594]"
	} else {
//line /snap/go/10455/src/net/udpsock.go:174
		_go_fuzz_dep_.CoverTab[529962]++
//line /snap/go/10455/src/net/udpsock.go:174
		_go_fuzz_dep_.CoverTab[8595]++
//line /snap/go/10455/src/net/udpsock.go:174
		// _ = "end of CoverTab[8595]"
//line /snap/go/10455/src/net/udpsock.go:174
	}
//line /snap/go/10455/src/net/udpsock.go:174
	// _ = "end of CoverTab[8591]"
//line /snap/go/10455/src/net/udpsock.go:174
	_go_fuzz_dep_.CoverTab[8592]++
						n, addr, err = c.readFromAddrPort(b)
						if err != nil {
//line /snap/go/10455/src/net/udpsock.go:176
		_go_fuzz_dep_.CoverTab[529963]++
//line /snap/go/10455/src/net/udpsock.go:176
		_go_fuzz_dep_.CoverTab[8596]++
							err = &OpError{Op: "read", Net: c.fd.net, Source: c.fd.laddr, Addr: c.fd.raddr, Err: err}
//line /snap/go/10455/src/net/udpsock.go:177
		// _ = "end of CoverTab[8596]"
	} else {
//line /snap/go/10455/src/net/udpsock.go:178
		_go_fuzz_dep_.CoverTab[529964]++
//line /snap/go/10455/src/net/udpsock.go:178
		_go_fuzz_dep_.CoverTab[8597]++
//line /snap/go/10455/src/net/udpsock.go:178
		// _ = "end of CoverTab[8597]"
//line /snap/go/10455/src/net/udpsock.go:178
	}
//line /snap/go/10455/src/net/udpsock.go:178
	// _ = "end of CoverTab[8592]"
//line /snap/go/10455/src/net/udpsock.go:178
	_go_fuzz_dep_.CoverTab[8593]++
						return n, addr, err
//line /snap/go/10455/src/net/udpsock.go:179
	// _ = "end of CoverTab[8593]"
}

// ReadMsgUDP reads a message from c, copying the payload into b and
//line /snap/go/10455/src/net/udpsock.go:182
// the associated out-of-band data into oob. It returns the number of
//line /snap/go/10455/src/net/udpsock.go:182
// bytes copied into b, the number of bytes copied into oob, the flags
//line /snap/go/10455/src/net/udpsock.go:182
// that were set on the message and the source address of the message.
//line /snap/go/10455/src/net/udpsock.go:182
//
//line /snap/go/10455/src/net/udpsock.go:182
// The packages golang.org/x/net/ipv4 and golang.org/x/net/ipv6 can be
//line /snap/go/10455/src/net/udpsock.go:182
// used to manipulate IP-level socket options in oob.
//line /snap/go/10455/src/net/udpsock.go:189
func (c *UDPConn) ReadMsgUDP(b, oob []byte) (n, oobn, flags int, addr *UDPAddr, err error) {
//line /snap/go/10455/src/net/udpsock.go:189
	_go_fuzz_dep_.CoverTab[8598]++
						var ap netip.AddrPort
						n, oobn, flags, ap, err = c.ReadMsgUDPAddrPort(b, oob)
						if ap.IsValid() {
//line /snap/go/10455/src/net/udpsock.go:192
		_go_fuzz_dep_.CoverTab[529965]++
//line /snap/go/10455/src/net/udpsock.go:192
		_go_fuzz_dep_.CoverTab[8600]++
							addr = UDPAddrFromAddrPort(ap)
//line /snap/go/10455/src/net/udpsock.go:193
		// _ = "end of CoverTab[8600]"
	} else {
//line /snap/go/10455/src/net/udpsock.go:194
		_go_fuzz_dep_.CoverTab[529966]++
//line /snap/go/10455/src/net/udpsock.go:194
		_go_fuzz_dep_.CoverTab[8601]++
//line /snap/go/10455/src/net/udpsock.go:194
		// _ = "end of CoverTab[8601]"
//line /snap/go/10455/src/net/udpsock.go:194
	}
//line /snap/go/10455/src/net/udpsock.go:194
	// _ = "end of CoverTab[8598]"
//line /snap/go/10455/src/net/udpsock.go:194
	_go_fuzz_dep_.CoverTab[8599]++
						return
//line /snap/go/10455/src/net/udpsock.go:195
	// _ = "end of CoverTab[8599]"
}

// ReadMsgUDPAddrPort is like ReadMsgUDP but returns an netip.AddrPort instead of a UDPAddr.
func (c *UDPConn) ReadMsgUDPAddrPort(b, oob []byte) (n, oobn, flags int, addr netip.AddrPort, err error) {
//line /snap/go/10455/src/net/udpsock.go:199
	_go_fuzz_dep_.CoverTab[8602]++
						if !c.ok() {
//line /snap/go/10455/src/net/udpsock.go:200
		_go_fuzz_dep_.CoverTab[529967]++
//line /snap/go/10455/src/net/udpsock.go:200
		_go_fuzz_dep_.CoverTab[8605]++
							return 0, 0, 0, netip.AddrPort{}, syscall.EINVAL
//line /snap/go/10455/src/net/udpsock.go:201
		// _ = "end of CoverTab[8605]"
	} else {
//line /snap/go/10455/src/net/udpsock.go:202
		_go_fuzz_dep_.CoverTab[529968]++
//line /snap/go/10455/src/net/udpsock.go:202
		_go_fuzz_dep_.CoverTab[8606]++
//line /snap/go/10455/src/net/udpsock.go:202
		// _ = "end of CoverTab[8606]"
//line /snap/go/10455/src/net/udpsock.go:202
	}
//line /snap/go/10455/src/net/udpsock.go:202
	// _ = "end of CoverTab[8602]"
//line /snap/go/10455/src/net/udpsock.go:202
	_go_fuzz_dep_.CoverTab[8603]++
						n, oobn, flags, addr, err = c.readMsg(b, oob)
						if err != nil {
//line /snap/go/10455/src/net/udpsock.go:204
		_go_fuzz_dep_.CoverTab[529969]++
//line /snap/go/10455/src/net/udpsock.go:204
		_go_fuzz_dep_.CoverTab[8607]++
							err = &OpError{Op: "read", Net: c.fd.net, Source: c.fd.laddr, Addr: c.fd.raddr, Err: err}
//line /snap/go/10455/src/net/udpsock.go:205
		// _ = "end of CoverTab[8607]"
	} else {
//line /snap/go/10455/src/net/udpsock.go:206
		_go_fuzz_dep_.CoverTab[529970]++
//line /snap/go/10455/src/net/udpsock.go:206
		_go_fuzz_dep_.CoverTab[8608]++
//line /snap/go/10455/src/net/udpsock.go:206
		// _ = "end of CoverTab[8608]"
//line /snap/go/10455/src/net/udpsock.go:206
	}
//line /snap/go/10455/src/net/udpsock.go:206
	// _ = "end of CoverTab[8603]"
//line /snap/go/10455/src/net/udpsock.go:206
	_go_fuzz_dep_.CoverTab[8604]++
						return
//line /snap/go/10455/src/net/udpsock.go:207
	// _ = "end of CoverTab[8604]"
}

// WriteToUDP acts like WriteTo but takes a UDPAddr.
func (c *UDPConn) WriteToUDP(b []byte, addr *UDPAddr) (int, error) {
//line /snap/go/10455/src/net/udpsock.go:211
	_go_fuzz_dep_.CoverTab[8609]++
						if !c.ok() {
//line /snap/go/10455/src/net/udpsock.go:212
		_go_fuzz_dep_.CoverTab[529971]++
//line /snap/go/10455/src/net/udpsock.go:212
		_go_fuzz_dep_.CoverTab[8612]++
							return 0, syscall.EINVAL
//line /snap/go/10455/src/net/udpsock.go:213
		// _ = "end of CoverTab[8612]"
	} else {
//line /snap/go/10455/src/net/udpsock.go:214
		_go_fuzz_dep_.CoverTab[529972]++
//line /snap/go/10455/src/net/udpsock.go:214
		_go_fuzz_dep_.CoverTab[8613]++
//line /snap/go/10455/src/net/udpsock.go:214
		// _ = "end of CoverTab[8613]"
//line /snap/go/10455/src/net/udpsock.go:214
	}
//line /snap/go/10455/src/net/udpsock.go:214
	// _ = "end of CoverTab[8609]"
//line /snap/go/10455/src/net/udpsock.go:214
	_go_fuzz_dep_.CoverTab[8610]++
						n, err := c.writeTo(b, addr)
						if err != nil {
//line /snap/go/10455/src/net/udpsock.go:216
		_go_fuzz_dep_.CoverTab[529973]++
//line /snap/go/10455/src/net/udpsock.go:216
		_go_fuzz_dep_.CoverTab[8614]++
							err = &OpError{Op: "write", Net: c.fd.net, Source: c.fd.laddr, Addr: addr.opAddr(), Err: err}
//line /snap/go/10455/src/net/udpsock.go:217
		// _ = "end of CoverTab[8614]"
	} else {
//line /snap/go/10455/src/net/udpsock.go:218
		_go_fuzz_dep_.CoverTab[529974]++
//line /snap/go/10455/src/net/udpsock.go:218
		_go_fuzz_dep_.CoverTab[8615]++
//line /snap/go/10455/src/net/udpsock.go:218
		// _ = "end of CoverTab[8615]"
//line /snap/go/10455/src/net/udpsock.go:218
	}
//line /snap/go/10455/src/net/udpsock.go:218
	// _ = "end of CoverTab[8610]"
//line /snap/go/10455/src/net/udpsock.go:218
	_go_fuzz_dep_.CoverTab[8611]++
						return n, err
//line /snap/go/10455/src/net/udpsock.go:219
	// _ = "end of CoverTab[8611]"
}

// WriteToUDPAddrPort acts like WriteTo but takes a netip.AddrPort.
func (c *UDPConn) WriteToUDPAddrPort(b []byte, addr netip.AddrPort) (int, error) {
//line /snap/go/10455/src/net/udpsock.go:223
	_go_fuzz_dep_.CoverTab[8616]++
						if !c.ok() {
//line /snap/go/10455/src/net/udpsock.go:224
		_go_fuzz_dep_.CoverTab[529975]++
//line /snap/go/10455/src/net/udpsock.go:224
		_go_fuzz_dep_.CoverTab[8619]++
							return 0, syscall.EINVAL
//line /snap/go/10455/src/net/udpsock.go:225
		// _ = "end of CoverTab[8619]"
	} else {
//line /snap/go/10455/src/net/udpsock.go:226
		_go_fuzz_dep_.CoverTab[529976]++
//line /snap/go/10455/src/net/udpsock.go:226
		_go_fuzz_dep_.CoverTab[8620]++
//line /snap/go/10455/src/net/udpsock.go:226
		// _ = "end of CoverTab[8620]"
//line /snap/go/10455/src/net/udpsock.go:226
	}
//line /snap/go/10455/src/net/udpsock.go:226
	// _ = "end of CoverTab[8616]"
//line /snap/go/10455/src/net/udpsock.go:226
	_go_fuzz_dep_.CoverTab[8617]++
						n, err := c.writeToAddrPort(b, addr)
						if err != nil {
//line /snap/go/10455/src/net/udpsock.go:228
		_go_fuzz_dep_.CoverTab[529977]++
//line /snap/go/10455/src/net/udpsock.go:228
		_go_fuzz_dep_.CoverTab[8621]++
							err = &OpError{Op: "write", Net: c.fd.net, Source: c.fd.laddr, Addr: addrPortUDPAddr{addr}, Err: err}
//line /snap/go/10455/src/net/udpsock.go:229
		// _ = "end of CoverTab[8621]"
	} else {
//line /snap/go/10455/src/net/udpsock.go:230
		_go_fuzz_dep_.CoverTab[529978]++
//line /snap/go/10455/src/net/udpsock.go:230
		_go_fuzz_dep_.CoverTab[8622]++
//line /snap/go/10455/src/net/udpsock.go:230
		// _ = "end of CoverTab[8622]"
//line /snap/go/10455/src/net/udpsock.go:230
	}
//line /snap/go/10455/src/net/udpsock.go:230
	// _ = "end of CoverTab[8617]"
//line /snap/go/10455/src/net/udpsock.go:230
	_go_fuzz_dep_.CoverTab[8618]++
						return n, err
//line /snap/go/10455/src/net/udpsock.go:231
	// _ = "end of CoverTab[8618]"
}

// WriteTo implements the PacketConn WriteTo method.
func (c *UDPConn) WriteTo(b []byte, addr Addr) (int, error) {
//line /snap/go/10455/src/net/udpsock.go:235
	_go_fuzz_dep_.CoverTab[8623]++
						if !c.ok() {
//line /snap/go/10455/src/net/udpsock.go:236
		_go_fuzz_dep_.CoverTab[529979]++
//line /snap/go/10455/src/net/udpsock.go:236
		_go_fuzz_dep_.CoverTab[8627]++
							return 0, syscall.EINVAL
//line /snap/go/10455/src/net/udpsock.go:237
		// _ = "end of CoverTab[8627]"
	} else {
//line /snap/go/10455/src/net/udpsock.go:238
		_go_fuzz_dep_.CoverTab[529980]++
//line /snap/go/10455/src/net/udpsock.go:238
		_go_fuzz_dep_.CoverTab[8628]++
//line /snap/go/10455/src/net/udpsock.go:238
		// _ = "end of CoverTab[8628]"
//line /snap/go/10455/src/net/udpsock.go:238
	}
//line /snap/go/10455/src/net/udpsock.go:238
	// _ = "end of CoverTab[8623]"
//line /snap/go/10455/src/net/udpsock.go:238
	_go_fuzz_dep_.CoverTab[8624]++
						a, ok := addr.(*UDPAddr)
						if !ok {
//line /snap/go/10455/src/net/udpsock.go:240
		_go_fuzz_dep_.CoverTab[529981]++
//line /snap/go/10455/src/net/udpsock.go:240
		_go_fuzz_dep_.CoverTab[8629]++
							return 0, &OpError{Op: "write", Net: c.fd.net, Source: c.fd.laddr, Addr: addr, Err: syscall.EINVAL}
//line /snap/go/10455/src/net/udpsock.go:241
		// _ = "end of CoverTab[8629]"
	} else {
//line /snap/go/10455/src/net/udpsock.go:242
		_go_fuzz_dep_.CoverTab[529982]++
//line /snap/go/10455/src/net/udpsock.go:242
		_go_fuzz_dep_.CoverTab[8630]++
//line /snap/go/10455/src/net/udpsock.go:242
		// _ = "end of CoverTab[8630]"
//line /snap/go/10455/src/net/udpsock.go:242
	}
//line /snap/go/10455/src/net/udpsock.go:242
	// _ = "end of CoverTab[8624]"
//line /snap/go/10455/src/net/udpsock.go:242
	_go_fuzz_dep_.CoverTab[8625]++
						n, err := c.writeTo(b, a)
						if err != nil {
//line /snap/go/10455/src/net/udpsock.go:244
		_go_fuzz_dep_.CoverTab[529983]++
//line /snap/go/10455/src/net/udpsock.go:244
		_go_fuzz_dep_.CoverTab[8631]++
							err = &OpError{Op: "write", Net: c.fd.net, Source: c.fd.laddr, Addr: a.opAddr(), Err: err}
//line /snap/go/10455/src/net/udpsock.go:245
		// _ = "end of CoverTab[8631]"
	} else {
//line /snap/go/10455/src/net/udpsock.go:246
		_go_fuzz_dep_.CoverTab[529984]++
//line /snap/go/10455/src/net/udpsock.go:246
		_go_fuzz_dep_.CoverTab[8632]++
//line /snap/go/10455/src/net/udpsock.go:246
		// _ = "end of CoverTab[8632]"
//line /snap/go/10455/src/net/udpsock.go:246
	}
//line /snap/go/10455/src/net/udpsock.go:246
	// _ = "end of CoverTab[8625]"
//line /snap/go/10455/src/net/udpsock.go:246
	_go_fuzz_dep_.CoverTab[8626]++
						return n, err
//line /snap/go/10455/src/net/udpsock.go:247
	// _ = "end of CoverTab[8626]"
}

// WriteMsgUDP writes a message to addr via c if c isn't connected, or
//line /snap/go/10455/src/net/udpsock.go:250
// to c's remote address if c is connected (in which case addr must be
//line /snap/go/10455/src/net/udpsock.go:250
// nil). The payload is copied from b and the associated out-of-band
//line /snap/go/10455/src/net/udpsock.go:250
// data is copied from oob. It returns the number of payload and
//line /snap/go/10455/src/net/udpsock.go:250
// out-of-band bytes written.
//line /snap/go/10455/src/net/udpsock.go:250
//
//line /snap/go/10455/src/net/udpsock.go:250
// The packages golang.org/x/net/ipv4 and golang.org/x/net/ipv6 can be
//line /snap/go/10455/src/net/udpsock.go:250
// used to manipulate IP-level socket options in oob.
//line /snap/go/10455/src/net/udpsock.go:258
func (c *UDPConn) WriteMsgUDP(b, oob []byte, addr *UDPAddr) (n, oobn int, err error) {
//line /snap/go/10455/src/net/udpsock.go:258
	_go_fuzz_dep_.CoverTab[8633]++
						if !c.ok() {
//line /snap/go/10455/src/net/udpsock.go:259
		_go_fuzz_dep_.CoverTab[529985]++
//line /snap/go/10455/src/net/udpsock.go:259
		_go_fuzz_dep_.CoverTab[8636]++
							return 0, 0, syscall.EINVAL
//line /snap/go/10455/src/net/udpsock.go:260
		// _ = "end of CoverTab[8636]"
	} else {
//line /snap/go/10455/src/net/udpsock.go:261
		_go_fuzz_dep_.CoverTab[529986]++
//line /snap/go/10455/src/net/udpsock.go:261
		_go_fuzz_dep_.CoverTab[8637]++
//line /snap/go/10455/src/net/udpsock.go:261
		// _ = "end of CoverTab[8637]"
//line /snap/go/10455/src/net/udpsock.go:261
	}
//line /snap/go/10455/src/net/udpsock.go:261
	// _ = "end of CoverTab[8633]"
//line /snap/go/10455/src/net/udpsock.go:261
	_go_fuzz_dep_.CoverTab[8634]++
						n, oobn, err = c.writeMsg(b, oob, addr)
						if err != nil {
//line /snap/go/10455/src/net/udpsock.go:263
		_go_fuzz_dep_.CoverTab[529987]++
//line /snap/go/10455/src/net/udpsock.go:263
		_go_fuzz_dep_.CoverTab[8638]++
							err = &OpError{Op: "write", Net: c.fd.net, Source: c.fd.laddr, Addr: addr.opAddr(), Err: err}
//line /snap/go/10455/src/net/udpsock.go:264
		// _ = "end of CoverTab[8638]"
	} else {
//line /snap/go/10455/src/net/udpsock.go:265
		_go_fuzz_dep_.CoverTab[529988]++
//line /snap/go/10455/src/net/udpsock.go:265
		_go_fuzz_dep_.CoverTab[8639]++
//line /snap/go/10455/src/net/udpsock.go:265
		// _ = "end of CoverTab[8639]"
//line /snap/go/10455/src/net/udpsock.go:265
	}
//line /snap/go/10455/src/net/udpsock.go:265
	// _ = "end of CoverTab[8634]"
//line /snap/go/10455/src/net/udpsock.go:265
	_go_fuzz_dep_.CoverTab[8635]++
						return
//line /snap/go/10455/src/net/udpsock.go:266
	// _ = "end of CoverTab[8635]"
}

// WriteMsgUDPAddrPort is like WriteMsgUDP but takes a netip.AddrPort instead of a UDPAddr.
func (c *UDPConn) WriteMsgUDPAddrPort(b, oob []byte, addr netip.AddrPort) (n, oobn int, err error) {
//line /snap/go/10455/src/net/udpsock.go:270
	_go_fuzz_dep_.CoverTab[8640]++
						if !c.ok() {
//line /snap/go/10455/src/net/udpsock.go:271
		_go_fuzz_dep_.CoverTab[529989]++
//line /snap/go/10455/src/net/udpsock.go:271
		_go_fuzz_dep_.CoverTab[8643]++
							return 0, 0, syscall.EINVAL
//line /snap/go/10455/src/net/udpsock.go:272
		// _ = "end of CoverTab[8643]"
	} else {
//line /snap/go/10455/src/net/udpsock.go:273
		_go_fuzz_dep_.CoverTab[529990]++
//line /snap/go/10455/src/net/udpsock.go:273
		_go_fuzz_dep_.CoverTab[8644]++
//line /snap/go/10455/src/net/udpsock.go:273
		// _ = "end of CoverTab[8644]"
//line /snap/go/10455/src/net/udpsock.go:273
	}
//line /snap/go/10455/src/net/udpsock.go:273
	// _ = "end of CoverTab[8640]"
//line /snap/go/10455/src/net/udpsock.go:273
	_go_fuzz_dep_.CoverTab[8641]++
						n, oobn, err = c.writeMsgAddrPort(b, oob, addr)
						if err != nil {
//line /snap/go/10455/src/net/udpsock.go:275
		_go_fuzz_dep_.CoverTab[529991]++
//line /snap/go/10455/src/net/udpsock.go:275
		_go_fuzz_dep_.CoverTab[8645]++
							err = &OpError{Op: "write", Net: c.fd.net, Source: c.fd.laddr, Addr: addrPortUDPAddr{addr}, Err: err}
//line /snap/go/10455/src/net/udpsock.go:276
		// _ = "end of CoverTab[8645]"
	} else {
//line /snap/go/10455/src/net/udpsock.go:277
		_go_fuzz_dep_.CoverTab[529992]++
//line /snap/go/10455/src/net/udpsock.go:277
		_go_fuzz_dep_.CoverTab[8646]++
//line /snap/go/10455/src/net/udpsock.go:277
		// _ = "end of CoverTab[8646]"
//line /snap/go/10455/src/net/udpsock.go:277
	}
//line /snap/go/10455/src/net/udpsock.go:277
	// _ = "end of CoverTab[8641]"
//line /snap/go/10455/src/net/udpsock.go:277
	_go_fuzz_dep_.CoverTab[8642]++
						return
//line /snap/go/10455/src/net/udpsock.go:278
	// _ = "end of CoverTab[8642]"
}

func newUDPConn(fd *netFD) *UDPConn {
//line /snap/go/10455/src/net/udpsock.go:281
	_go_fuzz_dep_.CoverTab[8647]++
//line /snap/go/10455/src/net/udpsock.go:281
	return &UDPConn{conn{fd}}
//line /snap/go/10455/src/net/udpsock.go:281
	// _ = "end of CoverTab[8647]"
//line /snap/go/10455/src/net/udpsock.go:281
}

// DialUDP acts like Dial for UDP networks.
//line /snap/go/10455/src/net/udpsock.go:283
//
//line /snap/go/10455/src/net/udpsock.go:283
// The network must be a UDP network name; see func Dial for details.
//line /snap/go/10455/src/net/udpsock.go:283
//
//line /snap/go/10455/src/net/udpsock.go:283
// If laddr is nil, a local address is automatically chosen.
//line /snap/go/10455/src/net/udpsock.go:283
// If the IP field of raddr is nil or an unspecified IP address, the
//line /snap/go/10455/src/net/udpsock.go:283
// local system is assumed.
//line /snap/go/10455/src/net/udpsock.go:290
func DialUDP(network string, laddr, raddr *UDPAddr) (*UDPConn, error) {
//line /snap/go/10455/src/net/udpsock.go:290
	_go_fuzz_dep_.CoverTab[8648]++
						switch network {
	case "udp", "udp4", "udp6":
//line /snap/go/10455/src/net/udpsock.go:292
		_go_fuzz_dep_.CoverTab[529993]++
//line /snap/go/10455/src/net/udpsock.go:292
		_go_fuzz_dep_.CoverTab[8652]++
//line /snap/go/10455/src/net/udpsock.go:292
		// _ = "end of CoverTab[8652]"
	default:
//line /snap/go/10455/src/net/udpsock.go:293
		_go_fuzz_dep_.CoverTab[529994]++
//line /snap/go/10455/src/net/udpsock.go:293
		_go_fuzz_dep_.CoverTab[8653]++
							return nil, &OpError{Op: "dial", Net: network, Source: laddr.opAddr(), Addr: raddr.opAddr(), Err: UnknownNetworkError(network)}
//line /snap/go/10455/src/net/udpsock.go:294
		// _ = "end of CoverTab[8653]"
	}
//line /snap/go/10455/src/net/udpsock.go:295
	// _ = "end of CoverTab[8648]"
//line /snap/go/10455/src/net/udpsock.go:295
	_go_fuzz_dep_.CoverTab[8649]++
						if raddr == nil {
//line /snap/go/10455/src/net/udpsock.go:296
		_go_fuzz_dep_.CoverTab[529995]++
//line /snap/go/10455/src/net/udpsock.go:296
		_go_fuzz_dep_.CoverTab[8654]++
							return nil, &OpError{Op: "dial", Net: network, Source: laddr.opAddr(), Addr: nil, Err: errMissingAddress}
//line /snap/go/10455/src/net/udpsock.go:297
		// _ = "end of CoverTab[8654]"
	} else {
//line /snap/go/10455/src/net/udpsock.go:298
		_go_fuzz_dep_.CoverTab[529996]++
//line /snap/go/10455/src/net/udpsock.go:298
		_go_fuzz_dep_.CoverTab[8655]++
//line /snap/go/10455/src/net/udpsock.go:298
		// _ = "end of CoverTab[8655]"
//line /snap/go/10455/src/net/udpsock.go:298
	}
//line /snap/go/10455/src/net/udpsock.go:298
	// _ = "end of CoverTab[8649]"
//line /snap/go/10455/src/net/udpsock.go:298
	_go_fuzz_dep_.CoverTab[8650]++
						sd := &sysDialer{network: network, address: raddr.String()}
						c, err := sd.dialUDP(context.Background(), laddr, raddr)
						if err != nil {
//line /snap/go/10455/src/net/udpsock.go:301
		_go_fuzz_dep_.CoverTab[529997]++
//line /snap/go/10455/src/net/udpsock.go:301
		_go_fuzz_dep_.CoverTab[8656]++
							return nil, &OpError{Op: "dial", Net: network, Source: laddr.opAddr(), Addr: raddr.opAddr(), Err: err}
//line /snap/go/10455/src/net/udpsock.go:302
		// _ = "end of CoverTab[8656]"
	} else {
//line /snap/go/10455/src/net/udpsock.go:303
		_go_fuzz_dep_.CoverTab[529998]++
//line /snap/go/10455/src/net/udpsock.go:303
		_go_fuzz_dep_.CoverTab[8657]++
//line /snap/go/10455/src/net/udpsock.go:303
		// _ = "end of CoverTab[8657]"
//line /snap/go/10455/src/net/udpsock.go:303
	}
//line /snap/go/10455/src/net/udpsock.go:303
	// _ = "end of CoverTab[8650]"
//line /snap/go/10455/src/net/udpsock.go:303
	_go_fuzz_dep_.CoverTab[8651]++
						return c, nil
//line /snap/go/10455/src/net/udpsock.go:304
	// _ = "end of CoverTab[8651]"
}

// ListenUDP acts like ListenPacket for UDP networks.
//line /snap/go/10455/src/net/udpsock.go:307
//
//line /snap/go/10455/src/net/udpsock.go:307
// The network must be a UDP network name; see func Dial for details.
//line /snap/go/10455/src/net/udpsock.go:307
//
//line /snap/go/10455/src/net/udpsock.go:307
// If the IP field of laddr is nil or an unspecified IP address,
//line /snap/go/10455/src/net/udpsock.go:307
// ListenUDP listens on all available IP addresses of the local system
//line /snap/go/10455/src/net/udpsock.go:307
// except multicast IP addresses.
//line /snap/go/10455/src/net/udpsock.go:307
// If the Port field of laddr is 0, a port number is automatically
//line /snap/go/10455/src/net/udpsock.go:307
// chosen.
//line /snap/go/10455/src/net/udpsock.go:316
func ListenUDP(network string, laddr *UDPAddr) (*UDPConn, error) {
//line /snap/go/10455/src/net/udpsock.go:316
	_go_fuzz_dep_.CoverTab[8658]++
						switch network {
	case "udp", "udp4", "udp6":
//line /snap/go/10455/src/net/udpsock.go:318
		_go_fuzz_dep_.CoverTab[529999]++
//line /snap/go/10455/src/net/udpsock.go:318
		_go_fuzz_dep_.CoverTab[8662]++
//line /snap/go/10455/src/net/udpsock.go:318
		// _ = "end of CoverTab[8662]"
	default:
//line /snap/go/10455/src/net/udpsock.go:319
		_go_fuzz_dep_.CoverTab[530000]++
//line /snap/go/10455/src/net/udpsock.go:319
		_go_fuzz_dep_.CoverTab[8663]++
							return nil, &OpError{Op: "listen", Net: network, Source: nil, Addr: laddr.opAddr(), Err: UnknownNetworkError(network)}
//line /snap/go/10455/src/net/udpsock.go:320
		// _ = "end of CoverTab[8663]"
	}
//line /snap/go/10455/src/net/udpsock.go:321
	// _ = "end of CoverTab[8658]"
//line /snap/go/10455/src/net/udpsock.go:321
	_go_fuzz_dep_.CoverTab[8659]++
						if laddr == nil {
//line /snap/go/10455/src/net/udpsock.go:322
		_go_fuzz_dep_.CoverTab[530001]++
//line /snap/go/10455/src/net/udpsock.go:322
		_go_fuzz_dep_.CoverTab[8664]++
							laddr = &UDPAddr{}
//line /snap/go/10455/src/net/udpsock.go:323
		// _ = "end of CoverTab[8664]"
	} else {
//line /snap/go/10455/src/net/udpsock.go:324
		_go_fuzz_dep_.CoverTab[530002]++
//line /snap/go/10455/src/net/udpsock.go:324
		_go_fuzz_dep_.CoverTab[8665]++
//line /snap/go/10455/src/net/udpsock.go:324
		// _ = "end of CoverTab[8665]"
//line /snap/go/10455/src/net/udpsock.go:324
	}
//line /snap/go/10455/src/net/udpsock.go:324
	// _ = "end of CoverTab[8659]"
//line /snap/go/10455/src/net/udpsock.go:324
	_go_fuzz_dep_.CoverTab[8660]++
						sl := &sysListener{network: network, address: laddr.String()}
						c, err := sl.listenUDP(context.Background(), laddr)
						if err != nil {
//line /snap/go/10455/src/net/udpsock.go:327
		_go_fuzz_dep_.CoverTab[530003]++
//line /snap/go/10455/src/net/udpsock.go:327
		_go_fuzz_dep_.CoverTab[8666]++
							return nil, &OpError{Op: "listen", Net: network, Source: nil, Addr: laddr.opAddr(), Err: err}
//line /snap/go/10455/src/net/udpsock.go:328
		// _ = "end of CoverTab[8666]"
	} else {
//line /snap/go/10455/src/net/udpsock.go:329
		_go_fuzz_dep_.CoverTab[530004]++
//line /snap/go/10455/src/net/udpsock.go:329
		_go_fuzz_dep_.CoverTab[8667]++
//line /snap/go/10455/src/net/udpsock.go:329
		// _ = "end of CoverTab[8667]"
//line /snap/go/10455/src/net/udpsock.go:329
	}
//line /snap/go/10455/src/net/udpsock.go:329
	// _ = "end of CoverTab[8660]"
//line /snap/go/10455/src/net/udpsock.go:329
	_go_fuzz_dep_.CoverTab[8661]++
						return c, nil
//line /snap/go/10455/src/net/udpsock.go:330
	// _ = "end of CoverTab[8661]"
}

// ListenMulticastUDP acts like ListenPacket for UDP networks but
//line /snap/go/10455/src/net/udpsock.go:333
// takes a group address on a specific network interface.
//line /snap/go/10455/src/net/udpsock.go:333
//
//line /snap/go/10455/src/net/udpsock.go:333
// The network must be a UDP network name; see func Dial for details.
//line /snap/go/10455/src/net/udpsock.go:333
//
//line /snap/go/10455/src/net/udpsock.go:333
// ListenMulticastUDP listens on all available IP addresses of the
//line /snap/go/10455/src/net/udpsock.go:333
// local system including the group, multicast IP address.
//line /snap/go/10455/src/net/udpsock.go:333
// If ifi is nil, ListenMulticastUDP uses the system-assigned
//line /snap/go/10455/src/net/udpsock.go:333
// multicast interface, although this is not recommended because the
//line /snap/go/10455/src/net/udpsock.go:333
// assignment depends on platforms and sometimes it might require
//line /snap/go/10455/src/net/udpsock.go:333
// routing configuration.
//line /snap/go/10455/src/net/udpsock.go:333
// If the Port field of gaddr is 0, a port number is automatically
//line /snap/go/10455/src/net/udpsock.go:333
// chosen.
//line /snap/go/10455/src/net/udpsock.go:333
//
//line /snap/go/10455/src/net/udpsock.go:333
// ListenMulticastUDP is just for convenience of simple, small
//line /snap/go/10455/src/net/udpsock.go:333
// applications. There are golang.org/x/net/ipv4 and
//line /snap/go/10455/src/net/udpsock.go:333
// golang.org/x/net/ipv6 packages for general purpose uses.
//line /snap/go/10455/src/net/udpsock.go:333
//
//line /snap/go/10455/src/net/udpsock.go:333
// Note that ListenMulticastUDP will set the IP_MULTICAST_LOOP socket option
//line /snap/go/10455/src/net/udpsock.go:333
// to 0 under IPPROTO_IP, to disable loopback of multicast packets.
//line /snap/go/10455/src/net/udpsock.go:353
func ListenMulticastUDP(network string, ifi *Interface, gaddr *UDPAddr) (*UDPConn, error) {
//line /snap/go/10455/src/net/udpsock.go:353
	_go_fuzz_dep_.CoverTab[8668]++
						switch network {
	case "udp", "udp4", "udp6":
//line /snap/go/10455/src/net/udpsock.go:355
		_go_fuzz_dep_.CoverTab[530005]++
//line /snap/go/10455/src/net/udpsock.go:355
		_go_fuzz_dep_.CoverTab[8672]++
//line /snap/go/10455/src/net/udpsock.go:355
		// _ = "end of CoverTab[8672]"
	default:
//line /snap/go/10455/src/net/udpsock.go:356
		_go_fuzz_dep_.CoverTab[530006]++
//line /snap/go/10455/src/net/udpsock.go:356
		_go_fuzz_dep_.CoverTab[8673]++
							return nil, &OpError{Op: "listen", Net: network, Source: nil, Addr: gaddr.opAddr(), Err: UnknownNetworkError(network)}
//line /snap/go/10455/src/net/udpsock.go:357
		// _ = "end of CoverTab[8673]"
	}
//line /snap/go/10455/src/net/udpsock.go:358
	// _ = "end of CoverTab[8668]"
//line /snap/go/10455/src/net/udpsock.go:358
	_go_fuzz_dep_.CoverTab[8669]++
						if gaddr == nil || func() bool {
//line /snap/go/10455/src/net/udpsock.go:359
		_go_fuzz_dep_.CoverTab[8674]++
//line /snap/go/10455/src/net/udpsock.go:359
		return gaddr.IP == nil
//line /snap/go/10455/src/net/udpsock.go:359
		// _ = "end of CoverTab[8674]"
//line /snap/go/10455/src/net/udpsock.go:359
	}() {
//line /snap/go/10455/src/net/udpsock.go:359
		_go_fuzz_dep_.CoverTab[530007]++
//line /snap/go/10455/src/net/udpsock.go:359
		_go_fuzz_dep_.CoverTab[8675]++
							return nil, &OpError{Op: "listen", Net: network, Source: nil, Addr: gaddr.opAddr(), Err: errMissingAddress}
//line /snap/go/10455/src/net/udpsock.go:360
		// _ = "end of CoverTab[8675]"
	} else {
//line /snap/go/10455/src/net/udpsock.go:361
		_go_fuzz_dep_.CoverTab[530008]++
//line /snap/go/10455/src/net/udpsock.go:361
		_go_fuzz_dep_.CoverTab[8676]++
//line /snap/go/10455/src/net/udpsock.go:361
		// _ = "end of CoverTab[8676]"
//line /snap/go/10455/src/net/udpsock.go:361
	}
//line /snap/go/10455/src/net/udpsock.go:361
	// _ = "end of CoverTab[8669]"
//line /snap/go/10455/src/net/udpsock.go:361
	_go_fuzz_dep_.CoverTab[8670]++
						sl := &sysListener{network: network, address: gaddr.String()}
						c, err := sl.listenMulticastUDP(context.Background(), ifi, gaddr)
						if err != nil {
//line /snap/go/10455/src/net/udpsock.go:364
		_go_fuzz_dep_.CoverTab[530009]++
//line /snap/go/10455/src/net/udpsock.go:364
		_go_fuzz_dep_.CoverTab[8677]++
							return nil, &OpError{Op: "listen", Net: network, Source: nil, Addr: gaddr.opAddr(), Err: err}
//line /snap/go/10455/src/net/udpsock.go:365
		// _ = "end of CoverTab[8677]"
	} else {
//line /snap/go/10455/src/net/udpsock.go:366
		_go_fuzz_dep_.CoverTab[530010]++
//line /snap/go/10455/src/net/udpsock.go:366
		_go_fuzz_dep_.CoverTab[8678]++
//line /snap/go/10455/src/net/udpsock.go:366
		// _ = "end of CoverTab[8678]"
//line /snap/go/10455/src/net/udpsock.go:366
	}
//line /snap/go/10455/src/net/udpsock.go:366
	// _ = "end of CoverTab[8670]"
//line /snap/go/10455/src/net/udpsock.go:366
	_go_fuzz_dep_.CoverTab[8671]++
						return c, nil
//line /snap/go/10455/src/net/udpsock.go:367
	// _ = "end of CoverTab[8671]"
}

//line /snap/go/10455/src/net/udpsock.go:368
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /snap/go/10455/src/net/udpsock.go:368
var _ = _go_fuzz_dep_.CoverTab
