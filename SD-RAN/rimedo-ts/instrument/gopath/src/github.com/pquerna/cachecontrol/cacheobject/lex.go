// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:5
package cacheobject

//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:5
)

//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:9
func isSeparator(c byte) bool {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:9
	_go_fuzz_dep_.CoverTab[183836]++
																switch c {
	case '(', ')', '<', '>', '@', ',', ';', ':', '\\', '"', '/', '[', ']', '?', '=', '{', '}', ' ', '\t':
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:11
		_go_fuzz_dep_.CoverTab[183838]++
																	return true
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:12
		// _ = "end of CoverTab[183838]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:12
	default:
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:12
		_go_fuzz_dep_.CoverTab[183839]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:12
		// _ = "end of CoverTab[183839]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:13
	// _ = "end of CoverTab[183836]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:13
	_go_fuzz_dep_.CoverTab[183837]++
																return false
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:14
	// _ = "end of CoverTab[183837]"
}

func isCtl(c byte) bool {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:17
	_go_fuzz_dep_.CoverTab[183840]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:17
	return (0 <= c && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:17
		_go_fuzz_dep_.CoverTab[183841]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:17
		return c <= 31
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:17
		// _ = "end of CoverTab[183841]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:17
	}()) || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:17
		_go_fuzz_dep_.CoverTab[183842]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:17
		return c == 127
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:17
		// _ = "end of CoverTab[183842]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:17
	}()
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:17
	// _ = "end of CoverTab[183840]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:17
}

func isChar(c byte) bool {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:19
	_go_fuzz_dep_.CoverTab[183843]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:19
	return 0 <= c && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:19
		_go_fuzz_dep_.CoverTab[183844]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:19
		return c <= 127
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:19
		// _ = "end of CoverTab[183844]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:19
	}()
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:19
	// _ = "end of CoverTab[183843]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:19
}

func isAnyText(c byte) bool {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:21
	_go_fuzz_dep_.CoverTab[183845]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:21
	return !isCtl(c)
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:21
	// _ = "end of CoverTab[183845]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:21
}

func isQdText(c byte) bool {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:23
	_go_fuzz_dep_.CoverTab[183846]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:23
	return isAnyText(c) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:23
		_go_fuzz_dep_.CoverTab[183847]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:23
		return c != '"'
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:23
		// _ = "end of CoverTab[183847]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:23
	}()
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:23
	// _ = "end of CoverTab[183846]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:23
}

func isToken(c byte) bool {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:25
	_go_fuzz_dep_.CoverTab[183848]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:25
	return isChar(c) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:25
		_go_fuzz_dep_.CoverTab[183849]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:25
		return !isCtl(c)
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:25
		// _ = "end of CoverTab[183849]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:25
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:25
		_go_fuzz_dep_.CoverTab[183850]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:25
		return !isSeparator(c)
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:25
		// _ = "end of CoverTab[183850]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:25
	}()
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:25
	// _ = "end of CoverTab[183848]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:25
}

// Valid escaped sequences are not specified in RFC 2616, so for now, we assume
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:27
// that they coincide with the common sense ones used by GO. Malformed
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:27
// characters should probably not be treated as errors by a robust (forgiving)
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:27
// parser, so we replace them with the '?' character.
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:31
func httpUnquotePair(b byte) byte {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:31
	_go_fuzz_dep_.CoverTab[183851]++

																switch b {
	case 'a':
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:34
		_go_fuzz_dep_.CoverTab[183853]++
																	return '\a'
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:35
		// _ = "end of CoverTab[183853]"
	case 'b':
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:36
		_go_fuzz_dep_.CoverTab[183854]++
																	return '\b'
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:37
		// _ = "end of CoverTab[183854]"
	case 'f':
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:38
		_go_fuzz_dep_.CoverTab[183855]++
																	return '\f'
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:39
		// _ = "end of CoverTab[183855]"
	case 'n':
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:40
		_go_fuzz_dep_.CoverTab[183856]++
																	return '\n'
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:41
		// _ = "end of CoverTab[183856]"
	case 'r':
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:42
		_go_fuzz_dep_.CoverTab[183857]++
																	return '\r'
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:43
		// _ = "end of CoverTab[183857]"
	case 't':
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:44
		_go_fuzz_dep_.CoverTab[183858]++
																	return '\t'
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:45
		// _ = "end of CoverTab[183858]"
	case 'v':
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:46
		_go_fuzz_dep_.CoverTab[183859]++
																	return '\v'
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:47
		// _ = "end of CoverTab[183859]"
	case '\\':
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:48
		_go_fuzz_dep_.CoverTab[183860]++
																	return '\\'
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:49
		// _ = "end of CoverTab[183860]"
	case '\'':
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:50
		_go_fuzz_dep_.CoverTab[183861]++
																	return '\''
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:51
		// _ = "end of CoverTab[183861]"
	case '"':
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:52
		_go_fuzz_dep_.CoverTab[183862]++
																	return '"'
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:53
		// _ = "end of CoverTab[183862]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:53
	default:
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:53
		_go_fuzz_dep_.CoverTab[183863]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:53
		// _ = "end of CoverTab[183863]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:54
	// _ = "end of CoverTab[183851]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:54
	_go_fuzz_dep_.CoverTab[183852]++
																return '?'
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:55
	// _ = "end of CoverTab[183852]"
}

// raw must begin with a valid quoted string. Only the first quoted string is
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:58
// parsed and is unquoted in result. eaten is the number of bytes parsed, or -1
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:58
// upon failure.
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:61
func httpUnquote(raw string) (eaten int, result string) {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:61
	_go_fuzz_dep_.CoverTab[183864]++
																buf := make([]byte, len(raw))
																if raw[0] != '"' {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:63
		_go_fuzz_dep_.CoverTab[183867]++
																	return -1, ""
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:64
		// _ = "end of CoverTab[183867]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:65
		_go_fuzz_dep_.CoverTab[183868]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:65
		// _ = "end of CoverTab[183868]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:65
	}
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:65
	// _ = "end of CoverTab[183864]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:65
	_go_fuzz_dep_.CoverTab[183865]++
																eaten = 1
																j := 0
																for i := 1; i < len(raw); i++ {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:68
		_go_fuzz_dep_.CoverTab[183869]++
																	switch b := raw[i]; b {
		case '"':
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:70
			_go_fuzz_dep_.CoverTab[183870]++
																		eaten++
																		buf = buf[0:j]
																		return i + 1, string(buf)
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:73
			// _ = "end of CoverTab[183870]"
		case '\\':
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:74
			_go_fuzz_dep_.CoverTab[183871]++
																		if len(raw) < i+2 {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:75
				_go_fuzz_dep_.CoverTab[183875]++
																			return -1, ""
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:76
				// _ = "end of CoverTab[183875]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:77
				_go_fuzz_dep_.CoverTab[183876]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:77
				// _ = "end of CoverTab[183876]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:77
			}
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:77
			// _ = "end of CoverTab[183871]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:77
			_go_fuzz_dep_.CoverTab[183872]++
																		buf[j] = httpUnquotePair(raw[i+1])
																		eaten += 2
																		j++
																		i++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:81
			// _ = "end of CoverTab[183872]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:82
			_go_fuzz_dep_.CoverTab[183873]++
																		if isQdText(b) {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:83
				_go_fuzz_dep_.CoverTab[183877]++
																			buf[j] = b
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:84
				// _ = "end of CoverTab[183877]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:85
				_go_fuzz_dep_.CoverTab[183878]++
																			buf[j] = '?'
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:86
				// _ = "end of CoverTab[183878]"
			}
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:87
			// _ = "end of CoverTab[183873]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:87
			_go_fuzz_dep_.CoverTab[183874]++
																		eaten++
																		j++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:89
			// _ = "end of CoverTab[183874]"
		}
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:90
		// _ = "end of CoverTab[183869]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:91
	// _ = "end of CoverTab[183865]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:91
	_go_fuzz_dep_.CoverTab[183866]++
																return -1, ""
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:92
	// _ = "end of CoverTab[183866]"
}

//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:93
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/lex.go:93
var _ = _go_fuzz_dep_.CoverTab
