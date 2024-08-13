// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:5
// Package strs provides string manipulation functionality specific to protobuf.
package strs

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:6
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:6
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:6
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:6
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:6
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:6
)

import (
	"go/token"
	"strings"
	"unicode"
	"unicode/utf8"

	"google.golang.org/protobuf/internal/flags"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// EnforceUTF8 reports whether to enforce strict UTF-8 validation.
func EnforceUTF8(fd protoreflect.FieldDescriptor) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:19
	_go_fuzz_dep_.CoverTab[49304]++
													if flags.ProtoLegacy {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:20
		_go_fuzz_dep_.CoverTab[49306]++
														if fd, ok := fd.(interface{ EnforceUTF8() bool }); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:21
			_go_fuzz_dep_.CoverTab[49307]++
															return fd.EnforceUTF8()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:22
			// _ = "end of CoverTab[49307]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:23
			_go_fuzz_dep_.CoverTab[49308]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:23
			// _ = "end of CoverTab[49308]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:23
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:23
		// _ = "end of CoverTab[49306]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:24
		_go_fuzz_dep_.CoverTab[49309]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:24
		// _ = "end of CoverTab[49309]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:24
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:24
	// _ = "end of CoverTab[49304]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:24
	_go_fuzz_dep_.CoverTab[49305]++
													return fd.Syntax() == protoreflect.Proto3
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:25
	// _ = "end of CoverTab[49305]"
}

// GoCamelCase camel-cases a protobuf name for use as a Go identifier.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:28
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:28
// If there is an interior underscore followed by a lower case letter,
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:28
// drop the underscore and convert the letter to upper case.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:32
func GoCamelCase(s string) string {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:32
	_go_fuzz_dep_.CoverTab[49310]++
	// Invariant: if the next letter is lower case, it must be converted
	// to upper case.
	// That is, we process a word at a time, where words are marked by _ or
	// upper case letter. Digits are treated as words.
	var b []byte
	for i := 0; i < len(s); i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:38
		_go_fuzz_dep_.CoverTab[49312]++
														c := s[i]
														switch {
		case c == '.' && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:41
			_go_fuzz_dep_.CoverTab[49320]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:41
			return i+1 < len(s)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:41
			// _ = "end of CoverTab[49320]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:41
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:41
			_go_fuzz_dep_.CoverTab[49321]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:41
			return isASCIILower(s[i+1])
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:41
			// _ = "end of CoverTab[49321]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:41
		}():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:41
			_go_fuzz_dep_.CoverTab[49313]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:41
			// _ = "end of CoverTab[49313]"

		case c == '.':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:43
			_go_fuzz_dep_.CoverTab[49314]++
															b = append(b, '_')
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:44
			// _ = "end of CoverTab[49314]"
		case c == '_' && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:45
			_go_fuzz_dep_.CoverTab[49322]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:45
			return (i == 0 || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:45
				_go_fuzz_dep_.CoverTab[49323]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:45
				return s[i-1] == '.'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:45
				// _ = "end of CoverTab[49323]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:45
			}())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:45
			// _ = "end of CoverTab[49322]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:45
		}():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:45
			_go_fuzz_dep_.CoverTab[49315]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:48
			b = append(b, 'X')
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:48
			// _ = "end of CoverTab[49315]"
		case c == '_' && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:49
			_go_fuzz_dep_.CoverTab[49324]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:49
			return i+1 < len(s)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:49
			// _ = "end of CoverTab[49324]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:49
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:49
			_go_fuzz_dep_.CoverTab[49325]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:49
			return isASCIILower(s[i+1])
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:49
			// _ = "end of CoverTab[49325]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:49
		}():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:49
			_go_fuzz_dep_.CoverTab[49316]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:49
			// _ = "end of CoverTab[49316]"

		case isASCIIDigit(c):
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:51
			_go_fuzz_dep_.CoverTab[49317]++
															b = append(b, c)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:52
			// _ = "end of CoverTab[49317]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:53
			_go_fuzz_dep_.CoverTab[49318]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:56
			if isASCIILower(c) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:56
				_go_fuzz_dep_.CoverTab[49326]++
																c -= 'a' - 'A'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:57
				// _ = "end of CoverTab[49326]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:58
				_go_fuzz_dep_.CoverTab[49327]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:58
				// _ = "end of CoverTab[49327]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:58
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:58
			// _ = "end of CoverTab[49318]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:58
			_go_fuzz_dep_.CoverTab[49319]++
															b = append(b, c)

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:62
			for ; i+1 < len(s) && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:62
				_go_fuzz_dep_.CoverTab[49328]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:62
				return isASCIILower(s[i+1])
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:62
				// _ = "end of CoverTab[49328]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:62
			}(); i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:62
				_go_fuzz_dep_.CoverTab[49329]++
																b = append(b, s[i+1])
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:63
				// _ = "end of CoverTab[49329]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:64
			// _ = "end of CoverTab[49319]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:65
		// _ = "end of CoverTab[49312]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:66
	// _ = "end of CoverTab[49310]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:66
	_go_fuzz_dep_.CoverTab[49311]++
													return string(b)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:67
	// _ = "end of CoverTab[49311]"
}

// GoSanitized converts a string to a valid Go identifier.
func GoSanitized(s string) string {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:71
	_go_fuzz_dep_.CoverTab[49330]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:74
	s = strings.Map(func(r rune) rune {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:74
		_go_fuzz_dep_.CoverTab[49333]++
														if unicode.IsLetter(r) || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:75
			_go_fuzz_dep_.CoverTab[49335]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:75
			return unicode.IsDigit(r)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:75
			// _ = "end of CoverTab[49335]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:75
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:75
			_go_fuzz_dep_.CoverTab[49336]++
															return r
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:76
			// _ = "end of CoverTab[49336]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:77
			_go_fuzz_dep_.CoverTab[49337]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:77
			// _ = "end of CoverTab[49337]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:77
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:77
		// _ = "end of CoverTab[49333]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:77
		_go_fuzz_dep_.CoverTab[49334]++
														return '_'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:78
		// _ = "end of CoverTab[49334]"
	}, s)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:79
	// _ = "end of CoverTab[49330]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:79
	_go_fuzz_dep_.CoverTab[49331]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:83
	r, _ := utf8.DecodeRuneInString(s)
	if token.Lookup(s).IsKeyword() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:84
		_go_fuzz_dep_.CoverTab[49338]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:84
		return !unicode.IsLetter(r)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:84
		// _ = "end of CoverTab[49338]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:84
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:84
		_go_fuzz_dep_.CoverTab[49339]++
														return "_" + s
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:85
		// _ = "end of CoverTab[49339]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:86
		_go_fuzz_dep_.CoverTab[49340]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:86
		// _ = "end of CoverTab[49340]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:86
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:86
	// _ = "end of CoverTab[49331]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:86
	_go_fuzz_dep_.CoverTab[49332]++
													return s
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:87
	// _ = "end of CoverTab[49332]"
}

// JSONCamelCase converts a snake_case identifier to a camelCase identifier,
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:90
// according to the protobuf JSON specification.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:92
func JSONCamelCase(s string) string {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:92
	_go_fuzz_dep_.CoverTab[49341]++
													var b []byte
													var wasUnderscore bool
													for i := 0; i < len(s); i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:95
		_go_fuzz_dep_.CoverTab[49343]++
														c := s[i]
														if c != '_' {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:97
			_go_fuzz_dep_.CoverTab[49345]++
															if wasUnderscore && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:98
				_go_fuzz_dep_.CoverTab[49347]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:98
				return isASCIILower(c)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:98
				// _ = "end of CoverTab[49347]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:98
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:98
				_go_fuzz_dep_.CoverTab[49348]++
																c -= 'a' - 'A'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:99
				// _ = "end of CoverTab[49348]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:100
				_go_fuzz_dep_.CoverTab[49349]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:100
				// _ = "end of CoverTab[49349]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:100
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:100
			// _ = "end of CoverTab[49345]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:100
			_go_fuzz_dep_.CoverTab[49346]++
															b = append(b, c)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:101
			// _ = "end of CoverTab[49346]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:102
			_go_fuzz_dep_.CoverTab[49350]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:102
			// _ = "end of CoverTab[49350]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:102
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:102
		// _ = "end of CoverTab[49343]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:102
		_go_fuzz_dep_.CoverTab[49344]++
														wasUnderscore = c == '_'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:103
		// _ = "end of CoverTab[49344]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:104
	// _ = "end of CoverTab[49341]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:104
	_go_fuzz_dep_.CoverTab[49342]++
													return string(b)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:105
	// _ = "end of CoverTab[49342]"
}

// JSONSnakeCase converts a camelCase identifier to a snake_case identifier,
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:108
// according to the protobuf JSON specification.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:110
func JSONSnakeCase(s string) string {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:110
	_go_fuzz_dep_.CoverTab[49351]++
													var b []byte
													for i := 0; i < len(s); i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:112
		_go_fuzz_dep_.CoverTab[49353]++
														c := s[i]
														if isASCIIUpper(c) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:114
			_go_fuzz_dep_.CoverTab[49355]++
															b = append(b, '_')
															c += 'a' - 'A'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:116
			// _ = "end of CoverTab[49355]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:117
			_go_fuzz_dep_.CoverTab[49356]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:117
			// _ = "end of CoverTab[49356]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:117
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:117
		// _ = "end of CoverTab[49353]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:117
		_go_fuzz_dep_.CoverTab[49354]++
														b = append(b, c)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:118
		// _ = "end of CoverTab[49354]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:119
	// _ = "end of CoverTab[49351]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:119
	_go_fuzz_dep_.CoverTab[49352]++
													return string(b)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:120
	// _ = "end of CoverTab[49352]"
}

// MapEntryName derives the name of the map entry message given the field name.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:123
// See protoc v3.8.0: src/google/protobuf/descriptor.cc:254-276,6057
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:125
func MapEntryName(s string) string {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:125
	_go_fuzz_dep_.CoverTab[49357]++
													var b []byte
													upperNext := true
													for _, c := range s {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:128
		_go_fuzz_dep_.CoverTab[49359]++
														switch {
		case c == '_':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:130
			_go_fuzz_dep_.CoverTab[49360]++
															upperNext = true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:131
			// _ = "end of CoverTab[49360]"
		case upperNext:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:132
			_go_fuzz_dep_.CoverTab[49361]++
															b = append(b, byte(unicode.ToUpper(c)))
															upperNext = false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:134
			// _ = "end of CoverTab[49361]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:135
			_go_fuzz_dep_.CoverTab[49362]++
															b = append(b, byte(c))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:136
			// _ = "end of CoverTab[49362]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:137
		// _ = "end of CoverTab[49359]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:138
	// _ = "end of CoverTab[49357]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:138
	_go_fuzz_dep_.CoverTab[49358]++
													b = append(b, "Entry"...)
													return string(b)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:140
	// _ = "end of CoverTab[49358]"
}

// EnumValueName derives the camel-cased enum value name.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:143
// See protoc v3.8.0: src/google/protobuf/descriptor.cc:297-313
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:145
func EnumValueName(s string) string {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:145
	_go_fuzz_dep_.CoverTab[49363]++
													var b []byte
													upperNext := true
													for _, c := range s {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:148
		_go_fuzz_dep_.CoverTab[49365]++
														switch {
		case c == '_':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:150
			_go_fuzz_dep_.CoverTab[49366]++
															upperNext = true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:151
			// _ = "end of CoverTab[49366]"
		case upperNext:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:152
			_go_fuzz_dep_.CoverTab[49367]++
															b = append(b, byte(unicode.ToUpper(c)))
															upperNext = false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:154
			// _ = "end of CoverTab[49367]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:155
			_go_fuzz_dep_.CoverTab[49368]++
															b = append(b, byte(unicode.ToLower(c)))
															upperNext = false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:157
			// _ = "end of CoverTab[49368]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:158
		// _ = "end of CoverTab[49365]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:159
	// _ = "end of CoverTab[49363]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:159
	_go_fuzz_dep_.CoverTab[49364]++
													return string(b)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:160
	// _ = "end of CoverTab[49364]"
}

// TrimEnumPrefix trims the enum name prefix from an enum value name,
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:163
// where the prefix is all lowercase without underscores.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:163
// See protoc v3.8.0: src/google/protobuf/descriptor.cc:330-375
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:166
func TrimEnumPrefix(s, prefix string) string {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:166
	_go_fuzz_dep_.CoverTab[49369]++
													s0 := s
													for len(s) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:168
		_go_fuzz_dep_.CoverTab[49373]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:168
		return len(prefix) > 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:168
		// _ = "end of CoverTab[49373]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:168
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:168
		_go_fuzz_dep_.CoverTab[49374]++
														if s[0] == '_' {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:169
			_go_fuzz_dep_.CoverTab[49377]++
															s = s[1:]
															continue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:171
			// _ = "end of CoverTab[49377]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:172
			_go_fuzz_dep_.CoverTab[49378]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:172
			// _ = "end of CoverTab[49378]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:172
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:172
		// _ = "end of CoverTab[49374]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:172
		_go_fuzz_dep_.CoverTab[49375]++
														if unicode.ToLower(rune(s[0])) != rune(prefix[0]) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:173
			_go_fuzz_dep_.CoverTab[49379]++
															return s0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:174
			// _ = "end of CoverTab[49379]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:175
			_go_fuzz_dep_.CoverTab[49380]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:175
			// _ = "end of CoverTab[49380]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:175
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:175
		// _ = "end of CoverTab[49375]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:175
		_go_fuzz_dep_.CoverTab[49376]++
														s, prefix = s[1:], prefix[1:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:176
		// _ = "end of CoverTab[49376]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:177
	// _ = "end of CoverTab[49369]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:177
	_go_fuzz_dep_.CoverTab[49370]++
													if len(prefix) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:178
		_go_fuzz_dep_.CoverTab[49381]++
														return s0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:179
		// _ = "end of CoverTab[49381]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:180
		_go_fuzz_dep_.CoverTab[49382]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:180
		// _ = "end of CoverTab[49382]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:180
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:180
	// _ = "end of CoverTab[49370]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:180
	_go_fuzz_dep_.CoverTab[49371]++
													s = strings.TrimLeft(s, "_")
													if len(s) == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:182
		_go_fuzz_dep_.CoverTab[49383]++
														return s0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:183
		// _ = "end of CoverTab[49383]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:184
		_go_fuzz_dep_.CoverTab[49384]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:184
		// _ = "end of CoverTab[49384]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:184
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:184
	// _ = "end of CoverTab[49371]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:184
	_go_fuzz_dep_.CoverTab[49372]++
													return s
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:185
	// _ = "end of CoverTab[49372]"
}

func isASCIILower(c byte) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:188
	_go_fuzz_dep_.CoverTab[49385]++
													return 'a' <= c && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:189
		_go_fuzz_dep_.CoverTab[49386]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:189
		return c <= 'z'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:189
		// _ = "end of CoverTab[49386]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:189
	}()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:189
	// _ = "end of CoverTab[49385]"
}
func isASCIIUpper(c byte) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:191
	_go_fuzz_dep_.CoverTab[49387]++
													return 'A' <= c && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:192
		_go_fuzz_dep_.CoverTab[49388]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:192
		return c <= 'Z'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:192
		// _ = "end of CoverTab[49388]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:192
	}()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:192
	// _ = "end of CoverTab[49387]"
}
func isASCIIDigit(c byte) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:194
	_go_fuzz_dep_.CoverTab[49389]++
													return '0' <= c && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:195
		_go_fuzz_dep_.CoverTab[49390]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:195
		return c <= '9'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:195
		// _ = "end of CoverTab[49390]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:195
	}()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:195
	// _ = "end of CoverTab[49389]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:196
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings.go:196
var _ = _go_fuzz_dep_.CoverTab
