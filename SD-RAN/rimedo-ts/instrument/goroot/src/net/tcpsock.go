// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/net/tcpsock.go:5
package net

//line /usr/local/go/src/net/tcpsock.go:5
import (
//line /usr/local/go/src/net/tcpsock.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/tcpsock.go:5
)
//line /usr/local/go/src/net/tcpsock.go:5
import (
//line /usr/local/go/src/net/tcpsock.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/tcpsock.go:5
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

//line /usr/local/go/src/net/tcpsock.go:20
// TCPAddr represents the address of a TCP end point.
type TCPAddr struct {
	IP	IP
	Port	int
	Zone	string	// IPv6 scoped addressing zone
}

// AddrPort returns the TCPAddr a as a netip.AddrPort.
//line /usr/local/go/src/net/tcpsock.go:27
//
//line /usr/local/go/src/net/tcpsock.go:27
// If a.Port does not fit in a uint16, it's silently truncated.
//line /usr/local/go/src/net/tcpsock.go:27
//
//line /usr/local/go/src/net/tcpsock.go:27
// If a is nil, a zero value is returned.
//line /usr/local/go/src/net/tcpsock.go:32
func (a *TCPAddr) AddrPort() netip.AddrPort {
//line /usr/local/go/src/net/tcpsock.go:32
	_go_fuzz_dep_.CoverTab[16391]++
						if a == nil {
//line /usr/local/go/src/net/tcpsock.go:33
		_go_fuzz_dep_.CoverTab[16393]++
							return netip.AddrPort{}
//line /usr/local/go/src/net/tcpsock.go:34
		// _ = "end of CoverTab[16393]"
	} else {
//line /usr/local/go/src/net/tcpsock.go:35
		_go_fuzz_dep_.CoverTab[16394]++
//line /usr/local/go/src/net/tcpsock.go:35
		// _ = "end of CoverTab[16394]"
//line /usr/local/go/src/net/tcpsock.go:35
	}
//line /usr/local/go/src/net/tcpsock.go:35
	// _ = "end of CoverTab[16391]"
//line /usr/local/go/src/net/tcpsock.go:35
	_go_fuzz_dep_.CoverTab[16392]++
						na, _ := netip.AddrFromSlice(a.IP)
						na = na.WithZone(a.Zone)
						return netip.AddrPortFrom(na, uint16(a.Port))
//line /usr/local/go/src/net/tcpsock.go:38
	// _ = "end of CoverTab[16392]"
}

// Network returns the address's network name, "tcp".
func (a *TCPAddr) Network() string {
//line /usr/local/go/src/net/tcpsock.go:42
	_go_fuzz_dep_.CoverTab[16395]++
//line /usr/local/go/src/net/tcpsock.go:42
	return "tcp"
//line /usr/local/go/src/net/tcpsock.go:42
	// _ = "end of CoverTab[16395]"
//line /usr/local/go/src/net/tcpsock.go:42
}

func (a *TCPAddr) String() string {
//line /usr/local/go/src/net/tcpsock.go:44
	_go_fuzz_dep_.CoverTab[16396]++
						if a == nil {
//line /usr/local/go/src/net/tcpsock.go:45
		_go_fuzz_dep_.CoverTab[16399]++
							return "<nil>"
//line /usr/local/go/src/net/tcpsock.go:46
		// _ = "end of CoverTab[16399]"
	} else {
//line /usr/local/go/src/net/tcpsock.go:47
		_go_fuzz_dep_.CoverTab[16400]++
//line /usr/local/go/src/net/tcpsock.go:47
		// _ = "end of CoverTab[16400]"
//line /usr/local/go/src/net/tcpsock.go:47
	}
//line /usr/local/go/src/net/tcpsock.go:47
	// _ = "end of CoverTab[16396]"
//line /usr/local/go/src/net/tcpsock.go:47
	_go_fuzz_dep_.CoverTab[16397]++
						ip := ipEmptyString(a.IP)
						if a.Zone != "" {
//line /usr/local/go/src/net/tcpsock.go:49
		_go_fuzz_dep_.CoverTab[16401]++
							return JoinHostPort(ip+"%"+a.Zone, itoa.Itoa(a.Port))
//line /usr/local/go/src/net/tcpsock.go:50
		// _ = "end of CoverTab[16401]"
	} else {
//line /usr/local/go/src/net/tcpsock.go:51
		_go_fuzz_dep_.CoverTab[16402]++
//line /usr/local/go/src/net/tcpsock.go:51
		// _ = "end of CoverTab[16402]"
//line /usr/local/go/src/net/tcpsock.go:51
	}
//line /usr/local/go/src/net/tcpsock.go:51
	// _ = "end of CoverTab[16397]"
//line /usr/local/go/src/net/tcpsock.go:51
	_go_fuzz_dep_.CoverTab[16398]++
						return JoinHostPort(ip, itoa.Itoa(a.Port))
//line /usr/local/go/src/net/tcpsock.go:52
	// _ = "end of CoverTab[16398]"
}

func (a *TCPAddr) isWildcard() bool {
//line /usr/local/go/src/net/tcpsock.go:55
	_go_fuzz_dep_.CoverTab[16403]++
						if a == nil || func() bool {
//line /usr/local/go/src/net/tcpsock.go:56
		_go_fuzz_dep_.CoverTab[16405]++
//line /usr/local/go/src/net/tcpsock.go:56
		return a.IP == nil
//line /usr/local/go/src/net/tcpsock.go:56
		// _ = "end of CoverTab[16405]"
//line /usr/local/go/src/net/tcpsock.go:56
	}() {
//line /usr/local/go/src/net/tcpsock.go:56
		_go_fuzz_dep_.CoverTab[16406]++
							return true
//line /usr/local/go/src/net/tcpsock.go:57
		// _ = "end of CoverTab[16406]"
	} else {
//line /usr/local/go/src/net/tcpsock.go:58
		_go_fuzz_dep_.CoverTab[16407]++
//line /usr/local/go/src/net/tcpsock.go:58
		// _ = "end of CoverTab[16407]"
//line /usr/local/go/src/net/tcpsock.go:58
	}
//line /usr/local/go/src/net/tcpsock.go:58
	// _ = "end of CoverTab[16403]"
//line /usr/local/go/src/net/tcpsock.go:58
	_go_fuzz_dep_.CoverTab[16404]++
						return a.IP.IsUnspecified()
//line /usr/local/go/src/net/tcpsock.go:59
	// _ = "end of CoverTab[16404]"
}

func (a *TCPAddr) opAddr() Addr {
//line /usr/local/go/src/net/tcpsock.go:62
	_go_fuzz_dep_.CoverTab[16408]++
						if a == nil {
//line /usr/local/go/src/net/tcpsock.go:63
		_go_fuzz_dep_.CoverTab[16410]++
							return nil
//line /usr/local/go/src/net/tcpsock.go:64
		// _ = "end of CoverTab[16410]"
	} else {
//line /usr/local/go/src/net/tcpsock.go:65
		_go_fuzz_dep_.CoverTab[16411]++
//line /usr/local/go/src/net/tcpsock.go:65
		// _ = "end of CoverTab[16411]"
//line /usr/local/go/src/net/tcpsock.go:65
	}
//line /usr/local/go/src/net/tcpsock.go:65
	// _ = "end of CoverTab[16408]"
//line /usr/local/go/src/net/tcpsock.go:65
	_go_fuzz_dep_.CoverTab[16409]++
						return a
//line /usr/local/go/src/net/tcpsock.go:66
	// _ = "end of CoverTab[16409]"
}

// ResolveTCPAddr returns an address of TCP end point.
//line /usr/local/go/src/net/tcpsock.go:69
//
//line /usr/local/go/src/net/tcpsock.go:69
// The network must be a TCP network name.
//line /usr/local/go/src/net/tcpsock.go:69
//
//line /usr/local/go/src/net/tcpsock.go:69
// If the host in the address parameter is not a literal IP address or
//line /usr/local/go/src/net/tcpsock.go:69
// the port is not a literal port number, ResolveTCPAddr resolves the
//line /usr/local/go/src/net/tcpsock.go:69
// address to an address of TCP end point.
//line /usr/local/go/src/net/tcpsock.go:69
// Otherwise, it parses the address as a pair of literal IP address
//line /usr/local/go/src/net/tcpsock.go:69
// and port number.
//line /usr/local/go/src/net/tcpsock.go:69
// The address parameter can use a host name, but this is not
//line /usr/local/go/src/net/tcpsock.go:69
// recommended, because it will return at most one of the host name's
//line /usr/local/go/src/net/tcpsock.go:69
// IP addresses.
//line /usr/local/go/src/net/tcpsock.go:69
//
//line /usr/local/go/src/net/tcpsock.go:69
// See func Dial for a description of the network and address
//line /usr/local/go/src/net/tcpsock.go:69
// parameters.
//line /usr/local/go/src/net/tcpsock.go:84
func ResolveTCPAddr(network, address string) (*TCPAddr, error) {
//line /usr/local/go/src/net/tcpsock.go:84
	_go_fuzz_dep_.CoverTab[16412]++
						switch network {
	case "tcp", "tcp4", "tcp6":
//line /usr/local/go/src/net/tcpsock.go:86
		_go_fuzz_dep_.CoverTab[16415]++
//line /usr/local/go/src/net/tcpsock.go:86
		// _ = "end of CoverTab[16415]"
	case "":
//line /usr/local/go/src/net/tcpsock.go:87
		_go_fuzz_dep_.CoverTab[16416]++
							network = "tcp"
//line /usr/local/go/src/net/tcpsock.go:88
		// _ = "end of CoverTab[16416]"
	default:
//line /usr/local/go/src/net/tcpsock.go:89
		_go_fuzz_dep_.CoverTab[16417]++
							return nil, UnknownNetworkError(network)
//line /usr/local/go/src/net/tcpsock.go:90
		// _ = "end of CoverTab[16417]"
	}
//line /usr/local/go/src/net/tcpsock.go:91
	// _ = "end of CoverTab[16412]"
//line /usr/local/go/src/net/tcpsock.go:91
	_go_fuzz_dep_.CoverTab[16413]++
						addrs, err := DefaultResolver.internetAddrList(context.Background(), network, address)
						if err != nil {
//line /usr/local/go/src/net/tcpsock.go:93
		_go_fuzz_dep_.CoverTab[16418]++
							return nil, err
//line /usr/local/go/src/net/tcpsock.go:94
		// _ = "end of CoverTab[16418]"
	} else {
//line /usr/local/go/src/net/tcpsock.go:95
		_go_fuzz_dep_.CoverTab[16419]++
//line /usr/local/go/src/net/tcpsock.go:95
		// _ = "end of CoverTab[16419]"
//line /usr/local/go/src/net/tcpsock.go:95
	}
//line /usr/local/go/src/net/tcpsock.go:95
	// _ = "end of CoverTab[16413]"
//line /usr/local/go/src/net/tcpsock.go:95
	_go_fuzz_dep_.CoverTab[16414]++
						return addrs.forResolve(network, address).(*TCPAddr), nil
//line /usr/local/go/src/net/tcpsock.go:96
	// _ = "end of CoverTab[16414]"
}

// TCPAddrFromAddrPort returns addr as a TCPAddr. If addr.IsValid() is false,
//line /usr/local/go/src/net/tcpsock.go:99
// then the returned TCPAddr will contain a nil IP field, indicating an
//line /usr/local/go/src/net/tcpsock.go:99
// address family-agnostic unspecified address.
//line /usr/local/go/src/net/tcpsock.go:102
func TCPAddrFromAddrPort(addr netip.AddrPort) *TCPAddr {
//line /usr/local/go/src/net/tcpsock.go:102
	_go_fuzz_dep_.CoverTab[16420]++
						return &TCPAddr{
		IP:	addr.Addr().AsSlice(),
		Zone:	addr.Addr().Zone(),
		Port:	int(addr.Port()),
	}
//line /usr/local/go/src/net/tcpsock.go:107
	// _ = "end of CoverTab[16420]"
}

// TCPConn is an implementation of the Conn interface for TCP network
//line /usr/local/go/src/net/tcpsock.go:110
// connections.
//line /usr/local/go/src/net/tcpsock.go:112
type TCPConn struct {
	conn
}

// SyscallConn returns a raw network connection.
//line /usr/local/go/src/net/tcpsock.go:116
// This implements the syscall.Conn interface.
//line /usr/local/go/src/net/tcpsock.go:118
func (c *TCPConn) SyscallConn() (syscall.RawConn, error) {
//line /usr/local/go/src/net/tcpsock.go:118
	_go_fuzz_dep_.CoverTab[16421]++
						if !c.ok() {
//line /usr/local/go/src/net/tcpsock.go:119
		_go_fuzz_dep_.CoverTab[16423]++
							return nil, syscall.EINVAL
//line /usr/local/go/src/net/tcpsock.go:120
		// _ = "end of CoverTab[16423]"
	} else {
//line /usr/local/go/src/net/tcpsock.go:121
		_go_fuzz_dep_.CoverTab[16424]++
//line /usr/local/go/src/net/tcpsock.go:121
		// _ = "end of CoverTab[16424]"
//line /usr/local/go/src/net/tcpsock.go:121
	}
//line /usr/local/go/src/net/tcpsock.go:121
	// _ = "end of CoverTab[16421]"
//line /usr/local/go/src/net/tcpsock.go:121
	_go_fuzz_dep_.CoverTab[16422]++
						return newRawConn(c.fd)
//line /usr/local/go/src/net/tcpsock.go:122
	// _ = "end of CoverTab[16422]"
}

// ReadFrom implements the io.ReaderFrom ReadFrom method.
func (c *TCPConn) ReadFrom(r io.Reader) (int64, error) {
//line /usr/local/go/src/net/tcpsock.go:126
	_go_fuzz_dep_.CoverTab[16425]++
						if !c.ok() {
//line /usr/local/go/src/net/tcpsock.go:127
		_go_fuzz_dep_.CoverTab[16428]++
							return 0, syscall.EINVAL
//line /usr/local/go/src/net/tcpsock.go:128
		// _ = "end of CoverTab[16428]"
	} else {
//line /usr/local/go/src/net/tcpsock.go:129
		_go_fuzz_dep_.CoverTab[16429]++
//line /usr/local/go/src/net/tcpsock.go:129
		// _ = "end of CoverTab[16429]"
//line /usr/local/go/src/net/tcpsock.go:129
	}
//line /usr/local/go/src/net/tcpsock.go:129
	// _ = "end of CoverTab[16425]"
//line /usr/local/go/src/net/tcpsock.go:129
	_go_fuzz_dep_.CoverTab[16426]++
						n, err := c.readFrom(r)
						if err != nil && func() bool {
//line /usr/local/go/src/net/tcpsock.go:131
		_go_fuzz_dep_.CoverTab[16430]++
//line /usr/local/go/src/net/tcpsock.go:131
		return err != io.EOF
//line /usr/local/go/src/net/tcpsock.go:131
		// _ = "end of CoverTab[16430]"
//line /usr/local/go/src/net/tcpsock.go:131
	}() {
//line /usr/local/go/src/net/tcpsock.go:131
		_go_fuzz_dep_.CoverTab[16431]++
							err = &OpError{Op: "readfrom", Net: c.fd.net, Source: c.fd.laddr, Addr: c.fd.raddr, Err: err}
//line /usr/local/go/src/net/tcpsock.go:132
		// _ = "end of CoverTab[16431]"
	} else {
//line /usr/local/go/src/net/tcpsock.go:133
		_go_fuzz_dep_.CoverTab[16432]++
//line /usr/local/go/src/net/tcpsock.go:133
		// _ = "end of CoverTab[16432]"
//line /usr/local/go/src/net/tcpsock.go:133
	}
//line /usr/local/go/src/net/tcpsock.go:133
	// _ = "end of CoverTab[16426]"
//line /usr/local/go/src/net/tcpsock.go:133
	_go_fuzz_dep_.CoverTab[16427]++
						return n, err
//line /usr/local/go/src/net/tcpsock.go:134
	// _ = "end of CoverTab[16427]"
}

// CloseRead shuts down the reading side of the TCP connection.
//line /usr/local/go/src/net/tcpsock.go:137
// Most callers should just use Close.
//line /usr/local/go/src/net/tcpsock.go:139
func (c *TCPConn) CloseRead() error {
//line /usr/local/go/src/net/tcpsock.go:139
	_go_fuzz_dep_.CoverTab[16433]++
						if !c.ok() {
//line /usr/local/go/src/net/tcpsock.go:140
		_go_fuzz_dep_.CoverTab[16436]++
							return syscall.EINVAL
//line /usr/local/go/src/net/tcpsock.go:141
		// _ = "end of CoverTab[16436]"
	} else {
//line /usr/local/go/src/net/tcpsock.go:142
		_go_fuzz_dep_.CoverTab[16437]++
//line /usr/local/go/src/net/tcpsock.go:142
		// _ = "end of CoverTab[16437]"
//line /usr/local/go/src/net/tcpsock.go:142
	}
//line /usr/local/go/src/net/tcpsock.go:142
	// _ = "end of CoverTab[16433]"
//line /usr/local/go/src/net/tcpsock.go:142
	_go_fuzz_dep_.CoverTab[16434]++
						if err := c.fd.closeRead(); err != nil {
//line /usr/local/go/src/net/tcpsock.go:143
		_go_fuzz_dep_.CoverTab[16438]++
							return &OpError{Op: "close", Net: c.fd.net, Source: c.fd.laddr, Addr: c.fd.raddr, Err: err}
//line /usr/local/go/src/net/tcpsock.go:144
		// _ = "end of CoverTab[16438]"
	} else {
//line /usr/local/go/src/net/tcpsock.go:145
		_go_fuzz_dep_.CoverTab[16439]++
//line /usr/local/go/src/net/tcpsock.go:145
		// _ = "end of CoverTab[16439]"
//line /usr/local/go/src/net/tcpsock.go:145
	}
//line /usr/local/go/src/net/tcpsock.go:145
	// _ = "end of CoverTab[16434]"
//line /usr/local/go/src/net/tcpsock.go:145
	_go_fuzz_dep_.CoverTab[16435]++
						return nil
//line /usr/local/go/src/net/tcpsock.go:146
	// _ = "end of CoverTab[16435]"
}

// CloseWrite shuts down the writing side of the TCP connection.
//line /usr/local/go/src/net/tcpsock.go:149
// Most callers should just use Close.
//line /usr/local/go/src/net/tcpsock.go:151
func (c *TCPConn) CloseWrite() error {
//line /usr/local/go/src/net/tcpsock.go:151
	_go_fuzz_dep_.CoverTab[16440]++
						if !c.ok() {
//line /usr/local/go/src/net/tcpsock.go:152
		_go_fuzz_dep_.CoverTab[16443]++
							return syscall.EINVAL
//line /usr/local/go/src/net/tcpsock.go:153
		// _ = "end of CoverTab[16443]"
	} else {
//line /usr/local/go/src/net/tcpsock.go:154
		_go_fuzz_dep_.CoverTab[16444]++
//line /usr/local/go/src/net/tcpsock.go:154
		// _ = "end of CoverTab[16444]"
//line /usr/local/go/src/net/tcpsock.go:154
	}
//line /usr/local/go/src/net/tcpsock.go:154
	// _ = "end of CoverTab[16440]"
//line /usr/local/go/src/net/tcpsock.go:154
	_go_fuzz_dep_.CoverTab[16441]++
						if err := c.fd.closeWrite(); err != nil {
//line /usr/local/go/src/net/tcpsock.go:155
		_go_fuzz_dep_.CoverTab[16445]++
							return &OpError{Op: "close", Net: c.fd.net, Source: c.fd.laddr, Addr: c.fd.raddr, Err: err}
//line /usr/local/go/src/net/tcpsock.go:156
		// _ = "end of CoverTab[16445]"
	} else {
//line /usr/local/go/src/net/tcpsock.go:157
		_go_fuzz_dep_.CoverTab[16446]++
//line /usr/local/go/src/net/tcpsock.go:157
		// _ = "end of CoverTab[16446]"
//line /usr/local/go/src/net/tcpsock.go:157
	}
//line /usr/local/go/src/net/tcpsock.go:157
	// _ = "end of CoverTab[16441]"
//line /usr/local/go/src/net/tcpsock.go:157
	_go_fuzz_dep_.CoverTab[16442]++
						return nil
//line /usr/local/go/src/net/tcpsock.go:158
	// _ = "end of CoverTab[16442]"
}

// SetLinger sets the behavior of Close on a connection which still
//line /usr/local/go/src/net/tcpsock.go:161
// has data waiting to be sent or to be acknowledged.
//line /usr/local/go/src/net/tcpsock.go:161
//
//line /usr/local/go/src/net/tcpsock.go:161
// If sec < 0 (the default), the operating system finishes sending the
//line /usr/local/go/src/net/tcpsock.go:161
// data in the background.
//line /usr/local/go/src/net/tcpsock.go:161
//
//line /usr/local/go/src/net/tcpsock.go:161
// If sec == 0, the operating system discards any unsent or
//line /usr/local/go/src/net/tcpsock.go:161
// unacknowledged data.
//line /usr/local/go/src/net/tcpsock.go:161
//
//line /usr/local/go/src/net/tcpsock.go:161
// If sec > 0, the data is sent in the background as with sec < 0. On
//line /usr/local/go/src/net/tcpsock.go:161
// some operating systems after sec seconds have elapsed any remaining
//line /usr/local/go/src/net/tcpsock.go:161
// unsent data may be discarded.
//line /usr/local/go/src/net/tcpsock.go:173
func (c *TCPConn) SetLinger(sec int) error {
//line /usr/local/go/src/net/tcpsock.go:173
	_go_fuzz_dep_.CoverTab[16447]++
						if !c.ok() {
//line /usr/local/go/src/net/tcpsock.go:174
		_go_fuzz_dep_.CoverTab[16450]++
							return syscall.EINVAL
//line /usr/local/go/src/net/tcpsock.go:175
		// _ = "end of CoverTab[16450]"
	} else {
//line /usr/local/go/src/net/tcpsock.go:176
		_go_fuzz_dep_.CoverTab[16451]++
//line /usr/local/go/src/net/tcpsock.go:176
		// _ = "end of CoverTab[16451]"
//line /usr/local/go/src/net/tcpsock.go:176
	}
//line /usr/local/go/src/net/tcpsock.go:176
	// _ = "end of CoverTab[16447]"
//line /usr/local/go/src/net/tcpsock.go:176
	_go_fuzz_dep_.CoverTab[16448]++
						if err := setLinger(c.fd, sec); err != nil {
//line /usr/local/go/src/net/tcpsock.go:177
		_go_fuzz_dep_.CoverTab[16452]++
							return &OpError{Op: "set", Net: c.fd.net, Source: c.fd.laddr, Addr: c.fd.raddr, Err: err}
//line /usr/local/go/src/net/tcpsock.go:178
		// _ = "end of CoverTab[16452]"
	} else {
//line /usr/local/go/src/net/tcpsock.go:179
		_go_fuzz_dep_.CoverTab[16453]++
//line /usr/local/go/src/net/tcpsock.go:179
		// _ = "end of CoverTab[16453]"
//line /usr/local/go/src/net/tcpsock.go:179
	}
//line /usr/local/go/src/net/tcpsock.go:179
	// _ = "end of CoverTab[16448]"
//line /usr/local/go/src/net/tcpsock.go:179
	_go_fuzz_dep_.CoverTab[16449]++
						return nil
//line /usr/local/go/src/net/tcpsock.go:180
	// _ = "end of CoverTab[16449]"
}

// SetKeepAlive sets whether the operating system should send
//line /usr/local/go/src/net/tcpsock.go:183
// keep-alive messages on the connection.
//line /usr/local/go/src/net/tcpsock.go:185
func (c *TCPConn) SetKeepAlive(keepalive bool) error {
//line /usr/local/go/src/net/tcpsock.go:185
	_go_fuzz_dep_.CoverTab[16454]++
						if !c.ok() {
//line /usr/local/go/src/net/tcpsock.go:186
		_go_fuzz_dep_.CoverTab[16457]++
							return syscall.EINVAL
//line /usr/local/go/src/net/tcpsock.go:187
		// _ = "end of CoverTab[16457]"
	} else {
//line /usr/local/go/src/net/tcpsock.go:188
		_go_fuzz_dep_.CoverTab[16458]++
//line /usr/local/go/src/net/tcpsock.go:188
		// _ = "end of CoverTab[16458]"
//line /usr/local/go/src/net/tcpsock.go:188
	}
//line /usr/local/go/src/net/tcpsock.go:188
	// _ = "end of CoverTab[16454]"
//line /usr/local/go/src/net/tcpsock.go:188
	_go_fuzz_dep_.CoverTab[16455]++
						if err := setKeepAlive(c.fd, keepalive); err != nil {
//line /usr/local/go/src/net/tcpsock.go:189
		_go_fuzz_dep_.CoverTab[16459]++
							return &OpError{Op: "set", Net: c.fd.net, Source: c.fd.laddr, Addr: c.fd.raddr, Err: err}
//line /usr/local/go/src/net/tcpsock.go:190
		// _ = "end of CoverTab[16459]"
	} else {
//line /usr/local/go/src/net/tcpsock.go:191
		_go_fuzz_dep_.CoverTab[16460]++
//line /usr/local/go/src/net/tcpsock.go:191
		// _ = "end of CoverTab[16460]"
//line /usr/local/go/src/net/tcpsock.go:191
	}
//line /usr/local/go/src/net/tcpsock.go:191
	// _ = "end of CoverTab[16455]"
//line /usr/local/go/src/net/tcpsock.go:191
	_go_fuzz_dep_.CoverTab[16456]++
						return nil
//line /usr/local/go/src/net/tcpsock.go:192
	// _ = "end of CoverTab[16456]"
}

// SetKeepAlivePeriod sets period between keep-alives.
func (c *TCPConn) SetKeepAlivePeriod(d time.Duration) error {
//line /usr/local/go/src/net/tcpsock.go:196
	_go_fuzz_dep_.CoverTab[16461]++
						if !c.ok() {
//line /usr/local/go/src/net/tcpsock.go:197
		_go_fuzz_dep_.CoverTab[16464]++
							return syscall.EINVAL
//line /usr/local/go/src/net/tcpsock.go:198
		// _ = "end of CoverTab[16464]"
	} else {
//line /usr/local/go/src/net/tcpsock.go:199
		_go_fuzz_dep_.CoverTab[16465]++
//line /usr/local/go/src/net/tcpsock.go:199
		// _ = "end of CoverTab[16465]"
//line /usr/local/go/src/net/tcpsock.go:199
	}
//line /usr/local/go/src/net/tcpsock.go:199
	// _ = "end of CoverTab[16461]"
//line /usr/local/go/src/net/tcpsock.go:199
	_go_fuzz_dep_.CoverTab[16462]++
						if err := setKeepAlivePeriod(c.fd, d); err != nil {
//line /usr/local/go/src/net/tcpsock.go:200
		_go_fuzz_dep_.CoverTab[16466]++
							return &OpError{Op: "set", Net: c.fd.net, Source: c.fd.laddr, Addr: c.fd.raddr, Err: err}
//line /usr/local/go/src/net/tcpsock.go:201
		// _ = "end of CoverTab[16466]"
	} else {
//line /usr/local/go/src/net/tcpsock.go:202
		_go_fuzz_dep_.CoverTab[16467]++
//line /usr/local/go/src/net/tcpsock.go:202
		// _ = "end of CoverTab[16467]"
//line /usr/local/go/src/net/tcpsock.go:202
	}
//line /usr/local/go/src/net/tcpsock.go:202
	// _ = "end of CoverTab[16462]"
//line /usr/local/go/src/net/tcpsock.go:202
	_go_fuzz_dep_.CoverTab[16463]++
						return nil
//line /usr/local/go/src/net/tcpsock.go:203
	// _ = "end of CoverTab[16463]"
}

// SetNoDelay controls whether the operating system should delay
//line /usr/local/go/src/net/tcpsock.go:206
// packet transmission in hopes of sending fewer packets (Nagle's
//line /usr/local/go/src/net/tcpsock.go:206
// algorithm).  The default is true (no delay), meaning that data is
//line /usr/local/go/src/net/tcpsock.go:206
// sent as soon as possible after a Write.
//line /usr/local/go/src/net/tcpsock.go:210
func (c *TCPConn) SetNoDelay(noDelay bool) error {
//line /usr/local/go/src/net/tcpsock.go:210
	_go_fuzz_dep_.CoverTab[16468]++
						if !c.ok() {
//line /usr/local/go/src/net/tcpsock.go:211
		_go_fuzz_dep_.CoverTab[16471]++
							return syscall.EINVAL
//line /usr/local/go/src/net/tcpsock.go:212
		// _ = "end of CoverTab[16471]"
	} else {
//line /usr/local/go/src/net/tcpsock.go:213
		_go_fuzz_dep_.CoverTab[16472]++
//line /usr/local/go/src/net/tcpsock.go:213
		// _ = "end of CoverTab[16472]"
//line /usr/local/go/src/net/tcpsock.go:213
	}
//line /usr/local/go/src/net/tcpsock.go:213
	// _ = "end of CoverTab[16468]"
//line /usr/local/go/src/net/tcpsock.go:213
	_go_fuzz_dep_.CoverTab[16469]++
						if err := setNoDelay(c.fd, noDelay); err != nil {
//line /usr/local/go/src/net/tcpsock.go:214
		_go_fuzz_dep_.CoverTab[16473]++
							return &OpError{Op: "set", Net: c.fd.net, Source: c.fd.laddr, Addr: c.fd.raddr, Err: err}
//line /usr/local/go/src/net/tcpsock.go:215
		// _ = "end of CoverTab[16473]"
	} else {
//line /usr/local/go/src/net/tcpsock.go:216
		_go_fuzz_dep_.CoverTab[16474]++
//line /usr/local/go/src/net/tcpsock.go:216
		// _ = "end of CoverTab[16474]"
//line /usr/local/go/src/net/tcpsock.go:216
	}
//line /usr/local/go/src/net/tcpsock.go:216
	// _ = "end of CoverTab[16469]"
//line /usr/local/go/src/net/tcpsock.go:216
	_go_fuzz_dep_.CoverTab[16470]++
						return nil
//line /usr/local/go/src/net/tcpsock.go:217
	// _ = "end of CoverTab[16470]"
}

func newTCPConn(fd *netFD, keepAlive time.Duration, keepAliveHook func(time.Duration)) *TCPConn {
//line /usr/local/go/src/net/tcpsock.go:220
	_go_fuzz_dep_.CoverTab[16475]++
						setNoDelay(fd, true)
						if keepAlive == 0 {
//line /usr/local/go/src/net/tcpsock.go:222
		_go_fuzz_dep_.CoverTab[16478]++
							keepAlive = defaultTCPKeepAlive
//line /usr/local/go/src/net/tcpsock.go:223
		// _ = "end of CoverTab[16478]"
	} else {
//line /usr/local/go/src/net/tcpsock.go:224
		_go_fuzz_dep_.CoverTab[16479]++
//line /usr/local/go/src/net/tcpsock.go:224
		// _ = "end of CoverTab[16479]"
//line /usr/local/go/src/net/tcpsock.go:224
	}
//line /usr/local/go/src/net/tcpsock.go:224
	// _ = "end of CoverTab[16475]"
//line /usr/local/go/src/net/tcpsock.go:224
	_go_fuzz_dep_.CoverTab[16476]++
						if keepAlive > 0 {
//line /usr/local/go/src/net/tcpsock.go:225
		_go_fuzz_dep_.CoverTab[16480]++
							setKeepAlive(fd, true)
							setKeepAlivePeriod(fd, keepAlive)
							if keepAliveHook != nil {
//line /usr/local/go/src/net/tcpsock.go:228
			_go_fuzz_dep_.CoverTab[16481]++
								keepAliveHook(keepAlive)
//line /usr/local/go/src/net/tcpsock.go:229
			// _ = "end of CoverTab[16481]"
		} else {
//line /usr/local/go/src/net/tcpsock.go:230
			_go_fuzz_dep_.CoverTab[16482]++
//line /usr/local/go/src/net/tcpsock.go:230
			// _ = "end of CoverTab[16482]"
//line /usr/local/go/src/net/tcpsock.go:230
		}
//line /usr/local/go/src/net/tcpsock.go:230
		// _ = "end of CoverTab[16480]"
	} else {
//line /usr/local/go/src/net/tcpsock.go:231
		_go_fuzz_dep_.CoverTab[16483]++
//line /usr/local/go/src/net/tcpsock.go:231
		// _ = "end of CoverTab[16483]"
//line /usr/local/go/src/net/tcpsock.go:231
	}
//line /usr/local/go/src/net/tcpsock.go:231
	// _ = "end of CoverTab[16476]"
//line /usr/local/go/src/net/tcpsock.go:231
	_go_fuzz_dep_.CoverTab[16477]++
						return &TCPConn{conn{fd}}
//line /usr/local/go/src/net/tcpsock.go:232
	// _ = "end of CoverTab[16477]"
}

// DialTCP acts like Dial for TCP networks.
//line /usr/local/go/src/net/tcpsock.go:235
//
//line /usr/local/go/src/net/tcpsock.go:235
// The network must be a TCP network name; see func Dial for details.
//line /usr/local/go/src/net/tcpsock.go:235
//
//line /usr/local/go/src/net/tcpsock.go:235
// If laddr is nil, a local address is automatically chosen.
//line /usr/local/go/src/net/tcpsock.go:235
// If the IP field of raddr is nil or an unspecified IP address, the
//line /usr/local/go/src/net/tcpsock.go:235
// local system is assumed.
//line /usr/local/go/src/net/tcpsock.go:242
func DialTCP(network string, laddr, raddr *TCPAddr) (*TCPConn, error) {
//line /usr/local/go/src/net/tcpsock.go:242
	_go_fuzz_dep_.CoverTab[16484]++
						switch network {
	case "tcp", "tcp4", "tcp6":
//line /usr/local/go/src/net/tcpsock.go:244
		_go_fuzz_dep_.CoverTab[16488]++
//line /usr/local/go/src/net/tcpsock.go:244
		// _ = "end of CoverTab[16488]"
	default:
//line /usr/local/go/src/net/tcpsock.go:245
		_go_fuzz_dep_.CoverTab[16489]++
							return nil, &OpError{Op: "dial", Net: network, Source: laddr.opAddr(), Addr: raddr.opAddr(), Err: UnknownNetworkError(network)}
//line /usr/local/go/src/net/tcpsock.go:246
		// _ = "end of CoverTab[16489]"
	}
//line /usr/local/go/src/net/tcpsock.go:247
	// _ = "end of CoverTab[16484]"
//line /usr/local/go/src/net/tcpsock.go:247
	_go_fuzz_dep_.CoverTab[16485]++
						if raddr == nil {
//line /usr/local/go/src/net/tcpsock.go:248
		_go_fuzz_dep_.CoverTab[16490]++
							return nil, &OpError{Op: "dial", Net: network, Source: laddr.opAddr(), Addr: nil, Err: errMissingAddress}
//line /usr/local/go/src/net/tcpsock.go:249
		// _ = "end of CoverTab[16490]"
	} else {
//line /usr/local/go/src/net/tcpsock.go:250
		_go_fuzz_dep_.CoverTab[16491]++
//line /usr/local/go/src/net/tcpsock.go:250
		// _ = "end of CoverTab[16491]"
//line /usr/local/go/src/net/tcpsock.go:250
	}
//line /usr/local/go/src/net/tcpsock.go:250
	// _ = "end of CoverTab[16485]"
//line /usr/local/go/src/net/tcpsock.go:250
	_go_fuzz_dep_.CoverTab[16486]++
						sd := &sysDialer{network: network, address: raddr.String()}
						c, err := sd.dialTCP(context.Background(), laddr, raddr)
						if err != nil {
//line /usr/local/go/src/net/tcpsock.go:253
		_go_fuzz_dep_.CoverTab[16492]++
							return nil, &OpError{Op: "dial", Net: network, Source: laddr.opAddr(), Addr: raddr.opAddr(), Err: err}
//line /usr/local/go/src/net/tcpsock.go:254
		// _ = "end of CoverTab[16492]"
	} else {
//line /usr/local/go/src/net/tcpsock.go:255
		_go_fuzz_dep_.CoverTab[16493]++
//line /usr/local/go/src/net/tcpsock.go:255
		// _ = "end of CoverTab[16493]"
//line /usr/local/go/src/net/tcpsock.go:255
	}
//line /usr/local/go/src/net/tcpsock.go:255
	// _ = "end of CoverTab[16486]"
//line /usr/local/go/src/net/tcpsock.go:255
	_go_fuzz_dep_.CoverTab[16487]++
						return c, nil
//line /usr/local/go/src/net/tcpsock.go:256
	// _ = "end of CoverTab[16487]"
}

// TCPListener is a TCP network listener. Clients should typically
//line /usr/local/go/src/net/tcpsock.go:259
// use variables of type Listener instead of assuming TCP.
//line /usr/local/go/src/net/tcpsock.go:261
type TCPListener struct {
	fd	*netFD
	lc	ListenConfig
}

// SyscallConn returns a raw network connection.
//line /usr/local/go/src/net/tcpsock.go:266
// This implements the syscall.Conn interface.
//line /usr/local/go/src/net/tcpsock.go:266
//
//line /usr/local/go/src/net/tcpsock.go:266
// The returned RawConn only supports calling Control. Read and
//line /usr/local/go/src/net/tcpsock.go:266
// Write return an error.
//line /usr/local/go/src/net/tcpsock.go:271
func (l *TCPListener) SyscallConn() (syscall.RawConn, error) {
//line /usr/local/go/src/net/tcpsock.go:271
	_go_fuzz_dep_.CoverTab[16494]++
						if !l.ok() {
//line /usr/local/go/src/net/tcpsock.go:272
		_go_fuzz_dep_.CoverTab[16496]++
							return nil, syscall.EINVAL
//line /usr/local/go/src/net/tcpsock.go:273
		// _ = "end of CoverTab[16496]"
	} else {
//line /usr/local/go/src/net/tcpsock.go:274
		_go_fuzz_dep_.CoverTab[16497]++
//line /usr/local/go/src/net/tcpsock.go:274
		// _ = "end of CoverTab[16497]"
//line /usr/local/go/src/net/tcpsock.go:274
	}
//line /usr/local/go/src/net/tcpsock.go:274
	// _ = "end of CoverTab[16494]"
//line /usr/local/go/src/net/tcpsock.go:274
	_go_fuzz_dep_.CoverTab[16495]++
						return newRawListener(l.fd)
//line /usr/local/go/src/net/tcpsock.go:275
	// _ = "end of CoverTab[16495]"
}

// AcceptTCP accepts the next incoming call and returns the new
//line /usr/local/go/src/net/tcpsock.go:278
// connection.
//line /usr/local/go/src/net/tcpsock.go:280
func (l *TCPListener) AcceptTCP() (*TCPConn, error) {
//line /usr/local/go/src/net/tcpsock.go:280
	_go_fuzz_dep_.CoverTab[16498]++
						if !l.ok() {
//line /usr/local/go/src/net/tcpsock.go:281
		_go_fuzz_dep_.CoverTab[16501]++
							return nil, syscall.EINVAL
//line /usr/local/go/src/net/tcpsock.go:282
		// _ = "end of CoverTab[16501]"
	} else {
//line /usr/local/go/src/net/tcpsock.go:283
		_go_fuzz_dep_.CoverTab[16502]++
//line /usr/local/go/src/net/tcpsock.go:283
		// _ = "end of CoverTab[16502]"
//line /usr/local/go/src/net/tcpsock.go:283
	}
//line /usr/local/go/src/net/tcpsock.go:283
	// _ = "end of CoverTab[16498]"
//line /usr/local/go/src/net/tcpsock.go:283
	_go_fuzz_dep_.CoverTab[16499]++
						c, err := l.accept()
						if err != nil {
//line /usr/local/go/src/net/tcpsock.go:285
		_go_fuzz_dep_.CoverTab[16503]++
							return nil, &OpError{Op: "accept", Net: l.fd.net, Source: nil, Addr: l.fd.laddr, Err: err}
//line /usr/local/go/src/net/tcpsock.go:286
		// _ = "end of CoverTab[16503]"
	} else {
//line /usr/local/go/src/net/tcpsock.go:287
		_go_fuzz_dep_.CoverTab[16504]++
//line /usr/local/go/src/net/tcpsock.go:287
		// _ = "end of CoverTab[16504]"
//line /usr/local/go/src/net/tcpsock.go:287
	}
//line /usr/local/go/src/net/tcpsock.go:287
	// _ = "end of CoverTab[16499]"
//line /usr/local/go/src/net/tcpsock.go:287
	_go_fuzz_dep_.CoverTab[16500]++
						return c, nil
//line /usr/local/go/src/net/tcpsock.go:288
	// _ = "end of CoverTab[16500]"
}

// Accept implements the Accept method in the Listener interface; it
//line /usr/local/go/src/net/tcpsock.go:291
// waits for the next call and returns a generic Conn.
//line /usr/local/go/src/net/tcpsock.go:293
func (l *TCPListener) Accept() (Conn, error) {
//line /usr/local/go/src/net/tcpsock.go:293
	_go_fuzz_dep_.CoverTab[16505]++
						if !l.ok() {
//line /usr/local/go/src/net/tcpsock.go:294
		_go_fuzz_dep_.CoverTab[16508]++
							return nil, syscall.EINVAL
//line /usr/local/go/src/net/tcpsock.go:295
		// _ = "end of CoverTab[16508]"
	} else {
//line /usr/local/go/src/net/tcpsock.go:296
		_go_fuzz_dep_.CoverTab[16509]++
//line /usr/local/go/src/net/tcpsock.go:296
		// _ = "end of CoverTab[16509]"
//line /usr/local/go/src/net/tcpsock.go:296
	}
//line /usr/local/go/src/net/tcpsock.go:296
	// _ = "end of CoverTab[16505]"
//line /usr/local/go/src/net/tcpsock.go:296
	_go_fuzz_dep_.CoverTab[16506]++
						c, err := l.accept()
						if err != nil {
//line /usr/local/go/src/net/tcpsock.go:298
		_go_fuzz_dep_.CoverTab[16510]++
							return nil, &OpError{Op: "accept", Net: l.fd.net, Source: nil, Addr: l.fd.laddr, Err: err}
//line /usr/local/go/src/net/tcpsock.go:299
		// _ = "end of CoverTab[16510]"
	} else {
//line /usr/local/go/src/net/tcpsock.go:300
		_go_fuzz_dep_.CoverTab[16511]++
//line /usr/local/go/src/net/tcpsock.go:300
		// _ = "end of CoverTab[16511]"
//line /usr/local/go/src/net/tcpsock.go:300
	}
//line /usr/local/go/src/net/tcpsock.go:300
	// _ = "end of CoverTab[16506]"
//line /usr/local/go/src/net/tcpsock.go:300
	_go_fuzz_dep_.CoverTab[16507]++
						return c, nil
//line /usr/local/go/src/net/tcpsock.go:301
	// _ = "end of CoverTab[16507]"
}

// Close stops listening on the TCP address.
//line /usr/local/go/src/net/tcpsock.go:304
// Already Accepted connections are not closed.
//line /usr/local/go/src/net/tcpsock.go:306
func (l *TCPListener) Close() error {
//line /usr/local/go/src/net/tcpsock.go:306
	_go_fuzz_dep_.CoverTab[16512]++
						if !l.ok() {
//line /usr/local/go/src/net/tcpsock.go:307
		_go_fuzz_dep_.CoverTab[16515]++
							return syscall.EINVAL
//line /usr/local/go/src/net/tcpsock.go:308
		// _ = "end of CoverTab[16515]"
	} else {
//line /usr/local/go/src/net/tcpsock.go:309
		_go_fuzz_dep_.CoverTab[16516]++
//line /usr/local/go/src/net/tcpsock.go:309
		// _ = "end of CoverTab[16516]"
//line /usr/local/go/src/net/tcpsock.go:309
	}
//line /usr/local/go/src/net/tcpsock.go:309
	// _ = "end of CoverTab[16512]"
//line /usr/local/go/src/net/tcpsock.go:309
	_go_fuzz_dep_.CoverTab[16513]++
						if err := l.close(); err != nil {
//line /usr/local/go/src/net/tcpsock.go:310
		_go_fuzz_dep_.CoverTab[16517]++
							return &OpError{Op: "close", Net: l.fd.net, Source: nil, Addr: l.fd.laddr, Err: err}
//line /usr/local/go/src/net/tcpsock.go:311
		// _ = "end of CoverTab[16517]"
	} else {
//line /usr/local/go/src/net/tcpsock.go:312
		_go_fuzz_dep_.CoverTab[16518]++
//line /usr/local/go/src/net/tcpsock.go:312
		// _ = "end of CoverTab[16518]"
//line /usr/local/go/src/net/tcpsock.go:312
	}
//line /usr/local/go/src/net/tcpsock.go:312
	// _ = "end of CoverTab[16513]"
//line /usr/local/go/src/net/tcpsock.go:312
	_go_fuzz_dep_.CoverTab[16514]++
						return nil
//line /usr/local/go/src/net/tcpsock.go:313
	// _ = "end of CoverTab[16514]"
}

// Addr returns the listener's network address, a *TCPAddr.
//line /usr/local/go/src/net/tcpsock.go:316
// The Addr returned is shared by all invocations of Addr, so
//line /usr/local/go/src/net/tcpsock.go:316
// do not modify it.
//line /usr/local/go/src/net/tcpsock.go:319
func (l *TCPListener) Addr() Addr {
//line /usr/local/go/src/net/tcpsock.go:319
	_go_fuzz_dep_.CoverTab[16519]++
//line /usr/local/go/src/net/tcpsock.go:319
	return l.fd.laddr
//line /usr/local/go/src/net/tcpsock.go:319
	// _ = "end of CoverTab[16519]"
//line /usr/local/go/src/net/tcpsock.go:319
}

// SetDeadline sets the deadline associated with the listener.
//line /usr/local/go/src/net/tcpsock.go:321
// A zero time value disables the deadline.
//line /usr/local/go/src/net/tcpsock.go:323
func (l *TCPListener) SetDeadline(t time.Time) error {
//line /usr/local/go/src/net/tcpsock.go:323
	_go_fuzz_dep_.CoverTab[16520]++
						if !l.ok() {
//line /usr/local/go/src/net/tcpsock.go:324
		_go_fuzz_dep_.CoverTab[16523]++
							return syscall.EINVAL
//line /usr/local/go/src/net/tcpsock.go:325
		// _ = "end of CoverTab[16523]"
	} else {
//line /usr/local/go/src/net/tcpsock.go:326
		_go_fuzz_dep_.CoverTab[16524]++
//line /usr/local/go/src/net/tcpsock.go:326
		// _ = "end of CoverTab[16524]"
//line /usr/local/go/src/net/tcpsock.go:326
	}
//line /usr/local/go/src/net/tcpsock.go:326
	// _ = "end of CoverTab[16520]"
//line /usr/local/go/src/net/tcpsock.go:326
	_go_fuzz_dep_.CoverTab[16521]++
						if err := l.fd.pfd.SetDeadline(t); err != nil {
//line /usr/local/go/src/net/tcpsock.go:327
		_go_fuzz_dep_.CoverTab[16525]++
							return &OpError{Op: "set", Net: l.fd.net, Source: nil, Addr: l.fd.laddr, Err: err}
//line /usr/local/go/src/net/tcpsock.go:328
		// _ = "end of CoverTab[16525]"
	} else {
//line /usr/local/go/src/net/tcpsock.go:329
		_go_fuzz_dep_.CoverTab[16526]++
//line /usr/local/go/src/net/tcpsock.go:329
		// _ = "end of CoverTab[16526]"
//line /usr/local/go/src/net/tcpsock.go:329
	}
//line /usr/local/go/src/net/tcpsock.go:329
	// _ = "end of CoverTab[16521]"
//line /usr/local/go/src/net/tcpsock.go:329
	_go_fuzz_dep_.CoverTab[16522]++
						return nil
//line /usr/local/go/src/net/tcpsock.go:330
	// _ = "end of CoverTab[16522]"
}

// File returns a copy of the underlying os.File.
//line /usr/local/go/src/net/tcpsock.go:333
// It is the caller's responsibility to close f when finished.
//line /usr/local/go/src/net/tcpsock.go:333
// Closing l does not affect f, and closing f does not affect l.
//line /usr/local/go/src/net/tcpsock.go:333
//
//line /usr/local/go/src/net/tcpsock.go:333
// The returned os.File's file descriptor is different from the
//line /usr/local/go/src/net/tcpsock.go:333
// connection's. Attempting to change properties of the original
//line /usr/local/go/src/net/tcpsock.go:333
// using this duplicate may or may not have the desired effect.
//line /usr/local/go/src/net/tcpsock.go:340
func (l *TCPListener) File() (f *os.File, err error) {
//line /usr/local/go/src/net/tcpsock.go:340
	_go_fuzz_dep_.CoverTab[16527]++
						if !l.ok() {
//line /usr/local/go/src/net/tcpsock.go:341
		_go_fuzz_dep_.CoverTab[16530]++
							return nil, syscall.EINVAL
//line /usr/local/go/src/net/tcpsock.go:342
		// _ = "end of CoverTab[16530]"
	} else {
//line /usr/local/go/src/net/tcpsock.go:343
		_go_fuzz_dep_.CoverTab[16531]++
//line /usr/local/go/src/net/tcpsock.go:343
		// _ = "end of CoverTab[16531]"
//line /usr/local/go/src/net/tcpsock.go:343
	}
//line /usr/local/go/src/net/tcpsock.go:343
	// _ = "end of CoverTab[16527]"
//line /usr/local/go/src/net/tcpsock.go:343
	_go_fuzz_dep_.CoverTab[16528]++
						f, err = l.file()
						if err != nil {
//line /usr/local/go/src/net/tcpsock.go:345
		_go_fuzz_dep_.CoverTab[16532]++
							return nil, &OpError{Op: "file", Net: l.fd.net, Source: nil, Addr: l.fd.laddr, Err: err}
//line /usr/local/go/src/net/tcpsock.go:346
		// _ = "end of CoverTab[16532]"
	} else {
//line /usr/local/go/src/net/tcpsock.go:347
		_go_fuzz_dep_.CoverTab[16533]++
//line /usr/local/go/src/net/tcpsock.go:347
		// _ = "end of CoverTab[16533]"
//line /usr/local/go/src/net/tcpsock.go:347
	}
//line /usr/local/go/src/net/tcpsock.go:347
	// _ = "end of CoverTab[16528]"
//line /usr/local/go/src/net/tcpsock.go:347
	_go_fuzz_dep_.CoverTab[16529]++
						return
//line /usr/local/go/src/net/tcpsock.go:348
	// _ = "end of CoverTab[16529]"
}

// ListenTCP acts like Listen for TCP networks.
//line /usr/local/go/src/net/tcpsock.go:351
//
//line /usr/local/go/src/net/tcpsock.go:351
// The network must be a TCP network name; see func Dial for details.
//line /usr/local/go/src/net/tcpsock.go:351
//
//line /usr/local/go/src/net/tcpsock.go:351
// If the IP field of laddr is nil or an unspecified IP address,
//line /usr/local/go/src/net/tcpsock.go:351
// ListenTCP listens on all available unicast and anycast IP addresses
//line /usr/local/go/src/net/tcpsock.go:351
// of the local system.
//line /usr/local/go/src/net/tcpsock.go:351
// If the Port field of laddr is 0, a port number is automatically
//line /usr/local/go/src/net/tcpsock.go:351
// chosen.
//line /usr/local/go/src/net/tcpsock.go:360
func ListenTCP(network string, laddr *TCPAddr) (*TCPListener, error) {
//line /usr/local/go/src/net/tcpsock.go:360
	_go_fuzz_dep_.CoverTab[16534]++
						switch network {
	case "tcp", "tcp4", "tcp6":
//line /usr/local/go/src/net/tcpsock.go:362
		_go_fuzz_dep_.CoverTab[16538]++
//line /usr/local/go/src/net/tcpsock.go:362
		// _ = "end of CoverTab[16538]"
	default:
//line /usr/local/go/src/net/tcpsock.go:363
		_go_fuzz_dep_.CoverTab[16539]++
							return nil, &OpError{Op: "listen", Net: network, Source: nil, Addr: laddr.opAddr(), Err: UnknownNetworkError(network)}
//line /usr/local/go/src/net/tcpsock.go:364
		// _ = "end of CoverTab[16539]"
	}
//line /usr/local/go/src/net/tcpsock.go:365
	// _ = "end of CoverTab[16534]"
//line /usr/local/go/src/net/tcpsock.go:365
	_go_fuzz_dep_.CoverTab[16535]++
						if laddr == nil {
//line /usr/local/go/src/net/tcpsock.go:366
		_go_fuzz_dep_.CoverTab[16540]++
							laddr = &TCPAddr{}
//line /usr/local/go/src/net/tcpsock.go:367
		// _ = "end of CoverTab[16540]"
	} else {
//line /usr/local/go/src/net/tcpsock.go:368
		_go_fuzz_dep_.CoverTab[16541]++
//line /usr/local/go/src/net/tcpsock.go:368
		// _ = "end of CoverTab[16541]"
//line /usr/local/go/src/net/tcpsock.go:368
	}
//line /usr/local/go/src/net/tcpsock.go:368
	// _ = "end of CoverTab[16535]"
//line /usr/local/go/src/net/tcpsock.go:368
	_go_fuzz_dep_.CoverTab[16536]++
						sl := &sysListener{network: network, address: laddr.String()}
						ln, err := sl.listenTCP(context.Background(), laddr)
						if err != nil {
//line /usr/local/go/src/net/tcpsock.go:371
		_go_fuzz_dep_.CoverTab[16542]++
							return nil, &OpError{Op: "listen", Net: network, Source: nil, Addr: laddr.opAddr(), Err: err}
//line /usr/local/go/src/net/tcpsock.go:372
		// _ = "end of CoverTab[16542]"
	} else {
//line /usr/local/go/src/net/tcpsock.go:373
		_go_fuzz_dep_.CoverTab[16543]++
//line /usr/local/go/src/net/tcpsock.go:373
		// _ = "end of CoverTab[16543]"
//line /usr/local/go/src/net/tcpsock.go:373
	}
//line /usr/local/go/src/net/tcpsock.go:373
	// _ = "end of CoverTab[16536]"
//line /usr/local/go/src/net/tcpsock.go:373
	_go_fuzz_dep_.CoverTab[16537]++
						return ln, nil
//line /usr/local/go/src/net/tcpsock.go:374
	// _ = "end of CoverTab[16537]"
}

// roundDurationUp rounds d to the next multiple of to.
func roundDurationUp(d time.Duration, to time.Duration) time.Duration {
//line /usr/local/go/src/net/tcpsock.go:378
	_go_fuzz_dep_.CoverTab[16544]++
						return (d + to - 1) / to
//line /usr/local/go/src/net/tcpsock.go:379
	// _ = "end of CoverTab[16544]"
}

//line /usr/local/go/src/net/tcpsock.go:380
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/tcpsock.go:380
var _ = _go_fuzz_dep_.CoverTab
