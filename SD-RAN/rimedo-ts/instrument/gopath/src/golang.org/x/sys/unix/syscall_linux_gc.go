// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux && gc
// +build linux,gc

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_gc.go:8
package unix

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_gc.go:8
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_gc.go:8
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_gc.go:8
)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_gc.go:8
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_gc.go:8
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_gc.go:8
)

// SyscallNoError may be used instead of Syscall for syscalls that don't fail.
func SyscallNoError(trap, a1, a2, a3 uintptr) (r1, r2 uintptr)

// RawSyscallNoError may be used instead of RawSyscall for syscalls that don't
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_gc.go:13
// fail.
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_gc.go:15
func RawSyscallNoError(trap, a1, a2, a3 uintptr) (r1, r2 uintptr)

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_gc.go:15
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_gc.go:15
var _ = _go_fuzz_dep_.CoverTab
