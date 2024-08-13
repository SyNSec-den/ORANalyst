// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:5
package proto

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:5
)

import (
	"bytes"
	"encoding"
	"fmt"
	"io"
	"math"
	"sort"
	"strings"

	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

const wrapTextMarshalV2 = false

// TextMarshaler is a configurable text format marshaler.
type TextMarshaler struct {
	Compact		bool	// use compact text format (one line)
	ExpandAny	bool	// expand google.protobuf.Any messages of known types
}

// Marshal writes the proto text format of m to w.
func (tm *TextMarshaler) Marshal(w io.Writer, m Message) error {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:32
	_go_fuzz_dep_.CoverTab[62119]++
												b, err := tm.marshal(m)
												if len(b) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:34
		_go_fuzz_dep_.CoverTab[62121]++
													if _, err := w.Write(b); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:35
			_go_fuzz_dep_.CoverTab[62122]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:36
			// _ = "end of CoverTab[62122]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:37
			_go_fuzz_dep_.CoverTab[62123]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:37
			// _ = "end of CoverTab[62123]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:37
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:37
		// _ = "end of CoverTab[62121]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:38
		_go_fuzz_dep_.CoverTab[62124]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:38
		// _ = "end of CoverTab[62124]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:38
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:38
	// _ = "end of CoverTab[62119]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:38
	_go_fuzz_dep_.CoverTab[62120]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:39
	// _ = "end of CoverTab[62120]"
}

// Text returns a proto text formatted string of m.
func (tm *TextMarshaler) Text(m Message) string {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:43
	_go_fuzz_dep_.CoverTab[62125]++
												b, _ := tm.marshal(m)
												return string(b)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:45
	// _ = "end of CoverTab[62125]"
}

func (tm *TextMarshaler) marshal(m Message) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:48
	_go_fuzz_dep_.CoverTab[62126]++
												mr := MessageReflect(m)
												if mr == nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:50
		_go_fuzz_dep_.CoverTab[62128]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:50
		return !mr.IsValid()
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:50
		// _ = "end of CoverTab[62128]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:50
	}() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:50
		_go_fuzz_dep_.CoverTab[62129]++
													return []byte("<nil>"), nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:51
		// _ = "end of CoverTab[62129]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:52
		_go_fuzz_dep_.CoverTab[62130]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:52
		// _ = "end of CoverTab[62130]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:52
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:52
	// _ = "end of CoverTab[62126]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:52
	_go_fuzz_dep_.CoverTab[62127]++

												if wrapTextMarshalV2 {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:54
		_go_fuzz_dep_.CoverTab[62131]++
													if m, ok := m.(encoding.TextMarshaler); ok {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:55
			_go_fuzz_dep_.CoverTab[62135]++
														return m.MarshalText()
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:56
			// _ = "end of CoverTab[62135]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:57
			_go_fuzz_dep_.CoverTab[62136]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:57
			// _ = "end of CoverTab[62136]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:57
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:57
		// _ = "end of CoverTab[62131]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:57
		_go_fuzz_dep_.CoverTab[62132]++

													opts := prototext.MarshalOptions{
			AllowPartial:	true,
			EmitUnknown:	true,
		}
		if !tm.Compact {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:63
			_go_fuzz_dep_.CoverTab[62137]++
														opts.Indent = "  "
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:64
			// _ = "end of CoverTab[62137]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:65
			_go_fuzz_dep_.CoverTab[62138]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:65
			// _ = "end of CoverTab[62138]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:65
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:65
		// _ = "end of CoverTab[62132]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:65
		_go_fuzz_dep_.CoverTab[62133]++
													if !tm.ExpandAny {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:66
			_go_fuzz_dep_.CoverTab[62139]++
														opts.Resolver = (*protoregistry.Types)(nil)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:67
			// _ = "end of CoverTab[62139]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:68
			_go_fuzz_dep_.CoverTab[62140]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:68
			// _ = "end of CoverTab[62140]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:68
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:68
		// _ = "end of CoverTab[62133]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:68
		_go_fuzz_dep_.CoverTab[62134]++
													return opts.Marshal(mr.Interface())
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:69
		// _ = "end of CoverTab[62134]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:70
		_go_fuzz_dep_.CoverTab[62141]++
													w := &textWriter{
			compact:	tm.Compact,
			expandAny:	tm.ExpandAny,
			complete:	true,
		}

		if m, ok := m.(encoding.TextMarshaler); ok {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:77
			_go_fuzz_dep_.CoverTab[62143]++
														b, err := m.MarshalText()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:79
				_go_fuzz_dep_.CoverTab[62145]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:80
				// _ = "end of CoverTab[62145]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:81
				_go_fuzz_dep_.CoverTab[62146]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:81
				// _ = "end of CoverTab[62146]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:81
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:81
			// _ = "end of CoverTab[62143]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:81
			_go_fuzz_dep_.CoverTab[62144]++
														w.Write(b)
														return w.buf, nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:83
			// _ = "end of CoverTab[62144]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:84
			_go_fuzz_dep_.CoverTab[62147]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:84
			// _ = "end of CoverTab[62147]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:84
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:84
		// _ = "end of CoverTab[62141]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:84
		_go_fuzz_dep_.CoverTab[62142]++

													err := w.writeMessage(mr)
													return w.buf, err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:87
		// _ = "end of CoverTab[62142]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:88
	// _ = "end of CoverTab[62127]"
}

var (
	defaultTextMarshaler	= TextMarshaler{}
	compactTextMarshaler	= TextMarshaler{Compact: true}
)

// MarshalText writes the proto text format of m to w.
func MarshalText(w io.Writer, m Message) error {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:97
	_go_fuzz_dep_.CoverTab[62148]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:97
	return defaultTextMarshaler.Marshal(w, m)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:97
	// _ = "end of CoverTab[62148]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:97
}

// MarshalTextString returns a proto text formatted string of m.
func MarshalTextString(m Message) string {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:100
	_go_fuzz_dep_.CoverTab[62149]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:100
	return defaultTextMarshaler.Text(m)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:100
	// _ = "end of CoverTab[62149]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:100
}

// CompactText writes the compact proto text format of m to w.
func CompactText(w io.Writer, m Message) error {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:103
	_go_fuzz_dep_.CoverTab[62150]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:103
	return compactTextMarshaler.Marshal(w, m)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:103
	// _ = "end of CoverTab[62150]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:103
}

// CompactTextString returns a compact proto text formatted string of m.
func CompactTextString(m Message) string {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:106
	_go_fuzz_dep_.CoverTab[62151]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:106
	return compactTextMarshaler.Text(m)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:106
	// _ = "end of CoverTab[62151]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:106
}

var (
	newline		= []byte("\n")
	endBraceNewline	= []byte("}\n")
	posInf		= []byte("inf")
	negInf		= []byte("-inf")
	nan		= []byte("nan")
)

// textWriter is an io.Writer that tracks its indentation level.
type textWriter struct {
	compact		bool	// same as TextMarshaler.Compact
	expandAny	bool	// same as TextMarshaler.ExpandAny
	complete	bool	// whether the current position is a complete line
	indent		int	// indentation level; never negative
	buf		[]byte
}

func (w *textWriter) Write(p []byte) (n int, _ error) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:125
	_go_fuzz_dep_.CoverTab[62152]++
												newlines := bytes.Count(p, newline)
												if newlines == 0 {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:127
		_go_fuzz_dep_.CoverTab[62156]++
													if !w.compact && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:128
			_go_fuzz_dep_.CoverTab[62158]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:128
			return w.complete
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:128
			// _ = "end of CoverTab[62158]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:128
		}() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:128
			_go_fuzz_dep_.CoverTab[62159]++
														w.writeIndent()
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:129
			// _ = "end of CoverTab[62159]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:130
			_go_fuzz_dep_.CoverTab[62160]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:130
			// _ = "end of CoverTab[62160]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:130
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:130
		// _ = "end of CoverTab[62156]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:130
		_go_fuzz_dep_.CoverTab[62157]++
													w.buf = append(w.buf, p...)
													w.complete = false
													return len(p), nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:133
		// _ = "end of CoverTab[62157]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:134
		_go_fuzz_dep_.CoverTab[62161]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:134
		// _ = "end of CoverTab[62161]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:134
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:134
	// _ = "end of CoverTab[62152]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:134
	_go_fuzz_dep_.CoverTab[62153]++

												frags := bytes.SplitN(p, newline, newlines+1)
												if w.compact {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:137
		_go_fuzz_dep_.CoverTab[62162]++
													for i, frag := range frags {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:138
			_go_fuzz_dep_.CoverTab[62164]++
														if i > 0 {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:139
				_go_fuzz_dep_.CoverTab[62166]++
															w.buf = append(w.buf, ' ')
															n++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:141
				// _ = "end of CoverTab[62166]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:142
				_go_fuzz_dep_.CoverTab[62167]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:142
				// _ = "end of CoverTab[62167]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:142
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:142
			// _ = "end of CoverTab[62164]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:142
			_go_fuzz_dep_.CoverTab[62165]++
														w.buf = append(w.buf, frag...)
														n += len(frag)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:144
			// _ = "end of CoverTab[62165]"
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:145
		// _ = "end of CoverTab[62162]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:145
		_go_fuzz_dep_.CoverTab[62163]++
													return n, nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:146
		// _ = "end of CoverTab[62163]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:147
		_go_fuzz_dep_.CoverTab[62168]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:147
		// _ = "end of CoverTab[62168]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:147
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:147
	// _ = "end of CoverTab[62153]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:147
	_go_fuzz_dep_.CoverTab[62154]++

												for i, frag := range frags {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:149
		_go_fuzz_dep_.CoverTab[62169]++
													if w.complete {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:150
			_go_fuzz_dep_.CoverTab[62171]++
														w.writeIndent()
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:151
			// _ = "end of CoverTab[62171]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:152
			_go_fuzz_dep_.CoverTab[62172]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:152
			// _ = "end of CoverTab[62172]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:152
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:152
		// _ = "end of CoverTab[62169]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:152
		_go_fuzz_dep_.CoverTab[62170]++
													w.buf = append(w.buf, frag...)
													n += len(frag)
													if i+1 < len(frags) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:155
			_go_fuzz_dep_.CoverTab[62173]++
														w.buf = append(w.buf, '\n')
														n++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:157
			// _ = "end of CoverTab[62173]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:158
			_go_fuzz_dep_.CoverTab[62174]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:158
			// _ = "end of CoverTab[62174]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:158
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:158
		// _ = "end of CoverTab[62170]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:159
	// _ = "end of CoverTab[62154]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:159
	_go_fuzz_dep_.CoverTab[62155]++
												w.complete = len(frags[len(frags)-1]) == 0
												return n, nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:161
	// _ = "end of CoverTab[62155]"
}

func (w *textWriter) WriteByte(c byte) error {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:164
	_go_fuzz_dep_.CoverTab[62175]++
												if w.compact && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:165
		_go_fuzz_dep_.CoverTab[62178]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:165
		return c == '\n'
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:165
		// _ = "end of CoverTab[62178]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:165
	}() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:165
		_go_fuzz_dep_.CoverTab[62179]++
													c = ' '
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:166
		// _ = "end of CoverTab[62179]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:167
		_go_fuzz_dep_.CoverTab[62180]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:167
		// _ = "end of CoverTab[62180]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:167
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:167
	// _ = "end of CoverTab[62175]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:167
	_go_fuzz_dep_.CoverTab[62176]++
												if !w.compact && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:168
		_go_fuzz_dep_.CoverTab[62181]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:168
		return w.complete
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:168
		// _ = "end of CoverTab[62181]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:168
	}() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:168
		_go_fuzz_dep_.CoverTab[62182]++
													w.writeIndent()
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:169
		// _ = "end of CoverTab[62182]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:170
		_go_fuzz_dep_.CoverTab[62183]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:170
		// _ = "end of CoverTab[62183]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:170
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:170
	// _ = "end of CoverTab[62176]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:170
	_go_fuzz_dep_.CoverTab[62177]++
												w.buf = append(w.buf, c)
												w.complete = c == '\n'
												return nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:173
	// _ = "end of CoverTab[62177]"
}

func (w *textWriter) writeName(fd protoreflect.FieldDescriptor) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:176
	_go_fuzz_dep_.CoverTab[62184]++
												if !w.compact && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:177
		_go_fuzz_dep_.CoverTab[62187]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:177
		return w.complete
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:177
		// _ = "end of CoverTab[62187]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:177
	}() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:177
		_go_fuzz_dep_.CoverTab[62188]++
													w.writeIndent()
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:178
		// _ = "end of CoverTab[62188]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:179
		_go_fuzz_dep_.CoverTab[62189]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:179
		// _ = "end of CoverTab[62189]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:179
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:179
	// _ = "end of CoverTab[62184]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:179
	_go_fuzz_dep_.CoverTab[62185]++
												w.complete = false

												if fd.Kind() != protoreflect.GroupKind {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:182
		_go_fuzz_dep_.CoverTab[62190]++
													w.buf = append(w.buf, fd.Name()...)
													w.WriteByte(':')
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:184
		// _ = "end of CoverTab[62190]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:185
		_go_fuzz_dep_.CoverTab[62191]++

													w.buf = append(w.buf, fd.Message().Name()...)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:187
		// _ = "end of CoverTab[62191]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:188
	// _ = "end of CoverTab[62185]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:188
	_go_fuzz_dep_.CoverTab[62186]++

												if !w.compact {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:190
		_go_fuzz_dep_.CoverTab[62192]++
													w.WriteByte(' ')
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:191
		// _ = "end of CoverTab[62192]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:192
		_go_fuzz_dep_.CoverTab[62193]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:192
		// _ = "end of CoverTab[62193]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:192
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:192
	// _ = "end of CoverTab[62186]"
}

func requiresQuotes(u string) bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:195
	_go_fuzz_dep_.CoverTab[62194]++

												for _, ch := range u {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:197
		_go_fuzz_dep_.CoverTab[62196]++
													switch {
		case ch == '.' || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:199
			_go_fuzz_dep_.CoverTab[62202]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:199
			return ch == '/'
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:199
			// _ = "end of CoverTab[62202]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:199
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:199
			_go_fuzz_dep_.CoverTab[62203]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:199
			return ch == '_'
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:199
			// _ = "end of CoverTab[62203]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:199
		}():
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:199
			_go_fuzz_dep_.CoverTab[62197]++
														continue
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:200
			// _ = "end of CoverTab[62197]"
		case '0' <= ch && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:201
			_go_fuzz_dep_.CoverTab[62204]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:201
			return ch <= '9'
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:201
			// _ = "end of CoverTab[62204]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:201
		}():
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:201
			_go_fuzz_dep_.CoverTab[62198]++
														continue
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:202
			// _ = "end of CoverTab[62198]"
		case 'A' <= ch && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:203
			_go_fuzz_dep_.CoverTab[62205]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:203
			return ch <= 'Z'
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:203
			// _ = "end of CoverTab[62205]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:203
		}():
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:203
			_go_fuzz_dep_.CoverTab[62199]++
														continue
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:204
			// _ = "end of CoverTab[62199]"
		case 'a' <= ch && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:205
			_go_fuzz_dep_.CoverTab[62206]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:205
			return ch <= 'z'
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:205
			// _ = "end of CoverTab[62206]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:205
		}():
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:205
			_go_fuzz_dep_.CoverTab[62200]++
														continue
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:206
			// _ = "end of CoverTab[62200]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:207
			_go_fuzz_dep_.CoverTab[62201]++
														return true
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:208
			// _ = "end of CoverTab[62201]"
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:209
		// _ = "end of CoverTab[62196]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:210
	// _ = "end of CoverTab[62194]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:210
	_go_fuzz_dep_.CoverTab[62195]++
												return false
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:211
	// _ = "end of CoverTab[62195]"
}

// writeProto3Any writes an expanded google.protobuf.Any message.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:214
//
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:214
// It returns (false, nil) if sv value can't be unmarshaled (e.g. because
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:214
// required messages are not linked in).
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:214
//
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:214
// It returns (true, error) when sv was written in expanded format or an error
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:214
// was encountered.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:221
func (w *textWriter) writeProto3Any(m protoreflect.Message) (bool, error) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:221
	_go_fuzz_dep_.CoverTab[62207]++
												md := m.Descriptor()
												fdURL := md.Fields().ByName("type_url")
												fdVal := md.Fields().ByName("value")

												url := m.Get(fdURL).String()
												mt, err := protoregistry.GlobalTypes.FindMessageByURL(url)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:228
		_go_fuzz_dep_.CoverTab[62214]++
													return false, nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:229
		// _ = "end of CoverTab[62214]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:230
		_go_fuzz_dep_.CoverTab[62215]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:230
		// _ = "end of CoverTab[62215]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:230
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:230
	// _ = "end of CoverTab[62207]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:230
	_go_fuzz_dep_.CoverTab[62208]++

												b := m.Get(fdVal).Bytes()
												m2 := mt.New()
												if err := proto.Unmarshal(b, m2.Interface()); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:234
		_go_fuzz_dep_.CoverTab[62216]++
													return false, nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:235
		// _ = "end of CoverTab[62216]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:236
		_go_fuzz_dep_.CoverTab[62217]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:236
		// _ = "end of CoverTab[62217]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:236
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:236
	// _ = "end of CoverTab[62208]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:236
	_go_fuzz_dep_.CoverTab[62209]++
												w.Write([]byte("["))
												if requiresQuotes(url) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:238
		_go_fuzz_dep_.CoverTab[62218]++
													w.writeQuotedString(url)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:239
		// _ = "end of CoverTab[62218]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:240
		_go_fuzz_dep_.CoverTab[62219]++
													w.Write([]byte(url))
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:241
		// _ = "end of CoverTab[62219]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:242
	// _ = "end of CoverTab[62209]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:242
	_go_fuzz_dep_.CoverTab[62210]++
												if w.compact {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:243
		_go_fuzz_dep_.CoverTab[62220]++
													w.Write([]byte("]:<"))
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:244
		// _ = "end of CoverTab[62220]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:245
		_go_fuzz_dep_.CoverTab[62221]++
													w.Write([]byte("]: <\n"))
													w.indent++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:247
		// _ = "end of CoverTab[62221]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:248
	// _ = "end of CoverTab[62210]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:248
	_go_fuzz_dep_.CoverTab[62211]++
												if err := w.writeMessage(m2); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:249
		_go_fuzz_dep_.CoverTab[62222]++
													return true, err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:250
		// _ = "end of CoverTab[62222]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:251
		_go_fuzz_dep_.CoverTab[62223]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:251
		// _ = "end of CoverTab[62223]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:251
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:251
	// _ = "end of CoverTab[62211]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:251
	_go_fuzz_dep_.CoverTab[62212]++
												if w.compact {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:252
		_go_fuzz_dep_.CoverTab[62224]++
													w.Write([]byte("> "))
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:253
		// _ = "end of CoverTab[62224]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:254
		_go_fuzz_dep_.CoverTab[62225]++
													w.indent--
													w.Write([]byte(">\n"))
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:256
		// _ = "end of CoverTab[62225]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:257
	// _ = "end of CoverTab[62212]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:257
	_go_fuzz_dep_.CoverTab[62213]++
												return true, nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:258
	// _ = "end of CoverTab[62213]"
}

func (w *textWriter) writeMessage(m protoreflect.Message) error {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:261
	_go_fuzz_dep_.CoverTab[62226]++
												md := m.Descriptor()
												if w.expandAny && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:263
		_go_fuzz_dep_.CoverTab[62230]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:263
		return md.FullName() == "google.protobuf.Any"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:263
		// _ = "end of CoverTab[62230]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:263
	}() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:263
		_go_fuzz_dep_.CoverTab[62231]++
													if canExpand, err := w.writeProto3Any(m); canExpand {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:264
			_go_fuzz_dep_.CoverTab[62232]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:265
			// _ = "end of CoverTab[62232]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:266
			_go_fuzz_dep_.CoverTab[62233]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:266
			// _ = "end of CoverTab[62233]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:266
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:266
		// _ = "end of CoverTab[62231]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:267
		_go_fuzz_dep_.CoverTab[62234]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:267
		// _ = "end of CoverTab[62234]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:267
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:267
	// _ = "end of CoverTab[62226]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:267
	_go_fuzz_dep_.CoverTab[62227]++

												fds := md.Fields()
												for i := 0; i < fds.Len(); {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:270
		_go_fuzz_dep_.CoverTab[62235]++
													fd := fds.Get(i)
													if od := fd.ContainingOneof(); od != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:272
			_go_fuzz_dep_.CoverTab[62238]++
														fd = m.WhichOneof(od)
														i += od.Fields().Len()
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:274
			// _ = "end of CoverTab[62238]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:275
			_go_fuzz_dep_.CoverTab[62239]++
														i++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:276
			// _ = "end of CoverTab[62239]"
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:277
		// _ = "end of CoverTab[62235]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:277
		_go_fuzz_dep_.CoverTab[62236]++
													if fd == nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:278
			_go_fuzz_dep_.CoverTab[62240]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:278
			return !m.Has(fd)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:278
			// _ = "end of CoverTab[62240]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:278
		}() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:278
			_go_fuzz_dep_.CoverTab[62241]++
														continue
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:279
			// _ = "end of CoverTab[62241]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:280
			_go_fuzz_dep_.CoverTab[62242]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:280
			// _ = "end of CoverTab[62242]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:280
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:280
		// _ = "end of CoverTab[62236]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:280
		_go_fuzz_dep_.CoverTab[62237]++

													switch {
		case fd.IsList():
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:283
			_go_fuzz_dep_.CoverTab[62243]++
														lv := m.Get(fd).List()
														for j := 0; j < lv.Len(); j++ {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:285
				_go_fuzz_dep_.CoverTab[62249]++
															w.writeName(fd)
															v := lv.Get(j)
															if err := w.writeSingularValue(v, fd); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:288
					_go_fuzz_dep_.CoverTab[62251]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:289
					// _ = "end of CoverTab[62251]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:290
					_go_fuzz_dep_.CoverTab[62252]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:290
					// _ = "end of CoverTab[62252]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:290
				}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:290
				// _ = "end of CoverTab[62249]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:290
				_go_fuzz_dep_.CoverTab[62250]++
															w.WriteByte('\n')
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:291
				// _ = "end of CoverTab[62250]"
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:292
			// _ = "end of CoverTab[62243]"
		case fd.IsMap():
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:293
			_go_fuzz_dep_.CoverTab[62244]++
														kfd := fd.MapKey()
														vfd := fd.MapValue()
														mv := m.Get(fd).Map()

														type entry struct{ key, val protoreflect.Value }
														var entries []entry
														mv.Range(func(k protoreflect.MapKey, v protoreflect.Value) bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:300
				_go_fuzz_dep_.CoverTab[62253]++
															entries = append(entries, entry{k.Value(), v})
															return true
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:302
				// _ = "end of CoverTab[62253]"
			})
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:303
			// _ = "end of CoverTab[62244]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:303
			_go_fuzz_dep_.CoverTab[62245]++
														sort.Slice(entries, func(i, j int) bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:304
				_go_fuzz_dep_.CoverTab[62254]++
															switch kfd.Kind() {
				case protoreflect.BoolKind:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:306
					_go_fuzz_dep_.CoverTab[62255]++
																return !entries[i].key.Bool() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:307
						_go_fuzz_dep_.CoverTab[62260]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:307
						return entries[j].key.Bool()
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:307
						// _ = "end of CoverTab[62260]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:307
					}()
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:307
					// _ = "end of CoverTab[62255]"
				case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind, protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:308
					_go_fuzz_dep_.CoverTab[62256]++
																return entries[i].key.Int() < entries[j].key.Int()
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:309
					// _ = "end of CoverTab[62256]"
				case protoreflect.Uint32Kind, protoreflect.Fixed32Kind, protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:310
					_go_fuzz_dep_.CoverTab[62257]++
																return entries[i].key.Uint() < entries[j].key.Uint()
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:311
					// _ = "end of CoverTab[62257]"
				case protoreflect.StringKind:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:312
					_go_fuzz_dep_.CoverTab[62258]++
																return entries[i].key.String() < entries[j].key.String()
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:313
					// _ = "end of CoverTab[62258]"
				default:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:314
					_go_fuzz_dep_.CoverTab[62259]++
																panic("invalid kind")
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:315
					// _ = "end of CoverTab[62259]"
				}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:316
				// _ = "end of CoverTab[62254]"
			})
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:317
			// _ = "end of CoverTab[62245]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:317
			_go_fuzz_dep_.CoverTab[62246]++
														for _, entry := range entries {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:318
				_go_fuzz_dep_.CoverTab[62261]++
															w.writeName(fd)
															w.WriteByte('<')
															if !w.compact {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:321
					_go_fuzz_dep_.CoverTab[62265]++
																w.WriteByte('\n')
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:322
					// _ = "end of CoverTab[62265]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:323
					_go_fuzz_dep_.CoverTab[62266]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:323
					// _ = "end of CoverTab[62266]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:323
				}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:323
				// _ = "end of CoverTab[62261]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:323
				_go_fuzz_dep_.CoverTab[62262]++
															w.indent++
															w.writeName(kfd)
															if err := w.writeSingularValue(entry.key, kfd); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:326
					_go_fuzz_dep_.CoverTab[62267]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:327
					// _ = "end of CoverTab[62267]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:328
					_go_fuzz_dep_.CoverTab[62268]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:328
					// _ = "end of CoverTab[62268]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:328
				}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:328
				// _ = "end of CoverTab[62262]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:328
				_go_fuzz_dep_.CoverTab[62263]++
															w.WriteByte('\n')
															w.writeName(vfd)
															if err := w.writeSingularValue(entry.val, vfd); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:331
					_go_fuzz_dep_.CoverTab[62269]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:332
					// _ = "end of CoverTab[62269]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:333
					_go_fuzz_dep_.CoverTab[62270]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:333
					// _ = "end of CoverTab[62270]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:333
				}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:333
				// _ = "end of CoverTab[62263]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:333
				_go_fuzz_dep_.CoverTab[62264]++
															w.WriteByte('\n')
															w.indent--
															w.WriteByte('>')
															w.WriteByte('\n')
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:337
				// _ = "end of CoverTab[62264]"
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:338
			// _ = "end of CoverTab[62246]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:339
			_go_fuzz_dep_.CoverTab[62247]++
														w.writeName(fd)
														if err := w.writeSingularValue(m.Get(fd), fd); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:341
				_go_fuzz_dep_.CoverTab[62271]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:342
				// _ = "end of CoverTab[62271]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:343
				_go_fuzz_dep_.CoverTab[62272]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:343
				// _ = "end of CoverTab[62272]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:343
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:343
			// _ = "end of CoverTab[62247]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:343
			_go_fuzz_dep_.CoverTab[62248]++
														w.WriteByte('\n')
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:344
			// _ = "end of CoverTab[62248]"
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:345
		// _ = "end of CoverTab[62237]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:346
	// _ = "end of CoverTab[62227]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:346
	_go_fuzz_dep_.CoverTab[62228]++

												if b := m.GetUnknown(); len(b) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:348
		_go_fuzz_dep_.CoverTab[62273]++
													w.writeUnknownFields(b)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:349
		// _ = "end of CoverTab[62273]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:350
		_go_fuzz_dep_.CoverTab[62274]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:350
		// _ = "end of CoverTab[62274]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:350
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:350
	// _ = "end of CoverTab[62228]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:350
	_go_fuzz_dep_.CoverTab[62229]++
												return w.writeExtensions(m)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:351
	// _ = "end of CoverTab[62229]"
}

func (w *textWriter) writeSingularValue(v protoreflect.Value, fd protoreflect.FieldDescriptor) error {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:354
	_go_fuzz_dep_.CoverTab[62275]++
												switch fd.Kind() {
	case protoreflect.FloatKind, protoreflect.DoubleKind:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:356
		_go_fuzz_dep_.CoverTab[62277]++
													switch vf := v.Float(); {
		case math.IsInf(vf, +1):
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:358
			_go_fuzz_dep_.CoverTab[62286]++
														w.Write(posInf)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:359
			// _ = "end of CoverTab[62286]"
		case math.IsInf(vf, -1):
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:360
			_go_fuzz_dep_.CoverTab[62287]++
														w.Write(negInf)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:361
			// _ = "end of CoverTab[62287]"
		case math.IsNaN(vf):
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:362
			_go_fuzz_dep_.CoverTab[62288]++
														w.Write(nan)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:363
			// _ = "end of CoverTab[62288]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:364
			_go_fuzz_dep_.CoverTab[62289]++
														fmt.Fprint(w, v.Interface())
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:365
			// _ = "end of CoverTab[62289]"
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:366
		// _ = "end of CoverTab[62277]"
	case protoreflect.StringKind:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:367
		_go_fuzz_dep_.CoverTab[62278]++

													w.writeQuotedString(string(v.String()))
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:369
		// _ = "end of CoverTab[62278]"
	case protoreflect.BytesKind:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:370
		_go_fuzz_dep_.CoverTab[62279]++
													w.writeQuotedString(string(v.Bytes()))
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:371
		// _ = "end of CoverTab[62279]"
	case protoreflect.MessageKind, protoreflect.GroupKind:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:372
		_go_fuzz_dep_.CoverTab[62280]++
													var bra, ket byte = '<', '>'
													if fd.Kind() == protoreflect.GroupKind {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:374
			_go_fuzz_dep_.CoverTab[62290]++
														bra, ket = '{', '}'
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:375
			// _ = "end of CoverTab[62290]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:376
			_go_fuzz_dep_.CoverTab[62291]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:376
			// _ = "end of CoverTab[62291]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:376
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:376
		// _ = "end of CoverTab[62280]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:376
		_go_fuzz_dep_.CoverTab[62281]++
													w.WriteByte(bra)
													if !w.compact {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:378
			_go_fuzz_dep_.CoverTab[62292]++
														w.WriteByte('\n')
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:379
			// _ = "end of CoverTab[62292]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:380
			_go_fuzz_dep_.CoverTab[62293]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:380
			// _ = "end of CoverTab[62293]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:380
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:380
		// _ = "end of CoverTab[62281]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:380
		_go_fuzz_dep_.CoverTab[62282]++
													w.indent++
													m := v.Message()
													if m2, ok := m.Interface().(encoding.TextMarshaler); ok {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:383
			_go_fuzz_dep_.CoverTab[62294]++
														b, err := m2.MarshalText()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:385
				_go_fuzz_dep_.CoverTab[62296]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:386
				// _ = "end of CoverTab[62296]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:387
				_go_fuzz_dep_.CoverTab[62297]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:387
				// _ = "end of CoverTab[62297]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:387
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:387
			// _ = "end of CoverTab[62294]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:387
			_go_fuzz_dep_.CoverTab[62295]++
														w.Write(b)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:388
			// _ = "end of CoverTab[62295]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:389
			_go_fuzz_dep_.CoverTab[62298]++
														w.writeMessage(m)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:390
			// _ = "end of CoverTab[62298]"
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:391
		// _ = "end of CoverTab[62282]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:391
		_go_fuzz_dep_.CoverTab[62283]++
													w.indent--
													w.WriteByte(ket)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:393
		// _ = "end of CoverTab[62283]"
	case protoreflect.EnumKind:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:394
		_go_fuzz_dep_.CoverTab[62284]++
													if ev := fd.Enum().Values().ByNumber(v.Enum()); ev != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:395
			_go_fuzz_dep_.CoverTab[62299]++
														fmt.Fprint(w, ev.Name())
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:396
			// _ = "end of CoverTab[62299]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:397
			_go_fuzz_dep_.CoverTab[62300]++
														fmt.Fprint(w, v.Enum())
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:398
			// _ = "end of CoverTab[62300]"
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:399
		// _ = "end of CoverTab[62284]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:400
		_go_fuzz_dep_.CoverTab[62285]++
													fmt.Fprint(w, v.Interface())
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:401
		// _ = "end of CoverTab[62285]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:402
	// _ = "end of CoverTab[62275]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:402
	_go_fuzz_dep_.CoverTab[62276]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:403
	// _ = "end of CoverTab[62276]"
}

// writeQuotedString writes a quoted string in the protocol buffer text format.
func (w *textWriter) writeQuotedString(s string) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:407
	_go_fuzz_dep_.CoverTab[62301]++
												w.WriteByte('"')
												for i := 0; i < len(s); i++ {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:409
		_go_fuzz_dep_.CoverTab[62303]++
													switch c := s[i]; c {
		case '\n':
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:411
			_go_fuzz_dep_.CoverTab[62304]++
														w.buf = append(w.buf, `\n`...)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:412
			// _ = "end of CoverTab[62304]"
		case '\r':
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:413
			_go_fuzz_dep_.CoverTab[62305]++
														w.buf = append(w.buf, `\r`...)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:414
			// _ = "end of CoverTab[62305]"
		case '\t':
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:415
			_go_fuzz_dep_.CoverTab[62306]++
														w.buf = append(w.buf, `\t`...)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:416
			// _ = "end of CoverTab[62306]"
		case '"':
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:417
			_go_fuzz_dep_.CoverTab[62307]++
														w.buf = append(w.buf, `\"`...)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:418
			// _ = "end of CoverTab[62307]"
		case '\\':
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:419
			_go_fuzz_dep_.CoverTab[62308]++
														w.buf = append(w.buf, `\\`...)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:420
			// _ = "end of CoverTab[62308]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:421
			_go_fuzz_dep_.CoverTab[62309]++
														if isPrint := c >= 0x20 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:422
				_go_fuzz_dep_.CoverTab[62310]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:422
				return c < 0x7f
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:422
				// _ = "end of CoverTab[62310]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:422
			}(); isPrint {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:422
				_go_fuzz_dep_.CoverTab[62311]++
															w.buf = append(w.buf, c)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:423
				// _ = "end of CoverTab[62311]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:424
				_go_fuzz_dep_.CoverTab[62312]++
															w.buf = append(w.buf, fmt.Sprintf(`\%03o`, c)...)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:425
				// _ = "end of CoverTab[62312]"
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:426
			// _ = "end of CoverTab[62309]"
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:427
		// _ = "end of CoverTab[62303]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:428
	// _ = "end of CoverTab[62301]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:428
	_go_fuzz_dep_.CoverTab[62302]++
												w.WriteByte('"')
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:429
	// _ = "end of CoverTab[62302]"
}

func (w *textWriter) writeUnknownFields(b []byte) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:432
	_go_fuzz_dep_.CoverTab[62313]++
												if !w.compact {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:433
		_go_fuzz_dep_.CoverTab[62315]++
													fmt.Fprintf(w, "/* %d unknown bytes */\n", len(b))
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:434
		// _ = "end of CoverTab[62315]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:435
		_go_fuzz_dep_.CoverTab[62316]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:435
		// _ = "end of CoverTab[62316]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:435
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:435
	// _ = "end of CoverTab[62313]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:435
	_go_fuzz_dep_.CoverTab[62314]++

												for len(b) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:437
		_go_fuzz_dep_.CoverTab[62317]++
													num, wtyp, n := protowire.ConsumeTag(b)
													if n < 0 {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:439
			_go_fuzz_dep_.CoverTab[62323]++
														return
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:440
			// _ = "end of CoverTab[62323]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:441
			_go_fuzz_dep_.CoverTab[62324]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:441
			// _ = "end of CoverTab[62324]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:441
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:441
		// _ = "end of CoverTab[62317]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:441
		_go_fuzz_dep_.CoverTab[62318]++
													b = b[n:]

													if wtyp == protowire.EndGroupType {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:444
			_go_fuzz_dep_.CoverTab[62325]++
														w.indent--
														w.Write(endBraceNewline)
														continue
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:447
			// _ = "end of CoverTab[62325]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:448
			_go_fuzz_dep_.CoverTab[62326]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:448
			// _ = "end of CoverTab[62326]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:448
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:448
		// _ = "end of CoverTab[62318]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:448
		_go_fuzz_dep_.CoverTab[62319]++
													fmt.Fprint(w, num)
													if wtyp != protowire.StartGroupType {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:450
			_go_fuzz_dep_.CoverTab[62327]++
														w.WriteByte(':')
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:451
			// _ = "end of CoverTab[62327]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:452
			_go_fuzz_dep_.CoverTab[62328]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:452
			// _ = "end of CoverTab[62328]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:452
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:452
		// _ = "end of CoverTab[62319]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:452
		_go_fuzz_dep_.CoverTab[62320]++
													if !w.compact || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:453
			_go_fuzz_dep_.CoverTab[62329]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:453
			return wtyp == protowire.StartGroupType
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:453
			// _ = "end of CoverTab[62329]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:453
		}() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:453
			_go_fuzz_dep_.CoverTab[62330]++
														w.WriteByte(' ')
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:454
			// _ = "end of CoverTab[62330]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:455
			_go_fuzz_dep_.CoverTab[62331]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:455
			// _ = "end of CoverTab[62331]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:455
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:455
		// _ = "end of CoverTab[62320]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:455
		_go_fuzz_dep_.CoverTab[62321]++
													switch wtyp {
		case protowire.VarintType:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:457
			_go_fuzz_dep_.CoverTab[62332]++
														v, n := protowire.ConsumeVarint(b)
														if n < 0 {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:459
				_go_fuzz_dep_.CoverTab[62342]++
															return
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:460
				// _ = "end of CoverTab[62342]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:461
				_go_fuzz_dep_.CoverTab[62343]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:461
				// _ = "end of CoverTab[62343]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:461
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:461
			// _ = "end of CoverTab[62332]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:461
			_go_fuzz_dep_.CoverTab[62333]++
														b = b[n:]
														fmt.Fprint(w, v)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:463
			// _ = "end of CoverTab[62333]"
		case protowire.Fixed32Type:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:464
			_go_fuzz_dep_.CoverTab[62334]++
														v, n := protowire.ConsumeFixed32(b)
														if n < 0 {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:466
				_go_fuzz_dep_.CoverTab[62344]++
															return
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:467
				// _ = "end of CoverTab[62344]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:468
				_go_fuzz_dep_.CoverTab[62345]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:468
				// _ = "end of CoverTab[62345]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:468
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:468
			// _ = "end of CoverTab[62334]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:468
			_go_fuzz_dep_.CoverTab[62335]++
														b = b[n:]
														fmt.Fprint(w, v)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:470
			// _ = "end of CoverTab[62335]"
		case protowire.Fixed64Type:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:471
			_go_fuzz_dep_.CoverTab[62336]++
														v, n := protowire.ConsumeFixed64(b)
														if n < 0 {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:473
				_go_fuzz_dep_.CoverTab[62346]++
															return
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:474
				// _ = "end of CoverTab[62346]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:475
				_go_fuzz_dep_.CoverTab[62347]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:475
				// _ = "end of CoverTab[62347]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:475
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:475
			// _ = "end of CoverTab[62336]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:475
			_go_fuzz_dep_.CoverTab[62337]++
														b = b[n:]
														fmt.Fprint(w, v)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:477
			// _ = "end of CoverTab[62337]"
		case protowire.BytesType:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:478
			_go_fuzz_dep_.CoverTab[62338]++
														v, n := protowire.ConsumeBytes(b)
														if n < 0 {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:480
				_go_fuzz_dep_.CoverTab[62348]++
															return
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:481
				// _ = "end of CoverTab[62348]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:482
				_go_fuzz_dep_.CoverTab[62349]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:482
				// _ = "end of CoverTab[62349]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:482
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:482
			// _ = "end of CoverTab[62338]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:482
			_go_fuzz_dep_.CoverTab[62339]++
														b = b[n:]
														fmt.Fprintf(w, "%q", v)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:484
			// _ = "end of CoverTab[62339]"
		case protowire.StartGroupType:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:485
			_go_fuzz_dep_.CoverTab[62340]++
														w.WriteByte('{')
														w.indent++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:487
			// _ = "end of CoverTab[62340]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:488
			_go_fuzz_dep_.CoverTab[62341]++
														fmt.Fprintf(w, "/* unknown wire type %d */", wtyp)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:489
			// _ = "end of CoverTab[62341]"
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:490
		// _ = "end of CoverTab[62321]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:490
		_go_fuzz_dep_.CoverTab[62322]++
													w.WriteByte('\n')
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:491
		// _ = "end of CoverTab[62322]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:492
	// _ = "end of CoverTab[62314]"
}

// writeExtensions writes all the extensions in m.
func (w *textWriter) writeExtensions(m protoreflect.Message) error {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:496
	_go_fuzz_dep_.CoverTab[62350]++
												md := m.Descriptor()
												if md.ExtensionRanges().Len() == 0 {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:498
		_go_fuzz_dep_.CoverTab[62355]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:499
		// _ = "end of CoverTab[62355]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:500
		_go_fuzz_dep_.CoverTab[62356]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:500
		// _ = "end of CoverTab[62356]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:500
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:500
	// _ = "end of CoverTab[62350]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:500
	_go_fuzz_dep_.CoverTab[62351]++

												type ext struct {
		desc	protoreflect.FieldDescriptor
		val	protoreflect.Value
	}
	var exts []ext
	m.Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:507
		_go_fuzz_dep_.CoverTab[62357]++
													if fd.IsExtension() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:508
			_go_fuzz_dep_.CoverTab[62359]++
														exts = append(exts, ext{fd, v})
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:509
			// _ = "end of CoverTab[62359]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:510
			_go_fuzz_dep_.CoverTab[62360]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:510
			// _ = "end of CoverTab[62360]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:510
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:510
		// _ = "end of CoverTab[62357]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:510
		_go_fuzz_dep_.CoverTab[62358]++
													return true
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:511
		// _ = "end of CoverTab[62358]"
	})
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:512
	// _ = "end of CoverTab[62351]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:512
	_go_fuzz_dep_.CoverTab[62352]++
												sort.Slice(exts, func(i, j int) bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:513
		_go_fuzz_dep_.CoverTab[62361]++
													return exts[i].desc.Number() < exts[j].desc.Number()
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:514
		// _ = "end of CoverTab[62361]"
	})
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:515
	// _ = "end of CoverTab[62352]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:515
	_go_fuzz_dep_.CoverTab[62353]++

												for _, ext := range exts {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:517
		_go_fuzz_dep_.CoverTab[62362]++

													name := string(ext.desc.FullName())
													if isMessageSet(ext.desc.ContainingMessage()) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:520
			_go_fuzz_dep_.CoverTab[62364]++
														name = strings.TrimSuffix(name, ".message_set_extension")
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:521
			// _ = "end of CoverTab[62364]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:522
			_go_fuzz_dep_.CoverTab[62365]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:522
			// _ = "end of CoverTab[62365]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:522
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:522
		// _ = "end of CoverTab[62362]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:522
		_go_fuzz_dep_.CoverTab[62363]++

													if !ext.desc.IsList() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:524
			_go_fuzz_dep_.CoverTab[62366]++
														if err := w.writeSingularExtension(name, ext.val, ext.desc); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:525
				_go_fuzz_dep_.CoverTab[62367]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:526
				// _ = "end of CoverTab[62367]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:527
				_go_fuzz_dep_.CoverTab[62368]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:527
				// _ = "end of CoverTab[62368]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:527
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:527
			// _ = "end of CoverTab[62366]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:528
			_go_fuzz_dep_.CoverTab[62369]++
														lv := ext.val.List()
														for i := 0; i < lv.Len(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:530
				_go_fuzz_dep_.CoverTab[62370]++
															if err := w.writeSingularExtension(name, lv.Get(i), ext.desc); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:531
					_go_fuzz_dep_.CoverTab[62371]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:532
					// _ = "end of CoverTab[62371]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:533
					_go_fuzz_dep_.CoverTab[62372]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:533
					// _ = "end of CoverTab[62372]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:533
				}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:533
				// _ = "end of CoverTab[62370]"
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:534
			// _ = "end of CoverTab[62369]"
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:535
		// _ = "end of CoverTab[62363]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:536
	// _ = "end of CoverTab[62353]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:536
	_go_fuzz_dep_.CoverTab[62354]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:537
	// _ = "end of CoverTab[62354]"
}

func (w *textWriter) writeSingularExtension(name string, v protoreflect.Value, fd protoreflect.FieldDescriptor) error {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:540
	_go_fuzz_dep_.CoverTab[62373]++
												fmt.Fprintf(w, "[%s]:", name)
												if !w.compact {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:542
		_go_fuzz_dep_.CoverTab[62376]++
													w.WriteByte(' ')
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:543
		// _ = "end of CoverTab[62376]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:544
		_go_fuzz_dep_.CoverTab[62377]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:544
		// _ = "end of CoverTab[62377]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:544
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:544
	// _ = "end of CoverTab[62373]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:544
	_go_fuzz_dep_.CoverTab[62374]++
												if err := w.writeSingularValue(v, fd); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:545
		_go_fuzz_dep_.CoverTab[62378]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:546
		// _ = "end of CoverTab[62378]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:547
		_go_fuzz_dep_.CoverTab[62379]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:547
		// _ = "end of CoverTab[62379]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:547
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:547
	// _ = "end of CoverTab[62374]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:547
	_go_fuzz_dep_.CoverTab[62375]++
												w.WriteByte('\n')
												return nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:549
	// _ = "end of CoverTab[62375]"
}

func (w *textWriter) writeIndent() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:552
	_go_fuzz_dep_.CoverTab[62380]++
												if !w.complete {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:553
		_go_fuzz_dep_.CoverTab[62383]++
													return
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:554
		// _ = "end of CoverTab[62383]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:555
		_go_fuzz_dep_.CoverTab[62384]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:555
		// _ = "end of CoverTab[62384]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:555
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:555
	// _ = "end of CoverTab[62380]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:555
	_go_fuzz_dep_.CoverTab[62381]++
												for i := 0; i < w.indent*2; i++ {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:556
		_go_fuzz_dep_.CoverTab[62385]++
													w.buf = append(w.buf, ' ')
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:557
		// _ = "end of CoverTab[62385]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:558
	// _ = "end of CoverTab[62381]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:558
	_go_fuzz_dep_.CoverTab[62382]++
												w.complete = false
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:559
	// _ = "end of CoverTab[62382]"
}

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:560
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_encode.go:560
var _ = _go_fuzz_dep_.CoverTab
