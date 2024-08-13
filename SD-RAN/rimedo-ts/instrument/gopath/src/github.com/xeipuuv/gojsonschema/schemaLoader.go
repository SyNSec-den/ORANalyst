// Copyright 2018 johandorland ( https://github.com/johandorland )
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

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:15
package gojsonschema

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:15
import (
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:15
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:15
)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:15
import (
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:15
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:15
)

import (
	"bytes"
	"errors"

	"github.com/xeipuuv/gojsonreference"
)

// SchemaLoader is used to load schemas
type SchemaLoader struct {
	pool		*schemaPool
	AutoDetect	bool
	Validate	bool
	Draft		Draft
}

// NewSchemaLoader creates a new NewSchemaLoader
func NewSchemaLoader() *SchemaLoader {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:33
	_go_fuzz_dep_.CoverTab[195535]++

												ps := &SchemaLoader{
		pool: &schemaPool{
			schemaPoolDocuments: make(map[string]*schemaPoolDocument),
		},
		AutoDetect:	true,
		Validate:	false,
		Draft:		Hybrid,
	}
												ps.pool.autoDetect = &ps.AutoDetect

												return ps
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:45
	// _ = "end of CoverTab[195535]"
}

func (sl *SchemaLoader) validateMetaschema(documentNode interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:48
	_go_fuzz_dep_.CoverTab[195536]++

												var (
		schema	string
		err	error
	)
	if sl.AutoDetect {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:54
		_go_fuzz_dep_.CoverTab[195541]++
													schema, _, err = parseSchemaURL(documentNode)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:56
			_go_fuzz_dep_.CoverTab[195542]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:57
			// _ = "end of CoverTab[195542]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:58
			_go_fuzz_dep_.CoverTab[195543]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:58
			// _ = "end of CoverTab[195543]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:58
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:58
		// _ = "end of CoverTab[195541]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:59
		_go_fuzz_dep_.CoverTab[195544]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:59
		// _ = "end of CoverTab[195544]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:59
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:59
	// _ = "end of CoverTab[195536]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:59
	_go_fuzz_dep_.CoverTab[195537]++

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:62
	if schema == "" {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:62
		_go_fuzz_dep_.CoverTab[195545]++
													if sl.Draft == Hybrid {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:63
			_go_fuzz_dep_.CoverTab[195547]++
														return nil
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:64
			// _ = "end of CoverTab[195547]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:65
			_go_fuzz_dep_.CoverTab[195548]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:65
			// _ = "end of CoverTab[195548]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:65
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:65
		// _ = "end of CoverTab[195545]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:65
		_go_fuzz_dep_.CoverTab[195546]++
													schema = drafts.GetSchemaURL(sl.Draft)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:66
		// _ = "end of CoverTab[195546]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:67
		_go_fuzz_dep_.CoverTab[195549]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:67
		// _ = "end of CoverTab[195549]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:67
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:67
	// _ = "end of CoverTab[195537]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:67
	_go_fuzz_dep_.CoverTab[195538]++

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:70
	sl.Validate = false

	metaSchema, err := sl.Compile(NewReferenceLoader(schema))

	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:74
		_go_fuzz_dep_.CoverTab[195550]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:75
		// _ = "end of CoverTab[195550]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:76
		_go_fuzz_dep_.CoverTab[195551]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:76
		// _ = "end of CoverTab[195551]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:76
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:76
	// _ = "end of CoverTab[195538]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:76
	_go_fuzz_dep_.CoverTab[195539]++

												sl.Validate = true

												result := metaSchema.validateDocument(documentNode)

												if !result.Valid() {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:82
		_go_fuzz_dep_.CoverTab[195552]++
													var res bytes.Buffer
													for _, err := range result.Errors() {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:84
			_go_fuzz_dep_.CoverTab[195554]++
														res.WriteString(err.String())
														res.WriteString("\n")
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:86
			// _ = "end of CoverTab[195554]"
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:87
		// _ = "end of CoverTab[195552]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:87
		_go_fuzz_dep_.CoverTab[195553]++
													return errors.New(res.String())
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:88
		// _ = "end of CoverTab[195553]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:89
		_go_fuzz_dep_.CoverTab[195555]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:89
		// _ = "end of CoverTab[195555]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:89
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:89
	// _ = "end of CoverTab[195539]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:89
	_go_fuzz_dep_.CoverTab[195540]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:91
	// _ = "end of CoverTab[195540]"
}

// AddSchemas adds an arbritrary amount of schemas to the schema cache. As this function does not require
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:94
// an explicit URL, every schema should contain an $id, so that it can be referenced by the main schema
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:96
func (sl *SchemaLoader) AddSchemas(loaders ...JSONLoader) error {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:96
	_go_fuzz_dep_.CoverTab[195556]++
												emptyRef, _ := gojsonreference.NewJsonReference("")

												for _, loader := range loaders {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:99
		_go_fuzz_dep_.CoverTab[195558]++
													doc, err := loader.LoadJSON()

													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:102
			_go_fuzz_dep_.CoverTab[195561]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:103
			// _ = "end of CoverTab[195561]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:104
			_go_fuzz_dep_.CoverTab[195562]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:104
			// _ = "end of CoverTab[195562]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:104
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:104
		// _ = "end of CoverTab[195558]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:104
		_go_fuzz_dep_.CoverTab[195559]++

													if sl.Validate {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:106
			_go_fuzz_dep_.CoverTab[195563]++
														if err := sl.validateMetaschema(doc); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:107
				_go_fuzz_dep_.CoverTab[195564]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:108
				// _ = "end of CoverTab[195564]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:109
				_go_fuzz_dep_.CoverTab[195565]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:109
				// _ = "end of CoverTab[195565]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:109
			}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:109
			// _ = "end of CoverTab[195563]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:110
			_go_fuzz_dep_.CoverTab[195566]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:110
			// _ = "end of CoverTab[195566]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:110
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:110
		// _ = "end of CoverTab[195559]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:110
		_go_fuzz_dep_.CoverTab[195560]++

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:114
		if err = sl.pool.parseReferences(doc, emptyRef, false); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:114
			_go_fuzz_dep_.CoverTab[195567]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:115
			// _ = "end of CoverTab[195567]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:116
			_go_fuzz_dep_.CoverTab[195568]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:116
			// _ = "end of CoverTab[195568]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:116
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:116
		// _ = "end of CoverTab[195560]"
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:117
	// _ = "end of CoverTab[195556]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:117
	_go_fuzz_dep_.CoverTab[195557]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:119
	// _ = "end of CoverTab[195557]"
}

// AddSchema adds a schema under the provided URL to the schema cache
func (sl *SchemaLoader) AddSchema(url string, loader JSONLoader) error {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:123
	_go_fuzz_dep_.CoverTab[195569]++

												ref, err := gojsonreference.NewJsonReference(url)

												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:127
		_go_fuzz_dep_.CoverTab[195573]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:128
		// _ = "end of CoverTab[195573]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:129
		_go_fuzz_dep_.CoverTab[195574]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:129
		// _ = "end of CoverTab[195574]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:129
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:129
	// _ = "end of CoverTab[195569]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:129
	_go_fuzz_dep_.CoverTab[195570]++

												doc, err := loader.LoadJSON()

												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:133
		_go_fuzz_dep_.CoverTab[195575]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:134
		// _ = "end of CoverTab[195575]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:135
		_go_fuzz_dep_.CoverTab[195576]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:135
		// _ = "end of CoverTab[195576]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:135
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:135
	// _ = "end of CoverTab[195570]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:135
	_go_fuzz_dep_.CoverTab[195571]++

												if sl.Validate {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:137
		_go_fuzz_dep_.CoverTab[195577]++
													if err := sl.validateMetaschema(doc); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:138
			_go_fuzz_dep_.CoverTab[195578]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:139
			// _ = "end of CoverTab[195578]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:140
			_go_fuzz_dep_.CoverTab[195579]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:140
			// _ = "end of CoverTab[195579]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:140
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:140
		// _ = "end of CoverTab[195577]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:141
		_go_fuzz_dep_.CoverTab[195580]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:141
		// _ = "end of CoverTab[195580]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:141
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:141
	// _ = "end of CoverTab[195571]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:141
	_go_fuzz_dep_.CoverTab[195572]++

												return sl.pool.parseReferences(doc, ref, true)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:143
	// _ = "end of CoverTab[195572]"
}

// Compile loads and compiles a schema
func (sl *SchemaLoader) Compile(rootSchema JSONLoader) (*Schema, error) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:147
	_go_fuzz_dep_.CoverTab[195581]++

												ref, err := rootSchema.JsonReference()

												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:151
		_go_fuzz_dep_.CoverTab[195587]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:152
		// _ = "end of CoverTab[195587]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:153
		_go_fuzz_dep_.CoverTab[195588]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:153
		// _ = "end of CoverTab[195588]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:153
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:153
	// _ = "end of CoverTab[195581]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:153
	_go_fuzz_dep_.CoverTab[195582]++

												d := Schema{}
												d.pool = sl.pool
												d.pool.jsonLoaderFactory = rootSchema.LoaderFactory()
												d.documentReference = ref
												d.referencePool = newSchemaReferencePool()

												var doc interface{}
												if ref.String() != "" {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:162
		_go_fuzz_dep_.CoverTab[195589]++

													spd, err := d.pool.GetDocument(d.documentReference)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:165
			_go_fuzz_dep_.CoverTab[195591]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:166
			// _ = "end of CoverTab[195591]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:167
			_go_fuzz_dep_.CoverTab[195592]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:167
			// _ = "end of CoverTab[195592]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:167
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:167
		// _ = "end of CoverTab[195589]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:167
		_go_fuzz_dep_.CoverTab[195590]++
													doc = spd.Document
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:168
		// _ = "end of CoverTab[195590]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:169
		_go_fuzz_dep_.CoverTab[195593]++

													doc, err = rootSchema.LoadJSON()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:172
			_go_fuzz_dep_.CoverTab[195595]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:173
			// _ = "end of CoverTab[195595]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:174
			_go_fuzz_dep_.CoverTab[195596]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:174
			// _ = "end of CoverTab[195596]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:174
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:174
		// _ = "end of CoverTab[195593]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:174
		_go_fuzz_dep_.CoverTab[195594]++

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:177
		err = sl.pool.parseReferences(doc, ref, true)
		if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:178
			_go_fuzz_dep_.CoverTab[195597]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:179
			// _ = "end of CoverTab[195597]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:180
			_go_fuzz_dep_.CoverTab[195598]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:180
			// _ = "end of CoverTab[195598]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:180
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:180
		// _ = "end of CoverTab[195594]"
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:181
	// _ = "end of CoverTab[195582]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:181
	_go_fuzz_dep_.CoverTab[195583]++

												if sl.Validate {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:183
		_go_fuzz_dep_.CoverTab[195599]++
													if err := sl.validateMetaschema(doc); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:184
			_go_fuzz_dep_.CoverTab[195600]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:185
			// _ = "end of CoverTab[195600]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:186
			_go_fuzz_dep_.CoverTab[195601]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:186
			// _ = "end of CoverTab[195601]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:186
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:186
		// _ = "end of CoverTab[195599]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:187
		_go_fuzz_dep_.CoverTab[195602]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:187
		// _ = "end of CoverTab[195602]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:187
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:187
	// _ = "end of CoverTab[195583]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:187
	_go_fuzz_dep_.CoverTab[195584]++

												draft := sl.Draft
												if sl.AutoDetect {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:190
		_go_fuzz_dep_.CoverTab[195603]++
													_, detectedDraft, err := parseSchemaURL(doc)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:192
			_go_fuzz_dep_.CoverTab[195605]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:193
			// _ = "end of CoverTab[195605]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:194
			_go_fuzz_dep_.CoverTab[195606]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:194
			// _ = "end of CoverTab[195606]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:194
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:194
		// _ = "end of CoverTab[195603]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:194
		_go_fuzz_dep_.CoverTab[195604]++
													if detectedDraft != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:195
			_go_fuzz_dep_.CoverTab[195607]++
														draft = *detectedDraft
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:196
			// _ = "end of CoverTab[195607]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:197
			_go_fuzz_dep_.CoverTab[195608]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:197
			// _ = "end of CoverTab[195608]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:197
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:197
		// _ = "end of CoverTab[195604]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:198
		_go_fuzz_dep_.CoverTab[195609]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:198
		// _ = "end of CoverTab[195609]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:198
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:198
	// _ = "end of CoverTab[195584]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:198
	_go_fuzz_dep_.CoverTab[195585]++

												err = d.parse(doc, draft)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:201
		_go_fuzz_dep_.CoverTab[195610]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:202
		// _ = "end of CoverTab[195610]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:203
		_go_fuzz_dep_.CoverTab[195611]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:203
		// _ = "end of CoverTab[195611]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:203
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:203
	// _ = "end of CoverTab[195585]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:203
	_go_fuzz_dep_.CoverTab[195586]++

												return &d, nil
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:205
	// _ = "end of CoverTab[195586]"
}

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:206
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaLoader.go:206
var _ = _go_fuzz_dep_.CoverTab
