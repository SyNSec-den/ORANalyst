// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix || (js && wasm) || wasip1

//line /snap/go/10455/src/net/hook_unix.go:7
package net

//line /snap/go/10455/src/net/hook_unix.go:7
import (
//line /snap/go/10455/src/net/hook_unix.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /snap/go/10455/src/net/hook_unix.go:7
)
//line /snap/go/10455/src/net/hook_unix.go:7
import (
//line /snap/go/10455/src/net/hook_unix.go:7
	_atomic_ "sync/atomic"
//line /snap/go/10455/src/net/hook_unix.go:7
)

import "syscall"

var (
	testHookDialChannel	= func() { _go_fuzz_dep_.CoverTab[5983]++; // _ = "end of CoverTab[5983]" }	// for golang.org/issue/5349
	testHookCanceledDial	= func() { _go_fuzz_dep_.CoverTab[5984]++; // _ = "end of CoverTab[5984]" }	// for golang.org/issue/16523

	// Placeholders for socket system calls.
	socketFunc		func(int, int, int) (int, error)	= syscall.Socket
	connectFunc		func(int, syscall.Sockaddr) error	= syscall.Connect
	listenFunc		func(int, int) error			= syscall.Listen
	getsockoptIntFunc	func(int, int, int) (int, error)	= syscall.GetsockoptInt
)
//line /snap/go/10455/src/net/hook_unix.go:20
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /snap/go/10455/src/net/hook_unix.go:20
var _ = _go_fuzz_dep_.CoverTab
