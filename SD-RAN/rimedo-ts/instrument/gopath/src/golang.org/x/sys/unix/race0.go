// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build aix || (darwin && !race) || (linux && !race) || (freebsd && !race) || netbsd || openbsd || solaris || dragonfly || zos
// +build aix darwin,!race linux,!race freebsd,!race netbsd openbsd solaris dragonfly zos

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/race0.go:8
package unix

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/race0.go:8
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/race0.go:8
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/race0.go:8
)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/race0.go:8
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/race0.go:8
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/race0.go:8
)

import (
	"unsafe"
)

const raceenabled = false

func raceAcquire(addr unsafe.Pointer) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/race0.go:16
	_go_fuzz_dep_.CoverTab[45854]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/race0.go:16
	// _ = "end of CoverTab[45854]"
}

func raceReleaseMerge(addr unsafe.Pointer) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/race0.go:19
	_go_fuzz_dep_.CoverTab[45855]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/race0.go:19
	// _ = "end of CoverTab[45855]"
}

func raceReadRange(addr unsafe.Pointer, len int) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/race0.go:22
	_go_fuzz_dep_.CoverTab[45856]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/race0.go:22
	// _ = "end of CoverTab[45856]"
}

func raceWriteRange(addr unsafe.Pointer, len int) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/race0.go:25
	_go_fuzz_dep_.CoverTab[45857]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/race0.go:25
	// _ = "end of CoverTab[45857]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/race0.go:26
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/race0.go:26
var _ = _go_fuzz_dep_.CoverTab
