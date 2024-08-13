// Protocol Buffers for Go with Gadgets
//
// Copyright (c) 2013, The GoGo Authors. All rights reserved.
// http://github.com/gogo/protobuf
//
// Go support for Protocol Buffers - Google's data interchange format
//
// Copyright 2010 The Go Authors.  All rights reserved.
// https://github.com/golang/protobuf
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//     * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//     * Neither the name of Google Inc. nor the names of its
// contributors may be used to endorse or promote products derived from
// this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:37
package proto

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:37
import (
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:37
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:37
)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:37
import (
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:37
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:37
)

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:42
import (
	"encoding"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

// Error string emitted when deserializing Any and fields are already set
const anyRepeatedlyUnpacked = "Any message unpacked multiple times, or %q already set"

type ParseError struct {
	Message	string
	Line	int	// 1-based line number
	Offset	int	// 0-based byte offset from start of input
}

func (p *ParseError) Error() string {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:62
	_go_fuzz_dep_.CoverTab[112998]++
												if p.Line == 1 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:63
		_go_fuzz_dep_.CoverTab[113000]++

													return fmt.Sprintf("line 1.%d: %v", p.Offset, p.Message)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:65
		// _ = "end of CoverTab[113000]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:66
		_go_fuzz_dep_.CoverTab[113001]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:66
		// _ = "end of CoverTab[113001]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:66
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:66
	// _ = "end of CoverTab[112998]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:66
	_go_fuzz_dep_.CoverTab[112999]++
												return fmt.Sprintf("line %d: %v", p.Line, p.Message)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:67
	// _ = "end of CoverTab[112999]"
}

type token struct {
	value		string
	err		*ParseError
	line		int	// line number
	offset		int	// byte number from start of input, not start of line
	unquoted	string	// the unquoted version of value, if it was a quoted string
}

func (t *token) String() string {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:78
	_go_fuzz_dep_.CoverTab[113002]++
												if t.err == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:79
		_go_fuzz_dep_.CoverTab[113004]++
													return fmt.Sprintf("%q (line=%d, offset=%d)", t.value, t.line, t.offset)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:80
		// _ = "end of CoverTab[113004]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:81
		_go_fuzz_dep_.CoverTab[113005]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:81
		// _ = "end of CoverTab[113005]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:81
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:81
	// _ = "end of CoverTab[113002]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:81
	_go_fuzz_dep_.CoverTab[113003]++
												return fmt.Sprintf("parse error: %v", t.err)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:82
	// _ = "end of CoverTab[113003]"
}

type textParser struct {
	s		string	// remaining input
	done		bool	// whether the parsing is finished (success or error)
	backed		bool	// whether back() was called
	offset, line	int
	cur		token
}

func newTextParser(s string) *textParser {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:93
	_go_fuzz_dep_.CoverTab[113006]++
												p := new(textParser)
												p.s = s
												p.line = 1
												p.cur.line = 1
												return p
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:98
	// _ = "end of CoverTab[113006]"
}

func (p *textParser) errorf(format string, a ...interface{}) *ParseError {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:101
	_go_fuzz_dep_.CoverTab[113007]++
												pe := &ParseError{fmt.Sprintf(format, a...), p.cur.line, p.cur.offset}
												p.cur.err = pe
												p.done = true
												return pe
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:105
	// _ = "end of CoverTab[113007]"
}

// Numbers and identifiers are matched by [-+._A-Za-z0-9]
func isIdentOrNumberChar(c byte) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:109
	_go_fuzz_dep_.CoverTab[113008]++
												switch {
	case 'A' <= c && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:111
		_go_fuzz_dep_.CoverTab[113014]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:111
		return c <= 'Z'
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:111
		// _ = "end of CoverTab[113014]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:111
	}(), 'a' <= c && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:111
		_go_fuzz_dep_.CoverTab[113015]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:111
		return c <= 'z'
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:111
		// _ = "end of CoverTab[113015]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:111
	}():
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:111
		_go_fuzz_dep_.CoverTab[113011]++
													return true
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:112
		// _ = "end of CoverTab[113011]"
	case '0' <= c && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:113
		_go_fuzz_dep_.CoverTab[113016]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:113
		return c <= '9'
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:113
		// _ = "end of CoverTab[113016]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:113
	}():
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:113
		_go_fuzz_dep_.CoverTab[113012]++
													return true
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:114
		// _ = "end of CoverTab[113012]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:114
	default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:114
		_go_fuzz_dep_.CoverTab[113013]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:114
		// _ = "end of CoverTab[113013]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:115
	// _ = "end of CoverTab[113008]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:115
	_go_fuzz_dep_.CoverTab[113009]++
												switch c {
	case '-', '+', '.', '_':
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:117
		_go_fuzz_dep_.CoverTab[113017]++
													return true
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:118
		// _ = "end of CoverTab[113017]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:118
	default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:118
		_go_fuzz_dep_.CoverTab[113018]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:118
		// _ = "end of CoverTab[113018]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:119
	// _ = "end of CoverTab[113009]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:119
	_go_fuzz_dep_.CoverTab[113010]++
												return false
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:120
	// _ = "end of CoverTab[113010]"
}

func isWhitespace(c byte) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:123
	_go_fuzz_dep_.CoverTab[113019]++
												switch c {
	case ' ', '\t', '\n', '\r':
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:125
		_go_fuzz_dep_.CoverTab[113021]++
													return true
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:126
		// _ = "end of CoverTab[113021]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:126
	default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:126
		_go_fuzz_dep_.CoverTab[113022]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:126
		// _ = "end of CoverTab[113022]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:127
	// _ = "end of CoverTab[113019]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:127
	_go_fuzz_dep_.CoverTab[113020]++
												return false
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:128
	// _ = "end of CoverTab[113020]"
}

func isQuote(c byte) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:131
	_go_fuzz_dep_.CoverTab[113023]++
												switch c {
	case '"', '\'':
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:133
		_go_fuzz_dep_.CoverTab[113025]++
													return true
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:134
		// _ = "end of CoverTab[113025]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:134
	default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:134
		_go_fuzz_dep_.CoverTab[113026]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:134
		// _ = "end of CoverTab[113026]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:135
	// _ = "end of CoverTab[113023]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:135
	_go_fuzz_dep_.CoverTab[113024]++
												return false
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:136
	// _ = "end of CoverTab[113024]"
}

func (p *textParser) skipWhitespace() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:139
	_go_fuzz_dep_.CoverTab[113027]++
												i := 0
												for i < len(p.s) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:141
		_go_fuzz_dep_.CoverTab[113029]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:141
		return (isWhitespace(p.s[i]) || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:141
			_go_fuzz_dep_.CoverTab[113030]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:141
			return p.s[i] == '#'
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:141
			// _ = "end of CoverTab[113030]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:141
		}())
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:141
		// _ = "end of CoverTab[113029]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:141
	}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:141
		_go_fuzz_dep_.CoverTab[113031]++
													if p.s[i] == '#' {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:142
			_go_fuzz_dep_.CoverTab[113034]++

														for i < len(p.s) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:144
				_go_fuzz_dep_.CoverTab[113036]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:144
				return p.s[i] != '\n'
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:144
				// _ = "end of CoverTab[113036]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:144
			}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:144
				_go_fuzz_dep_.CoverTab[113037]++
															i++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:145
				// _ = "end of CoverTab[113037]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:146
			// _ = "end of CoverTab[113034]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:146
			_go_fuzz_dep_.CoverTab[113035]++
														if i == len(p.s) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:147
				_go_fuzz_dep_.CoverTab[113038]++
															break
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:148
				// _ = "end of CoverTab[113038]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:149
				_go_fuzz_dep_.CoverTab[113039]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:149
				// _ = "end of CoverTab[113039]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:149
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:149
			// _ = "end of CoverTab[113035]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:150
			_go_fuzz_dep_.CoverTab[113040]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:150
			// _ = "end of CoverTab[113040]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:150
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:150
		// _ = "end of CoverTab[113031]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:150
		_go_fuzz_dep_.CoverTab[113032]++
													if p.s[i] == '\n' {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:151
			_go_fuzz_dep_.CoverTab[113041]++
														p.line++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:152
			// _ = "end of CoverTab[113041]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:153
			_go_fuzz_dep_.CoverTab[113042]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:153
			// _ = "end of CoverTab[113042]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:153
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:153
		// _ = "end of CoverTab[113032]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:153
		_go_fuzz_dep_.CoverTab[113033]++
													i++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:154
		// _ = "end of CoverTab[113033]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:155
	// _ = "end of CoverTab[113027]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:155
	_go_fuzz_dep_.CoverTab[113028]++
												p.offset += i
												p.s = p.s[i:len(p.s)]
												if len(p.s) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:158
		_go_fuzz_dep_.CoverTab[113043]++
													p.done = true
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:159
		// _ = "end of CoverTab[113043]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:160
		_go_fuzz_dep_.CoverTab[113044]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:160
		// _ = "end of CoverTab[113044]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:160
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:160
	// _ = "end of CoverTab[113028]"
}

func (p *textParser) advance() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:163
	_go_fuzz_dep_.CoverTab[113045]++

												p.skipWhitespace()
												if p.done {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:166
		_go_fuzz_dep_.CoverTab[113048]++
													return
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:167
		// _ = "end of CoverTab[113048]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:168
		_go_fuzz_dep_.CoverTab[113049]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:168
		// _ = "end of CoverTab[113049]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:168
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:168
	// _ = "end of CoverTab[113045]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:168
	_go_fuzz_dep_.CoverTab[113046]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:171
	p.cur.err = nil
	p.cur.offset, p.cur.line = p.offset, p.line
	p.cur.unquoted = ""
	switch p.s[0] {
	case '<', '>', '{', '}', ':', '[', ']', ';', ',', '/':
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:175
		_go_fuzz_dep_.CoverTab[113050]++

													p.cur.value, p.s = p.s[0:1], p.s[1:len(p.s)]
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:177
		// _ = "end of CoverTab[113050]"
	case '"', '\'':
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:178
		_go_fuzz_dep_.CoverTab[113051]++

													i := 1
													for i < len(p.s) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:181
			_go_fuzz_dep_.CoverTab[113058]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:181
			return p.s[i] != p.s[0]
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:181
			// _ = "end of CoverTab[113058]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:181
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:181
			_go_fuzz_dep_.CoverTab[113059]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:181
			return p.s[i] != '\n'
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:181
			// _ = "end of CoverTab[113059]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:181
		}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:181
			_go_fuzz_dep_.CoverTab[113060]++
														if p.s[i] == '\\' && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:182
				_go_fuzz_dep_.CoverTab[113062]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:182
				return i+1 < len(p.s)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:182
				// _ = "end of CoverTab[113062]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:182
			}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:182
				_go_fuzz_dep_.CoverTab[113063]++

															i++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:184
				// _ = "end of CoverTab[113063]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:185
				_go_fuzz_dep_.CoverTab[113064]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:185
				// _ = "end of CoverTab[113064]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:185
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:185
			// _ = "end of CoverTab[113060]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:185
			_go_fuzz_dep_.CoverTab[113061]++
														i++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:186
			// _ = "end of CoverTab[113061]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:187
		// _ = "end of CoverTab[113051]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:187
		_go_fuzz_dep_.CoverTab[113052]++
													if i >= len(p.s) || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:188
			_go_fuzz_dep_.CoverTab[113065]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:188
			return p.s[i] != p.s[0]
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:188
			// _ = "end of CoverTab[113065]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:188
		}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:188
			_go_fuzz_dep_.CoverTab[113066]++
														p.errorf("unmatched quote")
														return
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:190
			// _ = "end of CoverTab[113066]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:191
			_go_fuzz_dep_.CoverTab[113067]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:191
			// _ = "end of CoverTab[113067]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:191
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:191
		// _ = "end of CoverTab[113052]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:191
		_go_fuzz_dep_.CoverTab[113053]++
													unq, err := unquoteC(p.s[1:i], rune(p.s[0]))
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:193
			_go_fuzz_dep_.CoverTab[113068]++
														p.errorf("invalid quoted string %s: %v", p.s[0:i+1], err)
														return
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:195
			// _ = "end of CoverTab[113068]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:196
			_go_fuzz_dep_.CoverTab[113069]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:196
			// _ = "end of CoverTab[113069]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:196
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:196
		// _ = "end of CoverTab[113053]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:196
		_go_fuzz_dep_.CoverTab[113054]++
													p.cur.value, p.s = p.s[0:i+1], p.s[i+1:len(p.s)]
													p.cur.unquoted = unq
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:198
		// _ = "end of CoverTab[113054]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:199
		_go_fuzz_dep_.CoverTab[113055]++
													i := 0
													for i < len(p.s) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:201
			_go_fuzz_dep_.CoverTab[113070]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:201
			return isIdentOrNumberChar(p.s[i])
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:201
			// _ = "end of CoverTab[113070]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:201
		}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:201
			_go_fuzz_dep_.CoverTab[113071]++
														i++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:202
			// _ = "end of CoverTab[113071]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:203
		// _ = "end of CoverTab[113055]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:203
		_go_fuzz_dep_.CoverTab[113056]++
													if i == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:204
			_go_fuzz_dep_.CoverTab[113072]++
														p.errorf("unexpected byte %#x", p.s[0])
														return
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:206
			// _ = "end of CoverTab[113072]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:207
			_go_fuzz_dep_.CoverTab[113073]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:207
			// _ = "end of CoverTab[113073]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:207
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:207
		// _ = "end of CoverTab[113056]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:207
		_go_fuzz_dep_.CoverTab[113057]++
													p.cur.value, p.s = p.s[0:i], p.s[i:len(p.s)]
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:208
		// _ = "end of CoverTab[113057]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:209
	// _ = "end of CoverTab[113046]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:209
	_go_fuzz_dep_.CoverTab[113047]++
												p.offset += len(p.cur.value)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:210
	// _ = "end of CoverTab[113047]"
}

var (
	errBadUTF8 = errors.New("proto: bad UTF-8")
)

func unquoteC(s string, quote rune) (string, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:217
	_go_fuzz_dep_.CoverTab[113074]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:223
	simple := true
	for _, r := range s {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:224
		_go_fuzz_dep_.CoverTab[113078]++
													if r == '\\' || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:225
			_go_fuzz_dep_.CoverTab[113079]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:225
			return r == quote
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:225
			// _ = "end of CoverTab[113079]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:225
		}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:225
			_go_fuzz_dep_.CoverTab[113080]++
														simple = false
														break
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:227
			// _ = "end of CoverTab[113080]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:228
			_go_fuzz_dep_.CoverTab[113081]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:228
			// _ = "end of CoverTab[113081]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:228
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:228
		// _ = "end of CoverTab[113078]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:229
	// _ = "end of CoverTab[113074]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:229
	_go_fuzz_dep_.CoverTab[113075]++
												if simple {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:230
		_go_fuzz_dep_.CoverTab[113082]++
													return s, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:231
		// _ = "end of CoverTab[113082]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:232
		_go_fuzz_dep_.CoverTab[113083]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:232
		// _ = "end of CoverTab[113083]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:232
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:232
	// _ = "end of CoverTab[113075]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:232
	_go_fuzz_dep_.CoverTab[113076]++

												buf := make([]byte, 0, 3*len(s)/2)
												for len(s) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:235
		_go_fuzz_dep_.CoverTab[113084]++
													r, n := utf8.DecodeRuneInString(s)
													if r == utf8.RuneError && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:237
			_go_fuzz_dep_.CoverTab[113088]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:237
			return n == 1
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:237
			// _ = "end of CoverTab[113088]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:237
		}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:237
			_go_fuzz_dep_.CoverTab[113089]++
														return "", errBadUTF8
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:238
			// _ = "end of CoverTab[113089]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:239
			_go_fuzz_dep_.CoverTab[113090]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:239
			// _ = "end of CoverTab[113090]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:239
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:239
		// _ = "end of CoverTab[113084]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:239
		_go_fuzz_dep_.CoverTab[113085]++
													s = s[n:]
													if r != '\\' {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:241
			_go_fuzz_dep_.CoverTab[113091]++
														if r < utf8.RuneSelf {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:242
				_go_fuzz_dep_.CoverTab[113093]++
															buf = append(buf, byte(r))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:243
				// _ = "end of CoverTab[113093]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:244
				_go_fuzz_dep_.CoverTab[113094]++
															buf = append(buf, string(r)...)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:245
				// _ = "end of CoverTab[113094]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:246
			// _ = "end of CoverTab[113091]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:246
			_go_fuzz_dep_.CoverTab[113092]++
														continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:247
			// _ = "end of CoverTab[113092]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:248
			_go_fuzz_dep_.CoverTab[113095]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:248
			// _ = "end of CoverTab[113095]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:248
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:248
		// _ = "end of CoverTab[113085]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:248
		_go_fuzz_dep_.CoverTab[113086]++

													ch, tail, err := unescape(s)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:251
			_go_fuzz_dep_.CoverTab[113096]++
														return "", err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:252
			// _ = "end of CoverTab[113096]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:253
			_go_fuzz_dep_.CoverTab[113097]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:253
			// _ = "end of CoverTab[113097]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:253
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:253
		// _ = "end of CoverTab[113086]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:253
		_go_fuzz_dep_.CoverTab[113087]++
													buf = append(buf, ch...)
													s = tail
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:255
		// _ = "end of CoverTab[113087]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:256
	// _ = "end of CoverTab[113076]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:256
	_go_fuzz_dep_.CoverTab[113077]++
												return string(buf), nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:257
	// _ = "end of CoverTab[113077]"
}

func unescape(s string) (ch string, tail string, err error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:260
	_go_fuzz_dep_.CoverTab[113098]++
												r, n := utf8.DecodeRuneInString(s)
												if r == utf8.RuneError && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:262
		_go_fuzz_dep_.CoverTab[113101]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:262
		return n == 1
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:262
		// _ = "end of CoverTab[113101]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:262
	}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:262
		_go_fuzz_dep_.CoverTab[113102]++
													return "", "", errBadUTF8
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:263
		// _ = "end of CoverTab[113102]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:264
		_go_fuzz_dep_.CoverTab[113103]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:264
		// _ = "end of CoverTab[113103]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:264
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:264
	// _ = "end of CoverTab[113098]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:264
	_go_fuzz_dep_.CoverTab[113099]++
												s = s[n:]
												switch r {
	case 'a':
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:267
		_go_fuzz_dep_.CoverTab[113104]++
													return "\a", s, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:268
		// _ = "end of CoverTab[113104]"
	case 'b':
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:269
		_go_fuzz_dep_.CoverTab[113105]++
													return "\b", s, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:270
		// _ = "end of CoverTab[113105]"
	case 'f':
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:271
		_go_fuzz_dep_.CoverTab[113106]++
													return "\f", s, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:272
		// _ = "end of CoverTab[113106]"
	case 'n':
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:273
		_go_fuzz_dep_.CoverTab[113107]++
													return "\n", s, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:274
		// _ = "end of CoverTab[113107]"
	case 'r':
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:275
		_go_fuzz_dep_.CoverTab[113108]++
													return "\r", s, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:276
		// _ = "end of CoverTab[113108]"
	case 't':
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:277
		_go_fuzz_dep_.CoverTab[113109]++
													return "\t", s, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:278
		// _ = "end of CoverTab[113109]"
	case 'v':
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:279
		_go_fuzz_dep_.CoverTab[113110]++
													return "\v", s, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:280
		// _ = "end of CoverTab[113110]"
	case '?':
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:281
		_go_fuzz_dep_.CoverTab[113111]++
													return "?", s, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:282
		// _ = "end of CoverTab[113111]"
	case '\'', '"', '\\':
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:283
		_go_fuzz_dep_.CoverTab[113112]++
													return string(r), s, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:284
		// _ = "end of CoverTab[113112]"
	case '0', '1', '2', '3', '4', '5', '6', '7':
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:285
		_go_fuzz_dep_.CoverTab[113113]++
													if len(s) < 2 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:286
			_go_fuzz_dep_.CoverTab[113123]++
														return "", "", fmt.Errorf(`\%c requires 2 following digits`, r)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:287
			// _ = "end of CoverTab[113123]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:288
			_go_fuzz_dep_.CoverTab[113124]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:288
			// _ = "end of CoverTab[113124]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:288
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:288
		// _ = "end of CoverTab[113113]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:288
		_go_fuzz_dep_.CoverTab[113114]++
													ss := string(r) + s[:2]
													s = s[2:]
													i, err := strconv.ParseUint(ss, 8, 8)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:292
			_go_fuzz_dep_.CoverTab[113125]++
														return "", "", fmt.Errorf(`\%s contains non-octal digits`, ss)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:293
			// _ = "end of CoverTab[113125]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:294
			_go_fuzz_dep_.CoverTab[113126]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:294
			// _ = "end of CoverTab[113126]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:294
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:294
		// _ = "end of CoverTab[113114]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:294
		_go_fuzz_dep_.CoverTab[113115]++
													return string([]byte{byte(i)}), s, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:295
		// _ = "end of CoverTab[113115]"
	case 'x', 'X', 'u', 'U':
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:296
		_go_fuzz_dep_.CoverTab[113116]++
													var n int
													switch r {
		case 'x', 'X':
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:299
			_go_fuzz_dep_.CoverTab[113127]++
														n = 2
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:300
			// _ = "end of CoverTab[113127]"
		case 'u':
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:301
			_go_fuzz_dep_.CoverTab[113128]++
														n = 4
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:302
			// _ = "end of CoverTab[113128]"
		case 'U':
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:303
			_go_fuzz_dep_.CoverTab[113129]++
														n = 8
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:304
			// _ = "end of CoverTab[113129]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:304
		default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:304
			_go_fuzz_dep_.CoverTab[113130]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:304
			// _ = "end of CoverTab[113130]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:305
		// _ = "end of CoverTab[113116]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:305
		_go_fuzz_dep_.CoverTab[113117]++
													if len(s) < n {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:306
			_go_fuzz_dep_.CoverTab[113131]++
														return "", "", fmt.Errorf(`\%c requires %d following digits`, r, n)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:307
			// _ = "end of CoverTab[113131]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:308
			_go_fuzz_dep_.CoverTab[113132]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:308
			// _ = "end of CoverTab[113132]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:308
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:308
		// _ = "end of CoverTab[113117]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:308
		_go_fuzz_dep_.CoverTab[113118]++
													ss := s[:n]
													s = s[n:]
													i, err := strconv.ParseUint(ss, 16, 64)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:312
			_go_fuzz_dep_.CoverTab[113133]++
														return "", "", fmt.Errorf(`\%c%s contains non-hexadecimal digits`, r, ss)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:313
			// _ = "end of CoverTab[113133]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:314
			_go_fuzz_dep_.CoverTab[113134]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:314
			// _ = "end of CoverTab[113134]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:314
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:314
		// _ = "end of CoverTab[113118]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:314
		_go_fuzz_dep_.CoverTab[113119]++
													if r == 'x' || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:315
			_go_fuzz_dep_.CoverTab[113135]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:315
			return r == 'X'
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:315
			// _ = "end of CoverTab[113135]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:315
		}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:315
			_go_fuzz_dep_.CoverTab[113136]++
														return string([]byte{byte(i)}), s, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:316
			// _ = "end of CoverTab[113136]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:317
			_go_fuzz_dep_.CoverTab[113137]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:317
			// _ = "end of CoverTab[113137]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:317
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:317
		// _ = "end of CoverTab[113119]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:317
		_go_fuzz_dep_.CoverTab[113120]++
													if i > utf8.MaxRune {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:318
			_go_fuzz_dep_.CoverTab[113138]++
														return "", "", fmt.Errorf(`\%c%s is not a valid Unicode code point`, r, ss)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:319
			// _ = "end of CoverTab[113138]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:320
			_go_fuzz_dep_.CoverTab[113139]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:320
			// _ = "end of CoverTab[113139]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:320
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:320
		// _ = "end of CoverTab[113120]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:320
		_go_fuzz_dep_.CoverTab[113121]++
													return string(rune(i)), s, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:321
		// _ = "end of CoverTab[113121]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:321
	default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:321
		_go_fuzz_dep_.CoverTab[113122]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:321
		// _ = "end of CoverTab[113122]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:322
	// _ = "end of CoverTab[113099]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:322
	_go_fuzz_dep_.CoverTab[113100]++
												return "", "", fmt.Errorf(`unknown escape \%c`, r)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:323
	// _ = "end of CoverTab[113100]"
}

// Back off the parser by one token. Can only be done between calls to next().
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:326
// It makes the next advance() a no-op.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:328
func (p *textParser) back() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:328
	_go_fuzz_dep_.CoverTab[113140]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:328
	p.backed = true
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:328
	// _ = "end of CoverTab[113140]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:328
}

// Advances the parser and returns the new current token.
func (p *textParser) next() *token {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:331
	_go_fuzz_dep_.CoverTab[113141]++
												if p.backed || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:332
		_go_fuzz_dep_.CoverTab[113144]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:332
		return p.done
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:332
		// _ = "end of CoverTab[113144]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:332
	}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:332
		_go_fuzz_dep_.CoverTab[113145]++
													p.backed = false
													return &p.cur
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:334
		// _ = "end of CoverTab[113145]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:335
		_go_fuzz_dep_.CoverTab[113146]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:335
		// _ = "end of CoverTab[113146]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:335
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:335
	// _ = "end of CoverTab[113141]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:335
	_go_fuzz_dep_.CoverTab[113142]++
												p.advance()
												if p.done {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:337
		_go_fuzz_dep_.CoverTab[113147]++
													p.cur.value = ""
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:338
		// _ = "end of CoverTab[113147]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:339
		_go_fuzz_dep_.CoverTab[113148]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:339
		if len(p.cur.value) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:339
			_go_fuzz_dep_.CoverTab[113149]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:339
			return isQuote(p.cur.value[0])
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:339
			// _ = "end of CoverTab[113149]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:339
		}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:339
			_go_fuzz_dep_.CoverTab[113150]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:342
			cat := p.cur
			for {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:343
				_go_fuzz_dep_.CoverTab[113152]++
															p.skipWhitespace()
															if p.done || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:345
					_go_fuzz_dep_.CoverTab[113155]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:345
					return !isQuote(p.s[0])
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:345
					// _ = "end of CoverTab[113155]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:345
				}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:345
					_go_fuzz_dep_.CoverTab[113156]++
																break
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:346
					// _ = "end of CoverTab[113156]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:347
					_go_fuzz_dep_.CoverTab[113157]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:347
					// _ = "end of CoverTab[113157]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:347
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:347
				// _ = "end of CoverTab[113152]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:347
				_go_fuzz_dep_.CoverTab[113153]++
															p.advance()
															if p.cur.err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:349
					_go_fuzz_dep_.CoverTab[113158]++
																return &p.cur
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:350
					// _ = "end of CoverTab[113158]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:351
					_go_fuzz_dep_.CoverTab[113159]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:351
					// _ = "end of CoverTab[113159]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:351
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:351
				// _ = "end of CoverTab[113153]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:351
				_go_fuzz_dep_.CoverTab[113154]++
															cat.value += " " + p.cur.value
															cat.unquoted += p.cur.unquoted
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:353
				// _ = "end of CoverTab[113154]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:354
			// _ = "end of CoverTab[113150]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:354
			_go_fuzz_dep_.CoverTab[113151]++
														p.done = false
														p.cur = cat
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:356
			// _ = "end of CoverTab[113151]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:357
			_go_fuzz_dep_.CoverTab[113160]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:357
			// _ = "end of CoverTab[113160]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:357
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:357
		// _ = "end of CoverTab[113148]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:357
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:357
	// _ = "end of CoverTab[113142]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:357
	_go_fuzz_dep_.CoverTab[113143]++
												return &p.cur
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:358
	// _ = "end of CoverTab[113143]"
}

func (p *textParser) consumeToken(s string) error {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:361
	_go_fuzz_dep_.CoverTab[113161]++
												tok := p.next()
												if tok.err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:363
		_go_fuzz_dep_.CoverTab[113164]++
													return tok.err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:364
		// _ = "end of CoverTab[113164]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:365
		_go_fuzz_dep_.CoverTab[113165]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:365
		// _ = "end of CoverTab[113165]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:365
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:365
	// _ = "end of CoverTab[113161]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:365
	_go_fuzz_dep_.CoverTab[113162]++
												if tok.value != s {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:366
		_go_fuzz_dep_.CoverTab[113166]++
													p.back()
													return p.errorf("expected %q, found %q", s, tok.value)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:368
		// _ = "end of CoverTab[113166]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:369
		_go_fuzz_dep_.CoverTab[113167]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:369
		// _ = "end of CoverTab[113167]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:369
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:369
	// _ = "end of CoverTab[113162]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:369
	_go_fuzz_dep_.CoverTab[113163]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:370
	// _ = "end of CoverTab[113163]"
}

// Return a RequiredNotSetError indicating which required field was not set.
func (p *textParser) missingRequiredFieldError(sv reflect.Value) *RequiredNotSetError {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:374
	_go_fuzz_dep_.CoverTab[113168]++
												st := sv.Type()
												sprops := GetProperties(st)
												for i := 0; i < st.NumField(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:377
		_go_fuzz_dep_.CoverTab[113170]++
													if !isNil(sv.Field(i)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:378
			_go_fuzz_dep_.CoverTab[113172]++
														continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:379
			// _ = "end of CoverTab[113172]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:380
			_go_fuzz_dep_.CoverTab[113173]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:380
			// _ = "end of CoverTab[113173]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:380
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:380
		// _ = "end of CoverTab[113170]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:380
		_go_fuzz_dep_.CoverTab[113171]++

													props := sprops.Prop[i]
													if props.Required {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:383
			_go_fuzz_dep_.CoverTab[113174]++
														return &RequiredNotSetError{fmt.Sprintf("%v.%v", st, props.OrigName)}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:384
			// _ = "end of CoverTab[113174]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:385
			_go_fuzz_dep_.CoverTab[113175]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:385
			// _ = "end of CoverTab[113175]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:385
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:385
		// _ = "end of CoverTab[113171]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:386
	// _ = "end of CoverTab[113168]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:386
	_go_fuzz_dep_.CoverTab[113169]++
												return &RequiredNotSetError{fmt.Sprintf("%v.<unknown field name>", st)}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:387
	// _ = "end of CoverTab[113169]"
}

// Returns the index in the struct for the named field, as well as the parsed tag properties.
func structFieldByName(sprops *StructProperties, name string) (int, *Properties, bool) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:391
	_go_fuzz_dep_.CoverTab[113176]++
												i, ok := sprops.decoderOrigNames[name]
												if ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:393
		_go_fuzz_dep_.CoverTab[113178]++
													return i, sprops.Prop[i], true
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:394
		// _ = "end of CoverTab[113178]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:395
		_go_fuzz_dep_.CoverTab[113179]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:395
		// _ = "end of CoverTab[113179]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:395
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:395
	// _ = "end of CoverTab[113176]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:395
	_go_fuzz_dep_.CoverTab[113177]++
												return -1, nil, false
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:396
	// _ = "end of CoverTab[113177]"
}

// Consume a ':' from the input stream (if the next token is a colon),
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:399
// returning an error if a colon is needed but not present.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:401
func (p *textParser) checkForColon(props *Properties, typ reflect.Type) *ParseError {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:401
	_go_fuzz_dep_.CoverTab[113180]++
												tok := p.next()
												if tok.err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:403
		_go_fuzz_dep_.CoverTab[113183]++
													return tok.err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:404
		// _ = "end of CoverTab[113183]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:405
		_go_fuzz_dep_.CoverTab[113184]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:405
		// _ = "end of CoverTab[113184]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:405
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:405
	// _ = "end of CoverTab[113180]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:405
	_go_fuzz_dep_.CoverTab[113181]++
												if tok.value != ":" {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:406
		_go_fuzz_dep_.CoverTab[113185]++

													needColon := true
													switch props.Wire {
		case "group":
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:410
			_go_fuzz_dep_.CoverTab[113188]++
														needColon = false
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:411
			// _ = "end of CoverTab[113188]"
		case "bytes":
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:412
			_go_fuzz_dep_.CoverTab[113189]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:416
			if typ.Kind() == reflect.Ptr {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:416
				_go_fuzz_dep_.CoverTab[113192]++

															if typ.Elem().Kind() == reflect.String {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:418
					_go_fuzz_dep_.CoverTab[113193]++
																break
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:419
					// _ = "end of CoverTab[113193]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:420
					_go_fuzz_dep_.CoverTab[113194]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:420
					// _ = "end of CoverTab[113194]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:420
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:420
				// _ = "end of CoverTab[113192]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:421
				_go_fuzz_dep_.CoverTab[113195]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:421
				if typ.Kind() == reflect.Slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:421
					_go_fuzz_dep_.CoverTab[113196]++

																if typ.Elem().Kind() != reflect.Ptr {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:423
						_go_fuzz_dep_.CoverTab[113197]++
																	break
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:424
						// _ = "end of CoverTab[113197]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:425
						_go_fuzz_dep_.CoverTab[113198]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:425
						// _ = "end of CoverTab[113198]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:425
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:425
					// _ = "end of CoverTab[113196]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:426
					_go_fuzz_dep_.CoverTab[113199]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:426
					if typ.Kind() == reflect.String {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:426
						_go_fuzz_dep_.CoverTab[113200]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:429
						break
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:429
						// _ = "end of CoverTab[113200]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:430
						_go_fuzz_dep_.CoverTab[113201]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:430
						// _ = "end of CoverTab[113201]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:430
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:430
					// _ = "end of CoverTab[113199]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:430
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:430
				// _ = "end of CoverTab[113195]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:430
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:430
			// _ = "end of CoverTab[113189]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:430
			_go_fuzz_dep_.CoverTab[113190]++
														needColon = false
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:431
			// _ = "end of CoverTab[113190]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:431
		default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:431
			_go_fuzz_dep_.CoverTab[113191]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:431
			// _ = "end of CoverTab[113191]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:432
		// _ = "end of CoverTab[113185]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:432
		_go_fuzz_dep_.CoverTab[113186]++
													if needColon {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:433
			_go_fuzz_dep_.CoverTab[113202]++
														return p.errorf("expected ':', found %q", tok.value)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:434
			// _ = "end of CoverTab[113202]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:435
			_go_fuzz_dep_.CoverTab[113203]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:435
			// _ = "end of CoverTab[113203]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:435
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:435
		// _ = "end of CoverTab[113186]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:435
		_go_fuzz_dep_.CoverTab[113187]++
													p.back()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:436
		// _ = "end of CoverTab[113187]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:437
		_go_fuzz_dep_.CoverTab[113204]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:437
		// _ = "end of CoverTab[113204]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:437
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:437
	// _ = "end of CoverTab[113181]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:437
	_go_fuzz_dep_.CoverTab[113182]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:438
	// _ = "end of CoverTab[113182]"
}

func (p *textParser) readStruct(sv reflect.Value, terminator string) error {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:441
	_go_fuzz_dep_.CoverTab[113205]++
												st := sv.Type()
												sprops := GetProperties(st)
												reqCount := sprops.reqCount
												var reqFieldErr error
												fieldSet := make(map[string]bool)

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:453
	for {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:453
		_go_fuzz_dep_.CoverTab[113208]++
													tok := p.next()
													if tok.err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:455
			_go_fuzz_dep_.CoverTab[113219]++
														return tok.err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:456
			// _ = "end of CoverTab[113219]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:457
			_go_fuzz_dep_.CoverTab[113220]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:457
			// _ = "end of CoverTab[113220]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:457
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:457
		// _ = "end of CoverTab[113208]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:457
		_go_fuzz_dep_.CoverTab[113209]++
													if tok.value == terminator {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:458
			_go_fuzz_dep_.CoverTab[113221]++
														break
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:459
			// _ = "end of CoverTab[113221]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:460
			_go_fuzz_dep_.CoverTab[113222]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:460
			// _ = "end of CoverTab[113222]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:460
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:460
		// _ = "end of CoverTab[113209]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:460
		_go_fuzz_dep_.CoverTab[113210]++
													if tok.value == "[" {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:461
			_go_fuzz_dep_.CoverTab[113223]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:466
			extName, err := p.consumeExtName()
			if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:467
				_go_fuzz_dep_.CoverTab[113233]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:468
				// _ = "end of CoverTab[113233]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:469
				_go_fuzz_dep_.CoverTab[113234]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:469
				// _ = "end of CoverTab[113234]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:469
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:469
			// _ = "end of CoverTab[113223]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:469
			_go_fuzz_dep_.CoverTab[113224]++

														if s := strings.LastIndex(extName, "/"); s >= 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:471
				_go_fuzz_dep_.CoverTab[113235]++

															messageName := extName[s+1:]
															mt := MessageType(messageName)
															if mt == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:475
					_go_fuzz_dep_.CoverTab[113244]++
																return p.errorf("unrecognized message %q in google.protobuf.Any", messageName)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:476
					// _ = "end of CoverTab[113244]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:477
					_go_fuzz_dep_.CoverTab[113245]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:477
					// _ = "end of CoverTab[113245]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:477
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:477
				// _ = "end of CoverTab[113235]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:477
				_go_fuzz_dep_.CoverTab[113236]++
															tok = p.next()
															if tok.err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:479
					_go_fuzz_dep_.CoverTab[113246]++
																return tok.err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:480
					// _ = "end of CoverTab[113246]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:481
					_go_fuzz_dep_.CoverTab[113247]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:481
					// _ = "end of CoverTab[113247]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:481
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:481
				// _ = "end of CoverTab[113236]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:481
				_go_fuzz_dep_.CoverTab[113237]++

															if tok.value == ":" {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:483
					_go_fuzz_dep_.CoverTab[113248]++
																tok = p.next()
																if tok.err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:485
						_go_fuzz_dep_.CoverTab[113249]++
																	return tok.err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:486
						// _ = "end of CoverTab[113249]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:487
						_go_fuzz_dep_.CoverTab[113250]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:487
						// _ = "end of CoverTab[113250]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:487
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:487
					// _ = "end of CoverTab[113248]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:488
					_go_fuzz_dep_.CoverTab[113251]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:488
					// _ = "end of CoverTab[113251]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:488
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:488
				// _ = "end of CoverTab[113237]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:488
				_go_fuzz_dep_.CoverTab[113238]++
															var terminator string
															switch tok.value {
				case "<":
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:491
					_go_fuzz_dep_.CoverTab[113252]++
																terminator = ">"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:492
					// _ = "end of CoverTab[113252]"
				case "{":
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:493
					_go_fuzz_dep_.CoverTab[113253]++
																terminator = "}"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:494
					// _ = "end of CoverTab[113253]"
				default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:495
					_go_fuzz_dep_.CoverTab[113254]++
																return p.errorf("expected '{' or '<', found %q", tok.value)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:496
					// _ = "end of CoverTab[113254]"
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:497
				// _ = "end of CoverTab[113238]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:497
				_go_fuzz_dep_.CoverTab[113239]++
															v := reflect.New(mt.Elem())
															if pe := p.readStruct(v.Elem(), terminator); pe != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:499
					_go_fuzz_dep_.CoverTab[113255]++
																return pe
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:500
					// _ = "end of CoverTab[113255]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:501
					_go_fuzz_dep_.CoverTab[113256]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:501
					// _ = "end of CoverTab[113256]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:501
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:501
				// _ = "end of CoverTab[113239]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:501
				_go_fuzz_dep_.CoverTab[113240]++
															b, err := Marshal(v.Interface().(Message))
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:503
					_go_fuzz_dep_.CoverTab[113257]++
																return p.errorf("failed to marshal message of type %q: %v", messageName, err)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:504
					// _ = "end of CoverTab[113257]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:505
					_go_fuzz_dep_.CoverTab[113258]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:505
					// _ = "end of CoverTab[113258]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:505
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:505
				// _ = "end of CoverTab[113240]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:505
				_go_fuzz_dep_.CoverTab[113241]++
															if fieldSet["type_url"] {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:506
					_go_fuzz_dep_.CoverTab[113259]++
																return p.errorf(anyRepeatedlyUnpacked, "type_url")
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:507
					// _ = "end of CoverTab[113259]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:508
					_go_fuzz_dep_.CoverTab[113260]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:508
					// _ = "end of CoverTab[113260]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:508
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:508
				// _ = "end of CoverTab[113241]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:508
				_go_fuzz_dep_.CoverTab[113242]++
															if fieldSet["value"] {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:509
					_go_fuzz_dep_.CoverTab[113261]++
																return p.errorf(anyRepeatedlyUnpacked, "value")
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:510
					// _ = "end of CoverTab[113261]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:511
					_go_fuzz_dep_.CoverTab[113262]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:511
					// _ = "end of CoverTab[113262]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:511
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:511
				// _ = "end of CoverTab[113242]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:511
				_go_fuzz_dep_.CoverTab[113243]++
															sv.FieldByName("TypeUrl").SetString(extName)
															sv.FieldByName("Value").SetBytes(b)
															fieldSet["type_url"] = true
															fieldSet["value"] = true
															continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:516
				// _ = "end of CoverTab[113243]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:517
				_go_fuzz_dep_.CoverTab[113263]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:517
				// _ = "end of CoverTab[113263]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:517
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:517
			// _ = "end of CoverTab[113224]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:517
			_go_fuzz_dep_.CoverTab[113225]++

														var desc *ExtensionDesc

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:522
			for _, d := range RegisteredExtensions(reflect.New(st).Interface().(Message)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:522
				_go_fuzz_dep_.CoverTab[113264]++
															if d.Name == extName {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:523
					_go_fuzz_dep_.CoverTab[113265]++
																desc = d
																break
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:525
					// _ = "end of CoverTab[113265]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:526
					_go_fuzz_dep_.CoverTab[113266]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:526
					// _ = "end of CoverTab[113266]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:526
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:526
				// _ = "end of CoverTab[113264]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:527
			// _ = "end of CoverTab[113225]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:527
			_go_fuzz_dep_.CoverTab[113226]++
														if desc == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:528
				_go_fuzz_dep_.CoverTab[113267]++
															return p.errorf("unrecognized extension %q", extName)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:529
				// _ = "end of CoverTab[113267]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:530
				_go_fuzz_dep_.CoverTab[113268]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:530
				// _ = "end of CoverTab[113268]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:530
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:530
			// _ = "end of CoverTab[113226]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:530
			_go_fuzz_dep_.CoverTab[113227]++

														props := &Properties{}
														props.Parse(desc.Tag)

														typ := reflect.TypeOf(desc.ExtensionType)
														if err := p.checkForColon(props, typ); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:536
				_go_fuzz_dep_.CoverTab[113269]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:537
				// _ = "end of CoverTab[113269]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:538
				_go_fuzz_dep_.CoverTab[113270]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:538
				// _ = "end of CoverTab[113270]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:538
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:538
			// _ = "end of CoverTab[113227]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:538
			_go_fuzz_dep_.CoverTab[113228]++

														rep := desc.repeated()

			// Read the extension structure, and set it in
			// the value we're constructing.
			var ext reflect.Value
			if !rep {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:545
				_go_fuzz_dep_.CoverTab[113271]++
															ext = reflect.New(typ).Elem()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:546
				// _ = "end of CoverTab[113271]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:547
				_go_fuzz_dep_.CoverTab[113272]++
															ext = reflect.New(typ.Elem()).Elem()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:548
				// _ = "end of CoverTab[113272]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:549
			// _ = "end of CoverTab[113228]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:549
			_go_fuzz_dep_.CoverTab[113229]++
														if err := p.readAny(ext, props); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:550
				_go_fuzz_dep_.CoverTab[113273]++
															if _, ok := err.(*RequiredNotSetError); !ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:551
					_go_fuzz_dep_.CoverTab[113275]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:552
					// _ = "end of CoverTab[113275]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:553
					_go_fuzz_dep_.CoverTab[113276]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:553
					// _ = "end of CoverTab[113276]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:553
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:553
				// _ = "end of CoverTab[113273]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:553
				_go_fuzz_dep_.CoverTab[113274]++
															reqFieldErr = err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:554
				// _ = "end of CoverTab[113274]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:555
				_go_fuzz_dep_.CoverTab[113277]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:555
				// _ = "end of CoverTab[113277]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:555
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:555
			// _ = "end of CoverTab[113229]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:555
			_go_fuzz_dep_.CoverTab[113230]++
														ep := sv.Addr().Interface().(Message)
														if !rep {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:557
				_go_fuzz_dep_.CoverTab[113278]++
															SetExtension(ep, desc, ext.Interface())
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:558
				// _ = "end of CoverTab[113278]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:559
				_go_fuzz_dep_.CoverTab[113279]++
															old, err := GetExtension(ep, desc)
															var sl reflect.Value
															if err == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:562
					_go_fuzz_dep_.CoverTab[113281]++
																sl = reflect.ValueOf(old)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:563
					// _ = "end of CoverTab[113281]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:564
					_go_fuzz_dep_.CoverTab[113282]++
																sl = reflect.MakeSlice(typ, 0, 1)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:565
					// _ = "end of CoverTab[113282]"
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:566
				// _ = "end of CoverTab[113279]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:566
				_go_fuzz_dep_.CoverTab[113280]++
															sl = reflect.Append(sl, ext)
															SetExtension(ep, desc, sl.Interface())
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:568
				// _ = "end of CoverTab[113280]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:569
			// _ = "end of CoverTab[113230]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:569
			_go_fuzz_dep_.CoverTab[113231]++
														if err := p.consumeOptionalSeparator(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:570
				_go_fuzz_dep_.CoverTab[113283]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:571
				// _ = "end of CoverTab[113283]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:572
				_go_fuzz_dep_.CoverTab[113284]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:572
				// _ = "end of CoverTab[113284]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:572
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:572
			// _ = "end of CoverTab[113231]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:572
			_go_fuzz_dep_.CoverTab[113232]++
														continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:573
			// _ = "end of CoverTab[113232]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:574
			_go_fuzz_dep_.CoverTab[113285]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:574
			// _ = "end of CoverTab[113285]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:574
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:574
		// _ = "end of CoverTab[113210]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:574
		_go_fuzz_dep_.CoverTab[113211]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:577
		name := tok.value
		var dst reflect.Value
		fi, props, ok := structFieldByName(sprops, name)
		if ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:580
			_go_fuzz_dep_.CoverTab[113286]++
														dst = sv.Field(fi)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:581
			// _ = "end of CoverTab[113286]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:582
			_go_fuzz_dep_.CoverTab[113287]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:582
			if oop, ok := sprops.OneofTypes[name]; ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:582
				_go_fuzz_dep_.CoverTab[113288]++

															props = oop.Prop
															nv := reflect.New(oop.Type.Elem())
															dst = nv.Elem().Field(0)
															field := sv.Field(oop.Field)
															if !field.IsNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:588
					_go_fuzz_dep_.CoverTab[113290]++
																return p.errorf("field '%s' would overwrite already parsed oneof '%s'", name, sv.Type().Field(oop.Field).Name)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:589
					// _ = "end of CoverTab[113290]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:590
					_go_fuzz_dep_.CoverTab[113291]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:590
					// _ = "end of CoverTab[113291]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:590
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:590
				// _ = "end of CoverTab[113288]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:590
				_go_fuzz_dep_.CoverTab[113289]++
															field.Set(nv)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:591
				// _ = "end of CoverTab[113289]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:592
				_go_fuzz_dep_.CoverTab[113292]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:592
				// _ = "end of CoverTab[113292]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:592
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:592
			// _ = "end of CoverTab[113287]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:592
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:592
		// _ = "end of CoverTab[113211]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:592
		_go_fuzz_dep_.CoverTab[113212]++
													if !dst.IsValid() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:593
			_go_fuzz_dep_.CoverTab[113293]++
														return p.errorf("unknown field name %q in %v", name, st)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:594
			// _ = "end of CoverTab[113293]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:595
			_go_fuzz_dep_.CoverTab[113294]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:595
			// _ = "end of CoverTab[113294]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:595
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:595
		// _ = "end of CoverTab[113212]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:595
		_go_fuzz_dep_.CoverTab[113213]++

													if dst.Kind() == reflect.Map {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:597
			_go_fuzz_dep_.CoverTab[113295]++

														if err := p.checkForColon(props, dst.Type()); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:599
				_go_fuzz_dep_.CoverTab[113300]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:600
				// _ = "end of CoverTab[113300]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:601
				_go_fuzz_dep_.CoverTab[113301]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:601
				// _ = "end of CoverTab[113301]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:601
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:601
			// _ = "end of CoverTab[113295]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:601
			_go_fuzz_dep_.CoverTab[113296]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:604
			if dst.IsNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:604
				_go_fuzz_dep_.CoverTab[113302]++
															dst.Set(reflect.MakeMap(dst.Type()))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:605
				// _ = "end of CoverTab[113302]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:606
				_go_fuzz_dep_.CoverTab[113303]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:606
				// _ = "end of CoverTab[113303]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:606
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:606
			// _ = "end of CoverTab[113296]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:606
			_go_fuzz_dep_.CoverTab[113297]++
														key := reflect.New(dst.Type().Key()).Elem()
														val := reflect.New(dst.Type().Elem()).Elem()

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:616
			tok := p.next()
			var terminator string
			switch tok.value {
			case "<":
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:619
				_go_fuzz_dep_.CoverTab[113304]++
															terminator = ">"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:620
				// _ = "end of CoverTab[113304]"
			case "{":
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:621
				_go_fuzz_dep_.CoverTab[113305]++
															terminator = "}"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:622
				// _ = "end of CoverTab[113305]"
			default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:623
				_go_fuzz_dep_.CoverTab[113306]++
															return p.errorf("expected '{' or '<', found %q", tok.value)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:624
				// _ = "end of CoverTab[113306]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:625
			// _ = "end of CoverTab[113297]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:625
			_go_fuzz_dep_.CoverTab[113298]++
														for {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:626
				_go_fuzz_dep_.CoverTab[113307]++
															tok := p.next()
															if tok.err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:628
					_go_fuzz_dep_.CoverTab[113310]++
																return tok.err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:629
					// _ = "end of CoverTab[113310]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:630
					_go_fuzz_dep_.CoverTab[113311]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:630
					// _ = "end of CoverTab[113311]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:630
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:630
				// _ = "end of CoverTab[113307]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:630
				_go_fuzz_dep_.CoverTab[113308]++
															if tok.value == terminator {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:631
					_go_fuzz_dep_.CoverTab[113312]++
																break
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:632
					// _ = "end of CoverTab[113312]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:633
					_go_fuzz_dep_.CoverTab[113313]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:633
					// _ = "end of CoverTab[113313]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:633
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:633
				// _ = "end of CoverTab[113308]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:633
				_go_fuzz_dep_.CoverTab[113309]++
															switch tok.value {
				case "key":
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:635
					_go_fuzz_dep_.CoverTab[113314]++
																if err := p.consumeToken(":"); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:636
						_go_fuzz_dep_.CoverTab[113321]++
																	return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:637
						// _ = "end of CoverTab[113321]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:638
						_go_fuzz_dep_.CoverTab[113322]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:638
						// _ = "end of CoverTab[113322]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:638
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:638
					// _ = "end of CoverTab[113314]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:638
					_go_fuzz_dep_.CoverTab[113315]++
																if err := p.readAny(key, props.MapKeyProp); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:639
						_go_fuzz_dep_.CoverTab[113323]++
																	return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:640
						// _ = "end of CoverTab[113323]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:641
						_go_fuzz_dep_.CoverTab[113324]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:641
						// _ = "end of CoverTab[113324]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:641
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:641
					// _ = "end of CoverTab[113315]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:641
					_go_fuzz_dep_.CoverTab[113316]++
																if err := p.consumeOptionalSeparator(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:642
						_go_fuzz_dep_.CoverTab[113325]++
																	return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:643
						// _ = "end of CoverTab[113325]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:644
						_go_fuzz_dep_.CoverTab[113326]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:644
						// _ = "end of CoverTab[113326]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:644
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:644
					// _ = "end of CoverTab[113316]"
				case "value":
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:645
					_go_fuzz_dep_.CoverTab[113317]++
																if err := p.checkForColon(props.MapValProp, dst.Type().Elem()); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:646
						_go_fuzz_dep_.CoverTab[113327]++
																	return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:647
						// _ = "end of CoverTab[113327]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:648
						_go_fuzz_dep_.CoverTab[113328]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:648
						// _ = "end of CoverTab[113328]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:648
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:648
					// _ = "end of CoverTab[113317]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:648
					_go_fuzz_dep_.CoverTab[113318]++
																if err := p.readAny(val, props.MapValProp); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:649
						_go_fuzz_dep_.CoverTab[113329]++
																	return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:650
						// _ = "end of CoverTab[113329]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:651
						_go_fuzz_dep_.CoverTab[113330]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:651
						// _ = "end of CoverTab[113330]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:651
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:651
					// _ = "end of CoverTab[113318]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:651
					_go_fuzz_dep_.CoverTab[113319]++
																if err := p.consumeOptionalSeparator(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:652
						_go_fuzz_dep_.CoverTab[113331]++
																	return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:653
						// _ = "end of CoverTab[113331]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:654
						_go_fuzz_dep_.CoverTab[113332]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:654
						// _ = "end of CoverTab[113332]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:654
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:654
					// _ = "end of CoverTab[113319]"
				default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:655
					_go_fuzz_dep_.CoverTab[113320]++
																p.back()
																return p.errorf(`expected "key", "value", or %q, found %q`, terminator, tok.value)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:657
					// _ = "end of CoverTab[113320]"
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:658
				// _ = "end of CoverTab[113309]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:659
			// _ = "end of CoverTab[113298]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:659
			_go_fuzz_dep_.CoverTab[113299]++

														dst.SetMapIndex(key, val)
														continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:662
			// _ = "end of CoverTab[113299]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:663
			_go_fuzz_dep_.CoverTab[113333]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:663
			// _ = "end of CoverTab[113333]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:663
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:663
		// _ = "end of CoverTab[113213]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:663
		_go_fuzz_dep_.CoverTab[113214]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:666
		if !props.Repeated && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:666
			_go_fuzz_dep_.CoverTab[113334]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:666
			return fieldSet[name]
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:666
			// _ = "end of CoverTab[113334]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:666
		}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:666
			_go_fuzz_dep_.CoverTab[113335]++
														return p.errorf("non-repeated field %q was repeated", name)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:667
			// _ = "end of CoverTab[113335]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:668
			_go_fuzz_dep_.CoverTab[113336]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:668
			// _ = "end of CoverTab[113336]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:668
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:668
		// _ = "end of CoverTab[113214]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:668
		_go_fuzz_dep_.CoverTab[113215]++

													if err := p.checkForColon(props, dst.Type()); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:670
			_go_fuzz_dep_.CoverTab[113337]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:671
			// _ = "end of CoverTab[113337]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:672
			_go_fuzz_dep_.CoverTab[113338]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:672
			// _ = "end of CoverTab[113338]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:672
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:672
		// _ = "end of CoverTab[113215]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:672
		_go_fuzz_dep_.CoverTab[113216]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:675
		fieldSet[name] = true
		if err := p.readAny(dst, props); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:676
			_go_fuzz_dep_.CoverTab[113339]++
														if _, ok := err.(*RequiredNotSetError); !ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:677
				_go_fuzz_dep_.CoverTab[113341]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:678
				// _ = "end of CoverTab[113341]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:679
				_go_fuzz_dep_.CoverTab[113342]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:679
				// _ = "end of CoverTab[113342]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:679
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:679
			// _ = "end of CoverTab[113339]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:679
			_go_fuzz_dep_.CoverTab[113340]++
														reqFieldErr = err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:680
			// _ = "end of CoverTab[113340]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:681
			_go_fuzz_dep_.CoverTab[113343]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:681
			// _ = "end of CoverTab[113343]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:681
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:681
		// _ = "end of CoverTab[113216]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:681
		_go_fuzz_dep_.CoverTab[113217]++
													if props.Required {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:682
			_go_fuzz_dep_.CoverTab[113344]++
														reqCount--
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:683
			// _ = "end of CoverTab[113344]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:684
			_go_fuzz_dep_.CoverTab[113345]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:684
			// _ = "end of CoverTab[113345]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:684
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:684
		// _ = "end of CoverTab[113217]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:684
		_go_fuzz_dep_.CoverTab[113218]++

													if err := p.consumeOptionalSeparator(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:686
			_go_fuzz_dep_.CoverTab[113346]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:687
			// _ = "end of CoverTab[113346]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:688
			_go_fuzz_dep_.CoverTab[113347]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:688
			// _ = "end of CoverTab[113347]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:688
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:688
		// _ = "end of CoverTab[113218]"

	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:690
	// _ = "end of CoverTab[113205]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:690
	_go_fuzz_dep_.CoverTab[113206]++

												if reqCount > 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:692
		_go_fuzz_dep_.CoverTab[113348]++
													return p.missingRequiredFieldError(sv)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:693
		// _ = "end of CoverTab[113348]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:694
		_go_fuzz_dep_.CoverTab[113349]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:694
		// _ = "end of CoverTab[113349]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:694
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:694
	// _ = "end of CoverTab[113206]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:694
	_go_fuzz_dep_.CoverTab[113207]++
												return reqFieldErr
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:695
	// _ = "end of CoverTab[113207]"
}

// consumeExtName consumes extension name or expanded Any type URL and the
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:698
// following ']'. It returns the name or URL consumed.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:700
func (p *textParser) consumeExtName() (string, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:700
	_go_fuzz_dep_.CoverTab[113350]++
												tok := p.next()
												if tok.err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:702
		_go_fuzz_dep_.CoverTab[113354]++
													return "", tok.err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:703
		// _ = "end of CoverTab[113354]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:704
		_go_fuzz_dep_.CoverTab[113355]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:704
		// _ = "end of CoverTab[113355]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:704
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:704
	// _ = "end of CoverTab[113350]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:704
	_go_fuzz_dep_.CoverTab[113351]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:707
	if len(tok.value) > 2 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:707
		_go_fuzz_dep_.CoverTab[113356]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:707
		return isQuote(tok.value[0])
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:707
		// _ = "end of CoverTab[113356]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:707
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:707
		_go_fuzz_dep_.CoverTab[113357]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:707
		return tok.value[len(tok.value)-1] == tok.value[0]
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:707
		// _ = "end of CoverTab[113357]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:707
	}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:707
		_go_fuzz_dep_.CoverTab[113358]++
													name, err := unquoteC(tok.value[1:len(tok.value)-1], rune(tok.value[0]))
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:709
			_go_fuzz_dep_.CoverTab[113360]++
														return "", err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:710
			// _ = "end of CoverTab[113360]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:711
			_go_fuzz_dep_.CoverTab[113361]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:711
			// _ = "end of CoverTab[113361]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:711
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:711
		// _ = "end of CoverTab[113358]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:711
		_go_fuzz_dep_.CoverTab[113359]++
													return name, p.consumeToken("]")
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:712
		// _ = "end of CoverTab[113359]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:713
		_go_fuzz_dep_.CoverTab[113362]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:713
		// _ = "end of CoverTab[113362]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:713
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:713
	// _ = "end of CoverTab[113351]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:713
	_go_fuzz_dep_.CoverTab[113352]++

	// Consume everything up to "]"
	var parts []string
	for tok.value != "]" {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:717
		_go_fuzz_dep_.CoverTab[113363]++
													parts = append(parts, tok.value)
													tok = p.next()
													if tok.err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:720
			_go_fuzz_dep_.CoverTab[113365]++
														return "", p.errorf("unrecognized type_url or extension name: %s", tok.err)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:721
			// _ = "end of CoverTab[113365]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:722
			_go_fuzz_dep_.CoverTab[113366]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:722
			// _ = "end of CoverTab[113366]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:722
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:722
		// _ = "end of CoverTab[113363]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:722
		_go_fuzz_dep_.CoverTab[113364]++
													if p.done && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:723
			_go_fuzz_dep_.CoverTab[113367]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:723
			return tok.value != "]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:723
			// _ = "end of CoverTab[113367]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:723
		}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:723
			_go_fuzz_dep_.CoverTab[113368]++
														return "", p.errorf("unclosed type_url or extension name")
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:724
			// _ = "end of CoverTab[113368]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:725
			_go_fuzz_dep_.CoverTab[113369]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:725
			// _ = "end of CoverTab[113369]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:725
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:725
		// _ = "end of CoverTab[113364]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:726
	// _ = "end of CoverTab[113352]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:726
	_go_fuzz_dep_.CoverTab[113353]++
												return strings.Join(parts, ""), nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:727
	// _ = "end of CoverTab[113353]"
}

// consumeOptionalSeparator consumes an optional semicolon or comma.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:730
// It is used in readStruct to provide backward compatibility.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:732
func (p *textParser) consumeOptionalSeparator() error {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:732
	_go_fuzz_dep_.CoverTab[113370]++
												tok := p.next()
												if tok.err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:734
		_go_fuzz_dep_.CoverTab[113373]++
													return tok.err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:735
		// _ = "end of CoverTab[113373]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:736
		_go_fuzz_dep_.CoverTab[113374]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:736
		// _ = "end of CoverTab[113374]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:736
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:736
	// _ = "end of CoverTab[113370]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:736
	_go_fuzz_dep_.CoverTab[113371]++
												if tok.value != ";" && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:737
		_go_fuzz_dep_.CoverTab[113375]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:737
		return tok.value != ","
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:737
		// _ = "end of CoverTab[113375]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:737
	}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:737
		_go_fuzz_dep_.CoverTab[113376]++
													p.back()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:738
		// _ = "end of CoverTab[113376]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:739
		_go_fuzz_dep_.CoverTab[113377]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:739
		// _ = "end of CoverTab[113377]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:739
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:739
	// _ = "end of CoverTab[113371]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:739
	_go_fuzz_dep_.CoverTab[113372]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:740
	// _ = "end of CoverTab[113372]"
}

func (p *textParser) readAny(v reflect.Value, props *Properties) error {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:743
	_go_fuzz_dep_.CoverTab[113378]++
												tok := p.next()
												if tok.err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:745
		_go_fuzz_dep_.CoverTab[113385]++
													return tok.err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:746
		// _ = "end of CoverTab[113385]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:747
		_go_fuzz_dep_.CoverTab[113386]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:747
		// _ = "end of CoverTab[113386]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:747
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:747
	// _ = "end of CoverTab[113378]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:747
	_go_fuzz_dep_.CoverTab[113379]++
												if tok.value == "" {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:748
		_go_fuzz_dep_.CoverTab[113387]++
													return p.errorf("unexpected EOF")
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:749
		// _ = "end of CoverTab[113387]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:750
		_go_fuzz_dep_.CoverTab[113388]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:750
		// _ = "end of CoverTab[113388]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:750
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:750
	// _ = "end of CoverTab[113379]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:750
	_go_fuzz_dep_.CoverTab[113380]++
												if len(props.CustomType) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:751
		_go_fuzz_dep_.CoverTab[113389]++
													if props.Repeated {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:752
			_go_fuzz_dep_.CoverTab[113392]++
														t := reflect.TypeOf(v.Interface())
														if t.Kind() == reflect.Slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:754
				_go_fuzz_dep_.CoverTab[113393]++
															tc := reflect.TypeOf(new(Marshaler))
															ok := t.Elem().Implements(tc.Elem())
															if ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:757
					_go_fuzz_dep_.CoverTab[113394]++
																fv := v
																flen := fv.Len()
																if flen == fv.Cap() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:760
						_go_fuzz_dep_.CoverTab[113396]++
																	nav := reflect.MakeSlice(v.Type(), flen, 2*flen+1)
																	reflect.Copy(nav, fv)
																	fv.Set(nav)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:763
						// _ = "end of CoverTab[113396]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:764
						_go_fuzz_dep_.CoverTab[113397]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:764
						// _ = "end of CoverTab[113397]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:764
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:764
					// _ = "end of CoverTab[113394]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:764
					_go_fuzz_dep_.CoverTab[113395]++
																fv.SetLen(flen + 1)

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:768
					p.back()
																return p.readAny(fv.Index(flen), props)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:769
					// _ = "end of CoverTab[113395]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:770
					_go_fuzz_dep_.CoverTab[113398]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:770
					// _ = "end of CoverTab[113398]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:770
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:770
				// _ = "end of CoverTab[113393]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:771
				_go_fuzz_dep_.CoverTab[113399]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:771
				// _ = "end of CoverTab[113399]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:771
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:771
			// _ = "end of CoverTab[113392]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:772
			_go_fuzz_dep_.CoverTab[113400]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:772
			// _ = "end of CoverTab[113400]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:772
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:772
		// _ = "end of CoverTab[113389]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:772
		_go_fuzz_dep_.CoverTab[113390]++
													if reflect.TypeOf(v.Interface()).Kind() == reflect.Ptr {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:773
			_go_fuzz_dep_.CoverTab[113401]++
														custom := reflect.New(props.ctype.Elem()).Interface().(Unmarshaler)
														err := custom.Unmarshal([]byte(tok.unquoted))
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:776
				_go_fuzz_dep_.CoverTab[113403]++
															return p.errorf("%v %v: %v", err, v.Type(), tok.value)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:777
				// _ = "end of CoverTab[113403]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:778
				_go_fuzz_dep_.CoverTab[113404]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:778
				// _ = "end of CoverTab[113404]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:778
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:778
			// _ = "end of CoverTab[113401]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:778
			_go_fuzz_dep_.CoverTab[113402]++
														v.Set(reflect.ValueOf(custom))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:779
			// _ = "end of CoverTab[113402]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:780
			_go_fuzz_dep_.CoverTab[113405]++
														custom := reflect.New(reflect.TypeOf(v.Interface())).Interface().(Unmarshaler)
														err := custom.Unmarshal([]byte(tok.unquoted))
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:783
				_go_fuzz_dep_.CoverTab[113407]++
															return p.errorf("%v %v: %v", err, v.Type(), tok.value)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:784
				// _ = "end of CoverTab[113407]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:785
				_go_fuzz_dep_.CoverTab[113408]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:785
				// _ = "end of CoverTab[113408]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:785
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:785
			// _ = "end of CoverTab[113405]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:785
			_go_fuzz_dep_.CoverTab[113406]++
														v.Set(reflect.Indirect(reflect.ValueOf(custom)))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:786
			// _ = "end of CoverTab[113406]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:787
		// _ = "end of CoverTab[113390]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:787
		_go_fuzz_dep_.CoverTab[113391]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:788
		// _ = "end of CoverTab[113391]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:789
		_go_fuzz_dep_.CoverTab[113409]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:789
		// _ = "end of CoverTab[113409]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:789
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:789
	// _ = "end of CoverTab[113380]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:789
	_go_fuzz_dep_.CoverTab[113381]++
												if props.StdTime {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:790
		_go_fuzz_dep_.CoverTab[113410]++
													fv := v
													p.back()
													props.StdTime = false
													tproto := &timestamp{}
													err := p.readAny(reflect.ValueOf(tproto).Elem(), props)
													props.StdTime = true
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:797
			_go_fuzz_dep_.CoverTab[113415]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:798
			// _ = "end of CoverTab[113415]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:799
			_go_fuzz_dep_.CoverTab[113416]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:799
			// _ = "end of CoverTab[113416]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:799
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:799
		// _ = "end of CoverTab[113410]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:799
		_go_fuzz_dep_.CoverTab[113411]++
													tim, err := timestampFromProto(tproto)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:801
			_go_fuzz_dep_.CoverTab[113417]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:802
			// _ = "end of CoverTab[113417]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:803
			_go_fuzz_dep_.CoverTab[113418]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:803
			// _ = "end of CoverTab[113418]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:803
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:803
		// _ = "end of CoverTab[113411]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:803
		_go_fuzz_dep_.CoverTab[113412]++
													if props.Repeated {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:804
			_go_fuzz_dep_.CoverTab[113419]++
														t := reflect.TypeOf(v.Interface())
														if t.Kind() == reflect.Slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:806
				_go_fuzz_dep_.CoverTab[113420]++
															if t.Elem().Kind() == reflect.Ptr {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:807
					_go_fuzz_dep_.CoverTab[113421]++
																ts := fv.Interface().([]*time.Time)
																ts = append(ts, &tim)
																fv.Set(reflect.ValueOf(ts))
																return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:811
					// _ = "end of CoverTab[113421]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:812
					_go_fuzz_dep_.CoverTab[113422]++
																ts := fv.Interface().([]time.Time)
																ts = append(ts, tim)
																fv.Set(reflect.ValueOf(ts))
																return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:816
					// _ = "end of CoverTab[113422]"
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:817
				// _ = "end of CoverTab[113420]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:818
				_go_fuzz_dep_.CoverTab[113423]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:818
				// _ = "end of CoverTab[113423]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:818
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:818
			// _ = "end of CoverTab[113419]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:819
			_go_fuzz_dep_.CoverTab[113424]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:819
			// _ = "end of CoverTab[113424]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:819
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:819
		// _ = "end of CoverTab[113412]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:819
		_go_fuzz_dep_.CoverTab[113413]++
													if reflect.TypeOf(v.Interface()).Kind() == reflect.Ptr {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:820
			_go_fuzz_dep_.CoverTab[113425]++
														v.Set(reflect.ValueOf(&tim))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:821
			// _ = "end of CoverTab[113425]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:822
			_go_fuzz_dep_.CoverTab[113426]++
														v.Set(reflect.Indirect(reflect.ValueOf(&tim)))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:823
			// _ = "end of CoverTab[113426]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:824
		// _ = "end of CoverTab[113413]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:824
		_go_fuzz_dep_.CoverTab[113414]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:825
		// _ = "end of CoverTab[113414]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:826
		_go_fuzz_dep_.CoverTab[113427]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:826
		// _ = "end of CoverTab[113427]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:826
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:826
	// _ = "end of CoverTab[113381]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:826
	_go_fuzz_dep_.CoverTab[113382]++
												if props.StdDuration {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:827
		_go_fuzz_dep_.CoverTab[113428]++
													fv := v
													p.back()
													props.StdDuration = false
													dproto := &duration{}
													err := p.readAny(reflect.ValueOf(dproto).Elem(), props)
													props.StdDuration = true
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:834
			_go_fuzz_dep_.CoverTab[113433]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:835
			// _ = "end of CoverTab[113433]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:836
			_go_fuzz_dep_.CoverTab[113434]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:836
			// _ = "end of CoverTab[113434]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:836
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:836
		// _ = "end of CoverTab[113428]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:836
		_go_fuzz_dep_.CoverTab[113429]++
													dur, err := durationFromProto(dproto)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:838
			_go_fuzz_dep_.CoverTab[113435]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:839
			// _ = "end of CoverTab[113435]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:840
			_go_fuzz_dep_.CoverTab[113436]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:840
			// _ = "end of CoverTab[113436]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:840
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:840
		// _ = "end of CoverTab[113429]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:840
		_go_fuzz_dep_.CoverTab[113430]++
													if props.Repeated {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:841
			_go_fuzz_dep_.CoverTab[113437]++
														t := reflect.TypeOf(v.Interface())
														if t.Kind() == reflect.Slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:843
				_go_fuzz_dep_.CoverTab[113438]++
															if t.Elem().Kind() == reflect.Ptr {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:844
					_go_fuzz_dep_.CoverTab[113439]++
																ds := fv.Interface().([]*time.Duration)
																ds = append(ds, &dur)
																fv.Set(reflect.ValueOf(ds))
																return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:848
					// _ = "end of CoverTab[113439]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:849
					_go_fuzz_dep_.CoverTab[113440]++
																ds := fv.Interface().([]time.Duration)
																ds = append(ds, dur)
																fv.Set(reflect.ValueOf(ds))
																return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:853
					// _ = "end of CoverTab[113440]"
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:854
				// _ = "end of CoverTab[113438]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:855
				_go_fuzz_dep_.CoverTab[113441]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:855
				// _ = "end of CoverTab[113441]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:855
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:855
			// _ = "end of CoverTab[113437]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:856
			_go_fuzz_dep_.CoverTab[113442]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:856
			// _ = "end of CoverTab[113442]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:856
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:856
		// _ = "end of CoverTab[113430]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:856
		_go_fuzz_dep_.CoverTab[113431]++
													if reflect.TypeOf(v.Interface()).Kind() == reflect.Ptr {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:857
			_go_fuzz_dep_.CoverTab[113443]++
														v.Set(reflect.ValueOf(&dur))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:858
			// _ = "end of CoverTab[113443]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:859
			_go_fuzz_dep_.CoverTab[113444]++
														v.Set(reflect.Indirect(reflect.ValueOf(&dur)))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:860
			// _ = "end of CoverTab[113444]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:861
		// _ = "end of CoverTab[113431]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:861
		_go_fuzz_dep_.CoverTab[113432]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:862
		// _ = "end of CoverTab[113432]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:863
		_go_fuzz_dep_.CoverTab[113445]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:863
		// _ = "end of CoverTab[113445]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:863
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:863
	// _ = "end of CoverTab[113382]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:863
	_go_fuzz_dep_.CoverTab[113383]++
												switch fv := v; fv.Kind() {
	case reflect.Slice:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:865
		_go_fuzz_dep_.CoverTab[113446]++
													at := v.Type()
													if at.Elem().Kind() == reflect.Uint8 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:867
			_go_fuzz_dep_.CoverTab[113469]++

														if tok.value[0] != '"' && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:869
				_go_fuzz_dep_.CoverTab[113471]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:869
				return tok.value[0] != '\''
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:869
				// _ = "end of CoverTab[113471]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:869
			}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:869
				_go_fuzz_dep_.CoverTab[113472]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:873
				return p.errorf("invalid string: %v", tok.value)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:873
				// _ = "end of CoverTab[113472]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:874
				_go_fuzz_dep_.CoverTab[113473]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:874
				// _ = "end of CoverTab[113473]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:874
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:874
			// _ = "end of CoverTab[113469]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:874
			_go_fuzz_dep_.CoverTab[113470]++
														bytes := []byte(tok.unquoted)
														fv.Set(reflect.ValueOf(bytes))
														return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:877
			// _ = "end of CoverTab[113470]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:878
			_go_fuzz_dep_.CoverTab[113474]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:878
			// _ = "end of CoverTab[113474]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:878
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:878
		// _ = "end of CoverTab[113446]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:878
		_go_fuzz_dep_.CoverTab[113447]++

													if tok.value == "[" {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:880
			_go_fuzz_dep_.CoverTab[113475]++

														for {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:882
				_go_fuzz_dep_.CoverTab[113477]++
															fv.Set(reflect.Append(fv, reflect.New(at.Elem()).Elem()))
															err := p.readAny(fv.Index(fv.Len()-1), props)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:885
					_go_fuzz_dep_.CoverTab[113481]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:886
					// _ = "end of CoverTab[113481]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:887
					_go_fuzz_dep_.CoverTab[113482]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:887
					// _ = "end of CoverTab[113482]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:887
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:887
				// _ = "end of CoverTab[113477]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:887
				_go_fuzz_dep_.CoverTab[113478]++
															ntok := p.next()
															if ntok.err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:889
					_go_fuzz_dep_.CoverTab[113483]++
																return ntok.err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:890
					// _ = "end of CoverTab[113483]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:891
					_go_fuzz_dep_.CoverTab[113484]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:891
					// _ = "end of CoverTab[113484]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:891
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:891
				// _ = "end of CoverTab[113478]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:891
				_go_fuzz_dep_.CoverTab[113479]++
															if ntok.value == "]" {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:892
					_go_fuzz_dep_.CoverTab[113485]++
																break
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:893
					// _ = "end of CoverTab[113485]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:894
					_go_fuzz_dep_.CoverTab[113486]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:894
					// _ = "end of CoverTab[113486]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:894
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:894
				// _ = "end of CoverTab[113479]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:894
				_go_fuzz_dep_.CoverTab[113480]++
															if ntok.value != "," {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:895
					_go_fuzz_dep_.CoverTab[113487]++
																return p.errorf("Expected ']' or ',' found %q", ntok.value)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:896
					// _ = "end of CoverTab[113487]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:897
					_go_fuzz_dep_.CoverTab[113488]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:897
					// _ = "end of CoverTab[113488]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:897
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:897
				// _ = "end of CoverTab[113480]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:898
			// _ = "end of CoverTab[113475]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:898
			_go_fuzz_dep_.CoverTab[113476]++
														return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:899
			// _ = "end of CoverTab[113476]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:900
			_go_fuzz_dep_.CoverTab[113489]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:900
			// _ = "end of CoverTab[113489]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:900
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:900
		// _ = "end of CoverTab[113447]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:900
		_go_fuzz_dep_.CoverTab[113448]++

													p.back()
													fv.Set(reflect.Append(fv, reflect.New(at.Elem()).Elem()))
													return p.readAny(fv.Index(fv.Len()-1), props)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:904
		// _ = "end of CoverTab[113448]"
	case reflect.Bool:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:905
		_go_fuzz_dep_.CoverTab[113449]++

													switch tok.value {
		case "true", "1", "t", "True":
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:908
			_go_fuzz_dep_.CoverTab[113490]++
														fv.SetBool(true)
														return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:910
			// _ = "end of CoverTab[113490]"
		case "false", "0", "f", "False":
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:911
			_go_fuzz_dep_.CoverTab[113491]++
														fv.SetBool(false)
														return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:913
			// _ = "end of CoverTab[113491]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:913
		default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:913
			_go_fuzz_dep_.CoverTab[113492]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:913
			// _ = "end of CoverTab[113492]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:914
		// _ = "end of CoverTab[113449]"
	case reflect.Float32, reflect.Float64:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:915
		_go_fuzz_dep_.CoverTab[113450]++
													v := tok.value

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:919
		if strings.HasSuffix(v, "f") && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:919
			_go_fuzz_dep_.CoverTab[113493]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:919
			return tok.value != "-inf"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:919
			// _ = "end of CoverTab[113493]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:919
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:919
			_go_fuzz_dep_.CoverTab[113494]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:919
			return tok.value != "inf"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:919
			// _ = "end of CoverTab[113494]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:919
		}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:919
			_go_fuzz_dep_.CoverTab[113495]++
														v = v[:len(v)-1]
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:920
			// _ = "end of CoverTab[113495]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:921
			_go_fuzz_dep_.CoverTab[113496]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:921
			// _ = "end of CoverTab[113496]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:921
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:921
		// _ = "end of CoverTab[113450]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:921
		_go_fuzz_dep_.CoverTab[113451]++
													if f, err := strconv.ParseFloat(v, fv.Type().Bits()); err == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:922
			_go_fuzz_dep_.CoverTab[113497]++
														fv.SetFloat(f)
														return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:924
			// _ = "end of CoverTab[113497]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:925
			_go_fuzz_dep_.CoverTab[113498]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:925
			// _ = "end of CoverTab[113498]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:925
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:925
		// _ = "end of CoverTab[113451]"
	case reflect.Int8:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:926
		_go_fuzz_dep_.CoverTab[113452]++
													if x, err := strconv.ParseInt(tok.value, 0, 8); err == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:927
			_go_fuzz_dep_.CoverTab[113499]++
														fv.SetInt(x)
														return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:929
			// _ = "end of CoverTab[113499]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:930
			_go_fuzz_dep_.CoverTab[113500]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:930
			// _ = "end of CoverTab[113500]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:930
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:930
		// _ = "end of CoverTab[113452]"
	case reflect.Int16:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:931
		_go_fuzz_dep_.CoverTab[113453]++
													if x, err := strconv.ParseInt(tok.value, 0, 16); err == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:932
			_go_fuzz_dep_.CoverTab[113501]++
														fv.SetInt(x)
														return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:934
			// _ = "end of CoverTab[113501]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:935
			_go_fuzz_dep_.CoverTab[113502]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:935
			// _ = "end of CoverTab[113502]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:935
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:935
		// _ = "end of CoverTab[113453]"
	case reflect.Int32:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:936
		_go_fuzz_dep_.CoverTab[113454]++
													if x, err := strconv.ParseInt(tok.value, 0, 32); err == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:937
			_go_fuzz_dep_.CoverTab[113503]++
														fv.SetInt(x)
														return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:939
			// _ = "end of CoverTab[113503]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:940
			_go_fuzz_dep_.CoverTab[113504]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:940
			// _ = "end of CoverTab[113504]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:940
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:940
		// _ = "end of CoverTab[113454]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:940
		_go_fuzz_dep_.CoverTab[113455]++

													if len(props.Enum) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:942
			_go_fuzz_dep_.CoverTab[113505]++
														break
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:943
			// _ = "end of CoverTab[113505]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:944
			_go_fuzz_dep_.CoverTab[113506]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:944
			// _ = "end of CoverTab[113506]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:944
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:944
		// _ = "end of CoverTab[113455]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:944
		_go_fuzz_dep_.CoverTab[113456]++
													m, ok := enumValueMaps[props.Enum]
													if !ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:946
			_go_fuzz_dep_.CoverTab[113507]++
														break
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:947
			// _ = "end of CoverTab[113507]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:948
			_go_fuzz_dep_.CoverTab[113508]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:948
			// _ = "end of CoverTab[113508]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:948
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:948
		// _ = "end of CoverTab[113456]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:948
		_go_fuzz_dep_.CoverTab[113457]++
													x, ok := m[tok.value]
													if !ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:950
			_go_fuzz_dep_.CoverTab[113509]++
														break
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:951
			// _ = "end of CoverTab[113509]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:952
			_go_fuzz_dep_.CoverTab[113510]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:952
			// _ = "end of CoverTab[113510]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:952
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:952
		// _ = "end of CoverTab[113457]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:952
		_go_fuzz_dep_.CoverTab[113458]++
													fv.SetInt(int64(x))
													return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:954
		// _ = "end of CoverTab[113458]"
	case reflect.Int64:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:955
		_go_fuzz_dep_.CoverTab[113459]++
													if x, err := strconv.ParseInt(tok.value, 0, 64); err == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:956
			_go_fuzz_dep_.CoverTab[113511]++
														fv.SetInt(x)
														return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:958
			// _ = "end of CoverTab[113511]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:959
			_go_fuzz_dep_.CoverTab[113512]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:959
			// _ = "end of CoverTab[113512]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:959
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:959
		// _ = "end of CoverTab[113459]"

	case reflect.Ptr:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:961
		_go_fuzz_dep_.CoverTab[113460]++

													p.back()
													fv.Set(reflect.New(fv.Type().Elem()))
													return p.readAny(fv.Elem(), props)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:965
		// _ = "end of CoverTab[113460]"
	case reflect.String:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:966
		_go_fuzz_dep_.CoverTab[113461]++
													if tok.value[0] == '"' || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:967
			_go_fuzz_dep_.CoverTab[113513]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:967
			return tok.value[0] == '\''
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:967
			// _ = "end of CoverTab[113513]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:967
		}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:967
			_go_fuzz_dep_.CoverTab[113514]++
														fv.SetString(tok.unquoted)
														return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:969
			// _ = "end of CoverTab[113514]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:970
			_go_fuzz_dep_.CoverTab[113515]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:970
			// _ = "end of CoverTab[113515]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:970
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:970
		// _ = "end of CoverTab[113461]"
	case reflect.Struct:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:971
		_go_fuzz_dep_.CoverTab[113462]++
													var terminator string
													switch tok.value {
		case "{":
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:974
			_go_fuzz_dep_.CoverTab[113516]++
														terminator = "}"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:975
			// _ = "end of CoverTab[113516]"
		case "<":
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:976
			_go_fuzz_dep_.CoverTab[113517]++
														terminator = ">"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:977
			// _ = "end of CoverTab[113517]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:978
			_go_fuzz_dep_.CoverTab[113518]++
														return p.errorf("expected '{' or '<', found %q", tok.value)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:979
			// _ = "end of CoverTab[113518]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:980
		// _ = "end of CoverTab[113462]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:980
		_go_fuzz_dep_.CoverTab[113463]++

													return p.readStruct(fv, terminator)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:982
		// _ = "end of CoverTab[113463]"
	case reflect.Uint8:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:983
		_go_fuzz_dep_.CoverTab[113464]++
													if x, err := strconv.ParseUint(tok.value, 0, 8); err == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:984
			_go_fuzz_dep_.CoverTab[113519]++
														fv.SetUint(x)
														return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:986
			// _ = "end of CoverTab[113519]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:987
			_go_fuzz_dep_.CoverTab[113520]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:987
			// _ = "end of CoverTab[113520]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:987
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:987
		// _ = "end of CoverTab[113464]"
	case reflect.Uint16:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:988
		_go_fuzz_dep_.CoverTab[113465]++
													if x, err := strconv.ParseUint(tok.value, 0, 16); err == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:989
			_go_fuzz_dep_.CoverTab[113521]++
														fv.SetUint(x)
														return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:991
			// _ = "end of CoverTab[113521]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:992
			_go_fuzz_dep_.CoverTab[113522]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:992
			// _ = "end of CoverTab[113522]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:992
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:992
		// _ = "end of CoverTab[113465]"
	case reflect.Uint32:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:993
		_go_fuzz_dep_.CoverTab[113466]++
													if x, err := strconv.ParseUint(tok.value, 0, 32); err == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:994
			_go_fuzz_dep_.CoverTab[113523]++
														fv.SetUint(uint64(x))
														return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:996
			// _ = "end of CoverTab[113523]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:997
			_go_fuzz_dep_.CoverTab[113524]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:997
			// _ = "end of CoverTab[113524]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:997
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:997
		// _ = "end of CoverTab[113466]"
	case reflect.Uint64:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:998
		_go_fuzz_dep_.CoverTab[113467]++
													if x, err := strconv.ParseUint(tok.value, 0, 64); err == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:999
			_go_fuzz_dep_.CoverTab[113525]++
														fv.SetUint(x)
														return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:1001
			// _ = "end of CoverTab[113525]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:1002
			_go_fuzz_dep_.CoverTab[113526]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:1002
			// _ = "end of CoverTab[113526]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:1002
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:1002
		// _ = "end of CoverTab[113467]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:1002
	default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:1002
		_go_fuzz_dep_.CoverTab[113468]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:1002
		// _ = "end of CoverTab[113468]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:1003
	// _ = "end of CoverTab[113383]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:1003
	_go_fuzz_dep_.CoverTab[113384]++
												return p.errorf("invalid %v: %v", v.Type(), tok.value)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:1004
	// _ = "end of CoverTab[113384]"
}

// UnmarshalText reads a protocol buffer in Text format. UnmarshalText resets pb
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:1007
// before starting to unmarshal, so any existing data in pb is always removed.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:1007
// If a required field is not set and no other error occurs,
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:1007
// UnmarshalText returns *RequiredNotSetError.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:1011
func UnmarshalText(s string, pb Message) error {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:1011
	_go_fuzz_dep_.CoverTab[113527]++
												if um, ok := pb.(encoding.TextUnmarshaler); ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:1012
		_go_fuzz_dep_.CoverTab[113529]++
													return um.UnmarshalText([]byte(s))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:1013
		// _ = "end of CoverTab[113529]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:1014
		_go_fuzz_dep_.CoverTab[113530]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:1014
		// _ = "end of CoverTab[113530]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:1014
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:1014
	// _ = "end of CoverTab[113527]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:1014
	_go_fuzz_dep_.CoverTab[113528]++
												pb.Reset()
												v := reflect.ValueOf(pb)
												return newTextParser(s).readStruct(v.Elem(), "")
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:1017
	// _ = "end of CoverTab[113528]"
}

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:1018
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_parser.go:1018
var _ = _go_fuzz_dep_.CoverTab
