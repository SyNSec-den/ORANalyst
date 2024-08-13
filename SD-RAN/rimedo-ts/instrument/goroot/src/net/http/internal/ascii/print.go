// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/net/http/internal/ascii/print.go:5
package ascii

//line /usr/local/go/src/net/http/internal/ascii/print.go:5
import (
//line /usr/local/go/src/net/http/internal/ascii/print.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/http/internal/ascii/print.go:5
)
//line /usr/local/go/src/net/http/internal/ascii/print.go:5
import (
//line /usr/local/go/src/net/http/internal/ascii/print.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/http/internal/ascii/print.go:5
)

import (
	"strings"
	"unicode"
)

// EqualFold is strings.EqualFold, ASCII only. It reports whether s and t
//line /usr/local/go/src/net/http/internal/ascii/print.go:12
// are equal, ASCII-case-insensitively.
//line /usr/local/go/src/net/http/internal/ascii/print.go:14
func EqualFold(s, t string) bool {
//line /usr/local/go/src/net/http/internal/ascii/print.go:14
	_go_fuzz_dep_.CoverTab[36571]++
								if len(s) != len(t) {
//line /usr/local/go/src/net/http/internal/ascii/print.go:15
		_go_fuzz_dep_.CoverTab[36574]++
									return false
//line /usr/local/go/src/net/http/internal/ascii/print.go:16
		// _ = "end of CoverTab[36574]"
	} else {
//line /usr/local/go/src/net/http/internal/ascii/print.go:17
		_go_fuzz_dep_.CoverTab[36575]++
//line /usr/local/go/src/net/http/internal/ascii/print.go:17
		// _ = "end of CoverTab[36575]"
//line /usr/local/go/src/net/http/internal/ascii/print.go:17
	}
//line /usr/local/go/src/net/http/internal/ascii/print.go:17
	// _ = "end of CoverTab[36571]"
//line /usr/local/go/src/net/http/internal/ascii/print.go:17
	_go_fuzz_dep_.CoverTab[36572]++
								for i := 0; i < len(s); i++ {
//line /usr/local/go/src/net/http/internal/ascii/print.go:18
		_go_fuzz_dep_.CoverTab[36576]++
									if lower(s[i]) != lower(t[i]) {
//line /usr/local/go/src/net/http/internal/ascii/print.go:19
			_go_fuzz_dep_.CoverTab[36577]++
										return false
//line /usr/local/go/src/net/http/internal/ascii/print.go:20
			// _ = "end of CoverTab[36577]"
		} else {
//line /usr/local/go/src/net/http/internal/ascii/print.go:21
			_go_fuzz_dep_.CoverTab[36578]++
//line /usr/local/go/src/net/http/internal/ascii/print.go:21
			// _ = "end of CoverTab[36578]"
//line /usr/local/go/src/net/http/internal/ascii/print.go:21
		}
//line /usr/local/go/src/net/http/internal/ascii/print.go:21
		// _ = "end of CoverTab[36576]"
	}
//line /usr/local/go/src/net/http/internal/ascii/print.go:22
	// _ = "end of CoverTab[36572]"
//line /usr/local/go/src/net/http/internal/ascii/print.go:22
	_go_fuzz_dep_.CoverTab[36573]++
								return true
//line /usr/local/go/src/net/http/internal/ascii/print.go:23
	// _ = "end of CoverTab[36573]"
}

// lower returns the ASCII lowercase version of b.
func lower(b byte) byte {
//line /usr/local/go/src/net/http/internal/ascii/print.go:27
	_go_fuzz_dep_.CoverTab[36579]++
								if 'A' <= b && func() bool {
//line /usr/local/go/src/net/http/internal/ascii/print.go:28
		_go_fuzz_dep_.CoverTab[36581]++
//line /usr/local/go/src/net/http/internal/ascii/print.go:28
		return b <= 'Z'
//line /usr/local/go/src/net/http/internal/ascii/print.go:28
		// _ = "end of CoverTab[36581]"
//line /usr/local/go/src/net/http/internal/ascii/print.go:28
	}() {
//line /usr/local/go/src/net/http/internal/ascii/print.go:28
		_go_fuzz_dep_.CoverTab[36582]++
									return b + ('a' - 'A')
//line /usr/local/go/src/net/http/internal/ascii/print.go:29
		// _ = "end of CoverTab[36582]"
	} else {
//line /usr/local/go/src/net/http/internal/ascii/print.go:30
		_go_fuzz_dep_.CoverTab[36583]++
//line /usr/local/go/src/net/http/internal/ascii/print.go:30
		// _ = "end of CoverTab[36583]"
//line /usr/local/go/src/net/http/internal/ascii/print.go:30
	}
//line /usr/local/go/src/net/http/internal/ascii/print.go:30
	// _ = "end of CoverTab[36579]"
//line /usr/local/go/src/net/http/internal/ascii/print.go:30
	_go_fuzz_dep_.CoverTab[36580]++
								return b
//line /usr/local/go/src/net/http/internal/ascii/print.go:31
	// _ = "end of CoverTab[36580]"
}

// IsPrint returns whether s is ASCII and printable according to
//line /usr/local/go/src/net/http/internal/ascii/print.go:34
// https://tools.ietf.org/html/rfc20#section-4.2.
//line /usr/local/go/src/net/http/internal/ascii/print.go:36
func IsPrint(s string) bool {
//line /usr/local/go/src/net/http/internal/ascii/print.go:36
	_go_fuzz_dep_.CoverTab[36584]++
								for i := 0; i < len(s); i++ {
//line /usr/local/go/src/net/http/internal/ascii/print.go:37
		_go_fuzz_dep_.CoverTab[36586]++
									if s[i] < ' ' || func() bool {
//line /usr/local/go/src/net/http/internal/ascii/print.go:38
			_go_fuzz_dep_.CoverTab[36587]++
//line /usr/local/go/src/net/http/internal/ascii/print.go:38
			return s[i] > '~'
//line /usr/local/go/src/net/http/internal/ascii/print.go:38
			// _ = "end of CoverTab[36587]"
//line /usr/local/go/src/net/http/internal/ascii/print.go:38
		}() {
//line /usr/local/go/src/net/http/internal/ascii/print.go:38
			_go_fuzz_dep_.CoverTab[36588]++
										return false
//line /usr/local/go/src/net/http/internal/ascii/print.go:39
			// _ = "end of CoverTab[36588]"
		} else {
//line /usr/local/go/src/net/http/internal/ascii/print.go:40
			_go_fuzz_dep_.CoverTab[36589]++
//line /usr/local/go/src/net/http/internal/ascii/print.go:40
			// _ = "end of CoverTab[36589]"
//line /usr/local/go/src/net/http/internal/ascii/print.go:40
		}
//line /usr/local/go/src/net/http/internal/ascii/print.go:40
		// _ = "end of CoverTab[36586]"
	}
//line /usr/local/go/src/net/http/internal/ascii/print.go:41
	// _ = "end of CoverTab[36584]"
//line /usr/local/go/src/net/http/internal/ascii/print.go:41
	_go_fuzz_dep_.CoverTab[36585]++
								return true
//line /usr/local/go/src/net/http/internal/ascii/print.go:42
	// _ = "end of CoverTab[36585]"
}

// Is returns whether s is ASCII.
func Is(s string) bool {
//line /usr/local/go/src/net/http/internal/ascii/print.go:46
	_go_fuzz_dep_.CoverTab[36590]++
								for i := 0; i < len(s); i++ {
//line /usr/local/go/src/net/http/internal/ascii/print.go:47
		_go_fuzz_dep_.CoverTab[36592]++
									if s[i] > unicode.MaxASCII {
//line /usr/local/go/src/net/http/internal/ascii/print.go:48
			_go_fuzz_dep_.CoverTab[36593]++
										return false
//line /usr/local/go/src/net/http/internal/ascii/print.go:49
			// _ = "end of CoverTab[36593]"
		} else {
//line /usr/local/go/src/net/http/internal/ascii/print.go:50
			_go_fuzz_dep_.CoverTab[36594]++
//line /usr/local/go/src/net/http/internal/ascii/print.go:50
			// _ = "end of CoverTab[36594]"
//line /usr/local/go/src/net/http/internal/ascii/print.go:50
		}
//line /usr/local/go/src/net/http/internal/ascii/print.go:50
		// _ = "end of CoverTab[36592]"
	}
//line /usr/local/go/src/net/http/internal/ascii/print.go:51
	// _ = "end of CoverTab[36590]"
//line /usr/local/go/src/net/http/internal/ascii/print.go:51
	_go_fuzz_dep_.CoverTab[36591]++
								return true
//line /usr/local/go/src/net/http/internal/ascii/print.go:52
	// _ = "end of CoverTab[36591]"
}

// ToLower returns the lowercase version of s if s is ASCII and printable.
func ToLower(s string) (lower string, ok bool) {
//line /usr/local/go/src/net/http/internal/ascii/print.go:56
	_go_fuzz_dep_.CoverTab[36595]++
								if !IsPrint(s) {
//line /usr/local/go/src/net/http/internal/ascii/print.go:57
		_go_fuzz_dep_.CoverTab[36597]++
									return "", false
//line /usr/local/go/src/net/http/internal/ascii/print.go:58
		// _ = "end of CoverTab[36597]"
	} else {
//line /usr/local/go/src/net/http/internal/ascii/print.go:59
		_go_fuzz_dep_.CoverTab[36598]++
//line /usr/local/go/src/net/http/internal/ascii/print.go:59
		// _ = "end of CoverTab[36598]"
//line /usr/local/go/src/net/http/internal/ascii/print.go:59
	}
//line /usr/local/go/src/net/http/internal/ascii/print.go:59
	// _ = "end of CoverTab[36595]"
//line /usr/local/go/src/net/http/internal/ascii/print.go:59
	_go_fuzz_dep_.CoverTab[36596]++
								return strings.ToLower(s), true
//line /usr/local/go/src/net/http/internal/ascii/print.go:60
	// _ = "end of CoverTab[36596]"
}

//line /usr/local/go/src/net/http/internal/ascii/print.go:61
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/http/internal/ascii/print.go:61
var _ = _go_fuzz_dep_.CoverTab
