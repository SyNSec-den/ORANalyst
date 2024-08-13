// Copyright 2015 xeipuuv ( https://github.com/xeipuuv )
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// author           xeipuuv
// author-github    https://github.com/xeipuuv
// author-mail      xeipuuv@gmail.com
//
// repository-name  gojsonschema
// repository-desc  An implementation of JSON Schema, based on IETF's draft v4 - Go language.
//
// description      Result and ResultError implementations.
//
// created          01-01-2015

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:26
package gojsonschema

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:26
import (
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:26
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:26
)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:26
import (
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:26
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:26
)

import (
	"fmt"
	"strings"
)

type (
	// ErrorDetails is a map of details specific to each error.
	// While the values will vary, every error will contain a "field" value
	ErrorDetails	map[string]interface{}

	// ResultError is the interface that library errors must implement
	ResultError	interface {
		// Field returns the field name without the root context
		// i.e. firstName or person.firstName instead of (root).firstName or (root).person.firstName
		Field() string
		// SetType sets the error-type
		SetType(string)
		// Type returns the error-type
		Type() string
		// SetContext sets the JSON-context for the error
		SetContext(*JsonContext)
		// Context returns the JSON-context of the error
		Context() *JsonContext
		// SetDescription sets a description for the error
		SetDescription(string)
		// Description returns the description of the error
		Description() string
		// SetDescriptionFormat sets the format for the description in the default text/template format
		SetDescriptionFormat(string)
		// DescriptionFormat returns the format for the description in the default text/template format
		DescriptionFormat() string
		// SetValue sets the value related to the error
		SetValue(interface{})
		// Value returns the value related to the error
		Value() interface{}
		// SetDetails sets the details specific to the error
		SetDetails(ErrorDetails)
		// Details returns details about the error
		Details() ErrorDetails
		// String returns a string representation of the error
		String() string
	}

	// ResultErrorFields holds the fields for each ResultError implementation.
	// ResultErrorFields implements the ResultError interface, so custom errors
	// can be defined by just embedding this type
	ResultErrorFields	struct {
		errorType		string		// A string with the type of error (i.e. invalid_type)
		context			*JsonContext	// Tree like notation of the part that failed the validation. ex (root).a.b ...
		description		string		// A human readable error message
		descriptionFormat	string		// A format for human readable error message
		value			interface{}	// Value given by the JSON file that is the source of the error
		details			ErrorDetails
	}

	// Result holds the result of a validation
	Result	struct {
		errors	[]ResultError
		// Scores how well the validation matched. Useful in generating
		// better error messages for anyOf and oneOf.
		score	int
	}
)

// Field returns the field name without the root context
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:92
// i.e. firstName or person.firstName instead of (root).firstName or (root).person.firstName
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:94
func (v *ResultErrorFields) Field() string {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:94
	_go_fuzz_dep_.CoverTab[195054]++
											return strings.TrimPrefix(v.context.String(), STRING_ROOT_SCHEMA_PROPERTY+".")
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:95
	// _ = "end of CoverTab[195054]"
}

// SetType sets the error-type
func (v *ResultErrorFields) SetType(errorType string) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:99
	_go_fuzz_dep_.CoverTab[195055]++
											v.errorType = errorType
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:100
	// _ = "end of CoverTab[195055]"
}

// Type returns the error-type
func (v *ResultErrorFields) Type() string {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:104
	_go_fuzz_dep_.CoverTab[195056]++
											return v.errorType
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:105
	// _ = "end of CoverTab[195056]"
}

// SetContext sets the JSON-context for the error
func (v *ResultErrorFields) SetContext(context *JsonContext) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:109
	_go_fuzz_dep_.CoverTab[195057]++
											v.context = context
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:110
	// _ = "end of CoverTab[195057]"
}

// Context returns the JSON-context of the error
func (v *ResultErrorFields) Context() *JsonContext {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:114
	_go_fuzz_dep_.CoverTab[195058]++
											return v.context
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:115
	// _ = "end of CoverTab[195058]"
}

// SetDescription sets a description for the error
func (v *ResultErrorFields) SetDescription(description string) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:119
	_go_fuzz_dep_.CoverTab[195059]++
											v.description = description
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:120
	// _ = "end of CoverTab[195059]"
}

// Description returns the description of the error
func (v *ResultErrorFields) Description() string {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:124
	_go_fuzz_dep_.CoverTab[195060]++
											return v.description
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:125
	// _ = "end of CoverTab[195060]"
}

// SetDescriptionFormat sets the format for the description in the default text/template format
func (v *ResultErrorFields) SetDescriptionFormat(descriptionFormat string) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:129
	_go_fuzz_dep_.CoverTab[195061]++
											v.descriptionFormat = descriptionFormat
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:130
	// _ = "end of CoverTab[195061]"
}

// DescriptionFormat returns the format for the description in the default text/template format
func (v *ResultErrorFields) DescriptionFormat() string {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:134
	_go_fuzz_dep_.CoverTab[195062]++
											return v.descriptionFormat
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:135
	// _ = "end of CoverTab[195062]"
}

// SetValue sets the value related to the error
func (v *ResultErrorFields) SetValue(value interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:139
	_go_fuzz_dep_.CoverTab[195063]++
											v.value = value
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:140
	// _ = "end of CoverTab[195063]"
}

// Value returns the value related to the error
func (v *ResultErrorFields) Value() interface{} {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:144
	_go_fuzz_dep_.CoverTab[195064]++
											return v.value
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:145
	// _ = "end of CoverTab[195064]"
}

// SetDetails sets the details specific to the error
func (v *ResultErrorFields) SetDetails(details ErrorDetails) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:149
	_go_fuzz_dep_.CoverTab[195065]++
											v.details = details
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:150
	// _ = "end of CoverTab[195065]"
}

// Details returns details about the error
func (v *ResultErrorFields) Details() ErrorDetails {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:154
	_go_fuzz_dep_.CoverTab[195066]++
											return v.details
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:155
	// _ = "end of CoverTab[195066]"
}

// String returns a string representation of the error
func (v ResultErrorFields) String() string {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:159
	_go_fuzz_dep_.CoverTab[195067]++

											valueString := fmt.Sprintf("%v", v.value)

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:164
	if v.value == nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:164
		_go_fuzz_dep_.CoverTab[195069]++
												valueString = TYPE_NULL
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:165
		// _ = "end of CoverTab[195069]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:166
		_go_fuzz_dep_.CoverTab[195070]++
												if vs, err := marshalToJSONString(v.value); err == nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:167
			_go_fuzz_dep_.CoverTab[195071]++
													if vs == nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:168
				_go_fuzz_dep_.CoverTab[195072]++
														valueString = TYPE_NULL
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:169
				// _ = "end of CoverTab[195072]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:170
				_go_fuzz_dep_.CoverTab[195073]++
														valueString = *vs
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:171
				// _ = "end of CoverTab[195073]"
			}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:172
			// _ = "end of CoverTab[195071]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:173
			_go_fuzz_dep_.CoverTab[195074]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:173
			// _ = "end of CoverTab[195074]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:173
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:173
		// _ = "end of CoverTab[195070]"
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:174
	// _ = "end of CoverTab[195067]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:174
	_go_fuzz_dep_.CoverTab[195068]++

											return formatErrorDescription(Locale.ErrorFormat(), ErrorDetails{
		"context":	v.context.String(),
		"description":	v.description,
		"value":	valueString,
		"field":	v.Field(),
	})
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:181
	// _ = "end of CoverTab[195068]"
}

// Valid indicates if no errors were found
func (v *Result) Valid() bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:185
	_go_fuzz_dep_.CoverTab[195075]++
											return len(v.errors) == 0
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:186
	// _ = "end of CoverTab[195075]"
}

// Errors returns the errors that were found
func (v *Result) Errors() []ResultError {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:190
	_go_fuzz_dep_.CoverTab[195076]++
											return v.errors
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:191
	// _ = "end of CoverTab[195076]"
}

// AddError appends a fully filled error to the error set
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:194
// SetDescription() will be called with the result of the parsed err.DescriptionFormat()
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:196
func (v *Result) AddError(err ResultError, details ErrorDetails) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:196
	_go_fuzz_dep_.CoverTab[195077]++
											if _, exists := details["context"]; !exists && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:197
		_go_fuzz_dep_.CoverTab[195079]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:197
		return err.Context() != nil
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:197
		// _ = "end of CoverTab[195079]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:197
	}() {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:197
		_go_fuzz_dep_.CoverTab[195080]++
												details["context"] = err.Context().String()
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:198
		// _ = "end of CoverTab[195080]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:199
		_go_fuzz_dep_.CoverTab[195081]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:199
		// _ = "end of CoverTab[195081]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:199
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:199
	// _ = "end of CoverTab[195077]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:199
	_go_fuzz_dep_.CoverTab[195078]++

											err.SetDescription(formatErrorDescription(err.DescriptionFormat(), details))

											v.errors = append(v.errors, err)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:203
	// _ = "end of CoverTab[195078]"
}

func (v *Result) addInternalError(err ResultError, context *JsonContext, value interface{}, details ErrorDetails) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:206
	_go_fuzz_dep_.CoverTab[195082]++
											newError(err, context, value, Locale, details)
											v.errors = append(v.errors, err)
											v.score -= 2
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:209
	// _ = "end of CoverTab[195082]"
}

// Used to copy errors from a sub-schema to the main one
func (v *Result) mergeErrors(otherResult *Result) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:213
	_go_fuzz_dep_.CoverTab[195083]++
											v.errors = append(v.errors, otherResult.Errors()...)
											v.score += otherResult.score
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:215
	// _ = "end of CoverTab[195083]"
}

func (v *Result) incrementScore() {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:218
	_go_fuzz_dep_.CoverTab[195084]++
											v.score++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:219
	// _ = "end of CoverTab[195084]"
}

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:220
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/result.go:220
var _ = _go_fuzz_dep_.CoverTab
