// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/encoding/json/encode.go:5
// Package json implements encoding and decoding of JSON as defined in
//line /usr/local/go/src/encoding/json/encode.go:5
// RFC 7159. The mapping between JSON and Go values is described
//line /usr/local/go/src/encoding/json/encode.go:5
// in the documentation for the Marshal and Unmarshal functions.
//line /usr/local/go/src/encoding/json/encode.go:5
//
//line /usr/local/go/src/encoding/json/encode.go:5
// See "JSON and Go" for an introduction to this package:
//line /usr/local/go/src/encoding/json/encode.go:5
// https://golang.org/doc/articles/json_and_go.html
//line /usr/local/go/src/encoding/json/encode.go:11
package json

//line /usr/local/go/src/encoding/json/encode.go:11
import (
//line /usr/local/go/src/encoding/json/encode.go:11
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/encoding/json/encode.go:11
)
//line /usr/local/go/src/encoding/json/encode.go:11
import (
//line /usr/local/go/src/encoding/json/encode.go:11
	_atomic_ "sync/atomic"
//line /usr/local/go/src/encoding/json/encode.go:11
)

import (
	"bytes"
	"encoding"
	"encoding/base64"
	"fmt"
	"math"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"sync"
	"unicode"
	"unicode/utf8"
)

// Marshal returns the JSON encoding of v.
//line /usr/local/go/src/encoding/json/encode.go:28
//
//line /usr/local/go/src/encoding/json/encode.go:28
// Marshal traverses the value v recursively.
//line /usr/local/go/src/encoding/json/encode.go:28
// If an encountered value implements the Marshaler interface
//line /usr/local/go/src/encoding/json/encode.go:28
// and is not a nil pointer, Marshal calls its MarshalJSON method
//line /usr/local/go/src/encoding/json/encode.go:28
// to produce JSON. If no MarshalJSON method is present but the
//line /usr/local/go/src/encoding/json/encode.go:28
// value implements encoding.TextMarshaler instead, Marshal calls
//line /usr/local/go/src/encoding/json/encode.go:28
// its MarshalText method and encodes the result as a JSON string.
//line /usr/local/go/src/encoding/json/encode.go:28
// The nil pointer exception is not strictly necessary
//line /usr/local/go/src/encoding/json/encode.go:28
// but mimics a similar, necessary exception in the behavior of
//line /usr/local/go/src/encoding/json/encode.go:28
// UnmarshalJSON.
//line /usr/local/go/src/encoding/json/encode.go:28
//
//line /usr/local/go/src/encoding/json/encode.go:28
// Otherwise, Marshal uses the following type-dependent default encodings:
//line /usr/local/go/src/encoding/json/encode.go:28
//
//line /usr/local/go/src/encoding/json/encode.go:28
// Boolean values encode as JSON booleans.
//line /usr/local/go/src/encoding/json/encode.go:28
//
//line /usr/local/go/src/encoding/json/encode.go:28
// Floating point, integer, and Number values encode as JSON numbers.
//line /usr/local/go/src/encoding/json/encode.go:28
//
//line /usr/local/go/src/encoding/json/encode.go:28
// String values encode as JSON strings coerced to valid UTF-8,
//line /usr/local/go/src/encoding/json/encode.go:28
// replacing invalid bytes with the Unicode replacement rune.
//line /usr/local/go/src/encoding/json/encode.go:28
// So that the JSON will be safe to embed inside HTML <script> tags,
//line /usr/local/go/src/encoding/json/encode.go:28
// the string is encoded using HTMLEscape,
//line /usr/local/go/src/encoding/json/encode.go:28
// which replaces "<", ">", "&", U+2028, and U+2029 are escaped
//line /usr/local/go/src/encoding/json/encode.go:28
// to "\u003c","\u003e", "\u0026", "\u2028", and "\u2029".
//line /usr/local/go/src/encoding/json/encode.go:28
// This replacement can be disabled when using an Encoder,
//line /usr/local/go/src/encoding/json/encode.go:28
// by calling SetEscapeHTML(false).
//line /usr/local/go/src/encoding/json/encode.go:28
//
//line /usr/local/go/src/encoding/json/encode.go:28
// Array and slice values encode as JSON arrays, except that
//line /usr/local/go/src/encoding/json/encode.go:28
// []byte encodes as a base64-encoded string, and a nil slice
//line /usr/local/go/src/encoding/json/encode.go:28
// encodes as the null JSON value.
//line /usr/local/go/src/encoding/json/encode.go:28
//
//line /usr/local/go/src/encoding/json/encode.go:28
// Struct values encode as JSON objects.
//line /usr/local/go/src/encoding/json/encode.go:28
// Each exported struct field becomes a member of the object, using the
//line /usr/local/go/src/encoding/json/encode.go:28
// field name as the object key, unless the field is omitted for one of the
//line /usr/local/go/src/encoding/json/encode.go:28
// reasons given below.
//line /usr/local/go/src/encoding/json/encode.go:28
//
//line /usr/local/go/src/encoding/json/encode.go:28
// The encoding of each struct field can be customized by the format string
//line /usr/local/go/src/encoding/json/encode.go:28
// stored under the "json" key in the struct field's tag.
//line /usr/local/go/src/encoding/json/encode.go:28
// The format string gives the name of the field, possibly followed by a
//line /usr/local/go/src/encoding/json/encode.go:28
// comma-separated list of options. The name may be empty in order to
//line /usr/local/go/src/encoding/json/encode.go:28
// specify options without overriding the default field name.
//line /usr/local/go/src/encoding/json/encode.go:28
//
//line /usr/local/go/src/encoding/json/encode.go:28
// The "omitempty" option specifies that the field should be omitted
//line /usr/local/go/src/encoding/json/encode.go:28
// from the encoding if the field has an empty value, defined as
//line /usr/local/go/src/encoding/json/encode.go:28
// false, 0, a nil pointer, a nil interface value, and any empty array,
//line /usr/local/go/src/encoding/json/encode.go:28
// slice, map, or string.
//line /usr/local/go/src/encoding/json/encode.go:28
//
//line /usr/local/go/src/encoding/json/encode.go:28
// As a special case, if the field tag is "-", the field is always omitted.
//line /usr/local/go/src/encoding/json/encode.go:28
// Note that a field with name "-" can still be generated using the tag "-,".
//line /usr/local/go/src/encoding/json/encode.go:28
//
//line /usr/local/go/src/encoding/json/encode.go:28
// Examples of struct field tags and their meanings:
//line /usr/local/go/src/encoding/json/encode.go:28
//
//line /usr/local/go/src/encoding/json/encode.go:28
//	// Field appears in JSON as key "myName".
//line /usr/local/go/src/encoding/json/encode.go:28
//	Field int `json:"myName"`
//line /usr/local/go/src/encoding/json/encode.go:28
//
//line /usr/local/go/src/encoding/json/encode.go:28
//	// Field appears in JSON as key "myName" and
//line /usr/local/go/src/encoding/json/encode.go:28
//	// the field is omitted from the object if its value is empty,
//line /usr/local/go/src/encoding/json/encode.go:28
//	// as defined above.
//line /usr/local/go/src/encoding/json/encode.go:28
//	Field int `json:"myName,omitempty"`
//line /usr/local/go/src/encoding/json/encode.go:28
//
//line /usr/local/go/src/encoding/json/encode.go:28
//	// Field appears in JSON as key "Field" (the default), but
//line /usr/local/go/src/encoding/json/encode.go:28
//	// the field is skipped if empty.
//line /usr/local/go/src/encoding/json/encode.go:28
//	// Note the leading comma.
//line /usr/local/go/src/encoding/json/encode.go:28
//	Field int `json:",omitempty"`
//line /usr/local/go/src/encoding/json/encode.go:28
//
//line /usr/local/go/src/encoding/json/encode.go:28
//	// Field is ignored by this package.
//line /usr/local/go/src/encoding/json/encode.go:28
//	Field int `json:"-"`
//line /usr/local/go/src/encoding/json/encode.go:28
//
//line /usr/local/go/src/encoding/json/encode.go:28
//	// Field appears in JSON as key "-".
//line /usr/local/go/src/encoding/json/encode.go:28
//	Field int `json:"-,"`
//line /usr/local/go/src/encoding/json/encode.go:28
//
//line /usr/local/go/src/encoding/json/encode.go:28
// The "string" option signals that a field is stored as JSON inside a
//line /usr/local/go/src/encoding/json/encode.go:28
// JSON-encoded string. It applies only to fields of string, floating point,
//line /usr/local/go/src/encoding/json/encode.go:28
// integer, or boolean types. This extra level of encoding is sometimes used
//line /usr/local/go/src/encoding/json/encode.go:28
// when communicating with JavaScript programs:
//line /usr/local/go/src/encoding/json/encode.go:28
//
//line /usr/local/go/src/encoding/json/encode.go:28
//	Int64String int64 `json:",string"`
//line /usr/local/go/src/encoding/json/encode.go:28
//
//line /usr/local/go/src/encoding/json/encode.go:28
// The key name will be used if it's a non-empty string consisting of
//line /usr/local/go/src/encoding/json/encode.go:28
// only Unicode letters, digits, and ASCII punctuation except quotation
//line /usr/local/go/src/encoding/json/encode.go:28
// marks, backslash, and comma.
//line /usr/local/go/src/encoding/json/encode.go:28
//
//line /usr/local/go/src/encoding/json/encode.go:28
// Anonymous struct fields are usually marshaled as if their inner exported fields
//line /usr/local/go/src/encoding/json/encode.go:28
// were fields in the outer struct, subject to the usual Go visibility rules amended
//line /usr/local/go/src/encoding/json/encode.go:28
// as described in the next paragraph.
//line /usr/local/go/src/encoding/json/encode.go:28
// An anonymous struct field with a name given in its JSON tag is treated as
//line /usr/local/go/src/encoding/json/encode.go:28
// having that name, rather than being anonymous.
//line /usr/local/go/src/encoding/json/encode.go:28
// An anonymous struct field of interface type is treated the same as having
//line /usr/local/go/src/encoding/json/encode.go:28
// that type as its name, rather than being anonymous.
//line /usr/local/go/src/encoding/json/encode.go:28
//
//line /usr/local/go/src/encoding/json/encode.go:28
// The Go visibility rules for struct fields are amended for JSON when
//line /usr/local/go/src/encoding/json/encode.go:28
// deciding which field to marshal or unmarshal. If there are
//line /usr/local/go/src/encoding/json/encode.go:28
// multiple fields at the same level, and that level is the least
//line /usr/local/go/src/encoding/json/encode.go:28
// nested (and would therefore be the nesting level selected by the
//line /usr/local/go/src/encoding/json/encode.go:28
// usual Go rules), the following extra rules apply:
//line /usr/local/go/src/encoding/json/encode.go:28
//
//line /usr/local/go/src/encoding/json/encode.go:28
// 1) Of those fields, if any are JSON-tagged, only tagged fields are considered,
//line /usr/local/go/src/encoding/json/encode.go:28
// even if there are multiple untagged fields that would otherwise conflict.
//line /usr/local/go/src/encoding/json/encode.go:28
//
//line /usr/local/go/src/encoding/json/encode.go:28
// 2) If there is exactly one field (tagged or not according to the first rule), that is selected.
//line /usr/local/go/src/encoding/json/encode.go:28
//
//line /usr/local/go/src/encoding/json/encode.go:28
// 3) Otherwise there are multiple fields, and all are ignored; no error occurs.
//line /usr/local/go/src/encoding/json/encode.go:28
//
//line /usr/local/go/src/encoding/json/encode.go:28
// Handling of anonymous struct fields is new in Go 1.1.
//line /usr/local/go/src/encoding/json/encode.go:28
// Prior to Go 1.1, anonymous struct fields were ignored. To force ignoring of
//line /usr/local/go/src/encoding/json/encode.go:28
// an anonymous struct field in both current and earlier versions, give the field
//line /usr/local/go/src/encoding/json/encode.go:28
// a JSON tag of "-".
//line /usr/local/go/src/encoding/json/encode.go:28
//
//line /usr/local/go/src/encoding/json/encode.go:28
// Map values encode as JSON objects. The map's key type must either be a
//line /usr/local/go/src/encoding/json/encode.go:28
// string, an integer type, or implement encoding.TextMarshaler. The map keys
//line /usr/local/go/src/encoding/json/encode.go:28
// are sorted and used as JSON object keys by applying the following rules,
//line /usr/local/go/src/encoding/json/encode.go:28
// subject to the UTF-8 coercion described for string values above:
//line /usr/local/go/src/encoding/json/encode.go:28
//   - keys of any string type are used directly
//line /usr/local/go/src/encoding/json/encode.go:28
//   - encoding.TextMarshalers are marshaled
//line /usr/local/go/src/encoding/json/encode.go:28
//   - integer keys are converted to strings
//line /usr/local/go/src/encoding/json/encode.go:28
//
//line /usr/local/go/src/encoding/json/encode.go:28
// Pointer values encode as the value pointed to.
//line /usr/local/go/src/encoding/json/encode.go:28
// A nil pointer encodes as the null JSON value.
//line /usr/local/go/src/encoding/json/encode.go:28
//
//line /usr/local/go/src/encoding/json/encode.go:28
// Interface values encode as the value contained in the interface.
//line /usr/local/go/src/encoding/json/encode.go:28
// A nil interface value encodes as the null JSON value.
//line /usr/local/go/src/encoding/json/encode.go:28
//
//line /usr/local/go/src/encoding/json/encode.go:28
// Channel, complex, and function values cannot be encoded in JSON.
//line /usr/local/go/src/encoding/json/encode.go:28
// Attempting to encode such a value causes Marshal to return
//line /usr/local/go/src/encoding/json/encode.go:28
// an UnsupportedTypeError.
//line /usr/local/go/src/encoding/json/encode.go:28
//
//line /usr/local/go/src/encoding/json/encode.go:28
// JSON cannot represent cyclic data structures and Marshal does not
//line /usr/local/go/src/encoding/json/encode.go:28
// handle them. Passing cyclic structures to Marshal will result in
//line /usr/local/go/src/encoding/json/encode.go:28
// an error.
//line /usr/local/go/src/encoding/json/encode.go:157
func Marshal(v any) ([]byte, error) {
//line /usr/local/go/src/encoding/json/encode.go:157
	_go_fuzz_dep_.CoverTab[27413]++
							e := newEncodeState()
							defer encodeStatePool.Put(e)

							err := e.marshal(v, encOpts{escapeHTML: true})
							if err != nil {
//line /usr/local/go/src/encoding/json/encode.go:162
		_go_fuzz_dep_.CoverTab[27415]++
								return nil, err
//line /usr/local/go/src/encoding/json/encode.go:163
		// _ = "end of CoverTab[27415]"
	} else {
//line /usr/local/go/src/encoding/json/encode.go:164
		_go_fuzz_dep_.CoverTab[27416]++
//line /usr/local/go/src/encoding/json/encode.go:164
		// _ = "end of CoverTab[27416]"
//line /usr/local/go/src/encoding/json/encode.go:164
	}
//line /usr/local/go/src/encoding/json/encode.go:164
	// _ = "end of CoverTab[27413]"
//line /usr/local/go/src/encoding/json/encode.go:164
	_go_fuzz_dep_.CoverTab[27414]++
							buf := append([]byte(nil), e.Bytes()...)

							return buf, nil
//line /usr/local/go/src/encoding/json/encode.go:167
	// _ = "end of CoverTab[27414]"
}

// MarshalIndent is like Marshal but applies Indent to format the output.
//line /usr/local/go/src/encoding/json/encode.go:170
// Each JSON element in the output will begin on a new line beginning with prefix
//line /usr/local/go/src/encoding/json/encode.go:170
// followed by one or more copies of indent according to the indentation nesting.
//line /usr/local/go/src/encoding/json/encode.go:173
func MarshalIndent(v any, prefix, indent string) ([]byte, error) {
//line /usr/local/go/src/encoding/json/encode.go:173
	_go_fuzz_dep_.CoverTab[27417]++
							b, err := Marshal(v)
							if err != nil {
//line /usr/local/go/src/encoding/json/encode.go:175
		_go_fuzz_dep_.CoverTab[27420]++
								return nil, err
//line /usr/local/go/src/encoding/json/encode.go:176
		// _ = "end of CoverTab[27420]"
	} else {
//line /usr/local/go/src/encoding/json/encode.go:177
		_go_fuzz_dep_.CoverTab[27421]++
//line /usr/local/go/src/encoding/json/encode.go:177
		// _ = "end of CoverTab[27421]"
//line /usr/local/go/src/encoding/json/encode.go:177
	}
//line /usr/local/go/src/encoding/json/encode.go:177
	// _ = "end of CoverTab[27417]"
//line /usr/local/go/src/encoding/json/encode.go:177
	_go_fuzz_dep_.CoverTab[27418]++
							var buf bytes.Buffer
							err = Indent(&buf, b, prefix, indent)
							if err != nil {
//line /usr/local/go/src/encoding/json/encode.go:180
		_go_fuzz_dep_.CoverTab[27422]++
								return nil, err
//line /usr/local/go/src/encoding/json/encode.go:181
		// _ = "end of CoverTab[27422]"
	} else {
//line /usr/local/go/src/encoding/json/encode.go:182
		_go_fuzz_dep_.CoverTab[27423]++
//line /usr/local/go/src/encoding/json/encode.go:182
		// _ = "end of CoverTab[27423]"
//line /usr/local/go/src/encoding/json/encode.go:182
	}
//line /usr/local/go/src/encoding/json/encode.go:182
	// _ = "end of CoverTab[27418]"
//line /usr/local/go/src/encoding/json/encode.go:182
	_go_fuzz_dep_.CoverTab[27419]++
							return buf.Bytes(), nil
//line /usr/local/go/src/encoding/json/encode.go:183
	// _ = "end of CoverTab[27419]"
}

// HTMLEscape appends to dst the JSON-encoded src with <, >, &, U+2028 and U+2029
//line /usr/local/go/src/encoding/json/encode.go:186
// characters inside string literals changed to \u003c, \u003e, \u0026, \u2028, \u2029
//line /usr/local/go/src/encoding/json/encode.go:186
// so that the JSON will be safe to embed inside HTML <script> tags.
//line /usr/local/go/src/encoding/json/encode.go:186
// For historical reasons, web browsers don't honor standard HTML
//line /usr/local/go/src/encoding/json/encode.go:186
// escaping within <script> tags, so an alternative JSON encoding must
//line /usr/local/go/src/encoding/json/encode.go:186
// be used.
//line /usr/local/go/src/encoding/json/encode.go:192
func HTMLEscape(dst *bytes.Buffer, src []byte) {
//line /usr/local/go/src/encoding/json/encode.go:192
	_go_fuzz_dep_.CoverTab[27424]++

//line /usr/local/go/src/encoding/json/encode.go:195
	start := 0
	for i, c := range src {
//line /usr/local/go/src/encoding/json/encode.go:196
		_go_fuzz_dep_.CoverTab[27426]++
								if c == '<' || func() bool {
//line /usr/local/go/src/encoding/json/encode.go:197
			_go_fuzz_dep_.CoverTab[27428]++
//line /usr/local/go/src/encoding/json/encode.go:197
			return c == '>'
//line /usr/local/go/src/encoding/json/encode.go:197
			// _ = "end of CoverTab[27428]"
//line /usr/local/go/src/encoding/json/encode.go:197
		}() || func() bool {
//line /usr/local/go/src/encoding/json/encode.go:197
			_go_fuzz_dep_.CoverTab[27429]++
//line /usr/local/go/src/encoding/json/encode.go:197
			return c == '&'
//line /usr/local/go/src/encoding/json/encode.go:197
			// _ = "end of CoverTab[27429]"
//line /usr/local/go/src/encoding/json/encode.go:197
		}() {
//line /usr/local/go/src/encoding/json/encode.go:197
			_go_fuzz_dep_.CoverTab[27430]++
									if start < i {
//line /usr/local/go/src/encoding/json/encode.go:198
				_go_fuzz_dep_.CoverTab[27432]++
										dst.Write(src[start:i])
//line /usr/local/go/src/encoding/json/encode.go:199
				// _ = "end of CoverTab[27432]"
			} else {
//line /usr/local/go/src/encoding/json/encode.go:200
				_go_fuzz_dep_.CoverTab[27433]++
//line /usr/local/go/src/encoding/json/encode.go:200
				// _ = "end of CoverTab[27433]"
//line /usr/local/go/src/encoding/json/encode.go:200
			}
//line /usr/local/go/src/encoding/json/encode.go:200
			// _ = "end of CoverTab[27430]"
//line /usr/local/go/src/encoding/json/encode.go:200
			_go_fuzz_dep_.CoverTab[27431]++
									dst.WriteString(`\u00`)
									dst.WriteByte(hex[c>>4])
									dst.WriteByte(hex[c&0xF])
									start = i + 1
//line /usr/local/go/src/encoding/json/encode.go:204
			// _ = "end of CoverTab[27431]"
		} else {
//line /usr/local/go/src/encoding/json/encode.go:205
			_go_fuzz_dep_.CoverTab[27434]++
//line /usr/local/go/src/encoding/json/encode.go:205
			// _ = "end of CoverTab[27434]"
//line /usr/local/go/src/encoding/json/encode.go:205
		}
//line /usr/local/go/src/encoding/json/encode.go:205
		// _ = "end of CoverTab[27426]"
//line /usr/local/go/src/encoding/json/encode.go:205
		_go_fuzz_dep_.CoverTab[27427]++

								if c == 0xE2 && func() bool {
//line /usr/local/go/src/encoding/json/encode.go:207
			_go_fuzz_dep_.CoverTab[27435]++
//line /usr/local/go/src/encoding/json/encode.go:207
			return i+2 < len(src)
//line /usr/local/go/src/encoding/json/encode.go:207
			// _ = "end of CoverTab[27435]"
//line /usr/local/go/src/encoding/json/encode.go:207
		}() && func() bool {
//line /usr/local/go/src/encoding/json/encode.go:207
			_go_fuzz_dep_.CoverTab[27436]++
//line /usr/local/go/src/encoding/json/encode.go:207
			return src[i+1] == 0x80
//line /usr/local/go/src/encoding/json/encode.go:207
			// _ = "end of CoverTab[27436]"
//line /usr/local/go/src/encoding/json/encode.go:207
		}() && func() bool {
//line /usr/local/go/src/encoding/json/encode.go:207
			_go_fuzz_dep_.CoverTab[27437]++
//line /usr/local/go/src/encoding/json/encode.go:207
			return src[i+2]&^1 == 0xA8
//line /usr/local/go/src/encoding/json/encode.go:207
			// _ = "end of CoverTab[27437]"
//line /usr/local/go/src/encoding/json/encode.go:207
		}() {
//line /usr/local/go/src/encoding/json/encode.go:207
			_go_fuzz_dep_.CoverTab[27438]++
									if start < i {
//line /usr/local/go/src/encoding/json/encode.go:208
				_go_fuzz_dep_.CoverTab[27440]++
										dst.Write(src[start:i])
//line /usr/local/go/src/encoding/json/encode.go:209
				// _ = "end of CoverTab[27440]"
			} else {
//line /usr/local/go/src/encoding/json/encode.go:210
				_go_fuzz_dep_.CoverTab[27441]++
//line /usr/local/go/src/encoding/json/encode.go:210
				// _ = "end of CoverTab[27441]"
//line /usr/local/go/src/encoding/json/encode.go:210
			}
//line /usr/local/go/src/encoding/json/encode.go:210
			// _ = "end of CoverTab[27438]"
//line /usr/local/go/src/encoding/json/encode.go:210
			_go_fuzz_dep_.CoverTab[27439]++
									dst.WriteString(`\u202`)
									dst.WriteByte(hex[src[i+2]&0xF])
									start = i + 3
//line /usr/local/go/src/encoding/json/encode.go:213
			// _ = "end of CoverTab[27439]"
		} else {
//line /usr/local/go/src/encoding/json/encode.go:214
			_go_fuzz_dep_.CoverTab[27442]++
//line /usr/local/go/src/encoding/json/encode.go:214
			// _ = "end of CoverTab[27442]"
//line /usr/local/go/src/encoding/json/encode.go:214
		}
//line /usr/local/go/src/encoding/json/encode.go:214
		// _ = "end of CoverTab[27427]"
	}
//line /usr/local/go/src/encoding/json/encode.go:215
	// _ = "end of CoverTab[27424]"
//line /usr/local/go/src/encoding/json/encode.go:215
	_go_fuzz_dep_.CoverTab[27425]++
							if start < len(src) {
//line /usr/local/go/src/encoding/json/encode.go:216
		_go_fuzz_dep_.CoverTab[27443]++
								dst.Write(src[start:])
//line /usr/local/go/src/encoding/json/encode.go:217
		// _ = "end of CoverTab[27443]"
	} else {
//line /usr/local/go/src/encoding/json/encode.go:218
		_go_fuzz_dep_.CoverTab[27444]++
//line /usr/local/go/src/encoding/json/encode.go:218
		// _ = "end of CoverTab[27444]"
//line /usr/local/go/src/encoding/json/encode.go:218
	}
//line /usr/local/go/src/encoding/json/encode.go:218
	// _ = "end of CoverTab[27425]"
}

// Marshaler is the interface implemented by types that
//line /usr/local/go/src/encoding/json/encode.go:221
// can marshal themselves into valid JSON.
//line /usr/local/go/src/encoding/json/encode.go:223
type Marshaler interface {
	MarshalJSON() ([]byte, error)
}

// An UnsupportedTypeError is returned by Marshal when attempting
//line /usr/local/go/src/encoding/json/encode.go:227
// to encode an unsupported value type.
//line /usr/local/go/src/encoding/json/encode.go:229
type UnsupportedTypeError struct {
	Type reflect.Type
}

func (e *UnsupportedTypeError) Error() string {
//line /usr/local/go/src/encoding/json/encode.go:233
	_go_fuzz_dep_.CoverTab[27445]++
							return "json: unsupported type: " + e.Type.String()
//line /usr/local/go/src/encoding/json/encode.go:234
	// _ = "end of CoverTab[27445]"
}

// An UnsupportedValueError is returned by Marshal when attempting
//line /usr/local/go/src/encoding/json/encode.go:237
// to encode an unsupported value.
//line /usr/local/go/src/encoding/json/encode.go:239
type UnsupportedValueError struct {
	Value	reflect.Value
	Str	string
}

func (e *UnsupportedValueError) Error() string {
//line /usr/local/go/src/encoding/json/encode.go:244
	_go_fuzz_dep_.CoverTab[27446]++
							return "json: unsupported value: " + e.Str
//line /usr/local/go/src/encoding/json/encode.go:245
	// _ = "end of CoverTab[27446]"
}

// Before Go 1.2, an InvalidUTF8Error was returned by Marshal when
//line /usr/local/go/src/encoding/json/encode.go:248
// attempting to encode a string value with invalid UTF-8 sequences.
//line /usr/local/go/src/encoding/json/encode.go:248
// As of Go 1.2, Marshal instead coerces the string to valid UTF-8 by
//line /usr/local/go/src/encoding/json/encode.go:248
// replacing invalid bytes with the Unicode replacement rune U+FFFD.
//line /usr/local/go/src/encoding/json/encode.go:248
//
//line /usr/local/go/src/encoding/json/encode.go:248
// Deprecated: No longer used; kept for compatibility.
//line /usr/local/go/src/encoding/json/encode.go:254
type InvalidUTF8Error struct {
	S string	// the whole string value that caused the error
}

func (e *InvalidUTF8Error) Error() string {
//line /usr/local/go/src/encoding/json/encode.go:258
	_go_fuzz_dep_.CoverTab[27447]++
							return "json: invalid UTF-8 in string: " + strconv.Quote(e.S)
//line /usr/local/go/src/encoding/json/encode.go:259
	// _ = "end of CoverTab[27447]"
}

// A MarshalerError represents an error from calling a MarshalJSON or MarshalText method.
type MarshalerError struct {
	Type		reflect.Type
	Err		error
	sourceFunc	string
}

func (e *MarshalerError) Error() string {
//line /usr/local/go/src/encoding/json/encode.go:269
	_go_fuzz_dep_.CoverTab[27448]++
							srcFunc := e.sourceFunc
							if srcFunc == "" {
//line /usr/local/go/src/encoding/json/encode.go:271
		_go_fuzz_dep_.CoverTab[27450]++
								srcFunc = "MarshalJSON"
//line /usr/local/go/src/encoding/json/encode.go:272
		// _ = "end of CoverTab[27450]"
	} else {
//line /usr/local/go/src/encoding/json/encode.go:273
		_go_fuzz_dep_.CoverTab[27451]++
//line /usr/local/go/src/encoding/json/encode.go:273
		// _ = "end of CoverTab[27451]"
//line /usr/local/go/src/encoding/json/encode.go:273
	}
//line /usr/local/go/src/encoding/json/encode.go:273
	// _ = "end of CoverTab[27448]"
//line /usr/local/go/src/encoding/json/encode.go:273
	_go_fuzz_dep_.CoverTab[27449]++
							return "json: error calling " + srcFunc +
		" for type " + e.Type.String() +
		": " + e.Err.Error()
//line /usr/local/go/src/encoding/json/encode.go:276
	// _ = "end of CoverTab[27449]"
}

// Unwrap returns the underlying error.
func (e *MarshalerError) Unwrap() error {
//line /usr/local/go/src/encoding/json/encode.go:280
	_go_fuzz_dep_.CoverTab[27452]++
//line /usr/local/go/src/encoding/json/encode.go:280
	return e.Err
//line /usr/local/go/src/encoding/json/encode.go:280
	// _ = "end of CoverTab[27452]"
//line /usr/local/go/src/encoding/json/encode.go:280
}

var hex = "0123456789abcdef"

// An encodeState encodes JSON into a bytes.Buffer.
type encodeState struct {
	bytes.Buffer	// accumulated output
	scratch		[64]byte

	// Keep track of what pointers we've seen in the current recursive call
	// path, to avoid cycles that could lead to a stack overflow. Only do
	// the relatively expensive map operations if ptrLevel is larger than
	// startDetectingCyclesAfter, so that we skip the work if we're within a
	// reasonable amount of nested pointers deep.
	ptrLevel	uint
	ptrSeen		map[any]struct{}
}

const startDetectingCyclesAfter = 1000

var encodeStatePool sync.Pool

func newEncodeState() *encodeState {
//line /usr/local/go/src/encoding/json/encode.go:302
	_go_fuzz_dep_.CoverTab[27453]++
							if v := encodeStatePool.Get(); v != nil {
//line /usr/local/go/src/encoding/json/encode.go:303
		_go_fuzz_dep_.CoverTab[27455]++
								e := v.(*encodeState)
								e.Reset()
								if len(e.ptrSeen) > 0 {
//line /usr/local/go/src/encoding/json/encode.go:306
			_go_fuzz_dep_.CoverTab[27457]++
									panic("ptrEncoder.encode should have emptied ptrSeen via defers")
//line /usr/local/go/src/encoding/json/encode.go:307
			// _ = "end of CoverTab[27457]"
		} else {
//line /usr/local/go/src/encoding/json/encode.go:308
			_go_fuzz_dep_.CoverTab[27458]++
//line /usr/local/go/src/encoding/json/encode.go:308
			// _ = "end of CoverTab[27458]"
//line /usr/local/go/src/encoding/json/encode.go:308
		}
//line /usr/local/go/src/encoding/json/encode.go:308
		// _ = "end of CoverTab[27455]"
//line /usr/local/go/src/encoding/json/encode.go:308
		_go_fuzz_dep_.CoverTab[27456]++
								e.ptrLevel = 0
								return e
//line /usr/local/go/src/encoding/json/encode.go:310
		// _ = "end of CoverTab[27456]"
	} else {
//line /usr/local/go/src/encoding/json/encode.go:311
		_go_fuzz_dep_.CoverTab[27459]++
//line /usr/local/go/src/encoding/json/encode.go:311
		// _ = "end of CoverTab[27459]"
//line /usr/local/go/src/encoding/json/encode.go:311
	}
//line /usr/local/go/src/encoding/json/encode.go:311
	// _ = "end of CoverTab[27453]"
//line /usr/local/go/src/encoding/json/encode.go:311
	_go_fuzz_dep_.CoverTab[27454]++
							return &encodeState{ptrSeen: make(map[any]struct{})}
//line /usr/local/go/src/encoding/json/encode.go:312
	// _ = "end of CoverTab[27454]"
}

// jsonError is an error wrapper type for internal use only.
//line /usr/local/go/src/encoding/json/encode.go:315
// Panics with errors are wrapped in jsonError so that the top-level recover
//line /usr/local/go/src/encoding/json/encode.go:315
// can distinguish intentional panics from this package.
//line /usr/local/go/src/encoding/json/encode.go:318
type jsonError struct{ error }

func (e *encodeState) marshal(v any, opts encOpts) (err error) {
//line /usr/local/go/src/encoding/json/encode.go:320
	_go_fuzz_dep_.CoverTab[27460]++
							defer func() {
//line /usr/local/go/src/encoding/json/encode.go:321
		_go_fuzz_dep_.CoverTab[27462]++
								if r := recover(); r != nil {
//line /usr/local/go/src/encoding/json/encode.go:322
			_go_fuzz_dep_.CoverTab[27463]++
									if je, ok := r.(jsonError); ok {
//line /usr/local/go/src/encoding/json/encode.go:323
				_go_fuzz_dep_.CoverTab[27464]++
										err = je.error
//line /usr/local/go/src/encoding/json/encode.go:324
				// _ = "end of CoverTab[27464]"
			} else {
//line /usr/local/go/src/encoding/json/encode.go:325
				_go_fuzz_dep_.CoverTab[27465]++
										panic(r)
//line /usr/local/go/src/encoding/json/encode.go:326
				// _ = "end of CoverTab[27465]"
			}
//line /usr/local/go/src/encoding/json/encode.go:327
			// _ = "end of CoverTab[27463]"
		} else {
//line /usr/local/go/src/encoding/json/encode.go:328
			_go_fuzz_dep_.CoverTab[27466]++
//line /usr/local/go/src/encoding/json/encode.go:328
			// _ = "end of CoverTab[27466]"
//line /usr/local/go/src/encoding/json/encode.go:328
		}
//line /usr/local/go/src/encoding/json/encode.go:328
		// _ = "end of CoverTab[27462]"
	}()
//line /usr/local/go/src/encoding/json/encode.go:329
	// _ = "end of CoverTab[27460]"
//line /usr/local/go/src/encoding/json/encode.go:329
	_go_fuzz_dep_.CoverTab[27461]++
							e.reflectValue(reflect.ValueOf(v), opts)
							return nil
//line /usr/local/go/src/encoding/json/encode.go:331
	// _ = "end of CoverTab[27461]"
}

// error aborts the encoding by panicking with err wrapped in jsonError.
func (e *encodeState) error(err error) {
//line /usr/local/go/src/encoding/json/encode.go:335
	_go_fuzz_dep_.CoverTab[27467]++
							panic(jsonError{err})
//line /usr/local/go/src/encoding/json/encode.go:336
	// _ = "end of CoverTab[27467]"
}

func isEmptyValue(v reflect.Value) bool {
//line /usr/local/go/src/encoding/json/encode.go:339
	_go_fuzz_dep_.CoverTab[27468]++
							switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
//line /usr/local/go/src/encoding/json/encode.go:341
		_go_fuzz_dep_.CoverTab[27470]++
								return v.Len() == 0
//line /usr/local/go/src/encoding/json/encode.go:342
		// _ = "end of CoverTab[27470]"
	case reflect.Bool:
//line /usr/local/go/src/encoding/json/encode.go:343
		_go_fuzz_dep_.CoverTab[27471]++
								return !v.Bool()
//line /usr/local/go/src/encoding/json/encode.go:344
		// _ = "end of CoverTab[27471]"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
//line /usr/local/go/src/encoding/json/encode.go:345
		_go_fuzz_dep_.CoverTab[27472]++
								return v.Int() == 0
//line /usr/local/go/src/encoding/json/encode.go:346
		// _ = "end of CoverTab[27472]"
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
//line /usr/local/go/src/encoding/json/encode.go:347
		_go_fuzz_dep_.CoverTab[27473]++
								return v.Uint() == 0
//line /usr/local/go/src/encoding/json/encode.go:348
		// _ = "end of CoverTab[27473]"
	case reflect.Float32, reflect.Float64:
//line /usr/local/go/src/encoding/json/encode.go:349
		_go_fuzz_dep_.CoverTab[27474]++
								return v.Float() == 0
//line /usr/local/go/src/encoding/json/encode.go:350
		// _ = "end of CoverTab[27474]"
	case reflect.Interface, reflect.Pointer:
//line /usr/local/go/src/encoding/json/encode.go:351
		_go_fuzz_dep_.CoverTab[27475]++
								return v.IsNil()
//line /usr/local/go/src/encoding/json/encode.go:352
		// _ = "end of CoverTab[27475]"
//line /usr/local/go/src/encoding/json/encode.go:352
	default:
//line /usr/local/go/src/encoding/json/encode.go:352
		_go_fuzz_dep_.CoverTab[27476]++
//line /usr/local/go/src/encoding/json/encode.go:352
		// _ = "end of CoverTab[27476]"
	}
//line /usr/local/go/src/encoding/json/encode.go:353
	// _ = "end of CoverTab[27468]"
//line /usr/local/go/src/encoding/json/encode.go:353
	_go_fuzz_dep_.CoverTab[27469]++
							return false
//line /usr/local/go/src/encoding/json/encode.go:354
	// _ = "end of CoverTab[27469]"
}

func (e *encodeState) reflectValue(v reflect.Value, opts encOpts) {
//line /usr/local/go/src/encoding/json/encode.go:357
	_go_fuzz_dep_.CoverTab[27477]++
							valueEncoder(v)(e, v, opts)
//line /usr/local/go/src/encoding/json/encode.go:358
	// _ = "end of CoverTab[27477]"
}

type encOpts struct {
	// quoted causes primitive fields to be encoded inside JSON strings.
	quoted	bool
	// escapeHTML causes '<', '>', and '&' to be escaped in JSON strings.
	escapeHTML	bool
}

type encoderFunc func(e *encodeState, v reflect.Value, opts encOpts)

var encoderCache sync.Map	// map[reflect.Type]encoderFunc

func valueEncoder(v reflect.Value) encoderFunc {
//line /usr/local/go/src/encoding/json/encode.go:372
	_go_fuzz_dep_.CoverTab[27478]++
							if !v.IsValid() {
//line /usr/local/go/src/encoding/json/encode.go:373
		_go_fuzz_dep_.CoverTab[27480]++
								return invalidValueEncoder
//line /usr/local/go/src/encoding/json/encode.go:374
		// _ = "end of CoverTab[27480]"
	} else {
//line /usr/local/go/src/encoding/json/encode.go:375
		_go_fuzz_dep_.CoverTab[27481]++
//line /usr/local/go/src/encoding/json/encode.go:375
		// _ = "end of CoverTab[27481]"
//line /usr/local/go/src/encoding/json/encode.go:375
	}
//line /usr/local/go/src/encoding/json/encode.go:375
	// _ = "end of CoverTab[27478]"
//line /usr/local/go/src/encoding/json/encode.go:375
	_go_fuzz_dep_.CoverTab[27479]++
							return typeEncoder(v.Type())
//line /usr/local/go/src/encoding/json/encode.go:376
	// _ = "end of CoverTab[27479]"
}

func typeEncoder(t reflect.Type) encoderFunc {
//line /usr/local/go/src/encoding/json/encode.go:379
	_go_fuzz_dep_.CoverTab[27482]++
							if fi, ok := encoderCache.Load(t); ok {
//line /usr/local/go/src/encoding/json/encode.go:380
		_go_fuzz_dep_.CoverTab[27486]++
								return fi.(encoderFunc)
//line /usr/local/go/src/encoding/json/encode.go:381
		// _ = "end of CoverTab[27486]"
	} else {
//line /usr/local/go/src/encoding/json/encode.go:382
		_go_fuzz_dep_.CoverTab[27487]++
//line /usr/local/go/src/encoding/json/encode.go:382
		// _ = "end of CoverTab[27487]"
//line /usr/local/go/src/encoding/json/encode.go:382
	}
//line /usr/local/go/src/encoding/json/encode.go:382
	// _ = "end of CoverTab[27482]"
//line /usr/local/go/src/encoding/json/encode.go:382
	_go_fuzz_dep_.CoverTab[27483]++

	// To deal with recursive types, populate the map with an
	// indirect func before we build it. This type waits on the
	// real func (f) to be ready and then calls it. This indirect
	// func is only used for recursive types.
	var (
		wg	sync.WaitGroup
		f	encoderFunc
	)
	wg.Add(1)
	fi, loaded := encoderCache.LoadOrStore(t, encoderFunc(func(e *encodeState, v reflect.Value, opts encOpts) {
//line /usr/local/go/src/encoding/json/encode.go:393
		_go_fuzz_dep_.CoverTab[27488]++
								wg.Wait()
								f(e, v, opts)
//line /usr/local/go/src/encoding/json/encode.go:395
		// _ = "end of CoverTab[27488]"
	}))
//line /usr/local/go/src/encoding/json/encode.go:396
	// _ = "end of CoverTab[27483]"
//line /usr/local/go/src/encoding/json/encode.go:396
	_go_fuzz_dep_.CoverTab[27484]++
							if loaded {
//line /usr/local/go/src/encoding/json/encode.go:397
		_go_fuzz_dep_.CoverTab[27489]++
								return fi.(encoderFunc)
//line /usr/local/go/src/encoding/json/encode.go:398
		// _ = "end of CoverTab[27489]"
	} else {
//line /usr/local/go/src/encoding/json/encode.go:399
		_go_fuzz_dep_.CoverTab[27490]++
//line /usr/local/go/src/encoding/json/encode.go:399
		// _ = "end of CoverTab[27490]"
//line /usr/local/go/src/encoding/json/encode.go:399
	}
//line /usr/local/go/src/encoding/json/encode.go:399
	// _ = "end of CoverTab[27484]"
//line /usr/local/go/src/encoding/json/encode.go:399
	_go_fuzz_dep_.CoverTab[27485]++

//line /usr/local/go/src/encoding/json/encode.go:402
	f = newTypeEncoder(t, true)
							wg.Done()
							encoderCache.Store(t, f)
							return f
//line /usr/local/go/src/encoding/json/encode.go:405
	// _ = "end of CoverTab[27485]"
}

var (
	marshalerType		= reflect.TypeOf((*Marshaler)(nil)).Elem()
	textMarshalerType	= reflect.TypeOf((*encoding.TextMarshaler)(nil)).Elem()
)

// newTypeEncoder constructs an encoderFunc for a type.
//line /usr/local/go/src/encoding/json/encode.go:413
// The returned encoder only checks CanAddr when allowAddr is true.
//line /usr/local/go/src/encoding/json/encode.go:415
func newTypeEncoder(t reflect.Type, allowAddr bool) encoderFunc {
//line /usr/local/go/src/encoding/json/encode.go:415
	_go_fuzz_dep_.CoverTab[27491]++

//line /usr/local/go/src/encoding/json/encode.go:420
	if t.Kind() != reflect.Pointer && func() bool {
//line /usr/local/go/src/encoding/json/encode.go:420
		_go_fuzz_dep_.CoverTab[27496]++
//line /usr/local/go/src/encoding/json/encode.go:420
		return allowAddr
//line /usr/local/go/src/encoding/json/encode.go:420
		// _ = "end of CoverTab[27496]"
//line /usr/local/go/src/encoding/json/encode.go:420
	}() && func() bool {
//line /usr/local/go/src/encoding/json/encode.go:420
		_go_fuzz_dep_.CoverTab[27497]++
//line /usr/local/go/src/encoding/json/encode.go:420
		return reflect.PointerTo(t).Implements(marshalerType)
//line /usr/local/go/src/encoding/json/encode.go:420
		// _ = "end of CoverTab[27497]"
//line /usr/local/go/src/encoding/json/encode.go:420
	}() {
//line /usr/local/go/src/encoding/json/encode.go:420
		_go_fuzz_dep_.CoverTab[27498]++
								return newCondAddrEncoder(addrMarshalerEncoder, newTypeEncoder(t, false))
//line /usr/local/go/src/encoding/json/encode.go:421
		// _ = "end of CoverTab[27498]"
	} else {
//line /usr/local/go/src/encoding/json/encode.go:422
		_go_fuzz_dep_.CoverTab[27499]++
//line /usr/local/go/src/encoding/json/encode.go:422
		// _ = "end of CoverTab[27499]"
//line /usr/local/go/src/encoding/json/encode.go:422
	}
//line /usr/local/go/src/encoding/json/encode.go:422
	// _ = "end of CoverTab[27491]"
//line /usr/local/go/src/encoding/json/encode.go:422
	_go_fuzz_dep_.CoverTab[27492]++
							if t.Implements(marshalerType) {
//line /usr/local/go/src/encoding/json/encode.go:423
		_go_fuzz_dep_.CoverTab[27500]++
								return marshalerEncoder
//line /usr/local/go/src/encoding/json/encode.go:424
		// _ = "end of CoverTab[27500]"
	} else {
//line /usr/local/go/src/encoding/json/encode.go:425
		_go_fuzz_dep_.CoverTab[27501]++
//line /usr/local/go/src/encoding/json/encode.go:425
		// _ = "end of CoverTab[27501]"
//line /usr/local/go/src/encoding/json/encode.go:425
	}
//line /usr/local/go/src/encoding/json/encode.go:425
	// _ = "end of CoverTab[27492]"
//line /usr/local/go/src/encoding/json/encode.go:425
	_go_fuzz_dep_.CoverTab[27493]++
							if t.Kind() != reflect.Pointer && func() bool {
//line /usr/local/go/src/encoding/json/encode.go:426
		_go_fuzz_dep_.CoverTab[27502]++
//line /usr/local/go/src/encoding/json/encode.go:426
		return allowAddr
//line /usr/local/go/src/encoding/json/encode.go:426
		// _ = "end of CoverTab[27502]"
//line /usr/local/go/src/encoding/json/encode.go:426
	}() && func() bool {
//line /usr/local/go/src/encoding/json/encode.go:426
		_go_fuzz_dep_.CoverTab[27503]++
//line /usr/local/go/src/encoding/json/encode.go:426
		return reflect.PointerTo(t).Implements(textMarshalerType)
//line /usr/local/go/src/encoding/json/encode.go:426
		// _ = "end of CoverTab[27503]"
//line /usr/local/go/src/encoding/json/encode.go:426
	}() {
//line /usr/local/go/src/encoding/json/encode.go:426
		_go_fuzz_dep_.CoverTab[27504]++
								return newCondAddrEncoder(addrTextMarshalerEncoder, newTypeEncoder(t, false))
//line /usr/local/go/src/encoding/json/encode.go:427
		// _ = "end of CoverTab[27504]"
	} else {
//line /usr/local/go/src/encoding/json/encode.go:428
		_go_fuzz_dep_.CoverTab[27505]++
//line /usr/local/go/src/encoding/json/encode.go:428
		// _ = "end of CoverTab[27505]"
//line /usr/local/go/src/encoding/json/encode.go:428
	}
//line /usr/local/go/src/encoding/json/encode.go:428
	// _ = "end of CoverTab[27493]"
//line /usr/local/go/src/encoding/json/encode.go:428
	_go_fuzz_dep_.CoverTab[27494]++
							if t.Implements(textMarshalerType) {
//line /usr/local/go/src/encoding/json/encode.go:429
		_go_fuzz_dep_.CoverTab[27506]++
								return textMarshalerEncoder
//line /usr/local/go/src/encoding/json/encode.go:430
		// _ = "end of CoverTab[27506]"
	} else {
//line /usr/local/go/src/encoding/json/encode.go:431
		_go_fuzz_dep_.CoverTab[27507]++
//line /usr/local/go/src/encoding/json/encode.go:431
		// _ = "end of CoverTab[27507]"
//line /usr/local/go/src/encoding/json/encode.go:431
	}
//line /usr/local/go/src/encoding/json/encode.go:431
	// _ = "end of CoverTab[27494]"
//line /usr/local/go/src/encoding/json/encode.go:431
	_go_fuzz_dep_.CoverTab[27495]++

							switch t.Kind() {
	case reflect.Bool:
//line /usr/local/go/src/encoding/json/encode.go:434
		_go_fuzz_dep_.CoverTab[27508]++
								return boolEncoder
//line /usr/local/go/src/encoding/json/encode.go:435
		// _ = "end of CoverTab[27508]"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
//line /usr/local/go/src/encoding/json/encode.go:436
		_go_fuzz_dep_.CoverTab[27509]++
								return intEncoder
//line /usr/local/go/src/encoding/json/encode.go:437
		// _ = "end of CoverTab[27509]"
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
//line /usr/local/go/src/encoding/json/encode.go:438
		_go_fuzz_dep_.CoverTab[27510]++
								return uintEncoder
//line /usr/local/go/src/encoding/json/encode.go:439
		// _ = "end of CoverTab[27510]"
	case reflect.Float32:
//line /usr/local/go/src/encoding/json/encode.go:440
		_go_fuzz_dep_.CoverTab[27511]++
								return float32Encoder
//line /usr/local/go/src/encoding/json/encode.go:441
		// _ = "end of CoverTab[27511]"
	case reflect.Float64:
//line /usr/local/go/src/encoding/json/encode.go:442
		_go_fuzz_dep_.CoverTab[27512]++
								return float64Encoder
//line /usr/local/go/src/encoding/json/encode.go:443
		// _ = "end of CoverTab[27512]"
	case reflect.String:
//line /usr/local/go/src/encoding/json/encode.go:444
		_go_fuzz_dep_.CoverTab[27513]++
								return stringEncoder
//line /usr/local/go/src/encoding/json/encode.go:445
		// _ = "end of CoverTab[27513]"
	case reflect.Interface:
//line /usr/local/go/src/encoding/json/encode.go:446
		_go_fuzz_dep_.CoverTab[27514]++
								return interfaceEncoder
//line /usr/local/go/src/encoding/json/encode.go:447
		// _ = "end of CoverTab[27514]"
	case reflect.Struct:
//line /usr/local/go/src/encoding/json/encode.go:448
		_go_fuzz_dep_.CoverTab[27515]++
								return newStructEncoder(t)
//line /usr/local/go/src/encoding/json/encode.go:449
		// _ = "end of CoverTab[27515]"
	case reflect.Map:
//line /usr/local/go/src/encoding/json/encode.go:450
		_go_fuzz_dep_.CoverTab[27516]++
								return newMapEncoder(t)
//line /usr/local/go/src/encoding/json/encode.go:451
		// _ = "end of CoverTab[27516]"
	case reflect.Slice:
//line /usr/local/go/src/encoding/json/encode.go:452
		_go_fuzz_dep_.CoverTab[27517]++
								return newSliceEncoder(t)
//line /usr/local/go/src/encoding/json/encode.go:453
		// _ = "end of CoverTab[27517]"
	case reflect.Array:
//line /usr/local/go/src/encoding/json/encode.go:454
		_go_fuzz_dep_.CoverTab[27518]++
								return newArrayEncoder(t)
//line /usr/local/go/src/encoding/json/encode.go:455
		// _ = "end of CoverTab[27518]"
	case reflect.Pointer:
//line /usr/local/go/src/encoding/json/encode.go:456
		_go_fuzz_dep_.CoverTab[27519]++
								return newPtrEncoder(t)
//line /usr/local/go/src/encoding/json/encode.go:457
		// _ = "end of CoverTab[27519]"
	default:
//line /usr/local/go/src/encoding/json/encode.go:458
		_go_fuzz_dep_.CoverTab[27520]++
								return unsupportedTypeEncoder
//line /usr/local/go/src/encoding/json/encode.go:459
		// _ = "end of CoverTab[27520]"
	}
//line /usr/local/go/src/encoding/json/encode.go:460
	// _ = "end of CoverTab[27495]"
}

func invalidValueEncoder(e *encodeState, v reflect.Value, _ encOpts) {
//line /usr/local/go/src/encoding/json/encode.go:463
	_go_fuzz_dep_.CoverTab[27521]++
							e.WriteString("null")
//line /usr/local/go/src/encoding/json/encode.go:464
	// _ = "end of CoverTab[27521]"
}

func marshalerEncoder(e *encodeState, v reflect.Value, opts encOpts) {
//line /usr/local/go/src/encoding/json/encode.go:467
	_go_fuzz_dep_.CoverTab[27522]++
							if v.Kind() == reflect.Pointer && func() bool {
//line /usr/local/go/src/encoding/json/encode.go:468
		_go_fuzz_dep_.CoverTab[27526]++
//line /usr/local/go/src/encoding/json/encode.go:468
		return v.IsNil()
//line /usr/local/go/src/encoding/json/encode.go:468
		// _ = "end of CoverTab[27526]"
//line /usr/local/go/src/encoding/json/encode.go:468
	}() {
//line /usr/local/go/src/encoding/json/encode.go:468
		_go_fuzz_dep_.CoverTab[27527]++
								e.WriteString("null")
								return
//line /usr/local/go/src/encoding/json/encode.go:470
		// _ = "end of CoverTab[27527]"
	} else {
//line /usr/local/go/src/encoding/json/encode.go:471
		_go_fuzz_dep_.CoverTab[27528]++
//line /usr/local/go/src/encoding/json/encode.go:471
		// _ = "end of CoverTab[27528]"
//line /usr/local/go/src/encoding/json/encode.go:471
	}
//line /usr/local/go/src/encoding/json/encode.go:471
	// _ = "end of CoverTab[27522]"
//line /usr/local/go/src/encoding/json/encode.go:471
	_go_fuzz_dep_.CoverTab[27523]++
							m, ok := v.Interface().(Marshaler)
							if !ok {
//line /usr/local/go/src/encoding/json/encode.go:473
		_go_fuzz_dep_.CoverTab[27529]++
								e.WriteString("null")
								return
//line /usr/local/go/src/encoding/json/encode.go:475
		// _ = "end of CoverTab[27529]"
	} else {
//line /usr/local/go/src/encoding/json/encode.go:476
		_go_fuzz_dep_.CoverTab[27530]++
//line /usr/local/go/src/encoding/json/encode.go:476
		// _ = "end of CoverTab[27530]"
//line /usr/local/go/src/encoding/json/encode.go:476
	}
//line /usr/local/go/src/encoding/json/encode.go:476
	// _ = "end of CoverTab[27523]"
//line /usr/local/go/src/encoding/json/encode.go:476
	_go_fuzz_dep_.CoverTab[27524]++
							b, err := m.MarshalJSON()
							if err == nil {
//line /usr/local/go/src/encoding/json/encode.go:478
		_go_fuzz_dep_.CoverTab[27531]++

								err = compact(&e.Buffer, b, opts.escapeHTML)
//line /usr/local/go/src/encoding/json/encode.go:480
		// _ = "end of CoverTab[27531]"
	} else {
//line /usr/local/go/src/encoding/json/encode.go:481
		_go_fuzz_dep_.CoverTab[27532]++
//line /usr/local/go/src/encoding/json/encode.go:481
		// _ = "end of CoverTab[27532]"
//line /usr/local/go/src/encoding/json/encode.go:481
	}
//line /usr/local/go/src/encoding/json/encode.go:481
	// _ = "end of CoverTab[27524]"
//line /usr/local/go/src/encoding/json/encode.go:481
	_go_fuzz_dep_.CoverTab[27525]++
							if err != nil {
//line /usr/local/go/src/encoding/json/encode.go:482
		_go_fuzz_dep_.CoverTab[27533]++
								e.error(&MarshalerError{v.Type(), err, "MarshalJSON"})
//line /usr/local/go/src/encoding/json/encode.go:483
		// _ = "end of CoverTab[27533]"
	} else {
//line /usr/local/go/src/encoding/json/encode.go:484
		_go_fuzz_dep_.CoverTab[27534]++
//line /usr/local/go/src/encoding/json/encode.go:484
		// _ = "end of CoverTab[27534]"
//line /usr/local/go/src/encoding/json/encode.go:484
	}
//line /usr/local/go/src/encoding/json/encode.go:484
	// _ = "end of CoverTab[27525]"
}

func addrMarshalerEncoder(e *encodeState, v reflect.Value, opts encOpts) {
//line /usr/local/go/src/encoding/json/encode.go:487
	_go_fuzz_dep_.CoverTab[27535]++
							va := v.Addr()
							if va.IsNil() {
//line /usr/local/go/src/encoding/json/encode.go:489
		_go_fuzz_dep_.CoverTab[27538]++
								e.WriteString("null")
								return
//line /usr/local/go/src/encoding/json/encode.go:491
		// _ = "end of CoverTab[27538]"
	} else {
//line /usr/local/go/src/encoding/json/encode.go:492
		_go_fuzz_dep_.CoverTab[27539]++
//line /usr/local/go/src/encoding/json/encode.go:492
		// _ = "end of CoverTab[27539]"
//line /usr/local/go/src/encoding/json/encode.go:492
	}
//line /usr/local/go/src/encoding/json/encode.go:492
	// _ = "end of CoverTab[27535]"
//line /usr/local/go/src/encoding/json/encode.go:492
	_go_fuzz_dep_.CoverTab[27536]++
							m := va.Interface().(Marshaler)
							b, err := m.MarshalJSON()
							if err == nil {
//line /usr/local/go/src/encoding/json/encode.go:495
		_go_fuzz_dep_.CoverTab[27540]++

								err = compact(&e.Buffer, b, opts.escapeHTML)
//line /usr/local/go/src/encoding/json/encode.go:497
		// _ = "end of CoverTab[27540]"
	} else {
//line /usr/local/go/src/encoding/json/encode.go:498
		_go_fuzz_dep_.CoverTab[27541]++
//line /usr/local/go/src/encoding/json/encode.go:498
		// _ = "end of CoverTab[27541]"
//line /usr/local/go/src/encoding/json/encode.go:498
	}
//line /usr/local/go/src/encoding/json/encode.go:498
	// _ = "end of CoverTab[27536]"
//line /usr/local/go/src/encoding/json/encode.go:498
	_go_fuzz_dep_.CoverTab[27537]++
							if err != nil {
//line /usr/local/go/src/encoding/json/encode.go:499
		_go_fuzz_dep_.CoverTab[27542]++
								e.error(&MarshalerError{v.Type(), err, "MarshalJSON"})
//line /usr/local/go/src/encoding/json/encode.go:500
		// _ = "end of CoverTab[27542]"
	} else {
//line /usr/local/go/src/encoding/json/encode.go:501
		_go_fuzz_dep_.CoverTab[27543]++
//line /usr/local/go/src/encoding/json/encode.go:501
		// _ = "end of CoverTab[27543]"
//line /usr/local/go/src/encoding/json/encode.go:501
	}
//line /usr/local/go/src/encoding/json/encode.go:501
	// _ = "end of CoverTab[27537]"
}

func textMarshalerEncoder(e *encodeState, v reflect.Value, opts encOpts) {
//line /usr/local/go/src/encoding/json/encode.go:504
	_go_fuzz_dep_.CoverTab[27544]++
							if v.Kind() == reflect.Pointer && func() bool {
//line /usr/local/go/src/encoding/json/encode.go:505
		_go_fuzz_dep_.CoverTab[27548]++
//line /usr/local/go/src/encoding/json/encode.go:505
		return v.IsNil()
//line /usr/local/go/src/encoding/json/encode.go:505
		// _ = "end of CoverTab[27548]"
//line /usr/local/go/src/encoding/json/encode.go:505
	}() {
//line /usr/local/go/src/encoding/json/encode.go:505
		_go_fuzz_dep_.CoverTab[27549]++
								e.WriteString("null")
								return
//line /usr/local/go/src/encoding/json/encode.go:507
		// _ = "end of CoverTab[27549]"
	} else {
//line /usr/local/go/src/encoding/json/encode.go:508
		_go_fuzz_dep_.CoverTab[27550]++
//line /usr/local/go/src/encoding/json/encode.go:508
		// _ = "end of CoverTab[27550]"
//line /usr/local/go/src/encoding/json/encode.go:508
	}
//line /usr/local/go/src/encoding/json/encode.go:508
	// _ = "end of CoverTab[27544]"
//line /usr/local/go/src/encoding/json/encode.go:508
	_go_fuzz_dep_.CoverTab[27545]++
							m, ok := v.Interface().(encoding.TextMarshaler)
							if !ok {
//line /usr/local/go/src/encoding/json/encode.go:510
		_go_fuzz_dep_.CoverTab[27551]++
								e.WriteString("null")
								return
//line /usr/local/go/src/encoding/json/encode.go:512
		// _ = "end of CoverTab[27551]"
	} else {
//line /usr/local/go/src/encoding/json/encode.go:513
		_go_fuzz_dep_.CoverTab[27552]++
//line /usr/local/go/src/encoding/json/encode.go:513
		// _ = "end of CoverTab[27552]"
//line /usr/local/go/src/encoding/json/encode.go:513
	}
//line /usr/local/go/src/encoding/json/encode.go:513
	// _ = "end of CoverTab[27545]"
//line /usr/local/go/src/encoding/json/encode.go:513
	_go_fuzz_dep_.CoverTab[27546]++
							b, err := m.MarshalText()
							if err != nil {
//line /usr/local/go/src/encoding/json/encode.go:515
		_go_fuzz_dep_.CoverTab[27553]++
								e.error(&MarshalerError{v.Type(), err, "MarshalText"})
//line /usr/local/go/src/encoding/json/encode.go:516
		// _ = "end of CoverTab[27553]"
	} else {
//line /usr/local/go/src/encoding/json/encode.go:517
		_go_fuzz_dep_.CoverTab[27554]++
//line /usr/local/go/src/encoding/json/encode.go:517
		// _ = "end of CoverTab[27554]"
//line /usr/local/go/src/encoding/json/encode.go:517
	}
//line /usr/local/go/src/encoding/json/encode.go:517
	// _ = "end of CoverTab[27546]"
//line /usr/local/go/src/encoding/json/encode.go:517
	_go_fuzz_dep_.CoverTab[27547]++
							e.stringBytes(b, opts.escapeHTML)
//line /usr/local/go/src/encoding/json/encode.go:518
	// _ = "end of CoverTab[27547]"
}

func addrTextMarshalerEncoder(e *encodeState, v reflect.Value, opts encOpts) {
//line /usr/local/go/src/encoding/json/encode.go:521
	_go_fuzz_dep_.CoverTab[27555]++
							va := v.Addr()
							if va.IsNil() {
//line /usr/local/go/src/encoding/json/encode.go:523
		_go_fuzz_dep_.CoverTab[27558]++
								e.WriteString("null")
								return
//line /usr/local/go/src/encoding/json/encode.go:525
		// _ = "end of CoverTab[27558]"
	} else {
//line /usr/local/go/src/encoding/json/encode.go:526
		_go_fuzz_dep_.CoverTab[27559]++
//line /usr/local/go/src/encoding/json/encode.go:526
		// _ = "end of CoverTab[27559]"
//line /usr/local/go/src/encoding/json/encode.go:526
	}
//line /usr/local/go/src/encoding/json/encode.go:526
	// _ = "end of CoverTab[27555]"
//line /usr/local/go/src/encoding/json/encode.go:526
	_go_fuzz_dep_.CoverTab[27556]++
							m := va.Interface().(encoding.TextMarshaler)
							b, err := m.MarshalText()
							if err != nil {
//line /usr/local/go/src/encoding/json/encode.go:529
		_go_fuzz_dep_.CoverTab[27560]++
								e.error(&MarshalerError{v.Type(), err, "MarshalText"})
//line /usr/local/go/src/encoding/json/encode.go:530
		// _ = "end of CoverTab[27560]"
	} else {
//line /usr/local/go/src/encoding/json/encode.go:531
		_go_fuzz_dep_.CoverTab[27561]++
//line /usr/local/go/src/encoding/json/encode.go:531
		// _ = "end of CoverTab[27561]"
//line /usr/local/go/src/encoding/json/encode.go:531
	}
//line /usr/local/go/src/encoding/json/encode.go:531
	// _ = "end of CoverTab[27556]"
//line /usr/local/go/src/encoding/json/encode.go:531
	_go_fuzz_dep_.CoverTab[27557]++
							e.stringBytes(b, opts.escapeHTML)
//line /usr/local/go/src/encoding/json/encode.go:532
	// _ = "end of CoverTab[27557]"
}

func boolEncoder(e *encodeState, v reflect.Value, opts encOpts) {
//line /usr/local/go/src/encoding/json/encode.go:535
	_go_fuzz_dep_.CoverTab[27562]++
							if opts.quoted {
//line /usr/local/go/src/encoding/json/encode.go:536
		_go_fuzz_dep_.CoverTab[27565]++
								e.WriteByte('"')
//line /usr/local/go/src/encoding/json/encode.go:537
		// _ = "end of CoverTab[27565]"
	} else {
//line /usr/local/go/src/encoding/json/encode.go:538
		_go_fuzz_dep_.CoverTab[27566]++
//line /usr/local/go/src/encoding/json/encode.go:538
		// _ = "end of CoverTab[27566]"
//line /usr/local/go/src/encoding/json/encode.go:538
	}
//line /usr/local/go/src/encoding/json/encode.go:538
	// _ = "end of CoverTab[27562]"
//line /usr/local/go/src/encoding/json/encode.go:538
	_go_fuzz_dep_.CoverTab[27563]++
							if v.Bool() {
//line /usr/local/go/src/encoding/json/encode.go:539
		_go_fuzz_dep_.CoverTab[27567]++
								e.WriteString("true")
//line /usr/local/go/src/encoding/json/encode.go:540
		// _ = "end of CoverTab[27567]"
	} else {
//line /usr/local/go/src/encoding/json/encode.go:541
		_go_fuzz_dep_.CoverTab[27568]++
								e.WriteString("false")
//line /usr/local/go/src/encoding/json/encode.go:542
		// _ = "end of CoverTab[27568]"
	}
//line /usr/local/go/src/encoding/json/encode.go:543
	// _ = "end of CoverTab[27563]"
//line /usr/local/go/src/encoding/json/encode.go:543
	_go_fuzz_dep_.CoverTab[27564]++
							if opts.quoted {
//line /usr/local/go/src/encoding/json/encode.go:544
		_go_fuzz_dep_.CoverTab[27569]++
								e.WriteByte('"')
//line /usr/local/go/src/encoding/json/encode.go:545
		// _ = "end of CoverTab[27569]"
	} else {
//line /usr/local/go/src/encoding/json/encode.go:546
		_go_fuzz_dep_.CoverTab[27570]++
//line /usr/local/go/src/encoding/json/encode.go:546
		// _ = "end of CoverTab[27570]"
//line /usr/local/go/src/encoding/json/encode.go:546
	}
//line /usr/local/go/src/encoding/json/encode.go:546
	// _ = "end of CoverTab[27564]"
}

func intEncoder(e *encodeState, v reflect.Value, opts encOpts) {
//line /usr/local/go/src/encoding/json/encode.go:549
	_go_fuzz_dep_.CoverTab[27571]++
							b := strconv.AppendInt(e.scratch[:0], v.Int(), 10)
							if opts.quoted {
//line /usr/local/go/src/encoding/json/encode.go:551
		_go_fuzz_dep_.CoverTab[27573]++
								e.WriteByte('"')
//line /usr/local/go/src/encoding/json/encode.go:552
		// _ = "end of CoverTab[27573]"
	} else {
//line /usr/local/go/src/encoding/json/encode.go:553
		_go_fuzz_dep_.CoverTab[27574]++
//line /usr/local/go/src/encoding/json/encode.go:553
		// _ = "end of CoverTab[27574]"
//line /usr/local/go/src/encoding/json/encode.go:553
	}
//line /usr/local/go/src/encoding/json/encode.go:553
	// _ = "end of CoverTab[27571]"
//line /usr/local/go/src/encoding/json/encode.go:553
	_go_fuzz_dep_.CoverTab[27572]++
							e.Write(b)
							if opts.quoted {
//line /usr/local/go/src/encoding/json/encode.go:555
		_go_fuzz_dep_.CoverTab[27575]++
								e.WriteByte('"')
//line /usr/local/go/src/encoding/json/encode.go:556
		// _ = "end of CoverTab[27575]"
	} else {
//line /usr/local/go/src/encoding/json/encode.go:557
		_go_fuzz_dep_.CoverTab[27576]++
//line /usr/local/go/src/encoding/json/encode.go:557
		// _ = "end of CoverTab[27576]"
//line /usr/local/go/src/encoding/json/encode.go:557
	}
//line /usr/local/go/src/encoding/json/encode.go:557
	// _ = "end of CoverTab[27572]"
}

func uintEncoder(e *encodeState, v reflect.Value, opts encOpts) {
//line /usr/local/go/src/encoding/json/encode.go:560
	_go_fuzz_dep_.CoverTab[27577]++
							b := strconv.AppendUint(e.scratch[:0], v.Uint(), 10)
							if opts.quoted {
//line /usr/local/go/src/encoding/json/encode.go:562
		_go_fuzz_dep_.CoverTab[27579]++
								e.WriteByte('"')
//line /usr/local/go/src/encoding/json/encode.go:563
		// _ = "end of CoverTab[27579]"
	} else {
//line /usr/local/go/src/encoding/json/encode.go:564
		_go_fuzz_dep_.CoverTab[27580]++
//line /usr/local/go/src/encoding/json/encode.go:564
		// _ = "end of CoverTab[27580]"
//line /usr/local/go/src/encoding/json/encode.go:564
	}
//line /usr/local/go/src/encoding/json/encode.go:564
	// _ = "end of CoverTab[27577]"
//line /usr/local/go/src/encoding/json/encode.go:564
	_go_fuzz_dep_.CoverTab[27578]++
							e.Write(b)
							if opts.quoted {
//line /usr/local/go/src/encoding/json/encode.go:566
		_go_fuzz_dep_.CoverTab[27581]++
								e.WriteByte('"')
//line /usr/local/go/src/encoding/json/encode.go:567
		// _ = "end of CoverTab[27581]"
	} else {
//line /usr/local/go/src/encoding/json/encode.go:568
		_go_fuzz_dep_.CoverTab[27582]++
//line /usr/local/go/src/encoding/json/encode.go:568
		// _ = "end of CoverTab[27582]"
//line /usr/local/go/src/encoding/json/encode.go:568
	}
//line /usr/local/go/src/encoding/json/encode.go:568
	// _ = "end of CoverTab[27578]"
}

type floatEncoder int	// number of bits

func (bits floatEncoder) encode(e *encodeState, v reflect.Value, opts encOpts) {
//line /usr/local/go/src/encoding/json/encode.go:573
	_go_fuzz_dep_.CoverTab[27583]++
							f := v.Float()
							if math.IsInf(f, 0) || func() bool {
//line /usr/local/go/src/encoding/json/encode.go:575
		_go_fuzz_dep_.CoverTab[27588]++
//line /usr/local/go/src/encoding/json/encode.go:575
		return math.IsNaN(f)
//line /usr/local/go/src/encoding/json/encode.go:575
		// _ = "end of CoverTab[27588]"
//line /usr/local/go/src/encoding/json/encode.go:575
	}() {
//line /usr/local/go/src/encoding/json/encode.go:575
		_go_fuzz_dep_.CoverTab[27589]++
								e.error(&UnsupportedValueError{v, strconv.FormatFloat(f, 'g', -1, int(bits))})
//line /usr/local/go/src/encoding/json/encode.go:576
		// _ = "end of CoverTab[27589]"
	} else {
//line /usr/local/go/src/encoding/json/encode.go:577
		_go_fuzz_dep_.CoverTab[27590]++
//line /usr/local/go/src/encoding/json/encode.go:577
		// _ = "end of CoverTab[27590]"
//line /usr/local/go/src/encoding/json/encode.go:577
	}
//line /usr/local/go/src/encoding/json/encode.go:577
	// _ = "end of CoverTab[27583]"
//line /usr/local/go/src/encoding/json/encode.go:577
	_go_fuzz_dep_.CoverTab[27584]++

//line /usr/local/go/src/encoding/json/encode.go:584
	b := e.scratch[:0]
	abs := math.Abs(f)
	fmt := byte('f')

	if abs != 0 {
//line /usr/local/go/src/encoding/json/encode.go:588
		_go_fuzz_dep_.CoverTab[27591]++
								if bits == 64 && func() bool {
//line /usr/local/go/src/encoding/json/encode.go:589
			_go_fuzz_dep_.CoverTab[27592]++
//line /usr/local/go/src/encoding/json/encode.go:589
			return (abs < 1e-6 || func() bool {
//line /usr/local/go/src/encoding/json/encode.go:589
				_go_fuzz_dep_.CoverTab[27593]++
//line /usr/local/go/src/encoding/json/encode.go:589
				return abs >= 1e21
//line /usr/local/go/src/encoding/json/encode.go:589
				// _ = "end of CoverTab[27593]"
//line /usr/local/go/src/encoding/json/encode.go:589
			}())
//line /usr/local/go/src/encoding/json/encode.go:589
			// _ = "end of CoverTab[27592]"
//line /usr/local/go/src/encoding/json/encode.go:589
		}() || func() bool {
//line /usr/local/go/src/encoding/json/encode.go:589
			_go_fuzz_dep_.CoverTab[27594]++
//line /usr/local/go/src/encoding/json/encode.go:589
			return bits == 32 && func() bool {
//line /usr/local/go/src/encoding/json/encode.go:589
				_go_fuzz_dep_.CoverTab[27595]++
//line /usr/local/go/src/encoding/json/encode.go:589
				return (float32(abs) < 1e-6 || func() bool {
//line /usr/local/go/src/encoding/json/encode.go:589
					_go_fuzz_dep_.CoverTab[27596]++
//line /usr/local/go/src/encoding/json/encode.go:589
					return float32(abs) >= 1e21
//line /usr/local/go/src/encoding/json/encode.go:589
					// _ = "end of CoverTab[27596]"
//line /usr/local/go/src/encoding/json/encode.go:589
				}())
//line /usr/local/go/src/encoding/json/encode.go:589
				// _ = "end of CoverTab[27595]"
//line /usr/local/go/src/encoding/json/encode.go:589
			}()
//line /usr/local/go/src/encoding/json/encode.go:589
			// _ = "end of CoverTab[27594]"
//line /usr/local/go/src/encoding/json/encode.go:589
		}() {
//line /usr/local/go/src/encoding/json/encode.go:589
			_go_fuzz_dep_.CoverTab[27597]++
									fmt = 'e'
//line /usr/local/go/src/encoding/json/encode.go:590
			// _ = "end of CoverTab[27597]"
		} else {
//line /usr/local/go/src/encoding/json/encode.go:591
			_go_fuzz_dep_.CoverTab[27598]++
//line /usr/local/go/src/encoding/json/encode.go:591
			// _ = "end of CoverTab[27598]"
//line /usr/local/go/src/encoding/json/encode.go:591
		}
//line /usr/local/go/src/encoding/json/encode.go:591
		// _ = "end of CoverTab[27591]"
	} else {
//line /usr/local/go/src/encoding/json/encode.go:592
		_go_fuzz_dep_.CoverTab[27599]++
//line /usr/local/go/src/encoding/json/encode.go:592
		// _ = "end of CoverTab[27599]"
//line /usr/local/go/src/encoding/json/encode.go:592
	}
//line /usr/local/go/src/encoding/json/encode.go:592
	// _ = "end of CoverTab[27584]"
//line /usr/local/go/src/encoding/json/encode.go:592
	_go_fuzz_dep_.CoverTab[27585]++
							b = strconv.AppendFloat(b, f, fmt, -1, int(bits))
							if fmt == 'e' {
//line /usr/local/go/src/encoding/json/encode.go:594
		_go_fuzz_dep_.CoverTab[27600]++

								n := len(b)
								if n >= 4 && func() bool {
//line /usr/local/go/src/encoding/json/encode.go:597
			_go_fuzz_dep_.CoverTab[27601]++
//line /usr/local/go/src/encoding/json/encode.go:597
			return b[n-4] == 'e'
//line /usr/local/go/src/encoding/json/encode.go:597
			// _ = "end of CoverTab[27601]"
//line /usr/local/go/src/encoding/json/encode.go:597
		}() && func() bool {
//line /usr/local/go/src/encoding/json/encode.go:597
			_go_fuzz_dep_.CoverTab[27602]++
//line /usr/local/go/src/encoding/json/encode.go:597
			return b[n-3] == '-'
//line /usr/local/go/src/encoding/json/encode.go:597
			// _ = "end of CoverTab[27602]"
//line /usr/local/go/src/encoding/json/encode.go:597
		}() && func() bool {
//line /usr/local/go/src/encoding/json/encode.go:597
			_go_fuzz_dep_.CoverTab[27603]++
//line /usr/local/go/src/encoding/json/encode.go:597
			return b[n-2] == '0'
//line /usr/local/go/src/encoding/json/encode.go:597
			// _ = "end of CoverTab[27603]"
//line /usr/local/go/src/encoding/json/encode.go:597
		}() {
//line /usr/local/go/src/encoding/json/encode.go:597
			_go_fuzz_dep_.CoverTab[27604]++
									b[n-2] = b[n-1]
									b = b[:n-1]
//line /usr/local/go/src/encoding/json/encode.go:599
			// _ = "end of CoverTab[27604]"
		} else {
//line /usr/local/go/src/encoding/json/encode.go:600
			_go_fuzz_dep_.CoverTab[27605]++
//line /usr/local/go/src/encoding/json/encode.go:600
			// _ = "end of CoverTab[27605]"
//line /usr/local/go/src/encoding/json/encode.go:600
		}
//line /usr/local/go/src/encoding/json/encode.go:600
		// _ = "end of CoverTab[27600]"
	} else {
//line /usr/local/go/src/encoding/json/encode.go:601
		_go_fuzz_dep_.CoverTab[27606]++
//line /usr/local/go/src/encoding/json/encode.go:601
		// _ = "end of CoverTab[27606]"
//line /usr/local/go/src/encoding/json/encode.go:601
	}
//line /usr/local/go/src/encoding/json/encode.go:601
	// _ = "end of CoverTab[27585]"
//line /usr/local/go/src/encoding/json/encode.go:601
	_go_fuzz_dep_.CoverTab[27586]++

							if opts.quoted {
//line /usr/local/go/src/encoding/json/encode.go:603
		_go_fuzz_dep_.CoverTab[27607]++
								e.WriteByte('"')
//line /usr/local/go/src/encoding/json/encode.go:604
		// _ = "end of CoverTab[27607]"
	} else {
//line /usr/local/go/src/encoding/json/encode.go:605
		_go_fuzz_dep_.CoverTab[27608]++
//line /usr/local/go/src/encoding/json/encode.go:605
		// _ = "end of CoverTab[27608]"
//line /usr/local/go/src/encoding/json/encode.go:605
	}
//line /usr/local/go/src/encoding/json/encode.go:605
	// _ = "end of CoverTab[27586]"
//line /usr/local/go/src/encoding/json/encode.go:605
	_go_fuzz_dep_.CoverTab[27587]++
							e.Write(b)
							if opts.quoted {
//line /usr/local/go/src/encoding/json/encode.go:607
		_go_fuzz_dep_.CoverTab[27609]++
								e.WriteByte('"')
//line /usr/local/go/src/encoding/json/encode.go:608
		// _ = "end of CoverTab[27609]"
	} else {
//line /usr/local/go/src/encoding/json/encode.go:609
		_go_fuzz_dep_.CoverTab[27610]++
//line /usr/local/go/src/encoding/json/encode.go:609
		// _ = "end of CoverTab[27610]"
//line /usr/local/go/src/encoding/json/encode.go:609
	}
//line /usr/local/go/src/encoding/json/encode.go:609
	// _ = "end of CoverTab[27587]"
}

var (
	float32Encoder	= (floatEncoder(32)).encode
	float64Encoder	= (floatEncoder(64)).encode
)

func stringEncoder(e *encodeState, v reflect.Value, opts encOpts) {
//line /usr/local/go/src/encoding/json/encode.go:617
	_go_fuzz_dep_.CoverTab[27611]++
							if v.Type() == numberType {
//line /usr/local/go/src/encoding/json/encode.go:618
		_go_fuzz_dep_.CoverTab[27613]++
								numStr := v.String()

//line /usr/local/go/src/encoding/json/encode.go:622
		if numStr == "" {
//line /usr/local/go/src/encoding/json/encode.go:622
			_go_fuzz_dep_.CoverTab[27618]++
									numStr = "0"
//line /usr/local/go/src/encoding/json/encode.go:623
			// _ = "end of CoverTab[27618]"
		} else {
//line /usr/local/go/src/encoding/json/encode.go:624
			_go_fuzz_dep_.CoverTab[27619]++
//line /usr/local/go/src/encoding/json/encode.go:624
			// _ = "end of CoverTab[27619]"
//line /usr/local/go/src/encoding/json/encode.go:624
		}
//line /usr/local/go/src/encoding/json/encode.go:624
		// _ = "end of CoverTab[27613]"
//line /usr/local/go/src/encoding/json/encode.go:624
		_go_fuzz_dep_.CoverTab[27614]++
								if !isValidNumber(numStr) {
//line /usr/local/go/src/encoding/json/encode.go:625
			_go_fuzz_dep_.CoverTab[27620]++
									e.error(fmt.Errorf("json: invalid number literal %q", numStr))
//line /usr/local/go/src/encoding/json/encode.go:626
			// _ = "end of CoverTab[27620]"
		} else {
//line /usr/local/go/src/encoding/json/encode.go:627
			_go_fuzz_dep_.CoverTab[27621]++
//line /usr/local/go/src/encoding/json/encode.go:627
			// _ = "end of CoverTab[27621]"
//line /usr/local/go/src/encoding/json/encode.go:627
		}
//line /usr/local/go/src/encoding/json/encode.go:627
		// _ = "end of CoverTab[27614]"
//line /usr/local/go/src/encoding/json/encode.go:627
		_go_fuzz_dep_.CoverTab[27615]++
								if opts.quoted {
//line /usr/local/go/src/encoding/json/encode.go:628
			_go_fuzz_dep_.CoverTab[27622]++
									e.WriteByte('"')
//line /usr/local/go/src/encoding/json/encode.go:629
			// _ = "end of CoverTab[27622]"
		} else {
//line /usr/local/go/src/encoding/json/encode.go:630
			_go_fuzz_dep_.CoverTab[27623]++
//line /usr/local/go/src/encoding/json/encode.go:630
			// _ = "end of CoverTab[27623]"
//line /usr/local/go/src/encoding/json/encode.go:630
		}
//line /usr/local/go/src/encoding/json/encode.go:630
		// _ = "end of CoverTab[27615]"
//line /usr/local/go/src/encoding/json/encode.go:630
		_go_fuzz_dep_.CoverTab[27616]++
								e.WriteString(numStr)
								if opts.quoted {
//line /usr/local/go/src/encoding/json/encode.go:632
			_go_fuzz_dep_.CoverTab[27624]++
									e.WriteByte('"')
//line /usr/local/go/src/encoding/json/encode.go:633
			// _ = "end of CoverTab[27624]"
		} else {
//line /usr/local/go/src/encoding/json/encode.go:634
			_go_fuzz_dep_.CoverTab[27625]++
//line /usr/local/go/src/encoding/json/encode.go:634
			// _ = "end of CoverTab[27625]"
//line /usr/local/go/src/encoding/json/encode.go:634
		}
//line /usr/local/go/src/encoding/json/encode.go:634
		// _ = "end of CoverTab[27616]"
//line /usr/local/go/src/encoding/json/encode.go:634
		_go_fuzz_dep_.CoverTab[27617]++
								return
//line /usr/local/go/src/encoding/json/encode.go:635
		// _ = "end of CoverTab[27617]"
	} else {
//line /usr/local/go/src/encoding/json/encode.go:636
		_go_fuzz_dep_.CoverTab[27626]++
//line /usr/local/go/src/encoding/json/encode.go:636
		// _ = "end of CoverTab[27626]"
//line /usr/local/go/src/encoding/json/encode.go:636
	}
//line /usr/local/go/src/encoding/json/encode.go:636
	// _ = "end of CoverTab[27611]"
//line /usr/local/go/src/encoding/json/encode.go:636
	_go_fuzz_dep_.CoverTab[27612]++
							if opts.quoted {
//line /usr/local/go/src/encoding/json/encode.go:637
		_go_fuzz_dep_.CoverTab[27627]++
								e2 := newEncodeState()

//line /usr/local/go/src/encoding/json/encode.go:641
		e2.string(v.String(), opts.escapeHTML)
								e.stringBytes(e2.Bytes(), false)
								encodeStatePool.Put(e2)
//line /usr/local/go/src/encoding/json/encode.go:643
		// _ = "end of CoverTab[27627]"
	} else {
//line /usr/local/go/src/encoding/json/encode.go:644
		_go_fuzz_dep_.CoverTab[27628]++
								e.string(v.String(), opts.escapeHTML)
//line /usr/local/go/src/encoding/json/encode.go:645
		// _ = "end of CoverTab[27628]"
	}
//line /usr/local/go/src/encoding/json/encode.go:646
	// _ = "end of CoverTab[27612]"
}

// isValidNumber reports whether s is a valid JSON number literal.
func isValidNumber(s string) bool {
//line /usr/local/go/src/encoding/json/encode.go:650
	_go_fuzz_dep_.CoverTab[27629]++

//line /usr/local/go/src/encoding/json/encode.go:655
	if s == "" {
//line /usr/local/go/src/encoding/json/encode.go:655
		_go_fuzz_dep_.CoverTab[27635]++
								return false
//line /usr/local/go/src/encoding/json/encode.go:656
		// _ = "end of CoverTab[27635]"
	} else {
//line /usr/local/go/src/encoding/json/encode.go:657
		_go_fuzz_dep_.CoverTab[27636]++
//line /usr/local/go/src/encoding/json/encode.go:657
		// _ = "end of CoverTab[27636]"
//line /usr/local/go/src/encoding/json/encode.go:657
	}
//line /usr/local/go/src/encoding/json/encode.go:657
	// _ = "end of CoverTab[27629]"
//line /usr/local/go/src/encoding/json/encode.go:657
	_go_fuzz_dep_.CoverTab[27630]++

//line /usr/local/go/src/encoding/json/encode.go:660
	if s[0] == '-' {
//line /usr/local/go/src/encoding/json/encode.go:660
		_go_fuzz_dep_.CoverTab[27637]++
								s = s[1:]
								if s == "" {
//line /usr/local/go/src/encoding/json/encode.go:662
			_go_fuzz_dep_.CoverTab[27638]++
									return false
//line /usr/local/go/src/encoding/json/encode.go:663
			// _ = "end of CoverTab[27638]"
		} else {
//line /usr/local/go/src/encoding/json/encode.go:664
			_go_fuzz_dep_.CoverTab[27639]++
//line /usr/local/go/src/encoding/json/encode.go:664
			// _ = "end of CoverTab[27639]"
//line /usr/local/go/src/encoding/json/encode.go:664
		}
//line /usr/local/go/src/encoding/json/encode.go:664
		// _ = "end of CoverTab[27637]"
	} else {
//line /usr/local/go/src/encoding/json/encode.go:665
		_go_fuzz_dep_.CoverTab[27640]++
//line /usr/local/go/src/encoding/json/encode.go:665
		// _ = "end of CoverTab[27640]"
//line /usr/local/go/src/encoding/json/encode.go:665
	}
//line /usr/local/go/src/encoding/json/encode.go:665
	// _ = "end of CoverTab[27630]"
//line /usr/local/go/src/encoding/json/encode.go:665
	_go_fuzz_dep_.CoverTab[27631]++

//line /usr/local/go/src/encoding/json/encode.go:668
	switch {
	default:
//line /usr/local/go/src/encoding/json/encode.go:669
		_go_fuzz_dep_.CoverTab[27641]++
								return false
//line /usr/local/go/src/encoding/json/encode.go:670
		// _ = "end of CoverTab[27641]"

	case s[0] == '0':
//line /usr/local/go/src/encoding/json/encode.go:672
		_go_fuzz_dep_.CoverTab[27642]++
								s = s[1:]
//line /usr/local/go/src/encoding/json/encode.go:673
		// _ = "end of CoverTab[27642]"

	case '1' <= s[0] && func() bool {
//line /usr/local/go/src/encoding/json/encode.go:675
		_go_fuzz_dep_.CoverTab[27644]++
//line /usr/local/go/src/encoding/json/encode.go:675
		return s[0] <= '9'
//line /usr/local/go/src/encoding/json/encode.go:675
		// _ = "end of CoverTab[27644]"
//line /usr/local/go/src/encoding/json/encode.go:675
	}():
//line /usr/local/go/src/encoding/json/encode.go:675
		_go_fuzz_dep_.CoverTab[27643]++
								s = s[1:]
								for len(s) > 0 && func() bool {
//line /usr/local/go/src/encoding/json/encode.go:677
			_go_fuzz_dep_.CoverTab[27645]++
//line /usr/local/go/src/encoding/json/encode.go:677
			return '0' <= s[0]
//line /usr/local/go/src/encoding/json/encode.go:677
			// _ = "end of CoverTab[27645]"
//line /usr/local/go/src/encoding/json/encode.go:677
		}() && func() bool {
//line /usr/local/go/src/encoding/json/encode.go:677
			_go_fuzz_dep_.CoverTab[27646]++
//line /usr/local/go/src/encoding/json/encode.go:677
			return s[0] <= '9'
//line /usr/local/go/src/encoding/json/encode.go:677
			// _ = "end of CoverTab[27646]"
//line /usr/local/go/src/encoding/json/encode.go:677
		}() {
//line /usr/local/go/src/encoding/json/encode.go:677
			_go_fuzz_dep_.CoverTab[27647]++
									s = s[1:]
//line /usr/local/go/src/encoding/json/encode.go:678
			// _ = "end of CoverTab[27647]"
		}
//line /usr/local/go/src/encoding/json/encode.go:679
		// _ = "end of CoverTab[27643]"
	}
//line /usr/local/go/src/encoding/json/encode.go:680
	// _ = "end of CoverTab[27631]"
//line /usr/local/go/src/encoding/json/encode.go:680
	_go_fuzz_dep_.CoverTab[27632]++

//line /usr/local/go/src/encoding/json/encode.go:683
	if len(s) >= 2 && func() bool {
//line /usr/local/go/src/encoding/json/encode.go:683
		_go_fuzz_dep_.CoverTab[27648]++
//line /usr/local/go/src/encoding/json/encode.go:683
		return s[0] == '.'
//line /usr/local/go/src/encoding/json/encode.go:683
		// _ = "end of CoverTab[27648]"
//line /usr/local/go/src/encoding/json/encode.go:683
	}() && func() bool {
//line /usr/local/go/src/encoding/json/encode.go:683
		_go_fuzz_dep_.CoverTab[27649]++
//line /usr/local/go/src/encoding/json/encode.go:683
		return '0' <= s[1]
//line /usr/local/go/src/encoding/json/encode.go:683
		// _ = "end of CoverTab[27649]"
//line /usr/local/go/src/encoding/json/encode.go:683
	}() && func() bool {
//line /usr/local/go/src/encoding/json/encode.go:683
		_go_fuzz_dep_.CoverTab[27650]++
//line /usr/local/go/src/encoding/json/encode.go:683
		return s[1] <= '9'
//line /usr/local/go/src/encoding/json/encode.go:683
		// _ = "end of CoverTab[27650]"
//line /usr/local/go/src/encoding/json/encode.go:683
	}() {
//line /usr/local/go/src/encoding/json/encode.go:683
		_go_fuzz_dep_.CoverTab[27651]++
								s = s[2:]
								for len(s) > 0 && func() bool {
//line /usr/local/go/src/encoding/json/encode.go:685
			_go_fuzz_dep_.CoverTab[27652]++
//line /usr/local/go/src/encoding/json/encode.go:685
			return '0' <= s[0]
//line /usr/local/go/src/encoding/json/encode.go:685
			// _ = "end of CoverTab[27652]"
//line /usr/local/go/src/encoding/json/encode.go:685
		}() && func() bool {
//line /usr/local/go/src/encoding/json/encode.go:685
			_go_fuzz_dep_.CoverTab[27653]++
//line /usr/local/go/src/encoding/json/encode.go:685
			return s[0] <= '9'
//line /usr/local/go/src/encoding/json/encode.go:685
			// _ = "end of CoverTab[27653]"
//line /usr/local/go/src/encoding/json/encode.go:685
		}() {
//line /usr/local/go/src/encoding/json/encode.go:685
			_go_fuzz_dep_.CoverTab[27654]++
									s = s[1:]
//line /usr/local/go/src/encoding/json/encode.go:686
			// _ = "end of CoverTab[27654]"
		}
//line /usr/local/go/src/encoding/json/encode.go:687
		// _ = "end of CoverTab[27651]"
	} else {
//line /usr/local/go/src/encoding/json/encode.go:688
		_go_fuzz_dep_.CoverTab[27655]++
//line /usr/local/go/src/encoding/json/encode.go:688
		// _ = "end of CoverTab[27655]"
//line /usr/local/go/src/encoding/json/encode.go:688
	}
//line /usr/local/go/src/encoding/json/encode.go:688
	// _ = "end of CoverTab[27632]"
//line /usr/local/go/src/encoding/json/encode.go:688
	_go_fuzz_dep_.CoverTab[27633]++

//line /usr/local/go/src/encoding/json/encode.go:692
	if len(s) >= 2 && func() bool {
//line /usr/local/go/src/encoding/json/encode.go:692
		_go_fuzz_dep_.CoverTab[27656]++
//line /usr/local/go/src/encoding/json/encode.go:692
		return (s[0] == 'e' || func() bool {
//line /usr/local/go/src/encoding/json/encode.go:692
			_go_fuzz_dep_.CoverTab[27657]++
//line /usr/local/go/src/encoding/json/encode.go:692
			return s[0] == 'E'
//line /usr/local/go/src/encoding/json/encode.go:692
			// _ = "end of CoverTab[27657]"
//line /usr/local/go/src/encoding/json/encode.go:692
		}())
//line /usr/local/go/src/encoding/json/encode.go:692
		// _ = "end of CoverTab[27656]"
//line /usr/local/go/src/encoding/json/encode.go:692
	}() {
//line /usr/local/go/src/encoding/json/encode.go:692
		_go_fuzz_dep_.CoverTab[27658]++
								s = s[1:]
								if s[0] == '+' || func() bool {
//line /usr/local/go/src/encoding/json/encode.go:694
			_go_fuzz_dep_.CoverTab[27660]++
//line /usr/local/go/src/encoding/json/encode.go:694
			return s[0] == '-'
//line /usr/local/go/src/encoding/json/encode.go:694
			// _ = "end of CoverTab[27660]"
//line /usr/local/go/src/encoding/json/encode.go:694
		}() {
//line /usr/local/go/src/encoding/json/encode.go:694
			_go_fuzz_dep_.CoverTab[27661]++
									s = s[1:]
									if s == "" {
//line /usr/local/go/src/encoding/json/encode.go:696
				_go_fuzz_dep_.CoverTab[27662]++
										return false
//line /usr/local/go/src/encoding/json/encode.go:697
				// _ = "end of CoverTab[27662]"
			} else {
//line /usr/local/go/src/encoding/json/encode.go:698
				_go_fuzz_dep_.CoverTab[27663]++
//line /usr/local/go/src/encoding/json/encode.go:698
				// _ = "end of CoverTab[27663]"
//line /usr/local/go/src/encoding/json/encode.go:698
			}
//line /usr/local/go/src/encoding/json/encode.go:698
			// _ = "end of CoverTab[27661]"
		} else {
//line /usr/local/go/src/encoding/json/encode.go:699
			_go_fuzz_dep_.CoverTab[27664]++
//line /usr/local/go/src/encoding/json/encode.go:699
			// _ = "end of CoverTab[27664]"
//line /usr/local/go/src/encoding/json/encode.go:699
		}
//line /usr/local/go/src/encoding/json/encode.go:699
		// _ = "end of CoverTab[27658]"
//line /usr/local/go/src/encoding/json/encode.go:699
		_go_fuzz_dep_.CoverTab[27659]++
								for len(s) > 0 && func() bool {
//line /usr/local/go/src/encoding/json/encode.go:700
			_go_fuzz_dep_.CoverTab[27665]++
//line /usr/local/go/src/encoding/json/encode.go:700
			return '0' <= s[0]
//line /usr/local/go/src/encoding/json/encode.go:700
			// _ = "end of CoverTab[27665]"
//line /usr/local/go/src/encoding/json/encode.go:700
		}() && func() bool {
//line /usr/local/go/src/encoding/json/encode.go:700
			_go_fuzz_dep_.CoverTab[27666]++
//line /usr/local/go/src/encoding/json/encode.go:700
			return s[0] <= '9'
//line /usr/local/go/src/encoding/json/encode.go:700
			// _ = "end of CoverTab[27666]"
//line /usr/local/go/src/encoding/json/encode.go:700
		}() {
//line /usr/local/go/src/encoding/json/encode.go:700
			_go_fuzz_dep_.CoverTab[27667]++
									s = s[1:]
//line /usr/local/go/src/encoding/json/encode.go:701
			// _ = "end of CoverTab[27667]"
		}
//line /usr/local/go/src/encoding/json/encode.go:702
		// _ = "end of CoverTab[27659]"
	} else {
//line /usr/local/go/src/encoding/json/encode.go:703
		_go_fuzz_dep_.CoverTab[27668]++
//line /usr/local/go/src/encoding/json/encode.go:703
		// _ = "end of CoverTab[27668]"
//line /usr/local/go/src/encoding/json/encode.go:703
	}
//line /usr/local/go/src/encoding/json/encode.go:703
	// _ = "end of CoverTab[27633]"
//line /usr/local/go/src/encoding/json/encode.go:703
	_go_fuzz_dep_.CoverTab[27634]++

//line /usr/local/go/src/encoding/json/encode.go:706
	return s == ""
//line /usr/local/go/src/encoding/json/encode.go:706
	// _ = "end of CoverTab[27634]"
}

func interfaceEncoder(e *encodeState, v reflect.Value, opts encOpts) {
//line /usr/local/go/src/encoding/json/encode.go:709
	_go_fuzz_dep_.CoverTab[27669]++
							if v.IsNil() {
//line /usr/local/go/src/encoding/json/encode.go:710
		_go_fuzz_dep_.CoverTab[27671]++
								e.WriteString("null")
								return
//line /usr/local/go/src/encoding/json/encode.go:712
		// _ = "end of CoverTab[27671]"
	} else {
//line /usr/local/go/src/encoding/json/encode.go:713
		_go_fuzz_dep_.CoverTab[27672]++
//line /usr/local/go/src/encoding/json/encode.go:713
		// _ = "end of CoverTab[27672]"
//line /usr/local/go/src/encoding/json/encode.go:713
	}
//line /usr/local/go/src/encoding/json/encode.go:713
	// _ = "end of CoverTab[27669]"
//line /usr/local/go/src/encoding/json/encode.go:713
	_go_fuzz_dep_.CoverTab[27670]++
							e.reflectValue(v.Elem(), opts)
//line /usr/local/go/src/encoding/json/encode.go:714
	// _ = "end of CoverTab[27670]"
}

func unsupportedTypeEncoder(e *encodeState, v reflect.Value, _ encOpts) {
//line /usr/local/go/src/encoding/json/encode.go:717
	_go_fuzz_dep_.CoverTab[27673]++
							e.error(&UnsupportedTypeError{v.Type()})
//line /usr/local/go/src/encoding/json/encode.go:718
	// _ = "end of CoverTab[27673]"
}

type structEncoder struct {
	fields structFields
}

type structFields struct {
	list		[]field
	nameIndex	map[string]int
}

func (se structEncoder) encode(e *encodeState, v reflect.Value, opts encOpts) {
//line /usr/local/go/src/encoding/json/encode.go:730
	_go_fuzz_dep_.CoverTab[27674]++
							next := byte('{')
FieldLoop:
	for i := range se.fields.list {
//line /usr/local/go/src/encoding/json/encode.go:733
		_go_fuzz_dep_.CoverTab[27676]++
								f := &se.fields.list[i]

//line /usr/local/go/src/encoding/json/encode.go:737
		fv := v
		for _, i := range f.index {
//line /usr/local/go/src/encoding/json/encode.go:738
			_go_fuzz_dep_.CoverTab[27680]++
									if fv.Kind() == reflect.Pointer {
//line /usr/local/go/src/encoding/json/encode.go:739
				_go_fuzz_dep_.CoverTab[27682]++
										if fv.IsNil() {
//line /usr/local/go/src/encoding/json/encode.go:740
					_go_fuzz_dep_.CoverTab[27684]++
											continue FieldLoop
//line /usr/local/go/src/encoding/json/encode.go:741
					// _ = "end of CoverTab[27684]"
				} else {
//line /usr/local/go/src/encoding/json/encode.go:742
					_go_fuzz_dep_.CoverTab[27685]++
//line /usr/local/go/src/encoding/json/encode.go:742
					// _ = "end of CoverTab[27685]"
//line /usr/local/go/src/encoding/json/encode.go:742
				}
//line /usr/local/go/src/encoding/json/encode.go:742
				// _ = "end of CoverTab[27682]"
//line /usr/local/go/src/encoding/json/encode.go:742
				_go_fuzz_dep_.CoverTab[27683]++
										fv = fv.Elem()
//line /usr/local/go/src/encoding/json/encode.go:743
				// _ = "end of CoverTab[27683]"
			} else {
//line /usr/local/go/src/encoding/json/encode.go:744
				_go_fuzz_dep_.CoverTab[27686]++
//line /usr/local/go/src/encoding/json/encode.go:744
				// _ = "end of CoverTab[27686]"
//line /usr/local/go/src/encoding/json/encode.go:744
			}
//line /usr/local/go/src/encoding/json/encode.go:744
			// _ = "end of CoverTab[27680]"
//line /usr/local/go/src/encoding/json/encode.go:744
			_go_fuzz_dep_.CoverTab[27681]++
									fv = fv.Field(i)
//line /usr/local/go/src/encoding/json/encode.go:745
			// _ = "end of CoverTab[27681]"
		}
//line /usr/local/go/src/encoding/json/encode.go:746
		// _ = "end of CoverTab[27676]"
//line /usr/local/go/src/encoding/json/encode.go:746
		_go_fuzz_dep_.CoverTab[27677]++

								if f.omitEmpty && func() bool {
//line /usr/local/go/src/encoding/json/encode.go:748
			_go_fuzz_dep_.CoverTab[27687]++
//line /usr/local/go/src/encoding/json/encode.go:748
			return isEmptyValue(fv)
//line /usr/local/go/src/encoding/json/encode.go:748
			// _ = "end of CoverTab[27687]"
//line /usr/local/go/src/encoding/json/encode.go:748
		}() {
//line /usr/local/go/src/encoding/json/encode.go:748
			_go_fuzz_dep_.CoverTab[27688]++
									continue
//line /usr/local/go/src/encoding/json/encode.go:749
			// _ = "end of CoverTab[27688]"
		} else {
//line /usr/local/go/src/encoding/json/encode.go:750
			_go_fuzz_dep_.CoverTab[27689]++
//line /usr/local/go/src/encoding/json/encode.go:750
			// _ = "end of CoverTab[27689]"
//line /usr/local/go/src/encoding/json/encode.go:750
		}
//line /usr/local/go/src/encoding/json/encode.go:750
		// _ = "end of CoverTab[27677]"
//line /usr/local/go/src/encoding/json/encode.go:750
		_go_fuzz_dep_.CoverTab[27678]++
								e.WriteByte(next)
								next = ','
								if opts.escapeHTML {
//line /usr/local/go/src/encoding/json/encode.go:753
			_go_fuzz_dep_.CoverTab[27690]++
									e.WriteString(f.nameEscHTML)
//line /usr/local/go/src/encoding/json/encode.go:754
			// _ = "end of CoverTab[27690]"
		} else {
//line /usr/local/go/src/encoding/json/encode.go:755
			_go_fuzz_dep_.CoverTab[27691]++
									e.WriteString(f.nameNonEsc)
//line /usr/local/go/src/encoding/json/encode.go:756
			// _ = "end of CoverTab[27691]"
		}
//line /usr/local/go/src/encoding/json/encode.go:757
		// _ = "end of CoverTab[27678]"
//line /usr/local/go/src/encoding/json/encode.go:757
		_go_fuzz_dep_.CoverTab[27679]++
								opts.quoted = f.quoted
								f.encoder(e, fv, opts)
//line /usr/local/go/src/encoding/json/encode.go:759
		// _ = "end of CoverTab[27679]"
	}
//line /usr/local/go/src/encoding/json/encode.go:760
	// _ = "end of CoverTab[27674]"
//line /usr/local/go/src/encoding/json/encode.go:760
	_go_fuzz_dep_.CoverTab[27675]++
							if next == '{' {
//line /usr/local/go/src/encoding/json/encode.go:761
		_go_fuzz_dep_.CoverTab[27692]++
								e.WriteString("{}")
//line /usr/local/go/src/encoding/json/encode.go:762
		// _ = "end of CoverTab[27692]"
	} else {
//line /usr/local/go/src/encoding/json/encode.go:763
		_go_fuzz_dep_.CoverTab[27693]++
								e.WriteByte('}')
//line /usr/local/go/src/encoding/json/encode.go:764
		// _ = "end of CoverTab[27693]"
	}
//line /usr/local/go/src/encoding/json/encode.go:765
	// _ = "end of CoverTab[27675]"
}

func newStructEncoder(t reflect.Type) encoderFunc {
//line /usr/local/go/src/encoding/json/encode.go:768
	_go_fuzz_dep_.CoverTab[27694]++
							se := structEncoder{fields: cachedTypeFields(t)}
							return se.encode
//line /usr/local/go/src/encoding/json/encode.go:770
	// _ = "end of CoverTab[27694]"
}

type mapEncoder struct {
	elemEnc encoderFunc
}

func (me mapEncoder) encode(e *encodeState, v reflect.Value, opts encOpts) {
//line /usr/local/go/src/encoding/json/encode.go:777
	_go_fuzz_dep_.CoverTab[27695]++
							if v.IsNil() {
//line /usr/local/go/src/encoding/json/encode.go:778
		_go_fuzz_dep_.CoverTab[27701]++
								e.WriteString("null")
								return
//line /usr/local/go/src/encoding/json/encode.go:780
		// _ = "end of CoverTab[27701]"
	} else {
//line /usr/local/go/src/encoding/json/encode.go:781
		_go_fuzz_dep_.CoverTab[27702]++
//line /usr/local/go/src/encoding/json/encode.go:781
		// _ = "end of CoverTab[27702]"
//line /usr/local/go/src/encoding/json/encode.go:781
	}
//line /usr/local/go/src/encoding/json/encode.go:781
	// _ = "end of CoverTab[27695]"
//line /usr/local/go/src/encoding/json/encode.go:781
	_go_fuzz_dep_.CoverTab[27696]++
							if e.ptrLevel++; e.ptrLevel > startDetectingCyclesAfter {
//line /usr/local/go/src/encoding/json/encode.go:782
		_go_fuzz_dep_.CoverTab[27703]++

//line /usr/local/go/src/encoding/json/encode.go:785
		ptr := v.UnsafePointer()
		if _, ok := e.ptrSeen[ptr]; ok {
//line /usr/local/go/src/encoding/json/encode.go:786
			_go_fuzz_dep_.CoverTab[27705]++
									e.error(&UnsupportedValueError{v, fmt.Sprintf("encountered a cycle via %s", v.Type())})
//line /usr/local/go/src/encoding/json/encode.go:787
			// _ = "end of CoverTab[27705]"
		} else {
//line /usr/local/go/src/encoding/json/encode.go:788
			_go_fuzz_dep_.CoverTab[27706]++
//line /usr/local/go/src/encoding/json/encode.go:788
			// _ = "end of CoverTab[27706]"
//line /usr/local/go/src/encoding/json/encode.go:788
		}
//line /usr/local/go/src/encoding/json/encode.go:788
		// _ = "end of CoverTab[27703]"
//line /usr/local/go/src/encoding/json/encode.go:788
		_go_fuzz_dep_.CoverTab[27704]++
								e.ptrSeen[ptr] = struct{}{}
								defer delete(e.ptrSeen, ptr)
//line /usr/local/go/src/encoding/json/encode.go:790
		// _ = "end of CoverTab[27704]"
	} else {
//line /usr/local/go/src/encoding/json/encode.go:791
		_go_fuzz_dep_.CoverTab[27707]++
//line /usr/local/go/src/encoding/json/encode.go:791
		// _ = "end of CoverTab[27707]"
//line /usr/local/go/src/encoding/json/encode.go:791
	}
//line /usr/local/go/src/encoding/json/encode.go:791
	// _ = "end of CoverTab[27696]"
//line /usr/local/go/src/encoding/json/encode.go:791
	_go_fuzz_dep_.CoverTab[27697]++
							e.WriteByte('{')

//line /usr/local/go/src/encoding/json/encode.go:795
	sv := make([]reflectWithString, v.Len())
	mi := v.MapRange()
	for i := 0; mi.Next(); i++ {
//line /usr/local/go/src/encoding/json/encode.go:797
		_go_fuzz_dep_.CoverTab[27708]++
								sv[i].k = mi.Key()
								sv[i].v = mi.Value()
								if err := sv[i].resolve(); err != nil {
//line /usr/local/go/src/encoding/json/encode.go:800
			_go_fuzz_dep_.CoverTab[27709]++
									e.error(fmt.Errorf("json: encoding error for type %q: %q", v.Type().String(), err.Error()))
//line /usr/local/go/src/encoding/json/encode.go:801
			// _ = "end of CoverTab[27709]"
		} else {
//line /usr/local/go/src/encoding/json/encode.go:802
			_go_fuzz_dep_.CoverTab[27710]++
//line /usr/local/go/src/encoding/json/encode.go:802
			// _ = "end of CoverTab[27710]"
//line /usr/local/go/src/encoding/json/encode.go:802
		}
//line /usr/local/go/src/encoding/json/encode.go:802
		// _ = "end of CoverTab[27708]"
	}
//line /usr/local/go/src/encoding/json/encode.go:803
	// _ = "end of CoverTab[27697]"
//line /usr/local/go/src/encoding/json/encode.go:803
	_go_fuzz_dep_.CoverTab[27698]++
							sort.Slice(sv, func(i, j int) bool {
//line /usr/local/go/src/encoding/json/encode.go:804
		_go_fuzz_dep_.CoverTab[27711]++
//line /usr/local/go/src/encoding/json/encode.go:804
		return sv[i].ks < sv[j].ks
//line /usr/local/go/src/encoding/json/encode.go:804
		// _ = "end of CoverTab[27711]"
//line /usr/local/go/src/encoding/json/encode.go:804
	})
//line /usr/local/go/src/encoding/json/encode.go:804
	// _ = "end of CoverTab[27698]"
//line /usr/local/go/src/encoding/json/encode.go:804
	_go_fuzz_dep_.CoverTab[27699]++

							for i, kv := range sv {
//line /usr/local/go/src/encoding/json/encode.go:806
		_go_fuzz_dep_.CoverTab[27712]++
								if i > 0 {
//line /usr/local/go/src/encoding/json/encode.go:807
			_go_fuzz_dep_.CoverTab[27714]++
									e.WriteByte(',')
//line /usr/local/go/src/encoding/json/encode.go:808
			// _ = "end of CoverTab[27714]"
		} else {
//line /usr/local/go/src/encoding/json/encode.go:809
			_go_fuzz_dep_.CoverTab[27715]++
//line /usr/local/go/src/encoding/json/encode.go:809
			// _ = "end of CoverTab[27715]"
//line /usr/local/go/src/encoding/json/encode.go:809
		}
//line /usr/local/go/src/encoding/json/encode.go:809
		// _ = "end of CoverTab[27712]"
//line /usr/local/go/src/encoding/json/encode.go:809
		_go_fuzz_dep_.CoverTab[27713]++
								e.string(kv.ks, opts.escapeHTML)
								e.WriteByte(':')
								me.elemEnc(e, kv.v, opts)
//line /usr/local/go/src/encoding/json/encode.go:812
		// _ = "end of CoverTab[27713]"
	}
//line /usr/local/go/src/encoding/json/encode.go:813
	// _ = "end of CoverTab[27699]"
//line /usr/local/go/src/encoding/json/encode.go:813
	_go_fuzz_dep_.CoverTab[27700]++
							e.WriteByte('}')
							e.ptrLevel--
//line /usr/local/go/src/encoding/json/encode.go:815
	// _ = "end of CoverTab[27700]"
}

func newMapEncoder(t reflect.Type) encoderFunc {
//line /usr/local/go/src/encoding/json/encode.go:818
	_go_fuzz_dep_.CoverTab[27716]++
							switch t.Key().Kind() {
	case reflect.String,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
//line /usr/local/go/src/encoding/json/encode.go:822
		_go_fuzz_dep_.CoverTab[27718]++
//line /usr/local/go/src/encoding/json/encode.go:822
		// _ = "end of CoverTab[27718]"
	default:
//line /usr/local/go/src/encoding/json/encode.go:823
		_go_fuzz_dep_.CoverTab[27719]++
								if !t.Key().Implements(textMarshalerType) {
//line /usr/local/go/src/encoding/json/encode.go:824
			_go_fuzz_dep_.CoverTab[27720]++
									return unsupportedTypeEncoder
//line /usr/local/go/src/encoding/json/encode.go:825
			// _ = "end of CoverTab[27720]"
		} else {
//line /usr/local/go/src/encoding/json/encode.go:826
			_go_fuzz_dep_.CoverTab[27721]++
//line /usr/local/go/src/encoding/json/encode.go:826
			// _ = "end of CoverTab[27721]"
//line /usr/local/go/src/encoding/json/encode.go:826
		}
//line /usr/local/go/src/encoding/json/encode.go:826
		// _ = "end of CoverTab[27719]"
	}
//line /usr/local/go/src/encoding/json/encode.go:827
	// _ = "end of CoverTab[27716]"
//line /usr/local/go/src/encoding/json/encode.go:827
	_go_fuzz_dep_.CoverTab[27717]++
							me := mapEncoder{typeEncoder(t.Elem())}
							return me.encode
//line /usr/local/go/src/encoding/json/encode.go:829
	// _ = "end of CoverTab[27717]"
}

func encodeByteSlice(e *encodeState, v reflect.Value, _ encOpts) {
//line /usr/local/go/src/encoding/json/encode.go:832
	_go_fuzz_dep_.CoverTab[27722]++
							if v.IsNil() {
//line /usr/local/go/src/encoding/json/encode.go:833
		_go_fuzz_dep_.CoverTab[27725]++
								e.WriteString("null")
								return
//line /usr/local/go/src/encoding/json/encode.go:835
		// _ = "end of CoverTab[27725]"
	} else {
//line /usr/local/go/src/encoding/json/encode.go:836
		_go_fuzz_dep_.CoverTab[27726]++
//line /usr/local/go/src/encoding/json/encode.go:836
		// _ = "end of CoverTab[27726]"
//line /usr/local/go/src/encoding/json/encode.go:836
	}
//line /usr/local/go/src/encoding/json/encode.go:836
	// _ = "end of CoverTab[27722]"
//line /usr/local/go/src/encoding/json/encode.go:836
	_go_fuzz_dep_.CoverTab[27723]++
							s := v.Bytes()
							e.WriteByte('"')
							encodedLen := base64.StdEncoding.EncodedLen(len(s))
							if encodedLen <= len(e.scratch) {
//line /usr/local/go/src/encoding/json/encode.go:840
		_go_fuzz_dep_.CoverTab[27727]++

//line /usr/local/go/src/encoding/json/encode.go:843
		dst := e.scratch[:encodedLen]
								base64.StdEncoding.Encode(dst, s)
								e.Write(dst)
//line /usr/local/go/src/encoding/json/encode.go:845
		// _ = "end of CoverTab[27727]"
	} else {
//line /usr/local/go/src/encoding/json/encode.go:846
		_go_fuzz_dep_.CoverTab[27728]++
//line /usr/local/go/src/encoding/json/encode.go:846
		if encodedLen <= 1024 {
//line /usr/local/go/src/encoding/json/encode.go:846
			_go_fuzz_dep_.CoverTab[27729]++

//line /usr/local/go/src/encoding/json/encode.go:849
			dst := make([]byte, encodedLen)
									base64.StdEncoding.Encode(dst, s)
									e.Write(dst)
//line /usr/local/go/src/encoding/json/encode.go:851
			// _ = "end of CoverTab[27729]"
		} else {
//line /usr/local/go/src/encoding/json/encode.go:852
			_go_fuzz_dep_.CoverTab[27730]++

//line /usr/local/go/src/encoding/json/encode.go:855
			enc := base64.NewEncoder(base64.StdEncoding, e)
									enc.Write(s)
									enc.Close()
//line /usr/local/go/src/encoding/json/encode.go:857
			// _ = "end of CoverTab[27730]"
		}
//line /usr/local/go/src/encoding/json/encode.go:858
		// _ = "end of CoverTab[27728]"
//line /usr/local/go/src/encoding/json/encode.go:858
	}
//line /usr/local/go/src/encoding/json/encode.go:858
	// _ = "end of CoverTab[27723]"
//line /usr/local/go/src/encoding/json/encode.go:858
	_go_fuzz_dep_.CoverTab[27724]++
							e.WriteByte('"')
//line /usr/local/go/src/encoding/json/encode.go:859
	// _ = "end of CoverTab[27724]"
}

// sliceEncoder just wraps an arrayEncoder, checking to make sure the value isn't nil.
type sliceEncoder struct {
	arrayEnc encoderFunc
}

func (se sliceEncoder) encode(e *encodeState, v reflect.Value, opts encOpts) {
//line /usr/local/go/src/encoding/json/encode.go:867
	_go_fuzz_dep_.CoverTab[27731]++
							if v.IsNil() {
//line /usr/local/go/src/encoding/json/encode.go:868
		_go_fuzz_dep_.CoverTab[27734]++
								e.WriteString("null")
								return
//line /usr/local/go/src/encoding/json/encode.go:870
		// _ = "end of CoverTab[27734]"
	} else {
//line /usr/local/go/src/encoding/json/encode.go:871
		_go_fuzz_dep_.CoverTab[27735]++
//line /usr/local/go/src/encoding/json/encode.go:871
		// _ = "end of CoverTab[27735]"
//line /usr/local/go/src/encoding/json/encode.go:871
	}
//line /usr/local/go/src/encoding/json/encode.go:871
	// _ = "end of CoverTab[27731]"
//line /usr/local/go/src/encoding/json/encode.go:871
	_go_fuzz_dep_.CoverTab[27732]++
							if e.ptrLevel++; e.ptrLevel > startDetectingCyclesAfter {
//line /usr/local/go/src/encoding/json/encode.go:872
		_go_fuzz_dep_.CoverTab[27736]++

//line /usr/local/go/src/encoding/json/encode.go:877
		ptr := struct {
			ptr	interface{}	// always an unsafe.Pointer, but avoids a dependency on package unsafe
			len	int
		}{v.UnsafePointer(), v.Len()}
		if _, ok := e.ptrSeen[ptr]; ok {
//line /usr/local/go/src/encoding/json/encode.go:881
			_go_fuzz_dep_.CoverTab[27738]++
									e.error(&UnsupportedValueError{v, fmt.Sprintf("encountered a cycle via %s", v.Type())})
//line /usr/local/go/src/encoding/json/encode.go:882
			// _ = "end of CoverTab[27738]"
		} else {
//line /usr/local/go/src/encoding/json/encode.go:883
			_go_fuzz_dep_.CoverTab[27739]++
//line /usr/local/go/src/encoding/json/encode.go:883
			// _ = "end of CoverTab[27739]"
//line /usr/local/go/src/encoding/json/encode.go:883
		}
//line /usr/local/go/src/encoding/json/encode.go:883
		// _ = "end of CoverTab[27736]"
//line /usr/local/go/src/encoding/json/encode.go:883
		_go_fuzz_dep_.CoverTab[27737]++
								e.ptrSeen[ptr] = struct{}{}
								defer delete(e.ptrSeen, ptr)
//line /usr/local/go/src/encoding/json/encode.go:885
		// _ = "end of CoverTab[27737]"
	} else {
//line /usr/local/go/src/encoding/json/encode.go:886
		_go_fuzz_dep_.CoverTab[27740]++
//line /usr/local/go/src/encoding/json/encode.go:886
		// _ = "end of CoverTab[27740]"
//line /usr/local/go/src/encoding/json/encode.go:886
	}
//line /usr/local/go/src/encoding/json/encode.go:886
	// _ = "end of CoverTab[27732]"
//line /usr/local/go/src/encoding/json/encode.go:886
	_go_fuzz_dep_.CoverTab[27733]++
							se.arrayEnc(e, v, opts)
							e.ptrLevel--
//line /usr/local/go/src/encoding/json/encode.go:888
	// _ = "end of CoverTab[27733]"
}

func newSliceEncoder(t reflect.Type) encoderFunc {
//line /usr/local/go/src/encoding/json/encode.go:891
	_go_fuzz_dep_.CoverTab[27741]++

							if t.Elem().Kind() == reflect.Uint8 {
//line /usr/local/go/src/encoding/json/encode.go:893
		_go_fuzz_dep_.CoverTab[27743]++
								p := reflect.PointerTo(t.Elem())
								if !p.Implements(marshalerType) && func() bool {
//line /usr/local/go/src/encoding/json/encode.go:895
			_go_fuzz_dep_.CoverTab[27744]++
//line /usr/local/go/src/encoding/json/encode.go:895
			return !p.Implements(textMarshalerType)
//line /usr/local/go/src/encoding/json/encode.go:895
			// _ = "end of CoverTab[27744]"
//line /usr/local/go/src/encoding/json/encode.go:895
		}() {
//line /usr/local/go/src/encoding/json/encode.go:895
			_go_fuzz_dep_.CoverTab[27745]++
									return encodeByteSlice
//line /usr/local/go/src/encoding/json/encode.go:896
			// _ = "end of CoverTab[27745]"
		} else {
//line /usr/local/go/src/encoding/json/encode.go:897
			_go_fuzz_dep_.CoverTab[27746]++
//line /usr/local/go/src/encoding/json/encode.go:897
			// _ = "end of CoverTab[27746]"
//line /usr/local/go/src/encoding/json/encode.go:897
		}
//line /usr/local/go/src/encoding/json/encode.go:897
		// _ = "end of CoverTab[27743]"
	} else {
//line /usr/local/go/src/encoding/json/encode.go:898
		_go_fuzz_dep_.CoverTab[27747]++
//line /usr/local/go/src/encoding/json/encode.go:898
		// _ = "end of CoverTab[27747]"
//line /usr/local/go/src/encoding/json/encode.go:898
	}
//line /usr/local/go/src/encoding/json/encode.go:898
	// _ = "end of CoverTab[27741]"
//line /usr/local/go/src/encoding/json/encode.go:898
	_go_fuzz_dep_.CoverTab[27742]++
							enc := sliceEncoder{newArrayEncoder(t)}
							return enc.encode
//line /usr/local/go/src/encoding/json/encode.go:900
	// _ = "end of CoverTab[27742]"
}

type arrayEncoder struct {
	elemEnc encoderFunc
}

func (ae arrayEncoder) encode(e *encodeState, v reflect.Value, opts encOpts) {
//line /usr/local/go/src/encoding/json/encode.go:907
	_go_fuzz_dep_.CoverTab[27748]++
							e.WriteByte('[')
							n := v.Len()
							for i := 0; i < n; i++ {
//line /usr/local/go/src/encoding/json/encode.go:910
		_go_fuzz_dep_.CoverTab[27750]++
								if i > 0 {
//line /usr/local/go/src/encoding/json/encode.go:911
			_go_fuzz_dep_.CoverTab[27752]++
									e.WriteByte(',')
//line /usr/local/go/src/encoding/json/encode.go:912
			// _ = "end of CoverTab[27752]"
		} else {
//line /usr/local/go/src/encoding/json/encode.go:913
			_go_fuzz_dep_.CoverTab[27753]++
//line /usr/local/go/src/encoding/json/encode.go:913
			// _ = "end of CoverTab[27753]"
//line /usr/local/go/src/encoding/json/encode.go:913
		}
//line /usr/local/go/src/encoding/json/encode.go:913
		// _ = "end of CoverTab[27750]"
//line /usr/local/go/src/encoding/json/encode.go:913
		_go_fuzz_dep_.CoverTab[27751]++
								ae.elemEnc(e, v.Index(i), opts)
//line /usr/local/go/src/encoding/json/encode.go:914
		// _ = "end of CoverTab[27751]"
	}
//line /usr/local/go/src/encoding/json/encode.go:915
	// _ = "end of CoverTab[27748]"
//line /usr/local/go/src/encoding/json/encode.go:915
	_go_fuzz_dep_.CoverTab[27749]++
							e.WriteByte(']')
//line /usr/local/go/src/encoding/json/encode.go:916
	// _ = "end of CoverTab[27749]"
}

func newArrayEncoder(t reflect.Type) encoderFunc {
//line /usr/local/go/src/encoding/json/encode.go:919
	_go_fuzz_dep_.CoverTab[27754]++
							enc := arrayEncoder{typeEncoder(t.Elem())}
							return enc.encode
//line /usr/local/go/src/encoding/json/encode.go:921
	// _ = "end of CoverTab[27754]"
}

type ptrEncoder struct {
	elemEnc encoderFunc
}

func (pe ptrEncoder) encode(e *encodeState, v reflect.Value, opts encOpts) {
//line /usr/local/go/src/encoding/json/encode.go:928
	_go_fuzz_dep_.CoverTab[27755]++
							if v.IsNil() {
//line /usr/local/go/src/encoding/json/encode.go:929
		_go_fuzz_dep_.CoverTab[27758]++
								e.WriteString("null")
								return
//line /usr/local/go/src/encoding/json/encode.go:931
		// _ = "end of CoverTab[27758]"
	} else {
//line /usr/local/go/src/encoding/json/encode.go:932
		_go_fuzz_dep_.CoverTab[27759]++
//line /usr/local/go/src/encoding/json/encode.go:932
		// _ = "end of CoverTab[27759]"
//line /usr/local/go/src/encoding/json/encode.go:932
	}
//line /usr/local/go/src/encoding/json/encode.go:932
	// _ = "end of CoverTab[27755]"
//line /usr/local/go/src/encoding/json/encode.go:932
	_go_fuzz_dep_.CoverTab[27756]++
							if e.ptrLevel++; e.ptrLevel > startDetectingCyclesAfter {
//line /usr/local/go/src/encoding/json/encode.go:933
		_go_fuzz_dep_.CoverTab[27760]++

//line /usr/local/go/src/encoding/json/encode.go:936
		ptr := v.Interface()
		if _, ok := e.ptrSeen[ptr]; ok {
//line /usr/local/go/src/encoding/json/encode.go:937
			_go_fuzz_dep_.CoverTab[27762]++
									e.error(&UnsupportedValueError{v, fmt.Sprintf("encountered a cycle via %s", v.Type())})
//line /usr/local/go/src/encoding/json/encode.go:938
			// _ = "end of CoverTab[27762]"
		} else {
//line /usr/local/go/src/encoding/json/encode.go:939
			_go_fuzz_dep_.CoverTab[27763]++
//line /usr/local/go/src/encoding/json/encode.go:939
			// _ = "end of CoverTab[27763]"
//line /usr/local/go/src/encoding/json/encode.go:939
		}
//line /usr/local/go/src/encoding/json/encode.go:939
		// _ = "end of CoverTab[27760]"
//line /usr/local/go/src/encoding/json/encode.go:939
		_go_fuzz_dep_.CoverTab[27761]++
								e.ptrSeen[ptr] = struct{}{}
								defer delete(e.ptrSeen, ptr)
//line /usr/local/go/src/encoding/json/encode.go:941
		// _ = "end of CoverTab[27761]"
	} else {
//line /usr/local/go/src/encoding/json/encode.go:942
		_go_fuzz_dep_.CoverTab[27764]++
//line /usr/local/go/src/encoding/json/encode.go:942
		// _ = "end of CoverTab[27764]"
//line /usr/local/go/src/encoding/json/encode.go:942
	}
//line /usr/local/go/src/encoding/json/encode.go:942
	// _ = "end of CoverTab[27756]"
//line /usr/local/go/src/encoding/json/encode.go:942
	_go_fuzz_dep_.CoverTab[27757]++
							pe.elemEnc(e, v.Elem(), opts)
							e.ptrLevel--
//line /usr/local/go/src/encoding/json/encode.go:944
	// _ = "end of CoverTab[27757]"
}

func newPtrEncoder(t reflect.Type) encoderFunc {
//line /usr/local/go/src/encoding/json/encode.go:947
	_go_fuzz_dep_.CoverTab[27765]++
							enc := ptrEncoder{typeEncoder(t.Elem())}
							return enc.encode
//line /usr/local/go/src/encoding/json/encode.go:949
	// _ = "end of CoverTab[27765]"
}

type condAddrEncoder struct {
	canAddrEnc, elseEnc encoderFunc
}

func (ce condAddrEncoder) encode(e *encodeState, v reflect.Value, opts encOpts) {
//line /usr/local/go/src/encoding/json/encode.go:956
	_go_fuzz_dep_.CoverTab[27766]++
							if v.CanAddr() {
//line /usr/local/go/src/encoding/json/encode.go:957
		_go_fuzz_dep_.CoverTab[27767]++
								ce.canAddrEnc(e, v, opts)
//line /usr/local/go/src/encoding/json/encode.go:958
		// _ = "end of CoverTab[27767]"
	} else {
//line /usr/local/go/src/encoding/json/encode.go:959
		_go_fuzz_dep_.CoverTab[27768]++
								ce.elseEnc(e, v, opts)
//line /usr/local/go/src/encoding/json/encode.go:960
		// _ = "end of CoverTab[27768]"
	}
//line /usr/local/go/src/encoding/json/encode.go:961
	// _ = "end of CoverTab[27766]"
}

// newCondAddrEncoder returns an encoder that checks whether its value
//line /usr/local/go/src/encoding/json/encode.go:964
// CanAddr and delegates to canAddrEnc if so, else to elseEnc.
//line /usr/local/go/src/encoding/json/encode.go:966
func newCondAddrEncoder(canAddrEnc, elseEnc encoderFunc) encoderFunc {
//line /usr/local/go/src/encoding/json/encode.go:966
	_go_fuzz_dep_.CoverTab[27769]++
							enc := condAddrEncoder{canAddrEnc: canAddrEnc, elseEnc: elseEnc}
							return enc.encode
//line /usr/local/go/src/encoding/json/encode.go:968
	// _ = "end of CoverTab[27769]"
}

func isValidTag(s string) bool {
//line /usr/local/go/src/encoding/json/encode.go:971
	_go_fuzz_dep_.CoverTab[27770]++
							if s == "" {
//line /usr/local/go/src/encoding/json/encode.go:972
		_go_fuzz_dep_.CoverTab[27773]++
								return false
//line /usr/local/go/src/encoding/json/encode.go:973
		// _ = "end of CoverTab[27773]"
	} else {
//line /usr/local/go/src/encoding/json/encode.go:974
		_go_fuzz_dep_.CoverTab[27774]++
//line /usr/local/go/src/encoding/json/encode.go:974
		// _ = "end of CoverTab[27774]"
//line /usr/local/go/src/encoding/json/encode.go:974
	}
//line /usr/local/go/src/encoding/json/encode.go:974
	// _ = "end of CoverTab[27770]"
//line /usr/local/go/src/encoding/json/encode.go:974
	_go_fuzz_dep_.CoverTab[27771]++
							for _, c := range s {
//line /usr/local/go/src/encoding/json/encode.go:975
		_go_fuzz_dep_.CoverTab[27775]++
								switch {
		case strings.ContainsRune("!#$%&()*+-./:;<=>?@[]^_{|}~ ", c):
//line /usr/local/go/src/encoding/json/encode.go:977
			_go_fuzz_dep_.CoverTab[27776]++
//line /usr/local/go/src/encoding/json/encode.go:977
			// _ = "end of CoverTab[27776]"

//line /usr/local/go/src/encoding/json/encode.go:981
		case !unicode.IsLetter(c) && func() bool {
//line /usr/local/go/src/encoding/json/encode.go:981
			_go_fuzz_dep_.CoverTab[27779]++
//line /usr/local/go/src/encoding/json/encode.go:981
			return !unicode.IsDigit(c)
//line /usr/local/go/src/encoding/json/encode.go:981
			// _ = "end of CoverTab[27779]"
//line /usr/local/go/src/encoding/json/encode.go:981
		}():
//line /usr/local/go/src/encoding/json/encode.go:981
			_go_fuzz_dep_.CoverTab[27777]++
									return false
//line /usr/local/go/src/encoding/json/encode.go:982
			// _ = "end of CoverTab[27777]"
//line /usr/local/go/src/encoding/json/encode.go:982
		default:
//line /usr/local/go/src/encoding/json/encode.go:982
			_go_fuzz_dep_.CoverTab[27778]++
//line /usr/local/go/src/encoding/json/encode.go:982
			// _ = "end of CoverTab[27778]"
		}
//line /usr/local/go/src/encoding/json/encode.go:983
		// _ = "end of CoverTab[27775]"
	}
//line /usr/local/go/src/encoding/json/encode.go:984
	// _ = "end of CoverTab[27771]"
//line /usr/local/go/src/encoding/json/encode.go:984
	_go_fuzz_dep_.CoverTab[27772]++
							return true
//line /usr/local/go/src/encoding/json/encode.go:985
	// _ = "end of CoverTab[27772]"
}

func typeByIndex(t reflect.Type, index []int) reflect.Type {
//line /usr/local/go/src/encoding/json/encode.go:988
	_go_fuzz_dep_.CoverTab[27780]++
							for _, i := range index {
//line /usr/local/go/src/encoding/json/encode.go:989
		_go_fuzz_dep_.CoverTab[27782]++
								if t.Kind() == reflect.Pointer {
//line /usr/local/go/src/encoding/json/encode.go:990
			_go_fuzz_dep_.CoverTab[27784]++
									t = t.Elem()
//line /usr/local/go/src/encoding/json/encode.go:991
			// _ = "end of CoverTab[27784]"
		} else {
//line /usr/local/go/src/encoding/json/encode.go:992
			_go_fuzz_dep_.CoverTab[27785]++
//line /usr/local/go/src/encoding/json/encode.go:992
			// _ = "end of CoverTab[27785]"
//line /usr/local/go/src/encoding/json/encode.go:992
		}
//line /usr/local/go/src/encoding/json/encode.go:992
		// _ = "end of CoverTab[27782]"
//line /usr/local/go/src/encoding/json/encode.go:992
		_go_fuzz_dep_.CoverTab[27783]++
								t = t.Field(i).Type
//line /usr/local/go/src/encoding/json/encode.go:993
		// _ = "end of CoverTab[27783]"
	}
//line /usr/local/go/src/encoding/json/encode.go:994
	// _ = "end of CoverTab[27780]"
//line /usr/local/go/src/encoding/json/encode.go:994
	_go_fuzz_dep_.CoverTab[27781]++
							return t
//line /usr/local/go/src/encoding/json/encode.go:995
	// _ = "end of CoverTab[27781]"
}

type reflectWithString struct {
	k	reflect.Value
	v	reflect.Value
	ks	string
}

func (w *reflectWithString) resolve() error {
//line /usr/local/go/src/encoding/json/encode.go:1004
	_go_fuzz_dep_.CoverTab[27786]++
							if w.k.Kind() == reflect.String {
//line /usr/local/go/src/encoding/json/encode.go:1005
		_go_fuzz_dep_.CoverTab[27790]++
								w.ks = w.k.String()
								return nil
//line /usr/local/go/src/encoding/json/encode.go:1007
		// _ = "end of CoverTab[27790]"
	} else {
//line /usr/local/go/src/encoding/json/encode.go:1008
		_go_fuzz_dep_.CoverTab[27791]++
//line /usr/local/go/src/encoding/json/encode.go:1008
		// _ = "end of CoverTab[27791]"
//line /usr/local/go/src/encoding/json/encode.go:1008
	}
//line /usr/local/go/src/encoding/json/encode.go:1008
	// _ = "end of CoverTab[27786]"
//line /usr/local/go/src/encoding/json/encode.go:1008
	_go_fuzz_dep_.CoverTab[27787]++
							if tm, ok := w.k.Interface().(encoding.TextMarshaler); ok {
//line /usr/local/go/src/encoding/json/encode.go:1009
		_go_fuzz_dep_.CoverTab[27792]++
								if w.k.Kind() == reflect.Pointer && func() bool {
//line /usr/local/go/src/encoding/json/encode.go:1010
			_go_fuzz_dep_.CoverTab[27794]++
//line /usr/local/go/src/encoding/json/encode.go:1010
			return w.k.IsNil()
//line /usr/local/go/src/encoding/json/encode.go:1010
			// _ = "end of CoverTab[27794]"
//line /usr/local/go/src/encoding/json/encode.go:1010
		}() {
//line /usr/local/go/src/encoding/json/encode.go:1010
			_go_fuzz_dep_.CoverTab[27795]++
									return nil
//line /usr/local/go/src/encoding/json/encode.go:1011
			// _ = "end of CoverTab[27795]"
		} else {
//line /usr/local/go/src/encoding/json/encode.go:1012
			_go_fuzz_dep_.CoverTab[27796]++
//line /usr/local/go/src/encoding/json/encode.go:1012
			// _ = "end of CoverTab[27796]"
//line /usr/local/go/src/encoding/json/encode.go:1012
		}
//line /usr/local/go/src/encoding/json/encode.go:1012
		// _ = "end of CoverTab[27792]"
//line /usr/local/go/src/encoding/json/encode.go:1012
		_go_fuzz_dep_.CoverTab[27793]++
								buf, err := tm.MarshalText()
								w.ks = string(buf)
								return err
//line /usr/local/go/src/encoding/json/encode.go:1015
		// _ = "end of CoverTab[27793]"
	} else {
//line /usr/local/go/src/encoding/json/encode.go:1016
		_go_fuzz_dep_.CoverTab[27797]++
//line /usr/local/go/src/encoding/json/encode.go:1016
		// _ = "end of CoverTab[27797]"
//line /usr/local/go/src/encoding/json/encode.go:1016
	}
//line /usr/local/go/src/encoding/json/encode.go:1016
	// _ = "end of CoverTab[27787]"
//line /usr/local/go/src/encoding/json/encode.go:1016
	_go_fuzz_dep_.CoverTab[27788]++
							switch w.k.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
//line /usr/local/go/src/encoding/json/encode.go:1018
		_go_fuzz_dep_.CoverTab[27798]++
								w.ks = strconv.FormatInt(w.k.Int(), 10)
								return nil
//line /usr/local/go/src/encoding/json/encode.go:1020
		// _ = "end of CoverTab[27798]"
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
//line /usr/local/go/src/encoding/json/encode.go:1021
		_go_fuzz_dep_.CoverTab[27799]++
								w.ks = strconv.FormatUint(w.k.Uint(), 10)
								return nil
//line /usr/local/go/src/encoding/json/encode.go:1023
		// _ = "end of CoverTab[27799]"
//line /usr/local/go/src/encoding/json/encode.go:1023
	default:
//line /usr/local/go/src/encoding/json/encode.go:1023
		_go_fuzz_dep_.CoverTab[27800]++
//line /usr/local/go/src/encoding/json/encode.go:1023
		// _ = "end of CoverTab[27800]"
	}
//line /usr/local/go/src/encoding/json/encode.go:1024
	// _ = "end of CoverTab[27788]"
//line /usr/local/go/src/encoding/json/encode.go:1024
	_go_fuzz_dep_.CoverTab[27789]++
							panic("unexpected map key type")
//line /usr/local/go/src/encoding/json/encode.go:1025
	// _ = "end of CoverTab[27789]"
}

// NOTE: keep in sync with stringBytes below.
func (e *encodeState) string(s string, escapeHTML bool) {
//line /usr/local/go/src/encoding/json/encode.go:1029
	_go_fuzz_dep_.CoverTab[27801]++
							e.WriteByte('"')
							start := 0
							for i := 0; i < len(s); {
//line /usr/local/go/src/encoding/json/encode.go:1032
		_go_fuzz_dep_.CoverTab[27804]++
								if b := s[i]; b < utf8.RuneSelf {
//line /usr/local/go/src/encoding/json/encode.go:1033
			_go_fuzz_dep_.CoverTab[27808]++
									if htmlSafeSet[b] || func() bool {
//line /usr/local/go/src/encoding/json/encode.go:1034
				_go_fuzz_dep_.CoverTab[27812]++
//line /usr/local/go/src/encoding/json/encode.go:1034
				return (!escapeHTML && func() bool {
//line /usr/local/go/src/encoding/json/encode.go:1034
					_go_fuzz_dep_.CoverTab[27813]++
//line /usr/local/go/src/encoding/json/encode.go:1034
					return safeSet[b]
//line /usr/local/go/src/encoding/json/encode.go:1034
					// _ = "end of CoverTab[27813]"
//line /usr/local/go/src/encoding/json/encode.go:1034
				}())
//line /usr/local/go/src/encoding/json/encode.go:1034
				// _ = "end of CoverTab[27812]"
//line /usr/local/go/src/encoding/json/encode.go:1034
			}() {
//line /usr/local/go/src/encoding/json/encode.go:1034
				_go_fuzz_dep_.CoverTab[27814]++
										i++
										continue
//line /usr/local/go/src/encoding/json/encode.go:1036
				// _ = "end of CoverTab[27814]"
			} else {
//line /usr/local/go/src/encoding/json/encode.go:1037
				_go_fuzz_dep_.CoverTab[27815]++
//line /usr/local/go/src/encoding/json/encode.go:1037
				// _ = "end of CoverTab[27815]"
//line /usr/local/go/src/encoding/json/encode.go:1037
			}
//line /usr/local/go/src/encoding/json/encode.go:1037
			// _ = "end of CoverTab[27808]"
//line /usr/local/go/src/encoding/json/encode.go:1037
			_go_fuzz_dep_.CoverTab[27809]++
									if start < i {
//line /usr/local/go/src/encoding/json/encode.go:1038
				_go_fuzz_dep_.CoverTab[27816]++
										e.WriteString(s[start:i])
//line /usr/local/go/src/encoding/json/encode.go:1039
				// _ = "end of CoverTab[27816]"
			} else {
//line /usr/local/go/src/encoding/json/encode.go:1040
				_go_fuzz_dep_.CoverTab[27817]++
//line /usr/local/go/src/encoding/json/encode.go:1040
				// _ = "end of CoverTab[27817]"
//line /usr/local/go/src/encoding/json/encode.go:1040
			}
//line /usr/local/go/src/encoding/json/encode.go:1040
			// _ = "end of CoverTab[27809]"
//line /usr/local/go/src/encoding/json/encode.go:1040
			_go_fuzz_dep_.CoverTab[27810]++
									e.WriteByte('\\')
									switch b {
			case '\\', '"':
//line /usr/local/go/src/encoding/json/encode.go:1043
				_go_fuzz_dep_.CoverTab[27818]++
										e.WriteByte(b)
//line /usr/local/go/src/encoding/json/encode.go:1044
				// _ = "end of CoverTab[27818]"
			case '\n':
//line /usr/local/go/src/encoding/json/encode.go:1045
				_go_fuzz_dep_.CoverTab[27819]++
										e.WriteByte('n')
//line /usr/local/go/src/encoding/json/encode.go:1046
				// _ = "end of CoverTab[27819]"
			case '\r':
//line /usr/local/go/src/encoding/json/encode.go:1047
				_go_fuzz_dep_.CoverTab[27820]++
										e.WriteByte('r')
//line /usr/local/go/src/encoding/json/encode.go:1048
				// _ = "end of CoverTab[27820]"
			case '\t':
//line /usr/local/go/src/encoding/json/encode.go:1049
				_go_fuzz_dep_.CoverTab[27821]++
										e.WriteByte('t')
//line /usr/local/go/src/encoding/json/encode.go:1050
				// _ = "end of CoverTab[27821]"
			default:
//line /usr/local/go/src/encoding/json/encode.go:1051
				_go_fuzz_dep_.CoverTab[27822]++

//line /usr/local/go/src/encoding/json/encode.go:1057
				e.WriteString(`u00`)
										e.WriteByte(hex[b>>4])
										e.WriteByte(hex[b&0xF])
//line /usr/local/go/src/encoding/json/encode.go:1059
				// _ = "end of CoverTab[27822]"
			}
//line /usr/local/go/src/encoding/json/encode.go:1060
			// _ = "end of CoverTab[27810]"
//line /usr/local/go/src/encoding/json/encode.go:1060
			_go_fuzz_dep_.CoverTab[27811]++
									i++
									start = i
									continue
//line /usr/local/go/src/encoding/json/encode.go:1063
			// _ = "end of CoverTab[27811]"
		} else {
//line /usr/local/go/src/encoding/json/encode.go:1064
			_go_fuzz_dep_.CoverTab[27823]++
//line /usr/local/go/src/encoding/json/encode.go:1064
			// _ = "end of CoverTab[27823]"
//line /usr/local/go/src/encoding/json/encode.go:1064
		}
//line /usr/local/go/src/encoding/json/encode.go:1064
		// _ = "end of CoverTab[27804]"
//line /usr/local/go/src/encoding/json/encode.go:1064
		_go_fuzz_dep_.CoverTab[27805]++
								c, size := utf8.DecodeRuneInString(s[i:])
								if c == utf8.RuneError && func() bool {
//line /usr/local/go/src/encoding/json/encode.go:1066
			_go_fuzz_dep_.CoverTab[27824]++
//line /usr/local/go/src/encoding/json/encode.go:1066
			return size == 1
//line /usr/local/go/src/encoding/json/encode.go:1066
			// _ = "end of CoverTab[27824]"
//line /usr/local/go/src/encoding/json/encode.go:1066
		}() {
//line /usr/local/go/src/encoding/json/encode.go:1066
			_go_fuzz_dep_.CoverTab[27825]++
									if start < i {
//line /usr/local/go/src/encoding/json/encode.go:1067
				_go_fuzz_dep_.CoverTab[27827]++
										e.WriteString(s[start:i])
//line /usr/local/go/src/encoding/json/encode.go:1068
				// _ = "end of CoverTab[27827]"
			} else {
//line /usr/local/go/src/encoding/json/encode.go:1069
				_go_fuzz_dep_.CoverTab[27828]++
//line /usr/local/go/src/encoding/json/encode.go:1069
				// _ = "end of CoverTab[27828]"
//line /usr/local/go/src/encoding/json/encode.go:1069
			}
//line /usr/local/go/src/encoding/json/encode.go:1069
			// _ = "end of CoverTab[27825]"
//line /usr/local/go/src/encoding/json/encode.go:1069
			_go_fuzz_dep_.CoverTab[27826]++
									e.WriteString(`\ufffd`)
									i += size
									start = i
									continue
//line /usr/local/go/src/encoding/json/encode.go:1073
			// _ = "end of CoverTab[27826]"
		} else {
//line /usr/local/go/src/encoding/json/encode.go:1074
			_go_fuzz_dep_.CoverTab[27829]++
//line /usr/local/go/src/encoding/json/encode.go:1074
			// _ = "end of CoverTab[27829]"
//line /usr/local/go/src/encoding/json/encode.go:1074
		}
//line /usr/local/go/src/encoding/json/encode.go:1074
		// _ = "end of CoverTab[27805]"
//line /usr/local/go/src/encoding/json/encode.go:1074
		_go_fuzz_dep_.CoverTab[27806]++

//line /usr/local/go/src/encoding/json/encode.go:1082
		if c == '\u2028' || func() bool {
//line /usr/local/go/src/encoding/json/encode.go:1082
			_go_fuzz_dep_.CoverTab[27830]++
//line /usr/local/go/src/encoding/json/encode.go:1082
			return c == '\u2029'
//line /usr/local/go/src/encoding/json/encode.go:1082
			// _ = "end of CoverTab[27830]"
//line /usr/local/go/src/encoding/json/encode.go:1082
		}() {
//line /usr/local/go/src/encoding/json/encode.go:1082
			_go_fuzz_dep_.CoverTab[27831]++
									if start < i {
//line /usr/local/go/src/encoding/json/encode.go:1083
				_go_fuzz_dep_.CoverTab[27833]++
										e.WriteString(s[start:i])
//line /usr/local/go/src/encoding/json/encode.go:1084
				// _ = "end of CoverTab[27833]"
			} else {
//line /usr/local/go/src/encoding/json/encode.go:1085
				_go_fuzz_dep_.CoverTab[27834]++
//line /usr/local/go/src/encoding/json/encode.go:1085
				// _ = "end of CoverTab[27834]"
//line /usr/local/go/src/encoding/json/encode.go:1085
			}
//line /usr/local/go/src/encoding/json/encode.go:1085
			// _ = "end of CoverTab[27831]"
//line /usr/local/go/src/encoding/json/encode.go:1085
			_go_fuzz_dep_.CoverTab[27832]++
									e.WriteString(`\u202`)
									e.WriteByte(hex[c&0xF])
									i += size
									start = i
									continue
//line /usr/local/go/src/encoding/json/encode.go:1090
			// _ = "end of CoverTab[27832]"
		} else {
//line /usr/local/go/src/encoding/json/encode.go:1091
			_go_fuzz_dep_.CoverTab[27835]++
//line /usr/local/go/src/encoding/json/encode.go:1091
			// _ = "end of CoverTab[27835]"
//line /usr/local/go/src/encoding/json/encode.go:1091
		}
//line /usr/local/go/src/encoding/json/encode.go:1091
		// _ = "end of CoverTab[27806]"
//line /usr/local/go/src/encoding/json/encode.go:1091
		_go_fuzz_dep_.CoverTab[27807]++
								i += size
//line /usr/local/go/src/encoding/json/encode.go:1092
		// _ = "end of CoverTab[27807]"
	}
//line /usr/local/go/src/encoding/json/encode.go:1093
	// _ = "end of CoverTab[27801]"
//line /usr/local/go/src/encoding/json/encode.go:1093
	_go_fuzz_dep_.CoverTab[27802]++
							if start < len(s) {
//line /usr/local/go/src/encoding/json/encode.go:1094
		_go_fuzz_dep_.CoverTab[27836]++
								e.WriteString(s[start:])
//line /usr/local/go/src/encoding/json/encode.go:1095
		// _ = "end of CoverTab[27836]"
	} else {
//line /usr/local/go/src/encoding/json/encode.go:1096
		_go_fuzz_dep_.CoverTab[27837]++
//line /usr/local/go/src/encoding/json/encode.go:1096
		// _ = "end of CoverTab[27837]"
//line /usr/local/go/src/encoding/json/encode.go:1096
	}
//line /usr/local/go/src/encoding/json/encode.go:1096
	// _ = "end of CoverTab[27802]"
//line /usr/local/go/src/encoding/json/encode.go:1096
	_go_fuzz_dep_.CoverTab[27803]++
							e.WriteByte('"')
//line /usr/local/go/src/encoding/json/encode.go:1097
	// _ = "end of CoverTab[27803]"
}

// NOTE: keep in sync with string above.
func (e *encodeState) stringBytes(s []byte, escapeHTML bool) {
//line /usr/local/go/src/encoding/json/encode.go:1101
	_go_fuzz_dep_.CoverTab[27838]++
							e.WriteByte('"')
							start := 0
							for i := 0; i < len(s); {
//line /usr/local/go/src/encoding/json/encode.go:1104
		_go_fuzz_dep_.CoverTab[27841]++
								if b := s[i]; b < utf8.RuneSelf {
//line /usr/local/go/src/encoding/json/encode.go:1105
			_go_fuzz_dep_.CoverTab[27845]++
									if htmlSafeSet[b] || func() bool {
//line /usr/local/go/src/encoding/json/encode.go:1106
				_go_fuzz_dep_.CoverTab[27849]++
//line /usr/local/go/src/encoding/json/encode.go:1106
				return (!escapeHTML && func() bool {
//line /usr/local/go/src/encoding/json/encode.go:1106
					_go_fuzz_dep_.CoverTab[27850]++
//line /usr/local/go/src/encoding/json/encode.go:1106
					return safeSet[b]
//line /usr/local/go/src/encoding/json/encode.go:1106
					// _ = "end of CoverTab[27850]"
//line /usr/local/go/src/encoding/json/encode.go:1106
				}())
//line /usr/local/go/src/encoding/json/encode.go:1106
				// _ = "end of CoverTab[27849]"
//line /usr/local/go/src/encoding/json/encode.go:1106
			}() {
//line /usr/local/go/src/encoding/json/encode.go:1106
				_go_fuzz_dep_.CoverTab[27851]++
										i++
										continue
//line /usr/local/go/src/encoding/json/encode.go:1108
				// _ = "end of CoverTab[27851]"
			} else {
//line /usr/local/go/src/encoding/json/encode.go:1109
				_go_fuzz_dep_.CoverTab[27852]++
//line /usr/local/go/src/encoding/json/encode.go:1109
				// _ = "end of CoverTab[27852]"
//line /usr/local/go/src/encoding/json/encode.go:1109
			}
//line /usr/local/go/src/encoding/json/encode.go:1109
			// _ = "end of CoverTab[27845]"
//line /usr/local/go/src/encoding/json/encode.go:1109
			_go_fuzz_dep_.CoverTab[27846]++
									if start < i {
//line /usr/local/go/src/encoding/json/encode.go:1110
				_go_fuzz_dep_.CoverTab[27853]++
										e.Write(s[start:i])
//line /usr/local/go/src/encoding/json/encode.go:1111
				// _ = "end of CoverTab[27853]"
			} else {
//line /usr/local/go/src/encoding/json/encode.go:1112
				_go_fuzz_dep_.CoverTab[27854]++
//line /usr/local/go/src/encoding/json/encode.go:1112
				// _ = "end of CoverTab[27854]"
//line /usr/local/go/src/encoding/json/encode.go:1112
			}
//line /usr/local/go/src/encoding/json/encode.go:1112
			// _ = "end of CoverTab[27846]"
//line /usr/local/go/src/encoding/json/encode.go:1112
			_go_fuzz_dep_.CoverTab[27847]++
									e.WriteByte('\\')
									switch b {
			case '\\', '"':
//line /usr/local/go/src/encoding/json/encode.go:1115
				_go_fuzz_dep_.CoverTab[27855]++
										e.WriteByte(b)
//line /usr/local/go/src/encoding/json/encode.go:1116
				// _ = "end of CoverTab[27855]"
			case '\n':
//line /usr/local/go/src/encoding/json/encode.go:1117
				_go_fuzz_dep_.CoverTab[27856]++
										e.WriteByte('n')
//line /usr/local/go/src/encoding/json/encode.go:1118
				// _ = "end of CoverTab[27856]"
			case '\r':
//line /usr/local/go/src/encoding/json/encode.go:1119
				_go_fuzz_dep_.CoverTab[27857]++
										e.WriteByte('r')
//line /usr/local/go/src/encoding/json/encode.go:1120
				// _ = "end of CoverTab[27857]"
			case '\t':
//line /usr/local/go/src/encoding/json/encode.go:1121
				_go_fuzz_dep_.CoverTab[27858]++
										e.WriteByte('t')
//line /usr/local/go/src/encoding/json/encode.go:1122
				// _ = "end of CoverTab[27858]"
			default:
//line /usr/local/go/src/encoding/json/encode.go:1123
				_go_fuzz_dep_.CoverTab[27859]++

//line /usr/local/go/src/encoding/json/encode.go:1129
				e.WriteString(`u00`)
										e.WriteByte(hex[b>>4])
										e.WriteByte(hex[b&0xF])
//line /usr/local/go/src/encoding/json/encode.go:1131
				// _ = "end of CoverTab[27859]"
			}
//line /usr/local/go/src/encoding/json/encode.go:1132
			// _ = "end of CoverTab[27847]"
//line /usr/local/go/src/encoding/json/encode.go:1132
			_go_fuzz_dep_.CoverTab[27848]++
									i++
									start = i
									continue
//line /usr/local/go/src/encoding/json/encode.go:1135
			// _ = "end of CoverTab[27848]"
		} else {
//line /usr/local/go/src/encoding/json/encode.go:1136
			_go_fuzz_dep_.CoverTab[27860]++
//line /usr/local/go/src/encoding/json/encode.go:1136
			// _ = "end of CoverTab[27860]"
//line /usr/local/go/src/encoding/json/encode.go:1136
		}
//line /usr/local/go/src/encoding/json/encode.go:1136
		// _ = "end of CoverTab[27841]"
//line /usr/local/go/src/encoding/json/encode.go:1136
		_go_fuzz_dep_.CoverTab[27842]++
								c, size := utf8.DecodeRune(s[i:])
								if c == utf8.RuneError && func() bool {
//line /usr/local/go/src/encoding/json/encode.go:1138
			_go_fuzz_dep_.CoverTab[27861]++
//line /usr/local/go/src/encoding/json/encode.go:1138
			return size == 1
//line /usr/local/go/src/encoding/json/encode.go:1138
			// _ = "end of CoverTab[27861]"
//line /usr/local/go/src/encoding/json/encode.go:1138
		}() {
//line /usr/local/go/src/encoding/json/encode.go:1138
			_go_fuzz_dep_.CoverTab[27862]++
									if start < i {
//line /usr/local/go/src/encoding/json/encode.go:1139
				_go_fuzz_dep_.CoverTab[27864]++
										e.Write(s[start:i])
//line /usr/local/go/src/encoding/json/encode.go:1140
				// _ = "end of CoverTab[27864]"
			} else {
//line /usr/local/go/src/encoding/json/encode.go:1141
				_go_fuzz_dep_.CoverTab[27865]++
//line /usr/local/go/src/encoding/json/encode.go:1141
				// _ = "end of CoverTab[27865]"
//line /usr/local/go/src/encoding/json/encode.go:1141
			}
//line /usr/local/go/src/encoding/json/encode.go:1141
			// _ = "end of CoverTab[27862]"
//line /usr/local/go/src/encoding/json/encode.go:1141
			_go_fuzz_dep_.CoverTab[27863]++
									e.WriteString(`\ufffd`)
									i += size
									start = i
									continue
//line /usr/local/go/src/encoding/json/encode.go:1145
			// _ = "end of CoverTab[27863]"
		} else {
//line /usr/local/go/src/encoding/json/encode.go:1146
			_go_fuzz_dep_.CoverTab[27866]++
//line /usr/local/go/src/encoding/json/encode.go:1146
			// _ = "end of CoverTab[27866]"
//line /usr/local/go/src/encoding/json/encode.go:1146
		}
//line /usr/local/go/src/encoding/json/encode.go:1146
		// _ = "end of CoverTab[27842]"
//line /usr/local/go/src/encoding/json/encode.go:1146
		_go_fuzz_dep_.CoverTab[27843]++

//line /usr/local/go/src/encoding/json/encode.go:1154
		if c == '\u2028' || func() bool {
//line /usr/local/go/src/encoding/json/encode.go:1154
			_go_fuzz_dep_.CoverTab[27867]++
//line /usr/local/go/src/encoding/json/encode.go:1154
			return c == '\u2029'
//line /usr/local/go/src/encoding/json/encode.go:1154
			// _ = "end of CoverTab[27867]"
//line /usr/local/go/src/encoding/json/encode.go:1154
		}() {
//line /usr/local/go/src/encoding/json/encode.go:1154
			_go_fuzz_dep_.CoverTab[27868]++
									if start < i {
//line /usr/local/go/src/encoding/json/encode.go:1155
				_go_fuzz_dep_.CoverTab[27870]++
										e.Write(s[start:i])
//line /usr/local/go/src/encoding/json/encode.go:1156
				// _ = "end of CoverTab[27870]"
			} else {
//line /usr/local/go/src/encoding/json/encode.go:1157
				_go_fuzz_dep_.CoverTab[27871]++
//line /usr/local/go/src/encoding/json/encode.go:1157
				// _ = "end of CoverTab[27871]"
//line /usr/local/go/src/encoding/json/encode.go:1157
			}
//line /usr/local/go/src/encoding/json/encode.go:1157
			// _ = "end of CoverTab[27868]"
//line /usr/local/go/src/encoding/json/encode.go:1157
			_go_fuzz_dep_.CoverTab[27869]++
									e.WriteString(`\u202`)
									e.WriteByte(hex[c&0xF])
									i += size
									start = i
									continue
//line /usr/local/go/src/encoding/json/encode.go:1162
			// _ = "end of CoverTab[27869]"
		} else {
//line /usr/local/go/src/encoding/json/encode.go:1163
			_go_fuzz_dep_.CoverTab[27872]++
//line /usr/local/go/src/encoding/json/encode.go:1163
			// _ = "end of CoverTab[27872]"
//line /usr/local/go/src/encoding/json/encode.go:1163
		}
//line /usr/local/go/src/encoding/json/encode.go:1163
		// _ = "end of CoverTab[27843]"
//line /usr/local/go/src/encoding/json/encode.go:1163
		_go_fuzz_dep_.CoverTab[27844]++
								i += size
//line /usr/local/go/src/encoding/json/encode.go:1164
		// _ = "end of CoverTab[27844]"
	}
//line /usr/local/go/src/encoding/json/encode.go:1165
	// _ = "end of CoverTab[27838]"
//line /usr/local/go/src/encoding/json/encode.go:1165
	_go_fuzz_dep_.CoverTab[27839]++
							if start < len(s) {
//line /usr/local/go/src/encoding/json/encode.go:1166
		_go_fuzz_dep_.CoverTab[27873]++
								e.Write(s[start:])
//line /usr/local/go/src/encoding/json/encode.go:1167
		// _ = "end of CoverTab[27873]"
	} else {
//line /usr/local/go/src/encoding/json/encode.go:1168
		_go_fuzz_dep_.CoverTab[27874]++
//line /usr/local/go/src/encoding/json/encode.go:1168
		// _ = "end of CoverTab[27874]"
//line /usr/local/go/src/encoding/json/encode.go:1168
	}
//line /usr/local/go/src/encoding/json/encode.go:1168
	// _ = "end of CoverTab[27839]"
//line /usr/local/go/src/encoding/json/encode.go:1168
	_go_fuzz_dep_.CoverTab[27840]++
							e.WriteByte('"')
//line /usr/local/go/src/encoding/json/encode.go:1169
	// _ = "end of CoverTab[27840]"
}

// A field represents a single field found in a struct.
type field struct {
	name		string
	nameBytes	[]byte			// []byte(name)
	equalFold	func(s, t []byte) bool	// bytes.EqualFold or equivalent

	nameNonEsc	string	// `"` + name + `":`
	nameEscHTML	string	// `"` + HTMLEscape(name) + `":`

	tag		bool
	index		[]int
	typ		reflect.Type
	omitEmpty	bool
	quoted		bool

	encoder	encoderFunc
}

// byIndex sorts field by index sequence.
type byIndex []field

func (x byIndex) Len() int {
//line /usr/local/go/src/encoding/json/encode.go:1193
	_go_fuzz_dep_.CoverTab[27875]++
//line /usr/local/go/src/encoding/json/encode.go:1193
	return len(x)
//line /usr/local/go/src/encoding/json/encode.go:1193
	// _ = "end of CoverTab[27875]"
//line /usr/local/go/src/encoding/json/encode.go:1193
}

func (x byIndex) Swap(i, j int) {
//line /usr/local/go/src/encoding/json/encode.go:1195
	_go_fuzz_dep_.CoverTab[27876]++
//line /usr/local/go/src/encoding/json/encode.go:1195
	x[i], x[j] = x[j], x[i]
//line /usr/local/go/src/encoding/json/encode.go:1195
	// _ = "end of CoverTab[27876]"
//line /usr/local/go/src/encoding/json/encode.go:1195
}

func (x byIndex) Less(i, j int) bool {
//line /usr/local/go/src/encoding/json/encode.go:1197
	_go_fuzz_dep_.CoverTab[27877]++
							for k, xik := range x[i].index {
//line /usr/local/go/src/encoding/json/encode.go:1198
		_go_fuzz_dep_.CoverTab[27879]++
								if k >= len(x[j].index) {
//line /usr/local/go/src/encoding/json/encode.go:1199
			_go_fuzz_dep_.CoverTab[27881]++
									return false
//line /usr/local/go/src/encoding/json/encode.go:1200
			// _ = "end of CoverTab[27881]"
		} else {
//line /usr/local/go/src/encoding/json/encode.go:1201
			_go_fuzz_dep_.CoverTab[27882]++
//line /usr/local/go/src/encoding/json/encode.go:1201
			// _ = "end of CoverTab[27882]"
//line /usr/local/go/src/encoding/json/encode.go:1201
		}
//line /usr/local/go/src/encoding/json/encode.go:1201
		// _ = "end of CoverTab[27879]"
//line /usr/local/go/src/encoding/json/encode.go:1201
		_go_fuzz_dep_.CoverTab[27880]++
								if xik != x[j].index[k] {
//line /usr/local/go/src/encoding/json/encode.go:1202
			_go_fuzz_dep_.CoverTab[27883]++
									return xik < x[j].index[k]
//line /usr/local/go/src/encoding/json/encode.go:1203
			// _ = "end of CoverTab[27883]"
		} else {
//line /usr/local/go/src/encoding/json/encode.go:1204
			_go_fuzz_dep_.CoverTab[27884]++
//line /usr/local/go/src/encoding/json/encode.go:1204
			// _ = "end of CoverTab[27884]"
//line /usr/local/go/src/encoding/json/encode.go:1204
		}
//line /usr/local/go/src/encoding/json/encode.go:1204
		// _ = "end of CoverTab[27880]"
	}
//line /usr/local/go/src/encoding/json/encode.go:1205
	// _ = "end of CoverTab[27877]"
//line /usr/local/go/src/encoding/json/encode.go:1205
	_go_fuzz_dep_.CoverTab[27878]++
							return len(x[i].index) < len(x[j].index)
//line /usr/local/go/src/encoding/json/encode.go:1206
	// _ = "end of CoverTab[27878]"
}

// typeFields returns a list of fields that JSON should recognize for the given type.
//line /usr/local/go/src/encoding/json/encode.go:1209
// The algorithm is breadth-first search over the set of structs to include - the top struct
//line /usr/local/go/src/encoding/json/encode.go:1209
// and then any reachable anonymous structs.
//line /usr/local/go/src/encoding/json/encode.go:1212
func typeFields(t reflect.Type) structFields {
//line /usr/local/go/src/encoding/json/encode.go:1212
	_go_fuzz_dep_.CoverTab[27885]++

							current := []field{}
							next := []field{{typ: t}}

							// Count of queued names for current level and the next.
							var count, nextCount map[reflect.Type]int

//line /usr/local/go/src/encoding/json/encode.go:1221
	visited := map[reflect.Type]bool{}

	// Fields found.
	var fields []field

	// Buffer to run HTMLEscape on field names.
	var nameEscBuf bytes.Buffer

	for len(next) > 0 {
//line /usr/local/go/src/encoding/json/encode.go:1229
		_go_fuzz_dep_.CoverTab[27891]++
								current, next = next, current[:0]
								count, nextCount = nextCount, map[reflect.Type]int{}

								for _, f := range current {
//line /usr/local/go/src/encoding/json/encode.go:1233
			_go_fuzz_dep_.CoverTab[27892]++
									if visited[f.typ] {
//line /usr/local/go/src/encoding/json/encode.go:1234
				_go_fuzz_dep_.CoverTab[27894]++
										continue
//line /usr/local/go/src/encoding/json/encode.go:1235
				// _ = "end of CoverTab[27894]"
			} else {
//line /usr/local/go/src/encoding/json/encode.go:1236
				_go_fuzz_dep_.CoverTab[27895]++
//line /usr/local/go/src/encoding/json/encode.go:1236
				// _ = "end of CoverTab[27895]"
//line /usr/local/go/src/encoding/json/encode.go:1236
			}
//line /usr/local/go/src/encoding/json/encode.go:1236
			// _ = "end of CoverTab[27892]"
//line /usr/local/go/src/encoding/json/encode.go:1236
			_go_fuzz_dep_.CoverTab[27893]++
									visited[f.typ] = true

//line /usr/local/go/src/encoding/json/encode.go:1240
			for i := 0; i < f.typ.NumField(); i++ {
//line /usr/local/go/src/encoding/json/encode.go:1240
				_go_fuzz_dep_.CoverTab[27896]++
										sf := f.typ.Field(i)
										if sf.Anonymous {
//line /usr/local/go/src/encoding/json/encode.go:1242
					_go_fuzz_dep_.CoverTab[27903]++
											t := sf.Type
											if t.Kind() == reflect.Pointer {
//line /usr/local/go/src/encoding/json/encode.go:1244
						_go_fuzz_dep_.CoverTab[27905]++
												t = t.Elem()
//line /usr/local/go/src/encoding/json/encode.go:1245
						// _ = "end of CoverTab[27905]"
					} else {
//line /usr/local/go/src/encoding/json/encode.go:1246
						_go_fuzz_dep_.CoverTab[27906]++
//line /usr/local/go/src/encoding/json/encode.go:1246
						// _ = "end of CoverTab[27906]"
//line /usr/local/go/src/encoding/json/encode.go:1246
					}
//line /usr/local/go/src/encoding/json/encode.go:1246
					// _ = "end of CoverTab[27903]"
//line /usr/local/go/src/encoding/json/encode.go:1246
					_go_fuzz_dep_.CoverTab[27904]++
											if !sf.IsExported() && func() bool {
//line /usr/local/go/src/encoding/json/encode.go:1247
						_go_fuzz_dep_.CoverTab[27907]++
//line /usr/local/go/src/encoding/json/encode.go:1247
						return t.Kind() != reflect.Struct
//line /usr/local/go/src/encoding/json/encode.go:1247
						// _ = "end of CoverTab[27907]"
//line /usr/local/go/src/encoding/json/encode.go:1247
					}() {
//line /usr/local/go/src/encoding/json/encode.go:1247
						_go_fuzz_dep_.CoverTab[27908]++

												continue
//line /usr/local/go/src/encoding/json/encode.go:1249
						// _ = "end of CoverTab[27908]"
					} else {
//line /usr/local/go/src/encoding/json/encode.go:1250
						_go_fuzz_dep_.CoverTab[27909]++
//line /usr/local/go/src/encoding/json/encode.go:1250
						// _ = "end of CoverTab[27909]"
//line /usr/local/go/src/encoding/json/encode.go:1250
					}
//line /usr/local/go/src/encoding/json/encode.go:1250
					// _ = "end of CoverTab[27904]"

//line /usr/local/go/src/encoding/json/encode.go:1253
				} else {
//line /usr/local/go/src/encoding/json/encode.go:1253
					_go_fuzz_dep_.CoverTab[27910]++
//line /usr/local/go/src/encoding/json/encode.go:1253
					if !sf.IsExported() {
//line /usr/local/go/src/encoding/json/encode.go:1253
						_go_fuzz_dep_.CoverTab[27911]++

												continue
//line /usr/local/go/src/encoding/json/encode.go:1255
						// _ = "end of CoverTab[27911]"
					} else {
//line /usr/local/go/src/encoding/json/encode.go:1256
						_go_fuzz_dep_.CoverTab[27912]++
//line /usr/local/go/src/encoding/json/encode.go:1256
						// _ = "end of CoverTab[27912]"
//line /usr/local/go/src/encoding/json/encode.go:1256
					}
//line /usr/local/go/src/encoding/json/encode.go:1256
					// _ = "end of CoverTab[27910]"
//line /usr/local/go/src/encoding/json/encode.go:1256
				}
//line /usr/local/go/src/encoding/json/encode.go:1256
				// _ = "end of CoverTab[27896]"
//line /usr/local/go/src/encoding/json/encode.go:1256
				_go_fuzz_dep_.CoverTab[27897]++
										tag := sf.Tag.Get("json")
										if tag == "-" {
//line /usr/local/go/src/encoding/json/encode.go:1258
					_go_fuzz_dep_.CoverTab[27913]++
											continue
//line /usr/local/go/src/encoding/json/encode.go:1259
					// _ = "end of CoverTab[27913]"
				} else {
//line /usr/local/go/src/encoding/json/encode.go:1260
					_go_fuzz_dep_.CoverTab[27914]++
//line /usr/local/go/src/encoding/json/encode.go:1260
					// _ = "end of CoverTab[27914]"
//line /usr/local/go/src/encoding/json/encode.go:1260
				}
//line /usr/local/go/src/encoding/json/encode.go:1260
				// _ = "end of CoverTab[27897]"
//line /usr/local/go/src/encoding/json/encode.go:1260
				_go_fuzz_dep_.CoverTab[27898]++
										name, opts := parseTag(tag)
										if !isValidTag(name) {
//line /usr/local/go/src/encoding/json/encode.go:1262
					_go_fuzz_dep_.CoverTab[27915]++
											name = ""
//line /usr/local/go/src/encoding/json/encode.go:1263
					// _ = "end of CoverTab[27915]"
				} else {
//line /usr/local/go/src/encoding/json/encode.go:1264
					_go_fuzz_dep_.CoverTab[27916]++
//line /usr/local/go/src/encoding/json/encode.go:1264
					// _ = "end of CoverTab[27916]"
//line /usr/local/go/src/encoding/json/encode.go:1264
				}
//line /usr/local/go/src/encoding/json/encode.go:1264
				// _ = "end of CoverTab[27898]"
//line /usr/local/go/src/encoding/json/encode.go:1264
				_go_fuzz_dep_.CoverTab[27899]++
										index := make([]int, len(f.index)+1)
										copy(index, f.index)
										index[len(f.index)] = i

										ft := sf.Type
										if ft.Name() == "" && func() bool {
//line /usr/local/go/src/encoding/json/encode.go:1270
					_go_fuzz_dep_.CoverTab[27917]++
//line /usr/local/go/src/encoding/json/encode.go:1270
					return ft.Kind() == reflect.Pointer
//line /usr/local/go/src/encoding/json/encode.go:1270
					// _ = "end of CoverTab[27917]"
//line /usr/local/go/src/encoding/json/encode.go:1270
				}() {
//line /usr/local/go/src/encoding/json/encode.go:1270
					_go_fuzz_dep_.CoverTab[27918]++

											ft = ft.Elem()
//line /usr/local/go/src/encoding/json/encode.go:1272
					// _ = "end of CoverTab[27918]"
				} else {
//line /usr/local/go/src/encoding/json/encode.go:1273
					_go_fuzz_dep_.CoverTab[27919]++
//line /usr/local/go/src/encoding/json/encode.go:1273
					// _ = "end of CoverTab[27919]"
//line /usr/local/go/src/encoding/json/encode.go:1273
				}
//line /usr/local/go/src/encoding/json/encode.go:1273
				// _ = "end of CoverTab[27899]"
//line /usr/local/go/src/encoding/json/encode.go:1273
				_go_fuzz_dep_.CoverTab[27900]++

//line /usr/local/go/src/encoding/json/encode.go:1276
				quoted := false
				if opts.Contains("string") {
//line /usr/local/go/src/encoding/json/encode.go:1277
					_go_fuzz_dep_.CoverTab[27920]++
											switch ft.Kind() {
					case reflect.Bool,
						reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
						reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
						reflect.Float32, reflect.Float64,
						reflect.String:
//line /usr/local/go/src/encoding/json/encode.go:1283
						_go_fuzz_dep_.CoverTab[27921]++
												quoted = true
//line /usr/local/go/src/encoding/json/encode.go:1284
						// _ = "end of CoverTab[27921]"
//line /usr/local/go/src/encoding/json/encode.go:1284
					default:
//line /usr/local/go/src/encoding/json/encode.go:1284
						_go_fuzz_dep_.CoverTab[27922]++
//line /usr/local/go/src/encoding/json/encode.go:1284
						// _ = "end of CoverTab[27922]"
					}
//line /usr/local/go/src/encoding/json/encode.go:1285
					// _ = "end of CoverTab[27920]"
				} else {
//line /usr/local/go/src/encoding/json/encode.go:1286
					_go_fuzz_dep_.CoverTab[27923]++
//line /usr/local/go/src/encoding/json/encode.go:1286
					// _ = "end of CoverTab[27923]"
//line /usr/local/go/src/encoding/json/encode.go:1286
				}
//line /usr/local/go/src/encoding/json/encode.go:1286
				// _ = "end of CoverTab[27900]"
//line /usr/local/go/src/encoding/json/encode.go:1286
				_go_fuzz_dep_.CoverTab[27901]++

//line /usr/local/go/src/encoding/json/encode.go:1289
				if name != "" || func() bool {
//line /usr/local/go/src/encoding/json/encode.go:1289
					_go_fuzz_dep_.CoverTab[27924]++
//line /usr/local/go/src/encoding/json/encode.go:1289
					return !sf.Anonymous
//line /usr/local/go/src/encoding/json/encode.go:1289
					// _ = "end of CoverTab[27924]"
//line /usr/local/go/src/encoding/json/encode.go:1289
				}() || func() bool {
//line /usr/local/go/src/encoding/json/encode.go:1289
					_go_fuzz_dep_.CoverTab[27925]++
//line /usr/local/go/src/encoding/json/encode.go:1289
					return ft.Kind() != reflect.Struct
//line /usr/local/go/src/encoding/json/encode.go:1289
					// _ = "end of CoverTab[27925]"
//line /usr/local/go/src/encoding/json/encode.go:1289
				}() {
//line /usr/local/go/src/encoding/json/encode.go:1289
					_go_fuzz_dep_.CoverTab[27926]++
											tagged := name != ""
											if name == "" {
//line /usr/local/go/src/encoding/json/encode.go:1291
						_go_fuzz_dep_.CoverTab[27929]++
												name = sf.Name
//line /usr/local/go/src/encoding/json/encode.go:1292
						// _ = "end of CoverTab[27929]"
					} else {
//line /usr/local/go/src/encoding/json/encode.go:1293
						_go_fuzz_dep_.CoverTab[27930]++
//line /usr/local/go/src/encoding/json/encode.go:1293
						// _ = "end of CoverTab[27930]"
//line /usr/local/go/src/encoding/json/encode.go:1293
					}
//line /usr/local/go/src/encoding/json/encode.go:1293
					// _ = "end of CoverTab[27926]"
//line /usr/local/go/src/encoding/json/encode.go:1293
					_go_fuzz_dep_.CoverTab[27927]++
											field := field{
						name:		name,
						tag:		tagged,
						index:		index,
						typ:		ft,
						omitEmpty:	opts.Contains("omitempty"),
						quoted:		quoted,
					}
											field.nameBytes = []byte(field.name)
											field.equalFold = foldFunc(field.nameBytes)

//line /usr/local/go/src/encoding/json/encode.go:1306
					nameEscBuf.Reset()
					nameEscBuf.WriteString(`"`)
					HTMLEscape(&nameEscBuf, field.nameBytes)
					nameEscBuf.WriteString(`":`)
					field.nameEscHTML = nameEscBuf.String()
					field.nameNonEsc = `"` + field.name + `":`

					fields = append(fields, field)
					if count[f.typ] > 1 {
//line /usr/local/go/src/encoding/json/encode.go:1314
						_go_fuzz_dep_.CoverTab[27931]++

//line /usr/local/go/src/encoding/json/encode.go:1319
						fields = append(fields, fields[len(fields)-1])
//line /usr/local/go/src/encoding/json/encode.go:1319
						// _ = "end of CoverTab[27931]"
					} else {
//line /usr/local/go/src/encoding/json/encode.go:1320
						_go_fuzz_dep_.CoverTab[27932]++
//line /usr/local/go/src/encoding/json/encode.go:1320
						// _ = "end of CoverTab[27932]"
//line /usr/local/go/src/encoding/json/encode.go:1320
					}
//line /usr/local/go/src/encoding/json/encode.go:1320
					// _ = "end of CoverTab[27927]"
//line /usr/local/go/src/encoding/json/encode.go:1320
					_go_fuzz_dep_.CoverTab[27928]++
											continue
//line /usr/local/go/src/encoding/json/encode.go:1321
					// _ = "end of CoverTab[27928]"
				} else {
//line /usr/local/go/src/encoding/json/encode.go:1322
					_go_fuzz_dep_.CoverTab[27933]++
//line /usr/local/go/src/encoding/json/encode.go:1322
					// _ = "end of CoverTab[27933]"
//line /usr/local/go/src/encoding/json/encode.go:1322
				}
//line /usr/local/go/src/encoding/json/encode.go:1322
				// _ = "end of CoverTab[27901]"
//line /usr/local/go/src/encoding/json/encode.go:1322
				_go_fuzz_dep_.CoverTab[27902]++

//line /usr/local/go/src/encoding/json/encode.go:1325
				nextCount[ft]++
				if nextCount[ft] == 1 {
//line /usr/local/go/src/encoding/json/encode.go:1326
					_go_fuzz_dep_.CoverTab[27934]++
											next = append(next, field{name: ft.Name(), index: index, typ: ft})
//line /usr/local/go/src/encoding/json/encode.go:1327
					// _ = "end of CoverTab[27934]"
				} else {
//line /usr/local/go/src/encoding/json/encode.go:1328
					_go_fuzz_dep_.CoverTab[27935]++
//line /usr/local/go/src/encoding/json/encode.go:1328
					// _ = "end of CoverTab[27935]"
//line /usr/local/go/src/encoding/json/encode.go:1328
				}
//line /usr/local/go/src/encoding/json/encode.go:1328
				// _ = "end of CoverTab[27902]"
			}
//line /usr/local/go/src/encoding/json/encode.go:1329
			// _ = "end of CoverTab[27893]"
		}
//line /usr/local/go/src/encoding/json/encode.go:1330
		// _ = "end of CoverTab[27891]"
	}
//line /usr/local/go/src/encoding/json/encode.go:1331
	// _ = "end of CoverTab[27885]"
//line /usr/local/go/src/encoding/json/encode.go:1331
	_go_fuzz_dep_.CoverTab[27886]++

							sort.Slice(fields, func(i, j int) bool {
//line /usr/local/go/src/encoding/json/encode.go:1333
		_go_fuzz_dep_.CoverTab[27936]++
								x := fields

//line /usr/local/go/src/encoding/json/encode.go:1338
		if x[i].name != x[j].name {
//line /usr/local/go/src/encoding/json/encode.go:1338
			_go_fuzz_dep_.CoverTab[27940]++
									return x[i].name < x[j].name
//line /usr/local/go/src/encoding/json/encode.go:1339
			// _ = "end of CoverTab[27940]"
		} else {
//line /usr/local/go/src/encoding/json/encode.go:1340
			_go_fuzz_dep_.CoverTab[27941]++
//line /usr/local/go/src/encoding/json/encode.go:1340
			// _ = "end of CoverTab[27941]"
//line /usr/local/go/src/encoding/json/encode.go:1340
		}
//line /usr/local/go/src/encoding/json/encode.go:1340
		// _ = "end of CoverTab[27936]"
//line /usr/local/go/src/encoding/json/encode.go:1340
		_go_fuzz_dep_.CoverTab[27937]++
								if len(x[i].index) != len(x[j].index) {
//line /usr/local/go/src/encoding/json/encode.go:1341
			_go_fuzz_dep_.CoverTab[27942]++
									return len(x[i].index) < len(x[j].index)
//line /usr/local/go/src/encoding/json/encode.go:1342
			// _ = "end of CoverTab[27942]"
		} else {
//line /usr/local/go/src/encoding/json/encode.go:1343
			_go_fuzz_dep_.CoverTab[27943]++
//line /usr/local/go/src/encoding/json/encode.go:1343
			// _ = "end of CoverTab[27943]"
//line /usr/local/go/src/encoding/json/encode.go:1343
		}
//line /usr/local/go/src/encoding/json/encode.go:1343
		// _ = "end of CoverTab[27937]"
//line /usr/local/go/src/encoding/json/encode.go:1343
		_go_fuzz_dep_.CoverTab[27938]++
								if x[i].tag != x[j].tag {
//line /usr/local/go/src/encoding/json/encode.go:1344
			_go_fuzz_dep_.CoverTab[27944]++
									return x[i].tag
//line /usr/local/go/src/encoding/json/encode.go:1345
			// _ = "end of CoverTab[27944]"
		} else {
//line /usr/local/go/src/encoding/json/encode.go:1346
			_go_fuzz_dep_.CoverTab[27945]++
//line /usr/local/go/src/encoding/json/encode.go:1346
			// _ = "end of CoverTab[27945]"
//line /usr/local/go/src/encoding/json/encode.go:1346
		}
//line /usr/local/go/src/encoding/json/encode.go:1346
		// _ = "end of CoverTab[27938]"
//line /usr/local/go/src/encoding/json/encode.go:1346
		_go_fuzz_dep_.CoverTab[27939]++
								return byIndex(x).Less(i, j)
//line /usr/local/go/src/encoding/json/encode.go:1347
		// _ = "end of CoverTab[27939]"
	})
//line /usr/local/go/src/encoding/json/encode.go:1348
	// _ = "end of CoverTab[27886]"
//line /usr/local/go/src/encoding/json/encode.go:1348
	_go_fuzz_dep_.CoverTab[27887]++

//line /usr/local/go/src/encoding/json/encode.go:1356
	out := fields[:0]
	for advance, i := 0, 0; i < len(fields); i += advance {
//line /usr/local/go/src/encoding/json/encode.go:1357
		_go_fuzz_dep_.CoverTab[27946]++

//line /usr/local/go/src/encoding/json/encode.go:1360
		fi := fields[i]
		name := fi.name
		for advance = 1; i+advance < len(fields); advance++ {
//line /usr/local/go/src/encoding/json/encode.go:1362
			_go_fuzz_dep_.CoverTab[27949]++
									fj := fields[i+advance]
									if fj.name != name {
//line /usr/local/go/src/encoding/json/encode.go:1364
				_go_fuzz_dep_.CoverTab[27950]++
										break
//line /usr/local/go/src/encoding/json/encode.go:1365
				// _ = "end of CoverTab[27950]"
			} else {
//line /usr/local/go/src/encoding/json/encode.go:1366
				_go_fuzz_dep_.CoverTab[27951]++
//line /usr/local/go/src/encoding/json/encode.go:1366
				// _ = "end of CoverTab[27951]"
//line /usr/local/go/src/encoding/json/encode.go:1366
			}
//line /usr/local/go/src/encoding/json/encode.go:1366
			// _ = "end of CoverTab[27949]"
		}
//line /usr/local/go/src/encoding/json/encode.go:1367
		// _ = "end of CoverTab[27946]"
//line /usr/local/go/src/encoding/json/encode.go:1367
		_go_fuzz_dep_.CoverTab[27947]++
								if advance == 1 {
//line /usr/local/go/src/encoding/json/encode.go:1368
			_go_fuzz_dep_.CoverTab[27952]++
									out = append(out, fi)
									continue
//line /usr/local/go/src/encoding/json/encode.go:1370
			// _ = "end of CoverTab[27952]"
		} else {
//line /usr/local/go/src/encoding/json/encode.go:1371
			_go_fuzz_dep_.CoverTab[27953]++
//line /usr/local/go/src/encoding/json/encode.go:1371
			// _ = "end of CoverTab[27953]"
//line /usr/local/go/src/encoding/json/encode.go:1371
		}
//line /usr/local/go/src/encoding/json/encode.go:1371
		// _ = "end of CoverTab[27947]"
//line /usr/local/go/src/encoding/json/encode.go:1371
		_go_fuzz_dep_.CoverTab[27948]++
								dominant, ok := dominantField(fields[i : i+advance])
								if ok {
//line /usr/local/go/src/encoding/json/encode.go:1373
			_go_fuzz_dep_.CoverTab[27954]++
									out = append(out, dominant)
//line /usr/local/go/src/encoding/json/encode.go:1374
			// _ = "end of CoverTab[27954]"
		} else {
//line /usr/local/go/src/encoding/json/encode.go:1375
			_go_fuzz_dep_.CoverTab[27955]++
//line /usr/local/go/src/encoding/json/encode.go:1375
			// _ = "end of CoverTab[27955]"
//line /usr/local/go/src/encoding/json/encode.go:1375
		}
//line /usr/local/go/src/encoding/json/encode.go:1375
		// _ = "end of CoverTab[27948]"
	}
//line /usr/local/go/src/encoding/json/encode.go:1376
	// _ = "end of CoverTab[27887]"
//line /usr/local/go/src/encoding/json/encode.go:1376
	_go_fuzz_dep_.CoverTab[27888]++

							fields = out
							sort.Sort(byIndex(fields))

							for i := range fields {
//line /usr/local/go/src/encoding/json/encode.go:1381
		_go_fuzz_dep_.CoverTab[27956]++
								f := &fields[i]
								f.encoder = typeEncoder(typeByIndex(t, f.index))
//line /usr/local/go/src/encoding/json/encode.go:1383
		// _ = "end of CoverTab[27956]"
	}
//line /usr/local/go/src/encoding/json/encode.go:1384
	// _ = "end of CoverTab[27888]"
//line /usr/local/go/src/encoding/json/encode.go:1384
	_go_fuzz_dep_.CoverTab[27889]++
							nameIndex := make(map[string]int, len(fields))
							for i, field := range fields {
//line /usr/local/go/src/encoding/json/encode.go:1386
		_go_fuzz_dep_.CoverTab[27957]++
								nameIndex[field.name] = i
//line /usr/local/go/src/encoding/json/encode.go:1387
		// _ = "end of CoverTab[27957]"
	}
//line /usr/local/go/src/encoding/json/encode.go:1388
	// _ = "end of CoverTab[27889]"
//line /usr/local/go/src/encoding/json/encode.go:1388
	_go_fuzz_dep_.CoverTab[27890]++
							return structFields{fields, nameIndex}
//line /usr/local/go/src/encoding/json/encode.go:1389
	// _ = "end of CoverTab[27890]"
}

// dominantField looks through the fields, all of which are known to
//line /usr/local/go/src/encoding/json/encode.go:1392
// have the same name, to find the single field that dominates the
//line /usr/local/go/src/encoding/json/encode.go:1392
// others using Go's embedding rules, modified by the presence of
//line /usr/local/go/src/encoding/json/encode.go:1392
// JSON tags. If there are multiple top-level fields, the boolean
//line /usr/local/go/src/encoding/json/encode.go:1392
// will be false: This condition is an error in Go and we skip all
//line /usr/local/go/src/encoding/json/encode.go:1392
// the fields.
//line /usr/local/go/src/encoding/json/encode.go:1398
func dominantField(fields []field) (field, bool) {
//line /usr/local/go/src/encoding/json/encode.go:1398
	_go_fuzz_dep_.CoverTab[27958]++

//line /usr/local/go/src/encoding/json/encode.go:1402
	if len(fields) > 1 && func() bool {
//line /usr/local/go/src/encoding/json/encode.go:1402
		_go_fuzz_dep_.CoverTab[27960]++
//line /usr/local/go/src/encoding/json/encode.go:1402
		return len(fields[0].index) == len(fields[1].index)
//line /usr/local/go/src/encoding/json/encode.go:1402
		// _ = "end of CoverTab[27960]"
//line /usr/local/go/src/encoding/json/encode.go:1402
	}() && func() bool {
//line /usr/local/go/src/encoding/json/encode.go:1402
		_go_fuzz_dep_.CoverTab[27961]++
//line /usr/local/go/src/encoding/json/encode.go:1402
		return fields[0].tag == fields[1].tag
//line /usr/local/go/src/encoding/json/encode.go:1402
		// _ = "end of CoverTab[27961]"
//line /usr/local/go/src/encoding/json/encode.go:1402
	}() {
//line /usr/local/go/src/encoding/json/encode.go:1402
		_go_fuzz_dep_.CoverTab[27962]++
								return field{}, false
//line /usr/local/go/src/encoding/json/encode.go:1403
		// _ = "end of CoverTab[27962]"
	} else {
//line /usr/local/go/src/encoding/json/encode.go:1404
		_go_fuzz_dep_.CoverTab[27963]++
//line /usr/local/go/src/encoding/json/encode.go:1404
		// _ = "end of CoverTab[27963]"
//line /usr/local/go/src/encoding/json/encode.go:1404
	}
//line /usr/local/go/src/encoding/json/encode.go:1404
	// _ = "end of CoverTab[27958]"
//line /usr/local/go/src/encoding/json/encode.go:1404
	_go_fuzz_dep_.CoverTab[27959]++
							return fields[0], true
//line /usr/local/go/src/encoding/json/encode.go:1405
	// _ = "end of CoverTab[27959]"
}

var fieldCache sync.Map	// map[reflect.Type]structFields

// cachedTypeFields is like typeFields but uses a cache to avoid repeated work.
func cachedTypeFields(t reflect.Type) structFields {
//line /usr/local/go/src/encoding/json/encode.go:1411
	_go_fuzz_dep_.CoverTab[27964]++
							if f, ok := fieldCache.Load(t); ok {
//line /usr/local/go/src/encoding/json/encode.go:1412
		_go_fuzz_dep_.CoverTab[27966]++
								return f.(structFields)
//line /usr/local/go/src/encoding/json/encode.go:1413
		// _ = "end of CoverTab[27966]"
	} else {
//line /usr/local/go/src/encoding/json/encode.go:1414
		_go_fuzz_dep_.CoverTab[27967]++
//line /usr/local/go/src/encoding/json/encode.go:1414
		// _ = "end of CoverTab[27967]"
//line /usr/local/go/src/encoding/json/encode.go:1414
	}
//line /usr/local/go/src/encoding/json/encode.go:1414
	// _ = "end of CoverTab[27964]"
//line /usr/local/go/src/encoding/json/encode.go:1414
	_go_fuzz_dep_.CoverTab[27965]++
							f, _ := fieldCache.LoadOrStore(t, typeFields(t))
							return f.(structFields)
//line /usr/local/go/src/encoding/json/encode.go:1416
	// _ = "end of CoverTab[27965]"
}

//line /usr/local/go/src/encoding/json/encode.go:1417
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/encoding/json/encode.go:1417
var _ = _go_fuzz_dep_.CoverTab
