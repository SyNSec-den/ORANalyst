// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /snap/go/10455/src/net/sockoptip_linux.go:5
package net

//line /snap/go/10455/src/net/sockoptip_linux.go:5
import (
//line /snap/go/10455/src/net/sockoptip_linux.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /snap/go/10455/src/net/sockoptip_linux.go:5
)
//line /snap/go/10455/src/net/sockoptip_linux.go:5
import (
//line /snap/go/10455/src/net/sockoptip_linux.go:5
	_atomic_ "sync/atomic"
//line /snap/go/10455/src/net/sockoptip_linux.go:5
)

import (
	"runtime"
	"syscall"
)

func setIPv4MulticastInterface(fd *netFD, ifi *Interface) error {
//line /snap/go/10455/src/net/sockoptip_linux.go:12
	_go_fuzz_dep_.CoverTab[8260]++
							var v int32
							if ifi != nil {
//line /snap/go/10455/src/net/sockoptip_linux.go:14
		_go_fuzz_dep_.CoverTab[529787]++
//line /snap/go/10455/src/net/sockoptip_linux.go:14
		_go_fuzz_dep_.CoverTab[8262]++
								v = int32(ifi.Index)
//line /snap/go/10455/src/net/sockoptip_linux.go:15
		// _ = "end of CoverTab[8262]"
	} else {
//line /snap/go/10455/src/net/sockoptip_linux.go:16
		_go_fuzz_dep_.CoverTab[529788]++
//line /snap/go/10455/src/net/sockoptip_linux.go:16
		_go_fuzz_dep_.CoverTab[8263]++
//line /snap/go/10455/src/net/sockoptip_linux.go:16
		// _ = "end of CoverTab[8263]"
//line /snap/go/10455/src/net/sockoptip_linux.go:16
	}
//line /snap/go/10455/src/net/sockoptip_linux.go:16
	// _ = "end of CoverTab[8260]"
//line /snap/go/10455/src/net/sockoptip_linux.go:16
	_go_fuzz_dep_.CoverTab[8261]++
							mreq := &syscall.IPMreqn{Ifindex: v}
							err := fd.pfd.SetsockoptIPMreqn(syscall.IPPROTO_IP, syscall.IP_MULTICAST_IF, mreq)
							runtime.KeepAlive(fd)
							return wrapSyscallError("setsockopt", err)
//line /snap/go/10455/src/net/sockoptip_linux.go:20
	// _ = "end of CoverTab[8261]"
}

func setIPv4MulticastLoopback(fd *netFD, v bool) error {
//line /snap/go/10455/src/net/sockoptip_linux.go:23
	_go_fuzz_dep_.CoverTab[8264]++
							err := fd.pfd.SetsockoptInt(syscall.IPPROTO_IP, syscall.IP_MULTICAST_LOOP, boolint(v))
							runtime.KeepAlive(fd)
							return wrapSyscallError("setsockopt", err)
//line /snap/go/10455/src/net/sockoptip_linux.go:26
	// _ = "end of CoverTab[8264]"
}

//line /snap/go/10455/src/net/sockoptip_linux.go:27
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /snap/go/10455/src/net/sockoptip_linux.go:27
var _ = _go_fuzz_dep_.CoverTab
