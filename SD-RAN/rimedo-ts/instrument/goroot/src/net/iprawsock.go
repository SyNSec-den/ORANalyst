// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/net/iprawsock.go:5
package net

//line /usr/local/go/src/net/iprawsock.go:5
import (
//line /usr/local/go/src/net/iprawsock.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/iprawsock.go:5
)
//line /usr/local/go/src/net/iprawsock.go:5
import (
//line /usr/local/go/src/net/iprawsock.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/iprawsock.go:5
)

import (
	"context"
	"syscall"
)

//line /usr/local/go/src/net/iprawsock.go:30
// IPAddr represents the address of an IP end point.
type IPAddr struct {
	IP	IP
	Zone	string	// IPv6 scoped addressing zone
}

// Network returns the address's network name, "ip".
func (a *IPAddr) Network() string {
//line /usr/local/go/src/net/iprawsock.go:37
	_go_fuzz_dep_.CoverTab[14659]++
//line /usr/local/go/src/net/iprawsock.go:37
	return "ip"
//line /usr/local/go/src/net/iprawsock.go:37
	// _ = "end of CoverTab[14659]"
//line /usr/local/go/src/net/iprawsock.go:37
}

func (a *IPAddr) String() string {
//line /usr/local/go/src/net/iprawsock.go:39
	_go_fuzz_dep_.CoverTab[14660]++
						if a == nil {
//line /usr/local/go/src/net/iprawsock.go:40
		_go_fuzz_dep_.CoverTab[14663]++
							return "<nil>"
//line /usr/local/go/src/net/iprawsock.go:41
		// _ = "end of CoverTab[14663]"
	} else {
//line /usr/local/go/src/net/iprawsock.go:42
		_go_fuzz_dep_.CoverTab[14664]++
//line /usr/local/go/src/net/iprawsock.go:42
		// _ = "end of CoverTab[14664]"
//line /usr/local/go/src/net/iprawsock.go:42
	}
//line /usr/local/go/src/net/iprawsock.go:42
	// _ = "end of CoverTab[14660]"
//line /usr/local/go/src/net/iprawsock.go:42
	_go_fuzz_dep_.CoverTab[14661]++
						ip := ipEmptyString(a.IP)
						if a.Zone != "" {
//line /usr/local/go/src/net/iprawsock.go:44
		_go_fuzz_dep_.CoverTab[14665]++
							return ip + "%" + a.Zone
//line /usr/local/go/src/net/iprawsock.go:45
		// _ = "end of CoverTab[14665]"
	} else {
//line /usr/local/go/src/net/iprawsock.go:46
		_go_fuzz_dep_.CoverTab[14666]++
//line /usr/local/go/src/net/iprawsock.go:46
		// _ = "end of CoverTab[14666]"
//line /usr/local/go/src/net/iprawsock.go:46
	}
//line /usr/local/go/src/net/iprawsock.go:46
	// _ = "end of CoverTab[14661]"
//line /usr/local/go/src/net/iprawsock.go:46
	_go_fuzz_dep_.CoverTab[14662]++
						return ip
//line /usr/local/go/src/net/iprawsock.go:47
	// _ = "end of CoverTab[14662]"
}

func (a *IPAddr) isWildcard() bool {
//line /usr/local/go/src/net/iprawsock.go:50
	_go_fuzz_dep_.CoverTab[14667]++
						if a == nil || func() bool {
//line /usr/local/go/src/net/iprawsock.go:51
		_go_fuzz_dep_.CoverTab[14669]++
//line /usr/local/go/src/net/iprawsock.go:51
		return a.IP == nil
//line /usr/local/go/src/net/iprawsock.go:51
		// _ = "end of CoverTab[14669]"
//line /usr/local/go/src/net/iprawsock.go:51
	}() {
//line /usr/local/go/src/net/iprawsock.go:51
		_go_fuzz_dep_.CoverTab[14670]++
							return true
//line /usr/local/go/src/net/iprawsock.go:52
		// _ = "end of CoverTab[14670]"
	} else {
//line /usr/local/go/src/net/iprawsock.go:53
		_go_fuzz_dep_.CoverTab[14671]++
//line /usr/local/go/src/net/iprawsock.go:53
		// _ = "end of CoverTab[14671]"
//line /usr/local/go/src/net/iprawsock.go:53
	}
//line /usr/local/go/src/net/iprawsock.go:53
	// _ = "end of CoverTab[14667]"
//line /usr/local/go/src/net/iprawsock.go:53
	_go_fuzz_dep_.CoverTab[14668]++
						return a.IP.IsUnspecified()
//line /usr/local/go/src/net/iprawsock.go:54
	// _ = "end of CoverTab[14668]"
}

func (a *IPAddr) opAddr() Addr {
//line /usr/local/go/src/net/iprawsock.go:57
	_go_fuzz_dep_.CoverTab[14672]++
						if a == nil {
//line /usr/local/go/src/net/iprawsock.go:58
		_go_fuzz_dep_.CoverTab[14674]++
							return nil
//line /usr/local/go/src/net/iprawsock.go:59
		// _ = "end of CoverTab[14674]"
	} else {
//line /usr/local/go/src/net/iprawsock.go:60
		_go_fuzz_dep_.CoverTab[14675]++
//line /usr/local/go/src/net/iprawsock.go:60
		// _ = "end of CoverTab[14675]"
//line /usr/local/go/src/net/iprawsock.go:60
	}
//line /usr/local/go/src/net/iprawsock.go:60
	// _ = "end of CoverTab[14672]"
//line /usr/local/go/src/net/iprawsock.go:60
	_go_fuzz_dep_.CoverTab[14673]++
						return a
//line /usr/local/go/src/net/iprawsock.go:61
	// _ = "end of CoverTab[14673]"
}

// ResolveIPAddr returns an address of IP end point.
//line /usr/local/go/src/net/iprawsock.go:64
//
//line /usr/local/go/src/net/iprawsock.go:64
// The network must be an IP network name.
//line /usr/local/go/src/net/iprawsock.go:64
//
//line /usr/local/go/src/net/iprawsock.go:64
// If the host in the address parameter is not a literal IP address,
//line /usr/local/go/src/net/iprawsock.go:64
// ResolveIPAddr resolves the address to an address of IP end point.
//line /usr/local/go/src/net/iprawsock.go:64
// Otherwise, it parses the address as a literal IP address.
//line /usr/local/go/src/net/iprawsock.go:64
// The address parameter can use a host name, but this is not
//line /usr/local/go/src/net/iprawsock.go:64
// recommended, because it will return at most one of the host name's
//line /usr/local/go/src/net/iprawsock.go:64
// IP addresses.
//line /usr/local/go/src/net/iprawsock.go:64
//
//line /usr/local/go/src/net/iprawsock.go:64
// See func Dial for a description of the network and address
//line /usr/local/go/src/net/iprawsock.go:64
// parameters.
//line /usr/local/go/src/net/iprawsock.go:77
func ResolveIPAddr(network, address string) (*IPAddr, error) {
//line /usr/local/go/src/net/iprawsock.go:77
	_go_fuzz_dep_.CoverTab[14676]++
						if network == "" {
//line /usr/local/go/src/net/iprawsock.go:78
		_go_fuzz_dep_.CoverTab[14681]++
							network = "ip"
//line /usr/local/go/src/net/iprawsock.go:79
		// _ = "end of CoverTab[14681]"
	} else {
//line /usr/local/go/src/net/iprawsock.go:80
		_go_fuzz_dep_.CoverTab[14682]++
//line /usr/local/go/src/net/iprawsock.go:80
		// _ = "end of CoverTab[14682]"
//line /usr/local/go/src/net/iprawsock.go:80
	}
//line /usr/local/go/src/net/iprawsock.go:80
	// _ = "end of CoverTab[14676]"
//line /usr/local/go/src/net/iprawsock.go:80
	_go_fuzz_dep_.CoverTab[14677]++
						afnet, _, err := parseNetwork(context.Background(), network, false)
						if err != nil {
//line /usr/local/go/src/net/iprawsock.go:82
		_go_fuzz_dep_.CoverTab[14683]++
							return nil, err
//line /usr/local/go/src/net/iprawsock.go:83
		// _ = "end of CoverTab[14683]"
	} else {
//line /usr/local/go/src/net/iprawsock.go:84
		_go_fuzz_dep_.CoverTab[14684]++
//line /usr/local/go/src/net/iprawsock.go:84
		// _ = "end of CoverTab[14684]"
//line /usr/local/go/src/net/iprawsock.go:84
	}
//line /usr/local/go/src/net/iprawsock.go:84
	// _ = "end of CoverTab[14677]"
//line /usr/local/go/src/net/iprawsock.go:84
	_go_fuzz_dep_.CoverTab[14678]++
						switch afnet {
	case "ip", "ip4", "ip6":
//line /usr/local/go/src/net/iprawsock.go:86
		_go_fuzz_dep_.CoverTab[14685]++
//line /usr/local/go/src/net/iprawsock.go:86
		// _ = "end of CoverTab[14685]"
	default:
//line /usr/local/go/src/net/iprawsock.go:87
		_go_fuzz_dep_.CoverTab[14686]++
							return nil, UnknownNetworkError(network)
//line /usr/local/go/src/net/iprawsock.go:88
		// _ = "end of CoverTab[14686]"
	}
//line /usr/local/go/src/net/iprawsock.go:89
	// _ = "end of CoverTab[14678]"
//line /usr/local/go/src/net/iprawsock.go:89
	_go_fuzz_dep_.CoverTab[14679]++
						addrs, err := DefaultResolver.internetAddrList(context.Background(), afnet, address)
						if err != nil {
//line /usr/local/go/src/net/iprawsock.go:91
		_go_fuzz_dep_.CoverTab[14687]++
							return nil, err
//line /usr/local/go/src/net/iprawsock.go:92
		// _ = "end of CoverTab[14687]"
	} else {
//line /usr/local/go/src/net/iprawsock.go:93
		_go_fuzz_dep_.CoverTab[14688]++
//line /usr/local/go/src/net/iprawsock.go:93
		// _ = "end of CoverTab[14688]"
//line /usr/local/go/src/net/iprawsock.go:93
	}
//line /usr/local/go/src/net/iprawsock.go:93
	// _ = "end of CoverTab[14679]"
//line /usr/local/go/src/net/iprawsock.go:93
	_go_fuzz_dep_.CoverTab[14680]++
						return addrs.forResolve(network, address).(*IPAddr), nil
//line /usr/local/go/src/net/iprawsock.go:94
	// _ = "end of CoverTab[14680]"
}

// IPConn is the implementation of the Conn and PacketConn interfaces
//line /usr/local/go/src/net/iprawsock.go:97
// for IP network connections.
//line /usr/local/go/src/net/iprawsock.go:99
type IPConn struct {
	conn
}

// SyscallConn returns a raw network connection.
//line /usr/local/go/src/net/iprawsock.go:103
// This implements the syscall.Conn interface.
//line /usr/local/go/src/net/iprawsock.go:105
func (c *IPConn) SyscallConn() (syscall.RawConn, error) {
//line /usr/local/go/src/net/iprawsock.go:105
	_go_fuzz_dep_.CoverTab[14689]++
						if !c.ok() {
//line /usr/local/go/src/net/iprawsock.go:106
		_go_fuzz_dep_.CoverTab[14691]++
							return nil, syscall.EINVAL
//line /usr/local/go/src/net/iprawsock.go:107
		// _ = "end of CoverTab[14691]"
	} else {
//line /usr/local/go/src/net/iprawsock.go:108
		_go_fuzz_dep_.CoverTab[14692]++
//line /usr/local/go/src/net/iprawsock.go:108
		// _ = "end of CoverTab[14692]"
//line /usr/local/go/src/net/iprawsock.go:108
	}
//line /usr/local/go/src/net/iprawsock.go:108
	// _ = "end of CoverTab[14689]"
//line /usr/local/go/src/net/iprawsock.go:108
	_go_fuzz_dep_.CoverTab[14690]++
						return newRawConn(c.fd)
//line /usr/local/go/src/net/iprawsock.go:109
	// _ = "end of CoverTab[14690]"
}

// ReadFromIP acts like ReadFrom but returns an IPAddr.
func (c *IPConn) ReadFromIP(b []byte) (int, *IPAddr, error) {
//line /usr/local/go/src/net/iprawsock.go:113
	_go_fuzz_dep_.CoverTab[14693]++
						if !c.ok() {
//line /usr/local/go/src/net/iprawsock.go:114
		_go_fuzz_dep_.CoverTab[14696]++
							return 0, nil, syscall.EINVAL
//line /usr/local/go/src/net/iprawsock.go:115
		// _ = "end of CoverTab[14696]"
	} else {
//line /usr/local/go/src/net/iprawsock.go:116
		_go_fuzz_dep_.CoverTab[14697]++
//line /usr/local/go/src/net/iprawsock.go:116
		// _ = "end of CoverTab[14697]"
//line /usr/local/go/src/net/iprawsock.go:116
	}
//line /usr/local/go/src/net/iprawsock.go:116
	// _ = "end of CoverTab[14693]"
//line /usr/local/go/src/net/iprawsock.go:116
	_go_fuzz_dep_.CoverTab[14694]++
						n, addr, err := c.readFrom(b)
						if err != nil {
//line /usr/local/go/src/net/iprawsock.go:118
		_go_fuzz_dep_.CoverTab[14698]++
							err = &OpError{Op: "read", Net: c.fd.net, Source: c.fd.laddr, Addr: c.fd.raddr, Err: err}
//line /usr/local/go/src/net/iprawsock.go:119
		// _ = "end of CoverTab[14698]"
	} else {
//line /usr/local/go/src/net/iprawsock.go:120
		_go_fuzz_dep_.CoverTab[14699]++
//line /usr/local/go/src/net/iprawsock.go:120
		// _ = "end of CoverTab[14699]"
//line /usr/local/go/src/net/iprawsock.go:120
	}
//line /usr/local/go/src/net/iprawsock.go:120
	// _ = "end of CoverTab[14694]"
//line /usr/local/go/src/net/iprawsock.go:120
	_go_fuzz_dep_.CoverTab[14695]++
						return n, addr, err
//line /usr/local/go/src/net/iprawsock.go:121
	// _ = "end of CoverTab[14695]"
}

// ReadFrom implements the PacketConn ReadFrom method.
func (c *IPConn) ReadFrom(b []byte) (int, Addr, error) {
//line /usr/local/go/src/net/iprawsock.go:125
	_go_fuzz_dep_.CoverTab[14700]++
						if !c.ok() {
//line /usr/local/go/src/net/iprawsock.go:126
		_go_fuzz_dep_.CoverTab[14704]++
							return 0, nil, syscall.EINVAL
//line /usr/local/go/src/net/iprawsock.go:127
		// _ = "end of CoverTab[14704]"
	} else {
//line /usr/local/go/src/net/iprawsock.go:128
		_go_fuzz_dep_.CoverTab[14705]++
//line /usr/local/go/src/net/iprawsock.go:128
		// _ = "end of CoverTab[14705]"
//line /usr/local/go/src/net/iprawsock.go:128
	}
//line /usr/local/go/src/net/iprawsock.go:128
	// _ = "end of CoverTab[14700]"
//line /usr/local/go/src/net/iprawsock.go:128
	_go_fuzz_dep_.CoverTab[14701]++
						n, addr, err := c.readFrom(b)
						if err != nil {
//line /usr/local/go/src/net/iprawsock.go:130
		_go_fuzz_dep_.CoverTab[14706]++
							err = &OpError{Op: "read", Net: c.fd.net, Source: c.fd.laddr, Addr: c.fd.raddr, Err: err}
//line /usr/local/go/src/net/iprawsock.go:131
		// _ = "end of CoverTab[14706]"
	} else {
//line /usr/local/go/src/net/iprawsock.go:132
		_go_fuzz_dep_.CoverTab[14707]++
//line /usr/local/go/src/net/iprawsock.go:132
		// _ = "end of CoverTab[14707]"
//line /usr/local/go/src/net/iprawsock.go:132
	}
//line /usr/local/go/src/net/iprawsock.go:132
	// _ = "end of CoverTab[14701]"
//line /usr/local/go/src/net/iprawsock.go:132
	_go_fuzz_dep_.CoverTab[14702]++
						if addr == nil {
//line /usr/local/go/src/net/iprawsock.go:133
		_go_fuzz_dep_.CoverTab[14708]++
							return n, nil, err
//line /usr/local/go/src/net/iprawsock.go:134
		// _ = "end of CoverTab[14708]"
	} else {
//line /usr/local/go/src/net/iprawsock.go:135
		_go_fuzz_dep_.CoverTab[14709]++
//line /usr/local/go/src/net/iprawsock.go:135
		// _ = "end of CoverTab[14709]"
//line /usr/local/go/src/net/iprawsock.go:135
	}
//line /usr/local/go/src/net/iprawsock.go:135
	// _ = "end of CoverTab[14702]"
//line /usr/local/go/src/net/iprawsock.go:135
	_go_fuzz_dep_.CoverTab[14703]++
						return n, addr, err
//line /usr/local/go/src/net/iprawsock.go:136
	// _ = "end of CoverTab[14703]"
}

// ReadMsgIP reads a message from c, copying the payload into b and
//line /usr/local/go/src/net/iprawsock.go:139
// the associated out-of-band data into oob. It returns the number of
//line /usr/local/go/src/net/iprawsock.go:139
// bytes copied into b, the number of bytes copied into oob, the flags
//line /usr/local/go/src/net/iprawsock.go:139
// that were set on the message and the source address of the message.
//line /usr/local/go/src/net/iprawsock.go:139
//
//line /usr/local/go/src/net/iprawsock.go:139
// The packages golang.org/x/net/ipv4 and golang.org/x/net/ipv6 can be
//line /usr/local/go/src/net/iprawsock.go:139
// used to manipulate IP-level socket options in oob.
//line /usr/local/go/src/net/iprawsock.go:146
func (c *IPConn) ReadMsgIP(b, oob []byte) (n, oobn, flags int, addr *IPAddr, err error) {
//line /usr/local/go/src/net/iprawsock.go:146
	_go_fuzz_dep_.CoverTab[14710]++
						if !c.ok() {
//line /usr/local/go/src/net/iprawsock.go:147
		_go_fuzz_dep_.CoverTab[14713]++
							return 0, 0, 0, nil, syscall.EINVAL
//line /usr/local/go/src/net/iprawsock.go:148
		// _ = "end of CoverTab[14713]"
	} else {
//line /usr/local/go/src/net/iprawsock.go:149
		_go_fuzz_dep_.CoverTab[14714]++
//line /usr/local/go/src/net/iprawsock.go:149
		// _ = "end of CoverTab[14714]"
//line /usr/local/go/src/net/iprawsock.go:149
	}
//line /usr/local/go/src/net/iprawsock.go:149
	// _ = "end of CoverTab[14710]"
//line /usr/local/go/src/net/iprawsock.go:149
	_go_fuzz_dep_.CoverTab[14711]++
						n, oobn, flags, addr, err = c.readMsg(b, oob)
						if err != nil {
//line /usr/local/go/src/net/iprawsock.go:151
		_go_fuzz_dep_.CoverTab[14715]++
							err = &OpError{Op: "read", Net: c.fd.net, Source: c.fd.laddr, Addr: c.fd.raddr, Err: err}
//line /usr/local/go/src/net/iprawsock.go:152
		// _ = "end of CoverTab[14715]"
	} else {
//line /usr/local/go/src/net/iprawsock.go:153
		_go_fuzz_dep_.CoverTab[14716]++
//line /usr/local/go/src/net/iprawsock.go:153
		// _ = "end of CoverTab[14716]"
//line /usr/local/go/src/net/iprawsock.go:153
	}
//line /usr/local/go/src/net/iprawsock.go:153
	// _ = "end of CoverTab[14711]"
//line /usr/local/go/src/net/iprawsock.go:153
	_go_fuzz_dep_.CoverTab[14712]++
						return
//line /usr/local/go/src/net/iprawsock.go:154
	// _ = "end of CoverTab[14712]"
}

// WriteToIP acts like WriteTo but takes an IPAddr.
func (c *IPConn) WriteToIP(b []byte, addr *IPAddr) (int, error) {
//line /usr/local/go/src/net/iprawsock.go:158
	_go_fuzz_dep_.CoverTab[14717]++
						if !c.ok() {
//line /usr/local/go/src/net/iprawsock.go:159
		_go_fuzz_dep_.CoverTab[14720]++
							return 0, syscall.EINVAL
//line /usr/local/go/src/net/iprawsock.go:160
		// _ = "end of CoverTab[14720]"
	} else {
//line /usr/local/go/src/net/iprawsock.go:161
		_go_fuzz_dep_.CoverTab[14721]++
//line /usr/local/go/src/net/iprawsock.go:161
		// _ = "end of CoverTab[14721]"
//line /usr/local/go/src/net/iprawsock.go:161
	}
//line /usr/local/go/src/net/iprawsock.go:161
	// _ = "end of CoverTab[14717]"
//line /usr/local/go/src/net/iprawsock.go:161
	_go_fuzz_dep_.CoverTab[14718]++
						n, err := c.writeTo(b, addr)
						if err != nil {
//line /usr/local/go/src/net/iprawsock.go:163
		_go_fuzz_dep_.CoverTab[14722]++
							err = &OpError{Op: "write", Net: c.fd.net, Source: c.fd.laddr, Addr: addr.opAddr(), Err: err}
//line /usr/local/go/src/net/iprawsock.go:164
		// _ = "end of CoverTab[14722]"
	} else {
//line /usr/local/go/src/net/iprawsock.go:165
		_go_fuzz_dep_.CoverTab[14723]++
//line /usr/local/go/src/net/iprawsock.go:165
		// _ = "end of CoverTab[14723]"
//line /usr/local/go/src/net/iprawsock.go:165
	}
//line /usr/local/go/src/net/iprawsock.go:165
	// _ = "end of CoverTab[14718]"
//line /usr/local/go/src/net/iprawsock.go:165
	_go_fuzz_dep_.CoverTab[14719]++
						return n, err
//line /usr/local/go/src/net/iprawsock.go:166
	// _ = "end of CoverTab[14719]"
}

// WriteTo implements the PacketConn WriteTo method.
func (c *IPConn) WriteTo(b []byte, addr Addr) (int, error) {
//line /usr/local/go/src/net/iprawsock.go:170
	_go_fuzz_dep_.CoverTab[14724]++
						if !c.ok() {
//line /usr/local/go/src/net/iprawsock.go:171
		_go_fuzz_dep_.CoverTab[14728]++
							return 0, syscall.EINVAL
//line /usr/local/go/src/net/iprawsock.go:172
		// _ = "end of CoverTab[14728]"
	} else {
//line /usr/local/go/src/net/iprawsock.go:173
		_go_fuzz_dep_.CoverTab[14729]++
//line /usr/local/go/src/net/iprawsock.go:173
		// _ = "end of CoverTab[14729]"
//line /usr/local/go/src/net/iprawsock.go:173
	}
//line /usr/local/go/src/net/iprawsock.go:173
	// _ = "end of CoverTab[14724]"
//line /usr/local/go/src/net/iprawsock.go:173
	_go_fuzz_dep_.CoverTab[14725]++
						a, ok := addr.(*IPAddr)
						if !ok {
//line /usr/local/go/src/net/iprawsock.go:175
		_go_fuzz_dep_.CoverTab[14730]++
							return 0, &OpError{Op: "write", Net: c.fd.net, Source: c.fd.laddr, Addr: addr, Err: syscall.EINVAL}
//line /usr/local/go/src/net/iprawsock.go:176
		// _ = "end of CoverTab[14730]"
	} else {
//line /usr/local/go/src/net/iprawsock.go:177
		_go_fuzz_dep_.CoverTab[14731]++
//line /usr/local/go/src/net/iprawsock.go:177
		// _ = "end of CoverTab[14731]"
//line /usr/local/go/src/net/iprawsock.go:177
	}
//line /usr/local/go/src/net/iprawsock.go:177
	// _ = "end of CoverTab[14725]"
//line /usr/local/go/src/net/iprawsock.go:177
	_go_fuzz_dep_.CoverTab[14726]++
						n, err := c.writeTo(b, a)
						if err != nil {
//line /usr/local/go/src/net/iprawsock.go:179
		_go_fuzz_dep_.CoverTab[14732]++
							err = &OpError{Op: "write", Net: c.fd.net, Source: c.fd.laddr, Addr: a.opAddr(), Err: err}
//line /usr/local/go/src/net/iprawsock.go:180
		// _ = "end of CoverTab[14732]"
	} else {
//line /usr/local/go/src/net/iprawsock.go:181
		_go_fuzz_dep_.CoverTab[14733]++
//line /usr/local/go/src/net/iprawsock.go:181
		// _ = "end of CoverTab[14733]"
//line /usr/local/go/src/net/iprawsock.go:181
	}
//line /usr/local/go/src/net/iprawsock.go:181
	// _ = "end of CoverTab[14726]"
//line /usr/local/go/src/net/iprawsock.go:181
	_go_fuzz_dep_.CoverTab[14727]++
						return n, err
//line /usr/local/go/src/net/iprawsock.go:182
	// _ = "end of CoverTab[14727]"
}

// WriteMsgIP writes a message to addr via c, copying the payload from
//line /usr/local/go/src/net/iprawsock.go:185
// b and the associated out-of-band data from oob. It returns the
//line /usr/local/go/src/net/iprawsock.go:185
// number of payload and out-of-band bytes written.
//line /usr/local/go/src/net/iprawsock.go:185
//
//line /usr/local/go/src/net/iprawsock.go:185
// The packages golang.org/x/net/ipv4 and golang.org/x/net/ipv6 can be
//line /usr/local/go/src/net/iprawsock.go:185
// used to manipulate IP-level socket options in oob.
//line /usr/local/go/src/net/iprawsock.go:191
func (c *IPConn) WriteMsgIP(b, oob []byte, addr *IPAddr) (n, oobn int, err error) {
//line /usr/local/go/src/net/iprawsock.go:191
	_go_fuzz_dep_.CoverTab[14734]++
						if !c.ok() {
//line /usr/local/go/src/net/iprawsock.go:192
		_go_fuzz_dep_.CoverTab[14737]++
							return 0, 0, syscall.EINVAL
//line /usr/local/go/src/net/iprawsock.go:193
		// _ = "end of CoverTab[14737]"
	} else {
//line /usr/local/go/src/net/iprawsock.go:194
		_go_fuzz_dep_.CoverTab[14738]++
//line /usr/local/go/src/net/iprawsock.go:194
		// _ = "end of CoverTab[14738]"
//line /usr/local/go/src/net/iprawsock.go:194
	}
//line /usr/local/go/src/net/iprawsock.go:194
	// _ = "end of CoverTab[14734]"
//line /usr/local/go/src/net/iprawsock.go:194
	_go_fuzz_dep_.CoverTab[14735]++
						n, oobn, err = c.writeMsg(b, oob, addr)
						if err != nil {
//line /usr/local/go/src/net/iprawsock.go:196
		_go_fuzz_dep_.CoverTab[14739]++
							err = &OpError{Op: "write", Net: c.fd.net, Source: c.fd.laddr, Addr: addr.opAddr(), Err: err}
//line /usr/local/go/src/net/iprawsock.go:197
		// _ = "end of CoverTab[14739]"
	} else {
//line /usr/local/go/src/net/iprawsock.go:198
		_go_fuzz_dep_.CoverTab[14740]++
//line /usr/local/go/src/net/iprawsock.go:198
		// _ = "end of CoverTab[14740]"
//line /usr/local/go/src/net/iprawsock.go:198
	}
//line /usr/local/go/src/net/iprawsock.go:198
	// _ = "end of CoverTab[14735]"
//line /usr/local/go/src/net/iprawsock.go:198
	_go_fuzz_dep_.CoverTab[14736]++
						return
//line /usr/local/go/src/net/iprawsock.go:199
	// _ = "end of CoverTab[14736]"
}

func newIPConn(fd *netFD) *IPConn {
//line /usr/local/go/src/net/iprawsock.go:202
	_go_fuzz_dep_.CoverTab[14741]++
//line /usr/local/go/src/net/iprawsock.go:202
	return &IPConn{conn{fd}}
//line /usr/local/go/src/net/iprawsock.go:202
	// _ = "end of CoverTab[14741]"
//line /usr/local/go/src/net/iprawsock.go:202
}

// DialIP acts like Dial for IP networks.
//line /usr/local/go/src/net/iprawsock.go:204
//
//line /usr/local/go/src/net/iprawsock.go:204
// The network must be an IP network name; see func Dial for details.
//line /usr/local/go/src/net/iprawsock.go:204
//
//line /usr/local/go/src/net/iprawsock.go:204
// If laddr is nil, a local address is automatically chosen.
//line /usr/local/go/src/net/iprawsock.go:204
// If the IP field of raddr is nil or an unspecified IP address, the
//line /usr/local/go/src/net/iprawsock.go:204
// local system is assumed.
//line /usr/local/go/src/net/iprawsock.go:211
func DialIP(network string, laddr, raddr *IPAddr) (*IPConn, error) {
//line /usr/local/go/src/net/iprawsock.go:211
	_go_fuzz_dep_.CoverTab[14742]++
						if raddr == nil {
//line /usr/local/go/src/net/iprawsock.go:212
		_go_fuzz_dep_.CoverTab[14745]++
							return nil, &OpError{Op: "dial", Net: network, Source: laddr.opAddr(), Addr: nil, Err: errMissingAddress}
//line /usr/local/go/src/net/iprawsock.go:213
		// _ = "end of CoverTab[14745]"
	} else {
//line /usr/local/go/src/net/iprawsock.go:214
		_go_fuzz_dep_.CoverTab[14746]++
//line /usr/local/go/src/net/iprawsock.go:214
		// _ = "end of CoverTab[14746]"
//line /usr/local/go/src/net/iprawsock.go:214
	}
//line /usr/local/go/src/net/iprawsock.go:214
	// _ = "end of CoverTab[14742]"
//line /usr/local/go/src/net/iprawsock.go:214
	_go_fuzz_dep_.CoverTab[14743]++
						sd := &sysDialer{network: network, address: raddr.String()}
						c, err := sd.dialIP(context.Background(), laddr, raddr)
						if err != nil {
//line /usr/local/go/src/net/iprawsock.go:217
		_go_fuzz_dep_.CoverTab[14747]++
							return nil, &OpError{Op: "dial", Net: network, Source: laddr.opAddr(), Addr: raddr.opAddr(), Err: err}
//line /usr/local/go/src/net/iprawsock.go:218
		// _ = "end of CoverTab[14747]"
	} else {
//line /usr/local/go/src/net/iprawsock.go:219
		_go_fuzz_dep_.CoverTab[14748]++
//line /usr/local/go/src/net/iprawsock.go:219
		// _ = "end of CoverTab[14748]"
//line /usr/local/go/src/net/iprawsock.go:219
	}
//line /usr/local/go/src/net/iprawsock.go:219
	// _ = "end of CoverTab[14743]"
//line /usr/local/go/src/net/iprawsock.go:219
	_go_fuzz_dep_.CoverTab[14744]++
						return c, nil
//line /usr/local/go/src/net/iprawsock.go:220
	// _ = "end of CoverTab[14744]"
}

// ListenIP acts like ListenPacket for IP networks.
//line /usr/local/go/src/net/iprawsock.go:223
//
//line /usr/local/go/src/net/iprawsock.go:223
// The network must be an IP network name; see func Dial for details.
//line /usr/local/go/src/net/iprawsock.go:223
//
//line /usr/local/go/src/net/iprawsock.go:223
// If the IP field of laddr is nil or an unspecified IP address,
//line /usr/local/go/src/net/iprawsock.go:223
// ListenIP listens on all available IP addresses of the local system
//line /usr/local/go/src/net/iprawsock.go:223
// except multicast IP addresses.
//line /usr/local/go/src/net/iprawsock.go:230
func ListenIP(network string, laddr *IPAddr) (*IPConn, error) {
//line /usr/local/go/src/net/iprawsock.go:230
	_go_fuzz_dep_.CoverTab[14749]++
						if laddr == nil {
//line /usr/local/go/src/net/iprawsock.go:231
		_go_fuzz_dep_.CoverTab[14752]++
							laddr = &IPAddr{}
//line /usr/local/go/src/net/iprawsock.go:232
		// _ = "end of CoverTab[14752]"
	} else {
//line /usr/local/go/src/net/iprawsock.go:233
		_go_fuzz_dep_.CoverTab[14753]++
//line /usr/local/go/src/net/iprawsock.go:233
		// _ = "end of CoverTab[14753]"
//line /usr/local/go/src/net/iprawsock.go:233
	}
//line /usr/local/go/src/net/iprawsock.go:233
	// _ = "end of CoverTab[14749]"
//line /usr/local/go/src/net/iprawsock.go:233
	_go_fuzz_dep_.CoverTab[14750]++
						sl := &sysListener{network: network, address: laddr.String()}
						c, err := sl.listenIP(context.Background(), laddr)
						if err != nil {
//line /usr/local/go/src/net/iprawsock.go:236
		_go_fuzz_dep_.CoverTab[14754]++
							return nil, &OpError{Op: "listen", Net: network, Source: nil, Addr: laddr.opAddr(), Err: err}
//line /usr/local/go/src/net/iprawsock.go:237
		// _ = "end of CoverTab[14754]"
	} else {
//line /usr/local/go/src/net/iprawsock.go:238
		_go_fuzz_dep_.CoverTab[14755]++
//line /usr/local/go/src/net/iprawsock.go:238
		// _ = "end of CoverTab[14755]"
//line /usr/local/go/src/net/iprawsock.go:238
	}
//line /usr/local/go/src/net/iprawsock.go:238
	// _ = "end of CoverTab[14750]"
//line /usr/local/go/src/net/iprawsock.go:238
	_go_fuzz_dep_.CoverTab[14751]++
						return c, nil
//line /usr/local/go/src/net/iprawsock.go:239
	// _ = "end of CoverTab[14751]"
}

//line /usr/local/go/src/net/iprawsock.go:240
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/iprawsock.go:240
var _ = _go_fuzz_dep_.CoverTab
