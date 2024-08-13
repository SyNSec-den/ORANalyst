// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /snap/go/10455/src/net/splice_linux.go:5
package net

//line /snap/go/10455/src/net/splice_linux.go:5
import (
//line /snap/go/10455/src/net/splice_linux.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /snap/go/10455/src/net/splice_linux.go:5
)
//line /snap/go/10455/src/net/splice_linux.go:5
import (
//line /snap/go/10455/src/net/splice_linux.go:5
	_atomic_ "sync/atomic"
//line /snap/go/10455/src/net/splice_linux.go:5
)

import (
	"internal/poll"
	"io"
)

// splice transfers data from r to c using the splice system call to minimize
//line /snap/go/10455/src/net/splice_linux.go:12
// copies from and to userspace. c must be a TCP connection. Currently, splice
//line /snap/go/10455/src/net/splice_linux.go:12
// is only enabled if r is a TCP or a stream-oriented Unix connection.
//line /snap/go/10455/src/net/splice_linux.go:12
//
//line /snap/go/10455/src/net/splice_linux.go:12
// If splice returns handled == false, it has performed no work.
//line /snap/go/10455/src/net/splice_linux.go:17
func splice(c *netFD, r io.Reader) (written int64, err error, handled bool) {
//line /snap/go/10455/src/net/splice_linux.go:17
	_go_fuzz_dep_.CoverTab[8278]++
							var remain int64 = 1<<63 - 1	// by default, copy until EOF
							lr, ok := r.(*io.LimitedReader)
							if ok {
//line /snap/go/10455/src/net/splice_linux.go:20
		_go_fuzz_dep_.CoverTab[529795]++
//line /snap/go/10455/src/net/splice_linux.go:20
		_go_fuzz_dep_.CoverTab[8282]++
								remain, r = lr.N, lr.R
								if remain <= 0 {
//line /snap/go/10455/src/net/splice_linux.go:22
			_go_fuzz_dep_.CoverTab[529797]++
//line /snap/go/10455/src/net/splice_linux.go:22
			_go_fuzz_dep_.CoverTab[8283]++
									return 0, nil, true
//line /snap/go/10455/src/net/splice_linux.go:23
			// _ = "end of CoverTab[8283]"
		} else {
//line /snap/go/10455/src/net/splice_linux.go:24
			_go_fuzz_dep_.CoverTab[529798]++
//line /snap/go/10455/src/net/splice_linux.go:24
			_go_fuzz_dep_.CoverTab[8284]++
//line /snap/go/10455/src/net/splice_linux.go:24
			// _ = "end of CoverTab[8284]"
//line /snap/go/10455/src/net/splice_linux.go:24
		}
//line /snap/go/10455/src/net/splice_linux.go:24
		// _ = "end of CoverTab[8282]"
	} else {
//line /snap/go/10455/src/net/splice_linux.go:25
		_go_fuzz_dep_.CoverTab[529796]++
//line /snap/go/10455/src/net/splice_linux.go:25
		_go_fuzz_dep_.CoverTab[8285]++
//line /snap/go/10455/src/net/splice_linux.go:25
		// _ = "end of CoverTab[8285]"
//line /snap/go/10455/src/net/splice_linux.go:25
	}
//line /snap/go/10455/src/net/splice_linux.go:25
	// _ = "end of CoverTab[8278]"
//line /snap/go/10455/src/net/splice_linux.go:25
	_go_fuzz_dep_.CoverTab[8279]++

							var s *netFD
							if tc, ok := r.(*TCPConn); ok {
//line /snap/go/10455/src/net/splice_linux.go:28
		_go_fuzz_dep_.CoverTab[529799]++
//line /snap/go/10455/src/net/splice_linux.go:28
		_go_fuzz_dep_.CoverTab[8286]++
								s = tc.fd
//line /snap/go/10455/src/net/splice_linux.go:29
		// _ = "end of CoverTab[8286]"
	} else {
//line /snap/go/10455/src/net/splice_linux.go:30
		_go_fuzz_dep_.CoverTab[529800]++
//line /snap/go/10455/src/net/splice_linux.go:30
		_go_fuzz_dep_.CoverTab[8287]++
//line /snap/go/10455/src/net/splice_linux.go:30
		if uc, ok := r.(*UnixConn); ok {
//line /snap/go/10455/src/net/splice_linux.go:30
			_go_fuzz_dep_.CoverTab[529801]++
//line /snap/go/10455/src/net/splice_linux.go:30
			_go_fuzz_dep_.CoverTab[8288]++
									if uc.fd.net != "unix" {
//line /snap/go/10455/src/net/splice_linux.go:31
				_go_fuzz_dep_.CoverTab[529803]++
//line /snap/go/10455/src/net/splice_linux.go:31
				_go_fuzz_dep_.CoverTab[8290]++
										return 0, nil, false
//line /snap/go/10455/src/net/splice_linux.go:32
				// _ = "end of CoverTab[8290]"
			} else {
//line /snap/go/10455/src/net/splice_linux.go:33
				_go_fuzz_dep_.CoverTab[529804]++
//line /snap/go/10455/src/net/splice_linux.go:33
				_go_fuzz_dep_.CoverTab[8291]++
//line /snap/go/10455/src/net/splice_linux.go:33
				// _ = "end of CoverTab[8291]"
//line /snap/go/10455/src/net/splice_linux.go:33
			}
//line /snap/go/10455/src/net/splice_linux.go:33
			// _ = "end of CoverTab[8288]"
//line /snap/go/10455/src/net/splice_linux.go:33
			_go_fuzz_dep_.CoverTab[8289]++
									s = uc.fd
//line /snap/go/10455/src/net/splice_linux.go:34
			// _ = "end of CoverTab[8289]"
		} else {
//line /snap/go/10455/src/net/splice_linux.go:35
			_go_fuzz_dep_.CoverTab[529802]++
//line /snap/go/10455/src/net/splice_linux.go:35
			_go_fuzz_dep_.CoverTab[8292]++
									return 0, nil, false
//line /snap/go/10455/src/net/splice_linux.go:36
			// _ = "end of CoverTab[8292]"
		}
//line /snap/go/10455/src/net/splice_linux.go:37
		// _ = "end of CoverTab[8287]"
//line /snap/go/10455/src/net/splice_linux.go:37
	}
//line /snap/go/10455/src/net/splice_linux.go:37
	// _ = "end of CoverTab[8279]"
//line /snap/go/10455/src/net/splice_linux.go:37
	_go_fuzz_dep_.CoverTab[8280]++

							written, handled, sc, err := poll.Splice(&c.pfd, &s.pfd, remain)
							if lr != nil {
//line /snap/go/10455/src/net/splice_linux.go:40
		_go_fuzz_dep_.CoverTab[529805]++
//line /snap/go/10455/src/net/splice_linux.go:40
		_go_fuzz_dep_.CoverTab[8293]++
								lr.N -= written
//line /snap/go/10455/src/net/splice_linux.go:41
		// _ = "end of CoverTab[8293]"
	} else {
//line /snap/go/10455/src/net/splice_linux.go:42
		_go_fuzz_dep_.CoverTab[529806]++
//line /snap/go/10455/src/net/splice_linux.go:42
		_go_fuzz_dep_.CoverTab[8294]++
//line /snap/go/10455/src/net/splice_linux.go:42
		// _ = "end of CoverTab[8294]"
//line /snap/go/10455/src/net/splice_linux.go:42
	}
//line /snap/go/10455/src/net/splice_linux.go:42
	// _ = "end of CoverTab[8280]"
//line /snap/go/10455/src/net/splice_linux.go:42
	_go_fuzz_dep_.CoverTab[8281]++
							return written, wrapSyscallError(sc, err), handled
//line /snap/go/10455/src/net/splice_linux.go:43
	// _ = "end of CoverTab[8281]"
}

//line /snap/go/10455/src/net/splice_linux.go:44
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /snap/go/10455/src/net/splice_linux.go:44
var _ = _go_fuzz_dep_.CoverTab
