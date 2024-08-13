// Copyright 2015 Unknwon
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:15
package ini

//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:15
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:15
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:15
)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:15
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:15
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:15
)

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

const minReaderBufferSize = 4096

var pythonMultiline = regexp.MustCompile(`^([\t\f ]+)(.*)`)

type parserOptions struct {
	IgnoreContinuation		bool
	IgnoreInlineComment		bool
	AllowPythonMultilineValues	bool
	SpaceBeforeInlineComment	bool
	UnescapeValueDoubleQuotes	bool
	UnescapeValueCommentSymbols	bool
	PreserveSurroundedQuote		bool
	DebugFunc			DebugFunc
	ReaderBufferSize		int
}

type parser struct {
	buf	*bufio.Reader
	options	parserOptions

	isEOF	bool
	count	int
	comment	*bytes.Buffer
}

func (p *parser) debug(format string, args ...interface{}) {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:53
	_go_fuzz_dep_.CoverTab[128809]++
									if p.options.DebugFunc != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:54
		_go_fuzz_dep_.CoverTab[128810]++
										p.options.DebugFunc(fmt.Sprintf(format, args...))
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:55
		// _ = "end of CoverTab[128810]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:56
		_go_fuzz_dep_.CoverTab[128811]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:56
		// _ = "end of CoverTab[128811]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:56
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:56
	// _ = "end of CoverTab[128809]"
}

func newParser(r io.Reader, opts parserOptions) *parser {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:59
	_go_fuzz_dep_.CoverTab[128812]++
									size := opts.ReaderBufferSize
									if size < minReaderBufferSize {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:61
		_go_fuzz_dep_.CoverTab[128814]++
										size = minReaderBufferSize
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:62
		// _ = "end of CoverTab[128814]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:63
		_go_fuzz_dep_.CoverTab[128815]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:63
		// _ = "end of CoverTab[128815]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:63
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:63
	// _ = "end of CoverTab[128812]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:63
	_go_fuzz_dep_.CoverTab[128813]++

									return &parser{
		buf:		bufio.NewReaderSize(r, size),
		options:	opts,
		count:		1,
		comment:	&bytes.Buffer{},
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:70
	// _ = "end of CoverTab[128813]"
}

// BOM handles header of UTF-8, UTF-16 LE and UTF-16 BE's BOM format.
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:73
// http://en.wikipedia.org/wiki/Byte_order_mark#Representations_of_byte_order_marks_by_encoding
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:75
func (p *parser) BOM() error {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:75
	_go_fuzz_dep_.CoverTab[128816]++
									mask, err := p.buf.Peek(2)
									if err != nil && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:77
		_go_fuzz_dep_.CoverTab[128819]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:77
		return err != io.EOF
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:77
		// _ = "end of CoverTab[128819]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:77
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:77
		_go_fuzz_dep_.CoverTab[128820]++
										return err
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:78
		// _ = "end of CoverTab[128820]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:79
		_go_fuzz_dep_.CoverTab[128821]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:79
		if len(mask) < 2 {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:79
			_go_fuzz_dep_.CoverTab[128822]++
											return nil
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:80
			// _ = "end of CoverTab[128822]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:81
			_go_fuzz_dep_.CoverTab[128823]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:81
			// _ = "end of CoverTab[128823]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:81
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:81
		// _ = "end of CoverTab[128821]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:81
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:81
	// _ = "end of CoverTab[128816]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:81
	_go_fuzz_dep_.CoverTab[128817]++

									switch {
	case mask[0] == 254 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:84
		_go_fuzz_dep_.CoverTab[128829]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:84
		return mask[1] == 255
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:84
		// _ = "end of CoverTab[128829]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:84
	}():
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:84
		_go_fuzz_dep_.CoverTab[128824]++
										fallthrough
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:85
		// _ = "end of CoverTab[128824]"
	case mask[0] == 255 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:86
		_go_fuzz_dep_.CoverTab[128830]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:86
		return mask[1] == 254
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:86
		// _ = "end of CoverTab[128830]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:86
	}():
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:86
		_go_fuzz_dep_.CoverTab[128825]++
										_, err = p.buf.Read(mask)
										if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:88
			_go_fuzz_dep_.CoverTab[128831]++
											return err
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:89
			// _ = "end of CoverTab[128831]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:90
			_go_fuzz_dep_.CoverTab[128832]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:90
			// _ = "end of CoverTab[128832]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:90
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:90
		// _ = "end of CoverTab[128825]"
	case mask[0] == 239 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:91
		_go_fuzz_dep_.CoverTab[128833]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:91
		return mask[1] == 187
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:91
		// _ = "end of CoverTab[128833]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:91
	}():
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:91
		_go_fuzz_dep_.CoverTab[128826]++
										mask, err := p.buf.Peek(3)
										if err != nil && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:93
			_go_fuzz_dep_.CoverTab[128834]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:93
			return err != io.EOF
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:93
			// _ = "end of CoverTab[128834]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:93
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:93
			_go_fuzz_dep_.CoverTab[128835]++
											return err
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:94
			// _ = "end of CoverTab[128835]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:95
			_go_fuzz_dep_.CoverTab[128836]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:95
			if len(mask) < 3 {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:95
				_go_fuzz_dep_.CoverTab[128837]++
												return nil
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:96
				// _ = "end of CoverTab[128837]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:97
				_go_fuzz_dep_.CoverTab[128838]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:97
				// _ = "end of CoverTab[128838]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:97
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:97
			// _ = "end of CoverTab[128836]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:97
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:97
		// _ = "end of CoverTab[128826]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:97
		_go_fuzz_dep_.CoverTab[128827]++
										if mask[2] == 191 {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:98
			_go_fuzz_dep_.CoverTab[128839]++
											_, err = p.buf.Read(mask)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:100
				_go_fuzz_dep_.CoverTab[128840]++
													return err
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:101
				// _ = "end of CoverTab[128840]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:102
				_go_fuzz_dep_.CoverTab[128841]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:102
				// _ = "end of CoverTab[128841]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:102
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:102
			// _ = "end of CoverTab[128839]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:103
			_go_fuzz_dep_.CoverTab[128842]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:103
			// _ = "end of CoverTab[128842]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:103
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:103
		// _ = "end of CoverTab[128827]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:103
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:103
		_go_fuzz_dep_.CoverTab[128828]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:103
		// _ = "end of CoverTab[128828]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:104
	// _ = "end of CoverTab[128817]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:104
	_go_fuzz_dep_.CoverTab[128818]++
										return nil
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:105
	// _ = "end of CoverTab[128818]"
}

func (p *parser) readUntil(delim byte) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:108
	_go_fuzz_dep_.CoverTab[128843]++
										data, err := p.buf.ReadBytes(delim)
										if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:110
		_go_fuzz_dep_.CoverTab[128845]++
											if err == io.EOF {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:111
			_go_fuzz_dep_.CoverTab[128846]++
												p.isEOF = true
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:112
			// _ = "end of CoverTab[128846]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:113
			_go_fuzz_dep_.CoverTab[128847]++
												return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:114
			// _ = "end of CoverTab[128847]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:115
		// _ = "end of CoverTab[128845]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:116
		_go_fuzz_dep_.CoverTab[128848]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:116
		// _ = "end of CoverTab[128848]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:116
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:116
	// _ = "end of CoverTab[128843]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:116
	_go_fuzz_dep_.CoverTab[128844]++
										return data, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:117
	// _ = "end of CoverTab[128844]"
}

func cleanComment(in []byte) ([]byte, bool) {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:120
	_go_fuzz_dep_.CoverTab[128849]++
										i := bytes.IndexAny(in, "#;")
										if i == -1 {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:122
		_go_fuzz_dep_.CoverTab[128851]++
											return nil, false
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:123
		// _ = "end of CoverTab[128851]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:124
		_go_fuzz_dep_.CoverTab[128852]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:124
		// _ = "end of CoverTab[128852]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:124
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:124
	// _ = "end of CoverTab[128849]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:124
	_go_fuzz_dep_.CoverTab[128850]++
										return in[i:], true
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:125
	// _ = "end of CoverTab[128850]"
}

func readKeyName(delimiters string, in []byte) (string, int, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:128
	_go_fuzz_dep_.CoverTab[128853]++
										line := string(in)

	// Check if key name surrounded by quotes.
	var keyQuote string
	if line[0] == '"' {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:133
		_go_fuzz_dep_.CoverTab[128857]++
											if len(line) > 6 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:134
			_go_fuzz_dep_.CoverTab[128858]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:134
			return line[0:3] == `"""`
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:134
			// _ = "end of CoverTab[128858]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:134
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:134
			_go_fuzz_dep_.CoverTab[128859]++
												keyQuote = `"""`
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:135
			// _ = "end of CoverTab[128859]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:136
			_go_fuzz_dep_.CoverTab[128860]++
												keyQuote = `"`
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:137
			// _ = "end of CoverTab[128860]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:138
		// _ = "end of CoverTab[128857]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:139
		_go_fuzz_dep_.CoverTab[128861]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:139
		if line[0] == '`' {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:139
			_go_fuzz_dep_.CoverTab[128862]++
												keyQuote = "`"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:140
			// _ = "end of CoverTab[128862]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:141
			_go_fuzz_dep_.CoverTab[128863]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:141
			// _ = "end of CoverTab[128863]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:141
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:141
		// _ = "end of CoverTab[128861]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:141
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:141
	// _ = "end of CoverTab[128853]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:141
	_go_fuzz_dep_.CoverTab[128854]++

	// Get out key name
	var endIdx int
	if len(keyQuote) > 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:145
		_go_fuzz_dep_.CoverTab[128864]++
											startIdx := len(keyQuote)

											pos := strings.Index(line[startIdx:], keyQuote)
											if pos == -1 {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:149
			_go_fuzz_dep_.CoverTab[128867]++
												return "", -1, fmt.Errorf("missing closing key quote: %s", line)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:150
			// _ = "end of CoverTab[128867]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:151
			_go_fuzz_dep_.CoverTab[128868]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:151
			// _ = "end of CoverTab[128868]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:151
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:151
		// _ = "end of CoverTab[128864]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:151
		_go_fuzz_dep_.CoverTab[128865]++
											pos += startIdx

//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:155
		i := strings.IndexAny(line[pos+startIdx:], delimiters)
		if i < 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:156
			_go_fuzz_dep_.CoverTab[128869]++
												return "", -1, ErrDelimiterNotFound{line}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:157
			// _ = "end of CoverTab[128869]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:158
			_go_fuzz_dep_.CoverTab[128870]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:158
			// _ = "end of CoverTab[128870]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:158
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:158
		// _ = "end of CoverTab[128865]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:158
		_go_fuzz_dep_.CoverTab[128866]++
											endIdx = pos + i
											return strings.TrimSpace(line[startIdx:pos]), endIdx + startIdx + 1, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:160
		// _ = "end of CoverTab[128866]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:161
		_go_fuzz_dep_.CoverTab[128871]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:161
		// _ = "end of CoverTab[128871]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:161
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:161
	// _ = "end of CoverTab[128854]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:161
	_go_fuzz_dep_.CoverTab[128855]++

										endIdx = strings.IndexAny(line, delimiters)
										if endIdx < 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:164
		_go_fuzz_dep_.CoverTab[128872]++
											return "", -1, ErrDelimiterNotFound{line}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:165
		// _ = "end of CoverTab[128872]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:166
		_go_fuzz_dep_.CoverTab[128873]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:166
		// _ = "end of CoverTab[128873]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:166
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:166
	// _ = "end of CoverTab[128855]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:166
	_go_fuzz_dep_.CoverTab[128856]++
										return strings.TrimSpace(line[0:endIdx]), endIdx + 1, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:167
	// _ = "end of CoverTab[128856]"
}

func (p *parser) readMultilines(line, val, valQuote string) (string, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:170
	_go_fuzz_dep_.CoverTab[128874]++
										for {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:171
		_go_fuzz_dep_.CoverTab[128876]++
											data, err := p.readUntil('\n')
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:173
			_go_fuzz_dep_.CoverTab[128879]++
												return "", err
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:174
			// _ = "end of CoverTab[128879]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:175
			_go_fuzz_dep_.CoverTab[128880]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:175
			// _ = "end of CoverTab[128880]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:175
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:175
		// _ = "end of CoverTab[128876]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:175
		_go_fuzz_dep_.CoverTab[128877]++
											next := string(data)

											pos := strings.LastIndex(next, valQuote)
											if pos > -1 {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:179
			_go_fuzz_dep_.CoverTab[128881]++
												val += next[:pos]

												comment, has := cleanComment([]byte(next[pos:]))
												if has {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:183
				_go_fuzz_dep_.CoverTab[128883]++
													p.comment.Write(bytes.TrimSpace(comment))
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:184
				// _ = "end of CoverTab[128883]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:185
				_go_fuzz_dep_.CoverTab[128884]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:185
				// _ = "end of CoverTab[128884]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:185
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:185
			// _ = "end of CoverTab[128881]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:185
			_go_fuzz_dep_.CoverTab[128882]++
												break
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:186
			// _ = "end of CoverTab[128882]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:187
			_go_fuzz_dep_.CoverTab[128885]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:187
			// _ = "end of CoverTab[128885]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:187
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:187
		// _ = "end of CoverTab[128877]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:187
		_go_fuzz_dep_.CoverTab[128878]++
											val += next
											if p.isEOF {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:189
			_go_fuzz_dep_.CoverTab[128886]++
												return "", fmt.Errorf("missing closing key quote from %q to %q", line, next)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:190
			// _ = "end of CoverTab[128886]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:191
			_go_fuzz_dep_.CoverTab[128887]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:191
			// _ = "end of CoverTab[128887]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:191
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:191
		// _ = "end of CoverTab[128878]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:192
	// _ = "end of CoverTab[128874]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:192
	_go_fuzz_dep_.CoverTab[128875]++
										return val, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:193
	// _ = "end of CoverTab[128875]"
}

func (p *parser) readContinuationLines(val string) (string, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:196
	_go_fuzz_dep_.CoverTab[128888]++
										for {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:197
		_go_fuzz_dep_.CoverTab[128890]++
											data, err := p.readUntil('\n')
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:199
			_go_fuzz_dep_.CoverTab[128894]++
												return "", err
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:200
			// _ = "end of CoverTab[128894]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:201
			_go_fuzz_dep_.CoverTab[128895]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:201
			// _ = "end of CoverTab[128895]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:201
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:201
		// _ = "end of CoverTab[128890]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:201
		_go_fuzz_dep_.CoverTab[128891]++
											next := strings.TrimSpace(string(data))

											if len(next) == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:204
			_go_fuzz_dep_.CoverTab[128896]++
												break
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:205
			// _ = "end of CoverTab[128896]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:206
			_go_fuzz_dep_.CoverTab[128897]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:206
			// _ = "end of CoverTab[128897]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:206
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:206
		// _ = "end of CoverTab[128891]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:206
		_go_fuzz_dep_.CoverTab[128892]++
											val += next
											if val[len(val)-1] != '\\' {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:208
			_go_fuzz_dep_.CoverTab[128898]++
												break
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:209
			// _ = "end of CoverTab[128898]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:210
			_go_fuzz_dep_.CoverTab[128899]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:210
			// _ = "end of CoverTab[128899]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:210
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:210
		// _ = "end of CoverTab[128892]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:210
		_go_fuzz_dep_.CoverTab[128893]++
											val = val[:len(val)-1]
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:211
		// _ = "end of CoverTab[128893]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:212
	// _ = "end of CoverTab[128888]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:212
	_go_fuzz_dep_.CoverTab[128889]++
										return val, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:213
	// _ = "end of CoverTab[128889]"
}

// hasSurroundedQuote check if and only if the first and last characters
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:216
// are quotes \" or \'.
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:216
// It returns false if any other parts also contain same kind of quotes.
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:219
func hasSurroundedQuote(in string, quote byte) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:219
	_go_fuzz_dep_.CoverTab[128900]++
										return len(in) >= 2 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:220
		_go_fuzz_dep_.CoverTab[128901]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:220
		return in[0] == quote
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:220
		// _ = "end of CoverTab[128901]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:220
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:220
		_go_fuzz_dep_.CoverTab[128902]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:220
		return in[len(in)-1] == quote
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:220
		// _ = "end of CoverTab[128902]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:220
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:220
		_go_fuzz_dep_.CoverTab[128903]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:220
		return strings.IndexByte(in[1:], quote) == len(in)-2
											// _ = "end of CoverTab[128903]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:221
	}()
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:221
	// _ = "end of CoverTab[128900]"
}

func (p *parser) readValue(in []byte, bufferSize int) (string, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:224
	_go_fuzz_dep_.CoverTab[128904]++

										line := strings.TrimLeftFunc(string(in), unicode.IsSpace)
										if len(line) == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:227
		_go_fuzz_dep_.CoverTab[128911]++
											if p.options.AllowPythonMultilineValues && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:228
			_go_fuzz_dep_.CoverTab[128913]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:228
			return len(in) > 0
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:228
			// _ = "end of CoverTab[128913]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:228
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:228
			_go_fuzz_dep_.CoverTab[128914]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:228
			return in[len(in)-1] == '\n'
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:228
			// _ = "end of CoverTab[128914]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:228
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:228
			_go_fuzz_dep_.CoverTab[128915]++
												return p.readPythonMultilines(line, bufferSize)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:229
			// _ = "end of CoverTab[128915]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:230
			_go_fuzz_dep_.CoverTab[128916]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:230
			// _ = "end of CoverTab[128916]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:230
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:230
		// _ = "end of CoverTab[128911]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:230
		_go_fuzz_dep_.CoverTab[128912]++
											return "", nil
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:231
		// _ = "end of CoverTab[128912]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:232
		_go_fuzz_dep_.CoverTab[128917]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:232
		// _ = "end of CoverTab[128917]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:232
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:232
	// _ = "end of CoverTab[128904]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:232
	_go_fuzz_dep_.CoverTab[128905]++

										var valQuote string
										if len(line) > 3 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:235
		_go_fuzz_dep_.CoverTab[128918]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:235
		return line[0:3] == `"""`
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:235
		// _ = "end of CoverTab[128918]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:235
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:235
		_go_fuzz_dep_.CoverTab[128919]++
											valQuote = `"""`
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:236
		// _ = "end of CoverTab[128919]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:237
		_go_fuzz_dep_.CoverTab[128920]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:237
		if line[0] == '`' {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:237
			_go_fuzz_dep_.CoverTab[128921]++
												valQuote = "`"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:238
			// _ = "end of CoverTab[128921]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:239
			_go_fuzz_dep_.CoverTab[128922]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:239
			if p.options.UnescapeValueDoubleQuotes && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:239
				_go_fuzz_dep_.CoverTab[128923]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:239
				return line[0] == '"'
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:239
				// _ = "end of CoverTab[128923]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:239
			}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:239
				_go_fuzz_dep_.CoverTab[128924]++
													valQuote = `"`
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:240
				// _ = "end of CoverTab[128924]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:241
				_go_fuzz_dep_.CoverTab[128925]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:241
				// _ = "end of CoverTab[128925]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:241
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:241
			// _ = "end of CoverTab[128922]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:241
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:241
		// _ = "end of CoverTab[128920]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:241
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:241
	// _ = "end of CoverTab[128905]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:241
	_go_fuzz_dep_.CoverTab[128906]++

										if len(valQuote) > 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:243
		_go_fuzz_dep_.CoverTab[128926]++
											startIdx := len(valQuote)
											pos := strings.LastIndex(line[startIdx:], valQuote)

											if pos == -1 {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:247
			_go_fuzz_dep_.CoverTab[128929]++
												return p.readMultilines(line, line[startIdx:], valQuote)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:248
			// _ = "end of CoverTab[128929]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:249
			_go_fuzz_dep_.CoverTab[128930]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:249
			// _ = "end of CoverTab[128930]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:249
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:249
		// _ = "end of CoverTab[128926]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:249
		_go_fuzz_dep_.CoverTab[128927]++

											if p.options.UnescapeValueDoubleQuotes && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:251
			_go_fuzz_dep_.CoverTab[128931]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:251
			return valQuote == `"`
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:251
			// _ = "end of CoverTab[128931]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:251
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:251
			_go_fuzz_dep_.CoverTab[128932]++
												return strings.Replace(line[startIdx:pos+startIdx], `\"`, `"`, -1), nil
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:252
			// _ = "end of CoverTab[128932]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:253
			_go_fuzz_dep_.CoverTab[128933]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:253
			// _ = "end of CoverTab[128933]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:253
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:253
		// _ = "end of CoverTab[128927]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:253
		_go_fuzz_dep_.CoverTab[128928]++
											return line[startIdx : pos+startIdx], nil
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:254
		// _ = "end of CoverTab[128928]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:255
		_go_fuzz_dep_.CoverTab[128934]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:255
		// _ = "end of CoverTab[128934]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:255
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:255
	// _ = "end of CoverTab[128906]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:255
	_go_fuzz_dep_.CoverTab[128907]++

										lastChar := line[len(line)-1]

										line = strings.TrimSpace(line)
										trimmedLastChar := line[len(line)-1]

//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:263
	if !p.options.IgnoreContinuation && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:263
		_go_fuzz_dep_.CoverTab[128935]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:263
		return trimmedLastChar == '\\'
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:263
		// _ = "end of CoverTab[128935]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:263
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:263
		_go_fuzz_dep_.CoverTab[128936]++
											return p.readContinuationLines(line[:len(line)-1])
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:264
		// _ = "end of CoverTab[128936]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:265
		_go_fuzz_dep_.CoverTab[128937]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:265
		// _ = "end of CoverTab[128937]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:265
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:265
	// _ = "end of CoverTab[128907]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:265
	_go_fuzz_dep_.CoverTab[128908]++

//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:268
	if !p.options.IgnoreInlineComment {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:268
		_go_fuzz_dep_.CoverTab[128938]++
											var i int
											if p.options.SpaceBeforeInlineComment {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:270
			_go_fuzz_dep_.CoverTab[128940]++
												i = strings.Index(line, " #")
												if i == -1 {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:272
				_go_fuzz_dep_.CoverTab[128941]++
													i = strings.Index(line, " ;")
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:273
				// _ = "end of CoverTab[128941]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:274
				_go_fuzz_dep_.CoverTab[128942]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:274
				// _ = "end of CoverTab[128942]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:274
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:274
			// _ = "end of CoverTab[128940]"

		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:276
			_go_fuzz_dep_.CoverTab[128943]++
												i = strings.IndexAny(line, "#;")
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:277
			// _ = "end of CoverTab[128943]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:278
		// _ = "end of CoverTab[128938]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:278
		_go_fuzz_dep_.CoverTab[128939]++

											if i > -1 {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:280
			_go_fuzz_dep_.CoverTab[128944]++
												p.comment.WriteString(line[i:])
												line = strings.TrimSpace(line[:i])
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:282
			// _ = "end of CoverTab[128944]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:283
			_go_fuzz_dep_.CoverTab[128945]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:283
			// _ = "end of CoverTab[128945]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:283
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:283
		// _ = "end of CoverTab[128939]"

	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:285
		_go_fuzz_dep_.CoverTab[128946]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:285
		// _ = "end of CoverTab[128946]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:285
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:285
	// _ = "end of CoverTab[128908]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:285
	_go_fuzz_dep_.CoverTab[128909]++

//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:288
	if (hasSurroundedQuote(line, '\'') || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:288
		_go_fuzz_dep_.CoverTab[128947]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:288
		return hasSurroundedQuote(line, '"')
											// _ = "end of CoverTab[128947]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:289
	}()) && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:289
		_go_fuzz_dep_.CoverTab[128948]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:289
		return !p.options.PreserveSurroundedQuote
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:289
		// _ = "end of CoverTab[128948]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:289
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:289
		_go_fuzz_dep_.CoverTab[128949]++
											line = line[1 : len(line)-1]
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:290
		// _ = "end of CoverTab[128949]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:291
		_go_fuzz_dep_.CoverTab[128950]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:291
		if len(valQuote) == 0 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:291
			_go_fuzz_dep_.CoverTab[128951]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:291
			return p.options.UnescapeValueCommentSymbols
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:291
			// _ = "end of CoverTab[128951]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:291
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:291
			_go_fuzz_dep_.CoverTab[128952]++
												line = strings.ReplaceAll(line, `\;`, ";")
												line = strings.ReplaceAll(line, `\#`, "#")
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:293
			// _ = "end of CoverTab[128952]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:294
			_go_fuzz_dep_.CoverTab[128953]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:294
			if p.options.AllowPythonMultilineValues && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:294
				_go_fuzz_dep_.CoverTab[128954]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:294
				return lastChar == '\n'
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:294
				// _ = "end of CoverTab[128954]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:294
			}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:294
				_go_fuzz_dep_.CoverTab[128955]++
													return p.readPythonMultilines(line, bufferSize)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:295
				// _ = "end of CoverTab[128955]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:296
				_go_fuzz_dep_.CoverTab[128956]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:296
				// _ = "end of CoverTab[128956]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:296
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:296
			// _ = "end of CoverTab[128953]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:296
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:296
		// _ = "end of CoverTab[128950]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:296
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:296
	// _ = "end of CoverTab[128909]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:296
	_go_fuzz_dep_.CoverTab[128910]++

										return line, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:298
	// _ = "end of CoverTab[128910]"
}

func (p *parser) readPythonMultilines(line string, bufferSize int) (string, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:301
	_go_fuzz_dep_.CoverTab[128957]++
										parserBufferPeekResult, _ := p.buf.Peek(bufferSize)
										peekBuffer := bytes.NewBuffer(parserBufferPeekResult)

										indentSize := 0
										for {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:306
		_go_fuzz_dep_.CoverTab[128958]++
											peekData, peekErr := peekBuffer.ReadBytes('\n')
											if peekErr != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:308
			_go_fuzz_dep_.CoverTab[128965]++
												if peekErr == io.EOF {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:309
				_go_fuzz_dep_.CoverTab[128967]++
													p.debug("readPythonMultilines: io.EOF, peekData: %q, line: %q", string(peekData), line)
													return line, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:311
				// _ = "end of CoverTab[128967]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:312
				_go_fuzz_dep_.CoverTab[128968]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:312
				// _ = "end of CoverTab[128968]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:312
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:312
			// _ = "end of CoverTab[128965]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:312
			_go_fuzz_dep_.CoverTab[128966]++

												p.debug("readPythonMultilines: failed to peek with error: %v", peekErr)
												return "", peekErr
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:315
			// _ = "end of CoverTab[128966]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:316
			_go_fuzz_dep_.CoverTab[128969]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:316
			// _ = "end of CoverTab[128969]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:316
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:316
		// _ = "end of CoverTab[128958]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:316
		_go_fuzz_dep_.CoverTab[128959]++

											p.debug("readPythonMultilines: parsing %q", string(peekData))

											peekMatches := pythonMultiline.FindStringSubmatch(string(peekData))
											p.debug("readPythonMultilines: matched %d parts", len(peekMatches))
											for n, v := range peekMatches {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:322
			_go_fuzz_dep_.CoverTab[128970]++
												p.debug("   %d: %q", n, v)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:323
			// _ = "end of CoverTab[128970]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:324
		// _ = "end of CoverTab[128959]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:324
		_go_fuzz_dep_.CoverTab[128960]++

//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:327
		if len(peekMatches) != 3 {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:327
			_go_fuzz_dep_.CoverTab[128971]++
												p.debug("readPythonMultilines: end of value, got: %q", line)
												return line, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:329
			// _ = "end of CoverTab[128971]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:330
			_go_fuzz_dep_.CoverTab[128972]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:330
			// _ = "end of CoverTab[128972]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:330
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:330
		// _ = "end of CoverTab[128960]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:330
		_go_fuzz_dep_.CoverTab[128961]++

//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:333
		currentIndentSize := len(peekMatches[1])
		if indentSize < 1 {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:334
			_go_fuzz_dep_.CoverTab[128973]++
												indentSize = currentIndentSize
												p.debug("readPythonMultilines: indent size is %d", indentSize)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:336
			// _ = "end of CoverTab[128973]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:337
			_go_fuzz_dep_.CoverTab[128974]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:337
			// _ = "end of CoverTab[128974]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:337
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:337
		// _ = "end of CoverTab[128961]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:337
		_go_fuzz_dep_.CoverTab[128962]++

//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:340
		if currentIndentSize < indentSize {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:340
			_go_fuzz_dep_.CoverTab[128975]++
												p.debug("readPythonMultilines: end of value, current indent: %d, expected indent: %d, line: %q", currentIndentSize, indentSize, line)
												return line, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:342
			// _ = "end of CoverTab[128975]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:343
			_go_fuzz_dep_.CoverTab[128976]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:343
			// _ = "end of CoverTab[128976]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:343
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:343
		// _ = "end of CoverTab[128962]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:343
		_go_fuzz_dep_.CoverTab[128963]++

//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:346
		_, err := p.buf.Discard(len(peekData))
		if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:347
			_go_fuzz_dep_.CoverTab[128977]++
												p.debug("readPythonMultilines: failed to skip to the end, returning error")
												return "", err
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:349
			// _ = "end of CoverTab[128977]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:350
			_go_fuzz_dep_.CoverTab[128978]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:350
			// _ = "end of CoverTab[128978]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:350
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:350
		// _ = "end of CoverTab[128963]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:350
		_go_fuzz_dep_.CoverTab[128964]++

//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:353
		line += "\n" + peekMatches[1][indentSize:] + peekMatches[2]
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:353
		// _ = "end of CoverTab[128964]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:354
	// _ = "end of CoverTab[128957]"
}

// parse parses data through an io.Reader.
func (f *File) parse(reader io.Reader) (err error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:358
	_go_fuzz_dep_.CoverTab[128979]++
										p := newParser(reader, parserOptions{
		IgnoreContinuation:		f.options.IgnoreContinuation,
		IgnoreInlineComment:		f.options.IgnoreInlineComment,
		AllowPythonMultilineValues:	f.options.AllowPythonMultilineValues,
		SpaceBeforeInlineComment:	f.options.SpaceBeforeInlineComment,
		UnescapeValueDoubleQuotes:	f.options.UnescapeValueDoubleQuotes,
		UnescapeValueCommentSymbols:	f.options.UnescapeValueCommentSymbols,
		PreserveSurroundedQuote:	f.options.PreserveSurroundedQuote,
		DebugFunc:			f.options.DebugFunc,
		ReaderBufferSize:		f.options.ReaderBufferSize,
	})
	if err = p.BOM(); err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:370
		_go_fuzz_dep_.CoverTab[128984]++
											return fmt.Errorf("BOM: %v", err)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:371
		// _ = "end of CoverTab[128984]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:372
		_go_fuzz_dep_.CoverTab[128985]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:372
		// _ = "end of CoverTab[128985]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:372
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:372
	// _ = "end of CoverTab[128979]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:372
	_go_fuzz_dep_.CoverTab[128980]++

//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:375
	name := DefaultSection
	if f.options.Insensitive || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:376
		_go_fuzz_dep_.CoverTab[128986]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:376
		return f.options.InsensitiveSections
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:376
		// _ = "end of CoverTab[128986]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:376
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:376
		_go_fuzz_dep_.CoverTab[128987]++
											name = strings.ToLower(DefaultSection)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:377
		// _ = "end of CoverTab[128987]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:378
		_go_fuzz_dep_.CoverTab[128988]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:378
		// _ = "end of CoverTab[128988]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:378
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:378
	// _ = "end of CoverTab[128980]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:378
	_go_fuzz_dep_.CoverTab[128981]++
										section, _ := f.NewSection(name)

										// This "last" is not strictly equivalent to "previous one" if current key is not the first nested key
										var isLastValueEmpty bool
										var lastRegularKey *Key

										var line []byte
										var inUnparseableSection bool

//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:391
	parserBufferSize := 0

	currentPeekSize := minReaderBufferSize

	if f.options.AllowPythonMultilineValues {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:395
		_go_fuzz_dep_.CoverTab[128989]++
											for {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:396
			_go_fuzz_dep_.CoverTab[128990]++
												peekBytes, _ := p.buf.Peek(currentPeekSize)
												peekBytesLength := len(peekBytes)

												if parserBufferSize >= peekBytesLength {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:400
				_go_fuzz_dep_.CoverTab[128992]++
													break
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:401
				// _ = "end of CoverTab[128992]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:402
				_go_fuzz_dep_.CoverTab[128993]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:402
				// _ = "end of CoverTab[128993]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:402
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:402
			// _ = "end of CoverTab[128990]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:402
			_go_fuzz_dep_.CoverTab[128991]++

												currentPeekSize *= 2
												parserBufferSize = peekBytesLength
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:405
			// _ = "end of CoverTab[128991]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:406
		// _ = "end of CoverTab[128989]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:407
		_go_fuzz_dep_.CoverTab[128994]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:407
		// _ = "end of CoverTab[128994]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:407
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:407
	// _ = "end of CoverTab[128981]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:407
	_go_fuzz_dep_.CoverTab[128982]++

										for !p.isEOF {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:409
		_go_fuzz_dep_.CoverTab[128995]++
											line, err = p.readUntil('\n')
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:411
			_go_fuzz_dep_.CoverTab[129006]++
												return err
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:412
			// _ = "end of CoverTab[129006]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:413
			_go_fuzz_dep_.CoverTab[129007]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:413
			// _ = "end of CoverTab[129007]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:413
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:413
		// _ = "end of CoverTab[128995]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:413
		_go_fuzz_dep_.CoverTab[128996]++

											if f.options.AllowNestedValues && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:415
			_go_fuzz_dep_.CoverTab[129008]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:415
			return isLastValueEmpty
												// _ = "end of CoverTab[129008]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:416
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:416
			_go_fuzz_dep_.CoverTab[129009]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:416
			return len(line) > 0
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:416
			// _ = "end of CoverTab[129009]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:416
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:416
			_go_fuzz_dep_.CoverTab[129010]++
												if line[0] == ' ' || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:417
				_go_fuzz_dep_.CoverTab[129011]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:417
				return line[0] == '\t'
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:417
				// _ = "end of CoverTab[129011]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:417
			}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:417
				_go_fuzz_dep_.CoverTab[129012]++
													err = lastRegularKey.addNestedValue(string(bytes.TrimSpace(line)))
													if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:419
					_go_fuzz_dep_.CoverTab[129014]++
														return err
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:420
					// _ = "end of CoverTab[129014]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:421
					_go_fuzz_dep_.CoverTab[129015]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:421
					// _ = "end of CoverTab[129015]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:421
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:421
				// _ = "end of CoverTab[129012]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:421
				_go_fuzz_dep_.CoverTab[129013]++
													continue
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:422
				// _ = "end of CoverTab[129013]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:423
				_go_fuzz_dep_.CoverTab[129016]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:423
				// _ = "end of CoverTab[129016]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:423
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:423
			// _ = "end of CoverTab[129010]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:424
			_go_fuzz_dep_.CoverTab[129017]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:424
			// _ = "end of CoverTab[129017]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:424
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:424
		// _ = "end of CoverTab[128996]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:424
		_go_fuzz_dep_.CoverTab[128997]++

											line = bytes.TrimLeftFunc(line, unicode.IsSpace)
											if len(line) == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:427
			_go_fuzz_dep_.CoverTab[129018]++
												continue
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:428
			// _ = "end of CoverTab[129018]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:429
			_go_fuzz_dep_.CoverTab[129019]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:429
			// _ = "end of CoverTab[129019]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:429
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:429
		// _ = "end of CoverTab[128997]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:429
		_go_fuzz_dep_.CoverTab[128998]++

//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:432
		if line[0] == '#' || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:432
			_go_fuzz_dep_.CoverTab[129020]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:432
			return line[0] == ';'
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:432
			// _ = "end of CoverTab[129020]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:432
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:432
			_go_fuzz_dep_.CoverTab[129021]++

//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:436
			p.comment.Write(line)
												continue
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:437
			// _ = "end of CoverTab[129021]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:438
			_go_fuzz_dep_.CoverTab[129022]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:438
			// _ = "end of CoverTab[129022]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:438
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:438
		// _ = "end of CoverTab[128998]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:438
		_go_fuzz_dep_.CoverTab[128999]++

//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:441
		if line[0] == '[' {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:441
			_go_fuzz_dep_.CoverTab[129023]++

												closeIdx := bytes.LastIndexByte(line, ']')
												if closeIdx == -1 {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:444
				_go_fuzz_dep_.CoverTab[129028]++
													return fmt.Errorf("unclosed section: %s", line)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:445
				// _ = "end of CoverTab[129028]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:446
				_go_fuzz_dep_.CoverTab[129029]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:446
				// _ = "end of CoverTab[129029]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:446
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:446
			// _ = "end of CoverTab[129023]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:446
			_go_fuzz_dep_.CoverTab[129024]++

												name := string(line[1:closeIdx])
												section, err = f.NewSection(name)
												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:450
				_go_fuzz_dep_.CoverTab[129030]++
													return err
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:451
				// _ = "end of CoverTab[129030]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:452
				_go_fuzz_dep_.CoverTab[129031]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:452
				// _ = "end of CoverTab[129031]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:452
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:452
			// _ = "end of CoverTab[129024]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:452
			_go_fuzz_dep_.CoverTab[129025]++

												comment, has := cleanComment(line[closeIdx+1:])
												if has {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:455
				_go_fuzz_dep_.CoverTab[129032]++
													p.comment.Write(comment)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:456
				// _ = "end of CoverTab[129032]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:457
				_go_fuzz_dep_.CoverTab[129033]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:457
				// _ = "end of CoverTab[129033]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:457
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:457
			// _ = "end of CoverTab[129025]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:457
			_go_fuzz_dep_.CoverTab[129026]++

												section.Comment = strings.TrimSpace(p.comment.String())

//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:462
			p.comment.Reset()
			p.count = 1

			inUnparseableSection = false
			for i := range f.options.UnparseableSections {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:466
				_go_fuzz_dep_.CoverTab[129034]++
													if f.options.UnparseableSections[i] == name || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:467
					_go_fuzz_dep_.CoverTab[129035]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:467
					return ((f.options.Insensitive || func() bool {
															_go_fuzz_dep_.CoverTab[129036]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:468
						return f.options.InsensitiveSections
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:468
						// _ = "end of CoverTab[129036]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:468
					}()) && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:468
						_go_fuzz_dep_.CoverTab[129037]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:468
						return strings.EqualFold(f.options.UnparseableSections[i], name)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:468
						// _ = "end of CoverTab[129037]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:468
					}())
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:468
					// _ = "end of CoverTab[129035]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:468
				}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:468
					_go_fuzz_dep_.CoverTab[129038]++
														inUnparseableSection = true
														continue
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:470
					// _ = "end of CoverTab[129038]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:471
					_go_fuzz_dep_.CoverTab[129039]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:471
					// _ = "end of CoverTab[129039]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:471
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:471
				// _ = "end of CoverTab[129034]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:472
			// _ = "end of CoverTab[129026]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:472
			_go_fuzz_dep_.CoverTab[129027]++
												continue
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:473
			// _ = "end of CoverTab[129027]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:474
			_go_fuzz_dep_.CoverTab[129040]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:474
			// _ = "end of CoverTab[129040]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:474
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:474
		// _ = "end of CoverTab[128999]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:474
		_go_fuzz_dep_.CoverTab[129000]++

											if inUnparseableSection {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:476
			_go_fuzz_dep_.CoverTab[129041]++
												section.isRawSection = true
												section.rawBody += string(line)
												continue
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:479
			// _ = "end of CoverTab[129041]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:480
			_go_fuzz_dep_.CoverTab[129042]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:480
			// _ = "end of CoverTab[129042]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:480
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:480
		// _ = "end of CoverTab[129000]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:480
		_go_fuzz_dep_.CoverTab[129001]++

											kname, offset, err := readKeyName(f.options.KeyValueDelimiters, line)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:483
			_go_fuzz_dep_.CoverTab[129043]++

												if IsErrDelimiterNotFound(err) {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:485
				_go_fuzz_dep_.CoverTab[129045]++
													switch {
				case f.options.AllowBooleanKeys:
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:487
					_go_fuzz_dep_.CoverTab[129046]++
														kname, err := p.readValue(line, parserBufferSize)
														if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:489
						_go_fuzz_dep_.CoverTab[129051]++
															return err
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:490
						// _ = "end of CoverTab[129051]"
					} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:491
						_go_fuzz_dep_.CoverTab[129052]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:491
						// _ = "end of CoverTab[129052]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:491
					}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:491
					// _ = "end of CoverTab[129046]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:491
					_go_fuzz_dep_.CoverTab[129047]++
														key, err := section.NewBooleanKey(kname)
														if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:493
						_go_fuzz_dep_.CoverTab[129053]++
															return err
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:494
						// _ = "end of CoverTab[129053]"
					} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:495
						_go_fuzz_dep_.CoverTab[129054]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:495
						// _ = "end of CoverTab[129054]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:495
					}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:495
					// _ = "end of CoverTab[129047]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:495
					_go_fuzz_dep_.CoverTab[129048]++
														key.Comment = strings.TrimSpace(p.comment.String())
														p.comment.Reset()
														continue
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:498
					// _ = "end of CoverTab[129048]"

				case f.options.SkipUnrecognizableLines:
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:500
					_go_fuzz_dep_.CoverTab[129049]++
														continue
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:501
					// _ = "end of CoverTab[129049]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:501
				default:
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:501
					_go_fuzz_dep_.CoverTab[129050]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:501
					// _ = "end of CoverTab[129050]"
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:502
				// _ = "end of CoverTab[129045]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:503
				_go_fuzz_dep_.CoverTab[129055]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:503
				// _ = "end of CoverTab[129055]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:503
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:503
			// _ = "end of CoverTab[129043]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:503
			_go_fuzz_dep_.CoverTab[129044]++
												return err
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:504
			// _ = "end of CoverTab[129044]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:505
			_go_fuzz_dep_.CoverTab[129056]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:505
			// _ = "end of CoverTab[129056]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:505
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:505
		// _ = "end of CoverTab[129001]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:505
		_go_fuzz_dep_.CoverTab[129002]++

//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:508
		isAutoIncr := false
		if kname == "-" {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:509
			_go_fuzz_dep_.CoverTab[129057]++
												isAutoIncr = true
												kname = "#" + strconv.Itoa(p.count)
												p.count++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:512
			// _ = "end of CoverTab[129057]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:513
			_go_fuzz_dep_.CoverTab[129058]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:513
			// _ = "end of CoverTab[129058]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:513
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:513
		// _ = "end of CoverTab[129002]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:513
		_go_fuzz_dep_.CoverTab[129003]++

											value, err := p.readValue(line[offset:], parserBufferSize)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:516
			_go_fuzz_dep_.CoverTab[129059]++
												return err
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:517
			// _ = "end of CoverTab[129059]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:518
			_go_fuzz_dep_.CoverTab[129060]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:518
			// _ = "end of CoverTab[129060]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:518
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:518
		// _ = "end of CoverTab[129003]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:518
		_go_fuzz_dep_.CoverTab[129004]++
											isLastValueEmpty = len(value) == 0

											key, err := section.NewKey(kname, value)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:522
			_go_fuzz_dep_.CoverTab[129061]++
												return err
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:523
			// _ = "end of CoverTab[129061]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:524
			_go_fuzz_dep_.CoverTab[129062]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:524
			// _ = "end of CoverTab[129062]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:524
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:524
		// _ = "end of CoverTab[129004]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:524
		_go_fuzz_dep_.CoverTab[129005]++
											key.isAutoIncrement = isAutoIncr
											key.Comment = strings.TrimSpace(p.comment.String())
											p.comment.Reset()
											lastRegularKey = key
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:528
		// _ = "end of CoverTab[129005]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:529
	// _ = "end of CoverTab[128982]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:529
	_go_fuzz_dep_.CoverTab[128983]++
										return nil
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:530
	// _ = "end of CoverTab[128983]"
}

//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:531
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/parser.go:531
var _ = _go_fuzz_dep_.CoverTab
