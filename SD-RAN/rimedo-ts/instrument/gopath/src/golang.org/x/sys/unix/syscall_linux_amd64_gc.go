// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build amd64 && linux && gc
// +build amd64,linux,gc

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64_gc.go:8
package unix

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64_gc.go:8
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64_gc.go:8
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64_gc.go:8
)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64_gc.go:8
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64_gc.go:8
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64_gc.go:8
)

import "syscall"

//go:noescape
func gettimeofday(tv *Timeval) (err syscall.Errno)

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64_gc.go:13
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64_gc.go:13
var _ = _go_fuzz_dep_.CoverTab
