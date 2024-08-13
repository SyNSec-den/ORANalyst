// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /snap/go/10455/src/net/unixsock.go:5
package net

//line /snap/go/10455/src/net/unixsock.go:5
import (
//line /snap/go/10455/src/net/unixsock.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /snap/go/10455/src/net/unixsock.go:5
)
//line /snap/go/10455/src/net/unixsock.go:5
import (
//line /snap/go/10455/src/net/unixsock.go:5
	_atomic_ "sync/atomic"
//line /snap/go/10455/src/net/unixsock.go:5
)

import (
	"context"
	"os"
	"sync"
	"syscall"
	"time"
)

//line /snap/go/10455/src/net/unixsock.go:21
// UnixAddr represents the address of a Unix domain socket end point.
type UnixAddr struct {
	Name	string
	Net	string
}

// Network returns the address's network name, "unix", "unixgram" or
//line /snap/go/10455/src/net/unixsock.go:27
// "unixpacket".
//line /snap/go/10455/src/net/unixsock.go:29
func (a *UnixAddr) Network() string {
//line /snap/go/10455/src/net/unixsock.go:29
	_go_fuzz_dep_.CoverTab[8843]++
						return a.Net
//line /snap/go/10455/src/net/unixsock.go:30
	// _ = "end of CoverTab[8843]"
}

func (a *UnixAddr) String() string {
//line /snap/go/10455/src/net/unixsock.go:33
	_go_fuzz_dep_.CoverTab[8844]++
						if a == nil {
//line /snap/go/10455/src/net/unixsock.go:34
		_go_fuzz_dep_.CoverTab[530113]++
//line /snap/go/10455/src/net/unixsock.go:34
		_go_fuzz_dep_.CoverTab[8846]++
							return "<nil>"
//line /snap/go/10455/src/net/unixsock.go:35
		// _ = "end of CoverTab[8846]"
	} else {
//line /snap/go/10455/src/net/unixsock.go:36
		_go_fuzz_dep_.CoverTab[530114]++
//line /snap/go/10455/src/net/unixsock.go:36
		_go_fuzz_dep_.CoverTab[8847]++
//line /snap/go/10455/src/net/unixsock.go:36
		// _ = "end of CoverTab[8847]"
//line /snap/go/10455/src/net/unixsock.go:36
	}
//line /snap/go/10455/src/net/unixsock.go:36
	// _ = "end of CoverTab[8844]"
//line /snap/go/10455/src/net/unixsock.go:36
	_go_fuzz_dep_.CoverTab[8845]++
						return a.Name
//line /snap/go/10455/src/net/unixsock.go:37
	// _ = "end of CoverTab[8845]"
}

func (a *UnixAddr) isWildcard() bool {
//line /snap/go/10455/src/net/unixsock.go:40
	_go_fuzz_dep_.CoverTab[8848]++
						return a == nil || func() bool {
//line /snap/go/10455/src/net/unixsock.go:41
		_go_fuzz_dep_.CoverTab[8849]++
//line /snap/go/10455/src/net/unixsock.go:41
		return a.Name == ""
//line /snap/go/10455/src/net/unixsock.go:41
		// _ = "end of CoverTab[8849]"
//line /snap/go/10455/src/net/unixsock.go:41
	}()
//line /snap/go/10455/src/net/unixsock.go:41
	// _ = "end of CoverTab[8848]"
}

func (a *UnixAddr) opAddr() Addr {
//line /snap/go/10455/src/net/unixsock.go:44
	_go_fuzz_dep_.CoverTab[8850]++
						if a == nil {
//line /snap/go/10455/src/net/unixsock.go:45
		_go_fuzz_dep_.CoverTab[530115]++
//line /snap/go/10455/src/net/unixsock.go:45
		_go_fuzz_dep_.CoverTab[8852]++
							return nil
//line /snap/go/10455/src/net/unixsock.go:46
		// _ = "end of CoverTab[8852]"
	} else {
//line /snap/go/10455/src/net/unixsock.go:47
		_go_fuzz_dep_.CoverTab[530116]++
//line /snap/go/10455/src/net/unixsock.go:47
		_go_fuzz_dep_.CoverTab[8853]++
//line /snap/go/10455/src/net/unixsock.go:47
		// _ = "end of CoverTab[8853]"
//line /snap/go/10455/src/net/unixsock.go:47
	}
//line /snap/go/10455/src/net/unixsock.go:47
	// _ = "end of CoverTab[8850]"
//line /snap/go/10455/src/net/unixsock.go:47
	_go_fuzz_dep_.CoverTab[8851]++
						return a
//line /snap/go/10455/src/net/unixsock.go:48
	// _ = "end of CoverTab[8851]"
}

// ResolveUnixAddr returns an address of Unix domain socket end point.
//line /snap/go/10455/src/net/unixsock.go:51
//
//line /snap/go/10455/src/net/unixsock.go:51
// The network must be a Unix network name.
//line /snap/go/10455/src/net/unixsock.go:51
//
//line /snap/go/10455/src/net/unixsock.go:51
// See func Dial for a description of the network and address
//line /snap/go/10455/src/net/unixsock.go:51
// parameters.
//line /snap/go/10455/src/net/unixsock.go:57
func ResolveUnixAddr(network, address string) (*UnixAddr, error) {
//line /snap/go/10455/src/net/unixsock.go:57
	_go_fuzz_dep_.CoverTab[8854]++
						switch network {
	case "unix", "unixgram", "unixpacket":
//line /snap/go/10455/src/net/unixsock.go:59
		_go_fuzz_dep_.CoverTab[530117]++
//line /snap/go/10455/src/net/unixsock.go:59
		_go_fuzz_dep_.CoverTab[8855]++
							return &UnixAddr{Name: address, Net: network}, nil
//line /snap/go/10455/src/net/unixsock.go:60
		// _ = "end of CoverTab[8855]"
	default:
//line /snap/go/10455/src/net/unixsock.go:61
		_go_fuzz_dep_.CoverTab[530118]++
//line /snap/go/10455/src/net/unixsock.go:61
		_go_fuzz_dep_.CoverTab[8856]++
							return nil, UnknownNetworkError(network)
//line /snap/go/10455/src/net/unixsock.go:62
		// _ = "end of CoverTab[8856]"
	}
//line /snap/go/10455/src/net/unixsock.go:63
	// _ = "end of CoverTab[8854]"
}

// UnixConn is an implementation of the Conn interface for connections
//line /snap/go/10455/src/net/unixsock.go:66
// to Unix domain sockets.
//line /snap/go/10455/src/net/unixsock.go:68
type UnixConn struct {
	conn
}

// SyscallConn returns a raw network connection.
//line /snap/go/10455/src/net/unixsock.go:72
// This implements the syscall.Conn interface.
//line /snap/go/10455/src/net/unixsock.go:74
func (c *UnixConn) SyscallConn() (syscall.RawConn, error) {
//line /snap/go/10455/src/net/unixsock.go:74
	_go_fuzz_dep_.CoverTab[8857]++
						if !c.ok() {
//line /snap/go/10455/src/net/unixsock.go:75
		_go_fuzz_dep_.CoverTab[530119]++
//line /snap/go/10455/src/net/unixsock.go:75
		_go_fuzz_dep_.CoverTab[8859]++
							return nil, syscall.EINVAL
//line /snap/go/10455/src/net/unixsock.go:76
		// _ = "end of CoverTab[8859]"
	} else {
//line /snap/go/10455/src/net/unixsock.go:77
		_go_fuzz_dep_.CoverTab[530120]++
//line /snap/go/10455/src/net/unixsock.go:77
		_go_fuzz_dep_.CoverTab[8860]++
//line /snap/go/10455/src/net/unixsock.go:77
		// _ = "end of CoverTab[8860]"
//line /snap/go/10455/src/net/unixsock.go:77
	}
//line /snap/go/10455/src/net/unixsock.go:77
	// _ = "end of CoverTab[8857]"
//line /snap/go/10455/src/net/unixsock.go:77
	_go_fuzz_dep_.CoverTab[8858]++
						return newRawConn(c.fd)
//line /snap/go/10455/src/net/unixsock.go:78
	// _ = "end of CoverTab[8858]"
}

// CloseRead shuts down the reading side of the Unix domain connection.
//line /snap/go/10455/src/net/unixsock.go:81
// Most callers should just use Close.
//line /snap/go/10455/src/net/unixsock.go:83
func (c *UnixConn) CloseRead() error {
//line /snap/go/10455/src/net/unixsock.go:83
	_go_fuzz_dep_.CoverTab[8861]++
						if !c.ok() {
//line /snap/go/10455/src/net/unixsock.go:84
		_go_fuzz_dep_.CoverTab[530121]++
//line /snap/go/10455/src/net/unixsock.go:84
		_go_fuzz_dep_.CoverTab[8864]++
							return syscall.EINVAL
//line /snap/go/10455/src/net/unixsock.go:85
		// _ = "end of CoverTab[8864]"
	} else {
//line /snap/go/10455/src/net/unixsock.go:86
		_go_fuzz_dep_.CoverTab[530122]++
//line /snap/go/10455/src/net/unixsock.go:86
		_go_fuzz_dep_.CoverTab[8865]++
//line /snap/go/10455/src/net/unixsock.go:86
		// _ = "end of CoverTab[8865]"
//line /snap/go/10455/src/net/unixsock.go:86
	}
//line /snap/go/10455/src/net/unixsock.go:86
	// _ = "end of CoverTab[8861]"
//line /snap/go/10455/src/net/unixsock.go:86
	_go_fuzz_dep_.CoverTab[8862]++
						if err := c.fd.closeRead(); err != nil {
//line /snap/go/10455/src/net/unixsock.go:87
		_go_fuzz_dep_.CoverTab[530123]++
//line /snap/go/10455/src/net/unixsock.go:87
		_go_fuzz_dep_.CoverTab[8866]++
							return &OpError{Op: "close", Net: c.fd.net, Source: c.fd.laddr, Addr: c.fd.raddr, Err: err}
//line /snap/go/10455/src/net/unixsock.go:88
		// _ = "end of CoverTab[8866]"
	} else {
//line /snap/go/10455/src/net/unixsock.go:89
		_go_fuzz_dep_.CoverTab[530124]++
//line /snap/go/10455/src/net/unixsock.go:89
		_go_fuzz_dep_.CoverTab[8867]++
//line /snap/go/10455/src/net/unixsock.go:89
		// _ = "end of CoverTab[8867]"
//line /snap/go/10455/src/net/unixsock.go:89
	}
//line /snap/go/10455/src/net/unixsock.go:89
	// _ = "end of CoverTab[8862]"
//line /snap/go/10455/src/net/unixsock.go:89
	_go_fuzz_dep_.CoverTab[8863]++
						return nil
//line /snap/go/10455/src/net/unixsock.go:90
	// _ = "end of CoverTab[8863]"
}

// CloseWrite shuts down the writing side of the Unix domain connection.
//line /snap/go/10455/src/net/unixsock.go:93
// Most callers should just use Close.
//line /snap/go/10455/src/net/unixsock.go:95
func (c *UnixConn) CloseWrite() error {
//line /snap/go/10455/src/net/unixsock.go:95
	_go_fuzz_dep_.CoverTab[8868]++
						if !c.ok() {
//line /snap/go/10455/src/net/unixsock.go:96
		_go_fuzz_dep_.CoverTab[530125]++
//line /snap/go/10455/src/net/unixsock.go:96
		_go_fuzz_dep_.CoverTab[8871]++
							return syscall.EINVAL
//line /snap/go/10455/src/net/unixsock.go:97
		// _ = "end of CoverTab[8871]"
	} else {
//line /snap/go/10455/src/net/unixsock.go:98
		_go_fuzz_dep_.CoverTab[530126]++
//line /snap/go/10455/src/net/unixsock.go:98
		_go_fuzz_dep_.CoverTab[8872]++
//line /snap/go/10455/src/net/unixsock.go:98
		// _ = "end of CoverTab[8872]"
//line /snap/go/10455/src/net/unixsock.go:98
	}
//line /snap/go/10455/src/net/unixsock.go:98
	// _ = "end of CoverTab[8868]"
//line /snap/go/10455/src/net/unixsock.go:98
	_go_fuzz_dep_.CoverTab[8869]++
						if err := c.fd.closeWrite(); err != nil {
//line /snap/go/10455/src/net/unixsock.go:99
		_go_fuzz_dep_.CoverTab[530127]++
//line /snap/go/10455/src/net/unixsock.go:99
		_go_fuzz_dep_.CoverTab[8873]++
							return &OpError{Op: "close", Net: c.fd.net, Source: c.fd.laddr, Addr: c.fd.raddr, Err: err}
//line /snap/go/10455/src/net/unixsock.go:100
		// _ = "end of CoverTab[8873]"
	} else {
//line /snap/go/10455/src/net/unixsock.go:101
		_go_fuzz_dep_.CoverTab[530128]++
//line /snap/go/10455/src/net/unixsock.go:101
		_go_fuzz_dep_.CoverTab[8874]++
//line /snap/go/10455/src/net/unixsock.go:101
		// _ = "end of CoverTab[8874]"
//line /snap/go/10455/src/net/unixsock.go:101
	}
//line /snap/go/10455/src/net/unixsock.go:101
	// _ = "end of CoverTab[8869]"
//line /snap/go/10455/src/net/unixsock.go:101
	_go_fuzz_dep_.CoverTab[8870]++
						return nil
//line /snap/go/10455/src/net/unixsock.go:102
	// _ = "end of CoverTab[8870]"
}

// ReadFromUnix acts like ReadFrom but returns a UnixAddr.
func (c *UnixConn) ReadFromUnix(b []byte) (int, *UnixAddr, error) {
//line /snap/go/10455/src/net/unixsock.go:106
	_go_fuzz_dep_.CoverTab[8875]++
						if !c.ok() {
//line /snap/go/10455/src/net/unixsock.go:107
		_go_fuzz_dep_.CoverTab[530129]++
//line /snap/go/10455/src/net/unixsock.go:107
		_go_fuzz_dep_.CoverTab[8878]++
							return 0, nil, syscall.EINVAL
//line /snap/go/10455/src/net/unixsock.go:108
		// _ = "end of CoverTab[8878]"
	} else {
//line /snap/go/10455/src/net/unixsock.go:109
		_go_fuzz_dep_.CoverTab[530130]++
//line /snap/go/10455/src/net/unixsock.go:109
		_go_fuzz_dep_.CoverTab[8879]++
//line /snap/go/10455/src/net/unixsock.go:109
		// _ = "end of CoverTab[8879]"
//line /snap/go/10455/src/net/unixsock.go:109
	}
//line /snap/go/10455/src/net/unixsock.go:109
	// _ = "end of CoverTab[8875]"
//line /snap/go/10455/src/net/unixsock.go:109
	_go_fuzz_dep_.CoverTab[8876]++
						n, addr, err := c.readFrom(b)
						if err != nil {
//line /snap/go/10455/src/net/unixsock.go:111
		_go_fuzz_dep_.CoverTab[530131]++
//line /snap/go/10455/src/net/unixsock.go:111
		_go_fuzz_dep_.CoverTab[8880]++
							err = &OpError{Op: "read", Net: c.fd.net, Source: c.fd.laddr, Addr: c.fd.raddr, Err: err}
//line /snap/go/10455/src/net/unixsock.go:112
		// _ = "end of CoverTab[8880]"
	} else {
//line /snap/go/10455/src/net/unixsock.go:113
		_go_fuzz_dep_.CoverTab[530132]++
//line /snap/go/10455/src/net/unixsock.go:113
		_go_fuzz_dep_.CoverTab[8881]++
//line /snap/go/10455/src/net/unixsock.go:113
		// _ = "end of CoverTab[8881]"
//line /snap/go/10455/src/net/unixsock.go:113
	}
//line /snap/go/10455/src/net/unixsock.go:113
	// _ = "end of CoverTab[8876]"
//line /snap/go/10455/src/net/unixsock.go:113
	_go_fuzz_dep_.CoverTab[8877]++
						return n, addr, err
//line /snap/go/10455/src/net/unixsock.go:114
	// _ = "end of CoverTab[8877]"
}

// ReadFrom implements the PacketConn ReadFrom method.
func (c *UnixConn) ReadFrom(b []byte) (int, Addr, error) {
//line /snap/go/10455/src/net/unixsock.go:118
	_go_fuzz_dep_.CoverTab[8882]++
						if !c.ok() {
//line /snap/go/10455/src/net/unixsock.go:119
		_go_fuzz_dep_.CoverTab[530133]++
//line /snap/go/10455/src/net/unixsock.go:119
		_go_fuzz_dep_.CoverTab[8886]++
							return 0, nil, syscall.EINVAL
//line /snap/go/10455/src/net/unixsock.go:120
		// _ = "end of CoverTab[8886]"
	} else {
//line /snap/go/10455/src/net/unixsock.go:121
		_go_fuzz_dep_.CoverTab[530134]++
//line /snap/go/10455/src/net/unixsock.go:121
		_go_fuzz_dep_.CoverTab[8887]++
//line /snap/go/10455/src/net/unixsock.go:121
		// _ = "end of CoverTab[8887]"
//line /snap/go/10455/src/net/unixsock.go:121
	}
//line /snap/go/10455/src/net/unixsock.go:121
	// _ = "end of CoverTab[8882]"
//line /snap/go/10455/src/net/unixsock.go:121
	_go_fuzz_dep_.CoverTab[8883]++
						n, addr, err := c.readFrom(b)
						if err != nil {
//line /snap/go/10455/src/net/unixsock.go:123
		_go_fuzz_dep_.CoverTab[530135]++
//line /snap/go/10455/src/net/unixsock.go:123
		_go_fuzz_dep_.CoverTab[8888]++
							err = &OpError{Op: "read", Net: c.fd.net, Source: c.fd.laddr, Addr: c.fd.raddr, Err: err}
//line /snap/go/10455/src/net/unixsock.go:124
		// _ = "end of CoverTab[8888]"
	} else {
//line /snap/go/10455/src/net/unixsock.go:125
		_go_fuzz_dep_.CoverTab[530136]++
//line /snap/go/10455/src/net/unixsock.go:125
		_go_fuzz_dep_.CoverTab[8889]++
//line /snap/go/10455/src/net/unixsock.go:125
		// _ = "end of CoverTab[8889]"
//line /snap/go/10455/src/net/unixsock.go:125
	}
//line /snap/go/10455/src/net/unixsock.go:125
	// _ = "end of CoverTab[8883]"
//line /snap/go/10455/src/net/unixsock.go:125
	_go_fuzz_dep_.CoverTab[8884]++
						if addr == nil {
//line /snap/go/10455/src/net/unixsock.go:126
		_go_fuzz_dep_.CoverTab[530137]++
//line /snap/go/10455/src/net/unixsock.go:126
		_go_fuzz_dep_.CoverTab[8890]++
							return n, nil, err
//line /snap/go/10455/src/net/unixsock.go:127
		// _ = "end of CoverTab[8890]"
	} else {
//line /snap/go/10455/src/net/unixsock.go:128
		_go_fuzz_dep_.CoverTab[530138]++
//line /snap/go/10455/src/net/unixsock.go:128
		_go_fuzz_dep_.CoverTab[8891]++
//line /snap/go/10455/src/net/unixsock.go:128
		// _ = "end of CoverTab[8891]"
//line /snap/go/10455/src/net/unixsock.go:128
	}
//line /snap/go/10455/src/net/unixsock.go:128
	// _ = "end of CoverTab[8884]"
//line /snap/go/10455/src/net/unixsock.go:128
	_go_fuzz_dep_.CoverTab[8885]++
						return n, addr, err
//line /snap/go/10455/src/net/unixsock.go:129
	// _ = "end of CoverTab[8885]"
}

// ReadMsgUnix reads a message from c, copying the payload into b and
//line /snap/go/10455/src/net/unixsock.go:132
// the associated out-of-band data into oob. It returns the number of
//line /snap/go/10455/src/net/unixsock.go:132
// bytes copied into b, the number of bytes copied into oob, the flags
//line /snap/go/10455/src/net/unixsock.go:132
// that were set on the message and the source address of the message.
//line /snap/go/10455/src/net/unixsock.go:132
//
//line /snap/go/10455/src/net/unixsock.go:132
// Note that if len(b) == 0 and len(oob) > 0, this function will still
//line /snap/go/10455/src/net/unixsock.go:132
// read (and discard) 1 byte from the connection.
//line /snap/go/10455/src/net/unixsock.go:139
func (c *UnixConn) ReadMsgUnix(b, oob []byte) (n, oobn, flags int, addr *UnixAddr, err error) {
//line /snap/go/10455/src/net/unixsock.go:139
	_go_fuzz_dep_.CoverTab[8892]++
						if !c.ok() {
//line /snap/go/10455/src/net/unixsock.go:140
		_go_fuzz_dep_.CoverTab[530139]++
//line /snap/go/10455/src/net/unixsock.go:140
		_go_fuzz_dep_.CoverTab[8895]++
							return 0, 0, 0, nil, syscall.EINVAL
//line /snap/go/10455/src/net/unixsock.go:141
		// _ = "end of CoverTab[8895]"
	} else {
//line /snap/go/10455/src/net/unixsock.go:142
		_go_fuzz_dep_.CoverTab[530140]++
//line /snap/go/10455/src/net/unixsock.go:142
		_go_fuzz_dep_.CoverTab[8896]++
//line /snap/go/10455/src/net/unixsock.go:142
		// _ = "end of CoverTab[8896]"
//line /snap/go/10455/src/net/unixsock.go:142
	}
//line /snap/go/10455/src/net/unixsock.go:142
	// _ = "end of CoverTab[8892]"
//line /snap/go/10455/src/net/unixsock.go:142
	_go_fuzz_dep_.CoverTab[8893]++
						n, oobn, flags, addr, err = c.readMsg(b, oob)
						if err != nil {
//line /snap/go/10455/src/net/unixsock.go:144
		_go_fuzz_dep_.CoverTab[530141]++
//line /snap/go/10455/src/net/unixsock.go:144
		_go_fuzz_dep_.CoverTab[8897]++
							err = &OpError{Op: "read", Net: c.fd.net, Source: c.fd.laddr, Addr: c.fd.raddr, Err: err}
//line /snap/go/10455/src/net/unixsock.go:145
		// _ = "end of CoverTab[8897]"
	} else {
//line /snap/go/10455/src/net/unixsock.go:146
		_go_fuzz_dep_.CoverTab[530142]++
//line /snap/go/10455/src/net/unixsock.go:146
		_go_fuzz_dep_.CoverTab[8898]++
//line /snap/go/10455/src/net/unixsock.go:146
		// _ = "end of CoverTab[8898]"
//line /snap/go/10455/src/net/unixsock.go:146
	}
//line /snap/go/10455/src/net/unixsock.go:146
	// _ = "end of CoverTab[8893]"
//line /snap/go/10455/src/net/unixsock.go:146
	_go_fuzz_dep_.CoverTab[8894]++
						return
//line /snap/go/10455/src/net/unixsock.go:147
	// _ = "end of CoverTab[8894]"
}

// WriteToUnix acts like WriteTo but takes a UnixAddr.
func (c *UnixConn) WriteToUnix(b []byte, addr *UnixAddr) (int, error) {
//line /snap/go/10455/src/net/unixsock.go:151
	_go_fuzz_dep_.CoverTab[8899]++
						if !c.ok() {
//line /snap/go/10455/src/net/unixsock.go:152
		_go_fuzz_dep_.CoverTab[530143]++
//line /snap/go/10455/src/net/unixsock.go:152
		_go_fuzz_dep_.CoverTab[8902]++
							return 0, syscall.EINVAL
//line /snap/go/10455/src/net/unixsock.go:153
		// _ = "end of CoverTab[8902]"
	} else {
//line /snap/go/10455/src/net/unixsock.go:154
		_go_fuzz_dep_.CoverTab[530144]++
//line /snap/go/10455/src/net/unixsock.go:154
		_go_fuzz_dep_.CoverTab[8903]++
//line /snap/go/10455/src/net/unixsock.go:154
		// _ = "end of CoverTab[8903]"
//line /snap/go/10455/src/net/unixsock.go:154
	}
//line /snap/go/10455/src/net/unixsock.go:154
	// _ = "end of CoverTab[8899]"
//line /snap/go/10455/src/net/unixsock.go:154
	_go_fuzz_dep_.CoverTab[8900]++
						n, err := c.writeTo(b, addr)
						if err != nil {
//line /snap/go/10455/src/net/unixsock.go:156
		_go_fuzz_dep_.CoverTab[530145]++
//line /snap/go/10455/src/net/unixsock.go:156
		_go_fuzz_dep_.CoverTab[8904]++
							err = &OpError{Op: "write", Net: c.fd.net, Source: c.fd.laddr, Addr: addr.opAddr(), Err: err}
//line /snap/go/10455/src/net/unixsock.go:157
		// _ = "end of CoverTab[8904]"
	} else {
//line /snap/go/10455/src/net/unixsock.go:158
		_go_fuzz_dep_.CoverTab[530146]++
//line /snap/go/10455/src/net/unixsock.go:158
		_go_fuzz_dep_.CoverTab[8905]++
//line /snap/go/10455/src/net/unixsock.go:158
		// _ = "end of CoverTab[8905]"
//line /snap/go/10455/src/net/unixsock.go:158
	}
//line /snap/go/10455/src/net/unixsock.go:158
	// _ = "end of CoverTab[8900]"
//line /snap/go/10455/src/net/unixsock.go:158
	_go_fuzz_dep_.CoverTab[8901]++
						return n, err
//line /snap/go/10455/src/net/unixsock.go:159
	// _ = "end of CoverTab[8901]"
}

// WriteTo implements the PacketConn WriteTo method.
func (c *UnixConn) WriteTo(b []byte, addr Addr) (int, error) {
//line /snap/go/10455/src/net/unixsock.go:163
	_go_fuzz_dep_.CoverTab[8906]++
						if !c.ok() {
//line /snap/go/10455/src/net/unixsock.go:164
		_go_fuzz_dep_.CoverTab[530147]++
//line /snap/go/10455/src/net/unixsock.go:164
		_go_fuzz_dep_.CoverTab[8910]++
							return 0, syscall.EINVAL
//line /snap/go/10455/src/net/unixsock.go:165
		// _ = "end of CoverTab[8910]"
	} else {
//line /snap/go/10455/src/net/unixsock.go:166
		_go_fuzz_dep_.CoverTab[530148]++
//line /snap/go/10455/src/net/unixsock.go:166
		_go_fuzz_dep_.CoverTab[8911]++
//line /snap/go/10455/src/net/unixsock.go:166
		// _ = "end of CoverTab[8911]"
//line /snap/go/10455/src/net/unixsock.go:166
	}
//line /snap/go/10455/src/net/unixsock.go:166
	// _ = "end of CoverTab[8906]"
//line /snap/go/10455/src/net/unixsock.go:166
	_go_fuzz_dep_.CoverTab[8907]++
						a, ok := addr.(*UnixAddr)
						if !ok {
//line /snap/go/10455/src/net/unixsock.go:168
		_go_fuzz_dep_.CoverTab[530149]++
//line /snap/go/10455/src/net/unixsock.go:168
		_go_fuzz_dep_.CoverTab[8912]++
							return 0, &OpError{Op: "write", Net: c.fd.net, Source: c.fd.laddr, Addr: addr, Err: syscall.EINVAL}
//line /snap/go/10455/src/net/unixsock.go:169
		// _ = "end of CoverTab[8912]"
	} else {
//line /snap/go/10455/src/net/unixsock.go:170
		_go_fuzz_dep_.CoverTab[530150]++
//line /snap/go/10455/src/net/unixsock.go:170
		_go_fuzz_dep_.CoverTab[8913]++
//line /snap/go/10455/src/net/unixsock.go:170
		// _ = "end of CoverTab[8913]"
//line /snap/go/10455/src/net/unixsock.go:170
	}
//line /snap/go/10455/src/net/unixsock.go:170
	// _ = "end of CoverTab[8907]"
//line /snap/go/10455/src/net/unixsock.go:170
	_go_fuzz_dep_.CoverTab[8908]++
						n, err := c.writeTo(b, a)
						if err != nil {
//line /snap/go/10455/src/net/unixsock.go:172
		_go_fuzz_dep_.CoverTab[530151]++
//line /snap/go/10455/src/net/unixsock.go:172
		_go_fuzz_dep_.CoverTab[8914]++
							err = &OpError{Op: "write", Net: c.fd.net, Source: c.fd.laddr, Addr: a.opAddr(), Err: err}
//line /snap/go/10455/src/net/unixsock.go:173
		// _ = "end of CoverTab[8914]"
	} else {
//line /snap/go/10455/src/net/unixsock.go:174
		_go_fuzz_dep_.CoverTab[530152]++
//line /snap/go/10455/src/net/unixsock.go:174
		_go_fuzz_dep_.CoverTab[8915]++
//line /snap/go/10455/src/net/unixsock.go:174
		// _ = "end of CoverTab[8915]"
//line /snap/go/10455/src/net/unixsock.go:174
	}
//line /snap/go/10455/src/net/unixsock.go:174
	// _ = "end of CoverTab[8908]"
//line /snap/go/10455/src/net/unixsock.go:174
	_go_fuzz_dep_.CoverTab[8909]++
						return n, err
//line /snap/go/10455/src/net/unixsock.go:175
	// _ = "end of CoverTab[8909]"
}

// WriteMsgUnix writes a message to addr via c, copying the payload
//line /snap/go/10455/src/net/unixsock.go:178
// from b and the associated out-of-band data from oob. It returns the
//line /snap/go/10455/src/net/unixsock.go:178
// number of payload and out-of-band bytes written.
//line /snap/go/10455/src/net/unixsock.go:178
//
//line /snap/go/10455/src/net/unixsock.go:178
// Note that if len(b) == 0 and len(oob) > 0, this function will still
//line /snap/go/10455/src/net/unixsock.go:178
// write 1 byte to the connection.
//line /snap/go/10455/src/net/unixsock.go:184
func (c *UnixConn) WriteMsgUnix(b, oob []byte, addr *UnixAddr) (n, oobn int, err error) {
//line /snap/go/10455/src/net/unixsock.go:184
	_go_fuzz_dep_.CoverTab[8916]++
						if !c.ok() {
//line /snap/go/10455/src/net/unixsock.go:185
		_go_fuzz_dep_.CoverTab[530153]++
//line /snap/go/10455/src/net/unixsock.go:185
		_go_fuzz_dep_.CoverTab[8919]++
							return 0, 0, syscall.EINVAL
//line /snap/go/10455/src/net/unixsock.go:186
		// _ = "end of CoverTab[8919]"
	} else {
//line /snap/go/10455/src/net/unixsock.go:187
		_go_fuzz_dep_.CoverTab[530154]++
//line /snap/go/10455/src/net/unixsock.go:187
		_go_fuzz_dep_.CoverTab[8920]++
//line /snap/go/10455/src/net/unixsock.go:187
		// _ = "end of CoverTab[8920]"
//line /snap/go/10455/src/net/unixsock.go:187
	}
//line /snap/go/10455/src/net/unixsock.go:187
	// _ = "end of CoverTab[8916]"
//line /snap/go/10455/src/net/unixsock.go:187
	_go_fuzz_dep_.CoverTab[8917]++
						n, oobn, err = c.writeMsg(b, oob, addr)
						if err != nil {
//line /snap/go/10455/src/net/unixsock.go:189
		_go_fuzz_dep_.CoverTab[530155]++
//line /snap/go/10455/src/net/unixsock.go:189
		_go_fuzz_dep_.CoverTab[8921]++
							err = &OpError{Op: "write", Net: c.fd.net, Source: c.fd.laddr, Addr: addr.opAddr(), Err: err}
//line /snap/go/10455/src/net/unixsock.go:190
		// _ = "end of CoverTab[8921]"
	} else {
//line /snap/go/10455/src/net/unixsock.go:191
		_go_fuzz_dep_.CoverTab[530156]++
//line /snap/go/10455/src/net/unixsock.go:191
		_go_fuzz_dep_.CoverTab[8922]++
//line /snap/go/10455/src/net/unixsock.go:191
		// _ = "end of CoverTab[8922]"
//line /snap/go/10455/src/net/unixsock.go:191
	}
//line /snap/go/10455/src/net/unixsock.go:191
	// _ = "end of CoverTab[8917]"
//line /snap/go/10455/src/net/unixsock.go:191
	_go_fuzz_dep_.CoverTab[8918]++
						return
//line /snap/go/10455/src/net/unixsock.go:192
	// _ = "end of CoverTab[8918]"
}

func newUnixConn(fd *netFD) *UnixConn {
//line /snap/go/10455/src/net/unixsock.go:195
	_go_fuzz_dep_.CoverTab[8923]++
//line /snap/go/10455/src/net/unixsock.go:195
	return &UnixConn{conn{fd}}
//line /snap/go/10455/src/net/unixsock.go:195
	// _ = "end of CoverTab[8923]"
//line /snap/go/10455/src/net/unixsock.go:195
}

// DialUnix acts like Dial for Unix networks.
//line /snap/go/10455/src/net/unixsock.go:197
//
//line /snap/go/10455/src/net/unixsock.go:197
// The network must be a Unix network name; see func Dial for details.
//line /snap/go/10455/src/net/unixsock.go:197
//
//line /snap/go/10455/src/net/unixsock.go:197
// If laddr is non-nil, it is used as the local address for the
//line /snap/go/10455/src/net/unixsock.go:197
// connection.
//line /snap/go/10455/src/net/unixsock.go:203
func DialUnix(network string, laddr, raddr *UnixAddr) (*UnixConn, error) {
//line /snap/go/10455/src/net/unixsock.go:203
	_go_fuzz_dep_.CoverTab[8924]++
						switch network {
	case "unix", "unixgram", "unixpacket":
//line /snap/go/10455/src/net/unixsock.go:205
		_go_fuzz_dep_.CoverTab[530157]++
//line /snap/go/10455/src/net/unixsock.go:205
		_go_fuzz_dep_.CoverTab[8927]++
//line /snap/go/10455/src/net/unixsock.go:205
		// _ = "end of CoverTab[8927]"
	default:
//line /snap/go/10455/src/net/unixsock.go:206
		_go_fuzz_dep_.CoverTab[530158]++
//line /snap/go/10455/src/net/unixsock.go:206
		_go_fuzz_dep_.CoverTab[8928]++
							return nil, &OpError{Op: "dial", Net: network, Source: laddr.opAddr(), Addr: raddr.opAddr(), Err: UnknownNetworkError(network)}
//line /snap/go/10455/src/net/unixsock.go:207
		// _ = "end of CoverTab[8928]"
	}
//line /snap/go/10455/src/net/unixsock.go:208
	// _ = "end of CoverTab[8924]"
//line /snap/go/10455/src/net/unixsock.go:208
	_go_fuzz_dep_.CoverTab[8925]++
						sd := &sysDialer{network: network, address: raddr.String()}
						c, err := sd.dialUnix(context.Background(), laddr, raddr)
						if err != nil {
//line /snap/go/10455/src/net/unixsock.go:211
		_go_fuzz_dep_.CoverTab[530159]++
//line /snap/go/10455/src/net/unixsock.go:211
		_go_fuzz_dep_.CoverTab[8929]++
							return nil, &OpError{Op: "dial", Net: network, Source: laddr.opAddr(), Addr: raddr.opAddr(), Err: err}
//line /snap/go/10455/src/net/unixsock.go:212
		// _ = "end of CoverTab[8929]"
	} else {
//line /snap/go/10455/src/net/unixsock.go:213
		_go_fuzz_dep_.CoverTab[530160]++
//line /snap/go/10455/src/net/unixsock.go:213
		_go_fuzz_dep_.CoverTab[8930]++
//line /snap/go/10455/src/net/unixsock.go:213
		// _ = "end of CoverTab[8930]"
//line /snap/go/10455/src/net/unixsock.go:213
	}
//line /snap/go/10455/src/net/unixsock.go:213
	// _ = "end of CoverTab[8925]"
//line /snap/go/10455/src/net/unixsock.go:213
	_go_fuzz_dep_.CoverTab[8926]++
						return c, nil
//line /snap/go/10455/src/net/unixsock.go:214
	// _ = "end of CoverTab[8926]"
}

// UnixListener is a Unix domain socket listener. Clients should
//line /snap/go/10455/src/net/unixsock.go:217
// typically use variables of type Listener instead of assuming Unix
//line /snap/go/10455/src/net/unixsock.go:217
// domain sockets.
//line /snap/go/10455/src/net/unixsock.go:220
type UnixListener struct {
	fd		*netFD
	path		string
	unlink		bool
	unlinkOnce	sync.Once
}

func (ln *UnixListener) ok() bool {
//line /snap/go/10455/src/net/unixsock.go:227
	_go_fuzz_dep_.CoverTab[8931]++
//line /snap/go/10455/src/net/unixsock.go:227
	return ln != nil && func() bool {
//line /snap/go/10455/src/net/unixsock.go:227
		_go_fuzz_dep_.CoverTab[8932]++
//line /snap/go/10455/src/net/unixsock.go:227
		return ln.fd != nil
//line /snap/go/10455/src/net/unixsock.go:227
		// _ = "end of CoverTab[8932]"
//line /snap/go/10455/src/net/unixsock.go:227
	}()
//line /snap/go/10455/src/net/unixsock.go:227
	// _ = "end of CoverTab[8931]"
//line /snap/go/10455/src/net/unixsock.go:227
}

// SyscallConn returns a raw network connection.
//line /snap/go/10455/src/net/unixsock.go:229
// This implements the syscall.Conn interface.
//line /snap/go/10455/src/net/unixsock.go:229
//
//line /snap/go/10455/src/net/unixsock.go:229
// The returned RawConn only supports calling Control. Read and
//line /snap/go/10455/src/net/unixsock.go:229
// Write return an error.
//line /snap/go/10455/src/net/unixsock.go:234
func (l *UnixListener) SyscallConn() (syscall.RawConn, error) {
//line /snap/go/10455/src/net/unixsock.go:234
	_go_fuzz_dep_.CoverTab[8933]++
						if !l.ok() {
//line /snap/go/10455/src/net/unixsock.go:235
		_go_fuzz_dep_.CoverTab[530161]++
//line /snap/go/10455/src/net/unixsock.go:235
		_go_fuzz_dep_.CoverTab[8935]++
							return nil, syscall.EINVAL
//line /snap/go/10455/src/net/unixsock.go:236
		// _ = "end of CoverTab[8935]"
	} else {
//line /snap/go/10455/src/net/unixsock.go:237
		_go_fuzz_dep_.CoverTab[530162]++
//line /snap/go/10455/src/net/unixsock.go:237
		_go_fuzz_dep_.CoverTab[8936]++
//line /snap/go/10455/src/net/unixsock.go:237
		// _ = "end of CoverTab[8936]"
//line /snap/go/10455/src/net/unixsock.go:237
	}
//line /snap/go/10455/src/net/unixsock.go:237
	// _ = "end of CoverTab[8933]"
//line /snap/go/10455/src/net/unixsock.go:237
	_go_fuzz_dep_.CoverTab[8934]++
						return newRawListener(l.fd)
//line /snap/go/10455/src/net/unixsock.go:238
	// _ = "end of CoverTab[8934]"
}

// AcceptUnix accepts the next incoming call and returns the new
//line /snap/go/10455/src/net/unixsock.go:241
// connection.
//line /snap/go/10455/src/net/unixsock.go:243
func (l *UnixListener) AcceptUnix() (*UnixConn, error) {
//line /snap/go/10455/src/net/unixsock.go:243
	_go_fuzz_dep_.CoverTab[8937]++
						if !l.ok() {
//line /snap/go/10455/src/net/unixsock.go:244
		_go_fuzz_dep_.CoverTab[530163]++
//line /snap/go/10455/src/net/unixsock.go:244
		_go_fuzz_dep_.CoverTab[8940]++
							return nil, syscall.EINVAL
//line /snap/go/10455/src/net/unixsock.go:245
		// _ = "end of CoverTab[8940]"
	} else {
//line /snap/go/10455/src/net/unixsock.go:246
		_go_fuzz_dep_.CoverTab[530164]++
//line /snap/go/10455/src/net/unixsock.go:246
		_go_fuzz_dep_.CoverTab[8941]++
//line /snap/go/10455/src/net/unixsock.go:246
		// _ = "end of CoverTab[8941]"
//line /snap/go/10455/src/net/unixsock.go:246
	}
//line /snap/go/10455/src/net/unixsock.go:246
	// _ = "end of CoverTab[8937]"
//line /snap/go/10455/src/net/unixsock.go:246
	_go_fuzz_dep_.CoverTab[8938]++
						c, err := l.accept()
						if err != nil {
//line /snap/go/10455/src/net/unixsock.go:248
		_go_fuzz_dep_.CoverTab[530165]++
//line /snap/go/10455/src/net/unixsock.go:248
		_go_fuzz_dep_.CoverTab[8942]++
							return nil, &OpError{Op: "accept", Net: l.fd.net, Source: nil, Addr: l.fd.laddr, Err: err}
//line /snap/go/10455/src/net/unixsock.go:249
		// _ = "end of CoverTab[8942]"
	} else {
//line /snap/go/10455/src/net/unixsock.go:250
		_go_fuzz_dep_.CoverTab[530166]++
//line /snap/go/10455/src/net/unixsock.go:250
		_go_fuzz_dep_.CoverTab[8943]++
//line /snap/go/10455/src/net/unixsock.go:250
		// _ = "end of CoverTab[8943]"
//line /snap/go/10455/src/net/unixsock.go:250
	}
//line /snap/go/10455/src/net/unixsock.go:250
	// _ = "end of CoverTab[8938]"
//line /snap/go/10455/src/net/unixsock.go:250
	_go_fuzz_dep_.CoverTab[8939]++
						return c, nil
//line /snap/go/10455/src/net/unixsock.go:251
	// _ = "end of CoverTab[8939]"
}

// Accept implements the Accept method in the Listener interface.
//line /snap/go/10455/src/net/unixsock.go:254
// Returned connections will be of type *UnixConn.
//line /snap/go/10455/src/net/unixsock.go:256
func (l *UnixListener) Accept() (Conn, error) {
//line /snap/go/10455/src/net/unixsock.go:256
	_go_fuzz_dep_.CoverTab[8944]++
						if !l.ok() {
//line /snap/go/10455/src/net/unixsock.go:257
		_go_fuzz_dep_.CoverTab[530167]++
//line /snap/go/10455/src/net/unixsock.go:257
		_go_fuzz_dep_.CoverTab[8947]++
							return nil, syscall.EINVAL
//line /snap/go/10455/src/net/unixsock.go:258
		// _ = "end of CoverTab[8947]"
	} else {
//line /snap/go/10455/src/net/unixsock.go:259
		_go_fuzz_dep_.CoverTab[530168]++
//line /snap/go/10455/src/net/unixsock.go:259
		_go_fuzz_dep_.CoverTab[8948]++
//line /snap/go/10455/src/net/unixsock.go:259
		// _ = "end of CoverTab[8948]"
//line /snap/go/10455/src/net/unixsock.go:259
	}
//line /snap/go/10455/src/net/unixsock.go:259
	// _ = "end of CoverTab[8944]"
//line /snap/go/10455/src/net/unixsock.go:259
	_go_fuzz_dep_.CoverTab[8945]++
						c, err := l.accept()
						if err != nil {
//line /snap/go/10455/src/net/unixsock.go:261
		_go_fuzz_dep_.CoverTab[530169]++
//line /snap/go/10455/src/net/unixsock.go:261
		_go_fuzz_dep_.CoverTab[8949]++
							return nil, &OpError{Op: "accept", Net: l.fd.net, Source: nil, Addr: l.fd.laddr, Err: err}
//line /snap/go/10455/src/net/unixsock.go:262
		// _ = "end of CoverTab[8949]"
	} else {
//line /snap/go/10455/src/net/unixsock.go:263
		_go_fuzz_dep_.CoverTab[530170]++
//line /snap/go/10455/src/net/unixsock.go:263
		_go_fuzz_dep_.CoverTab[8950]++
//line /snap/go/10455/src/net/unixsock.go:263
		// _ = "end of CoverTab[8950]"
//line /snap/go/10455/src/net/unixsock.go:263
	}
//line /snap/go/10455/src/net/unixsock.go:263
	// _ = "end of CoverTab[8945]"
//line /snap/go/10455/src/net/unixsock.go:263
	_go_fuzz_dep_.CoverTab[8946]++
						return c, nil
//line /snap/go/10455/src/net/unixsock.go:264
	// _ = "end of CoverTab[8946]"
}

// Close stops listening on the Unix address. Already accepted
//line /snap/go/10455/src/net/unixsock.go:267
// connections are not closed.
//line /snap/go/10455/src/net/unixsock.go:269
func (l *UnixListener) Close() error {
//line /snap/go/10455/src/net/unixsock.go:269
	_go_fuzz_dep_.CoverTab[8951]++
						if !l.ok() {
//line /snap/go/10455/src/net/unixsock.go:270
		_go_fuzz_dep_.CoverTab[530171]++
//line /snap/go/10455/src/net/unixsock.go:270
		_go_fuzz_dep_.CoverTab[8954]++
							return syscall.EINVAL
//line /snap/go/10455/src/net/unixsock.go:271
		// _ = "end of CoverTab[8954]"
	} else {
//line /snap/go/10455/src/net/unixsock.go:272
		_go_fuzz_dep_.CoverTab[530172]++
//line /snap/go/10455/src/net/unixsock.go:272
		_go_fuzz_dep_.CoverTab[8955]++
//line /snap/go/10455/src/net/unixsock.go:272
		// _ = "end of CoverTab[8955]"
//line /snap/go/10455/src/net/unixsock.go:272
	}
//line /snap/go/10455/src/net/unixsock.go:272
	// _ = "end of CoverTab[8951]"
//line /snap/go/10455/src/net/unixsock.go:272
	_go_fuzz_dep_.CoverTab[8952]++
						if err := l.close(); err != nil {
//line /snap/go/10455/src/net/unixsock.go:273
		_go_fuzz_dep_.CoverTab[530173]++
//line /snap/go/10455/src/net/unixsock.go:273
		_go_fuzz_dep_.CoverTab[8956]++
							return &OpError{Op: "close", Net: l.fd.net, Source: nil, Addr: l.fd.laddr, Err: err}
//line /snap/go/10455/src/net/unixsock.go:274
		// _ = "end of CoverTab[8956]"
	} else {
//line /snap/go/10455/src/net/unixsock.go:275
		_go_fuzz_dep_.CoverTab[530174]++
//line /snap/go/10455/src/net/unixsock.go:275
		_go_fuzz_dep_.CoverTab[8957]++
//line /snap/go/10455/src/net/unixsock.go:275
		// _ = "end of CoverTab[8957]"
//line /snap/go/10455/src/net/unixsock.go:275
	}
//line /snap/go/10455/src/net/unixsock.go:275
	// _ = "end of CoverTab[8952]"
//line /snap/go/10455/src/net/unixsock.go:275
	_go_fuzz_dep_.CoverTab[8953]++
						return nil
//line /snap/go/10455/src/net/unixsock.go:276
	// _ = "end of CoverTab[8953]"
}

// Addr returns the listener's network address.
//line /snap/go/10455/src/net/unixsock.go:279
// The Addr returned is shared by all invocations of Addr, so
//line /snap/go/10455/src/net/unixsock.go:279
// do not modify it.
//line /snap/go/10455/src/net/unixsock.go:282
func (l *UnixListener) Addr() Addr {
//line /snap/go/10455/src/net/unixsock.go:282
	_go_fuzz_dep_.CoverTab[8958]++
//line /snap/go/10455/src/net/unixsock.go:282
	return l.fd.laddr
//line /snap/go/10455/src/net/unixsock.go:282
	// _ = "end of CoverTab[8958]"
//line /snap/go/10455/src/net/unixsock.go:282
}

// SetDeadline sets the deadline associated with the listener.
//line /snap/go/10455/src/net/unixsock.go:284
// A zero time value disables the deadline.
//line /snap/go/10455/src/net/unixsock.go:286
func (l *UnixListener) SetDeadline(t time.Time) error {
//line /snap/go/10455/src/net/unixsock.go:286
	_go_fuzz_dep_.CoverTab[8959]++
						if !l.ok() {
//line /snap/go/10455/src/net/unixsock.go:287
		_go_fuzz_dep_.CoverTab[530175]++
//line /snap/go/10455/src/net/unixsock.go:287
		_go_fuzz_dep_.CoverTab[8962]++
							return syscall.EINVAL
//line /snap/go/10455/src/net/unixsock.go:288
		// _ = "end of CoverTab[8962]"
	} else {
//line /snap/go/10455/src/net/unixsock.go:289
		_go_fuzz_dep_.CoverTab[530176]++
//line /snap/go/10455/src/net/unixsock.go:289
		_go_fuzz_dep_.CoverTab[8963]++
//line /snap/go/10455/src/net/unixsock.go:289
		// _ = "end of CoverTab[8963]"
//line /snap/go/10455/src/net/unixsock.go:289
	}
//line /snap/go/10455/src/net/unixsock.go:289
	// _ = "end of CoverTab[8959]"
//line /snap/go/10455/src/net/unixsock.go:289
	_go_fuzz_dep_.CoverTab[8960]++
						if err := l.fd.pfd.SetDeadline(t); err != nil {
//line /snap/go/10455/src/net/unixsock.go:290
		_go_fuzz_dep_.CoverTab[530177]++
//line /snap/go/10455/src/net/unixsock.go:290
		_go_fuzz_dep_.CoverTab[8964]++
							return &OpError{Op: "set", Net: l.fd.net, Source: nil, Addr: l.fd.laddr, Err: err}
//line /snap/go/10455/src/net/unixsock.go:291
		// _ = "end of CoverTab[8964]"
	} else {
//line /snap/go/10455/src/net/unixsock.go:292
		_go_fuzz_dep_.CoverTab[530178]++
//line /snap/go/10455/src/net/unixsock.go:292
		_go_fuzz_dep_.CoverTab[8965]++
//line /snap/go/10455/src/net/unixsock.go:292
		// _ = "end of CoverTab[8965]"
//line /snap/go/10455/src/net/unixsock.go:292
	}
//line /snap/go/10455/src/net/unixsock.go:292
	// _ = "end of CoverTab[8960]"
//line /snap/go/10455/src/net/unixsock.go:292
	_go_fuzz_dep_.CoverTab[8961]++
						return nil
//line /snap/go/10455/src/net/unixsock.go:293
	// _ = "end of CoverTab[8961]"
}

// File returns a copy of the underlying os.File.
//line /snap/go/10455/src/net/unixsock.go:296
// It is the caller's responsibility to close f when finished.
//line /snap/go/10455/src/net/unixsock.go:296
// Closing l does not affect f, and closing f does not affect l.
//line /snap/go/10455/src/net/unixsock.go:296
//
//line /snap/go/10455/src/net/unixsock.go:296
// The returned os.File's file descriptor is different from the
//line /snap/go/10455/src/net/unixsock.go:296
// connection's. Attempting to change properties of the original
//line /snap/go/10455/src/net/unixsock.go:296
// using this duplicate may or may not have the desired effect.
//line /snap/go/10455/src/net/unixsock.go:303
func (l *UnixListener) File() (f *os.File, err error) {
//line /snap/go/10455/src/net/unixsock.go:303
	_go_fuzz_dep_.CoverTab[8966]++
						if !l.ok() {
//line /snap/go/10455/src/net/unixsock.go:304
		_go_fuzz_dep_.CoverTab[530179]++
//line /snap/go/10455/src/net/unixsock.go:304
		_go_fuzz_dep_.CoverTab[8969]++
							return nil, syscall.EINVAL
//line /snap/go/10455/src/net/unixsock.go:305
		// _ = "end of CoverTab[8969]"
	} else {
//line /snap/go/10455/src/net/unixsock.go:306
		_go_fuzz_dep_.CoverTab[530180]++
//line /snap/go/10455/src/net/unixsock.go:306
		_go_fuzz_dep_.CoverTab[8970]++
//line /snap/go/10455/src/net/unixsock.go:306
		// _ = "end of CoverTab[8970]"
//line /snap/go/10455/src/net/unixsock.go:306
	}
//line /snap/go/10455/src/net/unixsock.go:306
	// _ = "end of CoverTab[8966]"
//line /snap/go/10455/src/net/unixsock.go:306
	_go_fuzz_dep_.CoverTab[8967]++
						f, err = l.file()
						if err != nil {
//line /snap/go/10455/src/net/unixsock.go:308
		_go_fuzz_dep_.CoverTab[530181]++
//line /snap/go/10455/src/net/unixsock.go:308
		_go_fuzz_dep_.CoverTab[8971]++
							err = &OpError{Op: "file", Net: l.fd.net, Source: nil, Addr: l.fd.laddr, Err: err}
//line /snap/go/10455/src/net/unixsock.go:309
		// _ = "end of CoverTab[8971]"
	} else {
//line /snap/go/10455/src/net/unixsock.go:310
		_go_fuzz_dep_.CoverTab[530182]++
//line /snap/go/10455/src/net/unixsock.go:310
		_go_fuzz_dep_.CoverTab[8972]++
//line /snap/go/10455/src/net/unixsock.go:310
		// _ = "end of CoverTab[8972]"
//line /snap/go/10455/src/net/unixsock.go:310
	}
//line /snap/go/10455/src/net/unixsock.go:310
	// _ = "end of CoverTab[8967]"
//line /snap/go/10455/src/net/unixsock.go:310
	_go_fuzz_dep_.CoverTab[8968]++
						return
//line /snap/go/10455/src/net/unixsock.go:311
	// _ = "end of CoverTab[8968]"
}

// ListenUnix acts like Listen for Unix networks.
//line /snap/go/10455/src/net/unixsock.go:314
//
//line /snap/go/10455/src/net/unixsock.go:314
// The network must be "unix" or "unixpacket".
//line /snap/go/10455/src/net/unixsock.go:317
func ListenUnix(network string, laddr *UnixAddr) (*UnixListener, error) {
//line /snap/go/10455/src/net/unixsock.go:317
	_go_fuzz_dep_.CoverTab[8973]++
						switch network {
	case "unix", "unixpacket":
//line /snap/go/10455/src/net/unixsock.go:319
		_go_fuzz_dep_.CoverTab[530183]++
//line /snap/go/10455/src/net/unixsock.go:319
		_go_fuzz_dep_.CoverTab[8977]++
//line /snap/go/10455/src/net/unixsock.go:319
		// _ = "end of CoverTab[8977]"
	default:
//line /snap/go/10455/src/net/unixsock.go:320
		_go_fuzz_dep_.CoverTab[530184]++
//line /snap/go/10455/src/net/unixsock.go:320
		_go_fuzz_dep_.CoverTab[8978]++
							return nil, &OpError{Op: "listen", Net: network, Source: nil, Addr: laddr.opAddr(), Err: UnknownNetworkError(network)}
//line /snap/go/10455/src/net/unixsock.go:321
		// _ = "end of CoverTab[8978]"
	}
//line /snap/go/10455/src/net/unixsock.go:322
	// _ = "end of CoverTab[8973]"
//line /snap/go/10455/src/net/unixsock.go:322
	_go_fuzz_dep_.CoverTab[8974]++
						if laddr == nil {
//line /snap/go/10455/src/net/unixsock.go:323
		_go_fuzz_dep_.CoverTab[530185]++
//line /snap/go/10455/src/net/unixsock.go:323
		_go_fuzz_dep_.CoverTab[8979]++
							return nil, &OpError{Op: "listen", Net: network, Source: nil, Addr: laddr.opAddr(), Err: errMissingAddress}
//line /snap/go/10455/src/net/unixsock.go:324
		// _ = "end of CoverTab[8979]"
	} else {
//line /snap/go/10455/src/net/unixsock.go:325
		_go_fuzz_dep_.CoverTab[530186]++
//line /snap/go/10455/src/net/unixsock.go:325
		_go_fuzz_dep_.CoverTab[8980]++
//line /snap/go/10455/src/net/unixsock.go:325
		// _ = "end of CoverTab[8980]"
//line /snap/go/10455/src/net/unixsock.go:325
	}
//line /snap/go/10455/src/net/unixsock.go:325
	// _ = "end of CoverTab[8974]"
//line /snap/go/10455/src/net/unixsock.go:325
	_go_fuzz_dep_.CoverTab[8975]++
						sl := &sysListener{network: network, address: laddr.String()}
						ln, err := sl.listenUnix(context.Background(), laddr)
						if err != nil {
//line /snap/go/10455/src/net/unixsock.go:328
		_go_fuzz_dep_.CoverTab[530187]++
//line /snap/go/10455/src/net/unixsock.go:328
		_go_fuzz_dep_.CoverTab[8981]++
							return nil, &OpError{Op: "listen", Net: network, Source: nil, Addr: laddr.opAddr(), Err: err}
//line /snap/go/10455/src/net/unixsock.go:329
		// _ = "end of CoverTab[8981]"
	} else {
//line /snap/go/10455/src/net/unixsock.go:330
		_go_fuzz_dep_.CoverTab[530188]++
//line /snap/go/10455/src/net/unixsock.go:330
		_go_fuzz_dep_.CoverTab[8982]++
//line /snap/go/10455/src/net/unixsock.go:330
		// _ = "end of CoverTab[8982]"
//line /snap/go/10455/src/net/unixsock.go:330
	}
//line /snap/go/10455/src/net/unixsock.go:330
	// _ = "end of CoverTab[8975]"
//line /snap/go/10455/src/net/unixsock.go:330
	_go_fuzz_dep_.CoverTab[8976]++
						return ln, nil
//line /snap/go/10455/src/net/unixsock.go:331
	// _ = "end of CoverTab[8976]"
}

// ListenUnixgram acts like ListenPacket for Unix networks.
//line /snap/go/10455/src/net/unixsock.go:334
//
//line /snap/go/10455/src/net/unixsock.go:334
// The network must be "unixgram".
//line /snap/go/10455/src/net/unixsock.go:337
func ListenUnixgram(network string, laddr *UnixAddr) (*UnixConn, error) {
//line /snap/go/10455/src/net/unixsock.go:337
	_go_fuzz_dep_.CoverTab[8983]++
						switch network {
	case "unixgram":
//line /snap/go/10455/src/net/unixsock.go:339
		_go_fuzz_dep_.CoverTab[530189]++
//line /snap/go/10455/src/net/unixsock.go:339
		_go_fuzz_dep_.CoverTab[8987]++
//line /snap/go/10455/src/net/unixsock.go:339
		// _ = "end of CoverTab[8987]"
	default:
//line /snap/go/10455/src/net/unixsock.go:340
		_go_fuzz_dep_.CoverTab[530190]++
//line /snap/go/10455/src/net/unixsock.go:340
		_go_fuzz_dep_.CoverTab[8988]++
							return nil, &OpError{Op: "listen", Net: network, Source: nil, Addr: laddr.opAddr(), Err: UnknownNetworkError(network)}
//line /snap/go/10455/src/net/unixsock.go:341
		// _ = "end of CoverTab[8988]"
	}
//line /snap/go/10455/src/net/unixsock.go:342
	// _ = "end of CoverTab[8983]"
//line /snap/go/10455/src/net/unixsock.go:342
	_go_fuzz_dep_.CoverTab[8984]++
						if laddr == nil {
//line /snap/go/10455/src/net/unixsock.go:343
		_go_fuzz_dep_.CoverTab[530191]++
//line /snap/go/10455/src/net/unixsock.go:343
		_go_fuzz_dep_.CoverTab[8989]++
							return nil, &OpError{Op: "listen", Net: network, Source: nil, Addr: nil, Err: errMissingAddress}
//line /snap/go/10455/src/net/unixsock.go:344
		// _ = "end of CoverTab[8989]"
	} else {
//line /snap/go/10455/src/net/unixsock.go:345
		_go_fuzz_dep_.CoverTab[530192]++
//line /snap/go/10455/src/net/unixsock.go:345
		_go_fuzz_dep_.CoverTab[8990]++
//line /snap/go/10455/src/net/unixsock.go:345
		// _ = "end of CoverTab[8990]"
//line /snap/go/10455/src/net/unixsock.go:345
	}
//line /snap/go/10455/src/net/unixsock.go:345
	// _ = "end of CoverTab[8984]"
//line /snap/go/10455/src/net/unixsock.go:345
	_go_fuzz_dep_.CoverTab[8985]++
						sl := &sysListener{network: network, address: laddr.String()}
						c, err := sl.listenUnixgram(context.Background(), laddr)
						if err != nil {
//line /snap/go/10455/src/net/unixsock.go:348
		_go_fuzz_dep_.CoverTab[530193]++
//line /snap/go/10455/src/net/unixsock.go:348
		_go_fuzz_dep_.CoverTab[8991]++
							return nil, &OpError{Op: "listen", Net: network, Source: nil, Addr: laddr.opAddr(), Err: err}
//line /snap/go/10455/src/net/unixsock.go:349
		// _ = "end of CoverTab[8991]"
	} else {
//line /snap/go/10455/src/net/unixsock.go:350
		_go_fuzz_dep_.CoverTab[530194]++
//line /snap/go/10455/src/net/unixsock.go:350
		_go_fuzz_dep_.CoverTab[8992]++
//line /snap/go/10455/src/net/unixsock.go:350
		// _ = "end of CoverTab[8992]"
//line /snap/go/10455/src/net/unixsock.go:350
	}
//line /snap/go/10455/src/net/unixsock.go:350
	// _ = "end of CoverTab[8985]"
//line /snap/go/10455/src/net/unixsock.go:350
	_go_fuzz_dep_.CoverTab[8986]++
						return c, nil
//line /snap/go/10455/src/net/unixsock.go:351
	// _ = "end of CoverTab[8986]"
}

//line /snap/go/10455/src/net/unixsock.go:352
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /snap/go/10455/src/net/unixsock.go:352
var _ = _go_fuzz_dep_.CoverTab
