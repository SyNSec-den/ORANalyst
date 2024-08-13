// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build aix || darwin || freebsd || linux || netbsd || openbsd || solaris || zos
// +build aix darwin freebsd linux netbsd openbsd solaris zos

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix_other.go:8
package unix

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix_other.go:8
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix_other.go:8
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix_other.go:8
)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix_other.go:8
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix_other.go:8
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix_other.go:8
)

import (
	"runtime"
)

// Round the length of a raw sockaddr up to align it properly.
func cmsgAlignOf(salen int) int {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix_other.go:15
	_go_fuzz_dep_.CoverTab[45908]++
												salign := SizeofPtr

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix_other.go:20
	switch runtime.GOOS {
	case "aix":
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix_other.go:21
		_go_fuzz_dep_.CoverTab[45910]++

													salign = 1
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix_other.go:23
		// _ = "end of CoverTab[45910]"
	case "darwin", "ios", "illumos", "solaris":
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix_other.go:24
		_go_fuzz_dep_.CoverTab[45911]++

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix_other.go:28
		if SizeofPtr == 8 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix_other.go:28
			_go_fuzz_dep_.CoverTab[45916]++
														salign = 4
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix_other.go:29
			// _ = "end of CoverTab[45916]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix_other.go:30
			_go_fuzz_dep_.CoverTab[45917]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix_other.go:30
			// _ = "end of CoverTab[45917]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix_other.go:30
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix_other.go:30
		// _ = "end of CoverTab[45911]"
	case "netbsd", "openbsd":
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix_other.go:31
		_go_fuzz_dep_.CoverTab[45912]++

													if runtime.GOARCH == "arm" {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix_other.go:33
			_go_fuzz_dep_.CoverTab[45918]++
														salign = 8
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix_other.go:34
			// _ = "end of CoverTab[45918]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix_other.go:35
			_go_fuzz_dep_.CoverTab[45919]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix_other.go:35
			// _ = "end of CoverTab[45919]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix_other.go:35
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix_other.go:35
		// _ = "end of CoverTab[45912]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix_other.go:35
		_go_fuzz_dep_.CoverTab[45913]++

													if runtime.GOOS == "netbsd" && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix_other.go:37
			_go_fuzz_dep_.CoverTab[45920]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix_other.go:37
			return runtime.GOARCH == "arm64"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix_other.go:37
			// _ = "end of CoverTab[45920]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix_other.go:37
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix_other.go:37
			_go_fuzz_dep_.CoverTab[45921]++
														salign = 16
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix_other.go:38
			// _ = "end of CoverTab[45921]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix_other.go:39
			_go_fuzz_dep_.CoverTab[45922]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix_other.go:39
			// _ = "end of CoverTab[45922]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix_other.go:39
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix_other.go:39
		// _ = "end of CoverTab[45913]"
	case "zos":
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix_other.go:40
		_go_fuzz_dep_.CoverTab[45914]++

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix_other.go:43
		salign = SizeofInt
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix_other.go:43
		// _ = "end of CoverTab[45914]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix_other.go:43
	default:
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix_other.go:43
		_go_fuzz_dep_.CoverTab[45915]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix_other.go:43
		// _ = "end of CoverTab[45915]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix_other.go:44
	// _ = "end of CoverTab[45908]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix_other.go:44
	_go_fuzz_dep_.CoverTab[45909]++

												return (salen + salign - 1) & ^(salign - 1)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix_other.go:46
	// _ = "end of CoverTab[45909]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix_other.go:47
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix_other.go:47
var _ = _go_fuzz_dep_.CoverTab
