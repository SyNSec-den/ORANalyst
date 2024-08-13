// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build (aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris || zos) && go1.9
// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris zos
// +build go1.9

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/aliases.go:9
package unix

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/aliases.go:9
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/aliases.go:9
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/aliases.go:9
)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/aliases.go:9
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/aliases.go:9
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/aliases.go:9
)

import "syscall"

type Signal = syscall.Signal
type Errno = syscall.Errno
type SysProcAttr = syscall.SysProcAttr

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/aliases.go:15
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/aliases.go:15
var _ = _go_fuzz_dep_.CoverTab
