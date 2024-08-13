// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/net/sockopt_linux.go:5
package net

//line /usr/local/go/src/net/sockopt_linux.go:5
import (
//line /usr/local/go/src/net/sockopt_linux.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/sockopt_linux.go:5
)
//line /usr/local/go/src/net/sockopt_linux.go:5
import (
//line /usr/local/go/src/net/sockopt_linux.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/sockopt_linux.go:5
)

import (
	"os"
	"syscall"
)

func setDefaultSockopts(s, family, sotype int, ipv6only bool) error {
//line /usr/local/go/src/net/sockopt_linux.go:12
	_go_fuzz_dep_.CoverTab[7891]++
							if family == syscall.AF_INET6 && func() bool {
//line /usr/local/go/src/net/sockopt_linux.go:13
		_go_fuzz_dep_.CoverTab[7894]++
//line /usr/local/go/src/net/sockopt_linux.go:13
		return sotype != syscall.SOCK_RAW
//line /usr/local/go/src/net/sockopt_linux.go:13
		// _ = "end of CoverTab[7894]"
//line /usr/local/go/src/net/sockopt_linux.go:13
	}() {
//line /usr/local/go/src/net/sockopt_linux.go:13
		_go_fuzz_dep_.CoverTab[7895]++

//line /usr/local/go/src/net/sockopt_linux.go:17
		syscall.SetsockoptInt(s, syscall.IPPROTO_IPV6, syscall.IPV6_V6ONLY, boolint(ipv6only))
//line /usr/local/go/src/net/sockopt_linux.go:17
		// _ = "end of CoverTab[7895]"
	} else {
//line /usr/local/go/src/net/sockopt_linux.go:18
		_go_fuzz_dep_.CoverTab[7896]++
//line /usr/local/go/src/net/sockopt_linux.go:18
		// _ = "end of CoverTab[7896]"
//line /usr/local/go/src/net/sockopt_linux.go:18
	}
//line /usr/local/go/src/net/sockopt_linux.go:18
	// _ = "end of CoverTab[7891]"
//line /usr/local/go/src/net/sockopt_linux.go:18
	_go_fuzz_dep_.CoverTab[7892]++
							if (sotype == syscall.SOCK_DGRAM || func() bool {
//line /usr/local/go/src/net/sockopt_linux.go:19
		_go_fuzz_dep_.CoverTab[7897]++
//line /usr/local/go/src/net/sockopt_linux.go:19
		return sotype == syscall.SOCK_RAW
//line /usr/local/go/src/net/sockopt_linux.go:19
		// _ = "end of CoverTab[7897]"
//line /usr/local/go/src/net/sockopt_linux.go:19
	}()) && func() bool {
//line /usr/local/go/src/net/sockopt_linux.go:19
		_go_fuzz_dep_.CoverTab[7898]++
//line /usr/local/go/src/net/sockopt_linux.go:19
		return family != syscall.AF_UNIX
//line /usr/local/go/src/net/sockopt_linux.go:19
		// _ = "end of CoverTab[7898]"
//line /usr/local/go/src/net/sockopt_linux.go:19
	}() {
//line /usr/local/go/src/net/sockopt_linux.go:19
		_go_fuzz_dep_.CoverTab[7899]++

								return os.NewSyscallError("setsockopt", syscall.SetsockoptInt(s, syscall.SOL_SOCKET, syscall.SO_BROADCAST, 1))
//line /usr/local/go/src/net/sockopt_linux.go:21
		// _ = "end of CoverTab[7899]"
	} else {
//line /usr/local/go/src/net/sockopt_linux.go:22
		_go_fuzz_dep_.CoverTab[7900]++
//line /usr/local/go/src/net/sockopt_linux.go:22
		// _ = "end of CoverTab[7900]"
//line /usr/local/go/src/net/sockopt_linux.go:22
	}
//line /usr/local/go/src/net/sockopt_linux.go:22
	// _ = "end of CoverTab[7892]"
//line /usr/local/go/src/net/sockopt_linux.go:22
	_go_fuzz_dep_.CoverTab[7893]++
							return nil
//line /usr/local/go/src/net/sockopt_linux.go:23
	// _ = "end of CoverTab[7893]"
}

func setDefaultListenerSockopts(s int) error {
//line /usr/local/go/src/net/sockopt_linux.go:26
	_go_fuzz_dep_.CoverTab[7901]++

							return os.NewSyscallError("setsockopt", syscall.SetsockoptInt(s, syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 1))
//line /usr/local/go/src/net/sockopt_linux.go:28
	// _ = "end of CoverTab[7901]"
}

func setDefaultMulticastSockopts(s int) error {
//line /usr/local/go/src/net/sockopt_linux.go:31
	_go_fuzz_dep_.CoverTab[7902]++

//line /usr/local/go/src/net/sockopt_linux.go:34
	return os.NewSyscallError("setsockopt", syscall.SetsockoptInt(s, syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 1))
//line /usr/local/go/src/net/sockopt_linux.go:34
	// _ = "end of CoverTab[7902]"
}

//line /usr/local/go/src/net/sockopt_linux.go:35
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/sockopt_linux.go:35
var _ = _go_fuzz_dep_.CoverTab
