// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /snap/go/10455/src/net/iprawsock.go:5
package net

//line /snap/go/10455/src/net/iprawsock.go:5
import (
//line /snap/go/10455/src/net/iprawsock.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /snap/go/10455/src/net/iprawsock.go:5
)
//line /snap/go/10455/src/net/iprawsock.go:5
import (
//line /snap/go/10455/src/net/iprawsock.go:5
	_atomic_ "sync/atomic"
//line /snap/go/10455/src/net/iprawsock.go:5
)

import (
	"context"
	"syscall"
)

//line /snap/go/10455/src/net/iprawsock.go:30
// IPAddr represents the address of an IP end point.
type IPAddr struct {
	IP	IP
	Zone	string	// IPv6 scoped addressing zone
}

// Network returns the address's network name, "ip".
func (a *IPAddr) Network() string {
//line /snap/go/10455/src/net/iprawsock.go:37
	_go_fuzz_dep_.CoverTab[6535]++
//line /snap/go/10455/src/net/iprawsock.go:37
	return "ip"
//line /snap/go/10455/src/net/iprawsock.go:37
	// _ = "end of CoverTab[6535]"
//line /snap/go/10455/src/net/iprawsock.go:37
}

func (a *IPAddr) String() string {
//line /snap/go/10455/src/net/iprawsock.go:39
	_go_fuzz_dep_.CoverTab[6536]++
						if a == nil {
//line /snap/go/10455/src/net/iprawsock.go:40
		_go_fuzz_dep_.CoverTab[528697]++
//line /snap/go/10455/src/net/iprawsock.go:40
		_go_fuzz_dep_.CoverTab[6539]++
							return "<nil>"
//line /snap/go/10455/src/net/iprawsock.go:41
		// _ = "end of CoverTab[6539]"
	} else {
//line /snap/go/10455/src/net/iprawsock.go:42
		_go_fuzz_dep_.CoverTab[528698]++
//line /snap/go/10455/src/net/iprawsock.go:42
		_go_fuzz_dep_.CoverTab[6540]++
//line /snap/go/10455/src/net/iprawsock.go:42
		// _ = "end of CoverTab[6540]"
//line /snap/go/10455/src/net/iprawsock.go:42
	}
//line /snap/go/10455/src/net/iprawsock.go:42
	// _ = "end of CoverTab[6536]"
//line /snap/go/10455/src/net/iprawsock.go:42
	_go_fuzz_dep_.CoverTab[6537]++
						ip := ipEmptyString(a.IP)
						if a.Zone != "" {
//line /snap/go/10455/src/net/iprawsock.go:44
		_go_fuzz_dep_.CoverTab[528699]++
//line /snap/go/10455/src/net/iprawsock.go:44
		_go_fuzz_dep_.CoverTab[6541]++
							return ip + "%" + a.Zone
//line /snap/go/10455/src/net/iprawsock.go:45
		// _ = "end of CoverTab[6541]"
	} else {
//line /snap/go/10455/src/net/iprawsock.go:46
		_go_fuzz_dep_.CoverTab[528700]++
//line /snap/go/10455/src/net/iprawsock.go:46
		_go_fuzz_dep_.CoverTab[6542]++
//line /snap/go/10455/src/net/iprawsock.go:46
		// _ = "end of CoverTab[6542]"
//line /snap/go/10455/src/net/iprawsock.go:46
	}
//line /snap/go/10455/src/net/iprawsock.go:46
	// _ = "end of CoverTab[6537]"
//line /snap/go/10455/src/net/iprawsock.go:46
	_go_fuzz_dep_.CoverTab[6538]++
						return ip
//line /snap/go/10455/src/net/iprawsock.go:47
	// _ = "end of CoverTab[6538]"
}

func (a *IPAddr) isWildcard() bool {
//line /snap/go/10455/src/net/iprawsock.go:50
	_go_fuzz_dep_.CoverTab[6543]++
						if a == nil || func() bool {
//line /snap/go/10455/src/net/iprawsock.go:51
		_go_fuzz_dep_.CoverTab[6545]++
//line /snap/go/10455/src/net/iprawsock.go:51
		return a.IP == nil
//line /snap/go/10455/src/net/iprawsock.go:51
		// _ = "end of CoverTab[6545]"
//line /snap/go/10455/src/net/iprawsock.go:51
	}() {
//line /snap/go/10455/src/net/iprawsock.go:51
		_go_fuzz_dep_.CoverTab[528701]++
//line /snap/go/10455/src/net/iprawsock.go:51
		_go_fuzz_dep_.CoverTab[6546]++
							return true
//line /snap/go/10455/src/net/iprawsock.go:52
		// _ = "end of CoverTab[6546]"
	} else {
//line /snap/go/10455/src/net/iprawsock.go:53
		_go_fuzz_dep_.CoverTab[528702]++
//line /snap/go/10455/src/net/iprawsock.go:53
		_go_fuzz_dep_.CoverTab[6547]++
//line /snap/go/10455/src/net/iprawsock.go:53
		// _ = "end of CoverTab[6547]"
//line /snap/go/10455/src/net/iprawsock.go:53
	}
//line /snap/go/10455/src/net/iprawsock.go:53
	// _ = "end of CoverTab[6543]"
//line /snap/go/10455/src/net/iprawsock.go:53
	_go_fuzz_dep_.CoverTab[6544]++
						return a.IP.IsUnspecified()
//line /snap/go/10455/src/net/iprawsock.go:54
	// _ = "end of CoverTab[6544]"
}

func (a *IPAddr) opAddr() Addr {
//line /snap/go/10455/src/net/iprawsock.go:57
	_go_fuzz_dep_.CoverTab[6548]++
						if a == nil {
//line /snap/go/10455/src/net/iprawsock.go:58
		_go_fuzz_dep_.CoverTab[528703]++
//line /snap/go/10455/src/net/iprawsock.go:58
		_go_fuzz_dep_.CoverTab[6550]++
							return nil
//line /snap/go/10455/src/net/iprawsock.go:59
		// _ = "end of CoverTab[6550]"
	} else {
//line /snap/go/10455/src/net/iprawsock.go:60
		_go_fuzz_dep_.CoverTab[528704]++
//line /snap/go/10455/src/net/iprawsock.go:60
		_go_fuzz_dep_.CoverTab[6551]++
//line /snap/go/10455/src/net/iprawsock.go:60
		// _ = "end of CoverTab[6551]"
//line /snap/go/10455/src/net/iprawsock.go:60
	}
//line /snap/go/10455/src/net/iprawsock.go:60
	// _ = "end of CoverTab[6548]"
//line /snap/go/10455/src/net/iprawsock.go:60
	_go_fuzz_dep_.CoverTab[6549]++
						return a
//line /snap/go/10455/src/net/iprawsock.go:61
	// _ = "end of CoverTab[6549]"
}

// ResolveIPAddr returns an address of IP end point.
//line /snap/go/10455/src/net/iprawsock.go:64
//
//line /snap/go/10455/src/net/iprawsock.go:64
// The network must be an IP network name.
//line /snap/go/10455/src/net/iprawsock.go:64
//
//line /snap/go/10455/src/net/iprawsock.go:64
// If the host in the address parameter is not a literal IP address,
//line /snap/go/10455/src/net/iprawsock.go:64
// ResolveIPAddr resolves the address to an address of IP end point.
//line /snap/go/10455/src/net/iprawsock.go:64
// Otherwise, it parses the address as a literal IP address.
//line /snap/go/10455/src/net/iprawsock.go:64
// The address parameter can use a host name, but this is not
//line /snap/go/10455/src/net/iprawsock.go:64
// recommended, because it will return at most one of the host name's
//line /snap/go/10455/src/net/iprawsock.go:64
// IP addresses.
//line /snap/go/10455/src/net/iprawsock.go:64
//
//line /snap/go/10455/src/net/iprawsock.go:64
// See func Dial for a description of the network and address
//line /snap/go/10455/src/net/iprawsock.go:64
// parameters.
//line /snap/go/10455/src/net/iprawsock.go:77
func ResolveIPAddr(network, address string) (*IPAddr, error) {
//line /snap/go/10455/src/net/iprawsock.go:77
	_go_fuzz_dep_.CoverTab[6552]++
						if network == "" {
//line /snap/go/10455/src/net/iprawsock.go:78
		_go_fuzz_dep_.CoverTab[528705]++
//line /snap/go/10455/src/net/iprawsock.go:78
		_go_fuzz_dep_.CoverTab[6557]++
							network = "ip"
//line /snap/go/10455/src/net/iprawsock.go:79
		// _ = "end of CoverTab[6557]"
	} else {
//line /snap/go/10455/src/net/iprawsock.go:80
		_go_fuzz_dep_.CoverTab[528706]++
//line /snap/go/10455/src/net/iprawsock.go:80
		_go_fuzz_dep_.CoverTab[6558]++
//line /snap/go/10455/src/net/iprawsock.go:80
		// _ = "end of CoverTab[6558]"
//line /snap/go/10455/src/net/iprawsock.go:80
	}
//line /snap/go/10455/src/net/iprawsock.go:80
	// _ = "end of CoverTab[6552]"
//line /snap/go/10455/src/net/iprawsock.go:80
	_go_fuzz_dep_.CoverTab[6553]++
						afnet, _, err := parseNetwork(context.Background(), network, false)
						if err != nil {
//line /snap/go/10455/src/net/iprawsock.go:82
		_go_fuzz_dep_.CoverTab[528707]++
//line /snap/go/10455/src/net/iprawsock.go:82
		_go_fuzz_dep_.CoverTab[6559]++
							return nil, err
//line /snap/go/10455/src/net/iprawsock.go:83
		// _ = "end of CoverTab[6559]"
	} else {
//line /snap/go/10455/src/net/iprawsock.go:84
		_go_fuzz_dep_.CoverTab[528708]++
//line /snap/go/10455/src/net/iprawsock.go:84
		_go_fuzz_dep_.CoverTab[6560]++
//line /snap/go/10455/src/net/iprawsock.go:84
		// _ = "end of CoverTab[6560]"
//line /snap/go/10455/src/net/iprawsock.go:84
	}
//line /snap/go/10455/src/net/iprawsock.go:84
	// _ = "end of CoverTab[6553]"
//line /snap/go/10455/src/net/iprawsock.go:84
	_go_fuzz_dep_.CoverTab[6554]++
						switch afnet {
	case "ip", "ip4", "ip6":
//line /snap/go/10455/src/net/iprawsock.go:86
		_go_fuzz_dep_.CoverTab[528709]++
//line /snap/go/10455/src/net/iprawsock.go:86
		_go_fuzz_dep_.CoverTab[6561]++
//line /snap/go/10455/src/net/iprawsock.go:86
		// _ = "end of CoverTab[6561]"
	default:
//line /snap/go/10455/src/net/iprawsock.go:87
		_go_fuzz_dep_.CoverTab[528710]++
//line /snap/go/10455/src/net/iprawsock.go:87
		_go_fuzz_dep_.CoverTab[6562]++
							return nil, UnknownNetworkError(network)
//line /snap/go/10455/src/net/iprawsock.go:88
		// _ = "end of CoverTab[6562]"
	}
//line /snap/go/10455/src/net/iprawsock.go:89
	// _ = "end of CoverTab[6554]"
//line /snap/go/10455/src/net/iprawsock.go:89
	_go_fuzz_dep_.CoverTab[6555]++
						addrs, err := DefaultResolver.internetAddrList(context.Background(), afnet, address)
						if err != nil {
//line /snap/go/10455/src/net/iprawsock.go:91
		_go_fuzz_dep_.CoverTab[528711]++
//line /snap/go/10455/src/net/iprawsock.go:91
		_go_fuzz_dep_.CoverTab[6563]++
							return nil, err
//line /snap/go/10455/src/net/iprawsock.go:92
		// _ = "end of CoverTab[6563]"
	} else {
//line /snap/go/10455/src/net/iprawsock.go:93
		_go_fuzz_dep_.CoverTab[528712]++
//line /snap/go/10455/src/net/iprawsock.go:93
		_go_fuzz_dep_.CoverTab[6564]++
//line /snap/go/10455/src/net/iprawsock.go:93
		// _ = "end of CoverTab[6564]"
//line /snap/go/10455/src/net/iprawsock.go:93
	}
//line /snap/go/10455/src/net/iprawsock.go:93
	// _ = "end of CoverTab[6555]"
//line /snap/go/10455/src/net/iprawsock.go:93
	_go_fuzz_dep_.CoverTab[6556]++
						return addrs.forResolve(network, address).(*IPAddr), nil
//line /snap/go/10455/src/net/iprawsock.go:94
	// _ = "end of CoverTab[6556]"
}

// IPConn is the implementation of the Conn and PacketConn interfaces
//line /snap/go/10455/src/net/iprawsock.go:97
// for IP network connections.
//line /snap/go/10455/src/net/iprawsock.go:99
type IPConn struct {
	conn
}

// SyscallConn returns a raw network connection.
//line /snap/go/10455/src/net/iprawsock.go:103
// This implements the syscall.Conn interface.
//line /snap/go/10455/src/net/iprawsock.go:105
func (c *IPConn) SyscallConn() (syscall.RawConn, error) {
//line /snap/go/10455/src/net/iprawsock.go:105
	_go_fuzz_dep_.CoverTab[6565]++
						if !c.ok() {
//line /snap/go/10455/src/net/iprawsock.go:106
		_go_fuzz_dep_.CoverTab[528713]++
//line /snap/go/10455/src/net/iprawsock.go:106
		_go_fuzz_dep_.CoverTab[6567]++
							return nil, syscall.EINVAL
//line /snap/go/10455/src/net/iprawsock.go:107
		// _ = "end of CoverTab[6567]"
	} else {
//line /snap/go/10455/src/net/iprawsock.go:108
		_go_fuzz_dep_.CoverTab[528714]++
//line /snap/go/10455/src/net/iprawsock.go:108
		_go_fuzz_dep_.CoverTab[6568]++
//line /snap/go/10455/src/net/iprawsock.go:108
		// _ = "end of CoverTab[6568]"
//line /snap/go/10455/src/net/iprawsock.go:108
	}
//line /snap/go/10455/src/net/iprawsock.go:108
	// _ = "end of CoverTab[6565]"
//line /snap/go/10455/src/net/iprawsock.go:108
	_go_fuzz_dep_.CoverTab[6566]++
						return newRawConn(c.fd)
//line /snap/go/10455/src/net/iprawsock.go:109
	// _ = "end of CoverTab[6566]"
}

// ReadFromIP acts like ReadFrom but returns an IPAddr.
func (c *IPConn) ReadFromIP(b []byte) (int, *IPAddr, error) {
//line /snap/go/10455/src/net/iprawsock.go:113
	_go_fuzz_dep_.CoverTab[6569]++
						if !c.ok() {
//line /snap/go/10455/src/net/iprawsock.go:114
		_go_fuzz_dep_.CoverTab[528715]++
//line /snap/go/10455/src/net/iprawsock.go:114
		_go_fuzz_dep_.CoverTab[6572]++
							return 0, nil, syscall.EINVAL
//line /snap/go/10455/src/net/iprawsock.go:115
		// _ = "end of CoverTab[6572]"
	} else {
//line /snap/go/10455/src/net/iprawsock.go:116
		_go_fuzz_dep_.CoverTab[528716]++
//line /snap/go/10455/src/net/iprawsock.go:116
		_go_fuzz_dep_.CoverTab[6573]++
//line /snap/go/10455/src/net/iprawsock.go:116
		// _ = "end of CoverTab[6573]"
//line /snap/go/10455/src/net/iprawsock.go:116
	}
//line /snap/go/10455/src/net/iprawsock.go:116
	// _ = "end of CoverTab[6569]"
//line /snap/go/10455/src/net/iprawsock.go:116
	_go_fuzz_dep_.CoverTab[6570]++
						n, addr, err := c.readFrom(b)
						if err != nil {
//line /snap/go/10455/src/net/iprawsock.go:118
		_go_fuzz_dep_.CoverTab[528717]++
//line /snap/go/10455/src/net/iprawsock.go:118
		_go_fuzz_dep_.CoverTab[6574]++
							err = &OpError{Op: "read", Net: c.fd.net, Source: c.fd.laddr, Addr: c.fd.raddr, Err: err}
//line /snap/go/10455/src/net/iprawsock.go:119
		// _ = "end of CoverTab[6574]"
	} else {
//line /snap/go/10455/src/net/iprawsock.go:120
		_go_fuzz_dep_.CoverTab[528718]++
//line /snap/go/10455/src/net/iprawsock.go:120
		_go_fuzz_dep_.CoverTab[6575]++
//line /snap/go/10455/src/net/iprawsock.go:120
		// _ = "end of CoverTab[6575]"
//line /snap/go/10455/src/net/iprawsock.go:120
	}
//line /snap/go/10455/src/net/iprawsock.go:120
	// _ = "end of CoverTab[6570]"
//line /snap/go/10455/src/net/iprawsock.go:120
	_go_fuzz_dep_.CoverTab[6571]++
						return n, addr, err
//line /snap/go/10455/src/net/iprawsock.go:121
	// _ = "end of CoverTab[6571]"
}

// ReadFrom implements the PacketConn ReadFrom method.
func (c *IPConn) ReadFrom(b []byte) (int, Addr, error) {
//line /snap/go/10455/src/net/iprawsock.go:125
	_go_fuzz_dep_.CoverTab[6576]++
						if !c.ok() {
//line /snap/go/10455/src/net/iprawsock.go:126
		_go_fuzz_dep_.CoverTab[528719]++
//line /snap/go/10455/src/net/iprawsock.go:126
		_go_fuzz_dep_.CoverTab[6580]++
							return 0, nil, syscall.EINVAL
//line /snap/go/10455/src/net/iprawsock.go:127
		// _ = "end of CoverTab[6580]"
	} else {
//line /snap/go/10455/src/net/iprawsock.go:128
		_go_fuzz_dep_.CoverTab[528720]++
//line /snap/go/10455/src/net/iprawsock.go:128
		_go_fuzz_dep_.CoverTab[6581]++
//line /snap/go/10455/src/net/iprawsock.go:128
		// _ = "end of CoverTab[6581]"
//line /snap/go/10455/src/net/iprawsock.go:128
	}
//line /snap/go/10455/src/net/iprawsock.go:128
	// _ = "end of CoverTab[6576]"
//line /snap/go/10455/src/net/iprawsock.go:128
	_go_fuzz_dep_.CoverTab[6577]++
						n, addr, err := c.readFrom(b)
						if err != nil {
//line /snap/go/10455/src/net/iprawsock.go:130
		_go_fuzz_dep_.CoverTab[528721]++
//line /snap/go/10455/src/net/iprawsock.go:130
		_go_fuzz_dep_.CoverTab[6582]++
							err = &OpError{Op: "read", Net: c.fd.net, Source: c.fd.laddr, Addr: c.fd.raddr, Err: err}
//line /snap/go/10455/src/net/iprawsock.go:131
		// _ = "end of CoverTab[6582]"
	} else {
//line /snap/go/10455/src/net/iprawsock.go:132
		_go_fuzz_dep_.CoverTab[528722]++
//line /snap/go/10455/src/net/iprawsock.go:132
		_go_fuzz_dep_.CoverTab[6583]++
//line /snap/go/10455/src/net/iprawsock.go:132
		// _ = "end of CoverTab[6583]"
//line /snap/go/10455/src/net/iprawsock.go:132
	}
//line /snap/go/10455/src/net/iprawsock.go:132
	// _ = "end of CoverTab[6577]"
//line /snap/go/10455/src/net/iprawsock.go:132
	_go_fuzz_dep_.CoverTab[6578]++
						if addr == nil {
//line /snap/go/10455/src/net/iprawsock.go:133
		_go_fuzz_dep_.CoverTab[528723]++
//line /snap/go/10455/src/net/iprawsock.go:133
		_go_fuzz_dep_.CoverTab[6584]++
							return n, nil, err
//line /snap/go/10455/src/net/iprawsock.go:134
		// _ = "end of CoverTab[6584]"
	} else {
//line /snap/go/10455/src/net/iprawsock.go:135
		_go_fuzz_dep_.CoverTab[528724]++
//line /snap/go/10455/src/net/iprawsock.go:135
		_go_fuzz_dep_.CoverTab[6585]++
//line /snap/go/10455/src/net/iprawsock.go:135
		// _ = "end of CoverTab[6585]"
//line /snap/go/10455/src/net/iprawsock.go:135
	}
//line /snap/go/10455/src/net/iprawsock.go:135
	// _ = "end of CoverTab[6578]"
//line /snap/go/10455/src/net/iprawsock.go:135
	_go_fuzz_dep_.CoverTab[6579]++
						return n, addr, err
//line /snap/go/10455/src/net/iprawsock.go:136
	// _ = "end of CoverTab[6579]"
}

// ReadMsgIP reads a message from c, copying the payload into b and
//line /snap/go/10455/src/net/iprawsock.go:139
// the associated out-of-band data into oob. It returns the number of
//line /snap/go/10455/src/net/iprawsock.go:139
// bytes copied into b, the number of bytes copied into oob, the flags
//line /snap/go/10455/src/net/iprawsock.go:139
// that were set on the message and the source address of the message.
//line /snap/go/10455/src/net/iprawsock.go:139
//
//line /snap/go/10455/src/net/iprawsock.go:139
// The packages golang.org/x/net/ipv4 and golang.org/x/net/ipv6 can be
//line /snap/go/10455/src/net/iprawsock.go:139
// used to manipulate IP-level socket options in oob.
//line /snap/go/10455/src/net/iprawsock.go:146
func (c *IPConn) ReadMsgIP(b, oob []byte) (n, oobn, flags int, addr *IPAddr, err error) {
//line /snap/go/10455/src/net/iprawsock.go:146
	_go_fuzz_dep_.CoverTab[6586]++
						if !c.ok() {
//line /snap/go/10455/src/net/iprawsock.go:147
		_go_fuzz_dep_.CoverTab[528725]++
//line /snap/go/10455/src/net/iprawsock.go:147
		_go_fuzz_dep_.CoverTab[6589]++
							return 0, 0, 0, nil, syscall.EINVAL
//line /snap/go/10455/src/net/iprawsock.go:148
		// _ = "end of CoverTab[6589]"
	} else {
//line /snap/go/10455/src/net/iprawsock.go:149
		_go_fuzz_dep_.CoverTab[528726]++
//line /snap/go/10455/src/net/iprawsock.go:149
		_go_fuzz_dep_.CoverTab[6590]++
//line /snap/go/10455/src/net/iprawsock.go:149
		// _ = "end of CoverTab[6590]"
//line /snap/go/10455/src/net/iprawsock.go:149
	}
//line /snap/go/10455/src/net/iprawsock.go:149
	// _ = "end of CoverTab[6586]"
//line /snap/go/10455/src/net/iprawsock.go:149
	_go_fuzz_dep_.CoverTab[6587]++
						n, oobn, flags, addr, err = c.readMsg(b, oob)
						if err != nil {
//line /snap/go/10455/src/net/iprawsock.go:151
		_go_fuzz_dep_.CoverTab[528727]++
//line /snap/go/10455/src/net/iprawsock.go:151
		_go_fuzz_dep_.CoverTab[6591]++
							err = &OpError{Op: "read", Net: c.fd.net, Source: c.fd.laddr, Addr: c.fd.raddr, Err: err}
//line /snap/go/10455/src/net/iprawsock.go:152
		// _ = "end of CoverTab[6591]"
	} else {
//line /snap/go/10455/src/net/iprawsock.go:153
		_go_fuzz_dep_.CoverTab[528728]++
//line /snap/go/10455/src/net/iprawsock.go:153
		_go_fuzz_dep_.CoverTab[6592]++
//line /snap/go/10455/src/net/iprawsock.go:153
		// _ = "end of CoverTab[6592]"
//line /snap/go/10455/src/net/iprawsock.go:153
	}
//line /snap/go/10455/src/net/iprawsock.go:153
	// _ = "end of CoverTab[6587]"
//line /snap/go/10455/src/net/iprawsock.go:153
	_go_fuzz_dep_.CoverTab[6588]++
						return
//line /snap/go/10455/src/net/iprawsock.go:154
	// _ = "end of CoverTab[6588]"
}

// WriteToIP acts like WriteTo but takes an IPAddr.
func (c *IPConn) WriteToIP(b []byte, addr *IPAddr) (int, error) {
//line /snap/go/10455/src/net/iprawsock.go:158
	_go_fuzz_dep_.CoverTab[6593]++
						if !c.ok() {
//line /snap/go/10455/src/net/iprawsock.go:159
		_go_fuzz_dep_.CoverTab[528729]++
//line /snap/go/10455/src/net/iprawsock.go:159
		_go_fuzz_dep_.CoverTab[6596]++
							return 0, syscall.EINVAL
//line /snap/go/10455/src/net/iprawsock.go:160
		// _ = "end of CoverTab[6596]"
	} else {
//line /snap/go/10455/src/net/iprawsock.go:161
		_go_fuzz_dep_.CoverTab[528730]++
//line /snap/go/10455/src/net/iprawsock.go:161
		_go_fuzz_dep_.CoverTab[6597]++
//line /snap/go/10455/src/net/iprawsock.go:161
		// _ = "end of CoverTab[6597]"
//line /snap/go/10455/src/net/iprawsock.go:161
	}
//line /snap/go/10455/src/net/iprawsock.go:161
	// _ = "end of CoverTab[6593]"
//line /snap/go/10455/src/net/iprawsock.go:161
	_go_fuzz_dep_.CoverTab[6594]++
						n, err := c.writeTo(b, addr)
						if err != nil {
//line /snap/go/10455/src/net/iprawsock.go:163
		_go_fuzz_dep_.CoverTab[528731]++
//line /snap/go/10455/src/net/iprawsock.go:163
		_go_fuzz_dep_.CoverTab[6598]++
							err = &OpError{Op: "write", Net: c.fd.net, Source: c.fd.laddr, Addr: addr.opAddr(), Err: err}
//line /snap/go/10455/src/net/iprawsock.go:164
		// _ = "end of CoverTab[6598]"
	} else {
//line /snap/go/10455/src/net/iprawsock.go:165
		_go_fuzz_dep_.CoverTab[528732]++
//line /snap/go/10455/src/net/iprawsock.go:165
		_go_fuzz_dep_.CoverTab[6599]++
//line /snap/go/10455/src/net/iprawsock.go:165
		// _ = "end of CoverTab[6599]"
//line /snap/go/10455/src/net/iprawsock.go:165
	}
//line /snap/go/10455/src/net/iprawsock.go:165
	// _ = "end of CoverTab[6594]"
//line /snap/go/10455/src/net/iprawsock.go:165
	_go_fuzz_dep_.CoverTab[6595]++
						return n, err
//line /snap/go/10455/src/net/iprawsock.go:166
	// _ = "end of CoverTab[6595]"
}

// WriteTo implements the PacketConn WriteTo method.
func (c *IPConn) WriteTo(b []byte, addr Addr) (int, error) {
//line /snap/go/10455/src/net/iprawsock.go:170
	_go_fuzz_dep_.CoverTab[6600]++
						if !c.ok() {
//line /snap/go/10455/src/net/iprawsock.go:171
		_go_fuzz_dep_.CoverTab[528733]++
//line /snap/go/10455/src/net/iprawsock.go:171
		_go_fuzz_dep_.CoverTab[6604]++
							return 0, syscall.EINVAL
//line /snap/go/10455/src/net/iprawsock.go:172
		// _ = "end of CoverTab[6604]"
	} else {
//line /snap/go/10455/src/net/iprawsock.go:173
		_go_fuzz_dep_.CoverTab[528734]++
//line /snap/go/10455/src/net/iprawsock.go:173
		_go_fuzz_dep_.CoverTab[6605]++
//line /snap/go/10455/src/net/iprawsock.go:173
		// _ = "end of CoverTab[6605]"
//line /snap/go/10455/src/net/iprawsock.go:173
	}
//line /snap/go/10455/src/net/iprawsock.go:173
	// _ = "end of CoverTab[6600]"
//line /snap/go/10455/src/net/iprawsock.go:173
	_go_fuzz_dep_.CoverTab[6601]++
						a, ok := addr.(*IPAddr)
						if !ok {
//line /snap/go/10455/src/net/iprawsock.go:175
		_go_fuzz_dep_.CoverTab[528735]++
//line /snap/go/10455/src/net/iprawsock.go:175
		_go_fuzz_dep_.CoverTab[6606]++
							return 0, &OpError{Op: "write", Net: c.fd.net, Source: c.fd.laddr, Addr: addr, Err: syscall.EINVAL}
//line /snap/go/10455/src/net/iprawsock.go:176
		// _ = "end of CoverTab[6606]"
	} else {
//line /snap/go/10455/src/net/iprawsock.go:177
		_go_fuzz_dep_.CoverTab[528736]++
//line /snap/go/10455/src/net/iprawsock.go:177
		_go_fuzz_dep_.CoverTab[6607]++
//line /snap/go/10455/src/net/iprawsock.go:177
		// _ = "end of CoverTab[6607]"
//line /snap/go/10455/src/net/iprawsock.go:177
	}
//line /snap/go/10455/src/net/iprawsock.go:177
	// _ = "end of CoverTab[6601]"
//line /snap/go/10455/src/net/iprawsock.go:177
	_go_fuzz_dep_.CoverTab[6602]++
						n, err := c.writeTo(b, a)
						if err != nil {
//line /snap/go/10455/src/net/iprawsock.go:179
		_go_fuzz_dep_.CoverTab[528737]++
//line /snap/go/10455/src/net/iprawsock.go:179
		_go_fuzz_dep_.CoverTab[6608]++
							err = &OpError{Op: "write", Net: c.fd.net, Source: c.fd.laddr, Addr: a.opAddr(), Err: err}
//line /snap/go/10455/src/net/iprawsock.go:180
		// _ = "end of CoverTab[6608]"
	} else {
//line /snap/go/10455/src/net/iprawsock.go:181
		_go_fuzz_dep_.CoverTab[528738]++
//line /snap/go/10455/src/net/iprawsock.go:181
		_go_fuzz_dep_.CoverTab[6609]++
//line /snap/go/10455/src/net/iprawsock.go:181
		// _ = "end of CoverTab[6609]"
//line /snap/go/10455/src/net/iprawsock.go:181
	}
//line /snap/go/10455/src/net/iprawsock.go:181
	// _ = "end of CoverTab[6602]"
//line /snap/go/10455/src/net/iprawsock.go:181
	_go_fuzz_dep_.CoverTab[6603]++
						return n, err
//line /snap/go/10455/src/net/iprawsock.go:182
	// _ = "end of CoverTab[6603]"
}

// WriteMsgIP writes a message to addr via c, copying the payload from
//line /snap/go/10455/src/net/iprawsock.go:185
// b and the associated out-of-band data from oob. It returns the
//line /snap/go/10455/src/net/iprawsock.go:185
// number of payload and out-of-band bytes written.
//line /snap/go/10455/src/net/iprawsock.go:185
//
//line /snap/go/10455/src/net/iprawsock.go:185
// The packages golang.org/x/net/ipv4 and golang.org/x/net/ipv6 can be
//line /snap/go/10455/src/net/iprawsock.go:185
// used to manipulate IP-level socket options in oob.
//line /snap/go/10455/src/net/iprawsock.go:191
func (c *IPConn) WriteMsgIP(b, oob []byte, addr *IPAddr) (n, oobn int, err error) {
//line /snap/go/10455/src/net/iprawsock.go:191
	_go_fuzz_dep_.CoverTab[6610]++
						if !c.ok() {
//line /snap/go/10455/src/net/iprawsock.go:192
		_go_fuzz_dep_.CoverTab[528739]++
//line /snap/go/10455/src/net/iprawsock.go:192
		_go_fuzz_dep_.CoverTab[6613]++
							return 0, 0, syscall.EINVAL
//line /snap/go/10455/src/net/iprawsock.go:193
		// _ = "end of CoverTab[6613]"
	} else {
//line /snap/go/10455/src/net/iprawsock.go:194
		_go_fuzz_dep_.CoverTab[528740]++
//line /snap/go/10455/src/net/iprawsock.go:194
		_go_fuzz_dep_.CoverTab[6614]++
//line /snap/go/10455/src/net/iprawsock.go:194
		// _ = "end of CoverTab[6614]"
//line /snap/go/10455/src/net/iprawsock.go:194
	}
//line /snap/go/10455/src/net/iprawsock.go:194
	// _ = "end of CoverTab[6610]"
//line /snap/go/10455/src/net/iprawsock.go:194
	_go_fuzz_dep_.CoverTab[6611]++
						n, oobn, err = c.writeMsg(b, oob, addr)
						if err != nil {
//line /snap/go/10455/src/net/iprawsock.go:196
		_go_fuzz_dep_.CoverTab[528741]++
//line /snap/go/10455/src/net/iprawsock.go:196
		_go_fuzz_dep_.CoverTab[6615]++
							err = &OpError{Op: "write", Net: c.fd.net, Source: c.fd.laddr, Addr: addr.opAddr(), Err: err}
//line /snap/go/10455/src/net/iprawsock.go:197
		// _ = "end of CoverTab[6615]"
	} else {
//line /snap/go/10455/src/net/iprawsock.go:198
		_go_fuzz_dep_.CoverTab[528742]++
//line /snap/go/10455/src/net/iprawsock.go:198
		_go_fuzz_dep_.CoverTab[6616]++
//line /snap/go/10455/src/net/iprawsock.go:198
		// _ = "end of CoverTab[6616]"
//line /snap/go/10455/src/net/iprawsock.go:198
	}
//line /snap/go/10455/src/net/iprawsock.go:198
	// _ = "end of CoverTab[6611]"
//line /snap/go/10455/src/net/iprawsock.go:198
	_go_fuzz_dep_.CoverTab[6612]++
						return
//line /snap/go/10455/src/net/iprawsock.go:199
	// _ = "end of CoverTab[6612]"
}

func newIPConn(fd *netFD) *IPConn {
//line /snap/go/10455/src/net/iprawsock.go:202
	_go_fuzz_dep_.CoverTab[6617]++
//line /snap/go/10455/src/net/iprawsock.go:202
	return &IPConn{conn{fd}}
//line /snap/go/10455/src/net/iprawsock.go:202
	// _ = "end of CoverTab[6617]"
//line /snap/go/10455/src/net/iprawsock.go:202
}

// DialIP acts like Dial for IP networks.
//line /snap/go/10455/src/net/iprawsock.go:204
//
//line /snap/go/10455/src/net/iprawsock.go:204
// The network must be an IP network name; see func Dial for details.
//line /snap/go/10455/src/net/iprawsock.go:204
//
//line /snap/go/10455/src/net/iprawsock.go:204
// If laddr is nil, a local address is automatically chosen.
//line /snap/go/10455/src/net/iprawsock.go:204
// If the IP field of raddr is nil or an unspecified IP address, the
//line /snap/go/10455/src/net/iprawsock.go:204
// local system is assumed.
//line /snap/go/10455/src/net/iprawsock.go:211
func DialIP(network string, laddr, raddr *IPAddr) (*IPConn, error) {
//line /snap/go/10455/src/net/iprawsock.go:211
	_go_fuzz_dep_.CoverTab[6618]++
						if raddr == nil {
//line /snap/go/10455/src/net/iprawsock.go:212
		_go_fuzz_dep_.CoverTab[528743]++
//line /snap/go/10455/src/net/iprawsock.go:212
		_go_fuzz_dep_.CoverTab[6621]++
							return nil, &OpError{Op: "dial", Net: network, Source: laddr.opAddr(), Addr: nil, Err: errMissingAddress}
//line /snap/go/10455/src/net/iprawsock.go:213
		// _ = "end of CoverTab[6621]"
	} else {
//line /snap/go/10455/src/net/iprawsock.go:214
		_go_fuzz_dep_.CoverTab[528744]++
//line /snap/go/10455/src/net/iprawsock.go:214
		_go_fuzz_dep_.CoverTab[6622]++
//line /snap/go/10455/src/net/iprawsock.go:214
		// _ = "end of CoverTab[6622]"
//line /snap/go/10455/src/net/iprawsock.go:214
	}
//line /snap/go/10455/src/net/iprawsock.go:214
	// _ = "end of CoverTab[6618]"
//line /snap/go/10455/src/net/iprawsock.go:214
	_go_fuzz_dep_.CoverTab[6619]++
						sd := &sysDialer{network: network, address: raddr.String()}
						c, err := sd.dialIP(context.Background(), laddr, raddr)
						if err != nil {
//line /snap/go/10455/src/net/iprawsock.go:217
		_go_fuzz_dep_.CoverTab[528745]++
//line /snap/go/10455/src/net/iprawsock.go:217
		_go_fuzz_dep_.CoverTab[6623]++
							return nil, &OpError{Op: "dial", Net: network, Source: laddr.opAddr(), Addr: raddr.opAddr(), Err: err}
//line /snap/go/10455/src/net/iprawsock.go:218
		// _ = "end of CoverTab[6623]"
	} else {
//line /snap/go/10455/src/net/iprawsock.go:219
		_go_fuzz_dep_.CoverTab[528746]++
//line /snap/go/10455/src/net/iprawsock.go:219
		_go_fuzz_dep_.CoverTab[6624]++
//line /snap/go/10455/src/net/iprawsock.go:219
		// _ = "end of CoverTab[6624]"
//line /snap/go/10455/src/net/iprawsock.go:219
	}
//line /snap/go/10455/src/net/iprawsock.go:219
	// _ = "end of CoverTab[6619]"
//line /snap/go/10455/src/net/iprawsock.go:219
	_go_fuzz_dep_.CoverTab[6620]++
						return c, nil
//line /snap/go/10455/src/net/iprawsock.go:220
	// _ = "end of CoverTab[6620]"
}

// ListenIP acts like ListenPacket for IP networks.
//line /snap/go/10455/src/net/iprawsock.go:223
//
//line /snap/go/10455/src/net/iprawsock.go:223
// The network must be an IP network name; see func Dial for details.
//line /snap/go/10455/src/net/iprawsock.go:223
//
//line /snap/go/10455/src/net/iprawsock.go:223
// If the IP field of laddr is nil or an unspecified IP address,
//line /snap/go/10455/src/net/iprawsock.go:223
// ListenIP listens on all available IP addresses of the local system
//line /snap/go/10455/src/net/iprawsock.go:223
// except multicast IP addresses.
//line /snap/go/10455/src/net/iprawsock.go:230
func ListenIP(network string, laddr *IPAddr) (*IPConn, error) {
//line /snap/go/10455/src/net/iprawsock.go:230
	_go_fuzz_dep_.CoverTab[6625]++
						if laddr == nil {
//line /snap/go/10455/src/net/iprawsock.go:231
		_go_fuzz_dep_.CoverTab[528747]++
//line /snap/go/10455/src/net/iprawsock.go:231
		_go_fuzz_dep_.CoverTab[6628]++
							laddr = &IPAddr{}
//line /snap/go/10455/src/net/iprawsock.go:232
		// _ = "end of CoverTab[6628]"
	} else {
//line /snap/go/10455/src/net/iprawsock.go:233
		_go_fuzz_dep_.CoverTab[528748]++
//line /snap/go/10455/src/net/iprawsock.go:233
		_go_fuzz_dep_.CoverTab[6629]++
//line /snap/go/10455/src/net/iprawsock.go:233
		// _ = "end of CoverTab[6629]"
//line /snap/go/10455/src/net/iprawsock.go:233
	}
//line /snap/go/10455/src/net/iprawsock.go:233
	// _ = "end of CoverTab[6625]"
//line /snap/go/10455/src/net/iprawsock.go:233
	_go_fuzz_dep_.CoverTab[6626]++
						sl := &sysListener{network: network, address: laddr.String()}
						c, err := sl.listenIP(context.Background(), laddr)
						if err != nil {
//line /snap/go/10455/src/net/iprawsock.go:236
		_go_fuzz_dep_.CoverTab[528749]++
//line /snap/go/10455/src/net/iprawsock.go:236
		_go_fuzz_dep_.CoverTab[6630]++
							return nil, &OpError{Op: "listen", Net: network, Source: nil, Addr: laddr.opAddr(), Err: err}
//line /snap/go/10455/src/net/iprawsock.go:237
		// _ = "end of CoverTab[6630]"
	} else {
//line /snap/go/10455/src/net/iprawsock.go:238
		_go_fuzz_dep_.CoverTab[528750]++
//line /snap/go/10455/src/net/iprawsock.go:238
		_go_fuzz_dep_.CoverTab[6631]++
//line /snap/go/10455/src/net/iprawsock.go:238
		// _ = "end of CoverTab[6631]"
//line /snap/go/10455/src/net/iprawsock.go:238
	}
//line /snap/go/10455/src/net/iprawsock.go:238
	// _ = "end of CoverTab[6626]"
//line /snap/go/10455/src/net/iprawsock.go:238
	_go_fuzz_dep_.CoverTab[6627]++
						return c, nil
//line /snap/go/10455/src/net/iprawsock.go:239
	// _ = "end of CoverTab[6627]"
}

//line /snap/go/10455/src/net/iprawsock.go:240
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /snap/go/10455/src/net/iprawsock.go:240
var _ = _go_fuzz_dep_.CoverTab
