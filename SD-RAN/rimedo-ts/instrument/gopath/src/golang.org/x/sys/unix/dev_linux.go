// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Functions to access/create device major and minor numbers matching the
// encoding used by the Linux kernel and glibc.
//
// The information below is extracted and adapted from bits/sysmacros.h in the
// glibc sources:
//
// dev_t in glibc is 64-bit, with 32-bit major and minor numbers. glibc's
// default encoding is MMMM Mmmm mmmM MMmm, where M is a hex digit of the major
// number and m is a hex digit of the minor number. This is backward compatible
// with legacy systems where dev_t is 16 bits wide, encoded as MMmm. It is also
// backward compatible with the Linux kernel, which for some architectures uses
// 32-bit dev_t, encoded as mmmM MMmm.

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dev_linux.go:18
package unix

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dev_linux.go:18
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dev_linux.go:18
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dev_linux.go:18
)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dev_linux.go:18
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dev_linux.go:18
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dev_linux.go:18
)

// Major returns the major component of a Linux device number.
func Major(dev uint64) uint32 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dev_linux.go:21
	_go_fuzz_dep_.CoverTab[45719]++
										major := uint32((dev & 0x00000000000fff00) >> 8)
										major |= uint32((dev & 0xfffff00000000000) >> 32)
										return major
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dev_linux.go:24
	// _ = "end of CoverTab[45719]"
}

// Minor returns the minor component of a Linux device number.
func Minor(dev uint64) uint32 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dev_linux.go:28
	_go_fuzz_dep_.CoverTab[45720]++
										minor := uint32((dev & 0x00000000000000ff) >> 0)
										minor |= uint32((dev & 0x00000ffffff00000) >> 12)
										return minor
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dev_linux.go:31
	// _ = "end of CoverTab[45720]"
}

// Mkdev returns a Linux device number generated from the given major and minor
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dev_linux.go:34
// components.
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dev_linux.go:36
func Mkdev(major, minor uint32) uint64 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dev_linux.go:36
	_go_fuzz_dep_.CoverTab[45721]++
										dev := (uint64(major) & 0x00000fff) << 8
										dev |= (uint64(major) & 0xfffff000) << 32
										dev |= (uint64(minor) & 0x000000ff) << 0
										dev |= (uint64(minor) & 0xffffff00) << 12
										return dev
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dev_linux.go:41
	// _ = "end of CoverTab[45721]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dev_linux.go:42
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dev_linux.go:42
var _ = _go_fuzz_dep_.CoverTab
