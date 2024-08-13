// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build aix || dragonfly || freebsd || linux || netbsd || openbsd
// +build aix dragonfly freebsd linux netbsd openbsd

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/readdirent_getdents.go:8
package unix

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/readdirent_getdents.go:8
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/readdirent_getdents.go:8
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/readdirent_getdents.go:8
)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/readdirent_getdents.go:8
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/readdirent_getdents.go:8
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/readdirent_getdents.go:8
)

// ReadDirent reads directory entries from fd and writes them into buf.
func ReadDirent(fd int, buf []byte) (n int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/readdirent_getdents.go:11
	_go_fuzz_dep_.CoverTab[45858]++
												return Getdents(fd, buf)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/readdirent_getdents.go:12
	// _ = "end of CoverTab[45858]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/readdirent_getdents.go:13
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/readdirent_getdents.go:13
var _ = _go_fuzz_dep_.CoverTab
