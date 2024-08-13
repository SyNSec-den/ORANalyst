// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux
// +build linux

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sysvshm_linux.go:8
package unix

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sysvshm_linux.go:8
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sysvshm_linux.go:8
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sysvshm_linux.go:8
)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sysvshm_linux.go:8
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sysvshm_linux.go:8
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sysvshm_linux.go:8
)

import "runtime"

// SysvShmCtl performs control operations on the shared memory segment
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sysvshm_linux.go:12
// specified by id.
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sysvshm_linux.go:14
func SysvShmCtl(id, cmd int, desc *SysvShmDesc) (result int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sysvshm_linux.go:14
	_go_fuzz_dep_.CoverTab[46838]++
											if runtime.GOARCH == "arm" || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sysvshm_linux.go:15
		_go_fuzz_dep_.CoverTab[46840]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sysvshm_linux.go:15
		return runtime.GOARCH == "mips64"
												// _ = "end of CoverTab[46840]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sysvshm_linux.go:16
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sysvshm_linux.go:16
		_go_fuzz_dep_.CoverTab[46841]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sysvshm_linux.go:16
		return runtime.GOARCH == "mips64le"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sysvshm_linux.go:16
		// _ = "end of CoverTab[46841]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sysvshm_linux.go:16
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sysvshm_linux.go:16
		_go_fuzz_dep_.CoverTab[46842]++
												cmd |= ipc_64
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sysvshm_linux.go:17
		// _ = "end of CoverTab[46842]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sysvshm_linux.go:18
		_go_fuzz_dep_.CoverTab[46843]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sysvshm_linux.go:18
		// _ = "end of CoverTab[46843]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sysvshm_linux.go:18
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sysvshm_linux.go:18
	// _ = "end of CoverTab[46838]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sysvshm_linux.go:18
	_go_fuzz_dep_.CoverTab[46839]++

											return shmctl(id, cmd, desc)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sysvshm_linux.go:20
	// _ = "end of CoverTab[46839]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sysvshm_linux.go:21
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sysvshm_linux.go:21
var _ = _go_fuzz_dep_.CoverTab
