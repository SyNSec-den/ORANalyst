// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/net/sendfile_linux.go:5
package net

//line /usr/local/go/src/net/sendfile_linux.go:5
import (
//line /usr/local/go/src/net/sendfile_linux.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/sendfile_linux.go:5
)
//line /usr/local/go/src/net/sendfile_linux.go:5
import (
//line /usr/local/go/src/net/sendfile_linux.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/sendfile_linux.go:5
)

import (
	"internal/poll"
	"io"
	"os"
)

// sendFile copies the contents of r to c using the sendfile
//line /usr/local/go/src/net/sendfile_linux.go:13
// system call to minimize copies.
//line /usr/local/go/src/net/sendfile_linux.go:13
//
//line /usr/local/go/src/net/sendfile_linux.go:13
// if handled == true, sendFile returns the number (potentially zero) of bytes
//line /usr/local/go/src/net/sendfile_linux.go:13
// copied and any non-EOF error.
//line /usr/local/go/src/net/sendfile_linux.go:13
//
//line /usr/local/go/src/net/sendfile_linux.go:13
// if handled == false, sendFile performed no work.
//line /usr/local/go/src/net/sendfile_linux.go:20
func sendFile(c *netFD, r io.Reader) (written int64, err error, handled bool) {
//line /usr/local/go/src/net/sendfile_linux.go:20
	_go_fuzz_dep_.CoverTab[7701]++
							var remain int64 = 1 << 62	// by default, copy until EOF

							lr, ok := r.(*io.LimitedReader)
							if ok {
//line /usr/local/go/src/net/sendfile_linux.go:24
		_go_fuzz_dep_.CoverTab[7708]++
								remain, r = lr.N, lr.R
								if remain <= 0 {
//line /usr/local/go/src/net/sendfile_linux.go:26
			_go_fuzz_dep_.CoverTab[7709]++
									return 0, nil, true
//line /usr/local/go/src/net/sendfile_linux.go:27
			// _ = "end of CoverTab[7709]"
		} else {
//line /usr/local/go/src/net/sendfile_linux.go:28
			_go_fuzz_dep_.CoverTab[7710]++
//line /usr/local/go/src/net/sendfile_linux.go:28
			// _ = "end of CoverTab[7710]"
//line /usr/local/go/src/net/sendfile_linux.go:28
		}
//line /usr/local/go/src/net/sendfile_linux.go:28
		// _ = "end of CoverTab[7708]"
	} else {
//line /usr/local/go/src/net/sendfile_linux.go:29
		_go_fuzz_dep_.CoverTab[7711]++
//line /usr/local/go/src/net/sendfile_linux.go:29
		// _ = "end of CoverTab[7711]"
//line /usr/local/go/src/net/sendfile_linux.go:29
	}
//line /usr/local/go/src/net/sendfile_linux.go:29
	// _ = "end of CoverTab[7701]"
//line /usr/local/go/src/net/sendfile_linux.go:29
	_go_fuzz_dep_.CoverTab[7702]++
							f, ok := r.(*os.File)
							if !ok {
//line /usr/local/go/src/net/sendfile_linux.go:31
		_go_fuzz_dep_.CoverTab[7712]++
								return 0, nil, false
//line /usr/local/go/src/net/sendfile_linux.go:32
		// _ = "end of CoverTab[7712]"
	} else {
//line /usr/local/go/src/net/sendfile_linux.go:33
		_go_fuzz_dep_.CoverTab[7713]++
//line /usr/local/go/src/net/sendfile_linux.go:33
		// _ = "end of CoverTab[7713]"
//line /usr/local/go/src/net/sendfile_linux.go:33
	}
//line /usr/local/go/src/net/sendfile_linux.go:33
	// _ = "end of CoverTab[7702]"
//line /usr/local/go/src/net/sendfile_linux.go:33
	_go_fuzz_dep_.CoverTab[7703]++

							sc, err := f.SyscallConn()
							if err != nil {
//line /usr/local/go/src/net/sendfile_linux.go:36
		_go_fuzz_dep_.CoverTab[7714]++
								return 0, nil, false
//line /usr/local/go/src/net/sendfile_linux.go:37
		// _ = "end of CoverTab[7714]"
	} else {
//line /usr/local/go/src/net/sendfile_linux.go:38
		_go_fuzz_dep_.CoverTab[7715]++
//line /usr/local/go/src/net/sendfile_linux.go:38
		// _ = "end of CoverTab[7715]"
//line /usr/local/go/src/net/sendfile_linux.go:38
	}
//line /usr/local/go/src/net/sendfile_linux.go:38
	// _ = "end of CoverTab[7703]"
//line /usr/local/go/src/net/sendfile_linux.go:38
	_go_fuzz_dep_.CoverTab[7704]++

							var werr error
							err = sc.Read(func(fd uintptr) bool {
//line /usr/local/go/src/net/sendfile_linux.go:41
		_go_fuzz_dep_.CoverTab[7716]++
								written, werr, handled = poll.SendFile(&c.pfd, int(fd), remain)
								return true
//line /usr/local/go/src/net/sendfile_linux.go:43
		// _ = "end of CoverTab[7716]"
	})
//line /usr/local/go/src/net/sendfile_linux.go:44
	// _ = "end of CoverTab[7704]"
//line /usr/local/go/src/net/sendfile_linux.go:44
	_go_fuzz_dep_.CoverTab[7705]++
							if err == nil {
//line /usr/local/go/src/net/sendfile_linux.go:45
		_go_fuzz_dep_.CoverTab[7717]++
								err = werr
//line /usr/local/go/src/net/sendfile_linux.go:46
		// _ = "end of CoverTab[7717]"
	} else {
//line /usr/local/go/src/net/sendfile_linux.go:47
		_go_fuzz_dep_.CoverTab[7718]++
//line /usr/local/go/src/net/sendfile_linux.go:47
		// _ = "end of CoverTab[7718]"
//line /usr/local/go/src/net/sendfile_linux.go:47
	}
//line /usr/local/go/src/net/sendfile_linux.go:47
	// _ = "end of CoverTab[7705]"
//line /usr/local/go/src/net/sendfile_linux.go:47
	_go_fuzz_dep_.CoverTab[7706]++

							if lr != nil {
//line /usr/local/go/src/net/sendfile_linux.go:49
		_go_fuzz_dep_.CoverTab[7719]++
								lr.N = remain - written
//line /usr/local/go/src/net/sendfile_linux.go:50
		// _ = "end of CoverTab[7719]"
	} else {
//line /usr/local/go/src/net/sendfile_linux.go:51
		_go_fuzz_dep_.CoverTab[7720]++
//line /usr/local/go/src/net/sendfile_linux.go:51
		// _ = "end of CoverTab[7720]"
//line /usr/local/go/src/net/sendfile_linux.go:51
	}
//line /usr/local/go/src/net/sendfile_linux.go:51
	// _ = "end of CoverTab[7706]"
//line /usr/local/go/src/net/sendfile_linux.go:51
	_go_fuzz_dep_.CoverTab[7707]++
							return written, wrapSyscallError("sendfile", err), handled
//line /usr/local/go/src/net/sendfile_linux.go:52
	// _ = "end of CoverTab[7707]"
}

//line /usr/local/go/src/net/sendfile_linux.go:53
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/sendfile_linux.go:53
var _ = _go_fuzz_dep_.CoverTab
