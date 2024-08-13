// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris || zos
// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris zos

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/fdset.go:8
package unix

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/fdset.go:8
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/fdset.go:8
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/fdset.go:8
)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/fdset.go:8
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/fdset.go:8
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/fdset.go:8
)

// Set adds fd to the set fds.
func (fds *FdSet) Set(fd int) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/fdset.go:11
	_go_fuzz_dep_.CoverTab[45781]++
										fds.Bits[fd/NFDBITS] |= (1 << (uintptr(fd) % NFDBITS))
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/fdset.go:12
	// _ = "end of CoverTab[45781]"
}

// Clear removes fd from the set fds.
func (fds *FdSet) Clear(fd int) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/fdset.go:16
	_go_fuzz_dep_.CoverTab[45782]++
										fds.Bits[fd/NFDBITS] &^= (1 << (uintptr(fd) % NFDBITS))
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/fdset.go:17
	// _ = "end of CoverTab[45782]"
}

// IsSet returns whether fd is in the set fds.
func (fds *FdSet) IsSet(fd int) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/fdset.go:21
	_go_fuzz_dep_.CoverTab[45783]++
										return fds.Bits[fd/NFDBITS]&(1<<(uintptr(fd)%NFDBITS)) != 0
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/fdset.go:22
	// _ = "end of CoverTab[45783]"
}

// Zero clears the set fds.
func (fds *FdSet) Zero() {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/fdset.go:26
	_go_fuzz_dep_.CoverTab[45784]++
										for i := range fds.Bits {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/fdset.go:27
		_go_fuzz_dep_.CoverTab[45785]++
											fds.Bits[i] = 0
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/fdset.go:28
		// _ = "end of CoverTab[45785]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/fdset.go:29
	// _ = "end of CoverTab[45784]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/fdset.go:30
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/fdset.go:30
var _ = _go_fuzz_dep_.CoverTab
