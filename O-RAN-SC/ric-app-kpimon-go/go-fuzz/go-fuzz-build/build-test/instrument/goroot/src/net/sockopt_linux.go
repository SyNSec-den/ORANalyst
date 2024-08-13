// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /snap/go/10455/src/net/sockopt_linux.go:5
package net

//line /snap/go/10455/src/net/sockopt_linux.go:5
import (
//line /snap/go/10455/src/net/sockopt_linux.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /snap/go/10455/src/net/sockopt_linux.go:5
)
//line /snap/go/10455/src/net/sockopt_linux.go:5
import (
//line /snap/go/10455/src/net/sockopt_linux.go:5
	_atomic_ "sync/atomic"
//line /snap/go/10455/src/net/sockopt_linux.go:5
)

import (
	"os"
	"syscall"
)

func setDefaultSockopts(s, family, sotype int, ipv6only bool) error {
//line /snap/go/10455/src/net/sockopt_linux.go:12
	_go_fuzz_dep_.CoverTab[8185]++
							if family == syscall.AF_INET6 && func() bool {
//line /snap/go/10455/src/net/sockopt_linux.go:13
		_go_fuzz_dep_.CoverTab[8188]++
//line /snap/go/10455/src/net/sockopt_linux.go:13
		return sotype != syscall.SOCK_RAW
//line /snap/go/10455/src/net/sockopt_linux.go:13
		// _ = "end of CoverTab[8188]"
//line /snap/go/10455/src/net/sockopt_linux.go:13
	}() {
//line /snap/go/10455/src/net/sockopt_linux.go:13
		_go_fuzz_dep_.CoverTab[529729]++
//line /snap/go/10455/src/net/sockopt_linux.go:13
		_go_fuzz_dep_.CoverTab[8189]++

//line /snap/go/10455/src/net/sockopt_linux.go:17
		syscall.SetsockoptInt(s, syscall.IPPROTO_IPV6, syscall.IPV6_V6ONLY, boolint(ipv6only))
//line /snap/go/10455/src/net/sockopt_linux.go:17
		// _ = "end of CoverTab[8189]"
	} else {
//line /snap/go/10455/src/net/sockopt_linux.go:18
		_go_fuzz_dep_.CoverTab[529730]++
//line /snap/go/10455/src/net/sockopt_linux.go:18
		_go_fuzz_dep_.CoverTab[8190]++
//line /snap/go/10455/src/net/sockopt_linux.go:18
		// _ = "end of CoverTab[8190]"
//line /snap/go/10455/src/net/sockopt_linux.go:18
	}
//line /snap/go/10455/src/net/sockopt_linux.go:18
	// _ = "end of CoverTab[8185]"
//line /snap/go/10455/src/net/sockopt_linux.go:18
	_go_fuzz_dep_.CoverTab[8186]++
							if (sotype == syscall.SOCK_DGRAM || func() bool {
//line /snap/go/10455/src/net/sockopt_linux.go:19
		_go_fuzz_dep_.CoverTab[8191]++
//line /snap/go/10455/src/net/sockopt_linux.go:19
		return sotype == syscall.SOCK_RAW
//line /snap/go/10455/src/net/sockopt_linux.go:19
		// _ = "end of CoverTab[8191]"
//line /snap/go/10455/src/net/sockopt_linux.go:19
	}()) && func() bool {
//line /snap/go/10455/src/net/sockopt_linux.go:19
		_go_fuzz_dep_.CoverTab[8192]++
//line /snap/go/10455/src/net/sockopt_linux.go:19
		return family != syscall.AF_UNIX
//line /snap/go/10455/src/net/sockopt_linux.go:19
		// _ = "end of CoverTab[8192]"
//line /snap/go/10455/src/net/sockopt_linux.go:19
	}() {
//line /snap/go/10455/src/net/sockopt_linux.go:19
		_go_fuzz_dep_.CoverTab[529731]++
//line /snap/go/10455/src/net/sockopt_linux.go:19
		_go_fuzz_dep_.CoverTab[8193]++

								return os.NewSyscallError("setsockopt", syscall.SetsockoptInt(s, syscall.SOL_SOCKET, syscall.SO_BROADCAST, 1))
//line /snap/go/10455/src/net/sockopt_linux.go:21
		// _ = "end of CoverTab[8193]"
	} else {
//line /snap/go/10455/src/net/sockopt_linux.go:22
		_go_fuzz_dep_.CoverTab[529732]++
//line /snap/go/10455/src/net/sockopt_linux.go:22
		_go_fuzz_dep_.CoverTab[8194]++
//line /snap/go/10455/src/net/sockopt_linux.go:22
		// _ = "end of CoverTab[8194]"
//line /snap/go/10455/src/net/sockopt_linux.go:22
	}
//line /snap/go/10455/src/net/sockopt_linux.go:22
	// _ = "end of CoverTab[8186]"
//line /snap/go/10455/src/net/sockopt_linux.go:22
	_go_fuzz_dep_.CoverTab[8187]++
							return nil
//line /snap/go/10455/src/net/sockopt_linux.go:23
	// _ = "end of CoverTab[8187]"
}

func setDefaultListenerSockopts(s int) error {
//line /snap/go/10455/src/net/sockopt_linux.go:26
	_go_fuzz_dep_.CoverTab[8195]++

							return os.NewSyscallError("setsockopt", syscall.SetsockoptInt(s, syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 1))
//line /snap/go/10455/src/net/sockopt_linux.go:28
	// _ = "end of CoverTab[8195]"
}

func setDefaultMulticastSockopts(s int) error {
//line /snap/go/10455/src/net/sockopt_linux.go:31
	_go_fuzz_dep_.CoverTab[8196]++

//line /snap/go/10455/src/net/sockopt_linux.go:34
	return os.NewSyscallError("setsockopt", syscall.SetsockoptInt(s, syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 1))
//line /snap/go/10455/src/net/sockopt_linux.go:34
	// _ = "end of CoverTab[8196]"
}

//line /snap/go/10455/src/net/sockopt_linux.go:35
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /snap/go/10455/src/net/sockopt_linux.go:35
var _ = _go_fuzz_dep_.CoverTab
