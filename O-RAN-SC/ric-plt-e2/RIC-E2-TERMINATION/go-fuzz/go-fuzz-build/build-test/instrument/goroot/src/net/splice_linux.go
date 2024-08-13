// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/net/splice_linux.go:5
package net

//line /usr/local/go/src/net/splice_linux.go:5
import (
//line /usr/local/go/src/net/splice_linux.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/splice_linux.go:5
)
//line /usr/local/go/src/net/splice_linux.go:5
import (
//line /usr/local/go/src/net/splice_linux.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/splice_linux.go:5
)

import (
	"internal/poll"
	"io"
)

// splice transfers data from r to c using the splice system call to minimize
//line /usr/local/go/src/net/splice_linux.go:12
// copies from and to userspace. c must be a TCP connection. Currently, splice
//line /usr/local/go/src/net/splice_linux.go:12
// is only enabled if r is a TCP or a stream-oriented Unix connection.
//line /usr/local/go/src/net/splice_linux.go:12
//
//line /usr/local/go/src/net/splice_linux.go:12
// If splice returns handled == false, it has performed no work.
//line /usr/local/go/src/net/splice_linux.go:17
func splice(c *netFD, r io.Reader) (written int64, err error, handled bool) {
//line /usr/local/go/src/net/splice_linux.go:17
	_go_fuzz_dep_.CoverTab[7984]++
							var remain int64 = 1 << 62	// by default, copy until EOF
							lr, ok := r.(*io.LimitedReader)
							if ok {
//line /usr/local/go/src/net/splice_linux.go:20
		_go_fuzz_dep_.CoverTab[7988]++
								remain, r = lr.N, lr.R
								if remain <= 0 {
//line /usr/local/go/src/net/splice_linux.go:22
			_go_fuzz_dep_.CoverTab[7989]++
									return 0, nil, true
//line /usr/local/go/src/net/splice_linux.go:23
			// _ = "end of CoverTab[7989]"
		} else {
//line /usr/local/go/src/net/splice_linux.go:24
			_go_fuzz_dep_.CoverTab[7990]++
//line /usr/local/go/src/net/splice_linux.go:24
			// _ = "end of CoverTab[7990]"
//line /usr/local/go/src/net/splice_linux.go:24
		}
//line /usr/local/go/src/net/splice_linux.go:24
		// _ = "end of CoverTab[7988]"
	} else {
//line /usr/local/go/src/net/splice_linux.go:25
		_go_fuzz_dep_.CoverTab[7991]++
//line /usr/local/go/src/net/splice_linux.go:25
		// _ = "end of CoverTab[7991]"
//line /usr/local/go/src/net/splice_linux.go:25
	}
//line /usr/local/go/src/net/splice_linux.go:25
	// _ = "end of CoverTab[7984]"
//line /usr/local/go/src/net/splice_linux.go:25
	_go_fuzz_dep_.CoverTab[7985]++

							var s *netFD
							if tc, ok := r.(*TCPConn); ok {
//line /usr/local/go/src/net/splice_linux.go:28
		_go_fuzz_dep_.CoverTab[7992]++
								s = tc.fd
//line /usr/local/go/src/net/splice_linux.go:29
		// _ = "end of CoverTab[7992]"
	} else {
//line /usr/local/go/src/net/splice_linux.go:30
		_go_fuzz_dep_.CoverTab[7993]++
//line /usr/local/go/src/net/splice_linux.go:30
		if uc, ok := r.(*UnixConn); ok {
//line /usr/local/go/src/net/splice_linux.go:30
			_go_fuzz_dep_.CoverTab[7994]++
									if uc.fd.net != "unix" {
//line /usr/local/go/src/net/splice_linux.go:31
				_go_fuzz_dep_.CoverTab[7996]++
										return 0, nil, false
//line /usr/local/go/src/net/splice_linux.go:32
				// _ = "end of CoverTab[7996]"
			} else {
//line /usr/local/go/src/net/splice_linux.go:33
				_go_fuzz_dep_.CoverTab[7997]++
//line /usr/local/go/src/net/splice_linux.go:33
				// _ = "end of CoverTab[7997]"
//line /usr/local/go/src/net/splice_linux.go:33
			}
//line /usr/local/go/src/net/splice_linux.go:33
			// _ = "end of CoverTab[7994]"
//line /usr/local/go/src/net/splice_linux.go:33
			_go_fuzz_dep_.CoverTab[7995]++
									s = uc.fd
//line /usr/local/go/src/net/splice_linux.go:34
			// _ = "end of CoverTab[7995]"
		} else {
//line /usr/local/go/src/net/splice_linux.go:35
			_go_fuzz_dep_.CoverTab[7998]++
									return 0, nil, false
//line /usr/local/go/src/net/splice_linux.go:36
			// _ = "end of CoverTab[7998]"
		}
//line /usr/local/go/src/net/splice_linux.go:37
		// _ = "end of CoverTab[7993]"
//line /usr/local/go/src/net/splice_linux.go:37
	}
//line /usr/local/go/src/net/splice_linux.go:37
	// _ = "end of CoverTab[7985]"
//line /usr/local/go/src/net/splice_linux.go:37
	_go_fuzz_dep_.CoverTab[7986]++

							written, handled, sc, err := poll.Splice(&c.pfd, &s.pfd, remain)
							if lr != nil {
//line /usr/local/go/src/net/splice_linux.go:40
		_go_fuzz_dep_.CoverTab[7999]++
								lr.N -= written
//line /usr/local/go/src/net/splice_linux.go:41
		// _ = "end of CoverTab[7999]"
	} else {
//line /usr/local/go/src/net/splice_linux.go:42
		_go_fuzz_dep_.CoverTab[8000]++
//line /usr/local/go/src/net/splice_linux.go:42
		// _ = "end of CoverTab[8000]"
//line /usr/local/go/src/net/splice_linux.go:42
	}
//line /usr/local/go/src/net/splice_linux.go:42
	// _ = "end of CoverTab[7986]"
//line /usr/local/go/src/net/splice_linux.go:42
	_go_fuzz_dep_.CoverTab[7987]++
							return written, wrapSyscallError(sc, err), handled
//line /usr/local/go/src/net/splice_linux.go:43
	// _ = "end of CoverTab[7987]"
}

//line /usr/local/go/src/net/splice_linux.go:44
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/splice_linux.go:44
var _ = _go_fuzz_dep_.CoverTab
