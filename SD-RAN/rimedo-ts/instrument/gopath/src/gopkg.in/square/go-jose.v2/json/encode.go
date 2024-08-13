// Copyright 2010 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:5
// Package json implements encoding and decoding of JSON objects as defined in
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:5
// RFC 4627. The mapping between JSON objects and Go values is described
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:5
// in the documentation for the Marshal and Unmarshal functions.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:5
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:5
// See "JSON and Go" for an introduction to this package:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:5
// https://golang.org/doc/articles/json_and_go.html
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:11
package json

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:11
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:11
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:11
)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:11
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:11
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:11
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
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
// Marshal traverses the value v recursively.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
// If an encountered value implements the Marshaler interface
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
// and is not a nil pointer, Marshal calls its MarshalJSON method
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
// to produce JSON. If no MarshalJSON method is present but the
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
// value implements encoding.TextMarshaler instead, Marshal calls
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
// its MarshalText method.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
// The nil pointer exception is not strictly necessary
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
// but mimics a similar, necessary exception in the behavior of
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
// UnmarshalJSON.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
// Otherwise, Marshal uses the following type-dependent default encodings:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
// Boolean values encode as JSON booleans.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
// Floating point, integer, and Number values encode as JSON numbers.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
// String values encode as JSON strings coerced to valid UTF-8,
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
// replacing invalid bytes with the Unicode replacement rune.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
// The angle brackets "<" and ">" are escaped to "\u003c" and "\u003e"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
// to keep some browsers from misinterpreting JSON output as HTML.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
// Ampersand "&" is also escaped to "\u0026" for the same reason.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
// Array and slice values encode as JSON arrays, except that
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
// []byte encodes as a base64-encoded string, and a nil slice
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
// encodes as the null JSON object.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
// Struct values encode as JSON objects. Each exported struct field
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
// becomes a member of the object unless
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
//   - the field's tag is "-", or
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
//   - the field is empty and its tag specifies the "omitempty" option.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
// The empty values are false, 0, any
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
// nil pointer or interface value, and any array, slice, map, or string of
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
// length zero. The object's default key string is the struct field name
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
// but can be specified in the struct field's tag value. The "json" key in
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
// the struct field's tag value is the key name, followed by an optional comma
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
// and options. Examples:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
//	// Field is ignored by this package.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
//	Field int `json:"-"`
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
//	// Field appears in JSON as key "myName".
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
//	Field int `json:"myName"`
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
//	// Field appears in JSON as key "myName" and
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
//	// the field is omitted from the object if its value is empty,
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
//	// as defined above.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
//	Field int `json:"myName,omitempty"`
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
//	// Field appears in JSON as key "Field" (the default), but
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
//	// the field is skipped if empty.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
//	// Note the leading comma.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
//	Field int `json:",omitempty"`
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
// The "string" option signals that a field is stored as JSON inside a
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
// JSON-encoded string. It applies only to fields of string, floating point,
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
// integer, or boolean types. This extra level of encoding is sometimes used
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
// when communicating with JavaScript programs:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
//	Int64String int64 `json:",string"`
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
// The key name will be used if it's a non-empty string consisting of
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
// only Unicode letters, digits, dollar signs, percent signs, hyphens,
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
// underscores and slashes.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
// Anonymous struct fields are usually marshaled as if their inner exported fields
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
// were fields in the outer struct, subject to the usual Go visibility rules amended
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
// as described in the next paragraph.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
// An anonymous struct field with a name given in its JSON tag is treated as
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
// having that name, rather than being anonymous.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
// An anonymous struct field of interface type is treated the same as having
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
// that type as its name, rather than being anonymous.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
// The Go visibility rules for struct fields are amended for JSON when
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
// deciding which field to marshal or unmarshal. If there are
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
// multiple fields at the same level, and that level is the least
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
// nested (and would therefore be the nesting level selected by the
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
// usual Go rules), the following extra rules apply:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
// 1) Of those fields, if any are JSON-tagged, only tagged fields are considered,
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
// even if there are multiple untagged fields that would otherwise conflict.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
// 2) If there is exactly one field (tagged or not according to the first rule), that is selected.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
// 3) Otherwise there are multiple fields, and all are ignored; no error occurs.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
// Handling of anonymous struct fields is new in Go 1.1.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
// Prior to Go 1.1, anonymous struct fields were ignored. To force ignoring of
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
// an anonymous struct field in both current and earlier versions, give the field
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
// a JSON tag of "-".
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
// Map values encode as JSON objects.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
// The map's key type must be string; the map keys are used as JSON object
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
// keys, subject to the UTF-8 coercion described for string values above.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
// Pointer values encode as the value pointed to.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
// A nil pointer encodes as the null JSON object.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
// Interface values encode as the value contained in the interface.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
// A nil interface value encodes as the null JSON object.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
// Channel, complex, and function values cannot be encoded in JSON.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
// Attempting to encode such a value causes Marshal to return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
// an UnsupportedTypeError.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
// JSON cannot represent cyclic data structures and Marshal does not
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
// handle them.  Passing cyclic structures to Marshal will result in
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:29
// an infinite recursion.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:137
func Marshal(v interface{}) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:137
	_go_fuzz_dep_.CoverTab[187931]++
											e := &encodeState{}
											err := e.marshal(v)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:140
		_go_fuzz_dep_.CoverTab[187933]++
												return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:141
		// _ = "end of CoverTab[187933]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:142
		_go_fuzz_dep_.CoverTab[187934]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:142
		// _ = "end of CoverTab[187934]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:142
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:142
	// _ = "end of CoverTab[187931]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:142
	_go_fuzz_dep_.CoverTab[187932]++
											return e.Bytes(), nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:143
	// _ = "end of CoverTab[187932]"
}

// MarshalIndent is like Marshal but applies Indent to format the output.
func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:147
	_go_fuzz_dep_.CoverTab[187935]++
											b, err := Marshal(v)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:149
		_go_fuzz_dep_.CoverTab[187938]++
												return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:150
		// _ = "end of CoverTab[187938]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:151
		_go_fuzz_dep_.CoverTab[187939]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:151
		// _ = "end of CoverTab[187939]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:151
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:151
	// _ = "end of CoverTab[187935]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:151
	_go_fuzz_dep_.CoverTab[187936]++
											var buf bytes.Buffer
											err = Indent(&buf, b, prefix, indent)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:154
		_go_fuzz_dep_.CoverTab[187940]++
												return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:155
		// _ = "end of CoverTab[187940]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:156
		_go_fuzz_dep_.CoverTab[187941]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:156
		// _ = "end of CoverTab[187941]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:156
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:156
	// _ = "end of CoverTab[187936]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:156
	_go_fuzz_dep_.CoverTab[187937]++
											return buf.Bytes(), nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:157
	// _ = "end of CoverTab[187937]"
}

// HTMLEscape appends to dst the JSON-encoded src with <, >, &, U+2028 and U+2029
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:160
// characters inside string literals changed to \u003c, \u003e, \u0026, \u2028, \u2029
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:160
// so that the JSON will be safe to embed inside HTML <script> tags.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:160
// For historical reasons, web browsers don't honor standard HTML
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:160
// escaping within <script> tags, so an alternative JSON encoding must
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:160
// be used.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:166
func HTMLEscape(dst *bytes.Buffer, src []byte) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:166
	_go_fuzz_dep_.CoverTab[187942]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:169
	start := 0
	for i, c := range src {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:170
		_go_fuzz_dep_.CoverTab[187944]++
												if c == '<' || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:171
			_go_fuzz_dep_.CoverTab[187946]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:171
			return c == '>'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:171
			// _ = "end of CoverTab[187946]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:171
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:171
			_go_fuzz_dep_.CoverTab[187947]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:171
			return c == '&'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:171
			// _ = "end of CoverTab[187947]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:171
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:171
			_go_fuzz_dep_.CoverTab[187948]++
													if start < i {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:172
				_go_fuzz_dep_.CoverTab[187950]++
														dst.Write(src[start:i])
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:173
				// _ = "end of CoverTab[187950]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:174
				_go_fuzz_dep_.CoverTab[187951]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:174
				// _ = "end of CoverTab[187951]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:174
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:174
			// _ = "end of CoverTab[187948]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:174
			_go_fuzz_dep_.CoverTab[187949]++
													dst.WriteString(`\u00`)
													dst.WriteByte(hex[c>>4])
													dst.WriteByte(hex[c&0xF])
													start = i + 1
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:178
			// _ = "end of CoverTab[187949]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:179
			_go_fuzz_dep_.CoverTab[187952]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:179
			// _ = "end of CoverTab[187952]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:179
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:179
		// _ = "end of CoverTab[187944]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:179
		_go_fuzz_dep_.CoverTab[187945]++

												if c == 0xE2 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:181
			_go_fuzz_dep_.CoverTab[187953]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:181
			return i+2 < len(src)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:181
			// _ = "end of CoverTab[187953]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:181
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:181
			_go_fuzz_dep_.CoverTab[187954]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:181
			return src[i+1] == 0x80
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:181
			// _ = "end of CoverTab[187954]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:181
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:181
			_go_fuzz_dep_.CoverTab[187955]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:181
			return src[i+2]&^1 == 0xA8
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:181
			// _ = "end of CoverTab[187955]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:181
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:181
			_go_fuzz_dep_.CoverTab[187956]++
													if start < i {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:182
				_go_fuzz_dep_.CoverTab[187958]++
														dst.Write(src[start:i])
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:183
				// _ = "end of CoverTab[187958]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:184
				_go_fuzz_dep_.CoverTab[187959]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:184
				// _ = "end of CoverTab[187959]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:184
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:184
			// _ = "end of CoverTab[187956]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:184
			_go_fuzz_dep_.CoverTab[187957]++
													dst.WriteString(`\u202`)
													dst.WriteByte(hex[src[i+2]&0xF])
													start = i + 3
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:187
			// _ = "end of CoverTab[187957]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:188
			_go_fuzz_dep_.CoverTab[187960]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:188
			// _ = "end of CoverTab[187960]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:188
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:188
		// _ = "end of CoverTab[187945]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:189
	// _ = "end of CoverTab[187942]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:189
	_go_fuzz_dep_.CoverTab[187943]++
											if start < len(src) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:190
		_go_fuzz_dep_.CoverTab[187961]++
												dst.Write(src[start:])
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:191
		// _ = "end of CoverTab[187961]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:192
		_go_fuzz_dep_.CoverTab[187962]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:192
		// _ = "end of CoverTab[187962]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:192
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:192
	// _ = "end of CoverTab[187943]"
}

// Marshaler is the interface implemented by objects that
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:195
// can marshal themselves into valid JSON.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:197
type Marshaler interface {
	MarshalJSON() ([]byte, error)
}

// An UnsupportedTypeError is returned by Marshal when attempting
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:201
// to encode an unsupported value type.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:203
type UnsupportedTypeError struct {
	Type reflect.Type
}

func (e *UnsupportedTypeError) Error() string {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:207
	_go_fuzz_dep_.CoverTab[187963]++
											return "json: unsupported type: " + e.Type.String()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:208
	// _ = "end of CoverTab[187963]"
}

type UnsupportedValueError struct {
	Value	reflect.Value
	Str	string
}

func (e *UnsupportedValueError) Error() string {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:216
	_go_fuzz_dep_.CoverTab[187964]++
											return "json: unsupported value: " + e.Str
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:217
	// _ = "end of CoverTab[187964]"
}

// Before Go 1.2, an InvalidUTF8Error was returned by Marshal when
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:220
// attempting to encode a string value with invalid UTF-8 sequences.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:220
// As of Go 1.2, Marshal instead coerces the string to valid UTF-8 by
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:220
// replacing invalid bytes with the Unicode replacement rune U+FFFD.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:220
// This error is no longer generated but is kept for backwards compatibility
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:220
// with programs that might mention it.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:226
type InvalidUTF8Error struct {
	S string	// the whole string value that caused the error
}

func (e *InvalidUTF8Error) Error() string {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:230
	_go_fuzz_dep_.CoverTab[187965]++
											return "json: invalid UTF-8 in string: " + strconv.Quote(e.S)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:231
	// _ = "end of CoverTab[187965]"
}

type MarshalerError struct {
	Type	reflect.Type
	Err	error
}

func (e *MarshalerError) Error() string {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:239
	_go_fuzz_dep_.CoverTab[187966]++
											return "json: error calling MarshalJSON for type " + e.Type.String() + ": " + e.Err.Error()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:240
	// _ = "end of CoverTab[187966]"
}

var hex = "0123456789abcdef"

// An encodeState encodes JSON into a bytes.Buffer.
type encodeState struct {
	bytes.Buffer	// accumulated output
	scratch		[64]byte
}

var encodeStatePool sync.Pool

func newEncodeState() *encodeState {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:253
	_go_fuzz_dep_.CoverTab[187967]++
											if v := encodeStatePool.Get(); v != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:254
		_go_fuzz_dep_.CoverTab[187969]++
												e := v.(*encodeState)
												e.Reset()
												return e
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:257
		// _ = "end of CoverTab[187969]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:258
		_go_fuzz_dep_.CoverTab[187970]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:258
		// _ = "end of CoverTab[187970]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:258
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:258
	// _ = "end of CoverTab[187967]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:258
	_go_fuzz_dep_.CoverTab[187968]++
											return new(encodeState)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:259
	// _ = "end of CoverTab[187968]"
}

func (e *encodeState) marshal(v interface{}) (err error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:262
	_go_fuzz_dep_.CoverTab[187971]++
											defer func() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:263
		_go_fuzz_dep_.CoverTab[187973]++
												if r := recover(); r != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:264
			_go_fuzz_dep_.CoverTab[187974]++
													if _, ok := r.(runtime.Error); ok {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:265
				_go_fuzz_dep_.CoverTab[187977]++
														panic(r)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:266
				// _ = "end of CoverTab[187977]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:267
				_go_fuzz_dep_.CoverTab[187978]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:267
				// _ = "end of CoverTab[187978]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:267
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:267
			// _ = "end of CoverTab[187974]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:267
			_go_fuzz_dep_.CoverTab[187975]++
													if s, ok := r.(string); ok {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:268
				_go_fuzz_dep_.CoverTab[187979]++
														panic(s)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:269
				// _ = "end of CoverTab[187979]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:270
				_go_fuzz_dep_.CoverTab[187980]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:270
				// _ = "end of CoverTab[187980]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:270
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:270
			// _ = "end of CoverTab[187975]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:270
			_go_fuzz_dep_.CoverTab[187976]++
													err = r.(error)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:271
			// _ = "end of CoverTab[187976]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:272
			_go_fuzz_dep_.CoverTab[187981]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:272
			// _ = "end of CoverTab[187981]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:272
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:272
		// _ = "end of CoverTab[187973]"
	}()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:273
	// _ = "end of CoverTab[187971]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:273
	_go_fuzz_dep_.CoverTab[187972]++
											e.reflectValue(reflect.ValueOf(v))
											return nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:275
	// _ = "end of CoverTab[187972]"
}

func (e *encodeState) error(err error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:278
	_go_fuzz_dep_.CoverTab[187982]++
											panic(err)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:279
	// _ = "end of CoverTab[187982]"
}

func isEmptyValue(v reflect.Value) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:282
	_go_fuzz_dep_.CoverTab[187983]++
											switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:284
		_go_fuzz_dep_.CoverTab[187985]++
												return v.Len() == 0
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:285
		// _ = "end of CoverTab[187985]"
	case reflect.Bool:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:286
		_go_fuzz_dep_.CoverTab[187986]++
												return !v.Bool()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:287
		// _ = "end of CoverTab[187986]"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:288
		_go_fuzz_dep_.CoverTab[187987]++
												return v.Int() == 0
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:289
		// _ = "end of CoverTab[187987]"
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:290
		_go_fuzz_dep_.CoverTab[187988]++
												return v.Uint() == 0
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:291
		// _ = "end of CoverTab[187988]"
	case reflect.Float32, reflect.Float64:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:292
		_go_fuzz_dep_.CoverTab[187989]++
												return v.Float() == 0
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:293
		// _ = "end of CoverTab[187989]"
	case reflect.Interface, reflect.Ptr:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:294
		_go_fuzz_dep_.CoverTab[187990]++
												return v.IsNil()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:295
		// _ = "end of CoverTab[187990]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:295
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:295
		_go_fuzz_dep_.CoverTab[187991]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:295
		// _ = "end of CoverTab[187991]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:296
	// _ = "end of CoverTab[187983]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:296
	_go_fuzz_dep_.CoverTab[187984]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:297
	// _ = "end of CoverTab[187984]"
}

func (e *encodeState) reflectValue(v reflect.Value) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:300
	_go_fuzz_dep_.CoverTab[187992]++
											valueEncoder(v)(e, v, false)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:301
	// _ = "end of CoverTab[187992]"
}

type encoderFunc func(e *encodeState, v reflect.Value, quoted bool)

var encoderCache struct {
	sync.RWMutex
	m	map[reflect.Type]encoderFunc
}

func valueEncoder(v reflect.Value) encoderFunc {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:311
	_go_fuzz_dep_.CoverTab[187993]++
											if !v.IsValid() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:312
		_go_fuzz_dep_.CoverTab[187995]++
												return invalidValueEncoder
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:313
		// _ = "end of CoverTab[187995]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:314
		_go_fuzz_dep_.CoverTab[187996]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:314
		// _ = "end of CoverTab[187996]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:314
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:314
	// _ = "end of CoverTab[187993]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:314
	_go_fuzz_dep_.CoverTab[187994]++
											return typeEncoder(v.Type())
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:315
	// _ = "end of CoverTab[187994]"
}

func typeEncoder(t reflect.Type) encoderFunc {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:318
	_go_fuzz_dep_.CoverTab[187997]++
											encoderCache.RLock()
											f := encoderCache.m[t]
											encoderCache.RUnlock()
											if f != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:322
		_go_fuzz_dep_.CoverTab[188001]++
												return f
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:323
		// _ = "end of CoverTab[188001]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:324
		_go_fuzz_dep_.CoverTab[188002]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:324
		// _ = "end of CoverTab[188002]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:324
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:324
	// _ = "end of CoverTab[187997]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:324
	_go_fuzz_dep_.CoverTab[187998]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:330
	encoderCache.Lock()
	if encoderCache.m == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:331
		_go_fuzz_dep_.CoverTab[188003]++
												encoderCache.m = make(map[reflect.Type]encoderFunc)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:332
		// _ = "end of CoverTab[188003]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:333
		_go_fuzz_dep_.CoverTab[188004]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:333
		// _ = "end of CoverTab[188004]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:333
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:333
	// _ = "end of CoverTab[187998]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:333
	_go_fuzz_dep_.CoverTab[187999]++
											var wg sync.WaitGroup
											wg.Add(1)
											encoderCache.m[t] = func(e *encodeState, v reflect.Value, quoted bool) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:336
		_go_fuzz_dep_.CoverTab[188005]++
												wg.Wait()
												f(e, v, quoted)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:338
		// _ = "end of CoverTab[188005]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:339
	// _ = "end of CoverTab[187999]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:339
	_go_fuzz_dep_.CoverTab[188000]++
											encoderCache.Unlock()

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:344
	f = newTypeEncoder(t, true)
											wg.Done()
											encoderCache.Lock()
											encoderCache.m[t] = f
											encoderCache.Unlock()
											return f
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:349
	// _ = "end of CoverTab[188000]"
}

var (
	marshalerType		= reflect.TypeOf(new(Marshaler)).Elem()
	textMarshalerType	= reflect.TypeOf(new(encoding.TextMarshaler)).Elem()
)

// newTypeEncoder constructs an encoderFunc for a type.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:357
// The returned encoder only checks CanAddr when allowAddr is true.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:359
func newTypeEncoder(t reflect.Type, allowAddr bool) encoderFunc {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:359
	_go_fuzz_dep_.CoverTab[188006]++
											if t.Implements(marshalerType) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:360
		_go_fuzz_dep_.CoverTab[188011]++
												return marshalerEncoder
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:361
		// _ = "end of CoverTab[188011]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:362
		_go_fuzz_dep_.CoverTab[188012]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:362
		// _ = "end of CoverTab[188012]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:362
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:362
	// _ = "end of CoverTab[188006]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:362
	_go_fuzz_dep_.CoverTab[188007]++
											if t.Kind() != reflect.Ptr && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:363
		_go_fuzz_dep_.CoverTab[188013]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:363
		return allowAddr
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:363
		// _ = "end of CoverTab[188013]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:363
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:363
		_go_fuzz_dep_.CoverTab[188014]++
												if reflect.PtrTo(t).Implements(marshalerType) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:364
			_go_fuzz_dep_.CoverTab[188015]++
													return newCondAddrEncoder(addrMarshalerEncoder, newTypeEncoder(t, false))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:365
			// _ = "end of CoverTab[188015]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:366
			_go_fuzz_dep_.CoverTab[188016]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:366
			// _ = "end of CoverTab[188016]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:366
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:366
		// _ = "end of CoverTab[188014]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:367
		_go_fuzz_dep_.CoverTab[188017]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:367
		// _ = "end of CoverTab[188017]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:367
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:367
	// _ = "end of CoverTab[188007]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:367
	_go_fuzz_dep_.CoverTab[188008]++

											if t.Implements(textMarshalerType) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:369
		_go_fuzz_dep_.CoverTab[188018]++
												return textMarshalerEncoder
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:370
		// _ = "end of CoverTab[188018]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:371
		_go_fuzz_dep_.CoverTab[188019]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:371
		// _ = "end of CoverTab[188019]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:371
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:371
	// _ = "end of CoverTab[188008]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:371
	_go_fuzz_dep_.CoverTab[188009]++
											if t.Kind() != reflect.Ptr && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:372
		_go_fuzz_dep_.CoverTab[188020]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:372
		return allowAddr
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:372
		// _ = "end of CoverTab[188020]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:372
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:372
		_go_fuzz_dep_.CoverTab[188021]++
												if reflect.PtrTo(t).Implements(textMarshalerType) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:373
			_go_fuzz_dep_.CoverTab[188022]++
													return newCondAddrEncoder(addrTextMarshalerEncoder, newTypeEncoder(t, false))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:374
			// _ = "end of CoverTab[188022]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:375
			_go_fuzz_dep_.CoverTab[188023]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:375
			// _ = "end of CoverTab[188023]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:375
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:375
		// _ = "end of CoverTab[188021]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:376
		_go_fuzz_dep_.CoverTab[188024]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:376
		// _ = "end of CoverTab[188024]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:376
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:376
	// _ = "end of CoverTab[188009]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:376
	_go_fuzz_dep_.CoverTab[188010]++

											switch t.Kind() {
	case reflect.Bool:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:379
		_go_fuzz_dep_.CoverTab[188025]++
												return boolEncoder
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:380
		// _ = "end of CoverTab[188025]"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:381
		_go_fuzz_dep_.CoverTab[188026]++
												return intEncoder
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:382
		// _ = "end of CoverTab[188026]"
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:383
		_go_fuzz_dep_.CoverTab[188027]++
												return uintEncoder
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:384
		// _ = "end of CoverTab[188027]"
	case reflect.Float32:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:385
		_go_fuzz_dep_.CoverTab[188028]++
												return float32Encoder
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:386
		// _ = "end of CoverTab[188028]"
	case reflect.Float64:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:387
		_go_fuzz_dep_.CoverTab[188029]++
												return float64Encoder
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:388
		// _ = "end of CoverTab[188029]"
	case reflect.String:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:389
		_go_fuzz_dep_.CoverTab[188030]++
												return stringEncoder
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:390
		// _ = "end of CoverTab[188030]"
	case reflect.Interface:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:391
		_go_fuzz_dep_.CoverTab[188031]++
												return interfaceEncoder
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:392
		// _ = "end of CoverTab[188031]"
	case reflect.Struct:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:393
		_go_fuzz_dep_.CoverTab[188032]++
												return newStructEncoder(t)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:394
		// _ = "end of CoverTab[188032]"
	case reflect.Map:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:395
		_go_fuzz_dep_.CoverTab[188033]++
												return newMapEncoder(t)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:396
		// _ = "end of CoverTab[188033]"
	case reflect.Slice:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:397
		_go_fuzz_dep_.CoverTab[188034]++
												return newSliceEncoder(t)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:398
		// _ = "end of CoverTab[188034]"
	case reflect.Array:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:399
		_go_fuzz_dep_.CoverTab[188035]++
												return newArrayEncoder(t)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:400
		// _ = "end of CoverTab[188035]"
	case reflect.Ptr:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:401
		_go_fuzz_dep_.CoverTab[188036]++
												return newPtrEncoder(t)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:402
		// _ = "end of CoverTab[188036]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:403
		_go_fuzz_dep_.CoverTab[188037]++
												return unsupportedTypeEncoder
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:404
		// _ = "end of CoverTab[188037]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:405
	// _ = "end of CoverTab[188010]"
}

func invalidValueEncoder(e *encodeState, v reflect.Value, quoted bool) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:408
	_go_fuzz_dep_.CoverTab[188038]++
											e.WriteString("null")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:409
	// _ = "end of CoverTab[188038]"
}

func marshalerEncoder(e *encodeState, v reflect.Value, quoted bool) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:412
	_go_fuzz_dep_.CoverTab[188039]++
											if v.Kind() == reflect.Ptr && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:413
		_go_fuzz_dep_.CoverTab[188042]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:413
		return v.IsNil()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:413
		// _ = "end of CoverTab[188042]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:413
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:413
		_go_fuzz_dep_.CoverTab[188043]++
												e.WriteString("null")
												return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:415
		// _ = "end of CoverTab[188043]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:416
		_go_fuzz_dep_.CoverTab[188044]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:416
		// _ = "end of CoverTab[188044]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:416
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:416
	// _ = "end of CoverTab[188039]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:416
	_go_fuzz_dep_.CoverTab[188040]++
											m := v.Interface().(Marshaler)
											b, err := m.MarshalJSON()
											if err == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:419
		_go_fuzz_dep_.CoverTab[188045]++

												err = compact(&e.Buffer, b, true)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:421
		// _ = "end of CoverTab[188045]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:422
		_go_fuzz_dep_.CoverTab[188046]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:422
		// _ = "end of CoverTab[188046]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:422
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:422
	// _ = "end of CoverTab[188040]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:422
	_go_fuzz_dep_.CoverTab[188041]++
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:423
		_go_fuzz_dep_.CoverTab[188047]++
												e.error(&MarshalerError{v.Type(), err})
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:424
		// _ = "end of CoverTab[188047]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:425
		_go_fuzz_dep_.CoverTab[188048]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:425
		// _ = "end of CoverTab[188048]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:425
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:425
	// _ = "end of CoverTab[188041]"
}

func addrMarshalerEncoder(e *encodeState, v reflect.Value, quoted bool) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:428
	_go_fuzz_dep_.CoverTab[188049]++
											va := v.Addr()
											if va.IsNil() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:430
		_go_fuzz_dep_.CoverTab[188052]++
												e.WriteString("null")
												return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:432
		// _ = "end of CoverTab[188052]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:433
		_go_fuzz_dep_.CoverTab[188053]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:433
		// _ = "end of CoverTab[188053]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:433
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:433
	// _ = "end of CoverTab[188049]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:433
	_go_fuzz_dep_.CoverTab[188050]++
											m := va.Interface().(Marshaler)
											b, err := m.MarshalJSON()
											if err == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:436
		_go_fuzz_dep_.CoverTab[188054]++

												err = compact(&e.Buffer, b, true)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:438
		// _ = "end of CoverTab[188054]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:439
		_go_fuzz_dep_.CoverTab[188055]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:439
		// _ = "end of CoverTab[188055]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:439
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:439
	// _ = "end of CoverTab[188050]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:439
	_go_fuzz_dep_.CoverTab[188051]++
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:440
		_go_fuzz_dep_.CoverTab[188056]++
												e.error(&MarshalerError{v.Type(), err})
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:441
		// _ = "end of CoverTab[188056]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:442
		_go_fuzz_dep_.CoverTab[188057]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:442
		// _ = "end of CoverTab[188057]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:442
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:442
	// _ = "end of CoverTab[188051]"
}

func textMarshalerEncoder(e *encodeState, v reflect.Value, quoted bool) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:445
	_go_fuzz_dep_.CoverTab[188058]++
											if v.Kind() == reflect.Ptr && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:446
		_go_fuzz_dep_.CoverTab[188061]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:446
		return v.IsNil()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:446
		// _ = "end of CoverTab[188061]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:446
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:446
		_go_fuzz_dep_.CoverTab[188062]++
												e.WriteString("null")
												return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:448
		// _ = "end of CoverTab[188062]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:449
		_go_fuzz_dep_.CoverTab[188063]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:449
		// _ = "end of CoverTab[188063]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:449
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:449
	// _ = "end of CoverTab[188058]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:449
	_go_fuzz_dep_.CoverTab[188059]++
											m := v.Interface().(encoding.TextMarshaler)
											b, err := m.MarshalText()
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:452
		_go_fuzz_dep_.CoverTab[188064]++
												e.error(&MarshalerError{v.Type(), err})
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:453
		// _ = "end of CoverTab[188064]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:454
		_go_fuzz_dep_.CoverTab[188065]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:454
		// _ = "end of CoverTab[188065]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:454
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:454
	// _ = "end of CoverTab[188059]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:454
	_go_fuzz_dep_.CoverTab[188060]++
											e.stringBytes(b)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:455
	// _ = "end of CoverTab[188060]"
}

func addrTextMarshalerEncoder(e *encodeState, v reflect.Value, quoted bool) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:458
	_go_fuzz_dep_.CoverTab[188066]++
											va := v.Addr()
											if va.IsNil() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:460
		_go_fuzz_dep_.CoverTab[188069]++
												e.WriteString("null")
												return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:462
		// _ = "end of CoverTab[188069]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:463
		_go_fuzz_dep_.CoverTab[188070]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:463
		// _ = "end of CoverTab[188070]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:463
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:463
	// _ = "end of CoverTab[188066]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:463
	_go_fuzz_dep_.CoverTab[188067]++
											m := va.Interface().(encoding.TextMarshaler)
											b, err := m.MarshalText()
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:466
		_go_fuzz_dep_.CoverTab[188071]++
												e.error(&MarshalerError{v.Type(), err})
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:467
		// _ = "end of CoverTab[188071]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:468
		_go_fuzz_dep_.CoverTab[188072]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:468
		// _ = "end of CoverTab[188072]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:468
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:468
	// _ = "end of CoverTab[188067]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:468
	_go_fuzz_dep_.CoverTab[188068]++
											e.stringBytes(b)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:469
	// _ = "end of CoverTab[188068]"
}

func boolEncoder(e *encodeState, v reflect.Value, quoted bool) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:472
	_go_fuzz_dep_.CoverTab[188073]++
											if quoted {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:473
		_go_fuzz_dep_.CoverTab[188076]++
												e.WriteByte('"')
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:474
		// _ = "end of CoverTab[188076]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:475
		_go_fuzz_dep_.CoverTab[188077]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:475
		// _ = "end of CoverTab[188077]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:475
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:475
	// _ = "end of CoverTab[188073]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:475
	_go_fuzz_dep_.CoverTab[188074]++
											if v.Bool() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:476
		_go_fuzz_dep_.CoverTab[188078]++
												e.WriteString("true")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:477
		// _ = "end of CoverTab[188078]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:478
		_go_fuzz_dep_.CoverTab[188079]++
												e.WriteString("false")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:479
		// _ = "end of CoverTab[188079]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:480
	// _ = "end of CoverTab[188074]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:480
	_go_fuzz_dep_.CoverTab[188075]++
											if quoted {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:481
		_go_fuzz_dep_.CoverTab[188080]++
												e.WriteByte('"')
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:482
		// _ = "end of CoverTab[188080]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:483
		_go_fuzz_dep_.CoverTab[188081]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:483
		// _ = "end of CoverTab[188081]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:483
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:483
	// _ = "end of CoverTab[188075]"
}

func intEncoder(e *encodeState, v reflect.Value, quoted bool) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:486
	_go_fuzz_dep_.CoverTab[188082]++
											b := strconv.AppendInt(e.scratch[:0], v.Int(), 10)
											if quoted {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:488
		_go_fuzz_dep_.CoverTab[188084]++
												e.WriteByte('"')
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:489
		// _ = "end of CoverTab[188084]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:490
		_go_fuzz_dep_.CoverTab[188085]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:490
		// _ = "end of CoverTab[188085]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:490
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:490
	// _ = "end of CoverTab[188082]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:490
	_go_fuzz_dep_.CoverTab[188083]++
											e.Write(b)
											if quoted {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:492
		_go_fuzz_dep_.CoverTab[188086]++
												e.WriteByte('"')
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:493
		// _ = "end of CoverTab[188086]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:494
		_go_fuzz_dep_.CoverTab[188087]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:494
		// _ = "end of CoverTab[188087]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:494
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:494
	// _ = "end of CoverTab[188083]"
}

func uintEncoder(e *encodeState, v reflect.Value, quoted bool) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:497
	_go_fuzz_dep_.CoverTab[188088]++
											b := strconv.AppendUint(e.scratch[:0], v.Uint(), 10)
											if quoted {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:499
		_go_fuzz_dep_.CoverTab[188090]++
												e.WriteByte('"')
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:500
		// _ = "end of CoverTab[188090]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:501
		_go_fuzz_dep_.CoverTab[188091]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:501
		// _ = "end of CoverTab[188091]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:501
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:501
	// _ = "end of CoverTab[188088]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:501
	_go_fuzz_dep_.CoverTab[188089]++
											e.Write(b)
											if quoted {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:503
		_go_fuzz_dep_.CoverTab[188092]++
												e.WriteByte('"')
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:504
		// _ = "end of CoverTab[188092]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:505
		_go_fuzz_dep_.CoverTab[188093]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:505
		// _ = "end of CoverTab[188093]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:505
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:505
	// _ = "end of CoverTab[188089]"
}

type floatEncoder int	// number of bits

func (bits floatEncoder) encode(e *encodeState, v reflect.Value, quoted bool) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:510
	_go_fuzz_dep_.CoverTab[188094]++
											f := v.Float()
											if math.IsInf(f, 0) || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:512
		_go_fuzz_dep_.CoverTab[188097]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:512
		return math.IsNaN(f)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:512
		// _ = "end of CoverTab[188097]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:512
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:512
		_go_fuzz_dep_.CoverTab[188098]++
												e.error(&UnsupportedValueError{v, strconv.FormatFloat(f, 'g', -1, int(bits))})
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:513
		// _ = "end of CoverTab[188098]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:514
		_go_fuzz_dep_.CoverTab[188099]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:514
		// _ = "end of CoverTab[188099]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:514
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:514
	// _ = "end of CoverTab[188094]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:514
	_go_fuzz_dep_.CoverTab[188095]++
											b := strconv.AppendFloat(e.scratch[:0], f, 'g', -1, int(bits))
											if quoted {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:516
		_go_fuzz_dep_.CoverTab[188100]++
												e.WriteByte('"')
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:517
		// _ = "end of CoverTab[188100]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:518
		_go_fuzz_dep_.CoverTab[188101]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:518
		// _ = "end of CoverTab[188101]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:518
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:518
	// _ = "end of CoverTab[188095]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:518
	_go_fuzz_dep_.CoverTab[188096]++
											e.Write(b)
											if quoted {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:520
		_go_fuzz_dep_.CoverTab[188102]++
												e.WriteByte('"')
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:521
		// _ = "end of CoverTab[188102]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:522
		_go_fuzz_dep_.CoverTab[188103]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:522
		// _ = "end of CoverTab[188103]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:522
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:522
	// _ = "end of CoverTab[188096]"
}

var (
	float32Encoder	= (floatEncoder(32)).encode
	float64Encoder	= (floatEncoder(64)).encode
)

func stringEncoder(e *encodeState, v reflect.Value, quoted bool) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:530
	_go_fuzz_dep_.CoverTab[188104]++
											if v.Type() == numberType {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:531
		_go_fuzz_dep_.CoverTab[188106]++
												numStr := v.String()

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:535
		if numStr == "" {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:535
			_go_fuzz_dep_.CoverTab[188109]++
													numStr = "0"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:536
			// _ = "end of CoverTab[188109]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:537
			_go_fuzz_dep_.CoverTab[188110]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:537
			// _ = "end of CoverTab[188110]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:537
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:537
		// _ = "end of CoverTab[188106]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:537
		_go_fuzz_dep_.CoverTab[188107]++
												if !isValidNumber(numStr) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:538
			_go_fuzz_dep_.CoverTab[188111]++
													e.error(fmt.Errorf("json: invalid number literal %q", numStr))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:539
			// _ = "end of CoverTab[188111]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:540
			_go_fuzz_dep_.CoverTab[188112]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:540
			// _ = "end of CoverTab[188112]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:540
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:540
		// _ = "end of CoverTab[188107]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:540
		_go_fuzz_dep_.CoverTab[188108]++
												e.WriteString(numStr)
												return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:542
		// _ = "end of CoverTab[188108]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:543
		_go_fuzz_dep_.CoverTab[188113]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:543
		// _ = "end of CoverTab[188113]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:543
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:543
	// _ = "end of CoverTab[188104]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:543
	_go_fuzz_dep_.CoverTab[188105]++
											if quoted {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:544
		_go_fuzz_dep_.CoverTab[188114]++
												sb, err := Marshal(v.String())
												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:546
			_go_fuzz_dep_.CoverTab[188116]++
													e.error(err)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:547
			// _ = "end of CoverTab[188116]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:548
			_go_fuzz_dep_.CoverTab[188117]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:548
			// _ = "end of CoverTab[188117]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:548
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:548
		// _ = "end of CoverTab[188114]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:548
		_go_fuzz_dep_.CoverTab[188115]++
												e.string(string(sb))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:549
		// _ = "end of CoverTab[188115]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:550
		_go_fuzz_dep_.CoverTab[188118]++
												e.string(v.String())
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:551
		// _ = "end of CoverTab[188118]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:552
	// _ = "end of CoverTab[188105]"
}

func interfaceEncoder(e *encodeState, v reflect.Value, quoted bool) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:555
	_go_fuzz_dep_.CoverTab[188119]++
											if v.IsNil() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:556
		_go_fuzz_dep_.CoverTab[188121]++
												e.WriteString("null")
												return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:558
		// _ = "end of CoverTab[188121]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:559
		_go_fuzz_dep_.CoverTab[188122]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:559
		// _ = "end of CoverTab[188122]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:559
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:559
	// _ = "end of CoverTab[188119]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:559
	_go_fuzz_dep_.CoverTab[188120]++
											e.reflectValue(v.Elem())
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:560
	// _ = "end of CoverTab[188120]"
}

func unsupportedTypeEncoder(e *encodeState, v reflect.Value, quoted bool) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:563
	_go_fuzz_dep_.CoverTab[188123]++
											e.error(&UnsupportedTypeError{v.Type()})
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:564
	// _ = "end of CoverTab[188123]"
}

type structEncoder struct {
	fields		[]field
	fieldEncs	[]encoderFunc
}

func (se *structEncoder) encode(e *encodeState, v reflect.Value, quoted bool) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:572
	_go_fuzz_dep_.CoverTab[188124]++
											e.WriteByte('{')
											first := true
											for i, f := range se.fields {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:575
		_go_fuzz_dep_.CoverTab[188126]++
												fv := fieldByIndex(v, f.index)
												if !fv.IsValid() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:577
			_go_fuzz_dep_.CoverTab[188129]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:577
			return f.omitEmpty && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:577
				_go_fuzz_dep_.CoverTab[188130]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:577
				return isEmptyValue(fv)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:577
				// _ = "end of CoverTab[188130]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:577
			}()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:577
			// _ = "end of CoverTab[188129]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:577
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:577
			_go_fuzz_dep_.CoverTab[188131]++
													continue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:578
			// _ = "end of CoverTab[188131]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:579
			_go_fuzz_dep_.CoverTab[188132]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:579
			// _ = "end of CoverTab[188132]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:579
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:579
		// _ = "end of CoverTab[188126]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:579
		_go_fuzz_dep_.CoverTab[188127]++
												if first {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:580
			_go_fuzz_dep_.CoverTab[188133]++
													first = false
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:581
			// _ = "end of CoverTab[188133]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:582
			_go_fuzz_dep_.CoverTab[188134]++
													e.WriteByte(',')
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:583
			// _ = "end of CoverTab[188134]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:584
		// _ = "end of CoverTab[188127]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:584
		_go_fuzz_dep_.CoverTab[188128]++
												e.string(f.name)
												e.WriteByte(':')
												se.fieldEncs[i](e, fv, f.quoted)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:587
		// _ = "end of CoverTab[188128]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:588
	// _ = "end of CoverTab[188124]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:588
	_go_fuzz_dep_.CoverTab[188125]++
											e.WriteByte('}')
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:589
	// _ = "end of CoverTab[188125]"
}

func newStructEncoder(t reflect.Type) encoderFunc {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:592
	_go_fuzz_dep_.CoverTab[188135]++
											fields := cachedTypeFields(t)
											se := &structEncoder{
		fields:		fields,
		fieldEncs:	make([]encoderFunc, len(fields)),
	}
	for i, f := range fields {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:598
		_go_fuzz_dep_.CoverTab[188137]++
												se.fieldEncs[i] = typeEncoder(typeByIndex(t, f.index))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:599
		// _ = "end of CoverTab[188137]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:600
	// _ = "end of CoverTab[188135]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:600
	_go_fuzz_dep_.CoverTab[188136]++
											return se.encode
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:601
	// _ = "end of CoverTab[188136]"
}

type mapEncoder struct {
	elemEnc encoderFunc
}

func (me *mapEncoder) encode(e *encodeState, v reflect.Value, _ bool) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:608
	_go_fuzz_dep_.CoverTab[188138]++
											if v.IsNil() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:609
		_go_fuzz_dep_.CoverTab[188141]++
												e.WriteString("null")
												return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:611
		// _ = "end of CoverTab[188141]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:612
		_go_fuzz_dep_.CoverTab[188142]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:612
		// _ = "end of CoverTab[188142]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:612
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:612
	// _ = "end of CoverTab[188138]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:612
	_go_fuzz_dep_.CoverTab[188139]++
											e.WriteByte('{')
											var sv stringValues = v.MapKeys()
											sort.Sort(sv)
											for i, k := range sv {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:616
		_go_fuzz_dep_.CoverTab[188143]++
												if i > 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:617
			_go_fuzz_dep_.CoverTab[188145]++
													e.WriteByte(',')
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:618
			// _ = "end of CoverTab[188145]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:619
			_go_fuzz_dep_.CoverTab[188146]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:619
			// _ = "end of CoverTab[188146]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:619
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:619
		// _ = "end of CoverTab[188143]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:619
		_go_fuzz_dep_.CoverTab[188144]++
												e.string(k.String())
												e.WriteByte(':')
												me.elemEnc(e, v.MapIndex(k), false)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:622
		// _ = "end of CoverTab[188144]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:623
	// _ = "end of CoverTab[188139]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:623
	_go_fuzz_dep_.CoverTab[188140]++
											e.WriteByte('}')
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:624
	// _ = "end of CoverTab[188140]"
}

func newMapEncoder(t reflect.Type) encoderFunc {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:627
	_go_fuzz_dep_.CoverTab[188147]++
											if t.Key().Kind() != reflect.String {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:628
		_go_fuzz_dep_.CoverTab[188149]++
												return unsupportedTypeEncoder
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:629
		// _ = "end of CoverTab[188149]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:630
		_go_fuzz_dep_.CoverTab[188150]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:630
		// _ = "end of CoverTab[188150]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:630
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:630
	// _ = "end of CoverTab[188147]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:630
	_go_fuzz_dep_.CoverTab[188148]++
											me := &mapEncoder{typeEncoder(t.Elem())}
											return me.encode
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:632
	// _ = "end of CoverTab[188148]"
}

func encodeByteSlice(e *encodeState, v reflect.Value, _ bool) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:635
	_go_fuzz_dep_.CoverTab[188151]++
											if v.IsNil() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:636
		_go_fuzz_dep_.CoverTab[188154]++
												e.WriteString("null")
												return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:638
		// _ = "end of CoverTab[188154]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:639
		_go_fuzz_dep_.CoverTab[188155]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:639
		// _ = "end of CoverTab[188155]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:639
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:639
	// _ = "end of CoverTab[188151]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:639
	_go_fuzz_dep_.CoverTab[188152]++
											s := v.Bytes()
											e.WriteByte('"')
											if len(s) < 1024 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:642
		_go_fuzz_dep_.CoverTab[188156]++

												dst := make([]byte, base64.StdEncoding.EncodedLen(len(s)))
												base64.StdEncoding.Encode(dst, s)
												e.Write(dst)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:646
		// _ = "end of CoverTab[188156]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:647
		_go_fuzz_dep_.CoverTab[188157]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:650
		enc := base64.NewEncoder(base64.StdEncoding, e)
												enc.Write(s)
												enc.Close()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:652
		// _ = "end of CoverTab[188157]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:653
	// _ = "end of CoverTab[188152]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:653
	_go_fuzz_dep_.CoverTab[188153]++
											e.WriteByte('"')
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:654
	// _ = "end of CoverTab[188153]"
}

// sliceEncoder just wraps an arrayEncoder, checking to make sure the value isn't nil.
type sliceEncoder struct {
	arrayEnc encoderFunc
}

func (se *sliceEncoder) encode(e *encodeState, v reflect.Value, _ bool) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:662
	_go_fuzz_dep_.CoverTab[188158]++
											if v.IsNil() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:663
		_go_fuzz_dep_.CoverTab[188160]++
												e.WriteString("null")
												return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:665
		// _ = "end of CoverTab[188160]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:666
		_go_fuzz_dep_.CoverTab[188161]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:666
		// _ = "end of CoverTab[188161]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:666
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:666
	// _ = "end of CoverTab[188158]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:666
	_go_fuzz_dep_.CoverTab[188159]++
											se.arrayEnc(e, v, false)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:667
	// _ = "end of CoverTab[188159]"
}

func newSliceEncoder(t reflect.Type) encoderFunc {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:670
	_go_fuzz_dep_.CoverTab[188162]++

											if t.Elem().Kind() == reflect.Uint8 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:672
		_go_fuzz_dep_.CoverTab[188164]++
												return encodeByteSlice
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:673
		// _ = "end of CoverTab[188164]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:674
		_go_fuzz_dep_.CoverTab[188165]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:674
		// _ = "end of CoverTab[188165]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:674
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:674
	// _ = "end of CoverTab[188162]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:674
	_go_fuzz_dep_.CoverTab[188163]++
											enc := &sliceEncoder{newArrayEncoder(t)}
											return enc.encode
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:676
	// _ = "end of CoverTab[188163]"
}

type arrayEncoder struct {
	elemEnc encoderFunc
}

func (ae *arrayEncoder) encode(e *encodeState, v reflect.Value, _ bool) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:683
	_go_fuzz_dep_.CoverTab[188166]++
											e.WriteByte('[')
											n := v.Len()
											for i := 0; i < n; i++ {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:686
		_go_fuzz_dep_.CoverTab[188168]++
												if i > 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:687
			_go_fuzz_dep_.CoverTab[188170]++
													e.WriteByte(',')
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:688
			// _ = "end of CoverTab[188170]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:689
			_go_fuzz_dep_.CoverTab[188171]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:689
			// _ = "end of CoverTab[188171]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:689
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:689
		// _ = "end of CoverTab[188168]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:689
		_go_fuzz_dep_.CoverTab[188169]++
												ae.elemEnc(e, v.Index(i), false)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:690
		// _ = "end of CoverTab[188169]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:691
	// _ = "end of CoverTab[188166]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:691
	_go_fuzz_dep_.CoverTab[188167]++
											e.WriteByte(']')
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:692
	// _ = "end of CoverTab[188167]"
}

func newArrayEncoder(t reflect.Type) encoderFunc {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:695
	_go_fuzz_dep_.CoverTab[188172]++
											enc := &arrayEncoder{typeEncoder(t.Elem())}
											return enc.encode
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:697
	// _ = "end of CoverTab[188172]"
}

type ptrEncoder struct {
	elemEnc encoderFunc
}

func (pe *ptrEncoder) encode(e *encodeState, v reflect.Value, quoted bool) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:704
	_go_fuzz_dep_.CoverTab[188173]++
											if v.IsNil() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:705
		_go_fuzz_dep_.CoverTab[188175]++
												e.WriteString("null")
												return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:707
		// _ = "end of CoverTab[188175]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:708
		_go_fuzz_dep_.CoverTab[188176]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:708
		// _ = "end of CoverTab[188176]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:708
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:708
	// _ = "end of CoverTab[188173]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:708
	_go_fuzz_dep_.CoverTab[188174]++
											pe.elemEnc(e, v.Elem(), quoted)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:709
	// _ = "end of CoverTab[188174]"
}

func newPtrEncoder(t reflect.Type) encoderFunc {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:712
	_go_fuzz_dep_.CoverTab[188177]++
											enc := &ptrEncoder{typeEncoder(t.Elem())}
											return enc.encode
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:714
	// _ = "end of CoverTab[188177]"
}

type condAddrEncoder struct {
	canAddrEnc, elseEnc encoderFunc
}

func (ce *condAddrEncoder) encode(e *encodeState, v reflect.Value, quoted bool) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:721
	_go_fuzz_dep_.CoverTab[188178]++
											if v.CanAddr() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:722
		_go_fuzz_dep_.CoverTab[188179]++
												ce.canAddrEnc(e, v, quoted)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:723
		// _ = "end of CoverTab[188179]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:724
		_go_fuzz_dep_.CoverTab[188180]++
												ce.elseEnc(e, v, quoted)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:725
		// _ = "end of CoverTab[188180]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:726
	// _ = "end of CoverTab[188178]"
}

// newCondAddrEncoder returns an encoder that checks whether its value
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:729
// CanAddr and delegates to canAddrEnc if so, else to elseEnc.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:731
func newCondAddrEncoder(canAddrEnc, elseEnc encoderFunc) encoderFunc {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:731
	_go_fuzz_dep_.CoverTab[188181]++
											enc := &condAddrEncoder{canAddrEnc: canAddrEnc, elseEnc: elseEnc}
											return enc.encode
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:733
	// _ = "end of CoverTab[188181]"
}

func isValidTag(s string) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:736
	_go_fuzz_dep_.CoverTab[188182]++
											if s == "" {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:737
		_go_fuzz_dep_.CoverTab[188185]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:738
		// _ = "end of CoverTab[188185]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:739
		_go_fuzz_dep_.CoverTab[188186]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:739
		// _ = "end of CoverTab[188186]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:739
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:739
	// _ = "end of CoverTab[188182]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:739
	_go_fuzz_dep_.CoverTab[188183]++
											for _, c := range s {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:740
		_go_fuzz_dep_.CoverTab[188187]++
												switch {
		case strings.ContainsRune("!#$%&()*+-./:<=>?@[]^_{|}~ ", c):
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:742
			_go_fuzz_dep_.CoverTab[188188]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:742
			// _ = "end of CoverTab[188188]"

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:746
		default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:746
			_go_fuzz_dep_.CoverTab[188189]++
													if !unicode.IsLetter(c) && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:747
				_go_fuzz_dep_.CoverTab[188190]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:747
				return !unicode.IsDigit(c)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:747
				// _ = "end of CoverTab[188190]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:747
			}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:747
				_go_fuzz_dep_.CoverTab[188191]++
														return false
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:748
				// _ = "end of CoverTab[188191]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:749
				_go_fuzz_dep_.CoverTab[188192]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:749
				// _ = "end of CoverTab[188192]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:749
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:749
			// _ = "end of CoverTab[188189]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:750
		// _ = "end of CoverTab[188187]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:751
	// _ = "end of CoverTab[188183]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:751
	_go_fuzz_dep_.CoverTab[188184]++
											return true
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:752
	// _ = "end of CoverTab[188184]"
}

func fieldByIndex(v reflect.Value, index []int) reflect.Value {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:755
	_go_fuzz_dep_.CoverTab[188193]++
											for _, i := range index {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:756
		_go_fuzz_dep_.CoverTab[188195]++
												if v.Kind() == reflect.Ptr {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:757
			_go_fuzz_dep_.CoverTab[188197]++
													if v.IsNil() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:758
				_go_fuzz_dep_.CoverTab[188199]++
														return reflect.Value{}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:759
				// _ = "end of CoverTab[188199]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:760
				_go_fuzz_dep_.CoverTab[188200]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:760
				// _ = "end of CoverTab[188200]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:760
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:760
			// _ = "end of CoverTab[188197]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:760
			_go_fuzz_dep_.CoverTab[188198]++
													v = v.Elem()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:761
			// _ = "end of CoverTab[188198]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:762
			_go_fuzz_dep_.CoverTab[188201]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:762
			// _ = "end of CoverTab[188201]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:762
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:762
		// _ = "end of CoverTab[188195]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:762
		_go_fuzz_dep_.CoverTab[188196]++
												v = v.Field(i)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:763
		// _ = "end of CoverTab[188196]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:764
	// _ = "end of CoverTab[188193]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:764
	_go_fuzz_dep_.CoverTab[188194]++
											return v
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:765
	// _ = "end of CoverTab[188194]"
}

func typeByIndex(t reflect.Type, index []int) reflect.Type {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:768
	_go_fuzz_dep_.CoverTab[188202]++
											for _, i := range index {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:769
		_go_fuzz_dep_.CoverTab[188204]++
												if t.Kind() == reflect.Ptr {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:770
			_go_fuzz_dep_.CoverTab[188206]++
													t = t.Elem()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:771
			// _ = "end of CoverTab[188206]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:772
			_go_fuzz_dep_.CoverTab[188207]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:772
			// _ = "end of CoverTab[188207]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:772
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:772
		// _ = "end of CoverTab[188204]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:772
		_go_fuzz_dep_.CoverTab[188205]++
												t = t.Field(i).Type
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:773
		// _ = "end of CoverTab[188205]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:774
	// _ = "end of CoverTab[188202]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:774
	_go_fuzz_dep_.CoverTab[188203]++
											return t
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:775
	// _ = "end of CoverTab[188203]"
}

// stringValues is a slice of reflect.Value holding *reflect.StringValue.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:778
// It implements the methods to sort by string.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:780
type stringValues []reflect.Value

func (sv stringValues) Len() int {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:782
	_go_fuzz_dep_.CoverTab[188208]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:782
	return len(sv)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:782
	// _ = "end of CoverTab[188208]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:782
}
func (sv stringValues) Swap(i, j int) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:783
	_go_fuzz_dep_.CoverTab[188209]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:783
	sv[i], sv[j] = sv[j], sv[i]
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:783
	// _ = "end of CoverTab[188209]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:783
}
func (sv stringValues) Less(i, j int) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:784
	_go_fuzz_dep_.CoverTab[188210]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:784
	return sv.get(i) < sv.get(j)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:784
	// _ = "end of CoverTab[188210]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:784
}
func (sv stringValues) get(i int) string {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:785
	_go_fuzz_dep_.CoverTab[188211]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:785
	return sv[i].String()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:785
	// _ = "end of CoverTab[188211]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:785
}

// NOTE: keep in sync with stringBytes below.
func (e *encodeState) string(s string) int {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:788
	_go_fuzz_dep_.CoverTab[188212]++
											len0 := e.Len()
											e.WriteByte('"')
											start := 0
											for i := 0; i < len(s); {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:792
		_go_fuzz_dep_.CoverTab[188215]++
												if b := s[i]; b < utf8.RuneSelf {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:793
			_go_fuzz_dep_.CoverTab[188219]++
													if 0x20 <= b && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:794
				_go_fuzz_dep_.CoverTab[188223]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:794
				return b != '\\'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:794
				// _ = "end of CoverTab[188223]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:794
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:794
				_go_fuzz_dep_.CoverTab[188224]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:794
				return b != '"'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:794
				// _ = "end of CoverTab[188224]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:794
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:794
				_go_fuzz_dep_.CoverTab[188225]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:794
				return b != '<'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:794
				// _ = "end of CoverTab[188225]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:794
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:794
				_go_fuzz_dep_.CoverTab[188226]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:794
				return b != '>'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:794
				// _ = "end of CoverTab[188226]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:794
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:794
				_go_fuzz_dep_.CoverTab[188227]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:794
				return b != '&'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:794
				// _ = "end of CoverTab[188227]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:794
			}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:794
				_go_fuzz_dep_.CoverTab[188228]++
														i++
														continue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:796
				// _ = "end of CoverTab[188228]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:797
				_go_fuzz_dep_.CoverTab[188229]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:797
				// _ = "end of CoverTab[188229]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:797
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:797
			// _ = "end of CoverTab[188219]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:797
			_go_fuzz_dep_.CoverTab[188220]++
													if start < i {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:798
				_go_fuzz_dep_.CoverTab[188230]++
														e.WriteString(s[start:i])
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:799
				// _ = "end of CoverTab[188230]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:800
				_go_fuzz_dep_.CoverTab[188231]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:800
				// _ = "end of CoverTab[188231]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:800
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:800
			// _ = "end of CoverTab[188220]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:800
			_go_fuzz_dep_.CoverTab[188221]++
													switch b {
			case '\\', '"':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:802
				_go_fuzz_dep_.CoverTab[188232]++
														e.WriteByte('\\')
														e.WriteByte(b)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:804
				// _ = "end of CoverTab[188232]"
			case '\n':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:805
				_go_fuzz_dep_.CoverTab[188233]++
														e.WriteByte('\\')
														e.WriteByte('n')
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:807
				// _ = "end of CoverTab[188233]"
			case '\r':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:808
				_go_fuzz_dep_.CoverTab[188234]++
														e.WriteByte('\\')
														e.WriteByte('r')
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:810
				// _ = "end of CoverTab[188234]"
			case '\t':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:811
				_go_fuzz_dep_.CoverTab[188235]++
														e.WriteByte('\\')
														e.WriteByte('t')
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:813
				// _ = "end of CoverTab[188235]"
			default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:814
				_go_fuzz_dep_.CoverTab[188236]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:819
				e.WriteString(`\u00`)
														e.WriteByte(hex[b>>4])
														e.WriteByte(hex[b&0xF])
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:821
				// _ = "end of CoverTab[188236]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:822
			// _ = "end of CoverTab[188221]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:822
			_go_fuzz_dep_.CoverTab[188222]++
													i++
													start = i
													continue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:825
			// _ = "end of CoverTab[188222]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:826
			_go_fuzz_dep_.CoverTab[188237]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:826
			// _ = "end of CoverTab[188237]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:826
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:826
		// _ = "end of CoverTab[188215]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:826
		_go_fuzz_dep_.CoverTab[188216]++
												c, size := utf8.DecodeRuneInString(s[i:])
												if c == utf8.RuneError && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:828
			_go_fuzz_dep_.CoverTab[188238]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:828
			return size == 1
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:828
			// _ = "end of CoverTab[188238]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:828
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:828
			_go_fuzz_dep_.CoverTab[188239]++
													if start < i {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:829
				_go_fuzz_dep_.CoverTab[188241]++
														e.WriteString(s[start:i])
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:830
				// _ = "end of CoverTab[188241]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:831
				_go_fuzz_dep_.CoverTab[188242]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:831
				// _ = "end of CoverTab[188242]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:831
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:831
			// _ = "end of CoverTab[188239]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:831
			_go_fuzz_dep_.CoverTab[188240]++
													e.WriteString(`\ufffd`)
													i += size
													start = i
													continue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:835
			// _ = "end of CoverTab[188240]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:836
			_go_fuzz_dep_.CoverTab[188243]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:836
			// _ = "end of CoverTab[188243]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:836
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:836
		// _ = "end of CoverTab[188216]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:836
		_go_fuzz_dep_.CoverTab[188217]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:844
		if c == '\u2028' || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:844
			_go_fuzz_dep_.CoverTab[188244]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:844
			return c == '\u2029'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:844
			// _ = "end of CoverTab[188244]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:844
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:844
			_go_fuzz_dep_.CoverTab[188245]++
													if start < i {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:845
				_go_fuzz_dep_.CoverTab[188247]++
														e.WriteString(s[start:i])
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:846
				// _ = "end of CoverTab[188247]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:847
				_go_fuzz_dep_.CoverTab[188248]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:847
				// _ = "end of CoverTab[188248]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:847
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:847
			// _ = "end of CoverTab[188245]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:847
			_go_fuzz_dep_.CoverTab[188246]++
													e.WriteString(`\u202`)
													e.WriteByte(hex[c&0xF])
													i += size
													start = i
													continue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:852
			// _ = "end of CoverTab[188246]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:853
			_go_fuzz_dep_.CoverTab[188249]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:853
			// _ = "end of CoverTab[188249]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:853
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:853
		// _ = "end of CoverTab[188217]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:853
		_go_fuzz_dep_.CoverTab[188218]++
												i += size
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:854
		// _ = "end of CoverTab[188218]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:855
	// _ = "end of CoverTab[188212]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:855
	_go_fuzz_dep_.CoverTab[188213]++
											if start < len(s) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:856
		_go_fuzz_dep_.CoverTab[188250]++
												e.WriteString(s[start:])
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:857
		// _ = "end of CoverTab[188250]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:858
		_go_fuzz_dep_.CoverTab[188251]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:858
		// _ = "end of CoverTab[188251]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:858
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:858
	// _ = "end of CoverTab[188213]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:858
	_go_fuzz_dep_.CoverTab[188214]++
											e.WriteByte('"')
											return e.Len() - len0
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:860
	// _ = "end of CoverTab[188214]"
}

// NOTE: keep in sync with string above.
func (e *encodeState) stringBytes(s []byte) int {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:864
	_go_fuzz_dep_.CoverTab[188252]++
											len0 := e.Len()
											e.WriteByte('"')
											start := 0
											for i := 0; i < len(s); {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:868
		_go_fuzz_dep_.CoverTab[188255]++
												if b := s[i]; b < utf8.RuneSelf {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:869
			_go_fuzz_dep_.CoverTab[188259]++
													if 0x20 <= b && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:870
				_go_fuzz_dep_.CoverTab[188263]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:870
				return b != '\\'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:870
				// _ = "end of CoverTab[188263]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:870
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:870
				_go_fuzz_dep_.CoverTab[188264]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:870
				return b != '"'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:870
				// _ = "end of CoverTab[188264]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:870
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:870
				_go_fuzz_dep_.CoverTab[188265]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:870
				return b != '<'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:870
				// _ = "end of CoverTab[188265]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:870
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:870
				_go_fuzz_dep_.CoverTab[188266]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:870
				return b != '>'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:870
				// _ = "end of CoverTab[188266]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:870
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:870
				_go_fuzz_dep_.CoverTab[188267]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:870
				return b != '&'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:870
				// _ = "end of CoverTab[188267]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:870
			}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:870
				_go_fuzz_dep_.CoverTab[188268]++
														i++
														continue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:872
				// _ = "end of CoverTab[188268]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:873
				_go_fuzz_dep_.CoverTab[188269]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:873
				// _ = "end of CoverTab[188269]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:873
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:873
			// _ = "end of CoverTab[188259]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:873
			_go_fuzz_dep_.CoverTab[188260]++
													if start < i {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:874
				_go_fuzz_dep_.CoverTab[188270]++
														e.Write(s[start:i])
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:875
				// _ = "end of CoverTab[188270]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:876
				_go_fuzz_dep_.CoverTab[188271]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:876
				// _ = "end of CoverTab[188271]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:876
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:876
			// _ = "end of CoverTab[188260]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:876
			_go_fuzz_dep_.CoverTab[188261]++
													switch b {
			case '\\', '"':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:878
				_go_fuzz_dep_.CoverTab[188272]++
														e.WriteByte('\\')
														e.WriteByte(b)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:880
				// _ = "end of CoverTab[188272]"
			case '\n':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:881
				_go_fuzz_dep_.CoverTab[188273]++
														e.WriteByte('\\')
														e.WriteByte('n')
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:883
				// _ = "end of CoverTab[188273]"
			case '\r':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:884
				_go_fuzz_dep_.CoverTab[188274]++
														e.WriteByte('\\')
														e.WriteByte('r')
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:886
				// _ = "end of CoverTab[188274]"
			case '\t':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:887
				_go_fuzz_dep_.CoverTab[188275]++
														e.WriteByte('\\')
														e.WriteByte('t')
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:889
				// _ = "end of CoverTab[188275]"
			default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:890
				_go_fuzz_dep_.CoverTab[188276]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:895
				e.WriteString(`\u00`)
														e.WriteByte(hex[b>>4])
														e.WriteByte(hex[b&0xF])
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:897
				// _ = "end of CoverTab[188276]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:898
			// _ = "end of CoverTab[188261]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:898
			_go_fuzz_dep_.CoverTab[188262]++
													i++
													start = i
													continue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:901
			// _ = "end of CoverTab[188262]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:902
			_go_fuzz_dep_.CoverTab[188277]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:902
			// _ = "end of CoverTab[188277]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:902
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:902
		// _ = "end of CoverTab[188255]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:902
		_go_fuzz_dep_.CoverTab[188256]++
												c, size := utf8.DecodeRune(s[i:])
												if c == utf8.RuneError && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:904
			_go_fuzz_dep_.CoverTab[188278]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:904
			return size == 1
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:904
			// _ = "end of CoverTab[188278]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:904
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:904
			_go_fuzz_dep_.CoverTab[188279]++
													if start < i {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:905
				_go_fuzz_dep_.CoverTab[188281]++
														e.Write(s[start:i])
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:906
				// _ = "end of CoverTab[188281]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:907
				_go_fuzz_dep_.CoverTab[188282]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:907
				// _ = "end of CoverTab[188282]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:907
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:907
			// _ = "end of CoverTab[188279]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:907
			_go_fuzz_dep_.CoverTab[188280]++
													e.WriteString(`\ufffd`)
													i += size
													start = i
													continue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:911
			// _ = "end of CoverTab[188280]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:912
			_go_fuzz_dep_.CoverTab[188283]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:912
			// _ = "end of CoverTab[188283]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:912
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:912
		// _ = "end of CoverTab[188256]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:912
		_go_fuzz_dep_.CoverTab[188257]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:920
		if c == '\u2028' || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:920
			_go_fuzz_dep_.CoverTab[188284]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:920
			return c == '\u2029'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:920
			// _ = "end of CoverTab[188284]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:920
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:920
			_go_fuzz_dep_.CoverTab[188285]++
													if start < i {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:921
				_go_fuzz_dep_.CoverTab[188287]++
														e.Write(s[start:i])
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:922
				// _ = "end of CoverTab[188287]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:923
				_go_fuzz_dep_.CoverTab[188288]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:923
				// _ = "end of CoverTab[188288]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:923
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:923
			// _ = "end of CoverTab[188285]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:923
			_go_fuzz_dep_.CoverTab[188286]++
													e.WriteString(`\u202`)
													e.WriteByte(hex[c&0xF])
													i += size
													start = i
													continue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:928
			// _ = "end of CoverTab[188286]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:929
			_go_fuzz_dep_.CoverTab[188289]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:929
			// _ = "end of CoverTab[188289]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:929
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:929
		// _ = "end of CoverTab[188257]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:929
		_go_fuzz_dep_.CoverTab[188258]++
												i += size
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:930
		// _ = "end of CoverTab[188258]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:931
	// _ = "end of CoverTab[188252]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:931
	_go_fuzz_dep_.CoverTab[188253]++
											if start < len(s) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:932
		_go_fuzz_dep_.CoverTab[188290]++
												e.Write(s[start:])
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:933
		// _ = "end of CoverTab[188290]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:934
		_go_fuzz_dep_.CoverTab[188291]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:934
		// _ = "end of CoverTab[188291]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:934
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:934
	// _ = "end of CoverTab[188253]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:934
	_go_fuzz_dep_.CoverTab[188254]++
											e.WriteByte('"')
											return e.Len() - len0
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:936
	// _ = "end of CoverTab[188254]"
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
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:951
	_go_fuzz_dep_.CoverTab[188292]++
											f.nameBytes = []byte(f.name)
											return f
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:953
	// _ = "end of CoverTab[188292]"
}

// byName sorts field by name, breaking ties with depth,
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:956
// then breaking ties with "name came from json tag", then
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:956
// breaking ties with index sequence.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:959
type byName []field

func (x byName) Len() int {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:961
	_go_fuzz_dep_.CoverTab[188293]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:961
	return len(x)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:961
	// _ = "end of CoverTab[188293]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:961
}

func (x byName) Swap(i, j int) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:963
	_go_fuzz_dep_.CoverTab[188294]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:963
	x[i], x[j] = x[j], x[i]
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:963
	// _ = "end of CoverTab[188294]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:963
}

func (x byName) Less(i, j int) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:965
	_go_fuzz_dep_.CoverTab[188295]++
											if x[i].name != x[j].name {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:966
		_go_fuzz_dep_.CoverTab[188299]++
												return x[i].name < x[j].name
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:967
		// _ = "end of CoverTab[188299]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:968
		_go_fuzz_dep_.CoverTab[188300]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:968
		// _ = "end of CoverTab[188300]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:968
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:968
	// _ = "end of CoverTab[188295]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:968
	_go_fuzz_dep_.CoverTab[188296]++
											if len(x[i].index) != len(x[j].index) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:969
		_go_fuzz_dep_.CoverTab[188301]++
												return len(x[i].index) < len(x[j].index)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:970
		// _ = "end of CoverTab[188301]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:971
		_go_fuzz_dep_.CoverTab[188302]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:971
		// _ = "end of CoverTab[188302]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:971
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:971
	// _ = "end of CoverTab[188296]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:971
	_go_fuzz_dep_.CoverTab[188297]++
											if x[i].tag != x[j].tag {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:972
		_go_fuzz_dep_.CoverTab[188303]++
												return x[i].tag
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:973
		// _ = "end of CoverTab[188303]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:974
		_go_fuzz_dep_.CoverTab[188304]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:974
		// _ = "end of CoverTab[188304]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:974
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:974
	// _ = "end of CoverTab[188297]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:974
	_go_fuzz_dep_.CoverTab[188298]++
											return byIndex(x).Less(i, j)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:975
	// _ = "end of CoverTab[188298]"
}

// byIndex sorts field by index sequence.
type byIndex []field

func (x byIndex) Len() int {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:981
	_go_fuzz_dep_.CoverTab[188305]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:981
	return len(x)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:981
	// _ = "end of CoverTab[188305]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:981
}

func (x byIndex) Swap(i, j int) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:983
	_go_fuzz_dep_.CoverTab[188306]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:983
	x[i], x[j] = x[j], x[i]
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:983
	// _ = "end of CoverTab[188306]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:983
}

func (x byIndex) Less(i, j int) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:985
	_go_fuzz_dep_.CoverTab[188307]++
											for k, xik := range x[i].index {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:986
		_go_fuzz_dep_.CoverTab[188309]++
												if k >= len(x[j].index) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:987
			_go_fuzz_dep_.CoverTab[188311]++
													return false
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:988
			// _ = "end of CoverTab[188311]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:989
			_go_fuzz_dep_.CoverTab[188312]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:989
			// _ = "end of CoverTab[188312]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:989
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:989
		// _ = "end of CoverTab[188309]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:989
		_go_fuzz_dep_.CoverTab[188310]++
												if xik != x[j].index[k] {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:990
			_go_fuzz_dep_.CoverTab[188313]++
													return xik < x[j].index[k]
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:991
			// _ = "end of CoverTab[188313]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:992
			_go_fuzz_dep_.CoverTab[188314]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:992
			// _ = "end of CoverTab[188314]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:992
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:992
		// _ = "end of CoverTab[188310]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:993
	// _ = "end of CoverTab[188307]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:993
	_go_fuzz_dep_.CoverTab[188308]++
											return len(x[i].index) < len(x[j].index)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:994
	// _ = "end of CoverTab[188308]"
}

// typeFields returns a list of fields that JSON should recognize for the given type.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:997
// The algorithm is breadth-first search over the set of structs to include - the top struct
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:997
// and then any reachable anonymous structs.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1000
func typeFields(t reflect.Type) []field {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1000
	_go_fuzz_dep_.CoverTab[188315]++

												current := []field{}
												next := []field{{typ: t}}

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1006
	count := map[reflect.Type]int{}
												nextCount := map[reflect.Type]int{}

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1010
	visited := map[reflect.Type]bool{}

	// Fields found.
	var fields []field

	for len(next) > 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1015
		_go_fuzz_dep_.CoverTab[188318]++
													current, next = next, current[:0]
													count, nextCount = nextCount, map[reflect.Type]int{}

													for _, f := range current {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1019
			_go_fuzz_dep_.CoverTab[188319]++
														if visited[f.typ] {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1020
				_go_fuzz_dep_.CoverTab[188321]++
															continue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1021
				// _ = "end of CoverTab[188321]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1022
				_go_fuzz_dep_.CoverTab[188322]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1022
				// _ = "end of CoverTab[188322]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1022
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1022
			// _ = "end of CoverTab[188319]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1022
			_go_fuzz_dep_.CoverTab[188320]++
														visited[f.typ] = true

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1026
			for i := 0; i < f.typ.NumField(); i++ {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1026
				_go_fuzz_dep_.CoverTab[188323]++
															sf := f.typ.Field(i)
															if sf.PkgPath != "" && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1028
					_go_fuzz_dep_.CoverTab[188330]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1028
					return !sf.Anonymous
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1028
					// _ = "end of CoverTab[188330]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1028
				}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1028
					_go_fuzz_dep_.CoverTab[188331]++
																continue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1029
					// _ = "end of CoverTab[188331]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1030
					_go_fuzz_dep_.CoverTab[188332]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1030
					// _ = "end of CoverTab[188332]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1030
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1030
				// _ = "end of CoverTab[188323]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1030
				_go_fuzz_dep_.CoverTab[188324]++
															tag := sf.Tag.Get("json")
															if tag == "-" {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1032
					_go_fuzz_dep_.CoverTab[188333]++
																continue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1033
					// _ = "end of CoverTab[188333]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1034
					_go_fuzz_dep_.CoverTab[188334]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1034
					// _ = "end of CoverTab[188334]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1034
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1034
				// _ = "end of CoverTab[188324]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1034
				_go_fuzz_dep_.CoverTab[188325]++
															name, opts := parseTag(tag)
															if !isValidTag(name) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1036
					_go_fuzz_dep_.CoverTab[188335]++
																name = ""
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1037
					// _ = "end of CoverTab[188335]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1038
					_go_fuzz_dep_.CoverTab[188336]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1038
					// _ = "end of CoverTab[188336]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1038
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1038
				// _ = "end of CoverTab[188325]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1038
				_go_fuzz_dep_.CoverTab[188326]++
															index := make([]int, len(f.index)+1)
															copy(index, f.index)
															index[len(f.index)] = i

															ft := sf.Type
															if ft.Name() == "" && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1044
					_go_fuzz_dep_.CoverTab[188337]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1044
					return ft.Kind() == reflect.Ptr
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1044
					// _ = "end of CoverTab[188337]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1044
				}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1044
					_go_fuzz_dep_.CoverTab[188338]++

																ft = ft.Elem()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1046
					// _ = "end of CoverTab[188338]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1047
					_go_fuzz_dep_.CoverTab[188339]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1047
					// _ = "end of CoverTab[188339]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1047
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1047
				// _ = "end of CoverTab[188326]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1047
				_go_fuzz_dep_.CoverTab[188327]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1050
				quoted := false
				if opts.Contains("string") {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1051
					_go_fuzz_dep_.CoverTab[188340]++
																switch ft.Kind() {
					case reflect.Bool,
						reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
						reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
						reflect.Float32, reflect.Float64,
						reflect.String:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1057
						_go_fuzz_dep_.CoverTab[188341]++
																	quoted = true
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1058
						// _ = "end of CoverTab[188341]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1058
					default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1058
						_go_fuzz_dep_.CoverTab[188342]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1058
						// _ = "end of CoverTab[188342]"
					}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1059
					// _ = "end of CoverTab[188340]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1060
					_go_fuzz_dep_.CoverTab[188343]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1060
					// _ = "end of CoverTab[188343]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1060
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1060
				// _ = "end of CoverTab[188327]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1060
				_go_fuzz_dep_.CoverTab[188328]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1063
				if name != "" || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1063
					_go_fuzz_dep_.CoverTab[188344]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1063
					return !sf.Anonymous
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1063
					// _ = "end of CoverTab[188344]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1063
				}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1063
					_go_fuzz_dep_.CoverTab[188345]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1063
					return ft.Kind() != reflect.Struct
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1063
					// _ = "end of CoverTab[188345]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1063
				}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1063
					_go_fuzz_dep_.CoverTab[188346]++
																tagged := name != ""
																if name == "" {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1065
						_go_fuzz_dep_.CoverTab[188349]++
																	name = sf.Name
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1066
						// _ = "end of CoverTab[188349]"
					} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1067
						_go_fuzz_dep_.CoverTab[188350]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1067
						// _ = "end of CoverTab[188350]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1067
					}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1067
					// _ = "end of CoverTab[188346]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1067
					_go_fuzz_dep_.CoverTab[188347]++
																fields = append(fields, fillField(field{
						name:		name,
						tag:		tagged,
						index:		index,
						typ:		ft,
						omitEmpty:	opts.Contains("omitempty"),
						quoted:		quoted,
					}))
					if count[f.typ] > 1 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1076
						_go_fuzz_dep_.CoverTab[188351]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1081
						fields = append(fields, fields[len(fields)-1])
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1081
						// _ = "end of CoverTab[188351]"
					} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1082
						_go_fuzz_dep_.CoverTab[188352]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1082
						// _ = "end of CoverTab[188352]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1082
					}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1082
					// _ = "end of CoverTab[188347]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1082
					_go_fuzz_dep_.CoverTab[188348]++
																continue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1083
					// _ = "end of CoverTab[188348]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1084
					_go_fuzz_dep_.CoverTab[188353]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1084
					// _ = "end of CoverTab[188353]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1084
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1084
				// _ = "end of CoverTab[188328]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1084
				_go_fuzz_dep_.CoverTab[188329]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1087
				nextCount[ft]++
				if nextCount[ft] == 1 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1088
					_go_fuzz_dep_.CoverTab[188354]++
																next = append(next, fillField(field{name: ft.Name(), index: index, typ: ft}))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1089
					// _ = "end of CoverTab[188354]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1090
					_go_fuzz_dep_.CoverTab[188355]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1090
					// _ = "end of CoverTab[188355]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1090
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1090
				// _ = "end of CoverTab[188329]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1091
			// _ = "end of CoverTab[188320]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1092
		// _ = "end of CoverTab[188318]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1093
	// _ = "end of CoverTab[188315]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1093
	_go_fuzz_dep_.CoverTab[188316]++

												sort.Sort(byName(fields))

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1103
	out := fields[:0]
	for advance, i := 0, 0; i < len(fields); i += advance {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1104
		_go_fuzz_dep_.CoverTab[188356]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1107
		fi := fields[i]
		name := fi.name
		for advance = 1; i+advance < len(fields); advance++ {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1109
			_go_fuzz_dep_.CoverTab[188359]++
														fj := fields[i+advance]
														if fj.name != name {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1111
				_go_fuzz_dep_.CoverTab[188360]++
															break
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1112
				// _ = "end of CoverTab[188360]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1113
				_go_fuzz_dep_.CoverTab[188361]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1113
				// _ = "end of CoverTab[188361]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1113
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1113
			// _ = "end of CoverTab[188359]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1114
		// _ = "end of CoverTab[188356]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1114
		_go_fuzz_dep_.CoverTab[188357]++
													if advance == 1 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1115
			_go_fuzz_dep_.CoverTab[188362]++
														out = append(out, fi)
														continue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1117
			// _ = "end of CoverTab[188362]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1118
			_go_fuzz_dep_.CoverTab[188363]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1118
			// _ = "end of CoverTab[188363]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1118
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1118
		// _ = "end of CoverTab[188357]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1118
		_go_fuzz_dep_.CoverTab[188358]++
													dominant, ok := dominantField(fields[i : i+advance])
													if ok {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1120
			_go_fuzz_dep_.CoverTab[188364]++
														out = append(out, dominant)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1121
			// _ = "end of CoverTab[188364]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1122
			_go_fuzz_dep_.CoverTab[188365]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1122
			// _ = "end of CoverTab[188365]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1122
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1122
		// _ = "end of CoverTab[188358]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1123
	// _ = "end of CoverTab[188316]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1123
	_go_fuzz_dep_.CoverTab[188317]++

												fields = out
												sort.Sort(byIndex(fields))

												return fields
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1128
	// _ = "end of CoverTab[188317]"
}

// dominantField looks through the fields, all of which are known to
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1131
// have the same name, to find the single field that dominates the
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1131
// others using Go's embedding rules, modified by the presence of
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1131
// JSON tags. If there are multiple top-level fields, the boolean
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1131
// will be false: This condition is an error in Go and we skip all
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1131
// the fields.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1137
func dominantField(fields []field) (field, bool) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1137
	_go_fuzz_dep_.CoverTab[188366]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1141
	length := len(fields[0].index)
	tagged := -1
	for i, f := range fields {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1143
		_go_fuzz_dep_.CoverTab[188370]++
													if len(f.index) > length {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1144
			_go_fuzz_dep_.CoverTab[188372]++
														fields = fields[:i]
														break
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1146
			// _ = "end of CoverTab[188372]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1147
			_go_fuzz_dep_.CoverTab[188373]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1147
			// _ = "end of CoverTab[188373]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1147
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1147
		// _ = "end of CoverTab[188370]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1147
		_go_fuzz_dep_.CoverTab[188371]++
													if f.tag {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1148
			_go_fuzz_dep_.CoverTab[188374]++
														if tagged >= 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1149
				_go_fuzz_dep_.CoverTab[188376]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1152
				return field{}, false
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1152
				// _ = "end of CoverTab[188376]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1153
				_go_fuzz_dep_.CoverTab[188377]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1153
				// _ = "end of CoverTab[188377]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1153
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1153
			// _ = "end of CoverTab[188374]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1153
			_go_fuzz_dep_.CoverTab[188375]++
														tagged = i
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1154
			// _ = "end of CoverTab[188375]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1155
			_go_fuzz_dep_.CoverTab[188378]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1155
			// _ = "end of CoverTab[188378]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1155
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1155
		// _ = "end of CoverTab[188371]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1156
	// _ = "end of CoverTab[188366]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1156
	_go_fuzz_dep_.CoverTab[188367]++
												if tagged >= 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1157
		_go_fuzz_dep_.CoverTab[188379]++
													return fields[tagged], true
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1158
		// _ = "end of CoverTab[188379]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1159
		_go_fuzz_dep_.CoverTab[188380]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1159
		// _ = "end of CoverTab[188380]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1159
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1159
	// _ = "end of CoverTab[188367]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1159
	_go_fuzz_dep_.CoverTab[188368]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1163
	if len(fields) > 1 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1163
		_go_fuzz_dep_.CoverTab[188381]++
													return field{}, false
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1164
		// _ = "end of CoverTab[188381]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1165
		_go_fuzz_dep_.CoverTab[188382]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1165
		// _ = "end of CoverTab[188382]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1165
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1165
	// _ = "end of CoverTab[188368]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1165
	_go_fuzz_dep_.CoverTab[188369]++
												return fields[0], true
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1166
	// _ = "end of CoverTab[188369]"
}

var fieldCache struct {
	sync.RWMutex
	m	map[reflect.Type][]field
}

// cachedTypeFields is like typeFields but uses a cache to avoid repeated work.
func cachedTypeFields(t reflect.Type) []field {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1175
	_go_fuzz_dep_.CoverTab[188383]++
												fieldCache.RLock()
												f := fieldCache.m[t]
												fieldCache.RUnlock()
												if f != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1179
		_go_fuzz_dep_.CoverTab[188387]++
													return f
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1180
		// _ = "end of CoverTab[188387]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1181
		_go_fuzz_dep_.CoverTab[188388]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1181
		// _ = "end of CoverTab[188388]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1181
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1181
	// _ = "end of CoverTab[188383]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1181
	_go_fuzz_dep_.CoverTab[188384]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1185
	f = typeFields(t)
	if f == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1186
		_go_fuzz_dep_.CoverTab[188389]++
													f = []field{}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1187
		// _ = "end of CoverTab[188389]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1188
		_go_fuzz_dep_.CoverTab[188390]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1188
		// _ = "end of CoverTab[188390]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1188
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1188
	// _ = "end of CoverTab[188384]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1188
	_go_fuzz_dep_.CoverTab[188385]++

												fieldCache.Lock()
												if fieldCache.m == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1191
		_go_fuzz_dep_.CoverTab[188391]++
													fieldCache.m = map[reflect.Type][]field{}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1192
		// _ = "end of CoverTab[188391]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1193
		_go_fuzz_dep_.CoverTab[188392]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1193
		// _ = "end of CoverTab[188392]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1193
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1193
	// _ = "end of CoverTab[188385]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1193
	_go_fuzz_dep_.CoverTab[188386]++
												fieldCache.m[t] = f
												fieldCache.Unlock()
												return f
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1196
	// _ = "end of CoverTab[188386]"
}

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1197
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/encode.go:1197
var _ = _go_fuzz_dep_.CoverTab
