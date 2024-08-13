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
// description		Defines resources pooling.
//                  Eases referencing and avoids downloading the same resource twice.
//
// created          26-02-2013

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:27
package gojsonschema

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:27
import (
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:27
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:27
)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:27
import (
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:27
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:27
)

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/xeipuuv/gojsonreference"
)

type schemaPoolDocument struct {
	Document	interface{}
	Draft		*Draft
}

type schemaPool struct {
	schemaPoolDocuments	map[string]*schemaPoolDocument
	jsonLoaderFactory	JSONLoaderFactory
	autoDetect		*bool
}

func (p *schemaPool) parseReferences(document interface{}, ref gojsonreference.JsonReference, pooled bool) error {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:48
	_go_fuzz_dep_.CoverTab[195612]++

												var (
		draft		*Draft
		err		error
		reference	= ref.String()
	)

	if _, ok := p.schemaPoolDocuments[reference]; pooled && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:56
		_go_fuzz_dep_.CoverTab[195616]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:56
		return ok
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:56
		// _ = "end of CoverTab[195616]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:56
	}() {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:56
		_go_fuzz_dep_.CoverTab[195617]++
													return fmt.Errorf("Reference already exists: \"%s\"", reference)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:57
		// _ = "end of CoverTab[195617]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:58
		_go_fuzz_dep_.CoverTab[195618]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:58
		// _ = "end of CoverTab[195618]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:58
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:58
	// _ = "end of CoverTab[195612]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:58
	_go_fuzz_dep_.CoverTab[195613]++

												if *p.autoDetect {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:60
		_go_fuzz_dep_.CoverTab[195619]++
													_, draft, err = parseSchemaURL(document)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:62
			_go_fuzz_dep_.CoverTab[195620]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:63
			// _ = "end of CoverTab[195620]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:64
			_go_fuzz_dep_.CoverTab[195621]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:64
			// _ = "end of CoverTab[195621]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:64
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:64
		// _ = "end of CoverTab[195619]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:65
		_go_fuzz_dep_.CoverTab[195622]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:65
		// _ = "end of CoverTab[195622]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:65
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:65
	// _ = "end of CoverTab[195613]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:65
	_go_fuzz_dep_.CoverTab[195614]++

												err = p.parseReferencesRecursive(document, ref, draft)

												if pooled {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:69
		_go_fuzz_dep_.CoverTab[195623]++
													p.schemaPoolDocuments[reference] = &schemaPoolDocument{Document: document, Draft: draft}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:70
		// _ = "end of CoverTab[195623]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:71
		_go_fuzz_dep_.CoverTab[195624]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:71
		// _ = "end of CoverTab[195624]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:71
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:71
	// _ = "end of CoverTab[195614]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:71
	_go_fuzz_dep_.CoverTab[195615]++

												return err
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:73
	// _ = "end of CoverTab[195615]"
}

func (p *schemaPool) parseReferencesRecursive(document interface{}, ref gojsonreference.JsonReference, draft *Draft) error {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:76
	_go_fuzz_dep_.CoverTab[195625]++

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:83
	switch m := document.(type) {
	case []interface{}:
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:84
		_go_fuzz_dep_.CoverTab[195627]++
													for _, v := range m {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:85
			_go_fuzz_dep_.CoverTab[195632]++
														p.parseReferencesRecursive(v, ref, draft)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:86
			// _ = "end of CoverTab[195632]"
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:87
		// _ = "end of CoverTab[195627]"
	case map[string]interface{}:
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:88
		_go_fuzz_dep_.CoverTab[195628]++
													localRef := &ref

													keyID := KEY_ID_NEW
													if existsMapKey(m, KEY_ID) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:92
			_go_fuzz_dep_.CoverTab[195633]++
														keyID = KEY_ID
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:93
			// _ = "end of CoverTab[195633]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:94
			_go_fuzz_dep_.CoverTab[195634]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:94
			// _ = "end of CoverTab[195634]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:94
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:94
		// _ = "end of CoverTab[195628]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:94
		_go_fuzz_dep_.CoverTab[195629]++
													if existsMapKey(m, keyID) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:95
			_go_fuzz_dep_.CoverTab[195635]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:95
			return isKind(m[keyID], reflect.String)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:95
			// _ = "end of CoverTab[195635]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:95
		}() {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:95
			_go_fuzz_dep_.CoverTab[195636]++
														jsonReference, err := gojsonreference.NewJsonReference(m[keyID].(string))
														if err == nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:97
				_go_fuzz_dep_.CoverTab[195637]++
															localRef, err = ref.Inherits(jsonReference)
															if err == nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:99
					_go_fuzz_dep_.CoverTab[195638]++
																if _, ok := p.schemaPoolDocuments[localRef.String()]; ok {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:100
						_go_fuzz_dep_.CoverTab[195640]++
																	return fmt.Errorf("Reference already exists: \"%s\"", localRef.String())
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:101
						// _ = "end of CoverTab[195640]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:102
						_go_fuzz_dep_.CoverTab[195641]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:102
						// _ = "end of CoverTab[195641]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:102
					}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:102
					// _ = "end of CoverTab[195638]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:102
					_go_fuzz_dep_.CoverTab[195639]++
																p.schemaPoolDocuments[localRef.String()] = &schemaPoolDocument{Document: document, Draft: draft}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:103
					// _ = "end of CoverTab[195639]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:104
					_go_fuzz_dep_.CoverTab[195642]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:104
					// _ = "end of CoverTab[195642]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:104
				}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:104
				// _ = "end of CoverTab[195637]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:105
				_go_fuzz_dep_.CoverTab[195643]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:105
				// _ = "end of CoverTab[195643]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:105
			}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:105
			// _ = "end of CoverTab[195636]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:106
			_go_fuzz_dep_.CoverTab[195644]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:106
			// _ = "end of CoverTab[195644]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:106
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:106
		// _ = "end of CoverTab[195629]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:106
		_go_fuzz_dep_.CoverTab[195630]++

													if existsMapKey(m, KEY_REF) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:108
			_go_fuzz_dep_.CoverTab[195645]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:108
			return isKind(m[KEY_REF], reflect.String)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:108
			// _ = "end of CoverTab[195645]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:108
		}() {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:108
			_go_fuzz_dep_.CoverTab[195646]++
														jsonReference, err := gojsonreference.NewJsonReference(m[KEY_REF].(string))
														if err == nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:110
				_go_fuzz_dep_.CoverTab[195647]++
															absoluteRef, err := localRef.Inherits(jsonReference)
															if err == nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:112
					_go_fuzz_dep_.CoverTab[195648]++
																m[KEY_REF] = absoluteRef.String()
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:113
					// _ = "end of CoverTab[195648]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:114
					_go_fuzz_dep_.CoverTab[195649]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:114
					// _ = "end of CoverTab[195649]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:114
				}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:114
				// _ = "end of CoverTab[195647]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:115
				_go_fuzz_dep_.CoverTab[195650]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:115
				// _ = "end of CoverTab[195650]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:115
			}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:115
			// _ = "end of CoverTab[195646]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:116
			_go_fuzz_dep_.CoverTab[195651]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:116
			// _ = "end of CoverTab[195651]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:116
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:116
		// _ = "end of CoverTab[195630]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:116
		_go_fuzz_dep_.CoverTab[195631]++

													for k, v := range m {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:118
			_go_fuzz_dep_.CoverTab[195652]++

														if k == KEY_CONST || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:120
				_go_fuzz_dep_.CoverTab[195654]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:120
				return k == KEY_ENUM
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:120
				// _ = "end of CoverTab[195654]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:120
			}() {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:120
				_go_fuzz_dep_.CoverTab[195655]++
															continue
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:121
				// _ = "end of CoverTab[195655]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:122
				_go_fuzz_dep_.CoverTab[195656]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:122
				// _ = "end of CoverTab[195656]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:122
			}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:122
			// _ = "end of CoverTab[195652]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:122
			_go_fuzz_dep_.CoverTab[195653]++

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:125
			if k == KEY_PROPERTIES || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:125
				_go_fuzz_dep_.CoverTab[195657]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:125
				return k == KEY_DEPENDENCIES
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:125
				// _ = "end of CoverTab[195657]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:125
			}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:125
				_go_fuzz_dep_.CoverTab[195658]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:125
				return k == KEY_PATTERN_PROPERTIES
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:125
				// _ = "end of CoverTab[195658]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:125
			}() {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:125
				_go_fuzz_dep_.CoverTab[195659]++
															if child, ok := v.(map[string]interface{}); ok {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:126
					_go_fuzz_dep_.CoverTab[195660]++
																for _, v := range child {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:127
						_go_fuzz_dep_.CoverTab[195661]++
																	p.parseReferencesRecursive(v, *localRef, draft)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:128
						// _ = "end of CoverTab[195661]"
					}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:129
					// _ = "end of CoverTab[195660]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:130
					_go_fuzz_dep_.CoverTab[195662]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:130
					// _ = "end of CoverTab[195662]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:130
				}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:130
				// _ = "end of CoverTab[195659]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:131
				_go_fuzz_dep_.CoverTab[195663]++
															p.parseReferencesRecursive(v, *localRef, draft)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:132
				// _ = "end of CoverTab[195663]"
			}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:133
			// _ = "end of CoverTab[195653]"
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:134
		// _ = "end of CoverTab[195631]"
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:135
	// _ = "end of CoverTab[195625]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:135
	_go_fuzz_dep_.CoverTab[195626]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:136
	// _ = "end of CoverTab[195626]"
}

func (p *schemaPool) GetDocument(reference gojsonreference.JsonReference) (*schemaPoolDocument, error) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:139
	_go_fuzz_dep_.CoverTab[195664]++

												var (
		spd	*schemaPoolDocument
		draft	*Draft
		ok	bool
		err	error
	)

	if internalLogEnabled {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:148
		_go_fuzz_dep_.CoverTab[195671]++
													internalLog("Get Document ( %s )", reference.String())
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:149
		// _ = "end of CoverTab[195671]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:150
		_go_fuzz_dep_.CoverTab[195672]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:150
		// _ = "end of CoverTab[195672]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:150
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:150
	// _ = "end of CoverTab[195664]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:150
	_go_fuzz_dep_.CoverTab[195665]++

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:153
	refToURL, _ := gojsonreference.NewJsonReference(reference.String())

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:158
	if spd, ok = p.schemaPoolDocuments[refToURL.String()]; ok {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:158
		_go_fuzz_dep_.CoverTab[195673]++
													if internalLogEnabled {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:159
			_go_fuzz_dep_.CoverTab[195675]++
														internalLog(" From pool")
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:160
			// _ = "end of CoverTab[195675]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:161
			_go_fuzz_dep_.CoverTab[195676]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:161
			// _ = "end of CoverTab[195676]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:161
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:161
		// _ = "end of CoverTab[195673]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:161
		_go_fuzz_dep_.CoverTab[195674]++
													return spd, nil
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:162
		// _ = "end of CoverTab[195674]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:163
		_go_fuzz_dep_.CoverTab[195677]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:163
		// _ = "end of CoverTab[195677]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:163
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:163
	// _ = "end of CoverTab[195665]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:163
	_go_fuzz_dep_.CoverTab[195666]++

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:168
	refToURL.GetUrl().Fragment = ""

	if cachedSpd, ok := p.schemaPoolDocuments[refToURL.String()]; ok {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:170
		_go_fuzz_dep_.CoverTab[195678]++
													document, _, err := reference.GetPointer().Get(cachedSpd.Document)

													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:173
			_go_fuzz_dep_.CoverTab[195681]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:174
			// _ = "end of CoverTab[195681]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:175
			_go_fuzz_dep_.CoverTab[195682]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:175
			// _ = "end of CoverTab[195682]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:175
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:175
		// _ = "end of CoverTab[195678]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:175
		_go_fuzz_dep_.CoverTab[195679]++

													if internalLogEnabled {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:177
			_go_fuzz_dep_.CoverTab[195683]++
														internalLog(" From pool")
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:178
			// _ = "end of CoverTab[195683]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:179
			_go_fuzz_dep_.CoverTab[195684]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:179
			// _ = "end of CoverTab[195684]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:179
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:179
		// _ = "end of CoverTab[195679]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:179
		_go_fuzz_dep_.CoverTab[195680]++

													spd = &schemaPoolDocument{Document: document, Draft: cachedSpd.Draft}
													p.schemaPoolDocuments[reference.String()] = spd

													return spd, nil
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:184
		// _ = "end of CoverTab[195680]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:185
		_go_fuzz_dep_.CoverTab[195685]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:185
		// _ = "end of CoverTab[195685]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:185
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:185
	// _ = "end of CoverTab[195666]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:185
	_go_fuzz_dep_.CoverTab[195667]++

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:188
	if !reference.IsCanonical() {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:188
		_go_fuzz_dep_.CoverTab[195686]++
													return nil, errors.New(formatErrorDescription(
			Locale.ReferenceMustBeCanonical(),
			ErrorDetails{"reference": reference.String()},
		))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:192
		// _ = "end of CoverTab[195686]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:193
		_go_fuzz_dep_.CoverTab[195687]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:193
		// _ = "end of CoverTab[195687]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:193
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:193
	// _ = "end of CoverTab[195667]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:193
	_go_fuzz_dep_.CoverTab[195668]++

												jsonReferenceLoader := p.jsonLoaderFactory.New(reference.String())
												document, err := jsonReferenceLoader.LoadJSON()

												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:198
		_go_fuzz_dep_.CoverTab[195688]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:199
		// _ = "end of CoverTab[195688]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:200
		_go_fuzz_dep_.CoverTab[195689]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:200
		// _ = "end of CoverTab[195689]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:200
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:200
	// _ = "end of CoverTab[195668]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:200
	_go_fuzz_dep_.CoverTab[195669]++

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:203
	p.parseReferences(document, refToURL, true)

												_, draft, _ = parseSchemaURL(document)

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:208
	document, _, err = reference.GetPointer().Get(document)

	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:210
		_go_fuzz_dep_.CoverTab[195690]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:211
		// _ = "end of CoverTab[195690]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:212
		_go_fuzz_dep_.CoverTab[195691]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:212
		// _ = "end of CoverTab[195691]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:212
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:212
	// _ = "end of CoverTab[195669]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:212
	_go_fuzz_dep_.CoverTab[195670]++

												return &schemaPoolDocument{Document: document, Draft: draft}, nil
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:214
	// _ = "end of CoverTab[195670]"
}

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:215
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaPool.go:215
var _ = _go_fuzz_dep_.CoverTab
