// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris
// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris

// For Unix, get the pagesize from the runtime.

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/pagesize_unix.go:10
package unix

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/pagesize_unix.go:10
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/pagesize_unix.go:10
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/pagesize_unix.go:10
)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/pagesize_unix.go:10
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/pagesize_unix.go:10
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/pagesize_unix.go:10
)

import "syscall"

func Getpagesize() int {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/pagesize_unix.go:14
	_go_fuzz_dep_.CoverTab[45853]++
											return syscall.Getpagesize()
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/pagesize_unix.go:15
	// _ = "end of CoverTab[45853]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/pagesize_unix.go:16
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/pagesize_unix.go:16
var _ = _go_fuzz_dep_.CoverTab
