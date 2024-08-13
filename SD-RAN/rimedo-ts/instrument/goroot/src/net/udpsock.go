// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/net/udpsock.go:5
package net

//line /usr/local/go/src/net/udpsock.go:5
import (
//line /usr/local/go/src/net/udpsock.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/udpsock.go:5
)
//line /usr/local/go/src/net/udpsock.go:5
import (
//line /usr/local/go/src/net/udpsock.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/udpsock.go:5
)

import (
	"context"
	"internal/itoa"
	"net/netip"
	"syscall"
)

//line /usr/local/go/src/net/udpsock.go:23
// UDPAddr represents the address of a UDP end point.
type UDPAddr struct {
	IP	IP
	Port	int
	Zone	string	// IPv6 scoped addressing zone
}

// AddrPort returns the UDPAddr a as a netip.AddrPort.
//line /usr/local/go/src/net/udpsock.go:30
//
//line /usr/local/go/src/net/udpsock.go:30
// If a.Port does not fit in a uint16, it's silently truncated.
//line /usr/local/go/src/net/udpsock.go:30
//
//line /usr/local/go/src/net/udpsock.go:30
// If a is nil, a zero value is returned.
//line /usr/local/go/src/net/udpsock.go:35
func (a *UDPAddr) AddrPort() netip.AddrPort {
//line /usr/local/go/src/net/udpsock.go:35
	_go_fuzz_dep_.CoverTab[16634]++
						if a == nil {
//line /usr/local/go/src/net/udpsock.go:36
		_go_fuzz_dep_.CoverTab[16636]++
							return netip.AddrPort{}
//line /usr/local/go/src/net/udpsock.go:37
		// _ = "end of CoverTab[16636]"
	} else {
//line /usr/local/go/src/net/udpsock.go:38
		_go_fuzz_dep_.CoverTab[16637]++
//line /usr/local/go/src/net/udpsock.go:38
		// _ = "end of CoverTab[16637]"
//line /usr/local/go/src/net/udpsock.go:38
	}
//line /usr/local/go/src/net/udpsock.go:38
	// _ = "end of CoverTab[16634]"
//line /usr/local/go/src/net/udpsock.go:38
	_go_fuzz_dep_.CoverTab[16635]++
						na, _ := netip.AddrFromSlice(a.IP)
						na = na.WithZone(a.Zone)
						return netip.AddrPortFrom(na, uint16(a.Port))
//line /usr/local/go/src/net/udpsock.go:41
	// _ = "end of CoverTab[16635]"
}

// Network returns the address's network name, "udp".
func (a *UDPAddr) Network() string {
//line /usr/local/go/src/net/udpsock.go:45
	_go_fuzz_dep_.CoverTab[16638]++
//line /usr/local/go/src/net/udpsock.go:45
	return "udp"
//line /usr/local/go/src/net/udpsock.go:45
	// _ = "end of CoverTab[16638]"
//line /usr/local/go/src/net/udpsock.go:45
}

func (a *UDPAddr) String() string {
//line /usr/local/go/src/net/udpsock.go:47
	_go_fuzz_dep_.CoverTab[16639]++
						if a == nil {
//line /usr/local/go/src/net/udpsock.go:48
		_go_fuzz_dep_.CoverTab[16642]++
							return "<nil>"
//line /usr/local/go/src/net/udpsock.go:49
		// _ = "end of CoverTab[16642]"
	} else {
//line /usr/local/go/src/net/udpsock.go:50
		_go_fuzz_dep_.CoverTab[16643]++
//line /usr/local/go/src/net/udpsock.go:50
		// _ = "end of CoverTab[16643]"
//line /usr/local/go/src/net/udpsock.go:50
	}
//line /usr/local/go/src/net/udpsock.go:50
	// _ = "end of CoverTab[16639]"
//line /usr/local/go/src/net/udpsock.go:50
	_go_fuzz_dep_.CoverTab[16640]++
						ip := ipEmptyString(a.IP)
						if a.Zone != "" {
//line /usr/local/go/src/net/udpsock.go:52
		_go_fuzz_dep_.CoverTab[16644]++
							return JoinHostPort(ip+"%"+a.Zone, itoa.Itoa(a.Port))
//line /usr/local/go/src/net/udpsock.go:53
		// _ = "end of CoverTab[16644]"
	} else {
//line /usr/local/go/src/net/udpsock.go:54
		_go_fuzz_dep_.CoverTab[16645]++
//line /usr/local/go/src/net/udpsock.go:54
		// _ = "end of CoverTab[16645]"
//line /usr/local/go/src/net/udpsock.go:54
	}
//line /usr/local/go/src/net/udpsock.go:54
	// _ = "end of CoverTab[16640]"
//line /usr/local/go/src/net/udpsock.go:54
	_go_fuzz_dep_.CoverTab[16641]++
						return JoinHostPort(ip, itoa.Itoa(a.Port))
//line /usr/local/go/src/net/udpsock.go:55
	// _ = "end of CoverTab[16641]"
}

func (a *UDPAddr) isWildcard() bool {
//line /usr/local/go/src/net/udpsock.go:58
	_go_fuzz_dep_.CoverTab[16646]++
						if a == nil || func() bool {
//line /usr/local/go/src/net/udpsock.go:59
		_go_fuzz_dep_.CoverTab[16648]++
//line /usr/local/go/src/net/udpsock.go:59
		return a.IP == nil
//line /usr/local/go/src/net/udpsock.go:59
		// _ = "end of CoverTab[16648]"
//line /usr/local/go/src/net/udpsock.go:59
	}() {
//line /usr/local/go/src/net/udpsock.go:59
		_go_fuzz_dep_.CoverTab[16649]++
							return true
//line /usr/local/go/src/net/udpsock.go:60
		// _ = "end of CoverTab[16649]"
	} else {
//line /usr/local/go/src/net/udpsock.go:61
		_go_fuzz_dep_.CoverTab[16650]++
//line /usr/local/go/src/net/udpsock.go:61
		// _ = "end of CoverTab[16650]"
//line /usr/local/go/src/net/udpsock.go:61
	}
//line /usr/local/go/src/net/udpsock.go:61
	// _ = "end of CoverTab[16646]"
//line /usr/local/go/src/net/udpsock.go:61
	_go_fuzz_dep_.CoverTab[16647]++
						return a.IP.IsUnspecified()
//line /usr/local/go/src/net/udpsock.go:62
	// _ = "end of CoverTab[16647]"
}

func (a *UDPAddr) opAddr() Addr {
//line /usr/local/go/src/net/udpsock.go:65
	_go_fuzz_dep_.CoverTab[16651]++
						if a == nil {
//line /usr/local/go/src/net/udpsock.go:66
		_go_fuzz_dep_.CoverTab[16653]++
							return nil
//line /usr/local/go/src/net/udpsock.go:67
		// _ = "end of CoverTab[16653]"
	} else {
//line /usr/local/go/src/net/udpsock.go:68
		_go_fuzz_dep_.CoverTab[16654]++
//line /usr/local/go/src/net/udpsock.go:68
		// _ = "end of CoverTab[16654]"
//line /usr/local/go/src/net/udpsock.go:68
	}
//line /usr/local/go/src/net/udpsock.go:68
	// _ = "end of CoverTab[16651]"
//line /usr/local/go/src/net/udpsock.go:68
	_go_fuzz_dep_.CoverTab[16652]++
						return a
//line /usr/local/go/src/net/udpsock.go:69
	// _ = "end of CoverTab[16652]"
}

// ResolveUDPAddr returns an address of UDP end point.
//line /usr/local/go/src/net/udpsock.go:72
//
//line /usr/local/go/src/net/udpsock.go:72
// The network must be a UDP network name.
//line /usr/local/go/src/net/udpsock.go:72
//
//line /usr/local/go/src/net/udpsock.go:72
// If the host in the address parameter is not a literal IP address or
//line /usr/local/go/src/net/udpsock.go:72
// the port is not a literal port number, ResolveUDPAddr resolves the
//line /usr/local/go/src/net/udpsock.go:72
// address to an address of UDP end point.
//line /usr/local/go/src/net/udpsock.go:72
// Otherwise, it parses the address as a pair of literal IP address
//line /usr/local/go/src/net/udpsock.go:72
// and port number.
//line /usr/local/go/src/net/udpsock.go:72
// The address parameter can use a host name, but this is not
//line /usr/local/go/src/net/udpsock.go:72
// recommended, because it will return at most one of the host name's
//line /usr/local/go/src/net/udpsock.go:72
// IP addresses.
//line /usr/local/go/src/net/udpsock.go:72
//
//line /usr/local/go/src/net/udpsock.go:72
// See func Dial for a description of the network and address
//line /usr/local/go/src/net/udpsock.go:72
// parameters.
//line /usr/local/go/src/net/udpsock.go:87
func ResolveUDPAddr(network, address string) (*UDPAddr, error) {
//line /usr/local/go/src/net/udpsock.go:87
	_go_fuzz_dep_.CoverTab[16655]++
						switch network {
	case "udp", "udp4", "udp6":
//line /usr/local/go/src/net/udpsock.go:89
		_go_fuzz_dep_.CoverTab[16658]++
//line /usr/local/go/src/net/udpsock.go:89
		// _ = "end of CoverTab[16658]"
	case "":
//line /usr/local/go/src/net/udpsock.go:90
		_go_fuzz_dep_.CoverTab[16659]++
							network = "udp"
//line /usr/local/go/src/net/udpsock.go:91
		// _ = "end of CoverTab[16659]"
	default:
//line /usr/local/go/src/net/udpsock.go:92
		_go_fuzz_dep_.CoverTab[16660]++
							return nil, UnknownNetworkError(network)
//line /usr/local/go/src/net/udpsock.go:93
		// _ = "end of CoverTab[16660]"
	}
//line /usr/local/go/src/net/udpsock.go:94
	// _ = "end of CoverTab[16655]"
//line /usr/local/go/src/net/udpsock.go:94
	_go_fuzz_dep_.CoverTab[16656]++
						addrs, err := DefaultResolver.internetAddrList(context.Background(), network, address)
						if err != nil {
//line /usr/local/go/src/net/udpsock.go:96
		_go_fuzz_dep_.CoverTab[16661]++
							return nil, err
//line /usr/local/go/src/net/udpsock.go:97
		// _ = "end of CoverTab[16661]"
	} else {
//line /usr/local/go/src/net/udpsock.go:98
		_go_fuzz_dep_.CoverTab[16662]++
//line /usr/local/go/src/net/udpsock.go:98
		// _ = "end of CoverTab[16662]"
//line /usr/local/go/src/net/udpsock.go:98
	}
//line /usr/local/go/src/net/udpsock.go:98
	// _ = "end of CoverTab[16656]"
//line /usr/local/go/src/net/udpsock.go:98
	_go_fuzz_dep_.CoverTab[16657]++
						return addrs.forResolve(network, address).(*UDPAddr), nil
//line /usr/local/go/src/net/udpsock.go:99
	// _ = "end of CoverTab[16657]"
}

// UDPAddrFromAddrPort returns addr as a UDPAddr. If addr.IsValid() is false,
//line /usr/local/go/src/net/udpsock.go:102
// then the returned UDPAddr will contain a nil IP field, indicating an
//line /usr/local/go/src/net/udpsock.go:102
// address family-agnostic unspecified address.
//line /usr/local/go/src/net/udpsock.go:105
func UDPAddrFromAddrPort(addr netip.AddrPort) *UDPAddr {
//line /usr/local/go/src/net/udpsock.go:105
	_go_fuzz_dep_.CoverTab[16663]++
						return &UDPAddr{
		IP:	addr.Addr().AsSlice(),
		Zone:	addr.Addr().Zone(),
		Port:	int(addr.Port()),
	}
//line /usr/local/go/src/net/udpsock.go:110
	// _ = "end of CoverTab[16663]"
}

// An addrPortUDPAddr is a netip.AddrPort-based UDP address that satisfies the Addr interface.
type addrPortUDPAddr struct {
	netip.AddrPort
}

func (addrPortUDPAddr) Network() string {
//line /usr/local/go/src/net/udpsock.go:118
	_go_fuzz_dep_.CoverTab[16664]++
//line /usr/local/go/src/net/udpsock.go:118
	return "udp"
//line /usr/local/go/src/net/udpsock.go:118
	// _ = "end of CoverTab[16664]"
//line /usr/local/go/src/net/udpsock.go:118
}

// UDPConn is the implementation of the Conn and PacketConn interfaces
//line /usr/local/go/src/net/udpsock.go:120
// for UDP network connections.
//line /usr/local/go/src/net/udpsock.go:122
type UDPConn struct {
	conn
}

// SyscallConn returns a raw network connection.
//line /usr/local/go/src/net/udpsock.go:126
// This implements the syscall.Conn interface.
//line /usr/local/go/src/net/udpsock.go:128
func (c *UDPConn) SyscallConn() (syscall.RawConn, error) {
//line /usr/local/go/src/net/udpsock.go:128
	_go_fuzz_dep_.CoverTab[16665]++
						if !c.ok() {
//line /usr/local/go/src/net/udpsock.go:129
		_go_fuzz_dep_.CoverTab[16667]++
							return nil, syscall.EINVAL
//line /usr/local/go/src/net/udpsock.go:130
		// _ = "end of CoverTab[16667]"
	} else {
//line /usr/local/go/src/net/udpsock.go:131
		_go_fuzz_dep_.CoverTab[16668]++
//line /usr/local/go/src/net/udpsock.go:131
		// _ = "end of CoverTab[16668]"
//line /usr/local/go/src/net/udpsock.go:131
	}
//line /usr/local/go/src/net/udpsock.go:131
	// _ = "end of CoverTab[16665]"
//line /usr/local/go/src/net/udpsock.go:131
	_go_fuzz_dep_.CoverTab[16666]++
						return newRawConn(c.fd)
//line /usr/local/go/src/net/udpsock.go:132
	// _ = "end of CoverTab[16666]"
}

// ReadFromUDP acts like ReadFrom but returns a UDPAddr.
func (c *UDPConn) ReadFromUDP(b []byte) (n int, addr *UDPAddr, err error) {
//line /usr/local/go/src/net/udpsock.go:136
	_go_fuzz_dep_.CoverTab[16669]++

//line /usr/local/go/src/net/udpsock.go:141
	return c.readFromUDP(b, &UDPAddr{})
//line /usr/local/go/src/net/udpsock.go:141
	// _ = "end of CoverTab[16669]"
}

// readFromUDP implements ReadFromUDP.
func (c *UDPConn) readFromUDP(b []byte, addr *UDPAddr) (int, *UDPAddr, error) {
//line /usr/local/go/src/net/udpsock.go:145
	_go_fuzz_dep_.CoverTab[16670]++
						if !c.ok() {
//line /usr/local/go/src/net/udpsock.go:146
		_go_fuzz_dep_.CoverTab[16673]++
							return 0, nil, syscall.EINVAL
//line /usr/local/go/src/net/udpsock.go:147
		// _ = "end of CoverTab[16673]"
	} else {
//line /usr/local/go/src/net/udpsock.go:148
		_go_fuzz_dep_.CoverTab[16674]++
//line /usr/local/go/src/net/udpsock.go:148
		// _ = "end of CoverTab[16674]"
//line /usr/local/go/src/net/udpsock.go:148
	}
//line /usr/local/go/src/net/udpsock.go:148
	// _ = "end of CoverTab[16670]"
//line /usr/local/go/src/net/udpsock.go:148
	_go_fuzz_dep_.CoverTab[16671]++
						n, addr, err := c.readFrom(b, addr)
						if err != nil {
//line /usr/local/go/src/net/udpsock.go:150
		_go_fuzz_dep_.CoverTab[16675]++
							err = &OpError{Op: "read", Net: c.fd.net, Source: c.fd.laddr, Addr: c.fd.raddr, Err: err}
//line /usr/local/go/src/net/udpsock.go:151
		// _ = "end of CoverTab[16675]"
	} else {
//line /usr/local/go/src/net/udpsock.go:152
		_go_fuzz_dep_.CoverTab[16676]++
//line /usr/local/go/src/net/udpsock.go:152
		// _ = "end of CoverTab[16676]"
//line /usr/local/go/src/net/udpsock.go:152
	}
//line /usr/local/go/src/net/udpsock.go:152
	// _ = "end of CoverTab[16671]"
//line /usr/local/go/src/net/udpsock.go:152
	_go_fuzz_dep_.CoverTab[16672]++
						return n, addr, err
//line /usr/local/go/src/net/udpsock.go:153
	// _ = "end of CoverTab[16672]"
}

// ReadFrom implements the PacketConn ReadFrom method.
func (c *UDPConn) ReadFrom(b []byte) (int, Addr, error) {
//line /usr/local/go/src/net/udpsock.go:157
	_go_fuzz_dep_.CoverTab[16677]++
						n, addr, err := c.readFromUDP(b, &UDPAddr{})
						if addr == nil {
//line /usr/local/go/src/net/udpsock.go:159
		_go_fuzz_dep_.CoverTab[16679]++

							return n, nil, err
//line /usr/local/go/src/net/udpsock.go:161
		// _ = "end of CoverTab[16679]"
	} else {
//line /usr/local/go/src/net/udpsock.go:162
		_go_fuzz_dep_.CoverTab[16680]++
//line /usr/local/go/src/net/udpsock.go:162
		// _ = "end of CoverTab[16680]"
//line /usr/local/go/src/net/udpsock.go:162
	}
//line /usr/local/go/src/net/udpsock.go:162
	// _ = "end of CoverTab[16677]"
//line /usr/local/go/src/net/udpsock.go:162
	_go_fuzz_dep_.CoverTab[16678]++
						return n, addr, err
//line /usr/local/go/src/net/udpsock.go:163
	// _ = "end of CoverTab[16678]"
}

// ReadFromUDPAddrPort acts like ReadFrom but returns a netip.AddrPort.
//line /usr/local/go/src/net/udpsock.go:166
//
//line /usr/local/go/src/net/udpsock.go:166
// If c is bound to an unspecified address, the returned
//line /usr/local/go/src/net/udpsock.go:166
// netip.AddrPort's address might be an IPv4-mapped IPv6 address.
//line /usr/local/go/src/net/udpsock.go:166
// Use netip.Addr.Unmap to get the address without the IPv6 prefix.
//line /usr/local/go/src/net/udpsock.go:171
func (c *UDPConn) ReadFromUDPAddrPort(b []byte) (n int, addr netip.AddrPort, err error) {
//line /usr/local/go/src/net/udpsock.go:171
	_go_fuzz_dep_.CoverTab[16681]++
						if !c.ok() {
//line /usr/local/go/src/net/udpsock.go:172
		_go_fuzz_dep_.CoverTab[16684]++
							return 0, netip.AddrPort{}, syscall.EINVAL
//line /usr/local/go/src/net/udpsock.go:173
		// _ = "end of CoverTab[16684]"
	} else {
//line /usr/local/go/src/net/udpsock.go:174
		_go_fuzz_dep_.CoverTab[16685]++
//line /usr/local/go/src/net/udpsock.go:174
		// _ = "end of CoverTab[16685]"
//line /usr/local/go/src/net/udpsock.go:174
	}
//line /usr/local/go/src/net/udpsock.go:174
	// _ = "end of CoverTab[16681]"
//line /usr/local/go/src/net/udpsock.go:174
	_go_fuzz_dep_.CoverTab[16682]++
						n, addr, err = c.readFromAddrPort(b)
						if err != nil {
//line /usr/local/go/src/net/udpsock.go:176
		_go_fuzz_dep_.CoverTab[16686]++
							err = &OpError{Op: "read", Net: c.fd.net, Source: c.fd.laddr, Addr: c.fd.raddr, Err: err}
//line /usr/local/go/src/net/udpsock.go:177
		// _ = "end of CoverTab[16686]"
	} else {
//line /usr/local/go/src/net/udpsock.go:178
		_go_fuzz_dep_.CoverTab[16687]++
//line /usr/local/go/src/net/udpsock.go:178
		// _ = "end of CoverTab[16687]"
//line /usr/local/go/src/net/udpsock.go:178
	}
//line /usr/local/go/src/net/udpsock.go:178
	// _ = "end of CoverTab[16682]"
//line /usr/local/go/src/net/udpsock.go:178
	_go_fuzz_dep_.CoverTab[16683]++
						return n, addr, err
//line /usr/local/go/src/net/udpsock.go:179
	// _ = "end of CoverTab[16683]"
}

// ReadMsgUDP reads a message from c, copying the payload into b and
//line /usr/local/go/src/net/udpsock.go:182
// the associated out-of-band data into oob. It returns the number of
//line /usr/local/go/src/net/udpsock.go:182
// bytes copied into b, the number of bytes copied into oob, the flags
//line /usr/local/go/src/net/udpsock.go:182
// that were set on the message and the source address of the message.
//line /usr/local/go/src/net/udpsock.go:182
//
//line /usr/local/go/src/net/udpsock.go:182
// The packages golang.org/x/net/ipv4 and golang.org/x/net/ipv6 can be
//line /usr/local/go/src/net/udpsock.go:182
// used to manipulate IP-level socket options in oob.
//line /usr/local/go/src/net/udpsock.go:189
func (c *UDPConn) ReadMsgUDP(b, oob []byte) (n, oobn, flags int, addr *UDPAddr, err error) {
//line /usr/local/go/src/net/udpsock.go:189
	_go_fuzz_dep_.CoverTab[16688]++
						var ap netip.AddrPort
						n, oobn, flags, ap, err = c.ReadMsgUDPAddrPort(b, oob)
						if ap.IsValid() {
//line /usr/local/go/src/net/udpsock.go:192
		_go_fuzz_dep_.CoverTab[16690]++
							addr = UDPAddrFromAddrPort(ap)
//line /usr/local/go/src/net/udpsock.go:193
		// _ = "end of CoverTab[16690]"
	} else {
//line /usr/local/go/src/net/udpsock.go:194
		_go_fuzz_dep_.CoverTab[16691]++
//line /usr/local/go/src/net/udpsock.go:194
		// _ = "end of CoverTab[16691]"
//line /usr/local/go/src/net/udpsock.go:194
	}
//line /usr/local/go/src/net/udpsock.go:194
	// _ = "end of CoverTab[16688]"
//line /usr/local/go/src/net/udpsock.go:194
	_go_fuzz_dep_.CoverTab[16689]++
						return
//line /usr/local/go/src/net/udpsock.go:195
	// _ = "end of CoverTab[16689]"
}

// ReadMsgUDPAddrPort is like ReadMsgUDP but returns an netip.AddrPort instead of a UDPAddr.
func (c *UDPConn) ReadMsgUDPAddrPort(b, oob []byte) (n, oobn, flags int, addr netip.AddrPort, err error) {
//line /usr/local/go/src/net/udpsock.go:199
	_go_fuzz_dep_.CoverTab[16692]++
						if !c.ok() {
//line /usr/local/go/src/net/udpsock.go:200
		_go_fuzz_dep_.CoverTab[16695]++
							return 0, 0, 0, netip.AddrPort{}, syscall.EINVAL
//line /usr/local/go/src/net/udpsock.go:201
		// _ = "end of CoverTab[16695]"
	} else {
//line /usr/local/go/src/net/udpsock.go:202
		_go_fuzz_dep_.CoverTab[16696]++
//line /usr/local/go/src/net/udpsock.go:202
		// _ = "end of CoverTab[16696]"
//line /usr/local/go/src/net/udpsock.go:202
	}
//line /usr/local/go/src/net/udpsock.go:202
	// _ = "end of CoverTab[16692]"
//line /usr/local/go/src/net/udpsock.go:202
	_go_fuzz_dep_.CoverTab[16693]++
						n, oobn, flags, addr, err = c.readMsg(b, oob)
						if err != nil {
//line /usr/local/go/src/net/udpsock.go:204
		_go_fuzz_dep_.CoverTab[16697]++
							err = &OpError{Op: "read", Net: c.fd.net, Source: c.fd.laddr, Addr: c.fd.raddr, Err: err}
//line /usr/local/go/src/net/udpsock.go:205
		// _ = "end of CoverTab[16697]"
	} else {
//line /usr/local/go/src/net/udpsock.go:206
		_go_fuzz_dep_.CoverTab[16698]++
//line /usr/local/go/src/net/udpsock.go:206
		// _ = "end of CoverTab[16698]"
//line /usr/local/go/src/net/udpsock.go:206
	}
//line /usr/local/go/src/net/udpsock.go:206
	// _ = "end of CoverTab[16693]"
//line /usr/local/go/src/net/udpsock.go:206
	_go_fuzz_dep_.CoverTab[16694]++
						return
//line /usr/local/go/src/net/udpsock.go:207
	// _ = "end of CoverTab[16694]"
}

// WriteToUDP acts like WriteTo but takes a UDPAddr.
func (c *UDPConn) WriteToUDP(b []byte, addr *UDPAddr) (int, error) {
//line /usr/local/go/src/net/udpsock.go:211
	_go_fuzz_dep_.CoverTab[16699]++
						if !c.ok() {
//line /usr/local/go/src/net/udpsock.go:212
		_go_fuzz_dep_.CoverTab[16702]++
							return 0, syscall.EINVAL
//line /usr/local/go/src/net/udpsock.go:213
		// _ = "end of CoverTab[16702]"
	} else {
//line /usr/local/go/src/net/udpsock.go:214
		_go_fuzz_dep_.CoverTab[16703]++
//line /usr/local/go/src/net/udpsock.go:214
		// _ = "end of CoverTab[16703]"
//line /usr/local/go/src/net/udpsock.go:214
	}
//line /usr/local/go/src/net/udpsock.go:214
	// _ = "end of CoverTab[16699]"
//line /usr/local/go/src/net/udpsock.go:214
	_go_fuzz_dep_.CoverTab[16700]++
						n, err := c.writeTo(b, addr)
						if err != nil {
//line /usr/local/go/src/net/udpsock.go:216
		_go_fuzz_dep_.CoverTab[16704]++
							err = &OpError{Op: "write", Net: c.fd.net, Source: c.fd.laddr, Addr: addr.opAddr(), Err: err}
//line /usr/local/go/src/net/udpsock.go:217
		// _ = "end of CoverTab[16704]"
	} else {
//line /usr/local/go/src/net/udpsock.go:218
		_go_fuzz_dep_.CoverTab[16705]++
//line /usr/local/go/src/net/udpsock.go:218
		// _ = "end of CoverTab[16705]"
//line /usr/local/go/src/net/udpsock.go:218
	}
//line /usr/local/go/src/net/udpsock.go:218
	// _ = "end of CoverTab[16700]"
//line /usr/local/go/src/net/udpsock.go:218
	_go_fuzz_dep_.CoverTab[16701]++
						return n, err
//line /usr/local/go/src/net/udpsock.go:219
	// _ = "end of CoverTab[16701]"
}

// WriteToUDPAddrPort acts like WriteTo but takes a netip.AddrPort.
func (c *UDPConn) WriteToUDPAddrPort(b []byte, addr netip.AddrPort) (int, error) {
//line /usr/local/go/src/net/udpsock.go:223
	_go_fuzz_dep_.CoverTab[16706]++
						if !c.ok() {
//line /usr/local/go/src/net/udpsock.go:224
		_go_fuzz_dep_.CoverTab[16709]++
							return 0, syscall.EINVAL
//line /usr/local/go/src/net/udpsock.go:225
		// _ = "end of CoverTab[16709]"
	} else {
//line /usr/local/go/src/net/udpsock.go:226
		_go_fuzz_dep_.CoverTab[16710]++
//line /usr/local/go/src/net/udpsock.go:226
		// _ = "end of CoverTab[16710]"
//line /usr/local/go/src/net/udpsock.go:226
	}
//line /usr/local/go/src/net/udpsock.go:226
	// _ = "end of CoverTab[16706]"
//line /usr/local/go/src/net/udpsock.go:226
	_go_fuzz_dep_.CoverTab[16707]++
						n, err := c.writeToAddrPort(b, addr)
						if err != nil {
//line /usr/local/go/src/net/udpsock.go:228
		_go_fuzz_dep_.CoverTab[16711]++
							err = &OpError{Op: "write", Net: c.fd.net, Source: c.fd.laddr, Addr: addrPortUDPAddr{addr}, Err: err}
//line /usr/local/go/src/net/udpsock.go:229
		// _ = "end of CoverTab[16711]"
	} else {
//line /usr/local/go/src/net/udpsock.go:230
		_go_fuzz_dep_.CoverTab[16712]++
//line /usr/local/go/src/net/udpsock.go:230
		// _ = "end of CoverTab[16712]"
//line /usr/local/go/src/net/udpsock.go:230
	}
//line /usr/local/go/src/net/udpsock.go:230
	// _ = "end of CoverTab[16707]"
//line /usr/local/go/src/net/udpsock.go:230
	_go_fuzz_dep_.CoverTab[16708]++
						return n, err
//line /usr/local/go/src/net/udpsock.go:231
	// _ = "end of CoverTab[16708]"
}

// WriteTo implements the PacketConn WriteTo method.
func (c *UDPConn) WriteTo(b []byte, addr Addr) (int, error) {
//line /usr/local/go/src/net/udpsock.go:235
	_go_fuzz_dep_.CoverTab[16713]++
						if !c.ok() {
//line /usr/local/go/src/net/udpsock.go:236
		_go_fuzz_dep_.CoverTab[16717]++
							return 0, syscall.EINVAL
//line /usr/local/go/src/net/udpsock.go:237
		// _ = "end of CoverTab[16717]"
	} else {
//line /usr/local/go/src/net/udpsock.go:238
		_go_fuzz_dep_.CoverTab[16718]++
//line /usr/local/go/src/net/udpsock.go:238
		// _ = "end of CoverTab[16718]"
//line /usr/local/go/src/net/udpsock.go:238
	}
//line /usr/local/go/src/net/udpsock.go:238
	// _ = "end of CoverTab[16713]"
//line /usr/local/go/src/net/udpsock.go:238
	_go_fuzz_dep_.CoverTab[16714]++
						a, ok := addr.(*UDPAddr)
						if !ok {
//line /usr/local/go/src/net/udpsock.go:240
		_go_fuzz_dep_.CoverTab[16719]++
							return 0, &OpError{Op: "write", Net: c.fd.net, Source: c.fd.laddr, Addr: addr, Err: syscall.EINVAL}
//line /usr/local/go/src/net/udpsock.go:241
		// _ = "end of CoverTab[16719]"
	} else {
//line /usr/local/go/src/net/udpsock.go:242
		_go_fuzz_dep_.CoverTab[16720]++
//line /usr/local/go/src/net/udpsock.go:242
		// _ = "end of CoverTab[16720]"
//line /usr/local/go/src/net/udpsock.go:242
	}
//line /usr/local/go/src/net/udpsock.go:242
	// _ = "end of CoverTab[16714]"
//line /usr/local/go/src/net/udpsock.go:242
	_go_fuzz_dep_.CoverTab[16715]++
						n, err := c.writeTo(b, a)
						if err != nil {
//line /usr/local/go/src/net/udpsock.go:244
		_go_fuzz_dep_.CoverTab[16721]++
							err = &OpError{Op: "write", Net: c.fd.net, Source: c.fd.laddr, Addr: a.opAddr(), Err: err}
//line /usr/local/go/src/net/udpsock.go:245
		// _ = "end of CoverTab[16721]"
	} else {
//line /usr/local/go/src/net/udpsock.go:246
		_go_fuzz_dep_.CoverTab[16722]++
//line /usr/local/go/src/net/udpsock.go:246
		// _ = "end of CoverTab[16722]"
//line /usr/local/go/src/net/udpsock.go:246
	}
//line /usr/local/go/src/net/udpsock.go:246
	// _ = "end of CoverTab[16715]"
//line /usr/local/go/src/net/udpsock.go:246
	_go_fuzz_dep_.CoverTab[16716]++
						return n, err
//line /usr/local/go/src/net/udpsock.go:247
	// _ = "end of CoverTab[16716]"
}

// WriteMsgUDP writes a message to addr via c if c isn't connected, or
//line /usr/local/go/src/net/udpsock.go:250
// to c's remote address if c is connected (in which case addr must be
//line /usr/local/go/src/net/udpsock.go:250
// nil). The payload is copied from b and the associated out-of-band
//line /usr/local/go/src/net/udpsock.go:250
// data is copied from oob. It returns the number of payload and
//line /usr/local/go/src/net/udpsock.go:250
// out-of-band bytes written.
//line /usr/local/go/src/net/udpsock.go:250
//
//line /usr/local/go/src/net/udpsock.go:250
// The packages golang.org/x/net/ipv4 and golang.org/x/net/ipv6 can be
//line /usr/local/go/src/net/udpsock.go:250
// used to manipulate IP-level socket options in oob.
//line /usr/local/go/src/net/udpsock.go:258
func (c *UDPConn) WriteMsgUDP(b, oob []byte, addr *UDPAddr) (n, oobn int, err error) {
//line /usr/local/go/src/net/udpsock.go:258
	_go_fuzz_dep_.CoverTab[16723]++
						if !c.ok() {
//line /usr/local/go/src/net/udpsock.go:259
		_go_fuzz_dep_.CoverTab[16726]++
							return 0, 0, syscall.EINVAL
//line /usr/local/go/src/net/udpsock.go:260
		// _ = "end of CoverTab[16726]"
	} else {
//line /usr/local/go/src/net/udpsock.go:261
		_go_fuzz_dep_.CoverTab[16727]++
//line /usr/local/go/src/net/udpsock.go:261
		// _ = "end of CoverTab[16727]"
//line /usr/local/go/src/net/udpsock.go:261
	}
//line /usr/local/go/src/net/udpsock.go:261
	// _ = "end of CoverTab[16723]"
//line /usr/local/go/src/net/udpsock.go:261
	_go_fuzz_dep_.CoverTab[16724]++
						n, oobn, err = c.writeMsg(b, oob, addr)
						if err != nil {
//line /usr/local/go/src/net/udpsock.go:263
		_go_fuzz_dep_.CoverTab[16728]++
							err = &OpError{Op: "write", Net: c.fd.net, Source: c.fd.laddr, Addr: addr.opAddr(), Err: err}
//line /usr/local/go/src/net/udpsock.go:264
		// _ = "end of CoverTab[16728]"
	} else {
//line /usr/local/go/src/net/udpsock.go:265
		_go_fuzz_dep_.CoverTab[16729]++
//line /usr/local/go/src/net/udpsock.go:265
		// _ = "end of CoverTab[16729]"
//line /usr/local/go/src/net/udpsock.go:265
	}
//line /usr/local/go/src/net/udpsock.go:265
	// _ = "end of CoverTab[16724]"
//line /usr/local/go/src/net/udpsock.go:265
	_go_fuzz_dep_.CoverTab[16725]++
						return
//line /usr/local/go/src/net/udpsock.go:266
	// _ = "end of CoverTab[16725]"
}

// WriteMsgUDPAddrPort is like WriteMsgUDP but takes a netip.AddrPort instead of a UDPAddr.
func (c *UDPConn) WriteMsgUDPAddrPort(b, oob []byte, addr netip.AddrPort) (n, oobn int, err error) {
//line /usr/local/go/src/net/udpsock.go:270
	_go_fuzz_dep_.CoverTab[16730]++
						if !c.ok() {
//line /usr/local/go/src/net/udpsock.go:271
		_go_fuzz_dep_.CoverTab[16733]++
							return 0, 0, syscall.EINVAL
//line /usr/local/go/src/net/udpsock.go:272
		// _ = "end of CoverTab[16733]"
	} else {
//line /usr/local/go/src/net/udpsock.go:273
		_go_fuzz_dep_.CoverTab[16734]++
//line /usr/local/go/src/net/udpsock.go:273
		// _ = "end of CoverTab[16734]"
//line /usr/local/go/src/net/udpsock.go:273
	}
//line /usr/local/go/src/net/udpsock.go:273
	// _ = "end of CoverTab[16730]"
//line /usr/local/go/src/net/udpsock.go:273
	_go_fuzz_dep_.CoverTab[16731]++
						n, oobn, err = c.writeMsgAddrPort(b, oob, addr)
						if err != nil {
//line /usr/local/go/src/net/udpsock.go:275
		_go_fuzz_dep_.CoverTab[16735]++
							err = &OpError{Op: "write", Net: c.fd.net, Source: c.fd.laddr, Addr: addrPortUDPAddr{addr}, Err: err}
//line /usr/local/go/src/net/udpsock.go:276
		// _ = "end of CoverTab[16735]"
	} else {
//line /usr/local/go/src/net/udpsock.go:277
		_go_fuzz_dep_.CoverTab[16736]++
//line /usr/local/go/src/net/udpsock.go:277
		// _ = "end of CoverTab[16736]"
//line /usr/local/go/src/net/udpsock.go:277
	}
//line /usr/local/go/src/net/udpsock.go:277
	// _ = "end of CoverTab[16731]"
//line /usr/local/go/src/net/udpsock.go:277
	_go_fuzz_dep_.CoverTab[16732]++
						return
//line /usr/local/go/src/net/udpsock.go:278
	// _ = "end of CoverTab[16732]"
}

func newUDPConn(fd *netFD) *UDPConn {
//line /usr/local/go/src/net/udpsock.go:281
	_go_fuzz_dep_.CoverTab[16737]++
//line /usr/local/go/src/net/udpsock.go:281
	return &UDPConn{conn{fd}}
//line /usr/local/go/src/net/udpsock.go:281
	// _ = "end of CoverTab[16737]"
//line /usr/local/go/src/net/udpsock.go:281
}

// DialUDP acts like Dial for UDP networks.
//line /usr/local/go/src/net/udpsock.go:283
//
//line /usr/local/go/src/net/udpsock.go:283
// The network must be a UDP network name; see func Dial for details.
//line /usr/local/go/src/net/udpsock.go:283
//
//line /usr/local/go/src/net/udpsock.go:283
// If laddr is nil, a local address is automatically chosen.
//line /usr/local/go/src/net/udpsock.go:283
// If the IP field of raddr is nil or an unspecified IP address, the
//line /usr/local/go/src/net/udpsock.go:283
// local system is assumed.
//line /usr/local/go/src/net/udpsock.go:290
func DialUDP(network string, laddr, raddr *UDPAddr) (*UDPConn, error) {
//line /usr/local/go/src/net/udpsock.go:290
	_go_fuzz_dep_.CoverTab[16738]++
						switch network {
	case "udp", "udp4", "udp6":
//line /usr/local/go/src/net/udpsock.go:292
		_go_fuzz_dep_.CoverTab[16742]++
//line /usr/local/go/src/net/udpsock.go:292
		// _ = "end of CoverTab[16742]"
	default:
//line /usr/local/go/src/net/udpsock.go:293
		_go_fuzz_dep_.CoverTab[16743]++
							return nil, &OpError{Op: "dial", Net: network, Source: laddr.opAddr(), Addr: raddr.opAddr(), Err: UnknownNetworkError(network)}
//line /usr/local/go/src/net/udpsock.go:294
		// _ = "end of CoverTab[16743]"
	}
//line /usr/local/go/src/net/udpsock.go:295
	// _ = "end of CoverTab[16738]"
//line /usr/local/go/src/net/udpsock.go:295
	_go_fuzz_dep_.CoverTab[16739]++
						if raddr == nil {
//line /usr/local/go/src/net/udpsock.go:296
		_go_fuzz_dep_.CoverTab[16744]++
							return nil, &OpError{Op: "dial", Net: network, Source: laddr.opAddr(), Addr: nil, Err: errMissingAddress}
//line /usr/local/go/src/net/udpsock.go:297
		// _ = "end of CoverTab[16744]"
	} else {
//line /usr/local/go/src/net/udpsock.go:298
		_go_fuzz_dep_.CoverTab[16745]++
//line /usr/local/go/src/net/udpsock.go:298
		// _ = "end of CoverTab[16745]"
//line /usr/local/go/src/net/udpsock.go:298
	}
//line /usr/local/go/src/net/udpsock.go:298
	// _ = "end of CoverTab[16739]"
//line /usr/local/go/src/net/udpsock.go:298
	_go_fuzz_dep_.CoverTab[16740]++
						sd := &sysDialer{network: network, address: raddr.String()}
						c, err := sd.dialUDP(context.Background(), laddr, raddr)
						if err != nil {
//line /usr/local/go/src/net/udpsock.go:301
		_go_fuzz_dep_.CoverTab[16746]++
							return nil, &OpError{Op: "dial", Net: network, Source: laddr.opAddr(), Addr: raddr.opAddr(), Err: err}
//line /usr/local/go/src/net/udpsock.go:302
		// _ = "end of CoverTab[16746]"
	} else {
//line /usr/local/go/src/net/udpsock.go:303
		_go_fuzz_dep_.CoverTab[16747]++
//line /usr/local/go/src/net/udpsock.go:303
		// _ = "end of CoverTab[16747]"
//line /usr/local/go/src/net/udpsock.go:303
	}
//line /usr/local/go/src/net/udpsock.go:303
	// _ = "end of CoverTab[16740]"
//line /usr/local/go/src/net/udpsock.go:303
	_go_fuzz_dep_.CoverTab[16741]++
						return c, nil
//line /usr/local/go/src/net/udpsock.go:304
	// _ = "end of CoverTab[16741]"
}

// ListenUDP acts like ListenPacket for UDP networks.
//line /usr/local/go/src/net/udpsock.go:307
//
//line /usr/local/go/src/net/udpsock.go:307
// The network must be a UDP network name; see func Dial for details.
//line /usr/local/go/src/net/udpsock.go:307
//
//line /usr/local/go/src/net/udpsock.go:307
// If the IP field of laddr is nil or an unspecified IP address,
//line /usr/local/go/src/net/udpsock.go:307
// ListenUDP listens on all available IP addresses of the local system
//line /usr/local/go/src/net/udpsock.go:307
// except multicast IP addresses.
//line /usr/local/go/src/net/udpsock.go:307
// If the Port field of laddr is 0, a port number is automatically
//line /usr/local/go/src/net/udpsock.go:307
// chosen.
//line /usr/local/go/src/net/udpsock.go:316
func ListenUDP(network string, laddr *UDPAddr) (*UDPConn, error) {
//line /usr/local/go/src/net/udpsock.go:316
	_go_fuzz_dep_.CoverTab[16748]++
						switch network {
	case "udp", "udp4", "udp6":
//line /usr/local/go/src/net/udpsock.go:318
		_go_fuzz_dep_.CoverTab[16752]++
//line /usr/local/go/src/net/udpsock.go:318
		// _ = "end of CoverTab[16752]"
	default:
//line /usr/local/go/src/net/udpsock.go:319
		_go_fuzz_dep_.CoverTab[16753]++
							return nil, &OpError{Op: "listen", Net: network, Source: nil, Addr: laddr.opAddr(), Err: UnknownNetworkError(network)}
//line /usr/local/go/src/net/udpsock.go:320
		// _ = "end of CoverTab[16753]"
	}
//line /usr/local/go/src/net/udpsock.go:321
	// _ = "end of CoverTab[16748]"
//line /usr/local/go/src/net/udpsock.go:321
	_go_fuzz_dep_.CoverTab[16749]++
						if laddr == nil {
//line /usr/local/go/src/net/udpsock.go:322
		_go_fuzz_dep_.CoverTab[16754]++
							laddr = &UDPAddr{}
//line /usr/local/go/src/net/udpsock.go:323
		// _ = "end of CoverTab[16754]"
	} else {
//line /usr/local/go/src/net/udpsock.go:324
		_go_fuzz_dep_.CoverTab[16755]++
//line /usr/local/go/src/net/udpsock.go:324
		// _ = "end of CoverTab[16755]"
//line /usr/local/go/src/net/udpsock.go:324
	}
//line /usr/local/go/src/net/udpsock.go:324
	// _ = "end of CoverTab[16749]"
//line /usr/local/go/src/net/udpsock.go:324
	_go_fuzz_dep_.CoverTab[16750]++
						sl := &sysListener{network: network, address: laddr.String()}
						c, err := sl.listenUDP(context.Background(), laddr)
						if err != nil {
//line /usr/local/go/src/net/udpsock.go:327
		_go_fuzz_dep_.CoverTab[16756]++
							return nil, &OpError{Op: "listen", Net: network, Source: nil, Addr: laddr.opAddr(), Err: err}
//line /usr/local/go/src/net/udpsock.go:328
		// _ = "end of CoverTab[16756]"
	} else {
//line /usr/local/go/src/net/udpsock.go:329
		_go_fuzz_dep_.CoverTab[16757]++
//line /usr/local/go/src/net/udpsock.go:329
		// _ = "end of CoverTab[16757]"
//line /usr/local/go/src/net/udpsock.go:329
	}
//line /usr/local/go/src/net/udpsock.go:329
	// _ = "end of CoverTab[16750]"
//line /usr/local/go/src/net/udpsock.go:329
	_go_fuzz_dep_.CoverTab[16751]++
						return c, nil
//line /usr/local/go/src/net/udpsock.go:330
	// _ = "end of CoverTab[16751]"
}

// ListenMulticastUDP acts like ListenPacket for UDP networks but
//line /usr/local/go/src/net/udpsock.go:333
// takes a group address on a specific network interface.
//line /usr/local/go/src/net/udpsock.go:333
//
//line /usr/local/go/src/net/udpsock.go:333
// The network must be a UDP network name; see func Dial for details.
//line /usr/local/go/src/net/udpsock.go:333
//
//line /usr/local/go/src/net/udpsock.go:333
// ListenMulticastUDP listens on all available IP addresses of the
//line /usr/local/go/src/net/udpsock.go:333
// local system including the group, multicast IP address.
//line /usr/local/go/src/net/udpsock.go:333
// If ifi is nil, ListenMulticastUDP uses the system-assigned
//line /usr/local/go/src/net/udpsock.go:333
// multicast interface, although this is not recommended because the
//line /usr/local/go/src/net/udpsock.go:333
// assignment depends on platforms and sometimes it might require
//line /usr/local/go/src/net/udpsock.go:333
// routing configuration.
//line /usr/local/go/src/net/udpsock.go:333
// If the Port field of gaddr is 0, a port number is automatically
//line /usr/local/go/src/net/udpsock.go:333
// chosen.
//line /usr/local/go/src/net/udpsock.go:333
//
//line /usr/local/go/src/net/udpsock.go:333
// ListenMulticastUDP is just for convenience of simple, small
//line /usr/local/go/src/net/udpsock.go:333
// applications. There are golang.org/x/net/ipv4 and
//line /usr/local/go/src/net/udpsock.go:333
// golang.org/x/net/ipv6 packages for general purpose uses.
//line /usr/local/go/src/net/udpsock.go:333
//
//line /usr/local/go/src/net/udpsock.go:333
// Note that ListenMulticastUDP will set the IP_MULTICAST_LOOP socket option
//line /usr/local/go/src/net/udpsock.go:333
// to 0 under IPPROTO_IP, to disable loopback of multicast packets.
//line /usr/local/go/src/net/udpsock.go:353
func ListenMulticastUDP(network string, ifi *Interface, gaddr *UDPAddr) (*UDPConn, error) {
//line /usr/local/go/src/net/udpsock.go:353
	_go_fuzz_dep_.CoverTab[16758]++
						switch network {
	case "udp", "udp4", "udp6":
//line /usr/local/go/src/net/udpsock.go:355
		_go_fuzz_dep_.CoverTab[16762]++
//line /usr/local/go/src/net/udpsock.go:355
		// _ = "end of CoverTab[16762]"
	default:
//line /usr/local/go/src/net/udpsock.go:356
		_go_fuzz_dep_.CoverTab[16763]++
							return nil, &OpError{Op: "listen", Net: network, Source: nil, Addr: gaddr.opAddr(), Err: UnknownNetworkError(network)}
//line /usr/local/go/src/net/udpsock.go:357
		// _ = "end of CoverTab[16763]"
	}
//line /usr/local/go/src/net/udpsock.go:358
	// _ = "end of CoverTab[16758]"
//line /usr/local/go/src/net/udpsock.go:358
	_go_fuzz_dep_.CoverTab[16759]++
						if gaddr == nil || func() bool {
//line /usr/local/go/src/net/udpsock.go:359
		_go_fuzz_dep_.CoverTab[16764]++
//line /usr/local/go/src/net/udpsock.go:359
		return gaddr.IP == nil
//line /usr/local/go/src/net/udpsock.go:359
		// _ = "end of CoverTab[16764]"
//line /usr/local/go/src/net/udpsock.go:359
	}() {
//line /usr/local/go/src/net/udpsock.go:359
		_go_fuzz_dep_.CoverTab[16765]++
							return nil, &OpError{Op: "listen", Net: network, Source: nil, Addr: gaddr.opAddr(), Err: errMissingAddress}
//line /usr/local/go/src/net/udpsock.go:360
		// _ = "end of CoverTab[16765]"
	} else {
//line /usr/local/go/src/net/udpsock.go:361
		_go_fuzz_dep_.CoverTab[16766]++
//line /usr/local/go/src/net/udpsock.go:361
		// _ = "end of CoverTab[16766]"
//line /usr/local/go/src/net/udpsock.go:361
	}
//line /usr/local/go/src/net/udpsock.go:361
	// _ = "end of CoverTab[16759]"
//line /usr/local/go/src/net/udpsock.go:361
	_go_fuzz_dep_.CoverTab[16760]++
						sl := &sysListener{network: network, address: gaddr.String()}
						c, err := sl.listenMulticastUDP(context.Background(), ifi, gaddr)
						if err != nil {
//line /usr/local/go/src/net/udpsock.go:364
		_go_fuzz_dep_.CoverTab[16767]++
							return nil, &OpError{Op: "listen", Net: network, Source: nil, Addr: gaddr.opAddr(), Err: err}
//line /usr/local/go/src/net/udpsock.go:365
		// _ = "end of CoverTab[16767]"
	} else {
//line /usr/local/go/src/net/udpsock.go:366
		_go_fuzz_dep_.CoverTab[16768]++
//line /usr/local/go/src/net/udpsock.go:366
		// _ = "end of CoverTab[16768]"
//line /usr/local/go/src/net/udpsock.go:366
	}
//line /usr/local/go/src/net/udpsock.go:366
	// _ = "end of CoverTab[16760]"
//line /usr/local/go/src/net/udpsock.go:366
	_go_fuzz_dep_.CoverTab[16761]++
						return c, nil
//line /usr/local/go/src/net/udpsock.go:367
	// _ = "end of CoverTab[16761]"
}

//line /usr/local/go/src/net/udpsock.go:368
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/udpsock.go:368
var _ = _go_fuzz_dep_.CoverTab
