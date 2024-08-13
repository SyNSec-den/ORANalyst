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
	_go_fuzz_dep_.CoverTab[16064]++
//line /usr/local/go/src/net/rawconn.go:25
	return c != nil && func() bool {
//line /usr/local/go/src/net/rawconn.go:25
		_go_fuzz_dep_.CoverTab[16065]++
//line /usr/local/go/src/net/rawconn.go:25
		return c.fd != nil
//line /usr/local/go/src/net/rawconn.go:25
		// _ = "end of CoverTab[16065]"
//line /usr/local/go/src/net/rawconn.go:25
	}()
//line /usr/local/go/src/net/rawconn.go:25
	// _ = "end of CoverTab[16064]"
//line /usr/local/go/src/net/rawconn.go:25
}

func (c *rawConn) Control(f func(uintptr)) error {
//line /usr/local/go/src/net/rawconn.go:27
	_go_fuzz_dep_.CoverTab[16066]++
						if !c.ok() {
//line /usr/local/go/src/net/rawconn.go:28
		_go_fuzz_dep_.CoverTab[16069]++
							return syscall.EINVAL
//line /usr/local/go/src/net/rawconn.go:29
		// _ = "end of CoverTab[16069]"
	} else {
//line /usr/local/go/src/net/rawconn.go:30
		_go_fuzz_dep_.CoverTab[16070]++
//line /usr/local/go/src/net/rawconn.go:30
		// _ = "end of CoverTab[16070]"
//line /usr/local/go/src/net/rawconn.go:30
	}
//line /usr/local/go/src/net/rawconn.go:30
	// _ = "end of CoverTab[16066]"
//line /usr/local/go/src/net/rawconn.go:30
	_go_fuzz_dep_.CoverTab[16067]++
						err := c.fd.pfd.RawControl(f)
						runtime.KeepAlive(c.fd)
						if err != nil {
//line /usr/local/go/src/net/rawconn.go:33
		_go_fuzz_dep_.CoverTab[16071]++
							err = &OpError{Op: "raw-control", Net: c.fd.net, Source: nil, Addr: c.fd.laddr, Err: err}
//line /usr/local/go/src/net/rawconn.go:34
		// _ = "end of CoverTab[16071]"
	} else {
//line /usr/local/go/src/net/rawconn.go:35
		_go_fuzz_dep_.CoverTab[16072]++
//line /usr/local/go/src/net/rawconn.go:35
		// _ = "end of CoverTab[16072]"
//line /usr/local/go/src/net/rawconn.go:35
	}
//line /usr/local/go/src/net/rawconn.go:35
	// _ = "end of CoverTab[16067]"
//line /usr/local/go/src/net/rawconn.go:35
	_go_fuzz_dep_.CoverTab[16068]++
						return err
//line /usr/local/go/src/net/rawconn.go:36
	// _ = "end of CoverTab[16068]"
}

func (c *rawConn) Read(f func(uintptr) bool) error {
//line /usr/local/go/src/net/rawconn.go:39
	_go_fuzz_dep_.CoverTab[16073]++
						if !c.ok() {
//line /usr/local/go/src/net/rawconn.go:40
		_go_fuzz_dep_.CoverTab[16076]++
							return syscall.EINVAL
//line /usr/local/go/src/net/rawconn.go:41
		// _ = "end of CoverTab[16076]"
	} else {
//line /usr/local/go/src/net/rawconn.go:42
		_go_fuzz_dep_.CoverTab[16077]++
//line /usr/local/go/src/net/rawconn.go:42
		// _ = "end of CoverTab[16077]"
//line /usr/local/go/src/net/rawconn.go:42
	}
//line /usr/local/go/src/net/rawconn.go:42
	// _ = "end of CoverTab[16073]"
//line /usr/local/go/src/net/rawconn.go:42
	_go_fuzz_dep_.CoverTab[16074]++
						err := c.fd.pfd.RawRead(f)
						runtime.KeepAlive(c.fd)
						if err != nil {
//line /usr/local/go/src/net/rawconn.go:45
		_go_fuzz_dep_.CoverTab[16078]++
							err = &OpError{Op: "raw-read", Net: c.fd.net, Source: c.fd.laddr, Addr: c.fd.raddr, Err: err}
//line /usr/local/go/src/net/rawconn.go:46
		// _ = "end of CoverTab[16078]"
	} else {
//line /usr/local/go/src/net/rawconn.go:47
		_go_fuzz_dep_.CoverTab[16079]++
//line /usr/local/go/src/net/rawconn.go:47
		// _ = "end of CoverTab[16079]"
//line /usr/local/go/src/net/rawconn.go:47
	}
//line /usr/local/go/src/net/rawconn.go:47
	// _ = "end of CoverTab[16074]"
//line /usr/local/go/src/net/rawconn.go:47
	_go_fuzz_dep_.CoverTab[16075]++
						return err
//line /usr/local/go/src/net/rawconn.go:48
	// _ = "end of CoverTab[16075]"
}

func (c *rawConn) Write(f func(uintptr) bool) error {
//line /usr/local/go/src/net/rawconn.go:51
	_go_fuzz_dep_.CoverTab[16080]++
						if !c.ok() {
//line /usr/local/go/src/net/rawconn.go:52
		_go_fuzz_dep_.CoverTab[16083]++
							return syscall.EINVAL
//line /usr/local/go/src/net/rawconn.go:53
		// _ = "end of CoverTab[16083]"
	} else {
//line /usr/local/go/src/net/rawconn.go:54
		_go_fuzz_dep_.CoverTab[16084]++
//line /usr/local/go/src/net/rawconn.go:54
		// _ = "end of CoverTab[16084]"
//line /usr/local/go/src/net/rawconn.go:54
	}
//line /usr/local/go/src/net/rawconn.go:54
	// _ = "end of CoverTab[16080]"
//line /usr/local/go/src/net/rawconn.go:54
	_go_fuzz_dep_.CoverTab[16081]++
						err := c.fd.pfd.RawWrite(f)
						runtime.KeepAlive(c.fd)
						if err != nil {
//line /usr/local/go/src/net/rawconn.go:57
		_go_fuzz_dep_.CoverTab[16085]++
							err = &OpError{Op: "raw-write", Net: c.fd.net, Source: c.fd.laddr, Addr: c.fd.raddr, Err: err}
//line /usr/local/go/src/net/rawconn.go:58
		// _ = "end of CoverTab[16085]"
	} else {
//line /usr/local/go/src/net/rawconn.go:59
		_go_fuzz_dep_.CoverTab[16086]++
//line /usr/local/go/src/net/rawconn.go:59
		// _ = "end of CoverTab[16086]"
//line /usr/local/go/src/net/rawconn.go:59
	}
//line /usr/local/go/src/net/rawconn.go:59
	// _ = "end of CoverTab[16081]"
//line /usr/local/go/src/net/rawconn.go:59
	_go_fuzz_dep_.CoverTab[16082]++
						return err
//line /usr/local/go/src/net/rawconn.go:60
	// _ = "end of CoverTab[16082]"
}

func newRawConn(fd *netFD) (*rawConn, error) {
//line /usr/local/go/src/net/rawconn.go:63
	_go_fuzz_dep_.CoverTab[16087]++
						return &rawConn{fd: fd}, nil
//line /usr/local/go/src/net/rawconn.go:64
	// _ = "end of CoverTab[16087]"
}

type rawListener struct {
	rawConn
}

func (l *rawListener) Read(func(uintptr) bool) error {
//line /usr/local/go/src/net/rawconn.go:71
	_go_fuzz_dep_.CoverTab[16088]++
						return syscall.EINVAL
//line /usr/local/go/src/net/rawconn.go:72
	// _ = "end of CoverTab[16088]"
}

func (l *rawListener) Write(func(uintptr) bool) error {
//line /usr/local/go/src/net/rawconn.go:75
	_go_fuzz_dep_.CoverTab[16089]++
						return syscall.EINVAL
//line /usr/local/go/src/net/rawconn.go:76
	// _ = "end of CoverTab[16089]"
}

func newRawListener(fd *netFD) (*rawListener, error) {
//line /usr/local/go/src/net/rawconn.go:79
	_go_fuzz_dep_.CoverTab[16090]++
						return &rawListener{rawConn{fd: fd}}, nil
//line /usr/local/go/src/net/rawconn.go:80
	// _ = "end of CoverTab[16090]"
}

//line /usr/local/go/src/net/rawconn.go:81
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/rawconn.go:81
var _ = _go_fuzz_dep_.CoverTab
