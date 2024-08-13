// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /snap/go/10455/src/net/file.go:5
package net

//line /snap/go/10455/src/net/file.go:5
import (
//line /snap/go/10455/src/net/file.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /snap/go/10455/src/net/file.go:5
)
//line /snap/go/10455/src/net/file.go:5
import (
//line /snap/go/10455/src/net/file.go:5
	_atomic_ "sync/atomic"
//line /snap/go/10455/src/net/file.go:5
)

import "os"

//line /snap/go/10455/src/net/file.go:12
type fileAddr string

func (fileAddr) Network() string {
//line /snap/go/10455/src/net/file.go:14
	_go_fuzz_dep_.CoverTab[5915]++
//line /snap/go/10455/src/net/file.go:14
	return "file+net"
//line /snap/go/10455/src/net/file.go:14
	// _ = "end of CoverTab[5915]"
//line /snap/go/10455/src/net/file.go:14
}
func (f fileAddr) String() string {
//line /snap/go/10455/src/net/file.go:15
	_go_fuzz_dep_.CoverTab[5916]++
//line /snap/go/10455/src/net/file.go:15
	return string(f)
//line /snap/go/10455/src/net/file.go:15
	// _ = "end of CoverTab[5916]"
//line /snap/go/10455/src/net/file.go:15
}

// FileConn returns a copy of the network connection corresponding to
//line /snap/go/10455/src/net/file.go:17
// the open file f.
//line /snap/go/10455/src/net/file.go:17
// It is the caller's responsibility to close f when finished.
//line /snap/go/10455/src/net/file.go:17
// Closing c does not affect f, and closing f does not affect c.
//line /snap/go/10455/src/net/file.go:21
func FileConn(f *os.File) (c Conn, err error) {
//line /snap/go/10455/src/net/file.go:21
	_go_fuzz_dep_.CoverTab[5917]++
						c, err = fileConn(f)
						if err != nil {
//line /snap/go/10455/src/net/file.go:23
		_go_fuzz_dep_.CoverTab[528291]++
//line /snap/go/10455/src/net/file.go:23
		_go_fuzz_dep_.CoverTab[5919]++
							err = &OpError{Op: "file", Net: "file+net", Source: nil, Addr: fileAddr(f.Name()), Err: err}
//line /snap/go/10455/src/net/file.go:24
		// _ = "end of CoverTab[5919]"
	} else {
//line /snap/go/10455/src/net/file.go:25
		_go_fuzz_dep_.CoverTab[528292]++
//line /snap/go/10455/src/net/file.go:25
		_go_fuzz_dep_.CoverTab[5920]++
//line /snap/go/10455/src/net/file.go:25
		// _ = "end of CoverTab[5920]"
//line /snap/go/10455/src/net/file.go:25
	}
//line /snap/go/10455/src/net/file.go:25
	// _ = "end of CoverTab[5917]"
//line /snap/go/10455/src/net/file.go:25
	_go_fuzz_dep_.CoverTab[5918]++
						return
//line /snap/go/10455/src/net/file.go:26
	// _ = "end of CoverTab[5918]"
}

// FileListener returns a copy of the network listener corresponding
//line /snap/go/10455/src/net/file.go:29
// to the open file f.
//line /snap/go/10455/src/net/file.go:29
// It is the caller's responsibility to close ln when finished.
//line /snap/go/10455/src/net/file.go:29
// Closing ln does not affect f, and closing f does not affect ln.
//line /snap/go/10455/src/net/file.go:33
func FileListener(f *os.File) (ln Listener, err error) {
//line /snap/go/10455/src/net/file.go:33
	_go_fuzz_dep_.CoverTab[5921]++
						ln, err = fileListener(f)
						if err != nil {
//line /snap/go/10455/src/net/file.go:35
		_go_fuzz_dep_.CoverTab[528293]++
//line /snap/go/10455/src/net/file.go:35
		_go_fuzz_dep_.CoverTab[5923]++
							err = &OpError{Op: "file", Net: "file+net", Source: nil, Addr: fileAddr(f.Name()), Err: err}
//line /snap/go/10455/src/net/file.go:36
		// _ = "end of CoverTab[5923]"
	} else {
//line /snap/go/10455/src/net/file.go:37
		_go_fuzz_dep_.CoverTab[528294]++
//line /snap/go/10455/src/net/file.go:37
		_go_fuzz_dep_.CoverTab[5924]++
//line /snap/go/10455/src/net/file.go:37
		// _ = "end of CoverTab[5924]"
//line /snap/go/10455/src/net/file.go:37
	}
//line /snap/go/10455/src/net/file.go:37
	// _ = "end of CoverTab[5921]"
//line /snap/go/10455/src/net/file.go:37
	_go_fuzz_dep_.CoverTab[5922]++
						return
//line /snap/go/10455/src/net/file.go:38
	// _ = "end of CoverTab[5922]"
}

// FilePacketConn returns a copy of the packet network connection
//line /snap/go/10455/src/net/file.go:41
// corresponding to the open file f.
//line /snap/go/10455/src/net/file.go:41
// It is the caller's responsibility to close f when finished.
//line /snap/go/10455/src/net/file.go:41
// Closing c does not affect f, and closing f does not affect c.
//line /snap/go/10455/src/net/file.go:45
func FilePacketConn(f *os.File) (c PacketConn, err error) {
//line /snap/go/10455/src/net/file.go:45
	_go_fuzz_dep_.CoverTab[5925]++
						c, err = filePacketConn(f)
						if err != nil {
//line /snap/go/10455/src/net/file.go:47
		_go_fuzz_dep_.CoverTab[528295]++
//line /snap/go/10455/src/net/file.go:47
		_go_fuzz_dep_.CoverTab[5927]++
							err = &OpError{Op: "file", Net: "file+net", Source: nil, Addr: fileAddr(f.Name()), Err: err}
//line /snap/go/10455/src/net/file.go:48
		// _ = "end of CoverTab[5927]"
	} else {
//line /snap/go/10455/src/net/file.go:49
		_go_fuzz_dep_.CoverTab[528296]++
//line /snap/go/10455/src/net/file.go:49
		_go_fuzz_dep_.CoverTab[5928]++
//line /snap/go/10455/src/net/file.go:49
		// _ = "end of CoverTab[5928]"
//line /snap/go/10455/src/net/file.go:49
	}
//line /snap/go/10455/src/net/file.go:49
	// _ = "end of CoverTab[5925]"
//line /snap/go/10455/src/net/file.go:49
	_go_fuzz_dep_.CoverTab[5926]++
						return
//line /snap/go/10455/src/net/file.go:50
	// _ = "end of CoverTab[5926]"
}

//line /snap/go/10455/src/net/file.go:51
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /snap/go/10455/src/net/file.go:51
var _ = _go_fuzz_dep_.CoverTab
