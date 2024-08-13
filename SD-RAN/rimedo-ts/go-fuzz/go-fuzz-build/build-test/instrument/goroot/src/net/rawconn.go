// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/net/rawconn.go:5
package net

//line /usr/local/go/src/net/rawconn.go:5
import (
//line /usr/local/go/src/net/rawconn.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/rawconn.go:5
)
//line /usr/local/go/src/net/rawconn.go:5
import (
//line /usr/local/go/src/net/rawconn.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/rawconn.go:5
)

import (
	"runtime"
	"syscall"
)

//line /usr/local/go/src/net/rawconn.go:21
type rawConn struct {
	fd *netFD
}

func (c *rawConn) ok() bool {
//line /usr/local/go/src/net/rawconn.go:25
	_go_fuzz_dep_.CoverTab[7674]++
//line /usr/local/go/src/net/rawconn.go:25
	return c != nil && func() bool {
//line /usr/local/go/src/net/rawconn.go:25
		_go_fuzz_dep_.CoverTab[7675]++
//line /usr/local/go/src/net/rawconn.go:25
		return c.fd != nil
//line /usr/local/go/src/net/rawconn.go:25
		// _ = "end of CoverTab[7675]"
//line /usr/local/go/src/net/rawconn.go:25
	}()
//line /usr/local/go/src/net/rawconn.go:25
	// _ = "end of CoverTab[7674]"
//line /usr/local/go/src/net/rawconn.go:25
}

func (c *rawConn) Control(f func(uintptr)) error {
//line /usr/local/go/src/net/rawconn.go:27
	_go_fuzz_dep_.CoverTab[7676]++
						if !c.ok() {
//line /usr/local/go/src/net/rawconn.go:28
		_go_fuzz_dep_.CoverTab[7679]++
							return syscall.EINVAL
//line /usr/local/go/src/net/rawconn.go:29
		// _ = "end of CoverTab[7679]"
	} else {
//line /usr/local/go/src/net/rawconn.go:30
		_go_fuzz_dep_.CoverTab[7680]++
//line /usr/local/go/src/net/rawconn.go:30
		// _ = "end of CoverTab[7680]"
//line /usr/local/go/src/net/rawconn.go:30
	}
//line /usr/local/go/src/net/rawconn.go:30
	// _ = "end of CoverTab[7676]"
//line /usr/local/go/src/net/rawconn.go:30
	_go_fuzz_dep_.CoverTab[7677]++
						err := c.fd.pfd.RawControl(f)
						runtime.KeepAlive(c.fd)
						if err != nil {
//line /usr/local/go/src/net/rawconn.go:33
		_go_fuzz_dep_.CoverTab[7681]++
							err = &OpError{Op: "raw-control", Net: c.fd.net, Source: nil, Addr: c.fd.laddr, Err: err}
//line /usr/local/go/src/net/rawconn.go:34
		// _ = "end of CoverTab[7681]"
	} else {
//line /usr/local/go/src/net/rawconn.go:35
		_go_fuzz_dep_.CoverTab[7682]++
//line /usr/local/go/src/net/rawconn.go:35
		// _ = "end of CoverTab[7682]"
//line /usr/local/go/src/net/rawconn.go:35
	}
//line /usr/local/go/src/net/rawconn.go:35
	// _ = "end of CoverTab[7677]"
//line /usr/local/go/src/net/rawconn.go:35
	_go_fuzz_dep_.CoverTab[7678]++
						return err
//line /usr/local/go/src/net/rawconn.go:36
	// _ = "end of CoverTab[7678]"
}

func (c *rawConn) Read(f func(uintptr) bool) error {
//line /usr/local/go/src/net/rawconn.go:39
	_go_fuzz_dep_.CoverTab[7683]++
						if !c.ok() {
//line /usr/local/go/src/net/rawconn.go:40
		_go_fuzz_dep_.CoverTab[7686]++
							return syscall.EINVAL
//line /usr/local/go/src/net/rawconn.go:41
		// _ = "end of CoverTab[7686]"
	} else {
//line /usr/local/go/src/net/rawconn.go:42
		_go_fuzz_dep_.CoverTab[7687]++
//line /usr/local/go/src/net/rawconn.go:42
		// _ = "end of CoverTab[7687]"
//line /usr/local/go/src/net/rawconn.go:42
	}
//line /usr/local/go/src/net/rawconn.go:42
	// _ = "end of CoverTab[7683]"
//line /usr/local/go/src/net/rawconn.go:42
	_go_fuzz_dep_.CoverTab[7684]++
						err := c.fd.pfd.RawRead(f)
						runtime.KeepAlive(c.fd)
						if err != nil {
//line /usr/local/go/src/net/rawconn.go:45
		_go_fuzz_dep_.CoverTab[7688]++
							err = &OpError{Op: "raw-read", Net: c.fd.net, Source: c.fd.laddr, Addr: c.fd.raddr, Err: err}
//line /usr/local/go/src/net/rawconn.go:46
		// _ = "end of CoverTab[7688]"
	} else {
//line /usr/local/go/src/net/rawconn.go:47
		_go_fuzz_dep_.CoverTab[7689]++
//line /usr/local/go/src/net/rawconn.go:47
		// _ = "end of CoverTab[7689]"
//line /usr/local/go/src/net/rawconn.go:47
	}
//line /usr/local/go/src/net/rawconn.go:47
	// _ = "end of CoverTab[7684]"
//line /usr/local/go/src/net/rawconn.go:47
	_go_fuzz_dep_.CoverTab[7685]++
						return err
//line /usr/local/go/src/net/rawconn.go:48
	// _ = "end of CoverTab[7685]"
}

func (c *rawConn) Write(f func(uintptr) bool) error {
//line /usr/local/go/src/net/rawconn.go:51
	_go_fuzz_dep_.CoverTab[7690]++
						if !c.ok() {
//line /usr/local/go/src/net/rawconn.go:52
		_go_fuzz_dep_.CoverTab[7693]++
							return syscall.EINVAL
//line /usr/local/go/src/net/rawconn.go:53
		// _ = "end of CoverTab[7693]"
	} else {
//line /usr/local/go/src/net/rawconn.go:54
		_go_fuzz_dep_.CoverTab[7694]++
//line /usr/local/go/src/net/rawconn.go:54
		// _ = "end of CoverTab[7694]"
//line /usr/local/go/src/net/rawconn.go:54
	}
//line /usr/local/go/src/net/rawconn.go:54
	// _ = "end of CoverTab[7690]"
//line /usr/local/go/src/net/rawconn.go:54
	_go_fuzz_dep_.CoverTab[7691]++
						err := c.fd.pfd.RawWrite(f)
						runtime.KeepAlive(c.fd)
						if err != nil {
//line /usr/local/go/src/net/rawconn.go:57
		_go_fuzz_dep_.CoverTab[7695]++
							err = &OpError{Op: "raw-write", Net: c.fd.net, Source: c.fd.laddr, Addr: c.fd.raddr, Err: err}
//line /usr/local/go/src/net/rawconn.go:58
		// _ = "end of CoverTab[7695]"
	} else {
//line /usr/local/go/src/net/rawconn.go:59
		_go_fuzz_dep_.CoverTab[7696]++
//line /usr/local/go/src/net/rawconn.go:59
		// _ = "end of CoverTab[7696]"
//line /usr/local/go/src/net/rawconn.go:59
	}
//line /usr/local/go/src/net/rawconn.go:59
	// _ = "end of CoverTab[7691]"
//line /usr/local/go/src/net/rawconn.go:59
	_go_fuzz_dep_.CoverTab[7692]++
						return err
//line /usr/local/go/src/net/rawconn.go:60
	// _ = "end of CoverTab[7692]"
}

func newRawConn(fd *netFD) (*rawConn, error) {
//line /usr/local/go/src/net/rawconn.go:63
	_go_fuzz_dep_.CoverTab[7697]++
						return &rawConn{fd: fd}, nil
//line /usr/local/go/src/net/rawconn.go:64
	// _ = "end of CoverTab[7697]"
}

type rawListener struct {
	rawConn
}

func (l *rawListener) Read(func(uintptr) bool) error {
//line /usr/local/go/src/net/rawconn.go:71
	_go_fuzz_dep_.CoverTab[7698]++
						return syscall.EINVAL
//line /usr/local/go/src/net/rawconn.go:72
	// _ = "end of CoverTab[7698]"
}

func (l *rawListener) Write(func(uintptr) bool) error {
//line /usr/local/go/src/net/rawconn.go:75
	_go_fuzz_dep_.CoverTab[7699]++
						return syscall.EINVAL
//line /usr/local/go/src/net/rawconn.go:76
	// _ = "end of CoverTab[7699]"
}

func newRawListener(fd *netFD) (*rawListener, error) {
//line /usr/local/go/src/net/rawconn.go:79
	_go_fuzz_dep_.CoverTab[7700]++
						return &rawListener{rawConn{fd: fd}}, nil
//line /usr/local/go/src/net/rawconn.go:80
	// _ = "end of CoverTab[7700]"
}

//line /usr/local/go/src/net/rawconn.go:81
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/rawconn.go:81
var _ = _go_fuzz_dep_.CoverTab
