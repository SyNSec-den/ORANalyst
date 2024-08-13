// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:5
package proto

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:5
)

import (
	"encoding"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unicode/utf8"

	"google.golang.org/protobuf/encoding/prototext"
	protoV2 "google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

const wrapTextUnmarshalV2 = false

// ParseError is returned by UnmarshalText.
type ParseError struct {
	Message	string

	// Deprecated: Do not use.
	Line, Offset	int
}

func (e *ParseError) Error() string {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:32
	_go_fuzz_dep_.CoverTab[61660]++
												if wrapTextUnmarshalV2 {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:33
		_go_fuzz_dep_.CoverTab[61663]++
													return e.Message
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:34
		// _ = "end of CoverTab[61663]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:35
		_go_fuzz_dep_.CoverTab[61664]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:35
		// _ = "end of CoverTab[61664]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:35
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:35
	// _ = "end of CoverTab[61660]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:35
	_go_fuzz_dep_.CoverTab[61661]++
												if e.Line == 1 {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:36
		_go_fuzz_dep_.CoverTab[61665]++
													return fmt.Sprintf("line 1.%d: %v", e.Offset, e.Message)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:37
		// _ = "end of CoverTab[61665]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:38
		_go_fuzz_dep_.CoverTab[61666]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:38
		// _ = "end of CoverTab[61666]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:38
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:38
	// _ = "end of CoverTab[61661]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:38
	_go_fuzz_dep_.CoverTab[61662]++
												return fmt.Sprintf("line %d: %v", e.Line, e.Message)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:39
	// _ = "end of CoverTab[61662]"
}

// UnmarshalText parses a proto text formatted string into m.
func UnmarshalText(s string, m Message) error {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:43
	_go_fuzz_dep_.CoverTab[61667]++
												if u, ok := m.(encoding.TextUnmarshaler); ok {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:44
		_go_fuzz_dep_.CoverTab[61669]++
													return u.UnmarshalText([]byte(s))
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:45
		// _ = "end of CoverTab[61669]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:46
		_go_fuzz_dep_.CoverTab[61670]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:46
		// _ = "end of CoverTab[61670]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:46
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:46
	// _ = "end of CoverTab[61667]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:46
	_go_fuzz_dep_.CoverTab[61668]++

												m.Reset()
												mi := MessageV2(m)

												if wrapTextUnmarshalV2 {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:51
		_go_fuzz_dep_.CoverTab[61671]++
													err := prototext.UnmarshalOptions{
			AllowPartial: true,
		}.Unmarshal([]byte(s), mi)
		if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:55
			_go_fuzz_dep_.CoverTab[61673]++
														return &ParseError{Message: err.Error()}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:56
			// _ = "end of CoverTab[61673]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:57
			_go_fuzz_dep_.CoverTab[61674]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:57
			// _ = "end of CoverTab[61674]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:57
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:57
		// _ = "end of CoverTab[61671]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:57
		_go_fuzz_dep_.CoverTab[61672]++
													return checkRequiredNotSet(mi)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:58
		// _ = "end of CoverTab[61672]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:59
		_go_fuzz_dep_.CoverTab[61675]++
													if err := newTextParser(s).unmarshalMessage(mi.ProtoReflect(), ""); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:60
			_go_fuzz_dep_.CoverTab[61677]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:61
			// _ = "end of CoverTab[61677]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:62
			_go_fuzz_dep_.CoverTab[61678]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:62
			// _ = "end of CoverTab[61678]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:62
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:62
		// _ = "end of CoverTab[61675]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:62
		_go_fuzz_dep_.CoverTab[61676]++
													return checkRequiredNotSet(mi)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:63
		// _ = "end of CoverTab[61676]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:64
	// _ = "end of CoverTab[61668]"
}

type textParser struct {
	s		string	// remaining input
	done		bool	// whether the parsing is finished (success or error)
	backed		bool	// whether back() was called
	offset, line	int
	cur		token
}

type token struct {
	value		string
	err		*ParseError
	line		int	// line number
	offset		int	// byte number from start of input, not start of line
	unquoted	string	// the unquoted version of value, if it was a quoted string
}

func newTextParser(s string) *textParser {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:83
	_go_fuzz_dep_.CoverTab[61679]++
												p := new(textParser)
												p.s = s
												p.line = 1
												p.cur.line = 1
												return p
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:88
	// _ = "end of CoverTab[61679]"
}

func (p *textParser) unmarshalMessage(m protoreflect.Message, terminator string) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:91
	_go_fuzz_dep_.CoverTab[61680]++
												md := m.Descriptor()
												fds := md.Fields()

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:101
	seen := make(map[protoreflect.FieldNumber]bool)
	for {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:102
		_go_fuzz_dep_.CoverTab[61682]++
													tok := p.next()
													if tok.err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:104
			_go_fuzz_dep_.CoverTab[61693]++
														return tok.err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:105
			// _ = "end of CoverTab[61693]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:106
			_go_fuzz_dep_.CoverTab[61694]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:106
			// _ = "end of CoverTab[61694]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:106
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:106
		// _ = "end of CoverTab[61682]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:106
		_go_fuzz_dep_.CoverTab[61683]++
													if tok.value == terminator {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:107
			_go_fuzz_dep_.CoverTab[61695]++
														break
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:108
			// _ = "end of CoverTab[61695]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:109
			_go_fuzz_dep_.CoverTab[61696]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:109
			// _ = "end of CoverTab[61696]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:109
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:109
		// _ = "end of CoverTab[61683]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:109
		_go_fuzz_dep_.CoverTab[61684]++
													if tok.value == "[" {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:110
			_go_fuzz_dep_.CoverTab[61697]++
														if err := p.unmarshalExtensionOrAny(m, seen); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:111
				_go_fuzz_dep_.CoverTab[61699]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:112
				// _ = "end of CoverTab[61699]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:113
				_go_fuzz_dep_.CoverTab[61700]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:113
				// _ = "end of CoverTab[61700]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:113
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:113
			// _ = "end of CoverTab[61697]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:113
			_go_fuzz_dep_.CoverTab[61698]++
														continue
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:114
			// _ = "end of CoverTab[61698]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:115
			_go_fuzz_dep_.CoverTab[61701]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:115
			// _ = "end of CoverTab[61701]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:115
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:115
		// _ = "end of CoverTab[61684]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:115
		_go_fuzz_dep_.CoverTab[61685]++

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:118
		name := protoreflect.Name(tok.value)
		fd := fds.ByName(name)
		switch {
		case fd == nil:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:121
			_go_fuzz_dep_.CoverTab[61702]++
														gd := fds.ByName(protoreflect.Name(strings.ToLower(string(name))))
														if gd != nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:123
				_go_fuzz_dep_.CoverTab[61706]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:123
				return gd.Kind() == protoreflect.GroupKind
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:123
				// _ = "end of CoverTab[61706]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:123
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:123
				_go_fuzz_dep_.CoverTab[61707]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:123
				return gd.Message().Name() == name
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:123
				// _ = "end of CoverTab[61707]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:123
			}() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:123
				_go_fuzz_dep_.CoverTab[61708]++
															fd = gd
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:124
				// _ = "end of CoverTab[61708]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:125
				_go_fuzz_dep_.CoverTab[61709]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:125
				// _ = "end of CoverTab[61709]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:125
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:125
			// _ = "end of CoverTab[61702]"
		case fd.Kind() == protoreflect.GroupKind && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:126
			_go_fuzz_dep_.CoverTab[61710]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:126
			return fd.Message().Name() != name
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:126
			// _ = "end of CoverTab[61710]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:126
		}():
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:126
			_go_fuzz_dep_.CoverTab[61703]++
														fd = nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:127
			// _ = "end of CoverTab[61703]"
		case fd.IsWeak() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:128
			_go_fuzz_dep_.CoverTab[61711]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:128
			return fd.Message().IsPlaceholder()
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:128
			// _ = "end of CoverTab[61711]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:128
		}():
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:128
			_go_fuzz_dep_.CoverTab[61704]++
														fd = nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:129
			// _ = "end of CoverTab[61704]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:129
		default:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:129
			_go_fuzz_dep_.CoverTab[61705]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:129
			// _ = "end of CoverTab[61705]"
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:130
		// _ = "end of CoverTab[61685]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:130
		_go_fuzz_dep_.CoverTab[61686]++
													if fd == nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:131
			_go_fuzz_dep_.CoverTab[61712]++
														typeName := string(md.FullName())
														if m, ok := m.Interface().(Message); ok {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:133
				_go_fuzz_dep_.CoverTab[61714]++
															t := reflect.TypeOf(m)
															if t.Kind() == reflect.Ptr {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:135
					_go_fuzz_dep_.CoverTab[61715]++
																typeName = t.Elem().String()
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:136
					// _ = "end of CoverTab[61715]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:137
					_go_fuzz_dep_.CoverTab[61716]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:137
					// _ = "end of CoverTab[61716]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:137
				}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:137
				// _ = "end of CoverTab[61714]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:138
				_go_fuzz_dep_.CoverTab[61717]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:138
				// _ = "end of CoverTab[61717]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:138
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:138
			// _ = "end of CoverTab[61712]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:138
			_go_fuzz_dep_.CoverTab[61713]++
														return p.errorf("unknown field name %q in %v", name, typeName)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:139
			// _ = "end of CoverTab[61713]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:140
			_go_fuzz_dep_.CoverTab[61718]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:140
			// _ = "end of CoverTab[61718]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:140
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:140
		// _ = "end of CoverTab[61686]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:140
		_go_fuzz_dep_.CoverTab[61687]++
													if od := fd.ContainingOneof(); od != nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:141
			_go_fuzz_dep_.CoverTab[61719]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:141
			return m.WhichOneof(od) != nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:141
			// _ = "end of CoverTab[61719]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:141
		}() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:141
			_go_fuzz_dep_.CoverTab[61720]++
														return p.errorf("field '%s' would overwrite already parsed oneof '%s'", name, od.Name())
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:142
			// _ = "end of CoverTab[61720]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:143
			_go_fuzz_dep_.CoverTab[61721]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:143
			// _ = "end of CoverTab[61721]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:143
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:143
		// _ = "end of CoverTab[61687]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:143
		_go_fuzz_dep_.CoverTab[61688]++
													if fd.Cardinality() != protoreflect.Repeated && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:144
			_go_fuzz_dep_.CoverTab[61722]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:144
			return seen[fd.Number()]
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:144
			// _ = "end of CoverTab[61722]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:144
		}() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:144
			_go_fuzz_dep_.CoverTab[61723]++
														return p.errorf("non-repeated field %q was repeated", fd.Name())
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:145
			// _ = "end of CoverTab[61723]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:146
			_go_fuzz_dep_.CoverTab[61724]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:146
			// _ = "end of CoverTab[61724]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:146
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:146
		// _ = "end of CoverTab[61688]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:146
		_go_fuzz_dep_.CoverTab[61689]++
													seen[fd.Number()] = true

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:150
		if err := p.checkForColon(fd); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:150
			_go_fuzz_dep_.CoverTab[61725]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:151
			// _ = "end of CoverTab[61725]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:152
			_go_fuzz_dep_.CoverTab[61726]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:152
			// _ = "end of CoverTab[61726]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:152
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:152
		// _ = "end of CoverTab[61689]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:152
		_go_fuzz_dep_.CoverTab[61690]++

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:155
		v := m.Get(fd)
		if !m.Has(fd) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:156
			_go_fuzz_dep_.CoverTab[61727]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:156
			return (fd.IsList() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:156
				_go_fuzz_dep_.CoverTab[61728]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:156
				return fd.IsMap()
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:156
				// _ = "end of CoverTab[61728]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:156
			}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:156
				_go_fuzz_dep_.CoverTab[61729]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:156
				return fd.Message() != nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:156
				// _ = "end of CoverTab[61729]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:156
			}())
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:156
			// _ = "end of CoverTab[61727]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:156
		}() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:156
			_go_fuzz_dep_.CoverTab[61730]++
														v = m.Mutable(fd)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:157
			// _ = "end of CoverTab[61730]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:158
			_go_fuzz_dep_.CoverTab[61731]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:158
			// _ = "end of CoverTab[61731]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:158
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:158
		// _ = "end of CoverTab[61690]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:158
		_go_fuzz_dep_.CoverTab[61691]++
													if v, err = p.unmarshalValue(v, fd); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:159
			_go_fuzz_dep_.CoverTab[61732]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:160
			// _ = "end of CoverTab[61732]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:161
			_go_fuzz_dep_.CoverTab[61733]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:161
			// _ = "end of CoverTab[61733]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:161
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:161
		// _ = "end of CoverTab[61691]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:161
		_go_fuzz_dep_.CoverTab[61692]++
													m.Set(fd, v)

													if err := p.consumeOptionalSeparator(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:164
			_go_fuzz_dep_.CoverTab[61734]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:165
			// _ = "end of CoverTab[61734]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:166
			_go_fuzz_dep_.CoverTab[61735]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:166
			// _ = "end of CoverTab[61735]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:166
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:166
		// _ = "end of CoverTab[61692]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:167
	// _ = "end of CoverTab[61680]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:167
	_go_fuzz_dep_.CoverTab[61681]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:168
	// _ = "end of CoverTab[61681]"
}

func (p *textParser) unmarshalExtensionOrAny(m protoreflect.Message, seen map[protoreflect.FieldNumber]bool) error {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:171
	_go_fuzz_dep_.CoverTab[61736]++
												name, err := p.consumeExtensionOrAnyName()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:173
		_go_fuzz_dep_.CoverTab[61745]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:174
		// _ = "end of CoverTab[61745]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:175
		_go_fuzz_dep_.CoverTab[61746]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:175
		// _ = "end of CoverTab[61746]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:175
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:175
	// _ = "end of CoverTab[61736]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:175
	_go_fuzz_dep_.CoverTab[61737]++

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:178
	if slashIdx := strings.LastIndex(name, "/"); slashIdx >= 0 {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:178
		_go_fuzz_dep_.CoverTab[61747]++
													tok := p.next()
													if tok.err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:180
			_go_fuzz_dep_.CoverTab[61756]++
														return tok.err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:181
			// _ = "end of CoverTab[61756]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:182
			_go_fuzz_dep_.CoverTab[61757]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:182
			// _ = "end of CoverTab[61757]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:182
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:182
		// _ = "end of CoverTab[61747]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:182
		_go_fuzz_dep_.CoverTab[61748]++

													if tok.value == ":" {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:184
			_go_fuzz_dep_.CoverTab[61758]++
														tok = p.next()
														if tok.err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:186
				_go_fuzz_dep_.CoverTab[61759]++
															return tok.err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:187
				// _ = "end of CoverTab[61759]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:188
				_go_fuzz_dep_.CoverTab[61760]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:188
				// _ = "end of CoverTab[61760]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:188
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:188
			// _ = "end of CoverTab[61758]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:189
			_go_fuzz_dep_.CoverTab[61761]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:189
			// _ = "end of CoverTab[61761]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:189
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:189
		// _ = "end of CoverTab[61748]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:189
		_go_fuzz_dep_.CoverTab[61749]++

													var terminator string
													switch tok.value {
		case "<":
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:193
			_go_fuzz_dep_.CoverTab[61762]++
														terminator = ">"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:194
			// _ = "end of CoverTab[61762]"
		case "{":
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:195
			_go_fuzz_dep_.CoverTab[61763]++
														terminator = "}"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:196
			// _ = "end of CoverTab[61763]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:197
			_go_fuzz_dep_.CoverTab[61764]++
														return p.errorf("expected '{' or '<', found %q", tok.value)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:198
			// _ = "end of CoverTab[61764]"
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:199
		// _ = "end of CoverTab[61749]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:199
		_go_fuzz_dep_.CoverTab[61750]++

													mt, err := protoregistry.GlobalTypes.FindMessageByURL(name)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:202
			_go_fuzz_dep_.CoverTab[61765]++
														return p.errorf("unrecognized message %q in google.protobuf.Any", name[slashIdx+len("/"):])
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:203
			// _ = "end of CoverTab[61765]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:204
			_go_fuzz_dep_.CoverTab[61766]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:204
			// _ = "end of CoverTab[61766]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:204
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:204
		// _ = "end of CoverTab[61750]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:204
		_go_fuzz_dep_.CoverTab[61751]++
													m2 := mt.New()
													if err := p.unmarshalMessage(m2, terminator); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:206
			_go_fuzz_dep_.CoverTab[61767]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:207
			// _ = "end of CoverTab[61767]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:208
			_go_fuzz_dep_.CoverTab[61768]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:208
			// _ = "end of CoverTab[61768]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:208
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:208
		// _ = "end of CoverTab[61751]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:208
		_go_fuzz_dep_.CoverTab[61752]++
													b, err := protoV2.Marshal(m2.Interface())
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:210
			_go_fuzz_dep_.CoverTab[61769]++
														return p.errorf("failed to marshal message of type %q: %v", name[slashIdx+len("/"):], err)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:211
			// _ = "end of CoverTab[61769]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:212
			_go_fuzz_dep_.CoverTab[61770]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:212
			// _ = "end of CoverTab[61770]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:212
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:212
		// _ = "end of CoverTab[61752]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:212
		_go_fuzz_dep_.CoverTab[61753]++

													urlFD := m.Descriptor().Fields().ByName("type_url")
													valFD := m.Descriptor().Fields().ByName("value")
													if seen[urlFD.Number()] {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:216
			_go_fuzz_dep_.CoverTab[61771]++
														return p.errorf("Any message unpacked multiple times, or %q already set", urlFD.Name())
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:217
			// _ = "end of CoverTab[61771]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:218
			_go_fuzz_dep_.CoverTab[61772]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:218
			// _ = "end of CoverTab[61772]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:218
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:218
		// _ = "end of CoverTab[61753]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:218
		_go_fuzz_dep_.CoverTab[61754]++
													if seen[valFD.Number()] {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:219
			_go_fuzz_dep_.CoverTab[61773]++
														return p.errorf("Any message unpacked multiple times, or %q already set", valFD.Name())
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:220
			// _ = "end of CoverTab[61773]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:221
			_go_fuzz_dep_.CoverTab[61774]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:221
			// _ = "end of CoverTab[61774]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:221
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:221
		// _ = "end of CoverTab[61754]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:221
		_go_fuzz_dep_.CoverTab[61755]++
													m.Set(urlFD, protoreflect.ValueOfString(name))
													m.Set(valFD, protoreflect.ValueOfBytes(b))
													seen[urlFD.Number()] = true
													seen[valFD.Number()] = true
													return nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:226
		// _ = "end of CoverTab[61755]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:227
		_go_fuzz_dep_.CoverTab[61775]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:227
		// _ = "end of CoverTab[61775]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:227
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:227
	// _ = "end of CoverTab[61737]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:227
	_go_fuzz_dep_.CoverTab[61738]++

												xname := protoreflect.FullName(name)
												xt, _ := protoregistry.GlobalTypes.FindExtensionByName(xname)
												if xt == nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:231
		_go_fuzz_dep_.CoverTab[61776]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:231
		return isMessageSet(m.Descriptor())
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:231
		// _ = "end of CoverTab[61776]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:231
	}() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:231
		_go_fuzz_dep_.CoverTab[61777]++
													xt, _ = protoregistry.GlobalTypes.FindExtensionByName(xname.Append("message_set_extension"))
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:232
		// _ = "end of CoverTab[61777]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:233
		_go_fuzz_dep_.CoverTab[61778]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:233
		// _ = "end of CoverTab[61778]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:233
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:233
	// _ = "end of CoverTab[61738]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:233
	_go_fuzz_dep_.CoverTab[61739]++
												if xt == nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:234
		_go_fuzz_dep_.CoverTab[61779]++
													return p.errorf("unrecognized extension %q", name)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:235
		// _ = "end of CoverTab[61779]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:236
		_go_fuzz_dep_.CoverTab[61780]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:236
		// _ = "end of CoverTab[61780]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:236
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:236
	// _ = "end of CoverTab[61739]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:236
	_go_fuzz_dep_.CoverTab[61740]++
												fd := xt.TypeDescriptor()
												if fd.ContainingMessage().FullName() != m.Descriptor().FullName() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:238
		_go_fuzz_dep_.CoverTab[61781]++
													return p.errorf("extension field %q does not extend message %q", name, m.Descriptor().FullName())
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:239
		// _ = "end of CoverTab[61781]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:240
		_go_fuzz_dep_.CoverTab[61782]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:240
		// _ = "end of CoverTab[61782]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:240
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:240
	// _ = "end of CoverTab[61740]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:240
	_go_fuzz_dep_.CoverTab[61741]++

												if err := p.checkForColon(fd); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:242
		_go_fuzz_dep_.CoverTab[61783]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:243
		// _ = "end of CoverTab[61783]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:244
		_go_fuzz_dep_.CoverTab[61784]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:244
		// _ = "end of CoverTab[61784]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:244
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:244
	// _ = "end of CoverTab[61741]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:244
	_go_fuzz_dep_.CoverTab[61742]++

												v := m.Get(fd)
												if !m.Has(fd) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:247
		_go_fuzz_dep_.CoverTab[61785]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:247
		return (fd.IsList() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:247
			_go_fuzz_dep_.CoverTab[61786]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:247
			return fd.IsMap()
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:247
			// _ = "end of CoverTab[61786]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:247
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:247
			_go_fuzz_dep_.CoverTab[61787]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:247
			return fd.Message() != nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:247
			// _ = "end of CoverTab[61787]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:247
		}())
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:247
		// _ = "end of CoverTab[61785]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:247
	}() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:247
		_go_fuzz_dep_.CoverTab[61788]++
													v = m.Mutable(fd)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:248
		// _ = "end of CoverTab[61788]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:249
		_go_fuzz_dep_.CoverTab[61789]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:249
		// _ = "end of CoverTab[61789]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:249
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:249
	// _ = "end of CoverTab[61742]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:249
	_go_fuzz_dep_.CoverTab[61743]++
												v, err = p.unmarshalValue(v, fd)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:251
		_go_fuzz_dep_.CoverTab[61790]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:252
		// _ = "end of CoverTab[61790]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:253
		_go_fuzz_dep_.CoverTab[61791]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:253
		// _ = "end of CoverTab[61791]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:253
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:253
	// _ = "end of CoverTab[61743]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:253
	_go_fuzz_dep_.CoverTab[61744]++
												m.Set(fd, v)
												return p.consumeOptionalSeparator()
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:255
	// _ = "end of CoverTab[61744]"
}

func (p *textParser) unmarshalValue(v protoreflect.Value, fd protoreflect.FieldDescriptor) (protoreflect.Value, error) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:258
	_go_fuzz_dep_.CoverTab[61792]++
												tok := p.next()
												if tok.err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:260
		_go_fuzz_dep_.CoverTab[61795]++
													return v, tok.err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:261
		// _ = "end of CoverTab[61795]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:262
		_go_fuzz_dep_.CoverTab[61796]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:262
		// _ = "end of CoverTab[61796]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:262
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:262
	// _ = "end of CoverTab[61792]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:262
	_go_fuzz_dep_.CoverTab[61793]++
												if tok.value == "" {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:263
		_go_fuzz_dep_.CoverTab[61797]++
													return v, p.errorf("unexpected EOF")
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:264
		// _ = "end of CoverTab[61797]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:265
		_go_fuzz_dep_.CoverTab[61798]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:265
		// _ = "end of CoverTab[61798]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:265
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:265
	// _ = "end of CoverTab[61793]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:265
	_go_fuzz_dep_.CoverTab[61794]++

												switch {
	case fd.IsList():
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:268
		_go_fuzz_dep_.CoverTab[61799]++
													lv := v.List()
													var err error
													if tok.value == "[" {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:271
			_go_fuzz_dep_.CoverTab[61806]++

														for {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:273
				_go_fuzz_dep_.CoverTab[61808]++
															vv := lv.NewElement()
															vv, err = p.unmarshalSingularValue(vv, fd)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:276
					_go_fuzz_dep_.CoverTab[61812]++
																return v, err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:277
					// _ = "end of CoverTab[61812]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:278
					_go_fuzz_dep_.CoverTab[61813]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:278
					// _ = "end of CoverTab[61813]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:278
				}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:278
				// _ = "end of CoverTab[61808]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:278
				_go_fuzz_dep_.CoverTab[61809]++
															lv.Append(vv)

															tok := p.next()
															if tok.err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:282
					_go_fuzz_dep_.CoverTab[61814]++
																return v, tok.err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:283
					// _ = "end of CoverTab[61814]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:284
					_go_fuzz_dep_.CoverTab[61815]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:284
					// _ = "end of CoverTab[61815]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:284
				}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:284
				// _ = "end of CoverTab[61809]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:284
				_go_fuzz_dep_.CoverTab[61810]++
															if tok.value == "]" {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:285
					_go_fuzz_dep_.CoverTab[61816]++
																break
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:286
					// _ = "end of CoverTab[61816]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:287
					_go_fuzz_dep_.CoverTab[61817]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:287
					// _ = "end of CoverTab[61817]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:287
				}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:287
				// _ = "end of CoverTab[61810]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:287
				_go_fuzz_dep_.CoverTab[61811]++
															if tok.value != "," {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:288
					_go_fuzz_dep_.CoverTab[61818]++
																return v, p.errorf("Expected ']' or ',' found %q", tok.value)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:289
					// _ = "end of CoverTab[61818]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:290
					_go_fuzz_dep_.CoverTab[61819]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:290
					// _ = "end of CoverTab[61819]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:290
				}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:290
				// _ = "end of CoverTab[61811]"
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:291
			// _ = "end of CoverTab[61806]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:291
			_go_fuzz_dep_.CoverTab[61807]++
														return v, nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:292
			// _ = "end of CoverTab[61807]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:293
			_go_fuzz_dep_.CoverTab[61820]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:293
			// _ = "end of CoverTab[61820]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:293
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:293
		// _ = "end of CoverTab[61799]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:293
		_go_fuzz_dep_.CoverTab[61800]++

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:296
		p.back()
		vv := lv.NewElement()
		vv, err = p.unmarshalSingularValue(vv, fd)
		if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:299
			_go_fuzz_dep_.CoverTab[61821]++
														return v, err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:300
			// _ = "end of CoverTab[61821]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:301
			_go_fuzz_dep_.CoverTab[61822]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:301
			// _ = "end of CoverTab[61822]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:301
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:301
		// _ = "end of CoverTab[61800]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:301
		_go_fuzz_dep_.CoverTab[61801]++
													lv.Append(vv)
													return v, nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:303
		// _ = "end of CoverTab[61801]"
	case fd.IsMap():
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:304
		_go_fuzz_dep_.CoverTab[61802]++
		// The map entry should be this sequence of tokens:
		//	< key : KEY value : VALUE >
		// However, implementations may omit key or value, and technically
		// we should support them in any order.
		var terminator string
		switch tok.value {
		case "<":
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:311
			_go_fuzz_dep_.CoverTab[61823]++
														terminator = ">"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:312
			// _ = "end of CoverTab[61823]"
		case "{":
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:313
			_go_fuzz_dep_.CoverTab[61824]++
														terminator = "}"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:314
			// _ = "end of CoverTab[61824]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:315
			_go_fuzz_dep_.CoverTab[61825]++
														return v, p.errorf("expected '{' or '<', found %q", tok.value)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:316
			// _ = "end of CoverTab[61825]"
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:317
		// _ = "end of CoverTab[61802]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:317
		_go_fuzz_dep_.CoverTab[61803]++

													keyFD := fd.MapKey()
													valFD := fd.MapValue()

													mv := v.Map()
													kv := keyFD.Default()
													vv := mv.NewValue()
													for {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:325
			_go_fuzz_dep_.CoverTab[61826]++
														tok := p.next()
														if tok.err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:327
				_go_fuzz_dep_.CoverTab[61829]++
															return v, tok.err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:328
				// _ = "end of CoverTab[61829]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:329
				_go_fuzz_dep_.CoverTab[61830]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:329
				// _ = "end of CoverTab[61830]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:329
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:329
			// _ = "end of CoverTab[61826]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:329
			_go_fuzz_dep_.CoverTab[61827]++
														if tok.value == terminator {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:330
				_go_fuzz_dep_.CoverTab[61831]++
															break
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:331
				// _ = "end of CoverTab[61831]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:332
				_go_fuzz_dep_.CoverTab[61832]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:332
				// _ = "end of CoverTab[61832]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:332
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:332
			// _ = "end of CoverTab[61827]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:332
			_go_fuzz_dep_.CoverTab[61828]++
														var err error
														switch tok.value {
			case "key":
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:335
				_go_fuzz_dep_.CoverTab[61833]++
															if err := p.consumeToken(":"); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:336
					_go_fuzz_dep_.CoverTab[61840]++
																return v, err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:337
					// _ = "end of CoverTab[61840]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:338
					_go_fuzz_dep_.CoverTab[61841]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:338
					// _ = "end of CoverTab[61841]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:338
				}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:338
				// _ = "end of CoverTab[61833]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:338
				_go_fuzz_dep_.CoverTab[61834]++
															if kv, err = p.unmarshalSingularValue(kv, keyFD); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:339
					_go_fuzz_dep_.CoverTab[61842]++
																return v, err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:340
					// _ = "end of CoverTab[61842]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:341
					_go_fuzz_dep_.CoverTab[61843]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:341
					// _ = "end of CoverTab[61843]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:341
				}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:341
				// _ = "end of CoverTab[61834]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:341
				_go_fuzz_dep_.CoverTab[61835]++
															if err := p.consumeOptionalSeparator(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:342
					_go_fuzz_dep_.CoverTab[61844]++
																return v, err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:343
					// _ = "end of CoverTab[61844]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:344
					_go_fuzz_dep_.CoverTab[61845]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:344
					// _ = "end of CoverTab[61845]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:344
				}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:344
				// _ = "end of CoverTab[61835]"
			case "value":
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:345
				_go_fuzz_dep_.CoverTab[61836]++
															if err := p.checkForColon(valFD); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:346
					_go_fuzz_dep_.CoverTab[61846]++
																return v, err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:347
					// _ = "end of CoverTab[61846]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:348
					_go_fuzz_dep_.CoverTab[61847]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:348
					// _ = "end of CoverTab[61847]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:348
				}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:348
				// _ = "end of CoverTab[61836]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:348
				_go_fuzz_dep_.CoverTab[61837]++
															if vv, err = p.unmarshalSingularValue(vv, valFD); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:349
					_go_fuzz_dep_.CoverTab[61848]++
																return v, err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:350
					// _ = "end of CoverTab[61848]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:351
					_go_fuzz_dep_.CoverTab[61849]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:351
					// _ = "end of CoverTab[61849]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:351
				}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:351
				// _ = "end of CoverTab[61837]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:351
				_go_fuzz_dep_.CoverTab[61838]++
															if err := p.consumeOptionalSeparator(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:352
					_go_fuzz_dep_.CoverTab[61850]++
																return v, err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:353
					// _ = "end of CoverTab[61850]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:354
					_go_fuzz_dep_.CoverTab[61851]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:354
					// _ = "end of CoverTab[61851]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:354
				}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:354
				// _ = "end of CoverTab[61838]"
			default:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:355
				_go_fuzz_dep_.CoverTab[61839]++
															p.back()
															return v, p.errorf(`expected "key", "value", or %q, found %q`, terminator, tok.value)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:357
				// _ = "end of CoverTab[61839]"
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:358
			// _ = "end of CoverTab[61828]"
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:359
		// _ = "end of CoverTab[61803]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:359
		_go_fuzz_dep_.CoverTab[61804]++
													mv.Set(kv.MapKey(), vv)
													return v, nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:361
		// _ = "end of CoverTab[61804]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:362
		_go_fuzz_dep_.CoverTab[61805]++
													p.back()
													return p.unmarshalSingularValue(v, fd)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:364
		// _ = "end of CoverTab[61805]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:365
	// _ = "end of CoverTab[61794]"
}

func (p *textParser) unmarshalSingularValue(v protoreflect.Value, fd protoreflect.FieldDescriptor) (protoreflect.Value, error) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:368
	_go_fuzz_dep_.CoverTab[61852]++
												tok := p.next()
												if tok.err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:370
		_go_fuzz_dep_.CoverTab[61856]++
													return v, tok.err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:371
		// _ = "end of CoverTab[61856]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:372
		_go_fuzz_dep_.CoverTab[61857]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:372
		// _ = "end of CoverTab[61857]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:372
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:372
	// _ = "end of CoverTab[61852]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:372
	_go_fuzz_dep_.CoverTab[61853]++
												if tok.value == "" {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:373
		_go_fuzz_dep_.CoverTab[61858]++
													return v, p.errorf("unexpected EOF")
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:374
		// _ = "end of CoverTab[61858]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:375
		_go_fuzz_dep_.CoverTab[61859]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:375
		// _ = "end of CoverTab[61859]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:375
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:375
	// _ = "end of CoverTab[61853]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:375
	_go_fuzz_dep_.CoverTab[61854]++

												switch fd.Kind() {
	case protoreflect.BoolKind:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:378
		_go_fuzz_dep_.CoverTab[61860]++
													switch tok.value {
		case "true", "1", "t", "True":
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:380
			_go_fuzz_dep_.CoverTab[61878]++
														return protoreflect.ValueOfBool(true), nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:381
			// _ = "end of CoverTab[61878]"
		case "false", "0", "f", "False":
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:382
			_go_fuzz_dep_.CoverTab[61879]++
														return protoreflect.ValueOfBool(false), nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:383
			// _ = "end of CoverTab[61879]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:383
		default:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:383
			_go_fuzz_dep_.CoverTab[61880]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:383
			// _ = "end of CoverTab[61880]"
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:384
		// _ = "end of CoverTab[61860]"
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:385
		_go_fuzz_dep_.CoverTab[61861]++
													if x, err := strconv.ParseInt(tok.value, 0, 32); err == nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:386
			_go_fuzz_dep_.CoverTab[61881]++
														return protoreflect.ValueOfInt32(int32(x)), nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:387
			// _ = "end of CoverTab[61881]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:388
			_go_fuzz_dep_.CoverTab[61882]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:388
			// _ = "end of CoverTab[61882]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:388
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:388
		// _ = "end of CoverTab[61861]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:388
		_go_fuzz_dep_.CoverTab[61862]++

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:393
		if strings.HasPrefix(tok.value, "0x") {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:393
			_go_fuzz_dep_.CoverTab[61883]++
														if x, err := strconv.ParseUint(tok.value, 0, 32); err == nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:394
				_go_fuzz_dep_.CoverTab[61884]++
															return protoreflect.ValueOfInt32(int32(-(int64(^x) + 1))), nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:395
				// _ = "end of CoverTab[61884]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:396
				_go_fuzz_dep_.CoverTab[61885]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:396
				// _ = "end of CoverTab[61885]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:396
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:396
			// _ = "end of CoverTab[61883]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:397
			_go_fuzz_dep_.CoverTab[61886]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:397
			// _ = "end of CoverTab[61886]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:397
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:397
		// _ = "end of CoverTab[61862]"
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:398
		_go_fuzz_dep_.CoverTab[61863]++
													if x, err := strconv.ParseInt(tok.value, 0, 64); err == nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:399
			_go_fuzz_dep_.CoverTab[61887]++
														return protoreflect.ValueOfInt64(int64(x)), nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:400
			// _ = "end of CoverTab[61887]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:401
			_go_fuzz_dep_.CoverTab[61888]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:401
			// _ = "end of CoverTab[61888]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:401
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:401
		// _ = "end of CoverTab[61863]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:401
		_go_fuzz_dep_.CoverTab[61864]++

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:406
		if strings.HasPrefix(tok.value, "0x") {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:406
			_go_fuzz_dep_.CoverTab[61889]++
														if x, err := strconv.ParseUint(tok.value, 0, 64); err == nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:407
				_go_fuzz_dep_.CoverTab[61890]++
															return protoreflect.ValueOfInt64(int64(-(int64(^x) + 1))), nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:408
				// _ = "end of CoverTab[61890]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:409
				_go_fuzz_dep_.CoverTab[61891]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:409
				// _ = "end of CoverTab[61891]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:409
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:409
			// _ = "end of CoverTab[61889]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:410
			_go_fuzz_dep_.CoverTab[61892]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:410
			// _ = "end of CoverTab[61892]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:410
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:410
		// _ = "end of CoverTab[61864]"
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:411
		_go_fuzz_dep_.CoverTab[61865]++
													if x, err := strconv.ParseUint(tok.value, 0, 32); err == nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:412
			_go_fuzz_dep_.CoverTab[61893]++
														return protoreflect.ValueOfUint32(uint32(x)), nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:413
			// _ = "end of CoverTab[61893]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:414
			_go_fuzz_dep_.CoverTab[61894]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:414
			// _ = "end of CoverTab[61894]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:414
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:414
		// _ = "end of CoverTab[61865]"
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:415
		_go_fuzz_dep_.CoverTab[61866]++
													if x, err := strconv.ParseUint(tok.value, 0, 64); err == nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:416
			_go_fuzz_dep_.CoverTab[61895]++
														return protoreflect.ValueOfUint64(uint64(x)), nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:417
			// _ = "end of CoverTab[61895]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:418
			_go_fuzz_dep_.CoverTab[61896]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:418
			// _ = "end of CoverTab[61896]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:418
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:418
		// _ = "end of CoverTab[61866]"
	case protoreflect.FloatKind:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:419
		_go_fuzz_dep_.CoverTab[61867]++

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:422
		v := tok.value
		if strings.HasSuffix(v, "f") && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:423
			_go_fuzz_dep_.CoverTab[61897]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:423
			return v != "-inf"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:423
			// _ = "end of CoverTab[61897]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:423
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:423
			_go_fuzz_dep_.CoverTab[61898]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:423
			return v != "inf"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:423
			// _ = "end of CoverTab[61898]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:423
		}() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:423
			_go_fuzz_dep_.CoverTab[61899]++
														v = v[:len(v)-len("f")]
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:424
			// _ = "end of CoverTab[61899]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:425
			_go_fuzz_dep_.CoverTab[61900]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:425
			// _ = "end of CoverTab[61900]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:425
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:425
		// _ = "end of CoverTab[61867]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:425
		_go_fuzz_dep_.CoverTab[61868]++
													if x, err := strconv.ParseFloat(v, 32); err == nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:426
			_go_fuzz_dep_.CoverTab[61901]++
														return protoreflect.ValueOfFloat32(float32(x)), nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:427
			// _ = "end of CoverTab[61901]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:428
			_go_fuzz_dep_.CoverTab[61902]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:428
			// _ = "end of CoverTab[61902]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:428
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:428
		// _ = "end of CoverTab[61868]"
	case protoreflect.DoubleKind:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:429
		_go_fuzz_dep_.CoverTab[61869]++

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:432
		v := tok.value
		if strings.HasSuffix(v, "f") && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:433
			_go_fuzz_dep_.CoverTab[61903]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:433
			return v != "-inf"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:433
			// _ = "end of CoverTab[61903]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:433
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:433
			_go_fuzz_dep_.CoverTab[61904]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:433
			return v != "inf"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:433
			// _ = "end of CoverTab[61904]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:433
		}() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:433
			_go_fuzz_dep_.CoverTab[61905]++
														v = v[:len(v)-len("f")]
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:434
			// _ = "end of CoverTab[61905]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:435
			_go_fuzz_dep_.CoverTab[61906]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:435
			// _ = "end of CoverTab[61906]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:435
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:435
		// _ = "end of CoverTab[61869]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:435
		_go_fuzz_dep_.CoverTab[61870]++
													if x, err := strconv.ParseFloat(v, 64); err == nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:436
			_go_fuzz_dep_.CoverTab[61907]++
														return protoreflect.ValueOfFloat64(float64(x)), nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:437
			// _ = "end of CoverTab[61907]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:438
			_go_fuzz_dep_.CoverTab[61908]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:438
			// _ = "end of CoverTab[61908]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:438
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:438
		// _ = "end of CoverTab[61870]"
	case protoreflect.StringKind:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:439
		_go_fuzz_dep_.CoverTab[61871]++
													if isQuote(tok.value[0]) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:440
			_go_fuzz_dep_.CoverTab[61909]++
														return protoreflect.ValueOfString(tok.unquoted), nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:441
			// _ = "end of CoverTab[61909]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:442
			_go_fuzz_dep_.CoverTab[61910]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:442
			// _ = "end of CoverTab[61910]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:442
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:442
		// _ = "end of CoverTab[61871]"
	case protoreflect.BytesKind:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:443
		_go_fuzz_dep_.CoverTab[61872]++
													if isQuote(tok.value[0]) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:444
			_go_fuzz_dep_.CoverTab[61911]++
														return protoreflect.ValueOfBytes([]byte(tok.unquoted)), nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:445
			// _ = "end of CoverTab[61911]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:446
			_go_fuzz_dep_.CoverTab[61912]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:446
			// _ = "end of CoverTab[61912]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:446
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:446
		// _ = "end of CoverTab[61872]"
	case protoreflect.EnumKind:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:447
		_go_fuzz_dep_.CoverTab[61873]++
													if x, err := strconv.ParseInt(tok.value, 0, 32); err == nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:448
			_go_fuzz_dep_.CoverTab[61913]++
														return protoreflect.ValueOfEnum(protoreflect.EnumNumber(x)), nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:449
			// _ = "end of CoverTab[61913]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:450
			_go_fuzz_dep_.CoverTab[61914]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:450
			// _ = "end of CoverTab[61914]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:450
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:450
		// _ = "end of CoverTab[61873]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:450
		_go_fuzz_dep_.CoverTab[61874]++
													vd := fd.Enum().Values().ByName(protoreflect.Name(tok.value))
													if vd != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:452
			_go_fuzz_dep_.CoverTab[61915]++
														return protoreflect.ValueOfEnum(vd.Number()), nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:453
			// _ = "end of CoverTab[61915]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:454
			_go_fuzz_dep_.CoverTab[61916]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:454
			// _ = "end of CoverTab[61916]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:454
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:454
		// _ = "end of CoverTab[61874]"
	case protoreflect.MessageKind, protoreflect.GroupKind:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:455
		_go_fuzz_dep_.CoverTab[61875]++
													var terminator string
													switch tok.value {
		case "{":
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:458
			_go_fuzz_dep_.CoverTab[61917]++
														terminator = "}"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:459
			// _ = "end of CoverTab[61917]"
		case "<":
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:460
			_go_fuzz_dep_.CoverTab[61918]++
														terminator = ">"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:461
			// _ = "end of CoverTab[61918]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:462
			_go_fuzz_dep_.CoverTab[61919]++
														return v, p.errorf("expected '{' or '<', found %q", tok.value)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:463
			// _ = "end of CoverTab[61919]"
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:464
		// _ = "end of CoverTab[61875]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:464
		_go_fuzz_dep_.CoverTab[61876]++
													err := p.unmarshalMessage(v.Message(), terminator)
													return v, err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:466
		// _ = "end of CoverTab[61876]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:467
		_go_fuzz_dep_.CoverTab[61877]++
													panic(fmt.Sprintf("invalid kind %v", fd.Kind()))
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:468
		// _ = "end of CoverTab[61877]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:469
	// _ = "end of CoverTab[61854]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:469
	_go_fuzz_dep_.CoverTab[61855]++
												return v, p.errorf("invalid %v: %v", fd.Kind(), tok.value)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:470
	// _ = "end of CoverTab[61855]"
}

// Consume a ':' from the input stream (if the next token is a colon),
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:473
// returning an error if a colon is needed but not present.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:475
func (p *textParser) checkForColon(fd protoreflect.FieldDescriptor) *ParseError {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:475
	_go_fuzz_dep_.CoverTab[61920]++
												tok := p.next()
												if tok.err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:477
		_go_fuzz_dep_.CoverTab[61923]++
													return tok.err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:478
		// _ = "end of CoverTab[61923]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:479
		_go_fuzz_dep_.CoverTab[61924]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:479
		// _ = "end of CoverTab[61924]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:479
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:479
	// _ = "end of CoverTab[61920]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:479
	_go_fuzz_dep_.CoverTab[61921]++
												if tok.value != ":" {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:480
		_go_fuzz_dep_.CoverTab[61925]++
													if fd.Message() == nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:481
			_go_fuzz_dep_.CoverTab[61927]++
														return p.errorf("expected ':', found %q", tok.value)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:482
			// _ = "end of CoverTab[61927]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:483
			_go_fuzz_dep_.CoverTab[61928]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:483
			// _ = "end of CoverTab[61928]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:483
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:483
		// _ = "end of CoverTab[61925]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:483
		_go_fuzz_dep_.CoverTab[61926]++
													p.back()
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:484
		// _ = "end of CoverTab[61926]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:485
		_go_fuzz_dep_.CoverTab[61929]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:485
		// _ = "end of CoverTab[61929]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:485
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:485
	// _ = "end of CoverTab[61921]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:485
	_go_fuzz_dep_.CoverTab[61922]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:486
	// _ = "end of CoverTab[61922]"
}

// consumeExtensionOrAnyName consumes an extension name or an Any type URL and
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:489
// the following ']'. It returns the name or URL consumed.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:491
func (p *textParser) consumeExtensionOrAnyName() (string, error) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:491
	_go_fuzz_dep_.CoverTab[61930]++
												tok := p.next()
												if tok.err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:493
		_go_fuzz_dep_.CoverTab[61934]++
													return "", tok.err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:494
		// _ = "end of CoverTab[61934]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:495
		_go_fuzz_dep_.CoverTab[61935]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:495
		// _ = "end of CoverTab[61935]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:495
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:495
	// _ = "end of CoverTab[61930]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:495
	_go_fuzz_dep_.CoverTab[61931]++

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:498
	if len(tok.value) > 2 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:498
		_go_fuzz_dep_.CoverTab[61936]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:498
		return isQuote(tok.value[0])
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:498
		// _ = "end of CoverTab[61936]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:498
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:498
		_go_fuzz_dep_.CoverTab[61937]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:498
		return tok.value[len(tok.value)-1] == tok.value[0]
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:498
		// _ = "end of CoverTab[61937]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:498
	}() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:498
		_go_fuzz_dep_.CoverTab[61938]++
													name, err := unquoteC(tok.value[1:len(tok.value)-1], rune(tok.value[0]))
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:500
			_go_fuzz_dep_.CoverTab[61940]++
														return "", err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:501
			// _ = "end of CoverTab[61940]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:502
			_go_fuzz_dep_.CoverTab[61941]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:502
			// _ = "end of CoverTab[61941]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:502
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:502
		// _ = "end of CoverTab[61938]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:502
		_go_fuzz_dep_.CoverTab[61939]++
													return name, p.consumeToken("]")
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:503
		// _ = "end of CoverTab[61939]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:504
		_go_fuzz_dep_.CoverTab[61942]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:504
		// _ = "end of CoverTab[61942]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:504
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:504
	// _ = "end of CoverTab[61931]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:504
	_go_fuzz_dep_.CoverTab[61932]++

	// Consume everything up to "]"
	var parts []string
	for tok.value != "]" {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:508
		_go_fuzz_dep_.CoverTab[61943]++
													parts = append(parts, tok.value)
													tok = p.next()
													if tok.err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:511
			_go_fuzz_dep_.CoverTab[61945]++
														return "", p.errorf("unrecognized type_url or extension name: %s", tok.err)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:512
			// _ = "end of CoverTab[61945]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:513
			_go_fuzz_dep_.CoverTab[61946]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:513
			// _ = "end of CoverTab[61946]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:513
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:513
		// _ = "end of CoverTab[61943]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:513
		_go_fuzz_dep_.CoverTab[61944]++
													if p.done && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:514
			_go_fuzz_dep_.CoverTab[61947]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:514
			return tok.value != "]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:514
			// _ = "end of CoverTab[61947]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:514
		}() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:514
			_go_fuzz_dep_.CoverTab[61948]++
														return "", p.errorf("unclosed type_url or extension name")
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:515
			// _ = "end of CoverTab[61948]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:516
			_go_fuzz_dep_.CoverTab[61949]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:516
			// _ = "end of CoverTab[61949]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:516
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:516
		// _ = "end of CoverTab[61944]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:517
	// _ = "end of CoverTab[61932]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:517
	_go_fuzz_dep_.CoverTab[61933]++
												return strings.Join(parts, ""), nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:518
	// _ = "end of CoverTab[61933]"
}

// consumeOptionalSeparator consumes an optional semicolon or comma.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:521
// It is used in unmarshalMessage to provide backward compatibility.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:523
func (p *textParser) consumeOptionalSeparator() error {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:523
	_go_fuzz_dep_.CoverTab[61950]++
												tok := p.next()
												if tok.err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:525
		_go_fuzz_dep_.CoverTab[61953]++
													return tok.err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:526
		// _ = "end of CoverTab[61953]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:527
		_go_fuzz_dep_.CoverTab[61954]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:527
		// _ = "end of CoverTab[61954]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:527
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:527
	// _ = "end of CoverTab[61950]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:527
	_go_fuzz_dep_.CoverTab[61951]++
												if tok.value != ";" && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:528
		_go_fuzz_dep_.CoverTab[61955]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:528
		return tok.value != ","
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:528
		// _ = "end of CoverTab[61955]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:528
	}() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:528
		_go_fuzz_dep_.CoverTab[61956]++
													p.back()
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:529
		// _ = "end of CoverTab[61956]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:530
		_go_fuzz_dep_.CoverTab[61957]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:530
		// _ = "end of CoverTab[61957]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:530
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:530
	// _ = "end of CoverTab[61951]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:530
	_go_fuzz_dep_.CoverTab[61952]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:531
	// _ = "end of CoverTab[61952]"
}

func (p *textParser) errorf(format string, a ...interface{}) *ParseError {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:534
	_go_fuzz_dep_.CoverTab[61958]++
												pe := &ParseError{fmt.Sprintf(format, a...), p.cur.line, p.cur.offset}
												p.cur.err = pe
												p.done = true
												return pe
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:538
	// _ = "end of CoverTab[61958]"
}

func (p *textParser) skipWhitespace() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:541
	_go_fuzz_dep_.CoverTab[61959]++
												i := 0
												for i < len(p.s) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:543
		_go_fuzz_dep_.CoverTab[61961]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:543
		return (isWhitespace(p.s[i]) || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:543
			_go_fuzz_dep_.CoverTab[61962]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:543
			return p.s[i] == '#'
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:543
			// _ = "end of CoverTab[61962]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:543
		}())
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:543
		// _ = "end of CoverTab[61961]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:543
	}() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:543
		_go_fuzz_dep_.CoverTab[61963]++
													if p.s[i] == '#' {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:544
			_go_fuzz_dep_.CoverTab[61966]++

														for i < len(p.s) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:546
				_go_fuzz_dep_.CoverTab[61968]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:546
				return p.s[i] != '\n'
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:546
				// _ = "end of CoverTab[61968]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:546
			}() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:546
				_go_fuzz_dep_.CoverTab[61969]++
															i++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:547
				// _ = "end of CoverTab[61969]"
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:548
			// _ = "end of CoverTab[61966]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:548
			_go_fuzz_dep_.CoverTab[61967]++
														if i == len(p.s) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:549
				_go_fuzz_dep_.CoverTab[61970]++
															break
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:550
				// _ = "end of CoverTab[61970]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:551
				_go_fuzz_dep_.CoverTab[61971]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:551
				// _ = "end of CoverTab[61971]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:551
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:551
			// _ = "end of CoverTab[61967]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:552
			_go_fuzz_dep_.CoverTab[61972]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:552
			// _ = "end of CoverTab[61972]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:552
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:552
		// _ = "end of CoverTab[61963]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:552
		_go_fuzz_dep_.CoverTab[61964]++
													if p.s[i] == '\n' {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:553
			_go_fuzz_dep_.CoverTab[61973]++
														p.line++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:554
			// _ = "end of CoverTab[61973]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:555
			_go_fuzz_dep_.CoverTab[61974]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:555
			// _ = "end of CoverTab[61974]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:555
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:555
		// _ = "end of CoverTab[61964]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:555
		_go_fuzz_dep_.CoverTab[61965]++
													i++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:556
		// _ = "end of CoverTab[61965]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:557
	// _ = "end of CoverTab[61959]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:557
	_go_fuzz_dep_.CoverTab[61960]++
												p.offset += i
												p.s = p.s[i:len(p.s)]
												if len(p.s) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:560
		_go_fuzz_dep_.CoverTab[61975]++
													p.done = true
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:561
		// _ = "end of CoverTab[61975]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:562
		_go_fuzz_dep_.CoverTab[61976]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:562
		// _ = "end of CoverTab[61976]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:562
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:562
	// _ = "end of CoverTab[61960]"
}

func (p *textParser) advance() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:565
	_go_fuzz_dep_.CoverTab[61977]++

												p.skipWhitespace()
												if p.done {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:568
		_go_fuzz_dep_.CoverTab[61980]++
													return
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:569
		// _ = "end of CoverTab[61980]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:570
		_go_fuzz_dep_.CoverTab[61981]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:570
		// _ = "end of CoverTab[61981]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:570
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:570
	// _ = "end of CoverTab[61977]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:570
	_go_fuzz_dep_.CoverTab[61978]++

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:573
	p.cur.err = nil
	p.cur.offset, p.cur.line = p.offset, p.line
	p.cur.unquoted = ""
	switch p.s[0] {
	case '<', '>', '{', '}', ':', '[', ']', ';', ',', '/':
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:577
		_go_fuzz_dep_.CoverTab[61982]++

													p.cur.value, p.s = p.s[0:1], p.s[1:len(p.s)]
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:579
		// _ = "end of CoverTab[61982]"
	case '"', '\'':
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:580
		_go_fuzz_dep_.CoverTab[61983]++

													i := 1
													for i < len(p.s) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:583
			_go_fuzz_dep_.CoverTab[61990]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:583
			return p.s[i] != p.s[0]
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:583
			// _ = "end of CoverTab[61990]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:583
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:583
			_go_fuzz_dep_.CoverTab[61991]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:583
			return p.s[i] != '\n'
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:583
			// _ = "end of CoverTab[61991]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:583
		}() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:583
			_go_fuzz_dep_.CoverTab[61992]++
														if p.s[i] == '\\' && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:584
				_go_fuzz_dep_.CoverTab[61994]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:584
				return i+1 < len(p.s)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:584
				// _ = "end of CoverTab[61994]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:584
			}() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:584
				_go_fuzz_dep_.CoverTab[61995]++

															i++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:586
				// _ = "end of CoverTab[61995]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:587
				_go_fuzz_dep_.CoverTab[61996]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:587
				// _ = "end of CoverTab[61996]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:587
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:587
			// _ = "end of CoverTab[61992]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:587
			_go_fuzz_dep_.CoverTab[61993]++
														i++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:588
			// _ = "end of CoverTab[61993]"
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:589
		// _ = "end of CoverTab[61983]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:589
		_go_fuzz_dep_.CoverTab[61984]++
													if i >= len(p.s) || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:590
			_go_fuzz_dep_.CoverTab[61997]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:590
			return p.s[i] != p.s[0]
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:590
			// _ = "end of CoverTab[61997]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:590
		}() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:590
			_go_fuzz_dep_.CoverTab[61998]++
														p.errorf("unmatched quote")
														return
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:592
			// _ = "end of CoverTab[61998]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:593
			_go_fuzz_dep_.CoverTab[61999]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:593
			// _ = "end of CoverTab[61999]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:593
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:593
		// _ = "end of CoverTab[61984]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:593
		_go_fuzz_dep_.CoverTab[61985]++
													unq, err := unquoteC(p.s[1:i], rune(p.s[0]))
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:595
			_go_fuzz_dep_.CoverTab[62000]++
														p.errorf("invalid quoted string %s: %v", p.s[0:i+1], err)
														return
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:597
			// _ = "end of CoverTab[62000]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:598
			_go_fuzz_dep_.CoverTab[62001]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:598
			// _ = "end of CoverTab[62001]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:598
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:598
		// _ = "end of CoverTab[61985]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:598
		_go_fuzz_dep_.CoverTab[61986]++
													p.cur.value, p.s = p.s[0:i+1], p.s[i+1:len(p.s)]
													p.cur.unquoted = unq
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:600
		// _ = "end of CoverTab[61986]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:601
		_go_fuzz_dep_.CoverTab[61987]++
													i := 0
													for i < len(p.s) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:603
			_go_fuzz_dep_.CoverTab[62002]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:603
			return isIdentOrNumberChar(p.s[i])
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:603
			// _ = "end of CoverTab[62002]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:603
		}() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:603
			_go_fuzz_dep_.CoverTab[62003]++
														i++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:604
			// _ = "end of CoverTab[62003]"
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:605
		// _ = "end of CoverTab[61987]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:605
		_go_fuzz_dep_.CoverTab[61988]++
													if i == 0 {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:606
			_go_fuzz_dep_.CoverTab[62004]++
														p.errorf("unexpected byte %#x", p.s[0])
														return
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:608
			// _ = "end of CoverTab[62004]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:609
			_go_fuzz_dep_.CoverTab[62005]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:609
			// _ = "end of CoverTab[62005]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:609
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:609
		// _ = "end of CoverTab[61988]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:609
		_go_fuzz_dep_.CoverTab[61989]++
													p.cur.value, p.s = p.s[0:i], p.s[i:len(p.s)]
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:610
		// _ = "end of CoverTab[61989]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:611
	// _ = "end of CoverTab[61978]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:611
	_go_fuzz_dep_.CoverTab[61979]++
												p.offset += len(p.cur.value)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:612
	// _ = "end of CoverTab[61979]"
}

// Back off the parser by one token. Can only be done between calls to next().
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:615
// It makes the next advance() a no-op.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:617
func (p *textParser) back() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:617
	_go_fuzz_dep_.CoverTab[62006]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:617
	p.backed = true
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:617
	// _ = "end of CoverTab[62006]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:617
}

// Advances the parser and returns the new current token.
func (p *textParser) next() *token {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:620
	_go_fuzz_dep_.CoverTab[62007]++
												if p.backed || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:621
		_go_fuzz_dep_.CoverTab[62010]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:621
		return p.done
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:621
		// _ = "end of CoverTab[62010]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:621
	}() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:621
		_go_fuzz_dep_.CoverTab[62011]++
													p.backed = false
													return &p.cur
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:623
		// _ = "end of CoverTab[62011]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:624
		_go_fuzz_dep_.CoverTab[62012]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:624
		// _ = "end of CoverTab[62012]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:624
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:624
	// _ = "end of CoverTab[62007]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:624
	_go_fuzz_dep_.CoverTab[62008]++
												p.advance()
												if p.done {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:626
		_go_fuzz_dep_.CoverTab[62013]++
													p.cur.value = ""
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:627
		// _ = "end of CoverTab[62013]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:628
		_go_fuzz_dep_.CoverTab[62014]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:628
		if len(p.cur.value) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:628
			_go_fuzz_dep_.CoverTab[62015]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:628
			return isQuote(p.cur.value[0])
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:628
			// _ = "end of CoverTab[62015]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:628
		}() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:628
			_go_fuzz_dep_.CoverTab[62016]++

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:631
			cat := p.cur
			for {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:632
				_go_fuzz_dep_.CoverTab[62018]++
															p.skipWhitespace()
															if p.done || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:634
					_go_fuzz_dep_.CoverTab[62021]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:634
					return !isQuote(p.s[0])
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:634
					// _ = "end of CoverTab[62021]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:634
				}() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:634
					_go_fuzz_dep_.CoverTab[62022]++
																break
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:635
					// _ = "end of CoverTab[62022]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:636
					_go_fuzz_dep_.CoverTab[62023]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:636
					// _ = "end of CoverTab[62023]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:636
				}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:636
				// _ = "end of CoverTab[62018]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:636
				_go_fuzz_dep_.CoverTab[62019]++
															p.advance()
															if p.cur.err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:638
					_go_fuzz_dep_.CoverTab[62024]++
																return &p.cur
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:639
					// _ = "end of CoverTab[62024]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:640
					_go_fuzz_dep_.CoverTab[62025]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:640
					// _ = "end of CoverTab[62025]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:640
				}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:640
				// _ = "end of CoverTab[62019]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:640
				_go_fuzz_dep_.CoverTab[62020]++
															cat.value += " " + p.cur.value
															cat.unquoted += p.cur.unquoted
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:642
				// _ = "end of CoverTab[62020]"
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:643
			// _ = "end of CoverTab[62016]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:643
			_go_fuzz_dep_.CoverTab[62017]++
														p.done = false
														p.cur = cat
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:645
			// _ = "end of CoverTab[62017]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:646
			_go_fuzz_dep_.CoverTab[62026]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:646
			// _ = "end of CoverTab[62026]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:646
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:646
		// _ = "end of CoverTab[62014]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:646
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:646
	// _ = "end of CoverTab[62008]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:646
	_go_fuzz_dep_.CoverTab[62009]++
												return &p.cur
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:647
	// _ = "end of CoverTab[62009]"
}

func (p *textParser) consumeToken(s string) error {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:650
	_go_fuzz_dep_.CoverTab[62027]++
												tok := p.next()
												if tok.err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:652
		_go_fuzz_dep_.CoverTab[62030]++
													return tok.err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:653
		// _ = "end of CoverTab[62030]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:654
		_go_fuzz_dep_.CoverTab[62031]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:654
		// _ = "end of CoverTab[62031]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:654
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:654
	// _ = "end of CoverTab[62027]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:654
	_go_fuzz_dep_.CoverTab[62028]++
												if tok.value != s {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:655
		_go_fuzz_dep_.CoverTab[62032]++
													p.back()
													return p.errorf("expected %q, found %q", s, tok.value)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:657
		// _ = "end of CoverTab[62032]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:658
		_go_fuzz_dep_.CoverTab[62033]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:658
		// _ = "end of CoverTab[62033]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:658
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:658
	// _ = "end of CoverTab[62028]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:658
	_go_fuzz_dep_.CoverTab[62029]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:659
	// _ = "end of CoverTab[62029]"
}

var errBadUTF8 = errors.New("proto: bad UTF-8")

func unquoteC(s string, quote rune) (string, error) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:664
	_go_fuzz_dep_.CoverTab[62034]++

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:670
	simple := true
	for _, r := range s {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:671
		_go_fuzz_dep_.CoverTab[62038]++
													if r == '\\' || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:672
			_go_fuzz_dep_.CoverTab[62039]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:672
			return r == quote
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:672
			// _ = "end of CoverTab[62039]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:672
		}() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:672
			_go_fuzz_dep_.CoverTab[62040]++
														simple = false
														break
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:674
			// _ = "end of CoverTab[62040]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:675
			_go_fuzz_dep_.CoverTab[62041]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:675
			// _ = "end of CoverTab[62041]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:675
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:675
		// _ = "end of CoverTab[62038]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:676
	// _ = "end of CoverTab[62034]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:676
	_go_fuzz_dep_.CoverTab[62035]++
												if simple {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:677
		_go_fuzz_dep_.CoverTab[62042]++
													return s, nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:678
		// _ = "end of CoverTab[62042]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:679
		_go_fuzz_dep_.CoverTab[62043]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:679
		// _ = "end of CoverTab[62043]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:679
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:679
	// _ = "end of CoverTab[62035]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:679
	_go_fuzz_dep_.CoverTab[62036]++

												buf := make([]byte, 0, 3*len(s)/2)
												for len(s) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:682
		_go_fuzz_dep_.CoverTab[62044]++
													r, n := utf8.DecodeRuneInString(s)
													if r == utf8.RuneError && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:684
			_go_fuzz_dep_.CoverTab[62048]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:684
			return n == 1
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:684
			// _ = "end of CoverTab[62048]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:684
		}() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:684
			_go_fuzz_dep_.CoverTab[62049]++
														return "", errBadUTF8
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:685
			// _ = "end of CoverTab[62049]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:686
			_go_fuzz_dep_.CoverTab[62050]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:686
			// _ = "end of CoverTab[62050]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:686
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:686
		// _ = "end of CoverTab[62044]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:686
		_go_fuzz_dep_.CoverTab[62045]++
													s = s[n:]
													if r != '\\' {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:688
			_go_fuzz_dep_.CoverTab[62051]++
														if r < utf8.RuneSelf {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:689
				_go_fuzz_dep_.CoverTab[62053]++
															buf = append(buf, byte(r))
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:690
				// _ = "end of CoverTab[62053]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:691
				_go_fuzz_dep_.CoverTab[62054]++
															buf = append(buf, string(r)...)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:692
				// _ = "end of CoverTab[62054]"
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:693
			// _ = "end of CoverTab[62051]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:693
			_go_fuzz_dep_.CoverTab[62052]++
														continue
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:694
			// _ = "end of CoverTab[62052]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:695
			_go_fuzz_dep_.CoverTab[62055]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:695
			// _ = "end of CoverTab[62055]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:695
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:695
		// _ = "end of CoverTab[62045]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:695
		_go_fuzz_dep_.CoverTab[62046]++

													ch, tail, err := unescape(s)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:698
			_go_fuzz_dep_.CoverTab[62056]++
														return "", err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:699
			// _ = "end of CoverTab[62056]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:700
			_go_fuzz_dep_.CoverTab[62057]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:700
			// _ = "end of CoverTab[62057]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:700
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:700
		// _ = "end of CoverTab[62046]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:700
		_go_fuzz_dep_.CoverTab[62047]++
													buf = append(buf, ch...)
													s = tail
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:702
		// _ = "end of CoverTab[62047]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:703
	// _ = "end of CoverTab[62036]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:703
	_go_fuzz_dep_.CoverTab[62037]++
												return string(buf), nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:704
	// _ = "end of CoverTab[62037]"
}

func unescape(s string) (ch string, tail string, err error) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:707
	_go_fuzz_dep_.CoverTab[62058]++
												r, n := utf8.DecodeRuneInString(s)
												if r == utf8.RuneError && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:709
		_go_fuzz_dep_.CoverTab[62061]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:709
		return n == 1
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:709
		// _ = "end of CoverTab[62061]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:709
	}() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:709
		_go_fuzz_dep_.CoverTab[62062]++
													return "", "", errBadUTF8
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:710
		// _ = "end of CoverTab[62062]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:711
		_go_fuzz_dep_.CoverTab[62063]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:711
		// _ = "end of CoverTab[62063]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:711
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:711
	// _ = "end of CoverTab[62058]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:711
	_go_fuzz_dep_.CoverTab[62059]++
												s = s[n:]
												switch r {
	case 'a':
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:714
		_go_fuzz_dep_.CoverTab[62064]++
													return "\a", s, nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:715
		// _ = "end of CoverTab[62064]"
	case 'b':
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:716
		_go_fuzz_dep_.CoverTab[62065]++
													return "\b", s, nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:717
		// _ = "end of CoverTab[62065]"
	case 'f':
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:718
		_go_fuzz_dep_.CoverTab[62066]++
													return "\f", s, nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:719
		// _ = "end of CoverTab[62066]"
	case 'n':
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:720
		_go_fuzz_dep_.CoverTab[62067]++
													return "\n", s, nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:721
		// _ = "end of CoverTab[62067]"
	case 'r':
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:722
		_go_fuzz_dep_.CoverTab[62068]++
													return "\r", s, nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:723
		// _ = "end of CoverTab[62068]"
	case 't':
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:724
		_go_fuzz_dep_.CoverTab[62069]++
													return "\t", s, nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:725
		// _ = "end of CoverTab[62069]"
	case 'v':
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:726
		_go_fuzz_dep_.CoverTab[62070]++
													return "\v", s, nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:727
		// _ = "end of CoverTab[62070]"
	case '?':
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:728
		_go_fuzz_dep_.CoverTab[62071]++
													return "?", s, nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:729
		// _ = "end of CoverTab[62071]"
	case '\'', '"', '\\':
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:730
		_go_fuzz_dep_.CoverTab[62072]++
													return string(r), s, nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:731
		// _ = "end of CoverTab[62072]"
	case '0', '1', '2', '3', '4', '5', '6', '7':
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:732
		_go_fuzz_dep_.CoverTab[62073]++
													if len(s) < 2 {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:733
			_go_fuzz_dep_.CoverTab[62083]++
														return "", "", fmt.Errorf(`\%c requires 2 following digits`, r)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:734
			// _ = "end of CoverTab[62083]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:735
			_go_fuzz_dep_.CoverTab[62084]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:735
			// _ = "end of CoverTab[62084]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:735
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:735
		// _ = "end of CoverTab[62073]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:735
		_go_fuzz_dep_.CoverTab[62074]++
													ss := string(r) + s[:2]
													s = s[2:]
													i, err := strconv.ParseUint(ss, 8, 8)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:739
			_go_fuzz_dep_.CoverTab[62085]++
														return "", "", fmt.Errorf(`\%s contains non-octal digits`, ss)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:740
			// _ = "end of CoverTab[62085]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:741
			_go_fuzz_dep_.CoverTab[62086]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:741
			// _ = "end of CoverTab[62086]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:741
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:741
		// _ = "end of CoverTab[62074]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:741
		_go_fuzz_dep_.CoverTab[62075]++
													return string([]byte{byte(i)}), s, nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:742
		// _ = "end of CoverTab[62075]"
	case 'x', 'X', 'u', 'U':
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:743
		_go_fuzz_dep_.CoverTab[62076]++
													var n int
													switch r {
		case 'x', 'X':
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:746
			_go_fuzz_dep_.CoverTab[62087]++
														n = 2
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:747
			// _ = "end of CoverTab[62087]"
		case 'u':
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:748
			_go_fuzz_dep_.CoverTab[62088]++
														n = 4
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:749
			// _ = "end of CoverTab[62088]"
		case 'U':
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:750
			_go_fuzz_dep_.CoverTab[62089]++
														n = 8
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:751
			// _ = "end of CoverTab[62089]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:751
		default:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:751
			_go_fuzz_dep_.CoverTab[62090]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:751
			// _ = "end of CoverTab[62090]"
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:752
		// _ = "end of CoverTab[62076]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:752
		_go_fuzz_dep_.CoverTab[62077]++
													if len(s) < n {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:753
			_go_fuzz_dep_.CoverTab[62091]++
														return "", "", fmt.Errorf(`\%c requires %d following digits`, r, n)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:754
			// _ = "end of CoverTab[62091]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:755
			_go_fuzz_dep_.CoverTab[62092]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:755
			// _ = "end of CoverTab[62092]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:755
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:755
		// _ = "end of CoverTab[62077]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:755
		_go_fuzz_dep_.CoverTab[62078]++
													ss := s[:n]
													s = s[n:]
													i, err := strconv.ParseUint(ss, 16, 64)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:759
			_go_fuzz_dep_.CoverTab[62093]++
														return "", "", fmt.Errorf(`\%c%s contains non-hexadecimal digits`, r, ss)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:760
			// _ = "end of CoverTab[62093]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:761
			_go_fuzz_dep_.CoverTab[62094]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:761
			// _ = "end of CoverTab[62094]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:761
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:761
		// _ = "end of CoverTab[62078]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:761
		_go_fuzz_dep_.CoverTab[62079]++
													if r == 'x' || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:762
			_go_fuzz_dep_.CoverTab[62095]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:762
			return r == 'X'
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:762
			// _ = "end of CoverTab[62095]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:762
		}() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:762
			_go_fuzz_dep_.CoverTab[62096]++
														return string([]byte{byte(i)}), s, nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:763
			// _ = "end of CoverTab[62096]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:764
			_go_fuzz_dep_.CoverTab[62097]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:764
			// _ = "end of CoverTab[62097]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:764
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:764
		// _ = "end of CoverTab[62079]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:764
		_go_fuzz_dep_.CoverTab[62080]++
													if i > utf8.MaxRune {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:765
			_go_fuzz_dep_.CoverTab[62098]++
														return "", "", fmt.Errorf(`\%c%s is not a valid Unicode code point`, r, ss)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:766
			// _ = "end of CoverTab[62098]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:767
			_go_fuzz_dep_.CoverTab[62099]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:767
			// _ = "end of CoverTab[62099]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:767
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:767
		// _ = "end of CoverTab[62080]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:767
		_go_fuzz_dep_.CoverTab[62081]++
													return string(rune(i)), s, nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:768
		// _ = "end of CoverTab[62081]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:768
	default:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:768
		_go_fuzz_dep_.CoverTab[62082]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:768
		// _ = "end of CoverTab[62082]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:769
	// _ = "end of CoverTab[62059]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:769
	_go_fuzz_dep_.CoverTab[62060]++
												return "", "", fmt.Errorf(`unknown escape \%c`, r)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:770
	// _ = "end of CoverTab[62060]"
}

func isIdentOrNumberChar(c byte) bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:773
	_go_fuzz_dep_.CoverTab[62100]++
												switch {
	case 'A' <= c && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:775
		_go_fuzz_dep_.CoverTab[62106]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:775
		return c <= 'Z'
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:775
		// _ = "end of CoverTab[62106]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:775
	}(), 'a' <= c && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:775
		_go_fuzz_dep_.CoverTab[62107]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:775
		return c <= 'z'
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:775
		// _ = "end of CoverTab[62107]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:775
	}():
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:775
		_go_fuzz_dep_.CoverTab[62103]++
													return true
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:776
		// _ = "end of CoverTab[62103]"
	case '0' <= c && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:777
		_go_fuzz_dep_.CoverTab[62108]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:777
		return c <= '9'
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:777
		// _ = "end of CoverTab[62108]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:777
	}():
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:777
		_go_fuzz_dep_.CoverTab[62104]++
													return true
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:778
		// _ = "end of CoverTab[62104]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:778
	default:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:778
		_go_fuzz_dep_.CoverTab[62105]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:778
		// _ = "end of CoverTab[62105]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:779
	// _ = "end of CoverTab[62100]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:779
	_go_fuzz_dep_.CoverTab[62101]++
												switch c {
	case '-', '+', '.', '_':
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:781
		_go_fuzz_dep_.CoverTab[62109]++
													return true
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:782
		// _ = "end of CoverTab[62109]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:782
	default:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:782
		_go_fuzz_dep_.CoverTab[62110]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:782
		// _ = "end of CoverTab[62110]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:783
	// _ = "end of CoverTab[62101]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:783
	_go_fuzz_dep_.CoverTab[62102]++
												return false
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:784
	// _ = "end of CoverTab[62102]"
}

func isWhitespace(c byte) bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:787
	_go_fuzz_dep_.CoverTab[62111]++
												switch c {
	case ' ', '\t', '\n', '\r':
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:789
		_go_fuzz_dep_.CoverTab[62113]++
													return true
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:790
		// _ = "end of CoverTab[62113]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:790
	default:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:790
		_go_fuzz_dep_.CoverTab[62114]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:790
		// _ = "end of CoverTab[62114]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:791
	// _ = "end of CoverTab[62111]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:791
	_go_fuzz_dep_.CoverTab[62112]++
												return false
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:792
	// _ = "end of CoverTab[62112]"
}

func isQuote(c byte) bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:795
	_go_fuzz_dep_.CoverTab[62115]++
												switch c {
	case '"', '\'':
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:797
		_go_fuzz_dep_.CoverTab[62117]++
													return true
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:798
		// _ = "end of CoverTab[62117]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:798
	default:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:798
		_go_fuzz_dep_.CoverTab[62118]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:798
		// _ = "end of CoverTab[62118]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:799
	// _ = "end of CoverTab[62115]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:799
	_go_fuzz_dep_.CoverTab[62116]++
												return false
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:800
	// _ = "end of CoverTab[62116]"
}

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:801
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/text_decode.go:801
var _ = _go_fuzz_dep_.CoverTab
