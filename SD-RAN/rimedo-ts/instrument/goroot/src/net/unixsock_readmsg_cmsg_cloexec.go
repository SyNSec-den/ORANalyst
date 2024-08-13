// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build dragonfly || linux || netbsd || openbsd

//line /usr/local/go/src/net/unixsock_readmsg_cmsg_cloexec.go:7
package net

//line /usr/local/go/src/net/unixsock_readmsg_cmsg_cloexec.go:7
import (
//line /usr/local/go/src/net/unixsock_readmsg_cmsg_cloexec.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/unixsock_readmsg_cmsg_cloexec.go:7
)
//line /usr/local/go/src/net/unixsock_readmsg_cmsg_cloexec.go:7
import (
//line /usr/local/go/src/net/unixsock_readmsg_cmsg_cloexec.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/unixsock_readmsg_cmsg_cloexec.go:7
)

import "syscall"

const readMsgFlags = syscall.MSG_CMSG_CLOEXEC

func setReadMsgCloseOnExec(oob []byte) {
//line /usr/local/go/src/net/unixsock_readmsg_cmsg_cloexec.go:13
	_go_fuzz_dep_.CoverTab[17207]++
//line /usr/local/go/src/net/unixsock_readmsg_cmsg_cloexec.go:13
	// _ = "end of CoverTab[17207]"
//line /usr/local/go/src/net/unixsock_readmsg_cmsg_cloexec.go:13
}

//line /usr/local/go/src/net/unixsock_readmsg_cmsg_cloexec.go:13
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/unixsock_readmsg_cmsg_cloexec.go:13
var _ = _go_fuzz_dep_.CoverTab
