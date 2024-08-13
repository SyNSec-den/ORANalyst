// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /snap/go/10455/src/net/tcpsock.go:5
package net

//line /snap/go/10455/src/net/tcpsock.go:5
import (
//line /snap/go/10455/src/net/tcpsock.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /snap/go/10455/src/net/tcpsock.go:5
)
//line /snap/go/10455/src/net/tcpsock.go:5
import (
//line /snap/go/10455/src/net/tcpsock.go:5
	_atomic_ "sync/atomic"
//line /snap/go/10455/src/net/tcpsock.go:5
)

import (
	"context"
	"internal/itoa"
	"io"
	"net/netip"
	"os"
	"syscall"
	"time"
)

//line /snap/go/10455/src/net/tcpsock.go:20
// TCPAddr represents the address of a TCP end point.
type TCPAddr struct {
	IP	IP
	Port	int
	Zone	string	// IPv6 scoped addressing zone
}

// AddrPort returns the TCPAddr a as a netip.AddrPort.
//line /snap/go/10455/src/net/tcpsock.go:27
//
//line /snap/go/10455/src/net/tcpsock.go:27
// If a.Port does not fit in a uint16, it's silently truncated.
//line /snap/go/10455/src/net/tcpsock.go:27
//
//line /snap/go/10455/src/net/tcpsock.go:27
// If a is nil, a zero value is returned.
//line /snap/go/10455/src/net/tcpsock.go:32
func (a *TCPAddr) AddrPort() netip.AddrPort {
//line /snap/go/10455/src/net/tcpsock.go:32
	_go_fuzz_dep_.CoverTab[8295]++
						if a == nil {
//line /snap/go/10455/src/net/tcpsock.go:33
		_go_fuzz_dep_.CoverTab[529807]++
//line /snap/go/10455/src/net/tcpsock.go:33
		_go_fuzz_dep_.CoverTab[8297]++
							return netip.AddrPort{}
//line /snap/go/10455/src/net/tcpsock.go:34
		// _ = "end of CoverTab[8297]"
	} else {
//line /snap/go/10455/src/net/tcpsock.go:35
		_go_fuzz_dep_.CoverTab[529808]++
//line /snap/go/10455/src/net/tcpsock.go:35
		_go_fuzz_dep_.CoverTab[8298]++
//line /snap/go/10455/src/net/tcpsock.go:35
		// _ = "end of CoverTab[8298]"
//line /snap/go/10455/src/net/tcpsock.go:35
	}
//line /snap/go/10455/src/net/tcpsock.go:35
	// _ = "end of CoverTab[8295]"
//line /snap/go/10455/src/net/tcpsock.go:35
	_go_fuzz_dep_.CoverTab[8296]++
						na, _ := netip.AddrFromSlice(a.IP)
						na = na.WithZone(a.Zone)
						return netip.AddrPortFrom(na, uint16(a.Port))
//line /snap/go/10455/src/net/tcpsock.go:38
	// _ = "end of CoverTab[8296]"
}

// Network returns the address's network name, "tcp".
func (a *TCPAddr) Network() string {
//line /snap/go/10455/src/net/tcpsock.go:42
	_go_fuzz_dep_.CoverTab[8299]++
//line /snap/go/10455/src/net/tcpsock.go:42
	return "tcp"
//line /snap/go/10455/src/net/tcpsock.go:42
	// _ = "end of CoverTab[8299]"
//line /snap/go/10455/src/net/tcpsock.go:42
}

func (a *TCPAddr) String() string {
//line /snap/go/10455/src/net/tcpsock.go:44
	_go_fuzz_dep_.CoverTab[8300]++
						if a == nil {
//line /snap/go/10455/src/net/tcpsock.go:45
		_go_fuzz_dep_.CoverTab[529809]++
//line /snap/go/10455/src/net/tcpsock.go:45
		_go_fuzz_dep_.CoverTab[8303]++
							return "<nil>"
//line /snap/go/10455/src/net/tcpsock.go:46
		// _ = "end of CoverTab[8303]"
	} else {
//line /snap/go/10455/src/net/tcpsock.go:47
		_go_fuzz_dep_.CoverTab[529810]++
//line /snap/go/10455/src/net/tcpsock.go:47
		_go_fuzz_dep_.CoverTab[8304]++
//line /snap/go/10455/src/net/tcpsock.go:47
		// _ = "end of CoverTab[8304]"
//line /snap/go/10455/src/net/tcpsock.go:47
	}
//line /snap/go/10455/src/net/tcpsock.go:47
	// _ = "end of CoverTab[8300]"
//line /snap/go/10455/src/net/tcpsock.go:47
	_go_fuzz_dep_.CoverTab[8301]++
						ip := ipEmptyString(a.IP)
						if a.Zone != "" {
//line /snap/go/10455/src/net/tcpsock.go:49
		_go_fuzz_dep_.CoverTab[529811]++
//line /snap/go/10455/src/net/tcpsock.go:49
		_go_fuzz_dep_.CoverTab[8305]++
							return JoinHostPort(ip+"%"+a.Zone, itoa.Itoa(a.Port))
//line /snap/go/10455/src/net/tcpsock.go:50
		// _ = "end of CoverTab[8305]"
	} else {
//line /snap/go/10455/src/net/tcpsock.go:51
		_go_fuzz_dep_.CoverTab[529812]++
//line /snap/go/10455/src/net/tcpsock.go:51
		_go_fuzz_dep_.CoverTab[8306]++
//line /snap/go/10455/src/net/tcpsock.go:51
		// _ = "end of CoverTab[8306]"
//line /snap/go/10455/src/net/tcpsock.go:51
	}
//line /snap/go/10455/src/net/tcpsock.go:51
	// _ = "end of CoverTab[8301]"
//line /snap/go/10455/src/net/tcpsock.go:51
	_go_fuzz_dep_.CoverTab[8302]++
						return JoinHostPort(ip, itoa.Itoa(a.Port))
//line /snap/go/10455/src/net/tcpsock.go:52
	// _ = "end of CoverTab[8302]"
}

func (a *TCPAddr) isWildcard() bool {
//line /snap/go/10455/src/net/tcpsock.go:55
	_go_fuzz_dep_.CoverTab[8307]++
						if a == nil || func() bool {
//line /snap/go/10455/src/net/tcpsock.go:56
		_go_fuzz_dep_.CoverTab[8309]++
//line /snap/go/10455/src/net/tcpsock.go:56
		return a.IP == nil
//line /snap/go/10455/src/net/tcpsock.go:56
		// _ = "end of CoverTab[8309]"
//line /snap/go/10455/src/net/tcpsock.go:56
	}() {
//line /snap/go/10455/src/net/tcpsock.go:56
		_go_fuzz_dep_.CoverTab[529813]++
//line /snap/go/10455/src/net/tcpsock.go:56
		_go_fuzz_dep_.CoverTab[8310]++
							return true
//line /snap/go/10455/src/net/tcpsock.go:57
		// _ = "end of CoverTab[8310]"
	} else {
//line /snap/go/10455/src/net/tcpsock.go:58
		_go_fuzz_dep_.CoverTab[529814]++
//line /snap/go/10455/src/net/tcpsock.go:58
		_go_fuzz_dep_.CoverTab[8311]++
//line /snap/go/10455/src/net/tcpsock.go:58
		// _ = "end of CoverTab[8311]"
//line /snap/go/10455/src/net/tcpsock.go:58
	}
//line /snap/go/10455/src/net/tcpsock.go:58
	// _ = "end of CoverTab[8307]"
//line /snap/go/10455/src/net/tcpsock.go:58
	_go_fuzz_dep_.CoverTab[8308]++
						return a.IP.IsUnspecified()
//line /snap/go/10455/src/net/tcpsock.go:59
	// _ = "end of CoverTab[8308]"
}

func (a *TCPAddr) opAddr() Addr {
//line /snap/go/10455/src/net/tcpsock.go:62
	_go_fuzz_dep_.CoverTab[8312]++
						if a == nil {
//line /snap/go/10455/src/net/tcpsock.go:63
		_go_fuzz_dep_.CoverTab[529815]++
//line /snap/go/10455/src/net/tcpsock.go:63
		_go_fuzz_dep_.CoverTab[8314]++
							return nil
//line /snap/go/10455/src/net/tcpsock.go:64
		// _ = "end of CoverTab[8314]"
	} else {
//line /snap/go/10455/src/net/tcpsock.go:65
		_go_fuzz_dep_.CoverTab[529816]++
//line /snap/go/10455/src/net/tcpsock.go:65
		_go_fuzz_dep_.CoverTab[8315]++
//line /snap/go/10455/src/net/tcpsock.go:65
		// _ = "end of CoverTab[8315]"
//line /snap/go/10455/src/net/tcpsock.go:65
	}
//line /snap/go/10455/src/net/tcpsock.go:65
	// _ = "end of CoverTab[8312]"
//line /snap/go/10455/src/net/tcpsock.go:65
	_go_fuzz_dep_.CoverTab[8313]++
						return a
//line /snap/go/10455/src/net/tcpsock.go:66
	// _ = "end of CoverTab[8313]"
}

// ResolveTCPAddr returns an address of TCP end point.
//line /snap/go/10455/src/net/tcpsock.go:69
//
//line /snap/go/10455/src/net/tcpsock.go:69
// The network must be a TCP network name.
//line /snap/go/10455/src/net/tcpsock.go:69
//
//line /snap/go/10455/src/net/tcpsock.go:69
// If the host in the address parameter is not a literal IP address or
//line /snap/go/10455/src/net/tcpsock.go:69
// the port is not a literal port number, ResolveTCPAddr resolves the
//line /snap/go/10455/src/net/tcpsock.go:69
// address to an address of TCP end point.
//line /snap/go/10455/src/net/tcpsock.go:69
// Otherwise, it parses the address as a pair of literal IP address
//line /snap/go/10455/src/net/tcpsock.go:69
// and port number.
//line /snap/go/10455/src/net/tcpsock.go:69
// The address parameter can use a host name, but this is not
//line /snap/go/10455/src/net/tcpsock.go:69
// recommended, because it will return at most one of the host name's
//line /snap/go/10455/src/net/tcpsock.go:69
// IP addresses.
//line /snap/go/10455/src/net/tcpsock.go:69
//
//line /snap/go/10455/src/net/tcpsock.go:69
// See func Dial for a description of the network and address
//line /snap/go/10455/src/net/tcpsock.go:69
// parameters.
//line /snap/go/10455/src/net/tcpsock.go:84
func ResolveTCPAddr(network, address string) (*TCPAddr, error) {
//line /snap/go/10455/src/net/tcpsock.go:84
	_go_fuzz_dep_.CoverTab[8316]++
						switch network {
	case "tcp", "tcp4", "tcp6":
//line /snap/go/10455/src/net/tcpsock.go:86
		_go_fuzz_dep_.CoverTab[529817]++
//line /snap/go/10455/src/net/tcpsock.go:86
		_go_fuzz_dep_.CoverTab[8319]++
//line /snap/go/10455/src/net/tcpsock.go:86
		// _ = "end of CoverTab[8319]"
	case "":
//line /snap/go/10455/src/net/tcpsock.go:87
		_go_fuzz_dep_.CoverTab[529818]++
//line /snap/go/10455/src/net/tcpsock.go:87
		_go_fuzz_dep_.CoverTab[8320]++
							network = "tcp"
//line /snap/go/10455/src/net/tcpsock.go:88
		// _ = "end of CoverTab[8320]"
	default:
//line /snap/go/10455/src/net/tcpsock.go:89
		_go_fuzz_dep_.CoverTab[529819]++
//line /snap/go/10455/src/net/tcpsock.go:89
		_go_fuzz_dep_.CoverTab[8321]++
							return nil, UnknownNetworkError(network)
//line /snap/go/10455/src/net/tcpsock.go:90
		// _ = "end of CoverTab[8321]"
	}
//line /snap/go/10455/src/net/tcpsock.go:91
	// _ = "end of CoverTab[8316]"
//line /snap/go/10455/src/net/tcpsock.go:91
	_go_fuzz_dep_.CoverTab[8317]++
						addrs, err := DefaultResolver.internetAddrList(context.Background(), network, address)
						if err != nil {
//line /snap/go/10455/src/net/tcpsock.go:93
		_go_fuzz_dep_.CoverTab[529820]++
//line /snap/go/10455/src/net/tcpsock.go:93
		_go_fuzz_dep_.CoverTab[8322]++
							return nil, err
//line /snap/go/10455/src/net/tcpsock.go:94
		// _ = "end of CoverTab[8322]"
	} else {
//line /snap/go/10455/src/net/tcpsock.go:95
		_go_fuzz_dep_.CoverTab[529821]++
//line /snap/go/10455/src/net/tcpsock.go:95
		_go_fuzz_dep_.CoverTab[8323]++
//line /snap/go/10455/src/net/tcpsock.go:95
		// _ = "end of CoverTab[8323]"
//line /snap/go/10455/src/net/tcpsock.go:95
	}
//line /snap/go/10455/src/net/tcpsock.go:95
	// _ = "end of CoverTab[8317]"
//line /snap/go/10455/src/net/tcpsock.go:95
	_go_fuzz_dep_.CoverTab[8318]++
						return addrs.forResolve(network, address).(*TCPAddr), nil
//line /snap/go/10455/src/net/tcpsock.go:96
	// _ = "end of CoverTab[8318]"
}

// TCPAddrFromAddrPort returns addr as a TCPAddr. If addr.IsValid() is false,
//line /snap/go/10455/src/net/tcpsock.go:99
// then the returned TCPAddr will contain a nil IP field, indicating an
//line /snap/go/10455/src/net/tcpsock.go:99
// address family-agnostic unspecified address.
//line /snap/go/10455/src/net/tcpsock.go:102
func TCPAddrFromAddrPort(addr netip.AddrPort) *TCPAddr {
//line /snap/go/10455/src/net/tcpsock.go:102
	_go_fuzz_dep_.CoverTab[8324]++
						return &TCPAddr{
		IP:	addr.Addr().AsSlice(),
		Zone:	addr.Addr().Zone(),
		Port:	int(addr.Port()),
	}
//line /snap/go/10455/src/net/tcpsock.go:107
	// _ = "end of CoverTab[8324]"
}

// TCPConn is an implementation of the Conn interface for TCP network
//line /snap/go/10455/src/net/tcpsock.go:110
// connections.
//line /snap/go/10455/src/net/tcpsock.go:112
type TCPConn struct {
	conn
}

// SyscallConn returns a raw network connection.
//line /snap/go/10455/src/net/tcpsock.go:116
// This implements the syscall.Conn interface.
//line /snap/go/10455/src/net/tcpsock.go:118
func (c *TCPConn) SyscallConn() (syscall.RawConn, error) {
//line /snap/go/10455/src/net/tcpsock.go:118
	_go_fuzz_dep_.CoverTab[8325]++
						if !c.ok() {
//line /snap/go/10455/src/net/tcpsock.go:119
		_go_fuzz_dep_.CoverTab[529822]++
//line /snap/go/10455/src/net/tcpsock.go:119
		_go_fuzz_dep_.CoverTab[8327]++
							return nil, syscall.EINVAL
//line /snap/go/10455/src/net/tcpsock.go:120
		// _ = "end of CoverTab[8327]"
	} else {
//line /snap/go/10455/src/net/tcpsock.go:121
		_go_fuzz_dep_.CoverTab[529823]++
//line /snap/go/10455/src/net/tcpsock.go:121
		_go_fuzz_dep_.CoverTab[8328]++
//line /snap/go/10455/src/net/tcpsock.go:121
		// _ = "end of CoverTab[8328]"
//line /snap/go/10455/src/net/tcpsock.go:121
	}
//line /snap/go/10455/src/net/tcpsock.go:121
	// _ = "end of CoverTab[8325]"
//line /snap/go/10455/src/net/tcpsock.go:121
	_go_fuzz_dep_.CoverTab[8326]++
						return newRawConn(c.fd)
//line /snap/go/10455/src/net/tcpsock.go:122
	// _ = "end of CoverTab[8326]"
}

// ReadFrom implements the io.ReaderFrom ReadFrom method.
func (c *TCPConn) ReadFrom(r io.Reader) (int64, error) {
//line /snap/go/10455/src/net/tcpsock.go:126
	_go_fuzz_dep_.CoverTab[8329]++
						if !c.ok() {
//line /snap/go/10455/src/net/tcpsock.go:127
		_go_fuzz_dep_.CoverTab[529824]++
//line /snap/go/10455/src/net/tcpsock.go:127
		_go_fuzz_dep_.CoverTab[8332]++
							return 0, syscall.EINVAL
//line /snap/go/10455/src/net/tcpsock.go:128
		// _ = "end of CoverTab[8332]"
	} else {
//line /snap/go/10455/src/net/tcpsock.go:129
		_go_fuzz_dep_.CoverTab[529825]++
//line /snap/go/10455/src/net/tcpsock.go:129
		_go_fuzz_dep_.CoverTab[8333]++
//line /snap/go/10455/src/net/tcpsock.go:129
		// _ = "end of CoverTab[8333]"
//line /snap/go/10455/src/net/tcpsock.go:129
	}
//line /snap/go/10455/src/net/tcpsock.go:129
	// _ = "end of CoverTab[8329]"
//line /snap/go/10455/src/net/tcpsock.go:129
	_go_fuzz_dep_.CoverTab[8330]++
						n, err := c.readFrom(r)
						if err != nil && func() bool {
//line /snap/go/10455/src/net/tcpsock.go:131
		_go_fuzz_dep_.CoverTab[8334]++
//line /snap/go/10455/src/net/tcpsock.go:131
		return err != io.EOF
//line /snap/go/10455/src/net/tcpsock.go:131
		// _ = "end of CoverTab[8334]"
//line /snap/go/10455/src/net/tcpsock.go:131
	}() {
//line /snap/go/10455/src/net/tcpsock.go:131
		_go_fuzz_dep_.CoverTab[529826]++
//line /snap/go/10455/src/net/tcpsock.go:131
		_go_fuzz_dep_.CoverTab[8335]++
							err = &OpError{Op: "readfrom", Net: c.fd.net, Source: c.fd.laddr, Addr: c.fd.raddr, Err: err}
//line /snap/go/10455/src/net/tcpsock.go:132
		// _ = "end of CoverTab[8335]"
	} else {
//line /snap/go/10455/src/net/tcpsock.go:133
		_go_fuzz_dep_.CoverTab[529827]++
//line /snap/go/10455/src/net/tcpsock.go:133
		_go_fuzz_dep_.CoverTab[8336]++
//line /snap/go/10455/src/net/tcpsock.go:133
		// _ = "end of CoverTab[8336]"
//line /snap/go/10455/src/net/tcpsock.go:133
	}
//line /snap/go/10455/src/net/tcpsock.go:133
	// _ = "end of CoverTab[8330]"
//line /snap/go/10455/src/net/tcpsock.go:133
	_go_fuzz_dep_.CoverTab[8331]++
						return n, err
//line /snap/go/10455/src/net/tcpsock.go:134
	// _ = "end of CoverTab[8331]"
}

// CloseRead shuts down the reading side of the TCP connection.
//line /snap/go/10455/src/net/tcpsock.go:137
// Most callers should just use Close.
//line /snap/go/10455/src/net/tcpsock.go:139
func (c *TCPConn) CloseRead() error {
//line /snap/go/10455/src/net/tcpsock.go:139
	_go_fuzz_dep_.CoverTab[8337]++
						if !c.ok() {
//line /snap/go/10455/src/net/tcpsock.go:140
		_go_fuzz_dep_.CoverTab[529828]++
//line /snap/go/10455/src/net/tcpsock.go:140
		_go_fuzz_dep_.CoverTab[8340]++
							return syscall.EINVAL
//line /snap/go/10455/src/net/tcpsock.go:141
		// _ = "end of CoverTab[8340]"
	} else {
//line /snap/go/10455/src/net/tcpsock.go:142
		_go_fuzz_dep_.CoverTab[529829]++
//line /snap/go/10455/src/net/tcpsock.go:142
		_go_fuzz_dep_.CoverTab[8341]++
//line /snap/go/10455/src/net/tcpsock.go:142
		// _ = "end of CoverTab[8341]"
//line /snap/go/10455/src/net/tcpsock.go:142
	}
//line /snap/go/10455/src/net/tcpsock.go:142
	// _ = "end of CoverTab[8337]"
//line /snap/go/10455/src/net/tcpsock.go:142
	_go_fuzz_dep_.CoverTab[8338]++
						if err := c.fd.closeRead(); err != nil {
//line /snap/go/10455/src/net/tcpsock.go:143
		_go_fuzz_dep_.CoverTab[529830]++
//line /snap/go/10455/src/net/tcpsock.go:143
		_go_fuzz_dep_.CoverTab[8342]++
							return &OpError{Op: "close", Net: c.fd.net, Source: c.fd.laddr, Addr: c.fd.raddr, Err: err}
//line /snap/go/10455/src/net/tcpsock.go:144
		// _ = "end of CoverTab[8342]"
	} else {
//line /snap/go/10455/src/net/tcpsock.go:145
		_go_fuzz_dep_.CoverTab[529831]++
//line /snap/go/10455/src/net/tcpsock.go:145
		_go_fuzz_dep_.CoverTab[8343]++
//line /snap/go/10455/src/net/tcpsock.go:145
		// _ = "end of CoverTab[8343]"
//line /snap/go/10455/src/net/tcpsock.go:145
	}
//line /snap/go/10455/src/net/tcpsock.go:145
	// _ = "end of CoverTab[8338]"
//line /snap/go/10455/src/net/tcpsock.go:145
	_go_fuzz_dep_.CoverTab[8339]++
						return nil
//line /snap/go/10455/src/net/tcpsock.go:146
	// _ = "end of CoverTab[8339]"
}

// CloseWrite shuts down the writing side of the TCP connection.
//line /snap/go/10455/src/net/tcpsock.go:149
// Most callers should just use Close.
//line /snap/go/10455/src/net/tcpsock.go:151
func (c *TCPConn) CloseWrite() error {
//line /snap/go/10455/src/net/tcpsock.go:151
	_go_fuzz_dep_.CoverTab[8344]++
						if !c.ok() {
//line /snap/go/10455/src/net/tcpsock.go:152
		_go_fuzz_dep_.CoverTab[529832]++
//line /snap/go/10455/src/net/tcpsock.go:152
		_go_fuzz_dep_.CoverTab[8347]++
							return syscall.EINVAL
//line /snap/go/10455/src/net/tcpsock.go:153
		// _ = "end of CoverTab[8347]"
	} else {
//line /snap/go/10455/src/net/tcpsock.go:154
		_go_fuzz_dep_.CoverTab[529833]++
//line /snap/go/10455/src/net/tcpsock.go:154
		_go_fuzz_dep_.CoverTab[8348]++
//line /snap/go/10455/src/net/tcpsock.go:154
		// _ = "end of CoverTab[8348]"
//line /snap/go/10455/src/net/tcpsock.go:154
	}
//line /snap/go/10455/src/net/tcpsock.go:154
	// _ = "end of CoverTab[8344]"
//line /snap/go/10455/src/net/tcpsock.go:154
	_go_fuzz_dep_.CoverTab[8345]++
						if err := c.fd.closeWrite(); err != nil {
//line /snap/go/10455/src/net/tcpsock.go:155
		_go_fuzz_dep_.CoverTab[529834]++
//line /snap/go/10455/src/net/tcpsock.go:155
		_go_fuzz_dep_.CoverTab[8349]++
							return &OpError{Op: "close", Net: c.fd.net, Source: c.fd.laddr, Addr: c.fd.raddr, Err: err}
//line /snap/go/10455/src/net/tcpsock.go:156
		// _ = "end of CoverTab[8349]"
	} else {
//line /snap/go/10455/src/net/tcpsock.go:157
		_go_fuzz_dep_.CoverTab[529835]++
//line /snap/go/10455/src/net/tcpsock.go:157
		_go_fuzz_dep_.CoverTab[8350]++
//line /snap/go/10455/src/net/tcpsock.go:157
		// _ = "end of CoverTab[8350]"
//line /snap/go/10455/src/net/tcpsock.go:157
	}
//line /snap/go/10455/src/net/tcpsock.go:157
	// _ = "end of CoverTab[8345]"
//line /snap/go/10455/src/net/tcpsock.go:157
	_go_fuzz_dep_.CoverTab[8346]++
						return nil
//line /snap/go/10455/src/net/tcpsock.go:158
	// _ = "end of CoverTab[8346]"
}

// SetLinger sets the behavior of Close on a connection which still
//line /snap/go/10455/src/net/tcpsock.go:161
// has data waiting to be sent or to be acknowledged.
//line /snap/go/10455/src/net/tcpsock.go:161
//
//line /snap/go/10455/src/net/tcpsock.go:161
// If sec < 0 (the default), the operating system finishes sending the
//line /snap/go/10455/src/net/tcpsock.go:161
// data in the background.
//line /snap/go/10455/src/net/tcpsock.go:161
//
//line /snap/go/10455/src/net/tcpsock.go:161
// If sec == 0, the operating system discards any unsent or
//line /snap/go/10455/src/net/tcpsock.go:161
// unacknowledged data.
//line /snap/go/10455/src/net/tcpsock.go:161
//
//line /snap/go/10455/src/net/tcpsock.go:161
// If sec > 0, the data is sent in the background as with sec < 0.
//line /snap/go/10455/src/net/tcpsock.go:161
// On some operating systems including Linux, this may cause Close to block
//line /snap/go/10455/src/net/tcpsock.go:161
// until all data has been sent or discarded.
//line /snap/go/10455/src/net/tcpsock.go:161
// On some operating systems after sec seconds have elapsed any remaining
//line /snap/go/10455/src/net/tcpsock.go:161
// unsent data may be discarded.
//line /snap/go/10455/src/net/tcpsock.go:175
func (c *TCPConn) SetLinger(sec int) error {
//line /snap/go/10455/src/net/tcpsock.go:175
	_go_fuzz_dep_.CoverTab[8351]++
						if !c.ok() {
//line /snap/go/10455/src/net/tcpsock.go:176
		_go_fuzz_dep_.CoverTab[529836]++
//line /snap/go/10455/src/net/tcpsock.go:176
		_go_fuzz_dep_.CoverTab[8354]++
							return syscall.EINVAL
//line /snap/go/10455/src/net/tcpsock.go:177
		// _ = "end of CoverTab[8354]"
	} else {
//line /snap/go/10455/src/net/tcpsock.go:178
		_go_fuzz_dep_.CoverTab[529837]++
//line /snap/go/10455/src/net/tcpsock.go:178
		_go_fuzz_dep_.CoverTab[8355]++
//line /snap/go/10455/src/net/tcpsock.go:178
		// _ = "end of CoverTab[8355]"
//line /snap/go/10455/src/net/tcpsock.go:178
	}
//line /snap/go/10455/src/net/tcpsock.go:178
	// _ = "end of CoverTab[8351]"
//line /snap/go/10455/src/net/tcpsock.go:178
	_go_fuzz_dep_.CoverTab[8352]++
						if err := setLinger(c.fd, sec); err != nil {
//line /snap/go/10455/src/net/tcpsock.go:179
		_go_fuzz_dep_.CoverTab[529838]++
//line /snap/go/10455/src/net/tcpsock.go:179
		_go_fuzz_dep_.CoverTab[8356]++
							return &OpError{Op: "set", Net: c.fd.net, Source: c.fd.laddr, Addr: c.fd.raddr, Err: err}
//line /snap/go/10455/src/net/tcpsock.go:180
		// _ = "end of CoverTab[8356]"
	} else {
//line /snap/go/10455/src/net/tcpsock.go:181
		_go_fuzz_dep_.CoverTab[529839]++
//line /snap/go/10455/src/net/tcpsock.go:181
		_go_fuzz_dep_.CoverTab[8357]++
//line /snap/go/10455/src/net/tcpsock.go:181
		// _ = "end of CoverTab[8357]"
//line /snap/go/10455/src/net/tcpsock.go:181
	}
//line /snap/go/10455/src/net/tcpsock.go:181
	// _ = "end of CoverTab[8352]"
//line /snap/go/10455/src/net/tcpsock.go:181
	_go_fuzz_dep_.CoverTab[8353]++
						return nil
//line /snap/go/10455/src/net/tcpsock.go:182
	// _ = "end of CoverTab[8353]"
}

// SetKeepAlive sets whether the operating system should send
//line /snap/go/10455/src/net/tcpsock.go:185
// keep-alive messages on the connection.
//line /snap/go/10455/src/net/tcpsock.go:187
func (c *TCPConn) SetKeepAlive(keepalive bool) error {
//line /snap/go/10455/src/net/tcpsock.go:187
	_go_fuzz_dep_.CoverTab[8358]++
						if !c.ok() {
//line /snap/go/10455/src/net/tcpsock.go:188
		_go_fuzz_dep_.CoverTab[529840]++
//line /snap/go/10455/src/net/tcpsock.go:188
		_go_fuzz_dep_.CoverTab[8361]++
							return syscall.EINVAL
//line /snap/go/10455/src/net/tcpsock.go:189
		// _ = "end of CoverTab[8361]"
	} else {
//line /snap/go/10455/src/net/tcpsock.go:190
		_go_fuzz_dep_.CoverTab[529841]++
//line /snap/go/10455/src/net/tcpsock.go:190
		_go_fuzz_dep_.CoverTab[8362]++
//line /snap/go/10455/src/net/tcpsock.go:190
		// _ = "end of CoverTab[8362]"
//line /snap/go/10455/src/net/tcpsock.go:190
	}
//line /snap/go/10455/src/net/tcpsock.go:190
	// _ = "end of CoverTab[8358]"
//line /snap/go/10455/src/net/tcpsock.go:190
	_go_fuzz_dep_.CoverTab[8359]++
						if err := setKeepAlive(c.fd, keepalive); err != nil {
//line /snap/go/10455/src/net/tcpsock.go:191
		_go_fuzz_dep_.CoverTab[529842]++
//line /snap/go/10455/src/net/tcpsock.go:191
		_go_fuzz_dep_.CoverTab[8363]++
							return &OpError{Op: "set", Net: c.fd.net, Source: c.fd.laddr, Addr: c.fd.raddr, Err: err}
//line /snap/go/10455/src/net/tcpsock.go:192
		// _ = "end of CoverTab[8363]"
	} else {
//line /snap/go/10455/src/net/tcpsock.go:193
		_go_fuzz_dep_.CoverTab[529843]++
//line /snap/go/10455/src/net/tcpsock.go:193
		_go_fuzz_dep_.CoverTab[8364]++
//line /snap/go/10455/src/net/tcpsock.go:193
		// _ = "end of CoverTab[8364]"
//line /snap/go/10455/src/net/tcpsock.go:193
	}
//line /snap/go/10455/src/net/tcpsock.go:193
	// _ = "end of CoverTab[8359]"
//line /snap/go/10455/src/net/tcpsock.go:193
	_go_fuzz_dep_.CoverTab[8360]++
						return nil
//line /snap/go/10455/src/net/tcpsock.go:194
	// _ = "end of CoverTab[8360]"
}

// SetKeepAlivePeriod sets period between keep-alives.
func (c *TCPConn) SetKeepAlivePeriod(d time.Duration) error {
//line /snap/go/10455/src/net/tcpsock.go:198
	_go_fuzz_dep_.CoverTab[8365]++
						if !c.ok() {
//line /snap/go/10455/src/net/tcpsock.go:199
		_go_fuzz_dep_.CoverTab[529844]++
//line /snap/go/10455/src/net/tcpsock.go:199
		_go_fuzz_dep_.CoverTab[8368]++
							return syscall.EINVAL
//line /snap/go/10455/src/net/tcpsock.go:200
		// _ = "end of CoverTab[8368]"
	} else {
//line /snap/go/10455/src/net/tcpsock.go:201
		_go_fuzz_dep_.CoverTab[529845]++
//line /snap/go/10455/src/net/tcpsock.go:201
		_go_fuzz_dep_.CoverTab[8369]++
//line /snap/go/10455/src/net/tcpsock.go:201
		// _ = "end of CoverTab[8369]"
//line /snap/go/10455/src/net/tcpsock.go:201
	}
//line /snap/go/10455/src/net/tcpsock.go:201
	// _ = "end of CoverTab[8365]"
//line /snap/go/10455/src/net/tcpsock.go:201
	_go_fuzz_dep_.CoverTab[8366]++
						if err := setKeepAlivePeriod(c.fd, d); err != nil {
//line /snap/go/10455/src/net/tcpsock.go:202
		_go_fuzz_dep_.CoverTab[529846]++
//line /snap/go/10455/src/net/tcpsock.go:202
		_go_fuzz_dep_.CoverTab[8370]++
							return &OpError{Op: "set", Net: c.fd.net, Source: c.fd.laddr, Addr: c.fd.raddr, Err: err}
//line /snap/go/10455/src/net/tcpsock.go:203
		// _ = "end of CoverTab[8370]"
	} else {
//line /snap/go/10455/src/net/tcpsock.go:204
		_go_fuzz_dep_.CoverTab[529847]++
//line /snap/go/10455/src/net/tcpsock.go:204
		_go_fuzz_dep_.CoverTab[8371]++
//line /snap/go/10455/src/net/tcpsock.go:204
		// _ = "end of CoverTab[8371]"
//line /snap/go/10455/src/net/tcpsock.go:204
	}
//line /snap/go/10455/src/net/tcpsock.go:204
	// _ = "end of CoverTab[8366]"
//line /snap/go/10455/src/net/tcpsock.go:204
	_go_fuzz_dep_.CoverTab[8367]++
						return nil
//line /snap/go/10455/src/net/tcpsock.go:205
	// _ = "end of CoverTab[8367]"
}

// SetNoDelay controls whether the operating system should delay
//line /snap/go/10455/src/net/tcpsock.go:208
// packet transmission in hopes of sending fewer packets (Nagle's
//line /snap/go/10455/src/net/tcpsock.go:208
// algorithm).  The default is true (no delay), meaning that data is
//line /snap/go/10455/src/net/tcpsock.go:208
// sent as soon as possible after a Write.
//line /snap/go/10455/src/net/tcpsock.go:212
func (c *TCPConn) SetNoDelay(noDelay bool) error {
//line /snap/go/10455/src/net/tcpsock.go:212
	_go_fuzz_dep_.CoverTab[8372]++
						if !c.ok() {
//line /snap/go/10455/src/net/tcpsock.go:213
		_go_fuzz_dep_.CoverTab[529848]++
//line /snap/go/10455/src/net/tcpsock.go:213
		_go_fuzz_dep_.CoverTab[8375]++
							return syscall.EINVAL
//line /snap/go/10455/src/net/tcpsock.go:214
		// _ = "end of CoverTab[8375]"
	} else {
//line /snap/go/10455/src/net/tcpsock.go:215
		_go_fuzz_dep_.CoverTab[529849]++
//line /snap/go/10455/src/net/tcpsock.go:215
		_go_fuzz_dep_.CoverTab[8376]++
//line /snap/go/10455/src/net/tcpsock.go:215
		// _ = "end of CoverTab[8376]"
//line /snap/go/10455/src/net/tcpsock.go:215
	}
//line /snap/go/10455/src/net/tcpsock.go:215
	// _ = "end of CoverTab[8372]"
//line /snap/go/10455/src/net/tcpsock.go:215
	_go_fuzz_dep_.CoverTab[8373]++
						if err := setNoDelay(c.fd, noDelay); err != nil {
//line /snap/go/10455/src/net/tcpsock.go:216
		_go_fuzz_dep_.CoverTab[529850]++
//line /snap/go/10455/src/net/tcpsock.go:216
		_go_fuzz_dep_.CoverTab[8377]++
							return &OpError{Op: "set", Net: c.fd.net, Source: c.fd.laddr, Addr: c.fd.raddr, Err: err}
//line /snap/go/10455/src/net/tcpsock.go:217
		// _ = "end of CoverTab[8377]"
	} else {
//line /snap/go/10455/src/net/tcpsock.go:218
		_go_fuzz_dep_.CoverTab[529851]++
//line /snap/go/10455/src/net/tcpsock.go:218
		_go_fuzz_dep_.CoverTab[8378]++
//line /snap/go/10455/src/net/tcpsock.go:218
		// _ = "end of CoverTab[8378]"
//line /snap/go/10455/src/net/tcpsock.go:218
	}
//line /snap/go/10455/src/net/tcpsock.go:218
	// _ = "end of CoverTab[8373]"
//line /snap/go/10455/src/net/tcpsock.go:218
	_go_fuzz_dep_.CoverTab[8374]++
						return nil
//line /snap/go/10455/src/net/tcpsock.go:219
	// _ = "end of CoverTab[8374]"
}

// MultipathTCP reports whether the ongoing connection is using MPTCP.
//line /snap/go/10455/src/net/tcpsock.go:222
//
//line /snap/go/10455/src/net/tcpsock.go:222
// If Multipath TCP is not supported by the host, by the other peer or
//line /snap/go/10455/src/net/tcpsock.go:222
// intentionally / accidentally filtered out by a device in between, a
//line /snap/go/10455/src/net/tcpsock.go:222
// fallback to TCP will be done. This method does its best to check if
//line /snap/go/10455/src/net/tcpsock.go:222
// MPTCP is still being used or not.
//line /snap/go/10455/src/net/tcpsock.go:222
//
//line /snap/go/10455/src/net/tcpsock.go:222
// On Linux, more conditions are verified on kernels >= v5.16, improving
//line /snap/go/10455/src/net/tcpsock.go:222
// the results.
//line /snap/go/10455/src/net/tcpsock.go:231
func (c *TCPConn) MultipathTCP() (bool, error) {
//line /snap/go/10455/src/net/tcpsock.go:231
	_go_fuzz_dep_.CoverTab[8379]++
						if !c.ok() {
//line /snap/go/10455/src/net/tcpsock.go:232
		_go_fuzz_dep_.CoverTab[529852]++
//line /snap/go/10455/src/net/tcpsock.go:232
		_go_fuzz_dep_.CoverTab[8381]++
							return false, syscall.EINVAL
//line /snap/go/10455/src/net/tcpsock.go:233
		// _ = "end of CoverTab[8381]"
	} else {
//line /snap/go/10455/src/net/tcpsock.go:234
		_go_fuzz_dep_.CoverTab[529853]++
//line /snap/go/10455/src/net/tcpsock.go:234
		_go_fuzz_dep_.CoverTab[8382]++
//line /snap/go/10455/src/net/tcpsock.go:234
		// _ = "end of CoverTab[8382]"
//line /snap/go/10455/src/net/tcpsock.go:234
	}
//line /snap/go/10455/src/net/tcpsock.go:234
	// _ = "end of CoverTab[8379]"
//line /snap/go/10455/src/net/tcpsock.go:234
	_go_fuzz_dep_.CoverTab[8380]++
						return isUsingMultipathTCP(c.fd), nil
//line /snap/go/10455/src/net/tcpsock.go:235
	// _ = "end of CoverTab[8380]"
}

func newTCPConn(fd *netFD, keepAlive time.Duration, keepAliveHook func(time.Duration)) *TCPConn {
//line /snap/go/10455/src/net/tcpsock.go:238
	_go_fuzz_dep_.CoverTab[8383]++
						setNoDelay(fd, true)
						if keepAlive == 0 {
//line /snap/go/10455/src/net/tcpsock.go:240
		_go_fuzz_dep_.CoverTab[529854]++
//line /snap/go/10455/src/net/tcpsock.go:240
		_go_fuzz_dep_.CoverTab[8386]++
							keepAlive = defaultTCPKeepAlive
//line /snap/go/10455/src/net/tcpsock.go:241
		// _ = "end of CoverTab[8386]"
	} else {
//line /snap/go/10455/src/net/tcpsock.go:242
		_go_fuzz_dep_.CoverTab[529855]++
//line /snap/go/10455/src/net/tcpsock.go:242
		_go_fuzz_dep_.CoverTab[8387]++
//line /snap/go/10455/src/net/tcpsock.go:242
		// _ = "end of CoverTab[8387]"
//line /snap/go/10455/src/net/tcpsock.go:242
	}
//line /snap/go/10455/src/net/tcpsock.go:242
	// _ = "end of CoverTab[8383]"
//line /snap/go/10455/src/net/tcpsock.go:242
	_go_fuzz_dep_.CoverTab[8384]++
						if keepAlive > 0 {
//line /snap/go/10455/src/net/tcpsock.go:243
		_go_fuzz_dep_.CoverTab[529856]++
//line /snap/go/10455/src/net/tcpsock.go:243
		_go_fuzz_dep_.CoverTab[8388]++
							setKeepAlive(fd, true)
							setKeepAlivePeriod(fd, keepAlive)
							if keepAliveHook != nil {
//line /snap/go/10455/src/net/tcpsock.go:246
			_go_fuzz_dep_.CoverTab[529858]++
//line /snap/go/10455/src/net/tcpsock.go:246
			_go_fuzz_dep_.CoverTab[8389]++
								keepAliveHook(keepAlive)
//line /snap/go/10455/src/net/tcpsock.go:247
			// _ = "end of CoverTab[8389]"
		} else {
//line /snap/go/10455/src/net/tcpsock.go:248
			_go_fuzz_dep_.CoverTab[529859]++
//line /snap/go/10455/src/net/tcpsock.go:248
			_go_fuzz_dep_.CoverTab[8390]++
//line /snap/go/10455/src/net/tcpsock.go:248
			// _ = "end of CoverTab[8390]"
//line /snap/go/10455/src/net/tcpsock.go:248
		}
//line /snap/go/10455/src/net/tcpsock.go:248
		// _ = "end of CoverTab[8388]"
	} else {
//line /snap/go/10455/src/net/tcpsock.go:249
		_go_fuzz_dep_.CoverTab[529857]++
//line /snap/go/10455/src/net/tcpsock.go:249
		_go_fuzz_dep_.CoverTab[8391]++
//line /snap/go/10455/src/net/tcpsock.go:249
		// _ = "end of CoverTab[8391]"
//line /snap/go/10455/src/net/tcpsock.go:249
	}
//line /snap/go/10455/src/net/tcpsock.go:249
	// _ = "end of CoverTab[8384]"
//line /snap/go/10455/src/net/tcpsock.go:249
	_go_fuzz_dep_.CoverTab[8385]++
						return &TCPConn{conn{fd}}
//line /snap/go/10455/src/net/tcpsock.go:250
	// _ = "end of CoverTab[8385]"
}

// DialTCP acts like Dial for TCP networks.
//line /snap/go/10455/src/net/tcpsock.go:253
//
//line /snap/go/10455/src/net/tcpsock.go:253
// The network must be a TCP network name; see func Dial for details.
//line /snap/go/10455/src/net/tcpsock.go:253
//
//line /snap/go/10455/src/net/tcpsock.go:253
// If laddr is nil, a local address is automatically chosen.
//line /snap/go/10455/src/net/tcpsock.go:253
// If the IP field of raddr is nil or an unspecified IP address, the
//line /snap/go/10455/src/net/tcpsock.go:253
// local system is assumed.
//line /snap/go/10455/src/net/tcpsock.go:260
func DialTCP(network string, laddr, raddr *TCPAddr) (*TCPConn, error) {
//line /snap/go/10455/src/net/tcpsock.go:260
	_go_fuzz_dep_.CoverTab[8392]++
						switch network {
	case "tcp", "tcp4", "tcp6":
//line /snap/go/10455/src/net/tcpsock.go:262
		_go_fuzz_dep_.CoverTab[529860]++
//line /snap/go/10455/src/net/tcpsock.go:262
		_go_fuzz_dep_.CoverTab[8396]++
//line /snap/go/10455/src/net/tcpsock.go:262
		// _ = "end of CoverTab[8396]"
	default:
//line /snap/go/10455/src/net/tcpsock.go:263
		_go_fuzz_dep_.CoverTab[529861]++
//line /snap/go/10455/src/net/tcpsock.go:263
		_go_fuzz_dep_.CoverTab[8397]++
							return nil, &OpError{Op: "dial", Net: network, Source: laddr.opAddr(), Addr: raddr.opAddr(), Err: UnknownNetworkError(network)}
//line /snap/go/10455/src/net/tcpsock.go:264
		// _ = "end of CoverTab[8397]"
	}
//line /snap/go/10455/src/net/tcpsock.go:265
	// _ = "end of CoverTab[8392]"
//line /snap/go/10455/src/net/tcpsock.go:265
	_go_fuzz_dep_.CoverTab[8393]++
						if raddr == nil {
//line /snap/go/10455/src/net/tcpsock.go:266
		_go_fuzz_dep_.CoverTab[529862]++
//line /snap/go/10455/src/net/tcpsock.go:266
		_go_fuzz_dep_.CoverTab[8398]++
							return nil, &OpError{Op: "dial", Net: network, Source: laddr.opAddr(), Addr: nil, Err: errMissingAddress}
//line /snap/go/10455/src/net/tcpsock.go:267
		// _ = "end of CoverTab[8398]"
	} else {
//line /snap/go/10455/src/net/tcpsock.go:268
		_go_fuzz_dep_.CoverTab[529863]++
//line /snap/go/10455/src/net/tcpsock.go:268
		_go_fuzz_dep_.CoverTab[8399]++
//line /snap/go/10455/src/net/tcpsock.go:268
		// _ = "end of CoverTab[8399]"
//line /snap/go/10455/src/net/tcpsock.go:268
	}
//line /snap/go/10455/src/net/tcpsock.go:268
	// _ = "end of CoverTab[8393]"
//line /snap/go/10455/src/net/tcpsock.go:268
	_go_fuzz_dep_.CoverTab[8394]++
						sd := &sysDialer{network: network, address: raddr.String()}
						c, err := sd.dialTCP(context.Background(), laddr, raddr)
						if err != nil {
//line /snap/go/10455/src/net/tcpsock.go:271
		_go_fuzz_dep_.CoverTab[529864]++
//line /snap/go/10455/src/net/tcpsock.go:271
		_go_fuzz_dep_.CoverTab[8400]++
							return nil, &OpError{Op: "dial", Net: network, Source: laddr.opAddr(), Addr: raddr.opAddr(), Err: err}
//line /snap/go/10455/src/net/tcpsock.go:272
		// _ = "end of CoverTab[8400]"
	} else {
//line /snap/go/10455/src/net/tcpsock.go:273
		_go_fuzz_dep_.CoverTab[529865]++
//line /snap/go/10455/src/net/tcpsock.go:273
		_go_fuzz_dep_.CoverTab[8401]++
//line /snap/go/10455/src/net/tcpsock.go:273
		// _ = "end of CoverTab[8401]"
//line /snap/go/10455/src/net/tcpsock.go:273
	}
//line /snap/go/10455/src/net/tcpsock.go:273
	// _ = "end of CoverTab[8394]"
//line /snap/go/10455/src/net/tcpsock.go:273
	_go_fuzz_dep_.CoverTab[8395]++
						return c, nil
//line /snap/go/10455/src/net/tcpsock.go:274
	// _ = "end of CoverTab[8395]"
}

// TCPListener is a TCP network listener. Clients should typically
//line /snap/go/10455/src/net/tcpsock.go:277
// use variables of type Listener instead of assuming TCP.
//line /snap/go/10455/src/net/tcpsock.go:279
type TCPListener struct {
	fd	*netFD
	lc	ListenConfig
}

// SyscallConn returns a raw network connection.
//line /snap/go/10455/src/net/tcpsock.go:284
// This implements the syscall.Conn interface.
//line /snap/go/10455/src/net/tcpsock.go:284
//
//line /snap/go/10455/src/net/tcpsock.go:284
// The returned RawConn only supports calling Control. Read and
//line /snap/go/10455/src/net/tcpsock.go:284
// Write return an error.
//line /snap/go/10455/src/net/tcpsock.go:289
func (l *TCPListener) SyscallConn() (syscall.RawConn, error) {
//line /snap/go/10455/src/net/tcpsock.go:289
	_go_fuzz_dep_.CoverTab[8402]++
						if !l.ok() {
//line /snap/go/10455/src/net/tcpsock.go:290
		_go_fuzz_dep_.CoverTab[529866]++
//line /snap/go/10455/src/net/tcpsock.go:290
		_go_fuzz_dep_.CoverTab[8404]++
							return nil, syscall.EINVAL
//line /snap/go/10455/src/net/tcpsock.go:291
		// _ = "end of CoverTab[8404]"
	} else {
//line /snap/go/10455/src/net/tcpsock.go:292
		_go_fuzz_dep_.CoverTab[529867]++
//line /snap/go/10455/src/net/tcpsock.go:292
		_go_fuzz_dep_.CoverTab[8405]++
//line /snap/go/10455/src/net/tcpsock.go:292
		// _ = "end of CoverTab[8405]"
//line /snap/go/10455/src/net/tcpsock.go:292
	}
//line /snap/go/10455/src/net/tcpsock.go:292
	// _ = "end of CoverTab[8402]"
//line /snap/go/10455/src/net/tcpsock.go:292
	_go_fuzz_dep_.CoverTab[8403]++
						return newRawListener(l.fd)
//line /snap/go/10455/src/net/tcpsock.go:293
	// _ = "end of CoverTab[8403]"
}

// AcceptTCP accepts the next incoming call and returns the new
//line /snap/go/10455/src/net/tcpsock.go:296
// connection.
//line /snap/go/10455/src/net/tcpsock.go:298
func (l *TCPListener) AcceptTCP() (*TCPConn, error) {
//line /snap/go/10455/src/net/tcpsock.go:298
	_go_fuzz_dep_.CoverTab[8406]++
						if !l.ok() {
//line /snap/go/10455/src/net/tcpsock.go:299
		_go_fuzz_dep_.CoverTab[529868]++
//line /snap/go/10455/src/net/tcpsock.go:299
		_go_fuzz_dep_.CoverTab[8409]++
							return nil, syscall.EINVAL
//line /snap/go/10455/src/net/tcpsock.go:300
		// _ = "end of CoverTab[8409]"
	} else {
//line /snap/go/10455/src/net/tcpsock.go:301
		_go_fuzz_dep_.CoverTab[529869]++
//line /snap/go/10455/src/net/tcpsock.go:301
		_go_fuzz_dep_.CoverTab[8410]++
//line /snap/go/10455/src/net/tcpsock.go:301
		// _ = "end of CoverTab[8410]"
//line /snap/go/10455/src/net/tcpsock.go:301
	}
//line /snap/go/10455/src/net/tcpsock.go:301
	// _ = "end of CoverTab[8406]"
//line /snap/go/10455/src/net/tcpsock.go:301
	_go_fuzz_dep_.CoverTab[8407]++
						c, err := l.accept()
						if err != nil {
//line /snap/go/10455/src/net/tcpsock.go:303
		_go_fuzz_dep_.CoverTab[529870]++
//line /snap/go/10455/src/net/tcpsock.go:303
		_go_fuzz_dep_.CoverTab[8411]++
							return nil, &OpError{Op: "accept", Net: l.fd.net, Source: nil, Addr: l.fd.laddr, Err: err}
//line /snap/go/10455/src/net/tcpsock.go:304
		// _ = "end of CoverTab[8411]"
	} else {
//line /snap/go/10455/src/net/tcpsock.go:305
		_go_fuzz_dep_.CoverTab[529871]++
//line /snap/go/10455/src/net/tcpsock.go:305
		_go_fuzz_dep_.CoverTab[8412]++
//line /snap/go/10455/src/net/tcpsock.go:305
		// _ = "end of CoverTab[8412]"
//line /snap/go/10455/src/net/tcpsock.go:305
	}
//line /snap/go/10455/src/net/tcpsock.go:305
	// _ = "end of CoverTab[8407]"
//line /snap/go/10455/src/net/tcpsock.go:305
	_go_fuzz_dep_.CoverTab[8408]++
						return c, nil
//line /snap/go/10455/src/net/tcpsock.go:306
	// _ = "end of CoverTab[8408]"
}

// Accept implements the Accept method in the Listener interface; it
//line /snap/go/10455/src/net/tcpsock.go:309
// waits for the next call and returns a generic Conn.
//line /snap/go/10455/src/net/tcpsock.go:311
func (l *TCPListener) Accept() (Conn, error) {
//line /snap/go/10455/src/net/tcpsock.go:311
	_go_fuzz_dep_.CoverTab[8413]++
						if !l.ok() {
//line /snap/go/10455/src/net/tcpsock.go:312
		_go_fuzz_dep_.CoverTab[529872]++
//line /snap/go/10455/src/net/tcpsock.go:312
		_go_fuzz_dep_.CoverTab[8416]++
							return nil, syscall.EINVAL
//line /snap/go/10455/src/net/tcpsock.go:313
		// _ = "end of CoverTab[8416]"
	} else {
//line /snap/go/10455/src/net/tcpsock.go:314
		_go_fuzz_dep_.CoverTab[529873]++
//line /snap/go/10455/src/net/tcpsock.go:314
		_go_fuzz_dep_.CoverTab[8417]++
//line /snap/go/10455/src/net/tcpsock.go:314
		// _ = "end of CoverTab[8417]"
//line /snap/go/10455/src/net/tcpsock.go:314
	}
//line /snap/go/10455/src/net/tcpsock.go:314
	// _ = "end of CoverTab[8413]"
//line /snap/go/10455/src/net/tcpsock.go:314
	_go_fuzz_dep_.CoverTab[8414]++
						c, err := l.accept()
						if err != nil {
//line /snap/go/10455/src/net/tcpsock.go:316
		_go_fuzz_dep_.CoverTab[529874]++
//line /snap/go/10455/src/net/tcpsock.go:316
		_go_fuzz_dep_.CoverTab[8418]++
							return nil, &OpError{Op: "accept", Net: l.fd.net, Source: nil, Addr: l.fd.laddr, Err: err}
//line /snap/go/10455/src/net/tcpsock.go:317
		// _ = "end of CoverTab[8418]"
	} else {
//line /snap/go/10455/src/net/tcpsock.go:318
		_go_fuzz_dep_.CoverTab[529875]++
//line /snap/go/10455/src/net/tcpsock.go:318
		_go_fuzz_dep_.CoverTab[8419]++
//line /snap/go/10455/src/net/tcpsock.go:318
		// _ = "end of CoverTab[8419]"
//line /snap/go/10455/src/net/tcpsock.go:318
	}
//line /snap/go/10455/src/net/tcpsock.go:318
	// _ = "end of CoverTab[8414]"
//line /snap/go/10455/src/net/tcpsock.go:318
	_go_fuzz_dep_.CoverTab[8415]++
						return c, nil
//line /snap/go/10455/src/net/tcpsock.go:319
	// _ = "end of CoverTab[8415]"
}

// Close stops listening on the TCP address.
//line /snap/go/10455/src/net/tcpsock.go:322
// Already Accepted connections are not closed.
//line /snap/go/10455/src/net/tcpsock.go:324
func (l *TCPListener) Close() error {
//line /snap/go/10455/src/net/tcpsock.go:324
	_go_fuzz_dep_.CoverTab[8420]++
						if !l.ok() {
//line /snap/go/10455/src/net/tcpsock.go:325
		_go_fuzz_dep_.CoverTab[529876]++
//line /snap/go/10455/src/net/tcpsock.go:325
		_go_fuzz_dep_.CoverTab[8423]++
							return syscall.EINVAL
//line /snap/go/10455/src/net/tcpsock.go:326
		// _ = "end of CoverTab[8423]"
	} else {
//line /snap/go/10455/src/net/tcpsock.go:327
		_go_fuzz_dep_.CoverTab[529877]++
//line /snap/go/10455/src/net/tcpsock.go:327
		_go_fuzz_dep_.CoverTab[8424]++
//line /snap/go/10455/src/net/tcpsock.go:327
		// _ = "end of CoverTab[8424]"
//line /snap/go/10455/src/net/tcpsock.go:327
	}
//line /snap/go/10455/src/net/tcpsock.go:327
	// _ = "end of CoverTab[8420]"
//line /snap/go/10455/src/net/tcpsock.go:327
	_go_fuzz_dep_.CoverTab[8421]++
						if err := l.close(); err != nil {
//line /snap/go/10455/src/net/tcpsock.go:328
		_go_fuzz_dep_.CoverTab[529878]++
//line /snap/go/10455/src/net/tcpsock.go:328
		_go_fuzz_dep_.CoverTab[8425]++
							return &OpError{Op: "close", Net: l.fd.net, Source: nil, Addr: l.fd.laddr, Err: err}
//line /snap/go/10455/src/net/tcpsock.go:329
		// _ = "end of CoverTab[8425]"
	} else {
//line /snap/go/10455/src/net/tcpsock.go:330
		_go_fuzz_dep_.CoverTab[529879]++
//line /snap/go/10455/src/net/tcpsock.go:330
		_go_fuzz_dep_.CoverTab[8426]++
//line /snap/go/10455/src/net/tcpsock.go:330
		// _ = "end of CoverTab[8426]"
//line /snap/go/10455/src/net/tcpsock.go:330
	}
//line /snap/go/10455/src/net/tcpsock.go:330
	// _ = "end of CoverTab[8421]"
//line /snap/go/10455/src/net/tcpsock.go:330
	_go_fuzz_dep_.CoverTab[8422]++
						return nil
//line /snap/go/10455/src/net/tcpsock.go:331
	// _ = "end of CoverTab[8422]"
}

// Addr returns the listener's network address, a *TCPAddr.
//line /snap/go/10455/src/net/tcpsock.go:334
// The Addr returned is shared by all invocations of Addr, so
//line /snap/go/10455/src/net/tcpsock.go:334
// do not modify it.
//line /snap/go/10455/src/net/tcpsock.go:337
func (l *TCPListener) Addr() Addr {
//line /snap/go/10455/src/net/tcpsock.go:337
	_go_fuzz_dep_.CoverTab[8427]++
//line /snap/go/10455/src/net/tcpsock.go:337
	return l.fd.laddr
//line /snap/go/10455/src/net/tcpsock.go:337
	// _ = "end of CoverTab[8427]"
//line /snap/go/10455/src/net/tcpsock.go:337
}

// SetDeadline sets the deadline associated with the listener.
//line /snap/go/10455/src/net/tcpsock.go:339
// A zero time value disables the deadline.
//line /snap/go/10455/src/net/tcpsock.go:341
func (l *TCPListener) SetDeadline(t time.Time) error {
//line /snap/go/10455/src/net/tcpsock.go:341
	_go_fuzz_dep_.CoverTab[8428]++
						if !l.ok() {
//line /snap/go/10455/src/net/tcpsock.go:342
		_go_fuzz_dep_.CoverTab[529880]++
//line /snap/go/10455/src/net/tcpsock.go:342
		_go_fuzz_dep_.CoverTab[8431]++
							return syscall.EINVAL
//line /snap/go/10455/src/net/tcpsock.go:343
		// _ = "end of CoverTab[8431]"
	} else {
//line /snap/go/10455/src/net/tcpsock.go:344
		_go_fuzz_dep_.CoverTab[529881]++
//line /snap/go/10455/src/net/tcpsock.go:344
		_go_fuzz_dep_.CoverTab[8432]++
//line /snap/go/10455/src/net/tcpsock.go:344
		// _ = "end of CoverTab[8432]"
//line /snap/go/10455/src/net/tcpsock.go:344
	}
//line /snap/go/10455/src/net/tcpsock.go:344
	// _ = "end of CoverTab[8428]"
//line /snap/go/10455/src/net/tcpsock.go:344
	_go_fuzz_dep_.CoverTab[8429]++
						if err := l.fd.pfd.SetDeadline(t); err != nil {
//line /snap/go/10455/src/net/tcpsock.go:345
		_go_fuzz_dep_.CoverTab[529882]++
//line /snap/go/10455/src/net/tcpsock.go:345
		_go_fuzz_dep_.CoverTab[8433]++
							return &OpError{Op: "set", Net: l.fd.net, Source: nil, Addr: l.fd.laddr, Err: err}
//line /snap/go/10455/src/net/tcpsock.go:346
		// _ = "end of CoverTab[8433]"
	} else {
//line /snap/go/10455/src/net/tcpsock.go:347
		_go_fuzz_dep_.CoverTab[529883]++
//line /snap/go/10455/src/net/tcpsock.go:347
		_go_fuzz_dep_.CoverTab[8434]++
//line /snap/go/10455/src/net/tcpsock.go:347
		// _ = "end of CoverTab[8434]"
//line /snap/go/10455/src/net/tcpsock.go:347
	}
//line /snap/go/10455/src/net/tcpsock.go:347
	// _ = "end of CoverTab[8429]"
//line /snap/go/10455/src/net/tcpsock.go:347
	_go_fuzz_dep_.CoverTab[8430]++
						return nil
//line /snap/go/10455/src/net/tcpsock.go:348
	// _ = "end of CoverTab[8430]"
}

// File returns a copy of the underlying os.File.
//line /snap/go/10455/src/net/tcpsock.go:351
// It is the caller's responsibility to close f when finished.
//line /snap/go/10455/src/net/tcpsock.go:351
// Closing l does not affect f, and closing f does not affect l.
//line /snap/go/10455/src/net/tcpsock.go:351
//
//line /snap/go/10455/src/net/tcpsock.go:351
// The returned os.File's file descriptor is different from the
//line /snap/go/10455/src/net/tcpsock.go:351
// connection's. Attempting to change properties of the original
//line /snap/go/10455/src/net/tcpsock.go:351
// using this duplicate may or may not have the desired effect.
//line /snap/go/10455/src/net/tcpsock.go:358
func (l *TCPListener) File() (f *os.File, err error) {
//line /snap/go/10455/src/net/tcpsock.go:358
	_go_fuzz_dep_.CoverTab[8435]++
						if !l.ok() {
//line /snap/go/10455/src/net/tcpsock.go:359
		_go_fuzz_dep_.CoverTab[529884]++
//line /snap/go/10455/src/net/tcpsock.go:359
		_go_fuzz_dep_.CoverTab[8438]++
							return nil, syscall.EINVAL
//line /snap/go/10455/src/net/tcpsock.go:360
		// _ = "end of CoverTab[8438]"
	} else {
//line /snap/go/10455/src/net/tcpsock.go:361
		_go_fuzz_dep_.CoverTab[529885]++
//line /snap/go/10455/src/net/tcpsock.go:361
		_go_fuzz_dep_.CoverTab[8439]++
//line /snap/go/10455/src/net/tcpsock.go:361
		// _ = "end of CoverTab[8439]"
//line /snap/go/10455/src/net/tcpsock.go:361
	}
//line /snap/go/10455/src/net/tcpsock.go:361
	// _ = "end of CoverTab[8435]"
//line /snap/go/10455/src/net/tcpsock.go:361
	_go_fuzz_dep_.CoverTab[8436]++
						f, err = l.file()
						if err != nil {
//line /snap/go/10455/src/net/tcpsock.go:363
		_go_fuzz_dep_.CoverTab[529886]++
//line /snap/go/10455/src/net/tcpsock.go:363
		_go_fuzz_dep_.CoverTab[8440]++
							return nil, &OpError{Op: "file", Net: l.fd.net, Source: nil, Addr: l.fd.laddr, Err: err}
//line /snap/go/10455/src/net/tcpsock.go:364
		// _ = "end of CoverTab[8440]"
	} else {
//line /snap/go/10455/src/net/tcpsock.go:365
		_go_fuzz_dep_.CoverTab[529887]++
//line /snap/go/10455/src/net/tcpsock.go:365
		_go_fuzz_dep_.CoverTab[8441]++
//line /snap/go/10455/src/net/tcpsock.go:365
		// _ = "end of CoverTab[8441]"
//line /snap/go/10455/src/net/tcpsock.go:365
	}
//line /snap/go/10455/src/net/tcpsock.go:365
	// _ = "end of CoverTab[8436]"
//line /snap/go/10455/src/net/tcpsock.go:365
	_go_fuzz_dep_.CoverTab[8437]++
						return
//line /snap/go/10455/src/net/tcpsock.go:366
	// _ = "end of CoverTab[8437]"
}

// ListenTCP acts like Listen for TCP networks.
//line /snap/go/10455/src/net/tcpsock.go:369
//
//line /snap/go/10455/src/net/tcpsock.go:369
// The network must be a TCP network name; see func Dial for details.
//line /snap/go/10455/src/net/tcpsock.go:369
//
//line /snap/go/10455/src/net/tcpsock.go:369
// If the IP field of laddr is nil or an unspecified IP address,
//line /snap/go/10455/src/net/tcpsock.go:369
// ListenTCP listens on all available unicast and anycast IP addresses
//line /snap/go/10455/src/net/tcpsock.go:369
// of the local system.
//line /snap/go/10455/src/net/tcpsock.go:369
// If the Port field of laddr is 0, a port number is automatically
//line /snap/go/10455/src/net/tcpsock.go:369
// chosen.
//line /snap/go/10455/src/net/tcpsock.go:378
func ListenTCP(network string, laddr *TCPAddr) (*TCPListener, error) {
//line /snap/go/10455/src/net/tcpsock.go:378
	_go_fuzz_dep_.CoverTab[8442]++
						switch network {
	case "tcp", "tcp4", "tcp6":
//line /snap/go/10455/src/net/tcpsock.go:380
		_go_fuzz_dep_.CoverTab[529888]++
//line /snap/go/10455/src/net/tcpsock.go:380
		_go_fuzz_dep_.CoverTab[8446]++
//line /snap/go/10455/src/net/tcpsock.go:380
		// _ = "end of CoverTab[8446]"
	default:
//line /snap/go/10455/src/net/tcpsock.go:381
		_go_fuzz_dep_.CoverTab[529889]++
//line /snap/go/10455/src/net/tcpsock.go:381
		_go_fuzz_dep_.CoverTab[8447]++
							return nil, &OpError{Op: "listen", Net: network, Source: nil, Addr: laddr.opAddr(), Err: UnknownNetworkError(network)}
//line /snap/go/10455/src/net/tcpsock.go:382
		// _ = "end of CoverTab[8447]"
	}
//line /snap/go/10455/src/net/tcpsock.go:383
	// _ = "end of CoverTab[8442]"
//line /snap/go/10455/src/net/tcpsock.go:383
	_go_fuzz_dep_.CoverTab[8443]++
						if laddr == nil {
//line /snap/go/10455/src/net/tcpsock.go:384
		_go_fuzz_dep_.CoverTab[529890]++
//line /snap/go/10455/src/net/tcpsock.go:384
		_go_fuzz_dep_.CoverTab[8448]++
							laddr = &TCPAddr{}
//line /snap/go/10455/src/net/tcpsock.go:385
		// _ = "end of CoverTab[8448]"
	} else {
//line /snap/go/10455/src/net/tcpsock.go:386
		_go_fuzz_dep_.CoverTab[529891]++
//line /snap/go/10455/src/net/tcpsock.go:386
		_go_fuzz_dep_.CoverTab[8449]++
//line /snap/go/10455/src/net/tcpsock.go:386
		// _ = "end of CoverTab[8449]"
//line /snap/go/10455/src/net/tcpsock.go:386
	}
//line /snap/go/10455/src/net/tcpsock.go:386
	// _ = "end of CoverTab[8443]"
//line /snap/go/10455/src/net/tcpsock.go:386
	_go_fuzz_dep_.CoverTab[8444]++
						sl := &sysListener{network: network, address: laddr.String()}
						ln, err := sl.listenTCP(context.Background(), laddr)
						if err != nil {
//line /snap/go/10455/src/net/tcpsock.go:389
		_go_fuzz_dep_.CoverTab[529892]++
//line /snap/go/10455/src/net/tcpsock.go:389
		_go_fuzz_dep_.CoverTab[8450]++
							return nil, &OpError{Op: "listen", Net: network, Source: nil, Addr: laddr.opAddr(), Err: err}
//line /snap/go/10455/src/net/tcpsock.go:390
		// _ = "end of CoverTab[8450]"
	} else {
//line /snap/go/10455/src/net/tcpsock.go:391
		_go_fuzz_dep_.CoverTab[529893]++
//line /snap/go/10455/src/net/tcpsock.go:391
		_go_fuzz_dep_.CoverTab[8451]++
//line /snap/go/10455/src/net/tcpsock.go:391
		// _ = "end of CoverTab[8451]"
//line /snap/go/10455/src/net/tcpsock.go:391
	}
//line /snap/go/10455/src/net/tcpsock.go:391
	// _ = "end of CoverTab[8444]"
//line /snap/go/10455/src/net/tcpsock.go:391
	_go_fuzz_dep_.CoverTab[8445]++
						return ln, nil
//line /snap/go/10455/src/net/tcpsock.go:392
	// _ = "end of CoverTab[8445]"
}

// roundDurationUp rounds d to the next multiple of to.
func roundDurationUp(d time.Duration, to time.Duration) time.Duration {
//line /snap/go/10455/src/net/tcpsock.go:396
	_go_fuzz_dep_.CoverTab[8452]++
						return (d + to - 1) / to
//line /snap/go/10455/src/net/tcpsock.go:397
	// _ = "end of CoverTab[8452]"
}

//line /snap/go/10455/src/net/tcpsock.go:398
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /snap/go/10455/src/net/tcpsock.go:398
var _ = _go_fuzz_dep_.CoverTab
