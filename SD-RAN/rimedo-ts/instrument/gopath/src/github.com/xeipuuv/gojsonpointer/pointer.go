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

// author  			xeipuuv
// author-github 	https://github.com/xeipuuv
// author-mail		xeipuuv@gmail.com
//
// repository-name	gojsonpointer
// repository-desc	An implementation of JSON Pointer - Go language
//
// description		Main and unique file.
//
// created      	25-02-2013

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:26
package gojsonpointer

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:26
import (
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:26
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:26
)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:26
import (
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:26
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:26
)

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

const (
	const_empty_pointer	= ``
	const_pointer_separator	= `/`

	const_invalid_start	= `JSON pointer must be empty or start with a "` + const_pointer_separator + `"`
)

type implStruct struct {
	mode	string	// "SET" or "GET"

	inDocument	interface{}

	setInValue	interface{}

	getOutNode	interface{}
	getOutKind	reflect.Kind
	outError	error
}

type JsonPointer struct {
	referenceTokens []string
}

// NewJsonPointer parses the given string JSON pointer and returns an object
func NewJsonPointer(jsonPointerString string) (p JsonPointer, err error) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:60
	_go_fuzz_dep_.CoverTab[194653]++

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:63
	if len(jsonPointerString) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:63
		_go_fuzz_dep_.CoverTab[194656]++

																return
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:65
		// _ = "end of CoverTab[194656]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:66
		_go_fuzz_dep_.CoverTab[194657]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:66
		// _ = "end of CoverTab[194657]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:66
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:66
	// _ = "end of CoverTab[194653]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:66
	_go_fuzz_dep_.CoverTab[194654]++
															if jsonPointerString[0] != '/' {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:67
		_go_fuzz_dep_.CoverTab[194658]++
																return p, errors.New(const_invalid_start)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:68
		// _ = "end of CoverTab[194658]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:69
		_go_fuzz_dep_.CoverTab[194659]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:69
		// _ = "end of CoverTab[194659]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:69
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:69
	// _ = "end of CoverTab[194654]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:69
	_go_fuzz_dep_.CoverTab[194655]++

															p.referenceTokens = strings.Split(jsonPointerString[1:], const_pointer_separator)
															return
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:72
	// _ = "end of CoverTab[194655]"
}

// Uses the pointer to retrieve a value from a JSON document
func (p *JsonPointer) Get(document interface{}) (interface{}, reflect.Kind, error) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:76
	_go_fuzz_dep_.CoverTab[194660]++

															is := &implStruct{mode: "GET", inDocument: document}
															p.implementation(is)
															return is.getOutNode, is.getOutKind, is.outError
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:80
	// _ = "end of CoverTab[194660]"

}

// Uses the pointer to update a value from a JSON document
func (p *JsonPointer) Set(document interface{}, value interface{}) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:85
	_go_fuzz_dep_.CoverTab[194661]++

															is := &implStruct{mode: "SET", inDocument: document, setInValue: value}
															p.implementation(is)
															return document, is.outError
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:89
	// _ = "end of CoverTab[194661]"

}

// Uses the pointer to delete a value from a JSON document
func (p *JsonPointer) Delete(document interface{}) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:94
	_go_fuzz_dep_.CoverTab[194662]++
															is := &implStruct{mode: "DEL", inDocument: document}
															p.implementation(is)
															return document, is.outError
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:97
	// _ = "end of CoverTab[194662]"
}

// Both Get and Set functions use the same implementation to avoid code duplication
func (p *JsonPointer) implementation(i *implStruct) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:101
	_go_fuzz_dep_.CoverTab[194663]++

															kind := reflect.Invalid

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:106
	if len(p.referenceTokens) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:106
		_go_fuzz_dep_.CoverTab[194666]++
																i.getOutNode = i.inDocument
																i.outError = nil
																i.getOutKind = kind
																i.outError = nil
																return
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:111
		// _ = "end of CoverTab[194666]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:112
		_go_fuzz_dep_.CoverTab[194667]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:112
		// _ = "end of CoverTab[194667]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:112
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:112
	// _ = "end of CoverTab[194663]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:112
	_go_fuzz_dep_.CoverTab[194664]++

															node := i.inDocument

															previousNodes := make([]interface{}, len(p.referenceTokens))
															previousTokens := make([]string, len(p.referenceTokens))

															for ti, token := range p.referenceTokens {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:119
		_go_fuzz_dep_.CoverTab[194668]++

																isLastToken := ti == len(p.referenceTokens)-1
																previousNodes[ti] = node
																previousTokens[ti] = token

																switch v := node.(type) {

		case map[string]interface{}:
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:127
			_go_fuzz_dep_.CoverTab[194669]++
																	decodedToken := decodeReferenceToken(token)
																	if _, ok := v[decodedToken]; ok {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:129
				_go_fuzz_dep_.CoverTab[194674]++
																		node = v[decodedToken]
																		if isLastToken && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:131
					_go_fuzz_dep_.CoverTab[194675]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:131
					return i.mode == "SET"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:131
					// _ = "end of CoverTab[194675]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:131
				}() {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:131
					_go_fuzz_dep_.CoverTab[194676]++
																			v[decodedToken] = i.setInValue
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:132
					// _ = "end of CoverTab[194676]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:133
					_go_fuzz_dep_.CoverTab[194677]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:133
					if isLastToken && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:133
						_go_fuzz_dep_.CoverTab[194678]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:133
						return i.mode == "DEL"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:133
						// _ = "end of CoverTab[194678]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:133
					}() {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:133
						_go_fuzz_dep_.CoverTab[194679]++
																				delete(v, decodedToken)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:134
						// _ = "end of CoverTab[194679]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:135
						_go_fuzz_dep_.CoverTab[194680]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:135
						// _ = "end of CoverTab[194680]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:135
					}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:135
					// _ = "end of CoverTab[194677]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:135
				}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:135
				// _ = "end of CoverTab[194674]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:136
				_go_fuzz_dep_.CoverTab[194681]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:136
				if isLastToken && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:136
					_go_fuzz_dep_.CoverTab[194682]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:136
					return i.mode == "SET"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:136
					// _ = "end of CoverTab[194682]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:136
				}() {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:136
					_go_fuzz_dep_.CoverTab[194683]++
																			v[decodedToken] = i.setInValue
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:137
					// _ = "end of CoverTab[194683]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:138
					_go_fuzz_dep_.CoverTab[194684]++
																			i.outError = fmt.Errorf("Object has no key '%s'", decodedToken)
																			i.getOutKind = reflect.Map
																			i.getOutNode = nil
																			return
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:142
					// _ = "end of CoverTab[194684]"
				}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:143
				// _ = "end of CoverTab[194681]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:143
			}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:143
			// _ = "end of CoverTab[194669]"

		case []interface{}:
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:145
			_go_fuzz_dep_.CoverTab[194670]++
																	tokenIndex, err := strconv.Atoi(token)
																	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:147
				_go_fuzz_dep_.CoverTab[194685]++
																		i.outError = fmt.Errorf("Invalid array index '%s'", token)
																		i.getOutKind = reflect.Slice
																		i.getOutNode = nil
																		return
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:151
				// _ = "end of CoverTab[194685]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:152
				_go_fuzz_dep_.CoverTab[194686]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:152
				// _ = "end of CoverTab[194686]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:152
			}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:152
			// _ = "end of CoverTab[194670]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:152
			_go_fuzz_dep_.CoverTab[194671]++
																	if tokenIndex < 0 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:153
				_go_fuzz_dep_.CoverTab[194687]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:153
				return tokenIndex >= len(v)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:153
				// _ = "end of CoverTab[194687]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:153
			}() {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:153
				_go_fuzz_dep_.CoverTab[194688]++
																		i.outError = fmt.Errorf("Out of bound array[0,%d] index '%d'", len(v), tokenIndex)
																		i.getOutKind = reflect.Slice
																		i.getOutNode = nil
																		return
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:157
				// _ = "end of CoverTab[194688]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:158
				_go_fuzz_dep_.CoverTab[194689]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:158
				// _ = "end of CoverTab[194689]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:158
			}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:158
			// _ = "end of CoverTab[194671]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:158
			_go_fuzz_dep_.CoverTab[194672]++

																	node = v[tokenIndex]
																	if isLastToken && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:161
				_go_fuzz_dep_.CoverTab[194690]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:161
				return i.mode == "SET"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:161
				// _ = "end of CoverTab[194690]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:161
			}() {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:161
				_go_fuzz_dep_.CoverTab[194691]++
																		v[tokenIndex] = i.setInValue
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:162
				// _ = "end of CoverTab[194691]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:163
				_go_fuzz_dep_.CoverTab[194692]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:163
				if isLastToken && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:163
					_go_fuzz_dep_.CoverTab[194693]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:163
					return i.mode == "DEL"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:163
					// _ = "end of CoverTab[194693]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:163
				}() {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:163
					_go_fuzz_dep_.CoverTab[194694]++
																			v[tokenIndex] = v[len(v)-1]
																			v[len(v)-1] = nil
																			v = v[:len(v)-1]
																			previousNodes[ti-1].(map[string]interface{})[previousTokens[ti-1]] = v
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:167
					// _ = "end of CoverTab[194694]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:168
					_go_fuzz_dep_.CoverTab[194695]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:168
					// _ = "end of CoverTab[194695]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:168
				}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:168
				// _ = "end of CoverTab[194692]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:168
			}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:168
			// _ = "end of CoverTab[194672]"

		default:
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:170
			_go_fuzz_dep_.CoverTab[194673]++
																	i.outError = fmt.Errorf("Invalid token reference '%s'", token)
																	i.getOutKind = reflect.ValueOf(node).Kind()
																	i.getOutNode = nil
																	return
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:174
			// _ = "end of CoverTab[194673]"
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:175
		// _ = "end of CoverTab[194668]"

	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:177
	// _ = "end of CoverTab[194664]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:177
	_go_fuzz_dep_.CoverTab[194665]++

															i.getOutNode = node
															i.getOutKind = reflect.ValueOf(node).Kind()
															i.outError = nil
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:181
	// _ = "end of CoverTab[194665]"
}

// Pointer to string representation function
func (p *JsonPointer) String() string {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:185
	_go_fuzz_dep_.CoverTab[194696]++

															if len(p.referenceTokens) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:187
		_go_fuzz_dep_.CoverTab[194698]++
																return const_empty_pointer
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:188
		// _ = "end of CoverTab[194698]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:189
		_go_fuzz_dep_.CoverTab[194699]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:189
		// _ = "end of CoverTab[194699]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:189
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:189
	// _ = "end of CoverTab[194696]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:189
	_go_fuzz_dep_.CoverTab[194697]++

															pointerString := const_pointer_separator + strings.Join(p.referenceTokens, const_pointer_separator)

															return pointerString
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:193
	// _ = "end of CoverTab[194697]"
}

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:201
func decodeReferenceToken(token string) string {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:201
	_go_fuzz_dep_.CoverTab[194700]++
															step1 := strings.Replace(token, `~1`, `/`, -1)
															step2 := strings.Replace(step1, `~0`, `~`, -1)
															return step2
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:204
	// _ = "end of CoverTab[194700]"
}

func encodeReferenceToken(token string) string {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:207
	_go_fuzz_dep_.CoverTab[194701]++
															step1 := strings.Replace(token, `~`, `~0`, -1)
															step2 := strings.Replace(step1, `/`, `~1`, -1)
															return step2
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:210
	// _ = "end of CoverTab[194701]"
}

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:211
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonpointer@v0.0.0-20180127040702-4e3ac2762d5f/pointer.go:211
var _ = _go_fuzz_dep_.CoverTab
