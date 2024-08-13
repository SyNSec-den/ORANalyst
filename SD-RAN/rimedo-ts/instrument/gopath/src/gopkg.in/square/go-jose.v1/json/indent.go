// Copyright 2010 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:5
package json

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:5
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:5
)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:5
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:5
)

import "bytes"

// Compact appends to dst the JSON-encoded src with
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:9
// insignificant space characters elided.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:11
func Compact(dst *bytes.Buffer, src []byte) error {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:11
	_go_fuzz_dep_.CoverTab[185335]++
											return compact(dst, src, false)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:12
	// _ = "end of CoverTab[185335]"
}

func compact(dst *bytes.Buffer, src []byte, escape bool) error {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:15
	_go_fuzz_dep_.CoverTab[185336]++
											origLen := dst.Len()
											var scan scanner
											scan.reset()
											start := 0
											for i, c := range src {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:20
		_go_fuzz_dep_.CoverTab[185340]++
												if escape && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:21
			_go_fuzz_dep_.CoverTab[185343]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:21
			return (c == '<' || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:21
				_go_fuzz_dep_.CoverTab[185344]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:21
				return c == '>'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:21
				// _ = "end of CoverTab[185344]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:21
			}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:21
				_go_fuzz_dep_.CoverTab[185345]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:21
				return c == '&'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:21
				// _ = "end of CoverTab[185345]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:21
			}())
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:21
			// _ = "end of CoverTab[185343]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:21
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:21
			_go_fuzz_dep_.CoverTab[185346]++
													if start < i {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:22
				_go_fuzz_dep_.CoverTab[185348]++
														dst.Write(src[start:i])
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:23
				// _ = "end of CoverTab[185348]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:24
				_go_fuzz_dep_.CoverTab[185349]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:24
				// _ = "end of CoverTab[185349]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:24
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:24
			// _ = "end of CoverTab[185346]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:24
			_go_fuzz_dep_.CoverTab[185347]++
													dst.WriteString(`\u00`)
													dst.WriteByte(hex[c>>4])
													dst.WriteByte(hex[c&0xF])
													start = i + 1
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:28
			// _ = "end of CoverTab[185347]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:29
			_go_fuzz_dep_.CoverTab[185350]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:29
			// _ = "end of CoverTab[185350]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:29
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:29
		// _ = "end of CoverTab[185340]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:29
		_go_fuzz_dep_.CoverTab[185341]++

												if c == 0xE2 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:31
			_go_fuzz_dep_.CoverTab[185351]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:31
			return i+2 < len(src)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:31
			// _ = "end of CoverTab[185351]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:31
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:31
			_go_fuzz_dep_.CoverTab[185352]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:31
			return src[i+1] == 0x80
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:31
			// _ = "end of CoverTab[185352]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:31
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:31
			_go_fuzz_dep_.CoverTab[185353]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:31
			return src[i+2]&^1 == 0xA8
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:31
			// _ = "end of CoverTab[185353]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:31
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:31
			_go_fuzz_dep_.CoverTab[185354]++
													if start < i {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:32
				_go_fuzz_dep_.CoverTab[185356]++
														dst.Write(src[start:i])
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:33
				// _ = "end of CoverTab[185356]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:34
				_go_fuzz_dep_.CoverTab[185357]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:34
				// _ = "end of CoverTab[185357]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:34
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:34
			// _ = "end of CoverTab[185354]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:34
			_go_fuzz_dep_.CoverTab[185355]++
													dst.WriteString(`\u202`)
													dst.WriteByte(hex[src[i+2]&0xF])
													start = i + 3
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:37
			// _ = "end of CoverTab[185355]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:38
			_go_fuzz_dep_.CoverTab[185358]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:38
			// _ = "end of CoverTab[185358]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:38
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:38
		// _ = "end of CoverTab[185341]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:38
		_go_fuzz_dep_.CoverTab[185342]++
												v := scan.step(&scan, c)
												if v >= scanSkipSpace {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:40
			_go_fuzz_dep_.CoverTab[185359]++
													if v == scanError {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:41
				_go_fuzz_dep_.CoverTab[185362]++
														break
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:42
				// _ = "end of CoverTab[185362]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:43
				_go_fuzz_dep_.CoverTab[185363]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:43
				// _ = "end of CoverTab[185363]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:43
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:43
			// _ = "end of CoverTab[185359]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:43
			_go_fuzz_dep_.CoverTab[185360]++
													if start < i {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:44
				_go_fuzz_dep_.CoverTab[185364]++
														dst.Write(src[start:i])
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:45
				// _ = "end of CoverTab[185364]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:46
				_go_fuzz_dep_.CoverTab[185365]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:46
				// _ = "end of CoverTab[185365]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:46
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:46
			// _ = "end of CoverTab[185360]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:46
			_go_fuzz_dep_.CoverTab[185361]++
													start = i + 1
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:47
			// _ = "end of CoverTab[185361]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:48
			_go_fuzz_dep_.CoverTab[185366]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:48
			// _ = "end of CoverTab[185366]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:48
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:48
		// _ = "end of CoverTab[185342]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:49
	// _ = "end of CoverTab[185336]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:49
	_go_fuzz_dep_.CoverTab[185337]++
											if scan.eof() == scanError {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:50
		_go_fuzz_dep_.CoverTab[185367]++
												dst.Truncate(origLen)
												return scan.err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:52
		// _ = "end of CoverTab[185367]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:53
		_go_fuzz_dep_.CoverTab[185368]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:53
		// _ = "end of CoverTab[185368]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:53
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:53
	// _ = "end of CoverTab[185337]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:53
	_go_fuzz_dep_.CoverTab[185338]++
											if start < len(src) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:54
		_go_fuzz_dep_.CoverTab[185369]++
												dst.Write(src[start:])
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:55
		// _ = "end of CoverTab[185369]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:56
		_go_fuzz_dep_.CoverTab[185370]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:56
		// _ = "end of CoverTab[185370]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:56
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:56
	// _ = "end of CoverTab[185338]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:56
	_go_fuzz_dep_.CoverTab[185339]++
											return nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:57
	// _ = "end of CoverTab[185339]"
}

func newline(dst *bytes.Buffer, prefix, indent string, depth int) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:60
	_go_fuzz_dep_.CoverTab[185371]++
											dst.WriteByte('\n')
											dst.WriteString(prefix)
											for i := 0; i < depth; i++ {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:63
		_go_fuzz_dep_.CoverTab[185372]++
												dst.WriteString(indent)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:64
		// _ = "end of CoverTab[185372]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:65
	// _ = "end of CoverTab[185371]"
}

// Indent appends to dst an indented form of the JSON-encoded src.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:68
// Each element in a JSON object or array begins on a new,
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:68
// indented line beginning with prefix followed by one or more
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:68
// copies of indent according to the indentation nesting.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:68
// The data appended to dst does not begin with the prefix nor
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:68
// any indentation, to make it easier to embed inside other formatted JSON data.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:68
// Although leading space characters (space, tab, carriage return, newline)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:68
// at the beginning of src are dropped, trailing space characters
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:68
// at the end of src are preserved and copied to dst.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:68
// For example, if src has no trailing spaces, neither will dst;
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:68
// if src ends in a trailing newline, so will dst.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:79
func Indent(dst *bytes.Buffer, src []byte, prefix, indent string) error {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:79
	_go_fuzz_dep_.CoverTab[185373]++
											origLen := dst.Len()
											var scan scanner
											scan.reset()
											needIndent := false
											depth := 0
											for _, c := range src {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:85
		_go_fuzz_dep_.CoverTab[185376]++
												scan.bytes++
												v := scan.step(&scan, c)
												if v == scanSkipSpace {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:88
			_go_fuzz_dep_.CoverTab[185381]++
													continue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:89
			// _ = "end of CoverTab[185381]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:90
			_go_fuzz_dep_.CoverTab[185382]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:90
			// _ = "end of CoverTab[185382]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:90
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:90
		// _ = "end of CoverTab[185376]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:90
		_go_fuzz_dep_.CoverTab[185377]++
												if v == scanError {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:91
			_go_fuzz_dep_.CoverTab[185383]++
													break
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:92
			// _ = "end of CoverTab[185383]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:93
			_go_fuzz_dep_.CoverTab[185384]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:93
			// _ = "end of CoverTab[185384]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:93
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:93
		// _ = "end of CoverTab[185377]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:93
		_go_fuzz_dep_.CoverTab[185378]++
												if needIndent && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:94
			_go_fuzz_dep_.CoverTab[185385]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:94
			return v != scanEndObject
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:94
			// _ = "end of CoverTab[185385]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:94
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:94
			_go_fuzz_dep_.CoverTab[185386]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:94
			return v != scanEndArray
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:94
			// _ = "end of CoverTab[185386]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:94
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:94
			_go_fuzz_dep_.CoverTab[185387]++
													needIndent = false
													depth++
													newline(dst, prefix, indent, depth)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:97
			// _ = "end of CoverTab[185387]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:98
			_go_fuzz_dep_.CoverTab[185388]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:98
			// _ = "end of CoverTab[185388]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:98
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:98
		// _ = "end of CoverTab[185378]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:98
		_go_fuzz_dep_.CoverTab[185379]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:102
		if v == scanContinue {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:102
			_go_fuzz_dep_.CoverTab[185389]++
													dst.WriteByte(c)
													continue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:104
			// _ = "end of CoverTab[185389]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:105
			_go_fuzz_dep_.CoverTab[185390]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:105
			// _ = "end of CoverTab[185390]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:105
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:105
		// _ = "end of CoverTab[185379]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:105
		_go_fuzz_dep_.CoverTab[185380]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:108
		switch c {
		case '{', '[':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:109
			_go_fuzz_dep_.CoverTab[185391]++

													needIndent = true
													dst.WriteByte(c)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:112
			// _ = "end of CoverTab[185391]"

		case ',':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:114
			_go_fuzz_dep_.CoverTab[185392]++
													dst.WriteByte(c)
													newline(dst, prefix, indent, depth)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:116
			// _ = "end of CoverTab[185392]"

		case ':':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:118
			_go_fuzz_dep_.CoverTab[185393]++
													dst.WriteByte(c)
													dst.WriteByte(' ')
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:120
			// _ = "end of CoverTab[185393]"

		case '}', ']':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:122
			_go_fuzz_dep_.CoverTab[185394]++
													if needIndent {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:123
				_go_fuzz_dep_.CoverTab[185397]++

														needIndent = false
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:125
				// _ = "end of CoverTab[185397]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:126
				_go_fuzz_dep_.CoverTab[185398]++
														depth--
														newline(dst, prefix, indent, depth)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:128
				// _ = "end of CoverTab[185398]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:129
			// _ = "end of CoverTab[185394]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:129
			_go_fuzz_dep_.CoverTab[185395]++
													dst.WriteByte(c)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:130
			// _ = "end of CoverTab[185395]"

		default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:132
			_go_fuzz_dep_.CoverTab[185396]++
													dst.WriteByte(c)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:133
			// _ = "end of CoverTab[185396]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:134
		// _ = "end of CoverTab[185380]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:135
	// _ = "end of CoverTab[185373]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:135
	_go_fuzz_dep_.CoverTab[185374]++
											if scan.eof() == scanError {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:136
		_go_fuzz_dep_.CoverTab[185399]++
												dst.Truncate(origLen)
												return scan.err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:138
		// _ = "end of CoverTab[185399]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:139
		_go_fuzz_dep_.CoverTab[185400]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:139
		// _ = "end of CoverTab[185400]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:139
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:139
	// _ = "end of CoverTab[185374]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:139
	_go_fuzz_dep_.CoverTab[185375]++
											return nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:140
	// _ = "end of CoverTab[185375]"
}

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:141
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/indent.go:141
var _ = _go_fuzz_dep_.CoverTab
