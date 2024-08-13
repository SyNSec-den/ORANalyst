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
	_go_fuzz_dep_.CoverTab[6269]++
//line /usr/local/go/src/net/iprawsock.go:37
	return "ip"
//line /usr/local/go/src/net/iprawsock.go:37
	// _ = "end of CoverTab[6269]"
//line /usr/local/go/src/net/iprawsock.go:37
}

func (a *IPAddr) String() string {
//line /usr/local/go/src/net/iprawsock.go:39
	_go_fuzz_dep_.CoverTab[6270]++
						if a == nil {
//line /usr/local/go/src/net/iprawsock.go:40
		_go_fuzz_dep_.CoverTab[6273]++
							return "<nil>"
//line /usr/local/go/src/net/iprawsock.go:41
		// _ = "end of CoverTab[6273]"
	} else {
//line /usr/local/go/src/net/iprawsock.go:42
		_go_fuzz_dep_.CoverTab[6274]++
//line /usr/local/go/src/net/iprawsock.go:42
		// _ = "end of CoverTab[6274]"
//line /usr/local/go/src/net/iprawsock.go:42
	}
//line /usr/local/go/src/net/iprawsock.go:42
	// _ = "end of CoverTab[6270]"
//line /usr/local/go/src/net/iprawsock.go:42
	_go_fuzz_dep_.CoverTab[6271]++
						ip := ipEmptyString(a.IP)
						if a.Zone != "" {
//line /usr/local/go/src/net/iprawsock.go:44
		_go_fuzz_dep_.CoverTab[6275]++
							return ip + "%" + a.Zone
//line /usr/local/go/src/net/iprawsock.go:45
		// _ = "end of CoverTab[6275]"
	} else {
//line /usr/local/go/src/net/iprawsock.go:46
		_go_fuzz_dep_.CoverTab[6276]++
//line /usr/local/go/src/net/iprawsock.go:46
		// _ = "end of CoverTab[6276]"
//line /usr/local/go/src/net/iprawsock.go:46
	}
//line /usr/local/go/src/net/iprawsock.go:46
	// _ = "end of CoverTab[6271]"
//line /usr/local/go/src/net/iprawsock.go:46
	_go_fuzz_dep_.CoverTab[6272]++
						return ip
//line /usr/local/go/src/net/iprawsock.go:47
	// _ = "end of CoverTab[6272]"
}

func (a *IPAddr) isWildcard() bool {
//line /usr/local/go/src/net/iprawsock.go:50
	_go_fuzz_dep_.CoverTab[6277]++
						if a == nil || func() bool {
//line /usr/local/go/src/net/iprawsock.go:51
		_go_fuzz_dep_.CoverTab[6279]++
//line /usr/local/go/src/net/iprawsock.go:51
		return a.IP == nil
//line /usr/local/go/src/net/iprawsock.go:51
		// _ = "end of CoverTab[6279]"
//line /usr/local/go/src/net/iprawsock.go:51
	}() {
//line /usr/local/go/src/net/iprawsock.go:51
		_go_fuzz_dep_.CoverTab[6280]++
							return true
//line /usr/local/go/src/net/iprawsock.go:52
		// _ = "end of CoverTab[6280]"
	} else {
//line /usr/local/go/src/net/iprawsock.go:53
		_go_fuzz_dep_.CoverTab[6281]++
//line /usr/local/go/src/net/iprawsock.go:53
		// _ = "end of CoverTab[6281]"
//line /usr/local/go/src/net/iprawsock.go:53
	}
//line /usr/local/go/src/net/iprawsock.go:53
	// _ = "end of CoverTab[6277]"
//line /usr/local/go/src/net/iprawsock.go:53
	_go_fuzz_dep_.CoverTab[6278]++
						return a.IP.IsUnspecified()
//line /usr/local/go/src/net/iprawsock.go:54
	// _ = "end of CoverTab[6278]"
}

func (a *IPAddr) opAddr() Addr {
//line /usr/local/go/src/net/iprawsock.go:57
	_go_fuzz_dep_.CoverTab[6282]++
						if a == nil {
//line /usr/local/go/src/net/iprawsock.go:58
		_go_fuzz_dep_.CoverTab[6284]++
							return nil
//line /usr/local/go/src/net/iprawsock.go:59
		// _ = "end of CoverTab[6284]"
	} else {
//line /usr/local/go/src/net/iprawsock.go:60
		_go_fuzz_dep_.CoverTab[6285]++
//line /usr/local/go/src/net/iprawsock.go:60
		// _ = "end of CoverTab[6285]"
//line /usr/local/go/src/net/iprawsock.go:60
	}
//line /usr/local/go/src/net/iprawsock.go:60
	// _ = "end of CoverTab[6282]"
//line /usr/local/go/src/net/iprawsock.go:60
	_go_fuzz_dep_.CoverTab[6283]++
						return a
//line /usr/local/go/src/net/iprawsock.go:61
	// _ = "end of CoverTab[6283]"
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
	_go_fuzz_dep_.CoverTab[6286]++
						if network == "" {
//line /usr/local/go/src/net/iprawsock.go:78
		_go_fuzz_dep_.CoverTab[6291]++
							network = "ip"
//line /usr/local/go/src/net/iprawsock.go:79
		// _ = "end of CoverTab[6291]"
	} else {
//line /usr/local/go/src/net/iprawsock.go:80
		_go_fuzz_dep_.CoverTab[6292]++
//line /usr/local/go/src/net/iprawsock.go:80
		// _ = "end of CoverTab[6292]"
//line /usr/local/go/src/net/iprawsock.go:80
	}
//line /usr/local/go/src/net/iprawsock.go:80
	// _ = "end of CoverTab[6286]"
//line /usr/local/go/src/net/iprawsock.go:80
	_go_fuzz_dep_.CoverTab[6287]++
						afnet, _, err := parseNetwork(context.Background(), network, false)
						if err != nil {
//line /usr/local/go/src/net/iprawsock.go:82
		_go_fuzz_dep_.CoverTab[6293]++
							return nil, err
//line /usr/local/go/src/net/iprawsock.go:83
		// _ = "end of CoverTab[6293]"
	} else {
//line /usr/local/go/src/net/iprawsock.go:84
		_go_fuzz_dep_.CoverTab[6294]++
//line /usr/local/go/src/net/iprawsock.go:84
		// _ = "end of CoverTab[6294]"
//line /usr/local/go/src/net/iprawsock.go:84
	}
//line /usr/local/go/src/net/iprawsock.go:84
	// _ = "end of CoverTab[6287]"
//line /usr/local/go/src/net/iprawsock.go:84
	_go_fuzz_dep_.CoverTab[6288]++
						switch afnet {
	case "ip", "ip4", "ip6":
//line /usr/local/go/src/net/iprawsock.go:86
		_go_fuzz_dep_.CoverTab[6295]++
//line /usr/local/go/src/net/iprawsock.go:86
		// _ = "end of CoverTab[6295]"
	default:
//line /usr/local/go/src/net/iprawsock.go:87
		_go_fuzz_dep_.CoverTab[6296]++
							return nil, UnknownNetworkError(network)
//line /usr/local/go/src/net/iprawsock.go:88
		// _ = "end of CoverTab[6296]"
	}
//line /usr/local/go/src/net/iprawsock.go:89
	// _ = "end of CoverTab[6288]"
//line /usr/local/go/src/net/iprawsock.go:89
	_go_fuzz_dep_.CoverTab[6289]++
						addrs, err := DefaultResolver.internetAddrList(context.Background(), afnet, address)
						if err != nil {
//line /usr/local/go/src/net/iprawsock.go:91
		_go_fuzz_dep_.CoverTab[6297]++
							return nil, err
//line /usr/local/go/src/net/iprawsock.go:92
		// _ = "end of CoverTab[6297]"
	} else {
//line /usr/local/go/src/net/iprawsock.go:93
		_go_fuzz_dep_.CoverTab[6298]++
//line /usr/local/go/src/net/iprawsock.go:93
		// _ = "end of CoverTab[6298]"
//line /usr/local/go/src/net/iprawsock.go:93
	}
//line /usr/local/go/src/net/iprawsock.go:93
	// _ = "end of CoverTab[6289]"
//line /usr/local/go/src/net/iprawsock.go:93
	_go_fuzz_dep_.CoverTab[6290]++
						return addrs.forResolve(network, address).(*IPAddr), nil
//line /usr/local/go/src/net/iprawsock.go:94
	// _ = "end of CoverTab[6290]"
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
	_go_fuzz_dep_.CoverTab[6299]++
						if !c.ok() {
//line /usr/local/go/src/net/iprawsock.go:106
		_go_fuzz_dep_.CoverTab[6301]++
							return nil, syscall.EINVAL
//line /usr/local/go/src/net/iprawsock.go:107
		// _ = "end of CoverTab[6301]"
	} else {
//line /usr/local/go/src/net/iprawsock.go:108
		_go_fuzz_dep_.CoverTab[6302]++
//line /usr/local/go/src/net/iprawsock.go:108
		// _ = "end of CoverTab[6302]"
//line /usr/local/go/src/net/iprawsock.go:108
	}
//line /usr/local/go/src/net/iprawsock.go:108
	// _ = "end of CoverTab[6299]"
//line /usr/local/go/src/net/iprawsock.go:108
	_go_fuzz_dep_.CoverTab[6300]++
						return newRawConn(c.fd)
//line /usr/local/go/src/net/iprawsock.go:109
	// _ = "end of CoverTab[6300]"
}

// ReadFromIP acts like ReadFrom but returns an IPAddr.
func (c *IPConn) ReadFromIP(b []byte) (int, *IPAddr, error) {
//line /usr/local/go/src/net/iprawsock.go:113
	_go_fuzz_dep_.CoverTab[6303]++
						if !c.ok() {
//line /usr/local/go/src/net/iprawsock.go:114
		_go_fuzz_dep_.CoverTab[6306]++
							return 0, nil, syscall.EINVAL
//line /usr/local/go/src/net/iprawsock.go:115
		// _ = "end of CoverTab[6306]"
	} else {
//line /usr/local/go/src/net/iprawsock.go:116
		_go_fuzz_dep_.CoverTab[6307]++
//line /usr/local/go/src/net/iprawsock.go:116
		// _ = "end of CoverTab[6307]"
//line /usr/local/go/src/net/iprawsock.go:116
	}
//line /usr/local/go/src/net/iprawsock.go:116
	// _ = "end of CoverTab[6303]"
//line /usr/local/go/src/net/iprawsock.go:116
	_go_fuzz_dep_.CoverTab[6304]++
						n, addr, err := c.readFrom(b)
						if err != nil {
//line /usr/local/go/src/net/iprawsock.go:118
		_go_fuzz_dep_.CoverTab[6308]++
							err = &OpError{Op: "read", Net: c.fd.net, Source: c.fd.laddr, Addr: c.fd.raddr, Err: err}
//line /usr/local/go/src/net/iprawsock.go:119
		// _ = "end of CoverTab[6308]"
	} else {
//line /usr/local/go/src/net/iprawsock.go:120
		_go_fuzz_dep_.CoverTab[6309]++
//line /usr/local/go/src/net/iprawsock.go:120
		// _ = "end of CoverTab[6309]"
//line /usr/local/go/src/net/iprawsock.go:120
	}
//line /usr/local/go/src/net/iprawsock.go:120
	// _ = "end of CoverTab[6304]"
//line /usr/local/go/src/net/iprawsock.go:120
	_go_fuzz_dep_.CoverTab[6305]++
						return n, addr, err
//line /usr/local/go/src/net/iprawsock.go:121
	// _ = "end of CoverTab[6305]"
}

// ReadFrom implements the PacketConn ReadFrom method.
func (c *IPConn) ReadFrom(b []byte) (int, Addr, error) {
//line /usr/local/go/src/net/iprawsock.go:125
	_go_fuzz_dep_.CoverTab[6310]++
						if !c.ok() {
//line /usr/local/go/src/net/iprawsock.go:126
		_go_fuzz_dep_.CoverTab[6314]++
							return 0, nil, syscall.EINVAL
//line /usr/local/go/src/net/iprawsock.go:127
		// _ = "end of CoverTab[6314]"
	} else {
//line /usr/local/go/src/net/iprawsock.go:128
		_go_fuzz_dep_.CoverTab[6315]++
//line /usr/local/go/src/net/iprawsock.go:128
		// _ = "end of CoverTab[6315]"
//line /usr/local/go/src/net/iprawsock.go:128
	}
//line /usr/local/go/src/net/iprawsock.go:128
	// _ = "end of CoverTab[6310]"
//line /usr/local/go/src/net/iprawsock.go:128
	_go_fuzz_dep_.CoverTab[6311]++
						n, addr, err := c.readFrom(b)
						if err != nil {
//line /usr/local/go/src/net/iprawsock.go:130
		_go_fuzz_dep_.CoverTab[6316]++
							err = &OpError{Op: "read", Net: c.fd.net, Source: c.fd.laddr, Addr: c.fd.raddr, Err: err}
//line /usr/local/go/src/net/iprawsock.go:131
		// _ = "end of CoverTab[6316]"
	} else {
//line /usr/local/go/src/net/iprawsock.go:132
		_go_fuzz_dep_.CoverTab[6317]++
//line /usr/local/go/src/net/iprawsock.go:132
		// _ = "end of CoverTab[6317]"
//line /usr/local/go/src/net/iprawsock.go:132
	}
//line /usr/local/go/src/net/iprawsock.go:132
	// _ = "end of CoverTab[6311]"
//line /usr/local/go/src/net/iprawsock.go:132
	_go_fuzz_dep_.CoverTab[6312]++
						if addr == nil {
//line /usr/local/go/src/net/iprawsock.go:133
		_go_fuzz_dep_.CoverTab[6318]++
							return n, nil, err
//line /usr/local/go/src/net/iprawsock.go:134
		// _ = "end of CoverTab[6318]"
	} else {
//line /usr/local/go/src/net/iprawsock.go:135
		_go_fuzz_dep_.CoverTab[6319]++
//line /usr/local/go/src/net/iprawsock.go:135
		// _ = "end of CoverTab[6319]"
//line /usr/local/go/src/net/iprawsock.go:135
	}
//line /usr/local/go/src/net/iprawsock.go:135
	// _ = "end of CoverTab[6312]"
//line /usr/local/go/src/net/iprawsock.go:135
	_go_fuzz_dep_.CoverTab[6313]++
						return n, addr, err
//line /usr/local/go/src/net/iprawsock.go:136
	// _ = "end of CoverTab[6313]"
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
	_go_fuzz_dep_.CoverTab[6320]++
						if !c.ok() {
//line /usr/local/go/src/net/iprawsock.go:147
		_go_fuzz_dep_.CoverTab[6323]++
							return 0, 0, 0, nil, syscall.EINVAL
//line /usr/local/go/src/net/iprawsock.go:148
		// _ = "end of CoverTab[6323]"
	} else {
//line /usr/local/go/src/net/iprawsock.go:149
		_go_fuzz_dep_.CoverTab[6324]++
//line /usr/local/go/src/net/iprawsock.go:149
		// _ = "end of CoverTab[6324]"
//line /usr/local/go/src/net/iprawsock.go:149
	}
//line /usr/local/go/src/net/iprawsock.go:149
	// _ = "end of CoverTab[6320]"
//line /usr/local/go/src/net/iprawsock.go:149
	_go_fuzz_dep_.CoverTab[6321]++
						n, oobn, flags, addr, err = c.readMsg(b, oob)
						if err != nil {
//line /usr/local/go/src/net/iprawsock.go:151
		_go_fuzz_dep_.CoverTab[6325]++
							err = &OpError{Op: "read", Net: c.fd.net, Source: c.fd.laddr, Addr: c.fd.raddr, Err: err}
//line /usr/local/go/src/net/iprawsock.go:152
		// _ = "end of CoverTab[6325]"
	} else {
//line /usr/local/go/src/net/iprawsock.go:153
		_go_fuzz_dep_.CoverTab[6326]++
//line /usr/local/go/src/net/iprawsock.go:153
		// _ = "end of CoverTab[6326]"
//line /usr/local/go/src/net/iprawsock.go:153
	}
//line /usr/local/go/src/net/iprawsock.go:153
	// _ = "end of CoverTab[6321]"
//line /usr/local/go/src/net/iprawsock.go:153
	_go_fuzz_dep_.CoverTab[6322]++
						return
//line /usr/local/go/src/net/iprawsock.go:154
	// _ = "end of CoverTab[6322]"
}

// WriteToIP acts like WriteTo but takes an IPAddr.
func (c *IPConn) WriteToIP(b []byte, addr *IPAddr) (int, error) {
//line /usr/local/go/src/net/iprawsock.go:158
	_go_fuzz_dep_.CoverTab[6327]++
						if !c.ok() {
//line /usr/local/go/src/net/iprawsock.go:159
		_go_fuzz_dep_.CoverTab[6330]++
							return 0, syscall.EINVAL
//line /usr/local/go/src/net/iprawsock.go:160
		// _ = "end of CoverTab[6330]"
	} else {
//line /usr/local/go/src/net/iprawsock.go:161
		_go_fuzz_dep_.CoverTab[6331]++
//line /usr/local/go/src/net/iprawsock.go:161
		// _ = "end of CoverTab[6331]"
//line /usr/local/go/src/net/iprawsock.go:161
	}
//line /usr/local/go/src/net/iprawsock.go:161
	// _ = "end of CoverTab[6327]"
//line /usr/local/go/src/net/iprawsock.go:161
	_go_fuzz_dep_.CoverTab[6328]++
						n, err := c.writeTo(b, addr)
						if err != nil {
//line /usr/local/go/src/net/iprawsock.go:163
		_go_fuzz_dep_.CoverTab[6332]++
							err = &OpError{Op: "write", Net: c.fd.net, Source: c.fd.laddr, Addr: addr.opAddr(), Err: err}
//line /usr/local/go/src/net/iprawsock.go:164
		// _ = "end of CoverTab[6332]"
	} else {
//line /usr/local/go/src/net/iprawsock.go:165
		_go_fuzz_dep_.CoverTab[6333]++
//line /usr/local/go/src/net/iprawsock.go:165
		// _ = "end of CoverTab[6333]"
//line /usr/local/go/src/net/iprawsock.go:165
	}
//line /usr/local/go/src/net/iprawsock.go:165
	// _ = "end of CoverTab[6328]"
//line /usr/local/go/src/net/iprawsock.go:165
	_go_fuzz_dep_.CoverTab[6329]++
						return n, err
//line /usr/local/go/src/net/iprawsock.go:166
	// _ = "end of CoverTab[6329]"
}

// WriteTo implements the PacketConn WriteTo method.
func (c *IPConn) WriteTo(b []byte, addr Addr) (int, error) {
//line /usr/local/go/src/net/iprawsock.go:170
	_go_fuzz_dep_.CoverTab[6334]++
						if !c.ok() {
//line /usr/local/go/src/net/iprawsock.go:171
		_go_fuzz_dep_.CoverTab[6338]++
							return 0, syscall.EINVAL
//line /usr/local/go/src/net/iprawsock.go:172
		// _ = "end of CoverTab[6338]"
	} else {
//line /usr/local/go/src/net/iprawsock.go:173
		_go_fuzz_dep_.CoverTab[6339]++
//line /usr/local/go/src/net/iprawsock.go:173
		// _ = "end of CoverTab[6339]"
//line /usr/local/go/src/net/iprawsock.go:173
	}
//line /usr/local/go/src/net/iprawsock.go:173
	// _ = "end of CoverTab[6334]"
//line /usr/local/go/src/net/iprawsock.go:173
	_go_fuzz_dep_.CoverTab[6335]++
						a, ok := addr.(*IPAddr)
						if !ok {
//line /usr/local/go/src/net/iprawsock.go:175
		_go_fuzz_dep_.CoverTab[6340]++
							return 0, &OpError{Op: "write", Net: c.fd.net, Source: c.fd.laddr, Addr: addr, Err: syscall.EINVAL}
//line /usr/local/go/src/net/iprawsock.go:176
		// _ = "end of CoverTab[6340]"
	} else {
//line /usr/local/go/src/net/iprawsock.go:177
		_go_fuzz_dep_.CoverTab[6341]++
//line /usr/local/go/src/net/iprawsock.go:177
		// _ = "end of CoverTab[6341]"
//line /usr/local/go/src/net/iprawsock.go:177
	}
//line /usr/local/go/src/net/iprawsock.go:177
	// _ = "end of CoverTab[6335]"
//line /usr/local/go/src/net/iprawsock.go:177
	_go_fuzz_dep_.CoverTab[6336]++
						n, err := c.writeTo(b, a)
						if err != nil {
//line /usr/local/go/src/net/iprawsock.go:179
		_go_fuzz_dep_.CoverTab[6342]++
							err = &OpError{Op: "write", Net: c.fd.net, Source: c.fd.laddr, Addr: a.opAddr(), Err: err}
//line /usr/local/go/src/net/iprawsock.go:180
		// _ = "end of CoverTab[6342]"
	} else {
//line /usr/local/go/src/net/iprawsock.go:181
		_go_fuzz_dep_.CoverTab[6343]++
//line /usr/local/go/src/net/iprawsock.go:181
		// _ = "end of CoverTab[6343]"
//line /usr/local/go/src/net/iprawsock.go:181
	}
//line /usr/local/go/src/net/iprawsock.go:181
	// _ = "end of CoverTab[6336]"
//line /usr/local/go/src/net/iprawsock.go:181
	_go_fuzz_dep_.CoverTab[6337]++
						return n, err
//line /usr/local/go/src/net/iprawsock.go:182
	// _ = "end of CoverTab[6337]"
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
	_go_fuzz_dep_.CoverTab[6344]++
						if !c.ok() {
//line /usr/local/go/src/net/iprawsock.go:192
		_go_fuzz_dep_.CoverTab[6347]++
							return 0, 0, syscall.EINVAL
//line /usr/local/go/src/net/iprawsock.go:193
		// _ = "end of CoverTab[6347]"
	} else {
//line /usr/local/go/src/net/iprawsock.go:194
		_go_fuzz_dep_.CoverTab[6348]++
//line /usr/local/go/src/net/iprawsock.go:194
		// _ = "end of CoverTab[6348]"
//line /usr/local/go/src/net/iprawsock.go:194
	}
//line /usr/local/go/src/net/iprawsock.go:194
	// _ = "end of CoverTab[6344]"
//line /usr/local/go/src/net/iprawsock.go:194
	_go_fuzz_dep_.CoverTab[6345]++
						n, oobn, err = c.writeMsg(b, oob, addr)
						if err != nil {
//line /usr/local/go/src/net/iprawsock.go:196
		_go_fuzz_dep_.CoverTab[6349]++
							err = &OpError{Op: "write", Net: c.fd.net, Source: c.fd.laddr, Addr: addr.opAddr(), Err: err}
//line /usr/local/go/src/net/iprawsock.go:197
		// _ = "end of CoverTab[6349]"
	} else {
//line /usr/local/go/src/net/iprawsock.go:198
		_go_fuzz_dep_.CoverTab[6350]++
//line /usr/local/go/src/net/iprawsock.go:198
		// _ = "end of CoverTab[6350]"
//line /usr/local/go/src/net/iprawsock.go:198
	}
//line /usr/local/go/src/net/iprawsock.go:198
	// _ = "end of CoverTab[6345]"
//line /usr/local/go/src/net/iprawsock.go:198
	_go_fuzz_dep_.CoverTab[6346]++
						return
//line /usr/local/go/src/net/iprawsock.go:199
	// _ = "end of CoverTab[6346]"
}

func newIPConn(fd *netFD) *IPConn {
//line /usr/local/go/src/net/iprawsock.go:202
	_go_fuzz_dep_.CoverTab[6351]++
//line /usr/local/go/src/net/iprawsock.go:202
	return &IPConn{conn{fd}}
//line /usr/local/go/src/net/iprawsock.go:202
	// _ = "end of CoverTab[6351]"
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
	_go_fuzz_dep_.CoverTab[6352]++
						if raddr == nil {
//line /usr/local/go/src/net/iprawsock.go:212
		_go_fuzz_dep_.CoverTab[6355]++
							return nil, &OpError{Op: "dial", Net: network, Source: laddr.opAddr(), Addr: nil, Err: errMissingAddress}
//line /usr/local/go/src/net/iprawsock.go:213
		// _ = "end of CoverTab[6355]"
	} else {
//line /usr/local/go/src/net/iprawsock.go:214
		_go_fuzz_dep_.CoverTab[6356]++
//line /usr/local/go/src/net/iprawsock.go:214
		// _ = "end of CoverTab[6356]"
//line /usr/local/go/src/net/iprawsock.go:214
	}
//line /usr/local/go/src/net/iprawsock.go:214
	// _ = "end of CoverTab[6352]"
//line /usr/local/go/src/net/iprawsock.go:214
	_go_fuzz_dep_.CoverTab[6353]++
						sd := &sysDialer{network: network, address: raddr.String()}
						c, err := sd.dialIP(context.Background(), laddr, raddr)
						if err != nil {
//line /usr/local/go/src/net/iprawsock.go:217
		_go_fuzz_dep_.CoverTab[6357]++
							return nil, &OpError{Op: "dial", Net: network, Source: laddr.opAddr(), Addr: raddr.opAddr(), Err: err}
//line /usr/local/go/src/net/iprawsock.go:218
		// _ = "end of CoverTab[6357]"
	} else {
//line /usr/local/go/src/net/iprawsock.go:219
		_go_fuzz_dep_.CoverTab[6358]++
//line /usr/local/go/src/net/iprawsock.go:219
		// _ = "end of CoverTab[6358]"
//line /usr/local/go/src/net/iprawsock.go:219
	}
//line /usr/local/go/src/net/iprawsock.go:219
	// _ = "end of CoverTab[6353]"
//line /usr/local/go/src/net/iprawsock.go:219
	_go_fuzz_dep_.CoverTab[6354]++
						return c, nil
//line /usr/local/go/src/net/iprawsock.go:220
	// _ = "end of CoverTab[6354]"
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
	_go_fuzz_dep_.CoverTab[6359]++
						if laddr == nil {
//line /usr/local/go/src/net/iprawsock.go:231
		_go_fuzz_dep_.CoverTab[6362]++
							laddr = &IPAddr{}
//line /usr/local/go/src/net/iprawsock.go:232
		// _ = "end of CoverTab[6362]"
	} else {
//line /usr/local/go/src/net/iprawsock.go:233
		_go_fuzz_dep_.CoverTab[6363]++
//line /usr/local/go/src/net/iprawsock.go:233
		// _ = "end of CoverTab[6363]"
//line /usr/local/go/src/net/iprawsock.go:233
	}
//line /usr/local/go/src/net/iprawsock.go:233
	// _ = "end of CoverTab[6359]"
//line /usr/local/go/src/net/iprawsock.go:233
	_go_fuzz_dep_.CoverTab[6360]++
						sl := &sysListener{network: network, address: laddr.String()}
						c, err := sl.listenIP(context.Background(), laddr)
						if err != nil {
//line /usr/local/go/src/net/iprawsock.go:236
		_go_fuzz_dep_.CoverTab[6364]++
							return nil, &OpError{Op: "listen", Net: network, Source: nil, Addr: laddr.opAddr(), Err: err}
//line /usr/local/go/src/net/iprawsock.go:237
		// _ = "end of CoverTab[6364]"
	} else {
//line /usr/local/go/src/net/iprawsock.go:238
		_go_fuzz_dep_.CoverTab[6365]++
//line /usr/local/go/src/net/iprawsock.go:238
		// _ = "end of CoverTab[6365]"
//line /usr/local/go/src/net/iprawsock.go:238
	}
//line /usr/local/go/src/net/iprawsock.go:238
	// _ = "end of CoverTab[6360]"
//line /usr/local/go/src/net/iprawsock.go:238
	_go_fuzz_dep_.CoverTab[6361]++
						return c, nil
//line /usr/local/go/src/net/iprawsock.go:239
	// _ = "end of CoverTab[6361]"
}

//line /usr/local/go/src/net/iprawsock.go:240
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/iprawsock.go:240
var _ = _go_fuzz_dep_.CoverTab
