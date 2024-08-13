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
	_go_fuzz_dep_.CoverTab[8244]++
						if a == nil {
//line /usr/local/go/src/net/udpsock.go:36
		_go_fuzz_dep_.CoverTab[8246]++
							return netip.AddrPort{}
//line /usr/local/go/src/net/udpsock.go:37
		// _ = "end of CoverTab[8246]"
	} else {
//line /usr/local/go/src/net/udpsock.go:38
		_go_fuzz_dep_.CoverTab[8247]++
//line /usr/local/go/src/net/udpsock.go:38
		// _ = "end of CoverTab[8247]"
//line /usr/local/go/src/net/udpsock.go:38
	}
//line /usr/local/go/src/net/udpsock.go:38
	// _ = "end of CoverTab[8244]"
//line /usr/local/go/src/net/udpsock.go:38
	_go_fuzz_dep_.CoverTab[8245]++
						na, _ := netip.AddrFromSlice(a.IP)
						na = na.WithZone(a.Zone)
						return netip.AddrPortFrom(na, uint16(a.Port))
//line /usr/local/go/src/net/udpsock.go:41
	// _ = "end of CoverTab[8245]"
}

// Network returns the address's network name, "udp".
func (a *UDPAddr) Network() string {
//line /usr/local/go/src/net/udpsock.go:45
	_go_fuzz_dep_.CoverTab[8248]++
//line /usr/local/go/src/net/udpsock.go:45
	return "udp"
//line /usr/local/go/src/net/udpsock.go:45
	// _ = "end of CoverTab[8248]"
//line /usr/local/go/src/net/udpsock.go:45
}

func (a *UDPAddr) String() string {
//line /usr/local/go/src/net/udpsock.go:47
	_go_fuzz_dep_.CoverTab[8249]++
						if a == nil {
//line /usr/local/go/src/net/udpsock.go:48
		_go_fuzz_dep_.CoverTab[8252]++
							return "<nil>"
//line /usr/local/go/src/net/udpsock.go:49
		// _ = "end of CoverTab[8252]"
	} else {
//line /usr/local/go/src/net/udpsock.go:50
		_go_fuzz_dep_.CoverTab[8253]++
//line /usr/local/go/src/net/udpsock.go:50
		// _ = "end of CoverTab[8253]"
//line /usr/local/go/src/net/udpsock.go:50
	}
//line /usr/local/go/src/net/udpsock.go:50
	// _ = "end of CoverTab[8249]"
//line /usr/local/go/src/net/udpsock.go:50
	_go_fuzz_dep_.CoverTab[8250]++
						ip := ipEmptyString(a.IP)
						if a.Zone != "" {
//line /usr/local/go/src/net/udpsock.go:52
		_go_fuzz_dep_.CoverTab[8254]++
							return JoinHostPort(ip+"%"+a.Zone, itoa.Itoa(a.Port))
//line /usr/local/go/src/net/udpsock.go:53
		// _ = "end of CoverTab[8254]"
	} else {
//line /usr/local/go/src/net/udpsock.go:54
		_go_fuzz_dep_.CoverTab[8255]++
//line /usr/local/go/src/net/udpsock.go:54
		// _ = "end of CoverTab[8255]"
//line /usr/local/go/src/net/udpsock.go:54
	}
//line /usr/local/go/src/net/udpsock.go:54
	// _ = "end of CoverTab[8250]"
//line /usr/local/go/src/net/udpsock.go:54
	_go_fuzz_dep_.CoverTab[8251]++
						return JoinHostPort(ip, itoa.Itoa(a.Port))
//line /usr/local/go/src/net/udpsock.go:55
	// _ = "end of CoverTab[8251]"
}

func (a *UDPAddr) isWildcard() bool {
//line /usr/local/go/src/net/udpsock.go:58
	_go_fuzz_dep_.CoverTab[8256]++
						if a == nil || func() bool {
//line /usr/local/go/src/net/udpsock.go:59
		_go_fuzz_dep_.CoverTab[8258]++
//line /usr/local/go/src/net/udpsock.go:59
		return a.IP == nil
//line /usr/local/go/src/net/udpsock.go:59
		// _ = "end of CoverTab[8258]"
//line /usr/local/go/src/net/udpsock.go:59
	}() {
//line /usr/local/go/src/net/udpsock.go:59
		_go_fuzz_dep_.CoverTab[8259]++
							return true
//line /usr/local/go/src/net/udpsock.go:60
		// _ = "end of CoverTab[8259]"
	} else {
//line /usr/local/go/src/net/udpsock.go:61
		_go_fuzz_dep_.CoverTab[8260]++
//line /usr/local/go/src/net/udpsock.go:61
		// _ = "end of CoverTab[8260]"
//line /usr/local/go/src/net/udpsock.go:61
	}
//line /usr/local/go/src/net/udpsock.go:61
	// _ = "end of CoverTab[8256]"
//line /usr/local/go/src/net/udpsock.go:61
	_go_fuzz_dep_.CoverTab[8257]++
						return a.IP.IsUnspecified()
//line /usr/local/go/src/net/udpsock.go:62
	// _ = "end of CoverTab[8257]"
}

func (a *UDPAddr) opAddr() Addr {
//line /usr/local/go/src/net/udpsock.go:65
	_go_fuzz_dep_.CoverTab[8261]++
						if a == nil {
//line /usr/local/go/src/net/udpsock.go:66
		_go_fuzz_dep_.CoverTab[8263]++
							return nil
//line /usr/local/go/src/net/udpsock.go:67
		// _ = "end of CoverTab[8263]"
	} else {
//line /usr/local/go/src/net/udpsock.go:68
		_go_fuzz_dep_.CoverTab[8264]++
//line /usr/local/go/src/net/udpsock.go:68
		// _ = "end of CoverTab[8264]"
//line /usr/local/go/src/net/udpsock.go:68
	}
//line /usr/local/go/src/net/udpsock.go:68
	// _ = "end of CoverTab[8261]"
//line /usr/local/go/src/net/udpsock.go:68
	_go_fuzz_dep_.CoverTab[8262]++
						return a
//line /usr/local/go/src/net/udpsock.go:69
	// _ = "end of CoverTab[8262]"
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
	_go_fuzz_dep_.CoverTab[8265]++
						switch network {
	case "udp", "udp4", "udp6":
//line /usr/local/go/src/net/udpsock.go:89
		_go_fuzz_dep_.CoverTab[8268]++
//line /usr/local/go/src/net/udpsock.go:89
		// _ = "end of CoverTab[8268]"
	case "":
//line /usr/local/go/src/net/udpsock.go:90
		_go_fuzz_dep_.CoverTab[8269]++
							network = "udp"
//line /usr/local/go/src/net/udpsock.go:91
		// _ = "end of CoverTab[8269]"
	default:
//line /usr/local/go/src/net/udpsock.go:92
		_go_fuzz_dep_.CoverTab[8270]++
							return nil, UnknownNetworkError(network)
//line /usr/local/go/src/net/udpsock.go:93
		// _ = "end of CoverTab[8270]"
	}
//line /usr/local/go/src/net/udpsock.go:94
	// _ = "end of CoverTab[8265]"
//line /usr/local/go/src/net/udpsock.go:94
	_go_fuzz_dep_.CoverTab[8266]++
						addrs, err := DefaultResolver.internetAddrList(context.Background(), network, address)
						if err != nil {
//line /usr/local/go/src/net/udpsock.go:96
		_go_fuzz_dep_.CoverTab[8271]++
							return nil, err
//line /usr/local/go/src/net/udpsock.go:97
		// _ = "end of CoverTab[8271]"
	} else {
//line /usr/local/go/src/net/udpsock.go:98
		_go_fuzz_dep_.CoverTab[8272]++
//line /usr/local/go/src/net/udpsock.go:98
		// _ = "end of CoverTab[8272]"
//line /usr/local/go/src/net/udpsock.go:98
	}
//line /usr/local/go/src/net/udpsock.go:98
	// _ = "end of CoverTab[8266]"
//line /usr/local/go/src/net/udpsock.go:98
	_go_fuzz_dep_.CoverTab[8267]++
						return addrs.forResolve(network, address).(*UDPAddr), nil
//line /usr/local/go/src/net/udpsock.go:99
	// _ = "end of CoverTab[8267]"
}

// UDPAddrFromAddrPort returns addr as a UDPAddr. If addr.IsValid() is false,
//line /usr/local/go/src/net/udpsock.go:102
// then the returned UDPAddr will contain a nil IP field, indicating an
//line /usr/local/go/src/net/udpsock.go:102
// address family-agnostic unspecified address.
//line /usr/local/go/src/net/udpsock.go:105
func UDPAddrFromAddrPort(addr netip.AddrPort) *UDPAddr {
//line /usr/local/go/src/net/udpsock.go:105
	_go_fuzz_dep_.CoverTab[8273]++
						return &UDPAddr{
		IP:	addr.Addr().AsSlice(),
		Zone:	addr.Addr().Zone(),
		Port:	int(addr.Port()),
	}
//line /usr/local/go/src/net/udpsock.go:110
	// _ = "end of CoverTab[8273]"
}

// An addrPortUDPAddr is a netip.AddrPort-based UDP address that satisfies the Addr interface.
type addrPortUDPAddr struct {
	netip.AddrPort
}

func (addrPortUDPAddr) Network() string {
//line /usr/local/go/src/net/udpsock.go:118
	_go_fuzz_dep_.CoverTab[8274]++
//line /usr/local/go/src/net/udpsock.go:118
	return "udp"
//line /usr/local/go/src/net/udpsock.go:118
	// _ = "end of CoverTab[8274]"
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
	_go_fuzz_dep_.CoverTab[8275]++
						if !c.ok() {
//line /usr/local/go/src/net/udpsock.go:129
		_go_fuzz_dep_.CoverTab[8277]++
							return nil, syscall.EINVAL
//line /usr/local/go/src/net/udpsock.go:130
		// _ = "end of CoverTab[8277]"
	} else {
//line /usr/local/go/src/net/udpsock.go:131
		_go_fuzz_dep_.CoverTab[8278]++
//line /usr/local/go/src/net/udpsock.go:131
		// _ = "end of CoverTab[8278]"
//line /usr/local/go/src/net/udpsock.go:131
	}
//line /usr/local/go/src/net/udpsock.go:131
	// _ = "end of CoverTab[8275]"
//line /usr/local/go/src/net/udpsock.go:131
	_go_fuzz_dep_.CoverTab[8276]++
						return newRawConn(c.fd)
//line /usr/local/go/src/net/udpsock.go:132
	// _ = "end of CoverTab[8276]"
}

// ReadFromUDP acts like ReadFrom but returns a UDPAddr.
func (c *UDPConn) ReadFromUDP(b []byte) (n int, addr *UDPAddr, err error) {
//line /usr/local/go/src/net/udpsock.go:136
	_go_fuzz_dep_.CoverTab[8279]++

//line /usr/local/go/src/net/udpsock.go:141
	return c.readFromUDP(b, &UDPAddr{})
//line /usr/local/go/src/net/udpsock.go:141
	// _ = "end of CoverTab[8279]"
}

// readFromUDP implements ReadFromUDP.
func (c *UDPConn) readFromUDP(b []byte, addr *UDPAddr) (int, *UDPAddr, error) {
//line /usr/local/go/src/net/udpsock.go:145
	_go_fuzz_dep_.CoverTab[8280]++
						if !c.ok() {
//line /usr/local/go/src/net/udpsock.go:146
		_go_fuzz_dep_.CoverTab[8283]++
							return 0, nil, syscall.EINVAL
//line /usr/local/go/src/net/udpsock.go:147
		// _ = "end of CoverTab[8283]"
	} else {
//line /usr/local/go/src/net/udpsock.go:148
		_go_fuzz_dep_.CoverTab[8284]++
//line /usr/local/go/src/net/udpsock.go:148
		// _ = "end of CoverTab[8284]"
//line /usr/local/go/src/net/udpsock.go:148
	}
//line /usr/local/go/src/net/udpsock.go:148
	// _ = "end of CoverTab[8280]"
//line /usr/local/go/src/net/udpsock.go:148
	_go_fuzz_dep_.CoverTab[8281]++
						n, addr, err := c.readFrom(b, addr)
						if err != nil {
//line /usr/local/go/src/net/udpsock.go:150
		_go_fuzz_dep_.CoverTab[8285]++
							err = &OpError{Op: "read", Net: c.fd.net, Source: c.fd.laddr, Addr: c.fd.raddr, Err: err}
//line /usr/local/go/src/net/udpsock.go:151
		// _ = "end of CoverTab[8285]"
	} else {
//line /usr/local/go/src/net/udpsock.go:152
		_go_fuzz_dep_.CoverTab[8286]++
//line /usr/local/go/src/net/udpsock.go:152
		// _ = "end of CoverTab[8286]"
//line /usr/local/go/src/net/udpsock.go:152
	}
//line /usr/local/go/src/net/udpsock.go:152
	// _ = "end of CoverTab[8281]"
//line /usr/local/go/src/net/udpsock.go:152
	_go_fuzz_dep_.CoverTab[8282]++
						return n, addr, err
//line /usr/local/go/src/net/udpsock.go:153
	// _ = "end of CoverTab[8282]"
}

// ReadFrom implements the PacketConn ReadFrom method.
func (c *UDPConn) ReadFrom(b []byte) (int, Addr, error) {
//line /usr/local/go/src/net/udpsock.go:157
	_go_fuzz_dep_.CoverTab[8287]++
						n, addr, err := c.readFromUDP(b, &UDPAddr{})
						if addr == nil {
//line /usr/local/go/src/net/udpsock.go:159
		_go_fuzz_dep_.CoverTab[8289]++

							return n, nil, err
//line /usr/local/go/src/net/udpsock.go:161
		// _ = "end of CoverTab[8289]"
	} else {
//line /usr/local/go/src/net/udpsock.go:162
		_go_fuzz_dep_.CoverTab[8290]++
//line /usr/local/go/src/net/udpsock.go:162
		// _ = "end of CoverTab[8290]"
//line /usr/local/go/src/net/udpsock.go:162
	}
//line /usr/local/go/src/net/udpsock.go:162
	// _ = "end of CoverTab[8287]"
//line /usr/local/go/src/net/udpsock.go:162
	_go_fuzz_dep_.CoverTab[8288]++
						return n, addr, err
//line /usr/local/go/src/net/udpsock.go:163
	// _ = "end of CoverTab[8288]"
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
	_go_fuzz_dep_.CoverTab[8291]++
						if !c.ok() {
//line /usr/local/go/src/net/udpsock.go:172
		_go_fuzz_dep_.CoverTab[8294]++
							return 0, netip.AddrPort{}, syscall.EINVAL
//line /usr/local/go/src/net/udpsock.go:173
		// _ = "end of CoverTab[8294]"
	} else {
//line /usr/local/go/src/net/udpsock.go:174
		_go_fuzz_dep_.CoverTab[8295]++
//line /usr/local/go/src/net/udpsock.go:174
		// _ = "end of CoverTab[8295]"
//line /usr/local/go/src/net/udpsock.go:174
	}
//line /usr/local/go/src/net/udpsock.go:174
	// _ = "end of CoverTab[8291]"
//line /usr/local/go/src/net/udpsock.go:174
	_go_fuzz_dep_.CoverTab[8292]++
						n, addr, err = c.readFromAddrPort(b)
						if err != nil {
//line /usr/local/go/src/net/udpsock.go:176
		_go_fuzz_dep_.CoverTab[8296]++
							err = &OpError{Op: "read", Net: c.fd.net, Source: c.fd.laddr, Addr: c.fd.raddr, Err: err}
//line /usr/local/go/src/net/udpsock.go:177
		// _ = "end of CoverTab[8296]"
	} else {
//line /usr/local/go/src/net/udpsock.go:178
		_go_fuzz_dep_.CoverTab[8297]++
//line /usr/local/go/src/net/udpsock.go:178
		// _ = "end of CoverTab[8297]"
//line /usr/local/go/src/net/udpsock.go:178
	}
//line /usr/local/go/src/net/udpsock.go:178
	// _ = "end of CoverTab[8292]"
//line /usr/local/go/src/net/udpsock.go:178
	_go_fuzz_dep_.CoverTab[8293]++
						return n, addr, err
//line /usr/local/go/src/net/udpsock.go:179
	// _ = "end of CoverTab[8293]"
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
	_go_fuzz_dep_.CoverTab[8298]++
						var ap netip.AddrPort
						n, oobn, flags, ap, err = c.ReadMsgUDPAddrPort(b, oob)
						if ap.IsValid() {
//line /usr/local/go/src/net/udpsock.go:192
		_go_fuzz_dep_.CoverTab[8300]++
							addr = UDPAddrFromAddrPort(ap)
//line /usr/local/go/src/net/udpsock.go:193
		// _ = "end of CoverTab[8300]"
	} else {
//line /usr/local/go/src/net/udpsock.go:194
		_go_fuzz_dep_.CoverTab[8301]++
//line /usr/local/go/src/net/udpsock.go:194
		// _ = "end of CoverTab[8301]"
//line /usr/local/go/src/net/udpsock.go:194
	}
//line /usr/local/go/src/net/udpsock.go:194
	// _ = "end of CoverTab[8298]"
//line /usr/local/go/src/net/udpsock.go:194
	_go_fuzz_dep_.CoverTab[8299]++
						return
//line /usr/local/go/src/net/udpsock.go:195
	// _ = "end of CoverTab[8299]"
}

// ReadMsgUDPAddrPort is like ReadMsgUDP but returns an netip.AddrPort instead of a UDPAddr.
func (c *UDPConn) ReadMsgUDPAddrPort(b, oob []byte) (n, oobn, flags int, addr netip.AddrPort, err error) {
//line /usr/local/go/src/net/udpsock.go:199
	_go_fuzz_dep_.CoverTab[8302]++
						if !c.ok() {
//line /usr/local/go/src/net/udpsock.go:200
		_go_fuzz_dep_.CoverTab[8305]++
							return 0, 0, 0, netip.AddrPort{}, syscall.EINVAL
//line /usr/local/go/src/net/udpsock.go:201
		// _ = "end of CoverTab[8305]"
	} else {
//line /usr/local/go/src/net/udpsock.go:202
		_go_fuzz_dep_.CoverTab[8306]++
//line /usr/local/go/src/net/udpsock.go:202
		// _ = "end of CoverTab[8306]"
//line /usr/local/go/src/net/udpsock.go:202
	}
//line /usr/local/go/src/net/udpsock.go:202
	// _ = "end of CoverTab[8302]"
//line /usr/local/go/src/net/udpsock.go:202
	_go_fuzz_dep_.CoverTab[8303]++
						n, oobn, flags, addr, err = c.readMsg(b, oob)
						if err != nil {
//line /usr/local/go/src/net/udpsock.go:204
		_go_fuzz_dep_.CoverTab[8307]++
							err = &OpError{Op: "read", Net: c.fd.net, Source: c.fd.laddr, Addr: c.fd.raddr, Err: err}
//line /usr/local/go/src/net/udpsock.go:205
		// _ = "end of CoverTab[8307]"
	} else {
//line /usr/local/go/src/net/udpsock.go:206
		_go_fuzz_dep_.CoverTab[8308]++
//line /usr/local/go/src/net/udpsock.go:206
		// _ = "end of CoverTab[8308]"
//line /usr/local/go/src/net/udpsock.go:206
	}
//line /usr/local/go/src/net/udpsock.go:206
	// _ = "end of CoverTab[8303]"
//line /usr/local/go/src/net/udpsock.go:206
	_go_fuzz_dep_.CoverTab[8304]++
						return
//line /usr/local/go/src/net/udpsock.go:207
	// _ = "end of CoverTab[8304]"
}

// WriteToUDP acts like WriteTo but takes a UDPAddr.
func (c *UDPConn) WriteToUDP(b []byte, addr *UDPAddr) (int, error) {
//line /usr/local/go/src/net/udpsock.go:211
	_go_fuzz_dep_.CoverTab[8309]++
						if !c.ok() {
//line /usr/local/go/src/net/udpsock.go:212
		_go_fuzz_dep_.CoverTab[8312]++
							return 0, syscall.EINVAL
//line /usr/local/go/src/net/udpsock.go:213
		// _ = "end of CoverTab[8312]"
	} else {
//line /usr/local/go/src/net/udpsock.go:214
		_go_fuzz_dep_.CoverTab[8313]++
//line /usr/local/go/src/net/udpsock.go:214
		// _ = "end of CoverTab[8313]"
//line /usr/local/go/src/net/udpsock.go:214
	}
//line /usr/local/go/src/net/udpsock.go:214
	// _ = "end of CoverTab[8309]"
//line /usr/local/go/src/net/udpsock.go:214
	_go_fuzz_dep_.CoverTab[8310]++
						n, err := c.writeTo(b, addr)
						if err != nil {
//line /usr/local/go/src/net/udpsock.go:216
		_go_fuzz_dep_.CoverTab[8314]++
							err = &OpError{Op: "write", Net: c.fd.net, Source: c.fd.laddr, Addr: addr.opAddr(), Err: err}
//line /usr/local/go/src/net/udpsock.go:217
		// _ = "end of CoverTab[8314]"
	} else {
//line /usr/local/go/src/net/udpsock.go:218
		_go_fuzz_dep_.CoverTab[8315]++
//line /usr/local/go/src/net/udpsock.go:218
		// _ = "end of CoverTab[8315]"
//line /usr/local/go/src/net/udpsock.go:218
	}
//line /usr/local/go/src/net/udpsock.go:218
	// _ = "end of CoverTab[8310]"
//line /usr/local/go/src/net/udpsock.go:218
	_go_fuzz_dep_.CoverTab[8311]++
						return n, err
//line /usr/local/go/src/net/udpsock.go:219
	// _ = "end of CoverTab[8311]"
}

// WriteToUDPAddrPort acts like WriteTo but takes a netip.AddrPort.
func (c *UDPConn) WriteToUDPAddrPort(b []byte, addr netip.AddrPort) (int, error) {
//line /usr/local/go/src/net/udpsock.go:223
	_go_fuzz_dep_.CoverTab[8316]++
						if !c.ok() {
//line /usr/local/go/src/net/udpsock.go:224
		_go_fuzz_dep_.CoverTab[8319]++
							return 0, syscall.EINVAL
//line /usr/local/go/src/net/udpsock.go:225
		// _ = "end of CoverTab[8319]"
	} else {
//line /usr/local/go/src/net/udpsock.go:226
		_go_fuzz_dep_.CoverTab[8320]++
//line /usr/local/go/src/net/udpsock.go:226
		// _ = "end of CoverTab[8320]"
//line /usr/local/go/src/net/udpsock.go:226
	}
//line /usr/local/go/src/net/udpsock.go:226
	// _ = "end of CoverTab[8316]"
//line /usr/local/go/src/net/udpsock.go:226
	_go_fuzz_dep_.CoverTab[8317]++
						n, err := c.writeToAddrPort(b, addr)
						if err != nil {
//line /usr/local/go/src/net/udpsock.go:228
		_go_fuzz_dep_.CoverTab[8321]++
							err = &OpError{Op: "write", Net: c.fd.net, Source: c.fd.laddr, Addr: addrPortUDPAddr{addr}, Err: err}
//line /usr/local/go/src/net/udpsock.go:229
		// _ = "end of CoverTab[8321]"
	} else {
//line /usr/local/go/src/net/udpsock.go:230
		_go_fuzz_dep_.CoverTab[8322]++
//line /usr/local/go/src/net/udpsock.go:230
		// _ = "end of CoverTab[8322]"
//line /usr/local/go/src/net/udpsock.go:230
	}
//line /usr/local/go/src/net/udpsock.go:230
	// _ = "end of CoverTab[8317]"
//line /usr/local/go/src/net/udpsock.go:230
	_go_fuzz_dep_.CoverTab[8318]++
						return n, err
//line /usr/local/go/src/net/udpsock.go:231
	// _ = "end of CoverTab[8318]"
}

// WriteTo implements the PacketConn WriteTo method.
func (c *UDPConn) WriteTo(b []byte, addr Addr) (int, error) {
//line /usr/local/go/src/net/udpsock.go:235
	_go_fuzz_dep_.CoverTab[8323]++
						if !c.ok() {
//line /usr/local/go/src/net/udpsock.go:236
		_go_fuzz_dep_.CoverTab[8327]++
							return 0, syscall.EINVAL
//line /usr/local/go/src/net/udpsock.go:237
		// _ = "end of CoverTab[8327]"
	} else {
//line /usr/local/go/src/net/udpsock.go:238
		_go_fuzz_dep_.CoverTab[8328]++
//line /usr/local/go/src/net/udpsock.go:238
		// _ = "end of CoverTab[8328]"
//line /usr/local/go/src/net/udpsock.go:238
	}
//line /usr/local/go/src/net/udpsock.go:238
	// _ = "end of CoverTab[8323]"
//line /usr/local/go/src/net/udpsock.go:238
	_go_fuzz_dep_.CoverTab[8324]++
						a, ok := addr.(*UDPAddr)
						if !ok {
//line /usr/local/go/src/net/udpsock.go:240
		_go_fuzz_dep_.CoverTab[8329]++
							return 0, &OpError{Op: "write", Net: c.fd.net, Source: c.fd.laddr, Addr: addr, Err: syscall.EINVAL}
//line /usr/local/go/src/net/udpsock.go:241
		// _ = "end of CoverTab[8329]"
	} else {
//line /usr/local/go/src/net/udpsock.go:242
		_go_fuzz_dep_.CoverTab[8330]++
//line /usr/local/go/src/net/udpsock.go:242
		// _ = "end of CoverTab[8330]"
//line /usr/local/go/src/net/udpsock.go:242
	}
//line /usr/local/go/src/net/udpsock.go:242
	// _ = "end of CoverTab[8324]"
//line /usr/local/go/src/net/udpsock.go:242
	_go_fuzz_dep_.CoverTab[8325]++
						n, err := c.writeTo(b, a)
						if err != nil {
//line /usr/local/go/src/net/udpsock.go:244
		_go_fuzz_dep_.CoverTab[8331]++
							err = &OpError{Op: "write", Net: c.fd.net, Source: c.fd.laddr, Addr: a.opAddr(), Err: err}
//line /usr/local/go/src/net/udpsock.go:245
		// _ = "end of CoverTab[8331]"
	} else {
//line /usr/local/go/src/net/udpsock.go:246
		_go_fuzz_dep_.CoverTab[8332]++
//line /usr/local/go/src/net/udpsock.go:246
		// _ = "end of CoverTab[8332]"
//line /usr/local/go/src/net/udpsock.go:246
	}
//line /usr/local/go/src/net/udpsock.go:246
	// _ = "end of CoverTab[8325]"
//line /usr/local/go/src/net/udpsock.go:246
	_go_fuzz_dep_.CoverTab[8326]++
						return n, err
//line /usr/local/go/src/net/udpsock.go:247
	// _ = "end of CoverTab[8326]"
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
	_go_fuzz_dep_.CoverTab[8333]++
						if !c.ok() {
//line /usr/local/go/src/net/udpsock.go:259
		_go_fuzz_dep_.CoverTab[8336]++
							return 0, 0, syscall.EINVAL
//line /usr/local/go/src/net/udpsock.go:260
		// _ = "end of CoverTab[8336]"
	} else {
//line /usr/local/go/src/net/udpsock.go:261
		_go_fuzz_dep_.CoverTab[8337]++
//line /usr/local/go/src/net/udpsock.go:261
		// _ = "end of CoverTab[8337]"
//line /usr/local/go/src/net/udpsock.go:261
	}
//line /usr/local/go/src/net/udpsock.go:261
	// _ = "end of CoverTab[8333]"
//line /usr/local/go/src/net/udpsock.go:261
	_go_fuzz_dep_.CoverTab[8334]++
						n, oobn, err = c.writeMsg(b, oob, addr)
						if err != nil {
//line /usr/local/go/src/net/udpsock.go:263
		_go_fuzz_dep_.CoverTab[8338]++
							err = &OpError{Op: "write", Net: c.fd.net, Source: c.fd.laddr, Addr: addr.opAddr(), Err: err}
//line /usr/local/go/src/net/udpsock.go:264
		// _ = "end of CoverTab[8338]"
	} else {
//line /usr/local/go/src/net/udpsock.go:265
		_go_fuzz_dep_.CoverTab[8339]++
//line /usr/local/go/src/net/udpsock.go:265
		// _ = "end of CoverTab[8339]"
//line /usr/local/go/src/net/udpsock.go:265
	}
//line /usr/local/go/src/net/udpsock.go:265
	// _ = "end of CoverTab[8334]"
//line /usr/local/go/src/net/udpsock.go:265
	_go_fuzz_dep_.CoverTab[8335]++
						return
//line /usr/local/go/src/net/udpsock.go:266
	// _ = "end of CoverTab[8335]"
}

// WriteMsgUDPAddrPort is like WriteMsgUDP but takes a netip.AddrPort instead of a UDPAddr.
func (c *UDPConn) WriteMsgUDPAddrPort(b, oob []byte, addr netip.AddrPort) (n, oobn int, err error) {
//line /usr/local/go/src/net/udpsock.go:270
	_go_fuzz_dep_.CoverTab[8340]++
						if !c.ok() {
//line /usr/local/go/src/net/udpsock.go:271
		_go_fuzz_dep_.CoverTab[8343]++
							return 0, 0, syscall.EINVAL
//line /usr/local/go/src/net/udpsock.go:272
		// _ = "end of CoverTab[8343]"
	} else {
//line /usr/local/go/src/net/udpsock.go:273
		_go_fuzz_dep_.CoverTab[8344]++
//line /usr/local/go/src/net/udpsock.go:273
		// _ = "end of CoverTab[8344]"
//line /usr/local/go/src/net/udpsock.go:273
	}
//line /usr/local/go/src/net/udpsock.go:273
	// _ = "end of CoverTab[8340]"
//line /usr/local/go/src/net/udpsock.go:273
	_go_fuzz_dep_.CoverTab[8341]++
						n, oobn, err = c.writeMsgAddrPort(b, oob, addr)
						if err != nil {
//line /usr/local/go/src/net/udpsock.go:275
		_go_fuzz_dep_.CoverTab[8345]++
							err = &OpError{Op: "write", Net: c.fd.net, Source: c.fd.laddr, Addr: addrPortUDPAddr{addr}, Err: err}
//line /usr/local/go/src/net/udpsock.go:276
		// _ = "end of CoverTab[8345]"
	} else {
//line /usr/local/go/src/net/udpsock.go:277
		_go_fuzz_dep_.CoverTab[8346]++
//line /usr/local/go/src/net/udpsock.go:277
		// _ = "end of CoverTab[8346]"
//line /usr/local/go/src/net/udpsock.go:277
	}
//line /usr/local/go/src/net/udpsock.go:277
	// _ = "end of CoverTab[8341]"
//line /usr/local/go/src/net/udpsock.go:277
	_go_fuzz_dep_.CoverTab[8342]++
						return
//line /usr/local/go/src/net/udpsock.go:278
	// _ = "end of CoverTab[8342]"
}

func newUDPConn(fd *netFD) *UDPConn {
//line /usr/local/go/src/net/udpsock.go:281
	_go_fuzz_dep_.CoverTab[8347]++
//line /usr/local/go/src/net/udpsock.go:281
	return &UDPConn{conn{fd}}
//line /usr/local/go/src/net/udpsock.go:281
	// _ = "end of CoverTab[8347]"
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
	_go_fuzz_dep_.CoverTab[8348]++
						switch network {
	case "udp", "udp4", "udp6":
//line /usr/local/go/src/net/udpsock.go:292
		_go_fuzz_dep_.CoverTab[8352]++
//line /usr/local/go/src/net/udpsock.go:292
		// _ = "end of CoverTab[8352]"
	default:
//line /usr/local/go/src/net/udpsock.go:293
		_go_fuzz_dep_.CoverTab[8353]++
							return nil, &OpError{Op: "dial", Net: network, Source: laddr.opAddr(), Addr: raddr.opAddr(), Err: UnknownNetworkError(network)}
//line /usr/local/go/src/net/udpsock.go:294
		// _ = "end of CoverTab[8353]"
	}
//line /usr/local/go/src/net/udpsock.go:295
	// _ = "end of CoverTab[8348]"
//line /usr/local/go/src/net/udpsock.go:295
	_go_fuzz_dep_.CoverTab[8349]++
						if raddr == nil {
//line /usr/local/go/src/net/udpsock.go:296
		_go_fuzz_dep_.CoverTab[8354]++
							return nil, &OpError{Op: "dial", Net: network, Source: laddr.opAddr(), Addr: nil, Err: errMissingAddress}
//line /usr/local/go/src/net/udpsock.go:297
		// _ = "end of CoverTab[8354]"
	} else {
//line /usr/local/go/src/net/udpsock.go:298
		_go_fuzz_dep_.CoverTab[8355]++
//line /usr/local/go/src/net/udpsock.go:298
		// _ = "end of CoverTab[8355]"
//line /usr/local/go/src/net/udpsock.go:298
	}
//line /usr/local/go/src/net/udpsock.go:298
	// _ = "end of CoverTab[8349]"
//line /usr/local/go/src/net/udpsock.go:298
	_go_fuzz_dep_.CoverTab[8350]++
						sd := &sysDialer{network: network, address: raddr.String()}
						c, err := sd.dialUDP(context.Background(), laddr, raddr)
						if err != nil {
//line /usr/local/go/src/net/udpsock.go:301
		_go_fuzz_dep_.CoverTab[8356]++
							return nil, &OpError{Op: "dial", Net: network, Source: laddr.opAddr(), Addr: raddr.opAddr(), Err: err}
//line /usr/local/go/src/net/udpsock.go:302
		// _ = "end of CoverTab[8356]"
	} else {
//line /usr/local/go/src/net/udpsock.go:303
		_go_fuzz_dep_.CoverTab[8357]++
//line /usr/local/go/src/net/udpsock.go:303
		// _ = "end of CoverTab[8357]"
//line /usr/local/go/src/net/udpsock.go:303
	}
//line /usr/local/go/src/net/udpsock.go:303
	// _ = "end of CoverTab[8350]"
//line /usr/local/go/src/net/udpsock.go:303
	_go_fuzz_dep_.CoverTab[8351]++
						return c, nil
//line /usr/local/go/src/net/udpsock.go:304
	// _ = "end of CoverTab[8351]"
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
	_go_fuzz_dep_.CoverTab[8358]++
						switch network {
	case "udp", "udp4", "udp6":
//line /usr/local/go/src/net/udpsock.go:318
		_go_fuzz_dep_.CoverTab[8362]++
//line /usr/local/go/src/net/udpsock.go:318
		// _ = "end of CoverTab[8362]"
	default:
//line /usr/local/go/src/net/udpsock.go:319
		_go_fuzz_dep_.CoverTab[8363]++
							return nil, &OpError{Op: "listen", Net: network, Source: nil, Addr: laddr.opAddr(), Err: UnknownNetworkError(network)}
//line /usr/local/go/src/net/udpsock.go:320
		// _ = "end of CoverTab[8363]"
	}
//line /usr/local/go/src/net/udpsock.go:321
	// _ = "end of CoverTab[8358]"
//line /usr/local/go/src/net/udpsock.go:321
	_go_fuzz_dep_.CoverTab[8359]++
						if laddr == nil {
//line /usr/local/go/src/net/udpsock.go:322
		_go_fuzz_dep_.CoverTab[8364]++
							laddr = &UDPAddr{}
//line /usr/local/go/src/net/udpsock.go:323
		// _ = "end of CoverTab[8364]"
	} else {
//line /usr/local/go/src/net/udpsock.go:324
		_go_fuzz_dep_.CoverTab[8365]++
//line /usr/local/go/src/net/udpsock.go:324
		// _ = "end of CoverTab[8365]"
//line /usr/local/go/src/net/udpsock.go:324
	}
//line /usr/local/go/src/net/udpsock.go:324
	// _ = "end of CoverTab[8359]"
//line /usr/local/go/src/net/udpsock.go:324
	_go_fuzz_dep_.CoverTab[8360]++
						sl := &sysListener{network: network, address: laddr.String()}
						c, err := sl.listenUDP(context.Background(), laddr)
						if err != nil {
//line /usr/local/go/src/net/udpsock.go:327
		_go_fuzz_dep_.CoverTab[8366]++
							return nil, &OpError{Op: "listen", Net: network, Source: nil, Addr: laddr.opAddr(), Err: err}
//line /usr/local/go/src/net/udpsock.go:328
		// _ = "end of CoverTab[8366]"
	} else {
//line /usr/local/go/src/net/udpsock.go:329
		_go_fuzz_dep_.CoverTab[8367]++
//line /usr/local/go/src/net/udpsock.go:329
		// _ = "end of CoverTab[8367]"
//line /usr/local/go/src/net/udpsock.go:329
	}
//line /usr/local/go/src/net/udpsock.go:329
	// _ = "end of CoverTab[8360]"
//line /usr/local/go/src/net/udpsock.go:329
	_go_fuzz_dep_.CoverTab[8361]++
						return c, nil
//line /usr/local/go/src/net/udpsock.go:330
	// _ = "end of CoverTab[8361]"
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
	_go_fuzz_dep_.CoverTab[8368]++
						switch network {
	case "udp", "udp4", "udp6":
//line /usr/local/go/src/net/udpsock.go:355
		_go_fuzz_dep_.CoverTab[8372]++
//line /usr/local/go/src/net/udpsock.go:355
		// _ = "end of CoverTab[8372]"
	default:
//line /usr/local/go/src/net/udpsock.go:356
		_go_fuzz_dep_.CoverTab[8373]++
							return nil, &OpError{Op: "listen", Net: network, Source: nil, Addr: gaddr.opAddr(), Err: UnknownNetworkError(network)}
//line /usr/local/go/src/net/udpsock.go:357
		// _ = "end of CoverTab[8373]"
	}
//line /usr/local/go/src/net/udpsock.go:358
	// _ = "end of CoverTab[8368]"
//line /usr/local/go/src/net/udpsock.go:358
	_go_fuzz_dep_.CoverTab[8369]++
						if gaddr == nil || func() bool {
//line /usr/local/go/src/net/udpsock.go:359
		_go_fuzz_dep_.CoverTab[8374]++
//line /usr/local/go/src/net/udpsock.go:359
		return gaddr.IP == nil
//line /usr/local/go/src/net/udpsock.go:359
		// _ = "end of CoverTab[8374]"
//line /usr/local/go/src/net/udpsock.go:359
	}() {
//line /usr/local/go/src/net/udpsock.go:359
		_go_fuzz_dep_.CoverTab[8375]++
							return nil, &OpError{Op: "listen", Net: network, Source: nil, Addr: gaddr.opAddr(), Err: errMissingAddress}
//line /usr/local/go/src/net/udpsock.go:360
		// _ = "end of CoverTab[8375]"
	} else {
//line /usr/local/go/src/net/udpsock.go:361
		_go_fuzz_dep_.CoverTab[8376]++
//line /usr/local/go/src/net/udpsock.go:361
		// _ = "end of CoverTab[8376]"
//line /usr/local/go/src/net/udpsock.go:361
	}
//line /usr/local/go/src/net/udpsock.go:361
	// _ = "end of CoverTab[8369]"
//line /usr/local/go/src/net/udpsock.go:361
	_go_fuzz_dep_.CoverTab[8370]++
						sl := &sysListener{network: network, address: gaddr.String()}
						c, err := sl.listenMulticastUDP(context.Background(), ifi, gaddr)
						if err != nil {
//line /usr/local/go/src/net/udpsock.go:364
		_go_fuzz_dep_.CoverTab[8377]++
							return nil, &OpError{Op: "listen", Net: network, Source: nil, Addr: gaddr.opAddr(), Err: err}
//line /usr/local/go/src/net/udpsock.go:365
		// _ = "end of CoverTab[8377]"
	} else {
//line /usr/local/go/src/net/udpsock.go:366
		_go_fuzz_dep_.CoverTab[8378]++
//line /usr/local/go/src/net/udpsock.go:366
		// _ = "end of CoverTab[8378]"
//line /usr/local/go/src/net/udpsock.go:366
	}
//line /usr/local/go/src/net/udpsock.go:366
	// _ = "end of CoverTab[8370]"
//line /usr/local/go/src/net/udpsock.go:366
	_go_fuzz_dep_.CoverTab[8371]++
						return c, nil
//line /usr/local/go/src/net/udpsock.go:367
	// _ = "end of CoverTab[8371]"
}

//line /usr/local/go/src/net/udpsock.go:368
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/udpsock.go:368
var _ = _go_fuzz_dep_.CoverTab
