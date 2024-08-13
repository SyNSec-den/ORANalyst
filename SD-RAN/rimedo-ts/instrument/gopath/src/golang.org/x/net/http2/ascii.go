// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/ascii.go:5
package http2

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/ascii.go:5
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/ascii.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/ascii.go:5
)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/ascii.go:5
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/ascii.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/ascii.go:5
)

import "strings"

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/ascii.go:13
// asciiEqualFold is strings.EqualFold, ASCII only. It reports whether s and t
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/ascii.go:13
// are equal, ASCII-case-insensitively.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/ascii.go:15
func asciiEqualFold(s, t string) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/ascii.go:15
	_go_fuzz_dep_.CoverTab[72215]++
										if len(s) != len(t) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/ascii.go:16
		_go_fuzz_dep_.CoverTab[72218]++
											return false
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/ascii.go:17
		// _ = "end of CoverTab[72218]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/ascii.go:18
		_go_fuzz_dep_.CoverTab[72219]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/ascii.go:18
		// _ = "end of CoverTab[72219]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/ascii.go:18
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/ascii.go:18
	// _ = "end of CoverTab[72215]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/ascii.go:18
	_go_fuzz_dep_.CoverTab[72216]++
										for i := 0; i < len(s); i++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/ascii.go:19
		_go_fuzz_dep_.CoverTab[72220]++
											if lower(s[i]) != lower(t[i]) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/ascii.go:20
			_go_fuzz_dep_.CoverTab[72221]++
												return false
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/ascii.go:21
			// _ = "end of CoverTab[72221]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/ascii.go:22
			_go_fuzz_dep_.CoverTab[72222]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/ascii.go:22
			// _ = "end of CoverTab[72222]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/ascii.go:22
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/ascii.go:22
		// _ = "end of CoverTab[72220]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/ascii.go:23
	// _ = "end of CoverTab[72216]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/ascii.go:23
	_go_fuzz_dep_.CoverTab[72217]++
										return true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/ascii.go:24
	// _ = "end of CoverTab[72217]"
}

// lower returns the ASCII lowercase version of b.
func lower(b byte) byte {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/ascii.go:28
	_go_fuzz_dep_.CoverTab[72223]++
										if 'A' <= b && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/ascii.go:29
		_go_fuzz_dep_.CoverTab[72225]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/ascii.go:29
		return b <= 'Z'
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/ascii.go:29
		// _ = "end of CoverTab[72225]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/ascii.go:29
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/ascii.go:29
		_go_fuzz_dep_.CoverTab[72226]++
											return b + ('a' - 'A')
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/ascii.go:30
		// _ = "end of CoverTab[72226]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/ascii.go:31
		_go_fuzz_dep_.CoverTab[72227]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/ascii.go:31
		// _ = "end of CoverTab[72227]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/ascii.go:31
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/ascii.go:31
	// _ = "end of CoverTab[72223]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/ascii.go:31
	_go_fuzz_dep_.CoverTab[72224]++
										return b
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/ascii.go:32
	// _ = "end of CoverTab[72224]"
}

// isASCIIPrint returns whether s is ASCII and printable according to
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/ascii.go:35
// https://tools.ietf.org/html/rfc20#section-4.2.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/ascii.go:37
func isASCIIPrint(s string) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/ascii.go:37
	_go_fuzz_dep_.CoverTab[72228]++
										for i := 0; i < len(s); i++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/ascii.go:38
		_go_fuzz_dep_.CoverTab[72230]++
											if s[i] < ' ' || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/ascii.go:39
			_go_fuzz_dep_.CoverTab[72231]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/ascii.go:39
			return s[i] > '~'
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/ascii.go:39
			// _ = "end of CoverTab[72231]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/ascii.go:39
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/ascii.go:39
			_go_fuzz_dep_.CoverTab[72232]++
												return false
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/ascii.go:40
			// _ = "end of CoverTab[72232]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/ascii.go:41
			_go_fuzz_dep_.CoverTab[72233]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/ascii.go:41
			// _ = "end of CoverTab[72233]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/ascii.go:41
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/ascii.go:41
		// _ = "end of CoverTab[72230]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/ascii.go:42
	// _ = "end of CoverTab[72228]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/ascii.go:42
	_go_fuzz_dep_.CoverTab[72229]++
										return true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/ascii.go:43
	// _ = "end of CoverTab[72229]"
}

// asciiToLower returns the lowercase version of s if s is ASCII and printable,
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/ascii.go:46
// and whether or not it was.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/ascii.go:48
func asciiToLower(s string) (lower string, ok bool) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/ascii.go:48
	_go_fuzz_dep_.CoverTab[72234]++
										if !isASCIIPrint(s) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/ascii.go:49
		_go_fuzz_dep_.CoverTab[72236]++
											return "", false
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/ascii.go:50
		// _ = "end of CoverTab[72236]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/ascii.go:51
		_go_fuzz_dep_.CoverTab[72237]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/ascii.go:51
		// _ = "end of CoverTab[72237]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/ascii.go:51
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/ascii.go:51
	// _ = "end of CoverTab[72234]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/ascii.go:51
	_go_fuzz_dep_.CoverTab[72235]++
										return strings.ToLower(s), true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/ascii.go:52
	// _ = "end of CoverTab[72235]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/ascii.go:53
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/ascii.go:53
var _ = _go_fuzz_dep_.CoverTab
