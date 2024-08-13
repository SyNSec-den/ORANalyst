// Copyright 2010 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:5
// Package json implements encoding and decoding of JSON objects as defined in
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:5
// RFC 4627. The mapping between JSON objects and Go values is described
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:5
// in the documentation for the Marshal and Unmarshal functions.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:5
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:5
// See "JSON and Go" for an introduction to this package:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:5
// https://golang.org/doc/articles/json_and_go.html
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:11
package json

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:11
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:11
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:11
)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:11
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:11
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:11
)

import (
	"bytes"
	"encoding"
	"encoding/base64"
	"fmt"
	"math"
	"reflect"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
	"unicode"
	"unicode/utf8"
)

// Marshal returns the JSON encoding of v.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
// Marshal traverses the value v recursively.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
// If an encountered value implements the Marshaler interface
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
// and is not a nil pointer, Marshal calls its MarshalJSON method
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
// to produce JSON. If no MarshalJSON method is present but the
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
// value implements encoding.TextMarshaler instead, Marshal calls
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
// its MarshalText method.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
// The nil pointer exception is not strictly necessary
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
// but mimics a similar, necessary exception in the behavior of
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
// UnmarshalJSON.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
// Otherwise, Marshal uses the following type-dependent default encodings:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
// Boolean values encode as JSON booleans.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
// Floating point, integer, and Number values encode as JSON numbers.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
// String values encode as JSON strings coerced to valid UTF-8,
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
// replacing invalid bytes with the Unicode replacement rune.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
// The angle brackets "<" and ">" are escaped to "\u003c" and "\u003e"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
// to keep some browsers from misinterpreting JSON output as HTML.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
// Ampersand "&" is also escaped to "\u0026" for the same reason.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
// Array and slice values encode as JSON arrays, except that
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
// []byte encodes as a base64-encoded string, and a nil slice
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
// encodes as the null JSON object.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
// Struct values encode as JSON objects. Each exported struct field
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
// becomes a member of the object unless
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
//   - the field's tag is "-", or
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
//   - the field is empty and its tag specifies the "omitempty" option.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
// The empty values are false, 0, any
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
// nil pointer or interface value, and any array, slice, map, or string of
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
// length zero. The object's default key string is the struct field name
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
// but can be specified in the struct field's tag value. The "json" key in
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
// the struct field's tag value is the key name, followed by an optional comma
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
// and options. Examples:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
//	// Field is ignored by this package.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
//	Field int `json:"-"`
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
//	// Field appears in JSON as key "myName".
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
//	Field int `json:"myName"`
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
//	// Field appears in JSON as key "myName" and
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
//	// the field is omitted from the object if its value is empty,
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
//	// as defined above.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
//	Field int `json:"myName,omitempty"`
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
//	// Field appears in JSON as key "Field" (the default), but
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
//	// the field is skipped if empty.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
//	// Note the leading comma.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
//	Field int `json:",omitempty"`
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
// The "string" option signals that a field is stored as JSON inside a
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
// JSON-encoded string. It applies only to fields of string, floating point,
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
// integer, or boolean types. This extra level of encoding is sometimes used
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
// when communicating with JavaScript programs:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
//	Int64String int64 `json:",string"`
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
// The key name will be used if it's a non-empty string consisting of
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
// only Unicode letters, digits, dollar signs, percent signs, hyphens,
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
// underscores and slashes.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
// Anonymous struct fields are usually marshaled as if their inner exported fields
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
// were fields in the outer struct, subject to the usual Go visibility rules amended
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
// as described in the next paragraph.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
// An anonymous struct field with a name given in its JSON tag is treated as
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
// having that name, rather than being anonymous.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
// An anonymous struct field of interface type is treated the same as having
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
// that type as its name, rather than being anonymous.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
// The Go visibility rules for struct fields are amended for JSON when
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
// deciding which field to marshal or unmarshal. If there are
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
// multiple fields at the same level, and that level is the least
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
// nested (and would therefore be the nesting level selected by the
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
// usual Go rules), the following extra rules apply:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
// 1) Of those fields, if any are JSON-tagged, only tagged fields are considered,
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
// even if there are multiple untagged fields that would otherwise conflict.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
// 2) If there is exactly one field (tagged or not according to the first rule), that is selected.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
// 3) Otherwise there are multiple fields, and all are ignored; no error occurs.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
// Handling of anonymous struct fields is new in Go 1.1.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
// Prior to Go 1.1, anonymous struct fields were ignored. To force ignoring of
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
// an anonymous struct field in both current and earlier versions, give the field
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
// a JSON tag of "-".
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
// Map values encode as JSON objects.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
// The map's key type must be string; the map keys are used as JSON object
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
// keys, subject to the UTF-8 coercion described for string values above.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
// Pointer values encode as the value pointed to.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
// A nil pointer encodes as the null JSON object.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
// Interface values encode as the value contained in the interface.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
// A nil interface value encodes as the null JSON object.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
// Channel, complex, and function values cannot be encoded in JSON.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
// Attempting to encode such a value causes Marshal to return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
// an UnsupportedTypeError.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
// JSON cannot represent cyclic data structures and Marshal does not
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
// handle them.  Passing cyclic structures to Marshal will result in
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:29
// an infinite recursion.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:137
func Marshal(v interface{}) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:137
	_go_fuzz_dep_.CoverTab[184873]++
											e := &encodeState{}
											err := e.marshal(v)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:140
		_go_fuzz_dep_.CoverTab[184875]++
												return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:141
		// _ = "end of CoverTab[184875]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:142
		_go_fuzz_dep_.CoverTab[184876]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:142
		// _ = "end of CoverTab[184876]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:142
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:142
	// _ = "end of CoverTab[184873]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:142
	_go_fuzz_dep_.CoverTab[184874]++
											return e.Bytes(), nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:143
	// _ = "end of CoverTab[184874]"
}

// MarshalIndent is like Marshal but applies Indent to format the output.
func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:147
	_go_fuzz_dep_.CoverTab[184877]++
											b, err := Marshal(v)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:149
		_go_fuzz_dep_.CoverTab[184880]++
												return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:150
		// _ = "end of CoverTab[184880]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:151
		_go_fuzz_dep_.CoverTab[184881]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:151
		// _ = "end of CoverTab[184881]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:151
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:151
	// _ = "end of CoverTab[184877]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:151
	_go_fuzz_dep_.CoverTab[184878]++
											var buf bytes.Buffer
											err = Indent(&buf, b, prefix, indent)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:154
		_go_fuzz_dep_.CoverTab[184882]++
												return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:155
		// _ = "end of CoverTab[184882]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:156
		_go_fuzz_dep_.CoverTab[184883]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:156
		// _ = "end of CoverTab[184883]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:156
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:156
	// _ = "end of CoverTab[184878]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:156
	_go_fuzz_dep_.CoverTab[184879]++
											return buf.Bytes(), nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:157
	// _ = "end of CoverTab[184879]"
}

// HTMLEscape appends to dst the JSON-encoded src with <, >, &, U+2028 and U+2029
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:160
// characters inside string literals changed to \u003c, \u003e, \u0026, \u2028, \u2029
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:160
// so that the JSON will be safe to embed inside HTML <script> tags.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:160
// For historical reasons, web browsers don't honor standard HTML
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:160
// escaping within <script> tags, so an alternative JSON encoding must
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:160
// be used.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:166
func HTMLEscape(dst *bytes.Buffer, src []byte) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:166
	_go_fuzz_dep_.CoverTab[184884]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:169
	start := 0
	for i, c := range src {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:170
		_go_fuzz_dep_.CoverTab[184886]++
												if c == '<' || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:171
			_go_fuzz_dep_.CoverTab[184888]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:171
			return c == '>'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:171
			// _ = "end of CoverTab[184888]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:171
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:171
			_go_fuzz_dep_.CoverTab[184889]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:171
			return c == '&'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:171
			// _ = "end of CoverTab[184889]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:171
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:171
			_go_fuzz_dep_.CoverTab[184890]++
													if start < i {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:172
				_go_fuzz_dep_.CoverTab[184892]++
														dst.Write(src[start:i])
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:173
				// _ = "end of CoverTab[184892]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:174
				_go_fuzz_dep_.CoverTab[184893]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:174
				// _ = "end of CoverTab[184893]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:174
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:174
			// _ = "end of CoverTab[184890]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:174
			_go_fuzz_dep_.CoverTab[184891]++
													dst.WriteString(`\u00`)
													dst.WriteByte(hex[c>>4])
													dst.WriteByte(hex[c&0xF])
													start = i + 1
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:178
			// _ = "end of CoverTab[184891]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:179
			_go_fuzz_dep_.CoverTab[184894]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:179
			// _ = "end of CoverTab[184894]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:179
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:179
		// _ = "end of CoverTab[184886]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:179
		_go_fuzz_dep_.CoverTab[184887]++

												if c == 0xE2 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:181
			_go_fuzz_dep_.CoverTab[184895]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:181
			return i+2 < len(src)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:181
			// _ = "end of CoverTab[184895]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:181
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:181
			_go_fuzz_dep_.CoverTab[184896]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:181
			return src[i+1] == 0x80
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:181
			// _ = "end of CoverTab[184896]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:181
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:181
			_go_fuzz_dep_.CoverTab[184897]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:181
			return src[i+2]&^1 == 0xA8
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:181
			// _ = "end of CoverTab[184897]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:181
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:181
			_go_fuzz_dep_.CoverTab[184898]++
													if start < i {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:182
				_go_fuzz_dep_.CoverTab[184900]++
														dst.Write(src[start:i])
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:183
				// _ = "end of CoverTab[184900]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:184
				_go_fuzz_dep_.CoverTab[184901]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:184
				// _ = "end of CoverTab[184901]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:184
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:184
			// _ = "end of CoverTab[184898]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:184
			_go_fuzz_dep_.CoverTab[184899]++
													dst.WriteString(`\u202`)
													dst.WriteByte(hex[src[i+2]&0xF])
													start = i + 3
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:187
			// _ = "end of CoverTab[184899]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:188
			_go_fuzz_dep_.CoverTab[184902]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:188
			// _ = "end of CoverTab[184902]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:188
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:188
		// _ = "end of CoverTab[184887]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:189
	// _ = "end of CoverTab[184884]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:189
	_go_fuzz_dep_.CoverTab[184885]++
											if start < len(src) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:190
		_go_fuzz_dep_.CoverTab[184903]++
												dst.Write(src[start:])
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:191
		// _ = "end of CoverTab[184903]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:192
		_go_fuzz_dep_.CoverTab[184904]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:192
		// _ = "end of CoverTab[184904]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:192
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:192
	// _ = "end of CoverTab[184885]"
}

// Marshaler is the interface implemented by objects that
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:195
// can marshal themselves into valid JSON.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:197
type Marshaler interface {
	MarshalJSON() ([]byte, error)
}

// An UnsupportedTypeError is returned by Marshal when attempting
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:201
// to encode an unsupported value type.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:203
type UnsupportedTypeError struct {
	Type reflect.Type
}

func (e *UnsupportedTypeError) Error() string {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:207
	_go_fuzz_dep_.CoverTab[184905]++
											return "json: unsupported type: " + e.Type.String()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:208
	// _ = "end of CoverTab[184905]"
}

type UnsupportedValueError struct {
	Value	reflect.Value
	Str	string
}

func (e *UnsupportedValueError) Error() string {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:216
	_go_fuzz_dep_.CoverTab[184906]++
											return "json: unsupported value: " + e.Str
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:217
	// _ = "end of CoverTab[184906]"
}

// Before Go 1.2, an InvalidUTF8Error was returned by Marshal when
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:220
// attempting to encode a string value with invalid UTF-8 sequences.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:220
// As of Go 1.2, Marshal instead coerces the string to valid UTF-8 by
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:220
// replacing invalid bytes with the Unicode replacement rune U+FFFD.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:220
// This error is no longer generated but is kept for backwards compatibility
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:220
// with programs that might mention it.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:226
type InvalidUTF8Error struct {
	S string	// the whole string value that caused the error
}

func (e *InvalidUTF8Error) Error() string {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:230
	_go_fuzz_dep_.CoverTab[184907]++
											return "json: invalid UTF-8 in string: " + strconv.Quote(e.S)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:231
	// _ = "end of CoverTab[184907]"
}

type MarshalerError struct {
	Type	reflect.Type
	Err	error
}

func (e *MarshalerError) Error() string {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:239
	_go_fuzz_dep_.CoverTab[184908]++
											return "json: error calling MarshalJSON for type " + e.Type.String() + ": " + e.Err.Error()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:240
	// _ = "end of CoverTab[184908]"
}

var hex = "0123456789abcdef"

// An encodeState encodes JSON into a bytes.Buffer.
type encodeState struct {
	bytes.Buffer	// accumulated output
	scratch		[64]byte
}

var encodeStatePool sync.Pool

func newEncodeState() *encodeState {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:253
	_go_fuzz_dep_.CoverTab[184909]++
											if v := encodeStatePool.Get(); v != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:254
		_go_fuzz_dep_.CoverTab[184911]++
												e := v.(*encodeState)
												e.Reset()
												return e
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:257
		// _ = "end of CoverTab[184911]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:258
		_go_fuzz_dep_.CoverTab[184912]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:258
		// _ = "end of CoverTab[184912]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:258
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:258
	// _ = "end of CoverTab[184909]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:258
	_go_fuzz_dep_.CoverTab[184910]++
											return new(encodeState)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:259
	// _ = "end of CoverTab[184910]"
}

func (e *encodeState) marshal(v interface{}) (err error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:262
	_go_fuzz_dep_.CoverTab[184913]++
											defer func() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:263
		_go_fuzz_dep_.CoverTab[184915]++
												if r := recover(); r != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:264
			_go_fuzz_dep_.CoverTab[184916]++
													if _, ok := r.(runtime.Error); ok {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:265
				_go_fuzz_dep_.CoverTab[184919]++
														panic(r)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:266
				// _ = "end of CoverTab[184919]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:267
				_go_fuzz_dep_.CoverTab[184920]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:267
				// _ = "end of CoverTab[184920]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:267
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:267
			// _ = "end of CoverTab[184916]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:267
			_go_fuzz_dep_.CoverTab[184917]++
													if s, ok := r.(string); ok {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:268
				_go_fuzz_dep_.CoverTab[184921]++
														panic(s)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:269
				// _ = "end of CoverTab[184921]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:270
				_go_fuzz_dep_.CoverTab[184922]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:270
				// _ = "end of CoverTab[184922]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:270
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:270
			// _ = "end of CoverTab[184917]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:270
			_go_fuzz_dep_.CoverTab[184918]++
													err = r.(error)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:271
			// _ = "end of CoverTab[184918]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:272
			_go_fuzz_dep_.CoverTab[184923]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:272
			// _ = "end of CoverTab[184923]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:272
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:272
		// _ = "end of CoverTab[184915]"
	}()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:273
	// _ = "end of CoverTab[184913]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:273
	_go_fuzz_dep_.CoverTab[184914]++
											e.reflectValue(reflect.ValueOf(v))
											return nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:275
	// _ = "end of CoverTab[184914]"
}

func (e *encodeState) error(err error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:278
	_go_fuzz_dep_.CoverTab[184924]++
											panic(err)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:279
	// _ = "end of CoverTab[184924]"
}

func isEmptyValue(v reflect.Value) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:282
	_go_fuzz_dep_.CoverTab[184925]++
											switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:284
		_go_fuzz_dep_.CoverTab[184927]++
												return v.Len() == 0
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:285
		// _ = "end of CoverTab[184927]"
	case reflect.Bool:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:286
		_go_fuzz_dep_.CoverTab[184928]++
												return !v.Bool()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:287
		// _ = "end of CoverTab[184928]"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:288
		_go_fuzz_dep_.CoverTab[184929]++
												return v.Int() == 0
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:289
		// _ = "end of CoverTab[184929]"
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:290
		_go_fuzz_dep_.CoverTab[184930]++
												return v.Uint() == 0
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:291
		// _ = "end of CoverTab[184930]"
	case reflect.Float32, reflect.Float64:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:292
		_go_fuzz_dep_.CoverTab[184931]++
												return v.Float() == 0
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:293
		// _ = "end of CoverTab[184931]"
	case reflect.Interface, reflect.Ptr:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:294
		_go_fuzz_dep_.CoverTab[184932]++
												return v.IsNil()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:295
		// _ = "end of CoverTab[184932]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:295
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:295
		_go_fuzz_dep_.CoverTab[184933]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:295
		// _ = "end of CoverTab[184933]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:296
	// _ = "end of CoverTab[184925]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:296
	_go_fuzz_dep_.CoverTab[184926]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:297
	// _ = "end of CoverTab[184926]"
}

func (e *encodeState) reflectValue(v reflect.Value) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:300
	_go_fuzz_dep_.CoverTab[184934]++
											valueEncoder(v)(e, v, false)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:301
	// _ = "end of CoverTab[184934]"
}

type encoderFunc func(e *encodeState, v reflect.Value, quoted bool)

var encoderCache struct {
	sync.RWMutex
	m	map[reflect.Type]encoderFunc
}

func valueEncoder(v reflect.Value) encoderFunc {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:311
	_go_fuzz_dep_.CoverTab[184935]++
											if !v.IsValid() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:312
		_go_fuzz_dep_.CoverTab[184937]++
												return invalidValueEncoder
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:313
		// _ = "end of CoverTab[184937]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:314
		_go_fuzz_dep_.CoverTab[184938]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:314
		// _ = "end of CoverTab[184938]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:314
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:314
	// _ = "end of CoverTab[184935]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:314
	_go_fuzz_dep_.CoverTab[184936]++
											return typeEncoder(v.Type())
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:315
	// _ = "end of CoverTab[184936]"
}

func typeEncoder(t reflect.Type) encoderFunc {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:318
	_go_fuzz_dep_.CoverTab[184939]++
											encoderCache.RLock()
											f := encoderCache.m[t]
											encoderCache.RUnlock()
											if f != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:322
		_go_fuzz_dep_.CoverTab[184943]++
												return f
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:323
		// _ = "end of CoverTab[184943]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:324
		_go_fuzz_dep_.CoverTab[184944]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:324
		// _ = "end of CoverTab[184944]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:324
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:324
	// _ = "end of CoverTab[184939]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:324
	_go_fuzz_dep_.CoverTab[184940]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:330
	encoderCache.Lock()
	if encoderCache.m == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:331
		_go_fuzz_dep_.CoverTab[184945]++
												encoderCache.m = make(map[reflect.Type]encoderFunc)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:332
		// _ = "end of CoverTab[184945]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:333
		_go_fuzz_dep_.CoverTab[184946]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:333
		// _ = "end of CoverTab[184946]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:333
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:333
	// _ = "end of CoverTab[184940]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:333
	_go_fuzz_dep_.CoverTab[184941]++
											var wg sync.WaitGroup
											wg.Add(1)
											encoderCache.m[t] = func(e *encodeState, v reflect.Value, quoted bool) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:336
		_go_fuzz_dep_.CoverTab[184947]++
												wg.Wait()
												f(e, v, quoted)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:338
		// _ = "end of CoverTab[184947]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:339
	// _ = "end of CoverTab[184941]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:339
	_go_fuzz_dep_.CoverTab[184942]++
											encoderCache.Unlock()

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:344
	f = newTypeEncoder(t, true)
											wg.Done()
											encoderCache.Lock()
											encoderCache.m[t] = f
											encoderCache.Unlock()
											return f
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:349
	// _ = "end of CoverTab[184942]"
}

var (
	marshalerType		= reflect.TypeOf(new(Marshaler)).Elem()
	textMarshalerType	= reflect.TypeOf(new(encoding.TextMarshaler)).Elem()
)

// newTypeEncoder constructs an encoderFunc for a type.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:357
// The returned encoder only checks CanAddr when allowAddr is true.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:359
func newTypeEncoder(t reflect.Type, allowAddr bool) encoderFunc {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:359
	_go_fuzz_dep_.CoverTab[184948]++
											if t.Implements(marshalerType) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:360
		_go_fuzz_dep_.CoverTab[184953]++
												return marshalerEncoder
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:361
		// _ = "end of CoverTab[184953]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:362
		_go_fuzz_dep_.CoverTab[184954]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:362
		// _ = "end of CoverTab[184954]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:362
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:362
	// _ = "end of CoverTab[184948]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:362
	_go_fuzz_dep_.CoverTab[184949]++
											if t.Kind() != reflect.Ptr && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:363
		_go_fuzz_dep_.CoverTab[184955]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:363
		return allowAddr
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:363
		// _ = "end of CoverTab[184955]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:363
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:363
		_go_fuzz_dep_.CoverTab[184956]++
												if reflect.PtrTo(t).Implements(marshalerType) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:364
			_go_fuzz_dep_.CoverTab[184957]++
													return newCondAddrEncoder(addrMarshalerEncoder, newTypeEncoder(t, false))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:365
			// _ = "end of CoverTab[184957]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:366
			_go_fuzz_dep_.CoverTab[184958]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:366
			// _ = "end of CoverTab[184958]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:366
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:366
		// _ = "end of CoverTab[184956]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:367
		_go_fuzz_dep_.CoverTab[184959]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:367
		// _ = "end of CoverTab[184959]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:367
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:367
	// _ = "end of CoverTab[184949]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:367
	_go_fuzz_dep_.CoverTab[184950]++

											if t.Implements(textMarshalerType) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:369
		_go_fuzz_dep_.CoverTab[184960]++
												return textMarshalerEncoder
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:370
		// _ = "end of CoverTab[184960]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:371
		_go_fuzz_dep_.CoverTab[184961]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:371
		// _ = "end of CoverTab[184961]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:371
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:371
	// _ = "end of CoverTab[184950]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:371
	_go_fuzz_dep_.CoverTab[184951]++
											if t.Kind() != reflect.Ptr && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:372
		_go_fuzz_dep_.CoverTab[184962]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:372
		return allowAddr
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:372
		// _ = "end of CoverTab[184962]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:372
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:372
		_go_fuzz_dep_.CoverTab[184963]++
												if reflect.PtrTo(t).Implements(textMarshalerType) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:373
			_go_fuzz_dep_.CoverTab[184964]++
													return newCondAddrEncoder(addrTextMarshalerEncoder, newTypeEncoder(t, false))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:374
			// _ = "end of CoverTab[184964]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:375
			_go_fuzz_dep_.CoverTab[184965]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:375
			// _ = "end of CoverTab[184965]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:375
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:375
		// _ = "end of CoverTab[184963]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:376
		_go_fuzz_dep_.CoverTab[184966]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:376
		// _ = "end of CoverTab[184966]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:376
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:376
	// _ = "end of CoverTab[184951]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:376
	_go_fuzz_dep_.CoverTab[184952]++

											switch t.Kind() {
	case reflect.Bool:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:379
		_go_fuzz_dep_.CoverTab[184967]++
												return boolEncoder
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:380
		// _ = "end of CoverTab[184967]"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:381
		_go_fuzz_dep_.CoverTab[184968]++
												return intEncoder
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:382
		// _ = "end of CoverTab[184968]"
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:383
		_go_fuzz_dep_.CoverTab[184969]++
												return uintEncoder
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:384
		// _ = "end of CoverTab[184969]"
	case reflect.Float32:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:385
		_go_fuzz_dep_.CoverTab[184970]++
												return float32Encoder
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:386
		// _ = "end of CoverTab[184970]"
	case reflect.Float64:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:387
		_go_fuzz_dep_.CoverTab[184971]++
												return float64Encoder
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:388
		// _ = "end of CoverTab[184971]"
	case reflect.String:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:389
		_go_fuzz_dep_.CoverTab[184972]++
												return stringEncoder
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:390
		// _ = "end of CoverTab[184972]"
	case reflect.Interface:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:391
		_go_fuzz_dep_.CoverTab[184973]++
												return interfaceEncoder
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:392
		// _ = "end of CoverTab[184973]"
	case reflect.Struct:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:393
		_go_fuzz_dep_.CoverTab[184974]++
												return newStructEncoder(t)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:394
		// _ = "end of CoverTab[184974]"
	case reflect.Map:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:395
		_go_fuzz_dep_.CoverTab[184975]++
												return newMapEncoder(t)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:396
		// _ = "end of CoverTab[184975]"
	case reflect.Slice:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:397
		_go_fuzz_dep_.CoverTab[184976]++
												return newSliceEncoder(t)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:398
		// _ = "end of CoverTab[184976]"
	case reflect.Array:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:399
		_go_fuzz_dep_.CoverTab[184977]++
												return newArrayEncoder(t)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:400
		// _ = "end of CoverTab[184977]"
	case reflect.Ptr:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:401
		_go_fuzz_dep_.CoverTab[184978]++
												return newPtrEncoder(t)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:402
		// _ = "end of CoverTab[184978]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:403
		_go_fuzz_dep_.CoverTab[184979]++
												return unsupportedTypeEncoder
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:404
		// _ = "end of CoverTab[184979]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:405
	// _ = "end of CoverTab[184952]"
}

func invalidValueEncoder(e *encodeState, v reflect.Value, quoted bool) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:408
	_go_fuzz_dep_.CoverTab[184980]++
											e.WriteString("null")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:409
	// _ = "end of CoverTab[184980]"
}

func marshalerEncoder(e *encodeState, v reflect.Value, quoted bool) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:412
	_go_fuzz_dep_.CoverTab[184981]++
											if v.Kind() == reflect.Ptr && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:413
		_go_fuzz_dep_.CoverTab[184984]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:413
		return v.IsNil()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:413
		// _ = "end of CoverTab[184984]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:413
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:413
		_go_fuzz_dep_.CoverTab[184985]++
												e.WriteString("null")
												return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:415
		// _ = "end of CoverTab[184985]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:416
		_go_fuzz_dep_.CoverTab[184986]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:416
		// _ = "end of CoverTab[184986]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:416
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:416
	// _ = "end of CoverTab[184981]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:416
	_go_fuzz_dep_.CoverTab[184982]++
											m := v.Interface().(Marshaler)
											b, err := m.MarshalJSON()
											if err == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:419
		_go_fuzz_dep_.CoverTab[184987]++

												err = compact(&e.Buffer, b, true)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:421
		// _ = "end of CoverTab[184987]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:422
		_go_fuzz_dep_.CoverTab[184988]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:422
		// _ = "end of CoverTab[184988]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:422
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:422
	// _ = "end of CoverTab[184982]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:422
	_go_fuzz_dep_.CoverTab[184983]++
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:423
		_go_fuzz_dep_.CoverTab[184989]++
												e.error(&MarshalerError{v.Type(), err})
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:424
		// _ = "end of CoverTab[184989]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:425
		_go_fuzz_dep_.CoverTab[184990]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:425
		// _ = "end of CoverTab[184990]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:425
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:425
	// _ = "end of CoverTab[184983]"
}

func addrMarshalerEncoder(e *encodeState, v reflect.Value, quoted bool) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:428
	_go_fuzz_dep_.CoverTab[184991]++
											va := v.Addr()
											if va.IsNil() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:430
		_go_fuzz_dep_.CoverTab[184994]++
												e.WriteString("null")
												return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:432
		// _ = "end of CoverTab[184994]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:433
		_go_fuzz_dep_.CoverTab[184995]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:433
		// _ = "end of CoverTab[184995]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:433
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:433
	// _ = "end of CoverTab[184991]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:433
	_go_fuzz_dep_.CoverTab[184992]++
											m := va.Interface().(Marshaler)
											b, err := m.MarshalJSON()
											if err == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:436
		_go_fuzz_dep_.CoverTab[184996]++

												err = compact(&e.Buffer, b, true)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:438
		// _ = "end of CoverTab[184996]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:439
		_go_fuzz_dep_.CoverTab[184997]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:439
		// _ = "end of CoverTab[184997]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:439
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:439
	// _ = "end of CoverTab[184992]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:439
	_go_fuzz_dep_.CoverTab[184993]++
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:440
		_go_fuzz_dep_.CoverTab[184998]++
												e.error(&MarshalerError{v.Type(), err})
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:441
		// _ = "end of CoverTab[184998]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:442
		_go_fuzz_dep_.CoverTab[184999]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:442
		// _ = "end of CoverTab[184999]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:442
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:442
	// _ = "end of CoverTab[184993]"
}

func textMarshalerEncoder(e *encodeState, v reflect.Value, quoted bool) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:445
	_go_fuzz_dep_.CoverTab[185000]++
											if v.Kind() == reflect.Ptr && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:446
		_go_fuzz_dep_.CoverTab[185003]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:446
		return v.IsNil()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:446
		// _ = "end of CoverTab[185003]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:446
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:446
		_go_fuzz_dep_.CoverTab[185004]++
												e.WriteString("null")
												return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:448
		// _ = "end of CoverTab[185004]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:449
		_go_fuzz_dep_.CoverTab[185005]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:449
		// _ = "end of CoverTab[185005]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:449
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:449
	// _ = "end of CoverTab[185000]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:449
	_go_fuzz_dep_.CoverTab[185001]++
											m := v.Interface().(encoding.TextMarshaler)
											b, err := m.MarshalText()
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:452
		_go_fuzz_dep_.CoverTab[185006]++
												e.error(&MarshalerError{v.Type(), err})
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:453
		// _ = "end of CoverTab[185006]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:454
		_go_fuzz_dep_.CoverTab[185007]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:454
		// _ = "end of CoverTab[185007]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:454
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:454
	// _ = "end of CoverTab[185001]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:454
	_go_fuzz_dep_.CoverTab[185002]++
											e.stringBytes(b)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:455
	// _ = "end of CoverTab[185002]"
}

func addrTextMarshalerEncoder(e *encodeState, v reflect.Value, quoted bool) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:458
	_go_fuzz_dep_.CoverTab[185008]++
											va := v.Addr()
											if va.IsNil() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:460
		_go_fuzz_dep_.CoverTab[185011]++
												e.WriteString("null")
												return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:462
		// _ = "end of CoverTab[185011]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:463
		_go_fuzz_dep_.CoverTab[185012]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:463
		// _ = "end of CoverTab[185012]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:463
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:463
	// _ = "end of CoverTab[185008]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:463
	_go_fuzz_dep_.CoverTab[185009]++
											m := va.Interface().(encoding.TextMarshaler)
											b, err := m.MarshalText()
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:466
		_go_fuzz_dep_.CoverTab[185013]++
												e.error(&MarshalerError{v.Type(), err})
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:467
		// _ = "end of CoverTab[185013]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:468
		_go_fuzz_dep_.CoverTab[185014]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:468
		// _ = "end of CoverTab[185014]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:468
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:468
	// _ = "end of CoverTab[185009]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:468
	_go_fuzz_dep_.CoverTab[185010]++
											e.stringBytes(b)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:469
	// _ = "end of CoverTab[185010]"
}

func boolEncoder(e *encodeState, v reflect.Value, quoted bool) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:472
	_go_fuzz_dep_.CoverTab[185015]++
											if quoted {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:473
		_go_fuzz_dep_.CoverTab[185018]++
												e.WriteByte('"')
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:474
		// _ = "end of CoverTab[185018]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:475
		_go_fuzz_dep_.CoverTab[185019]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:475
		// _ = "end of CoverTab[185019]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:475
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:475
	// _ = "end of CoverTab[185015]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:475
	_go_fuzz_dep_.CoverTab[185016]++
											if v.Bool() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:476
		_go_fuzz_dep_.CoverTab[185020]++
												e.WriteString("true")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:477
		// _ = "end of CoverTab[185020]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:478
		_go_fuzz_dep_.CoverTab[185021]++
												e.WriteString("false")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:479
		// _ = "end of CoverTab[185021]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:480
	// _ = "end of CoverTab[185016]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:480
	_go_fuzz_dep_.CoverTab[185017]++
											if quoted {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:481
		_go_fuzz_dep_.CoverTab[185022]++
												e.WriteByte('"')
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:482
		// _ = "end of CoverTab[185022]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:483
		_go_fuzz_dep_.CoverTab[185023]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:483
		// _ = "end of CoverTab[185023]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:483
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:483
	// _ = "end of CoverTab[185017]"
}

func intEncoder(e *encodeState, v reflect.Value, quoted bool) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:486
	_go_fuzz_dep_.CoverTab[185024]++
											b := strconv.AppendInt(e.scratch[:0], v.Int(), 10)
											if quoted {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:488
		_go_fuzz_dep_.CoverTab[185026]++
												e.WriteByte('"')
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:489
		// _ = "end of CoverTab[185026]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:490
		_go_fuzz_dep_.CoverTab[185027]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:490
		// _ = "end of CoverTab[185027]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:490
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:490
	// _ = "end of CoverTab[185024]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:490
	_go_fuzz_dep_.CoverTab[185025]++
											e.Write(b)
											if quoted {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:492
		_go_fuzz_dep_.CoverTab[185028]++
												e.WriteByte('"')
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:493
		// _ = "end of CoverTab[185028]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:494
		_go_fuzz_dep_.CoverTab[185029]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:494
		// _ = "end of CoverTab[185029]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:494
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:494
	// _ = "end of CoverTab[185025]"
}

func uintEncoder(e *encodeState, v reflect.Value, quoted bool) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:497
	_go_fuzz_dep_.CoverTab[185030]++
											b := strconv.AppendUint(e.scratch[:0], v.Uint(), 10)
											if quoted {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:499
		_go_fuzz_dep_.CoverTab[185032]++
												e.WriteByte('"')
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:500
		// _ = "end of CoverTab[185032]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:501
		_go_fuzz_dep_.CoverTab[185033]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:501
		// _ = "end of CoverTab[185033]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:501
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:501
	// _ = "end of CoverTab[185030]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:501
	_go_fuzz_dep_.CoverTab[185031]++
											e.Write(b)
											if quoted {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:503
		_go_fuzz_dep_.CoverTab[185034]++
												e.WriteByte('"')
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:504
		// _ = "end of CoverTab[185034]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:505
		_go_fuzz_dep_.CoverTab[185035]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:505
		// _ = "end of CoverTab[185035]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:505
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:505
	// _ = "end of CoverTab[185031]"
}

type floatEncoder int	// number of bits

func (bits floatEncoder) encode(e *encodeState, v reflect.Value, quoted bool) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:510
	_go_fuzz_dep_.CoverTab[185036]++
											f := v.Float()
											if math.IsInf(f, 0) || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:512
		_go_fuzz_dep_.CoverTab[185039]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:512
		return math.IsNaN(f)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:512
		// _ = "end of CoverTab[185039]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:512
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:512
		_go_fuzz_dep_.CoverTab[185040]++
												e.error(&UnsupportedValueError{v, strconv.FormatFloat(f, 'g', -1, int(bits))})
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:513
		// _ = "end of CoverTab[185040]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:514
		_go_fuzz_dep_.CoverTab[185041]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:514
		// _ = "end of CoverTab[185041]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:514
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:514
	// _ = "end of CoverTab[185036]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:514
	_go_fuzz_dep_.CoverTab[185037]++
											b := strconv.AppendFloat(e.scratch[:0], f, 'g', -1, int(bits))
											if quoted {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:516
		_go_fuzz_dep_.CoverTab[185042]++
												e.WriteByte('"')
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:517
		// _ = "end of CoverTab[185042]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:518
		_go_fuzz_dep_.CoverTab[185043]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:518
		// _ = "end of CoverTab[185043]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:518
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:518
	// _ = "end of CoverTab[185037]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:518
	_go_fuzz_dep_.CoverTab[185038]++
											e.Write(b)
											if quoted {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:520
		_go_fuzz_dep_.CoverTab[185044]++
												e.WriteByte('"')
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:521
		// _ = "end of CoverTab[185044]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:522
		_go_fuzz_dep_.CoverTab[185045]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:522
		// _ = "end of CoverTab[185045]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:522
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:522
	// _ = "end of CoverTab[185038]"
}

var (
	float32Encoder	= (floatEncoder(32)).encode
	float64Encoder	= (floatEncoder(64)).encode
)

func stringEncoder(e *encodeState, v reflect.Value, quoted bool) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:530
	_go_fuzz_dep_.CoverTab[185046]++
											if v.Type() == numberType {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:531
		_go_fuzz_dep_.CoverTab[185048]++
												numStr := v.String()

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:535
		if numStr == "" {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:535
			_go_fuzz_dep_.CoverTab[185051]++
													numStr = "0"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:536
			// _ = "end of CoverTab[185051]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:537
			_go_fuzz_dep_.CoverTab[185052]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:537
			// _ = "end of CoverTab[185052]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:537
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:537
		// _ = "end of CoverTab[185048]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:537
		_go_fuzz_dep_.CoverTab[185049]++
												if !isValidNumber(numStr) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:538
			_go_fuzz_dep_.CoverTab[185053]++
													e.error(fmt.Errorf("json: invalid number literal %q", numStr))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:539
			// _ = "end of CoverTab[185053]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:540
			_go_fuzz_dep_.CoverTab[185054]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:540
			// _ = "end of CoverTab[185054]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:540
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:540
		// _ = "end of CoverTab[185049]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:540
		_go_fuzz_dep_.CoverTab[185050]++
												e.WriteString(numStr)
												return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:542
		// _ = "end of CoverTab[185050]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:543
		_go_fuzz_dep_.CoverTab[185055]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:543
		// _ = "end of CoverTab[185055]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:543
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:543
	// _ = "end of CoverTab[185046]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:543
	_go_fuzz_dep_.CoverTab[185047]++
											if quoted {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:544
		_go_fuzz_dep_.CoverTab[185056]++
												sb, err := Marshal(v.String())
												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:546
			_go_fuzz_dep_.CoverTab[185058]++
													e.error(err)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:547
			// _ = "end of CoverTab[185058]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:548
			_go_fuzz_dep_.CoverTab[185059]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:548
			// _ = "end of CoverTab[185059]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:548
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:548
		// _ = "end of CoverTab[185056]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:548
		_go_fuzz_dep_.CoverTab[185057]++
												e.string(string(sb))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:549
		// _ = "end of CoverTab[185057]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:550
		_go_fuzz_dep_.CoverTab[185060]++
												e.string(v.String())
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:551
		// _ = "end of CoverTab[185060]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:552
	// _ = "end of CoverTab[185047]"
}

func interfaceEncoder(e *encodeState, v reflect.Value, quoted bool) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:555
	_go_fuzz_dep_.CoverTab[185061]++
											if v.IsNil() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:556
		_go_fuzz_dep_.CoverTab[185063]++
												e.WriteString("null")
												return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:558
		// _ = "end of CoverTab[185063]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:559
		_go_fuzz_dep_.CoverTab[185064]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:559
		// _ = "end of CoverTab[185064]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:559
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:559
	// _ = "end of CoverTab[185061]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:559
	_go_fuzz_dep_.CoverTab[185062]++
											e.reflectValue(v.Elem())
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:560
	// _ = "end of CoverTab[185062]"
}

func unsupportedTypeEncoder(e *encodeState, v reflect.Value, quoted bool) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:563
	_go_fuzz_dep_.CoverTab[185065]++
											e.error(&UnsupportedTypeError{v.Type()})
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:564
	// _ = "end of CoverTab[185065]"
}

type structEncoder struct {
	fields		[]field
	fieldEncs	[]encoderFunc
}

func (se *structEncoder) encode(e *encodeState, v reflect.Value, quoted bool) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:572
	_go_fuzz_dep_.CoverTab[185066]++
											e.WriteByte('{')
											first := true
											for i, f := range se.fields {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:575
		_go_fuzz_dep_.CoverTab[185068]++
												fv := fieldByIndex(v, f.index)
												if !fv.IsValid() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:577
			_go_fuzz_dep_.CoverTab[185071]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:577
			return f.omitEmpty && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:577
				_go_fuzz_dep_.CoverTab[185072]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:577
				return isEmptyValue(fv)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:577
				// _ = "end of CoverTab[185072]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:577
			}()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:577
			// _ = "end of CoverTab[185071]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:577
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:577
			_go_fuzz_dep_.CoverTab[185073]++
													continue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:578
			// _ = "end of CoverTab[185073]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:579
			_go_fuzz_dep_.CoverTab[185074]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:579
			// _ = "end of CoverTab[185074]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:579
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:579
		// _ = "end of CoverTab[185068]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:579
		_go_fuzz_dep_.CoverTab[185069]++
												if first {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:580
			_go_fuzz_dep_.CoverTab[185075]++
													first = false
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:581
			// _ = "end of CoverTab[185075]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:582
			_go_fuzz_dep_.CoverTab[185076]++
													e.WriteByte(',')
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:583
			// _ = "end of CoverTab[185076]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:584
		// _ = "end of CoverTab[185069]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:584
		_go_fuzz_dep_.CoverTab[185070]++
												e.string(f.name)
												e.WriteByte(':')
												se.fieldEncs[i](e, fv, f.quoted)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:587
		// _ = "end of CoverTab[185070]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:588
	// _ = "end of CoverTab[185066]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:588
	_go_fuzz_dep_.CoverTab[185067]++
											e.WriteByte('}')
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:589
	// _ = "end of CoverTab[185067]"
}

func newStructEncoder(t reflect.Type) encoderFunc {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:592
	_go_fuzz_dep_.CoverTab[185077]++
											fields := cachedTypeFields(t)
											se := &structEncoder{
		fields:		fields,
		fieldEncs:	make([]encoderFunc, len(fields)),
	}
	for i, f := range fields {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:598
		_go_fuzz_dep_.CoverTab[185079]++
												se.fieldEncs[i] = typeEncoder(typeByIndex(t, f.index))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:599
		// _ = "end of CoverTab[185079]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:600
	// _ = "end of CoverTab[185077]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:600
	_go_fuzz_dep_.CoverTab[185078]++
											return se.encode
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:601
	// _ = "end of CoverTab[185078]"
}

type mapEncoder struct {
	elemEnc encoderFunc
}

func (me *mapEncoder) encode(e *encodeState, v reflect.Value, _ bool) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:608
	_go_fuzz_dep_.CoverTab[185080]++
											if v.IsNil() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:609
		_go_fuzz_dep_.CoverTab[185083]++
												e.WriteString("null")
												return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:611
		// _ = "end of CoverTab[185083]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:612
		_go_fuzz_dep_.CoverTab[185084]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:612
		// _ = "end of CoverTab[185084]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:612
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:612
	// _ = "end of CoverTab[185080]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:612
	_go_fuzz_dep_.CoverTab[185081]++
											e.WriteByte('{')
											var sv stringValues = v.MapKeys()
											sort.Sort(sv)
											for i, k := range sv {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:616
		_go_fuzz_dep_.CoverTab[185085]++
												if i > 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:617
			_go_fuzz_dep_.CoverTab[185087]++
													e.WriteByte(',')
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:618
			// _ = "end of CoverTab[185087]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:619
			_go_fuzz_dep_.CoverTab[185088]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:619
			// _ = "end of CoverTab[185088]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:619
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:619
		// _ = "end of CoverTab[185085]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:619
		_go_fuzz_dep_.CoverTab[185086]++
												e.string(k.String())
												e.WriteByte(':')
												me.elemEnc(e, v.MapIndex(k), false)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:622
		// _ = "end of CoverTab[185086]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:623
	// _ = "end of CoverTab[185081]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:623
	_go_fuzz_dep_.CoverTab[185082]++
											e.WriteByte('}')
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:624
	// _ = "end of CoverTab[185082]"
}

func newMapEncoder(t reflect.Type) encoderFunc {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:627
	_go_fuzz_dep_.CoverTab[185089]++
											if t.Key().Kind() != reflect.String {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:628
		_go_fuzz_dep_.CoverTab[185091]++
												return unsupportedTypeEncoder
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:629
		// _ = "end of CoverTab[185091]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:630
		_go_fuzz_dep_.CoverTab[185092]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:630
		// _ = "end of CoverTab[185092]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:630
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:630
	// _ = "end of CoverTab[185089]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:630
	_go_fuzz_dep_.CoverTab[185090]++
											me := &mapEncoder{typeEncoder(t.Elem())}
											return me.encode
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:632
	// _ = "end of CoverTab[185090]"
}

func encodeByteSlice(e *encodeState, v reflect.Value, _ bool) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:635
	_go_fuzz_dep_.CoverTab[185093]++
											if v.IsNil() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:636
		_go_fuzz_dep_.CoverTab[185096]++
												e.WriteString("null")
												return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:638
		// _ = "end of CoverTab[185096]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:639
		_go_fuzz_dep_.CoverTab[185097]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:639
		// _ = "end of CoverTab[185097]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:639
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:639
	// _ = "end of CoverTab[185093]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:639
	_go_fuzz_dep_.CoverTab[185094]++
											s := v.Bytes()
											e.WriteByte('"')
											if len(s) < 1024 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:642
		_go_fuzz_dep_.CoverTab[185098]++

												dst := make([]byte, base64.StdEncoding.EncodedLen(len(s)))
												base64.StdEncoding.Encode(dst, s)
												e.Write(dst)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:646
		// _ = "end of CoverTab[185098]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:647
		_go_fuzz_dep_.CoverTab[185099]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:650
		enc := base64.NewEncoder(base64.StdEncoding, e)
												enc.Write(s)
												enc.Close()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:652
		// _ = "end of CoverTab[185099]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:653
	// _ = "end of CoverTab[185094]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:653
	_go_fuzz_dep_.CoverTab[185095]++
											e.WriteByte('"')
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:654
	// _ = "end of CoverTab[185095]"
}

// sliceEncoder just wraps an arrayEncoder, checking to make sure the value isn't nil.
type sliceEncoder struct {
	arrayEnc encoderFunc
}

func (se *sliceEncoder) encode(e *encodeState, v reflect.Value, _ bool) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:662
	_go_fuzz_dep_.CoverTab[185100]++
											if v.IsNil() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:663
		_go_fuzz_dep_.CoverTab[185102]++
												e.WriteString("null")
												return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:665
		// _ = "end of CoverTab[185102]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:666
		_go_fuzz_dep_.CoverTab[185103]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:666
		// _ = "end of CoverTab[185103]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:666
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:666
	// _ = "end of CoverTab[185100]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:666
	_go_fuzz_dep_.CoverTab[185101]++
											se.arrayEnc(e, v, false)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:667
	// _ = "end of CoverTab[185101]"
}

func newSliceEncoder(t reflect.Type) encoderFunc {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:670
	_go_fuzz_dep_.CoverTab[185104]++

											if t.Elem().Kind() == reflect.Uint8 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:672
		_go_fuzz_dep_.CoverTab[185106]++
												return encodeByteSlice
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:673
		// _ = "end of CoverTab[185106]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:674
		_go_fuzz_dep_.CoverTab[185107]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:674
		// _ = "end of CoverTab[185107]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:674
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:674
	// _ = "end of CoverTab[185104]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:674
	_go_fuzz_dep_.CoverTab[185105]++
											enc := &sliceEncoder{newArrayEncoder(t)}
											return enc.encode
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:676
	// _ = "end of CoverTab[185105]"
}

type arrayEncoder struct {
	elemEnc encoderFunc
}

func (ae *arrayEncoder) encode(e *encodeState, v reflect.Value, _ bool) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:683
	_go_fuzz_dep_.CoverTab[185108]++
											e.WriteByte('[')
											n := v.Len()
											for i := 0; i < n; i++ {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:686
		_go_fuzz_dep_.CoverTab[185110]++
												if i > 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:687
			_go_fuzz_dep_.CoverTab[185112]++
													e.WriteByte(',')
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:688
			// _ = "end of CoverTab[185112]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:689
			_go_fuzz_dep_.CoverTab[185113]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:689
			// _ = "end of CoverTab[185113]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:689
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:689
		// _ = "end of CoverTab[185110]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:689
		_go_fuzz_dep_.CoverTab[185111]++
												ae.elemEnc(e, v.Index(i), false)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:690
		// _ = "end of CoverTab[185111]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:691
	// _ = "end of CoverTab[185108]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:691
	_go_fuzz_dep_.CoverTab[185109]++
											e.WriteByte(']')
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:692
	// _ = "end of CoverTab[185109]"
}

func newArrayEncoder(t reflect.Type) encoderFunc {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:695
	_go_fuzz_dep_.CoverTab[185114]++
											enc := &arrayEncoder{typeEncoder(t.Elem())}
											return enc.encode
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:697
	// _ = "end of CoverTab[185114]"
}

type ptrEncoder struct {
	elemEnc encoderFunc
}

func (pe *ptrEncoder) encode(e *encodeState, v reflect.Value, quoted bool) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:704
	_go_fuzz_dep_.CoverTab[185115]++
											if v.IsNil() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:705
		_go_fuzz_dep_.CoverTab[185117]++
												e.WriteString("null")
												return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:707
		// _ = "end of CoverTab[185117]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:708
		_go_fuzz_dep_.CoverTab[185118]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:708
		// _ = "end of CoverTab[185118]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:708
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:708
	// _ = "end of CoverTab[185115]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:708
	_go_fuzz_dep_.CoverTab[185116]++
											pe.elemEnc(e, v.Elem(), quoted)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:709
	// _ = "end of CoverTab[185116]"
}

func newPtrEncoder(t reflect.Type) encoderFunc {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:712
	_go_fuzz_dep_.CoverTab[185119]++
											enc := &ptrEncoder{typeEncoder(t.Elem())}
											return enc.encode
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:714
	// _ = "end of CoverTab[185119]"
}

type condAddrEncoder struct {
	canAddrEnc, elseEnc encoderFunc
}

func (ce *condAddrEncoder) encode(e *encodeState, v reflect.Value, quoted bool) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:721
	_go_fuzz_dep_.CoverTab[185120]++
											if v.CanAddr() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:722
		_go_fuzz_dep_.CoverTab[185121]++
												ce.canAddrEnc(e, v, quoted)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:723
		// _ = "end of CoverTab[185121]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:724
		_go_fuzz_dep_.CoverTab[185122]++
												ce.elseEnc(e, v, quoted)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:725
		// _ = "end of CoverTab[185122]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:726
	// _ = "end of CoverTab[185120]"
}

// newCondAddrEncoder returns an encoder that checks whether its value
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:729
// CanAddr and delegates to canAddrEnc if so, else to elseEnc.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:731
func newCondAddrEncoder(canAddrEnc, elseEnc encoderFunc) encoderFunc {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:731
	_go_fuzz_dep_.CoverTab[185123]++
											enc := &condAddrEncoder{canAddrEnc: canAddrEnc, elseEnc: elseEnc}
											return enc.encode
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:733
	// _ = "end of CoverTab[185123]"
}

func isValidTag(s string) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:736
	_go_fuzz_dep_.CoverTab[185124]++
											if s == "" {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:737
		_go_fuzz_dep_.CoverTab[185127]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:738
		// _ = "end of CoverTab[185127]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:739
		_go_fuzz_dep_.CoverTab[185128]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:739
		// _ = "end of CoverTab[185128]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:739
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:739
	// _ = "end of CoverTab[185124]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:739
	_go_fuzz_dep_.CoverTab[185125]++
											for _, c := range s {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:740
		_go_fuzz_dep_.CoverTab[185129]++
												switch {
		case strings.ContainsRune("!#$%&()*+-./:<=>?@[]^_{|}~ ", c):
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:742
			_go_fuzz_dep_.CoverTab[185130]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:742
			// _ = "end of CoverTab[185130]"

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:746
		default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:746
			_go_fuzz_dep_.CoverTab[185131]++
													if !unicode.IsLetter(c) && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:747
				_go_fuzz_dep_.CoverTab[185132]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:747
				return !unicode.IsDigit(c)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:747
				// _ = "end of CoverTab[185132]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:747
			}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:747
				_go_fuzz_dep_.CoverTab[185133]++
														return false
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:748
				// _ = "end of CoverTab[185133]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:749
				_go_fuzz_dep_.CoverTab[185134]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:749
				// _ = "end of CoverTab[185134]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:749
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:749
			// _ = "end of CoverTab[185131]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:750
		// _ = "end of CoverTab[185129]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:751
	// _ = "end of CoverTab[185125]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:751
	_go_fuzz_dep_.CoverTab[185126]++
											return true
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:752
	// _ = "end of CoverTab[185126]"
}

func fieldByIndex(v reflect.Value, index []int) reflect.Value {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:755
	_go_fuzz_dep_.CoverTab[185135]++
											for _, i := range index {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:756
		_go_fuzz_dep_.CoverTab[185137]++
												if v.Kind() == reflect.Ptr {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:757
			_go_fuzz_dep_.CoverTab[185139]++
													if v.IsNil() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:758
				_go_fuzz_dep_.CoverTab[185141]++
														return reflect.Value{}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:759
				// _ = "end of CoverTab[185141]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:760
				_go_fuzz_dep_.CoverTab[185142]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:760
				// _ = "end of CoverTab[185142]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:760
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:760
			// _ = "end of CoverTab[185139]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:760
			_go_fuzz_dep_.CoverTab[185140]++
													v = v.Elem()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:761
			// _ = "end of CoverTab[185140]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:762
			_go_fuzz_dep_.CoverTab[185143]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:762
			// _ = "end of CoverTab[185143]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:762
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:762
		// _ = "end of CoverTab[185137]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:762
		_go_fuzz_dep_.CoverTab[185138]++
												v = v.Field(i)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:763
		// _ = "end of CoverTab[185138]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:764
	// _ = "end of CoverTab[185135]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:764
	_go_fuzz_dep_.CoverTab[185136]++
											return v
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:765
	// _ = "end of CoverTab[185136]"
}

func typeByIndex(t reflect.Type, index []int) reflect.Type {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:768
	_go_fuzz_dep_.CoverTab[185144]++
											for _, i := range index {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:769
		_go_fuzz_dep_.CoverTab[185146]++
												if t.Kind() == reflect.Ptr {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:770
			_go_fuzz_dep_.CoverTab[185148]++
													t = t.Elem()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:771
			// _ = "end of CoverTab[185148]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:772
			_go_fuzz_dep_.CoverTab[185149]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:772
			// _ = "end of CoverTab[185149]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:772
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:772
		// _ = "end of CoverTab[185146]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:772
		_go_fuzz_dep_.CoverTab[185147]++
												t = t.Field(i).Type
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:773
		// _ = "end of CoverTab[185147]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:774
	// _ = "end of CoverTab[185144]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:774
	_go_fuzz_dep_.CoverTab[185145]++
											return t
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:775
	// _ = "end of CoverTab[185145]"
}

// stringValues is a slice of reflect.Value holding *reflect.StringValue.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:778
// It implements the methods to sort by string.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:780
type stringValues []reflect.Value

func (sv stringValues) Len() int {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:782
	_go_fuzz_dep_.CoverTab[185150]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:782
	return len(sv)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:782
	// _ = "end of CoverTab[185150]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:782
}
func (sv stringValues) Swap(i, j int) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:783
	_go_fuzz_dep_.CoverTab[185151]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:783
	sv[i], sv[j] = sv[j], sv[i]
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:783
	// _ = "end of CoverTab[185151]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:783
}
func (sv stringValues) Less(i, j int) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:784
	_go_fuzz_dep_.CoverTab[185152]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:784
	return sv.get(i) < sv.get(j)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:784
	// _ = "end of CoverTab[185152]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:784
}
func (sv stringValues) get(i int) string {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:785
	_go_fuzz_dep_.CoverTab[185153]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:785
	return sv[i].String()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:785
	// _ = "end of CoverTab[185153]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:785
}

// NOTE: keep in sync with stringBytes below.
func (e *encodeState) string(s string) int {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:788
	_go_fuzz_dep_.CoverTab[185154]++
											len0 := e.Len()
											e.WriteByte('"')
											start := 0
											for i := 0; i < len(s); {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:792
		_go_fuzz_dep_.CoverTab[185157]++
												if b := s[i]; b < utf8.RuneSelf {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:793
			_go_fuzz_dep_.CoverTab[185161]++
													if 0x20 <= b && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:794
				_go_fuzz_dep_.CoverTab[185165]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:794
				return b != '\\'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:794
				// _ = "end of CoverTab[185165]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:794
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:794
				_go_fuzz_dep_.CoverTab[185166]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:794
				return b != '"'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:794
				// _ = "end of CoverTab[185166]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:794
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:794
				_go_fuzz_dep_.CoverTab[185167]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:794
				return b != '<'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:794
				// _ = "end of CoverTab[185167]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:794
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:794
				_go_fuzz_dep_.CoverTab[185168]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:794
				return b != '>'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:794
				// _ = "end of CoverTab[185168]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:794
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:794
				_go_fuzz_dep_.CoverTab[185169]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:794
				return b != '&'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:794
				// _ = "end of CoverTab[185169]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:794
			}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:794
				_go_fuzz_dep_.CoverTab[185170]++
														i++
														continue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:796
				// _ = "end of CoverTab[185170]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:797
				_go_fuzz_dep_.CoverTab[185171]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:797
				// _ = "end of CoverTab[185171]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:797
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:797
			// _ = "end of CoverTab[185161]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:797
			_go_fuzz_dep_.CoverTab[185162]++
													if start < i {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:798
				_go_fuzz_dep_.CoverTab[185172]++
														e.WriteString(s[start:i])
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:799
				// _ = "end of CoverTab[185172]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:800
				_go_fuzz_dep_.CoverTab[185173]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:800
				// _ = "end of CoverTab[185173]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:800
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:800
			// _ = "end of CoverTab[185162]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:800
			_go_fuzz_dep_.CoverTab[185163]++
													switch b {
			case '\\', '"':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:802
				_go_fuzz_dep_.CoverTab[185174]++
														e.WriteByte('\\')
														e.WriteByte(b)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:804
				// _ = "end of CoverTab[185174]"
			case '\n':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:805
				_go_fuzz_dep_.CoverTab[185175]++
														e.WriteByte('\\')
														e.WriteByte('n')
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:807
				// _ = "end of CoverTab[185175]"
			case '\r':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:808
				_go_fuzz_dep_.CoverTab[185176]++
														e.WriteByte('\\')
														e.WriteByte('r')
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:810
				// _ = "end of CoverTab[185176]"
			case '\t':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:811
				_go_fuzz_dep_.CoverTab[185177]++
														e.WriteByte('\\')
														e.WriteByte('t')
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:813
				// _ = "end of CoverTab[185177]"
			default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:814
				_go_fuzz_dep_.CoverTab[185178]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:819
				e.WriteString(`\u00`)
														e.WriteByte(hex[b>>4])
														e.WriteByte(hex[b&0xF])
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:821
				// _ = "end of CoverTab[185178]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:822
			// _ = "end of CoverTab[185163]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:822
			_go_fuzz_dep_.CoverTab[185164]++
													i++
													start = i
													continue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:825
			// _ = "end of CoverTab[185164]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:826
			_go_fuzz_dep_.CoverTab[185179]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:826
			// _ = "end of CoverTab[185179]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:826
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:826
		// _ = "end of CoverTab[185157]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:826
		_go_fuzz_dep_.CoverTab[185158]++
												c, size := utf8.DecodeRuneInString(s[i:])
												if c == utf8.RuneError && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:828
			_go_fuzz_dep_.CoverTab[185180]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:828
			return size == 1
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:828
			// _ = "end of CoverTab[185180]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:828
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:828
			_go_fuzz_dep_.CoverTab[185181]++
													if start < i {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:829
				_go_fuzz_dep_.CoverTab[185183]++
														e.WriteString(s[start:i])
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:830
				// _ = "end of CoverTab[185183]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:831
				_go_fuzz_dep_.CoverTab[185184]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:831
				// _ = "end of CoverTab[185184]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:831
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:831
			// _ = "end of CoverTab[185181]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:831
			_go_fuzz_dep_.CoverTab[185182]++
													e.WriteString(`\ufffd`)
													i += size
													start = i
													continue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:835
			// _ = "end of CoverTab[185182]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:836
			_go_fuzz_dep_.CoverTab[185185]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:836
			// _ = "end of CoverTab[185185]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:836
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:836
		// _ = "end of CoverTab[185158]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:836
		_go_fuzz_dep_.CoverTab[185159]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:844
		if c == '\u2028' || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:844
			_go_fuzz_dep_.CoverTab[185186]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:844
			return c == '\u2029'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:844
			// _ = "end of CoverTab[185186]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:844
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:844
			_go_fuzz_dep_.CoverTab[185187]++
													if start < i {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:845
				_go_fuzz_dep_.CoverTab[185189]++
														e.WriteString(s[start:i])
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:846
				// _ = "end of CoverTab[185189]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:847
				_go_fuzz_dep_.CoverTab[185190]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:847
				// _ = "end of CoverTab[185190]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:847
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:847
			// _ = "end of CoverTab[185187]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:847
			_go_fuzz_dep_.CoverTab[185188]++
													e.WriteString(`\u202`)
													e.WriteByte(hex[c&0xF])
													i += size
													start = i
													continue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:852
			// _ = "end of CoverTab[185188]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:853
			_go_fuzz_dep_.CoverTab[185191]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:853
			// _ = "end of CoverTab[185191]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:853
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:853
		// _ = "end of CoverTab[185159]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:853
		_go_fuzz_dep_.CoverTab[185160]++
												i += size
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:854
		// _ = "end of CoverTab[185160]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:855
	// _ = "end of CoverTab[185154]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:855
	_go_fuzz_dep_.CoverTab[185155]++
											if start < len(s) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:856
		_go_fuzz_dep_.CoverTab[185192]++
												e.WriteString(s[start:])
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:857
		// _ = "end of CoverTab[185192]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:858
		_go_fuzz_dep_.CoverTab[185193]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:858
		// _ = "end of CoverTab[185193]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:858
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:858
	// _ = "end of CoverTab[185155]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:858
	_go_fuzz_dep_.CoverTab[185156]++
											e.WriteByte('"')
											return e.Len() - len0
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:860
	// _ = "end of CoverTab[185156]"
}

// NOTE: keep in sync with string above.
func (e *encodeState) stringBytes(s []byte) int {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:864
	_go_fuzz_dep_.CoverTab[185194]++
											len0 := e.Len()
											e.WriteByte('"')
											start := 0
											for i := 0; i < len(s); {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:868
		_go_fuzz_dep_.CoverTab[185197]++
												if b := s[i]; b < utf8.RuneSelf {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:869
			_go_fuzz_dep_.CoverTab[185201]++
													if 0x20 <= b && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:870
				_go_fuzz_dep_.CoverTab[185205]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:870
				return b != '\\'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:870
				// _ = "end of CoverTab[185205]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:870
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:870
				_go_fuzz_dep_.CoverTab[185206]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:870
				return b != '"'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:870
				// _ = "end of CoverTab[185206]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:870
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:870
				_go_fuzz_dep_.CoverTab[185207]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:870
				return b != '<'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:870
				// _ = "end of CoverTab[185207]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:870
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:870
				_go_fuzz_dep_.CoverTab[185208]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:870
				return b != '>'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:870
				// _ = "end of CoverTab[185208]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:870
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:870
				_go_fuzz_dep_.CoverTab[185209]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:870
				return b != '&'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:870
				// _ = "end of CoverTab[185209]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:870
			}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:870
				_go_fuzz_dep_.CoverTab[185210]++
														i++
														continue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:872
				// _ = "end of CoverTab[185210]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:873
				_go_fuzz_dep_.CoverTab[185211]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:873
				// _ = "end of CoverTab[185211]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:873
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:873
			// _ = "end of CoverTab[185201]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:873
			_go_fuzz_dep_.CoverTab[185202]++
													if start < i {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:874
				_go_fuzz_dep_.CoverTab[185212]++
														e.Write(s[start:i])
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:875
				// _ = "end of CoverTab[185212]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:876
				_go_fuzz_dep_.CoverTab[185213]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:876
				// _ = "end of CoverTab[185213]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:876
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:876
			// _ = "end of CoverTab[185202]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:876
			_go_fuzz_dep_.CoverTab[185203]++
													switch b {
			case '\\', '"':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:878
				_go_fuzz_dep_.CoverTab[185214]++
														e.WriteByte('\\')
														e.WriteByte(b)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:880
				// _ = "end of CoverTab[185214]"
			case '\n':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:881
				_go_fuzz_dep_.CoverTab[185215]++
														e.WriteByte('\\')
														e.WriteByte('n')
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:883
				// _ = "end of CoverTab[185215]"
			case '\r':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:884
				_go_fuzz_dep_.CoverTab[185216]++
														e.WriteByte('\\')
														e.WriteByte('r')
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:886
				// _ = "end of CoverTab[185216]"
			case '\t':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:887
				_go_fuzz_dep_.CoverTab[185217]++
														e.WriteByte('\\')
														e.WriteByte('t')
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:889
				// _ = "end of CoverTab[185217]"
			default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:890
				_go_fuzz_dep_.CoverTab[185218]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:895
				e.WriteString(`\u00`)
														e.WriteByte(hex[b>>4])
														e.WriteByte(hex[b&0xF])
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:897
				// _ = "end of CoverTab[185218]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:898
			// _ = "end of CoverTab[185203]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:898
			_go_fuzz_dep_.CoverTab[185204]++
													i++
													start = i
													continue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:901
			// _ = "end of CoverTab[185204]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:902
			_go_fuzz_dep_.CoverTab[185219]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:902
			// _ = "end of CoverTab[185219]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:902
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:902
		// _ = "end of CoverTab[185197]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:902
		_go_fuzz_dep_.CoverTab[185198]++
												c, size := utf8.DecodeRune(s[i:])
												if c == utf8.RuneError && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:904
			_go_fuzz_dep_.CoverTab[185220]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:904
			return size == 1
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:904
			// _ = "end of CoverTab[185220]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:904
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:904
			_go_fuzz_dep_.CoverTab[185221]++
													if start < i {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:905
				_go_fuzz_dep_.CoverTab[185223]++
														e.Write(s[start:i])
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:906
				// _ = "end of CoverTab[185223]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:907
				_go_fuzz_dep_.CoverTab[185224]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:907
				// _ = "end of CoverTab[185224]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:907
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:907
			// _ = "end of CoverTab[185221]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:907
			_go_fuzz_dep_.CoverTab[185222]++
													e.WriteString(`\ufffd`)
													i += size
													start = i
													continue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:911
			// _ = "end of CoverTab[185222]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:912
			_go_fuzz_dep_.CoverTab[185225]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:912
			// _ = "end of CoverTab[185225]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:912
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:912
		// _ = "end of CoverTab[185198]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:912
		_go_fuzz_dep_.CoverTab[185199]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:920
		if c == '\u2028' || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:920
			_go_fuzz_dep_.CoverTab[185226]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:920
			return c == '\u2029'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:920
			// _ = "end of CoverTab[185226]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:920
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:920
			_go_fuzz_dep_.CoverTab[185227]++
													if start < i {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:921
				_go_fuzz_dep_.CoverTab[185229]++
														e.Write(s[start:i])
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:922
				// _ = "end of CoverTab[185229]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:923
				_go_fuzz_dep_.CoverTab[185230]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:923
				// _ = "end of CoverTab[185230]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:923
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:923
			// _ = "end of CoverTab[185227]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:923
			_go_fuzz_dep_.CoverTab[185228]++
													e.WriteString(`\u202`)
													e.WriteByte(hex[c&0xF])
													i += size
													start = i
													continue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:928
			// _ = "end of CoverTab[185228]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:929
			_go_fuzz_dep_.CoverTab[185231]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:929
			// _ = "end of CoverTab[185231]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:929
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:929
		// _ = "end of CoverTab[185199]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:929
		_go_fuzz_dep_.CoverTab[185200]++
												i += size
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:930
		// _ = "end of CoverTab[185200]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:931
	// _ = "end of CoverTab[185194]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:931
	_go_fuzz_dep_.CoverTab[185195]++
											if start < len(s) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:932
		_go_fuzz_dep_.CoverTab[185232]++
												e.Write(s[start:])
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:933
		// _ = "end of CoverTab[185232]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:934
		_go_fuzz_dep_.CoverTab[185233]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:934
		// _ = "end of CoverTab[185233]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:934
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:934
	// _ = "end of CoverTab[185195]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:934
	_go_fuzz_dep_.CoverTab[185196]++
											e.WriteByte('"')
											return e.Len() - len0
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:936
	// _ = "end of CoverTab[185196]"
}

// A field represents a single field found in a struct.
type field struct {
	name		string
	nameBytes	[]byte	// []byte(name)

	tag		bool
	index		[]int
	typ		reflect.Type
	omitEmpty	bool
	quoted		bool
}

func fillField(f field) field {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:951
	_go_fuzz_dep_.CoverTab[185234]++
											f.nameBytes = []byte(f.name)
											return f
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:953
	// _ = "end of CoverTab[185234]"
}

// byName sorts field by name, breaking ties with depth,
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:956
// then breaking ties with "name came from json tag", then
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:956
// breaking ties with index sequence.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:959
type byName []field

func (x byName) Len() int {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:961
	_go_fuzz_dep_.CoverTab[185235]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:961
	return len(x)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:961
	// _ = "end of CoverTab[185235]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:961
}

func (x byName) Swap(i, j int) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:963
	_go_fuzz_dep_.CoverTab[185236]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:963
	x[i], x[j] = x[j], x[i]
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:963
	// _ = "end of CoverTab[185236]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:963
}

func (x byName) Less(i, j int) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:965
	_go_fuzz_dep_.CoverTab[185237]++
											if x[i].name != x[j].name {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:966
		_go_fuzz_dep_.CoverTab[185241]++
												return x[i].name < x[j].name
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:967
		// _ = "end of CoverTab[185241]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:968
		_go_fuzz_dep_.CoverTab[185242]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:968
		// _ = "end of CoverTab[185242]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:968
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:968
	// _ = "end of CoverTab[185237]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:968
	_go_fuzz_dep_.CoverTab[185238]++
											if len(x[i].index) != len(x[j].index) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:969
		_go_fuzz_dep_.CoverTab[185243]++
												return len(x[i].index) < len(x[j].index)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:970
		// _ = "end of CoverTab[185243]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:971
		_go_fuzz_dep_.CoverTab[185244]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:971
		// _ = "end of CoverTab[185244]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:971
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:971
	// _ = "end of CoverTab[185238]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:971
	_go_fuzz_dep_.CoverTab[185239]++
											if x[i].tag != x[j].tag {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:972
		_go_fuzz_dep_.CoverTab[185245]++
												return x[i].tag
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:973
		// _ = "end of CoverTab[185245]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:974
		_go_fuzz_dep_.CoverTab[185246]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:974
		// _ = "end of CoverTab[185246]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:974
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:974
	// _ = "end of CoverTab[185239]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:974
	_go_fuzz_dep_.CoverTab[185240]++
											return byIndex(x).Less(i, j)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:975
	// _ = "end of CoverTab[185240]"
}

// byIndex sorts field by index sequence.
type byIndex []field

func (x byIndex) Len() int {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:981
	_go_fuzz_dep_.CoverTab[185247]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:981
	return len(x)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:981
	// _ = "end of CoverTab[185247]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:981
}

func (x byIndex) Swap(i, j int) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:983
	_go_fuzz_dep_.CoverTab[185248]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:983
	x[i], x[j] = x[j], x[i]
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:983
	// _ = "end of CoverTab[185248]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:983
}

func (x byIndex) Less(i, j int) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:985
	_go_fuzz_dep_.CoverTab[185249]++
											for k, xik := range x[i].index {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:986
		_go_fuzz_dep_.CoverTab[185251]++
												if k >= len(x[j].index) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:987
			_go_fuzz_dep_.CoverTab[185253]++
													return false
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:988
			// _ = "end of CoverTab[185253]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:989
			_go_fuzz_dep_.CoverTab[185254]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:989
			// _ = "end of CoverTab[185254]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:989
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:989
		// _ = "end of CoverTab[185251]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:989
		_go_fuzz_dep_.CoverTab[185252]++
												if xik != x[j].index[k] {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:990
			_go_fuzz_dep_.CoverTab[185255]++
													return xik < x[j].index[k]
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:991
			// _ = "end of CoverTab[185255]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:992
			_go_fuzz_dep_.CoverTab[185256]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:992
			// _ = "end of CoverTab[185256]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:992
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:992
		// _ = "end of CoverTab[185252]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:993
	// _ = "end of CoverTab[185249]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:993
	_go_fuzz_dep_.CoverTab[185250]++
											return len(x[i].index) < len(x[j].index)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:994
	// _ = "end of CoverTab[185250]"
}

// typeFields returns a list of fields that JSON should recognize for the given type.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:997
// The algorithm is breadth-first search over the set of structs to include - the top struct
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:997
// and then any reachable anonymous structs.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1000
func typeFields(t reflect.Type) []field {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1000
	_go_fuzz_dep_.CoverTab[185257]++

												current := []field{}
												next := []field{{typ: t}}

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1006
	count := map[reflect.Type]int{}
												nextCount := map[reflect.Type]int{}

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1010
	visited := map[reflect.Type]bool{}

	// Fields found.
	var fields []field

	for len(next) > 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1015
		_go_fuzz_dep_.CoverTab[185260]++
													current, next = next, current[:0]
													count, nextCount = nextCount, map[reflect.Type]int{}

													for _, f := range current {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1019
			_go_fuzz_dep_.CoverTab[185261]++
														if visited[f.typ] {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1020
				_go_fuzz_dep_.CoverTab[185263]++
															continue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1021
				// _ = "end of CoverTab[185263]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1022
				_go_fuzz_dep_.CoverTab[185264]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1022
				// _ = "end of CoverTab[185264]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1022
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1022
			// _ = "end of CoverTab[185261]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1022
			_go_fuzz_dep_.CoverTab[185262]++
														visited[f.typ] = true

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1026
			for i := 0; i < f.typ.NumField(); i++ {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1026
				_go_fuzz_dep_.CoverTab[185265]++
															sf := f.typ.Field(i)
															if sf.PkgPath != "" && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1028
					_go_fuzz_dep_.CoverTab[185272]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1028
					return !sf.Anonymous
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1028
					// _ = "end of CoverTab[185272]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1028
				}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1028
					_go_fuzz_dep_.CoverTab[185273]++
																continue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1029
					// _ = "end of CoverTab[185273]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1030
					_go_fuzz_dep_.CoverTab[185274]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1030
					// _ = "end of CoverTab[185274]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1030
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1030
				// _ = "end of CoverTab[185265]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1030
				_go_fuzz_dep_.CoverTab[185266]++
															tag := sf.Tag.Get("json")
															if tag == "-" {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1032
					_go_fuzz_dep_.CoverTab[185275]++
																continue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1033
					// _ = "end of CoverTab[185275]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1034
					_go_fuzz_dep_.CoverTab[185276]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1034
					// _ = "end of CoverTab[185276]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1034
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1034
				// _ = "end of CoverTab[185266]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1034
				_go_fuzz_dep_.CoverTab[185267]++
															name, opts := parseTag(tag)
															if !isValidTag(name) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1036
					_go_fuzz_dep_.CoverTab[185277]++
																name = ""
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1037
					// _ = "end of CoverTab[185277]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1038
					_go_fuzz_dep_.CoverTab[185278]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1038
					// _ = "end of CoverTab[185278]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1038
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1038
				// _ = "end of CoverTab[185267]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1038
				_go_fuzz_dep_.CoverTab[185268]++
															index := make([]int, len(f.index)+1)
															copy(index, f.index)
															index[len(f.index)] = i

															ft := sf.Type
															if ft.Name() == "" && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1044
					_go_fuzz_dep_.CoverTab[185279]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1044
					return ft.Kind() == reflect.Ptr
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1044
					// _ = "end of CoverTab[185279]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1044
				}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1044
					_go_fuzz_dep_.CoverTab[185280]++

																ft = ft.Elem()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1046
					// _ = "end of CoverTab[185280]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1047
					_go_fuzz_dep_.CoverTab[185281]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1047
					// _ = "end of CoverTab[185281]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1047
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1047
				// _ = "end of CoverTab[185268]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1047
				_go_fuzz_dep_.CoverTab[185269]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1050
				quoted := false
				if opts.Contains("string") {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1051
					_go_fuzz_dep_.CoverTab[185282]++
																switch ft.Kind() {
					case reflect.Bool,
						reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
						reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
						reflect.Float32, reflect.Float64,
						reflect.String:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1057
						_go_fuzz_dep_.CoverTab[185283]++
																	quoted = true
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1058
						// _ = "end of CoverTab[185283]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1058
					default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1058
						_go_fuzz_dep_.CoverTab[185284]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1058
						// _ = "end of CoverTab[185284]"
					}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1059
					// _ = "end of CoverTab[185282]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1060
					_go_fuzz_dep_.CoverTab[185285]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1060
					// _ = "end of CoverTab[185285]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1060
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1060
				// _ = "end of CoverTab[185269]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1060
				_go_fuzz_dep_.CoverTab[185270]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1063
				if name != "" || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1063
					_go_fuzz_dep_.CoverTab[185286]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1063
					return !sf.Anonymous
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1063
					// _ = "end of CoverTab[185286]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1063
				}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1063
					_go_fuzz_dep_.CoverTab[185287]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1063
					return ft.Kind() != reflect.Struct
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1063
					// _ = "end of CoverTab[185287]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1063
				}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1063
					_go_fuzz_dep_.CoverTab[185288]++
																tagged := name != ""
																if name == "" {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1065
						_go_fuzz_dep_.CoverTab[185291]++
																	name = sf.Name
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1066
						// _ = "end of CoverTab[185291]"
					} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1067
						_go_fuzz_dep_.CoverTab[185292]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1067
						// _ = "end of CoverTab[185292]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1067
					}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1067
					// _ = "end of CoverTab[185288]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1067
					_go_fuzz_dep_.CoverTab[185289]++
																fields = append(fields, fillField(field{
						name:		name,
						tag:		tagged,
						index:		index,
						typ:		ft,
						omitEmpty:	opts.Contains("omitempty"),
						quoted:		quoted,
					}))
					if count[f.typ] > 1 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1076
						_go_fuzz_dep_.CoverTab[185293]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1081
						fields = append(fields, fields[len(fields)-1])
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1081
						// _ = "end of CoverTab[185293]"
					} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1082
						_go_fuzz_dep_.CoverTab[185294]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1082
						// _ = "end of CoverTab[185294]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1082
					}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1082
					// _ = "end of CoverTab[185289]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1082
					_go_fuzz_dep_.CoverTab[185290]++
																continue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1083
					// _ = "end of CoverTab[185290]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1084
					_go_fuzz_dep_.CoverTab[185295]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1084
					// _ = "end of CoverTab[185295]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1084
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1084
				// _ = "end of CoverTab[185270]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1084
				_go_fuzz_dep_.CoverTab[185271]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1087
				nextCount[ft]++
				if nextCount[ft] == 1 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1088
					_go_fuzz_dep_.CoverTab[185296]++
																next = append(next, fillField(field{name: ft.Name(), index: index, typ: ft}))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1089
					// _ = "end of CoverTab[185296]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1090
					_go_fuzz_dep_.CoverTab[185297]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1090
					// _ = "end of CoverTab[185297]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1090
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1090
				// _ = "end of CoverTab[185271]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1091
			// _ = "end of CoverTab[185262]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1092
		// _ = "end of CoverTab[185260]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1093
	// _ = "end of CoverTab[185257]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1093
	_go_fuzz_dep_.CoverTab[185258]++

												sort.Sort(byName(fields))

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1103
	out := fields[:0]
	for advance, i := 0, 0; i < len(fields); i += advance {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1104
		_go_fuzz_dep_.CoverTab[185298]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1107
		fi := fields[i]
		name := fi.name
		for advance = 1; i+advance < len(fields); advance++ {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1109
			_go_fuzz_dep_.CoverTab[185301]++
														fj := fields[i+advance]
														if fj.name != name {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1111
				_go_fuzz_dep_.CoverTab[185302]++
															break
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1112
				// _ = "end of CoverTab[185302]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1113
				_go_fuzz_dep_.CoverTab[185303]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1113
				// _ = "end of CoverTab[185303]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1113
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1113
			// _ = "end of CoverTab[185301]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1114
		// _ = "end of CoverTab[185298]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1114
		_go_fuzz_dep_.CoverTab[185299]++
													if advance == 1 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1115
			_go_fuzz_dep_.CoverTab[185304]++
														out = append(out, fi)
														continue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1117
			// _ = "end of CoverTab[185304]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1118
			_go_fuzz_dep_.CoverTab[185305]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1118
			// _ = "end of CoverTab[185305]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1118
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1118
		// _ = "end of CoverTab[185299]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1118
		_go_fuzz_dep_.CoverTab[185300]++
													dominant, ok := dominantField(fields[i : i+advance])
													if ok {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1120
			_go_fuzz_dep_.CoverTab[185306]++
														out = append(out, dominant)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1121
			// _ = "end of CoverTab[185306]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1122
			_go_fuzz_dep_.CoverTab[185307]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1122
			// _ = "end of CoverTab[185307]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1122
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1122
		// _ = "end of CoverTab[185300]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1123
	// _ = "end of CoverTab[185258]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1123
	_go_fuzz_dep_.CoverTab[185259]++

												fields = out
												sort.Sort(byIndex(fields))

												return fields
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1128
	// _ = "end of CoverTab[185259]"
}

// dominantField looks through the fields, all of which are known to
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1131
// have the same name, to find the single field that dominates the
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1131
// others using Go's embedding rules, modified by the presence of
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1131
// JSON tags. If there are multiple top-level fields, the boolean
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1131
// will be false: This condition is an error in Go and we skip all
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1131
// the fields.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1137
func dominantField(fields []field) (field, bool) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1137
	_go_fuzz_dep_.CoverTab[185308]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1141
	length := len(fields[0].index)
	tagged := -1
	for i, f := range fields {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1143
		_go_fuzz_dep_.CoverTab[185312]++
													if len(f.index) > length {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1144
			_go_fuzz_dep_.CoverTab[185314]++
														fields = fields[:i]
														break
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1146
			// _ = "end of CoverTab[185314]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1147
			_go_fuzz_dep_.CoverTab[185315]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1147
			// _ = "end of CoverTab[185315]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1147
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1147
		// _ = "end of CoverTab[185312]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1147
		_go_fuzz_dep_.CoverTab[185313]++
													if f.tag {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1148
			_go_fuzz_dep_.CoverTab[185316]++
														if tagged >= 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1149
				_go_fuzz_dep_.CoverTab[185318]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1152
				return field{}, false
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1152
				// _ = "end of CoverTab[185318]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1153
				_go_fuzz_dep_.CoverTab[185319]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1153
				// _ = "end of CoverTab[185319]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1153
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1153
			// _ = "end of CoverTab[185316]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1153
			_go_fuzz_dep_.CoverTab[185317]++
														tagged = i
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1154
			// _ = "end of CoverTab[185317]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1155
			_go_fuzz_dep_.CoverTab[185320]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1155
			// _ = "end of CoverTab[185320]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1155
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1155
		// _ = "end of CoverTab[185313]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1156
	// _ = "end of CoverTab[185308]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1156
	_go_fuzz_dep_.CoverTab[185309]++
												if tagged >= 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1157
		_go_fuzz_dep_.CoverTab[185321]++
													return fields[tagged], true
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1158
		// _ = "end of CoverTab[185321]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1159
		_go_fuzz_dep_.CoverTab[185322]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1159
		// _ = "end of CoverTab[185322]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1159
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1159
	// _ = "end of CoverTab[185309]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1159
	_go_fuzz_dep_.CoverTab[185310]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1163
	if len(fields) > 1 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1163
		_go_fuzz_dep_.CoverTab[185323]++
													return field{}, false
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1164
		// _ = "end of CoverTab[185323]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1165
		_go_fuzz_dep_.CoverTab[185324]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1165
		// _ = "end of CoverTab[185324]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1165
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1165
	// _ = "end of CoverTab[185310]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1165
	_go_fuzz_dep_.CoverTab[185311]++
												return fields[0], true
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1166
	// _ = "end of CoverTab[185311]"
}

var fieldCache struct {
	sync.RWMutex
	m	map[reflect.Type][]field
}

// cachedTypeFields is like typeFields but uses a cache to avoid repeated work.
func cachedTypeFields(t reflect.Type) []field {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1175
	_go_fuzz_dep_.CoverTab[185325]++
												fieldCache.RLock()
												f := fieldCache.m[t]
												fieldCache.RUnlock()
												if f != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1179
		_go_fuzz_dep_.CoverTab[185329]++
													return f
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1180
		// _ = "end of CoverTab[185329]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1181
		_go_fuzz_dep_.CoverTab[185330]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1181
		// _ = "end of CoverTab[185330]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1181
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1181
	// _ = "end of CoverTab[185325]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1181
	_go_fuzz_dep_.CoverTab[185326]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1185
	f = typeFields(t)
	if f == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1186
		_go_fuzz_dep_.CoverTab[185331]++
													f = []field{}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1187
		// _ = "end of CoverTab[185331]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1188
		_go_fuzz_dep_.CoverTab[185332]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1188
		// _ = "end of CoverTab[185332]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1188
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1188
	// _ = "end of CoverTab[185326]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1188
	_go_fuzz_dep_.CoverTab[185327]++

												fieldCache.Lock()
												if fieldCache.m == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1191
		_go_fuzz_dep_.CoverTab[185333]++
													fieldCache.m = map[reflect.Type][]field{}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1192
		// _ = "end of CoverTab[185333]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1193
		_go_fuzz_dep_.CoverTab[185334]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1193
		// _ = "end of CoverTab[185334]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1193
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1193
	// _ = "end of CoverTab[185327]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1193
	_go_fuzz_dep_.CoverTab[185328]++
												fieldCache.m[t] = f
												fieldCache.Unlock()
												return f
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1196
	// _ = "end of CoverTab[185328]"
}

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1197
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/encode.go:1197
var _ = _go_fuzz_dep_.CoverTab
