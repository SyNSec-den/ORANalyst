// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /snap/go/10455/src/strings/compare.go:5
package strings

//line /snap/go/10455/src/strings/compare.go:5
import (
//line /snap/go/10455/src/strings/compare.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /snap/go/10455/src/strings/compare.go:5
)
//line /snap/go/10455/src/strings/compare.go:5
import (
//line /snap/go/10455/src/strings/compare.go:5
	_atomic_ "sync/atomic"
//line /snap/go/10455/src/strings/compare.go:5
)

// Compare returns an integer comparing two strings lexicographically.
//line /snap/go/10455/src/strings/compare.go:7
// The result will be 0 if a == b, -1 if a < b, and +1 if a > b.
//line /snap/go/10455/src/strings/compare.go:7
//
//line /snap/go/10455/src/strings/compare.go:7
// Compare is included only for symmetry with package bytes.
//line /snap/go/10455/src/strings/compare.go:7
// It is usually clearer and always faster to use the built-in
//line /snap/go/10455/src/strings/compare.go:7
// string comparison operators ==, <, >, and so on.
//line /snap/go/10455/src/strings/compare.go:13
func Compare(a, b string) int {
//line /snap/go/10455/src/strings/compare.go:13
	_go_fuzz_dep_.CoverTab[902]++

//line /snap/go/10455/src/strings/compare.go:21
	if a == b {
//line /snap/go/10455/src/strings/compare.go:21
		_go_fuzz_dep_.CoverTab[524939]++
//line /snap/go/10455/src/strings/compare.go:21
		_go_fuzz_dep_.CoverTab[905]++
								return 0
//line /snap/go/10455/src/strings/compare.go:22
		// _ = "end of CoverTab[905]"
	} else {
//line /snap/go/10455/src/strings/compare.go:23
		_go_fuzz_dep_.CoverTab[524940]++
//line /snap/go/10455/src/strings/compare.go:23
		_go_fuzz_dep_.CoverTab[906]++
//line /snap/go/10455/src/strings/compare.go:23
		// _ = "end of CoverTab[906]"
//line /snap/go/10455/src/strings/compare.go:23
	}
//line /snap/go/10455/src/strings/compare.go:23
	// _ = "end of CoverTab[902]"
//line /snap/go/10455/src/strings/compare.go:23
	_go_fuzz_dep_.CoverTab[903]++
							if a < b {
//line /snap/go/10455/src/strings/compare.go:24
		_go_fuzz_dep_.CoverTab[524941]++
//line /snap/go/10455/src/strings/compare.go:24
		_go_fuzz_dep_.CoverTab[907]++
								return -1
//line /snap/go/10455/src/strings/compare.go:25
		// _ = "end of CoverTab[907]"
	} else {
//line /snap/go/10455/src/strings/compare.go:26
		_go_fuzz_dep_.CoverTab[524942]++
//line /snap/go/10455/src/strings/compare.go:26
		_go_fuzz_dep_.CoverTab[908]++
//line /snap/go/10455/src/strings/compare.go:26
		// _ = "end of CoverTab[908]"
//line /snap/go/10455/src/strings/compare.go:26
	}
//line /snap/go/10455/src/strings/compare.go:26
	// _ = "end of CoverTab[903]"
//line /snap/go/10455/src/strings/compare.go:26
	_go_fuzz_dep_.CoverTab[904]++
							return +1
//line /snap/go/10455/src/strings/compare.go:27
	// _ = "end of CoverTab[904]"
}

//line /snap/go/10455/src/strings/compare.go:28
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /snap/go/10455/src/strings/compare.go:28
var _ = _go_fuzz_dep_.CoverTab
