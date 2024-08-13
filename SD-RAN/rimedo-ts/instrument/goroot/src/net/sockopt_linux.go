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
	_go_fuzz_dep_.CoverTab[16281]++
							if family == syscall.AF_INET6 && func() bool {
//line /usr/local/go/src/net/sockopt_linux.go:13
		_go_fuzz_dep_.CoverTab[16284]++
//line /usr/local/go/src/net/sockopt_linux.go:13
		return sotype != syscall.SOCK_RAW
//line /usr/local/go/src/net/sockopt_linux.go:13
		// _ = "end of CoverTab[16284]"
//line /usr/local/go/src/net/sockopt_linux.go:13
	}() {
//line /usr/local/go/src/net/sockopt_linux.go:13
		_go_fuzz_dep_.CoverTab[16285]++

//line /usr/local/go/src/net/sockopt_linux.go:17
		syscall.SetsockoptInt(s, syscall.IPPROTO_IPV6, syscall.IPV6_V6ONLY, boolint(ipv6only))
//line /usr/local/go/src/net/sockopt_linux.go:17
		// _ = "end of CoverTab[16285]"
	} else {
//line /usr/local/go/src/net/sockopt_linux.go:18
		_go_fuzz_dep_.CoverTab[16286]++
//line /usr/local/go/src/net/sockopt_linux.go:18
		// _ = "end of CoverTab[16286]"
//line /usr/local/go/src/net/sockopt_linux.go:18
	}
//line /usr/local/go/src/net/sockopt_linux.go:18
	// _ = "end of CoverTab[16281]"
//line /usr/local/go/src/net/sockopt_linux.go:18
	_go_fuzz_dep_.CoverTab[16282]++
							if (sotype == syscall.SOCK_DGRAM || func() bool {
//line /usr/local/go/src/net/sockopt_linux.go:19
		_go_fuzz_dep_.CoverTab[16287]++
//line /usr/local/go/src/net/sockopt_linux.go:19
		return sotype == syscall.SOCK_RAW
//line /usr/local/go/src/net/sockopt_linux.go:19
		// _ = "end of CoverTab[16287]"
//line /usr/local/go/src/net/sockopt_linux.go:19
	}()) && func() bool {
//line /usr/local/go/src/net/sockopt_linux.go:19
		_go_fuzz_dep_.CoverTab[16288]++
//line /usr/local/go/src/net/sockopt_linux.go:19
		return family != syscall.AF_UNIX
//line /usr/local/go/src/net/sockopt_linux.go:19
		// _ = "end of CoverTab[16288]"
//line /usr/local/go/src/net/sockopt_linux.go:19
	}() {
//line /usr/local/go/src/net/sockopt_linux.go:19
		_go_fuzz_dep_.CoverTab[16289]++

								return os.NewSyscallError("setsockopt", syscall.SetsockoptInt(s, syscall.SOL_SOCKET, syscall.SO_BROADCAST, 1))
//line /usr/local/go/src/net/sockopt_linux.go:21
		// _ = "end of CoverTab[16289]"
	} else {
//line /usr/local/go/src/net/sockopt_linux.go:22
		_go_fuzz_dep_.CoverTab[16290]++
//line /usr/local/go/src/net/sockopt_linux.go:22
		// _ = "end of CoverTab[16290]"
//line /usr/local/go/src/net/sockopt_linux.go:22
	}
//line /usr/local/go/src/net/sockopt_linux.go:22
	// _ = "end of CoverTab[16282]"
//line /usr/local/go/src/net/sockopt_linux.go:22
	_go_fuzz_dep_.CoverTab[16283]++
							return nil
//line /usr/local/go/src/net/sockopt_linux.go:23
	// _ = "end of CoverTab[16283]"
}

func setDefaultListenerSockopts(s int) error {
//line /usr/local/go/src/net/sockopt_linux.go:26
	_go_fuzz_dep_.CoverTab[16291]++

							return os.NewSyscallError("setsockopt", syscall.SetsockoptInt(s, syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 1))
//line /usr/local/go/src/net/sockopt_linux.go:28
	// _ = "end of CoverTab[16291]"
}

func setDefaultMulticastSockopts(s int) error {
//line /usr/local/go/src/net/sockopt_linux.go:31
	_go_fuzz_dep_.CoverTab[16292]++

//line /usr/local/go/src/net/sockopt_linux.go:34
	return os.NewSyscallError("setsockopt", syscall.SetsockoptInt(s, syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 1))
//line /usr/local/go/src/net/sockopt_linux.go:34
	// _ = "end of CoverTab[16292]"
}

//line /usr/local/go/src/net/sockopt_linux.go:35
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/sockopt_linux.go:35
var _ = _go_fuzz_dep_.CoverTab
