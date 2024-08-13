// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Represents JSON data structure using native Go types: booleans, floats,
// strings, arrays, and maps.

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:8
package json

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:8
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:8
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:8
)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:8
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:8
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:8
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
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:24
// in the value pointed to by v.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:24
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:24
// Unmarshal uses the inverse of the encodings that
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:24
// Marshal uses, allocating maps, slices, and pointers as necessary,
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:24
// with the following additional rules:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:24
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:24
// To unmarshal JSON into a pointer, Unmarshal first handles the case of
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:24
// the JSON being the JSON literal null.  In that case, Unmarshal sets
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:24
// the pointer to nil.  Otherwise, Unmarshal unmarshals the JSON into
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:24
// the value pointed at by the pointer.  If the pointer is nil, Unmarshal
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:24
// allocates a new value for it to point to.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:24
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:24
// To unmarshal JSON into a struct, Unmarshal matches incoming object
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:24
// keys to the keys used by Marshal (either the struct field name or its tag),
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:24
// preferring an exact match but also accepting a case-insensitive match.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:24
// Unmarshal will only set exported fields of the struct.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:24
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:24
// To unmarshal JSON into an interface value,
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:24
// Unmarshal stores one of these in the interface value:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:24
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:24
//	bool, for JSON booleans
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:24
//	float64, for JSON numbers
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:24
//	string, for JSON strings
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:24
//	[]interface{}, for JSON arrays
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:24
//	map[string]interface{}, for JSON objects
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:24
//	nil for JSON null
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:24
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:24
// To unmarshal a JSON array into a slice, Unmarshal resets the slice length
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:24
// to zero and then appends each element to the slice.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:24
// As a special case, to unmarshal an empty JSON array into a slice,
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:24
// Unmarshal replaces the slice with a new empty slice.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:24
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:24
// To unmarshal a JSON array into a Go array, Unmarshal decodes
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:24
// JSON array elements into corresponding Go array elements.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:24
// If the Go array is smaller than the JSON array,
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:24
// the additional JSON array elements are discarded.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:24
// If the JSON array is smaller than the Go array,
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:24
// the additional Go array elements are set to zero values.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:24
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:24
// To unmarshal a JSON object into a string-keyed map, Unmarshal first
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:24
// establishes a map to use, If the map is nil, Unmarshal allocates a new map.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:24
// Otherwise Unmarshal reuses the existing map, keeping existing entries.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:24
// Unmarshal then stores key-value pairs from the JSON object into the map.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:24
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:24
// If a JSON value is not appropriate for a given target type,
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:24
// or if a JSON number overflows the target type, Unmarshal
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:24
// skips that field and completes the unmarshaling as best it can.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:24
// If no more serious errors are encountered, Unmarshal returns
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:24
// an UnmarshalTypeError describing the earliest such error.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:24
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:24
// The JSON null value unmarshals into an interface, map, pointer, or slice
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:24
// by setting that Go value to nil. Because null is often used in JSON to mean
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:24
// “not present,” unmarshaling a JSON null into any other Go type has no effect
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:24
// on the value and produces no error.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:24
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:24
// When unmarshaling quoted strings, invalid UTF-8 or
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:24
// invalid UTF-16 surrogate pairs are not treated as an error.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:24
// Instead, they are replaced by the Unicode replacement
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:24
// character U+FFFD.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:85
func Unmarshal(data []byte, v interface{}) error {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:85
	_go_fuzz_dep_.CoverTab[184375]++
	// Check for well-formedness.
	// Avoids filling out half a data structure
	// before discovering a JSON syntax error.
	var d decodeState
	err := checkValid(data, &d.scan)
	if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:91
		_go_fuzz_dep_.CoverTab[184377]++
												return err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:92
		// _ = "end of CoverTab[184377]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:93
		_go_fuzz_dep_.CoverTab[184378]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:93
		// _ = "end of CoverTab[184378]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:93
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:93
	// _ = "end of CoverTab[184375]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:93
	_go_fuzz_dep_.CoverTab[184376]++

											d.init(data)
											return d.unmarshal(v)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:96
	// _ = "end of CoverTab[184376]"
}

// Unmarshaler is the interface implemented by objects
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:99
// that can unmarshal a JSON description of themselves.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:99
// The input can be assumed to be a valid encoding of
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:99
// a JSON value. UnmarshalJSON must copy the JSON data
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:99
// if it wishes to retain the data after returning.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:104
type Unmarshaler interface {
	UnmarshalJSON([]byte) error
}

// An UnmarshalTypeError describes a JSON value that was
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:108
// not appropriate for a value of a specific Go type.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:110
type UnmarshalTypeError struct {
	Value	string		// description of JSON value - "bool", "array", "number -5"
	Type	reflect.Type	// type of Go value it could not be assigned to
	Offset	int64		// error occurred after reading Offset bytes
}

func (e *UnmarshalTypeError) Error() string {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:116
	_go_fuzz_dep_.CoverTab[184379]++
											return "json: cannot unmarshal " + e.Value + " into Go value of type " + e.Type.String()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:117
	// _ = "end of CoverTab[184379]"
}

// An UnmarshalFieldError describes a JSON object key that
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:120
// led to an unexported (and therefore unwritable) struct field.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:120
// (No longer used; kept for compatibility.)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:123
type UnmarshalFieldError struct {
	Key	string
	Type	reflect.Type
	Field	reflect.StructField
}

func (e *UnmarshalFieldError) Error() string {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:129
	_go_fuzz_dep_.CoverTab[184380]++
											return "json: cannot unmarshal object key " + strconv.Quote(e.Key) + " into unexported field " + e.Field.Name + " of type " + e.Type.String()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:130
	// _ = "end of CoverTab[184380]"
}

// An InvalidUnmarshalError describes an invalid argument passed to Unmarshal.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:133
// (The argument to Unmarshal must be a non-nil pointer.)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:135
type InvalidUnmarshalError struct {
	Type reflect.Type
}

func (e *InvalidUnmarshalError) Error() string {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:139
	_go_fuzz_dep_.CoverTab[184381]++
											if e.Type == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:140
		_go_fuzz_dep_.CoverTab[184384]++
												return "json: Unmarshal(nil)"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:141
		// _ = "end of CoverTab[184384]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:142
		_go_fuzz_dep_.CoverTab[184385]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:142
		// _ = "end of CoverTab[184385]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:142
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:142
	// _ = "end of CoverTab[184381]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:142
	_go_fuzz_dep_.CoverTab[184382]++

											if e.Type.Kind() != reflect.Ptr {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:144
		_go_fuzz_dep_.CoverTab[184386]++
												return "json: Unmarshal(non-pointer " + e.Type.String() + ")"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:145
		// _ = "end of CoverTab[184386]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:146
		_go_fuzz_dep_.CoverTab[184387]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:146
		// _ = "end of CoverTab[184387]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:146
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:146
	// _ = "end of CoverTab[184382]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:146
	_go_fuzz_dep_.CoverTab[184383]++
											return "json: Unmarshal(nil " + e.Type.String() + ")"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:147
	// _ = "end of CoverTab[184383]"
}

func (d *decodeState) unmarshal(v interface{}) (err error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:150
	_go_fuzz_dep_.CoverTab[184388]++
											defer func() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:151
		_go_fuzz_dep_.CoverTab[184391]++
												if r := recover(); r != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:152
			_go_fuzz_dep_.CoverTab[184392]++
													if _, ok := r.(runtime.Error); ok {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:153
				_go_fuzz_dep_.CoverTab[184394]++
														panic(r)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:154
				// _ = "end of CoverTab[184394]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:155
				_go_fuzz_dep_.CoverTab[184395]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:155
				// _ = "end of CoverTab[184395]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:155
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:155
			// _ = "end of CoverTab[184392]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:155
			_go_fuzz_dep_.CoverTab[184393]++
													err = r.(error)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:156
			// _ = "end of CoverTab[184393]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:157
			_go_fuzz_dep_.CoverTab[184396]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:157
			// _ = "end of CoverTab[184396]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:157
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:157
		// _ = "end of CoverTab[184391]"
	}()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:158
	// _ = "end of CoverTab[184388]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:158
	_go_fuzz_dep_.CoverTab[184389]++

											rv := reflect.ValueOf(v)
											if rv.Kind() != reflect.Ptr || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:161
		_go_fuzz_dep_.CoverTab[184397]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:161
		return rv.IsNil()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:161
		// _ = "end of CoverTab[184397]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:161
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:161
		_go_fuzz_dep_.CoverTab[184398]++
												return &InvalidUnmarshalError{reflect.TypeOf(v)}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:162
		// _ = "end of CoverTab[184398]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:163
		_go_fuzz_dep_.CoverTab[184399]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:163
		// _ = "end of CoverTab[184399]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:163
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:163
	// _ = "end of CoverTab[184389]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:163
	_go_fuzz_dep_.CoverTab[184390]++

											d.scan.reset()

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:168
	d.value(rv)
											return d.savedError
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:169
	// _ = "end of CoverTab[184390]"
}

// A Number represents a JSON number literal.
type Number string

// String returns the literal text of the number.
func (n Number) String() string {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:176
	_go_fuzz_dep_.CoverTab[184400]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:176
	return string(n)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:176
	// _ = "end of CoverTab[184400]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:176
}

// Float64 returns the number as a float64.
func (n Number) Float64() (float64, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:179
	_go_fuzz_dep_.CoverTab[184401]++
											return strconv.ParseFloat(string(n), 64)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:180
	// _ = "end of CoverTab[184401]"
}

// Int64 returns the number as an int64.
func (n Number) Int64() (int64, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:184
	_go_fuzz_dep_.CoverTab[184402]++
											return strconv.ParseInt(string(n), 10, 64)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:185
	// _ = "end of CoverTab[184402]"
}

// isValidNumber reports whether s is a valid JSON number literal.
func isValidNumber(s string) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:189
	_go_fuzz_dep_.CoverTab[184403]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:194
	if s == "" {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:194
		_go_fuzz_dep_.CoverTab[184409]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:195
		// _ = "end of CoverTab[184409]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:196
		_go_fuzz_dep_.CoverTab[184410]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:196
		// _ = "end of CoverTab[184410]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:196
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:196
	// _ = "end of CoverTab[184403]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:196
	_go_fuzz_dep_.CoverTab[184404]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:199
	if s[0] == '-' {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:199
		_go_fuzz_dep_.CoverTab[184411]++
												s = s[1:]
												if s == "" {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:201
			_go_fuzz_dep_.CoverTab[184412]++
													return false
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:202
			// _ = "end of CoverTab[184412]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:203
			_go_fuzz_dep_.CoverTab[184413]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:203
			// _ = "end of CoverTab[184413]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:203
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:203
		// _ = "end of CoverTab[184411]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:204
		_go_fuzz_dep_.CoverTab[184414]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:204
		// _ = "end of CoverTab[184414]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:204
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:204
	// _ = "end of CoverTab[184404]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:204
	_go_fuzz_dep_.CoverTab[184405]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:207
	switch {
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:208
		_go_fuzz_dep_.CoverTab[184415]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:209
		// _ = "end of CoverTab[184415]"

	case s[0] == '0':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:211
		_go_fuzz_dep_.CoverTab[184416]++
												s = s[1:]
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:212
		// _ = "end of CoverTab[184416]"

	case '1' <= s[0] && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:214
		_go_fuzz_dep_.CoverTab[184418]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:214
		return s[0] <= '9'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:214
		// _ = "end of CoverTab[184418]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:214
	}():
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:214
		_go_fuzz_dep_.CoverTab[184417]++
												s = s[1:]
												for len(s) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:216
			_go_fuzz_dep_.CoverTab[184419]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:216
			return '0' <= s[0]
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:216
			// _ = "end of CoverTab[184419]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:216
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:216
			_go_fuzz_dep_.CoverTab[184420]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:216
			return s[0] <= '9'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:216
			// _ = "end of CoverTab[184420]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:216
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:216
			_go_fuzz_dep_.CoverTab[184421]++
													s = s[1:]
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:217
			// _ = "end of CoverTab[184421]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:218
		// _ = "end of CoverTab[184417]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:219
	// _ = "end of CoverTab[184405]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:219
	_go_fuzz_dep_.CoverTab[184406]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:222
	if len(s) >= 2 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:222
		_go_fuzz_dep_.CoverTab[184422]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:222
		return s[0] == '.'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:222
		// _ = "end of CoverTab[184422]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:222
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:222
		_go_fuzz_dep_.CoverTab[184423]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:222
		return '0' <= s[1]
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:222
		// _ = "end of CoverTab[184423]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:222
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:222
		_go_fuzz_dep_.CoverTab[184424]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:222
		return s[1] <= '9'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:222
		// _ = "end of CoverTab[184424]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:222
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:222
		_go_fuzz_dep_.CoverTab[184425]++
												s = s[2:]
												for len(s) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:224
			_go_fuzz_dep_.CoverTab[184426]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:224
			return '0' <= s[0]
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:224
			// _ = "end of CoverTab[184426]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:224
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:224
			_go_fuzz_dep_.CoverTab[184427]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:224
			return s[0] <= '9'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:224
			// _ = "end of CoverTab[184427]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:224
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:224
			_go_fuzz_dep_.CoverTab[184428]++
													s = s[1:]
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:225
			// _ = "end of CoverTab[184428]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:226
		// _ = "end of CoverTab[184425]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:227
		_go_fuzz_dep_.CoverTab[184429]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:227
		// _ = "end of CoverTab[184429]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:227
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:227
	// _ = "end of CoverTab[184406]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:227
	_go_fuzz_dep_.CoverTab[184407]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:231
	if len(s) >= 2 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:231
		_go_fuzz_dep_.CoverTab[184430]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:231
		return (s[0] == 'e' || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:231
			_go_fuzz_dep_.CoverTab[184431]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:231
			return s[0] == 'E'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:231
			// _ = "end of CoverTab[184431]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:231
		}())
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:231
		// _ = "end of CoverTab[184430]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:231
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:231
		_go_fuzz_dep_.CoverTab[184432]++
												s = s[1:]
												if s[0] == '+' || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:233
			_go_fuzz_dep_.CoverTab[184434]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:233
			return s[0] == '-'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:233
			// _ = "end of CoverTab[184434]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:233
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:233
			_go_fuzz_dep_.CoverTab[184435]++
													s = s[1:]
													if s == "" {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:235
				_go_fuzz_dep_.CoverTab[184436]++
														return false
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:236
				// _ = "end of CoverTab[184436]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:237
				_go_fuzz_dep_.CoverTab[184437]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:237
				// _ = "end of CoverTab[184437]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:237
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:237
			// _ = "end of CoverTab[184435]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:238
			_go_fuzz_dep_.CoverTab[184438]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:238
			// _ = "end of CoverTab[184438]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:238
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:238
		// _ = "end of CoverTab[184432]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:238
		_go_fuzz_dep_.CoverTab[184433]++
												for len(s) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:239
			_go_fuzz_dep_.CoverTab[184439]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:239
			return '0' <= s[0]
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:239
			// _ = "end of CoverTab[184439]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:239
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:239
			_go_fuzz_dep_.CoverTab[184440]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:239
			return s[0] <= '9'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:239
			// _ = "end of CoverTab[184440]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:239
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:239
			_go_fuzz_dep_.CoverTab[184441]++
													s = s[1:]
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:240
			// _ = "end of CoverTab[184441]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:241
		// _ = "end of CoverTab[184433]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:242
		_go_fuzz_dep_.CoverTab[184442]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:242
		// _ = "end of CoverTab[184442]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:242
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:242
	// _ = "end of CoverTab[184407]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:242
	_go_fuzz_dep_.CoverTab[184408]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:245
	return s == ""
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:245
	// _ = "end of CoverTab[184408]"
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
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:258
// there is a bug in the JSON decoder or something is editing
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:258
// the data slice while the decoder executes.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:261
var errPhase = errors.New("JSON decoder out of sync - data changing underfoot?")

func (d *decodeState) init(data []byte) *decodeState {
	d.data = data
	d.off = 0
	d.savedError = nil
	return d
}

// error aborts the decoding by panicking with err.
func (d *decodeState) error(err error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:271
	_go_fuzz_dep_.CoverTab[184443]++
											panic(err)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:272
	// _ = "end of CoverTab[184443]"
}

// saveError saves the first err it is called with,
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:275
// for reporting at the end of the unmarshal.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:277
func (d *decodeState) saveError(err error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:277
	_go_fuzz_dep_.CoverTab[184444]++
											if d.savedError == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:278
		_go_fuzz_dep_.CoverTab[184445]++
												d.savedError = err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:279
		// _ = "end of CoverTab[184445]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:280
		_go_fuzz_dep_.CoverTab[184446]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:280
		// _ = "end of CoverTab[184446]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:280
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:280
	// _ = "end of CoverTab[184444]"
}

// next cuts off and returns the next full JSON value in d.data[d.off:].
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:283
// The next value is known to be an object or array, not a literal.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:285
func (d *decodeState) next() []byte {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:285
	_go_fuzz_dep_.CoverTab[184447]++
											c := d.data[d.off]
											item, rest, err := nextValue(d.data[d.off:], &d.nextscan)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:288
		_go_fuzz_dep_.CoverTab[184450]++
												d.error(err)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:289
		// _ = "end of CoverTab[184450]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:290
		_go_fuzz_dep_.CoverTab[184451]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:290
		// _ = "end of CoverTab[184451]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:290
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:290
	// _ = "end of CoverTab[184447]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:290
	_go_fuzz_dep_.CoverTab[184448]++
											d.off = len(d.data) - len(rest)

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:296
	if c == '{' {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:296
		_go_fuzz_dep_.CoverTab[184452]++
												d.scan.step(&d.scan, '}')
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:297
		// _ = "end of CoverTab[184452]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:298
		_go_fuzz_dep_.CoverTab[184453]++
												d.scan.step(&d.scan, ']')
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:299
		// _ = "end of CoverTab[184453]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:300
	// _ = "end of CoverTab[184448]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:300
	_go_fuzz_dep_.CoverTab[184449]++

											return item
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:302
	// _ = "end of CoverTab[184449]"
}

// scanWhile processes bytes in d.data[d.off:] until it
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:305
// receives a scan code not equal to op.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:305
// It updates d.off and returns the new scan code.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:308
func (d *decodeState) scanWhile(op int) int {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:308
	_go_fuzz_dep_.CoverTab[184454]++
											var newOp int
											for {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:310
		_go_fuzz_dep_.CoverTab[184456]++
												if d.off >= len(d.data) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:311
			_go_fuzz_dep_.CoverTab[184458]++
													newOp = d.scan.eof()
													d.off = len(d.data) + 1
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:313
			// _ = "end of CoverTab[184458]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:314
			_go_fuzz_dep_.CoverTab[184459]++
													c := d.data[d.off]
													d.off++
													newOp = d.scan.step(&d.scan, c)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:317
			// _ = "end of CoverTab[184459]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:318
		// _ = "end of CoverTab[184456]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:318
		_go_fuzz_dep_.CoverTab[184457]++
												if newOp != op {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:319
			_go_fuzz_dep_.CoverTab[184460]++
													break
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:320
			// _ = "end of CoverTab[184460]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:321
			_go_fuzz_dep_.CoverTab[184461]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:321
			// _ = "end of CoverTab[184461]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:321
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:321
		// _ = "end of CoverTab[184457]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:322
	// _ = "end of CoverTab[184454]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:322
	_go_fuzz_dep_.CoverTab[184455]++
											return newOp
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:323
	// _ = "end of CoverTab[184455]"
}

// value decodes a JSON value from d.data[d.off:] into the value.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:326
// it updates d.off to point past the decoded value.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:328
func (d *decodeState) value(v reflect.Value) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:328
	_go_fuzz_dep_.CoverTab[184462]++
											if !v.IsValid() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:329
		_go_fuzz_dep_.CoverTab[184464]++
												_, rest, err := nextValue(d.data[d.off:], &d.nextscan)
												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:331
			_go_fuzz_dep_.CoverTab[184468]++
													d.error(err)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:332
			// _ = "end of CoverTab[184468]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:333
			_go_fuzz_dep_.CoverTab[184469]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:333
			// _ = "end of CoverTab[184469]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:333
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:333
		// _ = "end of CoverTab[184464]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:333
		_go_fuzz_dep_.CoverTab[184465]++
												d.off = len(d.data) - len(rest)

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:339
		if d.scan.redo {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:339
			_go_fuzz_dep_.CoverTab[184470]++

													d.scan.redo = false
													d.scan.step = stateBeginValue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:342
			// _ = "end of CoverTab[184470]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:343
			_go_fuzz_dep_.CoverTab[184471]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:343
			// _ = "end of CoverTab[184471]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:343
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:343
		// _ = "end of CoverTab[184465]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:343
		_go_fuzz_dep_.CoverTab[184466]++
												d.scan.step(&d.scan, '"')
												d.scan.step(&d.scan, '"')

												n := len(d.scan.parseState)
												if n > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:348
			_go_fuzz_dep_.CoverTab[184472]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:348
			return d.scan.parseState[n-1] == parseObjectKey
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:348
			// _ = "end of CoverTab[184472]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:348
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:348
			_go_fuzz_dep_.CoverTab[184473]++

													d.scan.step(&d.scan, ':')
													d.scan.step(&d.scan, '"')
													d.scan.step(&d.scan, '"')
													d.scan.step(&d.scan, '}')
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:353
			// _ = "end of CoverTab[184473]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:354
			_go_fuzz_dep_.CoverTab[184474]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:354
			// _ = "end of CoverTab[184474]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:354
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:354
		// _ = "end of CoverTab[184466]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:354
		_go_fuzz_dep_.CoverTab[184467]++

												return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:356
		// _ = "end of CoverTab[184467]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:357
		_go_fuzz_dep_.CoverTab[184475]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:357
		// _ = "end of CoverTab[184475]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:357
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:357
	// _ = "end of CoverTab[184462]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:357
	_go_fuzz_dep_.CoverTab[184463]++

											switch op := d.scanWhile(scanSkipSpace); op {
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:360
		_go_fuzz_dep_.CoverTab[184476]++
												d.error(errPhase)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:361
		// _ = "end of CoverTab[184476]"

	case scanBeginArray:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:363
		_go_fuzz_dep_.CoverTab[184477]++
												d.array(v)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:364
		// _ = "end of CoverTab[184477]"

	case scanBeginObject:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:366
		_go_fuzz_dep_.CoverTab[184478]++
												d.object(v)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:367
		// _ = "end of CoverTab[184478]"

	case scanBeginLiteral:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:369
		_go_fuzz_dep_.CoverTab[184479]++
												d.literal(v)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:370
		// _ = "end of CoverTab[184479]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:371
	// _ = "end of CoverTab[184463]"
}

type unquotedValue struct{}

// valueQuoted is like value but decodes a
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:376
// quoted string literal or literal null into an interface value.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:376
// If it finds anything other than a quoted string literal or null,
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:376
// valueQuoted returns unquotedValue{}.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:380
func (d *decodeState) valueQuoted() interface{} {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:380
	_go_fuzz_dep_.CoverTab[184480]++
											switch op := d.scanWhile(scanSkipSpace); op {
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:382
		_go_fuzz_dep_.CoverTab[184482]++
												d.error(errPhase)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:383
		// _ = "end of CoverTab[184482]"

	case scanBeginArray:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:385
		_go_fuzz_dep_.CoverTab[184483]++
												d.array(reflect.Value{})
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:386
		// _ = "end of CoverTab[184483]"

	case scanBeginObject:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:388
		_go_fuzz_dep_.CoverTab[184484]++
												d.object(reflect.Value{})
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:389
		// _ = "end of CoverTab[184484]"

	case scanBeginLiteral:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:391
		_go_fuzz_dep_.CoverTab[184485]++
												switch v := d.literalInterface().(type) {
		case nil, string:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:393
			_go_fuzz_dep_.CoverTab[184486]++
													return v
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:394
			// _ = "end of CoverTab[184486]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:395
		// _ = "end of CoverTab[184485]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:396
	// _ = "end of CoverTab[184480]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:396
	_go_fuzz_dep_.CoverTab[184481]++
											return unquotedValue{}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:397
	// _ = "end of CoverTab[184481]"
}

// indirect walks down v allocating pointers as needed,
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:400
// until it gets to a non-pointer.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:400
// if it encounters an Unmarshaler, indirect stops and returns that.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:400
// if decodingNull is true, indirect stops at the last pointer so it can be set to nil.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:404
func (d *decodeState) indirect(v reflect.Value, decodingNull bool) (Unmarshaler, encoding.TextUnmarshaler, reflect.Value) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:404
	_go_fuzz_dep_.CoverTab[184487]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:408
	if v.Kind() != reflect.Ptr && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:408
		_go_fuzz_dep_.CoverTab[184490]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:408
		return v.Type().Name() != ""
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:408
		// _ = "end of CoverTab[184490]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:408
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:408
		_go_fuzz_dep_.CoverTab[184491]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:408
		return v.CanAddr()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:408
		// _ = "end of CoverTab[184491]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:408
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:408
		_go_fuzz_dep_.CoverTab[184492]++
												v = v.Addr()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:409
		// _ = "end of CoverTab[184492]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:410
		_go_fuzz_dep_.CoverTab[184493]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:410
		// _ = "end of CoverTab[184493]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:410
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:410
	// _ = "end of CoverTab[184487]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:410
	_go_fuzz_dep_.CoverTab[184488]++
											for {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:411
		_go_fuzz_dep_.CoverTab[184494]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:414
		if v.Kind() == reflect.Interface && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:414
			_go_fuzz_dep_.CoverTab[184500]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:414
			return !v.IsNil()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:414
			// _ = "end of CoverTab[184500]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:414
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:414
			_go_fuzz_dep_.CoverTab[184501]++
													e := v.Elem()
													if e.Kind() == reflect.Ptr && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:416
				_go_fuzz_dep_.CoverTab[184502]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:416
				return !e.IsNil()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:416
				// _ = "end of CoverTab[184502]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:416
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:416
				_go_fuzz_dep_.CoverTab[184503]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:416
				return (!decodingNull || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:416
					_go_fuzz_dep_.CoverTab[184504]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:416
					return e.Elem().Kind() == reflect.Ptr
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:416
					// _ = "end of CoverTab[184504]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:416
				}())
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:416
				// _ = "end of CoverTab[184503]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:416
			}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:416
				_go_fuzz_dep_.CoverTab[184505]++
														v = e
														continue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:418
				// _ = "end of CoverTab[184505]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:419
				_go_fuzz_dep_.CoverTab[184506]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:419
				// _ = "end of CoverTab[184506]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:419
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:419
			// _ = "end of CoverTab[184501]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:420
			_go_fuzz_dep_.CoverTab[184507]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:420
			// _ = "end of CoverTab[184507]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:420
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:420
		// _ = "end of CoverTab[184494]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:420
		_go_fuzz_dep_.CoverTab[184495]++

												if v.Kind() != reflect.Ptr {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:422
			_go_fuzz_dep_.CoverTab[184508]++
													break
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:423
			// _ = "end of CoverTab[184508]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:424
			_go_fuzz_dep_.CoverTab[184509]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:424
			// _ = "end of CoverTab[184509]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:424
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:424
		// _ = "end of CoverTab[184495]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:424
		_go_fuzz_dep_.CoverTab[184496]++

												if v.Elem().Kind() != reflect.Ptr && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:426
			_go_fuzz_dep_.CoverTab[184510]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:426
			return decodingNull
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:426
			// _ = "end of CoverTab[184510]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:426
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:426
			_go_fuzz_dep_.CoverTab[184511]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:426
			return v.CanSet()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:426
			// _ = "end of CoverTab[184511]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:426
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:426
			_go_fuzz_dep_.CoverTab[184512]++
													break
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:427
			// _ = "end of CoverTab[184512]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:428
			_go_fuzz_dep_.CoverTab[184513]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:428
			// _ = "end of CoverTab[184513]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:428
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:428
		// _ = "end of CoverTab[184496]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:428
		_go_fuzz_dep_.CoverTab[184497]++
												if v.IsNil() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:429
			_go_fuzz_dep_.CoverTab[184514]++
													v.Set(reflect.New(v.Type().Elem()))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:430
			// _ = "end of CoverTab[184514]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:431
			_go_fuzz_dep_.CoverTab[184515]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:431
			// _ = "end of CoverTab[184515]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:431
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:431
		// _ = "end of CoverTab[184497]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:431
		_go_fuzz_dep_.CoverTab[184498]++
												if v.Type().NumMethod() > 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:432
			_go_fuzz_dep_.CoverTab[184516]++
													if u, ok := v.Interface().(Unmarshaler); ok {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:433
				_go_fuzz_dep_.CoverTab[184518]++
														return u, nil, reflect.Value{}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:434
				// _ = "end of CoverTab[184518]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:435
				_go_fuzz_dep_.CoverTab[184519]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:435
				// _ = "end of CoverTab[184519]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:435
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:435
			// _ = "end of CoverTab[184516]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:435
			_go_fuzz_dep_.CoverTab[184517]++
													if u, ok := v.Interface().(encoding.TextUnmarshaler); ok {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:436
				_go_fuzz_dep_.CoverTab[184520]++
														return nil, u, reflect.Value{}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:437
				// _ = "end of CoverTab[184520]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:438
				_go_fuzz_dep_.CoverTab[184521]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:438
				// _ = "end of CoverTab[184521]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:438
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:438
			// _ = "end of CoverTab[184517]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:439
			_go_fuzz_dep_.CoverTab[184522]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:439
			// _ = "end of CoverTab[184522]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:439
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:439
		// _ = "end of CoverTab[184498]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:439
		_go_fuzz_dep_.CoverTab[184499]++
												v = v.Elem()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:440
		// _ = "end of CoverTab[184499]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:441
	// _ = "end of CoverTab[184488]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:441
	_go_fuzz_dep_.CoverTab[184489]++
											return nil, nil, v
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:442
	// _ = "end of CoverTab[184489]"
}

// array consumes an array from d.data[d.off-1:], decoding into the value v.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:445
// the first byte of the array ('[') has been read already.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:447
func (d *decodeState) array(v reflect.Value) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:447
	_go_fuzz_dep_.CoverTab[184523]++

											u, ut, pv := d.indirect(v, false)
											if u != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:450
		_go_fuzz_dep_.CoverTab[184529]++
												d.off--
												err := u.UnmarshalJSON(d.next())
												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:453
			_go_fuzz_dep_.CoverTab[184531]++
													d.error(err)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:454
			// _ = "end of CoverTab[184531]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:455
			_go_fuzz_dep_.CoverTab[184532]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:455
			// _ = "end of CoverTab[184532]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:455
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:455
		// _ = "end of CoverTab[184529]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:455
		_go_fuzz_dep_.CoverTab[184530]++
												return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:456
		// _ = "end of CoverTab[184530]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:457
		_go_fuzz_dep_.CoverTab[184533]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:457
		// _ = "end of CoverTab[184533]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:457
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:457
	// _ = "end of CoverTab[184523]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:457
	_go_fuzz_dep_.CoverTab[184524]++
											if ut != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:458
		_go_fuzz_dep_.CoverTab[184534]++
												d.saveError(&UnmarshalTypeError{"array", v.Type(), int64(d.off)})
												d.off--
												d.next()
												return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:462
		// _ = "end of CoverTab[184534]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:463
		_go_fuzz_dep_.CoverTab[184535]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:463
		// _ = "end of CoverTab[184535]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:463
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:463
	// _ = "end of CoverTab[184524]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:463
	_go_fuzz_dep_.CoverTab[184525]++

											v = pv

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:468
	switch v.Kind() {
	case reflect.Interface:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:469
		_go_fuzz_dep_.CoverTab[184536]++
												if v.NumMethod() == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:470
			_go_fuzz_dep_.CoverTab[184541]++

													v.Set(reflect.ValueOf(d.arrayInterface()))
													return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:473
			// _ = "end of CoverTab[184541]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:474
			_go_fuzz_dep_.CoverTab[184542]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:474
			// _ = "end of CoverTab[184542]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:474
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:474
		// _ = "end of CoverTab[184536]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:474
		_go_fuzz_dep_.CoverTab[184537]++

												fallthrough
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:476
		// _ = "end of CoverTab[184537]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:477
		_go_fuzz_dep_.CoverTab[184538]++
												d.saveError(&UnmarshalTypeError{"array", v.Type(), int64(d.off)})
												d.off--
												d.next()
												return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:481
		// _ = "end of CoverTab[184538]"
	case reflect.Array:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:482
		_go_fuzz_dep_.CoverTab[184539]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:482
		// _ = "end of CoverTab[184539]"
	case reflect.Slice:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:483
		_go_fuzz_dep_.CoverTab[184540]++
												break
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:484
		// _ = "end of CoverTab[184540]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:485
	// _ = "end of CoverTab[184525]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:485
	_go_fuzz_dep_.CoverTab[184526]++

											i := 0
											for {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:488
		_go_fuzz_dep_.CoverTab[184543]++

												op := d.scanWhile(scanSkipSpace)
												if op == scanEndArray {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:491
			_go_fuzz_dep_.CoverTab[184548]++
													break
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:492
			// _ = "end of CoverTab[184548]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:493
			_go_fuzz_dep_.CoverTab[184549]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:493
			// _ = "end of CoverTab[184549]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:493
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:493
		// _ = "end of CoverTab[184543]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:493
		_go_fuzz_dep_.CoverTab[184544]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:496
		d.off--
												d.scan.undo(op)

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:500
		if v.Kind() == reflect.Slice {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:500
			_go_fuzz_dep_.CoverTab[184550]++

													if i >= v.Cap() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:502
				_go_fuzz_dep_.CoverTab[184552]++
														newcap := v.Cap() + v.Cap()/2
														if newcap < 4 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:504
					_go_fuzz_dep_.CoverTab[184554]++
															newcap = 4
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:505
					// _ = "end of CoverTab[184554]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:506
					_go_fuzz_dep_.CoverTab[184555]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:506
					// _ = "end of CoverTab[184555]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:506
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:506
				// _ = "end of CoverTab[184552]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:506
				_go_fuzz_dep_.CoverTab[184553]++
														newv := reflect.MakeSlice(v.Type(), v.Len(), newcap)
														reflect.Copy(newv, v)
														v.Set(newv)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:509
				// _ = "end of CoverTab[184553]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:510
				_go_fuzz_dep_.CoverTab[184556]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:510
				// _ = "end of CoverTab[184556]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:510
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:510
			// _ = "end of CoverTab[184550]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:510
			_go_fuzz_dep_.CoverTab[184551]++
													if i >= v.Len() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:511
				_go_fuzz_dep_.CoverTab[184557]++
														v.SetLen(i + 1)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:512
				// _ = "end of CoverTab[184557]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:513
				_go_fuzz_dep_.CoverTab[184558]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:513
				// _ = "end of CoverTab[184558]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:513
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:513
			// _ = "end of CoverTab[184551]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:514
			_go_fuzz_dep_.CoverTab[184559]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:514
			// _ = "end of CoverTab[184559]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:514
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:514
		// _ = "end of CoverTab[184544]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:514
		_go_fuzz_dep_.CoverTab[184545]++

												if i < v.Len() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:516
			_go_fuzz_dep_.CoverTab[184560]++

													d.value(v.Index(i))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:518
			// _ = "end of CoverTab[184560]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:519
			_go_fuzz_dep_.CoverTab[184561]++

													d.value(reflect.Value{})
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:521
			// _ = "end of CoverTab[184561]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:522
		// _ = "end of CoverTab[184545]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:522
		_go_fuzz_dep_.CoverTab[184546]++
												i++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:526
		op = d.scanWhile(scanSkipSpace)
		if op == scanEndArray {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:527
			_go_fuzz_dep_.CoverTab[184562]++
													break
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:528
			// _ = "end of CoverTab[184562]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:529
			_go_fuzz_dep_.CoverTab[184563]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:529
			// _ = "end of CoverTab[184563]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:529
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:529
		// _ = "end of CoverTab[184546]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:529
		_go_fuzz_dep_.CoverTab[184547]++
												if op != scanArrayValue {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:530
			_go_fuzz_dep_.CoverTab[184564]++
													d.error(errPhase)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:531
			// _ = "end of CoverTab[184564]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:532
			_go_fuzz_dep_.CoverTab[184565]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:532
			// _ = "end of CoverTab[184565]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:532
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:532
		// _ = "end of CoverTab[184547]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:533
	// _ = "end of CoverTab[184526]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:533
	_go_fuzz_dep_.CoverTab[184527]++

											if i < v.Len() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:535
		_go_fuzz_dep_.CoverTab[184566]++
												if v.Kind() == reflect.Array {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:536
			_go_fuzz_dep_.CoverTab[184567]++

													z := reflect.Zero(v.Type().Elem())
													for ; i < v.Len(); i++ {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:539
				_go_fuzz_dep_.CoverTab[184568]++
														v.Index(i).Set(z)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:540
				// _ = "end of CoverTab[184568]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:541
			// _ = "end of CoverTab[184567]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:542
			_go_fuzz_dep_.CoverTab[184569]++
													v.SetLen(i)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:543
			// _ = "end of CoverTab[184569]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:544
		// _ = "end of CoverTab[184566]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:545
		_go_fuzz_dep_.CoverTab[184570]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:545
		// _ = "end of CoverTab[184570]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:545
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:545
	// _ = "end of CoverTab[184527]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:545
	_go_fuzz_dep_.CoverTab[184528]++
											if i == 0 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:546
		_go_fuzz_dep_.CoverTab[184571]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:546
		return v.Kind() == reflect.Slice
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:546
		// _ = "end of CoverTab[184571]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:546
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:546
		_go_fuzz_dep_.CoverTab[184572]++
												v.Set(reflect.MakeSlice(v.Type(), 0, 0))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:547
		// _ = "end of CoverTab[184572]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:548
		_go_fuzz_dep_.CoverTab[184573]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:548
		// _ = "end of CoverTab[184573]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:548
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:548
	// _ = "end of CoverTab[184528]"
}

var nullLiteral = []byte("null")

// object consumes an object from d.data[d.off-1:], decoding into the value v.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:553
// the first byte ('{') of the object has been read already.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:555
func (d *decodeState) object(v reflect.Value) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:555
	_go_fuzz_dep_.CoverTab[184574]++

											u, ut, pv := d.indirect(v, false)
											if u != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:558
		_go_fuzz_dep_.CoverTab[184579]++
												d.off--
												err := u.UnmarshalJSON(d.next())
												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:561
			_go_fuzz_dep_.CoverTab[184581]++
													d.error(err)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:562
			// _ = "end of CoverTab[184581]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:563
			_go_fuzz_dep_.CoverTab[184582]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:563
			// _ = "end of CoverTab[184582]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:563
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:563
		// _ = "end of CoverTab[184579]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:563
		_go_fuzz_dep_.CoverTab[184580]++
												return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:564
		// _ = "end of CoverTab[184580]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:565
		_go_fuzz_dep_.CoverTab[184583]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:565
		// _ = "end of CoverTab[184583]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:565
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:565
	// _ = "end of CoverTab[184574]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:565
	_go_fuzz_dep_.CoverTab[184575]++
											if ut != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:566
		_go_fuzz_dep_.CoverTab[184584]++
												d.saveError(&UnmarshalTypeError{"object", v.Type(), int64(d.off)})
												d.off--
												d.next()
												return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:570
		// _ = "end of CoverTab[184584]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:571
		_go_fuzz_dep_.CoverTab[184585]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:571
		// _ = "end of CoverTab[184585]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:571
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:571
	// _ = "end of CoverTab[184575]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:571
	_go_fuzz_dep_.CoverTab[184576]++
											v = pv

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:575
	if v.Kind() == reflect.Interface && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:575
		_go_fuzz_dep_.CoverTab[184586]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:575
		return v.NumMethod() == 0
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:575
		// _ = "end of CoverTab[184586]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:575
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:575
		_go_fuzz_dep_.CoverTab[184587]++
												v.Set(reflect.ValueOf(d.objectInterface()))
												return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:577
		// _ = "end of CoverTab[184587]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:578
		_go_fuzz_dep_.CoverTab[184588]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:578
		// _ = "end of CoverTab[184588]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:578
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:578
	// _ = "end of CoverTab[184576]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:578
	_go_fuzz_dep_.CoverTab[184577]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:581
	switch v.Kind() {
	case reflect.Map:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:582
		_go_fuzz_dep_.CoverTab[184589]++

												t := v.Type()
												if t.Key().Kind() != reflect.String {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:585
			_go_fuzz_dep_.CoverTab[184593]++
													d.saveError(&UnmarshalTypeError{"object", v.Type(), int64(d.off)})
													d.off--
													d.next()
													return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:589
			// _ = "end of CoverTab[184593]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:590
			_go_fuzz_dep_.CoverTab[184594]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:590
			// _ = "end of CoverTab[184594]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:590
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:590
		// _ = "end of CoverTab[184589]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:590
		_go_fuzz_dep_.CoverTab[184590]++
												if v.IsNil() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:591
			_go_fuzz_dep_.CoverTab[184595]++
													v.Set(reflect.MakeMap(t))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:592
			// _ = "end of CoverTab[184595]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:593
			_go_fuzz_dep_.CoverTab[184596]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:593
			// _ = "end of CoverTab[184596]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:593
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:593
		// _ = "end of CoverTab[184590]"
	case reflect.Struct:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:594
		_go_fuzz_dep_.CoverTab[184591]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:594
		// _ = "end of CoverTab[184591]"

	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:596
		_go_fuzz_dep_.CoverTab[184592]++
												d.saveError(&UnmarshalTypeError{"object", v.Type(), int64(d.off)})
												d.off--
												d.next()
												return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:600
		// _ = "end of CoverTab[184592]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:601
	// _ = "end of CoverTab[184577]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:601
	_go_fuzz_dep_.CoverTab[184578]++

											var mapElem reflect.Value
											keys := map[string]bool{}

											for {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:606
		_go_fuzz_dep_.CoverTab[184597]++

												op := d.scanWhile(scanSkipSpace)
												if op == scanEndObject {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:609
			_go_fuzz_dep_.CoverTab[184608]++

													break
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:611
			// _ = "end of CoverTab[184608]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:612
			_go_fuzz_dep_.CoverTab[184609]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:612
			// _ = "end of CoverTab[184609]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:612
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:612
		// _ = "end of CoverTab[184597]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:612
		_go_fuzz_dep_.CoverTab[184598]++
												if op != scanBeginLiteral {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:613
			_go_fuzz_dep_.CoverTab[184610]++
													d.error(errPhase)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:614
			// _ = "end of CoverTab[184610]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:615
			_go_fuzz_dep_.CoverTab[184611]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:615
			// _ = "end of CoverTab[184611]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:615
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:615
		// _ = "end of CoverTab[184598]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:615
		_go_fuzz_dep_.CoverTab[184599]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:618
		start := d.off - 1
		op = d.scanWhile(scanContinue)
		item := d.data[start : d.off-1]
		key, ok := unquote(item)
		if !ok {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:622
			_go_fuzz_dep_.CoverTab[184612]++
													d.error(errPhase)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:623
			// _ = "end of CoverTab[184612]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:624
			_go_fuzz_dep_.CoverTab[184613]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:624
			// _ = "end of CoverTab[184613]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:624
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:624
		// _ = "end of CoverTab[184599]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:624
		_go_fuzz_dep_.CoverTab[184600]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:627
		_, ok = keys[key]
		if !ok {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:628
			_go_fuzz_dep_.CoverTab[184614]++
													keys[key] = true
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:629
			// _ = "end of CoverTab[184614]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:630
			_go_fuzz_dep_.CoverTab[184615]++
													d.error(fmt.Errorf("json: duplicate key '%s' in object", key))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:631
			// _ = "end of CoverTab[184615]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:632
		// _ = "end of CoverTab[184600]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:632
		_go_fuzz_dep_.CoverTab[184601]++

		// Figure out field corresponding to key.
		var subv reflect.Value
		destring := false

		if v.Kind() == reflect.Map {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:638
			_go_fuzz_dep_.CoverTab[184616]++
													elemType := v.Type().Elem()
													if !mapElem.IsValid() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:640
				_go_fuzz_dep_.CoverTab[184618]++
														mapElem = reflect.New(elemType).Elem()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:641
				// _ = "end of CoverTab[184618]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:642
				_go_fuzz_dep_.CoverTab[184619]++
														mapElem.Set(reflect.Zero(elemType))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:643
				// _ = "end of CoverTab[184619]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:644
			// _ = "end of CoverTab[184616]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:644
			_go_fuzz_dep_.CoverTab[184617]++
													subv = mapElem
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:645
			// _ = "end of CoverTab[184617]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:646
			_go_fuzz_dep_.CoverTab[184620]++
													var f *field
													fields := cachedTypeFields(v.Type())
													for i := range fields {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:649
				_go_fuzz_dep_.CoverTab[184622]++
														ff := &fields[i]
														if bytes.Equal(ff.nameBytes, []byte(key)) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:651
					_go_fuzz_dep_.CoverTab[184623]++
															f = ff
															break
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:653
					// _ = "end of CoverTab[184623]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:654
					_go_fuzz_dep_.CoverTab[184624]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:654
					// _ = "end of CoverTab[184624]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:654
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:654
				// _ = "end of CoverTab[184622]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:655
			// _ = "end of CoverTab[184620]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:655
			_go_fuzz_dep_.CoverTab[184621]++
													if f != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:656
				_go_fuzz_dep_.CoverTab[184625]++
														subv = v
														destring = f.quoted
														for _, i := range f.index {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:659
					_go_fuzz_dep_.CoverTab[184626]++
															if subv.Kind() == reflect.Ptr {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:660
						_go_fuzz_dep_.CoverTab[184628]++
																if subv.IsNil() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:661
							_go_fuzz_dep_.CoverTab[184630]++
																	subv.Set(reflect.New(subv.Type().Elem()))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:662
							// _ = "end of CoverTab[184630]"
						} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:663
							_go_fuzz_dep_.CoverTab[184631]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:663
							// _ = "end of CoverTab[184631]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:663
						}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:663
						// _ = "end of CoverTab[184628]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:663
						_go_fuzz_dep_.CoverTab[184629]++
																subv = subv.Elem()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:664
						// _ = "end of CoverTab[184629]"
					} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:665
						_go_fuzz_dep_.CoverTab[184632]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:665
						// _ = "end of CoverTab[184632]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:665
					}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:665
					// _ = "end of CoverTab[184626]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:665
					_go_fuzz_dep_.CoverTab[184627]++
															subv = subv.Field(i)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:666
					// _ = "end of CoverTab[184627]"
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:667
				// _ = "end of CoverTab[184625]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:668
				_go_fuzz_dep_.CoverTab[184633]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:668
				// _ = "end of CoverTab[184633]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:668
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:668
			// _ = "end of CoverTab[184621]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:669
		// _ = "end of CoverTab[184601]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:669
		_go_fuzz_dep_.CoverTab[184602]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:672
		if op == scanSkipSpace {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:672
			_go_fuzz_dep_.CoverTab[184634]++
													op = d.scanWhile(scanSkipSpace)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:673
			// _ = "end of CoverTab[184634]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:674
			_go_fuzz_dep_.CoverTab[184635]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:674
			// _ = "end of CoverTab[184635]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:674
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:674
		// _ = "end of CoverTab[184602]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:674
		_go_fuzz_dep_.CoverTab[184603]++
												if op != scanObjectKey {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:675
			_go_fuzz_dep_.CoverTab[184636]++
													d.error(errPhase)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:676
			// _ = "end of CoverTab[184636]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:677
			_go_fuzz_dep_.CoverTab[184637]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:677
			// _ = "end of CoverTab[184637]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:677
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:677
		// _ = "end of CoverTab[184603]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:677
		_go_fuzz_dep_.CoverTab[184604]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:680
		if destring {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:680
			_go_fuzz_dep_.CoverTab[184638]++
													switch qv := d.valueQuoted().(type) {
			case nil:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:682
				_go_fuzz_dep_.CoverTab[184639]++
														d.literalStore(nullLiteral, subv, false)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:683
				// _ = "end of CoverTab[184639]"
			case string:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:684
				_go_fuzz_dep_.CoverTab[184640]++
														d.literalStore([]byte(qv), subv, true)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:685
				// _ = "end of CoverTab[184640]"
			default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:686
				_go_fuzz_dep_.CoverTab[184641]++
														d.saveError(fmt.Errorf("json: invalid use of ,string struct tag, trying to unmarshal unquoted value into %v", subv.Type()))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:687
				// _ = "end of CoverTab[184641]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:688
			// _ = "end of CoverTab[184638]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:689
			_go_fuzz_dep_.CoverTab[184642]++
													d.value(subv)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:690
			// _ = "end of CoverTab[184642]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:691
		// _ = "end of CoverTab[184604]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:691
		_go_fuzz_dep_.CoverTab[184605]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:695
		if v.Kind() == reflect.Map {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:695
			_go_fuzz_dep_.CoverTab[184643]++
													kv := reflect.ValueOf(key).Convert(v.Type().Key())
													v.SetMapIndex(kv, subv)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:697
			// _ = "end of CoverTab[184643]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:698
			_go_fuzz_dep_.CoverTab[184644]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:698
			// _ = "end of CoverTab[184644]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:698
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:698
		// _ = "end of CoverTab[184605]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:698
		_go_fuzz_dep_.CoverTab[184606]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:701
		op = d.scanWhile(scanSkipSpace)
		if op == scanEndObject {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:702
			_go_fuzz_dep_.CoverTab[184645]++
													break
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:703
			// _ = "end of CoverTab[184645]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:704
			_go_fuzz_dep_.CoverTab[184646]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:704
			// _ = "end of CoverTab[184646]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:704
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:704
		// _ = "end of CoverTab[184606]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:704
		_go_fuzz_dep_.CoverTab[184607]++
												if op != scanObjectValue {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:705
			_go_fuzz_dep_.CoverTab[184647]++
													d.error(errPhase)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:706
			// _ = "end of CoverTab[184647]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:707
			_go_fuzz_dep_.CoverTab[184648]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:707
			// _ = "end of CoverTab[184648]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:707
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:707
		// _ = "end of CoverTab[184607]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:708
	// _ = "end of CoverTab[184578]"
}

// literal consumes a literal from d.data[d.off-1:], decoding into the value v.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:711
// The first byte of the literal has been read already
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:711
// (that's how the caller knows it's a literal).
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:714
func (d *decodeState) literal(v reflect.Value) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:714
	_go_fuzz_dep_.CoverTab[184649]++

											start := d.off - 1
											op := d.scanWhile(scanContinue)

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:720
	d.off--
											d.scan.undo(op)

											d.literalStore(d.data[start:d.off], v, false)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:723
	// _ = "end of CoverTab[184649]"
}

// convertNumber converts the number literal s to a float64 or a Number
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:726
// depending on the setting of d.useNumber.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:728
func (d *decodeState) convertNumber(s string) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:728
	_go_fuzz_dep_.CoverTab[184650]++
											if d.useNumber {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:729
		_go_fuzz_dep_.CoverTab[184653]++
												return Number(s), nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:730
		// _ = "end of CoverTab[184653]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:731
		_go_fuzz_dep_.CoverTab[184654]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:731
		// _ = "end of CoverTab[184654]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:731
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:731
	// _ = "end of CoverTab[184650]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:731
	_go_fuzz_dep_.CoverTab[184651]++
											f, err := strconv.ParseFloat(s, 64)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:733
		_go_fuzz_dep_.CoverTab[184655]++
												return nil, &UnmarshalTypeError{"number " + s, reflect.TypeOf(0.0), int64(d.off)}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:734
		// _ = "end of CoverTab[184655]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:735
		_go_fuzz_dep_.CoverTab[184656]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:735
		// _ = "end of CoverTab[184656]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:735
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:735
	// _ = "end of CoverTab[184651]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:735
	_go_fuzz_dep_.CoverTab[184652]++
											return f, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:736
	// _ = "end of CoverTab[184652]"
}

var numberType = reflect.TypeOf(Number(""))

// literalStore decodes a literal stored in item into v.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:741
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:741
// fromQuoted indicates whether this literal came from unwrapping a
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:741
// string from the ",string" struct tag option. this is used only to
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:741
// produce more helpful error messages.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:746
func (d *decodeState) literalStore(item []byte, v reflect.Value, fromQuoted bool) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:746
	_go_fuzz_dep_.CoverTab[184657]++

											if len(item) == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:748
		_go_fuzz_dep_.CoverTab[184661]++

												d.saveError(fmt.Errorf("json: invalid use of ,string struct tag, trying to unmarshal %q into %v", item, v.Type()))
												return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:751
		// _ = "end of CoverTab[184661]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:752
		_go_fuzz_dep_.CoverTab[184662]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:752
		// _ = "end of CoverTab[184662]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:752
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:752
	// _ = "end of CoverTab[184657]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:752
	_go_fuzz_dep_.CoverTab[184658]++
											wantptr := item[0] == 'n'
											u, ut, pv := d.indirect(v, wantptr)
											if u != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:755
		_go_fuzz_dep_.CoverTab[184663]++
												err := u.UnmarshalJSON(item)
												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:757
			_go_fuzz_dep_.CoverTab[184665]++
													d.error(err)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:758
			// _ = "end of CoverTab[184665]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:759
			_go_fuzz_dep_.CoverTab[184666]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:759
			// _ = "end of CoverTab[184666]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:759
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:759
		// _ = "end of CoverTab[184663]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:759
		_go_fuzz_dep_.CoverTab[184664]++
												return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:760
		// _ = "end of CoverTab[184664]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:761
		_go_fuzz_dep_.CoverTab[184667]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:761
		// _ = "end of CoverTab[184667]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:761
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:761
	// _ = "end of CoverTab[184658]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:761
	_go_fuzz_dep_.CoverTab[184659]++
											if ut != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:762
		_go_fuzz_dep_.CoverTab[184668]++
												if item[0] != '"' {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:763
			_go_fuzz_dep_.CoverTab[184672]++
													if fromQuoted {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:764
				_go_fuzz_dep_.CoverTab[184674]++
														d.saveError(fmt.Errorf("json: invalid use of ,string struct tag, trying to unmarshal %q into %v", item, v.Type()))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:765
				// _ = "end of CoverTab[184674]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:766
				_go_fuzz_dep_.CoverTab[184675]++
														d.saveError(&UnmarshalTypeError{"string", v.Type(), int64(d.off)})
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:767
				// _ = "end of CoverTab[184675]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:768
			// _ = "end of CoverTab[184672]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:768
			_go_fuzz_dep_.CoverTab[184673]++
													return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:769
			// _ = "end of CoverTab[184673]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:770
			_go_fuzz_dep_.CoverTab[184676]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:770
			// _ = "end of CoverTab[184676]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:770
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:770
		// _ = "end of CoverTab[184668]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:770
		_go_fuzz_dep_.CoverTab[184669]++
												s, ok := unquoteBytes(item)
												if !ok {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:772
			_go_fuzz_dep_.CoverTab[184677]++
													if fromQuoted {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:773
				_go_fuzz_dep_.CoverTab[184678]++
														d.error(fmt.Errorf("json: invalid use of ,string struct tag, trying to unmarshal %q into %v", item, v.Type()))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:774
				// _ = "end of CoverTab[184678]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:775
				_go_fuzz_dep_.CoverTab[184679]++
														d.error(errPhase)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:776
				// _ = "end of CoverTab[184679]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:777
			// _ = "end of CoverTab[184677]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:778
			_go_fuzz_dep_.CoverTab[184680]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:778
			// _ = "end of CoverTab[184680]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:778
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:778
		// _ = "end of CoverTab[184669]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:778
		_go_fuzz_dep_.CoverTab[184670]++
												err := ut.UnmarshalText(s)
												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:780
			_go_fuzz_dep_.CoverTab[184681]++
													d.error(err)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:781
			// _ = "end of CoverTab[184681]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:782
			_go_fuzz_dep_.CoverTab[184682]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:782
			// _ = "end of CoverTab[184682]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:782
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:782
		// _ = "end of CoverTab[184670]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:782
		_go_fuzz_dep_.CoverTab[184671]++
												return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:783
		// _ = "end of CoverTab[184671]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:784
		_go_fuzz_dep_.CoverTab[184683]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:784
		// _ = "end of CoverTab[184683]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:784
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:784
	// _ = "end of CoverTab[184659]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:784
	_go_fuzz_dep_.CoverTab[184660]++

											v = pv

											switch c := item[0]; c {
	case 'n':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:789
		_go_fuzz_dep_.CoverTab[184684]++
												switch v.Kind() {
		case reflect.Interface, reflect.Ptr, reflect.Map, reflect.Slice:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:791
			_go_fuzz_dep_.CoverTab[184690]++
													v.Set(reflect.Zero(v.Type()))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:792
			// _ = "end of CoverTab[184690]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:792
		default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:792
			_go_fuzz_dep_.CoverTab[184691]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:792
			// _ = "end of CoverTab[184691]"

		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:794
		// _ = "end of CoverTab[184684]"
	case 't', 'f':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:795
		_go_fuzz_dep_.CoverTab[184685]++
												value := c == 't'
												switch v.Kind() {
		default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:798
			_go_fuzz_dep_.CoverTab[184692]++
													if fromQuoted {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:799
				_go_fuzz_dep_.CoverTab[184695]++
														d.saveError(fmt.Errorf("json: invalid use of ,string struct tag, trying to unmarshal %q into %v", item, v.Type()))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:800
				// _ = "end of CoverTab[184695]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:801
				_go_fuzz_dep_.CoverTab[184696]++
														d.saveError(&UnmarshalTypeError{"bool", v.Type(), int64(d.off)})
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:802
				// _ = "end of CoverTab[184696]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:803
			// _ = "end of CoverTab[184692]"
		case reflect.Bool:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:804
			_go_fuzz_dep_.CoverTab[184693]++
													v.SetBool(value)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:805
			// _ = "end of CoverTab[184693]"
		case reflect.Interface:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:806
			_go_fuzz_dep_.CoverTab[184694]++
													if v.NumMethod() == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:807
				_go_fuzz_dep_.CoverTab[184697]++
														v.Set(reflect.ValueOf(value))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:808
				// _ = "end of CoverTab[184697]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:809
				_go_fuzz_dep_.CoverTab[184698]++
														d.saveError(&UnmarshalTypeError{"bool", v.Type(), int64(d.off)})
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:810
				// _ = "end of CoverTab[184698]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:811
			// _ = "end of CoverTab[184694]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:812
		// _ = "end of CoverTab[184685]"

	case '"':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:814
		_go_fuzz_dep_.CoverTab[184686]++
												s, ok := unquoteBytes(item)
												if !ok {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:816
			_go_fuzz_dep_.CoverTab[184699]++
													if fromQuoted {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:817
				_go_fuzz_dep_.CoverTab[184700]++
														d.error(fmt.Errorf("json: invalid use of ,string struct tag, trying to unmarshal %q into %v", item, v.Type()))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:818
				// _ = "end of CoverTab[184700]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:819
				_go_fuzz_dep_.CoverTab[184701]++
														d.error(errPhase)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:820
				// _ = "end of CoverTab[184701]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:821
			// _ = "end of CoverTab[184699]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:822
			_go_fuzz_dep_.CoverTab[184702]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:822
			// _ = "end of CoverTab[184702]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:822
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:822
		// _ = "end of CoverTab[184686]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:822
		_go_fuzz_dep_.CoverTab[184687]++
												switch v.Kind() {
		default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:824
			_go_fuzz_dep_.CoverTab[184703]++
													d.saveError(&UnmarshalTypeError{"string", v.Type(), int64(d.off)})
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:825
			// _ = "end of CoverTab[184703]"
		case reflect.Slice:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:826
			_go_fuzz_dep_.CoverTab[184704]++
													if v.Type().Elem().Kind() != reflect.Uint8 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:827
				_go_fuzz_dep_.CoverTab[184709]++
														d.saveError(&UnmarshalTypeError{"string", v.Type(), int64(d.off)})
														break
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:829
				// _ = "end of CoverTab[184709]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:830
				_go_fuzz_dep_.CoverTab[184710]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:830
				// _ = "end of CoverTab[184710]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:830
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:830
			// _ = "end of CoverTab[184704]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:830
			_go_fuzz_dep_.CoverTab[184705]++
													b := make([]byte, base64.StdEncoding.DecodedLen(len(s)))
													n, err := base64.StdEncoding.Decode(b, s)
													if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:833
				_go_fuzz_dep_.CoverTab[184711]++
														d.saveError(err)
														break
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:835
				// _ = "end of CoverTab[184711]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:836
				_go_fuzz_dep_.CoverTab[184712]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:836
				// _ = "end of CoverTab[184712]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:836
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:836
			// _ = "end of CoverTab[184705]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:836
			_go_fuzz_dep_.CoverTab[184706]++
													v.SetBytes(b[:n])
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:837
			// _ = "end of CoverTab[184706]"
		case reflect.String:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:838
			_go_fuzz_dep_.CoverTab[184707]++
													v.SetString(string(s))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:839
			// _ = "end of CoverTab[184707]"
		case reflect.Interface:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:840
			_go_fuzz_dep_.CoverTab[184708]++
													if v.NumMethod() == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:841
				_go_fuzz_dep_.CoverTab[184713]++
														v.Set(reflect.ValueOf(string(s)))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:842
				// _ = "end of CoverTab[184713]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:843
				_go_fuzz_dep_.CoverTab[184714]++
														d.saveError(&UnmarshalTypeError{"string", v.Type(), int64(d.off)})
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:844
				// _ = "end of CoverTab[184714]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:845
			// _ = "end of CoverTab[184708]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:846
		// _ = "end of CoverTab[184687]"

	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:848
		_go_fuzz_dep_.CoverTab[184688]++
												if c != '-' && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:849
			_go_fuzz_dep_.CoverTab[184715]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:849
			return (c < '0' || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:849
				_go_fuzz_dep_.CoverTab[184716]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:849
				return c > '9'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:849
				// _ = "end of CoverTab[184716]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:849
			}())
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:849
			// _ = "end of CoverTab[184715]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:849
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:849
			_go_fuzz_dep_.CoverTab[184717]++
													if fromQuoted {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:850
				_go_fuzz_dep_.CoverTab[184718]++
														d.error(fmt.Errorf("json: invalid use of ,string struct tag, trying to unmarshal %q into %v", item, v.Type()))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:851
				// _ = "end of CoverTab[184718]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:852
				_go_fuzz_dep_.CoverTab[184719]++
														d.error(errPhase)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:853
				// _ = "end of CoverTab[184719]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:854
			// _ = "end of CoverTab[184717]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:855
			_go_fuzz_dep_.CoverTab[184720]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:855
			// _ = "end of CoverTab[184720]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:855
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:855
		// _ = "end of CoverTab[184688]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:855
		_go_fuzz_dep_.CoverTab[184689]++
												s := string(item)
												switch v.Kind() {
		default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:858
			_go_fuzz_dep_.CoverTab[184721]++
													if v.Kind() == reflect.String && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:859
				_go_fuzz_dep_.CoverTab[184732]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:859
				return v.Type() == numberType
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:859
				// _ = "end of CoverTab[184732]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:859
			}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:859
				_go_fuzz_dep_.CoverTab[184733]++
														v.SetString(s)
														if !isValidNumber(s) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:861
					_go_fuzz_dep_.CoverTab[184735]++
															d.error(fmt.Errorf("json: invalid number literal, trying to unmarshal %q into Number", item))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:862
					// _ = "end of CoverTab[184735]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:863
					_go_fuzz_dep_.CoverTab[184736]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:863
					// _ = "end of CoverTab[184736]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:863
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:863
				// _ = "end of CoverTab[184733]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:863
				_go_fuzz_dep_.CoverTab[184734]++
														break
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:864
				// _ = "end of CoverTab[184734]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:865
				_go_fuzz_dep_.CoverTab[184737]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:865
				// _ = "end of CoverTab[184737]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:865
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:865
			// _ = "end of CoverTab[184721]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:865
			_go_fuzz_dep_.CoverTab[184722]++
													if fromQuoted {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:866
				_go_fuzz_dep_.CoverTab[184738]++
														d.error(fmt.Errorf("json: invalid use of ,string struct tag, trying to unmarshal %q into %v", item, v.Type()))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:867
				// _ = "end of CoverTab[184738]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:868
				_go_fuzz_dep_.CoverTab[184739]++
														d.error(&UnmarshalTypeError{"number", v.Type(), int64(d.off)})
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:869
				// _ = "end of CoverTab[184739]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:870
			// _ = "end of CoverTab[184722]"
		case reflect.Interface:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:871
			_go_fuzz_dep_.CoverTab[184723]++
													n, err := d.convertNumber(s)
													if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:873
				_go_fuzz_dep_.CoverTab[184740]++
														d.saveError(err)
														break
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:875
				// _ = "end of CoverTab[184740]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:876
				_go_fuzz_dep_.CoverTab[184741]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:876
				// _ = "end of CoverTab[184741]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:876
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:876
			// _ = "end of CoverTab[184723]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:876
			_go_fuzz_dep_.CoverTab[184724]++
													if v.NumMethod() != 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:877
				_go_fuzz_dep_.CoverTab[184742]++
														d.saveError(&UnmarshalTypeError{"number", v.Type(), int64(d.off)})
														break
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:879
				// _ = "end of CoverTab[184742]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:880
				_go_fuzz_dep_.CoverTab[184743]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:880
				// _ = "end of CoverTab[184743]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:880
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:880
			// _ = "end of CoverTab[184724]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:880
			_go_fuzz_dep_.CoverTab[184725]++
													v.Set(reflect.ValueOf(n))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:881
			// _ = "end of CoverTab[184725]"

		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:883
			_go_fuzz_dep_.CoverTab[184726]++
													n, err := strconv.ParseInt(s, 10, 64)
													if err != nil || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:885
				_go_fuzz_dep_.CoverTab[184744]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:885
				return v.OverflowInt(n)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:885
				// _ = "end of CoverTab[184744]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:885
			}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:885
				_go_fuzz_dep_.CoverTab[184745]++
														d.saveError(&UnmarshalTypeError{"number " + s, v.Type(), int64(d.off)})
														break
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:887
				// _ = "end of CoverTab[184745]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:888
				_go_fuzz_dep_.CoverTab[184746]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:888
				// _ = "end of CoverTab[184746]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:888
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:888
			// _ = "end of CoverTab[184726]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:888
			_go_fuzz_dep_.CoverTab[184727]++
													v.SetInt(n)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:889
			// _ = "end of CoverTab[184727]"

		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:891
			_go_fuzz_dep_.CoverTab[184728]++
													n, err := strconv.ParseUint(s, 10, 64)
													if err != nil || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:893
				_go_fuzz_dep_.CoverTab[184747]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:893
				return v.OverflowUint(n)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:893
				// _ = "end of CoverTab[184747]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:893
			}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:893
				_go_fuzz_dep_.CoverTab[184748]++
														d.saveError(&UnmarshalTypeError{"number " + s, v.Type(), int64(d.off)})
														break
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:895
				// _ = "end of CoverTab[184748]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:896
				_go_fuzz_dep_.CoverTab[184749]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:896
				// _ = "end of CoverTab[184749]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:896
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:896
			// _ = "end of CoverTab[184728]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:896
			_go_fuzz_dep_.CoverTab[184729]++
													v.SetUint(n)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:897
			// _ = "end of CoverTab[184729]"

		case reflect.Float32, reflect.Float64:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:899
			_go_fuzz_dep_.CoverTab[184730]++
													n, err := strconv.ParseFloat(s, v.Type().Bits())
													if err != nil || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:901
				_go_fuzz_dep_.CoverTab[184750]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:901
				return v.OverflowFloat(n)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:901
				// _ = "end of CoverTab[184750]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:901
			}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:901
				_go_fuzz_dep_.CoverTab[184751]++
														d.saveError(&UnmarshalTypeError{"number " + s, v.Type(), int64(d.off)})
														break
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:903
				// _ = "end of CoverTab[184751]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:904
				_go_fuzz_dep_.CoverTab[184752]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:904
				// _ = "end of CoverTab[184752]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:904
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:904
			// _ = "end of CoverTab[184730]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:904
			_go_fuzz_dep_.CoverTab[184731]++
													v.SetFloat(n)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:905
			// _ = "end of CoverTab[184731]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:906
		// _ = "end of CoverTab[184689]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:907
	// _ = "end of CoverTab[184660]"
}

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:914
// valueInterface is like value but returns interface{}
func (d *decodeState) valueInterface() interface{} {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:915
	_go_fuzz_dep_.CoverTab[184753]++
											switch d.scanWhile(scanSkipSpace) {
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:917
		_go_fuzz_dep_.CoverTab[184754]++
												d.error(errPhase)
												panic("unreachable")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:919
		// _ = "end of CoverTab[184754]"
	case scanBeginArray:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:920
		_go_fuzz_dep_.CoverTab[184755]++
												return d.arrayInterface()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:921
		// _ = "end of CoverTab[184755]"
	case scanBeginObject:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:922
		_go_fuzz_dep_.CoverTab[184756]++
												return d.objectInterface()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:923
		// _ = "end of CoverTab[184756]"
	case scanBeginLiteral:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:924
		_go_fuzz_dep_.CoverTab[184757]++
												return d.literalInterface()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:925
		// _ = "end of CoverTab[184757]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:926
	// _ = "end of CoverTab[184753]"
}

// arrayInterface is like array but returns []interface{}.
func (d *decodeState) arrayInterface() []interface{} {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:930
	_go_fuzz_dep_.CoverTab[184758]++
											var v = make([]interface{}, 0)
											for {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:932
		_go_fuzz_dep_.CoverTab[184760]++

												op := d.scanWhile(scanSkipSpace)
												if op == scanEndArray {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:935
			_go_fuzz_dep_.CoverTab[184763]++
													break
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:936
			// _ = "end of CoverTab[184763]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:937
			_go_fuzz_dep_.CoverTab[184764]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:937
			// _ = "end of CoverTab[184764]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:937
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:937
		// _ = "end of CoverTab[184760]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:937
		_go_fuzz_dep_.CoverTab[184761]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:940
		d.off--
												d.scan.undo(op)

												v = append(v, d.valueInterface())

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:946
		op = d.scanWhile(scanSkipSpace)
		if op == scanEndArray {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:947
			_go_fuzz_dep_.CoverTab[184765]++
													break
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:948
			// _ = "end of CoverTab[184765]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:949
			_go_fuzz_dep_.CoverTab[184766]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:949
			// _ = "end of CoverTab[184766]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:949
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:949
		// _ = "end of CoverTab[184761]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:949
		_go_fuzz_dep_.CoverTab[184762]++
												if op != scanArrayValue {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:950
			_go_fuzz_dep_.CoverTab[184767]++
													d.error(errPhase)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:951
			// _ = "end of CoverTab[184767]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:952
			_go_fuzz_dep_.CoverTab[184768]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:952
			// _ = "end of CoverTab[184768]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:952
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:952
		// _ = "end of CoverTab[184762]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:953
	// _ = "end of CoverTab[184758]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:953
	_go_fuzz_dep_.CoverTab[184759]++
											return v
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:954
	// _ = "end of CoverTab[184759]"
}

// objectInterface is like object but returns map[string]interface{}.
func (d *decodeState) objectInterface() map[string]interface{} {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:958
	_go_fuzz_dep_.CoverTab[184769]++
											m := make(map[string]interface{})
											keys := map[string]bool{}

											for {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:962
		_go_fuzz_dep_.CoverTab[184771]++

												op := d.scanWhile(scanSkipSpace)
												if op == scanEndObject {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:965
			_go_fuzz_dep_.CoverTab[184779]++

													break
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:967
			// _ = "end of CoverTab[184779]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:968
			_go_fuzz_dep_.CoverTab[184780]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:968
			// _ = "end of CoverTab[184780]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:968
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:968
		// _ = "end of CoverTab[184771]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:968
		_go_fuzz_dep_.CoverTab[184772]++
												if op != scanBeginLiteral {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:969
			_go_fuzz_dep_.CoverTab[184781]++
													d.error(errPhase)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:970
			// _ = "end of CoverTab[184781]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:971
			_go_fuzz_dep_.CoverTab[184782]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:971
			// _ = "end of CoverTab[184782]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:971
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:971
		// _ = "end of CoverTab[184772]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:971
		_go_fuzz_dep_.CoverTab[184773]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:974
		start := d.off - 1
		op = d.scanWhile(scanContinue)
		item := d.data[start : d.off-1]
		key, ok := unquote(item)
		if !ok {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:978
			_go_fuzz_dep_.CoverTab[184783]++
													d.error(errPhase)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:979
			// _ = "end of CoverTab[184783]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:980
			_go_fuzz_dep_.CoverTab[184784]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:980
			// _ = "end of CoverTab[184784]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:980
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:980
		// _ = "end of CoverTab[184773]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:980
		_go_fuzz_dep_.CoverTab[184774]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:983
		_, ok = keys[key]
		if !ok {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:984
			_go_fuzz_dep_.CoverTab[184785]++
													keys[key] = true
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:985
			// _ = "end of CoverTab[184785]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:986
			_go_fuzz_dep_.CoverTab[184786]++
													d.error(fmt.Errorf("json: duplicate key '%s' in object", key))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:987
			// _ = "end of CoverTab[184786]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:988
		// _ = "end of CoverTab[184774]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:988
		_go_fuzz_dep_.CoverTab[184775]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:991
		if op == scanSkipSpace {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:991
			_go_fuzz_dep_.CoverTab[184787]++
													op = d.scanWhile(scanSkipSpace)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:992
			// _ = "end of CoverTab[184787]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:993
			_go_fuzz_dep_.CoverTab[184788]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:993
			// _ = "end of CoverTab[184788]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:993
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:993
		// _ = "end of CoverTab[184775]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:993
		_go_fuzz_dep_.CoverTab[184776]++
												if op != scanObjectKey {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:994
			_go_fuzz_dep_.CoverTab[184789]++
													d.error(errPhase)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:995
			// _ = "end of CoverTab[184789]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:996
			_go_fuzz_dep_.CoverTab[184790]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:996
			// _ = "end of CoverTab[184790]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:996
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:996
		// _ = "end of CoverTab[184776]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:996
		_go_fuzz_dep_.CoverTab[184777]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:999
		m[key] = d.valueInterface()

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1002
		op = d.scanWhile(scanSkipSpace)
		if op == scanEndObject {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1003
			_go_fuzz_dep_.CoverTab[184791]++
														break
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1004
			// _ = "end of CoverTab[184791]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1005
			_go_fuzz_dep_.CoverTab[184792]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1005
			// _ = "end of CoverTab[184792]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1005
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1005
		// _ = "end of CoverTab[184777]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1005
		_go_fuzz_dep_.CoverTab[184778]++
													if op != scanObjectValue {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1006
			_go_fuzz_dep_.CoverTab[184793]++
														d.error(errPhase)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1007
			// _ = "end of CoverTab[184793]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1008
			_go_fuzz_dep_.CoverTab[184794]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1008
			// _ = "end of CoverTab[184794]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1008
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1008
		// _ = "end of CoverTab[184778]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1009
	// _ = "end of CoverTab[184769]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1009
	_go_fuzz_dep_.CoverTab[184770]++
												return m
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1010
	// _ = "end of CoverTab[184770]"
}

// literalInterface is like literal but returns an interface value.
func (d *decodeState) literalInterface() interface{} {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1014
	_go_fuzz_dep_.CoverTab[184795]++

												start := d.off - 1
												op := d.scanWhile(scanContinue)

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1020
	d.off--
	d.scan.undo(op)
	item := d.data[start:d.off]

	switch c := item[0]; c {
	case 'n':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1025
		_go_fuzz_dep_.CoverTab[184796]++
													return nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1026
		// _ = "end of CoverTab[184796]"

	case 't', 'f':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1028
		_go_fuzz_dep_.CoverTab[184797]++
													return c == 't'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1029
		// _ = "end of CoverTab[184797]"

	case '"':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1031
		_go_fuzz_dep_.CoverTab[184798]++
													s, ok := unquote(item)
													if !ok {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1033
			_go_fuzz_dep_.CoverTab[184803]++
														d.error(errPhase)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1034
			// _ = "end of CoverTab[184803]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1035
			_go_fuzz_dep_.CoverTab[184804]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1035
			// _ = "end of CoverTab[184804]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1035
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1035
		// _ = "end of CoverTab[184798]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1035
		_go_fuzz_dep_.CoverTab[184799]++
													return s
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1036
		// _ = "end of CoverTab[184799]"

	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1038
		_go_fuzz_dep_.CoverTab[184800]++
													if c != '-' && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1039
			_go_fuzz_dep_.CoverTab[184805]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1039
			return (c < '0' || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1039
				_go_fuzz_dep_.CoverTab[184806]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1039
				return c > '9'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1039
				// _ = "end of CoverTab[184806]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1039
			}())
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1039
			// _ = "end of CoverTab[184805]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1039
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1039
			_go_fuzz_dep_.CoverTab[184807]++
														d.error(errPhase)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1040
			// _ = "end of CoverTab[184807]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1041
			_go_fuzz_dep_.CoverTab[184808]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1041
			// _ = "end of CoverTab[184808]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1041
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1041
		// _ = "end of CoverTab[184800]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1041
		_go_fuzz_dep_.CoverTab[184801]++
													n, err := d.convertNumber(string(item))
													if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1043
			_go_fuzz_dep_.CoverTab[184809]++
														d.saveError(err)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1044
			// _ = "end of CoverTab[184809]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1045
			_go_fuzz_dep_.CoverTab[184810]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1045
			// _ = "end of CoverTab[184810]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1045
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1045
		// _ = "end of CoverTab[184801]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1045
		_go_fuzz_dep_.CoverTab[184802]++
													return n
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1046
		// _ = "end of CoverTab[184802]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1047
	// _ = "end of CoverTab[184795]"
}

// getu4 decodes \uXXXX from the beginning of s, returning the hex value,
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1050
// or it returns -1.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1052
func getu4(s []byte) rune {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1052
	_go_fuzz_dep_.CoverTab[184811]++
												if len(s) < 6 || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1053
		_go_fuzz_dep_.CoverTab[184814]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1053
		return s[0] != '\\'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1053
		// _ = "end of CoverTab[184814]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1053
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1053
		_go_fuzz_dep_.CoverTab[184815]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1053
		return s[1] != 'u'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1053
		// _ = "end of CoverTab[184815]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1053
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1053
		_go_fuzz_dep_.CoverTab[184816]++
													return -1
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1054
		// _ = "end of CoverTab[184816]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1055
		_go_fuzz_dep_.CoverTab[184817]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1055
		// _ = "end of CoverTab[184817]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1055
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1055
	// _ = "end of CoverTab[184811]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1055
	_go_fuzz_dep_.CoverTab[184812]++
												r, err := strconv.ParseUint(string(s[2:6]), 16, 64)
												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1057
		_go_fuzz_dep_.CoverTab[184818]++
													return -1
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1058
		// _ = "end of CoverTab[184818]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1059
		_go_fuzz_dep_.CoverTab[184819]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1059
		// _ = "end of CoverTab[184819]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1059
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1059
	// _ = "end of CoverTab[184812]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1059
	_go_fuzz_dep_.CoverTab[184813]++
												return rune(r)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1060
	// _ = "end of CoverTab[184813]"
}

// unquote converts a quoted JSON string literal s into an actual string t.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1063
// The rules are different than for Go, so cannot use strconv.Unquote.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1065
func unquote(s []byte) (t string, ok bool) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1065
	_go_fuzz_dep_.CoverTab[184820]++
												s, ok = unquoteBytes(s)
												t = string(s)
												return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1068
	// _ = "end of CoverTab[184820]"
}

func unquoteBytes(s []byte) (t []byte, ok bool) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1071
	_go_fuzz_dep_.CoverTab[184821]++
												if len(s) < 2 || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1072
		_go_fuzz_dep_.CoverTab[184826]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1072
		return s[0] != '"'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1072
		// _ = "end of CoverTab[184826]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1072
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1072
		_go_fuzz_dep_.CoverTab[184827]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1072
		return s[len(s)-1] != '"'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1072
		// _ = "end of CoverTab[184827]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1072
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1072
		_go_fuzz_dep_.CoverTab[184828]++
													return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1073
		// _ = "end of CoverTab[184828]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1074
		_go_fuzz_dep_.CoverTab[184829]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1074
		// _ = "end of CoverTab[184829]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1074
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1074
	// _ = "end of CoverTab[184821]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1074
	_go_fuzz_dep_.CoverTab[184822]++
												s = s[1 : len(s)-1]

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1080
	r := 0
	for r < len(s) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1081
		_go_fuzz_dep_.CoverTab[184830]++
													c := s[r]
													if c == '\\' || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1083
			_go_fuzz_dep_.CoverTab[184834]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1083
			return c == '"'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1083
			// _ = "end of CoverTab[184834]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1083
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1083
			_go_fuzz_dep_.CoverTab[184835]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1083
			return c < ' '
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1083
			// _ = "end of CoverTab[184835]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1083
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1083
			_go_fuzz_dep_.CoverTab[184836]++
														break
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1084
			// _ = "end of CoverTab[184836]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1085
			_go_fuzz_dep_.CoverTab[184837]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1085
			// _ = "end of CoverTab[184837]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1085
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1085
		// _ = "end of CoverTab[184830]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1085
		_go_fuzz_dep_.CoverTab[184831]++
													if c < utf8.RuneSelf {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1086
			_go_fuzz_dep_.CoverTab[184838]++
														r++
														continue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1088
			// _ = "end of CoverTab[184838]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1089
			_go_fuzz_dep_.CoverTab[184839]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1089
			// _ = "end of CoverTab[184839]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1089
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1089
		// _ = "end of CoverTab[184831]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1089
		_go_fuzz_dep_.CoverTab[184832]++
													rr, size := utf8.DecodeRune(s[r:])
													if rr == utf8.RuneError && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1091
			_go_fuzz_dep_.CoverTab[184840]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1091
			return size == 1
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1091
			// _ = "end of CoverTab[184840]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1091
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1091
			_go_fuzz_dep_.CoverTab[184841]++
														break
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1092
			// _ = "end of CoverTab[184841]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1093
			_go_fuzz_dep_.CoverTab[184842]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1093
			// _ = "end of CoverTab[184842]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1093
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1093
		// _ = "end of CoverTab[184832]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1093
		_go_fuzz_dep_.CoverTab[184833]++
													r += size
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1094
		// _ = "end of CoverTab[184833]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1095
	// _ = "end of CoverTab[184822]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1095
	_go_fuzz_dep_.CoverTab[184823]++
												if r == len(s) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1096
		_go_fuzz_dep_.CoverTab[184843]++
													return s, true
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1097
		// _ = "end of CoverTab[184843]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1098
		_go_fuzz_dep_.CoverTab[184844]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1098
		// _ = "end of CoverTab[184844]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1098
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1098
	// _ = "end of CoverTab[184823]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1098
	_go_fuzz_dep_.CoverTab[184824]++

												b := make([]byte, len(s)+2*utf8.UTFMax)
												w := copy(b, s[0:r])
												for r < len(s) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1102
		_go_fuzz_dep_.CoverTab[184845]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1106
		if w >= len(b)-2*utf8.UTFMax {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1106
			_go_fuzz_dep_.CoverTab[184847]++
														nb := make([]byte, (len(b)+utf8.UTFMax)*2)
														copy(nb, b[0:w])
														b = nb
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1109
			// _ = "end of CoverTab[184847]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1110
			_go_fuzz_dep_.CoverTab[184848]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1110
			// _ = "end of CoverTab[184848]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1110
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1110
		// _ = "end of CoverTab[184845]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1110
		_go_fuzz_dep_.CoverTab[184846]++
													switch c := s[r]; {
		case c == '\\':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1112
			_go_fuzz_dep_.CoverTab[184849]++
														r++
														if r >= len(s) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1114
				_go_fuzz_dep_.CoverTab[184854]++
															return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1115
				// _ = "end of CoverTab[184854]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1116
				_go_fuzz_dep_.CoverTab[184855]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1116
				// _ = "end of CoverTab[184855]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1116
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1116
			// _ = "end of CoverTab[184849]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1116
			_go_fuzz_dep_.CoverTab[184850]++
														switch s[r] {
			default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1118
				_go_fuzz_dep_.CoverTab[184856]++
															return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1119
				// _ = "end of CoverTab[184856]"
			case '"', '\\', '/', '\'':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1120
				_go_fuzz_dep_.CoverTab[184857]++
															b[w] = s[r]
															r++
															w++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1123
				// _ = "end of CoverTab[184857]"
			case 'b':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1124
				_go_fuzz_dep_.CoverTab[184858]++
															b[w] = '\b'
															r++
															w++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1127
				// _ = "end of CoverTab[184858]"
			case 'f':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1128
				_go_fuzz_dep_.CoverTab[184859]++
															b[w] = '\f'
															r++
															w++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1131
				// _ = "end of CoverTab[184859]"
			case 'n':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1132
				_go_fuzz_dep_.CoverTab[184860]++
															b[w] = '\n'
															r++
															w++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1135
				// _ = "end of CoverTab[184860]"
			case 'r':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1136
				_go_fuzz_dep_.CoverTab[184861]++
															b[w] = '\r'
															r++
															w++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1139
				// _ = "end of CoverTab[184861]"
			case 't':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1140
				_go_fuzz_dep_.CoverTab[184862]++
															b[w] = '\t'
															r++
															w++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1143
				// _ = "end of CoverTab[184862]"
			case 'u':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1144
				_go_fuzz_dep_.CoverTab[184863]++
															r--
															rr := getu4(s[r:])
															if rr < 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1147
					_go_fuzz_dep_.CoverTab[184866]++
																return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1148
					// _ = "end of CoverTab[184866]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1149
					_go_fuzz_dep_.CoverTab[184867]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1149
					// _ = "end of CoverTab[184867]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1149
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1149
				// _ = "end of CoverTab[184863]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1149
				_go_fuzz_dep_.CoverTab[184864]++
															r += 6
															if utf16.IsSurrogate(rr) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1151
					_go_fuzz_dep_.CoverTab[184868]++
																rr1 := getu4(s[r:])
																if dec := utf16.DecodeRune(rr, rr1); dec != unicode.ReplacementChar {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1153
						_go_fuzz_dep_.CoverTab[184870]++

																	r += 6
																	w += utf8.EncodeRune(b[w:], dec)
																	break
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1157
						// _ = "end of CoverTab[184870]"
					} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1158
						_go_fuzz_dep_.CoverTab[184871]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1158
						// _ = "end of CoverTab[184871]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1158
					}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1158
					// _ = "end of CoverTab[184868]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1158
					_go_fuzz_dep_.CoverTab[184869]++

																rr = unicode.ReplacementChar
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1160
					// _ = "end of CoverTab[184869]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1161
					_go_fuzz_dep_.CoverTab[184872]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1161
					// _ = "end of CoverTab[184872]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1161
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1161
				// _ = "end of CoverTab[184864]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1161
				_go_fuzz_dep_.CoverTab[184865]++
															w += utf8.EncodeRune(b[w:], rr)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1162
				// _ = "end of CoverTab[184865]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1163
			// _ = "end of CoverTab[184850]"

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1166
		case c == '"', c < ' ':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1166
			_go_fuzz_dep_.CoverTab[184851]++
														return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1167
			// _ = "end of CoverTab[184851]"

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1170
		case c < utf8.RuneSelf:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1170
			_go_fuzz_dep_.CoverTab[184852]++
														b[w] = c
														r++
														w++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1173
			// _ = "end of CoverTab[184852]"

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1176
		default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1176
			_go_fuzz_dep_.CoverTab[184853]++
														rr, size := utf8.DecodeRune(s[r:])
														r += size
														w += utf8.EncodeRune(b[w:], rr)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1179
			// _ = "end of CoverTab[184853]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1180
		// _ = "end of CoverTab[184846]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1181
	// _ = "end of CoverTab[184824]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1181
	_go_fuzz_dep_.CoverTab[184825]++
												return b[0:w], true
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1182
	// _ = "end of CoverTab[184825]"
}

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1183
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/decode.go:1183
var _ = _go_fuzz_dep_.CoverTab
