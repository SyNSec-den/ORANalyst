// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/net/unixsock.go:5
package net

//line /usr/local/go/src/net/unixsock.go:5
import (
//line /usr/local/go/src/net/unixsock.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/unixsock.go:5
)
//line /usr/local/go/src/net/unixsock.go:5
import (
//line /usr/local/go/src/net/unixsock.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/unixsock.go:5
)

import (
	"context"
	"os"
	"sync"
	"syscall"
	"time"
)

//line /usr/local/go/src/net/unixsock.go:21
// UnixAddr represents the address of a Unix domain socket end point.
type UnixAddr struct {
	Name	string
	Net	string
}

// Network returns the address's network name, "unix", "unixgram" or
//line /usr/local/go/src/net/unixsock.go:27
// "unixpacket".
//line /usr/local/go/src/net/unixsock.go:29
func (a *UnixAddr) Network() string {
//line /usr/local/go/src/net/unixsock.go:29
	_go_fuzz_dep_.CoverTab[16933]++
						return a.Net
//line /usr/local/go/src/net/unixsock.go:30
	// _ = "end of CoverTab[16933]"
}

func (a *UnixAddr) String() string {
//line /usr/local/go/src/net/unixsock.go:33
	_go_fuzz_dep_.CoverTab[16934]++
						if a == nil {
//line /usr/local/go/src/net/unixsock.go:34
		_go_fuzz_dep_.CoverTab[16936]++
							return "<nil>"
//line /usr/local/go/src/net/unixsock.go:35
		// _ = "end of CoverTab[16936]"
	} else {
//line /usr/local/go/src/net/unixsock.go:36
		_go_fuzz_dep_.CoverTab[16937]++
//line /usr/local/go/src/net/unixsock.go:36
		// _ = "end of CoverTab[16937]"
//line /usr/local/go/src/net/unixsock.go:36
	}
//line /usr/local/go/src/net/unixsock.go:36
	// _ = "end of CoverTab[16934]"
//line /usr/local/go/src/net/unixsock.go:36
	_go_fuzz_dep_.CoverTab[16935]++
						return a.Name
//line /usr/local/go/src/net/unixsock.go:37
	// _ = "end of CoverTab[16935]"
}

func (a *UnixAddr) isWildcard() bool {
//line /usr/local/go/src/net/unixsock.go:40
	_go_fuzz_dep_.CoverTab[16938]++
						return a == nil || func() bool {
//line /usr/local/go/src/net/unixsock.go:41
		_go_fuzz_dep_.CoverTab[16939]++
//line /usr/local/go/src/net/unixsock.go:41
		return a.Name == ""
//line /usr/local/go/src/net/unixsock.go:41
		// _ = "end of CoverTab[16939]"
//line /usr/local/go/src/net/unixsock.go:41
	}()
//line /usr/local/go/src/net/unixsock.go:41
	// _ = "end of CoverTab[16938]"
}

func (a *UnixAddr) opAddr() Addr {
//line /usr/local/go/src/net/unixsock.go:44
	_go_fuzz_dep_.CoverTab[16940]++
						if a == nil {
//line /usr/local/go/src/net/unixsock.go:45
		_go_fuzz_dep_.CoverTab[16942]++
							return nil
//line /usr/local/go/src/net/unixsock.go:46
		// _ = "end of CoverTab[16942]"
	} else {
//line /usr/local/go/src/net/unixsock.go:47
		_go_fuzz_dep_.CoverTab[16943]++
//line /usr/local/go/src/net/unixsock.go:47
		// _ = "end of CoverTab[16943]"
//line /usr/local/go/src/net/unixsock.go:47
	}
//line /usr/local/go/src/net/unixsock.go:47
	// _ = "end of CoverTab[16940]"
//line /usr/local/go/src/net/unixsock.go:47
	_go_fuzz_dep_.CoverTab[16941]++
						return a
//line /usr/local/go/src/net/unixsock.go:48
	// _ = "end of CoverTab[16941]"
}

// ResolveUnixAddr returns an address of Unix domain socket end point.
//line /usr/local/go/src/net/unixsock.go:51
//
//line /usr/local/go/src/net/unixsock.go:51
// The network must be a Unix network name.
//line /usr/local/go/src/net/unixsock.go:51
//
//line /usr/local/go/src/net/unixsock.go:51
// See func Dial for a description of the network and address
//line /usr/local/go/src/net/unixsock.go:51
// parameters.
//line /usr/local/go/src/net/unixsock.go:57
func ResolveUnixAddr(network, address string) (*UnixAddr, error) {
//line /usr/local/go/src/net/unixsock.go:57
	_go_fuzz_dep_.CoverTab[16944]++
						switch network {
	case "unix", "unixgram", "unixpacket":
//line /usr/local/go/src/net/unixsock.go:59
		_go_fuzz_dep_.CoverTab[16945]++
							return &UnixAddr{Name: address, Net: network}, nil
//line /usr/local/go/src/net/unixsock.go:60
		// _ = "end of CoverTab[16945]"
	default:
//line /usr/local/go/src/net/unixsock.go:61
		_go_fuzz_dep_.CoverTab[16946]++
							return nil, UnknownNetworkError(network)
//line /usr/local/go/src/net/unixsock.go:62
		// _ = "end of CoverTab[16946]"
	}
//line /usr/local/go/src/net/unixsock.go:63
	// _ = "end of CoverTab[16944]"
}

// UnixConn is an implementation of the Conn interface for connections
//line /usr/local/go/src/net/unixsock.go:66
// to Unix domain sockets.
//line /usr/local/go/src/net/unixsock.go:68
type UnixConn struct {
	conn
}

// SyscallConn returns a raw network connection.
//line /usr/local/go/src/net/unixsock.go:72
// This implements the syscall.Conn interface.
//line /usr/local/go/src/net/unixsock.go:74
func (c *UnixConn) SyscallConn() (syscall.RawConn, error) {
//line /usr/local/go/src/net/unixsock.go:74
	_go_fuzz_dep_.CoverTab[16947]++
						if !c.ok() {
//line /usr/local/go/src/net/unixsock.go:75
		_go_fuzz_dep_.CoverTab[16949]++
							return nil, syscall.EINVAL
//line /usr/local/go/src/net/unixsock.go:76
		// _ = "end of CoverTab[16949]"
	} else {
//line /usr/local/go/src/net/unixsock.go:77
		_go_fuzz_dep_.CoverTab[16950]++
//line /usr/local/go/src/net/unixsock.go:77
		// _ = "end of CoverTab[16950]"
//line /usr/local/go/src/net/unixsock.go:77
	}
//line /usr/local/go/src/net/unixsock.go:77
	// _ = "end of CoverTab[16947]"
//line /usr/local/go/src/net/unixsock.go:77
	_go_fuzz_dep_.CoverTab[16948]++
						return newRawConn(c.fd)
//line /usr/local/go/src/net/unixsock.go:78
	// _ = "end of CoverTab[16948]"
}

// CloseRead shuts down the reading side of the Unix domain connection.
//line /usr/local/go/src/net/unixsock.go:81
// Most callers should just use Close.
//line /usr/local/go/src/net/unixsock.go:83
func (c *UnixConn) CloseRead() error {
//line /usr/local/go/src/net/unixsock.go:83
	_go_fuzz_dep_.CoverTab[16951]++
						if !c.ok() {
//line /usr/local/go/src/net/unixsock.go:84
		_go_fuzz_dep_.CoverTab[16954]++
							return syscall.EINVAL
//line /usr/local/go/src/net/unixsock.go:85
		// _ = "end of CoverTab[16954]"
	} else {
//line /usr/local/go/src/net/unixsock.go:86
		_go_fuzz_dep_.CoverTab[16955]++
//line /usr/local/go/src/net/unixsock.go:86
		// _ = "end of CoverTab[16955]"
//line /usr/local/go/src/net/unixsock.go:86
	}
//line /usr/local/go/src/net/unixsock.go:86
	// _ = "end of CoverTab[16951]"
//line /usr/local/go/src/net/unixsock.go:86
	_go_fuzz_dep_.CoverTab[16952]++
						if err := c.fd.closeRead(); err != nil {
//line /usr/local/go/src/net/unixsock.go:87
		_go_fuzz_dep_.CoverTab[16956]++
							return &OpError{Op: "close", Net: c.fd.net, Source: c.fd.laddr, Addr: c.fd.raddr, Err: err}
//line /usr/local/go/src/net/unixsock.go:88
		// _ = "end of CoverTab[16956]"
	} else {
//line /usr/local/go/src/net/unixsock.go:89
		_go_fuzz_dep_.CoverTab[16957]++
//line /usr/local/go/src/net/unixsock.go:89
		// _ = "end of CoverTab[16957]"
//line /usr/local/go/src/net/unixsock.go:89
	}
//line /usr/local/go/src/net/unixsock.go:89
	// _ = "end of CoverTab[16952]"
//line /usr/local/go/src/net/unixsock.go:89
	_go_fuzz_dep_.CoverTab[16953]++
						return nil
//line /usr/local/go/src/net/unixsock.go:90
	// _ = "end of CoverTab[16953]"
}

// CloseWrite shuts down the writing side of the Unix domain connection.
//line /usr/local/go/src/net/unixsock.go:93
// Most callers should just use Close.
//line /usr/local/go/src/net/unixsock.go:95
func (c *UnixConn) CloseWrite() error {
//line /usr/local/go/src/net/unixsock.go:95
	_go_fuzz_dep_.CoverTab[16958]++
						if !c.ok() {
//line /usr/local/go/src/net/unixsock.go:96
		_go_fuzz_dep_.CoverTab[16961]++
							return syscall.EINVAL
//line /usr/local/go/src/net/unixsock.go:97
		// _ = "end of CoverTab[16961]"
	} else {
//line /usr/local/go/src/net/unixsock.go:98
		_go_fuzz_dep_.CoverTab[16962]++
//line /usr/local/go/src/net/unixsock.go:98
		// _ = "end of CoverTab[16962]"
//line /usr/local/go/src/net/unixsock.go:98
	}
//line /usr/local/go/src/net/unixsock.go:98
	// _ = "end of CoverTab[16958]"
//line /usr/local/go/src/net/unixsock.go:98
	_go_fuzz_dep_.CoverTab[16959]++
						if err := c.fd.closeWrite(); err != nil {
//line /usr/local/go/src/net/unixsock.go:99
		_go_fuzz_dep_.CoverTab[16963]++
							return &OpError{Op: "close", Net: c.fd.net, Source: c.fd.laddr, Addr: c.fd.raddr, Err: err}
//line /usr/local/go/src/net/unixsock.go:100
		// _ = "end of CoverTab[16963]"
	} else {
//line /usr/local/go/src/net/unixsock.go:101
		_go_fuzz_dep_.CoverTab[16964]++
//line /usr/local/go/src/net/unixsock.go:101
		// _ = "end of CoverTab[16964]"
//line /usr/local/go/src/net/unixsock.go:101
	}
//line /usr/local/go/src/net/unixsock.go:101
	// _ = "end of CoverTab[16959]"
//line /usr/local/go/src/net/unixsock.go:101
	_go_fuzz_dep_.CoverTab[16960]++
						return nil
//line /usr/local/go/src/net/unixsock.go:102
	// _ = "end of CoverTab[16960]"
}

// ReadFromUnix acts like ReadFrom but returns a UnixAddr.
func (c *UnixConn) ReadFromUnix(b []byte) (int, *UnixAddr, error) {
//line /usr/local/go/src/net/unixsock.go:106
	_go_fuzz_dep_.CoverTab[16965]++
						if !c.ok() {
//line /usr/local/go/src/net/unixsock.go:107
		_go_fuzz_dep_.CoverTab[16968]++
							return 0, nil, syscall.EINVAL
//line /usr/local/go/src/net/unixsock.go:108
		// _ = "end of CoverTab[16968]"
	} else {
//line /usr/local/go/src/net/unixsock.go:109
		_go_fuzz_dep_.CoverTab[16969]++
//line /usr/local/go/src/net/unixsock.go:109
		// _ = "end of CoverTab[16969]"
//line /usr/local/go/src/net/unixsock.go:109
	}
//line /usr/local/go/src/net/unixsock.go:109
	// _ = "end of CoverTab[16965]"
//line /usr/local/go/src/net/unixsock.go:109
	_go_fuzz_dep_.CoverTab[16966]++
						n, addr, err := c.readFrom(b)
						if err != nil {
//line /usr/local/go/src/net/unixsock.go:111
		_go_fuzz_dep_.CoverTab[16970]++
							err = &OpError{Op: "read", Net: c.fd.net, Source: c.fd.laddr, Addr: c.fd.raddr, Err: err}
//line /usr/local/go/src/net/unixsock.go:112
		// _ = "end of CoverTab[16970]"
	} else {
//line /usr/local/go/src/net/unixsock.go:113
		_go_fuzz_dep_.CoverTab[16971]++
//line /usr/local/go/src/net/unixsock.go:113
		// _ = "end of CoverTab[16971]"
//line /usr/local/go/src/net/unixsock.go:113
	}
//line /usr/local/go/src/net/unixsock.go:113
	// _ = "end of CoverTab[16966]"
//line /usr/local/go/src/net/unixsock.go:113
	_go_fuzz_dep_.CoverTab[16967]++
						return n, addr, err
//line /usr/local/go/src/net/unixsock.go:114
	// _ = "end of CoverTab[16967]"
}

// ReadFrom implements the PacketConn ReadFrom method.
func (c *UnixConn) ReadFrom(b []byte) (int, Addr, error) {
//line /usr/local/go/src/net/unixsock.go:118
	_go_fuzz_dep_.CoverTab[16972]++
						if !c.ok() {
//line /usr/local/go/src/net/unixsock.go:119
		_go_fuzz_dep_.CoverTab[16976]++
							return 0, nil, syscall.EINVAL
//line /usr/local/go/src/net/unixsock.go:120
		// _ = "end of CoverTab[16976]"
	} else {
//line /usr/local/go/src/net/unixsock.go:121
		_go_fuzz_dep_.CoverTab[16977]++
//line /usr/local/go/src/net/unixsock.go:121
		// _ = "end of CoverTab[16977]"
//line /usr/local/go/src/net/unixsock.go:121
	}
//line /usr/local/go/src/net/unixsock.go:121
	// _ = "end of CoverTab[16972]"
//line /usr/local/go/src/net/unixsock.go:121
	_go_fuzz_dep_.CoverTab[16973]++
						n, addr, err := c.readFrom(b)
						if err != nil {
//line /usr/local/go/src/net/unixsock.go:123
		_go_fuzz_dep_.CoverTab[16978]++
							err = &OpError{Op: "read", Net: c.fd.net, Source: c.fd.laddr, Addr: c.fd.raddr, Err: err}
//line /usr/local/go/src/net/unixsock.go:124
		// _ = "end of CoverTab[16978]"
	} else {
//line /usr/local/go/src/net/unixsock.go:125
		_go_fuzz_dep_.CoverTab[16979]++
//line /usr/local/go/src/net/unixsock.go:125
		// _ = "end of CoverTab[16979]"
//line /usr/local/go/src/net/unixsock.go:125
	}
//line /usr/local/go/src/net/unixsock.go:125
	// _ = "end of CoverTab[16973]"
//line /usr/local/go/src/net/unixsock.go:125
	_go_fuzz_dep_.CoverTab[16974]++
						if addr == nil {
//line /usr/local/go/src/net/unixsock.go:126
		_go_fuzz_dep_.CoverTab[16980]++
							return n, nil, err
//line /usr/local/go/src/net/unixsock.go:127
		// _ = "end of CoverTab[16980]"
	} else {
//line /usr/local/go/src/net/unixsock.go:128
		_go_fuzz_dep_.CoverTab[16981]++
//line /usr/local/go/src/net/unixsock.go:128
		// _ = "end of CoverTab[16981]"
//line /usr/local/go/src/net/unixsock.go:128
	}
//line /usr/local/go/src/net/unixsock.go:128
	// _ = "end of CoverTab[16974]"
//line /usr/local/go/src/net/unixsock.go:128
	_go_fuzz_dep_.CoverTab[16975]++
						return n, addr, err
//line /usr/local/go/src/net/unixsock.go:129
	// _ = "end of CoverTab[16975]"
}

// ReadMsgUnix reads a message from c, copying the payload into b and
//line /usr/local/go/src/net/unixsock.go:132
// the associated out-of-band data into oob. It returns the number of
//line /usr/local/go/src/net/unixsock.go:132
// bytes copied into b, the number of bytes copied into oob, the flags
//line /usr/local/go/src/net/unixsock.go:132
// that were set on the message and the source address of the message.
//line /usr/local/go/src/net/unixsock.go:132
//
//line /usr/local/go/src/net/unixsock.go:132
// Note that if len(b) == 0 and len(oob) > 0, this function will still
//line /usr/local/go/src/net/unixsock.go:132
// read (and discard) 1 byte from the connection.
//line /usr/local/go/src/net/unixsock.go:139
func (c *UnixConn) ReadMsgUnix(b, oob []byte) (n, oobn, flags int, addr *UnixAddr, err error) {
//line /usr/local/go/src/net/unixsock.go:139
	_go_fuzz_dep_.CoverTab[16982]++
						if !c.ok() {
//line /usr/local/go/src/net/unixsock.go:140
		_go_fuzz_dep_.CoverTab[16985]++
							return 0, 0, 0, nil, syscall.EINVAL
//line /usr/local/go/src/net/unixsock.go:141
		// _ = "end of CoverTab[16985]"
	} else {
//line /usr/local/go/src/net/unixsock.go:142
		_go_fuzz_dep_.CoverTab[16986]++
//line /usr/local/go/src/net/unixsock.go:142
		// _ = "end of CoverTab[16986]"
//line /usr/local/go/src/net/unixsock.go:142
	}
//line /usr/local/go/src/net/unixsock.go:142
	// _ = "end of CoverTab[16982]"
//line /usr/local/go/src/net/unixsock.go:142
	_go_fuzz_dep_.CoverTab[16983]++
						n, oobn, flags, addr, err = c.readMsg(b, oob)
						if err != nil {
//line /usr/local/go/src/net/unixsock.go:144
		_go_fuzz_dep_.CoverTab[16987]++
							err = &OpError{Op: "read", Net: c.fd.net, Source: c.fd.laddr, Addr: c.fd.raddr, Err: err}
//line /usr/local/go/src/net/unixsock.go:145
		// _ = "end of CoverTab[16987]"
	} else {
//line /usr/local/go/src/net/unixsock.go:146
		_go_fuzz_dep_.CoverTab[16988]++
//line /usr/local/go/src/net/unixsock.go:146
		// _ = "end of CoverTab[16988]"
//line /usr/local/go/src/net/unixsock.go:146
	}
//line /usr/local/go/src/net/unixsock.go:146
	// _ = "end of CoverTab[16983]"
//line /usr/local/go/src/net/unixsock.go:146
	_go_fuzz_dep_.CoverTab[16984]++
						return
//line /usr/local/go/src/net/unixsock.go:147
	// _ = "end of CoverTab[16984]"
}

// WriteToUnix acts like WriteTo but takes a UnixAddr.
func (c *UnixConn) WriteToUnix(b []byte, addr *UnixAddr) (int, error) {
//line /usr/local/go/src/net/unixsock.go:151
	_go_fuzz_dep_.CoverTab[16989]++
						if !c.ok() {
//line /usr/local/go/src/net/unixsock.go:152
		_go_fuzz_dep_.CoverTab[16992]++
							return 0, syscall.EINVAL
//line /usr/local/go/src/net/unixsock.go:153
		// _ = "end of CoverTab[16992]"
	} else {
//line /usr/local/go/src/net/unixsock.go:154
		_go_fuzz_dep_.CoverTab[16993]++
//line /usr/local/go/src/net/unixsock.go:154
		// _ = "end of CoverTab[16993]"
//line /usr/local/go/src/net/unixsock.go:154
	}
//line /usr/local/go/src/net/unixsock.go:154
	// _ = "end of CoverTab[16989]"
//line /usr/local/go/src/net/unixsock.go:154
	_go_fuzz_dep_.CoverTab[16990]++
						n, err := c.writeTo(b, addr)
						if err != nil {
//line /usr/local/go/src/net/unixsock.go:156
		_go_fuzz_dep_.CoverTab[16994]++
							err = &OpError{Op: "write", Net: c.fd.net, Source: c.fd.laddr, Addr: addr.opAddr(), Err: err}
//line /usr/local/go/src/net/unixsock.go:157
		// _ = "end of CoverTab[16994]"
	} else {
//line /usr/local/go/src/net/unixsock.go:158
		_go_fuzz_dep_.CoverTab[16995]++
//line /usr/local/go/src/net/unixsock.go:158
		// _ = "end of CoverTab[16995]"
//line /usr/local/go/src/net/unixsock.go:158
	}
//line /usr/local/go/src/net/unixsock.go:158
	// _ = "end of CoverTab[16990]"
//line /usr/local/go/src/net/unixsock.go:158
	_go_fuzz_dep_.CoverTab[16991]++
						return n, err
//line /usr/local/go/src/net/unixsock.go:159
	// _ = "end of CoverTab[16991]"
}

// WriteTo implements the PacketConn WriteTo method.
func (c *UnixConn) WriteTo(b []byte, addr Addr) (int, error) {
//line /usr/local/go/src/net/unixsock.go:163
	_go_fuzz_dep_.CoverTab[16996]++
						if !c.ok() {
//line /usr/local/go/src/net/unixsock.go:164
		_go_fuzz_dep_.CoverTab[17000]++
							return 0, syscall.EINVAL
//line /usr/local/go/src/net/unixsock.go:165
		// _ = "end of CoverTab[17000]"
	} else {
//line /usr/local/go/src/net/unixsock.go:166
		_go_fuzz_dep_.CoverTab[17001]++
//line /usr/local/go/src/net/unixsock.go:166
		// _ = "end of CoverTab[17001]"
//line /usr/local/go/src/net/unixsock.go:166
	}
//line /usr/local/go/src/net/unixsock.go:166
	// _ = "end of CoverTab[16996]"
//line /usr/local/go/src/net/unixsock.go:166
	_go_fuzz_dep_.CoverTab[16997]++
						a, ok := addr.(*UnixAddr)
						if !ok {
//line /usr/local/go/src/net/unixsock.go:168
		_go_fuzz_dep_.CoverTab[17002]++
							return 0, &OpError{Op: "write", Net: c.fd.net, Source: c.fd.laddr, Addr: addr, Err: syscall.EINVAL}
//line /usr/local/go/src/net/unixsock.go:169
		// _ = "end of CoverTab[17002]"
	} else {
//line /usr/local/go/src/net/unixsock.go:170
		_go_fuzz_dep_.CoverTab[17003]++
//line /usr/local/go/src/net/unixsock.go:170
		// _ = "end of CoverTab[17003]"
//line /usr/local/go/src/net/unixsock.go:170
	}
//line /usr/local/go/src/net/unixsock.go:170
	// _ = "end of CoverTab[16997]"
//line /usr/local/go/src/net/unixsock.go:170
	_go_fuzz_dep_.CoverTab[16998]++
						n, err := c.writeTo(b, a)
						if err != nil {
//line /usr/local/go/src/net/unixsock.go:172
		_go_fuzz_dep_.CoverTab[17004]++
							err = &OpError{Op: "write", Net: c.fd.net, Source: c.fd.laddr, Addr: a.opAddr(), Err: err}
//line /usr/local/go/src/net/unixsock.go:173
		// _ = "end of CoverTab[17004]"
	} else {
//line /usr/local/go/src/net/unixsock.go:174
		_go_fuzz_dep_.CoverTab[17005]++
//line /usr/local/go/src/net/unixsock.go:174
		// _ = "end of CoverTab[17005]"
//line /usr/local/go/src/net/unixsock.go:174
	}
//line /usr/local/go/src/net/unixsock.go:174
	// _ = "end of CoverTab[16998]"
//line /usr/local/go/src/net/unixsock.go:174
	_go_fuzz_dep_.CoverTab[16999]++
						return n, err
//line /usr/local/go/src/net/unixsock.go:175
	// _ = "end of CoverTab[16999]"
}

// WriteMsgUnix writes a message to addr via c, copying the payload
//line /usr/local/go/src/net/unixsock.go:178
// from b and the associated out-of-band data from oob. It returns the
//line /usr/local/go/src/net/unixsock.go:178
// number of payload and out-of-band bytes written.
//line /usr/local/go/src/net/unixsock.go:178
//
//line /usr/local/go/src/net/unixsock.go:178
// Note that if len(b) == 0 and len(oob) > 0, this function will still
//line /usr/local/go/src/net/unixsock.go:178
// write 1 byte to the connection.
//line /usr/local/go/src/net/unixsock.go:184
func (c *UnixConn) WriteMsgUnix(b, oob []byte, addr *UnixAddr) (n, oobn int, err error) {
//line /usr/local/go/src/net/unixsock.go:184
	_go_fuzz_dep_.CoverTab[17006]++
						if !c.ok() {
//line /usr/local/go/src/net/unixsock.go:185
		_go_fuzz_dep_.CoverTab[17009]++
							return 0, 0, syscall.EINVAL
//line /usr/local/go/src/net/unixsock.go:186
		// _ = "end of CoverTab[17009]"
	} else {
//line /usr/local/go/src/net/unixsock.go:187
		_go_fuzz_dep_.CoverTab[17010]++
//line /usr/local/go/src/net/unixsock.go:187
		// _ = "end of CoverTab[17010]"
//line /usr/local/go/src/net/unixsock.go:187
	}
//line /usr/local/go/src/net/unixsock.go:187
	// _ = "end of CoverTab[17006]"
//line /usr/local/go/src/net/unixsock.go:187
	_go_fuzz_dep_.CoverTab[17007]++
						n, oobn, err = c.writeMsg(b, oob, addr)
						if err != nil {
//line /usr/local/go/src/net/unixsock.go:189
		_go_fuzz_dep_.CoverTab[17011]++
							err = &OpError{Op: "write", Net: c.fd.net, Source: c.fd.laddr, Addr: addr.opAddr(), Err: err}
//line /usr/local/go/src/net/unixsock.go:190
		// _ = "end of CoverTab[17011]"
	} else {
//line /usr/local/go/src/net/unixsock.go:191
		_go_fuzz_dep_.CoverTab[17012]++
//line /usr/local/go/src/net/unixsock.go:191
		// _ = "end of CoverTab[17012]"
//line /usr/local/go/src/net/unixsock.go:191
	}
//line /usr/local/go/src/net/unixsock.go:191
	// _ = "end of CoverTab[17007]"
//line /usr/local/go/src/net/unixsock.go:191
	_go_fuzz_dep_.CoverTab[17008]++
						return
//line /usr/local/go/src/net/unixsock.go:192
	// _ = "end of CoverTab[17008]"
}

func newUnixConn(fd *netFD) *UnixConn {
//line /usr/local/go/src/net/unixsock.go:195
	_go_fuzz_dep_.CoverTab[17013]++
//line /usr/local/go/src/net/unixsock.go:195
	return &UnixConn{conn{fd}}
//line /usr/local/go/src/net/unixsock.go:195
	// _ = "end of CoverTab[17013]"
//line /usr/local/go/src/net/unixsock.go:195
}

// DialUnix acts like Dial for Unix networks.
//line /usr/local/go/src/net/unixsock.go:197
//
//line /usr/local/go/src/net/unixsock.go:197
// The network must be a Unix network name; see func Dial for details.
//line /usr/local/go/src/net/unixsock.go:197
//
//line /usr/local/go/src/net/unixsock.go:197
// If laddr is non-nil, it is used as the local address for the
//line /usr/local/go/src/net/unixsock.go:197
// connection.
//line /usr/local/go/src/net/unixsock.go:203
func DialUnix(network string, laddr, raddr *UnixAddr) (*UnixConn, error) {
//line /usr/local/go/src/net/unixsock.go:203
	_go_fuzz_dep_.CoverTab[17014]++
						switch network {
	case "unix", "unixgram", "unixpacket":
//line /usr/local/go/src/net/unixsock.go:205
		_go_fuzz_dep_.CoverTab[17017]++
//line /usr/local/go/src/net/unixsock.go:205
		// _ = "end of CoverTab[17017]"
	default:
//line /usr/local/go/src/net/unixsock.go:206
		_go_fuzz_dep_.CoverTab[17018]++
							return nil, &OpError{Op: "dial", Net: network, Source: laddr.opAddr(), Addr: raddr.opAddr(), Err: UnknownNetworkError(network)}
//line /usr/local/go/src/net/unixsock.go:207
		// _ = "end of CoverTab[17018]"
	}
//line /usr/local/go/src/net/unixsock.go:208
	// _ = "end of CoverTab[17014]"
//line /usr/local/go/src/net/unixsock.go:208
	_go_fuzz_dep_.CoverTab[17015]++
						sd := &sysDialer{network: network, address: raddr.String()}
						c, err := sd.dialUnix(context.Background(), laddr, raddr)
						if err != nil {
//line /usr/local/go/src/net/unixsock.go:211
		_go_fuzz_dep_.CoverTab[17019]++
							return nil, &OpError{Op: "dial", Net: network, Source: laddr.opAddr(), Addr: raddr.opAddr(), Err: err}
//line /usr/local/go/src/net/unixsock.go:212
		// _ = "end of CoverTab[17019]"
	} else {
//line /usr/local/go/src/net/unixsock.go:213
		_go_fuzz_dep_.CoverTab[17020]++
//line /usr/local/go/src/net/unixsock.go:213
		// _ = "end of CoverTab[17020]"
//line /usr/local/go/src/net/unixsock.go:213
	}
//line /usr/local/go/src/net/unixsock.go:213
	// _ = "end of CoverTab[17015]"
//line /usr/local/go/src/net/unixsock.go:213
	_go_fuzz_dep_.CoverTab[17016]++
						return c, nil
//line /usr/local/go/src/net/unixsock.go:214
	// _ = "end of CoverTab[17016]"
}

// UnixListener is a Unix domain socket listener. Clients should
//line /usr/local/go/src/net/unixsock.go:217
// typically use variables of type Listener instead of assuming Unix
//line /usr/local/go/src/net/unixsock.go:217
// domain sockets.
//line /usr/local/go/src/net/unixsock.go:220
type UnixListener struct {
	fd		*netFD
	path		string
	unlink		bool
	unlinkOnce	sync.Once
}

func (ln *UnixListener) ok() bool {
//line /usr/local/go/src/net/unixsock.go:227
	_go_fuzz_dep_.CoverTab[17021]++
//line /usr/local/go/src/net/unixsock.go:227
	return ln != nil && func() bool {
//line /usr/local/go/src/net/unixsock.go:227
		_go_fuzz_dep_.CoverTab[17022]++
//line /usr/local/go/src/net/unixsock.go:227
		return ln.fd != nil
//line /usr/local/go/src/net/unixsock.go:227
		// _ = "end of CoverTab[17022]"
//line /usr/local/go/src/net/unixsock.go:227
	}()
//line /usr/local/go/src/net/unixsock.go:227
	// _ = "end of CoverTab[17021]"
//line /usr/local/go/src/net/unixsock.go:227
}

// SyscallConn returns a raw network connection.
//line /usr/local/go/src/net/unixsock.go:229
// This implements the syscall.Conn interface.
//line /usr/local/go/src/net/unixsock.go:229
//
//line /usr/local/go/src/net/unixsock.go:229
// The returned RawConn only supports calling Control. Read and
//line /usr/local/go/src/net/unixsock.go:229
// Write return an error.
//line /usr/local/go/src/net/unixsock.go:234
func (l *UnixListener) SyscallConn() (syscall.RawConn, error) {
//line /usr/local/go/src/net/unixsock.go:234
	_go_fuzz_dep_.CoverTab[17023]++
						if !l.ok() {
//line /usr/local/go/src/net/unixsock.go:235
		_go_fuzz_dep_.CoverTab[17025]++
							return nil, syscall.EINVAL
//line /usr/local/go/src/net/unixsock.go:236
		// _ = "end of CoverTab[17025]"
	} else {
//line /usr/local/go/src/net/unixsock.go:237
		_go_fuzz_dep_.CoverTab[17026]++
//line /usr/local/go/src/net/unixsock.go:237
		// _ = "end of CoverTab[17026]"
//line /usr/local/go/src/net/unixsock.go:237
	}
//line /usr/local/go/src/net/unixsock.go:237
	// _ = "end of CoverTab[17023]"
//line /usr/local/go/src/net/unixsock.go:237
	_go_fuzz_dep_.CoverTab[17024]++
						return newRawListener(l.fd)
//line /usr/local/go/src/net/unixsock.go:238
	// _ = "end of CoverTab[17024]"
}

// AcceptUnix accepts the next incoming call and returns the new
//line /usr/local/go/src/net/unixsock.go:241
// connection.
//line /usr/local/go/src/net/unixsock.go:243
func (l *UnixListener) AcceptUnix() (*UnixConn, error) {
//line /usr/local/go/src/net/unixsock.go:243
	_go_fuzz_dep_.CoverTab[17027]++
						if !l.ok() {
//line /usr/local/go/src/net/unixsock.go:244
		_go_fuzz_dep_.CoverTab[17030]++
							return nil, syscall.EINVAL
//line /usr/local/go/src/net/unixsock.go:245
		// _ = "end of CoverTab[17030]"
	} else {
//line /usr/local/go/src/net/unixsock.go:246
		_go_fuzz_dep_.CoverTab[17031]++
//line /usr/local/go/src/net/unixsock.go:246
		// _ = "end of CoverTab[17031]"
//line /usr/local/go/src/net/unixsock.go:246
	}
//line /usr/local/go/src/net/unixsock.go:246
	// _ = "end of CoverTab[17027]"
//line /usr/local/go/src/net/unixsock.go:246
	_go_fuzz_dep_.CoverTab[17028]++
						c, err := l.accept()
						if err != nil {
//line /usr/local/go/src/net/unixsock.go:248
		_go_fuzz_dep_.CoverTab[17032]++
							return nil, &OpError{Op: "accept", Net: l.fd.net, Source: nil, Addr: l.fd.laddr, Err: err}
//line /usr/local/go/src/net/unixsock.go:249
		// _ = "end of CoverTab[17032]"
	} else {
//line /usr/local/go/src/net/unixsock.go:250
		_go_fuzz_dep_.CoverTab[17033]++
//line /usr/local/go/src/net/unixsock.go:250
		// _ = "end of CoverTab[17033]"
//line /usr/local/go/src/net/unixsock.go:250
	}
//line /usr/local/go/src/net/unixsock.go:250
	// _ = "end of CoverTab[17028]"
//line /usr/local/go/src/net/unixsock.go:250
	_go_fuzz_dep_.CoverTab[17029]++
						return c, nil
//line /usr/local/go/src/net/unixsock.go:251
	// _ = "end of CoverTab[17029]"
}

// Accept implements the Accept method in the Listener interface.
//line /usr/local/go/src/net/unixsock.go:254
// Returned connections will be of type *UnixConn.
//line /usr/local/go/src/net/unixsock.go:256
func (l *UnixListener) Accept() (Conn, error) {
//line /usr/local/go/src/net/unixsock.go:256
	_go_fuzz_dep_.CoverTab[17034]++
						if !l.ok() {
//line /usr/local/go/src/net/unixsock.go:257
		_go_fuzz_dep_.CoverTab[17037]++
							return nil, syscall.EINVAL
//line /usr/local/go/src/net/unixsock.go:258
		// _ = "end of CoverTab[17037]"
	} else {
//line /usr/local/go/src/net/unixsock.go:259
		_go_fuzz_dep_.CoverTab[17038]++
//line /usr/local/go/src/net/unixsock.go:259
		// _ = "end of CoverTab[17038]"
//line /usr/local/go/src/net/unixsock.go:259
	}
//line /usr/local/go/src/net/unixsock.go:259
	// _ = "end of CoverTab[17034]"
//line /usr/local/go/src/net/unixsock.go:259
	_go_fuzz_dep_.CoverTab[17035]++
						c, err := l.accept()
						if err != nil {
//line /usr/local/go/src/net/unixsock.go:261
		_go_fuzz_dep_.CoverTab[17039]++
							return nil, &OpError{Op: "accept", Net: l.fd.net, Source: nil, Addr: l.fd.laddr, Err: err}
//line /usr/local/go/src/net/unixsock.go:262
		// _ = "end of CoverTab[17039]"
	} else {
//line /usr/local/go/src/net/unixsock.go:263
		_go_fuzz_dep_.CoverTab[17040]++
//line /usr/local/go/src/net/unixsock.go:263
		// _ = "end of CoverTab[17040]"
//line /usr/local/go/src/net/unixsock.go:263
	}
//line /usr/local/go/src/net/unixsock.go:263
	// _ = "end of CoverTab[17035]"
//line /usr/local/go/src/net/unixsock.go:263
	_go_fuzz_dep_.CoverTab[17036]++
						return c, nil
//line /usr/local/go/src/net/unixsock.go:264
	// _ = "end of CoverTab[17036]"
}

// Close stops listening on the Unix address. Already accepted
//line /usr/local/go/src/net/unixsock.go:267
// connections are not closed.
//line /usr/local/go/src/net/unixsock.go:269
func (l *UnixListener) Close() error {
//line /usr/local/go/src/net/unixsock.go:269
	_go_fuzz_dep_.CoverTab[17041]++
						if !l.ok() {
//line /usr/local/go/src/net/unixsock.go:270
		_go_fuzz_dep_.CoverTab[17044]++
							return syscall.EINVAL
//line /usr/local/go/src/net/unixsock.go:271
		// _ = "end of CoverTab[17044]"
	} else {
//line /usr/local/go/src/net/unixsock.go:272
		_go_fuzz_dep_.CoverTab[17045]++
//line /usr/local/go/src/net/unixsock.go:272
		// _ = "end of CoverTab[17045]"
//line /usr/local/go/src/net/unixsock.go:272
	}
//line /usr/local/go/src/net/unixsock.go:272
	// _ = "end of CoverTab[17041]"
//line /usr/local/go/src/net/unixsock.go:272
	_go_fuzz_dep_.CoverTab[17042]++
						if err := l.close(); err != nil {
//line /usr/local/go/src/net/unixsock.go:273
		_go_fuzz_dep_.CoverTab[17046]++
							return &OpError{Op: "close", Net: l.fd.net, Source: nil, Addr: l.fd.laddr, Err: err}
//line /usr/local/go/src/net/unixsock.go:274
		// _ = "end of CoverTab[17046]"
	} else {
//line /usr/local/go/src/net/unixsock.go:275
		_go_fuzz_dep_.CoverTab[17047]++
//line /usr/local/go/src/net/unixsock.go:275
		// _ = "end of CoverTab[17047]"
//line /usr/local/go/src/net/unixsock.go:275
	}
//line /usr/local/go/src/net/unixsock.go:275
	// _ = "end of CoverTab[17042]"
//line /usr/local/go/src/net/unixsock.go:275
	_go_fuzz_dep_.CoverTab[17043]++
						return nil
//line /usr/local/go/src/net/unixsock.go:276
	// _ = "end of CoverTab[17043]"
}

// Addr returns the listener's network address.
//line /usr/local/go/src/net/unixsock.go:279
// The Addr returned is shared by all invocations of Addr, so
//line /usr/local/go/src/net/unixsock.go:279
// do not modify it.
//line /usr/local/go/src/net/unixsock.go:282
func (l *UnixListener) Addr() Addr {
//line /usr/local/go/src/net/unixsock.go:282
	_go_fuzz_dep_.CoverTab[17048]++
//line /usr/local/go/src/net/unixsock.go:282
	return l.fd.laddr
//line /usr/local/go/src/net/unixsock.go:282
	// _ = "end of CoverTab[17048]"
//line /usr/local/go/src/net/unixsock.go:282
}

// SetDeadline sets the deadline associated with the listener.
//line /usr/local/go/src/net/unixsock.go:284
// A zero time value disables the deadline.
//line /usr/local/go/src/net/unixsock.go:286
func (l *UnixListener) SetDeadline(t time.Time) error {
//line /usr/local/go/src/net/unixsock.go:286
	_go_fuzz_dep_.CoverTab[17049]++
						if !l.ok() {
//line /usr/local/go/src/net/unixsock.go:287
		_go_fuzz_dep_.CoverTab[17052]++
							return syscall.EINVAL
//line /usr/local/go/src/net/unixsock.go:288
		// _ = "end of CoverTab[17052]"
	} else {
//line /usr/local/go/src/net/unixsock.go:289
		_go_fuzz_dep_.CoverTab[17053]++
//line /usr/local/go/src/net/unixsock.go:289
		// _ = "end of CoverTab[17053]"
//line /usr/local/go/src/net/unixsock.go:289
	}
//line /usr/local/go/src/net/unixsock.go:289
	// _ = "end of CoverTab[17049]"
//line /usr/local/go/src/net/unixsock.go:289
	_go_fuzz_dep_.CoverTab[17050]++
						if err := l.fd.pfd.SetDeadline(t); err != nil {
//line /usr/local/go/src/net/unixsock.go:290
		_go_fuzz_dep_.CoverTab[17054]++
							return &OpError{Op: "set", Net: l.fd.net, Source: nil, Addr: l.fd.laddr, Err: err}
//line /usr/local/go/src/net/unixsock.go:291
		// _ = "end of CoverTab[17054]"
	} else {
//line /usr/local/go/src/net/unixsock.go:292
		_go_fuzz_dep_.CoverTab[17055]++
//line /usr/local/go/src/net/unixsock.go:292
		// _ = "end of CoverTab[17055]"
//line /usr/local/go/src/net/unixsock.go:292
	}
//line /usr/local/go/src/net/unixsock.go:292
	// _ = "end of CoverTab[17050]"
//line /usr/local/go/src/net/unixsock.go:292
	_go_fuzz_dep_.CoverTab[17051]++
						return nil
//line /usr/local/go/src/net/unixsock.go:293
	// _ = "end of CoverTab[17051]"
}

// File returns a copy of the underlying os.File.
//line /usr/local/go/src/net/unixsock.go:296
// It is the caller's responsibility to close f when finished.
//line /usr/local/go/src/net/unixsock.go:296
// Closing l does not affect f, and closing f does not affect l.
//line /usr/local/go/src/net/unixsock.go:296
//
//line /usr/local/go/src/net/unixsock.go:296
// The returned os.File's file descriptor is different from the
//line /usr/local/go/src/net/unixsock.go:296
// connection's. Attempting to change properties of the original
//line /usr/local/go/src/net/unixsock.go:296
// using this duplicate may or may not have the desired effect.
//line /usr/local/go/src/net/unixsock.go:303
func (l *UnixListener) File() (f *os.File, err error) {
//line /usr/local/go/src/net/unixsock.go:303
	_go_fuzz_dep_.CoverTab[17056]++
						if !l.ok() {
//line /usr/local/go/src/net/unixsock.go:304
		_go_fuzz_dep_.CoverTab[17059]++
							return nil, syscall.EINVAL
//line /usr/local/go/src/net/unixsock.go:305
		// _ = "end of CoverTab[17059]"
	} else {
//line /usr/local/go/src/net/unixsock.go:306
		_go_fuzz_dep_.CoverTab[17060]++
//line /usr/local/go/src/net/unixsock.go:306
		// _ = "end of CoverTab[17060]"
//line /usr/local/go/src/net/unixsock.go:306
	}
//line /usr/local/go/src/net/unixsock.go:306
	// _ = "end of CoverTab[17056]"
//line /usr/local/go/src/net/unixsock.go:306
	_go_fuzz_dep_.CoverTab[17057]++
						f, err = l.file()
						if err != nil {
//line /usr/local/go/src/net/unixsock.go:308
		_go_fuzz_dep_.CoverTab[17061]++
							err = &OpError{Op: "file", Net: l.fd.net, Source: nil, Addr: l.fd.laddr, Err: err}
//line /usr/local/go/src/net/unixsock.go:309
		// _ = "end of CoverTab[17061]"
	} else {
//line /usr/local/go/src/net/unixsock.go:310
		_go_fuzz_dep_.CoverTab[17062]++
//line /usr/local/go/src/net/unixsock.go:310
		// _ = "end of CoverTab[17062]"
//line /usr/local/go/src/net/unixsock.go:310
	}
//line /usr/local/go/src/net/unixsock.go:310
	// _ = "end of CoverTab[17057]"
//line /usr/local/go/src/net/unixsock.go:310
	_go_fuzz_dep_.CoverTab[17058]++
						return
//line /usr/local/go/src/net/unixsock.go:311
	// _ = "end of CoverTab[17058]"
}

// ListenUnix acts like Listen for Unix networks.
//line /usr/local/go/src/net/unixsock.go:314
//
//line /usr/local/go/src/net/unixsock.go:314
// The network must be "unix" or "unixpacket".
//line /usr/local/go/src/net/unixsock.go:317
func ListenUnix(network string, laddr *UnixAddr) (*UnixListener, error) {
//line /usr/local/go/src/net/unixsock.go:317
	_go_fuzz_dep_.CoverTab[17063]++
						switch network {
	case "unix", "unixpacket":
//line /usr/local/go/src/net/unixsock.go:319
		_go_fuzz_dep_.CoverTab[17067]++
//line /usr/local/go/src/net/unixsock.go:319
		// _ = "end of CoverTab[17067]"
	default:
//line /usr/local/go/src/net/unixsock.go:320
		_go_fuzz_dep_.CoverTab[17068]++
							return nil, &OpError{Op: "listen", Net: network, Source: nil, Addr: laddr.opAddr(), Err: UnknownNetworkError(network)}
//line /usr/local/go/src/net/unixsock.go:321
		// _ = "end of CoverTab[17068]"
	}
//line /usr/local/go/src/net/unixsock.go:322
	// _ = "end of CoverTab[17063]"
//line /usr/local/go/src/net/unixsock.go:322
	_go_fuzz_dep_.CoverTab[17064]++
						if laddr == nil {
//line /usr/local/go/src/net/unixsock.go:323
		_go_fuzz_dep_.CoverTab[17069]++
							return nil, &OpError{Op: "listen", Net: network, Source: nil, Addr: laddr.opAddr(), Err: errMissingAddress}
//line /usr/local/go/src/net/unixsock.go:324
		// _ = "end of CoverTab[17069]"
	} else {
//line /usr/local/go/src/net/unixsock.go:325
		_go_fuzz_dep_.CoverTab[17070]++
//line /usr/local/go/src/net/unixsock.go:325
		// _ = "end of CoverTab[17070]"
//line /usr/local/go/src/net/unixsock.go:325
	}
//line /usr/local/go/src/net/unixsock.go:325
	// _ = "end of CoverTab[17064]"
//line /usr/local/go/src/net/unixsock.go:325
	_go_fuzz_dep_.CoverTab[17065]++
						sl := &sysListener{network: network, address: laddr.String()}
						ln, err := sl.listenUnix(context.Background(), laddr)
						if err != nil {
//line /usr/local/go/src/net/unixsock.go:328
		_go_fuzz_dep_.CoverTab[17071]++
							return nil, &OpError{Op: "listen", Net: network, Source: nil, Addr: laddr.opAddr(), Err: err}
//line /usr/local/go/src/net/unixsock.go:329
		// _ = "end of CoverTab[17071]"
	} else {
//line /usr/local/go/src/net/unixsock.go:330
		_go_fuzz_dep_.CoverTab[17072]++
//line /usr/local/go/src/net/unixsock.go:330
		// _ = "end of CoverTab[17072]"
//line /usr/local/go/src/net/unixsock.go:330
	}
//line /usr/local/go/src/net/unixsock.go:330
	// _ = "end of CoverTab[17065]"
//line /usr/local/go/src/net/unixsock.go:330
	_go_fuzz_dep_.CoverTab[17066]++
						return ln, nil
//line /usr/local/go/src/net/unixsock.go:331
	// _ = "end of CoverTab[17066]"
}

// ListenUnixgram acts like ListenPacket for Unix networks.
//line /usr/local/go/src/net/unixsock.go:334
//
//line /usr/local/go/src/net/unixsock.go:334
// The network must be "unixgram".
//line /usr/local/go/src/net/unixsock.go:337
func ListenUnixgram(network string, laddr *UnixAddr) (*UnixConn, error) {
//line /usr/local/go/src/net/unixsock.go:337
	_go_fuzz_dep_.CoverTab[17073]++
						switch network {
	case "unixgram":
//line /usr/local/go/src/net/unixsock.go:339
		_go_fuzz_dep_.CoverTab[17077]++
//line /usr/local/go/src/net/unixsock.go:339
		// _ = "end of CoverTab[17077]"
	default:
//line /usr/local/go/src/net/unixsock.go:340
		_go_fuzz_dep_.CoverTab[17078]++
							return nil, &OpError{Op: "listen", Net: network, Source: nil, Addr: laddr.opAddr(), Err: UnknownNetworkError(network)}
//line /usr/local/go/src/net/unixsock.go:341
		// _ = "end of CoverTab[17078]"
	}
//line /usr/local/go/src/net/unixsock.go:342
	// _ = "end of CoverTab[17073]"
//line /usr/local/go/src/net/unixsock.go:342
	_go_fuzz_dep_.CoverTab[17074]++
						if laddr == nil {
//line /usr/local/go/src/net/unixsock.go:343
		_go_fuzz_dep_.CoverTab[17079]++
							return nil, &OpError{Op: "listen", Net: network, Source: nil, Addr: nil, Err: errMissingAddress}
//line /usr/local/go/src/net/unixsock.go:344
		// _ = "end of CoverTab[17079]"
	} else {
//line /usr/local/go/src/net/unixsock.go:345
		_go_fuzz_dep_.CoverTab[17080]++
//line /usr/local/go/src/net/unixsock.go:345
		// _ = "end of CoverTab[17080]"
//line /usr/local/go/src/net/unixsock.go:345
	}
//line /usr/local/go/src/net/unixsock.go:345
	// _ = "end of CoverTab[17074]"
//line /usr/local/go/src/net/unixsock.go:345
	_go_fuzz_dep_.CoverTab[17075]++
						sl := &sysListener{network: network, address: laddr.String()}
						c, err := sl.listenUnixgram(context.Background(), laddr)
						if err != nil {
//line /usr/local/go/src/net/unixsock.go:348
		_go_fuzz_dep_.CoverTab[17081]++
							return nil, &OpError{Op: "listen", Net: network, Source: nil, Addr: laddr.opAddr(), Err: err}
//line /usr/local/go/src/net/unixsock.go:349
		// _ = "end of CoverTab[17081]"
	} else {
//line /usr/local/go/src/net/unixsock.go:350
		_go_fuzz_dep_.CoverTab[17082]++
//line /usr/local/go/src/net/unixsock.go:350
		// _ = "end of CoverTab[17082]"
//line /usr/local/go/src/net/unixsock.go:350
	}
//line /usr/local/go/src/net/unixsock.go:350
	// _ = "end of CoverTab[17075]"
//line /usr/local/go/src/net/unixsock.go:350
	_go_fuzz_dep_.CoverTab[17076]++
						return c, nil
//line /usr/local/go/src/net/unixsock.go:351
	// _ = "end of CoverTab[17076]"
}

//line /usr/local/go/src/net/unixsock.go:352
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/unixsock.go:352
var _ = _go_fuzz_dep_.CoverTab
