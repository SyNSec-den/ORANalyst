// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build (darwin || dragonfly || freebsd || (linux && !ppc64 && !ppc64le) || netbsd || openbsd || solaris) && gc
// +build darwin dragonfly freebsd linux,!ppc64,!ppc64le netbsd openbsd solaris
// +build gc

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix_gc.go:9
package unix

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix_gc.go:9
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix_gc.go:9
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix_gc.go:9
)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix_gc.go:9
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix_gc.go:9
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix_gc.go:9
)

import "syscall"

func Syscall(trap, a1, a2, a3 uintptr) (r1, r2 uintptr, err syscall.Errno)
func Syscall6(trap, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err syscall.Errno)
func RawSyscall(trap, a1, a2, a3 uintptr) (r1, r2 uintptr, err syscall.Errno)
func RawSyscall6(trap, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err syscall.Errno)

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix_gc.go:16
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix_gc.go:16
var _ = _go_fuzz_dep_.CoverTab
