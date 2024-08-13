// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build (darwin && !ios) || linux
// +build darwin,!ios linux

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sysvshm_unix.go:8
package unix

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sysvshm_unix.go:8
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sysvshm_unix.go:8
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sysvshm_unix.go:8
)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sysvshm_unix.go:8
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sysvshm_unix.go:8
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sysvshm_unix.go:8
)

import "unsafe"

// SysvShmAttach attaches the Sysv shared memory segment associated with the
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sysvshm_unix.go:12
// shared memory identifier id.
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sysvshm_unix.go:14
func SysvShmAttach(id int, addr uintptr, flag int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sysvshm_unix.go:14
	_go_fuzz_dep_.CoverTab[46844]++
											addr, errno := shmat(id, addr, flag)
											if errno != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sysvshm_unix.go:16
		_go_fuzz_dep_.CoverTab[46847]++
												return nil, errno
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sysvshm_unix.go:17
		// _ = "end of CoverTab[46847]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sysvshm_unix.go:18
		_go_fuzz_dep_.CoverTab[46848]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sysvshm_unix.go:18
		// _ = "end of CoverTab[46848]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sysvshm_unix.go:18
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sysvshm_unix.go:18
	// _ = "end of CoverTab[46844]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sysvshm_unix.go:18
	_go_fuzz_dep_.CoverTab[46845]++

	// Retrieve the size of the shared memory to enable slice creation
	var info SysvShmDesc

	_, err := SysvShmCtl(id, IPC_STAT, &info)
	if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sysvshm_unix.go:24
		_go_fuzz_dep_.CoverTab[46849]++

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sysvshm_unix.go:28
		shmdt(addr)
												return nil, err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sysvshm_unix.go:29
		// _ = "end of CoverTab[46849]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sysvshm_unix.go:30
		_go_fuzz_dep_.CoverTab[46850]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sysvshm_unix.go:30
		// _ = "end of CoverTab[46850]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sysvshm_unix.go:30
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sysvshm_unix.go:30
	// _ = "end of CoverTab[46845]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sysvshm_unix.go:30
	_go_fuzz_dep_.CoverTab[46846]++

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sysvshm_unix.go:33
	b := unsafe.Slice((*byte)(unsafe.Pointer(addr)), int(info.Segsz))
											return b, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sysvshm_unix.go:34
	// _ = "end of CoverTab[46846]"
}

// SysvShmDetach unmaps the shared memory slice returned from SysvShmAttach.
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sysvshm_unix.go:37
//
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sysvshm_unix.go:37
// It is not safe to use the slice after calling this function.
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sysvshm_unix.go:40
func SysvShmDetach(data []byte) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sysvshm_unix.go:40
	_go_fuzz_dep_.CoverTab[46851]++
											if len(data) == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sysvshm_unix.go:41
		_go_fuzz_dep_.CoverTab[46853]++
												return EINVAL
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sysvshm_unix.go:42
		// _ = "end of CoverTab[46853]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sysvshm_unix.go:43
		_go_fuzz_dep_.CoverTab[46854]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sysvshm_unix.go:43
		// _ = "end of CoverTab[46854]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sysvshm_unix.go:43
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sysvshm_unix.go:43
	// _ = "end of CoverTab[46851]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sysvshm_unix.go:43
	_go_fuzz_dep_.CoverTab[46852]++

											return shmdt(uintptr(unsafe.Pointer(&data[0])))
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sysvshm_unix.go:45
	// _ = "end of CoverTab[46852]"
}

// SysvShmGet returns the Sysv shared memory identifier associated with key.
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sysvshm_unix.go:48
// If the IPC_CREAT flag is specified a new segment is created.
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sysvshm_unix.go:50
func SysvShmGet(key, size, flag int) (id int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sysvshm_unix.go:50
	_go_fuzz_dep_.CoverTab[46855]++
											return shmget(key, size, flag)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sysvshm_unix.go:51
	// _ = "end of CoverTab[46855]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sysvshm_unix.go:52
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sysvshm_unix.go:52
var _ = _go_fuzz_dep_.CoverTab
