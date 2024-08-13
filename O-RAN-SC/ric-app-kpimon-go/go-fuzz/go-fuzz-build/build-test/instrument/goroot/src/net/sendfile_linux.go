// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /snap/go/10455/src/net/sendfile_linux.go:5
package net

//line /snap/go/10455/src/net/sendfile_linux.go:5
import (
//line /snap/go/10455/src/net/sendfile_linux.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /snap/go/10455/src/net/sendfile_linux.go:5
)
//line /snap/go/10455/src/net/sendfile_linux.go:5
import (
//line /snap/go/10455/src/net/sendfile_linux.go:5
	_atomic_ "sync/atomic"
//line /snap/go/10455/src/net/sendfile_linux.go:5
)

import (
	"internal/poll"
	"io"
	"os"
)

// sendFile copies the contents of r to c using the sendfile
//line /snap/go/10455/src/net/sendfile_linux.go:13
// system call to minimize copies.
//line /snap/go/10455/src/net/sendfile_linux.go:13
//
//line /snap/go/10455/src/net/sendfile_linux.go:13
// if handled == true, sendFile returns the number (potentially zero) of bytes
//line /snap/go/10455/src/net/sendfile_linux.go:13
// copied and any non-EOF error.
//line /snap/go/10455/src/net/sendfile_linux.go:13
//
//line /snap/go/10455/src/net/sendfile_linux.go:13
// if handled == false, sendFile performed no work.
//line /snap/go/10455/src/net/sendfile_linux.go:20
func sendFile(c *netFD, r io.Reader) (written int64, err error, handled bool) {
//line /snap/go/10455/src/net/sendfile_linux.go:20
	_go_fuzz_dep_.CoverTab[7985]++
							var remain int64 = 1<<63 - 1	// by default, copy until EOF

							lr, ok := r.(*io.LimitedReader)
							if ok {
//line /snap/go/10455/src/net/sendfile_linux.go:24
		_go_fuzz_dep_.CoverTab[529596]++
//line /snap/go/10455/src/net/sendfile_linux.go:24
		_go_fuzz_dep_.CoverTab[7992]++
								remain, r = lr.N, lr.R
								if remain <= 0 {
//line /snap/go/10455/src/net/sendfile_linux.go:26
			_go_fuzz_dep_.CoverTab[529598]++
//line /snap/go/10455/src/net/sendfile_linux.go:26
			_go_fuzz_dep_.CoverTab[7993]++
									return 0, nil, true
//line /snap/go/10455/src/net/sendfile_linux.go:27
			// _ = "end of CoverTab[7993]"
		} else {
//line /snap/go/10455/src/net/sendfile_linux.go:28
			_go_fuzz_dep_.CoverTab[529599]++
//line /snap/go/10455/src/net/sendfile_linux.go:28
			_go_fuzz_dep_.CoverTab[7994]++
//line /snap/go/10455/src/net/sendfile_linux.go:28
			// _ = "end of CoverTab[7994]"
//line /snap/go/10455/src/net/sendfile_linux.go:28
		}
//line /snap/go/10455/src/net/sendfile_linux.go:28
		// _ = "end of CoverTab[7992]"
	} else {
//line /snap/go/10455/src/net/sendfile_linux.go:29
		_go_fuzz_dep_.CoverTab[529597]++
//line /snap/go/10455/src/net/sendfile_linux.go:29
		_go_fuzz_dep_.CoverTab[7995]++
//line /snap/go/10455/src/net/sendfile_linux.go:29
		// _ = "end of CoverTab[7995]"
//line /snap/go/10455/src/net/sendfile_linux.go:29
	}
//line /snap/go/10455/src/net/sendfile_linux.go:29
	// _ = "end of CoverTab[7985]"
//line /snap/go/10455/src/net/sendfile_linux.go:29
	_go_fuzz_dep_.CoverTab[7986]++
							f, ok := r.(*os.File)
							if !ok {
//line /snap/go/10455/src/net/sendfile_linux.go:31
		_go_fuzz_dep_.CoverTab[529600]++
//line /snap/go/10455/src/net/sendfile_linux.go:31
		_go_fuzz_dep_.CoverTab[7996]++
								return 0, nil, false
//line /snap/go/10455/src/net/sendfile_linux.go:32
		// _ = "end of CoverTab[7996]"
	} else {
//line /snap/go/10455/src/net/sendfile_linux.go:33
		_go_fuzz_dep_.CoverTab[529601]++
//line /snap/go/10455/src/net/sendfile_linux.go:33
		_go_fuzz_dep_.CoverTab[7997]++
//line /snap/go/10455/src/net/sendfile_linux.go:33
		// _ = "end of CoverTab[7997]"
//line /snap/go/10455/src/net/sendfile_linux.go:33
	}
//line /snap/go/10455/src/net/sendfile_linux.go:33
	// _ = "end of CoverTab[7986]"
//line /snap/go/10455/src/net/sendfile_linux.go:33
	_go_fuzz_dep_.CoverTab[7987]++

							sc, err := f.SyscallConn()
							if err != nil {
//line /snap/go/10455/src/net/sendfile_linux.go:36
		_go_fuzz_dep_.CoverTab[529602]++
//line /snap/go/10455/src/net/sendfile_linux.go:36
		_go_fuzz_dep_.CoverTab[7998]++
								return 0, nil, false
//line /snap/go/10455/src/net/sendfile_linux.go:37
		// _ = "end of CoverTab[7998]"
	} else {
//line /snap/go/10455/src/net/sendfile_linux.go:38
		_go_fuzz_dep_.CoverTab[529603]++
//line /snap/go/10455/src/net/sendfile_linux.go:38
		_go_fuzz_dep_.CoverTab[7999]++
//line /snap/go/10455/src/net/sendfile_linux.go:38
		// _ = "end of CoverTab[7999]"
//line /snap/go/10455/src/net/sendfile_linux.go:38
	}
//line /snap/go/10455/src/net/sendfile_linux.go:38
	// _ = "end of CoverTab[7987]"
//line /snap/go/10455/src/net/sendfile_linux.go:38
	_go_fuzz_dep_.CoverTab[7988]++

							var werr error
							err = sc.Read(func(fd uintptr) bool {
//line /snap/go/10455/src/net/sendfile_linux.go:41
		_go_fuzz_dep_.CoverTab[8000]++
								written, werr, handled = poll.SendFile(&c.pfd, int(fd), remain)
								return true
//line /snap/go/10455/src/net/sendfile_linux.go:43
		// _ = "end of CoverTab[8000]"
	})
//line /snap/go/10455/src/net/sendfile_linux.go:44
	// _ = "end of CoverTab[7988]"
//line /snap/go/10455/src/net/sendfile_linux.go:44
	_go_fuzz_dep_.CoverTab[7989]++
							if err == nil {
//line /snap/go/10455/src/net/sendfile_linux.go:45
		_go_fuzz_dep_.CoverTab[529604]++
//line /snap/go/10455/src/net/sendfile_linux.go:45
		_go_fuzz_dep_.CoverTab[8001]++
								err = werr
//line /snap/go/10455/src/net/sendfile_linux.go:46
		// _ = "end of CoverTab[8001]"
	} else {
//line /snap/go/10455/src/net/sendfile_linux.go:47
		_go_fuzz_dep_.CoverTab[529605]++
//line /snap/go/10455/src/net/sendfile_linux.go:47
		_go_fuzz_dep_.CoverTab[8002]++
//line /snap/go/10455/src/net/sendfile_linux.go:47
		// _ = "end of CoverTab[8002]"
//line /snap/go/10455/src/net/sendfile_linux.go:47
	}
//line /snap/go/10455/src/net/sendfile_linux.go:47
	// _ = "end of CoverTab[7989]"
//line /snap/go/10455/src/net/sendfile_linux.go:47
	_go_fuzz_dep_.CoverTab[7990]++

							if lr != nil {
//line /snap/go/10455/src/net/sendfile_linux.go:49
		_go_fuzz_dep_.CoverTab[529606]++
//line /snap/go/10455/src/net/sendfile_linux.go:49
		_go_fuzz_dep_.CoverTab[8003]++
								lr.N = remain - written
//line /snap/go/10455/src/net/sendfile_linux.go:50
		// _ = "end of CoverTab[8003]"
	} else {
//line /snap/go/10455/src/net/sendfile_linux.go:51
		_go_fuzz_dep_.CoverTab[529607]++
//line /snap/go/10455/src/net/sendfile_linux.go:51
		_go_fuzz_dep_.CoverTab[8004]++
//line /snap/go/10455/src/net/sendfile_linux.go:51
		// _ = "end of CoverTab[8004]"
//line /snap/go/10455/src/net/sendfile_linux.go:51
	}
//line /snap/go/10455/src/net/sendfile_linux.go:51
	// _ = "end of CoverTab[7990]"
//line /snap/go/10455/src/net/sendfile_linux.go:51
	_go_fuzz_dep_.CoverTab[7991]++
							return written, wrapSyscallError("sendfile", err), handled
//line /snap/go/10455/src/net/sendfile_linux.go:52
	// _ = "end of CoverTab[7991]"
}

//line /snap/go/10455/src/net/sendfile_linux.go:53
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /snap/go/10455/src/net/sendfile_linux.go:53
var _ = _go_fuzz_dep_.CoverTab
