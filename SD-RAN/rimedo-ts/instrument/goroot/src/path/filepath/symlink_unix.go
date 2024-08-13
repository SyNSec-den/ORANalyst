// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !windows && !plan9

//line /usr/local/go/src/path/filepath/symlink_unix.go:7
package filepath

//line /usr/local/go/src/path/filepath/symlink_unix.go:7
import (
//line /usr/local/go/src/path/filepath/symlink_unix.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/path/filepath/symlink_unix.go:7
)
//line /usr/local/go/src/path/filepath/symlink_unix.go:7
import (
//line /usr/local/go/src/path/filepath/symlink_unix.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/path/filepath/symlink_unix.go:7
)

func evalSymlinks(path string) (string, error) {
//line /usr/local/go/src/path/filepath/symlink_unix.go:9
	_go_fuzz_dep_.CoverTab[18313]++
								return walkSymlinks(path)
//line /usr/local/go/src/path/filepath/symlink_unix.go:10
	// _ = "end of CoverTab[18313]"
}

//line /usr/local/go/src/path/filepath/symlink_unix.go:11
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/path/filepath/symlink_unix.go:11
var _ = _go_fuzz_dep_.CoverTab
