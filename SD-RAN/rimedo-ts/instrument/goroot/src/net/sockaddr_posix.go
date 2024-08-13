// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix || (js && wasm) || windows

//line /usr/local/go/src/net/sockaddr_posix.go:7
package net

//line /usr/local/go/src/net/sockaddr_posix.go:7
import (
//line /usr/local/go/src/net/sockaddr_posix.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/sockaddr_posix.go:7
)
//line /usr/local/go/src/net/sockaddr_posix.go:7
import (
//line /usr/local/go/src/net/sockaddr_posix.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/sockaddr_posix.go:7
)

import (
	"syscall"
)

// A sockaddr represents a TCP, UDP, IP or Unix network endpoint
//line /usr/local/go/src/net/sockaddr_posix.go:13
// address that can be converted into a syscall.Sockaddr.
//line /usr/local/go/src/net/sockaddr_posix.go:15
type sockaddr interface {
	Addr

	// family returns the platform-dependent address family
	// identifier.
	family() int

	// isWildcard reports whether the address is a wildcard
	// address.
	isWildcard() bool

	// sockaddr returns the address converted into a syscall
	// sockaddr type that implements syscall.Sockaddr
	// interface. It returns a nil interface when the address is
	// nil.
	sockaddr(family int) (syscall.Sockaddr, error)

	// toLocal maps the zero address to a local system address (127.0.0.1 or ::1)
	toLocal(net string) sockaddr
}

//line /usr/local/go/src/net/sockaddr_posix.go:34
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/sockaddr_posix.go:34
var _ = _go_fuzz_dep_.CoverTab
