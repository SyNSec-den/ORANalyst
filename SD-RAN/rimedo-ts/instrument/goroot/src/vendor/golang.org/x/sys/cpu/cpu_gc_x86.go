// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build (386 || amd64 || amd64p32) && gc
// +build 386 amd64 amd64p32
// +build gc

//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_gc_x86.go:9
package cpu

//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_gc_x86.go:9
import (
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_gc_x86.go:9
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_gc_x86.go:9
)
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_gc_x86.go:9
import (
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_gc_x86.go:9
	_atomic_ "sync/atomic"
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_gc_x86.go:9
)

// cpuid is implemented in cpu_x86.s for gc compiler
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_gc_x86.go:11
// and in cpu_gccgo.c for gccgo.
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_gc_x86.go:13
func cpuid(eaxArg, ecxArg uint32) (eax, ebx, ecx, edx uint32)

// xgetbv with ecx = 0 is implemented in cpu_x86.s for gc compiler
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_gc_x86.go:15
// and in cpu_gccgo.c for gccgo.
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_gc_x86.go:17
func xgetbv() (eax, edx uint32)

//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_gc_x86.go:17
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu_gc_x86.go:17
var _ = _go_fuzz_dep_.CoverTab
