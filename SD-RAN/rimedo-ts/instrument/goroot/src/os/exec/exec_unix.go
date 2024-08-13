// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !plan9 && !windows

//line /usr/local/go/src/os/exec/exec_unix.go:7
package exec

//line /usr/local/go/src/os/exec/exec_unix.go:7
import (
//line /usr/local/go/src/os/exec/exec_unix.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/os/exec/exec_unix.go:7
)
//line /usr/local/go/src/os/exec/exec_unix.go:7
import (
//line /usr/local/go/src/os/exec/exec_unix.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/os/exec/exec_unix.go:7
)

import (
	"io/fs"
	"syscall"
)

// skipStdinCopyError optionally specifies a function which reports
//line /usr/local/go/src/os/exec/exec_unix.go:14
// whether the provided stdin copy error should be ignored.
//line /usr/local/go/src/os/exec/exec_unix.go:16
func skipStdinCopyError(err error) bool {
//line /usr/local/go/src/os/exec/exec_unix.go:16
	_go_fuzz_dep_.CoverTab[107470]++

//line /usr/local/go/src/os/exec/exec_unix.go:20
	pe, ok := err.(*fs.PathError)
	return ok && func() bool {
//line /usr/local/go/src/os/exec/exec_unix.go:21
		_go_fuzz_dep_.CoverTab[107471]++
//line /usr/local/go/src/os/exec/exec_unix.go:21
		return pe.Op == "write"
								// _ = "end of CoverTab[107471]"
//line /usr/local/go/src/os/exec/exec_unix.go:22
	}() && func() bool {
//line /usr/local/go/src/os/exec/exec_unix.go:22
		_go_fuzz_dep_.CoverTab[107472]++
//line /usr/local/go/src/os/exec/exec_unix.go:22
		return pe.Path == "|1"
//line /usr/local/go/src/os/exec/exec_unix.go:22
		// _ = "end of CoverTab[107472]"
//line /usr/local/go/src/os/exec/exec_unix.go:22
	}() && func() bool {
//line /usr/local/go/src/os/exec/exec_unix.go:22
		_go_fuzz_dep_.CoverTab[107473]++
//line /usr/local/go/src/os/exec/exec_unix.go:22
		return pe.Err == syscall.EPIPE
								// _ = "end of CoverTab[107473]"
//line /usr/local/go/src/os/exec/exec_unix.go:23
	}()
//line /usr/local/go/src/os/exec/exec_unix.go:23
	// _ = "end of CoverTab[107470]"
}

//line /usr/local/go/src/os/exec/exec_unix.go:24
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/os/exec/exec_unix.go:24
var _ = _go_fuzz_dep_.CoverTab
