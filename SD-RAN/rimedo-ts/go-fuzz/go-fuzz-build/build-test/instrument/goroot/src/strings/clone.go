// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/strings/clone.go:5
package strings

//line /usr/local/go/src/strings/clone.go:5
import (
//line /usr/local/go/src/strings/clone.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/strings/clone.go:5
)
//line /usr/local/go/src/strings/clone.go:5
import (
//line /usr/local/go/src/strings/clone.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/strings/clone.go:5
)

import (
	"unsafe"
)

// Clone returns a fresh copy of s.
//line /usr/local/go/src/strings/clone.go:11
// It guarantees to make a copy of s into a new allocation,
//line /usr/local/go/src/strings/clone.go:11
// which can be important when retaining only a small substring
//line /usr/local/go/src/strings/clone.go:11
// of a much larger string. Using Clone can help such programs
//line /usr/local/go/src/strings/clone.go:11
// use less memory. Of course, since using Clone makes a copy,
//line /usr/local/go/src/strings/clone.go:11
// overuse of Clone can make programs use more memory.
//line /usr/local/go/src/strings/clone.go:11
// Clone should typically be used only rarely, and only when
//line /usr/local/go/src/strings/clone.go:11
// profiling indicates that it is needed.
//line /usr/local/go/src/strings/clone.go:11
// For strings of length zero the string "" will be returned
//line /usr/local/go/src/strings/clone.go:11
// and no allocation is made.
//line /usr/local/go/src/strings/clone.go:21
func Clone(s string) string {
//line /usr/local/go/src/strings/clone.go:21
	_go_fuzz_dep_.CoverTab[885]++
						if len(s) == 0 {
//line /usr/local/go/src/strings/clone.go:22
		_go_fuzz_dep_.CoverTab[887]++
							return ""
//line /usr/local/go/src/strings/clone.go:23
		// _ = "end of CoverTab[887]"
	} else {
//line /usr/local/go/src/strings/clone.go:24
		_go_fuzz_dep_.CoverTab[888]++
//line /usr/local/go/src/strings/clone.go:24
		// _ = "end of CoverTab[888]"
//line /usr/local/go/src/strings/clone.go:24
	}
//line /usr/local/go/src/strings/clone.go:24
	// _ = "end of CoverTab[885]"
//line /usr/local/go/src/strings/clone.go:24
	_go_fuzz_dep_.CoverTab[886]++
						b := make([]byte, len(s))
						copy(b, s)
						return unsafe.String(&b[0], len(b))
//line /usr/local/go/src/strings/clone.go:27
	// _ = "end of CoverTab[886]"
}

//line /usr/local/go/src/strings/clone.go:28
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/strings/clone.go:28
var _ = _go_fuzz_dep_.CoverTab
