// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux && !arm && !arm64 && !mips64 && !mips64le && !ppc64 && !ppc64le && !s390x
// +build linux,!arm,!arm64,!mips64,!mips64le,!ppc64,!ppc64le,!s390x

//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_linux_noinit.go:8
package cpu

//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_linux_noinit.go:8
import (
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_linux_noinit.go:8
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_linux_noinit.go:8
)
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_linux_noinit.go:8
import (
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_linux_noinit.go:8
	_atomic_ "sync/atomic"
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_linux_noinit.go:8
)

func doinit()	{ _go_fuzz_dep_.CoverTab[20869]++; // _ = "end of CoverTab[20869]" }

//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_linux_noinit.go:10
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_linux_noinit.go:10
var _ = _go_fuzz_dep_.CoverTab
