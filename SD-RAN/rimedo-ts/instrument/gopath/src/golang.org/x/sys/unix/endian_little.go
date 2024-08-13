// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
//go:build 386 || amd64 || amd64p32 || alpha || arm || arm64 || loong64 || mipsle || mips64le || mips64p32le || nios2 || ppc64le || riscv || riscv64 || sh
// +build 386 amd64 amd64p32 alpha arm arm64 loong64 mipsle mips64le mips64p32le nios2 ppc64le riscv riscv64 sh

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/endian_little.go:8
package unix

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/endian_little.go:8
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/endian_little.go:8
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/endian_little.go:8
)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/endian_little.go:8
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/endian_little.go:8
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/endian_little.go:8
)

const isBigEndian = false

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/endian_little.go:10
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/endian_little.go:10
var _ = _go_fuzz_dep_.CoverTab
