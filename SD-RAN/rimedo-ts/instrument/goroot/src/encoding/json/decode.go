// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Represents JSON data structure using native Go types: booleans, floats,
// strings, arrays, and maps.

//line /usr/local/go/src/encoding/json/decode.go:8
package json

//line /usr/local/go/src/encoding/json/decode.go:8
import (
//line /usr/local/go/src/encoding/json/decode.go:8
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/encoding/json/decode.go:8
)
//line /usr/local/go/src/encoding/json/decode.go:8
import (
//line /usr/local/go/src/encoding/json/decode.go:8
	_atomic_ "sync/atomic"
//line /usr/local/go/src/encoding/json/decode.go:8
)

import (
	"encoding"
	"encoding/base64"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf16"
	"unicode/utf8"
)

// Unmarshal parses the JSON-encoded data and stores the result
//line /usr/local/go/src/encoding/json/decode.go:22
// in the value pointed to by v. If v is nil or not a pointer,
//line /usr/local/go/src/encoding/json/decode.go:22
// Unmarshal returns an InvalidUnmarshalError.
//line /usr/local/go/src/encoding/json/decode.go:22
//
//line /usr/local/go/src/encoding/json/decode.go:22
// Unmarshal uses the inverse of the encodings that
//line /usr/local/go/src/encoding/json/decode.go:22
// Marshal uses, allocating maps, slices, and pointers as necessary,
//line /usr/local/go/src/encoding/json/decode.go:22
// with the following additional rules:
//line /usr/local/go/src/encoding/json/decode.go:22
//
//line /usr/local/go/src/encoding/json/decode.go:22
// To unmarshal JSON into a pointer, Unmarshal first handles the case of
//line /usr/local/go/src/encoding/json/decode.go:22
// the JSON being the JSON literal null. In that case, Unmarshal sets
//line /usr/local/go/src/encoding/json/decode.go:22
// the pointer to nil. Otherwise, Unmarshal unmarshals the JSON into
//line /usr/local/go/src/encoding/json/decode.go:22
// the value pointed at by the pointer. If the pointer is nil, Unmarshal
//line /usr/local/go/src/encoding/json/decode.go:22
// allocates a new value for it to point to.
//line /usr/local/go/src/encoding/json/decode.go:22
//
//line /usr/local/go/src/encoding/json/decode.go:22
// To unmarshal JSON into a value implementing the Unmarshaler interface,
//line /usr/local/go/src/encoding/json/decode.go:22
// Unmarshal calls that value's UnmarshalJSON method, including
//line /usr/local/go/src/encoding/json/decode.go:22
// when the input is a JSON null.
//line /usr/local/go/src/encoding/json/decode.go:22
// Otherwise, if the value implements encoding.TextUnmarshaler
//line /usr/local/go/src/encoding/json/decode.go:22
// and the input is a JSON quoted string, Unmarshal calls that value's
//line /usr/local/go/src/encoding/json/decode.go:22
// UnmarshalText method with the unquoted form of the string.
//line /usr/local/go/src/encoding/json/decode.go:22
//
//line /usr/local/go/src/encoding/json/decode.go:22
// To unmarshal JSON into a struct, Unmarshal matches incoming object
//line /usr/local/go/src/encoding/json/decode.go:22
// keys to the keys used by Marshal (either the struct field name or its tag),
//line /usr/local/go/src/encoding/json/decode.go:22
// preferring an exact match but also accepting a case-insensitive match. By
//line /usr/local/go/src/encoding/json/decode.go:22
// default, object keys which don't have a corresponding struct field are
//line /usr/local/go/src/encoding/json/decode.go:22
// ignored (see Decoder.DisallowUnknownFields for an alternative).
//line /usr/local/go/src/encoding/json/decode.go:22
//
//line /usr/local/go/src/encoding/json/decode.go:22
// To unmarshal JSON into an interface value,
//line /usr/local/go/src/encoding/json/decode.go:22
// Unmarshal stores one of these in the interface value:
//line /usr/local/go/src/encoding/json/decode.go:22
//
//line /usr/local/go/src/encoding/json/decode.go:22
//	bool, for JSON booleans
//line /usr/local/go/src/encoding/json/decode.go:22
//	float64, for JSON numbers
//line /usr/local/go/src/encoding/json/decode.go:22
//	string, for JSON strings
//line /usr/local/go/src/encoding/json/decode.go:22
//	[]interface{}, for JSON arrays
//line /usr/local/go/src/encoding/json/decode.go:22
//	map[string]interface{}, for JSON objects
//line /usr/local/go/src/encoding/json/decode.go:22
//	nil for JSON null
//line /usr/local/go/src/encoding/json/decode.go:22
//
//line /usr/local/go/src/encoding/json/decode.go:22
// To unmarshal a JSON array into a slice, Unmarshal resets the slice length
//line /usr/local/go/src/encoding/json/decode.go:22
// to zero and then appends each element to the slice.
//line /usr/local/go/src/encoding/json/decode.go:22
// As a special case, to unmarshal an empty JSON array into a slice,
//line /usr/local/go/src/encoding/json/decode.go:22
// Unmarshal replaces the slice with a new empty slice.
//line /usr/local/go/src/encoding/json/decode.go:22
//
//line /usr/local/go/src/encoding/json/decode.go:22
// To unmarshal a JSON array into a Go array, Unmarshal decodes
//line /usr/local/go/src/encoding/json/decode.go:22
// JSON array elements into corresponding Go array elements.
//line /usr/local/go/src/encoding/json/decode.go:22
// If the Go array is smaller than the JSON array,
//line /usr/local/go/src/encoding/json/decode.go:22
// the additional JSON array elements are discarded.
//line /usr/local/go/src/encoding/json/decode.go:22
// If the JSON array is smaller than the Go array,
//line /usr/local/go/src/encoding/json/decode.go:22
// the additional Go array elements are set to zero values.
//line /usr/local/go/src/encoding/json/decode.go:22
//
//line /usr/local/go/src/encoding/json/decode.go:22
// To unmarshal a JSON object into a map, Unmarshal first establishes a map to
//line /usr/local/go/src/encoding/json/decode.go:22
// use. If the map is nil, Unmarshal allocates a new map. Otherwise Unmarshal
//line /usr/local/go/src/encoding/json/decode.go:22
// reuses the existing map, keeping existing entries. Unmarshal then stores
//line /usr/local/go/src/encoding/json/decode.go:22
// key-value pairs from the JSON object into the map. The map's key type must
//line /usr/local/go/src/encoding/json/decode.go:22
// either be any string type, an integer, implement json.Unmarshaler, or
//line /usr/local/go/src/encoding/json/decode.go:22
// implement encoding.TextUnmarshaler.
//line /usr/local/go/src/encoding/json/decode.go:22
//
//line /usr/local/go/src/encoding/json/decode.go:22
// If the JSON-encoded data contain a syntax error, Unmarshal returns a SyntaxError.
//line /usr/local/go/src/encoding/json/decode.go:22
//
//line /usr/local/go/src/encoding/json/decode.go:22
// If a JSON value is not appropriate for a given target type,
//line /usr/local/go/src/encoding/json/decode.go:22
// or if a JSON number overflows the target type, Unmarshal
//line /usr/local/go/src/encoding/json/decode.go:22
// skips that field and completes the unmarshaling as best it can.
//line /usr/local/go/src/encoding/json/decode.go:22
// If no more serious errors are encountered, Unmarshal returns
//line /usr/local/go/src/encoding/json/decode.go:22
// an UnmarshalTypeError describing the earliest such error. In any
//line /usr/local/go/src/encoding/json/decode.go:22
// case, it's not guaranteed that all the remaining fields following
//line /usr/local/go/src/encoding/json/decode.go:22
// the problematic one will be unmarshaled into the target object.
//line /usr/local/go/src/encoding/json/decode.go:22
//
//line /usr/local/go/src/encoding/json/decode.go:22
// The JSON null value unmarshals into an interface, map, pointer, or slice
//line /usr/local/go/src/encoding/json/decode.go:22
// by setting that Go value to nil. Because null is often used in JSON to mean
//line /usr/local/go/src/encoding/json/decode.go:22
// “not present,” unmarshaling a JSON null into any other Go type has no effect
//line /usr/local/go/src/encoding/json/decode.go:22
// on the value and produces no error.
//line /usr/local/go/src/encoding/json/decode.go:22
//
//line /usr/local/go/src/encoding/json/decode.go:22
// When unmarshaling quoted strings, invalid UTF-8 or
//line /usr/local/go/src/encoding/json/decode.go:22
// invalid UTF-16 surrogate pairs are not treated as an error.
//line /usr/local/go/src/encoding/json/decode.go:22
// Instead, they are replaced by the Unicode replacement
//line /usr/local/go/src/encoding/json/decode.go:22
// character U+FFFD.
//line /usr/local/go/src/encoding/json/decode.go:97
func Unmarshal(data []byte, v any) error {
//line /usr/local/go/src/encoding/json/decode.go:97
	_go_fuzz_dep_.CoverTab[26854]++
	// Check for well-formedness.
	// Avoids filling out half a data structure
	// before discovering a JSON syntax error.
	var d decodeState
	err := checkValid(data, &d.scan)
	if err != nil {
//line /usr/local/go/src/encoding/json/decode.go:103
		_go_fuzz_dep_.CoverTab[26856]++
								return err
//line /usr/local/go/src/encoding/json/decode.go:104
		// _ = "end of CoverTab[26856]"
	} else {
//line /usr/local/go/src/encoding/json/decode.go:105
		_go_fuzz_dep_.CoverTab[26857]++
//line /usr/local/go/src/encoding/json/decode.go:105
		// _ = "end of CoverTab[26857]"
//line /usr/local/go/src/encoding/json/decode.go:105
	}
//line /usr/local/go/src/encoding/json/decode.go:105
	// _ = "end of CoverTab[26854]"
//line /usr/local/go/src/encoding/json/decode.go:105
	_go_fuzz_dep_.CoverTab[26855]++

							d.init(data)
							return d.unmarshal(v)
//line /usr/local/go/src/encoding/json/decode.go:108
	// _ = "end of CoverTab[26855]"
}

// Unmarshaler is the interface implemented by types
//line /usr/local/go/src/encoding/json/decode.go:111
// that can unmarshal a JSON description of themselves.
//line /usr/local/go/src/encoding/json/decode.go:111
// The input can be assumed to be a valid encoding of
//line /usr/local/go/src/encoding/json/decode.go:111
// a JSON value. UnmarshalJSON must copy the JSON data
//line /usr/local/go/src/encoding/json/decode.go:111
// if it wishes to retain the data after returning.
//line /usr/local/go/src/encoding/json/decode.go:111
//
//line /usr/local/go/src/encoding/json/decode.go:111
// By convention, to approximate the behavior of Unmarshal itself,
//line /usr/local/go/src/encoding/json/decode.go:111
// Unmarshalers implement UnmarshalJSON([]byte("null")) as a no-op.
//line /usr/local/go/src/encoding/json/decode.go:119
type Unmarshaler interface {
	UnmarshalJSON([]byte) error
}

// An UnmarshalTypeError describes a JSON value that was
//line /usr/local/go/src/encoding/json/decode.go:123
// not appropriate for a value of a specific Go type.
//line /usr/local/go/src/encoding/json/decode.go:125
type UnmarshalTypeError struct {
	Value	string		// description of JSON value - "bool", "array", "number -5"
	Type	reflect.Type	// type of Go value it could not be assigned to
	Offset	int64		// error occurred after reading Offset bytes
	Struct	string		// name of the struct type containing the field
	Field	string		// the full path from root node to the field
}

func (e *UnmarshalTypeError) Error() string {
//line /usr/local/go/src/encoding/json/decode.go:133
	_go_fuzz_dep_.CoverTab[26858]++
							if e.Struct != "" || func() bool {
//line /usr/local/go/src/encoding/json/decode.go:134
		_go_fuzz_dep_.CoverTab[26860]++
//line /usr/local/go/src/encoding/json/decode.go:134
		return e.Field != ""
//line /usr/local/go/src/encoding/json/decode.go:134
		// _ = "end of CoverTab[26860]"
//line /usr/local/go/src/encoding/json/decode.go:134
	}() {
//line /usr/local/go/src/encoding/json/decode.go:134
		_go_fuzz_dep_.CoverTab[26861]++
								return "json: cannot unmarshal " + e.Value + " into Go struct field " + e.Struct + "." + e.Field + " of type " + e.Type.String()
//line /usr/local/go/src/encoding/json/decode.go:135
		// _ = "end of CoverTab[26861]"
	} else {
//line /usr/local/go/src/encoding/json/decode.go:136
		_go_fuzz_dep_.CoverTab[26862]++
//line /usr/local/go/src/encoding/json/decode.go:136
		// _ = "end of CoverTab[26862]"
//line /usr/local/go/src/encoding/json/decode.go:136
	}
//line /usr/local/go/src/encoding/json/decode.go:136
	// _ = "end of CoverTab[26858]"
//line /usr/local/go/src/encoding/json/decode.go:136
	_go_fuzz_dep_.CoverTab[26859]++
							return "json: cannot unmarshal " + e.Value + " into Go value of type " + e.Type.String()
//line /usr/local/go/src/encoding/json/decode.go:137
	// _ = "end of CoverTab[26859]"
}

// An UnmarshalFieldError describes a JSON object key that
//line /usr/local/go/src/encoding/json/decode.go:140
// led to an unexported (and therefore unwritable) struct field.
//line /usr/local/go/src/encoding/json/decode.go:140
//
//line /usr/local/go/src/encoding/json/decode.go:140
// Deprecated: No longer used; kept for compatibility.
//line /usr/local/go/src/encoding/json/decode.go:144
type UnmarshalFieldError struct {
	Key	string
	Type	reflect.Type
	Field	reflect.StructField
}

func (e *UnmarshalFieldError) Error() string {
//line /usr/local/go/src/encoding/json/decode.go:150
	_go_fuzz_dep_.CoverTab[26863]++
							return "json: cannot unmarshal object key " + strconv.Quote(e.Key) + " into unexported field " + e.Field.Name + " of type " + e.Type.String()
//line /usr/local/go/src/encoding/json/decode.go:151
	// _ = "end of CoverTab[26863]"
}

// An InvalidUnmarshalError describes an invalid argument passed to Unmarshal.
//line /usr/local/go/src/encoding/json/decode.go:154
// (The argument to Unmarshal must be a non-nil pointer.)
//line /usr/local/go/src/encoding/json/decode.go:156
type InvalidUnmarshalError struct {
	Type reflect.Type
}

func (e *InvalidUnmarshalError) Error() string {
//line /usr/local/go/src/encoding/json/decode.go:160
	_go_fuzz_dep_.CoverTab[26864]++
							if e.Type == nil {
//line /usr/local/go/src/encoding/json/decode.go:161
		_go_fuzz_dep_.CoverTab[26867]++
								return "json: Unmarshal(nil)"
//line /usr/local/go/src/encoding/json/decode.go:162
		// _ = "end of CoverTab[26867]"
	} else {
//line /usr/local/go/src/encoding/json/decode.go:163
		_go_fuzz_dep_.CoverTab[26868]++
//line /usr/local/go/src/encoding/json/decode.go:163
		// _ = "end of CoverTab[26868]"
//line /usr/local/go/src/encoding/json/decode.go:163
	}
//line /usr/local/go/src/encoding/json/decode.go:163
	// _ = "end of CoverTab[26864]"
//line /usr/local/go/src/encoding/json/decode.go:163
	_go_fuzz_dep_.CoverTab[26865]++

							if e.Type.Kind() != reflect.Pointer {
//line /usr/local/go/src/encoding/json/decode.go:165
		_go_fuzz_dep_.CoverTab[26869]++
								return "json: Unmarshal(non-pointer " + e.Type.String() + ")"
//line /usr/local/go/src/encoding/json/decode.go:166
		// _ = "end of CoverTab[26869]"
	} else {
//line /usr/local/go/src/encoding/json/decode.go:167
		_go_fuzz_dep_.CoverTab[26870]++
//line /usr/local/go/src/encoding/json/decode.go:167
		// _ = "end of CoverTab[26870]"
//line /usr/local/go/src/encoding/json/decode.go:167
	}
//line /usr/local/go/src/encoding/json/decode.go:167
	// _ = "end of CoverTab[26865]"
//line /usr/local/go/src/encoding/json/decode.go:167
	_go_fuzz_dep_.CoverTab[26866]++
							return "json: Unmarshal(nil " + e.Type.String() + ")"
//line /usr/local/go/src/encoding/json/decode.go:168
	// _ = "end of CoverTab[26866]"
}

func (d *decodeState) unmarshal(v any) error {
//line /usr/local/go/src/encoding/json/decode.go:171
	_go_fuzz_dep_.CoverTab[26871]++
							rv := reflect.ValueOf(v)
							if rv.Kind() != reflect.Pointer || func() bool {
//line /usr/local/go/src/encoding/json/decode.go:173
		_go_fuzz_dep_.CoverTab[26874]++
//line /usr/local/go/src/encoding/json/decode.go:173
		return rv.IsNil()
//line /usr/local/go/src/encoding/json/decode.go:173
		// _ = "end of CoverTab[26874]"
//line /usr/local/go/src/encoding/json/decode.go:173
	}() {
//line /usr/local/go/src/encoding/json/decode.go:173
		_go_fuzz_dep_.CoverTab[26875]++
								return &InvalidUnmarshalError{reflect.TypeOf(v)}
//line /usr/local/go/src/encoding/json/decode.go:174
		// _ = "end of CoverTab[26875]"
	} else {
//line /usr/local/go/src/encoding/json/decode.go:175
		_go_fuzz_dep_.CoverTab[26876]++
//line /usr/local/go/src/encoding/json/decode.go:175
		// _ = "end of CoverTab[26876]"
//line /usr/local/go/src/encoding/json/decode.go:175
	}
//line /usr/local/go/src/encoding/json/decode.go:175
	// _ = "end of CoverTab[26871]"
//line /usr/local/go/src/encoding/json/decode.go:175
	_go_fuzz_dep_.CoverTab[26872]++

							d.scan.reset()
							d.scanWhile(scanSkipSpace)

//line /usr/local/go/src/encoding/json/decode.go:181
	err := d.value(rv)
	if err != nil {
//line /usr/local/go/src/encoding/json/decode.go:182
		_go_fuzz_dep_.CoverTab[26877]++
								return d.addErrorContext(err)
//line /usr/local/go/src/encoding/json/decode.go:183
		// _ = "end of CoverTab[26877]"
	} else {
//line /usr/local/go/src/encoding/json/decode.go:184
		_go_fuzz_dep_.CoverTab[26878]++
//line /usr/local/go/src/encoding/json/decode.go:184
		// _ = "end of CoverTab[26878]"
//line /usr/local/go/src/encoding/json/decode.go:184
	}
//line /usr/local/go/src/encoding/json/decode.go:184
	// _ = "end of CoverTab[26872]"
//line /usr/local/go/src/encoding/json/decode.go:184
	_go_fuzz_dep_.CoverTab[26873]++
							return d.savedError
//line /usr/local/go/src/encoding/json/decode.go:185
	// _ = "end of CoverTab[26873]"
}

// A Number represents a JSON number literal.
type Number string

// String returns the literal text of the number.
func (n Number) String() string {
//line /usr/local/go/src/encoding/json/decode.go:192
	_go_fuzz_dep_.CoverTab[26879]++
//line /usr/local/go/src/encoding/json/decode.go:192
	return string(n)
//line /usr/local/go/src/encoding/json/decode.go:192
	// _ = "end of CoverTab[26879]"
//line /usr/local/go/src/encoding/json/decode.go:192
}

// Float64 returns the number as a float64.
func (n Number) Float64() (float64, error) {
//line /usr/local/go/src/encoding/json/decode.go:195
	_go_fuzz_dep_.CoverTab[26880]++
							return strconv.ParseFloat(string(n), 64)
//line /usr/local/go/src/encoding/json/decode.go:196
	// _ = "end of CoverTab[26880]"
}

// Int64 returns the number as an int64.
func (n Number) Int64() (int64, error) {
//line /usr/local/go/src/encoding/json/decode.go:200
	_go_fuzz_dep_.CoverTab[26881]++
							return strconv.ParseInt(string(n), 10, 64)
//line /usr/local/go/src/encoding/json/decode.go:201
	// _ = "end of CoverTab[26881]"
}

// An errorContext provides context for type errors during decoding.
type errorContext struct {
	Struct		reflect.Type
	FieldStack	[]string
}

// decodeState represents the state while decoding a JSON value.
type decodeState struct {
	data			[]byte
	off			int	// next read offset in data
	opcode			int	// last read result
	scan			scanner
	errorContext		*errorContext
	savedError		error
	useNumber		bool
	disallowUnknownFields	bool
}

// readIndex returns the position of the last byte read.
func (d *decodeState) readIndex() int {
//line /usr/local/go/src/encoding/json/decode.go:223
	_go_fuzz_dep_.CoverTab[26882]++
							return d.off - 1
//line /usr/local/go/src/encoding/json/decode.go:224
	// _ = "end of CoverTab[26882]"
}

// phasePanicMsg is used as a panic message when we end up with something that
//line /usr/local/go/src/encoding/json/decode.go:227
// shouldn't happen. It can indicate a bug in the JSON decoder, or that
//line /usr/local/go/src/encoding/json/decode.go:227
// something is editing the data slice while the decoder executes.
//line /usr/local/go/src/encoding/json/decode.go:230
const phasePanicMsg = "JSON decoder out of sync - data changing underfoot?"

func (d *decodeState) init(data []byte) *decodeState {
	d.data = data
	d.off = 0
	d.savedError = nil
	if d.errorContext != nil {
		d.errorContext.Struct = nil

		d.errorContext.FieldStack = d.errorContext.FieldStack[:0]
	}
	return d
}

// saveError saves the first err it is called with,
//line /usr/local/go/src/encoding/json/decode.go:244
// for reporting at the end of the unmarshal.
//line /usr/local/go/src/encoding/json/decode.go:246
func (d *decodeState) saveError(err error) {
//line /usr/local/go/src/encoding/json/decode.go:246
	_go_fuzz_dep_.CoverTab[26883]++
							if d.savedError == nil {
//line /usr/local/go/src/encoding/json/decode.go:247
		_go_fuzz_dep_.CoverTab[26884]++
								d.savedError = d.addErrorContext(err)
//line /usr/local/go/src/encoding/json/decode.go:248
		// _ = "end of CoverTab[26884]"
	} else {
//line /usr/local/go/src/encoding/json/decode.go:249
		_go_fuzz_dep_.CoverTab[26885]++
//line /usr/local/go/src/encoding/json/decode.go:249
		// _ = "end of CoverTab[26885]"
//line /usr/local/go/src/encoding/json/decode.go:249
	}
//line /usr/local/go/src/encoding/json/decode.go:249
	// _ = "end of CoverTab[26883]"
}

// addErrorContext returns a new error enhanced with information from d.errorContext
func (d *decodeState) addErrorContext(err error) error {
//line /usr/local/go/src/encoding/json/decode.go:253
	_go_fuzz_dep_.CoverTab[26886]++
							if d.errorContext != nil && func() bool {
//line /usr/local/go/src/encoding/json/decode.go:254
		_go_fuzz_dep_.CoverTab[26888]++
//line /usr/local/go/src/encoding/json/decode.go:254
		return (d.errorContext.Struct != nil || func() bool {
//line /usr/local/go/src/encoding/json/decode.go:254
			_go_fuzz_dep_.CoverTab[26889]++
//line /usr/local/go/src/encoding/json/decode.go:254
			return len(d.errorContext.FieldStack) > 0
//line /usr/local/go/src/encoding/json/decode.go:254
			// _ = "end of CoverTab[26889]"
//line /usr/local/go/src/encoding/json/decode.go:254
		}())
//line /usr/local/go/src/encoding/json/decode.go:254
		// _ = "end of CoverTab[26888]"
//line /usr/local/go/src/encoding/json/decode.go:254
	}() {
//line /usr/local/go/src/encoding/json/decode.go:254
		_go_fuzz_dep_.CoverTab[26890]++
								switch err := err.(type) {
		case *UnmarshalTypeError:
//line /usr/local/go/src/encoding/json/decode.go:256
			_go_fuzz_dep_.CoverTab[26891]++
									err.Struct = d.errorContext.Struct.Name()
									err.Field = strings.Join(d.errorContext.FieldStack, ".")
//line /usr/local/go/src/encoding/json/decode.go:258
			// _ = "end of CoverTab[26891]"
		}
//line /usr/local/go/src/encoding/json/decode.go:259
		// _ = "end of CoverTab[26890]"
	} else {
//line /usr/local/go/src/encoding/json/decode.go:260
		_go_fuzz_dep_.CoverTab[26892]++
//line /usr/local/go/src/encoding/json/decode.go:260
		// _ = "end of CoverTab[26892]"
//line /usr/local/go/src/encoding/json/decode.go:260
	}
//line /usr/local/go/src/encoding/json/decode.go:260
	// _ = "end of CoverTab[26886]"
//line /usr/local/go/src/encoding/json/decode.go:260
	_go_fuzz_dep_.CoverTab[26887]++
							return err
//line /usr/local/go/src/encoding/json/decode.go:261
	// _ = "end of CoverTab[26887]"
}

// skip scans to the end of what was started.
func (d *decodeState) skip() {
//line /usr/local/go/src/encoding/json/decode.go:265
	_go_fuzz_dep_.CoverTab[26893]++
							s, data, i := &d.scan, d.data, d.off
							depth := len(s.parseState)
							for {
//line /usr/local/go/src/encoding/json/decode.go:268
		_go_fuzz_dep_.CoverTab[26894]++
								op := s.step(s, data[i])
								i++
								if len(s.parseState) < depth {
//line /usr/local/go/src/encoding/json/decode.go:271
			_go_fuzz_dep_.CoverTab[26895]++
									d.off = i
									d.opcode = op
									return
//line /usr/local/go/src/encoding/json/decode.go:274
			// _ = "end of CoverTab[26895]"
		} else {
//line /usr/local/go/src/encoding/json/decode.go:275
			_go_fuzz_dep_.CoverTab[26896]++
//line /usr/local/go/src/encoding/json/decode.go:275
			// _ = "end of CoverTab[26896]"
//line /usr/local/go/src/encoding/json/decode.go:275
		}
//line /usr/local/go/src/encoding/json/decode.go:275
		// _ = "end of CoverTab[26894]"
	}
//line /usr/local/go/src/encoding/json/decode.go:276
	// _ = "end of CoverTab[26893]"
}

// scanNext processes the byte at d.data[d.off].
func (d *decodeState) scanNext() {
//line /usr/local/go/src/encoding/json/decode.go:280
	_go_fuzz_dep_.CoverTab[26897]++
							if d.off < len(d.data) {
//line /usr/local/go/src/encoding/json/decode.go:281
		_go_fuzz_dep_.CoverTab[26898]++
								d.opcode = d.scan.step(&d.scan, d.data[d.off])
								d.off++
//line /usr/local/go/src/encoding/json/decode.go:283
		// _ = "end of CoverTab[26898]"
	} else {
//line /usr/local/go/src/encoding/json/decode.go:284
		_go_fuzz_dep_.CoverTab[26899]++
								d.opcode = d.scan.eof()
								d.off = len(d.data) + 1
//line /usr/local/go/src/encoding/json/decode.go:286
		// _ = "end of CoverTab[26899]"
	}
//line /usr/local/go/src/encoding/json/decode.go:287
	// _ = "end of CoverTab[26897]"
}

// scanWhile processes bytes in d.data[d.off:] until it
//line /usr/local/go/src/encoding/json/decode.go:290
// receives a scan code not equal to op.
//line /usr/local/go/src/encoding/json/decode.go:292
func (d *decodeState) scanWhile(op int) {
//line /usr/local/go/src/encoding/json/decode.go:292
	_go_fuzz_dep_.CoverTab[26900]++
							s, data, i := &d.scan, d.data, d.off
							for i < len(data) {
//line /usr/local/go/src/encoding/json/decode.go:294
		_go_fuzz_dep_.CoverTab[26902]++
								newOp := s.step(s, data[i])
								i++
								if newOp != op {
//line /usr/local/go/src/encoding/json/decode.go:297
			_go_fuzz_dep_.CoverTab[26903]++
									d.opcode = newOp
									d.off = i
									return
//line /usr/local/go/src/encoding/json/decode.go:300
			// _ = "end of CoverTab[26903]"
		} else {
//line /usr/local/go/src/encoding/json/decode.go:301
			_go_fuzz_dep_.CoverTab[26904]++
//line /usr/local/go/src/encoding/json/decode.go:301
			// _ = "end of CoverTab[26904]"
//line /usr/local/go/src/encoding/json/decode.go:301
		}
//line /usr/local/go/src/encoding/json/decode.go:301
		// _ = "end of CoverTab[26902]"
	}
//line /usr/local/go/src/encoding/json/decode.go:302
	// _ = "end of CoverTab[26900]"
//line /usr/local/go/src/encoding/json/decode.go:302
	_go_fuzz_dep_.CoverTab[26901]++

							d.off = len(data) + 1
							d.opcode = d.scan.eof()
//line /usr/local/go/src/encoding/json/decode.go:305
	// _ = "end of CoverTab[26901]"
}

// rescanLiteral is similar to scanWhile(scanContinue), but it specialises the
//line /usr/local/go/src/encoding/json/decode.go:308
// common case where we're decoding a literal. The decoder scans the input
//line /usr/local/go/src/encoding/json/decode.go:308
// twice, once for syntax errors and to check the length of the value, and the
//line /usr/local/go/src/encoding/json/decode.go:308
// second to perform the decoding.
//line /usr/local/go/src/encoding/json/decode.go:308
//
//line /usr/local/go/src/encoding/json/decode.go:308
// Only in the second step do we use decodeState to tokenize literals, so we
//line /usr/local/go/src/encoding/json/decode.go:308
// know there aren't any syntax errors. We can take advantage of that knowledge,
//line /usr/local/go/src/encoding/json/decode.go:308
// and scan a literal's bytes much more quickly.
//line /usr/local/go/src/encoding/json/decode.go:316
func (d *decodeState) rescanLiteral() {
//line /usr/local/go/src/encoding/json/decode.go:316
	_go_fuzz_dep_.CoverTab[26905]++
							data, i := d.data, d.off
Switch:
	switch data[i-1] {
	case '"':
//line /usr/local/go/src/encoding/json/decode.go:320
		_go_fuzz_dep_.CoverTab[26908]++
								for ; i < len(data); i++ {
//line /usr/local/go/src/encoding/json/decode.go:321
			_go_fuzz_dep_.CoverTab[26914]++
									switch data[i] {
			case '\\':
//line /usr/local/go/src/encoding/json/decode.go:323
				_go_fuzz_dep_.CoverTab[26915]++
										i++
//line /usr/local/go/src/encoding/json/decode.go:324
				// _ = "end of CoverTab[26915]"
			case '"':
//line /usr/local/go/src/encoding/json/decode.go:325
				_go_fuzz_dep_.CoverTab[26916]++
										i++
										break Switch
//line /usr/local/go/src/encoding/json/decode.go:327
				// _ = "end of CoverTab[26916]"
//line /usr/local/go/src/encoding/json/decode.go:327
			default:
//line /usr/local/go/src/encoding/json/decode.go:327
				_go_fuzz_dep_.CoverTab[26917]++
//line /usr/local/go/src/encoding/json/decode.go:327
				// _ = "end of CoverTab[26917]"
			}
//line /usr/local/go/src/encoding/json/decode.go:328
			// _ = "end of CoverTab[26914]"
		}
//line /usr/local/go/src/encoding/json/decode.go:329
		// _ = "end of CoverTab[26908]"
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '-':
//line /usr/local/go/src/encoding/json/decode.go:330
		_go_fuzz_dep_.CoverTab[26909]++
								for ; i < len(data); i++ {
//line /usr/local/go/src/encoding/json/decode.go:331
			_go_fuzz_dep_.CoverTab[26918]++
									switch data[i] {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
				'.', 'e', 'E', '+', '-':
//line /usr/local/go/src/encoding/json/decode.go:334
				_go_fuzz_dep_.CoverTab[26919]++
//line /usr/local/go/src/encoding/json/decode.go:334
				// _ = "end of CoverTab[26919]"
			default:
//line /usr/local/go/src/encoding/json/decode.go:335
				_go_fuzz_dep_.CoverTab[26920]++
										break Switch
//line /usr/local/go/src/encoding/json/decode.go:336
				// _ = "end of CoverTab[26920]"
			}
//line /usr/local/go/src/encoding/json/decode.go:337
			// _ = "end of CoverTab[26918]"
		}
//line /usr/local/go/src/encoding/json/decode.go:338
		// _ = "end of CoverTab[26909]"
	case 't':
//line /usr/local/go/src/encoding/json/decode.go:339
		_go_fuzz_dep_.CoverTab[26910]++
								i += len("rue")
//line /usr/local/go/src/encoding/json/decode.go:340
		// _ = "end of CoverTab[26910]"
	case 'f':
//line /usr/local/go/src/encoding/json/decode.go:341
		_go_fuzz_dep_.CoverTab[26911]++
								i += len("alse")
//line /usr/local/go/src/encoding/json/decode.go:342
		// _ = "end of CoverTab[26911]"
	case 'n':
//line /usr/local/go/src/encoding/json/decode.go:343
		_go_fuzz_dep_.CoverTab[26912]++
								i += len("ull")
//line /usr/local/go/src/encoding/json/decode.go:344
		// _ = "end of CoverTab[26912]"
//line /usr/local/go/src/encoding/json/decode.go:344
	default:
//line /usr/local/go/src/encoding/json/decode.go:344
		_go_fuzz_dep_.CoverTab[26913]++
//line /usr/local/go/src/encoding/json/decode.go:344
		// _ = "end of CoverTab[26913]"
	}
//line /usr/local/go/src/encoding/json/decode.go:345
	// _ = "end of CoverTab[26905]"
//line /usr/local/go/src/encoding/json/decode.go:345
	_go_fuzz_dep_.CoverTab[26906]++
							if i < len(data) {
//line /usr/local/go/src/encoding/json/decode.go:346
		_go_fuzz_dep_.CoverTab[26921]++
								d.opcode = stateEndValue(&d.scan, data[i])
//line /usr/local/go/src/encoding/json/decode.go:347
		// _ = "end of CoverTab[26921]"
	} else {
//line /usr/local/go/src/encoding/json/decode.go:348
		_go_fuzz_dep_.CoverTab[26922]++
								d.opcode = scanEnd
//line /usr/local/go/src/encoding/json/decode.go:349
		// _ = "end of CoverTab[26922]"
	}
//line /usr/local/go/src/encoding/json/decode.go:350
	// _ = "end of CoverTab[26906]"
//line /usr/local/go/src/encoding/json/decode.go:350
	_go_fuzz_dep_.CoverTab[26907]++
							d.off = i + 1
//line /usr/local/go/src/encoding/json/decode.go:351
	// _ = "end of CoverTab[26907]"
}

// value consumes a JSON value from d.data[d.off-1:], decoding into v, and
//line /usr/local/go/src/encoding/json/decode.go:354
// reads the following byte ahead. If v is invalid, the value is discarded.
//line /usr/local/go/src/encoding/json/decode.go:354
// The first byte of the value has been read already.
//line /usr/local/go/src/encoding/json/decode.go:357
func (d *decodeState) value(v reflect.Value) error {
//line /usr/local/go/src/encoding/json/decode.go:357
	_go_fuzz_dep_.CoverTab[26923]++
							switch d.opcode {
	default:
//line /usr/local/go/src/encoding/json/decode.go:359
		_go_fuzz_dep_.CoverTab[26925]++
								panic(phasePanicMsg)
//line /usr/local/go/src/encoding/json/decode.go:360
		// _ = "end of CoverTab[26925]"

	case scanBeginArray:
//line /usr/local/go/src/encoding/json/decode.go:362
		_go_fuzz_dep_.CoverTab[26926]++
								if v.IsValid() {
//line /usr/local/go/src/encoding/json/decode.go:363
			_go_fuzz_dep_.CoverTab[26931]++
									if err := d.array(v); err != nil {
//line /usr/local/go/src/encoding/json/decode.go:364
				_go_fuzz_dep_.CoverTab[26932]++
										return err
//line /usr/local/go/src/encoding/json/decode.go:365
				// _ = "end of CoverTab[26932]"
			} else {
//line /usr/local/go/src/encoding/json/decode.go:366
				_go_fuzz_dep_.CoverTab[26933]++
//line /usr/local/go/src/encoding/json/decode.go:366
				// _ = "end of CoverTab[26933]"
//line /usr/local/go/src/encoding/json/decode.go:366
			}
//line /usr/local/go/src/encoding/json/decode.go:366
			// _ = "end of CoverTab[26931]"
		} else {
//line /usr/local/go/src/encoding/json/decode.go:367
			_go_fuzz_dep_.CoverTab[26934]++
									d.skip()
//line /usr/local/go/src/encoding/json/decode.go:368
			// _ = "end of CoverTab[26934]"
		}
//line /usr/local/go/src/encoding/json/decode.go:369
		// _ = "end of CoverTab[26926]"
//line /usr/local/go/src/encoding/json/decode.go:369
		_go_fuzz_dep_.CoverTab[26927]++
								d.scanNext()
//line /usr/local/go/src/encoding/json/decode.go:370
		// _ = "end of CoverTab[26927]"

	case scanBeginObject:
//line /usr/local/go/src/encoding/json/decode.go:372
		_go_fuzz_dep_.CoverTab[26928]++
								if v.IsValid() {
//line /usr/local/go/src/encoding/json/decode.go:373
			_go_fuzz_dep_.CoverTab[26935]++
									if err := d.object(v); err != nil {
//line /usr/local/go/src/encoding/json/decode.go:374
				_go_fuzz_dep_.CoverTab[26936]++
										return err
//line /usr/local/go/src/encoding/json/decode.go:375
				// _ = "end of CoverTab[26936]"
			} else {
//line /usr/local/go/src/encoding/json/decode.go:376
				_go_fuzz_dep_.CoverTab[26937]++
//line /usr/local/go/src/encoding/json/decode.go:376
				// _ = "end of CoverTab[26937]"
//line /usr/local/go/src/encoding/json/decode.go:376
			}
//line /usr/local/go/src/encoding/json/decode.go:376
			// _ = "end of CoverTab[26935]"
		} else {
//line /usr/local/go/src/encoding/json/decode.go:377
			_go_fuzz_dep_.CoverTab[26938]++
									d.skip()
//line /usr/local/go/src/encoding/json/decode.go:378
			// _ = "end of CoverTab[26938]"
		}
//line /usr/local/go/src/encoding/json/decode.go:379
		// _ = "end of CoverTab[26928]"
//line /usr/local/go/src/encoding/json/decode.go:379
		_go_fuzz_dep_.CoverTab[26929]++
								d.scanNext()
//line /usr/local/go/src/encoding/json/decode.go:380
		// _ = "end of CoverTab[26929]"

	case scanBeginLiteral:
//line /usr/local/go/src/encoding/json/decode.go:382
		_go_fuzz_dep_.CoverTab[26930]++

								start := d.readIndex()
								d.rescanLiteral()

								if v.IsValid() {
//line /usr/local/go/src/encoding/json/decode.go:387
			_go_fuzz_dep_.CoverTab[26939]++
									if err := d.literalStore(d.data[start:d.readIndex()], v, false); err != nil {
//line /usr/local/go/src/encoding/json/decode.go:388
				_go_fuzz_dep_.CoverTab[26940]++
										return err
//line /usr/local/go/src/encoding/json/decode.go:389
				// _ = "end of CoverTab[26940]"
			} else {
//line /usr/local/go/src/encoding/json/decode.go:390
				_go_fuzz_dep_.CoverTab[26941]++
//line /usr/local/go/src/encoding/json/decode.go:390
				// _ = "end of CoverTab[26941]"
//line /usr/local/go/src/encoding/json/decode.go:390
			}
//line /usr/local/go/src/encoding/json/decode.go:390
			// _ = "end of CoverTab[26939]"
		} else {
//line /usr/local/go/src/encoding/json/decode.go:391
			_go_fuzz_dep_.CoverTab[26942]++
//line /usr/local/go/src/encoding/json/decode.go:391
			// _ = "end of CoverTab[26942]"
//line /usr/local/go/src/encoding/json/decode.go:391
		}
//line /usr/local/go/src/encoding/json/decode.go:391
		// _ = "end of CoverTab[26930]"
	}
//line /usr/local/go/src/encoding/json/decode.go:392
	// _ = "end of CoverTab[26923]"
//line /usr/local/go/src/encoding/json/decode.go:392
	_go_fuzz_dep_.CoverTab[26924]++
							return nil
//line /usr/local/go/src/encoding/json/decode.go:393
	// _ = "end of CoverTab[26924]"
}

type unquotedValue struct{}

// valueQuoted is like value but decodes a
//line /usr/local/go/src/encoding/json/decode.go:398
// quoted string literal or literal null into an interface value.
//line /usr/local/go/src/encoding/json/decode.go:398
// If it finds anything other than a quoted string literal or null,
//line /usr/local/go/src/encoding/json/decode.go:398
// valueQuoted returns unquotedValue{}.
//line /usr/local/go/src/encoding/json/decode.go:402
func (d *decodeState) valueQuoted() any {
//line /usr/local/go/src/encoding/json/decode.go:402
	_go_fuzz_dep_.CoverTab[26943]++
							switch d.opcode {
	default:
//line /usr/local/go/src/encoding/json/decode.go:404
		_go_fuzz_dep_.CoverTab[26945]++
								panic(phasePanicMsg)
//line /usr/local/go/src/encoding/json/decode.go:405
		// _ = "end of CoverTab[26945]"

	case scanBeginArray, scanBeginObject:
//line /usr/local/go/src/encoding/json/decode.go:407
		_go_fuzz_dep_.CoverTab[26946]++
								d.skip()
								d.scanNext()
//line /usr/local/go/src/encoding/json/decode.go:409
		// _ = "end of CoverTab[26946]"

	case scanBeginLiteral:
//line /usr/local/go/src/encoding/json/decode.go:411
		_go_fuzz_dep_.CoverTab[26947]++
								v := d.literalInterface()
								switch v.(type) {
		case nil, string:
//line /usr/local/go/src/encoding/json/decode.go:414
			_go_fuzz_dep_.CoverTab[26948]++
									return v
//line /usr/local/go/src/encoding/json/decode.go:415
			// _ = "end of CoverTab[26948]"
		}
//line /usr/local/go/src/encoding/json/decode.go:416
		// _ = "end of CoverTab[26947]"
	}
//line /usr/local/go/src/encoding/json/decode.go:417
	// _ = "end of CoverTab[26943]"
//line /usr/local/go/src/encoding/json/decode.go:417
	_go_fuzz_dep_.CoverTab[26944]++
							return unquotedValue{}
//line /usr/local/go/src/encoding/json/decode.go:418
	// _ = "end of CoverTab[26944]"
}

// indirect walks down v allocating pointers as needed,
//line /usr/local/go/src/encoding/json/decode.go:421
// until it gets to a non-pointer.
//line /usr/local/go/src/encoding/json/decode.go:421
// If it encounters an Unmarshaler, indirect stops and returns that.
//line /usr/local/go/src/encoding/json/decode.go:421
// If decodingNull is true, indirect stops at the first settable pointer so it
//line /usr/local/go/src/encoding/json/decode.go:421
// can be set to nil.
//line /usr/local/go/src/encoding/json/decode.go:426
func indirect(v reflect.Value, decodingNull bool) (Unmarshaler, encoding.TextUnmarshaler, reflect.Value) {
//line /usr/local/go/src/encoding/json/decode.go:426
	_go_fuzz_dep_.CoverTab[26949]++

//line /usr/local/go/src/encoding/json/decode.go:438
	v0 := v
							haveAddr := false

//line /usr/local/go/src/encoding/json/decode.go:444
	if v.Kind() != reflect.Pointer && func() bool {
//line /usr/local/go/src/encoding/json/decode.go:444
		_go_fuzz_dep_.CoverTab[26952]++
//line /usr/local/go/src/encoding/json/decode.go:444
		return v.Type().Name() != ""
//line /usr/local/go/src/encoding/json/decode.go:444
		// _ = "end of CoverTab[26952]"
//line /usr/local/go/src/encoding/json/decode.go:444
	}() && func() bool {
//line /usr/local/go/src/encoding/json/decode.go:444
		_go_fuzz_dep_.CoverTab[26953]++
//line /usr/local/go/src/encoding/json/decode.go:444
		return v.CanAddr()
//line /usr/local/go/src/encoding/json/decode.go:444
		// _ = "end of CoverTab[26953]"
//line /usr/local/go/src/encoding/json/decode.go:444
	}() {
//line /usr/local/go/src/encoding/json/decode.go:444
		_go_fuzz_dep_.CoverTab[26954]++
								haveAddr = true
								v = v.Addr()
//line /usr/local/go/src/encoding/json/decode.go:446
		// _ = "end of CoverTab[26954]"
	} else {
//line /usr/local/go/src/encoding/json/decode.go:447
		_go_fuzz_dep_.CoverTab[26955]++
//line /usr/local/go/src/encoding/json/decode.go:447
		// _ = "end of CoverTab[26955]"
//line /usr/local/go/src/encoding/json/decode.go:447
	}
//line /usr/local/go/src/encoding/json/decode.go:447
	// _ = "end of CoverTab[26949]"
//line /usr/local/go/src/encoding/json/decode.go:447
	_go_fuzz_dep_.CoverTab[26950]++
							for {
//line /usr/local/go/src/encoding/json/decode.go:448
		_go_fuzz_dep_.CoverTab[26956]++

//line /usr/local/go/src/encoding/json/decode.go:451
		if v.Kind() == reflect.Interface && func() bool {
//line /usr/local/go/src/encoding/json/decode.go:451
			_go_fuzz_dep_.CoverTab[26963]++
//line /usr/local/go/src/encoding/json/decode.go:451
			return !v.IsNil()
//line /usr/local/go/src/encoding/json/decode.go:451
			// _ = "end of CoverTab[26963]"
//line /usr/local/go/src/encoding/json/decode.go:451
		}() {
//line /usr/local/go/src/encoding/json/decode.go:451
			_go_fuzz_dep_.CoverTab[26964]++
									e := v.Elem()
									if e.Kind() == reflect.Pointer && func() bool {
//line /usr/local/go/src/encoding/json/decode.go:453
				_go_fuzz_dep_.CoverTab[26965]++
//line /usr/local/go/src/encoding/json/decode.go:453
				return !e.IsNil()
//line /usr/local/go/src/encoding/json/decode.go:453
				// _ = "end of CoverTab[26965]"
//line /usr/local/go/src/encoding/json/decode.go:453
			}() && func() bool {
//line /usr/local/go/src/encoding/json/decode.go:453
				_go_fuzz_dep_.CoverTab[26966]++
//line /usr/local/go/src/encoding/json/decode.go:453
				return (!decodingNull || func() bool {
//line /usr/local/go/src/encoding/json/decode.go:453
					_go_fuzz_dep_.CoverTab[26967]++
//line /usr/local/go/src/encoding/json/decode.go:453
					return e.Elem().Kind() == reflect.Pointer
//line /usr/local/go/src/encoding/json/decode.go:453
					// _ = "end of CoverTab[26967]"
//line /usr/local/go/src/encoding/json/decode.go:453
				}())
//line /usr/local/go/src/encoding/json/decode.go:453
				// _ = "end of CoverTab[26966]"
//line /usr/local/go/src/encoding/json/decode.go:453
			}() {
//line /usr/local/go/src/encoding/json/decode.go:453
				_go_fuzz_dep_.CoverTab[26968]++
										haveAddr = false
										v = e
										continue
//line /usr/local/go/src/encoding/json/decode.go:456
				// _ = "end of CoverTab[26968]"
			} else {
//line /usr/local/go/src/encoding/json/decode.go:457
				_go_fuzz_dep_.CoverTab[26969]++
//line /usr/local/go/src/encoding/json/decode.go:457
				// _ = "end of CoverTab[26969]"
//line /usr/local/go/src/encoding/json/decode.go:457
			}
//line /usr/local/go/src/encoding/json/decode.go:457
			// _ = "end of CoverTab[26964]"
		} else {
//line /usr/local/go/src/encoding/json/decode.go:458
			_go_fuzz_dep_.CoverTab[26970]++
//line /usr/local/go/src/encoding/json/decode.go:458
			// _ = "end of CoverTab[26970]"
//line /usr/local/go/src/encoding/json/decode.go:458
		}
//line /usr/local/go/src/encoding/json/decode.go:458
		// _ = "end of CoverTab[26956]"
//line /usr/local/go/src/encoding/json/decode.go:458
		_go_fuzz_dep_.CoverTab[26957]++

								if v.Kind() != reflect.Pointer {
//line /usr/local/go/src/encoding/json/decode.go:460
			_go_fuzz_dep_.CoverTab[26971]++
									break
//line /usr/local/go/src/encoding/json/decode.go:461
			// _ = "end of CoverTab[26971]"
		} else {
//line /usr/local/go/src/encoding/json/decode.go:462
			_go_fuzz_dep_.CoverTab[26972]++
//line /usr/local/go/src/encoding/json/decode.go:462
			// _ = "end of CoverTab[26972]"
//line /usr/local/go/src/encoding/json/decode.go:462
		}
//line /usr/local/go/src/encoding/json/decode.go:462
		// _ = "end of CoverTab[26957]"
//line /usr/local/go/src/encoding/json/decode.go:462
		_go_fuzz_dep_.CoverTab[26958]++

								if decodingNull && func() bool {
//line /usr/local/go/src/encoding/json/decode.go:464
			_go_fuzz_dep_.CoverTab[26973]++
//line /usr/local/go/src/encoding/json/decode.go:464
			return v.CanSet()
//line /usr/local/go/src/encoding/json/decode.go:464
			// _ = "end of CoverTab[26973]"
//line /usr/local/go/src/encoding/json/decode.go:464
		}() {
//line /usr/local/go/src/encoding/json/decode.go:464
			_go_fuzz_dep_.CoverTab[26974]++
									break
//line /usr/local/go/src/encoding/json/decode.go:465
			// _ = "end of CoverTab[26974]"
		} else {
//line /usr/local/go/src/encoding/json/decode.go:466
			_go_fuzz_dep_.CoverTab[26975]++
//line /usr/local/go/src/encoding/json/decode.go:466
			// _ = "end of CoverTab[26975]"
//line /usr/local/go/src/encoding/json/decode.go:466
		}
//line /usr/local/go/src/encoding/json/decode.go:466
		// _ = "end of CoverTab[26958]"
//line /usr/local/go/src/encoding/json/decode.go:466
		_go_fuzz_dep_.CoverTab[26959]++

//line /usr/local/go/src/encoding/json/decode.go:471
		if v.Elem().Kind() == reflect.Interface && func() bool {
//line /usr/local/go/src/encoding/json/decode.go:471
			_go_fuzz_dep_.CoverTab[26976]++
//line /usr/local/go/src/encoding/json/decode.go:471
			return v.Elem().Elem() == v
//line /usr/local/go/src/encoding/json/decode.go:471
			// _ = "end of CoverTab[26976]"
//line /usr/local/go/src/encoding/json/decode.go:471
		}() {
//line /usr/local/go/src/encoding/json/decode.go:471
			_go_fuzz_dep_.CoverTab[26977]++
									v = v.Elem()
									break
//line /usr/local/go/src/encoding/json/decode.go:473
			// _ = "end of CoverTab[26977]"
		} else {
//line /usr/local/go/src/encoding/json/decode.go:474
			_go_fuzz_dep_.CoverTab[26978]++
//line /usr/local/go/src/encoding/json/decode.go:474
			// _ = "end of CoverTab[26978]"
//line /usr/local/go/src/encoding/json/decode.go:474
		}
//line /usr/local/go/src/encoding/json/decode.go:474
		// _ = "end of CoverTab[26959]"
//line /usr/local/go/src/encoding/json/decode.go:474
		_go_fuzz_dep_.CoverTab[26960]++
								if v.IsNil() {
//line /usr/local/go/src/encoding/json/decode.go:475
			_go_fuzz_dep_.CoverTab[26979]++
									v.Set(reflect.New(v.Type().Elem()))
//line /usr/local/go/src/encoding/json/decode.go:476
			// _ = "end of CoverTab[26979]"
		} else {
//line /usr/local/go/src/encoding/json/decode.go:477
			_go_fuzz_dep_.CoverTab[26980]++
//line /usr/local/go/src/encoding/json/decode.go:477
			// _ = "end of CoverTab[26980]"
//line /usr/local/go/src/encoding/json/decode.go:477
		}
//line /usr/local/go/src/encoding/json/decode.go:477
		// _ = "end of CoverTab[26960]"
//line /usr/local/go/src/encoding/json/decode.go:477
		_go_fuzz_dep_.CoverTab[26961]++
								if v.Type().NumMethod() > 0 && func() bool {
//line /usr/local/go/src/encoding/json/decode.go:478
			_go_fuzz_dep_.CoverTab[26981]++
//line /usr/local/go/src/encoding/json/decode.go:478
			return v.CanInterface()
//line /usr/local/go/src/encoding/json/decode.go:478
			// _ = "end of CoverTab[26981]"
//line /usr/local/go/src/encoding/json/decode.go:478
		}() {
//line /usr/local/go/src/encoding/json/decode.go:478
			_go_fuzz_dep_.CoverTab[26982]++
									if u, ok := v.Interface().(Unmarshaler); ok {
//line /usr/local/go/src/encoding/json/decode.go:479
				_go_fuzz_dep_.CoverTab[26984]++
										return u, nil, reflect.Value{}
//line /usr/local/go/src/encoding/json/decode.go:480
				// _ = "end of CoverTab[26984]"
			} else {
//line /usr/local/go/src/encoding/json/decode.go:481
				_go_fuzz_dep_.CoverTab[26985]++
//line /usr/local/go/src/encoding/json/decode.go:481
				// _ = "end of CoverTab[26985]"
//line /usr/local/go/src/encoding/json/decode.go:481
			}
//line /usr/local/go/src/encoding/json/decode.go:481
			// _ = "end of CoverTab[26982]"
//line /usr/local/go/src/encoding/json/decode.go:481
			_go_fuzz_dep_.CoverTab[26983]++
									if !decodingNull {
//line /usr/local/go/src/encoding/json/decode.go:482
				_go_fuzz_dep_.CoverTab[26986]++
										if u, ok := v.Interface().(encoding.TextUnmarshaler); ok {
//line /usr/local/go/src/encoding/json/decode.go:483
					_go_fuzz_dep_.CoverTab[26987]++
											return nil, u, reflect.Value{}
//line /usr/local/go/src/encoding/json/decode.go:484
					// _ = "end of CoverTab[26987]"
				} else {
//line /usr/local/go/src/encoding/json/decode.go:485
					_go_fuzz_dep_.CoverTab[26988]++
//line /usr/local/go/src/encoding/json/decode.go:485
					// _ = "end of CoverTab[26988]"
//line /usr/local/go/src/encoding/json/decode.go:485
				}
//line /usr/local/go/src/encoding/json/decode.go:485
				// _ = "end of CoverTab[26986]"
			} else {
//line /usr/local/go/src/encoding/json/decode.go:486
				_go_fuzz_dep_.CoverTab[26989]++
//line /usr/local/go/src/encoding/json/decode.go:486
				// _ = "end of CoverTab[26989]"
//line /usr/local/go/src/encoding/json/decode.go:486
			}
//line /usr/local/go/src/encoding/json/decode.go:486
			// _ = "end of CoverTab[26983]"
		} else {
//line /usr/local/go/src/encoding/json/decode.go:487
			_go_fuzz_dep_.CoverTab[26990]++
//line /usr/local/go/src/encoding/json/decode.go:487
			// _ = "end of CoverTab[26990]"
//line /usr/local/go/src/encoding/json/decode.go:487
		}
//line /usr/local/go/src/encoding/json/decode.go:487
		// _ = "end of CoverTab[26961]"
//line /usr/local/go/src/encoding/json/decode.go:487
		_go_fuzz_dep_.CoverTab[26962]++

								if haveAddr {
//line /usr/local/go/src/encoding/json/decode.go:489
			_go_fuzz_dep_.CoverTab[26991]++
									v = v0
									haveAddr = false
//line /usr/local/go/src/encoding/json/decode.go:491
			// _ = "end of CoverTab[26991]"
		} else {
//line /usr/local/go/src/encoding/json/decode.go:492
			_go_fuzz_dep_.CoverTab[26992]++
									v = v.Elem()
//line /usr/local/go/src/encoding/json/decode.go:493
			// _ = "end of CoverTab[26992]"
		}
//line /usr/local/go/src/encoding/json/decode.go:494
		// _ = "end of CoverTab[26962]"
	}
//line /usr/local/go/src/encoding/json/decode.go:495
	// _ = "end of CoverTab[26950]"
//line /usr/local/go/src/encoding/json/decode.go:495
	_go_fuzz_dep_.CoverTab[26951]++
							return nil, nil, v
//line /usr/local/go/src/encoding/json/decode.go:496
	// _ = "end of CoverTab[26951]"
}

// array consumes an array from d.data[d.off-1:], decoding into v.
//line /usr/local/go/src/encoding/json/decode.go:499
// The first byte of the array ('[') has been read already.
//line /usr/local/go/src/encoding/json/decode.go:501
func (d *decodeState) array(v reflect.Value) error {
//line /usr/local/go/src/encoding/json/decode.go:501
	_go_fuzz_dep_.CoverTab[26993]++

							u, ut, pv := indirect(v, false)
							if u != nil {
//line /usr/local/go/src/encoding/json/decode.go:504
		_go_fuzz_dep_.CoverTab[27000]++
								start := d.readIndex()
								d.skip()
								return u.UnmarshalJSON(d.data[start:d.off])
//line /usr/local/go/src/encoding/json/decode.go:507
		// _ = "end of CoverTab[27000]"
	} else {
//line /usr/local/go/src/encoding/json/decode.go:508
		_go_fuzz_dep_.CoverTab[27001]++
//line /usr/local/go/src/encoding/json/decode.go:508
		// _ = "end of CoverTab[27001]"
//line /usr/local/go/src/encoding/json/decode.go:508
	}
//line /usr/local/go/src/encoding/json/decode.go:508
	// _ = "end of CoverTab[26993]"
//line /usr/local/go/src/encoding/json/decode.go:508
	_go_fuzz_dep_.CoverTab[26994]++
							if ut != nil {
//line /usr/local/go/src/encoding/json/decode.go:509
		_go_fuzz_dep_.CoverTab[27002]++
								d.saveError(&UnmarshalTypeError{Value: "array", Type: v.Type(), Offset: int64(d.off)})
								d.skip()
								return nil
//line /usr/local/go/src/encoding/json/decode.go:512
		// _ = "end of CoverTab[27002]"
	} else {
//line /usr/local/go/src/encoding/json/decode.go:513
		_go_fuzz_dep_.CoverTab[27003]++
//line /usr/local/go/src/encoding/json/decode.go:513
		// _ = "end of CoverTab[27003]"
//line /usr/local/go/src/encoding/json/decode.go:513
	}
//line /usr/local/go/src/encoding/json/decode.go:513
	// _ = "end of CoverTab[26994]"
//line /usr/local/go/src/encoding/json/decode.go:513
	_go_fuzz_dep_.CoverTab[26995]++
							v = pv

//line /usr/local/go/src/encoding/json/decode.go:517
	switch v.Kind() {
	case reflect.Interface:
//line /usr/local/go/src/encoding/json/decode.go:518
		_go_fuzz_dep_.CoverTab[27004]++
								if v.NumMethod() == 0 {
//line /usr/local/go/src/encoding/json/decode.go:519
			_go_fuzz_dep_.CoverTab[27008]++

									ai := d.arrayInterface()
									v.Set(reflect.ValueOf(ai))
									return nil
//line /usr/local/go/src/encoding/json/decode.go:523
			// _ = "end of CoverTab[27008]"
		} else {
//line /usr/local/go/src/encoding/json/decode.go:524
			_go_fuzz_dep_.CoverTab[27009]++
//line /usr/local/go/src/encoding/json/decode.go:524
			// _ = "end of CoverTab[27009]"
//line /usr/local/go/src/encoding/json/decode.go:524
		}
//line /usr/local/go/src/encoding/json/decode.go:524
		// _ = "end of CoverTab[27004]"
//line /usr/local/go/src/encoding/json/decode.go:524
		_go_fuzz_dep_.CoverTab[27005]++

								fallthrough
//line /usr/local/go/src/encoding/json/decode.go:526
		// _ = "end of CoverTab[27005]"
	default:
//line /usr/local/go/src/encoding/json/decode.go:527
		_go_fuzz_dep_.CoverTab[27006]++
								d.saveError(&UnmarshalTypeError{Value: "array", Type: v.Type(), Offset: int64(d.off)})
								d.skip()
								return nil
//line /usr/local/go/src/encoding/json/decode.go:530
		// _ = "end of CoverTab[27006]"
	case reflect.Array, reflect.Slice:
//line /usr/local/go/src/encoding/json/decode.go:531
		_go_fuzz_dep_.CoverTab[27007]++
								break
//line /usr/local/go/src/encoding/json/decode.go:532
		// _ = "end of CoverTab[27007]"
	}
//line /usr/local/go/src/encoding/json/decode.go:533
	// _ = "end of CoverTab[26995]"
//line /usr/local/go/src/encoding/json/decode.go:533
	_go_fuzz_dep_.CoverTab[26996]++

							i := 0
							for {
//line /usr/local/go/src/encoding/json/decode.go:536
		_go_fuzz_dep_.CoverTab[27010]++

								d.scanWhile(scanSkipSpace)
								if d.opcode == scanEndArray {
//line /usr/local/go/src/encoding/json/decode.go:539
			_go_fuzz_dep_.CoverTab[27016]++
									break
//line /usr/local/go/src/encoding/json/decode.go:540
			// _ = "end of CoverTab[27016]"
		} else {
//line /usr/local/go/src/encoding/json/decode.go:541
			_go_fuzz_dep_.CoverTab[27017]++
//line /usr/local/go/src/encoding/json/decode.go:541
			// _ = "end of CoverTab[27017]"
//line /usr/local/go/src/encoding/json/decode.go:541
		}
//line /usr/local/go/src/encoding/json/decode.go:541
		// _ = "end of CoverTab[27010]"
//line /usr/local/go/src/encoding/json/decode.go:541
		_go_fuzz_dep_.CoverTab[27011]++

//line /usr/local/go/src/encoding/json/decode.go:544
		if v.Kind() == reflect.Slice {
//line /usr/local/go/src/encoding/json/decode.go:544
			_go_fuzz_dep_.CoverTab[27018]++

									if i >= v.Cap() {
//line /usr/local/go/src/encoding/json/decode.go:546
				_go_fuzz_dep_.CoverTab[27020]++
										newcap := v.Cap() + v.Cap()/2
										if newcap < 4 {
//line /usr/local/go/src/encoding/json/decode.go:548
					_go_fuzz_dep_.CoverTab[27022]++
											newcap = 4
//line /usr/local/go/src/encoding/json/decode.go:549
					// _ = "end of CoverTab[27022]"
				} else {
//line /usr/local/go/src/encoding/json/decode.go:550
					_go_fuzz_dep_.CoverTab[27023]++
//line /usr/local/go/src/encoding/json/decode.go:550
					// _ = "end of CoverTab[27023]"
//line /usr/local/go/src/encoding/json/decode.go:550
				}
//line /usr/local/go/src/encoding/json/decode.go:550
				// _ = "end of CoverTab[27020]"
//line /usr/local/go/src/encoding/json/decode.go:550
				_go_fuzz_dep_.CoverTab[27021]++
										newv := reflect.MakeSlice(v.Type(), v.Len(), newcap)
										reflect.Copy(newv, v)
										v.Set(newv)
//line /usr/local/go/src/encoding/json/decode.go:553
				// _ = "end of CoverTab[27021]"
			} else {
//line /usr/local/go/src/encoding/json/decode.go:554
				_go_fuzz_dep_.CoverTab[27024]++
//line /usr/local/go/src/encoding/json/decode.go:554
				// _ = "end of CoverTab[27024]"
//line /usr/local/go/src/encoding/json/decode.go:554
			}
//line /usr/local/go/src/encoding/json/decode.go:554
			// _ = "end of CoverTab[27018]"
//line /usr/local/go/src/encoding/json/decode.go:554
			_go_fuzz_dep_.CoverTab[27019]++
									if i >= v.Len() {
//line /usr/local/go/src/encoding/json/decode.go:555
				_go_fuzz_dep_.CoverTab[27025]++
										v.SetLen(i + 1)
//line /usr/local/go/src/encoding/json/decode.go:556
				// _ = "end of CoverTab[27025]"
			} else {
//line /usr/local/go/src/encoding/json/decode.go:557
				_go_fuzz_dep_.CoverTab[27026]++
//line /usr/local/go/src/encoding/json/decode.go:557
				// _ = "end of CoverTab[27026]"
//line /usr/local/go/src/encoding/json/decode.go:557
			}
//line /usr/local/go/src/encoding/json/decode.go:557
			// _ = "end of CoverTab[27019]"
		} else {
//line /usr/local/go/src/encoding/json/decode.go:558
			_go_fuzz_dep_.CoverTab[27027]++
//line /usr/local/go/src/encoding/json/decode.go:558
			// _ = "end of CoverTab[27027]"
//line /usr/local/go/src/encoding/json/decode.go:558
		}
//line /usr/local/go/src/encoding/json/decode.go:558
		// _ = "end of CoverTab[27011]"
//line /usr/local/go/src/encoding/json/decode.go:558
		_go_fuzz_dep_.CoverTab[27012]++

								if i < v.Len() {
//line /usr/local/go/src/encoding/json/decode.go:560
			_go_fuzz_dep_.CoverTab[27028]++

									if err := d.value(v.Index(i)); err != nil {
//line /usr/local/go/src/encoding/json/decode.go:562
				_go_fuzz_dep_.CoverTab[27029]++
										return err
//line /usr/local/go/src/encoding/json/decode.go:563
				// _ = "end of CoverTab[27029]"
			} else {
//line /usr/local/go/src/encoding/json/decode.go:564
				_go_fuzz_dep_.CoverTab[27030]++
//line /usr/local/go/src/encoding/json/decode.go:564
				// _ = "end of CoverTab[27030]"
//line /usr/local/go/src/encoding/json/decode.go:564
			}
//line /usr/local/go/src/encoding/json/decode.go:564
			// _ = "end of CoverTab[27028]"
		} else {
//line /usr/local/go/src/encoding/json/decode.go:565
			_go_fuzz_dep_.CoverTab[27031]++

									if err := d.value(reflect.Value{}); err != nil {
//line /usr/local/go/src/encoding/json/decode.go:567
				_go_fuzz_dep_.CoverTab[27032]++
										return err
//line /usr/local/go/src/encoding/json/decode.go:568
				// _ = "end of CoverTab[27032]"
			} else {
//line /usr/local/go/src/encoding/json/decode.go:569
				_go_fuzz_dep_.CoverTab[27033]++
//line /usr/local/go/src/encoding/json/decode.go:569
				// _ = "end of CoverTab[27033]"
//line /usr/local/go/src/encoding/json/decode.go:569
			}
//line /usr/local/go/src/encoding/json/decode.go:569
			// _ = "end of CoverTab[27031]"
		}
//line /usr/local/go/src/encoding/json/decode.go:570
		// _ = "end of CoverTab[27012]"
//line /usr/local/go/src/encoding/json/decode.go:570
		_go_fuzz_dep_.CoverTab[27013]++
								i++

//line /usr/local/go/src/encoding/json/decode.go:574
		if d.opcode == scanSkipSpace {
//line /usr/local/go/src/encoding/json/decode.go:574
			_go_fuzz_dep_.CoverTab[27034]++
									d.scanWhile(scanSkipSpace)
//line /usr/local/go/src/encoding/json/decode.go:575
			// _ = "end of CoverTab[27034]"
		} else {
//line /usr/local/go/src/encoding/json/decode.go:576
			_go_fuzz_dep_.CoverTab[27035]++
//line /usr/local/go/src/encoding/json/decode.go:576
			// _ = "end of CoverTab[27035]"
//line /usr/local/go/src/encoding/json/decode.go:576
		}
//line /usr/local/go/src/encoding/json/decode.go:576
		// _ = "end of CoverTab[27013]"
//line /usr/local/go/src/encoding/json/decode.go:576
		_go_fuzz_dep_.CoverTab[27014]++
								if d.opcode == scanEndArray {
//line /usr/local/go/src/encoding/json/decode.go:577
			_go_fuzz_dep_.CoverTab[27036]++
									break
//line /usr/local/go/src/encoding/json/decode.go:578
			// _ = "end of CoverTab[27036]"
		} else {
//line /usr/local/go/src/encoding/json/decode.go:579
			_go_fuzz_dep_.CoverTab[27037]++
//line /usr/local/go/src/encoding/json/decode.go:579
			// _ = "end of CoverTab[27037]"
//line /usr/local/go/src/encoding/json/decode.go:579
		}
//line /usr/local/go/src/encoding/json/decode.go:579
		// _ = "end of CoverTab[27014]"
//line /usr/local/go/src/encoding/json/decode.go:579
		_go_fuzz_dep_.CoverTab[27015]++
								if d.opcode != scanArrayValue {
//line /usr/local/go/src/encoding/json/decode.go:580
			_go_fuzz_dep_.CoverTab[27038]++
									panic(phasePanicMsg)
//line /usr/local/go/src/encoding/json/decode.go:581
			// _ = "end of CoverTab[27038]"
		} else {
//line /usr/local/go/src/encoding/json/decode.go:582
			_go_fuzz_dep_.CoverTab[27039]++
//line /usr/local/go/src/encoding/json/decode.go:582
			// _ = "end of CoverTab[27039]"
//line /usr/local/go/src/encoding/json/decode.go:582
		}
//line /usr/local/go/src/encoding/json/decode.go:582
		// _ = "end of CoverTab[27015]"
	}
//line /usr/local/go/src/encoding/json/decode.go:583
	// _ = "end of CoverTab[26996]"
//line /usr/local/go/src/encoding/json/decode.go:583
	_go_fuzz_dep_.CoverTab[26997]++

							if i < v.Len() {
//line /usr/local/go/src/encoding/json/decode.go:585
		_go_fuzz_dep_.CoverTab[27040]++
								if v.Kind() == reflect.Array {
//line /usr/local/go/src/encoding/json/decode.go:586
			_go_fuzz_dep_.CoverTab[27041]++

									z := reflect.Zero(v.Type().Elem())
									for ; i < v.Len(); i++ {
//line /usr/local/go/src/encoding/json/decode.go:589
				_go_fuzz_dep_.CoverTab[27042]++
										v.Index(i).Set(z)
//line /usr/local/go/src/encoding/json/decode.go:590
				// _ = "end of CoverTab[27042]"
			}
//line /usr/local/go/src/encoding/json/decode.go:591
			// _ = "end of CoverTab[27041]"
		} else {
//line /usr/local/go/src/encoding/json/decode.go:592
			_go_fuzz_dep_.CoverTab[27043]++
									v.SetLen(i)
//line /usr/local/go/src/encoding/json/decode.go:593
			// _ = "end of CoverTab[27043]"
		}
//line /usr/local/go/src/encoding/json/decode.go:594
		// _ = "end of CoverTab[27040]"
	} else {
//line /usr/local/go/src/encoding/json/decode.go:595
		_go_fuzz_dep_.CoverTab[27044]++
//line /usr/local/go/src/encoding/json/decode.go:595
		// _ = "end of CoverTab[27044]"
//line /usr/local/go/src/encoding/json/decode.go:595
	}
//line /usr/local/go/src/encoding/json/decode.go:595
	// _ = "end of CoverTab[26997]"
//line /usr/local/go/src/encoding/json/decode.go:595
	_go_fuzz_dep_.CoverTab[26998]++
							if i == 0 && func() bool {
//line /usr/local/go/src/encoding/json/decode.go:596
		_go_fuzz_dep_.CoverTab[27045]++
//line /usr/local/go/src/encoding/json/decode.go:596
		return v.Kind() == reflect.Slice
//line /usr/local/go/src/encoding/json/decode.go:596
		// _ = "end of CoverTab[27045]"
//line /usr/local/go/src/encoding/json/decode.go:596
	}() {
//line /usr/local/go/src/encoding/json/decode.go:596
		_go_fuzz_dep_.CoverTab[27046]++
								v.Set(reflect.MakeSlice(v.Type(), 0, 0))
//line /usr/local/go/src/encoding/json/decode.go:597
		// _ = "end of CoverTab[27046]"
	} else {
//line /usr/local/go/src/encoding/json/decode.go:598
		_go_fuzz_dep_.CoverTab[27047]++
//line /usr/local/go/src/encoding/json/decode.go:598
		// _ = "end of CoverTab[27047]"
//line /usr/local/go/src/encoding/json/decode.go:598
	}
//line /usr/local/go/src/encoding/json/decode.go:598
	// _ = "end of CoverTab[26998]"
//line /usr/local/go/src/encoding/json/decode.go:598
	_go_fuzz_dep_.CoverTab[26999]++
							return nil
//line /usr/local/go/src/encoding/json/decode.go:599
	// _ = "end of CoverTab[26999]"
}

var nullLiteral = []byte("null")
var textUnmarshalerType = reflect.TypeOf((*encoding.TextUnmarshaler)(nil)).Elem()

// object consumes an object from d.data[d.off-1:], decoding into v.
//line /usr/local/go/src/encoding/json/decode.go:605
// The first byte ('{') of the object has been read already.
//line /usr/local/go/src/encoding/json/decode.go:607
func (d *decodeState) object(v reflect.Value) error {
//line /usr/local/go/src/encoding/json/decode.go:607
	_go_fuzz_dep_.CoverTab[27048]++

							u, ut, pv := indirect(v, false)
							if u != nil {
//line /usr/local/go/src/encoding/json/decode.go:610
		_go_fuzz_dep_.CoverTab[27055]++
								start := d.readIndex()
								d.skip()
								return u.UnmarshalJSON(d.data[start:d.off])
//line /usr/local/go/src/encoding/json/decode.go:613
		// _ = "end of CoverTab[27055]"
	} else {
//line /usr/local/go/src/encoding/json/decode.go:614
		_go_fuzz_dep_.CoverTab[27056]++
//line /usr/local/go/src/encoding/json/decode.go:614
		// _ = "end of CoverTab[27056]"
//line /usr/local/go/src/encoding/json/decode.go:614
	}
//line /usr/local/go/src/encoding/json/decode.go:614
	// _ = "end of CoverTab[27048]"
//line /usr/local/go/src/encoding/json/decode.go:614
	_go_fuzz_dep_.CoverTab[27049]++
							if ut != nil {
//line /usr/local/go/src/encoding/json/decode.go:615
		_go_fuzz_dep_.CoverTab[27057]++
								d.saveError(&UnmarshalTypeError{Value: "object", Type: v.Type(), Offset: int64(d.off)})
								d.skip()
								return nil
//line /usr/local/go/src/encoding/json/decode.go:618
		// _ = "end of CoverTab[27057]"
	} else {
//line /usr/local/go/src/encoding/json/decode.go:619
		_go_fuzz_dep_.CoverTab[27058]++
//line /usr/local/go/src/encoding/json/decode.go:619
		// _ = "end of CoverTab[27058]"
//line /usr/local/go/src/encoding/json/decode.go:619
	}
//line /usr/local/go/src/encoding/json/decode.go:619
	// _ = "end of CoverTab[27049]"
//line /usr/local/go/src/encoding/json/decode.go:619
	_go_fuzz_dep_.CoverTab[27050]++
							v = pv
							t := v.Type()

//line /usr/local/go/src/encoding/json/decode.go:624
	if v.Kind() == reflect.Interface && func() bool {
//line /usr/local/go/src/encoding/json/decode.go:624
		_go_fuzz_dep_.CoverTab[27059]++
//line /usr/local/go/src/encoding/json/decode.go:624
		return v.NumMethod() == 0
//line /usr/local/go/src/encoding/json/decode.go:624
		// _ = "end of CoverTab[27059]"
//line /usr/local/go/src/encoding/json/decode.go:624
	}() {
//line /usr/local/go/src/encoding/json/decode.go:624
		_go_fuzz_dep_.CoverTab[27060]++
								oi := d.objectInterface()
								v.Set(reflect.ValueOf(oi))
								return nil
//line /usr/local/go/src/encoding/json/decode.go:627
		// _ = "end of CoverTab[27060]"
	} else {
//line /usr/local/go/src/encoding/json/decode.go:628
		_go_fuzz_dep_.CoverTab[27061]++
//line /usr/local/go/src/encoding/json/decode.go:628
		// _ = "end of CoverTab[27061]"
//line /usr/local/go/src/encoding/json/decode.go:628
	}
//line /usr/local/go/src/encoding/json/decode.go:628
	// _ = "end of CoverTab[27050]"
//line /usr/local/go/src/encoding/json/decode.go:628
	_go_fuzz_dep_.CoverTab[27051]++

							var fields structFields

//line /usr/local/go/src/encoding/json/decode.go:636
	switch v.Kind() {
	case reflect.Map:
//line /usr/local/go/src/encoding/json/decode.go:637
		_go_fuzz_dep_.CoverTab[27062]++

//line /usr/local/go/src/encoding/json/decode.go:640
		switch t.Key().Kind() {
		case reflect.String,
			reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
//line /usr/local/go/src/encoding/json/decode.go:643
			_go_fuzz_dep_.CoverTab[27066]++
//line /usr/local/go/src/encoding/json/decode.go:643
			// _ = "end of CoverTab[27066]"
		default:
//line /usr/local/go/src/encoding/json/decode.go:644
			_go_fuzz_dep_.CoverTab[27067]++
									if !reflect.PointerTo(t.Key()).Implements(textUnmarshalerType) {
//line /usr/local/go/src/encoding/json/decode.go:645
				_go_fuzz_dep_.CoverTab[27068]++
										d.saveError(&UnmarshalTypeError{Value: "object", Type: t, Offset: int64(d.off)})
										d.skip()
										return nil
//line /usr/local/go/src/encoding/json/decode.go:648
				// _ = "end of CoverTab[27068]"
			} else {
//line /usr/local/go/src/encoding/json/decode.go:649
				_go_fuzz_dep_.CoverTab[27069]++
//line /usr/local/go/src/encoding/json/decode.go:649
				// _ = "end of CoverTab[27069]"
//line /usr/local/go/src/encoding/json/decode.go:649
			}
//line /usr/local/go/src/encoding/json/decode.go:649
			// _ = "end of CoverTab[27067]"
		}
//line /usr/local/go/src/encoding/json/decode.go:650
		// _ = "end of CoverTab[27062]"
//line /usr/local/go/src/encoding/json/decode.go:650
		_go_fuzz_dep_.CoverTab[27063]++
								if v.IsNil() {
//line /usr/local/go/src/encoding/json/decode.go:651
			_go_fuzz_dep_.CoverTab[27070]++
									v.Set(reflect.MakeMap(t))
//line /usr/local/go/src/encoding/json/decode.go:652
			// _ = "end of CoverTab[27070]"
		} else {
//line /usr/local/go/src/encoding/json/decode.go:653
			_go_fuzz_dep_.CoverTab[27071]++
//line /usr/local/go/src/encoding/json/decode.go:653
			// _ = "end of CoverTab[27071]"
//line /usr/local/go/src/encoding/json/decode.go:653
		}
//line /usr/local/go/src/encoding/json/decode.go:653
		// _ = "end of CoverTab[27063]"
	case reflect.Struct:
//line /usr/local/go/src/encoding/json/decode.go:654
		_go_fuzz_dep_.CoverTab[27064]++
								fields = cachedTypeFields(t)
//line /usr/local/go/src/encoding/json/decode.go:655
		// _ = "end of CoverTab[27064]"

	default:
//line /usr/local/go/src/encoding/json/decode.go:657
		_go_fuzz_dep_.CoverTab[27065]++
								d.saveError(&UnmarshalTypeError{Value: "object", Type: t, Offset: int64(d.off)})
								d.skip()
								return nil
//line /usr/local/go/src/encoding/json/decode.go:660
		// _ = "end of CoverTab[27065]"
	}
//line /usr/local/go/src/encoding/json/decode.go:661
	// _ = "end of CoverTab[27051]"
//line /usr/local/go/src/encoding/json/decode.go:661
	_go_fuzz_dep_.CoverTab[27052]++

							var mapElem reflect.Value
							var origErrorContext errorContext
							if d.errorContext != nil {
//line /usr/local/go/src/encoding/json/decode.go:665
		_go_fuzz_dep_.CoverTab[27072]++
								origErrorContext = *d.errorContext
//line /usr/local/go/src/encoding/json/decode.go:666
		// _ = "end of CoverTab[27072]"
	} else {
//line /usr/local/go/src/encoding/json/decode.go:667
		_go_fuzz_dep_.CoverTab[27073]++
//line /usr/local/go/src/encoding/json/decode.go:667
		// _ = "end of CoverTab[27073]"
//line /usr/local/go/src/encoding/json/decode.go:667
	}
//line /usr/local/go/src/encoding/json/decode.go:667
	// _ = "end of CoverTab[27052]"
//line /usr/local/go/src/encoding/json/decode.go:667
	_go_fuzz_dep_.CoverTab[27053]++

							for {
//line /usr/local/go/src/encoding/json/decode.go:669
		_go_fuzz_dep_.CoverTab[27074]++

								d.scanWhile(scanSkipSpace)
								if d.opcode == scanEndObject {
//line /usr/local/go/src/encoding/json/decode.go:672
			_go_fuzz_dep_.CoverTab[27086]++

									break
//line /usr/local/go/src/encoding/json/decode.go:674
			// _ = "end of CoverTab[27086]"
		} else {
//line /usr/local/go/src/encoding/json/decode.go:675
			_go_fuzz_dep_.CoverTab[27087]++
//line /usr/local/go/src/encoding/json/decode.go:675
			// _ = "end of CoverTab[27087]"
//line /usr/local/go/src/encoding/json/decode.go:675
		}
//line /usr/local/go/src/encoding/json/decode.go:675
		// _ = "end of CoverTab[27074]"
//line /usr/local/go/src/encoding/json/decode.go:675
		_go_fuzz_dep_.CoverTab[27075]++
								if d.opcode != scanBeginLiteral {
//line /usr/local/go/src/encoding/json/decode.go:676
			_go_fuzz_dep_.CoverTab[27088]++
									panic(phasePanicMsg)
//line /usr/local/go/src/encoding/json/decode.go:677
			// _ = "end of CoverTab[27088]"
		} else {
//line /usr/local/go/src/encoding/json/decode.go:678
			_go_fuzz_dep_.CoverTab[27089]++
//line /usr/local/go/src/encoding/json/decode.go:678
			// _ = "end of CoverTab[27089]"
//line /usr/local/go/src/encoding/json/decode.go:678
		}
//line /usr/local/go/src/encoding/json/decode.go:678
		// _ = "end of CoverTab[27075]"
//line /usr/local/go/src/encoding/json/decode.go:678
		_go_fuzz_dep_.CoverTab[27076]++

//line /usr/local/go/src/encoding/json/decode.go:681
		start := d.readIndex()
		d.rescanLiteral()
		item := d.data[start:d.readIndex()]
		key, ok := unquoteBytes(item)
		if !ok {
//line /usr/local/go/src/encoding/json/decode.go:685
			_go_fuzz_dep_.CoverTab[27090]++
									panic(phasePanicMsg)
//line /usr/local/go/src/encoding/json/decode.go:686
			// _ = "end of CoverTab[27090]"
		} else {
//line /usr/local/go/src/encoding/json/decode.go:687
			_go_fuzz_dep_.CoverTab[27091]++
//line /usr/local/go/src/encoding/json/decode.go:687
			// _ = "end of CoverTab[27091]"
//line /usr/local/go/src/encoding/json/decode.go:687
		}
//line /usr/local/go/src/encoding/json/decode.go:687
		// _ = "end of CoverTab[27076]"
//line /usr/local/go/src/encoding/json/decode.go:687
		_go_fuzz_dep_.CoverTab[27077]++

		// Figure out field corresponding to key.
		var subv reflect.Value
		destring := false

		if v.Kind() == reflect.Map {
//line /usr/local/go/src/encoding/json/decode.go:693
			_go_fuzz_dep_.CoverTab[27092]++
									elemType := t.Elem()
									if !mapElem.IsValid() {
//line /usr/local/go/src/encoding/json/decode.go:695
				_go_fuzz_dep_.CoverTab[27094]++
										mapElem = reflect.New(elemType).Elem()
//line /usr/local/go/src/encoding/json/decode.go:696
				// _ = "end of CoverTab[27094]"
			} else {
//line /usr/local/go/src/encoding/json/decode.go:697
				_go_fuzz_dep_.CoverTab[27095]++
										mapElem.Set(reflect.Zero(elemType))
//line /usr/local/go/src/encoding/json/decode.go:698
				// _ = "end of CoverTab[27095]"
			}
//line /usr/local/go/src/encoding/json/decode.go:699
			// _ = "end of CoverTab[27092]"
//line /usr/local/go/src/encoding/json/decode.go:699
			_go_fuzz_dep_.CoverTab[27093]++
									subv = mapElem
//line /usr/local/go/src/encoding/json/decode.go:700
			// _ = "end of CoverTab[27093]"
		} else {
//line /usr/local/go/src/encoding/json/decode.go:701
			_go_fuzz_dep_.CoverTab[27096]++
									var f *field
									if i, ok := fields.nameIndex[string(key)]; ok {
//line /usr/local/go/src/encoding/json/decode.go:703
				_go_fuzz_dep_.CoverTab[27098]++

										f = &fields.list[i]
//line /usr/local/go/src/encoding/json/decode.go:705
				// _ = "end of CoverTab[27098]"
			} else {
//line /usr/local/go/src/encoding/json/decode.go:706
				_go_fuzz_dep_.CoverTab[27099]++

//line /usr/local/go/src/encoding/json/decode.go:709
				for i := range fields.list {
//line /usr/local/go/src/encoding/json/decode.go:709
					_go_fuzz_dep_.CoverTab[27100]++
											ff := &fields.list[i]
											if ff.equalFold(ff.nameBytes, key) {
//line /usr/local/go/src/encoding/json/decode.go:711
						_go_fuzz_dep_.CoverTab[27101]++
												f = ff
												break
//line /usr/local/go/src/encoding/json/decode.go:713
						// _ = "end of CoverTab[27101]"
					} else {
//line /usr/local/go/src/encoding/json/decode.go:714
						_go_fuzz_dep_.CoverTab[27102]++
//line /usr/local/go/src/encoding/json/decode.go:714
						// _ = "end of CoverTab[27102]"
//line /usr/local/go/src/encoding/json/decode.go:714
					}
//line /usr/local/go/src/encoding/json/decode.go:714
					// _ = "end of CoverTab[27100]"
				}
//line /usr/local/go/src/encoding/json/decode.go:715
				// _ = "end of CoverTab[27099]"
			}
//line /usr/local/go/src/encoding/json/decode.go:716
			// _ = "end of CoverTab[27096]"
//line /usr/local/go/src/encoding/json/decode.go:716
			_go_fuzz_dep_.CoverTab[27097]++
									if f != nil {
//line /usr/local/go/src/encoding/json/decode.go:717
				_go_fuzz_dep_.CoverTab[27103]++
										subv = v
										destring = f.quoted
										for _, i := range f.index {
//line /usr/local/go/src/encoding/json/decode.go:720
					_go_fuzz_dep_.CoverTab[27106]++
											if subv.Kind() == reflect.Pointer {
//line /usr/local/go/src/encoding/json/decode.go:721
						_go_fuzz_dep_.CoverTab[27108]++
												if subv.IsNil() {
//line /usr/local/go/src/encoding/json/decode.go:722
							_go_fuzz_dep_.CoverTab[27110]++

//line /usr/local/go/src/encoding/json/decode.go:728
							if !subv.CanSet() {
//line /usr/local/go/src/encoding/json/decode.go:728
								_go_fuzz_dep_.CoverTab[27112]++
														d.saveError(fmt.Errorf("json: cannot set embedded pointer to unexported struct: %v", subv.Type().Elem()))

//line /usr/local/go/src/encoding/json/decode.go:732
								subv = reflect.Value{}
														destring = false
														break
//line /usr/local/go/src/encoding/json/decode.go:734
								// _ = "end of CoverTab[27112]"
							} else {
//line /usr/local/go/src/encoding/json/decode.go:735
								_go_fuzz_dep_.CoverTab[27113]++
//line /usr/local/go/src/encoding/json/decode.go:735
								// _ = "end of CoverTab[27113]"
//line /usr/local/go/src/encoding/json/decode.go:735
							}
//line /usr/local/go/src/encoding/json/decode.go:735
							// _ = "end of CoverTab[27110]"
//line /usr/local/go/src/encoding/json/decode.go:735
							_go_fuzz_dep_.CoverTab[27111]++
													subv.Set(reflect.New(subv.Type().Elem()))
//line /usr/local/go/src/encoding/json/decode.go:736
							// _ = "end of CoverTab[27111]"
						} else {
//line /usr/local/go/src/encoding/json/decode.go:737
							_go_fuzz_dep_.CoverTab[27114]++
//line /usr/local/go/src/encoding/json/decode.go:737
							// _ = "end of CoverTab[27114]"
//line /usr/local/go/src/encoding/json/decode.go:737
						}
//line /usr/local/go/src/encoding/json/decode.go:737
						// _ = "end of CoverTab[27108]"
//line /usr/local/go/src/encoding/json/decode.go:737
						_go_fuzz_dep_.CoverTab[27109]++
												subv = subv.Elem()
//line /usr/local/go/src/encoding/json/decode.go:738
						// _ = "end of CoverTab[27109]"
					} else {
//line /usr/local/go/src/encoding/json/decode.go:739
						_go_fuzz_dep_.CoverTab[27115]++
//line /usr/local/go/src/encoding/json/decode.go:739
						// _ = "end of CoverTab[27115]"
//line /usr/local/go/src/encoding/json/decode.go:739
					}
//line /usr/local/go/src/encoding/json/decode.go:739
					// _ = "end of CoverTab[27106]"
//line /usr/local/go/src/encoding/json/decode.go:739
					_go_fuzz_dep_.CoverTab[27107]++
											subv = subv.Field(i)
//line /usr/local/go/src/encoding/json/decode.go:740
					// _ = "end of CoverTab[27107]"
				}
//line /usr/local/go/src/encoding/json/decode.go:741
				// _ = "end of CoverTab[27103]"
//line /usr/local/go/src/encoding/json/decode.go:741
				_go_fuzz_dep_.CoverTab[27104]++
										if d.errorContext == nil {
//line /usr/local/go/src/encoding/json/decode.go:742
					_go_fuzz_dep_.CoverTab[27116]++
											d.errorContext = new(errorContext)
//line /usr/local/go/src/encoding/json/decode.go:743
					// _ = "end of CoverTab[27116]"
				} else {
//line /usr/local/go/src/encoding/json/decode.go:744
					_go_fuzz_dep_.CoverTab[27117]++
//line /usr/local/go/src/encoding/json/decode.go:744
					// _ = "end of CoverTab[27117]"
//line /usr/local/go/src/encoding/json/decode.go:744
				}
//line /usr/local/go/src/encoding/json/decode.go:744
				// _ = "end of CoverTab[27104]"
//line /usr/local/go/src/encoding/json/decode.go:744
				_go_fuzz_dep_.CoverTab[27105]++
										d.errorContext.FieldStack = append(d.errorContext.FieldStack, f.name)
										d.errorContext.Struct = t
//line /usr/local/go/src/encoding/json/decode.go:746
				// _ = "end of CoverTab[27105]"
			} else {
//line /usr/local/go/src/encoding/json/decode.go:747
				_go_fuzz_dep_.CoverTab[27118]++
//line /usr/local/go/src/encoding/json/decode.go:747
				if d.disallowUnknownFields {
//line /usr/local/go/src/encoding/json/decode.go:747
					_go_fuzz_dep_.CoverTab[27119]++
											d.saveError(fmt.Errorf("json: unknown field %q", key))
//line /usr/local/go/src/encoding/json/decode.go:748
					// _ = "end of CoverTab[27119]"
				} else {
//line /usr/local/go/src/encoding/json/decode.go:749
					_go_fuzz_dep_.CoverTab[27120]++
//line /usr/local/go/src/encoding/json/decode.go:749
					// _ = "end of CoverTab[27120]"
//line /usr/local/go/src/encoding/json/decode.go:749
				}
//line /usr/local/go/src/encoding/json/decode.go:749
				// _ = "end of CoverTab[27118]"
//line /usr/local/go/src/encoding/json/decode.go:749
			}
//line /usr/local/go/src/encoding/json/decode.go:749
			// _ = "end of CoverTab[27097]"
		}
//line /usr/local/go/src/encoding/json/decode.go:750
		// _ = "end of CoverTab[27077]"
//line /usr/local/go/src/encoding/json/decode.go:750
		_go_fuzz_dep_.CoverTab[27078]++

//line /usr/local/go/src/encoding/json/decode.go:753
		if d.opcode == scanSkipSpace {
//line /usr/local/go/src/encoding/json/decode.go:753
			_go_fuzz_dep_.CoverTab[27121]++
									d.scanWhile(scanSkipSpace)
//line /usr/local/go/src/encoding/json/decode.go:754
			// _ = "end of CoverTab[27121]"
		} else {
//line /usr/local/go/src/encoding/json/decode.go:755
			_go_fuzz_dep_.CoverTab[27122]++
//line /usr/local/go/src/encoding/json/decode.go:755
			// _ = "end of CoverTab[27122]"
//line /usr/local/go/src/encoding/json/decode.go:755
		}
//line /usr/local/go/src/encoding/json/decode.go:755
		// _ = "end of CoverTab[27078]"
//line /usr/local/go/src/encoding/json/decode.go:755
		_go_fuzz_dep_.CoverTab[27079]++
								if d.opcode != scanObjectKey {
//line /usr/local/go/src/encoding/json/decode.go:756
			_go_fuzz_dep_.CoverTab[27123]++
									panic(phasePanicMsg)
//line /usr/local/go/src/encoding/json/decode.go:757
			// _ = "end of CoverTab[27123]"
		} else {
//line /usr/local/go/src/encoding/json/decode.go:758
			_go_fuzz_dep_.CoverTab[27124]++
//line /usr/local/go/src/encoding/json/decode.go:758
			// _ = "end of CoverTab[27124]"
//line /usr/local/go/src/encoding/json/decode.go:758
		}
//line /usr/local/go/src/encoding/json/decode.go:758
		// _ = "end of CoverTab[27079]"
//line /usr/local/go/src/encoding/json/decode.go:758
		_go_fuzz_dep_.CoverTab[27080]++
								d.scanWhile(scanSkipSpace)

								if destring {
//line /usr/local/go/src/encoding/json/decode.go:761
			_go_fuzz_dep_.CoverTab[27125]++
									switch qv := d.valueQuoted().(type) {
			case nil:
//line /usr/local/go/src/encoding/json/decode.go:763
				_go_fuzz_dep_.CoverTab[27126]++
										if err := d.literalStore(nullLiteral, subv, false); err != nil {
//line /usr/local/go/src/encoding/json/decode.go:764
					_go_fuzz_dep_.CoverTab[27129]++
											return err
//line /usr/local/go/src/encoding/json/decode.go:765
					// _ = "end of CoverTab[27129]"
				} else {
//line /usr/local/go/src/encoding/json/decode.go:766
					_go_fuzz_dep_.CoverTab[27130]++
//line /usr/local/go/src/encoding/json/decode.go:766
					// _ = "end of CoverTab[27130]"
//line /usr/local/go/src/encoding/json/decode.go:766
				}
//line /usr/local/go/src/encoding/json/decode.go:766
				// _ = "end of CoverTab[27126]"
			case string:
//line /usr/local/go/src/encoding/json/decode.go:767
				_go_fuzz_dep_.CoverTab[27127]++
										if err := d.literalStore([]byte(qv), subv, true); err != nil {
//line /usr/local/go/src/encoding/json/decode.go:768
					_go_fuzz_dep_.CoverTab[27131]++
											return err
//line /usr/local/go/src/encoding/json/decode.go:769
					// _ = "end of CoverTab[27131]"
				} else {
//line /usr/local/go/src/encoding/json/decode.go:770
					_go_fuzz_dep_.CoverTab[27132]++
//line /usr/local/go/src/encoding/json/decode.go:770
					// _ = "end of CoverTab[27132]"
//line /usr/local/go/src/encoding/json/decode.go:770
				}
//line /usr/local/go/src/encoding/json/decode.go:770
				// _ = "end of CoverTab[27127]"
			default:
//line /usr/local/go/src/encoding/json/decode.go:771
				_go_fuzz_dep_.CoverTab[27128]++
										d.saveError(fmt.Errorf("json: invalid use of ,string struct tag, trying to unmarshal unquoted value into %v", subv.Type()))
//line /usr/local/go/src/encoding/json/decode.go:772
				// _ = "end of CoverTab[27128]"
			}
//line /usr/local/go/src/encoding/json/decode.go:773
			// _ = "end of CoverTab[27125]"
		} else {
//line /usr/local/go/src/encoding/json/decode.go:774
			_go_fuzz_dep_.CoverTab[27133]++
									if err := d.value(subv); err != nil {
//line /usr/local/go/src/encoding/json/decode.go:775
				_go_fuzz_dep_.CoverTab[27134]++
										return err
//line /usr/local/go/src/encoding/json/decode.go:776
				// _ = "end of CoverTab[27134]"
			} else {
//line /usr/local/go/src/encoding/json/decode.go:777
				_go_fuzz_dep_.CoverTab[27135]++
//line /usr/local/go/src/encoding/json/decode.go:777
				// _ = "end of CoverTab[27135]"
//line /usr/local/go/src/encoding/json/decode.go:777
			}
//line /usr/local/go/src/encoding/json/decode.go:777
			// _ = "end of CoverTab[27133]"
		}
//line /usr/local/go/src/encoding/json/decode.go:778
		// _ = "end of CoverTab[27080]"
//line /usr/local/go/src/encoding/json/decode.go:778
		_go_fuzz_dep_.CoverTab[27081]++

//line /usr/local/go/src/encoding/json/decode.go:782
		if v.Kind() == reflect.Map {
//line /usr/local/go/src/encoding/json/decode.go:782
			_go_fuzz_dep_.CoverTab[27136]++
									kt := t.Key()
									var kv reflect.Value
									switch {
			case reflect.PointerTo(kt).Implements(textUnmarshalerType):
//line /usr/local/go/src/encoding/json/decode.go:786
				_go_fuzz_dep_.CoverTab[27138]++
										kv = reflect.New(kt)
										if err := d.literalStore(item, kv, true); err != nil {
//line /usr/local/go/src/encoding/json/decode.go:788
					_go_fuzz_dep_.CoverTab[27142]++
											return err
//line /usr/local/go/src/encoding/json/decode.go:789
					// _ = "end of CoverTab[27142]"
				} else {
//line /usr/local/go/src/encoding/json/decode.go:790
					_go_fuzz_dep_.CoverTab[27143]++
//line /usr/local/go/src/encoding/json/decode.go:790
					// _ = "end of CoverTab[27143]"
//line /usr/local/go/src/encoding/json/decode.go:790
				}
//line /usr/local/go/src/encoding/json/decode.go:790
				// _ = "end of CoverTab[27138]"
//line /usr/local/go/src/encoding/json/decode.go:790
				_go_fuzz_dep_.CoverTab[27139]++
										kv = kv.Elem()
//line /usr/local/go/src/encoding/json/decode.go:791
				// _ = "end of CoverTab[27139]"
			case kt.Kind() == reflect.String:
//line /usr/local/go/src/encoding/json/decode.go:792
				_go_fuzz_dep_.CoverTab[27140]++
										kv = reflect.ValueOf(key).Convert(kt)
//line /usr/local/go/src/encoding/json/decode.go:793
				// _ = "end of CoverTab[27140]"
			default:
//line /usr/local/go/src/encoding/json/decode.go:794
				_go_fuzz_dep_.CoverTab[27141]++
										switch kt.Kind() {
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
//line /usr/local/go/src/encoding/json/decode.go:796
					_go_fuzz_dep_.CoverTab[27144]++
											s := string(key)
											n, err := strconv.ParseInt(s, 10, 64)
											if err != nil || func() bool {
//line /usr/local/go/src/encoding/json/decode.go:799
						_go_fuzz_dep_.CoverTab[27149]++
//line /usr/local/go/src/encoding/json/decode.go:799
						return reflect.Zero(kt).OverflowInt(n)
//line /usr/local/go/src/encoding/json/decode.go:799
						// _ = "end of CoverTab[27149]"
//line /usr/local/go/src/encoding/json/decode.go:799
					}() {
//line /usr/local/go/src/encoding/json/decode.go:799
						_go_fuzz_dep_.CoverTab[27150]++
												d.saveError(&UnmarshalTypeError{Value: "number " + s, Type: kt, Offset: int64(start + 1)})
												break
//line /usr/local/go/src/encoding/json/decode.go:801
						// _ = "end of CoverTab[27150]"
					} else {
//line /usr/local/go/src/encoding/json/decode.go:802
						_go_fuzz_dep_.CoverTab[27151]++
//line /usr/local/go/src/encoding/json/decode.go:802
						// _ = "end of CoverTab[27151]"
//line /usr/local/go/src/encoding/json/decode.go:802
					}
//line /usr/local/go/src/encoding/json/decode.go:802
					// _ = "end of CoverTab[27144]"
//line /usr/local/go/src/encoding/json/decode.go:802
					_go_fuzz_dep_.CoverTab[27145]++
											kv = reflect.ValueOf(n).Convert(kt)
//line /usr/local/go/src/encoding/json/decode.go:803
					// _ = "end of CoverTab[27145]"
				case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
//line /usr/local/go/src/encoding/json/decode.go:804
					_go_fuzz_dep_.CoverTab[27146]++
											s := string(key)
											n, err := strconv.ParseUint(s, 10, 64)
											if err != nil || func() bool {
//line /usr/local/go/src/encoding/json/decode.go:807
						_go_fuzz_dep_.CoverTab[27152]++
//line /usr/local/go/src/encoding/json/decode.go:807
						return reflect.Zero(kt).OverflowUint(n)
//line /usr/local/go/src/encoding/json/decode.go:807
						// _ = "end of CoverTab[27152]"
//line /usr/local/go/src/encoding/json/decode.go:807
					}() {
//line /usr/local/go/src/encoding/json/decode.go:807
						_go_fuzz_dep_.CoverTab[27153]++
												d.saveError(&UnmarshalTypeError{Value: "number " + s, Type: kt, Offset: int64(start + 1)})
												break
//line /usr/local/go/src/encoding/json/decode.go:809
						// _ = "end of CoverTab[27153]"
					} else {
//line /usr/local/go/src/encoding/json/decode.go:810
						_go_fuzz_dep_.CoverTab[27154]++
//line /usr/local/go/src/encoding/json/decode.go:810
						// _ = "end of CoverTab[27154]"
//line /usr/local/go/src/encoding/json/decode.go:810
					}
//line /usr/local/go/src/encoding/json/decode.go:810
					// _ = "end of CoverTab[27146]"
//line /usr/local/go/src/encoding/json/decode.go:810
					_go_fuzz_dep_.CoverTab[27147]++
											kv = reflect.ValueOf(n).Convert(kt)
//line /usr/local/go/src/encoding/json/decode.go:811
					// _ = "end of CoverTab[27147]"
				default:
//line /usr/local/go/src/encoding/json/decode.go:812
					_go_fuzz_dep_.CoverTab[27148]++
											panic("json: Unexpected key type")
//line /usr/local/go/src/encoding/json/decode.go:813
					// _ = "end of CoverTab[27148]"
				}
//line /usr/local/go/src/encoding/json/decode.go:814
				// _ = "end of CoverTab[27141]"
			}
//line /usr/local/go/src/encoding/json/decode.go:815
			// _ = "end of CoverTab[27136]"
//line /usr/local/go/src/encoding/json/decode.go:815
			_go_fuzz_dep_.CoverTab[27137]++
									if kv.IsValid() {
//line /usr/local/go/src/encoding/json/decode.go:816
				_go_fuzz_dep_.CoverTab[27155]++
										v.SetMapIndex(kv, subv)
//line /usr/local/go/src/encoding/json/decode.go:817
				// _ = "end of CoverTab[27155]"
			} else {
//line /usr/local/go/src/encoding/json/decode.go:818
				_go_fuzz_dep_.CoverTab[27156]++
//line /usr/local/go/src/encoding/json/decode.go:818
				// _ = "end of CoverTab[27156]"
//line /usr/local/go/src/encoding/json/decode.go:818
			}
//line /usr/local/go/src/encoding/json/decode.go:818
			// _ = "end of CoverTab[27137]"
		} else {
//line /usr/local/go/src/encoding/json/decode.go:819
			_go_fuzz_dep_.CoverTab[27157]++
//line /usr/local/go/src/encoding/json/decode.go:819
			// _ = "end of CoverTab[27157]"
//line /usr/local/go/src/encoding/json/decode.go:819
		}
//line /usr/local/go/src/encoding/json/decode.go:819
		// _ = "end of CoverTab[27081]"
//line /usr/local/go/src/encoding/json/decode.go:819
		_go_fuzz_dep_.CoverTab[27082]++

//line /usr/local/go/src/encoding/json/decode.go:822
		if d.opcode == scanSkipSpace {
//line /usr/local/go/src/encoding/json/decode.go:822
			_go_fuzz_dep_.CoverTab[27158]++
									d.scanWhile(scanSkipSpace)
//line /usr/local/go/src/encoding/json/decode.go:823
			// _ = "end of CoverTab[27158]"
		} else {
//line /usr/local/go/src/encoding/json/decode.go:824
			_go_fuzz_dep_.CoverTab[27159]++
//line /usr/local/go/src/encoding/json/decode.go:824
			// _ = "end of CoverTab[27159]"
//line /usr/local/go/src/encoding/json/decode.go:824
		}
//line /usr/local/go/src/encoding/json/decode.go:824
		// _ = "end of CoverTab[27082]"
//line /usr/local/go/src/encoding/json/decode.go:824
		_go_fuzz_dep_.CoverTab[27083]++
								if d.errorContext != nil {
//line /usr/local/go/src/encoding/json/decode.go:825
			_go_fuzz_dep_.CoverTab[27160]++

//line /usr/local/go/src/encoding/json/decode.go:829
			d.errorContext.FieldStack = d.errorContext.FieldStack[:len(origErrorContext.FieldStack)]
									d.errorContext.Struct = origErrorContext.Struct
//line /usr/local/go/src/encoding/json/decode.go:830
			// _ = "end of CoverTab[27160]"
		} else {
//line /usr/local/go/src/encoding/json/decode.go:831
			_go_fuzz_dep_.CoverTab[27161]++
//line /usr/local/go/src/encoding/json/decode.go:831
			// _ = "end of CoverTab[27161]"
//line /usr/local/go/src/encoding/json/decode.go:831
		}
//line /usr/local/go/src/encoding/json/decode.go:831
		// _ = "end of CoverTab[27083]"
//line /usr/local/go/src/encoding/json/decode.go:831
		_go_fuzz_dep_.CoverTab[27084]++
								if d.opcode == scanEndObject {
//line /usr/local/go/src/encoding/json/decode.go:832
			_go_fuzz_dep_.CoverTab[27162]++
									break
//line /usr/local/go/src/encoding/json/decode.go:833
			// _ = "end of CoverTab[27162]"
		} else {
//line /usr/local/go/src/encoding/json/decode.go:834
			_go_fuzz_dep_.CoverTab[27163]++
//line /usr/local/go/src/encoding/json/decode.go:834
			// _ = "end of CoverTab[27163]"
//line /usr/local/go/src/encoding/json/decode.go:834
		}
//line /usr/local/go/src/encoding/json/decode.go:834
		// _ = "end of CoverTab[27084]"
//line /usr/local/go/src/encoding/json/decode.go:834
		_go_fuzz_dep_.CoverTab[27085]++
								if d.opcode != scanObjectValue {
//line /usr/local/go/src/encoding/json/decode.go:835
			_go_fuzz_dep_.CoverTab[27164]++
									panic(phasePanicMsg)
//line /usr/local/go/src/encoding/json/decode.go:836
			// _ = "end of CoverTab[27164]"
		} else {
//line /usr/local/go/src/encoding/json/decode.go:837
			_go_fuzz_dep_.CoverTab[27165]++
//line /usr/local/go/src/encoding/json/decode.go:837
			// _ = "end of CoverTab[27165]"
//line /usr/local/go/src/encoding/json/decode.go:837
		}
//line /usr/local/go/src/encoding/json/decode.go:837
		// _ = "end of CoverTab[27085]"
	}
//line /usr/local/go/src/encoding/json/decode.go:838
	// _ = "end of CoverTab[27053]"
//line /usr/local/go/src/encoding/json/decode.go:838
	_go_fuzz_dep_.CoverTab[27054]++
							return nil
//line /usr/local/go/src/encoding/json/decode.go:839
	// _ = "end of CoverTab[27054]"
}

// convertNumber converts the number literal s to a float64 or a Number
//line /usr/local/go/src/encoding/json/decode.go:842
// depending on the setting of d.useNumber.
//line /usr/local/go/src/encoding/json/decode.go:844
func (d *decodeState) convertNumber(s string) (any, error) {
//line /usr/local/go/src/encoding/json/decode.go:844
	_go_fuzz_dep_.CoverTab[27166]++
							if d.useNumber {
//line /usr/local/go/src/encoding/json/decode.go:845
		_go_fuzz_dep_.CoverTab[27169]++
								return Number(s), nil
//line /usr/local/go/src/encoding/json/decode.go:846
		// _ = "end of CoverTab[27169]"
	} else {
//line /usr/local/go/src/encoding/json/decode.go:847
		_go_fuzz_dep_.CoverTab[27170]++
//line /usr/local/go/src/encoding/json/decode.go:847
		// _ = "end of CoverTab[27170]"
//line /usr/local/go/src/encoding/json/decode.go:847
	}
//line /usr/local/go/src/encoding/json/decode.go:847
	// _ = "end of CoverTab[27166]"
//line /usr/local/go/src/encoding/json/decode.go:847
	_go_fuzz_dep_.CoverTab[27167]++
							f, err := strconv.ParseFloat(s, 64)
							if err != nil {
//line /usr/local/go/src/encoding/json/decode.go:849
		_go_fuzz_dep_.CoverTab[27171]++
								return nil, &UnmarshalTypeError{Value: "number " + s, Type: reflect.TypeOf(0.0), Offset: int64(d.off)}
//line /usr/local/go/src/encoding/json/decode.go:850
		// _ = "end of CoverTab[27171]"
	} else {
//line /usr/local/go/src/encoding/json/decode.go:851
		_go_fuzz_dep_.CoverTab[27172]++
//line /usr/local/go/src/encoding/json/decode.go:851
		// _ = "end of CoverTab[27172]"
//line /usr/local/go/src/encoding/json/decode.go:851
	}
//line /usr/local/go/src/encoding/json/decode.go:851
	// _ = "end of CoverTab[27167]"
//line /usr/local/go/src/encoding/json/decode.go:851
	_go_fuzz_dep_.CoverTab[27168]++
							return f, nil
//line /usr/local/go/src/encoding/json/decode.go:852
	// _ = "end of CoverTab[27168]"
}

var numberType = reflect.TypeOf(Number(""))

// literalStore decodes a literal stored in item into v.
//line /usr/local/go/src/encoding/json/decode.go:857
//
//line /usr/local/go/src/encoding/json/decode.go:857
// fromQuoted indicates whether this literal came from unwrapping a
//line /usr/local/go/src/encoding/json/decode.go:857
// string from the ",string" struct tag option. this is used only to
//line /usr/local/go/src/encoding/json/decode.go:857
// produce more helpful error messages.
//line /usr/local/go/src/encoding/json/decode.go:862
func (d *decodeState) literalStore(item []byte, v reflect.Value, fromQuoted bool) error {
//line /usr/local/go/src/encoding/json/decode.go:862
	_go_fuzz_dep_.CoverTab[27173]++

							if len(item) == 0 {
//line /usr/local/go/src/encoding/json/decode.go:864
		_go_fuzz_dep_.CoverTab[27178]++

								d.saveError(fmt.Errorf("json: invalid use of ,string struct tag, trying to unmarshal %q into %v", item, v.Type()))
								return nil
//line /usr/local/go/src/encoding/json/decode.go:867
		// _ = "end of CoverTab[27178]"
	} else {
//line /usr/local/go/src/encoding/json/decode.go:868
		_go_fuzz_dep_.CoverTab[27179]++
//line /usr/local/go/src/encoding/json/decode.go:868
		// _ = "end of CoverTab[27179]"
//line /usr/local/go/src/encoding/json/decode.go:868
	}
//line /usr/local/go/src/encoding/json/decode.go:868
	// _ = "end of CoverTab[27173]"
//line /usr/local/go/src/encoding/json/decode.go:868
	_go_fuzz_dep_.CoverTab[27174]++
							isNull := item[0] == 'n'
							u, ut, pv := indirect(v, isNull)
							if u != nil {
//line /usr/local/go/src/encoding/json/decode.go:871
		_go_fuzz_dep_.CoverTab[27180]++
								return u.UnmarshalJSON(item)
//line /usr/local/go/src/encoding/json/decode.go:872
		// _ = "end of CoverTab[27180]"
	} else {
//line /usr/local/go/src/encoding/json/decode.go:873
		_go_fuzz_dep_.CoverTab[27181]++
//line /usr/local/go/src/encoding/json/decode.go:873
		// _ = "end of CoverTab[27181]"
//line /usr/local/go/src/encoding/json/decode.go:873
	}
//line /usr/local/go/src/encoding/json/decode.go:873
	// _ = "end of CoverTab[27174]"
//line /usr/local/go/src/encoding/json/decode.go:873
	_go_fuzz_dep_.CoverTab[27175]++
							if ut != nil {
//line /usr/local/go/src/encoding/json/decode.go:874
		_go_fuzz_dep_.CoverTab[27182]++
								if item[0] != '"' {
//line /usr/local/go/src/encoding/json/decode.go:875
			_go_fuzz_dep_.CoverTab[27185]++
									if fromQuoted {
//line /usr/local/go/src/encoding/json/decode.go:876
				_go_fuzz_dep_.CoverTab[27188]++
										d.saveError(fmt.Errorf("json: invalid use of ,string struct tag, trying to unmarshal %q into %v", item, v.Type()))
										return nil
//line /usr/local/go/src/encoding/json/decode.go:878
				// _ = "end of CoverTab[27188]"
			} else {
//line /usr/local/go/src/encoding/json/decode.go:879
				_go_fuzz_dep_.CoverTab[27189]++
//line /usr/local/go/src/encoding/json/decode.go:879
				// _ = "end of CoverTab[27189]"
//line /usr/local/go/src/encoding/json/decode.go:879
			}
//line /usr/local/go/src/encoding/json/decode.go:879
			// _ = "end of CoverTab[27185]"
//line /usr/local/go/src/encoding/json/decode.go:879
			_go_fuzz_dep_.CoverTab[27186]++
									val := "number"
									switch item[0] {
			case 'n':
//line /usr/local/go/src/encoding/json/decode.go:882
				_go_fuzz_dep_.CoverTab[27190]++
										val = "null"
//line /usr/local/go/src/encoding/json/decode.go:883
				// _ = "end of CoverTab[27190]"
			case 't', 'f':
//line /usr/local/go/src/encoding/json/decode.go:884
				_go_fuzz_dep_.CoverTab[27191]++
										val = "bool"
//line /usr/local/go/src/encoding/json/decode.go:885
				// _ = "end of CoverTab[27191]"
//line /usr/local/go/src/encoding/json/decode.go:885
			default:
//line /usr/local/go/src/encoding/json/decode.go:885
				_go_fuzz_dep_.CoverTab[27192]++
//line /usr/local/go/src/encoding/json/decode.go:885
				// _ = "end of CoverTab[27192]"
			}
//line /usr/local/go/src/encoding/json/decode.go:886
			// _ = "end of CoverTab[27186]"
//line /usr/local/go/src/encoding/json/decode.go:886
			_go_fuzz_dep_.CoverTab[27187]++
									d.saveError(&UnmarshalTypeError{Value: val, Type: v.Type(), Offset: int64(d.readIndex())})
									return nil
//line /usr/local/go/src/encoding/json/decode.go:888
			// _ = "end of CoverTab[27187]"
		} else {
//line /usr/local/go/src/encoding/json/decode.go:889
			_go_fuzz_dep_.CoverTab[27193]++
//line /usr/local/go/src/encoding/json/decode.go:889
			// _ = "end of CoverTab[27193]"
//line /usr/local/go/src/encoding/json/decode.go:889
		}
//line /usr/local/go/src/encoding/json/decode.go:889
		// _ = "end of CoverTab[27182]"
//line /usr/local/go/src/encoding/json/decode.go:889
		_go_fuzz_dep_.CoverTab[27183]++
								s, ok := unquoteBytes(item)
								if !ok {
//line /usr/local/go/src/encoding/json/decode.go:891
			_go_fuzz_dep_.CoverTab[27194]++
									if fromQuoted {
//line /usr/local/go/src/encoding/json/decode.go:892
				_go_fuzz_dep_.CoverTab[27196]++
										return fmt.Errorf("json: invalid use of ,string struct tag, trying to unmarshal %q into %v", item, v.Type())
//line /usr/local/go/src/encoding/json/decode.go:893
				// _ = "end of CoverTab[27196]"
			} else {
//line /usr/local/go/src/encoding/json/decode.go:894
				_go_fuzz_dep_.CoverTab[27197]++
//line /usr/local/go/src/encoding/json/decode.go:894
				// _ = "end of CoverTab[27197]"
//line /usr/local/go/src/encoding/json/decode.go:894
			}
//line /usr/local/go/src/encoding/json/decode.go:894
			// _ = "end of CoverTab[27194]"
//line /usr/local/go/src/encoding/json/decode.go:894
			_go_fuzz_dep_.CoverTab[27195]++
									panic(phasePanicMsg)
//line /usr/local/go/src/encoding/json/decode.go:895
			// _ = "end of CoverTab[27195]"
		} else {
//line /usr/local/go/src/encoding/json/decode.go:896
			_go_fuzz_dep_.CoverTab[27198]++
//line /usr/local/go/src/encoding/json/decode.go:896
			// _ = "end of CoverTab[27198]"
//line /usr/local/go/src/encoding/json/decode.go:896
		}
//line /usr/local/go/src/encoding/json/decode.go:896
		// _ = "end of CoverTab[27183]"
//line /usr/local/go/src/encoding/json/decode.go:896
		_go_fuzz_dep_.CoverTab[27184]++
								return ut.UnmarshalText(s)
//line /usr/local/go/src/encoding/json/decode.go:897
		// _ = "end of CoverTab[27184]"
	} else {
//line /usr/local/go/src/encoding/json/decode.go:898
		_go_fuzz_dep_.CoverTab[27199]++
//line /usr/local/go/src/encoding/json/decode.go:898
		// _ = "end of CoverTab[27199]"
//line /usr/local/go/src/encoding/json/decode.go:898
	}
//line /usr/local/go/src/encoding/json/decode.go:898
	// _ = "end of CoverTab[27175]"
//line /usr/local/go/src/encoding/json/decode.go:898
	_go_fuzz_dep_.CoverTab[27176]++

							v = pv

							switch c := item[0]; c {
	case 'n':
//line /usr/local/go/src/encoding/json/decode.go:903
		_go_fuzz_dep_.CoverTab[27200]++

//line /usr/local/go/src/encoding/json/decode.go:906
		if fromQuoted && func() bool {
//line /usr/local/go/src/encoding/json/decode.go:906
			_go_fuzz_dep_.CoverTab[27208]++
//line /usr/local/go/src/encoding/json/decode.go:906
			return string(item) != "null"
//line /usr/local/go/src/encoding/json/decode.go:906
			// _ = "end of CoverTab[27208]"
//line /usr/local/go/src/encoding/json/decode.go:906
		}() {
//line /usr/local/go/src/encoding/json/decode.go:906
			_go_fuzz_dep_.CoverTab[27209]++
									d.saveError(fmt.Errorf("json: invalid use of ,string struct tag, trying to unmarshal %q into %v", item, v.Type()))
									break
//line /usr/local/go/src/encoding/json/decode.go:908
			// _ = "end of CoverTab[27209]"
		} else {
//line /usr/local/go/src/encoding/json/decode.go:909
			_go_fuzz_dep_.CoverTab[27210]++
//line /usr/local/go/src/encoding/json/decode.go:909
			// _ = "end of CoverTab[27210]"
//line /usr/local/go/src/encoding/json/decode.go:909
		}
//line /usr/local/go/src/encoding/json/decode.go:909
		// _ = "end of CoverTab[27200]"
//line /usr/local/go/src/encoding/json/decode.go:909
		_go_fuzz_dep_.CoverTab[27201]++
								switch v.Kind() {
		case reflect.Interface, reflect.Pointer, reflect.Map, reflect.Slice:
//line /usr/local/go/src/encoding/json/decode.go:911
			_go_fuzz_dep_.CoverTab[27211]++
									v.Set(reflect.Zero(v.Type()))
//line /usr/local/go/src/encoding/json/decode.go:912
			// _ = "end of CoverTab[27211]"
//line /usr/local/go/src/encoding/json/decode.go:912
		default:
//line /usr/local/go/src/encoding/json/decode.go:912
			_go_fuzz_dep_.CoverTab[27212]++
//line /usr/local/go/src/encoding/json/decode.go:912
			// _ = "end of CoverTab[27212]"

		}
//line /usr/local/go/src/encoding/json/decode.go:914
		// _ = "end of CoverTab[27201]"
	case 't', 'f':
//line /usr/local/go/src/encoding/json/decode.go:915
		_go_fuzz_dep_.CoverTab[27202]++
								value := item[0] == 't'

//line /usr/local/go/src/encoding/json/decode.go:919
		if fromQuoted && func() bool {
//line /usr/local/go/src/encoding/json/decode.go:919
			_go_fuzz_dep_.CoverTab[27213]++
//line /usr/local/go/src/encoding/json/decode.go:919
			return string(item) != "true"
//line /usr/local/go/src/encoding/json/decode.go:919
			// _ = "end of CoverTab[27213]"
//line /usr/local/go/src/encoding/json/decode.go:919
		}() && func() bool {
//line /usr/local/go/src/encoding/json/decode.go:919
			_go_fuzz_dep_.CoverTab[27214]++
//line /usr/local/go/src/encoding/json/decode.go:919
			return string(item) != "false"
//line /usr/local/go/src/encoding/json/decode.go:919
			// _ = "end of CoverTab[27214]"
//line /usr/local/go/src/encoding/json/decode.go:919
		}() {
//line /usr/local/go/src/encoding/json/decode.go:919
			_go_fuzz_dep_.CoverTab[27215]++
									d.saveError(fmt.Errorf("json: invalid use of ,string struct tag, trying to unmarshal %q into %v", item, v.Type()))
									break
//line /usr/local/go/src/encoding/json/decode.go:921
			// _ = "end of CoverTab[27215]"
		} else {
//line /usr/local/go/src/encoding/json/decode.go:922
			_go_fuzz_dep_.CoverTab[27216]++
//line /usr/local/go/src/encoding/json/decode.go:922
			// _ = "end of CoverTab[27216]"
//line /usr/local/go/src/encoding/json/decode.go:922
		}
//line /usr/local/go/src/encoding/json/decode.go:922
		// _ = "end of CoverTab[27202]"
//line /usr/local/go/src/encoding/json/decode.go:922
		_go_fuzz_dep_.CoverTab[27203]++
								switch v.Kind() {
		default:
//line /usr/local/go/src/encoding/json/decode.go:924
			_go_fuzz_dep_.CoverTab[27217]++
									if fromQuoted {
//line /usr/local/go/src/encoding/json/decode.go:925
				_go_fuzz_dep_.CoverTab[27220]++
										d.saveError(fmt.Errorf("json: invalid use of ,string struct tag, trying to unmarshal %q into %v", item, v.Type()))
//line /usr/local/go/src/encoding/json/decode.go:926
				// _ = "end of CoverTab[27220]"
			} else {
//line /usr/local/go/src/encoding/json/decode.go:927
				_go_fuzz_dep_.CoverTab[27221]++
										d.saveError(&UnmarshalTypeError{Value: "bool", Type: v.Type(), Offset: int64(d.readIndex())})
//line /usr/local/go/src/encoding/json/decode.go:928
				// _ = "end of CoverTab[27221]"
			}
//line /usr/local/go/src/encoding/json/decode.go:929
			// _ = "end of CoverTab[27217]"
		case reflect.Bool:
//line /usr/local/go/src/encoding/json/decode.go:930
			_go_fuzz_dep_.CoverTab[27218]++
									v.SetBool(value)
//line /usr/local/go/src/encoding/json/decode.go:931
			// _ = "end of CoverTab[27218]"
		case reflect.Interface:
//line /usr/local/go/src/encoding/json/decode.go:932
			_go_fuzz_dep_.CoverTab[27219]++
									if v.NumMethod() == 0 {
//line /usr/local/go/src/encoding/json/decode.go:933
				_go_fuzz_dep_.CoverTab[27222]++
										v.Set(reflect.ValueOf(value))
//line /usr/local/go/src/encoding/json/decode.go:934
				// _ = "end of CoverTab[27222]"
			} else {
//line /usr/local/go/src/encoding/json/decode.go:935
				_go_fuzz_dep_.CoverTab[27223]++
										d.saveError(&UnmarshalTypeError{Value: "bool", Type: v.Type(), Offset: int64(d.readIndex())})
//line /usr/local/go/src/encoding/json/decode.go:936
				// _ = "end of CoverTab[27223]"
			}
//line /usr/local/go/src/encoding/json/decode.go:937
			// _ = "end of CoverTab[27219]"
		}
//line /usr/local/go/src/encoding/json/decode.go:938
		// _ = "end of CoverTab[27203]"

	case '"':
//line /usr/local/go/src/encoding/json/decode.go:940
		_go_fuzz_dep_.CoverTab[27204]++
								s, ok := unquoteBytes(item)
								if !ok {
//line /usr/local/go/src/encoding/json/decode.go:942
			_go_fuzz_dep_.CoverTab[27224]++
									if fromQuoted {
//line /usr/local/go/src/encoding/json/decode.go:943
				_go_fuzz_dep_.CoverTab[27226]++
										return fmt.Errorf("json: invalid use of ,string struct tag, trying to unmarshal %q into %v", item, v.Type())
//line /usr/local/go/src/encoding/json/decode.go:944
				// _ = "end of CoverTab[27226]"
			} else {
//line /usr/local/go/src/encoding/json/decode.go:945
				_go_fuzz_dep_.CoverTab[27227]++
//line /usr/local/go/src/encoding/json/decode.go:945
				// _ = "end of CoverTab[27227]"
//line /usr/local/go/src/encoding/json/decode.go:945
			}
//line /usr/local/go/src/encoding/json/decode.go:945
			// _ = "end of CoverTab[27224]"
//line /usr/local/go/src/encoding/json/decode.go:945
			_go_fuzz_dep_.CoverTab[27225]++
									panic(phasePanicMsg)
//line /usr/local/go/src/encoding/json/decode.go:946
			// _ = "end of CoverTab[27225]"
		} else {
//line /usr/local/go/src/encoding/json/decode.go:947
			_go_fuzz_dep_.CoverTab[27228]++
//line /usr/local/go/src/encoding/json/decode.go:947
			// _ = "end of CoverTab[27228]"
//line /usr/local/go/src/encoding/json/decode.go:947
		}
//line /usr/local/go/src/encoding/json/decode.go:947
		// _ = "end of CoverTab[27204]"
//line /usr/local/go/src/encoding/json/decode.go:947
		_go_fuzz_dep_.CoverTab[27205]++
								switch v.Kind() {
		default:
//line /usr/local/go/src/encoding/json/decode.go:949
			_go_fuzz_dep_.CoverTab[27229]++
									d.saveError(&UnmarshalTypeError{Value: "string", Type: v.Type(), Offset: int64(d.readIndex())})
//line /usr/local/go/src/encoding/json/decode.go:950
			// _ = "end of CoverTab[27229]"
		case reflect.Slice:
//line /usr/local/go/src/encoding/json/decode.go:951
			_go_fuzz_dep_.CoverTab[27230]++
									if v.Type().Elem().Kind() != reflect.Uint8 {
//line /usr/local/go/src/encoding/json/decode.go:952
				_go_fuzz_dep_.CoverTab[27236]++
										d.saveError(&UnmarshalTypeError{Value: "string", Type: v.Type(), Offset: int64(d.readIndex())})
										break
//line /usr/local/go/src/encoding/json/decode.go:954
				// _ = "end of CoverTab[27236]"
			} else {
//line /usr/local/go/src/encoding/json/decode.go:955
				_go_fuzz_dep_.CoverTab[27237]++
//line /usr/local/go/src/encoding/json/decode.go:955
				// _ = "end of CoverTab[27237]"
//line /usr/local/go/src/encoding/json/decode.go:955
			}
//line /usr/local/go/src/encoding/json/decode.go:955
			// _ = "end of CoverTab[27230]"
//line /usr/local/go/src/encoding/json/decode.go:955
			_go_fuzz_dep_.CoverTab[27231]++
									b := make([]byte, base64.StdEncoding.DecodedLen(len(s)))
									n, err := base64.StdEncoding.Decode(b, s)
									if err != nil {
//line /usr/local/go/src/encoding/json/decode.go:958
				_go_fuzz_dep_.CoverTab[27238]++
										d.saveError(err)
										break
//line /usr/local/go/src/encoding/json/decode.go:960
				// _ = "end of CoverTab[27238]"
			} else {
//line /usr/local/go/src/encoding/json/decode.go:961
				_go_fuzz_dep_.CoverTab[27239]++
//line /usr/local/go/src/encoding/json/decode.go:961
				// _ = "end of CoverTab[27239]"
//line /usr/local/go/src/encoding/json/decode.go:961
			}
//line /usr/local/go/src/encoding/json/decode.go:961
			// _ = "end of CoverTab[27231]"
//line /usr/local/go/src/encoding/json/decode.go:961
			_go_fuzz_dep_.CoverTab[27232]++
									v.SetBytes(b[:n])
//line /usr/local/go/src/encoding/json/decode.go:962
			// _ = "end of CoverTab[27232]"
		case reflect.String:
//line /usr/local/go/src/encoding/json/decode.go:963
			_go_fuzz_dep_.CoverTab[27233]++
									if v.Type() == numberType && func() bool {
//line /usr/local/go/src/encoding/json/decode.go:964
				_go_fuzz_dep_.CoverTab[27240]++
//line /usr/local/go/src/encoding/json/decode.go:964
				return !isValidNumber(string(s))
//line /usr/local/go/src/encoding/json/decode.go:964
				// _ = "end of CoverTab[27240]"
//line /usr/local/go/src/encoding/json/decode.go:964
			}() {
//line /usr/local/go/src/encoding/json/decode.go:964
				_go_fuzz_dep_.CoverTab[27241]++
										return fmt.Errorf("json: invalid number literal, trying to unmarshal %q into Number", item)
//line /usr/local/go/src/encoding/json/decode.go:965
				// _ = "end of CoverTab[27241]"
			} else {
//line /usr/local/go/src/encoding/json/decode.go:966
				_go_fuzz_dep_.CoverTab[27242]++
//line /usr/local/go/src/encoding/json/decode.go:966
				// _ = "end of CoverTab[27242]"
//line /usr/local/go/src/encoding/json/decode.go:966
			}
//line /usr/local/go/src/encoding/json/decode.go:966
			// _ = "end of CoverTab[27233]"
//line /usr/local/go/src/encoding/json/decode.go:966
			_go_fuzz_dep_.CoverTab[27234]++
									v.SetString(string(s))
//line /usr/local/go/src/encoding/json/decode.go:967
			// _ = "end of CoverTab[27234]"
		case reflect.Interface:
//line /usr/local/go/src/encoding/json/decode.go:968
			_go_fuzz_dep_.CoverTab[27235]++
									if v.NumMethod() == 0 {
//line /usr/local/go/src/encoding/json/decode.go:969
				_go_fuzz_dep_.CoverTab[27243]++
										v.Set(reflect.ValueOf(string(s)))
//line /usr/local/go/src/encoding/json/decode.go:970
				// _ = "end of CoverTab[27243]"
			} else {
//line /usr/local/go/src/encoding/json/decode.go:971
				_go_fuzz_dep_.CoverTab[27244]++
										d.saveError(&UnmarshalTypeError{Value: "string", Type: v.Type(), Offset: int64(d.readIndex())})
//line /usr/local/go/src/encoding/json/decode.go:972
				// _ = "end of CoverTab[27244]"
			}
//line /usr/local/go/src/encoding/json/decode.go:973
			// _ = "end of CoverTab[27235]"
		}
//line /usr/local/go/src/encoding/json/decode.go:974
		// _ = "end of CoverTab[27205]"

	default:
//line /usr/local/go/src/encoding/json/decode.go:976
		_go_fuzz_dep_.CoverTab[27206]++
								if c != '-' && func() bool {
//line /usr/local/go/src/encoding/json/decode.go:977
			_go_fuzz_dep_.CoverTab[27245]++
//line /usr/local/go/src/encoding/json/decode.go:977
			return (c < '0' || func() bool {
//line /usr/local/go/src/encoding/json/decode.go:977
				_go_fuzz_dep_.CoverTab[27246]++
//line /usr/local/go/src/encoding/json/decode.go:977
				return c > '9'
//line /usr/local/go/src/encoding/json/decode.go:977
				// _ = "end of CoverTab[27246]"
//line /usr/local/go/src/encoding/json/decode.go:977
			}())
//line /usr/local/go/src/encoding/json/decode.go:977
			// _ = "end of CoverTab[27245]"
//line /usr/local/go/src/encoding/json/decode.go:977
		}() {
//line /usr/local/go/src/encoding/json/decode.go:977
			_go_fuzz_dep_.CoverTab[27247]++
									if fromQuoted {
//line /usr/local/go/src/encoding/json/decode.go:978
				_go_fuzz_dep_.CoverTab[27249]++
										return fmt.Errorf("json: invalid use of ,string struct tag, trying to unmarshal %q into %v", item, v.Type())
//line /usr/local/go/src/encoding/json/decode.go:979
				// _ = "end of CoverTab[27249]"
			} else {
//line /usr/local/go/src/encoding/json/decode.go:980
				_go_fuzz_dep_.CoverTab[27250]++
//line /usr/local/go/src/encoding/json/decode.go:980
				// _ = "end of CoverTab[27250]"
//line /usr/local/go/src/encoding/json/decode.go:980
			}
//line /usr/local/go/src/encoding/json/decode.go:980
			// _ = "end of CoverTab[27247]"
//line /usr/local/go/src/encoding/json/decode.go:980
			_go_fuzz_dep_.CoverTab[27248]++
									panic(phasePanicMsg)
//line /usr/local/go/src/encoding/json/decode.go:981
			// _ = "end of CoverTab[27248]"
		} else {
//line /usr/local/go/src/encoding/json/decode.go:982
			_go_fuzz_dep_.CoverTab[27251]++
//line /usr/local/go/src/encoding/json/decode.go:982
			// _ = "end of CoverTab[27251]"
//line /usr/local/go/src/encoding/json/decode.go:982
		}
//line /usr/local/go/src/encoding/json/decode.go:982
		// _ = "end of CoverTab[27206]"
//line /usr/local/go/src/encoding/json/decode.go:982
		_go_fuzz_dep_.CoverTab[27207]++
								s := string(item)
								switch v.Kind() {
		default:
//line /usr/local/go/src/encoding/json/decode.go:985
			_go_fuzz_dep_.CoverTab[27252]++
									if v.Kind() == reflect.String && func() bool {
//line /usr/local/go/src/encoding/json/decode.go:986
				_go_fuzz_dep_.CoverTab[27264]++
//line /usr/local/go/src/encoding/json/decode.go:986
				return v.Type() == numberType
//line /usr/local/go/src/encoding/json/decode.go:986
				// _ = "end of CoverTab[27264]"
//line /usr/local/go/src/encoding/json/decode.go:986
			}() {
//line /usr/local/go/src/encoding/json/decode.go:986
				_go_fuzz_dep_.CoverTab[27265]++

//line /usr/local/go/src/encoding/json/decode.go:989
				v.SetString(s)
										break
//line /usr/local/go/src/encoding/json/decode.go:990
				// _ = "end of CoverTab[27265]"
			} else {
//line /usr/local/go/src/encoding/json/decode.go:991
				_go_fuzz_dep_.CoverTab[27266]++
//line /usr/local/go/src/encoding/json/decode.go:991
				// _ = "end of CoverTab[27266]"
//line /usr/local/go/src/encoding/json/decode.go:991
			}
//line /usr/local/go/src/encoding/json/decode.go:991
			// _ = "end of CoverTab[27252]"
//line /usr/local/go/src/encoding/json/decode.go:991
			_go_fuzz_dep_.CoverTab[27253]++
									if fromQuoted {
//line /usr/local/go/src/encoding/json/decode.go:992
				_go_fuzz_dep_.CoverTab[27267]++
										return fmt.Errorf("json: invalid use of ,string struct tag, trying to unmarshal %q into %v", item, v.Type())
//line /usr/local/go/src/encoding/json/decode.go:993
				// _ = "end of CoverTab[27267]"
			} else {
//line /usr/local/go/src/encoding/json/decode.go:994
				_go_fuzz_dep_.CoverTab[27268]++
//line /usr/local/go/src/encoding/json/decode.go:994
				// _ = "end of CoverTab[27268]"
//line /usr/local/go/src/encoding/json/decode.go:994
			}
//line /usr/local/go/src/encoding/json/decode.go:994
			// _ = "end of CoverTab[27253]"
//line /usr/local/go/src/encoding/json/decode.go:994
			_go_fuzz_dep_.CoverTab[27254]++
									d.saveError(&UnmarshalTypeError{Value: "number", Type: v.Type(), Offset: int64(d.readIndex())})
//line /usr/local/go/src/encoding/json/decode.go:995
			// _ = "end of CoverTab[27254]"
		case reflect.Interface:
//line /usr/local/go/src/encoding/json/decode.go:996
			_go_fuzz_dep_.CoverTab[27255]++
									n, err := d.convertNumber(s)
									if err != nil {
//line /usr/local/go/src/encoding/json/decode.go:998
				_go_fuzz_dep_.CoverTab[27269]++
										d.saveError(err)
										break
//line /usr/local/go/src/encoding/json/decode.go:1000
				// _ = "end of CoverTab[27269]"
			} else {
//line /usr/local/go/src/encoding/json/decode.go:1001
				_go_fuzz_dep_.CoverTab[27270]++
//line /usr/local/go/src/encoding/json/decode.go:1001
				// _ = "end of CoverTab[27270]"
//line /usr/local/go/src/encoding/json/decode.go:1001
			}
//line /usr/local/go/src/encoding/json/decode.go:1001
			// _ = "end of CoverTab[27255]"
//line /usr/local/go/src/encoding/json/decode.go:1001
			_go_fuzz_dep_.CoverTab[27256]++
									if v.NumMethod() != 0 {
//line /usr/local/go/src/encoding/json/decode.go:1002
				_go_fuzz_dep_.CoverTab[27271]++
										d.saveError(&UnmarshalTypeError{Value: "number", Type: v.Type(), Offset: int64(d.readIndex())})
										break
//line /usr/local/go/src/encoding/json/decode.go:1004
				// _ = "end of CoverTab[27271]"
			} else {
//line /usr/local/go/src/encoding/json/decode.go:1005
				_go_fuzz_dep_.CoverTab[27272]++
//line /usr/local/go/src/encoding/json/decode.go:1005
				// _ = "end of CoverTab[27272]"
//line /usr/local/go/src/encoding/json/decode.go:1005
			}
//line /usr/local/go/src/encoding/json/decode.go:1005
			// _ = "end of CoverTab[27256]"
//line /usr/local/go/src/encoding/json/decode.go:1005
			_go_fuzz_dep_.CoverTab[27257]++
									v.Set(reflect.ValueOf(n))
//line /usr/local/go/src/encoding/json/decode.go:1006
			// _ = "end of CoverTab[27257]"

		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
//line /usr/local/go/src/encoding/json/decode.go:1008
			_go_fuzz_dep_.CoverTab[27258]++
									n, err := strconv.ParseInt(s, 10, 64)
									if err != nil || func() bool {
//line /usr/local/go/src/encoding/json/decode.go:1010
				_go_fuzz_dep_.CoverTab[27273]++
//line /usr/local/go/src/encoding/json/decode.go:1010
				return v.OverflowInt(n)
//line /usr/local/go/src/encoding/json/decode.go:1010
				// _ = "end of CoverTab[27273]"
//line /usr/local/go/src/encoding/json/decode.go:1010
			}() {
//line /usr/local/go/src/encoding/json/decode.go:1010
				_go_fuzz_dep_.CoverTab[27274]++
										d.saveError(&UnmarshalTypeError{Value: "number " + s, Type: v.Type(), Offset: int64(d.readIndex())})
										break
//line /usr/local/go/src/encoding/json/decode.go:1012
				// _ = "end of CoverTab[27274]"
			} else {
//line /usr/local/go/src/encoding/json/decode.go:1013
				_go_fuzz_dep_.CoverTab[27275]++
//line /usr/local/go/src/encoding/json/decode.go:1013
				// _ = "end of CoverTab[27275]"
//line /usr/local/go/src/encoding/json/decode.go:1013
			}
//line /usr/local/go/src/encoding/json/decode.go:1013
			// _ = "end of CoverTab[27258]"
//line /usr/local/go/src/encoding/json/decode.go:1013
			_go_fuzz_dep_.CoverTab[27259]++
									v.SetInt(n)
//line /usr/local/go/src/encoding/json/decode.go:1014
			// _ = "end of CoverTab[27259]"

		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
//line /usr/local/go/src/encoding/json/decode.go:1016
			_go_fuzz_dep_.CoverTab[27260]++
									n, err := strconv.ParseUint(s, 10, 64)
									if err != nil || func() bool {
//line /usr/local/go/src/encoding/json/decode.go:1018
				_go_fuzz_dep_.CoverTab[27276]++
//line /usr/local/go/src/encoding/json/decode.go:1018
				return v.OverflowUint(n)
//line /usr/local/go/src/encoding/json/decode.go:1018
				// _ = "end of CoverTab[27276]"
//line /usr/local/go/src/encoding/json/decode.go:1018
			}() {
//line /usr/local/go/src/encoding/json/decode.go:1018
				_go_fuzz_dep_.CoverTab[27277]++
										d.saveError(&UnmarshalTypeError{Value: "number " + s, Type: v.Type(), Offset: int64(d.readIndex())})
										break
//line /usr/local/go/src/encoding/json/decode.go:1020
				// _ = "end of CoverTab[27277]"
			} else {
//line /usr/local/go/src/encoding/json/decode.go:1021
				_go_fuzz_dep_.CoverTab[27278]++
//line /usr/local/go/src/encoding/json/decode.go:1021
				// _ = "end of CoverTab[27278]"
//line /usr/local/go/src/encoding/json/decode.go:1021
			}
//line /usr/local/go/src/encoding/json/decode.go:1021
			// _ = "end of CoverTab[27260]"
//line /usr/local/go/src/encoding/json/decode.go:1021
			_go_fuzz_dep_.CoverTab[27261]++
									v.SetUint(n)
//line /usr/local/go/src/encoding/json/decode.go:1022
			// _ = "end of CoverTab[27261]"

		case reflect.Float32, reflect.Float64:
//line /usr/local/go/src/encoding/json/decode.go:1024
			_go_fuzz_dep_.CoverTab[27262]++
									n, err := strconv.ParseFloat(s, v.Type().Bits())
									if err != nil || func() bool {
//line /usr/local/go/src/encoding/json/decode.go:1026
				_go_fuzz_dep_.CoverTab[27279]++
//line /usr/local/go/src/encoding/json/decode.go:1026
				return v.OverflowFloat(n)
//line /usr/local/go/src/encoding/json/decode.go:1026
				// _ = "end of CoverTab[27279]"
//line /usr/local/go/src/encoding/json/decode.go:1026
			}() {
//line /usr/local/go/src/encoding/json/decode.go:1026
				_go_fuzz_dep_.CoverTab[27280]++
										d.saveError(&UnmarshalTypeError{Value: "number " + s, Type: v.Type(), Offset: int64(d.readIndex())})
										break
//line /usr/local/go/src/encoding/json/decode.go:1028
				// _ = "end of CoverTab[27280]"
			} else {
//line /usr/local/go/src/encoding/json/decode.go:1029
				_go_fuzz_dep_.CoverTab[27281]++
//line /usr/local/go/src/encoding/json/decode.go:1029
				// _ = "end of CoverTab[27281]"
//line /usr/local/go/src/encoding/json/decode.go:1029
			}
//line /usr/local/go/src/encoding/json/decode.go:1029
			// _ = "end of CoverTab[27262]"
//line /usr/local/go/src/encoding/json/decode.go:1029
			_go_fuzz_dep_.CoverTab[27263]++
									v.SetFloat(n)
//line /usr/local/go/src/encoding/json/decode.go:1030
			// _ = "end of CoverTab[27263]"
		}
//line /usr/local/go/src/encoding/json/decode.go:1031
		// _ = "end of CoverTab[27207]"
	}
//line /usr/local/go/src/encoding/json/decode.go:1032
	// _ = "end of CoverTab[27176]"
//line /usr/local/go/src/encoding/json/decode.go:1032
	_go_fuzz_dep_.CoverTab[27177]++
							return nil
//line /usr/local/go/src/encoding/json/decode.go:1033
	// _ = "end of CoverTab[27177]"
}

//line /usr/local/go/src/encoding/json/decode.go:1040
// valueInterface is like value but returns interface{}
func (d *decodeState) valueInterface() (val any) {
//line /usr/local/go/src/encoding/json/decode.go:1041
	_go_fuzz_dep_.CoverTab[27282]++
							switch d.opcode {
	default:
//line /usr/local/go/src/encoding/json/decode.go:1043
		_go_fuzz_dep_.CoverTab[27284]++
								panic(phasePanicMsg)
//line /usr/local/go/src/encoding/json/decode.go:1044
		// _ = "end of CoverTab[27284]"
	case scanBeginArray:
//line /usr/local/go/src/encoding/json/decode.go:1045
		_go_fuzz_dep_.CoverTab[27285]++
								val = d.arrayInterface()
								d.scanNext()
//line /usr/local/go/src/encoding/json/decode.go:1047
		// _ = "end of CoverTab[27285]"
	case scanBeginObject:
//line /usr/local/go/src/encoding/json/decode.go:1048
		_go_fuzz_dep_.CoverTab[27286]++
								val = d.objectInterface()
								d.scanNext()
//line /usr/local/go/src/encoding/json/decode.go:1050
		// _ = "end of CoverTab[27286]"
	case scanBeginLiteral:
//line /usr/local/go/src/encoding/json/decode.go:1051
		_go_fuzz_dep_.CoverTab[27287]++
								val = d.literalInterface()
//line /usr/local/go/src/encoding/json/decode.go:1052
		// _ = "end of CoverTab[27287]"
	}
//line /usr/local/go/src/encoding/json/decode.go:1053
	// _ = "end of CoverTab[27282]"
//line /usr/local/go/src/encoding/json/decode.go:1053
	_go_fuzz_dep_.CoverTab[27283]++
							return
//line /usr/local/go/src/encoding/json/decode.go:1054
	// _ = "end of CoverTab[27283]"
}

// arrayInterface is like array but returns []interface{}.
func (d *decodeState) arrayInterface() []any {
//line /usr/local/go/src/encoding/json/decode.go:1058
	_go_fuzz_dep_.CoverTab[27288]++
							var v = make([]any, 0)
							for {
//line /usr/local/go/src/encoding/json/decode.go:1060
		_go_fuzz_dep_.CoverTab[27290]++

								d.scanWhile(scanSkipSpace)
								if d.opcode == scanEndArray {
//line /usr/local/go/src/encoding/json/decode.go:1063
			_go_fuzz_dep_.CoverTab[27294]++
									break
//line /usr/local/go/src/encoding/json/decode.go:1064
			// _ = "end of CoverTab[27294]"
		} else {
//line /usr/local/go/src/encoding/json/decode.go:1065
			_go_fuzz_dep_.CoverTab[27295]++
//line /usr/local/go/src/encoding/json/decode.go:1065
			// _ = "end of CoverTab[27295]"
//line /usr/local/go/src/encoding/json/decode.go:1065
		}
//line /usr/local/go/src/encoding/json/decode.go:1065
		// _ = "end of CoverTab[27290]"
//line /usr/local/go/src/encoding/json/decode.go:1065
		_go_fuzz_dep_.CoverTab[27291]++

								v = append(v, d.valueInterface())

//line /usr/local/go/src/encoding/json/decode.go:1070
		if d.opcode == scanSkipSpace {
//line /usr/local/go/src/encoding/json/decode.go:1070
			_go_fuzz_dep_.CoverTab[27296]++
									d.scanWhile(scanSkipSpace)
//line /usr/local/go/src/encoding/json/decode.go:1071
			// _ = "end of CoverTab[27296]"
		} else {
//line /usr/local/go/src/encoding/json/decode.go:1072
			_go_fuzz_dep_.CoverTab[27297]++
//line /usr/local/go/src/encoding/json/decode.go:1072
			// _ = "end of CoverTab[27297]"
//line /usr/local/go/src/encoding/json/decode.go:1072
		}
//line /usr/local/go/src/encoding/json/decode.go:1072
		// _ = "end of CoverTab[27291]"
//line /usr/local/go/src/encoding/json/decode.go:1072
		_go_fuzz_dep_.CoverTab[27292]++
								if d.opcode == scanEndArray {
//line /usr/local/go/src/encoding/json/decode.go:1073
			_go_fuzz_dep_.CoverTab[27298]++
									break
//line /usr/local/go/src/encoding/json/decode.go:1074
			// _ = "end of CoverTab[27298]"
		} else {
//line /usr/local/go/src/encoding/json/decode.go:1075
			_go_fuzz_dep_.CoverTab[27299]++
//line /usr/local/go/src/encoding/json/decode.go:1075
			// _ = "end of CoverTab[27299]"
//line /usr/local/go/src/encoding/json/decode.go:1075
		}
//line /usr/local/go/src/encoding/json/decode.go:1075
		// _ = "end of CoverTab[27292]"
//line /usr/local/go/src/encoding/json/decode.go:1075
		_go_fuzz_dep_.CoverTab[27293]++
								if d.opcode != scanArrayValue {
//line /usr/local/go/src/encoding/json/decode.go:1076
			_go_fuzz_dep_.CoverTab[27300]++
									panic(phasePanicMsg)
//line /usr/local/go/src/encoding/json/decode.go:1077
			// _ = "end of CoverTab[27300]"
		} else {
//line /usr/local/go/src/encoding/json/decode.go:1078
			_go_fuzz_dep_.CoverTab[27301]++
//line /usr/local/go/src/encoding/json/decode.go:1078
			// _ = "end of CoverTab[27301]"
//line /usr/local/go/src/encoding/json/decode.go:1078
		}
//line /usr/local/go/src/encoding/json/decode.go:1078
		// _ = "end of CoverTab[27293]"
	}
//line /usr/local/go/src/encoding/json/decode.go:1079
	// _ = "end of CoverTab[27288]"
//line /usr/local/go/src/encoding/json/decode.go:1079
	_go_fuzz_dep_.CoverTab[27289]++
							return v
//line /usr/local/go/src/encoding/json/decode.go:1080
	// _ = "end of CoverTab[27289]"
}

// objectInterface is like object but returns map[string]interface{}.
func (d *decodeState) objectInterface() map[string]any {
//line /usr/local/go/src/encoding/json/decode.go:1084
	_go_fuzz_dep_.CoverTab[27302]++
							m := make(map[string]any)
							for {
//line /usr/local/go/src/encoding/json/decode.go:1086
		_go_fuzz_dep_.CoverTab[27304]++

								d.scanWhile(scanSkipSpace)
								if d.opcode == scanEndObject {
//line /usr/local/go/src/encoding/json/decode.go:1089
			_go_fuzz_dep_.CoverTab[27312]++

									break
//line /usr/local/go/src/encoding/json/decode.go:1091
			// _ = "end of CoverTab[27312]"
		} else {
//line /usr/local/go/src/encoding/json/decode.go:1092
			_go_fuzz_dep_.CoverTab[27313]++
//line /usr/local/go/src/encoding/json/decode.go:1092
			// _ = "end of CoverTab[27313]"
//line /usr/local/go/src/encoding/json/decode.go:1092
		}
//line /usr/local/go/src/encoding/json/decode.go:1092
		// _ = "end of CoverTab[27304]"
//line /usr/local/go/src/encoding/json/decode.go:1092
		_go_fuzz_dep_.CoverTab[27305]++
								if d.opcode != scanBeginLiteral {
//line /usr/local/go/src/encoding/json/decode.go:1093
			_go_fuzz_dep_.CoverTab[27314]++
									panic(phasePanicMsg)
//line /usr/local/go/src/encoding/json/decode.go:1094
			// _ = "end of CoverTab[27314]"
		} else {
//line /usr/local/go/src/encoding/json/decode.go:1095
			_go_fuzz_dep_.CoverTab[27315]++
//line /usr/local/go/src/encoding/json/decode.go:1095
			// _ = "end of CoverTab[27315]"
//line /usr/local/go/src/encoding/json/decode.go:1095
		}
//line /usr/local/go/src/encoding/json/decode.go:1095
		// _ = "end of CoverTab[27305]"
//line /usr/local/go/src/encoding/json/decode.go:1095
		_go_fuzz_dep_.CoverTab[27306]++

//line /usr/local/go/src/encoding/json/decode.go:1098
		start := d.readIndex()
		d.rescanLiteral()
		item := d.data[start:d.readIndex()]
		key, ok := unquote(item)
		if !ok {
//line /usr/local/go/src/encoding/json/decode.go:1102
			_go_fuzz_dep_.CoverTab[27316]++
									panic(phasePanicMsg)
//line /usr/local/go/src/encoding/json/decode.go:1103
			// _ = "end of CoverTab[27316]"
		} else {
//line /usr/local/go/src/encoding/json/decode.go:1104
			_go_fuzz_dep_.CoverTab[27317]++
//line /usr/local/go/src/encoding/json/decode.go:1104
			// _ = "end of CoverTab[27317]"
//line /usr/local/go/src/encoding/json/decode.go:1104
		}
//line /usr/local/go/src/encoding/json/decode.go:1104
		// _ = "end of CoverTab[27306]"
//line /usr/local/go/src/encoding/json/decode.go:1104
		_go_fuzz_dep_.CoverTab[27307]++

//line /usr/local/go/src/encoding/json/decode.go:1107
		if d.opcode == scanSkipSpace {
//line /usr/local/go/src/encoding/json/decode.go:1107
			_go_fuzz_dep_.CoverTab[27318]++
									d.scanWhile(scanSkipSpace)
//line /usr/local/go/src/encoding/json/decode.go:1108
			// _ = "end of CoverTab[27318]"
		} else {
//line /usr/local/go/src/encoding/json/decode.go:1109
			_go_fuzz_dep_.CoverTab[27319]++
//line /usr/local/go/src/encoding/json/decode.go:1109
			// _ = "end of CoverTab[27319]"
//line /usr/local/go/src/encoding/json/decode.go:1109
		}
//line /usr/local/go/src/encoding/json/decode.go:1109
		// _ = "end of CoverTab[27307]"
//line /usr/local/go/src/encoding/json/decode.go:1109
		_go_fuzz_dep_.CoverTab[27308]++
								if d.opcode != scanObjectKey {
//line /usr/local/go/src/encoding/json/decode.go:1110
			_go_fuzz_dep_.CoverTab[27320]++
									panic(phasePanicMsg)
//line /usr/local/go/src/encoding/json/decode.go:1111
			// _ = "end of CoverTab[27320]"
		} else {
//line /usr/local/go/src/encoding/json/decode.go:1112
			_go_fuzz_dep_.CoverTab[27321]++
//line /usr/local/go/src/encoding/json/decode.go:1112
			// _ = "end of CoverTab[27321]"
//line /usr/local/go/src/encoding/json/decode.go:1112
		}
//line /usr/local/go/src/encoding/json/decode.go:1112
		// _ = "end of CoverTab[27308]"
//line /usr/local/go/src/encoding/json/decode.go:1112
		_go_fuzz_dep_.CoverTab[27309]++
								d.scanWhile(scanSkipSpace)

//line /usr/local/go/src/encoding/json/decode.go:1116
		m[key] = d.valueInterface()

//line /usr/local/go/src/encoding/json/decode.go:1119
		if d.opcode == scanSkipSpace {
//line /usr/local/go/src/encoding/json/decode.go:1119
			_go_fuzz_dep_.CoverTab[27322]++
									d.scanWhile(scanSkipSpace)
//line /usr/local/go/src/encoding/json/decode.go:1120
			// _ = "end of CoverTab[27322]"
		} else {
//line /usr/local/go/src/encoding/json/decode.go:1121
			_go_fuzz_dep_.CoverTab[27323]++
//line /usr/local/go/src/encoding/json/decode.go:1121
			// _ = "end of CoverTab[27323]"
//line /usr/local/go/src/encoding/json/decode.go:1121
		}
//line /usr/local/go/src/encoding/json/decode.go:1121
		// _ = "end of CoverTab[27309]"
//line /usr/local/go/src/encoding/json/decode.go:1121
		_go_fuzz_dep_.CoverTab[27310]++
								if d.opcode == scanEndObject {
//line /usr/local/go/src/encoding/json/decode.go:1122
			_go_fuzz_dep_.CoverTab[27324]++
									break
//line /usr/local/go/src/encoding/json/decode.go:1123
			// _ = "end of CoverTab[27324]"
		} else {
//line /usr/local/go/src/encoding/json/decode.go:1124
			_go_fuzz_dep_.CoverTab[27325]++
//line /usr/local/go/src/encoding/json/decode.go:1124
			// _ = "end of CoverTab[27325]"
//line /usr/local/go/src/encoding/json/decode.go:1124
		}
//line /usr/local/go/src/encoding/json/decode.go:1124
		// _ = "end of CoverTab[27310]"
//line /usr/local/go/src/encoding/json/decode.go:1124
		_go_fuzz_dep_.CoverTab[27311]++
								if d.opcode != scanObjectValue {
//line /usr/local/go/src/encoding/json/decode.go:1125
			_go_fuzz_dep_.CoverTab[27326]++
									panic(phasePanicMsg)
//line /usr/local/go/src/encoding/json/decode.go:1126
			// _ = "end of CoverTab[27326]"
		} else {
//line /usr/local/go/src/encoding/json/decode.go:1127
			_go_fuzz_dep_.CoverTab[27327]++
//line /usr/local/go/src/encoding/json/decode.go:1127
			// _ = "end of CoverTab[27327]"
//line /usr/local/go/src/encoding/json/decode.go:1127
		}
//line /usr/local/go/src/encoding/json/decode.go:1127
		// _ = "end of CoverTab[27311]"
	}
//line /usr/local/go/src/encoding/json/decode.go:1128
	// _ = "end of CoverTab[27302]"
//line /usr/local/go/src/encoding/json/decode.go:1128
	_go_fuzz_dep_.CoverTab[27303]++
							return m
//line /usr/local/go/src/encoding/json/decode.go:1129
	// _ = "end of CoverTab[27303]"
}

// literalInterface consumes and returns a literal from d.data[d.off-1:] and
//line /usr/local/go/src/encoding/json/decode.go:1132
// it reads the following byte ahead. The first byte of the literal has been
//line /usr/local/go/src/encoding/json/decode.go:1132
// read already (that's how the caller knows it's a literal).
//line /usr/local/go/src/encoding/json/decode.go:1135
func (d *decodeState) literalInterface() any {
//line /usr/local/go/src/encoding/json/decode.go:1135
	_go_fuzz_dep_.CoverTab[27328]++

							start := d.readIndex()
							d.rescanLiteral()

							item := d.data[start:d.readIndex()]

							switch c := item[0]; c {
	case 'n':
//line /usr/local/go/src/encoding/json/decode.go:1143
		_go_fuzz_dep_.CoverTab[27329]++
								return nil
//line /usr/local/go/src/encoding/json/decode.go:1144
		// _ = "end of CoverTab[27329]"

	case 't', 'f':
//line /usr/local/go/src/encoding/json/decode.go:1146
		_go_fuzz_dep_.CoverTab[27330]++
								return c == 't'
//line /usr/local/go/src/encoding/json/decode.go:1147
		// _ = "end of CoverTab[27330]"

	case '"':
//line /usr/local/go/src/encoding/json/decode.go:1149
		_go_fuzz_dep_.CoverTab[27331]++
								s, ok := unquote(item)
								if !ok {
//line /usr/local/go/src/encoding/json/decode.go:1151
			_go_fuzz_dep_.CoverTab[27336]++
									panic(phasePanicMsg)
//line /usr/local/go/src/encoding/json/decode.go:1152
			// _ = "end of CoverTab[27336]"
		} else {
//line /usr/local/go/src/encoding/json/decode.go:1153
			_go_fuzz_dep_.CoverTab[27337]++
//line /usr/local/go/src/encoding/json/decode.go:1153
			// _ = "end of CoverTab[27337]"
//line /usr/local/go/src/encoding/json/decode.go:1153
		}
//line /usr/local/go/src/encoding/json/decode.go:1153
		// _ = "end of CoverTab[27331]"
//line /usr/local/go/src/encoding/json/decode.go:1153
		_go_fuzz_dep_.CoverTab[27332]++
								return s
//line /usr/local/go/src/encoding/json/decode.go:1154
		// _ = "end of CoverTab[27332]"

	default:
//line /usr/local/go/src/encoding/json/decode.go:1156
		_go_fuzz_dep_.CoverTab[27333]++
								if c != '-' && func() bool {
//line /usr/local/go/src/encoding/json/decode.go:1157
			_go_fuzz_dep_.CoverTab[27338]++
//line /usr/local/go/src/encoding/json/decode.go:1157
			return (c < '0' || func() bool {
//line /usr/local/go/src/encoding/json/decode.go:1157
				_go_fuzz_dep_.CoverTab[27339]++
//line /usr/local/go/src/encoding/json/decode.go:1157
				return c > '9'
//line /usr/local/go/src/encoding/json/decode.go:1157
				// _ = "end of CoverTab[27339]"
//line /usr/local/go/src/encoding/json/decode.go:1157
			}())
//line /usr/local/go/src/encoding/json/decode.go:1157
			// _ = "end of CoverTab[27338]"
//line /usr/local/go/src/encoding/json/decode.go:1157
		}() {
//line /usr/local/go/src/encoding/json/decode.go:1157
			_go_fuzz_dep_.CoverTab[27340]++
									panic(phasePanicMsg)
//line /usr/local/go/src/encoding/json/decode.go:1158
			// _ = "end of CoverTab[27340]"
		} else {
//line /usr/local/go/src/encoding/json/decode.go:1159
			_go_fuzz_dep_.CoverTab[27341]++
//line /usr/local/go/src/encoding/json/decode.go:1159
			// _ = "end of CoverTab[27341]"
//line /usr/local/go/src/encoding/json/decode.go:1159
		}
//line /usr/local/go/src/encoding/json/decode.go:1159
		// _ = "end of CoverTab[27333]"
//line /usr/local/go/src/encoding/json/decode.go:1159
		_go_fuzz_dep_.CoverTab[27334]++
								n, err := d.convertNumber(string(item))
								if err != nil {
//line /usr/local/go/src/encoding/json/decode.go:1161
			_go_fuzz_dep_.CoverTab[27342]++
									d.saveError(err)
//line /usr/local/go/src/encoding/json/decode.go:1162
			// _ = "end of CoverTab[27342]"
		} else {
//line /usr/local/go/src/encoding/json/decode.go:1163
			_go_fuzz_dep_.CoverTab[27343]++
//line /usr/local/go/src/encoding/json/decode.go:1163
			// _ = "end of CoverTab[27343]"
//line /usr/local/go/src/encoding/json/decode.go:1163
		}
//line /usr/local/go/src/encoding/json/decode.go:1163
		// _ = "end of CoverTab[27334]"
//line /usr/local/go/src/encoding/json/decode.go:1163
		_go_fuzz_dep_.CoverTab[27335]++
								return n
//line /usr/local/go/src/encoding/json/decode.go:1164
		// _ = "end of CoverTab[27335]"
	}
//line /usr/local/go/src/encoding/json/decode.go:1165
	// _ = "end of CoverTab[27328]"
}

// getu4 decodes \uXXXX from the beginning of s, returning the hex value,
//line /usr/local/go/src/encoding/json/decode.go:1168
// or it returns -1.
//line /usr/local/go/src/encoding/json/decode.go:1170
func getu4(s []byte) rune {
//line /usr/local/go/src/encoding/json/decode.go:1170
	_go_fuzz_dep_.CoverTab[27344]++
							if len(s) < 6 || func() bool {
//line /usr/local/go/src/encoding/json/decode.go:1171
		_go_fuzz_dep_.CoverTab[27347]++
//line /usr/local/go/src/encoding/json/decode.go:1171
		return s[0] != '\\'
//line /usr/local/go/src/encoding/json/decode.go:1171
		// _ = "end of CoverTab[27347]"
//line /usr/local/go/src/encoding/json/decode.go:1171
	}() || func() bool {
//line /usr/local/go/src/encoding/json/decode.go:1171
		_go_fuzz_dep_.CoverTab[27348]++
//line /usr/local/go/src/encoding/json/decode.go:1171
		return s[1] != 'u'
//line /usr/local/go/src/encoding/json/decode.go:1171
		// _ = "end of CoverTab[27348]"
//line /usr/local/go/src/encoding/json/decode.go:1171
	}() {
//line /usr/local/go/src/encoding/json/decode.go:1171
		_go_fuzz_dep_.CoverTab[27349]++
								return -1
//line /usr/local/go/src/encoding/json/decode.go:1172
		// _ = "end of CoverTab[27349]"
	} else {
//line /usr/local/go/src/encoding/json/decode.go:1173
		_go_fuzz_dep_.CoverTab[27350]++
//line /usr/local/go/src/encoding/json/decode.go:1173
		// _ = "end of CoverTab[27350]"
//line /usr/local/go/src/encoding/json/decode.go:1173
	}
//line /usr/local/go/src/encoding/json/decode.go:1173
	// _ = "end of CoverTab[27344]"
//line /usr/local/go/src/encoding/json/decode.go:1173
	_go_fuzz_dep_.CoverTab[27345]++
							var r rune
							for _, c := range s[2:6] {
//line /usr/local/go/src/encoding/json/decode.go:1175
		_go_fuzz_dep_.CoverTab[27351]++
								switch {
		case '0' <= c && func() bool {
//line /usr/local/go/src/encoding/json/decode.go:1177
			_go_fuzz_dep_.CoverTab[27357]++
//line /usr/local/go/src/encoding/json/decode.go:1177
			return c <= '9'
//line /usr/local/go/src/encoding/json/decode.go:1177
			// _ = "end of CoverTab[27357]"
//line /usr/local/go/src/encoding/json/decode.go:1177
		}():
//line /usr/local/go/src/encoding/json/decode.go:1177
			_go_fuzz_dep_.CoverTab[27353]++
									c = c - '0'
//line /usr/local/go/src/encoding/json/decode.go:1178
			// _ = "end of CoverTab[27353]"
		case 'a' <= c && func() bool {
//line /usr/local/go/src/encoding/json/decode.go:1179
			_go_fuzz_dep_.CoverTab[27358]++
//line /usr/local/go/src/encoding/json/decode.go:1179
			return c <= 'f'
//line /usr/local/go/src/encoding/json/decode.go:1179
			// _ = "end of CoverTab[27358]"
//line /usr/local/go/src/encoding/json/decode.go:1179
		}():
//line /usr/local/go/src/encoding/json/decode.go:1179
			_go_fuzz_dep_.CoverTab[27354]++
									c = c - 'a' + 10
//line /usr/local/go/src/encoding/json/decode.go:1180
			// _ = "end of CoverTab[27354]"
		case 'A' <= c && func() bool {
//line /usr/local/go/src/encoding/json/decode.go:1181
			_go_fuzz_dep_.CoverTab[27359]++
//line /usr/local/go/src/encoding/json/decode.go:1181
			return c <= 'F'
//line /usr/local/go/src/encoding/json/decode.go:1181
			// _ = "end of CoverTab[27359]"
//line /usr/local/go/src/encoding/json/decode.go:1181
		}():
//line /usr/local/go/src/encoding/json/decode.go:1181
			_go_fuzz_dep_.CoverTab[27355]++
									c = c - 'A' + 10
//line /usr/local/go/src/encoding/json/decode.go:1182
			// _ = "end of CoverTab[27355]"
		default:
//line /usr/local/go/src/encoding/json/decode.go:1183
			_go_fuzz_dep_.CoverTab[27356]++
									return -1
//line /usr/local/go/src/encoding/json/decode.go:1184
			// _ = "end of CoverTab[27356]"
		}
//line /usr/local/go/src/encoding/json/decode.go:1185
		// _ = "end of CoverTab[27351]"
//line /usr/local/go/src/encoding/json/decode.go:1185
		_go_fuzz_dep_.CoverTab[27352]++
								r = r*16 + rune(c)
//line /usr/local/go/src/encoding/json/decode.go:1186
		// _ = "end of CoverTab[27352]"
	}
//line /usr/local/go/src/encoding/json/decode.go:1187
	// _ = "end of CoverTab[27345]"
//line /usr/local/go/src/encoding/json/decode.go:1187
	_go_fuzz_dep_.CoverTab[27346]++
							return r
//line /usr/local/go/src/encoding/json/decode.go:1188
	// _ = "end of CoverTab[27346]"
}

// unquote converts a quoted JSON string literal s into an actual string t.
//line /usr/local/go/src/encoding/json/decode.go:1191
// The rules are different than for Go, so cannot use strconv.Unquote.
//line /usr/local/go/src/encoding/json/decode.go:1193
func unquote(s []byte) (t string, ok bool) {
//line /usr/local/go/src/encoding/json/decode.go:1193
	_go_fuzz_dep_.CoverTab[27360]++
							s, ok = unquoteBytes(s)
							t = string(s)
							return
//line /usr/local/go/src/encoding/json/decode.go:1196
	// _ = "end of CoverTab[27360]"
}

func unquoteBytes(s []byte) (t []byte, ok bool) {
//line /usr/local/go/src/encoding/json/decode.go:1199
	_go_fuzz_dep_.CoverTab[27361]++
							if len(s) < 2 || func() bool {
//line /usr/local/go/src/encoding/json/decode.go:1200
		_go_fuzz_dep_.CoverTab[27366]++
//line /usr/local/go/src/encoding/json/decode.go:1200
		return s[0] != '"'
//line /usr/local/go/src/encoding/json/decode.go:1200
		// _ = "end of CoverTab[27366]"
//line /usr/local/go/src/encoding/json/decode.go:1200
	}() || func() bool {
//line /usr/local/go/src/encoding/json/decode.go:1200
		_go_fuzz_dep_.CoverTab[27367]++
//line /usr/local/go/src/encoding/json/decode.go:1200
		return s[len(s)-1] != '"'
//line /usr/local/go/src/encoding/json/decode.go:1200
		// _ = "end of CoverTab[27367]"
//line /usr/local/go/src/encoding/json/decode.go:1200
	}() {
//line /usr/local/go/src/encoding/json/decode.go:1200
		_go_fuzz_dep_.CoverTab[27368]++
								return
//line /usr/local/go/src/encoding/json/decode.go:1201
		// _ = "end of CoverTab[27368]"
	} else {
//line /usr/local/go/src/encoding/json/decode.go:1202
		_go_fuzz_dep_.CoverTab[27369]++
//line /usr/local/go/src/encoding/json/decode.go:1202
		// _ = "end of CoverTab[27369]"
//line /usr/local/go/src/encoding/json/decode.go:1202
	}
//line /usr/local/go/src/encoding/json/decode.go:1202
	// _ = "end of CoverTab[27361]"
//line /usr/local/go/src/encoding/json/decode.go:1202
	_go_fuzz_dep_.CoverTab[27362]++
							s = s[1 : len(s)-1]

//line /usr/local/go/src/encoding/json/decode.go:1208
	r := 0
	for r < len(s) {
//line /usr/local/go/src/encoding/json/decode.go:1209
		_go_fuzz_dep_.CoverTab[27370]++
								c := s[r]
								if c == '\\' || func() bool {
//line /usr/local/go/src/encoding/json/decode.go:1211
			_go_fuzz_dep_.CoverTab[27374]++
//line /usr/local/go/src/encoding/json/decode.go:1211
			return c == '"'
//line /usr/local/go/src/encoding/json/decode.go:1211
			// _ = "end of CoverTab[27374]"
//line /usr/local/go/src/encoding/json/decode.go:1211
		}() || func() bool {
//line /usr/local/go/src/encoding/json/decode.go:1211
			_go_fuzz_dep_.CoverTab[27375]++
//line /usr/local/go/src/encoding/json/decode.go:1211
			return c < ' '
//line /usr/local/go/src/encoding/json/decode.go:1211
			// _ = "end of CoverTab[27375]"
//line /usr/local/go/src/encoding/json/decode.go:1211
		}() {
//line /usr/local/go/src/encoding/json/decode.go:1211
			_go_fuzz_dep_.CoverTab[27376]++
									break
//line /usr/local/go/src/encoding/json/decode.go:1212
			// _ = "end of CoverTab[27376]"
		} else {
//line /usr/local/go/src/encoding/json/decode.go:1213
			_go_fuzz_dep_.CoverTab[27377]++
//line /usr/local/go/src/encoding/json/decode.go:1213
			// _ = "end of CoverTab[27377]"
//line /usr/local/go/src/encoding/json/decode.go:1213
		}
//line /usr/local/go/src/encoding/json/decode.go:1213
		// _ = "end of CoverTab[27370]"
//line /usr/local/go/src/encoding/json/decode.go:1213
		_go_fuzz_dep_.CoverTab[27371]++
								if c < utf8.RuneSelf {
//line /usr/local/go/src/encoding/json/decode.go:1214
			_go_fuzz_dep_.CoverTab[27378]++
									r++
									continue
//line /usr/local/go/src/encoding/json/decode.go:1216
			// _ = "end of CoverTab[27378]"
		} else {
//line /usr/local/go/src/encoding/json/decode.go:1217
			_go_fuzz_dep_.CoverTab[27379]++
//line /usr/local/go/src/encoding/json/decode.go:1217
			// _ = "end of CoverTab[27379]"
//line /usr/local/go/src/encoding/json/decode.go:1217
		}
//line /usr/local/go/src/encoding/json/decode.go:1217
		// _ = "end of CoverTab[27371]"
//line /usr/local/go/src/encoding/json/decode.go:1217
		_go_fuzz_dep_.CoverTab[27372]++
								rr, size := utf8.DecodeRune(s[r:])
								if rr == utf8.RuneError && func() bool {
//line /usr/local/go/src/encoding/json/decode.go:1219
			_go_fuzz_dep_.CoverTab[27380]++
//line /usr/local/go/src/encoding/json/decode.go:1219
			return size == 1
//line /usr/local/go/src/encoding/json/decode.go:1219
			// _ = "end of CoverTab[27380]"
//line /usr/local/go/src/encoding/json/decode.go:1219
		}() {
//line /usr/local/go/src/encoding/json/decode.go:1219
			_go_fuzz_dep_.CoverTab[27381]++
									break
//line /usr/local/go/src/encoding/json/decode.go:1220
			// _ = "end of CoverTab[27381]"
		} else {
//line /usr/local/go/src/encoding/json/decode.go:1221
			_go_fuzz_dep_.CoverTab[27382]++
//line /usr/local/go/src/encoding/json/decode.go:1221
			// _ = "end of CoverTab[27382]"
//line /usr/local/go/src/encoding/json/decode.go:1221
		}
//line /usr/local/go/src/encoding/json/decode.go:1221
		// _ = "end of CoverTab[27372]"
//line /usr/local/go/src/encoding/json/decode.go:1221
		_go_fuzz_dep_.CoverTab[27373]++
								r += size
//line /usr/local/go/src/encoding/json/decode.go:1222
		// _ = "end of CoverTab[27373]"
	}
//line /usr/local/go/src/encoding/json/decode.go:1223
	// _ = "end of CoverTab[27362]"
//line /usr/local/go/src/encoding/json/decode.go:1223
	_go_fuzz_dep_.CoverTab[27363]++
							if r == len(s) {
//line /usr/local/go/src/encoding/json/decode.go:1224
		_go_fuzz_dep_.CoverTab[27383]++
								return s, true
//line /usr/local/go/src/encoding/json/decode.go:1225
		// _ = "end of CoverTab[27383]"
	} else {
//line /usr/local/go/src/encoding/json/decode.go:1226
		_go_fuzz_dep_.CoverTab[27384]++
//line /usr/local/go/src/encoding/json/decode.go:1226
		// _ = "end of CoverTab[27384]"
//line /usr/local/go/src/encoding/json/decode.go:1226
	}
//line /usr/local/go/src/encoding/json/decode.go:1226
	// _ = "end of CoverTab[27363]"
//line /usr/local/go/src/encoding/json/decode.go:1226
	_go_fuzz_dep_.CoverTab[27364]++

							b := make([]byte, len(s)+2*utf8.UTFMax)
							w := copy(b, s[0:r])
							for r < len(s) {
//line /usr/local/go/src/encoding/json/decode.go:1230
		_go_fuzz_dep_.CoverTab[27385]++

//line /usr/local/go/src/encoding/json/decode.go:1234
		if w >= len(b)-2*utf8.UTFMax {
//line /usr/local/go/src/encoding/json/decode.go:1234
			_go_fuzz_dep_.CoverTab[27387]++
									nb := make([]byte, (len(b)+utf8.UTFMax)*2)
									copy(nb, b[0:w])
									b = nb
//line /usr/local/go/src/encoding/json/decode.go:1237
			// _ = "end of CoverTab[27387]"
		} else {
//line /usr/local/go/src/encoding/json/decode.go:1238
			_go_fuzz_dep_.CoverTab[27388]++
//line /usr/local/go/src/encoding/json/decode.go:1238
			// _ = "end of CoverTab[27388]"
//line /usr/local/go/src/encoding/json/decode.go:1238
		}
//line /usr/local/go/src/encoding/json/decode.go:1238
		// _ = "end of CoverTab[27385]"
//line /usr/local/go/src/encoding/json/decode.go:1238
		_go_fuzz_dep_.CoverTab[27386]++
								switch c := s[r]; {
		case c == '\\':
//line /usr/local/go/src/encoding/json/decode.go:1240
			_go_fuzz_dep_.CoverTab[27389]++
									r++
									if r >= len(s) {
//line /usr/local/go/src/encoding/json/decode.go:1242
				_go_fuzz_dep_.CoverTab[27394]++
										return
//line /usr/local/go/src/encoding/json/decode.go:1243
				// _ = "end of CoverTab[27394]"
			} else {
//line /usr/local/go/src/encoding/json/decode.go:1244
				_go_fuzz_dep_.CoverTab[27395]++
//line /usr/local/go/src/encoding/json/decode.go:1244
				// _ = "end of CoverTab[27395]"
//line /usr/local/go/src/encoding/json/decode.go:1244
			}
//line /usr/local/go/src/encoding/json/decode.go:1244
			// _ = "end of CoverTab[27389]"
//line /usr/local/go/src/encoding/json/decode.go:1244
			_go_fuzz_dep_.CoverTab[27390]++
									switch s[r] {
			default:
//line /usr/local/go/src/encoding/json/decode.go:1246
				_go_fuzz_dep_.CoverTab[27396]++
										return
//line /usr/local/go/src/encoding/json/decode.go:1247
				// _ = "end of CoverTab[27396]"
			case '"', '\\', '/', '\'':
//line /usr/local/go/src/encoding/json/decode.go:1248
				_go_fuzz_dep_.CoverTab[27397]++
										b[w] = s[r]
										r++
										w++
//line /usr/local/go/src/encoding/json/decode.go:1251
				// _ = "end of CoverTab[27397]"
			case 'b':
//line /usr/local/go/src/encoding/json/decode.go:1252
				_go_fuzz_dep_.CoverTab[27398]++
										b[w] = '\b'
										r++
										w++
//line /usr/local/go/src/encoding/json/decode.go:1255
				// _ = "end of CoverTab[27398]"
			case 'f':
//line /usr/local/go/src/encoding/json/decode.go:1256
				_go_fuzz_dep_.CoverTab[27399]++
										b[w] = '\f'
										r++
										w++
//line /usr/local/go/src/encoding/json/decode.go:1259
				// _ = "end of CoverTab[27399]"
			case 'n':
//line /usr/local/go/src/encoding/json/decode.go:1260
				_go_fuzz_dep_.CoverTab[27400]++
										b[w] = '\n'
										r++
										w++
//line /usr/local/go/src/encoding/json/decode.go:1263
				// _ = "end of CoverTab[27400]"
			case 'r':
//line /usr/local/go/src/encoding/json/decode.go:1264
				_go_fuzz_dep_.CoverTab[27401]++
										b[w] = '\r'
										r++
										w++
//line /usr/local/go/src/encoding/json/decode.go:1267
				// _ = "end of CoverTab[27401]"
			case 't':
//line /usr/local/go/src/encoding/json/decode.go:1268
				_go_fuzz_dep_.CoverTab[27402]++
										b[w] = '\t'
										r++
										w++
//line /usr/local/go/src/encoding/json/decode.go:1271
				// _ = "end of CoverTab[27402]"
			case 'u':
//line /usr/local/go/src/encoding/json/decode.go:1272
				_go_fuzz_dep_.CoverTab[27403]++
										r--
										rr := getu4(s[r:])
										if rr < 0 {
//line /usr/local/go/src/encoding/json/decode.go:1275
					_go_fuzz_dep_.CoverTab[27406]++
											return
//line /usr/local/go/src/encoding/json/decode.go:1276
					// _ = "end of CoverTab[27406]"
				} else {
//line /usr/local/go/src/encoding/json/decode.go:1277
					_go_fuzz_dep_.CoverTab[27407]++
//line /usr/local/go/src/encoding/json/decode.go:1277
					// _ = "end of CoverTab[27407]"
//line /usr/local/go/src/encoding/json/decode.go:1277
				}
//line /usr/local/go/src/encoding/json/decode.go:1277
				// _ = "end of CoverTab[27403]"
//line /usr/local/go/src/encoding/json/decode.go:1277
				_go_fuzz_dep_.CoverTab[27404]++
										r += 6
										if utf16.IsSurrogate(rr) {
//line /usr/local/go/src/encoding/json/decode.go:1279
					_go_fuzz_dep_.CoverTab[27408]++
											rr1 := getu4(s[r:])
											if dec := utf16.DecodeRune(rr, rr1); dec != unicode.ReplacementChar {
//line /usr/local/go/src/encoding/json/decode.go:1281
						_go_fuzz_dep_.CoverTab[27410]++

												r += 6
												w += utf8.EncodeRune(b[w:], dec)
												break
//line /usr/local/go/src/encoding/json/decode.go:1285
						// _ = "end of CoverTab[27410]"
					} else {
//line /usr/local/go/src/encoding/json/decode.go:1286
						_go_fuzz_dep_.CoverTab[27411]++
//line /usr/local/go/src/encoding/json/decode.go:1286
						// _ = "end of CoverTab[27411]"
//line /usr/local/go/src/encoding/json/decode.go:1286
					}
//line /usr/local/go/src/encoding/json/decode.go:1286
					// _ = "end of CoverTab[27408]"
//line /usr/local/go/src/encoding/json/decode.go:1286
					_go_fuzz_dep_.CoverTab[27409]++

											rr = unicode.ReplacementChar
//line /usr/local/go/src/encoding/json/decode.go:1288
					// _ = "end of CoverTab[27409]"
				} else {
//line /usr/local/go/src/encoding/json/decode.go:1289
					_go_fuzz_dep_.CoverTab[27412]++
//line /usr/local/go/src/encoding/json/decode.go:1289
					// _ = "end of CoverTab[27412]"
//line /usr/local/go/src/encoding/json/decode.go:1289
				}
//line /usr/local/go/src/encoding/json/decode.go:1289
				// _ = "end of CoverTab[27404]"
//line /usr/local/go/src/encoding/json/decode.go:1289
				_go_fuzz_dep_.CoverTab[27405]++
										w += utf8.EncodeRune(b[w:], rr)
//line /usr/local/go/src/encoding/json/decode.go:1290
				// _ = "end of CoverTab[27405]"
			}
//line /usr/local/go/src/encoding/json/decode.go:1291
			// _ = "end of CoverTab[27390]"

//line /usr/local/go/src/encoding/json/decode.go:1294
		case c == '"', c < ' ':
//line /usr/local/go/src/encoding/json/decode.go:1294
			_go_fuzz_dep_.CoverTab[27391]++
									return
//line /usr/local/go/src/encoding/json/decode.go:1295
			// _ = "end of CoverTab[27391]"

//line /usr/local/go/src/encoding/json/decode.go:1298
		case c < utf8.RuneSelf:
//line /usr/local/go/src/encoding/json/decode.go:1298
			_go_fuzz_dep_.CoverTab[27392]++
									b[w] = c
									r++
									w++
//line /usr/local/go/src/encoding/json/decode.go:1301
			// _ = "end of CoverTab[27392]"

//line /usr/local/go/src/encoding/json/decode.go:1304
		default:
//line /usr/local/go/src/encoding/json/decode.go:1304
			_go_fuzz_dep_.CoverTab[27393]++
									rr, size := utf8.DecodeRune(s[r:])
									r += size
									w += utf8.EncodeRune(b[w:], rr)
//line /usr/local/go/src/encoding/json/decode.go:1307
			// _ = "end of CoverTab[27393]"
		}
//line /usr/local/go/src/encoding/json/decode.go:1308
		// _ = "end of CoverTab[27386]"
	}
//line /usr/local/go/src/encoding/json/decode.go:1309
	// _ = "end of CoverTab[27364]"
//line /usr/local/go/src/encoding/json/decode.go:1309
	_go_fuzz_dep_.CoverTab[27365]++
							return b[0:w], true
//line /usr/local/go/src/encoding/json/decode.go:1310
	// _ = "end of CoverTab[27365]"
}

//line /usr/local/go/src/encoding/json/decode.go:1311
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/encoding/json/decode.go:1311
var _ = _go_fuzz_dep_.CoverTab
