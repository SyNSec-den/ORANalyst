// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix

//line /snap/go/10455/src/net/writev_unix.go:7
package net

//line /snap/go/10455/src/net/writev_unix.go:7
import (
//line /snap/go/10455/src/net/writev_unix.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /snap/go/10455/src/net/writev_unix.go:7
)
//line /snap/go/10455/src/net/writev_unix.go:7
import (
//line /snap/go/10455/src/net/writev_unix.go:7
	_atomic_ "sync/atomic"
//line /snap/go/10455/src/net/writev_unix.go:7
)

import (
	"runtime"
	"syscall"
)

func (c *conn) writeBuffers(v *Buffers) (int64, error) {
//line /snap/go/10455/src/net/writev_unix.go:14
	_go_fuzz_dep_.CoverTab[9118]++
							if !c.ok() {
//line /snap/go/10455/src/net/writev_unix.go:15
		_go_fuzz_dep_.CoverTab[530260]++
//line /snap/go/10455/src/net/writev_unix.go:15
		_go_fuzz_dep_.CoverTab[9121]++
								return 0, syscall.EINVAL
//line /snap/go/10455/src/net/writev_unix.go:16
		// _ = "end of CoverTab[9121]"
	} else {
//line /snap/go/10455/src/net/writev_unix.go:17
		_go_fuzz_dep_.CoverTab[530261]++
//line /snap/go/10455/src/net/writev_unix.go:17
		_go_fuzz_dep_.CoverTab[9122]++
//line /snap/go/10455/src/net/writev_unix.go:17
		// _ = "end of CoverTab[9122]"
//line /snap/go/10455/src/net/writev_unix.go:17
	}
//line /snap/go/10455/src/net/writev_unix.go:17
	// _ = "end of CoverTab[9118]"
//line /snap/go/10455/src/net/writev_unix.go:17
	_go_fuzz_dep_.CoverTab[9119]++
							n, err := c.fd.writeBuffers(v)
							if err != nil {
//line /snap/go/10455/src/net/writev_unix.go:19
		_go_fuzz_dep_.CoverTab[530262]++
//line /snap/go/10455/src/net/writev_unix.go:19
		_go_fuzz_dep_.CoverTab[9123]++
								return n, &OpError{Op: "writev", Net: c.fd.net, Source: c.fd.laddr, Addr: c.fd.raddr, Err: err}
//line /snap/go/10455/src/net/writev_unix.go:20
		// _ = "end of CoverTab[9123]"
	} else {
//line /snap/go/10455/src/net/writev_unix.go:21
		_go_fuzz_dep_.CoverTab[530263]++
//line /snap/go/10455/src/net/writev_unix.go:21
		_go_fuzz_dep_.CoverTab[9124]++
//line /snap/go/10455/src/net/writev_unix.go:21
		// _ = "end of CoverTab[9124]"
//line /snap/go/10455/src/net/writev_unix.go:21
	}
//line /snap/go/10455/src/net/writev_unix.go:21
	// _ = "end of CoverTab[9119]"
//line /snap/go/10455/src/net/writev_unix.go:21
	_go_fuzz_dep_.CoverTab[9120]++
							return n, nil
//line /snap/go/10455/src/net/writev_unix.go:22
	// _ = "end of CoverTab[9120]"
}

func (fd *netFD) writeBuffers(v *Buffers) (n int64, err error) {
//line /snap/go/10455/src/net/writev_unix.go:25
	_go_fuzz_dep_.CoverTab[9125]++
							n, err = fd.pfd.Writev((*[][]byte)(v))
							runtime.KeepAlive(fd)
							return n, wrapSyscallError("writev", err)
//line /snap/go/10455/src/net/writev_unix.go:28
	// _ = "end of CoverTab[9125]"
}

//line /snap/go/10455/src/net/writev_unix.go:29
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /snap/go/10455/src/net/writev_unix.go:29
var _ = _go_fuzz_dep_.CoverTab
