// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/mime/grammar.go:5
package mime

//line /usr/local/go/src/mime/grammar.go:5
import (
//line /usr/local/go/src/mime/grammar.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/mime/grammar.go:5
)
//line /usr/local/go/src/mime/grammar.go:5
import (
//line /usr/local/go/src/mime/grammar.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/mime/grammar.go:5
)

import (
	"strings"
)

// isTSpecial reports whether rune is in 'tspecials' as defined by RFC
//line /usr/local/go/src/mime/grammar.go:11
// 1521 and RFC 2045.
//line /usr/local/go/src/mime/grammar.go:13
func isTSpecial(r rune) bool {
//line /usr/local/go/src/mime/grammar.go:13
	_go_fuzz_dep_.CoverTab[35600]++
						return strings.ContainsRune(`()<>@,;:\"/[]?=`, r)
//line /usr/local/go/src/mime/grammar.go:14
	// _ = "end of CoverTab[35600]"
}

// isTokenChar reports whether rune is in 'token' as defined by RFC
//line /usr/local/go/src/mime/grammar.go:17
// 1521 and RFC 2045.
//line /usr/local/go/src/mime/grammar.go:19
func isTokenChar(r rune) bool {
//line /usr/local/go/src/mime/grammar.go:19
	_go_fuzz_dep_.CoverTab[35601]++

//line /usr/local/go/src/mime/grammar.go:22
	return r > 0x20 && func() bool {
//line /usr/local/go/src/mime/grammar.go:22
		_go_fuzz_dep_.CoverTab[35602]++
//line /usr/local/go/src/mime/grammar.go:22
		return r < 0x7f
//line /usr/local/go/src/mime/grammar.go:22
		// _ = "end of CoverTab[35602]"
//line /usr/local/go/src/mime/grammar.go:22
	}() && func() bool {
//line /usr/local/go/src/mime/grammar.go:22
		_go_fuzz_dep_.CoverTab[35603]++
//line /usr/local/go/src/mime/grammar.go:22
		return !isTSpecial(r)
//line /usr/local/go/src/mime/grammar.go:22
		// _ = "end of CoverTab[35603]"
//line /usr/local/go/src/mime/grammar.go:22
	}()
//line /usr/local/go/src/mime/grammar.go:22
	// _ = "end of CoverTab[35601]"
}

// isToken reports whether s is a 'token' as defined by RFC 1521
//line /usr/local/go/src/mime/grammar.go:25
// and RFC 2045.
//line /usr/local/go/src/mime/grammar.go:27
func isToken(s string) bool {
//line /usr/local/go/src/mime/grammar.go:27
	_go_fuzz_dep_.CoverTab[35604]++
						if s == "" {
//line /usr/local/go/src/mime/grammar.go:28
		_go_fuzz_dep_.CoverTab[35606]++
							return false
//line /usr/local/go/src/mime/grammar.go:29
		// _ = "end of CoverTab[35606]"
	} else {
//line /usr/local/go/src/mime/grammar.go:30
		_go_fuzz_dep_.CoverTab[35607]++
//line /usr/local/go/src/mime/grammar.go:30
		// _ = "end of CoverTab[35607]"
//line /usr/local/go/src/mime/grammar.go:30
	}
//line /usr/local/go/src/mime/grammar.go:30
	// _ = "end of CoverTab[35604]"
//line /usr/local/go/src/mime/grammar.go:30
	_go_fuzz_dep_.CoverTab[35605]++
						return strings.IndexFunc(s, isNotTokenChar) < 0
//line /usr/local/go/src/mime/grammar.go:31
	// _ = "end of CoverTab[35605]"
}

//line /usr/local/go/src/mime/grammar.go:32
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/mime/grammar.go:32
var _ = _go_fuzz_dep_.CoverTab
