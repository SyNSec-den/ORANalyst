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

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:37
package proto

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:37
import (
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:37
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:37
)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:37
import (
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:37
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:37
)

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:41
import (
	"bufio"
	"bytes"
	"encoding"
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"reflect"
	"sort"
	"strings"
	"sync"
	"time"
)

var (
	newline		= []byte("\n")
	spaces		= []byte("                                        ")
	endBraceNewline	= []byte("}\n")
	backslashN	= []byte{'\\', 'n'}
	backslashR	= []byte{'\\', 'r'}
	backslashT	= []byte{'\\', 't'}
	backslashDQ	= []byte{'\\', '"'}
	backslashBS	= []byte{'\\', '\\'}
	posInf		= []byte("inf")
	negInf		= []byte("-inf")
	nan		= []byte("nan")
)

type writer interface {
	io.Writer
	WriteByte(byte) error
}

// textWriter is an io.Writer that tracks its indentation level.
type textWriter struct {
	ind		int
	complete	bool	// if the current position is a complete line
	compact		bool	// whether to write out as a one-liner
	w		writer
}

func (w *textWriter) WriteString(s string) (n int, err error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:84
	_go_fuzz_dep_.CoverTab[112473]++
											if !strings.Contains(s, "\n") {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:85
		_go_fuzz_dep_.CoverTab[112475]++
												if !w.compact && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:86
			_go_fuzz_dep_.CoverTab[112477]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:86
			return w.complete
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:86
			// _ = "end of CoverTab[112477]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:86
		}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:86
			_go_fuzz_dep_.CoverTab[112478]++
													w.writeIndent()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:87
			// _ = "end of CoverTab[112478]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:88
			_go_fuzz_dep_.CoverTab[112479]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:88
			// _ = "end of CoverTab[112479]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:88
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:88
		// _ = "end of CoverTab[112475]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:88
		_go_fuzz_dep_.CoverTab[112476]++
												w.complete = false
												return io.WriteString(w.w, s)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:90
		// _ = "end of CoverTab[112476]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:91
		_go_fuzz_dep_.CoverTab[112480]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:91
		// _ = "end of CoverTab[112480]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:91
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:91
	// _ = "end of CoverTab[112473]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:91
	_go_fuzz_dep_.CoverTab[112474]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:95
	return w.Write([]byte(s))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:95
	// _ = "end of CoverTab[112474]"
}

func (w *textWriter) Write(p []byte) (n int, err error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:98
	_go_fuzz_dep_.CoverTab[112481]++
											newlines := bytes.Count(p, newline)
											if newlines == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:100
		_go_fuzz_dep_.CoverTab[112485]++
												if !w.compact && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:101
			_go_fuzz_dep_.CoverTab[112487]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:101
			return w.complete
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:101
			// _ = "end of CoverTab[112487]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:101
		}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:101
			_go_fuzz_dep_.CoverTab[112488]++
													w.writeIndent()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:102
			// _ = "end of CoverTab[112488]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:103
			_go_fuzz_dep_.CoverTab[112489]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:103
			// _ = "end of CoverTab[112489]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:103
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:103
		// _ = "end of CoverTab[112485]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:103
		_go_fuzz_dep_.CoverTab[112486]++
												n, err = w.w.Write(p)
												w.complete = false
												return n, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:106
		// _ = "end of CoverTab[112486]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:107
		_go_fuzz_dep_.CoverTab[112490]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:107
		// _ = "end of CoverTab[112490]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:107
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:107
	// _ = "end of CoverTab[112481]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:107
	_go_fuzz_dep_.CoverTab[112482]++

											frags := bytes.SplitN(p, newline, newlines+1)
											if w.compact {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:110
		_go_fuzz_dep_.CoverTab[112491]++
												for i, frag := range frags {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:111
			_go_fuzz_dep_.CoverTab[112493]++
													if i > 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:112
				_go_fuzz_dep_.CoverTab[112495]++
														if err := w.w.WriteByte(' '); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:113
					_go_fuzz_dep_.CoverTab[112497]++
															return n, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:114
					// _ = "end of CoverTab[112497]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:115
					_go_fuzz_dep_.CoverTab[112498]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:115
					// _ = "end of CoverTab[112498]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:115
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:115
				// _ = "end of CoverTab[112495]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:115
				_go_fuzz_dep_.CoverTab[112496]++
														n++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:116
				// _ = "end of CoverTab[112496]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:117
				_go_fuzz_dep_.CoverTab[112499]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:117
				// _ = "end of CoverTab[112499]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:117
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:117
			// _ = "end of CoverTab[112493]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:117
			_go_fuzz_dep_.CoverTab[112494]++
													nn, err := w.w.Write(frag)
													n += nn
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:120
				_go_fuzz_dep_.CoverTab[112500]++
														return n, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:121
				// _ = "end of CoverTab[112500]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:122
				_go_fuzz_dep_.CoverTab[112501]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:122
				// _ = "end of CoverTab[112501]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:122
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:122
			// _ = "end of CoverTab[112494]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:123
		// _ = "end of CoverTab[112491]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:123
		_go_fuzz_dep_.CoverTab[112492]++
												return n, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:124
		// _ = "end of CoverTab[112492]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:125
		_go_fuzz_dep_.CoverTab[112502]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:125
		// _ = "end of CoverTab[112502]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:125
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:125
	// _ = "end of CoverTab[112482]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:125
	_go_fuzz_dep_.CoverTab[112483]++

											for i, frag := range frags {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:127
		_go_fuzz_dep_.CoverTab[112503]++
												if w.complete {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:128
			_go_fuzz_dep_.CoverTab[112506]++
													w.writeIndent()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:129
			// _ = "end of CoverTab[112506]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:130
			_go_fuzz_dep_.CoverTab[112507]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:130
			// _ = "end of CoverTab[112507]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:130
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:130
		// _ = "end of CoverTab[112503]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:130
		_go_fuzz_dep_.CoverTab[112504]++
												nn, err := w.w.Write(frag)
												n += nn
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:133
			_go_fuzz_dep_.CoverTab[112508]++
													return n, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:134
			// _ = "end of CoverTab[112508]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:135
			_go_fuzz_dep_.CoverTab[112509]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:135
			// _ = "end of CoverTab[112509]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:135
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:135
		// _ = "end of CoverTab[112504]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:135
		_go_fuzz_dep_.CoverTab[112505]++
												if i+1 < len(frags) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:136
			_go_fuzz_dep_.CoverTab[112510]++
													if err := w.w.WriteByte('\n'); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:137
				_go_fuzz_dep_.CoverTab[112512]++
														return n, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:138
				// _ = "end of CoverTab[112512]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:139
				_go_fuzz_dep_.CoverTab[112513]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:139
				// _ = "end of CoverTab[112513]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:139
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:139
			// _ = "end of CoverTab[112510]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:139
			_go_fuzz_dep_.CoverTab[112511]++
													n++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:140
			// _ = "end of CoverTab[112511]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:141
			_go_fuzz_dep_.CoverTab[112514]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:141
			// _ = "end of CoverTab[112514]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:141
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:141
		// _ = "end of CoverTab[112505]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:142
	// _ = "end of CoverTab[112483]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:142
	_go_fuzz_dep_.CoverTab[112484]++
											w.complete = len(frags[len(frags)-1]) == 0
											return n, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:144
	// _ = "end of CoverTab[112484]"
}

func (w *textWriter) WriteByte(c byte) error {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:147
	_go_fuzz_dep_.CoverTab[112515]++
											if w.compact && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:148
		_go_fuzz_dep_.CoverTab[112518]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:148
		return c == '\n'
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:148
		// _ = "end of CoverTab[112518]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:148
	}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:148
		_go_fuzz_dep_.CoverTab[112519]++
												c = ' '
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:149
		// _ = "end of CoverTab[112519]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:150
		_go_fuzz_dep_.CoverTab[112520]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:150
		// _ = "end of CoverTab[112520]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:150
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:150
	// _ = "end of CoverTab[112515]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:150
	_go_fuzz_dep_.CoverTab[112516]++
											if !w.compact && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:151
		_go_fuzz_dep_.CoverTab[112521]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:151
		return w.complete
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:151
		// _ = "end of CoverTab[112521]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:151
	}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:151
		_go_fuzz_dep_.CoverTab[112522]++
												w.writeIndent()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:152
		// _ = "end of CoverTab[112522]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:153
		_go_fuzz_dep_.CoverTab[112523]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:153
		// _ = "end of CoverTab[112523]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:153
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:153
	// _ = "end of CoverTab[112516]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:153
	_go_fuzz_dep_.CoverTab[112517]++
											err := w.w.WriteByte(c)
											w.complete = c == '\n'
											return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:156
	// _ = "end of CoverTab[112517]"
}

func (w *textWriter) indent() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:159
	_go_fuzz_dep_.CoverTab[112524]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:159
	w.ind++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:159
	// _ = "end of CoverTab[112524]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:159
}

func (w *textWriter) unindent() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:161
	_go_fuzz_dep_.CoverTab[112525]++
											if w.ind == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:162
		_go_fuzz_dep_.CoverTab[112527]++
												log.Print("proto: textWriter unindented too far")
												return
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:164
		// _ = "end of CoverTab[112527]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:165
		_go_fuzz_dep_.CoverTab[112528]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:165
		// _ = "end of CoverTab[112528]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:165
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:165
	// _ = "end of CoverTab[112525]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:165
	_go_fuzz_dep_.CoverTab[112526]++
											w.ind--
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:166
	// _ = "end of CoverTab[112526]"
}

func writeName(w *textWriter, props *Properties) error {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:169
	_go_fuzz_dep_.CoverTab[112529]++
											if _, err := w.WriteString(props.OrigName); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:170
		_go_fuzz_dep_.CoverTab[112532]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:171
		// _ = "end of CoverTab[112532]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:172
		_go_fuzz_dep_.CoverTab[112533]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:172
		// _ = "end of CoverTab[112533]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:172
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:172
	// _ = "end of CoverTab[112529]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:172
	_go_fuzz_dep_.CoverTab[112530]++
											if props.Wire != "group" {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:173
		_go_fuzz_dep_.CoverTab[112534]++
												return w.WriteByte(':')
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:174
		// _ = "end of CoverTab[112534]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:175
		_go_fuzz_dep_.CoverTab[112535]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:175
		// _ = "end of CoverTab[112535]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:175
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:175
	// _ = "end of CoverTab[112530]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:175
	_go_fuzz_dep_.CoverTab[112531]++
											return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:176
	// _ = "end of CoverTab[112531]"
}

func requiresQuotes(u string) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:179
	_go_fuzz_dep_.CoverTab[112536]++

											for _, ch := range u {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:181
		_go_fuzz_dep_.CoverTab[112538]++
												switch {
		case ch == '.' || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:183
			_go_fuzz_dep_.CoverTab[112544]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:183
			return ch == '/'
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:183
			// _ = "end of CoverTab[112544]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:183
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:183
			_go_fuzz_dep_.CoverTab[112545]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:183
			return ch == '_'
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:183
			// _ = "end of CoverTab[112545]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:183
		}():
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:183
			_go_fuzz_dep_.CoverTab[112539]++
													continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:184
			// _ = "end of CoverTab[112539]"
		case '0' <= ch && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:185
			_go_fuzz_dep_.CoverTab[112546]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:185
			return ch <= '9'
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:185
			// _ = "end of CoverTab[112546]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:185
		}():
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:185
			_go_fuzz_dep_.CoverTab[112540]++
													continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:186
			// _ = "end of CoverTab[112540]"
		case 'A' <= ch && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:187
			_go_fuzz_dep_.CoverTab[112547]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:187
			return ch <= 'Z'
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:187
			// _ = "end of CoverTab[112547]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:187
		}():
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:187
			_go_fuzz_dep_.CoverTab[112541]++
													continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:188
			// _ = "end of CoverTab[112541]"
		case 'a' <= ch && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:189
			_go_fuzz_dep_.CoverTab[112548]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:189
			return ch <= 'z'
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:189
			// _ = "end of CoverTab[112548]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:189
		}():
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:189
			_go_fuzz_dep_.CoverTab[112542]++
													continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:190
			// _ = "end of CoverTab[112542]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:191
			_go_fuzz_dep_.CoverTab[112543]++
													return true
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:192
			// _ = "end of CoverTab[112543]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:193
		// _ = "end of CoverTab[112538]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:194
	// _ = "end of CoverTab[112536]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:194
	_go_fuzz_dep_.CoverTab[112537]++
											return false
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:195
	// _ = "end of CoverTab[112537]"
}

// isAny reports whether sv is a google.protobuf.Any message
func isAny(sv reflect.Value) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:199
	_go_fuzz_dep_.CoverTab[112549]++
											type wkt interface {
		XXX_WellKnownType() string
	}
	t, ok := sv.Addr().Interface().(wkt)
	return ok && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:204
		_go_fuzz_dep_.CoverTab[112550]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:204
		return t.XXX_WellKnownType() == "Any"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:204
		// _ = "end of CoverTab[112550]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:204
	}()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:204
	// _ = "end of CoverTab[112549]"
}

// writeProto3Any writes an expanded google.protobuf.Any message.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:207
//
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:207
// It returns (false, nil) if sv value can't be unmarshaled (e.g. because
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:207
// required messages are not linked in).
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:207
//
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:207
// It returns (true, error) when sv was written in expanded format or an error
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:207
// was encountered.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:214
func (tm *TextMarshaler) writeProto3Any(w *textWriter, sv reflect.Value) (bool, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:214
	_go_fuzz_dep_.CoverTab[112551]++
											turl := sv.FieldByName("TypeUrl")
											val := sv.FieldByName("Value")
											if !turl.IsValid() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:217
		_go_fuzz_dep_.CoverTab[112560]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:217
		return !val.IsValid()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:217
		// _ = "end of CoverTab[112560]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:217
	}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:217
		_go_fuzz_dep_.CoverTab[112561]++
												return true, errors.New("proto: invalid google.protobuf.Any message")
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:218
		// _ = "end of CoverTab[112561]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:219
		_go_fuzz_dep_.CoverTab[112562]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:219
		// _ = "end of CoverTab[112562]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:219
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:219
	// _ = "end of CoverTab[112551]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:219
	_go_fuzz_dep_.CoverTab[112552]++

											b, ok := val.Interface().([]byte)
											if !ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:222
		_go_fuzz_dep_.CoverTab[112563]++
												return true, errors.New("proto: invalid google.protobuf.Any message")
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:223
		// _ = "end of CoverTab[112563]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:224
		_go_fuzz_dep_.CoverTab[112564]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:224
		// _ = "end of CoverTab[112564]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:224
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:224
	// _ = "end of CoverTab[112552]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:224
	_go_fuzz_dep_.CoverTab[112553]++

											parts := strings.Split(turl.String(), "/")
											mt := MessageType(parts[len(parts)-1])
											if mt == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:228
		_go_fuzz_dep_.CoverTab[112565]++
												return false, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:229
		// _ = "end of CoverTab[112565]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:230
		_go_fuzz_dep_.CoverTab[112566]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:230
		// _ = "end of CoverTab[112566]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:230
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:230
	// _ = "end of CoverTab[112553]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:230
	_go_fuzz_dep_.CoverTab[112554]++
											m := reflect.New(mt.Elem())
											if err := Unmarshal(b, m.Interface().(Message)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:232
		_go_fuzz_dep_.CoverTab[112567]++
												return false, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:233
		// _ = "end of CoverTab[112567]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:234
		_go_fuzz_dep_.CoverTab[112568]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:234
		// _ = "end of CoverTab[112568]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:234
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:234
	// _ = "end of CoverTab[112554]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:234
	_go_fuzz_dep_.CoverTab[112555]++
											w.Write([]byte("["))
											u := turl.String()
											if requiresQuotes(u) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:237
		_go_fuzz_dep_.CoverTab[112569]++
												writeString(w, u)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:238
		// _ = "end of CoverTab[112569]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:239
		_go_fuzz_dep_.CoverTab[112570]++
												w.Write([]byte(u))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:240
		// _ = "end of CoverTab[112570]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:241
	// _ = "end of CoverTab[112555]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:241
	_go_fuzz_dep_.CoverTab[112556]++
											if w.compact {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:242
		_go_fuzz_dep_.CoverTab[112571]++
												w.Write([]byte("]:<"))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:243
		// _ = "end of CoverTab[112571]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:244
		_go_fuzz_dep_.CoverTab[112572]++
												w.Write([]byte("]: <\n"))
												w.ind++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:246
		// _ = "end of CoverTab[112572]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:247
	// _ = "end of CoverTab[112556]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:247
	_go_fuzz_dep_.CoverTab[112557]++
											if err := tm.writeStruct(w, m.Elem()); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:248
		_go_fuzz_dep_.CoverTab[112573]++
												return true, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:249
		// _ = "end of CoverTab[112573]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:250
		_go_fuzz_dep_.CoverTab[112574]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:250
		// _ = "end of CoverTab[112574]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:250
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:250
	// _ = "end of CoverTab[112557]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:250
	_go_fuzz_dep_.CoverTab[112558]++
											if w.compact {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:251
		_go_fuzz_dep_.CoverTab[112575]++
												w.Write([]byte("> "))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:252
		// _ = "end of CoverTab[112575]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:253
		_go_fuzz_dep_.CoverTab[112576]++
												w.ind--
												w.Write([]byte(">\n"))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:255
		// _ = "end of CoverTab[112576]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:256
	// _ = "end of CoverTab[112558]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:256
	_go_fuzz_dep_.CoverTab[112559]++
											return true, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:257
	// _ = "end of CoverTab[112559]"
}

func (tm *TextMarshaler) writeStruct(w *textWriter, sv reflect.Value) error {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:260
	_go_fuzz_dep_.CoverTab[112577]++
											if tm.ExpandAny && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:261
		_go_fuzz_dep_.CoverTab[112582]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:261
		return isAny(sv)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:261
		// _ = "end of CoverTab[112582]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:261
	}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:261
		_go_fuzz_dep_.CoverTab[112583]++
												if canExpand, err := tm.writeProto3Any(w, sv); canExpand {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:262
			_go_fuzz_dep_.CoverTab[112584]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:263
			// _ = "end of CoverTab[112584]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:264
			_go_fuzz_dep_.CoverTab[112585]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:264
			// _ = "end of CoverTab[112585]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:264
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:264
		// _ = "end of CoverTab[112583]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:265
		_go_fuzz_dep_.CoverTab[112586]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:265
		// _ = "end of CoverTab[112586]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:265
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:265
	// _ = "end of CoverTab[112577]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:265
	_go_fuzz_dep_.CoverTab[112578]++
											st := sv.Type()
											sprops := GetProperties(st)
											for i := 0; i < sv.NumField(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:268
		_go_fuzz_dep_.CoverTab[112587]++
												fv := sv.Field(i)
												props := sprops.Prop[i]
												name := st.Field(i).Name

												if name == "XXX_NoUnkeyedLiteral" {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:273
			_go_fuzz_dep_.CoverTab[112600]++
													continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:274
			// _ = "end of CoverTab[112600]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:275
			_go_fuzz_dep_.CoverTab[112601]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:275
			// _ = "end of CoverTab[112601]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:275
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:275
		// _ = "end of CoverTab[112587]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:275
		_go_fuzz_dep_.CoverTab[112588]++

												if strings.HasPrefix(name, "XXX_") {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:277
			_go_fuzz_dep_.CoverTab[112602]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:283
			if name == "XXX_unrecognized" && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:283
				_go_fuzz_dep_.CoverTab[112604]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:283
				return !fv.IsNil()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:283
				// _ = "end of CoverTab[112604]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:283
			}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:283
				_go_fuzz_dep_.CoverTab[112605]++
														if err := writeUnknownStruct(w, fv.Interface().([]byte)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:284
					_go_fuzz_dep_.CoverTab[112606]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:285
					// _ = "end of CoverTab[112606]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:286
					_go_fuzz_dep_.CoverTab[112607]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:286
					// _ = "end of CoverTab[112607]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:286
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:286
				// _ = "end of CoverTab[112605]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:287
				_go_fuzz_dep_.CoverTab[112608]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:287
				// _ = "end of CoverTab[112608]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:287
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:287
			// _ = "end of CoverTab[112602]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:287
			_go_fuzz_dep_.CoverTab[112603]++
													continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:288
			// _ = "end of CoverTab[112603]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:289
			_go_fuzz_dep_.CoverTab[112609]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:289
			// _ = "end of CoverTab[112609]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:289
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:289
		// _ = "end of CoverTab[112588]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:289
		_go_fuzz_dep_.CoverTab[112589]++
												if fv.Kind() == reflect.Ptr && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:290
			_go_fuzz_dep_.CoverTab[112610]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:290
			return fv.IsNil()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:290
			// _ = "end of CoverTab[112610]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:290
		}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:290
			_go_fuzz_dep_.CoverTab[112611]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:294
			continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:294
			// _ = "end of CoverTab[112611]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:295
			_go_fuzz_dep_.CoverTab[112612]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:295
			// _ = "end of CoverTab[112612]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:295
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:295
		// _ = "end of CoverTab[112589]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:295
		_go_fuzz_dep_.CoverTab[112590]++
												if fv.Kind() == reflect.Slice && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:296
			_go_fuzz_dep_.CoverTab[112613]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:296
			return fv.IsNil()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:296
			// _ = "end of CoverTab[112613]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:296
		}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:296
			_go_fuzz_dep_.CoverTab[112614]++

													continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:298
			// _ = "end of CoverTab[112614]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:299
			_go_fuzz_dep_.CoverTab[112615]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:299
			// _ = "end of CoverTab[112615]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:299
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:299
		// _ = "end of CoverTab[112590]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:299
		_go_fuzz_dep_.CoverTab[112591]++

												if props.Repeated && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:301
			_go_fuzz_dep_.CoverTab[112616]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:301
			return fv.Kind() == reflect.Slice
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:301
			// _ = "end of CoverTab[112616]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:301
		}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:301
			_go_fuzz_dep_.CoverTab[112617]++

													for j := 0; j < fv.Len(); j++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:303
				_go_fuzz_dep_.CoverTab[112619]++
														if err := writeName(w, props); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:304
					_go_fuzz_dep_.CoverTab[112624]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:305
					// _ = "end of CoverTab[112624]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:306
					_go_fuzz_dep_.CoverTab[112625]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:306
					// _ = "end of CoverTab[112625]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:306
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:306
				// _ = "end of CoverTab[112619]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:306
				_go_fuzz_dep_.CoverTab[112620]++
														if !w.compact {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:307
					_go_fuzz_dep_.CoverTab[112626]++
															if err := w.WriteByte(' '); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:308
						_go_fuzz_dep_.CoverTab[112627]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:309
						// _ = "end of CoverTab[112627]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:310
						_go_fuzz_dep_.CoverTab[112628]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:310
						// _ = "end of CoverTab[112628]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:310
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:310
					// _ = "end of CoverTab[112626]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:311
					_go_fuzz_dep_.CoverTab[112629]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:311
					// _ = "end of CoverTab[112629]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:311
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:311
				// _ = "end of CoverTab[112620]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:311
				_go_fuzz_dep_.CoverTab[112621]++
														v := fv.Index(j)
														if v.Kind() == reflect.Ptr && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:313
					_go_fuzz_dep_.CoverTab[112630]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:313
					return v.IsNil()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:313
					// _ = "end of CoverTab[112630]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:313
				}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:313
					_go_fuzz_dep_.CoverTab[112631]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:316
					if _, err := w.Write([]byte("<nil>\n")); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:316
						_go_fuzz_dep_.CoverTab[112633]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:317
						// _ = "end of CoverTab[112633]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:318
						_go_fuzz_dep_.CoverTab[112634]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:318
						// _ = "end of CoverTab[112634]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:318
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:318
					// _ = "end of CoverTab[112631]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:318
					_go_fuzz_dep_.CoverTab[112632]++
															continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:319
					// _ = "end of CoverTab[112632]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:320
					_go_fuzz_dep_.CoverTab[112635]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:320
					// _ = "end of CoverTab[112635]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:320
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:320
				// _ = "end of CoverTab[112621]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:320
				_go_fuzz_dep_.CoverTab[112622]++
														if len(props.Enum) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:321
					_go_fuzz_dep_.CoverTab[112636]++
															if err := tm.writeEnum(w, v, props); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:322
						_go_fuzz_dep_.CoverTab[112637]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:323
						// _ = "end of CoverTab[112637]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:324
						_go_fuzz_dep_.CoverTab[112638]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:324
						// _ = "end of CoverTab[112638]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:324
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:324
					// _ = "end of CoverTab[112636]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:325
					_go_fuzz_dep_.CoverTab[112639]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:325
					if err := tm.writeAny(w, v, props); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:325
						_go_fuzz_dep_.CoverTab[112640]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:326
						// _ = "end of CoverTab[112640]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:327
						_go_fuzz_dep_.CoverTab[112641]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:327
						// _ = "end of CoverTab[112641]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:327
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:327
					// _ = "end of CoverTab[112639]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:327
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:327
				// _ = "end of CoverTab[112622]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:327
				_go_fuzz_dep_.CoverTab[112623]++
														if err := w.WriteByte('\n'); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:328
					_go_fuzz_dep_.CoverTab[112642]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:329
					// _ = "end of CoverTab[112642]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:330
					_go_fuzz_dep_.CoverTab[112643]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:330
					// _ = "end of CoverTab[112643]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:330
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:330
				// _ = "end of CoverTab[112623]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:331
			// _ = "end of CoverTab[112617]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:331
			_go_fuzz_dep_.CoverTab[112618]++
													continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:332
			// _ = "end of CoverTab[112618]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:333
			_go_fuzz_dep_.CoverTab[112644]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:333
			// _ = "end of CoverTab[112644]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:333
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:333
		// _ = "end of CoverTab[112591]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:333
		_go_fuzz_dep_.CoverTab[112592]++
												if fv.Kind() == reflect.Map {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:334
			_go_fuzz_dep_.CoverTab[112645]++

													keys := fv.MapKeys()
													sort.Sort(mapKeys(keys))
													for _, key := range keys {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:338
				_go_fuzz_dep_.CoverTab[112647]++
														val := fv.MapIndex(key)
														if err := writeName(w, props); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:340
					_go_fuzz_dep_.CoverTab[112658]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:341
					// _ = "end of CoverTab[112658]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:342
					_go_fuzz_dep_.CoverTab[112659]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:342
					// _ = "end of CoverTab[112659]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:342
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:342
				// _ = "end of CoverTab[112647]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:342
				_go_fuzz_dep_.CoverTab[112648]++
														if !w.compact {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:343
					_go_fuzz_dep_.CoverTab[112660]++
															if err := w.WriteByte(' '); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:344
						_go_fuzz_dep_.CoverTab[112661]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:345
						// _ = "end of CoverTab[112661]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:346
						_go_fuzz_dep_.CoverTab[112662]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:346
						// _ = "end of CoverTab[112662]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:346
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:346
					// _ = "end of CoverTab[112660]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:347
					_go_fuzz_dep_.CoverTab[112663]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:347
					// _ = "end of CoverTab[112663]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:347
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:347
				// _ = "end of CoverTab[112648]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:347
				_go_fuzz_dep_.CoverTab[112649]++

														if err := w.WriteByte('<'); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:349
					_go_fuzz_dep_.CoverTab[112664]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:350
					// _ = "end of CoverTab[112664]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:351
					_go_fuzz_dep_.CoverTab[112665]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:351
					// _ = "end of CoverTab[112665]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:351
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:351
				// _ = "end of CoverTab[112649]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:351
				_go_fuzz_dep_.CoverTab[112650]++
														if !w.compact {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:352
					_go_fuzz_dep_.CoverTab[112666]++
															if err := w.WriteByte('\n'); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:353
						_go_fuzz_dep_.CoverTab[112667]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:354
						// _ = "end of CoverTab[112667]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:355
						_go_fuzz_dep_.CoverTab[112668]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:355
						// _ = "end of CoverTab[112668]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:355
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:355
					// _ = "end of CoverTab[112666]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:356
					_go_fuzz_dep_.CoverTab[112669]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:356
					// _ = "end of CoverTab[112669]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:356
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:356
				// _ = "end of CoverTab[112650]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:356
				_go_fuzz_dep_.CoverTab[112651]++
														w.indent()

														if _, err := w.WriteString("key:"); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:359
					_go_fuzz_dep_.CoverTab[112670]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:360
					// _ = "end of CoverTab[112670]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:361
					_go_fuzz_dep_.CoverTab[112671]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:361
					// _ = "end of CoverTab[112671]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:361
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:361
				// _ = "end of CoverTab[112651]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:361
				_go_fuzz_dep_.CoverTab[112652]++
														if !w.compact {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:362
					_go_fuzz_dep_.CoverTab[112672]++
															if err := w.WriteByte(' '); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:363
						_go_fuzz_dep_.CoverTab[112673]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:364
						// _ = "end of CoverTab[112673]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:365
						_go_fuzz_dep_.CoverTab[112674]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:365
						// _ = "end of CoverTab[112674]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:365
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:365
					// _ = "end of CoverTab[112672]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:366
					_go_fuzz_dep_.CoverTab[112675]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:366
					// _ = "end of CoverTab[112675]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:366
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:366
				// _ = "end of CoverTab[112652]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:366
				_go_fuzz_dep_.CoverTab[112653]++
														if err := tm.writeAny(w, key, props.MapKeyProp); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:367
					_go_fuzz_dep_.CoverTab[112676]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:368
					// _ = "end of CoverTab[112676]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:369
					_go_fuzz_dep_.CoverTab[112677]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:369
					// _ = "end of CoverTab[112677]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:369
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:369
				// _ = "end of CoverTab[112653]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:369
				_go_fuzz_dep_.CoverTab[112654]++
														if err := w.WriteByte('\n'); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:370
					_go_fuzz_dep_.CoverTab[112678]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:371
					// _ = "end of CoverTab[112678]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:372
					_go_fuzz_dep_.CoverTab[112679]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:372
					// _ = "end of CoverTab[112679]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:372
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:372
				// _ = "end of CoverTab[112654]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:372
				_go_fuzz_dep_.CoverTab[112655]++

														if val.Kind() != reflect.Ptr || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:374
					_go_fuzz_dep_.CoverTab[112680]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:374
					return !val.IsNil()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:374
					// _ = "end of CoverTab[112680]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:374
				}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:374
					_go_fuzz_dep_.CoverTab[112681]++

															if _, err := w.WriteString("value:"); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:376
						_go_fuzz_dep_.CoverTab[112685]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:377
						// _ = "end of CoverTab[112685]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:378
						_go_fuzz_dep_.CoverTab[112686]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:378
						// _ = "end of CoverTab[112686]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:378
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:378
					// _ = "end of CoverTab[112681]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:378
					_go_fuzz_dep_.CoverTab[112682]++
															if !w.compact {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:379
						_go_fuzz_dep_.CoverTab[112687]++
																if err := w.WriteByte(' '); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:380
							_go_fuzz_dep_.CoverTab[112688]++
																	return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:381
							// _ = "end of CoverTab[112688]"
						} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:382
							_go_fuzz_dep_.CoverTab[112689]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:382
							// _ = "end of CoverTab[112689]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:382
						}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:382
						// _ = "end of CoverTab[112687]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:383
						_go_fuzz_dep_.CoverTab[112690]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:383
						// _ = "end of CoverTab[112690]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:383
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:383
					// _ = "end of CoverTab[112682]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:383
					_go_fuzz_dep_.CoverTab[112683]++
															if err := tm.writeAny(w, val, props.MapValProp); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:384
						_go_fuzz_dep_.CoverTab[112691]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:385
						// _ = "end of CoverTab[112691]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:386
						_go_fuzz_dep_.CoverTab[112692]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:386
						// _ = "end of CoverTab[112692]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:386
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:386
					// _ = "end of CoverTab[112683]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:386
					_go_fuzz_dep_.CoverTab[112684]++
															if err := w.WriteByte('\n'); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:387
						_go_fuzz_dep_.CoverTab[112693]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:388
						// _ = "end of CoverTab[112693]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:389
						_go_fuzz_dep_.CoverTab[112694]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:389
						// _ = "end of CoverTab[112694]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:389
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:389
					// _ = "end of CoverTab[112684]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:390
					_go_fuzz_dep_.CoverTab[112695]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:390
					// _ = "end of CoverTab[112695]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:390
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:390
				// _ = "end of CoverTab[112655]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:390
				_go_fuzz_dep_.CoverTab[112656]++

														w.unindent()
														if err := w.WriteByte('>'); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:393
					_go_fuzz_dep_.CoverTab[112696]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:394
					// _ = "end of CoverTab[112696]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:395
					_go_fuzz_dep_.CoverTab[112697]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:395
					// _ = "end of CoverTab[112697]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:395
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:395
				// _ = "end of CoverTab[112656]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:395
				_go_fuzz_dep_.CoverTab[112657]++
														if err := w.WriteByte('\n'); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:396
					_go_fuzz_dep_.CoverTab[112698]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:397
					// _ = "end of CoverTab[112698]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:398
					_go_fuzz_dep_.CoverTab[112699]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:398
					// _ = "end of CoverTab[112699]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:398
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:398
				// _ = "end of CoverTab[112657]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:399
			// _ = "end of CoverTab[112645]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:399
			_go_fuzz_dep_.CoverTab[112646]++
													continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:400
			// _ = "end of CoverTab[112646]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:401
			_go_fuzz_dep_.CoverTab[112700]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:401
			// _ = "end of CoverTab[112700]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:401
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:401
		// _ = "end of CoverTab[112592]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:401
		_go_fuzz_dep_.CoverTab[112593]++
												if props.proto3 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:402
			_go_fuzz_dep_.CoverTab[112701]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:402
			return fv.Kind() == reflect.Slice
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:402
			// _ = "end of CoverTab[112701]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:402
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:402
			_go_fuzz_dep_.CoverTab[112702]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:402
			return fv.Len() == 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:402
			// _ = "end of CoverTab[112702]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:402
		}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:402
			_go_fuzz_dep_.CoverTab[112703]++

													continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:404
			// _ = "end of CoverTab[112703]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:405
			_go_fuzz_dep_.CoverTab[112704]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:405
			// _ = "end of CoverTab[112704]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:405
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:405
		// _ = "end of CoverTab[112593]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:405
		_go_fuzz_dep_.CoverTab[112594]++
												if props.proto3 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:406
			_go_fuzz_dep_.CoverTab[112705]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:406
			return fv.Kind() != reflect.Ptr
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:406
			// _ = "end of CoverTab[112705]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:406
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:406
			_go_fuzz_dep_.CoverTab[112706]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:406
			return fv.Kind() != reflect.Slice
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:406
			// _ = "end of CoverTab[112706]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:406
		}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:406
			_go_fuzz_dep_.CoverTab[112707]++

													if isProto3Zero(fv) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:408
				_go_fuzz_dep_.CoverTab[112708]++
														continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:409
				// _ = "end of CoverTab[112708]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:410
				_go_fuzz_dep_.CoverTab[112709]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:410
				// _ = "end of CoverTab[112709]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:410
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:410
			// _ = "end of CoverTab[112707]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:411
			_go_fuzz_dep_.CoverTab[112710]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:411
			// _ = "end of CoverTab[112710]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:411
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:411
		// _ = "end of CoverTab[112594]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:411
		_go_fuzz_dep_.CoverTab[112595]++

												if fv.Kind() == reflect.Interface {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:413
			_go_fuzz_dep_.CoverTab[112711]++

													if st.Field(i).Tag.Get("protobuf_oneof") != "" {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:415
				_go_fuzz_dep_.CoverTab[112712]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:419
				if fv.IsNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:419
					_go_fuzz_dep_.CoverTab[112714]++
															continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:420
					// _ = "end of CoverTab[112714]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:421
					_go_fuzz_dep_.CoverTab[112715]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:421
					// _ = "end of CoverTab[112715]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:421
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:421
				// _ = "end of CoverTab[112712]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:421
				_go_fuzz_dep_.CoverTab[112713]++
														inner := fv.Elem().Elem()
														tag := inner.Type().Field(0).Tag.Get("protobuf")
														props = new(Properties)
														props.Parse(tag)

														fv = inner.Field(0)

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:432
				if fv.Kind() == reflect.Ptr && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:432
					_go_fuzz_dep_.CoverTab[112716]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:432
					return fv.IsNil()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:432
					// _ = "end of CoverTab[112716]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:432
				}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:432
					_go_fuzz_dep_.CoverTab[112717]++

															msg := errors.New("/* nil */")
															fv = reflect.ValueOf(&msg).Elem()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:435
					// _ = "end of CoverTab[112717]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:436
					_go_fuzz_dep_.CoverTab[112718]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:436
					// _ = "end of CoverTab[112718]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:436
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:436
				// _ = "end of CoverTab[112713]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:437
				_go_fuzz_dep_.CoverTab[112719]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:437
				// _ = "end of CoverTab[112719]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:437
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:437
			// _ = "end of CoverTab[112711]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:438
			_go_fuzz_dep_.CoverTab[112720]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:438
			// _ = "end of CoverTab[112720]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:438
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:438
		// _ = "end of CoverTab[112595]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:438
		_go_fuzz_dep_.CoverTab[112596]++

												if err := writeName(w, props); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:440
			_go_fuzz_dep_.CoverTab[112721]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:441
			// _ = "end of CoverTab[112721]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:442
			_go_fuzz_dep_.CoverTab[112722]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:442
			// _ = "end of CoverTab[112722]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:442
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:442
		// _ = "end of CoverTab[112596]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:442
		_go_fuzz_dep_.CoverTab[112597]++
												if !w.compact {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:443
			_go_fuzz_dep_.CoverTab[112723]++
													if err := w.WriteByte(' '); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:444
				_go_fuzz_dep_.CoverTab[112724]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:445
				// _ = "end of CoverTab[112724]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:446
				_go_fuzz_dep_.CoverTab[112725]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:446
				// _ = "end of CoverTab[112725]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:446
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:446
			// _ = "end of CoverTab[112723]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:447
			_go_fuzz_dep_.CoverTab[112726]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:447
			// _ = "end of CoverTab[112726]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:447
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:447
		// _ = "end of CoverTab[112597]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:447
		_go_fuzz_dep_.CoverTab[112598]++

												if len(props.Enum) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:449
			_go_fuzz_dep_.CoverTab[112727]++
													if err := tm.writeEnum(w, fv, props); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:450
				_go_fuzz_dep_.CoverTab[112728]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:451
				// _ = "end of CoverTab[112728]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:452
				_go_fuzz_dep_.CoverTab[112729]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:452
				// _ = "end of CoverTab[112729]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:452
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:452
			// _ = "end of CoverTab[112727]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:453
			_go_fuzz_dep_.CoverTab[112730]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:453
			if err := tm.writeAny(w, fv, props); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:453
				_go_fuzz_dep_.CoverTab[112731]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:454
				// _ = "end of CoverTab[112731]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:455
				_go_fuzz_dep_.CoverTab[112732]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:455
				// _ = "end of CoverTab[112732]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:455
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:455
			// _ = "end of CoverTab[112730]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:455
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:455
		// _ = "end of CoverTab[112598]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:455
		_go_fuzz_dep_.CoverTab[112599]++

												if err := w.WriteByte('\n'); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:457
			_go_fuzz_dep_.CoverTab[112733]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:458
			// _ = "end of CoverTab[112733]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:459
			_go_fuzz_dep_.CoverTab[112734]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:459
			// _ = "end of CoverTab[112734]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:459
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:459
		// _ = "end of CoverTab[112599]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:460
	// _ = "end of CoverTab[112578]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:460
	_go_fuzz_dep_.CoverTab[112579]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:463
	pv := sv
	if pv.CanAddr() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:464
		_go_fuzz_dep_.CoverTab[112735]++
												pv = sv.Addr()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:465
		// _ = "end of CoverTab[112735]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:466
		_go_fuzz_dep_.CoverTab[112736]++
												pv = reflect.New(sv.Type())
												pv.Elem().Set(sv)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:468
		// _ = "end of CoverTab[112736]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:469
	// _ = "end of CoverTab[112579]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:469
	_go_fuzz_dep_.CoverTab[112580]++
											if _, err := extendable(pv.Interface()); err == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:470
		_go_fuzz_dep_.CoverTab[112737]++
												if err := tm.writeExtensions(w, pv); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:471
			_go_fuzz_dep_.CoverTab[112738]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:472
			// _ = "end of CoverTab[112738]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:473
			_go_fuzz_dep_.CoverTab[112739]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:473
			// _ = "end of CoverTab[112739]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:473
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:473
		// _ = "end of CoverTab[112737]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:474
		_go_fuzz_dep_.CoverTab[112740]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:474
		// _ = "end of CoverTab[112740]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:474
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:474
	// _ = "end of CoverTab[112580]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:474
	_go_fuzz_dep_.CoverTab[112581]++

											return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:476
	// _ = "end of CoverTab[112581]"
}

var textMarshalerType = reflect.TypeOf((*encoding.TextMarshaler)(nil)).Elem()

// writeAny writes an arbitrary field.
func (tm *TextMarshaler) writeAny(w *textWriter, v reflect.Value, props *Properties) error {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:482
	_go_fuzz_dep_.CoverTab[112741]++
											v = reflect.Indirect(v)

											if props != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:485
		_go_fuzz_dep_.CoverTab[112745]++
												if len(props.CustomType) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:486
			_go_fuzz_dep_.CoverTab[112746]++
													custom, ok := v.Interface().(Marshaler)
													if ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:488
				_go_fuzz_dep_.CoverTab[112747]++
														data, err := custom.Marshal()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:490
					_go_fuzz_dep_.CoverTab[112750]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:491
					// _ = "end of CoverTab[112750]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:492
					_go_fuzz_dep_.CoverTab[112751]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:492
					// _ = "end of CoverTab[112751]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:492
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:492
				// _ = "end of CoverTab[112747]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:492
				_go_fuzz_dep_.CoverTab[112748]++
														if err := writeString(w, string(data)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:493
					_go_fuzz_dep_.CoverTab[112752]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:494
					// _ = "end of CoverTab[112752]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:495
					_go_fuzz_dep_.CoverTab[112753]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:495
					// _ = "end of CoverTab[112753]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:495
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:495
				// _ = "end of CoverTab[112748]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:495
				_go_fuzz_dep_.CoverTab[112749]++
														return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:496
				// _ = "end of CoverTab[112749]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:497
				_go_fuzz_dep_.CoverTab[112754]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:497
				// _ = "end of CoverTab[112754]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:497
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:497
			// _ = "end of CoverTab[112746]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:498
			_go_fuzz_dep_.CoverTab[112755]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:498
			if len(props.CastType) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:498
				_go_fuzz_dep_.CoverTab[112756]++
														if _, ok := v.Interface().(interface {
					String() string
				}); ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:501
					_go_fuzz_dep_.CoverTab[112757]++
															switch v.Kind() {
					case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
						reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:504
						_go_fuzz_dep_.CoverTab[112758]++
																_, err := fmt.Fprintf(w, "%d", v.Interface())
																return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:506
						// _ = "end of CoverTab[112758]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:506
					default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:506
						_go_fuzz_dep_.CoverTab[112759]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:506
						// _ = "end of CoverTab[112759]"
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:507
					// _ = "end of CoverTab[112757]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:508
					_go_fuzz_dep_.CoverTab[112760]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:508
					// _ = "end of CoverTab[112760]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:508
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:508
				// _ = "end of CoverTab[112756]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:509
				_go_fuzz_dep_.CoverTab[112761]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:509
				if props.StdTime {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:509
					_go_fuzz_dep_.CoverTab[112762]++
															t, ok := v.Interface().(time.Time)
															if !ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:511
						_go_fuzz_dep_.CoverTab[112765]++
																return fmt.Errorf("stdtime is not time.Time, but %T", v.Interface())
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:512
						// _ = "end of CoverTab[112765]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:513
						_go_fuzz_dep_.CoverTab[112766]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:513
						// _ = "end of CoverTab[112766]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:513
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:513
					// _ = "end of CoverTab[112762]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:513
					_go_fuzz_dep_.CoverTab[112763]++
															tproto, err := timestampProto(t)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:515
						_go_fuzz_dep_.CoverTab[112767]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:516
						// _ = "end of CoverTab[112767]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:517
						_go_fuzz_dep_.CoverTab[112768]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:517
						// _ = "end of CoverTab[112768]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:517
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:517
					// _ = "end of CoverTab[112763]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:517
					_go_fuzz_dep_.CoverTab[112764]++
															propsCopy := *props
															propsCopy.StdTime = false
															err = tm.writeAny(w, reflect.ValueOf(tproto), &propsCopy)
															return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:521
					// _ = "end of CoverTab[112764]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:522
					_go_fuzz_dep_.CoverTab[112769]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:522
					if props.StdDuration {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:522
						_go_fuzz_dep_.CoverTab[112770]++
																d, ok := v.Interface().(time.Duration)
																if !ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:524
							_go_fuzz_dep_.CoverTab[112772]++
																	return fmt.Errorf("stdtime is not time.Duration, but %T", v.Interface())
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:525
							// _ = "end of CoverTab[112772]"
						} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:526
							_go_fuzz_dep_.CoverTab[112773]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:526
							// _ = "end of CoverTab[112773]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:526
						}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:526
						// _ = "end of CoverTab[112770]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:526
						_go_fuzz_dep_.CoverTab[112771]++
																dproto := durationProto(d)
																propsCopy := *props
																propsCopy.StdDuration = false
																err := tm.writeAny(w, reflect.ValueOf(dproto), &propsCopy)
																return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:531
						// _ = "end of CoverTab[112771]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:532
						_go_fuzz_dep_.CoverTab[112774]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:532
						// _ = "end of CoverTab[112774]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:532
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:532
					// _ = "end of CoverTab[112769]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:532
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:532
				// _ = "end of CoverTab[112761]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:532
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:532
			// _ = "end of CoverTab[112755]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:532
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:532
		// _ = "end of CoverTab[112745]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:533
		_go_fuzz_dep_.CoverTab[112775]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:533
		// _ = "end of CoverTab[112775]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:533
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:533
	// _ = "end of CoverTab[112741]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:533
	_go_fuzz_dep_.CoverTab[112742]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:536
	if v.Kind() == reflect.Float32 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:536
		_go_fuzz_dep_.CoverTab[112776]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:536
		return v.Kind() == reflect.Float64
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:536
		// _ = "end of CoverTab[112776]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:536
	}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:536
		_go_fuzz_dep_.CoverTab[112777]++
												x := v.Float()
												var b []byte
												switch {
		case math.IsInf(x, 1):
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:540
			_go_fuzz_dep_.CoverTab[112779]++
													b = posInf
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:541
			// _ = "end of CoverTab[112779]"
		case math.IsInf(x, -1):
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:542
			_go_fuzz_dep_.CoverTab[112780]++
													b = negInf
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:543
			// _ = "end of CoverTab[112780]"
		case math.IsNaN(x):
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:544
			_go_fuzz_dep_.CoverTab[112781]++
													b = nan
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:545
			// _ = "end of CoverTab[112781]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:545
		default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:545
			_go_fuzz_dep_.CoverTab[112782]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:545
			// _ = "end of CoverTab[112782]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:546
		// _ = "end of CoverTab[112777]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:546
		_go_fuzz_dep_.CoverTab[112778]++
												if b != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:547
			_go_fuzz_dep_.CoverTab[112783]++
													_, err := w.Write(b)
													return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:549
			// _ = "end of CoverTab[112783]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:550
			_go_fuzz_dep_.CoverTab[112784]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:550
			// _ = "end of CoverTab[112784]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:550
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:550
		// _ = "end of CoverTab[112778]"

	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:552
		_go_fuzz_dep_.CoverTab[112785]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:552
		// _ = "end of CoverTab[112785]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:552
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:552
	// _ = "end of CoverTab[112742]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:552
	_go_fuzz_dep_.CoverTab[112743]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:556
	switch v.Kind() {
	case reflect.Slice:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:557
		_go_fuzz_dep_.CoverTab[112786]++

												if err := writeString(w, string(v.Bytes())); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:559
			_go_fuzz_dep_.CoverTab[112795]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:560
			// _ = "end of CoverTab[112795]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:561
			_go_fuzz_dep_.CoverTab[112796]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:561
			// _ = "end of CoverTab[112796]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:561
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:561
		// _ = "end of CoverTab[112786]"
	case reflect.String:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:562
		_go_fuzz_dep_.CoverTab[112787]++
												if err := writeString(w, v.String()); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:563
			_go_fuzz_dep_.CoverTab[112797]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:564
			// _ = "end of CoverTab[112797]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:565
			_go_fuzz_dep_.CoverTab[112798]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:565
			// _ = "end of CoverTab[112798]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:565
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:565
		// _ = "end of CoverTab[112787]"
	case reflect.Struct:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:566
		_go_fuzz_dep_.CoverTab[112788]++
		// Required/optional group/message.
		var bra, ket byte = '<', '>'
		if props != nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:569
			_go_fuzz_dep_.CoverTab[112799]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:569
			return props.Wire == "group"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:569
			// _ = "end of CoverTab[112799]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:569
		}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:569
			_go_fuzz_dep_.CoverTab[112800]++
													bra, ket = '{', '}'
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:570
			// _ = "end of CoverTab[112800]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:571
			_go_fuzz_dep_.CoverTab[112801]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:571
			// _ = "end of CoverTab[112801]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:571
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:571
		// _ = "end of CoverTab[112788]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:571
		_go_fuzz_dep_.CoverTab[112789]++
												if err := w.WriteByte(bra); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:572
			_go_fuzz_dep_.CoverTab[112802]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:573
			// _ = "end of CoverTab[112802]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:574
			_go_fuzz_dep_.CoverTab[112803]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:574
			// _ = "end of CoverTab[112803]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:574
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:574
		// _ = "end of CoverTab[112789]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:574
		_go_fuzz_dep_.CoverTab[112790]++
												if !w.compact {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:575
			_go_fuzz_dep_.CoverTab[112804]++
													if err := w.WriteByte('\n'); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:576
				_go_fuzz_dep_.CoverTab[112805]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:577
				// _ = "end of CoverTab[112805]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:578
				_go_fuzz_dep_.CoverTab[112806]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:578
				// _ = "end of CoverTab[112806]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:578
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:578
			// _ = "end of CoverTab[112804]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:579
			_go_fuzz_dep_.CoverTab[112807]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:579
			// _ = "end of CoverTab[112807]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:579
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:579
		// _ = "end of CoverTab[112790]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:579
		_go_fuzz_dep_.CoverTab[112791]++
												w.indent()
												if v.CanAddr() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:581
			_go_fuzz_dep_.CoverTab[112808]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:592
			v = v.Addr()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:592
			// _ = "end of CoverTab[112808]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:593
			_go_fuzz_dep_.CoverTab[112809]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:593
			// _ = "end of CoverTab[112809]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:593
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:593
		// _ = "end of CoverTab[112791]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:593
		_go_fuzz_dep_.CoverTab[112792]++
												if v.Type().Implements(textMarshalerType) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:594
			_go_fuzz_dep_.CoverTab[112810]++
													text, err := v.Interface().(encoding.TextMarshaler).MarshalText()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:596
				_go_fuzz_dep_.CoverTab[112812]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:597
				// _ = "end of CoverTab[112812]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:598
				_go_fuzz_dep_.CoverTab[112813]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:598
				// _ = "end of CoverTab[112813]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:598
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:598
			// _ = "end of CoverTab[112810]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:598
			_go_fuzz_dep_.CoverTab[112811]++
													if _, err = w.Write(text); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:599
				_go_fuzz_dep_.CoverTab[112814]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:600
				// _ = "end of CoverTab[112814]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:601
				_go_fuzz_dep_.CoverTab[112815]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:601
				// _ = "end of CoverTab[112815]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:601
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:601
			// _ = "end of CoverTab[112811]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:602
			_go_fuzz_dep_.CoverTab[112816]++
													if v.Kind() == reflect.Ptr {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:603
				_go_fuzz_dep_.CoverTab[112818]++
														v = v.Elem()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:604
				// _ = "end of CoverTab[112818]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:605
				_go_fuzz_dep_.CoverTab[112819]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:605
				// _ = "end of CoverTab[112819]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:605
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:605
			// _ = "end of CoverTab[112816]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:605
			_go_fuzz_dep_.CoverTab[112817]++
													if err := tm.writeStruct(w, v); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:606
				_go_fuzz_dep_.CoverTab[112820]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:607
				// _ = "end of CoverTab[112820]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:608
				_go_fuzz_dep_.CoverTab[112821]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:608
				// _ = "end of CoverTab[112821]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:608
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:608
			// _ = "end of CoverTab[112817]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:609
		// _ = "end of CoverTab[112792]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:609
		_go_fuzz_dep_.CoverTab[112793]++
												w.unindent()
												if err := w.WriteByte(ket); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:611
			_go_fuzz_dep_.CoverTab[112822]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:612
			// _ = "end of CoverTab[112822]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:613
			_go_fuzz_dep_.CoverTab[112823]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:613
			// _ = "end of CoverTab[112823]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:613
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:613
		// _ = "end of CoverTab[112793]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:614
		_go_fuzz_dep_.CoverTab[112794]++
												_, err := fmt.Fprint(w, v.Interface())
												return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:616
		// _ = "end of CoverTab[112794]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:617
	// _ = "end of CoverTab[112743]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:617
	_go_fuzz_dep_.CoverTab[112744]++
											return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:618
	// _ = "end of CoverTab[112744]"
}

// equivalent to C's isprint.
func isprint(c byte) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:622
	_go_fuzz_dep_.CoverTab[112824]++
											return c >= 0x20 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:623
		_go_fuzz_dep_.CoverTab[112825]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:623
		return c < 0x7f
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:623
		// _ = "end of CoverTab[112825]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:623
	}()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:623
	// _ = "end of CoverTab[112824]"
}

// writeString writes a string in the protocol buffer text format.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:626
// It is similar to strconv.Quote except we don't use Go escape sequences,
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:626
// we treat the string as a byte sequence, and we use octal escapes.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:626
// These differences are to maintain interoperability with the other
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:626
// languages' implementations of the text format.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:631
func writeString(w *textWriter, s string) error {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:631
	_go_fuzz_dep_.CoverTab[112826]++

											if err := w.WriteByte('"'); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:633
		_go_fuzz_dep_.CoverTab[112829]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:634
		// _ = "end of CoverTab[112829]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:635
		_go_fuzz_dep_.CoverTab[112830]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:635
		// _ = "end of CoverTab[112830]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:635
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:635
	// _ = "end of CoverTab[112826]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:635
	_go_fuzz_dep_.CoverTab[112827]++

											for i := 0; i < len(s); i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:637
		_go_fuzz_dep_.CoverTab[112831]++
												var err error

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:642
		switch c := s[i]; c {
		case '\n':
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:643
			_go_fuzz_dep_.CoverTab[112833]++
													_, err = w.w.Write(backslashN)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:644
			// _ = "end of CoverTab[112833]"
		case '\r':
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:645
			_go_fuzz_dep_.CoverTab[112834]++
													_, err = w.w.Write(backslashR)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:646
			// _ = "end of CoverTab[112834]"
		case '\t':
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:647
			_go_fuzz_dep_.CoverTab[112835]++
													_, err = w.w.Write(backslashT)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:648
			// _ = "end of CoverTab[112835]"
		case '"':
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:649
			_go_fuzz_dep_.CoverTab[112836]++
													_, err = w.w.Write(backslashDQ)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:650
			// _ = "end of CoverTab[112836]"
		case '\\':
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:651
			_go_fuzz_dep_.CoverTab[112837]++
													_, err = w.w.Write(backslashBS)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:652
			// _ = "end of CoverTab[112837]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:653
			_go_fuzz_dep_.CoverTab[112838]++
													if isprint(c) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:654
				_go_fuzz_dep_.CoverTab[112839]++
														err = w.w.WriteByte(c)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:655
				// _ = "end of CoverTab[112839]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:656
				_go_fuzz_dep_.CoverTab[112840]++
														_, err = fmt.Fprintf(w.w, "\\%03o", c)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:657
				// _ = "end of CoverTab[112840]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:658
			// _ = "end of CoverTab[112838]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:659
		// _ = "end of CoverTab[112831]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:659
		_go_fuzz_dep_.CoverTab[112832]++
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:660
			_go_fuzz_dep_.CoverTab[112841]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:661
			// _ = "end of CoverTab[112841]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:662
			_go_fuzz_dep_.CoverTab[112842]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:662
			// _ = "end of CoverTab[112842]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:662
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:662
		// _ = "end of CoverTab[112832]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:663
	// _ = "end of CoverTab[112827]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:663
	_go_fuzz_dep_.CoverTab[112828]++
											return w.WriteByte('"')
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:664
	// _ = "end of CoverTab[112828]"
}

func writeUnknownStruct(w *textWriter, data []byte) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:667
	_go_fuzz_dep_.CoverTab[112843]++
											if !w.compact {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:668
		_go_fuzz_dep_.CoverTab[112846]++
												if _, err := fmt.Fprintf(w, "/* %d unknown bytes */\n", len(data)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:669
			_go_fuzz_dep_.CoverTab[112847]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:670
			// _ = "end of CoverTab[112847]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:671
			_go_fuzz_dep_.CoverTab[112848]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:671
			// _ = "end of CoverTab[112848]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:671
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:671
		// _ = "end of CoverTab[112846]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:672
		_go_fuzz_dep_.CoverTab[112849]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:672
		// _ = "end of CoverTab[112849]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:672
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:672
	// _ = "end of CoverTab[112843]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:672
	_go_fuzz_dep_.CoverTab[112844]++
											b := NewBuffer(data)
											for b.index < len(b.buf) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:674
		_go_fuzz_dep_.CoverTab[112850]++
												x, err := b.DecodeVarint()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:676
			_go_fuzz_dep_.CoverTab[112858]++
													_, ferr := fmt.Fprintf(w, "/* %v */\n", err)
													return ferr
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:678
			// _ = "end of CoverTab[112858]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:679
			_go_fuzz_dep_.CoverTab[112859]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:679
			// _ = "end of CoverTab[112859]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:679
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:679
		// _ = "end of CoverTab[112850]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:679
		_go_fuzz_dep_.CoverTab[112851]++
												wire, tag := x&7, x>>3
												if wire == WireEndGroup {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:681
			_go_fuzz_dep_.CoverTab[112860]++
													w.unindent()
													if _, werr := w.Write(endBraceNewline); werr != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:683
				_go_fuzz_dep_.CoverTab[112862]++
														return werr
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:684
				// _ = "end of CoverTab[112862]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:685
				_go_fuzz_dep_.CoverTab[112863]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:685
				// _ = "end of CoverTab[112863]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:685
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:685
			// _ = "end of CoverTab[112860]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:685
			_go_fuzz_dep_.CoverTab[112861]++
													continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:686
			// _ = "end of CoverTab[112861]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:687
			_go_fuzz_dep_.CoverTab[112864]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:687
			// _ = "end of CoverTab[112864]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:687
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:687
		// _ = "end of CoverTab[112851]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:687
		_go_fuzz_dep_.CoverTab[112852]++
												if _, ferr := fmt.Fprint(w, tag); ferr != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:688
			_go_fuzz_dep_.CoverTab[112865]++
													return ferr
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:689
			// _ = "end of CoverTab[112865]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:690
			_go_fuzz_dep_.CoverTab[112866]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:690
			// _ = "end of CoverTab[112866]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:690
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:690
		// _ = "end of CoverTab[112852]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:690
		_go_fuzz_dep_.CoverTab[112853]++
												if wire != WireStartGroup {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:691
			_go_fuzz_dep_.CoverTab[112867]++
													if err = w.WriteByte(':'); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:692
				_go_fuzz_dep_.CoverTab[112868]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:693
				// _ = "end of CoverTab[112868]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:694
				_go_fuzz_dep_.CoverTab[112869]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:694
				// _ = "end of CoverTab[112869]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:694
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:694
			// _ = "end of CoverTab[112867]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:695
			_go_fuzz_dep_.CoverTab[112870]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:695
			// _ = "end of CoverTab[112870]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:695
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:695
		// _ = "end of CoverTab[112853]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:695
		_go_fuzz_dep_.CoverTab[112854]++
												if !w.compact || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:696
			_go_fuzz_dep_.CoverTab[112871]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:696
			return wire == WireStartGroup
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:696
			// _ = "end of CoverTab[112871]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:696
		}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:696
			_go_fuzz_dep_.CoverTab[112872]++
													if err = w.WriteByte(' '); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:697
				_go_fuzz_dep_.CoverTab[112873]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:698
				// _ = "end of CoverTab[112873]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:699
				_go_fuzz_dep_.CoverTab[112874]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:699
				// _ = "end of CoverTab[112874]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:699
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:699
			// _ = "end of CoverTab[112872]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:700
			_go_fuzz_dep_.CoverTab[112875]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:700
			// _ = "end of CoverTab[112875]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:700
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:700
		// _ = "end of CoverTab[112854]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:700
		_go_fuzz_dep_.CoverTab[112855]++
												switch wire {
		case WireBytes:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:702
			_go_fuzz_dep_.CoverTab[112876]++
													buf, e := b.DecodeRawBytes(false)
													if e == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:704
				_go_fuzz_dep_.CoverTab[112882]++
														_, err = fmt.Fprintf(w, "%q", buf)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:705
				// _ = "end of CoverTab[112882]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:706
				_go_fuzz_dep_.CoverTab[112883]++
														_, err = fmt.Fprintf(w, "/* %v */", e)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:707
				// _ = "end of CoverTab[112883]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:708
			// _ = "end of CoverTab[112876]"
		case WireFixed32:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:709
			_go_fuzz_dep_.CoverTab[112877]++
													x, err = b.DecodeFixed32()
													err = writeUnknownInt(w, x, err)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:711
			// _ = "end of CoverTab[112877]"
		case WireFixed64:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:712
			_go_fuzz_dep_.CoverTab[112878]++
													x, err = b.DecodeFixed64()
													err = writeUnknownInt(w, x, err)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:714
			// _ = "end of CoverTab[112878]"
		case WireStartGroup:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:715
			_go_fuzz_dep_.CoverTab[112879]++
													err = w.WriteByte('{')
													w.indent()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:717
			// _ = "end of CoverTab[112879]"
		case WireVarint:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:718
			_go_fuzz_dep_.CoverTab[112880]++
													x, err = b.DecodeVarint()
													err = writeUnknownInt(w, x, err)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:720
			// _ = "end of CoverTab[112880]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:721
			_go_fuzz_dep_.CoverTab[112881]++
													_, err = fmt.Fprintf(w, "/* unknown wire type %d */", wire)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:722
			// _ = "end of CoverTab[112881]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:723
		// _ = "end of CoverTab[112855]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:723
		_go_fuzz_dep_.CoverTab[112856]++
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:724
			_go_fuzz_dep_.CoverTab[112884]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:725
			// _ = "end of CoverTab[112884]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:726
			_go_fuzz_dep_.CoverTab[112885]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:726
			// _ = "end of CoverTab[112885]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:726
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:726
		// _ = "end of CoverTab[112856]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:726
		_go_fuzz_dep_.CoverTab[112857]++
												if err := w.WriteByte('\n'); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:727
			_go_fuzz_dep_.CoverTab[112886]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:728
			// _ = "end of CoverTab[112886]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:729
			_go_fuzz_dep_.CoverTab[112887]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:729
			// _ = "end of CoverTab[112887]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:729
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:729
		// _ = "end of CoverTab[112857]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:730
	// _ = "end of CoverTab[112844]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:730
	_go_fuzz_dep_.CoverTab[112845]++
											return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:731
	// _ = "end of CoverTab[112845]"
}

func writeUnknownInt(w *textWriter, x uint64, err error) error {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:734
	_go_fuzz_dep_.CoverTab[112888]++
											if err == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:735
		_go_fuzz_dep_.CoverTab[112890]++
												_, err = fmt.Fprint(w, x)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:736
		// _ = "end of CoverTab[112890]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:737
		_go_fuzz_dep_.CoverTab[112891]++
												_, err = fmt.Fprintf(w, "/* %v */", err)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:738
		// _ = "end of CoverTab[112891]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:739
	// _ = "end of CoverTab[112888]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:739
	_go_fuzz_dep_.CoverTab[112889]++
											return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:740
	// _ = "end of CoverTab[112889]"
}

type int32Slice []int32

func (s int32Slice) Len() int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:745
	_go_fuzz_dep_.CoverTab[112892]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:745
	return len(s)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:745
	// _ = "end of CoverTab[112892]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:745
}
func (s int32Slice) Less(i, j int) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:746
	_go_fuzz_dep_.CoverTab[112893]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:746
	return s[i] < s[j]
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:746
	// _ = "end of CoverTab[112893]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:746
}
func (s int32Slice) Swap(i, j int) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:747
	_go_fuzz_dep_.CoverTab[112894]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:747
	s[i], s[j] = s[j], s[i]
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:747
	// _ = "end of CoverTab[112894]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:747
}

// writeExtensions writes all the extensions in pv.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:749
// pv is assumed to be a pointer to a protocol message struct that is extendable.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:751
func (tm *TextMarshaler) writeExtensions(w *textWriter, pv reflect.Value) error {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:751
	_go_fuzz_dep_.CoverTab[112895]++
											emap := extensionMaps[pv.Type().Elem()]
											e := pv.Interface().(Message)

											var m map[int32]Extension
											var mu sync.Locker
											if em, ok := e.(extensionsBytes); ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:757
		_go_fuzz_dep_.CoverTab[112899]++
												eb := em.GetExtensions()
												var err error
												m, err = BytesToExtensionsMap(*eb)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:761
			_go_fuzz_dep_.CoverTab[112901]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:762
			// _ = "end of CoverTab[112901]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:763
			_go_fuzz_dep_.CoverTab[112902]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:763
			// _ = "end of CoverTab[112902]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:763
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:763
		// _ = "end of CoverTab[112899]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:763
		_go_fuzz_dep_.CoverTab[112900]++
												mu = notLocker{}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:764
		// _ = "end of CoverTab[112900]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:765
		_go_fuzz_dep_.CoverTab[112903]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:765
		if _, ok := e.(extendableProto); ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:765
			_go_fuzz_dep_.CoverTab[112904]++
													ep, _ := extendable(e)
													m, mu = ep.extensionsRead()
													if m == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:768
				_go_fuzz_dep_.CoverTab[112905]++
														return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:769
				// _ = "end of CoverTab[112905]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:770
				_go_fuzz_dep_.CoverTab[112906]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:770
				// _ = "end of CoverTab[112906]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:770
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:770
			// _ = "end of CoverTab[112904]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:771
			_go_fuzz_dep_.CoverTab[112907]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:771
			// _ = "end of CoverTab[112907]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:771
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:771
		// _ = "end of CoverTab[112903]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:771
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:771
	// _ = "end of CoverTab[112895]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:771
	_go_fuzz_dep_.CoverTab[112896]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:777
	mu.Lock()
	ids := make([]int32, 0, len(m))
	for id := range m {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:779
		_go_fuzz_dep_.CoverTab[112908]++
												ids = append(ids, id)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:780
		// _ = "end of CoverTab[112908]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:781
	// _ = "end of CoverTab[112896]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:781
	_go_fuzz_dep_.CoverTab[112897]++
											sort.Sort(int32Slice(ids))
											mu.Unlock()

											for _, extNum := range ids {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:785
		_go_fuzz_dep_.CoverTab[112909]++
												ext := m[extNum]
												var desc *ExtensionDesc
												if emap != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:788
			_go_fuzz_dep_.CoverTab[112913]++
													desc = emap[extNum]
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:789
			// _ = "end of CoverTab[112913]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:790
			_go_fuzz_dep_.CoverTab[112914]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:790
			// _ = "end of CoverTab[112914]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:790
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:790
		// _ = "end of CoverTab[112909]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:790
		_go_fuzz_dep_.CoverTab[112910]++
												if desc == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:791
			_go_fuzz_dep_.CoverTab[112915]++

													if err := writeUnknownStruct(w, ext.enc); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:793
				_go_fuzz_dep_.CoverTab[112917]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:794
				// _ = "end of CoverTab[112917]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:795
				_go_fuzz_dep_.CoverTab[112918]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:795
				// _ = "end of CoverTab[112918]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:795
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:795
			// _ = "end of CoverTab[112915]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:795
			_go_fuzz_dep_.CoverTab[112916]++
													continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:796
			// _ = "end of CoverTab[112916]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:797
			_go_fuzz_dep_.CoverTab[112919]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:797
			// _ = "end of CoverTab[112919]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:797
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:797
		// _ = "end of CoverTab[112910]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:797
		_go_fuzz_dep_.CoverTab[112911]++

												pb, err := GetExtension(e, desc)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:800
			_go_fuzz_dep_.CoverTab[112920]++
													return fmt.Errorf("failed getting extension: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:801
			// _ = "end of CoverTab[112920]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:802
			_go_fuzz_dep_.CoverTab[112921]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:802
			// _ = "end of CoverTab[112921]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:802
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:802
		// _ = "end of CoverTab[112911]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:802
		_go_fuzz_dep_.CoverTab[112912]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:805
		if !desc.repeated() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:805
			_go_fuzz_dep_.CoverTab[112922]++
													if err := tm.writeExtension(w, desc.Name, pb); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:806
				_go_fuzz_dep_.CoverTab[112923]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:807
				// _ = "end of CoverTab[112923]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:808
				_go_fuzz_dep_.CoverTab[112924]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:808
				// _ = "end of CoverTab[112924]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:808
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:808
			// _ = "end of CoverTab[112922]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:809
			_go_fuzz_dep_.CoverTab[112925]++
													v := reflect.ValueOf(pb)
													for i := 0; i < v.Len(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:811
				_go_fuzz_dep_.CoverTab[112926]++
														if err := tm.writeExtension(w, desc.Name, v.Index(i).Interface()); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:812
					_go_fuzz_dep_.CoverTab[112927]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:813
					// _ = "end of CoverTab[112927]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:814
					_go_fuzz_dep_.CoverTab[112928]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:814
					// _ = "end of CoverTab[112928]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:814
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:814
				// _ = "end of CoverTab[112926]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:815
			// _ = "end of CoverTab[112925]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:816
		// _ = "end of CoverTab[112912]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:817
	// _ = "end of CoverTab[112897]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:817
	_go_fuzz_dep_.CoverTab[112898]++
											return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:818
	// _ = "end of CoverTab[112898]"
}

func (tm *TextMarshaler) writeExtension(w *textWriter, name string, pb interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:821
	_go_fuzz_dep_.CoverTab[112929]++
											if _, err := fmt.Fprintf(w, "[%s]:", name); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:822
		_go_fuzz_dep_.CoverTab[112934]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:823
		// _ = "end of CoverTab[112934]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:824
		_go_fuzz_dep_.CoverTab[112935]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:824
		// _ = "end of CoverTab[112935]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:824
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:824
	// _ = "end of CoverTab[112929]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:824
	_go_fuzz_dep_.CoverTab[112930]++
											if !w.compact {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:825
		_go_fuzz_dep_.CoverTab[112936]++
												if err := w.WriteByte(' '); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:826
			_go_fuzz_dep_.CoverTab[112937]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:827
			// _ = "end of CoverTab[112937]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:828
			_go_fuzz_dep_.CoverTab[112938]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:828
			// _ = "end of CoverTab[112938]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:828
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:828
		// _ = "end of CoverTab[112936]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:829
		_go_fuzz_dep_.CoverTab[112939]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:829
		// _ = "end of CoverTab[112939]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:829
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:829
	// _ = "end of CoverTab[112930]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:829
	_go_fuzz_dep_.CoverTab[112931]++
											if err := tm.writeAny(w, reflect.ValueOf(pb), nil); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:830
		_go_fuzz_dep_.CoverTab[112940]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:831
		// _ = "end of CoverTab[112940]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:832
		_go_fuzz_dep_.CoverTab[112941]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:832
		// _ = "end of CoverTab[112941]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:832
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:832
	// _ = "end of CoverTab[112931]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:832
	_go_fuzz_dep_.CoverTab[112932]++
											if err := w.WriteByte('\n'); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:833
		_go_fuzz_dep_.CoverTab[112942]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:834
		// _ = "end of CoverTab[112942]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:835
		_go_fuzz_dep_.CoverTab[112943]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:835
		// _ = "end of CoverTab[112943]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:835
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:835
	// _ = "end of CoverTab[112932]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:835
	_go_fuzz_dep_.CoverTab[112933]++
											return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:836
	// _ = "end of CoverTab[112933]"
}

func (w *textWriter) writeIndent() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:839
	_go_fuzz_dep_.CoverTab[112944]++
											if !w.complete {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:840
		_go_fuzz_dep_.CoverTab[112947]++
												return
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:841
		// _ = "end of CoverTab[112947]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:842
		_go_fuzz_dep_.CoverTab[112948]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:842
		// _ = "end of CoverTab[112948]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:842
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:842
	// _ = "end of CoverTab[112944]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:842
	_go_fuzz_dep_.CoverTab[112945]++
											remain := w.ind * 2
											for remain > 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:844
		_go_fuzz_dep_.CoverTab[112949]++
												n := remain
												if n > len(spaces) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:846
			_go_fuzz_dep_.CoverTab[112951]++
													n = len(spaces)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:847
			// _ = "end of CoverTab[112951]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:848
			_go_fuzz_dep_.CoverTab[112952]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:848
			// _ = "end of CoverTab[112952]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:848
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:848
		// _ = "end of CoverTab[112949]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:848
		_go_fuzz_dep_.CoverTab[112950]++
												w.w.Write(spaces[:n])
												remain -= n
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:850
		// _ = "end of CoverTab[112950]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:851
	// _ = "end of CoverTab[112945]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:851
	_go_fuzz_dep_.CoverTab[112946]++
											w.complete = false
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:852
	// _ = "end of CoverTab[112946]"
}

// TextMarshaler is a configurable text format marshaler.
type TextMarshaler struct {
	Compact		bool	// use compact text format (one line).
	ExpandAny	bool	// expand google.protobuf.Any messages of known types
}

// Marshal writes a given protocol buffer in text format.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:861
// The only errors returned are from w.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:863
func (tm *TextMarshaler) Marshal(w io.Writer, pb Message) error {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:863
	_go_fuzz_dep_.CoverTab[112953]++
											val := reflect.ValueOf(pb)
											if pb == nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:865
		_go_fuzz_dep_.CoverTab[112959]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:865
		return val.IsNil()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:865
		// _ = "end of CoverTab[112959]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:865
	}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:865
		_go_fuzz_dep_.CoverTab[112960]++
												w.Write([]byte("<nil>"))
												return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:867
		// _ = "end of CoverTab[112960]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:868
		_go_fuzz_dep_.CoverTab[112961]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:868
		// _ = "end of CoverTab[112961]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:868
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:868
	// _ = "end of CoverTab[112953]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:868
	_go_fuzz_dep_.CoverTab[112954]++
											var bw *bufio.Writer
											ww, ok := w.(writer)
											if !ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:871
		_go_fuzz_dep_.CoverTab[112962]++
												bw = bufio.NewWriter(w)
												ww = bw
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:873
		// _ = "end of CoverTab[112962]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:874
		_go_fuzz_dep_.CoverTab[112963]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:874
		// _ = "end of CoverTab[112963]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:874
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:874
	// _ = "end of CoverTab[112954]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:874
	_go_fuzz_dep_.CoverTab[112955]++
											aw := &textWriter{
		w:		ww,
		complete:	true,
		compact:	tm.Compact,
	}

	if etm, ok := pb.(encoding.TextMarshaler); ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:881
		_go_fuzz_dep_.CoverTab[112964]++
												text, err := etm.MarshalText()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:883
			_go_fuzz_dep_.CoverTab[112968]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:884
			// _ = "end of CoverTab[112968]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:885
			_go_fuzz_dep_.CoverTab[112969]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:885
			// _ = "end of CoverTab[112969]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:885
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:885
		// _ = "end of CoverTab[112964]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:885
		_go_fuzz_dep_.CoverTab[112965]++
												if _, err = aw.Write(text); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:886
			_go_fuzz_dep_.CoverTab[112970]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:887
			// _ = "end of CoverTab[112970]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:888
			_go_fuzz_dep_.CoverTab[112971]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:888
			// _ = "end of CoverTab[112971]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:888
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:888
		// _ = "end of CoverTab[112965]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:888
		_go_fuzz_dep_.CoverTab[112966]++
												if bw != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:889
			_go_fuzz_dep_.CoverTab[112972]++
													return bw.Flush()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:890
			// _ = "end of CoverTab[112972]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:891
			_go_fuzz_dep_.CoverTab[112973]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:891
			// _ = "end of CoverTab[112973]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:891
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:891
		// _ = "end of CoverTab[112966]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:891
		_go_fuzz_dep_.CoverTab[112967]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:892
		// _ = "end of CoverTab[112967]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:893
		_go_fuzz_dep_.CoverTab[112974]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:893
		// _ = "end of CoverTab[112974]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:893
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:893
	// _ = "end of CoverTab[112955]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:893
	_go_fuzz_dep_.CoverTab[112956]++

											v := reflect.Indirect(val)
											if err := tm.writeStruct(aw, v); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:896
		_go_fuzz_dep_.CoverTab[112975]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:897
		// _ = "end of CoverTab[112975]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:898
		_go_fuzz_dep_.CoverTab[112976]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:898
		// _ = "end of CoverTab[112976]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:898
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:898
	// _ = "end of CoverTab[112956]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:898
	_go_fuzz_dep_.CoverTab[112957]++
											if bw != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:899
		_go_fuzz_dep_.CoverTab[112977]++
												return bw.Flush()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:900
		// _ = "end of CoverTab[112977]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:901
		_go_fuzz_dep_.CoverTab[112978]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:901
		// _ = "end of CoverTab[112978]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:901
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:901
	// _ = "end of CoverTab[112957]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:901
	_go_fuzz_dep_.CoverTab[112958]++
											return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:902
	// _ = "end of CoverTab[112958]"
}

// Text is the same as Marshal, but returns the string directly.
func (tm *TextMarshaler) Text(pb Message) string {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:906
	_go_fuzz_dep_.CoverTab[112979]++
											var buf bytes.Buffer
											tm.Marshal(&buf, pb)
											return buf.String()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:909
	// _ = "end of CoverTab[112979]"
}

var (
	defaultTextMarshaler	= TextMarshaler{}
	compactTextMarshaler	= TextMarshaler{Compact: true}
)

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:919
// MarshalText writes a given protocol buffer in text format.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:919
// The only errors returned are from w.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:921
func MarshalText(w io.Writer, pb Message) error {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:921
	_go_fuzz_dep_.CoverTab[112980]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:921
	return defaultTextMarshaler.Marshal(w, pb)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:921
	// _ = "end of CoverTab[112980]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:921
}

// MarshalTextString is the same as MarshalText, but returns the string directly.
func MarshalTextString(pb Message) string {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:924
	_go_fuzz_dep_.CoverTab[112981]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:924
	return defaultTextMarshaler.Text(pb)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:924
	// _ = "end of CoverTab[112981]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:924
}

// CompactText writes a given protocol buffer in compact text format (one line).
func CompactText(w io.Writer, pb Message) error {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:927
	_go_fuzz_dep_.CoverTab[112982]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:927
	return compactTextMarshaler.Marshal(w, pb)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:927
	// _ = "end of CoverTab[112982]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:927
}

// CompactTextString is the same as CompactText, but returns the string directly.
func CompactTextString(pb Message) string {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:930
	_go_fuzz_dep_.CoverTab[112983]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:930
	return compactTextMarshaler.Text(pb)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:930
	// _ = "end of CoverTab[112983]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:930
}

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:930
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text.go:930
var _ = _go_fuzz_dep_.CoverTab
