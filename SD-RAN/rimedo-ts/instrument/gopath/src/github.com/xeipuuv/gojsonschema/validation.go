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
// description      Extends Schema and subSchema, implements the validation phase.
//
// created          28-02-2013

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:26
package gojsonschema

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:26
import (
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:26
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:26
)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:26
import (
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:26
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:26
)

import (
	"encoding/json"
	"math/big"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

// Validate loads and validates a JSON schema
func Validate(ls JSONLoader, ld JSONLoader) (*Result, error) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:39
	_go_fuzz_dep_.CoverTab[195790]++

												schema, err := NewSchema(ls)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:42
		_go_fuzz_dep_.CoverTab[195792]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:43
		// _ = "end of CoverTab[195792]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:44
		_go_fuzz_dep_.CoverTab[195793]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:44
		// _ = "end of CoverTab[195793]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:44
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:44
	// _ = "end of CoverTab[195790]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:44
	_go_fuzz_dep_.CoverTab[195791]++
												return schema.Validate(ld)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:45
	// _ = "end of CoverTab[195791]"
}

// Validate loads and validates a JSON document
func (v *Schema) Validate(l JSONLoader) (*Result, error) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:49
	_go_fuzz_dep_.CoverTab[195794]++
												root, err := l.LoadJSON()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:51
		_go_fuzz_dep_.CoverTab[195796]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:52
		// _ = "end of CoverTab[195796]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:53
		_go_fuzz_dep_.CoverTab[195797]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:53
		// _ = "end of CoverTab[195797]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:53
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:53
	// _ = "end of CoverTab[195794]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:53
	_go_fuzz_dep_.CoverTab[195795]++
												return v.validateDocument(root), nil
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:54
	// _ = "end of CoverTab[195795]"
}

func (v *Schema) validateDocument(root interface{}) *Result {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:57
	_go_fuzz_dep_.CoverTab[195798]++
												result := &Result{}
												context := NewJsonContext(STRING_CONTEXT_ROOT, nil)
												v.rootSchema.validateRecursive(v.rootSchema, root, result, context)
												return result
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:61
	// _ = "end of CoverTab[195798]"
}

func (v *subSchema) subValidateWithContext(document interface{}, context *JsonContext) *Result {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:64
	_go_fuzz_dep_.CoverTab[195799]++
												result := &Result{}
												v.validateRecursive(v, document, result, context)
												return result
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:67
	// _ = "end of CoverTab[195799]"
}

// Walker function to validate the json recursively against the subSchema
func (v *subSchema) validateRecursive(currentSubSchema *subSchema, currentNode interface{}, result *Result, context *JsonContext) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:71
	_go_fuzz_dep_.CoverTab[195800]++

												if internalLogEnabled {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:73
		_go_fuzz_dep_.CoverTab[195805]++
													internalLog("validateRecursive %s", context.String())
													internalLog(" %v", currentNode)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:75
		// _ = "end of CoverTab[195805]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:76
		_go_fuzz_dep_.CoverTab[195806]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:76
		// _ = "end of CoverTab[195806]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:76
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:76
	// _ = "end of CoverTab[195800]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:76
	_go_fuzz_dep_.CoverTab[195801]++

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:79
	if currentSubSchema.pass != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:79
		_go_fuzz_dep_.CoverTab[195807]++
													if !*currentSubSchema.pass {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:80
			_go_fuzz_dep_.CoverTab[195809]++
														result.addInternalError(
				new(FalseError),
				context,
				currentNode,
				ErrorDetails{},
			)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:86
			// _ = "end of CoverTab[195809]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:87
			_go_fuzz_dep_.CoverTab[195810]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:87
			// _ = "end of CoverTab[195810]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:87
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:87
		// _ = "end of CoverTab[195807]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:87
		_go_fuzz_dep_.CoverTab[195808]++
													return
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:88
		// _ = "end of CoverTab[195808]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:89
		_go_fuzz_dep_.CoverTab[195811]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:89
		// _ = "end of CoverTab[195811]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:89
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:89
	// _ = "end of CoverTab[195801]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:89
	_go_fuzz_dep_.CoverTab[195802]++

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:92
	if currentSubSchema.refSchema != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:92
		_go_fuzz_dep_.CoverTab[195812]++
													v.validateRecursive(currentSubSchema.refSchema, currentNode, result, context)
													return
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:94
		// _ = "end of CoverTab[195812]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:95
		_go_fuzz_dep_.CoverTab[195813]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:95
		// _ = "end of CoverTab[195813]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:95
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:95
	// _ = "end of CoverTab[195802]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:95
	_go_fuzz_dep_.CoverTab[195803]++

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:98
	if currentNode == nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:98
		_go_fuzz_dep_.CoverTab[195814]++
													if currentSubSchema.types.IsTyped() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:99
			_go_fuzz_dep_.CoverTab[195816]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:99
			return !currentSubSchema.types.Contains(TYPE_NULL)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:99
			// _ = "end of CoverTab[195816]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:99
		}() {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:99
			_go_fuzz_dep_.CoverTab[195817]++
														result.addInternalError(
				new(InvalidTypeError),
				context,
				currentNode,
				ErrorDetails{
					"expected":	currentSubSchema.types.String(),
					"given":	TYPE_NULL,
				},
			)
														return
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:109
			// _ = "end of CoverTab[195817]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:110
			_go_fuzz_dep_.CoverTab[195818]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:110
			// _ = "end of CoverTab[195818]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:110
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:110
		// _ = "end of CoverTab[195814]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:110
		_go_fuzz_dep_.CoverTab[195815]++

													currentSubSchema.validateSchema(currentSubSchema, currentNode, result, context)
													v.validateCommon(currentSubSchema, currentNode, result, context)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:113
		// _ = "end of CoverTab[195815]"

	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:115
		_go_fuzz_dep_.CoverTab[195819]++

													if isJSONNumber(currentNode) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:117
			_go_fuzz_dep_.CoverTab[195820]++

														value := currentNode.(json.Number)

														isInt := checkJSONInteger(value)

														validType := currentSubSchema.types.Contains(TYPE_NUMBER) || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:123
				_go_fuzz_dep_.CoverTab[195822]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:123
				return (isInt && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:123
					_go_fuzz_dep_.CoverTab[195823]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:123
					return currentSubSchema.types.Contains(TYPE_INTEGER)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:123
					// _ = "end of CoverTab[195823]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:123
				}())
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:123
				// _ = "end of CoverTab[195822]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:123
			}()

														if currentSubSchema.types.IsTyped() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:125
				_go_fuzz_dep_.CoverTab[195824]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:125
				return !validType
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:125
				// _ = "end of CoverTab[195824]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:125
			}() {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:125
				_go_fuzz_dep_.CoverTab[195825]++

															givenType := TYPE_INTEGER
															if !isInt {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:128
					_go_fuzz_dep_.CoverTab[195827]++
																givenType = TYPE_NUMBER
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:129
					// _ = "end of CoverTab[195827]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:130
					_go_fuzz_dep_.CoverTab[195828]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:130
					// _ = "end of CoverTab[195828]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:130
				}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:130
				// _ = "end of CoverTab[195825]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:130
				_go_fuzz_dep_.CoverTab[195826]++

															result.addInternalError(
					new(InvalidTypeError),
					context,
					currentNode,
					ErrorDetails{
						"expected":	currentSubSchema.types.String(),
						"given":	givenType,
					},
				)
															return
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:141
				// _ = "end of CoverTab[195826]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:142
				_go_fuzz_dep_.CoverTab[195829]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:142
				// _ = "end of CoverTab[195829]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:142
			}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:142
			// _ = "end of CoverTab[195820]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:142
			_go_fuzz_dep_.CoverTab[195821]++

														currentSubSchema.validateSchema(currentSubSchema, value, result, context)
														v.validateNumber(currentSubSchema, value, result, context)
														v.validateCommon(currentSubSchema, value, result, context)
														v.validateString(currentSubSchema, value, result, context)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:147
			// _ = "end of CoverTab[195821]"

		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:149
			_go_fuzz_dep_.CoverTab[195830]++

														rValue := reflect.ValueOf(currentNode)
														rKind := rValue.Kind()

														switch rKind {

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:158
			case reflect.Slice:
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:158
				_go_fuzz_dep_.CoverTab[195831]++

															if currentSubSchema.types.IsTyped() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:160
					_go_fuzz_dep_.CoverTab[195841]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:160
					return !currentSubSchema.types.Contains(TYPE_ARRAY)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:160
					// _ = "end of CoverTab[195841]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:160
				}() {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:160
					_go_fuzz_dep_.CoverTab[195842]++
																result.addInternalError(
						new(InvalidTypeError),
						context,
						currentNode,
						ErrorDetails{
							"expected":	currentSubSchema.types.String(),
							"given":	TYPE_ARRAY,
						},
					)
																return
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:170
					// _ = "end of CoverTab[195842]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:171
					_go_fuzz_dep_.CoverTab[195843]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:171
					// _ = "end of CoverTab[195843]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:171
				}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:171
				// _ = "end of CoverTab[195831]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:171
				_go_fuzz_dep_.CoverTab[195832]++

															castCurrentNode := currentNode.([]interface{})

															currentSubSchema.validateSchema(currentSubSchema, castCurrentNode, result, context)

															v.validateArray(currentSubSchema, castCurrentNode, result, context)
															v.validateCommon(currentSubSchema, castCurrentNode, result, context)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:178
				// _ = "end of CoverTab[195832]"

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:182
			case reflect.Map:
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:182
				_go_fuzz_dep_.CoverTab[195833]++
															if currentSubSchema.types.IsTyped() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:183
					_go_fuzz_dep_.CoverTab[195844]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:183
					return !currentSubSchema.types.Contains(TYPE_OBJECT)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:183
					// _ = "end of CoverTab[195844]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:183
				}() {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:183
					_go_fuzz_dep_.CoverTab[195845]++
																result.addInternalError(
						new(InvalidTypeError),
						context,
						currentNode,
						ErrorDetails{
							"expected":	currentSubSchema.types.String(),
							"given":	TYPE_OBJECT,
						},
					)
																return
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:193
					// _ = "end of CoverTab[195845]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:194
					_go_fuzz_dep_.CoverTab[195846]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:194
					// _ = "end of CoverTab[195846]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:194
				}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:194
				// _ = "end of CoverTab[195833]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:194
				_go_fuzz_dep_.CoverTab[195834]++

															castCurrentNode, ok := currentNode.(map[string]interface{})
															if !ok {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:197
					_go_fuzz_dep_.CoverTab[195847]++
																castCurrentNode = convertDocumentNode(currentNode).(map[string]interface{})
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:198
					// _ = "end of CoverTab[195847]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:199
					_go_fuzz_dep_.CoverTab[195848]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:199
					// _ = "end of CoverTab[195848]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:199
				}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:199
				// _ = "end of CoverTab[195834]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:199
				_go_fuzz_dep_.CoverTab[195835]++

															currentSubSchema.validateSchema(currentSubSchema, castCurrentNode, result, context)

															v.validateObject(currentSubSchema, castCurrentNode, result, context)
															v.validateCommon(currentSubSchema, castCurrentNode, result, context)

															for _, pSchema := range currentSubSchema.propertiesChildren {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:206
					_go_fuzz_dep_.CoverTab[195849]++
																nextNode, ok := castCurrentNode[pSchema.property]
																if ok {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:208
						_go_fuzz_dep_.CoverTab[195850]++
																	subContext := NewJsonContext(pSchema.property, context)
																	v.validateRecursive(pSchema, nextNode, result, subContext)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:210
						// _ = "end of CoverTab[195850]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:211
						_go_fuzz_dep_.CoverTab[195851]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:211
						// _ = "end of CoverTab[195851]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:211
					}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:211
					// _ = "end of CoverTab[195849]"
				}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:212
				// _ = "end of CoverTab[195835]"

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:216
			case reflect.Bool:
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:216
				_go_fuzz_dep_.CoverTab[195836]++

															if currentSubSchema.types.IsTyped() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:218
					_go_fuzz_dep_.CoverTab[195852]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:218
					return !currentSubSchema.types.Contains(TYPE_BOOLEAN)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:218
					// _ = "end of CoverTab[195852]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:218
				}() {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:218
					_go_fuzz_dep_.CoverTab[195853]++
																result.addInternalError(
						new(InvalidTypeError),
						context,
						currentNode,
						ErrorDetails{
							"expected":	currentSubSchema.types.String(),
							"given":	TYPE_BOOLEAN,
						},
					)
																return
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:228
					// _ = "end of CoverTab[195853]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:229
					_go_fuzz_dep_.CoverTab[195854]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:229
					// _ = "end of CoverTab[195854]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:229
				}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:229
				// _ = "end of CoverTab[195836]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:229
				_go_fuzz_dep_.CoverTab[195837]++

															value := currentNode.(bool)

															currentSubSchema.validateSchema(currentSubSchema, value, result, context)
															v.validateNumber(currentSubSchema, value, result, context)
															v.validateCommon(currentSubSchema, value, result, context)
															v.validateString(currentSubSchema, value, result, context)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:236
				// _ = "end of CoverTab[195837]"

			case reflect.String:
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:238
				_go_fuzz_dep_.CoverTab[195838]++

															if currentSubSchema.types.IsTyped() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:240
					_go_fuzz_dep_.CoverTab[195855]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:240
					return !currentSubSchema.types.Contains(TYPE_STRING)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:240
					// _ = "end of CoverTab[195855]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:240
				}() {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:240
					_go_fuzz_dep_.CoverTab[195856]++
																result.addInternalError(
						new(InvalidTypeError),
						context,
						currentNode,
						ErrorDetails{
							"expected":	currentSubSchema.types.String(),
							"given":	TYPE_STRING,
						},
					)
																return
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:250
					// _ = "end of CoverTab[195856]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:251
					_go_fuzz_dep_.CoverTab[195857]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:251
					// _ = "end of CoverTab[195857]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:251
				}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:251
				// _ = "end of CoverTab[195838]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:251
				_go_fuzz_dep_.CoverTab[195839]++

															value := currentNode.(string)

															currentSubSchema.validateSchema(currentSubSchema, value, result, context)
															v.validateNumber(currentSubSchema, value, result, context)
															v.validateCommon(currentSubSchema, value, result, context)
															v.validateString(currentSubSchema, value, result, context)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:258
				// _ = "end of CoverTab[195839]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:258
			default:
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:258
				_go_fuzz_dep_.CoverTab[195840]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:258
				// _ = "end of CoverTab[195840]"

			}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:260
			// _ = "end of CoverTab[195830]"

		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:262
		// _ = "end of CoverTab[195819]"

	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:264
	// _ = "end of CoverTab[195803]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:264
	_go_fuzz_dep_.CoverTab[195804]++

												result.incrementScore()
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:266
	// _ = "end of CoverTab[195804]"
}

// Different kinds of validation there, subSchema / common / array / object / string...
func (v *subSchema) validateSchema(currentSubSchema *subSchema, currentNode interface{}, result *Result, context *JsonContext) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:270
	_go_fuzz_dep_.CoverTab[195858]++

												if internalLogEnabled {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:272
		_go_fuzz_dep_.CoverTab[195866]++
													internalLog("validateSchema %s", context.String())
													internalLog(" %v", currentNode)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:274
		// _ = "end of CoverTab[195866]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:275
		_go_fuzz_dep_.CoverTab[195867]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:275
		// _ = "end of CoverTab[195867]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:275
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:275
	// _ = "end of CoverTab[195858]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:275
	_go_fuzz_dep_.CoverTab[195859]++

												if len(currentSubSchema.anyOf) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:277
		_go_fuzz_dep_.CoverTab[195868]++

													validatedAnyOf := false
													var bestValidationResult *Result

													for _, anyOfSchema := range currentSubSchema.anyOf {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:282
			_go_fuzz_dep_.CoverTab[195870]++
														if !validatedAnyOf {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:283
				_go_fuzz_dep_.CoverTab[195871]++
															validationResult := anyOfSchema.subValidateWithContext(currentNode, context)
															validatedAnyOf = validationResult.Valid()

															if !validatedAnyOf && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:287
					_go_fuzz_dep_.CoverTab[195872]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:287
					return (bestValidationResult == nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:287
						_go_fuzz_dep_.CoverTab[195873]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:287
						return validationResult.score > bestValidationResult.score
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:287
						// _ = "end of CoverTab[195873]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:287
					}())
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:287
					// _ = "end of CoverTab[195872]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:287
				}() {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:287
					_go_fuzz_dep_.CoverTab[195874]++
																bestValidationResult = validationResult
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:288
					// _ = "end of CoverTab[195874]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:289
					_go_fuzz_dep_.CoverTab[195875]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:289
					// _ = "end of CoverTab[195875]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:289
				}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:289
				// _ = "end of CoverTab[195871]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:290
				_go_fuzz_dep_.CoverTab[195876]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:290
				// _ = "end of CoverTab[195876]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:290
			}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:290
			// _ = "end of CoverTab[195870]"
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:291
		// _ = "end of CoverTab[195868]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:291
		_go_fuzz_dep_.CoverTab[195869]++
													if !validatedAnyOf {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:292
			_go_fuzz_dep_.CoverTab[195877]++

														result.addInternalError(new(NumberAnyOfError), context, currentNode, ErrorDetails{})

														if bestValidationResult != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:296
				_go_fuzz_dep_.CoverTab[195878]++

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:299
				result.mergeErrors(bestValidationResult)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:299
				// _ = "end of CoverTab[195878]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:300
				_go_fuzz_dep_.CoverTab[195879]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:300
				// _ = "end of CoverTab[195879]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:300
			}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:300
			// _ = "end of CoverTab[195877]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:301
			_go_fuzz_dep_.CoverTab[195880]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:301
			// _ = "end of CoverTab[195880]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:301
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:301
		// _ = "end of CoverTab[195869]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:302
		_go_fuzz_dep_.CoverTab[195881]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:302
		// _ = "end of CoverTab[195881]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:302
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:302
	// _ = "end of CoverTab[195859]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:302
	_go_fuzz_dep_.CoverTab[195860]++

												if len(currentSubSchema.oneOf) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:304
		_go_fuzz_dep_.CoverTab[195882]++

													nbValidated := 0
													var bestValidationResult *Result

													for _, oneOfSchema := range currentSubSchema.oneOf {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:309
			_go_fuzz_dep_.CoverTab[195884]++
														validationResult := oneOfSchema.subValidateWithContext(currentNode, context)
														if validationResult.Valid() {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:311
				_go_fuzz_dep_.CoverTab[195885]++
															nbValidated++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:312
				// _ = "end of CoverTab[195885]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:313
				_go_fuzz_dep_.CoverTab[195886]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:313
				if nbValidated == 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:313
					_go_fuzz_dep_.CoverTab[195887]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:313
					return (bestValidationResult == nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:313
						_go_fuzz_dep_.CoverTab[195888]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:313
						return validationResult.score > bestValidationResult.score
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:313
						// _ = "end of CoverTab[195888]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:313
					}())
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:313
					// _ = "end of CoverTab[195887]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:313
				}() {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:313
					_go_fuzz_dep_.CoverTab[195889]++
																bestValidationResult = validationResult
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:314
					// _ = "end of CoverTab[195889]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:315
					_go_fuzz_dep_.CoverTab[195890]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:315
					// _ = "end of CoverTab[195890]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:315
				}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:315
				// _ = "end of CoverTab[195886]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:315
			}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:315
			// _ = "end of CoverTab[195884]"
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:316
		// _ = "end of CoverTab[195882]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:316
		_go_fuzz_dep_.CoverTab[195883]++

													if nbValidated != 1 {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:318
			_go_fuzz_dep_.CoverTab[195891]++

														result.addInternalError(new(NumberOneOfError), context, currentNode, ErrorDetails{})

														if nbValidated == 0 {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:322
				_go_fuzz_dep_.CoverTab[195892]++

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:325
				result.mergeErrors(bestValidationResult)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:325
				// _ = "end of CoverTab[195892]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:326
				_go_fuzz_dep_.CoverTab[195893]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:326
				// _ = "end of CoverTab[195893]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:326
			}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:326
			// _ = "end of CoverTab[195891]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:327
			_go_fuzz_dep_.CoverTab[195894]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:327
			// _ = "end of CoverTab[195894]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:327
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:327
		// _ = "end of CoverTab[195883]"

	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:329
		_go_fuzz_dep_.CoverTab[195895]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:329
		// _ = "end of CoverTab[195895]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:329
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:329
	// _ = "end of CoverTab[195860]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:329
	_go_fuzz_dep_.CoverTab[195861]++

												if len(currentSubSchema.allOf) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:331
		_go_fuzz_dep_.CoverTab[195896]++
													nbValidated := 0

													for _, allOfSchema := range currentSubSchema.allOf {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:334
			_go_fuzz_dep_.CoverTab[195898]++
														validationResult := allOfSchema.subValidateWithContext(currentNode, context)
														if validationResult.Valid() {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:336
				_go_fuzz_dep_.CoverTab[195900]++
															nbValidated++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:337
				// _ = "end of CoverTab[195900]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:338
				_go_fuzz_dep_.CoverTab[195901]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:338
				// _ = "end of CoverTab[195901]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:338
			}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:338
			// _ = "end of CoverTab[195898]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:338
			_go_fuzz_dep_.CoverTab[195899]++
														result.mergeErrors(validationResult)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:339
			// _ = "end of CoverTab[195899]"
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:340
		// _ = "end of CoverTab[195896]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:340
		_go_fuzz_dep_.CoverTab[195897]++

													if nbValidated != len(currentSubSchema.allOf) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:342
			_go_fuzz_dep_.CoverTab[195902]++
														result.addInternalError(new(NumberAllOfError), context, currentNode, ErrorDetails{})
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:343
			// _ = "end of CoverTab[195902]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:344
			_go_fuzz_dep_.CoverTab[195903]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:344
			// _ = "end of CoverTab[195903]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:344
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:344
		// _ = "end of CoverTab[195897]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:345
		_go_fuzz_dep_.CoverTab[195904]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:345
		// _ = "end of CoverTab[195904]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:345
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:345
	// _ = "end of CoverTab[195861]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:345
	_go_fuzz_dep_.CoverTab[195862]++

												if currentSubSchema.not != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:347
		_go_fuzz_dep_.CoverTab[195905]++
													validationResult := currentSubSchema.not.subValidateWithContext(currentNode, context)
													if validationResult.Valid() {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:349
			_go_fuzz_dep_.CoverTab[195906]++
														result.addInternalError(new(NumberNotError), context, currentNode, ErrorDetails{})
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:350
			// _ = "end of CoverTab[195906]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:351
			_go_fuzz_dep_.CoverTab[195907]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:351
			// _ = "end of CoverTab[195907]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:351
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:351
		// _ = "end of CoverTab[195905]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:352
		_go_fuzz_dep_.CoverTab[195908]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:352
		// _ = "end of CoverTab[195908]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:352
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:352
	// _ = "end of CoverTab[195862]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:352
	_go_fuzz_dep_.CoverTab[195863]++

												if currentSubSchema.dependencies != nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:354
		_go_fuzz_dep_.CoverTab[195909]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:354
		return len(currentSubSchema.dependencies) > 0
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:354
		// _ = "end of CoverTab[195909]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:354
	}() {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:354
		_go_fuzz_dep_.CoverTab[195910]++
													if isKind(currentNode, reflect.Map) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:355
			_go_fuzz_dep_.CoverTab[195911]++
														for elementKey := range currentNode.(map[string]interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:356
				_go_fuzz_dep_.CoverTab[195912]++
															if dependency, ok := currentSubSchema.dependencies[elementKey]; ok {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:357
					_go_fuzz_dep_.CoverTab[195913]++
																switch dependency := dependency.(type) {

					case []string:
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:360
						_go_fuzz_dep_.CoverTab[195914]++
																	for _, dependOnKey := range dependency {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:361
							_go_fuzz_dep_.CoverTab[195916]++
																		if _, dependencyResolved := currentNode.(map[string]interface{})[dependOnKey]; !dependencyResolved {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:362
								_go_fuzz_dep_.CoverTab[195917]++
																			result.addInternalError(
									new(MissingDependencyError),
									context,
									currentNode,
									ErrorDetails{"dependency": dependOnKey},
								)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:368
								// _ = "end of CoverTab[195917]"
							} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:369
								_go_fuzz_dep_.CoverTab[195918]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:369
								// _ = "end of CoverTab[195918]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:369
							}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:369
							// _ = "end of CoverTab[195916]"
						}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:370
						// _ = "end of CoverTab[195914]"

					case *subSchema:
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:372
						_go_fuzz_dep_.CoverTab[195915]++
																	dependency.validateRecursive(dependency, currentNode, result, context)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:373
						// _ = "end of CoverTab[195915]"
					}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:374
					// _ = "end of CoverTab[195913]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:375
					_go_fuzz_dep_.CoverTab[195919]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:375
					// _ = "end of CoverTab[195919]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:375
				}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:375
				// _ = "end of CoverTab[195912]"
			}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:376
			// _ = "end of CoverTab[195911]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:377
			_go_fuzz_dep_.CoverTab[195920]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:377
			// _ = "end of CoverTab[195920]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:377
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:377
		// _ = "end of CoverTab[195910]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:378
		_go_fuzz_dep_.CoverTab[195921]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:378
		// _ = "end of CoverTab[195921]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:378
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:378
	// _ = "end of CoverTab[195863]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:378
	_go_fuzz_dep_.CoverTab[195864]++

												if currentSubSchema._if != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:380
		_go_fuzz_dep_.CoverTab[195922]++
													validationResultIf := currentSubSchema._if.subValidateWithContext(currentNode, context)
													if currentSubSchema._then != nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:382
			_go_fuzz_dep_.CoverTab[195924]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:382
			return validationResultIf.Valid()
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:382
			// _ = "end of CoverTab[195924]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:382
		}() {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:382
			_go_fuzz_dep_.CoverTab[195925]++
														validationResultThen := currentSubSchema._then.subValidateWithContext(currentNode, context)
														if !validationResultThen.Valid() {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:384
				_go_fuzz_dep_.CoverTab[195926]++
															result.addInternalError(new(ConditionThenError), context, currentNode, ErrorDetails{})
															result.mergeErrors(validationResultThen)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:386
				// _ = "end of CoverTab[195926]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:387
				_go_fuzz_dep_.CoverTab[195927]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:387
				// _ = "end of CoverTab[195927]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:387
			}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:387
			// _ = "end of CoverTab[195925]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:388
			_go_fuzz_dep_.CoverTab[195928]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:388
			// _ = "end of CoverTab[195928]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:388
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:388
		// _ = "end of CoverTab[195922]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:388
		_go_fuzz_dep_.CoverTab[195923]++
													if currentSubSchema._else != nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:389
			_go_fuzz_dep_.CoverTab[195929]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:389
			return !validationResultIf.Valid()
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:389
			// _ = "end of CoverTab[195929]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:389
		}() {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:389
			_go_fuzz_dep_.CoverTab[195930]++
														validationResultElse := currentSubSchema._else.subValidateWithContext(currentNode, context)
														if !validationResultElse.Valid() {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:391
				_go_fuzz_dep_.CoverTab[195931]++
															result.addInternalError(new(ConditionElseError), context, currentNode, ErrorDetails{})
															result.mergeErrors(validationResultElse)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:393
				// _ = "end of CoverTab[195931]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:394
				_go_fuzz_dep_.CoverTab[195932]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:394
				// _ = "end of CoverTab[195932]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:394
			}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:394
			// _ = "end of CoverTab[195930]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:395
			_go_fuzz_dep_.CoverTab[195933]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:395
			// _ = "end of CoverTab[195933]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:395
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:395
		// _ = "end of CoverTab[195923]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:396
		_go_fuzz_dep_.CoverTab[195934]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:396
		// _ = "end of CoverTab[195934]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:396
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:396
	// _ = "end of CoverTab[195864]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:396
	_go_fuzz_dep_.CoverTab[195865]++

												result.incrementScore()
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:398
	// _ = "end of CoverTab[195865]"
}

func (v *subSchema) validateCommon(currentSubSchema *subSchema, value interface{}, result *Result, context *JsonContext) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:401
	_go_fuzz_dep_.CoverTab[195935]++

												if internalLogEnabled {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:403
		_go_fuzz_dep_.CoverTab[195939]++
													internalLog("validateCommon %s", context.String())
													internalLog(" %v", value)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:405
		// _ = "end of CoverTab[195939]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:406
		_go_fuzz_dep_.CoverTab[195940]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:406
		// _ = "end of CoverTab[195940]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:406
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:406
	// _ = "end of CoverTab[195935]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:406
	_go_fuzz_dep_.CoverTab[195936]++

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:409
	if currentSubSchema._const != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:409
		_go_fuzz_dep_.CoverTab[195941]++
													vString, err := marshalWithoutNumber(value)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:411
			_go_fuzz_dep_.CoverTab[195943]++
														result.addInternalError(new(InternalError), context, value, ErrorDetails{"error": err})
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:412
			// _ = "end of CoverTab[195943]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:413
			_go_fuzz_dep_.CoverTab[195944]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:413
			// _ = "end of CoverTab[195944]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:413
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:413
		// _ = "end of CoverTab[195941]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:413
		_go_fuzz_dep_.CoverTab[195942]++
													if *vString != *currentSubSchema._const {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:414
			_go_fuzz_dep_.CoverTab[195945]++
														result.addInternalError(new(ConstError),
				context,
				value,
				ErrorDetails{
					"allowed": *currentSubSchema._const,
				},
			)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:421
			// _ = "end of CoverTab[195945]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:422
			_go_fuzz_dep_.CoverTab[195946]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:422
			// _ = "end of CoverTab[195946]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:422
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:422
		// _ = "end of CoverTab[195942]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:423
		_go_fuzz_dep_.CoverTab[195947]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:423
		// _ = "end of CoverTab[195947]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:423
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:423
	// _ = "end of CoverTab[195936]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:423
	_go_fuzz_dep_.CoverTab[195937]++

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:426
	if len(currentSubSchema.enum) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:426
		_go_fuzz_dep_.CoverTab[195948]++
													vString, err := marshalWithoutNumber(value)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:428
			_go_fuzz_dep_.CoverTab[195950]++
														result.addInternalError(new(InternalError), context, value, ErrorDetails{"error": err})
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:429
			// _ = "end of CoverTab[195950]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:430
			_go_fuzz_dep_.CoverTab[195951]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:430
			// _ = "end of CoverTab[195951]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:430
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:430
		// _ = "end of CoverTab[195948]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:430
		_go_fuzz_dep_.CoverTab[195949]++
													if !isStringInSlice(currentSubSchema.enum, *vString) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:431
			_go_fuzz_dep_.CoverTab[195952]++
														result.addInternalError(
				new(EnumError),
				context,
				value,
				ErrorDetails{
					"allowed": strings.Join(currentSubSchema.enum, ", "),
				},
			)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:439
			// _ = "end of CoverTab[195952]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:440
			_go_fuzz_dep_.CoverTab[195953]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:440
			// _ = "end of CoverTab[195953]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:440
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:440
		// _ = "end of CoverTab[195949]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:441
		_go_fuzz_dep_.CoverTab[195954]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:441
		// _ = "end of CoverTab[195954]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:441
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:441
	// _ = "end of CoverTab[195937]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:441
	_go_fuzz_dep_.CoverTab[195938]++

												result.incrementScore()
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:443
	// _ = "end of CoverTab[195938]"
}

func (v *subSchema) validateArray(currentSubSchema *subSchema, value []interface{}, result *Result, context *JsonContext) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:446
	_go_fuzz_dep_.CoverTab[195955]++

												if internalLogEnabled {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:448
		_go_fuzz_dep_.CoverTab[195962]++
													internalLog("validateArray %s", context.String())
													internalLog(" %v", value)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:450
		// _ = "end of CoverTab[195962]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:451
		_go_fuzz_dep_.CoverTab[195963]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:451
		// _ = "end of CoverTab[195963]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:451
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:451
	// _ = "end of CoverTab[195955]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:451
	_go_fuzz_dep_.CoverTab[195956]++

												nbValues := len(value)

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:456
	if currentSubSchema.itemsChildrenIsSingleSchema {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:456
		_go_fuzz_dep_.CoverTab[195964]++
													for i := range value {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:457
			_go_fuzz_dep_.CoverTab[195965]++
														subContext := NewJsonContext(strconv.Itoa(i), context)
														validationResult := currentSubSchema.itemsChildren[0].subValidateWithContext(value[i], subContext)
														result.mergeErrors(validationResult)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:460
			// _ = "end of CoverTab[195965]"
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:461
		// _ = "end of CoverTab[195964]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:462
		_go_fuzz_dep_.CoverTab[195966]++
													if currentSubSchema.itemsChildren != nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:463
			_go_fuzz_dep_.CoverTab[195967]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:463
			return len(currentSubSchema.itemsChildren) > 0
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:463
			// _ = "end of CoverTab[195967]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:463
		}() {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:463
			_go_fuzz_dep_.CoverTab[195968]++

														nbItems := len(currentSubSchema.itemsChildren)

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:468
			for i := 0; i != nbItems && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:468
				_go_fuzz_dep_.CoverTab[195970]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:468
				return i != nbValues
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:468
				// _ = "end of CoverTab[195970]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:468
			}(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:468
				_go_fuzz_dep_.CoverTab[195971]++
															subContext := NewJsonContext(strconv.Itoa(i), context)
															validationResult := currentSubSchema.itemsChildren[i].subValidateWithContext(value[i], subContext)
															result.mergeErrors(validationResult)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:471
				// _ = "end of CoverTab[195971]"
			}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:472
			// _ = "end of CoverTab[195968]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:472
			_go_fuzz_dep_.CoverTab[195969]++

														if nbItems < nbValues {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:474
				_go_fuzz_dep_.CoverTab[195972]++

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:478
				switch currentSubSchema.additionalItems.(type) {
				case bool:
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:479
					_go_fuzz_dep_.CoverTab[195973]++
																if !currentSubSchema.additionalItems.(bool) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:480
						_go_fuzz_dep_.CoverTab[195975]++
																	result.addInternalError(new(ArrayNoAdditionalItemsError), context, value, ErrorDetails{})
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:481
						// _ = "end of CoverTab[195975]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:482
						_go_fuzz_dep_.CoverTab[195976]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:482
						// _ = "end of CoverTab[195976]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:482
					}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:482
					// _ = "end of CoverTab[195973]"
				case *subSchema:
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:483
					_go_fuzz_dep_.CoverTab[195974]++
																additionalItemSchema := currentSubSchema.additionalItems.(*subSchema)
																for i := nbItems; i != nbValues; i++ {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:485
						_go_fuzz_dep_.CoverTab[195977]++
																	subContext := NewJsonContext(strconv.Itoa(i), context)
																	validationResult := additionalItemSchema.subValidateWithContext(value[i], subContext)
																	result.mergeErrors(validationResult)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:488
						// _ = "end of CoverTab[195977]"
					}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:489
					// _ = "end of CoverTab[195974]"
				}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:490
				// _ = "end of CoverTab[195972]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:491
				_go_fuzz_dep_.CoverTab[195978]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:491
				// _ = "end of CoverTab[195978]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:491
			}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:491
			// _ = "end of CoverTab[195969]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:492
			_go_fuzz_dep_.CoverTab[195979]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:492
			// _ = "end of CoverTab[195979]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:492
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:492
		// _ = "end of CoverTab[195966]"
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:493
	// _ = "end of CoverTab[195956]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:493
	_go_fuzz_dep_.CoverTab[195957]++

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:496
	if currentSubSchema.minItems != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:496
		_go_fuzz_dep_.CoverTab[195980]++
													if nbValues < int(*currentSubSchema.minItems) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:497
			_go_fuzz_dep_.CoverTab[195981]++
														result.addInternalError(
				new(ArrayMinItemsError),
				context,
				value,
				ErrorDetails{"min": *currentSubSchema.minItems},
			)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:503
			// _ = "end of CoverTab[195981]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:504
			_go_fuzz_dep_.CoverTab[195982]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:504
			// _ = "end of CoverTab[195982]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:504
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:504
		// _ = "end of CoverTab[195980]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:505
		_go_fuzz_dep_.CoverTab[195983]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:505
		// _ = "end of CoverTab[195983]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:505
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:505
	// _ = "end of CoverTab[195957]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:505
	_go_fuzz_dep_.CoverTab[195958]++
												if currentSubSchema.maxItems != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:506
		_go_fuzz_dep_.CoverTab[195984]++
													if nbValues > int(*currentSubSchema.maxItems) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:507
			_go_fuzz_dep_.CoverTab[195985]++
														result.addInternalError(
				new(ArrayMaxItemsError),
				context,
				value,
				ErrorDetails{"max": *currentSubSchema.maxItems},
			)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:513
			// _ = "end of CoverTab[195985]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:514
			_go_fuzz_dep_.CoverTab[195986]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:514
			// _ = "end of CoverTab[195986]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:514
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:514
		// _ = "end of CoverTab[195984]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:515
		_go_fuzz_dep_.CoverTab[195987]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:515
		// _ = "end of CoverTab[195987]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:515
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:515
	// _ = "end of CoverTab[195958]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:515
	_go_fuzz_dep_.CoverTab[195959]++

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:518
	if currentSubSchema.uniqueItems {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:518
		_go_fuzz_dep_.CoverTab[195988]++
													var stringifiedItems = make(map[string]int)
													for j, v := range value {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:520
			_go_fuzz_dep_.CoverTab[195989]++
														vString, err := marshalWithoutNumber(v)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:522
				_go_fuzz_dep_.CoverTab[195992]++
															result.addInternalError(new(InternalError), context, value, ErrorDetails{"err": err})
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:523
				// _ = "end of CoverTab[195992]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:524
				_go_fuzz_dep_.CoverTab[195993]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:524
				// _ = "end of CoverTab[195993]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:524
			}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:524
			// _ = "end of CoverTab[195989]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:524
			_go_fuzz_dep_.CoverTab[195990]++
														if i, ok := stringifiedItems[*vString]; ok {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:525
				_go_fuzz_dep_.CoverTab[195994]++
															result.addInternalError(
					new(ItemsMustBeUniqueError),
					context,
					value,
					ErrorDetails{"type": TYPE_ARRAY, "i": i, "j": j},
				)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:531
				// _ = "end of CoverTab[195994]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:532
				_go_fuzz_dep_.CoverTab[195995]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:532
				// _ = "end of CoverTab[195995]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:532
			}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:532
			// _ = "end of CoverTab[195990]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:532
			_go_fuzz_dep_.CoverTab[195991]++
														stringifiedItems[*vString] = j
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:533
			// _ = "end of CoverTab[195991]"
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:534
		// _ = "end of CoverTab[195988]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:535
		_go_fuzz_dep_.CoverTab[195996]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:535
		// _ = "end of CoverTab[195996]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:535
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:535
	// _ = "end of CoverTab[195959]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:535
	_go_fuzz_dep_.CoverTab[195960]++

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:539
	if currentSubSchema.contains != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:539
		_go_fuzz_dep_.CoverTab[195997]++
													validatedOne := false
													var bestValidationResult *Result

													for i, v := range value {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:543
			_go_fuzz_dep_.CoverTab[195999]++
														subContext := NewJsonContext(strconv.Itoa(i), context)

														validationResult := currentSubSchema.contains.subValidateWithContext(v, subContext)
														if validationResult.Valid() {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:547
				_go_fuzz_dep_.CoverTab[196000]++
															validatedOne = true
															break
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:549
				// _ = "end of CoverTab[196000]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:550
				_go_fuzz_dep_.CoverTab[196001]++
															if bestValidationResult == nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:551
					_go_fuzz_dep_.CoverTab[196002]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:551
					return validationResult.score > bestValidationResult.score
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:551
					// _ = "end of CoverTab[196002]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:551
				}() {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:551
					_go_fuzz_dep_.CoverTab[196003]++
																bestValidationResult = validationResult
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:552
					// _ = "end of CoverTab[196003]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:553
					_go_fuzz_dep_.CoverTab[196004]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:553
					// _ = "end of CoverTab[196004]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:553
				}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:553
				// _ = "end of CoverTab[196001]"
			}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:554
			// _ = "end of CoverTab[195999]"
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:555
		// _ = "end of CoverTab[195997]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:555
		_go_fuzz_dep_.CoverTab[195998]++
													if !validatedOne {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:556
			_go_fuzz_dep_.CoverTab[196005]++
														result.addInternalError(
				new(ArrayContainsError),
				context,
				value,
				ErrorDetails{},
			)
			if bestValidationResult != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:563
				_go_fuzz_dep_.CoverTab[196006]++
															result.mergeErrors(bestValidationResult)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:564
				// _ = "end of CoverTab[196006]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:565
				_go_fuzz_dep_.CoverTab[196007]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:565
				// _ = "end of CoverTab[196007]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:565
			}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:565
			// _ = "end of CoverTab[196005]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:566
			_go_fuzz_dep_.CoverTab[196008]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:566
			// _ = "end of CoverTab[196008]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:566
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:566
		// _ = "end of CoverTab[195998]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:567
		_go_fuzz_dep_.CoverTab[196009]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:567
		// _ = "end of CoverTab[196009]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:567
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:567
	// _ = "end of CoverTab[195960]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:567
	_go_fuzz_dep_.CoverTab[195961]++

												result.incrementScore()
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:569
	// _ = "end of CoverTab[195961]"
}

func (v *subSchema) validateObject(currentSubSchema *subSchema, value map[string]interface{}, result *Result, context *JsonContext) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:572
	_go_fuzz_dep_.CoverTab[196010]++

												if internalLogEnabled {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:574
		_go_fuzz_dep_.CoverTab[196017]++
													internalLog("validateObject %s", context.String())
													internalLog(" %v", value)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:576
		// _ = "end of CoverTab[196017]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:577
		_go_fuzz_dep_.CoverTab[196018]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:577
		// _ = "end of CoverTab[196018]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:577
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:577
	// _ = "end of CoverTab[196010]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:577
	_go_fuzz_dep_.CoverTab[196011]++

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:580
	if currentSubSchema.minProperties != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:580
		_go_fuzz_dep_.CoverTab[196019]++
													if len(value) < int(*currentSubSchema.minProperties) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:581
			_go_fuzz_dep_.CoverTab[196020]++
														result.addInternalError(
				new(ArrayMinPropertiesError),
				context,
				value,
				ErrorDetails{"min": *currentSubSchema.minProperties},
			)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:587
			// _ = "end of CoverTab[196020]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:588
			_go_fuzz_dep_.CoverTab[196021]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:588
			// _ = "end of CoverTab[196021]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:588
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:588
		// _ = "end of CoverTab[196019]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:589
		_go_fuzz_dep_.CoverTab[196022]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:589
		// _ = "end of CoverTab[196022]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:589
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:589
	// _ = "end of CoverTab[196011]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:589
	_go_fuzz_dep_.CoverTab[196012]++
												if currentSubSchema.maxProperties != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:590
		_go_fuzz_dep_.CoverTab[196023]++
													if len(value) > int(*currentSubSchema.maxProperties) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:591
			_go_fuzz_dep_.CoverTab[196024]++
														result.addInternalError(
				new(ArrayMaxPropertiesError),
				context,
				value,
				ErrorDetails{"max": *currentSubSchema.maxProperties},
			)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:597
			// _ = "end of CoverTab[196024]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:598
			_go_fuzz_dep_.CoverTab[196025]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:598
			// _ = "end of CoverTab[196025]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:598
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:598
		// _ = "end of CoverTab[196023]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:599
		_go_fuzz_dep_.CoverTab[196026]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:599
		// _ = "end of CoverTab[196026]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:599
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:599
	// _ = "end of CoverTab[196012]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:599
	_go_fuzz_dep_.CoverTab[196013]++

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:602
	for _, requiredProperty := range currentSubSchema.required {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:602
		_go_fuzz_dep_.CoverTab[196027]++
													_, ok := value[requiredProperty]
													if ok {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:604
			_go_fuzz_dep_.CoverTab[196028]++
														result.incrementScore()
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:605
			// _ = "end of CoverTab[196028]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:606
			_go_fuzz_dep_.CoverTab[196029]++
														result.addInternalError(
				new(RequiredError),
				context,
				value,
				ErrorDetails{"property": requiredProperty},
			)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:612
			// _ = "end of CoverTab[196029]"
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:613
		// _ = "end of CoverTab[196027]"
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:614
	// _ = "end of CoverTab[196013]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:614
	_go_fuzz_dep_.CoverTab[196014]++

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:617
	for pk := range value {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:617
		_go_fuzz_dep_.CoverTab[196030]++

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:620
		found := false
		for _, spValue := range currentSubSchema.propertiesChildren {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:621
			_go_fuzz_dep_.CoverTab[196032]++
														if pk == spValue.property {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:622
				_go_fuzz_dep_.CoverTab[196033]++
															found = true
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:623
				// _ = "end of CoverTab[196033]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:624
				_go_fuzz_dep_.CoverTab[196034]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:624
				// _ = "end of CoverTab[196034]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:624
			}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:624
			// _ = "end of CoverTab[196032]"
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:625
		// _ = "end of CoverTab[196030]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:625
		_go_fuzz_dep_.CoverTab[196031]++

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:628
		ppMatch := v.validatePatternProperty(currentSubSchema, pk, value[pk], result, context)

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:631
		if !found && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:631
			_go_fuzz_dep_.CoverTab[196035]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:631
			return !ppMatch
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:631
			// _ = "end of CoverTab[196035]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:631
		}() {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:631
			_go_fuzz_dep_.CoverTab[196036]++
														switch ap := currentSubSchema.additionalProperties.(type) {
			case bool:
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:633
				_go_fuzz_dep_.CoverTab[196037]++

															if !ap {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:635
					_go_fuzz_dep_.CoverTab[196039]++
																result.addInternalError(
						new(AdditionalPropertyNotAllowedError),
						context,
						value[pk],
						ErrorDetails{"property": pk},
					)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:641
					// _ = "end of CoverTab[196039]"

				} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:643
					_go_fuzz_dep_.CoverTab[196040]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:643
					// _ = "end of CoverTab[196040]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:643
				}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:643
				// _ = "end of CoverTab[196037]"
			case *subSchema:
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:644
				_go_fuzz_dep_.CoverTab[196038]++
															validationResult := ap.subValidateWithContext(value[pk], NewJsonContext(pk, context))
															result.mergeErrors(validationResult)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:646
				// _ = "end of CoverTab[196038]"
			}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:647
			// _ = "end of CoverTab[196036]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:648
			_go_fuzz_dep_.CoverTab[196041]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:648
			// _ = "end of CoverTab[196041]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:648
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:648
		// _ = "end of CoverTab[196031]"
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:649
	// _ = "end of CoverTab[196014]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:649
	_go_fuzz_dep_.CoverTab[196015]++

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:652
	if currentSubSchema.propertyNames != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:652
		_go_fuzz_dep_.CoverTab[196042]++
													for pk := range value {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:653
			_go_fuzz_dep_.CoverTab[196043]++
														validationResult := currentSubSchema.propertyNames.subValidateWithContext(pk, context)
														if !validationResult.Valid() {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:655
				_go_fuzz_dep_.CoverTab[196044]++
															result.addInternalError(new(InvalidPropertyNameError),
					context,
					value, ErrorDetails{
						"property": pk,
					})
															result.mergeErrors(validationResult)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:661
				// _ = "end of CoverTab[196044]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:662
				_go_fuzz_dep_.CoverTab[196045]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:662
				// _ = "end of CoverTab[196045]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:662
			}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:662
			// _ = "end of CoverTab[196043]"
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:663
		// _ = "end of CoverTab[196042]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:664
		_go_fuzz_dep_.CoverTab[196046]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:664
		// _ = "end of CoverTab[196046]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:664
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:664
	// _ = "end of CoverTab[196015]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:664
	_go_fuzz_dep_.CoverTab[196016]++

												result.incrementScore()
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:666
	// _ = "end of CoverTab[196016]"
}

func (v *subSchema) validatePatternProperty(currentSubSchema *subSchema, key string, value interface{}, result *Result, context *JsonContext) bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:669
	_go_fuzz_dep_.CoverTab[196047]++

												if internalLogEnabled {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:671
		_go_fuzz_dep_.CoverTab[196051]++
													internalLog("validatePatternProperty %s", context.String())
													internalLog(" %s %v", key, value)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:673
		// _ = "end of CoverTab[196051]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:674
		_go_fuzz_dep_.CoverTab[196052]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:674
		// _ = "end of CoverTab[196052]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:674
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:674
	// _ = "end of CoverTab[196047]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:674
	_go_fuzz_dep_.CoverTab[196048]++

												validated := false

												for pk, pv := range currentSubSchema.patternProperties {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:678
		_go_fuzz_dep_.CoverTab[196053]++
													if matches, _ := regexp.MatchString(pk, key); matches {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:679
			_go_fuzz_dep_.CoverTab[196054]++
														validated = true
														subContext := NewJsonContext(key, context)
														validationResult := pv.subValidateWithContext(value, subContext)
														result.mergeErrors(validationResult)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:683
			// _ = "end of CoverTab[196054]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:684
			_go_fuzz_dep_.CoverTab[196055]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:684
			// _ = "end of CoverTab[196055]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:684
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:684
		// _ = "end of CoverTab[196053]"
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:685
	// _ = "end of CoverTab[196048]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:685
	_go_fuzz_dep_.CoverTab[196049]++

												if !validated {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:687
		_go_fuzz_dep_.CoverTab[196056]++
													return false
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:688
		// _ = "end of CoverTab[196056]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:689
		_go_fuzz_dep_.CoverTab[196057]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:689
		// _ = "end of CoverTab[196057]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:689
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:689
	// _ = "end of CoverTab[196049]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:689
	_go_fuzz_dep_.CoverTab[196050]++

												result.incrementScore()
												return true
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:692
	// _ = "end of CoverTab[196050]"
}

func (v *subSchema) validateString(currentSubSchema *subSchema, value interface{}, result *Result, context *JsonContext) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:695
	_go_fuzz_dep_.CoverTab[196058]++

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:698
	if isJSONNumber(value) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:698
		_go_fuzz_dep_.CoverTab[196066]++
													return
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:699
		// _ = "end of CoverTab[196066]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:700
		_go_fuzz_dep_.CoverTab[196067]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:700
		// _ = "end of CoverTab[196067]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:700
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:700
	// _ = "end of CoverTab[196058]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:700
	_go_fuzz_dep_.CoverTab[196059]++

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:703
	if !isKind(value, reflect.String) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:703
		_go_fuzz_dep_.CoverTab[196068]++
													return
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:704
		// _ = "end of CoverTab[196068]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:705
		_go_fuzz_dep_.CoverTab[196069]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:705
		// _ = "end of CoverTab[196069]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:705
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:705
	// _ = "end of CoverTab[196059]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:705
	_go_fuzz_dep_.CoverTab[196060]++

												if internalLogEnabled {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:707
		_go_fuzz_dep_.CoverTab[196070]++
													internalLog("validateString %s", context.String())
													internalLog(" %v", value)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:709
		// _ = "end of CoverTab[196070]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:710
		_go_fuzz_dep_.CoverTab[196071]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:710
		// _ = "end of CoverTab[196071]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:710
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:710
	// _ = "end of CoverTab[196060]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:710
	_go_fuzz_dep_.CoverTab[196061]++

												stringValue := value.(string)

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:715
	if currentSubSchema.minLength != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:715
		_go_fuzz_dep_.CoverTab[196072]++
													if utf8.RuneCount([]byte(stringValue)) < int(*currentSubSchema.minLength) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:716
			_go_fuzz_dep_.CoverTab[196073]++
														result.addInternalError(
				new(StringLengthGTEError),
				context,
				value,
				ErrorDetails{"min": *currentSubSchema.minLength},
			)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:722
			// _ = "end of CoverTab[196073]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:723
			_go_fuzz_dep_.CoverTab[196074]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:723
			// _ = "end of CoverTab[196074]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:723
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:723
		// _ = "end of CoverTab[196072]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:724
		_go_fuzz_dep_.CoverTab[196075]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:724
		// _ = "end of CoverTab[196075]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:724
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:724
	// _ = "end of CoverTab[196061]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:724
	_go_fuzz_dep_.CoverTab[196062]++
												if currentSubSchema.maxLength != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:725
		_go_fuzz_dep_.CoverTab[196076]++
													if utf8.RuneCount([]byte(stringValue)) > int(*currentSubSchema.maxLength) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:726
			_go_fuzz_dep_.CoverTab[196077]++
														result.addInternalError(
				new(StringLengthLTEError),
				context,
				value,
				ErrorDetails{"max": *currentSubSchema.maxLength},
			)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:732
			// _ = "end of CoverTab[196077]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:733
			_go_fuzz_dep_.CoverTab[196078]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:733
			// _ = "end of CoverTab[196078]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:733
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:733
		// _ = "end of CoverTab[196076]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:734
		_go_fuzz_dep_.CoverTab[196079]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:734
		// _ = "end of CoverTab[196079]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:734
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:734
	// _ = "end of CoverTab[196062]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:734
	_go_fuzz_dep_.CoverTab[196063]++

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:737
	if currentSubSchema.pattern != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:737
		_go_fuzz_dep_.CoverTab[196080]++
													if !currentSubSchema.pattern.MatchString(stringValue) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:738
			_go_fuzz_dep_.CoverTab[196081]++
														result.addInternalError(
				new(DoesNotMatchPatternError),
				context,
				value,
				ErrorDetails{"pattern": currentSubSchema.pattern},
			)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:744
			// _ = "end of CoverTab[196081]"

		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:746
			_go_fuzz_dep_.CoverTab[196082]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:746
			// _ = "end of CoverTab[196082]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:746
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:746
		// _ = "end of CoverTab[196080]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:747
		_go_fuzz_dep_.CoverTab[196083]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:747
		// _ = "end of CoverTab[196083]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:747
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:747
	// _ = "end of CoverTab[196063]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:747
	_go_fuzz_dep_.CoverTab[196064]++

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:750
	if currentSubSchema.format != "" {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:750
		_go_fuzz_dep_.CoverTab[196084]++
													if !FormatCheckers.IsFormat(currentSubSchema.format, stringValue) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:751
			_go_fuzz_dep_.CoverTab[196085]++
														result.addInternalError(
				new(DoesNotMatchFormatError),
				context,
				value,
				ErrorDetails{"format": currentSubSchema.format},
			)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:757
			// _ = "end of CoverTab[196085]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:758
			_go_fuzz_dep_.CoverTab[196086]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:758
			// _ = "end of CoverTab[196086]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:758
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:758
		// _ = "end of CoverTab[196084]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:759
		_go_fuzz_dep_.CoverTab[196087]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:759
		// _ = "end of CoverTab[196087]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:759
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:759
	// _ = "end of CoverTab[196064]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:759
	_go_fuzz_dep_.CoverTab[196065]++

												result.incrementScore()
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:761
	// _ = "end of CoverTab[196065]"
}

func (v *subSchema) validateNumber(currentSubSchema *subSchema, value interface{}, result *Result, context *JsonContext) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:764
	_go_fuzz_dep_.CoverTab[196088]++

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:767
	if !isJSONNumber(value) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:767
		_go_fuzz_dep_.CoverTab[196097]++
													return
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:768
		// _ = "end of CoverTab[196097]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:769
		_go_fuzz_dep_.CoverTab[196098]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:769
		// _ = "end of CoverTab[196098]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:769
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:769
	// _ = "end of CoverTab[196088]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:769
	_go_fuzz_dep_.CoverTab[196089]++

												if internalLogEnabled {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:771
		_go_fuzz_dep_.CoverTab[196099]++
													internalLog("validateNumber %s", context.String())
													internalLog(" %v", value)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:773
		// _ = "end of CoverTab[196099]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:774
		_go_fuzz_dep_.CoverTab[196100]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:774
		// _ = "end of CoverTab[196100]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:774
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:774
	// _ = "end of CoverTab[196089]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:774
	_go_fuzz_dep_.CoverTab[196090]++

												number := value.(json.Number)
												float64Value, _ := new(big.Rat).SetString(string(number))

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:780
	if currentSubSchema.multipleOf != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:780
		_go_fuzz_dep_.CoverTab[196101]++
													if q := new(big.Rat).Quo(float64Value, currentSubSchema.multipleOf); !q.IsInt() {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:781
			_go_fuzz_dep_.CoverTab[196102]++
														result.addInternalError(
				new(MultipleOfError),
				context,
				number,
				ErrorDetails{
					"multiple": new(big.Float).SetRat(currentSubSchema.multipleOf),
				},
			)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:789
			// _ = "end of CoverTab[196102]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:790
			_go_fuzz_dep_.CoverTab[196103]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:790
			// _ = "end of CoverTab[196103]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:790
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:790
		// _ = "end of CoverTab[196101]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:791
		_go_fuzz_dep_.CoverTab[196104]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:791
		// _ = "end of CoverTab[196104]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:791
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:791
	// _ = "end of CoverTab[196090]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:791
	_go_fuzz_dep_.CoverTab[196091]++

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:794
	if currentSubSchema.maximum != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:794
		_go_fuzz_dep_.CoverTab[196105]++
													if float64Value.Cmp(currentSubSchema.maximum) == 1 {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:795
			_go_fuzz_dep_.CoverTab[196106]++
														result.addInternalError(
				new(NumberLTEError),
				context,
				number,
				ErrorDetails{
					"max": new(big.Float).SetRat(currentSubSchema.maximum),
				},
			)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:803
			// _ = "end of CoverTab[196106]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:804
			_go_fuzz_dep_.CoverTab[196107]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:804
			// _ = "end of CoverTab[196107]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:804
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:804
		// _ = "end of CoverTab[196105]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:805
		_go_fuzz_dep_.CoverTab[196108]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:805
		// _ = "end of CoverTab[196108]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:805
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:805
	// _ = "end of CoverTab[196091]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:805
	_go_fuzz_dep_.CoverTab[196092]++
												if currentSubSchema.exclusiveMaximum != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:806
		_go_fuzz_dep_.CoverTab[196109]++
													if float64Value.Cmp(currentSubSchema.exclusiveMaximum) >= 0 {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:807
			_go_fuzz_dep_.CoverTab[196110]++
														result.addInternalError(
				new(NumberLTError),
				context,
				number,
				ErrorDetails{
					"max": new(big.Float).SetRat(currentSubSchema.exclusiveMaximum),
				},
			)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:815
			// _ = "end of CoverTab[196110]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:816
			_go_fuzz_dep_.CoverTab[196111]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:816
			// _ = "end of CoverTab[196111]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:816
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:816
		// _ = "end of CoverTab[196109]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:817
		_go_fuzz_dep_.CoverTab[196112]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:817
		// _ = "end of CoverTab[196112]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:817
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:817
	// _ = "end of CoverTab[196092]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:817
	_go_fuzz_dep_.CoverTab[196093]++

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:820
	if currentSubSchema.minimum != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:820
		_go_fuzz_dep_.CoverTab[196113]++
													if float64Value.Cmp(currentSubSchema.minimum) == -1 {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:821
			_go_fuzz_dep_.CoverTab[196114]++
														result.addInternalError(
				new(NumberGTEError),
				context,
				number,
				ErrorDetails{
					"min": new(big.Float).SetRat(currentSubSchema.minimum),
				},
			)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:829
			// _ = "end of CoverTab[196114]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:830
			_go_fuzz_dep_.CoverTab[196115]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:830
			// _ = "end of CoverTab[196115]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:830
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:830
		// _ = "end of CoverTab[196113]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:831
		_go_fuzz_dep_.CoverTab[196116]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:831
		// _ = "end of CoverTab[196116]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:831
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:831
	// _ = "end of CoverTab[196093]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:831
	_go_fuzz_dep_.CoverTab[196094]++
												if currentSubSchema.exclusiveMinimum != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:832
		_go_fuzz_dep_.CoverTab[196117]++
													if float64Value.Cmp(currentSubSchema.exclusiveMinimum) <= 0 {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:833
			_go_fuzz_dep_.CoverTab[196118]++
														result.addInternalError(
				new(NumberGTError),
				context,
				number,
				ErrorDetails{
					"min": new(big.Float).SetRat(currentSubSchema.exclusiveMinimum),
				},
			)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:841
			// _ = "end of CoverTab[196118]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:842
			_go_fuzz_dep_.CoverTab[196119]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:842
			// _ = "end of CoverTab[196119]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:842
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:842
		// _ = "end of CoverTab[196117]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:843
		_go_fuzz_dep_.CoverTab[196120]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:843
		// _ = "end of CoverTab[196120]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:843
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:843
	// _ = "end of CoverTab[196094]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:843
	_go_fuzz_dep_.CoverTab[196095]++

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:846
	if currentSubSchema.format != "" {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:846
		_go_fuzz_dep_.CoverTab[196121]++
													if !FormatCheckers.IsFormat(currentSubSchema.format, float64Value) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:847
			_go_fuzz_dep_.CoverTab[196122]++
														result.addInternalError(
				new(DoesNotMatchFormatError),
				context,
				value,
				ErrorDetails{"format": currentSubSchema.format},
			)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:853
			// _ = "end of CoverTab[196122]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:854
			_go_fuzz_dep_.CoverTab[196123]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:854
			// _ = "end of CoverTab[196123]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:854
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:854
		// _ = "end of CoverTab[196121]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:855
		_go_fuzz_dep_.CoverTab[196124]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:855
		// _ = "end of CoverTab[196124]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:855
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:855
	// _ = "end of CoverTab[196095]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:855
	_go_fuzz_dep_.CoverTab[196096]++

												result.incrementScore()
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:857
	// _ = "end of CoverTab[196096]"
}

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:858
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/validation.go:858
var _ = _go_fuzz_dep_.CoverTab
