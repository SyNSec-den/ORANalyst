// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix

//line /usr/local/go/src/net/writev_unix.go:7
package net

//line /usr/local/go/src/net/writev_unix.go:7
import (
//line /usr/local/go/src/net/writev_unix.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/writev_unix.go:7
)
//line /usr/local/go/src/net/writev_unix.go:7
import (
//line /usr/local/go/src/net/writev_unix.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/writev_unix.go:7
)

import (
	"runtime"
	"syscall"
)

func (c *conn) writeBuffers(v *Buffers) (int64, error) {
//line /usr/local/go/src/net/writev_unix.go:14
	_go_fuzz_dep_.CoverTab[8818]++
						if !c.ok() {
//line /usr/local/go/src/net/writev_unix.go:15
		_go_fuzz_dep_.CoverTab[8821]++
							return 0, syscall.EINVAL
//line /usr/local/go/src/net/writev_unix.go:16
		// _ = "end of CoverTab[8821]"
	} else {
//line /usr/local/go/src/net/writev_unix.go:17
		_go_fuzz_dep_.CoverTab[8822]++
//line /usr/local/go/src/net/writev_unix.go:17
		// _ = "end of CoverTab[8822]"
//line /usr/local/go/src/net/writev_unix.go:17
	}
//line /usr/local/go/src/net/writev_unix.go:17
	// _ = "end of CoverTab[8818]"
//line /usr/local/go/src/net/writev_unix.go:17
	_go_fuzz_dep_.CoverTab[8819]++
						n, err := c.fd.writeBuffers(v)
						if err != nil {
//line /usr/local/go/src/net/writev_unix.go:19
		_go_fuzz_dep_.CoverTab[8823]++
							return n, &OpError{Op: "writev", Net: c.fd.net, Source: c.fd.laddr, Addr: c.fd.raddr, Err: err}
//line /usr/local/go/src/net/writev_unix.go:20
		// _ = "end of CoverTab[8823]"
	} else {
//line /usr/local/go/src/net/writev_unix.go:21
		_go_fuzz_dep_.CoverTab[8824]++
//line /usr/local/go/src/net/writev_unix.go:21
		// _ = "end of CoverTab[8824]"
//line /usr/local/go/src/net/writev_unix.go:21
	}
//line /usr/local/go/src/net/writev_unix.go:21
	// _ = "end of CoverTab[8819]"
//line /usr/local/go/src/net/writev_unix.go:21
	_go_fuzz_dep_.CoverTab[8820]++
						return n, nil
//line /usr/local/go/src/net/writev_unix.go:22
	// _ = "end of CoverTab[8820]"
}

func (fd *netFD) writeBuffers(v *Buffers) (n int64, err error) {
//line /usr/local/go/src/net/writev_unix.go:25
	_go_fuzz_dep_.CoverTab[8825]++
						n, err = fd.pfd.Writev((*[][]byte)(v))
						runtime.KeepAlive(fd)
						return n, wrapSyscallError("writev", err)
//line /usr/local/go/src/net/writev_unix.go:28
	// _ = "end of CoverTab[8825]"
}

//line /usr/local/go/src/net/writev_unix.go:29
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/writev_unix.go:29
var _ = _go_fuzz_dep_.CoverTab
