// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements sysSocket for platforms that provide a fast path for
// setting SetNonblock and CloseOnExec.

//go:build dragonfly || freebsd || linux || netbsd || openbsd || solaris

//line /snap/go/10455/src/net/sock_cloexec.go:10
package net

//line /snap/go/10455/src/net/sock_cloexec.go:10
import (
//line /snap/go/10455/src/net/sock_cloexec.go:10
	_go_fuzz_dep_ "go-fuzz-dep"
//line /snap/go/10455/src/net/sock_cloexec.go:10
)
//line /snap/go/10455/src/net/sock_cloexec.go:10
import (
//line /snap/go/10455/src/net/sock_cloexec.go:10
	_atomic_ "sync/atomic"
//line /snap/go/10455/src/net/sock_cloexec.go:10
)

import (
	"internal/poll"
	"os"
	"syscall"
)

// Wrapper around the socket system call that marks the returned file
//line /snap/go/10455/src/net/sock_cloexec.go:18
// descriptor as nonblocking and close-on-exec.
//line /snap/go/10455/src/net/sock_cloexec.go:20
func sysSocket(family, sotype, proto int) (int, error) {
//line /snap/go/10455/src/net/sock_cloexec.go:20
	_go_fuzz_dep_.CoverTab[8005]++
							s, err := socketFunc(family, sotype|syscall.SOCK_NONBLOCK|syscall.SOCK_CLOEXEC, proto)

//line /snap/go/10455/src/net/sock_cloexec.go:25
	switch err {
	case nil:
//line /snap/go/10455/src/net/sock_cloexec.go:26
		_go_fuzz_dep_.CoverTab[529608]++
//line /snap/go/10455/src/net/sock_cloexec.go:26
		_go_fuzz_dep_.CoverTab[8010]++
								return s, nil
//line /snap/go/10455/src/net/sock_cloexec.go:27
		// _ = "end of CoverTab[8010]"
	default:
//line /snap/go/10455/src/net/sock_cloexec.go:28
		_go_fuzz_dep_.CoverTab[529609]++
//line /snap/go/10455/src/net/sock_cloexec.go:28
		_go_fuzz_dep_.CoverTab[8011]++
								return -1, os.NewSyscallError("socket", err)
//line /snap/go/10455/src/net/sock_cloexec.go:29
		// _ = "end of CoverTab[8011]"
	case syscall.EPROTONOSUPPORT, syscall.EINVAL:
//line /snap/go/10455/src/net/sock_cloexec.go:30
		_go_fuzz_dep_.CoverTab[529610]++
//line /snap/go/10455/src/net/sock_cloexec.go:30
		_go_fuzz_dep_.CoverTab[8012]++
//line /snap/go/10455/src/net/sock_cloexec.go:30
		// _ = "end of CoverTab[8012]"
	}
//line /snap/go/10455/src/net/sock_cloexec.go:31
	// _ = "end of CoverTab[8005]"
//line /snap/go/10455/src/net/sock_cloexec.go:31
	_go_fuzz_dep_.CoverTab[8006]++

//line /snap/go/10455/src/net/sock_cloexec.go:34
	syscall.ForkLock.RLock()
	s, err = socketFunc(family, sotype, proto)
	if err == nil {
//line /snap/go/10455/src/net/sock_cloexec.go:36
		_go_fuzz_dep_.CoverTab[529611]++
//line /snap/go/10455/src/net/sock_cloexec.go:36
		_go_fuzz_dep_.CoverTab[8013]++
								syscall.CloseOnExec(s)
//line /snap/go/10455/src/net/sock_cloexec.go:37
		// _ = "end of CoverTab[8013]"
	} else {
//line /snap/go/10455/src/net/sock_cloexec.go:38
		_go_fuzz_dep_.CoverTab[529612]++
//line /snap/go/10455/src/net/sock_cloexec.go:38
		_go_fuzz_dep_.CoverTab[8014]++
//line /snap/go/10455/src/net/sock_cloexec.go:38
		// _ = "end of CoverTab[8014]"
//line /snap/go/10455/src/net/sock_cloexec.go:38
	}
//line /snap/go/10455/src/net/sock_cloexec.go:38
	// _ = "end of CoverTab[8006]"
//line /snap/go/10455/src/net/sock_cloexec.go:38
	_go_fuzz_dep_.CoverTab[8007]++
							syscall.ForkLock.RUnlock()
							if err != nil {
//line /snap/go/10455/src/net/sock_cloexec.go:40
		_go_fuzz_dep_.CoverTab[529613]++
//line /snap/go/10455/src/net/sock_cloexec.go:40
		_go_fuzz_dep_.CoverTab[8015]++
								return -1, os.NewSyscallError("socket", err)
//line /snap/go/10455/src/net/sock_cloexec.go:41
		// _ = "end of CoverTab[8015]"
	} else {
//line /snap/go/10455/src/net/sock_cloexec.go:42
		_go_fuzz_dep_.CoverTab[529614]++
//line /snap/go/10455/src/net/sock_cloexec.go:42
		_go_fuzz_dep_.CoverTab[8016]++
//line /snap/go/10455/src/net/sock_cloexec.go:42
		// _ = "end of CoverTab[8016]"
//line /snap/go/10455/src/net/sock_cloexec.go:42
	}
//line /snap/go/10455/src/net/sock_cloexec.go:42
	// _ = "end of CoverTab[8007]"
//line /snap/go/10455/src/net/sock_cloexec.go:42
	_go_fuzz_dep_.CoverTab[8008]++
							if err = syscall.SetNonblock(s, true); err != nil {
//line /snap/go/10455/src/net/sock_cloexec.go:43
		_go_fuzz_dep_.CoverTab[529615]++
//line /snap/go/10455/src/net/sock_cloexec.go:43
		_go_fuzz_dep_.CoverTab[8017]++
								poll.CloseFunc(s)
								return -1, os.NewSyscallError("setnonblock", err)
//line /snap/go/10455/src/net/sock_cloexec.go:45
		// _ = "end of CoverTab[8017]"
	} else {
//line /snap/go/10455/src/net/sock_cloexec.go:46
		_go_fuzz_dep_.CoverTab[529616]++
//line /snap/go/10455/src/net/sock_cloexec.go:46
		_go_fuzz_dep_.CoverTab[8018]++
//line /snap/go/10455/src/net/sock_cloexec.go:46
		// _ = "end of CoverTab[8018]"
//line /snap/go/10455/src/net/sock_cloexec.go:46
	}
//line /snap/go/10455/src/net/sock_cloexec.go:46
	// _ = "end of CoverTab[8008]"
//line /snap/go/10455/src/net/sock_cloexec.go:46
	_go_fuzz_dep_.CoverTab[8009]++
							return s, nil
//line /snap/go/10455/src/net/sock_cloexec.go:47
	// _ = "end of CoverTab[8009]"
}

//line /snap/go/10455/src/net/sock_cloexec.go:48
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /snap/go/10455/src/net/sock_cloexec.go:48
var _ = _go_fuzz_dep_.CoverTab
