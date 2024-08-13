// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris || zos
// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris zos

// Unix environment variables.

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/env_unix.go:10
package unix

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/env_unix.go:10
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/env_unix.go:10
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/env_unix.go:10
)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/env_unix.go:10
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/env_unix.go:10
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/env_unix.go:10
)

import "syscall"

func Getenv(key string) (value string, found bool) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/env_unix.go:14
	_go_fuzz_dep_.CoverTab[45767]++
										return syscall.Getenv(key)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/env_unix.go:15
	// _ = "end of CoverTab[45767]"
}

func Setenv(key, value string) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/env_unix.go:18
	_go_fuzz_dep_.CoverTab[45768]++
										return syscall.Setenv(key, value)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/env_unix.go:19
	// _ = "end of CoverTab[45768]"
}

func Clearenv() {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/env_unix.go:22
	_go_fuzz_dep_.CoverTab[45769]++
										syscall.Clearenv()
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/env_unix.go:23
	// _ = "end of CoverTab[45769]"
}

func Environ() []string {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/env_unix.go:26
	_go_fuzz_dep_.CoverTab[45770]++
										return syscall.Environ()
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/env_unix.go:27
	// _ = "end of CoverTab[45770]"
}

func Unsetenv(key string) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/env_unix.go:30
	_go_fuzz_dep_.CoverTab[45771]++
										return syscall.Unsetenv(key)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/env_unix.go:31
	// _ = "end of CoverTab[45771]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/env_unix.go:32
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/env_unix.go:32
var _ = _go_fuzz_dep_.CoverTab
