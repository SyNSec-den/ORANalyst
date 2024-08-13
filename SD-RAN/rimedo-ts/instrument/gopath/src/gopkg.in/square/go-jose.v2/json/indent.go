// Copyright 2010 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:5
package json

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:5
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:5
)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:5
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:5
)

import "bytes"

// Compact appends to dst the JSON-encoded src with
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:9
// insignificant space characters elided.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:11
func Compact(dst *bytes.Buffer, src []byte) error {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:11
	_go_fuzz_dep_.CoverTab[188393]++
											return compact(dst, src, false)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:12
	// _ = "end of CoverTab[188393]"
}

func compact(dst *bytes.Buffer, src []byte, escape bool) error {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:15
	_go_fuzz_dep_.CoverTab[188394]++
											origLen := dst.Len()
											var scan scanner
											scan.reset()
											start := 0
											for i, c := range src {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:20
		_go_fuzz_dep_.CoverTab[188398]++
												if escape && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:21
			_go_fuzz_dep_.CoverTab[188401]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:21
			return (c == '<' || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:21
				_go_fuzz_dep_.CoverTab[188402]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:21
				return c == '>'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:21
				// _ = "end of CoverTab[188402]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:21
			}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:21
				_go_fuzz_dep_.CoverTab[188403]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:21
				return c == '&'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:21
				// _ = "end of CoverTab[188403]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:21
			}())
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:21
			// _ = "end of CoverTab[188401]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:21
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:21
			_go_fuzz_dep_.CoverTab[188404]++
													if start < i {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:22
				_go_fuzz_dep_.CoverTab[188406]++
														dst.Write(src[start:i])
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:23
				// _ = "end of CoverTab[188406]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:24
				_go_fuzz_dep_.CoverTab[188407]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:24
				// _ = "end of CoverTab[188407]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:24
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:24
			// _ = "end of CoverTab[188404]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:24
			_go_fuzz_dep_.CoverTab[188405]++
													dst.WriteString(`\u00`)
													dst.WriteByte(hex[c>>4])
													dst.WriteByte(hex[c&0xF])
													start = i + 1
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:28
			// _ = "end of CoverTab[188405]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:29
			_go_fuzz_dep_.CoverTab[188408]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:29
			// _ = "end of CoverTab[188408]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:29
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:29
		// _ = "end of CoverTab[188398]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:29
		_go_fuzz_dep_.CoverTab[188399]++

												if c == 0xE2 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:31
			_go_fuzz_dep_.CoverTab[188409]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:31
			return i+2 < len(src)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:31
			// _ = "end of CoverTab[188409]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:31
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:31
			_go_fuzz_dep_.CoverTab[188410]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:31
			return src[i+1] == 0x80
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:31
			// _ = "end of CoverTab[188410]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:31
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:31
			_go_fuzz_dep_.CoverTab[188411]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:31
			return src[i+2]&^1 == 0xA8
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:31
			// _ = "end of CoverTab[188411]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:31
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:31
			_go_fuzz_dep_.CoverTab[188412]++
													if start < i {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:32
				_go_fuzz_dep_.CoverTab[188414]++
														dst.Write(src[start:i])
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:33
				// _ = "end of CoverTab[188414]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:34
				_go_fuzz_dep_.CoverTab[188415]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:34
				// _ = "end of CoverTab[188415]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:34
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:34
			// _ = "end of CoverTab[188412]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:34
			_go_fuzz_dep_.CoverTab[188413]++
													dst.WriteString(`\u202`)
													dst.WriteByte(hex[src[i+2]&0xF])
													start = i + 3
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:37
			// _ = "end of CoverTab[188413]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:38
			_go_fuzz_dep_.CoverTab[188416]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:38
			// _ = "end of CoverTab[188416]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:38
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:38
		// _ = "end of CoverTab[188399]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:38
		_go_fuzz_dep_.CoverTab[188400]++
												v := scan.step(&scan, c)
												if v >= scanSkipSpace {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:40
			_go_fuzz_dep_.CoverTab[188417]++
													if v == scanError {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:41
				_go_fuzz_dep_.CoverTab[188420]++
														break
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:42
				// _ = "end of CoverTab[188420]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:43
				_go_fuzz_dep_.CoverTab[188421]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:43
				// _ = "end of CoverTab[188421]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:43
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:43
			// _ = "end of CoverTab[188417]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:43
			_go_fuzz_dep_.CoverTab[188418]++
													if start < i {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:44
				_go_fuzz_dep_.CoverTab[188422]++
														dst.Write(src[start:i])
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:45
				// _ = "end of CoverTab[188422]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:46
				_go_fuzz_dep_.CoverTab[188423]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:46
				// _ = "end of CoverTab[188423]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:46
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:46
			// _ = "end of CoverTab[188418]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:46
			_go_fuzz_dep_.CoverTab[188419]++
													start = i + 1
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:47
			// _ = "end of CoverTab[188419]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:48
			_go_fuzz_dep_.CoverTab[188424]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:48
			// _ = "end of CoverTab[188424]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:48
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:48
		// _ = "end of CoverTab[188400]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:49
	// _ = "end of CoverTab[188394]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:49
	_go_fuzz_dep_.CoverTab[188395]++
											if scan.eof() == scanError {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:50
		_go_fuzz_dep_.CoverTab[188425]++
												dst.Truncate(origLen)
												return scan.err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:52
		// _ = "end of CoverTab[188425]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:53
		_go_fuzz_dep_.CoverTab[188426]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:53
		// _ = "end of CoverTab[188426]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:53
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:53
	// _ = "end of CoverTab[188395]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:53
	_go_fuzz_dep_.CoverTab[188396]++
											if start < len(src) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:54
		_go_fuzz_dep_.CoverTab[188427]++
												dst.Write(src[start:])
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:55
		// _ = "end of CoverTab[188427]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:56
		_go_fuzz_dep_.CoverTab[188428]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:56
		// _ = "end of CoverTab[188428]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:56
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:56
	// _ = "end of CoverTab[188396]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:56
	_go_fuzz_dep_.CoverTab[188397]++
											return nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:57
	// _ = "end of CoverTab[188397]"
}

func newline(dst *bytes.Buffer, prefix, indent string, depth int) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:60
	_go_fuzz_dep_.CoverTab[188429]++
											dst.WriteByte('\n')
											dst.WriteString(prefix)
											for i := 0; i < depth; i++ {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:63
		_go_fuzz_dep_.CoverTab[188430]++
												dst.WriteString(indent)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:64
		// _ = "end of CoverTab[188430]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:65
	// _ = "end of CoverTab[188429]"
}

// Indent appends to dst an indented form of the JSON-encoded src.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:68
// Each element in a JSON object or array begins on a new,
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:68
// indented line beginning with prefix followed by one or more
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:68
// copies of indent according to the indentation nesting.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:68
// The data appended to dst does not begin with the prefix nor
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:68
// any indentation, to make it easier to embed inside other formatted JSON data.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:68
// Although leading space characters (space, tab, carriage return, newline)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:68
// at the beginning of src are dropped, trailing space characters
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:68
// at the end of src are preserved and copied to dst.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:68
// For example, if src has no trailing spaces, neither will dst;
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:68
// if src ends in a trailing newline, so will dst.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:79
func Indent(dst *bytes.Buffer, src []byte, prefix, indent string) error {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:79
	_go_fuzz_dep_.CoverTab[188431]++
											origLen := dst.Len()
											var scan scanner
											scan.reset()
											needIndent := false
											depth := 0
											for _, c := range src {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:85
		_go_fuzz_dep_.CoverTab[188434]++
												scan.bytes++
												v := scan.step(&scan, c)
												if v == scanSkipSpace {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:88
			_go_fuzz_dep_.CoverTab[188439]++
													continue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:89
			// _ = "end of CoverTab[188439]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:90
			_go_fuzz_dep_.CoverTab[188440]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:90
			// _ = "end of CoverTab[188440]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:90
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:90
		// _ = "end of CoverTab[188434]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:90
		_go_fuzz_dep_.CoverTab[188435]++
												if v == scanError {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:91
			_go_fuzz_dep_.CoverTab[188441]++
													break
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:92
			// _ = "end of CoverTab[188441]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:93
			_go_fuzz_dep_.CoverTab[188442]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:93
			// _ = "end of CoverTab[188442]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:93
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:93
		// _ = "end of CoverTab[188435]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:93
		_go_fuzz_dep_.CoverTab[188436]++
												if needIndent && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:94
			_go_fuzz_dep_.CoverTab[188443]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:94
			return v != scanEndObject
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:94
			// _ = "end of CoverTab[188443]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:94
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:94
			_go_fuzz_dep_.CoverTab[188444]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:94
			return v != scanEndArray
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:94
			// _ = "end of CoverTab[188444]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:94
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:94
			_go_fuzz_dep_.CoverTab[188445]++
													needIndent = false
													depth++
													newline(dst, prefix, indent, depth)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:97
			// _ = "end of CoverTab[188445]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:98
			_go_fuzz_dep_.CoverTab[188446]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:98
			// _ = "end of CoverTab[188446]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:98
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:98
		// _ = "end of CoverTab[188436]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:98
		_go_fuzz_dep_.CoverTab[188437]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:102
		if v == scanContinue {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:102
			_go_fuzz_dep_.CoverTab[188447]++
													dst.WriteByte(c)
													continue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:104
			// _ = "end of CoverTab[188447]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:105
			_go_fuzz_dep_.CoverTab[188448]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:105
			// _ = "end of CoverTab[188448]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:105
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:105
		// _ = "end of CoverTab[188437]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:105
		_go_fuzz_dep_.CoverTab[188438]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:108
		switch c {
		case '{', '[':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:109
			_go_fuzz_dep_.CoverTab[188449]++

													needIndent = true
													dst.WriteByte(c)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:112
			// _ = "end of CoverTab[188449]"

		case ',':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:114
			_go_fuzz_dep_.CoverTab[188450]++
													dst.WriteByte(c)
													newline(dst, prefix, indent, depth)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:116
			// _ = "end of CoverTab[188450]"

		case ':':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:118
			_go_fuzz_dep_.CoverTab[188451]++
													dst.WriteByte(c)
													dst.WriteByte(' ')
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:120
			// _ = "end of CoverTab[188451]"

		case '}', ']':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:122
			_go_fuzz_dep_.CoverTab[188452]++
													if needIndent {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:123
				_go_fuzz_dep_.CoverTab[188455]++

														needIndent = false
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:125
				// _ = "end of CoverTab[188455]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:126
				_go_fuzz_dep_.CoverTab[188456]++
														depth--
														newline(dst, prefix, indent, depth)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:128
				// _ = "end of CoverTab[188456]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:129
			// _ = "end of CoverTab[188452]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:129
			_go_fuzz_dep_.CoverTab[188453]++
													dst.WriteByte(c)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:130
			// _ = "end of CoverTab[188453]"

		default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:132
			_go_fuzz_dep_.CoverTab[188454]++
													dst.WriteByte(c)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:133
			// _ = "end of CoverTab[188454]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:134
		// _ = "end of CoverTab[188438]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:135
	// _ = "end of CoverTab[188431]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:135
	_go_fuzz_dep_.CoverTab[188432]++
											if scan.eof() == scanError {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:136
		_go_fuzz_dep_.CoverTab[188457]++
												dst.Truncate(origLen)
												return scan.err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:138
		// _ = "end of CoverTab[188457]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:139
		_go_fuzz_dep_.CoverTab[188458]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:139
		// _ = "end of CoverTab[188458]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:139
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:139
	// _ = "end of CoverTab[188432]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:139
	_go_fuzz_dep_.CoverTab[188433]++
											return nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:140
	// _ = "end of CoverTab[188433]"
}

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:141
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/indent.go:141
var _ = _go_fuzz_dep_.CoverTab
