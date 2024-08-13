// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Represents JSON data structure using native Go types: booleans, floats,
// strings, arrays, and maps.

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:8
package json

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:8
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:8
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:8
)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:8
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:8
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:8
)

import (
	"bytes"
	"encoding"
	"encoding/base64"
	"errors"
	"fmt"
	"reflect"
	"runtime"
	"strconv"
	"unicode"
	"unicode/utf16"
	"unicode/utf8"
)

// Unmarshal parses the JSON-encoded data and stores the result
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:24
// in the value pointed to by v.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:24
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:24
// Unmarshal uses the inverse of the encodings that
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:24
// Marshal uses, allocating maps, slices, and pointers as necessary,
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:24
// with the following additional rules:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:24
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:24
// To unmarshal JSON into a pointer, Unmarshal first handles the case of
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:24
// the JSON being the JSON literal null.  In that case, Unmarshal sets
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:24
// the pointer to nil.  Otherwise, Unmarshal unmarshals the JSON into
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:24
// the value pointed at by the pointer.  If the pointer is nil, Unmarshal
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:24
// allocates a new value for it to point to.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:24
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:24
// To unmarshal JSON into a struct, Unmarshal matches incoming object
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:24
// keys to the keys used by Marshal (either the struct field name or its tag),
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:24
// preferring an exact match but also accepting a case-insensitive match.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:24
// Unmarshal will only set exported fields of the struct.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:24
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:24
// To unmarshal JSON into an interface value,
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:24
// Unmarshal stores one of these in the interface value:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:24
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:24
//	bool, for JSON booleans
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:24
//	float64, for JSON numbers
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:24
//	string, for JSON strings
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:24
//	[]interface{}, for JSON arrays
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:24
//	map[string]interface{}, for JSON objects
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:24
//	nil for JSON null
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:24
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:24
// To unmarshal a JSON array into a slice, Unmarshal resets the slice length
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:24
// to zero and then appends each element to the slice.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:24
// As a special case, to unmarshal an empty JSON array into a slice,
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:24
// Unmarshal replaces the slice with a new empty slice.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:24
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:24
// To unmarshal a JSON array into a Go array, Unmarshal decodes
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:24
// JSON array elements into corresponding Go array elements.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:24
// If the Go array is smaller than the JSON array,
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:24
// the additional JSON array elements are discarded.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:24
// If the JSON array is smaller than the Go array,
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:24
// the additional Go array elements are set to zero values.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:24
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:24
// To unmarshal a JSON object into a string-keyed map, Unmarshal first
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:24
// establishes a map to use, If the map is nil, Unmarshal allocates a new map.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:24
// Otherwise Unmarshal reuses the existing map, keeping existing entries.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:24
// Unmarshal then stores key-value pairs from the JSON object into the map.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:24
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:24
// If a JSON value is not appropriate for a given target type,
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:24
// or if a JSON number overflows the target type, Unmarshal
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:24
// skips that field and completes the unmarshaling as best it can.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:24
// If no more serious errors are encountered, Unmarshal returns
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:24
// an UnmarshalTypeError describing the earliest such error.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:24
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:24
// The JSON null value unmarshals into an interface, map, pointer, or slice
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:24
// by setting that Go value to nil. Because null is often used in JSON to mean
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:24
// “not present,” unmarshaling a JSON null into any other Go type has no effect
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:24
// on the value and produces no error.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:24
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:24
// When unmarshaling quoted strings, invalid UTF-8 or
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:24
// invalid UTF-16 surrogate pairs are not treated as an error.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:24
// Instead, they are replaced by the Unicode replacement
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:24
// character U+FFFD.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:85
func Unmarshal(data []byte, v interface{}) error {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:85
	_go_fuzz_dep_.CoverTab[187433]++
	// Check for well-formedness.
	// Avoids filling out half a data structure
	// before discovering a JSON syntax error.
	var d decodeState
	err := checkValid(data, &d.scan)
	if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:91
		_go_fuzz_dep_.CoverTab[187435]++
												return err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:92
		// _ = "end of CoverTab[187435]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:93
		_go_fuzz_dep_.CoverTab[187436]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:93
		// _ = "end of CoverTab[187436]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:93
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:93
	// _ = "end of CoverTab[187433]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:93
	_go_fuzz_dep_.CoverTab[187434]++

											d.init(data)
											return d.unmarshal(v)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:96
	// _ = "end of CoverTab[187434]"
}

// Unmarshaler is the interface implemented by objects
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:99
// that can unmarshal a JSON description of themselves.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:99
// The input can be assumed to be a valid encoding of
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:99
// a JSON value. UnmarshalJSON must copy the JSON data
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:99
// if it wishes to retain the data after returning.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:104
type Unmarshaler interface {
	UnmarshalJSON([]byte) error
}

// An UnmarshalTypeError describes a JSON value that was
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:108
// not appropriate for a value of a specific Go type.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:110
type UnmarshalTypeError struct {
	Value	string		// description of JSON value - "bool", "array", "number -5"
	Type	reflect.Type	// type of Go value it could not be assigned to
	Offset	int64		// error occurred after reading Offset bytes
}

func (e *UnmarshalTypeError) Error() string {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:116
	_go_fuzz_dep_.CoverTab[187437]++
											return "json: cannot unmarshal " + e.Value + " into Go value of type " + e.Type.String()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:117
	// _ = "end of CoverTab[187437]"
}

// An UnmarshalFieldError describes a JSON object key that
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:120
// led to an unexported (and therefore unwritable) struct field.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:120
// (No longer used; kept for compatibility.)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:123
type UnmarshalFieldError struct {
	Key	string
	Type	reflect.Type
	Field	reflect.StructField
}

func (e *UnmarshalFieldError) Error() string {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:129
	_go_fuzz_dep_.CoverTab[187438]++
											return "json: cannot unmarshal object key " + strconv.Quote(e.Key) + " into unexported field " + e.Field.Name + " of type " + e.Type.String()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:130
	// _ = "end of CoverTab[187438]"
}

// An InvalidUnmarshalError describes an invalid argument passed to Unmarshal.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:133
// (The argument to Unmarshal must be a non-nil pointer.)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:135
type InvalidUnmarshalError struct {
	Type reflect.Type
}

func (e *InvalidUnmarshalError) Error() string {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:139
	_go_fuzz_dep_.CoverTab[187439]++
											if e.Type == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:140
		_go_fuzz_dep_.CoverTab[187442]++
												return "json: Unmarshal(nil)"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:141
		// _ = "end of CoverTab[187442]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:142
		_go_fuzz_dep_.CoverTab[187443]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:142
		// _ = "end of CoverTab[187443]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:142
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:142
	// _ = "end of CoverTab[187439]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:142
	_go_fuzz_dep_.CoverTab[187440]++

											if e.Type.Kind() != reflect.Ptr {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:144
		_go_fuzz_dep_.CoverTab[187444]++
												return "json: Unmarshal(non-pointer " + e.Type.String() + ")"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:145
		// _ = "end of CoverTab[187444]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:146
		_go_fuzz_dep_.CoverTab[187445]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:146
		// _ = "end of CoverTab[187445]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:146
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:146
	// _ = "end of CoverTab[187440]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:146
	_go_fuzz_dep_.CoverTab[187441]++
											return "json: Unmarshal(nil " + e.Type.String() + ")"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:147
	// _ = "end of CoverTab[187441]"
}

func (d *decodeState) unmarshal(v interface{}) (err error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:150
	_go_fuzz_dep_.CoverTab[187446]++
											defer func() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:151
		_go_fuzz_dep_.CoverTab[187449]++
												if r := recover(); r != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:152
			_go_fuzz_dep_.CoverTab[187450]++
													if _, ok := r.(runtime.Error); ok {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:153
				_go_fuzz_dep_.CoverTab[187452]++
														panic(r)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:154
				// _ = "end of CoverTab[187452]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:155
				_go_fuzz_dep_.CoverTab[187453]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:155
				// _ = "end of CoverTab[187453]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:155
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:155
			// _ = "end of CoverTab[187450]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:155
			_go_fuzz_dep_.CoverTab[187451]++
													err = r.(error)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:156
			// _ = "end of CoverTab[187451]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:157
			_go_fuzz_dep_.CoverTab[187454]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:157
			// _ = "end of CoverTab[187454]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:157
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:157
		// _ = "end of CoverTab[187449]"
	}()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:158
	// _ = "end of CoverTab[187446]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:158
	_go_fuzz_dep_.CoverTab[187447]++

											rv := reflect.ValueOf(v)
											if rv.Kind() != reflect.Ptr || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:161
		_go_fuzz_dep_.CoverTab[187455]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:161
		return rv.IsNil()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:161
		// _ = "end of CoverTab[187455]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:161
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:161
		_go_fuzz_dep_.CoverTab[187456]++
												return &InvalidUnmarshalError{reflect.TypeOf(v)}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:162
		// _ = "end of CoverTab[187456]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:163
		_go_fuzz_dep_.CoverTab[187457]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:163
		// _ = "end of CoverTab[187457]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:163
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:163
	// _ = "end of CoverTab[187447]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:163
	_go_fuzz_dep_.CoverTab[187448]++

											d.scan.reset()

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:168
	d.value(rv)
											return d.savedError
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:169
	// _ = "end of CoverTab[187448]"
}

// A Number represents a JSON number literal.
type Number string

// String returns the literal text of the number.
func (n Number) String() string {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:176
	_go_fuzz_dep_.CoverTab[187458]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:176
	return string(n)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:176
	// _ = "end of CoverTab[187458]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:176
}

// Float64 returns the number as a float64.
func (n Number) Float64() (float64, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:179
	_go_fuzz_dep_.CoverTab[187459]++
											return strconv.ParseFloat(string(n), 64)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:180
	// _ = "end of CoverTab[187459]"
}

// Int64 returns the number as an int64.
func (n Number) Int64() (int64, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:184
	_go_fuzz_dep_.CoverTab[187460]++
											return strconv.ParseInt(string(n), 10, 64)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:185
	// _ = "end of CoverTab[187460]"
}

// isValidNumber reports whether s is a valid JSON number literal.
func isValidNumber(s string) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:189
	_go_fuzz_dep_.CoverTab[187461]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:194
	if s == "" {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:194
		_go_fuzz_dep_.CoverTab[187467]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:195
		// _ = "end of CoverTab[187467]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:196
		_go_fuzz_dep_.CoverTab[187468]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:196
		// _ = "end of CoverTab[187468]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:196
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:196
	// _ = "end of CoverTab[187461]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:196
	_go_fuzz_dep_.CoverTab[187462]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:199
	if s[0] == '-' {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:199
		_go_fuzz_dep_.CoverTab[187469]++
												s = s[1:]
												if s == "" {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:201
			_go_fuzz_dep_.CoverTab[187470]++
													return false
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:202
			// _ = "end of CoverTab[187470]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:203
			_go_fuzz_dep_.CoverTab[187471]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:203
			// _ = "end of CoverTab[187471]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:203
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:203
		// _ = "end of CoverTab[187469]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:204
		_go_fuzz_dep_.CoverTab[187472]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:204
		// _ = "end of CoverTab[187472]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:204
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:204
	// _ = "end of CoverTab[187462]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:204
	_go_fuzz_dep_.CoverTab[187463]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:207
	switch {
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:208
		_go_fuzz_dep_.CoverTab[187473]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:209
		// _ = "end of CoverTab[187473]"

	case s[0] == '0':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:211
		_go_fuzz_dep_.CoverTab[187474]++
												s = s[1:]
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:212
		// _ = "end of CoverTab[187474]"

	case '1' <= s[0] && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:214
		_go_fuzz_dep_.CoverTab[187476]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:214
		return s[0] <= '9'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:214
		// _ = "end of CoverTab[187476]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:214
	}():
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:214
		_go_fuzz_dep_.CoverTab[187475]++
												s = s[1:]
												for len(s) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:216
			_go_fuzz_dep_.CoverTab[187477]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:216
			return '0' <= s[0]
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:216
			// _ = "end of CoverTab[187477]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:216
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:216
			_go_fuzz_dep_.CoverTab[187478]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:216
			return s[0] <= '9'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:216
			// _ = "end of CoverTab[187478]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:216
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:216
			_go_fuzz_dep_.CoverTab[187479]++
													s = s[1:]
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:217
			// _ = "end of CoverTab[187479]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:218
		// _ = "end of CoverTab[187475]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:219
	// _ = "end of CoverTab[187463]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:219
	_go_fuzz_dep_.CoverTab[187464]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:222
	if len(s) >= 2 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:222
		_go_fuzz_dep_.CoverTab[187480]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:222
		return s[0] == '.'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:222
		// _ = "end of CoverTab[187480]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:222
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:222
		_go_fuzz_dep_.CoverTab[187481]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:222
		return '0' <= s[1]
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:222
		// _ = "end of CoverTab[187481]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:222
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:222
		_go_fuzz_dep_.CoverTab[187482]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:222
		return s[1] <= '9'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:222
		// _ = "end of CoverTab[187482]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:222
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:222
		_go_fuzz_dep_.CoverTab[187483]++
												s = s[2:]
												for len(s) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:224
			_go_fuzz_dep_.CoverTab[187484]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:224
			return '0' <= s[0]
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:224
			// _ = "end of CoverTab[187484]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:224
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:224
			_go_fuzz_dep_.CoverTab[187485]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:224
			return s[0] <= '9'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:224
			// _ = "end of CoverTab[187485]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:224
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:224
			_go_fuzz_dep_.CoverTab[187486]++
													s = s[1:]
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:225
			// _ = "end of CoverTab[187486]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:226
		// _ = "end of CoverTab[187483]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:227
		_go_fuzz_dep_.CoverTab[187487]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:227
		// _ = "end of CoverTab[187487]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:227
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:227
	// _ = "end of CoverTab[187464]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:227
	_go_fuzz_dep_.CoverTab[187465]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:231
	if len(s) >= 2 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:231
		_go_fuzz_dep_.CoverTab[187488]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:231
		return (s[0] == 'e' || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:231
			_go_fuzz_dep_.CoverTab[187489]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:231
			return s[0] == 'E'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:231
			// _ = "end of CoverTab[187489]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:231
		}())
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:231
		// _ = "end of CoverTab[187488]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:231
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:231
		_go_fuzz_dep_.CoverTab[187490]++
												s = s[1:]
												if s[0] == '+' || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:233
			_go_fuzz_dep_.CoverTab[187492]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:233
			return s[0] == '-'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:233
			// _ = "end of CoverTab[187492]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:233
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:233
			_go_fuzz_dep_.CoverTab[187493]++
													s = s[1:]
													if s == "" {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:235
				_go_fuzz_dep_.CoverTab[187494]++
														return false
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:236
				// _ = "end of CoverTab[187494]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:237
				_go_fuzz_dep_.CoverTab[187495]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:237
				// _ = "end of CoverTab[187495]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:237
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:237
			// _ = "end of CoverTab[187493]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:238
			_go_fuzz_dep_.CoverTab[187496]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:238
			// _ = "end of CoverTab[187496]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:238
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:238
		// _ = "end of CoverTab[187490]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:238
		_go_fuzz_dep_.CoverTab[187491]++
												for len(s) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:239
			_go_fuzz_dep_.CoverTab[187497]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:239
			return '0' <= s[0]
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:239
			// _ = "end of CoverTab[187497]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:239
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:239
			_go_fuzz_dep_.CoverTab[187498]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:239
			return s[0] <= '9'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:239
			// _ = "end of CoverTab[187498]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:239
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:239
			_go_fuzz_dep_.CoverTab[187499]++
													s = s[1:]
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:240
			// _ = "end of CoverTab[187499]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:241
		// _ = "end of CoverTab[187491]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:242
		_go_fuzz_dep_.CoverTab[187500]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:242
		// _ = "end of CoverTab[187500]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:242
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:242
	// _ = "end of CoverTab[187465]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:242
	_go_fuzz_dep_.CoverTab[187466]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:245
	return s == ""
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:245
	// _ = "end of CoverTab[187466]"
}

// decodeState represents the state while decoding a JSON value.
type decodeState struct {
	data		[]byte
	off		int	// read offset in data
	scan		scanner
	nextscan	scanner	// for calls to nextValue
	savedError	error
	useNumber	bool
}

// errPhase is used for errors that should not happen unless
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:258
// there is a bug in the JSON decoder or something is editing
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:258
// the data slice while the decoder executes.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:261
var errPhase = errors.New("JSON decoder out of sync - data changing underfoot?")

func (d *decodeState) init(data []byte) *decodeState {
	d.data = data
	d.off = 0
	d.savedError = nil
	return d
}

// error aborts the decoding by panicking with err.
func (d *decodeState) error(err error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:271
	_go_fuzz_dep_.CoverTab[187501]++
											panic(err)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:272
	// _ = "end of CoverTab[187501]"
}

// saveError saves the first err it is called with,
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:275
// for reporting at the end of the unmarshal.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:277
func (d *decodeState) saveError(err error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:277
	_go_fuzz_dep_.CoverTab[187502]++
											if d.savedError == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:278
		_go_fuzz_dep_.CoverTab[187503]++
												d.savedError = err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:279
		// _ = "end of CoverTab[187503]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:280
		_go_fuzz_dep_.CoverTab[187504]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:280
		// _ = "end of CoverTab[187504]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:280
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:280
	// _ = "end of CoverTab[187502]"
}

// next cuts off and returns the next full JSON value in d.data[d.off:].
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:283
// The next value is known to be an object or array, not a literal.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:285
func (d *decodeState) next() []byte {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:285
	_go_fuzz_dep_.CoverTab[187505]++
											c := d.data[d.off]
											item, rest, err := nextValue(d.data[d.off:], &d.nextscan)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:288
		_go_fuzz_dep_.CoverTab[187508]++
												d.error(err)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:289
		// _ = "end of CoverTab[187508]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:290
		_go_fuzz_dep_.CoverTab[187509]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:290
		// _ = "end of CoverTab[187509]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:290
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:290
	// _ = "end of CoverTab[187505]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:290
	_go_fuzz_dep_.CoverTab[187506]++
											d.off = len(d.data) - len(rest)

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:296
	if c == '{' {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:296
		_go_fuzz_dep_.CoverTab[187510]++
												d.scan.step(&d.scan, '}')
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:297
		// _ = "end of CoverTab[187510]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:298
		_go_fuzz_dep_.CoverTab[187511]++
												d.scan.step(&d.scan, ']')
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:299
		// _ = "end of CoverTab[187511]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:300
	// _ = "end of CoverTab[187506]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:300
	_go_fuzz_dep_.CoverTab[187507]++

											return item
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:302
	// _ = "end of CoverTab[187507]"
}

// scanWhile processes bytes in d.data[d.off:] until it
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:305
// receives a scan code not equal to op.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:305
// It updates d.off and returns the new scan code.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:308
func (d *decodeState) scanWhile(op int) int {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:308
	_go_fuzz_dep_.CoverTab[187512]++
											var newOp int
											for {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:310
		_go_fuzz_dep_.CoverTab[187514]++
												if d.off >= len(d.data) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:311
			_go_fuzz_dep_.CoverTab[187516]++
													newOp = d.scan.eof()
													d.off = len(d.data) + 1
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:313
			// _ = "end of CoverTab[187516]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:314
			_go_fuzz_dep_.CoverTab[187517]++
													c := d.data[d.off]
													d.off++
													newOp = d.scan.step(&d.scan, c)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:317
			// _ = "end of CoverTab[187517]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:318
		// _ = "end of CoverTab[187514]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:318
		_go_fuzz_dep_.CoverTab[187515]++
												if newOp != op {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:319
			_go_fuzz_dep_.CoverTab[187518]++
													break
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:320
			// _ = "end of CoverTab[187518]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:321
			_go_fuzz_dep_.CoverTab[187519]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:321
			// _ = "end of CoverTab[187519]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:321
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:321
		// _ = "end of CoverTab[187515]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:322
	// _ = "end of CoverTab[187512]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:322
	_go_fuzz_dep_.CoverTab[187513]++
											return newOp
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:323
	// _ = "end of CoverTab[187513]"
}

// value decodes a JSON value from d.data[d.off:] into the value.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:326
// it updates d.off to point past the decoded value.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:328
func (d *decodeState) value(v reflect.Value) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:328
	_go_fuzz_dep_.CoverTab[187520]++
											if !v.IsValid() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:329
		_go_fuzz_dep_.CoverTab[187522]++
												_, rest, err := nextValue(d.data[d.off:], &d.nextscan)
												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:331
			_go_fuzz_dep_.CoverTab[187526]++
													d.error(err)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:332
			// _ = "end of CoverTab[187526]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:333
			_go_fuzz_dep_.CoverTab[187527]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:333
			// _ = "end of CoverTab[187527]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:333
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:333
		// _ = "end of CoverTab[187522]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:333
		_go_fuzz_dep_.CoverTab[187523]++
												d.off = len(d.data) - len(rest)

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:339
		if d.scan.redo {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:339
			_go_fuzz_dep_.CoverTab[187528]++

													d.scan.redo = false
													d.scan.step = stateBeginValue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:342
			// _ = "end of CoverTab[187528]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:343
			_go_fuzz_dep_.CoverTab[187529]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:343
			// _ = "end of CoverTab[187529]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:343
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:343
		// _ = "end of CoverTab[187523]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:343
		_go_fuzz_dep_.CoverTab[187524]++
												d.scan.step(&d.scan, '"')
												d.scan.step(&d.scan, '"')

												n := len(d.scan.parseState)
												if n > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:348
			_go_fuzz_dep_.CoverTab[187530]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:348
			return d.scan.parseState[n-1] == parseObjectKey
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:348
			// _ = "end of CoverTab[187530]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:348
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:348
			_go_fuzz_dep_.CoverTab[187531]++

													d.scan.step(&d.scan, ':')
													d.scan.step(&d.scan, '"')
													d.scan.step(&d.scan, '"')
													d.scan.step(&d.scan, '}')
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:353
			// _ = "end of CoverTab[187531]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:354
			_go_fuzz_dep_.CoverTab[187532]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:354
			// _ = "end of CoverTab[187532]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:354
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:354
		// _ = "end of CoverTab[187524]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:354
		_go_fuzz_dep_.CoverTab[187525]++

												return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:356
		// _ = "end of CoverTab[187525]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:357
		_go_fuzz_dep_.CoverTab[187533]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:357
		// _ = "end of CoverTab[187533]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:357
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:357
	// _ = "end of CoverTab[187520]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:357
	_go_fuzz_dep_.CoverTab[187521]++

											switch op := d.scanWhile(scanSkipSpace); op {
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:360
		_go_fuzz_dep_.CoverTab[187534]++
												d.error(errPhase)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:361
		// _ = "end of CoverTab[187534]"

	case scanBeginArray:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:363
		_go_fuzz_dep_.CoverTab[187535]++
												d.array(v)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:364
		// _ = "end of CoverTab[187535]"

	case scanBeginObject:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:366
		_go_fuzz_dep_.CoverTab[187536]++
												d.object(v)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:367
		// _ = "end of CoverTab[187536]"

	case scanBeginLiteral:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:369
		_go_fuzz_dep_.CoverTab[187537]++
												d.literal(v)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:370
		// _ = "end of CoverTab[187537]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:371
	// _ = "end of CoverTab[187521]"
}

type unquotedValue struct{}

// valueQuoted is like value but decodes a
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:376
// quoted string literal or literal null into an interface value.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:376
// If it finds anything other than a quoted string literal or null,
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:376
// valueQuoted returns unquotedValue{}.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:380
func (d *decodeState) valueQuoted() interface{} {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:380
	_go_fuzz_dep_.CoverTab[187538]++
											switch op := d.scanWhile(scanSkipSpace); op {
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:382
		_go_fuzz_dep_.CoverTab[187540]++
												d.error(errPhase)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:383
		// _ = "end of CoverTab[187540]"

	case scanBeginArray:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:385
		_go_fuzz_dep_.CoverTab[187541]++
												d.array(reflect.Value{})
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:386
		// _ = "end of CoverTab[187541]"

	case scanBeginObject:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:388
		_go_fuzz_dep_.CoverTab[187542]++
												d.object(reflect.Value{})
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:389
		// _ = "end of CoverTab[187542]"

	case scanBeginLiteral:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:391
		_go_fuzz_dep_.CoverTab[187543]++
												switch v := d.literalInterface().(type) {
		case nil, string:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:393
			_go_fuzz_dep_.CoverTab[187544]++
													return v
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:394
			// _ = "end of CoverTab[187544]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:395
		// _ = "end of CoverTab[187543]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:396
	// _ = "end of CoverTab[187538]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:396
	_go_fuzz_dep_.CoverTab[187539]++
											return unquotedValue{}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:397
	// _ = "end of CoverTab[187539]"
}

// indirect walks down v allocating pointers as needed,
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:400
// until it gets to a non-pointer.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:400
// if it encounters an Unmarshaler, indirect stops and returns that.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:400
// if decodingNull is true, indirect stops at the last pointer so it can be set to nil.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:404
func (d *decodeState) indirect(v reflect.Value, decodingNull bool) (Unmarshaler, encoding.TextUnmarshaler, reflect.Value) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:404
	_go_fuzz_dep_.CoverTab[187545]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:408
	if v.Kind() != reflect.Ptr && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:408
		_go_fuzz_dep_.CoverTab[187548]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:408
		return v.Type().Name() != ""
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:408
		// _ = "end of CoverTab[187548]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:408
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:408
		_go_fuzz_dep_.CoverTab[187549]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:408
		return v.CanAddr()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:408
		// _ = "end of CoverTab[187549]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:408
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:408
		_go_fuzz_dep_.CoverTab[187550]++
												v = v.Addr()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:409
		// _ = "end of CoverTab[187550]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:410
		_go_fuzz_dep_.CoverTab[187551]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:410
		// _ = "end of CoverTab[187551]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:410
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:410
	// _ = "end of CoverTab[187545]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:410
	_go_fuzz_dep_.CoverTab[187546]++
											for {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:411
		_go_fuzz_dep_.CoverTab[187552]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:414
		if v.Kind() == reflect.Interface && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:414
			_go_fuzz_dep_.CoverTab[187558]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:414
			return !v.IsNil()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:414
			// _ = "end of CoverTab[187558]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:414
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:414
			_go_fuzz_dep_.CoverTab[187559]++
													e := v.Elem()
													if e.Kind() == reflect.Ptr && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:416
				_go_fuzz_dep_.CoverTab[187560]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:416
				return !e.IsNil()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:416
				// _ = "end of CoverTab[187560]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:416
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:416
				_go_fuzz_dep_.CoverTab[187561]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:416
				return (!decodingNull || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:416
					_go_fuzz_dep_.CoverTab[187562]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:416
					return e.Elem().Kind() == reflect.Ptr
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:416
					// _ = "end of CoverTab[187562]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:416
				}())
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:416
				// _ = "end of CoverTab[187561]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:416
			}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:416
				_go_fuzz_dep_.CoverTab[187563]++
														v = e
														continue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:418
				// _ = "end of CoverTab[187563]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:419
				_go_fuzz_dep_.CoverTab[187564]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:419
				// _ = "end of CoverTab[187564]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:419
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:419
			// _ = "end of CoverTab[187559]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:420
			_go_fuzz_dep_.CoverTab[187565]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:420
			// _ = "end of CoverTab[187565]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:420
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:420
		// _ = "end of CoverTab[187552]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:420
		_go_fuzz_dep_.CoverTab[187553]++

												if v.Kind() != reflect.Ptr {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:422
			_go_fuzz_dep_.CoverTab[187566]++
													break
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:423
			// _ = "end of CoverTab[187566]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:424
			_go_fuzz_dep_.CoverTab[187567]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:424
			// _ = "end of CoverTab[187567]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:424
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:424
		// _ = "end of CoverTab[187553]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:424
		_go_fuzz_dep_.CoverTab[187554]++

												if v.Elem().Kind() != reflect.Ptr && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:426
			_go_fuzz_dep_.CoverTab[187568]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:426
			return decodingNull
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:426
			// _ = "end of CoverTab[187568]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:426
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:426
			_go_fuzz_dep_.CoverTab[187569]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:426
			return v.CanSet()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:426
			// _ = "end of CoverTab[187569]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:426
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:426
			_go_fuzz_dep_.CoverTab[187570]++
													break
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:427
			// _ = "end of CoverTab[187570]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:428
			_go_fuzz_dep_.CoverTab[187571]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:428
			// _ = "end of CoverTab[187571]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:428
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:428
		// _ = "end of CoverTab[187554]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:428
		_go_fuzz_dep_.CoverTab[187555]++
												if v.IsNil() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:429
			_go_fuzz_dep_.CoverTab[187572]++
													v.Set(reflect.New(v.Type().Elem()))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:430
			// _ = "end of CoverTab[187572]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:431
			_go_fuzz_dep_.CoverTab[187573]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:431
			// _ = "end of CoverTab[187573]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:431
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:431
		// _ = "end of CoverTab[187555]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:431
		_go_fuzz_dep_.CoverTab[187556]++
												if v.Type().NumMethod() > 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:432
			_go_fuzz_dep_.CoverTab[187574]++
													if u, ok := v.Interface().(Unmarshaler); ok {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:433
				_go_fuzz_dep_.CoverTab[187576]++
														return u, nil, reflect.Value{}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:434
				// _ = "end of CoverTab[187576]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:435
				_go_fuzz_dep_.CoverTab[187577]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:435
				// _ = "end of CoverTab[187577]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:435
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:435
			// _ = "end of CoverTab[187574]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:435
			_go_fuzz_dep_.CoverTab[187575]++
													if u, ok := v.Interface().(encoding.TextUnmarshaler); ok {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:436
				_go_fuzz_dep_.CoverTab[187578]++
														return nil, u, reflect.Value{}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:437
				// _ = "end of CoverTab[187578]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:438
				_go_fuzz_dep_.CoverTab[187579]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:438
				// _ = "end of CoverTab[187579]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:438
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:438
			// _ = "end of CoverTab[187575]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:439
			_go_fuzz_dep_.CoverTab[187580]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:439
			// _ = "end of CoverTab[187580]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:439
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:439
		// _ = "end of CoverTab[187556]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:439
		_go_fuzz_dep_.CoverTab[187557]++
												v = v.Elem()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:440
		// _ = "end of CoverTab[187557]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:441
	// _ = "end of CoverTab[187546]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:441
	_go_fuzz_dep_.CoverTab[187547]++
											return nil, nil, v
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:442
	// _ = "end of CoverTab[187547]"
}

// array consumes an array from d.data[d.off-1:], decoding into the value v.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:445
// the first byte of the array ('[') has been read already.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:447
func (d *decodeState) array(v reflect.Value) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:447
	_go_fuzz_dep_.CoverTab[187581]++

											u, ut, pv := d.indirect(v, false)
											if u != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:450
		_go_fuzz_dep_.CoverTab[187587]++
												d.off--
												err := u.UnmarshalJSON(d.next())
												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:453
			_go_fuzz_dep_.CoverTab[187589]++
													d.error(err)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:454
			// _ = "end of CoverTab[187589]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:455
			_go_fuzz_dep_.CoverTab[187590]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:455
			// _ = "end of CoverTab[187590]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:455
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:455
		// _ = "end of CoverTab[187587]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:455
		_go_fuzz_dep_.CoverTab[187588]++
												return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:456
		// _ = "end of CoverTab[187588]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:457
		_go_fuzz_dep_.CoverTab[187591]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:457
		// _ = "end of CoverTab[187591]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:457
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:457
	// _ = "end of CoverTab[187581]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:457
	_go_fuzz_dep_.CoverTab[187582]++
											if ut != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:458
		_go_fuzz_dep_.CoverTab[187592]++
												d.saveError(&UnmarshalTypeError{"array", v.Type(), int64(d.off)})
												d.off--
												d.next()
												return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:462
		// _ = "end of CoverTab[187592]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:463
		_go_fuzz_dep_.CoverTab[187593]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:463
		// _ = "end of CoverTab[187593]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:463
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:463
	// _ = "end of CoverTab[187582]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:463
	_go_fuzz_dep_.CoverTab[187583]++

											v = pv

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:468
	switch v.Kind() {
	case reflect.Interface:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:469
		_go_fuzz_dep_.CoverTab[187594]++
												if v.NumMethod() == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:470
			_go_fuzz_dep_.CoverTab[187599]++

													v.Set(reflect.ValueOf(d.arrayInterface()))
													return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:473
			// _ = "end of CoverTab[187599]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:474
			_go_fuzz_dep_.CoverTab[187600]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:474
			// _ = "end of CoverTab[187600]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:474
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:474
		// _ = "end of CoverTab[187594]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:474
		_go_fuzz_dep_.CoverTab[187595]++

												fallthrough
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:476
		// _ = "end of CoverTab[187595]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:477
		_go_fuzz_dep_.CoverTab[187596]++
												d.saveError(&UnmarshalTypeError{"array", v.Type(), int64(d.off)})
												d.off--
												d.next()
												return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:481
		// _ = "end of CoverTab[187596]"
	case reflect.Array:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:482
		_go_fuzz_dep_.CoverTab[187597]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:482
		// _ = "end of CoverTab[187597]"
	case reflect.Slice:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:483
		_go_fuzz_dep_.CoverTab[187598]++
												break
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:484
		// _ = "end of CoverTab[187598]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:485
	// _ = "end of CoverTab[187583]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:485
	_go_fuzz_dep_.CoverTab[187584]++

											i := 0
											for {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:488
		_go_fuzz_dep_.CoverTab[187601]++

												op := d.scanWhile(scanSkipSpace)
												if op == scanEndArray {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:491
			_go_fuzz_dep_.CoverTab[187606]++
													break
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:492
			// _ = "end of CoverTab[187606]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:493
			_go_fuzz_dep_.CoverTab[187607]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:493
			// _ = "end of CoverTab[187607]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:493
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:493
		// _ = "end of CoverTab[187601]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:493
		_go_fuzz_dep_.CoverTab[187602]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:496
		d.off--
												d.scan.undo(op)

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:500
		if v.Kind() == reflect.Slice {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:500
			_go_fuzz_dep_.CoverTab[187608]++

													if i >= v.Cap() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:502
				_go_fuzz_dep_.CoverTab[187610]++
														newcap := v.Cap() + v.Cap()/2
														if newcap < 4 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:504
					_go_fuzz_dep_.CoverTab[187612]++
															newcap = 4
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:505
					// _ = "end of CoverTab[187612]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:506
					_go_fuzz_dep_.CoverTab[187613]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:506
					// _ = "end of CoverTab[187613]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:506
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:506
				// _ = "end of CoverTab[187610]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:506
				_go_fuzz_dep_.CoverTab[187611]++
														newv := reflect.MakeSlice(v.Type(), v.Len(), newcap)
														reflect.Copy(newv, v)
														v.Set(newv)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:509
				// _ = "end of CoverTab[187611]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:510
				_go_fuzz_dep_.CoverTab[187614]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:510
				// _ = "end of CoverTab[187614]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:510
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:510
			// _ = "end of CoverTab[187608]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:510
			_go_fuzz_dep_.CoverTab[187609]++
													if i >= v.Len() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:511
				_go_fuzz_dep_.CoverTab[187615]++
														v.SetLen(i + 1)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:512
				// _ = "end of CoverTab[187615]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:513
				_go_fuzz_dep_.CoverTab[187616]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:513
				// _ = "end of CoverTab[187616]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:513
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:513
			// _ = "end of CoverTab[187609]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:514
			_go_fuzz_dep_.CoverTab[187617]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:514
			// _ = "end of CoverTab[187617]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:514
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:514
		// _ = "end of CoverTab[187602]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:514
		_go_fuzz_dep_.CoverTab[187603]++

												if i < v.Len() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:516
			_go_fuzz_dep_.CoverTab[187618]++

													d.value(v.Index(i))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:518
			// _ = "end of CoverTab[187618]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:519
			_go_fuzz_dep_.CoverTab[187619]++

													d.value(reflect.Value{})
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:521
			// _ = "end of CoverTab[187619]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:522
		// _ = "end of CoverTab[187603]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:522
		_go_fuzz_dep_.CoverTab[187604]++
												i++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:526
		op = d.scanWhile(scanSkipSpace)
		if op == scanEndArray {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:527
			_go_fuzz_dep_.CoverTab[187620]++
													break
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:528
			// _ = "end of CoverTab[187620]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:529
			_go_fuzz_dep_.CoverTab[187621]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:529
			// _ = "end of CoverTab[187621]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:529
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:529
		// _ = "end of CoverTab[187604]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:529
		_go_fuzz_dep_.CoverTab[187605]++
												if op != scanArrayValue {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:530
			_go_fuzz_dep_.CoverTab[187622]++
													d.error(errPhase)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:531
			// _ = "end of CoverTab[187622]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:532
			_go_fuzz_dep_.CoverTab[187623]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:532
			// _ = "end of CoverTab[187623]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:532
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:532
		// _ = "end of CoverTab[187605]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:533
	// _ = "end of CoverTab[187584]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:533
	_go_fuzz_dep_.CoverTab[187585]++

											if i < v.Len() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:535
		_go_fuzz_dep_.CoverTab[187624]++
												if v.Kind() == reflect.Array {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:536
			_go_fuzz_dep_.CoverTab[187625]++

													z := reflect.Zero(v.Type().Elem())
													for ; i < v.Len(); i++ {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:539
				_go_fuzz_dep_.CoverTab[187626]++
														v.Index(i).Set(z)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:540
				// _ = "end of CoverTab[187626]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:541
			// _ = "end of CoverTab[187625]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:542
			_go_fuzz_dep_.CoverTab[187627]++
													v.SetLen(i)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:543
			// _ = "end of CoverTab[187627]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:544
		// _ = "end of CoverTab[187624]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:545
		_go_fuzz_dep_.CoverTab[187628]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:545
		// _ = "end of CoverTab[187628]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:545
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:545
	// _ = "end of CoverTab[187585]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:545
	_go_fuzz_dep_.CoverTab[187586]++
											if i == 0 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:546
		_go_fuzz_dep_.CoverTab[187629]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:546
		return v.Kind() == reflect.Slice
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:546
		// _ = "end of CoverTab[187629]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:546
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:546
		_go_fuzz_dep_.CoverTab[187630]++
												v.Set(reflect.MakeSlice(v.Type(), 0, 0))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:547
		// _ = "end of CoverTab[187630]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:548
		_go_fuzz_dep_.CoverTab[187631]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:548
		// _ = "end of CoverTab[187631]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:548
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:548
	// _ = "end of CoverTab[187586]"
}

var nullLiteral = []byte("null")

// object consumes an object from d.data[d.off-1:], decoding into the value v.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:553
// the first byte ('{') of the object has been read already.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:555
func (d *decodeState) object(v reflect.Value) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:555
	_go_fuzz_dep_.CoverTab[187632]++

											u, ut, pv := d.indirect(v, false)
											if u != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:558
		_go_fuzz_dep_.CoverTab[187637]++
												d.off--
												err := u.UnmarshalJSON(d.next())
												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:561
			_go_fuzz_dep_.CoverTab[187639]++
													d.error(err)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:562
			// _ = "end of CoverTab[187639]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:563
			_go_fuzz_dep_.CoverTab[187640]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:563
			// _ = "end of CoverTab[187640]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:563
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:563
		// _ = "end of CoverTab[187637]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:563
		_go_fuzz_dep_.CoverTab[187638]++
												return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:564
		// _ = "end of CoverTab[187638]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:565
		_go_fuzz_dep_.CoverTab[187641]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:565
		// _ = "end of CoverTab[187641]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:565
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:565
	// _ = "end of CoverTab[187632]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:565
	_go_fuzz_dep_.CoverTab[187633]++
											if ut != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:566
		_go_fuzz_dep_.CoverTab[187642]++
												d.saveError(&UnmarshalTypeError{"object", v.Type(), int64(d.off)})
												d.off--
												d.next()
												return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:570
		// _ = "end of CoverTab[187642]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:571
		_go_fuzz_dep_.CoverTab[187643]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:571
		// _ = "end of CoverTab[187643]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:571
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:571
	// _ = "end of CoverTab[187633]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:571
	_go_fuzz_dep_.CoverTab[187634]++
											v = pv

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:575
	if v.Kind() == reflect.Interface && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:575
		_go_fuzz_dep_.CoverTab[187644]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:575
		return v.NumMethod() == 0
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:575
		// _ = "end of CoverTab[187644]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:575
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:575
		_go_fuzz_dep_.CoverTab[187645]++
												v.Set(reflect.ValueOf(d.objectInterface()))
												return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:577
		// _ = "end of CoverTab[187645]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:578
		_go_fuzz_dep_.CoverTab[187646]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:578
		// _ = "end of CoverTab[187646]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:578
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:578
	// _ = "end of CoverTab[187634]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:578
	_go_fuzz_dep_.CoverTab[187635]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:581
	switch v.Kind() {
	case reflect.Map:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:582
		_go_fuzz_dep_.CoverTab[187647]++

												t := v.Type()
												if t.Key().Kind() != reflect.String {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:585
			_go_fuzz_dep_.CoverTab[187651]++
													d.saveError(&UnmarshalTypeError{"object", v.Type(), int64(d.off)})
													d.off--
													d.next()
													return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:589
			// _ = "end of CoverTab[187651]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:590
			_go_fuzz_dep_.CoverTab[187652]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:590
			// _ = "end of CoverTab[187652]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:590
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:590
		// _ = "end of CoverTab[187647]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:590
		_go_fuzz_dep_.CoverTab[187648]++
												if v.IsNil() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:591
			_go_fuzz_dep_.CoverTab[187653]++
													v.Set(reflect.MakeMap(t))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:592
			// _ = "end of CoverTab[187653]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:593
			_go_fuzz_dep_.CoverTab[187654]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:593
			// _ = "end of CoverTab[187654]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:593
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:593
		// _ = "end of CoverTab[187648]"
	case reflect.Struct:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:594
		_go_fuzz_dep_.CoverTab[187649]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:594
		// _ = "end of CoverTab[187649]"

	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:596
		_go_fuzz_dep_.CoverTab[187650]++
												d.saveError(&UnmarshalTypeError{"object", v.Type(), int64(d.off)})
												d.off--
												d.next()
												return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:600
		// _ = "end of CoverTab[187650]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:601
	// _ = "end of CoverTab[187635]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:601
	_go_fuzz_dep_.CoverTab[187636]++

											var mapElem reflect.Value
											keys := map[string]bool{}

											for {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:606
		_go_fuzz_dep_.CoverTab[187655]++

												op := d.scanWhile(scanSkipSpace)
												if op == scanEndObject {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:609
			_go_fuzz_dep_.CoverTab[187666]++

													break
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:611
			// _ = "end of CoverTab[187666]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:612
			_go_fuzz_dep_.CoverTab[187667]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:612
			// _ = "end of CoverTab[187667]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:612
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:612
		// _ = "end of CoverTab[187655]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:612
		_go_fuzz_dep_.CoverTab[187656]++
												if op != scanBeginLiteral {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:613
			_go_fuzz_dep_.CoverTab[187668]++
													d.error(errPhase)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:614
			// _ = "end of CoverTab[187668]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:615
			_go_fuzz_dep_.CoverTab[187669]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:615
			// _ = "end of CoverTab[187669]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:615
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:615
		// _ = "end of CoverTab[187656]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:615
		_go_fuzz_dep_.CoverTab[187657]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:618
		start := d.off - 1
		op = d.scanWhile(scanContinue)
		item := d.data[start : d.off-1]
		key, ok := unquote(item)
		if !ok {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:622
			_go_fuzz_dep_.CoverTab[187670]++
													d.error(errPhase)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:623
			// _ = "end of CoverTab[187670]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:624
			_go_fuzz_dep_.CoverTab[187671]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:624
			// _ = "end of CoverTab[187671]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:624
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:624
		// _ = "end of CoverTab[187657]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:624
		_go_fuzz_dep_.CoverTab[187658]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:627
		_, ok = keys[key]
		if !ok {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:628
			_go_fuzz_dep_.CoverTab[187672]++
													keys[key] = true
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:629
			// _ = "end of CoverTab[187672]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:630
			_go_fuzz_dep_.CoverTab[187673]++
													d.error(fmt.Errorf("json: duplicate key '%s' in object", key))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:631
			// _ = "end of CoverTab[187673]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:632
		// _ = "end of CoverTab[187658]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:632
		_go_fuzz_dep_.CoverTab[187659]++

		// Figure out field corresponding to key.
		var subv reflect.Value
		destring := false

		if v.Kind() == reflect.Map {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:638
			_go_fuzz_dep_.CoverTab[187674]++
													elemType := v.Type().Elem()
													if !mapElem.IsValid() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:640
				_go_fuzz_dep_.CoverTab[187676]++
														mapElem = reflect.New(elemType).Elem()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:641
				// _ = "end of CoverTab[187676]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:642
				_go_fuzz_dep_.CoverTab[187677]++
														mapElem.Set(reflect.Zero(elemType))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:643
				// _ = "end of CoverTab[187677]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:644
			// _ = "end of CoverTab[187674]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:644
			_go_fuzz_dep_.CoverTab[187675]++
													subv = mapElem
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:645
			// _ = "end of CoverTab[187675]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:646
			_go_fuzz_dep_.CoverTab[187678]++
													var f *field
													fields := cachedTypeFields(v.Type())
													for i := range fields {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:649
				_go_fuzz_dep_.CoverTab[187680]++
														ff := &fields[i]
														if bytes.Equal(ff.nameBytes, []byte(key)) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:651
					_go_fuzz_dep_.CoverTab[187681]++
															f = ff
															break
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:653
					// _ = "end of CoverTab[187681]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:654
					_go_fuzz_dep_.CoverTab[187682]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:654
					// _ = "end of CoverTab[187682]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:654
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:654
				// _ = "end of CoverTab[187680]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:655
			// _ = "end of CoverTab[187678]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:655
			_go_fuzz_dep_.CoverTab[187679]++
													if f != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:656
				_go_fuzz_dep_.CoverTab[187683]++
														subv = v
														destring = f.quoted
														for _, i := range f.index {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:659
					_go_fuzz_dep_.CoverTab[187684]++
															if subv.Kind() == reflect.Ptr {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:660
						_go_fuzz_dep_.CoverTab[187686]++
																if subv.IsNil() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:661
							_go_fuzz_dep_.CoverTab[187688]++
																	subv.Set(reflect.New(subv.Type().Elem()))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:662
							// _ = "end of CoverTab[187688]"
						} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:663
							_go_fuzz_dep_.CoverTab[187689]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:663
							// _ = "end of CoverTab[187689]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:663
						}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:663
						// _ = "end of CoverTab[187686]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:663
						_go_fuzz_dep_.CoverTab[187687]++
																subv = subv.Elem()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:664
						// _ = "end of CoverTab[187687]"
					} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:665
						_go_fuzz_dep_.CoverTab[187690]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:665
						// _ = "end of CoverTab[187690]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:665
					}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:665
					// _ = "end of CoverTab[187684]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:665
					_go_fuzz_dep_.CoverTab[187685]++
															subv = subv.Field(i)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:666
					// _ = "end of CoverTab[187685]"
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:667
				// _ = "end of CoverTab[187683]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:668
				_go_fuzz_dep_.CoverTab[187691]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:668
				// _ = "end of CoverTab[187691]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:668
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:668
			// _ = "end of CoverTab[187679]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:669
		// _ = "end of CoverTab[187659]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:669
		_go_fuzz_dep_.CoverTab[187660]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:672
		if op == scanSkipSpace {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:672
			_go_fuzz_dep_.CoverTab[187692]++
													op = d.scanWhile(scanSkipSpace)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:673
			// _ = "end of CoverTab[187692]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:674
			_go_fuzz_dep_.CoverTab[187693]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:674
			// _ = "end of CoverTab[187693]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:674
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:674
		// _ = "end of CoverTab[187660]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:674
		_go_fuzz_dep_.CoverTab[187661]++
												if op != scanObjectKey {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:675
			_go_fuzz_dep_.CoverTab[187694]++
													d.error(errPhase)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:676
			// _ = "end of CoverTab[187694]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:677
			_go_fuzz_dep_.CoverTab[187695]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:677
			// _ = "end of CoverTab[187695]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:677
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:677
		// _ = "end of CoverTab[187661]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:677
		_go_fuzz_dep_.CoverTab[187662]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:680
		if destring {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:680
			_go_fuzz_dep_.CoverTab[187696]++
													switch qv := d.valueQuoted().(type) {
			case nil:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:682
				_go_fuzz_dep_.CoverTab[187697]++
														d.literalStore(nullLiteral, subv, false)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:683
				// _ = "end of CoverTab[187697]"
			case string:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:684
				_go_fuzz_dep_.CoverTab[187698]++
														d.literalStore([]byte(qv), subv, true)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:685
				// _ = "end of CoverTab[187698]"
			default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:686
				_go_fuzz_dep_.CoverTab[187699]++
														d.saveError(fmt.Errorf("json: invalid use of ,string struct tag, trying to unmarshal unquoted value into %v", subv.Type()))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:687
				// _ = "end of CoverTab[187699]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:688
			// _ = "end of CoverTab[187696]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:689
			_go_fuzz_dep_.CoverTab[187700]++
													d.value(subv)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:690
			// _ = "end of CoverTab[187700]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:691
		// _ = "end of CoverTab[187662]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:691
		_go_fuzz_dep_.CoverTab[187663]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:695
		if v.Kind() == reflect.Map {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:695
			_go_fuzz_dep_.CoverTab[187701]++
													kv := reflect.ValueOf(key).Convert(v.Type().Key())
													v.SetMapIndex(kv, subv)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:697
			// _ = "end of CoverTab[187701]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:698
			_go_fuzz_dep_.CoverTab[187702]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:698
			// _ = "end of CoverTab[187702]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:698
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:698
		// _ = "end of CoverTab[187663]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:698
		_go_fuzz_dep_.CoverTab[187664]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:701
		op = d.scanWhile(scanSkipSpace)
		if op == scanEndObject {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:702
			_go_fuzz_dep_.CoverTab[187703]++
													break
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:703
			// _ = "end of CoverTab[187703]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:704
			_go_fuzz_dep_.CoverTab[187704]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:704
			// _ = "end of CoverTab[187704]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:704
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:704
		// _ = "end of CoverTab[187664]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:704
		_go_fuzz_dep_.CoverTab[187665]++
												if op != scanObjectValue {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:705
			_go_fuzz_dep_.CoverTab[187705]++
													d.error(errPhase)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:706
			// _ = "end of CoverTab[187705]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:707
			_go_fuzz_dep_.CoverTab[187706]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:707
			// _ = "end of CoverTab[187706]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:707
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:707
		// _ = "end of CoverTab[187665]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:708
	// _ = "end of CoverTab[187636]"
}

// literal consumes a literal from d.data[d.off-1:], decoding into the value v.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:711
// The first byte of the literal has been read already
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:711
// (that's how the caller knows it's a literal).
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:714
func (d *decodeState) literal(v reflect.Value) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:714
	_go_fuzz_dep_.CoverTab[187707]++

											start := d.off - 1
											op := d.scanWhile(scanContinue)

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:720
	d.off--
											d.scan.undo(op)

											d.literalStore(d.data[start:d.off], v, false)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:723
	// _ = "end of CoverTab[187707]"
}

// convertNumber converts the number literal s to a float64 or a Number
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:726
// depending on the setting of d.useNumber.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:728
func (d *decodeState) convertNumber(s string) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:728
	_go_fuzz_dep_.CoverTab[187708]++
											if d.useNumber {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:729
		_go_fuzz_dep_.CoverTab[187711]++
												return Number(s), nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:730
		// _ = "end of CoverTab[187711]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:731
		_go_fuzz_dep_.CoverTab[187712]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:731
		// _ = "end of CoverTab[187712]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:731
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:731
	// _ = "end of CoverTab[187708]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:731
	_go_fuzz_dep_.CoverTab[187709]++
											f, err := strconv.ParseFloat(s, 64)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:733
		_go_fuzz_dep_.CoverTab[187713]++
												return nil, &UnmarshalTypeError{"number " + s, reflect.TypeOf(0.0), int64(d.off)}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:734
		// _ = "end of CoverTab[187713]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:735
		_go_fuzz_dep_.CoverTab[187714]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:735
		// _ = "end of CoverTab[187714]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:735
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:735
	// _ = "end of CoverTab[187709]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:735
	_go_fuzz_dep_.CoverTab[187710]++
											return f, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:736
	// _ = "end of CoverTab[187710]"
}

var numberType = reflect.TypeOf(Number(""))

// literalStore decodes a literal stored in item into v.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:741
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:741
// fromQuoted indicates whether this literal came from unwrapping a
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:741
// string from the ",string" struct tag option. this is used only to
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:741
// produce more helpful error messages.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:746
func (d *decodeState) literalStore(item []byte, v reflect.Value, fromQuoted bool) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:746
	_go_fuzz_dep_.CoverTab[187715]++

											if len(item) == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:748
		_go_fuzz_dep_.CoverTab[187719]++

												d.saveError(fmt.Errorf("json: invalid use of ,string struct tag, trying to unmarshal %q into %v", item, v.Type()))
												return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:751
		// _ = "end of CoverTab[187719]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:752
		_go_fuzz_dep_.CoverTab[187720]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:752
		// _ = "end of CoverTab[187720]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:752
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:752
	// _ = "end of CoverTab[187715]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:752
	_go_fuzz_dep_.CoverTab[187716]++
											wantptr := item[0] == 'n'
											u, ut, pv := d.indirect(v, wantptr)
											if u != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:755
		_go_fuzz_dep_.CoverTab[187721]++
												err := u.UnmarshalJSON(item)
												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:757
			_go_fuzz_dep_.CoverTab[187723]++
													d.error(err)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:758
			// _ = "end of CoverTab[187723]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:759
			_go_fuzz_dep_.CoverTab[187724]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:759
			// _ = "end of CoverTab[187724]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:759
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:759
		// _ = "end of CoverTab[187721]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:759
		_go_fuzz_dep_.CoverTab[187722]++
												return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:760
		// _ = "end of CoverTab[187722]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:761
		_go_fuzz_dep_.CoverTab[187725]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:761
		// _ = "end of CoverTab[187725]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:761
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:761
	// _ = "end of CoverTab[187716]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:761
	_go_fuzz_dep_.CoverTab[187717]++
											if ut != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:762
		_go_fuzz_dep_.CoverTab[187726]++
												if item[0] != '"' {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:763
			_go_fuzz_dep_.CoverTab[187730]++
													if fromQuoted {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:764
				_go_fuzz_dep_.CoverTab[187732]++
														d.saveError(fmt.Errorf("json: invalid use of ,string struct tag, trying to unmarshal %q into %v", item, v.Type()))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:765
				// _ = "end of CoverTab[187732]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:766
				_go_fuzz_dep_.CoverTab[187733]++
														d.saveError(&UnmarshalTypeError{"string", v.Type(), int64(d.off)})
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:767
				// _ = "end of CoverTab[187733]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:768
			// _ = "end of CoverTab[187730]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:768
			_go_fuzz_dep_.CoverTab[187731]++
													return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:769
			// _ = "end of CoverTab[187731]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:770
			_go_fuzz_dep_.CoverTab[187734]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:770
			// _ = "end of CoverTab[187734]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:770
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:770
		// _ = "end of CoverTab[187726]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:770
		_go_fuzz_dep_.CoverTab[187727]++
												s, ok := unquoteBytes(item)
												if !ok {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:772
			_go_fuzz_dep_.CoverTab[187735]++
													if fromQuoted {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:773
				_go_fuzz_dep_.CoverTab[187736]++
														d.error(fmt.Errorf("json: invalid use of ,string struct tag, trying to unmarshal %q into %v", item, v.Type()))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:774
				// _ = "end of CoverTab[187736]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:775
				_go_fuzz_dep_.CoverTab[187737]++
														d.error(errPhase)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:776
				// _ = "end of CoverTab[187737]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:777
			// _ = "end of CoverTab[187735]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:778
			_go_fuzz_dep_.CoverTab[187738]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:778
			// _ = "end of CoverTab[187738]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:778
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:778
		// _ = "end of CoverTab[187727]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:778
		_go_fuzz_dep_.CoverTab[187728]++
												err := ut.UnmarshalText(s)
												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:780
			_go_fuzz_dep_.CoverTab[187739]++
													d.error(err)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:781
			// _ = "end of CoverTab[187739]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:782
			_go_fuzz_dep_.CoverTab[187740]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:782
			// _ = "end of CoverTab[187740]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:782
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:782
		// _ = "end of CoverTab[187728]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:782
		_go_fuzz_dep_.CoverTab[187729]++
												return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:783
		// _ = "end of CoverTab[187729]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:784
		_go_fuzz_dep_.CoverTab[187741]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:784
		// _ = "end of CoverTab[187741]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:784
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:784
	// _ = "end of CoverTab[187717]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:784
	_go_fuzz_dep_.CoverTab[187718]++

											v = pv

											switch c := item[0]; c {
	case 'n':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:789
		_go_fuzz_dep_.CoverTab[187742]++
												switch v.Kind() {
		case reflect.Interface, reflect.Ptr, reflect.Map, reflect.Slice:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:791
			_go_fuzz_dep_.CoverTab[187748]++
													v.Set(reflect.Zero(v.Type()))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:792
			// _ = "end of CoverTab[187748]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:792
		default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:792
			_go_fuzz_dep_.CoverTab[187749]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:792
			// _ = "end of CoverTab[187749]"

		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:794
		// _ = "end of CoverTab[187742]"
	case 't', 'f':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:795
		_go_fuzz_dep_.CoverTab[187743]++
												value := c == 't'
												switch v.Kind() {
		default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:798
			_go_fuzz_dep_.CoverTab[187750]++
													if fromQuoted {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:799
				_go_fuzz_dep_.CoverTab[187753]++
														d.saveError(fmt.Errorf("json: invalid use of ,string struct tag, trying to unmarshal %q into %v", item, v.Type()))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:800
				// _ = "end of CoverTab[187753]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:801
				_go_fuzz_dep_.CoverTab[187754]++
														d.saveError(&UnmarshalTypeError{"bool", v.Type(), int64(d.off)})
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:802
				// _ = "end of CoverTab[187754]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:803
			// _ = "end of CoverTab[187750]"
		case reflect.Bool:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:804
			_go_fuzz_dep_.CoverTab[187751]++
													v.SetBool(value)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:805
			// _ = "end of CoverTab[187751]"
		case reflect.Interface:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:806
			_go_fuzz_dep_.CoverTab[187752]++
													if v.NumMethod() == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:807
				_go_fuzz_dep_.CoverTab[187755]++
														v.Set(reflect.ValueOf(value))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:808
				// _ = "end of CoverTab[187755]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:809
				_go_fuzz_dep_.CoverTab[187756]++
														d.saveError(&UnmarshalTypeError{"bool", v.Type(), int64(d.off)})
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:810
				// _ = "end of CoverTab[187756]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:811
			// _ = "end of CoverTab[187752]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:812
		// _ = "end of CoverTab[187743]"

	case '"':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:814
		_go_fuzz_dep_.CoverTab[187744]++
												s, ok := unquoteBytes(item)
												if !ok {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:816
			_go_fuzz_dep_.CoverTab[187757]++
													if fromQuoted {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:817
				_go_fuzz_dep_.CoverTab[187758]++
														d.error(fmt.Errorf("json: invalid use of ,string struct tag, trying to unmarshal %q into %v", item, v.Type()))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:818
				// _ = "end of CoverTab[187758]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:819
				_go_fuzz_dep_.CoverTab[187759]++
														d.error(errPhase)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:820
				// _ = "end of CoverTab[187759]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:821
			// _ = "end of CoverTab[187757]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:822
			_go_fuzz_dep_.CoverTab[187760]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:822
			// _ = "end of CoverTab[187760]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:822
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:822
		// _ = "end of CoverTab[187744]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:822
		_go_fuzz_dep_.CoverTab[187745]++
												switch v.Kind() {
		default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:824
			_go_fuzz_dep_.CoverTab[187761]++
													d.saveError(&UnmarshalTypeError{"string", v.Type(), int64(d.off)})
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:825
			// _ = "end of CoverTab[187761]"
		case reflect.Slice:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:826
			_go_fuzz_dep_.CoverTab[187762]++
													if v.Type().Elem().Kind() != reflect.Uint8 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:827
				_go_fuzz_dep_.CoverTab[187767]++
														d.saveError(&UnmarshalTypeError{"string", v.Type(), int64(d.off)})
														break
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:829
				// _ = "end of CoverTab[187767]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:830
				_go_fuzz_dep_.CoverTab[187768]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:830
				// _ = "end of CoverTab[187768]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:830
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:830
			// _ = "end of CoverTab[187762]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:830
			_go_fuzz_dep_.CoverTab[187763]++
													b := make([]byte, base64.StdEncoding.DecodedLen(len(s)))
													n, err := base64.StdEncoding.Decode(b, s)
													if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:833
				_go_fuzz_dep_.CoverTab[187769]++
														d.saveError(err)
														break
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:835
				// _ = "end of CoverTab[187769]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:836
				_go_fuzz_dep_.CoverTab[187770]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:836
				// _ = "end of CoverTab[187770]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:836
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:836
			// _ = "end of CoverTab[187763]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:836
			_go_fuzz_dep_.CoverTab[187764]++
													v.SetBytes(b[:n])
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:837
			// _ = "end of CoverTab[187764]"
		case reflect.String:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:838
			_go_fuzz_dep_.CoverTab[187765]++
													v.SetString(string(s))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:839
			// _ = "end of CoverTab[187765]"
		case reflect.Interface:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:840
			_go_fuzz_dep_.CoverTab[187766]++
													if v.NumMethod() == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:841
				_go_fuzz_dep_.CoverTab[187771]++
														v.Set(reflect.ValueOf(string(s)))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:842
				// _ = "end of CoverTab[187771]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:843
				_go_fuzz_dep_.CoverTab[187772]++
														d.saveError(&UnmarshalTypeError{"string", v.Type(), int64(d.off)})
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:844
				// _ = "end of CoverTab[187772]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:845
			// _ = "end of CoverTab[187766]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:846
		// _ = "end of CoverTab[187745]"

	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:848
		_go_fuzz_dep_.CoverTab[187746]++
												if c != '-' && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:849
			_go_fuzz_dep_.CoverTab[187773]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:849
			return (c < '0' || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:849
				_go_fuzz_dep_.CoverTab[187774]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:849
				return c > '9'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:849
				// _ = "end of CoverTab[187774]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:849
			}())
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:849
			// _ = "end of CoverTab[187773]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:849
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:849
			_go_fuzz_dep_.CoverTab[187775]++
													if fromQuoted {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:850
				_go_fuzz_dep_.CoverTab[187776]++
														d.error(fmt.Errorf("json: invalid use of ,string struct tag, trying to unmarshal %q into %v", item, v.Type()))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:851
				// _ = "end of CoverTab[187776]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:852
				_go_fuzz_dep_.CoverTab[187777]++
														d.error(errPhase)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:853
				// _ = "end of CoverTab[187777]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:854
			// _ = "end of CoverTab[187775]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:855
			_go_fuzz_dep_.CoverTab[187778]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:855
			// _ = "end of CoverTab[187778]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:855
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:855
		// _ = "end of CoverTab[187746]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:855
		_go_fuzz_dep_.CoverTab[187747]++
												s := string(item)
												switch v.Kind() {
		default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:858
			_go_fuzz_dep_.CoverTab[187779]++
													if v.Kind() == reflect.String && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:859
				_go_fuzz_dep_.CoverTab[187790]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:859
				return v.Type() == numberType
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:859
				// _ = "end of CoverTab[187790]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:859
			}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:859
				_go_fuzz_dep_.CoverTab[187791]++
														v.SetString(s)
														if !isValidNumber(s) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:861
					_go_fuzz_dep_.CoverTab[187793]++
															d.error(fmt.Errorf("json: invalid number literal, trying to unmarshal %q into Number", item))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:862
					// _ = "end of CoverTab[187793]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:863
					_go_fuzz_dep_.CoverTab[187794]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:863
					// _ = "end of CoverTab[187794]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:863
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:863
				// _ = "end of CoverTab[187791]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:863
				_go_fuzz_dep_.CoverTab[187792]++
														break
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:864
				// _ = "end of CoverTab[187792]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:865
				_go_fuzz_dep_.CoverTab[187795]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:865
				// _ = "end of CoverTab[187795]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:865
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:865
			// _ = "end of CoverTab[187779]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:865
			_go_fuzz_dep_.CoverTab[187780]++
													if fromQuoted {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:866
				_go_fuzz_dep_.CoverTab[187796]++
														d.error(fmt.Errorf("json: invalid use of ,string struct tag, trying to unmarshal %q into %v", item, v.Type()))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:867
				// _ = "end of CoverTab[187796]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:868
				_go_fuzz_dep_.CoverTab[187797]++
														d.error(&UnmarshalTypeError{"number", v.Type(), int64(d.off)})
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:869
				// _ = "end of CoverTab[187797]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:870
			// _ = "end of CoverTab[187780]"
		case reflect.Interface:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:871
			_go_fuzz_dep_.CoverTab[187781]++
													n, err := d.convertNumber(s)
													if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:873
				_go_fuzz_dep_.CoverTab[187798]++
														d.saveError(err)
														break
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:875
				// _ = "end of CoverTab[187798]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:876
				_go_fuzz_dep_.CoverTab[187799]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:876
				// _ = "end of CoverTab[187799]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:876
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:876
			// _ = "end of CoverTab[187781]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:876
			_go_fuzz_dep_.CoverTab[187782]++
													if v.NumMethod() != 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:877
				_go_fuzz_dep_.CoverTab[187800]++
														d.saveError(&UnmarshalTypeError{"number", v.Type(), int64(d.off)})
														break
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:879
				// _ = "end of CoverTab[187800]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:880
				_go_fuzz_dep_.CoverTab[187801]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:880
				// _ = "end of CoverTab[187801]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:880
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:880
			// _ = "end of CoverTab[187782]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:880
			_go_fuzz_dep_.CoverTab[187783]++
													v.Set(reflect.ValueOf(n))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:881
			// _ = "end of CoverTab[187783]"

		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:883
			_go_fuzz_dep_.CoverTab[187784]++
													n, err := strconv.ParseInt(s, 10, 64)
													if err != nil || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:885
				_go_fuzz_dep_.CoverTab[187802]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:885
				return v.OverflowInt(n)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:885
				// _ = "end of CoverTab[187802]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:885
			}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:885
				_go_fuzz_dep_.CoverTab[187803]++
														d.saveError(&UnmarshalTypeError{"number " + s, v.Type(), int64(d.off)})
														break
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:887
				// _ = "end of CoverTab[187803]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:888
				_go_fuzz_dep_.CoverTab[187804]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:888
				// _ = "end of CoverTab[187804]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:888
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:888
			// _ = "end of CoverTab[187784]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:888
			_go_fuzz_dep_.CoverTab[187785]++
													v.SetInt(n)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:889
			// _ = "end of CoverTab[187785]"

		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:891
			_go_fuzz_dep_.CoverTab[187786]++
													n, err := strconv.ParseUint(s, 10, 64)
													if err != nil || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:893
				_go_fuzz_dep_.CoverTab[187805]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:893
				return v.OverflowUint(n)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:893
				// _ = "end of CoverTab[187805]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:893
			}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:893
				_go_fuzz_dep_.CoverTab[187806]++
														d.saveError(&UnmarshalTypeError{"number " + s, v.Type(), int64(d.off)})
														break
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:895
				// _ = "end of CoverTab[187806]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:896
				_go_fuzz_dep_.CoverTab[187807]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:896
				// _ = "end of CoverTab[187807]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:896
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:896
			// _ = "end of CoverTab[187786]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:896
			_go_fuzz_dep_.CoverTab[187787]++
													v.SetUint(n)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:897
			// _ = "end of CoverTab[187787]"

		case reflect.Float32, reflect.Float64:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:899
			_go_fuzz_dep_.CoverTab[187788]++
													n, err := strconv.ParseFloat(s, v.Type().Bits())
													if err != nil || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:901
				_go_fuzz_dep_.CoverTab[187808]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:901
				return v.OverflowFloat(n)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:901
				// _ = "end of CoverTab[187808]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:901
			}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:901
				_go_fuzz_dep_.CoverTab[187809]++
														d.saveError(&UnmarshalTypeError{"number " + s, v.Type(), int64(d.off)})
														break
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:903
				// _ = "end of CoverTab[187809]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:904
				_go_fuzz_dep_.CoverTab[187810]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:904
				// _ = "end of CoverTab[187810]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:904
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:904
			// _ = "end of CoverTab[187788]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:904
			_go_fuzz_dep_.CoverTab[187789]++
													v.SetFloat(n)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:905
			// _ = "end of CoverTab[187789]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:906
		// _ = "end of CoverTab[187747]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:907
	// _ = "end of CoverTab[187718]"
}

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:914
// valueInterface is like value but returns interface{}
func (d *decodeState) valueInterface() interface{} {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:915
	_go_fuzz_dep_.CoverTab[187811]++
											switch d.scanWhile(scanSkipSpace) {
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:917
		_go_fuzz_dep_.CoverTab[187812]++
												d.error(errPhase)
												panic("unreachable")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:919
		// _ = "end of CoverTab[187812]"
	case scanBeginArray:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:920
		_go_fuzz_dep_.CoverTab[187813]++
												return d.arrayInterface()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:921
		// _ = "end of CoverTab[187813]"
	case scanBeginObject:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:922
		_go_fuzz_dep_.CoverTab[187814]++
												return d.objectInterface()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:923
		// _ = "end of CoverTab[187814]"
	case scanBeginLiteral:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:924
		_go_fuzz_dep_.CoverTab[187815]++
												return d.literalInterface()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:925
		// _ = "end of CoverTab[187815]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:926
	// _ = "end of CoverTab[187811]"
}

// arrayInterface is like array but returns []interface{}.
func (d *decodeState) arrayInterface() []interface{} {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:930
	_go_fuzz_dep_.CoverTab[187816]++
											var v = make([]interface{}, 0)
											for {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:932
		_go_fuzz_dep_.CoverTab[187818]++

												op := d.scanWhile(scanSkipSpace)
												if op == scanEndArray {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:935
			_go_fuzz_dep_.CoverTab[187821]++
													break
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:936
			// _ = "end of CoverTab[187821]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:937
			_go_fuzz_dep_.CoverTab[187822]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:937
			// _ = "end of CoverTab[187822]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:937
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:937
		// _ = "end of CoverTab[187818]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:937
		_go_fuzz_dep_.CoverTab[187819]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:940
		d.off--
												d.scan.undo(op)

												v = append(v, d.valueInterface())

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:946
		op = d.scanWhile(scanSkipSpace)
		if op == scanEndArray {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:947
			_go_fuzz_dep_.CoverTab[187823]++
													break
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:948
			// _ = "end of CoverTab[187823]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:949
			_go_fuzz_dep_.CoverTab[187824]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:949
			// _ = "end of CoverTab[187824]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:949
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:949
		// _ = "end of CoverTab[187819]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:949
		_go_fuzz_dep_.CoverTab[187820]++
												if op != scanArrayValue {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:950
			_go_fuzz_dep_.CoverTab[187825]++
													d.error(errPhase)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:951
			// _ = "end of CoverTab[187825]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:952
			_go_fuzz_dep_.CoverTab[187826]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:952
			// _ = "end of CoverTab[187826]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:952
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:952
		// _ = "end of CoverTab[187820]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:953
	// _ = "end of CoverTab[187816]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:953
	_go_fuzz_dep_.CoverTab[187817]++
											return v
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:954
	// _ = "end of CoverTab[187817]"
}

// objectInterface is like object but returns map[string]interface{}.
func (d *decodeState) objectInterface() map[string]interface{} {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:958
	_go_fuzz_dep_.CoverTab[187827]++
											m := make(map[string]interface{})
											keys := map[string]bool{}

											for {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:962
		_go_fuzz_dep_.CoverTab[187829]++

												op := d.scanWhile(scanSkipSpace)
												if op == scanEndObject {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:965
			_go_fuzz_dep_.CoverTab[187837]++

													break
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:967
			// _ = "end of CoverTab[187837]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:968
			_go_fuzz_dep_.CoverTab[187838]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:968
			// _ = "end of CoverTab[187838]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:968
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:968
		// _ = "end of CoverTab[187829]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:968
		_go_fuzz_dep_.CoverTab[187830]++
												if op != scanBeginLiteral {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:969
			_go_fuzz_dep_.CoverTab[187839]++
													d.error(errPhase)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:970
			// _ = "end of CoverTab[187839]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:971
			_go_fuzz_dep_.CoverTab[187840]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:971
			// _ = "end of CoverTab[187840]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:971
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:971
		// _ = "end of CoverTab[187830]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:971
		_go_fuzz_dep_.CoverTab[187831]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:974
		start := d.off - 1
		op = d.scanWhile(scanContinue)
		item := d.data[start : d.off-1]
		key, ok := unquote(item)
		if !ok {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:978
			_go_fuzz_dep_.CoverTab[187841]++
													d.error(errPhase)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:979
			// _ = "end of CoverTab[187841]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:980
			_go_fuzz_dep_.CoverTab[187842]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:980
			// _ = "end of CoverTab[187842]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:980
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:980
		// _ = "end of CoverTab[187831]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:980
		_go_fuzz_dep_.CoverTab[187832]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:983
		_, ok = keys[key]
		if !ok {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:984
			_go_fuzz_dep_.CoverTab[187843]++
													keys[key] = true
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:985
			// _ = "end of CoverTab[187843]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:986
			_go_fuzz_dep_.CoverTab[187844]++
													d.error(fmt.Errorf("json: duplicate key '%s' in object", key))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:987
			// _ = "end of CoverTab[187844]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:988
		// _ = "end of CoverTab[187832]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:988
		_go_fuzz_dep_.CoverTab[187833]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:991
		if op == scanSkipSpace {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:991
			_go_fuzz_dep_.CoverTab[187845]++
													op = d.scanWhile(scanSkipSpace)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:992
			// _ = "end of CoverTab[187845]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:993
			_go_fuzz_dep_.CoverTab[187846]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:993
			// _ = "end of CoverTab[187846]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:993
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:993
		// _ = "end of CoverTab[187833]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:993
		_go_fuzz_dep_.CoverTab[187834]++
												if op != scanObjectKey {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:994
			_go_fuzz_dep_.CoverTab[187847]++
													d.error(errPhase)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:995
			// _ = "end of CoverTab[187847]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:996
			_go_fuzz_dep_.CoverTab[187848]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:996
			// _ = "end of CoverTab[187848]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:996
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:996
		// _ = "end of CoverTab[187834]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:996
		_go_fuzz_dep_.CoverTab[187835]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:999
		m[key] = d.valueInterface()

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1002
		op = d.scanWhile(scanSkipSpace)
		if op == scanEndObject {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1003
			_go_fuzz_dep_.CoverTab[187849]++
														break
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1004
			// _ = "end of CoverTab[187849]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1005
			_go_fuzz_dep_.CoverTab[187850]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1005
			// _ = "end of CoverTab[187850]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1005
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1005
		// _ = "end of CoverTab[187835]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1005
		_go_fuzz_dep_.CoverTab[187836]++
													if op != scanObjectValue {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1006
			_go_fuzz_dep_.CoverTab[187851]++
														d.error(errPhase)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1007
			// _ = "end of CoverTab[187851]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1008
			_go_fuzz_dep_.CoverTab[187852]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1008
			// _ = "end of CoverTab[187852]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1008
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1008
		// _ = "end of CoverTab[187836]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1009
	// _ = "end of CoverTab[187827]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1009
	_go_fuzz_dep_.CoverTab[187828]++
												return m
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1010
	// _ = "end of CoverTab[187828]"
}

// literalInterface is like literal but returns an interface value.
func (d *decodeState) literalInterface() interface{} {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1014
	_go_fuzz_dep_.CoverTab[187853]++

												start := d.off - 1
												op := d.scanWhile(scanContinue)

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1020
	d.off--
	d.scan.undo(op)
	item := d.data[start:d.off]

	switch c := item[0]; c {
	case 'n':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1025
		_go_fuzz_dep_.CoverTab[187854]++
													return nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1026
		// _ = "end of CoverTab[187854]"

	case 't', 'f':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1028
		_go_fuzz_dep_.CoverTab[187855]++
													return c == 't'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1029
		// _ = "end of CoverTab[187855]"

	case '"':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1031
		_go_fuzz_dep_.CoverTab[187856]++
													s, ok := unquote(item)
													if !ok {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1033
			_go_fuzz_dep_.CoverTab[187861]++
														d.error(errPhase)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1034
			// _ = "end of CoverTab[187861]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1035
			_go_fuzz_dep_.CoverTab[187862]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1035
			// _ = "end of CoverTab[187862]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1035
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1035
		// _ = "end of CoverTab[187856]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1035
		_go_fuzz_dep_.CoverTab[187857]++
													return s
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1036
		// _ = "end of CoverTab[187857]"

	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1038
		_go_fuzz_dep_.CoverTab[187858]++
													if c != '-' && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1039
			_go_fuzz_dep_.CoverTab[187863]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1039
			return (c < '0' || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1039
				_go_fuzz_dep_.CoverTab[187864]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1039
				return c > '9'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1039
				// _ = "end of CoverTab[187864]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1039
			}())
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1039
			// _ = "end of CoverTab[187863]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1039
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1039
			_go_fuzz_dep_.CoverTab[187865]++
														d.error(errPhase)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1040
			// _ = "end of CoverTab[187865]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1041
			_go_fuzz_dep_.CoverTab[187866]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1041
			// _ = "end of CoverTab[187866]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1041
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1041
		// _ = "end of CoverTab[187858]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1041
		_go_fuzz_dep_.CoverTab[187859]++
													n, err := d.convertNumber(string(item))
													if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1043
			_go_fuzz_dep_.CoverTab[187867]++
														d.saveError(err)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1044
			// _ = "end of CoverTab[187867]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1045
			_go_fuzz_dep_.CoverTab[187868]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1045
			// _ = "end of CoverTab[187868]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1045
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1045
		// _ = "end of CoverTab[187859]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1045
		_go_fuzz_dep_.CoverTab[187860]++
													return n
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1046
		// _ = "end of CoverTab[187860]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1047
	// _ = "end of CoverTab[187853]"
}

// getu4 decodes \uXXXX from the beginning of s, returning the hex value,
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1050
// or it returns -1.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1052
func getu4(s []byte) rune {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1052
	_go_fuzz_dep_.CoverTab[187869]++
												if len(s) < 6 || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1053
		_go_fuzz_dep_.CoverTab[187872]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1053
		return s[0] != '\\'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1053
		// _ = "end of CoverTab[187872]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1053
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1053
		_go_fuzz_dep_.CoverTab[187873]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1053
		return s[1] != 'u'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1053
		// _ = "end of CoverTab[187873]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1053
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1053
		_go_fuzz_dep_.CoverTab[187874]++
													return -1
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1054
		// _ = "end of CoverTab[187874]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1055
		_go_fuzz_dep_.CoverTab[187875]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1055
		// _ = "end of CoverTab[187875]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1055
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1055
	// _ = "end of CoverTab[187869]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1055
	_go_fuzz_dep_.CoverTab[187870]++
												r, err := strconv.ParseUint(string(s[2:6]), 16, 64)
												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1057
		_go_fuzz_dep_.CoverTab[187876]++
													return -1
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1058
		// _ = "end of CoverTab[187876]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1059
		_go_fuzz_dep_.CoverTab[187877]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1059
		// _ = "end of CoverTab[187877]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1059
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1059
	// _ = "end of CoverTab[187870]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1059
	_go_fuzz_dep_.CoverTab[187871]++
												return rune(r)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1060
	// _ = "end of CoverTab[187871]"
}

// unquote converts a quoted JSON string literal s into an actual string t.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1063
// The rules are different than for Go, so cannot use strconv.Unquote.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1065
func unquote(s []byte) (t string, ok bool) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1065
	_go_fuzz_dep_.CoverTab[187878]++
												s, ok = unquoteBytes(s)
												t = string(s)
												return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1068
	// _ = "end of CoverTab[187878]"
}

func unquoteBytes(s []byte) (t []byte, ok bool) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1071
	_go_fuzz_dep_.CoverTab[187879]++
												if len(s) < 2 || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1072
		_go_fuzz_dep_.CoverTab[187884]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1072
		return s[0] != '"'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1072
		// _ = "end of CoverTab[187884]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1072
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1072
		_go_fuzz_dep_.CoverTab[187885]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1072
		return s[len(s)-1] != '"'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1072
		// _ = "end of CoverTab[187885]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1072
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1072
		_go_fuzz_dep_.CoverTab[187886]++
													return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1073
		// _ = "end of CoverTab[187886]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1074
		_go_fuzz_dep_.CoverTab[187887]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1074
		// _ = "end of CoverTab[187887]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1074
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1074
	// _ = "end of CoverTab[187879]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1074
	_go_fuzz_dep_.CoverTab[187880]++
												s = s[1 : len(s)-1]

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1080
	r := 0
	for r < len(s) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1081
		_go_fuzz_dep_.CoverTab[187888]++
													c := s[r]
													if c == '\\' || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1083
			_go_fuzz_dep_.CoverTab[187892]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1083
			return c == '"'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1083
			// _ = "end of CoverTab[187892]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1083
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1083
			_go_fuzz_dep_.CoverTab[187893]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1083
			return c < ' '
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1083
			// _ = "end of CoverTab[187893]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1083
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1083
			_go_fuzz_dep_.CoverTab[187894]++
														break
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1084
			// _ = "end of CoverTab[187894]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1085
			_go_fuzz_dep_.CoverTab[187895]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1085
			// _ = "end of CoverTab[187895]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1085
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1085
		// _ = "end of CoverTab[187888]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1085
		_go_fuzz_dep_.CoverTab[187889]++
													if c < utf8.RuneSelf {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1086
			_go_fuzz_dep_.CoverTab[187896]++
														r++
														continue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1088
			// _ = "end of CoverTab[187896]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1089
			_go_fuzz_dep_.CoverTab[187897]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1089
			// _ = "end of CoverTab[187897]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1089
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1089
		// _ = "end of CoverTab[187889]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1089
		_go_fuzz_dep_.CoverTab[187890]++
													rr, size := utf8.DecodeRune(s[r:])
													if rr == utf8.RuneError && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1091
			_go_fuzz_dep_.CoverTab[187898]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1091
			return size == 1
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1091
			// _ = "end of CoverTab[187898]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1091
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1091
			_go_fuzz_dep_.CoverTab[187899]++
														break
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1092
			// _ = "end of CoverTab[187899]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1093
			_go_fuzz_dep_.CoverTab[187900]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1093
			// _ = "end of CoverTab[187900]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1093
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1093
		// _ = "end of CoverTab[187890]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1093
		_go_fuzz_dep_.CoverTab[187891]++
													r += size
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1094
		// _ = "end of CoverTab[187891]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1095
	// _ = "end of CoverTab[187880]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1095
	_go_fuzz_dep_.CoverTab[187881]++
												if r == len(s) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1096
		_go_fuzz_dep_.CoverTab[187901]++
													return s, true
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1097
		// _ = "end of CoverTab[187901]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1098
		_go_fuzz_dep_.CoverTab[187902]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1098
		// _ = "end of CoverTab[187902]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1098
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1098
	// _ = "end of CoverTab[187881]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1098
	_go_fuzz_dep_.CoverTab[187882]++

												b := make([]byte, len(s)+2*utf8.UTFMax)
												w := copy(b, s[0:r])
												for r < len(s) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1102
		_go_fuzz_dep_.CoverTab[187903]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1106
		if w >= len(b)-2*utf8.UTFMax {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1106
			_go_fuzz_dep_.CoverTab[187905]++
														nb := make([]byte, (len(b)+utf8.UTFMax)*2)
														copy(nb, b[0:w])
														b = nb
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1109
			// _ = "end of CoverTab[187905]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1110
			_go_fuzz_dep_.CoverTab[187906]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1110
			// _ = "end of CoverTab[187906]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1110
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1110
		// _ = "end of CoverTab[187903]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1110
		_go_fuzz_dep_.CoverTab[187904]++
													switch c := s[r]; {
		case c == '\\':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1112
			_go_fuzz_dep_.CoverTab[187907]++
														r++
														if r >= len(s) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1114
				_go_fuzz_dep_.CoverTab[187912]++
															return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1115
				// _ = "end of CoverTab[187912]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1116
				_go_fuzz_dep_.CoverTab[187913]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1116
				// _ = "end of CoverTab[187913]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1116
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1116
			// _ = "end of CoverTab[187907]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1116
			_go_fuzz_dep_.CoverTab[187908]++
														switch s[r] {
			default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1118
				_go_fuzz_dep_.CoverTab[187914]++
															return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1119
				// _ = "end of CoverTab[187914]"
			case '"', '\\', '/', '\'':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1120
				_go_fuzz_dep_.CoverTab[187915]++
															b[w] = s[r]
															r++
															w++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1123
				// _ = "end of CoverTab[187915]"
			case 'b':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1124
				_go_fuzz_dep_.CoverTab[187916]++
															b[w] = '\b'
															r++
															w++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1127
				// _ = "end of CoverTab[187916]"
			case 'f':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1128
				_go_fuzz_dep_.CoverTab[187917]++
															b[w] = '\f'
															r++
															w++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1131
				// _ = "end of CoverTab[187917]"
			case 'n':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1132
				_go_fuzz_dep_.CoverTab[187918]++
															b[w] = '\n'
															r++
															w++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1135
				// _ = "end of CoverTab[187918]"
			case 'r':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1136
				_go_fuzz_dep_.CoverTab[187919]++
															b[w] = '\r'
															r++
															w++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1139
				// _ = "end of CoverTab[187919]"
			case 't':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1140
				_go_fuzz_dep_.CoverTab[187920]++
															b[w] = '\t'
															r++
															w++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1143
				// _ = "end of CoverTab[187920]"
			case 'u':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1144
				_go_fuzz_dep_.CoverTab[187921]++
															r--
															rr := getu4(s[r:])
															if rr < 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1147
					_go_fuzz_dep_.CoverTab[187924]++
																return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1148
					// _ = "end of CoverTab[187924]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1149
					_go_fuzz_dep_.CoverTab[187925]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1149
					// _ = "end of CoverTab[187925]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1149
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1149
				// _ = "end of CoverTab[187921]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1149
				_go_fuzz_dep_.CoverTab[187922]++
															r += 6
															if utf16.IsSurrogate(rr) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1151
					_go_fuzz_dep_.CoverTab[187926]++
																rr1 := getu4(s[r:])
																if dec := utf16.DecodeRune(rr, rr1); dec != unicode.ReplacementChar {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1153
						_go_fuzz_dep_.CoverTab[187928]++

																	r += 6
																	w += utf8.EncodeRune(b[w:], dec)
																	break
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1157
						// _ = "end of CoverTab[187928]"
					} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1158
						_go_fuzz_dep_.CoverTab[187929]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1158
						// _ = "end of CoverTab[187929]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1158
					}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1158
					// _ = "end of CoverTab[187926]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1158
					_go_fuzz_dep_.CoverTab[187927]++

																rr = unicode.ReplacementChar
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1160
					// _ = "end of CoverTab[187927]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1161
					_go_fuzz_dep_.CoverTab[187930]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1161
					// _ = "end of CoverTab[187930]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1161
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1161
				// _ = "end of CoverTab[187922]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1161
				_go_fuzz_dep_.CoverTab[187923]++
															w += utf8.EncodeRune(b[w:], rr)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1162
				// _ = "end of CoverTab[187923]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1163
			// _ = "end of CoverTab[187908]"

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1166
		case c == '"', c < ' ':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1166
			_go_fuzz_dep_.CoverTab[187909]++
														return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1167
			// _ = "end of CoverTab[187909]"

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1170
		case c < utf8.RuneSelf:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1170
			_go_fuzz_dep_.CoverTab[187910]++
														b[w] = c
														r++
														w++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1173
			// _ = "end of CoverTab[187910]"

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1176
		default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1176
			_go_fuzz_dep_.CoverTab[187911]++
														rr, size := utf8.DecodeRune(s[r:])
														r += size
														w += utf8.EncodeRune(b[w:], rr)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1179
			// _ = "end of CoverTab[187911]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1180
		// _ = "end of CoverTab[187904]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1181
	// _ = "end of CoverTab[187882]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1181
	_go_fuzz_dep_.CoverTab[187883]++
												return b[0:w], true
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1182
	// _ = "end of CoverTab[187883]"
}

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1183
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/decode.go:1183
var _ = _go_fuzz_dep_.CoverTab
