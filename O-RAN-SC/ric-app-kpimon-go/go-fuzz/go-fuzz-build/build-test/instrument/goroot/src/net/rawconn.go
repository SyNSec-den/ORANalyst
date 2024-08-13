// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /snap/go/10455/src/net/rawconn.go:5
package net

//line /snap/go/10455/src/net/rawconn.go:5
import (
//line /snap/go/10455/src/net/rawconn.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /snap/go/10455/src/net/rawconn.go:5
)
//line /snap/go/10455/src/net/rawconn.go:5
import (
//line /snap/go/10455/src/net/rawconn.go:5
	_atomic_ "sync/atomic"
//line /snap/go/10455/src/net/rawconn.go:5
)

import (
	"internal/poll"
	"runtime"
	"syscall"
)

//line /snap/go/10455/src/net/rawconn.go:22
type rawConn struct {
	fd *netFD
}

func (c *rawConn) ok() bool {
//line /snap/go/10455/src/net/rawconn.go:26
	_go_fuzz_dep_.CoverTab[7954]++
//line /snap/go/10455/src/net/rawconn.go:26
	return c != nil && func() bool {
//line /snap/go/10455/src/net/rawconn.go:26
		_go_fuzz_dep_.CoverTab[7955]++
//line /snap/go/10455/src/net/rawconn.go:26
		return c.fd != nil
//line /snap/go/10455/src/net/rawconn.go:26
		// _ = "end of CoverTab[7955]"
//line /snap/go/10455/src/net/rawconn.go:26
	}()
//line /snap/go/10455/src/net/rawconn.go:26
	// _ = "end of CoverTab[7954]"
//line /snap/go/10455/src/net/rawconn.go:26
}

func (c *rawConn) Control(f func(uintptr)) error {
//line /snap/go/10455/src/net/rawconn.go:28
	_go_fuzz_dep_.CoverTab[7956]++
						if !c.ok() {
//line /snap/go/10455/src/net/rawconn.go:29
		_go_fuzz_dep_.CoverTab[529582]++
//line /snap/go/10455/src/net/rawconn.go:29
		_go_fuzz_dep_.CoverTab[7959]++
							return syscall.EINVAL
//line /snap/go/10455/src/net/rawconn.go:30
		// _ = "end of CoverTab[7959]"
	} else {
//line /snap/go/10455/src/net/rawconn.go:31
		_go_fuzz_dep_.CoverTab[529583]++
//line /snap/go/10455/src/net/rawconn.go:31
		_go_fuzz_dep_.CoverTab[7960]++
//line /snap/go/10455/src/net/rawconn.go:31
		// _ = "end of CoverTab[7960]"
//line /snap/go/10455/src/net/rawconn.go:31
	}
//line /snap/go/10455/src/net/rawconn.go:31
	// _ = "end of CoverTab[7956]"
//line /snap/go/10455/src/net/rawconn.go:31
	_go_fuzz_dep_.CoverTab[7957]++
						err := c.fd.pfd.RawControl(f)
						runtime.KeepAlive(c.fd)
						if err != nil {
//line /snap/go/10455/src/net/rawconn.go:34
		_go_fuzz_dep_.CoverTab[529584]++
//line /snap/go/10455/src/net/rawconn.go:34
		_go_fuzz_dep_.CoverTab[7961]++
							err = &OpError{Op: "raw-control", Net: c.fd.net, Source: nil, Addr: c.fd.laddr, Err: err}
//line /snap/go/10455/src/net/rawconn.go:35
		// _ = "end of CoverTab[7961]"
	} else {
//line /snap/go/10455/src/net/rawconn.go:36
		_go_fuzz_dep_.CoverTab[529585]++
//line /snap/go/10455/src/net/rawconn.go:36
		_go_fuzz_dep_.CoverTab[7962]++
//line /snap/go/10455/src/net/rawconn.go:36
		// _ = "end of CoverTab[7962]"
//line /snap/go/10455/src/net/rawconn.go:36
	}
//line /snap/go/10455/src/net/rawconn.go:36
	// _ = "end of CoverTab[7957]"
//line /snap/go/10455/src/net/rawconn.go:36
	_go_fuzz_dep_.CoverTab[7958]++
						return err
//line /snap/go/10455/src/net/rawconn.go:37
	// _ = "end of CoverTab[7958]"
}

func (c *rawConn) Read(f func(uintptr) bool) error {
//line /snap/go/10455/src/net/rawconn.go:40
	_go_fuzz_dep_.CoverTab[7963]++
						if !c.ok() {
//line /snap/go/10455/src/net/rawconn.go:41
		_go_fuzz_dep_.CoverTab[529586]++
//line /snap/go/10455/src/net/rawconn.go:41
		_go_fuzz_dep_.CoverTab[7966]++
							return syscall.EINVAL
//line /snap/go/10455/src/net/rawconn.go:42
		// _ = "end of CoverTab[7966]"
	} else {
//line /snap/go/10455/src/net/rawconn.go:43
		_go_fuzz_dep_.CoverTab[529587]++
//line /snap/go/10455/src/net/rawconn.go:43
		_go_fuzz_dep_.CoverTab[7967]++
//line /snap/go/10455/src/net/rawconn.go:43
		// _ = "end of CoverTab[7967]"
//line /snap/go/10455/src/net/rawconn.go:43
	}
//line /snap/go/10455/src/net/rawconn.go:43
	// _ = "end of CoverTab[7963]"
//line /snap/go/10455/src/net/rawconn.go:43
	_go_fuzz_dep_.CoverTab[7964]++
						err := c.fd.pfd.RawRead(f)
						runtime.KeepAlive(c.fd)
						if err != nil {
//line /snap/go/10455/src/net/rawconn.go:46
		_go_fuzz_dep_.CoverTab[529588]++
//line /snap/go/10455/src/net/rawconn.go:46
		_go_fuzz_dep_.CoverTab[7968]++
							err = &OpError{Op: "raw-read", Net: c.fd.net, Source: c.fd.laddr, Addr: c.fd.raddr, Err: err}
//line /snap/go/10455/src/net/rawconn.go:47
		// _ = "end of CoverTab[7968]"
	} else {
//line /snap/go/10455/src/net/rawconn.go:48
		_go_fuzz_dep_.CoverTab[529589]++
//line /snap/go/10455/src/net/rawconn.go:48
		_go_fuzz_dep_.CoverTab[7969]++
//line /snap/go/10455/src/net/rawconn.go:48
		// _ = "end of CoverTab[7969]"
//line /snap/go/10455/src/net/rawconn.go:48
	}
//line /snap/go/10455/src/net/rawconn.go:48
	// _ = "end of CoverTab[7964]"
//line /snap/go/10455/src/net/rawconn.go:48
	_go_fuzz_dep_.CoverTab[7965]++
						return err
//line /snap/go/10455/src/net/rawconn.go:49
	// _ = "end of CoverTab[7965]"
}

func (c *rawConn) Write(f func(uintptr) bool) error {
//line /snap/go/10455/src/net/rawconn.go:52
	_go_fuzz_dep_.CoverTab[7970]++
						if !c.ok() {
//line /snap/go/10455/src/net/rawconn.go:53
		_go_fuzz_dep_.CoverTab[529590]++
//line /snap/go/10455/src/net/rawconn.go:53
		_go_fuzz_dep_.CoverTab[7973]++
							return syscall.EINVAL
//line /snap/go/10455/src/net/rawconn.go:54
		// _ = "end of CoverTab[7973]"
	} else {
//line /snap/go/10455/src/net/rawconn.go:55
		_go_fuzz_dep_.CoverTab[529591]++
//line /snap/go/10455/src/net/rawconn.go:55
		_go_fuzz_dep_.CoverTab[7974]++
//line /snap/go/10455/src/net/rawconn.go:55
		// _ = "end of CoverTab[7974]"
//line /snap/go/10455/src/net/rawconn.go:55
	}
//line /snap/go/10455/src/net/rawconn.go:55
	// _ = "end of CoverTab[7970]"
//line /snap/go/10455/src/net/rawconn.go:55
	_go_fuzz_dep_.CoverTab[7971]++
						err := c.fd.pfd.RawWrite(f)
						runtime.KeepAlive(c.fd)
						if err != nil {
//line /snap/go/10455/src/net/rawconn.go:58
		_go_fuzz_dep_.CoverTab[529592]++
//line /snap/go/10455/src/net/rawconn.go:58
		_go_fuzz_dep_.CoverTab[7975]++
							err = &OpError{Op: "raw-write", Net: c.fd.net, Source: c.fd.laddr, Addr: c.fd.raddr, Err: err}
//line /snap/go/10455/src/net/rawconn.go:59
		// _ = "end of CoverTab[7975]"
	} else {
//line /snap/go/10455/src/net/rawconn.go:60
		_go_fuzz_dep_.CoverTab[529593]++
//line /snap/go/10455/src/net/rawconn.go:60
		_go_fuzz_dep_.CoverTab[7976]++
//line /snap/go/10455/src/net/rawconn.go:60
		// _ = "end of CoverTab[7976]"
//line /snap/go/10455/src/net/rawconn.go:60
	}
//line /snap/go/10455/src/net/rawconn.go:60
	// _ = "end of CoverTab[7971]"
//line /snap/go/10455/src/net/rawconn.go:60
	_go_fuzz_dep_.CoverTab[7972]++
						return err
//line /snap/go/10455/src/net/rawconn.go:61
	// _ = "end of CoverTab[7972]"
}

// PollFD returns the poll.FD of the underlying connection.
//line /snap/go/10455/src/net/rawconn.go:64
//
//line /snap/go/10455/src/net/rawconn.go:64
// Other packages in std that also import internal/poll (such as os)
//line /snap/go/10455/src/net/rawconn.go:64
// can use a type assertion to access this extension method so that
//line /snap/go/10455/src/net/rawconn.go:64
// they can pass the *poll.FD to functions like poll.Splice.
//line /snap/go/10455/src/net/rawconn.go:64
//
//line /snap/go/10455/src/net/rawconn.go:64
// PollFD is not intended for use outside the standard library.
//line /snap/go/10455/src/net/rawconn.go:71
func (c *rawConn) PollFD() *poll.FD {
//line /snap/go/10455/src/net/rawconn.go:71
	_go_fuzz_dep_.CoverTab[7977]++
						if !c.ok() {
//line /snap/go/10455/src/net/rawconn.go:72
		_go_fuzz_dep_.CoverTab[529594]++
//line /snap/go/10455/src/net/rawconn.go:72
		_go_fuzz_dep_.CoverTab[7979]++
							return nil
//line /snap/go/10455/src/net/rawconn.go:73
		// _ = "end of CoverTab[7979]"
	} else {
//line /snap/go/10455/src/net/rawconn.go:74
		_go_fuzz_dep_.CoverTab[529595]++
//line /snap/go/10455/src/net/rawconn.go:74
		_go_fuzz_dep_.CoverTab[7980]++
//line /snap/go/10455/src/net/rawconn.go:74
		// _ = "end of CoverTab[7980]"
//line /snap/go/10455/src/net/rawconn.go:74
	}
//line /snap/go/10455/src/net/rawconn.go:74
	// _ = "end of CoverTab[7977]"
//line /snap/go/10455/src/net/rawconn.go:74
	_go_fuzz_dep_.CoverTab[7978]++
						return &c.fd.pfd
//line /snap/go/10455/src/net/rawconn.go:75
	// _ = "end of CoverTab[7978]"
}

func newRawConn(fd *netFD) (*rawConn, error) {
//line /snap/go/10455/src/net/rawconn.go:78
	_go_fuzz_dep_.CoverTab[7981]++
						return &rawConn{fd: fd}, nil
//line /snap/go/10455/src/net/rawconn.go:79
	// _ = "end of CoverTab[7981]"
}

type rawListener struct {
	rawConn
}

func (l *rawListener) Read(func(uintptr) bool) error {
//line /snap/go/10455/src/net/rawconn.go:86
	_go_fuzz_dep_.CoverTab[7982]++
						return syscall.EINVAL
//line /snap/go/10455/src/net/rawconn.go:87
	// _ = "end of CoverTab[7982]"
}

func (l *rawListener) Write(func(uintptr) bool) error {
//line /snap/go/10455/src/net/rawconn.go:90
	_go_fuzz_dep_.CoverTab[7983]++
						return syscall.EINVAL
//line /snap/go/10455/src/net/rawconn.go:91
	// _ = "end of CoverTab[7983]"
}

func newRawListener(fd *netFD) (*rawListener, error) {
//line /snap/go/10455/src/net/rawconn.go:94
	_go_fuzz_dep_.CoverTab[7984]++
						return &rawListener{rawConn{fd: fd}}, nil
//line /snap/go/10455/src/net/rawconn.go:95
	// _ = "end of CoverTab[7984]"
}

//line /snap/go/10455/src/net/rawconn.go:96
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /snap/go/10455/src/net/rawconn.go:96
var _ = _go_fuzz_dep_.CoverTab
