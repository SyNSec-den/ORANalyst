// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix || (js && wasm)

//line /usr/local/go/src/path/filepath/path_unix.go:7
package filepath

//line /usr/local/go/src/path/filepath/path_unix.go:7
import (
//line /usr/local/go/src/path/filepath/path_unix.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/path/filepath/path_unix.go:7
)
//line /usr/local/go/src/path/filepath/path_unix.go:7
import (
//line /usr/local/go/src/path/filepath/path_unix.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/path/filepath/path_unix.go:7
)

import "strings"

func isLocal(path string) bool {
//line /usr/local/go/src/path/filepath/path_unix.go:11
	_go_fuzz_dep_.CoverTab[18225]++
							return unixIsLocal(path)
//line /usr/local/go/src/path/filepath/path_unix.go:12
	// _ = "end of CoverTab[18225]"
}

// IsAbs reports whether the path is absolute.
func IsAbs(path string) bool {
//line /usr/local/go/src/path/filepath/path_unix.go:16
	_go_fuzz_dep_.CoverTab[18226]++
							return strings.HasPrefix(path, "/")
//line /usr/local/go/src/path/filepath/path_unix.go:17
	// _ = "end of CoverTab[18226]"
}

// volumeNameLen returns length of the leading volume name on Windows.
//line /usr/local/go/src/path/filepath/path_unix.go:20
// It returns 0 elsewhere.
//line /usr/local/go/src/path/filepath/path_unix.go:22
func volumeNameLen(path string) int {
//line /usr/local/go/src/path/filepath/path_unix.go:22
	_go_fuzz_dep_.CoverTab[18227]++
							return 0
//line /usr/local/go/src/path/filepath/path_unix.go:23
	// _ = "end of CoverTab[18227]"
}

// HasPrefix exists for historical compatibility and should not be used.
//line /usr/local/go/src/path/filepath/path_unix.go:26
//
//line /usr/local/go/src/path/filepath/path_unix.go:26
// Deprecated: HasPrefix does not respect path boundaries and
//line /usr/local/go/src/path/filepath/path_unix.go:26
// does not ignore case when required.
//line /usr/local/go/src/path/filepath/path_unix.go:30
func HasPrefix(p, prefix string) bool {
//line /usr/local/go/src/path/filepath/path_unix.go:30
	_go_fuzz_dep_.CoverTab[18228]++
							return strings.HasPrefix(p, prefix)
//line /usr/local/go/src/path/filepath/path_unix.go:31
	// _ = "end of CoverTab[18228]"
}

func splitList(path string) []string {
//line /usr/local/go/src/path/filepath/path_unix.go:34
	_go_fuzz_dep_.CoverTab[18229]++
							if path == "" {
//line /usr/local/go/src/path/filepath/path_unix.go:35
		_go_fuzz_dep_.CoverTab[18231]++
								return []string{}
//line /usr/local/go/src/path/filepath/path_unix.go:36
		// _ = "end of CoverTab[18231]"
	} else {
//line /usr/local/go/src/path/filepath/path_unix.go:37
		_go_fuzz_dep_.CoverTab[18232]++
//line /usr/local/go/src/path/filepath/path_unix.go:37
		// _ = "end of CoverTab[18232]"
//line /usr/local/go/src/path/filepath/path_unix.go:37
	}
//line /usr/local/go/src/path/filepath/path_unix.go:37
	// _ = "end of CoverTab[18229]"
//line /usr/local/go/src/path/filepath/path_unix.go:37
	_go_fuzz_dep_.CoverTab[18230]++
							return strings.Split(path, string(ListSeparator))
//line /usr/local/go/src/path/filepath/path_unix.go:38
	// _ = "end of CoverTab[18230]"
}

func abs(path string) (string, error) {
//line /usr/local/go/src/path/filepath/path_unix.go:41
	_go_fuzz_dep_.CoverTab[18233]++
							return unixAbs(path)
//line /usr/local/go/src/path/filepath/path_unix.go:42
	// _ = "end of CoverTab[18233]"
}

func join(elem []string) string {
//line /usr/local/go/src/path/filepath/path_unix.go:45
	_go_fuzz_dep_.CoverTab[18234]++

							for i, e := range elem {
//line /usr/local/go/src/path/filepath/path_unix.go:47
		_go_fuzz_dep_.CoverTab[18236]++
								if e != "" {
//line /usr/local/go/src/path/filepath/path_unix.go:48
			_go_fuzz_dep_.CoverTab[18237]++
									return Clean(strings.Join(elem[i:], string(Separator)))
//line /usr/local/go/src/path/filepath/path_unix.go:49
			// _ = "end of CoverTab[18237]"
		} else {
//line /usr/local/go/src/path/filepath/path_unix.go:50
			_go_fuzz_dep_.CoverTab[18238]++
//line /usr/local/go/src/path/filepath/path_unix.go:50
			// _ = "end of CoverTab[18238]"
//line /usr/local/go/src/path/filepath/path_unix.go:50
		}
//line /usr/local/go/src/path/filepath/path_unix.go:50
		// _ = "end of CoverTab[18236]"
	}
//line /usr/local/go/src/path/filepath/path_unix.go:51
	// _ = "end of CoverTab[18234]"
//line /usr/local/go/src/path/filepath/path_unix.go:51
	_go_fuzz_dep_.CoverTab[18235]++
							return ""
//line /usr/local/go/src/path/filepath/path_unix.go:52
	// _ = "end of CoverTab[18235]"
}

func sameWord(a, b string) bool {
//line /usr/local/go/src/path/filepath/path_unix.go:55
	_go_fuzz_dep_.CoverTab[18239]++
							return a == b
//line /usr/local/go/src/path/filepath/path_unix.go:56
	// _ = "end of CoverTab[18239]"
}

//line /usr/local/go/src/path/filepath/path_unix.go:57
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/path/filepath/path_unix.go:57
var _ = _go_fuzz_dep_.CoverTab
