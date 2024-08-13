// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/net/sockoptip_linux.go:5
package net

//line /usr/local/go/src/net/sockoptip_linux.go:5
import (
//line /usr/local/go/src/net/sockoptip_linux.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/sockoptip_linux.go:5
)
//line /usr/local/go/src/net/sockoptip_linux.go:5
import (
//line /usr/local/go/src/net/sockoptip_linux.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/sockoptip_linux.go:5
)

import (
	"runtime"
	"syscall"
)

func setIPv4MulticastInterface(fd *netFD, ifi *Interface) error {
//line /usr/local/go/src/net/sockoptip_linux.go:12
	_go_fuzz_dep_.CoverTab[16356]++
							var v int32
							if ifi != nil {
//line /usr/local/go/src/net/sockoptip_linux.go:14
		_go_fuzz_dep_.CoverTab[16358]++
								v = int32(ifi.Index)
//line /usr/local/go/src/net/sockoptip_linux.go:15
		// _ = "end of CoverTab[16358]"
	} else {
//line /usr/local/go/src/net/sockoptip_linux.go:16
		_go_fuzz_dep_.CoverTab[16359]++
//line /usr/local/go/src/net/sockoptip_linux.go:16
		// _ = "end of CoverTab[16359]"
//line /usr/local/go/src/net/sockoptip_linux.go:16
	}
//line /usr/local/go/src/net/sockoptip_linux.go:16
	// _ = "end of CoverTab[16356]"
//line /usr/local/go/src/net/sockoptip_linux.go:16
	_go_fuzz_dep_.CoverTab[16357]++
							mreq := &syscall.IPMreqn{Ifindex: v}
							err := fd.pfd.SetsockoptIPMreqn(syscall.IPPROTO_IP, syscall.IP_MULTICAST_IF, mreq)
							runtime.KeepAlive(fd)
							return wrapSyscallError("setsockopt", err)
//line /usr/local/go/src/net/sockoptip_linux.go:20
	// _ = "end of CoverTab[16357]"
}

func setIPv4MulticastLoopback(fd *netFD, v bool) error {
//line /usr/local/go/src/net/sockoptip_linux.go:23
	_go_fuzz_dep_.CoverTab[16360]++
							err := fd.pfd.SetsockoptInt(syscall.IPPROTO_IP, syscall.IP_MULTICAST_LOOP, boolint(v))
							runtime.KeepAlive(fd)
							return wrapSyscallError("setsockopt", err)
//line /usr/local/go/src/net/sockoptip_linux.go:26
	// _ = "end of CoverTab[16360]"
}

//line /usr/local/go/src/net/sockoptip_linux.go:27
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/sockoptip_linux.go:27
var _ = _go_fuzz_dep_.CoverTab
