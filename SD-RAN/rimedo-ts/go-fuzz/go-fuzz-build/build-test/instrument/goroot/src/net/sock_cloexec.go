// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements sysSocket for platforms that provide a fast path for
// setting SetNonblock and CloseOnExec.

//go:build dragonfly || freebsd || linux || netbsd || openbsd || solaris

//line /usr/local/go/src/net/sock_cloexec.go:10
package net

//line /usr/local/go/src/net/sock_cloexec.go:10
import (
//line /usr/local/go/src/net/sock_cloexec.go:10
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/sock_cloexec.go:10
)
//line /usr/local/go/src/net/sock_cloexec.go:10
import (
//line /usr/local/go/src/net/sock_cloexec.go:10
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/sock_cloexec.go:10
)

import (
	"os"
	"syscall"
)

// Wrapper around the socket system call that marks the returned file
//line /usr/local/go/src/net/sock_cloexec.go:17
// descriptor as nonblocking and close-on-exec.
//line /usr/local/go/src/net/sock_cloexec.go:19
func sysSocket(family, sotype, proto int) (int, error) {
//line /usr/local/go/src/net/sock_cloexec.go:19
	_go_fuzz_dep_.CoverTab[7721]++
							s, err := socketFunc(family, sotype|syscall.SOCK_NONBLOCK|syscall.SOCK_CLOEXEC, proto)
							if err != nil {
//line /usr/local/go/src/net/sock_cloexec.go:21
		_go_fuzz_dep_.CoverTab[7723]++
								return -1, os.NewSyscallError("socket", err)
//line /usr/local/go/src/net/sock_cloexec.go:22
		// _ = "end of CoverTab[7723]"
	} else {
//line /usr/local/go/src/net/sock_cloexec.go:23
		_go_fuzz_dep_.CoverTab[7724]++
//line /usr/local/go/src/net/sock_cloexec.go:23
		// _ = "end of CoverTab[7724]"
//line /usr/local/go/src/net/sock_cloexec.go:23
	}
//line /usr/local/go/src/net/sock_cloexec.go:23
	// _ = "end of CoverTab[7721]"
//line /usr/local/go/src/net/sock_cloexec.go:23
	_go_fuzz_dep_.CoverTab[7722]++
							return s, nil
//line /usr/local/go/src/net/sock_cloexec.go:24
	// _ = "end of CoverTab[7722]"
}

//line /usr/local/go/src/net/sock_cloexec.go:25
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/sock_cloexec.go:25
var _ = _go_fuzz_dep_.CoverTab
