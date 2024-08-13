// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/strings/compare.go:5
package strings

//line /usr/local/go/src/strings/compare.go:5
import (
//line /usr/local/go/src/strings/compare.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/strings/compare.go:5
)
//line /usr/local/go/src/strings/compare.go:5
import (
//line /usr/local/go/src/strings/compare.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/strings/compare.go:5
)

// Compare returns an integer comparing two strings lexicographically.
//line /usr/local/go/src/strings/compare.go:7
// The result will be 0 if a == b, -1 if a < b, and +1 if a > b.
//line /usr/local/go/src/strings/compare.go:7
//
//line /usr/local/go/src/strings/compare.go:7
// Compare is included only for symmetry with package bytes.
//line /usr/local/go/src/strings/compare.go:7
// It is usually clearer and always faster to use the built-in
//line /usr/local/go/src/strings/compare.go:7
// string comparison operators ==, <, >, and so on.
//line /usr/local/go/src/strings/compare.go:13
func Compare(a, b string) int {
//line /usr/local/go/src/strings/compare.go:13
	_go_fuzz_dep_.CoverTab[889]++

//line /usr/local/go/src/strings/compare.go:21
	if a == b {
//line /usr/local/go/src/strings/compare.go:21
		_go_fuzz_dep_.CoverTab[892]++
							return 0
//line /usr/local/go/src/strings/compare.go:22
		// _ = "end of CoverTab[892]"
	} else {
//line /usr/local/go/src/strings/compare.go:23
		_go_fuzz_dep_.CoverTab[893]++
//line /usr/local/go/src/strings/compare.go:23
		// _ = "end of CoverTab[893]"
//line /usr/local/go/src/strings/compare.go:23
	}
//line /usr/local/go/src/strings/compare.go:23
	// _ = "end of CoverTab[889]"
//line /usr/local/go/src/strings/compare.go:23
	_go_fuzz_dep_.CoverTab[890]++
						if a < b {
//line /usr/local/go/src/strings/compare.go:24
		_go_fuzz_dep_.CoverTab[894]++
							return -1
//line /usr/local/go/src/strings/compare.go:25
		// _ = "end of CoverTab[894]"
	} else {
//line /usr/local/go/src/strings/compare.go:26
		_go_fuzz_dep_.CoverTab[895]++
//line /usr/local/go/src/strings/compare.go:26
		// _ = "end of CoverTab[895]"
//line /usr/local/go/src/strings/compare.go:26
	}
//line /usr/local/go/src/strings/compare.go:26
	// _ = "end of CoverTab[890]"
//line /usr/local/go/src/strings/compare.go:26
	_go_fuzz_dep_.CoverTab[891]++
						return +1
//line /usr/local/go/src/strings/compare.go:27
	// _ = "end of CoverTab[891]"
}

//line /usr/local/go/src/strings/compare.go:28
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/strings/compare.go:28
var _ = _go_fuzz_dep_.CoverTab
