// Copyright 2018 Frank Schroeder. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:5
package properties

//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:5
)

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// Decode assigns property values to exported fields of a struct.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
//
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
// Decode traverses v recursively and returns an error if a value cannot be
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
// converted to the field type or a required value is missing for a field.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
//
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
// The following type dependent decodings are used:
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
//
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
// String, boolean, numeric fields have the value of the property key assigned.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
// The property key name is the name of the field. A different key and a default
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
// value can be set in the field's tag. Fields without default value are
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
// required. If the value cannot be converted to the field type an error is
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
// returned.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
//
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
// time.Duration fields have the result of time.ParseDuration() assigned.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
//
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
// time.Time fields have the vaule of time.Parse() assigned. The default layout
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
// is time.RFC3339 but can be set in the field's tag.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
//
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
// Arrays and slices of string, boolean, numeric, time.Duration and time.Time
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
// fields have the value interpreted as a comma separated list of values. The
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
// individual values are trimmed of whitespace and empty values are ignored. A
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
// default value can be provided as a semicolon separated list in the field's
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
// tag.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
//
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
// Struct fields are decoded recursively using the field name plus "." as
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
// prefix. The prefix (without dot) can be overridden in the field's tag.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
// Default values are not supported in the field's tag. Specify them on the
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
// fields of the inner struct instead.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
//
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
// Map fields must have a key of type string and are decoded recursively by
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
// using the field's name plus ".' as prefix and the next element of the key
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
// name as map key. The prefix (without dot) can be overridden in the field's
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
// tag. Default values are not supported.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
//
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
// Examples:
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
//
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
//	// Field is ignored.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
//	Field int `properties:"-"`
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
//
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
//	// Field is assigned value of 'Field'.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
//	Field int
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
//
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
//	// Field is assigned value of 'myName'.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
//	Field int `properties:"myName"`
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
//
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
//	// Field is assigned value of key 'myName' and has a default
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
//	// value 15 if the key does not exist.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
//	Field int `properties:"myName,default=15"`
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
//
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
//	// Field is assigned value of key 'Field' and has a default
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
//	// value 15 if the key does not exist.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
//	Field int `properties:",default=15"`
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
//
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
//	// Field is assigned value of key 'date' and the date
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
//	// is in format 2006-01-02
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
//	Field time.Time `properties:"date,layout=2006-01-02"`
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
//
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
//	// Field is assigned the non-empty and whitespace trimmed
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
//	// values of key 'Field' split by commas.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
//	Field []string
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
//
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
//	// Field is assigned the non-empty and whitespace trimmed
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
//	// values of key 'Field' split by commas and has a default
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
//	// value ["a", "b", "c"] if the key does not exist.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
//	Field []string `properties:",default=a;b;c"`
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
//
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
//	// Field is decoded recursively with "Field." as key prefix.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
//	Field SomeStruct
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
//
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
//	// Field is decoded recursively with "myName." as key prefix.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
//	Field SomeStruct `properties:"myName"`
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
//
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
//	// Field is decoded recursively with "Field." as key prefix
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
//	// and the next dotted element of the key as map key.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
//	Field map[string]string
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
//
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
//	// Field is decoded recursively with "myName." as key prefix
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
//	// and the next dotted element of the key as map key.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:15
//	Field map[string]string `properties:"myName"`
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:94
func (p *Properties) Decode(x interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:94
	_go_fuzz_dep_.CoverTab[115511]++
											t, v := reflect.TypeOf(x), reflect.ValueOf(x)
											if t.Kind() != reflect.Ptr || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:96
		_go_fuzz_dep_.CoverTab[115514]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:96
		return v.Elem().Type().Kind() != reflect.Struct
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:96
		// _ = "end of CoverTab[115514]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:96
	}() {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:96
		_go_fuzz_dep_.CoverTab[115515]++
												return fmt.Errorf("not a pointer to struct: %s", t)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:97
		// _ = "end of CoverTab[115515]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:98
		_go_fuzz_dep_.CoverTab[115516]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:98
		// _ = "end of CoverTab[115516]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:98
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:98
	// _ = "end of CoverTab[115511]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:98
	_go_fuzz_dep_.CoverTab[115512]++
											if err := dec(p, "", nil, nil, v); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:99
			_go_fuzz_dep_.CoverTab[115517]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:100
		// _ = "end of CoverTab[115517]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:101
		_go_fuzz_dep_.CoverTab[115518]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:101
		// _ = "end of CoverTab[115518]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:101
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:101
	// _ = "end of CoverTab[115512]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:101
	_go_fuzz_dep_.CoverTab[115513]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:102
	// _ = "end of CoverTab[115513]"
}

func dec(p *Properties, key string, def *string, opts map[string]string, v reflect.Value) error {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:105
	_go_fuzz_dep_.CoverTab[115519]++
												t := v.Type()

//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:109
	value := func() (string, error) {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:109
		_go_fuzz_dep_.CoverTab[115524]++
													if val, ok := p.Get(key); ok {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:110
			_go_fuzz_dep_.CoverTab[115527]++
														return val, nil
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:111
			// _ = "end of CoverTab[115527]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:112
			_go_fuzz_dep_.CoverTab[115528]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:112
			// _ = "end of CoverTab[115528]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:112
		}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:112
		// _ = "end of CoverTab[115524]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:112
		_go_fuzz_dep_.CoverTab[115525]++
													if def != nil {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:113
			_go_fuzz_dep_.CoverTab[115529]++
														return *def, nil
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:114
			// _ = "end of CoverTab[115529]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:115
			_go_fuzz_dep_.CoverTab[115530]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:115
			// _ = "end of CoverTab[115530]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:115
		}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:115
		// _ = "end of CoverTab[115525]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:115
		_go_fuzz_dep_.CoverTab[115526]++
													return "", fmt.Errorf("missing required key %s", key)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:116
		// _ = "end of CoverTab[115526]"
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:117
	// _ = "end of CoverTab[115519]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:117
	_go_fuzz_dep_.CoverTab[115520]++

//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:120
	conv := func(s string, t reflect.Type) (val reflect.Value, err error) {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:120
		_go_fuzz_dep_.CoverTab[115531]++
													var v interface{}

													switch {
		case isDuration(t):
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:124
			_go_fuzz_dep_.CoverTab[115534]++
														v, err = time.ParseDuration(s)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:125
			// _ = "end of CoverTab[115534]"

		case isTime(t):
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:127
			_go_fuzz_dep_.CoverTab[115535]++
														layout := opts["layout"]
														if layout == "" {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:129
				_go_fuzz_dep_.CoverTab[115543]++
															layout = time.RFC3339
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:130
				// _ = "end of CoverTab[115543]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:131
				_go_fuzz_dep_.CoverTab[115544]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:131
				// _ = "end of CoverTab[115544]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:131
			}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:131
			// _ = "end of CoverTab[115535]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:131
			_go_fuzz_dep_.CoverTab[115536]++
														v, err = time.Parse(layout, s)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:132
			// _ = "end of CoverTab[115536]"

		case isBool(t):
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:134
			_go_fuzz_dep_.CoverTab[115537]++
														v, err = boolVal(s), nil
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:135
			// _ = "end of CoverTab[115537]"

		case isString(t):
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:137
			_go_fuzz_dep_.CoverTab[115538]++
														v, err = s, nil
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:138
			// _ = "end of CoverTab[115538]"

		case isFloat(t):
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:140
			_go_fuzz_dep_.CoverTab[115539]++
														v, err = strconv.ParseFloat(s, 64)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:141
			// _ = "end of CoverTab[115539]"

		case isInt(t):
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:143
			_go_fuzz_dep_.CoverTab[115540]++
														v, err = strconv.ParseInt(s, 10, 64)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:144
			// _ = "end of CoverTab[115540]"

		case isUint(t):
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:146
			_go_fuzz_dep_.CoverTab[115541]++
														v, err = strconv.ParseUint(s, 10, 64)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:147
			// _ = "end of CoverTab[115541]"

		default:
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:149
			_go_fuzz_dep_.CoverTab[115542]++
														return reflect.Zero(t), fmt.Errorf("unsupported type %s", t)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:150
			// _ = "end of CoverTab[115542]"
		}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:151
		// _ = "end of CoverTab[115531]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:151
		_go_fuzz_dep_.CoverTab[115532]++
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:152
			_go_fuzz_dep_.CoverTab[115545]++
														return reflect.Zero(t), err
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:153
			// _ = "end of CoverTab[115545]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:154
			_go_fuzz_dep_.CoverTab[115546]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:154
			// _ = "end of CoverTab[115546]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:154
		}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:154
		// _ = "end of CoverTab[115532]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:154
		_go_fuzz_dep_.CoverTab[115533]++
													return reflect.ValueOf(v).Convert(t), nil
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:155
		// _ = "end of CoverTab[115533]"
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:156
	// _ = "end of CoverTab[115520]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:156
	_go_fuzz_dep_.CoverTab[115521]++

//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:160
	keydef := func(f reflect.StructField) (string, *string, map[string]string) {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:160
		_go_fuzz_dep_.CoverTab[115547]++
													_key, _opts := parseTag(f.Tag.Get("properties"))

													var _def *string
													if d, ok := _opts["default"]; ok {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:164
			_go_fuzz_dep_.CoverTab[115550]++
														_def = &d
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:165
			// _ = "end of CoverTab[115550]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:166
			_go_fuzz_dep_.CoverTab[115551]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:166
			// _ = "end of CoverTab[115551]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:166
		}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:166
		// _ = "end of CoverTab[115547]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:166
		_go_fuzz_dep_.CoverTab[115548]++
													if _key != "" {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:167
			_go_fuzz_dep_.CoverTab[115552]++
														return _key, _def, _opts
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:168
			// _ = "end of CoverTab[115552]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:169
			_go_fuzz_dep_.CoverTab[115553]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:169
			// _ = "end of CoverTab[115553]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:169
		}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:169
		// _ = "end of CoverTab[115548]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:169
		_go_fuzz_dep_.CoverTab[115549]++
													return f.Name, _def, _opts
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:170
		// _ = "end of CoverTab[115549]"
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:171
	// _ = "end of CoverTab[115521]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:171
	_go_fuzz_dep_.CoverTab[115522]++

												switch {
	case isDuration(t) || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:174
		_go_fuzz_dep_.CoverTab[115566]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:174
		return isTime(t)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:174
		// _ = "end of CoverTab[115566]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:174
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:174
		_go_fuzz_dep_.CoverTab[115567]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:174
		return isBool(t)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:174
		// _ = "end of CoverTab[115567]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:174
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:174
		_go_fuzz_dep_.CoverTab[115568]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:174
		return isString(t)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:174
		// _ = "end of CoverTab[115568]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:174
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:174
		_go_fuzz_dep_.CoverTab[115569]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:174
		return isFloat(t)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:174
		// _ = "end of CoverTab[115569]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:174
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:174
		_go_fuzz_dep_.CoverTab[115570]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:174
		return isInt(t)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:174
		// _ = "end of CoverTab[115570]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:174
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:174
		_go_fuzz_dep_.CoverTab[115571]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:174
		return isUint(t)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:174
		// _ = "end of CoverTab[115571]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:174
	}():
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:174
		_go_fuzz_dep_.CoverTab[115554]++
													s, err := value()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:176
			_go_fuzz_dep_.CoverTab[115572]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:177
			// _ = "end of CoverTab[115572]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:178
			_go_fuzz_dep_.CoverTab[115573]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:178
			// _ = "end of CoverTab[115573]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:178
		}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:178
		// _ = "end of CoverTab[115554]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:178
		_go_fuzz_dep_.CoverTab[115555]++
													val, err := conv(s, t)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:180
			_go_fuzz_dep_.CoverTab[115574]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:181
			// _ = "end of CoverTab[115574]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:182
			_go_fuzz_dep_.CoverTab[115575]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:182
			// _ = "end of CoverTab[115575]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:182
		}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:182
		// _ = "end of CoverTab[115555]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:182
		_go_fuzz_dep_.CoverTab[115556]++
													v.Set(val)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:183
		// _ = "end of CoverTab[115556]"

	case isPtr(t):
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:185
		_go_fuzz_dep_.CoverTab[115557]++
													return dec(p, key, def, opts, v.Elem())
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:186
		// _ = "end of CoverTab[115557]"

	case isStruct(t):
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:188
		_go_fuzz_dep_.CoverTab[115558]++
													for i := 0; i < v.NumField(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:189
			_go_fuzz_dep_.CoverTab[115576]++
														fv := v.Field(i)
														fk, def, opts := keydef(t.Field(i))
														if !fv.CanSet() {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:192
				_go_fuzz_dep_.CoverTab[115580]++
															return fmt.Errorf("cannot set %s", t.Field(i).Name)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:193
				// _ = "end of CoverTab[115580]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:194
				_go_fuzz_dep_.CoverTab[115581]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:194
				// _ = "end of CoverTab[115581]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:194
			}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:194
			// _ = "end of CoverTab[115576]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:194
			_go_fuzz_dep_.CoverTab[115577]++
														if fk == "-" {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:195
				_go_fuzz_dep_.CoverTab[115582]++
															continue
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:196
				// _ = "end of CoverTab[115582]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:197
				_go_fuzz_dep_.CoverTab[115583]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:197
				// _ = "end of CoverTab[115583]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:197
			}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:197
			// _ = "end of CoverTab[115577]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:197
			_go_fuzz_dep_.CoverTab[115578]++
														if key != "" {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:198
				_go_fuzz_dep_.CoverTab[115584]++
															fk = key + "." + fk
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:199
				// _ = "end of CoverTab[115584]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:200
				_go_fuzz_dep_.CoverTab[115585]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:200
				// _ = "end of CoverTab[115585]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:200
			}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:200
			// _ = "end of CoverTab[115578]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:200
			_go_fuzz_dep_.CoverTab[115579]++
														if err := dec(p, fk, def, opts, fv); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:201
				_go_fuzz_dep_.CoverTab[115586]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:202
				// _ = "end of CoverTab[115586]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:203
				_go_fuzz_dep_.CoverTab[115587]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:203
				// _ = "end of CoverTab[115587]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:203
			}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:203
			// _ = "end of CoverTab[115579]"
		}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:204
		// _ = "end of CoverTab[115558]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:204
		_go_fuzz_dep_.CoverTab[115559]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:205
		// _ = "end of CoverTab[115559]"

	case isArray(t):
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:207
		_go_fuzz_dep_.CoverTab[115560]++
													val, err := value()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:209
			_go_fuzz_dep_.CoverTab[115588]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:210
			// _ = "end of CoverTab[115588]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:211
			_go_fuzz_dep_.CoverTab[115589]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:211
			// _ = "end of CoverTab[115589]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:211
		}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:211
		// _ = "end of CoverTab[115560]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:211
		_go_fuzz_dep_.CoverTab[115561]++
													vals := split(val, ";")
													a := reflect.MakeSlice(t, 0, len(vals))
													for _, s := range vals {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:214
			_go_fuzz_dep_.CoverTab[115590]++
														val, err := conv(s, t.Elem())
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:216
				_go_fuzz_dep_.CoverTab[115592]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:217
				// _ = "end of CoverTab[115592]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:218
				_go_fuzz_dep_.CoverTab[115593]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:218
				// _ = "end of CoverTab[115593]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:218
			}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:218
			// _ = "end of CoverTab[115590]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:218
			_go_fuzz_dep_.CoverTab[115591]++
														a = reflect.Append(a, val)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:219
			// _ = "end of CoverTab[115591]"
		}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:220
		// _ = "end of CoverTab[115561]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:220
		_go_fuzz_dep_.CoverTab[115562]++
													v.Set(a)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:221
		// _ = "end of CoverTab[115562]"

	case isMap(t):
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:223
		_go_fuzz_dep_.CoverTab[115563]++
													valT := t.Elem()
													m := reflect.MakeMap(t)
													for postfix := range p.FilterStripPrefix(key + ".").m {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:226
			_go_fuzz_dep_.CoverTab[115594]++
														pp := strings.SplitN(postfix, ".", 2)
														mk, mv := pp[0], reflect.New(valT)
														if err := dec(p, key+"."+mk, nil, nil, mv); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:229
				_go_fuzz_dep_.CoverTab[115596]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:230
				// _ = "end of CoverTab[115596]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:231
				_go_fuzz_dep_.CoverTab[115597]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:231
				// _ = "end of CoverTab[115597]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:231
			}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:231
			// _ = "end of CoverTab[115594]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:231
			_go_fuzz_dep_.CoverTab[115595]++
														m.SetMapIndex(reflect.ValueOf(mk), mv.Elem())
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:232
			// _ = "end of CoverTab[115595]"
		}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:233
		// _ = "end of CoverTab[115563]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:233
		_go_fuzz_dep_.CoverTab[115564]++
													v.Set(m)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:234
		// _ = "end of CoverTab[115564]"

	default:
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:236
		_go_fuzz_dep_.CoverTab[115565]++
													return fmt.Errorf("unsupported type %s", t)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:237
		// _ = "end of CoverTab[115565]"
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:238
	// _ = "end of CoverTab[115522]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:238
	_go_fuzz_dep_.CoverTab[115523]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:239
	// _ = "end of CoverTab[115523]"
}

// split splits a string on sep, trims whitespace of elements
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:242
// and omits empty elements
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:244
func split(s string, sep string) []string {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:244
	_go_fuzz_dep_.CoverTab[115598]++
												var a []string
												for _, v := range strings.Split(s, sep) {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:246
		_go_fuzz_dep_.CoverTab[115600]++
													if v = strings.TrimSpace(v); v != "" {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:247
			_go_fuzz_dep_.CoverTab[115601]++
														a = append(a, v)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:248
			// _ = "end of CoverTab[115601]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:249
			_go_fuzz_dep_.CoverTab[115602]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:249
			// _ = "end of CoverTab[115602]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:249
		}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:249
		// _ = "end of CoverTab[115600]"
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:250
	// _ = "end of CoverTab[115598]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:250
	_go_fuzz_dep_.CoverTab[115599]++
												return a
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:251
	// _ = "end of CoverTab[115599]"
}

// parseTag parses a "key,k=v,k=v,..."
func parseTag(tag string) (key string, opts map[string]string) {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:255
	_go_fuzz_dep_.CoverTab[115603]++
												opts = map[string]string{}
												for i, s := range strings.Split(tag, ",") {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:257
		_go_fuzz_dep_.CoverTab[115605]++
													if i == 0 {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:258
			_go_fuzz_dep_.CoverTab[115607]++
														key = s
														continue
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:260
			// _ = "end of CoverTab[115607]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:261
			_go_fuzz_dep_.CoverTab[115608]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:261
			// _ = "end of CoverTab[115608]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:261
		}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:261
		// _ = "end of CoverTab[115605]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:261
		_go_fuzz_dep_.CoverTab[115606]++

													pp := strings.SplitN(s, "=", 2)
													if len(pp) == 1 {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:264
			_go_fuzz_dep_.CoverTab[115609]++
														opts[pp[0]] = ""
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:265
			// _ = "end of CoverTab[115609]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:266
			_go_fuzz_dep_.CoverTab[115610]++
														opts[pp[0]] = pp[1]
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:267
			// _ = "end of CoverTab[115610]"
		}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:268
		// _ = "end of CoverTab[115606]"
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:269
	// _ = "end of CoverTab[115603]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:269
	_go_fuzz_dep_.CoverTab[115604]++
												return key, opts
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:270
	// _ = "end of CoverTab[115604]"
}

func isArray(t reflect.Type) bool {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:273
	_go_fuzz_dep_.CoverTab[115611]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:273
	return t.Kind() == reflect.Array || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:273
		_go_fuzz_dep_.CoverTab[115612]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:273
		return t.Kind() == reflect.Slice
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:273
		// _ = "end of CoverTab[115612]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:273
	}()
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:273
	// _ = "end of CoverTab[115611]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:273
}
func isBool(t reflect.Type) bool {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:274
	_go_fuzz_dep_.CoverTab[115613]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:274
	return t.Kind() == reflect.Bool
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:274
	// _ = "end of CoverTab[115613]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:274
}
func isDuration(t reflect.Type) bool {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:275
	_go_fuzz_dep_.CoverTab[115614]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:275
	return t == reflect.TypeOf(time.Second)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:275
	// _ = "end of CoverTab[115614]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:275
}
func isMap(t reflect.Type) bool {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:276
	_go_fuzz_dep_.CoverTab[115615]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:276
	return t.Kind() == reflect.Map
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:276
	// _ = "end of CoverTab[115615]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:276
}
func isPtr(t reflect.Type) bool {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:277
	_go_fuzz_dep_.CoverTab[115616]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:277
	return t.Kind() == reflect.Ptr
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:277
	// _ = "end of CoverTab[115616]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:277
}
func isString(t reflect.Type) bool {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:278
	_go_fuzz_dep_.CoverTab[115617]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:278
	return t.Kind() == reflect.String
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:278
	// _ = "end of CoverTab[115617]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:278
}
func isStruct(t reflect.Type) bool {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:279
	_go_fuzz_dep_.CoverTab[115618]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:279
	return t.Kind() == reflect.Struct
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:279
	// _ = "end of CoverTab[115618]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:279
}
func isTime(t reflect.Type) bool {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:280
	_go_fuzz_dep_.CoverTab[115619]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:280
	return t == reflect.TypeOf(time.Time{})
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:280
	// _ = "end of CoverTab[115619]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:280
}
func isFloat(t reflect.Type) bool {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:281
	_go_fuzz_dep_.CoverTab[115620]++
												return t.Kind() == reflect.Float32 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:282
		_go_fuzz_dep_.CoverTab[115621]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:282
		return t.Kind() == reflect.Float64
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:282
		// _ = "end of CoverTab[115621]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:282
	}()
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:282
	// _ = "end of CoverTab[115620]"
}
func isInt(t reflect.Type) bool {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:284
	_go_fuzz_dep_.CoverTab[115622]++
												return t.Kind() == reflect.Int || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:285
		_go_fuzz_dep_.CoverTab[115623]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:285
		return t.Kind() == reflect.Int8
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:285
		// _ = "end of CoverTab[115623]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:285
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:285
		_go_fuzz_dep_.CoverTab[115624]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:285
		return t.Kind() == reflect.Int16
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:285
		// _ = "end of CoverTab[115624]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:285
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:285
		_go_fuzz_dep_.CoverTab[115625]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:285
		return t.Kind() == reflect.Int32
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:285
		// _ = "end of CoverTab[115625]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:285
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:285
		_go_fuzz_dep_.CoverTab[115626]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:285
		return t.Kind() == reflect.Int64
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:285
		// _ = "end of CoverTab[115626]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:285
	}()
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:285
	// _ = "end of CoverTab[115622]"
}
func isUint(t reflect.Type) bool {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:287
	_go_fuzz_dep_.CoverTab[115627]++
												return t.Kind() == reflect.Uint || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:288
		_go_fuzz_dep_.CoverTab[115628]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:288
		return t.Kind() == reflect.Uint8
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:288
		// _ = "end of CoverTab[115628]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:288
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:288
		_go_fuzz_dep_.CoverTab[115629]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:288
		return t.Kind() == reflect.Uint16
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:288
		// _ = "end of CoverTab[115629]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:288
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:288
		_go_fuzz_dep_.CoverTab[115630]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:288
		return t.Kind() == reflect.Uint32
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:288
		// _ = "end of CoverTab[115630]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:288
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:288
		_go_fuzz_dep_.CoverTab[115631]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:288
		return t.Kind() == reflect.Uint64
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:288
		// _ = "end of CoverTab[115631]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:288
	}()
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:288
	// _ = "end of CoverTab[115627]"
}

//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:289
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/decode.go:289
var _ = _go_fuzz_dep_.CoverTab
