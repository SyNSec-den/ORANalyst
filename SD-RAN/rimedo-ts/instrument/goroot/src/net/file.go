// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/net/file.go:5
package net

//line /usr/local/go/src/net/file.go:5
import (
//line /usr/local/go/src/net/file.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/file.go:5
)
//line /usr/local/go/src/net/file.go:5
import (
//line /usr/local/go/src/net/file.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/file.go:5
)

import "os"

//line /usr/local/go/src/net/file.go:12
type fileAddr string

func (fileAddr) Network() string {
//line /usr/local/go/src/net/file.go:14
	_go_fuzz_dep_.CoverTab[13929]++
//line /usr/local/go/src/net/file.go:14
	return "file+net"
//line /usr/local/go/src/net/file.go:14
	// _ = "end of CoverTab[13929]"
//line /usr/local/go/src/net/file.go:14
}
func (f fileAddr) String() string {
//line /usr/local/go/src/net/file.go:15
	_go_fuzz_dep_.CoverTab[13930]++
//line /usr/local/go/src/net/file.go:15
	return string(f)
//line /usr/local/go/src/net/file.go:15
	// _ = "end of CoverTab[13930]"
//line /usr/local/go/src/net/file.go:15
}

// FileConn returns a copy of the network connection corresponding to
//line /usr/local/go/src/net/file.go:17
// the open file f.
//line /usr/local/go/src/net/file.go:17
// It is the caller's responsibility to close f when finished.
//line /usr/local/go/src/net/file.go:17
// Closing c does not affect f, and closing f does not affect c.
//line /usr/local/go/src/net/file.go:21
func FileConn(f *os.File) (c Conn, err error) {
//line /usr/local/go/src/net/file.go:21
	_go_fuzz_dep_.CoverTab[13931]++
						c, err = fileConn(f)
						if err != nil {
//line /usr/local/go/src/net/file.go:23
		_go_fuzz_dep_.CoverTab[13933]++
							err = &OpError{Op: "file", Net: "file+net", Source: nil, Addr: fileAddr(f.Name()), Err: err}
//line /usr/local/go/src/net/file.go:24
		// _ = "end of CoverTab[13933]"
	} else {
//line /usr/local/go/src/net/file.go:25
		_go_fuzz_dep_.CoverTab[13934]++
//line /usr/local/go/src/net/file.go:25
		// _ = "end of CoverTab[13934]"
//line /usr/local/go/src/net/file.go:25
	}
//line /usr/local/go/src/net/file.go:25
	// _ = "end of CoverTab[13931]"
//line /usr/local/go/src/net/file.go:25
	_go_fuzz_dep_.CoverTab[13932]++
						return
//line /usr/local/go/src/net/file.go:26
	// _ = "end of CoverTab[13932]"
}

// FileListener returns a copy of the network listener corresponding
//line /usr/local/go/src/net/file.go:29
// to the open file f.
//line /usr/local/go/src/net/file.go:29
// It is the caller's responsibility to close ln when finished.
//line /usr/local/go/src/net/file.go:29
// Closing ln does not affect f, and closing f does not affect ln.
//line /usr/local/go/src/net/file.go:33
func FileListener(f *os.File) (ln Listener, err error) {
//line /usr/local/go/src/net/file.go:33
	_go_fuzz_dep_.CoverTab[13935]++
						ln, err = fileListener(f)
						if err != nil {
//line /usr/local/go/src/net/file.go:35
		_go_fuzz_dep_.CoverTab[13937]++
							err = &OpError{Op: "file", Net: "file+net", Source: nil, Addr: fileAddr(f.Name()), Err: err}
//line /usr/local/go/src/net/file.go:36
		// _ = "end of CoverTab[13937]"
	} else {
//line /usr/local/go/src/net/file.go:37
		_go_fuzz_dep_.CoverTab[13938]++
//line /usr/local/go/src/net/file.go:37
		// _ = "end of CoverTab[13938]"
//line /usr/local/go/src/net/file.go:37
	}
//line /usr/local/go/src/net/file.go:37
	// _ = "end of CoverTab[13935]"
//line /usr/local/go/src/net/file.go:37
	_go_fuzz_dep_.CoverTab[13936]++
						return
//line /usr/local/go/src/net/file.go:38
	// _ = "end of CoverTab[13936]"
}

// FilePacketConn returns a copy of the packet network connection
//line /usr/local/go/src/net/file.go:41
// corresponding to the open file f.
//line /usr/local/go/src/net/file.go:41
// It is the caller's responsibility to close f when finished.
//line /usr/local/go/src/net/file.go:41
// Closing c does not affect f, and closing f does not affect c.
//line /usr/local/go/src/net/file.go:45
func FilePacketConn(f *os.File) (c PacketConn, err error) {
//line /usr/local/go/src/net/file.go:45
	_go_fuzz_dep_.CoverTab[13939]++
						c, err = filePacketConn(f)
						if err != nil {
//line /usr/local/go/src/net/file.go:47
		_go_fuzz_dep_.CoverTab[13941]++
							err = &OpError{Op: "file", Net: "file+net", Source: nil, Addr: fileAddr(f.Name()), Err: err}
//line /usr/local/go/src/net/file.go:48
		// _ = "end of CoverTab[13941]"
	} else {
//line /usr/local/go/src/net/file.go:49
		_go_fuzz_dep_.CoverTab[13942]++
//line /usr/local/go/src/net/file.go:49
		// _ = "end of CoverTab[13942]"
//line /usr/local/go/src/net/file.go:49
	}
//line /usr/local/go/src/net/file.go:49
	// _ = "end of CoverTab[13939]"
//line /usr/local/go/src/net/file.go:49
	_go_fuzz_dep_.CoverTab[13940]++
						return
//line /usr/local/go/src/net/file.go:50
	// _ = "end of CoverTab[13940]"
}

//line /usr/local/go/src/net/file.go:51
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/file.go:51
var _ = _go_fuzz_dep_.CoverTab